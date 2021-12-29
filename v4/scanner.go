// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"bytes"
	"io"
	"io/fs"
	"math"
	"os"
	"unicode"
	"unicode/utf8"

	"modernc.org/token"
)

// Static type asserts.
var (
	_ Node = (*Token)(nil)
)

// Rune range used for IDENTIFIER and similar non-characters.
const (
	eof                     = -1
	unicodePrivateAreaFirst = 0xe000
	unicodePrivateAreaLast  = 0xf8ff
)

// errHandler is a function called on error.
type errHandler func(msg string, args ...interface{})

// tokCh enables using stringer.
type tokCh int

// Values of Token.Rune representing non-character categories.
const (
	_ tokCh = iota + unicodePrivateAreaFirst - 1

	ACCUM                  // _Accum
	ADDASSIGN              // +=
	ALIGNAS                // _Alignas
	ALIGNOF                // _Alignof
	ANDAND                 // &&
	ANDASSIGN              // &=
	ARROW                  // ->
	ASM                    // __asm__
	ATOMIC                 // _Atomic
	ATTRIBUTE              // __attribute__
	AUTO                   // auto
	BOOL                   // _Bool
	BREAK                  // break
	BUILTINCHOOSEEXPR      // __builtin_choose_expr
	BUILTINTYPESCOMPATIBLE // __builtin_types_compatible_p
	CASE                   // case
	CHAR                   // char
	CHARCONST              // 'a'
	COMPLEX                // _Complex
	CONST                  // const
	CONTINUE               // continue
	DDD                    // ...
	DEC                    // --
	DECIMAL128             // _Decimal128
	DECIMAL32              // _Decimal32
	DECIMAL64              // _Decimal64
	DEFAULT                // default
	DIVASSIGN              // /=
	DO                     // do
	DOUBLE                 // double
	ELSE                   // else
	ENUM                   // enum
	ENUMCONST              // foo in enum x { foo, bar };
	EQ                     // ==
	EXTERN                 // extern
	FLOAT                  // float
	FLOAT128               // _Float128
	FLOAT16                // __fp16
	FLOAT32                // _Float32
	FLOAT32X               // _Float32x
	FLOAT64                // _Float64
	FLOAT64X               // _Float64x
	FLOAT80                // __float80
	FLOATCONST             // 1.23
	FOR                    // for
	FRACT                  // _Fract
	GEQ                    // >=
	GOTO                   // goto
	HEADER_NAME            // <foo> or "bar"
	IDENTIFIER             // foo
	IF                     // if
	IMAG                   // __imag__
	INC                    // ++
	INLINE                 // inline
	INT                    // int
	INT128                 // __int128
	INT16                  // __int16
	INT32                  // __int32
	INT64                  // __int64
	INT8                   // __int8
	INTCONST               // 42
	LABEL                  // __label__
	LEQ                    // <=
	LONG                   // long
	LONGCHARCONST          // L'a'
	LONGSTRINGLITERAL      // L"foo"
	LSH                    // <<
	LSHASSIGN              // <<=
	MODASSIGN              // %=
	MULASSIGN              // *=
	NEQ                    // !=
	NORETURN               // _Noreturn
	ORASSIGN               // |=
	OROR                   // ||
	PPNUMBER               // .32e.
	PPPASTE                // ##
	PRAGMASTDC             // __pragma_stdc
	REAL                   // __real__
	REGISTER               // register
	RESTRICT               // restrict
	RETURN                 // return
	RSH                    // >>
	RSHASSIGN              // >>=
	SAT                    // _Sat
	SHORT                  // short
	SIGNED                 // signed
	SIZEOF                 // sizeof
	STATIC                 // static
	STRINGLITERAL          // "foo"
	STRUCT                 // struct
	SUBASSIGN              // -=
	SWITCH                 // switch
	THREADLOCAL            // _Thread_local
	TYPEDEF                // typedef
	TYPENAME               // int_t in typedef int int_t;
	TYPEOF                 // typeof
	UNION                  // union
	UNSIGNED               // unsigned
	VOID                   // void
	VOLATILE               // volatile
	WHILE                  // while
	XORASSIGN              // ^=
)

// Node is implemented by Token and AST nodes.
type Node interface {
	Position() token.Position
}

// Token is the lexical token produced by the scanner.
type Token struct { // 32 bytes on a 64 bit machine.
	s   *scannerSource
	Ch  rune   // '*' or IDENTIFIER etc.
	pos uint32 // Index into ss.buf of the original token.
	seq uint32 // Sequence number, determines scope boundaries.
	sep uint32 // Index into .ss.buf of the preceding white space, including comments. Length is .src-.sep.
	src uint32 // Index into .ss.buf, length is in .len.
	len uint32 // Length of the source representation (.src).
}

// newToken returns a newly created Token. The pos field is set equal to src.
func newToken(s *scannerSource, ch rune, sep, src, len uint32) Token {
	return Token{
		s:   s,
		Ch:  ch,
		pos: src,
		sep: sep,
		src: src,
		len: len,
	}
}

// String implements fmt.Stringer.
func (t *Token) String() string { return PrettyString(t) }

// Name returns a human readable representation of t.Ch.
func (t *Token) Name() string { return runeName(t.Ch) }

// Seq returns t's sequence number in the token stream. The information
// determines scope boundaries.
func (t *Token) Seq() int { return int(t.seq) }

// Sep returns any white space, including comments, that precede t. The result
// is R/O but it's okay to append to it.
func (t *Token) Sep() []byte {
	n := t.src - t.sep
	return t.s.buf[t.sep : t.sep+n : t.sep+n]
}

// Src returns the source representation of t. The result is R/O but it's okay
// to append to it.
func (t *Token) Src() []byte {
	n := t.len
	return t.s.buf[t.src : t.src+n : t.src+n]
}

// Set sets the values Sep and Src will report.
func (t *Token) Set(sep, src []byte) error {
	if uint64(len(sep)+len(src)) > math.MaxUint32 {
		return errorf("Token.Set: argument values too long")
	}

	if uint64(len(t.s.buf)+len(sep)+len(src)) > math.MaxUint32 {
		return errorf("Token.Set: underlying scanner source buffer overflow: size is already %v", len(t.s.buf))
	}

	t.sep = uint32(len(t.s.buf))
	t.s.buf = append(t.s.buf, sep...)
	t.src = uint32(len(t.s.buf))
	t.s.buf = append(t.s.buf, src...)
	t.len = uint32(len(src))
	return nil
}

// Position implements Node.
func (t *Token) Position() token.Position { return t.s.pos(t.pos) }

// scannerSource captures source code and the associated position information.
type scannerSource struct {
	buf  []byte
	file *token.File

	len  uint32 // Len of source code in .buf.
	pos0 token.Pos
}

// newScannerSource returns a new scanner source.
func newScannerSource(src Source) (s *scannerSource, err error) {
	s = &scannerSource{}
	switch x := src.Value.(type) {
	case io.ReadCloser:
		defer func() {
			if e := x.Close(); e != nil && err == nil {
				err = e
			}
		}()

		if s.buf, err = io.ReadAll(x); err != nil {
			return nil, errorf("", err)
		}
	case fs.File:
		fi, err := x.Stat()
		if err != nil {
			return nil, errorf("", err)
		}

		if sz := fi.Size(); sz >= math.MaxUint32 {
			return nil, errorf("source too big: %v bytes", sz)
		}

		return newScannerSource(Source{src.Name, io.ReadCloser(x)})
	case []byte:
		s.buf = x
	case io.Reader:
		if s.buf, err = io.ReadAll(x); err != nil {
			return nil, errorf("", err)
		}
	case nil:
		f, err := os.Open(src.Name)
		if err != nil {
			return nil, errorf("", err)
		}

		return newScannerSource(Source{src.Name, fs.File(f)})
	case string:
		s.buf = []byte(x)
	default:
		return nil, errorf("invalid value type: %T", x)
	}
	if uint64(len(s.buf)) >= math.MaxUint32 {
		return nil, errorf("source too big: %v bytes", len(s.buf))
	}

	// [0]5.1.1.2, 2: A source file that is not empty shall end in a new-line character, ...
	if len(s.buf) != 0 && s.buf[len(s.buf)-1] != '\n' {
		s.buf = append(s.buf, '\n')
	}
	s.len = uint32(len(s.buf))
	s.file = token.NewFile(src.Name, int(s.len))
	s.pos0 = s.file.Pos(0)
	return s, nil
}

// pos returns a token.Position representing off.
func (s *scannerSource) pos(off uint32) token.Position {
	return s.file.PositionFor(token.Pos(off)+s.pos0, true)
}

// scanner tokenizes a scannerSource.
type scanner struct {
	s  *scannerSource
	eh errHandler

	chSize int
	state  int // cppAtLineStart, ...

	ch  rune // eof at EOF or when consumed.
	off uint32
	sep uint32 // .off on .cppScan invocation.
	src uint32 // .off on .cppScan finding the start of a token after skipping the separator, if any.

	closed      bool
	joinedLines bool // Translation phase 2 skipped "\\\n".
}

// newScanner returns a new scanner. The errHandler function is invoked on
// scanner errors.
func newScanner(src Source, eh errHandler) (*scanner, error) {
	s, err := newScannerSource(src)
	if err != nil {
		return nil, err
	}

	return &scanner{
		s:  s,
		ch: eof,
		eh: eh,
	}, nil
}

// close causes all subsequent calls to .scan to return an EOF token.
func (s *scanner) close() { s.closed = true }

// c returns the current rune s is positioned on. If s is at EOF or closed, c
// returns eof.
func (s *scanner) c() rune {
	if s.closed {
		return eof
	}

	if s.ch < 0 {
		if s.off >= s.s.len {
			s.closed = true
			return eof
		}
	}

	if s.peek(0) == '\\' {
		switch {
		case s.peek(1) == '\n':
			s.joinedLines = true
			s.off += 2
		case s.peek(1) == '\r' && s.peek(2) == '\n':
			s.joinedLines = true
			s.off += 3
		}
	}

	s.ch, s.chSize = utf8.DecodeRune(s.s.buf[s.off:])
	return s.ch
}

func (s *scanner) peek(delta uint32) int {
	if s.off+delta < s.s.len {
		return int(s.s.buf[s.off+uint32(delta)])
	}

	return eof
}

// next moves s to and returns the next rune. If s is at EOF or closed, next
// returns eof.
func (s *scanner) next() rune {
	if s.closed || s.ch < 0 {
		return eof
	}

	s.off += uint32(s.chSize)
	prev := s.ch
	s.ch = eof
	r := s.c()
	if prev == '\n' {
		s.s.file.AddLine(int(s.off))
	}
	return r
}

// pos returns a token.Position representing off.
func (s *scanner) pos(off uint32) token.Position { return s.s.pos(off) }

// cppScan0 returns the next preprocessing token, see [0]6.4. If s is at EOF or
// closed, cppScan0 returns eof.
func (s *scanner) cppScan0() (tok Token) {
	s.scanSep()
	s.src = s.off
	c := s.c()
	switch c {
	case

		'(', ')', ',', ':', ';',
		'?', '[', '\n', ']', '{',
		'}', '~':

		s.next()
		return s.newToken(c)
	case '*':
		switch s.next() {
		case '=':
			s.next()
			return s.newToken(rune(MULASSIGN))
		default:
			return s.newToken(c)
		}
	case '=':
		switch s.next() {
		case '=':
			s.next()
			return s.newToken(rune(EQ))
		default:
			return s.newToken(c)
		}
	case '^':
		switch s.next() {
		case '=':
			s.next()
			return s.newToken(rune(XORASSIGN))
		default:
			return s.newToken(c)
		}
	case '+':
		switch s.next() {
		case '=':
			s.next()
			return s.newToken(rune(ADDASSIGN))
		case '+':
			s.next()
			return s.newToken(rune(INC))
		default:
			return s.newToken(c)
		}
	case '&':
		switch s.next() {
		case '=':
			s.next()
			return s.newToken(rune(ANDASSIGN))
		default:
			return s.newToken(c)
		}
	case '-':
		switch s.next() {
		case '=':
			s.next()
			return s.newToken(rune(SUBASSIGN))
		case '-':
			s.next()
			return s.newToken(rune(DEC))
		case '>':
			s.next()
			return s.newToken(rune(ARROW))
		default:
			return s.newToken(c)
		}
	case '|':
		switch s.next() {
		case '=':
			s.next()
			return s.newToken(rune(ORASSIGN))
		default:
			return s.newToken(c)
		}
	case '%':
		switch s.next() {
		case '=':
			s.next()
			return s.newToken(rune(MODASSIGN))
		default:
			return s.newToken(c)
		}
	case '/':
		switch s.next() {
		case '=':
			s.next()
			return s.newToken(rune(DIVASSIGN))
		default:
			return s.newToken(c)
		}
	case '!':
		switch s.next() {
		case '=':
			s.next()
			return s.newToken(rune(NEQ))
		default:
			return s.newToken(c)
		}
	case '#':
		switch s.next() {
		case '#':
			s.next()
			return s.newToken(rune(PPPASTE))
		default:
			return s.newToken(c)
		}
	case '<':
		switch s.next() {
		case '<':
			switch s.next() {
			case '=':
				s.next()
				return s.newToken(rune(LSHASSIGN))
			default:
				return s.newToken(rune(LSH))
			}
		case '=':
			s.next()
			return s.newToken(rune(LEQ))
		default:
			return s.newToken(c)
		}
	case '>':
		switch s.next() {
		case '>':
			switch s.next() {
			case '=':
				s.next()
				return s.newToken(rune(RSHASSIGN))
			default:
				return s.newToken(rune(RSH))
			}
		case '=':
			s.next()
			return s.newToken(rune(GEQ))
		default:
			return s.newToken(c)
		}
	case '.':
		s.next()
		switch c2 := s.c(); {
		case c2 >= '0' && c2 <= '9':
			return s.ppnumber()
		default:
			return s.newToken(c)
		}
	case '"':
		return s.stringLiteral(rune(STRINGLITERAL))
	case '\'':
		return s.characterConstant(rune(CHARCONST))
	case 'L':
		switch s.peek(1) {
		case '\'':
			s.next()
			return s.characterConstant(rune(LONGCHARCONST))
		case '"':
			s.next()
			return s.stringLiteral(rune(LONGSTRINGLITERAL))
		}
	case eof:
		s.closed = true
		return s.newToken(c)
	}

	switch {
	case unicode.IsLetter(c) || c == '_':
		return s.identifier()
	case c >= '0' && c <= '9':
		return s.ppnumber()
	default:
		s.next()
		return s.newToken(c)
	}
}

// characterConstant scans a character-constant.
//
// [0]A.1.5
func (s *scanner) characterConstant(c rune) Token {
	s.next() // '\''
	for s.cChar() {
	}
	switch s.c() {
	case '\'':
		s.next()
		return s.newToken(c)
	default:
		s.eh("%v: character constant not terminated", s.pos(s.src))
		return s.fail()
	}
}

// cChar scans c-char.
//
// [0]A.1.5
func (s *scanner) cChar() bool {
	switch s.c() {
	case '\'', '\n':
		return false
	case '\\':
		s.next()
		return s.escapeSequence()
	default:
		s.next()
		return true
	}
}

// stringLiteral scans a string-literal.
//
// [0]A.1.6
func (s *scanner) stringLiteral(c rune) Token {
	s.next() // '"'
	for s.sChar() {
	}
	switch s.c() {
	case '"':
		s.next()
		return s.newToken(c)
	default:
		s.eh("%v: string literal not terminated", s.pos(s.src))
		return s.fail()
	}
}

// sChar scans s-char.
//
// [0]A.1.6
func (s *scanner) sChar() bool {
	switch s.c() {
	case '"', '\n':
		return false
	case '\\':
		s.next()
		return s.escapeSequence()
	default:
		s.next()
		return true
	}
}

// escapeSequence scans escape-sequence.
//
// [0]A.1.5
func (s *scanner) escapeSequence() bool {
	// '\\' already consumed
	return s.simpleEscapeSequence() ||
		s.octalEscapeSequence() ||
		s.hexadecimalEscapeSequence() ||
		s.universalCharacterName()
}

// universalCharacterName scans universal-character-name.
//
// [0]A.1.4
func (s *scanner) universalCharacterName() bool {
	// '\\' already consumed
	switch s.c() {
	case 'u':
		s.hexQuad()
		return true
	case 'U':
		if s.hexQuad() {
			s.hexQuad()
		}
		return true
	}
	return false
}

// hexQuad scans hex-quad.
//
// [0]A.1.4
func (s *scanner) hexQuad() (r bool) {
	for i := 0; i < 4; i++ {
		if !s.hexadecimalDigit() {
			s.eh("%v: expected hexadecimal digit", s.pos(s.off))
			return r
		}

		r = true
	}
	return r
}

// hexadecimalEscapeSequence scans hexadecimal-escape-sequence.
//
// [0]A.1.5
func (s *scanner) hexadecimalEscapeSequence() bool {
	// '\\' already consumed
	switch s.c() {
	case 'x', 'X':
		s.next()
		ok := false
		for s.hexadecimalDigit() {
			ok = true
		}
		if !ok {
			s.eh("%v: expected hexadecimal digit", s.pos(s.off))
		}
		return true
	}
	return false
}

// hexadecimalDigit scans hexadecimal-digit.
//
// [0]A.1.5
func (s *scanner) hexadecimalDigit() bool {
	switch c := s.c(); {
	case
		c >= '0' && c <= '9',
		c >= 'a' && c <= 'f',
		c >= 'A' && c <= 'F':

		s.next()
		return true
	}
	return false
}

// octalEscapeSequence scans octal-escape-sequence.
//
// [0]A.1.5
func (s *scanner) octalEscapeSequence() bool {
	// '\\' already consumed
	switch c := s.c(); {
	case c >= '0' && c <= '7':
		for s.octalDigit() {
		}
		return true
	}
	return false
}

// octalDigit scans octal-digit.
//
// [0]A.1.5
func (s *scanner) octalDigit() bool {
	switch c := s.c(); {
	case c >= '0' && c <= '7':
		s.next()
		return true
	}
	return false
}

// simpleEscapeSequence scans simple-escape-sequence.
//
// [0]A.1.5
func (s *scanner) simpleEscapeSequence() bool {
	// '\\' already consumed
	switch s.c() {
	case '\'', '"', '?', '\\', 'a', 'b', 'e', 'f', 'n', 'r', 't', 'v':
		s.next()
		return true
	}
	return false
}

// newToken returns a newly created Token, conveniently filling in the usual defaults.
func (s *scanner) newToken(c rune) Token {
	switch {
	case s.joinedLines:
		s.joinedLines = false
		tok := newToken(s.s, c, s.sep, s.src, s.off-s.src)
		tok.Set(
			bytes.ReplaceAll(tok.Sep(), []byte("\\\n"), nil),
			bytes.ReplaceAll(tok.Src(), []byte("\\\n"), nil),
		)
		return tok
	default:
		return newToken(s.s, c, s.sep, s.src, s.off-s.src)

	}
}

// ppnumber scans preprocessor numbers.
//
// [0]A.1.9
func (s *scanner) ppnumber() Token {
	s.next()
	for {
		switch c := s.c(); c {
		case '.':
			s.next()
		case 'e', 'E', 'p', 'P':
			s.next()
			s.sign(false)
		default:
			switch {
			case c >= '0' && c <= '9':
				s.next()
			case unicode.IsLetter(c):
				s.identifier()
			default:
				return s.newToken(rune(PPNUMBER))
			}
		}
	}
}

// sign scans sign.
//
// [0]A.1.5
func (s *scanner) sign(must bool) {
	switch s.c() {
	case '+', '-':
		s.next()
	default:
		if must {
			s.next()
			s.eh("%v: expected sign", s.pos(s.off))
		}
	}
}

// scanSep will set s.sep to s.off and scan until end the of a separator, if
// any.
func (s *scanner) scanSep() {
	s.sep = s.off
	for {
		switch s.c() {
		case ' ', '\t', '\f', '\v', '\r':
			s.next()
		case '/':
			off := s.off
			switch s.next() {
			case '*':
				s.comment(off)
			case '/':
				s.lineComment(off)
				return
			default:
				s.off = off
				s.ch = eof
				return
			}
		default:
			return
		}
	}
}

// lineComment scans until the end of a //-style comment. The leading '/' is
// already consumed.
func (s *scanner) lineComment(start uint32) {
	s.next() // '/'
	for {
		switch s.c() {
		case eof, '\n':
			return
		}
		s.next()
	}
}

// identifier scans an identifier
func (s *scanner) identifier() Token {
	s.next()
	for {
		switch c := s.c(); {
		case unicode.IsLetter(c) || c == '_' || unicode.IsDigit(c):
			s.next()
		default:
			return s.newToken(rune(IDENTIFIER))
		}
	}
}

// fail reports an error at current position, closes s and returns an EOF
// Token.
func (s *scanner) fail() Token {
	s.eh("%v: unexpected rune: %s (%s)", s.pos(s.off), runeName(s.c()), origin(2))
	s.closed = true
	return newToken(s.s, eof, s.off, s.off, 0)
}

// comment scans until the end of a /*-style comment. The leading '/' is
// already consumed.
func (s *scanner) comment(start uint32) {
	s.next() // '*'
	for {
		switch s.c() {
		case eof:
			s.eh("%v: comment not terminated", s.pos(start))
			return
		case '*':
			switch s.next() {
			case '/':
				s.next()
				return
			}
		default:
			s.next()
		}
	}
}

const (
	cppScanAtLineStart = iota
	cppScanHash        // Returned a first token off a line and it was '#'.
	cppScanOther
	cppScanHeaderName
)

// cppScan returns the next preprocessing token, see [0]6.4. If s is at EOF or
// closed, cppScan returns eof.
func (s *scanner) cppScan() (tok Token) {
	switch s.state {
	case cppScanAtLineStart:
		switch tok = s.cppScan0(); tok.Ch {
		case '\n':
			s.state = cppScanAtLineStart
		case '#':
			s.state = cppScanHash
		default:
			s.state = cppScanOther
		}
		return tok
	case cppScanHash:
		switch tok = s.cppScan0(); {
		case tok.Ch == rune(IDENTIFIER) && bytes.Equal(tok.Src(), []byte("include")):
			s.state = cppScanHeaderName
		default:
			s.state = cppScanOther
		}
		return tok
	case cppScanHeaderName:
		s.state = cppScanOther
		return s.headerName()
	case cppScanOther:
		switch tok = s.cppScan0(); tok.Ch {
		case '\n':
			s.state = cppScanAtLineStart
		}
		return tok
	default:
		s.eh("%v: internal error, scanner state: %v", s.pos(s.off), s.state)
		s.state = cppScanOther
		return s.cppScan0()
	}
}

// headerName scans preprocessor header-name.
//
// [0]A.1.8
func (s *scanner) headerName() Token {
	s.scanSep()
	s.src = s.off
	switch s.c() {
	case '<':
		for {
			switch s.next() {
			case '>':
				s.next()
				s.state = cppScanOther
				return s.newToken(rune(HEADER_NAME))
			case '\n':
				s.state = cppScanAtLineStart
				return s.fail()
			}
		}
	case '"':
		for {
			switch s.next() {
			case '"':
				s.next()
				s.state = cppScanOther
				return s.newToken(rune(HEADER_NAME))
			case '\n':
				s.state = cppScanAtLineStart
				return s.fail()
			}
		}
	default:
		return s.cppScan0()
	}
}
