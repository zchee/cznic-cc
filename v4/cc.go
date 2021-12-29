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
package cc // import "modernc.org/cc/v4"

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
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
func HostConfig(cpp string, opts ...string) (predefined string, includePaths, sysIncludePaths []string, err error) {
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

// Source is a named part of a translation unit. The name argument is used for
// reporting Token positions.  The value argument can be a string, []byte,
// fs.File, io.ReadCloser, io.Reader or nil. If the value argument is nil an
// attempt is made to open a file using the name argument.
//
// When the value argument is an *os.File, io.ReadCloser or fs.File,
// value.Close() is called before returning.
type Source struct {
	Name  string
	Value interface{}
}
