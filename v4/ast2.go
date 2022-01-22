// Copyright 2022 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

type AST struct {
	TranslationUnit *TranslationUnit
	EOF             Token
}

func (n *Declarator) Name() string {
	return n.DirectDeclarator.Name()
}

func (n *DirectDeclarator) Name() string {
	for n.DirectDeclarator != nil {
		n = n.DirectDeclarator
	}
	switch n.Case {
	case DirectDeclaratorIdent:
		return string(n.Token.Src())
	case DirectDeclaratorDecl:
		return n.Declarator.Name()
	default:
		panic(todo("internal error"))
	}
}
