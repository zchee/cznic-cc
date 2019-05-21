// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestBOM(t *testing.T) {
	for i, v := range []struct {
		src string
		err string
	}{
		{"int main() {}", ""},
		{"\xEF\xBB\xBFint main() {}", ""},
	} {
		switch _, err := Parse(&Config{}, nil, nil, []Source{{Value: v.src}}); {
		case v.err == "" && err != nil:
			t.Errorf("%v: unexpected error %v", i, err)
		case v.err != "" && err == nil:
			t.Errorf("%v: unexpected success, expected error matching %v", i, v.err)
		case v.err != "":
			if !regexp.MustCompile(v.err).MatchString(err.Error()) {
				t.Errorf("%v: error %v does not match %v", i, err, v.err)
			}
		}
	}
}

func TestStrCatSep(t *testing.T) {
	cfg := &Config{Config3: Config3{PreserveWhiteSpace: true}}
	for i, v := range []struct {
		src         string
		lit         string
		sep         string
		trailingSep string
	}{
		{`int f() {  "a";}`, "a", "  ", "\n"},
		{`int f() { "a" "b";}`, "ab", "  ", "\n"},
		{`int f() { "a""b";}`, "ab", " ", "\n"},
		{`int f() { "a"` + "\n\t" + `"b"; }`, "ab", " \n\t", "\n"},
		{`int f() { "a";}`, "a", " ", "\n"},
		{`int f() { /*x*/ /*y*/ "a";}`, "a", " /*x*/ /*y*/ ", "\n"},
		{`int f() { /*x*/` + "\n" + `/*y*/ "a";}`, "a", " /*x*/\n/*y*/ ", "\n"},
		{`int f() { //x` + "\n" + ` "a";}`, "a", " //x\n ", "\n"},
		{`int f() { //x` + "\n" + `"a";}`, "a", " //x\n", "\n"},
		{`int f() { ` + "\n" + ` "a";}`, "a", " \n ", "\n"},
		{`int f() { ` + "\n" + `"a";}`, "a", " \n", "\n"},
		{`int f() {"a" "b";}`, "ab", " ", "\n"},
		{`int f() {"a"/*y*/"b";}`, "ab", "/*y*/", "\n"},
		{`int f() {"a";} /*x*/ `, "a", "", " /*x*/ \n"},
		{`int f() {"a";} /*x*/`, "a", "", " /*x*/\n"},
		{`int f() {"a";} /*x` + "\n" + `*/ `, "a", "", " /*x\n*/ \n"},
		{`int f() {"a";} `, "a", "", " \n"},
		{`int f() {"a";}/*x*/`, "a", "", "/*x*/\n"},
		{`int f() {"a";}` + "\n", "a", "", "\n"},
		{`int f() {"a";}`, "a", "", "\n"},
		{`int f() {/*x*/ /*y*/ "a";}`, "a", "/*x*/ /*y*/ ", "\n"},
		{`int f() {/*x*/"a""b";}`, "ab", "/*x*/", "\n"},
		{`int f() {/*x*/"a"/*y*/"b";}`, "ab", "/*x*//*y*/", "\n"},
		{`int f() {/*x*/"a";}`, "a", "/*x*/", "\n"},
		{`int f() {/*x*//*y*/ "a";}`, "a", "/*x*//*y*/ ", "\n"},
		{`int f() {/*x*//*y*/"a";}`, "a", "/*x*//*y*/", "\n"},
		{`int f() {//` + "\n" + `"a";}`, "a", "//\n", "\n"},
		{`int f() {//x` + "\n" + `"a";}`, "a", "//x\n", "\n"},
		{`int f() {` + "\n" + ` "a";}`, "a", "\n ", "\n"},
		{`int f() {` + "\n" + `"a";}`, "a", "\n", "\n"},
	} {
		ast, err := Parse(cfg, nil, nil, []Source{{Name: "test", Value: v.src}})
		if err != nil {
			t.Errorf("%v: %v", i, err)
			continue
		}

		tok := ast.
			TranslationUnit.
			ExternalDeclaration.
			FunctionDefinition.
			CompoundStatement.
			BlockItemList.
			BlockItem.
			Statement.
			ExpressionStatement.
			Expression.
			AssignmentExpression.
			ConditionalExpression.
			LogicalOrExpression.
			LogicalAndExpression.
			InclusiveOrExpression.
			ExclusiveOrExpression.
			AndExpression.
			EqualityExpression.
			RelationalExpression.
			ShiftExpression.
			AdditiveExpression.
			MultiplicativeExpression.
			CastExpression.
			UnaryExpression.
			PostfixExpression.
			PrimaryExpression.
			Token
		if g, e := tok.String(), v.lit; g != e {
			t.Errorf("%v: %q %q", i, g, e)
		}
		if g, e := tok.Sep.String(), v.sep; g != e {
			t.Errorf("%v: %q %q", i, g, e)
		}
		if g, e := ast.TrailingSeperator.String(), v.trailingSep; g != e {
			t.Errorf("%v: %q %q", i, g, e)
		}
	}
}

func TestParseJhjourdan(t *testing.T) {
	mustFail := map[string]struct{}{
		"dangling_else_misleading.fail.c": {},
		"atomic_parenthesis.c":            {}, // See [3], 3.5, pg. 20.
		//TODO Typechecking needed to actually fail "bitfield_declaration_ambiguity.fail.c":  {},
	}
	var re *regexp.Regexp
	if s := *oRE; s != "" {
		re = regexp.MustCompile(s)
	}

	cfg := &Config{
		ignoreIncludes:             true,
		ignoreUndefinedIdentifiers: true,
	}
	var ok, n int
	if err := filepath.Walk(filepath.Join(testWD, filepath.FromSlash("testdata/jhjourdan/")), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".c") {
			return nil
		}

		if re != nil && !re.MatchString(path) {
			return nil
		}

		n++
		_, expectFail := mustFail[filepath.Base(path)]
		switch _, err := Parse(cfg, nil, nil, []Source{{Name: path}}); {
		case err != nil:
			if !expectFail {
				t.Errorf("FAIL: unexpected error: %v", err)
				return nil
			}
		default:
			if expectFail {
				t.Errorf("FAIL: %v: unexpected success", path)
				return nil
			}
		}

		ok++
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	t.Logf("jhjourdan parse: ok %v/%v\n", ok, n)
}

func TestParseSQLite(t *testing.T) {
	cfg := &Config{}
	root := filepath.Join(testWD, filepath.FromSlash(sqliteDir))
	t.Run("shell.c", func(t *testing.T) { testParse(t, cfg, testPredef, filepath.Join(root, "shell.c")) })
	t.Run("shell.c/gnu", func(t *testing.T) { testParse(t, cfg, testPredefGNU, filepath.Join(root, "shell.c")) })
	t.Run("sqlite3.c", func(t *testing.T) { testParse(t, cfg, testPredef, filepath.Join(root, "sqlite3.c")) })
	t.Run("sqlite3.c/gnu", func(t *testing.T) { testParse(t, cfg, testPredefGNU, filepath.Join(root, "sqlite3.c")) })
}

var testParseAST *AST

func testParse(t *testing.T, cfg *Config, predef string, files ...string) {
	testParseAST = nil
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
	if testParseAST, err = parse(ctx, testIncludes, testSysIncludes, sources); err != nil {
		t.Error(err)
	}
	d := time.Since(t0)
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m1)
	t.Logf("sources %v, bytes %v, %v, %v B/s, mem %v",
		h(ctx.tuSources), h(ctx.tuSize), d, h(float64(time.Second)*float64(ctx.tuSize)/float64(d)), h(m1.Alloc-m0.Alloc))
}

func BenchmarkParseSQLite(b *testing.B) {
	cfg := &Config{}
	root := filepath.Join(testWD, filepath.FromSlash(sqliteDir))
	b.Run("shell.c", func(b *testing.B) { benchmarkParseSQLite(b, cfg, testPredef, filepath.Join(root, "shell.c")) })
	b.Run("shell.c/gnu", func(b *testing.B) { benchmarkParseSQLite(b, cfg, testPredefGNU, filepath.Join(root, "shell.c")) })
	b.Run("sqlite3.c", func(b *testing.B) { benchmarkParseSQLite(b, cfg, testPredef, filepath.Join(root, "sqlite3.c")) })
	b.Run("sqlite3.c/gnu", func(b *testing.B) { benchmarkParseSQLite(b, cfg, testPredefGNU, filepath.Join(root, "sqlite3.c")) })
}

func benchmarkParseSQLite(b *testing.B, cfg *Config, predef string, files ...string) {
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

	sz := ctx.tuSize
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Parse(cfg, testIncludes, testSysIncludes, sources); err != nil {
			b.Error(err)
			return
		}
	}
	b.SetBytes(sz)
}

// jnml@4670:~/src/modernc.org/cc/v3$ date ; go test -timeout 24h -v -dev -run DevParse -maxFiles -1 | tee log
// Sun Apr 14 22:57:51 CEST 2019
// === RUN   TestDevParse
// === RUN   TestDevParse/.c/gnu
// --- PASS: TestDevParse (304.52s)
//     --- PASS: TestDevParse/.c/gnu (304.52s)
//         ---- pass at least 1000 files
//           5678/5713   99.39% gcc-8.3.0/gcc/testsuite/gcc.target/i386
//           4027/4293   93.80% gcc-8.3.0/gcc/testsuite/gcc.dg
//           1759/1759  100.00% gcc-8.3.0/gcc/testsuite/gcc.c-torture/compile
//           1558/1564   99.62% gcc-8.3.0/gcc/testsuite/gcc.dg/tree-ssa
//           1475/1475  100.00% gcc-8.3.0/gcc/testsuite/gcc.c-torture/execute
//           1076/1094   98.35% gcc-8.3.0/gcc/testsuite/gcc.dg/torture
//           1041/1041  100.00% gcc-8.3.0/gcc/testsuite/gcc.dg/vect
//         files 32,435, sources 932,447, bytes 11,434,304,333, ok 26,123, 5m3.674664225s, 37,653,138 B/s, mem 2,350,780,288
// PASS
// ok  	modernc.org/cc/v3	304.923s
// jnml@4670:~/src/modernc.org/cc/v3$

// ==== jnml@e5-1650:~/src/modernc.org/cc/v3> date |& tee log ; go test -timeout 24h -v -dev -run DevParse -maxFiles -1 |& tee -a log
// Wed Apr 10 10:01:57 CEST 2019
// === RUN   TestDevParse
// === RUN   TestDevParse/.c
// === RUN   TestDevParse/.c/gnu
// --- PASS: TestDevParse (638.45s)
//     --- PASS: TestDevParse/.c (316.75s)
//         ---- pass at least 1000 files
//           5684/5713   99.49% gcc-8.3.0/gcc/testsuite/gcc.target/i386
//           4025/4293   93.76% gcc-8.3.0/gcc/testsuite/gcc.dg
//           1759/1759  100.00% gcc-8.3.0/gcc/testsuite/gcc.c-torture/compile
//           1560/1564   99.74% gcc-8.3.0/gcc/testsuite/gcc.dg/tree-ssa
//           1475/1475  100.00% gcc-8.3.0/gcc/testsuite/gcc.c-torture/execute
//           1041/1041  100.00% gcc-8.3.0/gcc/testsuite/gcc.dg/vect
//           1040/1094   95.06% gcc-8.3.0/gcc/testsuite/gcc.dg/torture
//         files 32,435, sources 916,514, bytes 11,159,198,946, ok 25,972, 5m15.837032153s, 35,332,142 B/s, mem 2,348,981,880
//     --- PASS: TestDevParse/.c/gnu (321.70s)
//         ---- pass at least 1000 files
//           5684/5713   99.49% gcc-8.3.0/gcc/testsuite/gcc.target/i386
//           4030/4293   93.87% gcc-8.3.0/gcc/testsuite/gcc.dg
//           1759/1759  100.00% gcc-8.3.0/gcc/testsuite/gcc.c-torture/compile
//           1560/1564   99.74% gcc-8.3.0/gcc/testsuite/gcc.dg/tree-ssa
//           1475/1475  100.00% gcc-8.3.0/gcc/testsuite/gcc.c-torture/execute
//           1041/1041  100.00% gcc-8.3.0/gcc/testsuite/gcc.dg/vect
//           1040/1094   95.06% gcc-8.3.0/gcc/testsuite/gcc.dg/torture
//         files 32,435, sources 908,733, bytes 11,154,082,341, ok 26,123, 5m19.950910638s, 34,861,855 B/s, mem 23,569,920
// PASS
// ok  	modernc.org/cc/v3	638.916s
// ==== jnml@e5-1650:~/src/modernc.org/cc/v3>

func TestDevParse(t *testing.T) {
	if !*oDev {
		t.Skip("-dev to enable")
		return
	}

	if testing.Short() {
		t.Skip("-short")
		return
	}

	t.Run(".c/gnu", func(t *testing.T) {
		testDevParse(t, testPredefGNU, func(s string) bool { return filepath.Ext(s) == ".c" }, 1000)
	})
}

func testDevParse(t *testing.T, predef string, filter func(string) bool, min int) {
	var re *regexp.Regexp
	if x := *oRE; x != "" {
		re = regexp.MustCompile(x)
	}

	cfg := &Config{
		ignoreUndefinedIdentifiers: true,
	}
	m := map[string]int{}
	n := map[string]int{}
	limit := *oMaxFiles
	var files, psources, ok int
	var bytes int64
	strip := len(*oWalkDir) + 1
	var m0, m1 runtime.MemStats
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m0)
	t0 := time.Now()
	if err := filepath.Walk(*oWalkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return skipDir(path)
		}

		if !filter(path) || info.Mode()&os.ModeType != 0 {
			return nil
		}

		if re != nil && !re.MatchString(path) {
			return nil
		}

		if limit == 0 {
			return nil
		}

		limit--
		files++
		key := filepath.Dir(path)[strip:]
		n[key]++
		sources := []Source{
			{Name: "<predefined>", Value: predef},
			{Name: "<built-in>", Value: parserTestBuiltin},
			{Name: path},
		}
		ctx := newContext(cfg)
		if *oTrace {
			fmt.Fprintln(os.Stderr, files, path)
		}
		if _, err := parse(ctx, testIncludes, testSysIncludes, sources); err == nil {
			ok++
			m[key]++
		}
		psources += ctx.tuSources
		bytes += ctx.tuSize
		return nil
	}); err != nil {
		t.Error(err)
	}
	d := time.Since(t0)
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m1)
	var a []string
	for k, v := range m {
		if v >= min {
			a = append(a, fmt.Sprintf("%6d/%-5d %6.2f%% %s", v, n[k], 100*float64(v)/float64(n[k]), k))
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(a)))
	if len(a) != 0 {
		t.Logf("---- pass at least %v files", min)
		for _, v := range a {
			t.Log(v)
		}
	}
	t.Logf("files %v, sources %v, bytes %v, ok %v, %v, %v B/s, mem %v",
		h(files), h(psources), h(bytes), h(ok), d, h(float64(time.Second)*float64(bytes)/float64(d)), h(m1.Alloc-m0.Alloc))
}

func BenchmarkDevParse(b *testing.B) {
	if !*oDev {
		b.Skip("-dev to enable")
		return
	}

	b.Run(".c/gnu", func(b *testing.B) { benchmarkDevParse(b, testPredefGNU) })
}

func benchmarkDevParse(b *testing.B, predef string) {
	var re *regexp.Regexp
	if x := *oRE; x != "" {
		re = regexp.MustCompile(x)
	}

	cfg := &Config{}
	limit := *oMaxFiles
	var bytes int64
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes = 0
		if err := filepath.Walk(*oWalkDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return skipDir(path)
			}

			if filepath.Ext(path) != ".c" || info.Mode()&os.ModeType != 0 {
				return nil
			}

			if re != nil && !re.MatchString(path) {
				return nil
			}

			if limit == 0 {
				return nil
			}

			limit--
			sources := []Source{
				{Name: "<predefined>", Value: predef},
				{Name: "<built-in>", Value: parserTestBuiltin},
				{Name: path},
			}
			ctx := newContext(cfg)
			parse(ctx, testIncludes, testSysIncludes, sources)
			bytes += ctx.tuSize
			return nil
		}); err != nil {
			b.Fatal(err)
		}
	}
	b.SetBytes(bytes)
}

func TestParseTCC(t *testing.T) {
	cfg := &Config{
		ignoreUndefinedIdentifiers: true,
	}
	root := filepath.Join(testWD, filepath.FromSlash(tccDir))
	if _, err := os.Stat(root); err != nil {
		t.Skipf("Missing resources in %s. Please run 'go test -download' to fix.", root)
	}

	ok := 0
	const dir = "tests/tests2"
	t.Run(dir, func(t *testing.T) {
		ok += testParseDir(t, cfg, testPredef, filepath.Join(root, filepath.FromSlash(dir)), false, true)
	})
	t.Run(dir+"/gnu", func(t *testing.T) {
		ok += testParseDir(t, cfg, testPredefGNU, filepath.Join(root, filepath.FromSlash(dir)), false, true)
	})
	t.Logf("ok %v", h(ok))
}

func TestParseGCC(t *testing.T) {
	if testing.Short() {
		t.Skip("-short")
		return
	}

	cfg := &Config{
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
			ok += testParseDir(t, cfg, testPredef, filepath.Join(root, filepath.FromSlash(v)), true, false)
		})
		t.Run(v+"/gnu", func(t *testing.T) {
			ok += testParseDir(t, cfg, testPredefGNU, filepath.Join(root, filepath.FromSlash(v)), true, true)
		})
	}
	t.Logf("ok %v", h(ok))
}

func testParseDir(t *testing.T, cfg *Config, predef, dir string, hfiles, must bool) (ok int) {
	blacklist := map[string]struct{}{ //TODO-
		"90_struct-init.c":  {},
		"94_generic.c":      {},
		"95_bitfields.c":    {},
		"95_bitfields_ms.c": {},
		"99_fastcall.c":     {},
	}
	var re *regexp.Regexp
	if s := *oRE; s != "" {
		re = regexp.MustCompile(s)
	}

	var files, psources int
	var bytes int64
	var m0, m1 runtime.MemStats
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m0)
	t0 := time.Now()
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
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
			ok++
			return nil
		}

		sources := []Source{
			{Name: "<predefined>", Value: predef},
			{Name: "<built-in>", Value: parserTestBuiltin},
			{Name: path},
		}
		ctx := newContext(cfg)

		defer func() {
			psources += ctx.tuSources
			bytes += ctx.tuSize
		}()

		if *oTrace {
			fmt.Fprintln(os.Stderr, files, path)
		}
		if _, err := parse(ctx, testIncludes, testSysIncludes, sources); err != nil {
			if must {
				t.Error(err)
			}
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
	if files != ok && must {
		t.Errorf("files %v, bytes %v, ok %v", files, bytes, ok)
	}
	return ok
}

func BenchmarkParseTCC(b *testing.B) {
	root := filepath.Join(testWD, filepath.FromSlash(tccDir))
	if _, err := os.Stat(root); err != nil {
		b.Skipf("Missing resources in %s. Please run 'go test -download' to fix.", root)
	}

	cfg := &Config{
		ignoreUndefinedIdentifiers: true,
	}
	const dir = "tests/tests2"
	b.Run(dir, func(b *testing.B) {
		benchmarkParseDir(b, cfg, testPredef, filepath.Join(root, filepath.FromSlash(dir)), false)
	})
	b.Run(dir+"/gnu", func(b *testing.B) {
		benchmarkParseDir(b, cfg, testPredefGNU, filepath.Join(root, filepath.FromSlash(dir)), false)
	})
}

func BenchmarkParseGCC(b *testing.B) {
	root := filepath.Join(testWD, filepath.FromSlash(gccDir))
	if _, err := os.Stat(root); err != nil {
		b.Skipf("Missing resources in %s. Please run 'go test -download -dev' to fix.", root)
	}

	cfg := &Config{
		ignoreUndefinedIdentifiers: true,
	}
	for _, v := range []string{
		"gcc/testsuite/gcc.c-torture/compile",
		"gcc/testsuite/gcc.c-torture/execute",
	} {
		b.Run(v, func(b *testing.B) {
			benchmarkParseDir(b, cfg, testPredef, filepath.Join(root, filepath.FromSlash(v)), false)
		})
		b.Run(v+"/gnu", func(b *testing.B) {
			benchmarkParseDir(b, cfg, testPredefGNU, filepath.Join(root, filepath.FromSlash(v)), true)
		})
	}
}

func benchmarkParseDir(b *testing.B, cfg *Config, predef, dir string, must bool) {
	var bytes int64
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes = 0
		if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return skipDir(path)
			}

			if filepath.Ext(path) != ".c" && filepath.Ext(path) != ".h" || info.Mode()&os.ModeType != 0 {
				return nil
			}

			sources := []Source{
				{Name: "<predefined>", Value: predef},
				{Name: "<built-in>", Value: parserTestBuiltin},
				{Name: path},
			}
			ctx := newContext(cfg)
			if _, err := parse(ctx, testIncludes, testSysIncludes, sources); err != nil {
				if must {
					b.Error(err)
				}
			}
			bytes += ctx.tuSize
			return nil
		}); err != nil {
			b.Error(err)
		}
	}
	b.SetBytes(bytes)
}

func ExampleInitDeclaratorList_uCN() {
	fmt.Println(exampleAST(0, `int a·z, a\u00b7z;`))
	// Output:
	// &cc.InitDeclaratorList{
	// · InitDeclarator: &cc.InitDeclarator{
	// · · Case: InitDeclaratorDecl,
	// · · Declarator: &cc.Declarator{
	// · · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · · Case: DirectDeclaratorIdent,
	// · · · · Token: example.c:1:5: IDENTIFIER "a·z",
	// · · · },
	// · · },
	// · },
	// · InitDeclaratorList: &cc.InitDeclaratorList{
	// · · InitDeclarator: &cc.InitDeclarator{
	// · · · Case: InitDeclaratorDecl,
	// · · · Declarator: &cc.Declarator{
	// · · · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · · · Case: DirectDeclaratorIdent,
	// · · · · · Token: example.c:1:11: IDENTIFIER "a·z",
	// · · · · },
	// · · · },
	// · · },
	// · · Token: example.c:1:9: ',' ",",
	// · },
	// }
}

func ExampleDirectDeclarator_line() {
	fmt.Println(exampleAST(0, "#line 1234\nint i;"))
	// Output:
	// &cc.DirectDeclarator{
	// · Case: DirectDeclaratorIdent,
	// · Token: example.c:1234:5: IDENTIFIER "i",
	// }
}

func ExampleDirectDeclarator_line2() {
	fmt.Println(exampleAST(0, "#line 1234 \"foo.c\"\nint i;"))
	// Output:
	// &cc.DirectDeclarator{
	// · Case: DirectDeclaratorIdent,
	// · Token: foo.c:1234:5: IDENTIFIER "i",
	// }
}
