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

	// More
	"__asm__": rune(ASM),
	"asm":     rune(ASM),
}

type parser struct {
	cpp *cpp
}

func newParser(cfg *Config, sources []Source) (*parser, error) {
	cpp, err := newCPP(cfg, sources, nil)
	if err != nil {
		return nil, err
	}

	cpp.rune()
	return &parser{cpp: cpp}, nil
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
		rune(ASM),
		rune(IDENTIFIER):

		return &BlockItem{Case: BlockItemStmt, Statement: p.statement()}
	case
		rune(DOUBLE),
		rune(INT):

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
func (p *parser) statement() *Statement {
	switch p.rune() {
	case rune(ASM):
		return &Statement{Case: StatementAsm, AsmStatement: p.asmStatement()}
	case rune(IDENTIFIER):
		switch p.peek(1).Ch {
		case ':':
			panic(todo("", &p.cpp.tok))
		default:
			return &Statement{Case: StatementExpr, ExpressionStatement: p.expressionStatement()}
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
	if p.rune() == ';' {
		return &ExpressionStatement{Token: p.shift()}
	}

	return &ExpressionStatement{Expression: p.expression(), Token: p.must(';')}
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

	return &AsmIndex{Token: p.shift(), Expression: p.expression(), Token2: p.must(']')}
}

// [0], 6.5.17 Comma operator
//
//  expression:
// 	assignment-expression
// 	expression , assignment-expression
func (p *parser) expression() (r *Expression) {
	r = &Expression{AssignmentExpression: p.assignmentExpression()}
	prev := r
	for p.rune() == ',' {
		ae := &Expression{Token: p.shift(), AssignmentExpression: p.assignmentExpression()}
		prev.Expression = ae
		prev = ae
	}
	return r
}

// [0], 6.7 Declarations
//
//  declaration:
// 	declaration-specifiers init-declarator-list_opt attribute-specifier-list_opt ;
func (p *parser) declaration(ds *DeclarationSpecifiers, d *Declarator) (r *Declaration) {
	if ds == nil {
		ds = p.declarationSpecifiers()
	}
	return &Declaration{DeclarationSpecifiers: ds, InitDeclaratorList: p.initDeclaratorListOpt(d), Token: p.must(';')}
}

//  init-declarator-list:
// 	init-declarator
// 	init-declarator-list , attribute-specifier-list_opt init-declarator
func (p *parser) initDeclaratorListOpt(d *Declarator) (r *InitDeclaratorList) {
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
	r = &InitDeclaratorList{InitDeclarator: p.initDeclarator(d)}
	prev := r
	for p.rune() == ',' {
		idl := &InitDeclaratorList{Token: p.shift(), InitDeclarator: p.initDeclarator(nil)}
		prev.InitDeclaratorList = idl
		prev = idl
	}
	return r
}

//  init-declarator:
// 	declarator attribute-specifier-list_opt
// 	declarator attribute-specifier-list_opt = initializer
func (p *parser) initDeclarator(d *Declarator) *InitDeclarator {
	if d == nil {
		d = p.declarator()
	}
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
		default:
			panic(todo("", &p.cpp.tok))
		}
	default:
		panic(todo("", &p.cpp.tok))
	}
	for p.rune() == ',' {
		panic(todo("", &p.cpp.tok))
	}
	return r
}

//  designation:
// 	designator-list =
func (p *parser) designation() (r *Designation) {
	r = &Designation{DesignatorList: p.designatorList()}
	if r.DesignatorList.Designator.Case != DesignatorField2 {
		r.Token = p.must('=')
	}
	return r
}

//  designator-list:
// 	designator
// 	designator-list designator
func (p *parser) designatorList() (r *DesignatorList) {
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
			return r
		}
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
		r = &ConditionalExpression{Case: ConditionalExpressionCond, LogicalOrExpression: lhs, Token: p.shift(), Expression: p.expression(), Token2: p.must(':')}
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
	switch p.rune() {
	case rune(ANDAND):
		panic(todo("", lhs, &p.cpp.tok))
	default:
		return &LogicalOrExpression{Case: LogicalOrExpressionLAnd, LogicalAndExpression: lhs}, u
	}
}

// [0], 6.5.13 Logical AND operator
//
//  logical-AND-expression:
// 	inclusive-OR-expression
// 	logical-AND-expression && inclusive-OR-expression
func (p *parser) logicalAndExpression() (r *LogicalAndExpression, u *UnaryExpression) {
	lhs, u := p.inclusiveOrExpression()
	switch p.rune() {
	case rune(OROR):
		panic(todo("", lhs, &p.cpp.tok))
	default:
		return &LogicalAndExpression{Case: LogicalAndExpressionOr, InclusiveOrExpression: lhs}, u
	}
}

// [0], 6.5.12 Bitwise inclusive OR operator
//
//  inclusive-OR-expression:
// 	exclusive-OR-expression
// 	inclusive-OR-expression | exclusive-OR-expression
func (p *parser) inclusiveOrExpression() (r *InclusiveOrExpression, u *UnaryExpression) {
	lhs, u := p.exclusiveOrExpression()
	switch p.rune() {
	case '|':
		panic(todo("", lhs, &p.cpp.tok))
	default:
		return &InclusiveOrExpression{Case: InclusiveOrExpressionXor, ExclusiveOrExpression: lhs}, u
	}
}

// [0], 6.5.11 Bitwise exclusive OR operator
//
//  exclusive-OR-expression:
// 	AND-expression
// 	exclusive-OR-expression ^ AND-expression
func (p *parser) exclusiveOrExpression() (r *ExclusiveOrExpression, u *UnaryExpression) {
	lhs, u := p.andExpression()
	switch p.rune() {
	case '^':
		panic(todo("", lhs, &p.cpp.tok))
	default:
		return &ExclusiveOrExpression{Case: ExclusiveOrExpressionAnd, AndExpression: lhs}, u
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
	switch p.rune() {
	case rune(EQ):
		panic(todo("", lhs, &p.cpp.tok))
	case rune(NEQ):
		panic(todo("", lhs, &p.cpp.tok))
	default:
		return &EqualityExpression{Case: EqualityExpressionRel, RelationalExpression: lhs}, u
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
	switch p.rune() {
	case '<':
		panic(todo("", lhs, &p.cpp.tok))
	case '>':
		panic(todo("", lhs, &p.cpp.tok))
	case rune(LEQ):
		panic(todo("", lhs, &p.cpp.tok))
	case rune(GEQ):
		panic(todo("", lhs, &p.cpp.tok))
	default:
		return &RelationalExpression{Case: RelationalExpressionShift, ShiftExpression: lhs}, u
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
	switch p.rune() {
	case rune(LSH):
		panic(todo("", lhs, &p.cpp.tok))
	case rune(RSH):
		panic(todo("", lhs, &p.cpp.tok))
	default:
		return &ShiftExpression{Case: ShiftExpressionAdd, AdditiveExpression: lhs}, u
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
		case '-':
			r = &AdditiveExpression{Case: AdditiveExpressionSub, AdditiveExpression: r, Token: p.shift()}
		default:
			return r, u
		}
		r.MultiplicativeExpression, _ = p.multiplicativeExpression()
		u = nil
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
	switch p.rune() {
	case '*':
		panic(todo("", lhs, &p.cpp.tok))
	case '/':
		panic(todo("", lhs, &p.cpp.tok))
	case '%':
		panic(todo("", lhs, &p.cpp.tok))
	default:
		return &MultiplicativeExpression{Case: MultiplicativeExpressionCast, CastExpression: lhs}, u
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
		r = &CastExpression{Case: CastExpressionCast, Token: p.shift(), TypeName: p.typeName(), Token2: p.shift()}
		r.CastExpression, _ = p.castExpression()
		return r, nil
	default:
		r = &CastExpression{Case: CastExpressionUnary, UnaryExpression: p.unaryExpression()}
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
	case ')':
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
			r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorFunc, Token: p.shift(), ParameterTypeList: p.parameterTypeList(), Token2: p.must(')')}
		default:
			t := p.peek(1)
			panic(todo("", &p.cpp.tok, &t))
		}
	case '[':
		switch p.peek(1).Ch {
		case rune(CONST):
			lbracket := p.shift()
			tql := p.typeQualifierList()
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
				r = &DirectAbstractDeclarator{Case: DirectAbstractDeclaratorStaticArr, Token: lbracket, Token2: p.shift(), TypeQualifiers: p.typeQualifierList(), AssignmentExpression: p.assignmentExpression(), Token3: p.must(']')}
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
	switch p.rune() {
	case ')':
		return r
	default:
		panic(todo("", &p.cpp.tok))
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
		case rune(INT):
			sql = &SpecifierQualifierList{Case: SpecifierQualifierListTypeSpec, TypeSpecifier: p.typeSpecifier()}
		case
			')',
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
func (p *parser) unaryExpression() (r *UnaryExpression) {
	switch p.rune() {
	case
		rune(FLOATCONST),
		rune(IDENTIFIER),
		rune(INTCONST):

		return &UnaryExpression{Case: UnaryExpressionPostfix, PostfixExpression: p.postfixExpression()}
	default:
		panic(todo("", &p.cpp.tok))
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
func (p *parser) postfixExpression() (r *PostfixExpression) {
	switch p.rune() {
	case
		rune(FLOATCONST),
		rune(IDENTIFIER),
		rune(INTCONST):

		r = &PostfixExpression{Case: PostfixExpressionPrimary, PrimaryExpression: p.primaryExpression()}
	default:
		panic(todo("", &p.cpp.tok))
	}
	for {
		switch p.rune() {
		case '[':
			panic(todo("", &p.cpp.tok))
		case '(':
			r = &PostfixExpression{Case: PostfixExpressionCall, PostfixExpression: r, Token: p.shift(), ArgumentExpressionList: p.argumentExpressionListOpt(), Token2: p.must(')')}
		case '.':
			panic(todo("", &p.cpp.tok))
		case rune(ARROW):
			panic(todo("", &p.cpp.tok))
		case rune(INC):
			panic(todo("", &p.cpp.tok))
		case rune(DEC):
			panic(todo("", &p.cpp.tok))
		default:
			return r
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
	case rune(FLOATCONST):
		return &PrimaryExpression{Case: PrimaryExpressionFloat, Token: p.shift()}
	case rune(IDENTIFIER):
		return &PrimaryExpression{Case: PrimaryExpressionIdent, Token: p.shift()}
	case rune(INTCONST):
		return &PrimaryExpression{Case: PrimaryExpressionInt, Token: p.shift()}
	default:
		panic(todo("", &p.cpp.tok))
	}
}

// [0], 6.7.5 Declarators
//
//  declarator:
// 	pointer_opt direct-declarator attribute-specifier-list_opt
func (p *parser) declarator() *Declarator {
	return &Declarator{Pointer: p.pointerOpt(), DirectDeclarator: p.directDeclarator()}
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
			default:
				t := p.peek(1)
				panic(todo("", &p.cpp.tok, &t))
			}
		case '(':
			switch p.peek(1).Ch {
			case rune(IDENTIFIER):
				r = &DirectDeclarator{Case: DirectDeclaratorFuncIdent, DirectDeclarator: r, Token: p.shift(), IdentifierList: p.identifierList(), Token2: p.must(')')}
			default:
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

	pl, comma := p.parameterList()
	switch comma.Ch {
	case ',':
		panic(todo("", pl, &p.cpp.tok))
	default:
		return &ParameterTypeList{Case: ParameterTypeListList, ParameterList: pl}
	}
}

//  parameter-list:
// 	parameter-declaration
// 	parameter-list , parameter-declaration
func (p *parser) parameterList() (r *ParameterList, t Token) {
	r = &ParameterList{ParameterDeclaration: p.parameterDeclaration()}
	for p.rune() == ',' {
		if p.peek(1).Ch == rune(DDD) {
			panic(todo("", &comma, &p.cpp.tok))
		}

		panic(todo("", &p.cpp.tok))
	}
	return r, t
}

//  parameter-declaration:
// 	declaration-specifiers declarator attribute-specifier-list_opt
// 	declaration-specifiers abstract-declarator_opt
func (p *parser) parameterDeclaration() *ParameterDeclaration {
	ds := p.declarationSpecifiers()
	switch x := p.declaratorOrAbstractDeclaratorOpt().(type) {
	case *AbstractDeclarator:
		return &ParameterDeclaration{Case: ParameterDeclarationAbstract, DeclarationSpecifiers: ds, AbstractDeclarator: x}
	case nil:
		return &ParameterDeclaration{Case: ParameterDeclarationAbstract, DeclarationSpecifiers: ds}
	default:
		panic(todo("%T %v %v", x, ds, &p.cpp.tok))
	}
}

func (p *parser) declaratorOrAbstractDeclaratorOpt() Node {
	switch p.rune() {
	case '*':
		ptr := p.pointerOpt()
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
	default:
		panic(todo("", &p.cpp.tok))
	}
}

//  type-qualifier-list:
// 	type-qualifier
// 	attribute-specifier
// 	type-qualifier-list type-qualifier
// 	type-qualifier-list attribute-specifier
func (p *parser) typeQualifierList() (r *TypeQualifiers) {
	var prev *TypeQualifiers
	for {
		var tq *TypeQualifiers
		switch p.rune() {
		case rune(CONST):
			tq = &TypeQualifiers{TypeQualifier: p.typeQualifier()}
		case
			rune(INTCONST),
			rune(STATIC):

			return r
		default:
			panic(todo("", &p.cpp.tok))
		}
		switch {
		case r == nil:
			r = tq
		default:
			prev.TypeQualifiers = tq
		}
		prev = tq
	}
}

//  pointer:
// 	* type-qualifier-list_opt
// 	* type-qualifier-list_opt pointer
//      ^ type-qualifier-list_opt
func (p *parser) pointerOpt() (r *Pointer) {
	switch p.rune() {
	case '*':
		switch p.peek(1).Ch {
		case
			')',
			rune(IDENTIFIER):

			return &Pointer{Case: PointerTypeQual, Token: p.shift()}
		default:
			panic(todo("", &p.cpp.tok))
		}
	default:
		return nil
	}
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
	for {
		switch p.rune() {
		case
			rune(CHAR),
			rune(DOUBLE),
			rune(INT),
			rune(STRUCT),
			rune(VOID):

			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: p.typeSpecifier()}
		case rune(STATIC):
			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersStorage, StorageClassSpecifier: p.storageClassSpecifier()}
		case rune(VOLATILE):
			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeQual, TypeQualifier: p.typeQualifier()}
		case rune(INLINE):
			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersFunc, FunctionSpecifier: p.functionSpecifier()}
		case
			'(',
			')',
			'*',
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
	case rune(STATIC):
		return &StorageClassSpecifier{Case: StorageClassSpecifierStatic, Token: p.shift()}
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
	case rune(INT):
		return &TypeSpecifier{Case: TypeSpecifierInt, Token: p.shift()}
	case rune(STRUCT):
		return &TypeSpecifier{Case: TypeSpecifierStructOrUnion, StructOrUnionSpecifier: p.structOrUnionSpecifier()}
	case rune(VOID):
		return &TypeSpecifier{Case: TypeSpecifierVoid, Token: p.shift()}
	default:
		panic(todo("", &p.cpp.tok))
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
		r = &StructOrUnionSpecifier{StructOrUnion: sou, Token2: p.shift()}
		switch p.rune() {
		case rune(IDENTIFIER):
			r.Case = StructOrUnionSpecifierTag
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
	for p.rune() == ',' {
		panic(todo("", &p.cpp.tok))
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
	default:
		panic(todo("", &p.cpp.tok))
	}
}
