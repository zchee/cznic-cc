// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"fmt"
	"sort"
	"strings"

	"modernc.org/token"
)

var (
	_ Type = (*ArrayType)(nil)
	_ Type = (*EnumType)(nil)
	_ Type = (*FunctionType)(nil)
	_ Type = (*InvalidType)(nil)
	_ Type = (*PointerType)(nil)
	_ Type = (*PredefinedType)(nil)
	_ Type = (*StructType)(nil)
	_ Type = (*UnionType)(nil)
)

var (
	// Invalid is a singleton representing an invalid/undetermined type.  Invalid
	// is comparable.
	Invalid Type = &InvalidType{}

	arithmeticKinds = [maxKind]bool{
		Bool:              true,
		Char:              true,
		ComplexChar:       true,
		ComplexDouble:     true,
		ComplexFloat:      true,
		ComplexInt:        true,
		ComplexLong:       true,
		ComplexLongDouble: true,
		ComplexLongLong:   true,
		ComplexShort:      true,
		ComplexUInt:       true,
		ComplexUShort:     true,
		Decimal128:        true,
		Decimal32:         true,
		Decimal64:         true,
		Double:            true,
		Enum:              true,
		Float128:          true,
		Float128x:         true,
		Float16:           true,
		Float32:           true,
		Float32x:          true,
		Float64:           true,
		Float64x:          true,
		Float:             true,
		Int128:            true,
		Int:               true,
		Long:              true,
		LongDouble:        true,
		LongLong:          true,
		SChar:             true,
		Short:             true,
		UChar:             true,
		UInt128:           true,
		UInt:              true,
		ULong:             true,
		ULongLong:         true,
		UShort:            true,
	}

	floatinPointKinds = [maxKind]bool{
		Double:     true,
		Float128:   true,
		Float128x:  true,
		Float16:    true,
		Float32:    true,
		Float32x:   true,
		Float64:    true,
		Float64x:   true,
		Float:      true,
		LongDouble: true,
	}

	integerKinds = [maxKind]bool{
		Bool:      true,
		Char:      true,
		Enum:      true,
		Int128:    true,
		Int:       true,
		Long:      true,
		LongLong:  true,
		SChar:     true,
		Short:     true,
		UChar:     true,
		UInt128:   true,
		UInt:      true,
		ULong:     true,
		ULongLong: true,
		UShort:    true,
	}

	complexKinds = [maxKind]bool{
		ComplexChar:       true,
		ComplexDouble:     true,
		ComplexFloat:      true,
		ComplexInt:        true,
		ComplexLong:       true,
		ComplexLongDouble: true,
		ComplexLongLong:   true,
		ComplexShort:      true,
		ComplexUInt:       true,
		ComplexUShort:     true,
	}

	correspondingRealKinds = [maxKind]Kind{
		ComplexChar:       Char,
		ComplexDouble:     Double,
		ComplexFloat:      Float,
		ComplexInt:        Int,
		ComplexLong:       Long,
		ComplexLongDouble: LongDouble,
		ComplexLongLong:   LongLong,
		ComplexShort:      Short,
		ComplexUInt:       UInt,
		ComplexUShort:     UShort,
	}

	// Keep Bool first and sorted by rank.
	intConvRanks = [maxKind]int{
		Bool:      1,
		Char:      2,
		SChar:     2,
		UChar:     2,
		Short:     3,
		UShort:    3,
		Int:       4,
		UInt:      4,
		Long:      5,
		ULong:     5,
		LongLong:  6,
		ULongLong: 6,
		Int128:    7,
		UInt128:   7,
	}

	realKinds = [maxKind]bool{
		Bool:       true,
		Char:       true,
		Decimal32:  true,
		Decimal64:  true,
		Double:     true,
		Enum:       true,
		Float:      true,
		Int128:     true,
		Int:        true,
		Long:       true,
		LongDouble: true,
		LongLong:   true,
		SChar:      true,
		Short:      true,
		UChar:      true,
		UInt:       true,
		ULong:      true,
		ULongLong:  true,
		UShort:     true,
		UInt128:    true,
	}
)

type Kind int

const (
	InvalidKind Kind = iota

	Array             // array
	Bool              // _Bool
	Char              // char
	ComplexChar       // _Complex char
	ComplexDouble     // _Complex double
	ComplexFloat      // _Complex float
	ComplexInt        // _Complex int
	ComplexLong       // _Complex long
	ComplexLongDouble // _Complex long double
	ComplexLongLong   // _Complex long long
	ComplexShort      // _Complex short
	ComplexUInt       // _Complex unsigned
	ComplexUShort     // _Complex unsigned short
	Decimal128        // _Decimal128
	Decimal32         // _Decimal32
	Decimal64         // _Decimal64
	Double            // double
	Enum              // enum
	Float             // float
	Float128          // _Float128
	Float128x         // _Float128x
	Float16           // _Float16
	Float32           // _Float32
	Float32x          // _Float32x
	Float64           // _Float64
	Float64x          // _Float64x
	Function          // function
	Int               // int
	Int128            // __int128
	Long              // long
	LongDouble        // long double
	LongLong          // long long
	Ptr               // pointer
	SChar             // signed char
	Short             // short
	Struct            // struct
	UChar             // unsigned char
	UInt              // unsigned
	UInt128           // unsigned __int128
	ULong             // unsigned long
	ULongLong         // unsigned long long
	UShort            // unsigned short
	Union             // union
	Void              // void

	maxKind
)

// Type is the representation of a C type.
//
// The dynamic type of a Type is one of
//
//  *ArrayType
//  *EnumType
//  *FunctionType
//  *InvalidType
//  *PointerType
//  *PredefinedType
//  *StructType
//  *UnionType
type Type interface {
	// Align reports the minimum alignment required by a type.
	Align() int

	Attributes() *Attributes

	// Decay returns a pointer to array element for array types, a pointer to a
	// function for function types and itself for all other type kinds.
	Decay() Type

	// FieldAlign reports the minimum alignment required by a type when it's used
	// in a struct/union.
	FieldAlign() int

	// IsIncomplete reports whether the size of a type is not determined.
	IsIncomplete() bool

	// Kind reports the kind of a type.
	Kind() Kind

	// Name returns the type name of a type or an empty string. C types are
	// associated with names using typedef.
	Name() string

	// Size reports the size of a type in bytes. Incomplete or invalid types may
	// report a negative size.
	Size() int64

	// String produces a human readable representation of a type or an
	// approximation of the same. The particular form is not specified and
	// may change. Namely, the returned value is not suitable for directly
	// determining type identity.
	String() string

	// Undecay reverses Decay() if the type is a pointer and was produced by
	// Decay() and the result was different than the Decay() receiver. Otherwise
	// Undecay() returns its receiver.
	Undecay() Type

	// VectorSize reports N from __attribute__((vector_size(N))). Valid if
	// > 0.
	VectorSize() int64

	isCompatible(Type) bool
	setAttr(*Attributes) Type
	setName(nm string) Type
	str(b *strings.Builder, useTag bool) *strings.Builder
}

type namer string

// Name implements Type
func (n namer) Name() string { return string(n) }

type InvalidType struct{}

// setAttr implements Type.
func (n InvalidType) setAttr(*Attributes) Type { return Invalid }

// VectorSize implements Type.
func (n InvalidType) VectorSize() int64 { return -1 }

// Attributes implements Type.
func (n *InvalidType) Attributes() *Attributes { return nil }

// Align implements Type.
func (n *InvalidType) Align() int { return 1 }

// Decay implements Type.
func (n *InvalidType) Decay() Type { return n }

// Undecay implements Type.
func (n *InvalidType) Undecay() Type { return n }

// FieldAlign implements Type.
func (n *InvalidType) FieldAlign() int { return 1 }

// Name implements Type.
func (n *InvalidType) Name() string { return "" }

// setName implements Type.
func (n *InvalidType) setName(nm string) Type { return n }

func (n *InvalidType) isCompatible(Type) bool { return false }

// String implements Type.
func (n *InvalidType) String() string { return "<invalid type>" }

func (n *InvalidType) str(b *strings.Builder, useTag bool) *strings.Builder {
	b.WriteString("<invalid type>")
	return b
}

// IsIncomplete implements Type.
func (n *InvalidType) IsIncomplete() bool { return true }

// Kind implements Type.
func (n *InvalidType) Kind() Kind { return InvalidKind }

// Size implements Type.
func (n *InvalidType) Size() int64 { return -1 }

type PredefinedType struct {
	attributer
	c    *ctx
	kind Kind
	namer
	vector *ArrayType
}

func (c *ctx) newPredefinedType(kind Kind) *PredefinedType {
	return &PredefinedType{c: c, kind: kind}
}

// VectorSize implements Type.
func (n *PredefinedType) VectorSize() int64 {
	if a := n.Attributes(); a != nil {
		return a.VectorSize()
	}

	return -1
}

// setAttr implements Type.
func (n *PredefinedType) setAttr(a *Attributes) Type {
	var vec *ArrayType
	if sz := a.VectorSize(); sz > 0 {
		vec = n.c.newArrayType(n, sz/n.Size(), nil)
	}
	m := *n
	m.vector = vec
	m.attributer.p = a
	return &m
}

func (n *PredefinedType) isCompatible(t Type) bool {
	switch x := t.(type) {
	case *PredefinedType:
		return n == x || n.Kind() == x.Kind()
	default:
		return false
	}
}

// setName implements Type.
func (n *PredefinedType) setName(nm string) Type {
	r := *n
	r.namer = namer(nm)
	return &r
}

// Align implements Type.
func (n *PredefinedType) Align() int {
	if n == nil {
		return 1
	}

	if n.attributer.p != nil {
		if v := n.attributer.p.Aligned(); v > 0 {
			return int(v)
		}
	}

	if x, ok := n.c.ast.ABI.types[n.kind]; ok {
		return x.align
	}

	return 1
}

// Decay implements Type.
func (n *PredefinedType) Decay() Type { return n }

// Undecay implements Type.
func (n *PredefinedType) Undecay() Type { return n }

// FieldAlign implements Type.
func (n *PredefinedType) FieldAlign() int {
	if n == nil {
		return 1
	}

	if x, ok := n.c.ast.ABI.types[n.kind]; ok {
		return x.fieldAlign
	}

	return 1
}

// String implements Type.
func (n *PredefinedType) String() string { return n.str(&strings.Builder{}, false).String() }

func (n *PredefinedType) str(b *strings.Builder, useTag bool) *strings.Builder {
	if s := n.Name(); s != "" {
		b.WriteString(s)
		return b
	}

	b.WriteString(n.kind.String())
	return b
}

// IsIncomplete implements Type.
func (n *PredefinedType) IsIncomplete() bool {
	if n == nil {
		return true
	}

	return n.Size() < 0
}

// Kind implements Type.
func (n *PredefinedType) Kind() Kind { return n.kind }

// Size implements Type.
func (n *PredefinedType) Size() int64 {
	if n == nil {
		return -1
	}

	if isIntegerType(n) || isFloatingPointType(n) {
		if v := n.VectorSize(); v > 0 {
			return v
		}
	}

	if x, ok := n.c.ast.ABI.types[n.kind]; ok {
		return x.size
	}

	return -1
}

// Parameter represents a function parameter.
type Parameter struct {
	Declarator         *Declarator         // Can be nil.
	AbstractDeclarator *AbstractDeclarator // Can be nil
	name               Token
	typer
	resolver
	visible
}

// Position implements Node.
func (n *Parameter) Position() (r token.Position) {
	if n.Declarator != nil {
		return n.Declarator.Position()
	}

	if n.AbstractDeclarator != nil {
		return n.AbstractDeclarator.Position()
	}

	return n.name.Position()
}

// Name returns the name of n. The result can be a zero value, like in
// `void f(int) { ... }`.
func (n *Parameter) Name() string {
	if n.Declarator != nil {
		return n.Declarator.DirectDeclarator.name().Token.SrcStr()
	}

	return n.name.SrcStr()
}

type FunctionType struct {
	attributer
	c  *ctx
	fp []*Parameter
	namer
	result typer
	vectorSizer

	minArgs int
	maxArgs int // -1: unlimited
}

func (c *ctx) newFunctionType(result Type, fp []*ParameterDeclaration, isVariadic bool) (r *FunctionType) {
	r = &FunctionType{c: c, result: newTyper(result), minArgs: len(fp), maxArgs: len(fp)}
	for _, n := range fp {
		p := &Parameter{}
		p.typ = n.Type()
		switch n.Case {
		case ParameterDeclarationDecl: // DeclarationSpecifiers Declarator
			p.Declarator = n.Declarator
		case ParameterDeclarationAbstract: // DeclarationSpecifiers AbstractDeclarator
			p.AbstractDeclarator = n.AbstractDeclarator
		default:
			c.errors.add(errorf("internal error: %v", n.Case))
		}
		r.fp = append(r.fp, p)
	}
	if isVariadic {
		r.maxArgs = -1
	}
	if len(fp) == 1 {
		if t := fp[0].Type(); t != nil && t.Kind() == Void {
			r.minArgs = 0
			r.maxArgs = 0
		}
	}
	return r
}

// setAttr implements Type.
func (n *FunctionType) setAttr(a *Attributes) Type {
	m := *n
	m.attributer.p = a
	return &m
}

func (n *FunctionType) isCompatible(t Type) bool {
	switch x := t.(type) {
	case *FunctionType:
		if n == x {
			return true
		}

		if len(n.fp) != len(x.fp) || n.minArgs != x.minArgs || n.maxArgs != x.maxArgs || !n.Result().isCompatible(x.Result()) {
			return false
		}

		for i, v := range n.fp {
			if !v.Type().isCompatible(x.fp[i].Type()) {
				return false
			}
		}

		return true
	default:
		return false
	}
}

// setName implements Type.
func (n *FunctionType) setName(nm string) Type {
	r := *n
	r.namer = namer(nm)
	return &r
}

// Result reports the result type of n.
func (n *FunctionType) Result() Type { return n.result.Type() }

// Parameters returns function type parameters.
func (n *FunctionType) Parameters() []*Parameter { return n.fp }

// Align implements Type.
func (n *FunctionType) Align() int {
	if n.attributer.p != nil {
		if v := n.attributer.p.Aligned(); v > 0 {
			return int(v)
		}
	}

	return 1
}

// Decay implements Type.
func (n *FunctionType) Decay() Type { return n.c.newPointerType2(n, n) }

// Undecay implements Type.
func (n *FunctionType) Undecay() Type { return n }

// FieldAlign implements Type.
func (n *FunctionType) FieldAlign() int { return 1 }

// IsIncomplete implements Type.
func (n *FunctionType) IsIncomplete() bool {
	if n == nil {
		return true
	}

	return false
}

// Kind implements Type.
func (n *FunctionType) Kind() Kind { return Function }

// Size implements Type.
func (n *FunctionType) Size() int64 {
	if n == nil {
		return -1
	}

	return 1
} // gcc compatibility

// String implements Type.
func (n *FunctionType) String() string { return n.str(&strings.Builder{}, false).String() }

func (n *FunctionType) str(b *strings.Builder, useTag bool) *strings.Builder {
	if s := n.Name(); s != "" {
		b.WriteString(s)
		return b
	}

	b.WriteString("function(")
	switch {
	case n.maxArgs == 0:
		b.WriteString("void")
	default:
		for i, v := range n.fp {
			v.Type().str(b, true)
			if i != len(n.fp)-1 {
				b.WriteString(", ")
			}
		}
	}
	b.WriteByte(')')
	if n.Result().Kind() != Void {
		b.WriteString(" returning ")
		n.Result().str(b, true)
	}
	return b
}

type PointerType struct {
	attributer
	c    *ctx
	elem typer
	namer
	undecay Type
	vectorSizer
}

func (c *ctx) newPointerType(elem Type) (r *PointerType) {
	r = &PointerType{c: c, elem: newTyper(elem)}
	r.undecay = r
	return r
}

// setAttr implements Type.
func (n *PointerType) setAttr(a *Attributes) Type {
	m := *n
	m.attributer.p = a
	return &m
}

func (n *PointerType) isCompatible(t Type) bool {
	switch x := t.(type) {
	case *PointerType:
		return n == x || n.Elem().isCompatible(x.Elem())
	default:
		return false
	}
}

// setName implements Type.
func (n *PointerType) setName(nm string) Type {
	r := *n
	r.namer = namer(nm)
	return &r
}

func (c *ctx) newPointerType2(elem, undecay Type) *PointerType {
	return &PointerType{c: c, elem: newTyper(elem), undecay: undecay}
}

// Elem returns the type n points to.
func (n *PointerType) Elem() Type { return n.elem.Type() }

// Align implements Type.
func (n *PointerType) Align() int {
	if n == nil {
		return 1
	}

	if n.attributer.p != nil {
		if v := n.attributer.p.Aligned(); v > 0 {
			return int(v)
		}
	}

	if x, ok := n.c.ast.ABI.types[Ptr]; ok {
		return x.align
	}

	return 1
}

// Decay implements Type.
func (n *PointerType) Decay() Type { return n }

// Undecay implements Type.
func (n *PointerType) Undecay() Type { return n.undecay }

// FieldAlign implements Type.
func (n *PointerType) FieldAlign() int {
	if n == nil {
		return 1
	}

	if x, ok := n.c.ast.ABI.types[Ptr]; ok {
		return x.fieldAlign
	}

	return 1
}

// IsIncomplete implements Type.
func (n *PointerType) IsIncomplete() bool {
	if n == nil {
		return true
	}

	return false
}

// Kind implements Type.
func (n *PointerType) Kind() Kind { return Ptr }

// Size implements Type.
func (n *PointerType) Size() int64 {
	if n == nil {
		return -1
	}

	if x, ok := n.c.ast.ABI.types[Ptr]; ok {
		return x.size
	}

	return -1
}

// String implements Type.
func (n *PointerType) String() string { return n.str(&strings.Builder{}, false).String() }

func (n *PointerType) str(b *strings.Builder, useTag bool) *strings.Builder {
	if s := n.Name(); s != "" {
		b.WriteString(s)
		return b
	}

	b.WriteString("pointer to ")
	n.elem.Type().str(b, true)
	return b
}

type Field struct {
	accessBytes int64       // accessBytes < typ.Size() -> bit field.
	declarator  *Declarator // Can be nil.
	mask        uint64      // Non zero only for bit fields.
	offsetBytes int64
	parent      *Field
	typ         typer
	valueBits   int64

	depth int
	// Additional bit offset to offset bytes. Non zero only for bit fields but can
	// be zero even for a bit field, for example, the first bit field after a non
	// bit field will have offsetBits zero.
	offsetBits int
	ordinal    int // index into .fields in structType

	isBitField bool
}

func (n *Field) path() (r []int) {
	if n.parent != nil {
		r = n.parent.path()
	}
	return append(r, n.ordinal)
}

// Type reports the type of f.
func (n *Field) Type() Type { return n.typ.Type() }

// Name reports the name of f, if any.
func (n *Field) Name() string {
	if d := n.declarator; d != nil {
		return d.Name()
	}

	return ""
}

func (n *Field) Offset() int64 { return n.offsetBytes }

type structType struct {
	fields []*Field
	m      map[string]*Field
	size   int64
	tag    string

	align   int
	isUnion bool
}

func (n *structType) isIncomplete() bool {
	for i, v := range n.fields {
		if v.Type().IsIncomplete() {
			if x, ok := v.Type().(*ArrayType); ok && x.IsVLA() {
				continue
			}

			if i == len(n.fields)-1 && v.Type().Kind() == Array { // Flexible array member.
				continue
			}

			return true
		}
	}
	return false
}

func (n *structType) isCompatible(m *structType) bool {
	if n == m {
		return true
	}

	if n.size != m.size || n.tag != m.tag || len(n.fields) != len(m.fields) {
		return false
	}

	for i, v := range n.fields {
		if w := m.fields[i]; v.Name() != w.Name() || !v.Type().isCompatible(w.Type()) {
			return false
		}
	}

	return true
}

func (n *structType) fieldByIndex(i int) *Field {
	if i >= 0 && i < len(n.fields) {
		return n.fields[i]
	}

	return nil
}

func (n *structType) fieldByName(nm string) *Field {
	if f := n.m[nm]; f != nil {
		return f
	}

	if n.m == nil {
		m := map[string][]*Field{}
		n.collectFields(m, nil, 0, 0)
		n.m = map[string]*Field{}
		for k, v := range m {
			sort.Slice(v, func(i, j int) bool { return v[i].depth < v[j].depth })
			switch {
			case len(v) != 1:
				if v[0].depth < v[1].depth {
					n.m[k] = v[0]
					break
				}

				if n.isUnion && v[0].depth == v[1].depth {
					p0 := v[0].path()
					p1 := v[1].path()
					if p0[0] == 0 && p1[0] != 0 {
						n.m[k] = v[0]
					}
				}
			default:
				n.m[k] = v[0]
			}
		}
	}
	return n.m[nm]
}

func (n *structType) collectFields(m map[string][]*Field, parent *Field, depth int, off int64) {
	for _, f := range n.fields {
		if nm := f.Name(); nm != "" {
			switch {
			case depth != 0:
				f2 := *f
				f2.offsetBytes += off
				f2.depth = depth
				f2.parent = parent
				m[nm] = append(m[nm], &f2)
			default:
				m[nm] = append(m[nm], f)
			}
		}
		switch f.Type().Kind() {
		case Struct:
			f.Type().(*StructType).collectFields(m, f, depth+1, f.offsetBytes)
		case Union:
			f.Type().(*UnionType).collectFields(m, f, depth+1, f.offsetBytes)
		}
	}
}

type StructType struct {
	attributer
	c       *ctx
	forward *StructOrUnionSpecifier
	namer
	structType
	vectorSizer
}

func (c *ctx) newStructType(tag string, fields []*Field, size int64, align int) (r *StructType) {
	r = &StructType{c: c, structType: structType{tag: tag, fields: fields, size: size, align: align}}
	return r
}

// setAttr implements Type.
func (n *StructType) setAttr(a *Attributes) Type {
	m := *n
	m.attributer.p = a
	return &m
}

func (n *StructType) isCompatible(t Type) bool {
	if n.forward != nil {
		return n.forward.Type().isCompatible(t)
	}

	switch x := t.(type) {
	case *StructType:
		if x.forward != nil {
			return n.isCompatible(x.forward.Type())
		}

		return n == x || n.structType.isCompatible(&x.structType)
	default:
		return false
	}
}

// setName implements Type.
func (n *StructType) setName(nm string) Type {
	r := *n
	r.namer = namer(nm)
	return &r
}

// FieldByIndex returns a member field by index, if any.
func (n *StructType) FieldByIndex(i int) *Field {
	if n == nil {
		return nil
	}

	if n.forward != nil {
		if x, ok := n.forward.typ.(*StructType); ok {
			return x.FieldByIndex(i)
		}

		return nil
	}

	return n.fieldByIndex(i)
}

// NamedFieldByIndex returns the first named member field at or after index, if any.
func (n *StructType) NamedFieldByIndex(i int) (r *Field) {
	for ; ; i++ {
		r = n.FieldByIndex(i)
		if r == nil {
			return nil
		}

		if r.Name() != "" {
			return r
		}
	}
}

// FieldByName returns the shallowest member field by name, if any.
func (n *StructType) FieldByName(nm string) *Field {
	if n == nil {
		return nil
	}

	if n.forward != nil {
		if x, ok := n.forward.typ.(*StructType); ok {
			return x.FieldByName(nm)
		}

		return nil
	}

	return n.fieldByName(nm)
}

// Align implements Type.
func (n *StructType) Align() int {
	if n == nil {
		return 1
	}

	if n.forward != nil {
		return n.forward.Type().Align()
	}

	if n.attributer.p != nil {
		if v := n.attributer.p.Aligned(); v > 0 {
			return int(v)
		}
	}

	if n.IsIncomplete() {
		return n.c.intT.Align()
	}

	return n.align
}

// Decay implements Type.
func (n *StructType) Decay() Type { return n }

// Undecay implements Type.
func (n *StructType) Undecay() Type { return n }

// FieldAlign implements Type.
func (n *StructType) FieldAlign() int {
	if n == nil {
		return 1
	}

	if n.forward != nil {
		return n.forward.Type().FieldAlign()
	}

	return n.align
}

// IsIncomplete implements Type.
func (n *StructType) IsIncomplete() bool {
	if n == nil {
		return true
	}

	if n.forward != nil {
		return n.forward.Type().IsIncomplete()
	}

	return n.isIncomplete()
}

// Kind implements Type.
func (n *StructType) Kind() Kind { return Struct }

// Size implements Type.
func (n *StructType) Size() int64 {
	if n == nil {
		return -1
	}

	if n.forward != nil {
		return n.forward.Type().Size()
	}

	return n.size
}

// String implements Type.
func (n *StructType) String() string { return n.str(&strings.Builder{}, false).String() }

func (n *StructType) str(b *strings.Builder, useTag bool) *strings.Builder {
	if n.forward != nil {
		b.WriteString(n.forward.Type().String())
		return b
	}

	if s := n.Name(); s != "" {
		b.WriteString(s)
		return b
	}

	b.WriteString("struct")
	if n.tag != "" {
		b.WriteByte(' ')
		b.WriteString(n.tag)
		if useTag {
			return b
		}
	}

	b.WriteString(" {")
	for i, v := range n.fields {
		if v.declarator != nil {
			b.WriteString(v.declarator.Name())
			b.WriteByte(' ')
		}
		v.Type().str(b, true)
		if i != len(n.fields)-1 {
			b.WriteString("; ")
		}
	}
	b.WriteByte('}')
	return b
}

type UnionType struct {
	attributer
	c       *ctx
	forward *StructOrUnionSpecifier
	namer
	structType
	vectorSizer
}

func (c *ctx) newUnionType(tag string, fields []*Field, size int64, align int) *UnionType {
	return &UnionType{c: c, structType: structType{tag: tag, fields: fields, size: size, align: align, isUnion: true}}
}

// setAttr implements Type.
func (n *UnionType) setAttr(a *Attributes) Type {
	m := *n
	m.attributer.p = a
	return &m
}

func (n *UnionType) isCompatible(t Type) bool {
	if n.forward != nil {
		return n.forward.Type().isCompatible(t)
	}

	switch x := t.(type) {
	case *UnionType:
		if x.forward != nil {
			return n.isCompatible(x.forward.Type())
		}

		return n == x || n.structType.isCompatible(&x.structType)
	default:
		return false
	}
}

// setName implements Type.
func (n *UnionType) setName(nm string) Type {
	r := *n
	r.namer = namer(nm)
	return &r
}

// FieldByIndex returns a member field by index, if any.
func (n *UnionType) FieldByIndex(i int) *Field {
	if n == nil {
		return nil
	}

	if n.forward != nil {
		if x, ok := n.forward.typ.(*UnionType); ok {
			return x.FieldByIndex(i)
		}

		return nil
	}

	return n.fieldByIndex(i)
}

// NamedFieldByIndex returns the first named member field at or after index, if any.
func (n *UnionType) NamedFieldByIndex(i int) (r *Field) {
	for ; ; i++ {
		r = n.FieldByIndex(i)
		if r == nil {
			return nil
		}

		if r.Name() != "" {
			return r
		}
	}
}

// FieldByName returns member field nm of n or nil if n does not have such member.
func (n *UnionType) FieldByName(nm string) *Field {
	if n == nil {
		return nil
	}

	if n.forward != nil {
		if x, ok := n.forward.typ.(*UnionType); ok {
			return x.FieldByName(nm)
		}

		return nil
	}

	return n.fieldByName(nm)
}

// Align implements Type.
func (n *UnionType) Align() int {
	if n == nil {
		return 1
	}

	if n.forward != nil {
		return n.forward.Type().Align()
	}

	if n.attributer.p != nil {
		if v := n.attributer.p.Aligned(); v > 0 {
			return int(v)
		}
	}

	if n.IsIncomplete() {
		return n.c.intT.Align()
	}

	return n.align
}

// Decay implements Type.
func (n *UnionType) Decay() Type { return n }

// Undecay implements Type.
func (n *UnionType) Undecay() Type { return n }

// FieldAlign implements Type.
func (n *UnionType) FieldAlign() int {
	if n == nil {
		return 1
	}

	if n.forward != nil {
		return n.forward.Type().FieldAlign()
	}

	return n.align
}

// IsIncomplete implements Type.
func (n *UnionType) IsIncomplete() bool {
	if n == nil {
		return true
	}

	if n.forward != nil {
		return n.forward.Type().IsIncomplete()
	}

	return n.isIncomplete()
}

// Kind implements Type.
func (n *UnionType) Kind() Kind { return Union }

// Size implements Type.
func (n *UnionType) Size() int64 {
	if n == nil {
		return -1
	}

	if n.forward != nil {
		return n.forward.Type().Size()
	}

	return n.size
}

// String implements Type.
func (n *UnionType) String() string { return n.str(&strings.Builder{}, false).String() }

func (n *UnionType) str(b *strings.Builder, useTag bool) *strings.Builder {
	if n.forward != nil {
		b.WriteString(n.forward.Type().String())
		return b
	}

	if s := n.Name(); s != "" {
		b.WriteString(s)
		return b
	}

	b.WriteString("union")
	if n.tag != "" {
		b.WriteByte(' ')
		b.WriteString(n.tag)
		if useTag {
			return b
		}
	}

	b.WriteString(" {")
	for i, v := range n.fields {
		if v.declarator != nil {
			b.WriteString(v.declarator.Name())
			b.WriteByte(' ')
		}
		v.Type().str(b, true)
		if i != len(n.fields)-1 {
			b.WriteString("; ")
		}
	}
	b.WriteByte('}')
	return b
}

type vectorSizer struct{}

// VectorSize implements Type.
func (vectorSizer) VectorSize() int64 { return -1 }

type ArrayType struct {
	attributer
	c     *ctx
	elem  typer
	elems int64
	expr  ExpressionNode
	namer
	vectorSizer
}

func (c *ctx) newArrayType(elem Type, elems int64, expr ExpressionNode) (r *ArrayType) {
	r = &ArrayType{c: c, elem: newTyper(elem), elems: elems, expr: expr}
	return r
}

// setAttr implements Type.
func (n *ArrayType) setAttr(a *Attributes) Type {
	m := *n
	m.attributer.p = a
	return &m
}

func (n *ArrayType) isCompatible(t Type) bool {
	switch x := t.(type) {
	case *ArrayType:
		return (n.Len() == x.Len() || n.Len() < 0 && x.Len() < 0) && n.Elem().isCompatible(x.Elem())
	default:
		return false
	}
}

// setName implements Type.
func (n *ArrayType) setName(nm string) Type {
	r := *n
	r.namer = namer(nm)
	return &r
}

func (n *ArrayType) IsVLA() bool {
	for {
		if n.elems < 0 && n.expr != nil && n.expr.Value() == Unknown {
			return true
		}

		x, ok := n.Elem().(*ArrayType)
		if !ok {
			break
		}

		n = x
	}
	return false
}

// Decay implements Type.
func (n *ArrayType) Decay() Type { return n.c.newPointerType2(n.Elem(), n) }

// Undecay implements Type.
func (n *ArrayType) Undecay() Type { return n }

// Elem reports the element type of n.
func (n *ArrayType) Elem() Type { return n.elem.Type() }

// Align implements Type.
func (n *ArrayType) Align() int {
	if n == nil {
		return 1
	}

	if n.attributer.p != nil {
		if v := n.attributer.p.Aligned(); v > 0 {
			return int(v)
		}
	}

	return n.elem.Type().Align()
}

// FieldAlign implements Type.
func (n *ArrayType) FieldAlign() int {
	if n == nil {
		return 1
	}

	return n.elem.Type().FieldAlign()
}

// IsIncomplete implements Type.
func (n *ArrayType) IsIncomplete() bool {
	if n == nil {
		return true
	}

	return n.Elem().IsIncomplete() || n.elems < 0
}

// Kind implements Type.
func (n *ArrayType) Kind() Kind { return Array }

// Len reports the number of elements in n or a negative value if n is incomplete.
func (n *ArrayType) Len() int64 { return n.elems }

// Size implements Type.
func (n *ArrayType) Size() int64 {
	if n == nil {
		return -1
	}

	if n.Elem().Kind() != InvalidKind {
		if a, b := n.elems, n.Elem().Size(); a >= 0 && b > 0 {
			return a * b
		}
	}

	return -1
}

// String implements Type.
func (n *ArrayType) String() string { return n.str(&strings.Builder{}, false).String() }

func (n *ArrayType) str(b *strings.Builder, useTag bool) *strings.Builder {
	if s := n.Name(); s != "" {
		b.WriteString(s)
		return b
	}

	b.WriteString("array of ")
	if !n.IsIncomplete() {
		fmt.Fprintf(b, "%d ", n.elems)
	}
	n.Elem().str(b, true)
	return b
}

type EnumType struct {
	attributer
	enums   []*Enumerator
	forward *EnumSpecifier
	namer
	tag string
	typ typer
	vectorSizer
}

func (c *ctx) newEnumType(tag string, typ Type, enums []*Enumerator) *EnumType {
	return &EnumType{tag: tag, typ: newTyper(typ), enums: enums}
}

// setAttr implements Type.
func (n *EnumType) setAttr(a *Attributes) Type {
	m := *n
	m.attributer.p = a
	return &m
}

func (n *EnumType) isCompatible(t Type) bool {
	if n.forward != nil {
		return n.forward.Type().isCompatible(t)
	}

	switch x := t.(type) {
	case *EnumType:
		if x.forward != nil {
			return n.isCompatible(x.forward.Type())
		}

		if n.tag != x.tag {
			return false
		}

		return n.isCompatible(x.typ.Type()) //TODO members and values must be the same
	default:
		return false
	}
}

// setName implements Type.
func (n *EnumType) setName(nm string) Type {
	r := *n
	r.namer = namer(nm)
	return &r
}

// Align implements Type.
func (n *EnumType) Align() int {
	if n == nil {
		return 1
	}

	if n.forward != nil {
		return n.forward.Type().Align()
	}

	if n.attributer.p != nil {
		if v := n.attributer.p.Aligned(); v > 0 {
			return int(v)
		}
	}

	return n.typ.Type().Align()
}

// Decay implements Type.
func (n *EnumType) Decay() Type { return n }

// Undecay implements Type.
func (n *EnumType) Undecay() Type { return n }

// FieldAlign implements Type.
func (n *EnumType) FieldAlign() int {
	if n == nil {
		return 1
	}

	if n.forward != nil {
		return n.forward.Type().FieldAlign()
	}

	return n.typ.Type().FieldAlign()
}

// IsIncomplete implements Type.
func (n *EnumType) IsIncomplete() bool {
	if n == nil {
		return true
	}

	if n.forward != nil {
		return n.forward.Type().IsIncomplete()
	}

	return n.typ.Type().IsIncomplete()
}

// Kind implements Type.
func (n *EnumType) Kind() Kind { return Enum }

// Size implements Type.
func (n *EnumType) Size() int64 {
	if n == nil {
		return -1
	}

	if n.forward != nil {
		return n.forward.Type().Size()
	}

	return n.typ.Type().Size()
}

// String implements Type.
func (n *EnumType) String() string { return n.str(&strings.Builder{}, false).String() }

func (n *EnumType) str(b *strings.Builder, useTag bool) *strings.Builder {
	if s := n.Name(); s != "" {
		b.WriteString(s)
		return b
	}

	b.WriteString("enum ")
	if n.tag != "" {
		b.WriteString(n.tag)
	}
	if n.forward != nil {
		return b
	}

	b.WriteString(" { ... }") //TODO
	return b
}

func isObjectType(t Type) bool { return t.Kind() != Function && !t.IsIncomplete() }

func isLvalue(t Type) bool { return t.Kind() != Function && t.Kind() != Void }

func isModifiableLvalue(t Type) bool {
	return t.Kind() != Function && t.Kind() != Void && t.Kind() != Array && !t.IsIncomplete()
}

func isPointerType(t Type) bool { return t.Kind() == Ptr }

func isVectorType(t Type) bool { a := t.Attributes(); return a != nil && a.vectorSize > 0 }

func isIntegerType(t Type) bool { return integerKinds[t.Kind()] }

func isFloatingPointType(t Type) bool { return floatinPointKinds[t.Kind()] }

func isComplexType(t Type) bool { return complexKinds[t.Kind()] }

func isScalarType(t Type) bool { return isArithmeticType(t) || t.Kind() == Ptr }

func isArithmeticType(t Type) bool { return arithmeticKinds[t.Kind()] }

func isRealType(t Type) bool { return realKinds[t.Kind()] }

// [0] 6.3.1.8 Usual arithmetic conversions
func usualArithmeticConversions(a, b Type) (r Type) {
	if a.Kind() == Enum {
		a = a.(*EnumType).typ.Type()
	}
	if b.Kind() == Enum {
		b = b.(*EnumType).typ.Type()
	}
	ak := a.Kind()
	bk := b.Kind()
	if ak == InvalidKind || bk == InvalidKind {
		return Invalid
	}

	// First, if the corresponding real type of either operand is long double, the
	// other operand is converted, without change of type domain, to a type whose
	// corresponding real type is long double.
	if ak == ComplexLongDouble {
		return a
	}

	if bk == ComplexLongDouble {
		return b
	}

	if ak == LongDouble {
		return a
	}

	if bk == LongDouble {
		return b
	}

	if ak == Float128 {
		return a
	}

	if bk == Float128 {
		return b
	}

	if ak == Decimal128 {
		return a
	}

	if bk == Decimal128 {
		return b
	}

	if ak == Decimal64 {
		return a
	}

	if bk == Decimal64 {
		return b
	}

	// Otherwise, if the corresponding real type of either operand is double, the
	// other operand is converted, without change of type domain, to a type whose
	// corresponding real type is double.

	if ak == ComplexDouble {
		return a
	}

	if bk == ComplexDouble {
		return b
	}

	if ak == Double {
		return a
	}

	if bk == Double {
		return b
	}

	// Otherwise, if the corresponding real type of either operand is float, the
	// other operand is converted, without change of type domain, to a type whose
	// corresponding real type is float.

	if ak == ComplexFloat {
		return a
	}

	if bk == ComplexFloat {
		return b
	}

	if ak == Float {
		return a
	}

	if bk == Float {
		return b
	}

	// ---- gcc complex integers
	if isComplexType(a) || isComplexType(b) {
		if a.Kind() == b.Kind() {
			return a
		}

		if a.Size() > b.Size() {
			return a
		}

		if b.Size() > a.Size() {
			return b
		}

		panic(todo("", a, b))
	}

	// Otherwise, the integer promotions are performed on both operands.

	if !isIntegerType(a) || !isIntegerType(b) {
		panic(todo("internal error: %s and %s", a, b))
	}

	ast := a.(*PredefinedType).c.ast
	abi := ast.ABI

	ak = integerPromotionKind(ak)
	bk = integerPromotionKind(bk)
	a = ast.kinds[ak]
	b = ast.kinds[bk]

	// Then the following rules are applied to the promoted operands:

	// If both operands have the same type, then no further conversion is needed.
	if ak == bk {
		return a
	}

	// Otherwise, if both operands have signed integer types or both have unsigned
	// integer types, the operand with the type of lesser integer conversion rank
	// is converted to the type of the operand with greater rank.
	if abi.isSignedInteger(ak) == abi.isSignedInteger(bk) {
		if intConvRanks[ak] > intConvRanks[bk] {
			return a
		}

		return b
	}

	// Otherwise, if the operand that has unsigned integer type has rank greater or
	// equal to the rank of the type of the other operand, then the operand with
	// signed integer type is converted to the type of the operand with unsigned
	// integer type.
	switch {
	case abi.isSignedInteger(ak):
		// b is unsigned
		if intConvRanks[bk] > intConvRanks[ak] {
			return b
		}
	default:
		// a is unsigned
		if intConvRanks[ak] > intConvRanks[bk] {
			return a
		}
	}

	// Otherwise, if the type of the operand with signed integer type can represent
	// all of the values of the type of the operand with unsigned integer type,
	// then the operand with unsigned integer type is converted to the type of the
	// operand with signed integer type.
	var signed Type
	switch {
	case abi.isSignedInteger(ak):
		// a is signed
		signed = a
		if a.Size() > b.Size() {
			return a
		}
	default:
		// b is signed
		signed = b
		if b.Size() > a.Size() {
			return b
		}
	}

	// Otherwise, both operands are converted to the unsigned integer type
	// corresponding to the type of the operand with signed integer type.
	switch signed.Kind() {
	case Char, SChar:
		return ast.kinds[UChar]
	case Short:
		return ast.kinds[UShort]
	case Int:
		return ast.kinds[UInt]
	case Long:
		return ast.kinds[ULong]
	case LongLong:
		return ast.kinds[ULongLong]
	case Int128:
		return ast.kinds[UInt128]
	}

	panic(todo(""))
}

// [0]6.3.1.1-2
//
// If an int can represent all values of the original type, the value is
// converted to an int; otherwise, it is converted to an unsigned int. These
// are called the integer promotions. All other types are unchanged by the
// integer promotions.
func integerPromotionKind(k Kind) Kind {
	switch k {
	case Char, SChar, UChar, Short, UShort:
		return Int
	default:
		return k
	}
}

func integerPromotion(t Type) Type {
	switch t.Kind() {
	case Char, SChar, UChar, Short, UShort:
		return t.(*PredefinedType).c.ast.kinds[Int]
	default:
		return t
	}
}

// Attributes represent selected values from __attribute__ constructs.
//
// See also https://gcc.gnu.org/onlinedocs/gcc/Attribute-Syntax.html
type Attributes struct {
	aligned    int64
	vectorSize int64

	isNonZero bool
}

func newAttributes() *Attributes {
	return &Attributes{
		aligned:    -1,
		vectorSize: -1,
	}
}

func (n *Attributes) setAligned(v int64)    { n.aligned = v; n.isNonZero = true }
func (n *Attributes) setVectorSize(v int64) { n.vectorSize = v; n.isNonZero = true }

// Aligned returns N from __attribute__(aligned(N)) or -1 if not
// present/valid.
func (n *Attributes) Aligned() int64 { return n.aligned }

// VectorSize returns N from __attribute__(vector_size(N)) or -1 if not
// present/valid.
//
// The vector_size attribute is only applicable to integral and floating
// scalars, otherwise it's ignored.
func (n *Attributes) VectorSize() int64 { return n.vectorSize }

type attributer struct{ p *Attributes }

// Attributes implemets Type.
func (n attributer) Attributes() *Attributes { return n.p }
