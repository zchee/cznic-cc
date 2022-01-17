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

func (p *parser) ast() (*AST, error) {
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
	r = &TranslationUnit{ExternalDeclaration: p.externalDeclaration()}
	for prev := r; p.rune() != eof; {
		tu := &TranslationUnit{ExternalDeclaration: p.externalDeclaration()}
		prev.TranslationUnit = tu
		prev = tu
	}
	return r
}

//  external-declaration:
// 	function-definition
// 	declaration
// 	asm-function-definition
// 	;
func (p *parser) externalDeclaration() (r *ExternalDeclaration) {
	switch p.rune() {
	default:
		ds := p.declarationSpecifiers()
		_ = ds
		switch p.rune() {
		case rune(IDENTIFIER):
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
			default:
				panic(todo("", d, &p.cpp.tok))
			}
		default:
			p.rune()
			panic(todo("", &p.cpp.tok))
		}
	}
}

// [0], 6.7.5 Declarators
//
//  declarator:
// 	pointer_opt direct-declarator attribute-specifier-list_opt
func (p *parser) declarator() (r *Declarator) {
	return &Declarator{Pointer: p.pointer(), DirectDeclarator: p.directDeclarator()}
}

//  pointer:
// 	* type-qualifier-list_opt
// 	* type-qualifier-list_opt pointer
//      ^ type-qualifier-list_opt
func (p *parser) pointer() (r *Pointer) {
	if p.rune() != '*' {
		return nil
	}

	r = &Pointer{Case: PointerPtr, Token: p.shift(), TypeQualifiers: p.typeQualifierList()}
	for prev := r; p.rune() == '*'; {
		_ = prev
		panic(todo("", &p.cpp.tok))
	}
	return r
}

//  type-qualifier-list:
// 	type-qualifier
// 	attribute-specifier
// 	type-qualifier-list type-qualifier
// 	type-qualifier-list attribute-specifier
func (p *parser) typeQualifierList() (r *TypeQualifiers) {
	tq := p.typeQualifier()
	if tq == nil {
		return nil
	}

	panic(todo("", &p.cpp.tok))
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
		return nil
	}
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
		p.rune()
		panic(todo("", &p.cpp.tok))
	default:
		p.rune()
		panic(todo("", &p.cpp.tok))
	}
	for prev := r; ; {
		switch p.rune() {
		case '(':
			paren := p.shift()
			switch p.rune() {
			case rune(IDENTIFIER):
				panic(todo("", r, &p.cpp.tok))
			default:
				d := &DirectDeclarator{Case: DirectDeclaratorFuncParam, Token: paren, ParameterTypeList: p.parameterTypeList(), Token2: p.must(')')}
				prev.DirectDeclarator = d
				prev = d
			}
		case ';':
			return r
		default:
			panic(todo("", r, &p.cpp.tok))
		}
	}
}

//  parameter-type-list:
// 	parameter-list
// 	parameter-list , ...
func (p *parser) parameterTypeList() (r *ParameterTypeList) {
	r = &ParameterTypeList{Case: ParameterTypeListList, ParameterList: p.parameterList()}
	if p.rune() == ',' {
		panic(todo("", &p.cpp.tok))
	}
	return r
}

//  parameter-list:
// 	parameter-declaration
// 	parameter-list , parameter-declaration
func (p *parser) parameterList() (r *ParameterList) {
	if p.rune() == ')' {
		return nil
	}

	r = &ParameterList{ParameterDeclaration: p.parameterDeclaration()}
	for prev := r; p.rune() == ','; {
		_ = prev
		panic(todo("", &p.cpp.tok))
	}
	return r
}

//  parameter-declaration:
// 	declaration-specifiers declarator attribute-specifier-list_opt
// 	declaration-specifiers abstract-declarator_opt
func (p *parser) parameterDeclaration() (r *ParameterDeclaration) {
	ds := p.declarationSpecifiers()
	switch p.rune() {
	case ',', ')':
		panic(todo("", ds, &p.cpp.tok))
	}

	switch x := p.declaratorOrAbtractDeclarator().(type) {
	case *AbstractDeclarator:
		return &ParameterDeclaration{Case: ParameterDeclarationAbstract, DeclarationSpecifiers: ds, AbstractDeclarator: x}
	default:
		panic(todo("%v %v %T", ds, &p.cpp.tok, x))
	}
}

func (p *parser) declaratorOrAbtractDeclarator() (r Node) {
	var ptr *Pointer
	switch p.rune() {
	case '*':
		ptr = p.pointer()
		switch p.rune() {
		case ')':
			return &AbstractDeclarator{Case: AbstractDeclaratorPtr, Pointer: ptr}
		default:
			panic(todo("", &p.cpp.tok))
		}
	case '(':
		paren := p.shift()
		p.rune()
		panic(todo("", &paren, &p.cpp.tok))
	default:
		panic(todo("", &p.cpp.tok))
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
	if r = p.declarationSpecifier(); r == nil {
		return nil
	}

	for prev := r; ; {
		ds := p.declarationSpecifier()
		if ds == nil {
			return r
		}

		prev.DeclarationSpecifiers = ds
		prev = ds
	}
}

func (p *parser) declarationSpecifier() (r *DeclarationSpecifiers) {
	switch p.rune() {
	case
		rune(AUTO),
		rune(EXTERN),
		rune(REGISTER),
		rune(STATIC),
		rune(TYPEDEF):

		p.rune()
		panic(todo("", &p.cpp.tok))
	case
		rune(CONST),
		rune(RESTRICT),
		rune(VOLATILE):

		p.rune()
		panic(todo("", &p.cpp.tok))
	case rune(INLINE):
		p.rune()
		panic(todo("", &p.cpp.tok))
	case rune(BOOL):
		return &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: &TypeSpecifier{Case: TypeSpecifierBool, Token: p.shift()}}
	case rune(CHAR):
		return &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: &TypeSpecifier{Case: TypeSpecifierChar, Token: p.shift()}}
	case rune(COMPLEX):
		return &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: &TypeSpecifier{Case: TypeSpecifierComplex, Token: p.shift()}}
	case rune(DOUBLE):
		return &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: &TypeSpecifier{Case: TypeSpecifierDouble, Token: p.shift()}}
	case rune(ENUM):
		p.rune()
		panic(todo("", &p.cpp.tok))
	case rune(FLOAT):
		return &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: &TypeSpecifier{Case: TypeSpecifierFloat, Token: p.shift()}}
	case rune(IMAGINARY):
		p.rune()
		panic(todo("", &p.cpp.tok))
	case rune(INT):
		return &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: &TypeSpecifier{Case: TypeSpecifierInt, Token: p.shift()}}
	case rune(LONG):
		return &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: &TypeSpecifier{Case: TypeSpecifierLong, Token: p.shift()}}
	case rune(SHORT):
		return &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: &TypeSpecifier{Case: TypeSpecifierShort, Token: p.shift()}}
	case rune(SIGNED):
		return &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: &TypeSpecifier{Case: TypeSpecifierSigned, Token: p.shift()}}
	case rune(STRUCT):
		p.rune()
		panic(todo("", &p.cpp.tok))
	case rune(TYPENAME):
		p.rune()
		panic(todo("", &p.cpp.tok))
	case rune(UNSIGNED):
		return &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: &TypeSpecifier{Case: TypeSpecifierUnsigned, Token: p.shift()}}
	case rune(VOID):
		return &DeclarationSpecifiers{Case: DeclarationSpecifiersTypeSpec, TypeSpecifier: &TypeSpecifier{Case: TypeSpecifierVoid, Token: p.shift()}}
	}
	return nil
}
