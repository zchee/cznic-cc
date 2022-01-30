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
	if n == nil {
		return ""
	}

	return n.DirectDeclarator.Name()
}

func (n *Declarator) isName() bool {
	if n == nil {
		return false
	}

	return n.DirectDeclarator.isName()
}

func (n *Declarator) isFn() bool {
	if n == nil {
		return false
	}

	return n.DirectDeclarator.isFn()
}

// Name returns the name of n.
func (n *DirectDeclarator) Name() string {
	if n == nil {
		return ""
	}

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

func (n *DirectDeclarator) isName() bool {
	if n == nil {
		return false
	}

	switch n.Case {
	case DirectDeclaratorIdent:
		return true
	case DirectDeclaratorDecl:
		return n.Declarator.isName()
	default:
		panic(todo("internal error"))
	}
}

func (n *DirectDeclarator) isFn() bool {
	if n == nil || !n.DirectDeclarator.isName() {
		return false
	}

	switch n.Case {
	case DirectDeclaratorFuncParam, DirectDeclaratorFuncIdent:
		return true
	}

	return false
}

func (n *DeclarationSpecifiers) attributeValueList() *AttributeSpecifierList {
	for ; n != nil; n = n.DeclarationSpecifiers {
		if n.Case == DeclarationSpecifiersAttr {
			return n.AttributeSpecifierList
		}
	}
	return nil
}
