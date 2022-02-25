// Copyright 2022 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"math/big"
)

var (
	_ Value = (*ComplexLongDoubleValue)(nil)
	_ Value = (*LongDoubleValue)(nil)
	_ Value = Complex128Value(0)
	_ Value = Complex64Value(0)
	_ Value = Float64Value(0)
	_ Value = Int64Value(0)
	_ Value = UInt64Value(0)
	_ Value = unknownValue(0)
)

var (
	UnknownValue unknownValue
)

type Value interface {
	isValue()
}

type unknownValue int

func (n unknownValue) isValue() {}

type ComplexLongDoubleValue struct {
	Re *big.Float
	Im *big.Float
}

func (n *ComplexLongDoubleValue) isValue() {}

type LongDoubleValue big.Float

func (n *LongDoubleValue) isValue() {}

type Complex128Value complex128

func (n Complex128Value) isValue() {}

type Complex64Value complex64

func (n Complex64Value) isValue() {}

type Float64Value float64

func (n Float64Value) isValue() {}

type Int64Value int64

func (n Int64Value) isValue() {}

type UInt64Value int64

func (n UInt64Value) isValue() {}

func convert(c *ctx, v Value, t Type) (r Value) {
	switch t.Kind() {
	case Int:
		m := uint64(1)<<t.Size() - 1
		switch x := v.(type) {
		case Int64Value:
			if x < 0 {
				return x | Int64Value(^m)
			}

			return x & Int64Value(m)
		default:
			c.errors.add(errorf("TODO TYPE %T", x))
		}
	default:
		c.errors.add(errorf("TODO %v", t.Kind()))
	}
	return UnknownValue
}

func (n *ConstantExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = n.ConditionalExpression.eval(c)
	}
	return n.Value()
}

func (n *ConditionalExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case ConditionalExpressionLOr: // LogicalOrExpression
			n.val = n.LogicalOrExpression.eval(c)
		case ConditionalExpressionCond: // LogicalOrExpression '?' Expression ':' ConditionalExpression
			switch val := n.LogicalOrExpression.eval(c); {
			case isZero(val):
				c.errors.add(errorf("TODO %v", n.Case))
			case isNonzero(val):
				c.errors.add(errorf("TODO %v", n.Case))
			}
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *LogicalOrExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case LogicalOrExpressionLAnd: // LogicalAndExpression
			n.val = n.LogicalAndExpression.eval(c)
		case LogicalOrExpressionLOr: // LogicalOrExpression "||" LogicalAndExpression
			c.errors.add(errorf("TODO %v", n.Case))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *LogicalAndExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case LogicalAndExpressionOr: // InclusiveOrExpression
			n.val = n.InclusiveOrExpression.eval(c)
		case LogicalAndExpressionLAnd: // LogicalAndExpression "&&" InclusiveOrExpression
			c.errors.add(errorf("TODO %v", n.Case))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *InclusiveOrExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case InclusiveOrExpressionXor: // ExclusiveOrExpression
			n.val = n.ExclusiveOrExpression.eval(c)
		case InclusiveOrExpressionOr: // InclusiveOrExpression '|' ExclusiveOrExpression
			switch x := n.InclusiveOrExpression.Value().(type) {
			case unknownValue:
				// nop
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *ExclusiveOrExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case ExclusiveOrExpressionAnd: // AndExpression
			n.val = n.AndExpression.eval(c)
		case ExclusiveOrExpressionXor: // ExclusiveOrExpression '^' AndExpression
			c.errors.add(errorf("TODO %v", n.Case))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *AndExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case AndExpressionEq: // EqualityExpression
			n.val = n.EqualityExpression.eval(c)
		case AndExpressionAnd: // AndExpression '&' EqualityExpression
			c.errors.add(errorf("TODO %v", n.Case))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *EqualityExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case EqualityExpressionRel: // RelationalExpression
			n.val = n.RelationalExpression.eval(c)
		case EqualityExpressionEq: // EqualityExpression "==" RelationalExpression
			c.errors.add(errorf("TODO %v", n.Case))
		case EqualityExpressionNeq: // EqualityExpression "!=" RelationalExpression
			c.errors.add(errorf("TODO %v", n.Case))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *RelationalExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case RelationalExpressionShift: // ShiftExpression
			n.val = n.ShiftExpression.eval(c)
		case RelationalExpressionLt: // RelationalExpression '<' ShiftExpression
			a, b := n.RelationalExpression.Type(), n.ShiftExpression.Type()
			if !isArithmeticType(a) || !isArithmeticType(b) {
				break
			}

			t := usualArithmeticConversions(a, b)
			switch x := convert(c, n.RelationalExpression.eval(c), t).(type) {
			case unknownValue:
				// ok
			case Int64Value:
				switch y := convert(c, n.ShiftExpression.eval(c), t).(type) {
				case unknownValue:
					// ok
				case Int64Value:
					if x < y {
						n.val = Int64Value(1)
						break
					}

					n.val = Int64Value(0)
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case RelationalExpressionGt: // RelationalExpression '>' ShiftExpression
			c.errors.add(errorf("TODO %v", n.Case))
		case RelationalExpressionLeq: // RelationalExpression "<=" ShiftExpression
			c.errors.add(errorf("TODO %v", n.Case))
		case RelationalExpressionGeq: // RelationalExpression ">=" ShiftExpression
			c.errors.add(errorf("TODO %v", n.Case))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *ShiftExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case ShiftExpressionAdd: // AdditiveExpression
			n.val = n.AdditiveExpression.eval(c)
		case ShiftExpressionLsh: // ShiftExpression "<<" AdditiveExpression
			switch x := n.ShiftExpression.Value().(type) {
			case unknownValue:
				// nop
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case ShiftExpressionRsh: // ShiftExpression ">>" AdditiveExpression
			c.errors.add(errorf("TODO %v", n.Case))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *AdditiveExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case AdditiveExpressionMul: // MultiplicativeExpression
			n.val = n.MultiplicativeExpression.eval(c)
		case AdditiveExpressionAdd: // AdditiveExpression '+' MultiplicativeExpression
			switch x := n.AdditiveExpression.Value().(type) {
			case unknownValue:
				// nop
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case AdditiveExpressionSub: // AdditiveExpression '-' MultiplicativeExpression
			switch x := n.AdditiveExpression.Value().(type) {
			case unknownValue:
				// nop
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *MultiplicativeExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case MultiplicativeExpressionCast: // CastExpression
			n.val = n.CastExpression.eval(c)
		case MultiplicativeExpressionMul: // MultiplicativeExpression '*' CastExpression
			c.errors.add(errorf("TODO %v", n.Case))
		case MultiplicativeExpressionDiv: // MultiplicativeExpression '/' CastExpression
			switch x := n.MultiplicativeExpression.Value().(type) {
			case unknownValue:
				// nop
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case MultiplicativeExpressionMod: // MultiplicativeExpression '%' CastExpression
			c.errors.add(errorf("TODO %v", n.Case))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *CastExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case CastExpressionUnary: // UnaryExpression
			n.val = n.UnaryExpression.eval(c)
		case CastExpressionCast: // '(' TypeName ')' CastExpression
			c.errors.add(errorf("TODO %v", n.Case))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *UnaryExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case UnaryExpressionPostfix: // PostfixExpression
			n.val = n.PostfixExpression.eval(c)
		case UnaryExpressionInc: // "++" UnaryExpression
			c.errors.add(errorf("TODO %v", n.Case))
		case UnaryExpressionDec: // "--" UnaryExpression
			c.errors.add(errorf("TODO %v", n.Case))
		case UnaryExpressionAddrof: // '&' CastExpression
			c.errors.add(errorf("TODO %v", n.Case))
		case UnaryExpressionDeref: // '*' CastExpression
			c.errors.add(errorf("TODO %v", n.Case))
		case UnaryExpressionPlus: // '+' CastExpression
			switch x := n.CastExpression.Value().(type) {
			case unknownValue:
				// nop
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case UnaryExpressionMinus: // '-' CastExpression
			switch x := n.CastExpression.Value().(type) {
			case unknownValue:
				// nop
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case UnaryExpressionCpl: // '~' CastExpression
			c.errors.add(errorf("TODO %v", n.Case))
		case UnaryExpressionNot: // '!' CastExpression
			c.errors.add(errorf("TODO %v", n.Case))
		case UnaryExpressionSizeofExpr: // "sizeof" UnaryExpression
			c.errors.add(errorf("TODO %v", n.Case))
		case UnaryExpressionSizeofType: // "sizeof" '(' TypeName ')'
			c.errors.add(errorf("TODO %v", n.Case))
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
	}
	return n.Value()
}

func (n *PostfixExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case PostfixExpressionPrimary: // PrimaryExpression
			n.val = n.PrimaryExpression.eval(c)
		case PostfixExpressionIndex: // PostfixExpression '[' Expression ']'
			c.errors.add(errorf("TODO %v", n.Case))
		case PostfixExpressionCall: // PostfixExpression '(' ArgumentExpressionList ')'
			c.errors.add(errorf("TODO %v", n.Case))
		case PostfixExpressionSelect: // PostfixExpression '.' IDENTIFIER
			c.errors.add(errorf("TODO %v", n.Case))
		case PostfixExpressionPSelect: // PostfixExpression "->" IDENTIFIER
			c.errors.add(errorf("TODO %v", n.Case))
		case PostfixExpressionInc: // PostfixExpression "++"
			c.errors.add(errorf("TODO %v", n.Case))
		case PostfixExpressionDec: // PostfixExpression "--"
			c.errors.add(errorf("TODO %v", n.Case))
		case PostfixExpressionComplit: // '(' TypeName ')' '{' InitializerList ',' '}'
			c.errors.add(errorf("TODO %v", n.Case))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *PrimaryExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case PrimaryExpressionIdent: // IDENTIFIER
			c.errors.add(errorf("TODO %v", n.Case))
		case PrimaryExpressionInt: // INTCONST
			c.errors.add(errorf("TODO %v", n.Case))
		case PrimaryExpressionFloat: // FLOATCONST
			c.errors.add(errorf("TODO %v", n.Case))
		case PrimaryExpressionChar: // CHARCONST
			c.errors.add(errorf("TODO %v", n.Case))
		case PrimaryExpressionLChar: // LONGCHARCONST
			c.errors.add(errorf("TODO %v", n.Case))
		case PrimaryExpressionString: // STRINGLITERAL
			c.errors.add(errorf("TODO %v", n.Case))
		case PrimaryExpressionLString: // LONGSTRINGLITERAL
			c.errors.add(errorf("TODO %v", n.Case))
		case PrimaryExpressionExpr: // '(' Expression ')'
			n.val = n.Expression.eval(c)
		case PrimaryExpressionStmt: // '(' CompoundStatement ')'
			c.errors.add(errorf("TODO %v", n.Case))
		case PrimaryExpressionGeneric: // GenericSelection
			c.errors.add(errorf("TODO %v", n.Case))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *Expression) eval(c *ctx) (r Value) {
	n0 := n
	if n.val == nil {
		n.val = UnknownValue
		for ; n != nil; n = n.Expression {
			n0.typ = n.AssignmentExpression.typ
			n0.val = n.AssignmentExpression.eval(c)
		}
	}
	return n0.Value()
}

func (n *AssignmentExpression) eval(c *ctx) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		switch n.Case {
		case AssignmentExpressionCond: // ConditionalExpression
			n.val = n.ConditionalExpression.eval(c)
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
	}
	return n.Value()
}

func isZero(v Value) bool {
	switch x := v.(type) {
	case unknownValue:
		return false
	case Int64Value:
		return x == 0
	case UInt64Value:
		return x == 0
	default:
		panic(todo("%T", x))
	}
}

func isNonzero(v Value) bool {
	switch x := v.(type) {
	case unknownValue:
		return false
	case Int64Value:
		return x != 0
	case UInt64Value:
		return x != 0
	default:
		panic(todo("%T", x))
	}
}
