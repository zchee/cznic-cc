// Copyright 2022 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"modernc.org/cc/v4"
	"modernc.org/opt"
)

const (
	builtin = `

#define __extension__
#define __restrict_arr restrict

#ifndef __builtin_va_list
#define __builtin_va_list __builtin_va_list
typedef void *__builtin_va_list;
#endif

#ifndef __builtin_va_arg
#define __builtin_va_arg __builtin_va_arg
#define __builtin_va_arg(va, type) (*(type*)__builtin_va_arg_impl(va))
#endif

#define __builtin_offsetof(type, member) ((size_t)&(((type*)0)->member))
#define __builtin_types_compatible_p(t1, t2) __builtin_types_compatible_p_impl()

#ifdef __SIZE_TYPE__
typedef __SIZE_TYPE__ size_t;
#else
#error __SIZE_TYPE__ undefined
#endif

#ifdef __WCHAR_TYPE__
typedef __WCHAR_TYPE__ wchar_t;
#else
#error __WCHAR_TYPE__ undefined
#endif

#ifdef __PTRDIFF_TYPE__
typedef __PTRDIFF_TYPE__ ptrdiff_t;
#else
#error __PTRDIFF_TYPE__ undefined
#endif

#define __FUNCTION__ __func__
#define __PRETTY_FUNCTION__ __func__

#ifdef __clang__
#define __builtin_convertvector(src, type) (*(type*)&src)
#endif

`
)

func fail(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

func main() {
	fn := os.Getenv("FAKE_CC_LOG")
	if fn == "" {
		fail("FAKE_CC_LOG not set\n")
	}

	CC := os.Getenv("FAKE_CC_CC")
	if CC == "" {
		fail("FAKE_CC_CC not set\n")
	}

	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_SYNC, 0660)
	if err != nil {
		fail("%v\n", err)
	}

	report2 := func(result, more string, rc int) {
		wd, err := os.Getwd()
		if err != nil {
			fail("%v\n", err)
		}

		a := append([]string{result, wd}, os.Args[1:]...)
		if more != "" {
			a = append(a, more)
		}
		var b bytes.Buffer
		if err := json.NewEncoder(&b).Encode(a); err != nil {
			fail("%v\n", err)
		}

		if _, err := f.Write(b.Bytes()); err != nil {
			fail("%v\n", err)
		}

		if rc >= 0 {
			os.Exit(rc)
		}
	}
	report := func(result string, rc int) { report2(result, "", rc) }

	cmd := exec.Command(CC, os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		report("SKIP", err.(*exec.ExitError).ExitCode())
	}

	set := opt.NewSet()
	var D, U, I []string
	var E bool
	set.Arg("D", true, func(opt, arg string) error { D = append(D, arg); return nil })
	set.Arg("I", true, func(opt, arg string) error { I = append(I, arg); return nil })
	set.Arg("U", true, func(opt, arg string) error { U = append(U, arg); return nil })
	set.Opt("E", func(opt string) error { E = true; return nil })
	var inputs []string
	if err := set.Parse(os.Args[1:], func(arg string) error {
		if strings.HasPrefix(arg, "-") {
			return nil
		}

		if strings.HasSuffix(arg, ".c") || strings.HasSuffix(arg, ".h") {
			inputs = append(inputs, arg)
		}

		return nil
	}); err != nil {
		fail("%v\n", err)
	}

	if len(inputs) == 0 {
		return
	}

	I = I[:len(I):len(I)]
	os.Setenv("CC", CC)
	cfg, err := cc.NewConfig(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		fail("%v\n", err)
	}

	cfg.IncludePaths = append([]string{""}, I...)
	cfg.IncludePaths = append(cfg.IncludePaths, cfg.HostIncludePaths...)
	cfg.IncludePaths = append(cfg.IncludePaths, cfg.HostSysIncludePaths...)
	cfg.SysIncludePaths = append(I, cfg.HostSysIncludePaths...)
	defs := buildDefs(D, U)
	for _, v := range inputs {
		if fi, err := os.Stat(v); err != nil || fi.Mode()&os.ModeIrregular != 0 {
			continue
		}

		sources := []cc.Source{
			{Name: "<predefined>", Value: cfg.Predefined},
			{Name: "<builtin>", Value: builtin},
		}
		if defs != "" {
			sources = append(sources, cc.Source{Name: "<command-line>", Value: defs})
		}
		sources = append(sources, cc.Source{Name: v})
		switch {
		case E:
			if err := cc.Preprocess(cfg, sources, io.Discard); err != nil {
				report("FAIL", -1)
				continue
			}
		default:
			if _, err := cc.Translate(cfg, sources); err != nil {
				src, _ := os.ReadFile(v)
				report2("FAIL", fmt.Sprintf("\nsrc\n----\n%s\n----\n%s", src, err), -1)
				continue
			}
		}

		report("PASS", -1)
	}
}

func buildDefs(D, U []string) string {
	var a []string
	for _, v := range D {
		if i := strings.IndexByte(v, '='); i > 0 {
			a = append(a, fmt.Sprintf("#define %s %s", v[:i], v[i+1:]))
			continue
		}

		a = append(a, fmt.Sprintf("#define %s 1", v))
	}
	for _, v := range U {
		a = append(a, fmt.Sprintf("#undef %s", v))
	}
	return strings.Join(a, "\n")
}
