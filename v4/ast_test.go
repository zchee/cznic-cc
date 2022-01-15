// Code generated by yy. DO NOT EDIT.

// Copyright 2022 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"fmt"
)

func ExampleAbstractDeclarator_ptr() {
	fmt.Println(exampleAST(193, "void f(int*);"))
	// Output:
	// TODO
}

func ExampleAbstractDeclarator_decl() {
	fmt.Println(exampleAST(194, "void f(int());"))
	// Output:
	// TODO
}

func ExampleAdditiveExpression_mul() {
	fmt.Println(exampleAST(45, "int i = x;"))
	// Output:
	// TODO
}

func ExampleAdditiveExpression_add() {
	fmt.Println(exampleAST(46, "int i = x+y;"))
	// Output:
	// TODO
}

func ExampleAdditiveExpression_sub() {
	fmt.Println(exampleAST(47, "int i = x-y;"))
	// Output:
	// TODO
}

func ExampleAlignmentSpecifier_alignasType() {
	fmt.Println(exampleAST(167, "_Alignas(double) char c;"))
	// Output:
	// TODO
}

func ExampleAlignmentSpecifier_alignasExpr() {
	fmt.Println(exampleAST(168, "_Alignas(0ll) char c;"))
	// Output:
	// TODO
}

func ExampleAndExpression_eq() {
	fmt.Println(exampleAST(59, "int i = x;"))
	// Output:
	// TODO
}

func ExampleAndExpression_and() {
	fmt.Println(exampleAST(60, "int i = x & y;"))
	// Output:
	// TODO
}

func ExampleArgumentExpressionList_case0() {
	fmt.Println(exampleAST(21, "int i = f(x);"))
	// Output:
	// TODO
}

func ExampleArgumentExpressionList_case1() {
	fmt.Println(exampleAST(22, "int i = f(x, y);"))
	// Output:
	// TODO
}

func ExampleAsm_case0() {
	fmt.Println(exampleAST(259, "__asm__(\"nop\");"))
	// Output:
	// TODO
}

func ExampleAsmArgList_case0() {
	fmt.Println(exampleAST(257, "__asm__(\"nop\": a);"))
	// Output:
	// TODO
}

func ExampleAsmArgList_case1() {
	fmt.Println(exampleAST(258, "__asm__(\"nop\": a : b);"))
	// Output:
	// TODO
}

func ExampleAsmExpressionList_case0() {
	fmt.Println(exampleAST(255, "__asm__(\"nop\": a);"))
	// Output:
	// TODO
}

func ExampleAsmExpressionList_case1() {
	fmt.Println(exampleAST(256, "__asm__(\"nop\": a, b);"))
	// Output:
	// TODO
}

func ExampleAsmFunctionDefinition_case0() {
	fmt.Println(exampleAST(261, "int f() __asm__(\"nop\");"))
	// Output:
	// TODO
}

func ExampleAsmIndex_case0() {
	fmt.Println(exampleAST(254, "__asm__(\"nop\": [a] b);"))
	// Output:
	// TODO
}

func ExampleAsmQualifier_volatile() {
	fmt.Println(exampleAST(262, "__asm__ volatile (\"nop\");"))
	// Output:
	// TODO
}

func ExampleAsmQualifier_inline() {
	fmt.Println(exampleAST(263, "__asm__ inline (\"nop\");"))
	// Output:
	// TODO
}

func ExampleAsmQualifier_goto() {
	fmt.Println(exampleAST(264, "__asm__ goto (\"nop\");"))
	// Output:
	// TODO
}

func ExampleAsmQualifierList_case0() {
	fmt.Println(exampleAST(265, "__asm__ inline (\"nop\");"))
	// Output:
	// TODO
}

func ExampleAsmQualifierList_case1() {
	fmt.Println(exampleAST(266, "__asm__ inline volatile (\"nop\");"))
	// Output:
	// TODO
}

func ExampleAsmStatement_case0() {
	fmt.Println(exampleAST(260, "void f() { __asm__(\"nop\"); }"))
	// Output:
	// TODO
}

func ExampleAssignmentExpression_cond() {
	fmt.Println(exampleAST(71, "int i = x; }"))
	// Output:
	// TODO
}

func ExampleAssignmentExpression_assign() {
	fmt.Println(exampleAST(72, "int f() { x = y; }"))
	// Output:
	// TODO
}

func ExampleAssignmentExpression_mul() {
	fmt.Println(exampleAST(73, "int f() { x *= y; }"))
	// Output:
	// TODO
}

func ExampleAssignmentExpression_div() {
	fmt.Println(exampleAST(74, "int f() { x /= y; }"))
	// Output:
	// TODO
}

func ExampleAssignmentExpression_mod() {
	fmt.Println(exampleAST(75, "int f() { x %= y; }"))
	// Output:
	// TODO
}

func ExampleAssignmentExpression_add() {
	fmt.Println(exampleAST(76, "int f() { x += y; }"))
	// Output:
	// TODO
}

func ExampleAssignmentExpression_sub() {
	fmt.Println(exampleAST(77, "int f() { x -= y; }"))
	// Output:
	// TODO
}

func ExampleAssignmentExpression_lsh() {
	fmt.Println(exampleAST(78, "int f() { x <<= y; }"))
	// Output:
	// TODO
}

func ExampleAssignmentExpression_rsh() {
	fmt.Println(exampleAST(79, "int f() { x >>= y; }"))
	// Output:
	// TODO
}

func ExampleAssignmentExpression_and() {
	fmt.Println(exampleAST(80, "int f() { x &= y; }"))
	// Output:
	// TODO
}

func ExampleAssignmentExpression_xor() {
	fmt.Println(exampleAST(81, "int f() { x ^= y; }"))
	// Output:
	// TODO
}

func ExampleAssignmentExpression_or() {
	fmt.Println(exampleAST(82, "int f() { x |= y; }"))
	// Output:
	// TODO
}

func ExampleAtomicTypeSpecifier_case0() {
	fmt.Println(exampleAST(159, "_Atomic(int) i;"))
	// Output:
	// TODO
}

func ExampleAttributeSpecifier_case0() {
	fmt.Println(exampleAST(274, "int i __attribute__((a));"))
	// Output:
	// TODO
}

func ExampleAttributeSpecifierList_case0() {
	fmt.Println(exampleAST(275, "int i __attribute__((a));"))
	// Output:
	// TODO
}

func ExampleAttributeSpecifierList_case1() {
	fmt.Println(exampleAST(276, "int i __attribute__((a)) __attribute((b));"))
	// Output:
	// TODO
}

func ExampleAttributeValue_ident() {
	fmt.Println(exampleAST(270, "int i __attribute__((a));"))
	// Output:
	// TODO
}

func ExampleAttributeValue_expr() {
	fmt.Println(exampleAST(271, "int i __attribute__((a(b)));"))
	// Output:
	// TODO
}

func ExampleAttributeValueList_case0() {
	fmt.Println(exampleAST(272, "int i __attribute__((a));"))
	// Output:
	// TODO
}

func ExampleAttributeValueList_case1() {
	fmt.Println(exampleAST(273, "int i __attribute__((a, b));"))
	// Output:
	// TODO
}

func ExampleBlockItem_decl() {
	fmt.Println(exampleAST(225, "int f() { int i; }"))
	// Output:
	// TODO
}

func ExampleBlockItem_stmt() {
	fmt.Println(exampleAST(226, "int f() { g(); }"))
	// Output:
	// TODO
}

func ExampleBlockItem_label() {
	fmt.Println(exampleAST(227, "int f() { __label__ L; }"))
	// Output:
	// TODO
}

func ExampleBlockItem_funcDef() {
	fmt.Println(exampleAST(228, "int f() { int g() {} }"))
	// Output:
	// TODO
}

func ExampleBlockItem_pragma() {
	fmt.Println(exampleAST(229, "int f() {\\n#pragma STDC FENV_ACCESS OFF\\n}"))
	// Output:
	// TODO
}

func ExampleBlockItemList_case0() {
	fmt.Println(exampleAST(223, "int f() { int i; }"))
	// Output:
	// TODO
}

func ExampleBlockItemList_case1() {
	fmt.Println(exampleAST(224, "int f() { int i; double j; }"))
	// Output:
	// TODO
}

func ExampleCastExpression_unary() {
	fmt.Println(exampleAST(39, "int i = 42;"))
	// Output:
	// TODO
}

func ExampleCastExpression_cast() {
	fmt.Println(exampleAST(40, "int i = (int)3.14;"))
	// Output:
	// TODO
}

func ExampleCompoundStatement_case0() {
	fmt.Println(exampleAST(222, "int f() { int i; }"))
	// Output:
	// TODO
}

func ExampleConditionalExpression_lOr() {
	fmt.Println(exampleAST(69, "int i = x;"))
	// Output:
	// TODO
}

func ExampleConditionalExpression_cond() {
	fmt.Println(exampleAST(70, "int i = x ? y : z;"))
	// Output:
	// TODO
}

func ExampleConstantExpression_case0() {
	fmt.Println(exampleAST(85, "struct { int i:3; };"))
	// Output:
	// TODO
}

func ExampleDeclaration_case0() {
	fmt.Println(exampleAST(86, "int i, j;"))
	// Output:
	// TODO
}

func ExampleDeclarationList_case0() {
	fmt.Println(exampleAST(252, "int f(i) int i; {}"))
	// Output:
	// TODO
}

func ExampleDeclarationList_case1() {
	fmt.Println(exampleAST(253, "int f(i, j) int i; int j; {}"))
	// Output:
	// TODO
}

func ExampleDeclarationSpecifiers_storage() {
	fmt.Println(exampleAST(87, "static int i;"))
	// Output:
	// TODO
}

func ExampleDeclarationSpecifiers_typeSpec() {
	fmt.Println(exampleAST(88, "int i;"))
	// Output:
	// TODO
}

func ExampleDeclarationSpecifiers_typeQual() {
	fmt.Println(exampleAST(89, "volatile int i;"))
	// Output:
	// TODO
}

func ExampleDeclarationSpecifiers_func() {
	fmt.Println(exampleAST(90, "inline int f() {}"))
	// Output:
	// TODO
}

func ExampleDeclarationSpecifiers_alignSpec() {
	fmt.Println(exampleAST(91, "_Alignas(double) int i;"))
	// Output:
	// TODO
}

func ExampleDeclarationSpecifiers_attribute() {
	fmt.Println(exampleAST(92, "int __attribute__((a)) i;"))
	// Output:
	// TODO
}

func ExampleDeclarator_case0() {
	fmt.Println(exampleAST(166, "int *p __attribute__ ((foo));"))
	// Output:
	// TODO
}

func ExampleDesignation_case0() {
	fmt.Println(exampleAST(205, "int a[] = { [42] = 314 };"))
	// Output:
	// TODO
}

func ExampleDesignator_index() {
	fmt.Println(exampleAST(208, "int a[] = { [42] = 314 };"))
	// Output:
	// TODO
}

func ExampleDesignator_field() {
	fmt.Println(exampleAST(209, "struct t s = { .fld = 314 };"))
	// Output:
	// TODO
}

func ExampleDesignator_field2() {
	fmt.Println(exampleAST(210, "struct t s = { fld: 314 };"))
	// Output:
	// TODO
}

func ExampleDesignatorList_case0() {
	fmt.Println(exampleAST(206, "int a[] = { [42] = 314 };"))
	// Output:
	// TODO
}

func ExampleDesignatorList_case1() {
	fmt.Println(exampleAST(207, "int a[100][] = { [42][12] = 314 };"))
	// Output:
	// TODO
}

func ExampleDirectAbstractDeclarator_decl() {
	fmt.Println(exampleAST(195, "void f(int());"))
	// Output:
	// TODO
}

func ExampleDirectAbstractDeclarator_arr() {
	fmt.Println(exampleAST(196, "void f(int[const 42]);"))
	// Output:
	// TODO
}

func ExampleDirectAbstractDeclarator_staticArr() {
	fmt.Println(exampleAST(197, "void f(int[static const 42]);"))
	// Output:
	// TODO
}

func ExampleDirectAbstractDeclarator_arrStatic() {
	fmt.Println(exampleAST(198, "void f(int[const static 42]);"))
	// Output:
	// TODO
}

func ExampleDirectAbstractDeclarator_arrStar() {
	fmt.Println(exampleAST(199, "void f(int[*]);"))
	// Output:
	// TODO
}

func ExampleDirectAbstractDeclarator_func() {
	fmt.Println(exampleAST(200, "void f(int(char));"))
	// Output:
	// TODO
}

func ExampleDirectDeclarator_ident() {
	fmt.Println(exampleAST(169, "int i;"))
	// Output:
	// TODO
}

func ExampleDirectDeclarator_decl() {
	fmt.Println(exampleAST(170, "int (f);"))
	// Output:
	// TODO
}

func ExampleDirectDeclarator_arr() {
	fmt.Println(exampleAST(171, "int i[const 42];"))
	// Output:
	// TODO
}

func ExampleDirectDeclarator_staticArr() {
	fmt.Println(exampleAST(172, "int i[static const 42];"))
	// Output:
	// TODO
}

func ExampleDirectDeclarator_arrStatic() {
	fmt.Println(exampleAST(173, "int i[const static 42];"))
	// Output:
	// TODO
}

func ExampleDirectDeclarator_star() {
	fmt.Println(exampleAST(174, "int i[const *];"))
	// Output:
	// TODO
}

func ExampleDirectDeclarator_funcParam() {
	fmt.Println(exampleAST(175, "int f(int i);"))
	// Output:
	// TODO
}

func ExampleDirectDeclarator_funcIdent() {
	fmt.Println(exampleAST(176, "int f(a);"))
	// Output:
	// TODO
}

func ExampleEnumSpecifier_def() {
	fmt.Println(exampleAST(153, "enum e {a};"))
	// Output:
	// TODO
}

func ExampleEnumSpecifier_tag() {
	fmt.Println(exampleAST(154, "enum e i;"))
	// Output:
	// TODO
}

func ExampleEnumerator_ident() {
	fmt.Println(exampleAST(157, "enum e {a};"))
	// Output:
	// TODO
}

func ExampleEnumerator_expr() {
	fmt.Println(exampleAST(158, "enum e {a = 42};"))
	// Output:
	// TODO
}

func ExampleEnumeratorList_case0() {
	fmt.Println(exampleAST(155, "enum e {a};"))
	// Output:
	// TODO
}

func ExampleEnumeratorList_case1() {
	fmt.Println(exampleAST(156, "enum e {a, b};"))
	// Output:
	// TODO
}

func ExampleEqualityExpression_rel() {
	fmt.Println(exampleAST(56, "int i = x;"))
	// Output:
	// TODO
}

func ExampleEqualityExpression_eq() {
	fmt.Println(exampleAST(57, "int i = x == y;"))
	// Output:
	// TODO
}

func ExampleEqualityExpression_neq() {
	fmt.Println(exampleAST(58, "int i = x != y;"))
	// Output:
	// TODO
}

func ExampleExclusiveOrExpression_and() {
	fmt.Println(exampleAST(61, "int i = x;"))
	// Output:
	// TODO
}

func ExampleExclusiveOrExpression_xor() {
	fmt.Println(exampleAST(62, "int i = x^y;"))
	// Output:
	// TODO
}

func ExampleExpression_assign() {
	fmt.Println(exampleAST(83, "int f() { i = x; };"))
	// Output:
	// TODO
}

func ExampleExpression_comma() {
	fmt.Println(exampleAST(84, "int f() { x, y; };"))
	// Output:
	// TODO
}

func ExampleExpressionList_case0() {
	fmt.Println(exampleAST(268, "int i __attribute__((a(b)));"))
	// Output:
	// TODO
}

func ExampleExpressionList_case1() {
	fmt.Println(exampleAST(269, "int i __attribute__((a(b, c)));"))
	// Output:
	// TODO
}

func ExampleExpressionStatement_case0() {
	fmt.Println(exampleAST(230, "int f() { g(); }"))
	// Output:
	// TODO
}

func ExampleExternalDeclaration_funcDef() {
	fmt.Println(exampleAST(245, "int f() {}"))
	// Output:
	// TODO
}

func ExampleExternalDeclaration_decl() {
	fmt.Println(exampleAST(246, "int i;"))
	// Output:
	// TODO
}

func ExampleExternalDeclaration_asm() {
	fmt.Println(exampleAST(247, "int f() __asm__(\"nop\");"))
	// Output:
	// TODO
}

func ExampleExternalDeclaration_asmStmt() {
	fmt.Println(exampleAST(248, "__asm__(\"nop\");"))
	// Output:
	// TODO
}

func ExampleExternalDeclaration_empty() {
	fmt.Println(exampleAST(249, ";"))
	// Output:
	// TODO
}

func ExampleExternalDeclaration_pragma() {
	fmt.Println(exampleAST(250, "#pragma STDC CX_LIMITED_RANGE DEFAULT"))
	// Output:
	// TODO
}

func ExampleFunctionDefinition_case0() {
	fmt.Println(exampleAST(251, "int f() {}"))
	// Output:
	// TODO
}

func ExampleFunctionSpecifier_inline() {
	fmt.Println(exampleAST(164, "inline int f() {}"))
	// Output:
	// TODO
}

func ExampleFunctionSpecifier_noreturn() {
	fmt.Println(exampleAST(165, "_Noreturn int f() {}"))
	// Output:
	// TODO
}

func ExampleIdentifierList_case0() {
	fmt.Println(exampleAST(190, "int f(i) int i; {}"))
	// Output:
	// TODO
}

func ExampleIdentifierList_case1() {
	fmt.Println(exampleAST(191, "int f(i, j) int i, j; {}"))
	// Output:
	// TODO
}

func ExampleInclusiveOrExpression_xor() {
	fmt.Println(exampleAST(63, "int i = x;"))
	// Output:
	// TODO
}

func ExampleInclusiveOrExpression_or() {
	fmt.Println(exampleAST(64, "int i = x|y;"))
	// Output:
	// TODO
}

func ExampleInitDeclarator_decl() {
	fmt.Println(exampleAST(95, "int i;"))
	// Output:
	// TODO
}

func ExampleInitDeclarator_init() {
	fmt.Println(exampleAST(96, "int i = x;"))
	// Output:
	// TODO
}

func ExampleInitDeclaratorList_case0() {
	fmt.Println(exampleAST(93, "int i;"))
	// Output:
	// TODO
}

func ExampleInitDeclaratorList_case1() {
	fmt.Println(exampleAST(94, "int i, j;"))
	// Output:
	// TODO
}

func ExampleInitializer_expr() {
	fmt.Println(exampleAST(201, "int i = x;"))
	// Output:
	// TODO
}

func ExampleInitializer_initList() {
	fmt.Println(exampleAST(202, "int i[] = { x };"))
	// Output:
	// TODO
}

func ExampleInitializerList_case0() {
	fmt.Println(exampleAST(203, "int i[] = { [10] = x };"))
	// Output:
	// TODO
}

func ExampleInitializerList_case1() {
	fmt.Println(exampleAST(204, "int i[] = { [10] = x, [20] = y };"))
	// Output:
	// TODO
}

func ExampleIterationStatement_while() {
	fmt.Println(exampleAST(234, "int f() { while(x) y(); }"))
	// Output:
	// TODO
}

func ExampleIterationStatement_do() {
	fmt.Println(exampleAST(235, "int f() { do x(); while(y); }"))
	// Output:
	// TODO
}

func ExampleIterationStatement_for() {
	fmt.Println(exampleAST(236, "int f() { for( i = 0; i < 10; i++) x(); }"))
	// Output:
	// TODO
}

func ExampleIterationStatement_forDecl() {
	fmt.Println(exampleAST(237, "int f() { for( int i = 0; i < 10; i++) x(); }"))
	// Output:
	// TODO
}

func ExampleJumpStatement_goto() {
	fmt.Println(exampleAST(238, "int f() { L: goto L; }"))
	// Output:
	// TODO
}

func ExampleJumpStatement_gotoExpr() {
	fmt.Println(exampleAST(239, "int f() { L: x(); void *p = &&L; goto *p; }"))
	// Output:
	// TODO
}

func ExampleJumpStatement_continue() {
	fmt.Println(exampleAST(240, "int f() { for(;;) if (i) continue; }"))
	// Output:
	// TODO
}

func ExampleJumpStatement_break() {
	fmt.Println(exampleAST(241, "int f() { for(;;) if (i) break; }"))
	// Output:
	// TODO
}

func ExampleJumpStatement_return() {
	fmt.Println(exampleAST(242, "int f() { if (i) return x; }"))
	// Output:
	// TODO
}

func ExampleLabelDeclaration_case0() {
	fmt.Println(exampleAST(267, "int f() { __label__ L; L: x(); }"))
	// Output:
	// TODO
}

func ExampleLabeledStatement_label() {
	fmt.Println(exampleAST(218, "int f() { L: goto L; }"))
	// Output:
	// TODO
}

func ExampleLabeledStatement_caseLabel() {
	fmt.Println(exampleAST(219, "int f() { switch(i) case 42: x(); }"))
	// Output:
	// TODO
}

func ExampleLabeledStatement_range() {
	fmt.Println(exampleAST(220, "int f() { switch(i) case 42 ... 56: x(); }"))
	// Output:
	// TODO
}

func ExampleLabeledStatement_default() {
	fmt.Println(exampleAST(221, "int f() { switch(i) default: x(); }"))
	// Output:
	// TODO
}

func ExampleLogicalAndExpression_or() {
	fmt.Println(exampleAST(65, "int i = x;"))
	// Output:
	// TODO
}

func ExampleLogicalAndExpression_lAnd() {
	fmt.Println(exampleAST(66, "int i = x && y;"))
	// Output:
	// TODO
}

func ExampleLogicalOrExpression_lAnd() {
	fmt.Println(exampleAST(67, "int i = x;"))
	// Output:
	// TODO
}

func ExampleLogicalOrExpression_lOr() {
	fmt.Println(exampleAST(68, "int i = x || y;"))
	// Output:
	// TODO
}

func ExampleMultiplicativeExpression_cast() {
	fmt.Println(exampleAST(41, "int i = x;"))
	// Output:
	// TODO
}

func ExampleMultiplicativeExpression_mul() {
	fmt.Println(exampleAST(42, "int i = x * y;"))
	// Output:
	// TODO
}

func ExampleMultiplicativeExpression_div() {
	fmt.Println(exampleAST(43, "int i = x / y;"))
	// Output:
	// TODO
}

func ExampleMultiplicativeExpression_mod() {
	fmt.Println(exampleAST(44, "int i = x % y;"))
	// Output:
	// TODO
}

func ExampleParameterDeclaration_decl() {
	fmt.Println(exampleAST(188, "int f(int i) {}"))
	// Output:
	// TODO
}

func ExampleParameterDeclaration_abstract() {
	fmt.Println(exampleAST(189, "int f(int*) {}"))
	// Output:
	// TODO
}

func ExampleParameterList_case0() {
	fmt.Println(exampleAST(186, "int f(int i) {}"))
	// Output:
	// TODO
}

func ExampleParameterList_case1() {
	fmt.Println(exampleAST(187, "int f(int i, int j) {}"))
	// Output:
	// TODO
}

func ExampleParameterTypeList_list() {
	fmt.Println(exampleAST(184, "int f(int i) {}"))
	// Output:
	// TODO
}

func ExampleParameterTypeList_var() {
	fmt.Println(exampleAST(185, "int f(int i, ...) {}"))
	// Output:
	// TODO
}

func ExamplePointer_typeQual() {
	fmt.Println(exampleAST(177, "int *p;"))
	// Output:
	// TODO
}

func ExamplePointer_ptr() {
	fmt.Println(exampleAST(178, "int **p;"))
	// Output:
	// TODO
}

func ExamplePointer_block() {
	fmt.Println(exampleAST(179, "int atexit_b(void (^ _Nonnull)(void));"))
	// Output:
	// TODO
}

func ExamplePostfixExpression_primary() {
	fmt.Println(exampleAST(11, "int i = x;"))
	// Output:
	// TODO
}

func ExamplePostfixExpression_index() {
	fmt.Println(exampleAST(12, "int i = x[y];"))
	// Output:
	// TODO
}

func ExamplePostfixExpression_call() {
	fmt.Println(exampleAST(13, "int i = x(y);"))
	// Output:
	// TODO
}

func ExamplePostfixExpression_select() {
	fmt.Println(exampleAST(14, "int i = x.y;"))
	// Output:
	// TODO
}

func ExamplePostfixExpression_pSelect() {
	fmt.Println(exampleAST(15, "int i = x->y;"))
	// Output:
	// TODO
}

func ExamplePostfixExpression_inc() {
	fmt.Println(exampleAST(16, "int i = x++;"))
	// Output:
	// TODO
}

func ExamplePostfixExpression_dec() {
	fmt.Println(exampleAST(17, "int i = x--;"))
	// Output:
	// TODO
}

func ExamplePostfixExpression_complit() {
	fmt.Println(exampleAST(18, "int i = (int[]){y};"))
	// Output:
	// TODO
}

func ExamplePostfixExpression_typeCmp() {
	fmt.Println(exampleAST(19, "int i = __builtin_types_compatible_p(int, double);"))
	// Output:
	// TODO
}

func ExamplePostfixExpression_chooseExpr() {
	fmt.Println(exampleAST(20, "int i = __builtin_choose_expr(1, 2, \"foo\");"))
	// Output:
	// TODO
}

func ExamplePragmaSTDC_case0() {
	fmt.Println(exampleAST(277, "_Pragma(\"STDC FP_CONTRACT ON\")"))
	// Output:
	// TODO
}

func ExamplePrimaryExpression_ident() {
	fmt.Println(exampleAST(1, "int i = x;"))
	// Output:
	// TODO
}

func ExamplePrimaryExpression_int() {
	fmt.Println(exampleAST(2, "int i = 42;"))
	// Output:
	// TODO
}

func ExamplePrimaryExpression_float() {
	fmt.Println(exampleAST(3, "int i = 3.14;"))
	// Output:
	// TODO
}

func ExamplePrimaryExpression_enum() {
	fmt.Println(exampleAST(4, "enum e {a}; int i = a;"))
	// Output:
	// TODO
}

func ExamplePrimaryExpression_char() {
	fmt.Println(exampleAST(5, "int i = 'x';"))
	// Output:
	// TODO
}

func ExamplePrimaryExpression_lChar() {
	fmt.Println(exampleAST(6, "int i = L'x';"))
	// Output:
	// TODO
}

func ExamplePrimaryExpression_string() {
	fmt.Println(exampleAST(7, "char *c = \"x\";"))
	// Output:
	// TODO
}

func ExamplePrimaryExpression_lString() {
	fmt.Println(exampleAST(8, "char *c = L\"x\";"))
	// Output:
	// TODO
}

func ExamplePrimaryExpression_expr() {
	fmt.Println(exampleAST(9, "int i = (x+y);"))
	// Output:
	// TODO
}

func ExamplePrimaryExpression_stmt() {
	fmt.Println(exampleAST(10, "int i = ({x();});"))
	// Output:
	// TODO
}

func ExampleRelationalExpression_shift() {
	fmt.Println(exampleAST(51, "int i = x;"))
	// Output:
	// TODO
}

func ExampleRelationalExpression_lt() {
	fmt.Println(exampleAST(52, "int i = x < y;"))
	// Output:
	// TODO
}

func ExampleRelationalExpression_gt() {
	fmt.Println(exampleAST(53, "int i = x > y;"))
	// Output:
	// TODO
}

func ExampleRelationalExpression_leq() {
	fmt.Println(exampleAST(54, "int i = x <= y;"))
	// Output:
	// TODO
}

func ExampleRelationalExpression_geq() {
	fmt.Println(exampleAST(55, "int i = x >= y;"))
	// Output:
	// TODO
}

func ExampleSelectionStatement_if() {
	fmt.Println(exampleAST(231, "int f() { if(x) y(); }"))
	// Output:
	// TODO
}

func ExampleSelectionStatement_ifElse() {
	fmt.Println(exampleAST(232, "int f() { if(x) y(); else z(); }"))
	// Output:
	// TODO
}

func ExampleSelectionStatement_switch() {
	fmt.Println(exampleAST(233, "int f() { switch(i) case 42: x(); }"))
	// Output:
	// TODO
}

func ExampleShiftExpression_add() {
	fmt.Println(exampleAST(48, "int i = x;"))
	// Output:
	// TODO
}

func ExampleShiftExpression_lsh() {
	fmt.Println(exampleAST(49, "int i = x << y;"))
	// Output:
	// TODO
}

func ExampleShiftExpression_rsh() {
	fmt.Println(exampleAST(50, "int i = x >> y;"))
	// Output:
	// TODO
}

func ExampleSpecifierQualifierList_typeSpec() {
	fmt.Println(exampleAST(145, "struct {int i;};"))
	// Output:
	// TODO
}

func ExampleSpecifierQualifierList_typeQual() {
	fmt.Println(exampleAST(146, "struct {const int i;};"))
	// Output:
	// TODO
}

func ExampleSpecifierQualifierList_alignSpec() {
	fmt.Println(exampleAST(147, "struct {_Alignas(double) int i;};"))
	// Output:
	// TODO
}

func ExampleSpecifierQualifierList_attribute() {
	fmt.Println(exampleAST(148, "struct {__attribute__((a)) int i;};"))
	// Output:
	// TODO
}

func ExampleStatement_labeled() {
	fmt.Println(exampleAST(211, "int f() { L: x(); }"))
	// Output:
	// TODO
}

func ExampleStatement_compound() {
	fmt.Println(exampleAST(212, "int f() { { y(); } }"))
	// Output:
	// TODO
}

func ExampleStatement_expr() {
	fmt.Println(exampleAST(213, "int f() { x(); }"))
	// Output:
	// TODO
}

func ExampleStatement_selection() {
	fmt.Println(exampleAST(214, "int f() { if(x) y(); }"))
	// Output:
	// TODO
}

func ExampleStatement_iteration() {
	fmt.Println(exampleAST(215, "int f() { for(;;) x(); }"))
	// Output:
	// TODO
}

func ExampleStatement_jump() {
	fmt.Println(exampleAST(216, "int f() { return x; }"))
	// Output:
	// TODO
}

func ExampleStatement_asm() {
	fmt.Println(exampleAST(217, "int f() { __asm__(\"nop\"); }"))
	// Output:
	// TODO
}

func ExampleStorageClassSpecifier_typedef() {
	fmt.Println(exampleAST(97, "typedef int int_t;"))
	// Output:
	// TODO
}

func ExampleStorageClassSpecifier_extern() {
	fmt.Println(exampleAST(98, "extern int i;"))
	// Output:
	// TODO
}

func ExampleStorageClassSpecifier_static() {
	fmt.Println(exampleAST(99, "static int i;"))
	// Output:
	// TODO
}

func ExampleStorageClassSpecifier_auto() {
	fmt.Println(exampleAST(100, "auto int i;"))
	// Output:
	// TODO
}

func ExampleStorageClassSpecifier_register() {
	fmt.Println(exampleAST(101, "register int i;"))
	// Output:
	// TODO
}

func ExampleStorageClassSpecifier_threadLocal() {
	fmt.Println(exampleAST(102, "_Thread_local int i;"))
	// Output:
	// TODO
}

func ExampleStructDeclaration_case0() {
	fmt.Println(exampleAST(144, "struct{ int i; }"))
	// Output:
	// TODO
}

func ExampleStructDeclarationList_case0() {
	fmt.Println(exampleAST(142, "struct{ int i; }"))
	// Output:
	// TODO
}

func ExampleStructDeclarationList_case1() {
	fmt.Println(exampleAST(143, "struct{ int i; double d; }"))
	// Output:
	// TODO
}

func ExampleStructDeclarator_decl() {
	fmt.Println(exampleAST(151, "struct{ int i; }"))
	// Output:
	// TODO
}

func ExampleStructDeclarator_bitField() {
	fmt.Println(exampleAST(152, "struct{ int i:3; }"))
	// Output:
	// TODO
}

func ExampleStructDeclaratorList_case0() {
	fmt.Println(exampleAST(149, "struct{ int i; }"))
	// Output:
	// TODO
}

func ExampleStructDeclaratorList_case1() {
	fmt.Println(exampleAST(150, "struct{ int i, j; }"))
	// Output:
	// TODO
}

func ExampleStructOrUnion_struct() {
	fmt.Println(exampleAST(140, "struct { int i; } s;"))
	// Output:
	// TODO
}

func ExampleStructOrUnion_union() {
	fmt.Println(exampleAST(141, "union { int i; double d; } u;"))
	// Output:
	// TODO
}

func ExampleStructOrUnionSpecifier_def() {
	fmt.Println(exampleAST(138, "struct s { int i; };"))
	// Output:
	// TODO
}

func ExampleStructOrUnionSpecifier_tag() {
	fmt.Println(exampleAST(139, "struct s v;"))
	// Output:
	// TODO
}

func ExampleTranslationUnit_case0() {
	fmt.Println(exampleAST(243, "int i;"))
	// Output:
	// TODO
}

func ExampleTranslationUnit_case1() {
	fmt.Println(exampleAST(244, "int i; int j;"))
	// Output:
	// TODO
}

func ExampleTypeName_case0() {
	fmt.Println(exampleAST(192, "int i = (int)x;"))
	// Output:
	// TODO
}

func ExampleTypeQualifier_const() {
	fmt.Println(exampleAST(160, "const int i;"))
	// Output:
	// TODO
}

func ExampleTypeQualifier_restrict() {
	fmt.Println(exampleAST(161, "restrict int i;"))
	// Output:
	// TODO
}

func ExampleTypeQualifier_volatile() {
	fmt.Println(exampleAST(162, "volatile int i;"))
	// Output:
	// TODO
}

func ExampleTypeQualifier_atomic() {
	fmt.Println(exampleAST(163, "_Atomic int i;"))
	// Output:
	// TODO
}

func ExampleTypeQualifiers_typeQual() {
	fmt.Println(exampleAST(180, "int * const i;"))
	// Output:
	// TODO
}

func ExampleTypeQualifiers_attribute() {
	fmt.Println(exampleAST(181, "int * __attribute__((a)) i;"))
	// Output:
	// TODO
}

func ExampleTypeQualifiers_case2() {
	fmt.Println(exampleAST(182, "int * const volatile i;"))
	// Output:
	// TODO
}

func ExampleTypeQualifiers_case3() {
	fmt.Println(exampleAST(183, "int * __attribute__((a)) __attribute__((b)) i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_void() {
	fmt.Println(exampleAST(103, "void i();"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_char() {
	fmt.Println(exampleAST(104, "char i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_short() {
	fmt.Println(exampleAST(105, "short i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_int() {
	fmt.Println(exampleAST(106, "int i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_int8() {
	fmt.Println(exampleAST(107, "__int8 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_int16() {
	fmt.Println(exampleAST(108, "__int16 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_int32() {
	fmt.Println(exampleAST(109, "__int32 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_int64() {
	fmt.Println(exampleAST(110, "__int64 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_int128() {
	fmt.Println(exampleAST(111, "__int128 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_long() {
	fmt.Println(exampleAST(112, "long i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_float() {
	fmt.Println(exampleAST(113, "float i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_float16() {
	fmt.Println(exampleAST(114, "__fp16 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_decimal32() {
	fmt.Println(exampleAST(115, "_Decimal32 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_decimal64() {
	fmt.Println(exampleAST(116, "_Decimal64 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_decimal128() {
	fmt.Println(exampleAST(117, "_Decimal128 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_float128() {
	fmt.Println(exampleAST(118, "_Float128 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_float80() {
	fmt.Println(exampleAST(119, "__float80 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_double() {
	fmt.Println(exampleAST(120, "double i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_signed() {
	fmt.Println(exampleAST(121, "signed i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_unsigned() {
	fmt.Println(exampleAST(122, "unsigned i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_bool() {
	fmt.Println(exampleAST(123, "_Bool i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_complex() {
	fmt.Println(exampleAST(124, "_Complex i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_structOrUnion() {
	fmt.Println(exampleAST(125, "struct s i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_enum() {
	fmt.Println(exampleAST(126, "enum e i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_typedefName() {
	fmt.Println(exampleAST(127, "typedef const T; T i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_typeofExpr() {
	fmt.Println(exampleAST(128, "typeof(42) i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_typeofType() {
	fmt.Println(exampleAST(129, "typedef const T; typeof(T) i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_atomic() {
	fmt.Println(exampleAST(130, "_Atomic(int) i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_fract() {
	fmt.Println(exampleAST(131, "_Fract i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_sat() {
	fmt.Println(exampleAST(132, "_Sat i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_accum() {
	fmt.Println(exampleAST(133, "_Accum i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_float32() {
	fmt.Println(exampleAST(134, "_Float32 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_float64() {
	fmt.Println(exampleAST(135, "_Float64 i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_float32x() {
	fmt.Println(exampleAST(136, "_Float32x i;"))
	// Output:
	// TODO
}

func ExampleTypeSpecifier_float64x() {
	fmt.Println(exampleAST(137, "_Float64x i;"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_postfix() {
	fmt.Println(exampleAST(23, "int i = x;"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_inc() {
	fmt.Println(exampleAST(24, "int i = ++x;"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_dec() {
	fmt.Println(exampleAST(25, "int i = --x;"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_addrof() {
	fmt.Println(exampleAST(26, "int *i = &x;"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_deref() {
	fmt.Println(exampleAST(27, "int i = *x;"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_plus() {
	fmt.Println(exampleAST(28, "int i = +x;"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_minus() {
	fmt.Println(exampleAST(29, "int i = -x;"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_cpl() {
	fmt.Println(exampleAST(30, "int i = ~x;"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_not() {
	fmt.Println(exampleAST(31, "int i = !x;"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_sizeofExpr() {
	fmt.Println(exampleAST(32, "int i = sizeof x;"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_sizeofType() {
	fmt.Println(exampleAST(33, "int i = sizeof(int);"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_labelAddr() {
	fmt.Println(exampleAST(34, "int f() { L: &&L; }"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_alignofExpr() {
	fmt.Println(exampleAST(35, "int i = _Alignof(x);"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_alignofType() {
	fmt.Println(exampleAST(36, "int i = _Alignof(int);"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_imag() {
	fmt.Println(exampleAST(37, "double i = __imag__ x;"))
	// Output:
	// TODO
}

func ExampleUnaryExpression_real() {
	fmt.Println(exampleAST(38, "double i = __real__ x;"))
	// Output:
	// TODO
}