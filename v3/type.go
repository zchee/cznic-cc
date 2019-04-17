// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Parts of the documentation are modified versions originating in the Go
// project, particularly the reflect package, license of which is reproduced
// below.
// ----------------------------------------------------------------------------
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the GO-LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

var (
	_ Type = (*arrayType)(nil)
	_ Type = (*attributedType)(nil)
	_ Type = (*pointerType)(nil)
	_ Type = (*typeBase)(nil)

	_ typeDescriptor = (*DeclarationSpecifiers)(nil)
	_ typeDescriptor = (*SpecifierQualifierList)(nil)
	_ typeDescriptor = (*TypeQualifiers)(nil)

	noType = &typeBase{}
)

// Type is the representation of a C type.
//
// Not all methods apply to all kinds of types. Restrictions, if any, are noted
// in the documentation for each method. Use the Kind method to find out the
// kind of type before calling kind-specific methods. Calling a method
// inappropriate to the kind of type causes a run-time panic.
type Type interface {

	// Align returns the alignment in bytes of a value of this type when
	// allocated in memory.
	Align() int

	// Attributes returns type's attributes, if any.
	Attributes() []*AttributeSpecifier

	// Elem returns a type's element type. It panics if the type's Kind is
	// not Array or Ptr.
	Elem() Type

	// FieldAlign returns the alignment in bytes of a value of this type
	// when used as a field in a struct.
	FieldAlign() int

	// Kind returns the specific kind of this type.
	Kind() Kind

	// Size returns the number of bytes needed to store a value of the
	// given type. It panics if type is incomplete.
	Size() uintptr

	// String implements fmt.Stringer.
	String() string

	// atomic reports whether type has type qualifier "_Atomic".
	atomic() bool

	// auto reports whether type has storage class specifier "auto".
	auto() bool

	// extern reports whether type has storage class specifier "extern".
	extern() bool

	// hasConst reports whether type has type qualifier "const".
	hasConst() bool

	// Incomplete reports whether type is incomplete.
	Incomplete() bool

	// inline reports whether type has function specifier "inline".
	inline() bool

	// noReturn reports whether type has function specifier "_NoReturn".
	noReturn() bool

	// register reports whether type has storage class specifier
	// "register".
	register() bool

	// restrict reports whether type has type qualifier "restrict".
	restrict() bool

	// static reports whether type has storage class specifier "static".
	static() bool

	string(*strings.Builder)

	// threadLocal reports whether type has storage class specifier
	// "_Thread_local".
	threadLocal() bool

	base() typeBase

	// typedef reports whether type has storage class specifier "typedef".
	typedef() bool

	// volatile reports whether type has type qualifier "volatile".
	volatile() bool
}

// A Kind represents the specific kind of type that a Type represents. The zero Kind is not a valid kind.
type Kind uint

const (
	maxTypeSpecifiers = 4 // eg. long long unsigned int
)

var (
	validTypeSpecifiers = map[[maxTypeSpecifiers]TypeSpecifierCase]byte{

		// [2], 6.7.2 Type specifiers, 2.

		//TODO atomic-type-specifier
		{TypeSpecifierBool}:                         byte(Bool),
		{TypeSpecifierChar, TypeSpecifierSigned}:    byte(SChar),
		{TypeSpecifierChar, TypeSpecifierUnsigned}:  byte(UChar),
		{TypeSpecifierChar}:                         byte(Char),
		{TypeSpecifierDouble, TypeSpecifierComplex}: byte(ComplexDouble),
		{TypeSpecifierDouble}:                       byte(Double),
		{TypeSpecifierEnum}:                         byte(Enum),
		{TypeSpecifierFloat, TypeSpecifierComplex}:  byte(ComplexFloat),
		{TypeSpecifierFloat}:                        byte(Float),
		{TypeSpecifierInt, TypeSpecifierLong, TypeSpecifierLong, TypeSpecifierSigned}:   byte(LongLong),
		{TypeSpecifierInt, TypeSpecifierLong, TypeSpecifierLong, TypeSpecifierUnsigned}: byte(ULongLong),
		{TypeSpecifierInt, TypeSpecifierLong, TypeSpecifierLong}:                        byte(LongLong),
		{TypeSpecifierInt, TypeSpecifierLong, TypeSpecifierSigned}:                      byte(Long),
		{TypeSpecifierInt, TypeSpecifierLong, TypeSpecifierUnsigned}:                    byte(ULong),
		{TypeSpecifierInt, TypeSpecifierLong}:                                           byte(Long),
		{TypeSpecifierInt, TypeSpecifierSigned}:                                         byte(Int),
		{TypeSpecifierInt, TypeSpecifierUnsigned}:                                       byte(UInt),
		{TypeSpecifierInt}: byte(Int),
		{TypeSpecifierLong, TypeSpecifierDouble, TypeSpecifierComplex}: byte(ComplexLongDouble),
		{TypeSpecifierLong, TypeSpecifierDouble}:                       byte(LongDouble),
		{TypeSpecifierLong, TypeSpecifierLong, TypeSpecifierSigned}:    byte(LongLong),
		{TypeSpecifierLong, TypeSpecifierLong, TypeSpecifierUnsigned}:  byte(ULongLong),
		{TypeSpecifierLong, TypeSpecifierLong}:                         byte(LongLong),
		{TypeSpecifierLong, TypeSpecifierSigned}:                       byte(Long),
		{TypeSpecifierLong, TypeSpecifierUnsigned}:                     byte(ULong),
		{TypeSpecifierLong}: byte(Long),
		{TypeSpecifierShort, TypeSpecifierInt, TypeSpecifierSigned}:   byte(Short),
		{TypeSpecifierShort, TypeSpecifierInt, TypeSpecifierUnsigned}: byte(UShort),
		{TypeSpecifierShort, TypeSpecifierInt}:                        byte(Short),
		{TypeSpecifierShort, TypeSpecifierSigned}:                     byte(Short),
		{TypeSpecifierShort, TypeSpecifierUnsigned}:                   byte(UShort),
		{TypeSpecifierShort}:                                          byte(Short),
		{TypeSpecifierSigned}:                                         byte(Int),
		{TypeSpecifierStruct}:                                         byte(Struct),
		{TypeSpecifierTypedefName}:                                    byte(TypedefName), //TODO
		{TypeSpecifierUnsigned}:                                       byte(UInt),
		{TypeSpecifierVoid}:                                           byte(Void),

		// GCC Extensions.

		{TypeSpecifierChar, TypeSpecifierComplex}:                         byte(ComplexChar),
		{TypeSpecifierComplex}:                                            byte(ComplexDouble),
		{TypeSpecifierDecimal128}:                                         byte(Decimal128),
		{TypeSpecifierDecimal32}:                                          byte(Decimal32),
		{TypeSpecifierDecimal64}:                                          byte(Decimal64),
		{TypeSpecifierFloat128}:                                           byte(Float128),
		{TypeSpecifierFloat32x}:                                           byte(Float32x),
		{TypeSpecifierFloat32}:                                            byte(Float32),
		{TypeSpecifierFloat64x}:                                           byte(Float64x),
		{TypeSpecifierFloat64}:                                            byte(Float64),
		{TypeSpecifierInt, TypeSpecifierComplex}:                          byte(ComplexInt),
		{TypeSpecifierInt, TypeSpecifierLong, TypeSpecifierComplex}:       byte(ComplexLong),
		{TypeSpecifierInt128, TypeSpecifierSigned}:                        byte(Int128),
		{TypeSpecifierInt128, TypeSpecifierUnsigned}:                      byte(UInt128),
		{TypeSpecifierInt128}:                                             byte(Int128),
		{TypeSpecifierLong, TypeSpecifierComplex}:                         byte(ComplexLong),
		{TypeSpecifierLong, TypeSpecifierLong, TypeSpecifierComplex}:      byte(ComplexLongLong),
		{TypeSpecifierShort, TypeSpecifierComplex}:                        byte(ComplexUShort),
		{TypeSpecifierShort, TypeSpecifierUnsigned, TypeSpecifierComplex}: byte(ComplexShort),
		{TypeSpecifierTypeofExpr}:                                         byte(typeofExpr), //TODO
		{TypeSpecifierTypeofType}:                                         byte(typeofType), //TODO
		{TypeSpecifierUnsigned, TypeSpecifierComplex}:                     byte(ComplexUInt),
	}

	integerTypes = [maxKind]bool{
		Bool:      true,
		Char:      true,
		Enum:      true,
		Int:       true,
		Long:      true,
		LongLong:  true,
		SChar:     true,
		Short:     true,
		UChar:     true,
		UInt:      true,
		ULong:     true,
		ULongLong: true,
		UShort:    true,
	}

	unsignedTypes = [maxKind]bool{
		Bool:      true,
		UChar:     true,
		UInt:      true,
		ULong:     true,
		ULongLong: true,
		UShort:    true,
	}
)

type typeDescriptor interface {
	Node
	isTypeDescriptor()
}

type flag uint16

const (
	// function specifier
	fInline flag = 1 << iota
	fNoReturn

	// storage class
	fAuto
	fExtern
	fRegister
	fStatic
	fThreadLocal
	fTypedef

	// type qualifier
	fAtomic
	fConst
	fRestrict
	fVolatile

	// other
	fIncomplete
)

type typeBase struct {
	size uintptr

	flags flag

	align      byte
	fieldAlign byte
	kind       byte
}

func (t *typeBase) check(ctx *context, td typeDescriptor) Type {
	var alignmentSpecifiers []*AlignmentSpecifier
	var attributeSpecifiers []*AttributeSpecifier
	var typeSpecifiers []*TypeSpecifier
	switch n := td.(type) {
	case nil:
		t.kind = byte(TypeSpecifierInt)
		return t //TODO configuration flag
	case *DeclarationSpecifiers:
		for ; n != nil; n = n.DeclarationSpecifiers {
			switch n.Case {
			case DeclarationSpecifiersStorage: // StorageClassSpecifier DeclarationSpecifiers
				// nop
			case DeclarationSpecifiersTypeSpec: // TypeSpecifier DeclarationSpecifiers
				typeSpecifiers = append(typeSpecifiers, n.TypeSpecifier)
			case DeclarationSpecifiersTypeQual: // TypeQualifier DeclarationSpecifiers
				// nop
			case DeclarationSpecifiersFunc: // FunctionSpecifier DeclarationSpecifiers
				// nop
			case DeclarationSpecifiersAlignSpec: // AlignmentSpecifier DeclarationSpecifiers
				alignmentSpecifiers = append(alignmentSpecifiers, n.AlignmentSpecifier)
			case DeclarationSpecifiersAttribute: // AttributeSpecifier DeclarationSpecifiers
				attributeSpecifiers = append(attributeSpecifiers, n.AttributeSpecifier)
			default:
				panic("internal error") //TODOOK
			}
		}
	case *SpecifierQualifierList:
		for ; n != nil; n = n.SpecifierQualifierList {
			switch n.Case {
			case SpecifierQualifierListTypeSpec: // TypeSpecifier SpecifierQualifierList
				typeSpecifiers = append(typeSpecifiers, n.TypeSpecifier)
			case SpecifierQualifierListTypeQual: // TypeQualifier SpecifierQualifierList
				// nop
			case SpecifierQualifierListAlignSpec: // AlignmentSpecifier SpecifierQualifierList
				alignmentSpecifiers = append(alignmentSpecifiers, n.AlignmentSpecifier)
			case SpecifierQualifierListAttribute: // AttributeSpecifier SpecifierQualifierList
				attributeSpecifiers = append(attributeSpecifiers, n.AttributeSpecifier)
			default:
				panic("internal error") //TODOOK
			}
		}
	case *TypeQualifiers:
		for ; n != nil; n = n.TypeQualifiers {
			if n.Case == TypeQualifiersAttribute {
				attributeSpecifiers = append(attributeSpecifiers, n.AttributeSpecifier)
			}
		}
	default:
		panic(fmt.Sprintf("internal error: %v: %T", n.Position(), n)) //TODOOK
	}

	if len(typeSpecifiers) > maxTypeSpecifiers {
		ctx.err(typeSpecifiers[maxTypeSpecifiers].Position(), "too many type specifiers")
		typeSpecifiers = typeSpecifiers[:maxTypeSpecifiers]
	}

	sort.Slice(typeSpecifiers, func(i, j int) bool {
		return typeSpecifiers[i].Case < typeSpecifiers[j].Case
	})
	var k [maxTypeSpecifiers]TypeSpecifierCase
	for i, v := range typeSpecifiers {
		k[i] = v.Case
	}
	if len(typeSpecifiers) == 0 {
		k[0] = TypeSpecifierInt //TODO configuration flag
	}
	var ok bool
	if t.kind, ok = validTypeSpecifiers[k]; !ok {
		// panic(fmt.Sprintf("%v: %v", td.Position(), k[:len(typeSpecifiers)]))
		ctx.err(td.Position(), "invalid type specifiers combination")
		k[0] = TypeSpecifierInt
		return t

	}
	switch len(alignmentSpecifiers) {
	case 0:
		//TODO set alignment from model
	case 1:
		align := alignmentSpecifiers[0].align()
		if align > math.MaxUint8 {
			panic("internal error") //TODOOK
		}
		t.align = byte(align)
		t.fieldAlign = t.align
	default:
		ctx.err(alignmentSpecifiers[1].Position(), "multiple alignment specifiers")
	}

	switch t.Kind() {
	case TypedefName:
		//TODO
	case typeofExpr:
		//TODO
	case typeofType:
		//TODO
	case Struct:
		//TODO
	case Void:
		// nop
	case Enum:
		//TODO
	default:
		if t.align != 0 {
			break
		}

		if v, ok := ctx.cfg.ABI.Types[t.Kind()]; ok {
			t.align = byte(v.Align)
			t.fieldAlign = byte(v.FieldAlign)
			break
		}

		ctx.err(td.Position(), "missing model item for %s", t.Kind())
	}

	if len(attributeSpecifiers) != 0 {
		return &attributedType{Type: t, attr: attributeSpecifiers}
	}

	return t
}

// atomic implements Type.
func (t *typeBase) atomic() bool { return t.flags&fAtomic != 0 }

// Attributes implements Type.
func (t *typeBase) Attributes() (a []*AttributeSpecifier) { return nil }

// auto implements Type.
func (t *typeBase) auto() bool { return t.flags&fAuto != 0 }

// Align implements Type.
func (t *typeBase) Align() int { return int(t.align) }

// base implements Type.
func (t *typeBase) base() typeBase { return *t }

// Elem implements Type.
func (t *typeBase) Elem() Type { panic("Elem(): inapropriate type") }

// extern implements Type.
func (t *typeBase) extern() bool { return t.flags&fExtern != 0 }

// hasConst implements Type.
func (t *typeBase) hasConst() bool { return t.flags&fConst != 0 }

// FieldAlign implements Type.
func (t *typeBase) FieldAlign() int { return int(t.fieldAlign) }

// Incomplete implements Type.
func (t *typeBase) Incomplete() bool { return t.flags&fIncomplete != 0 }

// inline implements Type.
func (t *typeBase) inline() bool { return t.flags&fInline != 0 }

// Kind implements Type.
func (t *typeBase) Kind() Kind { return Kind(t.kind) }

// noReturn implements Type.
func (t *typeBase) noReturn() bool { return t.flags&fNoReturn != 0 }

// register implements Type.
func (t *typeBase) register() bool { return t.flags&fRegister != 0 }

// restrict implements Type.
func (t *typeBase) restrict() bool { return t.flags&fRestrict != 0 }

// Size implements Type.
func (t *typeBase) Size() uintptr {
	if t.Incomplete() {
		panic("Size(): incomplete type") //TODOOK
	}

	if t.size == 0 {
		//TODO panic("TODO")
	}
	return t.size
}

// static implements Type.
func (t *typeBase) static() bool { return t.flags&fStatic != 0 }

// threadLocal implements Type.
func (t *typeBase) threadLocal() bool { return t.flags&fThreadLocal != 0 }

// typedef implements Type.
func (t *typeBase) typedef() bool { return t.flags&fTypedef != 0 }

// volatile implements Type.
func (t *typeBase) volatile() bool { return t.flags&fVolatile != 0 }

// String implements Type.
func (t *typeBase) String() string {
	var b strings.Builder
	t.string(&b)
	return strings.TrimSpace(b.String())
}

// string implements Type.
func (t *typeBase) string(b *strings.Builder) {
	if t.atomic() {
		b.WriteString("atomic ")
	}
	if t.auto() {
		b.WriteString("auto ")
	}
	if t.extern() {
		b.WriteString("extern ")
	}
	if t.hasConst() {
		b.WriteString("const ")
	}
	if t.inline() {
		b.WriteString("inline ")
	}
	if t.noReturn() {
		b.WriteString("_NoReturn ")
	}
	if t.register() {
		b.WriteString("register ")
	}
	if t.restrict() {
		b.WriteString("restrict ")
	}
	if t.static() {
		b.WriteString("static ")
	}
	if t.threadLocal() {
		b.WriteString("_Thread_local ")
	}
	if t.typedef() {
		b.WriteString("typedef ")
	}
	if t.volatile() {
		b.WriteString("volatile ")
	}
	switch k := t.Kind(); k {
	case Enum:
		b.WriteString("enum ")
	case Invalid:
		// nop
	case Struct:
		b.WriteString("struct ")
	case typeofExpr, typeofType, Ptr:
		panic("internal error") // TODOOK
	default:
		b.WriteString(k.String())
		b.WriteByte(' ')
	}
}

type attributedType struct {
	Type
	attr []*AttributeSpecifier
}

// String implements Type.
func (t *attributedType) String() string {
	var b strings.Builder
	t.string(&b)
	return strings.TrimSpace(b.String())
}

// string implements Type.
func (t *attributedType) string(b *strings.Builder) {
	for _, v := range t.attr {
		panic(v.Position())
	}
	t.Type.string(b)
}

// Attributes implements Type.
func (t *attributedType) Attributes() []*AttributeSpecifier { return t.attr }

type field struct {
	offset uintptr // In bytes from start of the struct.
	typ    Type

	name StringID // Can be zero.

	isBitField bool // Following fields are valid only if this field is true.

	bitOffset byte // In bits from bit 0 within the field.
	bits      byte // Width of the bit field. Valid only when isBitField is true.
	mask      byte // bits: 3, bitOffset: 2 -> 0x1c.
}

type structType struct { //TODO implement Type
	fields []field
	tag    StringID
}

type pointerType struct {
	typeBase

	elem           Type
	typeQualifiers Type
}

// Attributes implements Type.
func (t *pointerType) Attributes() (a []*AttributeSpecifier) { return t.elem.Attributes() }

// Elem implements Type.
func (t *pointerType) Elem() Type { return t.elem }

// String implements Type.
func (t *pointerType) String() string {
	var b strings.Builder
	t.string(&b)
	return strings.TrimSpace(b.String())
}

// string implements Type.
func (t *pointerType) string(b *strings.Builder) {
	if t := t.typeQualifiers; t != nil {
		t.string(b)
	}
	b.WriteString("pointer to ")
	t.elem.string(b)
}

type arrayType struct {
	typeBase

	elem Type
}

// Attributes implements Type.
func (t *arrayType) Attributes() (a []*AttributeSpecifier) { return t.elem.Attributes() }

// Elem implements Type.
func (t *arrayType) Elem() Type { return t.elem }
