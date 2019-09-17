// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v3"

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

// Source is a named part of a translation unit. If Value is empty, Name is
// interpreted as a path to file containing the source code.
type Source struct {
	Name       string
	Value      string
	DoNotCache bool // Disable caching of this source
}

// Promote returns the type the operands of a binary operation are promoted to
// or the type and argument passed in a function call is promoted.
func (n *AssignmentExpression) Promote() Type { return n.promote }

type StructInfo struct {
	Size uintptr

	Align int
}

// AST represents a translation unit and its related data.
type AST struct {
	// Alignment and size of every struct/union defined in the translation
	// unit. Valid only after Translate.
	PtrdiffType Type
	Scope       Scope // File scope.
	SizeType    Type
	Structs     map[StructInfo]struct{}
	// TLD contains pruned file scope declarators, ie. either the first one
	// or the first one that has an initializer.
	TLD               map[*Declarator]struct{}
	TrailingSeperator StringID // White space and/or comments preceding EOF.
	TranslationUnit   *TranslationUnit
	WideCharType      Type
	cfg               *Config
}

// Parse preprocesses and parses a translation unit and returns an *AST or
// error, if any.
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
// Declarators (*Declarator) and StructDeclarators (*StructDeclarator) are
// inserted in the appropriate scopes.
//
// Tagged struct/union specifier definitions (*StructOrUnionSpecifier) are
// inserted in the appropriate scopes.
//
// Tagged enum specifier definitions (*EnumSpecifier) and enumeration constants
// (*Enumerator) are inserted in the appropriate scopes.
//
// Labels (*LabeledStatement) are inserted in the appropriate scopes.
func Parse(cfg *Config, includePaths, sysIncludePaths []string, sources []Source) (*AST, error) {
	return parse(newContext(cfg), includePaths, sysIncludePaths, sources)
}

func parse(ctx *context, includePaths, sysIncludePaths []string, sources []Source) (*AST, error) {
	if debugWorkingDir || ctx.cfg.DebugWorkingDir {
		switch wd, err := os.Getwd(); err {
		case nil:
			fmt.Fprintf(os.Stderr, "OS working dir: %s\n", wd)
		default:
			fmt.Fprintf(os.Stderr, "OS working dir: error %s\n", err)
		}
		fmt.Fprintf(os.Stderr, "Config.WorkingDir: %s\n", ctx.cfg.WorkingDir)
	}
	if debugIncludePaths || ctx.cfg.DebugIncludePaths {
		fmt.Fprintf(os.Stderr, "include paths: %v\n", includePaths)
		fmt.Fprintf(os.Stderr, "system include paths: %v\n", sysIncludePaths)
	}
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
		cpp := newCPP(ctx)

		defer func() {
			close(p.in)
			ctx.intMaxWidth = cpp.intMaxWidth()
		}()

		toks := tokenPool.Get().(*[]Token)
		*toks = (*toks)[:0]
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
					t.file = tok.file
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
					p.in <- translationPhase5(ctx, toks)
					toks = tokenPool.Get().(*[]Token)
					*toks = (*toks)[:0]
				}
			}
		}
		if len(*toks) != 0 {
			p.in <- translationPhase5(ctx, toks)
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
		panic(internalErrorf("invalid scope nesting but no error reported"))
	}

	return &AST{
		Scope:             p.declScope,
		TLD:               map[*Declarator]struct{}{},
		TrailingSeperator: sep,
		TranslationUnit:   tu,
		cfg:               ctx.cfg,
	}, nil
}

func translationPhase5(ctx *context, toks *[]Token) *[]Token {
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
			cpt.file = tok.file
			cpt.pos = tok.pos
			(*toks)[i].Value = dict.sid(stringConst(ctx, cpt))
		case CHARCONST, LONGCHARCONST:
			var cpt cppToken
			cpt.char = tok.Rune
			cpt.value = tok.Value
			cpt.file = tok.file
			cpt.pos = tok.pos
			switch r := charConst(ctx, cpt); {
			case r <= 255:
				(*toks)[i].Value = dict.sid(string(r))
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
	return toks
}

// Preprocess preprocesses a translation unit and outputs the result to w.
//
// Please see Parse for the documentation of the other parameters.
func Preprocess(cfg *Config, includePaths, sysIncludePaths []string, sources []Source, w io.Writer) error {
	ctx := newContext(cfg)
	if debugWorkingDir || ctx.cfg.DebugWorkingDir {
		switch wd, err := os.Getwd(); err {
		case nil:
			fmt.Fprintf(os.Stderr, "OS working dir: %s\n", wd)
		default:
			fmt.Fprintf(os.Stderr, "OS working dir: error %s\n", err)
		}
		fmt.Fprintf(os.Stderr, "Config.WorkingDir: %s\n", ctx.cfg.WorkingDir)
	}
	if debugIncludePaths || ctx.cfg.DebugIncludePaths {
		fmt.Fprintf(os.Stderr, "include paths: %v\n", includePaths)
		fmt.Fprintf(os.Stderr, "system include paths: %v\n", sysIncludePaths)
	}
	ctx.includePaths = includePaths
	ctx.sysIncludePaths = sysIncludePaths
	var in []source
	for _, v := range sources {
		ts, err := cache.get(ctx, v)
		if err != nil {
			return err
		}

		in = append(in, ts)
	}

	var sep StringID
	cpp := newCPP(ctx)
	toks := tokenPool.Get().(*[]Token)
	*toks = (*toks)[:0]
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
				t.file = tok.file
				t.pos = tok.pos
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
				for _, v := range *translationPhase5(ctx, toks) {
					if err := wTok(w, v); err != nil {
						return err
					}
				}
				toks = tokenPool.Get().(*[]Token)
				*toks = (*toks)[:0]
			}
		}
	}
	if len(*toks) != 0 {
		for _, v := range *translationPhase5(ctx, toks) {
			if err := wTok(w, v); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}
	return ctx.Err()
}

func wTok(w io.Writer, tok Token) (err error) {
	switch tok.Rune {
	case STRINGLITERAL, LONGSTRINGLITERAL:
		_, err = fmt.Fprintf(w, `%s"%s"`, tok.Sep, cQuotedString(tok.String()))
	case CHARCONST, LONGCHARCONST:
		_, err = fmt.Fprintf(w, `%s'%s'`, tok.Sep, cQuotedString(tok.String()))
	default:
		_, err = fmt.Fprintf(w, "%s%s", tok.Sep, tok)
	}
	return err
}

func cQuotedString(s string) []byte {
	var b []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '\b':
			b = append(b, '\\', 'b')
			continue
		case '\f':
			b = append(b, '\\', 'f')
			continue
		case '\n':
			b = append(b, '\\', 'n')
			continue
		case '\r':
			b = append(b, '\\', 'r')
			continue
		case '\t':
			b = append(b, '\\', 't')
			continue
		case '\\':
			b = append(b, '\\', '\\')
			continue
		case '"':
			b = append(b, '\\', '"')
			continue
		}

		switch {
		case c < ' ' || c >= 0x7f:
			b = append(b, '\\', octal(c>>6), octal(c>>3), octal(c))
		default:
			b = append(b, c)
		}
	}
	return b
}

func octal(b byte) byte { return '0' + b&7 }

var trcSource = Source{"<builtin-trc>", `
extern void *stderr;
int fflush(void *stream);
int fprintf(void *stream, const char *format, ...);
`, false}

// Translate parses and typechecks a translation unit  and returns an *AST or
// error, if any.
//
// Please see Parse for the documentation of the parameters.
func Translate(cfg *Config, includePaths, sysIncludePaths []string, sources []Source) (*AST, error) {
	if cfg.InjectTracingCode {
		for i, v := range sources {
			if filepath.Ext(v.Name) == ".c" {
				sources = append(append(append([]Source(nil), sources[:i]...), trcSource), sources[i:]...)
			}
		}
	}
	return translate(newContext(cfg), includePaths, sysIncludePaths, sources)
}

func translate(ctx *context, includePaths, sysIncludePaths []string, sources []Source) (*AST, error) {
	ast, err := parse(ctx, includePaths, sysIncludePaths, sources)
	if err != nil {
		return nil, err
	}

	if ctx, err = ast.typecheck(); err != nil {
		return nil, err
	}

	ast.PtrdiffType = ptrdiffT(ctx, ast.Scope, Token{})
	ast.SizeType = sizeT(ctx, ast.Scope, Token{})
	ast.WideCharType = wcharT(ctx, ast.Scope, Token{})
	return ast, nil
}

// Typecheck determines types of objects and expressions and verifies types are
// valid in the context they are used.
func (n *AST) Typecheck() error {
	_, err := n.typecheck()
	return err
}

func (n *AST) typecheck() (*context, error) {
	ctx := newContext(n.cfg)
	if err := ctx.cfg.ABI.sanityCheck(ctx, int(ctx.intMaxWidth), n.Scope); err != nil {
		return nil, err
	}

	ctx.intBits = int(ctx.cfg.ABI.Types[Int].Size) * 8
	ctx.ast = n
	n.TranslationUnit.check(ctx)
	n.Structs = ctx.structs
	var a []int
	for k := range n.Scope {
		a = append(a, int(k))
	}
	sort.Ints(a)
	for _, v := range a {
		nm := StringID(v)
		defs := n.Scope[nm]
		var pruned *Declarator
		for _, v := range defs {
			switch x := v.(type) {
			case *Declarator:
				if x.IsExtern() {
					continue
				}

				if pruned == nil {
					pruned = x
					continue
				}

				//TODO check compatible types
				switch {
				case pruned.hasInitalizer:
					if x.hasInitalizer {
						ctx.errNode(x, "multiple initializers for the same symbol")
						continue
					}
				default:
					if x.hasInitalizer {
						pruned = x
					}
				}
			}
		}
		if pruned == nil {
			continue
		}

		n.TLD[pruned] = struct{}{}
	}
	return ctx, ctx.Err()
}

func (n *AlignmentSpecifier) align() int {
	switch n.Case {
	case AlignmentSpecifierAlignasType: // "_Alignas" '(' TypeName ')'
		return n.TypeName.Type().Align()
	case AlignmentSpecifierAlignasExpr: // "_Alignas" '(' ConstantExpression ')'
		return n.ConstantExpression.Operand.Type().Align()
	default:
		panic(internalError())
	}
}

// Closure reports the variables closed over by a nested function (case
// BlockItemFuncDef).
func (n *BlockItem) Closure() map[StringID]struct{} { return n.closure }

// FunctionDefinition returns the nested function (case BlockItemFuncDef).
func (n *BlockItem) FunctionDefinition() *FunctionDefinition { return n.fn }

func (n *Declarator) IsStatic() bool          { return n.td.static() }
func (n *Declarator) isVisible(at int32) bool { return at == 0 || n.DirectDeclarator.ends() < at }

// NameTok returns n's declaring name token.
func (n *Declarator) NameTok() (r Token) {
	if n == nil || n.DirectDeclarator == nil {
		return r
	}

	return n.DirectDeclarator.NameTok()
}

// LexicalScope returns the lexical scope of n.
func (n *Declarator) LexicalScope() Scope { return n.DirectDeclarator.lexicalScope }

// Name returns n's declared name.
func (n *Declarator) Name() StringID {
	if n == nil || n.DirectDeclarator == nil {
		return 0
	}

	return n.DirectDeclarator.Name()
}

// ParamScope returns the scope in which n's function parameters are declared
// if the underlying type of n is a function or nil otherwise. If n is part of
// a function definition the scope is the same as the scope of the function
// body.
func (n *Declarator) ParamScope() Scope {
	if n == nil {
		return nil
	}

	return n.DirectDeclarator.ParamScope()
}

// Type returns the type of n.
func (n *Declarator) Type() Type { return n.typ }

// IsExtern reports whether n was declared with storage class specifier 'extern'.
func (n *Declarator) IsExtern() bool { return n.td.extern() }

func (n *DeclarationSpecifiers) auto() bool        { return n != nil && n.class&fAuto != 0 }
func (n *DeclarationSpecifiers) extern() bool      { return n != nil && n.class&fExtern != 0 }
func (n *DeclarationSpecifiers) register() bool    { return n != nil && n.class&fRegister != 0 }
func (n *DeclarationSpecifiers) static() bool      { return n != nil && n.class&fStatic != 0 }
func (n *DeclarationSpecifiers) threadLocal() bool { return n != nil && n.class&fThreadLocal != 0 }
func (n *DeclarationSpecifiers) typedef() bool     { return n != nil && n.class&fTypedef != 0 }

func (n *DirectAbstractDeclarator) TypeQualifier() Type { return n.typeQualifiers }

func (n *DirectDeclarator) ends() int32 {
	switch n.Case {
	case DirectDeclaratorIdent: // IDENTIFIER
		return n.Token.seq
	case DirectDeclaratorDecl: // '(' Declarator ')'
		return n.Token2.seq
	case DirectDeclaratorArr: // DirectDeclarator '[' TypeQualifierList AssignmentExpression ']'
		return n.Token2.seq
	case DirectDeclaratorStaticArr: // DirectDeclarator '[' "static" TypeQualifierList AssignmentExpression ']'
		return n.Token3.seq
	case DirectDeclaratorArrStatic: // DirectDeclarator '[' TypeQualifierList "static" AssignmentExpression ']'
		return n.Token3.seq
	case DirectDeclaratorStar: // DirectDeclarator '[' TypeQualifierList '*' ']'
		return n.Token3.seq
	case DirectDeclaratorFuncParam: // DirectDeclarator '(' ParameterTypeList ')'
		return n.Token2.seq
	case DirectDeclaratorFuncIdent: // DirectDeclarator '(' IdentifierList ')'
		return n.Token2.seq
	default:
		panic(internalError())
	}
}

func (n *DirectDeclarator) TypeQualifier() Type { return n.typeQualifiers }

// NameTok returns n's declarin name token.
func (n *DirectDeclarator) NameTok() (r Token) {
	for {
		if n == nil {
			return r
		}

		switch n.Case {
		case DirectDeclaratorIdent: // IDENTIFIER
			return n.Token
		case DirectDeclaratorDecl: // '(' Declarator ')'
			return n.Declarator.NameTok()
		default:
			n = n.DirectDeclarator
		}
	}
}

// Name returns n's declared name.
func (n *DirectDeclarator) Name() StringID {
	for {
		if n == nil {
			return 0
		}

		switch n.Case {
		case DirectDeclaratorIdent: // IDENTIFIER
			return n.Token.Value
		case DirectDeclaratorDecl: // '(' Declarator ')'
			return n.Declarator.Name()
		default:
			n = n.DirectDeclarator
		}
	}
}

// ParamScope returns the innermost scope in which function parameters are
// declared for Case DirectDeclaratorFuncParam or DirectDeclaratorFuncIdent or
// nil otherwise.
func (n *DirectDeclarator) ParamScope() Scope {
	if n == nil {
		return nil
	}

	switch n.Case {
	case DirectDeclaratorIdent: // IDENTIFIER
		return nil
	case DirectDeclaratorDecl: // '(' Declarator ')'
		return n.Declarator.ParamScope()
	case DirectDeclaratorArr: // DirectDeclarator '[' TypeQualifierList AssignmentExpression ']'
		return n.DirectDeclarator.ParamScope()
	case DirectDeclaratorStaticArr: // DirectDeclarator '[' "static" TypeQualifierList AssignmentExpression ']'
		return n.DirectDeclarator.ParamScope()
	case DirectDeclaratorArrStatic: // DirectDeclarator '[' TypeQualifierList "static" AssignmentExpression ']'
		return n.DirectDeclarator.ParamScope()
	case DirectDeclaratorStar: // DirectDeclarator '[' TypeQualifierList '*' ']'
		return n.DirectDeclarator.ParamScope()
	case DirectDeclaratorFuncParam: // DirectDeclarator '(' ParameterTypeList ')'
		if s := n.DirectDeclarator.ParamScope(); s != nil {
			return s
		}

		return n.paramScope
	case DirectDeclaratorFuncIdent: // DirectDeclarator '(' IdentifierList ')'
		if s := n.DirectDeclarator.ParamScope(); s != nil {
			return s
		}

		return n.paramScope
	default:
		panic(internalError())
	}
}

func (n *Enumerator) isVisible(at int32) bool { return n.Token.seq < at }

func (n *EnumSpecifier) Type() Type { return n.typ }

// Promote returns the type the operands of the binary operation are promoted to.
func (n *EqualityExpression) Promote() Type { return n.promote }

func (n *InitDeclarator) Value() *InitializerValue { return n.initializer }

// IsConst reports whether n is constant.
func (n *Initializer) IsConst() bool { return n == nil || n.isConst }

// List returns n as a flattened list of all items that are case
// InitializerExpr.
func (n *Initializer) List() []*Initializer { return n.list }

// Type returns the type this initializer initializes.
func (n *Initializer) Type() Type { return n.typ }

// IsConst reports whether n is constant.
func (n *InitializerList) IsConst() bool { return n == nil || n.isConst }

// List returns n as a flattened list of all items that are case
// InitializerExpr.
func (n *InitializerList) List() []*Initializer {
	if n == nil {
		return nil
	}

	return n.list
}

// LexicalScope returns the lexical scope of n.
func (n *JumpStatement) LexicalScope() Scope { return n.lexicalScope }

// LexicalScope returns the lexical scope of n.
func (n *LabeledStatement) LexicalScope() Scope { return n.lexicalScope }

func (n *ParameterDeclaration) Type() Type { return n.typ }

func (n *Pointer) TypeQualifier() Type { return n.typeQualifiers }

// ResolvedIn reports which scope the identifier of cases
// PrimaryExpressionIdent, PrimaryExpressionEnum were resolved in, if any.
func (n *PrimaryExpression) ResolvedIn() Scope { return n.resolvedIn }

// Promote returns the type the operands of the binary operation are promoted to.
func (n *RelationalExpression) Promote() Type { return n.promote }

// Cases returns the cases a switch statement consist of.
func (n *SelectionStatement) Cases() []*LabeledStatement { return n.cases }

// Promote returns the type the shift count operand is promoted to.
func (n *ShiftExpression) Promote() Type { return n.promote }

func (n *StructOrUnionSpecifier) Type() Type { return n.typ }

// Promote returns the type the type the switch expression is promoted to.
func (n *SelectionStatement) Promote() Type { return n.promote }

// Type returns the type of n.
func (n *TypeName) Type() Type { return n.typ }

// // LexicalScope returns the lexical scope of n.
// func (n *AttributeValue) LexicalScope() Scope { return n.lexicalScope }

// // Scope returns n's scope.
// func (n *CompoundStatement) Scope() Scope { return n.scope }

// // LexicalScope returns the lexical scope of n.
// func (n *Designator) LexicalScope() Scope { return n.lexicalScope }

// // LexicalScope returns the lexical scope of n.
// func (n *DirectDeclarator) LexicalScope() Scope { return n.lexicalScope }

// // LexicalScope returns the lexical scope of n.
// func (n *EnumSpecifier) LexicalScope() Scope { return n.lexicalScope }

// // LexicalScope returns the lexical scope of n.
// func (n *IdentifierList) LexicalScope() Scope { return n.lexicalScope }

// // LexicalScope returns the lexical scope of n.
// func (n *PrimaryExpression) LexicalScope() Scope { return n.lexicalScope }

// // LexicalScope returns the lexical scope of n.
// func (n *StructOrUnionSpecifier) LexicalScope() Scope { return n.lexicalScope }

// // ResolvedIn reports which scope the identifier of case
// // TypeSpecifierTypedefName was resolved in, if any.
// func (n *TypeSpecifier) ResolvedIn() Scope { return n.resolvedIn }

// // LexicalScope returns the lexical scope of n.
// func (n *UnaryExpression) LexicalScope() Scope { return n.lexicalScope }
