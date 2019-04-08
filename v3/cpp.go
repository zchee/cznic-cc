// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"modernc.org/token"
)

const (
	maxIncludeLevel = 200 // gcc, std is at least 15.
)

var (
	_ tokenReader = (*cpp)(nil)
	_ tokenWriter = (*cpp)(nil)

	idComma          = dict.sid(",")
	idCxLimitedRange = dict.sid("CX_LIMITED_RANGE")
	idDefault        = dict.sid("DEFAULT")
	idDefined        = dict.sid("defined")
	idFILE           = dict.sid("__FILE__")
	idFPContract     = dict.sid("FP_CONTRACT")
	idFenvAccess     = dict.sid("FENV_ACCESS")
	idIntMaxWidth    = dict.sid("__INTMAX_WIDTH__")
	idLINE           = dict.sid("__LINE__")
	idNL             = dict.sid("\n")
	idOff            = dict.sid("OFF")
	idOn             = dict.sid("ON")
	idOne            = dict.sid("1")
	idPragmaSTDC     = dict.sid("__pragma_stdc")
	idSTDC           = dict.sid("STDC")
	idVaArgs         = dict.sid("__VA_ARGS__")
	idZero           = dict.sid("0")

	protectedMacros = hideSet{ // [0], 6.10.8, 4
		dict.sid("__DATE__"):                 {},
		dict.sid("__STDC_HOSTED__"):          {},
		dict.sid("__STDC_IEC_559_COMPLEX__"): {},
		dict.sid("__STDC_IEC_559__"):         {},
		dict.sid("__STDC_ISO_10646__"):       {},
		dict.sid("__STDC_MB_MIGHT_NEQ_WC__"): {},
		dict.sid("__STDC_VERSION__"):         {},
		dict.sid("__STDC__"):                 {},
		dict.sid("__TIME__"):                 {},
		idFILE:                               {},
		idLINE:                               {},
	}
)

type tokenReader interface {
	read() (cppToken, bool)
	unget(cppToken)
	ungets([]cppToken)
}

type tokenWriter interface {
	write(cppToken)
	writes([]cppToken)
}

// token4 is produced by translation phase 4.
type token4 struct {
	fileID int32
	token3
}

func (t *token4) Position() (r token.Position) {
	if t.pos != 0 && t.fileID != 0 {
		if f := cache.file(t.fileID); f != nil {
			r = f.PositionFor(token.Pos(t.pos), true)
		}
	}
	return r
}

type hideSet map[StringID]struct{}

type cppToken struct {
	token4
	hs hideSet
}

func (t *cppToken) has(nm StringID) bool { _, ok := t.hs[nm]; return ok }

type cppWriter struct {
	toks []cppToken
}

func (w *cppWriter) write(tok cppToken)     { w.toks = append(w.toks, tok) }
func (w *cppWriter) writes(toks []cppToken) { w.toks = append(w.toks, toks...) }

type ungetBuf []cppToken

func (u *ungetBuf) unget(t cppToken) { *u = append(*u, t) }

func (u *ungetBuf) read() (t cppToken) {
	s := *u
	n := len(s) - 1
	t = s[n]
	*u = s[:n]
	return t
}
func (u *ungetBuf) ungets(toks []cppToken) {
	s := *u
	for i := len(toks) - 1; i >= 0; i-- {
		s = append(s, toks[i])
	}
	*u = s
}

func cppToksStr(toks []cppToken, sep string) string {
	var b strings.Builder
	for i, v := range toks {
		if i != 0 {
			b.WriteString(sep)
		}
		b.WriteString(v.String())
	}
	return b.String()
}

type cppReader struct {
	buf []cppToken
	ungetBuf
}

func (r *cppReader) read() (tok cppToken, ok bool) {
	if len(r.ungetBuf) != 0 {
		return r.ungetBuf.read(), true
	}

	if len(r.buf) == 0 {
		return tok, false
	}

	tok = r.buf[0]
	r.buf = r.buf[1:]
	return tok, true
}

type cppScanner []cppToken

func (s *cppScanner) peek() (r cppToken) {
	r.char = -1
	if len(*s) == 0 {
		return r
	}

	return (*s)[0]
}

func (s *cppScanner) next() (r cppToken) {
	r.char = -1
	if len(*s) == 0 {
		return r
	}

	*s = (*s)[1:]
	return s.peek()
}

func (s *cppScanner) Pos() token.Pos {
	if len(*s) == 0 {
		return 0
	}

	return (*s)[0].Pos()
}

type macro struct {
	fp   []StringID
	repl []token3

	name token4

	isFnLike      bool
	namedVariadic bool // foo..., note no comma before ellipsis.
	variadic      bool
}

func (m *macro) param(ap [][]cppToken, nm StringID, out *[]cppToken) bool {
	*out = nil
	if nm == idVaArgs || m.namedVariadic && nm == m.fp[len(m.fp)-1] {
		if !m.variadic {
			return false
		}

		i := len(m.fp)
		if m.namedVariadic {
			i--
		}
		if i < len(ap) {
			var o []cppToken
			sel := ap[i:]
			for i, v := range sel {
				switch {
				case i == 0:
					switch {
					case len(v) != 0 && v[0].char == ' ':
						o = append(o, v[1:]...)
					default:
						o = append(o, v...)
					}
				default:
					o = append(o, v...)
				}
				if i != len(sel)-1 {
					var t cppToken
					t.char = ','
					t.value = idComma
					o = append(o, t)
				}
			}
			switch {
			case len(o) != 0 && o[0].char == ' ':
				*out = o[1:]
			default:
				*out = o
			}
		}
		return true
	}

	for i, v := range m.fp {
		if v == nm {
			switch {
			case i >= len(ap):
				// nop
			case len(ap[i]) != 0 && ap[i][0].char == ' ':
				*out = ap[i][1:]
			default:
				*out = ap[i]
			}
			return true
		}
	}
	return false
}

// --------------------------------------------------------------- Preprocessor

type cpp struct {
	ctx          *context
	file         *token.File
	fileMacro    macro
	in           chan []token3
	inBuf        []token3
	includeLevel int
	lineMacro    macro
	macros       map[StringID]*macro
	out          chan *[]token4
	outBuf       *[]token4
	rq           chan struct{}
	ungetBuf

	fileID int32

	intmaxChecked bool
	nonFirstRead  bool
	seenEOF       bool
}

func newCPP(ctx *context) *cpp {
	b := token4Pool.Get().(*[]token4)
	*b = (*b)[:0]
	r := &cpp{
		ctx:    ctx,
		macros: map[StringID]*macro{},
		outBuf: b,
	}
	r.fileMacro = macro{repl: []token3{{char: STRINGLITERAL}}}
	r.lineMacro = macro{repl: []token3{{char: PPNUMBER}}}
	r.macros = map[StringID]*macro{
		idFILE: &r.fileMacro,
		idLINE: &r.lineMacro,
	}
	return r
}

func (c *cpp) cppToks(toks []token3) (r []cppToken) {
	r = make([]cppToken, len(toks))
	for i, v := range toks {
		r[i].token4.token3 = v
	}
	return r
}

func (c *cpp) err(n node, msg string, args ...interface{}) (stop bool) {
	var position token.Position
	switch x := n.(type) {
	case nil:
	case token4:
		position = x.Position()
	default:
		if p := n.Pos(); p.IsValid() {
			position = c.file.PositionFor(p, true)
		}
	}
	return c.ctx.err(position, msg, args...)
}

func (c *cpp) read() (cppToken, bool) {
	if len(c.ungetBuf) != 0 {
		return c.ungetBuf.read(), true
	}

	if len(c.inBuf) == 0 {
		if c.seenEOF {
			return cppToken{}, false
		}

		if c.nonFirstRead {
			c.rq <- struct{}{}
		}
		c.nonFirstRead = true

		var ok bool
		if c.inBuf, ok = <-c.in; !ok {
			c.seenEOF = true
			return cppToken{}, false
		}
	}

	tok := c.inBuf[0]
	c.inBuf = c.inBuf[1:]
	return cppToken{token4{c.fileID, tok}, nil}, true
}

func (c *cpp) write(tok cppToken) {
	*c.outBuf = append(*c.outBuf, tok.token4)
	if tok.char == '\n' {
		c.out <- c.outBuf
		b := token4Pool.Get().(*[]token4)
		*b = (*b)[:0]
		c.outBuf = b
	}
}

func (c *cpp) writes(toks []cppToken) {
	for _, v := range toks {
		c.write(v)
	}
}

// [1]pg 1.
//
// expand(TS) /* recur, substitute, pushback, rescan */
// {
// 	if TS is {} then
//		// ---------------------------------------------------------- A
// 		return {};
//
// 	else if TS is T^HS • TS’ and T is in HS then
//		//----------------------------------------------------------- B
// 		return T^HS • expand(TS’);
//
// 	else if TS is T^HS • TS’ and T is a "()-less macro" then
//		// ---------------------------------------------------------- C
// 		return expand(subst(ts(T), {}, {}, HS \cup {T}, {}) • TS’ );
//
// 	else if TS is T^HS •(•TS’ and T is a "()’d macro" then
//		// ---------------------------------------------------------- D
// 		check TS’ is actuals • )^HS’ • TS’’ and actuals are "correct for T"
// 		return expand(subst(ts(T), fp(T), actuals,(HS \cap HS’) \cup {T }, {}) • TS’’);
//
//	// ------------------------------------------------------------------ E
// 	note TS must be T^HS • TS’
// 	return T^HS • expand(TS’);
// }
func (c *cpp) expand(ts tokenReader, w tokenWriter, expandDefined bool) {
start:
	tok, ok := ts.read()
	tok.fileID = c.fileID
	// First, if TS is the empty set, the result is the empty set.
	if !ok {
		// ---------------------------------------------------------- A
		// return {};
		return
	}

	if tok.char == IDENTIFIER {
		nm := tok.value
		if nm == idDefined && expandDefined {
			c.parseDefined(tok, ts, w)
			goto start
		}

		// Otherwise, if the token sequence begins with a token whose
		// hide set contains that token, then the result is the token
		// sequence beginning with that token (including its hide set)
		// followed by the result of expand on the rest of the token
		// sequence.
		if tok.has(nm) {
			// -------------------------------------------------- B
			// return T^HS • expand(TS’);
			w.write(tok)
			goto start
		}

		m := c.macros[nm]
		if m != nil && !m.isFnLike {
			// Otherwise, if the token sequence begins with an
			// object-like macro, the result is the expansion of
			// the rest of the token sequence beginning with the
			// sequence returned by subst invoked with the
			// replacement token sequence for the macro, two empty
			// sets, the union of the macro’s hide set and the
			// macro itself, and an empty set.
			if nm == idLINE {
				c.lineMacro.repl[0].value = dict.sid(fmt.Sprint(tok.Position().Line))
			}
			// -------------------------------------------------- C
			// return expand(subst(ts(T), {}, {}, HS \cup {T}, {}) • TS’ );
			hs := hideSet{nm: {}}
			for k, v := range tok.hs {
				hs[k] = v
			}
			toks := c.subst(m, c.cppToks(m.repl), nil, nil, hs, nil, expandDefined)
			for i := range toks {
				toks[i].pos = tok.pos
			}
			ts.ungets(toks)
			goto start
		}

		if m != nil && m.isFnLike {
			// -------------------------------------------------- D
			// check TS’ is actuals • )^HS’ • TS’’ and actuals are "correct for T"
			// return expand(subst(ts(T), fp(T), actuals,(HS \cap HS’) \cup {T }, {}) • TS’’);
			hs := tok.hs
		again:
			t2, ok := ts.read()
			if !ok {
				w.write(tok)
				goto start
			}

			switch t2.char {
			case '\n', ' ':
				goto again
			case '(':
				// ok
			default:
				w.write(tok)
				ts.unget(t2)
				goto start
			}

			ap, hs2 := c.actuals(m, ts)
			switch {
			case len(hs2) == 0:
				hs2 = hideSet{nm: {}}
			default:
				nhs := hideSet{}
				for k := range hs {
					if _, ok := hs2[k]; ok {
						nhs[k] = struct{}{}
					}
				}
				nhs[nm] = struct{}{}
				hs2 = nhs
			}
			toks := c.subst(m, c.cppToks(m.repl), m.fp, ap, hs2, nil, expandDefined)
			for i := range toks {
				toks[i].pos = tok.pos
			}
			ts.ungets(toks)
			goto start
		}
	}

	// ------------------------------------------------------------------ E
	// note TS must be T^HS • TS’
	// return T^HS • expand(TS’);
	w.write(tok)
	goto start
}

func (c *cpp) actuals(m *macro, r tokenReader) (out [][]cppToken, hs hideSet) {
	var lvl, n int
	for {
		t, ok := r.read()
		if !ok {
			c.err(t, "unexpected EOF")
			return nil, nil
		}

		switch t.char {
		case ',':
			if lvl == 0 {
				n++
				continue
			}
		case ')':
			if lvl == 0 {
				for len(out) < len(m.fp) {
					out = append(out, nil)
				}
				return out, t.hs
			}
			lvl--
		case '(':
			lvl++
		}
		for len(out) <= n {
			out = append(out, []cppToken{})
		}
		out[n] = append(out[n], t)
	}
}

// [1]pg 2.
//
// subst(IS, FP, AP, HS, OS) /* substitute args, handle stringize and paste */
// {
// 	if IS is {} then
//		// ---------------------------------------------------------- A
// 		return hsadd(HS, OS);
//
// 	else if IS is # • T • IS’ and T is FP[i] then
//		// ---------------------------------------------------------- B
// 		return subst(IS’, FP, AP, HS, OS • stringize(select(i, AP)));
//
// 	else if IS is ## • T • IS’ and T is FP[i] then
//	{
//		// ---------------------------------------------------------- C
// 		if select(i, AP) is {} then /* only if actuals can be empty */
//			// -------------------------------------------------- D
// 			return subst(IS’, FP, AP, HS, OS);
// 		else
//			// -------------------------------------------------- E
// 			return subst(IS’, FP, AP, HS, glue(OS, select(i, AP)));
// 	}
//
// 	else if IS is ## • T^HS’ • IS’ then
//		// ---------------------------------------------------------- F
// 		return subst(IS’, FP, AP, HS, glue(OS, T^HS’));
//
// 	else if IS is T • ##^HS’ • IS’ and T is FP[i] then
//	{
//		// ---------------------------------------------------------- G
// 		if select(i, AP) is {} then /* only if actuals can be empty */
//		{
//			// -------------------------------------------------- H
// 			if IS’ is T’ • IS’’ and T’ is FP[j] then
//				// ------------------------------------------ I
// 				return subst(IS’’, FP, AP, HS, OS • select(j, AP));
// 			else
//				// ------------------------------------------ J
// 				return subst(IS’, FP, AP, HS, OS);
// 		}
//		else
//			// -------------------------------------------------- K
// 			return subst(##^HS’ • IS’, FP, AP, HS, OS • select(i, AP));
//
//	}
//
// 	else if IS is T • IS’ and T is FP[i] then
//		// ---------------------------------------------------------- L
// 		return subst(IS’, FP, AP, HS, OS • expand(select(i, AP)));
//
//	// ------------------------------------------------------------------ M
// 	note IS must be T^HS’ • IS’
// 	return subst(IS’, FP, AP, HS, OS • T^HS’);
// }
//
// A quick overview of subst is that it walks through the input sequence, IS,
// building up an output sequence, OS, by handling each token from left to
// right. (The order that this operation takes is left to the implementation
// also, walking from left to right is more natural since the rest of the
// algorithm is constrained to this ordering.) Stringizing is easy, pasting
// requires trickier handling because the operation has a bunch of
// combinations. After the entire input sequence is finished, the updated hide
// set is applied to the output sequence, and that is the result of subst.
func (c *cpp) subst(m *macro, is []cppToken, fp []StringID, ap [][]cppToken, hs hideSet, os []cppToken, expandDefined bool) (r []cppToken) {
	// var a []string
	// for _, v := range ap {
	// 	a = append(a, fmt.Sprintf("%q", cppToksStr(v, "|")))
	// }
	// dbg("==== subst: is %q, fp %v ap %v", cppToksStr(is, "|"), fp, a)
start:
	// dbg("start: %q", cppToksStr(is, "|"))
	if len(is) == 0 {
		// ---------------------------------------------------------- A
		// return hsadd(HS, OS);
		// dbg("subst returns %q", cppToksStr(os, "|"))
		return c.hsAdd(hs, os)
	}

	tok := is[0]
	var arg []cppToken
	if tok.char == '#' {
		var nm cppToken
		n := 0
		switch {
		case len(is) > 1 && is[1].char == IDENTIFIER:
			nm = is[1]
			n = 2
		case len(is) > 2 && is[1].char == ' ' && is[2].char == IDENTIFIER:
			nm = is[2]
			n = 3
		}
		if n != 0 && m.param(ap, nm.value, &arg) {
			// -------------------------------------------------- B
			// return subst(IS’, FP, AP, HS, OS • stringize(select(i, AP)));
			os = append(os, c.stringize(arg))
			is = is[n:]
			goto start
		}
	}

	if tok.char == PPPASTE {
		var nm cppToken
		n := 0
		switch {
		case len(is) > 1 && is[1].char == IDENTIFIER:
			nm = is[1]
			n = 2
		case len(is) > 2 && is[1].char == ' ' && is[2].char == IDENTIFIER:
			nm = is[2]
			n = 3
		}
		if n != 0 && m.param(ap, nm.value, &arg) {
			// -------------------------------------------------- C
			if len(arg) == 0 {
				// TODO "only if actuals can be empty"
				// ------------------------------------------ D
				// return subst(IS’, FP, AP, HS, OS);
				is = is[n:]
				goto start
			}

			// -------------------------------------------------- E
			// return subst(IS’, FP, AP, HS, glue(OS, select(i, AP)));
			os = c.glue(os, arg)
			is = is[n:]
			goto start
		}

		n = 0
		switch {
		case len(is) > 1 && is[1].char != ' ':
			n = 1
		case len(is) > 2 && is[1].char == ' ':
			n = 2
		}

		if n != 0 {
			// -------------------------------------------------- F
			// return subst(IS’, FP, AP, HS, glue(OS, T^HS’));
			os = c.glue(os, is[n:n+1])
			is = is[n+1:]
			goto start
		}
	}

	if tok.char == IDENTIFIER &&
		((len(is) > 1 && is[1].char == PPPASTE) || (len(is) > 2 && is[1].char == ' ' && is[2].char == PPPASTE)) &&
		m.param(ap, tok.value, &arg) {
		// ---------------------------------------------------------- G
		if len(arg) == 0 {
			// TODO "only if actuals can be empty"
			// -------------------------------------------------- H
			is2 := is[1:]                            // skip T
			if len(is2) != 0 && is2[0].char == ' ' { // skip ' '
				is2 = is2[1:]
			}
			is2 = is2[1:]                            // skip ##
			if len(is2) != 0 && is2[0].char == ' ' { // skip ' '
				is2 = is2[1:]
			}
			if len(is) > 0 && is[0].char == IDENTIFIER && m.param(ap, is2[0].value, &arg) {
				// -------------------------------------------------- I
				// return subst(IS’’, FP, AP, HS, OS • select(j, AP));
				os = append(os, arg...)
				is = is2[1:]
				goto start
			} else {
				// -------------------------------------------------- J
				// return subst(IS’, FP, AP, HS, OS);
				is = is2
				goto start
			}
		}

		// ---------------------------------------------------------- K
		// return subst(##^HS’ • IS’, FP, AP, HS, OS • select(i, AP));
		os = append(os, arg...)
		is = is[1:]
		goto start
	}

	if tok.char == IDENTIFIER && m.param(ap, tok.value, &arg) {
		// ------------------------------------------ L
		// return subst(IS’, FP, AP, HS, OS • expand(select(i, AP)));
		sel := cppReader{buf: arg}
		var w cppWriter
		c.expand(&sel, &w, expandDefined)
		os = append(os, w.toks...)
		is = is[1:]
		goto start
	}

	// ------------------------------------------------------------------ M
	// note IS must be T^HS’ • IS’
	// return subst(IS’, FP, AP, HS, OS • T^HS’);
	if len(is) < 2 || tok.char != ' ' || is[1].char != PPPASTE {
		os = append(os, tok)
	}
	is = is[1:]
	goto start
}

// paste last of left side with first of right side
//
// [1] pg. 3
func (c *cpp) glue(ls, rs []cppToken) (out []cppToken) {
	if len(rs) == 0 {
		return ls
	}

	if len(ls) == 0 {
		return rs
	}

	l := ls[len(ls)-1]
	if l.char == ',' { // testsuite/gcc.target/avr/torture/isr-01-simple.c
		return append(ls, rs...)
	}

	ls = ls[:len(ls)-1]
	r := rs[0]
	rs = rs[1:]

	l.value = dict.sid(l.String() + r.String())
	return append(append(ls, l), rs...)
}

// Given a token sequence, stringize returns a single string literal token
// containing the concatenated spellings of the tokens.
//
// [1] pg. 3
func (c *cpp) stringize(s []cppToken) (r cppToken) {
	var a []string
	for i, v := range s {
		if v.char == ' ' && i < len(s)-2 && s[i+1].char == '\n' || v.char == '\n' {
			continue
		}

		s := v.String()
		switch v.char {
		case CHARCONST:
			s = strings.Replace(s, `\`, `\\`, -1)
		case STRINGLITERAL:
			s = strings.Replace(s, `\`, `\\`, -1)
			s = `\"` + s[1:len(s)-1] + `\"`
		}
		a = append(a, s)
	}
	if len(s) != 0 {
		r = s[0]
	}
	r.hs = nil
	r.char = STRINGLITERAL
	r.value = dict.sid(`"` + strings.Join(a, "") + `"`)
	return r
}

func (c *cpp) hsAdd(hs hideSet, toks []cppToken) []cppToken {
	for i, v := range toks {
		if v.hs == nil {
			v.hs = hideSet{}
		}
		for k, w := range hs {
			v.hs[k] = w
		}
		v.fileID = c.fileID
		toks[i] = v
	}
	return toks
}

func (c *cpp) parseDefined(tok cppToken, r tokenReader, w tokenWriter) {
	toks := []cppToken{tok}
	if tok = c.scanToNonBlankToken(&toks, r, w); tok.char < 0 {
		return
	}

	switch tok.char {
	case IDENTIFIER:
		// ok
	case '(':
		if tok = c.scanToNonBlankToken(&toks, r, w); tok.char < 0 {
			return
		}

		if tok.char != IDENTIFIER {
			w.writes(toks)
			return
		}

		tok2 := c.scanToNonBlankToken(&toks, r, w)
		if tok2.char < 0 {
			return
		}

		if tok2.char != ')' {
			w.writes(toks)
			return
		}
	}

	tok.char = PPNUMBER
	switch _, ok := c.macros[tok.value]; {
	case ok:
		tok.value = idOne
	default:
		tok.value = idZero
	}
	w.write(tok)
}

func (c *cpp) scanToNonBlankToken(toks *[]cppToken, r tokenReader, w tokenWriter) cppToken {
	tok, ok := r.read()
	if !ok {
		w.writes(*toks)
		tok.char = -1
		return tok
	}

	*toks = append(*toks, tok)
	if tok.char == ' ' || tok.char == '\n' {
		if tok, ok = r.read(); !ok {
			w.writes(*toks)
			tok.char = -1
			return tok
		}

		*toks = append(*toks, tok)
	}
	return (*toks)[len(*toks)-1]
}

// [0], 6.10.1
func (c *cpp) evalInclusionCondition(expr []token3) (r bool) {
	if !c.intmaxChecked {
		if m := c.macros[idIntMaxWidth]; m != nil && len(m.repl) != 0 {
			val := c.eval(m.repl)
			if val != nil && val != int64(64) {
				c.err(m.name, "%s is %v, but only 64 is supported", idIntMaxWidth, val)
			}
		}
		c.intmaxChecked = true
	}

	val := c.eval(expr)
	return val != nil && c.isNonZero(val)
}

func (c *cpp) eval(expr []token3) interface{} {
	toks := make([]cppToken, len(expr))
	for i, v := range expr {
		toks[i] = cppToken{token4{token3: v}, nil}
	}
	var w cppWriter
	c.expand(&cppReader{buf: toks}, &w, true)
	toks = w.toks
	p := 0
	for _, v := range toks {
		switch v.char {
		case ' ', '\n':
			// nop
		default:
			toks[p] = v
			p++
		}
	}
	toks = toks[:p]
	s := cppScanner(toks)
	val := c.conditionalExpression(&s, true)
	switch s.peek().char {
	case -1, '#':
		// ok
	default:
		t := s.peek()
		c.err(t, "unexpected %s", tokName(t.char))
		return nil
	}
	return val
}

// [0], 6.5.15 Conditional operator
//
//  conditional-expression:
//		logical-OR-expression
//		logical-OR-expression ? expression : conditional-expression
func (c *cpp) conditionalExpression(s *cppScanner, eval bool) interface{} {
	expr := c.logicalOrExpression(s, eval)
	if s.peek().char == '?' {
		s.next()
		exprIsNonZero := c.isNonZero(expr)
		expr2 := c.conditionalExpression(s, exprIsNonZero)
		if tok := s.peek(); tok.char != ':' {
			c.err(tok, "expected ':'")
			return expr
		}

		s.next()
		expr3 := c.conditionalExpression(s, !exprIsNonZero)
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
func (c *cpp) logicalOrExpression(s *cppScanner, eval bool) interface{} {
	lhs := c.logicalAndExpression(s, eval)
	for s.peek().char == OROR {
		s.next()
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
func (c *cpp) logicalAndExpression(s *cppScanner, eval bool) interface{} {
	lhs := c.inclusiveOrExpression(s, eval)
	for s.peek().char == ANDAND {
		s.next()
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

func (c *cpp) isZero(val interface{}) bool {
	switch x := val.(type) {
	case int64:
		return x == 0
	case uint64:
		return x == 0
	}
	panic("internal error") //TODOOK
}

// [0], 6.5.12 Bitwise inclusive OR operator
//
//  inclusive-OR-expression:
//		exclusive-OR-expression
//		inclusive-OR-expression | exclusive-OR-expression
func (c *cpp) inclusiveOrExpression(s *cppScanner, eval bool) interface{} {
	lhs := c.exclusiveOrExpression(s, eval)
	for s.peek().char == '|' {
		s.next()
		rhs := c.exclusiveOrExpression(s, eval)
		if eval {
			switch x := lhs.(type) {
			case int64:
				switch y := rhs.(type) {
				case int64:
					lhs = x | y
				case uint64:
					lhs = uint64(x) | y
				}
			case uint64:
				switch y := rhs.(type) {
				case int64:
					lhs = x | uint64(y)
				case uint64:
					lhs = x | y
				}
			}
		}
	}
	return lhs
}

// [0], 6.5.11 Bitwise exclusive OR operator
//
//  exclusive-OR-expression:
//		AND-expression
//		exclusive-OR-expression ^ AND-expression
func (c *cpp) exclusiveOrExpression(s *cppScanner, eval bool) interface{} {
	lhs := c.andExpression(s, eval)
	for s.peek().char == '^' {
		s.next()
		rhs := c.andExpression(s, eval)
		if eval {
			switch x := lhs.(type) {
			case int64:
				switch y := rhs.(type) {
				case int64:
					lhs = x ^ y
				case uint64:
					lhs = uint64(x) ^ y
				}
			case uint64:
				switch y := rhs.(type) {
				case int64:
					lhs = x ^ uint64(y)
				case uint64:
					lhs = x ^ y
				}
			}
		}
	}
	return lhs
}

// [0], 6.5.10 Bitwise AND operator
//
//  AND-expression:
// 		equality-expression
// 		AND-expression & equality-expression
func (c *cpp) andExpression(s *cppScanner, eval bool) interface{} {
	lhs := c.equalityExpression(s, eval)
	for s.peek().char == '&' {
		s.next()
		rhs := c.equalityExpression(s, eval)
		if eval {
			switch x := lhs.(type) {
			case int64:
				switch y := rhs.(type) {
				case int64:
					lhs = x & y
				case uint64:
					lhs = uint64(x) & y
				}
			case uint64:
				switch y := rhs.(type) {
				case int64:
					lhs = x & uint64(y)
				case uint64:
					lhs = x & y
				}
			}
		}
	}
	return lhs
}

// [0], 6.5.9 Equality operators
//
//  equality-expression:
//		relational-expression
//		equality-expression == relational-expression
//		equality-expression != relational-expression
func (c *cpp) equalityExpression(s *cppScanner, eval bool) interface{} {
	lhs := c.relationalExpression(s, eval)
	for {
		var v bool
		switch s.peek().char {
		case EQ:
			s.next()
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
		case NEQ:
			s.next()
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
func (c *cpp) relationalExpression(s *cppScanner, eval bool) interface{} {
	lhs := c.shiftExpression(s, eval)
	for {
		var v bool
		switch s.peek().char {
		case '<':
			s.next()
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
			s.next()
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
		case LEQ:
			s.next()
			rhs := c.shiftExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						v = x <= y
					case uint64:
						v = uint64(x) <= y
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						v = x <= uint64(y)
					case uint64:
						v = x <= y
					}
				}
			}
		case GEQ:
			s.next()
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
func (c *cpp) shiftExpression(s *cppScanner, eval bool) interface{} {
	lhs := c.additiveExpression(s, eval)
	for {
		switch s.peek().char {
		case LSH:
			s.next()
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
		case RSH:
			s.next()
			rhs := c.additiveExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						lhs = x >> uint(y)
					case uint64:
						lhs = uint64(x) >> uint(y)
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						lhs = x >> uint(y)
					case uint64:
						lhs = x >> uint(y)
					}
				}
			}
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
func (c *cpp) additiveExpression(s *cppScanner, eval bool) interface{} {
	lhs := c.multiplicativeExpression(s, eval)
	for {
		switch s.peek().char {
		case '+':
			s.next()
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
			s.next()
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
func (c *cpp) multiplicativeExpression(s *cppScanner, eval bool) interface{} {
	lhs := c.unaryExpression(s, eval)
	for {
		switch s.peek().char {
		case '*':
			s.next()
			rhs := c.unaryExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						lhs = x * y
					case uint64:
						lhs = uint64(x) * y
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						lhs = x * uint64(y)
					case uint64:
						lhs = x * y
					}
				}
			}
		case '/':
			tok := s.next()
			rhs := c.unaryExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						if y == 0 {
							c.err(tok, "division by zero")
							break
						}

						lhs = x / y
					case uint64:
						if y == 0 {
							c.err(tok, "division by zero")
							break
						}

						lhs = uint64(x) / y
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						if y == 0 {
							c.err(tok, "division by zero")
							break
						}

						lhs = x / uint64(y)
					case uint64:
						if y == 0 {
							c.err(tok, "division by zero")
							break
						}

						lhs = x / y
					}
				}
			}
		case '%':
			tok := s.next()
			rhs := c.unaryExpression(s, eval)
			if eval {
				switch x := lhs.(type) {
				case int64:
					switch y := rhs.(type) {
					case int64:
						if y == 0 {
							c.err(tok, "division by zero")
							break
						}

						lhs = x % y
					case uint64:
						if y == 0 {
							c.err(tok, "division by zero")
							break
						}

						lhs = uint64(x) % y
					}
				case uint64:
					switch y := rhs.(type) {
					case int64:
						if y == 0 {
							c.err(tok, "division by zero")
							break
						}

						lhs = x % uint64(y)
					case uint64:
						if y == 0 {
							c.err(tok, "division by zero")
							break
						}

						lhs = x % y
					}
				}
			}
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
func (c *cpp) unaryExpression(s *cppScanner, eval bool) interface{} {
	switch s.peek().char {
	case '+':
		s.next()
		return c.unaryExpression(s, eval)
	case '-':
		s.next()
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
		s.next()
		expr := c.unaryExpression(s, eval)
		if eval {
			switch x := expr.(type) {
			case int64:
				expr = ^x
			case uint64:
				expr = ^x
			}
		}
		return expr
	case '!':
		s.next()
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
func (c *cpp) primaryExpression(s *cppScanner, eval bool) interface{} {
	switch tok := s.peek(); tok.char {
	case CHARCONST, LONGCHARCONST:
		s.next()
		r := charConst(c.ctx, tok)
		if r < 0 {
			r = -r
		}
		return int64(r)
	case IDENTIFIER:
		s.next()
		if s.peek().char == '(' {
			s.next()
			n := 1
		loop:
			for n != 0 {
				switch s.peek().char {
				case '(':
					n++
				case ')':
					n--
				case -1:
					c.err(s.peek(), "expected )")
					break loop
				}
				s.next()
			}
		}
		return int64(0)
	case PPNUMBER:
		s.next()
		return c.intConst(tok)
	case '(':
		s.next()
		expr := c.conditionalExpression(s, eval)
		if s.peek().char == ')' {
			s.next()
		}
		return expr
	default:
		return int64(0)
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
func (c *cpp) intConst(tok cppToken) (r interface{}) {
	var n uint64
	s0 := tok.String()
	s := strings.TrimRight(s0, "uUlL")
	switch {
	case strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X"):
		var err error
		if n, err = strconv.ParseUint(s[2:], 16, 64); err != nil {
			c.err(tok, "%v", err)
			return int64(0)
		}
	case strings.HasPrefix(s, "0"):
		var err error
		if n, err = strconv.ParseUint(s, 8, 64); err != nil {
			c.err(tok, "%v", err)
			return int64(0)
		}
	default:
		var err error
		if n, err = strconv.ParseUint(s, 10, 64); err != nil {
			c.err(tok, "%v", err)
			return int64(0)
		}
	}

	suffix := s0[len(s):]
	if suffix == "" {
		if n > math.MaxInt64 {
			return n
		}

		return int64(n)
	}

	switch suffix = strings.ToLower(suffix); suffix {
	default:
		c.err(tok, "invalid suffix: %v", s0)
		fallthrough
	case
		"l",
		"ll":

		if n > math.MaxInt64 {
			return n
		}

		return int64(n)
	case
		"llu",
		"lu",
		"u",
		"ul",
		"ull":

		return n
	}
}

func charConst(ctx *context, tok cppToken) rune {
	switch tok.char {
	case CHARCONST:
		s := tok.String()
		s = s[1 : len(s)-1] // Remove outer 's.
		if len(s) == 1 {
			return rune(s[0])
		}

		runes := []rune(s)
		var r rune
		switch runes[0] {
		case '\\':
			r, _ = decodeEscapeSequence(runes)
		default:
			r = runes[0]
		}
		return r
	case LONGCHARCONST:
		s := tok.String()
		var buf bytes.Buffer
		s = s[2 : len(s)-1]
		runes := []rune(s)
		for i := 0; i < len(runes); {
			switch r := runes[i]; {
			case r == '\\':
				r, n := decodeEscapeSequence(runes[i:])
				switch {
				case r < 0:
					buf.WriteByte(byte(-r))
				default:
					buf.WriteRune(r)
				}
				i += n
			default:
				buf.WriteByte(byte(r))
				i++
			}
		}
		s = buf.String()
		runes = []rune(s)
		if len(runes) != 1 {
			ctx.err(tok.Position(), "invalid character constant %s", &tok)
		}
		return runes[0]
	}
	panic("internal error") //TODOOK
}

// escape-sequence		{simple-sequence}|{octal-escape-sequence}|{hexadecimal-escape-sequence}|{universal-character-name}
// simple-sequence		\\['\x22?\\abfnrtv]
// octal-escape-sequence	\\{octal-digit}{octal-digit}?{octal-digit}?
// hexadecimal-escape-sequence	\\x{hexadecimal-digit}+
func decodeEscapeSequence(runes []rune) (rune, int) {
	if runes[0] != '\\' {
		panic("internal error") //TODOOK
	}

	if len(runes) < 2 {
		return runes[0], 1
	}

	r := runes[1]
	switch r {
	case '\'', '"', '?', '\\':
		return r, 2
	case 'a':
		return 7, 2
	case 'b':
		return 8, 2
	case 'f':
		return 12, 2
	case 'n':
		return 10, 2
	case 'r':
		return 13, 2
	case 't':
		return 9, 2
	case 'v':
		return 11, 2
	case 'x':
		v, n := 0, 2
	loop2:
		for _, r := range runes[2:] {
			switch {
			case r >= '0' && r <= '9', r >= 'a' && r <= 'f', r >= 'A' && r <= 'F':
				v = v<<4 | decodeHex(r)
				n++
			default:
				break loop2
			}
		}
		return -rune(v & 0xff), n
	case 'u', 'U':
		return decodeUCN(runes)
	}

	if r < '0' || r > '7' {
		panic("internal error") //TODOOK
	}

	v, n := 0, 1
loop:
	for _, r := range runes[1:] {
		switch {
		case r >= '0' && r <= '7':
			v = v<<3 | (int(r) - '0')
			n++
		default:
			break loop
		}
	}
	return -rune(v), n
}

// universal-character-name	\\u{hex-quad}|\\U{hex-quad}{hex-quad}
func decodeUCN(runes []rune) (rune, int) {
	if runes[0] != '\\' {
		panic("internal error") //TODOOK
	}
	runes = runes[1:]
	switch runes[0] {
	case 'u':
		return rune(decodeHexQuad(runes[1:])), 6
	case 'U':
		return rune(decodeHexQuad(runes[1:])<<16 | decodeHexQuad(runes[5:])), 10
	}
	panic("internal error") //TODOOK
}

// hex-quad	{hexadecimal-digit}{hexadecimal-digit}{hexadecimal-digit}{hexadecimal-digit}
func decodeHexQuad(runes []rune) int {
	n := 0
	for _, r := range runes[:4] {
		n = n<<4 | decodeHex(r)
	}
	return n
}

func decodeHex(r rune) int {
	switch {
	case r >= '0' && r <= '9':
		return int(r) - '0'
	default:
		x := int(r) &^ 0x20
		return x - 'A' + 10
	}
}

func (c *cpp) isNonZero(val interface{}) bool {
	switch x := val.(type) {
	case int64:
		return x != 0
	case uint64:
		return x != 0
	}
	panic("internal error") //TODOOK
}

type ppLine interface {
	getToks() []token3
}

type ppIfGroupDirective interface {
	evalInclusionCondition(*cpp) bool
}

type ppElifDirective struct {
	toks []token3
	expr []token3
}

func (n *ppElifDirective) getToks() []token3 { return n.toks }

type ppElseDirective struct {
	toks []token3
}

func (n *ppElseDirective) getToks() []token3 { return n.toks }

type ppEndifDirective struct {
	toks []token3
}

func (n *ppEndifDirective) getToks() []token3 { return n.toks }

type ppEmptyDirective struct {
	toks []token3
}

func (n *ppEmptyDirective) getToks() []token3 { return n.toks }

func (n *ppEmptyDirective) translationPhase4(c *cpp) {
	// nop
}

type ppIncludeDirective struct {
	arg  []token3
	toks []token3

	includeNext bool // false: #include, true: #include_next
}

func (n *ppIncludeDirective) getToks() []token3 { return n.toks }

func (n *ppIncludeDirective) translationPhase4(c *cpp) {
	if c.ctx.cfg.ignoreIncludes {
		return
	}

	args := make([]cppToken, len(n.arg))
	for i, v := range n.arg {
		args[i] = cppToken{token4{token3: v}, nil}
	}
	for len(args) != 0 && args[0].char == ' ' {
		args = args[1:]
	}
	var sb strings.Builder
	for _, v := range args {
		sb.WriteString(v.String())
	}
	nm := strings.TrimSpace(sb.String())
	if nm == "" {
		c.err(n.toks[0], "invalid empty include argument")
		return
	}

	switch nm[0] {
	case '"', '<':
		// ok
	default:
		var w cppWriter
		c.expand(&cppReader{buf: args}, &w, false)
		nm = strings.TrimSpace(cppToksStr(w.toks, ""))
	}
	toks := n.toks
	for toks[0].char == ' ' {
		toks = toks[1:]
	}
	if c.ctx.cfg.RejectIncludeNext {
		c.err(toks[0], "#include_next is a GCC extension")
		return
	}

	if c.ctx.cfg.fakeIncludes {
		c.send([]token3{{char: STRINGLITERAL, value: dict.sid(nm)}, {char: '\n', value: idNL}})
		return
	}

	if c.includeLevel == maxIncludeLevel {
		c.err(toks[0], "too many include levels")
		return
	}

	c.includeLevel++

	defer func() { c.includeLevel-- }()

	var b byte
	var paths []string
	switch {
	case nm != "" && nm[0] == '"':
		paths = c.ctx.includePaths
		b = '"'
	case nm != "" && nm[0] == '<':
		paths = c.ctx.sysIncludePaths
		b = '>'
	case nm == "":
		c.err(toks[0], "invalid empty include argument")
		return
	default:
		c.err(toks[0], "invalid include argument %s", nm)
		return
	}

	x := strings.IndexByte(nm[1:], b)
	if x < 0 {
		c.err(toks[0], "invalid include argument %s", nm)
		return
	}

	nm = filepath.FromSlash(nm[1 : x+1])
	dir := filepath.Dir(c.file.Name())
	if n.includeNext {
		nmDir, _ := filepath.Split(nm)
		for i, v := range paths {
			if w, err := filepath.Abs(v); err == nil {
				v = w
			}
			v = filepath.Join(v, nmDir)
			if v == dir {
				paths = paths[i+1:]
				break
			}
		}
	}
	var path string
	for _, v := range paths {
		if v == "@" {
			v = dir
		}

		var p string
		switch {
		case strings.HasPrefix(nm, "./"):
			wd := c.ctx.cfg.WorkingDir
			if wd == "" {
				var err error
				if wd, err = os.Getwd(); err != nil {
					c.err(toks[0], "cannot determine working dir: %v", err)
					return
				}
			}
			p = filepath.Join(wd, nm)
		default:
			p = filepath.Join(v, nm)
		}
		fi, err := os.Stat(p)
		if err != nil || fi.IsDir() {
			continue
		}

		path = p
		break
	}

	if path == "" {
		c.err(toks[0], "include file not found: %s\nsearch paths:\n\t%s", nm, strings.Join(paths, "\n\t"))
		return
	}

	cf, err := cache.getFile(c.ctx, path)
	if err != nil {
		c.err(toks[0], "%s: %v", path, err)
		return
	}

	pf, err := cf.ppFile()
	if err != nil {
		c.err(toks[0], "%s: %v", path, err)
		return
	}

	saveFile := c.file
	saveFileIF := c.fileID
	saveFileMacro := c.fileMacro.repl[0].value

	c.file = pf.file
	c.fileID = cf.fileID()
	c.fileMacro.repl[0].value = dict.sid(c.file.Name())

	defer func() {
		c.file = saveFile
		c.fileID = saveFileIF
		c.fileMacro.repl[0].value = saveFileMacro
	}()

	pf.translationPhase4(c)
}

func (c *cpp) send(toks []token3) {
	c.in <- toks
	<-c.rq
}

func (c *cpp) identicalReplacementLists(a, b []token3) bool {
	for len(a) != 0 && a[0].char == ' ' {
		a = a[1:]
	}
	for len(b) != 0 && b[0].char == ' ' {
		b = b[1:]
	}
	for len(a) != 0 && a[len(a)-1].char == ' ' {
		a = a[:len(a)-1]
	}
	for len(b) != 0 && b[len(b)-1].char == ' ' {
		b = b[:len(b)-1]
	}
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		w := b[i]
		if v.char != w.char || v.value != w.value {
			return false
		}
	}
	return true
}

func stringConst(t cppToken) string {
	s := t.String()
	switch t.char {
	case LONGSTRINGLITERAL:
		s = s[1:] // Remove leading 'L'.
		fallthrough
	case STRINGLITERAL:
		var buf bytes.Buffer
		s = s[1 : len(s)-1] // Remove outer "s.
		runes := []rune(s)
		for i := 0; i < len(runes); {
			switch r := runes[i]; {
			case r == '\\':
				r, n := decodeEscapeSequence(runes[i:])
				switch {
				case r < 0:
					buf.WriteByte(byte(-r))
				default:
					buf.WriteRune(r)
				}
				i += n
			default:
				buf.WriteByte(byte(r))
				i++
			}
		}
		return buf.String()
	}
	panic("internal error") //TODOOK
}

// -------------------------------------------------------- Translation phase 4

//TODO _Pragma

// [0], 5.1.1.2, 4
//
// Preprocessing directives are executed, macro invocations are expanded, and
// _Pragma unary operator expressions are executed. If a character sequence
// that matches the syntax of a universal character name is produced by token
// concatenation (6.10.3.3), the behavior is undefined. A #include
// preprocessing directive causes the named header or source file to be
// processed from phase 1 through phase 4, recursively. All preprocessing
// directives are then deleted.
func (c *cpp) translationPhase4(in []source) chan *[]token4 {
	c.rq = make(chan struct{})       // Must be unbufferred
	c.in = make(chan []token3)       // Must be unbufferred
	c.out = make(chan *[]token4, 10) //DONE benchmark tuned

	go func() {
		defer close(c.out)

		c.expand(c, c, false)
	}()

	go func() {
		defer close(c.in)

		for _, v := range in {
			pf, err := v.ppFile()
			if err != nil {
				c.err(nil, "%s", err)
				break
			}

			c.file = pf.file
			c.fileID = v.fileID()
			c.fileMacro.repl[0].value = dict.sid(c.file.Name())
			pf.translationPhase4(c)
		}
	}()

	return c.out
}

type ppErrorDirective struct {
	toks []token3
	msg  []token3
}

func (n *ppErrorDirective) getToks() []token3 { return n.toks }

func (n *ppErrorDirective) translationPhase4(c *cpp) {
	var b strings.Builder
	for _, v := range n.msg {
		b.WriteString(v.String())
	}
	c.err(n.toks[0], "%s", strings.TrimSpace(b.String()))
}

type ppPragmaDirective struct {
	toks []token3
	args []token3
}

func (n *ppPragmaDirective) getToks() []token3 { return n.toks }

func (n *ppPragmaDirective) translationPhase4(c *cpp) { parsePragma(c, n.args) }

func parsePragma(c *cpp, args0 []token3) {
	if len(args0) == 1 { // \n
		return
	}

	if t := args0[0]; t.char == IDENTIFIER && t.value == idSTDC {
		p := t
		p.char = PRAGMASTDC
		p.value = idPragmaSTDC
		send := []token3{p}
		args := ltrim(args0[1:])
		if len(args) == 0 {
			c.err(args[0], "expected argument of STDC")
			return
		}

		if t = args[0]; t.char != IDENTIFIER {
			c.err(t, "expected identifier")
			return
		}

		switch t.value {
		case idFPContract, idFenvAccess, idCxLimitedRange:
			// ok
		default:
			c.err(t, "expected FP_CONTRACT or FENV_ACCESS or CX_LIMITED_RANGE")
			return
		}

		args = ltrim(args[1:])
		if len(args) == 0 {
			c.err(args[0], "expected ON or OFF or DEFAULT")
			return
		}

		if t = args[0]; t.char != IDENTIFIER {
			c.err(t, "expected identifier")
			return
		}

		switch t.value {
		case idOn, idOff, idDefault:
			c.send(append(send, args0...))
		default:
			c.err(t, "expected ON or OFF or DEFAULT")
			return
		}
	}

	if c.ctx.cfg.PragmaHandler == nil {
		return
	}

	toks := expandArgs(c, args0[:len(args0)-1])
	if len(toks) == 0 {
		return
	}

	var toks2 []Token
	var sep StringID
	for _, tok := range toks {
		switch tok.char {
		case ' ', '\n':
			switch {
			case sep != 0:
				sep = dict.sid(sep.String() + tok.String())
			default:
				sep = tok.value
			}
		default:
			var t Token
			t.Rune = tok.char
			t.Sep = sep
			t.Value = tok.value
			t.fileID = tok.fileID
			t.pos = tok.pos
			toks2 = append(toks2, t)
			sep = 0
		}
	}
	if len(toks2) == 0 {
		return
	}

	// dbg("%v: %q", c.file.PositionFor(args0[0].Pos(), true), tokStr(toks2, "|"))
	c.ctx.cfg.PragmaHandler(&pragma{tok: toks[0], c: c}, toks2)
}

type ppNonDirective struct {
	toks []token3
}

func (n *ppNonDirective) getToks() []token3 { return n.toks }

func (n *ppNonDirective) translationPhase4(c *cpp) {
	// nop
}

type ppTextLine struct {
	toks []token3
}

func (n *ppTextLine) getToks() []token3 { return n.toks }

func (n *ppTextLine) translationPhase4(c *cpp) { c.send(n.toks) }

type ppLineDirective struct {
	toks []token3
	args []token3
}

func (n *ppLineDirective) getToks() []token3 { return n.toks }

func (n *ppLineDirective) translationPhase4(c *cpp) {
	toks := expandArgs(c, n.args)
	if len(toks) == 0 {
		return
	}

	switch t := toks[0]; t.char {
	case PPNUMBER:
		ln, err := strconv.ParseInt(t.String(), 10, 31)
		if err != nil || ln < 1 {
			c.err(t, "expected positive integer less or equal 2147483647")
			return
		}

		for len(toks) != 0 && toks[0].char == ' ' {
			toks = toks[1:]
		}
		if len(toks) == 1 {
			c.file.AddLineInfo(int(n.toks[len(n.toks)-1].pos), c.file.Name(), int(ln))
			return
		}

		toks = toks[1:]
		for len(toks) != 0 && toks[0].char == ' ' {
			toks = toks[1:]
		}
		if len(toks) == 0 {
			c.file.AddLineInfo(int(n.toks[len(n.toks)-1].pos), c.file.Name(), int(ln))
			return
		}

		switch t := toks[0]; t.char {
		case STRINGLITERAL:
			s := t.String()
			s = s[1 : len(s)-1]
			c.file.AddLineInfo(int(n.toks[len(n.toks)-1].pos), s, int(ln))
			for len(toks) != 0 && toks[0].char == ' ' {
				toks = toks[1:]
			}
			if len(toks) != 0 && c.ctx.cfg.RejectLineExtraTokens {
				c.err(toks[0], "expected new-line")
			}
		default:
			c.err(t, "expected string literal")
			return
		}
	default:
		c.err(toks[0], "expected integer literal")
		return
	}
}

func expandArgs(c *cpp, args []token3) []cppToken {
	var w cppWriter
	var toks []cppToken
	for _, v := range args {
		toks = append(toks, cppToken{token4: token4{fileID: c.fileID, token3: v}})
	}
	c.expand(&cppReader{buf: toks}, &w, true)
	return w.toks
}

type ppUndefDirective struct {
	name token3
	toks []token3
}

func (n *ppUndefDirective) getToks() []token3 { return n.toks }

func (n *ppUndefDirective) translationPhase4(c *cpp) {
	nm := n.name.value
	if _, ok := protectedMacros[nm]; ok || nm == idDefined {
		c.err(n.name, "cannot undefine a protected name")
		return
	}

	// dbg("#undef %s", nm)
	delete(c.macros, nm)
}

type ppIfdefDirective struct {
	name StringID
	toks []token3
}

func (n *ppIfdefDirective) evalInclusionCondition(c *cpp) bool { _, ok := c.macros[n.name]; return ok }

func (n *ppIfdefDirective) getToks() []token3 { return n.toks }

type ppIfndefDirective struct {
	name StringID
	toks []token3
}

func (n *ppIfndefDirective) evalInclusionCondition(c *cpp) bool { _, ok := c.macros[n.name]; return !ok }
func (n *ppIfndefDirective) getToks() []token3                  { return n.toks }

type ppIfDirective struct {
	toks []token3
	expr []token3
}

func (n *ppIfDirective) getToks() []token3 { return n.toks }

func (n *ppIfDirective) evalInclusionCondition(c *cpp) bool {
	return c.evalInclusionCondition(n.expr)
}

type ppDefineObjectMacroDirective struct {
	name            token3
	toks            []token3
	replacementList []token3
}

func (n *ppDefineObjectMacroDirective) getToks() []token3 { return n.toks }

func (n *ppDefineObjectMacroDirective) translationPhase4(c *cpp) {
	nm := n.name.value
	m := c.macros[nm]
	if m != nil {
		if _, ok := protectedMacros[nm]; ok || nm == idDefined {
			c.err(n.name, "cannot define protected name")
			return
		}

		if m.isFnLike {
			c.err(n.name, "redefinition of a function-like macro with an object-like one")
		}

		if !c.identicalReplacementLists(n.replacementList, m.repl) {
			c.err(n.name, "redefinition with different replacement list")
			return
		}
	}
	// dbg("#define %s %s", n.Name, tokStr(n.ReplacementList))
	c.macros[nm] = &macro{name: token4{token3: n.name, fileID: c.fileID}, repl: n.replacementList}
}

type ppDefineFunctionMacroDirective struct {
	identifierList  []token3
	toks            []token3
	replacementList []token3

	name token3

	namedVariadic bool // foo..., note no comma before ellipsis.
	variadic      bool
}

func (n *ppDefineFunctionMacroDirective) getToks() []token3 { return n.toks }

func (n *ppDefineFunctionMacroDirective) translationPhase4(c *cpp) {
	nm := n.name.value
	m := c.macros[nm]
	if m != nil {
		if _, ok := protectedMacros[nm]; ok || nm == idDefined {
			c.err(n.name, "cannot define protected name")
			return
		}

		if !m.isFnLike {
			c.err(n.name, "redefinition of an object-like macro with a function-like one")
			return
		}

		ok := len(m.fp) == len(n.identifierList)
		if ok {
			for i, v := range m.fp {
				if v != n.identifierList[i].value {
					ok = false
					break
				}
			}
		}
		if !ok && (len(n.replacementList) != 0 || len(m.repl) != 0) {
			c.err(n.name, "redefinition with different formal parameters")
			return
		}

		if !c.identicalReplacementLists(n.replacementList, m.repl) {
			c.err(n.name, "redefinition with different replacement list")
			return
		}

		if m.variadic != n.variadic {
			c.err(n.name, "redefinition differs in being variadic")
			return
		}
	}
	nms := map[StringID]struct{}{}
	for _, v := range n.identifierList {
		if _, ok := nms[v.value]; ok {
			c.err(v, "duplicate identifier %s", v.value)
		}
	}
	var fp []StringID
	for _, v := range n.identifierList {
		fp = append(fp, v.value)
	}
	// dbg("#define %s(%v) %s", n.Name, fp, tokStr(n.ReplacementList))
	c.macros[nm] = &macro{fp: fp, isFnLike: true, name: token4{token3: n.name, fileID: c.fileID}, repl: n.replacementList, variadic: n.variadic, namedVariadic: n.namedVariadic}
}

// [0], 6.10.1
//
//  elif-group:
//  		# elif constant-expression new-line group_opt
type ppElifGroup struct {
	elif   *ppElifDirective
	groups []ppGroup
}

func (n *ppElifGroup) evalInclusionCondition(c *cpp) bool {
	if !c.evalInclusionCondition(n.elif.expr) {
		return false
	}

	for _, v := range n.groups {
		v.translationPhase4(c)
	}
	return true
}

// [0], 6.10.1
//
//  else-group:
//  		# else new-line group_opt
type ppElseGroup struct {
	elseLine *ppElseDirective
	groups   []ppGroup
}

func (n *ppElseGroup) translationPhase4(c *cpp) {
	if n == nil {
		return
	}

	for _, v := range n.groups {
		v.translationPhase4(c)
	}
}

// [0], 6.10.1
//
//  PreprocessingFile:
//  		GroupOpt
type ppFile struct {
	file   *token.File
	groups []ppGroup
}

func (n *ppFile) translationPhase4(c *cpp) {
	c.ctx.tuSources++
	if f := n.file; f != nil {
		c.ctx.tuSize += int64(f.Size())
	}
	for _, v := range n.groups {
		v.translationPhase4(c)
	}
}

// [0], 6.10.1
//
//  group-part:
//  		if-section
//  		control-line
//  		text-line
//  		# non-directive
type ppGroup interface {
	translationPhase4(*cpp)
}

// [0], 6.10.1
//
//  if-group:
//  		# if constant-expression new-line group opt
//  		# ifdef identifier new-line group opt
//  		# ifndef identifier new-line group opt
type ppIfGroup struct {
	directive ppIfGroupDirective
	groups    []ppGroup
}

func (n *ppIfGroup) evalInclusionCondition(c *cpp) bool {
	if !n.directive.evalInclusionCondition(c) {
		return false
	}

	for _, v := range n.groups {
		v.translationPhase4(c)
	}
	return true
}

// [0], 6.10.1
//
// if-section:
// 		if-group elif-groups_opt else-group_opt endif-line
type ppIfSection struct {
	ifGroup    *ppIfGroup
	elifGroups []*ppElifGroup
	elseGroup  *ppElseGroup
	endifLine  *ppEndifDirective
}

func (n *ppIfSection) translationPhase4(c *cpp) {
	if !n.ifGroup.evalInclusionCondition(c) {
		for _, v := range n.elifGroups {
			if v.evalInclusionCondition(c) {
				return
			}
		}

		n.elseGroup.translationPhase4(c)
	}
}
