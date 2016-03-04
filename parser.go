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
	"github.com/cznic/golex/lex"
	"github.com/cznic/xc"
)

type yySymType struct {
	yys       int
	Token     xc.Token
	groupPart node
	node      node
	toks      PPTokenList
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault           = 57447
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
	yyTabOfs   = -283
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (277x)
		42:      1,   // '*' (248x)
		57375:   2,   // IDENTIFIER (202x)
		38:      3,   // '&' (198x)
		43:      4,   // '+' (198x)
		45:      5,   // '-' (198x)
		57361:   6,   // DEC (198x)
		57378:   7,   // INC (198x)
		41:      8,   // ')' (176x)
		59:      9,   // ';' (175x)
		44:      10,  // ',' (168x)
		91:      11,  // '[' (149x)
		33:      12,  // '!' (131x)
		126:     13,  // '~' (131x)
		57356:   14,  // CHARCONST (131x)
		57371:   15,  // FLOATCONST (131x)
		57381:   16,  // INTCONST (131x)
		57384:   17,  // LONGCHARCONST (131x)
		57385:   18,  // LONGSTRINGLITERAL (131x)
		57419:   19,  // SIZEOF (131x)
		57421:   20,  // STRINGLITERAL (131x)
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
		57483:   49,  // Expression (87x)
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
		57533:   106, // TypeQualifier (28x)
		57484:   107, // ExpressionList (26x)
		57366:   108, // ELSE (22x)
		57507:   109, // PPTokenList (22x)
		57509:   110, // PPTokens (22x)
		57479:   111, // EnumSpecifier (20x)
		57528:   112, // StructOrUnion (20x)
		57529:   113, // StructOrUnionSpecifier (20x)
		57536:   114, // TypeSpecifier (20x)
		57485:   115, // ExpressionListOpt (18x)
		57508:   116, // PPTokenListOpt (16x)
		57462:   117, // DeclarationSpecifiers (15x)
		57490:   118, // FunctionSpecifier (15x)
		57523:   119, // StorageClassSpecifier (15x)
		57456:   120, // CompoundStatement (13x)
		57487:   121, // ExpressionStatement (12x)
		57504:   122, // IterationStatement (12x)
		57505:   123, // JumpStatement (12x)
		57506:   124, // LabeledStatement (12x)
		57518:   125, // SelectionStatement (12x)
		57522:   126, // Statement (12x)
		57514:   127, // Pointer (11x)
		57515:   128, // PointerOpt (10x)
		57458:   129, // ControlLine (8x)
		57464:   130, // Declarator (8x)
		57493:   131, // GroupPart (8x)
		57497:   132, // IfGroup (8x)
		57498:   133, // IfSection (8x)
		57530:   134, // TextLine (8x)
		57459:   135, // Declaration (7x)
		57491:   136, // GroupList (6x)
		57517:   137, // ReplacementList (6x)
		57457:   138, // ConstantExpression (5x)
		57360:   139, // DDD (5x)
		57492:   140, // GroupListOpt (5x)
		57519:   141, // SpecifierQualifierList (5x)
		57534:   142, // TypeQualifierList (5x)
		57448:   143, // AbstractDeclarator (4x)
		57463:   144, // DeclarationSpecifiersOpt (4x)
		57468:   145, // Designator (4x)
		57510:   146, // ParameterDeclaration (4x)
		57535:   147, // TypeQualifierListOpt (4x)
		57455:   148, // CommaOpt (3x)
		57466:   149, // Designation (3x)
		57467:   150, // DesignationOpt (3x)
		57469:   151, // DesignatorList (3x)
		57486:   152, // ExpressionOpt (3x)
		57499:   153, // InitDeclarator (3x)
		57502:   154, // Initializer (3x)
		57511:   155, // ParameterList (3x)
		57512:   156, // ParameterTypeList (3x)
		57435:   157, // $@10 (2x)
		57441:   158, // $@4 (2x)
		57446:   159, // $@9 (2x)
		57449:   160, // AbstractDeclaratorOpt (2x)
		57452:   161, // BlockItem (2x)
		57465:   162, // DeclaratorOpt (2x)
		57470:   163, // DirectAbstractDeclarator (2x)
		57471:   164, // DirectAbstractDeclaratorOpt (2x)
		57472:   165, // DirectDeclarator (2x)
		57473:   166, // ElifGroup (2x)
		57480:   167, // EnumerationConstant (2x)
		57481:   168, // Enumerator (2x)
		57488:   169, // ExternalDeclaration (2x)
		57489:   170, // FunctionDefinition (2x)
		57494:   171, // IdentifierList (2x)
		57495:   172, // IdentifierListOpt (2x)
		57496:   173, // IdentifierOpt (2x)
		57500:   174, // InitDeclaratorList (2x)
		57501:   175, // InitDeclaratorListOpt (2x)
		57503:   176, // InitializerList (2x)
		57513:   177, // ParameterTypeListOpt (2x)
		57520:   178, // SpecifierQualifierListOpt (2x)
		57524:   179, // StructDeclaration (2x)
		57526:   180, // StructDeclarator (2x)
		57532:   181, // TypeName (2x)
		57434:   182, // $@1 (1x)
		57436:   183, // $@11 (1x)
		57437:   184, // $@12 (1x)
		57438:   185, // $@13 (1x)
		57439:   186, // $@2 (1x)
		57440:   187, // $@3 (1x)
		57442:   188, // $@5 (1x)
		57443:   189, // $@6 (1x)
		57444:   190, // $@7 (1x)
		57445:   191, // $@8 (1x)
		57450:   192, // ArgumentExpressionList (1x)
		57451:   193, // ArgumentExpressionListOpt (1x)
		57453:   194, // BlockItemList (1x)
		57454:   195, // BlockItemListOpt (1x)
		1048577: 196, // CONSTANT_EXPRESSION (1x)
		57460:   197, // DeclarationList (1x)
		57461:   198, // DeclarationListOpt (1x)
		57474:   199, // ElifGroupList (1x)
		57475:   200, // ElifGroupListOpt (1x)
		57476:   201, // ElseGroup (1x)
		57477:   202, // ElseGroupOpt (1x)
		57478:   203, // EndifLine (1x)
		57482:   204, // EnumeratorList (1x)
		57376:   205, // IDENTIFIER_LPAREN (1x)
		1048576: 206, // PREPROCESSING_FILE (1x)
		57516:   207, // PreprocessingFile (1x)
		57521:   208, // Start (1x)
		57525:   209, // StructDeclarationList (1x)
		57527:   210, // StructDeclaratorList (1x)
		1048578: 211, // TRANSLATION_UNIT (1x)
		57531:   212, // TranslationUnit (1x)
		57447:   213, // $default (0x)
		57354:   214, // CAST (0x)
		57345:   215, // error (0x)
		57391:   216, // NOELSE (0x)
		57400:   217, // PPHEADER_NAME (0x)
		57407:   218, // PPNUMBER (0x)
		57409:   219, // PPPASTE (0x)
		57427:   220, // UNARY (0x)
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
		"$@4",
		"$@9",
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
		"$@11",
		"$@12",
		"$@13",
		"$@2",
		"$@3",
		"$@5",
		"$@6",
		"$@7",
		"$@8",
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
		1:   {182, 0},
		2:   {208, 3},
		3:   {186, 0},
		4:   {208, 3},
		5:   {187, 0},
		6:   {208, 3},
		7:   {167, 1},
		8:   {192, 1},
		9:   {192, 3},
		10:  {193, 0},
		11:  {193, 1},
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
		68:  {152, 0},
		69:  {152, 1},
		70:  {107, 1},
		71:  {107, 3},
		72:  {115, 0},
		73:  {115, 1},
		74:  {138, 1},
		75:  {135, 3},
		76:  {117, 2},
		77:  {117, 2},
		78:  {117, 2},
		79:  {117, 2},
		80:  {144, 0},
		81:  {144, 1},
		82:  {174, 1},
		83:  {174, 3},
		84:  {175, 0},
		85:  {175, 1},
		86:  {153, 1},
		87:  {158, 0},
		88:  {153, 4},
		89:  {119, 1},
		90:  {119, 1},
		91:  {119, 1},
		92:  {119, 1},
		93:  {119, 1},
		94:  {114, 1},
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
		108: {188, 0},
		109: {189, 0},
		110: {113, 7},
		111: {113, 2},
		112: {112, 1},
		113: {112, 1},
		114: {209, 1},
		115: {209, 2},
		116: {179, 3},
		117: {141, 2},
		118: {141, 2},
		119: {178, 0},
		120: {178, 1},
		121: {210, 1},
		122: {210, 3},
		123: {180, 1},
		124: {180, 3},
		125: {148, 0},
		126: {148, 1},
		127: {190, 0},
		128: {111, 7},
		129: {111, 2},
		130: {204, 1},
		131: {204, 3},
		132: {168, 1},
		133: {168, 3},
		134: {106, 1},
		135: {106, 1},
		136: {106, 1},
		137: {118, 1},
		138: {130, 2},
		139: {162, 0},
		140: {162, 1},
		141: {165, 1},
		142: {165, 3},
		143: {165, 5},
		144: {165, 6},
		145: {165, 6},
		146: {165, 5},
		147: {191, 0},
		148: {165, 5},
		149: {165, 4},
		150: {127, 2},
		151: {127, 3},
		152: {128, 0},
		153: {128, 1},
		154: {142, 1},
		155: {142, 2},
		156: {147, 0},
		157: {147, 1},
		158: {156, 1},
		159: {156, 3},
		160: {177, 0},
		161: {177, 1},
		162: {155, 1},
		163: {155, 3},
		164: {146, 2},
		165: {146, 2},
		166: {171, 1},
		167: {171, 3},
		168: {172, 0},
		169: {172, 1},
		170: {173, 0},
		171: {173, 1},
		172: {159, 0},
		173: {181, 3},
		174: {143, 1},
		175: {143, 2},
		176: {160, 0},
		177: {160, 1},
		178: {163, 3},
		179: {163, 4},
		180: {163, 5},
		181: {163, 6},
		182: {163, 6},
		183: {163, 4},
		184: {157, 0},
		185: {163, 4},
		186: {183, 0},
		187: {163, 5},
		188: {164, 0},
		189: {164, 1},
		190: {154, 1},
		191: {154, 4},
		192: {176, 2},
		193: {176, 4},
		194: {149, 2},
		195: {150, 0},
		196: {150, 1},
		197: {151, 1},
		198: {151, 2},
		199: {145, 3},
		200: {145, 2},
		201: {126, 1},
		202: {126, 1},
		203: {126, 1},
		204: {126, 1},
		205: {126, 1},
		206: {126, 1},
		207: {124, 3},
		208: {124, 4},
		209: {124, 3},
		210: {184, 0},
		211: {120, 4},
		212: {194, 1},
		213: {194, 2},
		214: {195, 0},
		215: {195, 1},
		216: {161, 1},
		217: {161, 1},
		218: {121, 2},
		219: {125, 5},
		220: {125, 7},
		221: {125, 5},
		222: {122, 5},
		223: {122, 7},
		224: {122, 9},
		225: {122, 8},
		226: {123, 3},
		227: {123, 2},
		228: {123, 2},
		229: {123, 3},
		230: {212, 1},
		231: {212, 2},
		232: {169, 1},
		233: {169, 1},
		234: {185, 0},
		235: {170, 5},
		236: {197, 1},
		237: {197, 2},
		238: {198, 0},
		239: {198, 1},
		240: {207, 1},
		241: {136, 1},
		242: {136, 2},
		243: {140, 0},
		244: {140, 1},
		245: {131, 1},
		246: {131, 1},
		247: {131, 3},
		248: {131, 1},
		249: {133, 4},
		250: {132, 4},
		251: {132, 4},
		252: {132, 4},
		253: {199, 1},
		254: {199, 2},
		255: {200, 0},
		256: {200, 1},
		257: {166, 4},
		258: {201, 3},
		259: {202, 0},
		260: {202, 1},
		261: {203, 1},
		262: {129, 3},
		263: {129, 5},
		264: {129, 7},
		265: {129, 5},
		266: {129, 2},
		267: {129, 1},
		268: {129, 3},
		269: {129, 3},
		270: {129, 2},
		271: {129, 3},
		272: {129, 6},
		273: {129, 8},
		274: {129, 2},
		275: {129, 4},
		276: {134, 1},
		277: {137, 1},
		278: {109, 1},
		279: {116, 1},
		280: {116, 2},
		281: {110, 1},
		282: {110, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 42}:   "invalid empty input",
		yyXError{486, -1}: "expected #endif",
		yyXError{488, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{404, -1}: "expected $end",
		yyXError{406, -1}: "expected $end",
		yyXError{337, -1}: "expected '('",
		yyXError{338, -1}: "expected '('",
		yyXError{339, -1}: "expected '('",
		yyXError{341, -1}: "expected '('",
		yyXError{367, -1}: "expected '('",
		yyXError{147, -1}: "expected ')'",
		yyXError{154, -1}: "expected ')'",
		yyXError{161, -1}: "expected ')'",
		yyXError{187, -1}: "expected ')'",
		yyXError{190, -1}: "expected ')'",
		yyXError{193, -1}: "expected ')'",
		yyXError{201, -1}: "expected ')'",
		yyXError{206, -1}: "expected ')'",
		yyXError{212, -1}: "expected ')'",
		yyXError{228, -1}: "expected ')'",
		yyXError{233, -1}: "expected ')'",
		yyXError{274, -1}: "expected ')'",
		yyXError{357, -1}: "expected ')'",
		yyXError{363, -1}: "expected ')'",
		yyXError{446, -1}: "expected ')'",
		yyXError{447, -1}: "expected ')'",
		yyXError{455, -1}: "expected ')'",
		yyXError{458, -1}: "expected ')'",
		yyXError{461, -1}: "expected ')'",
		yyXError{291, -1}: "expected ':'",
		yyXError{330, -1}: "expected ':'",
		yyXError{391, -1}: "expected ':'",
		yyXError{307, -1}: "expected ';'",
		yyXError{336, -1}: "expected ';'",
		yyXError{343, -1}: "expected ';'",
		yyXError{344, -1}: "expected ';'",
		yyXError{346, -1}: "expected ';'",
		yyXError{350, -1}: "expected ';'",
		yyXError{353, -1}: "expected ';'",
		yyXError{355, -1}: "expected ';'",
		yyXError{361, -1}: "expected ';'",
		yyXError{370, -1}: "expected ';'",
		yyXError{312, -1}: "expected '='",
		yyXError{166, -1}: "expected '['",
		yyXError{427, -1}: "expected '\\n'",
		yyXError{433, -1}: "expected '\\n'",
		yyXError{436, -1}: "expected '\\n'",
		yyXError{438, -1}: "expected '\\n'",
		yyXError{465, -1}: "expected '\\n'",
		yyXError{470, -1}: "expected '\\n'",
		yyXError{473, -1}: "expected '\\n'",
		yyXError{480, -1}: "expected '\\n'",
		yyXError{485, -1}: "expected '\\n'",
		yyXError{491, -1}: "expected '\\n'",
		yyXError{172, -1}: "expected ']'",
		yyXError{180, -1}: "expected ']'",
		yyXError{224, -1}: "expected ']'",
		yyXError{251, -1}: "expected ']'",
		yyXError{42, -1}:  "expected '{'",
		yyXError{44, -1}:  "expected '{'",
		yyXError{280, -1}: "expected '{'",
		yyXError{282, -1}: "expected '{'",
		yyXError{260, -1}: "expected '}'",
		yyXError{264, -1}: "expected '}'",
		yyXError{277, -1}: "expected '}'",
		yyXError{331, -1}: "expected '}'",
		yyXError{47, -1}:  "expected CommaOpt or one of [',', '}']",
		yyXError{243, -1}: "expected CommaOpt or one of [',', '}']",
		yyXError{258, -1}: "expected CommaOpt or one of [',', '}']",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{200, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{165, -1}: "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{333, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{313, -1}: "expected compound statement or '{'",
		yyXError{317, -1}: "expected compound statement or '{'",
		yyXError{310, -1}: "expected compound statement or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{50, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{248, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{295, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{329, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{403, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{315, -1}: "expected declaration or one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{352, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{294, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{192, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{245, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{195, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{162, -1}: "expected direct abstract declarator or one of ['(', '[']",
		yyXError{292, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{478, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{484, -1}: "expected endif line or #endif",
		yyXError{413, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{476, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{45, -1}:  "expected enumerator list or identifier",
		yyXError{276, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{72, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{96, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{368, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{372, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{376, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{380, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{58, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{70, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{240, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		yyXError{169, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{223, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{275, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{60, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{61, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{62, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{63, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{64, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{65, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{66, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{67, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{68, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{78, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{97, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{121, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{122, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{149, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{175, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{181, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{217, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{220, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{173, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{215, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{6, -1}:   "expected external declaration or one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{467, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{407, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{74, -1}:  "expected identifier",
		yyXError{75, -1}:  "expected identifier",
		yyXError{209, -1}: "expected identifier",
		yyXError{249, -1}: "expected identifier",
		yyXError{342, -1}: "expected identifier",
		yyXError{415, -1}: "expected identifier",
		yyXError{416, -1}: "expected identifier",
		yyXError{423, -1}: "expected identifier",
		yyXError{442, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{399, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{242, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{256, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{244, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{262, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{396, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{259, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{51, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{52, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{53, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{54, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{55, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{56, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{57, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{71, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{76, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{77, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{108, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
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
		yyXError{124, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
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
		yyXError{148, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{152, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{185, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{241, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{265, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{266, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{267, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{268, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{269, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{270, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{271, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{272, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{273, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{59, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{119, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{123, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{145, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{150, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{321, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{255, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{168, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{176, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{182, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{218, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{221, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{408, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{409, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{410, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{412, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{419, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{424, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{426, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{429, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{432, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{434, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{435, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{437, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{439, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{440, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{443, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{449, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{450, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{452, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{457, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{460, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{463, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{464, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{469, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{489, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{490, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{492, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{468, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{472, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{475, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{477, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{482, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{483, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{388, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{401, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{39, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{40, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{41, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{319, -1}: "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{402, -1}: "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{35, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{36, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{170, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{178, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{323, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{324, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{325, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{326, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{327, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{328, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{347, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{348, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{349, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{351, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{359, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{365, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{371, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{375, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{379, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{383, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{385, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{386, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{390, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{393, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{395, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{332, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{334, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{335, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{387, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{246, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{253, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{43, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{281, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
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
		yyXError{278, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{300, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{12, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{38, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{302, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{303, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{304, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{305, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{306, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{237, -1}: "expected one of ['(', ')', '*', ':', '[', identifier]",
		yyXError{238, -1}: "expected one of ['(', ')', '*', ':', '[', identifier]",
		yyXError{239, -1}: "expected one of ['(', ')', '*', ':', '[', identifier]",
		yyXError{198, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{199, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{202, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{211, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{213, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{219, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{222, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{225, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{226, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{160, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{236, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{164, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{177, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{179, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{183, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{184, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{186, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{194, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{230, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{234, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{293, -1}: "expected one of ['(', identifier]",
		yyXError{322, -1}: "expected one of [')', ',', ';']",
		yyXError{444, -1}: "expected one of [')', ',', ...]",
		yyXError{454, -1}: "expected one of [')', ',', ...]",
		yyXError{146, -1}: "expected one of [')', ',']",
		yyXError{153, -1}: "expected one of [')', ',']",
		yyXError{163, -1}: "expected one of [')', ',']",
		yyXError{189, -1}: "expected one of [')', ',']",
		yyXError{191, -1}: "expected one of [')', ',']",
		yyXError{196, -1}: "expected one of [')', ',']",
		yyXError{197, -1}: "expected one of [')', ',']",
		yyXError{207, -1}: "expected one of [')', ',']",
		yyXError{208, -1}: "expected one of [')', ',']",
		yyXError{210, -1}: "expected one of [')', ',']",
		yyXError{229, -1}: "expected one of [')', ',']",
		yyXError{369, -1}: "expected one of [')', ',']",
		yyXError{373, -1}: "expected one of [')', ',']",
		yyXError{377, -1}: "expected one of [')', ',']",
		yyXError{381, -1}: "expected one of [')', ',']",
		yyXError{445, -1}: "expected one of [')', ',']",
		yyXError{290, -1}: "expected one of [',', ':', ';']",
		yyXError{120, -1}: "expected one of [',', ':']",
		yyXError{398, -1}: "expected one of [',', ';', '=']",
		yyXError{261, -1}: "expected one of [',', ';', '}']",
		yyXError{288, -1}: "expected one of [',', ';']",
		yyXError{289, -1}: "expected one of [',', ';']",
		yyXError{296, -1}: "expected one of [',', ';']",
		yyXError{299, -1}: "expected one of [',', ';']",
		yyXError{308, -1}: "expected one of [',', ';']",
		yyXError{309, -1}: "expected one of [',', ';']",
		yyXError{397, -1}: "expected one of [',', ';']",
		yyXError{400, -1}: "expected one of [',', ';']",
		yyXError{46, -1}:  "expected one of [',', '=', '}']",
		yyXError{49, -1}:  "expected one of [',', '=', '}']",
		yyXError{151, -1}: "expected one of [',', ']']",
		yyXError{48, -1}:  "expected one of [',', '}']",
		yyXError{69, -1}:  "expected one of [',', '}']",
		yyXError{257, -1}: "expected one of [',', '}']",
		yyXError{263, -1}: "expected one of [',', '}']",
		yyXError{279, -1}: "expected one of [',', '}']",
		yyXError{247, -1}: "expected one of ['.', '=', '[']",
		yyXError{250, -1}: "expected one of ['.', '=', '[']",
		yyXError{252, -1}: "expected one of ['.', '=', '[']",
		yyXError{254, -1}: "expected one of ['.', '=', '[']",
		yyXError{417, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		yyXError{425, -1}: "expected one of ['\\n', ppother]",
		yyXError{428, -1}: "expected one of ['\\n', ppother]",
		yyXError{430, -1}: "expected one of ['\\n', ppother]",
		yyXError{314, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{316, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{32, -1}:  "expected one of ['{', identifier]",
		yyXError{33, -1}:  "expected one of ['{', identifier]",
		yyXError{286, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{297, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{301, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{453, -1}: "expected one of [..., identifier]",
		yyXError{158, -1}: "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{73, -1}:  "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{318, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{320, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{8, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{356, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{362, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{345, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{354, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{360, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{214, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{203, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{167, -1}: "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{171, -1}: "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{466, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{471, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{474, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{481, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{487, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{204, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{31, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{34, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{311, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{188, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{231, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{232, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{156, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{157, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{418, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{422, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{159, -1}: "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{227, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{205, -1}: "expected parameter type list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{235, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{405, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{441, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{448, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{451, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{456, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{459, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{462, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{155, -1}: "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{340, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{358, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{364, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{374, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{378, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{382, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{384, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{389, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{392, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{394, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{283, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{284, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{285, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{287, -1}: "expected struct declarator list or one of ['(', '*', ':', identifier]",
		yyXError{298, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{431, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{411, -1}: "expected token list or ppother",
		yyXError{414, -1}: "expected token list or ppother",
		yyXError{420, -1}: "expected token list or ppother",
		yyXError{421, -1}: "expected token list or ppother",
		yyXError{479, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{5, -1}:   "expected translation unit or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{174, -1}: "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{216, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{366, -1}: "expected while",
		yyXError{3, 42}:   "unexpected EOF",
		yyXError{2, 42}:   "unexpected EOF",
		yyXError{4, 42}:   "unexpected EOF",
	}

	yyParseTab = [493][]uint16{
		// 0
		{196: 286, 206: 285, 208: 284, 211: 287},
		{42: 283},
		{79: 282, 282, 95: 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 182: 688},
		{280, 280, 280, 280, 280, 280, 280, 280, 12: 280, 280, 280, 280, 280, 280, 280, 280, 280, 186: 686},
		{21: 278, 278, 278, 26: 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 43: 278, 278, 278, 278, 278, 187: 288},
		// 5
		{21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 106: 293, 111: 312, 314, 311, 292, 117: 290, 294, 291, 135: 324, 169: 322, 323, 212: 289},
		{21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 277, 298, 296, 321, 299, 295, 106: 293, 111: 312, 314, 311, 292, 117: 290, 294, 291, 135: 324, 169: 685, 323},
		{131, 442, 131, 9: 199, 127: 576, 575, 130: 593, 153: 591, 174: 592, 590},
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 106: 293, 111: 312, 314, 311, 292, 117: 586, 294, 291, 144: 589},
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 106: 293, 111: 312, 314, 311, 292, 117: 586, 294, 291, 144: 588},
		// 10
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 106: 293, 111: 312, 314, 311, 292, 117: 586, 294, 291, 144: 587},
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 106: 293, 111: 312, 314, 311, 292, 117: 586, 294, 291, 144: 585},
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
		{2: 564, 52: 113, 173: 563},
		{2: 171, 52: 171},
		{2: 170, 52: 170},
		{2: 326, 52: 113, 173: 325},
		// 35
		{149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 25: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 43: 149, 149, 149, 149, 149, 50: 149},
		{148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 25: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 43: 148, 148, 148, 148, 148, 50: 148},
		{147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 25: 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 43: 147, 147, 147, 147, 147, 50: 147},
		{146, 146, 146, 8: 146, 146, 146, 146, 21: 146, 146, 146, 26: 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 43: 146, 146, 146, 146, 146},
		{21: 53, 53, 53, 26: 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53},
		// 40
		{21: 51, 51, 51, 26: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51},
		{21: 50, 50, 50, 26: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50},
		{52: 156, 190: 327},
		{154, 154, 154, 8: 154, 154, 154, 154, 21: 154, 154, 154, 25: 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 43: 154, 154, 154, 154, 154, 52: 112},
		{52: 328},
		// 45
		{2: 329, 167: 332, 331, 204: 330},
		{10: 276, 24: 276, 48: 276},
		{10: 559, 24: 158, 148: 560},
		{10: 153, 24: 153},
		{10: 151, 24: 151, 48: 333},
		// 50
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 342, 138: 352},
		{271, 271, 3: 271, 271, 271, 271, 271, 271, 271, 271, 271, 24: 271, 271, 42: 271, 48: 271, 50: 271, 271, 53: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{270, 270, 3: 270, 270, 270, 270, 270, 270, 270, 270, 270, 24: 270, 270, 42: 270, 48: 270, 50: 270, 270, 53: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270},
		{269, 269, 3: 269, 269, 269, 269, 269, 269, 269, 269, 269, 24: 269, 269, 42: 269, 48: 269, 50: 269, 269, 53: 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269},
		{268, 268, 3: 268, 268, 268, 268, 268, 268, 268, 268, 268, 24: 268, 268, 42: 268, 48: 268, 50: 268, 268, 53: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268},
		// 55
		{267, 267, 3: 267, 267, 267, 267, 267, 267, 267, 267, 267, 24: 267, 267, 42: 267, 48: 267, 50: 267, 267, 53: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267},
		{266, 266, 3: 266, 266, 266, 266, 266, 266, 266, 266, 266, 24: 266, 266, 42: 266, 48: 266, 50: 266, 266, 53: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266},
		{265, 265, 3: 265, 265, 265, 265, 265, 265, 265, 265, 265, 24: 265, 265, 42: 265, 48: 265, 50: 265, 265, 53: 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 111, 111, 111, 26: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 49: 402, 107: 436, 159: 438, 181: 557},
		{356, 361, 3: 374, 364, 365, 360, 359, 9: 209, 209, 355, 24: 209, 209, 42: 209, 48: 380, 50: 209, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		// 60
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 556},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 555},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 554},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 468},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 553},
		// 65
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 552},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 551},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 550},
		{353, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 354},
		{10: 150, 24: 150},
		// 70
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 111, 111, 111, 26: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 49: 402, 107: 436, 159: 438, 181: 437},
		{356, 248, 3: 248, 248, 248, 360, 359, 248, 248, 248, 355, 24: 248, 248, 42: 248, 48: 248, 50: 248, 357, 53: 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 358, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 107: 434},
		{341, 346, 334, 345, 347, 348, 344, 343, 273, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 428, 192: 429, 430},
		{2: 427},
		// 75
		{2: 426},
		{259, 259, 3: 259, 259, 259, 259, 259, 259, 259, 259, 259, 24: 259, 259, 42: 259, 48: 259, 50: 259, 259, 53: 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259},
		{258, 258, 3: 258, 258, 258, 258, 258, 258, 258, 258, 258, 24: 258, 258, 42: 258, 48: 258, 50: 258, 258, 53: 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 425},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 424},
		// 80
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 423},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 422},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 421},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 420},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 419},
		// 85
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 418},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 417},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 416},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 415},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 414},
		// 90
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 413},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 412},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 411},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 410},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 409},
		// 95
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 408},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 107: 403},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 401},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 400},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 399},
		// 100
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 398},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 397},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 396},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 395},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 394},
		// 105
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 393},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 392},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 391},
		{356, 361, 3: 374, 364, 365, 360, 359, 216, 216, 216, 355, 24: 216, 216, 42: 216, 48: 380, 50: 216, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{356, 361, 3: 374, 364, 365, 360, 359, 217, 217, 217, 355, 24: 217, 217, 42: 217, 48: 380, 50: 217, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		// 110
		{356, 361, 3: 374, 364, 365, 360, 359, 218, 218, 218, 355, 24: 218, 218, 42: 218, 48: 380, 50: 218, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{356, 361, 3: 374, 364, 365, 360, 359, 219, 219, 219, 355, 24: 219, 219, 42: 219, 48: 380, 50: 219, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{356, 361, 3: 374, 364, 365, 360, 359, 220, 220, 220, 355, 24: 220, 220, 42: 220, 48: 380, 50: 220, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{356, 361, 3: 374, 364, 365, 360, 359, 221, 221, 221, 355, 24: 221, 221, 42: 221, 48: 380, 50: 221, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{356, 361, 3: 374, 364, 365, 360, 359, 222, 222, 222, 355, 24: 222, 222, 42: 222, 48: 380, 50: 222, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		// 115
		{356, 361, 3: 374, 364, 365, 360, 359, 223, 223, 223, 355, 24: 223, 223, 42: 223, 48: 380, 50: 223, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{356, 361, 3: 374, 364, 365, 360, 359, 224, 224, 224, 355, 24: 224, 224, 42: 224, 48: 380, 50: 224, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{356, 361, 3: 374, 364, 365, 360, 359, 225, 225, 225, 355, 24: 225, 225, 42: 225, 48: 380, 50: 225, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{356, 361, 3: 374, 364, 365, 360, 359, 226, 226, 226, 355, 24: 226, 226, 42: 226, 48: 380, 50: 226, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{356, 361, 3: 374, 364, 365, 360, 359, 213, 213, 213, 355, 25: 213, 48: 380, 50: 213, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		// 120
		{10: 405, 25: 404},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 407},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 406},
		{356, 361, 3: 374, 364, 365, 360, 359, 212, 212, 212, 355, 25: 212, 48: 380, 50: 212, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{356, 361, 3: 374, 364, 365, 360, 359, 227, 227, 227, 355, 24: 227, 227, 42: 227, 48: 227, 50: 227, 357, 53: 363, 362, 368, 369, 379, 375, 376, 227, 377, 227, 358, 227, 372, 371, 370, 366, 227, 227, 227, 373, 227, 378, 367, 227, 227, 227},
		// 125
		{356, 361, 3: 374, 364, 365, 360, 359, 228, 228, 228, 355, 24: 228, 228, 42: 228, 48: 228, 50: 228, 357, 53: 363, 362, 368, 369, 228, 375, 376, 228, 377, 228, 358, 228, 372, 371, 370, 366, 228, 228, 228, 373, 228, 228, 367, 228, 228, 228},
		{356, 361, 3: 374, 364, 365, 360, 359, 229, 229, 229, 355, 24: 229, 229, 42: 229, 48: 229, 50: 229, 357, 53: 363, 362, 368, 369, 229, 375, 376, 229, 229, 229, 358, 229, 372, 371, 370, 366, 229, 229, 229, 373, 229, 229, 367, 229, 229, 229},
		{356, 361, 3: 374, 364, 365, 360, 359, 230, 230, 230, 355, 24: 230, 230, 42: 230, 48: 230, 50: 230, 357, 53: 363, 362, 368, 369, 230, 375, 230, 230, 230, 230, 358, 230, 372, 371, 370, 366, 230, 230, 230, 373, 230, 230, 367, 230, 230, 230},
		{356, 361, 3: 374, 364, 365, 360, 359, 231, 231, 231, 355, 24: 231, 231, 42: 231, 48: 231, 50: 231, 357, 53: 363, 362, 368, 369, 231, 231, 231, 231, 231, 231, 358, 231, 372, 371, 370, 366, 231, 231, 231, 373, 231, 231, 367, 231, 231, 231},
		{356, 361, 3: 232, 364, 365, 360, 359, 232, 232, 232, 355, 24: 232, 232, 42: 232, 48: 232, 50: 232, 357, 53: 363, 362, 368, 369, 232, 232, 232, 232, 232, 232, 358, 232, 372, 371, 370, 366, 232, 232, 232, 373, 232, 232, 367, 232, 232, 232},
		// 130
		{356, 361, 3: 233, 364, 365, 360, 359, 233, 233, 233, 355, 24: 233, 233, 42: 233, 48: 233, 50: 233, 357, 53: 363, 362, 368, 369, 233, 233, 233, 233, 233, 233, 358, 233, 233, 371, 370, 366, 233, 233, 233, 233, 233, 233, 367, 233, 233, 233},
		{356, 361, 3: 234, 364, 365, 360, 359, 234, 234, 234, 355, 24: 234, 234, 42: 234, 48: 234, 50: 234, 357, 53: 363, 362, 368, 369, 234, 234, 234, 234, 234, 234, 358, 234, 234, 371, 370, 366, 234, 234, 234, 234, 234, 234, 367, 234, 234, 234},
		{356, 361, 3: 235, 364, 365, 360, 359, 235, 235, 235, 355, 24: 235, 235, 42: 235, 48: 235, 50: 235, 357, 53: 363, 362, 235, 235, 235, 235, 235, 235, 235, 235, 358, 235, 235, 235, 235, 366, 235, 235, 235, 235, 235, 235, 367, 235, 235, 235},
		{356, 361, 3: 236, 364, 365, 360, 359, 236, 236, 236, 355, 24: 236, 236, 42: 236, 48: 236, 50: 236, 357, 53: 363, 362, 236, 236, 236, 236, 236, 236, 236, 236, 358, 236, 236, 236, 236, 366, 236, 236, 236, 236, 236, 236, 367, 236, 236, 236},
		{356, 361, 3: 237, 364, 365, 360, 359, 237, 237, 237, 355, 24: 237, 237, 42: 237, 48: 237, 50: 237, 357, 53: 363, 362, 237, 237, 237, 237, 237, 237, 237, 237, 358, 237, 237, 237, 237, 366, 237, 237, 237, 237, 237, 237, 367, 237, 237, 237},
		// 135
		{356, 361, 3: 238, 364, 365, 360, 359, 238, 238, 238, 355, 24: 238, 238, 42: 238, 48: 238, 50: 238, 357, 53: 363, 362, 238, 238, 238, 238, 238, 238, 238, 238, 358, 238, 238, 238, 238, 366, 238, 238, 238, 238, 238, 238, 367, 238, 238, 238},
		{356, 361, 3: 239, 364, 365, 360, 359, 239, 239, 239, 355, 24: 239, 239, 42: 239, 48: 239, 50: 239, 357, 53: 363, 362, 239, 239, 239, 239, 239, 239, 239, 239, 358, 239, 239, 239, 239, 239, 239, 239, 239, 239, 239, 239, 239, 239, 239, 239},
		{356, 361, 3: 240, 364, 365, 360, 359, 240, 240, 240, 355, 24: 240, 240, 42: 240, 48: 240, 50: 240, 357, 53: 363, 362, 240, 240, 240, 240, 240, 240, 240, 240, 358, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240},
		{356, 361, 3: 241, 241, 241, 360, 359, 241, 241, 241, 355, 24: 241, 241, 42: 241, 48: 241, 50: 241, 357, 53: 363, 362, 241, 241, 241, 241, 241, 241, 241, 241, 358, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241},
		{356, 361, 3: 242, 242, 242, 360, 359, 242, 242, 242, 355, 24: 242, 242, 42: 242, 48: 242, 50: 242, 357, 53: 363, 362, 242, 242, 242, 242, 242, 242, 242, 242, 358, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242},
		// 140
		{356, 243, 3: 243, 243, 243, 360, 359, 243, 243, 243, 355, 24: 243, 243, 42: 243, 48: 243, 50: 243, 357, 53: 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 358, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{356, 244, 3: 244, 244, 244, 360, 359, 244, 244, 244, 355, 24: 244, 244, 42: 244, 48: 244, 50: 244, 357, 53: 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 358, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244},
		{356, 245, 3: 245, 245, 245, 360, 359, 245, 245, 245, 355, 24: 245, 245, 42: 245, 48: 245, 50: 245, 357, 53: 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 358, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245},
		{260, 260, 3: 260, 260, 260, 260, 260, 260, 260, 260, 260, 24: 260, 260, 42: 260, 48: 260, 50: 260, 260, 53: 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260},
		{261, 261, 3: 261, 261, 261, 261, 261, 261, 261, 261, 261, 24: 261, 261, 42: 261, 48: 261, 50: 261, 261, 53: 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261},
		// 145
		{356, 361, 3: 374, 364, 365, 360, 359, 275, 10: 275, 355, 48: 380, 51: 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{8: 272, 10: 432},
		{8: 431},
		{262, 262, 3: 262, 262, 262, 262, 262, 262, 262, 262, 262, 24: 262, 262, 42: 262, 48: 262, 50: 262, 262, 53: 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 433},
		// 150
		{356, 361, 3: 374, 364, 365, 360, 359, 274, 10: 274, 355, 48: 380, 51: 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{10: 405, 50: 435},
		{263, 263, 3: 263, 263, 263, 263, 263, 263, 263, 263, 263, 24: 263, 263, 42: 263, 48: 263, 50: 263, 263, 53: 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263},
		{8: 549, 10: 405},
		{8: 523},
		// 155
		{21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 106: 440, 111: 312, 314, 311, 439, 141: 441},
		{164, 164, 164, 8: 164, 11: 164, 21: 318, 319, 320, 25: 164, 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 106: 440, 111: 312, 314, 311, 439, 141: 521, 178: 522},
		{164, 164, 164, 8: 164, 11: 164, 21: 318, 319, 320, 25: 164, 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 106: 440, 111: 312, 314, 311, 439, 141: 521, 178: 520},
		{131, 442, 8: 107, 11: 131, 127: 443, 445, 143: 446, 160: 444},
		{127, 127, 127, 8: 127, 10: 127, 127, 21: 318, 319, 320, 106: 453, 142: 457, 147: 518},
		// 160
		{130, 2: 130, 8: 109, 10: 109, 130},
		{8: 110},
		{448, 11: 95, 163: 447, 449},
		{8: 106, 10: 106},
		{514, 8: 108, 10: 108, 94},
		// 165
		{131, 442, 8: 99, 11: 131, 21: 99, 99, 99, 26: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 43: 99, 99, 99, 99, 99, 127: 443, 445, 143: 470, 157: 471},
		{11: 450},
		{341, 452, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 318, 319, 320, 41: 456, 49: 451, 215, 106: 453, 142: 454, 152: 455},
		{356, 361, 3: 374, 364, 365, 360, 359, 11: 355, 48: 380, 50: 214, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 468, 469},
		// 170
		{129, 129, 129, 129, 129, 129, 129, 129, 129, 10: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 41: 129, 50: 129},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 318, 319, 320, 41: 464, 49: 451, 215, 106: 461, 152: 463},
		{50: 462},
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 318, 319, 320, 106: 453, 142: 457, 147: 458},
		{126, 126, 126, 126, 126, 126, 126, 126, 126, 10: 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 318, 319, 320, 106: 461},
		// 175
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 459},
		{356, 361, 3: 374, 364, 365, 360, 359, 11: 355, 48: 380, 50: 460, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{102, 8: 102, 10: 102, 102},
		{128, 128, 128, 128, 128, 128, 128, 128, 128, 10: 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 41: 128, 50: 128},
		{104, 8: 104, 10: 104, 104},
		// 180
		{50: 467},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 465},
		{356, 361, 3: 374, 364, 365, 360, 359, 11: 355, 48: 380, 50: 466, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{101, 8: 101, 10: 101, 101},
		{103, 8: 103, 10: 103, 103},
		// 185
		{356, 253, 3: 253, 253, 253, 360, 359, 253, 253, 253, 355, 24: 253, 253, 42: 253, 48: 253, 50: 253, 357, 53: 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 358, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253},
		{100, 8: 100, 10: 100, 100},
		{8: 513},
		{8: 123, 21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 106: 293, 111: 312, 314, 311, 292, 117: 475, 294, 291, 146: 474, 155: 472, 473, 177: 476},
		{8: 125, 10: 510},
		// 190
		{8: 122},
		{8: 121, 10: 121},
		{131, 442, 131, 8: 107, 10: 107, 131, 127: 443, 478, 130: 479, 143: 446, 160: 480},
		{8: 477},
		{98, 8: 98, 10: 98, 98},
		// 195
		{483, 2: 482, 11: 95, 163: 447, 449, 481},
		{8: 119, 10: 119},
		{8: 118, 10: 118},
		{487, 8: 145, 145, 145, 486, 21: 145, 145, 145, 25: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 43: 145, 145, 145, 145, 145, 145, 52: 145},
		{142, 8: 142, 142, 142, 142, 21: 142, 142, 142, 25: 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 43: 142, 142, 142, 142, 142, 142, 52: 142},
		// 200
		{131, 442, 131, 8: 99, 11: 131, 21: 99, 99, 99, 26: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 43: 99, 99, 99, 99, 99, 127: 443, 478, 130: 484, 143: 470, 157: 471},
		{8: 485},
		{141, 8: 141, 141, 141, 141, 21: 141, 141, 141, 25: 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 43: 141, 141, 141, 141, 141, 141, 52: 141},
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 318, 319, 320, 41: 498, 50: 127, 106: 453, 142: 499, 147: 497},
		{2: 490, 8: 115, 21: 136, 136, 136, 26: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 43: 136, 136, 136, 136, 136, 171: 491, 489, 191: 488},
		// 205
		{21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 106: 293, 111: 312, 314, 311, 292, 117: 475, 294, 291, 146: 474, 155: 472, 495},
		{8: 494},
		{8: 117, 10: 117},
		{8: 114, 10: 492},
		{2: 493},
		// 210
		{8: 116, 10: 116},
		{134, 8: 134, 134, 134, 134, 21: 134, 134, 134, 25: 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 43: 134, 134, 134, 134, 134, 134, 52: 134},
		{8: 496},
		{135, 8: 135, 135, 135, 135, 21: 135, 135, 135, 25: 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 43: 135, 135, 135, 135, 135, 135, 52: 135},
		{341, 506, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 451, 215, 152: 507},
		// 215
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 318, 319, 320, 106: 453, 142: 457, 147: 503},
		{126, 126, 126, 126, 126, 126, 126, 126, 12: 126, 126, 126, 126, 126, 126, 126, 126, 126, 318, 319, 320, 41: 500, 50: 126, 106: 461},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 501},
		{356, 361, 3: 374, 364, 365, 360, 359, 11: 355, 48: 380, 50: 502, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{138, 8: 138, 138, 138, 138, 21: 138, 138, 138, 25: 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 43: 138, 138, 138, 138, 138, 138, 52: 138},
		// 220
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 504},
		{356, 361, 3: 374, 364, 365, 360, 359, 11: 355, 48: 380, 50: 505, 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{139, 8: 139, 139, 139, 139, 21: 139, 139, 139, 25: 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 43: 139, 139, 139, 139, 139, 139, 52: 139},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 468, 509},
		{50: 508},
		// 225
		{140, 8: 140, 140, 140, 140, 21: 140, 140, 140, 25: 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 43: 140, 140, 140, 140, 140, 140, 52: 140},
		{137, 8: 137, 137, 137, 137, 21: 137, 137, 137, 25: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 43: 137, 137, 137, 137, 137, 137, 52: 137},
		{21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 106: 293, 111: 312, 314, 311, 292, 117: 475, 294, 291, 139: 511, 146: 512},
		{8: 124},
		{8: 120, 10: 120},
		// 230
		{105, 8: 105, 10: 105, 105},
		{8: 97, 21: 97, 97, 97, 26: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 43: 97, 97, 97, 97, 97, 183: 515},
		{8: 123, 21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 106: 293, 111: 312, 314, 311, 292, 117: 475, 294, 291, 146: 474, 155: 472, 473, 177: 516},
		{8: 517},
		{96, 8: 96, 10: 96, 96},
		// 235
		{133, 442, 133, 8: 133, 10: 133, 133, 127: 519},
		{132, 2: 132, 8: 132, 10: 132, 132},
		{165, 165, 165, 8: 165, 11: 165, 25: 165},
		{163, 163, 163, 8: 163, 11: 163, 25: 163},
		{166, 166, 166, 8: 166, 11: 166, 25: 166},
		// 240
		{341, 247, 334, 247, 247, 247, 344, 343, 247, 247, 247, 247, 350, 349, 335, 336, 337, 338, 339, 351, 340, 24: 247, 247, 42: 247, 48: 247, 524, 247, 247, 525, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247},
		{356, 246, 3: 246, 246, 246, 360, 359, 246, 246, 246, 355, 24: 246, 246, 42: 246, 48: 246, 50: 246, 357, 53: 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 358, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246},
		{88, 88, 88, 88, 88, 88, 88, 88, 11: 531, 88, 88, 88, 88, 88, 88, 88, 88, 88, 51: 532, 88, 145: 530, 149: 529, 527, 528, 176: 526},
		{10: 542, 24: 158, 148: 547},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 538, 52: 539, 154: 540},
		// 245
		{11: 531, 48: 536, 51: 532, 145: 537},
		{87, 87, 87, 87, 87, 87, 87, 87, 12: 87, 87, 87, 87, 87, 87, 87, 87, 87, 52: 87},
		{11: 86, 48: 86, 51: 86},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 342, 138: 534},
		{2: 533},
		// 250
		{11: 83, 48: 83, 51: 83},
		{50: 535},
		{11: 84, 48: 84, 51: 84},
		{89, 89, 89, 89, 89, 89, 89, 89, 12: 89, 89, 89, 89, 89, 89, 89, 89, 89, 52: 89},
		{11: 85, 48: 85, 51: 85},
		// 255
		{356, 361, 3: 374, 364, 365, 360, 359, 9: 93, 93, 355, 24: 93, 48: 380, 51: 357, 53: 363, 362, 368, 369, 379, 375, 376, 384, 377, 388, 358, 382, 372, 371, 370, 366, 386, 383, 381, 373, 390, 378, 367, 387, 385, 389},
		{88, 88, 88, 88, 88, 88, 88, 88, 11: 531, 88, 88, 88, 88, 88, 88, 88, 88, 88, 51: 532, 88, 145: 530, 149: 529, 527, 528, 176: 541},
		{10: 91, 24: 91},
		{10: 542, 24: 158, 148: 543},
		{88, 88, 88, 88, 88, 88, 88, 88, 11: 531, 88, 88, 88, 88, 88, 88, 88, 88, 88, 24: 157, 51: 532, 88, 145: 530, 149: 529, 545, 528},
		// 260
		{24: 544},
		{9: 92, 92, 24: 92},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 538, 52: 539, 154: 546},
		{10: 90, 24: 90},
		{24: 548},
		// 265
		{257, 257, 3: 257, 257, 257, 257, 257, 257, 257, 257, 257, 24: 257, 257, 42: 257, 48: 257, 50: 257, 257, 53: 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257},
		{264, 264, 3: 264, 264, 264, 264, 264, 264, 264, 264, 264, 24: 264, 264, 42: 264, 48: 264, 50: 264, 264, 53: 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264},
		{356, 249, 3: 249, 249, 249, 360, 359, 249, 249, 249, 355, 24: 249, 249, 42: 249, 48: 249, 50: 249, 357, 53: 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 358, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249},
		{356, 250, 3: 250, 250, 250, 360, 359, 250, 250, 250, 355, 24: 250, 250, 42: 250, 48: 250, 50: 250, 357, 53: 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 358, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250},
		{356, 251, 3: 251, 251, 251, 360, 359, 251, 251, 251, 355, 24: 251, 251, 42: 251, 48: 251, 50: 251, 357, 53: 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 358, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251},
		// 270
		{356, 252, 3: 252, 252, 252, 360, 359, 252, 252, 252, 355, 24: 252, 252, 42: 252, 48: 252, 50: 252, 357, 53: 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 358, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252},
		{356, 254, 3: 254, 254, 254, 360, 359, 254, 254, 254, 355, 24: 254, 254, 42: 254, 48: 254, 50: 254, 357, 53: 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 358, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254},
		{356, 255, 3: 255, 255, 255, 360, 359, 255, 255, 255, 355, 24: 255, 255, 42: 255, 48: 255, 50: 255, 357, 53: 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 358, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{356, 256, 3: 256, 256, 256, 360, 359, 256, 256, 256, 355, 24: 256, 256, 42: 256, 48: 256, 50: 256, 357, 53: 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 358, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256},
		{8: 558},
		// 275
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 524, 52: 525},
		{2: 329, 24: 157, 167: 332, 562},
		{24: 561},
		{155, 155, 155, 8: 155, 155, 155, 155, 21: 155, 155, 155, 25: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 43: 155, 155, 155, 155, 155},
		{10: 152, 24: 152},
		// 280
		{52: 175, 188: 565},
		{172, 172, 172, 8: 172, 172, 172, 172, 21: 172, 172, 172, 25: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 43: 172, 172, 172, 172, 172, 52: 112},
		{52: 566},
		{21: 174, 174, 174, 26: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 189: 567},
		{21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 106: 440, 111: 312, 314, 311, 439, 141: 570, 179: 569, 209: 568},
		// 285
		{21: 318, 319, 320, 583, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 106: 440, 111: 312, 314, 311, 439, 141: 570, 179: 584},
		{21: 169, 169, 169, 169, 26: 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169},
		{131, 442, 131, 25: 144, 127: 576, 575, 130: 573, 162: 574, 180: 572, 210: 571},
		{9: 580, 581},
		{9: 162, 162},
		// 290
		{9: 160, 160, 25: 143},
		{25: 578},
		{577, 2: 482, 165: 481},
		{130, 2: 130},
		{131, 442, 131, 127: 576, 575, 130: 484},
		// 295
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 342, 138: 579},
		{9: 159, 159},
		{21: 167, 167, 167, 167, 26: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167},
		{131, 442, 131, 25: 144, 127: 576, 575, 130: 573, 162: 574, 180: 582},
		{9: 161, 161},
		// 300
		{173, 173, 173, 8: 173, 173, 173, 173, 21: 173, 173, 173, 25: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 43: 173, 173, 173, 173, 173},
		{21: 168, 168, 168, 168, 26: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168},
		{204, 204, 204, 8: 204, 204, 204, 204},
		{202, 202, 202, 8: 202, 202, 202, 202},
		{205, 205, 205, 8: 205, 205, 205, 205},
		// 305
		{206, 206, 206, 8: 206, 206, 206, 206},
		{207, 207, 207, 8: 207, 207, 207, 207},
		{9: 684},
		{9: 201, 201},
		{9: 198, 682},
		// 310
		{9: 197, 197, 21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 196, 52: 45, 106: 293, 111: 312, 314, 311, 292, 117: 594, 294, 291, 135: 597, 158: 595, 197: 598, 596},
		{131, 442, 131, 9: 199, 127: 576, 575, 130: 681, 153: 591, 174: 592, 590},
		{48: 679},
		{52: 49, 185: 600},
		{21: 47, 47, 47, 26: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 43: 47, 47, 47, 47, 47, 52: 47},
		// 315
		{21: 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 52: 44, 106: 293, 111: 312, 314, 311, 292, 117: 594, 294, 291, 135: 599},
		{21: 46, 46, 46, 26: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 43: 46, 46, 46, 46, 46, 52: 46},
		{52: 601, 120: 602},
		{73, 73, 73, 73, 73, 73, 73, 73, 9: 73, 12: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 26: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 43: 73, 73, 73, 73, 73, 52: 73, 82: 73, 73, 73, 73, 73, 73, 73, 73, 73, 92: 73, 73, 184: 603},
		{21: 48, 48, 48, 26: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48},
		// 320
		{341, 346, 604, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 318, 319, 320, 69, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 49: 402, 52: 601, 82: 622, 627, 612, 626, 613, 623, 624, 625, 620, 92: 628, 621, 106: 293, 605, 111: 312, 314, 311, 292, 619, 117: 594, 294, 291, 607, 608, 610, 611, 606, 609, 618, 135: 617, 161: 615, 194: 616, 614},
		{271, 271, 3: 271, 271, 271, 271, 271, 9: 271, 271, 271, 25: 677, 48: 271, 51: 271, 53: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{8: 210, 210, 405},
		{82, 82, 82, 82, 82, 82, 82, 82, 9: 82, 12: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 26: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 43: 82, 82, 82, 82, 82, 52: 82, 82: 82, 82, 82, 82, 82, 82, 82, 82, 82, 92: 82, 82, 108: 82},
		{81, 81, 81, 81, 81, 81, 81, 81, 9: 81, 12: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 26: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 43: 81, 81, 81, 81, 81, 52: 81, 82: 81, 81, 81, 81, 81, 81, 81, 81, 81, 92: 81, 81, 108: 81},
		// 325
		{80, 80, 80, 80, 80, 80, 80, 80, 9: 80, 12: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 26: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 43: 80, 80, 80, 80, 80, 52: 80, 82: 80, 80, 80, 80, 80, 80, 80, 80, 80, 92: 80, 80, 108: 80},
		{79, 79, 79, 79, 79, 79, 79, 79, 9: 79, 12: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 26: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 43: 79, 79, 79, 79, 79, 52: 79, 82: 79, 79, 79, 79, 79, 79, 79, 79, 79, 92: 79, 79, 108: 79},
		{78, 78, 78, 78, 78, 78, 78, 78, 9: 78, 12: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 26: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 43: 78, 78, 78, 78, 78, 52: 78, 82: 78, 78, 78, 78, 78, 78, 78, 78, 78, 92: 78, 78, 108: 78},
		{77, 77, 77, 77, 77, 77, 77, 77, 9: 77, 12: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 26: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 43: 77, 77, 77, 77, 77, 52: 77, 82: 77, 77, 77, 77, 77, 77, 77, 77, 77, 92: 77, 77, 108: 77},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 342, 138: 674},
		// 330
		{25: 672},
		{24: 671},
		{71, 71, 71, 71, 71, 71, 71, 71, 9: 71, 12: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 26: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 43: 71, 71, 71, 71, 71, 52: 71, 82: 71, 71, 71, 71, 71, 71, 71, 71, 71, 92: 71, 71},
		{341, 346, 604, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 318, 319, 320, 68, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 49: 402, 52: 601, 82: 622, 627, 612, 626, 613, 623, 624, 625, 620, 92: 628, 621, 106: 293, 605, 111: 312, 314, 311, 292, 619, 117: 594, 294, 291, 607, 608, 610, 611, 606, 609, 618, 135: 617, 161: 670},
		{67, 67, 67, 67, 67, 67, 67, 67, 9: 67, 12: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 26: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 43: 67, 67, 67, 67, 67, 52: 67, 82: 67, 67, 67, 67, 67, 67, 67, 67, 67, 92: 67, 67},
		// 335
		{66, 66, 66, 66, 66, 66, 66, 66, 9: 66, 12: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 26: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 43: 66, 66, 66, 66, 66, 52: 66, 82: 66, 66, 66, 66, 66, 66, 66, 66, 66, 92: 66, 66},
		{9: 669},
		{663},
		{659},
		{655},
		// 340
		{341, 346, 604, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 52: 601, 82: 622, 627, 612, 626, 613, 623, 624, 625, 620, 92: 628, 621, 107: 605, 115: 619, 120: 607, 608, 610, 611, 606, 609, 649},
		{635},
		{2: 633},
		{9: 632},
		{9: 631},
		// 345
		{341, 346, 334, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 107: 605, 115: 629},
		{9: 630},
		{54, 54, 54, 54, 54, 54, 54, 54, 9: 54, 12: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 26: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 43: 54, 54, 54, 54, 54, 52: 54, 82: 54, 54, 54, 54, 54, 54, 54, 54, 54, 92: 54, 54, 108: 54},
		{55, 55, 55, 55, 55, 55, 55, 55, 9: 55, 12: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 26: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 43: 55, 55, 55, 55, 55, 52: 55, 82: 55, 55, 55, 55, 55, 55, 55, 55, 55, 92: 55, 55, 108: 55},
		{56, 56, 56, 56, 56, 56, 56, 56, 9: 56, 12: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 26: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 43: 56, 56, 56, 56, 56, 52: 56, 82: 56, 56, 56, 56, 56, 56, 56, 56, 56, 92: 56, 56, 108: 56},
		// 350
		{9: 634},
		{57, 57, 57, 57, 57, 57, 57, 57, 9: 57, 12: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 26: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 43: 57, 57, 57, 57, 57, 52: 57, 82: 57, 57, 57, 57, 57, 57, 57, 57, 57, 92: 57, 57, 108: 57},
		{341, 346, 334, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 318, 319, 320, 26: 309, 301, 310, 306, 317, 305, 303, 304, 302, 307, 315, 313, 316, 308, 300, 297, 43: 298, 296, 321, 299, 295, 49: 402, 106: 293, 605, 111: 312, 314, 311, 292, 636, 117: 594, 294, 291, 135: 637},
		{9: 643},
		{341, 346, 334, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 107: 605, 115: 638},
		// 355
		{9: 639},
		{341, 346, 334, 345, 347, 348, 344, 343, 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 107: 605, 115: 640},
		{8: 641},
		{341, 346, 604, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 52: 601, 82: 622, 627, 612, 626, 613, 623, 624, 625, 620, 92: 628, 621, 107: 605, 115: 619, 120: 607, 608, 610, 611, 606, 609, 642},
		{58, 58, 58, 58, 58, 58, 58, 58, 9: 58, 12: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 26: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 43: 58, 58, 58, 58, 58, 52: 58, 82: 58, 58, 58, 58, 58, 58, 58, 58, 58, 92: 58, 58, 108: 58},
		// 360
		{341, 346, 334, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 107: 605, 115: 644},
		{9: 645},
		{341, 346, 334, 345, 347, 348, 344, 343, 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 107: 605, 115: 646},
		{8: 647},
		{341, 346, 604, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 52: 601, 82: 622, 627, 612, 626, 613, 623, 624, 625, 620, 92: 628, 621, 107: 605, 115: 619, 120: 607, 608, 610, 611, 606, 609, 648},
		// 365
		{59, 59, 59, 59, 59, 59, 59, 59, 9: 59, 12: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 26: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 43: 59, 59, 59, 59, 59, 52: 59, 82: 59, 59, 59, 59, 59, 59, 59, 59, 59, 92: 59, 59, 108: 59},
		{82: 650},
		{651},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 107: 652},
		{8: 653, 10: 405},
		// 370
		{9: 654},
		{60, 60, 60, 60, 60, 60, 60, 60, 9: 60, 12: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 26: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 43: 60, 60, 60, 60, 60, 52: 60, 82: 60, 60, 60, 60, 60, 60, 60, 60, 60, 92: 60, 60, 108: 60},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 107: 656},
		{8: 657, 10: 405},
		{341, 346, 604, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 52: 601, 82: 622, 627, 612, 626, 613, 623, 624, 625, 620, 92: 628, 621, 107: 605, 115: 619, 120: 607, 608, 610, 611, 606, 609, 658},
		// 375
		{61, 61, 61, 61, 61, 61, 61, 61, 9: 61, 12: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 26: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 43: 61, 61, 61, 61, 61, 52: 61, 82: 61, 61, 61, 61, 61, 61, 61, 61, 61, 92: 61, 61, 108: 61},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 107: 660},
		{8: 661, 10: 405},
		{341, 346, 604, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 52: 601, 82: 622, 627, 612, 626, 613, 623, 624, 625, 620, 92: 628, 621, 107: 605, 115: 619, 120: 607, 608, 610, 611, 606, 609, 662},
		{62, 62, 62, 62, 62, 62, 62, 62, 9: 62, 12: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 26: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 43: 62, 62, 62, 62, 62, 52: 62, 82: 62, 62, 62, 62, 62, 62, 62, 62, 62, 92: 62, 62, 108: 62},
		// 380
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 107: 664},
		{8: 665, 10: 405},
		{341, 346, 604, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 52: 601, 82: 622, 627, 612, 626, 613, 623, 624, 625, 620, 92: 628, 621, 107: 605, 115: 619, 120: 607, 608, 610, 611, 606, 609, 666},
		{64, 64, 64, 64, 64, 64, 64, 64, 9: 64, 12: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 26: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 43: 64, 64, 64, 64, 64, 52: 64, 82: 64, 64, 64, 64, 64, 64, 64, 64, 64, 92: 64, 64, 108: 667},
		{341, 346, 604, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 52: 601, 82: 622, 627, 612, 626, 613, 623, 624, 625, 620, 92: 628, 621, 107: 605, 115: 619, 120: 607, 608, 610, 611, 606, 609, 668},
		// 385
		{63, 63, 63, 63, 63, 63, 63, 63, 9: 63, 12: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 26: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 43: 63, 63, 63, 63, 63, 52: 63, 82: 63, 63, 63, 63, 63, 63, 63, 63, 63, 92: 63, 63, 108: 63},
		{65, 65, 65, 65, 65, 65, 65, 65, 9: 65, 12: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 26: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 43: 65, 65, 65, 65, 65, 52: 65, 82: 65, 65, 65, 65, 65, 65, 65, 65, 65, 92: 65, 65, 108: 65},
		{70, 70, 70, 70, 70, 70, 70, 70, 9: 70, 12: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 26: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 43: 70, 70, 70, 70, 70, 52: 70, 82: 70, 70, 70, 70, 70, 70, 70, 70, 70, 92: 70, 70},
		{72, 72, 72, 72, 72, 72, 72, 72, 9: 72, 12: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 26: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 52: 72, 82: 72, 72, 72, 72, 72, 72, 72, 72, 72, 92: 72, 72, 108: 72},
		{341, 346, 604, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 52: 601, 82: 622, 627, 612, 626, 613, 623, 624, 625, 620, 92: 628, 621, 107: 605, 115: 619, 120: 607, 608, 610, 611, 606, 609, 673},
		// 390
		{74, 74, 74, 74, 74, 74, 74, 74, 9: 74, 12: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 26: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 43: 74, 74, 74, 74, 74, 52: 74, 82: 74, 74, 74, 74, 74, 74, 74, 74, 74, 92: 74, 74, 108: 74},
		{25: 675},
		{341, 346, 604, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 52: 601, 82: 622, 627, 612, 626, 613, 623, 624, 625, 620, 92: 628, 621, 107: 605, 115: 619, 120: 607, 608, 610, 611, 606, 609, 676},
		{75, 75, 75, 75, 75, 75, 75, 75, 9: 75, 12: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 26: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 43: 75, 75, 75, 75, 75, 52: 75, 82: 75, 75, 75, 75, 75, 75, 75, 75, 75, 92: 75, 75, 108: 75},
		{341, 346, 604, 345, 347, 348, 344, 343, 9: 211, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 402, 52: 601, 82: 622, 627, 612, 626, 613, 623, 624, 625, 620, 92: 628, 621, 107: 605, 115: 619, 120: 607, 608, 610, 611, 606, 609, 678},
		// 395
		{76, 76, 76, 76, 76, 76, 76, 76, 9: 76, 12: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 26: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 43: 76, 76, 76, 76, 76, 52: 76, 82: 76, 76, 76, 76, 76, 76, 76, 76, 76, 92: 76, 76, 108: 76},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 538, 52: 539, 154: 680},
		{9: 195, 195},
		{9: 197, 197, 48: 196, 158: 595},
		{131, 442, 131, 127: 576, 575, 130: 681, 153: 683},
		// 400
		{9: 200, 200},
		{208, 208, 208, 208, 208, 208, 208, 208, 9: 208, 12: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 26: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 52: 208, 82: 208, 208, 208, 208, 208, 208, 208, 208, 208, 92: 208, 208},
		{21: 52, 52, 52, 26: 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52},
		{341, 346, 334, 345, 347, 348, 344, 343, 12: 350, 349, 335, 336, 337, 338, 339, 351, 340, 49: 342, 138: 687},
		{42: 279},
		// 405
		{79: 709, 711, 95: 700, 701, 702, 697, 698, 699, 703, 704, 694, 705, 706, 109: 710, 708, 116: 707, 129: 692, 131: 691, 696, 693, 695, 136: 690, 207: 689},
		{42: 281},
		{42: 43, 79: 709, 711, 95: 700, 701, 702, 697, 698, 699, 703, 704, 694, 705, 706, 109: 710, 708, 116: 707, 129: 692, 131: 752, 696, 693, 695},
		{42: 42, 79: 42, 42, 42, 91: 42, 94: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{42: 38, 79: 38, 38, 38, 91: 38, 94: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		// 410
		{42: 37, 79: 37, 37, 37, 91: 37, 94: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{80: 711, 109: 774, 708},
		{42: 35, 79: 35, 35, 35, 91: 35, 94: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{81: 28, 91: 28, 94: 762, 166: 760, 199: 761, 759},
		{80: 711, 109: 756, 708},
		// 415
		{2: 753},
		{2: 748},
		{2: 724, 79: 726, 205: 725},
		{79: 709, 711, 109: 710, 708, 116: 723},
		{42: 16, 79: 16, 16, 16, 91: 16, 94: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		// 420
		{80: 711, 109: 721, 708},
		{80: 711, 109: 719, 708},
		{79: 709, 711, 109: 710, 708, 116: 718},
		{2: 714},
		{42: 7, 79: 7, 7, 7, 91: 7, 94: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		// 425
		{79: 5, 713},
		{42: 4, 79: 4, 4, 4, 91: 4, 94: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{79: 712},
		{79: 2, 2},
		{42: 3, 79: 3, 3, 3, 91: 3, 94: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		// 430
		{79: 1, 1},
		{79: 715, 711, 109: 716, 708},
		{42: 12, 79: 12, 12, 12, 91: 12, 94: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{79: 717},
		{42: 8, 79: 8, 8, 8, 91: 8, 94: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		// 435
		{42: 13, 79: 13, 13, 13, 91: 13, 94: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{79: 720},
		{42: 14, 79: 14, 14, 14, 91: 14, 94: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{79: 722},
		{42: 15, 79: 15, 15, 15, 91: 15, 94: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		// 440
		{42: 17, 79: 17, 17, 17, 91: 17, 94: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{79: 709, 711, 109: 710, 708, 116: 733, 137: 747},
		{2: 727, 8: 115, 139: 729, 171: 728, 730},
		{42: 9, 79: 9, 9, 9, 91: 9, 94: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{8: 117, 10: 117, 139: 744},
		// 445
		{8: 114, 10: 736},
		{8: 734},
		{8: 731},
		{79: 709, 711, 109: 710, 708, 116: 733, 137: 732},
		{42: 18, 79: 18, 18, 18, 91: 18, 94: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		// 450
		{42: 6, 79: 6, 6, 6, 91: 6, 94: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{79: 709, 711, 109: 710, 708, 116: 733, 137: 735},
		{42: 20, 79: 20, 20, 20, 91: 20, 94: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{2: 737, 139: 738},
		{8: 116, 10: 116, 139: 741},
		// 455
		{8: 739},
		{79: 709, 711, 109: 710, 708, 116: 733, 137: 740},
		{42: 19, 79: 19, 19, 19, 91: 19, 94: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{8: 742},
		{79: 709, 711, 109: 710, 708, 116: 733, 137: 743},
		// 460
		{42: 10, 79: 10, 10, 10, 91: 10, 94: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{8: 745},
		{79: 709, 711, 109: 710, 708, 116: 733, 137: 746},
		{42: 11, 79: 11, 11, 11, 91: 11, 94: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{42: 21, 79: 21, 21, 21, 91: 21, 94: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		// 465
		{79: 749},
		{79: 709, 711, 40, 91: 40, 94: 40, 700, 701, 702, 697, 698, 699, 703, 704, 694, 705, 706, 109: 710, 708, 116: 707, 129: 692, 131: 691, 696, 693, 695, 136: 750, 140: 751},
		{79: 709, 711, 39, 91: 39, 94: 39, 700, 701, 702, 697, 698, 699, 703, 704, 694, 705, 706, 109: 710, 708, 116: 707, 129: 692, 131: 752, 696, 693, 695},
		{81: 31, 91: 31, 94: 31},
		{42: 41, 79: 41, 41, 41, 91: 41, 94: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		// 470
		{79: 754},
		{79: 709, 711, 40, 91: 40, 94: 40, 700, 701, 702, 697, 698, 699, 703, 704, 694, 705, 706, 109: 710, 708, 116: 707, 129: 692, 131: 691, 696, 693, 695, 136: 750, 140: 755},
		{81: 32, 91: 32, 94: 32},
		{79: 757},
		{79: 709, 711, 40, 91: 40, 94: 40, 700, 701, 702, 697, 698, 699, 703, 704, 694, 705, 706, 109: 710, 708, 116: 707, 129: 692, 131: 691, 696, 693, 695, 136: 750, 140: 758},
		// 475
		{81: 33, 91: 33, 94: 33},
		{81: 24, 91: 768, 201: 769, 767},
		{81: 30, 91: 30, 94: 30},
		{81: 27, 91: 27, 94: 762, 166: 766},
		{80: 711, 109: 763, 708},
		// 480
		{79: 764},
		{79: 709, 711, 40, 91: 40, 94: 40, 700, 701, 702, 697, 698, 699, 703, 704, 694, 705, 706, 109: 710, 708, 116: 707, 129: 692, 131: 691, 696, 693, 695, 136: 750, 140: 765},
		{81: 26, 91: 26, 94: 26},
		{81: 29, 91: 29, 94: 29},
		{81: 773, 203: 772},
		// 485
		{79: 770},
		{81: 23},
		{79: 709, 711, 40, 95: 700, 701, 702, 697, 698, 699, 703, 704, 694, 705, 706, 109: 710, 708, 116: 707, 129: 692, 131: 691, 696, 693, 695, 136: 750, 140: 771},
		{81: 25},
		{42: 34, 79: 34, 34, 34, 91: 34, 94: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		// 490
		{42: 22, 79: 22, 22, 22, 91: 22, 94: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{79: 775},
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
	const yyError = 215

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
			lhs := &ConstantExpression{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Value, lhs.Type = lhs.Expression.eval(lx)
			if lhs.Value == nil {
				lx.report.Err(lhs.Pos(), "not a constant expression")
			}
		}
	case 75:
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
	case 76:
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
	case 77:
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
	case 78:
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
	case 79:
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
	case 80:
		{
			yyVAL.node = (*DeclarationSpecifiersOpt)(nil)
		}
	case 81:
		{
			lhs := &DeclarationSpecifiersOpt{
				DeclarationSpecifiers: yyS[yypt-0].node.(*DeclarationSpecifiers),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.DeclarationSpecifiers.attr
			lhs.typeSpecifier = lhs.DeclarationSpecifiers.typeSpecifier
		}
	case 82:
		{
			yyVAL.node = &InitDeclaratorList{
				InitDeclarator: yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 83:
		{
			yyVAL.node = &InitDeclaratorList{
				Case:               1,
				InitDeclaratorList: yyS[yypt-2].node.(*InitDeclaratorList),
				Token:              yyS[yypt-1].Token,
				InitDeclarator:     yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 84:
		{
			yyVAL.node = (*InitDeclaratorListOpt)(nil)
		}
	case 85:
		{
			yyVAL.node = &InitDeclaratorListOpt{
				InitDeclaratorList: yyS[yypt-0].node.(*InitDeclaratorList).reverse(),
			}
		}
	case 86:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
		}
	case 87:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-0].node.(*Declarator)
			d.setFull(lx)
		}
	case 88:
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
	case 89:
		{
			lhs := &StorageClassSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saTypedef
		}
	case 90:
		{
			lhs := &StorageClassSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saExtern
		}
	case 91:
		{
			lhs := &StorageClassSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saStatic
		}
	case 92:
		{
			lhs := &StorageClassSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saAuto
		}
	case 93:
		{
			lhs := &StorageClassSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRegister
		}
	case 94:
		{
			lhs := &TypeSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsVoid)
		}
	case 95:
		{
			lhs := &TypeSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsChar)
		}
	case 96:
		{
			lhs := &TypeSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsShort)
		}
	case 97:
		{
			lhs := &TypeSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsInt)
		}
	case 98:
		{
			lhs := &TypeSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsLong)
		}
	case 99:
		{
			lhs := &TypeSpecifier{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsFloat)
		}
	case 100:
		{
			lhs := &TypeSpecifier{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsDouble)
		}
	case 101:
		{
			lhs := &TypeSpecifier{
				Case:  7,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsSigned)
		}
	case 102:
		{
			lhs := &TypeSpecifier{
				Case:  8,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsUnsigned)
		}
	case 103:
		{
			lhs := &TypeSpecifier{
				Case:  9,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsBool)
		}
	case 104:
		{
			lhs := &TypeSpecifier{
				Case:  10,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsComplex)
		}
	case 105:
		{
			lhs := &TypeSpecifier{
				Case: 11,
				StructOrUnionSpecifier: yyS[yypt-0].node.(*StructOrUnionSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(lhs.StructOrUnionSpecifier.typeSpecifiers())
		}
	case 106:
		{
			lhs := &TypeSpecifier{
				Case:          12,
				EnumSpecifier: yyS[yypt-0].node.(*EnumSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsEnumSpecifier)
		}
	case 107:
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
	case 108:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareStructTag(o.Token, lx.report)
			}
		}
	case 109:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeMembers)
			lx.scope.isUnion = yyS[yypt-3].node.(*StructOrUnion).Case == 1 // "union"
			lx.scope.prevStructDeclarator = nil
		}
	case 110:
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
	case 111:
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
	case 112:
		{
			yyVAL.node = &StructOrUnion{
				Token: yyS[yypt-0].Token,
			}
		}
	case 113:
		{
			yyVAL.node = &StructOrUnion{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 114:
		{
			yyVAL.node = &StructDeclarationList{
				StructDeclaration: yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 115:
		{
			yyVAL.node = &StructDeclarationList{
				Case: 1,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList),
				StructDeclaration:     yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 116:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclaration{
				SpecifierQualifierList: yyS[yypt-2].node.(*SpecifierQualifierList),
				StructDeclaratorList:   yyS[yypt-1].node.(*StructDeclaratorList).reverse(),
				Token:                  yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			ts := tsDecode(lhs.SpecifierQualifierList.typeSpecifier)
			ok := false
			for _, v := range ts {
				if v == tsStructSpecifier || v == tsUnionSpecifier {
					ok = true
					break
				}
			}
			if !ok {
				break
			}

			s := lhs.SpecifierQualifierList
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
	case 117:
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
	case 118:
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
	case 119:
		{
			yyVAL.node = (*SpecifierQualifierListOpt)(nil)
		}
	case 120:
		{
			lhs := &SpecifierQualifierListOpt{
				SpecifierQualifierList: yyS[yypt-0].node.(*SpecifierQualifierList),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.SpecifierQualifierList.attr
			lhs.typeSpecifier = lhs.SpecifierQualifierList.typeSpecifier
		}
	case 121:
		{
			yyVAL.node = &StructDeclaratorList{
				StructDeclarator: yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 122:
		{
			yyVAL.node = &StructDeclaratorList{
				Case:                 1,
				StructDeclaratorList: yyS[yypt-2].node.(*StructDeclaratorList),
				Token:                yyS[yypt-1].Token,
				StructDeclarator:     yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 123:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.post(lx)
		}
	case 124:
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
	case 125:
		{
			yyVAL.node = (*CommaOpt)(nil)
		}
	case 126:
		{
			yyVAL.node = &CommaOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 127:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareEnumTag(o.Token, lx.report)
			}
			lx.iota = 0
		}
	case 128:
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
	case 129:
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
	case 130:
		{
			yyVAL.node = &EnumeratorList{
				Enumerator: yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 131:
		{
			yyVAL.node = &EnumeratorList{
				Case:           1,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList),
				Token:          yyS[yypt-1].Token,
				Enumerator:     yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 132:
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
	case 133:
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
	case 134:
		{
			lhs := &TypeQualifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saConst
		}
	case 135:
		{
			lhs := &TypeQualifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRestrict
		}
	case 136:
		{
			lhs := &TypeQualifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saVolatile
		}
	case 137:
		{
			lhs := &FunctionSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saInline
		}
	case 138:
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
	case 139:
		{
			yyVAL.node = (*DeclaratorOpt)(nil)
		}
	case 140:
		{
			yyVAL.node = &DeclaratorOpt{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
		}
	case 141:
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
	case 142:
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
	case 143:
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
	case 144:
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
	case 145:
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
	case 146:
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
	case 147:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 148:
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
	case 149:
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
	case 150:
		{
			yyVAL.node = &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].node.(*TypeQualifierListOpt),
			}
		}
	case 151:
		{
			yyVAL.node = &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].node.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].node.(*Pointer),
			}
		}
	case 152:
		{
			yyVAL.node = (*PointerOpt)(nil)
		}
	case 153:
		{
			yyVAL.node = &PointerOpt{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
		}
	case 154:
		{
			lhs := &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.TypeQualifier.attr
		}
	case 155:
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
	case 156:
		{
			yyVAL.node = (*TypeQualifierListOpt)(nil)
		}
	case 157:
		{
			yyVAL.node = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].node.(*TypeQualifierList).reverse(),
			}
		}
	case 158:
		{
			lhs := &ParameterTypeList{
				ParameterList: yyS[yypt-0].node.(*ParameterList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 159:
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
	case 160:
		{
			yyVAL.node = (*ParameterTypeListOpt)(nil)
		}
	case 161:
		{
			yyVAL.node = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].node.(*ParameterTypeList),
			}
		}
	case 162:
		{
			yyVAL.node = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 163:
		{
			yyVAL.node = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].node.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 164:
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
	case 165:
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
	case 166:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 167:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 168:
		{
			yyVAL.node = (*IdentifierListOpt)(nil)
		}
	case 169:
		{
			yyVAL.node = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 170:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 171:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 172:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeBlock)
		}
	case 173:
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
	case 174:
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
	case 175:
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
	case 176:
		{
			yyVAL.node = (*AbstractDeclaratorOpt)(nil)
		}
	case 177:
		{
			yyVAL.node = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].node.(*AbstractDeclarator),
			}
		}
	case 178:
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
	case 179:
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
	case 180:
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
	case 181:
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
	case 182:
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
	case 183:
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
	case 184:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 185:
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
	case 186:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 187:
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
	case 188:
		{
			yyVAL.node = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 189:
		{
			yyVAL.node = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
		}
	case 190:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 191:
		{
			yyVAL.node = &Initializer{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 192:
		{
			yyVAL.node = &InitializerList{
				DesignationOpt: yyS[yypt-1].node.(*DesignationOpt),
				Initializer:    yyS[yypt-0].node.(*Initializer),
			}
		}
	case 193:
		{
			yyVAL.node = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].node.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].node.(*DesignationOpt),
				Initializer:     yyS[yypt-0].node.(*Initializer),
			}
		}
	case 194:
		{
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 195:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 196:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 197:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 198:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 199:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 200:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 201:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 202:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 203:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 204:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 205:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 206:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 207:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 208:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 209:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 210:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 211:
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
	case 212:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 213:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 214:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 215:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 216:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 217:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 218:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 219:
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
	case 220:
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
	case 221:
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
	case 222:
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
	case 223:
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
	case 224:
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
	case 225:
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
	case 226:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 227:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 228:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 229:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 230:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 231:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 232:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 233:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 234:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-1].node.(*Declarator)
			d.setFull(lx)
			if k := d.Type.Kind(); k != Function {
				lx.report.Err(d.Pos(), "declarator is not a function (have '%s': %v)", d.Type, k)
			}
			lx.scope.mergeScope = d.DirectDeclarator.paramsScope
		}
	case 235:
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
	case 236:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 237:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 238:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 239:
		{
			yyVAL.node = &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
		}
	case 240:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 241:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 242:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 243:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 244:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 245:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(node)
		}
	case 246:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(node)
		}
	case 247:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 248:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 249:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 250:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 251:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 252:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 253:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 254:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 255:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 256:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 257:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 258:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 259:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 260:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 261:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 262:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 263:
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
	case 264:
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
	case 265:
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
	case 266:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 267:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 268:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 269:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 270:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 271:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 272:
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
	case 273:
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
	case 274:
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
	case 275:
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
	case 278:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 279:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
