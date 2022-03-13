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
		Double:            true,
		Enum:              true,
		Float:             true,
		Int:               true,
		Long:              true,
		LongDouble:        true,
		LongLong:          true,
		SChar:             true,
		Short:             true,
		UChar:             true,
		UInt:              true,
		ULong:             true,
		ULongLong:         true,
		UShort:            true,
		Int128:            true,
		UInt128:           true,
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
	Double            // double
	Enum              // enum
	Float             // float
	Float128          // _Float128
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

	// FieldAlign reports the minimum alignment required by a type when it's used
	// in a struct/union.
	FieldAlign() int

	// IsIncomplete reports whether the size of a type is not determined.
	IsIncomplete() bool

	// Kind reports the kind of a type.
	Kind() Kind

	// Size reports the size of a type in bytes. Incomplete or invalid types may
	// report a negative size.
	Size() int64

	String() string

	str(b *strings.Builder, useTag bool) *strings.Builder
}

type InvalidType struct{}

// Align implements Type.
func (n *InvalidType) Align() int { return 1 }

// FieldAlign implements Type.
func (n *InvalidType) FieldAlign() int { return 1 }

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
	ast  *AST
	kind Kind
}

func newPredefinedType(ast *AST, kind Kind) *PredefinedType {
	return &PredefinedType{ast: ast, kind: kind}
}

// Align implements Type.
func (n *PredefinedType) Align() int {
	if n == nil {
		return 1
	}

	if x, ok := n.ast.ABI.types[n.kind]; ok {
		return x.align
	}

	return 1
}

// FieldAlign implements Type.
func (n *PredefinedType) FieldAlign() int {
	if n == nil {
		return 1
	}

	if x, ok := n.ast.ABI.types[n.kind]; ok {
		return x.fieldAlign
	}

	return 1
}

// String implements Type.
func (n *PredefinedType) String() string { return n.str(&strings.Builder{}, false).String() }

func (n *PredefinedType) str(b *strings.Builder, useTag bool) *strings.Builder {
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

	if x, ok := n.ast.ABI.types[n.kind]; ok {
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

// Name returns the name token of n. The result can a zero value, like in `void
// f(int) { ... }`.
func (n *Parameter) Name() Token {
	if n.Declarator != nil {
		return n.Declarator.DirectDeclarator.name().Token
	}

	return n.name
}

type FunctionType struct {
	result  typer
	fp      []*Parameter
	minArgs int
	maxArgs int // -1: unlimited
}

func newFunctionType(c *ctx, result Type, fp []*ParameterDeclaration, isVariadic bool) (r *FunctionType) {
	r = &FunctionType{result: newTyper(result), minArgs: len(fp), maxArgs: len(fp)}
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

// Result reports the result type of n.
func (n *FunctionType) Result() Type { return n.result.Type() }

// Align implements Type.
func (n *FunctionType) Align() int { return 1 }

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
	ast  *AST
	elem typer
}

func newPointerType(ast *AST, elem Type) *PointerType {
	return &PointerType{ast: ast, elem: newTyper(elem)}
}

// Elem returns the type n points to.
func (n *PointerType) Elem() Type { return n.elem.Type() }

// Align implements Type.
func (n *PointerType) Align() int {
	if n == nil {
		return 1
	}

	if x, ok := n.ast.ABI.types[Ptr]; ok {
		return x.align
	}

	return 1
}

// FieldAlign implements Type.
func (n *PointerType) FieldAlign() int {
	if n == nil {
		return 1
	}

	if x, ok := n.ast.ABI.types[Ptr]; ok {
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

	if x, ok := n.ast.ABI.types[Ptr]; ok {
		return x.size
	}

	return -1
}

// String implements Type.
func (n *PointerType) String() string { return n.str(&strings.Builder{}, false).String() }

func (n *PointerType) str(b *strings.Builder, useTag bool) *strings.Builder {
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

	// Additional bit offset to offset bytes. Non zero only for bit fields but can
	// be zero even for a bit field, for example, the first bit field after a non
	// bit field will have offsetBits zero.
	depth      int
	offsetBits int

	isBitField bool
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

	align int
}

func (n *structType) field(nm string) *Field {
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
	forward *StructOrUnionSpecifier
	structType
}

func newStructType(tag string, fields []*Field, size int64, align int) (r *StructType) {
	r = &StructType{structType: structType{tag: tag, fields: fields, size: size, align: align}}
	return r
}

// Field returns member field nm of n or nil if n does not have such member.
func (n *StructType) Field(nm string) *Field {
	if n == nil {
		return nil
	}

	if n.forward != nil {
		if x, ok := n.forward.typ.(*StructType); ok {
			return x.Field(nm)
		}

		return nil
	}

	return n.field(nm)
}

// Align implements Type.
func (n *StructType) Align() int {
	if n == nil {
		return 1
	}

	if n.forward != nil {
		return n.forward.Type().Align()
	}

	return n.align
}

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

	return n.Size() < 0
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
	b.WriteString("struct")
	if n.tag != "" {
		b.WriteByte(' ')
		b.WriteString(n.tag)
	}
	if n.forward != nil || useTag {
		return b
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
	forward *StructOrUnionSpecifier
	structType
}

func newUnionType(tag string, fields []*Field, size int64, align int) *UnionType {
	return &UnionType{structType: structType{tag: tag, fields: fields, size: size, align: align}}
}

// Field returns member field nm of n or nil if n does not have such member.
func (n *UnionType) Field(nm string) *Field {
	if n == nil {
		return nil
	}

	if n.forward != nil {
		if x, ok := n.forward.typ.(*StructType); ok {
			return x.Field(nm)
		}

		return nil
	}

	return n.field(nm)
}

// Align implements Type.
func (n *UnionType) Align() int {
	if n == nil {
		return 1
	}

	if n.forward != nil {
		return n.forward.Type().Align()
	}

	return n.align
}

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

	return n.Size() < 0
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
	b.WriteString("union")
	if n.tag != "" {
		b.WriteByte(' ')
		b.WriteString(n.tag)
	}
	if n.forward != nil || useTag {
		return b
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

type ArrayType struct {
	elem  typer
	elems int64
}

func newArrayType(elem Type, elems int64) (r *ArrayType) {
	r = &ArrayType{elem: newTyper(elem), elems: elems}
	return r
}

// Elem reports the element type of n.
func (n *ArrayType) Elem() Type { return n.elem.Type() }

// Align implements Type.
func (n *ArrayType) Align() int {
	if n == nil {
		return 1
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
	b.WriteString("array of ")
	if !n.IsIncomplete() {
		fmt.Fprintf(b, "%d ", n.elems)
	}
	n.Elem().str(b, true)
	return b
}

type EnumType struct {
	enums   []*Enumerator
	forward *EnumSpecifier
	tag     string
	typ     typer
}

func newEnumType(tag string, typ Type, enums []*Enumerator) *EnumType {
	return &EnumType{tag: tag, typ: newTyper(typ), enums: enums}
}

// Align implements Type.
func (n *EnumType) Align() int {
	if n == nil {
		return 1
	}

	if n.forward != nil {
		return n.forward.Type().Align()
	}

	return n.typ.Type().Align()
}

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

func isIntegerType(t Type) bool { return integerKinds[t.Kind()] }

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

	ast := a.(*PredefinedType).ast
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
		return t.(*PredefinedType).ast.kinds[Int]
	default:
		return t
	}
}
