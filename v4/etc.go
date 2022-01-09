// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"modernc.org/strutil"
)

var (
	extendedErrors = true // true: Errors will include origin info.
)

// origin returns caller's short position, skipping skip frames.
func origin(skip int) string {
	pc, fn, fl, _ := runtime.Caller(skip)
	f := runtime.FuncForPC(pc)
	var fns string
	if f != nil {
		fns = f.Name()
		if x := strings.LastIndex(fns, "."); x > 0 {
			fns = fns[x+1:]
		}
		if strings.HasPrefix(fns, "func") {
			num := true
			for _, c := range fns[len("func"):] {
				if c < '0' || c > '9' {
					num = false
					break
				}
			}
			if num {
				return origin(skip + 2)
			}
		}
	}
	return fmt.Sprintf("%s:%d:%s", filepath.Base(fn), fl, fns)
}

// todo prints and return caller's position and an optional message tagged with TODO. Output goes to stderr.
func todo(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	r := fmt.Sprintf("%s\n\tTODO %s", origin(2), s)
	fmt.Fprintf(os.Stderr, "%s\n", r)
	os.Stdout.Sync()
	return r
}

// trc prints and return caller's position and an optional message tagged with TRC. Output goes to stderr.
func trc(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	r := fmt.Sprintf("%s: TRC %s", origin(2), s)
	fmt.Fprintf(os.Stderr, "%s\n", r)
	os.Stderr.Sync()
	return r
}

// errorf constructs and error value. If extendedErrors is true, the error will
// continue its origin.
func errorf(s string, args ...interface{}) error {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	switch {
	case extendedErrors:
		return fmt.Errorf("%s (%v:)", s, origin(2))
	default:
		return fmt.Errorf("%s", s)
	}
}

// printHooks configure strutil.PrettyString for pretty printing Token values.
var printHooks = strutil.PrettyPrintHooks{
	reflect.TypeOf((*Token)(nil)): func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		t := v.(*Token)
		if t == nil {
			return
		}

		f.Format(prefix)
		if p := t.Position(); p.IsValid() {
			f.Format("%v: ", p)
		}
		f.Format("%s %q", t.Name(), t.Src())
		f.Format(suffix)
	},
}

// PrettyString returns a formatted representation of Tokens and AST nodes.
func PrettyString(v interface{}) string {
	return strutil.PrettyString(v, "", "", printHooks)
}

// runeName returns a human readable representation of ch.
func runeName(ch rune) string {
	switch {
	case ch < 0:
		return "<EOF>"
	case ch >= unicodePrivateAreaFirst && ch <= unicodePrivateAreaLast:
		return tokCh(ch).String()
	default:
		return fmt.Sprintf("%+q", ch)
	}
}

// env returns os.Getenv("key") or defaultVal if getenv returns an empty string.
func env(key, defaultVal string) string {
	if s := os.Getenv(key); s != "" {
		return s
	}

	return defaultVal
}

func toksDump(v interface{}) string {
	var a []string
	switch x := v.(type) {
	case []cppToken:
		return toksDump(cppTokens(x))
	case *cppTokens:
		return toksDump(*x)
	case cppTokens:
		if len(x) != 0 {
			p := x[0].Position()
			p.Filename = filepath.Base(p.Filename)
			a = append(a, p.String())
		}
		for _, v := range x {
			s := string(v.Src())
			// if hs := v.hs.String(); hs != "[]" {
			// 	s = fmt.Sprintf("%s^%s", s, hs)
			// }
			a = append(a, s)
		}
	case []cppTokens:
		var a []string
		for _, v := range x {
			a = append(a, toksDump(v))
		}
		return fmt.Sprintf("%v.%d", a, len(a))
	case controlLine:
		return toksDump([]Token(x))
	case *tokenizer:
		t := x.peek(0)
		return fmt.Sprintf("[%T %v]", x, &t)
	case []Token:
		for _, v := range x {
			a = append(a, string(v.Src()))
		}
	case *cat:
		return fmt.Sprintf("(%s · %s)", toksDump(x.head), toksDump(x.tail))
	case textLine:
		return toksDump([]Token(x))
	default:
		panic(todo("%T", x))
	}
	return fmt.Sprintf("%q.%d", a, len(a))
}

func tokens2CppTokens(s []Token, skipFirstSep bool) (r []cppToken) {
	for i, v := range s {
		if (i != 0 || !skipFirstSep) && len(v.Sep()) != 0 {
			r = append(r, cppToken{spTok, nil})
		}
		r = append(r, cppToken{v, nil})
	}
	return r
}

func preprocessingTokens2Tokens(s []cppToken) (r []Token) {
	for _, v := range s {
		if v.Ch != ' ' {
			r = append(r, v.Token)
		}
	}
	return r
}

func toksTrim(s cppTokens) cppTokens {
	for len(s) != 0 && s[0].Ch == ' ' {
		s = s[1:]
	}
	for len(s) != 0 && s[len(s)-1].Ch == ' ' {
		s = s[:len(s)-1]
	}
	return s
}
