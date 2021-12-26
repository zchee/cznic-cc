// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate rm -f ast.go
//go:generate yy -o /dev/null -position -astImport "\"fmt\"\n\n\"modernc.org/token\"" -prettyString PrettyString -kind Case -noListKind -noPrivateHelpers -forceOptPos parser.yy
//go:generate stringer -output stringer.go -type=tokCh
//go:generate sh -c "go test -run ^Example |fe"

// Package cc is a C99 compiler front end.
//
// Online documentation
//
// See https://godoc.org/modernc.org/cc/v3.
//
// Links
//
// Referenced from elsewhere:
//
//  [0]: http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf
package cc // import "modernc.org/cc/v4"
