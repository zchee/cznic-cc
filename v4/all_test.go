// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"modernc.org/ccorpus"
	"modernc.org/httpfs"
	"modernc.org/token"
)

// Produce the AST used in examples documentation.
func exampleAST(rule int, src string) interface{} {
	return "TODO"
	// src = strings.Replace(src, "\\n", "\n", -1)
	// cfg := &Config{ignoreErrors: true, PreprocessOnly: true}
	// ctx := newContext(cfg)
	// ctx.keywords = gccKeywords
	// ast, _ := parse(ctx, nil, nil, []Source{{Name: "example.c", Value: src, DoNotCache: true}})
	// if ast == nil {
	// 	return "FAIL"
	// }

	// pc, _, _, _ := runtime.Caller(1)
	// typ := runtime.FuncForPC(pc - 1).Name()
	// i := strings.LastIndexByte(typ, '.')
	// typ = typ[i+1+len("Example"):]
	// i = strings.LastIndexByte(typ, '_')
	// typ = typ[:i]
	// var node Node
	// depth := mathutil.MaxInt
	// findNode(typ, ast.TranslationUnit, 0, &node, &depth)
	// return node
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

func testScannerSource(t *testing.T, name string, value interface{}, exp []byte, mustErr bool) {
	ss, err := newScannerSource(name, value)
	if err != nil != mustErr {
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
	s, err := newScannerSource("test", `abc
def
 ghi
`)
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
		{newToken(s, -1, 13, 13, 0), 3, 6, -1, "", ""}, // EOF token
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

func TestCPPScan(t *testing.T) {
	fs := ccorpus.FileSystem()
	var files, tokens, chars int
	walk(fs, "/", func(pth string, fi os.FileInfo) error {
		if fi.IsDir() {
			return nil
		}

		switch filepath.Ext(pth) {
		case ".c", ".h":
			f, err := fs.Open(pth)
			if err != nil {
				t.Fatalf("%q: %v", pth, err)
			}

			fi, err := f.Stat()
			if err != nil {
				t.Fatalf("%q: %v", pth, err)
			}

			chars += int(fi.Size())
			src, err := newScannerSource(pth, f)
			if err != nil {
				t.Fatalf("%q: %v", pth, err)
			}

			s := newScanner(src, func(pos token.Position, msg string, args ...interface{}) {
				t.Fatalf("%v: %v", pos, fmt.Sprintf(msg, args...))
			})
			files++
			for {
				tok := s.cppScan()
				if tok.Ch < 0 {
					break
				}

				tokens++
			}
		}
		return nil
	})
	t.Logf("files %v, tokens %v, chars %v", files, tokens, chars)
}

func walk(fs *httpfs.FileSystem, dir string, f func(pth string, fi os.FileInfo) error) error {
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	root, err := fs.Open(dir)
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
			if err = walk(fs, v.Name(), f); err != nil {
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
