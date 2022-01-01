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
	_ tokenSequence = (*tokens)(nil)
)

type tokens []Token

func (p tokens) peek(index int) (tok Token) {
	if index < len(p) {
		return p[index]
	}

	tok.Ch = eof
	return tok
}

func (p *tokens) read() (tok Token, hs hideSet, ok bool) {
	s := *p
	if len(s) == 0 {
		return tok, nil, false
	}

	tok = s[0]
	if tok.Ch != eof {
		s = s[1:]
		*p = s
	}
	return tok, nil, true
}

// Macro represents a preprocessor #define.
type Macro struct {
	Name            Token
	Params          []Token
	ReplacementList []Token
	params          formalParams

	minArgs int

	IsFnLike  bool
	IsVarArgs bool // #define foo(...), #define foo(bar, ...) etc.
}

func newMacro(nm Token, params, replList []Token, minArgs int, isFnLike, isVarArgs bool) (*Macro, error) {
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
		IsVarArgs:       isVarArgs,
		Name:            nm,
		Params:          params,
		ReplacementList: replList,
		minArgs:         minArgs,
		params:          fp,
	}, nil
}

func (m *Macro) ts() tokenSequence {
	r := m.ReplacementList
	return (*tokens)(&r)
}

func (m *Macro) fp() formalParams {
	if !m.IsFnLike {
		return nil
	}

	return m.params
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

func (h hideSet) add(s string) hideSet {
	r := hideSet{}
	for k := range h {
		r[k] = struct{}{}
	}
	r[s] = struct{}{}
	return r
}

func (h hideSet) has(s string) bool {
	if h == nil {
		return false
	}

	for k := range h {
		if s == k {
			return true
		}
	}

	return false
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

type tokenizer struct {
	c    *cpp
	line textLine
	toks []preprocessingToken
}

func newTokenizer(c *cpp) *tokenizer {
	return &tokenizer{c: c}
}

func (t *tokenizer) peek(index int) (tok Token) {
	for {
		n := len(t.line)
		if index < n {
			return t.line[index]
		}

		if n != 0 && t.line[n-1].Ch == eof {
			return t.line[n-1]
		}

		t.line = append(t.line, t.c.nextLine()...)
	}
}

func (t *tokenizer) read() (tok Token, hs hideSet, ok bool) {
	if len(t.line) == 0 {
		t.line = t.c.nextLine()
	}

	tok = t.line[0]
	if tok.Ch != eof {
		t.line = t.line[1:]
	}
	return tok, nil, true // The hide set starts empty
}

func (t *tokenizer) token() (tok Token) {
	if len(t.toks) == 0 {
		t.toks = t.c.expand(t)
	}
	tok = t.toks[0].Token
	if tok.Ch != eof {
		t.toks = t.toks[1:]
	}
	return tok
}

type tokenSequence interface {
	peek(int) Token
	read() (tok Token, hs hideSet, ok bool)
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

func (c *cat) peek(index int) (tok Token) {
	if index < len(c.head) {
		return c.head.peek(index)
	}

	return c.tail.peek(index - len(c.head))
}

func (c *cat) read() (tok Token, hs hideSet, ok bool) {
	if len(c.head) != 0 {
		return c.head.read()
	}

	return c.tail.read()
}

type preprocessingTokens []preprocessingToken

func (p preprocessingTokens) peek(index int) (tok Token) {
	if index < len(p) {
		return p[index].Token
	}

	tok.Ch = eof
	return tok
}

func (p *preprocessingTokens) read() (tok Token, hs hideSet, ok bool) {
	s := *p
	if len(s) == 0 {
		return tok, nil, false
	}

	tok = s[0].Token
	hs = s[0].hs
	if tok.Ch != eof {
		s = s[1:]
		*p = s
	}
	return tok, hs, true
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

// consume returns c.tok and invalidates c.tok.Ch.
func (c *cpp) consume() (r Token) {
	r = c.tok
	c.tok.Ch = eof
	return r
}

// nextLine returns the next input textLine.
func (c *cpp) nextLine() textLine {
	for {
		switch x := c.tos.(type) {
		case nil:
			if len(c.sources) == 0 {
				return nil
			}

			src := c.sources[0]
			c.sources = c.sources[1:]
			c.push(c.group(src))
		case group:
			c.pop()
			if len(x) == 0 {
				break
			}

			c.tos = x[1:]
			c.push(x[0])
		case controlLine:
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
			c.pop()
			return x
		case eofLine:
			// Don't pop it, EOF is sticky
			return textLine([]Token{Token(x)})
		default:
			panic(todo("internal error: %T", x))
		}
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
		ln = ln[1:] // '('
		// # define identifier ( args_opt ) replacement-list new-line
		//                       ^ln[0]
		c.defineFnMacro(nm, ln)
	default:
		// # define identifier replacement-list new-line
		//                     ^ln[0]
		c.defineObjectMacro(nm, ln[:len(ln)-1]) // strip new-line
	}
}

func (c *cpp) defineFnMacro(nm Token, line []Token) {
	minArgs := 0
	var fp []Token
	isVarArgs := false
	switch line[0].Ch {
	case rune(DDD): // ...
		panic(todo("", toksDump(line)))
	case rune(IDENTIFIER):
		fp = append(fp, line[0])
		minArgs++
		line = line[1:] // remove the FP
	out:
		for {
			switch line[0].Ch {
			case ',':
				line = line[1:] // ','
			case rune(IDENTIFIER):
				fp = append(fp, line[0])
				minArgs++
				line = line[1:] // remove the FP
			case ')':
				line = line[1:] // remove the final FP list paren
				break out
			case rune(DDD):
				ddd := line[0]
				line = line[1:] // ...
				isVarArgs = true
				for len(line) != 0 && line[0].Ch == ' ' {
					line = line[1:]
				}
				if len(line) == 0 || line[0].Ch != ')' {
					c.eh("%v: ... must be followd by ')'", ddd.Position())
					return
				}

				line = line[1:] // ')'
				break out
			default:
				panic(todo("", toksDump(line)))
			}
		}
	case ')':
		panic(todo("", toksDump(line)))
	default:
		panic(todo("", toksDump(line)))
	}
	line = line[:len(line)-1] // strip new-line
	rl := c.parseReplacementList(line)
	switch {
	case c.macros[string(nm.Src())] != nil:
		panic(todo("", nm.Position(), &nm))
	}
	c.macros[string(nm.Src())] = c.newMacro(nm, fp, rl, minArgs, true, isVarArgs)
}

func (c *cpp) newMacro(nm Token, params, replList []Token, minArgs int, isFnLike, isVarArgs bool) *Macro {
	m, err := newMacro(nm, params, replList, minArgs, isFnLike, isVarArgs)
	if err != nil {
		c.eh("", err)
	}
	return m
}

func (c *cpp) defineObjectMacro(nm Token, replacementList []Token) {
	rl := c.parseReplacementList(replacementList)
	switch {
	case c.macros[string(nm.Src())] != nil:
		panic(todo("", nm.Position(), &nm))
	default:
		c.macros[string(nm.Src())] = c.newMacro(nm, nil, rl, -1, false, false)
	}
}

// parseReplacementList transforms s into preprocessing tokens that have separate
// tokens for white space.
func (c *cpp) parseReplacementList(s []Token) (r []Token) {
	s = c.toksTrim(s)
	for i, v := range s {
		if b := v.Sep(); len(b) != 0 {
			switch {
			case i == 0:
				v.Set(nil, v.Src())
			default:
				w := v
				w.Ch = ' '
				w.len = uint32(len(b))
				w.src = w.sep
				w.sep = w.src
				r = append(r, w)
			}
		}
		r = append(r, v)
	}
	w := 0
	for i, v := range r {
		switch v.Ch {
		case ' ':
			if i != 0 {
				switch r[i-1].Ch {
				case ' ', '#', rune(PPPASTE):
					continue
				}
			}
			if i+1 < len(r) {
				switch r[i+1].Ch {
				case rune(PPPASTE):
					continue
				}
			}
		}

		r[w] = v
		w++
	}
	return r[:w]
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

func (c *cpp) push(v interface{}) {
	if c.tos != nil {
		c.stack = append(c.stack, c.tos)
	}
	c.tos = v
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
func (c *cpp) expand(TS tokenSequence) (r preprocessingTokens) {
	// trc("  %s%v (%v)", c.indent(), toksDump(TS), origin(2))
	// defer func() { trc("->%s%v", c.undent(), toksDump(r)) }()
	var ptok preprocessingToken
	var ok bool
	ptok.Token, ptok.hs, ok = TS.read()
	if !ok {
		// if TS is {} then
		//	return {};
		return nil
	}

	// trc("%s^%s, ok %v", &ptok.Token, &ptok.hs, ok)
	T := ptok.Token
	HS := ptok.hs
	if T.Ch == eof {
		return preprocessingTokens{ptok}
	}

	src := string(T.Src())
	if HS.has(src) {
		// if TS is T^HS • TS’ and T is in HS then
		//	return T^HS • expand(TS’);
		return append(preprocessingTokens{ptok}, c.expand(TS)...)
	}

	if m := c.macros[src]; m != nil {
		va := -1
	out:
		switch {
		default:
			// if TS is T^HS • TS’ and T is a "()-less macro" then
			//	return expand(subst(ts(T),{},{},HS∪{T},{}) • TS’);
			// trc("  %s<%s is a ()-less macro, expanding to %s>", c.indent(), T.Src(), toksDump(m.ts()))
			// defer func() { trc("  %s<%s expanded>", c.undent(), T.Src()) }()
			seq := c.subst(m.ts(), nil, nil, va, HS.add(src), nil)
			return c.expand(newCat(seq, TS))
		case m.IsFnLike:
			if TS.peek(0).Ch == '(' {
				// if TS is T^HS • ( • TS’ and T is a "()’d macro" then
				//	check TS’ is actuals • )^HS’ • TS’’ and actuals are "correct for T"
				//	return expand(subst(ts(T),fp(T),actuals,(HS∩HS’)∪{T},{}) • TS’’);
				// trc("  %s<%s is a ()'d macro, expanding to %s>", c.indent(), T.Src(), toksDump(m.ts()))
				// defer func() { trc("  %s<%s expanded>", c.undent(), T.Src()) }()
				args, rparen, ok := c.parseMacroArgs(TS)
				if !ok {
					return nil
				}

				switch {
				case len(args) < m.minArgs:
					c.eh("%v: not enough macro arguments", rparen.Position())
					break out
				case len(args) > m.minArgs && !m.IsVarArgs:
					c.eh("%v: too many macro arguments", rparen.Position())
					break out
				}
				if m.IsVarArgs {
					va = m.minArgs
				}

				return c.expand(newCat(c.subst(m.ts(), m.fp(), args, va, HS.cap(rparen.hs).add(src), nil), TS))
			}
		}
	}

	// note TS must be T HS • TS’
	// return T HS • expand(TS’);
	return append(preprocessingTokens{ptok}, c.expand(TS)...)
}

// [1], pg 2.
func (c *cpp) subst(IS tokenSequence, FP formalParams, AP []preprocessingTokens, va int, HS hideSet, OS preprocessingTokens) (r preprocessingTokens) {
	// trc("  %s%v, HS %v, FP %v, AP %v, OS %v (%v)", c.indent(), toksDump(IS), &HS, FP, toksDump(AP), toksDump(OS), origin(2))
	// defer func() { trc("->%s%v", c.undent(), toksDump(r)) }()
	var ptok preprocessingToken
	var ok bool
	ptok.Token, ptok.hs, ok = IS.read()
	if !ok {
		// if IS is {} then
		//	return hsadd(HS,OS);
		return c.hsAdd(HS, OS)
	}

	if ptok.Ch == '#' {
		switch tok, skip := c.peekNonBlank(IS); {
		case tok.Ch == rune(IDENTIFIER):
			if i := FP.is(string(tok.Src())); i >= 0 {
				// if IS is # • T • IS’ and T is FP[i] then
				//	return subst(IS’,FP,AP,HS,OS • stringize(select(i,AP )));
				c.skip(IS, skip+1)
				return c.subst(IS, FP, AP, va, HS, append(OS, c.stringize(AP[i])))
			}
		}
	}

	if ptok.Ch == rune(PPPASTE) {
		switch tok, skip := c.peekNonBlank(IS); {
		case tok.Ch == rune(IDENTIFIER):
			if i := FP.is(string(tok.Src())); i >= 0 {
				// if IS is ## • T • IS’ and T is FP[i] then
				if len(AP[i]) == 0 {
					//	if select(i,AP ) is {} then /* only if actuals can be empty */
					//		return subst(IS’,FP,AP,HS,OS );
					panic(todo("", toksDump(IS), skip, i))
				} else {
					//	else
					//		return subst(IS’,FP,AP,HS,glue(OS,select(i,AP )));
					c.skip(IS, skip+1)
					return c.subst(IS, FP, AP, va, HS, c.glue(OS, AP[i]))
				}
			}
		}
		// else if IS is ## • T HS’ • IS’ then
		//	return subst(IS’,FP,AP,HS,glue(OS,T^HS’));
		if IS.peek(0).Ch != eof {
			var t2 preprocessingToken
			t2.Token, t2.hs, _ = IS.read()
			return c.subst(IS, FP, AP, va, HS, c.glue(OS, preprocessingTokens{t2}))
		}
	}

	if ppasteTok, skip := c.peekNonBlank(IS); ppasteTok.Ch == rune(PPPASTE) {
		if ptok.Ch == rune(IDENTIFIER) {
			if i := FP.is(string(ptok.Src())); i >= 0 {
				// if IS is T • ##^HS’ • IS’ and T is FP[i] then
				//	if select(i,AP ) is {} then /* only if actuals can be empty */
				if len(AP[i]) == 0 {
					//	{
					//		if IS’ is T’ • IS’’ and T’ is FP[j] then
					//			return subst(IS’’,FP,AP,HS,OS • select(j,AP));
					//		else
					//			return subst(IS’,FP,AP,HS,OS);
					//	}
					trc("", &ppasteTok, skip)
					panic(todo("", toksDump(IS)))
				} else {
					//	else
					//		return subst(##^HS’ • IS’,FP,AP,HS,OS • select(i,AP));
					c.skip(IS, skip)
					return c.subst(IS, FP, AP, va, HS, append(OS, AP[i]...))
				}
			}
		}
	}

	if len(FP) != 0 {
		if i := FP.is(string(ptok.Src())); i >= 0 {
			// if IS is T • IS’ and T is FP[i] then
			//	return subst(IS’,FP,AP,HS,OS • expand(select(i,AP )));
			return c.subst(IS, FP, AP, va, HS, append(OS, c.expand(&AP[i])...))
		}
	}

	if ptok.Ch == rune(IDENTIFIER) && string(ptok.Src()) == "__VA_ARGS__" && va < len(AP) {
		var comma preprocessingTokens
		if len(AP)-va > 1 {
			t := ptok
			t.Set(nil, []byte{','})
			t.Ch = ','
			t2 := t
			t2.Set(nil, []byte{' '})
			t2.Ch = ' '
			comma = preprocessingTokens{t, t2}
		}
		var a preprocessingTokens
		for i, v := range AP[va:] {
			if i != 0 {
				a = append(a, comma...)
			}
			a = append(a, v...)
		}
		return c.subst(IS, FP, AP, va, HS, append(OS, c.expand(&a)...))
	}

	// note IS must be T HS’ • IS’
	// return subst(IS’,FP,AP,HS,OS • T HS’ );
	return c.subst(IS, FP, AP, va, HS, append(OS, ptok))
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

func (c *cpp) skip(s tokenSequence, n int) {
	for ; n != 0; n-- {
		s.read()
	}
}

func (c *cpp) peekNonBlank(s tokenSequence) (tok Token, skip int) {
	for {
		tok = s.peek(skip)
		if tok.Ch != ' ' {
			return tok, skip
		}

		skip++
	}
}

func (c *cpp) parseMacroArgs(TS tokenSequence) (args []preprocessingTokens, rparen preprocessingToken, ok bool) {
	// trc("  %sin  %v (%v)", c.indent(), toksDump(TS), origin(2))
	// defer func() { trc("->%sout %v", c.undent(), toksDump(args)) }()
	var tok preprocessingToken
	if tok.Token, tok.hs, ok = TS.read(); !ok || tok.Ch != '(' {
		c.eh("%v: expected macro argument list starting with '('", tok.Position())
		return nil, preprocessingToken{}, false
	}

	var arg preprocessingTokens
	level := 0
out:
	for {
		switch t := TS.peek(0); t.Ch {
		case ',':
			if level != 0 {
				tok.Token, tok.hs, _ = TS.read()
				arg = append(arg, tok)
				break
			}

			args = append(args, arg)
			TS.read() // discard the ','
			arg = nil
		case '(':
			tok.Token, tok.hs, _ = TS.read()
			arg = append(arg, tok)
			level++
		case ')':
			if level == 0 {
				args = append(args, arg)
				break out
			}

			tok.Token, tok.hs, _ = TS.read()
			arg = append(arg, tok)
			level--
		case eof:
			panic(todo("", &t))
		default:
			tok.Token, tok.hs, _ = TS.read()
			arg = append(arg, tok)
		}
	}

	if tok.Token, tok.hs, ok = TS.read(); tok.Ch != ')' {
		c.eh("%v: expected macro argument list terminating ')'", tok.Position())
		return nil, preprocessingToken{}, false
	}

	for i, arg0 := range args {
		arg0 = c.cppToksTrim(arg0)
		var arg preprocessingTokens
		for j, v := range arg0 {
			if b := v.Sep(); len(b) != 0 {
				switch {
				case j == 0:
					v.Set(nil, v.Src())
				default:
					w := v
					w.Ch = ' '
					w.len = uint32(len(b))
					w.src = w.sep
					w.sep = w.src
					arg = append(arg, w)
				}
			}
			arg = append(arg, v)
		}
		w := 0
		for i, v := range arg {
			switch v.Ch {
			case ' ':
				if i != 0 {
					switch arg[i-1].Ch {
					case ' ', '#', rune(PPPASTE):
						continue
					}
				}
				if i+1 < len(arg) {
					switch arg[i+1].Ch {
					case rune(PPPASTE):
						continue
					}
				}
			}

			arg[w] = v
			w++
		}
		args[i] = arg[:w]
	}
	return args, tok, true
}

func (c *cpp) cppToksTrim(s preprocessingTokens) preprocessingTokens {
	for len(s) != 0 && s[0].Ch == ' ' {
		s = s[1:]
	}
	for len(s) != 0 && s[len(s)-1].Ch == ' ' {
		s = s[:len(s)-1]
	}
	return s
}

func (c *cpp) hsAdd(HS hideSet, TS preprocessingTokens) (r preprocessingTokens) {
	if len(TS) == 0 {
		// if TS is {} then
		//	return {};
		return r
	}

	// note TS must be T^HS’ • TS’
	// return T^(HS∪HS’) • hsadd(HS,TS’ );
	for _, v := range TS {
		v.hs = v.hs.cup(HS)
		r = append(r, v)
	}
	return r
}

// c returns the current token rune the preprocessor is positioned on.
func (c *cpp) c() rune {
	if c.tok.Ch != eof || c.closed {
		return c.tok.Ch
	}

	if c.tok = c.tokenizer.token(); c.tok.Ch == eof {
		c.closed = true
	}
	return c.tok.Ch
}
