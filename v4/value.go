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
	oneValue     = Int64Value(1)
	zeroValue    = Int64Value(0)
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
				n.val = c.convert(n.Expression.eval(c), n.Type())
			case isNonzero(val):
				n.val = c.convert(n.ConditionalExpression.eval(c), n.Type())
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
			switch v := c.convert(n.LogicalAndExpression.eval(c), n.Type()); {
			case isZero(v):
				c.errors.add(errorf("TODO %v", n.Case))
			case isNonzero(v):
				n.val = oneValue
			default:
				c.errors.add(errorf("TODO %v", n.Case))
			}
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
			switch v := n.LogicalAndExpression.eval(c); {
			case isZero(v):
				c.errors.add(errorf("TODO %v", n.Case))
			case isNonzero(v):
				switch w := n.InclusiveOrExpression.eval(c); {
				case isZero(w):
					c.errors.add(errorf("TODO %v", n.Case))
				case isNonzero(w):
					n.val = oneValue
				default:
					c.errors.add(errorf("TODO %v", n.Case))
				}
			default:
				c.errors.add(errorf("TODO %v", n.Case))
			}
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
			switch x := c.convert(n.InclusiveOrExpression.eval(c), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				switch y := c.convert(n.InclusiveOrExpression.eval(c), n.Type()).(type) {
				case unknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x|y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
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
			if !isArithmeticType(n.EqualityExpression.Type()) || !isArithmeticType(n.RelationalExpression.Type()) {
				break
			}

			switch x := c.convert(n.EqualityExpression.eval(c), n.Type()).(type) {
			case unknownValue:
				// ok
			case Int64Value:
				switch y := c.convert(n.RelationalExpression.eval(c), n.Type()).(type) {
				case unknownValue:
					// ok
				case Int64Value:
					n.val = bool2int(x == y)
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
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
			if !isArithmeticType(n.RelationalExpression.Type()) || !isArithmeticType(n.ShiftExpression.Type()) {
				break
			}

			switch x := c.convert(n.RelationalExpression.eval(c), n.Type()).(type) {
			case unknownValue:
				// ok
			case Int64Value:
				switch y := c.convert(n.ShiftExpression.eval(c), n.Type()).(type) {
				case unknownValue:
					// ok
				case Int64Value:
					if x < y {
						n.val = oneValue
						break
					}

					n.val = zeroValue
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
			switch x := c.convert(n.ShiftExpression.eval(c), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				switch y := n.AdditiveExpression.eval(c).(type) {
				case unknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x<<y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case ShiftExpressionRsh: // ShiftExpression ">>" AdditiveExpression
			switch x := c.convert(n.ShiftExpression.eval(c), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				switch y := n.AdditiveExpression.eval(c).(type) {
				case unknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x>>y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
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
			if !isArithmeticType(n.AdditiveExpression.Type()) || !isArithmeticType(n.MultiplicativeExpression.Type()) {
				break
			}

			switch x := c.convert(n.AdditiveExpression.eval(c), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				switch y := c.convert(n.MultiplicativeExpression.eval(c), n.Type()).(type) {
				case unknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x+y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case UInt64Value:
				switch y := c.convert(n.MultiplicativeExpression.eval(c), n.Type()).(type) {
				case unknownValue:
					// nop
				case UInt64Value:
					n.val = c.convert(x+y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case AdditiveExpressionSub: // AdditiveExpression '-' MultiplicativeExpression
			if !isArithmeticType(n.AdditiveExpression.Type()) || !isArithmeticType(n.MultiplicativeExpression.Type()) {
				break
			}

			switch x := c.convert(n.AdditiveExpression.eval(c), n.Type()).(type) {
			case unknownValue:
				// nop
			case UInt64Value:
				switch y := c.convert(n.MultiplicativeExpression.eval(c), n.Type()).(type) {
				case unknownValue:
					// nop
				case UInt64Value:
					n.val = c.convert(x-y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case Int64Value:
				switch y := c.convert(n.MultiplicativeExpression.eval(c), n.Type()).(type) {
				case unknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x-y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
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
			switch x := c.convert(n.MultiplicativeExpression.eval(c), n.Type()).(type) {
			case unknownValue:
				// nop
			case UInt64Value:
				switch y := c.convert(n.CastExpression.eval(c), n.Type()).(type) {
				case unknownValue:
					// nop
				case UInt64Value:
					n.val = c.convert(x*y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case Int64Value:
				switch y := c.convert(n.CastExpression.eval(c), n.Type()).(type) {
				case unknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x*y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case MultiplicativeExpressionDiv: // MultiplicativeExpression '/' CastExpression
			switch x := c.convert(n.MultiplicativeExpression.eval(c), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				switch y := c.convert(n.CastExpression.eval(c), n.Type()).(type) {
				case unknownValue:
					// nop
				case Int64Value:
					if y != 0 {
						n.val = c.convert(x/y, n.Type())
						break
					}

					c.errors.add(errorf("%v: integer division by zero", n.Token.Position()))
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case UInt64Value:
				switch y := c.convert(n.CastExpression.eval(c), n.Type()).(type) {
				case unknownValue:
					// nop
				case UInt64Value:
					if y != 0 {
						n.val = c.convert(x/y, n.Type())
						break
					}

					c.errors.add(errorf("%v: integer division by zero", n.Token.Position()))
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case MultiplicativeExpressionMod: // MultiplicativeExpression '%' CastExpression
			switch x := c.convert(n.MultiplicativeExpression.eval(c), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				switch y := c.convert(n.CastExpression.eval(c), n.Type()).(type) {
				case unknownValue:
					// nop
				case Int64Value:
					if y != 0 {
						n.val = c.convert(x%y, n.Type())
						break
					}

					c.errors.add(errorf("%v: integer division by zero", n.Token.Position()))
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
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
			n.val = c.convert(n.CastExpression.eval(c), n.TypeName.Type())
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
			switch x := c.convert(n.CastExpression.eval(c), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				n.val = c.convert(x, n.Type())
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case UnaryExpressionMinus: // '-' CastExpression
			switch x := c.convert(n.CastExpression.eval(c), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				n.val = c.convert(-x, n.Type())
			case UInt64Value:
				n.val = c.convert(-x, n.Type())
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case UnaryExpressionCpl: // '~' CastExpression
			c.errors.add(errorf("TODO %v", n.Case))
		case UnaryExpressionNot: // '!' CastExpression
			switch x := n.CastExpression.eval(c).(type) {
			case unknownValue:
				// nop
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
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
			// nop
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
