// Copyright 2019 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"strconv"
	"testing"
)

// eqExpr implements a limited deep equality for expressions.
// Namely, it ignores types and identifier matching (only matches by the name).
func eqExpr(e1, e2 Expr) bool {
	switch e1 := e1.(type) {
	case IdentExpr:
		e2, ok := e2.(IdentExpr)
		return ok && e1.Ident.Name == e2.Ident.Name
	case *Literal:
		e2, ok := e2.(*Literal)
		return ok && e1.kind == e2.kind && e1.value == e2.value
	case *ParenExpr:
		e2, ok := e2.(*ParenExpr)
		return ok && eqExpr(e1.X, e2.X)
	case *UnaryExpr:
		e2, ok := e2.(*UnaryExpr)
		return ok && e1.Op == e2.Op && eqExpr(e1.X, e2.X)
	case *IncDecExpr:
		e2, ok := e2.(*IncDecExpr)
		return ok && e1.Op == e2.Op && eqExpr(e1.X, e2.X)
	case *BinaryExpr:
		e2, ok := e2.(*BinaryExpr)
		return ok && e1.Op == e2.Op && eqExpr(e1.X, e2.X) && eqExpr(e1.Y, e2.Y)
	case *AssignExpr:
		e2, ok := e2.(*AssignExpr)
		return ok && e1.Op == e2.Op && eqExpr(e1.Left, e2.Left) && eqExpr(e1.Right, e2.Right)
	case *IndexExpr:
		e2, ok := e2.(*IndexExpr)
		return ok && eqExpr(e1.X, e2.X) && eqExpr(e1.Ind, e2.Ind)
	case *SelectExpr:
		e2, ok := e2.(*SelectExpr)
		return ok && e1.Ptr == e2.Ptr && eqExpr(e1.X, e2.X) && eqExpr(IdentExpr{e1.Sel}, IdentExpr{e2.Sel})
	case *CondExpr:
		e2, ok := e2.(*CondExpr)
		return ok && eqExpr(e1.Cond, e2.Cond) && eqExpr(e1.Then, e2.Then) && eqExpr(e1.Else, e2.Else)
	case *CallExpr:
		e2, ok := e2.(*CallExpr)
		if !ok || len(e1.Args) != len(e2.Args) || !eqExpr(e1.Func, e2.Func) {
			return false
		}
		for i := range e1.Args {
			if !eqExpr(e1.Args[i], e2.Args[i]) {
				return false
			}
		}
		return true
	case CommaExpr:
		e2, ok := e2.(CommaExpr)
		if !ok || len(e1) != len(e2) {
			return false
		}
		for i := range e1 {
			if !eqExpr(e1[i], e2[i]) {
				return false
			}
		}
		return true
	}
	panic(e1)
}

func ident(s string) *Ident {
	return &Ident{Name: s}
}

func identExpr(s string) IdentExpr {
	return IdentExpr{ident(s)}
}

func lit(kind LiteralKind, v string) Expr {
	return &Literal{kind: kind, value: v}
}

func intLit(v int) Expr {
	return lit(LiteralInt, strconv.Itoa(v))
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
			exp:  lit(LiteralInt, `42`),
		},
		{
			name: "int lit hex",
			src:  `int f() { 0x42; }`,
			exp:  lit(LiteralInt, `0x42`),
		},
		{
			name: "int lit suffix",
			src:  `int f() { 42u; }`,
			exp:  lit(LiteralInt, `42u`),
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
			exp:  lit(LiteralFloat, `1.0`),
		},
		{
			name: "float lit exp",
			src:  `int f() { 1e5; }`,
			exp:  lit(LiteralFloat, `1e5`),
		},
		{
			name: "float lit pos",
			src:  `int f() { +1.0; }`,
			exp:  &UnaryExpr{Op: UnaryPlus, X: lit(LiteralFloat, `1.0`)},
		},
		{
			name: "float lit neg",
			src:  `int f() { -1.0; }`,
			exp:  &UnaryExpr{Op: UnaryMinus, X: lit(LiteralFloat, `1.0`)},
		},
		{
			name: "char lit",
			src:  `int f() { 'a'; }`,
			exp:  lit(LiteralChar, `a`),
		},
		{
			name: "wide char lit",
			src:  `int f() { L'a'; }`,
			exp:  lit(LiteralWChar, `a`),
		},
		{
			name: "string lit",
			src:  `int f() { "a"; }`,
			exp:  lit(LiteralString, `a`),
		},
		{
			name: "wide string lit",
			src:  `int f() { L"a"; }`,
			exp:  lit(LiteralWString, `a`),
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
