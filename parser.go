// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2016 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.10. Substantial portions of expression AST size
// optimizations are from [2], license of which follows.

// ----------------------------------------------------------------------------

// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This grammar is derived from the C grammar in the 'ansitize'
// program, which carried this notice:
//
// Copyright (c) 2006 Russ Cox,
// 	Massachusetts Institute of Technology
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the
// Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute,
// sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall
// be included in all copies or substantial portions of the
// Software.
//
// The software is provided "as is", without warranty of any
// kind, express or implied, including but not limited to the
// warranties of merchantability, fitness for a particular
// purpose and noninfringement.  In no event shall the authors
// or copyright holders be liable for any claim, damages or
// other liability, whether in an action of contract, tort or
// otherwise, arising from, out of or in connection with the
// software or the use or other dealings in the software.

package cc

import __yyfmt__ "fmt"

import (
	"fmt"

	"github.com/cznic/golex/lex"
	"github.com/cznic/xc"
)

type yySymType struct {
	yys       int
	Token     xc.Token
	groupPart Node
	node      Node
	toks      PPTokenList
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault           = 57448
	yyEofCode           = 57344
	ADDASSIGN           = 57346
	ANDAND              = 57347
	ANDASSIGN           = 57348
	ARROW               = 57349
	AUTO                = 57350
	BOOL                = 57351
	BREAK               = 57352
	CASE                = 57353
	CAST                = 57354
	CHAR                = 57355
	CHARCONST           = 57356
	COMPLEX             = 57357
	CONST               = 57358
	CONSTANT_EXPRESSION = 1048577
	CONTINUE            = 57359
	DDD                 = 57360
	DEC                 = 57361
	DEFAULT             = 57362
	DIVASSIGN           = 57363
	DO                  = 57364
	DOUBLE              = 57365
	ELSE                = 57366
	ENUM                = 57367
	EQ                  = 57368
	EXTERN              = 57369
	FLOAT               = 57370
	FLOATCONST          = 57371
	FOR                 = 57372
	GEQ                 = 57373
	GOTO                = 57374
	IDENTIFIER          = 57375
	IDENTIFIER_LPAREN   = 57376
	IF                  = 57377
	INC                 = 57378
	INLINE              = 57379
	INT                 = 57380
	INTCONST            = 57381
	LEQ                 = 57382
	LONG                = 57383
	LONGCHARCONST       = 57384
	LONGSTRINGLITERAL   = 57385
	LSH                 = 57386
	LSHASSIGN           = 57387
	MODASSIGN           = 57388
	MULASSIGN           = 57389
	NEQ                 = 57390
	NOELSE              = 57391
	ORASSIGN            = 57392
	OROR                = 57393
	PPDEFINE            = 57394
	PPELIF              = 57395
	PPELSE              = 57396
	PPENDIF             = 57397
	PPERROR             = 57398
	PPHASH_NL           = 57399
	PPHEADER_NAME       = 57400
	PPIF                = 57401
	PPIFDEF             = 57402
	PPIFNDEF            = 57403
	PPINCLUDE           = 57404
	PPLINE              = 57405
	PPNONDIRECTIVE      = 57406
	PPNUMBER            = 57407
	PPOTHER             = 57408
	PPPASTE             = 57409
	PPPRAGMA            = 57410
	PPUNDEF             = 57411
	PREPROCESSING_FILE  = 1048576
	REGISTER            = 57412
	RESTRICT            = 57413
	RETURN              = 57414
	RSH                 = 57415
	RSHASSIGN           = 57416
	SHORT               = 57417
	SIGNED              = 57418
	SIZEOF              = 57419
	STATIC              = 57420
	STRINGLITERAL       = 57421
	STRUCT              = 57422
	SUBASSIGN           = 57423
	SWITCH              = 57424
	TRANSLATION_UNIT    = 1048578
	TYPEDEF             = 57425
	TYPEDEFNAME         = 57426
	UNARY               = 57427
	UNION               = 57428
	UNSIGNED            = 57429
	VOID                = 57430
	VOLATILE            = 57431
	WHILE               = 57432
	XORASSIGN           = 57433
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -285
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (278x)
		42:      1,   // '*' (249x)
		57375:   2,   // IDENTIFIER (203x)
		38:      3,   // '&' (199x)
		43:      4,   // '+' (199x)
		45:      5,   // '-' (199x)
		57361:   6,   // DEC (199x)
		57378:   7,   // INC (199x)
		59:      8,   // ';' (181x)
		41:      9,   // ')' (176x)
		44:      10,  // ',' (168x)
		91:      11,  // '[' (149x)
		33:      12,  // '!' (132x)
		126:     13,  // '~' (132x)
		57356:   14,  // CHARCONST (132x)
		57371:   15,  // FLOATCONST (132x)
		57381:   16,  // INTCONST (132x)
		57384:   17,  // LONGCHARCONST (132x)
		57385:   18,  // LONGSTRINGLITERAL (132x)
		57419:   19,  // SIZEOF (132x)
		57421:   20,  // STRINGLITERAL (132x)
		57358:   21,  // CONST (113x)
		57413:   22,  // RESTRICT (113x)
		57431:   23,  // VOLATILE (113x)
		125:     24,  // '}' (110x)
		57351:   25,  // BOOL (103x)
		57355:   26,  // CHAR (103x)
		57357:   27,  // COMPLEX (103x)
		57365:   28,  // DOUBLE (103x)
		57367:   29,  // ENUM (103x)
		57370:   30,  // FLOAT (103x)
		57380:   31,  // INT (103x)
		57383:   32,  // LONG (103x)
		57417:   33,  // SHORT (103x)
		57418:   34,  // SIGNED (103x)
		57422:   35,  // STRUCT (103x)
		57426:   36,  // TYPEDEFNAME (103x)
		57428:   37,  // UNION (103x)
		57429:   38,  // UNSIGNED (103x)
		57430:   39,  // VOID (103x)
		58:      40,  // ':' (102x)
		57420:   41,  // STATIC (97x)
		57344:   42,  // $end (95x)
		57350:   43,  // AUTO (91x)
		57369:   44,  // EXTERN (91x)
		57379:   45,  // INLINE (91x)
		57412:   46,  // REGISTER (91x)
		57425:   47,  // TYPEDEF (91x)
		61:      48,  // '=' (87x)
		57484:   49,  // Expression (83x)
		93:      50,  // ']' (81x)
		46:      51,  // '.' (76x)
		123:     52,  // '{' (75x)
		37:      53,  // '%' (68x)
		47:      54,  // '/' (68x)
		60:      55,  // '<' (68x)
		62:      56,  // '>' (68x)
		63:      57,  // '?' (68x)
		94:      58,  // '^' (68x)
		124:     59,  // '|' (68x)
		57346:   60,  // ADDASSIGN (68x)
		57347:   61,  // ANDAND (68x)
		57348:   62,  // ANDASSIGN (68x)
		57349:   63,  // ARROW (68x)
		57363:   64,  // DIVASSIGN (68x)
		57368:   65,  // EQ (68x)
		57373:   66,  // GEQ (68x)
		57382:   67,  // LEQ (68x)
		57386:   68,  // LSH (68x)
		57387:   69,  // LSHASSIGN (68x)
		57388:   70,  // MODASSIGN (68x)
		57389:   71,  // MULASSIGN (68x)
		57390:   72,  // NEQ (68x)
		57392:   73,  // ORASSIGN (68x)
		57393:   74,  // OROR (68x)
		57415:   75,  // RSH (68x)
		57416:   76,  // RSHASSIGN (68x)
		57423:   77,  // SUBASSIGN (68x)
		57433:   78,  // XORASSIGN (68x)
		10:      79,  // '\n' (58x)
		57408:   80,  // PPOTHER (52x)
		57397:   81,  // PPENDIF (44x)
		57432:   82,  // WHILE (41x)
		57352:   83,  // BREAK (40x)
		57353:   84,  // CASE (40x)
		57359:   85,  // CONTINUE (40x)
		57362:   86,  // DEFAULT (40x)
		57364:   87,  // DO (40x)
		57372:   88,  // FOR (40x)
		57374:   89,  // GOTO (40x)
		57377:   90,  // IF (40x)
		57396:   91,  // PPELSE (40x)
		57414:   92,  // RETURN (40x)
		57424:   93,  // SWITCH (40x)
		57395:   94,  // PPELIF (39x)
		57394:   95,  // PPDEFINE (35x)
		57398:   96,  // PPERROR (35x)
		57399:   97,  // PPHASH_NL (35x)
		57401:   98,  // PPIF (35x)
		57402:   99,  // PPIFDEF (35x)
		57403:   100, // PPIFNDEF (35x)
		57404:   101, // PPINCLUDE (35x)
		57405:   102, // PPLINE (35x)
		57406:   103, // PPNONDIRECTIVE (35x)
		57410:   104, // PPPRAGMA (35x)
		57411:   105, // PPUNDEF (35x)
		57534:   106, // TypeQualifier (28x)
		57485:   107, // ExpressionList (26x)
		57366:   108, // ELSE (22x)
		57508:   109, // PPTokenList (22x)
		57510:   110, // PPTokens (22x)
		57480:   111, // EnumSpecifier (20x)
		57529:   112, // StructOrUnion (20x)
		57530:   113, // StructOrUnionSpecifier (20x)
		57537:   114, // TypeSpecifier (20x)
		57486:   115, // ExpressionListOpt (18x)
		57509:   116, // PPTokenListOpt (16x)
		57463:   117, // DeclarationSpecifiers (15x)
		57491:   118, // FunctionSpecifier (15x)
		57524:   119, // StorageClassSpecifier (15x)
		57457:   120, // CompoundStatement (13x)
		57488:   121, // ExpressionStatement (12x)
		57505:   122, // IterationStatement (12x)
		57506:   123, // JumpStatement (12x)
		57507:   124, // LabeledStatement (12x)
		57519:   125, // SelectionStatement (12x)
		57523:   126, // Statement (12x)
		57515:   127, // Pointer (11x)
		57516:   128, // PointerOpt (10x)
		57459:   129, // ControlLine (8x)
		57465:   130, // Declarator (8x)
		57494:   131, // GroupPart (8x)
		57498:   132, // IfGroup (8x)
		57499:   133, // IfSection (8x)
		57531:   134, // TextLine (8x)
		57460:   135, // Declaration (7x)
		57492:   136, // GroupList (6x)
		57518:   137, // ReplacementList (6x)
		57442:   138, // $@4 (5x)
		57458:   139, // ConstantExpression (5x)
		57360:   140, // DDD (5x)
		57493:   141, // GroupListOpt (5x)
		57520:   142, // SpecifierQualifierList (5x)
		57535:   143, // TypeQualifierList (5x)
		57449:   144, // AbstractDeclarator (4x)
		57464:   145, // DeclarationSpecifiersOpt (4x)
		57469:   146, // Designator (4x)
		57511:   147, // ParameterDeclaration (4x)
		57536:   148, // TypeQualifierListOpt (4x)
		57456:   149, // CommaOpt (3x)
		57467:   150, // Designation (3x)
		57468:   151, // DesignationOpt (3x)
		57470:   152, // DesignatorList (3x)
		57487:   153, // ExpressionOpt (3x)
		57500:   154, // InitDeclarator (3x)
		57503:   155, // Initializer (3x)
		57512:   156, // ParameterList (3x)
		57513:   157, // ParameterTypeList (3x)
		57435:   158, // $@10 (2x)
		57436:   159, // $@11 (2x)
		57443:   160, // $@5 (2x)
		57450:   161, // AbstractDeclaratorOpt (2x)
		57453:   162, // BlockItem (2x)
		57466:   163, // DeclaratorOpt (2x)
		57471:   164, // DirectAbstractDeclarator (2x)
		57472:   165, // DirectAbstractDeclaratorOpt (2x)
		57473:   166, // DirectDeclarator (2x)
		57474:   167, // ElifGroup (2x)
		57481:   168, // EnumerationConstant (2x)
		57482:   169, // Enumerator (2x)
		57489:   170, // ExternalDeclaration (2x)
		57490:   171, // FunctionDefinition (2x)
		57495:   172, // IdentifierList (2x)
		57496:   173, // IdentifierListOpt (2x)
		57497:   174, // IdentifierOpt (2x)
		57501:   175, // InitDeclaratorList (2x)
		57502:   176, // InitDeclaratorListOpt (2x)
		57504:   177, // InitializerList (2x)
		57514:   178, // ParameterTypeListOpt (2x)
		57521:   179, // SpecifierQualifierListOpt (2x)
		57525:   180, // StructDeclaration (2x)
		57527:   181, // StructDeclarator (2x)
		57533:   182, // TypeName (2x)
		57434:   183, // $@1 (1x)
		57437:   184, // $@12 (1x)
		57438:   185, // $@13 (1x)
		57439:   186, // $@14 (1x)
		57440:   187, // $@2 (1x)
		57441:   188, // $@3 (1x)
		57444:   189, // $@6 (1x)
		57445:   190, // $@7 (1x)
		57446:   191, // $@8 (1x)
		57447:   192, // $@9 (1x)
		57451:   193, // ArgumentExpressionList (1x)
		57452:   194, // ArgumentExpressionListOpt (1x)
		57454:   195, // BlockItemList (1x)
		57455:   196, // BlockItemListOpt (1x)
		1048577: 197, // CONSTANT_EXPRESSION (1x)
		57461:   198, // DeclarationList (1x)
		57462:   199, // DeclarationListOpt (1x)
		57475:   200, // ElifGroupList (1x)
		57476:   201, // ElifGroupListOpt (1x)
		57477:   202, // ElseGroup (1x)
		57478:   203, // ElseGroupOpt (1x)
		57479:   204, // EndifLine (1x)
		57483:   205, // EnumeratorList (1x)
		57376:   206, // IDENTIFIER_LPAREN (1x)
		1048576: 207, // PREPROCESSING_FILE (1x)
		57517:   208, // PreprocessingFile (1x)
		57522:   209, // Start (1x)
		57526:   210, // StructDeclarationList (1x)
		57528:   211, // StructDeclaratorList (1x)
		1048578: 212, // TRANSLATION_UNIT (1x)
		57532:   213, // TranslationUnit (1x)
		57448:   214, // $default (0x)
		57354:   215, // CAST (0x)
		57345:   216, // error (0x)
		57391:   217, // NOELSE (0x)
		57400:   218, // PPHEADER_NAME (0x)
		57407:   219, // PPNUMBER (0x)
		57409:   220, // PPPASTE (0x)
		57427:   221, // UNARY (0x)
	}

	yySymNames = []string{
		"'('",
		"'*'",
		"IDENTIFIER",
		"'&'",
		"'+'",
		"'-'",
		"DEC",
		"INC",
		"';'",
		"')'",
		"','",
		"'['",
		"'!'",
		"'~'",
		"CHARCONST",
		"FLOATCONST",
		"INTCONST",
		"LONGCHARCONST",
		"LONGSTRINGLITERAL",
		"SIZEOF",
		"STRINGLITERAL",
		"CONST",
		"RESTRICT",
		"VOLATILE",
		"'}'",
		"BOOL",
		"CHAR",
		"COMPLEX",
		"DOUBLE",
		"ENUM",
		"FLOAT",
		"INT",
		"LONG",
		"SHORT",
		"SIGNED",
		"STRUCT",
		"TYPEDEFNAME",
		"UNION",
		"UNSIGNED",
		"VOID",
		"':'",
		"STATIC",
		"$end",
		"AUTO",
		"EXTERN",
		"INLINE",
		"REGISTER",
		"TYPEDEF",
		"'='",
		"Expression",
		"']'",
		"'.'",
		"'{'",
		"'%'",
		"'/'",
		"'<'",
		"'>'",
		"'?'",
		"'^'",
		"'|'",
		"ADDASSIGN",
		"ANDAND",
		"ANDASSIGN",
		"ARROW",
		"DIVASSIGN",
		"EQ",
		"GEQ",
		"LEQ",
		"LSH",
		"LSHASSIGN",
		"MODASSIGN",
		"MULASSIGN",
		"NEQ",
		"ORASSIGN",
		"OROR",
		"RSH",
		"RSHASSIGN",
		"SUBASSIGN",
		"XORASSIGN",
		"'\\n'",
		"PPOTHER",
		"PPENDIF",
		"WHILE",
		"BREAK",
		"CASE",
		"CONTINUE",
		"DEFAULT",
		"DO",
		"FOR",
		"GOTO",
		"IF",
		"PPELSE",
		"RETURN",
		"SWITCH",
		"PPELIF",
		"PPDEFINE",
		"PPERROR",
		"PPHASH_NL",
		"PPIF",
		"PPIFDEF",
		"PPIFNDEF",
		"PPINCLUDE",
		"PPLINE",
		"PPNONDIRECTIVE",
		"PPPRAGMA",
		"PPUNDEF",
		"TypeQualifier",
		"ExpressionList",
		"ELSE",
		"PPTokenList",
		"PPTokens",
		"EnumSpecifier",
		"StructOrUnion",
		"StructOrUnionSpecifier",
		"TypeSpecifier",
		"ExpressionListOpt",
		"PPTokenListOpt",
		"DeclarationSpecifiers",
		"FunctionSpecifier",
		"StorageClassSpecifier",
		"CompoundStatement",
		"ExpressionStatement",
		"IterationStatement",
		"JumpStatement",
		"LabeledStatement",
		"SelectionStatement",
		"Statement",
		"Pointer",
		"PointerOpt",
		"ControlLine",
		"Declarator",
		"GroupPart",
		"IfGroup",
		"IfSection",
		"TextLine",
		"Declaration",
		"GroupList",
		"ReplacementList",
		"$@4",
		"ConstantExpression",
		"DDD",
		"GroupListOpt",
		"SpecifierQualifierList",
		"TypeQualifierList",
		"AbstractDeclarator",
		"DeclarationSpecifiersOpt",
		"Designator",
		"ParameterDeclaration",
		"TypeQualifierListOpt",
		"CommaOpt",
		"Designation",
		"DesignationOpt",
		"DesignatorList",
		"ExpressionOpt",
		"InitDeclarator",
		"Initializer",
		"ParameterList",
		"ParameterTypeList",
		"$@10",
		"$@11",
		"$@5",
		"AbstractDeclaratorOpt",
		"BlockItem",
		"DeclaratorOpt",
		"DirectAbstractDeclarator",
		"DirectAbstractDeclaratorOpt",
		"DirectDeclarator",
		"ElifGroup",
		"EnumerationConstant",
		"Enumerator",
		"ExternalDeclaration",
		"FunctionDefinition",
		"IdentifierList",
		"IdentifierListOpt",
		"IdentifierOpt",
		"InitDeclaratorList",
		"InitDeclaratorListOpt",
		"InitializerList",
		"ParameterTypeListOpt",
		"SpecifierQualifierListOpt",
		"StructDeclaration",
		"StructDeclarator",
		"TypeName",
		"$@1",
		"$@12",
		"$@13",
		"$@14",
		"$@2",
		"$@3",
		"$@6",
		"$@7",
		"$@8",
		"$@9",
		"ArgumentExpressionList",
		"ArgumentExpressionListOpt",
		"BlockItemList",
		"BlockItemListOpt",
		"CONSTANT_EXPRESSION",
		"DeclarationList",
		"DeclarationListOpt",
		"ElifGroupList",
		"ElifGroupListOpt",
		"ElseGroup",
		"ElseGroupOpt",
		"EndifLine",
		"EnumeratorList",
		"IDENTIFIER_LPAREN",
		"PREPROCESSING_FILE",
		"PreprocessingFile",
		"Start",
		"StructDeclarationList",
		"StructDeclaratorList",
		"TRANSLATION_UNIT",
		"TranslationUnit",
		"$default",
		"CAST",
		"error",
		"NOELSE",
		"PPHEADER_NAME",
		"PPNUMBER",
		"PPPASTE",
		"UNARY",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {183, 0},
		2:   {209, 3},
		3:   {187, 0},
		4:   {209, 3},
		5:   {188, 0},
		6:   {209, 3},
		7:   {168, 1},
		8:   {193, 1},
		9:   {193, 3},
		10:  {194, 0},
		11:  {194, 1},
		12:  {49, 1},
		13:  {49, 1},
		14:  {49, 1},
		15:  {49, 1},
		16:  {49, 1},
		17:  {49, 1},
		18:  {49, 1},
		19:  {49, 3},
		20:  {49, 4},
		21:  {49, 4},
		22:  {49, 3},
		23:  {49, 3},
		24:  {49, 2},
		25:  {49, 2},
		26:  {49, 7},
		27:  {49, 2},
		28:  {49, 2},
		29:  {49, 2},
		30:  {49, 2},
		31:  {49, 2},
		32:  {49, 2},
		33:  {49, 2},
		34:  {49, 2},
		35:  {49, 2},
		36:  {49, 4},
		37:  {49, 4},
		38:  {49, 3},
		39:  {49, 3},
		40:  {49, 3},
		41:  {49, 3},
		42:  {49, 3},
		43:  {49, 3},
		44:  {49, 3},
		45:  {49, 3},
		46:  {49, 3},
		47:  {49, 3},
		48:  {49, 3},
		49:  {49, 3},
		50:  {49, 3},
		51:  {49, 3},
		52:  {49, 3},
		53:  {49, 3},
		54:  {49, 3},
		55:  {49, 3},
		56:  {49, 5},
		57:  {49, 3},
		58:  {49, 3},
		59:  {49, 3},
		60:  {49, 3},
		61:  {49, 3},
		62:  {49, 3},
		63:  {49, 3},
		64:  {49, 3},
		65:  {49, 3},
		66:  {49, 3},
		67:  {49, 3},
		68:  {153, 0},
		69:  {153, 1},
		70:  {107, 1},
		71:  {107, 3},
		72:  {115, 0},
		73:  {115, 1},
		74:  {138, 0},
		75:  {139, 2},
		76:  {135, 3},
		77:  {117, 2},
		78:  {117, 2},
		79:  {117, 2},
		80:  {117, 2},
		81:  {145, 0},
		82:  {145, 1},
		83:  {175, 1},
		84:  {175, 3},
		85:  {176, 0},
		86:  {176, 1},
		87:  {154, 1},
		88:  {160, 0},
		89:  {154, 4},
		90:  {119, 1},
		91:  {119, 1},
		92:  {119, 1},
		93:  {119, 1},
		94:  {119, 1},
		95:  {114, 1},
		96:  {114, 1},
		97:  {114, 1},
		98:  {114, 1},
		99:  {114, 1},
		100: {114, 1},
		101: {114, 1},
		102: {114, 1},
		103: {114, 1},
		104: {114, 1},
		105: {114, 1},
		106: {114, 1},
		107: {114, 1},
		108: {114, 1},
		109: {189, 0},
		110: {190, 0},
		111: {113, 7},
		112: {113, 2},
		113: {112, 1},
		114: {112, 1},
		115: {210, 1},
		116: {210, 2},
		117: {180, 3},
		118: {180, 2},
		119: {142, 2},
		120: {142, 2},
		121: {179, 0},
		122: {179, 1},
		123: {211, 1},
		124: {211, 3},
		125: {181, 1},
		126: {181, 3},
		127: {149, 0},
		128: {149, 1},
		129: {191, 0},
		130: {111, 7},
		131: {111, 2},
		132: {205, 1},
		133: {205, 3},
		134: {169, 1},
		135: {169, 3},
		136: {106, 1},
		137: {106, 1},
		138: {106, 1},
		139: {118, 1},
		140: {130, 2},
		141: {163, 0},
		142: {163, 1},
		143: {166, 1},
		144: {166, 3},
		145: {166, 5},
		146: {166, 6},
		147: {166, 6},
		148: {166, 5},
		149: {192, 0},
		150: {166, 5},
		151: {166, 4},
		152: {127, 2},
		153: {127, 3},
		154: {128, 0},
		155: {128, 1},
		156: {143, 1},
		157: {143, 2},
		158: {148, 0},
		159: {148, 1},
		160: {157, 1},
		161: {157, 3},
		162: {178, 0},
		163: {178, 1},
		164: {156, 1},
		165: {156, 3},
		166: {147, 2},
		167: {147, 2},
		168: {172, 1},
		169: {172, 3},
		170: {173, 0},
		171: {173, 1},
		172: {174, 0},
		173: {174, 1},
		174: {158, 0},
		175: {182, 3},
		176: {144, 1},
		177: {144, 2},
		178: {161, 0},
		179: {161, 1},
		180: {164, 3},
		181: {164, 4},
		182: {164, 5},
		183: {164, 6},
		184: {164, 6},
		185: {164, 4},
		186: {159, 0},
		187: {164, 4},
		188: {184, 0},
		189: {164, 5},
		190: {165, 0},
		191: {165, 1},
		192: {155, 1},
		193: {155, 4},
		194: {177, 2},
		195: {177, 4},
		196: {150, 2},
		197: {151, 0},
		198: {151, 1},
		199: {152, 1},
		200: {152, 2},
		201: {146, 3},
		202: {146, 2},
		203: {126, 1},
		204: {126, 1},
		205: {126, 1},
		206: {126, 1},
		207: {126, 1},
		208: {126, 1},
		209: {124, 3},
		210: {124, 4},
		211: {124, 3},
		212: {185, 0},
		213: {120, 4},
		214: {195, 1},
		215: {195, 2},
		216: {196, 0},
		217: {196, 1},
		218: {162, 1},
		219: {162, 1},
		220: {121, 2},
		221: {125, 5},
		222: {125, 7},
		223: {125, 5},
		224: {122, 5},
		225: {122, 7},
		226: {122, 9},
		227: {122, 8},
		228: {123, 3},
		229: {123, 2},
		230: {123, 2},
		231: {123, 3},
		232: {213, 1},
		233: {213, 2},
		234: {170, 1},
		235: {170, 1},
		236: {186, 0},
		237: {171, 5},
		238: {198, 1},
		239: {198, 2},
		240: {199, 0},
		241: {199, 1},
		242: {208, 1},
		243: {136, 1},
		244: {136, 2},
		245: {141, 0},
		246: {141, 1},
		247: {131, 1},
		248: {131, 1},
		249: {131, 3},
		250: {131, 1},
		251: {133, 4},
		252: {132, 4},
		253: {132, 4},
		254: {132, 4},
		255: {200, 1},
		256: {200, 2},
		257: {201, 0},
		258: {201, 1},
		259: {167, 4},
		260: {202, 3},
		261: {203, 0},
		262: {203, 1},
		263: {204, 1},
		264: {129, 3},
		265: {129, 5},
		266: {129, 7},
		267: {129, 5},
		268: {129, 2},
		269: {129, 1},
		270: {129, 3},
		271: {129, 3},
		272: {129, 2},
		273: {129, 3},
		274: {129, 6},
		275: {129, 8},
		276: {129, 2},
		277: {129, 4},
		278: {134, 1},
		279: {137, 1},
		280: {109, 1},
		281: {116, 1},
		282: {116, 2},
		283: {110, 1},
		284: {110, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 42}:   "invalid empty input",
		yyXError{488, -1}: "expected #endif",
		yyXError{490, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{406, -1}: "expected $end",
		yyXError{408, -1}: "expected $end",
		yyXError{339, -1}: "expected '('",
		yyXError{340, -1}: "expected '('",
		yyXError{341, -1}: "expected '('",
		yyXError{343, -1}: "expected '('",
		yyXError{369, -1}: "expected '('",
		yyXError{148, -1}: "expected ')'",
		yyXError{155, -1}: "expected ')'",
		yyXError{162, -1}: "expected ')'",
		yyXError{188, -1}: "expected ')'",
		yyXError{191, -1}: "expected ')'",
		yyXError{194, -1}: "expected ')'",
		yyXError{202, -1}: "expected ')'",
		yyXError{207, -1}: "expected ')'",
		yyXError{213, -1}: "expected ')'",
		yyXError{229, -1}: "expected ')'",
		yyXError{234, -1}: "expected ')'",
		yyXError{275, -1}: "expected ')'",
		yyXError{359, -1}: "expected ')'",
		yyXError{365, -1}: "expected ')'",
		yyXError{448, -1}: "expected ')'",
		yyXError{449, -1}: "expected ')'",
		yyXError{457, -1}: "expected ')'",
		yyXError{460, -1}: "expected ')'",
		yyXError{463, -1}: "expected ')'",
		yyXError{293, -1}: "expected ':'",
		yyXError{332, -1}: "expected ':'",
		yyXError{393, -1}: "expected ':'",
		yyXError{309, -1}: "expected ';'",
		yyXError{338, -1}: "expected ';'",
		yyXError{345, -1}: "expected ';'",
		yyXError{346, -1}: "expected ';'",
		yyXError{348, -1}: "expected ';'",
		yyXError{352, -1}: "expected ';'",
		yyXError{355, -1}: "expected ';'",
		yyXError{357, -1}: "expected ';'",
		yyXError{363, -1}: "expected ';'",
		yyXError{372, -1}: "expected ';'",
		yyXError{314, -1}: "expected '='",
		yyXError{167, -1}: "expected '['",
		yyXError{429, -1}: "expected '\\n'",
		yyXError{435, -1}: "expected '\\n'",
		yyXError{438, -1}: "expected '\\n'",
		yyXError{440, -1}: "expected '\\n'",
		yyXError{467, -1}: "expected '\\n'",
		yyXError{472, -1}: "expected '\\n'",
		yyXError{475, -1}: "expected '\\n'",
		yyXError{482, -1}: "expected '\\n'",
		yyXError{487, -1}: "expected '\\n'",
		yyXError{493, -1}: "expected '\\n'",
		yyXError{173, -1}: "expected ']'",
		yyXError{181, -1}: "expected ']'",
		yyXError{225, -1}: "expected ']'",
		yyXError{252, -1}: "expected ']'",
		yyXError{42, -1}:  "expected '{'",
		yyXError{44, -1}:  "expected '{'",
		yyXError{281, -1}: "expected '{'",
		yyXError{283, -1}: "expected '{'",
		yyXError{261, -1}: "expected '}'",
		yyXError{265, -1}: "expected '}'",
		yyXError{278, -1}: "expected '}'",
		yyXError{333, -1}: "expected '}'",
		yyXError{47, -1}:  "expected CommaOpt or one of [',', '}']",
		yyXError{244, -1}: "expected CommaOpt or one of [',', '}']",
		yyXError{259, -1}: "expected CommaOpt or one of [',', '}']",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{201, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{166, -1}: "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{335, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{315, -1}: "expected compound statement or '{'",
		yyXError{319, -1}: "expected compound statement or '{'",
		yyXError{312, -1}: "expected compound statement or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{50, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{249, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{297, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{331, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{405, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{317, -1}: "expected declaration or one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{354, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{296, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{193, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{246, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{196, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{163, -1}: "expected direct abstract declarator or one of ['(', '[']",
		yyXError{294, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{480, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{486, -1}: "expected endif line or #endif",
		yyXError{415, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{478, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{45, -1}:  "expected enumerator list or identifier",
		yyXError{277, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{73, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{97, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{370, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{374, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{378, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{382, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{60, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{71, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{241, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		yyXError{170, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{224, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{276, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{51, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{62, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{63, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{64, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{65, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{66, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{67, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{68, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{69, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{70, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{79, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{80, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{81, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{82, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{83, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{84, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{85, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{86, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{87, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{88, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{89, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{90, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{91, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{92, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{93, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{94, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{95, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{96, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{98, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{99, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{100, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{101, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{102, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{103, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{104, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{105, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{106, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{107, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{108, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{122, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{123, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{150, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{176, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{182, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{218, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{221, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{174, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{216, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{6, -1}:   "expected external declaration or one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{469, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{409, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{75, -1}:  "expected identifier",
		yyXError{76, -1}:  "expected identifier",
		yyXError{210, -1}: "expected identifier",
		yyXError{250, -1}: "expected identifier",
		yyXError{344, -1}: "expected identifier",
		yyXError{417, -1}: "expected identifier",
		yyXError{418, -1}: "expected identifier",
		yyXError{425, -1}: "expected identifier",
		yyXError{444, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{401, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{243, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{257, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{245, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{263, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{398, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{260, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{53, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{54, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{55, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{56, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{57, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{58, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{59, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{72, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{77, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{78, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{109, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{110, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{111, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{112, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{113, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{114, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{115, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{116, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{117, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{118, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{119, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{125, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{126, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{127, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{128, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{129, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{130, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{131, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{132, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{133, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{134, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{135, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{136, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{137, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{138, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{139, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{140, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{141, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{142, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{143, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{144, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{145, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{149, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{153, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{186, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{242, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{266, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{267, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{268, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{269, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{270, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{271, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{272, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{273, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{274, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{61, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{120, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{124, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{146, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{151, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{323, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{256, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{169, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{177, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{183, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{219, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{222, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{410, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{411, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{412, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{414, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{421, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{426, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{428, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{431, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{434, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{436, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{437, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{439, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{441, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{442, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{445, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{451, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{452, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{454, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{459, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{462, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{465, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{466, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{471, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{491, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{492, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{494, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{470, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{474, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{477, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{479, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{484, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{485, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{390, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{403, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{39, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{40, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{41, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{321, -1}: "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{404, -1}: "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{35, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{36, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{171, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{179, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{325, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{326, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{327, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{328, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{329, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{330, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{349, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{350, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{351, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{353, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{361, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{367, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{373, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{377, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{381, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{385, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{387, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{388, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{392, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{395, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{397, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{334, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{336, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{337, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{389, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{247, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{254, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{43, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{282, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{17, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{18, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{19, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{20, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{21, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{22, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{23, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{24, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{25, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{26, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{27, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{28, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{29, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{30, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{279, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{302, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{12, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{38, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{304, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{305, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{306, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{307, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{308, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{238, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{239, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{240, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{199, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{200, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{203, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{212, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{214, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{220, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{223, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{226, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{227, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{161, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{237, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{165, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{178, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{180, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{184, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{185, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{187, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{195, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{231, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{235, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{295, -1}: "expected one of ['(', identifier]",
		yyXError{324, -1}: "expected one of [')', ',', ';']",
		yyXError{446, -1}: "expected one of [')', ',', ...]",
		yyXError{456, -1}: "expected one of [')', ',', ...]",
		yyXError{147, -1}: "expected one of [')', ',']",
		yyXError{154, -1}: "expected one of [')', ',']",
		yyXError{164, -1}: "expected one of [')', ',']",
		yyXError{190, -1}: "expected one of [')', ',']",
		yyXError{192, -1}: "expected one of [')', ',']",
		yyXError{197, -1}: "expected one of [')', ',']",
		yyXError{198, -1}: "expected one of [')', ',']",
		yyXError{208, -1}: "expected one of [')', ',']",
		yyXError{209, -1}: "expected one of [')', ',']",
		yyXError{211, -1}: "expected one of [')', ',']",
		yyXError{230, -1}: "expected one of [')', ',']",
		yyXError{371, -1}: "expected one of [')', ',']",
		yyXError{375, -1}: "expected one of [')', ',']",
		yyXError{379, -1}: "expected one of [')', ',']",
		yyXError{383, -1}: "expected one of [')', ',']",
		yyXError{447, -1}: "expected one of [')', ',']",
		yyXError{292, -1}: "expected one of [',', ':', ';']",
		yyXError{121, -1}: "expected one of [',', ':']",
		yyXError{400, -1}: "expected one of [',', ';', '=']",
		yyXError{262, -1}: "expected one of [',', ';', '}']",
		yyXError{289, -1}: "expected one of [',', ';']",
		yyXError{291, -1}: "expected one of [',', ';']",
		yyXError{298, -1}: "expected one of [',', ';']",
		yyXError{301, -1}: "expected one of [',', ';']",
		yyXError{310, -1}: "expected one of [',', ';']",
		yyXError{311, -1}: "expected one of [',', ';']",
		yyXError{399, -1}: "expected one of [',', ';']",
		yyXError{402, -1}: "expected one of [',', ';']",
		yyXError{46, -1}:  "expected one of [',', '=', '}']",
		yyXError{49, -1}:  "expected one of [',', '=', '}']",
		yyXError{152, -1}: "expected one of [',', ']']",
		yyXError{48, -1}:  "expected one of [',', '}']",
		yyXError{52, -1}:  "expected one of [',', '}']",
		yyXError{258, -1}: "expected one of [',', '}']",
		yyXError{264, -1}: "expected one of [',', '}']",
		yyXError{280, -1}: "expected one of [',', '}']",
		yyXError{248, -1}: "expected one of ['.', '=', '[']",
		yyXError{251, -1}: "expected one of ['.', '=', '[']",
		yyXError{253, -1}: "expected one of ['.', '=', '[']",
		yyXError{255, -1}: "expected one of ['.', '=', '[']",
		yyXError{419, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		yyXError{427, -1}: "expected one of ['\\n', ppother]",
		yyXError{430, -1}: "expected one of ['\\n', ppother]",
		yyXError{432, -1}: "expected one of ['\\n', ppother]",
		yyXError{316, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{318, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{32, -1}:  "expected one of ['{', identifier]",
		yyXError{33, -1}:  "expected one of ['{', identifier]",
		yyXError{287, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{290, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{299, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{303, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{455, -1}: "expected one of [..., identifier]",
		yyXError{159, -1}: "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{74, -1}:  "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{320, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{322, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{8, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{358, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{364, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{347, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{356, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{362, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{215, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{204, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{168, -1}: "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{172, -1}: "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{468, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{473, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{476, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{483, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{489, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{205, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{31, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{34, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{313, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{189, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{232, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{233, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{157, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{158, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{420, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{424, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{160, -1}: "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{228, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{206, -1}: "expected parameter type list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{236, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{407, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{443, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{450, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{453, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{458, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{461, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{464, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{156, -1}: "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{342, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{360, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{366, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{376, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{380, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{384, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{386, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{391, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{394, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{396, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{284, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{285, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{286, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{288, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		yyXError{300, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{433, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{413, -1}: "expected token list or ppother",
		yyXError{416, -1}: "expected token list or ppother",
		yyXError{422, -1}: "expected token list or ppother",
		yyXError{423, -1}: "expected token list or ppother",
		yyXError{481, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{5, -1}:   "expected translation unit or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{175, -1}: "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{217, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{368, -1}: "expected while",
		yyXError{3, 42}:   "unexpected EOF",
		yyXError{2, 42}:   "unexpected EOF",
		yyXError{4, 42}:   "unexpected EOF",
	}

	yyParseTab = [495][]uint16{
		// 0
		{197: 288, 207: 287, 209: 286, 212: 289},
		{42: 285},
		{79: 284, 284, 95: 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 183: 692},
		{282, 282, 282, 282, 282, 282, 282, 282, 12: 282, 282, 282, 282, 282, 282, 282, 282, 282, 187: 690},
		{21: 280, 280, 280, 25: 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 41: 280, 43: 280, 280, 280, 280, 280, 188: 290},
		// 5
		{21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 106: 295, 111: 314, 316, 313, 294, 117: 292, 296, 293, 135: 326, 170: 324, 325, 213: 291},
		{21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 279, 300, 298, 323, 301, 297, 106: 295, 111: 314, 316, 313, 294, 117: 292, 296, 293, 135: 326, 170: 689, 325},
		{131, 445, 131, 8: 200, 127: 580, 579, 130: 597, 154: 595, 175: 596, 594},
		{204, 204, 204, 8: 204, 204, 204, 204, 21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 106: 295, 111: 314, 316, 313, 294, 117: 590, 296, 293, 145: 593},
		{204, 204, 204, 8: 204, 204, 204, 204, 21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 106: 295, 111: 314, 316, 313, 294, 117: 590, 296, 293, 145: 592},
		// 10
		{204, 204, 204, 8: 204, 204, 204, 204, 21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 106: 295, 111: 314, 316, 313, 294, 117: 590, 296, 293, 145: 591},
		{204, 204, 204, 8: 204, 204, 204, 204, 21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 106: 295, 111: 314, 316, 313, 294, 117: 590, 296, 293, 145: 589},
		{195, 195, 195, 8: 195, 195, 195, 195, 21: 195, 195, 195, 25: 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 41: 195, 43: 195, 195, 195, 195, 195},
		{194, 194, 194, 8: 194, 194, 194, 194, 21: 194, 194, 194, 25: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 41: 194, 43: 194, 194, 194, 194, 194},
		{193, 193, 193, 8: 193, 193, 193, 193, 21: 193, 193, 193, 25: 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 41: 193, 43: 193, 193, 193, 193, 193},
		// 15
		{192, 192, 192, 8: 192, 192, 192, 192, 21: 192, 192, 192, 25: 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 41: 192, 43: 192, 192, 192, 192, 192},
		{191, 191, 191, 8: 191, 191, 191, 191, 21: 191, 191, 191, 25: 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 41: 191, 43: 191, 191, 191, 191, 191},
		{190, 190, 190, 8: 190, 190, 190, 190, 21: 190, 190, 190, 25: 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 43: 190, 190, 190, 190, 190},
		{189, 189, 189, 8: 189, 189, 189, 189, 21: 189, 189, 189, 25: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 43: 189, 189, 189, 189, 189},
		{188, 188, 188, 8: 188, 188, 188, 188, 21: 188, 188, 188, 25: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 43: 188, 188, 188, 188, 188},
		// 20
		{187, 187, 187, 8: 187, 187, 187, 187, 21: 187, 187, 187, 25: 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 43: 187, 187, 187, 187, 187},
		{186, 186, 186, 8: 186, 186, 186, 186, 21: 186, 186, 186, 25: 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 43: 186, 186, 186, 186, 186},
		{185, 185, 185, 8: 185, 185, 185, 185, 21: 185, 185, 185, 25: 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 43: 185, 185, 185, 185, 185},
		{184, 184, 184, 8: 184, 184, 184, 184, 21: 184, 184, 184, 25: 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 43: 184, 184, 184, 184, 184},
		{183, 183, 183, 8: 183, 183, 183, 183, 21: 183, 183, 183, 25: 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 43: 183, 183, 183, 183, 183},
		// 25
		{182, 182, 182, 8: 182, 182, 182, 182, 21: 182, 182, 182, 25: 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 43: 182, 182, 182, 182, 182},
		{181, 181, 181, 8: 181, 181, 181, 181, 21: 181, 181, 181, 25: 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 43: 181, 181, 181, 181, 181},
		{180, 180, 180, 8: 180, 180, 180, 180, 21: 180, 180, 180, 25: 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 43: 180, 180, 180, 180, 180},
		{179, 179, 179, 8: 179, 179, 179, 179, 21: 179, 179, 179, 25: 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 43: 179, 179, 179, 179, 179},
		{178, 178, 178, 8: 178, 178, 178, 178, 21: 178, 178, 178, 25: 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 43: 178, 178, 178, 178, 178},
		// 30
		{177, 177, 177, 8: 177, 177, 177, 177, 21: 177, 177, 177, 25: 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 43: 177, 177, 177, 177, 177},
		{2: 567, 52: 113, 174: 566},
		{2: 172, 52: 172},
		{2: 171, 52: 171},
		{2: 328, 52: 113, 174: 327},
		// 35
		{149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 25: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 43: 149, 149, 149, 149, 149, 50: 149},
		{148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 25: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 43: 148, 148, 148, 148, 148, 50: 148},
		{147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 25: 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 43: 147, 147, 147, 147, 147, 50: 147},
		{146, 146, 146, 8: 146, 146, 146, 146, 21: 146, 146, 146, 25: 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 41: 146, 43: 146, 146, 146, 146, 146},
		{21: 53, 53, 53, 25: 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 41: 53, 53, 53, 53, 53, 53, 53},
		// 40
		{21: 51, 51, 51, 25: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 41: 51, 51, 51, 51, 51, 51, 51},
		{21: 50, 50, 50, 25: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 41: 50, 50, 50, 50, 50, 50, 50},
		{52: 156, 191: 329},
		{154, 154, 154, 8: 154, 154, 154, 154, 21: 154, 154, 154, 25: 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 43: 154, 154, 154, 154, 154, 52: 112},
		{52: 330},
		// 45
		{2: 331, 168: 334, 333, 205: 332},
		{10: 278, 24: 278, 48: 278},
		{10: 562, 24: 158, 149: 563},
		{10: 153, 24: 153},
		{10: 151, 24: 151, 48: 335},
		// 50
		{211, 211, 211, 211, 211, 211, 211, 211, 12: 211, 211, 211, 211, 211, 211, 211, 211, 211, 138: 336, 337},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 346},
		{10: 150, 24: 150},
		{273, 273, 3: 273, 273, 273, 273, 273, 273, 273, 273, 273, 24: 273, 40: 273, 42: 273, 48: 273, 50: 273, 273, 53: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273},
		{272, 272, 3: 272, 272, 272, 272, 272, 272, 272, 272, 272, 24: 272, 40: 272, 42: 272, 48: 272, 50: 272, 272, 53: 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272},
		// 55
		{271, 271, 3: 271, 271, 271, 271, 271, 271, 271, 271, 271, 24: 271, 40: 271, 42: 271, 48: 271, 50: 271, 271, 53: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{270, 270, 3: 270, 270, 270, 270, 270, 270, 270, 270, 270, 24: 270, 40: 270, 42: 270, 48: 270, 50: 270, 270, 53: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270},
		{269, 269, 3: 269, 269, 269, 269, 269, 269, 269, 269, 269, 24: 269, 40: 269, 42: 269, 48: 269, 50: 269, 269, 53: 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269},
		{268, 268, 3: 268, 268, 268, 268, 268, 268, 268, 268, 268, 24: 268, 40: 268, 42: 268, 48: 268, 50: 268, 268, 53: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268},
		{267, 267, 3: 267, 267, 267, 267, 267, 267, 267, 267, 267, 24: 267, 40: 267, 42: 267, 48: 267, 50: 267, 267, 53: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267},
		// 60
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 111, 111, 111, 25: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 49: 405, 107: 439, 158: 441, 182: 560},
		{359, 364, 3: 377, 367, 368, 363, 362, 210, 10: 210, 358, 24: 210, 40: 210, 42: 210, 48: 383, 50: 210, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 559},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 558},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 557},
		// 65
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 471},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 556},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 555},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 554},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 553},
		// 70
		{356, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 357},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 111, 111, 111, 25: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 49: 405, 107: 439, 158: 441, 182: 440},
		{359, 250, 3: 250, 250, 250, 363, 362, 250, 250, 250, 358, 24: 250, 40: 250, 42: 250, 48: 250, 50: 250, 360, 53: 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 361, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 107: 437},
		{345, 350, 338, 349, 351, 352, 348, 347, 9: 275, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 431, 193: 432, 433},
		// 75
		{2: 430},
		{2: 429},
		{261, 261, 3: 261, 261, 261, 261, 261, 261, 261, 261, 261, 24: 261, 40: 261, 42: 261, 48: 261, 50: 261, 261, 53: 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261},
		{260, 260, 3: 260, 260, 260, 260, 260, 260, 260, 260, 260, 24: 260, 40: 260, 42: 260, 48: 260, 50: 260, 260, 53: 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 428},
		// 80
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 427},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 426},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 425},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 424},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 423},
		// 85
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 422},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 421},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 420},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 419},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 418},
		// 90
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 417},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 416},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 415},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 414},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 413},
		// 95
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 412},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 411},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 107: 406},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 404},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 403},
		// 100
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 402},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 401},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 400},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 399},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 398},
		// 105
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 397},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 396},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 395},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 394},
		{359, 364, 3: 377, 367, 368, 363, 362, 218, 218, 218, 358, 24: 218, 40: 218, 42: 218, 48: 383, 50: 218, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		// 110
		{359, 364, 3: 377, 367, 368, 363, 362, 219, 219, 219, 358, 24: 219, 40: 219, 42: 219, 48: 383, 50: 219, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{359, 364, 3: 377, 367, 368, 363, 362, 220, 220, 220, 358, 24: 220, 40: 220, 42: 220, 48: 383, 50: 220, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{359, 364, 3: 377, 367, 368, 363, 362, 221, 221, 221, 358, 24: 221, 40: 221, 42: 221, 48: 383, 50: 221, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{359, 364, 3: 377, 367, 368, 363, 362, 222, 222, 222, 358, 24: 222, 40: 222, 42: 222, 48: 383, 50: 222, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{359, 364, 3: 377, 367, 368, 363, 362, 223, 223, 223, 358, 24: 223, 40: 223, 42: 223, 48: 383, 50: 223, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		// 115
		{359, 364, 3: 377, 367, 368, 363, 362, 224, 224, 224, 358, 24: 224, 40: 224, 42: 224, 48: 383, 50: 224, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{359, 364, 3: 377, 367, 368, 363, 362, 225, 225, 225, 358, 24: 225, 40: 225, 42: 225, 48: 383, 50: 225, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{359, 364, 3: 377, 367, 368, 363, 362, 226, 226, 226, 358, 24: 226, 40: 226, 42: 226, 48: 383, 50: 226, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{359, 364, 3: 377, 367, 368, 363, 362, 227, 227, 227, 358, 24: 227, 40: 227, 42: 227, 48: 383, 50: 227, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{359, 364, 3: 377, 367, 368, 363, 362, 228, 228, 228, 358, 24: 228, 40: 228, 42: 228, 48: 383, 50: 228, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		// 120
		{359, 364, 3: 377, 367, 368, 363, 362, 215, 215, 215, 358, 40: 215, 48: 383, 50: 215, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{10: 408, 40: 407},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 410},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 409},
		{359, 364, 3: 377, 367, 368, 363, 362, 214, 214, 214, 358, 40: 214, 48: 383, 50: 214, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		// 125
		{359, 364, 3: 377, 367, 368, 363, 362, 229, 229, 229, 358, 24: 229, 40: 229, 42: 229, 48: 229, 50: 229, 360, 53: 366, 365, 371, 372, 382, 378, 379, 229, 380, 229, 361, 229, 375, 374, 373, 369, 229, 229, 229, 376, 229, 381, 370, 229, 229, 229},
		{359, 364, 3: 377, 367, 368, 363, 362, 230, 230, 230, 358, 24: 230, 40: 230, 42: 230, 48: 230, 50: 230, 360, 53: 366, 365, 371, 372, 230, 378, 379, 230, 380, 230, 361, 230, 375, 374, 373, 369, 230, 230, 230, 376, 230, 230, 370, 230, 230, 230},
		{359, 364, 3: 377, 367, 368, 363, 362, 231, 231, 231, 358, 24: 231, 40: 231, 42: 231, 48: 231, 50: 231, 360, 53: 366, 365, 371, 372, 231, 378, 379, 231, 231, 231, 361, 231, 375, 374, 373, 369, 231, 231, 231, 376, 231, 231, 370, 231, 231, 231},
		{359, 364, 3: 377, 367, 368, 363, 362, 232, 232, 232, 358, 24: 232, 40: 232, 42: 232, 48: 232, 50: 232, 360, 53: 366, 365, 371, 372, 232, 378, 232, 232, 232, 232, 361, 232, 375, 374, 373, 369, 232, 232, 232, 376, 232, 232, 370, 232, 232, 232},
		{359, 364, 3: 377, 367, 368, 363, 362, 233, 233, 233, 358, 24: 233, 40: 233, 42: 233, 48: 233, 50: 233, 360, 53: 366, 365, 371, 372, 233, 233, 233, 233, 233, 233, 361, 233, 375, 374, 373, 369, 233, 233, 233, 376, 233, 233, 370, 233, 233, 233},
		// 130
		{359, 364, 3: 234, 367, 368, 363, 362, 234, 234, 234, 358, 24: 234, 40: 234, 42: 234, 48: 234, 50: 234, 360, 53: 366, 365, 371, 372, 234, 234, 234, 234, 234, 234, 361, 234, 375, 374, 373, 369, 234, 234, 234, 376, 234, 234, 370, 234, 234, 234},
		{359, 364, 3: 235, 367, 368, 363, 362, 235, 235, 235, 358, 24: 235, 40: 235, 42: 235, 48: 235, 50: 235, 360, 53: 366, 365, 371, 372, 235, 235, 235, 235, 235, 235, 361, 235, 235, 374, 373, 369, 235, 235, 235, 235, 235, 235, 370, 235, 235, 235},
		{359, 364, 3: 236, 367, 368, 363, 362, 236, 236, 236, 358, 24: 236, 40: 236, 42: 236, 48: 236, 50: 236, 360, 53: 366, 365, 371, 372, 236, 236, 236, 236, 236, 236, 361, 236, 236, 374, 373, 369, 236, 236, 236, 236, 236, 236, 370, 236, 236, 236},
		{359, 364, 3: 237, 367, 368, 363, 362, 237, 237, 237, 358, 24: 237, 40: 237, 42: 237, 48: 237, 50: 237, 360, 53: 366, 365, 237, 237, 237, 237, 237, 237, 237, 237, 361, 237, 237, 237, 237, 369, 237, 237, 237, 237, 237, 237, 370, 237, 237, 237},
		{359, 364, 3: 238, 367, 368, 363, 362, 238, 238, 238, 358, 24: 238, 40: 238, 42: 238, 48: 238, 50: 238, 360, 53: 366, 365, 238, 238, 238, 238, 238, 238, 238, 238, 361, 238, 238, 238, 238, 369, 238, 238, 238, 238, 238, 238, 370, 238, 238, 238},
		// 135
		{359, 364, 3: 239, 367, 368, 363, 362, 239, 239, 239, 358, 24: 239, 40: 239, 42: 239, 48: 239, 50: 239, 360, 53: 366, 365, 239, 239, 239, 239, 239, 239, 239, 239, 361, 239, 239, 239, 239, 369, 239, 239, 239, 239, 239, 239, 370, 239, 239, 239},
		{359, 364, 3: 240, 367, 368, 363, 362, 240, 240, 240, 358, 24: 240, 40: 240, 42: 240, 48: 240, 50: 240, 360, 53: 366, 365, 240, 240, 240, 240, 240, 240, 240, 240, 361, 240, 240, 240, 240, 369, 240, 240, 240, 240, 240, 240, 370, 240, 240, 240},
		{359, 364, 3: 241, 367, 368, 363, 362, 241, 241, 241, 358, 24: 241, 40: 241, 42: 241, 48: 241, 50: 241, 360, 53: 366, 365, 241, 241, 241, 241, 241, 241, 241, 241, 361, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241},
		{359, 364, 3: 242, 367, 368, 363, 362, 242, 242, 242, 358, 24: 242, 40: 242, 42: 242, 48: 242, 50: 242, 360, 53: 366, 365, 242, 242, 242, 242, 242, 242, 242, 242, 361, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242},
		{359, 364, 3: 243, 243, 243, 363, 362, 243, 243, 243, 358, 24: 243, 40: 243, 42: 243, 48: 243, 50: 243, 360, 53: 366, 365, 243, 243, 243, 243, 243, 243, 243, 243, 361, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		// 140
		{359, 364, 3: 244, 244, 244, 363, 362, 244, 244, 244, 358, 24: 244, 40: 244, 42: 244, 48: 244, 50: 244, 360, 53: 366, 365, 244, 244, 244, 244, 244, 244, 244, 244, 361, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244},
		{359, 245, 3: 245, 245, 245, 363, 362, 245, 245, 245, 358, 24: 245, 40: 245, 42: 245, 48: 245, 50: 245, 360, 53: 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 361, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245},
		{359, 246, 3: 246, 246, 246, 363, 362, 246, 246, 246, 358, 24: 246, 40: 246, 42: 246, 48: 246, 50: 246, 360, 53: 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 361, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246},
		{359, 247, 3: 247, 247, 247, 363, 362, 247, 247, 247, 358, 24: 247, 40: 247, 42: 247, 48: 247, 50: 247, 360, 53: 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 361, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247},
		{262, 262, 3: 262, 262, 262, 262, 262, 262, 262, 262, 262, 24: 262, 40: 262, 42: 262, 48: 262, 50: 262, 262, 53: 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262},
		// 145
		{263, 263, 3: 263, 263, 263, 263, 263, 263, 263, 263, 263, 24: 263, 40: 263, 42: 263, 48: 263, 50: 263, 263, 53: 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263},
		{359, 364, 3: 377, 367, 368, 363, 362, 9: 277, 277, 358, 48: 383, 51: 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{9: 274, 435},
		{9: 434},
		{264, 264, 3: 264, 264, 264, 264, 264, 264, 264, 264, 264, 24: 264, 40: 264, 42: 264, 48: 264, 50: 264, 264, 53: 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264},
		// 150
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 436},
		{359, 364, 3: 377, 367, 368, 363, 362, 9: 276, 276, 358, 48: 383, 51: 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{10: 408, 50: 438},
		{265, 265, 3: 265, 265, 265, 265, 265, 265, 265, 265, 265, 24: 265, 40: 265, 42: 265, 48: 265, 50: 265, 265, 53: 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265},
		{9: 552, 408},
		// 155
		{9: 526},
		{21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 106: 443, 111: 314, 316, 313, 442, 142: 444},
		{164, 164, 164, 8: 164, 164, 11: 164, 21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 164, 106: 443, 111: 314, 316, 313, 442, 142: 524, 179: 525},
		{164, 164, 164, 8: 164, 164, 11: 164, 21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 164, 106: 443, 111: 314, 316, 313, 442, 142: 524, 179: 523},
		{131, 445, 9: 107, 11: 131, 127: 446, 448, 144: 449, 161: 447},
		// 160
		{127, 127, 127, 9: 127, 127, 127, 21: 320, 321, 322, 106: 456, 143: 460, 148: 521},
		{130, 2: 130, 9: 109, 109, 130},
		{9: 110},
		{451, 11: 95, 164: 450, 452},
		{9: 106, 106},
		// 165
		{517, 9: 108, 108, 94},
		{131, 445, 9: 99, 11: 131, 21: 99, 99, 99, 25: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 41: 99, 43: 99, 99, 99, 99, 99, 127: 446, 448, 144: 473, 159: 474},
		{11: 453},
		{345, 455, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 320, 321, 322, 41: 459, 49: 454, 217, 106: 456, 143: 457, 153: 458},
		{359, 364, 3: 377, 367, 368, 363, 362, 11: 358, 48: 383, 50: 216, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		// 170
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 471, 472},
		{129, 129, 129, 129, 129, 129, 129, 129, 9: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 41: 129, 50: 129},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 320, 321, 322, 41: 467, 49: 454, 217, 106: 464, 153: 466},
		{50: 465},
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 320, 321, 322, 106: 456, 143: 460, 148: 461},
		// 175
		{126, 126, 126, 126, 126, 126, 126, 126, 9: 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 320, 321, 322, 106: 464},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 462},
		{359, 364, 3: 377, 367, 368, 363, 362, 11: 358, 48: 383, 50: 463, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{102, 9: 102, 102, 102},
		{128, 128, 128, 128, 128, 128, 128, 128, 9: 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 41: 128, 50: 128},
		// 180
		{104, 9: 104, 104, 104},
		{50: 470},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 468},
		{359, 364, 3: 377, 367, 368, 363, 362, 11: 358, 48: 383, 50: 469, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{101, 9: 101, 101, 101},
		// 185
		{103, 9: 103, 103, 103},
		{359, 255, 3: 255, 255, 255, 363, 362, 255, 255, 255, 358, 24: 255, 40: 255, 42: 255, 48: 255, 50: 255, 360, 53: 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 361, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{100, 9: 100, 100, 100},
		{9: 516},
		{9: 123, 21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 106: 295, 111: 314, 316, 313, 294, 117: 478, 296, 293, 147: 477, 156: 475, 476, 178: 479},
		// 190
		{9: 125, 513},
		{9: 122},
		{9: 121, 121},
		{131, 445, 131, 9: 107, 107, 131, 127: 446, 481, 130: 482, 144: 449, 161: 483},
		{9: 480},
		// 195
		{98, 9: 98, 98, 98},
		{486, 2: 485, 11: 95, 164: 450, 452, 484},
		{9: 119, 119},
		{9: 118, 118},
		{490, 8: 145, 145, 145, 489, 21: 145, 145, 145, 25: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 43: 145, 145, 145, 145, 145, 145, 52: 145},
		// 200
		{142, 8: 142, 142, 142, 142, 21: 142, 142, 142, 25: 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 43: 142, 142, 142, 142, 142, 142, 52: 142},
		{131, 445, 131, 9: 99, 11: 131, 21: 99, 99, 99, 25: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 41: 99, 43: 99, 99, 99, 99, 99, 127: 446, 481, 130: 487, 144: 473, 159: 474},
		{9: 488},
		{141, 8: 141, 141, 141, 141, 21: 141, 141, 141, 25: 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 43: 141, 141, 141, 141, 141, 141, 52: 141},
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 320, 321, 322, 41: 501, 50: 127, 106: 456, 143: 502, 148: 500},
		// 205
		{2: 493, 9: 115, 21: 136, 136, 136, 25: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 41: 136, 43: 136, 136, 136, 136, 136, 172: 494, 492, 192: 491},
		{21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 106: 295, 111: 314, 316, 313, 294, 117: 478, 296, 293, 147: 477, 156: 475, 498},
		{9: 497},
		{9: 117, 117},
		{9: 114, 495},
		// 210
		{2: 496},
		{9: 116, 116},
		{134, 8: 134, 134, 134, 134, 21: 134, 134, 134, 25: 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 43: 134, 134, 134, 134, 134, 134, 52: 134},
		{9: 499},
		{135, 8: 135, 135, 135, 135, 21: 135, 135, 135, 25: 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 43: 135, 135, 135, 135, 135, 135, 52: 135},
		// 215
		{345, 509, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 454, 217, 153: 510},
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 320, 321, 322, 106: 456, 143: 460, 148: 506},
		{126, 126, 126, 126, 126, 126, 126, 126, 12: 126, 126, 126, 126, 126, 126, 126, 126, 126, 320, 321, 322, 41: 503, 50: 126, 106: 464},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 504},
		{359, 364, 3: 377, 367, 368, 363, 362, 11: 358, 48: 383, 50: 505, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		// 220
		{138, 8: 138, 138, 138, 138, 21: 138, 138, 138, 25: 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 43: 138, 138, 138, 138, 138, 138, 52: 138},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 507},
		{359, 364, 3: 377, 367, 368, 363, 362, 11: 358, 48: 383, 50: 508, 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{139, 8: 139, 139, 139, 139, 21: 139, 139, 139, 25: 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 43: 139, 139, 139, 139, 139, 139, 52: 139},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 471, 512},
		// 225
		{50: 511},
		{140, 8: 140, 140, 140, 140, 21: 140, 140, 140, 25: 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 43: 140, 140, 140, 140, 140, 140, 52: 140},
		{137, 8: 137, 137, 137, 137, 21: 137, 137, 137, 25: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 43: 137, 137, 137, 137, 137, 137, 52: 137},
		{21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 106: 295, 111: 314, 316, 313, 294, 117: 478, 296, 293, 140: 514, 147: 515},
		{9: 124},
		// 230
		{9: 120, 120},
		{105, 9: 105, 105, 105},
		{9: 97, 21: 97, 97, 97, 25: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 41: 97, 43: 97, 97, 97, 97, 97, 184: 518},
		{9: 123, 21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 106: 295, 111: 314, 316, 313, 294, 117: 478, 296, 293, 147: 477, 156: 475, 476, 178: 519},
		{9: 520},
		// 235
		{96, 9: 96, 96, 96},
		{133, 445, 133, 9: 133, 133, 133, 127: 522},
		{132, 2: 132, 9: 132, 132, 132},
		{165, 165, 165, 8: 165, 165, 11: 165, 40: 165},
		{163, 163, 163, 8: 163, 163, 11: 163, 40: 163},
		// 240
		{166, 166, 166, 8: 166, 166, 11: 166, 40: 166},
		{345, 249, 338, 249, 249, 249, 348, 347, 249, 249, 249, 249, 354, 353, 339, 340, 341, 342, 343, 355, 344, 24: 249, 40: 249, 42: 249, 48: 249, 527, 249, 249, 528, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249},
		{359, 248, 3: 248, 248, 248, 363, 362, 248, 248, 248, 358, 24: 248, 40: 248, 42: 248, 48: 248, 50: 248, 360, 53: 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 361, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248},
		{88, 88, 88, 88, 88, 88, 88, 88, 11: 534, 88, 88, 88, 88, 88, 88, 88, 88, 88, 51: 535, 88, 146: 533, 150: 532, 530, 531, 177: 529},
		{10: 545, 24: 158, 149: 550},
		// 245
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 541, 52: 542, 155: 543},
		{11: 534, 48: 539, 51: 535, 146: 540},
		{87, 87, 87, 87, 87, 87, 87, 87, 12: 87, 87, 87, 87, 87, 87, 87, 87, 87, 52: 87},
		{11: 86, 48: 86, 51: 86},
		{211, 211, 211, 211, 211, 211, 211, 211, 12: 211, 211, 211, 211, 211, 211, 211, 211, 211, 138: 336, 537},
		// 250
		{2: 536},
		{11: 83, 48: 83, 51: 83},
		{50: 538},
		{11: 84, 48: 84, 51: 84},
		{89, 89, 89, 89, 89, 89, 89, 89, 12: 89, 89, 89, 89, 89, 89, 89, 89, 89, 52: 89},
		// 255
		{11: 85, 48: 85, 51: 85},
		{359, 364, 3: 377, 367, 368, 363, 362, 93, 10: 93, 358, 24: 93, 48: 383, 51: 360, 53: 366, 365, 371, 372, 382, 378, 379, 387, 380, 391, 361, 385, 375, 374, 373, 369, 389, 386, 384, 376, 393, 381, 370, 390, 388, 392},
		{88, 88, 88, 88, 88, 88, 88, 88, 11: 534, 88, 88, 88, 88, 88, 88, 88, 88, 88, 51: 535, 88, 146: 533, 150: 532, 530, 531, 177: 544},
		{10: 91, 24: 91},
		{10: 545, 24: 158, 149: 546},
		// 260
		{88, 88, 88, 88, 88, 88, 88, 88, 11: 534, 88, 88, 88, 88, 88, 88, 88, 88, 88, 24: 157, 51: 535, 88, 146: 533, 150: 532, 548, 531},
		{24: 547},
		{8: 92, 10: 92, 24: 92},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 541, 52: 542, 155: 549},
		{10: 90, 24: 90},
		// 265
		{24: 551},
		{259, 259, 3: 259, 259, 259, 259, 259, 259, 259, 259, 259, 24: 259, 40: 259, 42: 259, 48: 259, 50: 259, 259, 53: 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259},
		{266, 266, 3: 266, 266, 266, 266, 266, 266, 266, 266, 266, 24: 266, 40: 266, 42: 266, 48: 266, 50: 266, 266, 53: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266},
		{359, 251, 3: 251, 251, 251, 363, 362, 251, 251, 251, 358, 24: 251, 40: 251, 42: 251, 48: 251, 50: 251, 360, 53: 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 361, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251},
		{359, 252, 3: 252, 252, 252, 363, 362, 252, 252, 252, 358, 24: 252, 40: 252, 42: 252, 48: 252, 50: 252, 360, 53: 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 361, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252},
		// 270
		{359, 253, 3: 253, 253, 253, 363, 362, 253, 253, 253, 358, 24: 253, 40: 253, 42: 253, 48: 253, 50: 253, 360, 53: 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 361, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253},
		{359, 254, 3: 254, 254, 254, 363, 362, 254, 254, 254, 358, 24: 254, 40: 254, 42: 254, 48: 254, 50: 254, 360, 53: 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 361, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254},
		{359, 256, 3: 256, 256, 256, 363, 362, 256, 256, 256, 358, 24: 256, 40: 256, 42: 256, 48: 256, 50: 256, 360, 53: 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 361, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256},
		{359, 257, 3: 257, 257, 257, 363, 362, 257, 257, 257, 358, 24: 257, 40: 257, 42: 257, 48: 257, 50: 257, 360, 53: 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 361, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257},
		{359, 258, 3: 258, 258, 258, 363, 362, 258, 258, 258, 358, 24: 258, 40: 258, 42: 258, 48: 258, 50: 258, 360, 53: 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 361, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258},
		// 275
		{9: 561},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 527, 52: 528},
		{2: 331, 24: 157, 168: 334, 565},
		{24: 564},
		{155, 155, 155, 8: 155, 155, 155, 155, 21: 155, 155, 155, 25: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 43: 155, 155, 155, 155, 155},
		// 280
		{10: 152, 24: 152},
		{52: 176, 189: 568},
		{173, 173, 173, 8: 173, 173, 173, 173, 21: 173, 173, 173, 25: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 43: 173, 173, 173, 173, 173, 52: 112},
		{52: 569},
		{21: 175, 175, 175, 25: 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 190: 570},
		// 285
		{21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 106: 443, 111: 314, 316, 313, 442, 142: 573, 180: 572, 210: 571},
		{21: 320, 321, 322, 587, 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 106: 443, 111: 314, 316, 313, 442, 142: 573, 180: 588},
		{21: 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170},
		{131, 445, 131, 8: 575, 40: 144, 127: 580, 579, 130: 577, 163: 578, 181: 576, 211: 574},
		{8: 584, 10: 585},
		// 290
		{21: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167},
		{8: 162, 10: 162},
		{8: 160, 10: 160, 40: 143},
		{40: 582},
		{581, 2: 485, 166: 484},
		// 295
		{130, 2: 130},
		{131, 445, 131, 127: 580, 579, 130: 487},
		{211, 211, 211, 211, 211, 211, 211, 211, 12: 211, 211, 211, 211, 211, 211, 211, 211, 211, 138: 336, 583},
		{8: 159, 10: 159},
		{21: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168},
		// 300
		{131, 445, 131, 40: 144, 127: 580, 579, 130: 577, 163: 578, 181: 586},
		{8: 161, 10: 161},
		{174, 174, 174, 8: 174, 174, 174, 174, 21: 174, 174, 174, 25: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 43: 174, 174, 174, 174, 174},
		{21: 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169},
		{205, 205, 205, 8: 205, 205, 205, 205},
		// 305
		{203, 203, 203, 8: 203, 203, 203, 203},
		{206, 206, 206, 8: 206, 206, 206, 206},
		{207, 207, 207, 8: 207, 207, 207, 207},
		{208, 208, 208, 8: 208, 208, 208, 208},
		{8: 688},
		// 310
		{8: 202, 10: 202},
		{8: 199, 10: 686},
		{8: 198, 10: 198, 21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 197, 52: 45, 106: 295, 111: 314, 316, 313, 294, 117: 598, 296, 293, 135: 601, 160: 599, 198: 602, 600},
		{131, 445, 131, 8: 200, 127: 580, 579, 130: 685, 154: 595, 175: 596, 594},
		{48: 683},
		// 315
		{52: 49, 186: 604},
		{21: 47, 47, 47, 25: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 41: 47, 43: 47, 47, 47, 47, 47, 52: 47},
		{21: 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 52: 44, 106: 295, 111: 314, 316, 313, 294, 117: 598, 296, 293, 135: 603},
		{21: 46, 46, 46, 25: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 41: 46, 43: 46, 46, 46, 46, 46, 52: 46},
		{52: 605, 120: 606},
		// 320
		{73, 73, 73, 73, 73, 73, 73, 73, 73, 12: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 41: 73, 43: 73, 73, 73, 73, 73, 52: 73, 82: 73, 73, 73, 73, 73, 73, 73, 73, 73, 92: 73, 73, 185: 607},
		{21: 48, 48, 48, 25: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 41: 48, 48, 48, 48, 48, 48, 48},
		{345, 350, 608, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 320, 321, 322, 69, 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 49: 405, 52: 605, 82: 626, 631, 616, 630, 617, 627, 628, 629, 624, 92: 632, 625, 106: 295, 609, 111: 314, 316, 313, 294, 623, 117: 598, 296, 293, 611, 612, 614, 615, 610, 613, 622, 135: 621, 162: 619, 195: 620, 618},
		{273, 273, 3: 273, 273, 273, 273, 273, 273, 10: 273, 273, 40: 681, 48: 273, 51: 273, 53: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273},
		{8: 212, 212, 408},
		// 325
		{82, 82, 82, 82, 82, 82, 82, 82, 82, 12: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 41: 82, 43: 82, 82, 82, 82, 82, 52: 82, 82: 82, 82, 82, 82, 82, 82, 82, 82, 82, 92: 82, 82, 108: 82},
		{81, 81, 81, 81, 81, 81, 81, 81, 81, 12: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 41: 81, 43: 81, 81, 81, 81, 81, 52: 81, 82: 81, 81, 81, 81, 81, 81, 81, 81, 81, 92: 81, 81, 108: 81},
		{80, 80, 80, 80, 80, 80, 80, 80, 80, 12: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 41: 80, 43: 80, 80, 80, 80, 80, 52: 80, 82: 80, 80, 80, 80, 80, 80, 80, 80, 80, 92: 80, 80, 108: 80},
		{79, 79, 79, 79, 79, 79, 79, 79, 79, 12: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 41: 79, 43: 79, 79, 79, 79, 79, 52: 79, 82: 79, 79, 79, 79, 79, 79, 79, 79, 79, 92: 79, 79, 108: 79},
		{78, 78, 78, 78, 78, 78, 78, 78, 78, 12: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 41: 78, 43: 78, 78, 78, 78, 78, 52: 78, 82: 78, 78, 78, 78, 78, 78, 78, 78, 78, 92: 78, 78, 108: 78},
		// 330
		{77, 77, 77, 77, 77, 77, 77, 77, 77, 12: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 41: 77, 43: 77, 77, 77, 77, 77, 52: 77, 82: 77, 77, 77, 77, 77, 77, 77, 77, 77, 92: 77, 77, 108: 77},
		{211, 211, 211, 211, 211, 211, 211, 211, 12: 211, 211, 211, 211, 211, 211, 211, 211, 211, 138: 336, 678},
		{40: 676},
		{24: 675},
		{71, 71, 71, 71, 71, 71, 71, 71, 71, 12: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 41: 71, 43: 71, 71, 71, 71, 71, 52: 71, 82: 71, 71, 71, 71, 71, 71, 71, 71, 71, 92: 71, 71},
		// 335
		{345, 350, 608, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 320, 321, 322, 68, 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 49: 405, 52: 605, 82: 626, 631, 616, 630, 617, 627, 628, 629, 624, 92: 632, 625, 106: 295, 609, 111: 314, 316, 313, 294, 623, 117: 598, 296, 293, 611, 612, 614, 615, 610, 613, 622, 135: 621, 162: 674},
		{67, 67, 67, 67, 67, 67, 67, 67, 67, 12: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 41: 67, 43: 67, 67, 67, 67, 67, 52: 67, 82: 67, 67, 67, 67, 67, 67, 67, 67, 67, 92: 67, 67},
		{66, 66, 66, 66, 66, 66, 66, 66, 66, 12: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 41: 66, 43: 66, 66, 66, 66, 66, 52: 66, 82: 66, 66, 66, 66, 66, 66, 66, 66, 66, 92: 66, 66},
		{8: 673},
		{667},
		// 340
		{663},
		{659},
		{345, 350, 608, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 52: 605, 82: 626, 631, 616, 630, 617, 627, 628, 629, 624, 92: 632, 625, 107: 609, 115: 623, 120: 611, 612, 614, 615, 610, 613, 653},
		{639},
		{2: 637},
		// 345
		{8: 636},
		{8: 635},
		{345, 350, 338, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 107: 609, 115: 633},
		{8: 634},
		{54, 54, 54, 54, 54, 54, 54, 54, 54, 12: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 41: 54, 43: 54, 54, 54, 54, 54, 52: 54, 82: 54, 54, 54, 54, 54, 54, 54, 54, 54, 92: 54, 54, 108: 54},
		// 350
		{55, 55, 55, 55, 55, 55, 55, 55, 55, 12: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 41: 55, 43: 55, 55, 55, 55, 55, 52: 55, 82: 55, 55, 55, 55, 55, 55, 55, 55, 55, 92: 55, 55, 108: 55},
		{56, 56, 56, 56, 56, 56, 56, 56, 56, 12: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 41: 56, 43: 56, 56, 56, 56, 56, 52: 56, 82: 56, 56, 56, 56, 56, 56, 56, 56, 56, 92: 56, 56, 108: 56},
		{8: 638},
		{57, 57, 57, 57, 57, 57, 57, 57, 57, 12: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 41: 57, 43: 57, 57, 57, 57, 57, 52: 57, 82: 57, 57, 57, 57, 57, 57, 57, 57, 57, 92: 57, 57, 108: 57},
		{345, 350, 338, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 320, 321, 322, 25: 311, 303, 312, 308, 319, 307, 305, 306, 304, 309, 317, 315, 318, 310, 302, 41: 299, 43: 300, 298, 323, 301, 297, 49: 405, 106: 295, 609, 111: 314, 316, 313, 294, 640, 117: 598, 296, 293, 135: 641},
		// 355
		{8: 647},
		{345, 350, 338, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 107: 609, 115: 642},
		{8: 643},
		{345, 350, 338, 349, 351, 352, 348, 347, 9: 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 107: 609, 115: 644},
		{9: 645},
		// 360
		{345, 350, 608, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 52: 605, 82: 626, 631, 616, 630, 617, 627, 628, 629, 624, 92: 632, 625, 107: 609, 115: 623, 120: 611, 612, 614, 615, 610, 613, 646},
		{58, 58, 58, 58, 58, 58, 58, 58, 58, 12: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 41: 58, 43: 58, 58, 58, 58, 58, 52: 58, 82: 58, 58, 58, 58, 58, 58, 58, 58, 58, 92: 58, 58, 108: 58},
		{345, 350, 338, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 107: 609, 115: 648},
		{8: 649},
		{345, 350, 338, 349, 351, 352, 348, 347, 9: 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 107: 609, 115: 650},
		// 365
		{9: 651},
		{345, 350, 608, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 52: 605, 82: 626, 631, 616, 630, 617, 627, 628, 629, 624, 92: 632, 625, 107: 609, 115: 623, 120: 611, 612, 614, 615, 610, 613, 652},
		{59, 59, 59, 59, 59, 59, 59, 59, 59, 12: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 41: 59, 43: 59, 59, 59, 59, 59, 52: 59, 82: 59, 59, 59, 59, 59, 59, 59, 59, 59, 92: 59, 59, 108: 59},
		{82: 654},
		{655},
		// 370
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 107: 656},
		{9: 657, 408},
		{8: 658},
		{60, 60, 60, 60, 60, 60, 60, 60, 60, 12: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 41: 60, 43: 60, 60, 60, 60, 60, 52: 60, 82: 60, 60, 60, 60, 60, 60, 60, 60, 60, 92: 60, 60, 108: 60},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 107: 660},
		// 375
		{9: 661, 408},
		{345, 350, 608, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 52: 605, 82: 626, 631, 616, 630, 617, 627, 628, 629, 624, 92: 632, 625, 107: 609, 115: 623, 120: 611, 612, 614, 615, 610, 613, 662},
		{61, 61, 61, 61, 61, 61, 61, 61, 61, 12: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 41: 61, 43: 61, 61, 61, 61, 61, 52: 61, 82: 61, 61, 61, 61, 61, 61, 61, 61, 61, 92: 61, 61, 108: 61},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 107: 664},
		{9: 665, 408},
		// 380
		{345, 350, 608, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 52: 605, 82: 626, 631, 616, 630, 617, 627, 628, 629, 624, 92: 632, 625, 107: 609, 115: 623, 120: 611, 612, 614, 615, 610, 613, 666},
		{62, 62, 62, 62, 62, 62, 62, 62, 62, 12: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 41: 62, 43: 62, 62, 62, 62, 62, 52: 62, 82: 62, 62, 62, 62, 62, 62, 62, 62, 62, 92: 62, 62, 108: 62},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 107: 668},
		{9: 669, 408},
		{345, 350, 608, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 52: 605, 82: 626, 631, 616, 630, 617, 627, 628, 629, 624, 92: 632, 625, 107: 609, 115: 623, 120: 611, 612, 614, 615, 610, 613, 670},
		// 385
		{64, 64, 64, 64, 64, 64, 64, 64, 64, 12: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 41: 64, 43: 64, 64, 64, 64, 64, 52: 64, 82: 64, 64, 64, 64, 64, 64, 64, 64, 64, 92: 64, 64, 108: 671},
		{345, 350, 608, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 52: 605, 82: 626, 631, 616, 630, 617, 627, 628, 629, 624, 92: 632, 625, 107: 609, 115: 623, 120: 611, 612, 614, 615, 610, 613, 672},
		{63, 63, 63, 63, 63, 63, 63, 63, 63, 12: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 41: 63, 43: 63, 63, 63, 63, 63, 52: 63, 82: 63, 63, 63, 63, 63, 63, 63, 63, 63, 92: 63, 63, 108: 63},
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 12: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 41: 65, 43: 65, 65, 65, 65, 65, 52: 65, 82: 65, 65, 65, 65, 65, 65, 65, 65, 65, 92: 65, 65, 108: 65},
		{70, 70, 70, 70, 70, 70, 70, 70, 70, 12: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 41: 70, 43: 70, 70, 70, 70, 70, 52: 70, 82: 70, 70, 70, 70, 70, 70, 70, 70, 70, 92: 70, 70},
		// 390
		{72, 72, 72, 72, 72, 72, 72, 72, 72, 12: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 41: 72, 72, 72, 72, 72, 72, 72, 52: 72, 82: 72, 72, 72, 72, 72, 72, 72, 72, 72, 92: 72, 72, 108: 72},
		{345, 350, 608, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 52: 605, 82: 626, 631, 616, 630, 617, 627, 628, 629, 624, 92: 632, 625, 107: 609, 115: 623, 120: 611, 612, 614, 615, 610, 613, 677},
		{74, 74, 74, 74, 74, 74, 74, 74, 74, 12: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 41: 74, 43: 74, 74, 74, 74, 74, 52: 74, 82: 74, 74, 74, 74, 74, 74, 74, 74, 74, 92: 74, 74, 108: 74},
		{40: 679},
		{345, 350, 608, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 52: 605, 82: 626, 631, 616, 630, 617, 627, 628, 629, 624, 92: 632, 625, 107: 609, 115: 623, 120: 611, 612, 614, 615, 610, 613, 680},
		// 395
		{75, 75, 75, 75, 75, 75, 75, 75, 75, 12: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 41: 75, 43: 75, 75, 75, 75, 75, 52: 75, 82: 75, 75, 75, 75, 75, 75, 75, 75, 75, 92: 75, 75, 108: 75},
		{345, 350, 608, 349, 351, 352, 348, 347, 213, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 405, 52: 605, 82: 626, 631, 616, 630, 617, 627, 628, 629, 624, 92: 632, 625, 107: 609, 115: 623, 120: 611, 612, 614, 615, 610, 613, 682},
		{76, 76, 76, 76, 76, 76, 76, 76, 76, 12: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 41: 76, 43: 76, 76, 76, 76, 76, 52: 76, 82: 76, 76, 76, 76, 76, 76, 76, 76, 76, 92: 76, 76, 108: 76},
		{345, 350, 338, 349, 351, 352, 348, 347, 12: 354, 353, 339, 340, 341, 342, 343, 355, 344, 49: 541, 52: 542, 155: 684},
		{8: 196, 10: 196},
		// 400
		{8: 198, 10: 198, 48: 197, 160: 599},
		{131, 445, 131, 127: 580, 579, 130: 685, 154: 687},
		{8: 201, 10: 201},
		{209, 209, 209, 209, 209, 209, 209, 209, 209, 12: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 41: 209, 209, 209, 209, 209, 209, 209, 52: 209, 82: 209, 209, 209, 209, 209, 209, 209, 209, 209, 92: 209, 209},
		{21: 52, 52, 52, 25: 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 41: 52, 52, 52, 52, 52, 52, 52},
		// 405
		{211, 211, 211, 211, 211, 211, 211, 211, 12: 211, 211, 211, 211, 211, 211, 211, 211, 211, 138: 336, 691},
		{42: 281},
		{79: 713, 715, 95: 704, 705, 706, 701, 702, 703, 707, 708, 698, 709, 710, 109: 714, 712, 116: 711, 129: 696, 131: 695, 700, 697, 699, 136: 694, 208: 693},
		{42: 283},
		{42: 43, 79: 713, 715, 95: 704, 705, 706, 701, 702, 703, 707, 708, 698, 709, 710, 109: 714, 712, 116: 711, 129: 696, 131: 756, 700, 697, 699},
		// 410
		{42: 42, 79: 42, 42, 42, 91: 42, 94: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{42: 38, 79: 38, 38, 38, 91: 38, 94: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{42: 37, 79: 37, 37, 37, 91: 37, 94: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{80: 715, 109: 778, 712},
		{42: 35, 79: 35, 35, 35, 91: 35, 94: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		// 415
		{81: 28, 91: 28, 94: 766, 167: 764, 200: 765, 763},
		{80: 715, 109: 760, 712},
		{2: 757},
		{2: 752},
		{2: 728, 79: 730, 206: 729},
		// 420
		{79: 713, 715, 109: 714, 712, 116: 727},
		{42: 16, 79: 16, 16, 16, 91: 16, 94: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{80: 715, 109: 725, 712},
		{80: 715, 109: 723, 712},
		{79: 713, 715, 109: 714, 712, 116: 722},
		// 425
		{2: 718},
		{42: 7, 79: 7, 7, 7, 91: 7, 94: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{79: 5, 717},
		{42: 4, 79: 4, 4, 4, 91: 4, 94: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{79: 716},
		// 430
		{79: 2, 2},
		{42: 3, 79: 3, 3, 3, 91: 3, 94: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{79: 1, 1},
		{79: 719, 715, 109: 720, 712},
		{42: 12, 79: 12, 12, 12, 91: 12, 94: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		// 435
		{79: 721},
		{42: 8, 79: 8, 8, 8, 91: 8, 94: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{42: 13, 79: 13, 13, 13, 91: 13, 94: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{79: 724},
		{42: 14, 79: 14, 14, 14, 91: 14, 94: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		// 440
		{79: 726},
		{42: 15, 79: 15, 15, 15, 91: 15, 94: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{42: 17, 79: 17, 17, 17, 91: 17, 94: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{79: 713, 715, 109: 714, 712, 116: 737, 137: 751},
		{2: 731, 9: 115, 140: 733, 172: 732, 734},
		// 445
		{42: 9, 79: 9, 9, 9, 91: 9, 94: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{9: 117, 117, 140: 748},
		{9: 114, 740},
		{9: 738},
		{9: 735},
		// 450
		{79: 713, 715, 109: 714, 712, 116: 737, 137: 736},
		{42: 18, 79: 18, 18, 18, 91: 18, 94: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{42: 6, 79: 6, 6, 6, 91: 6, 94: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{79: 713, 715, 109: 714, 712, 116: 737, 137: 739},
		{42: 20, 79: 20, 20, 20, 91: 20, 94: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		// 455
		{2: 741, 140: 742},
		{9: 116, 116, 140: 745},
		{9: 743},
		{79: 713, 715, 109: 714, 712, 116: 737, 137: 744},
		{42: 19, 79: 19, 19, 19, 91: 19, 94: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		// 460
		{9: 746},
		{79: 713, 715, 109: 714, 712, 116: 737, 137: 747},
		{42: 10, 79: 10, 10, 10, 91: 10, 94: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{9: 749},
		{79: 713, 715, 109: 714, 712, 116: 737, 137: 750},
		// 465
		{42: 11, 79: 11, 11, 11, 91: 11, 94: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{42: 21, 79: 21, 21, 21, 91: 21, 94: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{79: 753},
		{79: 713, 715, 40, 91: 40, 94: 40, 704, 705, 706, 701, 702, 703, 707, 708, 698, 709, 710, 109: 714, 712, 116: 711, 129: 696, 131: 695, 700, 697, 699, 136: 754, 141: 755},
		{79: 713, 715, 39, 91: 39, 94: 39, 704, 705, 706, 701, 702, 703, 707, 708, 698, 709, 710, 109: 714, 712, 116: 711, 129: 696, 131: 756, 700, 697, 699},
		// 470
		{81: 31, 91: 31, 94: 31},
		{42: 41, 79: 41, 41, 41, 91: 41, 94: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		{79: 758},
		{79: 713, 715, 40, 91: 40, 94: 40, 704, 705, 706, 701, 702, 703, 707, 708, 698, 709, 710, 109: 714, 712, 116: 711, 129: 696, 131: 695, 700, 697, 699, 136: 754, 141: 759},
		{81: 32, 91: 32, 94: 32},
		// 475
		{79: 761},
		{79: 713, 715, 40, 91: 40, 94: 40, 704, 705, 706, 701, 702, 703, 707, 708, 698, 709, 710, 109: 714, 712, 116: 711, 129: 696, 131: 695, 700, 697, 699, 136: 754, 141: 762},
		{81: 33, 91: 33, 94: 33},
		{81: 24, 91: 772, 202: 773, 771},
		{81: 30, 91: 30, 94: 30},
		// 480
		{81: 27, 91: 27, 94: 766, 167: 770},
		{80: 715, 109: 767, 712},
		{79: 768},
		{79: 713, 715, 40, 91: 40, 94: 40, 704, 705, 706, 701, 702, 703, 707, 708, 698, 709, 710, 109: 714, 712, 116: 711, 129: 696, 131: 695, 700, 697, 699, 136: 754, 141: 769},
		{81: 26, 91: 26, 94: 26},
		// 485
		{81: 29, 91: 29, 94: 29},
		{81: 777, 204: 776},
		{79: 774},
		{81: 23},
		{79: 713, 715, 40, 95: 704, 705, 706, 701, 702, 703, 707, 708, 698, 709, 710, 109: 714, 712, 116: 711, 129: 696, 131: 695, 700, 697, 699, 136: 754, 141: 775},
		// 490
		{81: 25},
		{42: 34, 79: 34, 34, 34, 91: 34, 94: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{42: 22, 79: 22, 22, 22, 91: 22, 94: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{79: 779},
		{42: 36, 79: 36, 36, 36, 91: 36, 94: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), PrettyString(lval.Token): %v\n", yySymName(n), n, n, PrettyString(lval.Token))
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 216

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if !ok || msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			lx := yylex.(*lexer)
			lx.preprocessingFile = nil
		}
	case 2:
		{
			lx := yylex.(*lexer)
			lx.preprocessingFile = yyS[yypt-0].node.(*PreprocessingFile)
		}
	case 3:
		{
			lx := yylex.(*lexer)
			lx.constantExpression = nil
		}
	case 4:
		{
			lx := yylex.(*lexer)
			lx.constantExpression = yyS[yypt-0].node.(*ConstantExpression)
		}
	case 5:
		{
			lx := yylex.(*lexer)
			lx.translationUnit = nil
		}
	case 6:
		{
			lx := yylex.(*lexer)
			if lx.report.Errors(false) == nil && lx.scope.kind != ScopeFile {
				panic("internal error")
			}

			lx.translationUnit = yyS[yypt-0].node.(*TranslationUnit).reverse()
			lx.translationUnit.Declarations = lx.scope
		}
	case 7:
		{
			yyVAL.node = &EnumerationConstant{
				Token: yyS[yypt-0].Token,
			}
		}
	case 8:
		{
			yyVAL.node = &ArgumentExpressionList{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 9:
		{
			yyVAL.node = &ArgumentExpressionList{
				Case: 1,
				ArgumentExpressionList: yyS[yypt-2].node.(*ArgumentExpressionList),
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 10:
		{
			yyVAL.node = (*ArgumentExpressionListOpt)(nil)
		}
	case 11:
		{
			yyVAL.node = &ArgumentExpressionListOpt{
				ArgumentExpressionList: yyS[yypt-0].node.(*ArgumentExpressionList).reverse(),
			}
		}
	case 12:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.scope = lx.scope
		}
	case 13:
		{
			yyVAL.node = &Expression{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 14:
		{
			yyVAL.node = &Expression{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
		}
	case 15:
		{
			yyVAL.node = &Expression{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 16:
		{
			yyVAL.node = &Expression{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
		}
	case 17:
		{
			yyVAL.node = &Expression{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 18:
		{
			yyVAL.node = &Expression{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
		}
	case 19:
		{
			yyVAL.node = &Expression{
				Case:           7,
				Token:          yyS[yypt-2].Token,
				ExpressionList: yyS[yypt-1].node.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 20:
		{
			yyVAL.node = &Expression{
				Case:           8,
				Expression:     yyS[yypt-3].node.(*Expression),
				Token:          yyS[yypt-2].Token,
				ExpressionList: yyS[yypt-1].node.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 21:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Case:       9,
				Expression: yyS[yypt-3].node.(*Expression),
				Token:      yyS[yypt-2].Token,
				ArgumentExpressionListOpt: yyS[yypt-1].node.(*ArgumentExpressionListOpt),
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			o := lhs.ArgumentExpressionListOpt
			if o == nil {
				break
			}

			lhs.Expression.eval(lx)
			s := lx.adjustFnArgs
			lx.adjustFnArgs = true
			for l := o.ArgumentExpressionList; l != nil; l = l.ArgumentExpressionList {
				l.Expression.eval(lx)
			}
			lx.adjustFnArgs = s
		}
	case 22:
		{
			yyVAL.node = &Expression{
				Case:       10,
				Expression: yyS[yypt-2].node.(*Expression),
				Token:      yyS[yypt-1].Token,
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 23:
		{
			yyVAL.node = &Expression{
				Case:       11,
				Expression: yyS[yypt-2].node.(*Expression),
				Token:      yyS[yypt-1].Token,
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 24:
		{
			yyVAL.node = &Expression{
				Case:       12,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 25:
		{
			yyVAL.node = &Expression{
				Case:       13,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 26:
		{
			yyVAL.node = &Expression{
				Case:            14,
				Token:           yyS[yypt-6].Token,
				TypeName:        yyS[yypt-5].node.(*TypeName),
				Token2:          yyS[yypt-4].Token,
				Token3:          yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token4:          yyS[yypt-0].Token,
			}
		}
	case 27:
		{
			yyVAL.node = &Expression{
				Case:       15,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 28:
		{
			yyVAL.node = &Expression{
				Case:       16,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 29:
		{
			yyVAL.node = &Expression{
				Case:       17,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 30:
		{
			yyVAL.node = &Expression{
				Case:       18,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 31:
		{
			yyVAL.node = &Expression{
				Case:       19,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 32:
		{
			yyVAL.node = &Expression{
				Case:       20,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 33:
		{
			yyVAL.node = &Expression{
				Case:       21,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 34:
		{
			yyVAL.node = &Expression{
				Case:       22,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 35:
		{
			yyVAL.node = &Expression{
				Case:       23,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 36:
		{
			yyVAL.node = &Expression{
				Case:     24,
				Token:    yyS[yypt-3].Token,
				Token2:   yyS[yypt-2].Token,
				TypeName: yyS[yypt-1].node.(*TypeName),
				Token3:   yyS[yypt-0].Token,
			}
		}
	case 37:
		{
			yyVAL.node = &Expression{
				Case:       25,
				Token:      yyS[yypt-3].Token,
				TypeName:   yyS[yypt-2].node.(*TypeName),
				Token2:     yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 38:
		{
			yyVAL.node = &Expression{
				Case:        26,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 39:
		{
			yyVAL.node = &Expression{
				Case:        27,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 40:
		{
			yyVAL.node = &Expression{
				Case:        28,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 41:
		{
			yyVAL.node = &Expression{
				Case:        29,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 42:
		{
			yyVAL.node = &Expression{
				Case:        30,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 43:
		{
			yyVAL.node = &Expression{
				Case:        31,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 44:
		{
			yyVAL.node = &Expression{
				Case:        32,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 45:
		{
			yyVAL.node = &Expression{
				Case:        33,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 46:
		{
			yyVAL.node = &Expression{
				Case:        34,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 47:
		{
			yyVAL.node = &Expression{
				Case:        35,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 48:
		{
			yyVAL.node = &Expression{
				Case:        36,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 49:
		{
			yyVAL.node = &Expression{
				Case:        37,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 50:
		{
			yyVAL.node = &Expression{
				Case:        38,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 51:
		{
			yyVAL.node = &Expression{
				Case:        39,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 52:
		{
			yyVAL.node = &Expression{
				Case:        40,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 53:
		{
			yyVAL.node = &Expression{
				Case:        41,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 54:
		{
			yyVAL.node = &Expression{
				Case:        42,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 55:
		{
			yyVAL.node = &Expression{
				Case:        43,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 56:
		{
			yyVAL.node = &Expression{
				Case:           44,
				Expression:     yyS[yypt-4].node.(*Expression),
				Token:          yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-1].Token,
				Expression2:    yyS[yypt-0].node.(*Expression),
			}
		}
	case 57:
		{
			yyVAL.node = &Expression{
				Case:        45,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 58:
		{
			yyVAL.node = &Expression{
				Case:        46,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 59:
		{
			yyVAL.node = &Expression{
				Case:        47,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 60:
		{
			yyVAL.node = &Expression{
				Case:        48,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 61:
		{
			yyVAL.node = &Expression{
				Case:        49,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 62:
		{
			yyVAL.node = &Expression{
				Case:        50,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 63:
		{
			yyVAL.node = &Expression{
				Case:        51,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 64:
		{
			yyVAL.node = &Expression{
				Case:        52,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 65:
		{
			yyVAL.node = &Expression{
				Case:        53,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 66:
		{
			yyVAL.node = &Expression{
				Case:        54,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 67:
		{
			yyVAL.node = &Expression{
				Case:        55,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 68:
		{
			yyVAL.node = (*ExpressionOpt)(nil)
		}
	case 69:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionOpt{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 70:
		{
			yyVAL.node = &ExpressionList{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 71:
		{
			yyVAL.node = &ExpressionList{
				Case:           1,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList),
				Token:          yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 72:
		{
			yyVAL.node = (*ExpressionListOpt)(nil)
		}
	case 73:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionListOpt{
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 74:
		{
			lx := yylex.(*lexer)
			lx.constExprToks = []xc.Token{lx.last}
		}
	case 75:
		{
			lx := yylex.(*lexer)
			lhs := &ConstantExpression{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Value, lhs.Type = lhs.Expression.eval(lx)
			if lhs.Value == nil {
				lx.report.Err(lhs.Pos(), "not a constant expression")
			}
			l := lx.constExprToks
			lhs.toks = l[:len(l)-1]
			lx.constExprToks = nil
		}
	case 76:
		{
			lx := yylex.(*lexer)
			lhs := &Declaration{
				DeclarationSpecifiers: yyS[yypt-2].node.(*DeclarationSpecifiers),
				InitDeclaratorListOpt: yyS[yypt-1].node.(*InitDeclaratorListOpt),
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			ts := tsDecode(lhs.DeclarationSpecifiers.typeSpecifiers())
			ok := false
			for _, v := range ts {
				if v == tsStructSpecifier || v == tsUnionSpecifier {
					ok = true
					break
				}
			}
			if ok {
				s := lhs.DeclarationSpecifiers
				d := &Declarator{specifier: s}
				dd := &DirectDeclarator{
					Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
					declarator: d,
					idScope:    lx.scope,
					specifier:  s,
				}
				d.DirectDeclarator = dd
				d.setFull(lx)
				for l := lhs.DeclarationSpecifiers; l != nil; {
					ts := l.TypeSpecifier
					if ts != nil && ts.Case == 11 && ts.StructOrUnionSpecifier.Case == 0 { // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
						ts.StructOrUnionSpecifier.declarator = d
						break
					}

					if o := l.DeclarationSpecifiersOpt; o != nil {
						l = o.DeclarationSpecifiers
						continue
					}

					break
				}
			}

			o := lhs.InitDeclaratorListOpt
			if o != nil {
				break
			}

			s := lhs.DeclarationSpecifiers
			d := &Declarator{specifier: s}
			dd := &DirectDeclarator{
				Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
				declarator: d,
				idScope:    lx.scope,
				specifier:  s,
			}
			d.DirectDeclarator = dd
			d.setFull(lx)
			lhs.declarator = d
		}
	case 77:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				StorageClassSpecifier:    yyS[yypt-1].node.(*StorageClassSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.StorageClassSpecifier
			b := lhs.DeclarationSpecifiersOpt
			if b == nil {
				lhs.attr = a.attr
				break
			}

			if a.attr&b.attr != 0 {
				lx.report.Err(a.Pos(), "invalid storage class specifier")
				break
			}

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
			if lhs.StorageClassSpecifier.Case != 0 /* "typedef" */ && lhs.IsTypedef() {
				lx.report.Err(a.Pos(), "invalid storage class specifier")
			}
		}
	case 78:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     1,
				TypeSpecifier:            yyS[yypt-1].node.(*TypeSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.TypeSpecifier
			b := lhs.DeclarationSpecifiersOpt
			if b == nil {
				lhs.typeSpecifier = a.typeSpecifier
				break
			}

			lhs.attr = b.attr
			ts := tsEncode(append(tsDecode(a.typeSpecifier), tsDecode(b.typeSpecifier)...)...)
			if _, ok := tsValid[ts]; !ok {
				ts = tsEncode(tsInt)
				lx.report.Err(a.Pos(), "invalid type specifier")
			}
			lhs.typeSpecifier = ts
		}
	case 79:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     2,
				TypeQualifier:            yyS[yypt-1].node.(*TypeQualifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.TypeQualifier
			b := lhs.DeclarationSpecifiersOpt
			if b == nil {
				lhs.attr = a.attr
				break
			}

			if a.attr&b.attr != 0 {
				lx.report.Err(a.Pos(), "invalid type qualifier")
				break
			}

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
			if lhs.IsTypedef() {
				lx.report.Err(a.Pos(), "invalid type qualifier")
			}
		}
	case 80:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     3,
				FunctionSpecifier:        yyS[yypt-1].node.(*FunctionSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.FunctionSpecifier
			b := lhs.DeclarationSpecifiersOpt
			if b == nil {
				lhs.attr = a.attr
				break
			}

			if a.attr&b.attr != 0 {
				lx.report.Err(a.Pos(), "invalid function specifier")
				break
			}

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
			if lhs.IsTypedef() {
				lx.report.Err(a.Pos(), "invalid function specifier")
			}
		}
	case 81:
		{
			yyVAL.node = (*DeclarationSpecifiersOpt)(nil)
		}
	case 82:
		{
			lhs := &DeclarationSpecifiersOpt{
				DeclarationSpecifiers: yyS[yypt-0].node.(*DeclarationSpecifiers),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.DeclarationSpecifiers.attr
			lhs.typeSpecifier = lhs.DeclarationSpecifiers.typeSpecifier
		}
	case 83:
		{
			yyVAL.node = &InitDeclaratorList{
				InitDeclarator: yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 84:
		{
			yyVAL.node = &InitDeclaratorList{
				Case:               1,
				InitDeclaratorList: yyS[yypt-2].node.(*InitDeclaratorList),
				Token:              yyS[yypt-1].Token,
				InitDeclarator:     yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 85:
		{
			yyVAL.node = (*InitDeclaratorListOpt)(nil)
		}
	case 86:
		{
			yyVAL.node = &InitDeclaratorListOpt{
				InitDeclaratorList: yyS[yypt-0].node.(*InitDeclaratorList).reverse(),
			}
		}
	case 87:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
		}
	case 88:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-0].node.(*Declarator)
			d.setFull(lx)
		}
	case 89:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Case:        1,
				Declarator:  yyS[yypt-3].node.(*Declarator),
				Token:       yyS[yypt-1].Token,
				Initializer: yyS[yypt-0].node.(*Initializer),
			}
			yyVAL.node = lhs
			switch i := lhs.Initializer; i.Case {
			case 0: // Expression
				s := i.Expression.Type
				d := lhs.Declarator.Type
				if !s.CanAssignTo(s) {
					lx.report.Err(i.Pos(), "incompatible types when initializing type '%s' using type ‘%s'", d, s)
				}
			case 1: // '{' InitializerList CommaOpt '}'
				limit := -1
				var checkType Type
				var mb []Member
				d := lhs.Declarator
				k := d.Type.Kind()
				switch k {
				case Array:
					checkType = d.Type.Element()
					limit = d.Type.Elements()
				case Ptr:
					checkType = d.Type.Element()
					d.Type = d.Type.(*ctype).setElements(i.InitializerList.Len())
				case Struct, Union:
					mb, _ = d.Type.Members()
					if mb == nil {
						panic("internal error")
					}

					limit = len(mb)
					if k == Union {
						limit = 1
					}
				default:
					//dbg("", position(d.Pos()), d.Type.Kind())
					panic("TODO")
				}

				values := 0
				for l := i.InitializerList; l != nil; l = l.InitializerList {
					values++
					l.Initializer.typeCheck(checkType, mb, values-1, limit, lx)
				}
			default:
				panic(i.Case)
			}
		}
	case 90:
		{
			lhs := &StorageClassSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saTypedef
		}
	case 91:
		{
			lhs := &StorageClassSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saExtern
		}
	case 92:
		{
			lhs := &StorageClassSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saStatic
		}
	case 93:
		{
			lhs := &StorageClassSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saAuto
		}
	case 94:
		{
			lhs := &StorageClassSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRegister
		}
	case 95:
		{
			lhs := &TypeSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsVoid)
		}
	case 96:
		{
			lhs := &TypeSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsChar)
		}
	case 97:
		{
			lhs := &TypeSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsShort)
		}
	case 98:
		{
			lhs := &TypeSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsInt)
		}
	case 99:
		{
			lhs := &TypeSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsLong)
		}
	case 100:
		{
			lhs := &TypeSpecifier{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsFloat)
		}
	case 101:
		{
			lhs := &TypeSpecifier{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsDouble)
		}
	case 102:
		{
			lhs := &TypeSpecifier{
				Case:  7,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsSigned)
		}
	case 103:
		{
			lhs := &TypeSpecifier{
				Case:  8,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsUnsigned)
		}
	case 104:
		{
			lhs := &TypeSpecifier{
				Case:  9,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsBool)
		}
	case 105:
		{
			lhs := &TypeSpecifier{
				Case:  10,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsComplex)
		}
	case 106:
		{
			lhs := &TypeSpecifier{
				Case: 11,
				StructOrUnionSpecifier: yyS[yypt-0].node.(*StructOrUnionSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(lhs.StructOrUnionSpecifier.typeSpecifiers())
		}
	case 107:
		{
			lhs := &TypeSpecifier{
				Case:          12,
				EnumSpecifier: yyS[yypt-0].node.(*EnumSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsEnumSpecifier)
		}
	case 108:
		{
			lx := yylex.(*lexer)
			lhs := &TypeSpecifier{
				Case:  13,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsTypedefName)
			_, lhs.scope = lx.scope.Lookup2(NSIdentifiers, lhs.Token.Val)
		}
	case 109:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareStructTag(o.Token, lx.report)
			}
		}
	case 110:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeMembers)
			lx.scope.isUnion = yyS[yypt-3].node.(*StructOrUnion).Case == 1 // "union"
			lx.scope.prevStructDeclarator = nil
		}
	case 111:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				StructOrUnion: yyS[yypt-6].node.(*StructOrUnion),
				IdentifierOpt: yyS[yypt-5].node.(*IdentifierOpt),
				Token:         yyS[yypt-3].Token,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList).reverse(),
				Token2:                yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			sc := lx.scope
			lhs.scope = sc
			if sc.bitOffset != 0 {
				finishBitField(lx)
			}

			i := 0
			var bt Type
			var d *Declarator
			for l := lhs.StructDeclarationList; l != nil; l = l.StructDeclarationList {
				for l := l.StructDeclaration.StructDeclaratorList; l != nil; l = l.StructDeclaratorList {
					switch sd := l.StructDeclarator; sd.Case {
					case 0: // Declarator
						d = sd.Declarator
					case 1: // DeclaratorOpt ':' ConstantExpression
						if o := sd.DeclaratorOpt; o != nil {
							x := o.Declarator
							if x.bitOffset == 0 {
								d = x
								bt = lx.scope.bitFieldTypes[i]
								i++
							}
							x.bitFieldType = bt
						}
					}
				}
			}
			lx.scope.bitFieldTypes = nil

			lhs.alignOf = sc.maxAlign
			switch {
			case sc.isUnion:
				lhs.sizeOf = align(sc.maxSize, sc.maxAlign)
			default:
				off := sc.offset
				lhs.sizeOf = align(sc.offset, sc.maxAlign)
				if d != nil {
					d.padding = lhs.sizeOf - off
				}
			}
			if lhs.sizeOf == 0 {
				panic("internal error")
			}

			lx.popScope(lhs.Token2)
			if o := lhs.IdentifierOpt; o != nil {
				lx.scope.defineStructTag(o.Token, lhs, lx.report)
			}
		}
	case 112:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				Case:          1,
				StructOrUnion: yyS[yypt-1].node.(*StructOrUnion),
				Token:         yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lx.scope.declareStructTag(lhs.Token, lx.report)
			lhs.scope = lx.scope
		}
	case 113:
		{
			yyVAL.node = &StructOrUnion{
				Token: yyS[yypt-0].Token,
			}
		}
	case 114:
		{
			yyVAL.node = &StructOrUnion{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 115:
		{
			yyVAL.node = &StructDeclarationList{
				StructDeclaration: yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 116:
		{
			yyVAL.node = &StructDeclarationList{
				Case: 1,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList),
				StructDeclaration:     yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 117:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclaration{
				SpecifierQualifierList: yyS[yypt-2].node.(*SpecifierQualifierList),
				StructDeclaratorList:   yyS[yypt-1].node.(*StructDeclaratorList).reverse(),
				Token:                  yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			s := lhs.SpecifierQualifierList
			if k := s.kind(); k != Struct && k != Union {
				break
			}

			d := &Declarator{specifier: s}
			dd := &DirectDeclarator{
				Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
				declarator: d,
				idScope:    lx.scope,
				specifier:  s,
			}
			d.DirectDeclarator = dd
			d.setFull(lx)
			for l := lhs.SpecifierQualifierList; l != nil; {
				ts := l.TypeSpecifier
				if ts != nil && ts.Case == 11 && ts.StructOrUnionSpecifier.Case == 0 { // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
					ts.StructOrUnionSpecifier.declarator = d
					break
				}

				if o := l.SpecifierQualifierListOpt; o != nil {
					l = o.SpecifierQualifierList
					continue
				}

				break
			}
		}
	case 118:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclaration{
				Case: 1,
				SpecifierQualifierList: yyS[yypt-1].node.(*SpecifierQualifierList),
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			s := lhs.SpecifierQualifierList
			if k := s.kind(); k != Struct && k != Union {
				lx.report.Err(lhs.Token.Pos(), "only unnamed structs and unions are allowed")
				break
			}

			if !lx.tweaks.enableAnonymousStructFields {
				lx.report.Err(lhs.Token.Pos(), "unnamed structs and unions not allowed")
			}

			d := &Declarator{specifier: s}
			dd := &DirectDeclarator{
				Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
				declarator: d,
				idScope:    lx.scope,
				specifier:  s,
			}
			d.DirectDeclarator = dd
			d.setFull(lx)
			for l := lhs.SpecifierQualifierList; l != nil; {
				ts := l.TypeSpecifier
				if ts != nil && ts.Case == 11 && ts.StructOrUnionSpecifier.Case == 0 { // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
					ts.StructOrUnionSpecifier.declarator = d
					break
				}

				if o := l.SpecifierQualifierListOpt; o != nil {
					l = o.SpecifierQualifierList
					continue
				}

				break
			}
		}
	case 119:
		{
			lx := yylex.(*lexer)
			lhs := &SpecifierQualifierList{
				TypeSpecifier:             yyS[yypt-1].node.(*TypeSpecifier),
				SpecifierQualifierListOpt: yyS[yypt-0].node.(*SpecifierQualifierListOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.TypeSpecifier
			b := lhs.SpecifierQualifierListOpt
			if b == nil {
				lhs.typeSpecifier = a.typeSpecifier
				break
			}

			lhs.attr = b.attr
			ts := tsEncode(append(tsDecode(a.typeSpecifier), tsDecode(b.typeSpecifier)...)...)
			if _, ok := tsValid[ts]; !ok {
				lx.report.Err(a.Pos(), "invalid type specifier")
				break
			}

			lhs.typeSpecifier = ts
		}
	case 120:
		{
			lx := yylex.(*lexer)
			lhs := &SpecifierQualifierList{
				Case:                      1,
				TypeQualifier:             yyS[yypt-1].node.(*TypeQualifier),
				SpecifierQualifierListOpt: yyS[yypt-0].node.(*SpecifierQualifierListOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.TypeQualifier
			b := lhs.SpecifierQualifierListOpt
			if b == nil {
				lhs.attr = a.attr
				break
			}

			if a.attr&b.attr != 0 {
				lx.report.Err(a.Pos(), "invalid type qualifier")
				break
			}

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
		}
	case 121:
		{
			yyVAL.node = (*SpecifierQualifierListOpt)(nil)
		}
	case 122:
		{
			lhs := &SpecifierQualifierListOpt{
				SpecifierQualifierList: yyS[yypt-0].node.(*SpecifierQualifierList),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.SpecifierQualifierList.attr
			lhs.typeSpecifier = lhs.SpecifierQualifierList.typeSpecifier
		}
	case 123:
		{
			yyVAL.node = &StructDeclaratorList{
				StructDeclarator: yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 124:
		{
			yyVAL.node = &StructDeclaratorList{
				Case:                 1,
				StructDeclaratorList: yyS[yypt-2].node.(*StructDeclaratorList),
				Token:                yyS[yypt-1].Token,
				StructDeclarator:     yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 125:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.post(lx)
		}
	case 126:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Case:               1,
				DeclaratorOpt:      yyS[yypt-2].node.(*DeclaratorOpt),
				Token:              yyS[yypt-1].Token,
				ConstantExpression: yyS[yypt-0].node.(*ConstantExpression),
			}
			yyVAL.node = lhs
			m := lx.model
			e := lhs.ConstantExpression
			if e.Value == nil {
				e.Value, e.Type = m.value2(1, m.IntType)
			}
			if e.Type.Kind() != Int {
				lx.report.Err(e.Pos(), "bit field width not an integer")
				e.Value, e.Type = m.value2(1, m.IntType)
			}
			if o := lhs.DeclaratorOpt; o != nil {
				o.Declarator.setFull(lx)
			}
			lhs.post(lx)
		}
	case 127:
		{
			yyVAL.node = (*CommaOpt)(nil)
		}
	case 128:
		{
			yyVAL.node = &CommaOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 129:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareEnumTag(o.Token, lx.report)
			}
			lx.iota = 0
		}
	case 130:
		{
			lx := yylex.(*lexer)
			lhs := &EnumSpecifier{
				Token:          yyS[yypt-6].Token,
				IdentifierOpt:  yyS[yypt-5].node.(*IdentifierOpt),
				Token2:         yyS[yypt-3].Token,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList).reverse(),
				CommaOpt:       yyS[yypt-1].node.(*CommaOpt),
				Token3:         yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if o := lhs.IdentifierOpt; o != nil {
				lx.scope.defineEnumTag(o.Token, lhs, lx.report)
			}
		}
	case 131:
		{
			lx := yylex.(*lexer)
			lhs := &EnumSpecifier{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lx.scope.declareEnumTag(lhs.Token2, lx.report)
		}
	case 132:
		{
			yyVAL.node = &EnumeratorList{
				Enumerator: yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 133:
		{
			yyVAL.node = &EnumeratorList{
				Case:           1,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList),
				Token:          yyS[yypt-1].Token,
				Enumerator:     yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 134:
		{
			lx := yylex.(*lexer)
			lhs := &Enumerator{
				EnumerationConstant: yyS[yypt-0].node.(*EnumerationConstant),
			}
			yyVAL.node = lhs
			m := lx.model
			v := m.MustConvert(lx.iota, m.IntType)
			lhs.enumVal = v
			lx.scope.defineEnumConst(lx, lhs.EnumerationConstant.Token, v)
		}
	case 135:
		{
			lx := yylex.(*lexer)
			lhs := &Enumerator{
				Case:                1,
				EnumerationConstant: yyS[yypt-2].node.(*EnumerationConstant),
				Token:               yyS[yypt-1].Token,
				ConstantExpression:  yyS[yypt-0].node.(*ConstantExpression),
			}
			yyVAL.node = lhs
			m := lx.model
			e := lhs.ConstantExpression
			var v interface{}
			// [0], 6.7.2.2
			// The expression that defines the value of an enumeration
			// constant shall be an integer constant expression that has a
			// value representable as an int.
			switch {
			case !IsIntType(e.Type):
				lx.report.Err(e.Pos(), "not an integer constant expression (have '%s')", e.Type)
				v = m.MustConvert(0, m.IntType)
			default:
				var ok bool
				if v, ok = m.toInt(e.Value); !ok {
					lx.report.Err(e.Pos(), "overflow in enumeration value: %v", e.Value)
				}
			}

			lhs.enumVal = v
			lx.scope.defineEnumConst(lx, lhs.EnumerationConstant.Token, v)
		}
	case 136:
		{
			lhs := &TypeQualifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saConst
		}
	case 137:
		{
			lhs := &TypeQualifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRestrict
		}
	case 138:
		{
			lhs := &TypeQualifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saVolatile
		}
	case 139:
		{
			lhs := &FunctionSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saInline
		}
	case 140:
		{
			lx := yylex.(*lexer)
			lhs := &Declarator{
				PointerOpt:       yyS[yypt-1].node.(*PointerOpt),
				DirectDeclarator: yyS[yypt-0].node.(*DirectDeclarator),
			}
			yyVAL.node = lhs
			lhs.specifier = lx.scope.specifier
			lhs.DirectDeclarator.declarator = lhs
		}
	case 141:
		{
			yyVAL.node = (*DeclaratorOpt)(nil)
		}
	case 142:
		{
			yyVAL.node = &DeclaratorOpt{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
		}
	case 143:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.specifier = lx.scope.specifier
			lx.scope.declareIdentifier(lhs.Token, lhs, lx.report)
			lhs.idScope = lx.scope
		}
	case 144:
		{
			lhs := &DirectDeclarator{
				Case:       1,
				Token:      yyS[yypt-2].Token,
				Declarator: yyS[yypt-1].node.(*Declarator),
				Token2:     yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Declarator.specifier = nil
			lhs.Declarator.DirectDeclarator.parent = lhs
		}
	case 145:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:                 2,
				DirectDeclarator:     yyS[yypt-4].node.(*DirectDeclarator),
				Token:                yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				ExpressionOpt:        yyS[yypt-1].node.(*ExpressionOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.elements = -1
			if o := lhs.ExpressionOpt; o != nil {
				o.Expression.eval(lx)
				var err error
				if lhs.elements, err = elements(o.Expression.Value); err != nil {
					lx.report.Err(o.Expression.Pos(), "%s", err)
				}

			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 146:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:                 3,
				DirectDeclarator:     yyS[yypt-5].node.(*DirectDeclarator),
				Token:                yyS[yypt-4].Token,
				Token2:               yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				Expression:           yyS[yypt-1].node.(*Expression),
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
			var err error
			if lhs.elements, err = elements(lhs.Expression.Value); err != nil {
				lx.report.Err(lhs.Expression.Pos(), "%s", err)
			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 147:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:              4,
				DirectDeclarator:  yyS[yypt-5].node.(*DirectDeclarator),
				Token:             yyS[yypt-4].Token,
				TypeQualifierList: yyS[yypt-3].node.(*TypeQualifierList).reverse(),
				Token2:            yyS[yypt-2].Token,
				Expression:        yyS[yypt-1].node.(*Expression),
				Token3:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
			var err error
			if lhs.elements, err = elements(lhs.Expression.Value); err != nil {
				lx.report.Err(lhs.Expression.Pos(), "%s", err)
			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 148:
		{
			lhs := &DirectDeclarator{
				Case:                 5,
				DirectDeclarator:     yyS[yypt-4].node.(*DirectDeclarator),
				Token:                yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				Token2:               yyS[yypt-1].Token,
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.DirectDeclarator.parent = lhs
			lhs.elements = -1
		}
	case 149:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 150:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:              6,
				DirectDeclarator:  yyS[yypt-4].node.(*DirectDeclarator),
				Token:             yyS[yypt-3].Token,
				ParameterTypeList: yyS[yypt-1].node.(*ParameterTypeList),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScope(lhs.Token2)
			lhs.DirectDeclarator.parent = lhs
		}
	case 151:
		{
			lhs := &DirectDeclarator{
				Case:              7,
				DirectDeclarator:  yyS[yypt-3].node.(*DirectDeclarator),
				Token:             yyS[yypt-2].Token,
				IdentifierListOpt: yyS[yypt-1].node.(*IdentifierListOpt),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.DirectDeclarator.parent = lhs
		}
	case 152:
		{
			yyVAL.node = &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].node.(*TypeQualifierListOpt),
			}
		}
	case 153:
		{
			yyVAL.node = &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].node.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].node.(*Pointer),
			}
		}
	case 154:
		{
			yyVAL.node = (*PointerOpt)(nil)
		}
	case 155:
		{
			yyVAL.node = &PointerOpt{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
		}
	case 156:
		{
			lhs := &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.TypeQualifier.attr
		}
	case 157:
		{
			lx := yylex.(*lexer)
			lhs := &TypeQualifierList{
				Case:              1,
				TypeQualifierList: yyS[yypt-1].node.(*TypeQualifierList),
				TypeQualifier:     yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			a := lhs.TypeQualifierList
			b := lhs.TypeQualifier
			if a.attr&b.attr != 0 {
				lx.report.Err(b.Pos(), "invalid type qualifier")
				break
			}

			lhs.attr = a.attr | b.attr
		}
	case 158:
		{
			yyVAL.node = (*TypeQualifierListOpt)(nil)
		}
	case 159:
		{
			yyVAL.node = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].node.(*TypeQualifierList).reverse(),
			}
		}
	case 160:
		{
			lhs := &ParameterTypeList{
				ParameterList: yyS[yypt-0].node.(*ParameterList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 161:
		{
			lhs := &ParameterTypeList{
				Case:          1,
				ParameterList: yyS[yypt-2].node.(*ParameterList).reverse(),
				Token:         yyS[yypt-1].Token,
				Token2:        yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 162:
		{
			yyVAL.node = (*ParameterTypeListOpt)(nil)
		}
	case 163:
		{
			yyVAL.node = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].node.(*ParameterTypeList),
			}
		}
	case 164:
		{
			yyVAL.node = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 165:
		{
			yyVAL.node = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].node.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 166:
		{
			lx := yylex.(*lexer)
			lhs := &ParameterDeclaration{
				DeclarationSpecifiers: yyS[yypt-1].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.declarator = lhs.Declarator
		}
	case 167:
		{
			lx := yylex.(*lexer)
			lhs := &ParameterDeclaration{
				Case: 1,
				DeclarationSpecifiers: yyS[yypt-1].node.(*DeclarationSpecifiers),
				AbstractDeclaratorOpt: yyS[yypt-0].node.(*AbstractDeclaratorOpt),
			}
			yyVAL.node = lhs
			if o := lhs.AbstractDeclaratorOpt; o != nil {
				lhs.declarator = o.AbstractDeclarator.declarator
				lhs.declarator.setFull(lx)
				break
			}

			d := &Declarator{
				specifier: lx.scope.specifier,
				DirectDeclarator: &DirectDeclarator{
					Case: 0, // IDENTIFIER
				},
			}
			d.DirectDeclarator.declarator = d
			lhs.declarator = d
			d.setFull(lx)
		}
	case 168:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 169:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 170:
		{
			yyVAL.node = (*IdentifierListOpt)(nil)
		}
	case 171:
		{
			yyVAL.node = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 172:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 173:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 174:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeBlock)
		}
	case 175:
		{
			lx := yylex.(*lexer)
			lhs := &TypeName{
				SpecifierQualifierList: yyS[yypt-1].node.(*SpecifierQualifierList),
				AbstractDeclaratorOpt:  yyS[yypt-0].node.(*AbstractDeclaratorOpt),
			}
			yyVAL.node = lhs
			if o := lhs.AbstractDeclaratorOpt; o != nil {
				lhs.declarator = o.AbstractDeclarator.declarator
			} else {
				d := &Declarator{
					specifier: lhs.SpecifierQualifierList,
					DirectDeclarator: &DirectDeclarator{
						Case:    0, // IDENTIFIER
						idScope: lx.scope,
					},
				}
				d.DirectDeclarator.declarator = d
				lhs.declarator = d
			}
			lhs.Type = lhs.declarator.setFull(lx)
			lx.popScope(xc.Token{})
		}
	case 176:
		{
			lx := yylex.(*lexer)
			lhs := &AbstractDeclarator{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
			yyVAL.node = lhs
			d := &Declarator{
				specifier: lx.scope.specifier,
				PointerOpt: &PointerOpt{
					Pointer: lhs.Pointer,
				},
				DirectDeclarator: &DirectDeclarator{
					Case:    0, // IDENTIFIER
					idScope: lx.scope,
				},
			}
			d.DirectDeclarator.declarator = d
			lhs.declarator = d
		}
	case 177:
		{
			lx := yylex.(*lexer)
			lhs := &AbstractDeclarator{
				Case:                     1,
				PointerOpt:               yyS[yypt-1].node.(*PointerOpt),
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
			yyVAL.node = lhs
			d := &Declarator{
				specifier:        lx.scope.specifier,
				PointerOpt:       lhs.PointerOpt,
				DirectDeclarator: lhs.DirectAbstractDeclarator.directDeclarator,
			}
			d.DirectDeclarator.declarator = d
			lhs.declarator = d
		}
	case 178:
		{
			yyVAL.node = (*AbstractDeclaratorOpt)(nil)
		}
	case 179:
		{
			yyVAL.node = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].node.(*AbstractDeclarator),
			}
		}
	case 180:
		{
			lhs := &DirectAbstractDeclarator{
				Token:              yyS[yypt-2].Token,
				AbstractDeclarator: yyS[yypt-1].node.(*AbstractDeclarator),
				Token2:             yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.AbstractDeclarator.declarator.specifier = nil
			lhs.directDeclarator = &DirectDeclarator{
				Case:       1, // '(' Declarator ')'
				Declarator: lhs.AbstractDeclarator.declarator,
			}
			lhs.AbstractDeclarator.declarator.DirectDeclarator.parent = lhs.directDeclarator
		}
	case 181:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 1,
				DirectAbstractDeclaratorOpt: yyS[yypt-3].node.(*DirectAbstractDeclaratorOpt),
				Token:         yyS[yypt-2].Token,
				ExpressionOpt: yyS[yypt-1].node.(*ExpressionOpt),
				Token2:        yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if o := lhs.ExpressionOpt; o != nil {
				o.Expression.eval(lx)
			}
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:             2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
				DirectDeclarator: dd,
				ExpressionOpt:    lhs.ExpressionOpt,
			}
			dd.parent = lhs.directDeclarator
		}
	case 182:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 2,
				DirectAbstractDeclaratorOpt: yyS[yypt-4].node.(*DirectAbstractDeclaratorOpt),
				Token:             yyS[yypt-3].Token,
				TypeQualifierList: yyS[yypt-2].node.(*TypeQualifierList).reverse(),
				ExpressionOpt:     yyS[yypt-1].node.(*ExpressionOpt),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if o := lhs.ExpressionOpt; o != nil {
				o.Expression.eval(lx)
			}
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:                 2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
				DirectDeclarator:     dd,
				TypeQualifierListOpt: &TypeQualifierListOpt{lhs.TypeQualifierList},
				ExpressionOpt:        lhs.ExpressionOpt,
			}
			dd.parent = lhs.directDeclarator
		}
	case 183:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 3,
				DirectAbstractDeclaratorOpt: yyS[yypt-5].node.(*DirectAbstractDeclaratorOpt),
				Token:                yyS[yypt-4].Token,
				Token2:               yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				Expression:           yyS[yypt-1].node.(*Expression),
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:                 2, // DirectDeclarator '[' "static" TypeQualifierListOpt Expression ']'
				DirectDeclarator:     dd,
				TypeQualifierListOpt: lhs.TypeQualifierListOpt,
				Expression:           lhs.Expression,
			}
			dd.parent = lhs.directDeclarator
		}
	case 184:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 4,
				DirectAbstractDeclaratorOpt: yyS[yypt-5].node.(*DirectAbstractDeclaratorOpt),
				Token:             yyS[yypt-4].Token,
				TypeQualifierList: yyS[yypt-3].node.(*TypeQualifierList).reverse(),
				Token2:            yyS[yypt-2].Token,
				Expression:        yyS[yypt-1].node.(*Expression),
				Token3:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:              4, // DirectDeclarator '[' TypeQualifierList "static" Expression ']'
				DirectDeclarator:  dd,
				TypeQualifierList: lhs.TypeQualifierList,
				Expression:        lhs.Expression,
			}
			dd.parent = lhs.directDeclarator
		}
	case 185:
		{
			lhs := &DirectAbstractDeclarator{
				Case: 5,
				DirectAbstractDeclaratorOpt: yyS[yypt-3].node.(*DirectAbstractDeclaratorOpt),
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:             5, // DirectDeclarator '[' TypeQualifierListOpt '*' ']'
				DirectDeclarator: dd,
			}
			dd.parent = lhs.directDeclarator
		}
	case 186:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 187:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case:                 6,
				Token:                yyS[yypt-3].Token,
				ParameterTypeListOpt: yyS[yypt-1].node.(*ParameterTypeListOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScope(lhs.Token2)
			switch o := lhs.ParameterTypeListOpt; {
			case o != nil:
				lhs.directDeclarator = &DirectDeclarator{
					Case: 6, // DirectDeclarator '(' ParameterTypeList ')'
					DirectDeclarator: &DirectDeclarator{
						Case: 0, // IDENTIFIER
					},
					ParameterTypeList: o.ParameterTypeList,
				}
			default:
				lhs.directDeclarator = &DirectDeclarator{
					Case: 7, // DirectDeclarator '(' IdentifierListOpt ')'
					DirectDeclarator: &DirectDeclarator{
						Case: 0, // IDENTIFIER
					},
				}
			}
			lhs.directDeclarator.DirectDeclarator.parent = lhs.directDeclarator
		}
	case 188:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 189:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 7,
				DirectAbstractDeclarator: yyS[yypt-4].node.(*DirectAbstractDeclarator),
				Token:                yyS[yypt-3].Token,
				ParameterTypeListOpt: yyS[yypt-1].node.(*ParameterTypeListOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScope(lhs.Token2)
			switch o := lhs.ParameterTypeListOpt; {
			case o != nil:
				lhs.directDeclarator = &DirectDeclarator{
					Case:              6, // DirectDeclarator '(' ParameterTypeList ')'
					DirectDeclarator:  lhs.DirectAbstractDeclarator.directDeclarator,
					ParameterTypeList: o.ParameterTypeList,
				}
			default:
				lhs.directDeclarator = &DirectDeclarator{
					Case:             7, // DirectDeclarator '(' IdentifierListOpt ')'
					DirectDeclarator: lhs.DirectAbstractDeclarator.directDeclarator,
				}
			}
			lhs.directDeclarator.DirectDeclarator.parent = lhs.directDeclarator
		}
	case 190:
		{
			yyVAL.node = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 191:
		{
			yyVAL.node = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
		}
	case 192:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 193:
		{
			yyVAL.node = &Initializer{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 194:
		{
			yyVAL.node = &InitializerList{
				DesignationOpt: yyS[yypt-1].node.(*DesignationOpt),
				Initializer:    yyS[yypt-0].node.(*Initializer),
			}
		}
	case 195:
		{
			yyVAL.node = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].node.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].node.(*DesignationOpt),
				Initializer:     yyS[yypt-0].node.(*Initializer),
			}
		}
	case 196:
		{
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 197:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 198:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 199:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 200:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 201:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 202:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 203:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 204:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 205:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 206:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 207:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 208:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 209:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 210:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 211:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 212:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 213:
		{
			lx := yylex.(*lexer)
			lhs := &CompoundStatement{
				Token:            yyS[yypt-3].Token,
				BlockItemListOpt: yyS[yypt-1].node.(*BlockItemListOpt),
				Token2:           yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.scope = lx.scope
			lx.popScope(lhs.Token2)
		}
	case 214:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 215:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 216:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 217:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 218:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 219:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 220:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 221:
		{
			lx := yylex.(*lexer)
			lhs := &SelectionStatement{
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 222:
		{
			lx := yylex.(*lexer)
			lhs := &SelectionStatement{
				Case:           1,
				Token:          yyS[yypt-6].Token,
				Token2:         yyS[yypt-5].Token,
				ExpressionList: yyS[yypt-4].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-3].Token,
				Statement:      yyS[yypt-2].node.(*Statement),
				Token4:         yyS[yypt-1].Token,
				Statement2:     yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 223:
		{
			lx := yylex.(*lexer)
			lhs := &SelectionStatement{
				Case:           2,
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 224:
		{
			lx := yylex.(*lexer)
			lhs := &IterationStatement{
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 225:
		{
			lx := yylex.(*lexer)
			lhs := &IterationStatement{
				Case:           1,
				Token:          yyS[yypt-6].Token,
				Statement:      yyS[yypt-5].node.(*Statement),
				Token2:         yyS[yypt-4].Token,
				Token3:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token4:         yyS[yypt-1].Token,
				Token5:         yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 226:
		{
			yyVAL.node = &IterationStatement{
				Case:               2,
				Token:              yyS[yypt-8].Token,
				Token2:             yyS[yypt-7].Token,
				ExpressionListOpt:  yyS[yypt-6].node.(*ExpressionListOpt),
				Token3:             yyS[yypt-5].Token,
				ExpressionListOpt2: yyS[yypt-4].node.(*ExpressionListOpt),
				Token4:             yyS[yypt-3].Token,
				ExpressionListOpt3: yyS[yypt-2].node.(*ExpressionListOpt),
				Token5:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 227:
		{
			yyVAL.node = &IterationStatement{
				Case:               3,
				Token:              yyS[yypt-7].Token,
				Token2:             yyS[yypt-6].Token,
				Declaration:        yyS[yypt-5].node.(*Declaration),
				ExpressionListOpt:  yyS[yypt-4].node.(*ExpressionListOpt),
				Token3:             yyS[yypt-3].Token,
				ExpressionListOpt2: yyS[yypt-2].node.(*ExpressionListOpt),
				Token4:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 228:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 229:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 230:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 231:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 232:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 233:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 234:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 235:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 236:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-1].node.(*Declarator)
			d.setFull(lx)
			if k := d.Type.Kind(); k != Function {
				lx.report.Err(d.Pos(), "declarator is not a function (have '%s': %v)", d.Type, k)
			}

			for dd := d.DirectDeclarator.bottom(); dd != nil; dd = dd.parent {
				if dd.Case == 6 { // DirectDeclarator '(' ParameterTypeList ')'
					lx.scope.mergeScope = dd.paramsScope
					break
				}
			}

			// Handle __func__, [0], 6.4.2.2.
			id, _ := d.Identifier()
			lx.injectFunc = []xc.Token{
				{lex.Char{Rune: STATIC}, idStatic},
				{lex.Char{Rune: CONST}, idConst},
				{lex.Char{Rune: CHAR}, idChar},
				{lex.Char{Rune: IDENTIFIER}, idMagicFunc},
				{lex.Char{Rune: '['}, 0},
				{lex.Char{Rune: ']'}, 0},
				{lex.Char{Rune: '='}, 0},
				{lex.Char{Rune: STRINGLITERAL}, xc.Dict.SID(fmt.Sprintf("%q", xc.Dict.S(id)))},
				{lex.Char{Rune: ';'}, 0},
			}
		}
	case 237:
		{
			lx := yylex.(*lexer)
			lhs := &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				CompoundStatement:     yyS[yypt-0].node.(*CompoundStatement),
			}
			yyVAL.node = lhs
			d := lhs.Declarator
			switch dd := d.DirectDeclarator; dd.Case {
			case 6: // DirectDeclarator '(' ParameterTypeList ')'
				if o := lhs.DeclarationListOpt; o != nil {
					lx.report.Err(o.Pos(), "declaration list not allowed in a function definition with parameter type list")
				}
			case 7: // DirectDeclarator '(' IdentifierListOpt ')'
				if o1, o2 := dd.IdentifierListOpt, lhs.DeclarationListOpt; o1 != nil && o2 == nil {
					lx.report.Err(o1.Pos(), "declaration list required in a function definition without a parameter type list")
				}
			default:
				lx.report.Err(lhs.Declarator.Pos(), "invalid function definition declarator")
			}
		}
	case 238:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 239:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 240:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 241:
		{
			yyVAL.node = &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
		}
	case 242:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 243:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 244:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 245:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 246:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 247:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 248:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 249:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 250:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 251:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 252:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 253:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 254:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 255:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 256:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 257:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 258:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 259:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 260:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 261:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 262:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 263:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 264:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 265:
		{
			yyVAL.node = &ControlLine{
				Case:            1,
				Token:           yyS[yypt-4].Token,
				Token2:          yyS[yypt-3].Token,
				Token3:          yyS[yypt-2].Token,
				Token4:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 266:
		{
			yyVAL.node = &ControlLine{
				Case:            2,
				Token:           yyS[yypt-6].Token,
				Token2:          yyS[yypt-5].Token,
				IdentifierList:  yyS[yypt-4].node.(*IdentifierList).reverse(),
				Token3:          yyS[yypt-3].Token,
				Token4:          yyS[yypt-2].Token,
				Token5:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 267:
		{
			yyVAL.node = &ControlLine{
				Case:              3,
				Token:             yyS[yypt-4].Token,
				Token2:            yyS[yypt-3].Token,
				IdentifierListOpt: yyS[yypt-2].node.(*IdentifierListOpt),
				Token3:            yyS[yypt-1].Token,
				ReplacementList:   yyS[yypt-0].toks,
			}
		}
	case 268:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 269:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 270:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 271:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 272:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 273:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 274:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:            10,
				Token:           yyS[yypt-5].Token,
				Token2:          yyS[yypt-4].Token,
				Token3:          yyS[yypt-3].Token,
				Token4:          yyS[yypt-2].Token,
				Token5:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableDefineOmitCommaBeforeDDD {
				lx.report.ErrTok(lhs.Token4, "missing comma before \"...\"")
			}
		}
	case 275:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:            11,
				Token:           yyS[yypt-7].Token,
				Token2:          yyS[yypt-6].Token,
				IdentifierList:  yyS[yypt-5].node.(*IdentifierList).reverse(),
				Token3:          yyS[yypt-4].Token,
				Token4:          yyS[yypt-3].Token,
				Token5:          yyS[yypt-2].Token,
				Token6:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableDefineOmitCommaBeforeDDD {
				lx.report.ErrTok(lhs.Token6, "missing comma before \"...\"")
			}
		}
	case 276:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:   12,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableEmptyDefine {
				lx.report.ErrTok(lhs.Token2, "expected identifier")
			}
		}
	case 277:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:        13,
				Token:       yyS[yypt-3].Token,
				Token2:      yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token3:      yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableUndefExtraTokens {
				lx.report.ErrTok(decodeTokens(lhs.PPTokenList, nil)[0], "extra tokens after #undef argument")
			}
		}
	case 280:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 281:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
