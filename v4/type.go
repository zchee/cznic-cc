// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

var (
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
	Char       Kind = iota // char
	Double                 // double
	Float                  // float
	Function               // function
	Int                    // int
	Long                   // long
	LongDouble             // long double
	LongLong               // long long
	Ptr                    // pointer
	Schar                  // signed char
	Short                  // short
	Struct                 // struct
	Uchar                  // unsigned char
	Uint                   // unsigned
	Ulong                  // unsigned long
	UlongLong              // unsigned long long
	Union                  // union
	Ushort                 // unsigned short
	Void                   // void
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
