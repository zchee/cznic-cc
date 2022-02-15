// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

var (
	_ Type = (*ArrayType)(nil)
	_ Type = (*EnumType)(nil)
	_ Type = (*FunctionType)(nil)
	_ Type = (*PointerType)(nil)
	_ Type = (*PredefinedType)(nil)
	_ Type = (*StructType)(nil)
	_ Type = (*UnionType)(nil)
)

// Type is the representation of a C type.
//
// Not all methods apply to all kinds of types. Restrictions, if any, are noted
// in the documentation for each method. Use the Kind method to find out the
// kind of type before calling kind-specific methods. Calling a method
// inappropriate to the kind of type causes a run-time panic.
//
// Calling a method on a type of kind Invalid yields an undefined result, but
// does not panic.
type Type interface {
	Kind() Kind
}

const (
	Array             Kind = iota // array
	Bool                          // _Bool
	Char                          // char
	ComplexChar                   // _Complex char
	ComplexDouble                 // _Complex double
	ComplexFloat                  // _Complex float
	ComplexInt                    // _Complex int
	ComplexLong                   // _Complex long
	ComplexLongDouble             // _Complex long double
	ComplexLongLong               // _Complex long long
	ComplexShort                  // _Complex short
	ComplexUInt                   // _Complex unsigned
	Double                        // double
	Enum                          // enum
	Float                         // float
	Float128                      // _Float128
	Float32                       // _Float32
	Float32x                      // _Float32x
	Float64                       // _Float64
	Float64x                      // _Float64x
	Function                      // function
	Int                           // int
	Int128                        // __int128
	Long                          // long
	LongDouble                    // long double
	LongLong                      // long long
	Ptr                           // pointer
	SChar                         // signed char
	Short                         // short
	Struct                        // struct
	UChar                         // unsigned char
	UInt                          // unsigned
	UInt128                       // unsigned __int128
	ULong                         // unsigned long
	ULongLong                     // unsigned long long
	UShort                        // unsigned short
	Union                         // union
	Void                          // void
)

type Kind int

type PredefinedType struct {
	ast  *AST
	kind Kind
}

func newPredefinedType(ast *AST, kind Kind) *PredefinedType { return &PredefinedType{kind: kind} }

func (n *PredefinedType) Kind() Kind { return n.kind }

type FunctionType struct {
	result  Type
	fp      []*ParameterDeclaration
	minArgs int
	maxArgs int // -1: unlimited
}

func newFunctionType(c *ctx, result Type, fp []*ParameterDeclaration, isVariadic bool) *FunctionType {
	r := &FunctionType{result: result, fp: fp, minArgs: len(fp), maxArgs: len(fp)}
	if isVariadic {
		r.maxArgs = -1
	}
	if len(fp) == 1 {
		if t := fp[0].typ; t != nil && t.Kind() == Void {
			r.minArgs = 0
			r.maxArgs = 0
		}
	}
	return r
}

func (n *FunctionType) Kind() Kind { return Function }

type PointerType struct {
	elem Type
}

func newPointerType(elem Type) *PointerType { return &PointerType{elem: elem} }

func (n *PointerType) Kind() Kind { return Ptr }

type StructType struct {
	//TODO
}

func (n *StructType) Kind() Kind { return Struct }

type UnionType struct {
	//TODO
}

func (n *UnionType) Kind() Kind { return Union }

type ArrayType struct {
	elem Type
	expr *AssignmentExpression
}

func newArrayType(expr *AssignmentExpression, elem Type) *ArrayType {
	return &ArrayType{elem: elem, expr: expr}
}

func (n *ArrayType) Kind() Kind { return Array }

type EnumType struct {
	//TODO
}

func (n *EnumType) Kind() Kind { return Enum }
