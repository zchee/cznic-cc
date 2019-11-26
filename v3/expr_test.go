// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"reflect"
	"strconv"
	"testing"
)

func eqExpr(e1, e2 Expr) bool {
	return reflect.DeepEqual(e1, e2)
}

func ident(s string) *Ident {
	return &Ident{Name: s}
}

func identExpr(s string) IdentExpr {
	return IdentExpr{ident(s)}
}

func intLit(v int) Expr {
	return &Literal{Kind: LiteralInt, Value: strconv.Itoa(v)}
}

func TestExpr(t *testing.T) {
	cfg := &Config{ABI: testABI}
	for _, v := range []struct {
		name string
		src  string
		exp  Expr
	}{
		{
			name: "int lit",
			src:  `int f() { 42; }`,
			exp:  &Literal{Kind: LiteralInt, Value: `42`},
		},
		{
			name: "int lit hex",
			src:  `int f() { 0x42; }`,
			exp:  &Literal{Kind: LiteralInt, Value: `0x42`},
		},
		{
			name: "int lit suffix",
			src:  `int f() { 42u; }`,
			exp:  &Literal{Kind: LiteralInt, Value: `42u`},
		},
		{
			name: "int lit pos",
			src:  `int f() { +42; }`,
			exp:  &UnaryExpr{Op: UnaryPlus, X: intLit(42)},
		},
		{
			name: "int lit neg",
			src:  `int f() { -42; }`,
			exp:  &UnaryExpr{Op: UnaryMinus, X: intLit(42)},
		},
		{
			name: "float lit",
			src:  `int f() { 1.0; }`,
			exp:  &Literal{Kind: LiteralFloat, Value: `1.0`},
		},
		{
			name: "float lit exp",
			src:  `int f() { 1e5; }`,
			exp:  &Literal{Kind: LiteralFloat, Value: `1e5`},
		},
		{
			name: "float lit pos",
			src:  `int f() { +1.0; }`,
			exp:  &UnaryExpr{Op: UnaryPlus, X: &Literal{Kind: LiteralFloat, Value: `1.0`}},
		},
		{
			name: "float lit neg",
			src:  `int f() { -1.0; }`,
			exp:  &UnaryExpr{Op: UnaryMinus, X: &Literal{Kind: LiteralFloat, Value: `1.0`}},
		},
		{
			name: "char lit",
			src:  `int f() { 'a'; }`,
			exp:  &Literal{Kind: LiteralChar, Value: `a`},
		},
		{
			name: "wide char lit",
			src:  `int f() { L'a'; }`,
			exp:  &Literal{Kind: LiteralWChar, Value: `a`},
		},
		{
			name: "string lit",
			src:  `int f() { "a"; }`,
			exp:  &Literal{Kind: LiteralString, Value: `a`},
		},
		{
			name: "wide string lit",
			src:  `int f() { L"a"; }`,
			exp:  &Literal{Kind: LiteralWString, Value: `a`},
		},
		{
			name: "parentheses",
			src:  `int f() { (1); }`,
			exp:  &ParenExpr{X: intLit(1)},
		},
		{
			name: "comma",
			src:  `int f() { 1,2,3; }`,
			exp:  CommaExpr{intLit(1), intLit(2), intLit(3)},
		},
		{
			name: "ident",
			src:  `int f(int x) { x; }`,
			exp:  identExpr("x"),
		},
		{
			name: "unary plus",
			src:  `int f(int x) { +x; }`,
			exp:  &UnaryExpr{Op: UnaryPlus, X: identExpr("x")},
		},
		{
			name: "unary minus",
			src:  `int f(int x) { -x; }`,
			exp:  &UnaryExpr{Op: UnaryMinus, X: identExpr("x")},
		},
		{
			name: "unary addr",
			src:  `int f(int x) { &x; }`,
			exp:  &UnaryExpr{Op: UnaryAddr, X: identExpr("x")},
		},
		{
			name: "unary deref",
			src:  `int f(int x) { *x; }`,
			exp:  &UnaryExpr{Op: UnaryDeref, X: identExpr("x")},
		},
		{
			name: "unary not",
			src:  `int f(int x) { !x; }`,
			exp:  &UnaryExpr{Op: UnaryNot, X: identExpr("x")},
		},
		{
			name: "unary invert",
			src:  `int f(int x) { ~x; }`,
			exp:  &UnaryExpr{Op: UnaryInvert, X: identExpr("x")},
		},
		{
			name: "binary add",
			src:  `int f(int x, int y) { x + y; }`,
			exp:  &BinaryExpr{X: identExpr("x"), Op: BinaryAdd, Y: identExpr("y")},
		},
		{
			name: "binary order 1",
			src:  `int f(int x, int y) { x * 2 + y; }`,
			exp:  &BinaryExpr{X: &BinaryExpr{X: identExpr("x"), Op: BinaryMul, Y: intLit(2)}, Op: BinaryAdd, Y: identExpr("y")},
		},
		{
			name: "binary order 2",
			src:  `int f(int x, int y) { x + y * 2; }`,
			exp:  &BinaryExpr{X: identExpr("x"), Op: BinaryAdd, Y: &BinaryExpr{X: identExpr("y"), Op: BinaryMul, Y: intLit(2)}},
		},
		{
			name: "assign",
			src:  `int f(int x) { x = 1; }`,
			exp:  &AssignExpr{Left: identExpr("x"), Op: BinaryNone, Right: intLit(1)},
		},
		{
			name: "assign op",
			src:  `int f(int x) { x += 1; }`,
			exp:  &AssignExpr{Left: identExpr("x"), Op: BinaryAdd, Right: intLit(1)},
		},
		{
			name: "assign chain",
			src:  `int f(int x, int y) { y = x = 1; }`,
			exp:  &AssignExpr{Left: identExpr("y"), Right: &AssignExpr{Left: identExpr("x"), Right: intLit(1)}},
		},
		{
			name: "inc post",
			src:  `int f(int x) { x++; }`,
			exp:  &IncDecExpr{X: identExpr("x"), Op: IncPost},
		},
		{
			name: "inc pre",
			src:  `int f(int x) { ++x; }`,
			exp:  &IncDecExpr{X: identExpr("x"), Op: IncPre},
		},
		{
			name: "inc add 1",
			src:  `int f(int x, int y) { x+++y; }`,
			exp:  &BinaryExpr{X: &IncDecExpr{X: identExpr("x"), Op: IncPost}, Op: BinaryAdd, Y: identExpr("y")},
		},
		{
			name: "inc add 2",
			src:  `int f(int x, int y) { ++x+y; }`,
			exp:  &BinaryExpr{X: &IncDecExpr{X: identExpr("x"), Op: IncPre}, Op: BinaryAdd, Y: identExpr("y")},
		},
		{
			name: "inc add 3",
			src:  `int f(int x, int y) { x+y++; }`,
			exp:  &BinaryExpr{X: identExpr("x"), Op: BinaryAdd, Y: &IncDecExpr{X: identExpr("y"), Op: IncPost}},
		},
		{
			name: "index",
			src:  `int f(int* x) { x[1]; }`,
			exp:  &IndexExpr{X: identExpr("x"), Ind: intLit(1)},
		},
		{
			name: "index chain",
			src:  `int f(int* x) { x[1][2]; }`,
			exp:  &IndexExpr{X: &IndexExpr{X: identExpr("x"), Ind: intLit(1)}, Ind: intLit(2)},
		},
		{
			name: "select",
			src:  `struct s { int y; }; int f(struct s x) { x.y; }`,
			exp:  &SelectExpr{X: identExpr("x"), Sel: ident("y"), Ptr: false},
		},
		{
			name: "select ptr",
			src:  `struct s { int y; }; int f(struct s* x) { x->y; }`,
			exp:  &SelectExpr{X: identExpr("x"), Sel: ident("y"), Ptr: true},
		},
		{
			name: "select chain 1",
			src:  `struct s1 { int z; }; struct s2 { struct s1 y; }; int f(struct s2 x) { x.y.z; }`,
			exp:  &SelectExpr{X: &SelectExpr{X: identExpr("x"), Sel: ident("y")}, Sel: ident("z")},
		},
		{
			name: "select chain 2",
			src:  `struct s1 { int z; }; struct s2 { struct s1* y; }; int f(struct s2* x) { x->y->z; }`,
			exp:  &SelectExpr{X: &SelectExpr{X: identExpr("x"), Sel: ident("y"), Ptr: true}, Sel: ident("z"), Ptr: true},
		},
		{
			name: "select chain 3",
			src:  `struct s1 { int z; }; struct s2 { struct s1 y; }; int f(struct s2* x) { x.y->z; }`,
			exp:  &SelectExpr{X: &SelectExpr{X: identExpr("x"), Sel: ident("y")}, Sel: ident("z"), Ptr: true},
		},
		{
			name: "select chain 4",
			src:  `struct s1 { int z; }; struct s2 { struct s1* y; }; int f(struct s2 x) { x->y.z; }`,
			exp:  &SelectExpr{X: &SelectExpr{X: identExpr("x"), Sel: ident("y"), Ptr: true}, Sel: ident("z")},
		},
		{
			name: "call",
			src:  `void x(int a1, int a2); int f() { x(1,2); }`,
			exp:  &CallExpr{Func: identExpr("x"), Args: []Expr{intLit(1), intLit(2)}},
		},
		{
			name: "cond",
			src:  `int f(int x) { x ? 1 : 2; }`,
			exp:  &CondExpr{Cond: identExpr("x"), Then: intLit(1), Else: intLit(2)},
		},
		{
			name: "cond chain",
			src:  `int f(int x, int y) { x ? 1 : y ? 2 : 3; }`,
			exp:  &CondExpr{Cond: identExpr("x"), Then: intLit(1), Else: &CondExpr{Cond: identExpr("y"), Then: intLit(2), Else: intLit(3)}},
		},
	} {
		t.Run(v.name, func(t *testing.T) {
			ast, err := Parse(cfg, nil, nil, []Source{{Name: "test", Value: v.src}})
			if err != nil {
				t.Fatal(err)
				return
			}
			tu := ast.TranslationUnit
			for ; tu.TranslationUnit != nil; tu = tu.TranslationUnit {
			}
			e := tu.
				ExternalDeclaration.
				FunctionDefinition.
				CompoundStatement.
				BlockItemList.
				BlockItemList.
				BlockItem.
				Statement.
				ExpressionStatement.
				Expression.
				Expr()
			if !eqExpr(v.exp, e) {
				t.Fatalf("unexpected expression: %#v", e)
			}
		})
	}
}
