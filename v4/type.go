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
	// Size reports the size of a type in bytes. Incomplete or invalid types may
	// report a negative size.
	Size() int64
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

// Kind implements Type.
func (n *PredefinedType) Kind() Kind { return n.kind }

// Size implements Type.
func (n *PredefinedType) Size() int64 {
	if x, ok := n.ast.ABI.types[n.kind]; ok {
		return x.size
	}

	return -1
}

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

// Kind implements Type.
func (n *FunctionType) Kind() Kind { return Function }

// Size implements Type.
func (n *FunctionType) Size() int64 { return 1 } // gcc compatibility

type PointerType struct {
	ast  *AST
	elem Type
}

func newPointerType(ast *AST, elem Type) *PointerType { return &PointerType{ast: ast, elem: elem} }

// Kind implements Type.
func (n *PointerType) Kind() Kind { return Ptr }

// Size implements Type.
func (n *PointerType) Size() int64 {
	if x, ok := n.ast.ABI.types[Ptr]; ok {
		return x.size
	}

	return -1
}

type StructType struct {
	//TODO
}

// Kind implements Type.
func (n *StructType) Kind() Kind { return Struct }

// Size implements Type.
func (n *StructType) Size() int64 {
	panic(todo(""))
}

type UnionType struct {
	//TODO
}

// Kind implements Type.
func (n *UnionType) Kind() Kind { return Union }

// Size implements Type.
func (n *UnionType) Size() int64 {
	panic(todo(""))
}

type ArrayType struct {
	elem  Type
	expr  *AssignmentExpression
	elems int64
}

func newArrayType(expr *AssignmentExpression, elem Type) (r *ArrayType) {
	r = &ArrayType{elem: elem, expr: expr, elems: -1}
	//TODO handle expr
	return r
}

// Kind implements Type.
func (n *ArrayType) Kind() Kind { return Array }

// Size implements Type.
func (n *ArrayType) Size() int64 {
	if n.elem != nil {
		if a, b := n.elems, n.elem.Size(); a >= 0 && b > 0 {
			return a * b
		}
	}

	return -1
}

type EnumType struct {
	typ Type
}

func newEnumType(typ Type) *EnumType { return &EnumType{typ: typ} }

// Kind implements Type.
func (n *EnumType) Kind() Kind { return Enum }

// Size implements Type.
func (n *EnumType) Size() int64 {
	if n.typ != nil {
		return n.typ.Size()
	}

	return -1
}

// [0] 6.3.1.8 Usual arithmetic conversions
func usualArithmeticConversions(a, b Type) (r, s Type) {
	panic(todo(""))
}
