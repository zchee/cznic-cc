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
)

var (
	_ Type = (*attributedType)(nil)
	_ Type = (*baseType)(nil)

	_ typeDescriptor = (*DeclarationSpecifiers)(nil)
	_ typeDescriptor = (*SpecifierQualifierList)(nil)
	_ typeDescriptor = (*TypeQualifiers)(nil)

	noType = &baseType{}
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

	// FieldAlign returns the alignment in bytes of a value of this type
	// when used as a field in a struct.
	FieldAlign() int

	// Kind returns the specific kind of this type.
	Kind() Kind

	// Size returns the number of bytes needed to store a value of the
	// given type.
	//TODO Size() uintptr

	// atomic reports whether type has type qualifier "_Atomic".
	atomic() bool

	// auto reports whether type has storage class specifier "auto".
	auto() bool

	// extern reports whether type has storage class specifier "extern".
	extern() bool

	// hasConst reports whether type has type qualifier "const".
	hasConst() bool

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

	// threadLocal reports whether type has storage class specifier
	// "_Thread_local".
	threadLocal() bool

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
)

type typeDescriptor interface {
	Node
	isTypeDescriptor()
}

type baseTypeFlag uint16

const (
	// function specifier
	btfInline baseTypeFlag = 1 << iota
	btfNoReturn

	// storage class
	btfAuto
	btfExtern
	btfRegister
	btfStatic
	btfThreadLocal
	btfTypedef

	// type qualifier
	btfAtomic
	btfConst
	btfRestrict
	btfVolatile
)

type baseType struct {
	flags baseTypeFlag

	align      byte
	fieldAlign byte
	kind       byte
}

func (t *baseType) check(ctx *context, td typeDescriptor) Type {
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
func (t *baseType) atomic() bool { return t.flags&btfAtomic != 0 }

// Attributes implements Type.
func (t *baseType) Attributes() (a []*AttributeSpecifier) { return nil }

// auto implements Type.
func (t *baseType) auto() bool { return t.flags&btfAuto != 0 }

// Align implements Type.
func (t *baseType) Align() int { return int(t.align) }

// extern implements Type.
func (t *baseType) extern() bool { return t.flags&btfExtern != 0 }

// hasConst implements Type.
func (t *baseType) hasConst() bool { return t.flags&btfConst != 0 }

// FieldAlign implements Type.
func (t *baseType) FieldAlign() int { return int(t.fieldAlign) }

// inline implements Type.
func (t *baseType) inline() bool { return t.flags&btfInline != 0 }

// Kind implements Type.
func (t *baseType) Kind() Kind { return Kind(t.kind) }

// noReturn implements Type.
func (t *baseType) noReturn() bool { return t.flags&btfNoReturn != 0 }

// register implements Type.
func (t *baseType) register() bool { return t.flags&btfRegister != 0 }

// restrict implements Type.
func (t *baseType) restrict() bool { return t.flags&btfTypedef != 0 }

// static implements Type.
func (t *baseType) static() bool { return t.flags&btfStatic != 0 }

// threadLocal implements Type.
func (t *baseType) threadLocal() bool { return t.flags&btfThreadLocal != 0 }

// typedef implements Type.
func (t *baseType) typedef() bool { return t.flags&btfTypedef != 0 }

// volatile implements Type.
func (t *baseType) volatile() bool { return t.flags&btfVolatile != 0 }

type attributedType struct {
	Type
	attr []*AttributeSpecifier
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
	typ            Type // T in T*
	typeQualifiers Type

	align      byte
	fieldAlign byte
}

// Align implements Type.
func (t *pointerType) Align() int { return int(t.align) }

// Attributes implements Type.
func (t *pointerType) Attributes() (a []*AttributeSpecifier) { return t.typ.Attributes() }

// FieldAlign implements Type.
func (t *pointerType) FieldAlign() int { return int(t.fieldAlign) }

// Kind implements Type.
func (t *pointerType) Kind() Kind { return Ptr }

// atomic implements Type.
func (t *pointerType) atomic() bool { return t.typ.atomic() }

// auto implements Type.
func (t *pointerType) auto() bool { return t.typ.auto() }

// extern implements Type.
func (t *pointerType) extern() bool { return t.typ.extern() }

// hasConst implements Type.
func (t *pointerType) hasConst() bool { return t.typ.hasConst() }

// inline implements Type.
func (t *pointerType) inline() bool { return t.typ.inline() }

// noReturn implements Type.
func (t *pointerType) noReturn() bool { return t.typ.noReturn() }

// register implements Type.
func (t *pointerType) register() bool { return t.typ.register() }

// restrict implements Type.
func (t *pointerType) restrict() bool { return t.typ.restrict() }

// static implements Type.
func (t *pointerType) static() bool { return t.typ.static() }

// threadLocal implements Type.
func (t *pointerType) threadLocal() bool { return t.typ.threadLocal() }

// typedef implements Type.
func (t *pointerType) typedef() bool { return t.typ.typedef() }

// volatile implements Type.
func (t *pointerType) volatile() bool { return t.typ.volatile() }
