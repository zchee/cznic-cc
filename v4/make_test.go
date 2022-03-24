// Copyright 2022 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || freebsd || linux || netbsd || openbsd
// +build darwin freebsd linux netbsd openbsd

package cc // import "modernc.org/cc/v4"

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"modernc.org/opt"
)

func TestMake(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	oldCC := os.Getenv("CC")

	defer os.Setenv("CC", oldCC)

	tmp := t.TempDir()
	cc := filepath.Join(tmp, "fakecc")
	if runtime.GOOS == "windows" {
		cc += ".exe"
	}
	mustShell(t, "go", "build", "-o", cc, "fakecc.go")
	os.Setenv("CC", cc)
	os.Setenv("FAKE_CC_CC", defaultCfg().CC)
	var files, ok, skip, fails int32
	for _, v := range []struct {
		archive string
		dir     string
	}{
		{"ftp.pcre.org/pub/pcre.tar.gz", "pcre"},
		{"ftp.pcre.org/pub/pcre2.tar.gz", "pcre2"},
		{"github.com/madler/zlib.tar.gz", "zlib"},
		//TODO {"sourceforge.net/projects/tcl/files/Tcl/tcl.tar.gz", "tcl/unix"},
	} {
		t.Run(v.dir, func(t *testing.T) {
			f, o, s, n := testMake(t, "assets/"+v.archive, v.dir)
			files += f
			ok += o
			skip += s
			fails += n
		})
	}
	t.Logf("TOTAL: files %v, skip %v, ok %v, fails %v", files, skip, ok, fails)
}

func testMake(t *testing.T, archive, dir string) (files, ok, skip, nfails int32) {
	tmp := t.TempDir()
	mustUntarFile(t, tmp, archive, nil)
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	newWd := filepath.Join(tmp, dir)
	if err := os.Chdir(newWd); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := os.Chdir(wd); err != nil {
			t.Fatal(err)
		}
	}()

	os.Setenv("FAKE_CC_LOG", os.DevNull)
	mustShell(t, "./configure")
	os.Setenv("FAKE_CC_LOG", filepath.Join(newWd, "_fake_cc.log"))
	mustShell(t, "make")
	logf, err := os.Open(os.Getenv("FAKE_CC_LOG"))
	if err != nil {
		t.Fatal(err)
	}

	defer logf.Close()

	sc := bufio.NewScanner(logf)
	var lines [][]string
	for sc.Scan() {
		b := sc.Bytes()
		var a []string
		if err := json.NewDecoder(bytes.NewReader(b)).Decode(&a); err != nil {
			t.Fatal(err)
		}

		lines = append(lines, a)
	}
	if sc.Err() != nil {
		t.Fatal(err)
	}

next:
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}

		wd := line[0]
		set := opt.NewSet()
		var D, U, I []string
		set.Arg("D", true, func(opt, arg string) error { D = append(D, arg); return nil })
		set.Arg("I", true, func(opt, arg string) error { I = append(I, arg); return nil })
		set.Arg("U", true, func(opt, arg string) error { U = append(U, arg); return nil })
		var inputs []string
		if err := set.Parse(line[1:], func(arg string) error {
			if strings.HasPrefix(arg, "-") {
				return nil
			}

			if strings.HasSuffix(arg, ".c") || strings.HasSuffix(arg, ".h") {
				inputs = append(inputs, arg)
			}

			return nil
		}); err != nil {
			t.Fatal(err)
		}
		if len(inputs) == 0 {
			continue
		}

		I = I[:len(I):len(I)]
		cfg := defaultCfg()
		cfg.IncludePaths = append([]string{""}, I...)
		cfg.IncludePaths = append(cfg.IncludePaths, cfg.HostIncludePaths...)
		cfg.IncludePaths = append(cfg.IncludePaths, cfg.HostSysIncludePaths...)
		cfg.SysIncludePaths = append(I, cfg.HostSysIncludePaths...)
		sources := []Source{
			{Name: "<predefined>", Value: cfg.Predefined},
			{Name: "<builtin>", Value: builtin},
		}
		if s := buildDefs(D, U); s != "" {
			sources = append(sources, Source{Name: "<command-line>", Value: s})
		}
		for _, v := range inputs {
			if fi, err := os.Stat(v); err != nil || fi.Mode()&os.ModeIrregular != 0 {
				continue next
			}

			sources = append(sources, Source{Name: v})
		}
		if err := os.Chdir(wd); err != nil {
			t.Fatal(err)
		}

		dir, _ := filepath.Rel(tmp, wd)
		if *oTrace {
			t.Logf("in %q -I %q -D %q -U %q %q", dir, I, D, U, inputs)
		}
		files++
		_, err := Translate(cfg, sources)
		if err != nil {
			t.Error(err)
			nfails++
			continue
		}

		ok++
	}
	t.Logf("files %v, skip %v, ok %v, fails %v", files, skip, ok, nfails)
	return files, ok, skip, nfails
}
