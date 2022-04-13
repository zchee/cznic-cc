// Copyright 2022 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

// Name returns the name of n.
func (n *Declarator) Name() string {
	if n == nil {
		return ""
	}

	if dn := n.DirectDeclarator.name(); dn != nil {
		return dn.Token.SrcStr()
	}

	return ""
}

func (n *Declarator) isFn() bool {
	if n == nil {
		return false
	}

	return n.DirectDeclarator.isFn()
}

// Linkage descibes linkage of identifiers ([0]6.2.2).
type Linkage int

// Values of type Linkage
const (
	External Linkage = iota
	Internal
	None
)

func (n *Declarator) Linkage() Linkage {
	if n.IsTypename() {
		return None
	}

	if n.IsStatic() && n.LexicalScope().Parent == nil {
		return Internal
	}

	if n.IsExtern() || n.LexicalScope().Parent == nil {
		return External
	}

	return None
}

// IsExtern reports whether the storage class specifier 'extern' was present in
// the declaration of n.
func (n *Declarator) IsExtern() bool { return n.isExtern }

// IsConst reports whether the type qualifier 'const' was present in
// the declaration of n.
func (n *Declarator) IsConst() bool { return n.isConst }

// IsInline reports whether the function specifier 'inline' was present in the
// declaration of n.
func (n *Declarator) IsInline() bool { return n.isInline }

// IsVolatile reports whether the type qualifier 'volatile' was present in
// the declaration of n.
func (n *Declarator) IsVolatile() bool { return n.isVolatile }

// IsRegister reports whether the storage class specifier 'register' was
// present in the declaration of n.
func (n *Declarator) IsRegister() bool { return n.isRegister }

// IsStatic reports whether the storage class specifier 'static' was present in
// the declaration of n.
func (n *Declarator) IsStatic() bool { return n.isStatic }

// IsAtomic reports whether the type specifier '_Atomic' was present in the
// declaration of n.
func (n *Declarator) IsAtomic() bool { return n.isAtomic }

// IsThreadLocal reports whether the storage class specifier '_Thread_local'
// was present in the declaration of n.
func (n *Declarator) IsThreadLocal() bool { return n.isThreadLocal }

// IsTypename reports whether n is a typedef.
func (n *Declarator) IsTypename() bool { return n.isTypename }

// Alignas reports whether the _Alignas specifier was present in the
// declaration specifiers of n, if non-zero.
func (n *Declarator) Alignas() int { return n.alignas }

// IsParam reports whether n is a function paramater.
func (n *Declarator) IsParam() bool { return n.isParam }

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

// ResolvedTo returns the node n resolved to when n.Case is
// PrimaryExpressionIdent.
func (n *PrimaryExpression) ResolvedTo() Node { return n.resolvedTo }

// Macro returns the single token, object-like, constant macro that produced
// this node, if any.
func (n *PrimaryExpression) Macro() *Macro { return n.m }

// Associated returns the selected association of n, if any.
func (n *GenericSelection) Associated() *GenericAssociation { return n.assoc }

// Offset returns the offset of n within it's containing type.
func (n *Initializer) Offset() int64 { return n.off }

// Len returns the number of array elements initialized. It's normally one, but
// can be more using the [lo ... hi] designator.
func (n *Initializer) Len() int64 { return n.nelems }

// Field reports the resolved field for cases PostfixExpressionSelect and
// PostfixExpressionPSelect.
func (n *PostfixExpression) Field() *Field { return n.field }
