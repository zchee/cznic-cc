// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

// Values of Kind
const (
	Invalid Kind = iota

	Bool              // _Bool
	Char              // char
	ComplexChar       // complex char
	ComplexDouble     // complex double
	ComplexFloat      // complex float
	ComplexInt        // complex int
	ComplexLong       // complex long
	ComplexLongDouble // complex long double
	ComplexLongLong   // complex long long
	ComplexShort      // complex short
	ComplexUInt       // complex unsigned
	ComplexUShort     // complex shor
	Decimal128        // _Decimal128
	Decimal32         // _Decimal32
	Decimal64         // _Decimal64
	Double            // double
	Enum              // enum
	Float             // float
	Float128          // _Float128
	Float32           // _Float32
	Float32x          // _Float32x
	Float64           // _Float64
	Float64x          // _Float64x
	Int               // int
	Int128            // __int128
	Long              // long
	LongDouble        // long double
	LongLong          // long long
	Ptr               // pointer
	SChar             // signed char
	Short             // chort
	Struct            // struct
	TypedefName       // typedef name
	UChar             // unsigned char
	UInt              // unsigned
	UInt128           // unsigned __int128
	ULong             // unsigned long
	ULongLong         // unsigned long long
	UShort            // unsigned short
	Void              // void

	typeofExpr
	typeofType
)
