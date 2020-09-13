// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"compress/bzip2"
	"compress/gzip"
	flags "flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"runtime/debug"
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
#define __builtin_expect(exp, c) exp
#define __builtin_mul_overflow(...) 0
#define __builtin_offsetof(type, member) ((__SIZE_TYPE__)&(((type*)0)->member))
#define __builtin_sub_overflow(...) 0
#define __builtin_va_arg(ap, type) (type)__builtin_va_arg_impl(ap)
#define __builtin_va_end(ap)
#define __builtin_constant_p(x) __builtin_constant_p_impl(0, x)
#define __builtin_va_start(ap, v)
#define __declspec(...)
#define __extension__
#define __sync_synchronize(...)
#define __word__
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

#if __SIZEOF_POINTER__ == 8
typedef void* __builtin_va_list[3];
#else
typedef void* __builtin_va_list[1];
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
void *__builtin_va_arg_impl(void* ap);
`
)

var (
	oBlackBox = flags.String("blackbox", "", "Record CSmith file to this file")
	oCSmith   = flags.Duration("csmith", 10*time.Second, "")
	oDev      = flags.Bool("dev", false, "Enable developer tests/downloads.")
	oDownload = flags.Bool("download", false, "Download missing testdata. Add -dev to download also 100+ MB of developer resources.")
	oMaxFiles = flags.Int("maxFiles", maxFiles, "")
	oRE       = flags.String("re", "", "")
	oSkipInit = flags.Bool("skipInit", false, "")
	oTrace    = flags.Bool("trc", false, "Print tested paths.")
	oWalkDir  = flags.String("walkDir", "testdata", "")

	gccDir    = filepath.FromSlash("testdata/gcc-9.1.0")
	sqliteDir = filepath.FromSlash("testdata/sqlite-amalgamation-3300100")
	tccDir    = filepath.FromSlash("testdata/tcc-0.9.27")

	downloads = []struct {
		dir, url string
		sz       int
		dev      bool
	}{
		{gccDir, "http://mirror.koddos.net/gcc/releases/gcc-9.1.0/gcc-9.1.0.tar.gz", 118000, true},
		{sqliteDir, "https://www.sqlite.org/2019/sqlite-amalgamation-3300100.zip", 2200, false},
		{tccDir, "http://download.savannah.gnu.org/releases/tinycc/tcc-0.9.27.tar.bz2", 620, false},
	}

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

	testBuiltinSource   *cachedPPFile
	testIncludes        []string
	testPredef          string
	testPredefGNU       string
	testPredefGNUSource *cachedPPFile
	testPredefSource    *cachedPPFile
	testSQLiteSource    *cachedPPFile
	testShellSource     *cachedPPFile
	testSysIncludes     []string
	testWD              string

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

	fmt.Printf("test binary compiled for %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("CC_TEST_CPP=%q\n", os.Getenv("CC_TEST_CPP"))

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

	if testPredefGNU, testIncludes, testSysIncludes, err = HostConfig(os.Getenv("CC_TEST_CPP")); err != nil {
		log.Fatal("Cannot acquire host cpp configuration.")
		return
	}

	isTestingMingw = detectMingw(testPredefGNU)
	testIncludes = append(testIncludes, filepath.FromSlash("/usr/include/csmith"))

	if *oSkipInit {
		return
	}

	cfg := &Config{}
	if testPredefGNUSource, err = cache.getValue(newContext(cfg), "<predefined>", testPredefGNU, false, false); err != nil {
		log.Fatal(err)
	}

	switch {
	case isTestingMingw:
		testPredef = testPredefGNU
	default:
		a := strings.Split(testPredefGNU, "\n")
		w := 0
		for _, v0 := range a {
			v := strings.TrimSpace(strings.ToLower(v0))
			if !strings.HasPrefix(v, "#define __gnu") && !strings.HasPrefix(v, "#define __gcc") {
				a[w] = v0
				w++
			}
		}
		testPredef = strings.Join(a[:w], "\n")
	}
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

	if *oDownload {
		download()
	}

	path := filepath.Join(testWD, sqliteDir)
	if _, err := os.Stat(path); err != nil {
		log.Fatalf("Missing resources in %s. Please run 'go test -download' to fix.", path)
		return
	}

	path = filepath.Join(testWD, sqliteDir, "shell.c")
	if testShellSource, err = cache.getFile(newContext(cfg), path, false, false); err != nil {
		log.Fatal(err)
	}

	path = filepath.Join(testWD, sqliteDir, "sqlite3.c")
	if testSQLiteSource, err = cache.getFile(newContext(cfg), path, false, false); err != nil {
		log.Fatal(err)
	}
}

func download() {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	defer os.RemoveAll(tmp)

	for _, v := range downloads {
		dir := filepath.FromSlash(v.dir)
		root := filepath.Dir(v.dir)
		fi, err := os.Stat(dir)
		switch {
		case err == nil:
			if !fi.IsDir() {
				fmt.Fprintf(os.Stderr, "expected %s to be a directory\n", dir)
			}
			continue
		default:
			if !os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "%s", err)
				continue
			}

			if v.dev && !*oDev {
				fmt.Printf("Not downloading (no -dev) %v MB from %s\n", float64(v.sz)/1000, v.url)
				continue
			}

		}

		if err := func() error {
			fmt.Printf("Downloading %v MB from %s\n", float64(v.sz)/1000, v.url)
			resp, err := http.Get(v.url)
			if err != nil {
				return err
			}

			defer resp.Body.Close()

			base := filepath.Base(v.url)
			name := filepath.Join(tmp, base)
			f, err := os.Create(name)
			if err != nil {
				return err
			}

			defer os.Remove(name)

			n, err := io.Copy(f, resp.Body)
			if err != nil {
				return err
			}

			if _, err := f.Seek(0, io.SeekStart); err != nil {
				return err
			}

			switch {
			case strings.HasSuffix(base, ".tar.bz2"):
				b2r := bzip2.NewReader(bufio.NewReader(f))
				tr := tar.NewReader(b2r)
				for {
					hdr, err := tr.Next()
					if err != nil {
						if err != io.EOF {
							return err
						}

						return nil
					}

					switch hdr.Typeflag {
					case tar.TypeDir:
						if err = os.MkdirAll(filepath.Join(root, hdr.Name), 0770); err != nil {
							return err
						}
					case tar.TypeReg, tar.TypeRegA:
						f, err := os.OpenFile(filepath.Join(root, hdr.Name), os.O_CREATE|os.O_WRONLY, os.FileMode(hdr.Mode))
						if err != nil {
							return err
						}

						w := bufio.NewWriter(f)
						if _, err = io.Copy(w, tr); err != nil {
							return err
						}

						if err := w.Flush(); err != nil {
							return err
						}

						if err := f.Close(); err != nil {
							return err
						}
					default:
						return fmt.Errorf("unexpected tar header typeflag %#02x", hdr.Typeflag)
					}
				}
			case strings.HasSuffix(base, ".tar.gz"):
				gr, err := gzip.NewReader(bufio.NewReader(f))
				if err != nil {
					return err
				}

				tr := tar.NewReader(gr)
				for {
					hdr, err := tr.Next()
					if err != nil {
						if err != io.EOF {
							return err
						}

						return nil
					}

					switch hdr.Typeflag {
					case tar.TypeDir:
						if err = os.MkdirAll(filepath.Join(root, hdr.Name), 0770); err != nil {
							return err
						}
					case tar.TypeReg, tar.TypeRegA:
						f, err := os.OpenFile(filepath.Join(root, hdr.Name), os.O_CREATE|os.O_WRONLY, os.FileMode(hdr.Mode))
						if err != nil {
							return err
						}

						w := bufio.NewWriter(f)
						if _, err = io.Copy(w, tr); err != nil {
							return err
						}

						if err := w.Flush(); err != nil {
							return err
						}

						if err := f.Close(); err != nil {
							return err
						}
					default:
						return fmt.Errorf("unexpected tar header typeflag %#02x", hdr.Typeflag)
					}
				}
			case strings.HasSuffix(base, ".zip"):
				r, err := zip.NewReader(f, n)
				if err != nil {
					return err
				}

				for _, f := range r.File {
					fi := f.FileInfo()
					if fi.IsDir() {
						if err := os.MkdirAll(filepath.Join(root, f.Name), 0770); err != nil {
							return err
						}

						continue
					}

					if err := func() error {
						rc, err := f.Open()
						if err != nil {
							return err
						}

						defer rc.Close()

						dname := filepath.Join(root, f.Name)
						g, err := os.Create(dname)
						if err != nil {
							return err
						}

						defer g.Close()

						n, err = io.Copy(g, rc)
						return err
					}(); err != nil {
						return err
					}
				}
				return nil
			}
			panic(internalError())
		}(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
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
	sp := filepath.ToSlash(path)
	if strings.Contains(sp, "/.") {
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
