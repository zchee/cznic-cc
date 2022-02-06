// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"fmt"
	"strings"
)

var keywords = map[string]rune{
	// C99
	"_Bool":      rune(BOOL),
	"_Complex":   rune(COMPLEX),
	"_Imaginary": rune(IMAGINARY),
	"auto":       rune(AUTO),
	"break":      rune(BREAK),
	"case":       rune(CASE),
	"char":       rune(CHAR),
	"const":      rune(CONST),
	"continue":   rune(CONTINUE),
	"default":    rune(DEFAULT),
	"do":         rune(DO),
	"double":     rune(DOUBLE),
	"else":       rune(ELSE),
	"enum":       rune(ENUM),
	"extern":     rune(EXTERN),
	"float":      rune(FLOAT),
	"for":        rune(FOR),
	"goto":       rune(GOTO),
	"if":         rune(IF),
	"inline":     rune(INLINE),
	"int":        rune(INT),
	"long":       rune(LONG),
	"register":   rune(REGISTER),
	"restrict":   rune(RESTRICT),
	"return":     rune(RETURN),
	"short":      rune(SHORT),
	"signed":     rune(SIGNED),
	"sizeof":     rune(SIZEOF),
	"static":     rune(STATIC),
	"struct":     rune(STRUCT),
	"switch":     rune(SWITCH),
	"typedef":    rune(TYPEDEF),
	"union":      rune(UNION),
	"unsigned":   rune(UNSIGNED),
	"void":       rune(VOID),
	"volatile":   rune(VOLATILE),
	"while":      rune(WHILE),

	// C11
	"_Alignas":      rune(ALIGNAS),
	"_Alignof":      rune(ALIGNOF),
	"_Atomic":       rune(ATOMIC),
	"_Generic":      rune(GENERIC),
	"_Noreturn":     rune(NORETURN),
	"_Thread_local": rune(THREADLOCAL),

	// GCC/clang/other extensions.
	"_Decimal64":    rune(DECIMAL64),
	"_Float128":     rune(FLOAT128),
	"_Float128x":    rune(FLOAT128X),
	"_Float16":      rune(FLOAT16),
	"_Float32":      rune(FLOAT32),
	"_Float32x":     rune(FLOAT32X),
	"_Float64":      rune(FLOAT64),
	"_Float64x":     rune(FLOAT64X),
	"_Nonnull":      rune(NONNULL),
	"__alignof":     rune(ALIGNOF),
	"__alignof__":   rune(ALIGNOF),
	"__asm":         rune(ASM),
	"__asm__":       rune(ASM),
	"__attribute":   rune(ATTRIBUTE),
	"__attribute__": rune(ATTRIBUTE),
	"__complex":     rune(COMPLEX),
	"__complex__":   rune(COMPLEX),
	"__const":       rune(CONST),
	"__declspec":    rune(DECLSPEC),
	"__float128":    rune(FLOAT128),
	"__imag":        rune(IMAG),
	"__imag__":      rune(IMAG),
	"__inline":      rune(INLINE),
	"__inline__":    rune(INLINE),
	"__int128":      rune(INT128),
	"__label__":     rune(LABEL),
	"__m128":        rune(M128),
	"__m256d":       rune(M256D),
	"__real":        rune(REAL),
	"__real__":      rune(REAL),
	"__restrict":    rune(RESTRICT),
	"__restrict__":  rune(RESTRICT),
	"__signed":      rune(SIGNED),
	"__signed__":    rune(SIGNED),
	"__thread":      rune(THREADLOCAL),
	"__typeof":      rune(TYPEOF),
	"__uint128_t":   rune(UINT128),
	"__volatile":    rune(VOLATILE),
	"__volatile__":  rune(VOLATILE),
	"asm":           rune(ASM),
	"typeof":        rune(TYPEOF),
}

type parser struct {
	cpp     *cpp
	scope   *Scope
	fnScope *Scope
	prevNL  Token
	toks    []Token

	seq int32
}

func newParser(cfg *Config, sources []Source) (*parser, error) {
	cpp, err := newCPP(cfg, sources, nil)
	if err != nil {
		return nil, err
	}

	cpp.rune()
	return &parser{
		cpp:   cpp,
		scope: &Scope{},
	}, nil
}

func (p *parser) newScope() { p.scope = newScope(p.scope) }

func (p *parser) closeScope() {
	if parent := p.scope.Parent; parent != nil {
		p.scope = parent
		return
	}

	p.rune()
	p.cpp.eh("%v: internal error (%v:)", p.toks[0].Position(), origin(1))
	if isTesting {
		trc("%v: internal error (%v:)", p.toks[0].Position(), origin(1))
	}
}

func (p *parser) isKeyword(s []byte) (r rune, ok bool) {
	r, ok = keywords[string(s)]
	return r, ok
}

func (p *parser) next() {
	for {
		p.cpp.rune()
		t := p.cpp.shift()
		switch t.Ch {
		case ' ':
			// nop
		case '\n':
			p.prevNL = t
		default:
			t.seq = p.seq
			p.seq++
			p.toks = append(p.toks, p.tok(t))
			return
		}
	}
}

func (p *parser) rune() (r rune) {
	if len(p.toks) == 0 {
		p.next()
	}
	return p.toks[0].Ch
}

func (p *parser) rune2() (r rune) {
	p.rune()
	t := p.toks[0]
	p.checkTypeName(&t)
	return t.Ch
}

func (p *parser) rune3(checkTypeName bool) (r rune) {
	if checkTypeName {
		return p.rune2()
	}

	return p.rune()
}

func (p *parser) checkTypeName(t *Token) (r bool) {
	if t.Ch != rune(IDENTIFIER) {
		return false
	}

	// defer func() { trc("%v %v (%v: %v: %v:)", t, r, origin(3), origin(4), origin(5)) }()
	nm := string(t.Src())
	for s := p.scope; s != nil; s = s.Parent {
		for _, v := range s.Nodes[nm] {
			switch x := v.(type) {
			case *Declarator:
				if x.typename && t.seq >= int32(x.visible) {
					// trc("%v: %q x.typename %v, t.seq %v, x.visible %v (scope %p)", x.Position(), nm, x.typename, t.seq, x.visible, s)
					t.Ch = rune(TYPENAME)
					return true
				}

				// trc("%v: %q x.typename %v, t.seq %v, x.visible %v (scope %p)", x.Position(), nm, x.typename, t.seq, x.visible, s)
				return false
			case *Enumerator:
				if t.seq >= int32(x.visible) {
					return false
				}
			}
		}
	}
	return false
}

func (p *parser) tok(t Token) (r Token) {
	r = t
	switch t.Ch {
	case rune(IDENTIFIER):
		if c, ok := p.isKeyword(r.Src()); ok {
			r.Ch = c
			break
		}
	case rune(PPNUMBER):
		switch s := string(r.Src()); {
		case strings.ContainsAny(s, ".+-ijpIJP"):
			r.Ch = rune(FLOATCONST)
		case strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X"):
			r.Ch = rune(INTCONST)
		case strings.ContainsAny(s, "Ee"):
			r.Ch = rune(FLOATCONST)
		default:
			r.Ch = rune(INTCONST)
		}
	}
	return r
}

func (p *parser) shift0() (r Token) {
	r = p.toks[0]
	if r.Ch != '\n' && p.prevNL.Ch != 0 {
		sep := append(p.prevNL.Sep(), p.prevNL.Src()...)
		sep = append(sep, r.Sep()...)
		r.Set(sep, r.Src())
	}
	r.seq = p.seq
	if r.Ch != eof {
		switch {
		case len(p.toks) == 1:
			p.toks = p.toks[:0]
		default:
			p.toks = p.toks[1:]
		}
		p.prevNL.Ch = 0
		p.seq++
		p.rune()
	}
	return r
}

func (p *parser) shift() (r Token) {
	r = p.shift0()
	// if r.Ch != ' ' && r.Ch != '\n' {
	// 	trc("%v (%v: %v: %v: %v:)", &r, origin(2), origin(3), origin(4), origin(5))
	// }
	switch {
	case r.Ch == rune(STRINGLITERAL) && p.rune() == rune(STRINGLITERAL):
		sep := r.Sep()
		src := r.Src()
		for p.rune() == rune(STRINGLITERAL) {
			t := p.shift0()
			sep = append(sep, t.Sep()...)
			src = append(src[:len(src)-1], t.Src()[1:]...)
		}
		r.Set(sep, src)
	case r.Ch == rune(LONGSTRINGLITERAL) && (p.rune() == rune(LONGSTRINGLITERAL) || p.rune() == rune(STRINGLITERAL)):
		sep := r.Sep()
		src := r.Src()
	out:
		for {
			switch p.rune() {
			case rune(LONGSTRINGLITERAL):
				t := p.shift0()
				sep = append(sep, t.Sep()...)
				src = append(src[:len(src)-1], t.Src()[2:]...)
			case rune(STRINGLITERAL):
				t := p.shift0()
				sep = append(sep, t.Sep()...)
				src = append(src[:len(src)-1], t.Src()[1:]...)
			default:
				break out
			}
		}
		r.Set(sep, src)
	}
	return r
}

func (p *parser) shift2() (r Token) {
	r = p.shift()
	p.checkTypeName(&r)
	return r
}

func (p *parser) peek(i int) (r Token) {
	for len(p.toks) <= i {
		p.next()
	}
	return p.toks[i]
}

func (p *parser) peek2(i int) (r Token) {
	r = p.peek(i)
	r.seq = p.seq + int32(i)
	p.checkTypeName(&r)
	return r
}

func (p *parser) must(r rune) Token {
	c := p.rune()
	t := p.shift()
	if c != r {
		p.cpp.eh("%v: unexpected %v, expected %v", t.Position(), runeName(c), runeName(r))
	}
	return t
}

func (p *parser) parse() (ast *AST, err error) {
	var errors errors
	p.cpp.eh = func(msg string, args ...interface{}) {
		s := fmt.Sprintf(msg, args...)
		if isTesting {
			s += fmt.Sprintf(" (%v: %v)", origin(2), origin(3))
			if isTesting && traceFails {
				trc("%s FAIL (%v)", s, origin(1))
				// trc("\n%s", debug.Stack())
			}
		}
		errors = append(errors, s)
	}
	tu := p.translationUnit()
	switch p.rune() {
	case eof:
		t := p.shift()
		t.Set(append(p.prevNL.Sep(), p.prevNL.Src()...), nil)
		return &AST{
			TranslationUnit: tu,
			EOF:             t,
		}, errors.err()
	default:
		panic(todo("", p.toks[0]))
	}
}

// [0], 6.9 External definitions
//
//  translation-unit:
// 	external-declaration
// 	translation-unit external-declaration
func (p *parser) translationUnit() (r *TranslationUnit) {
	var prev *TranslationUnit
	for p.rune() != eof {
		tu := &TranslationUnit{ExternalDeclaration: p.externalDeclaration()}
		if tu.ExternalDeclaration == nil {
			return r
		}

		switch {
		case r == nil:
			r = tu
		default:
			prev.TranslationUnit = tu
		}
		prev = tu
	}
	return r
}

//  external-declaration:
// 	function-definition
// 	declaration
// 	asm-function-definition
//	asm-statement
// 	;
func (p *parser) externalDeclaration() *ExternalDeclaration {
again:
	switch p.rune() {
	case eof:
		return nil
	case rune(ASM):
		return &ExternalDeclaration{Case: ExternalDeclarationAsmStmt, AsmStatement: p.asmStatement()}
	}

	ds, ok := p.declarationSpecifiers()
	if !ok {
		goto again
	}

	var d *Declarator
	if p.rune() != ';' {
		d = p.declarator(nil, ds, true)
	}
	switch p.rune() {
	case
		',',
		';',
		'=',
		rune(ASM),
		rune(ATTRIBUTE):

		return &ExternalDeclaration{Case: ExternalDeclarationDecl, Declaration: p.declaration(ds, d)}
	case '{':
		return &ExternalDeclaration{Case: ExternalDeclarationFuncDef, FunctionDefinition: p.functionDefinition(ds, d)}
	default:
		if d.isFn() {
			return &ExternalDeclaration{Case: ExternalDeclarationFuncDef, FunctionDefinition: p.functionDefinition(ds, d)}
		}

		panic(todo("", p.toks[0]))
	}
}

// [0], 6.9.1 Function definitions
//
//  function-definition:
// 	declaration-specifiers declarator declaration-list_opt compound-statement
func (p *parser) functionDefinition(ds *DeclarationSpecifiers, d *Declarator) (r *FunctionDefinition) {
	defer func() { p.fnScope = nil }()
	return &FunctionDefinition{DeclarationSpecifiers: ds, Declarator: d, DeclarationList: p.declarationListOpt(), CompoundStatement: p.compoundStatement(true, d.fnParams())}
}

//  declaration-list:
// 	declaration
// 	declaration-list declaration
func (p *parser) declarationListOpt() (r *DeclarationList) {
	var prev *DeclarationList
	for {
		switch p.rune() {
		case '{', eof:
			return r
		}

		dl := &DeclarationList{Declaration: p.declaration(nil, nil)}
		switch {
		case r == nil:
			r = dl
		default:
			prev.DeclarationList = dl
		}
		prev = dl
	}
	return r
}

// [0], 6.8.2 Compound statement
//
//  compound-statement:
// 	{ label-declaration_opt block-item-list_opt }
func (p *parser) compoundStatement(isFnScope bool, params *Scope) (r *CompoundStatement) {
	p.newScope()

	defer p.closeScope()

	if isFnScope {
		p.fnScope = p.scope
		if params != nil {
			for nm, v := range params.Nodes {
				for _, n := range v {
					p.scope.declare(nm, n)
				}
			}
		}
	}
	return &CompoundStatement{Token: p.must('{'), LabelDeclarationList: p.labelDeclarationListOpt(), BlockItemList: p.blockItemListOpt(), Token2: p.must('}')}
}

//  label-declaration-list
//	label-declaration
//	label-declaration-list label-declaration
func (p *parser) labelDeclarationListOpt() (r *LabelDeclarationList) {
	var prev *LabelDeclarationList
	for p.rune() == rune(LABEL) {
		ldl := &LabelDeclarationList{LabelDeclaration: p.labelDeclaration()}
		switch {
		case r == nil:
			r = ldl
		default:
			prev.LabelDeclarationList = ldl
		}
		prev = ldl
	}
	return r
}

//  label-declaration
// 	__label__ identifier-list ;
func (p *parser) labelDeclaration() (r *LabelDeclaration) {
	r = &LabelDeclaration{Token: p.shift(), IdentifierList: p.identifierList(), Token2: p.must(';')}
	for l := r.IdentifierList; l != nil; l = l.IdentifierList {
		p.scope.declare(string(l.Token2.Src()), r)
	}
	return r
}

//  block-item-list:
// 	block-item
// 	block-item-list block-item
func (p *parser) blockItemListOpt() (r *BlockItemList) {
	var prev *BlockItemList
	for {
		switch p.rune() {
		case '}', eof:
			return r
		default:
			bil := &BlockItemList{BlockItem: p.blockItem()}
			switch {
			case r == nil:
				r = bil
			default:
				prev.BlockItemList = bil
			}
			prev = bil
		}
	}
}

//  block-item:
// 	declaration
// 	statement
// 	label-declaration
// 	declaration-specifiers declarator compound-statement
func (p *parser) blockItem() *BlockItem {
again:
	switch ch := p.rune2(); {
	case p.isStatement(ch) || p.isExpression(ch) || ch == rune(TYPENAME) && p.peek(1).Ch == ':':
		return &BlockItem{Case: BlockItemStmt, Statement: p.statement(false)}
	case p.isDeclarationSpecifier(ch, true):
		r0 := p.rune()
		ds, ok := p.declarationSpecifiers()
		if !ok {
			goto again
		}

		if r0 == rune(ATTRIBUTE) && p.rune() == ';' {
			return &BlockItem{
				Case: BlockItemStmt,
				Statement: &Statement{
					Case:                StatementExpr,
					ExpressionStatement: &ExpressionStatement{declarationSpecifiers: ds, AttributeSpecifierList: p.attributeSpecifierListOpt(), Token: p.shift()},
				},
			}
		}

		var d *Declarator
		if p.rune() != ';' {
			d = p.declarator(nil, ds, true)
		}
		switch p.rune() {
		case '{':
			return &BlockItem{Case: BlockItemFuncDef, DeclarationSpecifiers: ds, Declarator: d, CompoundStatement: p.compoundStatement(true, d.fnParams())}
		default:
			return &BlockItem{Case: BlockItemDecl, Declaration: p.declaration(ds, d)}
		}
	default:
		panic(todo("", p.toks[0]))
	}
}

// [0], 6.8 Statements and blocks
//
//  statement:
// 	labeled-statement
// 	compound-statement
// 	expression-statement
// 	selection-statement
// 	iteration-statement
// 	jump-statement
//	asm-statement
func (p *parser) statement(newBlock bool) *Statement {
	if newBlock {
		p.newScope()

		defer p.closeScope()
	}
	switch ch := p.rune(); {
	case ch == rune(IDENTIFIER):
		switch p.peek(1).Ch {
		case ':':
			return &Statement{Case: StatementLabeled, LabeledStatement: p.labeledStatement()}
		default:
			return &Statement{Case: StatementExpr, ExpressionStatement: p.expressionStatement()}
		}
	case p.isExpression(ch) || ch == ';' || ch == rune(ATTRIBUTE):
		return &Statement{Case: StatementExpr, ExpressionStatement: p.expressionStatement()}
	}

	switch p.rune2() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case '{':
		return &Statement{Case: StatementCompound, CompoundStatement: p.compoundStatement(false, nil)}
	case rune(ASM):
		return &Statement{Case: StatementAsm, AsmStatement: p.asmStatement()}
	case
		rune(DO),
		rune(FOR),
		rune(WHILE):

		return &Statement{Case: StatementIteration, IterationStatement: p.iterationStatement()}
	case
		rune(BREAK),
		rune(CONTINUE),
		rune(GOTO),
		rune(RETURN):

		return &Statement{Case: StatementJump, JumpStatement: p.jumpStatement()}
	case
		rune(IF),
		rune(SWITCH):

		return &Statement{Case: StatementSelection, SelectionStatement: p.selectionStatement()}
	case
		rune(CASE),
		rune(DEFAULT):

		return &Statement{Case: StatementLabeled, LabeledStatement: p.labeledStatement()}
	default:
		panic(todo("", p.toks[0]))
	}
}

func (p *parser) isStatement(ch rune) bool {
	switch ch {
	case
		'(',
		'*',
		';',
		'{',
		rune(ASM),
		rune(BREAK),
		rune(CASE),
		rune(CONTINUE),
		rune(DEC),
		rune(DEFAULT),
		rune(DO),
		rune(FOR),
		rune(GOTO),
		rune(IDENTIFIER),
		rune(IF),
		rune(INC),
		rune(LONGSTRINGLITERAL),
		rune(RETURN),
		rune(STRINGLITERAL),
		rune(SWITCH),
		rune(WHILE):

		return true
	}
	return false
}

// [0], 6.8.4 Selection statements
//
//  selection-statement:
//  	if ( expression ) statement
//  	if ( expression ) statement else statement
//  	switch ( expression ) statement
func (p *parser) selectionStatement() (r *SelectionStatement) {
	p.newScope()

	defer p.closeScope()

	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(IF):
		r = &SelectionStatement{Case: SelectionStatementIf, Token: p.shift(), Token2: p.must('('), Expression: p.expression(false), Token3: p.must(')'), Statement: p.statement(true)}
		switch p.rune() {
		case rune(ELSE):
			r.Case = SelectionStatementIfElse
			r.Token4 = p.shift()
			r.Statement2 = p.statement(true)
			return r
		default:
			return r
		}
	case rune(SWITCH):
		return &SelectionStatement{Case: SelectionStatementSwitch, Token: p.shift(), Token2: p.must('('), Expression: p.expression(false), Token3: p.must(')'), Statement: p.statement(true)}
	default:
		panic(todo("", p.toks[0]))
	}
}

// [0], 6.8.6 Jump statements
//
//  jump-statement:
// 	goto identifier ;
// 	goto * expression ;
// 	continue ;
// 	break ;
// 	return expression_opt ;
func (p *parser) jumpStatement() *JumpStatement {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(BREAK):
		return &JumpStatement{Case: JumpStatementBreak, Token: p.shift(), Token2: p.must(';')}
	case rune(CONTINUE):
		return &JumpStatement{Case: JumpStatementContinue, Token: p.shift(), Token2: p.must(';')}
	case rune(GOTO):
		switch p.peek(1).Ch {
		case '*':
			return &JumpStatement{Case: JumpStatementGotoExpr, Token: p.shift(), Token2: p.shift(), Expression: p.expression(false), Token3: p.must(';')}
		default:
			return &JumpStatement{Case: JumpStatementGoto, Token: p.shift(), Token2: p.must(rune(IDENTIFIER)), Token3: p.must(';')}
		}
	case rune(RETURN):
		return &JumpStatement{Case: JumpStatementReturn, Token: p.shift(), Expression: p.expression(true), Token2: p.must(';')}
	default:
		panic(todo("", p.toks[0]))
	}
}

// [0], 6.8.1 Labeled statements
//
//  labeled-statement:
// 	identifier : statement
// 	case constant-expression : statement
// 	case constant-expression ... constant-expression : statement
// 	default : statement
func (p *parser) labeledStatement() (r *LabeledStatement) {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(CASE):
		r = &LabeledStatement{Token: p.shift(), ConstantExpression: p.constantExpression()}
		switch p.rune() {
		case rune(DDD):
			r.Case = LabeledStatementRange
			r.Token2 = p.shift()
			r.ConstantExpression2 = p.constantExpression()
			r.Token3 = p.must(':')
			r.Statement = p.statement(false)
			return r
		default:
			r.Case = LabeledStatementCaseLabel
			r.Token2 = p.must(':')
			r.Statement = p.statement(false)
			return r
		}
	case rune(DEFAULT):
		return &LabeledStatement{Case: LabeledStatementDefault, Token: p.shift(), Token2: p.must(':'), Statement: p.statement(false)}
	case rune(IDENTIFIER):
		r = &LabeledStatement{Case: LabeledStatementLabel, Token: p.shift(), Token2: p.must(':'), Statement: p.statement(false)}
		switch {
		case p.fnScope == nil:
			panic(todo("", &r.Token))
		default:
			p.fnScope.declare(string(r.Token.Src()), r)
		}
		return r
	default:
		panic(todo("", p.toks[0]))
	}
}

// [0], 6.8.5 Iteration statements
//
//  iteration-statement:
// 	while ( expression ) statement
// 	do statement while ( expression ) ;
// 	for ( expression_opt ; expression_opt ; expression_opt ) statement
// 	for ( declaration expression_opt ; expression_opt ) statement
func (p *parser) iterationStatement() (r *IterationStatement) {
	p.newScope()

	defer p.closeScope()

	switch p.rune() {
	case rune(WHILE):
		return &IterationStatement{Case: IterationStatementWhile, Token: p.shift(), Token2: p.must('('), Expression: p.expression(false), Token3: p.must(')'), Statement: p.statement(true)}
	case rune(DO):
		return &IterationStatement{Case: IterationStatementDo, Token: p.shift(), Statement: p.statement(true), Token2: p.must(rune(WHILE)), Token3: p.must('('), Expression: p.expression(false), Token4: p.must(')'), Token5: p.must(';')}
	case rune(FOR):
		switch ch := p.peek2(2).Ch; {
		case p.isDeclarationSpecifier(ch, true):
			return &IterationStatement{Case: IterationStatementForDecl, Token: p.shift(), Token2: p.must('('), Declaration: p.declaration(nil, nil), Expression: p.expression(true), Token3: p.must(';'), Expression2: p.expression(true), Token4: p.must(')'), Statement: p.statement(true)}
		}

		return &IterationStatement{Case: IterationStatementFor, Token: p.shift(), Token2: p.must('('), Expression: p.expression(true), Token3: p.must(';'), Expression2: p.expression(true), Token4: p.must(';'), Expression3: p.expression(true), Token5: p.must(')'), Statement: p.statement(true)}
	default:
		panic(todo("", p.toks[0]))
	}
}

// [0], 6.8.3 Expression and null statements
//
//  expression-statement:
// 	expression_opt;
func (p *parser) expressionStatement() *ExpressionStatement {
	switch p.rune() {
	case rune(ATTRIBUTE):
		panic(todo("", p.toks[0]))
	default:
		return &ExpressionStatement{Expression: p.expression(true), Token: p.must(';')}
	}
}

//  asm-statement:
//  	asm ;
func (p *parser) asmStatement() *AsmStatement {
	return &AsmStatement{Asm: p.asm(), Token: p.must(';')}
}

//  asm:
// 	asm asm-qualifier-list_opt ( string-literal asm-arg-list_opt )
func (p *parser) asm() *Asm {
	return &Asm{Token: p.must(rune(ASM)), AsmQualifierList: p.asmQualifierListOpt(), Token2: p.must('('), Token3: p.must(rune(STRINGLITERAL)), AsmArgList: p.asmArgListOpt(), Token4: p.must(')')}
}

//  asm-qualifier-list:
// 	asm-qualifier
// 	asm-qualifier-list asm-qualifier
func (p *parser) asmQualifierListOpt() (r *AsmQualifierList) {
	var prev *AsmQualifierList
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return r
		case rune(GOTO), rune(INLINE), rune(VOLATILE):
			aql := &AsmQualifierList{AsmQualifier: p.asmQualifier()}
			switch {
			case r == nil:
				r = aql
			default:
				prev.AsmQualifierList = aql
			}
			prev = aql
		default:
			return r
		}
	}
}

//  asm-qualifier:
// 	volatile
//  	inline
// 	goto"
func (p *parser) asmQualifier() *AsmQualifier {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(VOLATILE):
		return &AsmQualifier{Case: AsmQualifierVolatile, Token: p.shift()}
	case rune(INLINE):
		return &AsmQualifier{Case: AsmQualifierInline, Token: p.shift()}
	case rune(GOTO):
		return &AsmQualifier{Case: AsmQualifierGoto, Token: p.shift()}
	default:
		panic(todo("", p.toks[0]))
	}
}

//  asm-arg-list:
// 	: asm-expression-list
// 	asm-arg-list : asm-expression-list
func (p *parser) asmArgListOpt() (r *AsmArgList) {
	var prev *AsmArgList
	for p.rune() == ':' {
		var aal *AsmArgList
		switch p.peek(1).Ch {
		case
			')',
			':':

			aal = &AsmArgList{Token: p.shift()}
		default:
			aal = &AsmArgList{Token: p.shift(), AsmExpressionList: p.asmExpressionList()}
		}
		switch {
		case r == nil:
			r = aal
		default:
			prev.AsmArgList = aal
		}
		prev = aal
	}
	return r
}

//  asm-expression-list:
// 	asm-index_opt assignment-expression
// 	asm-expression-list , asm-index_opt assignment-expression
func (p *parser) asmExpressionList() (r *AsmExpressionList) {
	r = &AsmExpressionList{AsmIndex: p.asmIndexOpt(), AssignmentExpression: p.assignmentExpression(true)}
	prev := r
	for p.rune() == ',' {
		ael := &AsmExpressionList{Token: p.shift(), AsmIndex: p.asmIndexOpt(), AssignmentExpression: p.assignmentExpression(true)}
		prev.AsmExpressionList = ael
		prev = ael
	}
	return r
}

//  asm-index:
// 	[ expression ]
func (p *parser) asmIndexOpt() *AsmIndex {
	if p.rune() != '[' {
		return nil
	}

	return &AsmIndex{Token: p.shift(), Expression: p.expression(false), Token2: p.must(']')}
}

// [0], 6.5.17 Comma operator
//
//  expression:
// 	assignment-expression
// 	expression , assignment-expression
func (p *parser) expression(opt bool) (r *Expression) {
	if p.isExpression(p.rune()) {
		r = &Expression{AssignmentExpression: p.assignmentExpression(true)}
		for p.rune() == ',' {
			r = &Expression{Expression: r, Token: p.shift(), AssignmentExpression: p.assignmentExpression(true)}
		}
		return r
	}

	if opt {
		return nil
	}

	panic(todo("", p.toks[0], opt))
}

func (p *parser) isExpression(ch rune) bool {
	switch ch {
	case
		'!',
		'&',
		'(',
		'*',
		'+',
		'-',
		'~',
		rune(ALIGNOF),
		rune(ANDAND),
		rune(CHARCONST),
		rune(DEC),
		rune(FLOATCONST),
		rune(GENERIC),
		rune(IDENTIFIER),
		rune(IMAG),
		rune(INC),
		rune(INTCONST),
		rune(LONGCHARCONST),
		rune(LONGSTRINGLITERAL),
		rune(REAL),
		rune(SIZEOF),
		rune(STRINGLITERAL):

		return true
	}
	return false
}

// [0], 6.7 Declarations
//
//  declaration:
// 	declaration-specifiers init-declarator-list_opt attribute-specifier-list_opt;
func (p *parser) declaration(ds *DeclarationSpecifiers, d *Declarator) (r *Declaration) {
	if ds == nil {
		var ok bool
		if ds, ok = p.declarationSpecifiers(); !ok {
			return nil
		}
	}

	return &Declaration{DeclarationSpecifiers: ds, InitDeclaratorList: p.initDeclaratorListOpt(ds, d), AttributeSpecifierList: p.attributeSpecifierListOpt(), Token: p.must(';')}
}

//  attribute-specifier-list:
// 	attribute-specifier
// 	attribute-specifier-list attribute-specifier
func (p *parser) attributeSpecifierListOpt() (r *AttributeSpecifierList) {
	if p.rune() != rune(ATTRIBUTE) {
		return nil
	}

	r = &AttributeSpecifierList{AttributeSpecifier: p.attributeSpecifier()}
	prev := r
	for p.rune() == rune(ATTRIBUTE) {
		asl := &AttributeSpecifierList{AttributeSpecifier: p.attributeSpecifier()}
		prev.AttributeSpecifierList = asl
		prev = asl
	}
	return r
}

//  attribute-specifier:
// 	__attribute__ (( attribute-value-list_opt ))
func (p *parser) attributeSpecifier() (r *AttributeSpecifier) {
	return &AttributeSpecifier{Token: p.must(rune(ATTRIBUTE)), Token2: p.must('('), Token3: p.must('('), AttributeValueList: p.attributeValueListOpt(), Token4: p.must(')'), Token5: p.must(')')}
}

//  attribute-value-list:
// 	attribute-value
// 	attribute-value-list , attribute-value
func (p *parser) attributeValueListOpt() (r *AttributeValueList) {
	switch p.rune() {
	case eof:
		return r
	case rune(IDENTIFIER):
		r = &AttributeValueList{AttributeValue: p.attributeValue()}
	case ')':
		return nil
	default:
		if _, ok := keywords[string(p.toks[0].Src())]; ok {
			r = &AttributeValueList{AttributeValue: p.attributeValue()}
			break
		}

		panic(todo("", p.toks[0]))
	}
	prev := r
	for p.rune() == ',' {
		avl := &AttributeValueList{Token: p.shift(), AttributeValue: p.attributeValue()}
		prev.AttributeValueList = avl
	}
	return r
}

//  attribute-value:
// 	identifier
// 	identifier ( expression-list_opt )
func (p *parser) attributeValue() (r *AttributeValue) {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(IDENTIFIER):
		switch p.peek(1).Ch {
		case '(':
			return &AttributeValue{Case: AttributeValueExpr, Token: p.shift(), Token2: p.shift(), ArgumentExpressionList: p.argumentExpressionListOpt(false), Token3: p.must(')')}
		default:
			return &AttributeValue{Case: AttributeValueIdent, Token: p.shift()}
		}
	default:
		if _, ok := keywords[string(p.toks[0].Src())]; ok {
			r = &AttributeValue{Case: AttributeValueIdent, Token: p.shift()}
			r.Token.Ch = rune(IDENTIFIER)
			return r
		}

		p.cpp.eh("%v: unexpected %v, expected identifier", p.toks[0].Position(), runeName(p.rune()))
		p.shift()
		return nil
	}
}

//  init-declarator-list:
// 	init-declarator
// 	init-declarator-list , init-declarator
func (p *parser) initDeclaratorListOpt(ds *DeclarationSpecifiers, d *Declarator) (r *InitDeclaratorList) {
	switch {
	case d == nil:
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil
		case ';':
			return nil
		case
			'(',
			'*',
			rune(IDENTIFIER):

			d = p.declarator(nil, ds, true)
		default:
			p.cpp.eh("%v: unexpected %v, expected direct-declarator", p.toks[0].Position(), runeName(p.rune()))
			p.shift()
			return nil
		}
	}
	r = &InitDeclaratorList{InitDeclarator: p.initDeclarator(ds, d)}
	prev := r
	for p.rune() == ',' {
		idl := &InitDeclaratorList{Token: p.shift(), AttributeSpecifierList: p.attributeSpecifierListOpt(), InitDeclarator: p.initDeclarator(ds, nil)}
		prev.InitDeclaratorList = idl
		prev = idl
	}
	return r
}

//  init-declarator:
// 	declarator asm-register-var_opt
// 	declarator asm-register-var_opt = initializer
func (p *parser) initDeclarator(ds *DeclarationSpecifiers, d *Declarator) (r *InitDeclarator) {
	if d == nil {
		d = p.declarator(nil, ds, true)
	}

	r = &InitDeclarator{Case: InitDeclaratorDecl, Declarator: d}
	if p.rune() == rune(ASM) {
		r.Asm = p.asm()
	}
	if p.rune() == rune(ATTRIBUTE) {
		r.AttributeSpecifierList = p.attributeSpecifierListOpt()
	}
	if p.rune() == '=' {
		r.Case = InitDeclaratorInit
		r.Token = p.shift()
		r.Initializer = p.initializer()
	}
	return r
}

// [0], 6.7.8 Initialization
//
//  initializer:
// 	assignment-expression
// 	{ initializer-list }
// 	{ initializer-list , }
func (p *parser) initializer() (r *Initializer) {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case '{':
		r = &Initializer{Token: p.shift(), InitializerList: p.initializerList()}
		if p.rune() == ',' {
			r.Token2 = p.shift()
		}

		r.Token3 = p.must('}')
		return r
	default:
		return &Initializer{Case: InitializerExpr, AssignmentExpression: p.assignmentExpression(true)}
	}
}

//  initializer-list:
// 	designation_opt initializer
// 	initializer-list , designation_opt initializer
func (p *parser) initializerList() (r *InitializerList) {
	switch p.rune2() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case '[', '.':
		r = &InitializerList{Designation: p.designation(), Initializer: p.initializer()}
	case rune(IDENTIFIER):
		switch p.peek(1).Ch {
		case ':':
			r = &InitializerList{Designation: p.designation(), Initializer: p.initializer()}
		default:
			r = &InitializerList{Initializer: p.initializer()}
		}
	case '{':
		r = &InitializerList{Initializer: p.initializer()}
	case '}':
		return nil
	default:
		if p.isExpression(p.rune2()) {
			r = &InitializerList{Initializer: p.initializer()}
			break
		}

		panic(todo("", p.toks[0]))
	}
	prev := r
	for p.rune() == ',' {
		var il *InitializerList
		switch p.peek2(1).Ch {
		case
			'.',
			'[':

			il = &InitializerList{Token: p.shift(), Designation: p.designation(), Initializer: p.initializer()}
		case rune(IDENTIFIER):
			switch p.peek(2).Ch {
			case ':':
				il = &InitializerList{Token: p.shift(), Designation: p.designation(), Initializer: p.initializer()}
			default:
				il = &InitializerList{Token: p.shift(), Initializer: p.initializer()}
			}
		case '}':
			return r
		default:
			il = &InitializerList{Token: p.shift(), Initializer: p.initializer()}
		}
		prev.InitializerList = il
		prev = il
	}
	return r
}

//  designation:
// 	designator-list =
func (p *parser) designation() (r *Designation) {
	p.rune()
	dl, last := p.designatorList()
	r = &Designation{DesignatorList: dl}
	if last != DesignatorField2 {
		r.Token = p.must('=')
	}
	return r
}

//  designator-list:
// 	designator
// 	designator-list designator
func (p *parser) designatorList() (r *DesignatorList, last DesignatorCase) {
	var prev *DesignatorList
	for {
		var dl *DesignatorList
		switch p.rune2() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil, 0
		case '[', '.':
			dl = &DesignatorList{Designator: p.designator()}
		case rune(IDENTIFIER):
			switch p.peek(1).Ch {
			case ':':
				dl = &DesignatorList{Designator: p.designator()}
			default:
				return r, last
			}
		default:
			return r, last
		}
		last = dl.Designator.Case
		switch {
		case r == nil:
			r = dl
		default:
			prev.DesignatorList = dl
		}
		prev = dl
	}
}

//  designator:
// 	[ constant-expression ]
// 	[ constant-expression ... constant-expression]
// 	. identifier
//	identifier :
func (p *parser) designator() (r *Designator) {
	switch p.rune2() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case '[':
		r = &Designator{Case: DesignatorIndex, Token: p.shift(), ConstantExpression: p.constantExpression()}
		switch p.rune() {
		case rune(DDD):
			r.Case = DesignatorIndex2
			r.Token2 = p.shift()
			r.ConstantExpression2 = p.constantExpression()
			r.Token3 = p.must(']')
			return r
		default:
			r.Token2 = p.must(']')
			return r
		}
	case '.':
		return &Designator{Case: DesignatorField, Token: p.shift(), Token2: p.must(rune(IDENTIFIER))}
	case rune(IDENTIFIER):
		switch p.peek(1).Ch {
		case ':':
			return &Designator{Case: DesignatorField2, Token: p.shift(), Token2: p.shift()}
		default:
			panic(todo("", p.toks[0]))
		}
	default:
		panic(todo("", p.toks[0]))
	}
}

// [0], 6.5.16 Assignment operators
//
//  assignment-expression:
// 	conditional-expression
// 	unary-expression assignment-operator assignment-expression
//
//  assignment-operator: one of
// 	= *= /= %= += -= <<= >>= &= ^= |=
func (p *parser) assignmentExpression(checkTypeName bool) (r *AssignmentExpression) {
	lhs, u := p.conditionalExpression(checkTypeName)
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case '=':
		r = &AssignmentExpression{Case: AssignmentExpressionAssign, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression(checkTypeName)}
	case rune(MULASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionMul, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression(checkTypeName)}
	case rune(DIVASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionDiv, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression(checkTypeName)}
	case rune(MODASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionMod, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression(checkTypeName)}
	case rune(ADDASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionAdd, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression(checkTypeName)}
	case rune(SUBASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionSub, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression(checkTypeName)}
	case rune(LSHASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionLsh, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression(checkTypeName)}
	case rune(RSHASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionRsh, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression(checkTypeName)}
	case rune(ANDASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionAnd, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression(checkTypeName)}
	case rune(XORASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionXor, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression(checkTypeName)}
	case rune(ORASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionOr, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression(checkTypeName)}
	default:
		return &AssignmentExpression{Case: AssignmentExpressionCond, ConditionalExpression: lhs}
	}
	if u == nil {
		p.cpp.eh("%v: LHS must be unary-expression", r.Token.Position())
	}
	return r
}

// [0], 6.5.15 Conditional operator
//
//  conditional-expression:
// 	logical-OR-expression
// 	logical-OR-expression ? expression : conditional-expression
func (p *parser) conditionalExpression(checkTypeName bool) (r *ConditionalExpression, u *UnaryExpression) {
	lhs, u := p.logicalOrExpression(checkTypeName)
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil, nil
	case '?':
		r = &ConditionalExpression{Case: ConditionalExpressionCond, LogicalOrExpression: lhs, Token: p.shift(), Expression: p.expression(true), Token2: p.must(':')}
		r.ConditionalExpression, _ = p.conditionalExpression(checkTypeName)
		return r, nil
	default:
		return &ConditionalExpression{Case: ConditionalExpressionLOr, LogicalOrExpression: lhs}, u
	}
}

// [0], 6.5.14 Logical OR operator
//
//  logical-OR-expression:
// 	logical-AND-expression
// 	logical-OR-expression || logical-AND-expression
func (p *parser) logicalOrExpression(checkTypeName bool) (r *LogicalOrExpression, u *UnaryExpression) {
	lhs, u := p.logicalAndExpression(checkTypeName)
	r = &LogicalOrExpression{Case: LogicalOrExpressionLAnd, LogicalAndExpression: lhs}
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil, nil
		case rune(OROR):
			r = &LogicalOrExpression{Case: LogicalOrExpressionLOr, LogicalOrExpression: r, Token: p.shift()}
			r.LogicalAndExpression, _ = p.logicalAndExpression(checkTypeName)
			u = nil
		default:
			return r, u
		}
	}
}

// [0], 6.5.13 Logical AND operator
//
//  logical-AND-expression:
// 	inclusive-OR-expression
// 	logical-AND-expression && inclusive-OR-expression
func (p *parser) logicalAndExpression(checkTypeName bool) (r *LogicalAndExpression, u *UnaryExpression) {
	lhs, u := p.inclusiveOrExpression(checkTypeName)
	r = &LogicalAndExpression{Case: LogicalAndExpressionOr, InclusiveOrExpression: lhs}
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil, nil
		case rune(ANDAND):
			r = &LogicalAndExpression{Case: LogicalAndExpressionLAnd, LogicalAndExpression: r, Token: p.shift()}
			r.InclusiveOrExpression, _ = p.inclusiveOrExpression(checkTypeName)
			u = nil
		default:
			return r, u
		}
	}
}

// [0], 6.5.12 Bitwise inclusive OR operator
//
//  inclusive-OR-expression:
// 	exclusive-OR-expression
// 	inclusive-OR-expression | exclusive-OR-expression
func (p *parser) inclusiveOrExpression(checkTypeName bool) (r *InclusiveOrExpression, u *UnaryExpression) {
	lhs, u := p.exclusiveOrExpression(checkTypeName)
	r = &InclusiveOrExpression{Case: InclusiveOrExpressionXor, ExclusiveOrExpression: lhs}
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil, nil
		case '|':
			r = &InclusiveOrExpression{Case: InclusiveOrExpressionOr, InclusiveOrExpression: r, Token: p.shift()}
			r.ExclusiveOrExpression, _ = p.exclusiveOrExpression(checkTypeName)
			u = nil
		default:
			return r, u
		}
	}
}

// [0], 6.5.11 Bitwise exclusive OR operator
//
//  exclusive-OR-expression:
// 	AND-expression
// 	exclusive-OR-expression ^ AND-expression
func (p *parser) exclusiveOrExpression(checkTypeName bool) (r *ExclusiveOrExpression, u *UnaryExpression) {
	lhs, u := p.andExpression(checkTypeName)
	r = &ExclusiveOrExpression{Case: ExclusiveOrExpressionAnd, AndExpression: lhs}
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil, nil
		case '^':
			r = &ExclusiveOrExpression{Case: ExclusiveOrExpressionXor, ExclusiveOrExpression: r, Token: p.shift()}
			r.AndExpression, _ = p.andExpression(checkTypeName)
			u = nil
		default:
			return r, u
		}
	}
}

// [0], 6.5.10 Bitwise AND operator
//
//  AND-expression:
// 	equality-expression
// 	AND-expression & equality-expression
func (p *parser) andExpression(checkTypeName bool) (r *AndExpression, u *UnaryExpression) {
	lhs, u := p.equalityExpression(checkTypeName)
	r = &AndExpression{Case: AndExpressionEq, EqualityExpression: lhs}
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil, nil
		case '&':
			r = &AndExpression{Case: AndExpressionAnd, AndExpression: r, Token: p.shift()}
			r.EqualityExpression, _ = p.equalityExpression(checkTypeName)
			u = nil
		default:
			return r, u
		}
	}
}

// [0], 6.5.9 Equality operators
//
//  equality-expression:
// 	relational-expression
// 	equality-expression == relational-expression
// 	equality-expression != relational-expression
func (p *parser) equalityExpression(checkTypeName bool) (r *EqualityExpression, u *UnaryExpression) {
	lhs, u := p.relationalExpression(checkTypeName)
	r = &EqualityExpression{Case: EqualityExpressionRel, RelationalExpression: lhs}
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil, nil
		case rune(EQ):
			r = &EqualityExpression{Case: EqualityExpressionEq, EqualityExpression: r, Token: p.shift()}
			r.RelationalExpression, _ = p.relationalExpression(checkTypeName)
			u = nil
		case rune(NEQ):
			r = &EqualityExpression{Case: EqualityExpressionNeq, EqualityExpression: r, Token: p.shift()}
			r.RelationalExpression, _ = p.relationalExpression(checkTypeName)
			u = nil
		default:
			return r, u
		}
	}
}

// [0], 6.5.8 Relational operators
//
//  relational-expression:
// 	shift-expression
// 	relational-expression <  shift-expression
// 	relational-expression >  shift-expression
// 	relational-expression <= shift-expression
// 	relational-expression >= shift-expression
func (p *parser) relationalExpression(checkTypeName bool) (r *RelationalExpression, u *UnaryExpression) {
	lhs, u := p.shiftExpression(checkTypeName)
	r = &RelationalExpression{Case: RelationalExpressionShift, ShiftExpression: lhs}
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil, nil
		case '<':
			r = &RelationalExpression{Case: RelationalExpressionLt, RelationalExpression: r, Token: p.shift()}
			r.ShiftExpression, _ = p.shiftExpression(checkTypeName)
			u = nil
		case '>':
			r = &RelationalExpression{Case: RelationalExpressionGt, RelationalExpression: r, Token: p.shift()}
			r.ShiftExpression, _ = p.shiftExpression(checkTypeName)
			u = nil
		case rune(LEQ):
			r = &RelationalExpression{Case: RelationalExpressionLeq, RelationalExpression: r, Token: p.shift()}
			r.ShiftExpression, _ = p.shiftExpression(checkTypeName)
			u = nil
		case rune(GEQ):
			r = &RelationalExpression{Case: RelationalExpressionGeq, RelationalExpression: r, Token: p.shift()}
			r.ShiftExpression, _ = p.shiftExpression(checkTypeName)
			u = nil
		default:
			return r, u
		}
	}
}

// [0], 6.5.7 Bitwise shift operators
//
//  shift-expression:
// 	additive-expression
// 	shift-expression << additive-expression
// 	shift-expression >> additive-expression
func (p *parser) shiftExpression(checkTypeName bool) (r *ShiftExpression, u *UnaryExpression) {
	lhs, u := p.additiveExpression(checkTypeName)
	r = &ShiftExpression{Case: ShiftExpressionAdd, AdditiveExpression: lhs}
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil, nil
		case rune(LSH):
			r = &ShiftExpression{Case: ShiftExpressionLsh, ShiftExpression: r, Token: p.shift()}
			r.AdditiveExpression, _ = p.additiveExpression(checkTypeName)
			u = nil
		case rune(RSH):
			r = &ShiftExpression{Case: ShiftExpressionRsh, ShiftExpression: r, Token: p.shift()}
			r.AdditiveExpression, _ = p.additiveExpression(checkTypeName)
			u = nil
		default:
			return r, u
		}
	}
}

// [0], 6.5.6 Additive operators
//
//  additive-expression:
// 	multiplicative-expression
// 	additive-expression + multiplicative-expression
// 	additive-expression - multiplicative-expression
func (p *parser) additiveExpression(checkTypeName bool) (r *AdditiveExpression, u *UnaryExpression) {
	lhs, u := p.multiplicativeExpression(checkTypeName)
	r = &AdditiveExpression{Case: AdditiveExpressionMul, MultiplicativeExpression: lhs}
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil, nil
		case '+':
			r = &AdditiveExpression{Case: AdditiveExpressionAdd, AdditiveExpression: r, Token: p.shift()}
			r.MultiplicativeExpression, _ = p.multiplicativeExpression(checkTypeName)
			u = nil
		case '-':
			r = &AdditiveExpression{Case: AdditiveExpressionSub, AdditiveExpression: r, Token: p.shift()}
			r.MultiplicativeExpression, _ = p.multiplicativeExpression(checkTypeName)
			u = nil
		default:
			return r, u
		}
	}
}

// [0], 6.5.5 Multiplicative operators
//
//  multiplicative-expression:
// 	cast-expression
// 	multiplicative-expression * cast-expression
// 	multiplicative-expression / cast-expression
// 	multiplicative-expression % cast-expression
func (p *parser) multiplicativeExpression(checkTypeName bool) (r *MultiplicativeExpression, u *UnaryExpression) {
	lhs, u := p.castExpression(checkTypeName)
	r = &MultiplicativeExpression{Case: MultiplicativeExpressionCast, CastExpression: lhs}
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil, nil
		case '*':
			r = &MultiplicativeExpression{Case: MultiplicativeExpressionMul, MultiplicativeExpression: r, Token: p.shift()}
			r.CastExpression, _ = p.castExpression(checkTypeName)
			u = nil
		case '/':
			r = &MultiplicativeExpression{Case: MultiplicativeExpressionDiv, MultiplicativeExpression: r, Token: p.shift()}
			r.CastExpression, _ = p.castExpression(checkTypeName)
			u = nil
		case '%':
			r = &MultiplicativeExpression{Case: MultiplicativeExpressionMod, MultiplicativeExpression: r, Token: p.shift()}
			r.CastExpression, _ = p.castExpression(checkTypeName)
			u = nil
		default:
			return r, u
		}
	}
}

// [0], 6.5.4 Cast operators
//
//  cast-expression:
// 	unary-expression
// 	( type-name ) cast-expression
func (p *parser) castExpression(checkTypeName bool) (r *CastExpression, u *UnaryExpression) {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil, nil
	case '(':
		switch ch := p.peek2(1).Ch; {
		case p.isSpecifierQualifer(ch, true) || ch == rune(ATTRIBUTE):
			lparen := p.shift()
			tn := p.typeName()
			rparen := p.must(')')
			switch p.rune() {
			case eof:
				p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
				return nil, nil
			case '{':
				u = p.unaryExpression(lparen, tn, rparen, checkTypeName)
				return &CastExpression{Case: CastExpressionUnary, UnaryExpression: u}, u
			default:
				r = &CastExpression{Case: CastExpressionCast, Token: lparen, TypeName: tn, Token2: rparen}
				r.CastExpression, _ = p.castExpression(checkTypeName)
				return r, nil
			}
		case p.isExpression(ch) || ch == '{':
			r = &CastExpression{Case: CastExpressionUnary, UnaryExpression: p.unaryExpression(Token{}, nil, Token{}, checkTypeName)}
			return r, r.UnaryExpression
		}

		panic(todo("", p.peek2(1)))
	default:
		r = &CastExpression{Case: CastExpressionUnary, UnaryExpression: p.unaryExpression(Token{}, nil, Token{}, checkTypeName)}
		return r, r.UnaryExpression
	}
}

// [0], 6.7.6 Type names
//
//  type-name:
// 	specifier-qualifier-list abstract-declarator_opt
func (p *parser) typeName() *TypeName {
	return &TypeName{SpecifierQualifierList: p.specifierQualifierList(), AbstractDeclarator: p.abstractDeclarator(nil, true)}
}

//  abstract-declarator:
// 	pointer
// 	pointer_opt direct-abstract-declarator
func (p *parser) abstractDeclarator(ptr *Pointer, opt bool) *AbstractDeclarator {
	if ptr == nil {
		p.attributeSpecifierListOpt() //TODO not stored yet
		ptr = p.pointer(true)
	}
	switch {
	case ptr == nil:
		switch p.rune() {
		case
			')',
			':':

			if opt {
				return nil
			}

			panic(todo("", p.toks[0]))
		case '[':
			return &AbstractDeclarator{Case: AbstractDeclaratorDecl, DirectAbstractDeclarator: p.directAbstractDeclarator()}
		case '(':
			switch p.peek2(1).Ch {
			case
				'*',
				rune(ATTRIBUTE):

				return &AbstractDeclarator{Case: AbstractDeclaratorPtr, DirectAbstractDeclarator: p.directAbstractDeclarator()}
			default:
				panic(todo("", p.peek2(1)))
			}
		default:
			panic(todo("", p.toks[0], opt))
		}
	default:
		switch p.rune() {
		case
			')',
			':':

			return &AbstractDeclarator{Case: AbstractDeclaratorPtr, Pointer: ptr}
		case '[':
			return &AbstractDeclarator{Case: AbstractDeclaratorDecl, Pointer: ptr, DirectAbstractDeclarator: p.directAbstractDeclarator()}
		case '(':
			switch ch := p.peek2(1).Ch; {
			case ch == '*' || p.isDeclarationSpecifier(ch, true):
				return &AbstractDeclarator{Case: AbstractDeclaratorPtr, Pointer: ptr, DirectAbstractDeclarator: p.directAbstractDeclarator()}
			default:
				panic(todo("", p.peek2(1)))
			}
		default:
			panic(todo("", p.toks[0]))
		}
	}
}

//  direct-abstract-declarator:
// 	( abstract-declarator )
// 	direct-abstract-declarator_opt [ type-qualifier-list_opt assignment-expression_opt ]
// 	direct-abstract-declarator_opt [ static type-qualifier-list_opt assignment-expression ]
// 	direct-abstract-declarator_opt [ type-qualifier-list static assignment-expression ]
// 	direct-abstract-declarator_opt [ * ]
// 	direct-abstract-declarator_opt ( parameter-type-list_opt )
func (p *parser) directAbstractDeclarator() (r *DirectAbstractDeclarator) {
out:
	switch p.rune() {
	case '[':
		switch ch := p.peek(1).Ch; {
		case p.isTypeQualifier(ch):
			lbracket := p.shift()
			tql := p.typeQualifierList(false, false)
			switch p.rune() {
			case eof:
				p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
				return nil
			case rune(INTCONST):
				r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorArr, Token: lbracket, TypeQualifiers: tql, AssignmentExpression: p.assignmentExpression(true), Token2: p.must(']')}
				break out
			case rune(STATIC):
				r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorArr, Token: lbracket, TypeQualifiers: tql, Token2: p.shift(), AssignmentExpression: p.assignmentExpression(true), Token3: p.must(']')}
				break out
			default:
				panic(todo("", &lbracket, tql, &p.toks[0]))
			}
		case ch != '*' && p.isExpression(ch):
			r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorArr, Token: p.shift(), AssignmentExpression: p.assignmentExpression(true), Token2: p.must(']')}
			break out
		}

		switch p.peek(1).Ch {
		case ']':
			r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorArr, Token: p.shift(), Token2: p.shift()}
		case rune(STATIC):
			lbracket := p.shift()
			switch {
			case p.isTypeQualifier(p.peek(1).Ch):
				r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorStaticArr, Token: lbracket, Token2: p.shift(), TypeQualifiers: p.typeQualifierList(false, false), AssignmentExpression: p.assignmentExpression(true), Token3: p.must(']')}
			default:
				panic(todo("", &lbracket, p.peek(1)))
			}
		case '*':
			switch p.peek(2).Ch {
			case ']':
				r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorArrStar, Token: p.shift(), Token2: p.shift(), Token3: p.shift()}
			default:
				panic(todo("", p.peek(2)))
			}
		default:
			panic(todo("", p.peek(1)))
		}
	case '(':
		switch ch := p.peek2(1).Ch; {
		case p.in(ch, '*', rune(ATTRIBUTE)):
			r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorDecl, Token: p.shift(), AbstractDeclarator: p.abstractDeclarator(nil, false), Token2: p.must(')')}
		case p.isDeclarationSpecifier(ch, true):
			r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorFunc, Token: p.shift(), ParameterTypeList: p.parameterTypeListOpt(), Token2: p.must(')')}
		default:
			panic(todo("", p.peek2(1)))
		}
	default:
		panic(todo("", p.toks[0]))
	}
	return p.directAbstractDeclarator2(r)
}

func (p *parser) directAbstractDeclarator2(dad *DirectAbstractDeclarator) (r *DirectAbstractDeclarator) {
	r = dad
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil
		case
			')',
			',',
			':':

			return r
		case '(':
			r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorFunc, DirectAbstractDeclarator: r, Token: p.shift(), ParameterTypeList: p.parameterTypeListOpt(), Token2: p.must(')')}
		case '[':
			ch := p.peek2(1).Ch
			switch {
			case p.isExpression(ch):
				r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorArr, DirectAbstractDeclarator: r, Token: p.shift(), AssignmentExpression: p.assignmentExpression(true), Token2: p.must(']')}
				continue
			default:
				panic(todo("", p.peek2(1)))
			}
		default:
			panic(todo("", &p.toks[0]))
		}
	}
}

//  specifier-qualifier-list:
// 	type-specifier specifier-qualifier-list_opt
// 	type-qualifier specifier-qualifier-list_opt
// 	alignment-specifier specifier-qualifier-list_opt
func (p *parser) specifierQualifierList() (r *SpecifierQualifierList) {
	var prev *SpecifierQualifierList
	acceptTypeName := true
	for {
		var sql *SpecifierQualifierList
		switch ch := p.rune(); {
		case ch == rune(IDENTIFIER):
			if !acceptTypeName || !p.checkTypeName(&p.toks[0]) {
				return r
			}

			sql = &SpecifierQualifierList{Case: SpecifierQualifierListTypeSpec, TypeSpecifier: p.typeSpecifier()}
			acceptTypeName = false
		case p.isTypeSpecifier(ch, false):
			sql = &SpecifierQualifierList{Case: SpecifierQualifierListTypeSpec, TypeSpecifier: p.typeSpecifier()}
			acceptTypeName = false
		case p.isTypeQualifier(ch) || ch == rune(ATTRIBUTE):
			sql = &SpecifierQualifierList{Case: SpecifierQualifierListTypeQual, TypeQualifier: p.typeQualifier(true)}
		case ch == rune(ALIGNAS):
			sql = &SpecifierQualifierList{Case: SpecifierQualifierListAlignSpec, AlignmentSpecifier: p.alignmentSpecifier()}
		default:
			return r
		}

		switch {
		case r == nil:
			r = sql
		default:
			prev.SpecifierQualifierList = sql
		}
		prev = sql
	}
}

func (p *parser) isSpecifierQualifer(ch rune, typenameOk bool) bool {
	return p.isTypeSpecifier(ch, typenameOk) || p.isTypeQualifier(ch)
}

// [0], 6.5.3 Unary operators
//
//  unary-expression:
// 	postfix-expression
// 	++ unary-expression
// 	-- unary-expression
// 	unary-operator cast-expression
// 	sizeof unary-expression
// 	sizeof ( type-name )
// 	&& identifier
// 	_Alignof unary-expression
// 	_Alignof ( type-name )
// 	__imag__ unary-expression
// 	__real__ unary-expression
//
//  unary-operator: one of
// 	& * + - ~ !
func (p *parser) unaryExpression(lp Token, tn *TypeName, rp Token, checkTypeName bool) (r *UnaryExpression) {
	if tn != nil {
		return &UnaryExpression{Case: UnaryExpressionPostfix, PostfixExpression: p.postfixExpression(lp, tn, rp, checkTypeName)}
	}

	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case '&':
		r = &UnaryExpression{Case: UnaryExpressionAddrof, Token: p.shift()}
		r.CastExpression, _ = p.castExpression(checkTypeName)
		return r
	case '*':
		r = &UnaryExpression{Case: UnaryExpressionDeref, Token: p.shift()}
		r.CastExpression, _ = p.castExpression(checkTypeName)
		return r
	case '+':
		r = &UnaryExpression{Case: UnaryExpressionPlus, Token: p.shift()}
		r.CastExpression, _ = p.castExpression(checkTypeName)
		return r
	case '-':
		r = &UnaryExpression{Case: UnaryExpressionMinus, Token: p.shift()}
		r.CastExpression, _ = p.castExpression(checkTypeName)
		return r
	case '~':
		r = &UnaryExpression{Case: UnaryExpressionCpl, Token: p.shift()}
		r.CastExpression, _ = p.castExpression(checkTypeName)
		return r
	case '!':
		r = &UnaryExpression{Case: UnaryExpressionNot, Token: p.shift()}
		r.CastExpression, _ = p.castExpression(checkTypeName)
		return r
	case rune(SIZEOF):
		switch p.peek(1).Ch {
		case '(':
			switch {
			case p.isSpecifierQualifer(p.peek2(2).Ch, true):
				sz := p.shift()
				lp := p.shift()
				tn := p.typeName()
				rp := p.must(')')
				switch p.rune() {
				case '{':
					return &UnaryExpression{Case: UnaryExpressionSizeofExpr, Token: sz, UnaryExpression: p.unaryExpression(lp, tn, rp, checkTypeName)}
				default:
					return &UnaryExpression{Case: UnaryExpressionSizeofType, Token: sz, Token2: lp, TypeName: tn, Token3: rp}
				}
			default:
				return &UnaryExpression{Case: UnaryExpressionSizeofExpr, Token: p.shift(), UnaryExpression: p.unaryExpression(Token{}, nil, Token{}, checkTypeName)}
			}
		default:
			return &UnaryExpression{Case: UnaryExpressionSizeofExpr, Token: p.shift(), UnaryExpression: p.unaryExpression(Token{}, nil, Token{}, checkTypeName)}
		}
	case
		'(',
		rune(CHARCONST),
		rune(FLOATCONST),
		rune(GENERIC),
		rune(IDENTIFIER),
		rune(INTCONST),
		rune(LONGCHARCONST),
		rune(LONGSTRINGLITERAL),
		rune(STRINGLITERAL):

		return &UnaryExpression{Case: UnaryExpressionPostfix, PostfixExpression: p.postfixExpression(Token{}, nil, Token{}, checkTypeName)}
	case rune(INC):
		return &UnaryExpression{Case: UnaryExpressionInc, Token: p.shift(), UnaryExpression: p.unaryExpression(Token{}, nil, Token{}, checkTypeName)}
	case rune(DEC):
		return &UnaryExpression{Case: UnaryExpressionDec, Token: p.shift(), UnaryExpression: p.unaryExpression(Token{}, nil, Token{}, checkTypeName)}
	case rune(ALIGNOF):
		switch p.peek(1).Ch {
		case '(':
			switch {
			case p.isSpecifierQualifer(p.peek2(2).Ch, true):
				return &UnaryExpression{Case: UnaryExpressionAlignofType, Token: p.shift(), Token2: p.shift(), TypeName: p.typeName(), Token3: p.must(')')}
			default:
				return &UnaryExpression{Case: UnaryExpressionAlignofExpr, Token: p.shift(), UnaryExpression: p.unaryExpression(Token{}, nil, Token{}, checkTypeName)}
			}
		default:
			panic(todo("", p.peek(1)))
		}
	case rune(ANDAND):
		return &UnaryExpression{Case: UnaryExpressionLabelAddr, Token: p.shift(), Token2: p.must(rune(IDENTIFIER))}
	case rune(REAL):
		return &UnaryExpression{Case: UnaryExpressionReal, Token: p.shift(), UnaryExpression: p.unaryExpression(Token{}, nil, Token{}, checkTypeName)}
	case rune(IMAG):
		return &UnaryExpression{Case: UnaryExpressionImag, Token: p.shift(), UnaryExpression: p.unaryExpression(Token{}, nil, Token{}, checkTypeName)}
	default:
		panic(todo("", p.toks[0]))
	}
}

// [0], 6.5.2 Postfix operators
//
//  postfix-expression:
// 	primary-expression
// 	postfix-expression [ expression ]
// 	postfix-expression ( argument-expression-list_opt )
// 	postfix-expression . identifier
// 	postfix-expression -> identifier
// 	postfix-expression ++
// 	postfix-expression --
// 	( type-name ) { initializer-list }
// 	( type-name ) { initializer-list , }
// 	__builtin_types_compatible_p ( type-name , type-name )
func (p *parser) postfixExpression(lp Token, tn *TypeName, rp Token, checkTypeName bool) (r *PostfixExpression) {
	switch {
	case tn != nil:
		r = &PostfixExpression{Case: PostfixExpressionComplit, Token: lp, TypeName: tn, Token2: rp, Token3: p.must('{'), InitializerList: p.initializerList()}
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil
		case ',':
			r.Token4 = p.shift()
			fallthrough
		default:
			r.Token5 = p.must('}')
		}
	default:
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil
		case '(':
			switch ch := p.peek2(1).Ch; {
			case p.isExpression(ch) || ch == '{':
				r = &PostfixExpression{Case: PostfixExpressionPrimary, PrimaryExpression: p.primaryExpression(checkTypeName)}
			case p.isSpecifierQualifer(ch, true):
				r = &PostfixExpression{Case: PostfixExpressionComplit, Token: p.shift(), TypeName: p.typeName(), Token2: p.must(')'), Token3: p.must('{'), InitializerList: p.initializerList()}
				switch p.rune() {
				case eof:
					p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
					return nil
				case ',':
					r.Token4 = p.shift()
					fallthrough
				default:
					r.Token5 = p.must('}')
					return r
				}
			default:
				panic(todo("", p.peek2(1)))
			}
		default:
			r = &PostfixExpression{Case: PostfixExpressionPrimary, PrimaryExpression: p.primaryExpression(checkTypeName)}
		}
	}
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil
		case '[':
			r = &PostfixExpression{Case: PostfixExpressionIndex, PostfixExpression: r, Token: p.shift(), Expression: p.expression(false), Token2: p.must(']')}
		case '(':
			switch ch := p.peek2(1).Ch; {
			case p.isExpression(ch) || ch == ')':
				r = &PostfixExpression{Case: PostfixExpressionCall, PostfixExpression: r, Token: p.shift(), ArgumentExpressionList: p.argumentExpressionListOpt(true), Token2: p.must(')')}
			default:
				panic(todo("", p.peek2(1)))
			}
		case '.':
			r = &PostfixExpression{Case: PostfixExpressionSelect, PostfixExpression: r, Token: p.shift(), Token2: p.must(rune(IDENTIFIER))}
		case rune(ARROW):
			r = &PostfixExpression{Case: PostfixExpressionPSelect, PostfixExpression: r, Token: p.shift(), Token2: p.must(rune(IDENTIFIER))}
		case rune(INC):
			r = &PostfixExpression{Case: PostfixExpressionInc, PostfixExpression: r, Token: p.shift()}
		case rune(DEC):
			r = &PostfixExpression{Case: PostfixExpressionDec, PostfixExpression: r, Token: p.shift()}
		default:
			return r
		}
	}
}

//  argument-expression-list:
// 	assignment-expression
// 	argument-expression-list , assignment-expression
func (p *parser) argumentExpressionListOpt(checkTypeName bool) (r *ArgumentExpressionList) {
	p.rune()
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case ')':
		return nil
	}

	r = &ArgumentExpressionList{AssignmentExpression: p.assignmentExpression(checkTypeName)}
	prev := r
	for p.rune() == ',' {
		prev.ArgumentExpressionList = &ArgumentExpressionList{Token: p.shift(), AssignmentExpression: p.assignmentExpression(checkTypeName)}
		p.rune()
	}
	p.rune()
	return r
}

// [0], 6.5.1 Primary expressions
//
//  primary-expression:
// 	identifier
// 	constant
// 	string-literal
// 	( expression )
// 	( compound-statement )
//	generic-selection
func (p *parser) primaryExpression(checkTypeName bool) (r *PrimaryExpression) {
	switch p.rune3(checkTypeName) {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(CHARCONST):
		return &PrimaryExpression{Case: PrimaryExpressionChar, Token: p.shift()}
	case rune(FLOATCONST):
		return &PrimaryExpression{Case: PrimaryExpressionFloat, Token: p.shift()}
	case rune(IDENTIFIER):
		return &PrimaryExpression{Case: PrimaryExpressionIdent, Token: p.shift()}
	case rune(INTCONST):
		return &PrimaryExpression{Case: PrimaryExpressionInt, Token: p.shift()}
	case rune(LONGCHARCONST):
		return &PrimaryExpression{Case: PrimaryExpressionLChar, Token: p.shift()}
	case rune(LONGSTRINGLITERAL):
		return &PrimaryExpression{Case: PrimaryExpressionLString, Token: p.shift()}
	case rune(STRINGLITERAL):
		return &PrimaryExpression{Case: PrimaryExpressionString, Token: p.shift()}
	case '(':
		switch p.peek(1).Ch {
		case '{':
			return &PrimaryExpression{Case: PrimaryExpressionStmt, Token: p.shift(), CompoundStatement: p.compoundStatement(false, nil), Token2: p.must(')')}
		default:
			return &PrimaryExpression{Case: PrimaryExpressionExpr, Token: p.shift(), Expression: p.expression(false), Token2: p.must(')')}
		}
	case rune(GENERIC):
		return &PrimaryExpression{Case: PrimaryExpressionGeneric, GenericSelection: p.genericSelection()}
	default:
		panic(todo("%v %v scope %p", p.toks[0], checkTypeName, p.scope))
	}
}

// generic-selection:
//	_Generic ( assignment-expression , generic-assoc-list )
func (p *parser) genericSelection() (r *GenericSelection) {
	return &GenericSelection{Token: p.must(rune(GENERIC)), Token2: p.must('('), AssignmentExpression: p.assignmentExpression(true), Token3: p.must(','), GenericAssociationList: p.genericAssociationList(), Token4: p.must(')')}
}

// generic-assoc-list:
// 	generic-association
// 	generic-assoc-list , generic-association
func (p *parser) genericAssociationList() (r *GenericAssociationList) {
	r = &GenericAssociationList{GenericAssociation: p.genericAssociation()}
	prev := r
	for p.rune() == ',' {
		gal := &GenericAssociationList{Token: p.shift(), GenericAssociation: p.genericAssociation()}
		prev.GenericAssociationList = gal
		prev = gal
	}
	return r
}

// generic-association:
//	type-name : assignment-expression
//	default : assignment-expression
func (p *parser) genericAssociation() (r *GenericAssociation) {
	switch {
	case p.rune() == rune(DEFAULT):
		return &GenericAssociation{Case: GenericAssociationDefault, Token: p.shift(), Token2: p.must(':'), AssignmentExpression: p.assignmentExpression(true)}
	default:
		return &GenericAssociation{Case: GenericAssociationType, TypeName: p.typeName(), Token: p.must(':'), AssignmentExpression: p.assignmentExpression(true)}
	}
}

// [0], 6.7.5 Declarators
//
//  declarator:
// 	pointer_opt direct-declarator
func (p *parser) declarator(ptr *Pointer, ds *DeclarationSpecifiers, declare bool) (r *Declarator) {
	if ptr == nil {
		ptr = p.pointer(true)
	}
	r = &Declarator{Pointer: ptr, DirectDeclarator: p.directDeclarator(declare)}
	if ds != nil {
		r.typename = ds.typedef
	}
	if declare {
		r.visible = visible(p.seq) // [0]6.2.1,7
		p.scope.declare(r.Name(), r)
		// if r.typename {
		// 	trc("%v: %s is a typename", r.Position(), r.Name())
		// }
	}
	return r
}

//  direct-declarator:
// 	identifier
// 	( declarator )
// 	direct-declarator [ type-qualifier-list_opt assignment-expression_opt ]
// 	direct-declarator [ static type-qualifier-list_opt assignment-expression ]
// 	direct-declarator [ type-qualifier-list static assignment-expression ]
// 	direct-declarator [ type-qualifier-list_opt * ]
// 	direct-declarator ( parameter-type-list_opt )
// 	direct-declarator ( identifier-list_opt )
func (p *parser) directDeclarator(declare bool) (r *DirectDeclarator) {
	switch p.rune() {
	case rune(IDENTIFIER):
		r = &DirectDeclarator{Case: DirectDeclaratorIdent, Token: p.shift()}
	case '(':
		r = &DirectDeclarator{Case: DirectDeclaratorDecl, Token: p.shift()}
		p.attributeSpecifierListOpt() //TODO not stored yet
		r.Declarator = p.declarator(nil, nil, false)
		r.Token2 = p.must(')')
	default:
		panic(todo("", p.toks[0]))
	}
	return p.directDeclarator2(r, declare)
}

func (p *parser) directDeclarator2(dd *DirectDeclarator, declare bool) (r *DirectDeclarator) {
	r = dd
	for {
		switch p.rune() {
		case '(':
			p.newScope()

			func() {
				defer p.closeScope()

				switch p.peek2(1).Ch {
				case rune(IDENTIFIER):
					r = &DirectDeclarator{Case: DirectDeclaratorFuncIdent, DirectDeclarator: r, Token: p.shift(), IdentifierList: p.identifierList(), Token2: p.must(')')}
				default:
					r = &DirectDeclarator{Case: DirectDeclaratorFuncParam, DirectDeclarator: r, Token: p.shift(), ParameterTypeList: p.parameterTypeListOpt(), Token2: p.must(')')}
				}
				r.params = p.scope
			}()
		case '[':
			ch := p.peek2(1).Ch
			switch {
			case p.isExpression(ch):
				r = &DirectDeclarator{Case: DirectDeclaratorArr, DirectDeclarator: r, Token: p.shift(), AssignmentExpression: p.assignmentExpression(true), Token2: p.must(']')}
				continue
			case p.isTypeQualifier(ch):
				lbracket := p.shift()
				tql := p.typeQualifierList(false, false)
				switch p.rune() {
				case rune(STATIC):
					r = &DirectDeclarator{Case: DirectDeclaratorArrStatic, DirectDeclarator: r, Token: lbracket, TypeQualifiers: tql, Token2: p.shift(), AssignmentExpression: p.assignmentExpression(true), Token3: p.must(']')}
				case '*':
					r = &DirectDeclarator{Case: DirectDeclaratorStar, DirectDeclarator: r, Token: lbracket, TypeQualifiers: tql, Token2: p.shift(), Token3: p.must(']')}
				default:
					r = &DirectDeclarator{Case: DirectDeclaratorArr, DirectDeclarator: r, Token: lbracket, TypeQualifiers: tql, AssignmentExpression: p.assignmentExpression(true), Token3: p.must(']')}
				}
				continue
			}

			switch ch {
			case ']':
				r = &DirectDeclarator{Case: DirectDeclaratorArr, DirectDeclarator: r, Token: p.shift(), Token2: p.shift()}
			case rune(STATIC):
				r = &DirectDeclarator{Case: DirectDeclaratorStaticArr, DirectDeclarator: r, Token: p.shift(), Token2: p.shift(), TypeQualifiers: p.typeQualifierList(true, false), AssignmentExpression: p.assignmentExpression(true), Token3: p.must(']')}
			default:
				panic(todo("", p.toks[0], p.peek2(1)))
			}
		default:
			return r
		}
	}
	return r
}

//  identifier-list:
// 	identifier
// 	identifier-list , identifier
func (p *parser) identifierList() (r *IdentifierList) {
	switch p.rune2() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(IDENTIFIER):
		r = &IdentifierList{Token2: p.shift()}
	default:
		panic(todo("", p.toks[0]))
	}

	prev := r
	for p.rune() == ',' {
		il := &IdentifierList{Token: p.shift(), Token2: p.must(rune(IDENTIFIER))}
		prev.IdentifierList = il
		prev = il
	}
	return r
}

//  parameter-type-list:
// 	parameter-list
// 	parameter-list , ...
func (p *parser) parameterTypeListOpt() (r *ParameterTypeList) {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case ')':
		return nil
	}

	pl := p.parameterList()
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case ',':
		return &ParameterTypeList{Case: ParameterTypeListVar, ParameterList: pl, Token: p.shift(), Token2: p.must(rune(DDD))}
	default:
		return &ParameterTypeList{Case: ParameterTypeListList, ParameterList: pl}
	}
}

//  parameter-list:
// 	parameter-declaration
// 	parameter-list , parameter-declaration
func (p *parser) parameterList() (r *ParameterList) {
	r = &ParameterList{ParameterDeclaration: p.parameterDeclaration()}
	prev := r
	for p.rune() == ',' {
		if p.peek(1).Ch == rune(DDD) {
			return r
		}

		pl := &ParameterList{Token: p.shift(), ParameterDeclaration: p.parameterDeclaration()}
		prev.ParameterList = pl
		prev = pl
	}
	return r
}

//  parameter-declaration:
// 	declaration-specifiers declarator
// 	declaration-specifiers abstract-declarator_opt
func (p *parser) parameterDeclaration() (r *ParameterDeclaration) {
	ds, ok := p.declarationSpecifiers()
	if !ok {
		return nil
	}

	switch p.rune() {
	case
		')',
		',':

		return &ParameterDeclaration{Case: ParameterDeclarationAbstract, DeclarationSpecifiers: ds}
	default:
		switch x := p.declaratorOrAbstractDeclarator(true).(type) {
		case *AbstractDeclarator:
			return &ParameterDeclaration{Case: ParameterDeclarationAbstract, DeclarationSpecifiers: ds, AbstractDeclarator: x}
		case *Declarator:
			return &ParameterDeclaration{Case: ParameterDeclarationDecl, DeclarationSpecifiers: ds, Declarator: x, AttributeSpecifierList: p.attributeSpecifierListOpt()}
		default:
			p.rune()
			panic(todo("%v %T", p.toks[0], x))
		}
	}
}

func (p *parser) declaratorOrAbstractDeclarator(declare bool) (r Node) {
	// state 2
	ptr0 := p.pointer(true)
	switch p.rune() {
	case
		')',
		',':

		if ptr0 != nil {
			return &AbstractDeclarator{Case: AbstractDeclaratorPtr, Pointer: ptr0}
		}

		panic(todo("", p.toks[0]))
	case '(':
		lparen := p.shift()
		p.attributeSpecifierListOpt() //TODO not stored yet
		if p.isSpecifierQualifer(p.rune2(), true) {
			dad := &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorFunc, Token: lparen, ParameterTypeList: p.parameterTypeListOpt(), Token2: p.must(')')}
			switch {
			case p.in(p.rune(), ',', ')'):
				return &AbstractDeclarator{Case: AbstractDeclaratorDecl, Pointer: ptr0, DirectAbstractDeclarator: dad}
			default:
				panic(todo("", p.toks[0]))
			}
		}

		switch p.rune2() {
		case ')':
			dad := &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorFunc, Token: lparen, Token2: p.shift()}
			return &AbstractDeclarator{Case: AbstractDeclaratorDecl, Pointer: ptr0, DirectAbstractDeclarator: p.directAbstractDeclarator2(dad)}
		case '(':
			switch x := p.declaratorOrAbstractDeclarator(declare).(type) {
			case *AbstractDeclarator:
				dad := &DirectAbstractDeclarator{
					Case:               DirectAbstractDeclaratorDecl,
					Token:              lparen,
					AbstractDeclarator: x,
					Token2:             p.must(')'),
				}
				return &AbstractDeclarator{Case: AbstractDeclaratorDecl, Pointer: ptr0, DirectAbstractDeclarator: p.directAbstractDeclarator2(dad)}
			case *Declarator:
				dd := &DirectDeclarator{
					Case:       DirectDeclaratorDecl,
					Token:      lparen,
					Declarator: x,
					Token2:     p.must(')'),
				}
				return &Declarator{Pointer: ptr0, DirectDeclarator: p.directDeclarator2(dd, declare)}
			default:
				p.rune()
				panic(todo("%v %T", p.toks[0], x))
			}
		case rune(IDENTIFIER):
			dd := p.directDeclarator(declare)
			switch p.rune() {
			case ')':
				dd = &DirectDeclarator{
					Case:  DirectDeclaratorDecl,
					Token: lparen,
					Declarator: &Declarator{
						DirectDeclarator: dd,
					},
					Token2: p.must(')'),
				}
				return &Declarator{Pointer: ptr0, DirectDeclarator: p.directDeclarator2(dd, declare)}
			default:
				panic(todo("", p.toks[0]))
			}
		case '*':
			ptr := p.pointer(false)
			switch p.rune() {
			case
				')',
				'[':

				dad := &DirectAbstractDeclarator{
					Case:               DirectAbstractDeclaratorDecl,
					Token:              lparen,
					AbstractDeclarator: p.abstractDeclarator(ptr, false),
					Token2:             p.must(')'),
				}
				return &AbstractDeclarator{Case: AbstractDeclaratorDecl, Pointer: ptr0, DirectAbstractDeclarator: p.directAbstractDeclarator2(dad)}
			case rune(IDENTIFIER):
				dd := p.directDeclarator(declare)
				switch p.rune() {
				case ')':
					dd = &DirectDeclarator{
						Case:  DirectDeclaratorDecl,
						Token: lparen,
						Declarator: &Declarator{
							Pointer:          ptr,
							DirectDeclarator: dd,
						},
						Token2: p.must(')'),
					}
					return &Declarator{Pointer: ptr0, DirectDeclarator: p.directDeclarator2(dd, declare)}
				default:
					panic(todo("", p.toks[0]))
				}
			case '(':
				switch x := p.declaratorOrAbstractDeclarator(declare).(type) {
				case *Declarator:
					x.Pointer = ptr
					dd := &DirectDeclarator{Case: DirectDeclaratorDecl, Token: lparen, Declarator: x, Token2: p.must(')')}
					return &Declarator{Pointer: ptr0, DirectDeclarator: dd}
				default:
					p.rune()
					panic(todo("%v %T", p.toks[0], x))
				}
			default:
				panic(todo("", p.toks[0]))
			}
		case '^':
			ptr := p.pointer(false)
			switch p.rune() {
			case ')':
				dad := &DirectAbstractDeclarator{
					Case:               DirectAbstractDeclaratorDecl,
					Token:              lparen,
					AbstractDeclarator: p.abstractDeclarator(ptr, false),
					Token2:             p.must(')'),
				}
				return &AbstractDeclarator{Case: AbstractDeclaratorDecl, Pointer: ptr0, DirectAbstractDeclarator: p.directAbstractDeclarator2(dad)}
			case rune(IDENTIFIER):
				dd := p.directDeclarator(declare)
				switch p.rune() {
				case ')':
					dd = &DirectDeclarator{
						Case:  DirectDeclaratorDecl,
						Token: lparen,
						Declarator: &Declarator{
							Pointer:          ptr,
							DirectDeclarator: dd,
						},
						Token2: p.must(')'),
					}
					return &Declarator{Pointer: ptr0, DirectDeclarator: p.directDeclarator2(dd, declare)}
				default:
					panic(todo("", p.toks[0]))
				}
			default:
				panic(todo("", p.toks[0]))
			}
		case '[':
			dad := &DirectAbstractDeclarator{
				Case:               DirectAbstractDeclaratorDecl,
				Token:              lparen,
				AbstractDeclarator: p.abstractDeclarator(nil, false),
				Token2:             p.must(')'),
			}
			return &AbstractDeclarator{Case: AbstractDeclaratorDecl, Pointer: ptr0, DirectAbstractDeclarator: p.directAbstractDeclarator2(dad)}
		default:
			panic(todo("", p.toks[0]))
		}
	case '[':
		return p.abstractDeclarator(ptr0, false)
	case rune(IDENTIFIER):
		return p.declarator(ptr0, nil, declare)
	default:
		t := p.toks[0]
		x := p.checkTypeName(&t)
		panic(todo("", p.toks[0], t, x))
	}
}

func (p *parser) in(ch rune, set ...rune) bool {
	for _, v := range set {
		if ch == v {
			return true
		}
	}
	return false
}

//  type-qualifier-list:
// 	type-qualifier
// 	type-qualifier-list type-qualifier
// 	type-qualifier-list
func (p *parser) typeQualifierList(opt, acceptAttributes bool) (r *TypeQualifiers) {
	switch ch := p.rune2(); {
	case p.isTypeQualifier(ch):
		r = &TypeQualifiers{Case: TypeQualifiersTypeQual, TypeQualifier: p.typeQualifier(false)}
	case p.in(ch, ':', ')', ',', '[', rune(STATIC)) || p.isExpression(ch) || ch == rune(TYPENAME):
		if opt {
			return nil
		}

		panic(todo("", p.toks[0]))
	case ch == rune(ATTRIBUTE):
		if acceptAttributes {
			r = &TypeQualifiers{Case: TypeQualifiersTypeQual, TypeQualifier: p.typeQualifier(true)}
			break
		}

		fallthrough
	default:
		panic(todo("", p.toks[0], opt))
	}
	for {
		switch ch := p.rune(); {
		case p.in(ch, ':', ')', ',', '[', rune(STATIC)) || p.isExpression(ch) || ch == rune(TYPENAME):
			return r
		case p.isTypeQualifier(ch):
			r = &TypeQualifiers{Case: TypeQualifiersTypeQual, TypeQualifiers: r, TypeQualifier: p.typeQualifier(acceptAttributes)}
		default:
			panic(todo("", p.toks[0]))
		}
	}
}

func (p *parser) isTypeQualifier(ch rune) bool {
	switch ch {
	case
		rune(ATOMIC),
		rune(CONST),
		rune(NONNULL),
		rune(RESTRICT),
		rune(VOLATILE):

		return true
	}
	return false
}

//  pointer:
// 	* type-qualifier-list_opt
// 	* type-qualifier-list_opt pointer
//      ^ type-qualifier-list_opt
func (p *parser) pointer(opt bool) (r *Pointer) {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case '*':
		r = &Pointer{Case: PointerTypeQual, Token: p.shift(), TypeQualifiers: p.typeQualifierList(true, true)}
	case '^':
		return &Pointer{Case: PointerBlock, Token: p.shift(), TypeQualifiers: p.typeQualifierList(true, true)}
	default:
		if opt {
			return nil
		}

		panic(todo("", p.toks[0]))
	}
	prev := r
	for p.rune() == '*' {
		p := p.pointer(false)
		prev.Pointer = p
		prev = p
	}
	return r
}

//  declaration-specifiers:
// 	storage-class-specifier declaration-specifiers_opt
// 	type-specifier declaration-specifiers_opt
// 	type-qualifier declaration-specifiers_opt
// 	function-specifier declaration-specifiers_opt
//	alignment-specifier declaration-specifiers_opt
func (p *parser) declarationSpecifiers() (r *DeclarationSpecifiers, ok bool) {
	var ds, prev *DeclarationSpecifiers
	var typedef bool
	acceptTypeName := true
	for {
		switch ch := p.rune(); {
		case ch == eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil, false
		case p.isTypeSpecifier(ch, false):
			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: p.typeSpecifier()}
			acceptTypeName = false
		case ch == rune(IDENTIFIER):
			if !acceptTypeName || !p.checkTypeName(&p.toks[0]) {
				return r, true
			}

			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: p.typeSpecifier()}
			acceptTypeName = false
		case ch == rune(ATOMIC):
			switch p.peek(1).Ch {
			case '(':
				ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: p.typeSpecifier()}
				acceptTypeName = false
			default:
				ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeQual, TypeQualifier: p.typeQualifier(false)}
			}
		case ch == rune(TYPEDEF):
			typedef = true
			fallthrough
		case p.isStorageClassSpecifier(ch):
			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersStorage, StorageClassSpecifier: p.storageClassSpecifier()}
		case p.isTypeQualifier(ch):
			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeQual, TypeQualifier: p.typeQualifier(false)}
		case p.in(ch, rune(INLINE), rune(NORETURN)):
			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersFunc, FunctionSpecifier: p.functionSpecifier()}
		case ch == rune(ATTRIBUTE):
			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersAttr, AttributeSpecifierList: p.attributeSpecifierListOpt()}
		case ch == rune(ALIGNAS):
			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersAlignSpec, AlignmentSpecifier: p.alignmentSpecifier()}
		default:
			return r, true
		}
		switch {
		case r == nil:
			r = ds
		default:
			prev.DeclarationSpecifiers = ds
		}
		prev = ds
		r.typedef = typedef
		if !acceptTypeName {
			r.typename = true
		}
	}
}

func (p *parser) isDeclarationSpecifier(ch rune, typenameOk bool) bool {
	return p.isTypeSpecifier(ch, typenameOk) || p.isTypeQualifier(ch) || p.isStorageClassSpecifier(ch) ||
		ch == rune(ATTRIBUTE) || ch == rune(ALIGNAS) || p.isFunctionSpecifier(ch)
}

// [2], 6.7.5 Alignment specifier
//
// alignment-specifier:
// 	_Alignas ( type-name )
// 	_Alignas ( constant-expression )
func (p *parser) alignmentSpecifier() *AlignmentSpecifier {
	switch ch := p.peek2(2).Ch; {
	case p.isTypeSpecifier(ch, true):
		return &AlignmentSpecifier{Case: AlignmentSpecifierType, Token: p.must(rune(ALIGNAS)), Token2: p.must('('), TypeName: p.typeName(), Token3: p.must(')')}
	case p.isExpression(ch):
		return &AlignmentSpecifier{Case: AlignmentSpecifierExpr, Token: p.must(rune(ALIGNAS)), Token2: p.must('('), ConstantExpression: p.constantExpression(), Token3: p.must(')')}
	default:
		panic(todo("", p.toks[0]))
	}
}

// [0], 6.7.4 Function specifiers
//
//  function-specifier:
// 	inline
// 	_Noreturn
func (p *parser) functionSpecifier() *FunctionSpecifier {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(INLINE):
		return &FunctionSpecifier{Case: FunctionSpecifierInline, Token: p.shift()}
	case rune(NORETURN):
		return &FunctionSpecifier{Case: FunctionSpecifierNoreturn, Token: p.shift()}
	default:
		panic(todo("", p.toks[0]))
	}
}

func (p *parser) isFunctionSpecifier(ch rune) bool { return ch == rune(INLINE) || ch == rune(NORETURN) }

// [0], 6.7.3 Type qualifiers
//
//  type-qualifier:
// 	const
// 	restrict
// 	volatile
// 	_Atomic
func (p *parser) typeQualifier(acceptAttributes bool) *TypeQualifier {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(CONST):
		return &TypeQualifier{Case: TypeQualifierConst, Token: p.shift()}
	case rune(RESTRICT):
		return &TypeQualifier{Case: TypeQualifierRestrict, Token: p.shift()}
	case rune(VOLATILE):
		return &TypeQualifier{Case: TypeQualifierVolatile, Token: p.shift()}
	case rune(ATOMIC):
		return &TypeQualifier{Case: TypeQualifierAtomic, Token: p.shift()}
	case rune(NONNULL):
		return &TypeQualifier{Case: TypeQualifierNonnull, Token: p.shift()}
	case rune(ATTRIBUTE):
		if acceptAttributes {
			return &TypeQualifier{Case: TypeQualifierAttr, AttributeSpecifierList: p.attributeSpecifierListOpt()}
		}

		fallthrough
	default:
		panic(todo("", p.toks[0]))
	}
}

// [0], 6.7.1 Storage-class specifiers
//
//  storage-class-specifier:
// 	typedef
// 	extern
// 	static
// 	auto
// 	register
//	__declspec ( ... )
func (p *parser) storageClassSpecifier() (r *StorageClassSpecifier) {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(AUTO):
		return &StorageClassSpecifier{Case: StorageClassSpecifierAuto, Token: p.shift()}
	case rune(EXTERN):
		return &StorageClassSpecifier{Case: StorageClassSpecifierExtern, Token: p.shift()}
	case rune(REGISTER):
		return &StorageClassSpecifier{Case: StorageClassSpecifierRegister, Token: p.shift()}
	case rune(STATIC):
		return &StorageClassSpecifier{Case: StorageClassSpecifierStatic, Token: p.shift()}
	case rune(TYPEDEF):
		return &StorageClassSpecifier{Case: StorageClassSpecifierTypedef, Token: p.shift()}
	case rune(THREADLOCAL):
		return &StorageClassSpecifier{Case: StorageClassSpecifierThreadLocal, Token: p.shift()}
	case rune(DECLSPEC):
		r = &StorageClassSpecifier{Case: StorageClassSpecifierDeclspec, Token: p.shift(), Token2: p.must('(')}
	out:
		for lvl := 0; ; {
			switch p.rune() {
			case '(':
				lvl++
			case ')':
				if lvl == 0 {
					break out
				}

				lvl--
			case eof:
				break out
			}
			r.Declspecs = append(r.Declspecs, p.shift())
		}
		r.Token3 = p.must(')')
		return r
	default:
		panic(todo("", p.toks[0]))
	}
}

func (p *parser) isStorageClassSpecifier(ch rune) bool {
	switch ch {
	case
		rune(AUTO),
		rune(DECLSPEC),
		rune(EXTERN),
		rune(REGISTER),
		rune(STATIC),
		rune(THREADLOCAL),
		rune(TYPEDEF):

		return true
	}
	return false
}

// [0], 6.7.2 Type specifiers
//
//  type-specifier:
//	atomic-type-specifier
// 	_Bool
// 	_Complex
// 	_Decimal64
// 	_Float128
// 	_Float32
// 	_Float64
// 	__float80
// 	__fp16
// 	__m256d
// 	char
// 	double
// 	enum-specifier
// 	float
// 	int
// 	long
// 	short
// 	signed
// 	struct-or-union-specifier
// 	typedef-name
// 	typeof ( expression )
// 	typeof ( type-name )
// 	unsigned
// 	void
func (p *parser) typeSpecifier() *TypeSpecifier {
	switch p.rune2() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(CHAR):
		return &TypeSpecifier{Case: TypeSpecifierChar, Token: p.shift()}
	case rune(DOUBLE):
		return &TypeSpecifier{Case: TypeSpecifierDouble, Token: p.shift()}
	case rune(ENUM):
		return &TypeSpecifier{Case: TypeSpecifierEnum, EnumSpecifier: p.enumSpecifier()}
	case rune(FLOAT):
		return &TypeSpecifier{Case: TypeSpecifierFloat, Token: p.shift()}
	case rune(INT):
		return &TypeSpecifier{Case: TypeSpecifierInt, Token: p.shift()}
	case rune(LONG):
		return &TypeSpecifier{Case: TypeSpecifierLong, Token: p.shift()}
	case rune(SIGNED):
		return &TypeSpecifier{Case: TypeSpecifierSigned, Token: p.shift()}
	case rune(SHORT):
		return &TypeSpecifier{Case: TypeSpecifierShort, Token: p.shift()}
	case rune(TYPENAME):
		return &TypeSpecifier{Case: TypeSpecifierTypeName, Token: p.shift2()}
	case rune(UNSIGNED):
		return &TypeSpecifier{Case: TypeSpecifierUnsigned, Token: p.shift()}
	case
		rune(STRUCT),
		rune(UNION):

		return &TypeSpecifier{Case: TypeSpecifierStructOrUnion, StructOrUnionSpecifier: p.structOrUnionSpecifier()}
	case rune(VOID):
		return &TypeSpecifier{Case: TypeSpecifierVoid, Token: p.shift()}
	case rune(BOOL):
		return &TypeSpecifier{Case: TypeSpecifierBool, Token: p.shift()}
	case rune(COMPLEX):
		return &TypeSpecifier{Case: TypeSpecifierComplex, Token: p.shift()}
	case rune(IMAGINARY):
		return &TypeSpecifier{Case: TypeSpecifierImaginary, Token: p.shift()}
	case rune(FLOAT16):
		return &TypeSpecifier{Case: TypeSpecifierFloat16, Token: p.shift()}
	case rune(FLOAT32):
		return &TypeSpecifier{Case: TypeSpecifierFloat32, Token: p.shift()}
	case rune(FLOAT32X):
		return &TypeSpecifier{Case: TypeSpecifierFloat32x, Token: p.shift()}
	case rune(FLOAT64):
		return &TypeSpecifier{Case: TypeSpecifierFloat64, Token: p.shift()}
	case rune(FLOAT64X):
		return &TypeSpecifier{Case: TypeSpecifierFloat64x, Token: p.shift()}
	case rune(FLOAT128):
		return &TypeSpecifier{Case: TypeSpecifierFloat128, Token: p.shift()}
	case rune(FLOAT128X):
		return &TypeSpecifier{Case: TypeSpecifierFloat128x, Token: p.shift()}
	case rune(ATOMIC):
		return &TypeSpecifier{Case: TypeSpecifierAtomic, AtomicTypeSpecifier: p.atomicTypeSpecifier()}
	case rune(UINT128):
		return &TypeSpecifier{Case: TypeSpecifierUint128, Token: p.shift()}
	case rune(INT128):
		return &TypeSpecifier{Case: TypeSpecifierInt128, Token: p.shift()}
	case rune(M128):
		return &TypeSpecifier{Case: TypeSpecifierM128, Token: p.shift()}
	case rune(M256D):
		return &TypeSpecifier{Case: TypeSpecifierM256d, Token: p.shift()}
	case rune(DECIMAL64):
		return &TypeSpecifier{Case: TypeSpecifierDecimal64, Token: p.shift()}
	case rune(TYPEOF):
		switch ch := p.peek2(2).Ch; {
		case p.isTypeSpecifier(ch, true):
			return &TypeSpecifier{Case: TypeSpecifierTypeofType, Token: p.shift(), Token2: p.must('('), TypeName: p.typeName(), Token3: p.must(')')}
		case p.isExpression(ch):
			return &TypeSpecifier{Case: TypeSpecifierTypeofExpr, Token: p.shift(), Token2: p.must('('), Expression: p.expression(false), Token3: p.must(')')}
		default:
			panic(todo("", p.toks[0]))
		}
	default:
		panic(todo("", p.toks[0]))
	}
}

func (p *parser) isTypeSpecifier(ch rune, typenameOk bool) bool {
	switch ch {
	case
		rune(BOOL),
		rune(CHAR),
		rune(COMPLEX),
		rune(DECIMAL64),
		rune(DOUBLE),
		rune(ENUM),
		rune(FLOAT),
		rune(FLOAT128),
		rune(FLOAT128X),
		rune(FLOAT16),
		rune(FLOAT32),
		rune(FLOAT32X),
		rune(FLOAT64),
		rune(FLOAT64X),
		rune(IMAGINARY),
		rune(INT),
		rune(INT128),
		rune(LONG),
		rune(M128),
		rune(M256D),
		rune(SHORT),
		rune(SIGNED),
		rune(STRUCT),
		rune(TYPEOF),
		rune(UINT128),
		rune(UNION),
		rune(UNSIGNED),
		rune(VOID):

		return true
	case rune(TYPENAME):
		return typenameOk
	}
	return false
}

// [2], 6.7.2.4 Atomic type specifiers
//
//  atomic-type-specifier:
// 	_Atomic ( type-name )
func (p *parser) atomicTypeSpecifier() (r *AtomicTypeSpecifier) {
	return &AtomicTypeSpecifier{Token: p.must(rune(ATOMIC)), Token2: p.must('('), TypeName: p.typeName(), Token3: p.must(')')}
}

// [0], 6.7.2.2 Enumeration specifiers
//
//  enum-specifier:
// 	enum identifier_opt { enumerator-list }
// 	enum identifier_opt { enumerator-list , }
// 	enum identifier
func (p *parser) enumSpecifier() (r *EnumSpecifier) {
	switch p.peek2(1).Ch {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(IDENTIFIER):
		switch p.peek2(2).Ch {
		case '{':
			r = &EnumSpecifier{Case: EnumSpecifierDef, Token: p.shift(), Token2: p.shift(), Token3: p.shift(), EnumeratorList: p.enumeratorList()}
			r.visible = visible(r.Token.seq + 1) // [0]6.2.1,7
			p.scope.declare(string(r.Token.Src()), r)
		default:
			return &EnumSpecifier{Case: EnumSpecifierTag, Token: p.shift(), Token2: p.shift()}
		}
	case '{':
		r = &EnumSpecifier{Case: EnumSpecifierDef, Token: p.shift(), Token2: p.shift(), EnumeratorList: p.enumeratorList()}
	default:
		panic(todo("", p.peek2(1)))
	}
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case ',':
		r.Token4 = p.shift()
		r.Token5 = p.must('}')
		return r
	case '}':
		r.Token5 = p.shift()
		return r
	default:
		panic(todo("", p.toks[0]))
	}
}

//  enumerator-list:
// 	enumerator
// 	enumerator-list , enumerator
func (p *parser) enumeratorList() (r *EnumeratorList) {
	r = &EnumeratorList{Enumerator: p.enumerator()}
	prev := r
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil
		case '}':
			return r
		case ',':
			switch p.peek2(1).Ch {
			case rune(IDENTIFIER):
				e := &EnumeratorList{Token: p.shift(), Enumerator: p.enumerator()}
				prev.EnumeratorList = e
				prev = e
			default:
				return r
			}
		default:
			panic(todo("", p.toks[0]))
		}
	}
}

//  enumerator:
// 	enumeration-constant
// 	enumeration-constant = constant-expression
func (p *parser) enumerator() (r *Enumerator) {
	switch p.peek(1).Ch {
	case '}', ',':
		r = &Enumerator{Case: EnumeratorIdent, Token: p.must(rune(IDENTIFIER))}
	case '=':
		r = &Enumerator{Case: EnumeratorExpr, Token: p.must(rune(IDENTIFIER)), Token2: p.shift(), ConstantExpression: p.constantExpression()}
	default:
		panic(todo("", p.peek(1)))
	}
	r.visible = visible(r.Token.seq + 1)
	p.scope.declare(string(r.Token.Src()), r)
	return r
}

// [0], 6.7.2.1 Structure and union specifiers
//
//  struct-or-union-specifier:
// 	struct-or-union identifier_opt { struct-declaration-list }
// 	struct-or-union identifier
func (p *parser) structOrUnionSpecifier() (r *StructOrUnionSpecifier) {
	sou := p.structOrUnion()
	attrs := p.attributeSpecifierListOpt()
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case '{':
		return &StructOrUnionSpecifier{Case: StructOrUnionSpecifierDef, StructOrUnion: sou, AttributeSpecifierList: attrs, Token2: p.shift(), StructDeclarationList: p.structDeclarationList(), Token3: p.must('}'), AttributeSpecifierList2: p.attributeSpecifierListOpt()}
	case rune(IDENTIFIER):
		r = &StructOrUnionSpecifier{StructOrUnion: sou, AttributeSpecifierList: attrs, Token: p.shift()}
		r.visible = visible(r.Token.seq + 1) // [0]6.2.1,7
		p.scope.declare(string(r.Token.Src()), r)
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil
		case '{':
			r.Case = StructOrUnionSpecifierDef
			r.Token2 = p.shift()
			r.StructDeclarationList = p.structDeclarationList()
			r.Token3 = p.must('}')
			r.AttributeSpecifierList = p.attributeSpecifierListOpt()
			return r
		default:
			r.Case = StructOrUnionSpecifierTag
			return r
		}
	default:
		panic(todo("", p.toks[0]))
	}
}

//  struct-declaration-list:
// 	struct-declaration
// 	struct-declaration-list struct-declaration
func (p *parser) structDeclarationList() (r *StructDeclarationList) {
	var prev *StructDeclarationList
	for {
		switch p.rune() {
		case eof:
			p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
			return nil
		case '}':
			return r
		default:
			sdl := &StructDeclarationList{StructDeclaration: p.structDeclaration()}
			switch {
			case r == nil:
				r = sdl
			default:
				prev.StructDeclarationList = sdl
			}
			prev = sdl
		}
	}
}

//  struct-declaration:
// 	specifier-qualifier-list struct-declarator-list_opt ;
func (p *parser) structDeclaration() (r *StructDeclaration) {
	r = &StructDeclaration{SpecifierQualifierList: p.specifierQualifierList(), StructDeclaratorList: p.structDeclaratorListOpt(), AttributeSpecifierList: p.attributeSpecifierListOpt()}
	switch p.rune() {
	case ';':
		r.Token = p.shift()
		return r
	case '}':
		return r
	default:
		t := p.toks[0]
		p.cpp.eh("%v: unexpected %v, expected ';'", t.Position(), runeName(t.Ch))
		return r
	}
}

//  struct-declarator-list:
// 	struct-declarator
// 	struct-declarator-list , struct-declarator
func (p *parser) structDeclaratorListOpt() (r *StructDeclaratorList) {
	switch p.rune() {
	case
		'(',
		'*',
		':',
		rune(IDENTIFIER):

		// ok
	case ';':
		return nil
	default:
		panic(todo("", p.toks[0]))
	}

	r = &StructDeclaratorList{StructDeclarator: p.structDeclarator()}
	prev := r
	for p.rune() == ',' {
		sdl := &StructDeclaratorList{Token: p.shift(), StructDeclarator: p.structDeclarator()}
		prev.StructDeclaratorList = sdl
		prev = sdl
	}
	return r
}

//  struct-declarator:
// 	declarator
// 	declarator_opt : constant-expression
func (p *parser) structDeclarator() (r *StructDeclarator) {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case ':':
		return &StructDeclarator{Case: StructDeclaratorBitField, Token: p.shift(), ConstantExpression: p.constantExpression()}
	default:
		r = &StructDeclarator{Case: StructDeclaratorDecl, Declarator: p.declarator(nil, nil, false)}
		if p.rune() == ':' {
			r.Case = StructDeclaratorBitField
			r.Token = p.shift()
			r.ConstantExpression = p.constantExpression()
		}
		return r
	}
}

// [0], 6.6 Constant expressions
//
//  constant-expression:
// 	conditional-expression
func (p *parser) constantExpression() (r *ConstantExpression) {
	r = &ConstantExpression{}
	r.ConditionalExpression, _ = p.conditionalExpression(true)
	return r
}

// struct-or-union:
//	struct
//	union
func (p *parser) structOrUnion() *StructOrUnion {
	switch p.rune() {
	case eof:
		p.cpp.eh("%v: unexpected EOF", p.toks[0].Position())
		return nil
	case rune(STRUCT):
		return &StructOrUnion{Case: StructOrUnionStruct, Token: p.shift()}
	case rune(UNION):
		return &StructOrUnion{Case: StructOrUnionUnion, Token: p.shift()}
	default:
		panic(todo("", p.toks[0]))
	}
}

// Scope binds names to declaring nodes.
type Scope struct {
	Nodes  map[string][]Node
	Parent *Scope
}

func newScope(parent *Scope) *Scope { return &Scope{Parent: parent} }

func (s *Scope) declare(nm string, n Node) {
	// trc("%v: %q %T, visible %v (scope %p)", n.Position(), nm, n, n.(interface{ Visible() int }).Visible(), s)
	if s.Nodes != nil {
		s.Nodes[nm] = append(s.Nodes[nm], n)
		return
	}

	s.Nodes = map[string][]Node{nm: {n}}
}

type scoper struct {
	s *Scope
}

// Nodes return nodes binded to nm.
func (s *scoper) Nodes(nm string) (r []Node) {
	if s.s != nil {
		r, _ = s.s.Nodes[nm]
		return r
	}

	return nil
}

// Parent returns scope parent of a node, if any.
func (s *scoper) Parent() (r *Scope) {
	if s.s != nil {
		return s.s.Parent
	}

	return nil
}

type visible int32

// Visible reports the token sequence number where the visibility of a name starts.
func (v visible) Visible() int { return int(v) }
