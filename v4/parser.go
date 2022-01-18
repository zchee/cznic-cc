// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"fmt"
)

var keywords = map[string]rune{
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
	switch r = p.cpp.rune(); r {
	case rune(IDENTIFIER):
		if r2, ok := p.isKeyword(p.cpp.tok.Src()); ok {
			p.cpp.tok.Ch = r2
			r = r2
		}
	case ' ', '\n':
		p.shift()
		goto more
	}
	return r
}

func (p *parser) shift() Token { return p.cpp.shift() }

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
		ed := p.externalDeclaration()
		if ed == nil {
			return r
		}

		tu := &TranslationUnit{ExternalDeclaration: ed}
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
// 	;
func (p *parser) externalDeclaration() *ExternalDeclaration {
	ds := p.declarationSpecifiers()
	d := p.declarator()
	switch p.rune() {
	case ';':
		return &ExternalDeclaration{
			Case: ExternalDeclarationDecl,
			Declaration: &Declaration{
				DeclarationSpecifiers: ds,
				InitDeclaratorList: &InitDeclaratorList{
					InitDeclarator: &InitDeclarator{
						Case:       InitDeclaratorDecl,
						Declarator: d,
					},
				},
				Token: p.shift(),
			},
		}
	case '=':
		return &ExternalDeclaration{Case: ExternalDeclarationDecl, Declaration: p.declaration(ds, d)}
	default:
		panic(todo("", ds, d, &p.cpp.tok))
	}
}

// [0], 6.7 Declarations
//
//  declaration:
// 	declaration-specifiers init-declarator-list_opt attribute-specifier-list_opt ;
func (p *parser) declaration(ds *DeclarationSpecifiers, d *Declarator) (r *Declaration) {
	if ds == nil {
		panic(todo("", &p.cpp.tok))
	}
	return &Declaration{DeclarationSpecifiers: ds, InitDeclaratorList: p.initDeclaratorListOpt(d), Token: p.must(';')}
}

//  init-declarator-list:
// 	init-declarator
// 	init-declarator-list , attribute-specifier-list_opt init-declarator
func (p *parser) initDeclaratorListOpt(d *Declarator) (r *InitDeclaratorList) {
	switch {
	case d == nil:
		panic(todo("", &p.cpp.tok))
	}
	r = &InitDeclaratorList{InitDeclarator: p.initDeclarator(d)}
	for p.rune() == ',' {
		panic(todo("", &p.cpp.tok))
	}
	return r
}

//  init-declarator:
// 	declarator attribute-specifier-list_opt
// 	declarator attribute-specifier-list_opt = initializer
func (p *parser) initDeclarator(d *Declarator) *InitDeclarator {
	if d == nil {
		panic(todo("", &p.cpp.tok))
	}

	switch p.rune() {
	case '=':
		return &InitDeclarator{Case: InitDeclaratorInit, Token: p.shift(), Initializer: p.initializer()}
	default:
		panic(todo("", &p.cpp.tok))
	}
}

// [0], 6.7.8 Initialization
//
//  initializer:
// 	assignment-expression
// 	{ initializer-list }
// 	{ initializer-list , }
func (p *parser) initializer() *Initializer {
	switch p.rune() {
	case '{':
		panic(todo("", &p.cpp.tok))
	default:
		return &Initializer{Case: InitializerExpr, AssignmentExpression: p.assignmentExpression()}
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
		panic(todo("", lhs, u, &p.cpp.tok))
	case rune(MULASSIGN):
		panic(todo("", lhs, u, &p.cpp.tok))
	case rune(DIVASSIGN):
		panic(todo("", lhs, u, &p.cpp.tok))
	case rune(MODASSIGN):
		panic(todo("", lhs, u, &p.cpp.tok))
	case rune(ADDASSIGN):
		panic(todo("", lhs, u, &p.cpp.tok))
	case rune(SUBASSIGN):
		panic(todo("", lhs, u, &p.cpp.tok))
	case rune(LSHASSIGN):
		panic(todo("", lhs, u, &p.cpp.tok))
	case rune(RSHASSIGN):
		panic(todo("", lhs, u, &p.cpp.tok))
	case rune(ANDASSIGN):
		panic(todo("", lhs, u, &p.cpp.tok))
	case rune(XORASSIGN):
		panic(todo("", lhs, u, &p.cpp.tok))
	case rune(ORASSIGN):
		panic(todo("", lhs, u, &p.cpp.tok))
	default:
		return &AssignmentExpression{Case: AssignmentExpressionCond, ConditionalExpression: lhs}
	}
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
		panic(todo("", lhs, &p.cpp.tok))
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
		panic(todo("", &p.cpp.tok))
	default:
		r = &CastExpression{Case: CastExpressionUnary, UnaryExpression: p.unaryExpression()}
		return r, r.UnaryExpression
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
	case rune(IDENTIFIER):
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
	case rune(IDENTIFIER):
		r = &PostfixExpression{Case: PostfixExpressionPrimary, PrimaryExpression: p.primaryExpression()}
	default:
		panic(todo("", &p.cpp.tok))
	}
	for {
		switch p.rune() {
		case '[':
			panic(todo("", &p.cpp.tok))
		case '(':
			panic(todo("", &p.cpp.tok))
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
	case rune(IDENTIFIER):
		return &PrimaryExpression{Case: PrimaryExpressionIdent, Token: p.shift()}
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
			panic(todo("", r, &p.cpp.tok))
		case '(':
			lparen := p.shift()
			switch p.rune() {
			case rune(IDENTIFIER):
				panic(todo("", r, &p.cpp.tok))
			default:
				r = &DirectDeclarator{Case: DirectDeclaratorFuncParam, DirectDeclarator: r, Token: lparen, ParameterTypeList: p.parameterTypeList(), Token2: p.must(')')}
			}
		default:
			return r
		}
	}
}

//  parameter-type-list:
// 	parameter-list
// 	parameter-list , ...
func (p *parser) parameterTypeList() (r *ParameterTypeList) {
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
		comma := p.shift()
		if p.rune() == rune(DDD) {
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
	case '(':
		return &AbstractDeclarator{Case: AbstractDeclaratorDecl, DirectAbstractDeclarator: p.directAbstractDeclarator()}
	default:
		panic(todo("", &p.cpp.tok))
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
		lparen := p.shift()
		switch p.rune() {
		case ')': // ()
			return &DirectAbstractDeclarator{
				Case:   DirectAbstractDeclaratorFunc,
				Token:  lparen,
				Token2: p.shift(),
			}
		default:
			panic(todo("", &lparen, &p.cpp.tok))
		}
	case '[':
		panic(todo("", &p.cpp.tok))
	default:
		panic(todo("", &p.cpp.tok))
	}
}

//  pointer:
// 	* type-qualifier-list_opt
// 	* type-qualifier-list_opt pointer
//      ^ type-qualifier-list_opt
func (p *parser) pointerOpt() (r *Pointer) {
	switch p.rune() {
	case '*':
		star := p.shift()
		switch p.rune() {
		case ')':
			return &Pointer{Case: PointerTypeQual, Token: star}
		default:
			panic(todo("", &star, &p.cpp.tok))
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
			rune(INT),
			rune(VOID):

			ds = &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: p.typeSpecifier()}
		case rune(IDENTIFIER):
			return r
		case '*', '(':
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
	case rune(INT):
		return &TypeSpecifier{Case: TypeSpecifierInt, Token: p.shift()}
	case rune(VOID):
		return &TypeSpecifier{Case: TypeSpecifierVoid, Token: p.shift()}
	default:
		panic(todo("", &p.cpp.tok))
	}
}
