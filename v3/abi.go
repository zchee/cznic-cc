// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

package cc // import "modernc.org/cc/v3"

import (
	"math"
)

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

	for _, k := range []Kind{
		Bool,
		Char,
		ComplexFloat,
		ComplexLongDouble,
		Double,
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
	} {
		v, ok := a.Types[k]
		if !ok {
			if ctx.err(noPos, "ABI is missing %s", k) {
				return ctx.Err()
			}

			continue
		}

		if v.Size == 0 || v.Align == 0 || v.FieldAlign == 0 ||
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
	}
	return ctx.Err()
}

func (a *ABI) align(ctx *context, n Node, k Kind) int {
	x, ok := a.Types[k]
	if !ok {
		ctx.errNode(n, "ABI is missing an ABIType for kind %s", k)
		return 1
	}

	return x.Align
}

func (a *ABI) fieldAlign(ctx *context, n Node, k Kind) int {
	x, ok := a.Types[k]
	if !ok {
		ctx.errNode(n, "ABI is missing an ABIType for kind %s", k)
		return 1
	}

	return x.FieldAlign
}

func (a *ABI) isSigned(k Kind) bool {
	if !a.isInt(k) {
		panic("internal error") //TODOOK
	}

	return !unsignedTypes[k] || k == Char && a.SignedChar
}

func (a *ABI) isInt(k Kind) bool { return integerTypes[k] }

func (a *ABI) size(k Kind) int {
	if x, ok := a.Types[k]; ok {
		return int(x.Size)
	}

	panic("internal error") //TODOOK
}

func (a *ABI) typ(ctx *context, n Node, k Kind) Type { //TODO singletons within ABI instance
	return &typeBase{
		align:      byte(a.align(ctx, n, k)),
		fieldAlign: byte(a.fieldAlign(ctx, n, k)),
		kind:       byte(k),
		size:       uintptr(a.size(k)),
	}
}

func (a *ABI) layout(ctx *context, t *structType) *structType {
	if t == nil {
		return nil
	}

	//TODO
	return t
}
