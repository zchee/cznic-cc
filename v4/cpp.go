// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

// cpp is the C preprocessor.
type cpp struct {
	cache   map[string]group
	eh      errHandler
	sources []Source
	tok     Token // phase 7

	closed bool
}

func newCPP(sources []Source, eh errHandler) *cpp {
	r := &cpp{
		eh:      eh,
		sources: sources,
	}
	r.tok.Ch = eof // Invalidate
	return r
}
