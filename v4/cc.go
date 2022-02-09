// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate rm -f ast.go
//go:generate yy -o /dev/null -position -astImport "\"fmt\"\n\n\"modernc.org/token\"" -prettyString PrettyString -kind Case -noListKind -noPrivateHelpers -forceOptPos parser.yy
//go:generate stringer -output stringer.go -type=tokCh
//go:generate sh -c "go test -run ^Example |fe"

// Package cc is a C99 compiler front end.
//
// Online documentation
//
// See https://godoc.org/modernc.org/cc/v4.
//
// Links
//
// Referenced from elsewhere:
//
//  [0]: http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf
//  [1]: https://www.spinellis.gr/blog/20060626/cpp.algo.pdf
//  [2]: https://jhjourdan.mketjh.fr/pdf/jourdan2017simple.pdf
package cc // import "modernc.org/cc/v4"

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"runtime"
	"strings"

	v3 "modernc.org/cc/v3"
)

var (
	isTesting  bool
	traceFails bool
)

// HostConfig returns the system C preprocessor/compiler configuration, or an
// error, if any.  The configuration is obtained by running the command named
// by the cpp argumnent or "cpp" when it's empty.  For the predefined macros
// list the '-dM' options is added. For the include paths lists, the option
// '-v' is added and the output is parsed to extract the "..." include and
// <...> include paths. To add any other options to cpp, list them in opts.
//
// The function relies on a POSIX/GCC compatible C preprocessor installed.
// Execution of HostConfig is not free, so caching of the results is
// recommended.
func HostConfig(cpp string, opts ...string) (predefined string, includePaths, sysIncludePaths []string, err error) { //TODO-
	if cpp == "" {
		cpp = "cpp"
	}
	args := append(append([]string{"-dM"}, opts...), os.DevNull)
	pre, err := exec.Command(cpp, args...).Output()
	if err != nil {
		return "", nil, nil, errorf("", err)
	}

	args = append(append([]string{"-v"}, opts...), os.DevNull)
	out, err := exec.Command(cpp, args...).CombinedOutput()
	if err != nil {
		return "", nil, nil, errorf("", err)
	}

	sep := "\n"
	if env("GOOS", runtime.GOOS) == "windows" {
		sep = "\r\n"
	}

	a := strings.Split(string(out), sep)
	for i := 0; i < len(a); {
		switch a[i] {
		case "#include \"...\" search starts here:":
		loop:
			for i = i + 1; i < len(a); {
				switch v := a[i]; {
				case strings.HasPrefix(v, "#") || v == "End of search list.":
					break loop
				default:
					includePaths = append(includePaths, strings.TrimSpace(v))
					i++
				}
			}
		case "#include <...> search starts here:":
			for i = i + 1; i < len(a); {
				switch v := a[i]; {
				case strings.HasPrefix(v, "#") || v == "End of search list.":
					return string(pre), includePaths, sysIncludePaths, nil
				default:
					sysIncludePaths = append(sysIncludePaths, strings.TrimSpace(v))
					i++
				}
			}
		default:
			i++
		}
	}
	return "", nil, nil, errorf("failed parsing the output of %s -v", cpp)
}

// NewConfig returns the system C compiler configuration, or an error, if
// any. The function will look for the compiler first in the environment
// variable CC, then it'll try other options. Usually that means looking for
// the "cc" and "gcc" binary, in that order.
//
// Execution of NewConfig is expensive, caching the results is recommended.
func NewConfig(goos, goarch string) (r *Config, err error) {
	cc, predefined, includePaths, sysIncludePaths, err := newConfig()
	if err != nil {
		return nil, fmt.Errorf("DefaultConfig: %v", err)
	}

	switch fmt.Sprintf("%s/%s", goos, goarch) {
	case "netbsd/amd64":
		sysIncludePaths = append(sysIncludePaths, "/usr/pkg/include")
	case "freebsd/386":
		sysIncludePaths = append(sysIncludePaths, "/usr/local/include")
	}
	sysIncludePaths = sysIncludePaths[:len(sysIncludePaths):len(sysIncludePaths)]
	includePaths = append([]string{""}, includePaths...)
	includePaths = append(includePaths, sysIncludePaths...)
	includePaths = includePaths[:len(includePaths):len(includePaths)]
	abi, err := NewABI(goos, goarch)
	if err != nil {
		return nil, err
	}

	return &Config{
		ABI:             abi,
		CC:              cc,
		IncludePaths:    includePaths,
		Predefined:      predefined,
		SysIncludePaths: sysIncludePaths,
	}, nil
}

func newConfig() (cc, predefined string, includePaths, sysIncludePaths []string, err error) {
	for _, cc = range []string{os.Getenv("CC"), "cc", "gcc"} {
		if cc == "" {
			continue
		}

		cc, err = exec.LookPath(cc)
		if err != nil {
			continue
		}

		pre, err := exec.Command(cc, "-dM", "-E", "-").Output()
		if err != nil {
			continue
		}

		out, err := exec.Command(cc, "-v", "-E", "-").CombinedOutput()
		if err != nil {
			continue
		}

		sep := "\n"
		if env("GOOS", runtime.GOOS) == "windows" {
			sep = "\r\n"
		}
		a := strings.Split(string(out), sep)
		for i := 0; i < len(a); {
			switch a[i] {
			case "#include \"...\" search starts here:":
			loop:
				for i = i + 1; i < len(a); {
					switch v := a[i]; {
					case strings.HasPrefix(v, "#") || v == "End of search list.":
						break loop
					default:
						includePaths = append(includePaths, strings.TrimSpace(v))
						i++
					}
				}
			case "#include <...> search starts here:":
				for i = i + 1; i < len(a); {
					switch v := a[i]; {
					case strings.HasPrefix(v, "#") || v == "End of search list.":
						return cc, string(pre), includePaths, sysIncludePaths, nil
					default:
						sysIncludePaths = append(sysIncludePaths, strings.TrimSpace(v))
						i++
					}
				}
			default:
				i++
			}
		}
	}
	return "", "", nil, nil, fmt.Errorf("cannot determine C compiler configuration")
}

// Source is a named part of a translation unit. The name argument is used for
// reporting Token positions.  The value argument can be a string, []byte,
// fs.File, io.ReadCloser, io.Reader or nil. If the value argument is nil an
// attempt is made to open a file using the name argument.
//
// When the value argument is an *os.File, io.ReadCloser or fs.File,
// Value.Close() is called before returning from Preprocess, Parse or
// Translate.
//
// If FS is not nil it overrides the FS from Config.
type Source struct {
	Name  string
	Value interface{}
	FS    fs.FS
}

// ABI describes selected parts of the Application Binary Interface.
type ABI struct {
	ByteOrder  binary.ByteOrder
	SignedChar bool
}

// NewABI creates an ABI for a given OS and architecture. The OS and
// architecture values are the same as used in Go. The ABI type map may miss
// advanced types like complex numbers, etc.
func NewABI(os, arch string) (*ABI, error) {
	abi, err := v3.NewABI(os, arch)
	if err != nil {
		return nil, err
	}

	return &ABI{abi.ByteOrder, abi.SignedChar}, nil
}

// Config configures the preprocessor, parser and type checker.
//
// Search paths listed in IncludePaths and SysIncludePaths are used to resolve
// #include "foo.h" and #include <foo.h> preprocessing directives respectively.
// A special search path "@" is interpreted as 'the same directory as where the
// file with the #include directive is'.
//
// If FS is nil, os.Open is used to open named files.
type Config struct {
	ABI             *ABI
	CC              string // The configured C compiler, filled by NewConfig.
	FS              fs.FS
	IncludePaths    []string
	PragmaHandler   func([]Token) error
	Predefined      string // The predefined macros from CC, filled by NewConfig.
	SysIncludePaths []string

	fakeIncludes bool // testing
}

type errors []string

// Error implements error.
func (e errors) Error() string { return strings.Join(e, "\n") }

func (e errors) err() error {
	w := 0
	for i, v := range e {
		if i != 0 {
			if prev, ok := extractPos(e[i-1]); ok {
				if cur, ok := extractPos(v); ok && prev.Filename == cur.Filename && prev.Line == cur.Line {
					continue
				}
			}
		}
		e[w] = v
		w++
	}
	e = e[:w]
	if len(e) == 0 {
		return nil
	}

	return e
}

// Preprocess preprocesses a translation unit, consisting of inputs in sources,
// and writes the result to w.
func Preprocess(cfg *Config, sources []Source, w io.Writer) (err error) {
	cpp, err := newCPP(cfg, sources, nil)
	if err != nil {
		return err
	}

	return preprocess(cpp, w)
}

func preprocess(cpp *cpp, w io.Writer) (err error) {
	var errors errors
	cpp.eh = func(msg string, args ...interface{}) { errors = append(errors, fmt.Sprintf(msg, args...)) }
	var prev rune
	for {
		if cpp.rune() == eof {
			return errors.err()
		}

		tok := cpp.shift()
		switch c := tok.Ch; {
		case
			// Prevent the textual form of certain adjacent tokens to form a "false" token,
			// not present in the source.
			c == '#' && prev == '#',
			c == '&' && prev == '&',
			c == '+' && prev == '+',
			c == '+' && prev == rune(PPNUMBER),
			c == '-' && prev == '-',
			c == '-' && prev == rune(PPNUMBER),
			c == '.' && prev == '.',
			c == '.' && prev == rune(PPNUMBER),
			c == '<' && prev == '<',
			c == '=' && prev == '!',
			c == '=' && prev == '%',
			c == '=' && prev == '&',
			c == '=' && prev == '*',
			c == '=' && prev == '+',
			c == '=' && prev == '/',
			c == '=' && prev == '<',
			c == '=' && prev == '=',
			c == '=' && prev == '^',
			c == '=' && prev == '|',
			c == '>' && prev == '-',
			c == '>' && prev == '>',
			c == '|' && prev == '|',
			c == rune(DEC) && prev == '-',
			c == rune(IDENTIFIER) && prev == rune(IDENTIFIER),
			c == rune(INC) && prev == '+':

			if _, err := w.Write(sp); err != nil {
				return err
			}
		}
		if prev == ' ' && tok.Ch == ' ' {
			continue
		}

		prev = tok.Ch
		if _, err = w.Write(tok.Src()); err != nil {
			return err
		}
	}
}

// Parse preprocesses and parsees a translation unit, consisting of inputs in
// sources, and writes the result to w.
func Parse(cfg *Config, sources []Source) (*AST, error) {
	p, err := newParser(cfg, sources)
	if err != nil {
		return nil, err
	}

	return p.parse()
}
