// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"fmt"
	"runtime/debug"
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

	// More
	"_Noreturn": rune(NORETURN),
	"__asm__":   rune(ASM),
	"asm":       rune(ASM),
}

type parser struct {
	cpp   *cpp
	scope *Scope

	seq int32
}

func newParser(cfg *Config, sources []Source) (*parser, error) {
	cpp, err := newCPP(cfg, sources, nil)
	if err != nil {
		return nil, err
	}

	cpp.rune()
	return &parser{cpp: cpp, scope: &Scope{}}, nil
}

func (p *parser) newScope() { p.scope = newScope(p.scope) }

func (p *parser) closeScope() {
	if parent := p.scope.Parent; parent != nil {
		p.scope = parent
		return
	}

	p.rune()
	p.cpp.eh("%v: internal error (%v:)", p.cpp.tok.Position(), origin(1))
}

func (p *parser) isKeyword(s []byte) (r rune, ok bool) {
	r, ok = keywords[string(s)]
	return r, ok
}

func (p *parser) rune() (r rune) {
more:
	p.cpp.rune()
	p.cpp.tok = p.tok(p.cpp.tok)
	switch p.cpp.tok.Ch {
	case ' ', '\n':
		p.shift()
		goto more
	default:
		return p.cpp.tok.Ch
	}

}

func (p *parser) tok(t Token) (r Token) {
	r = t
	switch t.Ch {
	case rune(IDENTIFIER):
		if c, ok := p.isKeyword(r.Src()); ok {
			r.Ch = c
			break
		}

		nm := string(r.Src())
		for s := p.scope; s != nil; s = s.Parent {
			for _, v := range s.Nodes[nm] {
				switch x := v.(type) {
				case *Declarator:
					if x.typedef {
						r.Ch = rune(TYPENAME)
						return r
					}
				}
			}
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

func (p *parser) shift() (r Token) {
	r = p.tok(p.cpp.shift())
	p.seq++
	r.seq = p.seq
	p.rune()
	return r
}

func (p *parser) peek(i int) (r Token) {
	n := 0
	for {
		r = p.tok(p.cpp.tokenizer.peek(n).Token)
		switch r.Ch {
		case ' ', '\n':
			n++
		case eof:
			return r
		default:
			if i--; i == 0 {
				return r
			}

			n++
		}
	}
}

func (p *parser) must(r rune) Token {
	c := p.rune()
	t := p.shift()
	if c != r {
		p.cpp.eh("%v: expected %v", t.Position(), runeName(r))
	}
	return t
}

func (p *parser) parse() (*AST, error) {
	var errors errors
	p.cpp.eh = func(msg string, args ...interface{}) { errors = append(errors, fmt.Sprintf(msg, args...)) }
	tu := p.translationUnit()
	switch p.rune() {
	case eof:
		return &AST{
			TranslationUnit: tu,
			EOF:             p.shift(),
		}, nil
	default:
		panic(todo("", &p.cpp.tok))
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
	switch p.rune() {
	case rune(ASM):
		return &ExternalDeclaration{Case: ExternalDeclarationAsmStmt, AsmStatement: p.asmStatement()}
	}

	ds := p.declarationSpecifiers()
	var d *Declarator
	switch p.rune() {
	case ';', eof:
	default:
		d = p.declarator()
	}
	switch p.rune() {
	case
		',',
		';',
		'=':

		return &ExternalDeclaration{Case: ExternalDeclarationDecl, Declaration: p.declaration(ds, d)}
	case rune(ASM):
		return &ExternalDeclaration{
			Case:                  ExternalDeclarationAsm,
			AsmFunctionDefinition: &AsmFunctionDefinition{DeclarationSpecifiers: ds, Declarator: d, AsmStatement: p.asmStatement()},
		}
	case '{':
		return &ExternalDeclaration{Case: ExternalDeclarationFuncDef, FunctionDefinition: p.functionDefinition(ds, d)}
	default:
		return &ExternalDeclaration{Case: ExternalDeclarationFuncDef, FunctionDefinition: p.functionDefinition(ds, d)}
	}
}

// [0], 6.9.1 Function definitions
//
//  function-definition:
// 	declaration-specifiers declarator declaration-list_opt compound-statement
func (p *parser) functionDefinition(ds *DeclarationSpecifiers, d *Declarator) (r *FunctionDefinition) {
	return &FunctionDefinition{DeclarationSpecifiers: ds, Declarator: d, DeclarationList: p.declarationListOpt(), CompoundStatement: p.compoundStatement()}
}

//  declaration-list:
// 	declaration
// 	declaration-list declaration
func (p *parser) declarationListOpt() (r *DeclarationList) {
	var prev *DeclarationList
	for p.rune() != '{' {
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
// 	{ block-item-list_opt }
func (p *parser) compoundStatement() (r *CompoundStatement) {
	p.newScope()
	defer p.closeScope()
	return &CompoundStatement{Token: p.must('{'), BlockItemList: p.blockItemListOpt(), Token2: p.must('}')}
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
	switch p.rune() {
	case
		'{',
		rune(ASM),
		rune(DO),
		rune(FOR),
		rune(IDENTIFIER),
		rune(IF),
		rune(RETURN),
		rune(SWITCH),
		rune(WHILE):

		return &BlockItem{Case: BlockItemStmt, Statement: p.statement(false)}
	case
		rune(DOUBLE),
		rune(INT),
		rune(VOID):

		ds := p.declarationSpecifiers()
		d := p.declarator()
		switch p.rune() {
		case '{':
			return &BlockItem{Case: BlockItemFuncDef, DeclarationSpecifiers: ds, Declarator: d, CompoundStatement: p.compoundStatement()}
		default:
			return &BlockItem{Case: BlockItemDecl, Declaration: p.declaration(ds, d)}
		}
	default:
		panic(todo("", &p.cpp.tok))
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
	switch p.rune() {
	case '{':
		return &Statement{Case: StatementCompound, CompoundStatement: p.compoundStatement()}
	case rune(ASM):
		return &Statement{Case: StatementAsm, AsmStatement: p.asmStatement()}
	case rune(IDENTIFIER):
		switch p.peek(1).Ch {
		case ':':
			return &Statement{Case: StatementLabeled, LabeledStatement: p.labeledStatement()}
		default:
			return &Statement{Case: StatementExpr, ExpressionStatement: p.expressionStatement()}
		}
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
		panic(todo("", &p.cpp.tok))
	}
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
		panic(todo("", &p.cpp.tok))
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
	case rune(BREAK):
		return &JumpStatement{Case: JumpStatementBreak, Token: p.shift(), Token2: p.must(';')}
	case rune(CONTINUE):
		return &JumpStatement{Case: JumpStatementContinue, Token: p.shift(), Token2: p.must(';')}
	case rune(GOTO):
		return &JumpStatement{Case: JumpStatementGoto, Token: p.shift(), Token2: p.must(rune(IDENTIFIER)), Token3: p.must(';')}
	case rune(RETURN):
		return &JumpStatement{Case: JumpStatementReturn, Token: p.shift(), Expression: p.expression(true), Token2: p.must(';')}
	default:
		panic(todo("", &p.cpp.tok))
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
	case rune(CASE):
		return &LabeledStatement{Case: LabeledStatementCaseLabel, Token: p.shift(), ConstantExpression: p.constantExpression(), Token2: p.must(':'), Statement: p.statement(false)}
	case rune(DEFAULT):
		return &LabeledStatement{Case: LabeledStatementDefault, Token: p.shift(), Token2: p.must(':'), Statement: p.statement(false)}
	case rune(IDENTIFIER):
		return &LabeledStatement{Case: LabeledStatementLabel, Token: p.shift(), Token2: p.must(':'), Statement: p.statement(false)}
	default:
		panic(todo("", &p.cpp.tok))
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
		switch p.peek(2).Ch {
		case
			';',
			rune(IDENTIFIER):

			return &IterationStatement{Case: IterationStatementFor, Token: p.shift(), Token2: p.must('('), Expression: p.expression(true), Token3: p.must(';'), Expression2: p.expression(true), Token4: p.must(';'), Expression3: p.expression(true), Token5: p.must(')'), Statement: p.statement(true)}
		case rune(INT):
			return &IterationStatement{Case: IterationStatementForDecl, Token: p.shift(), Token2: p.must('('), Declaration: p.declaration(nil, nil), Expression: p.expression(true), Token3: p.must(';'), Expression2: p.expression(true), Token4: p.must(')'), Statement: p.statement(true)}
		default:
			t := p.peek(2)
			panic(todo("", &t))
		}
	default:
		panic(todo("", &p.cpp.tok))
	}
}

// [0], 6.8.3 Expression and null statements
//
//  expression-statement:
// 	expression_opt attribute-specifier-list_opt;
func (p *parser) expressionStatement() *ExpressionStatement {
	return &ExpressionStatement{Expression: p.expression(true), Token: p.must(';')}
}

//  asm-statement:
//  	asm attribute-specifier-list_opt ;
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
	case rune(VOLATILE):
		return &AsmQualifier{Case: AsmQualifierVolatile, Token: p.shift()}
	case rune(INLINE):
		return &AsmQualifier{Case: AsmQualifierInline, Token: p.shift()}
	case rune(GOTO):
		return &AsmQualifier{Case: AsmQualifierGoto, Token: p.shift()}
	default:
		panic(todo("", &p.cpp.tok))
	}
}

//  asm-arg-list:
// 	: asm-expression-list
// 	asm-arg-list : asm-expression-list
func (p *parser) asmArgListOpt() (r *AsmArgList) {
	var prev *AsmArgList
	for p.rune() == ':' {
		aal := &AsmArgList{Token: p.shift(), AsmExpressionList: p.asmExpressionList()}
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
	r = &AsmExpressionList{AsmIndex: p.asmIndexOpt(), AssignmentExpression: p.assignmentExpression()}
	prev := r
	for p.rune() == ',' {
		ael := &AsmExpressionList{Token: p.shift(), AsmIndex: p.asmIndexOpt(), AssignmentExpression: p.assignmentExpression()}
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
	switch p.rune() {
	case rune(IDENTIFIER):
		// ok
	case
		')',
		';':
		if opt {
			return nil
		}

		panic(todo("", &p.cpp.tok, opt))
	default:
		panic(todo("", &p.cpp.tok, opt))
	}
	r = &Expression{AssignmentExpression: p.assignmentExpression()}
	for {
		switch p.rune() {
		case ',':
			r = &Expression{Expression: r, Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
		default:
			return r
		}
	}
}

// [0], 6.7 Declarations
//
//  declaration:
// 	declaration-specifiers init-declarator-list_opt attribute-specifier-list_opt ;
func (p *parser) declaration(ds *DeclarationSpecifiers, d *Declarator) (r *Declaration) {
	if ds == nil {
		ds = p.declarationSpecifiers()
	}
	return &Declaration{DeclarationSpecifiers: ds, InitDeclaratorList: p.initDeclaratorListOpt(ds, d), Token: p.must(';')}
}

//  init-declarator-list:
// 	init-declarator
// 	init-declarator-list , attribute-specifier-list_opt init-declarator
func (p *parser) initDeclaratorListOpt(ds *DeclarationSpecifiers, d *Declarator) (r *InitDeclaratorList) {
	switch {
	case d == nil:
		switch p.rune() {
		case ';':
			return nil
		case rune(IDENTIFIER):
			d = p.declarator()
		default:
			panic(todo("", &p.cpp.tok))
		}
	}
	r = &InitDeclaratorList{InitDeclarator: p.initDeclarator(ds, d)}
	prev := r
	for p.rune() == ',' {
		idl := &InitDeclaratorList{Token: p.shift(), InitDeclarator: p.initDeclarator(ds, nil)}
		prev.InitDeclaratorList = idl
		prev = idl
	}
	return r
}

//  init-declarator:
// 	declarator attribute-specifier-list_opt
// 	declarator attribute-specifier-list_opt = initializer
func (p *parser) initDeclarator(ds *DeclarationSpecifiers, d *Declarator) *InitDeclarator {
	if d == nil {
		d = p.declarator()
	}
	d.typedef = ds.typedef
	switch p.rune() {
	case '=':
		return &InitDeclarator{Case: InitDeclaratorInit, Declarator: d, Token: p.shift(), Initializer: p.initializer()}
	default:
		return &InitDeclarator{Case: InitDeclaratorDecl, Declarator: d}
	}
}

// [0], 6.7.8 Initialization
//
//  initializer:
// 	assignment-expression
// 	{ initializer-list }
// 	{ initializer-list , }
func (p *parser) initializer() (r *Initializer) {
	switch p.rune() {
	case '{':
		r = &Initializer{Token: p.shift(), InitializerList: p.initializerList()}
		if p.rune() == ',' {
			panic(todo("", &p.cpp.tok))
		}

		r.Token2 = p.must('}')
		return r
	default:
		return &Initializer{Case: InitializerExpr, AssignmentExpression: p.assignmentExpression()}
	}
}

//  initializer-list:
// 	designation_opt initializer
// 	initializer-list , designation_opt initializer
func (p *parser) initializerList() (r *InitializerList) {
	switch p.rune() {
	case '[', '.':
		r = &InitializerList{Designation: p.designation(), Initializer: p.initializer()}
	case rune(IDENTIFIER):
		switch p.peek(1).Ch {
		case ':':
			r = &InitializerList{Designation: p.designation(), Initializer: p.initializer()}
		case '}':
			return &InitializerList{Initializer: p.initializer()}
		default:
			t := p.peek(1)
			panic(todo("", &t))
		}
	default:
		panic(todo("", &p.cpp.tok))
	}
	prev := r
	for p.rune() == ',' {
		var il *InitializerList
		switch p.peek(1).Ch {
		case '[':
			il = &InitializerList{Token: p.shift(), Designation: p.designation(), Initializer: p.initializer()}
		default:
			t := p.peek(1)
			panic(todo("", &t))
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
		switch p.rune() {
		case '[', '.':
			dl = &DesignatorList{Designator: p.designator()}
		case rune(IDENTIFIER):
			switch p.peek(1).Ch {
			case ':':
				dl = &DesignatorList{Designator: p.designator()}
			default:
				panic(todo("", &p.cpp.tok))
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
// 	. identifier
//	identifier :
func (p *parser) designator() (r *Designator) {
	switch p.rune() {
	case '[':
		return &Designator{Case: DesignatorIndex, Token: p.shift(), ConstantExpression: p.constantExpression(), Token2: p.must(']')}
	case '.':
		return &Designator{Case: DesignatorField, Token: p.shift(), Token2: p.must(rune(IDENTIFIER))}
	case rune(IDENTIFIER):
		switch p.peek(1).Ch {
		case ':':
			return &Designator{Case: DesignatorField2, Token: p.shift(), Token2: p.shift()}
		default:
			panic(todo("", &p.cpp.tok))
		}
	default:
		panic(todo("", &p.cpp.tok))
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
func (p *parser) assignmentExpression() (r *AssignmentExpression) {
	lhs, u := p.conditionalExpression()
	switch p.rune() {
	case '=':
		r = &AssignmentExpression{Case: AssignmentExpressionAssign, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
	case rune(MULASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionMul, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
	case rune(DIVASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionDiv, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
	case rune(MODASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionMod, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
	case rune(ADDASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionAdd, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
	case rune(SUBASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionSub, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
	case rune(LSHASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionLsh, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
	case rune(RSHASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionRsh, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
	case rune(ANDASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionAnd, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
	case rune(XORASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionXor, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
	case rune(ORASSIGN):
		r = &AssignmentExpression{Case: AssignmentExpressionOr, UnaryExpression: u, Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
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
func (p *parser) conditionalExpression() (r *ConditionalExpression, u *UnaryExpression) {
	lhs, u := p.logicalOrExpression()
	switch p.rune() {
	case '?':
		r = &ConditionalExpression{Case: ConditionalExpressionCond, LogicalOrExpression: lhs, Token: p.shift(), Expression: p.expression(false), Token2: p.must(':')}
		r.ConditionalExpression, _ = p.conditionalExpression()
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
func (p *parser) logicalOrExpression() (r *LogicalOrExpression, u *UnaryExpression) {
	lhs, u := p.logicalAndExpression()
	r = &LogicalOrExpression{Case: LogicalOrExpressionLAnd, LogicalAndExpression: lhs}
	for {
		switch p.rune() {
		case rune(OROR):
			r = &LogicalOrExpression{Case: LogicalOrExpressionLOr, LogicalOrExpression: r, Token: p.shift()}
			r.LogicalAndExpression, _ = p.logicalAndExpression()
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
func (p *parser) logicalAndExpression() (r *LogicalAndExpression, u *UnaryExpression) {
	lhs, u := p.inclusiveOrExpression()
	r = &LogicalAndExpression{Case: LogicalAndExpressionOr, InclusiveOrExpression: lhs}
	for {
		switch p.rune() {
		case rune(ANDAND):
			r = &LogicalAndExpression{Case: LogicalAndExpressionLAnd, LogicalAndExpression: r, Token: p.shift()}
			r.InclusiveOrExpression, _ = p.inclusiveOrExpression()
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
func (p *parser) inclusiveOrExpression() (r *InclusiveOrExpression, u *UnaryExpression) {
	lhs, u := p.exclusiveOrExpression()
	r = &InclusiveOrExpression{Case: InclusiveOrExpressionXor, ExclusiveOrExpression: lhs}
	for {
		switch p.rune() {
		case '|':
			r = &InclusiveOrExpression{Case: InclusiveOrExpressionOr, InclusiveOrExpression: r, Token: p.shift()}
			r.ExclusiveOrExpression, _ = p.exclusiveOrExpression()
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
func (p *parser) exclusiveOrExpression() (r *ExclusiveOrExpression, u *UnaryExpression) {
	lhs, u := p.andExpression()
	r = &ExclusiveOrExpression{Case: ExclusiveOrExpressionAnd, AndExpression: lhs}
	for {
		switch p.rune() {
		case '^':
			r = &ExclusiveOrExpression{Case: ExclusiveOrExpressionXor, ExclusiveOrExpression: r, Token: p.shift()}
			r.AndExpression, _ = p.andExpression()
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
func (p *parser) andExpression() (r *AndExpression, u *UnaryExpression) {
	lhs, u := p.equalityExpression()
	r = &AndExpression{Case: AndExpressionEq, EqualityExpression: lhs}
	for {
		switch p.rune() {
		case '&':
			r = &AndExpression{Case: AndExpressionAnd, AndExpression: r, Token: p.shift()}
			r.EqualityExpression, _ = p.equalityExpression()
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
func (p *parser) equalityExpression() (r *EqualityExpression, u *UnaryExpression) {
	lhs, u := p.relationalExpression()
	r = &EqualityExpression{Case: EqualityExpressionRel, RelationalExpression: lhs}
	for {
		switch p.rune() {
		case rune(EQ):
			r = &EqualityExpression{Case: EqualityExpressionEq, EqualityExpression: r, Token: p.shift()}
			r.RelationalExpression, _ = p.relationalExpression()
			u = nil
		case rune(NEQ):
			r = &EqualityExpression{Case: EqualityExpressionNeq, EqualityExpression: r, Token: p.shift()}
			r.RelationalExpression, _ = p.relationalExpression()
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
func (p *parser) relationalExpression() (r *RelationalExpression, u *UnaryExpression) {
	lhs, u := p.shiftExpression()
	r = &RelationalExpression{Case: RelationalExpressionShift, ShiftExpression: lhs}
	for {
		switch p.rune() {
		case '<':
			r = &RelationalExpression{Case: RelationalExpressionLt, RelationalExpression: r, Token: p.shift()}
			r.ShiftExpression, _ = p.shiftExpression()
			u = nil
		case '>':
			r = &RelationalExpression{Case: RelationalExpressionGt, RelationalExpression: r, Token: p.shift()}
			r.ShiftExpression, _ = p.shiftExpression()
			u = nil
		case rune(LEQ):
			r = &RelationalExpression{Case: RelationalExpressionLeq, RelationalExpression: r, Token: p.shift()}
			r.ShiftExpression, _ = p.shiftExpression()
			u = nil
		case rune(GEQ):
			r = &RelationalExpression{Case: RelationalExpressionGeq, RelationalExpression: r, Token: p.shift()}
			r.ShiftExpression, _ = p.shiftExpression()
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
func (p *parser) shiftExpression() (r *ShiftExpression, u *UnaryExpression) {
	lhs, u := p.additiveExpression()
	r = &ShiftExpression{Case: ShiftExpressionAdd, AdditiveExpression: lhs}
	for {
		switch p.rune() {
		case rune(LSH):
			r = &ShiftExpression{Case: ShiftExpressionLsh, ShiftExpression: r, Token: p.shift()}
			r.AdditiveExpression, _ = p.additiveExpression()
			u = nil
		case rune(RSH):
			r = &ShiftExpression{Case: ShiftExpressionRsh, ShiftExpression: r, Token: p.shift()}
			r.AdditiveExpression, _ = p.additiveExpression()
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
func (p *parser) additiveExpression() (r *AdditiveExpression, u *UnaryExpression) {
	lhs, u := p.multiplicativeExpression()
	r = &AdditiveExpression{Case: AdditiveExpressionMul, MultiplicativeExpression: lhs}
	for {
		switch p.rune() {
		case '+':
			r = &AdditiveExpression{Case: AdditiveExpressionAdd, AdditiveExpression: r, Token: p.shift()}
			r.MultiplicativeExpression, _ = p.multiplicativeExpression()
			u = nil
		case '-':
			r = &AdditiveExpression{Case: AdditiveExpressionSub, AdditiveExpression: r, Token: p.shift()}
			r.MultiplicativeExpression, _ = p.multiplicativeExpression()
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
func (p *parser) multiplicativeExpression() (r *MultiplicativeExpression, u *UnaryExpression) {
	lhs, u := p.castExpression()
	r = &MultiplicativeExpression{Case: MultiplicativeExpressionCast, CastExpression: lhs}
	for {
		switch p.rune() {
		case '*':
			r = &MultiplicativeExpression{Case: MultiplicativeExpressionMul, MultiplicativeExpression: r, Token: p.shift()}
			r.CastExpression, _ = p.castExpression()
			u = nil
		case '/':
			r = &MultiplicativeExpression{Case: MultiplicativeExpressionDiv, MultiplicativeExpression: r, Token: p.shift()}
			r.CastExpression, _ = p.castExpression()
			u = nil
		case '%':
			r = &MultiplicativeExpression{Case: MultiplicativeExpressionMod, MultiplicativeExpression: r, Token: p.shift()}
			r.CastExpression, _ = p.castExpression()
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
func (p *parser) castExpression() (r *CastExpression, u *UnaryExpression) {
	switch p.rune() {
	case '(':
		switch p.peek(1).Ch {
		case rune(INT):
			lparen := p.shift()
			tn := p.typeName()
			rparen := p.shift()
			switch p.rune() {
			case '{':
				u = p.unaryExpression(lparen, tn, rparen)
				return &CastExpression{Case: CastExpressionUnary, UnaryExpression: u}, u
			default:
				r = &CastExpression{Case: CastExpressionCast, Token: lparen, TypeName: tn, Token2: rparen}
				r.CastExpression, _ = p.castExpression()
				return r, nil
			}
		case
			'{',
			rune(IDENTIFIER):

			r = &CastExpression{Case: CastExpressionUnary, UnaryExpression: p.unaryExpression(Token{}, nil, Token{})}
			return r, r.UnaryExpression
		default:
			t := p.peek(1)
			panic(todo("", &t))
		}
	default:
		r = &CastExpression{Case: CastExpressionUnary, UnaryExpression: p.unaryExpression(Token{}, nil, Token{})}
		return r, r.UnaryExpression
	}
}

// [0], 6.7.6 Type names
//
//  type-name:
// 	specifier-qualifier-list abstract-declarator_opt
func (p *parser) typeName() *TypeName {
	return &TypeName{SpecifierQualifierList: p.specifierQualifierList(), AbstractDeclarator: p.abstractDeclarator(true)}
}

//  abstract-declarator:
// 	pointer
// 	pointer_opt direct-abstract-declarator
func (p *parser) abstractDeclarator(opt bool) *AbstractDeclarator {
	switch p.rune() {
	case
		'(',
		'[':

		return &AbstractDeclarator{Case: AbstractDeclaratorDecl, DirectAbstractDeclarator: p.directAbstractDeclarator()}
	case
		')',
		rune(IDENTIFIER):

		if opt {
			return nil
		}

		panic(todo("", opt, &p.cpp.tok))
	default:
		panic(todo("", opt, &p.cpp.tok))
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
	switch p.rune() {
	case '(':
		switch p.peek(1).Ch {
		case ')': // ()
			r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorFunc, Token: p.shift(), Token2: p.shift()}
		case rune(CHAR):
			p.newScope()
			defer func() {
				r.params = p.scope
				p.closeScope()
			}()
			r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorFunc, Token: p.shift(), ParameterTypeList: p.parameterTypeList(), Token2: p.must(')')}
		default:
			t := p.peek(1)
			panic(todo("", &p.cpp.tok, &t))
		}
	case '[':
		switch p.peek(1).Ch {
		case ']':
			r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorArr, Token: p.shift(), Token2: p.shift()}
		case rune(CONST):
			lbracket := p.shift()
			tql := p.typeQualifierList(false)
			switch p.rune() {
			case rune(INTCONST):
				r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorArr, Token: lbracket, TypeQualifiers: tql, AssignmentExpression: p.assignmentExpression(), Token2: p.must(']')}
			case rune(STATIC):
				r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorArr, Token: lbracket, TypeQualifiers: tql, Token2: p.shift(), AssignmentExpression: p.assignmentExpression(), Token3: p.must(']')}
			default:
				panic(todo("", &lbracket, tql, &p.cpp.tok))
			}
		case rune(STATIC):
			lbracket := p.shift()
			switch p.peek(1).Ch {
			case rune(CONST):
				r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorStaticArr, Token: lbracket, Token2: p.shift(), TypeQualifiers: p.typeQualifierList(false), AssignmentExpression: p.assignmentExpression(), Token3: p.must(']')}
			default:
				t := p.peek(1)
				panic(todo("", &lbracket, &t))
			}
		case '*':
			switch p.peek(2).Ch {
			case ']':
				r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorArrStar, Token: p.shift(), Token2: p.shift(), Token3: p.shift()}
			default:
				t := p.peek(2)
				panic(todo("", &t))
			}
		default:
			t := p.peek(1)
			panic(todo("", &t))
		}
	default:
		panic(todo("", &p.cpp.tok))
	}
	for {
		switch p.rune() {
		case ')':
			return r
		default:
			panic(todo("", &p.cpp.tok))
		}
	}
}

//  specifier-qualifier-list:
// 	type-specifier specifier-qualifier-list_opt
// 	type-qualifier specifier-qualifier-list_opt
// 	alignment-specifier specifier-qualifier-list_opt
func (p *parser) specifierQualifierList() (r *SpecifierQualifierList) {
	var prev *SpecifierQualifierList
	for {
		var sql *SpecifierQualifierList
		switch p.rune() {
		case
			rune(DOUBLE),
			rune(INT):

			sql = &SpecifierQualifierList{Case: SpecifierQualifierListTypeSpec, TypeSpecifier: p.typeSpecifier()}
		case rune(CONST):
			sql = &SpecifierQualifierList{Case: SpecifierQualifierListTypeSpec, TypeQualifier: p.typeQualifier()}
		case
			')',
			'[',
			rune(IDENTIFIER):

			return r
		default:
			panic(todo("", &p.cpp.tok))
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
func (p *parser) unaryExpression(lp Token, tn *TypeName, rp Token) (r *UnaryExpression) {
	if tn != nil {
		return &UnaryExpression{Case: UnaryExpressionPostfix, PostfixExpression: p.postfixExpression(lp, tn, rp)}
	}

	switch p.rune() {
	case '&':
		r = &UnaryExpression{Case: UnaryExpressionAddrof, Token: p.shift()}
		r.CastExpression, _ = p.castExpression()
		return r
	case '*':
		r = &UnaryExpression{Case: UnaryExpressionDeref, Token: p.shift()}
		r.CastExpression, _ = p.castExpression()
		return r
	case '+':
		r = &UnaryExpression{Case: UnaryExpressionPlus, Token: p.shift()}
		r.CastExpression, _ = p.castExpression()
		return r
	case '-':
		r = &UnaryExpression{Case: UnaryExpressionMinus, Token: p.shift()}
		r.CastExpression, _ = p.castExpression()
		return r
	case '~':
		r = &UnaryExpression{Case: UnaryExpressionCpl, Token: p.shift()}
		r.CastExpression, _ = p.castExpression()
		return r
	case '!':
		r = &UnaryExpression{Case: UnaryExpressionNot, Token: p.shift()}
		r.CastExpression, _ = p.castExpression()
		return r
	case rune(SIZEOF):
		switch p.peek(1).Ch {
		case '(':
			return &UnaryExpression{Case: UnaryExpressionSizeofType, Token: p.shift(), Token2: p.shift(), TypeName: p.typeName(), Token3: p.must(')')}
		default:
			return &UnaryExpression{Case: UnaryExpressionSizeofExpr, Token: p.shift(), UnaryExpression: p.unaryExpression(Token{}, nil, Token{})}
		}
	case
		'(',
		rune(CHARCONST),
		rune(FLOATCONST),
		rune(IDENTIFIER),
		rune(INTCONST),
		rune(LONGCHARCONST),
		rune(LONGSTRINGLITERAL),
		rune(STRINGLITERAL):

		return &UnaryExpression{Case: UnaryExpressionPostfix, PostfixExpression: p.postfixExpression(Token{}, nil, Token{})}
	case rune(INC):
		return &UnaryExpression{Case: UnaryExpressionInc, Token: p.shift(), UnaryExpression: p.unaryExpression(Token{}, nil, Token{})}
	case rune(DEC):
		return &UnaryExpression{Case: UnaryExpressionDec, Token: p.shift(), UnaryExpression: p.unaryExpression(Token{}, nil, Token{})}
	default:
		panic(todo("%v\n%s", &p.cpp.tok, debug.Stack()))
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
func (p *parser) postfixExpression(lp Token, tn *TypeName, rp Token) (r *PostfixExpression) {
	switch {
	case tn != nil:
		r = &PostfixExpression{Case: PostfixExpressionComplit, Token: lp, TypeName: tn, Token2: rp, Token3: p.must('{'), InitializerList: p.initializerList()}
		switch p.rune() {
		case ',':
			panic(todo("", &p.cpp.tok))
		default:
			r.Token5 = p.must('}')
		}
	default:
		switch p.rune() {
		case '(':
			switch p.peek(1).Ch {
			case
				'{',
				rune(IDENTIFIER):

				r = &PostfixExpression{Case: PostfixExpressionPrimary, PrimaryExpression: p.primaryExpression()}
			default:
				t := p.peek(1)
				panic(todo("", &t))
			}
		default:
			r = &PostfixExpression{Case: PostfixExpressionPrimary, PrimaryExpression: p.primaryExpression()}
		}
	}
	for {
		switch p.rune() {
		case '[':
			r = &PostfixExpression{Case: PostfixExpressionIndex, PostfixExpression: r, Token: p.shift(), Expression: p.expression(false), Token2: p.must(']')}
		case '(':
			r = &PostfixExpression{Case: PostfixExpressionCall, PostfixExpression: r, Token: p.shift(), ArgumentExpressionList: p.argumentExpressionListOpt(), Token2: p.must(')')}
		case '.':
			r = &PostfixExpression{Case: PostfixExpressionSelect, PostfixExpression: r, Token: p.shift(), Token2: p.must(rune(IDENTIFIER))}
		case rune(ARROW):
			r = &PostfixExpression{Case: PostfixExpressionPSelect, PostfixExpression: r, Token: p.shift(), Token2: p.must(rune(IDENTIFIER))}
		case rune(INC):
			r = &PostfixExpression{Case: PostfixExpressionInc, PostfixExpression: r, Token: p.shift()}
		case rune(DEC):
			r = &PostfixExpression{Case: PostfixExpressionDec, PostfixExpression: r, Token: p.shift()}
		case
			'%',
			'&',
			')',
			'*',
			'+',
			',',
			'-',
			'/',
			':',
			';',
			'<',
			'=',
			'>',
			'?',
			']',
			'^',
			'|',
			'}',
			rune(ADDASSIGN),
			rune(ANDAND),
			rune(ANDASSIGN),
			rune(DIVASSIGN),
			rune(EQ),
			rune(GEQ),
			rune(LEQ),
			rune(LSH),
			rune(LSHASSIGN),
			rune(MODASSIGN),
			rune(MULASSIGN),
			rune(NEQ),
			rune(ORASSIGN),
			rune(OROR),
			rune(RSH),
			rune(RSHASSIGN),
			rune(SUBASSIGN),
			rune(XORASSIGN):

			return r
		default:
			panic(todo("", &p.cpp.tok))
		}
	}
}

//  argument-expression-list:
// 	assignment-expression
// 	argument-expression-list , assignment-expression
func (p *parser) argumentExpressionListOpt() (r *ArgumentExpressionList) {
	if p.rune() == ')' {
		return nil
	}

	r = &ArgumentExpressionList{AssignmentExpression: p.assignmentExpression()}
	prev := r
	for p.rune() == ',' {
		prev.ArgumentExpressionList = &ArgumentExpressionList{Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
	}
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
func (p *parser) primaryExpression() *PrimaryExpression {
	switch p.rune() {
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
		case rune(IDENTIFIER):
			return &PrimaryExpression{Case: PrimaryExpressionExpr, Token: p.shift(), Expression: p.expression(false), Token2: p.must(')')}
		case '{':
			return &PrimaryExpression{Case: PrimaryExpressionStmt, Token: p.shift(), CompoundStatement: p.compoundStatement(), Token2: p.must(')')}
		default:
			t := p.peek(1)
			panic(todo("", &t))
		}
	default:
		panic(todo("", &p.cpp.tok))
	}
}

// [0], 6.7.5 Declarators
//
//  declarator:
// 	pointer_opt direct-declarator attribute-specifier-list_opt
func (p *parser) declarator() (r *Declarator) {
	r = &Declarator{Pointer: p.pointer(true), DirectDeclarator: p.directDeclarator()}
	p.scope.insert(r.Name(), r)
	return r
}

//  direct-declarator:
// 	identifier asm_opt
// 	( attribute-specifier-list_opt declarator )
// 	direct-declarator [ type-qualifier-list_opt assignment-expression_opt ]
// 	direct-declarator [ static type-qualifier-list_opt assignment-expression ]
// 	direct-declarator [ type-qualifier-list static assignment-expression ]
// 	direct-declarator [ type-qualifier-list_opt * ]
// 	direct-declarator ( parameter-type-list )
// 	direct-declarator ( identifier-list_opt )
func (p *parser) directDeclarator() (r *DirectDeclarator) {
	switch p.rune() {
	case rune(IDENTIFIER):
		r = &DirectDeclarator{Case: DirectDeclaratorIdent, Token: p.shift()}
	case '(':
		r = &DirectDeclarator{Case: DirectDeclaratorDecl, Token: p.shift(), Declarator: p.declarator(), Token2: p.must(')')}
	default:
		panic(todo("", &p.cpp.tok))
	}
	for {
		switch p.rune() {
		case '[':
			switch p.peek(1).Ch {
			case ']':
				r = &DirectDeclarator{Case: DirectDeclaratorArr, DirectDeclarator: r, Token: p.shift(), Token2: p.shift()}
			case rune(INTCONST):
				r = &DirectDeclarator{Case: DirectDeclaratorArr, DirectDeclarator: r, Token: p.shift(), AssignmentExpression: p.assignmentExpression(), Token2: p.shift()}
			case rune(CONST):
				lbracket := p.shift()
				tql := p.typeQualifierList(false)
				switch p.rune() {
				case rune(INTCONST):
					r = &DirectDeclarator{Case: DirectDeclaratorArr, DirectDeclarator: r, Token: lbracket, TypeQualifiers: tql, AssignmentExpression: p.assignmentExpression(), Token2: p.must(']')}
				case rune(STATIC):
					r = &DirectDeclarator{Case: DirectDeclaratorArrStatic, DirectDeclarator: r, Token: lbracket, TypeQualifiers: tql, Token2: p.shift(), AssignmentExpression: p.assignmentExpression(), Token3: p.must(']')}
				case '*':
					r = &DirectDeclarator{Case: DirectDeclaratorStar, DirectDeclarator: r, Token: lbracket, TypeQualifiers: tql, Token2: p.shift(), Token3: p.must(']')}
				default:
					panic(todo("", &p.cpp.tok))
				}
			case rune(STATIC):
				switch p.peek(2).Ch {
				case rune(CONST):
					r = &DirectDeclarator{Case: DirectDeclaratorStaticArr, DirectDeclarator: r, Token: p.shift(), Token2: p.shift(), TypeQualifiers: p.typeQualifierList(false), AssignmentExpression: p.assignmentExpression(), Token3: p.must(']')}
				default:
					t := p.peek(2)
					panic(todo("", &t))
				}
			default:
				t := p.peek(1)
				panic(todo("", &p.cpp.tok, &t))
			}
		case '(':
			switch p.peek(1).Ch {
			case rune(IDENTIFIER):
				r = &DirectDeclarator{Case: DirectDeclaratorFuncIdent, DirectDeclarator: r, Token: p.shift(), IdentifierList: p.identifierList(), Token2: p.must(')')}
			default:
				p.newScope()
				defer func() {
					r.params = p.scope
					p.closeScope()
				}()
				r = &DirectDeclarator{Case: DirectDeclaratorFuncParam, DirectDeclarator: r, Token: p.shift(), ParameterTypeList: p.parameterTypeList(), Token2: p.must(')')}
			}
		default:
			return r
		}
	}
}

//  identifier-list:
// 	identifier
// 	identifier-list , identifier
func (p *parser) identifierList() (r *IdentifierList) {
	switch p.rune() {
	case rune(IDENTIFIER):
		r = &IdentifierList{Token: p.shift()}
	default:
		panic(todo("", &p.cpp.tok))
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
func (p *parser) parameterTypeList() (r *ParameterTypeList) {
	if p.rune() == ')' {
		return nil
	}

	pl := p.parameterList()
	switch p.rune() {
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
// 	declaration-specifiers declarator attribute-specifier-list_opt
// 	declaration-specifiers abstract-declarator_opt
func (p *parser) parameterDeclaration() *ParameterDeclaration {
	ds := p.declarationSpecifiers()
	switch x := p.declaratorOrAbstractDeclaratorOpt().(type) {
	case *AbstractDeclarator:
		return &ParameterDeclaration{Case: ParameterDeclarationAbstract, DeclarationSpecifiers: ds, AbstractDeclarator: x}
	case *Declarator:
		return &ParameterDeclaration{Case: ParameterDeclarationDecl, DeclarationSpecifiers: ds, Declarator: x}
	case nil:
		return &ParameterDeclaration{Case: ParameterDeclarationAbstract, DeclarationSpecifiers: ds}
	default:
		panic(todo("%T %v %v", x, ds, &p.cpp.tok))
	}
}

func (p *parser) declaratorOrAbstractDeclaratorOpt() Node {
	switch p.rune() {
	case '*':
		ptr := p.pointer(false)
		switch p.rune() {
		case ')':
			return &AbstractDeclarator{Case: AbstractDeclaratorPtr, Pointer: ptr}
		default:
			panic(todo("", ptr, &p.cpp.tok))
		}
	case
		'(',
		'[':

		return p.abstractDeclarator(false)
	case ')':
		return nil
	case rune(IDENTIFIER):
		return p.declarator()
	default:
		panic(todo("", &p.cpp.tok))
	}
}

//  type-qualifier-list:
// 	type-qualifier
// 	attribute-specifier
// 	type-qualifier-list type-qualifier
// 	type-qualifier-list attribute-specifier
func (p *parser) typeQualifierList(opt bool) (r *TypeQualifiers) {
	switch p.rune() {
	case rune(CONST):
		r = &TypeQualifiers{Case: TypeQualifiersTypeQual, TypeQualifier: p.typeQualifier()}
	case
		')',
		'*',
		rune(IDENTIFIER):

		if opt {
			return nil
		}

		panic(todo("", &p.cpp.tok))
	default:
		panic(todo("", &p.cpp.tok, opt))
	}
	for {
		switch p.rune() {
		case
			'*',
			rune(IDENTIFIER),
			rune(INTCONST),
			rune(STATIC):

			return r
		case rune(VOLATILE):
			r = &TypeQualifiers{Case: TypeQualifiersTypeQual, TypeQualifiers: r, TypeQualifier: p.typeQualifier()}
		default:
			panic(todo("", &p.cpp.tok))
		}
	}
}

//  pointer:
// 	* type-qualifier-list_opt
// 	* type-qualifier-list_opt pointer
//      ^ type-qualifier-list_opt
func (p *parser) pointer(opt bool) (r *Pointer) {
	switch p.rune() {
	case '*':
		r = &Pointer{Case: PointerTypeQual, Token: p.shift(), TypeQualifiers: p.typeQualifierList(true)}
	default:
		if opt {
			return nil
		}

		panic(todo("", &p.cpp.tok))
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
//	attribute-specifier declaration-specifiers_opt
func (p *parser) declarationSpecifiers() (r *DeclarationSpecifiers) {
	var ds, prev *DeclarationSpecifiers
	var typedef bool
	for {
		switch p.rune() {
		case
			rune(CHAR),
			rune(DOUBLE),
			rune(ENUM),
			rune(FLOAT),
			rune(INT),
			rune(LONG),
			rune(SHORT),
			rune(SIGNED),
			rune(STRUCT),
			rune(TYPENAME),
			rune(UNION),
			rune(UNSIGNED),
			rune(VOID):

			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: p.typeSpecifier()}
		case rune(TYPEDEF):
			typedef = true
			fallthrough
		case
			rune(AUTO),
			rune(EXTERN),
			rune(REGISTER),
			rune(STATIC):

			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersStorage, StorageClassSpecifier: p.storageClassSpecifier()}
		case
			rune(CONST),
			rune(RESTRICT),
			rune(VOLATILE):

			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeQual, TypeQualifier: p.typeQualifier()}
		case
			rune(INLINE),
			rune(NORETURN):

			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersFunc, FunctionSpecifier: p.functionSpecifier()}
		case
			'(',
			')',
			'*',
			':',
			';',
			'[',
			'}',
			rune(IDENTIFIER):

			return r
		default:
			panic(todo("", ds, &p.cpp.tok))
		}
		switch {
		case r == nil:
			r = ds
		default:
			prev.DeclarationSpecifiers = ds
		}
		prev = ds
		r.typedef = typedef
	}
}

// [0], 6.7.4 Function specifiers
//
//  function-specifier:
// 	inline
// 	_Noreturn
func (p *parser) functionSpecifier() *FunctionSpecifier {
	switch p.rune() {
	case rune(INLINE):
		return &FunctionSpecifier{Case: FunctionSpecifierInline, Token: p.shift()}
	case rune(NORETURN):
		return &FunctionSpecifier{Case: FunctionSpecifierNoreturn, Token: p.shift()}
	default:
		panic(todo("", &p.cpp.tok))
	}
}

// [0], 6.7.3 Type qualifiers
//
//  type-qualifier:
// 	const
// 	restrict
// 	volatile
// 	_Atomic
func (p *parser) typeQualifier() *TypeQualifier {
	switch p.rune() {
	case rune(CONST):
		return &TypeQualifier{Case: TypeQualifierConst, Token: p.shift()}
	case rune(RESTRICT):
		return &TypeQualifier{Case: TypeQualifierRestrict, Token: p.shift()}
	case rune(VOLATILE):
		return &TypeQualifier{Case: TypeQualifierVolatile, Token: p.shift()}
	default:
		panic(todo("", &p.cpp.tok))
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
func (p *parser) storageClassSpecifier() *StorageClassSpecifier {
	switch p.rune() {
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
	default:
		panic(todo("", &p.cpp.tok))
	}
}

// [0], 6.7.2 Type specifiers
//
//  type-specifier:
// 	void
// 	char
// 	short
// 	int
// 	long
// 	float
// 	__fp16
// 	__float80
// 	double
// 	signed
// 	unsigned
// 	_Bool
// 	_Complex
// 	_Float128
// 	struct-or-union-specifier
// 	enum-specifier
// 	typedef-name
// 	typeof ( expression )
// 	typeof ( type-name )
//	atomic-type-specifier
//	_Frac
//	_Sat
//	_Accum
// 	_Float32
func (p *parser) typeSpecifier() *TypeSpecifier {
	switch p.rune() {
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
		return &TypeSpecifier{Case: TypeSpecifierTypeName, Token: p.shift()}
	case rune(UNSIGNED):
		return &TypeSpecifier{Case: TypeSpecifierUnsigned, Token: p.shift()}
	case
		rune(STRUCT),
		rune(UNION):

		return &TypeSpecifier{Case: TypeSpecifierStructOrUnion, StructOrUnionSpecifier: p.structOrUnionSpecifier()}
	case rune(VOID):
		return &TypeSpecifier{Case: TypeSpecifierVoid, Token: p.shift()}
	default:
		panic(todo("", &p.cpp.tok))
	}
}

// [0], 6.7.2.2 Enumeration specifiers
//
//  enum-specifier:
// 	enum attribute-specifier-list_opt identifier_opt { enumerator-list }
// 	enum attribute-specifier-list_opt identifier_opt { enumerator-list , }
// 	enum attribute-specifier-list_opt identifier
func (p *parser) enumSpecifier() (r *EnumSpecifier) {
	switch p.peek(1).Ch {
	case rune(IDENTIFIER):
		switch p.peek(2).Ch {
		case '{':
			r = &EnumSpecifier{Case: EnumSpecifierDef, Token: p.shift(), Token2: p.shift(), Token3: p.shift(), EnumeratorList: p.enumeratorList()}
			switch p.rune() {
			case '}':
				r.Token3 = p.shift()
				return r
			default:
				panic(todo("", &p.cpp.tok))
			}
		case rune(IDENTIFIER):
			return &EnumSpecifier{Case: EnumSpecifierTag, Token: p.shift(), Token2: p.shift()}
		default:
			t := p.peek(2)
			panic(todo("", &t))
		}
	default:
		t := p.peek(1)
		panic(todo("", &t))
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
		case '}':
			return r
		case ',':
			switch p.peek(1).Ch {
			case rune(IDENTIFIER):
				e := &EnumeratorList{Token: p.shift(), Enumerator: p.enumerator()}
				prev.EnumeratorList = e
				prev = e
			default:
				t := p.peek(1)
				panic(todo("", &t))
			}
		default:
			panic(todo("", &p.cpp.tok))
		}
	}
}

//  enumerator:
// 	enumeration-constant attribute-specifier-list_opt
// 	enumeration-constant attribute-specifier-list_opt = constant-expression
func (p *parser) enumerator() (r *Enumerator) {
	switch p.peek(1).Ch {
	case '}', ',':
		return &Enumerator{Case: EnumeratorIdent, Token: p.must(rune(IDENTIFIER))}
	case '=':
		return &Enumerator{Case: EnumeratorExpr, Token: p.must(rune(IDENTIFIER)), Token2: p.shift(), ConstantExpression: p.constantExpression()}
	default:
		t := p.peek(1)
		panic(todo("", &t))
	}
}

// [0], 6.7.2.1 Structure and union specifiers
//
//  struct-or-union-specifier:
// 	struct-or-union attribute-specifier-list_opt identifier_opt { struct-declaration-list }
// 	struct-or-union attribute-specifier-list_opt identifier
func (p *parser) structOrUnionSpecifier() (r *StructOrUnionSpecifier) {
	sou := p.structOrUnion()
	switch p.rune() {
	case '{':
		return &StructOrUnionSpecifier{Case: StructOrUnionSpecifierDef, StructOrUnion: sou, Token2: p.shift(), StructDeclarationList: p.structDeclarationList(), Token3: p.must('}')}
	case rune(IDENTIFIER):
		r = &StructOrUnionSpecifier{StructOrUnion: sou, Token: p.shift()}
		switch p.rune() {
		case rune(IDENTIFIER):
			r.Case = StructOrUnionSpecifierTag
			return r
		case '{':
			r.Case = StructOrUnionSpecifierDef
			r.Token2 = p.shift()
			r.StructDeclarationList = p.structDeclarationList()
			r.Token3 = p.must('}')
			return r
		default:
			panic(todo("", &p.cpp.tok))
		}
	default:
		panic(todo("", sou, &p.cpp.tok))
	}
}

//  struct-declaration-list:
// 	struct-declaration
// 	struct-declaration-list struct-declaration
func (p *parser) structDeclarationList() (r *StructDeclarationList) {
	p.newScope()
	defer p.closeScope()
	var prev *StructDeclarationList
	for {
		switch p.rune() {
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
// 	specifier-qualifier-list struct-declarator-list ;
func (p *parser) structDeclaration() (r *StructDeclaration) {
	return &StructDeclaration{SpecifierQualifierList: p.specifierQualifierList(), StructDeclaratorList: p.structDeclaratorList(), Token: p.must(';')}
}

//  struct-declarator-list:
// 	struct-declarator
// 	struct-declarator-list , struct-declarator
func (p *parser) structDeclaratorList() (r *StructDeclaratorList) {
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
// 	declarator_opt : constant-expression attribute-specifier-list_opt
func (p *parser) structDeclarator() (r *StructDeclarator) {
	switch p.rune() {
	case ':':
		panic(todo("", &p.cpp.tok))
	default:
		r = &StructDeclarator{Case: StructDeclaratorDecl, Declarator: p.declarator()}
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
	r.ConditionalExpression, _ = p.conditionalExpression()
	return r
}

// struct-or-union:
//	struct
//	union
func (p *parser) structOrUnion() *StructOrUnion {
	switch p.rune() {
	case rune(STRUCT):
		return &StructOrUnion{Case: StructOrUnionStruct, Token: p.shift()}
	case rune(UNION):
		return &StructOrUnion{Case: StructOrUnionUnion, Token: p.shift()}
	default:
		panic(todo("", &p.cpp.tok))
	}
}

// Scope binds names to declaring nodes.
type Scope struct {
	Nodes  map[string][]Node
	Parent *Scope
}

func newScope(parent *Scope) *Scope { return &Scope{Parent: parent} }

func (s *Scope) insert(nm string, n Node) {
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
