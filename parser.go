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
	yyDefault           = 57450
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
	IDENTIFIER_NONREPL  = 57377
	IF                  = 57378
	INC                 = 57379
	INLINE              = 57380
	INT                 = 57381
	INTCONST            = 57382
	LEQ                 = 57383
	LONG                = 57384
	LONGCHARCONST       = 57385
	LONGSTRINGLITERAL   = 57386
	LSH                 = 57387
	LSHASSIGN           = 57388
	MODASSIGN           = 57389
	MULASSIGN           = 57390
	NEQ                 = 57391
	NOELSE              = 57392
	ORASSIGN            = 57393
	OROR                = 57394
	PPDEFINE            = 57395
	PPELIF              = 57396
	PPELSE              = 57397
	PPENDIF             = 57398
	PPERROR             = 57399
	PPHASH_NL           = 57400
	PPHEADER_NAME       = 57401
	PPIF                = 57402
	PPIFDEF             = 57403
	PPIFNDEF            = 57404
	PPINCLUDE           = 57405
	PPINCLUDE_NEXT      = 57406
	PPLINE              = 57407
	PPNONDIRECTIVE      = 57408
	PPNUMBER            = 57409
	PPOTHER             = 57410
	PPPASTE             = 57411
	PPPRAGMA            = 57412
	PPUNDEF             = 57413
	PREPROCESSING_FILE  = 1048576
	REGISTER            = 57414
	RESTRICT            = 57415
	RETURN              = 57416
	RSH                 = 57417
	RSHASSIGN           = 57418
	SHORT               = 57419
	SIGNED              = 57420
	SIZEOF              = 57421
	STATIC              = 57422
	STRINGLITERAL       = 57423
	STRUCT              = 57424
	SUBASSIGN           = 57425
	SWITCH              = 57426
	TRANSLATION_UNIT    = 1048578
	TYPEDEF             = 57427
	TYPEDEFNAME         = 57428
	UNARY               = 57429
	UNION               = 57430
	UNSIGNED            = 57431
	VOID                = 57432
	VOLATILE            = 57433
	WHILE               = 57434
	XORASSIGN           = 57435
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -286
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
		57379:   7,   // INC (199x)
		59:      8,   // ';' (181x)
		41:      9,   // ')' (176x)
		44:      10,  // ',' (168x)
		91:      11,  // '[' (149x)
		33:      12,  // '!' (132x)
		126:     13,  // '~' (132x)
		57356:   14,  // CHARCONST (132x)
		57371:   15,  // FLOATCONST (132x)
		57382:   16,  // INTCONST (132x)
		57385:   17,  // LONGCHARCONST (132x)
		57386:   18,  // LONGSTRINGLITERAL (132x)
		57421:   19,  // SIZEOF (132x)
		57423:   20,  // STRINGLITERAL (132x)
		57358:   21,  // CONST (113x)
		57415:   22,  // RESTRICT (113x)
		57433:   23,  // VOLATILE (113x)
		125:     24,  // '}' (110x)
		57351:   25,  // BOOL (103x)
		57355:   26,  // CHAR (103x)
		57357:   27,  // COMPLEX (103x)
		57365:   28,  // DOUBLE (103x)
		57367:   29,  // ENUM (103x)
		57370:   30,  // FLOAT (103x)
		57381:   31,  // INT (103x)
		57384:   32,  // LONG (103x)
		57419:   33,  // SHORT (103x)
		57420:   34,  // SIGNED (103x)
		57424:   35,  // STRUCT (103x)
		57428:   36,  // TYPEDEFNAME (103x)
		57430:   37,  // UNION (103x)
		57431:   38,  // UNSIGNED (103x)
		57432:   39,  // VOID (103x)
		58:      40,  // ':' (102x)
		57422:   41,  // STATIC (97x)
		57344:   42,  // $end (96x)
		57350:   43,  // AUTO (91x)
		57369:   44,  // EXTERN (91x)
		57380:   45,  // INLINE (91x)
		57414:   46,  // REGISTER (91x)
		57427:   47,  // TYPEDEF (91x)
		61:      48,  // '=' (87x)
		57486:   49,  // Expression (83x)
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
		57383:   67,  // LEQ (68x)
		57387:   68,  // LSH (68x)
		57388:   69,  // LSHASSIGN (68x)
		57389:   70,  // MODASSIGN (68x)
		57390:   71,  // MULASSIGN (68x)
		57391:   72,  // NEQ (68x)
		57393:   73,  // ORASSIGN (68x)
		57394:   74,  // OROR (68x)
		57417:   75,  // RSH (68x)
		57418:   76,  // RSHASSIGN (68x)
		57425:   77,  // SUBASSIGN (68x)
		57435:   78,  // XORASSIGN (68x)
		10:      79,  // '\n' (60x)
		57410:   80,  // PPOTHER (54x)
		57398:   81,  // PPENDIF (45x)
		57397:   82,  // PPELSE (41x)
		57434:   83,  // WHILE (41x)
		57352:   84,  // BREAK (40x)
		57353:   85,  // CASE (40x)
		57359:   86,  // CONTINUE (40x)
		57362:   87,  // DEFAULT (40x)
		57364:   88,  // DO (40x)
		57372:   89,  // FOR (40x)
		57374:   90,  // GOTO (40x)
		57378:   91,  // IF (40x)
		57396:   92,  // PPELIF (40x)
		57416:   93,  // RETURN (40x)
		57426:   94,  // SWITCH (40x)
		57395:   95,  // PPDEFINE (36x)
		57399:   96,  // PPERROR (36x)
		57400:   97,  // PPHASH_NL (36x)
		57402:   98,  // PPIF (36x)
		57403:   99,  // PPIFDEF (36x)
		57404:   100, // PPIFNDEF (36x)
		57405:   101, // PPINCLUDE (36x)
		57406:   102, // PPINCLUDE_NEXT (36x)
		57407:   103, // PPLINE (36x)
		57408:   104, // PPNONDIRECTIVE (36x)
		57412:   105, // PPPRAGMA (36x)
		57413:   106, // PPUNDEF (36x)
		57536:   107, // TypeQualifier (28x)
		57487:   108, // ExpressionList (26x)
		57510:   109, // PPTokenList (23x)
		57512:   110, // PPTokens (23x)
		57366:   111, // ELSE (22x)
		57482:   112, // EnumSpecifier (20x)
		57531:   113, // StructOrUnion (20x)
		57532:   114, // StructOrUnionSpecifier (20x)
		57539:   115, // TypeSpecifier (20x)
		57488:   116, // ExpressionListOpt (18x)
		57511:   117, // PPTokenListOpt (16x)
		57465:   118, // DeclarationSpecifiers (15x)
		57493:   119, // FunctionSpecifier (15x)
		57526:   120, // StorageClassSpecifier (15x)
		57459:   121, // CompoundStatement (13x)
		57490:   122, // ExpressionStatement (12x)
		57507:   123, // IterationStatement (12x)
		57508:   124, // JumpStatement (12x)
		57509:   125, // LabeledStatement (12x)
		57521:   126, // SelectionStatement (12x)
		57525:   127, // Statement (12x)
		57517:   128, // Pointer (11x)
		57518:   129, // PointerOpt (10x)
		57461:   130, // ControlLine (8x)
		57467:   131, // Declarator (8x)
		57496:   132, // GroupPart (8x)
		57500:   133, // IfGroup (8x)
		57501:   134, // IfSection (8x)
		57533:   135, // TextLine (8x)
		57462:   136, // Declaration (7x)
		57494:   137, // GroupList (6x)
		57520:   138, // ReplacementList (6x)
		57444:   139, // $@4 (5x)
		57460:   140, // ConstantExpression (5x)
		57360:   141, // DDD (5x)
		57495:   142, // GroupListOpt (5x)
		57522:   143, // SpecifierQualifierList (5x)
		57537:   144, // TypeQualifierList (5x)
		57451:   145, // AbstractDeclarator (4x)
		57466:   146, // DeclarationSpecifiersOpt (4x)
		57471:   147, // Designator (4x)
		57513:   148, // ParameterDeclaration (4x)
		57538:   149, // TypeQualifierListOpt (4x)
		57458:   150, // CommaOpt (3x)
		57469:   151, // Designation (3x)
		57470:   152, // DesignationOpt (3x)
		57472:   153, // DesignatorList (3x)
		57489:   154, // ExpressionOpt (3x)
		57502:   155, // InitDeclarator (3x)
		57505:   156, // Initializer (3x)
		57514:   157, // ParameterList (3x)
		57515:   158, // ParameterTypeList (3x)
		57437:   159, // $@10 (2x)
		57438:   160, // $@11 (2x)
		57445:   161, // $@5 (2x)
		57452:   162, // AbstractDeclaratorOpt (2x)
		57455:   163, // BlockItem (2x)
		57468:   164, // DeclaratorOpt (2x)
		57473:   165, // DirectAbstractDeclarator (2x)
		57474:   166, // DirectAbstractDeclaratorOpt (2x)
		57475:   167, // DirectDeclarator (2x)
		57476:   168, // ElifGroup (2x)
		57483:   169, // EnumerationConstant (2x)
		57484:   170, // Enumerator (2x)
		57491:   171, // ExternalDeclaration (2x)
		57492:   172, // FunctionDefinition (2x)
		57497:   173, // IdentifierList (2x)
		57498:   174, // IdentifierListOpt (2x)
		57499:   175, // IdentifierOpt (2x)
		57503:   176, // InitDeclaratorList (2x)
		57504:   177, // InitDeclaratorListOpt (2x)
		57506:   178, // InitializerList (2x)
		57516:   179, // ParameterTypeListOpt (2x)
		57523:   180, // SpecifierQualifierListOpt (2x)
		57527:   181, // StructDeclaration (2x)
		57529:   182, // StructDeclarator (2x)
		57535:   183, // TypeName (2x)
		57436:   184, // $@1 (1x)
		57439:   185, // $@12 (1x)
		57440:   186, // $@13 (1x)
		57441:   187, // $@14 (1x)
		57442:   188, // $@2 (1x)
		57443:   189, // $@3 (1x)
		57446:   190, // $@6 (1x)
		57447:   191, // $@7 (1x)
		57448:   192, // $@8 (1x)
		57449:   193, // $@9 (1x)
		57453:   194, // ArgumentExpressionList (1x)
		57454:   195, // ArgumentExpressionListOpt (1x)
		57456:   196, // BlockItemList (1x)
		57457:   197, // BlockItemListOpt (1x)
		1048577: 198, // CONSTANT_EXPRESSION (1x)
		57463:   199, // DeclarationList (1x)
		57464:   200, // DeclarationListOpt (1x)
		57477:   201, // ElifGroupList (1x)
		57478:   202, // ElifGroupListOpt (1x)
		57479:   203, // ElseGroup (1x)
		57480:   204, // ElseGroupOpt (1x)
		57481:   205, // EndifLine (1x)
		57485:   206, // EnumeratorList (1x)
		57376:   207, // IDENTIFIER_LPAREN (1x)
		1048576: 208, // PREPROCESSING_FILE (1x)
		57519:   209, // PreprocessingFile (1x)
		57524:   210, // Start (1x)
		57528:   211, // StructDeclarationList (1x)
		57530:   212, // StructDeclaratorList (1x)
		1048578: 213, // TRANSLATION_UNIT (1x)
		57534:   214, // TranslationUnit (1x)
		57450:   215, // $default (0x)
		57354:   216, // CAST (0x)
		57345:   217, // error (0x)
		57377:   218, // IDENTIFIER_NONREPL (0x)
		57392:   219, // NOELSE (0x)
		57401:   220, // PPHEADER_NAME (0x)
		57409:   221, // PPNUMBER (0x)
		57411:   222, // PPPASTE (0x)
		57429:   223, // UNARY (0x)
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
		"PPELSE",
		"WHILE",
		"BREAK",
		"CASE",
		"CONTINUE",
		"DEFAULT",
		"DO",
		"FOR",
		"GOTO",
		"IF",
		"PPELIF",
		"RETURN",
		"SWITCH",
		"PPDEFINE",
		"PPERROR",
		"PPHASH_NL",
		"PPIF",
		"PPIFDEF",
		"PPIFNDEF",
		"PPINCLUDE",
		"PPINCLUDE_NEXT",
		"PPLINE",
		"PPNONDIRECTIVE",
		"PPPRAGMA",
		"PPUNDEF",
		"TypeQualifier",
		"ExpressionList",
		"PPTokenList",
		"PPTokens",
		"ELSE",
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
		"IDENTIFIER_NONREPL",
		"NOELSE",
		"PPHEADER_NAME",
		"PPNUMBER",
		"PPPASTE",
		"UNARY",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {184, 0},
		2:   {210, 3},
		3:   {188, 0},
		4:   {210, 3},
		5:   {189, 0},
		6:   {210, 3},
		7:   {169, 1},
		8:   {194, 1},
		9:   {194, 3},
		10:  {195, 0},
		11:  {195, 1},
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
		68:  {154, 0},
		69:  {154, 1},
		70:  {108, 1},
		71:  {108, 3},
		72:  {116, 0},
		73:  {116, 1},
		74:  {139, 0},
		75:  {140, 2},
		76:  {136, 3},
		77:  {118, 2},
		78:  {118, 2},
		79:  {118, 2},
		80:  {118, 2},
		81:  {146, 0},
		82:  {146, 1},
		83:  {176, 1},
		84:  {176, 3},
		85:  {177, 0},
		86:  {177, 1},
		87:  {155, 1},
		88:  {161, 0},
		89:  {155, 4},
		90:  {120, 1},
		91:  {120, 1},
		92:  {120, 1},
		93:  {120, 1},
		94:  {120, 1},
		95:  {115, 1},
		96:  {115, 1},
		97:  {115, 1},
		98:  {115, 1},
		99:  {115, 1},
		100: {115, 1},
		101: {115, 1},
		102: {115, 1},
		103: {115, 1},
		104: {115, 1},
		105: {115, 1},
		106: {115, 1},
		107: {115, 1},
		108: {115, 1},
		109: {190, 0},
		110: {191, 0},
		111: {114, 7},
		112: {114, 2},
		113: {113, 1},
		114: {113, 1},
		115: {211, 1},
		116: {211, 2},
		117: {181, 3},
		118: {181, 2},
		119: {143, 2},
		120: {143, 2},
		121: {180, 0},
		122: {180, 1},
		123: {212, 1},
		124: {212, 3},
		125: {182, 1},
		126: {182, 3},
		127: {150, 0},
		128: {150, 1},
		129: {192, 0},
		130: {112, 7},
		131: {112, 2},
		132: {206, 1},
		133: {206, 3},
		134: {170, 1},
		135: {170, 3},
		136: {107, 1},
		137: {107, 1},
		138: {107, 1},
		139: {119, 1},
		140: {131, 2},
		141: {164, 0},
		142: {164, 1},
		143: {167, 1},
		144: {167, 3},
		145: {167, 5},
		146: {167, 6},
		147: {167, 6},
		148: {167, 5},
		149: {193, 0},
		150: {167, 5},
		151: {167, 4},
		152: {128, 2},
		153: {128, 3},
		154: {129, 0},
		155: {129, 1},
		156: {144, 1},
		157: {144, 2},
		158: {149, 0},
		159: {149, 1},
		160: {158, 1},
		161: {158, 3},
		162: {179, 0},
		163: {179, 1},
		164: {157, 1},
		165: {157, 3},
		166: {148, 2},
		167: {148, 2},
		168: {173, 1},
		169: {173, 3},
		170: {174, 0},
		171: {174, 1},
		172: {175, 0},
		173: {175, 1},
		174: {159, 0},
		175: {183, 3},
		176: {145, 1},
		177: {145, 2},
		178: {162, 0},
		179: {162, 1},
		180: {165, 3},
		181: {165, 4},
		182: {165, 5},
		183: {165, 6},
		184: {165, 6},
		185: {165, 4},
		186: {160, 0},
		187: {165, 4},
		188: {185, 0},
		189: {165, 5},
		190: {166, 0},
		191: {166, 1},
		192: {156, 1},
		193: {156, 4},
		194: {178, 2},
		195: {178, 4},
		196: {151, 2},
		197: {152, 0},
		198: {152, 1},
		199: {153, 1},
		200: {153, 2},
		201: {147, 3},
		202: {147, 2},
		203: {127, 1},
		204: {127, 1},
		205: {127, 1},
		206: {127, 1},
		207: {127, 1},
		208: {127, 1},
		209: {125, 3},
		210: {125, 4},
		211: {125, 3},
		212: {186, 0},
		213: {121, 4},
		214: {196, 1},
		215: {196, 2},
		216: {197, 0},
		217: {197, 1},
		218: {163, 1},
		219: {163, 1},
		220: {122, 2},
		221: {126, 5},
		222: {126, 7},
		223: {126, 5},
		224: {123, 5},
		225: {123, 7},
		226: {123, 9},
		227: {123, 8},
		228: {124, 3},
		229: {124, 2},
		230: {124, 2},
		231: {124, 3},
		232: {214, 1},
		233: {214, 2},
		234: {171, 1},
		235: {171, 1},
		236: {187, 0},
		237: {172, 5},
		238: {199, 1},
		239: {199, 2},
		240: {200, 0},
		241: {200, 1},
		242: {209, 1},
		243: {137, 1},
		244: {137, 2},
		245: {142, 0},
		246: {142, 1},
		247: {132, 1},
		248: {132, 1},
		249: {132, 3},
		250: {132, 1},
		251: {134, 4},
		252: {133, 4},
		253: {133, 4},
		254: {133, 4},
		255: {201, 1},
		256: {201, 2},
		257: {202, 0},
		258: {202, 1},
		259: {168, 4},
		260: {203, 3},
		261: {204, 0},
		262: {204, 1},
		263: {205, 1},
		264: {130, 3},
		265: {130, 5},
		266: {130, 7},
		267: {130, 5},
		268: {130, 2},
		269: {130, 1},
		270: {130, 3},
		271: {130, 3},
		272: {130, 2},
		273: {130, 3},
		274: {130, 6},
		275: {130, 8},
		276: {130, 2},
		277: {130, 4},
		278: {130, 3},
		279: {135, 1},
		280: {138, 1},
		281: {109, 1},
		282: {117, 1},
		283: {117, 2},
		284: {110, 1},
		285: {110, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 42}:   "invalid empty input",
		yyXError{491, -1}: "expected #endif",
		yyXError{493, -1}: "expected #endif",
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
		yyXError{451, -1}: "expected ')'",
		yyXError{452, -1}: "expected ')'",
		yyXError{460, -1}: "expected ')'",
		yyXError{463, -1}: "expected ')'",
		yyXError{466, -1}: "expected ')'",
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
		yyXError{430, -1}: "expected '\\n'",
		yyXError{434, -1}: "expected '\\n'",
		yyXError{438, -1}: "expected '\\n'",
		yyXError{441, -1}: "expected '\\n'",
		yyXError{443, -1}: "expected '\\n'",
		yyXError{470, -1}: "expected '\\n'",
		yyXError{475, -1}: "expected '\\n'",
		yyXError{478, -1}: "expected '\\n'",
		yyXError{485, -1}: "expected '\\n'",
		yyXError{490, -1}: "expected '\\n'",
		yyXError{496, -1}: "expected '\\n'",
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
		yyXError{483, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{489, -1}: "expected endif line or #endif",
		yyXError{415, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{481, -1}: "expected endif line or optional else group or one of [#else, #endif]",
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
		yyXError{472, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{409, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{75, -1}:  "expected identifier",
		yyXError{76, -1}:  "expected identifier",
		yyXError{210, -1}: "expected identifier",
		yyXError{250, -1}: "expected identifier",
		yyXError{344, -1}: "expected identifier",
		yyXError{417, -1}: "expected identifier",
		yyXError{418, -1}: "expected identifier",
		yyXError{425, -1}: "expected identifier",
		yyXError{447, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
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
		yyXError{410, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{411, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{412, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{414, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{421, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{427, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{429, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{432, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{435, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{437, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{439, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{440, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{442, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{444, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{445, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{448, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{454, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{455, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{457, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{462, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{465, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{468, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{469, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{474, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{494, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{495, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{497, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{473, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{477, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{480, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{482, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{487, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{488, -1}: "expected one of [#elif, #else, #endif]",
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
		yyXError{449, -1}: "expected one of [')', ',', ...]",
		yyXError{459, -1}: "expected one of [')', ',', ...]",
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
		yyXError{450, -1}: "expected one of [')', ',']",
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
		yyXError{428, -1}: "expected one of ['\\n', ppother]",
		yyXError{431, -1}: "expected one of ['\\n', ppother]",
		yyXError{433, -1}: "expected one of ['\\n', ppother]",
		yyXError{316, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{318, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{32, -1}:  "expected one of ['{', identifier]",
		yyXError{33, -1}:  "expected one of ['{', identifier]",
		yyXError{287, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{290, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{299, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{303, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{458, -1}: "expected one of [..., identifier]",
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
		yyXError{471, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{476, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{479, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{486, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{492, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
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
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{407, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{446, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{453, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{456, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{461, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{464, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{467, -1}: "expected replacement list or one of ['\\n', ppother]",
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
		yyXError{436, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{413, -1}: "expected token list or ppother",
		yyXError{416, -1}: "expected token list or ppother",
		yyXError{422, -1}: "expected token list or ppother",
		yyXError{423, -1}: "expected token list or ppother",
		yyXError{426, -1}: "expected token list or ppother",
		yyXError{484, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{5, -1}:   "expected translation unit or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{175, -1}: "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{217, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{368, -1}: "expected while",
		yyXError{3, 42}:   "unexpected EOF",
		yyXError{2, 42}:   "unexpected EOF",
		yyXError{4, 42}:   "unexpected EOF",
	}

	yyParseTab = [498][]uint16{
		// 0
		{198: 289, 208: 288, 210: 287, 213: 290},
		{42: 286},
		{79: 285, 285, 95: 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 184: 693},
		{283, 283, 283, 283, 283, 283, 283, 283, 12: 283, 283, 283, 283, 283, 283, 283, 283, 283, 188: 691},
		{21: 281, 281, 281, 25: 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 41: 281, 43: 281, 281, 281, 281, 281, 189: 291},
		// 5
		{21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 107: 296, 112: 315, 317, 314, 295, 118: 293, 297, 294, 136: 327, 171: 325, 326, 214: 292},
		{21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 280, 301, 299, 324, 302, 298, 107: 296, 112: 315, 317, 314, 295, 118: 293, 297, 294, 136: 327, 171: 690, 326},
		{132, 446, 132, 8: 201, 128: 581, 580, 131: 598, 155: 596, 176: 597, 595},
		{205, 205, 205, 8: 205, 205, 205, 205, 21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 107: 296, 112: 315, 317, 314, 295, 118: 591, 297, 294, 146: 594},
		{205, 205, 205, 8: 205, 205, 205, 205, 21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 107: 296, 112: 315, 317, 314, 295, 118: 591, 297, 294, 146: 593},
		// 10
		{205, 205, 205, 8: 205, 205, 205, 205, 21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 107: 296, 112: 315, 317, 314, 295, 118: 591, 297, 294, 146: 592},
		{205, 205, 205, 8: 205, 205, 205, 205, 21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 107: 296, 112: 315, 317, 314, 295, 118: 591, 297, 294, 146: 590},
		{196, 196, 196, 8: 196, 196, 196, 196, 21: 196, 196, 196, 25: 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 41: 196, 43: 196, 196, 196, 196, 196},
		{195, 195, 195, 8: 195, 195, 195, 195, 21: 195, 195, 195, 25: 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 41: 195, 43: 195, 195, 195, 195, 195},
		{194, 194, 194, 8: 194, 194, 194, 194, 21: 194, 194, 194, 25: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 41: 194, 43: 194, 194, 194, 194, 194},
		// 15
		{193, 193, 193, 8: 193, 193, 193, 193, 21: 193, 193, 193, 25: 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 41: 193, 43: 193, 193, 193, 193, 193},
		{192, 192, 192, 8: 192, 192, 192, 192, 21: 192, 192, 192, 25: 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 41: 192, 43: 192, 192, 192, 192, 192},
		{191, 191, 191, 8: 191, 191, 191, 191, 21: 191, 191, 191, 25: 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 43: 191, 191, 191, 191, 191},
		{190, 190, 190, 8: 190, 190, 190, 190, 21: 190, 190, 190, 25: 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 43: 190, 190, 190, 190, 190},
		{189, 189, 189, 8: 189, 189, 189, 189, 21: 189, 189, 189, 25: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 43: 189, 189, 189, 189, 189},
		// 20
		{188, 188, 188, 8: 188, 188, 188, 188, 21: 188, 188, 188, 25: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 43: 188, 188, 188, 188, 188},
		{187, 187, 187, 8: 187, 187, 187, 187, 21: 187, 187, 187, 25: 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 43: 187, 187, 187, 187, 187},
		{186, 186, 186, 8: 186, 186, 186, 186, 21: 186, 186, 186, 25: 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 43: 186, 186, 186, 186, 186},
		{185, 185, 185, 8: 185, 185, 185, 185, 21: 185, 185, 185, 25: 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 43: 185, 185, 185, 185, 185},
		{184, 184, 184, 8: 184, 184, 184, 184, 21: 184, 184, 184, 25: 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 43: 184, 184, 184, 184, 184},
		// 25
		{183, 183, 183, 8: 183, 183, 183, 183, 21: 183, 183, 183, 25: 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 43: 183, 183, 183, 183, 183},
		{182, 182, 182, 8: 182, 182, 182, 182, 21: 182, 182, 182, 25: 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 43: 182, 182, 182, 182, 182},
		{181, 181, 181, 8: 181, 181, 181, 181, 21: 181, 181, 181, 25: 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 43: 181, 181, 181, 181, 181},
		{180, 180, 180, 8: 180, 180, 180, 180, 21: 180, 180, 180, 25: 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 43: 180, 180, 180, 180, 180},
		{179, 179, 179, 8: 179, 179, 179, 179, 21: 179, 179, 179, 25: 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 43: 179, 179, 179, 179, 179},
		// 30
		{178, 178, 178, 8: 178, 178, 178, 178, 21: 178, 178, 178, 25: 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 43: 178, 178, 178, 178, 178},
		{2: 568, 52: 114, 175: 567},
		{2: 173, 52: 173},
		{2: 172, 52: 172},
		{2: 329, 52: 114, 175: 328},
		// 35
		{150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 25: 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 43: 150, 150, 150, 150, 150, 50: 150},
		{149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 25: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 43: 149, 149, 149, 149, 149, 50: 149},
		{148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 25: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 43: 148, 148, 148, 148, 148, 50: 148},
		{147, 147, 147, 8: 147, 147, 147, 147, 21: 147, 147, 147, 25: 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 41: 147, 43: 147, 147, 147, 147, 147},
		{21: 54, 54, 54, 25: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 41: 54, 54, 54, 54, 54, 54, 54},
		// 40
		{21: 52, 52, 52, 25: 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 41: 52, 52, 52, 52, 52, 52, 52},
		{21: 51, 51, 51, 25: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 41: 51, 51, 51, 51, 51, 51, 51},
		{52: 157, 192: 330},
		{155, 155, 155, 8: 155, 155, 155, 155, 21: 155, 155, 155, 25: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 43: 155, 155, 155, 155, 155, 52: 113},
		{52: 331},
		// 45
		{2: 332, 169: 335, 334, 206: 333},
		{10: 279, 24: 279, 48: 279},
		{10: 563, 24: 159, 150: 564},
		{10: 154, 24: 154},
		{10: 152, 24: 152, 48: 336},
		// 50
		{212, 212, 212, 212, 212, 212, 212, 212, 12: 212, 212, 212, 212, 212, 212, 212, 212, 212, 139: 337, 338},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 347},
		{10: 151, 24: 151},
		{274, 274, 3: 274, 274, 274, 274, 274, 274, 274, 274, 274, 24: 274, 40: 274, 42: 274, 48: 274, 50: 274, 274, 53: 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274},
		{273, 273, 3: 273, 273, 273, 273, 273, 273, 273, 273, 273, 24: 273, 40: 273, 42: 273, 48: 273, 50: 273, 273, 53: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273},
		// 55
		{272, 272, 3: 272, 272, 272, 272, 272, 272, 272, 272, 272, 24: 272, 40: 272, 42: 272, 48: 272, 50: 272, 272, 53: 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272},
		{271, 271, 3: 271, 271, 271, 271, 271, 271, 271, 271, 271, 24: 271, 40: 271, 42: 271, 48: 271, 50: 271, 271, 53: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{270, 270, 3: 270, 270, 270, 270, 270, 270, 270, 270, 270, 24: 270, 40: 270, 42: 270, 48: 270, 50: 270, 270, 53: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270},
		{269, 269, 3: 269, 269, 269, 269, 269, 269, 269, 269, 269, 24: 269, 40: 269, 42: 269, 48: 269, 50: 269, 269, 53: 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269},
		{268, 268, 3: 268, 268, 268, 268, 268, 268, 268, 268, 268, 24: 268, 40: 268, 42: 268, 48: 268, 50: 268, 268, 53: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268},
		// 60
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 112, 112, 112, 25: 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 49: 406, 108: 440, 159: 442, 183: 561},
		{360, 365, 3: 378, 368, 369, 364, 363, 211, 10: 211, 359, 24: 211, 40: 211, 42: 211, 48: 384, 50: 211, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 560},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 559},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 558},
		// 65
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 472},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 557},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 556},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 555},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 554},
		// 70
		{357, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 358},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 112, 112, 112, 25: 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 49: 406, 108: 440, 159: 442, 183: 441},
		{360, 251, 3: 251, 251, 251, 364, 363, 251, 251, 251, 359, 24: 251, 40: 251, 42: 251, 48: 251, 50: 251, 361, 53: 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 362, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 108: 438},
		{346, 351, 339, 350, 352, 353, 349, 348, 9: 276, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 432, 194: 433, 434},
		// 75
		{2: 431},
		{2: 430},
		{262, 262, 3: 262, 262, 262, 262, 262, 262, 262, 262, 262, 24: 262, 40: 262, 42: 262, 48: 262, 50: 262, 262, 53: 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262},
		{261, 261, 3: 261, 261, 261, 261, 261, 261, 261, 261, 261, 24: 261, 40: 261, 42: 261, 48: 261, 50: 261, 261, 53: 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 429},
		// 80
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 428},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 427},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 426},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 425},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 424},
		// 85
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 423},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 422},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 421},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 420},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 419},
		// 90
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 418},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 417},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 416},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 415},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 414},
		// 95
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 413},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 412},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 108: 407},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 405},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 404},
		// 100
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 403},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 402},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 401},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 400},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 399},
		// 105
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 398},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 397},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 396},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 395},
		{360, 365, 3: 378, 368, 369, 364, 363, 219, 219, 219, 359, 24: 219, 40: 219, 42: 219, 48: 384, 50: 219, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		// 110
		{360, 365, 3: 378, 368, 369, 364, 363, 220, 220, 220, 359, 24: 220, 40: 220, 42: 220, 48: 384, 50: 220, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{360, 365, 3: 378, 368, 369, 364, 363, 221, 221, 221, 359, 24: 221, 40: 221, 42: 221, 48: 384, 50: 221, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{360, 365, 3: 378, 368, 369, 364, 363, 222, 222, 222, 359, 24: 222, 40: 222, 42: 222, 48: 384, 50: 222, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{360, 365, 3: 378, 368, 369, 364, 363, 223, 223, 223, 359, 24: 223, 40: 223, 42: 223, 48: 384, 50: 223, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{360, 365, 3: 378, 368, 369, 364, 363, 224, 224, 224, 359, 24: 224, 40: 224, 42: 224, 48: 384, 50: 224, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		// 115
		{360, 365, 3: 378, 368, 369, 364, 363, 225, 225, 225, 359, 24: 225, 40: 225, 42: 225, 48: 384, 50: 225, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{360, 365, 3: 378, 368, 369, 364, 363, 226, 226, 226, 359, 24: 226, 40: 226, 42: 226, 48: 384, 50: 226, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{360, 365, 3: 378, 368, 369, 364, 363, 227, 227, 227, 359, 24: 227, 40: 227, 42: 227, 48: 384, 50: 227, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{360, 365, 3: 378, 368, 369, 364, 363, 228, 228, 228, 359, 24: 228, 40: 228, 42: 228, 48: 384, 50: 228, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{360, 365, 3: 378, 368, 369, 364, 363, 229, 229, 229, 359, 24: 229, 40: 229, 42: 229, 48: 384, 50: 229, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		// 120
		{360, 365, 3: 378, 368, 369, 364, 363, 216, 216, 216, 359, 40: 216, 48: 384, 50: 216, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{10: 409, 40: 408},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 411},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 410},
		{360, 365, 3: 378, 368, 369, 364, 363, 215, 215, 215, 359, 40: 215, 48: 384, 50: 215, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		// 125
		{360, 365, 3: 378, 368, 369, 364, 363, 230, 230, 230, 359, 24: 230, 40: 230, 42: 230, 48: 230, 50: 230, 361, 53: 367, 366, 372, 373, 383, 379, 380, 230, 381, 230, 362, 230, 376, 375, 374, 370, 230, 230, 230, 377, 230, 382, 371, 230, 230, 230},
		{360, 365, 3: 378, 368, 369, 364, 363, 231, 231, 231, 359, 24: 231, 40: 231, 42: 231, 48: 231, 50: 231, 361, 53: 367, 366, 372, 373, 231, 379, 380, 231, 381, 231, 362, 231, 376, 375, 374, 370, 231, 231, 231, 377, 231, 231, 371, 231, 231, 231},
		{360, 365, 3: 378, 368, 369, 364, 363, 232, 232, 232, 359, 24: 232, 40: 232, 42: 232, 48: 232, 50: 232, 361, 53: 367, 366, 372, 373, 232, 379, 380, 232, 232, 232, 362, 232, 376, 375, 374, 370, 232, 232, 232, 377, 232, 232, 371, 232, 232, 232},
		{360, 365, 3: 378, 368, 369, 364, 363, 233, 233, 233, 359, 24: 233, 40: 233, 42: 233, 48: 233, 50: 233, 361, 53: 367, 366, 372, 373, 233, 379, 233, 233, 233, 233, 362, 233, 376, 375, 374, 370, 233, 233, 233, 377, 233, 233, 371, 233, 233, 233},
		{360, 365, 3: 378, 368, 369, 364, 363, 234, 234, 234, 359, 24: 234, 40: 234, 42: 234, 48: 234, 50: 234, 361, 53: 367, 366, 372, 373, 234, 234, 234, 234, 234, 234, 362, 234, 376, 375, 374, 370, 234, 234, 234, 377, 234, 234, 371, 234, 234, 234},
		// 130
		{360, 365, 3: 235, 368, 369, 364, 363, 235, 235, 235, 359, 24: 235, 40: 235, 42: 235, 48: 235, 50: 235, 361, 53: 367, 366, 372, 373, 235, 235, 235, 235, 235, 235, 362, 235, 376, 375, 374, 370, 235, 235, 235, 377, 235, 235, 371, 235, 235, 235},
		{360, 365, 3: 236, 368, 369, 364, 363, 236, 236, 236, 359, 24: 236, 40: 236, 42: 236, 48: 236, 50: 236, 361, 53: 367, 366, 372, 373, 236, 236, 236, 236, 236, 236, 362, 236, 236, 375, 374, 370, 236, 236, 236, 236, 236, 236, 371, 236, 236, 236},
		{360, 365, 3: 237, 368, 369, 364, 363, 237, 237, 237, 359, 24: 237, 40: 237, 42: 237, 48: 237, 50: 237, 361, 53: 367, 366, 372, 373, 237, 237, 237, 237, 237, 237, 362, 237, 237, 375, 374, 370, 237, 237, 237, 237, 237, 237, 371, 237, 237, 237},
		{360, 365, 3: 238, 368, 369, 364, 363, 238, 238, 238, 359, 24: 238, 40: 238, 42: 238, 48: 238, 50: 238, 361, 53: 367, 366, 238, 238, 238, 238, 238, 238, 238, 238, 362, 238, 238, 238, 238, 370, 238, 238, 238, 238, 238, 238, 371, 238, 238, 238},
		{360, 365, 3: 239, 368, 369, 364, 363, 239, 239, 239, 359, 24: 239, 40: 239, 42: 239, 48: 239, 50: 239, 361, 53: 367, 366, 239, 239, 239, 239, 239, 239, 239, 239, 362, 239, 239, 239, 239, 370, 239, 239, 239, 239, 239, 239, 371, 239, 239, 239},
		// 135
		{360, 365, 3: 240, 368, 369, 364, 363, 240, 240, 240, 359, 24: 240, 40: 240, 42: 240, 48: 240, 50: 240, 361, 53: 367, 366, 240, 240, 240, 240, 240, 240, 240, 240, 362, 240, 240, 240, 240, 370, 240, 240, 240, 240, 240, 240, 371, 240, 240, 240},
		{360, 365, 3: 241, 368, 369, 364, 363, 241, 241, 241, 359, 24: 241, 40: 241, 42: 241, 48: 241, 50: 241, 361, 53: 367, 366, 241, 241, 241, 241, 241, 241, 241, 241, 362, 241, 241, 241, 241, 370, 241, 241, 241, 241, 241, 241, 371, 241, 241, 241},
		{360, 365, 3: 242, 368, 369, 364, 363, 242, 242, 242, 359, 24: 242, 40: 242, 42: 242, 48: 242, 50: 242, 361, 53: 367, 366, 242, 242, 242, 242, 242, 242, 242, 242, 362, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242},
		{360, 365, 3: 243, 368, 369, 364, 363, 243, 243, 243, 359, 24: 243, 40: 243, 42: 243, 48: 243, 50: 243, 361, 53: 367, 366, 243, 243, 243, 243, 243, 243, 243, 243, 362, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{360, 365, 3: 244, 244, 244, 364, 363, 244, 244, 244, 359, 24: 244, 40: 244, 42: 244, 48: 244, 50: 244, 361, 53: 367, 366, 244, 244, 244, 244, 244, 244, 244, 244, 362, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244},
		// 140
		{360, 365, 3: 245, 245, 245, 364, 363, 245, 245, 245, 359, 24: 245, 40: 245, 42: 245, 48: 245, 50: 245, 361, 53: 367, 366, 245, 245, 245, 245, 245, 245, 245, 245, 362, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245},
		{360, 246, 3: 246, 246, 246, 364, 363, 246, 246, 246, 359, 24: 246, 40: 246, 42: 246, 48: 246, 50: 246, 361, 53: 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 362, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246},
		{360, 247, 3: 247, 247, 247, 364, 363, 247, 247, 247, 359, 24: 247, 40: 247, 42: 247, 48: 247, 50: 247, 361, 53: 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 362, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247},
		{360, 248, 3: 248, 248, 248, 364, 363, 248, 248, 248, 359, 24: 248, 40: 248, 42: 248, 48: 248, 50: 248, 361, 53: 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 362, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248},
		{263, 263, 3: 263, 263, 263, 263, 263, 263, 263, 263, 263, 24: 263, 40: 263, 42: 263, 48: 263, 50: 263, 263, 53: 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263},
		// 145
		{264, 264, 3: 264, 264, 264, 264, 264, 264, 264, 264, 264, 24: 264, 40: 264, 42: 264, 48: 264, 50: 264, 264, 53: 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264},
		{360, 365, 3: 378, 368, 369, 364, 363, 9: 278, 278, 359, 48: 384, 51: 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{9: 275, 436},
		{9: 435},
		{265, 265, 3: 265, 265, 265, 265, 265, 265, 265, 265, 265, 24: 265, 40: 265, 42: 265, 48: 265, 50: 265, 265, 53: 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265},
		// 150
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 437},
		{360, 365, 3: 378, 368, 369, 364, 363, 9: 277, 277, 359, 48: 384, 51: 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{10: 409, 50: 439},
		{266, 266, 3: 266, 266, 266, 266, 266, 266, 266, 266, 266, 24: 266, 40: 266, 42: 266, 48: 266, 50: 266, 266, 53: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266},
		{9: 553, 409},
		// 155
		{9: 527},
		{21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 107: 444, 112: 315, 317, 314, 443, 143: 445},
		{165, 165, 165, 8: 165, 165, 11: 165, 21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 165, 107: 444, 112: 315, 317, 314, 443, 143: 525, 180: 526},
		{165, 165, 165, 8: 165, 165, 11: 165, 21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 165, 107: 444, 112: 315, 317, 314, 443, 143: 525, 180: 524},
		{132, 446, 9: 108, 11: 132, 128: 447, 449, 145: 450, 162: 448},
		// 160
		{128, 128, 128, 9: 128, 128, 128, 21: 321, 322, 323, 107: 457, 144: 461, 149: 522},
		{131, 2: 131, 9: 110, 110, 131},
		{9: 111},
		{452, 11: 96, 165: 451, 453},
		{9: 107, 107},
		// 165
		{518, 9: 109, 109, 95},
		{132, 446, 9: 100, 11: 132, 21: 100, 100, 100, 25: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 41: 100, 43: 100, 100, 100, 100, 100, 128: 447, 449, 145: 474, 160: 475},
		{11: 454},
		{346, 456, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 321, 322, 323, 41: 460, 49: 455, 218, 107: 457, 144: 458, 154: 459},
		{360, 365, 3: 378, 368, 369, 364, 363, 11: 359, 48: 384, 50: 217, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		// 170
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 472, 473},
		{130, 130, 130, 130, 130, 130, 130, 130, 9: 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 41: 130, 50: 130},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 321, 322, 323, 41: 468, 49: 455, 218, 107: 465, 154: 467},
		{50: 466},
		{128, 128, 128, 128, 128, 128, 128, 128, 12: 128, 128, 128, 128, 128, 128, 128, 128, 128, 321, 322, 323, 107: 457, 144: 461, 149: 462},
		// 175
		{127, 127, 127, 127, 127, 127, 127, 127, 9: 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 321, 322, 323, 107: 465},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 463},
		{360, 365, 3: 378, 368, 369, 364, 363, 11: 359, 48: 384, 50: 464, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{103, 9: 103, 103, 103},
		{129, 129, 129, 129, 129, 129, 129, 129, 9: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 41: 129, 50: 129},
		// 180
		{105, 9: 105, 105, 105},
		{50: 471},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 469},
		{360, 365, 3: 378, 368, 369, 364, 363, 11: 359, 48: 384, 50: 470, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{102, 9: 102, 102, 102},
		// 185
		{104, 9: 104, 104, 104},
		{360, 256, 3: 256, 256, 256, 364, 363, 256, 256, 256, 359, 24: 256, 40: 256, 42: 256, 48: 256, 50: 256, 361, 53: 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 362, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256},
		{101, 9: 101, 101, 101},
		{9: 517},
		{9: 124, 21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 107: 296, 112: 315, 317, 314, 295, 118: 479, 297, 294, 148: 478, 157: 476, 477, 179: 480},
		// 190
		{9: 126, 514},
		{9: 123},
		{9: 122, 122},
		{132, 446, 132, 9: 108, 108, 132, 128: 447, 482, 131: 483, 145: 450, 162: 484},
		{9: 481},
		// 195
		{99, 9: 99, 99, 99},
		{487, 2: 486, 11: 96, 165: 451, 453, 485},
		{9: 120, 120},
		{9: 119, 119},
		{491, 8: 146, 146, 146, 490, 21: 146, 146, 146, 25: 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 43: 146, 146, 146, 146, 146, 146, 52: 146},
		// 200
		{143, 8: 143, 143, 143, 143, 21: 143, 143, 143, 25: 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 43: 143, 143, 143, 143, 143, 143, 52: 143},
		{132, 446, 132, 9: 100, 11: 132, 21: 100, 100, 100, 25: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 41: 100, 43: 100, 100, 100, 100, 100, 128: 447, 482, 131: 488, 145: 474, 160: 475},
		{9: 489},
		{142, 8: 142, 142, 142, 142, 21: 142, 142, 142, 25: 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 43: 142, 142, 142, 142, 142, 142, 52: 142},
		{128, 128, 128, 128, 128, 128, 128, 128, 12: 128, 128, 128, 128, 128, 128, 128, 128, 128, 321, 322, 323, 41: 502, 50: 128, 107: 457, 144: 503, 149: 501},
		// 205
		{2: 494, 9: 116, 21: 137, 137, 137, 25: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 41: 137, 43: 137, 137, 137, 137, 137, 173: 495, 493, 193: 492},
		{21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 107: 296, 112: 315, 317, 314, 295, 118: 479, 297, 294, 148: 478, 157: 476, 499},
		{9: 498},
		{9: 118, 118},
		{9: 115, 496},
		// 210
		{2: 497},
		{9: 117, 117},
		{135, 8: 135, 135, 135, 135, 21: 135, 135, 135, 25: 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 43: 135, 135, 135, 135, 135, 135, 52: 135},
		{9: 500},
		{136, 8: 136, 136, 136, 136, 21: 136, 136, 136, 25: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 43: 136, 136, 136, 136, 136, 136, 52: 136},
		// 215
		{346, 510, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 455, 218, 154: 511},
		{128, 128, 128, 128, 128, 128, 128, 128, 12: 128, 128, 128, 128, 128, 128, 128, 128, 128, 321, 322, 323, 107: 457, 144: 461, 149: 507},
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 321, 322, 323, 41: 504, 50: 127, 107: 465},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 505},
		{360, 365, 3: 378, 368, 369, 364, 363, 11: 359, 48: 384, 50: 506, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		// 220
		{139, 8: 139, 139, 139, 139, 21: 139, 139, 139, 25: 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 43: 139, 139, 139, 139, 139, 139, 52: 139},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 508},
		{360, 365, 3: 378, 368, 369, 364, 363, 11: 359, 48: 384, 50: 509, 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{140, 8: 140, 140, 140, 140, 21: 140, 140, 140, 25: 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 43: 140, 140, 140, 140, 140, 140, 52: 140},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 472, 513},
		// 225
		{50: 512},
		{141, 8: 141, 141, 141, 141, 21: 141, 141, 141, 25: 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 43: 141, 141, 141, 141, 141, 141, 52: 141},
		{138, 8: 138, 138, 138, 138, 21: 138, 138, 138, 25: 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 43: 138, 138, 138, 138, 138, 138, 52: 138},
		{21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 107: 296, 112: 315, 317, 314, 295, 118: 479, 297, 294, 141: 515, 148: 516},
		{9: 125},
		// 230
		{9: 121, 121},
		{106, 9: 106, 106, 106},
		{9: 98, 21: 98, 98, 98, 25: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 41: 98, 43: 98, 98, 98, 98, 98, 185: 519},
		{9: 124, 21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 107: 296, 112: 315, 317, 314, 295, 118: 479, 297, 294, 148: 478, 157: 476, 477, 179: 520},
		{9: 521},
		// 235
		{97, 9: 97, 97, 97},
		{134, 446, 134, 9: 134, 134, 134, 128: 523},
		{133, 2: 133, 9: 133, 133, 133},
		{166, 166, 166, 8: 166, 166, 11: 166, 40: 166},
		{164, 164, 164, 8: 164, 164, 11: 164, 40: 164},
		// 240
		{167, 167, 167, 8: 167, 167, 11: 167, 40: 167},
		{346, 250, 339, 250, 250, 250, 349, 348, 250, 250, 250, 250, 355, 354, 340, 341, 342, 343, 344, 356, 345, 24: 250, 40: 250, 42: 250, 48: 250, 528, 250, 250, 529, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250},
		{360, 249, 3: 249, 249, 249, 364, 363, 249, 249, 249, 359, 24: 249, 40: 249, 42: 249, 48: 249, 50: 249, 361, 53: 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 362, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249},
		{89, 89, 89, 89, 89, 89, 89, 89, 11: 535, 89, 89, 89, 89, 89, 89, 89, 89, 89, 51: 536, 89, 147: 534, 151: 533, 531, 532, 178: 530},
		{10: 546, 24: 159, 150: 551},
		// 245
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 542, 52: 543, 156: 544},
		{11: 535, 48: 540, 51: 536, 147: 541},
		{88, 88, 88, 88, 88, 88, 88, 88, 12: 88, 88, 88, 88, 88, 88, 88, 88, 88, 52: 88},
		{11: 87, 48: 87, 51: 87},
		{212, 212, 212, 212, 212, 212, 212, 212, 12: 212, 212, 212, 212, 212, 212, 212, 212, 212, 139: 337, 538},
		// 250
		{2: 537},
		{11: 84, 48: 84, 51: 84},
		{50: 539},
		{11: 85, 48: 85, 51: 85},
		{90, 90, 90, 90, 90, 90, 90, 90, 12: 90, 90, 90, 90, 90, 90, 90, 90, 90, 52: 90},
		// 255
		{11: 86, 48: 86, 51: 86},
		{360, 365, 3: 378, 368, 369, 364, 363, 94, 10: 94, 359, 24: 94, 48: 384, 51: 361, 53: 367, 366, 372, 373, 383, 379, 380, 388, 381, 392, 362, 386, 376, 375, 374, 370, 390, 387, 385, 377, 394, 382, 371, 391, 389, 393},
		{89, 89, 89, 89, 89, 89, 89, 89, 11: 535, 89, 89, 89, 89, 89, 89, 89, 89, 89, 51: 536, 89, 147: 534, 151: 533, 531, 532, 178: 545},
		{10: 92, 24: 92},
		{10: 546, 24: 159, 150: 547},
		// 260
		{89, 89, 89, 89, 89, 89, 89, 89, 11: 535, 89, 89, 89, 89, 89, 89, 89, 89, 89, 24: 158, 51: 536, 89, 147: 534, 151: 533, 549, 532},
		{24: 548},
		{8: 93, 10: 93, 24: 93},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 542, 52: 543, 156: 550},
		{10: 91, 24: 91},
		// 265
		{24: 552},
		{260, 260, 3: 260, 260, 260, 260, 260, 260, 260, 260, 260, 24: 260, 40: 260, 42: 260, 48: 260, 50: 260, 260, 53: 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260},
		{267, 267, 3: 267, 267, 267, 267, 267, 267, 267, 267, 267, 24: 267, 40: 267, 42: 267, 48: 267, 50: 267, 267, 53: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267},
		{360, 252, 3: 252, 252, 252, 364, 363, 252, 252, 252, 359, 24: 252, 40: 252, 42: 252, 48: 252, 50: 252, 361, 53: 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 362, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252},
		{360, 253, 3: 253, 253, 253, 364, 363, 253, 253, 253, 359, 24: 253, 40: 253, 42: 253, 48: 253, 50: 253, 361, 53: 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 362, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253},
		// 270
		{360, 254, 3: 254, 254, 254, 364, 363, 254, 254, 254, 359, 24: 254, 40: 254, 42: 254, 48: 254, 50: 254, 361, 53: 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 362, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254},
		{360, 255, 3: 255, 255, 255, 364, 363, 255, 255, 255, 359, 24: 255, 40: 255, 42: 255, 48: 255, 50: 255, 361, 53: 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 362, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{360, 257, 3: 257, 257, 257, 364, 363, 257, 257, 257, 359, 24: 257, 40: 257, 42: 257, 48: 257, 50: 257, 361, 53: 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 362, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257},
		{360, 258, 3: 258, 258, 258, 364, 363, 258, 258, 258, 359, 24: 258, 40: 258, 42: 258, 48: 258, 50: 258, 361, 53: 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 362, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258},
		{360, 259, 3: 259, 259, 259, 364, 363, 259, 259, 259, 359, 24: 259, 40: 259, 42: 259, 48: 259, 50: 259, 361, 53: 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 362, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259},
		// 275
		{9: 562},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 528, 52: 529},
		{2: 332, 24: 158, 169: 335, 566},
		{24: 565},
		{156, 156, 156, 8: 156, 156, 156, 156, 21: 156, 156, 156, 25: 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 43: 156, 156, 156, 156, 156},
		// 280
		{10: 153, 24: 153},
		{52: 177, 190: 569},
		{174, 174, 174, 8: 174, 174, 174, 174, 21: 174, 174, 174, 25: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 43: 174, 174, 174, 174, 174, 52: 113},
		{52: 570},
		{21: 176, 176, 176, 25: 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 191: 571},
		// 285
		{21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 107: 444, 112: 315, 317, 314, 443, 143: 574, 181: 573, 211: 572},
		{21: 321, 322, 323, 588, 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 107: 444, 112: 315, 317, 314, 443, 143: 574, 181: 589},
		{21: 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171},
		{132, 446, 132, 8: 576, 40: 145, 128: 581, 580, 131: 578, 164: 579, 182: 577, 212: 575},
		{8: 585, 10: 586},
		// 290
		{21: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168},
		{8: 163, 10: 163},
		{8: 161, 10: 161, 40: 144},
		{40: 583},
		{582, 2: 486, 167: 485},
		// 295
		{131, 2: 131},
		{132, 446, 132, 128: 581, 580, 131: 488},
		{212, 212, 212, 212, 212, 212, 212, 212, 12: 212, 212, 212, 212, 212, 212, 212, 212, 212, 139: 337, 584},
		{8: 160, 10: 160},
		{21: 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169},
		// 300
		{132, 446, 132, 40: 145, 128: 581, 580, 131: 578, 164: 579, 182: 587},
		{8: 162, 10: 162},
		{175, 175, 175, 8: 175, 175, 175, 175, 21: 175, 175, 175, 25: 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 43: 175, 175, 175, 175, 175},
		{21: 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170},
		{206, 206, 206, 8: 206, 206, 206, 206},
		// 305
		{204, 204, 204, 8: 204, 204, 204, 204},
		{207, 207, 207, 8: 207, 207, 207, 207},
		{208, 208, 208, 8: 208, 208, 208, 208},
		{209, 209, 209, 8: 209, 209, 209, 209},
		{8: 689},
		// 310
		{8: 203, 10: 203},
		{8: 200, 10: 687},
		{8: 199, 10: 199, 21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 198, 52: 46, 107: 296, 112: 315, 317, 314, 295, 118: 599, 297, 294, 136: 602, 161: 600, 199: 603, 601},
		{132, 446, 132, 8: 201, 128: 581, 580, 131: 686, 155: 596, 176: 597, 595},
		{48: 684},
		// 315
		{52: 50, 187: 605},
		{21: 48, 48, 48, 25: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 41: 48, 43: 48, 48, 48, 48, 48, 52: 48},
		{21: 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 52: 45, 107: 296, 112: 315, 317, 314, 295, 118: 599, 297, 294, 136: 604},
		{21: 47, 47, 47, 25: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 41: 47, 43: 47, 47, 47, 47, 47, 52: 47},
		{52: 606, 121: 607},
		// 320
		{74, 74, 74, 74, 74, 74, 74, 74, 74, 12: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 41: 74, 43: 74, 74, 74, 74, 74, 52: 74, 83: 74, 74, 74, 74, 74, 74, 74, 74, 74, 93: 74, 74, 186: 608},
		{21: 49, 49, 49, 25: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 41: 49, 49, 49, 49, 49, 49, 49},
		{346, 351, 609, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 321, 322, 323, 70, 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 49: 406, 52: 606, 83: 627, 632, 617, 631, 618, 628, 629, 630, 625, 93: 633, 626, 107: 296, 610, 112: 315, 317, 314, 295, 624, 118: 599, 297, 294, 612, 613, 615, 616, 611, 614, 623, 136: 622, 163: 620, 196: 621, 619},
		{274, 274, 3: 274, 274, 274, 274, 274, 274, 10: 274, 274, 40: 682, 48: 274, 51: 274, 53: 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274},
		{8: 213, 213, 409},
		// 325
		{83, 83, 83, 83, 83, 83, 83, 83, 83, 12: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 41: 83, 43: 83, 83, 83, 83, 83, 52: 83, 83: 83, 83, 83, 83, 83, 83, 83, 83, 83, 93: 83, 83, 111: 83},
		{82, 82, 82, 82, 82, 82, 82, 82, 82, 12: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 41: 82, 43: 82, 82, 82, 82, 82, 52: 82, 83: 82, 82, 82, 82, 82, 82, 82, 82, 82, 93: 82, 82, 111: 82},
		{81, 81, 81, 81, 81, 81, 81, 81, 81, 12: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 41: 81, 43: 81, 81, 81, 81, 81, 52: 81, 83: 81, 81, 81, 81, 81, 81, 81, 81, 81, 93: 81, 81, 111: 81},
		{80, 80, 80, 80, 80, 80, 80, 80, 80, 12: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 41: 80, 43: 80, 80, 80, 80, 80, 52: 80, 83: 80, 80, 80, 80, 80, 80, 80, 80, 80, 93: 80, 80, 111: 80},
		{79, 79, 79, 79, 79, 79, 79, 79, 79, 12: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 41: 79, 43: 79, 79, 79, 79, 79, 52: 79, 83: 79, 79, 79, 79, 79, 79, 79, 79, 79, 93: 79, 79, 111: 79},
		// 330
		{78, 78, 78, 78, 78, 78, 78, 78, 78, 12: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 41: 78, 43: 78, 78, 78, 78, 78, 52: 78, 83: 78, 78, 78, 78, 78, 78, 78, 78, 78, 93: 78, 78, 111: 78},
		{212, 212, 212, 212, 212, 212, 212, 212, 12: 212, 212, 212, 212, 212, 212, 212, 212, 212, 139: 337, 679},
		{40: 677},
		{24: 676},
		{72, 72, 72, 72, 72, 72, 72, 72, 72, 12: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 41: 72, 43: 72, 72, 72, 72, 72, 52: 72, 83: 72, 72, 72, 72, 72, 72, 72, 72, 72, 93: 72, 72},
		// 335
		{346, 351, 609, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 321, 322, 323, 69, 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 49: 406, 52: 606, 83: 627, 632, 617, 631, 618, 628, 629, 630, 625, 93: 633, 626, 107: 296, 610, 112: 315, 317, 314, 295, 624, 118: 599, 297, 294, 612, 613, 615, 616, 611, 614, 623, 136: 622, 163: 675},
		{68, 68, 68, 68, 68, 68, 68, 68, 68, 12: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 41: 68, 43: 68, 68, 68, 68, 68, 52: 68, 83: 68, 68, 68, 68, 68, 68, 68, 68, 68, 93: 68, 68},
		{67, 67, 67, 67, 67, 67, 67, 67, 67, 12: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 41: 67, 43: 67, 67, 67, 67, 67, 52: 67, 83: 67, 67, 67, 67, 67, 67, 67, 67, 67, 93: 67, 67},
		{8: 674},
		{668},
		// 340
		{664},
		{660},
		{346, 351, 609, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 52: 606, 83: 627, 632, 617, 631, 618, 628, 629, 630, 625, 93: 633, 626, 108: 610, 116: 624, 121: 612, 613, 615, 616, 611, 614, 654},
		{640},
		{2: 638},
		// 345
		{8: 637},
		{8: 636},
		{346, 351, 339, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 108: 610, 116: 634},
		{8: 635},
		{55, 55, 55, 55, 55, 55, 55, 55, 55, 12: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 41: 55, 43: 55, 55, 55, 55, 55, 52: 55, 83: 55, 55, 55, 55, 55, 55, 55, 55, 55, 93: 55, 55, 111: 55},
		// 350
		{56, 56, 56, 56, 56, 56, 56, 56, 56, 12: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 41: 56, 43: 56, 56, 56, 56, 56, 52: 56, 83: 56, 56, 56, 56, 56, 56, 56, 56, 56, 93: 56, 56, 111: 56},
		{57, 57, 57, 57, 57, 57, 57, 57, 57, 12: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 41: 57, 43: 57, 57, 57, 57, 57, 52: 57, 83: 57, 57, 57, 57, 57, 57, 57, 57, 57, 93: 57, 57, 111: 57},
		{8: 639},
		{58, 58, 58, 58, 58, 58, 58, 58, 58, 12: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 41: 58, 43: 58, 58, 58, 58, 58, 52: 58, 83: 58, 58, 58, 58, 58, 58, 58, 58, 58, 93: 58, 58, 111: 58},
		{346, 351, 339, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 321, 322, 323, 25: 312, 304, 313, 309, 320, 308, 306, 307, 305, 310, 318, 316, 319, 311, 303, 41: 300, 43: 301, 299, 324, 302, 298, 49: 406, 107: 296, 610, 112: 315, 317, 314, 295, 641, 118: 599, 297, 294, 136: 642},
		// 355
		{8: 648},
		{346, 351, 339, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 108: 610, 116: 643},
		{8: 644},
		{346, 351, 339, 350, 352, 353, 349, 348, 9: 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 108: 610, 116: 645},
		{9: 646},
		// 360
		{346, 351, 609, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 52: 606, 83: 627, 632, 617, 631, 618, 628, 629, 630, 625, 93: 633, 626, 108: 610, 116: 624, 121: 612, 613, 615, 616, 611, 614, 647},
		{59, 59, 59, 59, 59, 59, 59, 59, 59, 12: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 41: 59, 43: 59, 59, 59, 59, 59, 52: 59, 83: 59, 59, 59, 59, 59, 59, 59, 59, 59, 93: 59, 59, 111: 59},
		{346, 351, 339, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 108: 610, 116: 649},
		{8: 650},
		{346, 351, 339, 350, 352, 353, 349, 348, 9: 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 108: 610, 116: 651},
		// 365
		{9: 652},
		{346, 351, 609, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 52: 606, 83: 627, 632, 617, 631, 618, 628, 629, 630, 625, 93: 633, 626, 108: 610, 116: 624, 121: 612, 613, 615, 616, 611, 614, 653},
		{60, 60, 60, 60, 60, 60, 60, 60, 60, 12: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 41: 60, 43: 60, 60, 60, 60, 60, 52: 60, 83: 60, 60, 60, 60, 60, 60, 60, 60, 60, 93: 60, 60, 111: 60},
		{83: 655},
		{656},
		// 370
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 108: 657},
		{9: 658, 409},
		{8: 659},
		{61, 61, 61, 61, 61, 61, 61, 61, 61, 12: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 41: 61, 43: 61, 61, 61, 61, 61, 52: 61, 83: 61, 61, 61, 61, 61, 61, 61, 61, 61, 93: 61, 61, 111: 61},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 108: 661},
		// 375
		{9: 662, 409},
		{346, 351, 609, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 52: 606, 83: 627, 632, 617, 631, 618, 628, 629, 630, 625, 93: 633, 626, 108: 610, 116: 624, 121: 612, 613, 615, 616, 611, 614, 663},
		{62, 62, 62, 62, 62, 62, 62, 62, 62, 12: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 41: 62, 43: 62, 62, 62, 62, 62, 52: 62, 83: 62, 62, 62, 62, 62, 62, 62, 62, 62, 93: 62, 62, 111: 62},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 108: 665},
		{9: 666, 409},
		// 380
		{346, 351, 609, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 52: 606, 83: 627, 632, 617, 631, 618, 628, 629, 630, 625, 93: 633, 626, 108: 610, 116: 624, 121: 612, 613, 615, 616, 611, 614, 667},
		{63, 63, 63, 63, 63, 63, 63, 63, 63, 12: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 41: 63, 43: 63, 63, 63, 63, 63, 52: 63, 83: 63, 63, 63, 63, 63, 63, 63, 63, 63, 93: 63, 63, 111: 63},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 108: 669},
		{9: 670, 409},
		{346, 351, 609, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 52: 606, 83: 627, 632, 617, 631, 618, 628, 629, 630, 625, 93: 633, 626, 108: 610, 116: 624, 121: 612, 613, 615, 616, 611, 614, 671},
		// 385
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 12: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 41: 65, 43: 65, 65, 65, 65, 65, 52: 65, 83: 65, 65, 65, 65, 65, 65, 65, 65, 65, 93: 65, 65, 111: 672},
		{346, 351, 609, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 52: 606, 83: 627, 632, 617, 631, 618, 628, 629, 630, 625, 93: 633, 626, 108: 610, 116: 624, 121: 612, 613, 615, 616, 611, 614, 673},
		{64, 64, 64, 64, 64, 64, 64, 64, 64, 12: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 41: 64, 43: 64, 64, 64, 64, 64, 52: 64, 83: 64, 64, 64, 64, 64, 64, 64, 64, 64, 93: 64, 64, 111: 64},
		{66, 66, 66, 66, 66, 66, 66, 66, 66, 12: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 41: 66, 43: 66, 66, 66, 66, 66, 52: 66, 83: 66, 66, 66, 66, 66, 66, 66, 66, 66, 93: 66, 66, 111: 66},
		{71, 71, 71, 71, 71, 71, 71, 71, 71, 12: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 41: 71, 43: 71, 71, 71, 71, 71, 52: 71, 83: 71, 71, 71, 71, 71, 71, 71, 71, 71, 93: 71, 71},
		// 390
		{73, 73, 73, 73, 73, 73, 73, 73, 73, 12: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 41: 73, 73, 73, 73, 73, 73, 73, 52: 73, 83: 73, 73, 73, 73, 73, 73, 73, 73, 73, 93: 73, 73, 111: 73},
		{346, 351, 609, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 52: 606, 83: 627, 632, 617, 631, 618, 628, 629, 630, 625, 93: 633, 626, 108: 610, 116: 624, 121: 612, 613, 615, 616, 611, 614, 678},
		{75, 75, 75, 75, 75, 75, 75, 75, 75, 12: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 41: 75, 43: 75, 75, 75, 75, 75, 52: 75, 83: 75, 75, 75, 75, 75, 75, 75, 75, 75, 93: 75, 75, 111: 75},
		{40: 680},
		{346, 351, 609, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 52: 606, 83: 627, 632, 617, 631, 618, 628, 629, 630, 625, 93: 633, 626, 108: 610, 116: 624, 121: 612, 613, 615, 616, 611, 614, 681},
		// 395
		{76, 76, 76, 76, 76, 76, 76, 76, 76, 12: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 41: 76, 43: 76, 76, 76, 76, 76, 52: 76, 83: 76, 76, 76, 76, 76, 76, 76, 76, 76, 93: 76, 76, 111: 76},
		{346, 351, 609, 350, 352, 353, 349, 348, 214, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 406, 52: 606, 83: 627, 632, 617, 631, 618, 628, 629, 630, 625, 93: 633, 626, 108: 610, 116: 624, 121: 612, 613, 615, 616, 611, 614, 683},
		{77, 77, 77, 77, 77, 77, 77, 77, 77, 12: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 41: 77, 43: 77, 77, 77, 77, 77, 52: 77, 83: 77, 77, 77, 77, 77, 77, 77, 77, 77, 93: 77, 77, 111: 77},
		{346, 351, 339, 350, 352, 353, 349, 348, 12: 355, 354, 340, 341, 342, 343, 344, 356, 345, 49: 542, 52: 543, 156: 685},
		{8: 197, 10: 197},
		// 400
		{8: 199, 10: 199, 48: 198, 161: 600},
		{132, 446, 132, 128: 581, 580, 131: 686, 155: 688},
		{8: 202, 10: 202},
		{210, 210, 210, 210, 210, 210, 210, 210, 210, 12: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 41: 210, 210, 210, 210, 210, 210, 210, 52: 210, 83: 210, 210, 210, 210, 210, 210, 210, 210, 210, 93: 210, 210},
		{21: 53, 53, 53, 25: 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 41: 53, 53, 53, 53, 53, 53, 53},
		// 405
		{212, 212, 212, 212, 212, 212, 212, 212, 12: 212, 212, 212, 212, 212, 212, 212, 212, 212, 139: 337, 692},
		{42: 282},
		{79: 715, 717, 95: 705, 706, 707, 702, 703, 704, 708, 712, 709, 699, 710, 711, 109: 716, 714, 117: 713, 130: 697, 132: 696, 701, 698, 700, 137: 695, 209: 694},
		{42: 284},
		{42: 44, 79: 715, 717, 95: 705, 706, 707, 702, 703, 704, 708, 712, 709, 699, 710, 711, 109: 716, 714, 117: 713, 130: 697, 132: 760, 701, 698, 700},
		// 410
		{42: 43, 79: 43, 43, 43, 43, 92: 43, 95: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43},
		{42: 39, 79: 39, 39, 39, 39, 92: 39, 95: 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39},
		{42: 38, 79: 38, 38, 38, 38, 92: 38, 95: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{80: 717, 109: 782, 714},
		{42: 36, 79: 36, 36, 36, 36, 92: 36, 95: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
		// 415
		{81: 29, 29, 92: 770, 168: 768, 201: 769, 767},
		{80: 717, 109: 764, 714},
		{2: 761},
		{2: 756},
		{2: 732, 79: 734, 207: 733},
		// 420
		{79: 715, 717, 109: 716, 714, 117: 731},
		{42: 17, 79: 17, 17, 17, 17, 92: 17, 95: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{80: 717, 109: 729, 714},
		{80: 717, 109: 727, 714},
		{79: 715, 717, 109: 716, 714, 117: 726},
		// 425
		{2: 722},
		{80: 717, 109: 720, 714},
		{42: 7, 79: 7, 7, 7, 7, 92: 7, 95: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{79: 5, 719},
		{42: 4, 79: 4, 4, 4, 4, 92: 4, 95: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		// 430
		{79: 718},
		{79: 2, 2},
		{42: 3, 79: 3, 3, 3, 3, 92: 3, 95: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{79: 1, 1},
		{79: 721},
		// 435
		{42: 8, 79: 8, 8, 8, 8, 92: 8, 95: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{79: 723, 717, 109: 724, 714},
		{42: 13, 79: 13, 13, 13, 13, 92: 13, 95: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{79: 725},
		{42: 9, 79: 9, 9, 9, 9, 92: 9, 95: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		// 440
		{42: 14, 79: 14, 14, 14, 14, 92: 14, 95: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{79: 728},
		{42: 15, 79: 15, 15, 15, 15, 92: 15, 95: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{79: 730},
		{42: 16, 79: 16, 16, 16, 16, 92: 16, 95: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		// 445
		{42: 18, 79: 18, 18, 18, 18, 92: 18, 95: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{79: 715, 717, 109: 716, 714, 117: 741, 138: 755},
		{2: 735, 9: 116, 141: 737, 173: 736, 738},
		{42: 10, 79: 10, 10, 10, 10, 92: 10, 95: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{9: 118, 118, 141: 752},
		// 450
		{9: 115, 744},
		{9: 742},
		{9: 739},
		{79: 715, 717, 109: 716, 714, 117: 741, 138: 740},
		{42: 19, 79: 19, 19, 19, 19, 92: 19, 95: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		// 455
		{42: 6, 79: 6, 6, 6, 6, 92: 6, 95: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{79: 715, 717, 109: 716, 714, 117: 741, 138: 743},
		{42: 21, 79: 21, 21, 21, 21, 92: 21, 95: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{2: 745, 141: 746},
		{9: 117, 117, 141: 749},
		// 460
		{9: 747},
		{79: 715, 717, 109: 716, 714, 117: 741, 138: 748},
		{42: 20, 79: 20, 20, 20, 20, 92: 20, 95: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{9: 750},
		{79: 715, 717, 109: 716, 714, 117: 741, 138: 751},
		// 465
		{42: 11, 79: 11, 11, 11, 11, 92: 11, 95: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{9: 753},
		{79: 715, 717, 109: 716, 714, 117: 741, 138: 754},
		{42: 12, 79: 12, 12, 12, 12, 92: 12, 95: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{42: 22, 79: 22, 22, 22, 22, 92: 22, 95: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		// 470
		{79: 757},
		{79: 715, 717, 41, 41, 92: 41, 95: 705, 706, 707, 702, 703, 704, 708, 712, 709, 699, 710, 711, 109: 716, 714, 117: 713, 130: 697, 132: 696, 701, 698, 700, 137: 758, 142: 759},
		{79: 715, 717, 40, 40, 92: 40, 95: 705, 706, 707, 702, 703, 704, 708, 712, 709, 699, 710, 711, 109: 716, 714, 117: 713, 130: 697, 132: 760, 701, 698, 700},
		{81: 32, 32, 92: 32},
		{42: 42, 79: 42, 42, 42, 42, 92: 42, 95: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		// 475
		{79: 762},
		{79: 715, 717, 41, 41, 92: 41, 95: 705, 706, 707, 702, 703, 704, 708, 712, 709, 699, 710, 711, 109: 716, 714, 117: 713, 130: 697, 132: 696, 701, 698, 700, 137: 758, 142: 763},
		{81: 33, 33, 92: 33},
		{79: 765},
		{79: 715, 717, 41, 41, 92: 41, 95: 705, 706, 707, 702, 703, 704, 708, 712, 709, 699, 710, 711, 109: 716, 714, 117: 713, 130: 697, 132: 696, 701, 698, 700, 137: 758, 142: 766},
		// 480
		{81: 34, 34, 92: 34},
		{81: 25, 776, 203: 777, 775},
		{81: 31, 31, 92: 31},
		{81: 28, 28, 92: 770, 168: 774},
		{80: 717, 109: 771, 714},
		// 485
		{79: 772},
		{79: 715, 717, 41, 41, 92: 41, 95: 705, 706, 707, 702, 703, 704, 708, 712, 709, 699, 710, 711, 109: 716, 714, 117: 713, 130: 697, 132: 696, 701, 698, 700, 137: 758, 142: 773},
		{81: 27, 27, 92: 27},
		{81: 30, 30, 92: 30},
		{81: 781, 205: 780},
		// 490
		{79: 778},
		{81: 24},
		{79: 715, 717, 41, 95: 705, 706, 707, 702, 703, 704, 708, 712, 709, 699, 710, 711, 109: 716, 714, 117: 713, 130: 697, 132: 696, 701, 698, 700, 137: 758, 142: 779},
		{81: 26},
		{42: 35, 79: 35, 35, 35, 35, 92: 35, 95: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		// 495
		{42: 23, 79: 23, 23, 23, 23, 92: 23, 95: 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23},
		{79: 783},
		{42: 37, 79: 37, 37, 37, 37, 92: 37, 95: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
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
	const yyError = 217

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
			if !lx.tweaks.enableAnonymousStructFields {
				lx.report.Err(lhs.Token.Pos(), "unnamed fields not allowed")
			} else if k := s.kind(); k != Struct && k != Union {
				lx.report.Err(lhs.Token.Pos(), "only unnamed structs and unions are allowed")
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
			toks := decodeTokens(lhs.PPTokenList, nil, false)
			if len(toks) == 0 {
				lhs.Case = 9 // PPUNDEF IDENTIFIER '\n'
				break
			}

			lx.report.ErrTok(toks[0], "extra tokens after #undef argument")
		}
	case 278:
		{
			yyVAL.node = &ControlLine{
				Case:        14,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 281:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 282:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
