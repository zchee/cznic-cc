// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"testing"

	"github.com/dustin/go-humanize"
	"github.com/pmezard/go-difflib/difflib"
	"modernc.org/ccorpus"
	"modernc.org/httpfs"
	"modernc.org/mathutil"
	//TODO "modernc.org/scannertest"
)

var (
	cfs         = ccorpus.FileSystem()
	cFS         = &corpusFS{cfs}
	corpus      = map[string][]byte{}
	corpusIndex []string
	re          *regexp.Regexp
	testCfg0    = &Config{}
	predefined  string
	builtin     = `
#define __extension__

#ifndef __builtin_va_list
#define __builtin_va_list __builtin_va_list
typedef void *__builtin_va_list;
#endif

#ifndef __builtin_va_arg
#define __builtin_va_arg(va, type) (*(type*)__builtin_va_arg_sink(0, va))
void *__builtin_va_arg_sink(int, ...);
#endif
`

	oTrace = flag.Bool("trc", false, "Print tested paths.")
)

func init() {
	isTesting = true
	var err error
	if predefined, testCfg0.IncludePaths, testCfg0.SysIncludePaths, err = HostConfig(""); err != nil {
		panic(errorf("cannot acquire host configuration: %v", err))
	}

	switch fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH) {
	case "netbsd/amd64":
		testCfg0.SysIncludePaths = append(testCfg0.SysIncludePaths, "/usr/pkg/include")
	case "freebsd/386":
		testCfg0.SysIncludePaths = append(testCfg0.SysIncludePaths, "/usr/local/include")
	}
	testCfg0.SysIncludePaths = testCfg0.SysIncludePaths[:len(testCfg0.SysIncludePaths):len(testCfg0.SysIncludePaths)]
	testCfg0.IncludePaths = append([]string{""}, testCfg0.IncludePaths...)
	testCfg0.IncludePaths = append(testCfg0.IncludePaths, testCfg0.SysIncludePaths...)
	testCfg0.IncludePaths = testCfg0.IncludePaths[:len(testCfg0.IncludePaths):len(testCfg0.IncludePaths)]
	if testCfg0.ABI, err = NewABI(runtime.GOOS, runtime.GOARCH); err != nil {
		panic(errorf("cannot configure ABI: %v", err))
	}

	var chars int
	if err := walk("/", func(pth string, fi os.FileInfo) error {
		if fi.IsDir() {
			return nil
		}

		f, err := cfs.Open(pth)
		if err != nil {
			return errorf("%v: %v", pth, err)
		}

		b, err := io.ReadAll(f)
		if err != nil {
			return errorf("%v: %v", pth, err)
		}

		switch filepath.Ext(pth) {
		case ".c", ".h":
			if len(b) != 0 && b[len(b)-1] != '\n' {
				b = append(b, '\n')
			}
		}
		chars += len(b)
		corpus[pth] = b
		corpusIndex = append(corpusIndex, pth)
		return nil
	}); err != nil {
		panic(err)
	}
}

type corpusFS struct {
	*httpfs.FileSystem
}

func (c *corpusFS) Open(name string) (fs.File, error) {
	name = filepath.ToSlash(name)
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	f, err := c.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}

	return fs.File(f), nil
}

func walk(dir string, f func(pth string, fi os.FileInfo) error) error {
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	root, err := cfs.Open(dir)
	if err != nil {
		return err
	}

	fi, err := root.Stat()
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return fmt.Errorf("%s: not a directory", fi.Name())
	}

	fis, err := root.Readdir(-1)
	if err != nil {
		return err
	}

	for _, v := range fis {
		switch {
		case v.IsDir():
			if err = walk(v.Name(), f); err != nil {
				return err
			}
		default:
			if err = f(v.Name(), v); err != nil {
				return err
			}
		}
	}
	return nil
}

// Produce the AST used in examples documentation.
func exampleAST(rule int, src string) (r interface{}) {
	defer func() {
		if err := recover(); err != nil {
			r = fmt.Sprintf("%v (%v:)", err, origin(5))
		}
	}()

	src = strings.Replace(src, "\\n", "\n", -1)
	cfg := &Config{}
	ast, _ := Parse(cfg, []Source{{Name: "example.c", Value: src}})
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

func TestMain(m *testing.M) {
	oRE := flag.String("re", "", "")
	flag.Parse()
	if *oRE != "" {
		re = regexp.MustCompile(*oRE)
	}
	os.Exit(m.Run())
}

func TestScannerSource(t *testing.T) {
	const fn = "all_test.go"
	exp, err := os.ReadFile(fn)
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(fn)
	if err != nil {
		t.Fatal(err)
	}

	testScannerSource(t, fn, f, exp, false)
	testScannerSource(t, fn, exp, exp, false)
	testScannerSource(t, fn, string(exp), exp, false)
	testScannerSource(t, fn, bytes.NewReader(exp), exp, false)
	testScannerSource(t, fn, nil, exp, false)
	testScannerSource(t, fn, 42, nil, true)
}

func testScannerSource(t *testing.T, name string, value interface{}, exp []byte, mustFail bool) {
	ss, err := newScannerSource(Source{name, value, nil})
	if err != nil != mustFail {
		t.Fatalf("(%q, %T): %v", name, value, err)
	}

	if err != nil {
		return
	}

	if !bytes.Equal(ss.buf, exp) {
		t.Fatal("buf does not match")
	}
}

func TestToken(t *testing.T) {
	s, err := newScannerSource(Source{"test", `abc
def
 ghi
`, nil})
	// abc\ndef\n ghi\n
	//             1
	// 0123 4567 89012
	if err != nil {
		t.Fatal(err)
	}

	s.file.AddLine(4)
	s.file.AddLine(8)
	s.file.AddLine(13)
	for itest, test := range []struct {
		Token
		line int
		col  int
		ch   rune
		sep  string
		src  string
	}{
		{newToken(s, 0, 0, 0, 3), 1, 1, 0, "", "abc"},
		{newToken(s, 1, 3, 4, 3), 2, 1, 1, "\n", "def"},
		{newToken(s, 2, 7, 9, 3), 3, 2, 2, "\n ", "ghi"},
		{newToken(s, eof, 13, 13, 0), 3, 6, eof, "", ""},
	} {
		tok := test.Token
		if g, e := tok.Position().Line, test.line; g != e {
			t.Fatal(itest, g, e)
		}
		if g, e := tok.Position().Column, test.col; g != e {
			t.Fatal(itest, g, e)
		}
		if g, e := tok.Ch, test.ch; g != e {
			t.Fatal(itest, g, e)
		}
		if g, e := string(tok.Sep()), test.sep; g != e {
			t.Fatalf("%v %q %q", itest, g, e)
		}
		if g, e := string(tok.Src()), test.src; g != e {
			t.Fatalf("%v %q %q", itest, g, e)
		}

		tok2 := tok
		tok2.Set([]byte("xyz0123"), []byte("456789"))
		if g, e := string(tok.Sep()), test.sep; g != e {
			t.Fatalf("%v %q %q", itest, g, e)
		}
		if g, e := string(tok.Src()), test.src; g != e {
			t.Fatalf("%v %q %q", itest, g, e)
		}
		if g, e := string(tok2.Sep()), "xyz0123"; g != e {
			t.Fatalf("%v %q %q", itest, g, e)
		}
		if g, e := string(tok2.Src()), "456789"; g != e {
			t.Fatalf("%v %q %q", itest, g, e)
		}
	}
}

type parallel struct {
	limit chan struct{}
	sync.Mutex
	wg sync.WaitGroup
}

func newParallel() *parallel {
	return &parallel{
		limit: make(chan struct{}, runtime.GOMAXPROCS(0)),
	}
}

func (p *parallel) exec(run func()) {
	p.limit <- struct{}{}
	p.wg.Add(1)

	go func() {
		defer func() {
			p.wg.Done()
			<-p.limit
		}()

		run()
	}()
}

var tokSink []Token

func TestScanner(t *testing.T) {
	defer func() { tokSink = nil }()

	var files, tokens, chars int64
	var m0, m runtime.MemStats
	p := newParallel()
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m0)
	for _, path := range corpusIndex {
		path := path
		switch filepath.Ext(path) {
		case ".c", ".h":
			files++
			p.exec(func() {
				var err error
				var chars0, tokens0 int64
				var toks []Token

				defer func() {
					p.Lock()
					chars += chars0
					tokens += tokens0
					tokSink = append(tokSink, toks...)
					if err != nil {
						t.Error(err)
					}
					p.Unlock()
				}()

				buf := corpus[path]
				chars0 += int64(len(buf))
				var s *scanner
				if s, err = newScanner(Source{path, buf, nil}, func(msg string, args ...interface{}) {
					s.close()
					err = fmt.Errorf(msg, args...)
				}); err != nil {
					err = fmt.Errorf("%v: %v", path, err)
					return
				}

				for {
					tok := s.cppScan()
					if tok.Ch == eof {
						return
					}

					toks = append(toks, tok)
					tokens0++
				}
			})
		}
	}
	p.wg.Wait()
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m)
	t.Logf("files %v; tokens %v; bytes %v; heap %v; alloc %v", h(files), h(tokens), h(chars), h(m.HeapAlloc-m0.HeapAlloc), h(m.TotalAlloc-m0.TotalAlloc))
}

func h(v interface{}) string {
	switch x := v.(type) {
	case int64:
		return humanize.Comma(x)
	case uint64:
		if x <= math.MaxInt64 {
			return humanize.Comma(int64(x))
		}
	}
	return fmt.Sprint(v)
}

func BenchmarkScanner(b *testing.B) {
	debug.FreeOSMemory()
	for i := 0; i < b.N; i++ {
		var chars int64
		for _, path := range corpusIndex {
			switch filepath.Ext(path) {
			case ".c", ".h":
				buf := corpus[path]
				chars += int64(len(buf))
				var s *scanner
				var err error
				if s, err = newScanner(Source{path, buf, nil}, func(msg string, args ...interface{}) {
					s.close()
					b.Fatalf(msg, args...)
				}); err != nil {
					b.Fatal(path, err)
				}
				for {
					tok := s.cppScan()
					if tok.Ch == eof {
						break
					}
				}
			}
		}
		b.SetBytes(chars)
	}
}

var (
	cppParseBlacklist = map[string]struct{}{
		"/github.com/vnmakarov/mir/c-tests/new/endif.c": {}, // 1:1: unexpected #endif
	}
	astSink []group
)

func TestCPPParse0(t *testing.T) {
	defer func() { astSink = nil }()

	var files, lines, chars int64
	var m0, m runtime.MemStats
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m0)
	for _, path := range corpusIndex {
		if _, ok := cppParseBlacklist[path]; ok {
			continue
		}

		switch filepath.Ext(path) {
		case ".c", ".h":
			buf := corpus[path]
			chars += int64(len(buf))
			var p *cppParser
			var err error
			if p, err = newCppParser(Source{path, buf, nil}, func(msg string, args ...interface{}) {
				p.close()
				t.Fatalf(msg, args...)
			}); err != nil {
				t.Fatal(path, err)
			}

			files++
			ast := p.preprocessingFile()
			if len(ast) == 0 {
				t.Fatalf("%v: empty AST", path)
			}

			eol := ast[len(ast)-1]
			x, ok := eol.(eofLine)
			if !ok {
				t.Fatalf("%v: AST not terminated: %T", p.pos(), eol)
			}

			eof := Token(x)
			lines += int64(eof.Position().Line)
			astSink = append(astSink, ast)
		}
	}
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m)
	astSink = nil
	t.Logf("files %v; lines %v bytes %v; heap %v; alloc %v", h(files), h(lines), h(chars), h(m.HeapAlloc-m0.HeapAlloc), h(m.TotalAlloc-m0.TotalAlloc))
}

func TestCPPParse(t *testing.T) {
	defer func() { astSink = nil }()

	var files, lines, chars int64
	var m0, m runtime.MemStats
	p := newParallel()
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m0)
	for _, path := range corpusIndex {
		path := path
		if _, ok := cppParseBlacklist[path]; ok {
			continue
		}

		switch filepath.Ext(path) {
		case ".c", ".h":
			files++
			p.exec(func() {
				buf := corpus[path]
				var err error
				var ast group
				var eof Token

				defer func() {
					p.Lock()
					chars += int64(len(buf))
					lines += int64(eof.Position().Line)
					astSink = append(astSink, ast)
					if err != nil {
						t.Error(err)
					}
					p.Unlock()
				}()

				var p *cppParser
				if p, err = newCppParser(Source{path, buf, nil}, func(msg string, args ...interface{}) {
					p.close()
					err = fmt.Errorf(msg, args...)
				}); err != nil {
					t.Fatal(path, err)
				}

				if ast = p.preprocessingFile(); len(ast) == 0 {
					t.Fatalf("%v: empty AST", path)
				}

				eol := ast[len(ast)-1]
				x, ok := eol.(eofLine)
				if !ok {
					err = fmt.Errorf("%v: AST not terminated: %T", p.pos(), eol)
					return
				}

				eof = Token(x)
			})
		}
	}
	p.wg.Wait()
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m)
	astSink = nil
	t.Logf("files %v; lines %v bytes %v; heap %v; alloc %v", h(files), h(lines), h(chars), h(m.HeapAlloc-m0.HeapAlloc), h(m.TotalAlloc-m0.TotalAlloc))
}

func BenchmarkCPPParse(b *testing.B) {
	debug.FreeOSMemory()
	for i := 0; i < b.N; i++ {
		var chars int64
		for _, path := range corpusIndex {
			if _, ok := cppParseBlacklist[path]; ok {
				continue
			}

			switch filepath.Ext(path) {
			case ".c", ".h":
				buf := corpus[path]
				chars += int64(len(buf))
				var p *cppParser
				var err error
				if p, err = newCppParser(Source{path, buf, nil}, func(msg string, args ...interface{}) {
					p.close()
					b.Fatalf(msg, args...)
				}); err != nil {
					b.Fatal(path, err)
				}

				ast := p.preprocessingFile()
				if len(ast) == 0 {
					b.Fatalf("%v: empty AST", path)
				}

				eol := ast[len(ast)-1]
				if _, ok := eol.(eofLine); !ok {
					b.Fatalf("%v: AST not terminated: %T", p.pos(), eol)
				}
			}
		}
		b.SetBytes(chars)
	}
}

func testCfg() *Config {
	c := *testCfg0
	return &c
}

func TestCPPExpand(t *testing.T) {
	testCPPExpand(t, "testdata/cpp-expand/", nil, true)
}

func testCPPExpand(t *testing.T, dir string, blacklist map[string]struct{}, fakeIncludes bool) {
	var fails []string
	var files, ok, skip int
	var c *cpp
	cfg := testCfg()
	cfg.fakeIncludes = fakeIncludes
	cfg.PragmaHandler = func(s []Token) error {
		a := textLine{pragmaTestTok}
		for i, v := range s {
			if i == 0 {
				v.Set(sp, v.Src())
			}
			a = append(a, v)
		}
		c.push(append(a, nlTok))
		return nil
	}
	err := filepath.Walk(filepath.FromSlash(dir), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || (!strings.HasSuffix(path, ".c") && !strings.HasSuffix(path, ".h")) {
			return nil
		}

		files++
		switch {
		case re != nil:
			if !re.MatchString(path) {
				skip++
				return nil
			}
		default:
			if _, ok := blacklist[filepath.Base(path)]; ok {
				skip++
				return nil
			}
		}

		if *oTrace {
			fmt.Fprintln(os.Stderr, path)
		}
		var b strings.Builder
		if c, err = newCPP(cfg, []Source{{path, nil, nil}}, nil); err != nil {
			t.Fatalf("%v: %v", path, err)
		}

		if err := preprocess(c, &b); err != nil {
			fails = append(fails, path)
			t.Fatalf("%v: %v", path, err)
		}

		if strings.Contains(filepath.ToSlash(path), "/mustfail/") {
			if err != nil {
				return nil
			}

			fails = append(fails, path)
			return fmt.Errorf("%v: unexpected success", path)
		}

		if err != nil {
			fails = append(fails, path)
			return err
		}

		expFn := path + ".expect"
		exp, err := os.ReadFile(expFn)
		if err != nil {
			fails = append(fails, path)
			t.Error(err)
		}

		g := strings.ReplaceAll(b.String(), "\r", "")
		g = strings.TrimSpace(g)
		e := strings.ReplaceAll(string(exp), "\r", "")
		e = strings.TrimSpace(e)
		if g != e {
			fails = append(fails, path)
			diff := difflib.UnifiedDiff{
				A:        difflib.SplitLines(e),
				B:        difflib.SplitLines(g),
				FromFile: expFn,
				ToFile:   path,
				Context:  0,
			}
			s, err := difflib.GetUnifiedDiffString(diff)
			if err != nil {
				t.Fatalf("%v: %v", path, err)
			}

			t.Errorf("%v\ngot\n%s\nexp\n%s\ngot\n%s\nexp\n%s", s, g, e, hex.Dump([]byte(g)), hex.Dump([]byte(e)))
			return nil
		}
		ok++
		return nil
	})
	for _, v := range fails {
		t.Log(v)
	}
	t.Logf("files %v, skip %v, ok %v, fails %v", files, skip, ok, len(fails))
	if err != nil {
		t.Fatal(err)
	}
}

func TestPreprocess(t *testing.T) {
	testCPPExpand(t, "testdata/preprocess/", nil, true)
}

func TestTCCExpand(t *testing.T) {
	testCPPExpand(t, "testdata/tcc-0.9.27/tests/pp/", map[string]struct{}{
		"11.c": {}, // https://gcc.gnu.org/onlinedocs/gcc/Variadic-Macros.html#Variadic-Macros
		"16.c": {}, // We don't produce warnings on macro redefinition.
	}, true)
}

func TestInclude(t *testing.T) {
	testCPPExpand(t, "testdata/include/", nil, false)
}

func TestTranslationPhase4(t *testing.T) {
	cfgGame := testCfg()
	cfgGame.FS = cFS
	cfgGame.SysIncludePaths = append(
		cfgGame.SysIncludePaths,
		"/Library/Developer/CommandLineTools/SDKs/MacOSX11.1.sdk/usr/include/libxml",
		"/Library/Developer/CommandLineTools/SDKs/MacOSX11.1.sdk/usr/include/malloc",
		"/Library/Developer/CommandLineTools/SDKs/MacOSX12.1.sdk/usr/include/libxml",
		"/Library/Developer/CommandLineTools/SDKs/MacOSX12.1.sdk/usr/include/malloc",
		"/benchmarksgame-team.pages.debian.net/Include",
		"/opt/homebrew/include",
		"/usr/include/sys",
		"/usr/lib/clang/11.1.0/include",
		"/usr/local/Cellar/gcc/11.2.0_1/lib/gcc/11/gcc/x86_64-apple-darwin19/11.2.0/include",
		"/usr/local/include",
	)
	cfgGame.IncludePaths = append(
		cfgGame.IncludePaths,
		"/opt/homebrew/include",
		"/usr/local/include",
	)
	cfg := testCfg()
	cfg.FS = cFS
	var blacklistCompCert, blacklistCxgo map[string]struct{}
	blacklistGame := map[string]struct{}{
		// Missing <apr_pools.h>
		"binary-trees-2.c": {},
		"binary-trees-3.c": {},

		"binary-trees-5.c":       {}, //TODO
		"fannkuchredux-2.c":      {}, //TODO
		"fannkuchredux-3.c":      {}, //TODO
		"fannkuchredux-4.c":      {}, //TODO
		"fasta-4.c":              {}, //TODO
		"fasta.c":                {}, //TODO
		"k-nucleotide.c":         {}, //TODO
		"mandelbrot-3.c":         {}, //TODO
		"mandelbrot-4.c":         {}, //TODO
		"mandelbrot-6.c":         {}, //TODO
		"mandelbrot-7.c":         {}, //TODO
		"mandelbrot.c":           {}, //TODO
		"nbody-4.c":              {}, //TODO
		"nbody-8.c":              {}, //TODO
		"nbody-9.c":              {}, //TODO
		"pidigits-2.c":           {}, //TODO
		"pidigits-6.c":           {}, //TODO
		"pidigits.c":             {}, //TODO
		"regex-redux-3.c":        {}, //TODO
		"regex-redux-4.c":        {}, //TODO
		"regex-redux-5.c":        {}, //TODO
		"reverse-complement-2.c": {}, //TODO
		"reverse-complement-4.c": {}, //TODO
		"reverse-complement-5.c": {}, //TODO
		"spectral-norm-4.c":      {}, //TODO
		"spectral-norm-5.c":      {}, //TODO
		"spectral-norm-6.c":      {}, //TODO
	}
	blacklistGCC := map[string]struct{}{
		// assertions are deprecated.
		"950919-1.c": {},

		// Need include files not in ccorpus.
		"pr88347.c": {},
		"pr88423.c": {},
	}
	blacklistVNMakarov := map[string]struct{}{
		// #endif without #if
		"endif.c": {},
	}
	blacklictTCC := map[string]struct{}{
		"11.c": {}, // https://gcc.gnu.org/onlinedocs/gcc/Variadic-Macros.html#Variadic-Macros
	}
	switch fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH) {
	case "linux/s390x":
		blacklistCompCert = map[string]struct{}{"aes.c": {}} // Unsupported endianness.
		fallthrough
	case "linux/arm", "linux/arm64":
		// Uses sse2 headers.
		blacklistGame["fannkuchredux-4.c"] = struct{}{}
		blacklistGame["mandelbrot-6.c"] = struct{}{}
		blacklistGame["nbody-4.c"] = struct{}{}
		blacklistGame["nbody-8.c"] = struct{}{}
		blacklistGame["nbody-9.c"] = struct{}{}
		blacklistGame["spectral-norm-5.c"] = struct{}{}
		blacklistGame["spectral-norm-6.c"] = struct{}{}
	case "freebsd/386", "darwin/amd64", "darwin/arm64", "freebsd/amd64":
		blacklistCompCert = map[string]struct{}{"aes.c": {}} // include file not found: "../endian.h"
	case "windows/amd64", "windows/386":
		blacklistCompCert = map[string]struct{}{"aes.c": {}} // include file not found: "../endian.h"
		blacklistCxgo = map[string]struct{}{"inet.c": {}}    // include file not found: <arpa/inet.h>
		blacklistGCC["loop-2f.c"] = struct{}{}               // include file not found: <sys/mman.h>
		blacklistGCC["loop-2g.c"] = struct{}{}               // include file not found: <sys/mman.h>
		blacklistGame["fasta-4.c"] = struct{}{}              // include file not found: <err.h>
		blacklistGame["pidigits-2.c"] = struct{}{}           // include file not found: <gmp.h>
		blacklistGame["pidigits-6.c"] = struct{}{}           // include file not found: <threads.h>
		blacklistGame["pidigits-9.c"] = struct{}{}           // include file not found: <gmp.h>
		blacklistGame["pidigits.c"] = struct{}{}             // include file not found: <gmp.h>
		blacklistGame["regex-redux-2.c"] = struct{}{}        // include file not found: <pcre.h>
		blacklistGame["regex-redux-3.c"] = struct{}{}        // include file not found: <pcre.h>
		blacklistGame["regex-redux-4.c"] = struct{}{}        // include file not found: <pcre.h>
		blacklistGame["regex-redux-5.c"] = struct{}{}        // include file not found: <pcre2.h>
	case "openbsd/amd64":
		blacklistCompCert = map[string]struct{}{"aes.c": {}} // include file not found: "../endian.h"
		blacklistGame["mandelbrot-7.c"] = struct{}{}         // include file not found: <omp.h>
		blacklistGame["pidigits-6.c"] = struct{}{}           // include file not found: <threads.h>
		blacklistGame["regex-redux-3.c"] = struct{}{}        // include file not found: <omp.h>
		blacklistGame["spectral-norm-4.c"] = struct{}{}      // include file not found: <omp.h>
	}
	var files, ok, skip, fails int
	for _, v := range []struct {
		cfg       *Config
		dir       string
		blacklist map[string]struct{}
	}{
		{cfg, "CompCert-3.6/test/c", blacklistCompCert},
		{cfg, "ccgo", nil},
		{cfg, "gcc-9.1.0/gcc/testsuite/gcc.c-torture", blacklistGCC},
		{cfg, "github.com/AbsInt/CompCert/test/c", blacklistCompCert},
		{cfg, "github.com/cxgo", blacklistCxgo},
		{cfg, "github.com/gcc-mirror/gcc/gcc/testsuite", blacklistGCC},
		{cfg, "github.com/vnmakarov", blacklistVNMakarov},
		{cfg, "sqlite-amalgamation-3370200", nil},
		{cfg, "tcc-0.9.27/tests", blacklictTCC},
		{cfgGame, "benchmarksgame-team.pages.debian.net", blacklistGame},
	} {
		t.Run(v.dir, func(t *testing.T) {
			f, o, s, n := testTranslationPhase4(t, v.cfg, "/"+v.dir, v.blacklist)
			files += f
			ok += o
			skip += s
			fails += n
		})
	}
	t.Logf("TOTAL: files %v, skip %v, ok %v, fails %v", files, skip, ok, fails)
}

func testTranslationPhase4(t *testing.T, cfg *Config, dir string, blacklist map[string]struct{}) (files, ok, skip, nfails int) {
	var fails []string
	p := newParallel()
	err := walk(dir, func(pth string, fi os.FileInfo) error {
		if fi.IsDir() {
			return nil
		}

		if filepath.Ext(pth) != ".c" {
			return nil
		}

		switch {
		case re != nil:
			if !re.MatchString(pth) {
				skip++
				return nil
			}
		default:
			if _, ok := blacklist[filepath.Base(pth)]; ok {
				skip++
				return nil
			}
		}

		files++
		p.exec(func() {
			if *oTrace {
				fmt.Fprintln(os.Stderr, pth)
			}
			err := Preprocess(
				cfg,
				[]Source{
					{Name: "<predefined>", Value: predefined},
					{Name: "<builtin>", Value: builtin},
					{Name: pth, FS: cFS},
				},
				io.Discard,
			)
			p.Lock()

			defer p.Unlock()

			if err != nil {
				fails = append(fails, pth)
				t.Errorf("%v: %v", pth, err)
			} else {
				ok++
			}
		})
		return nil
	})
	if err != nil {
		t.Error(err)
	}

	p.wg.Wait()
	for _, v := range fails {
		t.Log(v)
	}
	t.Logf("files %v, skip %v, ok %v, fails %v", files, skip, ok, len(fails))
	return files, ok, skip, len(fails)
}

// https://gitlab.com/cznic/cc/-/issues/127
func TestIssue127(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := os.Chdir(wd); err != nil {
			t.Fatal(err)
		}
	}()

	if err := os.Chdir(filepath.FromSlash("testdata/issue127/")); err != nil {
		t.Error(err)
		return
	}

	cd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("working directory: %s", cd)
	cfg := testCfg()
	cfg.IncludePaths = append(cfg.IncludePaths, "include")
	if err := Preprocess(
		cfg,
		[]Source{{Name: "main.c"}},
		io.Discard,
	); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBOM(t *testing.T) {
	for i, v := range []struct {
		src string
		err string
	}{
		{"int main() {}", ""},
		{"\xEF\xBB\xBFint main() {}", ""},
	} {
		switch _, err := Parse(testCfg(), []Source{{Value: v.src}}); {
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
	for i, v := range []struct {
		src         string
		lit         string
		sep         string
		trailingSep string
	}{
		{`int f() {  "a";}`, `"a"`, "  ", "\n"},
		{`int f() {  L"a";}`, `L"a"`, "  ", "\n"},
		{`int f() { "a" "b";}`, `"ab"`, "  ", "\n"},
		{`int f() { "a""b";}`, `"ab"`, " ", "\n"},
		{`int f() { "a";}`, `"a"`, " ", "\n"},
		{`int f() { "a"` + "\n\t" + `"b"; }`, `"ab"`, " \n\t", "\n"},
		{`int f() { /*x*/ /*y*/ "a";}`, `"a"`, " /*x*/ /*y*/ ", "\n"},
		{`int f() { /*x*/` + "\n" + `/*y*/ "a";}`, `"a"`, " /*x*/\n/*y*/ ", "\n"},
		{`int f() { //x` + "\n" + ` "a";}`, `"a"`, " //x\n ", "\n"},
		{`int f() { //x` + "\n" + `"a";}`, `"a"`, " //x\n", "\n"},
		{`int f() { L"a" L"b";}`, `L"ab"`, "  ", "\n"},
		{`int f() { ` + "\n" + ` "a";}`, `"a"`, " \n ", "\n"},
		{`int f() { ` + "\n" + `"a";}`, `"a"`, " \n", "\n"},
		{`int f() {"a" "b";}`, `"ab"`, " ", "\n"},
		{`int f() {"a"/*y*/"b";}`, `"ab"`, "/*y*/", "\n"},
		{`int f() {"a";} /*x*/ `, `"a"`, "", " /*x*/ \n"},
		{`int f() {"a";} /*x*/`, `"a"`, "", " /*x*/\n"},
		{`int f() {"a";} /*x` + "\n" + `*/ `, `"a"`, "", " /*x\n*/ \n"},
		{`int f() {"a";} `, `"a"`, "", " \n"},
		{`int f() {"a";}/*x*/`, `"a"`, "", "/*x*/\n"},
		{`int f() {"a";}` + "\n", `"a"`, "", "\n"},
		{`int f() {"a";}`, `"a"`, "", "\n"},
		{`int f() {/*x*/ /*y*/ "a";}`, `"a"`, "/*x*/ /*y*/ ", "\n"},
		{`int f() {/*x*/"a""b";}`, `"ab"`, "/*x*/", "\n"},
		{`int f() {/*x*/"a"/*y*/"b";}`, `"ab"`, "/*x*//*y*/", "\n"},
		{`int f() {/*x*/"a";}`, `"a"`, "/*x*/", "\n"},
		{`int f() {/*x*//*y*/ "a";}`, `"a"`, "/*x*//*y*/ ", "\n"},
		{`int f() {/*x*//*y*/"a";}`, `"a"`, "/*x*//*y*/", "\n"},
		{`int f() {//` + "\n" + `"a";}`, `"a"`, "//\n", "\n"},
		{`int f() {//x` + "\n" + `"a";}`, `"a"`, "//x\n", "\n"},
		{`int f() {` + "\n" + ` "a";}`, `"a"`, "\n ", "\n"},
		{`int f() {` + "\n" + `"a";}`, `"a"`, "\n", "\n"},
	} {
		ast, err := Parse(&Config{}, []Source{{Name: "test", Value: v.src}})
		if err != nil {
			t.Errorf("%v: %v", i, err)
			continue
		}

		var n Node
		depth := mathutil.MaxInt
		findNode("PrimaryExpression", ast.TranslationUnit, 0, &n, &depth)
		tok := n.(*PrimaryExpression).Token
		if g, e := string(tok.Src()), v.lit; g != e {
			t.Errorf("%v: %q %q", i, g, e)
		}
		if g, e := string(tok.Sep()), v.sep; g != e {
			t.Errorf("%v: %q %q", i, g, e)
		}
		if g, e := string(ast.EOF.Sep()), v.trailingSep; g != e {
			t.Errorf("%v: %q %q", i, g, e)
		}
	}
}

func TestParserBug(t *testing.T) {
	cfg := testCfg()
	var fails []string
	var files, ok, skip int
	err := filepath.Walk(filepath.FromSlash("testdata/parser/bug"), func(pth string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		if filepath.Ext(pth) != ".c" {
			return nil
		}

		switch {
		case re != nil:
			if !re.MatchString(pth) {
				skip++
				return nil
			}
		}

		files++
		if *oTrace {
			fmt.Fprintln(os.Stderr, pth)
		}
		var ast *AST
		ast, err = Parse(
			cfg,
			[]Source{
				{Name: "<predefined>", Value: predefined},
				{Name: "<builtin>", Value: builtin},
				{Name: pth, FS: cFS},
			},
		)
		_ = ast //TODO-
		if err != nil {
			fails = append(fails, pth)
			t.Errorf("%v: %v", pth, err)
		} else {
			ok++
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}

	for _, v := range fails {
		t.Log(v)
	}
	t.Logf("files %v, skip %v, ok %v, fails %v", files, skip, ok, len(fails))
}

func TestParse(t *testing.T) {
	cfgGame := testCfg()
	cfgGame.FS = cFS
	cfgGame.SysIncludePaths = append(
		cfgGame.SysIncludePaths,
		"/Library/Developer/CommandLineTools/SDKs/MacOSX11.1.sdk/usr/include/libxml",
		"/Library/Developer/CommandLineTools/SDKs/MacOSX11.1.sdk/usr/include/malloc",
		"/Library/Developer/CommandLineTools/SDKs/MacOSX12.1.sdk/usr/include/libxml",
		"/Library/Developer/CommandLineTools/SDKs/MacOSX12.1.sdk/usr/include/malloc",
		"/benchmarksgame-team.pages.debian.net/Include",
		"/opt/homebrew/include",
		"/usr/include/sys",
		"/usr/lib/clang/11.1.0/include",
		"/usr/local/Cellar/gcc/11.2.0_1/lib/gcc/11/gcc/x86_64-apple-darwin19/11.2.0/include",
		"/usr/local/include",
	)
	cfgGame.IncludePaths = append(
		cfgGame.IncludePaths,
		"/opt/homebrew/include",
		"/usr/local/include",
	)
	cfg := testCfg()
	cfg.FS = cFS
	var blacklistCompCert map[string]struct{}
	blacklistGame := map[string]struct{}{
		// Missing <apr_pools.h>
		"binary-trees-2.c": {},
		"binary-trees-3.c": {},

		"binary-trees-5.c":       {}, //TODO
		"fannkuchredux-2.c":      {}, //TODO
		"fannkuchredux-3.c":      {}, //TODO
		"fannkuchredux-4.c":      {}, //TODO
		"fasta-4.c":              {}, //TODO
		"fasta.c":                {}, //TODO
		"k-nucleotide.c":         {}, //TODO
		"mandelbrot-3.c":         {}, //TODO
		"mandelbrot-4.c":         {}, //TODO
		"mandelbrot-6.c":         {}, //TODO
		"mandelbrot-7.c":         {}, //TODO
		"mandelbrot.c":           {}, //TODO
		"nbody-4.c":              {}, //TODO
		"nbody-8.c":              {}, //TODO
		"nbody-9.c":              {}, //TODO
		"pidigits-2.c":           {}, //TODO
		"pidigits-6.c":           {}, //TODO
		"pidigits.c":             {}, //TODO
		"regex-redux-3.c":        {}, //TODO
		"regex-redux-4.c":        {}, //TODO
		"regex-redux-5.c":        {}, //TODO
		"reverse-complement-2.c": {}, //TODO
		"reverse-complement-4.c": {}, //TODO
		"reverse-complement-5.c": {}, //TODO
		"spectral-norm-4.c":      {}, //TODO
		"spectral-norm-5.c":      {}, //TODO
		"spectral-norm-6.c":      {}, //TODO
	}
	blacklistGCC := map[string]struct{}{
		// assertions are deprecated.
		"950919-1.c": {},

		// Need include files not in ccorpus.
		"pr88347.c": {},
		"pr88423.c": {},

		"20000205-1.c":                 {}, //TODO
		"20000403-1.c":                 {}, //TODO
		"20000605-2.c":                 {}, //TODO
		"20000605-3.c":                 {}, //TODO
		"20000914-1.c":                 {}, //TODO
		"20000917-1.c":                 {}, //TODO
		"20001026-1.c":                 {}, //TODO
		"20001203-2.c":                 {}, //TODO
		"20010114-1.c":                 {}, //TODO
		"20010118-1.c":                 {}, //TODO
		"20010122-1.c":                 {}, //TODO
		"20010124-1-lib.c":             {}, //TODO
		"20010124-1.c":                 {}, //TODO
		"20010226-1.c":                 {}, //TODO
		"20010313-1.c":                 {}, //TODO
		"20010325-1.c":                 {}, //TODO
		"20010328-1.c":                 {}, //TODO
		"20010518-1.c":                 {}, //TODO
		"20010605-1.c":                 {}, //TODO
		"20010605-2.c":                 {}, //TODO
		"20010701-1.c":                 {}, //TODO
		"20010714-1.c":                 {}, //TODO
		"20010903-1.c":                 {}, //TODO
		"20010903-2.c":                 {}, //TODO
		"20010911-1.c":                 {}, //TODO
		"20011113-1.c":                 {}, //TODO
		"20020226-1.c":                 {}, //TODO
		"20020303-1.c":                 {}, //TODO
		"20020411-1.c":                 {}, //TODO
		"20020412-1.c":                 {}, //TODO
		"20020508-1.c":                 {}, //TODO
		"20020508-2.c":                 {}, //TODO
		"20020508-3.c":                 {}, //TODO
		"20020706-2.c":                 {}, //TODO
		"20020709-1.c":                 {}, //TODO
		"20020807-1.c":                 {}, //TODO
		"20030105-1.c":                 {}, //TODO
		"20030125-1.c":                 {}, //TODO
		"20030219-1.c":                 {}, //TODO
		"20030314-1.c":                 {}, //TODO
		"20030405-1.c":                 {}, //TODO
		"20030903-1.c":                 {}, //TODO
		"20030910-1.c":                 {}, //TODO
		"20031010-1.c":                 {}, //TODO
		"20031020-1.c":                 {}, //TODO
		"20031023-1.c":                 {}, //TODO
		"20031023-2.c":                 {}, //TODO
		"20031023-3.c":                 {}, //TODO
		"20031023-4.c":                 {}, //TODO
		"20031112-1.c":                 {}, //TODO
		"20031113-1.c":                 {}, //TODO
		"20040101-1.c":                 {}, //TODO
		"20040130-1.c":                 {}, //TODO
		"20040220-1.c":                 {}, //TODO
		"20040302-1.c":                 {}, //TODO
		"20040308-1.c":                 {}, //TODO
		"20040311-1.c":                 {}, //TODO
		"20040614-1.c":                 {}, //TODO
		"20040625-1.c":                 {}, //TODO
		"20040703-1.c":                 {}, //TODO
		"20040709-2.c":                 {}, //TODO
		"20040709-3.c":                 {}, //TODO
		"20040805-1.c":                 {}, //TODO
		"20040901-1.c":                 {}, //TODO
		"20041011-1.c":                 {}, //TODO
		"20041119-1.c":                 {}, //TODO
		"20041124-1.c":                 {}, //TODO
		"20041213-2.c":                 {}, //TODO
		"20041214-1.c":                 {}, //TODO
		"20041218-2.c":                 {}, //TODO
		"20050107-1.c":                 {}, //TODO
		"20050119-1.c":                 {}, //TODO
		"20050119-2.c":                 {}, //TODO
		"20050121-1.c":                 {}, //TODO
		"20050125-1.c":                 {}, //TODO
		"20050203-1.c":                 {}, //TODO
		"20050218-1.c":                 {}, //TODO
		"20050410-1.c":                 {}, //TODO
		"20050502-1.c":                 {}, //TODO
		"20050510-1.c":                 {}, //TODO
		"20050516-1.c":                 {}, //TODO
		"20050801-1.c":                 {}, //TODO
		"20050826-2.c":                 {}, //TODO
		"20051215-1.c":                 {}, //TODO
		"20051216-1.c":                 {}, //TODO
		"20060102-1.c":                 {}, //TODO
		"20060109-1.c":                 {}, //TODO
		"20060202-1.c":                 {}, //TODO
		"20060420-1.c":                 {}, //TODO
		"20060609-1.c":                 {}, //TODO
		"20060930-2.c":                 {}, //TODO
		"20061031-1.c":                 {}, //TODO
		"20070603-1.c":                 {}, //TODO
		"20070623-1.c":                 {}, //TODO
		"20070905-1.c":                 {}, //TODO
		"20070919-1.c":                 {}, //TODO
		"20071018-1.c":                 {}, //TODO
		"20071029-1.c":                 {}, //TODO
		"20071108-1.c":                 {}, //TODO
		"20071120-1.c":                 {}, //TODO
		"20071202-1.c":                 {}, //TODO
		"20071210-1.c":                 {}, //TODO
		"20071216-1.c":                 {}, //TODO
		"20071219-1.c":                 {}, //TODO
		"20071220-1.c":                 {}, //TODO
		"20071220-2.c":                 {}, //TODO
		"20080424-1.c":                 {}, //TODO
		"20080502-1.c":                 {}, //TODO
		"20080506-2.c":                 {}, //TODO
		"20080519-1.c":                 {}, //TODO
		"20080522-1.c":                 {}, //TODO
		"20080604-1.c":                 {}, //TODO
		"20081108-1.c":                 {}, //TODO
		"20081112-1.c":                 {}, //TODO
		"20081117-1.c":                 {}, //TODO
		"20081218-1.c":                 {}, //TODO
		"20090113-2.c":                 {}, //TODO
		"20090113-3.c":                 {}, //TODO
		"20090711-1.c":                 {}, //TODO
		"20090814-1.c":                 {}, //TODO
		"20090907-1.c":                 {}, //TODO
		"20100316-1.c":                 {}, //TODO
		"20100430-1.c":                 {}, //TODO
		"20100609-1.c":                 {}, //TODO
		"20100805-1.c":                 {}, //TODO
		"20100827-1.c":                 {}, //TODO
		"20100915-1.c":                 {}, //TODO
		"20101011-1.c":                 {}, //TODO
		"20101013-1.c":                 {}, //TODO
		"20101216-1.c":                 {}, //TODO
		"20110126-1.c":                 {}, //TODO
		"20110131-1.c":                 {}, //TODO
		"20110902.c":                   {}, //TODO
		"20110906-1.c":                 {}, //TODO
		"20111208-1.c":                 {}, //TODO
		"20111212-1.c":                 {}, //TODO
		"20111227-1.c":                 {}, //TODO
		"20120105-1.c":                 {}, //TODO
		"20120207-1.c":                 {}, //TODO
		"20120615-1.c":                 {}, //TODO
		"20121107-1.c":                 {}, //TODO
		"20121220-1.c":                 {}, //TODO
		"20140425-1.c":                 {}, //TODO
		"20140622-1.c":                 {}, //TODO
		"20170401-1.c":                 {}, //TODO
		"20180921-1.c":                 {}, //TODO
		"20190820-1.c":                 {}, //TODO
		"20190901-1.c":                 {}, //TODO
		"920301-1.c":                   {}, //TODO
		"920302-1.c":                   {}, //TODO
		"920415-1.c":                   {}, //TODO
		"920428-2.c":                   {}, //TODO
		"920428-3.c":                   {}, //TODO
		"920501-1.c":                   {}, //TODO
		"920501-13.c":                  {}, //TODO
		"920501-16.c":                  {}, //TODO
		"920501-3.c":                   {}, //TODO
		"920501-4.c":                   {}, //TODO
		"920501-5.c":                   {}, //TODO
		"920501-7.c":                   {}, //TODO
		"920502-1.c":                   {}, //TODO
		"920502-2.c":                   {}, //TODO
		"920625-1.c":                   {}, //TODO
		"920711-1.c":                   {}, //TODO
		"920721-3.c":                   {}, //TODO
		"920721-4.c":                   {}, //TODO
		"920817-1.c":                   {}, //TODO
		"920826-1.c":                   {}, //TODO
		"920831-1.c":                   {}, //TODO
		"920917-1.c":                   {}, //TODO
		"920928-4.c":                   {}, //TODO
		"920928-5.c":                   {}, //TODO
		"921012-1.c":                   {}, //TODO
		"921118-1.c":                   {}, //TODO
		"921123-1.c":                   {}, //TODO
		"921215-1.c":                   {}, //TODO
		"930126-1.c":                   {}, //TODO
		"930325-1.c":                   {}, //TODO
		"930510-1.c":                   {}, //TODO
		"930513-1.c":                   {}, //TODO
		"930926-1.c":                   {}, //TODO
		"931009-1.c":                   {}, //TODO
		"940712-1.c":                   {}, //TODO
		"940718-1.c":                   {}, //TODO
		"941014-1.c":                   {}, //TODO
		"941014-2.c":                   {}, //TODO
		"941202-1.c":                   {}, //TODO
		"950613-1.c":                   {}, //TODO
		"951204-1.c":                   {}, //TODO
		"960116-1.c":                   {}, //TODO
		"960301-1.c":                   {}, //TODO
		"961031-1.c":                   {}, //TODO
		"980505-1.c":                   {}, //TODO
		"980526-1.c":                   {}, //TODO
		"980618-1.c":                   {}, //TODO
		"980929-1.c":                   {}, //TODO
		"981001-2.c":                   {}, //TODO
		"981001-4.c":                   {}, //TODO
		"981223-1.c":                   {}, //TODO
		"990106-2.c":                   {}, //TODO
		"990208-1.c":                   {}, //TODO
		"990326-1.c":                   {}, //TODO
		"990413-2.c":                   {}, //TODO
		"990523-1.c":                   {}, //TODO
		"990628-1.c":                   {}, //TODO
		"990928-1.c":                   {}, //TODO
		"991026-2.c":                   {}, //TODO
		"991213-1.c":                   {}, //TODO
		"991213-2.c":                   {}, //TODO
		"991213-3.c":                   {}, //TODO
		"DFcmp.c":                      {}, //TODO
		"HIcmp.c":                      {}, //TODO
		"HIset.c":                      {}, //TODO
		"SFset.c":                      {}, //TODO
		"SIcmp.c":                      {}, //TODO
		"SIset.c":                      {}, //TODO
		"UHIcmp.c":                     {}, //TODO
		"USIcmp.c":                     {}, //TODO
		"abs-2-lib.c":                  {}, //TODO
		"abs-3-lib.c":                  {}, //TODO
		"abs.c":                        {}, //TODO
		"alias-1.c":                    {}, //TODO
		"alias-3.c":                    {}, //TODO
		"align-3.c":                    {}, //TODO
		"align-nest.c":                 {}, //TODO
		"alloca-1.c":                   {}, //TODO
		"arith-1.c":                    {}, //TODO
		"band.c":                       {}, //TODO
		"bcp-1.c":                      {}, //TODO
		"bf-pack-1.c":                  {}, //TODO
		"bfill.c":                      {}, //TODO
		"bitfld-5.c":                   {}, //TODO
		"bswap-2.c":                    {}, //TODO
		"built-in-setjmp.c":            {}, //TODO
		"builtin-bitops-1.c":           {}, //TODO
		"builtin-prefetch-2.c":         {}, //TODO
		"builtin-prefetch-3.c":         {}, //TODO
		"builtin-prefetch-4.c":         {}, //TODO
		"builtin-types-compatible-p.c": {}, //TODO
		"builtin_constant_p.c":         {}, //TODO
		"bzero.c":                      {}, //TODO
		"callind.c":                    {}, //TODO
		"chk.c":                        {}, //TODO
		"comp-goto-1.c":                {}, //TODO
		"comp-goto-2.c":                {}, //TODO
		"complex-1-lib.c":              {}, //TODO
		"complex-1.c":                  {}, //TODO
		"complex-2.c":                  {}, //TODO
		"complex-3.c":                  {}, //TODO
		"complex-4.c":                  {}, //TODO
		"complex-7.c":                  {}, //TODO
		"compound-literal-3.c":         {}, //TODO
		"conv.c":                       {}, //TODO
		"conversion.c":                 {}, //TODO
		"copysign1.c":                  {}, //TODO
		"cp.c":                         {}, //TODO
		"dll.c":                        {}, //TODO
		"eeprof-1.c":                   {}, //TODO
		"enum-3.c":                     {}, //TODO
		"ffs-1.c":                      {}, //TODO
		"ffs-2.c":                      {}, //TODO
		"fp-cmp-4.c":                   {}, //TODO
		"fp-cmp-4f.c":                  {}, //TODO
		"fp-cmp-4l.c":                  {}, //TODO
		"fp-cmp-5.c":                   {}, //TODO
		"fp-cmp-8.c":                   {}, //TODO
		"fp-cmp-8f.c":                  {}, //TODO
		"fp-cmp-8l.c":                  {}, //TODO
		"fprintf-chk-1.c":              {}, //TODO
		"fprintf-lib.c":                {}, //TODO
		"fprintf.c":                    {}, //TODO
		"fputs.c":                      {}, //TODO
		"icfmatch.c":                   {}, //TODO
		"ifcvt-onecmpl-abs-1.c":        {}, //TODO
		"ipa-sra-1.c":                  {}, //TODO
		"ipa-sra-2.c":                  {}, //TODO
		"labels-2.c":                   {}, //TODO
		"labels-3.c":                   {}, //TODO
		"libcall-1.c":                  {}, //TODO
		"loop-12.c":                    {}, //TODO
		"lto-tbaa-1.c":                 {}, //TODO
		"mangle-1.c":                   {}, //TODO
		"mayalias-1.c":                 {}, //TODO
		"mayalias-2.c":                 {}, //TODO
		"mayalias-3.c":                 {}, //TODO
		"memchr-lib.c":                 {}, //TODO
		"memchr.c":                     {}, //TODO
		"memcmp-lib.c":                 {}, //TODO
		"memcmp.c":                     {}, //TODO
		"memcpy-2.c":                   {}, //TODO
		"memcpy-chk-lib.c":             {}, //TODO
		"memcpy-chk.c":                 {}, //TODO
		"memmove-2-lib.c":              {}, //TODO
		"memmove-chk-lib.c":            {}, //TODO
		"memmove-chk.c":                {}, //TODO
		"memmove-lib.c":                {}, //TODO
		"memmove.c":                    {}, //TODO
		"memops-asm-lib.c":             {}, //TODO
		"memops-asm.c":                 {}, //TODO
		"mempcpy-2-lib.c":              {}, //TODO
		"mempcpy-2.c":                  {}, //TODO
		"mempcpy-chk-lib.c":            {}, //TODO
		"mempcpy-chk.c":                {}, //TODO
		"mempcpy-lib.c":                {}, //TODO
		"mempcpy.c":                    {}, //TODO
		"memset-1.c":                   {}, //TODO
		"memset-2.c":                   {}, //TODO
		"memset-3.c":                   {}, //TODO
		"memset-chk-lib.c":             {}, //TODO
		"memset-chk.c":                 {}, //TODO
		"memset-lib.c":                 {}, //TODO
		"memset.c":                     {}, //TODO
		"mipscop-1.c":                  {}, //TODO
		"mipscop-2.c":                  {}, //TODO
		"mipscop-3.c":                  {}, //TODO
		"mipscop-4.c":                  {}, //TODO
		"mzero2.c":                     {}, //TODO
		"nest-align-1.c":               {}, //TODO
		"nested-1.c":                   {}, //TODO
		"nestfunc-2.c":                 {}, //TODO
		"nestfunc-3.c":                 {}, //TODO
		"nestfunc-6.c":                 {}, //TODO
		"packed-aligned.c":             {}, //TODO
		"pc44485.c":                    {}, //TODO
		"postmod-1.c":                  {}, //TODO
		"pr17397.c":                    {}, //TODO
		"pr17529.c":                    {}, //TODO
		"pr17913.c":                    {}, //TODO
		"pr18903.c":                    {}, //TODO
		"pr19449.c":                    {}, //TODO
		"pr20601-1.c":                  {}, //TODO
		"pr21021.c":                    {}, //TODO
		"pr21293.c":                    {}, //TODO
		"pr21356.c":                    {}, //TODO
		"pr21728.c":                    {}, //TODO
		"pr22013-1.c":                  {}, //TODO
		"pr22061-3.c":                  {}, //TODO
		"pr22098-1.c":                  {}, //TODO
		"pr22098-2.c":                  {}, //TODO
		"pr22098-3.c":                  {}, //TODO
		"pr22141-1.c":                  {}, //TODO
		"pr22141-2.c":                  {}, //TODO
		"pr22379.c":                    {}, //TODO
		"pr22422.c":                    {}, //TODO
		"pr22429.c":                    {}, //TODO
		"pr23135.c":                    {}, //TODO
		"pr23233-1.c":                  {}, //TODO
		"pr23237.c":                    {}, //TODO
		"pr23467.c":                    {}, //TODO
		"pr23484-chk-lib.c":            {}, //TODO
		"pr23484-chk.c":                {}, //TODO
		"pr24135.c":                    {}, //TODO
		"pr25224.c":                    {}, //TODO
		"pr25860.c":                    {}, //TODO
		"pr26781-1.c":                  {}, //TODO
		"pr27073.c":                    {}, //TODO
		"pr27285.c":                    {}, //TODO
		"pr27341-1.c":                  {}, //TODO
		"pr27671-1.c":                  {}, //TODO
		"pr27863.c":                    {}, //TODO
		"pr27889.c":                    {}, //TODO
		"pr28289.c":                    {}, //TODO
		"pr28403.c":                    {}, //TODO
		"pr28489.c":                    {}, //TODO
		"pr28651.c":                    {}, //TODO
		"pr28865.c":                    {}, //TODO
		"pr28982a.c":                   {}, //TODO
		"pr28982b.c":                   {}, //TODO
		"pr29006.c":                    {}, //TODO
		"pr29128.c":                    {}, //TODO
		"pr29241.c":                    {}, //TODO
		"pr30778.c":                    {}, //TODO
		"pr30984.c":                    {}, //TODO
		"pr31169.c":                    {}, //TODO
		"pr32139.c":                    {}, //TODO
		"pr32482.c":                    {}, //TODO
		"pr32584.c":                    {}, //TODO
		"pr32919.c":                    {}, //TODO
		"pr33142.c":                    {}, //TODO
		"pr33173.c":                    {}, //TODO
		"pr33382.c":                    {}, //TODO
		"pr33870-1.c":                  {}, //TODO
		"pr33870.c":                    {}, //TODO
		"pr33992.c":                    {}, //TODO
		"pr34154.c":                    {}, //TODO
		"pr34176.c":                    {}, //TODO
		"pr34334.c":                    {}, //TODO
		"pr34415.c":                    {}, //TODO
		"pr34448.c":                    {}, //TODO
		"pr34456.c":                    {}, //TODO
		"pr34648.c":                    {}, //TODO
		"pr34688.c":                    {}, //TODO
		"pr34768-1.c":                  {}, //TODO
		"pr34768-2.c":                  {}, //TODO
		"pr34856.c":                    {}, //TODO
		"pr34885.c":                    {}, //TODO
		"pr34966.c":                    {}, //TODO
		"pr34993.c":                    {}, //TODO
		"pr35231.c":                    {}, //TODO
		"pr35456.c":                    {}, //TODO
		"pr35472.c":                    {}, //TODO
		"pr35869.c":                    {}, //TODO
		"pr36034-1.c":                  {}, //TODO
		"pr36034-2.c":                  {}, //TODO
		"pr36141.c":                    {}, //TODO
		"pr36339.c":                    {}, //TODO
		"pr36343.c":                    {}, //TODO
		"pr36666.c":                    {}, //TODO
		"pr36765.c":                    {}, //TODO
		"pr37026.c":                    {}, //TODO
		"pr37056.c":                    {}, //TODO
		"pr37102.c":                    {}, //TODO
		"pr37207.c":                    {}, //TODO
		"pr37381.c":                    {}, //TODO
		"pr37418-1.c":                  {}, //TODO
		"pr37418-2.c":                  {}, //TODO
		"pr37418-3.c":                  {}, //TODO
		"pr37418-4.c":                  {}, //TODO
		"pr37573.c":                    {}, //TODO
		"pr37669.c":                    {}, //TODO
		"pr37780.c":                    {}, //TODO
		"pr37913.c":                    {}, //TODO
		"pr38016.c":                    {}, //TODO
		"pr38051.c":                    {}, //TODO
		"pr38151.c":                    {}, //TODO
		"pr38212.c":                    {}, //TODO
		"pr38236.c":                    {}, //TODO
		"pr38422.c":                    {}, //TODO
		"pr38554.c":                    {}, //TODO
		"pr38771.c":                    {}, //TODO
		"pr38819.c":                    {}, //TODO
		"pr38969.c":                    {}, //TODO
		"pr39100.c":                    {}, //TODO
		"pr39120.c":                    {}, //TODO
		"pr39228.c":                    {}, //TODO
		"pr39233.c":                    {}, //TODO
		"pr39240.c":                    {}, //TODO
		"pr39339.c":                    {}, //TODO
		"pr39501.c":                    {}, //TODO
		"pr39845.c":                    {}, //TODO
		"pr39937.c":                    {}, //TODO
		"pr40022.c":                    {}, //TODO
		"pr40057.c":                    {}, //TODO
		"pr40233.c":                    {}, //TODO
		"pr40386.c":                    {}, //TODO
		"pr40579.c":                    {}, //TODO
		"pr40657.c":                    {}, //TODO
		"pr41239.c":                    {}, //TODO
		"pr41395-1.c":                  {}, //TODO
		"pr41395-2.c":                  {}, //TODO
		"pr41463.c":                    {}, //TODO
		"pr41750.c":                    {}, //TODO
		"pr41935.c":                    {}, //TODO
		"pr42142.c":                    {}, //TODO
		"pr42154.c":                    {}, //TODO
		"pr42164.c":                    {}, //TODO
		"pr42196-1.c":                  {}, //TODO
		"pr42196-2.c":                  {}, //TODO
		"pr42196-3.c":                  {}, //TODO
		"pr42231.c":                    {}, //TODO
		"pr42237.c":                    {}, //TODO
		"pr42269-2.c":                  {}, //TODO
		"pr42398.c":                    {}, //TODO
		"pr42632.c":                    {}, //TODO
		"pr42691.c":                    {}, //TODO
		"pr42716.c":                    {}, //TODO
		"pr42717.c":                    {}, //TODO
		"pr42833.c":                    {}, //TODO
		"pr43008.c":                    {}, //TODO
		"pr43164.c":                    {}, //TODO
		"pr43188.c":                    {}, //TODO
		"pr43191.c":                    {}, //TODO
		"pr43255.c":                    {}, //TODO
		"pr43269.c":                    {}, //TODO
		"pr43385.c":                    {}, //TODO
		"pr43560.c":                    {}, //TODO
		"pr43614.c":                    {}, //TODO
		"pr43661.c":                    {}, //TODO
		"pr43679.c":                    {}, //TODO
		"pr43783.c":                    {}, //TODO
		"pr43784.c":                    {}, //TODO
		"pr43791.c":                    {}, //TODO
		"pr43845.c":                    {}, //TODO
		"pr44043.c":                    {}, //TODO
		"pr44063.c":                    {}, //TODO
		"pr44119.c":                    {}, //TODO
		"pr44164.c":                    {}, //TODO
		"pr44197.c":                    {}, //TODO
		"pr44202-1.c":                  {}, //TODO
		"pr44468.c":                    {}, //TODO
		"pr44683.c":                    {}, //TODO
		"pr44784.c":                    {}, //TODO
		"pr44852.c":                    {}, //TODO
		"pr44858.c":                    {}, //TODO
		"pr45034.c":                    {}, //TODO
		"pr45070.c":                    {}, //TODO
		"pr45109.c":                    {}, //TODO
		"pr45695.c":                    {}, //TODO
		"pr46107.c":                    {}, //TODO
		"pr46309.c":                    {}, //TODO
		"pr46316.c":                    {}, //TODO
		"pr46360.c":                    {}, //TODO
		"pr46866.c":                    {}, //TODO
		"pr46909-1.c":                  {}, //TODO
		"pr46909-2.c":                  {}, //TODO
		"pr47299.c":                    {}, //TODO
		"pr47428.c":                    {}, //TODO
		"pr47538.c":                    {}, //TODO
		"pr47925.c":                    {}, //TODO
		"pr48517.c":                    {}, //TODO
		"pr48571-1.c":                  {}, //TODO
		"pr48717.c":                    {}, //TODO
		"pr48814-1.c":                  {}, //TODO
		"pr49039.c":                    {}, //TODO
		"pr49049.c":                    {}, //TODO
		"pr49161.c":                    {}, //TODO
		"pr49218.c":                    {}, //TODO
		"pr49279.c":                    {}, //TODO
		"pr49281.c":                    {}, //TODO
		"pr49390.c":                    {}, //TODO
		"pr49474.c":                    {}, //TODO
		"pr49768.c":                    {}, //TODO
		"pr49886.c":                    {}, //TODO
		"pr50380.c":                    {}, //TODO
		"pr51069.c":                    {}, //TODO
		"pr51323.c":                    {}, //TODO
		"pr51354.c":                    {}, //TODO
		"pr51447.c":                    {}, //TODO
		"pr51466.c":                    {}, //TODO
		"pr51495.c":                    {}, //TODO
		"pr51581-1.c":                  {}, //TODO
		"pr51581-2.c":                  {}, //TODO
		"pr51694.c":                    {}, //TODO
		"pr51767.c":                    {}, //TODO
		"pr51877.c":                    {}, //TODO
		"pr51933.c":                    {}, //TODO
		"pr52129.c":                    {}, //TODO
		"pr52555.c":                    {}, //TODO
		"pr52760.c":                    {}, //TODO
		"pr52891-2.c":                  {}, //TODO
		"pr52979-1.c":                  {}, //TODO
		"pr52979-2.c":                  {}, //TODO
		"pr53084.c":                    {}, //TODO
		"pr53410-2.c":                  {}, //TODO
		"pr53519.c":                    {}, //TODO
		"pr53645-2.c":                  {}, //TODO
		"pr53645.c":                    {}, //TODO
		"pr53688.c":                    {}, //TODO
		"pr53748.c":                    {}, //TODO
		"pr54103-1.c":                  {}, //TODO
		"pr54103-2.c":                  {}, //TODO
		"pr54103-3.c":                  {}, //TODO
		"pr54103-4.c":                  {}, //TODO
		"pr54103-5.c":                  {}, //TODO
		"pr54103-6.c":                  {}, //TODO
		"pr54471.c":                    {}, //TODO
		"pr54552-1.c":                  {}, //TODO
		"pr54713-1.c":                  {}, //TODO
		"pr54713-2.c":                  {}, //TODO
		"pr54713-3.c":                  {}, //TODO
		"pr54937.c":                    {}, //TODO
		"pr54985.c":                    {}, //TODO
		"pr55750.c":                    {}, //TODO
		"pr55875.c":                    {}, //TODO
		"pr55921.c":                    {}, //TODO
		"pr56051.c":                    {}, //TODO
		"pr56205.c":                    {}, //TODO
		"pr56405.c":                    {}, //TODO
		"pr56571.c":                    {}, //TODO
		"pr56799.c":                    {}, //TODO
		"pr56837.c":                    {}, //TODO
		"pr56866.c":                    {}, //TODO
		"pr56899.c":                    {}, //TODO
		"pr56962.c":                    {}, //TODO
		"pr57124.c":                    {}, //TODO
		"pr57130.c":                    {}, //TODO
		"pr57144.c":                    {}, //TODO
		"pr57344-1.c":                  {}, //TODO
		"pr57344-2.c":                  {}, //TODO
		"pr57344-3.c":                  {}, //TODO
		"pr57344-4.c":                  {}, //TODO
		"pr57698.c":                    {}, //TODO
		"pr57829.c":                    {}, //TODO
		"pr58164.c":                    {}, //TODO
		"pr58277-1.c":                  {}, //TODO
		"pr58332.c":                    {}, //TODO
		"pr58419.c":                    {}, //TODO
		"pr58574.c":                    {}, //TODO
		"pr58831.c":                    {}, //TODO
		"pr58984.c":                    {}, //TODO
		"pr59014-2.c":                  {}, //TODO
		"pr59101.c":                    {}, //TODO
		"pr59229.c":                    {}, //TODO
		"pr59358.c":                    {}, //TODO
		"pr59417.c":                    {}, //TODO
		"pr59643.c":                    {}, //TODO
		"pr59747.c":                    {}, //TODO
		"pr59919.c":                    {}, //TODO
		"pr60003.c":                    {}, //TODO
		"pr60062.c":                    {}, //TODO
		"pr60072.c":                    {}, //TODO
		"pr60454.c":                    {}, //TODO
		"pr60655-1.c":                  {}, //TODO
		"pr60655-2.c":                  {}, //TODO
		"pr60822.c":                    {}, //TODO
		"pr60960.c":                    {}, //TODO
		"pr61306-1.c":                  {}, //TODO
		"pr61306-2.c":                  {}, //TODO
		"pr61375.c":                    {}, //TODO
		"pr61673.c":                    {}, //TODO
		"pr63209.c":                    {}, //TODO
		"pr63302.c":                    {}, //TODO
		"pr63641.c":                    {}, //TODO
		"pr63843.c":                    {}, //TODO
		"pr64006.c":                    {}, //TODO
		"pr64242.c":                    {}, //TODO
		"pr64255.c":                    {}, //TODO
		"pr64682.c":                    {}, //TODO
		"pr64718.c":                    {}, //TODO
		"pr64957.c":                    {}, //TODO
		"pr64979.c":                    {}, //TODO
		"pr65053-1.c":                  {}, //TODO
		"pr65170.c":                    {}, //TODO
		"pr65215-1.c":                  {}, //TODO
		"pr65215-2.c":                  {}, //TODO
		"pr65215-3.c":                  {}, //TODO
		"pr65215-4.c":                  {}, //TODO
		"pr65215-5.c":                  {}, //TODO
		"pr65369.c":                    {}, //TODO
		"pr65401.c":                    {}, //TODO
		"pr65418-1.c":                  {}, //TODO
		"pr65418-2.c":                  {}, //TODO
		"pr65427.c":                    {}, //TODO
		"pr65595.c":                    {}, //TODO
		"pr65648.c":                    {}, //TODO
		"pr65873.c":                    {}, //TODO
		"pr65956.c":                    {}, //TODO
		"pr66233.c":                    {}, //TODO
		"pr66940.c":                    {}, //TODO
		"pr67037.c":                    {}, //TODO
		"pr67218.c":                    {}, //TODO
		"pr67226.c":                    {}, //TODO
		"pr67929_1.c":                  {}, //TODO
		"pr68143_1.c":                  {}, //TODO
		"pr68249.c":                    {}, //TODO
		"pr68250.c":                    {}, //TODO
		"pr68328.c":                    {}, //TODO
		"pr68376-2.c":                  {}, //TODO
		"pr68381.c":                    {}, //TODO
		"pr68390.c":                    {}, //TODO
		"pr68532.c":                    {}, //TODO
		"pr68648.c":                    {}, //TODO
		"pr69097-2.c":                  {}, //TODO
		"pr69403.c":                    {}, //TODO
		"pr69447.c":                    {}, //TODO
		"pr69691.c":                    {}, //TODO
		"pr70127.c":                    {}, //TODO
		"pr70190.c":                    {}, //TODO
		"pr70199.c":                    {}, //TODO
		"pr70222-1.c":                  {}, //TODO
		"pr70222-2.c":                  {}, //TODO
		"pr70240.c":                    {}, //TODO
		"pr70355.c":                    {}, //TODO
		"pr70429.c":                    {}, //TODO
		"pr70460.c":                    {}, //TODO
		"pr70566.c":                    {}, //TODO
		"pr70602.c":                    {}, //TODO
		"pr70903.c":                    {}, //TODO
		"pr70916.c":                    {}, //TODO
		"pr71083.c":                    {}, //TODO
		"pr71494.c":                    {}, //TODO
		"pr71554.c":                    {}, //TODO
		"pr71626-1.c":                  {}, //TODO
		"pr71626-2.c":                  {}, //TODO
		"pr71872.c":                    {}, //TODO
		"pr72749.c":                    {}, //TODO
		"pr72802.c":                    {}, //TODO
		"pr77718.c":                    {}, //TODO
		"pr77754-1.c":                  {}, //TODO
		"pr77754-2.c":                  {}, //TODO
		"pr77754-3.c":                  {}, //TODO
		"pr77754-4.c":                  {}, //TODO
		"pr77754-5.c":                  {}, //TODO
		"pr77754-6.c":                  {}, //TODO
		"pr78378.c":                    {}, //TODO
		"pr78436.c":                    {}, //TODO
		"pr78438.c":                    {}, //TODO
		"pr78617.c":                    {}, //TODO
		"pr78622.c":                    {}, //TODO
		"pr78675.c":                    {}, //TODO
		"pr78694.c":                    {}, //TODO
		"pr78720.c":                    {}, //TODO
		"pr78726.c":                    {}, //TODO
		"pr78791.c":                    {}, //TODO
		"pr79043.c":                    {}, //TODO
		"pr79121.c":                    {}, //TODO
		"pr79354.c":                    {}, //TODO
		"pr79388.c":                    {}, //TODO
		"pr80357.c":                    {}, //TODO
		"pr80421.c":                    {}, //TODO
		"pr80501.c":                    {}, //TODO
		"pr80692.c":                    {}, //TODO
		"pr81360.c":                    {}, //TODO
		"pr81423.c":                    {}, //TODO
		"pr81556.c":                    {}, //TODO
		"pr81588.c":                    {}, //TODO
		"pr82052.c":                    {}, //TODO
		"pr82192.c":                    {}, //TODO
		"pr82210.c":                    {}, //TODO
		"pr82337.c":                    {}, //TODO
		"pr82524.c":                    {}, //TODO
		"pr82564.c":                    {}, //TODO
		"pr82879.c":                    {}, //TODO
		"pr82954.c":                    {}, //TODO
		"pr83051-2.c":                  {}, //TODO
		"pr83362.c":                    {}, //TODO
		"pr83487.c":                    {}, //TODO
		"pr83547.c":                    {}, //TODO
		"pr84136.c":                    {}, //TODO
		"pr84169.c":                    {}, //TODO
		"pr84195.c":                    {}, //TODO
		"pr84305.c":                    {}, //TODO
		"pr84339.c":                    {}, //TODO
		"pr84425.c":                    {}, //TODO
		"pr84478.c":                    {}, //TODO
		"pr84524.c":                    {}, //TODO
		"pr84748.c":                    {}, //TODO
		"pr84960.c":                    {}, //TODO
		"pr85095.c":                    {}, //TODO
		"pr85156.c":                    {}, //TODO
		"pr85169.c":                    {}, //TODO
		"pr85331.c":                    {}, //TODO
		"pr85529-2.c":                  {}, //TODO
		"pr85582-2.c":                  {}, //TODO
		"pr85582-3.c":                  {}, //TODO
		"pr85756.c":                    {}, //TODO
		"pr86231.c":                    {}, //TODO
		"pr86492.c":                    {}, //TODO
		"pr86528.c":                    {}, //TODO
		"pr86659-1.c":                  {}, //TODO
		"pr86659-2.c":                  {}, //TODO
		"pr86844.c":                    {}, //TODO
		"pr87110.c":                    {}, //TODO
		"pr87290.c":                    {}, //TODO
		"pr88693.c":                    {}, //TODO
		"pr88714.c":                    {}, //TODO
		"pr88739.c":                    {}, //TODO
		"pr88904.c":                    {}, //TODO
		"pr89195.c":                    {}, //TODO
		"pr89280.c":                    {}, //TODO
		"pr89369.c":                    {}, //TODO
		"pr89634.c":                    {}, //TODO
		"pr89655.c":                    {}, //TODO
		"pr90025.c":                    {}, //TODO
		"pr90139.c":                    {}, //TODO
		"pr90949.c":                    {}, //TODO
		"pr91137.c":                    {}, //TODO
		"pr91450-1.c":                  {}, //TODO
		"pr91450-2.c":                  {}, //TODO
		"pr91597.c":                    {}, //TODO
		"pr91632.c":                    {}, //TODO
		"pr91635.c":                    {}, //TODO
		"pr92140.c":                    {}, //TODO
		"pr92618.c":                    {}, //TODO
		"pr92904.c":                    {}, //TODO
		"pr93213.c":                    {}, //TODO
		"pr93402.c":                    {}, //TODO
		"pr93582.c":                    {}, //TODO
		"pr93908.c":                    {}, //TODO
		"pr93945.c":                    {}, //TODO
		"pr94524-1.c":                  {}, //TODO
		"pr94524-2.c":                  {}, //TODO
		"pr94591.c":                    {}, //TODO
		"pr94734.c":                    {}, //TODO
		"pr95731.c":                    {}, //TODO
		"pr97073.c":                    {}, //TODO
		"pr97386-1.c":                  {}, //TODO
		"pr97386-2.c":                  {}, //TODO
		"pr97421-1.c":                  {}, //TODO
		"pr97888-2.c":                  {}, //TODO
		"pr98474.c":                    {}, //TODO
		"pr98681.c":                    {}, //TODO
		"pr98727.c":                    {}, //TODO
		"pr98853-1.c":                  {}, //TODO
		"pr98853-2.c":                  {}, //TODO
		"pr99079.c":                    {}, //TODO
		"printf-2.c":                   {}, //TODO
		"printf-chk-1.c":               {}, //TODO
		"printf-lib.c":                 {}, //TODO
		"printf.c":                     {}, //TODO
		"pta-field-1.c":                {}, //TODO
		"pta-field-2.c":                {}, //TODO
		"restrict-1.c":                 {}, //TODO
		"return-addr.c":                {}, //TODO
		"scal-to-vec1.c":               {}, //TODO
		"scal-to-vec2.c":               {}, //TODO
		"scal-to-vec3.c":               {}, //TODO
		"section.c":                    {}, //TODO
		"simd-1.c":                     {}, //TODO
		"simd-2.c":                     {}, //TODO
		"simd-3.c":                     {}, //TODO
		"simd-4.c":                     {}, //TODO
		"simd-5.c":                     {}, //TODO
		"simd-6.c":                     {}, //TODO
		"snprintf-chk-lib.c":           {}, //TODO
		"snprintf-chk.c":               {}, //TODO
		"sprintf-chk-lib.c":            {}, //TODO
		"sprintf-chk.c":                {}, //TODO
		"sprintf-lib.c":                {}, //TODO
		"sprintf.c":                    {}, //TODO
		"sra-1.c":                      {}, //TODO
		"stack-check-1.c":              {}, //TODO
		"stdarg-4.c":                   {}, //TODO
		"stkalign.c":                   {}, //TODO
		"stpcpy-chk-lib.c":             {}, //TODO
		"stpcpy-chk.c":                 {}, //TODO
		"stpcpy.c":                     {}, //TODO
		"stpncpy-chk-lib.c":            {}, //TODO
		"stpncpy-chk.c":                {}, //TODO
		"strcat-chk-lib.c":             {}, //TODO
		"strcat-chk.c":                 {}, //TODO
		"strcat-lib.c":                 {}, //TODO
		"strcat.c":                     {}, //TODO
		"strchr-lib.c":                 {}, //TODO
		"strchr.c":                     {}, //TODO
		"strcmp-1.c":                   {}, //TODO
		"strcmp-lib.c":                 {}, //TODO
		"strcmp.c":                     {}, //TODO
		"strcpy-1.c":                   {}, //TODO
		"strcpy-2-lib.c":               {}, //TODO
		"strcpy-2.c":                   {}, //TODO
		"strcpy-chk-lib.c":             {}, //TODO
		"strcpy-chk.c":                 {}, //TODO
		"strcpy-lib.c":                 {}, //TODO
		"strcpy.c":                     {}, //TODO
		"strcspn-lib.c":                {}, //TODO
		"strcspn.c":                    {}, //TODO
		"string-opt-5.c":               {}, //TODO
		"strlen-1.c":                   {}, //TODO
		"strlen-2-lib.c":               {}, //TODO
		"strlen-3-lib.c":               {}, //TODO
		"strlen-7.c":                   {}, //TODO
		"strlen-lib.c":                 {}, //TODO
		"strlen.c":                     {}, //TODO
		"strncat-chk-lib.c":            {}, //TODO
		"strncat-chk.c":                {}, //TODO
		"strncat-lib.c":                {}, //TODO
		"strncat.c":                    {}, //TODO
		"strncmp-1.c":                  {}, //TODO
		"strncmp-2-lib.c":              {}, //TODO
		"strncmp-2.c":                  {}, //TODO
		"strncmp-lib.c":                {}, //TODO
		"strncmp.c":                    {}, //TODO
		"strncpy-chk-lib.c":            {}, //TODO
		"strncpy-chk.c":                {}, //TODO
		"strncpy-lib.c":                {}, //TODO
		"strncpy.c":                    {}, //TODO
		"strnlen-lib.c":                {}, //TODO
		"strnlen.c":                    {}, //TODO
		"strpbrk-lib.c":                {}, //TODO
		"strpbrk.c":                    {}, //TODO
		"strpcpy-2-lib.c":              {}, //TODO
		"strpcpy-2.c":                  {}, //TODO
		"strpcpy-lib.c":                {}, //TODO
		"strrchr-lib.c":                {}, //TODO
		"strrchr.c":                    {}, //TODO
		"strspn-lib.c":                 {}, //TODO
		"strspn.c":                     {}, //TODO
		"strstr-asm-lib.c":             {}, //TODO
		"strstr-asm.c":                 {}, //TODO
		"strstr-lib.c":                 {}, //TODO
		"strstr.c":                     {}, //TODO
		"struct-aliasing-1.c":          {}, //TODO
		"struct-ini-2.c":               {}, //TODO
		"unalign-1.c":                  {}, //TODO
		"unsafe-fp-assoc-1.c":          {}, //TODO
		"user-printf.c":                {}, //TODO
		"usmul.c":                      {}, //TODO
		"va-arg-21.c":                  {}, //TODO
		"va-arg-22.c":                  {}, //TODO
		"va-arg-pack-1.c":              {}, //TODO
		"vector-1.c":                   {}, //TODO
		"vector-2.c":                   {}, //TODO
		"vector-3.c":                   {}, //TODO
		"vector-4.c":                   {}, //TODO
		"vfprintf-chk-1.c":             {}, //TODO
		"vprintf-chk-1.c":              {}, //TODO
		"vrp-7.c":                      {}, //TODO
		"vsnprintf-chk-lib.c":          {}, //TODO
		"vsnprintf-chk.c":              {}, //TODO
		"vsprintf-chk-lib.c":           {}, //TODO
		"vsprintf-chk.c":               {}, //TODO
		"widechar-1.c":                 {}, //TODO
		"zerolen-2.c":                  {}, //TODO
	}
	blacklistVNMakarov := map[string]struct{}{
		// #endif without #if
		"endif.c": {},

		"funnkuch-reduce.c":         {}, //TODO
		"0009-breakcont1.c":         {}, //TODO
		"0013-struct5.c":            {}, //TODO
		"0022-namespaces1.c":        {}, //TODO
		"anonymous-members.c":       {}, //TODO
		"anonymous-struct.c":        {}, //TODO
		"bitfield-basic.c":          {}, //TODO
		"bitfield-packing.c":        {}, //TODO
		"bitfield-reset-align.c":    {}, //TODO
		"bitfield-trailing-zero.c":  {}, //TODO
		"bitfield-types-init.c":     {}, //TODO
		"bitfield-types.c":          {}, //TODO
		"declaration-default-int.c": {}, //TODO
		"declarator-abstract.c":     {}, //TODO
		"expression.c":              {}, //TODO
		"for-empty-expr.c":          {}, //TODO
		"initialize-call.c":         {}, //TODO
		"initialize-object.c":       {}, //TODO
		"offsetof.c":                {}, //TODO
		"short-circuit-comma.c":     {}, //TODO
		"sizeof.c":                  {}, //TODO
		"struct-padding.c":          {}, //TODO
		"unary-plus.c":              {}, //TODO
		"enum_test.c":               {}, //TODO
		"fermian-2.c":               {}, //TODO
		"typedef-member-scope.c":    {}, //TODO
		"typedef.c":                 {}, //TODO
	}
	blacklictTCC := map[string]struct{}{
		"11.c": {}, // https://gcc.gnu.org/onlinedocs/gcc/Variadic-Macros.html#Variadic-Macros

		"02.c":                        {}, //TODO
		"03.c":                        {}, //TODO
		"04.c":                        {}, //TODO
		"06.c":                        {}, //TODO
		"07.c":                        {}, //TODO
		"08.c":                        {}, //TODO
		"09.c":                        {}, //TODO
		"10.c":                        {}, //TODO
		"14.c":                        {}, //TODO
		"15.c":                        {}, //TODO
		"17.c":                        {}, //TODO
		"18.c":                        {}, //TODO
		"19.c":                        {}, //TODO
		"20.c":                        {}, //TODO
		"21.c":                        {}, //TODO
		"pp-counter.c":                {}, //TODO
		"17_enum.c":                   {}, //TODO
		"39_typedef.c":                {}, //TODO
		"46_grep.c":                   {}, //TODO
		"76_dollars_in_identifiers.c": {}, //TODO
		"79_vla_continue.c":           {}, //TODO
		"81_types.c":                  {}, //TODO
		"82_attribs_position.c":       {}, //TODO
		"88_codeopt.c":                {}, //TODO
		"90_struct-init.c":            {}, //TODO
		"92_enum_bitfield.c":          {}, //TODO
		"93_integer_promotion.c":      {}, //TODO
		"94_generic.c":                {}, //TODO
		"95_bitfields.c":              {}, //TODO
		"95_bitfields_ms.c":           {}, //TODO
		"99_fastcall.c":               {}, //TODO
	}
	blacklistCxgo := map[string]struct{}{
		"forward enum.c":      {}, //TODO
		"inet.c":              {}, //TODO
		"literal statement.c": {}, //TODO
	}
	switch fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH) {
	case "linux/s390x":
		blacklistCompCert = map[string]struct{}{"aes.c": {}} // Unsupported endianness.
		fallthrough
	case "linux/arm", "linux/arm64":
		// Uses sse2 headers.
		blacklistGame["fannkuchredux-4.c"] = struct{}{}
		blacklistGame["mandelbrot-6.c"] = struct{}{}
		blacklistGame["nbody-4.c"] = struct{}{}
		blacklistGame["nbody-8.c"] = struct{}{}
		blacklistGame["nbody-9.c"] = struct{}{}
		blacklistGame["spectral-norm-5.c"] = struct{}{}
		blacklistGame["spectral-norm-6.c"] = struct{}{}
	case "freebsd/386", "darwin/amd64", "darwin/arm64", "freebsd/amd64":
		blacklistCompCert = map[string]struct{}{"aes.c": {}} // include file not found: "../endian.h"
	case "windows/amd64", "windows/386":
		blacklistCompCert = map[string]struct{}{"aes.c": {}} // include file not found: "../endian.h"
		blacklistCxgo["inet.c"] = struct{}{}                 // include file not found: <arpa/inet.h>
		blacklistGCC["loop-2f.c"] = struct{}{}               // include file not found: <sys/mman.h>
		blacklistGCC["loop-2g.c"] = struct{}{}               // include file not found: <sys/mman.h>
		blacklistGame["fasta-4.c"] = struct{}{}              // include file not found: <err.h>
		blacklistGame["pidigits-2.c"] = struct{}{}           // include file not found: <gmp.h>
		blacklistGame["pidigits-6.c"] = struct{}{}           // include file not found: <threads.h>
		blacklistGame["pidigits-9.c"] = struct{}{}           // include file not found: <gmp.h>
		blacklistGame["pidigits.c"] = struct{}{}             // include file not found: <gmp.h>
		blacklistGame["regex-redux-2.c"] = struct{}{}        // include file not found: <pcre.h>
		blacklistGame["regex-redux-3.c"] = struct{}{}        // include file not found: <pcre.h>
		blacklistGame["regex-redux-4.c"] = struct{}{}        // include file not found: <pcre.h>
		blacklistGame["regex-redux-5.c"] = struct{}{}        // include file not found: <pcre2.h>
	case "openbsd/amd64":
		blacklistCompCert = map[string]struct{}{"aes.c": {}} // include file not found: "../endian.h"
		blacklistGame["mandelbrot-7.c"] = struct{}{}         // include file not found: <omp.h>
		blacklistGame["pidigits-6.c"] = struct{}{}           // include file not found: <threads.h>
		blacklistGame["regex-redux-3.c"] = struct{}{}        // include file not found: <omp.h>
		blacklistGame["spectral-norm-4.c"] = struct{}{}      // include file not found: <omp.h>
	}
	var files, ok, skip, fails int
	for _, v := range []struct {
		cfg       *Config
		dir       string
		blacklist map[string]struct{}
	}{
		{cfg, "CompCert-3.6/test/c", blacklistCompCert},
		{cfg, "ccgo", nil},
		{cfg, "gcc-9.1.0/gcc/testsuite/gcc.c-torture", blacklistGCC},
		{cfg, "github.com/AbsInt/CompCert/test/c", blacklistCompCert},
		{cfg, "github.com/cxgo", blacklistCxgo},
		{cfg, "github.com/gcc-mirror/gcc/gcc/testsuite", blacklistGCC},
		{cfg, "github.com/vnmakarov", blacklistVNMakarov},
		//TODO {cfg, "sqlite-amalgamation-3370200", nil},
		{cfg, "tcc-0.9.27/tests", blacklictTCC},
		{cfgGame, "benchmarksgame-team.pages.debian.net", blacklistGame},
	} {
		t.Run(v.dir, func(t *testing.T) {
			f, o, s, n := testParse(t, v.cfg, "/"+v.dir, v.blacklist)
			files += f
			ok += o
			skip += s
			fails += n
		})
	}
	t.Logf("TOTAL: files %v, skip %v, ok %v, fails %v", files, skip, ok, fails)
}

func testParse(t *testing.T, cfg *Config, dir string, blacklist map[string]struct{}) (files, ok, skip, nfails int) {
	var fails []string
	p := newParallel()
	err := walk(dir, func(pth string, fi os.FileInfo) error {
		if fi.IsDir() {
			return nil
		}

		if filepath.Ext(pth) != ".c" {
			return nil
		}

		files++
		switch {
		case re != nil:
			if !re.MatchString(pth) {
				skip++
				return nil
			}
		default:
			if _, ok := blacklist[filepath.Base(pth)]; ok {
				skip++
				return nil
			}
		}

		// Preprocess( //TODO-
		// 	cfg,
		// 	[]Source{
		// 		{Name: "<predefined>", Value: predefined},
		// 		{Name: "<builtin>", Value: builtin},
		// 		{Name: pth, FS: cFS},
		// 	},
		// 	os.Stdout,
		// )
		p.exec(func() {
			if *oTrace {
				fmt.Fprintln(os.Stderr, pth)
			}

			var err error

			func() {
				defer func() {
					if e := recover(); e != nil && err == nil {
						err = fmt.Errorf("%v: PANIC", pth)
					}
				}()

				_, err = Parse(
					cfg,
					[]Source{
						{Name: "<predefined>", Value: predefined},
						{Name: "<builtin>", Value: builtin},
						{Name: pth, FS: cFS},
					},
				)
			}()
			p.Lock()

			defer p.Unlock()

			if err != nil {
				fails = append(fails, pth)
				t.Errorf("%v: %v", pth, err)
			} else {
				ok++
			}
		})
		return nil
	})
	if err != nil {
		t.Error(err)
	}

	p.wg.Wait()
	for _, v := range fails {
		t.Log(v)
	}
	// fmt.Fprintf(os.Stderr, "%v: files %v, skip %v, ok %v, fails %v\n", dir, files, skip, ok, len(fails))
	t.Logf("files %v, skip %v, ok %v, fails %v", files, skip, ok, len(fails))
	return files, ok, skip, len(fails)
}
