// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"bufio"
	"bytes"
	"fmt"
	goscanner "go/scanner"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"

	"modernc.org/mathutil"
	"modernc.org/token"
)

const (
	clsEOF = iota + 0x80
	clsOther
)

const maxASCII = 0x7f

var (
	bom = []byte{0xEF, 0xBB, 0xBF}

	idDefine      = dict.sid("define")
	idElif        = dict.sid("elif")
	idElse        = dict.sid("else")
	idEndif       = dict.sid("endif")
	idError       = dict.sid("error")
	idIf          = dict.sid("if")
	idIfdef       = dict.sid("ifdef")
	idIfndef      = dict.sid("ifndef")
	idInclude     = dict.sid("include")
	idIncludeNext = dict.sid("include_next")
	idLine        = dict.sid("line")
	idPragma      = dict.sid("pragma")
	idSpace       = dict.sid(" ")
	idUndef       = dict.sid("undef")

	trigraphPrefix = []byte("??")
	trigraphs      = []struct{ from, to []byte }{
		{[]byte("??="), []byte{'#'}},
		{[]byte("??("), []byte{'['}},
		{[]byte("??/"), []byte{'\\'}},
		{[]byte("??)"), []byte{']'}},
		{[]byte("??'"), []byte{'^'}},
		{[]byte("??<"), []byte{'{'}},
		{[]byte("??!"), []byte{'|'}},
		{[]byte("??>"), []byte{'}'}},
		{[]byte("??-"), []byte{'~'}},
	}
)

type node interface {
	Pos() token.Pos
}

type dictionary struct {
	mu      sync.RWMutex
	m       map[string]StringID
	strings []string
}

func newDictionary() (r *dictionary) {
	r = &dictionary{m: map[string]StringID{}}
	b := make([]byte, 1)
	for i := 0; i < 128; i++ {
		var s string
		if i != 0 {
			b[0] = byte(i)
			s = string(b)
		}
		r.m[s] = StringID(i)
		r.strings = append(r.strings, s)
		dictStrings[i] = s
	}
	return r
}

func (d *dictionary) id(key []byte) StringID {
	switch len(key) {
	case 0:
		return 0
	case 1:
		if c := key[0]; c != 0 && c < 128 {
			return StringID(c)
		}
	}

	d.mu.Lock()
	if n, ok := d.m[string(key)]; ok {
		d.mu.Unlock()
		return n
	}

	n := StringID(len(d.strings))
	s := string(key)
	if int(n) < 256 {
		dictStrings[n] = s
	}
	d.strings = append(d.strings, s)
	d.m[s] = n
	d.mu.Unlock()
	return n
}

func (d *dictionary) sid(key string) StringID {
	switch len(key) {
	case 0:
		return 0
	case 1:
		if c := key[0]; c != 0 && c < 128 {
			return StringID(c)
		}
	}

	d.mu.Lock()
	if n, ok := d.m[key]; ok {
		d.mu.Unlock()
		return n
	}

	n := StringID(len(d.strings))
	if int(n) < 256 {
		dictStrings[n] = key
	}
	d.strings = append(d.strings, key)
	d.m[key] = n
	d.mu.Unlock()
	return n
}

type char struct {
	pos int32
	c   byte
}

// token3 is produced by translation phase 3.
type token3 struct {
	char  rune
	pos   int32
	value StringID
}

func (t token3) Pos() token.Pos { return token.Pos(t.pos) }
func (t token3) String() string { return t.value.String() }

type scanner struct {
	bomFix        int
	bytesBuf      []byte
	charBuf       []char
	ctx           *context
	file          *token.File
	fileOffset    int
	firstPos      token.Pos
	lineBuf       []byte
	lookaheadChar char
	lookaheadLine ppLine
	mark          int
	pos           token.Pos
	r             *bufio.Reader
	tokenBuf      []token3
	ungetBuf      []char

	tok token3

	closed             bool
	preserveWhiteSpace bool
}

func newScanner(ctx *context, r io.Reader, file *token.File) *scanner {
	bufSize := 1 << 17 // emulate gcc
	if n := ctx.cfg.MaxSourceLine; n > 4096 {
		bufSize = n
	}
	s := &scanner{
		ctx:  ctx,
		file: file,
		r:    bufio.NewReaderSize(r, bufSize),
	}
	if r != nil {
		s.init()
	}
	return s
}

func (s *scanner) abort() (r byte, b bool) {
	if s.mark >= 0 {
		if len(s.charBuf) > s.mark {
			s.unget(s.lookaheadChar)
			for i := len(s.charBuf) - 1; i >= s.mark; i-- {
				s.unget(s.charBuf[i])
			}
		}
		s.charBuf = s.charBuf[:s.mark]
		return 0, false
	}

	switch n := len(s.charBuf); n {
	case 0: // [] z
		c := s.lookaheadChar
		s.next()
		return s.class(c.c), true
	case 1: // [a] z
		return s.class(s.charBuf[0].c), true
	default: // [a, b, ...], z
		c := s.charBuf[0]        // a
		s.unget(s.lookaheadChar) // z
		for i := n - 1; i > 1; i-- {
			s.unget(s.charBuf[i]) // ...
		}
		s.lookaheadChar = s.charBuf[1] // b
		s.charBuf = s.charBuf[:1]
		return s.class(c.c), true
	}
}

func (s *scanner) class(b byte) byte {
	switch {
	case b == 0:
		return clsEOF
	case b > maxASCII:
		return clsOther
	default:
		return b
	}
}

func (s *scanner) err(n node, msg string, args ...interface{}) { s.errPos(n.Pos(), msg, args...) }

func (s *scanner) errLine(x interface{}, msg string, args ...interface{}) {
	var toks []token3
	switch x := x.(type) {
	case nil:
		toks = []token3{{}}
	case ppLine:
		toks = x.toks()
	default:
		panic(fmt.Sprintf("internal error %T", x)) //TODOOK
	}
	var b strings.Builder
	for _, v := range toks {
		switch v.char {
		case '\n':
			// nop
		case ' ':
			b.WriteByte(' ')
		default:
			b.WriteString(v.String())
		}
	}
	s.err(toks[0], "%s"+msg, append([]interface{}{b.String()}, args...)...)
}

func (s *scanner) errPos(pos token.Pos, msg string, args ...interface{}) {
	if s.ctx.err(s.file.Position(pos), msg, args...) {
		s.r.Reset(nil)
		s.closed = true
	}
}

func (s *scanner) init() *scanner {
	if s.r == nil {
		return s
	}

	b, err := s.r.Peek(3)
	if err == nil && bytes.Equal(b, bom) {
		s.bomFix, _ = s.r.Discard(3)
	}
	return s
}

func (s *scanner) initScan() (r byte) {
	if s.lookaheadChar.pos == 0 {
		s.next()
	}
	s.firstPos = token.Pos(s.lookaheadChar.pos)
	s.mark = -1
	if len(s.charBuf) > 1<<18 { //DONE benchmark tuned
		s.bytesBuf = nil
		s.charBuf = nil
	} else {
		s.bytesBuf = s.bytesBuf[:0]
		s.charBuf = s.charBuf[:0]
	}
	return s.class(s.lookaheadChar.c)
}

func (s *scanner) lex() {
	s.tok.char = s.scan()
	s.tok.pos = int32(s.firstPos)
	switch {
	case s.tok.char == ' ' && !s.preserveWhiteSpace && !s.ctx.cfg.PreserveWhiteSpace:
		s.tok.value = idSpace
	case s.tok.char == IDENTIFIER:
		for i := 0; i < len(s.charBuf); {
			c := s.charBuf[i].c
			if c != '\\' {
				s.bytesBuf = append(s.bytesBuf, c)
				i++
				continue
			}

			i++ // Skip '\\'
			var n int
			switch s.charBuf[i].c {
			case 'u':
				n = 4
			case 'U':
				n = 4
			default:
				panic("internal error") //TODOOK
			}
			i++ // Skip 'u' or 'U'
			l := len(s.bytesBuf)
			for i0 := i; i < i0+n; i++ {
				s.bytesBuf = append(s.bytesBuf, s.charBuf[i].c)
			}
			r, err := strconv.ParseUint(string(s.bytesBuf[l:l+n]), 16, 32)
			if err != nil {
				panic("internal error") //TODOOK
			}

			n2 := utf8.EncodeRune(s.bytesBuf[l:], rune(r))
			s.bytesBuf = s.bytesBuf[:l+n2]
		}
		s.tok.value = dict.id(s.bytesBuf)
	default:
		for _, v := range s.charBuf {
			s.bytesBuf = append(s.bytesBuf, v.c)
		}
		s.tok.value = dict.id(s.bytesBuf)
	}
	switch s.tok.char {
	case clsEOF:
		s.tok.char = -1
		s.tok.pos = int32(s.file.Pos(s.file.Size()))
	}
}

func (s *scanner) next() (r byte) {
	if s.lookaheadChar.pos > 0 {
		s.charBuf = append(s.charBuf, s.lookaheadChar)
	}
	if n := len(s.ungetBuf); n != 0 {
		s.lookaheadChar = s.ungetBuf[n-1]
		s.ungetBuf = s.ungetBuf[:n-1]
		return s.class(s.lookaheadChar.c)
	}

	if len(s.lineBuf) == 0 {
	more:
		if s.closed || s.fileOffset == s.file.Size() {
			s.lookaheadChar.c = 0
			s.lookaheadChar.pos = 0
			return clsEOF
		}

		b, err := s.r.ReadSlice('\n')
		if err != nil {
			if err != io.EOF {
				s.errPos(s.pos, "error while reading %s: %s", s.file.Name(), err)
			}
			if len(b) == 0 {
				return clsEOF
			}
		}

		s.file.AddLine(s.fileOffset)
		s.fileOffset += s.bomFix
		s.bomFix = 0
		s.pos = token.Pos(s.fileOffset)
		s.fileOffset += len(b)

		// [0], 5.1.1.2, 1.1
		//
		// Physical source file multibyte characters are mapped, in an
		// implementation- defined manner, to the source character set
		// (introducing new-line characters for end-of-line indicators)
		// if necessary. Trigraph sequences are replaced by
		// corresponding single-character internal representations.
		if bytes.Contains(b, trigraphPrefix) {
			for _, v := range trigraphs {
				b = bytes.Replace(b, v.from, v.to, -1)
			}
		}

		// [0], 5.1.1.2, 2
		//
		// Each instance of a backslash character (\) immediately
		// followed by a new-line character is deleted, splicing
		// physical source lines to form logical source lines.  Only
		// the last backslash on any physical source line shall be
		// eligible for being part of such a splice. A source file that
		// is not empty shall end in a new-line character, which shall
		// not be immediately preceded by a backslash character before
		// any such splicing takes place.
		s.lineBuf = b
		n := len(b)
		switch {
		case b[n-1] != '\n':
			if s.ctx.cfg.RejectMissingFinalNewline {
				s.errPos(s.pos+token.Pos(n), "non empty source file shall end in a new-line character")
			}
			b = append(b[:n:n], '\n') // bufio.Reader owns the bytes
		case n > 1 && b[n-2] == '\\':
			if n == 2 {
				goto more
			}

			b = b[:n-2]
			n = len(b)
			if s.fileOffset == s.file.Size() {
				if s.ctx.cfg.RejectFinalBackslash {
					s.errPos(s.pos+token.Pos(n+1), "source file final new-line character shall not be preceded by a backslash character")
				}
				b = append(b[:n:n], '\n') // bufio.Reader owns the bytes
			}
		}
		s.lineBuf = b
	}
	s.pos++
	s.lookaheadChar = char{int32(s.pos), s.lineBuf[0]}
	s.lineBuf = s.lineBuf[1:]
	return s.class(s.lookaheadChar.c)
}

func (s *scanner) unget(c ...char) {
	s.ungetBuf = append(s.ungetBuf, c...)
	s.lookaheadChar.pos = 0 // Must invalidate lookahead.
}

func (s *scanner) unterminatedComment() rune {
	s.errPos(token.Pos(s.file.Size()), "unterminated comment")
	n := len(s.charBuf)
	s.unget(s.charBuf[n-1]) // \n
	s.charBuf = s.charBuf[:n-1]
	return ' '
}

// -------------------------------------------------------- Translation phase 3

// [0], 5.1.1.2, 3
//
// The source file is decomposed into preprocessing tokens and sequences of
// white-space characters (including comments). A source file shall not end in
// a partial preprocessing token or in a partial comment. Each comment is
// replaced by one space character. New-line characters are retained. Whether
// each nonempty sequence of white-space characters other than new-line is
// retained or replaced by one space character is implementation-defined.
func (s *scanner) translationPhase3() *ppFile {
	r := &ppFile{file: s.file}
	if s.file.Size() == 0 {
		s.r.Reset(nil)
		return r
	}

	s.nextLine()
	r.groups = s.parseGroup()
	return r
}

func (s *scanner) parseGroup() (r []ppGroup) {
	for {
		switch x := s.lookaheadLine.(type) {
		case ppGroup:
			r = append(r, x)
			s.nextLine()
		case ppIfGroupDirective:
			r = append(r, s.parseIfSection())
		default:
			return r
		}
	}
}

func (s *scanner) parseIfSection() *ppIfSection {
	return &ppIfSection{
		ifGroup:    s.parseIfGroup(),
		elifGroups: s.parseElifGroup(),
		elseGroup:  s.parseElseGroup(),
		endifLine:  s.parseEndifLine(),
	}
}

func (s *scanner) parseEndifLine() *ppEndifDirective {
	switch x := s.lookaheadLine.(type) {
	case *ppEndifDirective:
		s.nextLine()
		return x
	default:
		s.errLine(x, fmt.Sprintf(": expected #endif (unexpected %T)", x))
		s.nextLine()
		return nil
	}
}

func (s *scanner) parseElseGroup() *ppElseGroup {
	switch x := s.lookaheadLine.(type) {
	case *ppElseDirective:
		r := &ppElseGroup{elseLine: x}
		s.nextLine()
		r.groups = s.parseGroup()
		return r
	default:
		return nil
	}
}

func (s *scanner) parseElifGroup() (r []*ppElifGroup) {
	for {
		var g ppElifGroup
		switch x := s.lookaheadLine.(type) {
		case *ppElifDirective:
			g.elif = x
			s.nextLine()
			g.groups = s.parseGroup()
			r = append(r, &g)
		default:
			return r
		}
	}
}

func (s *scanner) parseIfGroup() *ppIfGroup {
	r := &ppIfGroup{}
	switch x := s.lookaheadLine.(type) {
	case ppIfGroupDirective:
		r.directive = x
	default:
		s.errLine(x, fmt.Sprintf(": expected if-group (unexpected %T)", x))
	}
	s.nextLine()
	r.groups = s.parseGroup()
	return r
}

func (s *scanner) nextLine() {
	s.lookaheadLine = s.scanLine()
}

func (s *scanner) scanLine() (r ppLine) {
	toks := s.scanToNonBlankToken(nil)
	if len(toks) == 0 {
		return nil
	}

	includeNext := false
	switch tok := toks[len(toks)-1]; tok.char {
	case '#':
		toks = s.scanToNonBlankToken(toks)
		switch tok := toks[len(toks)-1]; tok.char {
		case '\n':
			return &ppEmptyDirective{Toks: toks}
		case IDENTIFIER:
			switch tok.value {
			case idDefine:
				return s.parseDefine(toks)
			case idElif:
				return s.parseElif(toks)
			case idElse:
				return s.parseElse(toks)
			case idEndif:
				return s.parseEndif(toks)
			case idIf:
				return s.parseIf(toks)
			case idIfdef:
				return s.parseIfdef(toks)
			case idIfndef:
				return s.parseIfndef(toks)
			case idIncludeNext:
				includeNext = true
				fallthrough
			case idInclude:
				// # include pp-tokens new-line
				//
				// Prevent aliasing of eg. <foo  bar.h> and <foo bar.h>.
				save := s.preserveWhiteSpace
				s.preserveWhiteSpace = true
				n := len(toks)
				toks := s.scanLineToEOL(toks)
				r := &ppIncludeDirective{Arg: toks[n : len(toks)-1], Toks: toks, IncludeNext: includeNext}
				s.preserveWhiteSpace = save
				return r
			case idUndef:
				return s.parseUndef(toks)
			case idLine:
				return s.parseLine(toks)
			case idError:
				// # error pp-tokens_opt new-line
				n := len(toks)
				toks := s.scanLineToEOL(toks)
				msg := toks[n : len(toks)-1]
				if len(msg) != 0 && msg[0].char == ' ' {
					msg = msg[1:]
				}
				return &ppErrorDirective{Toks: toks, Msg: msg}
			case idPragma:
				// pragma pp-tokens_opt new-line
				return &ppPragmaDirective{Toks: s.scanLineToEOL(toks)}
			}
		}

		// # non-directive
		return &ppNonDirective{Toks: s.scanLineToEOL(toks)}
	case '\n':
		return &ppTextLine{Toks: toks}
	default:
		return &ppTextLine{Toks: s.scanLineToEOL(toks)}
	}
}

// # line pp-tokens new-line
func (s *scanner) parseLine(toks []token3) *ppLineDirective {
	toks = s.scanToNonBlankToken(toks)
	switch tok := toks[len(toks)-1]; tok.char {
	case '\n':
		s.err(tok, "unexpected new-line")
		return &ppLineDirective{Toks: toks}
	default:
		toks := s.scanLineToEOL(toks)
		r := &ppLineDirective{Toks: toks}
		toks = toks[:len(toks)-1] // sans new-line
		toks = ltrim(toks)
		toks = toks[1:] // Skip '#'
		toks = ltrim(toks)
		toks = toks[1:] // Skip "line"
		r.Args = ltrim(toks)
		return r
	}
}

func ltrim(toks []token3) []token3 {
	for i, v := range toks {
		if v.char != ' ' {
			return toks[i:]
		}
	}
	return toks
}

// # undef identifier new-line
func (s *scanner) parseUndef(toks []token3) *ppUndefDirective {
	toks = s.scanToNonBlankToken(toks)
	switch tok := toks[len(toks)-1]; tok.char {
	case '\n':
		s.err(&tok, "expected identifier")
		return &ppUndefDirective{Toks: toks}
	case IDENTIFIER:
		name := tok
		toks = s.scanToNonBlankToken(toks)
		switch tok := toks[len(toks)-1]; tok.char {
		case '\n':
			return &ppUndefDirective{Name: name, Toks: toks}
		default:
			if s.ctx.cfg.RejectUndefExtraTokens {
				s.err(&tok, "extra tokens after #undef")
			}
			return &ppUndefDirective{Name: name, Toks: s.scanLineToEOL(toks)}
		}
	default:
		s.err(&tok, "expected identifier")
		return &ppUndefDirective{Toks: s.scanLineToEOL(toks)}
	}
}

func (s *scanner) scanLineToEOL(toks []token3) []token3 {
	n := len(s.tokenBuf) - len(toks)
	for {
		s.lex()
		s.tokenBuf = append(s.tokenBuf, s.tok)
		if s.tok.char == '\n' {
			return s.tokenBuf[n:]
		}
	}
}

// # ifndef identifier new-line
func (s *scanner) parseIfndef(toks []token3) *ppIfndefDirective {
	var name StringID
	toks = s.scanToNonBlankToken(toks)
	switch tok := toks[len(toks)-1]; tok.char {
	case IDENTIFIER:
		name = tok.value
		toks = s.scanToNonBlankToken(toks)
		switch tok := toks[len(toks)-1]; tok.char {
		case '\n':
			return &ppIfndefDirective{Name: name, Toks: toks}
		default:
			if s.ctx.cfg.RejectIfndefExtraTokens {
				s.err(&tok, "extra tokens after #ifndef")
			}
			return &ppIfndefDirective{Name: name, Toks: s.scanLineToEOL(toks)}
		}
	case '\n':
		s.err(tok, "expected identifier")
		return &ppIfndefDirective{Name: name, Toks: toks}
	default:
		s.err(tok, "expected identifier")
		return &ppIfndefDirective{Name: name, Toks: s.scanLineToEOL(toks)}
	}
}

// # ifdef identifier new-line
func (s *scanner) parseIfdef(toks []token3) *ppIfdefDirective {
	var name StringID
	toks = s.scanToNonBlankToken(toks)
	switch tok := toks[len(toks)-1]; tok.char {
	case IDENTIFIER:
		name = tok.value
		toks = s.scanToNonBlankToken(toks)
		switch tok := toks[len(toks)-1]; tok.char {
		case '\n':
			return &ppIfdefDirective{Name: name, Toks: toks}
		default:
			if s.ctx.cfg.RejectIfdefExtraTokens {
				s.err(&tok, "extra tokens after #ifdef")
			}
			return &ppIfdefDirective{Name: name, Toks: s.scanLineToEOL(toks)}
		}
	case '\n':
		s.err(tok, "expected identifier")
		return &ppIfdefDirective{Name: name, Toks: toks}
	default:
		s.err(tok, "expected identifier")
		return &ppIfdefDirective{Name: name, Toks: s.scanLineToEOL(toks)}
	}
}

// # if constant-expression new-line
func (s *scanner) parseIf(toks []token3) *ppIfDirective {
	n := len(toks)
	toks = s.scanToNonBlankToken(toks)
	switch tok := toks[len(toks)-1]; tok.char {
	case '\n':
		s.err(tok, "expected expression")
		return &ppIfDirective{Toks: toks}
	default:
		toks = s.scanLineToEOL(toks)
		expr := toks[n:]
		if expr[0].char == ' ' { // sans leading space
			expr = expr[1:]
		}
		expr = expr[:len(expr)-1] // sans '\n'
		return &ppIfDirective{Toks: toks, Expr: expr}
	}
}

// # endif new-line
func (s *scanner) parseEndif(toks []token3) *ppEndifDirective {
	toks = s.scanToNonBlankToken(toks)
	switch tok := toks[len(toks)-1]; tok.char {
	case '\n':
		return &ppEndifDirective{toks}
	default:
		if s.ctx.cfg.RejectEndifExtraTokens {
			s.err(&tok, "extra tokens after #else")
		}
		return &ppEndifDirective{s.scanLineToEOL(toks)}
	}
}

// # else new-line
func (s *scanner) parseElse(toks []token3) *ppElseDirective {
	toks = s.scanToNonBlankToken(toks)
	switch tok := toks[len(toks)-1]; tok.char {
	case '\n':
		return &ppElseDirective{toks}
	default:
		if s.ctx.cfg.RejectElseExtraTokens {
			s.err(&tok, "extra tokens after #else")
		}
		return &ppElseDirective{s.scanLineToEOL(toks)}
	}
}

// # elif constant-expression new-line
func (s *scanner) parseElif(toks []token3) *ppElifDirective {
	n := len(toks)
	toks = s.scanToNonBlankToken(toks)
	switch tok := toks[len(toks)-1]; tok.char {
	case '\n':
		s.err(tok, "expected expression")
		return &ppElifDirective{toks, nil}
	default:
		toks = s.scanLineToEOL(toks)
		expr := toks[n:]
		if expr[0].char == ' ' { // sans leading space
			expr = expr[1:]
		}
		expr = expr[:len(expr)-1] // sans '\n'
		return &ppElifDirective{toks, expr}
	}
}

func (s *scanner) parseDefine(toks []token3) ppLine {
	toks = s.scanToNonBlankToken(toks)
	switch tok := toks[len(toks)-1]; tok.char {
	case IDENTIFIER:
		name := tok
		n := len(toks)
		toks = s.scanToNonBlankToken(toks)
		switch tok := toks[len(toks)-1]; tok.char {
		case '\n':
			return &ppDefineObjectMacroDirective{Name: name, Toks: toks}
		case '(':
			if toks[n].char == ' ' {
				return s.parseDefineObjectMacro(n, name, toks)
			}

			return s.parseDefineFunctionMacro(name, toks)
		default:
			return s.parseDefineObjectMacro(n, name, toks)
		}
	case '\n':
		s.err(tok, "expected identifier")
		return &ppDefineObjectMacroDirective{Toks: toks}
	default:
		s.err(tok, "expected identifier")
		return &ppDefineObjectMacroDirective{Toks: s.scanLineToEOL(toks)}
	}
}

// # define identifier lparen identifier-list_opt ) replacement-list new-line
// # define identifier lparen ... ) replacement-list new-line
// # define identifier lparen identifier-list , ... ) replacement-list new-line
func (s *scanner) parseDefineFunctionMacro(name token3, toks []token3) *ppDefineFunctionMacroDirective {
	// Parse parameters after "#define name(".
	var list []token3
	variadic := false
	namedVariadic := false
again:
	toks = s.scanToNonBlankToken(toks)
	switch tok := toks[len(toks)-1]; tok.char {
	case IDENTIFIER:
	more:
		list = append(list, tok)
		toks = s.scanToNonBlankToken(toks)
		switch tok = toks[len(toks)-1]; tok.char {
		case ',':
			toks = s.scanToNonBlankToken(toks)
			switch tok = toks[len(toks)-1]; tok.char {
			case IDENTIFIER:
				goto more
			case DDD:
				if toks, variadic = s.parseDDD(toks); !variadic {
					goto again
				}
			case ')':
				s.err(tok, "expected parameter name")
			default:
				s.err(tok, "unexpected %q", &tok)
			}
		case DDD:
			namedVariadic = true
			if s.ctx.cfg.RejectInvalidVariadicMacros {
				s.err(tok, "expected comma")
			}
			if toks, variadic = s.parseDDD(toks); !variadic {
				goto again
			}
		case ')':
			// ok
		case '\n':
			s.err(tok, "unexpected new-line")
			return &ppDefineFunctionMacroDirective{Toks: toks}
		case IDENTIFIER:
			s.err(tok, "expected comma")
			goto more
		default:
			s.err(tok, "unexpected %q", &tok)
		}
	case DDD:
		if toks, variadic = s.parseDDD(toks); !variadic {
			goto again
		}
	case ',':
		s.err(tok, "expected parameter name")
		goto again
	case ')':
		// ok
	default:
		s.err(tok, "expected parameter name")
		goto again
	}
	// Parse replacement list.
	n := len(toks)
	toks = s.scanToNonBlankToken(toks)
	switch tok := toks[len(toks)-1]; tok.char {
	case '\n':
		if s.ctx.cfg.RejectFunctionMacroEmptyReplacementList {
			s.err(tok, "expected replacement list")
		}
		return &ppDefineFunctionMacroDirective{Name: name, IdentifierList: list, Toks: toks, Variadic: variadic, NamedVariadic: namedVariadic}
	default:
		toks = s.scanLineToEOL(toks)
		repl := toks[n:] // sans #define identifier
		if repl[0].char == ' ' {
			repl = repl[1:]
		}
		repl = repl[:len(repl)-1] // sans '\n'
		return &ppDefineFunctionMacroDirective{Name: name, IdentifierList: list, Toks: toks, ReplacementList: repl, Variadic: variadic, NamedVariadic: namedVariadic}
	}
}

func (s *scanner) parseDDD(toks []token3) ([]token3, bool) {
	toks = s.scanToNonBlankToken(toks)
	switch tok := toks[len(toks)-1]; tok.char {
	case ')':
		return toks, true
	default:
		s.err(tok, "expected right parenthesis")
		return toks, false
	}
}

// # define identifier replacement-list new-line
func (s *scanner) parseDefineObjectMacro(n int, name token3, toks []token3) *ppDefineObjectMacroDirective {
	toks = s.scanLineToEOL(toks)
	repl := toks[n:] // sans #define identifier
	if repl[0].char == ' ' {
		repl = repl[1:]
	}
	repl = repl[:len(repl)-1] // sans '\n'
	return &ppDefineObjectMacroDirective{Name: name, Toks: toks, ReplacementList: repl}
}

// Return {}, {x} or {' ', x}
func (s *scanner) scanToNonBlankToken(toks []token3) []token3 {
	n := len(s.tokenBuf) - len(toks)
	for {
		s.lex()
		if s.tok.char < 0 {
			return s.tokenBuf[n:]
		}

		s.tokenBuf = append(s.tokenBuf, s.tok)
		if s.tok.char != ' ' {
			return s.tokenBuf[n:]
		}
	}
}

// ---------------------------------------------------------------------- Cache

// Translation phase4 source.
type source interface {
	ppFile() (*ppFile, error)
	fileID() int32
}

type cachedPPFile struct {
	err     error
	errs    goscanner.ErrorList
	modTime int64 // time.Time.UnixNano()
	pf      *ppFile
	readyCh chan struct{}
	size    int

	id int32
}

func (c *cachedPPFile) fileID() int32                   { return c.id }
func (c *cachedPPFile) ready() *cachedPPFile            { close(c.readyCh); return c }
func (c *cachedPPFile) waitFor() (*cachedPPFile, error) { <-c.readyCh; return c, c.err }

func (c *cachedPPFile) ppFile() (*ppFile, error) {
	c.waitFor()
	if c.err == nil {
		return c.pf, nil
	}

	return nil, c.err
}

type cacheKey struct {
	name  StringID
	value StringID
	Config3
}

type ppCache struct {
	mu sync.RWMutex
	f  map[int32]*token.File
	m  map[cacheKey]*cachedPPFile

	fileID int32
}

func newPPCache() *ppCache {
	return &ppCache{
		f: map[int32]*token.File{},
		m: map[cacheKey]*cachedPPFile{},
	}
}

func (c *ppCache) allocID(tf *token.File) (r int32) {
	c.mu.Lock()
	c.fileID++
	r = c.fileID
	c.f[r] = tf
	c.mu.Unlock()
	return r
}

func (c *ppCache) file(id int32) *token.File {
	c.mu.RLock()
	r := c.f[id]
	c.mu.RUnlock()
	return r
}

func (c *ppCache) get(ctx *context, src Source) (source, error) {
	if src.Value != "" {
		return c.getValue(ctx, src.Name, src.Value)
	}

	return c.getFile(ctx, src.Name)
}

func (c *ppCache) getFile(ctx *context, name string) (*cachedPPFile, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return nil, err
	}

	if !fi.Mode().IsRegular() {
		return nil, fmt.Errorf("%s is not a regular file", name)
	}

	if fi.Size() > mathutil.MaxInt {
		return nil, fmt.Errorf("%s: file too big", name)
	}

	size := int(fi.Size())
	if !filepath.IsAbs(name) { // Never cache relative paths
		if isTesting {
			panic("internal error") //TODOOK
		}

		f, err := os.Open(name)
		if err != nil {
			return nil, err
		}

		defer f.Close()

		tf := token.NewFile(name, size)
		ppFile := newScanner(ctx, f, tf).translationPhase3()
		cf := &cachedPPFile{pf: ppFile, id: c.allocID(tf), readyCh: make(chan struct{})}
		c.allocID(tf)
		cf.ready()
		return cf, nil
	}

	modTime := fi.ModTime().UnixNano()
	key := cacheKey{dict.sid(name), 0, ctx.cfg.Config3}
	c.mu.Lock()
	if cf, ok := c.m[key]; ok {
		if modTime <= cf.modTime && size == cf.size {
			c.mu.Unlock()
			if cf.err != nil {
				return nil, cf.err
			}

			r, err := cf.waitFor()
			ctx.errs(cf.errs)
			return r, err
		}

		delete(c.f, cf.id)
		delete(c.m, key)
	}

	c.fileID++
	tf := token.NewFile(name, size)
	c.f[c.fileID] = tf
	cf := &cachedPPFile{modTime: modTime, size: size, id: c.fileID, readyCh: make(chan struct{})}
	c.m[key] = cf
	c.mu.Unlock()

	go func() {
		defer cf.ready()

		f, err := os.Open(name)
		if err != nil {
			cf.err = err
			return
		}

		defer f.Close()

		ctx2 := newContext(ctx.cfg)
		cf.pf = newScanner(ctx2, f, tf).translationPhase3()
		cf.errs = ctx2.ErrorList
		ctx.errs(cf.errs)
	}()

	return cf.waitFor()
}

func (c *ppCache) getValue(ctx *context, name, value string) (*cachedPPFile, error) {
	key := cacheKey{dict.sid(name), dict.sid(value), ctx.cfg.Config3}
	c.mu.Lock()
	if cf, ok := c.m[key]; ok {
		c.mu.Unlock()
		if cf.err != nil {
			return nil, cf.err
		}

		r, err := cf.waitFor()
		ctx.errs(cf.errs)
		return r, err
	}

	c.fileID++
	fileID := c.fileID
	tf := token.NewFile(name, len(value))
	c.f[fileID] = tf
	cf := &cachedPPFile{id: fileID, readyCh: make(chan struct{})}
	c.m[key] = cf
	c.mu.Unlock()
	ctx2 := newContext(ctx.cfg)
	cf.pf = newScanner(ctx2, strings.NewReader(value), tf).translationPhase3()
	cf.errs = ctx2.ErrorList
	ctx.errs(cf.errs)
	cf.ready()
	return cf.waitFor()
}
