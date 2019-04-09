// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"modernc.org/token"
)

var (
	_ Node              = (*AST)(nil)
	_ TypeSpecification = (*DeclarationSpecifiers)(nil)
	_ TypeSpecification = (*SpecifierQualifierList)(nil)
)

// TypeSpecification is either nil or one of: *DeclarationSpecifiers,
// *SpecifierQualifierList.
type TypeSpecification interface{ isTypeSpecification() }

// LexicalScope returns the lexical scope of n.
func (n *AttributeValue) LexicalScope() Scope { return n.lexicalScope }

// AST represents a translation unit with related data.
type AST struct {
	Scope             Scope    // File scope.
	TrailingSeperator StringID // White space and/or comments preceding EOF.
	TranslationUnit   *TranslationUnit
}

// Position implements Node.
func (n *AST) Position() token.Position { return n.TranslationUnit.Position() }

// Scope returns n's scope.
func (n *CompoundStatement) Scope() Scope { return n.scope }

func (n *Declarator) isVisible(at int32) bool { return n.DirectDeclarator.ends() < at }

// ParamScope returns the scope in which n's function parameters are declared
// if the underlying type of n is a function or nil otherwise. If n is part of
// a function definition the scope is the same as the scope of the function
// body.
func (n *Declarator) ParamScope() Scope {
	if n == nil {
		return nil
	}

	return n.DirectDeclarator.ParamScope()
}

// Name returns n's declared name.
func (n *Declarator) Name() StringID {
	if n.DirectDeclarator == nil {
		return 0
	}

	return n.DirectDeclarator.Name()
}

// TypeSpecification returns n's type specification.
func (n *Declarator) TypeSpecification() TypeSpecification { return n.typeSpecification }

// LexicalScope returns the lexical scope of n.
func (n *Designator) LexicalScope() Scope { return n.lexicalScope }

// Name returns n's declared name.
func (n *DirectDeclarator) Name() StringID {
	for {
		if n == nil {
			return 0
		}

		switch n.Case {
		case DirectDeclaratorIdent: // IDENTIFIER
			return n.Token.Value
		case DirectDeclaratorDecl: // '(' Declarator ')'
			return n.Declarator.Name()
		default:
			n = n.DirectDeclarator
		}
	}
}

func (n *DirectDeclarator) ends() int32 {
	switch n.Case {
	case DirectDeclaratorIdent: // IDENTIFIER
		return n.Token.seq
	case DirectDeclaratorDecl: // '(' Declarator ')'
		return n.Token2.seq
	case DirectDeclaratorArr: // DirectDeclarator '[' TypeQualifierList AssignmentExpression ']'
		return n.Token2.seq
	case DirectDeclaratorStaticArr: // DirectDeclarator '[' "static" TypeQualifierList AssignmentExpression ']'
		return n.Token3.seq
	case DirectDeclaratorArrStatic: // DirectDeclarator '[' TypeQualifierList "static" AssignmentExpression ']'
		return n.Token3.seq
	case DirectDeclaratorStar: // DirectDeclarator '[' TypeQualifierList '*' ']'
		return n.Token3.seq
	case DirectDeclaratorFuncParam: // DirectDeclarator '(' ParameterTypeList ')'
		return n.Token2.seq
	case DirectDeclaratorFuncIdent: // DirectDeclarator '(' IdentifierList ')'
		return n.Token2.seq
	default:
		panic("internal error") //TODOOK
	}
}

// ParamScope returns the innermost scope in which function parameters are
// declared for Case DirectDeclaratorFuncParam or DirectDeclaratorFuncIdent or
// nil otherwise.
func (n *DirectDeclarator) ParamScope() Scope {
	if n == nil {
		return nil
	}

	switch n.Case {
	case DirectDeclaratorIdent: // IDENTIFIER
		return nil
	case DirectDeclaratorDecl: // '(' Declarator ')'
		return n.Declarator.ParamScope()
	case DirectDeclaratorArr: // DirectDeclarator '[' TypeQualifierList AssignmentExpression ']'
		return n.DirectDeclarator.ParamScope()
	case DirectDeclaratorStaticArr: // DirectDeclarator '[' "static" TypeQualifierList AssignmentExpression ']'
		return n.DirectDeclarator.ParamScope()
	case DirectDeclaratorArrStatic: // DirectDeclarator '[' TypeQualifierList "static" AssignmentExpression ']'
		return n.DirectDeclarator.ParamScope()
	case DirectDeclaratorStar: // DirectDeclarator '[' TypeQualifierList '*' ']'
		return n.DirectDeclarator.ParamScope()
	case DirectDeclaratorFuncParam: // DirectDeclarator '(' ParameterTypeList ')'
		if s := n.DirectDeclarator.ParamScope(); s != nil {
			return s
		}

		return n.paramScope
	case DirectDeclaratorFuncIdent: // DirectDeclarator '(' IdentifierList ')'
		if s := n.DirectDeclarator.ParamScope(); s != nil {
			return s
		}

		return n.paramScope
	default:
		panic("internal error") //TODOOK
	}
}

// LexicalScope returns the lexical scope of n.
func (n *DirectDeclarator) LexicalScope() Scope { return n.lexicalScope }

func (n *DeclarationSpecifiers) isTypeSpecification() {}

func (n *DeclarationSpecifiers) isTypedef() bool {
	for n != nil {
		if n.StorageClassSpecifier.isTypedef() {
			return true
		}

		n = n.DeclarationSpecifiers
	}
	return false
}

func (n *Enumerator) isVisible(at int32) bool { return n.Token.seq < at }

// LexicalScope returns the lexical scope of n.
func (n *EnumSpecifier) LexicalScope() Scope { return n.lexicalScope }

// LexicalScope returns the lexical scope of n.
func (n *IdentifierList) LexicalScope() Scope { return n.lexicalScope }

// LexicalScope returns the lexical scope of n.
func (n *JumpStatement) LexicalScope() Scope { return n.lexicalScope }

// LexicalScope returns the lexical scope of n.
func (n *LabeledStatement) LexicalScope() Scope { return n.lexicalScope }

// ResolvedIn reports which scope the identifier of cases
// PrimaryExpressionIdent, PrimaryExpressionEnum were resolved in, if any.
func (n *PrimaryExpression) ResolvedIn() Scope { return n.resolvedIn }

// LexicalScope returns the lexical scope of n.
func (n *PrimaryExpression) LexicalScope() Scope { return n.lexicalScope }

func (n *SpecifierQualifierList) isTypeSpecification() {}

func (n *StorageClassSpecifier) isTypedef() bool {
	return n != nil && n.Case == StorageClassSpecifierTypedef
}

// LexicalScope returns the lexical scope of n.
func (n *StructOrUnionSpecifier) LexicalScope() Scope { return n.lexicalScope }

// ResolvedIn reports which scope the identifier of case
// TypeSpecifierTypedefName was resolved in, if any.
func (n *TypeSpecifier) ResolvedIn() Scope { return n.resolvedIn }

// LexicalScope returns the lexical scope of n.
func (n *UnaryExpression) LexicalScope() Scope { return n.lexicalScope }
