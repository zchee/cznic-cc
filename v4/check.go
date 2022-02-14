// Copyright 2022 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"sort"
	"strings"
)

type ctx struct {
	ast          *AST
	builtinTypes map[string]Type
	errors       errors
	intType      Type
}

func newCtx(ast *AST) *ctx {
	c := &ctx{ast: ast}
	c.builtinTypes = map[string]Type{
		ts2String([]TypeSpecifierCase{TypeSpecifierChar, TypeSpecifierSigned}):   c.newPredefinedType(Schar),
		ts2String([]TypeSpecifierCase{TypeSpecifierChar, TypeSpecifierUnsigned}): c.newPredefinedType(Uchar),
		ts2String([]TypeSpecifierCase{TypeSpecifierChar}):                        c.newPredefinedType(Char),
		ts2String([]TypeSpecifierCase{TypeSpecifierDouble, TypeSpecifierLong}):   c.newPredefinedType(LongDouble),
		ts2String([]TypeSpecifierCase{TypeSpecifierDouble}):                      c.newPredefinedType(Double),
		ts2String([]TypeSpecifierCase{TypeSpecifierFloat}):                       c.newPredefinedType(Float),
		ts2String([]TypeSpecifierCase{TypeSpecifierInt, TypeSpecifierLong}):      c.newPredefinedType(Long),
		ts2String([]TypeSpecifierCase{TypeSpecifierInt}):                         c.newPredefinedType(Int),
		ts2String([]TypeSpecifierCase{TypeSpecifierVoid}):                        c.newPredefinedType(Void),
	}
	c.intType = c.newPredefinedType(Int)
	return c
}

func (c *ctx) newPredefinedType(kind Kind) *PredefinedType { return newPredefinedType(c.ast, kind) }

type AST struct {
	TranslationUnit *TranslationUnit
	EOF             Token
}

func (n *AST) check() error {
	c := newCtx(n)
	for l := n.TranslationUnit; l != nil; l = l.TranslationUnit {
		l.ExternalDeclaration.check(c)
	}
	return c.errors.err()
}

func (n *ExternalDeclaration) check(c *ctx) {
	switch n.Case {
	case ExternalDeclarationFuncDef: // FunctionDefinition
		n.FunctionDefinition.check(c)
	case ExternalDeclarationDecl: // Declaration
		n.Declaration.check(c)
	case ExternalDeclarationAsmStmt: // AsmStatement
		c.errors.add(errorf("TODO %v", n.Case))
	case ExternalDeclarationEmpty: // ';'
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *FunctionDefinition) check(c *ctx) {
	n.Declarator.check(c, n.DeclarationSpecifiers.check(c, &n.Declarator.isExtern))
	n.CompoundStatement.check(c)
}

func (n *CompoundStatement) check(c *ctx) {
	if n == nil {
		return
	}

	//TODO c.errors.add(errorf("TODO"))
	// for l := n.BlockItemList; l != nil ; l = l.BlockItemList {
	// }
}

func (n *Declaration) check(c *ctx) {
	var isExtern bool
	t := n.DeclarationSpecifiers.check(c, &isExtern)
	for l := n.InitDeclaratorList; l != nil; l = l.InitDeclaratorList {
		l.InitDeclarator.check(c, t, isExtern)
	}
}

func (n *InitDeclarator) check(c *ctx, t Type, isExtern bool) {
	n.Declarator.isExtern = isExtern
	n.Declarator.check(c, t)
	if n.Asm != nil {
		c.errors.add(errorf("TODO"))
		return
	}

	n.Initializer.check(c)
}

func (n *Initializer) check(c *ctx) {
	if n == nil {
		return
	}

	switch n.Case {
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *Declarator) check(c *ctx, t Type) Type {
	n.typ = n.DirectDeclarator.check(c, n.Pointer.check(c, t))
	return n.typ
}

func (n *DirectDeclarator) check(c *ctx, t Type) Type {
	switch n.Case {
	case DirectDeclaratorIdent: // IDENTIFIER
		return t
	case DirectDeclaratorDecl: // '(' Declarator ')'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectDeclaratorArr: // DirectDeclarator '[' TypeQualifiers AssignmentExpression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectDeclaratorStaticArr: // DirectDeclarator '[' "static" TypeQualifiers AssignmentExpression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectDeclaratorArrStatic: // DirectDeclarator '[' TypeQualifiers "static" AssignmentExpression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectDeclaratorStar: // DirectDeclarator '[' TypeQualifiers '*' ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectDeclaratorFuncParam: // DirectDeclarator '(' ParameterTypeList ')'
		fp, isVariadic := n.ParameterTypeList.check(c)
		return n.DirectDeclarator.check(c, newFunctionType(c, t, fp, isVariadic))
	case DirectDeclaratorFuncIdent: // DirectDeclarator '(' IdentifierList ')'
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return t //TODO-
}

func (n *ParameterTypeList) check(c *ctx) (fp []*ParameterDeclaration, isVariadic bool) {
	if n == nil {
		return nil, false
	}

	fp = n.ParameterList.check(c)
	return fp, n.Case == ParameterTypeListVar
}

func (n *ParameterList) check(c *ctx) (fp []*ParameterDeclaration) {
	for ; n != nil; n = n.ParameterList {
		n.ParameterDeclaration.check(c)
		fp = append(fp, n.ParameterDeclaration)
	}
	return fp
}

func (n *ParameterDeclaration) check(c *ctx) {
	var isExtern bool
	switch n.Case {
	case ParameterDeclarationDecl: // DeclarationSpecifiers Declarator
		n.Declarator.isParam = true
		n.typ = n.Declarator.check(c, n.DeclarationSpecifiers.check(c, &isExtern))
		if isExtern {
			c.errors.add(errorf("%v: storage class specified for parameter: abc", n.Declarator.Position()))
		}
	case ParameterDeclarationAbstract: // DeclarationSpecifiers AbstractDeclarator
		n.typ = n.AbstractDeclarator.check(c, n.DeclarationSpecifiers.check(c, &isExtern))
		if isExtern {
			c.errors.add(errorf("%v: storage class specified for unnamed parameter", n.Position()))
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *AbstractDeclarator) check(c *ctx, t Type) Type {
	if n == nil {
		return t
	}

	n.typ = n.DirectAbstractDeclarator.check(c, n.Pointer.check(c, t))
	return n.typ
}

func (n *DirectAbstractDeclarator) check(c *ctx, t Type) Type {
	if n == nil {
		return t
	}

	switch n.Case {
	case DirectAbstractDeclaratorDecl: // '(' AbstractDeclarator ')'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectAbstractDeclaratorArr: // DirectAbstractDeclarator '[' TypeQualifiers AssignmentExpression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectAbstractDeclaratorStaticArr: // DirectAbstractDeclarator '[' "static" TypeQualifiers AssignmentExpression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectAbstractDeclaratorArrStatic: // DirectAbstractDeclarator '[' TypeQualifiers "static" AssignmentExpression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectAbstractDeclaratorArrStar: // DirectAbstractDeclarator '[' '*' ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectAbstractDeclaratorFunc: // DirectAbstractDeclarator '(' ParameterTypeList ')'
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return t //TODO-
}

func (n *Pointer) check(c *ctx, t Type) Type {
	if n == nil {
		return t
	}

	switch n.Case {
	case PointerTypeQual: // '*' TypeQualifiers
		return newPointerType(t)
	case PointerPtr: // '*' TypeQualifiers Pointer
		c.errors.add(errorf("TODO %v", n.Case))
	case PointerBlock: // '^' TypeQualifiers
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return t //TODO-
}

func ts2String(a []TypeSpecifierCase) string {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	var b strings.Builder
	for _, v := range a {
		b.WriteByte(byte(v))
	}
	return b.String()
}

func (n *DeclarationSpecifiers) check(c *ctx, isExtern *bool) (r Type) {
	if n == nil {
		return c.intType
	}

	var ts []TypeSpecifierCase
	var firstTypeSpecifier *TypeSpecifier
	for ; n != nil; n = n.DeclarationSpecifiers {
		switch n.Case {
		case DeclarationSpecifiersStorage: // StorageClassSpecifier DeclarationSpecifiers
			n.StorageClassSpecifier.check(c, isExtern)
		case DeclarationSpecifiersTypeSpec: // TypeSpecifier DeclarationSpecifiers
			ts = append(ts, n.TypeSpecifier.Case)
			if firstTypeSpecifier == nil {
				firstTypeSpecifier = n.TypeSpecifier
			}
			r = n.TypeSpecifier.check(c)
		case DeclarationSpecifiersTypeQual: // TypeQualifier DeclarationSpecifiers
			//TODO c.errors.add(errorf("TODO %v", n.Case))
		case DeclarationSpecifiersFunc: // FunctionSpecifier DeclarationSpecifiers
			//TODO c.errors.add(errorf("TODO %v", n.Case))
		case DeclarationSpecifiersAlignSpec: // AlignmentSpecifier DeclarationSpecifiers
			//TODO c.errors.add(errorf("TODO %v", n.Case))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	t, ok := c.builtinTypes[ts2String(ts)]
	if ok {
		return t
	}

	if len(ts) == 1 {
		switch ts[0] {
		case
			TypeSpecifierStructOrUnion,
			TypeSpecifierTypeName:

			return r
		default:
			c.errors.add(errorf("TODO %v", ts[0]))
			return nil
		}
	}

	c.errors.add(errorf("TODO %v", ts))
	return nil
}

func (n *StorageClassSpecifier) check(c *ctx, isExtern *bool) {
	switch n.Case {
	case StorageClassSpecifierTypedef: // "typedef"
		// ok
	case StorageClassSpecifierExtern: // "extern"
		*isExtern = true
	case StorageClassSpecifierStatic: // "static"
		c.errors.add(errorf("TODO %v", n.Case))
	case StorageClassSpecifierAuto: // "auto"
		c.errors.add(errorf("TODO %v", n.Case))
	case StorageClassSpecifierRegister: // "register"
		c.errors.add(errorf("TODO %v", n.Case))
	case StorageClassSpecifierThreadLocal: // "_Thread_local"
		c.errors.add(errorf("TODO %v", n.Case))
	case StorageClassSpecifierDeclspec: // "__declspec" '(' ')'
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *TypeSpecifier) check(c *ctx) Type {
	switch n.Case {
	case TypeSpecifierVoid: // "void"
		// ok
	case TypeSpecifierChar: // "char"
		// ok
	case TypeSpecifierShort: // "short"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierInt: // "int"
		// ok
	case TypeSpecifierInt128: // "__int128"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierUint128: // "__uint128_t"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierLong: // "long"
		// ok
	case TypeSpecifierFloat: // "float"
		// ok
	case TypeSpecifierFloat16: // "_Float16"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierDecimal64: // "_Decimal64"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierFloat128: // "_Float128"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierFloat128x: // "_Float128x"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierDouble: // "double"
		// ok
	case TypeSpecifierSigned: // "signed"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierUnsigned: // "unsigned"
		// ok
	case TypeSpecifierBool: // "_Bool"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierComplex: // "_Complex"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierImaginary: // "_Imaginary"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierStructOrUnion: // StructOrUnionSpecifier
		return n.StructOrUnionSpecifier.check(c)
	case TypeSpecifierEnum: // EnumSpecifier
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierTypeName: // TYPENAME
		if x, ok := n.resolutionScope.ident(n.Token).(*Declarator); ok && x.typename {
			return x.typ
		}

		c.errors.add(errorf("%v: undefined type name", n.Position()))
	case TypeSpecifierTypeofExpr: // "typeof" '(' Expression ')'
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierTypeofType: // "typeof" '(' TypeName ')'
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierAtomic: // AtomicTypeSpecifier
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierFloat32: // "_Float32"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierFloat64: // "_Float64"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierFloat32x: // "_Float32x"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierFloat64x: // "_Float64x"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierM256d: // "__m256d"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierM128: // "__m128"
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return nil
}

func (n *StructOrUnionSpecifier) check(c *ctx) Type {
	switch n.Case {
	case StructOrUnionSpecifierDef: // StructOrUnion IDENTIFIER '{' StructDeclarationList '}'
		n.typ = n.StructDeclarationList.check(c)
		return n.typ
	case StructOrUnionSpecifierTag: // StructOrUnion IDENTIFIER
		if x := n.resolutionScope.structOrUnion(n.Token); x != nil {
			if n.StructOrUnion.Case != x.StructOrUnion.Case {
				c.errors.add(errorf("%v: mismateched struct/union tag", n.Token.Position()))
				break
			}

			return x.typ
		}

		c.errors.add(errorf("%v: undefined struct/union tag", n.Token.Position()))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return nil //TODO-
}

func (n *StructDeclarationList) check(c *ctx) Type {
	c.errors.add(errorf("TODO"))
	return nil
}
