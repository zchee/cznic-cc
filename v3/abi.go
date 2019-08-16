// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"encoding/binary"
	"math"
)

var (
	complexTypedefs = map[StringID]Kind{
		dict.sid("__COMPLEX_CHAR_TYPE__"):               ComplexChar,
		dict.sid("__COMPLEX_DOUBLE_TYPE__"):             ComplexDouble,
		dict.sid("__COMPLEX_FLOAT_TYPE__"):              ComplexFloat,
		dict.sid("__COMPLEX_INT_TYPE__"):                ComplexInt,
		dict.sid("__COMPLEX_LONG_TYPE__"):               ComplexLong,
		dict.sid("__COMPLEX_LONG_DOUBLE_TYPE__"):        ComplexLongDouble,
		dict.sid("__COMPLEX_LONG_LONG_TYPE__"):          ComplexLongLong,
		dict.sid("__COMPLEX_SHORT_TYPE__"):              ComplexShort,
		dict.sid("__COMPLEX_UNSIGNED_TYPE__"):           ComplexUInt,
		dict.sid("__COMPLEX_LONG_UNSIGNED_TYPE__"):      ComplexULong,
		dict.sid("__COMPLEX_LONG_LONG_UNSIGNED_TYPE__"): ComplexULongLong,
		dict.sid("__COMPLEX_SHORT_UNSIGNED_TYPE__"):     ComplexUShort,
	}
)

// ABIType describes properties of a non-aggregate type.
type ABIType struct {
	Size       uintptr
	Align      int
	FieldAlign int
}

// ABI describes selected parts of the Application Binary Interface.
type ABI struct {
	ByteOrder binary.ByteOrder
	Types     map[Kind]ABIType
	types     map[Kind]Type

	SignedChar bool
}

func (a *ABI) sanityCheck(ctx *context, intMaxWidth int, s Scope) error {
	if intMaxWidth == 0 {
		intMaxWidth = 64
	}

	a.types = map[Kind]Type{}
	for _, k := range []Kind{
		Bool,
		Char,
		Double,
		Enum,
		Float,
		Int,
		Long,
		LongDouble,
		LongLong,
		Ptr,
		SChar,
		Short,
		UChar,
		UInt,
		ULong,
		ULongLong,
		UShort,
		Void,
	} {
		v, ok := a.Types[k]
		if !ok {
			if ctx.err(noPos, "ABI is missing %s", k) {
				return ctx.Err()
			}

			continue
		}

		if (k != Void && v.Size == 0) || v.Align == 0 || v.FieldAlign == 0 ||
			v.Align > math.MaxUint8 || v.FieldAlign > math.MaxUint8 {
			if ctx.err(noPos, "invalid ABI type %s: %+v", k, v) {
				return ctx.Err()
			}
		}

		if integerTypes[k] && v.Size > 8 {
			if ctx.err(noPos, "invalid ABI type %s size: %v, must be <= 8", k, v.Size) {
				return ctx.Err()
			}
		}
		var f flag
		if integerTypes[k] && a.isSignedInteger(k) {
			f = fSigned
		}
		t := &typeBase{
			align:      byte(a.align(k)),
			fieldAlign: byte(a.fieldAlign(k)),
			flags:      f,
			kind:       byte(k),
			size:       uintptr(a.size(k)),
		}
		a.types[k] = t
	}
	return ctx.Err()
}

func (a *ABI) Type(k Kind) Type { return a.types[k] }

func (a *ABI) align(k Kind) int      { return a.Types[k].Align }
func (a *ABI) fieldAlign(k Kind) int { return a.Types[k].FieldAlign }
func (a *ABI) size(k Kind) int       { return int(a.Types[k].Size) }

func (a *ABI) isSignedInteger(k Kind) bool {
	if !integerTypes[k] {
		internalError()
	}

	switch k {
	case Bool, UChar, UInt, ULong, ULongLong, UShort:
		return false
	case Char:
		return a.SignedChar
	default:
		return true
	}
}

func roundup(n, to int64) int64 {
	if r := n % to; r != 0 {
		return n + to - r
	}

	return n
}

func (a *ABI) layout(ctx *context, n Node, t *structType) *structType {
	if t == nil {
		return nil
	}

	var off int64 // bit offset
	align := 1

	switch {
	case t.Kind() == Union:
		for _, f := range t.fields {
			ft := f.Type()
			sz := ft.Size()
			if n := int64(8 * sz); n > off {
				off = n
			}
			al := ft.FieldAlign()
			if al == 0 {
				al = 1
			}
			if al > align {
				align = al
			}

			if f.isBitField {
				f.bitFieldMask = 1<<f.bitFieldWidth - 1
			}
			f.promote = integerPromotion(ctx, ft)
		}
		t.align = byte(align)
		t.fieldAlign = byte(align)
		off = roundup(off, int64(align))
		t.size = uintptr(off >> 3)
		ctx.structs[StructInfo{Size: t.size, Align: t.Align()}] = struct{}{}
	default:
		var i int
		var f *field
		lzero := false
		for i, f = range t.fields {
			ft := f.Type()
			var sz uintptr
			switch {
			case ft.Kind() == Array && i == len(t.fields)-1:
				if ft.IsIncomplete() || ft.Len() == 0 {
					break
				}

				fallthrough
			default:
				sz = ft.Size()
			}

			bitSize := 8 * int(sz)
			al := ft.FieldAlign()
			if al == 0 {
				al = 1
			}
			if al > align {
				align = al
			}

			switch {
			case f.isBitField:
				down := off &^ (8*int64(al) - 1)
				bitoff := off - down
				downMax := off &^ (int64(bitSize) - 1)
				switch {
				case lzero || int(off-downMax)+int(f.bitFieldWidth) > bitSize:
					off = roundup(off, 8*int64(al))
					f.offset = uintptr(off >> 3)
					f.bitFieldOffset = 0
					f.bitFieldMask = 1<<f.bitFieldWidth - 1
					off += int64(f.bitFieldWidth)
					if f.bitFieldWidth == 0 {
						continue
					}
				default:
					f.offset = uintptr(down >> 3)
					f.bitFieldOffset = byte(bitoff)
					f.bitFieldMask = (1<<f.bitFieldWidth - 1) << byte(bitoff)
					off += int64(f.bitFieldWidth)
				}
			default:
				off0 := off
				off = roundup(off, 8*int64(al))
				f.pad = byte(off-off0) >> 3
				f.offset = uintptr(off) >> 3
				off += 8 * int64(sz)
			}
			f.promote = integerPromotion(ctx, ft)
			lzero = false
		}
		t.align = byte(align)
		t.fieldAlign = byte(align)
		off0 := off
		off = roundup(off, 8*int64(align))
		if f != nil && !f.IsBitField() {
			f.pad = byte(off-off0) >> 3
		}
		t.size = uintptr(off >> 3)
		ctx.structs[StructInfo{Size: t.size, Align: t.Align()}] = struct{}{}
	}
	return t
}

func (a *ABI) Ptr(n Node, t Type) Type {
	base := t.base()
	base.align = byte(a.align(Ptr))
	base.fieldAlign = byte(a.fieldAlign(Ptr))
	base.kind = byte(Ptr)
	base.size = uintptr(a.size(Ptr))
	base.flags &^= fIncomplete
	return &pointerType{
		elem:     t,
		typeBase: base,
	}
}
