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
	complexdouble := c.newPredefinedType(ComplexDouble)
	complexlong := c.newPredefinedType(ComplexLong)
	int := c.newPredefinedType(Int)
	long := c.newPredefinedType(Long)
	longlong := c.newPredefinedType(LongLong)
	short := c.newPredefinedType(Short)
	uint := c.newPredefinedType(UInt)
	uint128 := c.newPredefinedType(UInt128)
	ulong := c.newPredefinedType(ULong)
	ulonglog := c.newPredefinedType(ULongLong)
	ushort := c.newPredefinedType(UShort)
	c.builtinTypes = map[string]Type{
		ts2String([]TypeSpecifierCase{TypeSpecifierBool}):                                                             c.newPredefinedType(Bool),
		ts2String([]TypeSpecifierCase{TypeSpecifierChar, TypeSpecifierSigned}):                                        c.newPredefinedType(SChar),
		ts2String([]TypeSpecifierCase{TypeSpecifierChar, TypeSpecifierUnsigned}):                                      c.newPredefinedType(UChar),
		ts2String([]TypeSpecifierCase{TypeSpecifierChar}):                                                             c.newPredefinedType(Char),
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierChar}):                                       c.newPredefinedType(ComplexChar),
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierDouble}):                                     complexdouble,
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierFloat}):                                      c.newPredefinedType(ComplexFloat),
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierInt, TypeSpecifierLong}):                     complexlong,
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierInt}):                                        c.newPredefinedType(ComplexInt),
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierLong, TypeSpecifierDouble}):                  c.newPredefinedType(ComplexLongDouble),
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierLong, TypeSpecifierLong}):                    c.newPredefinedType(ComplexLongLong),
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierLong}):                                       complexlong,
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierShort}):                                      c.newPredefinedType(ComplexShort),
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierUnsigned}):                                   c.newPredefinedType(ComplexUInt),
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex}):                                                          complexdouble,
		ts2String([]TypeSpecifierCase{TypeSpecifierDouble, TypeSpecifierLong}):                                        c.newPredefinedType(LongDouble),
		ts2String([]TypeSpecifierCase{TypeSpecifierDouble}):                                                           c.newPredefinedType(Double),
		ts2String([]TypeSpecifierCase{TypeSpecifierFloat128}):                                                         c.newPredefinedType(Float128),
		ts2String([]TypeSpecifierCase{TypeSpecifierFloat32x}):                                                         c.newPredefinedType(Float32x),
		ts2String([]TypeSpecifierCase{TypeSpecifierFloat32}):                                                          c.newPredefinedType(Float32),
		ts2String([]TypeSpecifierCase{TypeSpecifierFloat64x}):                                                         c.newPredefinedType(Float64x),
		ts2String([]TypeSpecifierCase{TypeSpecifierFloat64}):                                                          c.newPredefinedType(Float64),
		ts2String([]TypeSpecifierCase{TypeSpecifierFloat}):                                                            c.newPredefinedType(Float),
		ts2String([]TypeSpecifierCase{TypeSpecifierInt, TypeSpecifierLong, TypeSpecifierLong, TypeSpecifierSigned}):   longlong,
		ts2String([]TypeSpecifierCase{TypeSpecifierInt, TypeSpecifierLong, TypeSpecifierLong, TypeSpecifierUnsigned}): ulonglog,
		ts2String([]TypeSpecifierCase{TypeSpecifierInt, TypeSpecifierLong, TypeSpecifierLong}):                        longlong,
		ts2String([]TypeSpecifierCase{TypeSpecifierInt, TypeSpecifierLong, TypeSpecifierSigned}):                      long,
		ts2String([]TypeSpecifierCase{TypeSpecifierInt, TypeSpecifierLong, TypeSpecifierUnsigned}):                    ulong,
		ts2String([]TypeSpecifierCase{TypeSpecifierInt, TypeSpecifierLong}):                                           long,
		ts2String([]TypeSpecifierCase{TypeSpecifierInt, TypeSpecifierShort, TypeSpecifierSigned}):                     short,
		ts2String([]TypeSpecifierCase{TypeSpecifierInt, TypeSpecifierShort, TypeSpecifierUnsigned}):                   ushort,
		ts2String([]TypeSpecifierCase{TypeSpecifierInt, TypeSpecifierShort}):                                          short,
		ts2String([]TypeSpecifierCase{TypeSpecifierInt, TypeSpecifierSigned}):                                         int,
		ts2String([]TypeSpecifierCase{TypeSpecifierInt, TypeSpecifierUnsigned}):                                       uint,
		ts2String([]TypeSpecifierCase{TypeSpecifierInt128, TypeSpecifierUnsigned}):                                    uint128,
		ts2String([]TypeSpecifierCase{TypeSpecifierInt128}):                                                           c.newPredefinedType(Int128),
		ts2String([]TypeSpecifierCase{TypeSpecifierInt}):                                                              int,
		ts2String([]TypeSpecifierCase{TypeSpecifierLong, TypeSpecifierLong, TypeSpecifierSigned}):                     longlong,
		ts2String([]TypeSpecifierCase{TypeSpecifierLong, TypeSpecifierLong, TypeSpecifierUnsigned}):                   ulonglog,
		ts2String([]TypeSpecifierCase{TypeSpecifierLong, TypeSpecifierLong}):                                          longlong,
		ts2String([]TypeSpecifierCase{TypeSpecifierLong, TypeSpecifierSigned}):                                        long,
		ts2String([]TypeSpecifierCase{TypeSpecifierLong, TypeSpecifierUnsigned}):                                      ulong,
		ts2String([]TypeSpecifierCase{TypeSpecifierLong}):                                                             long,
		ts2String([]TypeSpecifierCase{TypeSpecifierShort, TypeSpecifierSigned}):                                       short,
		ts2String([]TypeSpecifierCase{TypeSpecifierShort, TypeSpecifierUnsigned}):                                     ushort,
		ts2String([]TypeSpecifierCase{TypeSpecifierShort}):                                                            short,
		ts2String([]TypeSpecifierCase{TypeSpecifierSigned}):                                                           int,
		ts2String([]TypeSpecifierCase{TypeSpecifierUint128}):                                                          uint128,
		ts2String([]TypeSpecifierCase{TypeSpecifierUnsigned}):                                                         uint,
		ts2String([]TypeSpecifierCase{TypeSpecifierVoid}):                                                             c.newPredefinedType(Void),
	}
	c.intType = int
	return c
}

func (c *ctx) newPredefinedType(kind Kind) *PredefinedType { return newPredefinedType(c.ast, kind) }

type typer struct{ typ Type }

func (t typer) Type() (r Type) { return t.typ }

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
		n.AsmStatement.check(c)
	case ExternalDeclarationEmpty: // ';'
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *AsmStatement) check(c *ctx) {
	n.Asm.check(c)
}

func (n *Asm) check(c *ctx) {
	n.AsmQualifierList.check(c)
	n.AsmArgList.check(c)
}

func (n *AsmArgList) check(c *ctx) {
	for ; n != nil; n = n.AsmArgList {
		n.AsmExpressionList.check(c)
	}
}

func (n *AsmExpressionList) check(c *ctx) {
	for ; n != nil; n = n.AsmExpressionList {
		n.AsmIndex.check(c)
		n.AssignmentExpression.check(c)
	}
}

func (n *AsmIndex) check(c *ctx) {
	n.Expression.check(c)
}

func (n *AsmQualifierList) check(c *ctx) {
	for ; n != nil; n = n.AsmQualifierList {
		n.AsmQualifier.check(c)
	}
}

func (n *AsmQualifier) check(c *ctx) {
	switch n.Case {
	case AsmQualifierVolatile: // "volatile"
		c.errors.add(errorf("TODO %v", n.Case))
	case AsmQualifierInline: // "inline"
		c.errors.add(errorf("TODO %v", n.Case))
	case AsmQualifierGoto: // "goto"
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *FunctionDefinition) check(c *ctx) {
	d := n.Declarator
	d.check(c, n.DeclarationSpecifiers.check(c, &d.isExtern, &d.isStatic, &d.isAtomic, &d.isThreadLocal, &d.isConst, &d.isVolatile, &d.isInline, &d.isRegister))
	n.DeclarationList.check(c)
	n.CompoundStatement.check(c)
}

func (n *DeclarationList) check(c *ctx) {
	for ; n != nil; n = n.DeclarationList {
		n.Declaration.check(c)
	}
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
	var isExtern, isStatic, isAtomic, isThreadLocal, isConst, isVolatile, isInline, isRegister bool
	t := n.DeclarationSpecifiers.check(c, &isExtern, &isStatic, &isAtomic, &isThreadLocal, &isConst, &isVolatile, &isInline, &isRegister)
	for l := n.InitDeclaratorList; l != nil; l = l.InitDeclaratorList {
		l.InitDeclarator.check(c, t, isExtern, isStatic, isAtomic, isThreadLocal, isConst, isVolatile, isInline, isRegister)
	}
}

func (n *InitDeclarator) check(c *ctx, t Type, isExtern, isStatic, isAtomic, isThreadLocal, isConst, isVolatile, isInline, isRegister bool) {
	n.Declarator.isExtern = isExtern
	n.Declarator.isStatic = isStatic
	n.Declarator.isAtomic = isAtomic
	n.Declarator.isThreadLocal = isThreadLocal
	n.Declarator.isConst = isConst
	n.Declarator.isVolatile = isVolatile
	n.Declarator.isInline = isInline
	n.Declarator.isRegister = isRegister
	t = n.Declarator.check(c, t)
	if n.Asm != nil {
		n.Asm.check(c)
		return
	}

	n.Initializer.check(c, t)
}

func (n *Initializer) check(c *ctx, t Type) {
	if n == nil {
		return
	}

	switch n.Case {
	case InitializerExpr: // AssignmentExpression
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	case InitializerInitList: // '{' InitializerList ',' '}'
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *Declarator) check(c *ctx, t Type) (r Type) {
	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	n.typ = n.DirectDeclarator.check(c, n.Pointer.check(c, t))
	return n.typ
}

func (n *DirectDeclarator) check(c *ctx, t Type) (r Type) {
	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case DirectDeclaratorIdent: // IDENTIFIER
		return t
	case DirectDeclaratorDecl: // '(' Declarator ')'
		return n.Declarator.check(c, t)
	case DirectDeclaratorArr: // DirectDeclarator '[' TypeQualifiers AssignmentExpression ']'
		n.AssignmentExpression.check(c)
		return n.DirectDeclarator.check(c, newArrayType(n.AssignmentExpression, t))
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
		return n.DirectDeclarator.check(c, newFunctionType(c, t, nil, false))
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
	var isExtern, isStatic, isAtomic, isThreadLocal, isConst, isVolatile, isInline, isRegister bool
	switch n.Case {
	case ParameterDeclarationDecl: // DeclarationSpecifiers Declarator
		n.Declarator.isParam = true
		n.typ = n.Declarator.check(c, n.DeclarationSpecifiers.check(c, &isExtern, &isStatic, &isAtomic, &isThreadLocal, &isConst, &isVolatile, &isInline, &isRegister))
		n.Declarator.isConst = isConst
		n.Declarator.isVolatile = isVolatile
		n.Declarator.isRegister = isRegister
		if isExtern || isStatic || isAtomic || isThreadLocal || isInline {
			c.errors.add(errorf("%v: storage class or atomic specified or function specifier for parameter: abc", n.Declarator.Position()))
		}
	case ParameterDeclarationAbstract: // DeclarationSpecifiers AbstractDeclarator
		n.typ = n.AbstractDeclarator.check(c, n.DeclarationSpecifiers.check(c, &isExtern, &isStatic, &isAtomic, &isThreadLocal, &isConst, &isVolatile, &isInline, &isRegister))
		if isExtern || isStatic || isAtomic || isThreadLocal || isInline {
			c.errors.add(errorf("%v: storage class or atomic or function specifier for unnamed parameter", n.Position()))
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *AbstractDeclarator) check(c *ctx, t Type) (r Type) {
	if n == nil {
		return t
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	n.typ = n.DirectAbstractDeclarator.check(c, n.Pointer.check(c, t))
	return n.typ
}

func (n *DirectAbstractDeclarator) check(c *ctx, t Type) (r Type) {
	if n == nil {
		return t
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case DirectAbstractDeclaratorDecl: // '(' AbstractDeclarator ')'
		return n.AbstractDeclarator.check(c, t)
	case DirectAbstractDeclaratorArr: // DirectAbstractDeclarator '[' TypeQualifiers AssignmentExpression ']'
		n.AssignmentExpression.check(c)
		return n.DirectAbstractDeclarator.check(c, newArrayType(n.AssignmentExpression, t))
	case DirectAbstractDeclaratorStaticArr: // DirectAbstractDeclarator '[' "static" TypeQualifiers AssignmentExpression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectAbstractDeclaratorArrStatic: // DirectAbstractDeclarator '[' TypeQualifiers "static" AssignmentExpression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectAbstractDeclaratorArrStar: // DirectAbstractDeclarator '[' '*' ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectAbstractDeclaratorFunc: // DirectAbstractDeclarator '(' ParameterTypeList ')'
		fp, isVariadic := n.ParameterTypeList.check(c)
		return n.DirectAbstractDeclarator.check(c, newFunctionType(c, t, fp, isVariadic))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return t //TODO-
}

func (n *Pointer) check(c *ctx, t Type) (r Type) {
	if n == nil {
		return t
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

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

func (n *DeclarationSpecifiers) check(c *ctx, isExtern, isStatic, isAtomic, isThreadLocal, isConst, isVolatile, isInline, isRegister *bool) (r Type) {
	if n == nil {
		return c.intType
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	n0 := n
	var ts []TypeSpecifierCase
	var firstTypeSpecifier *TypeSpecifier
	for ; n != nil; n = n.DeclarationSpecifiers {
		switch n.Case {
		case DeclarationSpecifiersStorage: // StorageClassSpecifier DeclarationSpecifiers
			n.StorageClassSpecifier.check(c, isExtern, isStatic, isThreadLocal, isRegister)
		case DeclarationSpecifiersTypeSpec: // TypeSpecifier DeclarationSpecifiers
			ts = append(ts, n.TypeSpecifier.Case)
			if firstTypeSpecifier == nil {
				firstTypeSpecifier = n.TypeSpecifier
			}
			r = n.TypeSpecifier.check(c, isAtomic)
		case DeclarationSpecifiersTypeQual: // TypeQualifier DeclarationSpecifiers
			n.TypeQualifier.check(c, isConst, isVolatile, isAtomic)
		case DeclarationSpecifiersFunc: // FunctionSpecifier DeclarationSpecifiers
			n.FunctionSpecifier.check(c, isInline)
		case DeclarationSpecifiersAlignSpec: // AlignmentSpecifier DeclarationSpecifiers
			n.AlignmentSpecifier.check(c)
		case DeclarationSpecifiersAttr:
			n.AttributeSpecifierList.check(c)
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	t, ok := c.builtinTypes[ts2String(ts)]
	if ok {
		return t
	}

	switch len(ts) {
	case 0:
		return c.intType
	case 1:
		switch ts[0] {
		case
			TypeSpecifierAtomic,
			TypeSpecifierEnum,
			TypeSpecifierStructOrUnion,
			TypeSpecifierTypeName,
			TypeSpecifierTypeofExpr:

			return r
		}
	}

	c.errors.add(errorf("TODO %v %v", ts, n0.Position()))
	return nil
}

func (n *AttributeSpecifierList) check(c *ctx) {
	for ; n != nil; n = n.AttributeSpecifierList {
		n.AttributeSpecifier.check(c)
	}
}

func (n *AttributeSpecifier) check(c *ctx) {
	n.AttributeValueList.check(c)
}

func (n *AttributeValueList) check(c *ctx) {
	for ; n != nil; n = n.AttributeValueList {
		n.AttributeValue.check(c)
	}
}

func (n *AttributeValue) check(c *ctx) {
	switch n.Case {
	case AttributeValueIdent: // IDENTIFIER
		// ok
	case AttributeValueExpr: // IDENTIFIER '(' ArgumentExpressionList ')'
		n.ArgumentExpressionList.check(c)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *ArgumentExpressionList) check(c *ctx) {
	for ; n != nil; n = n.ArgumentExpressionList {
		n.AssignmentExpression.check(c)
	}
}

func (n *AlignmentSpecifier) check(c *ctx) {
	switch n.Case {
	case AlignmentSpecifierType: // "_Alignas" '(' TypeName ')'
		c.errors.add(errorf("TODO %v", n.Case))
	case AlignmentSpecifierExpr: // "_Alignas" '(' ConstantExpression ')'
		n.ConstantExpression.check(c)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *FunctionSpecifier) check(c *ctx, isInline *bool) {
	switch n.Case {
	case FunctionSpecifierInline: // "inline"
		*isInline = true
	case FunctionSpecifierNoreturn: // "_Noreturn"
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *TypeQualifier) check(c *ctx, isConst, isVolatile, isAtomic *bool) {
	switch n.Case {
	case TypeQualifierConst: // "const"
		*isConst = true
	case TypeQualifierRestrict: // "restrict"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeQualifierVolatile: // "volatile"
		*isVolatile = true
	case TypeQualifierAtomic: // "_Atomic"
		*isAtomic = true
	case TypeQualifierNonnull: // "_Nonnull"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeQualifierAttr: // AttributeSpecifierList
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *StorageClassSpecifier) check(c *ctx, isExtern, isStatic, isThreadLocal, isRegister *bool) {
	switch n.Case {
	case StorageClassSpecifierTypedef: // "typedef"
		// ok
	case StorageClassSpecifierExtern: // "extern"
		*isExtern = true
	case StorageClassSpecifierStatic: // "static"
		*isStatic = true
	case StorageClassSpecifierAuto: // "auto"
		c.errors.add(errorf("TODO %v", n.Case))
	case StorageClassSpecifierRegister: // "register"
		*isRegister = true
	case StorageClassSpecifierThreadLocal: // "_Thread_local"
		*isThreadLocal = true
	case StorageClassSpecifierDeclspec: // "__declspec" '(' ')'
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *TypeSpecifier) check(c *ctx, isAtomic *bool) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case TypeSpecifierVoid: // "void"
		// ok
	case TypeSpecifierChar: // "char"
		// ok
	case TypeSpecifierShort: // "short"
		// ok
	case TypeSpecifierInt: // "int"
		// ok
	case TypeSpecifierInt128: // "__int128"
		// ok
	case TypeSpecifierUint128: // "__uint128_t"
		// ok
	case TypeSpecifierLong: // "long"
		// ok
	case TypeSpecifierFloat: // "float"
		// ok
	case TypeSpecifierFloat16: // "_Float16"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierDecimal64: // "_Decimal64"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierFloat128: // "_Float128"
		// ok
	case TypeSpecifierFloat128x: // "_Float128x"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierDouble: // "double"
		// ok
	case TypeSpecifierSigned: // "signed"
		// ok
	case TypeSpecifierUnsigned: // "unsigned"
		// ok
	case TypeSpecifierBool: // "_Bool"
		// ok
	case TypeSpecifierComplex: // "_Complex"
		// ok
	case TypeSpecifierImaginary: // "_Imaginary"
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierStructOrUnion: // StructOrUnionSpecifier
		return n.StructOrUnionSpecifier.check(c)
	case TypeSpecifierEnum: // EnumSpecifier
		return n.EnumSpecifier.check(c)
	case TypeSpecifierTypeName: // TYPENAME
		if x, ok := n.resolutionScope.ident(n.Token).(*Declarator); ok && x.isTypename {
			return x.typ
		}

		c.errors.add(errorf("%v: undefined type name", n.Position()))
	case TypeSpecifierTypeofExpr: // "typeof" '(' Expression ')'
		return n.Expression.check(c)
	case TypeSpecifierTypeofType: // "typeof" '(' TypeName ')'
		c.errors.add(errorf("TODO %v", n.Case))
	case TypeSpecifierAtomic: // AtomicTypeSpecifier
		*isAtomic = true
	case TypeSpecifierFloat32: // "_Float32"
		// ok
	case TypeSpecifierFloat64: // "_Float64"
		// ok
	case TypeSpecifierFloat32x: // "_Float32x"
		// ok
	case TypeSpecifierFloat64x: // "_Float64x"
		// ok
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return nil
}

func (n *EnumSpecifier) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case EnumSpecifierDef: // "enum" IDENTIFIER '{' EnumeratorList ',' '}'
		n.typ = n.EnumeratorList.check(c)
	case EnumSpecifierTag: // "enum" IDENTIFIER
		// No error is reported when the type is not found. It's sometimes valid to
		// have an enum type that never becomes complete:
		//
		// enum E *e;
		if x := n.resolutionScope.enum(n.Token2); x != nil {
			n.typ = x.typ
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *EnumeratorList) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	return nil
}

func (n *StructOrUnionSpecifier) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	switch n.Case {
	case StructOrUnionSpecifierDef: // StructOrUnion IDENTIFIER '{' StructDeclarationList '}'
		// defer func() {
		// 	if r == nil {
		// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
		// 	}
		// }()

		n.typ = n.StructDeclarationList.check(c)
	case StructOrUnionSpecifierTag: // StructOrUnion IDENTIFIER
		// No error is reported when the type is not found. It's sometimes valid to
		// have a struct type that never becomes complete:
		//
		// struct x *p;
		if x := n.resolutionScope.structOrUnion(n.Token); x != nil {
			if n.StructOrUnion.Case != x.StructOrUnion.Case {
				c.errors.add(errorf("%v: mismateched struct/union tag", n.Token.Position()))
				break
			}

			n.typ = x.typ
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *StructDeclarationList) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	return nil
}

func (n *AssignmentExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case AssignmentExpressionCond: // ConditionalExpression
		n.typ = n.ConditionalExpression.check(c)
	case AssignmentExpressionAssign: // UnaryExpression '=' AssignmentExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case AssignmentExpressionMul: // UnaryExpression "*=" AssignmentExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case AssignmentExpressionDiv: // UnaryExpression "/=" AssignmentExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case AssignmentExpressionMod: // UnaryExpression "%=" AssignmentExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case AssignmentExpressionAdd: // UnaryExpression "+=" AssignmentExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case AssignmentExpressionSub: // UnaryExpression "-=" AssignmentExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case AssignmentExpressionLsh: // UnaryExpression "<<=" AssignmentExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case AssignmentExpressionRsh: // UnaryExpression ">>=" AssignmentExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case AssignmentExpressionAnd: // UnaryExpression "&=" AssignmentExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case AssignmentExpressionXor: // UnaryExpression "^=" AssignmentExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case AssignmentExpressionOr: // UnaryExpression "|=" AssignmentExpression
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *ConditionalExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case ConditionalExpressionLOr: // LogicalOrExpression
		n.typ = n.LogicalOrExpression.check(c)
	case ConditionalExpressionCond: // LogicalOrExpression '?' Expression ':' ConditionalExpression
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *LogicalOrExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case LogicalOrExpressionLAnd: // LogicalAndExpression
		n.typ = n.LogicalAndExpression.check(c)
	case LogicalOrExpressionLOr: // LogicalOrExpression "||" LogicalAndExpression
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *LogicalAndExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case LogicalAndExpressionOr: // InclusiveOrExpression
		n.typ = n.InclusiveOrExpression.check(c)
	case LogicalAndExpressionLAnd: // LogicalAndExpression "&&" InclusiveOrExpression
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *InclusiveOrExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case InclusiveOrExpressionXor: // ExclusiveOrExpression
		n.typ = n.ExclusiveOrExpression.check(c)
	case InclusiveOrExpressionOr: // InclusiveOrExpression '|' ExclusiveOrExpression
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *ExclusiveOrExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case ExclusiveOrExpressionAnd: // AndExpression
		n.typ = n.AndExpression.check(c)
	case ExclusiveOrExpressionXor: // ExclusiveOrExpression '^' AndExpression
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *AndExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case AndExpressionEq: // EqualityExpression
		n.typ = n.EqualityExpression.check(c)
	case AndExpressionAnd: // AndExpression '&' EqualityExpression
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *EqualityExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case EqualityExpressionRel: // RelationalExpression
		n.typ = n.RelationalExpression.check(c)
	case EqualityExpressionEq: // EqualityExpression "==" RelationalExpression
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	case EqualityExpressionNeq: // EqualityExpression "!=" RelationalExpression
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *RelationalExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case RelationalExpressionShift: // ShiftExpression
		n.typ = n.ShiftExpression.check(c)
	case RelationalExpressionLt: // RelationalExpression '<' ShiftExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case RelationalExpressionGt: // RelationalExpression '>' ShiftExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case RelationalExpressionLeq: // RelationalExpression "<=" ShiftExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case RelationalExpressionGeq: // RelationalExpression ">=" ShiftExpression
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *ShiftExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case ShiftExpressionAdd: // AdditiveExpression
		n.typ = n.AdditiveExpression.check(c)
	case ShiftExpressionLsh: // ShiftExpression "<<" AdditiveExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case ShiftExpressionRsh: // ShiftExpression ">>" AdditiveExpression
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *AdditiveExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case AdditiveExpressionMul: // MultiplicativeExpression
		n.typ = n.MultiplicativeExpression.check(c)
	case AdditiveExpressionAdd: // AdditiveExpression '+' MultiplicativeExpression
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	case AdditiveExpressionSub: // AdditiveExpression '-' MultiplicativeExpression
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *MultiplicativeExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case MultiplicativeExpressionCast: // CastExpression
		n.typ = n.CastExpression.check(c)
	case MultiplicativeExpressionMul: // MultiplicativeExpression '*' CastExpression
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	case MultiplicativeExpressionDiv: // MultiplicativeExpression '/' CastExpression
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	case MultiplicativeExpressionMod: // MultiplicativeExpression '%' CastExpression
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *CastExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case CastExpressionUnary: // UnaryExpression
		n.typ = n.UnaryExpression.check(c)
	case CastExpressionCast: // '(' TypeName ')' CastExpression
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *UnaryExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case UnaryExpressionPostfix: // PostfixExpression
		n.typ = n.PostfixExpression.check(c)
	case UnaryExpressionInc: // "++" UnaryExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionDec: // "--" UnaryExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionAddrof: // '&' CastExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionDeref: // '*' CastExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionPlus: // '+' CastExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionMinus: // '-' CastExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionCpl: // '~' CastExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionNot: // '!' CastExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionSizeofExpr: // "sizeof" UnaryExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionSizeofType: // "sizeof" '(' TypeName ')'
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionLabelAddr: // "&&" IDENTIFIER
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionAlignofExpr: // "_Alignof" UnaryExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionAlignofType: // "_Alignof" '(' TypeName ')'
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionImag: // "__imag__" UnaryExpression
		c.errors.add(errorf("TODO %v", n.Case))
	case UnaryExpressionReal: // "__real__" UnaryExpression
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *PostfixExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case PostfixExpressionPrimary: // PrimaryExpression
		n.typ = n.PrimaryExpression.check(c)
	case PostfixExpressionIndex: // PostfixExpression '[' Expression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case PostfixExpressionCall: // PostfixExpression '(' ArgumentExpressionList ')'
		n.PostfixExpression.check(c)
		n.ArgumentExpressionList.check(c)
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	case PostfixExpressionSelect: // PostfixExpression '.' IDENTIFIER
		c.errors.add(errorf("TODO %v", n.Case))
	case PostfixExpressionPSelect: // PostfixExpression "->" IDENTIFIER
		c.errors.add(errorf("TODO %v", n.Case))
	case PostfixExpressionInc: // PostfixExpression "++"
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	case PostfixExpressionDec: // PostfixExpression "--"
		c.errors.add(errorf("TODO %v", n.Case))
	case PostfixExpressionComplit: // '(' TypeName ')' '{' InitializerList ',' '}'
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *PrimaryExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	switch n.Case {
	case PrimaryExpressionIdent: // IDENTIFIER
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	case PrimaryExpressionInt: // INTCONST
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	case PrimaryExpressionFloat: // FLOATCONST
		c.errors.add(errorf("TODO %v", n.Case))
	case PrimaryExpressionEnum: // ENUMCONST
		c.errors.add(errorf("TODO %v", n.Case))
	case PrimaryExpressionChar: // CHARCONST
		c.errors.add(errorf("TODO %v", n.Case))
	case PrimaryExpressionLChar: // LONGCHARCONST
		c.errors.add(errorf("TODO %v", n.Case))
	case PrimaryExpressionString: // STRINGLITERAL
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	case PrimaryExpressionLString: // LONGSTRINGLITERAL
		c.errors.add(errorf("TODO %v", n.Case))
	case PrimaryExpressionExpr: // '(' Expression ')'
		n.typ = n.Expression.check(c)
	case PrimaryExpressionStmt: // '(' CompoundStatement ')'
		c.errors.add(errorf("TODO %v", n.Case))
	case PrimaryExpressionGeneric: // GenericSelection
		c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.typ
}

func (n *Expression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	n0 := n
	for ; n != nil; n = n.Expression {
		n0.typ = n.AssignmentExpression.check(c)
	}
	return n0.typ
}

func (n *ConstantExpression) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	// defer func() {
	// 	if r == nil {
	// 		c.errors.add(errorf("TODO missed type check %v:", n.Position()))
	// 	}
	// }()

	n.typ = n.ConditionalExpression.check(c)
	//TODO eval
	return n.typ
}
