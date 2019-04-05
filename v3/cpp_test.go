// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestCPPExpand(t *testing.T) {
	var re *regexp.Regexp
	if s := *oRE; s != "" {
		re = regexp.MustCompile(s)
	}

	cfg := &Config{fakeIncludes: true}
	if err := filepath.Walk(filepath.Join(testWD, filepath.FromSlash("testdata/cpp-expand/")), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || (!strings.HasSuffix(path, ".c") && !strings.HasSuffix(path, ".h")) {
			return nil
		}

		if re != nil && !re.MatchString(path) {
			return nil
		}

		ctx := newContext(cfg)
		cf, err := cache.getFile(ctx, path)
		if err != nil {
			return err
		}

		cpp := newCPP(ctx)
		var b strings.Builder
		expParth := path + ".expect"
		for line := range cpp.translationPhase4([]source{cf}) {
			for _, tok := range *line {
				b.WriteString(tok.String())
			}
			token4Pool.Put(line)
		}

		switch {
		case strings.Contains(filepath.ToSlash(path), "/mustfail/"):
			if err := ctx.Err(); err != nil {
				return nil
			}

			t.Fatalf("unexpected success: %s", path)
		default:
			if err := ctx.Err(); err != nil {
				t.Error(err)
			}
		}

		exp, err := ioutil.ReadFile(expParth)
		if err != nil {
			t.Error(err)
		}

		if g, e := b.String(), string(exp); g != e {
			a := strings.Split(g, "\n")
			b := strings.Split(e, "\n")
			n := len(a)
			if len(b) > n {
				n = len(b)
			}
			for i := 0; i < n; i++ {
				var x, y string
				if i < len(a) {
					x = a[i]
				}
				if i < len(b) {
					y = b[i]
				}
				if x != y {
					t.Errorf("%s:%v: %v", path, i+1, cmp.Diff(y, x))
				}
			}
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}

func TestTranslationPhase4(t *testing.T) {
	t.Run("shell.c", func(t *testing.T) { testTranslationPhase4(t, testPredefSource, testShellSource) })
	t.Run("shell.c/gnu", func(t *testing.T) { testTranslationPhase4(t, testPredefGNUSource, testShellSource) })
	t.Run("sqlite3.c", func(t *testing.T) { testTranslationPhase4(t, testPredefSource, testSQLiteSource) })
	t.Run("sqlite3.c/gnu", func(t *testing.T) { testTranslationPhase4(t, testPredefGNUSource, testSQLiteSource) })
}

func testTranslationPhase4(t *testing.T, predef, src source) {
	sources := []source{predef, testBuiltinSource, src}
	cfg := &Config{}
	ctx := newContext(cfg)
	ctx.includePaths = testIncludes
	ctx.sysIncludePaths = testSysIncludes
	cpp := newCPP(ctx)
	var m0, m1 runtime.MemStats
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m0)
	t0 := time.Now()
	for line := range cpp.translationPhase4(sources) {
		token4Pool.Put(line)
	}
	if err := ctx.Err(); err != nil {
		t.Error(err)
	}
	d := time.Since(t0)
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m1)
	t.Logf("sources %v, bytes %v, %v, %v B/s, mem %v",
		h(ctx.tuSources), h(ctx.tuSize), d, h(float64(time.Second)*float64(ctx.tuSize)/float64(d)), h(m1.Alloc-m0.Alloc))
}

func BenchmarkTranslationPhase4(b *testing.B) {
	b.Run("shell.c", func(b *testing.B) { benchmarkTranslationPhase4(b, testPredefSource, testShellSource) })
	b.Run("shell.c/gnu", func(b *testing.B) { benchmarkTranslationPhase4(b, testPredefGNUSource, testShellSource) })
	b.Run("sqlite3.c", func(b *testing.B) { benchmarkTranslationPhase4(b, testPredefSource, testSQLiteSource) })
	b.Run("sqlite3.c/gnu", func(b *testing.B) { benchmarkTranslationPhase4(b, testPredefGNUSource, testSQLiteSource) })
}

func benchmarkTranslationPhase4(b *testing.B, predef, src source) {
	sources := []source{predef, testBuiltinSource, src}
	cfg := &Config{}
	var ctx *context
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx = newContext(cfg)
		ctx.includePaths = testIncludes
		ctx.sysIncludePaths = testSysIncludes
		cpp := newCPP(ctx)
		for line := range cpp.translationPhase4(sources) {
			token4Pool.Put(line)
		}
		if err := ctx.Err(); err != nil {
			b.Error(err)
		}
	}
	b.SetBytes(ctx.tuSize)
}
