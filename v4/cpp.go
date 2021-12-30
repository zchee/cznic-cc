// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"fmt"
)

var (
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

	IsFnLike  bool
	IsVarArgs bool // #define foo(...), #define foo(bar, ...) etc.
}

func (m *Macro) ts() tokenSequence {
	r := m.ReplacementList
	return (*tokens)(&r)
}

func (m *Macro) fp() formalParams {
	if !m.IsFnLike {
		return nil
	}

	panic(todo(""))
}

type hideSet map[string]struct{}

func (h *hideSet) String() string {
	hs := *h
	var a []string
	for k := range hs {
		a = append(a, k)
	}
	return fmt.Sprintf("%q", a)
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

func (h hideSet) union(src hideSet) hideSet {
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

func (t *tokenizer) peek(index int) (tok Token) { panic(todo("internal error")) }

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
	eh        errHandler
	groups    map[string]group
	macros    map[string]*Macro
	sources   []Source
	stack     []interface{}
	tok       Token
	tokenizer *tokenizer
	tos       interface{}

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
		ln = ln[1:]
		c.defineFnMacro(nm, ln)
	default:
		// # define identifier replacement-list new-line
		//                     ^ln[0]
		c.defineObjectMacro(nm, ln[:len(ln)-1]) // strip new-line
	}
}

func (c *cpp) defineFnMacro(nm Token, line []Token) {
	fp := []Token{line[0]}
	switch line[0].Ch {
	case rune(DDD): // ...
		panic(todo("", toksDump(line)))
	case rune(IDENTIFIER):
		line = line[1:]
	out:
		for {
			switch line[0].Ch {
			case ')':
				line = line[1:]
				break out
			default:
				panic(todo("", toksDump(line)))
			}
		}
	case '(':
		panic(todo("", toksDump(line)))
	default:
		panic(todo("", toksDump(line)))
	}
	line = line[:len(line)-1] // strip new-line
	rl := c.replacementList(line)
	switch {
	case c.macros[string(nm.Src())] != nil:
		panic(todo("", nm.Position(), &nm))
	default:
		trc("", toksDump(fp))
		trc("", toksDump(rl))
		c.macros[string(nm.Src())] = &Macro{Name: nm, Params: fp, ReplacementList: rl, IsFnLike: true}
	}
}

func (c *cpp) defineObjectMacro(nm Token, replacementList []Token) {
	rl := c.replacementList(replacementList)
	switch {
	case c.macros[string(nm.Src())] != nil:
		panic(todo("", nm.Position(), &nm))
	default:
		c.macros[string(nm.Src())] = &Macro{Name: nm, ReplacementList: rl}
	}
}

// replacementList transforms s into preprocessing tokens that have separate
// tokens for white space.
func (c *cpp) replacementList(s []Token) (r []Token) {
	for len(s) != 0 && s[0].Ch == ' ' {
		s = s[1:]
	}
	for len(s) != 0 && s[len(s)-1].Ch == ' ' {
		s = s[:len(s)-1]
	}
	for i, v := range s {
		switch {
		case i == 0:
			if b := v.Sep(); len(b) != 0 {
				v.Set(nil, v.Src())
			}
		default:
			if b := v.Sep(); len(b) != 0 {
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
	return r
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

// [1], pg 1.
func (c *cpp) expand(TS tokenSequence) (r preprocessingTokens) {
	trc("TS %v", toksDump(TS))
	defer func() { trc("expand -> %v", toksDump(r)) }()
	var ptok preprocessingToken
	var ok bool
	ptok.Token, ptok.hs, ok = TS.read()
	if !ok {
		// if TS is {} then
		//	return {};
		return nil
	}

	T := ptok.Token
	HS := ptok.hs
	if T.Ch == eof {
		r = append(r, ptok)
		return r
	}

	src := string(T.Src())
	if HS.has(src) {
		// if TS is T^HS • TS’ and T is in HS then
		//	return T^HS • expand(TS’);
		panic(todo("", &T))
	}

	if m := c.macros[src]; m != nil {
		switch {
		default:
			// if TS is T^HS • TS’ and T is a "()-less macro" then
			//	return expand(subst(ts(T),{},{},HS∪{T},{}) • TS’);
			seq := c.subst(m.ts(), nil, nil, HS.add(src), nil)
			return c.expand(&seq)
		case m.IsFnLike:
			ts, args, ok := c.parseMacroArgs()
			_ = ts
			_ = args
			_ = ok
			trc("", toksDump(ts))
			trc("", toksDump(args))
			trc("", ok)
			panic(todo("", &T))
		}
	}

	// note TS must be T HS • TS’
	// return T HS • expand(TS’);
	return append(append(r, ptok), c.expand(TS)...)
}

func (c *cpp) parseMacroArgs() (consumed []Token, args [][]Token, ok bool) {
	panic(todo(""))
}

// [1], pg 2.
func (c *cpp) subst(IS tokenSequence, FP formalParams, AP interface{}, HS hideSet, OS preprocessingTokens) (r preprocessingTokens) {
	trc("IS %v, HS %v, OS %v", toksDump(IS), &HS, toksDump(OS))
	defer func() { trc("subst -> %v", toksDump(r)) }()
	var ptok preprocessingToken
	var ok bool
	ptok.Token, ptok.hs, ok = IS.read()
	if !ok {
		// if IS is {} then
		//	return hsadd(HS,OS);
		return c.hsAdd(HS, OS)
	}

	if ptok.Ch == '#' {
		// if IS is # • T • IS’ and T is FP[i] then
		//	return subst(IS’,FP,AP,HS,OS • stringize(select(i,AP )));
		panic(todo("", &ptok))
	}

	if ptok.Ch == rune(PPPASTE) {
		// if IS is ## • T • IS’ and T is FP[i] then
		panic(todo(""))
	}

	if IS.peek(1).Ch == rune(PPPASTE) {
		// if IS is T • ## HS’ • IS’ and T is FP[i] then
		panic(todo(""))
	}

	if FP != nil {
		if i := FP.is(string(ptok.Src())); i >= 0 {
			// if IS is T • IS’ and T is FP[i] then
			//	return subst(IS’,FP,AP,HS,OS • expand(select(i,AP )));
			panic(todo(""))
		}
	}

	// note IS must be T HS’ • IS’
	// return subst(IS’,FP,AP,HS,OS • T HS’ );
	return c.subst(IS, FP, AP, HS, append(OS, ptok))
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
		v.hs.union(HS)
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
