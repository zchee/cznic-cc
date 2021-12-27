// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

// cppParser produces the AST based on [0]6.10.
type cppParser struct {
	s   *scanner
	eof []Token

	closed bool
}

// newCppParser returns a newly created cppParser.
func newCppParser(s *scanner) *cppParser {
	return &cppParser{s: s}
}

// scanLine scans a line, up to and including the final '\n' token. After all
// lines are produced, scanLine always returns line consisting of a single EOF
// token, with .Ch set to -1. This EOF token still has proper position and
// separator information.
func (p *cppParser) line() (r []Token) {
	if p.closed {
		return p.eof
	}

	var tok Token
	for tok.Ch != '\n' {
		tok = p.s.cppScan()
		if tok.Ch < 0 {
			p.eof = []Token{tok}
			p.closed = true
			if len(r) != 0 {
				if r[len(r)-1].Ch != '\n' {
					nl := tok
					nl.Ch = '\n'
					nl.Set(nil, []byte{'\n'})
					r = append(r, nl)
				}
				return r
			}

			return p.eof
		}

		r = append(r, tok)
	}
	return r
}
