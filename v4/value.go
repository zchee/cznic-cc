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
	_ Value = (*UnknownValue)(nil)
	_ Value = (*ZeroValue)(nil)
	_ Value = Complex128Value(0)
	_ Value = Complex64Value(0)
	_ Value = Float64Value(0)
	_ Value = Int64Value(0)
	_ Value = StringValue("")
	_ Value = UInt64Value(0)
	_ Value = UTF16StringValue(nil)
	_ Value = UTF32StringValue(nil)
	_ Value = VoidValue{}
)

var (
	// Unknown is a singleton representing an undetermined value.  Unknown is
	// comparable.
	Unknown Value = &UnknownValue{}

	// Zero is a singleton representing a zero value of a type. Zero is comparable.
	Zero Value = &ZeroValue{}

	int1 = Int64Value(1)
	int0 = Int64Value(0)
)

type Value interface {
	isValue()
}

type UnknownValue struct{}

func (*UnknownValue) isValue() {}

func (*UnknownValue) String() string { return "<unknown value>" }

type ZeroValue struct{}

func (*ZeroValue) isValue() {}

func (*ZeroValue) String() string { return "{}" }

type ComplexLongDoubleValue struct {
	Re *big.Float
	Im *big.Float
}

func (n *ComplexLongDoubleValue) isValue() {}

type LongDoubleValue big.Float

func (*LongDoubleValue) isValue() {}

type Complex128Value complex128

func (Complex128Value) isValue() {}

type Complex64Value complex64

func (Complex64Value) isValue() {}

type Float64Value float64

func (Float64Value) isValue() {}

type Int64Value int64

func (Int64Value) isValue() {}

type UInt64Value int64

func (UInt64Value) isValue() {}

type VoidValue struct{}

func (VoidValue) isValue() {}

type StringValue string

func (StringValue) isValue() {}

type UTF16StringValue []uint16

func (UTF16StringValue) isValue() {}

type UTF32StringValue []rune

func (UTF32StringValue) isValue() {}

func (n *ConstantExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = n.ConditionalExpression.eval(c, mode)
	}
	return n.Value()
}

func (n *ConditionalExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

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
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case LogicalOrExpressionLAnd: // LogicalAndExpression
			n.val = n.LogicalAndExpression.eval(c, mode)
		case LogicalOrExpressionLOr: // LogicalOrExpression "||" LogicalAndExpression
			switch v := c.convert(n.LogicalOrExpression.eval(c, mode), n.Type()); {
			case isZero(v):
				switch v := c.convert(n.LogicalAndExpression.eval(c, mode), n.Type()); {
				case isZero(v):
					n.val = int0
				case isNonzero(v):
					n.val = int1
				}
			case isNonzero(v):
				n.val = int1
			}
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *LogicalAndExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

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
				n.val = int0
			case isNonzero(v):
				switch w := n.InclusiveOrExpression.eval(c, mode); {
				case isZero(w):
					n.val = int0
				case isNonzero(w):
					n.val = int1
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
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case InclusiveOrExpressionXor: // ExclusiveOrExpression
			n.val = n.ExclusiveOrExpression.eval(c, mode)
		case InclusiveOrExpressionOr: // InclusiveOrExpression '|' ExclusiveOrExpression
			switch x := c.convert(n.InclusiveOrExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// nop
			case Int64Value:
				switch y := c.convert(n.InclusiveOrExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x|y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case UInt64Value:
				switch y := c.convert(n.InclusiveOrExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// nop
				case UInt64Value:
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
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case ExclusiveOrExpressionAnd: // AndExpression
			n.val = n.AndExpression.eval(c, mode)
		case ExclusiveOrExpressionXor: // ExclusiveOrExpression '^' AndExpression
			switch x := c.convert(n.ExclusiveOrExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// ok
			case Int64Value:
				switch y := c.convert(n.AndExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// ok
				case Int64Value:
					n.val = c.convert(x^y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case UInt64Value:
				switch y := c.convert(n.AndExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// ok
				case UInt64Value:
					n.val = c.convert(x^y, n.Type())
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

func (n *AndExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case AndExpressionEq: // EqualityExpression
			n.val = n.EqualityExpression.eval(c, mode)
		case AndExpressionAnd: // AndExpression '&' EqualityExpression
			switch x := c.convert(n.AndExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// ok
			case Int64Value:
				switch y := c.convert(n.EqualityExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// ok
				case Int64Value:
					n.val = c.convert(x&y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case UInt64Value:
				switch y := c.convert(n.EqualityExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// ok
				case UInt64Value:
					n.val = c.convert(x&y, n.Type())
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

func (n *EqualityExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case EqualityExpressionRel: // RelationalExpression
			n.val = n.RelationalExpression.eval(c, mode)
		case EqualityExpressionEq: // EqualityExpression "==" RelationalExpression
			switch x := c.convert(n.EqualityExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// ok
			case Int64Value:
				switch y := c.convert(n.RelationalExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
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
			switch x := c.convert(n.EqualityExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// ok
			case Int64Value:
				switch y := c.convert(n.RelationalExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// ok
				case Int64Value:
					n.val = bool2int(x != y)
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

func (n *RelationalExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case RelationalExpressionShift: // ShiftExpression
			n.val = n.ShiftExpression.eval(c, mode)
		case RelationalExpressionLt: // RelationalExpression '<' ShiftExpression
			switch x := c.convert(n.RelationalExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// ok
			case Int64Value:
				switch y := c.convert(n.ShiftExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// ok
				case Int64Value:
					if x < y {
						n.val = int1
						break
					}

					n.val = int0
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case RelationalExpressionGt: // RelationalExpression '>' ShiftExpression
			switch x := c.convert(n.RelationalExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// ok
			case Int64Value:
				switch y := c.convert(n.ShiftExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// ok
				case Int64Value:
					if x > y {
						n.val = int1
						break
					}

					n.val = int0
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case RelationalExpressionLeq: // RelationalExpression "<=" ShiftExpression
			switch x := c.convert(n.RelationalExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// ok
			case Int64Value:
				switch y := c.convert(n.ShiftExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// ok
				case Int64Value:
					if x <= y {
						n.val = int1
						break
					}

					n.val = int0
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case RelationalExpressionGeq: // RelationalExpression ">=" ShiftExpression
			switch x := c.convert(n.RelationalExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// ok
			case Int64Value:
				switch y := c.convert(n.ShiftExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// ok
				case Int64Value:
					if x >= y {
						n.val = int1
						break
					}

					n.val = int0
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

func (n *ShiftExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case ShiftExpressionAdd: // AdditiveExpression
			n.val = n.AdditiveExpression.eval(c, mode)
		case ShiftExpressionLsh: // ShiftExpression "<<" AdditiveExpression
			switch x := c.convert(n.ShiftExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// nop
			case Int64Value:
				switch y := n.AdditiveExpression.eval(c, mode).(type) {
				case *UnknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x<<y, n.Type())
				case UInt64Value:
					n.val = c.convert(x<<y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case UInt64Value:
				switch y := n.AdditiveExpression.eval(c, mode).(type) {
				case *UnknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x<<y, n.Type())
				case UInt64Value:
					n.val = c.convert(x<<y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case ShiftExpressionRsh: // ShiftExpression ">>" AdditiveExpression
			switch x := c.convert(n.ShiftExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// nop
			case Int64Value:
				switch y := n.AdditiveExpression.eval(c, mode).(type) {
				case *UnknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x>>y, n.Type())
				case UInt64Value:
					n.val = c.convert(x>>y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case UInt64Value:
				switch y := n.AdditiveExpression.eval(c, mode).(type) {
				case *UnknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x>>y, n.Type())
				case UInt64Value:
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
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case AdditiveExpressionMul: // MultiplicativeExpression
			n.val = n.MultiplicativeExpression.eval(c, mode)
		case AdditiveExpressionAdd: // AdditiveExpression '+' MultiplicativeExpression
			switch x := c.convert(n.AdditiveExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// nop
			case Int64Value:
				switch y := c.convert(n.MultiplicativeExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// nop
				case Int64Value:
					n.val = c.convert(x+y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case UInt64Value:
				switch y := c.convert(n.MultiplicativeExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// nop
				case UInt64Value:
					n.val = c.convert(x+y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case Float64Value:
				// nop
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case AdditiveExpressionSub: // AdditiveExpression '-' MultiplicativeExpression
			switch x := c.convert(n.AdditiveExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// nop
			case UInt64Value:
				switch y := c.convert(n.MultiplicativeExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// nop
				case UInt64Value:
					n.val = c.convert(x-y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case Int64Value:
				switch y := c.convert(n.MultiplicativeExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
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
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case MultiplicativeExpressionCast: // CastExpression
			n.val = n.CastExpression.eval(c, mode)
		case MultiplicativeExpressionMul: // MultiplicativeExpression '*' CastExpression
			switch x := c.convert(n.MultiplicativeExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// nop
			case UInt64Value:
				switch y := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// nop
				case UInt64Value:
					n.val = c.convert(x*y, n.Type())
				default:
					c.errors.add(errorf("TODO %v TYPE %T", n.Case, y))
				}
			case Int64Value:
				switch y := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
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
			case *UnknownValue:
				// nop
			case Int64Value:
				switch y := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
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
				case *UnknownValue:
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
			case *UnknownValue:
				// nop
			case Int64Value:
				switch y := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
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
			case UInt64Value:
				switch y := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
				case *UnknownValue:
					// nop
				case UInt64Value:
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
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

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
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case UnaryExpressionPostfix: // PostfixExpression
			n.val = n.PostfixExpression.eval(c, mode)
		case UnaryExpressionInc: // "++" UnaryExpression
			// nop
		case UnaryExpressionDec: // "--" UnaryExpression
			// nop
		case UnaryExpressionAddrof: // '&' CastExpression
			n.val = c.convert(n.CastExpression.eval(c, mode.add(addrOf)), n.Type())
		case UnaryExpressionDeref: // '*' CastExpression
			// nop
		case UnaryExpressionPlus: // '+' CastExpression
			switch x := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// nop
			case Int64Value:
				n.val = c.convert(x, n.Type())
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case UnaryExpressionMinus: // '-' CastExpression
			switch x := c.convert(n.CastExpression.eval(c, mode), n.Type()).(type) {
			case *UnknownValue:
				// nop
			case Int64Value:
				n.val = c.convert(-x, n.Type())
			case UInt64Value:
				n.val = c.convert(-x, n.Type())
			case Float64Value:
				// nop
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case UnaryExpressionCpl: // '~' CastExpression
			switch x := n.CastExpression.eval(c, mode).(type) {
			case *UnknownValue:
				// nop
			case Int64Value:
				n.val = c.convert(^x, n.Type())
			case UInt64Value:
				n.val = c.convert(^x, n.Type())
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case UnaryExpressionNot: // '!' CastExpression
			switch x := n.CastExpression.eval(c, mode).(type) {
			case *UnknownValue:
				// nop
			case Int64Value:
				n.val = c.convert(bool2int(x == 0), n.Type())
			case UInt64Value:
				n.val = c.convert(bool2int(x == 0), n.Type())
			default:
				c.errors.add(errorf("TODO %v TYPE %T", n.Case, x))
			}
		case UnaryExpressionSizeofExpr: // "sizeof" UnaryExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionSizeofType: // "sizeof" '(' TypeName ')'
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionLabelAddr: // "&&" IDENTIFIER
			// nop
		case UnaryExpressionAlignofExpr: // "_Alignof" UnaryExpression
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionAlignofType: // "_Alignof" '(' TypeName ')'
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		case UnaryExpressionImag: // "__imag__" UnaryExpression
			// nop
		case UnaryExpressionReal: // "__real__" UnaryExpression
			// nop
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *PostfixExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

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
				switch x := n.PostfixExpression.eval(c, mode.del(addrOf)).(type) {
				case UInt64Value:
					switch y := n.PostfixExpression.Type().(*PointerType).Elem().(type) {
					case *StructType:
						if f := y.FieldByName(n.Token2.SrcStr()); f != nil {
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
			switch {
			case isPointerType(n.PostfixExpression.Type()) && isIntegerType(n.ExpressionList.Type()):
				switch v := n.PostfixExpression.eval(c, mode).(type) {
				case *UnknownValue:
					// nop
				case StringValue:
					switch x := n.ExpressionList.eval(c, 0).(type) {
					case *UnknownValue:
						// nop
					case Int64Value:
						if x >= 0 && x < Int64Value(len(v)) {
							n.val = c.convert(Int64Value(v[x]), n.Type())
						}
					case UInt64Value:
						if x < UInt64Value(len(v)) {
							n.val = c.convert(Int64Value(v[x]), n.Type())
						}
					default:
						c.errors.add(errorf("TODO %v %T", n.Case, x))
					}
				case UTF32StringValue:
					switch x := n.ExpressionList.eval(c, 0).(type) {
					case *UnknownValue:
						// nop
					case Int64Value:
						if x >= 0 && x < Int64Value(len(v)) {
							n.val = c.convert(Int64Value(v[x]), n.Type())
						}
					case UInt64Value:
						if x < UInt64Value(len(v)) {
							n.val = c.convert(Int64Value(v[x]), n.Type())
						}
					default:
						c.errors.add(errorf("TODO %v %T", n.Case, x))
					}
				case UInt64Value:
					// nop
				default:
					// trc("%v: %v %v [%v %v] %T", n.Token.Position(), n.PostfixExpression.Value(), n.PostfixExpression.Type(), n.ExpressionList.Value(), n.ExpressionList.Type(), v)
					c.errors.add(errorf("TODO %v %T", n.Case, v))
				}
			case isIntegerType(n.PostfixExpression.Type()) && isPointerType(n.ExpressionList.Type()):
				switch v := n.ExpressionList.eval(c, mode).(type) {
				case *UnknownValue:
					// nop
				case StringValue:
					switch x := n.PostfixExpression.eval(c, 0).(type) {
					case *UnknownValue:
						// nop
					case Int64Value:
						if x >= 0 && x < Int64Value(len(v)) {
							n.val = c.convert(Int64Value(v[x]), n.Type())
						}
					case UInt64Value:
						if x < UInt64Value(len(v)) {
							n.val = c.convert(Int64Value(v[x]), n.Type())
						}
					default:
						c.errors.add(errorf("TODO %v %T", n.Case, x))
					}
				default:
					// trc("%v: %v %v [%v %v] %T", n.Token.Position(), n.PostfixExpression.Value(), n.PostfixExpression.Type(), n.ExpressionList.Value(), n.ExpressionList.Type(), v)
					c.errors.add(errorf("TODO %v %T", n.Case, v))
				}
			}
		case PostfixExpressionCall: // PostfixExpression '(' ArgumentExpressionList ')'
			// nop
		case PostfixExpressionSelect: // PostfixExpression '.' IDENTIFIER
			// nop
		case PostfixExpressionPSelect: // PostfixExpression "->" IDENTIFIER
			// nop
		case PostfixExpressionInc: // PostfixExpression "++"
			// nop
		case PostfixExpressionDec: // PostfixExpression "--"
			// nop
		case PostfixExpressionComplit: // '(' TypeName ')' '{' InitializerList ',' '}'
			if n.InitializerList == nil || n.InitializerList.InitializerList != nil || n.InitializerList.Initializer.Case != InitializerExpr {
				break
			}

			v := c.convert(n.InitializerList.Initializer.AssignmentExpression.eval(c, mode), n.Type())
			switch n.Type().(type) {
			case *PredefinedType:
				n.val = v
			case *EnumType:
				n.val = v
			case *PointerType:
				n.val = v
			}
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *PrimaryExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

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
			switch n.resolvedTo.(type) {
			case *Declarator, *Parameter:
				// ok
			default:
				c.errors.add(errorf("TODO %v %T", n.Case, n.resolvedTo))
			}
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
			// ok
		case PrimaryExpressionGeneric: // GenericSelection
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func (n *ExpressionList) eval(c *ctx, mode flags) (r Value) {
	if n == nil {
		return Unknown
	}

	n0 := n
	if n.val == nil {
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

		for ; n != nil; n = n.ExpressionList {
			n0.typ = n.AssignmentExpression.Type()
			n0.val = n.AssignmentExpression.eval(c, mode)
		}
	}
	return n0.Value()
}

func (n *AssignmentExpression) eval(c *ctx, mode flags) (r Value) {
	if n.val == nil {
		n.val = Unknown
		if mode.has(dontEval) {
			return n.Value()
		}

		if mode.has(addrOf) {
			c.errors.add(errorf("TODO %v %v", n.Case, mode.has(addrOf)))
			return n.Value()
		}

		switch n.Case {
		case AssignmentExpressionCond: // ConditionalExpression
			n.val = n.ConditionalExpression.eval(c, mode)
		case AssignmentExpressionAssign: // UnaryExpression '=' AssignmentExpression
			n.val = c.convert(n.AssignmentExpression.eval(c, mode), n.UnaryExpression.Type())
		case AssignmentExpressionMul, // UnaryExpression "*=" AssignmentExpression
			AssignmentExpressionDiv, // UnaryExpression "/=" AssignmentExpression
			AssignmentExpressionMod, // UnaryExpression "%=" AssignmentExpression
			AssignmentExpressionAdd, // UnaryExpression "+=" AssignmentExpression
			AssignmentExpressionSub, // UnaryExpression "-=" AssignmentExpression
			AssignmentExpressionLsh, // UnaryExpression "<<=" AssignmentExpression
			AssignmentExpressionRsh, // UnaryExpression ">>=" AssignmentExpression
			AssignmentExpressionAnd, // UnaryExpression "&=" AssignmentExpression
			AssignmentExpressionXor, // UnaryExpression "^=" AssignmentExpression
			AssignmentExpressionOr:  // UnaryExpression "|=" AssignmentExpression

			n.UnaryExpression.eval(c, mode)
			n.AssignmentExpression.eval(c, mode)
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
	}
	return n.Value()
}

func isZero(v Value) bool {
	switch x := v.(type) {
	case *UnknownValue:
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
	case *UnknownValue:
		return false
	case Int64Value:
		return x != 0
	case UInt64Value:
		return x != 0
	default:
		panic(todo("%T", x))
	}
}
