// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	flags "flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
	"modernc.org/mathutil"
)

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "# caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# \tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	pc, fn, fl, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	fmt.Fprintf(os.Stderr, "# dbg %s:%d:%s: ", path.Base(fn), fl, f.Name())
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func TODO(...interface{}) string { //TODOOK
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("# TODO: %s:%d:\n", path.Base(fn), fl) //TODOOK
}

func stack() []byte { return debug.Stack() }

func use(...interface{}) {}

func init() {
	use(caller, dbg, TODO, stack, tokStr, shift) //TODOOK
}

// ----------------------------------------------------------------------------

const (
	maxFiles = 2000

	parserTestBuiltin = `
#define __DI__
#define __FUNCTION__ __func__
#define __HI__
#define __PRETTY_FUNCTION__ __func__
#define __QI__
#define __SI__
#define __builtin_add_overflow(...) 0
#define __builtin_constant_p(x) (__builtin_constant_p_impl(0, x))
#define __builtin_expect(exp, c) (exp)
#define __builtin_mul_overflow(...) 0
#define __builtin_offsetof(type, member) ((__SIZE_TYPE__)&(((type*)0)->member))
#define __builtin_sub_overflow(...) 0
#define __builtin_va_arg(ap, type) ((type)__builtin_va_arg_impl(ap))
#define __builtin_va_end(ap)
#define __builtin_va_start(ap, v)
#define __declspec(...)
#define __extension__
#define __read_only__ 0
#define __signed signed
#define __sync_synchronize(...)
#define __word__
#define __write_only__ 0
#define asm __asm__

#ifndef __GNUC__
#define __attribute__(x)
#endif

#ifdef __PTRDIFF_TYPE__
typedef __PTRDIFF_TYPE__ ptrdiff_t;
#endif

#ifdef __SIZE_TYPE__
typedef __SIZE_TYPE__ size_t;
#endif

#ifdef __WCHAR_TYPE__
typedef __WCHAR_TYPE__ wchar_t;
#endif

typedef void *__builtin_va_list;

#if __SIZEOF_POINTER__ == 8
#else
typedef long double __float128;
#endif

#ifdef __UINT16_TYPE__
__UINT16_TYPE__ __builtin_bswap16 (__UINT16_TYPE__ x);
#endif

#ifdef __UINT32_TYPE__
__UINT32_TYPE__ __builtin_bswap32 (__UINT32_TYPE__ x);
#endif

#ifdef __UINT64_TYPE__
__UINT64_TYPE__ __builtin_bswap64 (__UINT64_TYPE__ x);
#endif

#ifdef __SIZEOF_INT128__
typedef __INT64_TYPE__ __int128_t[2];	//TODO
typedef __UINT64_TYPE__ __uint128_t[2];	//TODO
#endif

#if defined(__MINGW32__) || defined(__MINGW64__)
int gnu_printf(const char *format, ...);
int gnu_scanf(const char *format, ...);
int ms_scanf(const char *format, ...);
int ms_printf(const char *format, ...);
#endif

typedef struct { char real, imag; } __COMPLEX_CHAR_TYPE__;
typedef struct { double real, imag; } __COMPLEX_DOUBLE_TYPE__;
typedef struct { float real, imag; } __COMPLEX_FLOAT_TYPE__;
typedef struct { int real, imag; } __COMPLEX_INT_TYPE__;
typedef struct { long double real, imag; } __COMPLEX_LONG_DOUBLE_TYPE__;
typedef struct { long real, imag; } __COMPLEX_LONG_TYPE__;
typedef struct { long long real, imag; } __COMPLEX_LONG_LONG_TYPE__;
typedef struct { long long unsigned real, imag; } __COMPLEX_LONG_LONG_UNSIGNED_TYPE__;
typedef struct { long unsigned real, imag; } __COMPLEX_LONG_UNSIGNED_TYPE__;
typedef struct { short real, imag; } __COMPLEX_SHORT_TYPE__;
typedef struct { unsigned real, imag; } __COMPLEX_UNSIGNED_TYPE__;
typedef struct { unsigned short real, imag; } __COMPLEX_SHORT_UNSIGNED_TYPE__;

int __builtin_clzll (unsigned long long);
int __builtin_constant_p_impl(int, ...);
int __printf__ ( const char * format, ... );
int __scanf__ ( const char *format, ... );
void *__builtin_alloca (size_t);
void *__builtin_extract_return_addr (void *addr);
void *__builtin_frame_address (unsigned int level);
void *__builtin_malloc(size_t);
void *__builtin_stack_save(void);
void *__builtin_va_arg_impl(void* ap);

`
)

var (
	oBlackBox = flags.String("blackbox", "", "Record CSmith file to this file")
	oCSmith   = flags.Duration("csmith", 10*time.Second, "")
	oMaxFiles = flags.Int("maxFiles", maxFiles, "")
	oRE       = flags.String("re", "", "")
	oSkipInit = flags.Bool("skipInit", false, "")
	oTrace    = flags.Bool("trc", false, "Print tested paths.")
	oWalkDir  = flags.String("walkDir", "testdata", "")

	gccDir    = filepath.FromSlash("testdata/gcc-9.1.0")
	sqliteDir = filepath.FromSlash("testdata/sqlite-amalgamation-3300100")
	tccDir    = filepath.FromSlash("testdata/tcc-0.9.27")

	csmithArgs = strings.Join([]string{
		"--bitfields",                     // --bitfields | --no-bitfields: enable | disable full-bitfields structs (disabled by default).
		"--no-const-pointers",             // --const-pointers | --no-const-pointers: enable | disable const pointers (enabled by default).
		"--no-consts",                     // --consts | --no-consts: enable | disable const qualifier (enabled by default).
		"--no-packed-struct",              // --packed-struct | --no-packed-struct: enable | disable packed structs by adding #pragma pack(1) before struct definition (disabled by default).
		"--no-volatile-pointers",          // --volatile-pointers | --no-volatile-pointers: enable | disable volatile pointers (enabled by default).
		"--no-volatiles",                  // --volatiles | --no-volatiles: enable | disable volatiles (enabled by default).
		"--paranoid",                      // --paranoid | --no-paranoid: enable | disable pointer-related assertions (disabled by default).
		"--max-nested-struct-level", "10", // --max-nested-struct-level <num>: limit maximum nested level of structs to <num>(default 0). Only works in the exhaustive mode.
	}, " ")

	testBuiltinSource *cachedPPFile
	testIncludes      []string
	testPredef        string
	testPredefSource  *cachedPPFile
	testSQLiteSource  *cachedPPFile
	testShellSource   *cachedPPFile
	testSysIncludes   []string
	testWD            string

	testABI = newTestABI()
)

func newTestABI() ABI {
	abi, err := NewABIFromEnv()
	if err != nil {
		panic(err)
	}
	return abi
}

func TestMain(m *testing.M) {
	defer func() {
		os.Exit(m.Run())
	}()

	// fmt.Printf("test binary compiled for %s/%s\n", runtime.GOOS, runtime.GOARCH)
	// fmt.Printf("CC_TEST_CPP=%q\n", os.Getenv("CC_TEST_CPP"))

	isTesting = true
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flags.BoolVar(&panicOnParserError, "panicOnParserError", false, "Panic on parser error.") //TODOOK
	flags.BoolVar(&debugWorkingDir, "dbgWorkingDir", false, "")
	flags.BoolVar(&debugIncludePaths, "dbgIncludePaths", false, "")

	flags.Parse()
	var err error
	if testWD, err = os.Getwd(); err != nil {
		log.Fatalf("Cannot determine working dir: %v", err)
	}

	if s := *oWalkDir; !filepath.IsAbs(s) {
		*oWalkDir = filepath.Join(testWD, s)
	}

	if testPredef, testIncludes, testSysIncludes, err = HostConfig(os.Getenv("CC_TEST_CPP")); err != nil {
		log.Fatal("Cannot acquire host cpp configuration.")
		return
	}

	if runtime.GOOS == "darwin" {
		switch runtime.GOARCH {
		case "amd64":
			testPredef += `
				#define TARGET_CPU_X86_64 1
				#define TARGET_OS_UNIX 1
			`
		case "arm64":
			testPredef += `
				#define TARGET_OS_UNIX 1
			`
		default:
			log.Fatalf("clang: unknown/unsupported GOARCH: %s", runtime.GOARCH)
		}
	}

	isTestingMingw = detectMingw(testPredef)
	if s := os.Getenv("CSMITH_PATH"); s != "" {
		testIncludes = append(testIncludes, s) //TODO nix only
	}
	testIncludes = append(testIncludes, filepath.FromSlash("/usr/include/csmith")) //TODO nix only

	if *oSkipInit {
		return
	}

	cfg := &Config{}
	if testPredefSource, err = cache.getValue(newContext(cfg), "<predefined>", testPredef, false, false); err != nil {
		log.Fatal(err)
	}

	if testBuiltinSource, err = cache.getValue(newContext(cfg), "<built-in>", parserTestBuiltin, false, false); err != nil {
		log.Fatal(err)
	}

	testSysIncludes = testSysIncludes[:len(testSysIncludes):len(testSysIncludes)]
	testIncludes = testIncludes[:len(testIncludes):len(testIncludes)]
	testIncludes = append(testIncludes, "@")
	testIncludes = append(testIncludes, testSysIncludes...)

	path := filepath.Join(testWD, sqliteDir, "shell.c")
	if testShellSource, err = cache.getFile(newContext(cfg), path, false, false); err != nil {
		log.Fatal(err)
	}

	path = filepath.Join(testWD, sqliteDir, "sqlite3.c")
	if testSQLiteSource, err = cache.getFile(newContext(cfg), path, false, false); err != nil {
		log.Fatal(err)
	}
}

func h(v interface{}) string {
	switch x := v.(type) {
	case int:
		return humanize.Comma(int64(x))
	case int64:
		return humanize.Comma(x)
	case uint64:
		return humanize.Comma(int64(x))
	case float64:
		return humanize.CommafWithDigits(x, 0)
	default:
		panic(internalErrorf("%T", x))
	}
}

func skipDir(path string) error {
	if strings.HasPrefix(filepath.Base(path), ".") {
		return filepath.SkipDir
	}

	return nil
}

func shift(tok Token) string {
	pc, _, _, _ := runtime.Caller(2)
	caller := runtime.FuncForPC(pc - 1)
	return fmt.Sprintf("# %v: %v", caller.Name(), PrettyString(tok))
}

func exampleAST(rule int, src string) interface{} {
	src = strings.Replace(src, "\\n", "\n", -1)
	cfg := &Config{ignoreErrors: true, PreprocessOnly: true}
	ctx := newContext(cfg)
	ctx.keywords = gccKeywords
	ast, _ := parse(ctx, nil, nil, []Source{{Name: "example.c", Value: src, DoNotCache: true}})
	if ast == nil {
		return "FAIL"
	}

	pc, _, _, _ := runtime.Caller(1)
	typ := runtime.FuncForPC(pc - 1).Name()
	i := strings.LastIndexByte(typ, '.')
	typ = typ[i+1+len("Example"):]
	i = strings.LastIndexByte(typ, '_')
	typ = typ[:i]
	var node Node
	depth := mathutil.MaxInt
	findNode(typ, ast.TranslationUnit, 0, &node, &depth)
	return node
}

func findNode(typ string, n Node, depth int, out *Node, pdepth *int) {
	if depth >= *pdepth {
		return
	}

	v := reflect.ValueOf(n)
	if v.Kind() != reflect.Ptr {
		return
	}

	elem := v.Elem()
	if elem.Kind() != reflect.Struct {
		return
	}

	t := reflect.TypeOf(elem.Interface())
	if t.Name() == typ {
		*pdepth = depth
		*out = n
		return
	}

	for i := 0; i < elem.NumField(); i++ {
		fld := t.Field(i)
		if nm := fld.Name[0]; nm < 'A' || nm > 'Z' {
			continue
		}

		if x, ok := elem.FieldByIndex([]int{i}).Interface().(Node); ok {
			findNode(typ, x, depth+1, out, pdepth)
		}
	}
}

type structInfo struct {
	offs      []uintptr
	flds      map[uintptr][]Field
	padBefore map[Field]int
	padAfter  int

	forceAlign bool
}

func newStructInfo(t Type) *structInfo {
	nf := t.NumField()
	flds := map[uintptr][]Field{}
	var maxAlign int
	for idx := []int{0}; idx[0] < nf; idx[0]++ {
		f := t.FieldByIndex(idx)
		if f.IsBitField() && f.BitFieldWidth() == 0 {
			continue
		}

		if a := f.Type().Align(); !f.IsBitField() && a > maxAlign {
			maxAlign = a
		}
		off := f.Offset()
		flds[off] = append(flds[off], f)
	}
	var offs []uintptr
	for k := range flds {
		offs = append(offs, k)
	}
	sort.Slice(offs, func(i, j int) bool { return offs[i] < offs[j] })
	var pads map[Field]int
	var pos uintptr
	var forceAlign bool
	for _, off := range offs {
		f := flds[off][0]
		ft := f.Type()
		//trc("%q off %d pos %d %v %v %v", f.Name(), off, pos, ft, ft.Kind(), ft.IsIncomplete())
		switch {
		case ft.IsBitFieldType():
			if f.BitFieldOffset() != 0 {
				break
			}

			if p := int(off - pos); p != 0 {
				if pads == nil {
					pads = map[Field]int{}
				}
				pads[f] = p
				pos = off
			}
			pos += uintptr(f.BitFieldBlockWidth()) >> 3
		default:
			var sz uintptr
			switch {
			case ft.Kind() != Array || ft.Len() != 0:
				sz = ft.Size()
			default:
				forceAlign = true
			}
			if p := int(off - pos); p != 0 {
				if pads == nil {
					pads = map[Field]int{}
				}
				pads[f] = p
				pos = off
			}
			pos += sz
		}
	}
	return &structInfo{
		offs:       offs,
		flds:       flds,
		padBefore:  pads,
		padAfter:   int(t.Size() - pos),
		forceAlign: forceAlign || maxAlign < t.Align(),
	}
}

func dumpLayout(t Type) string {
	info := newStructInfo(t)
	switch t.Kind() {
	case Struct, Union:
		// ok
	default:
		return t.String()
	}

	nf := t.NumField()
	var a []string
	w := 0
	for i := 0; i < nf; i++ {
		if n := len(t.FieldByIndex([]int{i}).Name().String()); n > w {
			w = n
		}
	}
	for i := 0; i < nf; i++ {
		f := t.FieldByIndex([]int{i})
		var bf StringID
		if f.IsBitField() {
			if bfbf := f.BitFieldBlockFirst(); bfbf != nil {
				bf = bfbf.Name()
			}
		}
		a = append(a, fmt.Sprintf("%3d: %*q: BitFieldOffset %3v, BitFieldWidth %3v, IsBitField %5v, Mask: %#016x, off: %3v, pad %2v, BitFieldBlockWidth: %2d, BitFieldBlockFirst: %s, %v",
			i, w+2, f.Name(), f.BitFieldOffset(), f.BitFieldWidth(),
			f.IsBitField(), f.Mask(), f.Offset(), f.Padding(),
			f.BitFieldBlockWidth(), bf, f.Type(),
		))
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%v\n%s\n----\n", t, strings.Join(a, "\n"))
	fmt.Fprintf(&b, "offs: %v\n", info.offs)
	a = a[:0]
	for k, v := range info.flds {
		var b []string
		for _, w := range v {
			b = append(b, fmt.Sprintf("%q padBefore: %d ", w.Name(), info.padBefore[w]))
		}
		a = append(a, fmt.Sprintf("%4d %s", k, b))
	}
	sort.Strings(a)
	for _, v := range a {
		fmt.Fprintf(&b, "%s\n", v)
	}
	fmt.Fprintf(&b, "padAfter: %v\n", info.padAfter)
	for i := 0; i < nf; i++ {
		f := t.FieldByIndex([]int{i})
		if x, ok := f.Type().(*structType); ok {
			s := dumpLayout(x)
			a := strings.Split(s, "\n")
			fmt.Fprintf(&b, "====\n")
			for _, v := range a {
				fmt.Fprintf(&b, "%s\n", v)
			}
		}
	}
	return b.String()
}
