// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	"modernc.org/mathutil"
)

func TestTranslateSQLite(t *testing.T) {
	cfg := &Config{ABI: testABI}
	root := filepath.Join(testWD, filepath.FromSlash(sqliteDir))
	t.Run("shell.c", func(t *testing.T) { testTranslate(t, cfg, testPredef, filepath.Join(root, "shell.c")) })
	t.Run("shell.c/gnu", func(t *testing.T) { testTranslate(t, cfg, testPredefGNU, filepath.Join(root, "shell.c")) })
	t.Run("sqlite3.c", func(t *testing.T) { testTranslate(t, cfg, testPredef, filepath.Join(root, "sqlite3.c")) })
	t.Run("sqlite3.c/gnu", func(t *testing.T) { testTranslate(t, cfg, testPredefGNU, filepath.Join(root, "sqlite3.c")) })
}

var (
	benchmarkTranslateSQLiteAST *AST
	testTranslateSQLiteAST      *AST
)

func testTranslate(t *testing.T, cfg *Config, predef string, files ...string) {
	testTranslateSQLiteAST = nil
	sources := []Source{
		{Name: "<predefined>", Value: predef},
		{Name: "<built-in>", Value: parserTestBuiltin},
	}
	for _, v := range files {
		sources = append(sources, Source{Name: v})
	}
	ctx := newContext(cfg)
	var m0, m1 runtime.MemStats
	var err error
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m0)
	t0 := time.Now()
	if testTranslateSQLiteAST, err = translate(ctx, testIncludes, testSysIncludes, sources); err != nil {
		t.Error(err)
	}
	d := time.Since(t0)
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m1)
	t.Logf("sources %v, bytes %v, %v, %v B/s, mem %v",
		h(ctx.tuSources()), h(ctx.tuSize()), d, h(float64(time.Second)*float64(ctx.tuSize())/float64(d)), h(m1.Alloc-m0.Alloc))
}

func BenchmarkTranslateSQLite(b *testing.B) {
	cfg := &Config{ABI: testABI}
	root := filepath.Join(testWD, filepath.FromSlash(sqliteDir))
	b.Run("shell.c", func(b *testing.B) { benchmarkTranslateSQLite(b, cfg, testPredef, filepath.Join(root, "shell.c")) })
	b.Run("shell.c/gnu", func(b *testing.B) { benchmarkTranslateSQLite(b, cfg, testPredefGNU, filepath.Join(root, "shell.c")) })
	b.Run("sqlite3.c", func(b *testing.B) { benchmarkTranslateSQLite(b, cfg, testPredef, filepath.Join(root, "sqlite3.c")) })
	b.Run("sqlite3.c/gnu", func(b *testing.B) { benchmarkTranslateSQLite(b, cfg, testPredefGNU, filepath.Join(root, "sqlite3.c")) })
}

func benchmarkTranslateSQLite(b *testing.B, cfg *Config, predef string, files ...string) {
	sources := []Source{
		{Name: "<predefined>", Value: predef},
		{Name: "<built-in>", Value: parserTestBuiltin},
	}
	for _, v := range files {
		sources = append(sources, Source{Name: v})
	}
	ctx := newContext(cfg)
	// Warm up the cache
	if _, err := parse(ctx, testIncludes, testSysIncludes, sources); err != nil {
		b.Error(err)
		return
	}

	sz := ctx.tuSize()
	var err error
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if benchmarkTranslateSQLiteAST, err = Translate(cfg, testIncludes, testSysIncludes, sources); err != nil {
			b.Error(err)
		}
	}
	b.SetBytes(sz)
}

var (
	benchmarkTranslateAST *AST
	testTranslateAST      *AST
)

func TestTranslateTCC(t *testing.T) {
	if testing.Short() {
		t.Skip("-short")
		return
	}

	cfg := &Config{
		ABI:                        testABI,
		ignoreUndefinedIdentifiers: true,
	}
	root := filepath.Join(testWD, filepath.FromSlash(tccDir))
	if _, err := os.Stat(root); err != nil {
		t.Skipf("Missing resources in %s. Please run 'go test -download' to fix.", root)
	}

	ok := 0
	const dir = "tests/tests2"
	t.Run(dir, func(t *testing.T) {
		ok += testTranslateDir(t, cfg, testPredef, filepath.Join(root, filepath.FromSlash(dir)), false, false)
	})
	t.Run(dir+"/gnu", func(t *testing.T) {
		ok += testTranslateDir(t, cfg, testPredefGNU, filepath.Join(root, filepath.FromSlash(dir)), false, false)
	})
	t.Logf("ok %v", h(ok))
}

func TestTranslateGCC(t *testing.T) {
	if testing.Short() {
		t.Skip("-short")
		return
	}

	cfg := &Config{
		ABI:                        testABI,
		ignoreUndefinedIdentifiers: true,
	}
	root := filepath.Join(testWD, filepath.FromSlash(gccDir))
	if _, err := os.Stat(root); err != nil {
		t.Skipf("Missing resources in %s. Please run 'go test -download -dev' to fix.", root)
	}

	ok := 0
	for _, v := range []string{
		"gcc/testsuite/gcc.c-torture/compile",
		"gcc/testsuite/gcc.c-torture/execute",
	} {
		t.Run(v, func(t *testing.T) {
			ok += testTranslateDir(t, cfg, testPredef, filepath.Join(root, filepath.FromSlash(v)), true, false)
		})
		t.Run(v+"/gnu", func(t *testing.T) {
			ok += testTranslateDir(t, cfg, testPredefGNU, filepath.Join(root, filepath.FromSlash(v)), true, true)
		})
	}
	t.Logf("ok %v", h(ok))
}

func testTranslateDir(t *testing.T, cfg *Config, predef, dir string, hfiles, gnuc bool) (ok int) {
	blacklist := map[string]struct{}{ //TODO-
		// TCC
		"34_array_assignment.c":     {}, // gcc: 16:6: error: assignment to expression with array type
		"75_array_in_struct_init.c": {}, //TODO initializer missing braces

		// GCC
		"20000120-2.c":                 {}, //TODO function redefinition
		"20000804-1.c":                 {}, //TODO 1: unsupported type: complex long long
		"20020810-1.c":                 {}, //TODO :17:16: missing braces around initializer
		"20021118-1.c":                 {}, //TODO array initializer
		"20021120-1.c":                 {}, //TODO function redefinition
		"20021120-2.c":                 {}, //TODO function redefinition
		"20030305-1.c":                 {}, //TODO array initializer
		"20030903-1.c":                 {}, //TODO 41: unsupported type: complex int
		"20040726-2.c":                 {}, //TODO array initializer
		"20041124-1.c":                 {}, //TODO 12: unsupported type: complex short
		"20041201-1.c":                 {}, //TODO 18: unsupported type: complex char
		"20050121-1.c":                 {}, //TODO 1: unsupported type: complex char
		"20050122-2.c":                 {}, //TODO goto from nested function to outer function label
		"20050215-1.c":                 {}, //TODO function redefinition
		"20050215-2.c":                 {}, //TODO function redefinition
		"20050215-3.c":                 {}, //TODO function redefinition
		"20070919-1.c":                 {}, //TODO :39:9: missing braces around initializer
		"20180921-1.c":                 {}, //TODO :129:27: missing braces around initializer
		"920415-1.c":                   {}, //TODO label l undefined
		"920428-2.c":                   {}, //TODO goto from nested function to outer function label
		"920428-4.c":                   {}, //TODO invalid declarator type
		"920501-16.c":                  {}, //TODO invalid declarator type
		"920501-7.c":                   {}, //TODO goto from nested function to outer function label
		"920611-2.c":                   {}, //TODO array initializer
		"920721-4.c":                   {}, //TODO label default_lab undefined
		"921017-1.c":                   {}, //TODO invalid declarator type
		"930510-1.c":                   {}, //TODO array initializer
		"991201-1.c":                   {}, //TODO :11:27: missing braces around initializer
		"builtin-types-compatible-p.c": {}, //TODO invalid declarator type
		"comp-goto-2.c":                {}, //TODO goto from nested function to outer function label
		"complex-1.c":                  {}, //TODO 9: unsupported type: complex int
		"complex-5.c":                  {}, //TODO 9: unsupported type: complex int
		"complex-6.c":                  {}, //TODO 3: unsupported type: complex int
		"limits-externdecl.c":          {}, //TODO :57:3: missing braces around initializer
		"nestfunc-5.c":                 {}, //TODO goto from nested function to outer function label
		"nestfunc-6.c":                 {}, //TODO goto from nested function to outer function label
		"pr21728.c":                    {}, //TODO goto from nested function to outer function label
		"pr24135.c":                    {}, //TODO goto from nested function to outer function label
		"pr27889.c":                    {}, //TODO 1: unsupported type: complex int
		"pr33631.c":                    {}, //TODO :10:53: missing braces around initializer
		"pr35431.c":                    {}, //TODO 3: unsupported type: complex int
		"pr38151.c":                    {}, //TODO 3: unsupported type: complex int
		"pr41987.c":                    {}, //TODO 3: unsupported type: complex char
		"pr42196-1.c":                  {}, //TODO 3: unsupported type: complex int
		"pr42196-2.c":                  {}, //TODO 3: unsupported type: complex int
		"pr42196-3.c":                  {}, //TODO 3: unsupported type: complex int
		"pr43191.c":                    {}, //TODO array initializer
		"pr48517.c":                    {}, //TODO array initializer
		"pr49218.c":                    {}, //TODO :11:9: missing braces around initializer
		"pr51447.c":                    {}, //TODO goto from nested function to outer function label
		"pr54471.c":                    {}, //TODO :15:22: missing braces around initializer
		"pr56448.c":                    {}, // Decimal64 literals
		"pr56837.c":                    {}, //TODO 1: unsupported type: complex int
		"pr61375.c":                    {}, //TODO :17:19: missing braces around initializer
		"pr63302.c":                    {}, //TODO :16:16: missing braces around initializer
		"pr64067.c":                    {}, //TODO & of composite literal
		"pr80692.c":                    {}, // Decimal64 literals
		"pr85582-2.c":                  {}, //TODO :38:9: missing braces around initializer
		"pr85582-3.c":                  {}, //TODO :32:9: missing braces around initializer
		"pr86122.c":                    {}, //TODO 1: unsupported type: complex int
		"pr86123.c":                    {}, //TODO 6: unsupported type: complex unsigned
		"pr87647.c":                    {}, //TODO :11:20: missing braces around initializer
		"pr89369.c":                    {}, //TODO array initializer
		"struct-ini-1.c":               {}, //TODO array initializer

	}
	if !gnuc { // vector extensions
		for k, v := range map[string]struct{}{
			"20050113-1.c":   {},
			"20050316-1.c":   {},
			"20050316-2.c":   {},
			"20050316-3.c":   {},
			"20050604-1.c":   {},
			"20050607-1.c":   {},
			"icfmatch.c":     {},
			"pr23135.c":      {},
			"pr33614.c":      {},
			"pr33617.c":      {},
			"pr52750.c":      {},
			"pr53410-2.c":    {},
			"pr53645-2.c":    {},
			"pr53645.c":      {},
			"pr53748.c":      {},
			"pr54713-1.c":    {},
			"pr54713-2.c":    {},
			"pr54713-3.c":    {},
			"pr60502.c":      {},
			"pr60960.c":      {},
			"pr65427.c":      {},
			"pr70061.c":      {},
			"pr70240.c":      {},
			"pr70355.c":      {},
			"pr70633.c":      {},
			"pr70903.c":      {},
			"pr71626-1.c":    {},
			"pr71626-2.c":    {},
			"pr72824-2.c":    {},
			"pr85169.c":      {},
			"pr85331.c":      {},
			"scal-to-vec1.c": {},
			"scal-to-vec2.c": {},
			"scal-to-vec3.c": {},
			"simd-1.c":       {},
			"simd-2.c":       {},
			"simd-3.c":       {},
			"simd-4.c":       {},
			"simd-5.c":       {},
			"simd-6.c":       {},
			"vector-3.c":     {},
			"vector-5.c":     {},
			"vector-6.c":     {},
		} {
			blacklist[k] = v
		}
	}
	var re *regexp.Regexp
	if s := *oRE; s != "" {
		re = regexp.MustCompile(s)
	}

	var files, psources int
	var bytes int64
	var m0, m1 runtime.MemStats
	testTranslateAST = nil
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m0)
	t0 := time.Now()
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if os.IsNotExist(err) {
				err = nil
			}
			return err
		}

		if info.IsDir() {
			return skipDir(path)
		}

		if filepath.Ext(path) != ".c" && (!hfiles || filepath.Ext(path) != ".h") || info.Mode()&os.ModeType != 0 {
			return nil
		}

		if _, ok := blacklist[filepath.Base(path)]; ok {
			return nil
		}

		files++

		if re != nil && !re.MatchString(path) {
			return nil
		}

		sources := []Source{
			{Name: "<predefined>", Value: predef},
			{Name: "<built-in>", Value: parserTestBuiltin},
			{Name: path},
		}
		ctx := newContext(cfg)

		defer func() {
			psources += ctx.tuSources()
			bytes += ctx.tuSize()
		}()

		if *oTrace {
			fmt.Fprintln(os.Stderr, files, path)
		}
		if testTranslateAST, err = parse(ctx, testIncludes, testSysIncludes, sources); err != nil {
			if gnuc {
				t.Error(err)
			}
			return nil
		}

		if err = testTranslateAST.Typecheck(); err != nil {
			t.Error(err)
			return nil
		}

		ok++
		return nil
	}); err != nil {
		t.Error(err)
	}
	d := time.Since(t0)
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m1)
	t.Logf("files %v, sources %v, bytes %v, ok %v, %v, %v B/s, mem %v",
		h(files), h(psources), h(bytes), h(ok), d, h(float64(time.Second)*float64(bytes)/float64(d)), h(m1.Alloc-m0.Alloc))
	if files != ok && gnuc {
		t.Errorf("files %v, bytes %v, ok %v", files, bytes, ok)
	}
	return ok
}

func BenchmarkTranslateTCC(b *testing.B) {
	root := filepath.Join(testWD, filepath.FromSlash(tccDir))
	if _, err := os.Stat(root); err != nil {
		b.Skipf("Missing resources in %s. Please run 'go test -download' to fix.", root)
	}

	cfg := &Config{
		ABI:                        testABI,
		ignoreUndefinedIdentifiers: true,
	}
	const dir = "tests/tests2"
	b.Run(dir, func(b *testing.B) {
		benchmarkTranslateDir(b, cfg, testPredef, filepath.Join(root, filepath.FromSlash(dir)), false)
	})
	b.Run(dir+"/gnu", func(b *testing.B) {
		benchmarkTranslateDir(b, cfg, testPredefGNU, filepath.Join(root, filepath.FromSlash(dir)), false)
	})
}

func BenchmarkTranslateGCC(b *testing.B) {
	root := filepath.Join(testWD, filepath.FromSlash(gccDir))
	if _, err := os.Stat(root); err != nil {
		b.Skipf("Missing resources in %s. Please run 'go test -download -dev' to fix.", root)
	}

	cfg := &Config{
		ABI:                        testABI,
		ignoreUndefinedIdentifiers: true,
	}
	for _, v := range []string{
		"gcc/testsuite/gcc.c-torture/compile",
		"gcc/testsuite/gcc.c-torture/execute",
	} {
		b.Run(v+"/gnu", func(b *testing.B) {
			benchmarkTranslateDir(b, cfg, testPredefGNU, filepath.Join(root, filepath.FromSlash(v)), true)
		})
	}
}

func benchmarkTranslateDir(b *testing.B, cfg *Config, predef, dir string, must bool) {
	blacklist := map[string]struct{}{ //TODO-
		// TCC
		"13_integer_literals.c":        {},
		"70_floating_point_literals.c": {},

		// GCC/exec
		"pr80692.c": {}, // Decimal64 literals
	}
	var bytes int64
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes = 0
		if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				if os.IsNotExist(err) {
					err = nil
				}
				return err
			}

			if info.IsDir() {
				return skipDir(path)
			}

			if filepath.Ext(path) != ".c" && filepath.Ext(path) != ".h" || info.Mode()&os.ModeType != 0 {
				return nil
			}

			if _, ok := blacklist[filepath.Base(path)]; ok {
				return nil
			}

			sources := []Source{
				{Name: "<predefined>", Value: predef},
				{Name: "<built-in>", Value: parserTestBuiltin},
				{Name: path},
			}
			ctx := newContext(cfg)
			if benchmarkTranslateAST, err = parse(ctx, testIncludes, testSysIncludes, sources); err != nil {
				if must {
					b.Error(err)
				}
				return nil
			}

			if err = benchmarkTranslateAST.Typecheck(); err != nil {
				if must {
					b.Error(err)
				}
				return nil
			}
			bytes += ctx.tuSize()
			return nil
		}); err != nil {
			b.Error(err)
		}
	}
	b.SetBytes(bytes)
}

func TestAbstractDeclarator(t *testing.T) { //TODO -> Example
	for i, test := range []struct{ src, typ string }{
		{"int i = sizeof(int);", "int"},                                                                                            // [0], 6.7.6, 3, (a)
		{"int i = sizeof(int*);", "pointer to int"},                                                                                // [0], 6.7.6, 3, (b)
		{"int i = sizeof(int*[3]);", "array of 3 pointer to int"},                                                                  // [0], 6.7.6, 3, (c)
		{"int i = sizeof(int(*)[3]);", "pointer to array of 3 int"},                                                                // [0], 6.7.6, 3, (d)
		{"int i = sizeof(int(*)[*]);", "pointer to array of int"},                                                                  // [0], 6.7.6, 3, (e)
		{"int i = sizeof(int *());", "function() returning pointer to int"},                                                        // [0], 6.7.6, 3, (f)
		{"int i = sizeof(int (*)(void));", "pointer to function(void) returning int"},                                              // [0], 6.7.6, 3, (g)
		{"int i = sizeof(int (*[])(unsigned int, ...));", "array of pointer to function(unsigned, ...) returning int"},             // [0], 6.7.6, 3, (h)
		{"int i = sizeof(int (*const [])(unsigned int, ...));", "array of const pointer to function(unsigned, ...) returning int"}, // [0], 6.7.6, 3, (h)
	} {
		letter := string(rune('a' + i))
		cfg := &Config{ABI: testABI, doNotSanityCheckComplexTypes: true}
		ast, err := Translate(cfg, nil, nil, []Source{
			{Name: "<built-in>", Value: "typedef long long unsigned size_t;"},
			{Name: "test", Value: test.src},
		})
		if err != nil {
			t.Errorf("(%v): %v", letter, err)
			continue
		}

		var node Node
		depth := mathutil.MaxInt
		findNode("TypeName", ast.TranslationUnit, 0, &node, &depth)
		if node == nil {
			t.Errorf("(%v): %q, TypeName node not found", letter, test.src)
			continue
		}

		g, e := node.(*TypeName).Type().String(), test.typ
		if *oTrace {
			t.Logf("(%v): %q, %q", letter, test.src, g)
		}
		if g != e {
			t.Errorf("(%v): %q, got %q, expected %q", letter, test.src, g, e)
		}
	}
}

func TestAbstractDeclarator2(t *testing.T) { //TODO -> Example
	for i, test := range []struct{ src, typ string }{
		{"void f(int);", "int"},                                                                                            // [0], 6.7.6, 3, (a)
		{"void f(int*);", "pointer to int"},                                                                                // [0], 6.7.6, 3, (b)
		{"void f(int*[3]);", "array of 3 pointer to int"},                                                                  // [0], 6.7.6, 3, (c)
		{"void f(int(*)[3]);", "pointer to array of 3 int"},                                                                // [0], 6.7.6, 3, (d)
		{"void f(int(*)[*]);", "pointer to array of int"},                                                                  // [0], 6.7.6, 3, (e)
		{"void f(int *());", "function() returning pointer to int"},                                                        // [0], 6.7.6, 3, (f)
		{"void f(int (*)(void));", "pointer to function(void) returning int"},                                              // [0], 6.7.6, 3, (g)
		{"void f(int (*[])(unsigned int, ...));", "array of pointer to function(unsigned, ...) returning int"},             // [0], 6.7.6, 3, (h)
		{"void f(int (*const [])(unsigned int, ...));", "array of const pointer to function(unsigned, ...) returning int"}, // [0], 6.7.6, 3, (h)
	} {
		letter := string(rune('a' + i))
		cfg := &Config{ABI: testABI, doNotSanityCheckComplexTypes: true}
		ast, err := Translate(cfg, nil, nil, []Source{{Name: "test", Value: test.src}})
		if err != nil {
			t.Errorf("(%v): %v", letter, err)
			continue
		}

		var node Node
		depth := mathutil.MaxInt
		findNode("ParameterDeclaration", ast.TranslationUnit, 0, &node, &depth)
		if node == nil {
			t.Errorf("(%v): %q, ParameterDeclaration node not found", letter, test.src)
			continue
		}

		g, e := node.(*ParameterDeclaration).Type().String(), test.typ
		if *oTrace {
			t.Logf("(%v): %q, %q", letter, test.src, g)
		}
		if g != e {
			t.Errorf("(%v): %q, got %q, expected %q", letter, test.src, g, e)
		}
	}
}

func TestDeclarator(t *testing.T) { //TODO -> Example
	for i, test := range []struct{ src, typ string }{
		{"int x;", "int"},                                                                                           // [0], 6.7.6, 3, (a)
		{"int *x;", "pointer to int"},                                                                               // [0], 6.7.6, 3, (b)
		{"int *x[3];", "array of 3 pointer to int"},                                                                 // [0], 6.7.6, 3, (c)
		{"int (*x)[3];", "pointer to array of 3 int"},                                                               // [0], 6.7.6, 3, (d)
		{"int (*x)[*];", "pointer to array of int"},                                                                 // [0], 6.7.6, 3, (e)
		{"int *x();", "function() returning pointer to int"},                                                        // [0], 6.7.6, 3, (f)
		{"int (*x)(void);", "pointer to function(void) returning int"},                                              // [0], 6.7.6, 3, (g)
		{"int (*x[])(unsigned int, ...);", "array of pointer to function(unsigned, ...) returning int"},             // [0], 6.7.6, 3, (h)
		{"int (*const x[])(unsigned int, ...);", "array of const pointer to function(unsigned, ...) returning int"}, // [0], 6.7.6, 3, (h)
	} {
		letter := string(rune('a' + i))
		cfg := &Config{ABI: testABI, doNotSanityCheckComplexTypes: true}
		ast, err := Translate(cfg, nil, nil, []Source{{Name: "test", Value: test.src}})
		if err != nil {
			t.Errorf("(%v): %v", letter, err)
			continue
		}

		var node Node
		depth := mathutil.MaxInt
		findNode("Declarator", ast.TranslationUnit, 0, &node, &depth)
		if node == nil {
			t.Errorf("(%v): %q, Declarator node not found", letter, test.src)
			continue
		}

		g, e := node.(*Declarator).Type().String(), test.typ
		if *oTrace {
			t.Logf("(%v): %q, %q", letter, test.src, g)
		}
		if g != e {
			t.Errorf("(%v): %q, got %q, expected %q", letter, test.src, g, e)
		}
	}
}

func TestDeclarator2(t *testing.T) {
	for i, test := range []struct{ src, typ string }{
		{"struct { int i; } s;", "struct {i int; }"},                                                                       // (a)
		{"union { int i; double d; } u;", "union {i int; d double; }"},                                                     // (b)
		{"typedef struct { unsigned char __c[8]; double __d; } s;", "struct {__c array of 8 unsigned char; __d double; }"}, // (c)
		{"typedef union { unsigned char __c[8]; double __d; } u;", "union {__c array of 8 unsigned char; __d double; }"},   // (d)
		{"struct s { int i;}; typeof(struct s) x;", "struct s"},                                                            // (e)
		{"typeof(42) x;", "int"},       // (f)
		{"typeof(42L) x;", "long"},     // (g)
		{"typeof(42U) x;", "unsigned"}, // (h)
		{"typeof(42.) x;", "double"},   // (i)
		{"#define __GNUC__\ntypedef int x __attribute__ ((vector_size (16)));", "vector of 4 int"}, // (j)
	} {
		letter := string(rune('a' + i))
		cfg := &Config{ABI: testABI, doNotSanityCheckComplexTypes: true}
		ast, err := Translate(cfg, nil, nil, []Source{{Name: "test", Value: test.src}})
		if err != nil {
			t.Errorf("(%v): %v", letter, err)
			continue
		}

		var node Node
		depth := mathutil.MaxInt
		findNode("Declarator", ast.TranslationUnit, 0, &node, &depth)
		if node == nil {
			t.Errorf("(%v): %q, Declarator node not found", letter, test.src)
			continue
		}

		g, e := node.(*Declarator).Type().String(), test.typ
		if *oTrace {
			t.Logf("(%v): %q, %q", letter, test.src, g)
		}
		if g != e {
			t.Errorf("(%v): %q, got %q, expected %q", letter, test.src, g, e)
		}
	}
}

func TestTranslateCSmith(t *testing.T) {
	if testing.Short() {
		t.Skip("-short")
		return
	}

	csmith, err := exec.LookPath("csmith")
	if err != nil {
		t.Logf("%v: skipping test", err)
		return
	}

	fixedBugs := []string{}
	ch := time.After(*oCSmith)
	t0 := time.Now()
	var files, ok int
	var size int64
out:
	for i := 0; ; i++ {
		extra := ""
		var args string
		switch {
		case i < len(fixedBugs):
			args += fixedBugs[i]
			a := strings.Split(fixedBugs[i], " ")
			extra = strings.Join(a[len(a)-2:], " ")
		default:
			select {
			case <-ch:
				break out
			default:
			}

			args += csmithArgs
		}
		out, err := exec.Command(csmith, strings.Split(args, " ")...).Output()
		if err != nil {
			t.Fatalf("%v\n%s", err, out)
		}

		if fn := *oBlackBox; fn != "" {
			if err := ioutil.WriteFile(fn, out, 0660); err != nil {
				t.Fatal(err)
			}
		}

		cfg := &Config{ABI: testABI, Config3: Config3{MaxSourceLine: 1 << 20}}
		ctx := newContext(cfg)
		files++
		size += int64(len(out))
		sources := []Source{
			{Name: "<predefined>", Value: testPredef},
			{Name: "<built-in>", Value: parserTestBuiltin},
			{Name: "test.c", Value: string(out), DoNotCache: true},
		}
		if err := func() (err error) {
			defer func() {
				if e := recover(); e != nil && err == nil {
					err = fmt.Errorf("%v", e)
				}
			}()

			var ast *AST
			if ast, err = parse(ctx, testIncludes, testSysIncludes, sources); err != nil {
				return err
			}

			err = ast.Typecheck()
			return err

		}(); err != nil {
			t.Errorf("%s\n%s\nTypecheck: %v", extra, out, err)
			break
		}

		ok++
		if *oTrace {
			fmt.Fprintln(os.Stderr, time.Since(t0), files, ok)
		}
	}
	d := time.Since(t0)
	t.Logf("files %v, bytes %v, ok %v in %v", h(files), h(size), h(ok), d)
}
