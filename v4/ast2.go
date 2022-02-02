// Copyright 2022 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

const (
	DeclarationSpecifiersAttr = DeclarationSpecifiersFunc + 1
	TypeQualifierAttr         = TypeQualifierAtomic + 1
)

type AST struct {
	TranslationUnit *TranslationUnit
	EOF             Token
}

// Name returns the name of n.
func (n *Declarator) Name() string {
	if dn := n.DirectDeclarator.name(); dn != nil {
		return string(dn.Token.Src())
	}

	return ""
}

func (n *Declarator) isFn() bool {
	if n == nil {
		return false
	}

	return n.DirectDeclarator.isFn()
}

func (n *Declarator) fnParams() *Scope { return n.DirectDeclarator.params }

func (n *DirectDeclarator) name() *DirectDeclarator {
	if n == nil {
		return nil
	}

	switch n.Case {
	case DirectDeclaratorIdent:
		return n
	case DirectDeclaratorDecl:
		return n.Declarator.DirectDeclarator.name()
	default:
		return n.DirectDeclarator.name()
	}
}

func (n *DirectDeclarator) isFn() bool {
	if n == nil {
		return false
	}

	switch n.Case {
	case DirectDeclaratorFuncParam, DirectDeclaratorFuncIdent:
		return true
	}

	return false
}
