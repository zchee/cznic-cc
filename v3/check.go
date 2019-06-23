// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"math"
	"math/big"
	"math/bits"
	"strconv"
	"strings"
)

type mode = int

var (
	idClosure = dict.sid("0closure") // Must be invalid indentifier.
)

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

type Parameter struct {
	d   *Declarator
	typ Type
}

func (p *Parameter) Declarator() *Declarator { return p.d }
func (p *Parameter) Name() StringID          { return p.d.Name() }
func (p *Parameter) Type() Type              { return p.typ }

func (n *TranslationUnit) check(ctx *context) {
	for n := n; n != nil; n = n.TranslationUnit {
		n.ExternalDeclaration.check(ctx)
	}
	for ; n != nil; n = n.TranslationUnit {
		n.ExternalDeclaration.checkFnBodies(ctx)
	}
}

func (n *ExternalDeclaration) checkFnBodies(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case ExternalDeclarationFuncDef: // FunctionDefinition
		n.FunctionDefinition.checkBody(ctx)
	}
}

// DeclarationSpecifiers Declarator DeclarationList CompoundStatement
func (n *FunctionDefinition) checkBody(ctx *context) {
	if n == nil {
		return
	}

	ctx.checkFn = n
	n.CompoundStatement.check(ctx)
	ctx.checkFn = nil
	for k, v := range n.Gotos {
		if _, ok := n.Labels[k]; !ok {
			ctx.errNode(v, "label %s undefined", k)
		}
	}
}

func (n *ExternalDeclaration) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case ExternalDeclarationFuncDef: // FunctionDefinition
		n.FunctionDefinition.checkDeclarator(ctx)
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
		panic(internalError())
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

func (n *InitDeclarator) check(ctx *context, td typeDescriptor, typ Type, tld bool) {
	if n == nil {
		return
	}

	if f := ctx.checkFn; f != nil {
		f.InitDeclarators = append(f.InitDeclarators, n)
	}
	switch n.Case {
	case InitDeclaratorDecl: // Declarator AttributeSpecifierList
		n.Declarator.check(ctx, td, typ, tld)
		n.AttributeSpecifierList.check(ctx)
	case InitDeclaratorInit: // Declarator AttributeSpecifierList '=' Initializer
		typ := n.Declarator.check(ctx, td, typ, tld)
		n.AttributeSpecifierList.check(ctx)
		n.Initializer.isConst = true
		length := n.Initializer.check(ctx, &n.Initializer.list, &n.Initializer.isConst, typ, 0)
		if typ.Kind() == Array && typ.Len() == 0 {
			typ.setLen(length)
		}
	default:
		panic(internalError())
	}
}

func (n *Initializer) check(ctx *context, list *[]*Initializer, isConst *bool, t Type, off uintptr) (r uintptr) {
	if n == nil {
		return 0
	}

	n.typ = t
	n.Offset = off
	switch n.Case {
	case InitializerExpr: // AssignmentExpression
		if n.AssignmentExpression.check(ctx).Value() == nil && n.AssignmentExpression.Operand.Declarator() == nil {
			*isConst = false
		}
		*list = append(*list, n)
	case InitializerInitList: // '{' InitializerList ',' '}'
		// nop
	default:
		panic(internalError())
	}

	// [0], 6.7.8 Initialization
	//
	// 11 - The initializer for a scalar shall be a single expression,
	// optionally enclosed in braces. The initial value of the object is
	// that of the expression (after conversion); the same type constraints
	// and conversions as for simple assignment apply, taking the type of
	// the scalar to be the unqualified version of its declared type.
	if t.IsScalarType() {
		return 0
	}

	// 12 - The rest of this subclause deals with initializers for objects
	// that have aggregate or union type.
	if t.Kind() == Array {
		// 14 - An array of character type may be initialized by a
		// character string literal, optionally enclosed in braces.
		// Successive characters of the character string literal
		// (including the terminating null character if there is room
		// or if the array is of unknown size) initialize the elements
		// of the array.
		if k := t.Elem().Kind(); k == Char || k == SChar || k == UChar {
			switch n.Case {
			case InitializerExpr: // AssignmentExpression
				switch x := n.AssignmentExpression.Operand.Value().(type) {
				case StringValue:
					return uintptr(len(StringID(x).String())) + 1
				}
			case InitializerInitList: // '{' InitializerList ',' '}'
				l := n.InitializerList
				if l == nil || l.Initializer.Case != InitializerExpr {
					break
				}

				switch l.Initializer.AssignmentExpression.check(ctx).Value().(type) {
				case StringValue:
					panic(internalErrorf("%v: TODO", n.Position()))
					return r
				}
			}
		}

		// 15 - An array with element type compatible with wchar_t may
		// be initialized by a wide string literal, optionally enclosed
		// in braces. Successive wide characters of the wide string
		// literal (including the terminating null wide character if
		// there is room or if the array is of unknown size) initialize
		// the elements of the array.
		if k := t.Elem().Kind(); ctx.wcharT != nil && k == ctx.wcharT.Kind() {
			switch n.Case {
			case InitializerExpr: // AssignmentExpression
				switch x := n.AssignmentExpression.Operand.Value().(type) {
				case WideStringValue:
					return uintptr(len([]rune(StringID(x).String()))) + 1
				}
			case InitializerInitList: // '{' InitializerList ',' '}'
				l := n.InitializerList
				if l == nil || l.Initializer.Case != InitializerExpr {
					break
				}

				switch l.Initializer.AssignmentExpression.check(ctx).Value().(type) {
				case WideStringValue:
					panic(internalErrorf("%v: TODO", n.Position()))
					return
				}
			}
		}

		// 16 - Otherwise, the initializer for an object that has
		// aggregate or union type shall be a brace- enclosed list of
		// initializers for the elements or named members.
		if n.Case != InitializerInitList {
			panic(internalErrorf("%v: TODO", n.Position()))
		}

		return n.InitializerList.checkArray(ctx, list, isConst, t, off)
	}

	// Struct or Union.

	// 13 - The initializer for a structure or union object that has
	// automatic storage duration shall be either an initializer list as
	// described below, or a single expression that has compatible
	// structure or union type. ...
	if n.Case == InitializerInitList {
		n.InitializerList.checkStruct(ctx, list, isConst, t, off)
	}
	return 0
}

func (n *InitializerList) checkArray(ctx *context, list *[]*Initializer, isConst *bool, t Type, off uintptr) (r uintptr) {
	elem := t.Elem()
	esz := elem.Size()
	length := t.Len()
	var i uintptr
	for ; n != nil; n = n.InitializerList {
		if n.Designation != nil {
			i, t2, off2 := n.Designation.checkArray(ctx, t)
			n.Initializer.check(ctx, list, isConst, t2, off+off2)
			i++
			if i > r {
				r = i
			}
			continue
		}

		if length != 0 && i >= length {
			panic(internalErrorf("%v: TODO", n.Position()))
		}

		n.Initializer.check(ctx, list, isConst, elem, off+i*esz)
		i++
		if i > r {
			r = i
		}
	}
	return r
}

func (n *Designation) checkArray(ctx *context, t Type) (ix uintptr, elem Type, off uintptr) {
	first := true
	for n := n.DesignatorList; n != nil; n = n.DesignatorList {
		d := n.Designator
		switch d.Case {
		case DesignatorIndex: // '[' ConstantExpression ']'
			switch x := d.ConstantExpression.check(ctx, ctx.mode|mIntConstExpr).Value().(type) {
			case Int64Value:
				if first {
					ix = uintptr(x)
				}
				elem = t.Elem()
				off += ix * elem.Size()
				t = elem
			default:
				panic(internalErrorf("%v: TODO", n.Position()))
			}
		default:
			panic(internalError())
		}
		first = false
	}
	return ix, elem, off
}

func (n *InitializerList) checkStruct(ctx *context, list *[]*Initializer, isConst *bool, t Type, off uintptr) {
	nf := t.NumField()
	union := t.Kind() == Union
	var i int
	var f Field
	for ; n != nil; n = n.InitializerList {
		if n.Designation != nil {
			off2, f := n.Designation.checkStruct(t)
			n.Initializer.check(ctx, list, isConst, f.Type(), off+off2+f.Offset())
			i++
			continue
		}

		// Skip anonymous fields. //TODO only anononymous bitfields?
		for ; ; i++ {
			if i >= nf {
				panic(internalErrorf("%v: TODO", n.Position()))
			}

			if union && i != 0 {
				panic(internalErrorf("%v: TODO", n.Position()))
			}

			f = t.FieldByIndex([]int{i})
			if f.Name() != 0 {
				break
			}
		}
		n.Initializer.check(ctx, list, isConst, f.Type(), off+f.Offset())
		i++
	}
}

func (n *Designation) checkStruct(t Type) (off uintptr, r Field) {
	for n := n.DesignatorList; n != nil; n = n.DesignatorList {
		if r != nil {
			off += r.Offset()
		}
		d := n.Designator
		var nm StringID
		switch d.Case {
		case DesignatorField: // '.' IDENTIFIER
			nm = d.Token2.Value
		case DesignatorField2: // IDENTIFIER ':'
			nm = d.Token.Value
		default:
			panic(internalError())
		}
		f, ok := t.FieldByName(nm)
		if !ok {
			panic(internalErrorf("%v: TODO", n.Position()))
		}

		r = f
		t = f.Type()
	}
	return off, r
}

func (n *InitializerList) check(ctx *context, list *[]*Initializer, isConst *bool, t Type, off uintptr) (r uintptr) {
	n2 := &Initializer{
		Case:            InitializerInitList,
		InitializerList: n,
	}
	return n2.check(ctx, list, isConst, t, off)
}

func (n *AssignmentExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	if n.Operand != nil {
		return n.Operand
	}

	//TODO check for "modifiable lvalue" in left operand
	n.Operand = noOperand
	switch n.Case {
	case AssignmentExpressionCond: // ConditionalExpression
		n.Operand = n.ConditionalExpression.check(ctx)
	case AssignmentExpressionAssign: // UnaryExpression '=' AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if d := n.UnaryExpression.Operand.Declarator(); d != nil {
			d.Write++
		}
		if !l.IsLValue() {
			//TODO ctx.errNode(n.UnaryExpression, "expected lvalue")
			break
		}

		r := n.AssignmentExpression.check(ctx)
		_ = r //TODO check assignability
		n.Operand = l.(*lvalue).Operand
	case AssignmentExpressionMul: // UnaryExpression "*=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if d := n.UnaryExpression.Operand.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		//TODO check assignability
		if l.Type().IsArithmeticType() {
			op, _ := usualArithmeticConversions(ctx, n, l, r)
			n.promote = op.Type()
		}
		n.Operand = &operand{typ: l.Type()}
	case AssignmentExpressionDiv: // UnaryExpression "/=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if d := n.UnaryExpression.Operand.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		//TODO check assignability
		if l.Type().IsArithmeticType() {
			op, _ := usualArithmeticConversions(ctx, n, l, r)
			n.promote = op.Type()
		}
		n.Operand = &operand{typ: l.Type()}
	case AssignmentExpressionMod: // UnaryExpression "%=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if d := n.UnaryExpression.Operand.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		//TODO check assignability
		if l.Type().IsArithmeticType() {
			op, _ := usualArithmeticConversions(ctx, n, l, r)
			n.promote = op.Type()
		}
		n.Operand = &operand{typ: l.Type()}
	case AssignmentExpressionAdd: // UnaryExpression "+=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if d := n.UnaryExpression.Operand.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		//TODO check assignability
		n.promote = n.UnaryExpression.Operand.Type()
		if l.Type().IsArithmeticType() {
			op, _ := usualArithmeticConversions(ctx, n, l, r)
			n.promote = op.Type()
		}
		n.Operand = &operand{typ: l.Type()}
	case AssignmentExpressionSub: // UnaryExpression "-=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if d := n.UnaryExpression.Operand.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		//TODO check assignability
		n.promote = n.UnaryExpression.Operand.Type()
		if l.Type().IsArithmeticType() {
			op, _ := usualArithmeticConversions(ctx, n, l, r)
			n.promote = op.Type()
		}
		n.Operand = &operand{typ: l.Type()}
	case AssignmentExpressionLsh: // UnaryExpression "<<=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if d := n.UnaryExpression.Operand.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		//TODO check assignability
		if !l.Type().IsIntegerType() || !r.Type().IsIntegerType() {
			//TODO report error
			break
		}

		n.promote = r.integerPromotion(ctx, n).Type()
		n.Operand = (&operand{typ: l.Type()}).integerPromotion(ctx, n)
	case AssignmentExpressionRsh: // UnaryExpression ">>=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if d := n.UnaryExpression.Operand.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		//TODO check assignability
		if !l.Type().IsIntegerType() || !r.Type().IsIntegerType() {
			//TODO report error
			break
		}

		n.promote = r.integerPromotion(ctx, n).Type()
		n.Operand = (&operand{typ: l.Type()}).integerPromotion(ctx, n)
	case AssignmentExpressionAnd: // UnaryExpression "&=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if d := n.UnaryExpression.Operand.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		//TODO check assignability
		if !l.Type().IsIntegerType() || !r.Type().IsIntegerType() {
			//TODO report error
			break
		}

		op, _ := usualArithmeticConversions(ctx, n, l, r)
		n.promote = op.Type()
		n.Operand = &operand{typ: l.Type()}
	case AssignmentExpressionXor: // UnaryExpression "^=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if d := n.UnaryExpression.Operand.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		//TODO check assignability
		if !l.Type().IsIntegerType() || !r.Type().IsIntegerType() {
			//TODO report error
			break
		}

		op, _ := usualArithmeticConversions(ctx, n, l, r)
		n.promote = op.Type()
		n.Operand = &operand{typ: l.Type()}
	case AssignmentExpressionOr: // UnaryExpression "|=" AssignmentExpression
		l := n.UnaryExpression.check(ctx)
		if d := n.UnaryExpression.Operand.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		if !l.IsLValue() {
			//TODO panic(n.Position().String()) // report error
			break
		}

		r := n.AssignmentExpression.check(ctx)
		//TODO check assignability
		if !l.Type().IsIntegerType() || !r.Type().IsIntegerType() {
			//TODO report error
			break
		}

		op, _ := usualArithmeticConversions(ctx, n, l, r)
		n.promote = op.Type()
		n.Operand = &operand{typ: l.Type()}
	default:
		panic(internalError())
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
		op := n.UnaryExpression.check(ctx)
		if d := op.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		n.Operand = &operand{typ: op.Type()}
	case UnaryExpressionDec: // "--" UnaryExpression
		op := n.UnaryExpression.check(ctx)
		if d := op.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		n.Operand = &operand{typ: op.Type()}
	case UnaryExpressionAddrof: // '&' CastExpression
		ctx.not(n, mIntConstExpr)
		op := n.CastExpression.addr(ctx)
		if op.Type().IsBitFieldType() {
			//TODO report error
			break
		}

		d := op.Declarator()
		if d != nil {
			d.AddressTaken = true
			if d.td.register() {
				//TODO report error
			}
		}

		// [0], 6.5.3.2
		//
		// The operand of the unary & operator shall be either a
		// function designator, the result of a [] or unary * operator,
		// or an lvalue that designates an object that is not a
		// bit-field and is not declared with the register
		// storage-class specifier.
		//TODO
		if x, ok := op.(*funcDesignator); ok {
			n.Operand = x
			break
		}

		n.Operand = op
	case UnaryExpressionDeref: // '*' CastExpression
		ctx.not(n, mIntConstExpr)
		op := n.CastExpression.check(ctx)
		if op.Type().Decay().Kind() != Ptr {
			//TODO report error
			break
		}

		//TODO- n.Operand = &lvalue{Operand: &operand{typ: op.Type().Elem(), value: op.Value()}}
		n.Operand = &lvalue{Operand: &operand{typ: op.Type().Elem()}}
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
		op := n.CastExpression.check(ctx)
		if !op.Type().IsIntegerType() {
			//TODO report error
			break
		}
		n.Operand = &operand{typ: op.integerPromotion(ctx, n).Type()}
	case UnaryExpressionNot: // '!' CastExpression
		n.CastExpression.check(ctx)
		n.Operand = &operand{typ: ctx.cfg.ABI.Type(Int)}
	case UnaryExpressionSizeofExpr: // "sizeof" UnaryExpression
		ctx.push(ctx.mode &^ mIntConstExpr)
		op := n.UnaryExpression.check(ctx)
		ctx.pop()
		if op.Type().Incomplete() {
			break
		}

		n.Operand = sizeT(ctx, n.lexicalScope, n.Token, uint64(op.Type().Size()))
	case UnaryExpressionSizeofType: // "sizeof" '(' TypeName ')'
		ctx.push(ctx.mode)
		if ctx.mode&mIntConstExpr != 0 {
			ctx.mode |= mIntConstExprAnyCast
		}
		t := n.TypeName.check(ctx)
		ctx.pop()
		if t.Incomplete() {
			break
		}

		n.Operand = sizeT(ctx, n.lexicalScope, n.Token, uint64(t.Size()))
	case UnaryExpressionLabelAddr: // "&&" IDENTIFIER
		ctx.not(n, mIntConstExpr)
		//TODO
	case UnaryExpressionAlignofExpr: // "_Alignof" UnaryExpression
		ctx.push(ctx.mode &^ mIntConstExpr)
		op := n.UnaryExpression.check(ctx)
		n.Operand = sizeT(ctx, n.lexicalScope, n.Token, uint64(op.Type().Align()))
		ctx.pop()
	case UnaryExpressionAlignofType: // "_Alignof" '(' TypeName ')'
		ctx.push(ctx.mode)
		if ctx.mode&mIntConstExpr != 0 {
			ctx.mode |= mIntConstExprAnyCast
		}
		t := n.TypeName.check(ctx)
		n.Operand = sizeT(ctx, n.lexicalScope, n.Token, uint64(t.Align()))
		ctx.pop()
	case UnaryExpressionImag: // "__imag__" UnaryExpression
		ctx.not(n, mIntConstExpr)
		n.UnaryExpression.check(ctx)
	case UnaryExpressionReal: // "__real__" UnaryExpression
		ctx.not(n, mIntConstExpr)
		n.UnaryExpression.check(ctx)
	default:
		panic(internalError())
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
		panic(internalError())
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
		n.Operand = n.UnaryExpression.addr(ctx)
	default:
		panic(internalError())
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
		op := n.check(ctx)
		switch op.Type().Kind() {
		case Array:
			n.Operand = op
		default:
			n.Operand = &lvalue{Operand: &operand{typ: ctx.cfg.ABI.Ptr(n, op.Type())}}
		}
	case PostfixExpressionCall: // PostfixExpression '(' ArgumentExpressionList ')'
		panic(n.Position().String())
	case PostfixExpressionSelect: // PostfixExpression '.' IDENTIFIER
		op := n.check(ctx)
		if op.Type().IsBitFieldType() {
			panic("TODO") //TODO report error
		}

		switch op.Type().Kind() {
		case Array:
			n.Operand = op
		default:
			n.Operand = &lvalue{Operand: &operand{typ: ctx.cfg.ABI.Ptr(n, op.Type())}}
		}
	case PostfixExpressionPSelect: // PostfixExpression "->" IDENTIFIER
		op := n.check(ctx)
		if op.Type().IsBitFieldType() {
			panic("TODO") //TODO report error
		}

		var v Value
		if p := n.PostfixExpression.Operand.Value(); p != nil {
			v = Uint64Value(uintptr(p.(Uint64Value)) + n.Field.Offset())
			switch op.Type().Kind() {
			case Array:
				n.Operand = &lvalue{Operand: &operand{typ: op.Type(), value: v}}
			default:
				n.Operand = &lvalue{Operand: &operand{typ: ctx.cfg.ABI.Ptr(n, op.Type()), value: v}}
			}
			break
		}

		switch op.Type().Kind() {
		case Array:
			n.Operand = op
		default:
			n.Operand = &lvalue{Operand: &operand{typ: ctx.cfg.ABI.Ptr(n, op.Type())}}
		}
	case PostfixExpressionInc: // PostfixExpression "++"
		panic(n.Position().String())
	case PostfixExpressionDec: // PostfixExpression "--"
		panic(n.Position().String())
	case PostfixExpressionComplit: // '(' TypeName ')' '{' InitializerList ',' '}'
		if f := ctx.checkFn; f != nil {
			f.CompositeLiterals = append(f.CompositeLiterals, n)
		}
		t := n.TypeName.check(ctx)
		if n.InitializerList != nil {
			n.InitializerList.isConst = true
			len := n.InitializerList.check(ctx, &n.InitializerList.list, &n.InitializerList.isConst, t, 0)
			if t.Kind() == Array && t.Incomplete() {
				t.setLen(len)
			}
		}
		n.Operand = &lvalue{Operand: &operand{typ: ctx.cfg.ABI.Ptr(n, t)}}
	case PostfixExpressionTypeCmp: // "__builtin_types_compatible_p" '(' TypeName ',' TypeName ')'
		panic(n.Position().String())
	default:
		panic(internalError())
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
				// nop //TODO ?
			default:
				d.AddressTaken = true
				n.Operand = &lvalue{Operand: &operand{typ: ctx.cfg.ABI.Ptr(n, d.Type())}, declarator: d}
			}
			return n.Operand
		}
		if ctx.cfg.RejectLateBinding && !ctx.cfg.ignoreUndefinedIdentifiers {
			ctx.errNode(n, "undefined: %s", n.Token.Value)
			return noOperand
		}

		//TODO
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
	}
	return n.Operand
}

func wcharT(ctx *context, s Scope, tok Token) Type { //TODO method of context?
	if t := ctx.wcharT; t != nil {
		return t
	}

	if d := s.typedef(idWCharT, tok); d != nil {
		t := d.Type()
		ctx.wcharT = t
		return t
	}

	panic(internalError())
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
	if st == nil {
		abi := ctx.cfg.ABI
		need := abi.size(Ptr)
		rank := maxRank
		for v, ok := range integerTypes {
			if !ok || abi.size(Kind(v)) < need || !abi.isSignedInteger(Kind(v)) || intConvRank[v] >= rank || Kind(v) == Enum {
				continue
			}

			st = abi.Type(Kind(v))
			rank = intConvRank[v]
		}
		if st == nil {
			panic(internalError())
		}
		ctx.ssizeT = st
	}
	return &operand{typ: st, value: v}
}

func sizeT(ctx *context, s Scope, tok Token, v uint64) Operand { //TODO method of context?
	if d := s.typedef(idSizeT, tok); d != nil {
		return &operand{
			typ:   &aliasType{nm: idSizeT, typ: d.Type()},
			value: Uint64Value(v),
		}
	}

	st := ctx.sizeT
	if st == nil {
		abi := ctx.cfg.ABI
		need := abi.size(Ptr)
		rank := maxRank
		for v, ok := range integerTypes {
			if !ok || abi.size(Kind(v)) < need || abi.isSignedInteger(Kind(v)) || intConvRank[v] >= rank || Kind(v) == Enum {
				continue
			}

			st = abi.Type(Kind(v))
			rank = intConvRank[v]
		}
		if st == nil {
			panic(internalError())
		}
		ctx.sizeT = st
	}
	return &operand{typ: st, value: Uint64Value(v)}
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
		return typ
	}

	n.typ = noType //TODO-
	switch n.Case {
	case AbstractDeclaratorPtr: // Pointer
		n.typ = n.Pointer.check(ctx, typ)
	case AbstractDeclaratorDecl: // Pointer DirectAbstractDeclarator
		typ = n.Pointer.check(ctx, typ)
		n.typ = n.DirectAbstractDeclarator.check(ctx, typ)
	default:
		panic(internalError())
	}
	return n.typ
}

func (n *DirectAbstractDeclarator) check(ctx *context, typ Type) Type {
	if n == nil {
		return typ
	}

	switch n.Case {
	case DirectAbstractDeclaratorDecl: // '(' AbstractDeclarator ')'
		if n.AbstractDeclarator == nil {
			// [0], 6.7.6, 128)
			//
			// As indicated by the syntax, empty parentheses in a
			// type name are interpreted as ‘‘function with no
			// parameter specification’’, rather than redundant
			// parentheses around the omitted identifier.
			panic(internalError()) //TODO
		}

		return n.AbstractDeclarator.check(ctx, typ)
	case DirectAbstractDeclaratorArr: // DirectAbstractDeclarator '[' TypeQualifiers AssignmentExpression ']'
		return n.DirectAbstractDeclarator.check(ctx, checkArray(ctx, &n.Token, typ, n.AssignmentExpression, true, false))
	case DirectAbstractDeclaratorStaticArr: // DirectAbstractDeclarator '[' "static" TypeQualifiers AssignmentExpression ']'
		return n.DirectAbstractDeclarator.check(ctx, checkArray(ctx, &n.Token, typ, n.AssignmentExpression, false, false))
	case DirectAbstractDeclaratorArrStatic: // DirectAbstractDeclarator '[' TypeQualifiers "static" AssignmentExpression ']'
		return n.DirectAbstractDeclarator.check(ctx, checkArray(ctx, &n.Token, typ, n.AssignmentExpression, false, false))
	case DirectAbstractDeclaratorArrStar: // DirectAbstractDeclarator '[' '*' ']'
		return n.DirectAbstractDeclarator.check(ctx, checkArray(ctx, &n.Token, typ, nil, true, true))
	case DirectAbstractDeclaratorFunc: // DirectAbstractDeclarator '(' ParameterTypeList ')'
		ft := &functionType{typeBase: typeBase{kind: byte(Function)}, result: typ}
		n.ParameterTypeList.check(ctx, ft)
		return n.DirectAbstractDeclarator.check(ctx, ft)
	default:
		panic(internalError())
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
		panic(internalError())
	}
}

func (n *ParameterList) check(ctx *context, ft *functionType) {
	for ; n != nil; n = n.ParameterList {
		p := n.ParameterDeclaration.check(ctx, ft)
		ft.params = append(ft.params, p)
	}
}

func (n *ParameterDeclaration) check(ctx *context, ft *functionType) *Parameter {
	if n == nil {
		return nil
	}

	switch n.Case {
	case ParameterDeclarationDecl: // DeclarationSpecifiers Declarator AttributeSpecifierList
		typ := n.DeclarationSpecifiers.check(ctx)
		if n.typ = n.Declarator.check(ctx, n.DeclarationSpecifiers, typ, false); n.typ.Kind() == Void {
			panic(n.Position().String())
		}
		if n.AttributeSpecifierList != nil {
			panic(n.Position().String())
		}
		n.AttributeSpecifierList.check(ctx)
		return &Parameter{d: n.Declarator, typ: n.Declarator.Type()}
	case ParameterDeclarationAbstract: // DeclarationSpecifiers AbstractDeclarator
		n.typ = n.DeclarationSpecifiers.check(ctx)
		if n.AbstractDeclarator != nil {
			n.typ = n.AbstractDeclarator.check(ctx, n.typ)
		}
		return &Parameter{typ: n.typ}
	default:
		panic(internalError())
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
		panic(internalError())
	}
	r := ctx.cfg.ABI.Ptr(n, typ).(*pointerType)
	if n.typeQualifiers != nil {
		r.typeQualifiers = n.typeQualifiers.check(ctx, (*DeclarationSpecifiers)(nil), false)
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
			panic(internalError())
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
		panic(internalError())
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
			panic(internalError())
		}
	}
	return typ.check(ctx, n0, true)
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
	case TypeSpecifierStructOrUnion: // StructOrUnionSpecifier
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
		panic(internalError())
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
		panic(internalError())
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
			panic(internalError())
		}
	default:
		panic(internalError())
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
		panic(internalError())
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

	if n.Declarator != nil {
		typ = n.Declarator.check(ctx, td, typ, false)
	}
	sf := &field{
		typ: typ,
	}
	switch n.Case {
	case StructDeclaratorDecl: // Declarator
		sf.name = n.Declarator.Name()
	case StructDeclaratorBitField: // Declarator ':' ConstantExpression AttributeSpecifierList
		sf.isBitField = true
		sf.typ = &bitFieldType{Type: typ, field: sf}
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
	case PostfixExpressionIndex: // PostfixExpression '[' Expression ']'
		n.PostfixExpression.check(ctx)
		n.Expression.check(ctx)
		switch t := n.PostfixExpression.Operand.Type(); {
		case t.Kind() == Array, t.Kind() == Ptr:
			if !n.Expression.Operand.Type().IsIntegerType() {
				//TODO report error
			}
			n.Operand = &lvalue{Operand: &operand{typ: t.Elem()}}
		case t.Kind() == Invalid:
			// nop
		case t.IsIntegerType():
			//TODO
		default:
			//TODO report error
		}
	case PostfixExpressionCall: // PostfixExpression '(' ArgumentExpressionList ')'
		op := n.PostfixExpression.check(ctx)
		args := n.ArgumentExpressionList.check(ctx)
		n.Operand = n.checkCall(ctx, n, op.Type(), args)
	case PostfixExpressionSelect: // PostfixExpression '.' IDENTIFIER
		op := n.PostfixExpression.check(ctx)
		st := op.Type()
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
		ft := f.Type()
		if f.IsBitField() {
			ft = &bitFieldType{Type: ft, field: f.(*field)}
		}
		n.Operand = &lvalue{Operand: &operand{typ: ft}}
	case PostfixExpressionPSelect: // PostfixExpression "->" IDENTIFIER
		op := n.PostfixExpression.check(ctx)
		t := op.Type()
		if k := t.Decay().Kind(); k != Ptr && k != Invalid {
			//TODO report error
			break
		}

		st := t.Elem()
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
		ft := f.Type()
		if f.IsBitField() {
			ft = &bitFieldType{Type: ft, field: f.(*field)}
		}
		n.Operand = &lvalue{Operand: &operand{typ: ft}}
	case PostfixExpressionInc: // PostfixExpression "++"
		op := n.PostfixExpression.check(ctx)
		if d := op.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		n.Operand = &operand{typ: op.Type()}
	case PostfixExpressionDec: // PostfixExpression "--"
		op := n.PostfixExpression.check(ctx)
		if d := op.Declarator(); d != nil {
			d.Read++
			d.Write++
		}
		n.Operand = &operand{typ: op.Type()}
	case PostfixExpressionComplit: // '(' TypeName ')' '{' InitializerList ',' '}'
		if f := ctx.checkFn; f != nil {
			f.CompositeLiterals = append(f.CompositeLiterals, n)
		}
		t := n.TypeName.check(ctx)
		if n.InitializerList != nil {
			n.InitializerList.isConst = true
			len := n.InitializerList.check(ctx, &n.InitializerList.list, &n.InitializerList.isConst, t, 0)
			if t.Kind() == Array && t.Incomplete() {
				t.setLen(len)
			}
		}
		n.Operand = &lvalue{Operand: &operand{typ: t}}
	case PostfixExpressionTypeCmp: // "__builtin_types_compatible_p" '(' TypeName ',' TypeName ')'
		n.TypeName.check(ctx)
		n.TypeName2.check(ctx)
		//TODO
	default:
		panic(internalError())
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
			n.Arguments = append(n.Arguments, params[i].Type().Decay())
		default:
			n.Arguments = append(n.Arguments, defaultArgumentPromotion(ctx, nd, arg).Type())
		}
	}
	return r
}

func defaultArgumentPromotion(ctx *context, n Node, op Operand) Operand {
	t := op.Type().Decay()
	if arithmeticTypes[t.Kind()] {
		if t.IsIntegerType() {
			return op.integerPromotion(ctx, n)
		}

		switch t.Kind() {
		case Float:
			return op.convertTo(ctx, n, ctx.cfg.ABI.Type(Double))
		}
	}
	return op
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
		n.Operand = n.AssignmentExpression.check(ctx)
	default:
		panic(internalError())
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
		return n.checkIdentifier(ctx)
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
		s := []rune(n.Token.Value.String())
		n.Operand = &operand{typ: wcharT(ctx, n.lexicalScope, n.Token), value: Int64Value(s[0])}
	case PrimaryExpressionString: // STRINGLITERAL
		ctx.not(n, mIntConstExpr)
		typ := ctx.cfg.ABI.Type(Char)
		b := typ.base()
		b.align = byte(typ.Align())
		b.fieldAlign = byte(typ.FieldAlign())
		b.kind = byte(Array)
		sz := uintptr(len(n.Token.Value.String())) + 1 //TODO set sz in cpp
		arr := &arrayType{typeBase: b, decay: ctx.cfg.ABI.Ptr(n, typ), elem: typ, length: sz}
		arr.setLen(sz)
		n.Operand = &operand{typ: arr, value: StringValue(n.Token.Value)}
	case PrimaryExpressionLString: // LONGSTRINGLITERAL
		ctx.not(n, mIntConstExpr)
		typ := wcharT(ctx, n.lexicalScope, n.Token)
		b := typ.base()
		b.align = byte(typ.Align())
		b.fieldAlign = byte(typ.FieldAlign())
		b.kind = byte(Array)
		sz := uintptr(len([]rune(n.Token.Value.String()))) + 1 //TODO set sz in cpp
		arr := &arrayType{typeBase: b, decay: ctx.cfg.ABI.Ptr(n, typ), elem: typ, length: sz}
		arr.setLen(sz)
		n.Operand = &operand{typ: arr, value: WideStringValue(n.Token.Value)}
	case PrimaryExpressionExpr: // '(' Expression ')'
		n.Operand = n.Expression.check(ctx)
	case PrimaryExpressionStmt: // '(' CompoundStatement ')'
		ctx.not(n, mIntConstExpr)
		n.Operand = n.CompoundStatement.check(ctx)
		if n.Operand == noOperand {
			n.Operand = &operand{typ: ctx.cfg.ABI.Type(Void)}
		}
	default:
		panic(internalError())
	}
	return n.Operand
}

func (n *PrimaryExpression) checkIdentifier(ctx *context) Operand {
	ctx.not(n, mIntConstExpr)
	var d *Declarator
	if n.resolvedIn == nil {
		if ctx.cfg.RejectLateBinding && !ctx.cfg.ignoreUndefinedIdentifiers {
			ctx.errNode(n, "undefined: %s", n.Token.Value)
			return noOperand
		}

		nm := n.Token.Value
	out:
		for s := n.lexicalScope; s != nil; s = s.Parent() { //TODO use s.declarator()
			for _, v := range s[nm] {
				switch x := v.(type) {
				case *Enumerator:
					break out
				case *Declarator:
					if x.IsTypedefName {
						break out
					}

					n.resolvedIn = s
					d = x
					break out
				case *EnumSpecifier, *StructOrUnionSpecifier, *StructDeclarator, *LabeledStatement:
					// nop
				default:
					panic(internalError())
				}
			}
		}
	}
	if d == nil {
		d = n.resolvedIn.identifier(n.Token.Value, n.Token) //TODO use .declarator()
	}
	if d == nil {
		if !ctx.cfg.ignoreUndefinedIdentifiers {
			ctx.errNode(n, "undefined: %s", n.Token.Value)
		}
		return noOperand
	}

	switch d.Linkage {
	case Internal:
		if d.IsStatic() {
			break
		}

		fallthrough
	case External:
		s := n.resolvedIn
		if s.Parent() == nil {
			break
		}

		for s.Parent() != nil {
			s = s.Parent()
		}

		if d2 := s.declarator(n.Token.Value, Token{}); d2 != nil {
			d = d2
		}
	}

	switch t := d.Type(); t.Kind() {
	case Function:
		n.Operand = &funcDesignator{Operand: &operand{typ: t}, declarator: d}
	default:
		d.Read++
		n.Operand = &lvalue{Operand: &operand{typ: t}, declarator: d}
	}
	if ctx.closure == nil {
		return n.Operand
	}

	for s := n.lexicalScope; s != nil; s = s.Parent() {
		if _, ok := s[idClosure]; !ok {
			continue
		}

		if ctx.closure == nil {
			ctx.closure = map[StringID]struct{}{}
		}
		ctx.closure[d.Name()] = struct{}{}
		return n.Operand
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
			return &operand{typ: ctx.cfg.ABI.Type(ComplexDouble), value: Complex128Value(complex(0, v))}
		default:
			return &operand{typ: ctx.cfg.ABI.Type(Double), value: Float64Value(v)}
		}
	case "f":
		switch {
		case cplx != "":
			return &operand{typ: ctx.cfg.ABI.Type(ComplexFloat), value: Complex64Value(complex(0, float32(v)))}
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

		a := n.Expression.check(ctx)
		b := n.ConditionalExpression.check(ctx)

		var val Value
		if op.Value() != nil && a.Value() != nil && b.Value() != nil {
			switch {
			case op.IsZero():
				val = b.Value()
			default:
				val = a.Value()
			}
		}

		if a.Type().Kind() == Invalid && b.Type().Kind() == Invalid {
			return noOperand
		}

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
			n.Operand = &operand{typ: op.Type(), value: val}
		// — both operands are pointers to qualified or unqualified versions of compatible types;
		case a.Type().Decay().Kind() == Ptr && b.Type().Decay().Kind() == Ptr:
			//TODO check compatible
			n.Operand = &operand{typ: n.Expression.Operand.Type(), value: val}
		// — both operands have void type;
		case a.Type().Kind() == Void && b.Type().Kind() == Void:
			n.Operand = &operand{typ: a.Type(), value: val}
		// — one operand is a pointer and the other is a null pointer constant;
		case (a.Type().Kind() == Ptr || a.Type().Kind() == Function) && b.IsZero():
			n.Operand = &operand{typ: a.Type(), value: val}
		case (b.Type().Kind() == Ptr || b.Type().Kind() == Function) && a.IsZero():
			n.Operand = &operand{typ: b.Type(), value: val}
		case a.Type().Kind() == Ptr && a.Type().Elem().Kind() == Function && b.Type().Kind() == Function:
			//TODO check compatible
			n.Operand = &operand{typ: a.Type(), value: val}
		case b.Type().Kind() == Ptr && b.Type().Elem().Kind() == Function && a.Type().Kind() == Function:
			//TODO check compatible
			n.Operand = &operand{typ: b.Type(), value: val}
		case a.Type().Kind() != Invalid:
			n.Operand = &operand{typ: a.Type(), value: val}
		case b.Type().Kind() != Invalid:
			n.Operand = &operand{typ: b.Type(), value: val}
		default:
			panic(internalError())
		}
	default:
		panic(internalError())
	}
	return n.Operand
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
		lop := n.LogicalOrExpression.check(ctx)
		rop := n.LogicalAndExpression.check(ctx)
		var v Value
		if lop.Value() != nil && rop.Value() != nil {
			switch {
			case n.LogicalOrExpression.Operand.IsNonZero() || n.LogicalAndExpression.Operand.IsNonZero():
				v = Int64Value(1)
			default:
				v = Int64Value(0)
			}
		}
		n.Operand = &operand{typ: ctx.cfg.ABI.Type(Int), value: v}
	default:
		panic(internalError())
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
		lop := n.LogicalAndExpression.check(ctx)
		rop := n.InclusiveOrExpression.check(ctx)
		var v Value
		if lop.Value() != nil && rop.Value() != nil {
			switch {
			case n.LogicalAndExpression.Operand.IsNonZero() && n.InclusiveOrExpression.Operand.IsNonZero():
				v = Int64Value(1)
			default:
				v = Int64Value(0)
			}
		}
		n.Operand = &operand{typ: ctx.cfg.ABI.Type(Int), value: v}
	default:
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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
	case
		EqualityExpressionEq,  // EqualityExpression "==" RelationalExpression
		EqualityExpressionNeq: // EqualityExpression "!=" RelationalExpression

		op := &operand{typ: ctx.cfg.ABI.Type(Int)}
		n.Operand = op
		lo := n.EqualityExpression.check(ctx)
		ro := n.RelationalExpression.check(ctx)
		lt := lo.Type().Decay()
		rt := ro.Type().Decay()
		n.promote = noType
		ok := false
		switch {
		case lt.IsArithmeticType() && rt.IsArithmeticType():
			op, _ := usualArithmeticConversions(ctx, n, lo, ro)
			n.promote = op.Type()
			ok = true
		case lt.Kind() == Ptr && (rt.Kind() == Ptr || rt.IsIntegerType()):
			n.promote = lt
			//TODO
		case (lt.Kind() == Ptr || lt.IsIntegerType()) && rt.Kind() == Ptr:
			n.promote = rt
			//TODO
		case lt.Kind() == Function || rt.Kind() == Function:
			n.promote = ctx.cfg.ABI.Type(Ptr)
		default:
			//TODO report error
		}
		if n.promote.Kind() == Invalid || !ok {
			break
		}

		lo = lo.convertTo(ctx, n, n.promote)
		ro = ro.convertTo(ctx, n, n.promote)
		if a, b := lo.Value(), ro.Value(); a != nil && b != nil {
			switch n.Case {
			case EqualityExpressionEq: // EqualityExpression "==" RelationalExpression
				op.value = a.eq(b)
			case EqualityExpressionNeq: // EqualityExpression "!=" RelationalExpression
				op.value = a.neq(b)
			}
		}
	default:
		panic(internalError())
	}
	return n.Operand
}

func (n *RelationalExpression) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.Operand = noOperand //TODO-
	switch n.Case {
	case RelationalExpressionShift: // ShiftExpression
		n.Operand = n.ShiftExpression.check(ctx)
	case
		RelationalExpressionLt,  // RelationalExpression '<' ShiftExpression
		RelationalExpressionGt,  // RelationalExpression '>' ShiftExpression
		RelationalExpressionLeq, // RelationalExpression "<=" ShiftExpression
		RelationalExpressionGeq: // RelationalExpression ">=" ShiftExpression

		op := &operand{typ: ctx.cfg.ABI.Type(Int)}
		n.Operand = op
		lo := n.RelationalExpression.check(ctx)
		ro := n.ShiftExpression.check(ctx)
		lt := lo.Type().Decay()
		rt := ro.Type().Decay()
		n.promote = noType
		ok := false
		switch {
		case lt.IsRealType() && rt.IsRealType():
			op, _ := usualArithmeticConversions(ctx, n, lo, ro)
			n.promote = op.Type()
			ok = true
		case lt.Kind() == Ptr && (rt.Kind() == Ptr || rt.IsIntegerType()):
			n.promote = lt
			//TODO
		case (lt.Kind() == Ptr || lt.IsIntegerType()) && rt.Kind() == Ptr:
			n.promote = rt
			//TODO
		default:
			//TODO report error
		}
		if n.promote.Kind() == Invalid || !ok {
			break
		}

		lo = lo.convertTo(ctx, n, n.promote)
		ro = ro.convertTo(ctx, n, n.promote)
		if a, b := lo.Value(), ro.Value(); a != nil && b != nil {
			switch n.Case {
			case RelationalExpressionLt: // RelationalExpression '<' ShiftExpression
				op.value = a.lt(b)
			case RelationalExpressionGt: // RelationalExpression '>' ShiftExpression
				op.value = a.gt(b)
			case RelationalExpressionLeq: // RelationalExpression "<=" ShiftExpression
				op.value = a.le(b)
			case RelationalExpressionGeq: // RelationalExpression ">=" ShiftExpression
				op.value = a.ge(b)
			}
		}
	default:
		panic(internalError())
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
		a := n.ShiftExpression.check(ctx)
		b := n.AdditiveExpression.check(ctx)
		if !a.Type().IsIntegerType() || !b.Type().IsIntegerType() {
			//TODO report err
			break
		}

		n.promote = b.integerPromotion(ctx, n).Type()
		if a.Value() == nil || b.Value() == nil {
			n.Operand = &operand{typ: a.integerPromotion(ctx, n).Type()}
			break
		}

		n.Operand = &operand{typ: a.Type(), value: a.Value().lsh(b.Value())}
	case ShiftExpressionRsh: // ShiftExpression ">>" AdditiveExpression
		a := n.ShiftExpression.check(ctx)
		b := n.AdditiveExpression.check(ctx)
		if !a.Type().IsIntegerType() || !b.Type().IsIntegerType() {
			//TODO report err
			break
		}

		n.promote = b.integerPromotion(ctx, n).Type()
		if a.Value() == nil || b.Value() == nil {
			n.Operand = &operand{typ: a.integerPromotion(ctx, n).Type()}
			break
		}

		n.Operand = &operand{typ: a.Type(), value: a.Value().rsh(b.Value())}
	default:
		panic(internalError())
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
		if t := a.Type().Decay(); t.Kind() == Ptr && b.Type().IsScalarType() {
			n.Operand = &operand{typ: t}
			break
		}

		if t := b.Type().Decay(); t.Kind() == Ptr && a.Type().IsScalarType() {
			n.Operand = &operand{typ: t}
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

		n.Operand = &operand{typ: a.Type(), value: a.Value().add(b.Value())}
	case AdditiveExpressionSub: // AdditiveExpression '-' MultiplicativeExpression
		a := n.AdditiveExpression.check(ctx)
		b := n.MultiplicativeExpression.check(ctx)
		if a.Type().Decay().Kind() == Ptr && b.Type().Decay().Kind() == Ptr {
			n.Operand = ssizeT(ctx, n.lexicalScope, n.Token, nil)
			break
		}

		if t := a.Type().Decay(); t.Kind() == Ptr && b.Type().IsScalarType() {
			n.Operand = &operand{typ: t}
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
		panic(internalError())
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
		panic(internalError())
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

	typ = n.typ
	if typ.Kind() == Array && typ.IsVLA() {
		if f := ctx.checkFn; f != nil {
			f.VLAs = append(f.VLAs, n)
		}
	}
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
		return n.DirectDeclarator.check(ctx, checkArray(ctx, &n.Token, typ, n.AssignmentExpression, true, false))
	case DirectDeclaratorStaticArr: // DirectDeclarator '[' "static" TypeQualifiers AssignmentExpression ']'
		return n.DirectDeclarator.check(ctx, checkArray(ctx, &n.Token, typ, n.AssignmentExpression, false, false))
	case DirectDeclaratorArrStatic: // DirectDeclarator '[' TypeQualifiers "static" AssignmentExpression ']'
		return n.DirectDeclarator.check(ctx, checkArray(ctx, &n.Token, typ, n.AssignmentExpression, false, false))
	case DirectDeclaratorStar: // DirectDeclarator '[' TypeQualifiers '*' ']'
		return n.DirectDeclarator.check(ctx, checkArray(ctx, &n.Token, typ, nil, true, true))
	case DirectDeclaratorFuncParam: // DirectDeclarator '(' ParameterTypeList ')'
		ft := &functionType{typeBase: typeBase{kind: byte(Function)}, result: typ}
		n.ParameterTypeList.check(ctx, ft)
		return n.DirectDeclarator.check(ctx, ft)
	case DirectDeclaratorFuncIdent: // DirectDeclarator '(' IdentifierList ')'
		ft := &functionType{typeBase: typeBase{kind: byte(Function)}, result: typ, paramList: n.IdentifierList.check(ctx)}
		return n.DirectDeclarator.check(ctx, ft)
	default:
		panic(internalError())
	}
	return noType //TODO-
}

func checkArray(ctx *context, n Node, typ Type, expr *AssignmentExpression, exprIsOptional, noExpr bool) Type { //TODO pass and use typeQualifiers
	b := typ.base()
	b.align = byte(typ.Align())
	b.fieldAlign = byte(typ.FieldAlign())
	b.kind = byte(Array)
	switch {
	case expr != nil && noExpr:
		panic(internalError())
	case expr != nil:
		op := expr.check(ctx)
		if op.Type().Kind() == Invalid {
			return noType
		}

		if !op.Type().IsIntegerType() {
			//TODO report err
			return noType
		}

		var length uintptr
		var vla bool
		var vlaExpr *AssignmentExpression
		switch x := op.Value().(type) {
		case nil:
			vla = true
			vlaExpr = expr
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
		return &arrayType{typeBase: b, decay: ctx.cfg.ABI.Ptr(n, typ), elem: typ, length: length, vla: vla, expr: vlaExpr}
	case noExpr:
		// nop
	case !exprIsOptional:
		panic(internalError())
	}
	b.flags |= fIncomplete
	return &arrayType{typeBase: b, decay: ctx.cfg.ABI.Ptr(n, typ), elem: typ}
}

func (n *IdentifierList) check(ctx *context) (r []StringID) {
	for ; n != nil; n = n.IdentifierList {
		tok := n.Token2.Value
		if tok == 0 {
			tok = n.Token.Value
		}
		r = append(r, tok)
	}
	return r
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
		panic(internalError())
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
		panic(internalError())
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
			panic(internalError())
		}
	}
	return typ.check(ctx, n0, true)
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
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

// DeclarationSpecifiers Declarator DeclarationList CompoundStatement
func (n *FunctionDefinition) checkDeclarator(ctx *context) {
	if n == nil {
		return
	}

	ctx.checkFn = n
	typ := n.DeclarationSpecifiers.check(ctx)
	typ = n.Declarator.check(ctx, n.DeclarationSpecifiers, typ, true)
	ctx.checkFn = nil
	n.DeclarationList.checkFn(ctx, typ)
}

func (n *DeclarationList) checkFn(ctx *context, typ Type) {
	if n == nil {
		return
	}

	n.check(ctx)
	ft, ok := typ.(*functionType)
	if !ok {
		return
	}

	if ft.params != nil {
		//TODO report error
		return
	}
	if len(ft.paramList) == 0 {
		//TODO report error
		return
	}

	m := make(map[StringID]int, len(ft.paramList))
	for i, v := range ft.paramList {
		if _, ok := m[v]; ok {
			//TODO report error
			return
		}

		m[v] = i
	}
	params := make([]*Parameter, len(m))
	i := 0
	for ; n != nil; n = n.DeclarationList {
		for n := n.Declaration.InitDeclaratorList; n != nil; n = n.InitDeclaratorList {
			n := n.InitDeclarator
			switch n.Case {
			case InitDeclaratorDecl: // Declarator AttributeSpecifierList
				nm := n.Declarator.Name()
				switch x, ok := m[nm]; {
				case ok:
					params[x] = &Parameter{d: n.Declarator, typ: n.Declarator.Type()}
					i++
				default:
					//TODO report error
				}
			case InitDeclaratorInit: // Declarator AttributeSpecifierList '=' Initializer
				//TODO report error
				return
			default:
				panic(internalError())
			}
		}
	}
	switch {
	case i > len(m):
		ctx.errNode(n, "expected %d declarations, got %d", len(m), i)
		return
	case i < len(m):
		//TODO synthetize missing declarator
		return //TODO-
	}

	ft.params = params
}

func (n *CompoundStatement) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	return n.BlockItemList.check(ctx)
}

func (n *BlockItemList) check(ctx *context) (r Operand) {
	r = noOperand
	for ; n != nil; n = n.BlockItemList {
		r = n.BlockItem.check(ctx)
	}
	return r
}

func (n *BlockItem) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	switch n.Case {
	case BlockItemDecl: // Declaration
		n.Declaration.check(ctx, false)
	case BlockItemStmt: // Statement
		return n.Statement.check(ctx)
	case BlockItemLabel: // LabelDeclaration
		n.LabelDeclaration.check(ctx)
	case BlockItemFuncDef: // DeclarationSpecifiers Declarator CompoundStatement
		ctxClosure := ctx.closure
		ctx.closure = nil
		ctxCheckFn := ctx.checkFn
		n.fn = &FunctionDefinition{
			DeclarationSpecifiers: n.DeclarationSpecifiers,
			Declarator:            n.Declarator,
			CompoundStatement:     n.CompoundStatement,
		}
		ctx.checkFn = n.fn
		n.CompoundStatement.scope.declare(idClosure, n)
		ctx.checkFn.checkDeclarator(ctx)
		ctx.checkFn.checkBody(ctx)
		delete(n.CompoundStatement.scope, idClosure)
		ctx.checkFn = ctxCheckFn
		ctx.closure = ctxClosure
	case BlockItemPragma: // PragmaSTDC
		n.PragmaSTDC.check(ctx)
	default:
		panic(internalError())
	}
	return noOperand
}

func (n *LabelDeclaration) check(ctx *context) {
	if n == nil {
		return
	}

	n.IdentifierList.check(ctx)
}

func (n *Statement) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	switch n.Case {
	case StatementLabeled: // LabeledStatement
		n.LabeledStatement.check(ctx)
	case StatementCompound: // CompoundStatement
		n.CompoundStatement.check(ctx)
	case StatementExpr: // ExpressionStatement
		return n.ExpressionStatement.check(ctx)
	case StatementSelection: // SelectionStatement
		n.SelectionStatement.check(ctx)
	case StatementIteration: // IterationStatement
		n.IterationStatement.check(ctx)
	case StatementJump: // JumpStatement
		n.JumpStatement.check(ctx)
	case StatementAsm: // AsmStatement
		n.AsmStatement.check(ctx)
	default:
		panic(internalError())
	}
	return noOperand
}

func (n *JumpStatement) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case JumpStatementGoto: // "goto" IDENTIFIER ';'
		if ctx.checkFn.Gotos == nil {
			ctx.checkFn.Gotos = map[StringID]*JumpStatement{}
		}
		ctx.checkFn.Gotos[n.Token2.Value] = n
	case JumpStatementGotoExpr: // "goto" '*' Expression ';'
		n.Expression.check(ctx)
		//TODO
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
		panic(internalError())
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
		panic(internalError())
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
		panic(internalError())
	}
}

func (n *ExpressionStatement) check(ctx *context) Operand {
	if n == nil {
		return noOperand
	}

	n.AttributeSpecifierList.check(ctx)
	return n.Expression.check(ctx)
}

func (n *LabeledStatement) check(ctx *context) {
	if n == nil {
		return
	}

	switch n.Case {
	case LabeledStatementLabel: // IDENTIFIER ':' AttributeSpecifierList Statement
		if ctx.checkFn.Labels == nil {
			ctx.checkFn.Labels = map[StringID]*LabeledStatement{}
		}
		if _, ok := ctx.checkFn.Labels[n.Token.Value]; ok {
			//TODO report redeclared
		}
		ctx.checkFn.Labels[n.Token.Value] = n
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
		panic(internalError())
	}
}

func (n *DeclarationList) check(ctx *context) {
	for ; n != nil; n = n.DeclarationList {
		n.Declaration.check(ctx, false)
	}
}
