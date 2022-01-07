// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	maxIncludeLevel = 200 // gcc, std is at least 15.
)

var (
	_ tokenSequence = (*cat)(nil)
	_ tokenSequence = (*preprocessingTokens)(nil)
	_ tokenSequence = (*tokenizer)(nil)

	commaTok, eofTok, hashTok, nlTok, oneTok, pragmaTok, pragmaTestTok, spTok, zeroTok, emptyStringTok Token

	emptyStringPreprocessingTok preprocessingToken
	eofPreprocessToken          preprocessingToken

	comma  = []byte{','}
	hash   = []byte{'#'}
	nl     = []byte{'\n'}
	one    = []byte{'1'}
	pragma = []byte("pragma")
	qq     = []byte(`""`)
	sp     = []byte{' '}
	zero   = []byte{'0'}

	protectedMacros = map[string]struct{}{
		"__DATE__":                 {},
		"__FILE__":                 {},
		"__LINE__":                 {},
		"__STDC_HOSTED__":          {},
		"__STDC_MB_MIGHT_NEQ_WC__": {},
		"__STDC_VERSION__":         {},
		"__STDC__":                 {},
		"__TIME__":                 {},
		"defined":                  {},
	}
)

func init() {
	s, err := newScannerSource(Source{"", []byte(nil), nil})
	if err != nil {
		panic(errorf("initialization error"))
	}

	commaTok.s = s
	commaTok.Ch = ','
	commaTok.Set(nil, comma)

	emptyStringTok.s = s
	emptyStringTok.Ch = rune(STRINGLITERAL)
	emptyStringTok.Set(nil, qq)

	hashTok.s = s
	hashTok.Ch = '#'
	hashTok.Set(nil, hash)

	eofTok.s = s
	eofTok.Ch = eof
	eofTok.Set(nil, nil)

	nlTok.s = s
	nlTok.Ch = '\n'
	nlTok.Set(nil, []byte{'\n'})

	pragmaTok.s = s
	pragmaTok.Ch = rune(IDENTIFIER)
	pragmaTok.Set(nil, pragma)

	pragmaTestTok.s = s
	pragmaTestTok.Ch = rune(IDENTIFIER)
	pragmaTestTok.Set(nil, []byte("__pragma"))

	oneTok.s = s
	oneTok.Ch = rune(PPNUMBER)
	oneTok.Set(nil, one)

	spTok.s = s
	spTok.Ch = ' '
	spTok.Set(nil, sp)

	zeroTok.s = s
	zeroTok.Ch = rune(PPNUMBER)
	zeroTok.Set(nil, zero)

	emptyStringPreprocessingTok = preprocessingToken{emptyStringTok, nil}
	eofPreprocessToken = preprocessingToken{eofTok, nil}
}

type hideSet map[string]struct{}

func (h *hideSet) String() string {
	if h == nil {
		return ""
	}

	hs := *h
	var a []string
	for k := range hs {
		a = append(a, k)
	}
	return fmt.Sprintf("%s", a)
}

func (h hideSet) has(s string) bool { _, ok := h[s]; return ok }

func (h hideSet) add(s string) hideSet {
	r := hideSet{}
	for k := range h {
		r[k] = struct{}{}
	}
	r[s] = struct{}{}
	return r
}

// ∪ - union.
func (h hideSet) cup(src hideSet) hideSet {
	if h == nil {
		return src
	}

	if src == nil {
		return h
	}

	r := hideSet{}
	for k := range h {
		r[k] = struct{}{}
	}
	for k := range src {
		r[k] = struct{}{}
	}
	return r
}

// ∩ - intersection.
func (h hideSet) cap(src hideSet) hideSet {
	if h == nil {
		return nil
	}

	if src == nil {
		return nil
	}

	r := hideSet{}
	for k := range h {
		if _, ok := src[k]; ok {
			r[k] = struct{}{}
		}
	}
	return r
}

type preprocessingToken struct {
	Token
	hs hideSet
}

// Macro represents a preprocessor #define.
type Macro struct {
	Name            Token
	Params          []Token
	params          []string
	replacementList preprocessingTokens

	MinArgs int // m x: 0, m() x: 0, m(...): 0, m(a) a: 1, m(a, ...): 1, m(a, b): 2, m(a, b, ...): 2.
	VarArg  int // m(a): -1, m(...): 0, m(a, ...): 1, m(a...): 0, m(a, b...): 1.

	IsFnLike bool // m: false, m(): true, m(x): true.
}

func newMacro(nm Token, params []Token, replList []preprocessingToken, minArgs, varArg int, isFnLike bool) (*Macro, error) {
	var fp []string
	for _, v := range params {
		fp = append(fp, string(v.Src()))
	}
	if len(fp) > 1 {
		m := map[string]struct{}{}
		for i, v := range fp {
			if _, ok := m[v]; ok {
				return nil, errorf("%v: duplicate parameter name", params[i].Position())
			}
		}
	}
	return &Macro{
		IsFnLike:        isFnLike,
		MinArgs:         minArgs,
		Name:            nm,
		Params:          params,
		VarArg:          varArg,
		params:          fp,
		replacementList: replList,
	}, nil
}

func (m *Macro) is(s string) int {
	for i, v := range m.params {
		if s == v {
			return i
		}
	}

	if m.VarArg >= 0 && s == "__VA_ARGS__" {
		return m.VarArg
	}

	return -1
}

// ReplacementList returns the tokens m is substitued with.
func (m *Macro) ReplacementList() (r []Token) {
	for _, v := range m.replacementList {
		if v.Ch != ' ' {
			r = append(r, v.Token)
		}
	}
	return r
}

func (m *Macro) fp() []string {
	if !m.IsFnLike {
		return nil
	}

	return m.params
}

func (m *Macro) ts() *preprocessingTokens {
	r := m.replacementList
	return &r
}

type tokenSequence interface {
	peek(int) preprocessingToken
	peekNonBlank() (preprocessingToken, int)
	read() preprocessingToken
	skip(int)
}

type tokenizer struct {
	c   *cpp
	in  []preprocessingToken
	out []preprocessingToken
}

func newTokenizer(c *cpp) *tokenizer {
	return &tokenizer{c: c}
}

func (t *tokenizer) peek(index int) (tok preprocessingToken) {
	for {
		if index < len(t.in) {
			return t.in[index]
		}

		t.in = append(t.in, tokens2PreprocessingTokens(t.c.nextLine(), false)...)
		continue
	}
}

func (t *tokenizer) peekNonBlank() (preprocessingToken, int) {
	for i := 0; ; i++ {
		if tok := t.peek(i); tok.Ch != ' ' && tok.Ch != '\n' {
			return tok, i
		}
	}
}

func (t *tokenizer) read() (tok preprocessingToken) {
	if len(t.in) == 0 {
		t.in = tokens2PreprocessingTokens(t.c.nextLine(), false)
	}

	tok = t.in[0]
	if tok.Ch != eof {
		t.in = t.in[1:]
	}
	return tok

}

func (t *tokenizer) skip(n int) {
	for ; n != 0; n-- {
		t.read()
	}
}

func (t *tokenizer) token() (tok Token) {
	for len(t.out) == 0 {
		t.out = t.c.expand(true, false, t)
	}
	tok = t.out[0].Token
	if tok.Ch != eof {
		t.out = t.out[1:]
	}
	return tok
}

type preprocessingTokens []preprocessingToken

func (p *preprocessingTokens) peek(index int) preprocessingToken {
	if index < len(*p) {
		return (*p)[index]
	}

	return eofPreprocessToken
}

func (p *preprocessingTokens) read() (tok preprocessingToken) {
	s := *p
	if len(s) == 0 {
		return eofPreprocessToken
	}

	tok = s[0]
	if tok.Ch != eof {
		s = s[1:]
		*p = s
	}
	return tok
}

func (p *preprocessingTokens) skipBlank() {
	s := *p
	for len(s) != 0 && s[0].Ch == ' ' {
		s = s[1:]
	}
	*p = s
}

func (p *preprocessingTokens) peekNonBlank() (preprocessingToken, int) {
	for i, v := range *p {
		if v.Ch != ' ' && v.Ch != '\n' {
			return v, i
		}
	}

	return eofPreprocessToken, -1
}

func (p *preprocessingTokens) skip(n int) {
	*p = (*p)[n:]
}

func (p *preprocessingTokens) c() Token {
	p.skipBlank()
	return p.peek(0).Token
}

type cat struct {
	head preprocessingTokens
	tail tokenSequence
}

func newCat(head preprocessingTokens, tail tokenSequence) tokenSequence {
	if len(head) == 0 {
		return tail
	}

	switch x := tail.(type) {
	case *tokenizer:
		return &cat{head, tail}
	case *cat:
		if len(x.head) == 0 {
			return &cat{head, x.tail}
		}

		return &cat{head, tail}
	case *preprocessingTokens:
		t := *x
		if len(t) == 0 {
			return &head
		}

		head = append(head[:len(head):len(head)], t...)
		return &head
	default:
		panic(todo("%T", x))
	}
}

func (c *cat) peek(index int) preprocessingToken {
	if index < len(c.head) {
		return c.head.peek(index)
	}

	return c.tail.peek(index - len(c.head))
}

func (c *cat) peekNonBlank() (preprocessingToken, int) {
	for i := 0; ; i++ {
		if t := c.peek(i); t.Ch != ' ' && t.Ch != '\n' {
			return t, i
		}
	}
}

func (c *cat) read() preprocessingToken {
	if len(c.head) != 0 {
		return c.head.read()
	}

	return c.tail.read()
}

func (c *cat) skip(n int) {
	for ; n != 0; n-- {
		c.read()
	}
}

// cpp is the C preprocessor.
type cpp struct {
	cfg         *Config
	eh          errHandler
	eof         eofLine
	groups      map[string]group
	indentLevel int // debug dumps
	macros      map[string]*Macro
	sources     []Source
	srcStack    []string
	stack       []interface{}
	tok         Token
	tokenizer   *tokenizer
	tos         interface{}

	includeLevel int

	closed bool
}

// newCPP returns a newly created cpp.
func newCPP(cfg *Config, sources []Source, eh errHandler) (*cpp, error) {
	m := map[string]struct{}{}
	for _, v := range sources {
		if _, ok := m[v.Name]; ok {
			return nil, errorf("duplicate source name: %q", v.Name)
		}

		if v.FS == nil {
			v.FS = cfg.FS
		}
		m[v.Name] = struct{}{}
	}
	c := &cpp{
		cfg:     cfg,
		groups:  map[string]group{},
		eh:      eh,
		macros:  map[string]*Macro{},
		sources: sources,
	}
	c.tokenizer = newTokenizer(c)
	c.tok.Ch = eof // Invalidate
	return c, nil
}

// c returns the current token rune the preprocessor is positioned on.
func (c *cpp) c() rune {
	if c.closed || c.tok.Ch != eof {
		return c.tok.Ch
	}

	if c.tok = c.tokenizer.token(); c.tok.Ch == eof {
		c.closed = true
	}
	return c.tok.Ch
}

// consume returns c.tok and invalidates c.tok.Ch.
func (c *cpp) consume() (r Token) {
	r = c.tok
	c.tok.Ch = eof
	return r
}

func (c *cpp) indent() string {
	c.indentLevel++
	return fmt.Sprintf("\t%s", strings.Repeat("· ", c.indentLevel-1))
}

func (c *cpp) undent() string {
	c.indentLevel--
	return fmt.Sprintf("\t%s", strings.Repeat("· ", c.indentLevel))
}

// [1], pg 1.
func (c *cpp) expand(outer, eval bool, TS tokenSequence) (r preprocessingTokens) {
	// trc("* %s%v outer %v (%v)", c.indent(), toksDump(TS), outer, origin(2))
	// defer func() { trc("->%s%v", c.undent(), toksDump(r)) }()
more:
	t := TS.read()
	// trc("  %[2]s%v", c.undent(), c.indent(), &t)
	if t.Ch == eof {
		// if TS is {} then
		//	return {};
		if outer {
			return []preprocessingToken{t}
		}

		return nil
	}

	T := t.Token
	HS := t.hs
	src := string(T.Src())
	if eval && src == "defined" {
		TS = c.parseDefined(TS)
		goto more
	}

	if src == "_Pragma" {
		c.parsePragma(TS)
		goto more
	}

	if HS.has(src) {
		// if TS is T^HS • TS’ and T is in HS then
		//	return T^HS • expand(TS’);
		if outer {
			panic(todo(""))
		}

		if x, ok := TS.(*cat); ok && len(x.head) == 0 {
			if _, ok := x.tail.(*tokenizer); ok {
				panic(todo(""))
				// return preprocessingTokens{t}
			}
		}

		return append(preprocessingTokens{t}, c.expand(false, eval, TS)...)
	}

	if m := c.macro(&T, src); m != nil {
	out:
		switch {
		default:
			// if TS is T^HS • TS’ and T is a "()-less macro" then
			//	return expand(subst(ts(T),{},{},HS∪{T},{}) • TS’);

			// trc("  %s<%s is a ()-less macro, expanding to %s>", c.indent(), T.Src(), toksDump(m.replacementList))
			// defer func() { trc("  %s<%s expanded>", c.undent(), T.Src()) }()
			subst := c.subst(eval, m, m.ts(), nil, nil, HS.add(src), nil)
			if len(subst) == 0 {
				if outer {
					return nil
				}

				return c.expand(false, eval, TS)
			}

			return c.expand(false, eval, newCat(subst, TS))
		case m.IsFnLike:
			t2, skip := TS.peekNonBlank()
			if t2.Ch == '(' {
				// if TS is T^HS • ( • TS’ and T is a "()’d macro" then
				//	check TS’ is actuals • )^HS’ • TS’’ and actuals are "correct for T"
				//	return expand(subst(ts(T),fp(T),actuals,(HS∩HS’)∪{T},{}) • TS’’);

				TS.skip(skip + 1)
				args, rparen, ok := c.parseMacroArgs(TS)
				if !ok {
					return nil
				}

				if len(args) > m.MinArgs && m.VarArg < 0 {
					c.eh("%v: too many macro arguments", rparen.Position())
					break out
				}

				// trc("  %s<%s is a (%v)'d macro, expanding to %s> using args %v", c.indent(), T.Src(), m.fp(), toksDump(m.replacementList), toksDump(args))
				// defer func() { trc("  %s<%s expanded>", c.undent(), T.Src()) }()
				return c.expand(false, eval, newCat(c.subst(eval, m, m.ts(), m.fp(), args, HS.cap(rparen.hs).add(src), nil), TS))
			}
		}

	}

	// note TS must be T HS • TS’
	// return T HS • expand(TS’);
	if outer {
		return preprocessingTokens{t}
	}

	if x, ok := TS.(*cat); ok && len(x.head) == 0 {
		if _, ok := x.tail.(*tokenizer); ok {
			return preprocessingTokens{t}
		}
	}

	return append(preprocessingTokens{t}, c.expand(false, eval, TS)...)
}

// [1], pg 2.
func (c *cpp) subst(eval bool, m *Macro, IS *preprocessingTokens, FP []string, AP []preprocessingTokens, HS hideSet, OS preprocessingTokens) (r preprocessingTokens) {
	// trc("* %s%v, HS %v, FP %v, AP %v, OS %v (%v)", c.indent(), toksDump(IS), &HS, FP, toksDump(AP), toksDump(OS), origin(2))
	// defer func() { trc("->%s%v", c.undent(), toksDump(r)) }()
	t := IS.read()
	// trc("  %[2]s%v", c.undent(), c.indent(), &t)
	if t.Ch == eof {
		// if IS is {} then
		//	return hsadd(HS,OS);
		return c.hsAdd(HS, OS)
	}

	if t.Ch == '#' {
		if t2, skip := IS.peekNonBlank(); t2.Ch == rune(IDENTIFIER) {
			if i := m.is(string(t2.Src())); i >= 0 {
				// if IS is # • T • IS’ and T is FP[i] then
				//	return subst(IS’,FP,AP,HS,OS • stringize(select(i,AP )));
				IS.skip(skip + 1)
				return c.subst(eval, m, IS, FP, AP, HS, append(OS, c.stringize(c.apSelect(m, t2, AP, i))))
			}
		}
	}

	if t.Ch == ' ' && IS.peek(0).Ch == rune(PPPASTE) {
		t = IS.read()
	}
	if t.Ch == rune(PPPASTE) {
		t2, skip := IS.peekNonBlank()
		if t2.Ch == rune(IDENTIFIER) {
			if i := m.is(string(t2.Src())); i >= 0 {
				// if IS is ## • T • IS’ and T is FP[i] then
				if i >= len(AP) || len(AP[i]) == 0 {
					//	if select(i,AP ) is {} then /* only if actuals can be empty */
					//		return subst(IS’,FP,AP,HS,OS );
					IS.skip(skip + 1)
					return c.subst(eval, m, IS, FP, AP, HS, OS)
				}

				//	else
				//		return subst(IS’,FP,AP,HS,glue(OS,select(i,AP )));
				IS.skip(skip + 1)
				return c.subst(eval, m, IS, FP, AP, HS, c.glue(OS, c.apSelect(m, t2, AP, i)))

			}
		}

		// else if IS is ## • T HS’ • IS’ then
		//	return subst(IS’,FP,AP,HS,glue(OS,T^HS’));
		IS.skip(skip + 1)
		return c.subst(eval, m, IS, FP, AP, HS, c.glue(OS, preprocessingTokens{t2}))
	}

	if t.Ch == rune(IDENTIFIER) {
		if t2, skip := IS.peekNonBlank(); t2.Ch == rune(PPPASTE) {
			if i := m.is(string(t.Src())); i >= 0 {
				// if IS is T • ##^HS’ • IS’ and T is FP[i] then
				//	if select(i,AP ) is {} then /* only if actuals can be empty */
				if i >= len(AP) || len(AP[i]) == 0 {
					IS.skip(skip + 1) // ##
					t2, skip := IS.peekNonBlank()
					if j := m.is(string(t2.Src())); j >= 0 {
						//		if IS’ is T’ • IS’’ and T’ is FP[j] then
						//			return subst(IS’’,FP,AP,HS,OS • select(j,AP));
						IS.skip(skip + 1)
						return c.subst(eval, m, IS, FP, AP, HS, append(OS, c.apSelect(m, t, AP, j)...))
					}

					//		else
					//			return subst(IS’,FP,AP,HS,OS);
					panic(todo(""))
				} else {
					//	else
					//		return subst(##^HS’ • IS’,FP,AP,HS,OS • select(i,AP));
					return c.subst(eval, m, IS, FP, AP, HS, append(OS, c.apSelect(m, t, AP, i)...))
				}
			}
		}
	}

	if len(FP) != 0 || m.VarArg >= 0 {
		if i := m.is(string(t.Src())); i >= 0 {
			// if IS is T • IS’ and T is FP[i] then
			//	return subst(IS’,FP,AP,HS,OS • expand(select(i,AP )));
			return c.subst(eval, m, IS, FP, AP, HS, append(OS, c.expand(false, eval, c.apSelectP(m, t, AP, i))...))
		}
	}

	// note IS must be T HS’ • IS’
	// return subst(IS’,FP,AP,HS,OS • T HS’ );
	return c.subst(eval, m, IS, FP, AP, HS, append(OS, t))

}

func (c *cpp) glue(LS, RS preprocessingTokens) (r preprocessingTokens) {
	// trc("  \t%sLS %v, RS %v (%v)", c.indent(), toksDump(LS), toksDump(RS), origin(2))
	// defer func() { trc("->\t%s%v", c.undent(), toksDump(r)) }()

	// if LS is L^HS and RS is R^HS’ • RS’ then
	//	return L&R^(HS∩HS’) • RS’; /* undefined if L&R is invalid */
	switch len(LS) {
	case 0:
		return RS
	case 1:
		t := LS[0]
		t.Set(nil, append(t.Src(), RS[0].Src()...))
		t.hs = t.hs.cap(RS[0].hs)
		return append(preprocessingTokens{t}, RS[1:]...)
	}

	// note LS must be L^HS • LS’
	// return L^HS • glue(LS’,RS );
	return append(LS[:1:1], c.glue(LS[1:], RS)...)
}

// Given a token sequence, stringize returns a single string literal token
// containing the concatenated spellings of the tokens.
//
// [1] pg. 3
func (c *cpp) stringize(s0 preprocessingTokens) (r preprocessingToken) {
	// trc("  %s%v (%v)", c.indent(), toksDump(s0), origin(2))
	// defer func() { trc("->%s%v: %s", c.undent(), toksDump(preprocessingTokens{r}), r.Src()) }()
	if len(s0) == 0 {
		return emptyStringPreprocessingTok
	}

	r = s0[0]

	// [0]6.10.3.2
	//
	// Each occurrence of white space between the argument’s preprocessing
	// tokens becomes a single space character in the character string
	// literal.
	s := make([]Token, 0, len(s0))
	var last rune
	for _, v := range s0 {
		if v.Ch == '\n' {
			v.Ch = ' '
			v.Set(nil, sp)
		}
		if v.Ch == ' ' {
			if last == ' ' {
				continue
			}
		}

		last = v.Ch
		s = append(s, v.Token)
	}

	// White space before the first preprocessing token and after the last
	// preprocessing token composing the argument is deleted.
	s = c.toksTrim(s)

	// The character string literal corresponding to an empty argument is
	// ""
	if len(s) == 0 {
		r.Ch = rune(STRINGLITERAL)
		r.Set(nil, []byte(`""`))
		return r
	}

	var a []string
	// Otherwise, the original spelling of each preprocessing token in the
	// argument is retained in the character string literal, except for
	// special handling for producing the spelling of string literals and
	// character constants: a \ character is inserted before each " and \
	// character of a character constant or string literal (including the
	// delimiting " characters), except that it is implementation-defined
	// whether a \ character is inserted before the \ character beginning a
	// universal character name.
	for _, v := range s {
		s := string(v.Src())
		switch v.Ch {
		case rune(CHARCONST), rune(STRINGLITERAL), rune(LONGCHARCONST), rune(LONGSTRINGLITERAL):
			s = strings.ReplaceAll(s, `\`, `\\`)
			s = strings.ReplaceAll(s, `"`, `\"`)
		}
		a = append(a, s)
	}
	r.Ch = rune(STRINGLITERAL)
	r.Set(nil, []byte(`"`+strings.Join(a, "")+`"`))
	return r
}

func (c *cpp) parsePragma(ts tokenSequence) {
	t, skip := ts.peekNonBlank()
	ts.skip(skip + 1)
	if t.Ch != '(' {
		c.eh("%v: expected '('", t.Position())
		return
	}

	t2, skip := ts.peekNonBlank()
	ts.skip(skip + 1)
	s := string(t2.Src())
	switch t2.Ch {
	case rune(STRINGLITERAL):
		// ok
	case rune(LONGSTRINGLITERAL):
		s = s[1:] // Remove leading 'L'
	default:
		c.eh("%v: expected string-literal", t2.Position())
		return
	}

	t, skip = ts.peekNonBlank()
	ts.skip(skip + 1)
	if t.Ch != ')' {
		c.eh("%v: expected '('", t.Position())
	}

	// [0]6.10.9 The string literal is destringized by deleting the L prefix, if
	// present, deleting the leading and trailing double-quotes, replacing each
	// escape sequence \" by a double-quote, and replacing each escape sequence \\
	// by a single backslash. The resulting sequence of characters is processed
	// through translation phase 3 to produce preprocessing tokens that are
	// executed as if they were the pp-tokens in a pragma directive.
	s = s[1 : len(s)-1]
	s = strings.ReplaceAll(s, `\"`, `"`)
	s = strings.ReplaceAll(s, `\\`, `\`)
	sc, err := newScanner(Source{"_Pragma_", s, nil}, c.eh)
	if err != nil {
		c.eh("%v: %v", t2.Position(), err)
		return
	}

	a := controlLine{hashTok, pragmaTok}
	for {
		t := sc.cppScan0()
		a = append(a, t)
		if t.Ch == '\n' {
			break
		}
	}
	c.push(a)
}

func (c *cpp) macro(n Node, nm string) *Macro {
	if m := c.macros[nm]; m != nil {
		return m
	}

	if _, ok := protectedMacros[nm]; !ok {
		return nil
	}

	panic(todo("%v: %q", n.Position(), nm))
}

func (c *cpp) parseDefined(ts tokenSequence) (r tokenSequence) {
	c.skipBlank(ts)
	var t preprocessingToken
	switch p := ts.peek(0); p.Ch {
	case '(':
		ts.read()
		c.skipBlank(ts)
		switch p = ts.peek(0); p.Ch {
		case rune(IDENTIFIER):
			t = ts.read()
			c.skipBlank(ts)
			if paren := ts.read(); paren.Ch != ')' {
				c.eh("%v: expected ')'", paren.Position())
			}
		default:
			c.eh("%v: operator \"defined\" requires an identifier", p.Position())
			ts.read()
			return ts
		}
	case rune(IDENTIFIER):
		t = ts.read()
	default:
		c.eh("%v: operator \"defined\" requires an identifier", p.Position())
		ts.read()
		return ts
	}

	switch c.macro(&t.Token, string(t.Src())) {
	case nil:
		t.Token = zeroTok
	default:
		t.Token = oneTok
	}

	s := preprocessingTokens(append([]preprocessingToken{t}, *ts.(*preprocessingTokens)...))
	return &s
}

func (c *cpp) skipBlank(ts tokenSequence) {
	for ts.peek(0).Ch == ' ' {
		ts.read()
	}
}

func (c *cpp) apSelectP(m *Macro, t preprocessingToken, AP []preprocessingTokens, i int) *preprocessingTokens {
	r := c.apSelect(m, t, AP, i)
	return &r
}

func (c *cpp) apSelect(m *Macro, t preprocessingToken, AP []preprocessingTokens, i int) preprocessingTokens {
	if m.VarArg < 0 || m.VarArg != i {
		if i < len(AP) {
			return AP[i]
		}

		return nil
	}

	return c.varArgs(t, AP[i:])
}

func (c *cpp) varArgsP(t preprocessingToken, AP []preprocessingTokens) *preprocessingTokens {
	r := c.varArgs(t, AP)
	return &r
}

func (c *cpp) varArgs(t preprocessingToken, AP []preprocessingTokens) (r preprocessingTokens) {
	// trc("  %s%v %v (%v)", c.indent(), &t, toksDump(AP), origin(2))
	// defer func() { trc("->%sout %v", c.undent(), toksDump(r)) }()
	for i, v := range AP {
		if i != 0 {
			r = append(r, preprocessingToken{commaTok, nil})
			if len(v) != 0 && len(v[0].Sep()) != 0 {
				r = append(r, preprocessingToken{spTok, nil})
			}
		}
		r = append(r, v...)
	}
	return r
}

func (c *cpp) parseMacroArgs(TS tokenSequence) (args []preprocessingTokens, rparen preprocessingToken, ok bool) {
	// trc("  %sin  %v (%v)", c.indent(), toksDump(TS), origin(2))
	// defer func() { trc("->%sout %v", c.undent(), toksDump(args)) }()
	var arg preprocessingTokens
	level := 0
	var t preprocessingToken
out:
	for {
		t = TS.read()
		switch t.Ch {
		case ',':
			if level != 0 {
				arg = append(arg, t)
				break
			}

			args = append(args, arg)
			arg = nil
		case '(':
			arg = append(arg, t)
			level++
		case ')':
			if level == 0 {
				args = append(args, arg)
				break out
			}

			arg = append(arg, t)
			level--
		case eof:
			panic(todo("", &t))
			// 			panic(todo("", &t))
		default:
			arg = append(arg, t)
		}
	}
	for i, arg := range args {
		args[i] = toksTrim(arg)
	}
	if len(args) == 1 && len(args[0]) == 0 {
		args = nil
	}
	return args, t, true
}

func (c *cpp) hsAdd(HS hideSet, TS preprocessingTokens) (r preprocessingTokens) {
	if len(TS) == 0 {
		// if TS is {} then
		//	return {};
		return r
	}

	// note TS must be T^HS’ • TS’
	// return T^(HS∪HS’) • hsadd(HS,TS’);
	for _, v := range TS {
		v.hs = v.hs.cup(HS)
		r = append(r, v)
	}
	return r
}

// nextLine returns the next input textLine.
func (c *cpp) nextLine() textLine {
	for {
		// a := []string{fmt.Sprintf("%T", c.tos)}
		// for _, v := range c.stack {
		// 	a = append(a, fmt.Sprintf("%T", v))
		// }
		// trc("STACK %v", a)
		switch x := c.tos.(type) {
		case nil:
			// trc("<nil>, len(c.sources): %v", len(c.sources))
			if len(c.sources) == 0 {
				return textLine{Token(c.eof)}
			}

			src := c.sources[0]
			c.sources = c.sources[1:]
			g, err := c.group(src)
			if err != nil {
				c.eh("%v", err)
				break
			}

			c.push(g)
		case group:
			// trc("group len %v", len(x))
			if len(x) == 0 {
				c.pop()
				break
			}

			c.tos = x[1:]
			c.push(x[0])
		case controlLine:
			// trc("%v: controlLine %v", x[0].Position(), toksDump(x))
			c.pop()
			switch string(x[1].Src()) {
			case "define":
				c.define(x)
			case "undef":
				c.undef(x)
			case "include":
				c.include(x)
			case "pragma":
				// eg.  ["#" "pragma" "STDC" "FP_CONTRACT" "ON" "\n"]
				if h := c.cfg.PragmaHandler; h != nil {
					if err := h(x[2 : len(x)-1]); err != nil {
						c.eh("%v:", err)
					}
				}
			default:
				panic(todo("%v: %q", x[0].Position(), x[1].Src()))
			}
		case textLine:
			// trc("%v: textLine %v", x[0].Position(), toksDump(x))
			c.pop()
			return x
		case eofLine:
			c.eof = x
			c.srcStack = c.srcStack[:len(c.srcStack)-1]
			// trc("%v: eofLine, len(c.stack): %v", (*Token)(&x).Position(), len(c.stack))
			if len(c.stack) == 0 {
				trc("EOF")
				return textLine{Token(x)}
			}

			c.pop()
		case *ifSection:
			switch {
			case x.ifGroup != nil:
				// trc("%v: ifSection/if %v", x.ifGroup.line[0].Position, toksDump(x.ifGroup.line))
				switch {
				case c.ifGroup(x.ifGroup):
					c.tos = x.ifGroup.group
				case len(x.elifGroups) != 0:
					y := *x
					y.ifGroup = nil
					c.tos = &y
				case x.elseGroup != nil:
					c.tos = x.elseGroup.group
				default:
					c.pop()
				}
			case len(x.elifGroups) != 0:
				// trc("%v: ifSection/elif %v", x.elifGroups[0].line[0].Position, toksDump(x.elifGroups[0].line))
				if eg := x.elifGroups[0]; c.elifGroup(eg) {
					c.tos = eg.group
					break
				}

				x.elifGroups = x.elifGroups[1:]
			case x.elseGroup != nil:
				//	# else new-line group_opt
				c.tos = x.elseGroup.group
			default:
				c.pop()
			}
		default:
			panic(todo("internal error: %T", x))
		}
	}
}

func (c *cpp) elifGroup(eg elifGroup) bool {
	//	# elif constant-expression new-line group_opt
	ln := eg.line[:len(eg.line)-1] // Remove new-line
	if len(ln) < 3 {
		c.eh("%v: expected expression", ln[1].Position())
		return false
	}

	return c.isNonZero(c.eval(ln[2:]))
}

// include executes an #include control-line, [0]6.10.
func (c *cpp) include(ln controlLine) {
	// eg. ["#" "include" "<stdio.h>" "\n"]
	if len(ln) < 3 {
		c.eh("%v: missing argument", ln[1].Position())
		return
	}

	if c.includeLevel == maxIncludeLevel {
		c.eh("%v: too many include levels", ln[1].Position())
		return
	}

	c.includeLevel++
	defer func() { c.includeLevel-- }()

	s := preprocessingTokens(tokens2PreprocessingTokens(ln[2:], true))
	s = c.expand(false, true, &s)
	if c.cfg.fakeIncludes {
		t := s[0].Token
		t.Set(nil, t.Src())
		c.push(textLine([]Token{
			t,
			ln[len(ln)-1],
		}))
		return
	}

	switch raw := string(s[0].Src()); {
	case strings.HasPrefix(raw, `"`):
		nm := raw[1 : len(raw)-1]
		for _, v := range c.cfg.IncludePaths {
			if v == "" {
				v, _ = filepath.Split(c.srcStack[len(c.srcStack)-1])
			}
			pth := filepath.Join(v, nm)
			if g, err := c.group(Source{pth, nil, c.cfg.FS}); err == nil {
				c.push(g)
				return
			}
		}

		c.eh("%v: include file not found: %s", s[0].Position(), raw)
	case strings.HasPrefix(raw, `<`):
		nm := raw[1 : len(raw)-1]
		for _, v := range c.cfg.SysIncludePaths {
			pth := filepath.Join(v, nm)
			if g, err := c.group(Source{pth, nil, c.cfg.FS}); err == nil {
				c.push(g)
				return
			}
		}

		c.eh("%v: include file not found: %s", s[0].Position(), raw)
	default:
		c.eh("%v: invalid argument", s[0].Position())
	}
}

func (c *cpp) sysInclue(n Node, nm string) {
	var err error
	var g group
	for _, v := range c.cfg.SysIncludePaths {
		pth := filepath.Join(v, nm)
		if g, err = c.group(Source{pth, nil, c.cfg.FS}); err == nil {
			c.push(g)
			return
		}
	}

	c.eh("%v: include file not found: %s", n.Position(), nm)
}

func (c *cpp) ifGroup(ig *ifGroup) bool {
	ln := ig.line[:len(ig.line)-1] // Remove new-line
	switch string(ln[1].Src()) {
	case "ifdef":
		if len(ln) < 3 { // '#' "ifdef" IDENTIFIER
			c.eh("%v: expected identifier", ln[1].Position())
			return false
		}

		t := ln[2]
		return c.macro(&t, string(t.Src())) != nil
	case "ifndef":
		if len(ln) < 3 { // '#' "ifndef" IDENTIFIER
			c.eh("%v: expected identifier", ln[1].Position())
			return false
		}

		t := ln[2]
		return c.macro(&t, string(t.Src())) == nil
	case "if":
		if len(ln) < 3 { // '#' "if" <expr>
			c.eh("%v: expected expression", ln[1].Position())
			return false
		}

		return c.isNonZero(c.eval(ln[2:]))
	default:
		panic(todo("", toksDump(ln)))
	}
}

func (c *cpp) eval(s0 []Token) interface{} {
	s := preprocessingTokens(tokens2PreprocessingTokens(s0, false))
	s = c.expand(false, true, &s)
	p := &s
	val := c.expression(p, true)
	switch t := p.c(); t.Ch {
	case eof, '#':
		// ok
	default:
		c.eh("%v: unexpected %s", t.Position(), runeName(t.Ch))
		return nil
	}

	return val
}

// [0], 6.5.17 Comma operator
//
//  expression:
// 	assignment-expression
// 	expression , assignment-expression
func (c *cpp) expression(s *preprocessingTokens, eval bool) interface{} {
	for {
		r := c.assignmentExpression(s, eval)
		if s.c().Ch != ',' {
			return r
		}

		s.read()
	}
}

// [0], 6.5.16 Assignment operators
//
//  assignment-expression:
// 	conditional-expression
// 	unary-expression assignment-operator assignment-expression
//
//  assignment-operator: one of
// 	= *= /= %= += -= <<= >>= &= ^= |=
func (c *cpp) assignmentExpression(s *preprocessingTokens, eval bool) interface{} {
	return c.conditionalExpression(s, eval)
}

// [0], 6.5.15 Conditional operator
//
//  conditional-expression:
//		logical-OR-expression
//		logical-OR-expression ? expression : conditional-expression
func (c *cpp) conditionalExpression(s *preprocessingTokens, eval bool) interface{} {
	expr := c.logicalOrExpression(s, eval)
	if s.c().Ch == '?' {
		s.read()
		exprIsNonZero := c.isNonZero(expr)
		expr2 := c.conditionalExpression(s, exprIsNonZero)
		if t := s.c(); t.Ch != ':' {
			c.eh("%v: expected ':'", t.Position())
			return expr
		}

		s.read()
		expr3 := c.conditionalExpression(s, !exprIsNonZero)

		// [0] 6.5.15
		//
		// 5. If both the second and third operands have arithmetic type, the result
		// type that would be determined by the usual arithmetic conversions, were they
		// applied to those two operands, is the type of the result.
		switch x := expr2.(type) {
		case int64:
			switch expr3.(type) {
			case uint64:
				expr2 = uint64(x)
			}
		case uint64:
			switch y := expr3.(type) {
			case int64:
				expr3 = uint64(y)
			}
		}
		switch {
		case exprIsNonZero:
			expr = expr2
		default:
			expr = expr3
		}
	}
	return expr
}

// [0], 6.5.14 Logical OR operator
//
//  logical-OR-expression:
//		logical-AND-expression
//		logical-OR-expression || logical-AND-expression
func (c *cpp) logicalOrExpression(s *preprocessingTokens, eval bool) interface{} {
	lhs := c.logicalAndExpression(s, eval)
	for s.c().Ch == rune(OROR) {
		s.read()
		if c.isNonZero(lhs) {
			eval = false
		}
		rhs := c.logicalAndExpression(s, eval)
		if c.isNonZero(lhs) || c.isNonZero(rhs) {
			lhs = int64(1)
		}
	}
	return lhs
}

// [0], 6.5.13 Logical AND operator
//
//  logical-AND-expression:
//		inclusive-OR-expression
//		logical-AND-expression && inclusive-OR-expression
func (c *cpp) logicalAndExpression(s *preprocessingTokens, eval bool) interface{} {
	lhs := c.inclusiveOrExpression(s, eval)
	for s.c().Ch == rune(ANDAND) {
		s.read()
		if c.isZero(lhs) {
			eval = false
		}
		rhs := c.inclusiveOrExpression(s, eval)
		if c.isZero(lhs) || c.isZero(rhs) {
			lhs = int64(0)
		}
	}
	return lhs
}

// [0], 6.5.12 Bitwise inclusive OR operator
//
//  inclusive-OR-expression:
//		exclusive-OR-expression
//		inclusive-OR-expression | exclusive-OR-expression
func (c *cpp) inclusiveOrExpression(s *preprocessingTokens, eval bool) interface{} {
	lhs := c.exclusiveOrExpression(s, eval)
	for s.c().Ch == '|' {
		panic(todo(""))
		// s.next()
		// rhs := c.exclusiveOrExpression(s, eval)
		// if eval {
		// 	switch x := lhs.(type) {
		// 	case int64:
		// 		switch y := rhs.(type) {
		// 		case int64:
		// 			lhs = x | y
		// 		case uint64:
		// 			lhs = uint64(x) | y
		// 		}
		// 	case uint64:
		// 		switch y := rhs.(type) {
		// 		case int64:
		// 			lhs = x | uint64(y)
		// 		case uint64:
		// 			lhs = x | y
		// 		}
		// 	}
		// }
	}
	return lhs
}

// [0], 6.5.11 Bitwise exclusive OR operator
//
//  exclusive-OR-expression:
//		AND-expression
//		exclusive-OR-expression ^ AND-expression
func (c *cpp) exclusiveOrExpression(s *preprocessingTokens, eval bool) interface{} {
	lhs := c.andExpression(s, eval)
	for s.c().Ch == '^' {
		panic(todo(""))
		// s.next()
		// rhs := c.andExpression(s, eval)
		// if eval {
		// 	switch x := lhs.(type) {
		// 	case int64:
		// 		switch y := rhs.(type) {
		// 		case int64:
		// 			lhs = x ^ y
		// 		case uint64:
		// 			lhs = uint64(x) ^ y
		// 		}
		// 	case uint64:
		// 		switch y := rhs.(type) {
		// 		case int64:
		// 			lhs = x ^ uint64(y)
		// 		case uint64:
		// 			lhs = x ^ y
		// 		}
		// 	}
		// }
	}
	return lhs
}

// [0], 6.5.10 Bitwise AND operator
//
//  AND-expression:
// 		equality-expression
// 		AND-expression & equality-expression
func (c *cpp) andExpression(s *preprocessingTokens, eval bool) interface{} {
	lhs := c.equalityExpression(s, eval)
	for s.c().Ch == '&' {
		panic(todo(""))
		// s.next()
		// rhs := c.equalityExpression(s, eval)
		// if eval {
		// 	switch x := lhs.(type) {
		// 	case int64:
		// 		switch y := rhs.(type) {
		// 		case int64:
		// 			lhs = x & y
		// 		case uint64:
		// 			lhs = uint64(x) & y
		// 		}
		// 	case uint64:
		// 		switch y := rhs.(type) {
		// 		case int64:
		// 			lhs = x & uint64(y)
		// 		case uint64:
		// 			lhs = x & y
		// 		}
		// 	}
		// }
	}
	return lhs
}

// [0], 6.5.9 Equality operators
//
//  equality-expression:
//		relational-expression
//		equality-expression == relational-expression
//		equality-expression != relational-expression
func (c *cpp) equalityExpression(s *preprocessingTokens, eval bool) interface{} {
	lhs := c.relationalExpression(s, eval)
	for {
		var v bool
		switch s.c().Ch {
		case rune(EQ):
			s.read()
			rhs := c.relationalExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						v = x == y
					case uint64:
						v = uint64(x) == y
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						v = x == uint64(y)
					case uint64:
						v = x == y
					}
				}
			}
		case rune(NEQ):
			s.read()
			rhs := c.relationalExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						v = x != y
					case uint64:
						v = uint64(x) != y
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						v = x != uint64(y)
					case uint64:
						v = x != y
					}
				}
			}
		default:
			return lhs
		}
		switch {
		case v:
			lhs = int64(1)
		default:
			lhs = int64(0)
		}
	}
}

// [0], 6.5.8 Relational operators
//
//  relational-expression:
//		shift-expression
//		relational-expression <  shift-expression
//		relational-expression >  shift-expression
//		relational-expression <= shift-expression
//		relational-expression >= shift-expression
func (c *cpp) relationalExpression(s *preprocessingTokens, eval bool) interface{} {
	lhs := c.shiftExpression(s, eval)
	for {
		var v bool
		switch s.c().Ch {
		case '<':
			s.read()
			rhs := c.shiftExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						v = x < y
					case uint64:
						v = uint64(x) < y
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						v = x < uint64(y)
					case uint64:
						v = x < y
					}
				}
			}
		case '>':
			s.read()
			rhs := c.shiftExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						v = x > y
					case uint64:
						v = uint64(x) > y
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						v = x > uint64(y)
					case uint64:
						v = x > y
					}
				}
			}
		case rune(LEQ):
			panic(todo(""))
			// s.next()
			// rhs := c.shiftExpression(s, eval)
			// if eval {
			// 	switch x := lhs.(type) {
			// 	case int64:
			// 		switch y := rhs.(type) {
			// 		case int64:
			// 			v = x <= y
			// 		case uint64:
			// 			v = uint64(x) <= y
			// 		}
			// 	case uint64:
			// 		switch y := rhs.(type) {
			// 		case int64:
			// 			v = x <= uint64(y)
			// 		case uint64:
			// 			v = x <= y
			// 		}
			// 	}
			// }
		case rune(GEQ):
			s.read()
			rhs := c.shiftExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						v = x >= y
					case uint64:
						v = uint64(x) >= y
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						v = x >= uint64(y)
					case uint64:
						v = x >= y
					}
				}
			}
		default:
			return lhs
		}
		switch {
		case v:
			lhs = int64(1)
		default:
			lhs = int64(0)
		}
	}
}

// [0], 6.5.7 Bitwise shift operators
//
//  shift-expression:
//		additive-expression
//		shift-expression << additive-expression
//		shift-expression >> additive-expression
func (c *cpp) shiftExpression(s *preprocessingTokens, eval bool) interface{} {
	lhs := c.additiveExpression(s, eval)
	for {
		switch s.c().Ch {
		case rune(LSH):
			s.read()
			rhs := c.additiveExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						lhs = x << uint(y)
					case uint64:
						lhs = uint64(x) << uint(y)
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						lhs = x << uint(y)
					case uint64:
						lhs = x << uint(y)
					}
				}
			}
		case rune(RSH):
			panic(todo(""))
			// s.next()
			// rhs := c.additiveExpression(s, eval)
			// if eval {
			// 	switch x := lhs.(type) {
			// 	case int64:
			// 		switch y := rhs.(type) {
			// 		case int64:
			// 			lhs = x >> uint(y)
			// 		case uint64:
			// 			lhs = uint64(x) >> uint(y)
			// 		}
			// 	case uint64:
			// 		switch y := rhs.(type) {
			// 		case int64:
			// 			lhs = x >> uint(y)
			// 		case uint64:
			// 			lhs = x >> uint(y)
			// 		}
			// 	}
			// }
		default:
			return lhs
		}
	}
}

// [0], 6.5.6 Additive operators
//
//  additive-expression:
//		multiplicative-expression
//		additive-expression + multiplicative-expression
//		additive-expression - multiplicative-expression
func (c *cpp) additiveExpression(s *preprocessingTokens, eval bool) interface{} {
	lhs := c.multiplicativeExpression(s, eval)
	for {
		switch s.c().Ch {
		case '+':
			s.read()
			rhs := c.multiplicativeExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						lhs = x + y
					case uint64:
						lhs = uint64(x) + y
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						lhs = x + uint64(y)
					case uint64:
						lhs = x + y
					}
				}
			}
		case '-':
			s.read()
			rhs := c.multiplicativeExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						lhs = x - y
					case uint64:
						lhs = uint64(x) - y
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						lhs = x - uint64(y)
					case uint64:
						lhs = x - y
					}
				}
			}
		default:
			return lhs
		}
	}
}

// [0], 6.5.5 Multiplicative operators
//
//  multiplicative-expression:
//		unary-expression // [0], 6.10.1, 1.
//		multiplicative-expression * unary-expression
//		multiplicative-expression / unary-expression
//		multiplicative-expression % unary-expression
func (c *cpp) multiplicativeExpression(s *preprocessingTokens, eval bool) interface{} {
	lhs := c.unaryExpression(s, eval)
	for {
		switch s.c().Ch {
		case '*':
			panic(todo(""))
			// s.next()
			// rhs := c.unaryExpression(s, eval)
			// if eval {
			// 	switch x := lhs.(type) {
			// 	case int64:
			// 		switch y := rhs.(type) {
			// 		case int64:
			// 			lhs = x * y
			// 		case uint64:
			// 			lhs = uint64(x) * y
			// 		}
			// 	case uint64:
			// 		switch y := rhs.(type) {
			// 		case int64:
			// 			lhs = x * uint64(y)
			// 		case uint64:
			// 			lhs = x * y
			// 		}
			// 	}
			// }
		case '/':
			t := s.read()
			rhs := c.unaryExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						if y == 0 {
							c.eh("%v: division by zero", t.Position())
							break
						}

						lhs = x / y
					case uint64:
						if y == 0 {
							c.eh("%v: division by zero", t.Position())
							break
						}

						lhs = uint64(x) / y
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						if y == 0 {
							c.eh("%v: division by zero", t.Position())
							break
						}

						lhs = x / uint64(y)
					case uint64:
						if y == 0 {
							c.eh("%v: division by zero", t.Position())
							break
						}

						lhs = x / y
					}
				}
			}
		case '%':
			panic(todo(""))
			// t := s.next()
			// rhs := c.unaryExpression(s, eval)
			// if eval {
			// 	switch x := lhs.(type) {
			// 	case int64:
			// 		switch y := rhs.(type) {
			// 		case int64:
			// 			if y == 0 {
			// 				c.err(t, "division by zero")
			// 				break
			// 			}

			// 			lhs = x % y
			// 		case uint64:
			// 			if y == 0 {
			// 				c.err(t, "division by zero")
			// 				break
			// 			}

			// 			lhs = uint64(x) % y
			// 		}
			// 	case uint64:
			// 		switch y := rhs.(type) {
			// 		case int64:
			// 			if y == 0 {
			// 				c.err(t, "division by zero")
			// 				break
			// 			}

			// 			lhs = x % uint64(y)
			// 		case uint64:
			// 			if y == 0 {
			// 				c.err(t, "division by zero")
			// 				break
			// 			}

			// 			lhs = x % y
			// 		}
			// 	}
			// }
		default:
			return lhs
		}
	}
}

// [0], 6.5.3 Unary operators
//
//  unary-expression:
//		primary-expression
//		unary-operator unary-expression
//
//  unary-operator: one of
//		+ - ~ !
func (c *cpp) unaryExpression(s *preprocessingTokens, eval bool) interface{} {
	switch s.c().Ch {
	case '+':
		panic(todo(""))
		// s.next()
		// return c.unaryExpression(s, eval)
	case '-':
		s.read()
		expr := c.unaryExpression(s, eval)
		if eval {
			switch x := expr.(type) {
			case int64:
				expr = -x
			case uint64:
				expr = -x
			}
		}
		return expr
	case '~':
		panic(todo(""))
		// s.next()
		// expr := c.unaryExpression(s, eval)
		// if eval {
		// 	switch x := expr.(type) {
		// 	case int64:
		// 		expr = ^x
		// 	case uint64:
		// 		expr = ^x
		// 	}
		// }
		// return expr
	case '!':
		s.read()
		expr := c.unaryExpression(s, eval)
		if eval {
			var v bool
			switch x := expr.(type) {
			case int64:
				v = x == 0
			case uint64:
				v = x == 0
			}
			switch {
			case v:
				expr = int64(1)
			default:
				expr = int64(0)
			}
		}
		return expr
	default:
		return c.primaryExpression(s, eval)
	}
}

// [0], 6.5.1 Primary expressions
//
//  primary-expression:
//		identifier
//		constant
//		( expression )
func (c *cpp) primaryExpression(s *preprocessingTokens, eval bool) interface{} {
	switch t := s.c(); t.Ch {
	case rune(CHARCONST), rune(LONGCHARCONST):
		panic(todo(""))
		// s.next()
		// r := charConst(c.ctx, t)
		// return int64(r)
	case rune(IDENTIFIER):
		s.read()
		// if s.peek().char == '(' {
		// 	s.next()
		// 	n := 1
		// loop:
		// 	for n != 0 {
		// 		switch s.peek().char {
		// 		case '(':
		// 			n++
		// 		case ')':
		// 			n--
		// 		case -1:
		// 			c.err(s.peek(), "expected )")
		// 			break loop
		// 		}
		// 		s.next()
		// 	}
		// }
		return int64(0)
	case rune(PPNUMBER):
		s.read()
		return c.intConst(t)
	case '(':
		s.read()
		expr := c.expression(s, eval)
		if s.c().Ch == ')' {
			s.read()
		} else {
			panic(todo(""))
		}
		return expr
	default:
		panic(todo(""))
		// return int64(0)
	}
}

// [0], 6.4.4.1 Integer constants
//
//  integer-constant:
//		decimal-constant integer-suffix_opt
//		octal-constant integer-suffix_opt
//		hexadecimal-constant integer-suffix_opt
//
//  decimal-constant:
//		nonzero-digit
//		decimal-constant digit
//
//  octal-constant:
//		0
//		octal-constant octal-digit
//
//  hexadecimal-prefix: one of
//		0x 0X
//
//  integer-suffix_opt: one of
//		u ul ull l lu ll llu
func (c *cpp) intConst(t Token) (r interface{}) {
	var n uint64
	s0 := string(t.Src())
	s := strings.TrimRight(s0, "uUlL")
	var err error
	switch {
	case strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X"):
		if n, err = strconv.ParseUint(s[2:], 16, 64); err != nil {
			c.eh("%v: %v", t.Position(), err)
			return int64(0)
		}
	case strings.HasPrefix(s, "0"):
		if n, err = strconv.ParseUint(s, 8, 64); err != nil {
			c.eh("%v: %v", t.Position(), err)
			return int64(0)
		}
	default:
		if n, err = strconv.ParseUint(s, 10, 64); err != nil {
			c.eh("%v: %v", t.Position(), err)
			return int64(0)
		}
	}

	switch suffix := strings.ToLower(s0[len(s):]); suffix {
	case "", "l", "ll":
		if n > math.MaxInt64 {
			return n
		}

		return int64(n)
	default:
		panic(todo(""))
		// c.err(t, "invalid suffix: %v", s0)
		// fallthrough
	case "llu", "lu", "u", "ul", "ull":
		return n
	}
}

func (c *cpp) isZero(val interface{}) bool {
	switch x := val.(type) {
	case int64:
		return x == 0
	case uint64:
		return x == 0
	default:
		c.eh("", errorf("internal error: %T", x))
		return false
	}
}

func (c *cpp) isNonZero(val interface{}) bool {
	switch x := val.(type) {
	case nil:
		return false
	case int64:
		return x != 0
	case uint64:
		return x != 0
	default:
		c.eh("", errorf("internal error: %T", x))
		return false
	}
}

// undef executes an #undef control-line, [0]6.10.
func (c *cpp) undef(ln controlLine) {
	// eg. ["#" "undef" "  x" "      \n"]
	if len(ln) < 3 {
		return
	}

	nm := string(ln[2].Src())
	if _, ok := protectedMacros[nm]; !ok {
		delete(c.macros, nm)
	}
}

// define executes a #define control-line, [0]6.10.
func (c *cpp) define(ln controlLine) {
	def := ln[1]
	ln = ln[2:] // Remove '#' and "define"
	if len(ln) == 0 {
		c.eh("%v: missing macro name", def.Position())
		return
	}

	nm := ln[0]
	ln = ln[1:]
	switch {
	case ln[0].Ch == '(' && len(ln[0].Sep()) == 0:
		// lparen: a ( character not immediately preceded by white-space
		// # define identifier ( args_opt ) replacement-list new-line
		//                     ^ln[0]
		c.defineFnMacro(nm, ln[:len(ln)-1]) // strip new-line
	default:
		// # define identifier replacement-list new-line
		//                     ^ln[0]
		c.defineObjectMacro(nm, ln[:len(ln)-1]) // strip new-line
	}
}

func (c *cpp) defineFnMacro(nm Token, ln []Token) {
	fp, rl0, minArgs, varArg := c.parseMacroParams(ln)
	rl := c.parseReplacementList(rl0)
	c.newMacro(nm, fp, rl, minArgs, varArg, true)
}

func (c *cpp) parseMacroParams(ln []Token) (fp, rl []Token, minArgs, vaArg int) {
	if len(ln) == 0 || ln[0].Ch != '(' {
		c.eh("internal error")
		return nil, nil, -1, -1
	}

	lpar := ln[0]
	ln = ln[1:] // remove '('
	vaArg = -1
	for {
		if len(ln) == 0 {
			// (A)
			c.eh("%v: macro paramater list is missing final ')'", lpar.Position())
			return nil, nil, -1, -1
		}

		switch ln[0].Ch {
		case ')':
			// (B)
			ln = ln[1:]
			return fp, ln, minArgs, vaArg
		case rune(IDENTIFIER):
			fp = append(fp, ln[0])
			minArgs++
			ln = ln[1:]
			if len(ln) == 0 {
				break // -> (A)
			}

			switch ln[0].Ch {
			case ')':
				// ok -> (B)
			case ',':
				ln = ln[1:]
			case rune(DDD):
				if vaArg >= 0 {
					c.eh("%v: multiple var arguments", ln[0].Position())
					return nil, nil, -1, -1
				}

				vaArg = len(fp) - 1
				ln = ln[1:]
			default:
				c.eh("%v: unexpected %v", ln[0].Position(), runeName(ln[0].Ch))
				return nil, nil, -1, -1
			}
		case rune(DDD):
			if vaArg >= 0 {
				c.eh("%v: multiple var arguments", ln[0].Position())
				return nil, nil, -1, -1
			}

			vaArg = len(fp)
			ln = ln[1:]
			if len(ln) == 0 {
				break // -> (A)
			}

			switch ln[0].Ch {
			case ')':
				// ok -> (B)
			default:
				ln = nil // -> (A)
			}
		default:
			c.eh("%v: unexpected %v", ln[0].Position(), runeName(ln[0].Ch))
			return nil, nil, -1, -1
		}
	}
}

func (c *cpp) defineObjectMacro(nm Token, ln []Token) {
	rl := c.parseReplacementList(ln)
	c.newMacro(nm, nil, rl, -1, -1, false)
}

func (c *cpp) newMacro(nm Token, params []Token, replList []preprocessingToken, minArgs, varArg int, isFnLike bool) {
	// trc("nm %q, params %v, replList %v, minArgs %v, varArg %v, isFnLike %v", nm.Src(), toksDump(params), toksDump(replList), minArgs, varArg, isFnLike)
	s := string(nm.Src())
	if _, ok := protectedMacros[s]; ok {
		if nm.Position().Filename != "<predefined>" {
			return
		}
	}

	m, err := newMacro(nm, params, replList, minArgs, varArg, isFnLike)
	if err != nil {
		c.eh("%v", err)
		return
	}

	c.macros[s] = m
}

// parseReplacementList transforms s into preprocessing tokens that have separate
// tokens for white space.
func (c *cpp) parseReplacementList(s []Token) (r []preprocessingToken) {
	return tokens2PreprocessingTokens(c.toksTrim(s), true)
}

func (c *cpp) toksTrim(s []Token) []Token {
	for len(s) != 0 && s[0].Ch == ' ' {
		s = s[1:]
	}
	for len(s) != 0 && s[len(s)-1].Ch == ' ' {
		s = s[:len(s)-1]
	}
	return s
}

func (c *cpp) pop() {
	if n := len(c.stack); n != 0 {
		c.tos = c.stack[n-1]
		c.stack = c.stack[:n-1]
		return
	}

	c.tos = nil
}

func (c *cpp) push(v interface{}) {
	if c.tos != nil {
		c.stack = append(c.stack, c.tos)
	}
	c.tos = v
}

func (c *cpp) group(src Source) (group, error) {
	c.srcStack = append(c.srcStack, src.Name)
	if g, ok := c.groups[src.Name]; ok {
		return g, nil
	}

	p, err := newCppParser(src, c.eh)
	if err != nil {
		return nil, err
	}

	g := p.group(false)
	c.groups[src.Name] = g
	return g, nil
}
