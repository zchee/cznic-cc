// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"math"
)

var (
	_ Value = Float64Value(0)
	_ Value = Int64Value(0)
	_ Value = UInt64Value(0)

	_ Operand = (*operand)(nil)

	noOperand = &operand{typ: noType}
)

type Operand interface {
	Type() Type
	Value() Value
	convertTo(*context, Node, Type) Operand
	integerPromotion(*context, Node) Operand
	isArithmeticType() bool
	isIntegerType() bool
	normalize(*context) Operand
}

type Value interface {
	add(b Value) Value
	div(b Value) Value
	isZero() bool
	mod(b Value) Value
	mul(b Value) Value
	sub(b Value) Value
}

type Int64Value int64

func (v Int64Value) add(b Value) Value { return v + b.(Int64Value) }
func (v Int64Value) isZero() bool      { return v == 0 }
func (v Int64Value) mul(b Value) Value { return v * b.(Int64Value) }
func (v Int64Value) sub(b Value) Value { return v - b.(Int64Value) }

func (v Int64Value) div(b Value) Value {
	if b.isZero() {
		return nil
	}

	return v / b.(Int64Value)
}

func (v Int64Value) mod(b Value) Value {
	if b.isZero() {
		return nil
	}

	return v % b.(Int64Value)
}

type UInt64Value uint64

func (v UInt64Value) add(b Value) Value { return v + b.(UInt64Value) }
func (v UInt64Value) isZero() bool      { return v == 0 }
func (v UInt64Value) mul(b Value) Value { return v * b.(UInt64Value) }
func (v UInt64Value) sub(b Value) Value { return v - b.(UInt64Value) }

func (v UInt64Value) div(b Value) Value {
	if b.isZero() {
		return nil
	}

	return v / b.(UInt64Value)
}

func (v UInt64Value) mod(b Value) Value {
	if b.isZero() {
		return nil
	}

	return v % b.(UInt64Value)
}

type Float64Value float64

func (v Float64Value) add(b Value) Value { return v + b.(Float64Value) }
func (v Float64Value) div(b Value) Value { return v / b.(Float64Value) }
func (v Float64Value) isZero() bool      { return v == 0 }
func (v Float64Value) mod(b Value) Value { panic("internal error") } //TODOOK
func (v Float64Value) mul(b Value) Value { return v * b.(Float64Value) }
func (v Float64Value) sub(b Value) Value { return v - b.(Float64Value) }

type operand struct {
	typ   Type
	value Value
}

func (o *operand) Type() Type   { return o.typ }
func (o *operand) Value() Value { return o.value }

func (o *operand) isArithmeticType() bool {
	return o.Type() != nil && isArithmeticType[o.Type().underlyingType().Kind()]
}

func (o *operand) isIntegerType() bool {
	return o.Type() != nil && integerTypes[o.Type().underlyingType().Kind()]
}

// [0]6.3.1.8
//
// Many operators that expect operands of arithmetic type cause conversions and
// yield result types in a similar way. The purpose is to determine a common
// real type for the operands and result. For the specified operands, each
// operand is converted, without change of type domain, to a type whose
// corresponding real type is the common real type. Unless explicitly stated
// otherwise, the common real type is also the corresponding real type of the
// result, whose type domain is the type domain of the operands if they are the
// same, and complex otherwise. This pattern is called the usual arithmetic
// conversions:
func usualArithmeticConversions(ctx *context, n Node, a, b Operand) (Operand, Operand) {
	if !a.isArithmeticType() || !b.isArithmeticType() {
		// dbg("", PrettyString(n))
		panic("internal error") //TODOOK
	}

	if a.Type() == nil || b.Type() == nil {
		return a, b
	}

	a = a.normalize(ctx)
	b = b.normalize(ctx)

	at := a.Type().underlyingType()
	bt := b.Type().underlyingType()

	// First, if the corresponding real type of either operand is long
	// double, the other operand is converted, without change of type
	// domain, to a type whose corresponding real type is long double.
	if at.Kind() == LongDouble || bt.Kind() == LongDouble {
		panic("TODO")
	}

	// Otherwise, if the corresponding real type of either operand is
	// double, the other operand is converted, without change of type
	// domain, to a type whose corresponding real type is double.
	if at.Kind() == Double || bt.Kind() == Double {
		panic("TODO")
	}

	// Otherwise, if the corresponding real type of either operand is
	// float, the other operand is converted, without change of type
	// domain, to a type whose corresponding real type is float.
	if at.Kind() == Float || bt.Kind() == Float {
		panic("TODO")
	}

	if !a.isIntegerType() || !b.isIntegerType() {
		panic("internal error") //TODOOK
	}

	// Otherwise, the integer promotions are performed on both operands.
	a = a.integerPromotion(ctx, n)
	b = b.integerPromotion(ctx, n)
	at = a.Type().underlyingType()
	bt = b.Type().underlyingType()

	// Then the following rules are applied to the promoted operands:

	// If both operands have the same type, then no further conversion is
	// needed.
	if at.Kind() == bt.Kind() {
		return a, b
	}

	// Otherwise, if both operands have signed integer types or both have
	// unsigned integer types, the operand with the type of lesser integer
	// conversion rank is converted to the type of the operand with greater
	// rank.
	abi := ctx.cfg.ABI
	if abi.isSigned(at.Kind()) == abi.isSigned(bt.Kind()) {
		t := a.Type()
		if intConvRank[bt.Kind()] > intConvRank[at.Kind()] {
			t = b.Type()
		}
		return a.convertTo(ctx, n, t), b.convertTo(ctx, n, t)

	}

	// Otherwise, if the operand that has unsigned integer type has rank
	// greater or equal to the rank of the type of the other operand, then
	// the operand with signed integer type is converted to the type of the
	// operand with unsigned integer type.
	var signed Type
	switch {
	case abi.isSigned(at.Kind()): // b is unsigned
		signed = a.Type()
		if intConvRank[bt.Kind()] >= intConvRank[at.Kind()] {
			return a.convertTo(ctx, n, b.Type()), b
		}
	case abi.isSigned(bt.Kind()): // a is unsigned
		signed = b.Type()
		if intConvRank[at.Kind()] >= intConvRank[bt.Kind()] {
			return a, b.convertTo(ctx, n, a.Type())
		}

	}

	// Otherwise, both operands are converted to the unsigned integer type
	// corresponding to the type of the operand with signed integer type.
	var typ Type
	switch signed.underlyingType().Kind() {
	case Int:
		//TODO if a.IsEnumConst || b.IsEnumConst {
		//TODO 	return a, b
		//TODO }

		typ = abi.typ(ctx, n, UInt)
	case Long:
		typ = abi.typ(ctx, n, ULong)
	case LongLong:
		typ = abi.typ(ctx, n, ULongLong)
	default:
		panic("internal error") //TODOOK
	}
	return a.convertTo(ctx, n, typ), b.convertTo(ctx, n, typ)
}

// [0]6.3.1.1-2
//
// If an int can represent all values of the original type, the value is
// converted to an int; otherwise, it is converted to an unsigned int. These
// are called the integer promotions. All other types are unchanged by the
// integer promotions.
func (o *operand) integerPromotion(ctx *context, n Node) Operand {
	//TODO
	// github.com/gcc-mirror/gcc/gcc/testsuite/gcc.c-torture/execute/bf-sign-2.c
	//
	// This test checks promotion of bitfields.  Bitfields
	// should be promoted very much like chars and shorts:
	//
	// Bitfields (signed or unsigned) should be promoted to
	// signed int if their value will fit in a signed int,
	// otherwise to an unsigned int if their value will fit
	// in an unsigned int, otherwise we don't promote them
	// (ANSI/ISO does not specify the behavior of bitfields
	// larger than an unsigned int).

	t := o.Type()
	if t == nil {
		return o
	}

	switch t.underlyingType().Kind() {
	case
		Char,
		SChar,
		Short,
		UChar,
		UShort:
		return o.convertTo(ctx, n, ctx.cfg.ABI.typ(ctx, n, Int))
	default:
		return o
	}
}

func (o *operand) convertTo(ctx *context, n Node, t Type) (r Operand) {
	if o.Type() == nil {
		return &operand{typ: t}
	}

	abi := ctx.cfg.ABI
	k0 := o.Type().underlyingType().Kind()
	if o.Value() == nil || !abi.isInt(k0) {
		return &operand{typ: t}
	}

	k := t.underlyingType().Kind()
	if k == Void {
		return &operand{typ: abi.typ(ctx, n, Void)} //TODO ABI singleton
	}

	if abi.isInt(k) {
		var i64 int64
		switch x := o.Value().(type) {
		case Int64Value:
			i64 = int64(x)
		case UInt64Value:
			i64 = int64(x)
		default:
			panic("internal error") //TODOOK
		}
		var v Value
		switch {
		case abi.isSigned(k):
			v = Int64Value(i64)
		default:
			v = UInt64Value(i64)
		}
		return (&operand{typ: t, value: v}).normalize(ctx)
	}

	if k == Ptr {
		// [0]6.3.2.3
		if o.Value().isZero() {
			// 3. An integer constant expression with the
			// value 0, or such an expression cast to type
			// void *, is called a null pointer constant.
			// If a null pointer constant is converted to a
			// pointer type, the resulting pointer, called
			// a null pointer, is guaranteed to compare
			// unequal to a pointer to any object or
			// function.
			return &operand{typ: t, value: UInt64Value(0)}
		}
	}
	return &operand{typ: t}
}

func (o *operand) normalize(ctx *context) Operand {
	switch x := o.Value().(type) {
	case Int64Value:
		if v := convertInt64(int64(x), o.Type(), ctx); v != int64(x) {
			return &operand{o.Type(), Int64Value(v)}
		}
	}
	return o
}

func convertInt64(n int64, t Type, ctx *context) int64 {
	abi := ctx.cfg.ABI
	k := t.underlyingType().Kind()
	if k == Enum {
		panic("TODO")
	}
	signed := abi.isSigned(k)
	switch sz := abi.size(k); sz {
	case 1:
		switch {
		case signed:
			switch {
			case int8(n) < 0:
				return n | ^math.MaxUint8
			default:
				return n & math.MaxUint8
			}
		default:
			return n & math.MaxUint8
		}
	case 2:
		switch {
		case signed:
			switch {
			case int16(n) < 0:
				return n | ^math.MaxUint16
			default:
				return n & math.MaxUint16
			}
		default:
			return n & math.MaxUint16
		}
	case 4:
		switch {
		case signed:
			switch {
			case int32(n) < 0:
				return n | ^math.MaxUint32
			default:
				return n & math.MaxUint32
			}
		default:
			return n & math.MaxUint32
		}
	case 8:
		return n
	default:
		panic("internal error") //TODOOK
	}
}
