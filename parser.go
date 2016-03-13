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
	yyDefault           = 57453
	yyEofCode           = 57344
	ADDASSIGN           = 57346
	ANDAND              = 57347
	ANDASSIGN           = 57348
	ARROW               = 57349
	ASM                 = 57350
	AUTO                = 57351
	BOOL                = 57352
	BREAK               = 57353
	CASE                = 57354
	CAST                = 57355
	CHAR                = 57356
	CHARCONST           = 57357
	COMPLEX             = 57358
	CONST               = 57359
	CONSTANT_EXPRESSION = 1048577
	CONTINUE            = 57360
	DDD                 = 57361
	DEC                 = 57362
	DEFAULT             = 57363
	DIVASSIGN           = 57364
	DO                  = 57365
	DOUBLE              = 57366
	ELSE                = 57367
	ENUM                = 57368
	EQ                  = 57369
	EXTERN              = 57370
	FLOAT               = 57371
	FLOATCONST          = 57372
	FOR                 = 57373
	GEQ                 = 57374
	GOTO                = 57375
	IDENTIFIER          = 57376
	IDENTIFIER_LPAREN   = 57377
	IDENTIFIER_NONREPL  = 57378
	IF                  = 57379
	INC                 = 57380
	INLINE              = 57381
	INT                 = 57382
	INTCONST            = 57383
	LEQ                 = 57384
	LONG                = 57385
	LONGCHARCONST       = 57386
	LONGSTRINGLITERAL   = 57387
	LSH                 = 57388
	LSHASSIGN           = 57389
	MODASSIGN           = 57390
	MULASSIGN           = 57391
	NEQ                 = 57392
	NOELSE              = 57393
	ORASSIGN            = 57394
	OROR                = 57395
	PPDEFINE            = 57396
	PPELIF              = 57397
	PPELSE              = 57398
	PPENDIF             = 57399
	PPERROR             = 57400
	PPHASH_NL           = 57401
	PPHEADER_NAME       = 57402
	PPIF                = 57403
	PPIFDEF             = 57404
	PPIFNDEF            = 57405
	PPINCLUDE           = 57406
	PPINCLUDE_NEXT      = 57407
	PPLINE              = 57408
	PPNONDIRECTIVE      = 57409
	PPNUMBER            = 57410
	PPOTHER             = 57411
	PPPASTE             = 57412
	PPPRAGMA            = 57413
	PPUNDEF             = 57414
	PREPROCESSING_FILE  = 1048576
	REGISTER            = 57415
	RESTRICT            = 57416
	RETURN              = 57417
	RSH                 = 57418
	RSHASSIGN           = 57419
	SHORT               = 57420
	SIGNED              = 57421
	SIZEOF              = 57422
	STATIC              = 57423
	STRINGLITERAL       = 57424
	STRUCT              = 57425
	SUBASSIGN           = 57426
	SWITCH              = 57427
	TRANSLATION_UNIT    = 1048578
	TYPEDEF             = 57428
	TYPEDEFNAME         = 57429
	TYPEOF              = 57430
	UNARY               = 57431
	UNION               = 57432
	UNSIGNED            = 57433
	VOID                = 57434
	VOLATILE            = 57435
	WHILE               = 57436
	XORASSIGN           = 57437
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -295
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (286x)
		42:      1,   // '*' (253x)
		57376:   2,   // IDENTIFIER (206x)
		38:      3,   // '&' (201x)
		43:      4,   // '+' (201x)
		45:      5,   // '-' (201x)
		57362:   6,   // DEC (201x)
		57380:   7,   // INC (201x)
		41:      8,   // ')' (183x)
		59:      9,   // ';' (183x)
		44:      10,  // ',' (170x)
		91:      11,  // '[' (152x)
		57424:   12,  // STRINGLITERAL (137x)
		33:      13,  // '!' (133x)
		126:     14,  // '~' (133x)
		57357:   15,  // CHARCONST (133x)
		57372:   16,  // FLOATCONST (133x)
		57383:   17,  // INTCONST (133x)
		57386:   18,  // LONGCHARCONST (133x)
		57387:   19,  // LONGSTRINGLITERAL (133x)
		57422:   20,  // SIZEOF (133x)
		57435:   21,  // VOLATILE (120x)
		57359:   22,  // CONST (119x)
		57416:   23,  // RESTRICT (119x)
		125:     24,  // '}' (110x)
		57352:   25,  // BOOL (109x)
		57356:   26,  // CHAR (109x)
		57358:   27,  // COMPLEX (109x)
		57366:   28,  // DOUBLE (109x)
		57368:   29,  // ENUM (109x)
		57371:   30,  // FLOAT (109x)
		57382:   31,  // INT (109x)
		57385:   32,  // LONG (109x)
		57420:   33,  // SHORT (109x)
		57421:   34,  // SIGNED (109x)
		57425:   35,  // STRUCT (109x)
		57429:   36,  // TYPEDEFNAME (109x)
		57430:   37,  // TYPEOF (109x)
		57432:   38,  // UNION (109x)
		57433:   39,  // UNSIGNED (109x)
		57434:   40,  // VOID (109x)
		58:      41,  // ':' (104x)
		57423:   42,  // STATIC (102x)
		57344:   43,  // $end (98x)
		57351:   44,  // AUTO (96x)
		57370:   45,  // EXTERN (96x)
		57381:   46,  // INLINE (96x)
		57415:   47,  // REGISTER (96x)
		57428:   48,  // TYPEDEF (96x)
		61:      49,  // '=' (88x)
		57491:   50,  // Expression (84x)
		93:      51,  // ']' (81x)
		46:      52,  // '.' (77x)
		123:     53,  // '{' (75x)
		37:      54,  // '%' (69x)
		47:      55,  // '/' (69x)
		60:      56,  // '<' (69x)
		62:      57,  // '>' (69x)
		63:      58,  // '?' (69x)
		94:      59,  // '^' (69x)
		124:     60,  // '|' (69x)
		57346:   61,  // ADDASSIGN (69x)
		57347:   62,  // ANDAND (69x)
		57348:   63,  // ANDASSIGN (69x)
		57349:   64,  // ARROW (69x)
		57364:   65,  // DIVASSIGN (69x)
		57369:   66,  // EQ (69x)
		57374:   67,  // GEQ (69x)
		57384:   68,  // LEQ (69x)
		57388:   69,  // LSH (69x)
		57389:   70,  // LSHASSIGN (69x)
		57390:   71,  // MODASSIGN (69x)
		57391:   72,  // MULASSIGN (69x)
		57392:   73,  // NEQ (69x)
		57394:   74,  // ORASSIGN (69x)
		57395:   75,  // OROR (69x)
		57418:   76,  // RSH (69x)
		57419:   77,  // RSHASSIGN (69x)
		57426:   78,  // SUBASSIGN (69x)
		57437:   79,  // XORASSIGN (69x)
		10:      80,  // '\n' (60x)
		57411:   81,  // PPOTHER (54x)
		57399:   82,  // PPENDIF (45x)
		57398:   83,  // PPELSE (41x)
		57436:   84,  // WHILE (41x)
		57353:   85,  // BREAK (40x)
		57354:   86,  // CASE (40x)
		57360:   87,  // CONTINUE (40x)
		57363:   88,  // DEFAULT (40x)
		57365:   89,  // DO (40x)
		57373:   90,  // FOR (40x)
		57375:   91,  // GOTO (40x)
		57379:   92,  // IF (40x)
		57397:   93,  // PPELIF (40x)
		57417:   94,  // RETURN (40x)
		57427:   95,  // SWITCH (40x)
		57396:   96,  // PPDEFINE (36x)
		57400:   97,  // PPERROR (36x)
		57401:   98,  // PPHASH_NL (36x)
		57403:   99,  // PPIF (36x)
		57404:   100, // PPIFDEF (36x)
		57405:   101, // PPIFNDEF (36x)
		57406:   102, // PPINCLUDE (36x)
		57407:   103, // PPINCLUDE_NEXT (36x)
		57408:   104, // PPLINE (36x)
		57409:   105, // PPNONDIRECTIVE (36x)
		57413:   106, // PPPRAGMA (36x)
		57414:   107, // PPUNDEF (36x)
		57541:   108, // TypeQualifier (28x)
		57492:   109, // ExpressionList (26x)
		57515:   110, // PPTokenList (23x)
		57517:   111, // PPTokens (23x)
		57367:   112, // ELSE (22x)
		57487:   113, // EnumSpecifier (20x)
		57536:   114, // StructOrUnion (20x)
		57537:   115, // StructOrUnionSpecifier (20x)
		57544:   116, // TypeSpecifier (20x)
		57493:   117, // ExpressionListOpt (18x)
		57516:   118, // PPTokenListOpt (16x)
		57470:   119, // DeclarationSpecifiers (15x)
		57498:   120, // FunctionSpecifier (15x)
		57531:   121, // StorageClassSpecifier (15x)
		57464:   122, // CompoundStatement (13x)
		57350:   123, // ASM (12x)
		57495:   124, // ExpressionStatement (12x)
		57512:   125, // IterationStatement (12x)
		57513:   126, // JumpStatement (12x)
		57514:   127, // LabeledStatement (12x)
		57526:   128, // SelectionStatement (12x)
		57530:   129, // Statement (12x)
		57522:   130, // Pointer (11x)
		57523:   131, // PointerOpt (10x)
		57466:   132, // ControlLine (8x)
		57472:   133, // Declarator (8x)
		57501:   134, // GroupPart (8x)
		57505:   135, // IfGroup (8x)
		57506:   136, // IfSection (8x)
		57538:   137, // TextLine (8x)
		57467:   138, // Declaration (7x)
		57499:   139, // GroupList (6x)
		57525:   140, // ReplacementList (6x)
		57447:   141, // $@4 (5x)
		57465:   142, // ConstantExpression (5x)
		57361:   143, // DDD (5x)
		57500:   144, // GroupListOpt (5x)
		57527:   145, // SpecifierQualifierList (5x)
		57542:   146, // TypeQualifierList (5x)
		57454:   147, // AbstractDeclarator (4x)
		57471:   148, // DeclarationSpecifiersOpt (4x)
		57476:   149, // Designator (4x)
		57518:   150, // ParameterDeclaration (4x)
		57543:   151, // TypeQualifierListOpt (4x)
		57439:   152, // $@10 (3x)
		57463:   153, // CommaOpt (3x)
		57474:   154, // Designation (3x)
		57475:   155, // DesignationOpt (3x)
		57477:   156, // DesignatorList (3x)
		57494:   157, // ExpressionOpt (3x)
		57507:   158, // InitDeclarator (3x)
		57510:   159, // Initializer (3x)
		57519:   160, // ParameterList (3x)
		57520:   161, // ParameterTypeList (3x)
		57540:   162, // TypeName (3x)
		57440:   163, // $@11 (2x)
		57448:   164, // $@5 (2x)
		57455:   165, // AbstractDeclaratorOpt (2x)
		57459:   166, // BasicAsmStatement (2x)
		57460:   167, // BlockItem (2x)
		57473:   168, // DeclaratorOpt (2x)
		57478:   169, // DirectAbstractDeclarator (2x)
		57479:   170, // DirectAbstractDeclaratorOpt (2x)
		57480:   171, // DirectDeclarator (2x)
		57481:   172, // ElifGroup (2x)
		57488:   173, // EnumerationConstant (2x)
		57489:   174, // Enumerator (2x)
		57496:   175, // ExternalDeclaration (2x)
		57497:   176, // FunctionDefinition (2x)
		57502:   177, // IdentifierList (2x)
		57503:   178, // IdentifierListOpt (2x)
		57504:   179, // IdentifierOpt (2x)
		57508:   180, // InitDeclaratorList (2x)
		57509:   181, // InitDeclaratorListOpt (2x)
		57511:   182, // InitializerList (2x)
		57521:   183, // ParameterTypeListOpt (2x)
		57528:   184, // SpecifierQualifierListOpt (2x)
		57532:   185, // StructDeclaration (2x)
		57534:   186, // StructDeclarator (2x)
		57438:   187, // $@1 (1x)
		57441:   188, // $@12 (1x)
		57442:   189, // $@13 (1x)
		57443:   190, // $@14 (1x)
		57444:   191, // $@15 (1x)
		57445:   192, // $@2 (1x)
		57446:   193, // $@3 (1x)
		57449:   194, // $@6 (1x)
		57450:   195, // $@7 (1x)
		57451:   196, // $@8 (1x)
		57452:   197, // $@9 (1x)
		57456:   198, // ArgumentExpressionList (1x)
		57457:   199, // ArgumentExpressionListOpt (1x)
		57458:   200, // AssemblerInstructions (1x)
		57461:   201, // BlockItemList (1x)
		57462:   202, // BlockItemListOpt (1x)
		1048577: 203, // CONSTANT_EXPRESSION (1x)
		57468:   204, // DeclarationList (1x)
		57469:   205, // DeclarationListOpt (1x)
		57482:   206, // ElifGroupList (1x)
		57483:   207, // ElifGroupListOpt (1x)
		57484:   208, // ElseGroup (1x)
		57485:   209, // ElseGroupOpt (1x)
		57486:   210, // EndifLine (1x)
		57490:   211, // EnumeratorList (1x)
		57377:   212, // IDENTIFIER_LPAREN (1x)
		1048576: 213, // PREPROCESSING_FILE (1x)
		57524:   214, // PreprocessingFile (1x)
		57529:   215, // Start (1x)
		57533:   216, // StructDeclarationList (1x)
		57535:   217, // StructDeclaratorList (1x)
		1048578: 218, // TRANSLATION_UNIT (1x)
		57539:   219, // TranslationUnit (1x)
		57545:   220, // VolatileOpt (1x)
		57453:   221, // $default (0x)
		57355:   222, // CAST (0x)
		57345:   223, // error (0x)
		57378:   224, // IDENTIFIER_NONREPL (0x)
		57393:   225, // NOELSE (0x)
		57402:   226, // PPHEADER_NAME (0x)
		57410:   227, // PPNUMBER (0x)
		57412:   228, // PPPASTE (0x)
		57431:   229, // UNARY (0x)
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
		"STRINGLITERAL",
		"'!'",
		"'~'",
		"CHARCONST",
		"FLOATCONST",
		"INTCONST",
		"LONGCHARCONST",
		"LONGSTRINGLITERAL",
		"SIZEOF",
		"VOLATILE",
		"CONST",
		"RESTRICT",
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
		"TYPEOF",
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
		"ASM",
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
		"$@10",
		"CommaOpt",
		"Designation",
		"DesignationOpt",
		"DesignatorList",
		"ExpressionOpt",
		"InitDeclarator",
		"Initializer",
		"ParameterList",
		"ParameterTypeList",
		"TypeName",
		"$@11",
		"$@5",
		"AbstractDeclaratorOpt",
		"BasicAsmStatement",
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
		"$@1",
		"$@12",
		"$@13",
		"$@14",
		"$@15",
		"$@2",
		"$@3",
		"$@6",
		"$@7",
		"$@8",
		"$@9",
		"ArgumentExpressionList",
		"ArgumentExpressionListOpt",
		"AssemblerInstructions",
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
		"VolatileOpt",
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
		1:   {187, 0},
		2:   {215, 3},
		3:   {192, 0},
		4:   {215, 3},
		5:   {193, 0},
		6:   {215, 3},
		7:   {173, 1},
		8:   {198, 1},
		9:   {198, 3},
		10:  {199, 0},
		11:  {199, 1},
		12:  {50, 1},
		13:  {50, 1},
		14:  {50, 1},
		15:  {50, 1},
		16:  {50, 1},
		17:  {50, 1},
		18:  {50, 1},
		19:  {50, 3},
		20:  {50, 4},
		21:  {50, 4},
		22:  {50, 3},
		23:  {50, 3},
		24:  {50, 2},
		25:  {50, 2},
		26:  {50, 7},
		27:  {50, 2},
		28:  {50, 2},
		29:  {50, 2},
		30:  {50, 2},
		31:  {50, 2},
		32:  {50, 2},
		33:  {50, 2},
		34:  {50, 2},
		35:  {50, 2},
		36:  {50, 4},
		37:  {50, 4},
		38:  {50, 3},
		39:  {50, 3},
		40:  {50, 3},
		41:  {50, 3},
		42:  {50, 3},
		43:  {50, 3},
		44:  {50, 3},
		45:  {50, 3},
		46:  {50, 3},
		47:  {50, 3},
		48:  {50, 3},
		49:  {50, 3},
		50:  {50, 3},
		51:  {50, 3},
		52:  {50, 3},
		53:  {50, 3},
		54:  {50, 3},
		55:  {50, 3},
		56:  {50, 5},
		57:  {50, 3},
		58:  {50, 3},
		59:  {50, 3},
		60:  {50, 3},
		61:  {50, 3},
		62:  {50, 3},
		63:  {50, 3},
		64:  {50, 3},
		65:  {50, 3},
		66:  {50, 3},
		67:  {50, 3},
		68:  {157, 0},
		69:  {157, 1},
		70:  {109, 1},
		71:  {109, 3},
		72:  {117, 0},
		73:  {117, 1},
		74:  {141, 0},
		75:  {142, 2},
		76:  {138, 3},
		77:  {119, 2},
		78:  {119, 2},
		79:  {119, 2},
		80:  {119, 2},
		81:  {148, 0},
		82:  {148, 1},
		83:  {180, 1},
		84:  {180, 3},
		85:  {181, 0},
		86:  {181, 1},
		87:  {158, 1},
		88:  {164, 0},
		89:  {158, 4},
		90:  {121, 1},
		91:  {121, 1},
		92:  {121, 1},
		93:  {121, 1},
		94:  {121, 1},
		95:  {116, 1},
		96:  {116, 1},
		97:  {116, 1},
		98:  {116, 1},
		99:  {116, 1},
		100: {116, 1},
		101: {116, 1},
		102: {116, 1},
		103: {116, 1},
		104: {116, 1},
		105: {116, 1},
		106: {116, 1},
		107: {116, 1},
		108: {116, 1},
		109: {116, 4},
		110: {116, 4},
		111: {194, 0},
		112: {195, 0},
		113: {115, 7},
		114: {115, 2},
		115: {114, 1},
		116: {114, 1},
		117: {216, 1},
		118: {216, 2},
		119: {185, 3},
		120: {185, 2},
		121: {145, 2},
		122: {145, 2},
		123: {184, 0},
		124: {184, 1},
		125: {217, 1},
		126: {217, 3},
		127: {186, 1},
		128: {186, 3},
		129: {153, 0},
		130: {153, 1},
		131: {196, 0},
		132: {113, 7},
		133: {113, 2},
		134: {211, 1},
		135: {211, 3},
		136: {174, 1},
		137: {174, 3},
		138: {108, 1},
		139: {108, 1},
		140: {108, 1},
		141: {120, 1},
		142: {133, 2},
		143: {168, 0},
		144: {168, 1},
		145: {171, 1},
		146: {171, 3},
		147: {171, 5},
		148: {171, 6},
		149: {171, 6},
		150: {171, 5},
		151: {197, 0},
		152: {171, 5},
		153: {171, 4},
		154: {130, 2},
		155: {130, 3},
		156: {131, 0},
		157: {131, 1},
		158: {146, 1},
		159: {146, 2},
		160: {151, 0},
		161: {151, 1},
		162: {161, 1},
		163: {161, 3},
		164: {183, 0},
		165: {183, 1},
		166: {160, 1},
		167: {160, 3},
		168: {150, 2},
		169: {150, 2},
		170: {177, 1},
		171: {177, 3},
		172: {178, 0},
		173: {178, 1},
		174: {179, 0},
		175: {179, 1},
		176: {152, 0},
		177: {162, 3},
		178: {147, 1},
		179: {147, 2},
		180: {165, 0},
		181: {165, 1},
		182: {169, 3},
		183: {169, 4},
		184: {169, 5},
		185: {169, 6},
		186: {169, 6},
		187: {169, 4},
		188: {163, 0},
		189: {169, 4},
		190: {188, 0},
		191: {169, 5},
		192: {170, 0},
		193: {170, 1},
		194: {159, 1},
		195: {159, 4},
		196: {182, 2},
		197: {182, 4},
		198: {154, 2},
		199: {155, 0},
		200: {155, 1},
		201: {156, 1},
		202: {156, 2},
		203: {149, 3},
		204: {149, 2},
		205: {129, 1},
		206: {129, 1},
		207: {129, 1},
		208: {129, 1},
		209: {129, 1},
		210: {129, 1},
		211: {127, 3},
		212: {127, 4},
		213: {127, 3},
		214: {189, 0},
		215: {122, 4},
		216: {201, 1},
		217: {201, 2},
		218: {202, 0},
		219: {202, 1},
		220: {167, 1},
		221: {167, 1},
		222: {124, 2},
		223: {128, 5},
		224: {128, 7},
		225: {128, 5},
		226: {125, 5},
		227: {125, 7},
		228: {125, 9},
		229: {125, 8},
		230: {126, 3},
		231: {126, 2},
		232: {126, 2},
		233: {126, 3},
		234: {219, 1},
		235: {219, 2},
		236: {175, 1},
		237: {175, 1},
		238: {175, 1},
		239: {190, 0},
		240: {176, 5},
		241: {204, 1},
		242: {204, 2},
		243: {205, 0},
		244: {191, 0},
		245: {205, 2},
		246: {200, 1},
		247: {200, 2},
		248: {166, 5},
		249: {220, 0},
		250: {220, 1},
		251: {214, 1},
		252: {139, 1},
		253: {139, 2},
		254: {144, 0},
		255: {144, 1},
		256: {134, 1},
		257: {134, 1},
		258: {134, 3},
		259: {134, 1},
		260: {136, 4},
		261: {135, 4},
		262: {135, 4},
		263: {135, 4},
		264: {206, 1},
		265: {206, 2},
		266: {207, 0},
		267: {207, 1},
		268: {172, 4},
		269: {208, 3},
		270: {209, 0},
		271: {209, 1},
		272: {210, 1},
		273: {132, 3},
		274: {132, 5},
		275: {132, 7},
		276: {132, 5},
		277: {132, 2},
		278: {132, 1},
		279: {132, 3},
		280: {132, 3},
		281: {132, 2},
		282: {132, 3},
		283: {132, 6},
		284: {132, 8},
		285: {132, 2},
		286: {132, 4},
		287: {132, 3},
		288: {137, 1},
		289: {140, 1},
		290: {110, 1},
		291: {118, 1},
		292: {118, 2},
		293: {111, 1},
		294: {111, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 43}:   "invalid empty input",
		yyXError{507, -1}: "expected #endif",
		yyXError{509, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{422, -1}: "expected $end",
		yyXError{424, -1}: "expected $end",
		yyXError{31, -1}:  "expected '('",
		yyXError{45, -1}:  "expected '('",
		yyXError{46, -1}:  "expected '('",
		yyXError{356, -1}: "expected '('",
		yyXError{357, -1}: "expected '('",
		yyXError{358, -1}: "expected '('",
		yyXError{360, -1}: "expected '('",
		yyXError{386, -1}: "expected '('",
		yyXError{158, -1}: "expected ')'",
		yyXError{165, -1}: "expected ')'",
		yyXError{172, -1}: "expected ')'",
		yyXError{198, -1}: "expected ')'",
		yyXError{201, -1}: "expected ')'",
		yyXError{204, -1}: "expected ')'",
		yyXError{212, -1}: "expected ')'",
		yyXError{217, -1}: "expected ')'",
		yyXError{223, -1}: "expected ')'",
		yyXError{239, -1}: "expected ')'",
		yyXError{244, -1}: "expected ')'",
		yyXError{285, -1}: "expected ')'",
		yyXError{316, -1}: "expected ')'",
		yyXError{376, -1}: "expected ')'",
		yyXError{382, -1}: "expected ')'",
		yyXError{467, -1}: "expected ')'",
		yyXError{468, -1}: "expected ')'",
		yyXError{476, -1}: "expected ')'",
		yyXError{479, -1}: "expected ')'",
		yyXError{482, -1}: "expected ')'",
		yyXError{303, -1}: "expected ':'",
		yyXError{349, -1}: "expected ':'",
		yyXError{410, -1}: "expected ':'",
		yyXError{324, -1}: "expected ';'",
		yyXError{355, -1}: "expected ';'",
		yyXError{362, -1}: "expected ';'",
		yyXError{363, -1}: "expected ';'",
		yyXError{365, -1}: "expected ';'",
		yyXError{369, -1}: "expected ';'",
		yyXError{372, -1}: "expected ';'",
		yyXError{374, -1}: "expected ';'",
		yyXError{380, -1}: "expected ';'",
		yyXError{389, -1}: "expected ';'",
		yyXError{328, -1}: "expected '='",
		yyXError{177, -1}: "expected '['",
		yyXError{446, -1}: "expected '\\n'",
		yyXError{450, -1}: "expected '\\n'",
		yyXError{454, -1}: "expected '\\n'",
		yyXError{457, -1}: "expected '\\n'",
		yyXError{459, -1}: "expected '\\n'",
		yyXError{486, -1}: "expected '\\n'",
		yyXError{491, -1}: "expected '\\n'",
		yyXError{494, -1}: "expected '\\n'",
		yyXError{501, -1}: "expected '\\n'",
		yyXError{506, -1}: "expected '\\n'",
		yyXError{512, -1}: "expected '\\n'",
		yyXError{183, -1}: "expected ']'",
		yyXError{191, -1}: "expected ']'",
		yyXError{235, -1}: "expected ']'",
		yyXError{262, -1}: "expected ']'",
		yyXError{52, -1}:  "expected '{'",
		yyXError{54, -1}:  "expected '{'",
		yyXError{291, -1}: "expected '{'",
		yyXError{293, -1}: "expected '{'",
		yyXError{271, -1}: "expected '}'",
		yyXError{275, -1}: "expected '}'",
		yyXError{288, -1}: "expected '}'",
		yyXError{350, -1}: "expected '}'",
		yyXError{57, -1}:  "expected CommaOpt or one of [',', '}']",
		yyXError{254, -1}: "expected CommaOpt or one of [',', '}']",
		yyXError{269, -1}: "expected CommaOpt or one of [',', '}']",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{211, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{176, -1}: "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{47, -1}:  "expected assembler instructions or string literal",
		yyXError{352, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{329, -1}: "expected compound statement or '{'",
		yyXError{336, -1}: "expected compound statement or '{'",
		yyXError{327, -1}: "expected compound statement or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{60, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{259, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{307, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{348, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{421, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{330, -1}: "expected declaration list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{333, -1}: "expected declaration or one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{371, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{306, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{203, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{256, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{206, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{173, -1}: "expected direct abstract declarator or one of ['(', '[']",
		yyXError{304, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{499, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{505, -1}: "expected endif line or #endif",
		yyXError{431, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{497, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{55, -1}:  "expected enumerator list or identifier",
		yyXError{287, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{83, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{107, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{387, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{391, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{395, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{399, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{70, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{81, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{251, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		yyXError{180, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{234, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{286, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{61, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{72, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{73, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{74, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{75, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{76, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{77, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{78, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{79, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{80, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{89, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{90, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{91, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{92, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{93, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{94, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{95, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{96, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{108, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{109, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{110, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{111, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{112, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{113, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{114, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{115, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{116, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{117, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{118, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{132, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{133, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{160, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{186, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{192, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{228, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{231, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{184, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{226, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{314, -1}: "expected expression or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{6, -1}:   "expected external declaration or one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{488, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{425, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{85, -1}:  "expected identifier",
		yyXError{86, -1}:  "expected identifier",
		yyXError{220, -1}: "expected identifier",
		yyXError{260, -1}: "expected identifier",
		yyXError{361, -1}: "expected identifier",
		yyXError{433, -1}: "expected identifier",
		yyXError{434, -1}: "expected identifier",
		yyXError{441, -1}: "expected identifier",
		yyXError{463, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{417, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{253, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{267, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{255, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{273, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{415, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{270, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{63, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{64, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{65, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{66, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{67, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{68, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{69, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{82, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{87, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{88, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{119, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{120, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{121, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{122, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{123, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{124, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{125, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{126, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{127, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{128, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{129, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
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
		yyXError{146, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{147, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{148, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{149, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{150, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{151, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{152, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{153, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{154, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{155, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{159, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{163, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{196, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{252, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{276, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{277, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{278, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{279, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{280, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{281, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{282, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{283, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{284, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{71, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{130, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{134, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{156, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{161, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{315, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{340, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{266, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{179, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{187, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{193, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{229, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{232, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{426, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{427, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{428, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{430, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{437, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{443, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{445, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{448, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{451, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{453, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{455, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{456, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{458, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{460, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{461, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{464, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{470, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{471, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{473, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{478, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{481, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{484, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{485, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{490, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{510, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{511, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{513, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{489, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{493, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{496, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{498, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{503, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{504, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{407, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{419, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{40, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{41, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{42, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{43, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{51, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{338, -1}: "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{420, -1}: "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{36, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{38, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{181, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{189, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{342, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{343, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{344, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{345, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{346, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{347, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{366, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{367, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{368, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{370, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{378, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{384, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{390, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{394, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{398, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{402, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{404, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{405, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{409, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{412, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{414, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{351, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{353, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{354, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{406, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{257, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{264, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{53, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{292, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{17, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{18, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{19, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{20, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{21, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{22, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{23, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{24, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{25, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{26, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{27, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{28, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{29, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{30, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{289, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{312, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{317, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{318, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{12, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{39, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{319, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{320, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{321, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{322, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{323, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{248, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{249, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{250, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{209, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{210, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{213, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{222, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{224, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{230, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{233, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{236, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{237, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{171, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{247, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{175, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{188, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{190, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{194, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{195, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{197, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{205, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{241, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{245, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{305, -1}: "expected one of ['(', identifier]",
		yyXError{341, -1}: "expected one of [')', ',', ';']",
		yyXError{465, -1}: "expected one of [')', ',', ...]",
		yyXError{475, -1}: "expected one of [')', ',', ...]",
		yyXError{157, -1}: "expected one of [')', ',']",
		yyXError{164, -1}: "expected one of [')', ',']",
		yyXError{174, -1}: "expected one of [')', ',']",
		yyXError{200, -1}: "expected one of [')', ',']",
		yyXError{202, -1}: "expected one of [')', ',']",
		yyXError{207, -1}: "expected one of [')', ',']",
		yyXError{208, -1}: "expected one of [')', ',']",
		yyXError{218, -1}: "expected one of [')', ',']",
		yyXError{219, -1}: "expected one of [')', ',']",
		yyXError{221, -1}: "expected one of [')', ',']",
		yyXError{240, -1}: "expected one of [')', ',']",
		yyXError{388, -1}: "expected one of [')', ',']",
		yyXError{392, -1}: "expected one of [')', ',']",
		yyXError{396, -1}: "expected one of [')', ',']",
		yyXError{400, -1}: "expected one of [')', ',']",
		yyXError{466, -1}: "expected one of [')', ',']",
		yyXError{48, -1}:  "expected one of [')', string literal]",
		yyXError{49, -1}:  "expected one of [')', string literal]",
		yyXError{50, -1}:  "expected one of [')', string literal]",
		yyXError{302, -1}: "expected one of [',', ':', ';']",
		yyXError{131, -1}: "expected one of [',', ':']",
		yyXError{335, -1}: "expected one of [',', ';', '=']",
		yyXError{272, -1}: "expected one of [',', ';', '}']",
		yyXError{299, -1}: "expected one of [',', ';']",
		yyXError{301, -1}: "expected one of [',', ';']",
		yyXError{308, -1}: "expected one of [',', ';']",
		yyXError{311, -1}: "expected one of [',', ';']",
		yyXError{325, -1}: "expected one of [',', ';']",
		yyXError{326, -1}: "expected one of [',', ';']",
		yyXError{416, -1}: "expected one of [',', ';']",
		yyXError{418, -1}: "expected one of [',', ';']",
		yyXError{56, -1}:  "expected one of [',', '=', '}']",
		yyXError{59, -1}:  "expected one of [',', '=', '}']",
		yyXError{162, -1}: "expected one of [',', ']']",
		yyXError{58, -1}:  "expected one of [',', '}']",
		yyXError{62, -1}:  "expected one of [',', '}']",
		yyXError{268, -1}: "expected one of [',', '}']",
		yyXError{274, -1}: "expected one of [',', '}']",
		yyXError{290, -1}: "expected one of [',', '}']",
		yyXError{258, -1}: "expected one of ['.', '=', '[']",
		yyXError{261, -1}: "expected one of ['.', '=', '[']",
		yyXError{263, -1}: "expected one of ['.', '=', '[']",
		yyXError{265, -1}: "expected one of ['.', '=', '[']",
		yyXError{435, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		yyXError{444, -1}: "expected one of ['\\n', ppother]",
		yyXError{447, -1}: "expected one of ['\\n', ppother]",
		yyXError{449, -1}: "expected one of ['\\n', ppother]",
		yyXError{332, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{334, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{33, -1}:  "expected one of ['{', identifier]",
		yyXError{34, -1}:  "expected one of ['{', identifier]",
		yyXError{297, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{300, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{309, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{313, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{474, -1}: "expected one of [..., identifier]",
		yyXError{169, -1}: "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{84, -1}:  "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{337, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{339, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{8, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{375, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{381, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{364, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{373, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{379, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{225, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{214, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{178, -1}: "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{182, -1}: "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{487, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{492, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{495, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{502, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{508, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{215, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{32, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{35, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{331, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{199, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{242, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{243, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{167, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{168, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{436, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{440, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{170, -1}: "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{44, -1}:  "expected optional volatile or one of ['(', volatile]",
		yyXError{238, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{216, -1}: "expected parameter type list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{246, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{423, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{462, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{469, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{472, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{477, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{480, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{483, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{166, -1}: "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{359, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{377, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{383, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{393, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{397, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{401, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{403, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{408, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{411, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{413, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{294, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{295, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{296, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{298, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		yyXError{310, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{452, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{429, -1}: "expected token list or ppother",
		yyXError{432, -1}: "expected token list or ppother",
		yyXError{438, -1}: "expected token list or ppother",
		yyXError{439, -1}: "expected token list or ppother",
		yyXError{442, -1}: "expected token list or ppother",
		yyXError{500, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [_Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{5, -1}:   "expected translation unit or one of [_Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{185, -1}: "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{227, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{385, -1}: "expected while",
		yyXError{3, 43}:   "unexpected EOF",
		yyXError{2, 43}:   "unexpected EOF",
		yyXError{4, 43}:   "unexpected EOF",
	}

	yyParseTab = [514][]uint16{
		// 0
		{203: 298, 213: 297, 215: 296, 218: 299},
		{43: 295},
		{80: 294, 294, 96: 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 187: 718},
		{292, 292, 292, 292, 292, 292, 292, 292, 12: 292, 292, 292, 292, 292, 292, 292, 292, 292, 192: 716},
		{21: 290, 290, 290, 25: 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 42: 290, 44: 290, 290, 290, 290, 290, 123: 290, 193: 300},
		// 5
		{21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 108: 305, 113: 324, 327, 323, 304, 119: 302, 306, 303, 123: 339, 138: 337, 166: 338, 175: 335, 336, 219: 301},
		{21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 289, 310, 308, 334, 311, 307, 108: 305, 113: 324, 327, 323, 304, 119: 302, 306, 303, 123: 339, 138: 337, 166: 338, 175: 715, 336},
		{139, 465, 139, 9: 210, 130: 600, 599, 133: 622, 158: 620, 180: 621, 619},
		{214, 214, 214, 8: 214, 214, 214, 214, 21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 108: 305, 113: 324, 327, 323, 304, 119: 615, 306, 303, 148: 618},
		{214, 214, 214, 8: 214, 214, 214, 214, 21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 108: 305, 113: 324, 327, 323, 304, 119: 615, 306, 303, 148: 617},
		// 10
		{214, 214, 214, 8: 214, 214, 214, 214, 21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 108: 305, 113: 324, 327, 323, 304, 119: 615, 306, 303, 148: 616},
		{214, 214, 214, 8: 214, 214, 214, 214, 21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 108: 305, 113: 324, 327, 323, 304, 119: 615, 306, 303, 148: 614},
		{205, 205, 205, 8: 205, 205, 205, 205, 21: 205, 205, 205, 25: 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 42: 205, 44: 205, 205, 205, 205, 205},
		{204, 204, 204, 8: 204, 204, 204, 204, 21: 204, 204, 204, 25: 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 42: 204, 44: 204, 204, 204, 204, 204},
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 203, 203, 203, 25: 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 42: 203, 44: 203, 203, 203, 203, 203},
		// 15
		{202, 202, 202, 8: 202, 202, 202, 202, 21: 202, 202, 202, 25: 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 42: 202, 44: 202, 202, 202, 202, 202},
		{201, 201, 201, 8: 201, 201, 201, 201, 21: 201, 201, 201, 25: 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 42: 201, 44: 201, 201, 201, 201, 201},
		{200, 200, 200, 8: 200, 200, 200, 200, 21: 200, 200, 200, 25: 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 44: 200, 200, 200, 200, 200},
		{199, 199, 199, 8: 199, 199, 199, 199, 21: 199, 199, 199, 25: 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 44: 199, 199, 199, 199, 199},
		{198, 198, 198, 8: 198, 198, 198, 198, 21: 198, 198, 198, 25: 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 44: 198, 198, 198, 198, 198},
		// 20
		{197, 197, 197, 8: 197, 197, 197, 197, 21: 197, 197, 197, 25: 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 44: 197, 197, 197, 197, 197},
		{196, 196, 196, 8: 196, 196, 196, 196, 21: 196, 196, 196, 25: 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 44: 196, 196, 196, 196, 196},
		{195, 195, 195, 8: 195, 195, 195, 195, 21: 195, 195, 195, 25: 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 44: 195, 195, 195, 195, 195},
		{194, 194, 194, 8: 194, 194, 194, 194, 21: 194, 194, 194, 25: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 44: 194, 194, 194, 194, 194},
		{193, 193, 193, 8: 193, 193, 193, 193, 21: 193, 193, 193, 25: 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 44: 193, 193, 193, 193, 193},
		// 25
		{192, 192, 192, 8: 192, 192, 192, 192, 21: 192, 192, 192, 25: 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 44: 192, 192, 192, 192, 192},
		{191, 191, 191, 8: 191, 191, 191, 191, 21: 191, 191, 191, 25: 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 44: 191, 191, 191, 191, 191},
		{190, 190, 190, 8: 190, 190, 190, 190, 21: 190, 190, 190, 25: 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 44: 190, 190, 190, 190, 190},
		{189, 189, 189, 8: 189, 189, 189, 189, 21: 189, 189, 189, 25: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 44: 189, 189, 189, 189, 189},
		{188, 188, 188, 8: 188, 188, 188, 188, 21: 188, 188, 188, 25: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 44: 188, 188, 188, 188, 188},
		// 30
		{187, 187, 187, 8: 187, 187, 187, 187, 21: 187, 187, 187, 25: 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 44: 187, 187, 187, 187, 187},
		{609},
		{2: 587, 53: 121, 179: 586},
		{2: 180, 53: 180},
		{2: 179, 53: 179},
		// 35
		{2: 348, 53: 121, 179: 347},
		{157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 25: 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 44: 157, 157, 157, 157, 157, 51: 157},
		{156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 25: 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 44: 156, 156, 156, 156, 156, 51: 156},
		{155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 25: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 44: 155, 155, 155, 155, 155, 51: 155},
		{154, 154, 154, 8: 154, 154, 154, 154, 21: 154, 154, 154, 25: 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 42: 154, 44: 154, 154, 154, 154, 154},
		// 40
		{21: 61, 61, 61, 25: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 42: 61, 61, 61, 61, 61, 61, 61, 123: 61},
		{21: 59, 59, 59, 25: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 42: 59, 59, 59, 59, 59, 59, 59, 123: 59},
		{21: 58, 58, 58, 25: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 42: 58, 58, 58, 58, 58, 58, 58, 123: 58},
		{21: 57, 57, 57, 25: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 42: 57, 57, 57, 57, 57, 57, 57, 123: 57},
		{46, 21: 341, 220: 340},
		// 45
		{342},
		{45},
		{12: 343, 200: 344},
		{8: 49, 12: 49},
		{8: 346, 12: 345},
		// 50
		{8: 48, 12: 48},
		{21: 47, 47, 47, 25: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 42: 47, 47, 47, 47, 47, 47, 47, 123: 47},
		{53: 164, 196: 349},
		{162, 162, 162, 8: 162, 162, 162, 162, 21: 162, 162, 162, 25: 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 44: 162, 162, 162, 162, 162, 53: 120},
		{53: 350},
		// 55
		{2: 351, 173: 354, 353, 211: 352},
		{10: 288, 24: 288, 49: 288},
		{10: 582, 24: 166, 153: 583},
		{10: 161, 24: 161},
		{10: 159, 24: 159, 49: 355},
		// 60
		{221, 221, 221, 221, 221, 221, 221, 221, 12: 221, 221, 221, 221, 221, 221, 221, 221, 221, 141: 356, 357},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 366},
		{10: 158, 24: 158},
		{283, 283, 3: 283, 283, 283, 283, 283, 283, 283, 283, 283, 24: 283, 41: 283, 43: 283, 49: 283, 51: 283, 283, 54: 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283},
		{282, 282, 3: 282, 282, 282, 282, 282, 282, 282, 282, 282, 24: 282, 41: 282, 43: 282, 49: 282, 51: 282, 282, 54: 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282},
		// 65
		{281, 281, 3: 281, 281, 281, 281, 281, 281, 281, 281, 281, 24: 281, 41: 281, 43: 281, 49: 281, 51: 281, 281, 54: 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281},
		{280, 280, 3: 280, 280, 280, 280, 280, 280, 280, 280, 280, 24: 280, 41: 280, 43: 280, 49: 280, 51: 280, 280, 54: 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280},
		{279, 279, 3: 279, 279, 279, 279, 279, 279, 279, 279, 279, 24: 279, 41: 279, 43: 279, 49: 279, 51: 279, 279, 54: 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279},
		{278, 278, 3: 278, 278, 278, 278, 278, 278, 278, 278, 278, 24: 278, 41: 278, 43: 278, 49: 278, 51: 278, 278, 54: 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278},
		{277, 277, 3: 277, 277, 277, 277, 277, 277, 277, 277, 277, 24: 277, 41: 277, 43: 277, 49: 277, 51: 277, 277, 54: 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277},
		// 70
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 119, 119, 119, 25: 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 50: 425, 109: 459, 152: 461, 162: 580},
		{379, 384, 3: 397, 387, 388, 383, 382, 9: 220, 220, 378, 24: 220, 41: 220, 43: 220, 49: 403, 51: 220, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 579},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 578},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 577},
		// 75
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 491},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 576},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 575},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 574},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 573},
		// 80
		{376, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 377},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 119, 119, 119, 25: 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 50: 425, 109: 459, 152: 461, 162: 460},
		{379, 260, 3: 260, 260, 260, 383, 382, 260, 260, 260, 378, 24: 260, 41: 260, 43: 260, 49: 260, 51: 260, 380, 54: 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 381, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 109: 457},
		{365, 370, 358, 369, 371, 372, 368, 367, 285, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 451, 198: 452, 453},
		// 85
		{2: 450},
		{2: 449},
		{271, 271, 3: 271, 271, 271, 271, 271, 271, 271, 271, 271, 24: 271, 41: 271, 43: 271, 49: 271, 51: 271, 271, 54: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{270, 270, 3: 270, 270, 270, 270, 270, 270, 270, 270, 270, 24: 270, 41: 270, 43: 270, 49: 270, 51: 270, 270, 54: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 448},
		// 90
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 447},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 446},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 445},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 444},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 443},
		// 95
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 442},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 441},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 440},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 439},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 438},
		// 100
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 437},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 436},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 435},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 434},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 433},
		// 105
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 432},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 431},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 109: 426},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 424},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 423},
		// 110
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 422},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 421},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 420},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 419},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 418},
		// 115
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 417},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 416},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 415},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 414},
		{379, 384, 3: 397, 387, 388, 383, 382, 228, 228, 228, 378, 24: 228, 41: 228, 43: 228, 49: 403, 51: 228, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		// 120
		{379, 384, 3: 397, 387, 388, 383, 382, 229, 229, 229, 378, 24: 229, 41: 229, 43: 229, 49: 403, 51: 229, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{379, 384, 3: 397, 387, 388, 383, 382, 230, 230, 230, 378, 24: 230, 41: 230, 43: 230, 49: 403, 51: 230, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{379, 384, 3: 397, 387, 388, 383, 382, 231, 231, 231, 378, 24: 231, 41: 231, 43: 231, 49: 403, 51: 231, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{379, 384, 3: 397, 387, 388, 383, 382, 232, 232, 232, 378, 24: 232, 41: 232, 43: 232, 49: 403, 51: 232, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{379, 384, 3: 397, 387, 388, 383, 382, 233, 233, 233, 378, 24: 233, 41: 233, 43: 233, 49: 403, 51: 233, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		// 125
		{379, 384, 3: 397, 387, 388, 383, 382, 234, 234, 234, 378, 24: 234, 41: 234, 43: 234, 49: 403, 51: 234, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{379, 384, 3: 397, 387, 388, 383, 382, 235, 235, 235, 378, 24: 235, 41: 235, 43: 235, 49: 403, 51: 235, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{379, 384, 3: 397, 387, 388, 383, 382, 236, 236, 236, 378, 24: 236, 41: 236, 43: 236, 49: 403, 51: 236, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{379, 384, 3: 397, 387, 388, 383, 382, 237, 237, 237, 378, 24: 237, 41: 237, 43: 237, 49: 403, 51: 237, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{379, 384, 3: 397, 387, 388, 383, 382, 238, 238, 238, 378, 24: 238, 41: 238, 43: 238, 49: 403, 51: 238, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		// 130
		{379, 384, 3: 397, 387, 388, 383, 382, 225, 225, 225, 378, 41: 225, 49: 403, 51: 225, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{10: 428, 41: 427},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 430},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 429},
		{379, 384, 3: 397, 387, 388, 383, 382, 224, 224, 224, 378, 41: 224, 49: 403, 51: 224, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		// 135
		{379, 384, 3: 397, 387, 388, 383, 382, 239, 239, 239, 378, 24: 239, 41: 239, 43: 239, 49: 239, 51: 239, 380, 54: 386, 385, 391, 392, 402, 398, 399, 239, 400, 239, 381, 239, 395, 394, 393, 389, 239, 239, 239, 396, 239, 401, 390, 239, 239, 239},
		{379, 384, 3: 397, 387, 388, 383, 382, 240, 240, 240, 378, 24: 240, 41: 240, 43: 240, 49: 240, 51: 240, 380, 54: 386, 385, 391, 392, 240, 398, 399, 240, 400, 240, 381, 240, 395, 394, 393, 389, 240, 240, 240, 396, 240, 240, 390, 240, 240, 240},
		{379, 384, 3: 397, 387, 388, 383, 382, 241, 241, 241, 378, 24: 241, 41: 241, 43: 241, 49: 241, 51: 241, 380, 54: 386, 385, 391, 392, 241, 398, 399, 241, 241, 241, 381, 241, 395, 394, 393, 389, 241, 241, 241, 396, 241, 241, 390, 241, 241, 241},
		{379, 384, 3: 397, 387, 388, 383, 382, 242, 242, 242, 378, 24: 242, 41: 242, 43: 242, 49: 242, 51: 242, 380, 54: 386, 385, 391, 392, 242, 398, 242, 242, 242, 242, 381, 242, 395, 394, 393, 389, 242, 242, 242, 396, 242, 242, 390, 242, 242, 242},
		{379, 384, 3: 397, 387, 388, 383, 382, 243, 243, 243, 378, 24: 243, 41: 243, 43: 243, 49: 243, 51: 243, 380, 54: 386, 385, 391, 392, 243, 243, 243, 243, 243, 243, 381, 243, 395, 394, 393, 389, 243, 243, 243, 396, 243, 243, 390, 243, 243, 243},
		// 140
		{379, 384, 3: 244, 387, 388, 383, 382, 244, 244, 244, 378, 24: 244, 41: 244, 43: 244, 49: 244, 51: 244, 380, 54: 386, 385, 391, 392, 244, 244, 244, 244, 244, 244, 381, 244, 395, 394, 393, 389, 244, 244, 244, 396, 244, 244, 390, 244, 244, 244},
		{379, 384, 3: 245, 387, 388, 383, 382, 245, 245, 245, 378, 24: 245, 41: 245, 43: 245, 49: 245, 51: 245, 380, 54: 386, 385, 391, 392, 245, 245, 245, 245, 245, 245, 381, 245, 245, 394, 393, 389, 245, 245, 245, 245, 245, 245, 390, 245, 245, 245},
		{379, 384, 3: 246, 387, 388, 383, 382, 246, 246, 246, 378, 24: 246, 41: 246, 43: 246, 49: 246, 51: 246, 380, 54: 386, 385, 391, 392, 246, 246, 246, 246, 246, 246, 381, 246, 246, 394, 393, 389, 246, 246, 246, 246, 246, 246, 390, 246, 246, 246},
		{379, 384, 3: 247, 387, 388, 383, 382, 247, 247, 247, 378, 24: 247, 41: 247, 43: 247, 49: 247, 51: 247, 380, 54: 386, 385, 247, 247, 247, 247, 247, 247, 247, 247, 381, 247, 247, 247, 247, 389, 247, 247, 247, 247, 247, 247, 390, 247, 247, 247},
		{379, 384, 3: 248, 387, 388, 383, 382, 248, 248, 248, 378, 24: 248, 41: 248, 43: 248, 49: 248, 51: 248, 380, 54: 386, 385, 248, 248, 248, 248, 248, 248, 248, 248, 381, 248, 248, 248, 248, 389, 248, 248, 248, 248, 248, 248, 390, 248, 248, 248},
		// 145
		{379, 384, 3: 249, 387, 388, 383, 382, 249, 249, 249, 378, 24: 249, 41: 249, 43: 249, 49: 249, 51: 249, 380, 54: 386, 385, 249, 249, 249, 249, 249, 249, 249, 249, 381, 249, 249, 249, 249, 389, 249, 249, 249, 249, 249, 249, 390, 249, 249, 249},
		{379, 384, 3: 250, 387, 388, 383, 382, 250, 250, 250, 378, 24: 250, 41: 250, 43: 250, 49: 250, 51: 250, 380, 54: 386, 385, 250, 250, 250, 250, 250, 250, 250, 250, 381, 250, 250, 250, 250, 389, 250, 250, 250, 250, 250, 250, 390, 250, 250, 250},
		{379, 384, 3: 251, 387, 388, 383, 382, 251, 251, 251, 378, 24: 251, 41: 251, 43: 251, 49: 251, 51: 251, 380, 54: 386, 385, 251, 251, 251, 251, 251, 251, 251, 251, 381, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251},
		{379, 384, 3: 252, 387, 388, 383, 382, 252, 252, 252, 378, 24: 252, 41: 252, 43: 252, 49: 252, 51: 252, 380, 54: 386, 385, 252, 252, 252, 252, 252, 252, 252, 252, 381, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252},
		{379, 384, 3: 253, 253, 253, 383, 382, 253, 253, 253, 378, 24: 253, 41: 253, 43: 253, 49: 253, 51: 253, 380, 54: 386, 385, 253, 253, 253, 253, 253, 253, 253, 253, 381, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253},
		// 150
		{379, 384, 3: 254, 254, 254, 383, 382, 254, 254, 254, 378, 24: 254, 41: 254, 43: 254, 49: 254, 51: 254, 380, 54: 386, 385, 254, 254, 254, 254, 254, 254, 254, 254, 381, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254},
		{379, 255, 3: 255, 255, 255, 383, 382, 255, 255, 255, 378, 24: 255, 41: 255, 43: 255, 49: 255, 51: 255, 380, 54: 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 381, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{379, 256, 3: 256, 256, 256, 383, 382, 256, 256, 256, 378, 24: 256, 41: 256, 43: 256, 49: 256, 51: 256, 380, 54: 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 381, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256},
		{379, 257, 3: 257, 257, 257, 383, 382, 257, 257, 257, 378, 24: 257, 41: 257, 43: 257, 49: 257, 51: 257, 380, 54: 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 381, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257},
		{272, 272, 3: 272, 272, 272, 272, 272, 272, 272, 272, 272, 24: 272, 41: 272, 43: 272, 49: 272, 51: 272, 272, 54: 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272},
		// 155
		{273, 273, 3: 273, 273, 273, 273, 273, 273, 273, 273, 273, 24: 273, 41: 273, 43: 273, 49: 273, 51: 273, 273, 54: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273},
		{379, 384, 3: 397, 387, 388, 383, 382, 287, 10: 287, 378, 49: 403, 52: 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{8: 284, 10: 455},
		{8: 454},
		{274, 274, 3: 274, 274, 274, 274, 274, 274, 274, 274, 274, 24: 274, 41: 274, 43: 274, 49: 274, 51: 274, 274, 54: 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274},
		// 160
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 456},
		{379, 384, 3: 397, 387, 388, 383, 382, 286, 10: 286, 378, 49: 403, 52: 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{10: 428, 51: 458},
		{275, 275, 3: 275, 275, 275, 275, 275, 275, 275, 275, 275, 24: 275, 41: 275, 43: 275, 49: 275, 51: 275, 275, 54: 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275},
		{8: 572, 10: 428},
		// 165
		{8: 546},
		{21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 108: 463, 113: 324, 327, 323, 462, 145: 464},
		{172, 172, 172, 8: 172, 172, 11: 172, 21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 172, 108: 463, 113: 324, 327, 323, 462, 145: 544, 184: 545},
		{172, 172, 172, 8: 172, 172, 11: 172, 21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 172, 108: 463, 113: 324, 327, 323, 462, 145: 544, 184: 543},
		{139, 465, 8: 115, 11: 139, 130: 466, 468, 147: 469, 165: 467},
		// 170
		{135, 135, 135, 8: 135, 10: 135, 135, 21: 333, 331, 332, 108: 476, 146: 480, 151: 541},
		{138, 2: 138, 8: 117, 10: 117, 138},
		{8: 118},
		{471, 11: 103, 169: 470, 472},
		{8: 114, 10: 114},
		// 175
		{537, 8: 116, 10: 116, 102},
		{139, 465, 8: 107, 11: 139, 21: 107, 107, 107, 25: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 42: 107, 44: 107, 107, 107, 107, 107, 130: 466, 468, 147: 493, 163: 494},
		{11: 473},
		{365, 475, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 333, 331, 332, 42: 479, 50: 474, 227, 108: 476, 146: 477, 157: 478},
		{379, 384, 3: 397, 387, 388, 383, 382, 11: 378, 49: 403, 51: 226, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		// 180
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 491, 492},
		{137, 137, 137, 137, 137, 137, 137, 137, 137, 10: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 42: 137, 51: 137},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 333, 331, 332, 42: 487, 50: 474, 227, 108: 484, 157: 486},
		{51: 485},
		{135, 135, 135, 135, 135, 135, 135, 135, 12: 135, 135, 135, 135, 135, 135, 135, 135, 135, 333, 331, 332, 108: 476, 146: 480, 151: 481},
		// 185
		{134, 134, 134, 134, 134, 134, 134, 134, 134, 10: 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 333, 331, 332, 108: 484},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 482},
		{379, 384, 3: 397, 387, 388, 383, 382, 11: 378, 49: 403, 51: 483, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{110, 8: 110, 10: 110, 110},
		{136, 136, 136, 136, 136, 136, 136, 136, 136, 10: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 42: 136, 51: 136},
		// 190
		{112, 8: 112, 10: 112, 112},
		{51: 490},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 488},
		{379, 384, 3: 397, 387, 388, 383, 382, 11: 378, 49: 403, 51: 489, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{109, 8: 109, 10: 109, 109},
		// 195
		{111, 8: 111, 10: 111, 111},
		{379, 265, 3: 265, 265, 265, 383, 382, 265, 265, 265, 378, 24: 265, 41: 265, 43: 265, 49: 265, 51: 265, 380, 54: 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 381, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265},
		{108, 8: 108, 10: 108, 108},
		{8: 536},
		{8: 131, 21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 108: 305, 113: 324, 327, 323, 304, 119: 498, 306, 303, 150: 497, 160: 495, 496, 183: 499},
		// 200
		{8: 133, 10: 533},
		{8: 130},
		{8: 129, 10: 129},
		{139, 465, 139, 8: 115, 10: 115, 139, 130: 466, 501, 133: 502, 147: 469, 165: 503},
		{8: 500},
		// 205
		{106, 8: 106, 10: 106, 106},
		{506, 2: 505, 11: 103, 169: 470, 472, 504},
		{8: 127, 10: 127},
		{8: 126, 10: 126},
		{510, 8: 153, 153, 153, 509, 21: 153, 153, 153, 25: 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 44: 153, 153, 153, 153, 153, 153, 53: 153},
		// 210
		{150, 8: 150, 150, 150, 150, 21: 150, 150, 150, 25: 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 44: 150, 150, 150, 150, 150, 150, 53: 150},
		{139, 465, 139, 8: 107, 11: 139, 21: 107, 107, 107, 25: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 42: 107, 44: 107, 107, 107, 107, 107, 130: 466, 501, 133: 507, 147: 493, 163: 494},
		{8: 508},
		{149, 8: 149, 149, 149, 149, 21: 149, 149, 149, 25: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 44: 149, 149, 149, 149, 149, 149, 53: 149},
		{135, 135, 135, 135, 135, 135, 135, 135, 12: 135, 135, 135, 135, 135, 135, 135, 135, 135, 333, 331, 332, 42: 521, 51: 135, 108: 476, 146: 522, 151: 520},
		// 215
		{2: 513, 8: 123, 21: 144, 144, 144, 25: 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 42: 144, 44: 144, 144, 144, 144, 144, 177: 514, 512, 197: 511},
		{21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 108: 305, 113: 324, 327, 323, 304, 119: 498, 306, 303, 150: 497, 160: 495, 518},
		{8: 517},
		{8: 125, 10: 125},
		{8: 122, 10: 515},
		// 220
		{2: 516},
		{8: 124, 10: 124},
		{142, 8: 142, 142, 142, 142, 21: 142, 142, 142, 25: 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 44: 142, 142, 142, 142, 142, 142, 53: 142},
		{8: 519},
		{143, 8: 143, 143, 143, 143, 21: 143, 143, 143, 25: 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 44: 143, 143, 143, 143, 143, 143, 53: 143},
		// 225
		{365, 529, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 474, 227, 157: 530},
		{135, 135, 135, 135, 135, 135, 135, 135, 12: 135, 135, 135, 135, 135, 135, 135, 135, 135, 333, 331, 332, 108: 476, 146: 480, 151: 526},
		{134, 134, 134, 134, 134, 134, 134, 134, 12: 134, 134, 134, 134, 134, 134, 134, 134, 134, 333, 331, 332, 42: 523, 51: 134, 108: 484},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 524},
		{379, 384, 3: 397, 387, 388, 383, 382, 11: 378, 49: 403, 51: 525, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		// 230
		{146, 8: 146, 146, 146, 146, 21: 146, 146, 146, 25: 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 44: 146, 146, 146, 146, 146, 146, 53: 146},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 527},
		{379, 384, 3: 397, 387, 388, 383, 382, 11: 378, 49: 403, 51: 528, 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{147, 8: 147, 147, 147, 147, 21: 147, 147, 147, 25: 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 44: 147, 147, 147, 147, 147, 147, 53: 147},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 491, 532},
		// 235
		{51: 531},
		{148, 8: 148, 148, 148, 148, 21: 148, 148, 148, 25: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 44: 148, 148, 148, 148, 148, 148, 53: 148},
		{145, 8: 145, 145, 145, 145, 21: 145, 145, 145, 25: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 44: 145, 145, 145, 145, 145, 145, 53: 145},
		{21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 108: 305, 113: 324, 327, 323, 304, 119: 498, 306, 303, 143: 534, 150: 535},
		{8: 132},
		// 240
		{8: 128, 10: 128},
		{113, 8: 113, 10: 113, 113},
		{8: 105, 21: 105, 105, 105, 25: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 42: 105, 44: 105, 105, 105, 105, 105, 188: 538},
		{8: 131, 21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 108: 305, 113: 324, 327, 323, 304, 119: 498, 306, 303, 150: 497, 160: 495, 496, 183: 539},
		{8: 540},
		// 245
		{104, 8: 104, 10: 104, 104},
		{141, 465, 141, 8: 141, 10: 141, 141, 130: 542},
		{140, 2: 140, 8: 140, 10: 140, 140},
		{173, 173, 173, 8: 173, 173, 11: 173, 41: 173},
		{171, 171, 171, 8: 171, 171, 11: 171, 41: 171},
		// 250
		{174, 174, 174, 8: 174, 174, 11: 174, 41: 174},
		{365, 259, 358, 259, 259, 259, 368, 367, 259, 259, 259, 259, 364, 374, 373, 359, 360, 361, 362, 363, 375, 24: 259, 41: 259, 43: 259, 49: 259, 547, 259, 259, 548, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259},
		{379, 258, 3: 258, 258, 258, 383, 382, 258, 258, 258, 378, 24: 258, 41: 258, 43: 258, 49: 258, 51: 258, 380, 54: 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 381, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258},
		{96, 96, 96, 96, 96, 96, 96, 96, 11: 554, 96, 96, 96, 96, 96, 96, 96, 96, 96, 52: 555, 96, 149: 553, 154: 552, 550, 551, 182: 549},
		{10: 565, 24: 166, 153: 570},
		// 255
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 561, 53: 562, 159: 563},
		{11: 554, 49: 559, 52: 555, 149: 560},
		{95, 95, 95, 95, 95, 95, 95, 95, 12: 95, 95, 95, 95, 95, 95, 95, 95, 95, 53: 95},
		{11: 94, 49: 94, 52: 94},
		{221, 221, 221, 221, 221, 221, 221, 221, 12: 221, 221, 221, 221, 221, 221, 221, 221, 221, 141: 356, 557},
		// 260
		{2: 556},
		{11: 91, 49: 91, 52: 91},
		{51: 558},
		{11: 92, 49: 92, 52: 92},
		{97, 97, 97, 97, 97, 97, 97, 97, 12: 97, 97, 97, 97, 97, 97, 97, 97, 97, 53: 97},
		// 265
		{11: 93, 49: 93, 52: 93},
		{379, 384, 3: 397, 387, 388, 383, 382, 9: 101, 101, 378, 24: 101, 49: 403, 52: 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{96, 96, 96, 96, 96, 96, 96, 96, 11: 554, 96, 96, 96, 96, 96, 96, 96, 96, 96, 52: 555, 96, 149: 553, 154: 552, 550, 551, 182: 564},
		{10: 99, 24: 99},
		{10: 565, 24: 166, 153: 566},
		// 270
		{96, 96, 96, 96, 96, 96, 96, 96, 11: 554, 96, 96, 96, 96, 96, 96, 96, 96, 96, 24: 165, 52: 555, 96, 149: 553, 154: 552, 568, 551},
		{24: 567},
		{9: 100, 100, 24: 100},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 561, 53: 562, 159: 569},
		{10: 98, 24: 98},
		// 275
		{24: 571},
		{269, 269, 3: 269, 269, 269, 269, 269, 269, 269, 269, 269, 24: 269, 41: 269, 43: 269, 49: 269, 51: 269, 269, 54: 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269},
		{276, 276, 3: 276, 276, 276, 276, 276, 276, 276, 276, 276, 24: 276, 41: 276, 43: 276, 49: 276, 51: 276, 276, 54: 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276},
		{379, 261, 3: 261, 261, 261, 383, 382, 261, 261, 261, 378, 24: 261, 41: 261, 43: 261, 49: 261, 51: 261, 380, 54: 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 381, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261},
		{379, 262, 3: 262, 262, 262, 383, 382, 262, 262, 262, 378, 24: 262, 41: 262, 43: 262, 49: 262, 51: 262, 380, 54: 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 381, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262},
		// 280
		{379, 263, 3: 263, 263, 263, 383, 382, 263, 263, 263, 378, 24: 263, 41: 263, 43: 263, 49: 263, 51: 263, 380, 54: 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 381, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263},
		{379, 264, 3: 264, 264, 264, 383, 382, 264, 264, 264, 378, 24: 264, 41: 264, 43: 264, 49: 264, 51: 264, 380, 54: 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 381, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264},
		{379, 266, 3: 266, 266, 266, 383, 382, 266, 266, 266, 378, 24: 266, 41: 266, 43: 266, 49: 266, 51: 266, 380, 54: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 381, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266},
		{379, 267, 3: 267, 267, 267, 383, 382, 267, 267, 267, 378, 24: 267, 41: 267, 43: 267, 49: 267, 51: 267, 380, 54: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 381, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267},
		{379, 268, 3: 268, 268, 268, 383, 382, 268, 268, 268, 378, 24: 268, 41: 268, 43: 268, 49: 268, 51: 268, 380, 54: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 381, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268},
		// 285
		{8: 581},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 547, 53: 548},
		{2: 351, 24: 165, 173: 354, 585},
		{24: 584},
		{163, 163, 163, 8: 163, 163, 163, 163, 21: 163, 163, 163, 25: 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 44: 163, 163, 163, 163, 163},
		// 290
		{10: 160, 24: 160},
		{53: 184, 194: 588},
		{181, 181, 181, 8: 181, 181, 181, 181, 21: 181, 181, 181, 25: 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 44: 181, 181, 181, 181, 181, 53: 120},
		{53: 589},
		{21: 183, 183, 183, 25: 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 195: 590},
		// 295
		{21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 108: 463, 113: 324, 327, 323, 462, 145: 593, 185: 592, 216: 591},
		{21: 333, 331, 332, 607, 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 108: 463, 113: 324, 327, 323, 462, 145: 593, 185: 608},
		{21: 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178},
		{139, 465, 139, 9: 595, 41: 152, 130: 600, 599, 133: 597, 168: 598, 186: 596, 217: 594},
		{9: 604, 605},
		// 300
		{21: 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175},
		{9: 170, 170},
		{9: 168, 168, 41: 151},
		{41: 602},
		{601, 2: 505, 171: 504},
		// 305
		{138, 2: 138},
		{139, 465, 139, 130: 600, 599, 133: 507},
		{221, 221, 221, 221, 221, 221, 221, 221, 12: 221, 221, 221, 221, 221, 221, 221, 221, 221, 141: 356, 603},
		{9: 167, 167},
		{21: 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176},
		// 310
		{139, 465, 139, 41: 152, 130: 600, 599, 133: 597, 168: 598, 186: 606},
		{9: 169, 169},
		{182, 182, 182, 8: 182, 182, 182, 182, 21: 182, 182, 182, 25: 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 44: 182, 182, 182, 182, 182},
		{21: 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 119, 119, 119, 25: 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 50: 610, 152: 461, 162: 611},
		// 315
		{379, 384, 3: 397, 387, 388, 383, 382, 613, 11: 378, 49: 403, 52: 380, 54: 386, 385, 391, 392, 402, 398, 399, 407, 400, 411, 381, 405, 395, 394, 393, 389, 409, 406, 404, 396, 413, 401, 390, 410, 408, 412},
		{8: 612},
		{185, 185, 185, 8: 185, 185, 185, 185, 21: 185, 185, 185, 25: 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 44: 185, 185, 185, 185, 185},
		{186, 186, 186, 8: 186, 186, 186, 186, 21: 186, 186, 186, 25: 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 44: 186, 186, 186, 186, 186},
		{215, 215, 215, 8: 215, 215, 215, 215},
		// 320
		{213, 213, 213, 8: 213, 213, 213, 213},
		{216, 216, 216, 8: 216, 216, 216, 216},
		{217, 217, 217, 8: 217, 217, 217, 217},
		{218, 218, 218, 8: 218, 218, 218, 218},
		{9: 714},
		// 325
		{9: 212, 212},
		{9: 209, 712},
		{9: 208, 208, 21: 51, 51, 51, 25: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 42: 51, 44: 51, 51, 51, 51, 51, 207, 53: 52, 164: 623, 191: 625, 205: 624},
		{49: 710},
		{53: 56, 190: 631},
		// 330
		{21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 108: 305, 113: 324, 327, 323, 304, 119: 626, 306, 303, 138: 627, 204: 628},
		{139, 465, 139, 9: 210, 130: 600, 599, 133: 630, 158: 620, 180: 621, 619},
		{21: 54, 54, 54, 25: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 42: 54, 44: 54, 54, 54, 54, 54, 53: 54},
		{21: 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 53: 50, 108: 305, 113: 324, 327, 323, 304, 119: 626, 306, 303, 138: 629},
		{21: 53, 53, 53, 25: 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 42: 53, 44: 53, 53, 53, 53, 53, 53: 53},
		// 335
		{9: 208, 208, 49: 207, 164: 623},
		{53: 632, 122: 633},
		{81, 81, 81, 81, 81, 81, 81, 81, 9: 81, 12: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 42: 81, 44: 81, 81, 81, 81, 81, 53: 81, 84: 81, 81, 81, 81, 81, 81, 81, 81, 81, 94: 81, 81, 189: 634},
		{21: 55, 55, 55, 25: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 42: 55, 55, 55, 55, 55, 55, 55, 123: 55},
		{365, 370, 635, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 333, 331, 332, 77, 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 50: 425, 53: 632, 84: 653, 658, 643, 657, 644, 654, 655, 656, 651, 94: 659, 652, 108: 305, 636, 113: 324, 327, 323, 304, 650, 119: 626, 306, 303, 638, 124: 639, 641, 642, 637, 640, 649, 138: 648, 167: 646, 201: 647, 645},
		// 340
		{283, 283, 3: 283, 283, 283, 283, 283, 9: 283, 283, 283, 41: 708, 49: 283, 52: 283, 54: 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283},
		{8: 222, 222, 428},
		{90, 90, 90, 90, 90, 90, 90, 90, 9: 90, 12: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 42: 90, 44: 90, 90, 90, 90, 90, 53: 90, 84: 90, 90, 90, 90, 90, 90, 90, 90, 90, 94: 90, 90, 112: 90},
		{89, 89, 89, 89, 89, 89, 89, 89, 9: 89, 12: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 42: 89, 44: 89, 89, 89, 89, 89, 53: 89, 84: 89, 89, 89, 89, 89, 89, 89, 89, 89, 94: 89, 89, 112: 89},
		{88, 88, 88, 88, 88, 88, 88, 88, 9: 88, 12: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 42: 88, 44: 88, 88, 88, 88, 88, 53: 88, 84: 88, 88, 88, 88, 88, 88, 88, 88, 88, 94: 88, 88, 112: 88},
		// 345
		{87, 87, 87, 87, 87, 87, 87, 87, 9: 87, 12: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 42: 87, 44: 87, 87, 87, 87, 87, 53: 87, 84: 87, 87, 87, 87, 87, 87, 87, 87, 87, 94: 87, 87, 112: 87},
		{86, 86, 86, 86, 86, 86, 86, 86, 9: 86, 12: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 42: 86, 44: 86, 86, 86, 86, 86, 53: 86, 84: 86, 86, 86, 86, 86, 86, 86, 86, 86, 94: 86, 86, 112: 86},
		{85, 85, 85, 85, 85, 85, 85, 85, 9: 85, 12: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 42: 85, 44: 85, 85, 85, 85, 85, 53: 85, 84: 85, 85, 85, 85, 85, 85, 85, 85, 85, 94: 85, 85, 112: 85},
		{221, 221, 221, 221, 221, 221, 221, 221, 12: 221, 221, 221, 221, 221, 221, 221, 221, 221, 141: 356, 705},
		{41: 703},
		// 350
		{24: 702},
		{79, 79, 79, 79, 79, 79, 79, 79, 9: 79, 12: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 42: 79, 44: 79, 79, 79, 79, 79, 53: 79, 84: 79, 79, 79, 79, 79, 79, 79, 79, 79, 94: 79, 79},
		{365, 370, 635, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 333, 331, 332, 76, 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 50: 425, 53: 632, 84: 653, 658, 643, 657, 644, 654, 655, 656, 651, 94: 659, 652, 108: 305, 636, 113: 324, 327, 323, 304, 650, 119: 626, 306, 303, 638, 124: 639, 641, 642, 637, 640, 649, 138: 648, 167: 701},
		{75, 75, 75, 75, 75, 75, 75, 75, 9: 75, 12: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 42: 75, 44: 75, 75, 75, 75, 75, 53: 75, 84: 75, 75, 75, 75, 75, 75, 75, 75, 75, 94: 75, 75},
		{74, 74, 74, 74, 74, 74, 74, 74, 9: 74, 12: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 42: 74, 44: 74, 74, 74, 74, 74, 53: 74, 84: 74, 74, 74, 74, 74, 74, 74, 74, 74, 94: 74, 74},
		// 355
		{9: 700},
		{694},
		{690},
		{686},
		{365, 370, 635, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 53: 632, 84: 653, 658, 643, 657, 644, 654, 655, 656, 651, 94: 659, 652, 109: 636, 117: 650, 122: 638, 124: 639, 641, 642, 637, 640, 680},
		// 360
		{666},
		{2: 664},
		{9: 663},
		{9: 662},
		{365, 370, 358, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 109: 636, 117: 660},
		// 365
		{9: 661},
		{62, 62, 62, 62, 62, 62, 62, 62, 9: 62, 12: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 42: 62, 44: 62, 62, 62, 62, 62, 53: 62, 84: 62, 62, 62, 62, 62, 62, 62, 62, 62, 94: 62, 62, 112: 62},
		{63, 63, 63, 63, 63, 63, 63, 63, 9: 63, 12: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 42: 63, 44: 63, 63, 63, 63, 63, 53: 63, 84: 63, 63, 63, 63, 63, 63, 63, 63, 63, 94: 63, 63, 112: 63},
		{64, 64, 64, 64, 64, 64, 64, 64, 9: 64, 12: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 42: 64, 44: 64, 64, 64, 64, 64, 53: 64, 84: 64, 64, 64, 64, 64, 64, 64, 64, 64, 94: 64, 64, 112: 64},
		{9: 665},
		// 370
		{65, 65, 65, 65, 65, 65, 65, 65, 9: 65, 12: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 42: 65, 44: 65, 65, 65, 65, 65, 53: 65, 84: 65, 65, 65, 65, 65, 65, 65, 65, 65, 94: 65, 65, 112: 65},
		{365, 370, 358, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 333, 331, 332, 25: 321, 313, 322, 318, 330, 317, 315, 316, 314, 319, 328, 325, 326, 329, 320, 312, 42: 309, 44: 310, 308, 334, 311, 307, 50: 425, 108: 305, 636, 113: 324, 327, 323, 304, 667, 119: 626, 306, 303, 138: 668},
		{9: 674},
		{365, 370, 358, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 109: 636, 117: 669},
		{9: 670},
		// 375
		{365, 370, 358, 369, 371, 372, 368, 367, 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 109: 636, 117: 671},
		{8: 672},
		{365, 370, 635, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 53: 632, 84: 653, 658, 643, 657, 644, 654, 655, 656, 651, 94: 659, 652, 109: 636, 117: 650, 122: 638, 124: 639, 641, 642, 637, 640, 673},
		{66, 66, 66, 66, 66, 66, 66, 66, 9: 66, 12: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 42: 66, 44: 66, 66, 66, 66, 66, 53: 66, 84: 66, 66, 66, 66, 66, 66, 66, 66, 66, 94: 66, 66, 112: 66},
		{365, 370, 358, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 109: 636, 117: 675},
		// 380
		{9: 676},
		{365, 370, 358, 369, 371, 372, 368, 367, 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 109: 636, 117: 677},
		{8: 678},
		{365, 370, 635, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 53: 632, 84: 653, 658, 643, 657, 644, 654, 655, 656, 651, 94: 659, 652, 109: 636, 117: 650, 122: 638, 124: 639, 641, 642, 637, 640, 679},
		{67, 67, 67, 67, 67, 67, 67, 67, 9: 67, 12: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 42: 67, 44: 67, 67, 67, 67, 67, 53: 67, 84: 67, 67, 67, 67, 67, 67, 67, 67, 67, 94: 67, 67, 112: 67},
		// 385
		{84: 681},
		{682},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 109: 683},
		{8: 684, 10: 428},
		{9: 685},
		// 390
		{68, 68, 68, 68, 68, 68, 68, 68, 9: 68, 12: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 42: 68, 44: 68, 68, 68, 68, 68, 53: 68, 84: 68, 68, 68, 68, 68, 68, 68, 68, 68, 94: 68, 68, 112: 68},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 109: 687},
		{8: 688, 10: 428},
		{365, 370, 635, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 53: 632, 84: 653, 658, 643, 657, 644, 654, 655, 656, 651, 94: 659, 652, 109: 636, 117: 650, 122: 638, 124: 639, 641, 642, 637, 640, 689},
		{69, 69, 69, 69, 69, 69, 69, 69, 9: 69, 12: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 42: 69, 44: 69, 69, 69, 69, 69, 53: 69, 84: 69, 69, 69, 69, 69, 69, 69, 69, 69, 94: 69, 69, 112: 69},
		// 395
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 109: 691},
		{8: 692, 10: 428},
		{365, 370, 635, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 53: 632, 84: 653, 658, 643, 657, 644, 654, 655, 656, 651, 94: 659, 652, 109: 636, 117: 650, 122: 638, 124: 639, 641, 642, 637, 640, 693},
		{70, 70, 70, 70, 70, 70, 70, 70, 9: 70, 12: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 42: 70, 44: 70, 70, 70, 70, 70, 53: 70, 84: 70, 70, 70, 70, 70, 70, 70, 70, 70, 94: 70, 70, 112: 70},
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 109: 695},
		// 400
		{8: 696, 10: 428},
		{365, 370, 635, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 53: 632, 84: 653, 658, 643, 657, 644, 654, 655, 656, 651, 94: 659, 652, 109: 636, 117: 650, 122: 638, 124: 639, 641, 642, 637, 640, 697},
		{72, 72, 72, 72, 72, 72, 72, 72, 9: 72, 12: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 42: 72, 44: 72, 72, 72, 72, 72, 53: 72, 84: 72, 72, 72, 72, 72, 72, 72, 72, 72, 94: 72, 72, 112: 698},
		{365, 370, 635, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 53: 632, 84: 653, 658, 643, 657, 644, 654, 655, 656, 651, 94: 659, 652, 109: 636, 117: 650, 122: 638, 124: 639, 641, 642, 637, 640, 699},
		{71, 71, 71, 71, 71, 71, 71, 71, 9: 71, 12: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 42: 71, 44: 71, 71, 71, 71, 71, 53: 71, 84: 71, 71, 71, 71, 71, 71, 71, 71, 71, 94: 71, 71, 112: 71},
		// 405
		{73, 73, 73, 73, 73, 73, 73, 73, 9: 73, 12: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 42: 73, 44: 73, 73, 73, 73, 73, 53: 73, 84: 73, 73, 73, 73, 73, 73, 73, 73, 73, 94: 73, 73, 112: 73},
		{78, 78, 78, 78, 78, 78, 78, 78, 9: 78, 12: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 42: 78, 44: 78, 78, 78, 78, 78, 53: 78, 84: 78, 78, 78, 78, 78, 78, 78, 78, 78, 94: 78, 78},
		{80, 80, 80, 80, 80, 80, 80, 80, 9: 80, 12: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 42: 80, 80, 80, 80, 80, 80, 80, 53: 80, 84: 80, 80, 80, 80, 80, 80, 80, 80, 80, 94: 80, 80, 112: 80, 123: 80},
		{365, 370, 635, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 53: 632, 84: 653, 658, 643, 657, 644, 654, 655, 656, 651, 94: 659, 652, 109: 636, 117: 650, 122: 638, 124: 639, 641, 642, 637, 640, 704},
		{82, 82, 82, 82, 82, 82, 82, 82, 9: 82, 12: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 42: 82, 44: 82, 82, 82, 82, 82, 53: 82, 84: 82, 82, 82, 82, 82, 82, 82, 82, 82, 94: 82, 82, 112: 82},
		// 410
		{41: 706},
		{365, 370, 635, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 53: 632, 84: 653, 658, 643, 657, 644, 654, 655, 656, 651, 94: 659, 652, 109: 636, 117: 650, 122: 638, 124: 639, 641, 642, 637, 640, 707},
		{83, 83, 83, 83, 83, 83, 83, 83, 9: 83, 12: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 42: 83, 44: 83, 83, 83, 83, 83, 53: 83, 84: 83, 83, 83, 83, 83, 83, 83, 83, 83, 94: 83, 83, 112: 83},
		{365, 370, 635, 369, 371, 372, 368, 367, 9: 223, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 425, 53: 632, 84: 653, 658, 643, 657, 644, 654, 655, 656, 651, 94: 659, 652, 109: 636, 117: 650, 122: 638, 124: 639, 641, 642, 637, 640, 709},
		{84, 84, 84, 84, 84, 84, 84, 84, 9: 84, 12: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 42: 84, 44: 84, 84, 84, 84, 84, 53: 84, 84: 84, 84, 84, 84, 84, 84, 84, 84, 84, 94: 84, 84, 112: 84},
		// 415
		{365, 370, 358, 369, 371, 372, 368, 367, 12: 364, 374, 373, 359, 360, 361, 362, 363, 375, 50: 561, 53: 562, 159: 711},
		{9: 206, 206},
		{139, 465, 139, 130: 600, 599, 133: 630, 158: 713},
		{9: 211, 211},
		{219, 219, 219, 219, 219, 219, 219, 219, 9: 219, 12: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 42: 219, 219, 219, 219, 219, 219, 219, 53: 219, 84: 219, 219, 219, 219, 219, 219, 219, 219, 219, 94: 219, 219, 123: 219},
		// 420
		{21: 60, 60, 60, 25: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 42: 60, 60, 60, 60, 60, 60, 60, 123: 60},
		{221, 221, 221, 221, 221, 221, 221, 221, 12: 221, 221, 221, 221, 221, 221, 221, 221, 221, 141: 356, 717},
		{43: 291},
		{80: 740, 742, 96: 730, 731, 732, 727, 728, 729, 733, 737, 734, 724, 735, 736, 110: 741, 739, 118: 738, 132: 722, 134: 721, 726, 723, 725, 139: 720, 214: 719},
		{43: 293},
		// 425
		{43: 44, 80: 740, 742, 96: 730, 731, 732, 727, 728, 729, 733, 737, 734, 724, 735, 736, 110: 741, 739, 118: 738, 132: 722, 134: 785, 726, 723, 725},
		{43: 43, 80: 43, 43, 43, 43, 93: 43, 96: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43},
		{43: 39, 80: 39, 39, 39, 39, 93: 39, 96: 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39},
		{43: 38, 80: 38, 38, 38, 38, 93: 38, 96: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{81: 742, 110: 807, 739},
		// 430
		{43: 36, 80: 36, 36, 36, 36, 93: 36, 96: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
		{82: 29, 29, 93: 795, 172: 793, 206: 794, 792},
		{81: 742, 110: 789, 739},
		{2: 786},
		{2: 781},
		// 435
		{2: 757, 80: 759, 212: 758},
		{80: 740, 742, 110: 741, 739, 118: 756},
		{43: 17, 80: 17, 17, 17, 17, 93: 17, 96: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{81: 742, 110: 754, 739},
		{81: 742, 110: 752, 739},
		// 440
		{80: 740, 742, 110: 741, 739, 118: 751},
		{2: 747},
		{81: 742, 110: 745, 739},
		{43: 7, 80: 7, 7, 7, 7, 93: 7, 96: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{80: 5, 744},
		// 445
		{43: 4, 80: 4, 4, 4, 4, 93: 4, 96: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{80: 743},
		{80: 2, 2},
		{43: 3, 80: 3, 3, 3, 3, 93: 3, 96: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{80: 1, 1},
		// 450
		{80: 746},
		{43: 8, 80: 8, 8, 8, 8, 93: 8, 96: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{80: 748, 742, 110: 749, 739},
		{43: 13, 80: 13, 13, 13, 13, 93: 13, 96: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{80: 750},
		// 455
		{43: 9, 80: 9, 9, 9, 9, 93: 9, 96: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{43: 14, 80: 14, 14, 14, 14, 93: 14, 96: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{80: 753},
		{43: 15, 80: 15, 15, 15, 15, 93: 15, 96: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{80: 755},
		// 460
		{43: 16, 80: 16, 16, 16, 16, 93: 16, 96: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{43: 18, 80: 18, 18, 18, 18, 93: 18, 96: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{80: 740, 742, 110: 741, 739, 118: 766, 140: 780},
		{2: 760, 8: 123, 143: 762, 177: 761, 763},
		{43: 10, 80: 10, 10, 10, 10, 93: 10, 96: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		// 465
		{8: 125, 10: 125, 143: 777},
		{8: 122, 10: 769},
		{8: 767},
		{8: 764},
		{80: 740, 742, 110: 741, 739, 118: 766, 140: 765},
		// 470
		{43: 19, 80: 19, 19, 19, 19, 93: 19, 96: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{43: 6, 80: 6, 6, 6, 6, 93: 6, 96: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{80: 740, 742, 110: 741, 739, 118: 766, 140: 768},
		{43: 21, 80: 21, 21, 21, 21, 93: 21, 96: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{2: 770, 143: 771},
		// 475
		{8: 124, 10: 124, 143: 774},
		{8: 772},
		{80: 740, 742, 110: 741, 739, 118: 766, 140: 773},
		{43: 20, 80: 20, 20, 20, 20, 93: 20, 96: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{8: 775},
		// 480
		{80: 740, 742, 110: 741, 739, 118: 766, 140: 776},
		{43: 11, 80: 11, 11, 11, 11, 93: 11, 96: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{8: 778},
		{80: 740, 742, 110: 741, 739, 118: 766, 140: 779},
		{43: 12, 80: 12, 12, 12, 12, 93: 12, 96: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		// 485
		{43: 22, 80: 22, 22, 22, 22, 93: 22, 96: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{80: 782},
		{80: 740, 742, 41, 41, 93: 41, 96: 730, 731, 732, 727, 728, 729, 733, 737, 734, 724, 735, 736, 110: 741, 739, 118: 738, 132: 722, 134: 721, 726, 723, 725, 139: 783, 144: 784},
		{80: 740, 742, 40, 40, 93: 40, 96: 730, 731, 732, 727, 728, 729, 733, 737, 734, 724, 735, 736, 110: 741, 739, 118: 738, 132: 722, 134: 785, 726, 723, 725},
		{82: 32, 32, 93: 32},
		// 490
		{43: 42, 80: 42, 42, 42, 42, 93: 42, 96: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{80: 787},
		{80: 740, 742, 41, 41, 93: 41, 96: 730, 731, 732, 727, 728, 729, 733, 737, 734, 724, 735, 736, 110: 741, 739, 118: 738, 132: 722, 134: 721, 726, 723, 725, 139: 783, 144: 788},
		{82: 33, 33, 93: 33},
		{80: 790},
		// 495
		{80: 740, 742, 41, 41, 93: 41, 96: 730, 731, 732, 727, 728, 729, 733, 737, 734, 724, 735, 736, 110: 741, 739, 118: 738, 132: 722, 134: 721, 726, 723, 725, 139: 783, 144: 791},
		{82: 34, 34, 93: 34},
		{82: 25, 801, 208: 802, 800},
		{82: 31, 31, 93: 31},
		{82: 28, 28, 93: 795, 172: 799},
		// 500
		{81: 742, 110: 796, 739},
		{80: 797},
		{80: 740, 742, 41, 41, 93: 41, 96: 730, 731, 732, 727, 728, 729, 733, 737, 734, 724, 735, 736, 110: 741, 739, 118: 738, 132: 722, 134: 721, 726, 723, 725, 139: 783, 144: 798},
		{82: 27, 27, 93: 27},
		{82: 30, 30, 93: 30},
		// 505
		{82: 806, 210: 805},
		{80: 803},
		{82: 24},
		{80: 740, 742, 41, 96: 730, 731, 732, 727, 728, 729, 733, 737, 734, 724, 735, 736, 110: 741, 739, 118: 738, 132: 722, 134: 721, 726, 723, 725, 139: 783, 144: 804},
		{82: 26},
		// 510
		{43: 35, 80: 35, 35, 35, 35, 93: 35, 96: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{43: 23, 80: 23, 23, 23, 23, 93: 23, 96: 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23},
		{80: 808},
		{43: 37, 80: 37, 37, 37, 37, 93: 37, 96: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
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
	const yyError = 223

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
					if l.DesignationOpt != nil {
						panic("TODO")
					}

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
			lhs := &TypeSpecifier{
				Case:       14,
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsTypeof)
			_, lhs.Type = lhs.Expression.eval(lx)
		}
	case 110:
		{
			lhs := &TypeSpecifier{
				Case:     15,
				Token:    yyS[yypt-3].Token,
				Token2:   yyS[yypt-2].Token,
				TypeName: yyS[yypt-1].node.(*TypeName),
				Token3:   yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsTypeof)
			lhs.Type = undefined
			if t := lhs.TypeName.Type; t != nil {
				lhs.Type = t
			}
		}
	case 111:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareStructTag(o.Token, lx.report)
			}
		}
	case 112:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeMembers)
			lx.scope.isUnion = yyS[yypt-3].node.(*StructOrUnion).Case == 1 // "union"
			lx.scope.prevStructDeclarator = nil
		}
	case 113:
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
	case 114:
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
	case 115:
		{
			yyVAL.node = &StructOrUnion{
				Token: yyS[yypt-0].Token,
			}
		}
	case 116:
		{
			yyVAL.node = &StructOrUnion{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 117:
		{
			yyVAL.node = &StructDeclarationList{
				StructDeclaration: yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 118:
		{
			yyVAL.node = &StructDeclarationList{
				Case: 1,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList),
				StructDeclaration:     yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 119:
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
	case 120:
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
	case 121:
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
	case 122:
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
	case 123:
		{
			yyVAL.node = (*SpecifierQualifierListOpt)(nil)
		}
	case 124:
		{
			lhs := &SpecifierQualifierListOpt{
				SpecifierQualifierList: yyS[yypt-0].node.(*SpecifierQualifierList),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.SpecifierQualifierList.attr
			lhs.typeSpecifier = lhs.SpecifierQualifierList.typeSpecifier
		}
	case 125:
		{
			yyVAL.node = &StructDeclaratorList{
				StructDeclarator: yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 126:
		{
			yyVAL.node = &StructDeclaratorList{
				Case:                 1,
				StructDeclaratorList: yyS[yypt-2].node.(*StructDeclaratorList),
				Token:                yyS[yypt-1].Token,
				StructDeclarator:     yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 127:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.post(lx)
		}
	case 128:
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
	case 129:
		{
			yyVAL.node = (*CommaOpt)(nil)
		}
	case 130:
		{
			yyVAL.node = &CommaOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 131:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareEnumTag(o.Token, lx.report)
			}
			lx.iota = 0
		}
	case 132:
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
	case 133:
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
	case 134:
		{
			yyVAL.node = &EnumeratorList{
				Enumerator: yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 135:
		{
			yyVAL.node = &EnumeratorList{
				Case:           1,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList),
				Token:          yyS[yypt-1].Token,
				Enumerator:     yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 136:
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
	case 137:
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
	case 138:
		{
			lhs := &TypeQualifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saConst
		}
	case 139:
		{
			lhs := &TypeQualifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRestrict
		}
	case 140:
		{
			lhs := &TypeQualifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saVolatile
		}
	case 141:
		{
			lhs := &FunctionSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saInline
		}
	case 142:
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
	case 143:
		{
			yyVAL.node = (*DeclaratorOpt)(nil)
		}
	case 144:
		{
			yyVAL.node = &DeclaratorOpt{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
		}
	case 145:
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
	case 146:
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
	case 147:
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
	case 148:
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
	case 149:
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
	case 150:
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
	case 151:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 152:
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
	case 153:
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
	case 154:
		{
			yyVAL.node = &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].node.(*TypeQualifierListOpt),
			}
		}
	case 155:
		{
			yyVAL.node = &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].node.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].node.(*Pointer),
			}
		}
	case 156:
		{
			yyVAL.node = (*PointerOpt)(nil)
		}
	case 157:
		{
			yyVAL.node = &PointerOpt{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
		}
	case 158:
		{
			lhs := &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.TypeQualifier.attr
		}
	case 159:
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
	case 160:
		{
			yyVAL.node = (*TypeQualifierListOpt)(nil)
		}
	case 161:
		{
			yyVAL.node = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].node.(*TypeQualifierList).reverse(),
			}
		}
	case 162:
		{
			lhs := &ParameterTypeList{
				ParameterList: yyS[yypt-0].node.(*ParameterList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 163:
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
	case 164:
		{
			yyVAL.node = (*ParameterTypeListOpt)(nil)
		}
	case 165:
		{
			yyVAL.node = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].node.(*ParameterTypeList),
			}
		}
	case 166:
		{
			yyVAL.node = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 167:
		{
			yyVAL.node = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].node.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 168:
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
	case 169:
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
	case 170:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 171:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 172:
		{
			yyVAL.node = (*IdentifierListOpt)(nil)
		}
	case 173:
		{
			yyVAL.node = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 174:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 175:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 176:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeBlock)
		}
	case 177:
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
	case 178:
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
	case 179:
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
	case 180:
		{
			yyVAL.node = (*AbstractDeclaratorOpt)(nil)
		}
	case 181:
		{
			yyVAL.node = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].node.(*AbstractDeclarator),
			}
		}
	case 182:
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
	case 183:
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
	case 184:
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
	case 185:
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
	case 186:
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
	case 187:
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
	case 188:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 189:
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
	case 190:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 191:
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
	case 192:
		{
			yyVAL.node = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 193:
		{
			yyVAL.node = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
		}
	case 194:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 195:
		{
			yyVAL.node = &Initializer{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 196:
		{
			yyVAL.node = &InitializerList{
				DesignationOpt: yyS[yypt-1].node.(*DesignationOpt),
				Initializer:    yyS[yypt-0].node.(*Initializer),
			}
		}
	case 197:
		{
			yyVAL.node = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].node.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].node.(*DesignationOpt),
				Initializer:     yyS[yypt-0].node.(*Initializer),
			}
		}
	case 198:
		{
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 199:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 200:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 201:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 202:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 203:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 204:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 205:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 206:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 207:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 208:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 209:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 210:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 211:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 212:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 213:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 214:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 215:
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
	case 216:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 217:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 218:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 219:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 220:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 221:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 222:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 223:
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
	case 224:
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
	case 225:
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
	case 226:
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
	case 227:
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
	case 228:
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
	case 229:
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
	case 230:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 231:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 232:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 233:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 234:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 235:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 236:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 237:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 238:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:              2,
				BasicAsmStatement: yyS[yypt-0].node.(*BasicAsmStatement),
			}
		}
	case 239:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-1].node.(*Declarator)
			lx.scope.mergeScope = nil
			dlo := yyS[yypt-0].node.(*DeclarationListOpt)
			done := false
			for dd := d.DirectDeclarator.bottom(); !done && dd != nil; dd = dd.parent {
				switch dd.Case {
				case 6: // DirectDeclarator '(' ParameterTypeList ')'
					done = true
					lx.scope.mergeScope = dd.paramsScope
					if dlo != nil {
						lx.report.Err(dlo.Pos(), "declaration list not allowed in a function definition with parameter type list")
					}
				case 7: // DirectDeclarator '(' IdentifierListOpt ')'
					done = true
					ilo := dd.IdentifierListOpt
					if ilo != nil && dlo == nil {
						lx.report.Err(ilo.Pos(), "missing parameter declaration list")
						break
					}

					if ilo == nil {
						if dlo != nil {
							lx.report.Err(dlo.Pos(), "unexpected parameter declaration list")
						}
						break
					}

					// ilo != nil && dlo != nil
					lx.scope.mergeScope = dlo.paramsScope
					ilo.post(lx.report, dlo.DeclarationList)
				}
			}
			d.setFull(lx)
			if !done {
				lx.report.Err(d.Pos(), "declarator is not a function (have '%s': %v)", d.Type, d.Type.Kind())
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
	case 240:
		{
			yyVAL.node = &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				CompoundStatement:     yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 241:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 242:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 243:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 244:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 245:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScopePos(lhs.Pos())
		}
	case 246:
		{
			yyVAL.node = &AssemblerInstructions{
				Token: yyS[yypt-0].Token,
			}
		}
	case 247:
		{
			yyVAL.node = &AssemblerInstructions{
				Case: 1,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions),
				Token: yyS[yypt-0].Token,
			}
		}
	case 248:
		{
			yyVAL.node = &BasicAsmStatement{
				Token:                 yyS[yypt-4].Token,
				VolatileOpt:           yyS[yypt-3].node.(*VolatileOpt),
				Token2:                yyS[yypt-2].Token,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-0].Token,
			}
		}
	case 249:
		{
			yyVAL.node = (*VolatileOpt)(nil)
		}
	case 250:
		{
			yyVAL.node = &VolatileOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 251:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 252:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 253:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 254:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 255:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 256:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 257:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 258:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 259:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 260:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 261:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 262:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 263:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 264:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 265:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 266:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 267:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 268:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 269:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 270:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 271:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 272:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 273:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 274:
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
	case 275:
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
	case 276:
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
	case 277:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 278:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 279:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 280:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 281:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 282:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 283:
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
	case 284:
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
	case 285:
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
	case 286:
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
	case 287:
		{
			yyVAL.node = &ControlLine{
				Case:        14,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 290:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 291:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
