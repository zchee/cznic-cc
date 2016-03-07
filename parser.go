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
	yyTabOfs   = -284
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
		41:      8,   // ')' (176x)
		59:      9,   // ';' (175x)
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
		57358:   21,  // CONST (112x)
		57413:   22,  // RESTRICT (112x)
		57431:   23,  // VOLATILE (112x)
		125:     24,  // '}' (109x)
		58:      25,  // ':' (102x)
		57351:   26,  // BOOL (102x)
		57355:   27,  // CHAR (102x)
		57357:   28,  // COMPLEX (102x)
		57365:   29,  // DOUBLE (102x)
		57367:   30,  // ENUM (102x)
		57370:   31,  // FLOAT (102x)
		57380:   32,  // INT (102x)
		57383:   33,  // LONG (102x)
		57417:   34,  // SHORT (102x)
		57418:   35,  // SIGNED (102x)
		57422:   36,  // STRUCT (102x)
		57426:   37,  // TYPEDEFNAME (102x)
		57428:   38,  // UNION (102x)
		57429:   39,  // UNSIGNED (102x)
		57430:   40,  // VOID (102x)
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
		"')'",
		"';'",
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
		"':'",
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
		118: {142, 2},
		119: {142, 2},
		120: {179, 0},
		121: {179, 1},
		122: {211, 1},
		123: {211, 3},
		124: {181, 1},
		125: {181, 3},
		126: {149, 0},
		127: {149, 1},
		128: {191, 0},
		129: {111, 7},
		130: {111, 2},
		131: {205, 1},
		132: {205, 3},
		133: {169, 1},
		134: {169, 3},
		135: {106, 1},
		136: {106, 1},
		137: {106, 1},
		138: {118, 1},
		139: {130, 2},
		140: {163, 0},
		141: {163, 1},
		142: {166, 1},
		143: {166, 3},
		144: {166, 5},
		145: {166, 6},
		146: {166, 6},
		147: {166, 5},
		148: {192, 0},
		149: {166, 5},
		150: {166, 4},
		151: {127, 2},
		152: {127, 3},
		153: {128, 0},
		154: {128, 1},
		155: {143, 1},
		156: {143, 2},
		157: {148, 0},
		158: {148, 1},
		159: {157, 1},
		160: {157, 3},
		161: {178, 0},
		162: {178, 1},
		163: {156, 1},
		164: {156, 3},
		165: {147, 2},
		166: {147, 2},
		167: {172, 1},
		168: {172, 3},
		169: {173, 0},
		170: {173, 1},
		171: {174, 0},
		172: {174, 1},
		173: {158, 0},
		174: {182, 3},
		175: {144, 1},
		176: {144, 2},
		177: {161, 0},
		178: {161, 1},
		179: {164, 3},
		180: {164, 4},
		181: {164, 5},
		182: {164, 6},
		183: {164, 6},
		184: {164, 4},
		185: {159, 0},
		186: {164, 4},
		187: {184, 0},
		188: {164, 5},
		189: {165, 0},
		190: {165, 1},
		191: {155, 1},
		192: {155, 4},
		193: {177, 2},
		194: {177, 4},
		195: {150, 2},
		196: {151, 0},
		197: {151, 1},
		198: {152, 1},
		199: {152, 2},
		200: {146, 3},
		201: {146, 2},
		202: {126, 1},
		203: {126, 1},
		204: {126, 1},
		205: {126, 1},
		206: {126, 1},
		207: {126, 1},
		208: {124, 3},
		209: {124, 4},
		210: {124, 3},
		211: {185, 0},
		212: {120, 4},
		213: {195, 1},
		214: {195, 2},
		215: {196, 0},
		216: {196, 1},
		217: {162, 1},
		218: {162, 1},
		219: {121, 2},
		220: {125, 5},
		221: {125, 7},
		222: {125, 5},
		223: {122, 5},
		224: {122, 7},
		225: {122, 9},
		226: {122, 8},
		227: {123, 3},
		228: {123, 2},
		229: {123, 2},
		230: {123, 3},
		231: {213, 1},
		232: {213, 2},
		233: {170, 1},
		234: {170, 1},
		235: {186, 0},
		236: {171, 5},
		237: {198, 1},
		238: {198, 2},
		239: {199, 0},
		240: {199, 1},
		241: {208, 1},
		242: {136, 1},
		243: {136, 2},
		244: {141, 0},
		245: {141, 1},
		246: {131, 1},
		247: {131, 1},
		248: {131, 3},
		249: {131, 1},
		250: {133, 4},
		251: {132, 4},
		252: {132, 4},
		253: {132, 4},
		254: {200, 1},
		255: {200, 2},
		256: {201, 0},
		257: {201, 1},
		258: {167, 4},
		259: {202, 3},
		260: {203, 0},
		261: {203, 1},
		262: {204, 1},
		263: {129, 3},
		264: {129, 5},
		265: {129, 7},
		266: {129, 5},
		267: {129, 2},
		268: {129, 1},
		269: {129, 3},
		270: {129, 3},
		271: {129, 2},
		272: {129, 3},
		273: {129, 6},
		274: {129, 8},
		275: {129, 2},
		276: {129, 4},
		277: {134, 1},
		278: {137, 1},
		279: {109, 1},
		280: {116, 1},
		281: {116, 2},
		282: {110, 1},
		283: {110, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 42}:   "invalid empty input",
		yyXError{487, -1}: "expected #endif",
		yyXError{489, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{405, -1}: "expected $end",
		yyXError{407, -1}: "expected $end",
		yyXError{338, -1}: "expected '('",
		yyXError{339, -1}: "expected '('",
		yyXError{340, -1}: "expected '('",
		yyXError{342, -1}: "expected '('",
		yyXError{368, -1}: "expected '('",
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
		yyXError{358, -1}: "expected ')'",
		yyXError{364, -1}: "expected ')'",
		yyXError{447, -1}: "expected ')'",
		yyXError{448, -1}: "expected ')'",
		yyXError{456, -1}: "expected ')'",
		yyXError{459, -1}: "expected ')'",
		yyXError{462, -1}: "expected ')'",
		yyXError{292, -1}: "expected ':'",
		yyXError{331, -1}: "expected ':'",
		yyXError{392, -1}: "expected ':'",
		yyXError{308, -1}: "expected ';'",
		yyXError{337, -1}: "expected ';'",
		yyXError{344, -1}: "expected ';'",
		yyXError{345, -1}: "expected ';'",
		yyXError{347, -1}: "expected ';'",
		yyXError{351, -1}: "expected ';'",
		yyXError{354, -1}: "expected ';'",
		yyXError{356, -1}: "expected ';'",
		yyXError{362, -1}: "expected ';'",
		yyXError{371, -1}: "expected ';'",
		yyXError{313, -1}: "expected '='",
		yyXError{167, -1}: "expected '['",
		yyXError{428, -1}: "expected '\\n'",
		yyXError{434, -1}: "expected '\\n'",
		yyXError{437, -1}: "expected '\\n'",
		yyXError{439, -1}: "expected '\\n'",
		yyXError{466, -1}: "expected '\\n'",
		yyXError{471, -1}: "expected '\\n'",
		yyXError{474, -1}: "expected '\\n'",
		yyXError{481, -1}: "expected '\\n'",
		yyXError{486, -1}: "expected '\\n'",
		yyXError{492, -1}: "expected '\\n'",
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
		yyXError{332, -1}: "expected '}'",
		yyXError{47, -1}:  "expected CommaOpt or one of [',', '}']",
		yyXError{244, -1}: "expected CommaOpt or one of [',', '}']",
		yyXError{259, -1}: "expected CommaOpt or one of [',', '}']",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{201, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{166, -1}: "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{334, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{314, -1}: "expected compound statement or '{'",
		yyXError{318, -1}: "expected compound statement or '{'",
		yyXError{311, -1}: "expected compound statement or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{50, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{249, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{296, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{330, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{404, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{316, -1}: "expected declaration or one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{353, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{295, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{193, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{246, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{196, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{163, -1}: "expected direct abstract declarator or one of ['(', '[']",
		yyXError{293, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{479, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{485, -1}: "expected endif line or #endif",
		yyXError{414, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{477, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{45, -1}:  "expected enumerator list or identifier",
		yyXError{277, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{73, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{97, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{369, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{373, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{377, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{381, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{468, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{408, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{75, -1}:  "expected identifier",
		yyXError{76, -1}:  "expected identifier",
		yyXError{210, -1}: "expected identifier",
		yyXError{250, -1}: "expected identifier",
		yyXError{343, -1}: "expected identifier",
		yyXError{416, -1}: "expected identifier",
		yyXError{417, -1}: "expected identifier",
		yyXError{424, -1}: "expected identifier",
		yyXError{443, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{400, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{243, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{257, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{245, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{263, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{397, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{322, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{256, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{169, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{177, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{183, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{219, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{222, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{409, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{410, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{411, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{413, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{420, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{425, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{427, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{430, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{433, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{435, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{436, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{438, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{440, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{441, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{444, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{450, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{451, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{453, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{458, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{461, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{464, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{465, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{470, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{490, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{491, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{493, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{469, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{473, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{476, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{478, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{483, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{484, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{389, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{402, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{39, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{40, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{41, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{320, -1}: "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{403, -1}: "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{35, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{36, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{171, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{179, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{324, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{325, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{326, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{327, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{328, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{329, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{348, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{349, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{350, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{352, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{360, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{366, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{372, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{376, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{380, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{384, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{386, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{387, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{391, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{394, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{396, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{333, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{335, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{336, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{388, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
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
		yyXError{301, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{12, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{38, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{303, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{304, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{305, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{306, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{307, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{238, -1}: "expected one of ['(', ')', '*', ':', '[', identifier]",
		yyXError{239, -1}: "expected one of ['(', ')', '*', ':', '[', identifier]",
		yyXError{240, -1}: "expected one of ['(', ')', '*', ':', '[', identifier]",
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
		yyXError{294, -1}: "expected one of ['(', identifier]",
		yyXError{323, -1}: "expected one of [')', ',', ';']",
		yyXError{445, -1}: "expected one of [')', ',', ...]",
		yyXError{455, -1}: "expected one of [')', ',', ...]",
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
		yyXError{370, -1}: "expected one of [')', ',']",
		yyXError{374, -1}: "expected one of [')', ',']",
		yyXError{378, -1}: "expected one of [')', ',']",
		yyXError{382, -1}: "expected one of [')', ',']",
		yyXError{446, -1}: "expected one of [')', ',']",
		yyXError{291, -1}: "expected one of [',', ':', ';']",
		yyXError{121, -1}: "expected one of [',', ':']",
		yyXError{399, -1}: "expected one of [',', ';', '=']",
		yyXError{262, -1}: "expected one of [',', ';', '}']",
		yyXError{289, -1}: "expected one of [',', ';']",
		yyXError{290, -1}: "expected one of [',', ';']",
		yyXError{297, -1}: "expected one of [',', ';']",
		yyXError{300, -1}: "expected one of [',', ';']",
		yyXError{309, -1}: "expected one of [',', ';']",
		yyXError{310, -1}: "expected one of [',', ';']",
		yyXError{398, -1}: "expected one of [',', ';']",
		yyXError{401, -1}: "expected one of [',', ';']",
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
		yyXError{418, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		yyXError{426, -1}: "expected one of ['\\n', ppother]",
		yyXError{429, -1}: "expected one of ['\\n', ppother]",
		yyXError{431, -1}: "expected one of ['\\n', ppother]",
		yyXError{315, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{317, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{32, -1}:  "expected one of ['{', identifier]",
		yyXError{33, -1}:  "expected one of ['{', identifier]",
		yyXError{287, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{298, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{302, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{454, -1}: "expected one of [..., identifier]",
		yyXError{159, -1}: "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{74, -1}:  "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{319, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{321, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{8, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{357, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{363, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{346, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{355, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{361, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{215, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{204, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{168, -1}: "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{172, -1}: "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{467, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{472, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{475, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{482, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{488, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{205, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{31, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{34, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{312, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{189, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{232, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{233, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{157, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{158, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{419, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{423, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{160, -1}: "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{228, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{206, -1}: "expected parameter type list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{236, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{406, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{442, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{449, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{452, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{457, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{460, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{463, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{156, -1}: "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{341, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{359, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{365, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{375, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{379, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{383, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{385, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{390, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{393, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{395, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{284, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{285, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{286, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{288, -1}: "expected struct declarator list or one of ['(', '*', ':', identifier]",
		yyXError{299, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{432, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{412, -1}: "expected token list or ppother",
		yyXError{415, -1}: "expected token list or ppother",
		yyXError{421, -1}: "expected token list or ppother",
		yyXError{422, -1}: "expected token list or ppother",
		yyXError{480, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{5, -1}:   "expected translation unit or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{175, -1}: "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{217, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{367, -1}: "expected while",
		yyXError{3, 42}:   "unexpected EOF",
		yyXError{2, 42}:   "unexpected EOF",
		yyXError{4, 42}:   "unexpected EOF",
	}

	yyParseTab = [494][]uint16{
		// 0
		{197: 287, 207: 286, 209: 285, 212: 288},
		{42: 284},
		{79: 283, 283, 95: 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 183: 690},
		{281, 281, 281, 281, 281, 281, 281, 281, 12: 281, 281, 281, 281, 281, 281, 281, 281, 281, 187: 688},
		{21: 279, 279, 279, 26: 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 43: 279, 279, 279, 279, 279, 188: 289},
		// 5
		{21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 106: 294, 111: 313, 315, 312, 293, 117: 291, 295, 292, 135: 325, 170: 323, 324, 213: 290},
		{21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 278, 299, 297, 322, 300, 296, 106: 294, 111: 313, 315, 312, 293, 117: 291, 295, 292, 135: 325, 170: 687, 324},
		{131, 444, 131, 9: 199, 127: 578, 577, 130: 595, 154: 593, 175: 594, 592},
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 106: 294, 111: 313, 315, 312, 293, 117: 588, 295, 292, 145: 591},
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 106: 294, 111: 313, 315, 312, 293, 117: 588, 295, 292, 145: 590},
		// 10
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 106: 294, 111: 313, 315, 312, 293, 117: 588, 295, 292, 145: 589},
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 106: 294, 111: 313, 315, 312, 293, 117: 588, 295, 292, 145: 587},
		{194, 194, 194, 8: 194, 194, 194, 194, 21: 194, 194, 194, 26: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 43: 194, 194, 194, 194, 194},
		{193, 193, 193, 8: 193, 193, 193, 193, 21: 193, 193, 193, 26: 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 43: 193, 193, 193, 193, 193},
		{192, 192, 192, 8: 192, 192, 192, 192, 21: 192, 192, 192, 26: 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 43: 192, 192, 192, 192, 192},
		// 15
		{191, 191, 191, 8: 191, 191, 191, 191, 21: 191, 191, 191, 26: 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 43: 191, 191, 191, 191, 191},
		{190, 190, 190, 8: 190, 190, 190, 190, 21: 190, 190, 190, 26: 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 43: 190, 190, 190, 190, 190},
		{189, 189, 189, 8: 189, 189, 189, 189, 21: 189, 189, 189, 25: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 43: 189, 189, 189, 189, 189},
		{188, 188, 188, 8: 188, 188, 188, 188, 21: 188, 188, 188, 25: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 43: 188, 188, 188, 188, 188},
		{187, 187, 187, 8: 187, 187, 187, 187, 21: 187, 187, 187, 25: 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 43: 187, 187, 187, 187, 187},
		// 20
		{186, 186, 186, 8: 186, 186, 186, 186, 21: 186, 186, 186, 25: 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 43: 186, 186, 186, 186, 186},
		{185, 185, 185, 8: 185, 185, 185, 185, 21: 185, 185, 185, 25: 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 43: 185, 185, 185, 185, 185},
		{184, 184, 184, 8: 184, 184, 184, 184, 21: 184, 184, 184, 25: 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 43: 184, 184, 184, 184, 184},
		{183, 183, 183, 8: 183, 183, 183, 183, 21: 183, 183, 183, 25: 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 43: 183, 183, 183, 183, 183},
		{182, 182, 182, 8: 182, 182, 182, 182, 21: 182, 182, 182, 25: 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 43: 182, 182, 182, 182, 182},
		// 25
		{181, 181, 181, 8: 181, 181, 181, 181, 21: 181, 181, 181, 25: 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 43: 181, 181, 181, 181, 181},
		{180, 180, 180, 8: 180, 180, 180, 180, 21: 180, 180, 180, 25: 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 43: 180, 180, 180, 180, 180},
		{179, 179, 179, 8: 179, 179, 179, 179, 21: 179, 179, 179, 25: 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 43: 179, 179, 179, 179, 179},
		{178, 178, 178, 8: 178, 178, 178, 178, 21: 178, 178, 178, 25: 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 43: 178, 178, 178, 178, 178},
		{177, 177, 177, 8: 177, 177, 177, 177, 21: 177, 177, 177, 25: 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 43: 177, 177, 177, 177, 177},
		// 30
		{176, 176, 176, 8: 176, 176, 176, 176, 21: 176, 176, 176, 25: 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 43: 176, 176, 176, 176, 176},
		{2: 566, 52: 113, 174: 565},
		{2: 171, 52: 171},
		{2: 170, 52: 170},
		{2: 327, 52: 113, 174: 326},
		// 35
		{149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 25: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 43: 149, 149, 149, 149, 149, 50: 149},
		{148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 25: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 43: 148, 148, 148, 148, 148, 50: 148},
		{147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 25: 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 43: 147, 147, 147, 147, 147, 50: 147},
		{146, 146, 146, 8: 146, 146, 146, 146, 21: 146, 146, 146, 26: 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 43: 146, 146, 146, 146, 146},
		{21: 53, 53, 53, 26: 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53},
		// 40
		{21: 51, 51, 51, 26: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51},
		{21: 50, 50, 50, 26: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50},
		{52: 156, 191: 328},
		{154, 154, 154, 8: 154, 154, 154, 154, 21: 154, 154, 154, 25: 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 43: 154, 154, 154, 154, 154, 52: 112},
		{52: 329},
		// 45
		{2: 330, 168: 333, 332, 205: 331},
		{10: 277, 24: 277, 48: 277},
		{10: 561, 24: 158, 149: 562},
		{10: 153, 24: 153},
		{10: 151, 24: 151, 48: 334},
		// 50
		{210, 210, 210, 210, 210, 210, 210, 210, 12: 210, 210, 210, 210, 210, 210, 210, 210, 210, 138: 335, 336},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 345},
		{10: 150, 24: 150},
		{272, 272, 3: 272, 272, 272, 272, 272, 272, 272, 272, 272, 24: 272, 272, 42: 272, 48: 272, 50: 272, 272, 53: 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272},
		{271, 271, 3: 271, 271, 271, 271, 271, 271, 271, 271, 271, 24: 271, 271, 42: 271, 48: 271, 50: 271, 271, 53: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		// 55
		{270, 270, 3: 270, 270, 270, 270, 270, 270, 270, 270, 270, 24: 270, 270, 42: 270, 48: 270, 50: 270, 270, 53: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270},
		{269, 269, 3: 269, 269, 269, 269, 269, 269, 269, 269, 269, 24: 269, 269, 42: 269, 48: 269, 50: 269, 269, 53: 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269},
		{268, 268, 3: 268, 268, 268, 268, 268, 268, 268, 268, 268, 24: 268, 268, 42: 268, 48: 268, 50: 268, 268, 53: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268},
		{267, 267, 3: 267, 267, 267, 267, 267, 267, 267, 267, 267, 24: 267, 267, 42: 267, 48: 267, 50: 267, 267, 53: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267},
		{266, 266, 3: 266, 266, 266, 266, 266, 266, 266, 266, 266, 24: 266, 266, 42: 266, 48: 266, 50: 266, 266, 53: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266},
		// 60
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 111, 111, 111, 26: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 49: 404, 107: 438, 158: 440, 182: 559},
		{358, 363, 3: 376, 366, 367, 362, 361, 9: 209, 209, 357, 24: 209, 209, 42: 209, 48: 382, 50: 209, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 558},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 557},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 556},
		// 65
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 470},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 555},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 554},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 553},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 552},
		// 70
		{355, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 356},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 111, 111, 111, 26: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 49: 404, 107: 438, 158: 440, 182: 439},
		{358, 249, 3: 249, 249, 249, 362, 361, 249, 249, 249, 357, 24: 249, 249, 42: 249, 48: 249, 50: 249, 359, 53: 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 360, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 107: 436},
		{344, 349, 337, 348, 350, 351, 347, 346, 274, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 430, 193: 431, 432},
		// 75
		{2: 429},
		{2: 428},
		{260, 260, 3: 260, 260, 260, 260, 260, 260, 260, 260, 260, 24: 260, 260, 42: 260, 48: 260, 50: 260, 260, 53: 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260},
		{259, 259, 3: 259, 259, 259, 259, 259, 259, 259, 259, 259, 24: 259, 259, 42: 259, 48: 259, 50: 259, 259, 53: 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 427},
		// 80
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 426},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 425},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 424},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 423},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 422},
		// 85
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 421},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 420},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 419},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 418},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 417},
		// 90
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 416},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 415},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 414},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 413},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 412},
		// 95
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 411},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 410},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 107: 405},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 403},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 402},
		// 100
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 401},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 400},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 399},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 398},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 397},
		// 105
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 396},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 395},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 394},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 393},
		{358, 363, 3: 376, 366, 367, 362, 361, 217, 217, 217, 357, 24: 217, 217, 42: 217, 48: 382, 50: 217, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		// 110
		{358, 363, 3: 376, 366, 367, 362, 361, 218, 218, 218, 357, 24: 218, 218, 42: 218, 48: 382, 50: 218, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{358, 363, 3: 376, 366, 367, 362, 361, 219, 219, 219, 357, 24: 219, 219, 42: 219, 48: 382, 50: 219, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{358, 363, 3: 376, 366, 367, 362, 361, 220, 220, 220, 357, 24: 220, 220, 42: 220, 48: 382, 50: 220, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{358, 363, 3: 376, 366, 367, 362, 361, 221, 221, 221, 357, 24: 221, 221, 42: 221, 48: 382, 50: 221, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{358, 363, 3: 376, 366, 367, 362, 361, 222, 222, 222, 357, 24: 222, 222, 42: 222, 48: 382, 50: 222, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		// 115
		{358, 363, 3: 376, 366, 367, 362, 361, 223, 223, 223, 357, 24: 223, 223, 42: 223, 48: 382, 50: 223, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{358, 363, 3: 376, 366, 367, 362, 361, 224, 224, 224, 357, 24: 224, 224, 42: 224, 48: 382, 50: 224, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{358, 363, 3: 376, 366, 367, 362, 361, 225, 225, 225, 357, 24: 225, 225, 42: 225, 48: 382, 50: 225, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{358, 363, 3: 376, 366, 367, 362, 361, 226, 226, 226, 357, 24: 226, 226, 42: 226, 48: 382, 50: 226, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{358, 363, 3: 376, 366, 367, 362, 361, 227, 227, 227, 357, 24: 227, 227, 42: 227, 48: 382, 50: 227, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		// 120
		{358, 363, 3: 376, 366, 367, 362, 361, 214, 214, 214, 357, 25: 214, 48: 382, 50: 214, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{10: 407, 25: 406},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 409},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 408},
		{358, 363, 3: 376, 366, 367, 362, 361, 213, 213, 213, 357, 25: 213, 48: 382, 50: 213, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		// 125
		{358, 363, 3: 376, 366, 367, 362, 361, 228, 228, 228, 357, 24: 228, 228, 42: 228, 48: 228, 50: 228, 359, 53: 365, 364, 370, 371, 381, 377, 378, 228, 379, 228, 360, 228, 374, 373, 372, 368, 228, 228, 228, 375, 228, 380, 369, 228, 228, 228},
		{358, 363, 3: 376, 366, 367, 362, 361, 229, 229, 229, 357, 24: 229, 229, 42: 229, 48: 229, 50: 229, 359, 53: 365, 364, 370, 371, 229, 377, 378, 229, 379, 229, 360, 229, 374, 373, 372, 368, 229, 229, 229, 375, 229, 229, 369, 229, 229, 229},
		{358, 363, 3: 376, 366, 367, 362, 361, 230, 230, 230, 357, 24: 230, 230, 42: 230, 48: 230, 50: 230, 359, 53: 365, 364, 370, 371, 230, 377, 378, 230, 230, 230, 360, 230, 374, 373, 372, 368, 230, 230, 230, 375, 230, 230, 369, 230, 230, 230},
		{358, 363, 3: 376, 366, 367, 362, 361, 231, 231, 231, 357, 24: 231, 231, 42: 231, 48: 231, 50: 231, 359, 53: 365, 364, 370, 371, 231, 377, 231, 231, 231, 231, 360, 231, 374, 373, 372, 368, 231, 231, 231, 375, 231, 231, 369, 231, 231, 231},
		{358, 363, 3: 376, 366, 367, 362, 361, 232, 232, 232, 357, 24: 232, 232, 42: 232, 48: 232, 50: 232, 359, 53: 365, 364, 370, 371, 232, 232, 232, 232, 232, 232, 360, 232, 374, 373, 372, 368, 232, 232, 232, 375, 232, 232, 369, 232, 232, 232},
		// 130
		{358, 363, 3: 233, 366, 367, 362, 361, 233, 233, 233, 357, 24: 233, 233, 42: 233, 48: 233, 50: 233, 359, 53: 365, 364, 370, 371, 233, 233, 233, 233, 233, 233, 360, 233, 374, 373, 372, 368, 233, 233, 233, 375, 233, 233, 369, 233, 233, 233},
		{358, 363, 3: 234, 366, 367, 362, 361, 234, 234, 234, 357, 24: 234, 234, 42: 234, 48: 234, 50: 234, 359, 53: 365, 364, 370, 371, 234, 234, 234, 234, 234, 234, 360, 234, 234, 373, 372, 368, 234, 234, 234, 234, 234, 234, 369, 234, 234, 234},
		{358, 363, 3: 235, 366, 367, 362, 361, 235, 235, 235, 357, 24: 235, 235, 42: 235, 48: 235, 50: 235, 359, 53: 365, 364, 370, 371, 235, 235, 235, 235, 235, 235, 360, 235, 235, 373, 372, 368, 235, 235, 235, 235, 235, 235, 369, 235, 235, 235},
		{358, 363, 3: 236, 366, 367, 362, 361, 236, 236, 236, 357, 24: 236, 236, 42: 236, 48: 236, 50: 236, 359, 53: 365, 364, 236, 236, 236, 236, 236, 236, 236, 236, 360, 236, 236, 236, 236, 368, 236, 236, 236, 236, 236, 236, 369, 236, 236, 236},
		{358, 363, 3: 237, 366, 367, 362, 361, 237, 237, 237, 357, 24: 237, 237, 42: 237, 48: 237, 50: 237, 359, 53: 365, 364, 237, 237, 237, 237, 237, 237, 237, 237, 360, 237, 237, 237, 237, 368, 237, 237, 237, 237, 237, 237, 369, 237, 237, 237},
		// 135
		{358, 363, 3: 238, 366, 367, 362, 361, 238, 238, 238, 357, 24: 238, 238, 42: 238, 48: 238, 50: 238, 359, 53: 365, 364, 238, 238, 238, 238, 238, 238, 238, 238, 360, 238, 238, 238, 238, 368, 238, 238, 238, 238, 238, 238, 369, 238, 238, 238},
		{358, 363, 3: 239, 366, 367, 362, 361, 239, 239, 239, 357, 24: 239, 239, 42: 239, 48: 239, 50: 239, 359, 53: 365, 364, 239, 239, 239, 239, 239, 239, 239, 239, 360, 239, 239, 239, 239, 368, 239, 239, 239, 239, 239, 239, 369, 239, 239, 239},
		{358, 363, 3: 240, 366, 367, 362, 361, 240, 240, 240, 357, 24: 240, 240, 42: 240, 48: 240, 50: 240, 359, 53: 365, 364, 240, 240, 240, 240, 240, 240, 240, 240, 360, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240},
		{358, 363, 3: 241, 366, 367, 362, 361, 241, 241, 241, 357, 24: 241, 241, 42: 241, 48: 241, 50: 241, 359, 53: 365, 364, 241, 241, 241, 241, 241, 241, 241, 241, 360, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241},
		{358, 363, 3: 242, 242, 242, 362, 361, 242, 242, 242, 357, 24: 242, 242, 42: 242, 48: 242, 50: 242, 359, 53: 365, 364, 242, 242, 242, 242, 242, 242, 242, 242, 360, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242},
		// 140
		{358, 363, 3: 243, 243, 243, 362, 361, 243, 243, 243, 357, 24: 243, 243, 42: 243, 48: 243, 50: 243, 359, 53: 365, 364, 243, 243, 243, 243, 243, 243, 243, 243, 360, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{358, 244, 3: 244, 244, 244, 362, 361, 244, 244, 244, 357, 24: 244, 244, 42: 244, 48: 244, 50: 244, 359, 53: 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 360, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244},
		{358, 245, 3: 245, 245, 245, 362, 361, 245, 245, 245, 357, 24: 245, 245, 42: 245, 48: 245, 50: 245, 359, 53: 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 360, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245},
		{358, 246, 3: 246, 246, 246, 362, 361, 246, 246, 246, 357, 24: 246, 246, 42: 246, 48: 246, 50: 246, 359, 53: 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 360, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246},
		{261, 261, 3: 261, 261, 261, 261, 261, 261, 261, 261, 261, 24: 261, 261, 42: 261, 48: 261, 50: 261, 261, 53: 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261},
		// 145
		{262, 262, 3: 262, 262, 262, 262, 262, 262, 262, 262, 262, 24: 262, 262, 42: 262, 48: 262, 50: 262, 262, 53: 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262},
		{358, 363, 3: 376, 366, 367, 362, 361, 276, 10: 276, 357, 48: 382, 51: 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{8: 273, 10: 434},
		{8: 433},
		{263, 263, 3: 263, 263, 263, 263, 263, 263, 263, 263, 263, 24: 263, 263, 42: 263, 48: 263, 50: 263, 263, 53: 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263},
		// 150
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 435},
		{358, 363, 3: 376, 366, 367, 362, 361, 275, 10: 275, 357, 48: 382, 51: 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{10: 407, 50: 437},
		{264, 264, 3: 264, 264, 264, 264, 264, 264, 264, 264, 264, 24: 264, 264, 42: 264, 48: 264, 50: 264, 264, 53: 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264},
		{8: 551, 10: 407},
		// 155
		{8: 525},
		{21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 106: 442, 111: 313, 315, 312, 441, 142: 443},
		{164, 164, 164, 8: 164, 11: 164, 21: 319, 320, 321, 25: 164, 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 106: 442, 111: 313, 315, 312, 441, 142: 523, 179: 524},
		{164, 164, 164, 8: 164, 11: 164, 21: 319, 320, 321, 25: 164, 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 106: 442, 111: 313, 315, 312, 441, 142: 523, 179: 522},
		{131, 444, 8: 107, 11: 131, 127: 445, 447, 144: 448, 161: 446},
		// 160
		{127, 127, 127, 8: 127, 10: 127, 127, 21: 319, 320, 321, 106: 455, 143: 459, 148: 520},
		{130, 2: 130, 8: 109, 10: 109, 130},
		{8: 110},
		{450, 11: 95, 164: 449, 451},
		{8: 106, 10: 106},
		// 165
		{516, 8: 108, 10: 108, 94},
		{131, 444, 8: 99, 11: 131, 21: 99, 99, 99, 26: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 43: 99, 99, 99, 99, 99, 127: 445, 447, 144: 472, 159: 473},
		{11: 452},
		{344, 454, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 319, 320, 321, 41: 458, 49: 453, 216, 106: 455, 143: 456, 153: 457},
		{358, 363, 3: 376, 366, 367, 362, 361, 11: 357, 48: 382, 50: 215, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		// 170
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 470, 471},
		{129, 129, 129, 129, 129, 129, 129, 129, 129, 10: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 41: 129, 50: 129},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 319, 320, 321, 41: 466, 49: 453, 216, 106: 463, 153: 465},
		{50: 464},
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 319, 320, 321, 106: 455, 143: 459, 148: 460},
		// 175
		{126, 126, 126, 126, 126, 126, 126, 126, 126, 10: 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 319, 320, 321, 106: 463},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 461},
		{358, 363, 3: 376, 366, 367, 362, 361, 11: 357, 48: 382, 50: 462, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{102, 8: 102, 10: 102, 102},
		{128, 128, 128, 128, 128, 128, 128, 128, 128, 10: 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 41: 128, 50: 128},
		// 180
		{104, 8: 104, 10: 104, 104},
		{50: 469},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 467},
		{358, 363, 3: 376, 366, 367, 362, 361, 11: 357, 48: 382, 50: 468, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{101, 8: 101, 10: 101, 101},
		// 185
		{103, 8: 103, 10: 103, 103},
		{358, 254, 3: 254, 254, 254, 362, 361, 254, 254, 254, 357, 24: 254, 254, 42: 254, 48: 254, 50: 254, 359, 53: 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 360, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254},
		{100, 8: 100, 10: 100, 100},
		{8: 515},
		{8: 123, 21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 106: 294, 111: 313, 315, 312, 293, 117: 477, 295, 292, 147: 476, 156: 474, 475, 178: 478},
		// 190
		{8: 125, 10: 512},
		{8: 122},
		{8: 121, 10: 121},
		{131, 444, 131, 8: 107, 10: 107, 131, 127: 445, 480, 130: 481, 144: 448, 161: 482},
		{8: 479},
		// 195
		{98, 8: 98, 10: 98, 98},
		{485, 2: 484, 11: 95, 164: 449, 451, 483},
		{8: 119, 10: 119},
		{8: 118, 10: 118},
		{489, 8: 145, 145, 145, 488, 21: 145, 145, 145, 25: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 43: 145, 145, 145, 145, 145, 145, 52: 145},
		// 200
		{142, 8: 142, 142, 142, 142, 21: 142, 142, 142, 25: 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 43: 142, 142, 142, 142, 142, 142, 52: 142},
		{131, 444, 131, 8: 99, 11: 131, 21: 99, 99, 99, 26: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 43: 99, 99, 99, 99, 99, 127: 445, 480, 130: 486, 144: 472, 159: 473},
		{8: 487},
		{141, 8: 141, 141, 141, 141, 21: 141, 141, 141, 25: 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 43: 141, 141, 141, 141, 141, 141, 52: 141},
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 319, 320, 321, 41: 500, 50: 127, 106: 455, 143: 501, 148: 499},
		// 205
		{2: 492, 8: 115, 21: 136, 136, 136, 26: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 43: 136, 136, 136, 136, 136, 172: 493, 491, 192: 490},
		{21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 106: 294, 111: 313, 315, 312, 293, 117: 477, 295, 292, 147: 476, 156: 474, 497},
		{8: 496},
		{8: 117, 10: 117},
		{8: 114, 10: 494},
		// 210
		{2: 495},
		{8: 116, 10: 116},
		{134, 8: 134, 134, 134, 134, 21: 134, 134, 134, 25: 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 43: 134, 134, 134, 134, 134, 134, 52: 134},
		{8: 498},
		{135, 8: 135, 135, 135, 135, 21: 135, 135, 135, 25: 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 43: 135, 135, 135, 135, 135, 135, 52: 135},
		// 215
		{344, 508, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 453, 216, 153: 509},
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 319, 320, 321, 106: 455, 143: 459, 148: 505},
		{126, 126, 126, 126, 126, 126, 126, 126, 12: 126, 126, 126, 126, 126, 126, 126, 126, 126, 319, 320, 321, 41: 502, 50: 126, 106: 463},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 503},
		{358, 363, 3: 376, 366, 367, 362, 361, 11: 357, 48: 382, 50: 504, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		// 220
		{138, 8: 138, 138, 138, 138, 21: 138, 138, 138, 25: 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 43: 138, 138, 138, 138, 138, 138, 52: 138},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 506},
		{358, 363, 3: 376, 366, 367, 362, 361, 11: 357, 48: 382, 50: 507, 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{139, 8: 139, 139, 139, 139, 21: 139, 139, 139, 25: 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 43: 139, 139, 139, 139, 139, 139, 52: 139},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 470, 511},
		// 225
		{50: 510},
		{140, 8: 140, 140, 140, 140, 21: 140, 140, 140, 25: 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 43: 140, 140, 140, 140, 140, 140, 52: 140},
		{137, 8: 137, 137, 137, 137, 21: 137, 137, 137, 25: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 43: 137, 137, 137, 137, 137, 137, 52: 137},
		{21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 106: 294, 111: 313, 315, 312, 293, 117: 477, 295, 292, 140: 513, 147: 514},
		{8: 124},
		// 230
		{8: 120, 10: 120},
		{105, 8: 105, 10: 105, 105},
		{8: 97, 21: 97, 97, 97, 26: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 43: 97, 97, 97, 97, 97, 184: 517},
		{8: 123, 21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 106: 294, 111: 313, 315, 312, 293, 117: 477, 295, 292, 147: 476, 156: 474, 475, 178: 518},
		{8: 519},
		// 235
		{96, 8: 96, 10: 96, 96},
		{133, 444, 133, 8: 133, 10: 133, 133, 127: 521},
		{132, 2: 132, 8: 132, 10: 132, 132},
		{165, 165, 165, 8: 165, 11: 165, 25: 165},
		{163, 163, 163, 8: 163, 11: 163, 25: 163},
		// 240
		{166, 166, 166, 8: 166, 11: 166, 25: 166},
		{344, 248, 337, 248, 248, 248, 347, 346, 248, 248, 248, 248, 353, 352, 338, 339, 340, 341, 342, 354, 343, 24: 248, 248, 42: 248, 48: 248, 526, 248, 248, 527, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248},
		{358, 247, 3: 247, 247, 247, 362, 361, 247, 247, 247, 357, 24: 247, 247, 42: 247, 48: 247, 50: 247, 359, 53: 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 360, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247},
		{88, 88, 88, 88, 88, 88, 88, 88, 11: 533, 88, 88, 88, 88, 88, 88, 88, 88, 88, 51: 534, 88, 146: 532, 150: 531, 529, 530, 177: 528},
		{10: 544, 24: 158, 149: 549},
		// 245
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 540, 52: 541, 155: 542},
		{11: 533, 48: 538, 51: 534, 146: 539},
		{87, 87, 87, 87, 87, 87, 87, 87, 12: 87, 87, 87, 87, 87, 87, 87, 87, 87, 52: 87},
		{11: 86, 48: 86, 51: 86},
		{210, 210, 210, 210, 210, 210, 210, 210, 12: 210, 210, 210, 210, 210, 210, 210, 210, 210, 138: 335, 536},
		// 250
		{2: 535},
		{11: 83, 48: 83, 51: 83},
		{50: 537},
		{11: 84, 48: 84, 51: 84},
		{89, 89, 89, 89, 89, 89, 89, 89, 12: 89, 89, 89, 89, 89, 89, 89, 89, 89, 52: 89},
		// 255
		{11: 85, 48: 85, 51: 85},
		{358, 363, 3: 376, 366, 367, 362, 361, 9: 93, 93, 357, 24: 93, 48: 382, 51: 359, 53: 365, 364, 370, 371, 381, 377, 378, 386, 379, 390, 360, 384, 374, 373, 372, 368, 388, 385, 383, 375, 392, 380, 369, 389, 387, 391},
		{88, 88, 88, 88, 88, 88, 88, 88, 11: 533, 88, 88, 88, 88, 88, 88, 88, 88, 88, 51: 534, 88, 146: 532, 150: 531, 529, 530, 177: 543},
		{10: 91, 24: 91},
		{10: 544, 24: 158, 149: 545},
		// 260
		{88, 88, 88, 88, 88, 88, 88, 88, 11: 533, 88, 88, 88, 88, 88, 88, 88, 88, 88, 24: 157, 51: 534, 88, 146: 532, 150: 531, 547, 530},
		{24: 546},
		{9: 92, 92, 24: 92},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 540, 52: 541, 155: 548},
		{10: 90, 24: 90},
		// 265
		{24: 550},
		{258, 258, 3: 258, 258, 258, 258, 258, 258, 258, 258, 258, 24: 258, 258, 42: 258, 48: 258, 50: 258, 258, 53: 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258},
		{265, 265, 3: 265, 265, 265, 265, 265, 265, 265, 265, 265, 24: 265, 265, 42: 265, 48: 265, 50: 265, 265, 53: 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265},
		{358, 250, 3: 250, 250, 250, 362, 361, 250, 250, 250, 357, 24: 250, 250, 42: 250, 48: 250, 50: 250, 359, 53: 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 360, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250},
		{358, 251, 3: 251, 251, 251, 362, 361, 251, 251, 251, 357, 24: 251, 251, 42: 251, 48: 251, 50: 251, 359, 53: 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 360, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251},
		// 270
		{358, 252, 3: 252, 252, 252, 362, 361, 252, 252, 252, 357, 24: 252, 252, 42: 252, 48: 252, 50: 252, 359, 53: 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 360, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252},
		{358, 253, 3: 253, 253, 253, 362, 361, 253, 253, 253, 357, 24: 253, 253, 42: 253, 48: 253, 50: 253, 359, 53: 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 360, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253},
		{358, 255, 3: 255, 255, 255, 362, 361, 255, 255, 255, 357, 24: 255, 255, 42: 255, 48: 255, 50: 255, 359, 53: 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 360, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{358, 256, 3: 256, 256, 256, 362, 361, 256, 256, 256, 357, 24: 256, 256, 42: 256, 48: 256, 50: 256, 359, 53: 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 360, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256},
		{358, 257, 3: 257, 257, 257, 362, 361, 257, 257, 257, 357, 24: 257, 257, 42: 257, 48: 257, 50: 257, 359, 53: 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 360, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257},
		// 275
		{8: 560},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 526, 52: 527},
		{2: 330, 24: 157, 168: 333, 564},
		{24: 563},
		{155, 155, 155, 8: 155, 155, 155, 155, 21: 155, 155, 155, 25: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 43: 155, 155, 155, 155, 155},
		// 280
		{10: 152, 24: 152},
		{52: 175, 189: 567},
		{172, 172, 172, 8: 172, 172, 172, 172, 21: 172, 172, 172, 25: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 43: 172, 172, 172, 172, 172, 52: 112},
		{52: 568},
		{21: 174, 174, 174, 26: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 190: 569},
		// 285
		{21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 106: 442, 111: 313, 315, 312, 441, 142: 572, 180: 571, 210: 570},
		{21: 319, 320, 321, 585, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 106: 442, 111: 313, 315, 312, 441, 142: 572, 180: 586},
		{21: 169, 169, 169, 169, 26: 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169},
		{131, 444, 131, 25: 144, 127: 578, 577, 130: 575, 163: 576, 181: 574, 211: 573},
		{9: 582, 583},
		// 290
		{9: 162, 162},
		{9: 160, 160, 25: 143},
		{25: 580},
		{579, 2: 484, 166: 483},
		{130, 2: 130},
		// 295
		{131, 444, 131, 127: 578, 577, 130: 486},
		{210, 210, 210, 210, 210, 210, 210, 210, 12: 210, 210, 210, 210, 210, 210, 210, 210, 210, 138: 335, 581},
		{9: 159, 159},
		{21: 167, 167, 167, 167, 26: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167},
		{131, 444, 131, 25: 144, 127: 578, 577, 130: 575, 163: 576, 181: 584},
		// 300
		{9: 161, 161},
		{173, 173, 173, 8: 173, 173, 173, 173, 21: 173, 173, 173, 25: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 43: 173, 173, 173, 173, 173},
		{21: 168, 168, 168, 168, 26: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168},
		{204, 204, 204, 8: 204, 204, 204, 204},
		{202, 202, 202, 8: 202, 202, 202, 202},
		// 305
		{205, 205, 205, 8: 205, 205, 205, 205},
		{206, 206, 206, 8: 206, 206, 206, 206},
		{207, 207, 207, 8: 207, 207, 207, 207},
		{9: 686},
		{9: 201, 201},
		// 310
		{9: 198, 684},
		{9: 197, 197, 21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 196, 52: 45, 106: 294, 111: 313, 315, 312, 293, 117: 596, 295, 292, 135: 599, 160: 597, 198: 600, 598},
		{131, 444, 131, 9: 199, 127: 578, 577, 130: 683, 154: 593, 175: 594, 592},
		{48: 681},
		{52: 49, 186: 602},
		// 315
		{21: 47, 47, 47, 26: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 43: 47, 47, 47, 47, 47, 52: 47},
		{21: 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 52: 44, 106: 294, 111: 313, 315, 312, 293, 117: 596, 295, 292, 135: 601},
		{21: 46, 46, 46, 26: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 43: 46, 46, 46, 46, 46, 52: 46},
		{52: 603, 120: 604},
		{73, 73, 73, 73, 73, 73, 73, 73, 9: 73, 12: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 26: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 43: 73, 73, 73, 73, 73, 52: 73, 82: 73, 73, 73, 73, 73, 73, 73, 73, 73, 92: 73, 73, 185: 605},
		// 320
		{21: 48, 48, 48, 26: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48},
		{344, 349, 606, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 319, 320, 321, 69, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 49: 404, 52: 603, 82: 624, 629, 614, 628, 615, 625, 626, 627, 622, 92: 630, 623, 106: 294, 607, 111: 313, 315, 312, 293, 621, 117: 596, 295, 292, 609, 610, 612, 613, 608, 611, 620, 135: 619, 162: 617, 195: 618, 616},
		{272, 272, 3: 272, 272, 272, 272, 272, 9: 272, 272, 272, 25: 679, 48: 272, 51: 272, 53: 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272},
		{8: 211, 211, 407},
		{82, 82, 82, 82, 82, 82, 82, 82, 9: 82, 12: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 26: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 43: 82, 82, 82, 82, 82, 52: 82, 82: 82, 82, 82, 82, 82, 82, 82, 82, 82, 92: 82, 82, 108: 82},
		// 325
		{81, 81, 81, 81, 81, 81, 81, 81, 9: 81, 12: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 26: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 43: 81, 81, 81, 81, 81, 52: 81, 82: 81, 81, 81, 81, 81, 81, 81, 81, 81, 92: 81, 81, 108: 81},
		{80, 80, 80, 80, 80, 80, 80, 80, 9: 80, 12: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 26: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 43: 80, 80, 80, 80, 80, 52: 80, 82: 80, 80, 80, 80, 80, 80, 80, 80, 80, 92: 80, 80, 108: 80},
		{79, 79, 79, 79, 79, 79, 79, 79, 9: 79, 12: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 26: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 43: 79, 79, 79, 79, 79, 52: 79, 82: 79, 79, 79, 79, 79, 79, 79, 79, 79, 92: 79, 79, 108: 79},
		{78, 78, 78, 78, 78, 78, 78, 78, 9: 78, 12: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 26: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 43: 78, 78, 78, 78, 78, 52: 78, 82: 78, 78, 78, 78, 78, 78, 78, 78, 78, 92: 78, 78, 108: 78},
		{77, 77, 77, 77, 77, 77, 77, 77, 9: 77, 12: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 26: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 43: 77, 77, 77, 77, 77, 52: 77, 82: 77, 77, 77, 77, 77, 77, 77, 77, 77, 92: 77, 77, 108: 77},
		// 330
		{210, 210, 210, 210, 210, 210, 210, 210, 12: 210, 210, 210, 210, 210, 210, 210, 210, 210, 138: 335, 676},
		{25: 674},
		{24: 673},
		{71, 71, 71, 71, 71, 71, 71, 71, 9: 71, 12: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 26: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 43: 71, 71, 71, 71, 71, 52: 71, 82: 71, 71, 71, 71, 71, 71, 71, 71, 71, 92: 71, 71},
		{344, 349, 606, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 319, 320, 321, 68, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 49: 404, 52: 603, 82: 624, 629, 614, 628, 615, 625, 626, 627, 622, 92: 630, 623, 106: 294, 607, 111: 313, 315, 312, 293, 621, 117: 596, 295, 292, 609, 610, 612, 613, 608, 611, 620, 135: 619, 162: 672},
		// 335
		{67, 67, 67, 67, 67, 67, 67, 67, 9: 67, 12: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 26: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 43: 67, 67, 67, 67, 67, 52: 67, 82: 67, 67, 67, 67, 67, 67, 67, 67, 67, 92: 67, 67},
		{66, 66, 66, 66, 66, 66, 66, 66, 9: 66, 12: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 26: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 43: 66, 66, 66, 66, 66, 52: 66, 82: 66, 66, 66, 66, 66, 66, 66, 66, 66, 92: 66, 66},
		{9: 671},
		{665},
		{661},
		// 340
		{657},
		{344, 349, 606, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 52: 603, 82: 624, 629, 614, 628, 615, 625, 626, 627, 622, 92: 630, 623, 107: 607, 115: 621, 120: 609, 610, 612, 613, 608, 611, 651},
		{637},
		{2: 635},
		{9: 634},
		// 345
		{9: 633},
		{344, 349, 337, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 107: 607, 115: 631},
		{9: 632},
		{54, 54, 54, 54, 54, 54, 54, 54, 9: 54, 12: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 26: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 43: 54, 54, 54, 54, 54, 52: 54, 82: 54, 54, 54, 54, 54, 54, 54, 54, 54, 92: 54, 54, 108: 54},
		{55, 55, 55, 55, 55, 55, 55, 55, 9: 55, 12: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 26: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 43: 55, 55, 55, 55, 55, 52: 55, 82: 55, 55, 55, 55, 55, 55, 55, 55, 55, 92: 55, 55, 108: 55},
		// 350
		{56, 56, 56, 56, 56, 56, 56, 56, 9: 56, 12: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 26: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 43: 56, 56, 56, 56, 56, 52: 56, 82: 56, 56, 56, 56, 56, 56, 56, 56, 56, 92: 56, 56, 108: 56},
		{9: 636},
		{57, 57, 57, 57, 57, 57, 57, 57, 9: 57, 12: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 26: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 43: 57, 57, 57, 57, 57, 52: 57, 82: 57, 57, 57, 57, 57, 57, 57, 57, 57, 92: 57, 57, 108: 57},
		{344, 349, 337, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 319, 320, 321, 26: 310, 302, 311, 307, 318, 306, 304, 305, 303, 308, 316, 314, 317, 309, 301, 298, 43: 299, 297, 322, 300, 296, 49: 404, 106: 294, 607, 111: 313, 315, 312, 293, 638, 117: 596, 295, 292, 135: 639},
		{9: 645},
		// 355
		{344, 349, 337, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 107: 607, 115: 640},
		{9: 641},
		{344, 349, 337, 348, 350, 351, 347, 346, 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 107: 607, 115: 642},
		{8: 643},
		{344, 349, 606, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 52: 603, 82: 624, 629, 614, 628, 615, 625, 626, 627, 622, 92: 630, 623, 107: 607, 115: 621, 120: 609, 610, 612, 613, 608, 611, 644},
		// 360
		{58, 58, 58, 58, 58, 58, 58, 58, 9: 58, 12: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 26: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 43: 58, 58, 58, 58, 58, 52: 58, 82: 58, 58, 58, 58, 58, 58, 58, 58, 58, 92: 58, 58, 108: 58},
		{344, 349, 337, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 107: 607, 115: 646},
		{9: 647},
		{344, 349, 337, 348, 350, 351, 347, 346, 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 107: 607, 115: 648},
		{8: 649},
		// 365
		{344, 349, 606, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 52: 603, 82: 624, 629, 614, 628, 615, 625, 626, 627, 622, 92: 630, 623, 107: 607, 115: 621, 120: 609, 610, 612, 613, 608, 611, 650},
		{59, 59, 59, 59, 59, 59, 59, 59, 9: 59, 12: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 26: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 43: 59, 59, 59, 59, 59, 52: 59, 82: 59, 59, 59, 59, 59, 59, 59, 59, 59, 92: 59, 59, 108: 59},
		{82: 652},
		{653},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 107: 654},
		// 370
		{8: 655, 10: 407},
		{9: 656},
		{60, 60, 60, 60, 60, 60, 60, 60, 9: 60, 12: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 26: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 43: 60, 60, 60, 60, 60, 52: 60, 82: 60, 60, 60, 60, 60, 60, 60, 60, 60, 92: 60, 60, 108: 60},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 107: 658},
		{8: 659, 10: 407},
		// 375
		{344, 349, 606, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 52: 603, 82: 624, 629, 614, 628, 615, 625, 626, 627, 622, 92: 630, 623, 107: 607, 115: 621, 120: 609, 610, 612, 613, 608, 611, 660},
		{61, 61, 61, 61, 61, 61, 61, 61, 9: 61, 12: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 26: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 43: 61, 61, 61, 61, 61, 52: 61, 82: 61, 61, 61, 61, 61, 61, 61, 61, 61, 92: 61, 61, 108: 61},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 107: 662},
		{8: 663, 10: 407},
		{344, 349, 606, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 52: 603, 82: 624, 629, 614, 628, 615, 625, 626, 627, 622, 92: 630, 623, 107: 607, 115: 621, 120: 609, 610, 612, 613, 608, 611, 664},
		// 380
		{62, 62, 62, 62, 62, 62, 62, 62, 9: 62, 12: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 26: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 43: 62, 62, 62, 62, 62, 52: 62, 82: 62, 62, 62, 62, 62, 62, 62, 62, 62, 92: 62, 62, 108: 62},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 107: 666},
		{8: 667, 10: 407},
		{344, 349, 606, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 52: 603, 82: 624, 629, 614, 628, 615, 625, 626, 627, 622, 92: 630, 623, 107: 607, 115: 621, 120: 609, 610, 612, 613, 608, 611, 668},
		{64, 64, 64, 64, 64, 64, 64, 64, 9: 64, 12: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 26: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 43: 64, 64, 64, 64, 64, 52: 64, 82: 64, 64, 64, 64, 64, 64, 64, 64, 64, 92: 64, 64, 108: 669},
		// 385
		{344, 349, 606, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 52: 603, 82: 624, 629, 614, 628, 615, 625, 626, 627, 622, 92: 630, 623, 107: 607, 115: 621, 120: 609, 610, 612, 613, 608, 611, 670},
		{63, 63, 63, 63, 63, 63, 63, 63, 9: 63, 12: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 26: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 43: 63, 63, 63, 63, 63, 52: 63, 82: 63, 63, 63, 63, 63, 63, 63, 63, 63, 92: 63, 63, 108: 63},
		{65, 65, 65, 65, 65, 65, 65, 65, 9: 65, 12: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 26: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 43: 65, 65, 65, 65, 65, 52: 65, 82: 65, 65, 65, 65, 65, 65, 65, 65, 65, 92: 65, 65, 108: 65},
		{70, 70, 70, 70, 70, 70, 70, 70, 9: 70, 12: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 26: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 43: 70, 70, 70, 70, 70, 52: 70, 82: 70, 70, 70, 70, 70, 70, 70, 70, 70, 92: 70, 70},
		{72, 72, 72, 72, 72, 72, 72, 72, 9: 72, 12: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 26: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 52: 72, 82: 72, 72, 72, 72, 72, 72, 72, 72, 72, 92: 72, 72, 108: 72},
		// 390
		{344, 349, 606, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 52: 603, 82: 624, 629, 614, 628, 615, 625, 626, 627, 622, 92: 630, 623, 107: 607, 115: 621, 120: 609, 610, 612, 613, 608, 611, 675},
		{74, 74, 74, 74, 74, 74, 74, 74, 9: 74, 12: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 26: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 43: 74, 74, 74, 74, 74, 52: 74, 82: 74, 74, 74, 74, 74, 74, 74, 74, 74, 92: 74, 74, 108: 74},
		{25: 677},
		{344, 349, 606, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 52: 603, 82: 624, 629, 614, 628, 615, 625, 626, 627, 622, 92: 630, 623, 107: 607, 115: 621, 120: 609, 610, 612, 613, 608, 611, 678},
		{75, 75, 75, 75, 75, 75, 75, 75, 9: 75, 12: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 26: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 43: 75, 75, 75, 75, 75, 52: 75, 82: 75, 75, 75, 75, 75, 75, 75, 75, 75, 92: 75, 75, 108: 75},
		// 395
		{344, 349, 606, 348, 350, 351, 347, 346, 9: 212, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 404, 52: 603, 82: 624, 629, 614, 628, 615, 625, 626, 627, 622, 92: 630, 623, 107: 607, 115: 621, 120: 609, 610, 612, 613, 608, 611, 680},
		{76, 76, 76, 76, 76, 76, 76, 76, 9: 76, 12: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 26: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 43: 76, 76, 76, 76, 76, 52: 76, 82: 76, 76, 76, 76, 76, 76, 76, 76, 76, 92: 76, 76, 108: 76},
		{344, 349, 337, 348, 350, 351, 347, 346, 12: 353, 352, 338, 339, 340, 341, 342, 354, 343, 49: 540, 52: 541, 155: 682},
		{9: 195, 195},
		{9: 197, 197, 48: 196, 160: 597},
		// 400
		{131, 444, 131, 127: 578, 577, 130: 683, 154: 685},
		{9: 200, 200},
		{208, 208, 208, 208, 208, 208, 208, 208, 9: 208, 12: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 26: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 52: 208, 82: 208, 208, 208, 208, 208, 208, 208, 208, 208, 92: 208, 208},
		{21: 52, 52, 52, 26: 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52},
		{210, 210, 210, 210, 210, 210, 210, 210, 12: 210, 210, 210, 210, 210, 210, 210, 210, 210, 138: 335, 689},
		// 405
		{42: 280},
		{79: 711, 713, 95: 702, 703, 704, 699, 700, 701, 705, 706, 696, 707, 708, 109: 712, 710, 116: 709, 129: 694, 131: 693, 698, 695, 697, 136: 692, 208: 691},
		{42: 282},
		{42: 43, 79: 711, 713, 95: 702, 703, 704, 699, 700, 701, 705, 706, 696, 707, 708, 109: 712, 710, 116: 709, 129: 694, 131: 754, 698, 695, 697},
		{42: 42, 79: 42, 42, 42, 91: 42, 94: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		// 410
		{42: 38, 79: 38, 38, 38, 91: 38, 94: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{42: 37, 79: 37, 37, 37, 91: 37, 94: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{80: 713, 109: 776, 710},
		{42: 35, 79: 35, 35, 35, 91: 35, 94: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{81: 28, 91: 28, 94: 764, 167: 762, 200: 763, 761},
		// 415
		{80: 713, 109: 758, 710},
		{2: 755},
		{2: 750},
		{2: 726, 79: 728, 206: 727},
		{79: 711, 713, 109: 712, 710, 116: 725},
		// 420
		{42: 16, 79: 16, 16, 16, 91: 16, 94: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{80: 713, 109: 723, 710},
		{80: 713, 109: 721, 710},
		{79: 711, 713, 109: 712, 710, 116: 720},
		{2: 716},
		// 425
		{42: 7, 79: 7, 7, 7, 91: 7, 94: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{79: 5, 715},
		{42: 4, 79: 4, 4, 4, 91: 4, 94: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{79: 714},
		{79: 2, 2},
		// 430
		{42: 3, 79: 3, 3, 3, 91: 3, 94: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{79: 1, 1},
		{79: 717, 713, 109: 718, 710},
		{42: 12, 79: 12, 12, 12, 91: 12, 94: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{79: 719},
		// 435
		{42: 8, 79: 8, 8, 8, 91: 8, 94: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{42: 13, 79: 13, 13, 13, 91: 13, 94: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{79: 722},
		{42: 14, 79: 14, 14, 14, 91: 14, 94: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{79: 724},
		// 440
		{42: 15, 79: 15, 15, 15, 91: 15, 94: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{42: 17, 79: 17, 17, 17, 91: 17, 94: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{79: 711, 713, 109: 712, 710, 116: 735, 137: 749},
		{2: 729, 8: 115, 140: 731, 172: 730, 732},
		{42: 9, 79: 9, 9, 9, 91: 9, 94: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		// 445
		{8: 117, 10: 117, 140: 746},
		{8: 114, 10: 738},
		{8: 736},
		{8: 733},
		{79: 711, 713, 109: 712, 710, 116: 735, 137: 734},
		// 450
		{42: 18, 79: 18, 18, 18, 91: 18, 94: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{42: 6, 79: 6, 6, 6, 91: 6, 94: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{79: 711, 713, 109: 712, 710, 116: 735, 137: 737},
		{42: 20, 79: 20, 20, 20, 91: 20, 94: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{2: 739, 140: 740},
		// 455
		{8: 116, 10: 116, 140: 743},
		{8: 741},
		{79: 711, 713, 109: 712, 710, 116: 735, 137: 742},
		{42: 19, 79: 19, 19, 19, 91: 19, 94: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{8: 744},
		// 460
		{79: 711, 713, 109: 712, 710, 116: 735, 137: 745},
		{42: 10, 79: 10, 10, 10, 91: 10, 94: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{8: 747},
		{79: 711, 713, 109: 712, 710, 116: 735, 137: 748},
		{42: 11, 79: 11, 11, 11, 91: 11, 94: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		// 465
		{42: 21, 79: 21, 21, 21, 91: 21, 94: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{79: 751},
		{79: 711, 713, 40, 91: 40, 94: 40, 702, 703, 704, 699, 700, 701, 705, 706, 696, 707, 708, 109: 712, 710, 116: 709, 129: 694, 131: 693, 698, 695, 697, 136: 752, 141: 753},
		{79: 711, 713, 39, 91: 39, 94: 39, 702, 703, 704, 699, 700, 701, 705, 706, 696, 707, 708, 109: 712, 710, 116: 709, 129: 694, 131: 754, 698, 695, 697},
		{81: 31, 91: 31, 94: 31},
		// 470
		{42: 41, 79: 41, 41, 41, 91: 41, 94: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		{79: 756},
		{79: 711, 713, 40, 91: 40, 94: 40, 702, 703, 704, 699, 700, 701, 705, 706, 696, 707, 708, 109: 712, 710, 116: 709, 129: 694, 131: 693, 698, 695, 697, 136: 752, 141: 757},
		{81: 32, 91: 32, 94: 32},
		{79: 759},
		// 475
		{79: 711, 713, 40, 91: 40, 94: 40, 702, 703, 704, 699, 700, 701, 705, 706, 696, 707, 708, 109: 712, 710, 116: 709, 129: 694, 131: 693, 698, 695, 697, 136: 752, 141: 760},
		{81: 33, 91: 33, 94: 33},
		{81: 24, 91: 770, 202: 771, 769},
		{81: 30, 91: 30, 94: 30},
		{81: 27, 91: 27, 94: 764, 167: 768},
		// 480
		{80: 713, 109: 765, 710},
		{79: 766},
		{79: 711, 713, 40, 91: 40, 94: 40, 702, 703, 704, 699, 700, 701, 705, 706, 696, 707, 708, 109: 712, 710, 116: 709, 129: 694, 131: 693, 698, 695, 697, 136: 752, 141: 767},
		{81: 26, 91: 26, 94: 26},
		{81: 29, 91: 29, 94: 29},
		// 485
		{81: 775, 204: 774},
		{79: 772},
		{81: 23},
		{79: 711, 713, 40, 95: 702, 703, 704, 699, 700, 701, 705, 706, 696, 707, 708, 109: 712, 710, 116: 709, 129: 694, 131: 693, 698, 695, 697, 136: 752, 141: 773},
		{81: 25},
		// 490
		{42: 34, 79: 34, 34, 34, 91: 34, 94: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{42: 22, 79: 22, 22, 22, 91: 22, 94: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{79: 777},
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
	case 119:
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
	case 120:
		{
			yyVAL.node = (*SpecifierQualifierListOpt)(nil)
		}
	case 121:
		{
			lhs := &SpecifierQualifierListOpt{
				SpecifierQualifierList: yyS[yypt-0].node.(*SpecifierQualifierList),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.SpecifierQualifierList.attr
			lhs.typeSpecifier = lhs.SpecifierQualifierList.typeSpecifier
		}
	case 122:
		{
			yyVAL.node = &StructDeclaratorList{
				StructDeclarator: yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 123:
		{
			yyVAL.node = &StructDeclaratorList{
				Case:                 1,
				StructDeclaratorList: yyS[yypt-2].node.(*StructDeclaratorList),
				Token:                yyS[yypt-1].Token,
				StructDeclarator:     yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 124:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.post(lx)
		}
	case 125:
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
	case 126:
		{
			yyVAL.node = (*CommaOpt)(nil)
		}
	case 127:
		{
			yyVAL.node = &CommaOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 128:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareEnumTag(o.Token, lx.report)
			}
			lx.iota = 0
		}
	case 129:
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
	case 130:
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
	case 131:
		{
			yyVAL.node = &EnumeratorList{
				Enumerator: yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 132:
		{
			yyVAL.node = &EnumeratorList{
				Case:           1,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList),
				Token:          yyS[yypt-1].Token,
				Enumerator:     yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 133:
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
	case 134:
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
	case 135:
		{
			lhs := &TypeQualifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saConst
		}
	case 136:
		{
			lhs := &TypeQualifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRestrict
		}
	case 137:
		{
			lhs := &TypeQualifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saVolatile
		}
	case 138:
		{
			lhs := &FunctionSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saInline
		}
	case 139:
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
	case 140:
		{
			yyVAL.node = (*DeclaratorOpt)(nil)
		}
	case 141:
		{
			yyVAL.node = &DeclaratorOpt{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
		}
	case 142:
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
	case 143:
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
	case 144:
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
	case 145:
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
	case 146:
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
	case 147:
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
	case 148:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 149:
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
	case 150:
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
	case 151:
		{
			yyVAL.node = &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].node.(*TypeQualifierListOpt),
			}
		}
	case 152:
		{
			yyVAL.node = &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].node.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].node.(*Pointer),
			}
		}
	case 153:
		{
			yyVAL.node = (*PointerOpt)(nil)
		}
	case 154:
		{
			yyVAL.node = &PointerOpt{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
		}
	case 155:
		{
			lhs := &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.TypeQualifier.attr
		}
	case 156:
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
	case 157:
		{
			yyVAL.node = (*TypeQualifierListOpt)(nil)
		}
	case 158:
		{
			yyVAL.node = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].node.(*TypeQualifierList).reverse(),
			}
		}
	case 159:
		{
			lhs := &ParameterTypeList{
				ParameterList: yyS[yypt-0].node.(*ParameterList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 160:
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
	case 161:
		{
			yyVAL.node = (*ParameterTypeListOpt)(nil)
		}
	case 162:
		{
			yyVAL.node = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].node.(*ParameterTypeList),
			}
		}
	case 163:
		{
			yyVAL.node = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 164:
		{
			yyVAL.node = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].node.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 165:
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
	case 166:
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
	case 167:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 168:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 169:
		{
			yyVAL.node = (*IdentifierListOpt)(nil)
		}
	case 170:
		{
			yyVAL.node = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 171:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 172:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 173:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeBlock)
		}
	case 174:
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
	case 175:
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
	case 176:
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
	case 177:
		{
			yyVAL.node = (*AbstractDeclaratorOpt)(nil)
		}
	case 178:
		{
			yyVAL.node = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].node.(*AbstractDeclarator),
			}
		}
	case 179:
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
	case 180:
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
	case 181:
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
	case 182:
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
	case 183:
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
	case 184:
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
	case 185:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 186:
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
	case 187:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 188:
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
	case 189:
		{
			yyVAL.node = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 190:
		{
			yyVAL.node = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
		}
	case 191:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 192:
		{
			yyVAL.node = &Initializer{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 193:
		{
			yyVAL.node = &InitializerList{
				DesignationOpt: yyS[yypt-1].node.(*DesignationOpt),
				Initializer:    yyS[yypt-0].node.(*Initializer),
			}
		}
	case 194:
		{
			yyVAL.node = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].node.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].node.(*DesignationOpt),
				Initializer:     yyS[yypt-0].node.(*Initializer),
			}
		}
	case 195:
		{
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 196:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 197:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 198:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 199:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 200:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 201:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 202:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 203:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 204:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 205:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 206:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 207:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 208:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 209:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 210:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 211:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 212:
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
	case 213:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 214:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 215:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 216:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 217:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 218:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 219:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 220:
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
	case 221:
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
	case 222:
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
	case 223:
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
	case 224:
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
	case 225:
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
	case 226:
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
	case 227:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 228:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 229:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 230:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 231:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 232:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 233:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 234:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 235:
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
	case 236:
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
	case 237:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 238:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 239:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 240:
		{
			yyVAL.node = &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
		}
	case 241:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 242:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 243:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 244:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 245:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 246:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 247:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 248:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 249:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 250:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 251:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 252:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 253:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 254:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 255:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 256:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 257:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 258:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 259:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 260:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 261:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 262:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 263:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 264:
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
	case 265:
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
	case 266:
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
	case 267:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 268:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 269:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 270:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 271:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 272:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 273:
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
	case 274:
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
	case 275:
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
	case 276:
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
	case 279:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 280:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
