// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import "fmt"

// Expr is an interface type for C expressions.
type Expr interface {
	// isExpr disallows custom implementation of this interface.
	isExpr()
	// Type returns a type of an expression.
	Type() Type
	// TODO(dennwc): operand, isConst, ...
}

type baseExpr struct{}

func (*baseExpr) isExpr() {}

// NewIdent creates a new identifier with a given type.
func NewIdent(name string, typ Type) *Ident {
	if typ == nil {
		typ = InvalidType()
	}
	return &Ident{Name: name, typ: typ}
}

// Ident is an identifier in C.
type Ident struct {
	Name string
	typ  Type
	// TODO(dennwc): defined Scope, usages?
}

// Type implements Expr.
func (e *Ident) Type() Type {
	return e.typ
}

// IdentExpr is an identifier expression in C.
type IdentExpr struct {
	*Ident
}

func (IdentExpr) isExpr() {}

// LiteralKind is an enum for literal kinds in C.
type LiteralKind int

const (
	// LiteralInt is a kind for integer literals in C: 42, 0x42, etc.
	LiteralInt = LiteralKind(iota)
	// LiteralFloat is a kind for float/double literals: 1.0, 1e5, etc.
	LiteralFloat
	// LiteralChar is a kind for char literals: 'a', '\n', etc.
	LiteralChar
	// LiteralWChar is a kind for long/wide char literals: L'a'.
	LiteralWChar
	// LiteralString is a kind for string literals: "abc".
	LiteralString
	// LiteralString is a kind for long/wide string literals: L"abc".
	LiteralWString
)

// Literal is a constant literal expression in C.
// See LiteralKind for details on specific kinds.
type Literal struct {
	baseExpr
	typ   Type
	kind  LiteralKind
	value string
}

// Kind returns a kind of a literal.
func (e *Literal) Kind() LiteralKind {
	return e.kind
}

// Kind returns a literal value.
func (e *Literal) Value() string {
	return e.value
}

// Type implements Expr.
func (e *Literal) Type() Type {
	// TODO: type-check if no type is set
	return e.typ
}

// ParenExpr is a parentheses expression in C: (x).
// It returns the expression value without changes. Used primarily to control evaluation order in binary expressions.
type ParenExpr struct {
	baseExpr
	X Expr
}

// Type implements Expr.
func (e *ParenExpr) Type() Type {
	return e.X.Type()
}

// CommaExpr is a comma expression in C: x1, x2, ..., xN.
// It evaluates all expressions in order and returns the result of the last one only. Other results are discarded.
type CommaExpr []Expr

func (e CommaExpr) isExpr() {}

// Type implements Expr.
func (e CommaExpr) Type() Type {
	if len(e) == 0 {
		return InvalidType()
	}
	return e[len(e)-1].Type()
}

// UnaryOp is an enum for unary operators in C.
type UnaryOp int

const (
	// UnaryPlus is a plus operator in C: +x.
	UnaryPlus = UnaryOp(iota)
	// UnaryMinus is a minus operator in C: -x.
	UnaryMinus
	// UnaryInvert is a bit inversion operator in C: ~x.
	UnaryInvert
	// UnaryNot is a not operator in C: !x.
	UnaryNot
	// UnaryAddr is a take address operator in C: &x.
	UnaryAddr
	// UnaryDeref is a pointer dereference operator in C: *x.
	UnaryDeref
)

// UnaryExpr is an unary expression in C: !x, *x, etc.
type UnaryExpr struct {
	baseExpr
	typ Type
	Op  UnaryOp
	X   Expr
}

// Type implements Expr.
func (e *UnaryExpr) Type() Type {
	// TODO: type-check if no type is set
	return e.typ
}

// BinaryOp is an enum for binary operators in C.
type BinaryOp int

const (
	// BinaryNone is an fake C binary operator used in assign expressions: x = y.
	BinaryNone = BinaryOp(iota)
	// BinaryAdd is an addition operator in C: x + y.
	BinaryAdd
	// BinarySub is a subtraction operator in C: x - y.
	BinarySub
	// BinaryMul is a multiplication operator in C: x * y.
	BinaryMul
	// BinaryDiv is a division operator in C: x / y.
	BinaryDiv
	// BinaryMod is a modulo operator in C: x % y.
	BinaryMod
	// BinaryLsh is a binary left shift operator in C: x << y.
	BinaryLsh
	// BinaryRsh is a binary right shift operator in C: x >> y.
	BinaryRsh
	// BinaryEqual is an equality operator in C: x == y.
	BinaryEqual
	// BinaryNotEqual is an inequality operator in C: x != y.
	BinaryNotEqual
	// BinaryLess is a less than operator in C: x < y.
	BinaryLess
	// BinaryGreater is a greater than operator in C: x > y.
	BinaryGreater
	// BinaryLessEqual is a less than or equal operator in C: x <= y.
	BinaryLessEqual
	// BinaryGreaterEqual is a greater than or equal operator in C: x >= y.
	BinaryGreaterEqual
	// BinaryAnd is a logical and operator in C: x && y.
	BinaryAnd
	// BinaryOr is a logical or operator in C: x || y.
	BinaryOr
	// BinaryBitAnd is a bit and operator in C: x & y.
	BinaryBitAnd
	// BinaryBitOr is a bit or operator in C: x | y.
	BinaryBitOr
	// BinaryBitXOr is a exclusive bit or operator in C: x ^ y.
	BinaryBitXOr
)

// BinaryExpr is a binary expression in C: x + y, x == y, etc.
type BinaryExpr struct {
	baseExpr
	typ Type
	X   Expr
	Op  BinaryOp
	Y   Expr
}

// Type implements Expr.
func (e *BinaryExpr) Type() Type {
	// TODO: type-check if no type is set
	return e.typ
}

// AssignExpr is an assignment expression in C: x = y, x += y, etc.
// It returns a value that was assigned, allowing chaining assignments: x = y = z.
type AssignExpr struct {
	baseExpr
	Left  Expr
	Op    BinaryOp
	Right Expr
}

// Type implements Expr.
func (e *AssignExpr) Type() Type {
	return e.Left.Type()
}

// IncDecOp is an enum for increment/decrement operators in C.
type IncDecOp int

const (
	// IncPost is a postfix increment operator in C: x++.
	IncPost = IncDecOp(iota)
	// DecPost is a postfix decrement operator in C: x--.
	DecPost
	// IncPre is a prefix increment operator in C: ++x.
	IncPre
	// DecPre is a prefix decrement operator in C: --x.
	DecPre
)

// IncDecExpr is an increment/decrement expression in C: x++, ++x, x--, etc.
// Prefix operators first increment the value and return a modified one,
// while postfix variant return an old value and then increment the value.
type IncDecExpr struct {
	baseExpr
	X  Expr
	Op IncDecOp
}

// Type implements Expr.
func (e *IncDecExpr) Type() Type {
	return e.X.Type()
}

// IndexExpr is an index expression in C: x[y].
// The left operand should be either an array or a pointer.
type IndexExpr struct {
	baseExpr
	X   Expr
	Ind Expr
}

// Type implements Expr.
func (e *IndexExpr) Type() Type {
	return e.X.Type().Elem()
}

// SelectExpr is field select expression in C: x.y, x->y.
type SelectExpr struct {
	baseExpr
	X   Expr
	Sel *Ident
	Ptr bool
}

// Type implements Expr.
func (e *SelectExpr) Type() Type {
	return e.Sel.Type()
}

// CallExpr is a function call expression in C: x(a1, a2, a3).
type CallExpr struct {
	baseExpr
	Func Expr
	Args []Expr
}

// Type implements Expr.
func (e *CallExpr) Type() Type {
	return e.Func.Type().Result()
}

// CondExpr is a conditional expression in C: x ? y : z.
// If condition evaluates to true, "then" expression is returned. Otherwise, the "else" expression is returned.
type CondExpr struct {
	baseExpr
	typ  Type
	Cond Expr
	Then Expr
	Else Expr
}

// Type implements Expr.
func (e *CondExpr) Type() Type {
	// TODO: type-check if no type is set
	return e.typ
}

func operandType(o Operand) Type {
	if o == nil {
		return InvalidType()
	}
	return o.Type()
}

func (n *PrimaryExpression) Expr() Expr {
	switch n.Case {
	case PrimaryExpressionIdent:
		// TODO(dennwc): "asm" expression
		id := NewIdent(n.Token.Value.String(), operandType(n.Operand))
		return IdentExpr{id}
	case PrimaryExpressionEnum:
		id := NewIdent(n.Token.Value.String(), operandType(n.Operand))
		return IdentExpr{id}
	case PrimaryExpressionInt:
		return &Literal{typ: operandType(n.Operand), kind: LiteralInt, value: n.Token.Value.String()}
	case PrimaryExpressionFloat:
		return &Literal{typ: operandType(n.Operand), kind: LiteralFloat, value: n.Token.Value.String()}
	case PrimaryExpressionChar:
		return &Literal{typ: operandType(n.Operand), kind: LiteralChar, value: n.Token.Value.String()}
	case PrimaryExpressionLChar:
		return &Literal{typ: operandType(n.Operand), kind: LiteralWChar, value: n.Token.Value.String()}
	case PrimaryExpressionString:
		return &Literal{typ: operandType(n.Operand), kind: LiteralString, value: n.Token.Value.String()}
	case PrimaryExpressionLString:
		return &Literal{typ: operandType(n.Operand), kind: LiteralWString, value: n.Token.Value.String()}
	case PrimaryExpressionExpr:
		return &ParenExpr{X: n.Expression.Expr()}
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
}

func (n *PostfixExpression) Expr() Expr {
	switch n.Case {
	case PostfixExpressionPrimary:
		return n.PrimaryExpression.Expr()
	case PostfixExpressionIndex:
		return &IndexExpr{
			X:   n.PostfixExpression.Expr(),
			Ind: n.Expression.Expr(),
		}
	case PostfixExpressionSelect, PostfixExpressionPSelect:
		return &SelectExpr{
			X:   n.PostfixExpression.Expr(),
			Sel: &Ident{Name: n.Token2.Value.String()},
			Ptr: n.Case == PostfixExpressionPSelect,
		}
	case PostfixExpressionCall:
		var args []Expr
		for it := n.ArgumentExpressionList; it != nil; it = it.ArgumentExpressionList {
			args = append(args, it.AssignmentExpression.Expr())
		}
		return &CallExpr{
			Func: n.PostfixExpression.Expr(),
			Args: args,
		}
	case PostfixExpressionInc:
		return &IncDecExpr{X: n.PostfixExpression.Expr(), Op: IncPost}
	case PostfixExpressionDec:
		return &IncDecExpr{X: n.PostfixExpression.Expr(), Op: DecPost}
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
}

func (n *UnaryExpression) Expr() Expr {
	switch n.Case {
	case UnaryExpressionPostfix:
		return n.PostfixExpression.Expr()
	case UnaryExpressionInc:
		return &IncDecExpr{Op: IncPre, X: n.UnaryExpression.Expr()}
	case UnaryExpressionDec:
		return &IncDecExpr{Op: DecPre, X: n.UnaryExpression.Expr()}
	}
	var op UnaryOp
	switch n.Case {
	case UnaryExpressionAddrof:
		op = UnaryAddr
	case UnaryExpressionDeref:
		op = UnaryDeref
	case UnaryExpressionPlus:
		op = UnaryPlus
	case UnaryExpressionMinus:
		op = UnaryMinus
	case UnaryExpressionCpl:
		op = UnaryInvert
	case UnaryExpressionNot:
		op = UnaryNot
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
	return &UnaryExpr{
		typ: operandType(n.Operand),
		Op:  op, X: n.CastExpression.Expr(),
	}
}

func (n *CastExpression) Expr() Expr {
	switch n.Case {
	case CastExpressionUnary:
		return n.UnaryExpression.Expr()
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
}

func (n *MultiplicativeExpression) Expr() Expr {
	switch n.Case {
	case MultiplicativeExpressionCast:
		return n.CastExpression.Expr()
	}
	x := n.MultiplicativeExpression.Expr()
	y := n.CastExpression.Expr()
	var op BinaryOp
	switch n.Case {
	case MultiplicativeExpressionMul:
		op = BinaryMul
	case MultiplicativeExpressionDiv:
		op = BinaryDiv
	case MultiplicativeExpressionMod:
		op = BinaryMod
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
	return &BinaryExpr{
		typ: operandType(n.Operand),
		X:   x, Op: op, Y: y,
	}
}

func (n *AdditiveExpression) Expr() Expr {
	switch n.Case {
	case AdditiveExpressionMul:
		return n.MultiplicativeExpression.Expr()
	}
	x := n.AdditiveExpression.Expr()
	y := n.MultiplicativeExpression.Expr()
	var op BinaryOp
	switch n.Case {
	case AdditiveExpressionAdd:
		op = BinaryAdd
	case AdditiveExpressionSub:
		op = BinarySub
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
	return &BinaryExpr{
		typ: operandType(n.Operand),
		X:   x, Op: op, Y: y,
	}
}

func (n *ShiftExpression) Expr() Expr {
	switch n.Case {
	case ShiftExpressionAdd:
		return n.AdditiveExpression.Expr()
	}
	x := n.ShiftExpression.Expr()
	y := n.AdditiveExpression.Expr()
	var op BinaryOp
	switch n.Case {
	case ShiftExpressionLsh:
		op = BinaryLsh
	case ShiftExpressionRsh:
		op = BinaryRsh
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
	return &BinaryExpr{
		typ: operandType(n.Operand),
		X:   x, Op: op, Y: y,
	}
}

func (n *RelationalExpression) Expr() Expr {
	switch n.Case {
	case RelationalExpressionShift:
		return n.ShiftExpression.Expr()
	}
	x := n.RelationalExpression.Expr()
	y := n.ShiftExpression.Expr()
	var op BinaryOp
	switch n.Case {
	case RelationalExpressionLt:
		op = BinaryLess
	case RelationalExpressionGt:
		op = BinaryGreater
	case RelationalExpressionLeq:
		op = BinaryLessEqual
	case RelationalExpressionGeq:
		op = BinaryGreaterEqual
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
	return &BinaryExpr{
		typ: operandType(n.Operand),
		X:   x, Op: op, Y: y,
	}
}

func (n *EqualityExpression) Expr() Expr {
	switch n.Case {
	case EqualityExpressionRel:
		return n.RelationalExpression.Expr()
	}
	x := n.EqualityExpression.Expr()
	y := n.RelationalExpression.Expr()
	var op BinaryOp
	switch n.Case {
	case EqualityExpressionEq:
		op = BinaryEqual
	case EqualityExpressionNeq:
		op = BinaryNotEqual
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
	return &BinaryExpr{
		typ: operandType(n.Operand),
		X:   x, Op: op, Y: y,
	}
}

func (n *AndExpression) Expr() Expr {
	switch n.Case {
	case AndExpressionEq:
		return n.EqualityExpression.Expr()
	}
	x := n.AndExpression.Expr()
	y := n.EqualityExpression.Expr()
	var op BinaryOp
	switch n.Case {
	case AndExpressionAnd:
		op = BinaryBitAnd
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
	return &BinaryExpr{
		typ: operandType(n.Operand),
		X:   x, Op: op, Y: y,
	}
}

func (n *ExclusiveOrExpression) Expr() Expr {
	switch n.Case {
	case ExclusiveOrExpressionAnd:
		return n.AndExpression.Expr()
	}
	x := n.ExclusiveOrExpression.Expr()
	y := n.AndExpression.Expr()
	var op BinaryOp
	switch n.Case {
	case ExclusiveOrExpressionXor:
		op = BinaryBitXOr
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
	return &BinaryExpr{
		typ: operandType(n.Operand),
		X:   x, Op: op, Y: y,
	}
}

func (n *InclusiveOrExpression) Expr() Expr {
	switch n.Case {
	case InclusiveOrExpressionXor:
		return n.ExclusiveOrExpression.Expr()
	}
	x := n.InclusiveOrExpression.Expr()
	y := n.ExclusiveOrExpression.Expr()
	var op BinaryOp
	switch n.Case {
	case InclusiveOrExpressionOr:
		op = BinaryBitOr
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
	return &BinaryExpr{
		typ: operandType(n.Operand),
		X:   x, Op: op, Y: y,
	}
}

func (n *LogicalAndExpression) Expr() Expr {
	switch n.Case {
	case LogicalAndExpressionOr:
		return n.InclusiveOrExpression.Expr()
	}
	x := n.LogicalAndExpression.Expr()
	y := n.InclusiveOrExpression.Expr()
	var op BinaryOp
	switch n.Case {
	case LogicalAndExpressionLAnd:
		op = BinaryAnd
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
	return &BinaryExpr{
		typ: operandType(n.Operand),
		X:   x, Op: op, Y: y,
	}
}

func (n *LogicalOrExpression) Expr() Expr {
	switch n.Case {
	case LogicalOrExpressionLAnd:
		return n.LogicalAndExpression.Expr()
	}
	x := n.LogicalOrExpression.Expr()
	y := n.LogicalAndExpression.Expr()
	var op BinaryOp
	switch n.Case {
	case LogicalOrExpressionLOr:
		op = BinaryOr
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
	return &BinaryExpr{
		typ: operandType(n.Operand),
		X:   x, Op: op, Y: y,
	}
}

func (n *ConditionalExpression) Expr() Expr {
	switch n.Case {
	case ConditionalExpressionLOr:
		return n.LogicalOrExpression.Expr()
	case ConditionalExpressionCond:
		return &CondExpr{
			typ:  operandType(n.Operand),
			Cond: n.LogicalOrExpression.Expr(),
			Then: n.Expression.Expr(),
			Else: n.ConditionalExpression.Expr(),
		}
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
}

func (n *AssignmentExpression) Expr() Expr {
	switch n.Case {
	case AssignmentExpressionCond:
		return n.ConditionalExpression.Expr()
	}
	left := n.UnaryExpression.Expr()
	right := n.AssignmentExpression.Expr()
	var op BinaryOp
	switch n.Case {
	case AssignmentExpressionAssign:
		op = BinaryNone
	case AssignmentExpressionMul:
		op = BinaryMul
	case AssignmentExpressionDiv:
		op = BinaryDiv
	case AssignmentExpressionMod:
		op = BinaryMod
	case AssignmentExpressionAdd:
		op = BinaryAdd
	case AssignmentExpressionSub:
		op = BinarySub
	case AssignmentExpressionLsh:
		op = BinaryLsh
	case AssignmentExpressionRsh:
		op = BinaryRsh
	case AssignmentExpressionAnd:
		op = BinaryBitAnd
	case AssignmentExpressionXor:
		op = BinaryBitXOr
	case AssignmentExpressionOr:
		op = BinaryBitOr
	default:
		panic(fmt.Errorf("TODO: case %v (%v)", n.Case, n.Position()))
	}
	return &AssignExpr{
		Left: left, Op: op, Right: right,
	}
}

func (n *Expression) Expr() Expr {
	if n.Expression == nil {
		return n.AssignmentExpression.Expr()
	}
	var arr []*AssignmentExpression
	for it := n; it != nil; it = it.Expression {
		arr = append(arr, it.AssignmentExpression)
	}
	expr := make(CommaExpr, 0, len(arr))
	// order is reversed: returned expression is the first in linked list
	for i := len(arr) - 1; i >= 0; i-- {
		expr = append(expr, arr[i].Expr())
	}
	return expr
}

func (n *ConstantExpression) Expr() Expr {
	return n.ConditionalExpression.Expr()
}
