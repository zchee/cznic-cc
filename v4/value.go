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

func (n *ConstantExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = n.ConditionalExpression.eval(c, mode)
	}
	return n.Value()
}

func (n *ConditionalExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case ConditionalExpressionLOr: // LogicalOrExpression
			n.val = n.LogicalOrExpression.eval(c, mode)
		case ConditionalExpressionCond: // LogicalOrExpression '?' ExpressionList ':' ConditionalExpression
			switch val := n.LogicalOrExpression.eval(c, mode); {
			case isZero(val):
				n.val = c.convert(n.ExpressionList.eval(c, mode), n.Type())
			case isNonzero(val):
				n.val = c.convert(n.ConditionalExpression.eval(c, mode), n.Type())
			}
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *LogicalOrExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case LogicalOrExpressionLAnd: // LogicalAndExpression
			n.val = n.LogicalAndExpression.eval(c, mode)
		case LogicalOrExpressionLOr: // LogicalOrExpression "||" LogicalAndExpression
			switch v := c.convert(n.LogicalAndExpression.eval(c, mode), n.Type()); {
			case isZero(v):
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case isNonzero(v):
				n.val = oneValue
			default:
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			}
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *LogicalAndExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case LogicalAndExpressionOr: // InclusiveOrExpression
			n.val = n.InclusiveOrExpression.eval(c, mode)
		case LogicalAndExpressionLAnd: // LogicalAndExpression "&&" InclusiveOrExpression
			switch v := n.LogicalAndExpression.eval(c, mode); {
			case isZero(v):
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case isNonzero(v):
				switch w := n.InclusiveOrExpression.eval(c, mode); {
				case isZero(w):
					c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
				case isNonzero(w):
					n.val = oneValue
				default:
					c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
				}
			}
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *InclusiveOrExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case InclusiveOrExpressionXor: // ExclusiveOrExpression
			n.val = n.ExclusiveOrExpression.eval(c, mode)
		case InclusiveOrExpressionOr: // InclusiveOrExpression '|' ExclusiveOrExpression
			switch x := c.convert(n.InclusiveOrExpression.eval(c, mode), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				switch y := c.convert(n.InclusiveOrExpression.eval(c, mode), n.Type()).(type) {
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

func (n *ExclusiveOrExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case ExclusiveOrExpressionAnd: // AndExpression
			n.val = n.AndExpression.eval(c, mode)
		case ExclusiveOrExpressionXor: // ExclusiveOrExpression '^' AndExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *AndExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case AndExpressionEq: // EqualityExpression
			n.val = n.EqualityExpression.eval(c, mode)
		case AndExpressionAnd: // AndExpression '&' EqualityExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *EqualityExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case EqualityExpressionRel: // RelationalExpression
			n.val = n.RelationalExpression.eval(c, mode)
		case EqualityExpressionEq: // EqualityExpression "==" RelationalExpression
			if !isArithmeticType(n.EqualityExpression.Type()) || !isArithmeticType(n.RelationalExpression.Type()) {
				break
			}

			switch x := c.convert(n.EqualityExpression.eval(c, mode), n.Type()).(type) {
			case unknownValue:
				// ok
			case Int64Value:
				switch y := c.convert(n.RelationalExpression.eval(c, mode), n.Type()).(type) {
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
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *RelationalExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case RelationalExpressionShift: // ShiftExpression
			n.val = n.ShiftExpression.eval(c, mode)
		case RelationalExpressionLt: // RelationalExpression '<' ShiftExpression
			if !isArithmeticType(n.RelationalExpression.Type()) || !isArithmeticType(n.ShiftExpression.Type()) {
				break
			}

			switch x := c.convert(n.RelationalExpression.eval(c, mode), n.Type()).(type) {
			case unknownValue:
				// ok
			case Int64Value:
				switch y := c.convert(n.ShiftExpression.eval(c, mode), n.Type()).(type) {
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
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case RelationalExpressionLeq: // RelationalExpression "<=" ShiftExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case RelationalExpressionGeq: // RelationalExpression ">=" ShiftExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *ShiftExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case ShiftExpressionAdd: // AdditiveExpression
			n.val = n.AdditiveExpression.eval(c, mode)
		case ShiftExpressionLsh: // ShiftExpression "<<" AdditiveExpression
			switch x := c.convert(n.ShiftExpression.eval(c, mode), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				switch y := n.AdditiveExpression.eval(c, mode).(type) {
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
			switch x := c.convert(n.ShiftExpression.eval(c, mode), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				switch y := n.AdditiveExpression.eval(c, mode).(type) {
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

func (n *AdditiveExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case AdditiveExpressionMul: // MultiplicativeExpression
			n.val = n.MultiplicativeExpression.eval(c, mode)
		case AdditiveExpressionAdd: // AdditiveExpression '+' MultiplicativeExpression
			if !isArithmeticType(n.AdditiveExpression.Type()) || !isArithmeticType(n.MultiplicativeExpression.Type()) {
				break
			}

			switch x := c.convert(n.AdditiveExpression.eval(c, mode), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				switch y := c.convert(n.MultiplicativeExpression.eval(c, mode), n.Type()).(type) {
				case unknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x+y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case UInt64Value:
				switch y := c.convert(n.MultiplicativeExpression.eval(c, mode), n.Type()).(type) {
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

			switch x := c.convert(n.AdditiveExpression.eval(c, mode), n.Type()).(type) {
			case unknownValue:
				// nop
			case UInt64Value:
				switch y := c.convert(n.MultiplicativeExpression.eval(c, mode), n.Type()).(type) {
				case unknownValue:
					// nop
				case UInt64Value:
					n.val = c.convert(x-y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case Int64Value:
				switch y := c.convert(n.MultiplicativeExpression.eval(c, mode), n.Type()).(type) {
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

func (n *MultiplicativeExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case MultiplicativeExpressionCast: // CastExpression
			n.val = n.CastExpression.eval(c, mode)
		case MultiplicativeExpressionMul: // MultiplicativeExpression '*' CastExpression
			switch x := c.convert(n.MultiplicativeExpression.eval(c, mode), n.Type()).(type) {
			case unknownValue:
				// nop
			case UInt64Value:
				switch y := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
				case unknownValue:
					// nop
				case UInt64Value:
					n.val = c.convert(x*y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case Int64Value:
				switch y := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
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
			switch x := c.convert(n.MultiplicativeExpression.eval(c, mode), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				switch y := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
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
				switch y := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
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
			switch x := c.convert(n.MultiplicativeExpression.eval(c, mode), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				switch y := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
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

func (n *CastExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			switch n.Case {
			case CastExpressionUnary: // UnaryExpression
				n.val = n.UnaryExpression.eval(c, mode)
			case CastExpressionCast: // '(' TypeName ')' CastExpression
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			default:
				c.errors.add(errorf("internal error: %v", n.Case))
			}
			return n.Value()
		}

		switch n.Case {
		case CastExpressionUnary: // UnaryExpression
			n.val = n.UnaryExpression.eval(c, mode)
		case CastExpressionCast: // '(' TypeName ')' CastExpression
			n.val = c.convert(n.CastExpression.eval(c, mode), n.TypeName.Type())
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *UnaryExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case UnaryExpressionPostfix: // PostfixExpression
			n.val = n.PostfixExpression.eval(c, mode)
		case UnaryExpressionInc: // "++" UnaryExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionDec: // "--" UnaryExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionAddrof: // '&' CastExpression
			n.val = c.convert(n.CastExpression.eval(c, mode.add(addrOf)), n.Type())
		case UnaryExpressionDeref: // '*' CastExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionPlus: // '+' CastExpression
			switch x := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
			case unknownValue:
				// nop
			case Int64Value:
				n.val = c.convert(x, n.Type())
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case UnaryExpressionMinus: // '-' CastExpression
			switch x := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
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
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionNot: // '!' CastExpression
			switch x := n.CastExpression.eval(c, mode).(type) {
			case unknownValue:
				// nop
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case UnaryExpressionSizeofExpr: // "sizeof" UnaryExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionSizeofType: // "sizeof" '(' TypeName ')'
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionLabelAddr: // "&&" IDENTIFIER
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionAlignofExpr: // "_Alignof" UnaryExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionAlignofType: // "_Alignof" '(' TypeName ')'
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionImag: // "__imag__" UnaryExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionReal: // "__real__" UnaryExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *PostfixExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			switch n.Case {
			case PostfixExpressionPrimary: // PrimaryExpression
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PostfixExpressionIndex: // PostfixExpression '[' Expression ']'
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PostfixExpressionCall: // PostfixExpression '(' ArgumentExpressionList ')'
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PostfixExpressionSelect: // PostfixExpression '.' IDENTIFIER
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PostfixExpressionPSelect: // PostfixExpression "->" IDENTIFIER
				if !isPointerType(n.PostfixExpression.Type()) {
					break
				}

				switch x := n.PostfixExpression.eval(c, mode.del(addrOf)).(type) {
				case UInt64Value:
					switch y := n.PostfixExpression.Type().(*PointerType).Elem().(type) {
					case *StructType:
						if f := y.Field(n.Token2.SrcStr()); f != nil {
							n.val = c.convert(x+UInt64Value(f.Offset()), n.Type())
						}
					default:
						c.errors.add(errorf("TODO %v %v %T %v: %T", n.Case, mode.has(addrOf), x, n.Token.Position(), y))
					}
				default:
					c.errors.add(errorf("TODO %v %v %T %v:", n.Case, mode.has(addrOf), x, n.Token.Position()))
				}
			case PostfixExpressionInc: // PostfixExpression "++"
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PostfixExpressionDec: // PostfixExpression "--"
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PostfixExpressionComplit: // '(' TypeName ')' '{' InitializerList ',' '}'
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			default:
				c.errors.add(errorf("internal error: %v", n.Case))
			}
			return n.Value()
		}

		switch n.Case {
		case PostfixExpressionPrimary: // PrimaryExpression
			n.val = n.PrimaryExpression.eval(c, mode)
		case PostfixExpressionIndex: // PostfixExpression '[' Expression ']'
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case PostfixExpressionCall: // PostfixExpression '(' ArgumentExpressionList ')'
			// nop
		case PostfixExpressionSelect: // PostfixExpression '.' IDENTIFIER
			// nop
		case PostfixExpressionPSelect: // PostfixExpression "->" IDENTIFIER
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case PostfixExpressionInc: // PostfixExpression "++"
			// nop
		case PostfixExpressionDec: // PostfixExpression "--"
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case PostfixExpressionComplit: // '(' TypeName ')' '{' InitializerList ',' '}'
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *PrimaryExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			switch n.Case {
			case PrimaryExpressionIdent: // IDENTIFIER
				// nop
			case PrimaryExpressionInt: // INTCONST
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PrimaryExpressionFloat: // FLOATCONST
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PrimaryExpressionChar: // CHARCONST
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PrimaryExpressionLChar: // LONGCHARCONST
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PrimaryExpressionString: // STRINGLITERAL
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PrimaryExpressionLString: // LONGSTRINGLITERAL
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PrimaryExpressionExpr: // '(' Expression ')'
				n.val = n.ExpressionList.eval(c, mode)
			case PrimaryExpressionStmt: // '(' CompoundStatement ')'
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			case PrimaryExpressionGeneric: // GenericSelection
				c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			default:
				c.errors.add(errorf("internal error: %v", n.Case))
			}
			return n.Value()
		}

		switch n.Case {
		case PrimaryExpressionIdent: // IDENTIFIER
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case PrimaryExpressionInt: // INTCONST
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case PrimaryExpressionFloat: // FLOATCONST
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case PrimaryExpressionChar: // CHARCONST
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case PrimaryExpressionLChar: // LONGCHARCONST
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case PrimaryExpressionString: // STRINGLITERAL
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case PrimaryExpressionLString: // LONGSTRINGLITERAL
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case PrimaryExpressionExpr: // '(' Expression ')'
			n.val = n.ExpressionList.eval(c, mode)
		case PrimaryExpressionStmt: // '(' CompoundStatement ')'
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case PrimaryExpressionGeneric: // GenericSelection
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *ExpressionList) eval(c *ctx, mode flags) (r Value) {
	n0 := n
	if n.val == nil {
		n.val = UnknownValue
		for ; n != nil; n = n.ExpressionList {
			n0.typ = n.AssignmentExpression.Type()
			n0.val = n.AssignmentExpression.eval(c, mode)
		}
	}
	return n0.Value()
}

func (n *AssignmentExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = UnknownValue
		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case AssignmentExpressionCond: // ConditionalExpression
			n.val = n.ConditionalExpression.eval(c, mode)
		case AssignmentExpressionAssign: // UnaryExpression '=' AssignmentExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case AssignmentExpressionMul: // UnaryExpression "*=" AssignmentExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case AssignmentExpressionDiv: // UnaryExpression "/=" AssignmentExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case AssignmentExpressionMod: // UnaryExpression "%=" AssignmentExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case AssignmentExpressionAdd: // UnaryExpression "+=" AssignmentExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case AssignmentExpressionSub: // UnaryExpression "-=" AssignmentExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case AssignmentExpressionLsh: // UnaryExpression "<<=" AssignmentExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case AssignmentExpressionRsh: // UnaryExpression ">>=" AssignmentExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case AssignmentExpressionAnd: // UnaryExpression "&=" AssignmentExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case AssignmentExpressionXor: // UnaryExpression "^=" AssignmentExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case AssignmentExpressionOr: // UnaryExpression "|=" AssignmentExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
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
