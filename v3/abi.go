// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	//TODO- "fmt" //TODO-
	"math"
)

// Bit field allocation
//
//	#include <stdio.h>
//
//	struct {
//		int f0:2, f1:3, f2:20, f3:10, f4;
//	} x;
//
//	int main() {
//		long long unsigned *p = (void*)&x;
//		printf("%llx\n", *p);
//		x.f0 = -1;
//		printf("%llx\n", *p);
//		x.f1 = -1;
//		printf("%llx\n", *p);
//		x.f2 = -1;
//		printf("%llx\n", *p);
//		x.f3 = -1;
//		printf("%llx\n", *p);
//	}
//
// linux/amd64: from bit 0 upwards, MaxPackedBitfieldWidth 32.
//
//	$ ./a.out
//	0
//	3
//	1f
//	1ffffff
//	3ff01ffffff
//	$
//
// ----------------------------------------------------------------------------
//
//	#include <stdio.h>
//
//	union {
//		int f0:2, f1:3, f2:20, f3:10, f4;
//	} x;
//
//	int main() {
//		long long unsigned *p = (void*)&x;
//		printf("%llx\n", *p);
//		x.f0 = -1;
//		printf("%llx\n", *p);
//		x.f1 = -1;
//		printf("%llx\n", *p);
//		x.f2 = -1;
//		printf("%llx\n", *p);
//		x.f3 = -1;
//		printf("%llx\n", *p);
//	}
//
// linux/amd64: all fields, including bitfields start at offset 0, bit offset 0.
//
//	$ ./a.out
//	0
//	3
//	7
//	fffff
//	fffff
//	$

// ABIType describes properties of a non-aggregate type.
type ABIType struct {
	Size       uintptr
	Align      int
	FieldAlign int
}

// ABI describes selected parts of the Application Binary Interface.
type ABI struct {
	MaxPackedBitfieldWidth int // In bits.
	Types                  map[Kind]ABIType
	types                  map[Kind]Type

	SignedChar bool
}

func (a *ABI) sanityCheck(ctx *context, intMaxWidth int) error {
	if intMaxWidth == 0 {
		intMaxWidth = 64
	}
	if a.MaxPackedBitfieldWidth < 0 || a.MaxPackedBitfieldWidth > intMaxWidth {
		ctx.err(noPos, "invalid ABI.MaxPackedBitfieldWidth value: %v", a.MaxPackedBitfieldWidth)
	}
	if _, ok := a.Types[Void]; !ok {
		if ctx.err(noPos, "ABI is missing %s", Void) {
			return ctx.Err()
		}
	}

	a.types = map[Kind]Type{}
	for _, k := range []Kind{
		Bool,
		Char,
		ComplexFloat,
		ComplexLongDouble,
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

		if k != Void && (v.Size == 0 || v.Align == 0 || v.FieldAlign == 0 ||
			v.Align > math.MaxUint8 || v.FieldAlign > math.MaxUint8) {
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

func (a *ABI) Type(k Kind) Type      { return a.types[k] }
func (a *ABI) align(k Kind) int      { return a.Types[k].Align }      //TODO export, -n, -ctx
func (a *ABI) fieldAlign(k Kind) int { return a.Types[k].FieldAlign } //TODO export, -n, -ctx
func (a *ABI) size(k Kind) int       { return int(a.Types[k].Size) }

func (a *ABI) isSignedInteger(k Kind) bool {
	if !integerTypes[k] {
		panic("internal error") //TODOOK
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

	//TODO- defer func() { //TODO-
	//TODO- 	fmt.Printf("%s, %v\n", t, t.Size())
	//TODO- }()
	var off int64 // bit offset
	var align, fieldAlign int

	switch {
	case t.Kind() == Union:
		for i, f := range t.fields {
			ft := f.Type()
			if ft.Kind() == Array && ft.Incomplete() && i == len(t.fields)-1 {
				continue
				//TODO
			}

			sz := ft.Size()
			if sz == 0 {
				return t //TODO-
				panic("TODO")
			}
			al := ft.FieldAlign()
			if al == 0 {
				panic("TODO")
			}

			switch {
			case f.isBitField:
				return t //TODO-
				panic("TODO")
			default:
				if al > align {
					align = al
				}
				if fal := ft.FieldAlign(); fal > fieldAlign {
					fieldAlign = fal
				}
				if n := 8 * int64(sz); n > off {
					off = n
				}
			}
		}
		if align == 0 { //TODO ok?
			align = 1
		}
		t.align = byte(align)
		t.fieldAlign = byte(fieldAlign)
		off = roundup(off, int64(align))
		t.size = uintptr(off >> 3)
		ctx.structs[StructInfo{Size: t.size, Align: t.Align()}] = struct{}{}
	default:
		var i int
		var f *field
		for i, f = range t.fields {
			ft := f.Type()
			if ft.Kind() == Array && ft.Incomplete() && i == len(t.fields)-1 {
				continue
				//TODO
			}

			sz := ft.Size()
			if sz == 0 {
				return t //TODO-
				panic("TODO")
			}
			al := ft.FieldAlign()
			if al == 0 {
				return t //TODO-
				panic(n.Position().String())
			}

			switch {
			case f.isBitField:
				return t //TODO-
				panic("TODO")
			default:
				if al > align {
					align = al
				}
				if fal := ft.FieldAlign(); fal > fieldAlign {
					fieldAlign = fal
				}
				off0 := off
				off = roundup(off, 8*int64(al))
				f.pad = byte(off-off0) >> 3
				f.offset = uintptr(off) >> 3
				off += 8 * int64(sz)
			}
		}
		if align == 0 { //TODO ok?
			align = 1
		}
		t.align = byte(align)
		t.fieldAlign = byte(fieldAlign)
		off0 := off
		off = roundup(off, int64(align))
		if f != nil {
			f.pad = byte(off-off0) >> 3
		}
		t.size = uintptr(off >> 3)
		ctx.structs[StructInfo{Size: t.size, Align: t.Align()}] = struct{}{}
	}
	//dbg("%v sz %v", t, t.Size())
	return t
}
