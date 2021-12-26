// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

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
	//TODO
}
