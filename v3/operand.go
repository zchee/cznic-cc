// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"fmt"
	"math"
)

var (
	_ Value = Complex128Value(0)
	_ Value = Float32Value(0)
	_ Value = Float64Value(0)
	_ Value = Int64Value(0)
	_ Value = StringValue(0)
	_ Value = Uint64Value(0)
	_ Value = WideStringValue(0)

	_ Operand = (*funcDesignator)(nil)
	_ Operand = (*lvalue)(nil)
	_ Operand = (*operand)(nil)
	_ Operand = noOperand

	noOperand = &operand{typ: noType}
)

type Operand interface {
	Declarator() *Declarator
	IsLValue() bool
	IsZero() bool
	Type() Type
	Value() Value
	convertTo(*context, Node, Type) Operand
	integerPromotion(*context, Node) Operand
	normalize(*context) Operand
}

type Value interface {
	add(b Value) Value
	and(b Value) Value
	div(b Value) Value
	isZero() bool
	lsh(b Value) Value
	mod(b Value) Value
	mul(b Value) Value
	neg() Value
	or(b Value) Value
	rsh(b Value) Value
	sub(b Value) Value
	xor(b Value) Value
	//TODO all comparisons
}

type WideStringValue StringID

func (v WideStringValue) add(b Value) Value { panic("internal error") } //TODOOK
func (v WideStringValue) and(b Value) Value { panic("internal error") } //TODOOK
func (v WideStringValue) div(b Value) Value { panic("internal error") } //TODOOK
func (v WideStringValue) isZero() bool      { return v == 0 }
func (v WideStringValue) lsh(b Value) Value { panic("internal error") } //TODOOK
func (v WideStringValue) mod(b Value) Value { panic("internal error") } //TODOOK
func (v WideStringValue) mul(b Value) Value { panic("internal error") } //TODOOK
func (v WideStringValue) neg() Value        { panic("internal error") } //TODOOK
func (v WideStringValue) or(b Value) Value  { panic("internal error") } //TODOOK
func (v WideStringValue) rsh(b Value) Value { panic("internal error") } //TODOOK
func (v WideStringValue) sub(b Value) Value { panic("internal error") } //TODOOK
func (v WideStringValue) xor(b Value) Value { panic("internal error") } //TODOOK

type StringValue StringID

func (v StringValue) add(b Value) Value { panic("internal error") } //TODOOK
func (v StringValue) and(b Value) Value { panic("internal error") } //TODOOK
func (v StringValue) div(b Value) Value { panic("internal error") } //TODOOK
func (v StringValue) isZero() bool      { return v == 0 }
func (v StringValue) lsh(b Value) Value { panic("internal error") } //TODOOK
func (v StringValue) mod(b Value) Value { panic("internal error") } //TODOOK
func (v StringValue) mul(b Value) Value { panic("internal error") } //TODOOK
func (v StringValue) neg() Value        { panic("internal error") } //TODOOK
func (v StringValue) or(b Value) Value  { panic("internal error") } //TODOOK
func (v StringValue) rsh(b Value) Value { panic("internal error") } //TODOOK
func (v StringValue) sub(b Value) Value { panic("internal error") } //TODOOK
func (v StringValue) xor(b Value) Value { panic("internal error") } //TODOOK

type Int64Value int64

func (v Int64Value) add(b Value) Value { return v + b.(Int64Value) }
func (v Int64Value) and(b Value) Value { return v & b.(Int64Value) }
func (v Int64Value) isZero() bool      { return v == 0 }
func (v Int64Value) mul(b Value) Value { return v * b.(Int64Value) }
func (v Int64Value) neg() Value        { return -v }
func (v Int64Value) or(b Value) Value  { return v | b.(Int64Value) }
func (v Int64Value) sub(b Value) Value { return v - b.(Int64Value) }
func (v Int64Value) xor(b Value) Value { return v ^ b.(Int64Value) }

func (v Int64Value) div(b Value) Value {
	if b.isZero() {
		return nil
	}

	return v / b.(Int64Value)
}

func (v Int64Value) lsh(b Value) Value {
	switch y := b.(type) {
	case Int64Value:
		return v << uint64(y)
	case Uint64Value:
		return v << y
	default:
		panic("internal error") //TODOOK
	}
}

func (v Int64Value) rsh(b Value) Value {
	switch y := b.(type) {
	case Int64Value:
		return v >> uint64(y)
	case Uint64Value:
		return v >> y
	default:
		panic("internal error") //TODOOK
	}
}

func (v Int64Value) mod(b Value) Value {
	if b.isZero() {
		return nil
	}

	return v % b.(Int64Value)
}

type Uint64Value uint64

func (v Uint64Value) add(b Value) Value { return v + b.(Uint64Value) }
func (v Uint64Value) and(b Value) Value { return v & b.(Uint64Value) }
func (v Uint64Value) isZero() bool      { return v == 0 }
func (v Uint64Value) mul(b Value) Value { return v * b.(Uint64Value) }
func (v Uint64Value) neg() Value        { return -v }
func (v Uint64Value) or(b Value) Value  { return v | b.(Uint64Value) }
func (v Uint64Value) sub(b Value) Value { return v - b.(Uint64Value) }
func (v Uint64Value) xor(b Value) Value { return v ^ b.(Uint64Value) }

func (v Uint64Value) div(b Value) Value {
	if b.isZero() {
		return nil
	}

	return v / b.(Uint64Value)
}

func (v Uint64Value) lsh(b Value) Value {
	switch y := b.(type) {
	case Int64Value:
		return v << uint64(y)
	case Uint64Value:
		return v << y
	default:
		panic("internal error") //TODOOK
	}
}

func (v Uint64Value) rsh(b Value) Value {
	switch y := b.(type) {
	case Int64Value:
		return v >> uint64(y)
	case Uint64Value:
		return v >> y
	default:
		panic("internal error") //TODOOK
	}
}

func (v Uint64Value) mod(b Value) Value {
	if b.isZero() {
		return nil
	}

	return v % b.(Uint64Value)
}

type Float32Value float32

func (v Float32Value) add(b Value) Value { return v + b.(Float32Value) }
func (v Float32Value) and(b Value) Value { panic("internal error") } //TODOOK
func (v Float32Value) div(b Value) Value { return v / b.(Float32Value) }
func (v Float32Value) isZero() bool      { return v == 0 }
func (v Float32Value) lsh(b Value) Value { panic("internal error") } //TODOOK
func (v Float32Value) mod(b Value) Value { panic("internal error") } //TODOOK
func (v Float32Value) mul(b Value) Value { return v * b.(Float32Value) }
func (v Float32Value) neg() Value        { return -v }
func (v Float32Value) or(b Value) Value  { panic("internal error") } //TODOOK
func (v Float32Value) rsh(b Value) Value { panic("internal error") } //TODOOK
func (v Float32Value) sub(b Value) Value { return v - b.(Float32Value) }
func (v Float32Value) xor(b Value) Value { panic("internal error") } //TODOOK

type Float64Value float64

func (v Float64Value) add(b Value) Value { return v + b.(Float64Value) }
func (v Float64Value) and(b Value) Value { panic("internal error") } //TODOOK
func (v Float64Value) div(b Value) Value { return v / b.(Float64Value) }
func (v Float64Value) isZero() bool      { return v == 0 }
func (v Float64Value) lsh(b Value) Value { panic("internal error") } //TODOOK
func (v Float64Value) mod(b Value) Value { panic("internal error") } //TODOOK
func (v Float64Value) mul(b Value) Value { return v * b.(Float64Value) }
func (v Float64Value) neg() Value        { return -v }
func (v Float64Value) or(b Value) Value  { panic("internal error") } //TODOOK
func (v Float64Value) rsh(b Value) Value { panic("internal error") } //TODOOK
func (v Float64Value) sub(b Value) Value { return v - b.(Float64Value) }
func (v Float64Value) xor(b Value) Value { panic("internal error") } //TODOOK

type Complex128Value complex128

func (v Complex128Value) add(b Value) Value { return v + b.(Complex128Value) }
func (v Complex128Value) and(b Value) Value { panic("internal error") } //TODOOK
func (v Complex128Value) div(b Value) Value { return v / b.(Complex128Value) }
func (v Complex128Value) isZero() bool      { return v == 0 }
func (v Complex128Value) lsh(b Value) Value { panic("internal error") } //TODOOK
func (v Complex128Value) mod(b Value) Value { panic("internal error") } //TODOOK
func (v Complex128Value) mul(b Value) Value { return v * b.(Complex128Value) }
func (v Complex128Value) neg() Value        { return -v }
func (v Complex128Value) or(b Value) Value  { panic("internal error") } //TODOOK
func (v Complex128Value) rsh(b Value) Value { panic("internal error") } //TODOOK
func (v Complex128Value) sub(b Value) Value { return v - b.(Complex128Value) }
func (v Complex128Value) xor(b Value) Value { panic("internal error") } //TODOOK

type lvalue struct {
	Operand
	declarator *Declarator
}

func (o *lvalue) Declarator() *Declarator { return o.declarator }
func (o *lvalue) IsLValue() bool          { return true }

type funcDesignator struct {
	Operand
	declarator *Declarator
}

func (o *funcDesignator) Declarator() *Declarator { return o.declarator }
func (o *funcDesignator) IsLValue() bool          { return false }

type operand struct {
	typ   Type
	value Value

	//TODO isLvalue bool or wrapper type
}

func (o *operand) Declarator() *Declarator { return nil }
func (o *operand) IsLValue() bool          { return false }
func (o *operand) IsZero() bool            { return o.value != nil && o.value.isZero() }
func (o *operand) Type() Type              { return o.typ }
func (o *operand) Value() Value            { return o.value }

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
	if a.Type().Kind() == Invalid || b.Type().Kind() == Invalid {
		return noOperand, noOperand
	}

	if !a.Type().IsArithmeticType() || !b.Type().IsArithmeticType() {
		panic("internal error") //TODOOK
	}

	if a.Type() == nil || b.Type() == nil {
		return a, b
	}

	a = a.normalize(ctx)
	b = b.normalize(ctx)

	at := a.Type()
	bt := b.Type()

	// First, if the corresponding real type of either operand is long
	// double, the other operand is converted, without change of type
	// domain, to a type whose corresponding real type is long double.
	if at.Kind() == ComplexLongDouble || bt.Kind() == ComplexLongDouble {
		return noOperand, noOperand //TODO
	}

	if at.Kind() == LongDouble || bt.Kind() == LongDouble {
		return noOperand, noOperand //TODO
	}

	// Otherwise, if the corresponding real type of either operand is
	// double, the other operand is converted, without change of type
	// domain, to a type whose corresponding real type is double.
	if at.Kind() == ComplexDouble || bt.Kind() == ComplexDouble {
		return noOperand, noOperand //TODO
	}

	if at.Kind() == Double || bt.Kind() == Double {
		return a.convertTo(ctx, n, ctx.cfg.ABI.Type(Double)), b.convertTo(ctx, n, ctx.cfg.ABI.Type(Double))
	}

	// Otherwise, if the corresponding real type of either operand is
	// float, the other operand is converted, without change of type
	// domain, to a type whose corresponding real type is float.
	if at.Kind() == ComplexFloat || bt.Kind() == ComplexFloat {
		return noOperand, noOperand //TODO
	}

	if at.Kind() == Float || bt.Kind() == Float {
		return a.convertTo(ctx, n, ctx.cfg.ABI.Type(Float)), b.convertTo(ctx, n, ctx.cfg.ABI.Type(Float))
	}

	if !a.Type().IsIntegerType() || !b.Type().IsIntegerType() {
		panic("internal error") //TODOOK
	}

	// Otherwise, the integer promotions are performed on both operands.
	a = a.integerPromotion(ctx, n)
	b = b.integerPromotion(ctx, n)
	at = a.Type()
	bt = b.Type()

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
	if abi.isSignedInteger(at.Kind()) == abi.isSignedInteger(bt.Kind()) {
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
	switch {
	case a.Type().IsSignedType(): // b is unsigned
		if intConvRank[bt.Kind()] >= intConvRank[a.Type().Kind()] {
			return a.convertTo(ctx, n, b.Type()), b
		}
	case b.Type().IsSignedType(): // a is unsigned
		if intConvRank[at.Kind()] >= intConvRank[b.Type().Kind()] {
			return a, b.convertTo(ctx, n, a.Type())
		}
	default:
		panic(fmt.Errorf("TODO %v %v", a, b))
	}

	// Otherwise, if the type of the operand with signed integer type can
	// represent all of the values of the type of the operand with unsigned
	// integer type, then the operand with unsigned integer type is
	// converted to the type of the operand with signed integer type.
	var signed Type
	switch {
	case abi.isSignedInteger(at.Kind()): // b is unsigned
		signed = a.Type()
		if intConvRank[bt.Kind()] >= intConvRank[at.Kind()] {
			return a.convertTo(ctx, n, b.Type()), b
		}
	case abi.isSignedInteger(bt.Kind()): // a is unsigned
		signed = b.Type()
		if intConvRank[at.Kind()] >= intConvRank[bt.Kind()] {
			return a, b.convertTo(ctx, n, a.Type())
		}

	}

	// Otherwise, both operands are converted to the unsigned integer type
	// corresponding to the type of the operand with signed integer type.
	var typ Type
	switch signed.Kind() {
	case Int:
		//TODO if a.IsEnumConst || b.IsEnumConst {
		//TODO 	return a, b
		//TODO }

		typ = abi.Type(UInt)
	case Long:
		typ = abi.Type(ULong)
	case LongLong:
		typ = abi.Type(ULongLong)
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
	t := o.Type()
	if t2 := integerPromotion(ctx, t); t2.Kind() != t.Kind() {
		return o.convertTo(ctx, n, t2)
	}

	return o
}

// [0]6.3.1.1-2
//
// If an int can represent all values of the original type, the value is
// converted to an int; otherwise, it is converted to an unsigned int. These
// are called the integer promotions. All other types are unchanged by the
// integer promotions.
func integerPromotion(ctx *context, t Type) Type {
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
	if t.IsBitFieldType() {
		f := t.BitField()
		intBits := int(ctx.cfg.ABI.Types[Int].Size) * 8
		switch {
		case t.IsSignedType():
			if f.BitFieldWidth() < intBits-1 {
				return ctx.cfg.ABI.Type(Int)
			}
		default:
			if f.BitFieldOffset() < intBits {
				return ctx.cfg.ABI.Type(Int)
			}
		}
		return t
	}

	switch t.Kind() {
	case Invalid:
		return t
	case Char, SChar, UChar, Short, UShort:
		return ctx.cfg.ABI.Type(Int)
	default:
		return t
	}
}

func (o *operand) convertTo(ctx *context, n Node, t Type) (r Operand) {
	if o.Type() == nil {
		return &operand{typ: t}
	}

	abi := ctx.cfg.ABI
	k0 := o.Type().Kind()
	if o.Value() == nil {
		return &operand{typ: t}
	}

	k := t.Kind()
	if k == Void {
		return &operand{typ: abi.Type(Void)} //TODO ABI singleton
	}

	if !integerTypes[k0] && k0 != Ptr {
		return &operand{typ: t}
	}

	if integerTypes[k] {
		var i64 int64
		switch x := o.Value().(type) {
		case Int64Value:
			i64 = int64(x)
		case Uint64Value:
			i64 = int64(x)
		case StringValue:
			if x != 0 {
				panic("internal error") //TODOOK
			}
		default:
			panic(fmt.Sprintf("%v: internal error: %T %v", n.Position(), x, x)) //TODOOK
		}
		var v Value
		switch {
		case abi.isSignedInteger(k):
			v = Int64Value(i64)
		default:
			v = Uint64Value(i64)
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
			return &operand{typ: t, value: Uint64Value(0)}
		}
	}
	return &operand{typ: t}
}

func (o *operand) normalize(ctx *context) (r Operand) {
	switch x := o.Value().(type) {
	case Int64Value:
		if v := convertInt64(int64(x), o.Type(), ctx); v != int64(x) {
			return &operand{o.Type(), Int64Value(v)}
		}
	case Uint64Value:
		v := uint64(x)
		switch o.Type().Size() {
		case 1:
			v &= math.MaxUint8
		case 2:
			v &= math.MaxUint16
		case 4:
			v &= math.MaxUint32
		}
		if v != uint64(x) {
			return &operand{o.Type(), Uint64Value(v)}
		}
	}
	return o
}

func convertInt64(n int64, t Type, ctx *context) int64 {
	abi := ctx.cfg.ABI
	k := t.Kind()
	if k == Enum {
		//TODO
	}
	signed := abi.isSignedInteger(k)
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
