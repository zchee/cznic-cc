// Copyright 2022 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"math"
	"math/big"
	"math/bits"
	"sort"
	"strconv"
	"strings"
	"unicode/utf16"

	"modernc.org/mathutil"
)

const (
	// type check
	decay flags = 1 << iota
	asmArgList
	implicitFuncDef
	ignoreUndefined

	// eval
	addrOf
	dontEval // Eg X in _Generic(X, ...)
)

type flags int

func (f flags) add(g flags) flags { return f | g }
func (f flags) del(g flags) flags { return f &^ g }
func (f flags) has(g flags) bool  { return f&g != 0 }

type ExpressionNode interface {
	Node
	Type() Type
	Value() Value
	check(*ctx, flags) Type
	eval(*ctx, flags) Value
}

const longDoublePrec = 256 // mantissa bits

type ctx struct {
	ast          *AST
	builtinTypes map[string]Type
	errors       errors
	fnScope      *Scope
	implicitFunc Type
	int64T       Type
	intT         Type
	pcharT       Type
	ptrDiffT0    Type
	pvoidT       Type
	pwcharT0     Type
	sizeT0       Type
	voidT        Type
	wcharT0      Type
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
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierShort, TypeSpecifierUnsigned}):               c.newPredefinedType(ComplexUShort),
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierShort}):                                      c.newPredefinedType(ComplexShort),
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex, TypeSpecifierUnsigned}):                                   c.newPredefinedType(ComplexUInt),
		ts2String([]TypeSpecifierCase{TypeSpecifierComplex}):                                                          complexdouble,
		ts2String([]TypeSpecifierCase{TypeSpecifierDecimal64}):                                                        c.newPredefinedType(Decimal64),
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
	c.ast.kinds = map[Kind]Type{}
	for _, v := range c.builtinTypes {
		c.ast.kinds[v.Kind()] = v
	}
	c.intT = c.ast.kinds[Int]
	c.int64T = c.ast.kinds[LongLong]
	c.pcharT = c.newPointerType(c.ast.kinds[Char])
	c.voidT = c.ast.kinds[Void]
	c.pvoidT = c.newPointerType(c.voidT)
	c.implicitFunc = c.newPointerType(c.newFunctionType(c.intT, nil, false))
	return c
}

func (c *ctx) convert(v Value, t Type) (r Value) {
	if v == nil || v == Unknown {
		return Unknown
	}

	switch t.Kind() {
	case Int, Long, LongLong, Char, SChar, Short:
		m := Int64Value(1)<<(8*t.Size()) - 1
		switch x := v.(type) {
		case Int64Value:
			if x < 0 {
				return x | ^m
			}

			return x & m
		case UInt64Value:
			y := Int64Value(x)
			if y < 0 {
				return y | ^m
			}

			return y & m
		}
	case ULong, UInt, ULongLong, UChar, UShort:
		m := UInt64Value(1)<<(8*t.Size()) - 1
		switch x := v.(type) {
		case Int64Value:
			return UInt64Value(x) & m
		case UInt64Value:
			return x & m
		}
	case Bool:
		switch x := v.(type) {
		case Int64Value:
			if x != 0 {
				return int1
			}

			return int0
		case UInt64Value:
			if x != 0 {
				return int1
			}

			return int0
		}
	case Ptr:
		switch x := v.(type) {
		case Int64Value:
			return UInt64Value(x)
		case UInt64Value:
			return x
		}
	case Void:
		return VoidValue{}
	}
	return Unknown
}

func (c *ctx) decay(t Type, mode flags) Type {
	if !mode.has(decay) || t == nil {
		return t
	}

	return t.Decay()
}

func (c *ctx) wcharT(n Node) Type {
	if c.wcharT0 == nil {
		if s := c.ast.scope.Nodes["wchar_t"]; len(s) != 0 {
			if d, ok := s[0].(*Declarator); ok && d.isTypename {
				c.wcharT0 = d.Type()
			}
		}
		if c.wcharT0 == nil {
			c.errors.add(errorf("%v: undefined type: wchar_t", n.Position()))
			c.wcharT0 = c.intT
		}
	}
	return c.wcharT0
}

func (c *ctx) ptrDiffT(n Node) Type {
	if c.ptrDiffT0 == nil {
		if s := c.ast.scope.Nodes["ptrdiff_t"]; len(s) != 0 {
			if d, ok := s[0].(*Declarator); ok && d.isTypename {
				c.ptrDiffT0 = d.Type()
			}
		}
		if c.ptrDiffT0 == nil {
			c.errors.add(errorf("%v: undefined type: ptrdiff_t", n.Position()))
			c.ptrDiffT0 = c.intT
		}
	}
	return c.ptrDiffT0
}

func (c *ctx) sizeT(n Node) Type {
	if c.sizeT0 == nil {
		if s := c.ast.scope.Nodes["size_t"]; len(s) != 0 {
			if d, ok := s[0].(*Declarator); ok && d.isTypename {
				c.sizeT0 = d.Type()
			}
		}
		if c.sizeT0 == nil {
			c.errors.add(errorf("%v: undefined type: size_t", n.Position()))
			c.sizeT0 = c.intT
		}
	}
	return c.sizeT0
}

func (c *ctx) pwcharT(n Node) Type {
	if c.pwcharT0 == nil {
		c.pwcharT0 = c.newPointerType(c.wcharT(n))
	}
	return c.pwcharT0
}

type typer struct{ typ Type }

func newTyper(t Type) typer { return typer{typ: t} }

// Type returns the type of a node or an *InvalidType type value, if the type
// is unknown/undetermined.
func (t typer) Type() Type {
	if t.typ != nil {
		return t.typ
	}

	return Invalid
}

type valuer struct{ val Value }

// Value returns the value of a node or UnknownValue if it is undetermined. The
// dynamic type of a Value is one of
//
//	*ComplexLongDoubleValue
//	*LongDoubleValue
//	*UnknownValue
//	*ZeroValue
//	Complex128Value
//	Complex64Value
//	Float64Value
//	Int64Value
//	StringValue
//	UInt64Value
//	UTF16StringValue
//	UTF32StringValue
//	VoidValue
func (v valuer) Value() Value {
	if v.val != nil {
		return v.val
	}

	return Unknown
}

type AST struct {
	ABI             *ABI
	EOF             Token
	TranslationUnit *TranslationUnit
	kinds           map[Kind]Type
	scope           *Scope
}

func (n *AST) check() error {
	c := newCtx(n)
	for l := n.TranslationUnit; l != nil; l = l.TranslationUnit {
		l.ExternalDeclaration.check(c)
	}
	return c.errors.err()
}

//  ExternalDeclaration:
//          FunctionDefinition  // Case ExternalDeclarationFuncDef
//  |       Declaration         // Case ExternalDeclarationDecl
//  |       AsmStatement        // Case ExternalDeclarationAsmStmt
//  |       ';'                 // Case ExternalDeclarationEmpty
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

//  AsmStatement:
//          Asm ';'
func (n *AsmStatement) check(c *ctx) {
	n.Asm.check(c)
}

//  Asm:
//          "asm" AsmQualifierList '(' STRINGLITERAL AsmArgList ')'
func (n *Asm) check(c *ctx) {
	n.AsmQualifierList.check(c)
	n.AsmArgList.check(c)
}

//  AsmArgList:
//          ':' AsmExpressionList
//  |       AsmArgList ':' AsmExpressionList
func (n *AsmArgList) check(c *ctx) {
	for ; n != nil; n = n.AsmArgList {
		n.AsmExpressionList.check(c)
	}
}

//  AsmExpressionList:
//          AsmIndex AssignmentExpression
//  |       AsmExpressionList ',' AsmIndex AssignmentExpression
func (n *AsmExpressionList) check(c *ctx) {
	for ; n != nil; n = n.AsmExpressionList {
		n.AsmIndex.check(c)
		n.AssignmentExpression.check(c, decay|asmArgList|ignoreUndefined)
	}
}

//  AsmIndex:
//          '[' ExpressionList ']'
func (n *AsmIndex) check(c *ctx) {
	if n == nil {
		return
	}

	n.ExpressionList.check(c, decay|asmArgList)
}

func (n *AsmQualifierList) check(c *ctx) {
	for ; n != nil; n = n.AsmQualifierList {
		n.AsmQualifier.check(c)
	}
}

//  AsmQualifierList:
//          AsmQualifier
//  |       AsmQualifierList AsmQualifier
func (n *AsmQualifier) check(c *ctx) {
	switch n.Case {
	case AsmQualifierVolatile: // "volatile"
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	case AsmQualifierInline: // "inline"
		c.errors.add(errorf("TODO %v", n.Case))
	case AsmQualifierGoto: // "goto"
		//TODO c.errors.add(errorf("TODO %v", n.Case))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

//  FunctionDefinition:
//          DeclarationSpecifiers Declarator DeclarationList CompoundStatement
func (n *FunctionDefinition) check(c *ctx) {
	d := n.Declarator
	d.check(c, n.DeclarationSpecifiers.check(c, &d.isExtern, &d.isStatic, &d.isAtomic, &d.isThreadLocal, &d.isConst, &d.isVolatile, &d.isInline, &d.isRegister, &d.isAuto, &d.isNoreturn))
	switch d.DirectDeclarator.Case {
	case DirectDeclaratorFuncIdent:
		ft, ok := d.Type().(*FunctionType)
		if !ok {
			break
		}

		m := n.DeclarationList.check(c)
		for _, param := range d.DirectDeclarator.IdentifierList.parameters {
			switch d := m[param.name.SrcStr()]; {
			case d == nil:
				param.typ = c.intT
			default:
				param.Declarator = d
				param.typ = d.Type()
			}
			ft.fp = append(ft.fp, param)
		}
	}
	c.fnScope = n.scope
	defer func() { c.fnScope = nil }()
	n.CompoundStatement.check(c)
}

//  DeclarationList:
//          Declaration
//  |       DeclarationList Declaration
func (n *DeclarationList) check(c *ctx) (m map[string]*Declarator) {
	for ; n != nil; n = n.DeclarationList {
		n.Declaration.check(c)
		if m == nil {
			m = map[string]*Declarator{}
		}
		for l := n.Declaration.InitDeclaratorList; l != nil; l = l.InitDeclaratorList {
			d := l.InitDeclarator.Declarator
			nm := d.Name()
			if x := m[nm]; x != nil {
				c.errors.add(errorf("%v: %s redeclared, previous declaration at %v:", d.Position(), nm, x.Position()))
				continue
			}

			m[nm] = d
		}
	}
	return m
}

//  CompoundStatement:
//          '{' LabelDeclarationList BlockItemList '}'
func (n *CompoundStatement) check(c *ctx) (r Type) {
	if n == nil {
		return
	}

	r = c.voidT
	for l := n.BlockItemList; l != nil; l = l.BlockItemList {
		r = l.BlockItem.check(c)
	}
	return r
}

func (n *BlockItem) check(c *ctx) (r Type) {
	if n == nil {
		return Invalid
	}

	switch n.Case {
	case BlockItemDecl: // Declaration
		n.Declaration.check(c)
	case BlockItemLabel:
		n.LabelDeclaration.check(c)
	case BlockItemStmt: // Statement
		return n.Statement.check(c)
	case BlockItemFuncDef: // DeclarationSpecifiers Declarator CompoundStatement
		var isExtern, isStatic, isAtomic, isThreadLocal, isConst, isVolatile, isInline, isRegister, isAuto, isNoreturn bool
		n.Declarator.check(c, n.DeclarationSpecifiers.check(c, &isExtern, &isStatic, &isAtomic, &isThreadLocal, &isConst, &isVolatile, &isInline, &isRegister, &isAuto, &isNoreturn))
		if isExtern || isStatic || isAtomic || isThreadLocal || isConst || isVolatile || isRegister || isAuto {
			c.errors.add(errorf("%v: invalid specifier/qualifer combination", n.Position()))
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return nil
}

func (n *LabelDeclaration) check(c *ctx) {
	//TODO
}

func (n *Statement) check(c *ctx) (r Type) {
	if n == nil {
		return Invalid
	}

	switch n.Case {
	case StatementLabeled: // LabeledStatement
		return n.LabeledStatement.check(c)
	case StatementCompound: // CompoundStatement
		return n.CompoundStatement.check(c)
	case StatementExpr: // ExpressionStatement
		return n.ExpressionStatement.check(c)
	case StatementSelection: // SelectionStatement
		return n.SelectionStatement.check(c)
	case StatementIteration: // IterationStatement
		return n.IterationStatement.check(c)
	case StatementJump: // JumpStatement
		return n.JumpStatement.check(c)
	case StatementAsm: // AsmStatement
		n.AsmStatement.check(c)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return Invalid
}

func (n *LabeledStatement) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	switch n.Case {
	case LabeledStatementLabel: // IDENTIFIER ':' Statement
		return n.Statement.check(c)
	case LabeledStatementCaseLabel: // "case" ConstantExpression ':' Statement
		n.ConstantExpression.check(c, decay)
		return n.Statement.check(c)
	case LabeledStatementRange: // "case" ConstantExpression "..." ConstantExpression ':' Statement
		n.ConstantExpression.check(c, decay)
		n.ConstantExpression2.check(c, decay)
		return n.Statement.check(c)
	case LabeledStatementDefault: // "default" ':' Statement
		return n.Statement.check(c)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return nil
}

func (n *IterationStatement) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	switch n.Case {
	case IterationStatementWhile: // "while" '(' ExpressionList ')' Statement
		n.ExpressionList.check(c, decay)
		return n.Statement.check(c)
	case IterationStatementDo: // "do" Statement "while" '(' ExpressionList ')' ';'
		r = n.Statement.check(c)
		n.ExpressionList.check(c, decay)
		return r
	case IterationStatementFor: // "for" '(' ExpressionList ';' ExpressionList ';' ExpressionList ')' Statement
		n.ExpressionList.check(c, decay)
		n.ExpressionList2.check(c, decay)
		n.ExpressionList3.check(c, decay)
		return n.Statement.check(c)
	case IterationStatementForDecl: // "for" '(' Declaration ExpressionList ';' ExpressionList ')' Statement
		n.Declaration.check(c)
		n.ExpressionList.check(c, decay)
		n.ExpressionList2.check(c, decay)
		return n.Statement.check(c)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return nil
}

func (n *JumpStatement) check(c *ctx) (r Type) {
	if n == nil {
		return
	}

out:
	switch n.Case {
	case JumpStatementGoto: // "goto" IDENTIFIER ';'
		for _, nd := range c.fnScope.Nodes[string(n.Token2.Src())] {
			switch nd.(type) {
			case *LabeledStatement:
				break out
			}
		}

		c.errors.add(errorf("%v: undefined label: %s", n.Token2.Position(), n.Token2.Src()))
	case JumpStatementGotoExpr: // "goto" '*' ExpressionList ';'
		n.ExpressionList.check(c, decay)
	case JumpStatementContinue: // "continue" ';'
		//TODO
	case JumpStatementBreak: // "break" ';'
		//TODO
	case JumpStatementReturn: // "return" ExpressionList ';'
		r = n.ExpressionList.check(c, decay)
		//TODO check assignable to fn result
		return r
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return c.voidT
}

func (n *SelectionStatement) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	switch n.Case {
	case SelectionStatementIf: // "if" '(' ExpressionList ')' Statement
		n.ExpressionList.check(c, decay)
		return n.Statement.check(c)
	case SelectionStatementIfElse: // "if" '(' ExpressionList ')' Statement "else" Statement
		n.ExpressionList.check(c, decay)
		r1 := n.Statement.check(c)
		r2 := n.Statement2.check(c)
		if r1 != nil && r1 != Invalid {
			return r1
		}

		return r2
	case SelectionStatementSwitch: // "switch" '(' ExpressionList ')' Statement
		n.ExpressionList.check(c, decay)
		return n.Statement.check(c)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return nil
}

func (n *ExpressionStatement) check(c *ctx) (r Type) {
	if n == nil {
		return nil
	}

	if n.ExpressionList == nil {
		return c.voidT
	}

	return n.ExpressionList.check(c, decay)
}

//  Declaration:
//          DeclarationSpecifiers InitDeclaratorList AttributeSpecifierList ';'
func (n *Declaration) check(c *ctx) {
	var isExtern, isStatic, isAtomic, isThreadLocal, isConst, isVolatile, isInline, isRegister, isAuto, isNoreturn bool
	t := n.DeclarationSpecifiers.check(c, &isExtern, &isStatic, &isAtomic, &isThreadLocal, &isConst, &isVolatile, &isInline, &isRegister, &isAuto, &isNoreturn)
	var attr *Attributes
	if n.InitDeclaratorList != nil && n.InitDeclaratorList.InitDeclaratorList == nil {
		if attr = n.InitDeclaratorList.InitDeclarator.AttributeSpecifierList.check(c); attr != nil {
			t = t.setAttr(attr)
		}
	}
	for l := n.InitDeclaratorList; l != nil; l = l.InitDeclaratorList {
		l.InitDeclarator.check(c, t, isExtern, isStatic, isAtomic, isThreadLocal, isConst, isVolatile, isInline, isRegister, isAuto)
	}
}

//  InitDeclarator:
//          Declarator Asm                  // Case InitDeclaratorDecl
//  |       Declarator Asm '=' Initializer  // Case InitDeclaratorInit
func (n *InitDeclarator) check(c *ctx, t Type, isExtern, isStatic, isAtomic, isThreadLocal, isConst, isVolatile, isInline, isRegister, isAuto bool) {
	if t == nil {
		c.errors.add(errorf("TODO %T internal error", n))
		return
	}

	if n == nil {
		return
	}

	n.Declarator.isAtomic = isAtomic
	n.Declarator.isAuto = isAuto
	n.Declarator.isConst = isConst
	n.Declarator.isExtern = isExtern
	n.Declarator.isInline = isInline
	n.Declarator.isRegister = isRegister
	n.Declarator.isStatic = isStatic
	n.Declarator.isThreadLocal = isThreadLocal
	n.Declarator.isVolatile = isVolatile
	t = n.Declarator.check(c, t)
	if n.Asm != nil {
		n.Asm.check(c)
		return
	}

	if n.Case == InitDeclaratorInit {
		n.Initializer.check(c, t, 0)
	}
}

//  Initializer:
//          AssignmentExpression         // Case InitializerExpr
//  |       '{' InitializerList ',' '}'  // Case InitializerInitList
func (n *Initializer) check(c *ctx, t Type, off int64) {
	if n == nil || t == nil {
		c.errors.add(errorf("internal error %T(%v) %T(%v)", n, n == nil, t, t == nil))
		return
	}

	// The type of the entity to be initialized shall be an array of unknown size
	// or an object type that is not a variable length array type.
	n.typ = t
	n.off = off
	t = n.Type()
	if x, ok := t.(*ArrayType); ok && x.IsVLA() {
		c.errors.add(errorf("%v: cannot initalize a variable length array", n.Position()))
		return
	}

	switch n.Case {
	case InitializerExpr: // AssignmentExpression
		exprT := n.AssignmentExpression.check(c, decay)
		if exprT == Invalid {
			n.val = Unknown
			c.errors.add(errorf("TODO %T <- %T", t, exprT))
			return
		}

		n.val = n.AssignmentExpression.Value()
		switch x := t.(type) {
		case *ArrayType:
			n.checkExprArray(c, x, exprT, off)
		case *EnumType:
			n.val = c.convert(n.Value(), t)
		case *PointerType:
			if isPointerType(exprT) {
				return
			}

			if isIntegerType(exprT) {
				n.val = c.convert(n.Value(), t)
				return
			}

			c.errors.add(errorf("TODO %T <- %T %T", x, n.Value(), exprT))
		case *PredefinedType:
			n.val = c.convert(n.Value(), t)
		case *StructType:
			n.checkExprStruct(c, x, exprT, off)
		case *UnionType:
			n.checkExprUnion(c, x, exprT, off)
		default:
			c.errors.add(errorf("TODO %T <- %T %T", x, n.Value(), exprT))
		}
	case InitializerInitList: // '{' InitializerList ',' '}'
		if n.InitializerList == nil {
			n.val = Zero
			return
		}

		n.InitializerList.check(c, t, off, true)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

func (n *Initializer) checkExprArray(c *ctx, t *ArrayType, exprT Type, off int64) {
	if n.Case != InitializerExpr {
		c.errors.add(errorf("internal error: %T %v", n, n.Case))
		return
	}

	v := n.Value()
	switch x := t.Elem().(type) {
	case *ArrayType:
		n.check(c, x.Elem(), off)
		if t.IsIncomplete() {
			t.elems = 1
		}
	case
		*EnumType,
		*PointerType,
		*StructType,
		*UnionType:

		n.check(c, x, off)
		if t.IsIncomplete() {
			t.elems = 1
		}
	case *PredefinedType:
		// [0]6.7.8/14 An array of character type may be initialized by a character
		// string literal, optionally enclosed in braces. Successive characters of the
		// character string literal (including the terminating null character if there
		// is room or if the array is of unknown size) initialize the elements of the
		// array.
		switch x.Kind() {
		case Char, SChar, UChar:
			if x, ok := v.(StringValue); ok {
				if t.IsIncomplete() {
					t.elems = int64(len(x))
				}
				if max, sz := t.Len(), int64(len(x)); sz > max {
					n.val = x[:max]
				}
				return
			}
		}

		// [0]6.7.8/15 An array with element type compatible with wchar_t may be
		// initialized by a wide string literal, optionally enclosed in braces.
		// Successive wide characters of the wide string literal (including the
		// terminating null wide character if there is room or if the array is of
		// unknown size) initialize the elements of the array.
		if isIntegerType(x) && x.Size() == c.wcharT(n).Size() {
			switch x := v.(type) {
			case UTF16StringValue:
				if t.IsIncomplete() {
					t.elems = int64(len(x))
				}
				if max, sz := t.Len(), int64(len(x)); sz > max {
					n.val = x[:max]
				}
				return
			case UTF32StringValue:
				if t.IsIncomplete() {
					t.elems = int64(len(x))
				}
				if max, sz := t.Len(), int64(len(x)); sz > max {
					n.val = x[:max]
				}
				return
			}
		}

		n.check(c, x, off)
		if t.IsIncomplete() {
			t.elems = 1
		}
	default:
		c.errors.add(errorf("TODO []%T <- %T %T", x, v, exprT))
	}
}

func (n *Initializer) checkExprStruct(c *ctx, t *StructType, exprT Type, off int64) {
	if n.Case != InitializerExpr {
		c.errors.add(errorf("internal error: %T %v", n, n.Case))
		return
	}

	switch x := exprT.(type) {
	case
		*EnumType,
		*PointerType,
		*PredefinedType:

		f := t.NamedFieldByIndex(0)
		if f == nil {
			c.errors.add(errorf("TODO %T <- %T %T", t, n.Value(), exprT))
			return
		}

		n.check(c, f.Type(), off)
	case *StructType:
		if !t.isCompatible(x) {
			c.errors.add(errorf("%v: incompatible types: %s and %s", n.AssignmentExpression.Position(), t, x))
		}
	default:
		c.errors.add(errorf("TODO %T <- %T %T", t, n.Value(), x))
	}

}

func (n *Initializer) checkExprUnion(c *ctx, t *UnionType, exprT Type, off int64) {
	if n.Case != InitializerExpr {
		c.errors.add(errorf("internal error: %T %v", n, n.Case))
		return
	}

	switch x := exprT.(type) {
	case
		*PointerType,
		*PredefinedType:

		f := t.NamedFieldByIndex(0)
		if f == nil {
			c.errors.add(errorf("TODO %T <- %T %T", t, n.Value(), exprT))
			return
		}

		n.check(c, f.Type(), off)
	case *UnionType:
		if !t.isCompatible(x) {
			c.errors.add(errorf("%v: incompatible types: %s and %s", n.AssignmentExpression.Position(), t, x))
		}
	default:
		c.errors.add(errorf("TODO %T <- %T %T", t, n.Value(), x))
	}
}

func (n *InitializerList) check(c *ctx, currObj Type, off int64, outer bool) *InitializerList {
	if n == nil || currObj == nil {
		c.errors.add(errorf("internal error: %T %T", n, currObj))
		return nil
	}

	switch x := currObj.(type) {
	case *ArrayType:
		return n.checkArray(c, x, off, outer)
	case *EnumType:
		return n.checkEnum(c, x, off, outer)
	case *PointerType:
		return n.checkPointer(c, x, off, outer)
	case *PredefinedType:
		if x.VectorSize() > 0 {
			return n.checkArray(c, x.vector, off, outer)
		}

		return n.checkPredefined(c, x, off, outer)
	case *StructType:
		return n.checkStruct(c, x, off, outer)
	case *UnionType:
		return n.checkUnion(c, x, off, outer)
	default:
		c.errors.add(errorf("TODO %T <- ...", x))
	}
	return nil
}

func (n *InitializerList) checkArray(c *ctx, t *ArrayType, off int64, outer bool) *InitializerList {
	elemT := t.Elem()
	switch {
	case t.IsIncomplete():
		var x, y int64
		for ; n != nil; n = n.InitializerList {
			if n.Designation != nil {
				if !outer {
					return n
				}

				dl := n.Designation.DesignatorList
				x, y = dl.Designator.index(c)
				if x < 0 {
					return nil
				}

				if dl := dl.DesignatorList; dl != nil {
					//TODO if n = n.checkDesignatorList(dl, c, f.Type(), off+f.Offset()); n == nil { return nil }
					c.errors.add(errorf("TODO %T", n))
					return nil
				}

				x = mathutil.MaxInt64(x, y)
			}
			t.elems = mathutil.MaxInt64(t.elems, x)
			n.Initializer.check(c, elemT, off+x*elemT.Size())
		}
		return nil
	default:
		var x, y int64
		for ; n != nil; n = n.InitializerList {
			if n.Designation != nil {
				if !outer {
					return n
				}

				dl := n.Designation.DesignatorList
				x, y = dl.Designator.index(c)
				if x < 0 {
					return nil
				}

				if x >= t.elems {
					c.errors.add(errorf("%v: index %v out of range for array of %v elements", dl.Designator.Position(), x, t.elems))
					return nil
				}

				if y >= t.elems {
					c.errors.add(errorf("%v: index %v out of range for array of %v elements", dl.Designator.Position(), y, t.elems))
					return nil
				}

				if dl := dl.DesignatorList; dl != nil {
					//TODO if n = n.checkDesignatorList(dl, c, f.Type(), off+f.Offset()); n == nil { return nil }
					c.errors.add(errorf("TODO %T", n))
					return nil
				}

				x = mathutil.MaxInt64(x, y)
			}
			if x >= t.elems {
				c.errors.add(errorf("%v: index %v out of range for array of %v elements", x, t.elems))
				return nil
			}

			n.Initializer.check(c, elemT, off+x*elemT.Size())
			x++
			if x == t.elems {
				return n.InitializerList
			}
		}
		return n
	}
}

func (n *Designator) index(c *ctx) (lo, hi int64) {
	switch n.Case {
	case DesignatorIndex: // '[' ConstantExpression ']'
		lo = arrayIndex(c, n.ConstantExpression)
		hi = arrayIndex(c, n.ConstantExpression2)
		if lo < 0 {
			c.errors.add(errorf("%v: array indices cannot be negative"))
			return -1, -1
		}

		return lo, -1
	case DesignatorIndex2: // '[' ConstantExpression "..." ConstantExpression ']'
		lo = arrayIndex(c, n.ConstantExpression)
		hi = arrayIndex(c, n.ConstantExpression2)
		if lo < 0 || hi < 0 {
			c.errors.add(errorf("%v: array indices cannot be negative"))
			return -1, -1
		}

		if lo >= 0 && hi >= 0 && lo >= hi {
			c.errors.add(errorf("%v: first index shall be smaller than second index"))
			return -1, -1
		}

		return lo, hi
	case
		DesignatorField,  // '.' IDENTIFIER
		DesignatorField2: // IDENTIFIER ':'
		c.errors.add(errorf("%v: expected bracketed array index"))
	}
	return -1, -1
}

func arrayIndex(c *ctx, n ExpressionNode) int64 {
	if n == nil {
		return -1
	}

	v, ok := int64Value(c, n)
	if !ok || v < 0 {
		c.errors.add(errorf("%v: invalid array index", n.Position()))
		return -1
	}

	return v
}

func int64Value(c *ctx, n ExpressionNode) (int64, bool) {
	if n == nil {
		return 0, false
	}

	switch t := n.check(c, decay); {
	case isIntegerType(t):
		switch x := n.eval(c, 0).(type) {
		case Int64Value:
			return int64(x), true
		case UInt64Value:
			if x <= math.MaxInt64 {
				return int64(x), true
			}
		}
	}
	return 0, false
}

func (n *InitializerList) checkEnum(c *ctx, t *EnumType, off int64, outer bool) *InitializerList {
	if n.Designation != nil {
		if !outer {
			return n
		}

		c.errors.add(errorf("%v: unexpected designation", n.Designation.Position()))
		return nil
	}

	if isScalarType(t) {
		n.Initializer.check(c, t, off)
		return n.InitializerList
	}

	c.errors.add(errorf("TODO %T <- ...", t))
	return nil
}

func (n *InitializerList) checkPointer(c *ctx, t *PointerType, off int64, outer bool) *InitializerList {
	if n.Designation != nil {
		if !outer {
			return n
		}

		c.errors.add(errorf("%v: unexpected designation", n.Designation.Position()))
		return nil
	}

	if isScalarType(t) {
		n.Initializer.check(c, t, off)
		return n.InitializerList
	}

	c.errors.add(errorf("TODO %T <- ...", t))
	return nil
}

func (n *InitializerList) checkPredefined(c *ctx, t *PredefinedType, off int64, outer bool) *InitializerList {
	if n.Designation != nil {
		if !outer {
			return n
		}

		c.errors.add(errorf("%v: unexpected designation", n.Designation.Position()))
		return nil
	}

	if isScalarType(t) {
		n.Initializer.check(c, t, off)
		return n.InitializerList
	}

	c.errors.add(errorf("TODO %T <- ...", t))
	return nil
}

func (n *InitializerList) checkStruct(c *ctx, t *StructType, off int64, outer bool) *InitializerList {
	f := t.NamedFieldByIndex(0)
	for f != nil && n != nil {
		if n.Designation != nil {
			if !outer {
				return n
			}

			dl := n.Designation.DesignatorList
			nm := dl.Designator.name(c)
			if nm == "" {
				return nil
			}

			if f = t.FieldByName(nm); f == nil {
				c.errors.add(errorf("%v: %v has no member %v", n.Position(), t, nm))
				return nil
			}

			if dl := dl.DesignatorList; dl != nil {
				if n = n.checkDesignatorList(dl, c, f.Type(), off+f.Offset()); n == nil {
					return nil
				}
			}
		}

		if f == nil {
			return nil
		}

		n.Initializer.check(c, f.Type(), off+f.Offset())
		n = n.InitializerList
		f = t.NamedFieldByIndex(f.ordinal + 1)
	}
	return n
}

func (n *InitializerList) checkDesignatorList(dl *DesignatorList, c *ctx, t Type, off int64) *InitializerList {
	for ; dl != nil; dl = dl.DesignatorList {
		switch x := t.(type) {
		case *StructType:
			nm := dl.Designator.name(c)
			if nm == "" {
				return nil
			}

			f := x.FieldByName(nm)
			if f == nil {
				c.errors.add(errorf("%v: type %s has no member %s", dl.Designator.Position(), t, nm))
				return nil
			}

			t = f.Type()
			off += f.Offset()
		default:
			c.errors.add(errorf("TODO %T", x))
			return nil
		}
	}
	n2 := *n
	n2.Designation = nil
	return n2.check(c, t, off, false)
}

func (n *Designator) name(c *ctx) string {
	switch n.Case {
	case
		DesignatorIndex,  // '[' ConstantExpression ']'
		DesignatorIndex2: // '[' ConstantExpression "..." ConstantExpression ']'

		c.errors.add(errorf("%v: expected field name"))
	case DesignatorField: // '.' IDENTIFIER
		return n.Token2.SrcStr()
	case DesignatorField2: // IDENTIFIER ':'
		return n.Token.SrcStr()
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return ""
}

func (n *InitializerList) checkUnion(c *ctx, t *UnionType, off int64, outer bool) *InitializerList {
	f := t.NamedFieldByIndex(0)
	if n.Designation != nil {
		if !outer {
			return n
		}

		dl := n.Designation.DesignatorList
		nm := dl.Designator.name(c)
		if nm == "" {
			return nil
		}

		if f = t.FieldByName(nm); f == nil {
			c.errors.add(errorf("%v: %v has no member %v", n.Position(), t, nm))
			return nil
		}

		if dl := dl.DesignatorList; dl != nil {
			return n.checkDesignatorList(dl, c, f.Type(), off+f.Offset())
		}
	}

	if f == nil {
		return nil
	}

	n.Initializer.check(c, f.Type(), off+f.Offset())
	return n.InitializerList
}

//  Declarator:
//          Pointer DirectDeclarator
func (n *Declarator) check(c *ctx, t Type) (r Type) {
	if t == nil {
		c.errors.add(errorf("TODO %T internal error", n))
		return
	}

	if n == nil {
		return t
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
		}
	}()

	r = n.DirectDeclarator.check(c, n.Pointer.check(c, t))
	if n.isTypename {
		r = r.setName(n.Name())
	}
	n.typ = r
	return n.Type()
}

//  DirectDeclarator:
//          IDENTIFIER                                                             // Case DirectDeclaratorIdent
//  |       '(' Declarator ')'                                                     // Case DirectDeclaratorDecl
//  |       DirectDeclarator '[' TypeQualifiers AssignmentExpression ']'           // Case DirectDeclaratorArr
//  |       DirectDeclarator '[' "static" TypeQualifiers AssignmentExpression ']'  // Case DirectDeclaratorStaticArr
//  |       DirectDeclarator '[' TypeQualifiers "static" AssignmentExpression ']'  // Case DirectDeclaratorArrStatic
//  |       DirectDeclarator '[' TypeQualifiers '*' ']'                            // Case DirectDeclaratorStar
//  |       DirectDeclarator '(' ParameterTypeList ')'                             // Case DirectDeclaratorFuncParam
//  |       DirectDeclarator '(' IdentifierList ')'                                // Case DirectDeclaratorFuncIdent
func (n *DirectDeclarator) check(c *ctx, t Type) (r Type) {
	if t == nil {
		c.errors.add(errorf("TODO %T internal error", n))
		return
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
		}
	}()

	switch n.Case {
	case DirectDeclaratorIdent: // IDENTIFIER
		return t
	case DirectDeclaratorDecl: // '(' Declarator ')'
		return n.Declarator.check(c, t)
	case DirectDeclaratorArr: // DirectDeclarator '[' TypeQualifiers AssignmentExpression ']'
		return n.DirectDeclarator.check(c, c.newArrayType(t, arraySize(c, n.AssignmentExpression), n.AssignmentExpression))
	case DirectDeclaratorStaticArr: // DirectDeclarator '[' "static" TypeQualifiers AssignmentExpression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectDeclaratorArrStatic: // DirectDeclarator '[' TypeQualifiers "static" AssignmentExpression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectDeclaratorStar: // DirectDeclarator '[' TypeQualifiers '*' ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectDeclaratorFuncParam: // DirectDeclarator '(' ParameterTypeList ')'
		fp, isVariadic := n.ParameterTypeList.check(c)
		return n.DirectDeclarator.check(c, c.newFunctionType(t, fp, isVariadic))
	case DirectDeclaratorFuncIdent: // DirectDeclarator '(' IdentifierList ')'
		return n.DirectDeclarator.check(c, c.newFunctionType(t, nil, false))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return t //TODO-
}

func arraySize(c *ctx, n ExpressionNode) int64 {
	if n == nil {
		return -1
	}

	v, ok := int64Value(c, n)
	if !ok { // VLA
		return -1
	}

	if v < 0 {
		c.errors.add(errorf("%v: invalid array size", n.Position()))
		return -1
	}

	return v
}

//  ParameterTypeList:
//          ParameterList            // Case ParameterTypeListList
//  |       ParameterList ',' "..."  // Case ParameterTypeListVar
func (n *ParameterTypeList) check(c *ctx) (fp []*ParameterDeclaration, isVariadic bool) {
	if n == nil {
		return nil, false
	}

	fp = n.ParameterList.check(c)
	return fp, n.Case == ParameterTypeListVar
}

//  ParameterList:
//          ParameterDeclaration
//  |       ParameterList ',' ParameterDeclaration
func (n *ParameterList) check(c *ctx) (fp []*ParameterDeclaration) {
	for ; n != nil; n = n.ParameterList {
		n.ParameterDeclaration.check(c)
		fp = append(fp, n.ParameterDeclaration)
	}
	return fp
}

//  ParameterDeclaration:
//          DeclarationSpecifiers Declarator          // Case ParameterDeclarationDecl
//  |       DeclarationSpecifiers AbstractDeclarator  // Case ParameterDeclarationAbstract
func (n *ParameterDeclaration) check(c *ctx) {
	var isExtern, isStatic, isAtomic, isThreadLocal, isConst, isVolatile, isInline, isRegister, isAuto, isNoreturn bool
	switch n.Case {
	case ParameterDeclarationDecl: // DeclarationSpecifiers Declarator
		n.Declarator.isParam = true
		n.typ = n.Declarator.check(c, n.DeclarationSpecifiers.check(c, &isExtern, &isStatic, &isAtomic, &isThreadLocal, &isConst, &isVolatile, &isInline, &isRegister, &isAuto, &isNoreturn))
		n.Declarator.isConst = isConst
		n.Declarator.isVolatile = isVolatile
		n.Declarator.isRegister = isRegister
		n.Declarator.isNoreturn = isNoreturn
		if isExtern || isStatic || isAtomic || isThreadLocal || isInline || isAuto {
			c.errors.add(errorf("%v: storage class or atomic specified or function specifier for parameter: abc", n.Declarator.Position()))
		}
	case ParameterDeclarationAbstract: // DeclarationSpecifiers AbstractDeclarator
		n.typ = n.AbstractDeclarator.check(c, n.DeclarationSpecifiers.check(c, &isExtern, &isStatic, &isAtomic, &isThreadLocal, &isConst, &isVolatile, &isInline, &isRegister, &isAuto, &isNoreturn))
		if isExtern || isStatic || isAtomic || isThreadLocal || isInline || isAuto || isNoreturn {
			c.errors.add(errorf("%v: storage class or atomic or function specifier for unnamed parameter", n.Position()))
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

//  AbstractDeclarator:
//          Pointer                           // Case AbstractDeclaratorPtr
//  |       Pointer DirectAbstractDeclarator  // Case AbstractDeclaratorDecl
func (n *AbstractDeclarator) check(c *ctx, t Type) (r Type) {
	if t == nil {
		c.errors.add(errorf("TODO %T internal error", n))
		return
	}

	if n == nil {
		return t
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
		}
	}()

	n.typ = n.DirectAbstractDeclarator.check(c, n.Pointer.check(c, t))
	return n.Type()
}

//  DirectAbstractDeclarator:
//          '(' AbstractDeclarator ')'                                                     // Case DirectAbstractDeclaratorDecl
//  |       DirectAbstractDeclarator '[' TypeQualifiers AssignmentExpression ']'           // Case DirectAbstractDeclaratorArr
//  |       DirectAbstractDeclarator '[' "static" TypeQualifiers AssignmentExpression ']'  // Case DirectAbstractDeclaratorStaticArr
//  |       DirectAbstractDeclarator '[' TypeQualifiers "static" AssignmentExpression ']'  // Case DirectAbstractDeclaratorArrStatic
//  |       DirectAbstractDeclarator '[' '*' ']'                                           // Case DirectAbstractDeclaratorArrStar
//  |       DirectAbstractDeclarator '(' ParameterTypeList ')'                             // Case DirectAbstractDeclaratorFunc
func (n *DirectAbstractDeclarator) check(c *ctx, t Type) (r Type) {
	if t == nil {
		c.errors.add(errorf("TODO %T internal error", n))
		return
	}

	if n == nil {
		return t
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
		}
	}()

	switch n.Case {
	case DirectAbstractDeclaratorDecl: // '(' AbstractDeclarator ')'
		return n.AbstractDeclarator.check(c, t)
	case DirectAbstractDeclaratorArr: // DirectAbstractDeclarator '[' TypeQualifiers AssignmentExpression ']'
		return n.DirectAbstractDeclarator.check(c, c.newArrayType(t, arraySize(c, n.AssignmentExpression), n.AssignmentExpression))
	case DirectAbstractDeclaratorStaticArr: // DirectAbstractDeclarator '[' "static" TypeQualifiers AssignmentExpression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectAbstractDeclaratorArrStatic: // DirectAbstractDeclarator '[' TypeQualifiers "static" AssignmentExpression ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectAbstractDeclaratorArrStar: // DirectAbstractDeclarator '[' '*' ']'
		c.errors.add(errorf("TODO %v", n.Case))
	case DirectAbstractDeclaratorFunc: // DirectAbstractDeclarator '(' ParameterTypeList ')'
		fp, isVariadic := n.ParameterTypeList.check(c)
		return n.DirectAbstractDeclarator.check(c, c.newFunctionType(t, fp, isVariadic))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return t //TODO-
}

//  Pointer:
//          '*' TypeQualifiers          // Case PointerTypeQual
//  |       '*' TypeQualifiers Pointer  // Case PointerPtr
//  |       '^' TypeQualifiers          // Case PointerBlock
func (n *Pointer) check(c *ctx, t Type) (r Type) {
	if t == nil {
		c.errors.add(errorf("TODO %T internal error", n))
		return
	}

	if n == nil {
		return t
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
		}
	}()

	switch n.Case {
	case PointerTypeQual: // '*' TypeQualifiers
		return c.newPointerType(t)
	case PointerPtr: // '*' TypeQualifiers Pointer
		return n.Pointer.check(c, c.newPointerType(t))
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

//  DeclarationSpecifiers:
//          StorageClassSpecifier DeclarationSpecifiers  // Case DeclarationSpecifiersStorage
//  |       TypeSpecifier DeclarationSpecifiers          // Case DeclarationSpecifiersTypeSpec
//  |       TypeQualifier DeclarationSpecifiers          // Case DeclarationSpecifiersTypeQual
//  |       FunctionSpecifier DeclarationSpecifiers      // Case DeclarationSpecifiersFunc
//  |       AlignmentSpecifier DeclarationSpecifiers     // Case DeclarationSpecifiersAlignSpec
//  |       "__attribute__"                              // Case DeclarationSpecifiersAttr
func (n *DeclarationSpecifiers) check(c *ctx, isExtern, isStatic, isAtomic, isThreadLocal, isConst, isVolatile, isInline, isRegister, isAuto, isNoreturn *bool) (r Type) {
	if n == nil {
		return c.intT
	}

	n0 := n
	var attr *Attributes
	var ts []TypeSpecifierCase

	defer func(n *DeclarationSpecifiers) {
		if r == nil || r == Invalid {
			//panic(todo("%v: %v %v", n.Position(), ts, TypeString(r)))
			c.errors.add(errorf("TODO %T missed/failed type check: %v", n, ts))
			return
		}

		if attr != nil {
			r = r.setAttr(attr)
		}
	}(n)

	var attrs []*Attributes
	for ; n != nil; n = n.DeclarationSpecifiers {
		switch n.Case {
		case DeclarationSpecifiersStorage: // StorageClassSpecifier DeclarationSpecifiers
			n.StorageClassSpecifier.check(c, isExtern, isStatic, isThreadLocal, isRegister, isAuto)
		case DeclarationSpecifiersTypeSpec: // TypeSpecifier DeclarationSpecifiers
			ts = append(ts, n.TypeSpecifier.Case)
			r = n.TypeSpecifier.check(c, isAtomic)
		case DeclarationSpecifiersTypeQual: // TypeQualifier DeclarationSpecifiers
			n.TypeQualifier.check(c, isConst, isVolatile, isAtomic)
		case DeclarationSpecifiersFunc: // FunctionSpecifier DeclarationSpecifiers
			n.FunctionSpecifier.check(c, isInline, isNoreturn)
		case DeclarationSpecifiersAlignSpec: // AlignmentSpecifier DeclarationSpecifiers
			n.AlignmentSpecifier.check(c)
			//TODO use returned type
		case DeclarationSpecifiersAttr:
			if attr := n.AttributeSpecifierList.check(c); attr != nil {
				attrs = append(attrs, attr)
			}
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	switch len(attrs) {
	case 0:
		// ok
	case 1:
		attr = attrs[0]
	default:
		c.errors.add(errorf("TODO %T", n0.Position()))
	}

	t, ok := c.builtinTypes[ts2String(ts)]
	if ok {
		return t
	}

	switch len(ts) {
	case 0:
		return c.intT
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

	return nil
}

//  AttributeSpecifierList:
//          AttributeSpecifier
//  |       AttributeSpecifierList AttributeSpecifier
func (n *AttributeSpecifierList) check(c *ctx) *Attributes {
	var attr Attributes
	for ; n != nil; n = n.AttributeSpecifierList {
		n.AttributeSpecifier.check(c, &attr)
	}
	if attr.isNonZero {
		return &attr
	}

	return nil
}

//  AttributeSpecifier:
//          "__attribute__" '(' '(' AttributeValueList ')' ')'
func (n *AttributeSpecifier) check(c *ctx, attr *Attributes) {
	n.AttributeValueList.check(c, attr)
}

//  AttributeValueList:
//          AttributeValue
//  |       AttributeValueList ',' AttributeValue
func (n *AttributeValueList) check(c *ctx, attr *Attributes) {
	for ; n != nil; n = n.AttributeValueList {
		n.AttributeValue.check(c, attr)
	}
}

//  AttributeValue:
//          IDENTIFIER                                 // Case AttributeValueIdent
//  |       IDENTIFIER '(' ArgumentExpressionList ')'  // Case AttributeValueExpr
func (n *AttributeValue) check(c *ctx, attr *Attributes) {
	switch n.Case {
	case AttributeValueIdent: // IDENTIFIER
		// ok
	case AttributeValueExpr: // IDENTIFIER '(' ArgumentExpressionList ')'
		n.ArgumentExpressionList.check(c, decay|ignoreUndefined)
		switch n.Token.SrcStr() {
		case "aligned":
			e := n.ArgumentExpressionList.AssignmentExpression
			if n.ArgumentExpressionList.ArgumentExpressionList != nil {
				c.errors.add(errorf("%v: expected one expression", e.Position()))
				break
			}

			v, ok := int64Value(c, e)
			if !ok {
				c.errors.add(errorf("%v: expected a constant integer value", e.Position()))
				return
			}

			if attr.Aligned() > 0 {
				c.errors.add(errorf("%v: multiple 'aligned' specifications", e.Position()))
				return
			}

			if v <= 0 {
				c.errors.add(errorf("%v: alignment must be positive", e.Position()))
				return
			}

			attr.setAligned(v)
		case
			"__vector_size__",
			"vector_size":

			e := n.ArgumentExpressionList.AssignmentExpression
			if n.ArgumentExpressionList.ArgumentExpressionList != nil {
				c.errors.add(errorf("%v: expected one expression", e.Position()))
				break
			}

			v, ok := int64Value(c, e)
			if !ok {
				c.errors.add(errorf("%v: expected a constant integer value", e.Position()))
				return
			}

			if attr.VectorSize() > 0 {
				c.errors.add(errorf("%v: multiple 'vector_size' specifications", e.Position()))
				return
			}

			if v <= 0 {
				c.errors.add(errorf("%v: vector size must be positive", e.Position()))
				return
			}

			if v&(v-1) != 0 {
				c.errors.add(errorf("%v: vector size must be a power of two: %v", e.Position(), v))
				return
			}

			attr.setVectorSize(v)
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

//  ArgumentExpressionList:
//          AssignmentExpression
//  |       ArgumentExpressionList ',' AssignmentExpression
func (n *ArgumentExpressionList) check(c *ctx, mode flags) {
	for ; n != nil; n = n.ArgumentExpressionList {
		n.AssignmentExpression.check(c, mode)
	}
}

//  AlignmentSpecifier:
//          "_Alignas" '(' TypeName ')'            // Case AlignmentSpecifierType
//  |       "_Alignas" '(' ConstantExpression ')'  // Case AlignmentSpecifierExpr
func (n *AlignmentSpecifier) check(c *ctx) (r Type) {
	switch n.Case {
	case AlignmentSpecifierType: // "_Alignas" '(' TypeName ')'
		return n.TypeName.check(c)
	case AlignmentSpecifierExpr: // "_Alignas" '(' ConstantExpression ')'
		return n.ConstantExpression.check(c, decay)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
		return nil
	}
}

//  FunctionSpecifier:
//          "inline"     // Case FunctionSpecifierInline
//  |       "_Noreturn"  // Case FunctionSpecifierNoreturn
func (n *FunctionSpecifier) check(c *ctx, isInline, isNoreturn *bool) {
	switch n.Case {
	case FunctionSpecifierInline: // "inline"
		*isInline = true
	case FunctionSpecifierNoreturn: // "_Noreturn"
		*isNoreturn = true
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

//  TypeQualifier:
//          "const"          // Case TypeQualifierConst
//  |       "restrict"       // Case TypeQualifierRestrict
//  |       "volatile"       // Case TypeQualifierVolatile
//  |       "_Atomic"        // Case TypeQualifierAtomic
//  |       "_Nonnull"       // Case TypeQualifierNonnull
//  |       "__attribute__"  // Case TypeQualifierAttr
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
		n.AttributeSpecifierList.check(c)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
}

//  StorageClassSpecifier:
//          "typedef"             // Case StorageClassSpecifierTypedef
//  |       "extern"              // Case StorageClassSpecifierExtern
//  |       "static"              // Case StorageClassSpecifierStatic
//  |       "auto"                // Case StorageClassSpecifierAuto
//  |       "register"            // Case StorageClassSpecifierRegister
//  |       "_Thread_local"       // Case StorageClassSpecifierThreadLocal
//  |       "__declspec" '(' ')'  // Case StorageClassSpecifierDeclspec
func (n *StorageClassSpecifier) check(c *ctx, isExtern, isStatic, isThreadLocal, isRegister, isAuto *bool) {
	switch n.Case {
	case StorageClassSpecifierTypedef: // "typedef"
		// ok
	case StorageClassSpecifierExtern: // "extern"
		*isExtern = true
	case StorageClassSpecifierStatic: // "static"
		*isStatic = true
	case StorageClassSpecifierAuto: // "auto"
		*isAuto = true
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

//  TypeSpecifier:
//          "void"                           // Case TypeSpecifierVoid
//  |       "char"                           // Case TypeSpecifierChar
//  |       "short"                          // Case TypeSpecifierShort
//  |       "int"                            // Case TypeSpecifierInt
//  |       "__int128"                       // Case TypeSpecifierInt128
//  |       "__uint128_t"                    // Case TypeSpecifierUint128
//  |       "long"                           // Case TypeSpecifierLong
//  |       "float"                          // Case TypeSpecifierFloat
//  |       "_Float16"                       // Case TypeSpecifierFloat16
//  |       "_Decimal64"                     // Case TypeSpecifierDecimal64
//  |       "_Float128"                      // Case TypeSpecifierFloat128
//  |       "_Float128x"                     // Case TypeSpecifierFloat128x
//  |       "double"                         // Case TypeSpecifierDouble
//  |       "signed"                         // Case TypeSpecifierSigned
//  |       "unsigned"                       // Case TypeSpecifierUnsigned
//  |       "_Bool"                          // Case TypeSpecifierBool
//  |       "_Complex"                       // Case TypeSpecifierComplex
//  |       "_Imaginary"                     // Case TypeSpecifierImaginary
//  |       StructOrUnionSpecifier           // Case TypeSpecifierStructOrUnion
//  |       EnumSpecifier                    // Case TypeSpecifierEnum
//  |       TYPENAME                         // Case TypeSpecifierTypeName
//  |       "typeof" '(' ExpressionList ')'  // Case TypeSpecifierTypeofExpr
//  |       "typeof" '(' TypeName ')'        // Case TypeSpecifierTypeofType
//  |       AtomicTypeSpecifier              // Case TypeSpecifierAtomic
//  |       "_Float32"                       // Case TypeSpecifierFloat32
//  |       "_Float64"                       // Case TypeSpecifierFloat64
//  |       "_Float32x"                      // Case TypeSpecifierFloat32x
//  |       "_Float64x"                      // Case TypeSpecifierFloat64x
func (n *TypeSpecifier) check(c *ctx, isAtomic *bool) (r Type) {
	if n == nil {
		return Invalid
	}

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
		// ok
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
			return x.Type()
		}

		c.errors.add(errorf("%v: undefined type name: %s", n.Position(), n.Token.Src()))
	case TypeSpecifierTypeofExpr: // "typeof" '(' ExpressionList ')'
		return n.ExpressionList.check(c, decay)
	case TypeSpecifierTypeofType: // "typeof" '(' TypeName ')'
		return n.TypeName.check(c)
	case TypeSpecifierAtomic: // AtomicTypeSpecifier
		*isAtomic = true
		return n.AtomicTypeSpecifier.check(c)
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

func (n *AtomicTypeSpecifier) check(c *ctx) (r Type) {
	return n.TypeName.check(c)
}

//  EnumSpecifier:
//          "enum" IDENTIFIER '{' EnumeratorList ',' '}'  // Case EnumSpecifierDef
//  |       "enum" IDENTIFIER                             // Case EnumSpecifierTag
func (n *EnumSpecifier) check(c *ctx) (r Type) {
	if n == nil {
		return Invalid
	}

	if n.typ != nil {
		return n.Type()
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
		}
	}()

	tag := ""
	if n.Token2.s != nil {
		tag = n.Token2.SrcStr()
	}
	switch n.Case {
	case EnumSpecifierDef: // "enum" IDENTIFIER '{' EnumeratorList ',' '}'
		var t Type
		var iota, min int64
		var max uint64
		var list []*Enumerator
		for l := n.EnumeratorList; l != nil; l = l.EnumeratorList {
			en := l.Enumerator
			list = append(list, en)
			iota = en.check(c, iota)
			switch x := en.Value().(type) {
			case Int64Value:
				v := int64(x)
				min = mathutil.MinInt64(min, v)
				if v >= 0 {
					max = mathutil.MaxUint64(max, uint64(v))
				}
			case UInt64Value:
				v := uint64(x)
				max = mathutil.MaxUint64(max, v)
			}
		}
		switch {
		case min >= math.MinInt32 && max <= math.MaxInt32:
			t = c.intT
		case min >= 0 && max <= math.MaxUint32:
			t = c.ast.kinds[UInt]
		case min >= math.MinInt64 && max < math.MaxInt64:
			t = c.ast.kinds[Long]
			if t.Size() < 8 {
				t = c.ast.kinds[LongLong]
			}
		default:
			t = c.ast.kinds[ULong]
			if t.Size() < 8 {
				t = c.ast.kinds[ULongLong]
			}
		}
		for _, v := range list {
			v.typ = t
		}
		n.typ = c.newEnumType(tag, t, list)
	case EnumSpecifierTag: // "enum" IDENTIFIER
		if x := n.resolutionScope.enum(n.Token2); x != nil {
			switch {
			case x.typ == nil:
				t := c.newEnumType(tag, nil, nil)
				t.forward = x
				n.typ = t
			default:
				n.typ = x.typ
			}
			break
		}

		n.typ = c.newEnumType(tag, nil, nil)
		c.ast.scope.declare(tag, n)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

func (n *Enumerator) check(c *ctx, iota int64) int64 {
	switch n.Case {
	case EnumeratorIdent: // IDENTIFIER
		n.typ = c.int64T
		n.val = Int64Value(iota)
	case EnumeratorExpr: // IDENTIFIER '=' ConstantExpression
		n.typ = n.ConstantExpression.check(c, decay)
		switch n.val = n.ConstantExpression.Value(); x := n.Value().(type) {
		case Int64Value:
			return int64(x) + 1
		case UInt64Value:
			return int64(x) + 1
		case *UnknownValue:
			// ok
		default:
			c.errors.add(errorf("internal error: %T", x))
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return iota + 1
}

//  StructOrUnionSpecifier:
//          StructOrUnion IDENTIFIER '{' StructDeclarationList '}'  // Case StructOrUnionSpecifierDef
//  |       StructOrUnion IDENTIFIER                                // Case StructOrUnionSpecifierTag
func (n *StructOrUnionSpecifier) check(c *ctx) (r Type) {
	if n == nil {
		return Invalid
	}

	if n.typ != nil {
		return n.Type()
	}

	tag := ""
	if n.Token.s != nil {
		tag = n.Token.SrcStr()
	}
	switch n.Case {
	case StructOrUnionSpecifierDef: // StructOrUnion IDENTIFIER '{' StructDeclarationList '}'
		defer func() {
			if r == nil || r == Invalid {
				c.errors.add(errorf("TODO %T missed/failed type check", n))
			}
		}()

		switch {
		case n.StructOrUnion.Case == StructOrUnionUnion:
			n.typ = c.newUnionType(tag, nil, -1, 1)
		default:
			n.typ = c.newStructType(tag, nil, -1, 1)
		}

		n.StructDeclarationList.check(c, n)
	case StructOrUnionSpecifierTag: // StructOrUnion IDENTIFIER
		if x := n.resolutionScope.structOrUnion(n.Token); x != nil {
			if n.StructOrUnion.Case != x.StructOrUnion.Case {
				c.errors.add(errorf("%v: mismatched struct/union tag", n.Token.Position()))
				break
			}

			switch {
			case x.typ == nil:
				switch {
				case n.StructOrUnion.Case == StructOrUnionUnion:
					t := c.newUnionType(tag, nil, -1, 1)
					t.forward = x
					n.typ = t
				default:
					t := c.newStructType(tag, nil, -1, 1)
					t.forward = x
					n.typ = t
				}
			default:
				n.typ = x.typ
			}
			break
		}

		switch {
		case n.StructOrUnion.Case == StructOrUnionUnion:
			n.typ = c.newUnionType(tag, nil, -1, 1)
		default:
			n.typ = c.newStructType(tag, nil, -1, 1)
		}
		c.ast.scope.declare(tag, n)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  StructDeclarationList:
//          StructDeclaration
//  |       StructDeclarationList StructDeclaration
func (n *StructDeclarationList) check(c *ctx, s *StructOrUnionSpecifier) {
	if n == nil {
		return
	}

	defer func() {
		if s.typ == nil || s.typ == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
		}
	}()

	var fields []*Field
	for l := n; l != nil; l = l.StructDeclarationList {
		fields = append(fields, l.StructDeclaration.check(c)...)
	}
	// for _, v := range fields {
	// 	switch {
	// 	case v.declarator == nil:
	// 		trc("-: - %s", v.Type())
	// 	default:
	// 		trc("%v: %s %s", v.declarator.Position(), v.Name(), v.Type())
	// 	}
	// }

	isUnion := s.StructOrUnion.Case == StructOrUnionUnion
	var brk, unionBits int64
	maxAlignBytes := 1
	for i, f := range fields {
		if f == nil {
			c.errors.add(errorf("TODO %T", n))
			return
		}

		f.ordinal = i

		switch {
		case f.isBitField:
			f.accessBytes = bits2AccessBytes(f.valueBits)
			switch {
			case isUnion:
				f.mask = (uint64(1)<<f.valueBits - 1)
				unionBits = mathutil.MaxInt64(unionBits, 8*f.accessBytes)
			default:
				brkOffBytes := brk >> 3
				if mod := brkOffBytes % f.accessBytes; mod != 0 {
					brkOffBytes += f.accessBytes - mod
					brk = brkOffBytes << 3
				}
				f.offsetBytes = brkOffBytes
				f.offsetBits = int(brk - 8*f.offsetBytes)
				f.mask = (uint64(1)<<f.valueBits - 1) << f.offsetBits
				brk += f.valueBits
			}
		default:
			sz := f.Type().Size()
			if f.Type().IsIncomplete() && f.Type().Kind() == Array { // Flexible array member
				sz = 0
			}
			al := f.Type().Align()
			if al > maxAlignBytes {
				maxAlignBytes = al
			}
			if !isUnion {
				brk = roundup(brk, 8*int64(al))
			}
			f.accessBytes = sz
			f.offsetBytes = brk >> 3
			f.valueBits = 8 * sz
			switch {
			case isUnion:
				unionBits = mathutil.MaxInt64(unionBits, 8*sz)
			default:
				brk += 8 * sz
			}
		}
	}
	brk = roundup(brk, int64(maxAlignBytes*8))
	switch {
	case isUnion:
		t := s.typ.(*UnionType)
		t.fields = fields
		t.size = unionBits >> 3
		t.align = maxAlignBytes
	default:
		t := s.typ.(*StructType)
		t.fields = fields
		t.size = brk >> 3
		t.align = maxAlignBytes
	}
}

func (n *StructDeclaration) check(c *ctx) (r []*Field) {
	var isAtomic, isConst, isVolatile bool
	t := n.SpecifierQualifierList.check(c, &isAtomic, &isConst, &isVolatile)
	switch {
	case n.StructDeclaratorList == nil:
		return []*Field{{typ: newTyper(t)}}
	default:
		return n.StructDeclaratorList.check(c, t, isAtomic, isConst, isVolatile)
	}
}

func (n *SpecifierQualifierList) check(c *ctx, isAtomic, isConst, isVolatile *bool) (r Type) {
	if n == nil {
		return c.intT
	}

	var ts []TypeSpecifierCase

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO missed/failed type check %v: %v %T", n.Position(), ts, n))
		}
	}()

	for ; n != nil; n = n.SpecifierQualifierList {
		switch n.Case {
		case SpecifierQualifierListTypeSpec: // TypeSpecifier SpecifierQualifierList
			ts = append(ts, n.TypeSpecifier.Case)
			r = n.TypeSpecifier.check(c, isAtomic)
		case SpecifierQualifierListTypeQual: // TypeQualifier SpecifierQualifierList
			n.TypeQualifier.check(c, isConst, isVolatile, isAtomic)
		case SpecifierQualifierListAlignSpec: // AlignmentSpecifier SpecifierQualifierList
			c.errors.add(errorf("TODO %v", n.Case))
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
		return c.intT
	case 1:
		switch ts[0] {
		case
			TypeSpecifierAtomic,
			TypeSpecifierEnum,
			TypeSpecifierStructOrUnion,
			TypeSpecifierTypeName,
			TypeSpecifierTypeofExpr,
			TypeSpecifierTypeofType:

			return r
		}
	}

	return nil
}

func (n *StructDeclaratorList) check(c *ctx, t Type, isAtomic, isConst, isVolatile bool) (r []*Field) {
	if t == nil {
		c.errors.add(errorf("TODO %T internal error", n))
		return
	}

	for ; n != nil; n = n.StructDeclaratorList {
		r = append(r, n.StructDeclarator.check(c, t, isAtomic, isConst, isVolatile))
	}
	return r
}

func (n *StructDeclarator) check(c *ctx, t Type, isAtomic, isConst, isVolatile bool) (r *Field) {
	if t == nil {
		c.errors.add(errorf("TODO %T internal error", n))
		return
	}

	if n.Declarator != nil {
		n.Declarator.isAtomic = isAtomic
		n.Declarator.isConst = isConst
		n.Declarator.isVolatile = isVolatile
	}
	switch n.Case {
	case StructDeclaratorDecl: // Declarator
		return &Field{declarator: n.Declarator, typ: newTyper(n.Declarator.check(c, t))}
	case StructDeclaratorBitField: // Declarator ':' ConstantExpression
		t := n.ConstantExpression.check(c, decay)
		if !isIntegerType(t) {
			c.errors.add(errorf("%v: expected integer expression: %s", n.ConstantExpression.Position(), t))
			break
		}

		var bits int64
		switch x := n.ConstantExpression.Value().(type) {
		case Int64Value:
			bits = int64(x)
		case UInt64Value:
			bits = int64(x)
		}
		if bits < 0 || bits > 64 {
			c.errors.add(errorf("%v: value out of range: %v", n.ConstantExpression.Position(), bits))
			break
		}

		return &Field{declarator: n.Declarator, typ: newTyper(n.Declarator.check(c, t)), valueBits: bits, isBitField: true}
	default:
		c.errors.add(errorf("internal error: %v %T", n.Case, n))
	}
	return nil //TODO-
}

//  AssignmentExpression:
//          ConditionalExpression                       // Case AssignmentExpressionCond
//  |       UnaryExpression '=' AssignmentExpression    // Case AssignmentExpressionAssign
//  |       UnaryExpression "*=" AssignmentExpression   // Case AssignmentExpressionMul
//  |       UnaryExpression "/=" AssignmentExpression   // Case AssignmentExpressionDiv
//  |       UnaryExpression "%=" AssignmentExpression   // Case AssignmentExpressionMod
//  |       UnaryExpression "+=" AssignmentExpression   // Case AssignmentExpressionAdd
//  |       UnaryExpression "-=" AssignmentExpression   // Case AssignmentExpressionSub
//  |       UnaryExpression "<<=" AssignmentExpression  // Case AssignmentExpressionLsh
//  |       UnaryExpression ">>=" AssignmentExpression  // Case AssignmentExpressionRsh
//  |       UnaryExpression "&=" AssignmentExpression   // Case AssignmentExpressionAnd
//  |       UnaryExpression "^=" AssignmentExpression   // Case AssignmentExpressionXor
//  |       UnaryExpression "|=" AssignmentExpression   // Case AssignmentExpressionOr
func (n *AssignmentExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
		}
	}()

	switch n.Case {
	case AssignmentExpressionCond: // ConditionalExpression
		n.typ = n.ConditionalExpression.check(c, mode)
		n.val = n.ConditionalExpression.eval(c, mode)
	case
		AssignmentExpressionAssign, // UnaryExpression '=' AssignmentExpression
		AssignmentExpressionMul,    // UnaryExpression "*=" AssignmentExpression
		AssignmentExpressionDiv,    // UnaryExpression "/=" AssignmentExpression
		AssignmentExpressionMod,    // UnaryExpression "%=" AssignmentExpression
		AssignmentExpressionAdd,    // UnaryExpression "+=" AssignmentExpression
		AssignmentExpressionSub,    // UnaryExpression "-=" AssignmentExpression
		AssignmentExpressionLsh,    // UnaryExpression "<<=" AssignmentExpression
		AssignmentExpressionRsh,    // UnaryExpression ">>=" AssignmentExpression
		AssignmentExpressionAnd,    // UnaryExpression "&=" AssignmentExpression
		AssignmentExpressionXor,    // UnaryExpression "^=" AssignmentExpression
		AssignmentExpressionOr:     // UnaryExpression "|=" AssignmentExpression

		n.typ = n.UnaryExpression.check(c, mode)
		a := n.Type()
		b := n.AssignmentExpression.check(c, mode)
		if !isModifiableLvalue(a) {
			c.errors.add(errorf("%v: left operand shall be a modifiable lvalue", n.UnaryExpression.Position()))
			break
		}

		switch {
		case
			// — the left operand has qualified or unqualified arithmetic type and the
			// right has arithmetic type;
			isArithmeticType(a) && isArithmeticType(b),

			// — the left operand has a qualified or unqualified version of a structure or
			// union type compatible with the type of the right;
			a.Kind() == Struct && b.Kind() == Struct || a.Kind() == Union && b.Kind() == Union,

			// — both operands are pointers to qualified or unqualified versions of
			// compatible types, and the type pointed to by the left has all the qualifiers
			// of the type pointed to by the right;
			//
			// — one operand is a pointer to an object or incomplete type and the other is
			// a pointer to a qualified or unqualified version of void, and the type
			// pointed to by the left has all the qualifiers of the type pointed to by the
			// right;
			isPointerType(a) && isPointerType(b),

			// — the left operand is a pointer and the right is a null pointer constant; or
			isPointerType(a) && isIntegerType(b),

			// — the left operand has type _Bool and the right is a pointer.
			a.Kind() == Bool && isPointerType(b):

			// ok
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  ConditionalExpression:
//          LogicalOrExpression                                           // Case ConditionalExpressionLOr
//  |       LogicalOrExpression '?' ExpressionList ':' ConditionalExpression  // Case ConditionalExpressionCond
func (n *ConditionalExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case ConditionalExpressionLOr: // LogicalOrExpression
		n.typ = n.LogicalOrExpression.check(c, mode)
	case ConditionalExpressionCond: // LogicalOrExpression '?' ExpressionList ':' ConditionalExpression
		mode = mode.add(decay)
		t1 := n.LogicalOrExpression.check(c, mode)
		if !isScalarType(t1) {
			c.errors.add(errorf("%v: operand shall have scalar type: %s", n.LogicalOrExpression.Position(), t1))
			break
		}

		t2 := t1
		if n.ExpressionList != nil {
			t2 = n.ExpressionList.check(c, mode)
		}
		switch t3 := n.ConditionalExpression.check(c, mode); {
		case
			// both operands have arithmetic type;
			isArithmeticType(t2) && isArithmeticType(t3):
			n.typ = usualArithmeticConversions(t2, t3)
		case
			// both operands have the same structure or union type;
			(t2.Kind() == Struct || t2.Kind() == Union) && t3.Kind() == t2.Kind():
			n.typ = t2
		case
			// both operands have void type;
			t2.Kind() == Void && t3.Kind() == Void:
			n.typ = t2
		case
			// both operands are pointers to qualified or unqualified versions of compatible types;
			isPointerType(t2) && isPointerType(t3):
			n.typ = t2
		case
			// one operand is a pointer and the other is a null pointer constant; or
			isPointerType(t2) && isIntegerType(t3):
			n.typ = t2
		case
			isIntegerType(t2) && isPointerType(t3):
			n.typ = t3
		case t2.Kind() == Void:
			n.typ = t2
		case t3.Kind() == Void:
			n.typ = t3
		default:
			c.errors.add(errorf("TODO %v, t1 %v, t2 %v, t3 %v", n.Case, t1, t2, t3))
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  LogicalOrExpression:
//          LogicalAndExpression                           // Case LogicalOrExpressionLAnd
//  |       LogicalOrExpression "||" LogicalAndExpression  // Case LogicalOrExpressionLOr
func (n *LogicalOrExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case LogicalOrExpressionLAnd: // LogicalAndExpression
		n.typ = n.LogicalAndExpression.check(c, mode)
	case LogicalOrExpressionLOr: // LogicalOrExpression "||" LogicalAndExpression
		mode = mode.add(decay)
		switch a, b := n.LogicalOrExpression.check(c, mode), n.LogicalAndExpression.check(c, mode); {
		case !isScalarType(a) || !isScalarType(b):
			c.errors.add(errorf("%v: operands shall be scalars: %s and %s", n.Token.Position(), a, b))
		default:
			n.typ = c.intT
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  LogicalAndExpression:
//          InclusiveOrExpression                            // Case LogicalAndExpressionOr
//  |       LogicalAndExpression "&&" InclusiveOrExpression  // Case LogicalAndExpressionLAnd
func (n *LogicalAndExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case LogicalAndExpressionOr: // InclusiveOrExpression
		n.typ = n.InclusiveOrExpression.check(c, mode)
	case LogicalAndExpressionLAnd: // LogicalAndExpression "&&" InclusiveOrExpression
		mode = mode.add(decay)
		switch a, b := n.LogicalAndExpression.check(c, mode), n.InclusiveOrExpression.check(c, mode); {
		case !isScalarType(a) || !isScalarType(b):
			c.errors.add(errorf("%v: operands shall be scalars: %s and %s", n.Token.Position(), a, b))
		default:
			n.typ = c.intT
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  InclusiveOrExpression:
//          ExclusiveOrExpression                            // Case InclusiveOrExpressionXor
//  |       InclusiveOrExpression '|' ExclusiveOrExpression  // Case InclusiveOrExpressionOr
func (n *InclusiveOrExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case InclusiveOrExpressionXor: // ExclusiveOrExpression
		n.typ = n.ExclusiveOrExpression.check(c, mode)
	case InclusiveOrExpressionOr: // InclusiveOrExpression '|' ExclusiveOrExpression
		mode = mode.add(decay)
		switch a, b := n.InclusiveOrExpression.check(c, mode), n.ExclusiveOrExpression.check(c, mode); {
		case !isIntegerType(a) || !isIntegerType(b):
			c.errors.add(errorf("%v: operands shall have integer type: %s and %s", n.Token.Position(), a, b))
		default:
			n.typ = usualArithmeticConversions(a, b)
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  ExclusiveOrExpression:
//          AndExpression                            // Case ExclusiveOrExpressionAnd
//  |       ExclusiveOrExpression '^' AndExpression  // Case ExclusiveOrExpressionXor
func (n *ExclusiveOrExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case ExclusiveOrExpressionAnd: // AndExpression
		n.typ = n.AndExpression.check(c, mode)
	case ExclusiveOrExpressionXor: // ExclusiveOrExpression '^' AndExpression
		mode = mode.add(decay)
		switch a, b := n.ExclusiveOrExpression.check(c, mode), n.AndExpression.check(c, mode); {
		case !isIntegerType(a) || !isIntegerType(b):
			c.errors.add(errorf("%v: operands shall have integer type: %s and %s", n.Token.Position(), a, b))
		default:
			n.typ = usualArithmeticConversions(a, b)
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  AndExpression:
//          EqualityExpression                    // Case AndExpressionEq
//  |       AndExpression '&' EqualityExpression  // Case AndExpressionAnd
func (n *AndExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case AndExpressionEq: // EqualityExpression
		n.typ = n.EqualityExpression.check(c, mode)
	case AndExpressionAnd: // AndExpression '&' EqualityExpression
		mode = mode.add(decay)
		switch a, b := n.AndExpression.check(c, mode), n.EqualityExpression.check(c, mode); {
		case !isIntegerType(a) || !isIntegerType(b):
			c.errors.add(errorf("%v: operands shall have integer type: %s and %s", n.Token.Position(), a, b))
		default:
			n.typ = usualArithmeticConversions(a, b)
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  EqualityExpression:
//          RelationalExpression                          // Case EqualityExpressionRel
//  |       EqualityExpression "==" RelationalExpression  // Case EqualityExpressionEq
//  |       EqualityExpression "!=" RelationalExpression  // Case EqualityExpressionNeq
func (n *EqualityExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case EqualityExpressionRel: // RelationalExpression
		n.typ = n.RelationalExpression.check(c, mode)
	case
		EqualityExpressionEq,  // EqualityExpression "==" RelationalExpression
		EqualityExpressionNeq: // EqualityExpression "!=" RelationalExpression

		mode = mode.add(decay)
		switch a, b := n.EqualityExpression.check(c, mode), n.RelationalExpression.check(c, mode); {
		case
			// both operands have arithmetic type;
			isArithmeticType(a) && isArithmeticType(b),

			// both operands are pointers to qualified or unqualified versions of
			// compatible types;
			//
			// one operand is a pointer to an object or incomplete type and the other is a
			// pointer to a qualified or unqualified version of void;
			isPointerType(a) && isPointerType(b),

			// one operand is a pointer and the other is a null pointer constant.
			isPointerType(a) && isIntegerType(b) || isPointerType(b) && isIntegerType(a):

			n.typ = c.intT
		default:
			c.errors.add(errorf("%v: invalid operands: %v and %v", n.Token.Position(), a, b))
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  RelationalExpression:
//          ShiftExpression                            // Case RelationalExpressionShift
//  |       RelationalExpression '<' ShiftExpression   // Case RelationalExpressionLt
//  |       RelationalExpression '>' ShiftExpression   // Case RelationalExpressionGt
//  |       RelationalExpression "<=" ShiftExpression  // Case RelationalExpressionLeq
//  |       RelationalExpression ">=" ShiftExpression  // Case RelationalExpressionGeq
func (n *RelationalExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case RelationalExpressionShift: // ShiftExpression
		n.typ = n.ShiftExpression.check(c, mode)
	case
		RelationalExpressionLt,  // RelationalExpression '<' ShiftExpression
		RelationalExpressionGt,  // RelationalExpression '>' ShiftExpression
		RelationalExpressionLeq, // RelationalExpression "<=" ShiftExpression
		RelationalExpressionGeq: // RelationalExpression ">=" ShiftExpression

		n.typ = c.intT
		mode = mode.add(decay)
		switch a, b := n.RelationalExpression.check(c, mode), n.ShiftExpression.check(c, mode); {
		case
			// both operands have real type;
			isRealType(a) && isRealType(b),
			// both operands are pointers to qualified or unqualified versions of
			// compatible object types;
			//
			// both operands are pointers to qualified or unqualified versions of
			// compatible incomplete types
			//
			// gcc allows mixing ints and pointers
			isScalarType(a) && isScalarType(b):

			// ok
		default:
			c.errors.add(errorf("%v: invalid operands: %s and %s", n.Token.Position(), a, b))
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  ShiftExpression:
//          AdditiveExpression                       // Case ShiftExpressionAdd
//  |       ShiftExpression "<<" AdditiveExpression  // Case ShiftExpressionLsh
//  |       ShiftExpression ">>" AdditiveExpression  // Case ShiftExpressionRsh
func (n *ShiftExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case ShiftExpressionAdd: // AdditiveExpression
		n.typ = n.AdditiveExpression.check(c, mode)
	case
		ShiftExpressionLsh, // ShiftExpression "<<" AdditiveExpression
		ShiftExpressionRsh: // ShiftExpression ">>" AdditiveExpression

		mode = mode.add(decay)
		switch a, b := n.ShiftExpression.check(c, mode), n.AdditiveExpression.check(c, mode); {
		case !isScalarType(a) || !isScalarType(b):
			c.errors.add(errorf("%v: operands shall be a scalars: %s and %s", n.Token.Position(), a, b))
		default:
			n.typ = integerPromotion(a)
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  AdditiveExpression:
//          MultiplicativeExpression                         // Case AdditiveExpressionMul
//  |       AdditiveExpression '+' MultiplicativeExpression  // Case AdditiveExpressionAdd
//  |       AdditiveExpression '-' MultiplicativeExpression  // Case AdditiveExpressionSub
func (n *AdditiveExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case AdditiveExpressionMul: // MultiplicativeExpression
		n.typ = n.MultiplicativeExpression.check(c, mode)
	case AdditiveExpressionAdd: // AdditiveExpression '+' MultiplicativeExpression
		mode = mode.add(decay)
		switch a, b := n.AdditiveExpression.check(c, mode), n.MultiplicativeExpression.check(c, mode); {
		case
			// For addition, either both operands shall have arithmetic type
			isArithmeticType(a) && isArithmeticType(b):
			n.typ = usualArithmeticConversions(a, b)
		case
			// or one operand shall be a pointer to an object type and the other shall have
			// integer type.
			isPointerType(a) && isIntegerType(b):
			n.typ = a
		case isIntegerType(a) && isPointerType(b):
			n.typ = b
		default:
			c.errors.add(errorf("%v: invalid operands: %s and %s", n.Token.Position(), a, b))
		}
	case AdditiveExpressionSub: // AdditiveExpression '-' MultiplicativeExpression
		mode = mode.add(decay)
		switch a, b := n.AdditiveExpression.check(c, mode), n.MultiplicativeExpression.check(c, mode); {
		case
			// both operands have arithmetic type;
			isArithmeticType(a) && isArithmeticType(b):
			n.typ = usualArithmeticConversions(a, b)
		case
			// both operands are pointers to qualified or unqualified versions of
			// compatible object types;
			isPointerType(a) && isPointerType(b):
			n.typ = c.ptrDiffT(n)
		case
			// the left operand is a pointer to an object type and the right operand has
			// integer type.
			isPointerType(a) && isIntegerType(b):
			n.typ = a
		default:
			c.errors.add(errorf("%v: invalid operands: %s and %s", n.Token.Position(), a, b))
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  MultiplicativeExpression:
//          CastExpression                               // Case MultiplicativeExpressionCast
//  |       MultiplicativeExpression '*' CastExpression  // Case MultiplicativeExpressionMul
//  |       MultiplicativeExpression '/' CastExpression  // Case MultiplicativeExpressionDiv
//  |       MultiplicativeExpression '%' CastExpression  // Case MultiplicativeExpressionMod
func (n *MultiplicativeExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case MultiplicativeExpressionCast: // CastExpression
		n.typ = n.CastExpression.check(c, mode)
	case
		MultiplicativeExpressionMul, // MultiplicativeExpression '*' CastExpression
		MultiplicativeExpressionDiv: // MultiplicativeExpression '/' CastExpression

		mode = mode.add(decay)
		switch a, b := n.MultiplicativeExpression.check(c, mode), n.CastExpression.check(c, mode); {
		case !isArithmeticType(a) || !isArithmeticType(b):
			c.errors.add(errorf("%v: operands shall have arithmetic type: %s and %s", n.Token.Position(), a, b))
		default:
			n.typ = usualArithmeticConversions(a, b)
		}
	case MultiplicativeExpressionMod: // MultiplicativeExpression '%' CastExpression
		mode = mode.add(decay)
		switch a, b := n.MultiplicativeExpression.check(c, mode), n.CastExpression.check(c, mode); {
		case !isIntegerType(a) || !isIntegerType(b):
			c.errors.add(errorf("%v: operands shall have integer type: %s and %s", n.Token.Position(), a, b))
		default:
			n.typ = usualArithmeticConversions(a, b)
		}
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  CastExpression:
//          UnaryExpression                  // Case CastExpressionUnary
//  |       '(' TypeName ')' CastExpression  // Case CastExpressionCast
func (n *CastExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case CastExpressionUnary: // UnaryExpression
		n.typ = n.UnaryExpression.check(c, mode)
	case CastExpressionCast: // '(' TypeName ')' CastExpression
		n.typ = n.TypeName.check(c)
		n.CastExpression.check(c, mode.add(decay))
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

func (n *TypeName) check(c *ctx) (r Type) {
	var dummy bool
	n.typ = n.AbstractDeclarator.check(c, n.SpecifierQualifierList.check(c, &dummy, &dummy, &dummy))
	return n.Type()
}

//  UnaryExpression:
//          PostfixExpression            // Case UnaryExpressionPostfix
//  |       "++" UnaryExpression         // Case UnaryExpressionInc
//  |       "--" UnaryExpression         // Case UnaryExpressionDec
//  |       '&' CastExpression           // Case UnaryExpressionAddrof
//  |       '*' CastExpression           // Case UnaryExpressionDeref
//  |       '+' CastExpression           // Case UnaryExpressionPlus
//  |       '-' CastExpression           // Case UnaryExpressionMinus
//  |       '~' CastExpression           // Case UnaryExpressionCpl
//  |       '!' CastExpression           // Case UnaryExpressionNot
//  |       "sizeof" UnaryExpression     // Case UnaryExpressionSizeofExpr
//  |       "sizeof" '(' TypeName ')'    // Case UnaryExpressionSizeofType
//  |       "&&" IDENTIFIER              // Case UnaryExpressionLabelAddr
//  |       "_Alignof" UnaryExpression   // Case UnaryExpressionAlignofExpr
//  |       "_Alignof" '(' TypeName ')'  // Case UnaryExpressionAlignofType
//  |       "__imag__" UnaryExpression   // Case UnaryExpressionImag
//  |       "__real__" UnaryExpression   // Case UnaryExpressionReal
func (n *UnaryExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
			return
		}

		n.eval(c, 0)
	}()

	switch n.Case {
	case UnaryExpressionPostfix: // PostfixExpression
		n.typ = n.PostfixExpression.check(c, mode)
	case
		UnaryExpressionInc, // "++" UnaryExpression
		UnaryExpressionDec: // "--" UnaryExpression

		n.typ = n.UnaryExpression.check(c, mode.add(decay))
		if !isRealType(n.Type()) && !isPointerType(n.Type()) {
			c.errors.add(errorf("%v: operand shall have real or pointer type: %s", n.UnaryExpression.Position(), n.Type()))
		}
	case UnaryExpressionAddrof: // '&' CastExpression
		switch t := n.CastExpression.check(c, mode.del(decay)); {
		case
			// The operand of the unary & operator shall be either a function designator,
			t.Kind() == Function,

			// the result of a [] or unary * operator, or an lvalue that designates an
			// object that is not a bit-field and is not declared with the register
			// storage-class specifier.
			isLvalue(t),

			// gcc extension: gcc-9.1.0/gcc/testsuite/gcc.c-torture/compile/20010327-1.c
			t.Kind() == Void:

			n.typ = c.newPointerType(t)
		default:
			c.errors.add(errorf("%v: invalid operand: %s", n.CastExpression.Position(), t))
		}
	case UnaryExpressionDeref: // '*' CastExpression
		switch t := n.CastExpression.check(c, mode.add(decay)); t.Kind() {
		case Ptr:
			n.typ = c.decay(t.(*PointerType).Elem(), mode)
		default:
			c.errors.add(errorf("%v: operand shall be a pointer: %s", n.CastExpression.Position(), t))
		}
	case
		UnaryExpressionPlus,  // '+' CastExpression
		UnaryExpressionMinus: // '-' CastExpression

		n.typ = integerPromotion(n.CastExpression.check(c, mode.add(decay)))
		if !isArithmeticType(n.Type()) {
			c.errors.add(errorf("%v: expected arithmetic type: %s", n.Position(), n.CastExpression.Type()))
		}
	case UnaryExpressionCpl: // '~' CastExpression
		t := n.CastExpression.check(c, mode.add(decay))
		switch {
		case isIntegerType(t):
			n.typ = integerPromotion(t)
		case isComplexType(t): // GCC extension, complex conjugate
			n.typ = t
		default:
			c.errors.add(errorf("%v: expected integer type: %s", n.Position(), n.CastExpression.Type()))
		}
	case UnaryExpressionNot: // '!' CastExpression
		t := n.CastExpression.check(c, mode.add(decay))
		if !isScalarType(t) {
			c.errors.add(errorf("%v: expected scalar type: %s", n.Position(), n.CastExpression.Type()))
		}
		n.typ = c.intT
	case UnaryExpressionSizeofExpr: // "sizeof" UnaryExpression
		n.typ = c.sizeT(n)
		t := n.UnaryExpression.check(c, mode.del(decay))
		if t.Kind() == Function {
			t = c.newPointerType(t)
		}
		if t.IsIncomplete() {
			if x, ok := t.(*ArrayType); ok && x.IsVLA() {
				break
			}

			c.errors.add(errorf("%v: sizeof incomplete type: %s", n.UnaryExpression.Position(), t))
		}
		n.val = UInt64Value(t.Size())
	case UnaryExpressionSizeofType: // "sizeof" '(' TypeName ')'
		n.typ = c.sizeT(n)
		t := n.TypeName.check(c)
		if t.IsIncomplete() {
			if x, ok := t.(*ArrayType); ok && x.IsVLA() {
				break
			}

			c.errors.add(errorf("%v: sizeof incomplete type: %s", n.TypeName.Position(), t))
		}
		n.val = UInt64Value(t.Size())
	case UnaryExpressionLabelAddr: // "&&" IDENTIFIER
		n.typ = c.pvoidT
	case UnaryExpressionAlignofExpr: // "_Alignof" UnaryExpression
		t := n.UnaryExpression.check(c, mode.add(decay))
		n.val, n.typ = UInt64Value(t.Align()), c.sizeT(n)
	case UnaryExpressionAlignofType: // "_Alignof" '(' TypeName ')'
		t := n.TypeName.check(c)
		n.val, n.typ = UInt64Value(t.Align()), c.sizeT(n)
	case
		UnaryExpressionImag, // "__imag__" UnaryExpression
		UnaryExpressionReal: // "__real__" UnaryExpression

		t := n.UnaryExpression.check(c, mode.add(decay))
		if !isComplexType(t) {
			c.errors.add(errorf("%v: expected complex type: %s", n.Position(), t))
			break
		}

		n.typ = c.ast.kinds[correspondingRealKinds[t.Kind()]]
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

//  PostfixExpression:
//          PrimaryExpression                                 // Case PostfixExpressionPrimary
//  |       PostfixExpression '[' ExpressionList ']'          ;// Case PostfixExpressionIndex
//  |       PostfixExpression '(' ArgumentExpressionList ')'  // Case PostfixExpressionCall
//  |       PostfixExpression '.' IDENTIFIER                  // Case PostfixExpressionSelect
//  |       PostfixExpression "->" IDENTIFIER                 // Case PostfixExpressionPSelect
//  |       PostfixExpression "++"                            // Case PostfixExpressionInc
//  |       PostfixExpression "--"                            // Case PostfixExpressionDec
//  |       '(' TypeName ')' '{' InitializerList ',' '}'      // Case PostfixExpressionComplit
func (n *PostfixExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check %v", n, n.Case))
			return
		}

		r = c.decay(r, mode)
		n.typ = r
		n.eval(c, 0)
	}()

	switch n.Case {
	case PostfixExpressionPrimary: // PrimaryExpression
		n.typ = n.PrimaryExpression.check(c, mode)
	case PostfixExpressionIndex: // PostfixExpression '[' ExpressionList ']'
		// One of the expressions shall have type ‘‘pointer to object type’’, the other
		// expression shall have integer type, and the result has type ‘‘type’’.
		mode = mode.add(decay)
		switch t1, t2 := n.PostfixExpression.check(c, mode), n.ExpressionList.check(c, mode); {
		case isPointerType(t1) && isIntegerType(t2):
			n.typ = t1.(*PointerType).Elem()
		case isPointerType(t2) && isIntegerType(t1):
			n.typ = t2.(*PointerType).Elem()
		case isVectorType(t1) && isIntegerType(t2):
			n.typ = t1
		case isVectorType(t2) && isIntegerType(t1):
			n.typ = t2
		default:
			c.errors.add(errorf("%v: one of the expressions shall be a pointer and the other shall have integer type: %s and %s", n.Token.Position(), t1, t2))
			n.typ = c.intT
		}
	case PostfixExpressionCall: // PostfixExpression '(' ArgumentExpressionList ')'
		if isName(n.PostfixExpression) {
			mode = mode.add(implicitFuncDef)
		}
		t := n.PostfixExpression.check(c, mode.add(decay))
		n.ArgumentExpressionList.check(c, decay)
		if t == nil || mode.has(asmArgList) {
			n.typ = c.intT
			break
		}

		if t.Kind() != Ptr {
			c.errors.add(errorf("%v: expected pointer to function: %s", n.Position(), t))
			break
		}

		pt := t.(*PointerType)
		if pt.Elem().Kind() != Function {
			c.errors.add(errorf("%v: expected pointer to function: %s", n.Position(), pt))
			break
		}

		n.typ = pt.Elem().(*FunctionType).Result()
		//TODO check args
	case PostfixExpressionSelect: // PostfixExpression '.' IDENTIFIER
		nm := n.Token2.SrcStr()
		switch t := n.PostfixExpression.check(c, mode.add(decay)); t.Kind() {
		case Struct:
			st := t.(*StructType)
			f := st.FieldByName(nm)
			if f == nil {
				c.errors.add(errorf("%v: type %s has no member named %s", n.Token.Position(), st, nm))
				break
			}

			n.typ = f.Type()
		case Union:
			st := t.(*UnionType)
			f := st.FieldByName(nm)
			if f == nil {
				c.errors.add(errorf("%v: type %s has no member named %s", n.Token.Position(), st, nm))
				break
			}

			n.typ = f.Type()
		default:
			c.errors.add(errorf("%v: expected a struct or union: %s", n.Token.Position(), t))
		}
	case PostfixExpressionPSelect: // PostfixExpression "->" IDENTIFIER
		nm := n.Token2.SrcStr()
		switch t := n.PostfixExpression.check(c, mode.add(decay)); t.Kind() {
		case Ptr:
			switch et := t.(*PointerType).Elem(); et.Kind() {
			case Struct:
				st := et.(*StructType)
				f := st.FieldByName(nm)
				if f == nil {
					c.errors.add(errorf("%v: type %s has no member named %s", n.Token.Position(), st, nm))
					break
				}

				n.typ = f.Type()
			case Union:
				st := et.(*UnionType)
				f := st.FieldByName(nm)
				if f == nil {
					c.errors.add(errorf("%v: type %s has no member named %s", n.Token.Position(), st, nm))
					break
				}

				n.typ = f.Type()
			default:
				c.errors.add(errorf("%v: expected a pointer to struct or union: %s", n.Token.Position(), t))
			}
		default:
			c.errors.add(errorf("%v: expected a pointer: %s", n.Token.Position(), t))
		}
	case
		PostfixExpressionInc, // PostfixExpression "++"
		PostfixExpressionDec: // PostfixExpression "--"
		switch t := n.PostfixExpression.check(c, mode.add(decay)); {
		case
			// The operand of the postfix increment or decrement operator shall have
			// qualified or unqualified real or pointer type and shall be a modifiable
			// lvalue.
			realKinds[t.Kind()] || isPointerType(t):

			n.typ = t
			if !isModifiableLvalue(t) {
				c.errors.add(errorf("%v: operand shall be a modifiable lvalue: %s", n.PostfixExpression.Position(), t))
			}
		default:
			c.errors.add(errorf("%v: invalid operand: %s", n.PostfixExpression.Position(), t))
		}
	case PostfixExpressionComplit: // '(' TypeName ')' '{' InitializerList ',' '}'
		n.typ = n.TypeName.check(c)
		if n.InitializerList == nil {
			n.val = Zero
			break
		}

		n.InitializerList.check(c, n.Type(), 0, true)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

func isName(n Node) bool {
	for {
		switch x := n.(type) {
		case *ExpressionList:
			if x.ExpressionList != nil {
				return false
			}

			n = x.AssignmentExpression
		case *PrimaryExpression:
			return x.Case == PrimaryExpressionIdent
		default:
			return false
		}
	}
}

//  PrimaryExpression:
//          IDENTIFIER                 // Case PrimaryExpressionIdent
//  |       INTCONST                   // Case PrimaryExpressionInt
//  |       FLOATCONST                 // Case PrimaryExpressionFloat
//  |       CHARCONST                  // Case PrimaryExpressionChar
//  |       LONGCHARCONST              // Case PrimaryExpressionLChar
//  |       STRINGLITERAL              // Case PrimaryExpressionString
//  |       LONGSTRINGLITERAL          // Case PrimaryExpressionLString
//  |       '(' ExpressionList ')'     // Case PrimaryExpressionExpr
//  |       '(' CompoundStatement ')'  // Case PrimaryExpressionStmt
//  |       GenericSelection           // Case PrimaryExpressionGeneric
func (n *PrimaryExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			if !mode.has(ignoreUndefined) {
				c.errors.add(errorf("TODO %T missed/failed type check %v", n, n.Case))
			}
			return
		}

		n.typ = c.decay(r, mode)
		r = n.Type()
	}()

out:
	switch n.Case {
	case PrimaryExpressionIdent: // IDENTIFIER
		var d *Declarator
		switch x := n.resolutionScope.ident(n.Token).(type) {
		case *Declarator:
			d = x
			n.typ = d.Type()
		case *Enumerator:
			n.resolvedTo = x
			n.val = x.val
			n.typ = x.Type()
			break out
		case *Parameter:
			n.resolvedTo = x
			n.typ = x.Type()
			break out
		default:
			d = n.resolutionScope.builtin(n.Token)
			if d == nil {
				if mode.has(implicitFuncDef) {
					n.typ = c.implicitFunc
					break out
				}

				if mode.has(ignoreUndefined) {
					n.typ = c.intT
					break out
				}

				c.errors.add(errorf("%v: undefined: %s", n.Position(), n.Token.Src()))
				break out
			}
		}

		n.resolvedTo = d
		n.typ = d.Type()
	case PrimaryExpressionInt: // INTCONST
		n.val, n.typ = n.intConst(c)
	case PrimaryExpressionFloat: // FLOATCONST
		n.val, n.typ = n.floatConst(c)
	case PrimaryExpressionChar: // CHARCONST
		n.val, n.typ = n.charConst(c)
	case PrimaryExpressionLChar: // LONGCHARCONST
		n.typ = c.wcharT(n)
		n.val, n.typ = n.charConst(c)
	case PrimaryExpressionString: // STRINGLITERAL
		n.val, n.typ = n.stringConst(c)
	case PrimaryExpressionLString: // LONGSTRINGLITERAL
		n.val, n.typ = n.stringConst(c)
	case PrimaryExpressionExpr: // '(' ExpressionList ')'
		n.typ = n.ExpressionList.check(c, mode)
	case PrimaryExpressionStmt: // '(' CompoundStatement ')'
		n.typ = n.CompoundStatement.check(c)
	case PrimaryExpressionGeneric: // GenericSelection
		n.typ = n.GenericSelection.check(c, mode)
	default:
		c.errors.add(errorf("internal error: %v", n.Case))
	}
	return n.Type()
}

func (n *GenericSelection) check(c *ctx, mode flags) (r Type) {
	n.assoc, n.typ = n.GenericAssociationList.check(c, mode, n.AssignmentExpression.check(c, mode.add(decay|dontEval)))
	return n.Type()
}

func (n *GenericAssociationList) check(c *ctx, mode flags, ctrl Type) (assoc *GenericAssociation, r Type) {
	n0 := n
	var deflt *GenericAssociation
	for ; n != nil; n = n.GenericAssociationList {
		switch assoc = n.GenericAssociation; assoc.Case {
		case GenericAssociationType: // TypeName ':' AssignmentExpression
			if t := assoc.TypeName.check(c); ctrl.isCompatible(t) {
				return assoc, assoc.AssignmentExpression.check(c, decay|dontEval)
			}
		case GenericAssociationDefault: //  "default" ':' AssignmentExpression
			if deflt != nil {
				c.errors.add(errorf("%v: multiple defaults, previous at %v", assoc.Position(), deflt.Position()))
				break
			}

			assoc.AssignmentExpression.check(c, decay|dontEval)
			deflt = assoc
		default:
			c.errors.add(errorf("TODO internal error: %v", assoc.Case))
		}
	}
	if deflt != nil {
		return deflt, deflt.AssignmentExpression.Type()
	}

	c.errors.add(errorf("%v: failed to find a matching association for %s", n0.Position(), ctrl))
	return nil, Invalid
}

func (n *PrimaryExpression) floatConst(c *ctx) (v Value, t Type) {
	s0 := n.Token.SrcStr()
	s := s0
	var cplx, suff string
out:
	for i := len(s) - 1; i > 0; i-- {
		switch s0[i] {
		case 'l', 'L':
			s = s[:i]
			suff += "l"
		case 'f', 'F':
			s = s[:i]
			suff += "f"
		case 'i', 'I', 'j', 'J':
			s = s[:i]
			cplx += "i"
		case 'd', 'D':
			suff += "d"
		default:
			break out
		}
	}

	if (len(suff) > 1 || len(cplx) > 1) && suff != "dd" {
		c.errors.add(errorf("%v: invalid number format", n.Position()))
		return nil, nil
	}

	var val float64
	var err error
	prec := uint(64)
	if suff == "l" || suff == "dd" {
		prec = longDoublePrec
	}
	var bf *big.Float
	switch {
	case suff == "l" || suff == "dd" || strings.Contains(s, "p") || strings.Contains(s, "P"):
		s = s[:len(s)-len(suff)]
		bf, _, err = big.ParseFloat(strings.ToLower(s), 0, prec, big.ToNearestEven)
		if err == nil {
			val, _ = bf.Float64()
		}
	default:
		val, err = strconv.ParseFloat(s, 64)
	}
	if err != nil {
		c.errors.add(errorf("%v: %v", n.Position(), err))
		return nil, nil
	}

	// [0]6.4.4.2
	switch suff {
	case "":
		switch {
		case cplx != "":
			return Complex128Value(complex(0, val)), c.ast.kinds[ComplexDouble]
		default:
			return Float64Value(val), c.ast.kinds[Double]
		}
	case "f":
		switch {
		case cplx != "":
			return Complex64Value(complex(0, float32(val))), c.ast.kinds[ComplexFloat]
		default:
			return Float64Value(val), c.ast.kinds[Float]
		}
	case "l":
		switch {
		case cplx != "":
			return &ComplexLongDoubleValue{big.NewFloat(0), bf}, c.ast.kinds[ComplexLongDouble]
		default:
			return (*LongDoubleValue)(bf), c.ast.kinds[LongDouble]
		}
	case "dd":
		return (*LongDoubleValue)(bf), c.ast.kinds[Decimal64]
	default:
		c.errors.add(errorf("TODO %v", n.Case))
	}
	return nil, nil
}

func (n *PrimaryExpression) stringConst(c *ctx) (v Value, t Type) {
	s := stringConst(func(msg string, args ...interface{}) {
		c.errors.add(errorf(msg, args...))
	}, n.Token)
	switch n.Case {
	case PrimaryExpressionString:
		return StringValue(s), c.newArrayType(c.ast.kinds[Char], int64(len(s)), nil)
	case PrimaryExpressionLString:
		switch t = c.wcharT(n); t.Size() {
		case 2:
			v := UTF16StringValue(utf16.Encode([]rune(s)))
			return v, c.newArrayType(t, int64(len(v)), nil)
		case 4:
			v := UTF32StringValue([]rune(s))
			return v, c.newArrayType(t, int64(len(v)), nil)
		}
	default:
		c.errors.add(errorf("TODO %v", n.Case))
	}
	return n.Value(), n.Type()
}

func (n *PrimaryExpression) charConst(c *ctx) (v Value, t Type) {
	n.typ = c.intT
	switch n.Case {
	case PrimaryExpressionLChar:
		n.typ = c.wcharT(n)
		fallthrough
	case PrimaryExpressionChar:
		r := charConst(func(msg string, args ...interface{}) {
			c.errors.add(errorf(msg, args...))
		}, n.Token)
		n.val = Int64Value(r)
	default:
		c.errors.add(errorf("TODO %v", n.Case))
	}
	return n.Value(), n.Type()
}

func (n *PrimaryExpression) intConst(c *ctx) (v Value, t Type) {
	s0 := n.Token.SrcStr()
	s := strings.TrimRight(s0, "uUlL")
	prefix := 0
	var base int
	switch {
	case strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X"):
		prefix = 2
		base = 16
	case strings.HasPrefix(s, "0b") || strings.HasPrefix(s, "0B"):
		prefix = 2
		base = 2
	case strings.HasPrefix(s, "0"):
		base = 8
	default:
		base = 10
	}
	s = s[prefix:]
	val, err := strconv.ParseUint(s, base, 64)
	if err != nil {
		c.errors.add(errorf("%v: %v", n.Position(), err))
		return Unknown, c.intT
	}

	suffix := s0[prefix+len(s):]
	switch suffix = strings.ToLower(suffix); suffix {
	case "":
		if base == 10 {
			return n.intConst2(c, s0, val, Int, Long, LongLong)
		}

		return n.intConst2(c, s0, val, Int, UInt, Long, ULong, LongLong, ULongLong)
	case "u":
		return n.intConst2(c, s0, val, UInt, ULong, ULongLong)
	case "l":
		if base == 10 {
			return n.intConst2(c, s0, val, Long, LongLong)
		}

		return n.intConst2(c, s0, val, Long, ULong, LongLong, ULongLong)
	case "lu", "ul":
		return n.intConst2(c, s0, val, ULong, ULongLong)
	case "ll":
		if base == 10 {
			return n.intConst2(c, s0, val, LongLong)
		}

		return n.intConst2(c, s0, val, LongLong, ULongLong)
	case "llu", "ull":
		return n.intConst2(c, s0, val, ULongLong)
	default:
		c.errors.add(errorf("%v: invalid suffix", n.Position()))
		return Unknown, c.intT
	}
}

func (n *PrimaryExpression) intConst2(c *ctx, s string, val uint64, list ...Kind) (v Value, t Type) {
	abi := c.ast.ABI
	b := bits.Len64(val)
	for _, k := range list {
		sign := 0
		if abi.isSignedInteger(k) {
			sign = 1
		}
		if int(abi.types[k].size)*8 >= b+sign {
			switch {
			case sign == 0:
				return UInt64Value(val), c.ast.kinds[k]
			default:
				return Int64Value(val), c.ast.kinds[k]
			}

		}
	}

	c.errors.add(errorf("%v: invalid integer constant", n.Position()))
	return Unknown, c.intT
}

//  ExpressionList:
//          AssignmentExpression
//  |       ExpressionList ',' AssignmentExpression
func (n *ExpressionList) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	n0 := n
	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
		}
	}()

	for ; n != nil; n = n.ExpressionList {
		n0.typ = n.AssignmentExpression.check(c, mode)
	}
	return n0.Type()
}

//  ConstantExpression:
//          ConditionalExpression
func (n *ConstantExpression) check(c *ctx, mode flags) (r Type) {
	if n == nil {
		return Invalid
	}

	defer func() {
		if r == nil || r == Invalid {
			c.errors.add(errorf("TODO %T missed/failed type check", n))
		}
	}()

	n.typ = n.ConditionalExpression.check(c, mode)
	if n.val = n.ConditionalExpression.eval(c, mode); n.Value() == Unknown {
		c.errors.add(errorf("%v: cannot evaluate constant expression", n.Position()))
	}
	return n.Type()
}
