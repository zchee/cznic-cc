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
	yyDefault           = 57452
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
	TYPEOF              = 57429
	UNARY               = 57430
	UNION               = 57431
	UNSIGNED            = 57432
	VOID                = 57433
	VOLATILE            = 57434
	WHILE               = 57435
	XORASSIGN           = 57436
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -289
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (283x)
		42:      1,   // '*' (253x)
		57375:   2,   // IDENTIFIER (206x)
		38:      3,   // '&' (201x)
		43:      4,   // '+' (201x)
		45:      5,   // '-' (201x)
		57361:   6,   // DEC (201x)
		57379:   7,   // INC (201x)
		59:      8,   // ';' (183x)
		41:      9,   // ')' (180x)
		44:      10,  // ',' (170x)
		91:      11,  // '[' (152x)
		33:      12,  // '!' (133x)
		126:     13,  // '~' (133x)
		57356:   14,  // CHARCONST (133x)
		57371:   15,  // FLOATCONST (133x)
		57382:   16,  // INTCONST (133x)
		57385:   17,  // LONGCHARCONST (133x)
		57386:   18,  // LONGSTRINGLITERAL (133x)
		57421:   19,  // SIZEOF (133x)
		57423:   20,  // STRINGLITERAL (133x)
		57358:   21,  // CONST (117x)
		57415:   22,  // RESTRICT (117x)
		57434:   23,  // VOLATILE (117x)
		125:     24,  // '}' (110x)
		57351:   25,  // BOOL (107x)
		57355:   26,  // CHAR (107x)
		57357:   27,  // COMPLEX (107x)
		57365:   28,  // DOUBLE (107x)
		57367:   29,  // ENUM (107x)
		57370:   30,  // FLOAT (107x)
		57381:   31,  // INT (107x)
		57384:   32,  // LONG (107x)
		57419:   33,  // SHORT (107x)
		57420:   34,  // SIGNED (107x)
		57424:   35,  // STRUCT (107x)
		57428:   36,  // TYPEDEFNAME (107x)
		57429:   37,  // TYPEOF (107x)
		57431:   38,  // UNION (107x)
		57432:   39,  // UNSIGNED (107x)
		57433:   40,  // VOID (107x)
		58:      41,  // ':' (104x)
		57422:   42,  // STATIC (100x)
		57344:   43,  // $end (96x)
		57350:   44,  // AUTO (94x)
		57369:   45,  // EXTERN (94x)
		57380:   46,  // INLINE (94x)
		57414:   47,  // REGISTER (94x)
		57427:   48,  // TYPEDEF (94x)
		61:      49,  // '=' (88x)
		57488:   50,  // Expression (84x)
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
		57363:   65,  // DIVASSIGN (69x)
		57368:   66,  // EQ (69x)
		57373:   67,  // GEQ (69x)
		57383:   68,  // LEQ (69x)
		57387:   69,  // LSH (69x)
		57388:   70,  // LSHASSIGN (69x)
		57389:   71,  // MODASSIGN (69x)
		57390:   72,  // MULASSIGN (69x)
		57391:   73,  // NEQ (69x)
		57393:   74,  // ORASSIGN (69x)
		57394:   75,  // OROR (69x)
		57417:   76,  // RSH (69x)
		57418:   77,  // RSHASSIGN (69x)
		57425:   78,  // SUBASSIGN (69x)
		57436:   79,  // XORASSIGN (69x)
		10:      80,  // '\n' (60x)
		57410:   81,  // PPOTHER (54x)
		57398:   82,  // PPENDIF (45x)
		57397:   83,  // PPELSE (41x)
		57435:   84,  // WHILE (41x)
		57352:   85,  // BREAK (40x)
		57353:   86,  // CASE (40x)
		57359:   87,  // CONTINUE (40x)
		57362:   88,  // DEFAULT (40x)
		57364:   89,  // DO (40x)
		57372:   90,  // FOR (40x)
		57374:   91,  // GOTO (40x)
		57378:   92,  // IF (40x)
		57396:   93,  // PPELIF (40x)
		57416:   94,  // RETURN (40x)
		57426:   95,  // SWITCH (40x)
		57395:   96,  // PPDEFINE (36x)
		57399:   97,  // PPERROR (36x)
		57400:   98,  // PPHASH_NL (36x)
		57402:   99,  // PPIF (36x)
		57403:   100, // PPIFDEF (36x)
		57404:   101, // PPIFNDEF (36x)
		57405:   102, // PPINCLUDE (36x)
		57406:   103, // PPINCLUDE_NEXT (36x)
		57407:   104, // PPLINE (36x)
		57408:   105, // PPNONDIRECTIVE (36x)
		57412:   106, // PPPRAGMA (36x)
		57413:   107, // PPUNDEF (36x)
		57538:   108, // TypeQualifier (28x)
		57489:   109, // ExpressionList (26x)
		57512:   110, // PPTokenList (23x)
		57514:   111, // PPTokens (23x)
		57366:   112, // ELSE (22x)
		57484:   113, // EnumSpecifier (20x)
		57533:   114, // StructOrUnion (20x)
		57534:   115, // StructOrUnionSpecifier (20x)
		57541:   116, // TypeSpecifier (20x)
		57490:   117, // ExpressionListOpt (18x)
		57513:   118, // PPTokenListOpt (16x)
		57467:   119, // DeclarationSpecifiers (15x)
		57495:   120, // FunctionSpecifier (15x)
		57528:   121, // StorageClassSpecifier (15x)
		57461:   122, // CompoundStatement (13x)
		57492:   123, // ExpressionStatement (12x)
		57509:   124, // IterationStatement (12x)
		57510:   125, // JumpStatement (12x)
		57511:   126, // LabeledStatement (12x)
		57523:   127, // SelectionStatement (12x)
		57527:   128, // Statement (12x)
		57519:   129, // Pointer (11x)
		57520:   130, // PointerOpt (10x)
		57463:   131, // ControlLine (8x)
		57469:   132, // Declarator (8x)
		57498:   133, // GroupPart (8x)
		57502:   134, // IfGroup (8x)
		57503:   135, // IfSection (8x)
		57535:   136, // TextLine (8x)
		57464:   137, // Declaration (7x)
		57496:   138, // GroupList (6x)
		57522:   139, // ReplacementList (6x)
		57446:   140, // $@4 (5x)
		57462:   141, // ConstantExpression (5x)
		57360:   142, // DDD (5x)
		57497:   143, // GroupListOpt (5x)
		57524:   144, // SpecifierQualifierList (5x)
		57539:   145, // TypeQualifierList (5x)
		57453:   146, // AbstractDeclarator (4x)
		57468:   147, // DeclarationSpecifiersOpt (4x)
		57473:   148, // Designator (4x)
		57515:   149, // ParameterDeclaration (4x)
		57540:   150, // TypeQualifierListOpt (4x)
		57438:   151, // $@10 (3x)
		57460:   152, // CommaOpt (3x)
		57471:   153, // Designation (3x)
		57472:   154, // DesignationOpt (3x)
		57474:   155, // DesignatorList (3x)
		57491:   156, // ExpressionOpt (3x)
		57504:   157, // InitDeclarator (3x)
		57507:   158, // Initializer (3x)
		57516:   159, // ParameterList (3x)
		57517:   160, // ParameterTypeList (3x)
		57537:   161, // TypeName (3x)
		57439:   162, // $@11 (2x)
		57447:   163, // $@5 (2x)
		57454:   164, // AbstractDeclaratorOpt (2x)
		57457:   165, // BlockItem (2x)
		57470:   166, // DeclaratorOpt (2x)
		57475:   167, // DirectAbstractDeclarator (2x)
		57476:   168, // DirectAbstractDeclaratorOpt (2x)
		57477:   169, // DirectDeclarator (2x)
		57478:   170, // ElifGroup (2x)
		57485:   171, // EnumerationConstant (2x)
		57486:   172, // Enumerator (2x)
		57493:   173, // ExternalDeclaration (2x)
		57494:   174, // FunctionDefinition (2x)
		57499:   175, // IdentifierList (2x)
		57500:   176, // IdentifierListOpt (2x)
		57501:   177, // IdentifierOpt (2x)
		57505:   178, // InitDeclaratorList (2x)
		57506:   179, // InitDeclaratorListOpt (2x)
		57508:   180, // InitializerList (2x)
		57518:   181, // ParameterTypeListOpt (2x)
		57525:   182, // SpecifierQualifierListOpt (2x)
		57529:   183, // StructDeclaration (2x)
		57531:   184, // StructDeclarator (2x)
		57437:   185, // $@1 (1x)
		57440:   186, // $@12 (1x)
		57441:   187, // $@13 (1x)
		57442:   188, // $@14 (1x)
		57443:   189, // $@15 (1x)
		57444:   190, // $@2 (1x)
		57445:   191, // $@3 (1x)
		57448:   192, // $@6 (1x)
		57449:   193, // $@7 (1x)
		57450:   194, // $@8 (1x)
		57451:   195, // $@9 (1x)
		57455:   196, // ArgumentExpressionList (1x)
		57456:   197, // ArgumentExpressionListOpt (1x)
		57458:   198, // BlockItemList (1x)
		57459:   199, // BlockItemListOpt (1x)
		1048577: 200, // CONSTANT_EXPRESSION (1x)
		57465:   201, // DeclarationList (1x)
		57466:   202, // DeclarationListOpt (1x)
		57479:   203, // ElifGroupList (1x)
		57480:   204, // ElifGroupListOpt (1x)
		57481:   205, // ElseGroup (1x)
		57482:   206, // ElseGroupOpt (1x)
		57483:   207, // EndifLine (1x)
		57487:   208, // EnumeratorList (1x)
		57376:   209, // IDENTIFIER_LPAREN (1x)
		1048576: 210, // PREPROCESSING_FILE (1x)
		57521:   211, // PreprocessingFile (1x)
		57526:   212, // Start (1x)
		57530:   213, // StructDeclarationList (1x)
		57532:   214, // StructDeclaratorList (1x)
		1048578: 215, // TRANSLATION_UNIT (1x)
		57536:   216, // TranslationUnit (1x)
		57452:   217, // $default (0x)
		57354:   218, // CAST (0x)
		57345:   219, // error (0x)
		57377:   220, // IDENTIFIER_NONREPL (0x)
		57392:   221, // NOELSE (0x)
		57401:   222, // PPHEADER_NAME (0x)
		57409:   223, // PPNUMBER (0x)
		57411:   224, // PPPASTE (0x)
		57430:   225, // UNARY (0x)
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
		1:   {185, 0},
		2:   {212, 3},
		3:   {190, 0},
		4:   {212, 3},
		5:   {191, 0},
		6:   {212, 3},
		7:   {171, 1},
		8:   {196, 1},
		9:   {196, 3},
		10:  {197, 0},
		11:  {197, 1},
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
		68:  {156, 0},
		69:  {156, 1},
		70:  {109, 1},
		71:  {109, 3},
		72:  {117, 0},
		73:  {117, 1},
		74:  {140, 0},
		75:  {141, 2},
		76:  {137, 3},
		77:  {119, 2},
		78:  {119, 2},
		79:  {119, 2},
		80:  {119, 2},
		81:  {147, 0},
		82:  {147, 1},
		83:  {178, 1},
		84:  {178, 3},
		85:  {179, 0},
		86:  {179, 1},
		87:  {157, 1},
		88:  {163, 0},
		89:  {157, 4},
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
		111: {192, 0},
		112: {193, 0},
		113: {115, 7},
		114: {115, 2},
		115: {114, 1},
		116: {114, 1},
		117: {213, 1},
		118: {213, 2},
		119: {183, 3},
		120: {183, 2},
		121: {144, 2},
		122: {144, 2},
		123: {182, 0},
		124: {182, 1},
		125: {214, 1},
		126: {214, 3},
		127: {184, 1},
		128: {184, 3},
		129: {152, 0},
		130: {152, 1},
		131: {194, 0},
		132: {113, 7},
		133: {113, 2},
		134: {208, 1},
		135: {208, 3},
		136: {172, 1},
		137: {172, 3},
		138: {108, 1},
		139: {108, 1},
		140: {108, 1},
		141: {120, 1},
		142: {132, 2},
		143: {166, 0},
		144: {166, 1},
		145: {169, 1},
		146: {169, 3},
		147: {169, 5},
		148: {169, 6},
		149: {169, 6},
		150: {169, 5},
		151: {195, 0},
		152: {169, 5},
		153: {169, 4},
		154: {129, 2},
		155: {129, 3},
		156: {130, 0},
		157: {130, 1},
		158: {145, 1},
		159: {145, 2},
		160: {150, 0},
		161: {150, 1},
		162: {160, 1},
		163: {160, 3},
		164: {181, 0},
		165: {181, 1},
		166: {159, 1},
		167: {159, 3},
		168: {149, 2},
		169: {149, 2},
		170: {175, 1},
		171: {175, 3},
		172: {176, 0},
		173: {176, 1},
		174: {177, 0},
		175: {177, 1},
		176: {151, 0},
		177: {161, 3},
		178: {146, 1},
		179: {146, 2},
		180: {164, 0},
		181: {164, 1},
		182: {167, 3},
		183: {167, 4},
		184: {167, 5},
		185: {167, 6},
		186: {167, 6},
		187: {167, 4},
		188: {162, 0},
		189: {167, 4},
		190: {186, 0},
		191: {167, 5},
		192: {168, 0},
		193: {168, 1},
		194: {158, 1},
		195: {158, 4},
		196: {180, 2},
		197: {180, 4},
		198: {153, 2},
		199: {154, 0},
		200: {154, 1},
		201: {155, 1},
		202: {155, 2},
		203: {148, 3},
		204: {148, 2},
		205: {128, 1},
		206: {128, 1},
		207: {128, 1},
		208: {128, 1},
		209: {128, 1},
		210: {128, 1},
		211: {126, 3},
		212: {126, 4},
		213: {126, 3},
		214: {187, 0},
		215: {122, 4},
		216: {198, 1},
		217: {198, 2},
		218: {199, 0},
		219: {199, 1},
		220: {165, 1},
		221: {165, 1},
		222: {123, 2},
		223: {127, 5},
		224: {127, 7},
		225: {127, 5},
		226: {124, 5},
		227: {124, 7},
		228: {124, 9},
		229: {124, 8},
		230: {125, 3},
		231: {125, 2},
		232: {125, 2},
		233: {125, 3},
		234: {216, 1},
		235: {216, 2},
		236: {173, 1},
		237: {173, 1},
		238: {188, 0},
		239: {174, 5},
		240: {201, 1},
		241: {201, 2},
		242: {202, 0},
		243: {189, 0},
		244: {202, 2},
		245: {211, 1},
		246: {138, 1},
		247: {138, 2},
		248: {143, 0},
		249: {143, 1},
		250: {133, 1},
		251: {133, 1},
		252: {133, 3},
		253: {133, 1},
		254: {135, 4},
		255: {134, 4},
		256: {134, 4},
		257: {134, 4},
		258: {203, 1},
		259: {203, 2},
		260: {204, 0},
		261: {204, 1},
		262: {170, 4},
		263: {205, 3},
		264: {206, 0},
		265: {206, 1},
		266: {207, 1},
		267: {131, 3},
		268: {131, 5},
		269: {131, 7},
		270: {131, 5},
		271: {131, 2},
		272: {131, 1},
		273: {131, 3},
		274: {131, 3},
		275: {131, 2},
		276: {131, 3},
		277: {131, 6},
		278: {131, 8},
		279: {131, 2},
		280: {131, 4},
		281: {131, 3},
		282: {136, 1},
		283: {139, 1},
		284: {110, 1},
		285: {118, 1},
		286: {118, 2},
		287: {111, 1},
		288: {111, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 43}:   "invalid empty input",
		yyXError{498, -1}: "expected #endif",
		yyXError{500, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{413, -1}: "expected $end",
		yyXError{415, -1}: "expected $end",
		yyXError{31, -1}:  "expected '('",
		yyXError{347, -1}: "expected '('",
		yyXError{348, -1}: "expected '('",
		yyXError{349, -1}: "expected '('",
		yyXError{351, -1}: "expected '('",
		yyXError{377, -1}: "expected '('",
		yyXError{149, -1}: "expected ')'",
		yyXError{156, -1}: "expected ')'",
		yyXError{163, -1}: "expected ')'",
		yyXError{189, -1}: "expected ')'",
		yyXError{192, -1}: "expected ')'",
		yyXError{195, -1}: "expected ')'",
		yyXError{203, -1}: "expected ')'",
		yyXError{208, -1}: "expected ')'",
		yyXError{214, -1}: "expected ')'",
		yyXError{230, -1}: "expected ')'",
		yyXError{235, -1}: "expected ')'",
		yyXError{276, -1}: "expected ')'",
		yyXError{307, -1}: "expected ')'",
		yyXError{367, -1}: "expected ')'",
		yyXError{373, -1}: "expected ')'",
		yyXError{458, -1}: "expected ')'",
		yyXError{459, -1}: "expected ')'",
		yyXError{467, -1}: "expected ')'",
		yyXError{470, -1}: "expected ')'",
		yyXError{473, -1}: "expected ')'",
		yyXError{294, -1}: "expected ':'",
		yyXError{340, -1}: "expected ':'",
		yyXError{401, -1}: "expected ':'",
		yyXError{315, -1}: "expected ';'",
		yyXError{346, -1}: "expected ';'",
		yyXError{353, -1}: "expected ';'",
		yyXError{354, -1}: "expected ';'",
		yyXError{356, -1}: "expected ';'",
		yyXError{360, -1}: "expected ';'",
		yyXError{363, -1}: "expected ';'",
		yyXError{365, -1}: "expected ';'",
		yyXError{371, -1}: "expected ';'",
		yyXError{380, -1}: "expected ';'",
		yyXError{319, -1}: "expected '='",
		yyXError{168, -1}: "expected '['",
		yyXError{437, -1}: "expected '\\n'",
		yyXError{441, -1}: "expected '\\n'",
		yyXError{445, -1}: "expected '\\n'",
		yyXError{448, -1}: "expected '\\n'",
		yyXError{450, -1}: "expected '\\n'",
		yyXError{477, -1}: "expected '\\n'",
		yyXError{482, -1}: "expected '\\n'",
		yyXError{485, -1}: "expected '\\n'",
		yyXError{492, -1}: "expected '\\n'",
		yyXError{497, -1}: "expected '\\n'",
		yyXError{503, -1}: "expected '\\n'",
		yyXError{174, -1}: "expected ']'",
		yyXError{182, -1}: "expected ']'",
		yyXError{226, -1}: "expected ']'",
		yyXError{253, -1}: "expected ']'",
		yyXError{43, -1}:  "expected '{'",
		yyXError{45, -1}:  "expected '{'",
		yyXError{282, -1}: "expected '{'",
		yyXError{284, -1}: "expected '{'",
		yyXError{262, -1}: "expected '}'",
		yyXError{266, -1}: "expected '}'",
		yyXError{279, -1}: "expected '}'",
		yyXError{341, -1}: "expected '}'",
		yyXError{48, -1}:  "expected CommaOpt or one of [',', '}']",
		yyXError{245, -1}: "expected CommaOpt or one of [',', '}']",
		yyXError{260, -1}: "expected CommaOpt or one of [',', '}']",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{202, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{167, -1}: "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{343, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{320, -1}: "expected compound statement or '{'",
		yyXError{327, -1}: "expected compound statement or '{'",
		yyXError{318, -1}: "expected compound statement or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{51, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{250, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{298, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{339, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{412, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{321, -1}: "expected declaration list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{324, -1}: "expected declaration or one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{362, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{297, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{194, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{247, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{197, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{164, -1}: "expected direct abstract declarator or one of ['(', '[']",
		yyXError{295, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{490, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{496, -1}: "expected endif line or #endif",
		yyXError{422, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{488, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{46, -1}:  "expected enumerator list or identifier",
		yyXError{278, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{74, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{98, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{378, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{382, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{386, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{390, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{61, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{72, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{242, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		yyXError{171, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{225, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{277, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{52, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{63, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{64, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{65, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{66, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{67, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{68, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{69, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{70, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{71, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{97, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{109, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{123, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{124, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{151, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{177, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{183, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{219, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{222, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{175, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{217, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{305, -1}: "expected expression or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{6, -1}:   "expected external declaration or one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{479, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{416, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{76, -1}:  "expected identifier",
		yyXError{77, -1}:  "expected identifier",
		yyXError{211, -1}: "expected identifier",
		yyXError{251, -1}: "expected identifier",
		yyXError{352, -1}: "expected identifier",
		yyXError{424, -1}: "expected identifier",
		yyXError{425, -1}: "expected identifier",
		yyXError{432, -1}: "expected identifier",
		yyXError{454, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{408, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{244, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{258, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{246, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{264, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{406, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{261, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{54, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{55, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{56, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{57, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{58, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{59, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{60, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{73, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{78, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{79, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
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
		yyXError{120, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
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
		yyXError{146, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{150, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{154, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{187, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{243, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{267, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{268, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{269, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{270, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{271, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{272, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{273, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{274, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{275, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{62, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{121, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{125, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{147, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{152, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{306, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{331, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{257, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{170, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{178, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{184, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{220, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{223, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{417, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{418, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{419, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{421, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{428, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{434, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{436, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{439, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{442, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{444, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{446, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{447, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{449, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{451, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{452, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{455, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{461, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{462, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{464, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{469, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{472, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{475, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{476, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{481, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{501, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{502, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{504, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{480, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{484, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{487, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{489, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{494, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{495, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{398, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{410, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{40, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{41, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{42, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{329, -1}: "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{411, -1}: "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{36, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{38, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{172, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{180, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{333, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{334, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{335, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{336, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{337, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{338, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{357, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{358, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{359, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{361, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{369, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{375, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{381, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{385, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{389, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{393, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{395, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{396, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{400, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{403, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{405, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{342, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{344, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{345, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{397, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{248, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{255, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{44, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{283, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
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
		yyXError{280, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{303, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{308, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{309, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{12, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{39, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{310, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{311, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{312, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{313, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{314, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{239, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{240, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{241, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{200, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{201, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{204, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{213, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{215, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{221, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{224, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{227, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{228, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{162, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{238, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{166, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{179, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{181, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{185, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{186, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{188, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{196, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{232, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{236, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{296, -1}: "expected one of ['(', identifier]",
		yyXError{332, -1}: "expected one of [')', ',', ';']",
		yyXError{456, -1}: "expected one of [')', ',', ...]",
		yyXError{466, -1}: "expected one of [')', ',', ...]",
		yyXError{148, -1}: "expected one of [')', ',']",
		yyXError{155, -1}: "expected one of [')', ',']",
		yyXError{165, -1}: "expected one of [')', ',']",
		yyXError{191, -1}: "expected one of [')', ',']",
		yyXError{193, -1}: "expected one of [')', ',']",
		yyXError{198, -1}: "expected one of [')', ',']",
		yyXError{199, -1}: "expected one of [')', ',']",
		yyXError{209, -1}: "expected one of [')', ',']",
		yyXError{210, -1}: "expected one of [')', ',']",
		yyXError{212, -1}: "expected one of [')', ',']",
		yyXError{231, -1}: "expected one of [')', ',']",
		yyXError{379, -1}: "expected one of [')', ',']",
		yyXError{383, -1}: "expected one of [')', ',']",
		yyXError{387, -1}: "expected one of [')', ',']",
		yyXError{391, -1}: "expected one of [')', ',']",
		yyXError{457, -1}: "expected one of [')', ',']",
		yyXError{293, -1}: "expected one of [',', ':', ';']",
		yyXError{122, -1}: "expected one of [',', ':']",
		yyXError{326, -1}: "expected one of [',', ';', '=']",
		yyXError{263, -1}: "expected one of [',', ';', '}']",
		yyXError{290, -1}: "expected one of [',', ';']",
		yyXError{292, -1}: "expected one of [',', ';']",
		yyXError{299, -1}: "expected one of [',', ';']",
		yyXError{302, -1}: "expected one of [',', ';']",
		yyXError{316, -1}: "expected one of [',', ';']",
		yyXError{317, -1}: "expected one of [',', ';']",
		yyXError{407, -1}: "expected one of [',', ';']",
		yyXError{409, -1}: "expected one of [',', ';']",
		yyXError{47, -1}:  "expected one of [',', '=', '}']",
		yyXError{50, -1}:  "expected one of [',', '=', '}']",
		yyXError{153, -1}: "expected one of [',', ']']",
		yyXError{49, -1}:  "expected one of [',', '}']",
		yyXError{53, -1}:  "expected one of [',', '}']",
		yyXError{259, -1}: "expected one of [',', '}']",
		yyXError{265, -1}: "expected one of [',', '}']",
		yyXError{281, -1}: "expected one of [',', '}']",
		yyXError{249, -1}: "expected one of ['.', '=', '[']",
		yyXError{252, -1}: "expected one of ['.', '=', '[']",
		yyXError{254, -1}: "expected one of ['.', '=', '[']",
		yyXError{256, -1}: "expected one of ['.', '=', '[']",
		yyXError{426, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		yyXError{435, -1}: "expected one of ['\\n', ppother]",
		yyXError{438, -1}: "expected one of ['\\n', ppother]",
		yyXError{440, -1}: "expected one of ['\\n', ppother]",
		yyXError{323, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{325, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{33, -1}:  "expected one of ['{', identifier]",
		yyXError{34, -1}:  "expected one of ['{', identifier]",
		yyXError{288, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{291, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{300, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{304, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{465, -1}: "expected one of [..., identifier]",
		yyXError{160, -1}: "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{75, -1}:  "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{328, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{330, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{8, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{366, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{372, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{355, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{364, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{370, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{216, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{205, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{169, -1}: "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{173, -1}: "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{478, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{483, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{486, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{493, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{499, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{206, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{32, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{35, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{322, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{190, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{233, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{234, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{158, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{159, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{427, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{431, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{161, -1}: "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{229, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{207, -1}: "expected parameter type list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{237, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{414, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{453, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{460, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{463, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{468, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{471, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{474, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{157, -1}: "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{350, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{368, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{374, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{384, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{388, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{392, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{394, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{399, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{402, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{404, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{285, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{286, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{287, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{289, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		yyXError{301, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{443, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{420, -1}: "expected token list or ppother",
		yyXError{423, -1}: "expected token list or ppother",
		yyXError{429, -1}: "expected token list or ppother",
		yyXError{430, -1}: "expected token list or ppother",
		yyXError{433, -1}: "expected token list or ppother",
		yyXError{491, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{5, -1}:   "expected translation unit or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{176, -1}: "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{218, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{376, -1}: "expected while",
		yyXError{3, 43}:   "unexpected EOF",
		yyXError{2, 43}:   "unexpected EOF",
		yyXError{4, 43}:   "unexpected EOF",
	}

	yyParseTab = [505][]uint16{
		// 0
		{200: 292, 210: 291, 212: 290, 215: 293},
		{43: 289},
		{80: 288, 288, 96: 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 185: 703},
		{286, 286, 286, 286, 286, 286, 286, 286, 12: 286, 286, 286, 286, 286, 286, 286, 286, 286, 190: 701},
		{21: 284, 284, 284, 25: 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 42: 284, 44: 284, 284, 284, 284, 284, 191: 294},
		// 5
		{21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 108: 299, 113: 318, 321, 317, 298, 119: 296, 300, 297, 137: 331, 173: 329, 330, 216: 295},
		{21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 283, 304, 302, 328, 305, 301, 108: 299, 113: 318, 321, 317, 298, 119: 296, 300, 297, 137: 331, 173: 700, 330},
		{133, 450, 133, 8: 204, 129: 585, 584, 132: 607, 157: 605, 178: 606, 604},
		{208, 208, 208, 8: 208, 208, 208, 208, 21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 108: 299, 113: 318, 321, 317, 298, 119: 600, 300, 297, 147: 603},
		{208, 208, 208, 8: 208, 208, 208, 208, 21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 108: 299, 113: 318, 321, 317, 298, 119: 600, 300, 297, 147: 602},
		// 10
		{208, 208, 208, 8: 208, 208, 208, 208, 21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 108: 299, 113: 318, 321, 317, 298, 119: 600, 300, 297, 147: 601},
		{208, 208, 208, 8: 208, 208, 208, 208, 21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 108: 299, 113: 318, 321, 317, 298, 119: 600, 300, 297, 147: 599},
		{199, 199, 199, 8: 199, 199, 199, 199, 21: 199, 199, 199, 25: 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 42: 199, 44: 199, 199, 199, 199, 199},
		{198, 198, 198, 8: 198, 198, 198, 198, 21: 198, 198, 198, 25: 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 42: 198, 44: 198, 198, 198, 198, 198},
		{197, 197, 197, 8: 197, 197, 197, 197, 21: 197, 197, 197, 25: 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 42: 197, 44: 197, 197, 197, 197, 197},
		// 15
		{196, 196, 196, 8: 196, 196, 196, 196, 21: 196, 196, 196, 25: 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 42: 196, 44: 196, 196, 196, 196, 196},
		{195, 195, 195, 8: 195, 195, 195, 195, 21: 195, 195, 195, 25: 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 42: 195, 44: 195, 195, 195, 195, 195},
		{194, 194, 194, 8: 194, 194, 194, 194, 21: 194, 194, 194, 25: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 44: 194, 194, 194, 194, 194},
		{193, 193, 193, 8: 193, 193, 193, 193, 21: 193, 193, 193, 25: 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 44: 193, 193, 193, 193, 193},
		{192, 192, 192, 8: 192, 192, 192, 192, 21: 192, 192, 192, 25: 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 44: 192, 192, 192, 192, 192},
		// 20
		{191, 191, 191, 8: 191, 191, 191, 191, 21: 191, 191, 191, 25: 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 44: 191, 191, 191, 191, 191},
		{190, 190, 190, 8: 190, 190, 190, 190, 21: 190, 190, 190, 25: 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 44: 190, 190, 190, 190, 190},
		{189, 189, 189, 8: 189, 189, 189, 189, 21: 189, 189, 189, 25: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 44: 189, 189, 189, 189, 189},
		{188, 188, 188, 8: 188, 188, 188, 188, 21: 188, 188, 188, 25: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 44: 188, 188, 188, 188, 188},
		{187, 187, 187, 8: 187, 187, 187, 187, 21: 187, 187, 187, 25: 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 44: 187, 187, 187, 187, 187},
		// 25
		{186, 186, 186, 8: 186, 186, 186, 186, 21: 186, 186, 186, 25: 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 44: 186, 186, 186, 186, 186},
		{185, 185, 185, 8: 185, 185, 185, 185, 21: 185, 185, 185, 25: 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 44: 185, 185, 185, 185, 185},
		{184, 184, 184, 8: 184, 184, 184, 184, 21: 184, 184, 184, 25: 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 44: 184, 184, 184, 184, 184},
		{183, 183, 183, 8: 183, 183, 183, 183, 21: 183, 183, 183, 25: 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 44: 183, 183, 183, 183, 183},
		{182, 182, 182, 8: 182, 182, 182, 182, 21: 182, 182, 182, 25: 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 44: 182, 182, 182, 182, 182},
		// 30
		{181, 181, 181, 8: 181, 181, 181, 181, 21: 181, 181, 181, 25: 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 44: 181, 181, 181, 181, 181},
		{594},
		{2: 572, 53: 115, 177: 571},
		{2: 174, 53: 174},
		{2: 173, 53: 173},
		// 35
		{2: 333, 53: 115, 177: 332},
		{151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 25: 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 44: 151, 151, 151, 151, 151, 51: 151},
		{150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 25: 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 44: 150, 150, 150, 150, 150, 51: 150},
		{149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 25: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 44: 149, 149, 149, 149, 149, 51: 149},
		{148, 148, 148, 8: 148, 148, 148, 148, 21: 148, 148, 148, 25: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 42: 148, 44: 148, 148, 148, 148, 148},
		// 40
		{21: 55, 55, 55, 25: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 42: 55, 55, 55, 55, 55, 55, 55},
		{21: 53, 53, 53, 25: 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 42: 53, 53, 53, 53, 53, 53, 53},
		{21: 52, 52, 52, 25: 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 42: 52, 52, 52, 52, 52, 52, 52},
		{53: 158, 194: 334},
		{156, 156, 156, 8: 156, 156, 156, 156, 21: 156, 156, 156, 25: 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 44: 156, 156, 156, 156, 156, 53: 114},
		// 45
		{53: 335},
		{2: 336, 171: 339, 338, 208: 337},
		{10: 282, 24: 282, 49: 282},
		{10: 567, 24: 160, 152: 568},
		{10: 155, 24: 155},
		// 50
		{10: 153, 24: 153, 49: 340},
		{215, 215, 215, 215, 215, 215, 215, 215, 12: 215, 215, 215, 215, 215, 215, 215, 215, 215, 140: 341, 342},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 351},
		{10: 152, 24: 152},
		{277, 277, 3: 277, 277, 277, 277, 277, 277, 277, 277, 277, 24: 277, 41: 277, 43: 277, 49: 277, 51: 277, 277, 54: 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277},
		// 55
		{276, 276, 3: 276, 276, 276, 276, 276, 276, 276, 276, 276, 24: 276, 41: 276, 43: 276, 49: 276, 51: 276, 276, 54: 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276},
		{275, 275, 3: 275, 275, 275, 275, 275, 275, 275, 275, 275, 24: 275, 41: 275, 43: 275, 49: 275, 51: 275, 275, 54: 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275},
		{274, 274, 3: 274, 274, 274, 274, 274, 274, 274, 274, 274, 24: 274, 41: 274, 43: 274, 49: 274, 51: 274, 274, 54: 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274},
		{273, 273, 3: 273, 273, 273, 273, 273, 273, 273, 273, 273, 24: 273, 41: 273, 43: 273, 49: 273, 51: 273, 273, 54: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273},
		{272, 272, 3: 272, 272, 272, 272, 272, 272, 272, 272, 272, 24: 272, 41: 272, 43: 272, 49: 272, 51: 272, 272, 54: 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272},
		// 60
		{271, 271, 3: 271, 271, 271, 271, 271, 271, 271, 271, 271, 24: 271, 41: 271, 43: 271, 49: 271, 51: 271, 271, 54: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 113, 113, 113, 25: 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 50: 410, 109: 444, 151: 446, 161: 565},
		{364, 369, 3: 382, 372, 373, 368, 367, 214, 10: 214, 363, 24: 214, 41: 214, 43: 214, 49: 388, 51: 214, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 564},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 563},
		// 65
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 562},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 476},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 561},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 560},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 559},
		// 70
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 558},
		{361, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 362},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 113, 113, 113, 25: 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 50: 410, 109: 444, 151: 446, 161: 445},
		{364, 254, 3: 254, 254, 254, 368, 367, 254, 254, 254, 363, 24: 254, 41: 254, 43: 254, 49: 254, 51: 254, 365, 54: 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 366, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 109: 442},
		// 75
		{350, 355, 343, 354, 356, 357, 353, 352, 9: 279, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 436, 196: 437, 438},
		{2: 435},
		{2: 434},
		{265, 265, 3: 265, 265, 265, 265, 265, 265, 265, 265, 265, 24: 265, 41: 265, 43: 265, 49: 265, 51: 265, 265, 54: 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265},
		{264, 264, 3: 264, 264, 264, 264, 264, 264, 264, 264, 264, 24: 264, 41: 264, 43: 264, 49: 264, 51: 264, 264, 54: 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264},
		// 80
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 433},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 432},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 431},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 430},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 429},
		// 85
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 428},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 427},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 426},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 425},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 424},
		// 90
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 423},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 422},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 421},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 420},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 419},
		// 95
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 418},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 417},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 416},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 109: 411},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 409},
		// 100
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 408},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 407},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 406},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 405},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 404},
		// 105
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 403},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 402},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 401},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 400},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 399},
		// 110
		{364, 369, 3: 382, 372, 373, 368, 367, 222, 222, 222, 363, 24: 222, 41: 222, 43: 222, 49: 388, 51: 222, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{364, 369, 3: 382, 372, 373, 368, 367, 223, 223, 223, 363, 24: 223, 41: 223, 43: 223, 49: 388, 51: 223, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{364, 369, 3: 382, 372, 373, 368, 367, 224, 224, 224, 363, 24: 224, 41: 224, 43: 224, 49: 388, 51: 224, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{364, 369, 3: 382, 372, 373, 368, 367, 225, 225, 225, 363, 24: 225, 41: 225, 43: 225, 49: 388, 51: 225, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{364, 369, 3: 382, 372, 373, 368, 367, 226, 226, 226, 363, 24: 226, 41: 226, 43: 226, 49: 388, 51: 226, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		// 115
		{364, 369, 3: 382, 372, 373, 368, 367, 227, 227, 227, 363, 24: 227, 41: 227, 43: 227, 49: 388, 51: 227, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{364, 369, 3: 382, 372, 373, 368, 367, 228, 228, 228, 363, 24: 228, 41: 228, 43: 228, 49: 388, 51: 228, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{364, 369, 3: 382, 372, 373, 368, 367, 229, 229, 229, 363, 24: 229, 41: 229, 43: 229, 49: 388, 51: 229, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{364, 369, 3: 382, 372, 373, 368, 367, 230, 230, 230, 363, 24: 230, 41: 230, 43: 230, 49: 388, 51: 230, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{364, 369, 3: 382, 372, 373, 368, 367, 231, 231, 231, 363, 24: 231, 41: 231, 43: 231, 49: 388, 51: 231, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		// 120
		{364, 369, 3: 382, 372, 373, 368, 367, 232, 232, 232, 363, 24: 232, 41: 232, 43: 232, 49: 388, 51: 232, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{364, 369, 3: 382, 372, 373, 368, 367, 219, 219, 219, 363, 41: 219, 49: 388, 51: 219, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{10: 413, 41: 412},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 415},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 414},
		// 125
		{364, 369, 3: 382, 372, 373, 368, 367, 218, 218, 218, 363, 41: 218, 49: 388, 51: 218, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{364, 369, 3: 382, 372, 373, 368, 367, 233, 233, 233, 363, 24: 233, 41: 233, 43: 233, 49: 233, 51: 233, 365, 54: 371, 370, 376, 377, 387, 383, 384, 233, 385, 233, 366, 233, 380, 379, 378, 374, 233, 233, 233, 381, 233, 386, 375, 233, 233, 233},
		{364, 369, 3: 382, 372, 373, 368, 367, 234, 234, 234, 363, 24: 234, 41: 234, 43: 234, 49: 234, 51: 234, 365, 54: 371, 370, 376, 377, 234, 383, 384, 234, 385, 234, 366, 234, 380, 379, 378, 374, 234, 234, 234, 381, 234, 234, 375, 234, 234, 234},
		{364, 369, 3: 382, 372, 373, 368, 367, 235, 235, 235, 363, 24: 235, 41: 235, 43: 235, 49: 235, 51: 235, 365, 54: 371, 370, 376, 377, 235, 383, 384, 235, 235, 235, 366, 235, 380, 379, 378, 374, 235, 235, 235, 381, 235, 235, 375, 235, 235, 235},
		{364, 369, 3: 382, 372, 373, 368, 367, 236, 236, 236, 363, 24: 236, 41: 236, 43: 236, 49: 236, 51: 236, 365, 54: 371, 370, 376, 377, 236, 383, 236, 236, 236, 236, 366, 236, 380, 379, 378, 374, 236, 236, 236, 381, 236, 236, 375, 236, 236, 236},
		// 130
		{364, 369, 3: 382, 372, 373, 368, 367, 237, 237, 237, 363, 24: 237, 41: 237, 43: 237, 49: 237, 51: 237, 365, 54: 371, 370, 376, 377, 237, 237, 237, 237, 237, 237, 366, 237, 380, 379, 378, 374, 237, 237, 237, 381, 237, 237, 375, 237, 237, 237},
		{364, 369, 3: 238, 372, 373, 368, 367, 238, 238, 238, 363, 24: 238, 41: 238, 43: 238, 49: 238, 51: 238, 365, 54: 371, 370, 376, 377, 238, 238, 238, 238, 238, 238, 366, 238, 380, 379, 378, 374, 238, 238, 238, 381, 238, 238, 375, 238, 238, 238},
		{364, 369, 3: 239, 372, 373, 368, 367, 239, 239, 239, 363, 24: 239, 41: 239, 43: 239, 49: 239, 51: 239, 365, 54: 371, 370, 376, 377, 239, 239, 239, 239, 239, 239, 366, 239, 239, 379, 378, 374, 239, 239, 239, 239, 239, 239, 375, 239, 239, 239},
		{364, 369, 3: 240, 372, 373, 368, 367, 240, 240, 240, 363, 24: 240, 41: 240, 43: 240, 49: 240, 51: 240, 365, 54: 371, 370, 376, 377, 240, 240, 240, 240, 240, 240, 366, 240, 240, 379, 378, 374, 240, 240, 240, 240, 240, 240, 375, 240, 240, 240},
		{364, 369, 3: 241, 372, 373, 368, 367, 241, 241, 241, 363, 24: 241, 41: 241, 43: 241, 49: 241, 51: 241, 365, 54: 371, 370, 241, 241, 241, 241, 241, 241, 241, 241, 366, 241, 241, 241, 241, 374, 241, 241, 241, 241, 241, 241, 375, 241, 241, 241},
		// 135
		{364, 369, 3: 242, 372, 373, 368, 367, 242, 242, 242, 363, 24: 242, 41: 242, 43: 242, 49: 242, 51: 242, 365, 54: 371, 370, 242, 242, 242, 242, 242, 242, 242, 242, 366, 242, 242, 242, 242, 374, 242, 242, 242, 242, 242, 242, 375, 242, 242, 242},
		{364, 369, 3: 243, 372, 373, 368, 367, 243, 243, 243, 363, 24: 243, 41: 243, 43: 243, 49: 243, 51: 243, 365, 54: 371, 370, 243, 243, 243, 243, 243, 243, 243, 243, 366, 243, 243, 243, 243, 374, 243, 243, 243, 243, 243, 243, 375, 243, 243, 243},
		{364, 369, 3: 244, 372, 373, 368, 367, 244, 244, 244, 363, 24: 244, 41: 244, 43: 244, 49: 244, 51: 244, 365, 54: 371, 370, 244, 244, 244, 244, 244, 244, 244, 244, 366, 244, 244, 244, 244, 374, 244, 244, 244, 244, 244, 244, 375, 244, 244, 244},
		{364, 369, 3: 245, 372, 373, 368, 367, 245, 245, 245, 363, 24: 245, 41: 245, 43: 245, 49: 245, 51: 245, 365, 54: 371, 370, 245, 245, 245, 245, 245, 245, 245, 245, 366, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245},
		{364, 369, 3: 246, 372, 373, 368, 367, 246, 246, 246, 363, 24: 246, 41: 246, 43: 246, 49: 246, 51: 246, 365, 54: 371, 370, 246, 246, 246, 246, 246, 246, 246, 246, 366, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246},
		// 140
		{364, 369, 3: 247, 247, 247, 368, 367, 247, 247, 247, 363, 24: 247, 41: 247, 43: 247, 49: 247, 51: 247, 365, 54: 371, 370, 247, 247, 247, 247, 247, 247, 247, 247, 366, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247},
		{364, 369, 3: 248, 248, 248, 368, 367, 248, 248, 248, 363, 24: 248, 41: 248, 43: 248, 49: 248, 51: 248, 365, 54: 371, 370, 248, 248, 248, 248, 248, 248, 248, 248, 366, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248},
		{364, 249, 3: 249, 249, 249, 368, 367, 249, 249, 249, 363, 24: 249, 41: 249, 43: 249, 49: 249, 51: 249, 365, 54: 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 366, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249},
		{364, 250, 3: 250, 250, 250, 368, 367, 250, 250, 250, 363, 24: 250, 41: 250, 43: 250, 49: 250, 51: 250, 365, 54: 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 366, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250},
		{364, 251, 3: 251, 251, 251, 368, 367, 251, 251, 251, 363, 24: 251, 41: 251, 43: 251, 49: 251, 51: 251, 365, 54: 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 366, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251},
		// 145
		{266, 266, 3: 266, 266, 266, 266, 266, 266, 266, 266, 266, 24: 266, 41: 266, 43: 266, 49: 266, 51: 266, 266, 54: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266},
		{267, 267, 3: 267, 267, 267, 267, 267, 267, 267, 267, 267, 24: 267, 41: 267, 43: 267, 49: 267, 51: 267, 267, 54: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267},
		{364, 369, 3: 382, 372, 373, 368, 367, 9: 281, 281, 363, 49: 388, 52: 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{9: 278, 440},
		{9: 439},
		// 150
		{268, 268, 3: 268, 268, 268, 268, 268, 268, 268, 268, 268, 24: 268, 41: 268, 43: 268, 49: 268, 51: 268, 268, 54: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 441},
		{364, 369, 3: 382, 372, 373, 368, 367, 9: 280, 280, 363, 49: 388, 52: 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{10: 413, 51: 443},
		{269, 269, 3: 269, 269, 269, 269, 269, 269, 269, 269, 269, 24: 269, 41: 269, 43: 269, 49: 269, 51: 269, 269, 54: 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269},
		// 155
		{9: 557, 413},
		{9: 531},
		{21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 108: 448, 113: 318, 321, 317, 447, 144: 449},
		{166, 166, 166, 8: 166, 166, 11: 166, 21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 166, 108: 448, 113: 318, 321, 317, 447, 144: 529, 182: 530},
		{166, 166, 166, 8: 166, 166, 11: 166, 21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 166, 108: 448, 113: 318, 321, 317, 447, 144: 529, 182: 528},
		// 160
		{133, 450, 9: 109, 11: 133, 129: 451, 453, 146: 454, 164: 452},
		{129, 129, 129, 9: 129, 129, 129, 21: 325, 326, 327, 108: 461, 145: 465, 150: 526},
		{132, 2: 132, 9: 111, 111, 132},
		{9: 112},
		{456, 11: 97, 167: 455, 457},
		// 165
		{9: 108, 108},
		{522, 9: 110, 110, 96},
		{133, 450, 9: 101, 11: 133, 21: 101, 101, 101, 25: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 42: 101, 44: 101, 101, 101, 101, 101, 129: 451, 453, 146: 478, 162: 479},
		{11: 458},
		{350, 460, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 325, 326, 327, 42: 464, 50: 459, 221, 108: 461, 145: 462, 156: 463},
		// 170
		{364, 369, 3: 382, 372, 373, 368, 367, 11: 363, 49: 388, 51: 220, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 476, 477},
		{131, 131, 131, 131, 131, 131, 131, 131, 9: 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 42: 131, 51: 131},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 325, 326, 327, 42: 472, 50: 459, 221, 108: 469, 156: 471},
		{51: 470},
		// 175
		{129, 129, 129, 129, 129, 129, 129, 129, 12: 129, 129, 129, 129, 129, 129, 129, 129, 129, 325, 326, 327, 108: 461, 145: 465, 150: 466},
		{128, 128, 128, 128, 128, 128, 128, 128, 9: 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 325, 326, 327, 108: 469},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 467},
		{364, 369, 3: 382, 372, 373, 368, 367, 11: 363, 49: 388, 51: 468, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{104, 9: 104, 104, 104},
		// 180
		{130, 130, 130, 130, 130, 130, 130, 130, 9: 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 42: 130, 51: 130},
		{106, 9: 106, 106, 106},
		{51: 475},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 473},
		{364, 369, 3: 382, 372, 373, 368, 367, 11: 363, 49: 388, 51: 474, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		// 185
		{103, 9: 103, 103, 103},
		{105, 9: 105, 105, 105},
		{364, 259, 3: 259, 259, 259, 368, 367, 259, 259, 259, 363, 24: 259, 41: 259, 43: 259, 49: 259, 51: 259, 365, 54: 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 366, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259},
		{102, 9: 102, 102, 102},
		{9: 521},
		// 190
		{9: 125, 21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 108: 299, 113: 318, 321, 317, 298, 119: 483, 300, 297, 149: 482, 159: 480, 481, 181: 484},
		{9: 127, 518},
		{9: 124},
		{9: 123, 123},
		{133, 450, 133, 9: 109, 109, 133, 129: 451, 486, 132: 487, 146: 454, 164: 488},
		// 195
		{9: 485},
		{100, 9: 100, 100, 100},
		{491, 2: 490, 11: 97, 167: 455, 457, 489},
		{9: 121, 121},
		{9: 120, 120},
		// 200
		{495, 8: 147, 147, 147, 494, 21: 147, 147, 147, 25: 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 44: 147, 147, 147, 147, 147, 147, 53: 147},
		{144, 8: 144, 144, 144, 144, 21: 144, 144, 144, 25: 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 44: 144, 144, 144, 144, 144, 144, 53: 144},
		{133, 450, 133, 9: 101, 11: 133, 21: 101, 101, 101, 25: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 42: 101, 44: 101, 101, 101, 101, 101, 129: 451, 486, 132: 492, 146: 478, 162: 479},
		{9: 493},
		{143, 8: 143, 143, 143, 143, 21: 143, 143, 143, 25: 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 44: 143, 143, 143, 143, 143, 143, 53: 143},
		// 205
		{129, 129, 129, 129, 129, 129, 129, 129, 12: 129, 129, 129, 129, 129, 129, 129, 129, 129, 325, 326, 327, 42: 506, 51: 129, 108: 461, 145: 507, 150: 505},
		{2: 498, 9: 117, 21: 138, 138, 138, 25: 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 42: 138, 44: 138, 138, 138, 138, 138, 175: 499, 497, 195: 496},
		{21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 108: 299, 113: 318, 321, 317, 298, 119: 483, 300, 297, 149: 482, 159: 480, 503},
		{9: 502},
		{9: 119, 119},
		// 210
		{9: 116, 500},
		{2: 501},
		{9: 118, 118},
		{136, 8: 136, 136, 136, 136, 21: 136, 136, 136, 25: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 44: 136, 136, 136, 136, 136, 136, 53: 136},
		{9: 504},
		// 215
		{137, 8: 137, 137, 137, 137, 21: 137, 137, 137, 25: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 44: 137, 137, 137, 137, 137, 137, 53: 137},
		{350, 514, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 459, 221, 156: 515},
		{129, 129, 129, 129, 129, 129, 129, 129, 12: 129, 129, 129, 129, 129, 129, 129, 129, 129, 325, 326, 327, 108: 461, 145: 465, 150: 511},
		{128, 128, 128, 128, 128, 128, 128, 128, 12: 128, 128, 128, 128, 128, 128, 128, 128, 128, 325, 326, 327, 42: 508, 51: 128, 108: 469},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 509},
		// 220
		{364, 369, 3: 382, 372, 373, 368, 367, 11: 363, 49: 388, 51: 510, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{140, 8: 140, 140, 140, 140, 21: 140, 140, 140, 25: 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 44: 140, 140, 140, 140, 140, 140, 53: 140},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 512},
		{364, 369, 3: 382, 372, 373, 368, 367, 11: 363, 49: 388, 51: 513, 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{141, 8: 141, 141, 141, 141, 21: 141, 141, 141, 25: 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 44: 141, 141, 141, 141, 141, 141, 53: 141},
		// 225
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 476, 517},
		{51: 516},
		{142, 8: 142, 142, 142, 142, 21: 142, 142, 142, 25: 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 44: 142, 142, 142, 142, 142, 142, 53: 142},
		{139, 8: 139, 139, 139, 139, 21: 139, 139, 139, 25: 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 44: 139, 139, 139, 139, 139, 139, 53: 139},
		{21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 108: 299, 113: 318, 321, 317, 298, 119: 483, 300, 297, 142: 519, 149: 520},
		// 230
		{9: 126},
		{9: 122, 122},
		{107, 9: 107, 107, 107},
		{9: 99, 21: 99, 99, 99, 25: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 42: 99, 44: 99, 99, 99, 99, 99, 186: 523},
		{9: 125, 21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 108: 299, 113: 318, 321, 317, 298, 119: 483, 300, 297, 149: 482, 159: 480, 481, 181: 524},
		// 235
		{9: 525},
		{98, 9: 98, 98, 98},
		{135, 450, 135, 9: 135, 135, 135, 129: 527},
		{134, 2: 134, 9: 134, 134, 134},
		{167, 167, 167, 8: 167, 167, 11: 167, 41: 167},
		// 240
		{165, 165, 165, 8: 165, 165, 11: 165, 41: 165},
		{168, 168, 168, 8: 168, 168, 11: 168, 41: 168},
		{350, 253, 343, 253, 253, 253, 353, 352, 253, 253, 253, 253, 359, 358, 344, 345, 346, 347, 348, 360, 349, 24: 253, 41: 253, 43: 253, 49: 253, 532, 253, 253, 533, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253},
		{364, 252, 3: 252, 252, 252, 368, 367, 252, 252, 252, 363, 24: 252, 41: 252, 43: 252, 49: 252, 51: 252, 365, 54: 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 366, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252},
		{90, 90, 90, 90, 90, 90, 90, 90, 11: 539, 90, 90, 90, 90, 90, 90, 90, 90, 90, 52: 540, 90, 148: 538, 153: 537, 535, 536, 180: 534},
		// 245
		{10: 550, 24: 160, 152: 555},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 546, 53: 547, 158: 548},
		{11: 539, 49: 544, 52: 540, 148: 545},
		{89, 89, 89, 89, 89, 89, 89, 89, 12: 89, 89, 89, 89, 89, 89, 89, 89, 89, 53: 89},
		{11: 88, 49: 88, 52: 88},
		// 250
		{215, 215, 215, 215, 215, 215, 215, 215, 12: 215, 215, 215, 215, 215, 215, 215, 215, 215, 140: 341, 542},
		{2: 541},
		{11: 85, 49: 85, 52: 85},
		{51: 543},
		{11: 86, 49: 86, 52: 86},
		// 255
		{91, 91, 91, 91, 91, 91, 91, 91, 12: 91, 91, 91, 91, 91, 91, 91, 91, 91, 53: 91},
		{11: 87, 49: 87, 52: 87},
		{364, 369, 3: 382, 372, 373, 368, 367, 95, 10: 95, 363, 24: 95, 49: 388, 52: 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{90, 90, 90, 90, 90, 90, 90, 90, 11: 539, 90, 90, 90, 90, 90, 90, 90, 90, 90, 52: 540, 90, 148: 538, 153: 537, 535, 536, 180: 549},
		{10: 93, 24: 93},
		// 260
		{10: 550, 24: 160, 152: 551},
		{90, 90, 90, 90, 90, 90, 90, 90, 11: 539, 90, 90, 90, 90, 90, 90, 90, 90, 90, 24: 159, 52: 540, 90, 148: 538, 153: 537, 553, 536},
		{24: 552},
		{8: 94, 10: 94, 24: 94},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 546, 53: 547, 158: 554},
		// 265
		{10: 92, 24: 92},
		{24: 556},
		{263, 263, 3: 263, 263, 263, 263, 263, 263, 263, 263, 263, 24: 263, 41: 263, 43: 263, 49: 263, 51: 263, 263, 54: 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263},
		{270, 270, 3: 270, 270, 270, 270, 270, 270, 270, 270, 270, 24: 270, 41: 270, 43: 270, 49: 270, 51: 270, 270, 54: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270},
		{364, 255, 3: 255, 255, 255, 368, 367, 255, 255, 255, 363, 24: 255, 41: 255, 43: 255, 49: 255, 51: 255, 365, 54: 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 366, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		// 270
		{364, 256, 3: 256, 256, 256, 368, 367, 256, 256, 256, 363, 24: 256, 41: 256, 43: 256, 49: 256, 51: 256, 365, 54: 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 366, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256},
		{364, 257, 3: 257, 257, 257, 368, 367, 257, 257, 257, 363, 24: 257, 41: 257, 43: 257, 49: 257, 51: 257, 365, 54: 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 366, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257},
		{364, 258, 3: 258, 258, 258, 368, 367, 258, 258, 258, 363, 24: 258, 41: 258, 43: 258, 49: 258, 51: 258, 365, 54: 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 366, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258},
		{364, 260, 3: 260, 260, 260, 368, 367, 260, 260, 260, 363, 24: 260, 41: 260, 43: 260, 49: 260, 51: 260, 365, 54: 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 366, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260},
		{364, 261, 3: 261, 261, 261, 368, 367, 261, 261, 261, 363, 24: 261, 41: 261, 43: 261, 49: 261, 51: 261, 365, 54: 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 366, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261},
		// 275
		{364, 262, 3: 262, 262, 262, 368, 367, 262, 262, 262, 363, 24: 262, 41: 262, 43: 262, 49: 262, 51: 262, 365, 54: 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 366, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262},
		{9: 566},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 532, 53: 533},
		{2: 336, 24: 159, 171: 339, 570},
		{24: 569},
		// 280
		{157, 157, 157, 8: 157, 157, 157, 157, 21: 157, 157, 157, 25: 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 44: 157, 157, 157, 157, 157},
		{10: 154, 24: 154},
		{53: 178, 192: 573},
		{175, 175, 175, 8: 175, 175, 175, 175, 21: 175, 175, 175, 25: 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 44: 175, 175, 175, 175, 175, 53: 114},
		{53: 574},
		// 285
		{21: 177, 177, 177, 25: 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 193: 575},
		{21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 108: 448, 113: 318, 321, 317, 447, 144: 578, 183: 577, 213: 576},
		{21: 325, 326, 327, 592, 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 108: 448, 113: 318, 321, 317, 447, 144: 578, 183: 593},
		{21: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172},
		{133, 450, 133, 8: 580, 41: 146, 129: 585, 584, 132: 582, 166: 583, 184: 581, 214: 579},
		// 290
		{8: 589, 10: 590},
		{21: 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169},
		{8: 164, 10: 164},
		{8: 162, 10: 162, 41: 145},
		{41: 587},
		// 295
		{586, 2: 490, 169: 489},
		{132, 2: 132},
		{133, 450, 133, 129: 585, 584, 132: 492},
		{215, 215, 215, 215, 215, 215, 215, 215, 12: 215, 215, 215, 215, 215, 215, 215, 215, 215, 140: 341, 588},
		{8: 161, 10: 161},
		// 300
		{21: 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170},
		{133, 450, 133, 41: 146, 129: 585, 584, 132: 582, 166: 583, 184: 591},
		{8: 163, 10: 163},
		{176, 176, 176, 8: 176, 176, 176, 176, 21: 176, 176, 176, 25: 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 44: 176, 176, 176, 176, 176},
		{21: 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171},
		// 305
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 113, 113, 113, 25: 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 50: 595, 151: 446, 161: 596},
		{364, 369, 3: 382, 372, 373, 368, 367, 9: 598, 11: 363, 49: 388, 52: 365, 54: 371, 370, 376, 377, 387, 383, 384, 392, 385, 396, 366, 390, 380, 379, 378, 374, 394, 391, 389, 381, 398, 386, 375, 395, 393, 397},
		{9: 597},
		{179, 179, 179, 8: 179, 179, 179, 179, 21: 179, 179, 179, 25: 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 44: 179, 179, 179, 179, 179},
		{180, 180, 180, 8: 180, 180, 180, 180, 21: 180, 180, 180, 25: 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 44: 180, 180, 180, 180, 180},
		// 310
		{209, 209, 209, 8: 209, 209, 209, 209},
		{207, 207, 207, 8: 207, 207, 207, 207},
		{210, 210, 210, 8: 210, 210, 210, 210},
		{211, 211, 211, 8: 211, 211, 211, 211},
		{212, 212, 212, 8: 212, 212, 212, 212},
		// 315
		{8: 699},
		{8: 206, 10: 206},
		{8: 203, 10: 697},
		{8: 202, 10: 202, 21: 46, 46, 46, 25: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 42: 46, 44: 46, 46, 46, 46, 46, 201, 53: 47, 163: 608, 189: 610, 202: 609},
		{49: 695},
		// 320
		{53: 51, 188: 616},
		{21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 108: 299, 113: 318, 321, 317, 298, 119: 611, 300, 297, 137: 612, 201: 613},
		{133, 450, 133, 8: 204, 129: 585, 584, 132: 615, 157: 605, 178: 606, 604},
		{21: 49, 49, 49, 25: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 42: 49, 44: 49, 49, 49, 49, 49, 53: 49},
		{21: 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 53: 45, 108: 299, 113: 318, 321, 317, 298, 119: 611, 300, 297, 137: 614},
		// 325
		{21: 48, 48, 48, 25: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 42: 48, 44: 48, 48, 48, 48, 48, 53: 48},
		{8: 202, 10: 202, 49: 201, 163: 608},
		{53: 617, 122: 618},
		{75, 75, 75, 75, 75, 75, 75, 75, 75, 12: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 42: 75, 44: 75, 75, 75, 75, 75, 53: 75, 84: 75, 75, 75, 75, 75, 75, 75, 75, 75, 94: 75, 75, 187: 619},
		{21: 50, 50, 50, 25: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 42: 50, 50, 50, 50, 50, 50, 50},
		// 330
		{350, 355, 620, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 325, 326, 327, 71, 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 50: 410, 53: 617, 84: 638, 643, 628, 642, 629, 639, 640, 641, 636, 94: 644, 637, 108: 299, 621, 113: 318, 321, 317, 298, 635, 119: 611, 300, 297, 623, 624, 626, 627, 622, 625, 634, 137: 633, 165: 631, 198: 632, 630},
		{277, 277, 3: 277, 277, 277, 277, 277, 277, 10: 277, 277, 41: 693, 49: 277, 52: 277, 54: 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277},
		{8: 216, 216, 413},
		{84, 84, 84, 84, 84, 84, 84, 84, 84, 12: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 42: 84, 44: 84, 84, 84, 84, 84, 53: 84, 84: 84, 84, 84, 84, 84, 84, 84, 84, 84, 94: 84, 84, 112: 84},
		{83, 83, 83, 83, 83, 83, 83, 83, 83, 12: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 42: 83, 44: 83, 83, 83, 83, 83, 53: 83, 84: 83, 83, 83, 83, 83, 83, 83, 83, 83, 94: 83, 83, 112: 83},
		// 335
		{82, 82, 82, 82, 82, 82, 82, 82, 82, 12: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 42: 82, 44: 82, 82, 82, 82, 82, 53: 82, 84: 82, 82, 82, 82, 82, 82, 82, 82, 82, 94: 82, 82, 112: 82},
		{81, 81, 81, 81, 81, 81, 81, 81, 81, 12: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 42: 81, 44: 81, 81, 81, 81, 81, 53: 81, 84: 81, 81, 81, 81, 81, 81, 81, 81, 81, 94: 81, 81, 112: 81},
		{80, 80, 80, 80, 80, 80, 80, 80, 80, 12: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 42: 80, 44: 80, 80, 80, 80, 80, 53: 80, 84: 80, 80, 80, 80, 80, 80, 80, 80, 80, 94: 80, 80, 112: 80},
		{79, 79, 79, 79, 79, 79, 79, 79, 79, 12: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 42: 79, 44: 79, 79, 79, 79, 79, 53: 79, 84: 79, 79, 79, 79, 79, 79, 79, 79, 79, 94: 79, 79, 112: 79},
		{215, 215, 215, 215, 215, 215, 215, 215, 12: 215, 215, 215, 215, 215, 215, 215, 215, 215, 140: 341, 690},
		// 340
		{41: 688},
		{24: 687},
		{73, 73, 73, 73, 73, 73, 73, 73, 73, 12: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 42: 73, 44: 73, 73, 73, 73, 73, 53: 73, 84: 73, 73, 73, 73, 73, 73, 73, 73, 73, 94: 73, 73},
		{350, 355, 620, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 325, 326, 327, 70, 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 50: 410, 53: 617, 84: 638, 643, 628, 642, 629, 639, 640, 641, 636, 94: 644, 637, 108: 299, 621, 113: 318, 321, 317, 298, 635, 119: 611, 300, 297, 623, 624, 626, 627, 622, 625, 634, 137: 633, 165: 686},
		{69, 69, 69, 69, 69, 69, 69, 69, 69, 12: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 42: 69, 44: 69, 69, 69, 69, 69, 53: 69, 84: 69, 69, 69, 69, 69, 69, 69, 69, 69, 94: 69, 69},
		// 345
		{68, 68, 68, 68, 68, 68, 68, 68, 68, 12: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 42: 68, 44: 68, 68, 68, 68, 68, 53: 68, 84: 68, 68, 68, 68, 68, 68, 68, 68, 68, 94: 68, 68},
		{8: 685},
		{679},
		{675},
		{671},
		// 350
		{350, 355, 620, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 53: 617, 84: 638, 643, 628, 642, 629, 639, 640, 641, 636, 94: 644, 637, 109: 621, 117: 635, 122: 623, 624, 626, 627, 622, 625, 665},
		{651},
		{2: 649},
		{8: 648},
		{8: 647},
		// 355
		{350, 355, 343, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 109: 621, 117: 645},
		{8: 646},
		{56, 56, 56, 56, 56, 56, 56, 56, 56, 12: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 42: 56, 44: 56, 56, 56, 56, 56, 53: 56, 84: 56, 56, 56, 56, 56, 56, 56, 56, 56, 94: 56, 56, 112: 56},
		{57, 57, 57, 57, 57, 57, 57, 57, 57, 12: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 42: 57, 44: 57, 57, 57, 57, 57, 53: 57, 84: 57, 57, 57, 57, 57, 57, 57, 57, 57, 94: 57, 57, 112: 57},
		{58, 58, 58, 58, 58, 58, 58, 58, 58, 12: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 42: 58, 44: 58, 58, 58, 58, 58, 53: 58, 84: 58, 58, 58, 58, 58, 58, 58, 58, 58, 94: 58, 58, 112: 58},
		// 360
		{8: 650},
		{59, 59, 59, 59, 59, 59, 59, 59, 59, 12: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 42: 59, 44: 59, 59, 59, 59, 59, 53: 59, 84: 59, 59, 59, 59, 59, 59, 59, 59, 59, 94: 59, 59, 112: 59},
		{350, 355, 343, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 325, 326, 327, 25: 315, 307, 316, 312, 324, 311, 309, 310, 308, 313, 322, 319, 320, 323, 314, 306, 42: 303, 44: 304, 302, 328, 305, 301, 50: 410, 108: 299, 621, 113: 318, 321, 317, 298, 652, 119: 611, 300, 297, 137: 653},
		{8: 659},
		{350, 355, 343, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 109: 621, 117: 654},
		// 365
		{8: 655},
		{350, 355, 343, 354, 356, 357, 353, 352, 9: 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 109: 621, 117: 656},
		{9: 657},
		{350, 355, 620, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 53: 617, 84: 638, 643, 628, 642, 629, 639, 640, 641, 636, 94: 644, 637, 109: 621, 117: 635, 122: 623, 624, 626, 627, 622, 625, 658},
		{60, 60, 60, 60, 60, 60, 60, 60, 60, 12: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 42: 60, 44: 60, 60, 60, 60, 60, 53: 60, 84: 60, 60, 60, 60, 60, 60, 60, 60, 60, 94: 60, 60, 112: 60},
		// 370
		{350, 355, 343, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 109: 621, 117: 660},
		{8: 661},
		{350, 355, 343, 354, 356, 357, 353, 352, 9: 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 109: 621, 117: 662},
		{9: 663},
		{350, 355, 620, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 53: 617, 84: 638, 643, 628, 642, 629, 639, 640, 641, 636, 94: 644, 637, 109: 621, 117: 635, 122: 623, 624, 626, 627, 622, 625, 664},
		// 375
		{61, 61, 61, 61, 61, 61, 61, 61, 61, 12: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 42: 61, 44: 61, 61, 61, 61, 61, 53: 61, 84: 61, 61, 61, 61, 61, 61, 61, 61, 61, 94: 61, 61, 112: 61},
		{84: 666},
		{667},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 109: 668},
		{9: 669, 413},
		// 380
		{8: 670},
		{62, 62, 62, 62, 62, 62, 62, 62, 62, 12: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 42: 62, 44: 62, 62, 62, 62, 62, 53: 62, 84: 62, 62, 62, 62, 62, 62, 62, 62, 62, 94: 62, 62, 112: 62},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 109: 672},
		{9: 673, 413},
		{350, 355, 620, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 53: 617, 84: 638, 643, 628, 642, 629, 639, 640, 641, 636, 94: 644, 637, 109: 621, 117: 635, 122: 623, 624, 626, 627, 622, 625, 674},
		// 385
		{63, 63, 63, 63, 63, 63, 63, 63, 63, 12: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 42: 63, 44: 63, 63, 63, 63, 63, 53: 63, 84: 63, 63, 63, 63, 63, 63, 63, 63, 63, 94: 63, 63, 112: 63},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 109: 676},
		{9: 677, 413},
		{350, 355, 620, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 53: 617, 84: 638, 643, 628, 642, 629, 639, 640, 641, 636, 94: 644, 637, 109: 621, 117: 635, 122: 623, 624, 626, 627, 622, 625, 678},
		{64, 64, 64, 64, 64, 64, 64, 64, 64, 12: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 42: 64, 44: 64, 64, 64, 64, 64, 53: 64, 84: 64, 64, 64, 64, 64, 64, 64, 64, 64, 94: 64, 64, 112: 64},
		// 390
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 109: 680},
		{9: 681, 413},
		{350, 355, 620, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 53: 617, 84: 638, 643, 628, 642, 629, 639, 640, 641, 636, 94: 644, 637, 109: 621, 117: 635, 122: 623, 624, 626, 627, 622, 625, 682},
		{66, 66, 66, 66, 66, 66, 66, 66, 66, 12: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 42: 66, 44: 66, 66, 66, 66, 66, 53: 66, 84: 66, 66, 66, 66, 66, 66, 66, 66, 66, 94: 66, 66, 112: 683},
		{350, 355, 620, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 53: 617, 84: 638, 643, 628, 642, 629, 639, 640, 641, 636, 94: 644, 637, 109: 621, 117: 635, 122: 623, 624, 626, 627, 622, 625, 684},
		// 395
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 12: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 42: 65, 44: 65, 65, 65, 65, 65, 53: 65, 84: 65, 65, 65, 65, 65, 65, 65, 65, 65, 94: 65, 65, 112: 65},
		{67, 67, 67, 67, 67, 67, 67, 67, 67, 12: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 42: 67, 44: 67, 67, 67, 67, 67, 53: 67, 84: 67, 67, 67, 67, 67, 67, 67, 67, 67, 94: 67, 67, 112: 67},
		{72, 72, 72, 72, 72, 72, 72, 72, 72, 12: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 42: 72, 44: 72, 72, 72, 72, 72, 53: 72, 84: 72, 72, 72, 72, 72, 72, 72, 72, 72, 94: 72, 72},
		{74, 74, 74, 74, 74, 74, 74, 74, 74, 12: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 42: 74, 74, 74, 74, 74, 74, 74, 53: 74, 84: 74, 74, 74, 74, 74, 74, 74, 74, 74, 94: 74, 74, 112: 74},
		{350, 355, 620, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 53: 617, 84: 638, 643, 628, 642, 629, 639, 640, 641, 636, 94: 644, 637, 109: 621, 117: 635, 122: 623, 624, 626, 627, 622, 625, 689},
		// 400
		{76, 76, 76, 76, 76, 76, 76, 76, 76, 12: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 42: 76, 44: 76, 76, 76, 76, 76, 53: 76, 84: 76, 76, 76, 76, 76, 76, 76, 76, 76, 94: 76, 76, 112: 76},
		{41: 691},
		{350, 355, 620, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 53: 617, 84: 638, 643, 628, 642, 629, 639, 640, 641, 636, 94: 644, 637, 109: 621, 117: 635, 122: 623, 624, 626, 627, 622, 625, 692},
		{77, 77, 77, 77, 77, 77, 77, 77, 77, 12: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 42: 77, 44: 77, 77, 77, 77, 77, 53: 77, 84: 77, 77, 77, 77, 77, 77, 77, 77, 77, 94: 77, 77, 112: 77},
		{350, 355, 620, 354, 356, 357, 353, 352, 217, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 410, 53: 617, 84: 638, 643, 628, 642, 629, 639, 640, 641, 636, 94: 644, 637, 109: 621, 117: 635, 122: 623, 624, 626, 627, 622, 625, 694},
		// 405
		{78, 78, 78, 78, 78, 78, 78, 78, 78, 12: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 42: 78, 44: 78, 78, 78, 78, 78, 53: 78, 84: 78, 78, 78, 78, 78, 78, 78, 78, 78, 94: 78, 78, 112: 78},
		{350, 355, 343, 354, 356, 357, 353, 352, 12: 359, 358, 344, 345, 346, 347, 348, 360, 349, 50: 546, 53: 547, 158: 696},
		{8: 200, 10: 200},
		{133, 450, 133, 129: 585, 584, 132: 615, 157: 698},
		{8: 205, 10: 205},
		// 410
		{213, 213, 213, 213, 213, 213, 213, 213, 213, 12: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 42: 213, 213, 213, 213, 213, 213, 213, 53: 213, 84: 213, 213, 213, 213, 213, 213, 213, 213, 213, 94: 213, 213},
		{21: 54, 54, 54, 25: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 42: 54, 54, 54, 54, 54, 54, 54},
		{215, 215, 215, 215, 215, 215, 215, 215, 12: 215, 215, 215, 215, 215, 215, 215, 215, 215, 140: 341, 702},
		{43: 285},
		{80: 725, 727, 96: 715, 716, 717, 712, 713, 714, 718, 722, 719, 709, 720, 721, 110: 726, 724, 118: 723, 131: 707, 133: 706, 711, 708, 710, 138: 705, 211: 704},
		// 415
		{43: 287},
		{43: 44, 80: 725, 727, 96: 715, 716, 717, 712, 713, 714, 718, 722, 719, 709, 720, 721, 110: 726, 724, 118: 723, 131: 707, 133: 770, 711, 708, 710},
		{43: 43, 80: 43, 43, 43, 43, 93: 43, 96: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43},
		{43: 39, 80: 39, 39, 39, 39, 93: 39, 96: 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39},
		{43: 38, 80: 38, 38, 38, 38, 93: 38, 96: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		// 420
		{81: 727, 110: 792, 724},
		{43: 36, 80: 36, 36, 36, 36, 93: 36, 96: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
		{82: 29, 29, 93: 780, 170: 778, 203: 779, 777},
		{81: 727, 110: 774, 724},
		{2: 771},
		// 425
		{2: 766},
		{2: 742, 80: 744, 209: 743},
		{80: 725, 727, 110: 726, 724, 118: 741},
		{43: 17, 80: 17, 17, 17, 17, 93: 17, 96: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{81: 727, 110: 739, 724},
		// 430
		{81: 727, 110: 737, 724},
		{80: 725, 727, 110: 726, 724, 118: 736},
		{2: 732},
		{81: 727, 110: 730, 724},
		{43: 7, 80: 7, 7, 7, 7, 93: 7, 96: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		// 435
		{80: 5, 729},
		{43: 4, 80: 4, 4, 4, 4, 93: 4, 96: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{80: 728},
		{80: 2, 2},
		{43: 3, 80: 3, 3, 3, 3, 93: 3, 96: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		// 440
		{80: 1, 1},
		{80: 731},
		{43: 8, 80: 8, 8, 8, 8, 93: 8, 96: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{80: 733, 727, 110: 734, 724},
		{43: 13, 80: 13, 13, 13, 13, 93: 13, 96: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		// 445
		{80: 735},
		{43: 9, 80: 9, 9, 9, 9, 93: 9, 96: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{43: 14, 80: 14, 14, 14, 14, 93: 14, 96: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{80: 738},
		{43: 15, 80: 15, 15, 15, 15, 93: 15, 96: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		// 450
		{80: 740},
		{43: 16, 80: 16, 16, 16, 16, 93: 16, 96: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{43: 18, 80: 18, 18, 18, 18, 93: 18, 96: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{80: 725, 727, 110: 726, 724, 118: 751, 139: 765},
		{2: 745, 9: 117, 142: 747, 175: 746, 748},
		// 455
		{43: 10, 80: 10, 10, 10, 10, 93: 10, 96: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{9: 119, 119, 142: 762},
		{9: 116, 754},
		{9: 752},
		{9: 749},
		// 460
		{80: 725, 727, 110: 726, 724, 118: 751, 139: 750},
		{43: 19, 80: 19, 19, 19, 19, 93: 19, 96: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{43: 6, 80: 6, 6, 6, 6, 93: 6, 96: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{80: 725, 727, 110: 726, 724, 118: 751, 139: 753},
		{43: 21, 80: 21, 21, 21, 21, 93: 21, 96: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		// 465
		{2: 755, 142: 756},
		{9: 118, 118, 142: 759},
		{9: 757},
		{80: 725, 727, 110: 726, 724, 118: 751, 139: 758},
		{43: 20, 80: 20, 20, 20, 20, 93: 20, 96: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		// 470
		{9: 760},
		{80: 725, 727, 110: 726, 724, 118: 751, 139: 761},
		{43: 11, 80: 11, 11, 11, 11, 93: 11, 96: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{9: 763},
		{80: 725, 727, 110: 726, 724, 118: 751, 139: 764},
		// 475
		{43: 12, 80: 12, 12, 12, 12, 93: 12, 96: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{43: 22, 80: 22, 22, 22, 22, 93: 22, 96: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{80: 767},
		{80: 725, 727, 41, 41, 93: 41, 96: 715, 716, 717, 712, 713, 714, 718, 722, 719, 709, 720, 721, 110: 726, 724, 118: 723, 131: 707, 133: 706, 711, 708, 710, 138: 768, 143: 769},
		{80: 725, 727, 40, 40, 93: 40, 96: 715, 716, 717, 712, 713, 714, 718, 722, 719, 709, 720, 721, 110: 726, 724, 118: 723, 131: 707, 133: 770, 711, 708, 710},
		// 480
		{82: 32, 32, 93: 32},
		{43: 42, 80: 42, 42, 42, 42, 93: 42, 96: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{80: 772},
		{80: 725, 727, 41, 41, 93: 41, 96: 715, 716, 717, 712, 713, 714, 718, 722, 719, 709, 720, 721, 110: 726, 724, 118: 723, 131: 707, 133: 706, 711, 708, 710, 138: 768, 143: 773},
		{82: 33, 33, 93: 33},
		// 485
		{80: 775},
		{80: 725, 727, 41, 41, 93: 41, 96: 715, 716, 717, 712, 713, 714, 718, 722, 719, 709, 720, 721, 110: 726, 724, 118: 723, 131: 707, 133: 706, 711, 708, 710, 138: 768, 143: 776},
		{82: 34, 34, 93: 34},
		{82: 25, 786, 205: 787, 785},
		{82: 31, 31, 93: 31},
		// 490
		{82: 28, 28, 93: 780, 170: 784},
		{81: 727, 110: 781, 724},
		{80: 782},
		{80: 725, 727, 41, 41, 93: 41, 96: 715, 716, 717, 712, 713, 714, 718, 722, 719, 709, 720, 721, 110: 726, 724, 118: 723, 131: 707, 133: 706, 711, 708, 710, 138: 768, 143: 783},
		{82: 27, 27, 93: 27},
		// 495
		{82: 30, 30, 93: 30},
		{82: 791, 207: 790},
		{80: 788},
		{82: 24},
		{80: 725, 727, 41, 96: 715, 716, 717, 712, 713, 714, 718, 722, 719, 709, 720, 721, 110: 726, 724, 118: 723, 131: 707, 133: 706, 711, 708, 710, 138: 768, 143: 789},
		// 500
		{82: 26},
		{43: 35, 80: 35, 35, 35, 35, 93: 35, 96: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{43: 23, 80: 23, 23, 23, 23, 93: 23, 96: 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23},
		{80: 793},
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
	const yyError = 219

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
	case 239:
		{
			yyVAL.node = &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				CompoundStatement:     yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 240:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 241:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 242:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 243:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 244:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScopePos(lhs.Pos())
		}
	case 245:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 246:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 247:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 248:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 249:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 250:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 251:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 252:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 253:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 254:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 255:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 256:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 257:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 258:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 259:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 260:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 261:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 262:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 263:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 264:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 265:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 266:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 267:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 268:
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
	case 269:
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
	case 270:
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
	case 271:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 272:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 273:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 274:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 275:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 276:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 277:
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
	case 278:
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
	case 279:
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
	case 280:
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
	case 281:
		{
			yyVAL.node = &ControlLine{
				Case:        14,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 284:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 285:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
