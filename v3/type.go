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

const maxRank = 6

var (
	_ Field = (*field)(nil)

	_ Type = (*aliasType)(nil)
	_ Type = (*arrayType)(nil)
	_ Type = (*attributedType)(nil)
	_ Type = (*bitFieldType)(nil)
	_ Type = (*functionType)(nil)
	_ Type = (*pointerType)(nil)
	_ Type = (*structType)(nil)
	_ Type = (*taggedType)(nil)
	_ Type = (*typeBase)(nil)
	_ Type = noType

	noType = &typeBase{}

	_ typeDescriptor = (*DeclarationSpecifiers)(nil)
	_ typeDescriptor = (*SpecifierQualifierList)(nil)
	_ typeDescriptor = (*TypeQualifiers)(nil)
	_ typeDescriptor = noTypeDescriptor

	noTypeDescriptor = &DeclarationSpecifiers{}

	_ Object = (*AbstractDeclarator)(nil)
	_ Object = (*Declarator)(nil)
	_ Object = (*abstractObject)(nil)

	// [0]6.3.1.1-1
	//
	// Every integer type has an integer conversion rank defined as
	// follows:
	intConvRank = [maxKind]int{
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

	arithmeticTypes = [maxKind]bool{
		Bool:              true,
		Char:              true,
		ComplexDouble:     true,
		ComplexFloat:      true,
		ComplexLongDouble: true,
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
	}

	realTypes = [maxKind]bool{
		Bool:       true,
		Char:       true,
		Double:     true,
		Enum:       true,
		Float:      true,
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
	}
)

type noStorageClass struct{}

func (noStorageClass) auto() bool        { return false }
func (noStorageClass) extern() bool      { return false }
func (noStorageClass) register() bool    { return false }
func (noStorageClass) static() bool      { return false }
func (noStorageClass) threadLocal() bool { return false }
func (noStorageClass) typedef() bool     { return false }

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
	//TODO bits()

	// Alias returns the type this type aliases. Non typedef types return
	// themselves.
	Alias() Type

	// Align returns the alignment in bytes of a value of this type when
	// allocated in memory.
	Align() int

	// Attributes returns type's attributes, if any.
	Attributes() []*AttributeSpecifier

	// Decay returns itself for non array types and the pointer to array
	// element otherwise.
	Decay() Type

	// Elem returns a type's element type. It panics if the type's Kind is
	// valid but not Array or Ptr.
	Elem() Type

	// FieldAlign returns the alignment in bytes of a value of this type
	// when used as a field in a struct.
	FieldAlign() int

	// FieldByIndex returns the nested field corresponding to the index
	// sequence. It is equivalent to calling Field successively for each
	// index i.  It panics if the type's Kind is valid but not Struct.
	FieldByIndex(index []int) Field

	// FieldByName returns the struct field with the given name and a
	// boolean indicating if the field was found.
	FieldByName(name StringID) (Field, bool)

	// Incomplete reports whether type is incomplete.
	Incomplete() bool

	// IsArithmeticType report whether a type is an arithmetic type.
	IsArithmeticType() bool

	// IsBitFieldType report whether a type is for a bit field.
	IsBitFieldType() bool //TODO-

	// IsIntegerType report whether a type is an integer type.
	IsIntegerType() bool

	// IsRealType report whether a type is a real type.
	IsRealType() bool

	// IsScalarType report whether a type is a scalar type.
	IsScalarType() bool

	// IsVariadic reports whether a function type is variadic. It panics if
	// the type's Kind is valid but not Function.
	IsVariadic() bool

	// IsVLA reports whether array is a variable length array. It panics if
	// the type's Kind is valid but not Array.
	IsVLA() bool

	// Kind returns the specific kind of this type.
	Kind() Kind

	// Len returns an array type's length.  It panics if the type's Kind is
	// valid but not Array.
	Len() uintptr

	// NumField returns a struct type's field count.  It panics if the
	// type's Kind is valid but not Struct.
	NumField() int

	// Parameters returns the parameters of a function type. It panics if
	// the type's Kind is valid but not Function.
	Parameters() []Object

	// Result returns the result type of a function type. It panics if the
	// type's Kind is valid but not Function.
	Result() Type

	// Size returns the number of bytes needed to store a value of the
	// given type. It panics if type is valid but incomplete.
	Size() uintptr

	// String implements fmt.Stringer.
	String() string

	// atomic reports whether type has type qualifier "_Atomic".
	atomic() bool

	// hasConst reports whether type has type qualifier "const".
	hasConst() bool

	// inline reports whether type has function specifier "inline".
	inline() bool

	IsSignedType() bool

	// noReturn reports whether type has function specifier "_NoReturn".
	noReturn() bool

	// restrict reports whether type has type qualifier "restrict".
	restrict() bool

	setLen(uintptr)

	string(*strings.Builder)

	base() typeBase

	underlyingType() Type

	// volatile reports whether type has type qualifier "volatile".
	volatile() bool
}

type Object interface {
	Declarator() *Declarator
	IsStatic() bool
	Name() StringID
	Type() Type
}

type abstractObject struct {
	typ Type
}

func (o *abstractObject) Declarator() *Declarator { return nil }
func (o *abstractObject) IsStatic() bool          { panic("TODO") } //TODO return o.Type().static() }
func (o *abstractObject) Name() StringID          { return 0 }
func (o *abstractObject) Type() Type              { return o.typ }

// A Field describes a single field in a struct/union.
type Field interface {
	BitFieldOffset() int
	BitFieldWidth() int
	IsBitField() bool
	Mask() uint64
	Offset() uintptr // In bytes from the beginning of the struct/union.
	Padding() int
	Type() Type // Field type.
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
		{TypeSpecifierStruct}:                                         byte(Struct),      //TODO
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
	auto() bool
	extern() bool
	register() bool
	static() bool
	threadLocal() bool
	typedef() bool
}

type storageClass byte

const (
	fAuto storageClass = 1 << iota
	fExtern
	fRegister
	fStatic
	fThreadLocal
	fTypedef
)

type flag uint16

const (
	// function specifier
	fInline flag = 1 << iota //TODO should go elsewhere
	fNoReturn

	// type qualifier
	fAtomic
	fConst
	fRestrict
	fVolatile

	// other
	fIncomplete
	fSigned // Valid only for integer types.
	fBitField
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
	if t.Kind() == Invalid {
		if t.kind, ok = validTypeSpecifiers[k]; !ok {
			// panic(fmt.Sprintf("%v: %v", td.Position(), k[:len(typeSpecifiers)]))
			ctx.err(td.Position(), "invalid type specifiers combination")
			k[0] = TypeSpecifierInt
			return t

		}
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

	abi := ctx.cfg.ABI
	switch k := t.Kind(); k {
	case typeofExpr:
		//TODO
	case typeofType:
		//TODO
	case Struct:
		//TODO
	case Union:
		//TODO
	case Void:
		// nop
	case Enum:
		if v, ok := abi.Types[Int]; ok {
			t.size = uintptr(abi.size(Int))
			t.align = byte(v.Align)
			t.fieldAlign = byte(v.FieldAlign)
		}
	default:
		if integerTypes[k] && abi.isSignedInteger(k) {
			t.flags |= fSigned
		}
		if v, ok := abi.Types[k]; ok {
			t.size = uintptr(abi.size(k))
			if t.align != 0 {
				break
			}

			t.align = byte(v.Align)
			t.fieldAlign = byte(v.FieldAlign)
			break
		}

		//TODO ctx.err(td.Position(), "missing model item for %s", t.Kind())
	}

	typ := Type(t)
	switch k := t.Kind(); k {
	case TypedefName:
		ts := typeSpecifiers[0]
		tok := ts.Token
		nm := tok.Value
		d := ts.resolvedIn.typedef(nm, tok)
		typ = &aliasType{nm: nm, typ: d.Type()}
	case Struct, Union:
		typ = typeSpecifiers[0].StructOrUnionSpecifier.typ
	case typeofExpr, typeofType:
		typ = typeSpecifiers[0].typ
	}

	if len(attributeSpecifiers) != 0 {
		typ = &attributedType{Type: typ, attr: attributeSpecifiers}
	}
	return typ
}

// atomic implements Type.
func (t *typeBase) atomic() bool { return t.flags&fAtomic != 0 }

// Attributes implements Type.
func (t *typeBase) Attributes() (a []*AttributeSpecifier) { return nil }

// Alias implements Type.
func (t *typeBase) Alias() Type { return t }

// Align implements Type.
func (t *typeBase) Align() int { return int(t.align) }

// base implements Type.
func (t *typeBase) base() typeBase { return *t }

// Decay implements Type.
func (t *typeBase) Decay() Type {
	if t.Kind() != Array {
		return t
	}

	panic(fmt.Errorf("%s: Decay of invalid type", t.Kind()))
}

// Elem implements Type.
func (t *typeBase) Elem() Type {
	if t.Kind() == Invalid {
		return t
	}

	panic(fmt.Errorf("%s: Elem of invalid type", t.Kind()))
}

// hasConst implements Type.
func (t *typeBase) hasConst() bool { return t.flags&fConst != 0 }

// FieldAlign implements Type.
func (t *typeBase) FieldAlign() int { return int(t.fieldAlign) }

// FieldByIndex implements Type.
func (t *typeBase) FieldByIndex([]int) Field {
	if t.Kind() == Invalid {
		return nil
	}

	panic(fmt.Errorf("%s: FieldByIndex of invalid type", t.Kind()))
}

// NumField implements Type.
func (t *typeBase) NumField() int {
	if t.Kind() == Invalid {
		return 0
	}

	panic(fmt.Errorf("%s: NumField of invalid type", t.Kind()))
}

// FieldByName implements Type.
func (t *typeBase) FieldByName(StringID) (Field, bool) {
	if t.Kind() == Invalid {
		return nil, false
	}

	panic(fmt.Errorf("%s: FieldByName of invalid type", t.Kind()))
}

// Incomplete implements Type.
func (t *typeBase) Incomplete() bool { return t.flags&fIncomplete != 0 }

// inline implements Type.
func (t *typeBase) inline() bool { return t.flags&fInline != 0 }

// IsIntegerType implements Type.
func (t *typeBase) IsIntegerType() bool { return integerTypes[t.kind] }

// IsArithmeticType implements Type.
func (t *typeBase) IsArithmeticType() bool { return arithmeticTypes[t.Kind()] }

// IsBitFieldType implements Type.
func (t *typeBase) IsBitFieldType() bool { return false }

// IsRealType implements Type.
func (t *typeBase) IsRealType() bool { return realTypes[t.Kind()] }

// IsScalarType implements Type.
func (t *typeBase) IsScalarType() bool { return t.IsArithmeticType() || t.Kind() == Ptr }

// IsSignedType implements Type.
func (t *typeBase) IsSignedType() bool {
	if !integerTypes[t.kind] {
		panic(fmt.Errorf("%s: IsSignedType of non-integer type", t.Kind()))
	}

	return t.flags&fSigned != 0
}

// IsVariadic implements Type.
func (t *typeBase) IsVariadic() bool {
	if t.Kind() == Invalid {
		return false
	}

	panic(fmt.Errorf("%s: IsVariadic of invalid type", t.Kind()))
}

// IsVLA implements Type.
func (t *typeBase) IsVLA() bool {
	if t.Kind() == Invalid {
		return false
	}

	panic(fmt.Errorf("%s: IsVLA of invalid type", t.Kind()))
}

// Kind implements Type.
func (t *typeBase) Kind() Kind { return Kind(t.kind) }

// Len implements Type.
func (t *typeBase) Len() uintptr { panic(fmt.Errorf("%s: Len of non-array type", t.Kind())) }

// noReturn implements Type.
func (t *typeBase) noReturn() bool { return t.flags&fNoReturn != 0 }

// restrict implements Type.
func (t *typeBase) restrict() bool { return t.flags&fRestrict != 0 }

// Parameters implements Type.
func (t *typeBase) Parameters() []Object {
	if t.Kind() == Invalid {
		return nil
	}

	panic(fmt.Errorf("%s: Parameters of invalid type", t.Kind()))
}

// Result implements Type.
func (t *typeBase) Result() Type {
	if t.Kind() == Invalid {
		return noType
	}

	panic(fmt.Errorf("%s: Result of invalid type", t.Kind()))
}

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

// setLen implements Type.
func (t *typeBase) setLen(uintptr) {
	if t.Kind() == Invalid {
		return
	}

	panic(fmt.Errorf("%s: setLen of non-array type", t.Kind()))
}

// underlyingType implements Type.
func (t *typeBase) underlyingType() Type { return t }

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
	spc := ""
	if t.atomic() {
		b.WriteString("atomic")
		spc = " "
	}
	if t.hasConst() {
		b.WriteString(spc)
		b.WriteString("const")
		spc = " "
	}
	if t.inline() {
		b.WriteString(spc)
		b.WriteString("inline")
		spc = " "
	}
	if t.noReturn() {
		b.WriteString(spc)
		b.WriteString("_NoReturn")
		spc = " "
	}
	if t.restrict() {
		b.WriteString(spc)
		b.WriteString("restrict")
		spc = " "
	}
	if t.volatile() {
		b.WriteString(spc)
		b.WriteString("volatile")
		spc = " "
	}
	b.WriteString(spc)
	switch k := t.Kind(); k {
	case Enum:
		b.WriteString("enum")
	case Invalid:
		// nop
	case Struct:
		b.WriteString("struct")
	case Union:
		b.WriteString("union")
	case Ptr:
		b.WriteString("pointer")
	case typeofExpr, typeofType:
		panic("internal error: " + k.String()) // TODOOK
	default:
		b.WriteString(k.String())
	}
}

type attributedType struct {
	Type
	attr []*AttributeSpecifier
}

// Alias implements Type.
func (t *attributedType) Alias() Type { return t }

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

type pointerType struct {
	typeBase

	elem           Type
	typeQualifiers Type
}

// Alias implements Type.
func (t *pointerType) Alias() Type { return t }

// Attributes implements Type.
func (t *pointerType) Attributes() (a []*AttributeSpecifier) { return t.elem.Attributes() }

// Decay implements Type.
func (t *pointerType) Decay() Type { return t }

// Elem implements Type.
func (t *pointerType) Elem() Type { return t.elem }

// underlyingType implements Type.
func (t *pointerType) underlyingType() Type { return t }

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
	t.Elem().string(b)
}

type arrayType struct {
	typeBase

	expr   *AssignmentExpression
	decay  Type
	elem   Type
	length uintptr

	vla bool
}

// Alias implements Type.
func (t *arrayType) Alias() Type { return t }

// IsVLA implements Type.
func (t *arrayType) IsVLA() bool { return t.vla }

// String implements Type.
func (t *arrayType) String() string {
	var b strings.Builder
	t.string(&b)
	return strings.TrimSpace(b.String())
}

// string implements Type.
func (t *arrayType) string(b *strings.Builder) {
	b.WriteString("array of ")
	if t.Len() != 0 {
		fmt.Fprintf(b, "%d ", t.Len())
	}
	t.Elem().string(b)
}

// Attributes implements Type.
func (t *arrayType) Attributes() (a []*AttributeSpecifier) { return t.elem.Attributes() }

// Decay implements Type.
func (t *arrayType) Decay() Type { return t.decay }

// Elem implements Type.
func (t *arrayType) Elem() Type { return t.elem }

// Len implements Type.
func (t *arrayType) Len() uintptr { return t.length }

// setLen implements Type.
func (t *arrayType) setLen(n uintptr) {
	t.typeBase.flags &^= fIncomplete
	t.length = n
	if t.Elem() != nil {
		t.size = t.length * t.Elem().Size()
	}
}

// underlyingType implements Type.
func (t *arrayType) underlyingType() Type { return t }

type aliasType struct {
	nm  StringID
	typ Type
}

// Alias implements Type.
func (t *aliasType) Alias() Type { return t.typ }

// Align implements Type.
func (t *aliasType) Align() int { return t.typ.Align() }

// Attributes implements Type.
func (t *aliasType) Attributes() (a []*AttributeSpecifier) { return nil }

// Decay implements Type.
func (t *aliasType) Decay() Type { return t.typ.Decay() }

// Elem implements Type.
func (t *aliasType) Elem() Type { return t.typ.Elem() }

// FieldAlign implements Type.
func (t *aliasType) FieldAlign() int { return t.typ.FieldAlign() }

// NumField implements Type.
func (t *aliasType) NumField() int { return t.typ.NumField() }

// FieldByIndex implements Type.
func (t *aliasType) FieldByIndex(i []int) Field { return t.typ.FieldByIndex(i) }

// FieldByName implements Type.
func (t *aliasType) FieldByName(s StringID) (Field, bool) { return t.typ.FieldByName(s) }

// Incomplete implements Type.
func (t *aliasType) Incomplete() bool { return t.typ.Incomplete() }

// IsArithmeticType implements Type.
func (t *aliasType) IsArithmeticType() bool { return t.typ.IsArithmeticType() }

// IsBitFieldType implements Type.
func (t *aliasType) IsBitFieldType() bool { return t.typ.IsBitFieldType() }

// IsIntegerType implements Type.
func (t *aliasType) IsIntegerType() bool { return t.typ.IsIntegerType() }

// IsRealType implements Type.
func (t *aliasType) IsRealType() bool { return t.typ.IsRealType() }

// IsScalarType implements Type.
func (t *aliasType) IsScalarType() bool { return t.typ.IsScalarType() }

// IsVLA implements Type.
func (t *aliasType) IsVLA() bool { return t.typ.IsVLA() }

// IsVariadic implements Type.
func (t *aliasType) IsVariadic() bool { return t.typ.IsVariadic() }

// Kind implements Type.
func (t *aliasType) Kind() Kind { return t.typ.Kind() }

// Len implements Type.
func (t *aliasType) Len() uintptr { return t.typ.Len() }

// Parameters implements Type.
func (t *aliasType) Parameters() []Object { return t.typ.Parameters() }

// Result implements Type.
func (t *aliasType) Result() Type { return t.typ.Result() }

// Size implements Type.
func (t *aliasType) Size() uintptr { return t.typ.Size() }

// String implements Type.
func (t *aliasType) String() string { return t.nm.String() }

// atomic implements Type.
func (t *aliasType) atomic() bool { return t.typ.atomic() }

// base implements Type.
func (t *aliasType) base() typeBase { return t.typ.base() }

// hasConst implements Type.
func (t *aliasType) hasConst() bool { return t.typ.hasConst() }

// inline implements Type.
func (t *aliasType) inline() bool { return t.typ.inline() }

// IsSignedType implements Type.
func (t *aliasType) IsSignedType() bool { return t.typ.IsSignedType() }

// noReturn implements Type.
func (t *aliasType) noReturn() bool { return t.typ.noReturn() }

// restrict implements Type.
func (t *aliasType) restrict() bool { return t.typ.restrict() }

// setLen implements Type.
func (t *aliasType) setLen(n uintptr) { t.typ.setLen(n) }

// string implements Type.
func (t *aliasType) string(b *strings.Builder) { b.WriteString(t.nm.String()) }

func (t *aliasType) underlyingType() Type { return t.typ.underlyingType() }

// volatile implements Type.
func (t *aliasType) volatile() bool { return t.typ.volatile() }

type field struct {
	bitFieldMask uint64  // bits: 3, bitOffset: 2 -> 0x1c. Valid only when isBitField is true.
	offset       uintptr // In bytes from start of the struct.
	typ          Type

	name StringID // Can be zero.

	isBitField bool

	bitFieldOffset byte // In bits from bit 0 within the field. Valid only when isBitField is true.
	bitFieldWidth  byte // Width of the bit field in bits. Valid only when isBitField is true.
	pad            byte
}

func (f *field) BitFieldOffset() int { return int(f.bitFieldOffset) }
func (f *field) BitFieldWidth() int  { return int(f.bitFieldWidth) }
func (f *field) IsBitField() bool    { return f.isBitField }
func (f *field) Mask() uint64        { return f.bitFieldMask }
func (f *field) Offset() uintptr     { return f.offset }
func (f *field) Padding() int        { return int(f.pad) } // N/A for bitfields
func (f *field) Type() Type          { return f.typ }

func (f *field) string(b *strings.Builder) {
	b.WriteString(f.name.String())
	if f.isBitField {
		fmt.Fprintf(b, ":%d", f.bitFieldWidth)
	}
	b.WriteByte(' ')
	f.typ.string(b)
}

type structType struct { //TODO implement Type
	*typeBase

	attr   []*AttributeSpecifier
	fields []*field
	m      map[StringID]*field

	tag StringID
}

// Alias implements Type.
func (t *structType) Alias() Type { return t }

func (t *structType) check(ctx *context, n Node) *structType {
	if t == nil {
		return nil
	}

	return ctx.cfg.ABI.layout(ctx, n, t)
}

// Decay implements Type.
func (t *structType) Decay() Type { return t }

func (t *structType) underlyingType() Type { return t }

// String implements Type.
func (t *structType) String() string {
	var b strings.Builder
	t.string(&b)
	return strings.TrimSpace(b.String())
}

// string implements Type.
func (t *structType) string(b *strings.Builder) {
	b.WriteString(t.Kind().String())
	b.WriteByte(' ')
	for _, v := range t.attr {
		panic("TODO")
		_ = v
	}
	if t.tag != 0 {
		b.WriteString(t.tag.String())
		b.WriteByte(' ')
	}
	b.WriteByte('{')
	for _, v := range t.fields {
		v.string(b)
		b.WriteString("; ")
	}
	b.WriteByte('}')
}

// FieldByIndex implements Type.
func (t *structType) FieldByIndex(i []int) Field {
	if len(i) > 1 {
		panic("TODO")
	}

	return t.fields[i[0]]
}

// FieldByName implements Type.
func (t *structType) FieldByName(name StringID) (Field, bool) {
	f, ok := t.m[name]
	return f, ok
}

func (t *structType) NumField() int { return len(t.fields) }

type taggedType struct {
	*typeBase
	resolutionScope Scope
	typ             Type

	tag StringID
}

// Alias implements Type.
func (t *taggedType) Alias() Type { return t }

// String implements Type.
func (t *taggedType) String() string {
	var b strings.Builder
	t.string(&b)
	return strings.TrimSpace(b.String())
}

// NumField implements Type.
func (t *taggedType) NumField() int { return t.underlyingType().NumField() }

// FieldByIndex implements Type.
func (t *taggedType) FieldByIndex(i []int) Field { return t.underlyingType().FieldByIndex(i) }

// FieldByName implements Type.
func (t *taggedType) FieldByName(s StringID) (Field, bool) { return t.underlyingType().FieldByName(s) }

// string implements Type.
func (t *taggedType) string(b *strings.Builder) {
	t.typeBase.string(b)
	b.WriteByte(' ')
	b.WriteString(t.tag.String())
}

func (t *taggedType) underlyingType() Type {
	if t.typ != nil {
		return t.typ
	}

	k := t.Kind()
	for s := t.resolutionScope; s != nil; s = s.Parent() {
		for _, v := range s[t.tag] {
			switch x := v.(type) {
			case *Declarator, *StructDeclarator:
			case *EnumSpecifier:
				if k == Enum {
					t.typ = x.Type()
					return t.typ
				}
			case *StructOrUnionSpecifier:
				switch k {
				case Struct:
					if typ := x.Type(); typ.Kind() == Struct {
						t.typ = typ
						return typ
					}
				case Union:
					if typ := x.Type(); typ.Kind() == Union {
						t.typ = typ
						return typ
					}
				}
			default:
				panic("internal error") //TODOOK
			}
		}
	}
	t.typ = noType
	return noType
}

// Size implements Type.
func (t *taggedType) Size() uintptr {
	return t.underlyingType().Size()
}

// Align implements Type.
func (t *taggedType) Align() int { return t.underlyingType().Align() }

// FieldAlign implements Type.
func (t *taggedType) FieldAlign() int { return t.underlyingType().FieldAlign() }

type functionType struct {
	typeBase
	params    []Object // *Declarator or *AbstractDeclarator
	paramList []StringID

	result Type

	variadic bool
}

// Alias implements Type.
func (t *functionType) Alias() Type { return t }

// Decay implements Type.
func (t *functionType) Decay() Type { return t }

// String implements Type.
func (t *functionType) String() string {
	var b strings.Builder
	t.string(&b)
	return strings.TrimSpace(b.String())
}

// string implements Type.
func (t *functionType) string(b *strings.Builder) {
	b.WriteString("function(")
	for i, v := range t.params {
		v.Type().string(b)
		if i < len(t.params)-1 {
			b.WriteString(", ")
		}
	}
	if t.variadic {
		b.WriteString(", ...")
	}
	b.WriteString(")")
	if t.result != nil && t.result.Kind() != Void {
		b.WriteString(" returning ")
		t.result.string(b)
	}
}

// Parameters implements Type.
func (t *functionType) Parameters() []Object { return t.params }

// Result implements Type.
func (t *functionType) Result() Type { return t.result }

// IsVariadic implements Type.
func (t *functionType) IsVariadic() bool { return t.variadic }

func mkPtr(ctx *context, n Node, t Type) *pointerType {
	base := t.base()
	base.align = byte(ctx.cfg.ABI.align(Ptr))
	base.fieldAlign = byte(ctx.cfg.ABI.fieldAlign(Ptr))
	base.kind = byte(Ptr)
	base.size = uintptr(ctx.cfg.ABI.size(Ptr))
	return &pointerType{
		elem:     t,
		typeBase: base,
	}
}

type bitFieldType struct {
	Type
}

// Alias implements Type.
func (t *bitFieldType) Alias() Type { return t }

// IsBitFieldType implements Type.
func (t *bitFieldType) IsBitFieldType() bool { return true } //TODO-
