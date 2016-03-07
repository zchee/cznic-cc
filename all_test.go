// Copyright 2016 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/scanner"
	"go/token"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
	"unicode"

	"github.com/cznic/golex/lex"
	"github.com/cznic/xc"
)

func printStack() { debug.PrintStack() }

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "\tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func TODO(...interface{}) string {
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("TODO: %s:%d:\n", path.Base(fn), fl)
}

func use(...interface{}) {}

// ============================================================================

var (
	o1        = flag.String("1", "", "Single file argument of TestPPParse1.")
	oAST      = flag.Bool("ast", false, "Show AST.")
	oFailFast = flag.Bool("ff", false, "crash on first reported error (in some tests.)")
	oLong     = flag.Bool("long", false, "Enable long tests. (false)")
	oRE       = flag.String("re", "", "Regexp filter.")

	includes = []string{}

	predefinedMacros = `
#define __STDC_HOSTED__ 1
#define __STDC_VERSION__ 199901L
#define __STDC__ 1

#define __MODEL64

void __GO__(char *s, ...);
`

	gccPredefined = `
#define __SSP_STRONG__ 3
#define __DBL_MIN_EXP__ (-1021)
#define __UINT_LEAST16_MAX__ 0xffff
#define __ATOMIC_ACQUIRE 2
#define __FLT_MIN__ 1.17549435082228750797e-38F
#define __GCC_IEC_559_COMPLEX 2
#define __UINT_LEAST8_TYPE__ unsigned char
#define __SIZEOF_FLOAT80__ 16
#define __INTMAX_C(c) c ## L
#define __CHAR_BIT__ 8
#define __UINT8_MAX__ 0xff
#define __WINT_MAX__ 0xffffffffU
#define __ORDER_LITTLE_ENDIAN__ 1234
#define __SIZE_MAX__ 0xffffffffffffffffUL
#define __WCHAR_MAX__ 0x7fffffff
#define __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 1
#define __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 1
#define __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 1
#define __DBL_DENORM_MIN__ ((double)4.94065645841246544177e-324L)
#define __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 1
#define __GCC_ATOMIC_CHAR_LOCK_FREE 2
#define __GCC_IEC_559 2
#define __FLT_EVAL_METHOD__ 0
#define __unix__ 1
#define __GCC_ATOMIC_CHAR32_T_LOCK_FREE 2
#define __x86_64 1
#define __UINT_FAST64_MAX__ 0xffffffffffffffffUL
#define __SIG_ATOMIC_TYPE__ int
#define __DBL_MIN_10_EXP__ (-307)
#define __FINITE_MATH_ONLY__ 0
#define __GNUC_PATCHLEVEL__ 1
#define __UINT_FAST8_MAX__ 0xff
#define __has_include(STR) __has_include__(STR)
#define __DEC64_MAX_EXP__ 385
#define __INT8_C(c) c
#define __UINT_LEAST64_MAX__ 0xffffffffffffffffUL
#define __SHRT_MAX__ 0x7fff
#define __LDBL_MAX__ 1.18973149535723176502e+4932L
#define __UINT_LEAST8_MAX__ 0xff
#define __GCC_ATOMIC_BOOL_LOCK_FREE 2
#define __UINTMAX_TYPE__ long unsigned int
#define __linux 1
#define __DEC32_EPSILON__ 1E-6DF
#define __unix 1
#define __UINT32_MAX__ 0xffffffffU
#define __LDBL_MAX_EXP__ 16384
#define __WINT_MIN__ 0U
#define __linux__ 1
#define __SCHAR_MAX__ 0x7f
#define __WCHAR_MIN__ (-__WCHAR_MAX__ - 1)
#define __INT64_C(c) c ## L
#define __DBL_DIG__ 15
#define __GCC_ATOMIC_POINTER_LOCK_FREE 2
#define __SIZEOF_INT__ 4
#define __SIZEOF_POINTER__ 8
#define __USER_LABEL_PREFIX__ 
#define __STDC_HOSTED__ 1
#define __LDBL_HAS_INFINITY__ 1
#define __FLT_EPSILON__ 1.19209289550781250000e-7F
#define __LDBL_MIN__ 3.36210314311209350626e-4932L
#define __STDC_UTF_16__ 1
#define __DEC32_MAX__ 9.999999E96DF
#define __INT32_MAX__ 0x7fffffff
#define __SIZEOF_LONG__ 8
#define __STDC_IEC_559__ 1
#define __STDC_ISO_10646__ 201103L
#define __UINT16_C(c) c
#define __DECIMAL_DIG__ 21
#define __gnu_linux__ 1
#define __has_include_next(STR) __has_include_next__(STR)
#define __LDBL_HAS_QUIET_NAN__ 1
#define __GNUC__ 5
#define __MMX__ 1
#define __FLT_HAS_DENORM__ 1
#define __SIZEOF_LONG_DOUBLE__ 16
#define __BIGGEST_ALIGNMENT__ 16
#define __DBL_MAX__ ((double)1.79769313486231570815e+308L)
#define __INT_FAST32_MAX__ 0x7fffffffffffffffL
#define __DBL_HAS_INFINITY__ 1
#define __DEC32_MIN_EXP__ (-94)
#define __INT_FAST16_TYPE__ long int
#define __LDBL_HAS_DENORM__ 1
#define __DEC128_MAX__ 9.999999999999999999999999999999999E6144DL
#define __INT_LEAST32_MAX__ 0x7fffffff
#define __DEC32_MIN__ 1E-95DF
#define __DBL_MAX_EXP__ 1024
#define __DEC128_EPSILON__ 1E-33DL
#define __SSE2_MATH__ 1
#define __ATOMIC_HLE_RELEASE 131072
#define __PTRDIFF_MAX__ 0x7fffffffffffffffL
#define __amd64 1
#define __STDC_NO_THREADS__ 1
#define __ATOMIC_HLE_ACQUIRE 65536
#define __LONG_LONG_MAX__ 0x7fffffffffffffffLL
#define __SIZEOF_SIZE_T__ 8
#define __SIZEOF_WINT_T__ 4
#define __GCC_HAVE_DWARF2_CFI_ASM 1
#define __GXX_ABI_VERSION 1009
#define __FLT_MIN_EXP__ (-125)
#define __INT_FAST64_TYPE__ long int
#define __DBL_MIN__ ((double)2.22507385850720138309e-308L)
#define __LP64__ 1
#define __DECIMAL_BID_FORMAT__ 1
#define __DEC128_MIN__ 1E-6143DL
#define __REGISTER_PREFIX__ 
#define __UINT16_MAX__ 0xffff
#define __DBL_HAS_DENORM__ 1
#define __UINT8_TYPE__ unsigned char
#define __NO_INLINE__ 1
#define __FLT_MANT_DIG__ 24
#define __VERSION__ "5.2.1 20151010"
#define __UINT64_C(c) c ## UL
#define _STDC_PREDEF_H 1
#define __GCC_ATOMIC_INT_LOCK_FREE 2
#define __FLOAT_WORD_ORDER__ __ORDER_LITTLE_ENDIAN__
#define __STDC_IEC_559_COMPLEX__ 1
#define __INT32_C(c) c
#define __DEC64_EPSILON__ 1E-15DD
#define __ORDER_PDP_ENDIAN__ 3412
#define __DEC128_MIN_EXP__ (-6142)
#define __INT_FAST32_TYPE__ long int
#define __UINT_LEAST16_TYPE__ short unsigned int
#define unix 1
#define __INT16_MAX__ 0x7fff
#define __SIZE_TYPE__ long unsigned int
#define __UINT64_MAX__ 0xffffffffffffffffUL
#define __INT8_TYPE__ signed char
#define __ELF__ 1
#define __FLT_RADIX__ 2
#define __INT_LEAST16_TYPE__ short int
#define __LDBL_EPSILON__ 1.08420217248550443401e-19L
#define __UINTMAX_C(c) c ## UL
#define __SSE_MATH__ 1
#define __k8 1
#define __SIG_ATOMIC_MAX__ 0x7fffffff
#define __GCC_ATOMIC_WCHAR_T_LOCK_FREE 2
#define __SIZEOF_PTRDIFF_T__ 8
#define __x86_64__ 1
#define __DEC32_SUBNORMAL_MIN__ 0.000001E-95DF
#define __INT_FAST16_MAX__ 0x7fffffffffffffffL
#define __UINT_FAST32_MAX__ 0xffffffffffffffffUL
#define __UINT_LEAST64_TYPE__ long unsigned int
#define __FLT_HAS_QUIET_NAN__ 1
#define __FLT_MAX_10_EXP__ 38
#define __LONG_MAX__ 0x7fffffffffffffffL
#define __DEC128_SUBNORMAL_MIN__ 0.000000000000000000000000000000001E-6143DL
#define __FLT_HAS_INFINITY__ 1
#define __UINT_FAST16_TYPE__ long unsigned int
#define __DEC64_MAX__ 9.999999999999999E384DD
#define __CHAR16_TYPE__ short unsigned int
#define __PRAGMA_REDEFINE_EXTNAME 1
#define __INT_LEAST16_MAX__ 0x7fff
#define __DEC64_MANT_DIG__ 16
#define __INT64_MAX__ 0x7fffffffffffffffL
#define __UINT_LEAST32_MAX__ 0xffffffffU
#define __GCC_ATOMIC_LONG_LOCK_FREE 2
#define __INT_LEAST64_TYPE__ long int
#define __INT16_TYPE__ short int
#define __INT_LEAST8_TYPE__ signed char
#define __STDC_VERSION__ 201112L
#define __DEC32_MAX_EXP__ 97
#define __INT_FAST8_MAX__ 0x7f
#define __INTPTR_MAX__ 0x7fffffffffffffffL
#define linux 1
#define __SSE2__ 1
#define __LDBL_MANT_DIG__ 64
#define __DBL_HAS_QUIET_NAN__ 1
#define __SIG_ATOMIC_MIN__ (-__SIG_ATOMIC_MAX__ - 1)
#define __code_model_small__ 1
#define __k8__ 1
#define __INTPTR_TYPE__ long int
#define __UINT16_TYPE__ short unsigned int
#define __WCHAR_TYPE__ int
#define __SIZEOF_FLOAT__ 4
#define __UINTPTR_MAX__ 0xffffffffffffffffUL
#define __DEC64_MIN_EXP__ (-382)
#define __INT_FAST64_MAX__ 0x7fffffffffffffffL
#define __GCC_ATOMIC_TEST_AND_SET_TRUEVAL 1
#define __FLT_DIG__ 6
#define __UINT_FAST64_TYPE__ long unsigned int
#define __INT_MAX__ 0x7fffffff
#define __amd64__ 1
#define __INT64_TYPE__ long int
#define __FLT_MAX_EXP__ 128
#define __ORDER_BIG_ENDIAN__ 4321
#define __DBL_MANT_DIG__ 53
#define __SIZEOF_FLOAT128__ 16
#define __INT_LEAST64_MAX__ 0x7fffffffffffffffL
#define __GCC_ATOMIC_CHAR16_T_LOCK_FREE 2
#define __DEC64_MIN__ 1E-383DD
#define __WINT_TYPE__ unsigned int
#define __UINT_LEAST32_TYPE__ unsigned int
#define __SIZEOF_SHORT__ 2
#define __SSE__ 1
#define __LDBL_MIN_EXP__ (-16381)
#define __INT_LEAST8_MAX__ 0x7f
#define __SIZEOF_INT128__ 16
#define __LDBL_MAX_10_EXP__ 4932
#define __ATOMIC_RELAXED 0
#define __DBL_EPSILON__ ((double)2.22044604925031308085e-16L)
#define _LP64 1
#define __UINT8_C(c) c
#define __INT_LEAST32_TYPE__ int
#define __SIZEOF_WCHAR_T__ 4
#define __UINT64_TYPE__ long unsigned int
#define __INT_FAST8_TYPE__ signed char
#define __GNUC_STDC_INLINE__ 1
#define __DBL_DECIMAL_DIG__ 17
#define __STDC_UTF_32__ 1
#define __FXSR__ 1
#define __DEC_EVAL_METHOD__ 2
#define __UINT32_C(c) c ## U
#define __INTMAX_MAX__ 0x7fffffffffffffffL
#define __BYTE_ORDER__ __ORDER_LITTLE_ENDIAN__
#define __FLT_DENORM_MIN__ 1.40129846432481707092e-45F
#define __INT8_MAX__ 0x7f
#define __UINT_FAST32_TYPE__ long unsigned int
#define __CHAR32_TYPE__ unsigned int
#define __FLT_MAX__ 3.40282346638528859812e+38F
#define __INT32_TYPE__ int
#define __SIZEOF_DOUBLE__ 8
#define __FLT_MIN_10_EXP__ (-37)
#define __INTMAX_TYPE__ long int
#define __DEC128_MAX_EXP__ 6145
#define __ATOMIC_CONSUME 1
#define __GNUC_MINOR__ 2
#define __UINTMAX_MAX__ 0xffffffffffffffffUL
#define __DEC32_MANT_DIG__ 7
#define __DBL_MAX_10_EXP__ 308
#define __LDBL_DENORM_MIN__ 3.64519953188247460253e-4951L
#define __INT16_C(c) c
#define __STDC__ 1
#define __PTRDIFF_TYPE__ long int
#define __ATOMIC_SEQ_CST 5
#define __UINT32_TYPE__ unsigned int
#define __UINTPTR_TYPE__ long unsigned int
#define __DEC64_SUBNORMAL_MIN__ 0.000000000000001E-383DD
#define __DEC128_MANT_DIG__ 34
#define __LDBL_MIN_10_EXP__ (-4931)
#define __SIZEOF_LONG_LONG__ 8
#define __GCC_ATOMIC_LLONG_LOCK_FREE 2
#define __LDBL_DIG__ 18
#define __FLT_DECIMAL_DIG__ 9
#define __UINT_FAST16_MAX__ 0xffffffffffffffffUL
#define __GCC_ATOMIC_SHORT_LOCK_FREE 2
#define __UINT_FAST8_TYPE__ unsigned char
#define __ATOMIC_ACQ_REL 4
#define __ATOMIC_RELEASE 3
`
	sysIncludes = []string{}

	testTweaks = &tweaks{
		enableDefineOmitCommaBeforeDDD: true,
		enableDlrInIdentifiers:         true,
		enableEmptyDefine:              true,
		enableUndefExtraTokens:         true,
	}
)

func newTestReport() *xc.Report {
	r := xc.NewReport()
	r.ErrLimit = -1
	if *oFailFast {
		r.PanicOnError = true
	}
	return r
}

func init() {
	isTesting = true
	log.SetFlags(log.Llongfile)
	flag.IntVar(&yyDebug, "yydebug", 0, "")
	flag.BoolVar(&isGenerating, "generating", false, "go generate is executing (false).")
	flag.BoolVar(&debugTypeStrings, "xtypes", false, "add debug info to type strings")
}

func newTestModel() *Model {
	return &Model{ // 64
		Items: map[Kind]ModelItem{
			Ptr:               {8, 8, 8, nil},
			UintPtr:           {8, 8, 8, nil},
			Void:              {0, 1, 1, nil},
			Char:              {1, 1, 1, nil},
			SChar:             {1, 1, 1, nil},
			UChar:             {1, 1, 1, nil},
			Short:             {2, 2, 2, nil},
			UShort:            {2, 2, 2, nil},
			Int:               {4, 4, 4, nil},
			UInt:              {4, 4, 4, nil},
			Long:              {8, 8, 8, nil},
			ULong:             {8, 8, 8, nil},
			LongLong:          {8, 8, 8, nil},
			ULongLong:         {8, 8, 8, nil},
			Float:             {4, 4, 4, nil},
			Double:            {8, 8, 8, nil},
			LongDouble:        {8, 8, 8, nil},
			Bool:              {1, 1, 1, nil},
			FloatComplex:      {8, 8, 8, nil},
			DoubleComplex:     {16, 16, 16, nil},
			LongDoubleComplex: {16, 16, 16, nil},
		},
	}
}

func printError(w io.Writer, pref string, err error) {
	switch x := err.(type) {
	case scanner.ErrorList:
		for _, v := range x {
			fmt.Fprintf(w, "%s%v\n", pref, v)
		}
	default:
		fmt.Fprintf(w, "%s%v\n", pref, err)
	}
}

func errString(err error) string {
	var b bytes.Buffer
	printError(&b, "", err)
	return b.String()
}

func testUCNTable(t *testing.T, tab []rune, fOk, fOther func(rune) bool, fcategory func(rune) bool, tag string) {
	m := map[rune]struct{}{}
	for i := 0; i < len(tab); i += 2 {
		l, h := tab[i], tab[i+1]
		if h == 0 {
			h = l
		}
		for r := rune(l); r <= rune(h); r++ {
			m[r] = struct{}{}
		}
	}
	for r := rune(0); r < 0xffff; r++ {
		_, ok := m[r]
		if g, e := fOk(r), ok; g != e {
			t.Errorf("%#04x %v %v", r, g, e)
		}

		if ok {
			if g, e := fOther(r), false; g != e {
				t.Errorf("%#04x %v %v", r, g, e)
			}
		}
	}
}

func TestUCNDigitsTable(t *testing.T) {
	tab := []rune{
		0x0660, 0x0669, 0x06F0, 0x06F9, 0x0966, 0x096F, 0x09E6, 0x09EF, 0x0A66, 0x0A6F,
		0x0AE6, 0x0AEF, 0x0B66, 0x0B6F, 0x0BE7, 0x0BEF, 0x0C66, 0x0C6F, 0x0CE6, 0x0CEF,
		0x0D66, 0x0D6F, 0x0E50, 0x0E59, 0x0ED0, 0x0ED9, 0x0F20, 0x0F33,
	}
	testUCNTable(t, tab, isUCNDigit, isUCNNonDigit, unicode.IsDigit, "unicode.IsDigit")
}

func TestUCNNonDigitsTable(t *testing.T) {
	tab := []rune{
		0x00AA, 0x0000, 0x00B5, 0x0000, 0x00B7, 0x0000, 0x00BA, 0x0000, 0x00C0, 0x00D6,
		0x00D8, 0x00F6, 0x00F8, 0x01F5, 0x01FA, 0x0217, 0x0250, 0x02A8, 0x02B0, 0x02B8,
		0x02BB, 0x0000, 0x02BD, 0x02C1, 0x02D0, 0x02D1, 0x02E0, 0x02E4, 0x037A, 0x0000,
		0x0386, 0x0000, 0x0388, 0x038A, 0x038C, 0x0000, 0x038E, 0x03A1, 0x03A3, 0x03CE,
		0x03D0, 0x03D6, 0x03DA, 0x0000, 0x03DC, 0x0000, 0x03DE, 0x0000, 0x03E0, 0x0000,
		0x03E2, 0x03F3, 0x0401, 0x040C, 0x040E, 0x044F, 0x0451, 0x045C, 0x045E, 0x0481,
		0x0490, 0x04C4, 0x04C7, 0x04C8, 0x04CB, 0x04CC, 0x04D0, 0x04EB, 0x04EE, 0x04F5,
		0x04F8, 0x04F9, 0x0531, 0x0556, 0x0559, 0x0000, 0x0561, 0x0587, 0x05B0, 0x05B9,
		0x05F0, 0x05F2, 0x0621, 0x063A, 0x0640, 0x0652, 0x0670, 0x06B7, 0x06BA, 0x06BE,
		0x06C0, 0x06CE, 0x06D0, 0x06DC, 0x06E5, 0x06E8, 0x06EA, 0x06ED, 0x0901, 0x0903,
		0x0905, 0x0939, 0x093D, 0x0000, 0x093E, 0x094D, 0x0950, 0x0952, 0x0958, 0x0963,
		0x0981, 0x0983, 0x0985, 0x098C, 0x098F, 0x0990, 0x0993, 0x09A8, 0x09AA, 0x09B0,
		0x09B2, 0x0000, 0x09B6, 0x09B9, 0x09BE, 0x09C4, 0x09C7, 0x09C8, 0x09CB, 0x09CD,
		0x09DC, 0x09DD, 0x09DF, 0x09E3, 0x09F0, 0x09F1, 0x0A02, 0x0000, 0x0A05, 0x0A0A,
		0x0A0F, 0x0A10, 0x0A13, 0x0A28, 0x0A2A, 0x0A30, 0x0A32, 0x0A33, 0x0A35, 0x0A36,
		0x0A38, 0x0A39, 0x0A3E, 0x0A42, 0x0A47, 0x0A48, 0x0A4B, 0x0A4D, 0x0A59, 0x0A5C,
		0x0A5E, 0x0000, 0x0A74, 0x0000, 0x0A81, 0x0A83, 0x0A85, 0x0A8B, 0x0A8D, 0x0000,
		0x0A8F, 0x0A91, 0x0A93, 0x0AA8, 0x0AAA, 0x0AB0, 0x0AB2, 0x0AB3, 0x0AB5, 0x0AB9,
		0x0ABD, 0x0AC5, 0x0AC7, 0x0AC9, 0x0ACB, 0x0ACD, 0x0AD0, 0x0000, 0x0AE0, 0x0000,
		0x0B01, 0x0B03, 0x0B05, 0x0B0C, 0x0B0F, 0x0B10, 0x0B13, 0x0B28, 0x0B2A, 0x0B30,
		0x0B32, 0x0B33, 0x0B36, 0x0B39, 0x0B3D, 0x0000, 0x0B3E, 0x0B43, 0x0B47, 0x0B48,
		0x0B4B, 0x0B4D, 0x0B5C, 0x0B5D, 0x0B5F, 0x0B61, 0x0B82, 0x0B83, 0x0B85, 0x0B8A,
		0x0B8E, 0x0B90, 0x0B92, 0x0B95, 0x0B99, 0x0B9A, 0x0B9C, 0x0000, 0x0B9E, 0x0B9F,
		0x0BA3, 0x0BA4, 0x0BA8, 0x0BAA, 0x0BAE, 0x0BB5, 0x0BB7, 0x0BB9, 0x0BBE, 0x0BC2,
		0x0BC6, 0x0BC8, 0x0BCA, 0x0BCD, 0x0C01, 0x0C03, 0x0C05, 0x0C0C, 0x0C0E, 0x0C10,
		0x0C12, 0x0C28, 0x0C2A, 0x0C33, 0x0C35, 0x0C39, 0x0C3E, 0x0C44, 0x0C46, 0x0C48,
		0x0C4A, 0x0C4D, 0x0C60, 0x0C61, 0x0C82, 0x0C83, 0x0C85, 0x0C8C, 0x0C8E, 0x0C90,
		0x0C92, 0x0CA8, 0x0CAA, 0x0CB3, 0x0CB5, 0x0CB9, 0x0CBE, 0x0CC4, 0x0CC6, 0x0CC8,
		0x0CCA, 0x0CCD, 0x0CDE, 0x0000, 0x0CE0, 0x0CE1, 0x0D02, 0x0D03, 0x0D05, 0x0D0C,
		0x0D0E, 0x0D10, 0x0D12, 0x0D28, 0x0D2A, 0x0D39, 0x0D3E, 0x0D43, 0x0D46, 0x0D48,
		0x0D4A, 0x0D4D, 0x0D60, 0x0D61, 0x0E01, 0x0E3A,

		// In [0], Annex D, Thai [0x0E40, 0x0E5B] overlaps with digits
		// [0x0E50, 0x0E59]. Exclude them.
		0x0E40, 0x0E4F,
		0x0E5A, 0x0E5B,

		0x0E81, 0x0E82,
		0x0E84, 0x0000, 0x0E87, 0x0E88, 0x0E8A, 0x0000, 0x0E8D, 0x0000, 0x0E94, 0x0E97,
		0x0E99, 0x0E9F, 0x0EA1, 0x0EA3, 0x0EA5, 0x0000, 0x0EA7, 0x0000, 0x0EAA, 0x0EAB,
		0x0EAD, 0x0EAE, 0x0EB0, 0x0EB9, 0x0EBB, 0x0EBD, 0x0EC0, 0x0EC4, 0x0EC6, 0x0000,
		0x0EC8, 0x0ECD, 0x0EDC, 0x0EDD, 0x0F00, 0x0000, 0x0F18, 0x0F19, 0x0F35, 0x0000,
		0x0F37, 0x0000, 0x0F39, 0x0000, 0x0F3E, 0x0F47, 0x0F49, 0x0F69, 0x0F71, 0x0F84,
		0x0F86, 0x0F8B, 0x0F90, 0x0F95, 0x0F97, 0x0000, 0x0F99, 0x0FAD, 0x0FB1, 0x0FB7,
		0x0FB9, 0x0000, 0x10A0, 0x10C5, 0x10D0, 0x10F6, 0x1E00, 0x1E9B, 0x1EA0, 0x1EF9,
		0x1F00, 0x1F15, 0x1F18, 0x1F1D, 0x1F20, 0x1F45, 0x1F48, 0x1F4D, 0x1F50, 0x1F57,
		0x1F59, 0x0000, 0x1F5B, 0x0000, 0x1F5D, 0x0000, 0x1F5F, 0x1F7D, 0x1F80, 0x1FB4,
		0x1FB6, 0x1FBC, 0x1FBE, 0x0000, 0x1FC2, 0x1FC4, 0x1FC6, 0x1FCC, 0x1FD0, 0x1FD3,
		0x1FD6, 0x1FDB, 0x1FE0, 0x1FEC, 0x1FF2, 0x1FF4, 0x1FF6, 0x1FFC, 0x203F, 0x2040,
		0x207F, 0x0000, 0x2102, 0x0000, 0x2107, 0x0000, 0x210A, 0x2113, 0x2115, 0x0000,
		0x2118, 0x211D, 0x2124, 0x0000, 0x2126, 0x0000, 0x2128, 0x0000, 0x212A, 0x2131,
		0x2133, 0x2138, 0x2160, 0x2182, 0x3005, 0x3007, 0x3021, 0x3029, 0x3041, 0x3093,
		0x309B, 0x309C, 0x30A1, 0x30F6, 0x30FB, 0x30FC, 0x3105, 0x312C, 0x4E00, 0x9FA5,
		0xAC00, 0xD7A3,
	}
	testUCNTable(t, tab, isUCNNonDigit, isUCNDigit, unicode.IsLetter, "unicode.IsLetter")
}

func charStr(c rune) string {
	return yySymName(int(c))
}

func charsStr(chars []lex.Char, delta token.Pos) (a []string) {
	for _, v := range chars {
		a = append(a, fmt.Sprintf("{%s %d}", charStr(v.Rune), v.Pos()-delta))
	}
	return a
}

type x []struct {
	c   rune
	pos token.Pos
}

type lexerTests []struct {
	src   string
	chars x
}

func testLexer(t *testing.T, newLexer func(i int, src string, report *xc.Report) (*lexer, error), tab lexerTests) {
nextTest:
	for ti, test := range tab {
		//dbg("==== %v", ti)
		report := xc.NewReport()
		lx, err := newLexer(ti, test.src, report)
		if err != nil {
			t.Fatal(err)
		}

		delta := token.Pos(lx.file.Base() - 1)
		var chars []lex.Char
		var c lex.Char
		for i := 0; c.Rune != ccEOF && i < len(test.src)+2; i++ {
			c = lx.scanChar()
			chars = append(chars, c)
		}
		if c.Rune != ccEOF {
			t.Errorf("%d: scanner stall %v", ti, charsStr(chars, delta))
			continue
		}

		if g, e := report.Errors(true), error(nil); g != e {
			t.Errorf("%d: lx.err %v %v %v", ti, g, e, charsStr(chars, delta))
			continue
		}

		if g, e := len(chars), len(test.chars); g != e {
			t.Errorf("%d: len(chars) %v %v %v", ti, g, e, charsStr(chars, delta))
			continue
		}

		for i, c := range chars {
			c = chars[i]
			e := test.chars[i]
			g := c.Rune
			if c.Rune == ccEOF {
				g = -1
			}
			if e := e.c; g != e {
				t.Errorf("%d: c[%d] %v %v %v", ti, i, charStr(g), charStr(e), charsStr(chars, delta))
				continue nextTest
			}

			if g, e := c.Pos()-delta, e.pos; g != e {
				t.Errorf("%d: pos[%d] %v %v %v", ti, i, g, e, charsStr(chars, delta))
				continue nextTest
			}
		}
	}
}

func TestLexer(t *testing.T) {
	testLexer(
		t,
		func(i int, src string, report *xc.Report) (*lexer, error) {
			return newLexer(fmt.Sprintf("TestLexer.%d", i), len(src), strings.NewReader(src), report, testTweaks)
		},
		lexerTests{
			{"", x{{-1, 1}}},
			{"%0", x{{'%', 1}, {INTCONST, 2}, {-1, 3}}},
			{"%:%:", x{{PPPASTE, 1}, {-1, 5}}},
			{"%>", x{{'}', 1}, {-1, 3}}},
			{"0", x{{INTCONST, 1}, {-1, 2}}},
			{"01", x{{INTCONST, 1}, {-1, 3}}},
			{"0??/1\n", x{{INTCONST, 1}, {'?', 2}, {'?', 3}, {'/', 4}, {INTCONST, 5}, {'\n', 6}, {-1, 7}}},
			{"0??/1\n2", x{{INTCONST, 1}, {'?', 2}, {'?', 3}, {'/', 4}, {INTCONST, 5}, {'\n', 6}, {INTCONST, 7}, {-1, 8}}},
			{"0??/\n", x{{INTCONST, 1}, {'?', 2}, {'?', 3}, {'/', 4}, {'\n', 5}, {-1, 6}}},
			{"0??/\n2", x{{INTCONST, 1}, {'?', 2}, {'?', 3}, {'/', 4}, {'\n', 5}, {INTCONST, 6}, {-1, 7}}},
			{"0\\1\n", x{{INTCONST, 1}, {'\\', 2}, {INTCONST, 3}, {'\n', 4}, {-1, 5}}},
			{"0\\1\n2", x{{INTCONST, 1}, {'\\', 2}, {INTCONST, 3}, {'\n', 4}, {INTCONST, 5}, {-1, 6}}},
			{"0\\\n", x{{INTCONST, 1}, {-1, 4}}},
			{"0\\\n2", x{{INTCONST, 1}, {-1, 5}}},
			{"0\x00", x{{INTCONST, 1}, {0, 2}, {-1, 3}}},
			{"0\x001", x{{INTCONST, 1}, {0, 2}, {INTCONST, 3}, {-1, 4}}},
			{":>", x{{']', 1}, {-1, 3}}},
			{"<%", x{{'{', 1}, {-1, 3}}},
			{"<:", x{{'[', 1}, {-1, 3}}},
			{"??!", x{{'?', 1}, {'?', 2}, {'!', 3}, {-1, 4}}},
			{"??!0", x{{'?', 1}, {'?', 2}, {'!', 3}, {INTCONST, 4}, {-1, 5}}},
			{"??!01", x{{'?', 1}, {'?', 2}, {'!', 3}, {INTCONST, 4}, {-1, 6}}},
			{"??!=", x{{'?', 1}, {'?', 2}, {NEQ, 3}, {-1, 5}}},
			{"??'", x{{'?', 1}, {'?', 2}, {'\'', 3}, {-1, 4}}},
			{"??(", x{{'?', 1}, {'?', 2}, {'(', 3}, {-1, 4}}},
			{"??)", x{{'?', 1}, {'?', 2}, {')', 3}, {-1, 4}}},
			{"??-", x{{'?', 1}, {'?', 2}, {'-', 3}, {-1, 4}}},
			{"??/", x{{'?', 1}, {'?', 2}, {'/', 3}, {-1, 4}}},
			{"??/1\n", x{{'?', 1}, {'?', 2}, {'/', 3}, {INTCONST, 4}, {'\n', 5}, {-1, 6}}},
			{"??/1\n2", x{{'?', 1}, {'?', 2}, {'/', 3}, {INTCONST, 4}, {'\n', 5}, {INTCONST, 6}, {-1, 7}}},
			{"??/\n", x{{'?', 1}, {'?', 2}, {'/', 3}, {'\n', 4}, {-1, 5}}},
			{"??/\n2", x{{'?', 1}, {'?', 2}, {'/', 3}, {'\n', 4}, {INTCONST, 5}, {-1, 6}}},
			{"??<", x{{'?', 1}, {'?', 2}, {'<', 3}, {-1, 4}}},
			{"??=??=", x{{'?', 1}, {'?', 2}, {'=', 3}, {'?', 4}, {'?', 5}, {'=', 6}, {-1, 7}}},
			{"??>", x{{'?', 1}, {'?', 2}, {'>', 3}, {-1, 4}}},
			{"???!", x{{'?', 1}, {'?', 2}, {'?', 3}, {'!', 4}, {-1, 5}}},
			{"???!0", x{{'?', 1}, {'?', 2}, {'?', 3}, {'!', 4}, {INTCONST, 5}, {-1, 6}}},
			{"???/\n2", x{{'?', 1}, {'?', 2}, {'?', 3}, {'/', 4}, {'\n', 5}, {INTCONST, 6}, {-1, 7}}},
			{"????!0", x{{'?', 1}, {'?', 2}, {'?', 3}, {'?', 4}, {'!', 5}, {INTCONST, 6}, {-1, 7}}},
			{"???x0", x{{'?', 1}, {'?', 2}, {'?', 3}, {IDENTIFIER, 4}, {-1, 6}}},
			{"???x??!0", x{{'?', 1}, {'?', 2}, {'?', 3}, {IDENTIFIER, 4}, {'?', 5}, {'?', 6}, {'!', 7}, {INTCONST, 8}, {-1, 9}}},
			{"??x0", x{{'?', 1}, {'?', 2}, {IDENTIFIER, 3}, {-1, 5}}},
			{"??x??!0", x{{'?', 1}, {'?', 2}, {IDENTIFIER, 3}, {'?', 4}, {'?', 5}, {'!', 6}, {INTCONST, 7}, {-1, 8}}},
			{"?x0", x{{'?', 1}, {IDENTIFIER, 2}, {-1, 4}}},
			{"?x??!0", x{{'?', 1}, {IDENTIFIER, 2}, {'?', 3}, {'?', 4}, {'!', 5}, {INTCONST, 6}, {-1, 7}}},
			{"@", x{{'@', 1}, {-1, 2}}},
			{"@%", x{{'@', 1}, {'%', 2}, {-1, 3}}},
			{"@%0", x{{'@', 1}, {'%', 2}, {INTCONST, 3}, {-1, 4}}},
			{"@%:", x{{'@', 1}, {'#', 2}, {-1, 4}}},
			{"@%:0", x{{'@', 1}, {'#', 2}, {INTCONST, 4}, {-1, 5}}},
			{"@%:01", x{{'@', 1}, {'#', 2}, {INTCONST, 4}, {-1, 6}}},
			{"@??=", x{{'@', 1}, {'?', 2}, {'?', 3}, {'=', 4}, {-1, 5}}},
			{"\"(a\\\nz", x{{'"', 1}, {'(', 2}, {IDENTIFIER, 3}, {-1, 7}}},
			{"\\1\n", x{{'\\', 1}, {INTCONST, 2}, {'\n', 3}, {-1, 4}}},
			{"\\1\n2", x{{'\\', 1}, {INTCONST, 2}, {'\n', 3}, {INTCONST, 4}, {-1, 5}}},
			{"\\\n", x{{-1, 3}}},
			{"\\\n2", x{{INTCONST, 3}, {-1, 4}}},
			{"\\\r\n", x{{-1, 4}}},
			{"\\\r\n2", x{{INTCONST, 4}, {-1, 5}}},
			{"\r", x{{-1, 2}}},
			{"\r0", x{{INTCONST, 2}, {-1, 3}}},
			{"\r01", x{{INTCONST, 2}, {-1, 4}}},
			{"\x00", x{{0, 1}, {-1, 2}}},
			{"\x000", x{{0, 1}, {INTCONST, 2}, {-1, 3}}},
		},
	)
}

func TestLexer2(t *testing.T) {
	testLexer(
		t,
		func(i int, src string, report *xc.Report) (*lexer, error) {
			tweaks := *testTweaks
			tweaks.enableTrigraphs = true
			return newLexer(fmt.Sprintf("TestLexer.%d", i), len(src), strings.NewReader(src), report, &tweaks)
		},
		lexerTests{
			{"", x{{-1, 1}}},
			{"%0", x{{'%', 1}, {INTCONST, 2}, {-1, 3}}},
			{"%:%:", x{{PPPASTE, 1}, {-1, 5}}},
			{"%>", x{{'}', 1}, {-1, 3}}},
			{"0", x{{INTCONST, 1}, {-1, 2}}},
			{"01", x{{INTCONST, 1}, {-1, 3}}},
			{"0??/1\n", x{{INTCONST, 1}, {'\\', 2}, {INTCONST, 5}, {'\n', 6}, {-1, 7}}},
			{"0??/1\n2", x{{INTCONST, 1}, {'\\', 2}, {INTCONST, 5}, {'\n', 6}, {INTCONST, 7}, {-1, 8}}},
			{"0??/\n", x{{INTCONST, 1}, {-1, 6}}},
			{"0??/\n2", x{{INTCONST, 1}, {-1, 7}}},
			{"0\\1\n", x{{INTCONST, 1}, {'\\', 2}, {INTCONST, 3}, {'\n', 4}, {-1, 5}}},
			{"0\\1\n2", x{{INTCONST, 1}, {'\\', 2}, {INTCONST, 3}, {'\n', 4}, {INTCONST, 5}, {-1, 6}}},
			{"0\\\n", x{{INTCONST, 1}, {-1, 4}}},
			{"0\\\n2", x{{INTCONST, 1}, {-1, 5}}},
			{"0\x00", x{{INTCONST, 1}, {0, 2}, {-1, 3}}},
			{"0\x001", x{{INTCONST, 1}, {0, 2}, {INTCONST, 3}, {-1, 4}}},
			{":>", x{{']', 1}, {-1, 3}}},
			{"<%", x{{'{', 1}, {-1, 3}}},
			{"<:", x{{'[', 1}, {-1, 3}}},
			{"??!", x{{'|', 1}, {-1, 4}}},
			{"??!0", x{{'|', 1}, {INTCONST, 4}, {-1, 5}}},
			{"??!01", x{{'|', 1}, {INTCONST, 4}, {-1, 6}}},
			{"??!=", x{{ORASSIGN, 1}, {-1, 5}}},
			{"??'", x{{'^', 1}, {-1, 4}}},
			{"??(", x{{'[', 1}, {-1, 4}}},
			{"??)", x{{']', 1}, {-1, 4}}},
			{"??-", x{{'~', 1}, {-1, 4}}},
			{"??/", x{{'\\', 1}, {-1, 4}}},
			{"??/1\n", x{{'\\', 1}, {INTCONST, 4}, {'\n', 5}, {-1, 6}}},
			{"??/1\n2", x{{'\\', 1}, {INTCONST, 4}, {'\n', 5}, {INTCONST, 6}, {-1, 7}}},
			{"??/\n", x{{-1, 5}}},
			{"??/\n2", x{{INTCONST, 5}, {-1, 6}}},
			{"??<", x{{'{', 1}, {-1, 4}}},
			{"??=??=", x{{PPPASTE, 1}, {-1, 7}}},
			{"??>", x{{'}', 1}, {-1, 4}}},
			{"???!", x{{'?', 1}, {'|', 2}, {-1, 5}}},
			{"???!0", x{{'?', 1}, {'|', 2}, {INTCONST, 5}, {-1, 6}}},
			{"???/\n2", x{{'?', 1}, {INTCONST, 6}, {-1, 7}}},
			{"????!0", x{{'?', 1}, {'?', 2}, {'|', 3}, {INTCONST, 6}, {-1, 7}}},
			{"???x0", x{{'?', 1}, {'?', 2}, {'?', 3}, {IDENTIFIER, 4}, {-1, 6}}},
			{"???x??!0", x{{'?', 1}, {'?', 2}, {'?', 3}, {IDENTIFIER, 4}, {'|', 5}, {INTCONST, 8}, {-1, 9}}},
			{"??x0", x{{'?', 1}, {'?', 2}, {IDENTIFIER, 3}, {-1, 5}}},
			{"??x??!0", x{{'?', 1}, {'?', 2}, {IDENTIFIER, 3}, {'|', 4}, {INTCONST, 7}, {-1, 8}}},
			{"?x0", x{{'?', 1}, {IDENTIFIER, 2}, {-1, 4}}},
			{"?x??!0", x{{'?', 1}, {IDENTIFIER, 2}, {'|', 3}, {INTCONST, 6}, {-1, 7}}},
			{"@", x{{'@', 1}, {-1, 2}}},
			{"@%", x{{'@', 1}, {'%', 2}, {-1, 3}}},
			{"@%0", x{{'@', 1}, {'%', 2}, {INTCONST, 3}, {-1, 4}}},
			{"@%:", x{{'@', 1}, {'#', 2}, {-1, 4}}},
			{"@%:0", x{{'@', 1}, {'#', 2}, {INTCONST, 4}, {-1, 5}}},
			{"@%:01", x{{'@', 1}, {'#', 2}, {INTCONST, 4}, {-1, 6}}},
			{"@??=", x{{'@', 1}, {'#', 2}, {-1, 5}}},
			{"\"(a\\\nz", x{{'"', 1}, {'(', 2}, {IDENTIFIER, 3}, {-1, 7}}},
			{"\\1\n", x{{'\\', 1}, {INTCONST, 2}, {'\n', 3}, {-1, 4}}},
			{"\\1\n2", x{{'\\', 1}, {INTCONST, 2}, {'\n', 3}, {INTCONST, 4}, {-1, 5}}},
			{"\\\n", x{{-1, 3}}},
			{"\\\n2", x{{INTCONST, 3}, {-1, 4}}},
			{"\\\r\n", x{{-1, 4}}},
			{"\\\r\n2", x{{INTCONST, 4}, {-1, 5}}},
			{"\r", x{{-1, 2}}},
			{"\r0", x{{INTCONST, 2}, {-1, 3}}},
			{"\r01", x{{INTCONST, 2}, {-1, 4}}},
			{"\x00", x{{0, 1}, {-1, 2}}},
			{"\x000", x{{0, 1}, {INTCONST, 2}, {-1, 3}}},
		},
	)
}

func TestPPParse1(t *testing.T) {
	path := *o1
	if path == "" {
		return
	}

	testReport := newTestReport()
	testReport.ClearErrors()
	_, err := ppParse(path, testReport, testTweaks)
	if err != nil {
		t.Fatal(err)
	}

	if err := testReport.Errors(true); err != nil {
		t.Fatal(errString(err))
	}
}

func testPreprocess1(t *testing.T, paths []string) {
	testReport := xc.NewReport()
	testReport.ClearErrors()
	predefined, err := ppParseString("", predefinedMacros, testReport, testTweaks)
	if err != nil {
		t.Fatal(err)
	}

	model := newTestModel()
	macros := newMacros()
	func() {
		ch := make(chan []xc.Token, 1000)
		defer close(ch)

		go func() {
			for range ch {
			}
		}()

		newPP(ch, includes, sysIncludes, macros, false, model, testReport, testTweaks).preprocessingFile(predefined)
		if err := testReport.Errors(true); err != nil {
			t.Fatal(errString(err))
		}

		for _, path := range paths {
			ast, err := ppParse(path, testReport, testTweaks)
			if err != nil {
				t.Fatal(err)
			}

			newPP(ch, includes, sysIncludes, macros, true, model, testReport, testTweaks).preprocessingFile(ast)
			if err := testReport.Errors(true); err != nil {
				t.Fatal(errString(err))
			}
		}
	}()
}

func TestPreprocess(t *testing.T) {
	const arith = "testdata/arith-1.c"
	a := []string{}
	if _, err := os.Stat(arith); err == nil {
		a = []string{arith}
	}

	for _, v := range a {
		testPreprocess1(t, []string{v})
	}
}

func TestFinalInjection(t *testing.T) {
	const src = "int f() {}"

	if strings.HasSuffix(src, "\n") {
		t.Fatal("internal error")
	}

	ast, err := ppParseString("test.c", src, xc.NewReport(), &tweaks{})
	if err != nil {
		t.Fatal(errString(err))
	}

	ast2, err := ppParseString("test.c", src+"\n", xc.NewReport(), &tweaks{})
	if err != nil {
		t.Fatal(errString(err))
	}

	if g, e := PrettyString(ast2), PrettyString(ast); g != e {
		t.Fatalf("got\n%s\nexpected\n%s", g, e)
	}

	t.Log(PrettyString(ast))
}

func TestRedecl(t *testing.T) {
	testParse(t, []string{"testdata/redecl.c"})
}

func testParse(t *testing.T, paths []string) *TranslationUnit {
	last := paths[len(paths)-1]
	ln := filepath.Base(last)
	f, err := os.Create("log-" + ln)
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	b := bufio.NewWriter(f)
	defer b.Flush()

	b.WriteString("// +build ignore\n\n")
	var a []string
	crash := nopOpt()
	if *oFailFast {
		crash = CrashOnError()
	}
	ast, err := Parse(
		predefinedMacros,
		paths,
		newTestModel(),
		IncludePaths(includes),
		SysIncludePaths(sysIncludes),
		Cpp(func(toks []xc.Token) {
			a = a[:0]
			for _, v := range toks {
				a = append(a, TokSrc(v))
			}
			fmt.Fprintf(b, "%s\n", strings.Join(a, " "))
		}),
		ErrLimit(-1),
		crash,
	)
	if err != nil {
		t.Fatal(errString(err))
	}

	t.Log(paths)
	return ast
}

func ddsStr(dds []*DirectDeclarator) string {
	buf := bytes.NewBufferString("|")
	for i, dd := range dds {
		switch dd.Case {
		case 0: // IDENTIFIER
			buf.WriteString("IDENTIFIER")
		case 1: // '(' Declarator ')'
			buf.WriteString("(")
			buf.WriteString(strings.Repeat("*", dd.Declarator.stars()))
			fmt.Fprintf(buf, "Declarator.%v)", dds[i-1].Case)
		case 2: // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
			fmt.Fprintf(buf, "DirectDeclarator[TypeQualifierListOpt ExpressionOpt.%v]", dd.elements)
		case 3: // DirectDeclarator '[' "static" TypeQualifierListOpt Expression ']'
			fmt.Fprintf(buf, "DirectDeclarator[static TypeQualifierListOpt Expression.%v]", dd.elements)
		case 4: // DirectDeclarator '[' TypeQualifierList "static" Expression ']'
			fmt.Fprintf(buf, "DirectDeclarator[TypeQualifierList static Expression.%v]", dd.elements)
		case 5: // DirectDeclarator '[' TypeQualifierListOpt '*' ']'
			fmt.Fprintf(buf, "DirectDeclarator[TypeQualifierListOpt*.%v]", dd.elements)
		case 6: // DirectDeclarator '(' ParameterTypeList ')'
			buf.WriteString("DirectDeclarator(ParameterTypeList)")
		case 7: // DirectDeclarator '(' IdentifierListOpt ')'
			buf.WriteString("DirectDeclarator(IdentifierListOpt)")
		}
		buf.WriteString("|")
	}
	return buf.String()
}

func (n *ctype) str() string {
	return fmt.Sprintf("R%v S%v %v", n.resultStars, n.stars, ddsStr(n.dds))
}

func TestIssue3(t *testing.T) {
	if _, err := Parse("", []string{"testdata/issue3.h"}, newTestModel()); err != nil {
		t.Fatal(errString(err))
	}
}

func TestIssue8(t *testing.T) {
	if _, err := Parse("", []string{"testdata/issue8.h"}, newTestModel()); err != nil {
		t.Fatal(errString(err))
	}
}

func TestIssue4(t *testing.T) {
	_, err := Parse("", []string{"testdata/issue4.c"}, newTestModel())
	if err == nil {
		t.Fatal("unexpected sucess")
	}

	l, ok := err.(scanner.ErrorList)
	if !ok {
		t.Fatalf("unexpected error type %T", err)
	}

	if g, e := l.Len(), 2; g != e {
		t.Fatal(g, e)
	}

	if g, e := l[0].Error(), "testdata/issue4.c:5:13: redeclaration of foo as different kind of symbol, previous declaration at testdata/issue4.c:4:5"; g != e {
		t.Fatal(g, e)
	}

	if g, e := l[1].Error(), "testdata/issue4.c:9:15: redeclaration of foo2 as different kind of symbol, previous declaration at testdata/issue4.c:8:7"; g != e {
		t.Fatal(g, e)
	}
}

func unpackType(typ Type) Type {
	for {
		switch typ.Kind() {
		case Ptr, Array:
			typ = typ.Element()
		default:
			return typ
		}
	}
}

func TestIssue9(t *testing.T) {
	const exp = `original:  typedef short[64] Array
unpacked:  typedef short Short
original: JBLOCK short(*)[64] Ptr
unpacked: JBLOCK short Short
original: JBLOCKROW short(**)[64] Ptr
unpacked: JBLOCKROW short Short
original: JBLOCKARRAY short(***)[64] Ptr
unpacked: JBLOCKARRAY short Short
original:  short[64] Array
unpacked:  short Short
original:  short[64] Array
unpacked:  short Short
original: JBLOCK short[64] Array
unpacked: JBLOCK short Short
original: JBLOCK short[64] Array
unpacked: JBLOCK short Short
original:  short(*)[64] Ptr
unpacked:  short Short
original:  short(*)[64] Ptr
unpacked:  short Short
original:  short(*)[64] Ptr
unpacked:  short Short
original: JBLOCKROW short(*)[64] Ptr
unpacked: JBLOCKROW short Short
original: JBLOCKROW short(*)[64] Ptr
unpacked: JBLOCKROW short Short
original:  short(**)[64] Ptr
unpacked:  short Short
original:  short(**)[64] Ptr
unpacked:  short Short
original:  short(**)[64] Ptr
unpacked:  short Short
original: JBLOCKARRAY short(**)[64] Ptr
unpacked: JBLOCKARRAY short Short
original: JBLOCKARRAY short(**)[64] Ptr
unpacked: JBLOCKARRAY short Short
original:  short(***)[64] Ptr
unpacked:  short Short
original:  short(***)[64] Ptr
unpacked:  short Short
original:  short(***)[64] Ptr
unpacked:  short Short
original: JBLOCKIMAGE short(***)[64] Ptr
unpacked: JBLOCKIMAGE short Short
original: JBLOCKIMAGE short(***)[64] Ptr
unpacked: JBLOCKIMAGE short Short
`

	tu, err := Parse("", []string{"testdata/issue9.c"}, newTestModel())
	if err != nil {
		t.Fatal(errString(err))
	}

	var buf bytes.Buffer
	for tu != nil {
		declr := tu.ExternalDeclaration.Declaration.InitDeclaratorListOpt.InitDeclaratorList.InitDeclarator.Declarator
		name := string(xc.Dict.S(declr.RawSpecifier().TypedefName()))
		fmt.Fprintln(&buf, "original:", name, declr.Type, declr.Type.Kind())
		unpacked := unpackType(declr.Type)
		fmt.Fprintln(&buf, "unpacked:", name, unpacked, unpacked.Kind())
		tu = tu.TranslationUnit
	}
	if g, e := buf.String(), exp; g != e {
		t.Fatalf("got:\n%s\nexp:\n%s", g, e)
	}
}

func TestEnumConstToks(t *testing.T) {
	tu, err := Parse("", []string{"testdata/enum.c"}, newTestModel())
	if err != nil {
		t.Fatal(errString(err))
	}

	sc := tu.Declarations
	foo := sc.Lookup(NSIdentifiers, xc.Dict.SID("foo"))
	if foo.Node == nil {
		t.Fatal()
	}

	switch x := foo.Node.(type) {
	case *DirectDeclarator:
		typ := x.TopDeclarator().Type
		if g, e := typ.Kind(), Enum; g != e {
			t.Fatal(g, e)
		}

		l := typ.EnumeratorList()
		if g, e := PrettyString(
			l),
			`[]cc.EnumConstant{ // len 2
· 0: cc.EnumConstant{
· · DefTok: testdata/enum.c:4:2: IDENTIFIER "c",
· · Value: 18,
· · Tokens: []xc.Token{ // len 3
· · · 0: testdata/enum.c:4:6: INTCONST "42",
· · · 1: testdata/enum.c:4:9: '-',
· · · 2: testdata/enum.c:4:11: INTCONST "24",
· · },
· },
· 1: cc.EnumConstant{
· · DefTok: testdata/enum.c:5:2: IDENTIFIER "d",
· · Value: 592,
· · Tokens: []xc.Token{ // len 3
· · · 0: testdata/enum.c:5:6: INTCONST "314",
· · · 1: testdata/enum.c:5:10: '+',
· · · 2: testdata/enum.c:1:11: INTCONST "278",
· · },
· },
}`; g != e {
			t.Fatalf("got\n%s\nexp\n%s", g, e)
		}
		var a []string
		for _, e := range l {
			var b []string
			for _, t := range e.Tokens {
				b = append(b, TokSrc(t))
			}
			a = append(a, strings.Join(b, " "))
		}
		if g, e := strings.Join(a, "\n"), "42 - 24\n314 + 278"; g != e {
			t.Fatalf("got\n%s\nexp\n%s", g, e)
		}
	default:
		t.Fatalf("%T", x)
	}
}

func TestPaste(t *testing.T) {
	testParse(t, []string{"testdata/paste.c"})
}

func TestGCCPredefs(t *testing.T) {
	if _, err := Parse(gccPredefined, []string{"testdata/gcc.c"}, newTestModel()); err != nil {
		t.Fatal(err)
	}
}

func TestPaste2(t *testing.T) {
	testParse(t, []string{"testdata/paste2.c"})
}

func TestFunc(t *testing.T) {
	testParse(t, []string{"testdata/func.c"})
}

func TestEmptyMacroArg(t *testing.T) {
	testParse(t, []string{"testdata/empty.c"})
}
