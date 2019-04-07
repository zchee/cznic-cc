// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//TODO testCC

//go:generate rm -f lexer.go
//go:generate golex -o lexer.go lexer.l

//go:generate rm -f ast.go
//go:generate yy -o /dev/null -position -astImport "\"fmt\"\n\n\"modernc.org/token\"" -prettyString PrettyString -kind Case -noListKind -noPrivateHelpers -forceOptPos parser.yy
//go:generate sh -c "go test -run ^Example |fe"

// Package cc is a C99 compiler front end.
//
// Links
//
// Referenced from elsewhere:
//
//  [0]: http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf
//  [1]: https://www.spinellis.gr/blog/20060626/cpp.algo.pdf
//  [2]: http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf
//  [3]: http://gallium.inria.fr/~fpottier/publis/jourdan-fpottier-2016.pdf
//  [4]: https://gcc.gnu.org/onlinedocs/gcc-8.3.0/gcc/Attribute-Syntax.html#Attribute-Syntax
package cc // import "modernc.org/cc/v3"

import (
	"fmt"
	goscanner "go/scanner"
	gotoken "go/token"
	"math"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"

	"modernc.org/strutil"
	"modernc.org/token"
)

const (
	scopeParent StringID = -1
)

var (
	cache       = newPPCache()
	dict        = newDictionary()
	dictStrings [math.MaxUint8 + 1]string
	isTesting   bool

	token4Pool = sync.Pool{New: func() interface{} { r := make([]token4, 0); return &r }} //DONE benchmrk tuned capacity
	tokenPool  = sync.Pool{New: func() interface{} { r := make([]Token, 0); return &r }}  //DONE benchmrk tuned capacity

	printHooks = strutil.PrettyPrintHooks{
		reflect.TypeOf(Token{}): func(f strutil.Formatter, v interface{}, prefix, suffix string) {
			t := v.(Token)
			if (t == Token{}) {
				return
			}

			f.Format(prefix)
			r := t.Rune
			if p := t.Position(); p.IsValid() {
				f.Format("%v: ", p)
			}
			s := tokName(r)
			if x := s[0]; x >= '0' && x <= '9' {
				s = strconv.QuoteRune(r)
			}
			f.Format("%s", s)
			if s := t.Value.String(); len(s) != 0 {
				f.Format(" %q", s)
			}
			f.Format(suffix)
		},
	}
)

// PrettyString returns a formatted representation of things produced by this package.
func PrettyString(v interface{}) string {
	return strutil.PrettyString(v, "", "", printHooks)
}

// StringID is a process-unique string numeric identifier. Its zero value
// represents an empty string.
type StringID int32

// String implements fmt.Stringer.
func (n StringID) String() (r string) {
	if n < 256 {
		return dictStrings[byte(n)]
	}

	dict.mu.RLock()
	r = dict.strings[n]
	dict.mu.RUnlock()
	return r
}

// Node is implemented by Token and all AST nodes.
type Node interface {
	Position() token.Position
}

type noder struct{}

func (noder) Position() token.Position { panic("internal error") } //TODOOK

// Scope maps identifiers to definitions.
type Scope map[StringID][]Node

func (s *Scope) new() (r Scope) {
	if *s == nil {
		*s = map[StringID][]Node{}
	}
	r = Scope{scopeParent: []Node{struct {
		noder
		Scope
	}{Scope: *s}}}
	return r
}

func (s *Scope) declare(nm StringID, n Node) {
	if *s == nil {
		*s = map[StringID][]Node{nm: {n}}
		// dbg("declared %s at %v in scope %p", nm, n.Position(), *s)
		return
	}

	(*s)[nm] = append((*s)[nm], n)
	// t := ""
	// if x, ok := n.(*Declarator); ok && x.IsTypedefName {
	// 	t = ", typedefname"
	// }
	// dbg("declared %s%s at %v in scope %p", nm, t, n.Position(), *s)
}

// Parent returns s's outer scope, if any.
func (s *Scope) Parent() Scope {
	if *s == nil {
		return nil
	}

	if x, ok := (*s)[scopeParent]; ok {
		return x[0].(struct {
			noder
			Scope
		}).Scope
	}

	return nil
}

// Config3 amends behavior of translation phases 1 to 3.
type Config3 struct {
	WorkingDir string // Overrides os.Getwd if non empty.

	MaxSourceLine int // Zero: Scanner will use default buffer. Non zero: Scanner will use max(default buffer size, MaxSourceLine).

	PreserveWhiteSpace                      bool // Including also comments.
	RejectAnonymousFields                   bool // Pedantic: do not silently accept "struct{int;}".
	RejectCaseRange                         bool // Pedantic: do not silently accept "case 'a'...'z':".
	RejectElseExtraTokens                   bool // Pedantic: do not silently accept "#else foo".
	RejectEmptyCompositeLiterals            bool // Pedantic: do not silently accept "foo = (T){}".
	RejectEmptyDeclarations                 bool // Pedantic: do not silently accept "int foo(){};".
	RejectEmptyInitializerList              bool // Pedantic: do not silently accept "foo f = {};".
	RejectEmptyStructs                      bool // Pedantic: do not silently accept "struct foo {};".
	RejectEndifExtraTokens                  bool // Pedantic: do not silently accept "#endif foo".
	RejectFinalBackslash                    bool // Pedantic: do not silently accept "foo\\\n".
	RejectFunctionMacroEmptyReplacementList bool // Pedantic: do not silently accept "#define foo(bar)\n".
	RejectIfdefExtraTokens                  bool // Pedantic: do not silently accept "#ifdef foo bar".
	RejectIfndefExtraTokens                 bool // Pedantic: do not silently accept "#ifndef foo bar".
	RejectIncludeNext                       bool // Pedantic: do not silently accept "#include_next".
	RejectInvalidVariadicMacros             bool // Pedantic: do not silently accept "#define foo(bar...)". Standard allows only #define foo(bar, ...)
	RejectLabelValues                       bool // Pedantic: do not silently accept "foo: bar(); void *ptr = &&foo;" or "goto *ptr".
	RejectLineExtraTokens                   bool // Pedantic: do not silently accept "#line 1234 \"foo.c\" bar".
	RejectMissingConditionalExpr            bool // Pedantic: do not silently accept "foo = bar ? : baz;".
	RejectMissingDeclarationSpecifiers      bool // Pedantic: do not silently accept "main() {}".
	RejectMissingFinalNewline               bool // Pedantic: do not silently accept "foo\nbar".
	RejectMissingFinalStructFieldSemicolon  bool // Pedantic: do not silently accept "struct{int i; int j}".
	RejectNestedFunctionDefinitions         bool // Pedantic: do not silently accept nested function definitons.
	RejectParamSemicolon                    bool // Pedantic: do not silently accept "int f(int a; int b)".
	RejectStatementExpressions              bool // Pedantic: do not silently accept "i = ({foo();})".
	RejectTypeof                            bool // Pedantic: do not silently accept "typeof foo" or "typeof(bar*)".
	RejectUndefExtraTokens                  bool // Pedantic: do not silently accept "#undef foo bar".
}

// Config amend behavior of translation phase 4 and above. Instances of Config
// are not mutated by this package and it's safe to share/reuse them.
//
type Config struct {
	Config3

	MaxErrors int // 0: default (10), < 0: unlimited, n: n.

	fakeIncludes               bool // Testing only.
	ignoreErrors               bool // Testing only.
	ignoreIncludes             bool // Testing only.
	ignoreUndefinedIdentifiers bool // Testing only.
}

type context struct {
	cfg *Config
	goscanner.ErrorList
	includePaths    []string
	maxErrors       int
	mu              sync.Mutex
	sysIncludePaths []string
	tuSize          int64 // Sum of sizes of processed inputs
	tuSources       int   // Number of processed inputs
}

func newContext(cfg *Config) *context {
	maxErrors := cfg.MaxErrors
	if maxErrors == 0 {
		maxErrors = 10
	}
	return &context{
		maxErrors: maxErrors,
		cfg:       cfg,
	}
}

func (c *context) err(pos token.Position, msg string, args ...interface{}) (stop bool) {
	if c.cfg.ignoreErrors {
		return false
	}

	s := fmt.Sprintf(msg, args...)
	c.mu.Lock()
	max := c.maxErrors
	switch {
	case max < 0 || max > len(c.ErrorList):
		c.ErrorList.Add(gotoken.Position(pos), s)
	default:
		stop = true
	}
	c.mu.Unlock()
	return stop
}

func (c *context) errs(list goscanner.ErrorList) (stop bool) {
	c.mu.Lock()

	defer c.mu.Unlock()

	max := c.maxErrors
	for _, v := range list {
		switch {
		case max < 0 || max > len(c.ErrorList):
			c.ErrorList = append(c.ErrorList, v)
		default:
			return true
		}
	}
	return false
}

func (c *context) Err() error {
	c.mu.Lock()
	switch x := c.ErrorList.Err().(type) {
	case goscanner.ErrorList:
		x = append(goscanner.ErrorList(nil), x...)
		c.mu.Unlock()
		sort.Slice(x, func(i, j int) bool {
			a := x[i]
			b := x[j]
			if a.Pos.Filename < b.Pos.Filename {
				return true
			}

			if a.Pos.Filename > b.Pos.Filename {
				return false
			}

			if a.Pos.Line < b.Pos.Line {
				return true
			}

			if a.Pos.Line > b.Pos.Line {
				return false
			}

			return a.Pos.Column < b.Pos.Column
		})
		a := make([]string, 0, len(x))
		for _, v := range x {
			s := v.Error()
			if n := len(a); n == 0 || a[n-1] != s {
				a = append(a, s)
			}
		}
		return fmt.Errorf("%s", strings.Join(a, "\n"))
	default:
		c.mu.Unlock()
		return x
	}
}

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
		return "", nil, nil, err
	}

	args = append(append([]string{"-v"}, opts...), os.DevNull)
	out, err := exec.Command(cpp, args...).CombinedOutput()
	if err != nil {
		return "", nil, nil, err
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
	return "", nil, nil, fmt.Errorf("failed parsing %s -v output", cpp)
}

func env(key, val string) string {
	if s := os.Getenv(key); s != "" {
		return s
	}

	return val
}

func translationPhase5and6(ctx *context, toks *[]Token) *[]Token {
	// [0], 5.1.1.2, 5
	//
	// Each source character set member and escape sequence in character
	// constants and string literals is converted to the corresponding
	// member of the execution character set; if there is no corresponding
	// member, it is converted to an implementation- defined member other
	// than the null (wide) character.
	for i, tok := range *toks {
		var cpt cppToken
		switch tok.Rune {
		case STRINGLITERAL, LONGSTRINGLITERAL:
			cpt.char = tok.Rune
			cpt.value = tok.Value
			cpt.fileID = tok.fileID
			cpt.pos = tok.pos
			(*toks)[i].Value = dict.sid(stringConst(cpt))
		case CHARCONST, LONGCHARCONST:
			var cpt cppToken
			cpt.char = tok.Rune
			cpt.value = tok.Value
			cpt.fileID = tok.fileID
			cpt.pos = tok.pos
			switch r := charConst(ctx, cpt); {
			case r >= -255 && r < 0:
				(*toks)[i].Value = dict.sid(string([]byte{byte(-r)}))
			case r <= 255:
				(*toks)[i].Value = dict.sid(string([]byte{byte(r)}))
			default:
				switch cpt.char {
				case CHARCONST:
					ctx.err(tok.Position(), "invalid character constant: %s", tok.Value)
				default:
					(*toks)[i].Value = dict.sid(string(r))
				}
			}
		}
	}
	// [0], 5.1.1.2, 6
	//
	// Adjacent string literal tokens are concatenated.
	w := 0
	for i, tok := range *toks {
		switch tok.Rune {
		case STRINGLITERAL, LONGSTRINGLITERAL:
			if i > 0 {
				switch (*toks)[i-1].Rune {
				case STRINGLITERAL, LONGSTRINGLITERAL:
					(*toks)[i-1].Value = dict.sid((*toks)[i-1].String() + tok.String())
					// /*x*/ "a" /*y*/ "b" -> "ab" with sep "/*x*/  /*y*/"
					(*toks)[i-1].Sep = dict.sid((*toks)[i-1].Sep.String() + tok.Sep.String())
					continue
				}
			}
			fallthrough
		default:
			(*toks)[w] = tok
			w++
		}
	}
	*toks = (*toks)[:w]
	return toks
}

// Token is a grammar terminal.
type Token struct {
	Rune   rune     // ';' or IDENTIFIER etc.
	Sep    StringID // If Config3.PreserveWhiteSpace is in effect: All preceding white space, if any, combined, including comments.
	Value  StringID // ";" or "foo" etc.
	fileID int32
	pos    int32
	seq    int32
}

// String implements fmt.Stringer.
func (t Token) String() string { return t.Value.String() }

// Position implements Node.
func (t *Token) Position() (r token.Position) {
	if t.pos != 0 && t.fileID > 0 {
		if file := cache.file(t.fileID); file != nil {
			r = file.PositionFor(token.Pos(t.pos), true)
		}
	}
	return r
}

// Source is a named part of a translation unit. If Value is empty, Name is
// interpreted as a path to file containing the source code.
type Source struct {
	Name  string
	Value string
}

// Parse preprocesses and parses a translation unit.
//
// Search paths listed in includePaths and sysIncludePaths are used to resolve
// #include "foo.h" and #include <foo.h> preprocessing directives respectively.
// A special search path "@" is interpreted as 'the same directory as where the
// file with the #include directive is'.
//
// The sources should typically provide, usually in this particular order:
//
// - predefined macros, eg.
//
//	#define __SIZE_TYPE__ long unsigned int
//
// - built-in declarations, eg.
//
//	int __builtin_printf(char *__format, ...);
//
// - command-line provided directives, eg.
//
//	#define FOO
//	#define BAR 42
//	#undef QUX
//
// - normal C sources, eg.
//
//	int main() {}
//
// All search and file paths should be absolute paths.
//
// If the preprocessed translation unit is empty, the function may return (nil,
// nil).
//
// The parser does only the minimum declarations/identifier resolving necessary
// for correct parsing. Redeclarations are not checked.
//
// Declarators are inserted in the appropriate scopes.
//
// Tagged struct/union specifier definitions and tagged enum specifier
// definitions are inserted in the appropriate scopes.
func Parse(cfg *Config, includePaths, sysIncludePaths []string, sources []Source) (*AST, error) {
	return parse(newContext(cfg), includePaths, sysIncludePaths, sources)
}

func parse(ctx *context, includePaths, sysIncludePaths []string, sources []Source) (*AST, error) {
	ctx.includePaths = includePaths
	ctx.sysIncludePaths = sysIncludePaths
	var in []source
	for _, v := range sources {
		ts, err := cache.get(ctx, v)
		if err != nil {
			return nil, err
		}

		in = append(in, ts)
	}

	p := newParser(ctx, make(chan *[]Token, 5000)) //DONE benchmark tuned
	var sep StringID
	var seq int32
	go func() {
		defer close(p.in)

		toks := tokenPool.Get().(*[]Token)
		*toks = (*toks)[:0]
		cpp := newCPP(ctx)
		for pline := range cpp.translationPhase4(in) {
			line := *pline
			for _, tok := range line {
				switch tok.char {
				case ' ', '\n':
					switch {
					case sep != 0:
						sep = dict.sid(sep.String() + tok.String())
					default:
						sep = tok.value
					}
				default:
					var t Token
					t.Rune = tok.char
					t.Sep = sep
					t.Value = tok.value
					t.fileID = tok.fileID
					t.pos = tok.pos
					seq++
					t.seq = seq
					*toks = append(*toks, t)
					sep = 0
				}
			}
			token4Pool.Put(pline)
			var c rune
			if n := len(*toks); n != 0 {
				c = (*toks)[n-1].Rune
			}
			switch c {
			case STRINGLITERAL, LONGSTRINGLITERAL:
				// nop
			default:
				if len(*toks) != 0 {
					p.in <- translationPhase5and6(ctx, toks)
					toks = tokenPool.Get().(*[]Token)
					*toks = (*toks)[:0]
				}
			}
		}
		if len(*toks) != 0 {
			p.in <- translationPhase5and6(ctx, toks)
		}
	}()

	tu := p.translationUnit()
	if p.errored { // Must drain
		go func() {
			for range p.in {
			}
		}()
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if p.errored && !ctx.cfg.ignoreErrors {
		return nil, fmt.Errorf("%v: syntax error", p.tok.Position())
	}

	if p.scopes != 0 {
		panic("internal error: invalid scope nesting but no error reported") //TODOOK
	}

	return &AST{
		Scope:             p.declScope,
		TrailingSeperator: sep,
		TranslationUnit:   tu,
	}, nil
}
