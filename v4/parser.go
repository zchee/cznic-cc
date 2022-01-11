// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"strconv"

	"modernc.org/token"
)

// cppParser produces a preprocessingFile.
type cppParser struct {
	s    *scanner
	eh   errHandler
	line []Token // nil when consumed

	closed bool
}

// newCppParser returns a newly created cppParser. The errHandler function is invoked on
// parser errors.
func newCppParser(src Source, eh errHandler) (*cppParser, error) {
	s, err := newScanner(src, eh)
	if err != nil {
		return nil, err
	}

	return &cppParser{s: s, eh: eh}, nil
}

// close causes all subsequent calls to .line to signal EOF.
func (p *cppParser) close() {
	if p.closed {
		return
	}

	if len(p.line) == 0 || p.line[0].Ch != eof {
		p.line = []Token{p.s.newToken(eof)}
	}
	p.closed = true
}

func (p *cppParser) pos() token.Position {
	p.c()
	return p.line[0].Position()
}

// c sets p.line to the line p is currently positioned on, up to and including
// its final '\n' token. After all lines are produced or p is closed, c always
// returns line consisting of a single EOF token, with .Ch set to eof.  This
// EOF token still has proper position and separator information if the end of
// file was reached normally, while parsing.
//
// c returns p.line[0].Ch.
func (p *cppParser) c() rune {
	if p.line != nil || p.closed {
		return p.line[0].Ch
	}

	var tok Token
	for tok.Ch != '\n' {
		tok = p.s.cppScan()
		if tok.Ch == eof {
			p.line = []Token{tok}
			p.closed = true
			break
		}

		p.line = append(p.line, tok)
	}
	return p.line[0].Ch
}

// consume returns p.line and if p is not closed, sets p.line to nil
func (p *cppParser) consume() (r []Token) {
	r = p.line
	if !p.closed {
		p.line = nil
	}
	return r
}

// preprocessingFile produces the AST based on [0]6.10.
//
// preprocessing-file: group_opt
func (p *cppParser) preprocessingFile() group { return p.group(false) }

// group:
//	group-part
//	group group-part
type group []groupPart

// group:
//	group-part
//	group group-part
func (p *cppParser) group(inIfSection bool) (r group) {
	for {
		g := p.groupPart(inIfSection)
		if g == nil {
			break
		}

		r = append(r, g)
		if _, ok := g.(eofLine); ok {
			break
		}
	}
	return r
}

// group-part:
// 	if-section
// 	control-line
// 	text-line
// 	# non-directive
// 	eof-line
type groupPart interface{}

// groupPart parses a group-part.
func (p *cppParser) groupPart(inIfSection bool) groupPart {
	switch p.c() {
	case '#':
		switch verb := string(p.line[1].Src()); verb {
		case "if", "ifdef", "ifndef":
			return p.ifSection()
		case "include", "include_next", "define", "undef", "line", "error", "pragma", "\n":
			gp := p.consume()
			switch {
			case verb == "line":
				// eg. ["#" "line" "1" "\"20010206-1.c\"" "\n"].5
				if len(gp) < 3 {
					break
				}

				ln, err := strconv.ParseUint(string(gp[2].Src()), 10, 31)
				if err != nil {
					break
				}

				fn := gp[0].Position().Filename
				if len(gp) >= 4 && gp[3].Ch == rune(STRINGLITERAL) {
					fn = string(gp[3].Src())
					fn = fn[1 : len(fn)-1]
				}

				nl := gp[len(gp)-1]
				pos := nl.Position()
				p.s.s.file.AddLineInfo(int(pos.Offset+1), fn, int(ln))
			}
			return controlLine(gp)
		case "elif", "else", "endif":
			if inIfSection {
				return nil
			}

			p.eh("%v: unexpected #%s", p.pos(), p.line[1].Src())
			return textLine(p.consume())
		default:
			return nonDirective(p.consume())
		}
	case eof:
		return eofLine(p.consume()[0])
	default:
		return textLine(p.consume())
	}
}

// controlLine parses an control-line. At unexpected eof it returns ok == false.
//
// control-line:
// 	# include pp-tokens new-line
// 	# define identifier replacement-list new-line
// 	# define identifier lparen identifier-list_opt ) replacement-list new-line
// 	# define identifier lparen ... ) replacement-list new-line
// 	# define identifier lparen identifier-list , ... ) replacement-list new-line
// 	# undef identifier new-line
// 	# line pp-tokens new-line
// 	# error pp-tokens_opt new-line
// 	# pragma pp-tokens_opt new-line
// 	# new-line
type controlLine []Token

// textLine is a groupPart representing a source line not starting with '#'.
type textLine []Token

// eofLine is a groupPart representing the end of a file
type eofLine Token

// non-directive is a groupPart representing a source line starting with '#'
// but not followed by any recognized token.
type nonDirective []Token

// if-section:
//	if-group elif-groups_opt else-group_opt endif-line
type ifSection struct {
	ifGroup    *ifGroup
	elifGroups []elifGroup
	elseGroup  *elseGroup
	endifLine  []Token
}

// ifSection parses an if-section.
func (p *cppParser) ifSection() *ifSection {
	return &ifSection{
		ifGroup:    p.ifGroup(),
		elifGroups: p.elifGroups(),
		elseGroup:  p.elseGroup(),
		endifLine:  p.endifLine(),
	}
}

// endifLine parses:

// endif-line:
// 	# endif new-line
func (p *cppParser) endifLine() []Token {
	if p.c() != '#' || string(p.line[1].Src()) != "endif" {
		p.eh("%v: expected #endif", p.pos())
		return nil
	}

	return p.consume()
}

// else-group:
//	# else new-line group_opt
type elseGroup struct {
	line  []Token
	group group
}

// elseGroup parses else-group.
func (p *cppParser) elseGroup() (r *elseGroup) {
	if p.c() == '#' && string(p.line[1].Src()) == "else" {
		return &elseGroup{p.consume(), p.group(true)}
	}

	return nil
}

// elif-group:
//	# elif constant-expression new-line group_opt
type elifGroup struct {
	line  []Token
	group group
}

// elifGroups parses:
//
// elif-groups:
//	elif-group
//	elif-groups elif-group
func (p *cppParser) elifGroups() (r []elifGroup) {
	for p.c() == '#' && string(p.line[1].Src()) == "elif" {
		r = append(r, elifGroup{p.consume(), p.group(true)})
	}
	return r
}

// if-group:
//	# if constant-expression new-line group_opt
//	# ifdef identifier new-line group_opt
//	# ifndef identifier new-line group_opt
type ifGroup struct {
	line  []Token
	group group
}

// ifGroup parses an if-group.
func (p *cppParser) ifGroup() (r *ifGroup) { return &ifGroup{p.consume(), p.group(true)} }
