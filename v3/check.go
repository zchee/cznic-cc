// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"fmt" //TODO-
	"math"
	"math/big"
	"math/bits"
	"strconv"
	"strings"
)

type mode = int

const (
	// [2], 6.6 Constant expressions, 6
	//
	// An integer constant expression shall have integer type and shall
	// only have operands that are integer constants, enumeration
	// constants, character constants, sizeof expressions whose results are
	// integer constants, _Alignof expressions, and floating constants that
	// are the immediate operands of casts. Cast operators in an integer
	// constant expression shall only convert arithmetic types to integer
	// types, except as part of an operand to the sizeof or _Alignof
	// operator.
	mIntConstExpr = 1 << iota

	mIntConstExprFloat   // As mIntConstExpr plus accept floating point constants.
	mIntConstExprAnyCast // As mIntConstExpr plus accept any cast.
)

func (n *TranslationUnit) check(ctx *context) {
	for ; n != nil; n = n.TranslationUnit {
		n.ExternalDeclaration.check(ctx)
	}
}

func (n *ExternalDeclaration) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case ExternalDeclarationFuncDef: // FunctionDefinition
		n.FunctionDefinition.check(ctx)
	case ExternalDeclarationDecl: // Declaration
		n.Declaration.check(ctx, true)
	case ExternalDeclarationAsm: // AsmFunctionDefinition
		n.AsmFunctionDefinition.check(ctx)
	case ExternalDeclarationAsmStmt: // AsmStatement
		n.AsmStatement.check(ctx)
	case ExternalDeclarationEmpty: // ';'
		// nop
	case ExternalDeclarationPragma: // PragmaSTDC
		n.PragmaSTDC.check(ctx)
	default:
		panic("internal error") //TODOOK
	}
}

func (n *PragmaSTDC) check(ctx *context) {
	// nop
}

func (n *AsmFunctionDefinition) check(ctx *context) {
	if n == nil {
		return
	}

	typ := n.DeclarationSpecifiers.check(ctx)
	n.Declarator.check(ctx, n.DeclarationSpecifiers, typ, true)
	n.AsmStatement.check(ctx)
}

func (n *AsmStatement) check(ctx *context) {
	if n == nil {
		return
	}

	n.Asm.check(ctx)
	n.AttributeSpecifierList.check(ctx)
}

func (n *Declaration) check(ctx *context, tld bool) {
	if n == nil {
		return
	}

	typ := n.DeclarationSpecifiers.check(ctx)
	n.InitDeclaratorList.check(ctx, n.DeclarationSpecifiers, typ, tld)
}

func (n *InitDeclaratorList) check(ctx *context, td typeDescriptor, typ Type, tld bool) {
	for ; n != nil; n = n.InitDeclaratorList {
		n.AttributeSpecifierList.check(ctx)
		n.InitDeclarator.check(ctx, td, typ, tld)
	}
}

func (n *InitDeclarator) check(ctx *context, td typeDescriptor, typ Type, tld bool) { //TODO- tld?
	if n == nil {
		return
	}

	switch n.Case {
	case InitDeclaratorDecl: // Declarator AttributeSpecifierList
		n.Declarator.check(ctx, td, typ, tld)
		n.AttributeSpecifierList.check(ctx)
	case InitDeclaratorInit: // Declarator AttributeSpecifierList '=' Initializer
		typ := n.Declarator.check(ctx, td, typ, tld)
		n.AttributeSpecifierList.check(ctx)
		length := n.Initializer.check(ctx)
		if t := typ.underlyingType(); t.Kind() == Array && t.Len() == 0 {
			if length == 0 {
				//TODO report error
			}
			typ.setLen(length)
			break
		}

		//TODO check length
	default:
		panic("internal error") //TODOOK
	}
}

func (n *Initializer) check(ctx *context) uintptr {
	if n == nil {
		return 0
	}

	switch n.Case {
	case InitializerExpr: // AssignmentExpression
		n.AssignmentExpression.check(ctx)
		return 0 //TODO handle string literal
	case InitializerInitList: // '{' InitializerList ',' '}'
		return n.InitializerList.check(ctx)
	default:
		panic("internal error") //TODOOK
	}
}

func (n *InitializerList) check(ctx *context) (r uintptr) {
	for ; n != nil; n = n.InitializerList {
		if x := n.Designation.check(ctx); x != 0 {
			r = x
		}
		n.Initializer.check(ctx)
		r++
	}
	return r
}

func (n *Designation) check(ctx *context) uintptr {
	if n == nil {
		return 0
	}

	return n.DesignatorList.check(ctx)
}

func (n *DesignatorList) check(ctx *context) (r uintptr) {
	for ; n != nil; n = n.DesignatorList {
		if x := n.Designator.check(ctx); x > r {
			r = x
		}
	}
	return r
}

func (n *Designator) check(ctx *context) uintptr {
	if n == nil {
		return 0
	}

	switch n.Case {
	case DesignatorIndex: // '[' ConstantExpression ']'
		op := n.ConstantExpression.check(ctx, ctx.mode|mIntConstExpr)
		switch op.Value().(type) {
		// TODO
		case nil:
			//TODO report error
		}
	case DesignatorField: // '.' IDENTIFIER
		//TODO
	case DesignatorField2: // IDENTIFIER ':'
		//TODO
	default:
		panic("internal error") //TODOOK
	}
	return 0
}

func (n *AssignmentExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	//TODO check for "modifiable lvalue" in left operand
	n.Operand = noOperand //TODO-
	lop := n.UnaryExpression.check(ctx)
	if lop != noOperand && !lop.IsLValue() {
		panic("TODO") // report error
	}

	switch n.Case {
	case AssignmentExpressionCond: // ConditionalExpression
		n.Operand = n.ConditionalExpression.check(ctx)
	case AssignmentExpressionAssign: // UnaryExpression '=' AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		_ = r //TODO check assignability
		n.Operand = l.(*lvalue).Operand
	case AssignmentExpressionMul: // UnaryExpression "*=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		_ = r //TODO check assignability
		if l.Type().IsArithmeticType() {
			op, _ := usualArithmeticConversions(ctx, n, l, r)
			n.promote = op.Type()
		}
		n.Operand = &operand{typ: l.Type()}
	case AssignmentExpressionDiv: // UnaryExpression "/=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		_ = r //TODO check assignability
		if l.Type().IsArithmeticType() {
			op, _ := usualArithmeticConversions(ctx, n, l, r)
			n.promote = op.Type()
		}
		n.Operand = &operand{typ: l.Type()}
	case AssignmentExpressionMod: // UnaryExpression "%=" AssignmentExpression
		n.UnaryExpression.check(ctx)
		n.AssignmentExpression.check(ctx)
	case AssignmentExpressionAdd: // UnaryExpression "+=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		_ = r //TODO check assignability
		n.promote = n.UnaryExpression.Operand.Type()
		if l.Type().IsArithmeticType() {
			op, _ := usualArithmeticConversions(ctx, n, l, r)
			n.promote = op.Type()
		}
		n.Operand = &operand{typ: l.Type()}
	case AssignmentExpressionSub: // UnaryExpression "-=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		_ = r //TODO check assignability
		n.promote = n.UnaryExpression.Operand.Type()
		if l.Type().IsArithmeticType() {
			op, _ := usualArithmeticConversions(ctx, n, l, r)
			n.promote = op.Type()
		}
		n.Operand = &operand{typ: l.Type()}
	case AssignmentExpressionLsh: // UnaryExpression "<<=" AssignmentExpression
		n.UnaryExpression.check(ctx)
		n.AssignmentExpression.check(ctx)
	case AssignmentExpressionRsh: // UnaryExpression ">>=" AssignmentExpression
		n.UnaryExpression.check(ctx)
		n.AssignmentExpression.check(ctx)
	case AssignmentExpressionAnd: // UnaryExpression "&=" AssignmentExpression
		n.UnaryExpression.check(ctx)
		n.AssignmentExpression.check(ctx)
	case AssignmentExpressionXor: // UnaryExpression "^=" AssignmentExpression
		n.UnaryExpression.check(ctx)
		n.AssignmentExpression.check(ctx)
	case AssignmentExpressionOr: // UnaryExpression "|=" AssignmentExpression
		n.UnaryExpression.check(ctx)
		n.AssignmentExpression.check(ctx)
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *UnaryExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case UnaryExpressionPostfix: // PostfixExpression
		n.Operand = n.PostfixExpression.check(ctx)
	case UnaryExpressionInc: // "++" UnaryExpression
		n.Operand = &operand{typ: n.UnaryExpression.check(ctx).Type()}
	case UnaryExpressionDec: // "--" UnaryExpression
		n.Operand = &operand{typ: n.UnaryExpression.check(ctx).Type()}
	case UnaryExpressionAddrof: // '&' CastExpression
		ctx.not(n, mIntConstExpr)
		op := n.CastExpression.addr(ctx)
		// [0], 6.5.3.2
		//
		// The operand of the unary & operator shall be either a
		// function designator, the result of a [] or unary * operator,
		// or an lvalue that designates an object that is not a
		// bit-field and is not declared with the register
		// storage-class specifier.
		if x, ok := op.(*funcDesignator); ok {
			n.Operand = x
			break

		}
		t := op.Type()
		if _, ok := t.(*bitFieldType); ok { //TODO func isBitField(t Type) bool
			panic(n.Position().String())
			//TODO report error
			break
		}

		if d := n.CastExpression.Operand.Declarator(); d != nil && d.td.register() {
			panic(n.Position().String())
			//TODO report error
			break
		}

		switch t.Kind() {
		case Ptr:
			if v := op.Value(); v != nil {
				n.Operand = &operand{typ: t, value: v}
				break
			}

			fallthrough
		default:
			if op != noOperand && !op.IsLValue() {
				if n.CastExpression.Case == CastExpressionUnary && n.CastExpression.UnaryExpression.Case == UnaryExpressionDeref {
					n.Operand = &operand{typ: mkPtr(ctx, n, t)}
					break
				}

				panic(n.Position()) // report error
			}

			n.Operand = &operand{typ: mkPtr(ctx, n, t)}
		}
	case UnaryExpressionDeref: // '*' CastExpression
		ctx.not(n, mIntConstExpr)
		op := n.CastExpression.check(ctx)
		if op.Type().Kind() != Ptr {
			//TODO report error
			break
		}

		n.Operand = &lvalue{Operand: &operand{typ: op.Type().Elem(), value: op.Value()}}
	case UnaryExpressionPlus: // '+' CastExpression
		op := n.CastExpression.check(ctx)
		if !op.Type().IsArithmeticType() {
			//TODO report error
			break
		}

		if op.Type().IsIntegerType() {
			op = op.integerPromotion(ctx, n)
		}
		n.Operand = op
	case UnaryExpressionMinus: // '-' CastExpression
		op := n.CastExpression.check(ctx)
		if !op.Type().IsArithmeticType() {
			//TODO report error
			break
		}

		if op.Type().IsIntegerType() {
			op = op.integerPromotion(ctx, n)
		}
		if v := op.Value(); v != nil {
			op = &operand{typ: op.Type(), value: v.neg()}
		}
		n.Operand = op
	case UnaryExpressionCpl: // '~' CastExpression
		n.CastExpression.check(ctx)
		//TODO
	case UnaryExpressionNot: // '!' CastExpression
		n.CastExpression.check(ctx)
		n.Operand = &operand{typ: ctx.cfg.ABI.Type(Int)}
	case UnaryExpressionSizeofExpr: // "sizeof" UnaryExpression
		ctx.push(ctx.mode &^ mIntConstExpr)
		op := n.UnaryExpression.check(ctx)
		ctx.pop()
		n.Operand = sizeT(ctx, n.lexicalScope, n.Token, op.Type())
	case UnaryExpressionSizeofType: // "sizeof" '(' TypeName ')'
		ctx.push(ctx.mode)
		if ctx.mode&mIntConstExpr != 0 {
			ctx.mode |= mIntConstExprAnyCast
		}
		t := n.TypeName.check(ctx)
		ctx.pop()
		n.Operand = sizeT(ctx, n.lexicalScope, n.Token, t)
	case UnaryExpressionLabelAddr: // "&&" IDENTIFIER
		ctx.not(n, mIntConstExpr)
		//TODO
	case UnaryExpressionAlignofExpr: // "_Alignof" UnaryExpression
		ctx.push(ctx.mode &^ mIntConstExpr)
		n.UnaryExpression.check(ctx)
		ctx.pop()
	case UnaryExpressionAlignofType: // "_Alignof" '(' TypeName ')'
		ctx.push(ctx.mode)
		if ctx.mode&mIntConstExpr != 0 {
			ctx.mode |= mIntConstExprAnyCast
		}
		n.TypeName.check(ctx)
		ctx.pop()
	case UnaryExpressionImag: // "__imag__" UnaryExpression
		ctx.not(n, mIntConstExpr)
		n.UnaryExpression.check(ctx)
	case UnaryExpressionReal: // "__real__" UnaryExpression
		ctx.not(n, mIntConstExpr)
		n.UnaryExpression.check(ctx)
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *CastExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case CastExpressionUnary: // UnaryExpression
		n.Operand = n.UnaryExpression.addr(ctx)
	case CastExpressionCast: // '(' TypeName ')' CastExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *UnaryExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case UnaryExpressionPostfix: // PostfixExpression
		n.Operand = n.PostfixExpression.addr(ctx)
	case UnaryExpressionInc: // "++" UnaryExpression
		panic(n.Position().String())
	case UnaryExpressionDec: // "--" UnaryExpression
		panic(n.Position().String())
	case UnaryExpressionAddrof: // '&' CastExpression
		panic(n.Position().String())
	case UnaryExpressionDeref: // '*' CastExpression
		n.Operand = n.CastExpression.check(ctx)
	case UnaryExpressionPlus: // '+' CastExpression
		panic(n.Position().String())
	case UnaryExpressionMinus: // '-' CastExpression
		panic(n.Position().String())
	case UnaryExpressionCpl: // '~' CastExpression
		panic(n.Position().String())
	case UnaryExpressionNot: // '!' CastExpression
		panic(n.Position().String())
	case UnaryExpressionSizeofExpr: // "sizeof" UnaryExpression
		panic(n.Position().String())
	case UnaryExpressionSizeofType: // "sizeof" '(' TypeName ')'
		panic(n.Position().String())
	case UnaryExpressionLabelAddr: // "&&" IDENTIFIER
		panic(n.Position().String())
	case UnaryExpressionAlignofExpr: // "_Alignof" UnaryExpression
		panic(n.Position().String())
	case UnaryExpressionAlignofType: // "_Alignof" '(' TypeName ')'
		panic(n.Position().String())
	case UnaryExpressionImag: // "__imag__" UnaryExpression
		panic(n.Position().String())
	case UnaryExpressionReal: // "__real__" UnaryExpression
		//TODO
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *PostfixExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case PostfixExpressionPrimary: // PrimaryExpression
		n.Operand = n.PrimaryExpression.addr(ctx)
	case PostfixExpressionIndex: // PostfixExpression '[' Expression ']'
		n.check(ctx)
		if n.Operand.Value() != nil {
			panic("TODO")
		}
	case PostfixExpressionCall: // PostfixExpression '(' ArgumentExpressionList ')'
		panic(n.Position().String())
	case PostfixExpressionSelect: // PostfixExpression '.' IDENTIFIER
		return n.check(ctx)
	case PostfixExpressionPSelect: // PostfixExpression "->" IDENTIFIER
		op := n.PostfixExpression.check(ctx)
		t := op.Type().underlyingType()
		if k := t.Kind(); k != Ptr && k != Array && k != Invalid {
			//TODO report error
			break
		}

		st := t.Elem()
		st = st.underlyingType()
		if k := st.Kind(); k != Struct && k != Union && k != Invalid {
			//TODO report error
			break
		}

		f, ok := st.FieldByName(n.Token2.Value)
		if !ok {
			//TODO report error
			break
		}

		rop := &operand{typ: f.Type()}
		r := &lvalue{Operand: rop}
		switch x := op.Value().(type) {
		case nil:
			// nop
		case Uint64Value:
			rop.value = x + Uint64Value(f.Offset())
		default:
			panic(n.Position().String())
		}
		n.Operand = r
	case PostfixExpressionInc: // PostfixExpression "++"
		panic(n.Position().String())
	case PostfixExpressionDec: // PostfixExpression "--"
		panic(n.Position().String())
	case PostfixExpressionComplit: // '(' TypeName ')' '{' InitializerList ',' '}'
		//TODO
	case PostfixExpressionTypeCmp: // "__builtin_types_compatible_p" '(' TypeName ',' TypeName ')'
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *PrimaryExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case PrimaryExpressionIdent: // IDENTIFIER
		n.check(ctx)
		if d := n.Operand.Declarator(); d != nil {
			switch d.Type().Kind() {
			case Function:
				// nop
			default:
				n.Operand = &lvalue{Operand: &operand{typ: mkPtr(ctx, n, d.Type())}, declarator: d}
			}
		}
	case PrimaryExpressionInt: // INTCONST
		panic(n.Position().String())
	case PrimaryExpressionFloat: // FLOATCONST
		panic(n.Position().String())
	case PrimaryExpressionEnum: // ENUMCONST
		panic(n.Position().String())
	case PrimaryExpressionChar: // CHARCONST
		panic(n.Position().String())
	case PrimaryExpressionLChar: // LONGCHARCONST
		panic(n.Position().String())
	case PrimaryExpressionString: // STRINGLITERAL
		panic(n.Position().String())
	case PrimaryExpressionLString: // LONGSTRINGLITERAL
		panic(n.Position().String())
	case PrimaryExpressionExpr: // '(' Expression ')'
		n.Operand = n.Expression.addr(ctx)
	case PrimaryExpressionStmt: // '(' CompoundStatement ')'
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *Expression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case ExpressionAssign: // AssignmentExpression
		n.Operand = n.AssignmentExpression.addr(ctx)
	case ExpressionComma: // Expression ',' AssignmentExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *AssignmentExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case AssignmentExpressionCond: // ConditionalExpression
		n.Operand = n.ConditionalExpression.addr(ctx)
	case AssignmentExpressionAssign: // UnaryExpression '=' AssignmentExpression
		panic(n.Position().String())
	case AssignmentExpressionMul: // UnaryExpression "*=" AssignmentExpression
		panic(n.Position().String())
	case AssignmentExpressionDiv: // UnaryExpression "/=" AssignmentExpression
		panic(n.Position().String())
	case AssignmentExpressionMod: // UnaryExpression "%=" AssignmentExpression
		panic(n.Position().String())
	case AssignmentExpressionAdd: // UnaryExpression "+=" AssignmentExpression
		panic(n.Position().String())
	case AssignmentExpressionSub: // UnaryExpression "-=" AssignmentExpression
		panic(n.Position().String())
	case AssignmentExpressionLsh: // UnaryExpression "<<=" AssignmentExpression
		panic(n.Position().String())
	case AssignmentExpressionRsh: // UnaryExpression ">>=" AssignmentExpression
		panic(n.Position().String())
	case AssignmentExpressionAnd: // UnaryExpression "&=" AssignmentExpression
		panic(n.Position().String())
	case AssignmentExpressionXor: // UnaryExpression "^=" AssignmentExpression
		panic(n.Position().String())
	case AssignmentExpressionOr: // UnaryExpression "|=" AssignmentExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *ConditionalExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case ConditionalExpressionLOr: // LogicalOrExpression
		n.Operand = n.LogicalOrExpression.addr(ctx)
	case ConditionalExpressionCond: // LogicalOrExpression '?' Expression ':' ConditionalExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *LogicalOrExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case LogicalOrExpressionLAnd: // LogicalAndExpression
		n.Operand = n.LogicalAndExpression.addr(ctx)
	case LogicalOrExpressionLOr: // LogicalOrExpression "||" LogicalAndExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *LogicalAndExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case LogicalAndExpressionOr: // InclusiveOrExpression
		n.Operand = n.InclusiveOrExpression.addr(ctx)
	case LogicalAndExpressionLAnd: // LogicalAndExpression "&&" InclusiveOrExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *InclusiveOrExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case InclusiveOrExpressionXor: // ExclusiveOrExpression
		n.Operand = n.ExclusiveOrExpression.addr(ctx)
	case InclusiveOrExpressionOr: // InclusiveOrExpression '|' ExclusiveOrExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *ExclusiveOrExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case ExclusiveOrExpressionAnd: // AndExpression
		n.Operand = n.AndExpression.addr(ctx)
	case ExclusiveOrExpressionXor: // ExclusiveOrExpression '^' AndExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *AndExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case AndExpressionEq: // EqualityExpression
		n.Operand = n.EqualityExpression.addr(ctx)
	case AndExpressionAnd: // AndExpression '&' EqualityExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *EqualityExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case EqualityExpressionRel: // RelationalExpression
		n.Operand = n.RelationalExpression.addr(ctx)
	case EqualityExpressionEq: // EqualityExpression "==" RelationalExpression
		panic(n.Position().String())
	case EqualityExpressionNeq: // EqualityExpression "!=" RelationalExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *RelationalExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case RelationalExpressionShift: // ShiftExpression
		n.Operand = n.ShiftExpression.addr(ctx)
	case RelationalExpressionLt: // RelationalExpression '<' ShiftExpression
		panic(n.Position().String())
	case RelationalExpressionGt: // RelationalExpression '>' ShiftExpression
		panic(n.Position().String())
	case RelationalExpressionLeq: // RelationalExpression "<=" ShiftExpression
		panic(n.Position().String())
	case RelationalExpressionGeq: // RelationalExpression ">=" ShiftExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *ShiftExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case ShiftExpressionAdd: // AdditiveExpression
		n.Operand = n.AdditiveExpression.addr(ctx)
	case ShiftExpressionLsh: // ShiftExpression "<<" AdditiveExpression
		panic(n.Position().String())
	case ShiftExpressionRsh: // ShiftExpression ">>" AdditiveExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *AdditiveExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case AdditiveExpressionMul: // MultiplicativeExpression
		n.Operand = n.MultiplicativeExpression.addr(ctx)
	case AdditiveExpressionAdd: // AdditiveExpression '+' MultiplicativeExpression
		panic(n.Position().String())
	case AdditiveExpressionSub: // AdditiveExpression '-' MultiplicativeExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *MultiplicativeExpression) addr(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case MultiplicativeExpressionCast: // CastExpression
		n.Operand = n.CastExpression.addr(ctx)
	case MultiplicativeExpressionMul: // MultiplicativeExpression '*' CastExpression
		panic(n.Position().String())
	case MultiplicativeExpressionDiv: // MultiplicativeExpression '/' CastExpression
		panic(n.Position().String())
	case MultiplicativeExpressionMod: // MultiplicativeExpression '%' CastExpression
		panic(n.Position().String())
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func ssizeT(ctx *context, s Scope, tok Token, t Type) Operand { //TODO method of context?
	if t != nil && t.Incomplete() {
		//TODO report error
		return noOperand
	}

	var v Value
	if t != nil {
		v = Int64Value(t.Size())
	}
	if d := s.typedef(idSSizeT, tok); d != nil {
		return &operand{
			typ:   &aliasType{nm: idSSizeT, typ: d.Type()},
			value: v,
		}
	}

	st := ctx.ssizeT
	rank := maxRank
	if st == nil {
		abi := ctx.cfg.ABI
		for v, ok := range integerTypes {
			if !ok || !abi.isSignedInteger(Kind(v)) || intConvRank[v] >= rank || Kind(v) == Enum {
				continue
			}

			st = abi.Type(Kind(v))
			rank = intConvRank[v]
		}
		if st == nil {
			panic("internal error") //TODOOK
		}
		ctx.ssizeT = st
	}
	return &operand{typ: st, value: v}
}

func sizeT(ctx *context, s Scope, tok Token, t Type) Operand { //TODO method of context?
	if t.Incomplete() {
		//TODO report error
		return noOperand
	}

	if d := s.typedef(idSizeT, tok); d != nil {
		return &operand{
			typ:   &aliasType{nm: idSizeT, typ: d.Type()},
			value: Uint64Value(t.Size()),
		}
	}

	st := ctx.sizeT
	rank := maxRank
	if st == nil {
		abi := ctx.cfg.ABI
		for v, ok := range integerTypes {
			if !ok || abi.isSignedInteger(Kind(v)) || intConvRank[v] >= rank || Kind(v) == Enum {
				continue
			}

			st = abi.Type(Kind(v))
			rank = intConvRank[v]
		}
		if st == nil {
			panic("internal error") //TODOOK
		}
		ctx.sizeT = st
	}
	return &operand{typ: st, value: Uint64Value(t.Size())}
}

func (n *TypeName) check(ctx *context) Type {
	if n == nil {
		return noType
	}

	n.typ = n.SpecifierQualifierList.check(ctx)
	if n.AbstractDeclarator != nil {
		n.typ = n.AbstractDeclarator.check(ctx, n.typ)
	}
	return n.typ
}

func (n *AbstractDeclarator) check(ctx *context, typ Type) Type {
	if n == nil {
		return noType
	}

	n.typ = noType //TODO-
	switch n.Case {
	case AbstractDeclaratorPtr: // Pointer
		n.typ = n.Pointer.check(ctx, typ)
	case AbstractDeclaratorDecl: // Pointer DirectAbstractDeclarator
		typ = n.Pointer.check(ctx, typ)
		n.typ = n.DirectAbstractDeclarator.check(ctx, typ)
	default:
		panic("internal error") //TODOOK
	}
	return n.typ
}

func (n *DirectAbstractDeclarator) check(ctx *context, typ Type) Type {
	if n == nil {
		return noType
	}

	switch n.Case {
	case DirectAbstractDeclaratorDecl: // '(' AbstractDeclarator ')'
		n.AbstractDeclarator.check(ctx, nil)
	case DirectAbstractDeclaratorArr: // DirectAbstractDeclarator '[' TypeQualifiers AssignmentExpression ']'
		return checkArray(ctx, n, n.DirectAbstractDeclarator.check(ctx, typ), n.AssignmentExpression, true)
	case DirectAbstractDeclaratorStaticArr: // DirectAbstractDeclarator '[' "static" TypeQualifiers AssignmentExpression ']'
		return checkArray(ctx, n, n.DirectAbstractDeclarator.check(ctx, typ), n.AssignmentExpression, false)
	case DirectAbstractDeclaratorArrStatic: // DirectAbstractDeclarator '[' TypeQualifiers "static" AssignmentExpression ']'
		return checkArray(ctx, n, n.DirectAbstractDeclarator.check(ctx, typ), n.AssignmentExpression, false)
	case DirectAbstractDeclaratorArrStar: // DirectAbstractDeclarator '[' '*' ']'
		return checkArray(ctx, n, n.DirectAbstractDeclarator.check(ctx, typ), n.AssignmentExpression, false)
	case DirectAbstractDeclaratorFunc: // DirectAbstractDeclarator '(' ParameterTypeList ')'
		ft := &functionType{typeBase: typeBase{kind: byte(Function)}, result: typ}
		n.ParameterTypeList.check(ctx, ft)
		return n.DirectAbstractDeclarator.check(ctx, ft)
	default:
		panic("internal error") //TODOOK
	}
	return noType //TODO-
}

func (n *ParameterTypeList) check(ctx *context, ft *functionType) {
	if n == nil {
		return
	}

	switch n.Case {
	case ParameterTypeListList: // ParameterList
		n.ParameterList.check(ctx, ft)
	case ParameterTypeListVar: // ParameterList ',' "..."
		ft.variadic = true
		n.ParameterList.check(ctx, ft)
	default:
		panic("internal error") //TODOOK
	}
}

func (n *ParameterList) check(ctx *context, ft *functionType) {
	for ; n != nil; n = n.ParameterList {
		d := n.ParameterDeclaration.check(ctx, ft)
		if d == (*AbstractDeclarator)(nil) { //TODO-
			panic(n.Position().String())
		}
		ft.params = append(ft.params, d)
	}
}

func (n *ParameterDeclaration) check(ctx *context, ft *functionType) Object {
	if n == nil {
		return nil
	}

	switch n.Case {
	case ParameterDeclarationDecl: // DeclarationSpecifiers Declarator AttributeSpecifierList
		typ := n.DeclarationSpecifiers.check(ctx)
		if n.Declarator.check(ctx, n.DeclarationSpecifiers, typ, false).Kind() == Void {
			panic(n.Position().String())
		}
		if n.AttributeSpecifierList != nil {
			panic(n.Position().String())
		}
		n.AttributeSpecifierList.check(ctx)
		return n.Declarator
	case ParameterDeclarationAbstract: // DeclarationSpecifiers AbstractDeclarator
		typ := n.DeclarationSpecifiers.check(ctx)
		if n.AbstractDeclarator == nil {
			return &abstractObject{typ}
		}

		n.AbstractDeclarator.check(ctx, typ)
		return n.AbstractDeclarator
	default:
		panic("internal error") //TODOOK
	}
}

func (n *Pointer) check(ctx *context, typ Type) (t Type) {
	if n == nil || typ == nil {
		return typ
	}

	switch n.Case {
	case PointerTypeQual: // '*' TypeQualifiers
		n.TypeQualifiers.check(ctx, &n.typeQualifiers)
	case PointerPtr: // '*' TypeQualifiers Pointer
		n.TypeQualifiers.check(ctx, &n.typeQualifiers)
		typ = n.Pointer.check(ctx, typ)
	default:
		panic("internal error") //TODOOK
	}
	r := mkPtr(ctx, n, typ)
	if n.typeQualifiers != nil {
		r.typeQualifiers = n.typeQualifiers.check(ctx, (*DeclarationSpecifiers)(nil))
	}
	return r
}

func (n *TypeQualifiers) check(ctx *context, typ **typeBase) {
	for ; n != nil; n = n.TypeQualifiers {
		switch n.Case {
		case TypeQualifiersTypeQual: // TypeQualifier
			if *typ == nil {
				*typ = &typeBase{}
			}
			n.TypeQualifier.check(ctx, *typ)
		case TypeQualifiersAttribute: // AttributeSpecifier
			n.AttributeSpecifier.check(ctx)
		default:
			panic("internal error") //TODOOK
		}
	}
}

func (n *TypeQualifier) check(ctx *context, typ *typeBase) {
	if n == nil {
		return
	}

	switch n.Case {
	case TypeQualifierConst: // "const"
		typ.flags |= fConst
	case TypeQualifierRestrict: // "restrict"
		typ.flags |= fRestrict
	case TypeQualifierVolatile: // "volatile"
		typ.flags |= fVolatile
	case TypeQualifierAtomic: // "_Atomic"
		typ.flags |= fAtomic
	default:
		panic("internal error") //TODOOK
	}
}

func (n *SpecifierQualifierList) check(ctx *context) Type {
	n0 := n
	typ := &typeBase{}
	for ; n != nil; n = n.SpecifierQualifierList {
		switch n.Case {
		case SpecifierQualifierListTypeSpec: // TypeSpecifier SpecifierQualifierList
			n.TypeSpecifier.check(ctx, typ)
		case SpecifierQualifierListTypeQual: // TypeQualifier SpecifierQualifierList
			n.TypeQualifier.check(ctx, typ)
		case SpecifierQualifierListAlignSpec: // AlignmentSpecifier SpecifierQualifierList
			n.AlignmentSpecifier.check(ctx)
		case SpecifierQualifierListAttribute: // AttributeSpecifier SpecifierQualifierList
			n.AttributeSpecifier.check(ctx)
		default:
			panic("internal error") //TODOOK
		}
	}
	return typ.check(ctx, n0)
}

func (n *TypeSpecifier) check(ctx *context, typ *typeBase) {
	if n == nil {
		return
	}

	switch n.Case {
	case
		TypeSpecifierVoid,       // "void"
		TypeSpecifierChar,       // "char"
		TypeSpecifierShort,      // "short"
		TypeSpecifierInt,        // "int"
		TypeSpecifierInt128,     // "__int128"
		TypeSpecifierLong,       // "long"
		TypeSpecifierFloat,      // "float"
		TypeSpecifierFloat16,    // "__fp16"
		TypeSpecifierDecimal32,  // "_Decimal32"
		TypeSpecifierDecimal64,  // "_Decimal64"
		TypeSpecifierDecimal128, // "_Decimal128"
		TypeSpecifierFloat32,    // "_Float32"
		TypeSpecifierFloat32x,   // "_Float32x"
		TypeSpecifierFloat64,    // "_Float64"
		TypeSpecifierFloat64x,   // "_Float64x"
		TypeSpecifierFloat128,   // "_Float128"
		TypeSpecifierFloat80,    // "__float80"
		TypeSpecifierDouble,     // "double"
		TypeSpecifierSigned,     // "signed"
		TypeSpecifierUnsigned,   // "unsigned"
		TypeSpecifierBool,       // "_Bool"
		TypeSpecifierComplex:    // "_Complex"
		// nop
	case TypeSpecifierStruct: // StructOrUnionSpecifier
		n.StructOrUnionSpecifier.check(ctx, typ)
	case TypeSpecifierEnum: // EnumSpecifier
		n.EnumSpecifier.check(ctx)
	case TypeSpecifierTypedefName: // TYPEDEFNAME
		// nop
	case TypeSpecifierTypeofExpr: // "typeof" '(' Expression ')'
		op := n.Expression.check(ctx)
		n.typ = op.Type()
	case TypeSpecifierTypeofType: // "typeof" '(' TypeName ')'
		n.typ = n.TypeName.check(ctx)
	case TypeSpecifierAtomic: // AtomicTypeSpecifier
		n.AtomicTypeSpecifier.check(ctx)
	case
		TypeSpecifierFract, // "_Fract"
		TypeSpecifierSat,   // "_Sat"
		TypeSpecifierAccum: // "_Accum"
		// nop
	default:
		panic("TODO") //TODOOK
	}
}

func (n *AtomicTypeSpecifier) check(ctx *context) {
	if n == nil {
		return
	}

	n.TypeName.check(ctx)
}

func (n *EnumSpecifier) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case EnumSpecifierDef: // "enum" AttributeSpecifierList IDENTIFIER '{' EnumeratorList ',' '}'
		n.AttributeSpecifierList.check(ctx)
		n.EnumeratorList.check(ctx)
	case EnumSpecifierTag: // "enum" AttributeSpecifierList IDENTIFIER
		n.AttributeSpecifierList.check(ctx)
	default:
		panic("internal error") //TODOOK
	}
}

func (n *EnumeratorList) check(ctx *context) {
	var v int64
	for ; n != nil; n = n.EnumeratorList {
		n.Enumerator.check(ctx, &v)
	}
}

func (n *Enumerator) check(ctx *context, v *int64) {
	if n == nil {
		return
	}

	switch n.Case {
	case EnumeratorIdent: // IDENTIFIER AttributeSpecifierList
		n.AttributeSpecifierList.check(ctx)
		n.Operand = &operand{typ: ctx.cfg.ABI.Type(Int), value: Int64Value(*v)}
		*v++
	case EnumeratorExpr: // IDENTIFIER AttributeSpecifierList '=' ConstantExpression
		n.AttributeSpecifierList.check(ctx)
		n.Operand = n.ConstantExpression.check(ctx, ctx.mode|mIntConstExpr)
		switch x := n.Operand.Value().(type) {
		case Uint64Value:
			*v = int64(x) + 1
		case Int64Value:
			*v = int64(x) + 1
		case nil:
			//TODO report error
		default:
			panic(fmt.Sprintf("%v: %T(%v)", n.Position().String(), x, x))
		}
	default:
		panic("internal error") //TODOOK
	}
}

func (n *ConstantExpression) check(ctx *context, mode mode) Operand {
	if n == nil {
		return noOperand
	}

	ctx.push(mode)
	n.Operand = n.ConditionalExpression.check(ctx)
	ctx.pop()
	return n.Operand
}

func (n *StructOrUnionSpecifier) check(ctx *context, typ *typeBase) Type {
	if n == nil {
		return noType
	}

	switch n.Case {
	case StructOrUnionSpecifierDef: // StructOrUnion AttributeSpecifierList IDENTIFIER '{' StructDeclarationList '}'
		typ.kind = byte(n.StructOrUnion.check(ctx))
		attr := n.AttributeSpecifierList.check(ctx)
		fields := n.StructDeclarationList.check(ctx)
		m := make(map[StringID]*field, len(fields))
		for _, v := range fields {
			if v.name != 0 {
				m[v.name] = v
			}
		}
		n.typ = (&structType{
			attr:     attr,
			fields:   fields,
			m:        m,
			tag:      n.Token.Value,
			typeBase: typ,
		}).check(ctx, n)
	case StructOrUnionSpecifierTag: // StructOrUnion AttributeSpecifierList IDENTIFIER
		typ.kind = byte(n.StructOrUnion.check(ctx))
		attr := n.AttributeSpecifierList.check(ctx)
		n.typ = &taggedType{
			resolutionScope: n.lexicalScope,
			tag:             n.Token.Value,
			typeBase:        typ,
		}
		if attr != nil {
			n.typ = &attributedType{n.typ, attr}
		}
	default:
		panic("interanal error") //TODOOK
	}
	return n.typ
}

func (n *StructDeclarationList) check(ctx *context) (s []*field) {
	for ; n != nil; n = n.StructDeclarationList {
		s = append(s, n.StructDeclaration.check(ctx)...)
	}
	return s
}

func (n *StructDeclaration) check(ctx *context) (s []*field) {
	if n == nil {
		return nil
	}

	typ := n.SpecifierQualifierList.check(ctx)
	return n.StructDeclaratorList.check(ctx, n.SpecifierQualifierList, typ)
}

func (n *StructDeclaratorList) check(ctx *context, td typeDescriptor, typ Type) (s []*field) {
	for ; n != nil; n = n.StructDeclaratorList {
		s = append(s, n.StructDeclarator.check(ctx, td, typ))
	}
	return s
}

func (n *StructDeclarator) check(ctx *context, td typeDescriptor, typ Type) *field {
	if n == nil {
		return nil
	}

	sf := &field{
		typ: typ,
	}
	switch n.Case {
	case StructDeclaratorDecl: // Declarator
		sf.typ = n.Declarator.check(ctx, td, typ, false)
		sf.name = n.Declarator.Name()
	case StructDeclaratorBitField: // Declarator ':' ConstantExpression AttributeSpecifierList
		sf.isBitField = true
		sf.typ = &bitFieldType{Type: n.Declarator.check(ctx, td, typ, false)}
		sf.name = n.Declarator.Name()
		if op := n.ConstantExpression.check(ctx, ctx.mode|mIntConstExpr); op.Type().IsIntegerType() {
			switch x := op.Value().(type) {
			case Int64Value:
				if x < 0 || x > 64 {
					panic("TODO")
				}
				sf.bitFieldWidth = byte(x)
			case Uint64Value:
				if x > 64 {
					panic("TODO")
				}
				sf.bitFieldWidth = byte(x)
			default:
				//dbg("%T", x)
				panic(PrettyString(op))
			}
		} else {
			//dbg("", n.ConstantExpression)
			panic(n.Declarator.Position())
		}
		n.AttributeSpecifierList.check(ctx)
	default:
		panic("internal error") //TODOOK
	}
	return sf
}

func (n *StructOrUnion) check(ctx *context) Kind {
	if n == nil {
		return Invalid
	}

	switch n.Case {
	case StructOrUnionStruct: // "struct"
		return Struct
	case StructOrUnionUnion: // "union"
		return Union
	default:
		panic("internal error") //TODOOK
	}
}

func (n *CastExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case CastExpressionUnary: // UnaryExpression
		n.Operand = n.UnaryExpression.check(ctx)
	case CastExpressionCast: // '(' TypeName ')' CastExpression
		t := n.TypeName.check(ctx)
		ctx.push(ctx.mode)
		if m := ctx.mode; m&mIntConstExpr != 0 && m&mIntConstExprAnyCast == 0 {
			if t := n.TypeName.Type(); t != nil && t.Kind() != Int {
				ctx.mode &^= mIntConstExpr
			}
			ctx.mode |= mIntConstExprFloat
		}
		op := n.CastExpression.check(ctx)
		ctx.pop()
		n.Operand = op.convertTo(ctx, n, t)
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *PostfixExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case PostfixExpressionPrimary: // PrimaryExpression
		n.Operand = n.PrimaryExpression.check(ctx)
		//TODO- var t Type
		//TODO- if n.Operand != nil {
		//TODO- 	t = n.Operand.Type()
		//TODO- }
		// dbg("==== %v: %v %v", n.Position(), n.Case, t)
	case PostfixExpressionIndex: // PostfixExpression '[' Expression ']'
		op := n.PostfixExpression.check(ctx)
		n.Expression.check(ctx)
		t := op.Type().underlyingType()
		if k := t.Kind(); k != Array && k != Ptr && k != Invalid {
			//TODO report error
			break
		}

		n.Operand = &lvalue{Operand: &operand{typ: t.Elem().underlyingType()}}
	case PostfixExpressionCall: // PostfixExpression '(' ArgumentExpressionList ')'
		op := n.PostfixExpression.check(ctx)
		args := n.ArgumentExpressionList.check(ctx)
		n.Operand = n.checkCall(ctx, n, op.Type(), args)
	case PostfixExpressionSelect: // PostfixExpression '.' IDENTIFIER
		op := n.PostfixExpression.check(ctx)
		st := op.Type().underlyingType()
		if k := st.Kind(); k != Struct && k != Union && k != Invalid {
			//TODO report error
			break
		}

		f, ok := st.FieldByName(n.Token2.Value)
		if !ok {
			//TODO report error
			break
		}

		n.Field = f
		n.Operand = &lvalue{Operand: &operand{typ: f.Type()}}
		//TODO
	case PostfixExpressionPSelect: // PostfixExpression "->" IDENTIFIER
		op := n.PostfixExpression.check(ctx)
		t := op.Type().underlyingType()
		if k := t.Kind(); k != Ptr && k != Array && k != Invalid {
			//TODO report error
			break
		}

		st := t.Elem()
		st = st.underlyingType()
		if k := st.Kind(); k != Struct && k != Union && k != Invalid {
			//TODO report error
			break
		}

		f, ok := st.FieldByName(n.Token2.Value)
		if !ok {
			//TODO report error
			break
		}

		n.Field = f
		n.Operand = &lvalue{Operand: &operand{typ: f.Type()}}
	case PostfixExpressionInc: // PostfixExpression "++"
		n.Operand = &operand{typ: n.PostfixExpression.check(ctx).Type()}
	case PostfixExpressionDec: // PostfixExpression "--"
		n.Operand = &operand{typ: n.PostfixExpression.check(ctx).Type()}
	case PostfixExpressionComplit: // '(' TypeName ')' '{' InitializerList ',' '}'
		n.TypeName.check(ctx)
		n.InitializerList.check(ctx)
		//TODO
	case PostfixExpressionTypeCmp: // "__builtin_types_compatible_p" '(' TypeName ',' TypeName ')'
		n.TypeName.check(ctx)
		n.TypeName2.check(ctx)
		//TODO
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *PostfixExpression) checkCall(ctx *context, nd Node, f Type, args []Operand) (r Operand) {
	r = noOperand
	switch f.Kind() {
	case Invalid:
		return noOperand
	case Function:
		// ok
	case Ptr:
		if e := f.Elem(); e.Kind() == Function {
			f = e
			break
		}

		fallthrough
	default:
		//TODO report error
		return r
	}

	r = &operand{typ: f.Result()}
	params := f.Parameters()
	if len(params) == 1 && params[0].Type().Kind() == Void {
		params = nil
		if len(args) != 0 {
			//TODO report error
			return r
		}
	}

	for i, arg := range args {
		switch {
		case i < len(params):
			//TODO check assignability
			switch {
			case arg.Type().Kind() == Array:
				n.Arguments = append(n.Arguments, mkPtr(ctx, n, arg.Type().Elem()))
			default:
				n.Arguments = append(n.Arguments, params[i].Type())
			}
		default:
			n.Arguments = append(n.Arguments, defaultArgumentPromotion(ctx, nd, arg).Type())
		}
	}
	//fmt.Printf("==== %v: %p %v %v\n", n.Position(), n, len(args), len(n.Arguments))
	return r
}

func (n *ArgumentExpressionList) check(ctx *context) (r []Operand) {
	for ; n != nil; n = n.ArgumentExpressionList {
		r = append(r, n.AssignmentExpression.check(ctx))
	}
	return r
}

func (n *Expression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case ExpressionAssign: // AssignmentExpression
		n.Operand = n.AssignmentExpression.check(ctx)
	case ExpressionComma: // Expression ',' AssignmentExpression
		n.Expression.check(ctx)
		n.AssignmentExpression.check(ctx)
		//TODO
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *PrimaryExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case PrimaryExpressionIdent: // IDENTIFIER
		ctx.not(n, mIntConstExpr)
		if d := n.resolvedIn.identifier(n.Token.Value, n.Token); d != nil {
			switch t := d.Type(); t.Kind() {
			case Function:
				n.Operand = &funcDesignator{Operand: &operand{typ: t}, declarator: d}
			default:
				n.Operand = &lvalue{Operand: &operand{typ: t}, declarator: d}
			}
			break
		}

		//TODO report err
	case PrimaryExpressionInt: // INTCONST
		n.Operand = n.intConst(ctx)
	case PrimaryExpressionFloat: // FLOATCONST
		if ctx.mode&mIntConstExpr != 0 && ctx.mode&mIntConstExprFloat == 0 {
			ctx.errNode(n, "invalid integer constant expression")
			break
		}

		n.Operand = n.floatConst(ctx)
	case PrimaryExpressionEnum: // ENUMCONST
		if e := n.resolvedIn.enumerator(n.Token.Value, n.Token); e != nil {
			n.Operand = e.Operand
			break
		}

		//TODO report err
	case PrimaryExpressionChar: // CHARCONST
		s := []rune(n.Token.Value.String())
		n.Operand = &operand{typ: ctx.cfg.ABI.Type(Int), value: Int64Value(s[0])}
	case PrimaryExpressionLChar: // LONGCHARCONST
		//TODO
	case PrimaryExpressionString: // STRINGLITERAL
		ctx.not(n, mIntConstExpr)
		n.Operand = &operand{typ: mkPtr(ctx, n, &typeBase{kind: byte(Char)}), value: StringValue(n.Token.Value)} //TODO ABI singleton
	case PrimaryExpressionLString: // LONGSTRINGLITERAL
		ctx.not(n, mIntConstExpr)
		//TODO
	case PrimaryExpressionExpr: // '(' Expression ')'
		n.Operand = n.Expression.check(ctx)
	case PrimaryExpressionStmt: // '(' CompoundStatement ')'
		ctx.not(n, mIntConstExpr)
		n.CompoundStatement.check(ctx)
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *PrimaryExpression) floatConst(ctx *context) Operand {
	s0 := n.Token.String()
	s := s0
	var cplx, suff string
loop2:
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
		default:
			break loop2
		}
	}

	if len(suff) > 1 || len(cplx) > 1 {
		ctx.errNode(n, "invalid number format")
		return noOperand
	}

	var v float64
	var err error
	switch {
	case strings.Contains(s, "p") || strings.Contains(s, "P"):
		var bf *big.Float
		bf, _, err = big.ParseFloat(strings.ToLower(s), 0, 53, big.ToNearestEven)
		if err == nil {
			v, _ = bf.Float64()
		}
	default:
		v, err = strconv.ParseFloat(s, 64)
	}
	if err != nil {
		switch {
		case !strings.HasPrefix(s, "-") && strings.Contains(err.Error(), "value out of range"):
			// linux_386/usr/include/math.h
			//
			// 	/* Value returned on overflow.  With IEEE 754 floating point, this is
			// 	   +Infinity, otherwise the largest representable positive value.  */
			// 	#if __GNUC_PREREQ (3, 3)
			// 	# define HUGE_VAL (__builtin_huge_val ())
			// 	#else
			// 	/* This may provoke compiler warnings, and may not be rounded to
			// 	   +Infinity in all IEEE 754 rounding modes, but is the best that can
			// 	   be done in ISO C while remaining a constant expression.  10,000 is
			// 	   greater than the maximum (decimal) exponent for all supported
			// 	   floating-point formats and widths.  */
			// 	# define HUGE_VAL 1e10000
			// 	#endif
			v = math.Inf(1)
		default:
			ctx.errNode(n, "%v", err)
			return noOperand
		}
	}

	// [0]6.4.4.2
	switch suff {
	case "", "l":
		switch {
		case cplx != "":
			return &operand{typ: ctx.cfg.ABI.Type(Double), value: Complex128Value(complex(0, v))}
		default:
			return &operand{typ: ctx.cfg.ABI.Type(Double), value: Float64Value(v)}
		}
	case "f":
		switch {
		case cplx != "":
			return &operand{typ: ctx.cfg.ABI.Type(Double), value: Complex128Value(complex(0, float32(v)))}
		default:
			return &operand{typ: ctx.cfg.ABI.Type(Float), value: Float32Value(float32(v))}
		}
	default:
		//dbg("%q %q %q %q %v", s0, s, suff, cplx, err)
		panic("TODO")
	}
}

func (n *PrimaryExpression) intConst(ctx *context) Operand {
	var val uint64
	s0 := n.Token.String()
	s := strings.TrimRight(s0, "uUlL")
	var decadic bool
	switch {
	case strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X"):
		var err error
		if val, err = strconv.ParseUint(s[2:], 16, 64); err != nil {
			ctx.errNode(n, "%v", err)
			return nil
		}
	case strings.HasPrefix(s, "0b") || strings.HasPrefix(s, "0B"):
		var err error
		if val, err = strconv.ParseUint(s[2:], 2, 64); err != nil {
			ctx.errNode(n, "%v", err)
			return nil
		}
	case strings.HasPrefix(s, "0"):
		var err error
		if val, err = strconv.ParseUint(s, 8, 64); err != nil {
			ctx.errNode(n, "%v", err)
			return nil
		}
	default:
		decadic = true
		var err error
		if val, err = strconv.ParseUint(s, 10, 64); err != nil {
			ctx.errNode(n, "%v", err)
			return nil
		}
	}

	suffix := s0[len(s):]
	switch suffix = strings.ToLower(suffix); suffix {
	case "":
		if decadic {
			return intConst(ctx, n, s0, val, Int, Long, LongLong)
		}

		return intConst(ctx, n, s0, val, Int, UInt, Long, ULong, LongLong, ULongLong)
	case "u":
		return intConst(ctx, n, s0, val, UInt, ULong, ULongLong)
	case "l":
		if decadic {
			return intConst(ctx, n, s0, val, Long, LongLong)
		}

		return intConst(ctx, n, s0, val, Long, ULong, LongLong, ULongLong)
	case "lu", "ul":
		return intConst(ctx, n, s0, val, ULong, ULongLong)
	case "ll":
		if decadic {
			return intConst(ctx, n, s0, val, LongLong)
		}

		return intConst(ctx, n, s0, val, LongLong, ULongLong)
	case "llu", "ull":
		return intConst(ctx, n, s0, val, ULongLong)
	default:
		ctx.errNode(n, "invalid suffix: %v", s0)
		return nil
	}
}

func intConst(ctx *context, n Node, s string, val uint64, list ...Kind) Operand {
	abi := ctx.cfg.ABI
	b := bits.Len64(val)
	for _, k := range list {
		sign := 0
		if abi.isSignedInteger(k) {
			sign = 1
		}
		if abi.size(k)*8 >= b+sign {
			switch {
			case sign == 0:
				return &operand{typ: abi.Type(k), value: Uint64Value(val)}
			default:
				return &operand{typ: abi.Type(k), value: Int64Value(val)}
			}
		}
	}

	k := list[len(list)-1]
	if abi.size(k)*8 == b {
		switch {
		case abi.isSignedInteger(k):
			return &operand{typ: abi.Type(k), value: Int64Value(val)}
		default:
			return &operand{typ: abi.Type(k), value: Uint64Value(val)}
		}
	}

	// dbg("%q %v %q", s, val, list)
	ctx.errNode(n, "invalid integer constant %v", s)
	return nil
}

func (n *ConditionalExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case ConditionalExpressionLOr: // LogicalOrExpression
		n.Operand = n.LogicalOrExpression.check(ctx)
	case ConditionalExpressionCond: // LogicalOrExpression '?' Expression ':' ConditionalExpression
		op := n.LogicalOrExpression.check(ctx)
		// The first operand shall have scalar type.
		if !op.Type().IsScalarType() {
			//TODO report errpr
			break
		}

		a := flat(ctx, n, n.Expression.check(ctx))
		b := flat(ctx, n, n.ConditionalExpression.check(ctx))
		// One of the following shall hold for the second and third
		// operands:
		//TODO — both operands have the same structure or union type;
		//TODO — one operand is a pointer to an object or incomplete type and the other is a pointer to a
		//TODO qualified or unqualified version of void.
		switch {
		// — both operands have arithmetic type;
		case a.Type().IsArithmeticType() && b.Type().IsArithmeticType():
			// If both the second and third operands have
			// arithmetic type, the result type that would be
			// determined by the usual arithmetic conversions, were
			// they applied to those two operands,
			// is the type of the result.
			op, _ := usualArithmeticConversions(ctx, n, a, b)
			n.Operand = &operand{typ: op.Type()}
		// — both operands are pointers to qualified or unqualified versions of compatible types;
		case a.Type().Kind() == Ptr && b.Type().Kind() == Ptr:
			switch {
			case a.Type().Kind() == b.Type().Kind():
				n.Operand = &operand{typ: n.Expression.Operand.Type()}
			default:
				panic(fmt.Errorf("%v: internal error: %v, %v", n.Token2.Position(), a.Type(), b.Type())) //TODOOK
			}
		// — both operands have void type;
		case a.Type().Kind() == Void && b.Type().Kind() == Void:
			n.Operand = &operand{typ: a.Type()}
			// — one operand is a pointer and the other is a null pointer constant;
		case a.Type().Kind() == Ptr && b.IsZero():
			n.Operand = &operand{typ: a.Type()}
		case b.Type().Kind() == Ptr && a.IsZero():
			n.Operand = &operand{typ: b.Type()}
		default:
			//TODO panic(fmt.Errorf("%v: internal error: %v, %v", n.Token2.Position(), a.Type(), b.Type())) //TODOOK
		}
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func flat(ctx *context, n Node, op Operand) Operand {
	if op.Type().Kind() != Array {
		return op
	}

	return &operand{typ: mkPtr(ctx, n, op.Type().Elem())}
}

func (n *LogicalOrExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case LogicalOrExpressionLAnd: // LogicalAndExpression
		n.Operand = n.LogicalAndExpression.check(ctx)
	case LogicalOrExpressionLOr: // LogicalOrExpression "||" LogicalAndExpression
		n.LogicalOrExpression.check(ctx)
		n.LogicalAndExpression.check(ctx)
		n.Operand = &operand{typ: ctx.cfg.ABI.Type(Int)}
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *LogicalAndExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case LogicalAndExpressionOr: // InclusiveOrExpression
		n.Operand = n.InclusiveOrExpression.check(ctx)
	case LogicalAndExpressionLAnd: // LogicalAndExpression "&&" InclusiveOrExpression
		n.LogicalAndExpression.check(ctx)
		n.InclusiveOrExpression.check(ctx)
		n.Operand = &operand{typ: ctx.cfg.ABI.Type(Int)}
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *InclusiveOrExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case InclusiveOrExpressionXor: // ExclusiveOrExpression
		n.Operand = n.ExclusiveOrExpression.check(ctx)
	case InclusiveOrExpressionOr: // InclusiveOrExpression '|' ExclusiveOrExpression
		a := n.InclusiveOrExpression.check(ctx)
		b := n.ExclusiveOrExpression.check(ctx)
		if !a.Type().IsIntegerType() || !b.Type().IsIntegerType() {
			//TODO report err
			break
		}

		a, b = usualArithmeticConversions(ctx, &n.Token, a, b)
		if a.Value() == nil || b.Value() == nil {
			n.Operand = &operand{typ: a.Type()}
			break
		}

		n.Operand = &operand{typ: a.Type(), value: a.Value().or(b.Value())}
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *ExclusiveOrExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case ExclusiveOrExpressionAnd: // AndExpression
		n.Operand = n.AndExpression.check(ctx)
	case ExclusiveOrExpressionXor: // ExclusiveOrExpression '^' AndExpression
		a := n.ExclusiveOrExpression.check(ctx)
		b := n.AndExpression.check(ctx)
		if !a.Type().IsIntegerType() || !b.Type().IsIntegerType() {
			//TODO report err
			break
		}

		a, b = usualArithmeticConversions(ctx, &n.Token, a, b)
		if a.Value() == nil || b.Value() == nil {
			n.Operand = &operand{typ: a.Type()}
			break
		}

		n.Operand = &operand{typ: a.Type(), value: a.Value().xor(b.Value())}
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *AndExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case AndExpressionEq: // EqualityExpression
		n.Operand = n.EqualityExpression.check(ctx)
	case AndExpressionAnd: // AndExpression '&' EqualityExpression
		a := n.AndExpression.check(ctx)
		b := n.EqualityExpression.check(ctx)
		if !a.Type().IsIntegerType() || !b.Type().IsIntegerType() {
			//TODO report err
			break
		}

		a, b = usualArithmeticConversions(ctx, &n.Token, a, b)
		if a.Value() == nil || b.Value() == nil {
			n.Operand = &operand{typ: a.Type()}
			break
		}

		n.Operand = &operand{typ: a.Type(), value: a.Value().and(b.Value())}
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *EqualityExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	switch n.Case {
	case EqualityExpressionRel: // RelationalExpression
		n.Operand = n.RelationalExpression.check(ctx)
		return n.Operand
	case
		EqualityExpressionEq,  // EqualityExpression "==" RelationalExpression
		EqualityExpressionNeq: // EqualityExpression "!=" RelationalExpression

		n.Operand = &operand{typ: ctx.cfg.ABI.Type(Int)}
		lo := n.EqualityExpression.check(ctx)
		ro := n.RelationalExpression.check(ctx)
		lt := lo.Type()
		rt := ro.Type()
		switch {
		case lt.IsArithmeticType() && rt.IsArithmeticType():
			op, _ := usualArithmeticConversions(ctx, n, lo, ro)
			n.promote = op.Type()
		case lt.Kind() == Ptr && (rt.Kind() == Ptr || rt.IsIntegerType()):
			n.promote = lt
			//TODO
		case (lt.Kind() == Ptr || lt.IsIntegerType()) && rt.Kind() == Ptr:
			n.promote = rt
			//TODO
		case lt.Kind() == Ptr && rt.Kind() == Array:
			n.promote = lt
			//TODO
		case lt.Kind() == Array && rt.Kind() == Ptr:
			n.promote = rt
			//TODO
		default:
			//TODO report error
		}
		return n.Operand
	default:
		panic("internal error") //TODOOK
	}
}

func (n *RelationalExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case RelationalExpressionShift: // ShiftExpression
		n.Operand = n.ShiftExpression.check(ctx)
	case RelationalExpressionLt: // RelationalExpression '<' ShiftExpression
		n.RelationalExpression.check(ctx)
		n.ShiftExpression.check(ctx)
	case RelationalExpressionGt: // RelationalExpression '>' ShiftExpression
		n.RelationalExpression.check(ctx)
		n.ShiftExpression.check(ctx)
	case RelationalExpressionLeq: // RelationalExpression "<=" ShiftExpression
		n.RelationalExpression.check(ctx)
		n.ShiftExpression.check(ctx)
	case RelationalExpressionGeq: // RelationalExpression ">=" ShiftExpression
		n.RelationalExpression.check(ctx)
		n.ShiftExpression.check(ctx)
	default:
		panic("internal error") //TODOOK
	}
	if n.RelationalExpression != nil {
		n.Operand = &operand{typ: ctx.cfg.ABI.Type(Int)}
		lo := n.RelationalExpression.Operand
		ro := n.ShiftExpression.Operand
		lt := lo.Type()
		rt := ro.Type()
		switch {
		case lt.IsRealType() && rt.IsRealType():
			op, _ := usualArithmeticConversions(ctx, n, lo, ro)
			n.promote = op.Type()
		case lt.Kind() == Ptr && (rt.Kind() == Ptr || rt.IsIntegerType()):
			n.promote = lt
			//TODO
		case (lt.Kind() == Ptr || lt.IsIntegerType()) && rt.Kind() == Ptr:
			n.promote = rt
			//TODO
		default:
			//TODO report error
		}
	}
	return n.Operand
}

func (n *ShiftExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case ShiftExpressionAdd: // AdditiveExpression
		n.Operand = n.AdditiveExpression.check(ctx)
	case ShiftExpressionLsh: // ShiftExpression "<<" AdditiveExpression
		lop := n.ShiftExpression.check(ctx)
		rop := n.AdditiveExpression.check(ctx)
		if !lop.Type().IsIntegerType() || !rop.Type().IsIntegerType() {
			//TODO report err
			break
		}

		n.Operand = &operand{typ: lop.integerPromotion(ctx, n).Type()}
		n.promote = rop.integerPromotion(ctx, n).Type()
	case ShiftExpressionRsh: // ShiftExpression ">>" AdditiveExpression
		lop := n.ShiftExpression.check(ctx)
		rop := n.AdditiveExpression.check(ctx)
		if !lop.Type().IsIntegerType() || !rop.Type().IsIntegerType() {
			//TODO report err
			break
		}

		n.Operand = &operand{typ: lop.integerPromotion(ctx, n).Type()}
		n.promote = rop.integerPromotion(ctx, n).Type()
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *AdditiveExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case AdditiveExpressionMul: // MultiplicativeExpression
		n.Operand = n.MultiplicativeExpression.check(ctx)
	case AdditiveExpressionAdd: // AdditiveExpression '+' MultiplicativeExpression
		a := n.AdditiveExpression.check(ctx)
		b := n.MultiplicativeExpression.check(ctx)
		if !a.Type().IsArithmeticType() || !b.Type().IsArithmeticType() {
			//TODO ptr and int sum
			break
		}

		a, b = usualArithmeticConversions(ctx, &n.Token, a, b)
		if a.Value() == nil || b.Value() == nil {
			n.Operand = &operand{typ: a.Type()}
			break
		}

		n.Operand = &operand{typ: a.Type(), value: a.Value().add(b.Value())}
	case AdditiveExpressionSub: // AdditiveExpression '-' MultiplicativeExpression
		a := n.AdditiveExpression.check(ctx)
		b := n.MultiplicativeExpression.check(ctx)
		if a.Type().Kind() == Ptr && b.Type().IsScalarType() {
			n.Operand = ssizeT(ctx, n.lexicalScope, n.Token, nil)
			break
		}

		if !a.Type().IsArithmeticType() || !b.Type().IsArithmeticType() {
			//TODO report error
			break
		}

		a, b = usualArithmeticConversions(ctx, &n.Token, a, b)
		if a.Value() == nil || b.Value() == nil {
			n.Operand = &operand{typ: a.Type()}
			break
		}

		n.Operand = &operand{typ: a.Type(), value: a.Value().sub(b.Value())}
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *MultiplicativeExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case MultiplicativeExpressionCast: // CastExpression
		n.Operand = n.CastExpression.check(ctx)
	case MultiplicativeExpressionMul: // MultiplicativeExpression '*' CastExpression
		a := n.MultiplicativeExpression.check(ctx)
		b := n.CastExpression.check(ctx)
		if !a.Type().IsArithmeticType() || !b.Type().IsArithmeticType() {
			break
		}

		a, b = usualArithmeticConversions(ctx, &n.Token, a, b)
		if a.Value() == nil || b.Value() == nil {
			n.Operand = &operand{typ: a.Type()}
			break
		}

		n.Operand = &operand{typ: a.Type(), value: a.Value().mul(b.Value())}
	case MultiplicativeExpressionDiv: // MultiplicativeExpression '/' CastExpression
		a := n.MultiplicativeExpression.check(ctx)
		b := n.CastExpression.check(ctx)
		if !a.Type().IsArithmeticType() || !b.Type().IsArithmeticType() {
			break
		}

		a, b = usualArithmeticConversions(ctx, &n.Token, a, b)
		if a.Value() == nil || b.Value() == nil {
			n.Operand = &operand{typ: a.Type()}
			break
		}

		n.Operand = &operand{typ: a.Type(), value: a.Value().div(b.Value())}
	case MultiplicativeExpressionMod: // MultiplicativeExpression '%' CastExpression
		a := n.MultiplicativeExpression.check(ctx)
		b := n.CastExpression.check(ctx)
		if !a.Type().IsArithmeticType() || !b.Type().IsArithmeticType() {
			break
		}

		if !a.Type().IsIntegerType() || !b.Type().IsIntegerType() {
			ctx.errNode(&n.Token, "the operands of the %% operator shall have integer type.") // [0] 6.5.5, 2
			break
		}

		a, b = usualArithmeticConversions(ctx, &n.Token, a, b)
		if a.Value() == nil || b.Value() == nil {
			n.Operand = &operand{typ: a.Type()}
			break
		}

		n.Operand = &operand{typ: a.Type(), value: a.Value().mod(b.Value())}
	default:
		panic("internal error") //TODOOK
	}
	return n.Operand
}

func (n *Declarator) check(ctx *context, td typeDescriptor, typ Type, tld bool) Type {
	if n == nil {
		return noType
	}

	typ = n.Pointer.check(ctx, typ)
	n.td = td
	n.typ = n.DirectDeclarator.check(ctx, typ)
	n.AttributeSpecifierList.check(ctx)

	hasStorageSpecifiers := td.typedef() || td.extern() || td.static() ||
		td.auto() || td.register() || td.threadLocal()

	// 6.2.2 Linkages of identifiers

	switch {
	case tld && td.static():
		// 3: If the declaration of a file scope identifier for an object or a
		// function contains the storage-class specifier static, the identifier
		// has internal linkage.
		n.Linkage = Internal

	case td.extern():
		//TODO
		//
		// 4: For an identifier declared with the storage-class
		// specifier extern in a scope in which a prior declaration of
		// that identifier is visible, 23) if the prior declaration
		// specifies internal or external linkage, the linkage of the
		// identifier at the later declaration is the same as the
		// linkage specified at the prior declaration. If no prior
		// declaration is visible, or if the prior declaration
		// specifies no linkage, then the identifier has external
		// linkage.
		n.Linkage = External
	case
		typ.Kind() == Function && !hasStorageSpecifiers,
		tld && !hasStorageSpecifiers:

		// 5: If the declaration of an identifier for a function has no
		// storage-class specifier, its linkage is determined exactly
		// as if it were declared with the storage-class specifier
		// extern.
		n.Linkage = External
	}
	//TODO- fmt.Printf("%v: %q typ %v, kind %v, hasStorageSpecifiers %v, linkage %v\n", n.Position(), n.Name(), n.Type(), n.Type().Kind(), hasStorageSpecifiers, n.Linkage) //TODO-
	//TODO- fmt.Printf("typedef %v, extern %v, static %v, auto %v, register %v, threadLocal %v\n", td.typedef(), td.extern(), td.static(), td.auto(), td.register(), td.threadLocal())
	return n.typ
}

func (n *DirectDeclarator) check(ctx *context, typ Type) Type {
	if n == nil {
		return noType
	}

	switch n.Case {
	case DirectDeclaratorIdent: // IDENTIFIER Asm
		n.Asm.check(ctx)
		return typ
	case DirectDeclaratorDecl: // '(' AttributeSpecifierList Declarator ')'
		n.AttributeSpecifierList.check(ctx)
		return n.Declarator.check(ctx, noTypeDescriptor, typ, false)
	case DirectDeclaratorArr: // DirectDeclarator '[' TypeQualifiers AssignmentExpression ']'
		return checkArray(ctx, n, n.DirectDeclarator.check(ctx, typ), n.AssignmentExpression, true)
	case DirectDeclaratorStaticArr: // DirectDeclarator '[' "static" TypeQualifiers AssignmentExpression ']'
		return checkArray(ctx, n, n.DirectDeclarator.check(ctx, typ), n.AssignmentExpression, false)
	case DirectDeclaratorArrStatic: // DirectDeclarator '[' TypeQualifiers "static" AssignmentExpression ']'
		return checkArray(ctx, n, n.DirectDeclarator.check(ctx, typ), n.AssignmentExpression, false)
	case DirectDeclaratorStar: // DirectDeclarator '[' TypeQualifiers '*' ']'
		return checkArray(ctx, n, n.DirectDeclarator.check(ctx, typ), n.AssignmentExpression, false)
	case DirectDeclaratorFuncParam: // DirectDeclarator '(' ParameterTypeList ')'
		ft := &functionType{typeBase: typeBase{kind: byte(Function)}, result: typ}
		n.ParameterTypeList.check(ctx, ft)
		return n.DirectDeclarator.check(ctx, ft)
	case DirectDeclaratorFuncIdent: // DirectDeclarator '(' IdentifierList ')'
		ft := &functionType{typeBase: typeBase{kind: byte(Function)}, result: typ}
		n.IdentifierList.check(ctx)
		//TODO
		return n.DirectDeclarator.check(ctx, ft)
		//TODO type
	default:
		panic("internal error") //TODOOK
	}
	return noType //TODO-
}

func checkArray(ctx *context, n Node, typ Type, expr *AssignmentExpression, exprIsOptional bool) Type { //TODO pass and use typeQualifiers
	b := typ.base()
	b.align = byte(typ.Align())
	b.kind = byte(Array)
	switch {
	case expr != nil:
		op := expr.check(ctx)
		if op.Type().Kind() == Invalid {
			return noType
		}

		if !op.Type().IsIntegerType() {
			//TODO repor err
			return noType
		}

		var length uintptr
		var vla bool
		switch x := op.Value().(type) {
		case nil:
			vla = true
		case Int64Value:
			length = uintptr(x)
		case Uint64Value:
			length = uintptr(x)
		}
		switch {
		case vla:
			b.size = ctx.cfg.ABI.Types[Ptr].Size
		default:
			if typ.Incomplete() {
				//TODO report error
				return noType
			}

			b.size = length * typ.Size()
		}
		return &arrayType{typeBase: b, elem: typ, length: length, vla: vla}
	case !exprIsOptional:
		panic("TODO")
	default:
		b.flags |= fIncomplete
		return &arrayType{typeBase: b, elem: typ}
	}
}

func (n *IdentifierList) check(ctx *context) {
	for ; n != nil; n = n.IdentifierList {
		//TODO
	}
}

func (n *Asm) check(ctx *context) {
	if n == nil {
		return
	}

	n.AsmQualifierList.check(ctx)
	n.AsmArgList.check(ctx)
}

func (n *AsmArgList) check(ctx *context) {
	for ; n != nil; n = n.AsmArgList {
		n.AsmExpressionList.check(ctx)
	}
}

func (n *AsmExpressionList) check(ctx *context) {
	for ; n != nil; n = n.AsmExpressionList {
		n.AsmIndex.check(ctx)
		n.AssignmentExpression.check(ctx)
	}
}

func (n *AsmIndex) check(ctx *context) {
	if n == nil {
		return
	}

	n.Expression.check(ctx)
}

func (n *AsmQualifierList) check(ctx *context) {
	for ; n != nil; n = n.AsmQualifierList {
		n.AsmQualifier.check(ctx)
	}
}

func (n *AsmQualifier) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case AsmQualifierVolatile: // "volatile"
		//TODO
	case AsmQualifierInline: // "inline"
		//TODO
	case AsmQualifierGoto: // "goto"
		//TODO
	default:
		panic("internal error") //TODOOK
	}
}

func (n *AttributeSpecifierList) check(ctx *context) (a []*AttributeSpecifier) {
	for ; n != nil; n = n.AttributeSpecifierList {
		a = append(a, n.AttributeSpecifier.check(ctx))
	}
	return a
}

func (n *AttributeSpecifier) check(ctx *context) *AttributeSpecifier {
	if n == nil {
		return nil
	}

	n.AttributeValueList.check(ctx)
	return n
}

func (n *AttributeValueList) check(ctx *context) {
	for ; n != nil; n = n.AttributeValueList {
		n.AttributeValue.check(ctx)
	}
}

func (n *AttributeValue) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case AttributeValueIdent: // IDENTIFIER
		//TODO
	case AttributeValueExpr: // IDENTIFIER '(' ExpressionList ')'
		n.ExpressionList.check(ctx)
	default:
		panic("internal error") //TODOOK
	}
}

func (n *ExpressionList) check(ctx *context) {
	for ; n != nil; n = n.ExpressionList {
		n.AssignmentExpression.check(ctx)
	}
}

func (n *DeclarationSpecifiers) check(ctx *context) Type {
	n0 := n
	typ := &typeBase{}
	for ; n != nil; n = n.DeclarationSpecifiers {
		switch n.Case {
		case DeclarationSpecifiersStorage: // StorageClassSpecifier DeclarationSpecifiers
			n.StorageClassSpecifier.check(ctx, n)
		case DeclarationSpecifiersTypeSpec: // TypeSpecifier DeclarationSpecifiers
			n.TypeSpecifier.check(ctx, typ)
		case DeclarationSpecifiersTypeQual: // TypeQualifier DeclarationSpecifiers
			n.TypeQualifier.check(ctx, typ)
		case DeclarationSpecifiersFunc: // FunctionSpecifier DeclarationSpecifiers
			n.FunctionSpecifier.check(ctx, typ)
		case DeclarationSpecifiersAlignSpec: // AlignmentSpecifier DeclarationSpecifiers
			n.AlignmentSpecifier.check(ctx)
		case DeclarationSpecifiersAttribute: // AttributeSpecifier DeclarationSpecifiers
			n.AttributeSpecifier.check(ctx)
		default:
			panic("internal error") //TODOOK
		}
	}
	return typ.check(ctx, n0)
}

func (n *AlignmentSpecifier) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case AlignmentSpecifierAlignasType: // "_Alignas" '(' TypeName ')'
		n.TypeName.check(ctx)
	case AlignmentSpecifierAlignasExpr: // "_Alignas" '(' ConstantExpression ')'
		n.ConstantExpression.check(ctx, ctx.mode|mIntConstExpr)
	default:
		panic("internal error") //TODOOK
	}
}

func (n *FunctionSpecifier) check(ctx *context, typ *typeBase) {
	if n == nil {
		return
	}

	switch n.Case {
	case FunctionSpecifierInline: // "inline"
		typ.flags |= fInline
	case FunctionSpecifierNoreturn: // "_Noreturn"
		typ.flags |= fNoReturn
	default:
		panic("internal error") //TODOOK
	}
}

func (n *StorageClassSpecifier) check(ctx *context, ds *DeclarationSpecifiers) {
	if n == nil {
		return
	}

	switch n.Case {
	case StorageClassSpecifierTypedef: // "typedef"
		ds.class |= fTypedef
	case StorageClassSpecifierExtern: // "extern"
		ds.class |= fExtern
	case StorageClassSpecifierStatic: // "static"
		ds.class |= fStatic
	case StorageClassSpecifierAuto: // "auto"
		ds.class |= fAuto
	case StorageClassSpecifierRegister: // "register"
		ds.class |= fRegister
	case StorageClassSpecifierThreadLocal: // "_Thread_local"
		ds.class |= fThreadLocal
	default:
		panic("internal error") //TODOOK
	}
	c := bits.OnesCount(uint(ds.class & (fTypedef | fExtern | fStatic | fAuto | fRegister | fThreadLocal)))
	if c == 1 {
		return
	}

	// [2], 6.7.1, 2
	if c == 2 && ds.class&fThreadLocal != 0 {
		if ds.class&(fStatic|fExtern) != 0 {
			return
		}
	}

	ctx.errNode(n, "at most, one storage-class specifier may be given in the declaration specifiers in a declaration")
}

func (n *FunctionDefinition) check(ctx *context) {
	if n == nil {
		return
	}

	typ := n.DeclarationSpecifiers.check(ctx)
	n.Declarator.check(ctx, n.DeclarationSpecifiers, typ, true) //TODO- (why - ?)
	n.DeclarationList.check(ctx)
	n.CompoundStatement.check(ctx)
}

func (n *CompoundStatement) check(ctx *context) {
	if n == nil {
		return
	}

	n.BlockItemList.check(ctx)
}

func (n *BlockItemList) check(ctx *context) {
	for ; n != nil; n = n.BlockItemList {
		n.BlockItem.check(ctx)
	}
}

func (n *BlockItem) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case BlockItemDecl: // Declaration
		n.Declaration.check(ctx, false)
	case BlockItemStmt: // Statement
		n.Statement.check(ctx)
	case BlockItemLabel: // LabelDeclaration
		n.LabelDeclaration.check(ctx)
	case BlockItemFuncDef: // DeclarationSpecifiers Declarator CompoundStatement
		typ := n.DeclarationSpecifiers.check(ctx)
		n.Declarator.check(ctx, n.DeclarationSpecifiers, typ, false) //TODO what's the linkage of nested fns?
		n.CompoundStatement.check(ctx)
	case BlockItemPragma: // PragmaSTDC
		n.PragmaSTDC.check(ctx)
	default:
		panic("internal error") //TODOOK
	}
}

func (n *LabelDeclaration) check(ctx *context) {
	if n == nil {
		return
	}

	n.IdentifierList.check(ctx)
}

func (n *Statement) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case StatementLabeled: // LabeledStatement
		n.LabeledStatement.check(ctx)
	case StatementCompound: // CompoundStatement
		n.CompoundStatement.check(ctx)
	case StatementExpr: // ExpressionStatement
		n.ExpressionStatement.check(ctx)
	case StatementSelection: // SelectionStatement
		n.SelectionStatement.check(ctx)
	case StatementIteration: // IterationStatement
		n.IterationStatement.check(ctx)
	case StatementJump: // JumpStatement
		n.JumpStatement.check(ctx)
	case StatementAsm: // AsmStatement
		n.AsmStatement.check(ctx)
	default:
		panic("internal error") //TODOOK
	}
}

func (n *JumpStatement) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case JumpStatementGoto: // "goto" IDENTIFIER ';'
		//TODO
	case JumpStatementGotoExpr: // "goto" '*' Expression ';'
		n.Expression.check(ctx)
	case JumpStatementContinue: // "continue" ';'
		if ctx.continues <= 0 {
			panic(n.Position().String())
		}
		//TODO
	case JumpStatementBreak: // "break" ';'
		if ctx.breaks <= 0 {
			panic(n.Position().String())
		}
		//TODO
	case JumpStatementReturn: // "return" Expression ';'
		n.Expression.check(ctx)
	default:
		panic("internal error") //TODOOK
	}
}

func (n *IterationStatement) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case IterationStatementWhile: // "while" '(' Expression ')' Statement
		n.Expression.check(ctx)
		ctx.breaks++
		ctx.continues++
		n.Statement.check(ctx)
		ctx.breaks--
		ctx.continues--
	case IterationStatementDo: // "do" Statement "while" '(' Expression ')' ';'
		ctx.breaks++
		ctx.continues++
		n.Statement.check(ctx)
		ctx.breaks--
		ctx.continues--
		n.Expression.check(ctx)
	case IterationStatementFor: // "for" '(' Expression ';' Expression ';' Expression ')' Statement
		n.Expression.check(ctx)
		n.Expression2.check(ctx)
		n.Expression3.check(ctx)
		ctx.breaks++
		ctx.continues++
		n.Statement.check(ctx)
		ctx.breaks--
		ctx.continues--
	case IterationStatementForDecl: // "for" '(' Declaration Expression ';' Expression ')' Statement
		n.Declaration.check(ctx, false)
		n.Expression.check(ctx)
		n.Expression2.check(ctx)
		ctx.breaks++
		ctx.continues++
		n.Statement.check(ctx)
		ctx.breaks--
		ctx.continues--
	default:
		panic("internal error") //TODOOK
	}
}

func (n *SelectionStatement) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case SelectionStatementIf: // "if" '(' Expression ')' Statement
		n.Expression.check(ctx)
		n.Statement.check(ctx)
	case SelectionStatementIfElse: // "if" '(' Expression ')' Statement "else" Statement
		n.Expression.check(ctx)
		n.Statement.check(ctx)
		n.Statement2.check(ctx)
		if !n.Expression.Operand.Type().IsScalarType() {
			//TODO report err
			break
		}
	case SelectionStatementSwitch: // "switch" '(' Expression ')' Statement
		n.Expression.check(ctx)
		cs := ctx.cases
		ctx.cases = nil
		ctx.switches++
		ctx.breaks++
		n.Statement.check(ctx)
		ctx.breaks--
		ctx.switches--
		n.cases = ctx.cases
		ctx.cases = cs
	default:
		panic("internal error") //TODOOK
	}
}

func (n *ExpressionStatement) check(ctx *context) {
	if n == nil {
		return
	}

	n.Expression.check(ctx)
	n.AttributeSpecifierList.check(ctx)
}

func (n *LabeledStatement) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case LabeledStatementLabel: // IDENTIFIER ':' AttributeSpecifierList Statement
		n.AttributeSpecifierList.check(ctx)
		n.Statement.check(ctx)
	case LabeledStatementCaseLabel: // "case" ConstantExpression ':' Statement
		if ctx.switches <= 0 {
			//TODO report error
			break
		}

		switch n.ConstantExpression.check(ctx, ctx.mode|mIntConstExpr).Value().(type) {
		case Int64Value, Uint64Value:
			// ok
		default:
			//TODO report error
		}
		n.Statement.check(ctx)
		ctx.cases = append(ctx.cases, n)
	case LabeledStatementRange: // "case" ConstantExpression "..." ConstantExpression ':' Statement
		if ctx.switches <= 0 {
			//TODO report error
			break
		}

		switch n.ConstantExpression.check(ctx, ctx.mode|mIntConstExpr).Value().(type) {
		case Int64Value, Uint64Value:
			// ok
		default:
			//TODO report error
		}
		switch n.ConstantExpression2.check(ctx, ctx.mode|mIntConstExpr).Value().(type) {
		case Int64Value, Uint64Value:
			// ok
		default:
			//TODO report error
		}
		n.Statement.check(ctx)
		ctx.cases = append(ctx.cases, n)
	case LabeledStatementDefault: // "default" ':' Statement
		if ctx.switches <= 0 {
			//TODO report error
			break
		}

		n.Statement.check(ctx)
		ctx.cases = append(ctx.cases, n)
	default:
		panic("internal error") //TODOOK
	}
}

func (n *DeclarationList) check(ctx *context) {
	for ; n != nil; n = n.DeclarationList {
		n.Declaration.check(ctx, false)
	}
}
