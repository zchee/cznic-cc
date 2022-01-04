// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"fmt"
	"strings"
)

var (
	_ tokenSequence = (*cat)(nil)
	_ tokenSequence = (*preprocessingTokens)(nil)
	_ tokenSequence = (*tokenizer)(nil)

	comma              = []byte{','}
	eofPreprocessToken preprocessingToken
	eofTok             Token
	sp                 = []byte{' '}
	spTok              Token
)

func init() {
	s, err := newScannerSource(Source{"", []byte(nil)})
	if err != nil {
		panic(errorf("initialization error"))
	}

	spTok.s = s
	spTok.Ch = ' '
	spTok.Set(nil, sp)
	eofTok.s = s
	eofTok.Ch = eof
	eofTok.Set(nil, nil)
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

type formalParams []string

func (fp formalParams) is(s string) int {
	for i, v := range fp {
		if s == v {
			return i
		}
	}

	return -1
}

// Macro represents a preprocessor #define.
type Macro struct {
	Name            Token
	Params          []Token
	params          formalParams
	replacementList preprocessingTokens

	MinArgs int // m x: 0, m() x: 0, m(...): 0, m(a) a: 1, m(a, ...): 1, m(a, b): 2, m(a, b, ...): 2.
	VarArg  int // m(a): -1, m(...): 0, m(a, ...): 1, m(a...): 0, m(a, b...): 1.

	IsFnLike bool // m: false, m(): true, m(x): true.
}

func newMacro(nm Token, params []Token, replList []preprocessingToken, minArgs, varArg int, isFnLike bool) (*Macro, error) {
	var fp formalParams
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

// ReplacementList returns the tokens m is substitued with.
func (m *Macro) ReplacementList() (r []Token) {
	for _, v := range m.replacementList {
		if v.Ch != ' ' {
			r = append(r, v.Token)
		}
	}
	return r
}

func (m *Macro) fp() formalParams {
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
	peek(int) Token
	read() preprocessingToken
}

type tokenizer struct {
	c   *cpp
	in  []preprocessingToken
	out []preprocessingToken
}

func newTokenizer(c *cpp) *tokenizer {
	return &tokenizer{c: c}
}

func (t *tokenizer) peek(index int) (tok Token) {
	if index < len(t.in) {
		return t.in[index].Token
	}

	t.c.eh("%v", errorf("internal error"))
	return eofTok
}

func (t *tokenizer) read() (tok preprocessingToken) {
	if len(t.in) == 0 {
		t.in = tokens2preprocessingTokens(t.c.nextLine())
	}

	tok = t.in[0]
	if tok.Ch != eof {
		t.in = t.in[1:]
	}
	return tok

}

func (t *tokenizer) token() (tok Token) {
	for len(t.out) == 0 {
		t.out = t.c.expand(true, t)
	}
	tok = t.out[0].Token
	if tok.Ch != eof {
		t.out = t.out[1:]
	}
	return tok
}

type preprocessingTokens []preprocessingToken

func (p *preprocessingTokens) peek(index int) (tok Token) {
	if index < len(*p) {
		return (*p)[index].Token
	}

	tok.Ch = eof
	return tok
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
	// case *tokens:
	// 	if len(*x) == 0 {
	// 		return &head
	// 	}

	// 	return &cat{head, x}
	default:
		panic(todo("%T", x))
	}
}

func (c *cat) peek(index int) (tok Token) {
	if index < len(c.head) {
		return c.head.peek(index)
	}

	return c.tail.peek(index - len(c.head))
}

func (c *cat) read() preprocessingToken {
	if len(c.head) != 0 {
		return c.head.read()
	}

	return c.tail.read()
}

// cpp is the C preprocessor.
type cpp struct {
	eh          errHandler
	groups      map[string]group
	indentLevel int // debug dumps
	macros      map[string]*Macro
	sources     []Source
	stack       []interface{}
	tok         Token
	tokenizer   *tokenizer
	tos         interface{}

	closed bool
}

// newCPP returns a newly created cpp.
func newCPP(sources []Source, eh errHandler) (*cpp, error) {
	m := map[string]struct{}{}
	for _, v := range sources {
		if _, ok := m[v.Name]; ok {
			return nil, errorf("duplicate source name: %q", v.Name)
		}

		m[v.Name] = struct{}{}
	}
	c := &cpp{
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
func (c *cpp) expand(outer bool, TS tokenSequence) (r preprocessingTokens) {
	// trc("  %s%v (%v)", c.indent(), toksDump(TS), origin(2))
	// defer func() { trc("->%s%v", c.undent(), toksDump(r)) }()
	t := TS.read()
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

		return append(preprocessingTokens{t}, c.expand(false, TS)...)
	}

	if m := c.macros[src]; m != nil {
	out:
		switch {
		default:
			// if TS is T^HS • TS’ and T is a "()-less macro" then
			//	return expand(subst(ts(T),{},{},HS∪{T},{}) • TS’);

			// trc("  %s<%s is a ()-less macro, expanding to %s>", c.indent(), T.Src(), toksDump(m.replacementList))
			// defer func() { trc("  %s<%s expanded>", c.undent(), T.Src()) }()
			return c.expand(false, newCat(c.subst(m, m.ts(), nil, nil, HS.add(src), nil), TS))
		case m.IsFnLike:
			t2 := TS.peek(0)
			if t2.Ch == ' ' {
				TS.read()
				t2 = TS.peek(0)
			}
			if t2.Ch == '(' {
				// if TS is T^HS • ( • TS’ and T is a "()’d macro" then
				//	check TS’ is actuals • )^HS’ • TS’’ and actuals are "correct for T"
				//	return expand(subst(ts(T),fp(T),actuals,(HS∩HS’)∪{T},{}) • TS’’);

				// trc("  %s<%s is a ()'d macro, expanding to %s>", c.indent(), T.Src(), toksDump(m.replacementList))
				// defer func() { trc("  %s<%s expanded>", c.undent(), T.Src()) }()
				TS.read() // '('
				args, rparen, ok := c.parseMacroArgs(TS)
				if !ok {
					return nil
				}

				switch {
				case len(args) < m.MinArgs:
					c.eh("%v: not enough macro arguments", rparen.Position())
					break out
				case len(args) > m.MinArgs && m.VarArg < 0:
					c.eh("%v: too many macro arguments", rparen.Position())
					break out
				}

				return c.expand(false, newCat(c.subst(m, m.ts(), m.fp(), args, HS.cap(rparen.hs).add(src), nil), TS))
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

	return append(preprocessingTokens{t}, c.expand(false, TS)...)
}

// [1], pg 2.
func (c *cpp) subst(m *Macro, IS tokenSequence, FP formalParams, AP []preprocessingTokens, HS hideSet, OS preprocessingTokens) (r preprocessingTokens) {
	// trc("  %s%v, HS %v, FP %v, AP %v, OS %v (%v)", c.indent(), toksDump(IS), &HS, FP, toksDump(AP), toksDump(OS), origin(2))
	// defer func() { trc("->%s%v", c.undent(), toksDump(r)) }()
	t := IS.read()
	if t.Ch == eof {
		// if IS is {} then
		//	return hsadd(HS,OS);
		return c.hsAdd(HS, OS)
	}

	if t.Ch == '#' {
		t2 := IS.read()
		if i := FP.is(string(t2.Src())); i >= 0 {
			// if IS is # • T • IS’ and T is FP[i] then
			//	return subst(IS’,FP,AP,HS,OS • stringize(select(i,AP )));
			return c.subst(m, IS, FP, AP, HS, append(OS, c.stringize(c.apSelect(m, t, AP, i))))
		}

		panic(todo("", &t2))
	}

	if t.Ch == rune(PPPASTE) {
		t2 := IS.read()
		if i := FP.is(string(t2.Src())); i >= 0 {
			// if IS is ## • T • IS’ and T is FP[i] then
			if len(AP[i]) == 0 {
				//	if select(i,AP ) is {} then /* only if actuals can be empty */
				//		return subst(IS’,FP,AP,HS,OS );
				panic(todo("%v %v %v", toksDump(IS), &t2, i))
			} else {
				//	else
				//		return subst(IS’,FP,AP,HS,glue(OS,select(i,AP )));
				return c.subst(m, IS, FP, AP, HS, c.glue(OS, c.apSelect(m, t2, AP, i)))
			}
		}

		// else if IS is ## • T HS’ • IS’ then
		//	return subst(IS’,FP,AP,HS,glue(OS,T^HS’));
		return c.subst(m, IS, FP, AP, HS, c.glue(OS, preprocessingTokens{t2}))
	}

	if IS.peek(0).Ch == rune(PPPASTE) {
		if t.Ch == rune(IDENTIFIER) {
			if i := FP.is(string(t.Src())); i >= 0 {
				// if IS is T • ##^HS’ • IS’ and T is FP[i] then
				//	if select(i,AP ) is {} then /* only if actuals can be empty */
				if len(AP[i]) == 0 {
					//	{
					//		if IS’ is T’ • IS’’ and T’ is FP[j] then
					//			return subst(IS’’,FP,AP,HS,OS • select(j,AP));
					//		else
					//			return subst(IS’,FP,AP,HS,OS);
					//	}
					panic(todo("", toksDump(IS)))
				} else {
					//	else
					//		return subst(##^HS’ • IS’,FP,AP,HS,OS • select(i,AP));
					return c.subst(m, IS, FP, AP, HS, append(OS, c.apSelect(m, t, AP, i)...))
				}
			}
		}
	}

	if len(FP) != 0 {
		if i := FP.is(string(t.Src())); i >= 0 {
			// if IS is T • IS’ and T is FP[i] then
			//	return subst(IS’,FP,AP,HS,OS • expand(select(i,AP )));
			return c.subst(m, IS, FP, AP, HS, append(OS, c.expand(false, c.apSelectP(m, t, AP, i))...))
		}
	}

	if va := m.VarArg; t.Ch == rune(IDENTIFIER) && va >= 0 && string(t.Src()) == "__VA_ARGS__" && va <= len(AP) {
		return c.subst(m, IS, FP, AP, HS, append(OS, c.expand(false, c.varArgsP(t, AP[va:]))...))
	}

	// note IS must be T HS’ • IS’
	// return subst(IS’,FP,AP,HS,OS • T HS’ );
	return c.subst(m, IS, FP, AP, HS, append(OS, t))

}

func (c *cpp) glue(LS, RS preprocessingTokens) (r preprocessingTokens) {
	// trc("  %sLS %v, RS %v (%v)", c.indent(), toksDump(LS), toksDump(RS), origin(2))
	// defer func() { trc("->%s%v", c.undent(), toksDump(r)) }()
	// if LS is L^HS and RS is R^HS’ • RS’ then
	//	return L&R^(HS∩HS’) • RS’; /* undefined if L&R is invalid */
	if len(LS) == 1 {
		tok := LS[0]
		tok.Set(nil, append(tok.Src(), RS[0].Src()...))
		tok.hs = tok.hs.cap(RS[0].hs)
		return append(preprocessingTokens{tok}, RS[1:]...)
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
		c.eh("", errorf("internal error"))
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

func (c *cpp) apSelectP(m *Macro, t preprocessingToken, AP []preprocessingTokens, i int) *preprocessingTokens {
	r := c.apSelect(m, t, AP, i)
	return &r
}

func (c *cpp) apSelect(m *Macro, t preprocessingToken, AP []preprocessingTokens, i int) preprocessingTokens {
	if m.VarArg < 0 || m.VarArg != i {
		return AP[i]
	}

	return c.varArgs(t, AP[i:])
}

func (c *cpp) varArgsP(t preprocessingToken, AP []preprocessingTokens) *preprocessingTokens {
	r := c.varArgs(t, AP)
	return &r
}

func (c *cpp) varArgs(t preprocessingToken, AP []preprocessingTokens) preprocessingTokens {
	var commaSP preprocessingTokens
	if len(AP) > 1 {
		t.Set(nil, comma)
		t.Ch = ','
		t2 := t
		t2.Set(nil, sp)
		t2.Ch = ' '
		commaSP = preprocessingTokens{t, t2}
	}
	var a preprocessingTokens
	for i, v := range AP {
		if i != 0 {
			a = append(a, commaSP...)
		}
		a = append(a, v...)
	}
	return a
}

func (c *cpp) parseMacroArgs(TS tokenSequence) (args []preprocessingTokens, rparen preprocessingToken, ok bool) {
	// trc("  %sin  %v (%v)", c.indent(), toksDump(TS), origin(2))
	// defer func() { trc("->%sout %v", c.undent(), toksDump(args)) }()
	var arg preprocessingTokens
	level := 0
	var t preprocessingToken
out:
	for {
		switch t = TS.read(); t.Ch {
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
		args[i] = normalizeHashWhitespace(toksTrim(arg))
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
			// trc("<nil>")
			if len(c.sources) == 0 {
				return nil
			}

			src := c.sources[0]
			c.sources = c.sources[1:]
			c.push(c.group(src))
		case group:
			// trc("group len %v", len(x))
			if len(x) == 0 {
				c.pop()
				break
			}

			c.tos = x[1:]
			c.push(x[0])
		case controlLine:
			// trc("controlLine %v", toksDump([]Token(x)))
			c.pop()
			switch string(x[1].Src()) {
			case "define":
				c.define(x)
			case "undef":
				c.undef(x)
			default:
				panic(todo("%v: %q", x[0].Position(), x[1].Src()))
			}
		case textLine:
			// trc("textLine %v", toksDump([]Token(x)))
			c.pop()
			return x
		case eofLine:
			// trc("eofLine")
			// Keep it on stack, EOF is sticky
			return textLine([]Token{Token(x)})
		case *ifSection:
			// trc("ifSection")
			switch {
			case x.ifGroup != nil:
				switch {
				case c.ifGroup(x.ifGroup):
					c.pop()
					c.push(x.ifGroup.group)
				default:
					panic(todo(""))
				}
			default:
				panic(todo(""))
			}
		default:
			panic(todo("internal error: %T", x))
		}
	}
}

func (c *cpp) ifGroup(ig *ifGroup) bool {
	ln := ig.line[:len(ig.line)-1] // Remove new-line
	switch string(ln[1].Src()) {
	case "ifndef":
		if len(ln) < 3 { // '#' "ifndef" IDENTIFIER
			c.eh("%v: expected identifier", ln[1].Position())
			return false
		}

		_, ok := c.macros[string(ln[2].Src())]
		return !ok
	case "if":
		if len(ln) < 3 { // '#' "if" <expr>
			c.eh("%v: expected expression", ln[1].Position())
			return false
		}

		panic(todo(""))
		// return c.isNonZero(c.eval(tokens(ln[2:])))
	default:
		panic(todo("", toksDump(ln)))
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

	delete(c.macros, string(ln[2].Src()))
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
	switch {
	case c.macros[string(nm.Src())] != nil:
		panic(todo("", nm.Position(), &nm))
	}
	c.macros[string(nm.Src())] = c.newMacro(nm, fp, rl, minArgs, varArg, true)
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
	switch {
	case c.macros[string(nm.Src())] != nil:
		panic(todo("", nm.Position(), &nm))
	default:
		c.macros[string(nm.Src())] = c.newMacro(nm, nil, rl, -1, -1, false)
	}
}

func (c *cpp) newMacro(nm Token, params []Token, replList []preprocessingToken, minArgs, varArg int, isFnLike bool) *Macro {
	// trc("nm %q, params %v, replList %v, minArgs %v, varArg %v, isFnLike %v", nm.Src(), toksDump(params), toksDump(replList), minArgs, varArg, isFnLike)
	if string(nm.Src()) == "defined" {
		c.eh("%v: \"defined\" cannot be used as a macro name", nm.Position) // gcc says so.
		return nil
	}

	m, err := newMacro(nm, params, replList, minArgs, varArg, isFnLike)
	if err != nil {
		c.eh("", err)
	}
	return m
}

// parseReplacementList transforms s into preprocessing tokens that have separate
// tokens for white space.
func (c *cpp) parseReplacementList(s []Token) (r []preprocessingToken) {
	return normalizeHashWhitespace(tokens2preprocessingTokens(c.toksTrim(s)))
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

func (c *cpp) group(src Source) group {
	if g, ok := c.groups[src.Name]; ok {
		return g
	}

	p, err := newCppParser(src, c.eh)
	if err != nil {
		c.eh("", err)
		return nil
	}

	g := p.group(false)
	c.groups[src.Name] = g
	return g
}
