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
	yyDefault           = 57458
	yyEofCode           = 57344
	ADDASSIGN           = 57346
	ALIGNOF             = 57347
	ANDAND              = 57348
	ANDASSIGN           = 57349
	ARROW               = 57350
	ASM                 = 57351
	AUTO                = 57352
	BOOL                = 57353
	BREAK               = 57354
	CASE                = 57355
	CAST                = 57356
	CHAR                = 57357
	CHARCONST           = 57358
	COMPLEX             = 57359
	CONST               = 57360
	CONSTANT_EXPRESSION = 1048577
	CONTINUE            = 57361
	DDD                 = 57362
	DEC                 = 57363
	DEFAULT             = 57364
	DIVASSIGN           = 57365
	DO                  = 57366
	DOUBLE              = 57367
	ELSE                = 57368
	ENUM                = 57369
	EQ                  = 57370
	EXTERN              = 57371
	FLOAT               = 57372
	FLOATCONST          = 57373
	FOR                 = 57374
	GEQ                 = 57375
	GOTO                = 57376
	IDENTIFIER          = 57377
	IDENTIFIER_LPAREN   = 57378
	IDENTIFIER_NONREPL  = 57379
	IF                  = 57380
	INC                 = 57381
	INLINE              = 57382
	INT                 = 57383
	INTCONST            = 57384
	LEQ                 = 57385
	LONG                = 57386
	LONGCHARCONST       = 57387
	LONGSTRINGLITERAL   = 57388
	LSH                 = 57389
	LSHASSIGN           = 57390
	MODASSIGN           = 57391
	MULASSIGN           = 57392
	NEQ                 = 57393
	NOELSE              = 57394
	NORETURN            = 57395
	ORASSIGN            = 57396
	OROR                = 57397
	PPDEFINE            = 57398
	PPELIF              = 57399
	PPELSE              = 57400
	PPENDIF             = 57401
	PPERROR             = 57402
	PPHASH_NL           = 57403
	PPHEADER_NAME       = 57404
	PPIF                = 57405
	PPIFDEF             = 57406
	PPIFNDEF            = 57407
	PPINCLUDE           = 57408
	PPINCLUDE_NEXT      = 57409
	PPLINE              = 57410
	PPNONDIRECTIVE      = 57411
	PPNUMBER            = 57412
	PPOTHER             = 57413
	PPPASTE             = 57414
	PPPRAGMA            = 57415
	PPUNDEF             = 57416
	PREPROCESSING_FILE  = 1048576
	REGISTER            = 57417
	RESTRICT            = 57418
	RETURN              = 57419
	RSH                 = 57420
	RSHASSIGN           = 57421
	SHORT               = 57422
	SIGNED              = 57423
	SIZEOF              = 57424
	STATIC              = 57425
	STATIC_ASSERT       = 57426
	STRINGLITERAL       = 57427
	STRUCT              = 57428
	SUBASSIGN           = 57429
	SWITCH              = 57430
	TRANSLATION_UNIT    = 1048578
	TYPEDEF             = 57431
	TYPEDEFNAME         = 57432
	TYPEOF              = 57433
	UNARY               = 57434
	UNION               = 57435
	UNSIGNED            = 57436
	VOID                = 57437
	VOLATILE            = 57438
	WHILE               = 57439
	XORASSIGN           = 57440
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -315
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (304x)
		42:      1,   // '*' (265x)
		57377:   2,   // IDENTIFIER (218x)
		38:      3,   // '&' (212x)
		43:      4,   // '+' (212x)
		45:      5,   // '-' (212x)
		57363:   6,   // DEC (212x)
		57381:   7,   // INC (212x)
		41:      8,   // ')' (195x)
		59:      9,   // ';' (195x)
		44:      10,  // ',' (182x)
		57427:   11,  // STRINGLITERAL (160x)
		91:      12,  // '[' (159x)
		33:      13,  // '!' (142x)
		126:     14,  // '~' (142x)
		57347:   15,  // ALIGNOF (142x)
		57358:   16,  // CHARCONST (142x)
		57373:   17,  // FLOATCONST (142x)
		57384:   18,  // INTCONST (142x)
		57387:   19,  // LONGCHARCONST (142x)
		57388:   20,  // LONGSTRINGLITERAL (142x)
		57424:   21,  // SIZEOF (142x)
		57438:   22,  // VOLATILE (133x)
		57360:   23,  // CONST (131x)
		57418:   24,  // RESTRICT (131x)
		57353:   25,  // BOOL (121x)
		57357:   26,  // CHAR (121x)
		57359:   27,  // COMPLEX (121x)
		57367:   28,  // DOUBLE (121x)
		57369:   29,  // ENUM (121x)
		57372:   30,  // FLOAT (121x)
		57383:   31,  // INT (121x)
		57386:   32,  // LONG (121x)
		57422:   33,  // SHORT (121x)
		57423:   34,  // SIGNED (121x)
		57428:   35,  // STRUCT (121x)
		57432:   36,  // TYPEDEFNAME (121x)
		57433:   37,  // TYPEOF (121x)
		57435:   38,  // UNION (121x)
		57436:   39,  // UNSIGNED (121x)
		57437:   40,  // VOID (121x)
		58:      41,  // ':' (119x)
		125:     42,  // '}' (118x)
		57425:   43,  // STATIC (113x)
		57352:   44,  // AUTO (107x)
		57371:   45,  // EXTERN (107x)
		57382:   46,  // INLINE (107x)
		57395:   47,  // NORETURN (107x)
		57417:   48,  // REGISTER (107x)
		57431:   49,  // TYPEDEF (107x)
		57344:   50,  // $end (101x)
		61:      51,  // '=' (90x)
		57501:   52,  // Expression (85x)
		93:      53,  // ']' (83x)
		123:     54,  // '{' (83x)
		46:      55,  // '.' (79x)
		57351:   56,  // ASM (76x)
		37:      57,  // '%' (71x)
		47:      58,  // '/' (71x)
		60:      59,  // '<' (71x)
		62:      60,  // '>' (71x)
		63:      61,  // '?' (71x)
		94:      62,  // '^' (71x)
		124:     63,  // '|' (71x)
		57346:   64,  // ADDASSIGN (71x)
		57348:   65,  // ANDAND (71x)
		57349:   66,  // ANDASSIGN (71x)
		57350:   67,  // ARROW (71x)
		57365:   68,  // DIVASSIGN (71x)
		57370:   69,  // EQ (71x)
		57375:   70,  // GEQ (71x)
		57385:   71,  // LEQ (71x)
		57389:   72,  // LSH (71x)
		57390:   73,  // LSHASSIGN (71x)
		57391:   74,  // MODASSIGN (71x)
		57392:   75,  // MULASSIGN (71x)
		57393:   76,  // NEQ (71x)
		57396:   77,  // ORASSIGN (71x)
		57397:   78,  // OROR (71x)
		57420:   79,  // RSH (71x)
		57421:   80,  // RSHASSIGN (71x)
		57429:   81,  // SUBASSIGN (71x)
		57440:   82,  // XORASSIGN (71x)
		10:      83,  // '\n' (58x)
		57413:   84,  // PPOTHER (52x)
		57376:   85,  // GOTO (50x)
		57439:   86,  // WHILE (48x)
		57354:   87,  // BREAK (47x)
		57355:   88,  // CASE (47x)
		57361:   89,  // CONTINUE (47x)
		57364:   90,  // DEFAULT (47x)
		57366:   91,  // DO (47x)
		57374:   92,  // FOR (47x)
		57380:   93,  // IF (47x)
		57419:   94,  // RETURN (47x)
		57430:   95,  // SWITCH (47x)
		57401:   96,  // PPENDIF (44x)
		57400:   97,  // PPELSE (40x)
		57399:   98,  // PPELIF (39x)
		57398:   99,  // PPDEFINE (35x)
		57402:   100, // PPERROR (35x)
		57403:   101, // PPHASH_NL (35x)
		57405:   102, // PPIF (35x)
		57406:   103, // PPIFDEF (35x)
		57407:   104, // PPIFNDEF (35x)
		57408:   105, // PPINCLUDE (35x)
		57409:   106, // PPINCLUDE_NEXT (35x)
		57410:   107, // PPLINE (35x)
		57411:   108, // PPNONDIRECTIVE (35x)
		57415:   109, // PPPRAGMA (35x)
		57416:   110, // PPUNDEF (35x)
		57368:   111, // ELSE (29x)
		57553:   112, // TypeQualifier (28x)
		57502:   113, // ExpressionList (26x)
		57526:   114, // PPTokenList (22x)
		57528:   115, // PPTokens (22x)
		57497:   116, // EnumSpecifier (20x)
		57548:   117, // StructOrUnion (20x)
		57549:   118, // StructOrUnionSpecifier (20x)
		57556:   119, // TypeSpecifier (20x)
		57503:   120, // ExpressionListOpt (18x)
		57468:   121, // BasicAssemblerStatement (15x)
		57480:   122, // DeclarationSpecifiers (15x)
		57509:   123, // FunctionSpecifier (15x)
		57527:   124, // PPTokenListOpt (15x)
		57426:   125, // STATIC_ASSERT (15x)
		57543:   126, // StorageClassSpecifier (15x)
		57466:   127, // AssemblerStatement (13x)
		57474:   128, // CompoundStatement (13x)
		57505:   129, // ExpressionStatement (12x)
		57523:   130, // IterationStatement (12x)
		57524:   131, // JumpStatement (12x)
		57525:   132, // LabeledStatement (12x)
		57537:   133, // SelectionStatement (12x)
		57541:   134, // Statement (12x)
		57533:   135, // Pointer (11x)
		57534:   136, // PointerOpt (10x)
		57476:   137, // ControlLine (8x)
		57482:   138, // Declarator (8x)
		57512:   139, // GroupPart (8x)
		57516:   140, // IfGroup (8x)
		57517:   141, // IfSection (8x)
		57550:   142, // TextLine (8x)
		57477:   143, // Declaration (7x)
		57452:   144, // $@4 (6x)
		57475:   145, // ConstantExpression (6x)
		57362:   146, // DDD (6x)
		57510:   147, // GroupList (6x)
		57511:   148, // GroupListOpt (5x)
		57536:   149, // ReplacementList (5x)
		57538:   150, // SpecifierQualifierList (5x)
		57554:   151, // TypeQualifierList (5x)
		57442:   152, // $@10 (4x)
		57459:   153, // AbstractDeclarator (4x)
		57464:   154, // AssemblerOperand (4x)
		57467:   155, // AssemblerSymbolicNameOpt (4x)
		57481:   156, // DeclarationSpecifiersOpt (4x)
		57486:   157, // Designator (4x)
		57529:   158, // ParameterDeclaration (4x)
		57552:   159, // TypeName (4x)
		57555:   160, // TypeQualifierListOpt (4x)
		57463:   161, // AssemblerInstructions (3x)
		57465:   162, // AssemblerOperands (3x)
		57473:   163, // CommaOpt (3x)
		57484:   164, // Designation (3x)
		57485:   165, // DesignationOpt (3x)
		57487:   166, // DesignatorList (3x)
		57504:   167, // ExpressionOpt (3x)
		57513:   168, // IdentifierList (3x)
		57518:   169, // InitDeclarator (3x)
		57521:   170, // Initializer (3x)
		57530:   171, // ParameterList (3x)
		57531:   172, // ParameterTypeList (3x)
		57443:   173, // $@11 (2x)
		57453:   174, // $@5 (2x)
		57460:   175, // AbstractDeclaratorOpt (2x)
		57469:   176, // BlockItem (2x)
		57472:   177, // Clobbers (2x)
		57483:   178, // DeclaratorOpt (2x)
		57488:   179, // DirectAbstractDeclarator (2x)
		57489:   180, // DirectAbstractDeclaratorOpt (2x)
		57490:   181, // DirectDeclarator (2x)
		57491:   182, // ElifGroup (2x)
		57498:   183, // EnumerationConstant (2x)
		57499:   184, // Enumerator (2x)
		57506:   185, // ExternalDeclaration (2x)
		57508:   186, // FunctionDefinition (2x)
		57514:   187, // IdentifierListOpt (2x)
		57515:   188, // IdentifierOpt (2x)
		57519:   189, // InitDeclaratorList (2x)
		57520:   190, // InitDeclaratorListOpt (2x)
		57522:   191, // InitializerList (2x)
		57532:   192, // ParameterTypeListOpt (2x)
		57539:   193, // SpecifierQualifierListOpt (2x)
		57542:   194, // StaticAssert (2x)
		57544:   195, // StructDeclaration (2x)
		57546:   196, // StructDeclarator (2x)
		57557:   197, // VolatileOpt (2x)
		57441:   198, // $@1 (1x)
		57444:   199, // $@12 (1x)
		57445:   200, // $@13 (1x)
		57446:   201, // $@14 (1x)
		57447:   202, // $@15 (1x)
		57448:   203, // $@16 (1x)
		57449:   204, // $@17 (1x)
		57450:   205, // $@2 (1x)
		57451:   206, // $@3 (1x)
		57454:   207, // $@6 (1x)
		57455:   208, // $@7 (1x)
		57456:   209, // $@8 (1x)
		57457:   210, // $@9 (1x)
		57461:   211, // ArgumentExpressionList (1x)
		57462:   212, // ArgumentExpressionListOpt (1x)
		57470:   213, // BlockItemList (1x)
		57471:   214, // BlockItemListOpt (1x)
		1048577: 215, // CONSTANT_EXPRESSION (1x)
		57478:   216, // DeclarationList (1x)
		57479:   217, // DeclarationListOpt (1x)
		57492:   218, // ElifGroupList (1x)
		57493:   219, // ElifGroupListOpt (1x)
		57494:   220, // ElseGroup (1x)
		57495:   221, // ElseGroupOpt (1x)
		57496:   222, // EndifLine (1x)
		57500:   223, // EnumeratorList (1x)
		57507:   224, // FunctionBody (1x)
		57378:   225, // IDENTIFIER_LPAREN (1x)
		1048576: 226, // PREPROCESSING_FILE (1x)
		57535:   227, // PreprocessingFile (1x)
		57540:   228, // Start (1x)
		57545:   229, // StructDeclarationList (1x)
		57547:   230, // StructDeclaratorList (1x)
		1048578: 231, // TRANSLATION_UNIT (1x)
		57551:   232, // TranslationUnit (1x)
		57458:   233, // $default (0x)
		57356:   234, // CAST (0x)
		57345:   235, // error (0x)
		57379:   236, // IDENTIFIER_NONREPL (0x)
		57394:   237, // NOELSE (0x)
		57404:   238, // PPHEADER_NAME (0x)
		57412:   239, // PPNUMBER (0x)
		57414:   240, // PPPASTE (0x)
		57434:   241, // UNARY (0x)
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
		"STRINGLITERAL",
		"'['",
		"'!'",
		"'~'",
		"ALIGNOF",
		"CHARCONST",
		"FLOATCONST",
		"INTCONST",
		"LONGCHARCONST",
		"LONGSTRINGLITERAL",
		"SIZEOF",
		"VOLATILE",
		"CONST",
		"RESTRICT",
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
		"'}'",
		"STATIC",
		"AUTO",
		"EXTERN",
		"INLINE",
		"NORETURN",
		"REGISTER",
		"TYPEDEF",
		"$end",
		"'='",
		"Expression",
		"']'",
		"'{'",
		"'.'",
		"ASM",
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
		"GOTO",
		"WHILE",
		"BREAK",
		"CASE",
		"CONTINUE",
		"DEFAULT",
		"DO",
		"FOR",
		"IF",
		"RETURN",
		"SWITCH",
		"PPENDIF",
		"PPELSE",
		"PPELIF",
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
		"ELSE",
		"TypeQualifier",
		"ExpressionList",
		"PPTokenList",
		"PPTokens",
		"EnumSpecifier",
		"StructOrUnion",
		"StructOrUnionSpecifier",
		"TypeSpecifier",
		"ExpressionListOpt",
		"BasicAssemblerStatement",
		"DeclarationSpecifiers",
		"FunctionSpecifier",
		"PPTokenListOpt",
		"STATIC_ASSERT",
		"StorageClassSpecifier",
		"AssemblerStatement",
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
		"$@4",
		"ConstantExpression",
		"DDD",
		"GroupList",
		"GroupListOpt",
		"ReplacementList",
		"SpecifierQualifierList",
		"TypeQualifierList",
		"$@10",
		"AbstractDeclarator",
		"AssemblerOperand",
		"AssemblerSymbolicNameOpt",
		"DeclarationSpecifiersOpt",
		"Designator",
		"ParameterDeclaration",
		"TypeName",
		"TypeQualifierListOpt",
		"AssemblerInstructions",
		"AssemblerOperands",
		"CommaOpt",
		"Designation",
		"DesignationOpt",
		"DesignatorList",
		"ExpressionOpt",
		"IdentifierList",
		"InitDeclarator",
		"Initializer",
		"ParameterList",
		"ParameterTypeList",
		"$@11",
		"$@5",
		"AbstractDeclaratorOpt",
		"BlockItem",
		"Clobbers",
		"DeclaratorOpt",
		"DirectAbstractDeclarator",
		"DirectAbstractDeclaratorOpt",
		"DirectDeclarator",
		"ElifGroup",
		"EnumerationConstant",
		"Enumerator",
		"ExternalDeclaration",
		"FunctionDefinition",
		"IdentifierListOpt",
		"IdentifierOpt",
		"InitDeclaratorList",
		"InitDeclaratorListOpt",
		"InitializerList",
		"ParameterTypeListOpt",
		"SpecifierQualifierListOpt",
		"StaticAssert",
		"StructDeclaration",
		"StructDeclarator",
		"VolatileOpt",
		"$@1",
		"$@12",
		"$@13",
		"$@14",
		"$@15",
		"$@16",
		"$@17",
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
		"FunctionBody",
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

	yyTokenLiteralStrings = map[int]string{
		57377:   "identifier",
		57363:   "--",
		57381:   "++",
		57427:   "string literal",
		57347:   "_Alignof",
		57358:   "character constant",
		57373:   "floating-point constant",
		57384:   "integer constant",
		57387:   "long character constant",
		57388:   "long string constant",
		57424:   "sizeof",
		57438:   "volatile",
		57360:   "const",
		57418:   "restrict",
		57353:   "_Bool",
		57357:   "char",
		57359:   "_Complex",
		57367:   "double",
		57369:   "enum",
		57372:   "float",
		57383:   "int",
		57386:   "long",
		57422:   "short",
		57423:   "signed",
		57428:   "struct",
		57432:   "typedefname",
		57433:   "typeof",
		57435:   "union",
		57436:   "unsigned",
		57437:   "void",
		57425:   "static",
		57352:   "auto",
		57371:   "extern",
		57382:   "inline",
		57395:   "_Noreturn",
		57417:   "register",
		57431:   "typedef",
		57351:   "asm",
		57346:   "+=",
		57348:   "&&",
		57349:   "&=",
		57350:   "->",
		57365:   "/=",
		57370:   "==",
		57375:   ">=",
		57385:   "<=",
		57389:   "<<",
		57390:   "<<=",
		57391:   "%=",
		57392:   "*=",
		57393:   "!=",
		57396:   "|=",
		57397:   "||",
		57420:   ">>",
		57421:   ">>=",
		57429:   "-=",
		57440:   "^=",
		57413:   "ppother",
		57376:   "goto",
		57439:   "while",
		57354:   "break",
		57355:   "case",
		57361:   "continue",
		57364:   "default",
		57366:   "do",
		57374:   "for",
		57380:   "if",
		57419:   "return",
		57430:   "switch",
		57401:   "#endif",
		57400:   "#else",
		57399:   "#elif",
		57398:   "#define",
		57402:   "#error",
		57403:   "#",
		57405:   "#if",
		57406:   "#ifdef",
		57407:   "#ifndef",
		57408:   "#include",
		57409:   "#include_next",
		57410:   "#line",
		57411:   "#foo",
		57415:   "#pragma",
		57416:   "#undef",
		57368:   "else",
		57426:   "_Static_assert",
		57362:   "...",
		1048577: "constant expression prefix",
		57378:   "identifier immediatelly followed by '('",
		1048576: "preprocessing file prefix",
		1048578: "translation unit prefix",
		57379:   "non replaceable identifier",
		57404:   "header name",
		57412:   "preprocessing number",
		57414:   "##",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {198, 0},
		2:   {228, 3},
		3:   {205, 0},
		4:   {228, 3},
		5:   {206, 0},
		6:   {228, 3},
		7:   {183, 1},
		8:   {211, 1},
		9:   {211, 3},
		10:  {212, 0},
		11:  {212, 1},
		12:  {52, 1},
		13:  {52, 1},
		14:  {52, 1},
		15:  {52, 1},
		16:  {52, 1},
		17:  {52, 1},
		18:  {52, 1},
		19:  {52, 3},
		20:  {52, 4},
		21:  {52, 4},
		22:  {52, 3},
		23:  {52, 3},
		24:  {52, 2},
		25:  {52, 2},
		26:  {52, 7},
		27:  {52, 2},
		28:  {52, 2},
		29:  {52, 2},
		30:  {52, 2},
		31:  {52, 2},
		32:  {52, 2},
		33:  {52, 2},
		34:  {52, 2},
		35:  {52, 2},
		36:  {52, 4},
		37:  {52, 4},
		38:  {52, 3},
		39:  {52, 3},
		40:  {52, 3},
		41:  {52, 3},
		42:  {52, 3},
		43:  {52, 3},
		44:  {52, 3},
		45:  {52, 3},
		46:  {52, 3},
		47:  {52, 3},
		48:  {52, 3},
		49:  {52, 3},
		50:  {52, 3},
		51:  {52, 3},
		52:  {52, 3},
		53:  {52, 3},
		54:  {52, 3},
		55:  {52, 3},
		56:  {52, 5},
		57:  {52, 3},
		58:  {52, 3},
		59:  {52, 3},
		60:  {52, 3},
		61:  {52, 3},
		62:  {52, 3},
		63:  {52, 3},
		64:  {52, 3},
		65:  {52, 3},
		66:  {52, 3},
		67:  {52, 3},
		68:  {52, 4},
		69:  {167, 0},
		70:  {167, 1},
		71:  {113, 1},
		72:  {113, 3},
		73:  {120, 0},
		74:  {120, 1},
		75:  {144, 0},
		76:  {145, 2},
		77:  {143, 3},
		78:  {122, 2},
		79:  {122, 2},
		80:  {122, 2},
		81:  {122, 2},
		82:  {156, 0},
		83:  {156, 1},
		84:  {189, 1},
		85:  {189, 3},
		86:  {190, 0},
		87:  {190, 1},
		88:  {169, 1},
		89:  {174, 0},
		90:  {169, 4},
		91:  {126, 1},
		92:  {126, 1},
		93:  {126, 1},
		94:  {126, 1},
		95:  {126, 1},
		96:  {119, 1},
		97:  {119, 1},
		98:  {119, 1},
		99:  {119, 1},
		100: {119, 1},
		101: {119, 1},
		102: {119, 1},
		103: {119, 1},
		104: {119, 1},
		105: {119, 1},
		106: {119, 1},
		107: {119, 1},
		108: {119, 1},
		109: {119, 1},
		110: {119, 4},
		111: {119, 4},
		112: {207, 0},
		113: {208, 0},
		114: {118, 7},
		115: {118, 2},
		116: {117, 1},
		117: {117, 1},
		118: {229, 1},
		119: {229, 2},
		120: {195, 3},
		121: {195, 2},
		122: {150, 2},
		123: {150, 2},
		124: {193, 0},
		125: {193, 1},
		126: {230, 1},
		127: {230, 3},
		128: {196, 1},
		129: {196, 3},
		130: {163, 0},
		131: {163, 1},
		132: {209, 0},
		133: {116, 7},
		134: {116, 2},
		135: {223, 1},
		136: {223, 3},
		137: {184, 1},
		138: {184, 3},
		139: {112, 1},
		140: {112, 1},
		141: {112, 1},
		142: {123, 1},
		143: {123, 1},
		144: {138, 2},
		145: {178, 0},
		146: {178, 1},
		147: {181, 1},
		148: {181, 3},
		149: {181, 5},
		150: {181, 6},
		151: {181, 6},
		152: {181, 5},
		153: {210, 0},
		154: {181, 5},
		155: {181, 4},
		156: {135, 2},
		157: {135, 3},
		158: {136, 0},
		159: {136, 1},
		160: {151, 1},
		161: {151, 2},
		162: {160, 0},
		163: {160, 1},
		164: {172, 1},
		165: {172, 3},
		166: {192, 0},
		167: {192, 1},
		168: {171, 1},
		169: {171, 3},
		170: {158, 2},
		171: {158, 2},
		172: {168, 1},
		173: {168, 3},
		174: {187, 0},
		175: {187, 1},
		176: {188, 0},
		177: {188, 1},
		178: {152, 0},
		179: {159, 3},
		180: {153, 1},
		181: {153, 2},
		182: {175, 0},
		183: {175, 1},
		184: {179, 3},
		185: {179, 4},
		186: {179, 5},
		187: {179, 6},
		188: {179, 6},
		189: {179, 4},
		190: {173, 0},
		191: {179, 4},
		192: {199, 0},
		193: {179, 5},
		194: {180, 0},
		195: {180, 1},
		196: {170, 1},
		197: {170, 4},
		198: {191, 2},
		199: {191, 4},
		200: {164, 2},
		201: {165, 0},
		202: {165, 1},
		203: {166, 1},
		204: {166, 2},
		205: {157, 3},
		206: {157, 2},
		207: {134, 1},
		208: {134, 1},
		209: {134, 1},
		210: {134, 1},
		211: {134, 1},
		212: {134, 1},
		213: {134, 1},
		214: {132, 3},
		215: {132, 4},
		216: {132, 3},
		217: {200, 0},
		218: {128, 4},
		219: {213, 1},
		220: {213, 2},
		221: {214, 0},
		222: {214, 1},
		223: {176, 1},
		224: {176, 1},
		225: {129, 2},
		226: {133, 5},
		227: {133, 7},
		228: {133, 5},
		229: {130, 5},
		230: {130, 7},
		231: {130, 9},
		232: {130, 8},
		233: {131, 3},
		234: {131, 2},
		235: {131, 2},
		236: {131, 3},
		237: {232, 1},
		238: {232, 2},
		239: {185, 1},
		240: {185, 1},
		241: {185, 2},
		242: {185, 1},
		243: {201, 0},
		244: {186, 5},
		245: {202, 0},
		246: {224, 2},
		247: {203, 0},
		248: {224, 3},
		249: {216, 1},
		250: {216, 2},
		251: {217, 0},
		252: {204, 0},
		253: {217, 2},
		254: {161, 1},
		255: {161, 2},
		256: {121, 5},
		257: {197, 0},
		258: {197, 1},
		259: {154, 5},
		260: {162, 1},
		261: {162, 3},
		262: {155, 0},
		263: {155, 3},
		264: {177, 1},
		265: {177, 3},
		266: {127, 1},
		267: {127, 7},
		268: {127, 9},
		269: {127, 11},
		270: {127, 13},
		271: {194, 7},
		272: {227, 1},
		273: {147, 1},
		274: {147, 2},
		275: {148, 0},
		276: {148, 1},
		277: {139, 1},
		278: {139, 1},
		279: {139, 3},
		280: {139, 1},
		281: {141, 4},
		282: {140, 4},
		283: {140, 4},
		284: {140, 4},
		285: {218, 1},
		286: {218, 2},
		287: {219, 0},
		288: {219, 1},
		289: {182, 4},
		290: {220, 3},
		291: {221, 0},
		292: {221, 1},
		293: {222, 1},
		294: {137, 3},
		295: {137, 5},
		296: {137, 7},
		297: {137, 5},
		298: {137, 2},
		299: {137, 1},
		300: {137, 3},
		301: {137, 3},
		302: {137, 2},
		303: {137, 3},
		304: {137, 6},
		305: {137, 2},
		306: {137, 4},
		307: {137, 3},
		308: {142, 1},
		309: {149, 1},
		310: {114, 1},
		311: {124, 1},
		312: {124, 2},
		313: {115, 1},
		314: {115, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 50}:   "invalid empty input",
		yyXError{561, -1}: "expected #endif",
		yyXError{563, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{481, -1}: "expected $end",
		yyXError{483, -1}: "expected $end",
		yyXError{31, -1}:  "expected '('",
		yyXError{47, -1}:  "expected '('",
		yyXError{73, -1}:  "expected '('",
		yyXError{283, -1}: "expected '('",
		yyXError{359, -1}: "expected '('",
		yyXError{380, -1}: "expected '('",
		yyXError{415, -1}: "expected '('",
		yyXError{416, -1}: "expected '('",
		yyXError{417, -1}: "expected '('",
		yyXError{419, -1}: "expected '('",
		yyXError{445, -1}: "expected '('",
		yyXError{52, -1}:  "expected ')'",
		yyXError{75, -1}:  "expected ')'",
		yyXError{82, -1}:  "expected ')'",
		yyXError{174, -1}: "expected ')'",
		yyXError{189, -1}: "expected ')'",
		yyXError{192, -1}: "expected ')'",
		yyXError{195, -1}: "expected ')'",
		yyXError{203, -1}: "expected ')'",
		yyXError{208, -1}: "expected ')'",
		yyXError{214, -1}: "expected ')'",
		yyXError{230, -1}: "expected ')'",
		yyXError{235, -1}: "expected ')'",
		yyXError{246, -1}: "expected ')'",
		yyXError{281, -1}: "expected ')'",
		yyXError{330, -1}: "expected ')'",
		yyXError{435, -1}: "expected ')'",
		yyXError{441, -1}: "expected ')'",
		yyXError{525, -1}: "expected ')'",
		yyXError{526, -1}: "expected ')'",
		yyXError{533, -1}: "expected ')'",
		yyXError{536, -1}: "expected ')'",
		yyXError{50, -1}:  "expected ','",
		yyXError{317, -1}: "expected ':'",
		yyXError{362, -1}: "expected ':'",
		yyXError{408, -1}: "expected ':'",
		yyXError{469, -1}: "expected ':'",
		yyXError{44, -1}:  "expected ';'",
		yyXError{53, -1}:  "expected ';'",
		yyXError{338, -1}: "expected ';'",
		yyXError{354, -1}: "expected ';'",
		yyXError{414, -1}: "expected ';'",
		yyXError{421, -1}: "expected ';'",
		yyXError{422, -1}: "expected ';'",
		yyXError{424, -1}: "expected ';'",
		yyXError{428, -1}: "expected ';'",
		yyXError{431, -1}: "expected ';'",
		yyXError{433, -1}: "expected ';'",
		yyXError{439, -1}: "expected ';'",
		yyXError{448, -1}: "expected ';'",
		yyXError{342, -1}: "expected '='",
		yyXError{87, -1}:  "expected '['",
		yyXError{505, -1}: "expected '\\n'",
		yyXError{509, -1}: "expected '\\n'",
		yyXError{513, -1}: "expected '\\n'",
		yyXError{516, -1}: "expected '\\n'",
		yyXError{518, -1}: "expected '\\n'",
		yyXError{540, -1}: "expected '\\n'",
		yyXError{545, -1}: "expected '\\n'",
		yyXError{548, -1}: "expected '\\n'",
		yyXError{555, -1}: "expected '\\n'",
		yyXError{560, -1}: "expected '\\n'",
		yyXError{566, -1}: "expected '\\n'",
		yyXError{93, -1}:  "expected ']'",
		yyXError{182, -1}: "expected ']'",
		yyXError{226, -1}: "expected ']'",
		yyXError{258, -1}: "expected ']'",
		yyXError{368, -1}: "expected ']'",
		yyXError{291, -1}: "expected '{'",
		yyXError{293, -1}: "expected '{'",
		yyXError{305, -1}: "expected '{'",
		yyXError{307, -1}: "expected '{'",
		yyXError{267, -1}: "expected '}'",
		yyXError{271, -1}: "expected '}'",
		yyXError{302, -1}: "expected '}'",
		yyXError{409, -1}: "expected '}'",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{202, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{86, -1}:  "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{285, -1}: "expected assembler instructions or string literal",
		yyXError{358, -1}: "expected assembler instructions or string literal",
		yyXError{360, -1}: "expected assembler instructions or string literal",
		yyXError{370, -1}: "expected assembler operand or one of ['[', string literal]",
		yyXError{363, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{385, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{388, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{353, -1}: "expected assembler statement or asm",
		yyXError{411, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{371, -1}: "expected clobbers or string literal",
		yyXError{391, -1}: "expected clobbers or string literal",
		yyXError{352, -1}: "expected compound statement or '{'",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{48, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{255, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{299, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{321, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{407, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{480, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{344, -1}: "expected declaration list or one of [_Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{347, -1}: "expected declaration or one of ['{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{430, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{320, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{194, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{252, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{197, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{83, -1}:  "expected direct abstract declarator or one of ['(', '[']",
		yyXError{318, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{553, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{559, -1}: "expected endif line or #endif",
		yyXError{490, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{551, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{294, -1}: "expected enumerator list or identifier",
		yyXError{301, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{98, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{122, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{446, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{450, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{454, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{458, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{62, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{243, -1}: "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{247, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		yyXError{90, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{225, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{282, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{49, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{64, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{65, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{66, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{67, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{68, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{69, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{70, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{71, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{72, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{96, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{104, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{105, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{106, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{107, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{108, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{109, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{110, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{111, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{112, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{113, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{114, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{115, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{116, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{117, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{118, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{119, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{120, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{121, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{123, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{124, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{125, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{126, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{127, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{128, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{129, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{130, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{131, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{132, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{133, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{148, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{149, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{176, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{183, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{219, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{222, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{381, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{94, -1}:  "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{217, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{328, -1}: "expected expression or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{6, -1}:   "expected external declaration or one of [$end, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{343, -1}: "expected function body or one of ['{', asm]",
		yyXError{350, -1}: "expected function body or one of ['{', asm]",
		yyXError{341, -1}: "expected function body or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{542, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{484, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{100, -1}: "expected identifier",
		yyXError{101, -1}: "expected identifier",
		yyXError{211, -1}: "expected identifier",
		yyXError{256, -1}: "expected identifier",
		yyXError{367, -1}: "expected identifier",
		yyXError{420, -1}: "expected identifier",
		yyXError{492, -1}: "expected identifier",
		yyXError{493, -1}: "expected identifier",
		yyXError{500, -1}: "expected identifier",
		yyXError{375, -1}: "expected identifier list or identifier",
		yyXError{522, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{476, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{249, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{263, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{251, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{269, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{474, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{266, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{55, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{56, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{57, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{58, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{59, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{60, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{61, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{102, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{103, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
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
		yyXError{151, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{152, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{153, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{154, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{155, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{156, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{157, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{158, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{159, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{160, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{161, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{162, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{163, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{164, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{165, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{166, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{167, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{168, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{169, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{170, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{171, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{175, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{179, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{187, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{242, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{244, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{248, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{272, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{273, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{274, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{275, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{276, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{277, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{278, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{279, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{280, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{63, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{146, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{150, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{172, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{177, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{329, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{382, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{398, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{262, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{89, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{97, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{184, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{220, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{223, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{485, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{486, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{487, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{489, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{496, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{502, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{504, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{507, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{510, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{512, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{514, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{515, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{517, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{519, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{520, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{523, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{528, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{529, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{531, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{535, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{538, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{539, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{544, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{564, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{565, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{567, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{543, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{547, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{550, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{552, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{557, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{558, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{466, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{478, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{41, -1}:  "expected one of [$end, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{42, -1}:  "expected one of [$end, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{43, -1}:  "expected one of [$end, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{45, -1}:  "expected one of [$end, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{54, -1}:  "expected one of [$end, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{290, -1}: "expected one of [$end, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{351, -1}: "expected one of [$end, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{394, -1}: "expected one of [$end, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{396, -1}: "expected one of [$end, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{479, -1}: "expected one of [$end, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{36, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{38, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{91, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{180, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{289, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{356, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{377, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{387, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{390, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{393, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{400, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{401, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{402, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{403, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{404, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{405, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{406, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{425, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{426, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{427, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{429, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{437, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{443, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{449, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{453, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{457, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{461, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{463, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{464, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{468, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{471, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{473, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{410, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{412, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{413, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{465, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{253, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{260, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{292, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{306, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{17, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{18, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{19, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{20, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{21, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{22, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{23, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{24, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{25, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{26, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{27, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{28, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{29, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{30, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{303, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{326, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{331, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{332, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{12, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{39, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{40, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{333, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{334, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{335, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{336, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{337, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{239, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{240, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{241, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{200, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{201, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{204, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{213, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{215, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{221, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{224, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{227, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{228, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{81, -1}:  "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{238, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{85, -1}:  "expected one of ['(', ')', ',', '[']",
		yyXError{134, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{181, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{185, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{186, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{188, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{196, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{232, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{236, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{284, -1}: "expected one of ['(', goto]",
		yyXError{357, -1}: "expected one of ['(', goto]",
		yyXError{319, -1}: "expected one of ['(', identifier]",
		yyXError{365, -1}: "expected one of [')', ',', ':']",
		yyXError{372, -1}: "expected one of [')', ',', ':']",
		yyXError{378, -1}: "expected one of [')', ',', ':']",
		yyXError{379, -1}: "expected one of [')', ',', ':']",
		yyXError{383, -1}: "expected one of [')', ',', ':']",
		yyXError{386, -1}: "expected one of [')', ',', ':']",
		yyXError{389, -1}: "expected one of [')', ',', ':']",
		yyXError{399, -1}: "expected one of [')', ',', ';']",
		yyXError{209, -1}: "expected one of [')', ',', ...]",
		yyXError{212, -1}: "expected one of [')', ',', ...]",
		yyXError{524, -1}: "expected one of [')', ',', ...]",
		yyXError{84, -1}:  "expected one of [')', ',']",
		yyXError{173, -1}: "expected one of [')', ',']",
		yyXError{191, -1}: "expected one of [')', ',']",
		yyXError{193, -1}: "expected one of [')', ',']",
		yyXError{198, -1}: "expected one of [')', ',']",
		yyXError{199, -1}: "expected one of [')', ',']",
		yyXError{210, -1}: "expected one of [')', ',']",
		yyXError{231, -1}: "expected one of [')', ',']",
		yyXError{245, -1}: "expected one of [')', ',']",
		yyXError{376, -1}: "expected one of [')', ',']",
		yyXError{392, -1}: "expected one of [')', ',']",
		yyXError{447, -1}: "expected one of [')', ',']",
		yyXError{451, -1}: "expected one of [')', ',']",
		yyXError{455, -1}: "expected one of [')', ',']",
		yyXError{459, -1}: "expected one of [')', ',']",
		yyXError{286, -1}: "expected one of [')', ':', string literal]",
		yyXError{288, -1}: "expected one of [')', ':', string literal]",
		yyXError{384, -1}: "expected one of [')', ':', string literal]",
		yyXError{287, -1}: "expected one of [')', string literal]",
		yyXError{316, -1}: "expected one of [',', ':', ';']",
		yyXError{147, -1}: "expected one of [',', ':']",
		yyXError{366, -1}: "expected one of [',', ':']",
		yyXError{373, -1}: "expected one of [',', ':']",
		yyXError{349, -1}: "expected one of [',', ';', '=']",
		yyXError{268, -1}: "expected one of [',', ';', '}']",
		yyXError{313, -1}: "expected one of [',', ';']",
		yyXError{315, -1}: "expected one of [',', ';']",
		yyXError{322, -1}: "expected one of [',', ';']",
		yyXError{325, -1}: "expected one of [',', ';']",
		yyXError{339, -1}: "expected one of [',', ';']",
		yyXError{340, -1}: "expected one of [',', ';']",
		yyXError{475, -1}: "expected one of [',', ';']",
		yyXError{477, -1}: "expected one of [',', ';']",
		yyXError{295, -1}: "expected one of [',', '=', '}']",
		yyXError{298, -1}: "expected one of [',', '=', '}']",
		yyXError{178, -1}: "expected one of [',', ']']",
		yyXError{264, -1}: "expected one of [',', '}']",
		yyXError{270, -1}: "expected one of [',', '}']",
		yyXError{297, -1}: "expected one of [',', '}']",
		yyXError{300, -1}: "expected one of [',', '}']",
		yyXError{304, -1}: "expected one of [',', '}']",
		yyXError{254, -1}: "expected one of ['.', '=', '[']",
		yyXError{257, -1}: "expected one of ['.', '=', '[']",
		yyXError{259, -1}: "expected one of ['.', '=', '[']",
		yyXError{261, -1}: "expected one of ['.', '=', '[']",
		yyXError{361, -1}: "expected one of [':', string literal]",
		yyXError{494, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		yyXError{503, -1}: "expected one of ['\\n', ppother]",
		yyXError{506, -1}: "expected one of ['\\n', ppother]",
		yyXError{508, -1}: "expected one of ['\\n', ppother]",
		yyXError{346, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{348, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{33, -1}:  "expected one of ['{', identifier]",
		yyXError{34, -1}:  "expected one of ['{', identifier]",
		yyXError{311, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{314, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{323, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{327, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{532, -1}: "expected one of [..., identifier]",
		yyXError{79, -1}:  "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{99, -1}:  "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{395, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{397, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{250, -1}: "expected optional comma or one of [',', '}']",
		yyXError{265, -1}: "expected optional comma or one of [',', '}']",
		yyXError{296, -1}: "expected optional comma or one of [',', '}']",
		yyXError{8, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{434, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{440, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{423, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{432, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{438, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{216, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{205, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{88, -1}:  "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{92, -1}:  "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{541, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{546, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{549, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{556, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{562, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{206, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{32, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{35, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{345, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{190, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{233, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{234, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{77, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{78, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{495, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{499, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{80, -1}:  "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{355, -1}: "expected optional volatile or one of ['(', goto, volatile]",
		yyXError{46, -1}:  "expected optional volatile or one of ['(', volatile]",
		yyXError{229, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{207, -1}: "expected parameter type list or one of [_Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{237, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{482, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{521, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{527, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{530, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{534, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{537, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{76, -1}:  "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{418, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{436, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{442, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{452, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{456, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{460, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{462, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{467, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{470, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{472, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{51, -1}:  "expected string literal",
		yyXError{364, -1}: "expected string literal",
		yyXError{369, -1}: "expected string literal",
		yyXError{374, -1}: "expected string literal",
		yyXError{308, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{309, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{310, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{312, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		yyXError{324, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{511, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{488, -1}: "expected token list or ppother",
		yyXError{491, -1}: "expected token list or ppother",
		yyXError{497, -1}: "expected token list or ppother",
		yyXError{498, -1}: "expected token list or ppother",
		yyXError{501, -1}: "expected token list or ppother",
		yyXError{554, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [_Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{5, -1}:   "expected translation unit or one of [_Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{74, -1}:  "expected type name or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{95, -1}:  "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{218, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{444, -1}: "expected while",
		yyXError{3, 50}:   "unexpected EOF",
		yyXError{2, 50}:   "unexpected EOF",
		yyXError{4, 50}:   "unexpected EOF",
	}

	yyParseTab = [568][]uint16{
		// 0
		{215: 318, 226: 317, 228: 316, 231: 319},
		{50: 315},
		{83: 314, 314, 99: 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 198: 797},
		{312, 312, 312, 312, 312, 312, 312, 312, 11: 312, 13: 312, 312, 312, 312, 312, 312, 312, 312, 312, 205: 795},
		{22: 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 43: 310, 310, 310, 310, 310, 310, 310, 56: 310, 125: 310, 206: 320},
		// 5
		{22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 56: 361, 112: 325, 116: 344, 347, 343, 324, 121: 359, 322, 326, 125: 362, 323, 143: 358, 185: 356, 357, 194: 360, 232: 321},
		{22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 309, 56: 361, 112: 325, 116: 344, 347, 343, 324, 121: 359, 322, 326, 125: 362, 323, 143: 358, 185: 794, 357, 194: 360},
		{157, 395, 157, 9: 229, 135: 634, 633, 138: 656, 169: 654, 189: 655, 653},
		{233, 233, 233, 8: 233, 233, 233, 12: 233, 22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 112: 325, 116: 344, 347, 343, 324, 122: 649, 326, 126: 323, 156: 652},
		{233, 233, 233, 8: 233, 233, 233, 12: 233, 22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 112: 325, 116: 344, 347, 343, 324, 122: 649, 326, 126: 323, 156: 651},
		// 10
		{233, 233, 233, 8: 233, 233, 233, 12: 233, 22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 112: 325, 116: 344, 347, 343, 324, 122: 649, 326, 126: 323, 156: 650},
		{233, 233, 233, 8: 233, 233, 233, 12: 233, 22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 112: 325, 116: 344, 347, 343, 324, 122: 649, 326, 126: 323, 156: 648},
		{224, 224, 224, 8: 224, 224, 224, 12: 224, 22: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 43: 224, 224, 224, 224, 224, 224, 224},
		{223, 223, 223, 8: 223, 223, 223, 12: 223, 22: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 43: 223, 223, 223, 223, 223, 223, 223},
		{222, 222, 222, 8: 222, 222, 222, 12: 222, 22: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 43: 222, 222, 222, 222, 222, 222, 222},
		// 15
		{221, 221, 221, 8: 221, 221, 221, 12: 221, 22: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 43: 221, 221, 221, 221, 221, 221, 221},
		{220, 220, 220, 8: 220, 220, 220, 12: 220, 22: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 43: 220, 220, 220, 220, 220, 220, 220},
		{219, 219, 219, 8: 219, 219, 219, 12: 219, 22: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 43: 219, 219, 219, 219, 219, 219, 219},
		{218, 218, 218, 8: 218, 218, 218, 12: 218, 22: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 43: 218, 218, 218, 218, 218, 218, 218},
		{217, 217, 217, 8: 217, 217, 217, 12: 217, 22: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 43: 217, 217, 217, 217, 217, 217, 217},
		// 20
		{216, 216, 216, 8: 216, 216, 216, 12: 216, 22: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 43: 216, 216, 216, 216, 216, 216, 216},
		{215, 215, 215, 8: 215, 215, 215, 12: 215, 22: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 43: 215, 215, 215, 215, 215, 215, 215},
		{214, 214, 214, 8: 214, 214, 214, 12: 214, 22: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 43: 214, 214, 214, 214, 214, 214, 214},
		{213, 213, 213, 8: 213, 213, 213, 12: 213, 22: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 43: 213, 213, 213, 213, 213, 213, 213},
		{212, 212, 212, 8: 212, 212, 212, 12: 212, 22: 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 43: 212, 212, 212, 212, 212, 212, 212},
		// 25
		{211, 211, 211, 8: 211, 211, 211, 12: 211, 22: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 43: 211, 211, 211, 211, 211, 211, 211},
		{210, 210, 210, 8: 210, 210, 210, 12: 210, 22: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 43: 210, 210, 210, 210, 210, 210, 210},
		{209, 209, 209, 8: 209, 209, 209, 12: 209, 22: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 43: 209, 209, 209, 209, 209, 209, 209},
		{208, 208, 208, 8: 208, 208, 208, 12: 208, 22: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 43: 208, 208, 208, 208, 208, 208, 208},
		{207, 207, 207, 8: 207, 207, 207, 12: 207, 22: 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 43: 207, 207, 207, 207, 207, 207, 207},
		// 30
		{206, 206, 206, 8: 206, 206, 206, 12: 206, 22: 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 43: 206, 206, 206, 206, 206, 206, 206},
		{643},
		{2: 621, 54: 139, 188: 620},
		{2: 199, 54: 199},
		{2: 198, 54: 198},
		// 35
		{2: 607, 54: 139, 188: 606},
		{176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 43: 176, 176, 176, 176, 176, 176, 176, 53: 176},
		{175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 43: 175, 175, 175, 175, 175, 175, 175, 53: 175},
		{174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 43: 174, 174, 174, 174, 174, 174, 174, 53: 174},
		{173, 173, 173, 8: 173, 173, 173, 12: 173, 22: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 43: 173, 173, 173, 173, 173, 173, 173},
		// 40
		{172, 172, 172, 8: 172, 172, 172, 12: 172, 22: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 43: 172, 172, 172, 172, 172, 172, 172},
		{22: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 43: 78, 78, 78, 78, 78, 78, 78, 78, 56: 78, 125: 78},
		{22: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 43: 76, 76, 76, 76, 76, 76, 76, 76, 56: 76, 125: 76},
		{22: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 43: 75, 75, 75, 75, 75, 75, 75, 75, 56: 75, 125: 75},
		{9: 605},
		// 45
		{22: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 43: 73, 73, 73, 73, 73, 73, 73, 73, 56: 73, 125: 73},
		{58, 22: 599, 197: 598},
		{363},
		{240, 240, 240, 240, 240, 240, 240, 240, 11: 240, 13: 240, 240, 240, 240, 240, 240, 240, 240, 240, 144: 364, 365},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 378},
		// 50
		{10: 366},
		{11: 367},
		{8: 368},
		{9: 369},
		{22: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 43: 44, 44, 44, 44, 44, 44, 44, 44, 56: 44, 125: 44},
		// 55
		{303, 303, 3: 303, 303, 303, 303, 303, 303, 303, 303, 12: 303, 41: 303, 303, 50: 303, 303, 53: 303, 55: 303, 57: 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303},
		{302, 302, 3: 302, 302, 302, 302, 302, 302, 302, 302, 12: 302, 41: 302, 302, 50: 302, 302, 53: 302, 55: 302, 57: 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302},
		{301, 301, 3: 301, 301, 301, 301, 301, 301, 301, 301, 12: 301, 41: 301, 301, 50: 301, 301, 53: 301, 55: 301, 57: 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301},
		{300, 300, 3: 300, 300, 300, 300, 300, 300, 300, 300, 12: 300, 41: 300, 300, 50: 300, 300, 53: 300, 55: 300, 57: 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300},
		{299, 299, 3: 299, 299, 299, 299, 299, 299, 299, 299, 12: 299, 41: 299, 299, 50: 299, 299, 53: 299, 55: 299, 57: 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299},
		// 60
		{298, 298, 3: 298, 298, 298, 298, 298, 298, 298, 298, 12: 298, 41: 298, 298, 50: 298, 298, 53: 298, 55: 298, 57: 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298},
		{297, 297, 3: 297, 297, 297, 297, 297, 297, 297, 297, 12: 297, 41: 297, 297, 50: 297, 297, 53: 297, 55: 297, 57: 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 52: 461, 113: 560, 152: 391, 159: 596},
		{414, 419, 3: 432, 422, 423, 418, 417, 9: 239, 239, 12: 413, 41: 239, 239, 50: 239, 438, 53: 239, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 595},
		// 65
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 594},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 593},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 502},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 592},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 591},
		// 70
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 590},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 589},
		{558, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 559},
		{389},
		{22: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 152: 391, 159: 390},
		// 75
		{8: 557},
		{22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 112: 393, 116: 344, 347, 343, 392, 150: 394},
		{191, 191, 191, 8: 191, 191, 12: 191, 22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 191, 112: 393, 116: 344, 347, 343, 392, 150: 555, 193: 556},
		{191, 191, 191, 8: 191, 191, 12: 191, 22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 191, 112: 393, 116: 344, 347, 343, 392, 150: 555, 193: 554},
		{157, 395, 8: 133, 12: 157, 135: 396, 398, 153: 399, 175: 397},
		// 80
		{153, 153, 153, 8: 153, 10: 153, 12: 153, 22: 353, 351, 352, 112: 406, 151: 410, 160: 552},
		{156, 2: 156, 8: 135, 10: 135, 12: 156},
		{8: 136},
		{401, 12: 121, 179: 400, 402},
		{8: 132, 10: 132},
		// 85
		{548, 8: 134, 10: 134, 12: 120},
		{157, 395, 8: 125, 12: 157, 22: 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 43: 125, 125, 125, 125, 125, 125, 125, 135: 396, 398, 153: 504, 173: 505},
		{12: 403},
		{377, 405, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 353, 351, 352, 43: 409, 52: 404, 246, 112: 406, 151: 407, 167: 408},
		{414, 419, 3: 432, 422, 423, 418, 417, 12: 413, 51: 438, 53: 245, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		// 90
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 502, 503},
		{155, 155, 155, 155, 155, 155, 155, 155, 155, 10: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 43: 155, 53: 155},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 353, 351, 352, 43: 498, 52: 404, 246, 112: 495, 167: 497},
		{53: 496},
		{153, 153, 153, 153, 153, 153, 153, 153, 11: 153, 13: 153, 153, 153, 153, 153, 153, 153, 153, 153, 353, 351, 352, 112: 406, 151: 410, 160: 411},
		// 95
		{152, 152, 152, 152, 152, 152, 152, 152, 152, 10: 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 353, 351, 352, 112: 495},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 412},
		{414, 419, 3: 432, 422, 423, 418, 417, 12: 413, 51: 438, 53: 449, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 113: 493},
		{377, 382, 370, 381, 383, 384, 380, 379, 305, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 487, 211: 488, 489},
		// 100
		{2: 486},
		{2: 485},
		{291, 291, 3: 291, 291, 291, 291, 291, 291, 291, 291, 12: 291, 41: 291, 291, 50: 291, 291, 53: 291, 55: 291, 57: 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291},
		{290, 290, 3: 290, 290, 290, 290, 290, 290, 290, 290, 12: 290, 41: 290, 290, 50: 290, 290, 53: 290, 55: 290, 57: 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 484},
		// 105
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 483},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 482},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 481},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 480},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 479},
		// 110
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 478},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 477},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 476},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 475},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 474},
		// 115
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 473},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 472},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 471},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 470},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 469},
		// 120
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 468},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 467},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 113: 462},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 460},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 459},
		// 125
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 458},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 457},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 456},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 455},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 454},
		// 130
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 453},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 452},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 451},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 450},
		{128, 8: 128, 10: 128, 12: 128},
		// 135
		{414, 419, 3: 432, 422, 423, 418, 417, 248, 248, 248, 12: 413, 41: 248, 248, 50: 248, 438, 53: 248, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{414, 419, 3: 432, 422, 423, 418, 417, 249, 249, 249, 12: 413, 41: 249, 249, 50: 249, 438, 53: 249, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{414, 419, 3: 432, 422, 423, 418, 417, 250, 250, 250, 12: 413, 41: 250, 250, 50: 250, 438, 53: 250, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{414, 419, 3: 432, 422, 423, 418, 417, 251, 251, 251, 12: 413, 41: 251, 251, 50: 251, 438, 53: 251, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{414, 419, 3: 432, 422, 423, 418, 417, 252, 252, 252, 12: 413, 41: 252, 252, 50: 252, 438, 53: 252, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		// 140
		{414, 419, 3: 432, 422, 423, 418, 417, 253, 253, 253, 12: 413, 41: 253, 253, 50: 253, 438, 53: 253, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{414, 419, 3: 432, 422, 423, 418, 417, 254, 254, 254, 12: 413, 41: 254, 254, 50: 254, 438, 53: 254, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{414, 419, 3: 432, 422, 423, 418, 417, 255, 255, 255, 12: 413, 41: 255, 255, 50: 255, 438, 53: 255, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{414, 419, 3: 432, 422, 423, 418, 417, 256, 256, 256, 12: 413, 41: 256, 256, 50: 256, 438, 53: 256, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{414, 419, 3: 432, 422, 423, 418, 417, 257, 257, 257, 12: 413, 41: 257, 257, 50: 257, 438, 53: 257, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		// 145
		{414, 419, 3: 432, 422, 423, 418, 417, 258, 258, 258, 12: 413, 41: 258, 258, 50: 258, 438, 53: 258, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{414, 419, 3: 432, 422, 423, 418, 417, 244, 244, 244, 12: 413, 41: 244, 51: 438, 53: 244, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{10: 464, 41: 463},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 466},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 465},
		// 150
		{414, 419, 3: 432, 422, 423, 418, 417, 243, 243, 243, 12: 413, 41: 243, 51: 438, 53: 243, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{414, 419, 3: 432, 422, 423, 418, 417, 259, 259, 259, 12: 413, 41: 259, 259, 50: 259, 259, 53: 259, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 259, 435, 259, 416, 259, 430, 429, 428, 424, 259, 259, 259, 431, 259, 436, 425, 259, 259, 259},
		{414, 419, 3: 432, 422, 423, 418, 417, 260, 260, 260, 12: 413, 41: 260, 260, 50: 260, 260, 53: 260, 55: 415, 57: 421, 420, 426, 427, 260, 433, 434, 260, 435, 260, 416, 260, 430, 429, 428, 424, 260, 260, 260, 431, 260, 260, 425, 260, 260, 260},
		{414, 419, 3: 432, 422, 423, 418, 417, 261, 261, 261, 12: 413, 41: 261, 261, 50: 261, 261, 53: 261, 55: 415, 57: 421, 420, 426, 427, 261, 433, 434, 261, 261, 261, 416, 261, 430, 429, 428, 424, 261, 261, 261, 431, 261, 261, 425, 261, 261, 261},
		{414, 419, 3: 432, 422, 423, 418, 417, 262, 262, 262, 12: 413, 41: 262, 262, 50: 262, 262, 53: 262, 55: 415, 57: 421, 420, 426, 427, 262, 433, 262, 262, 262, 262, 416, 262, 430, 429, 428, 424, 262, 262, 262, 431, 262, 262, 425, 262, 262, 262},
		// 155
		{414, 419, 3: 432, 422, 423, 418, 417, 263, 263, 263, 12: 413, 41: 263, 263, 50: 263, 263, 53: 263, 55: 415, 57: 421, 420, 426, 427, 263, 263, 263, 263, 263, 263, 416, 263, 430, 429, 428, 424, 263, 263, 263, 431, 263, 263, 425, 263, 263, 263},
		{414, 419, 3: 264, 422, 423, 418, 417, 264, 264, 264, 12: 413, 41: 264, 264, 50: 264, 264, 53: 264, 55: 415, 57: 421, 420, 426, 427, 264, 264, 264, 264, 264, 264, 416, 264, 430, 429, 428, 424, 264, 264, 264, 431, 264, 264, 425, 264, 264, 264},
		{414, 419, 3: 265, 422, 423, 418, 417, 265, 265, 265, 12: 413, 41: 265, 265, 50: 265, 265, 53: 265, 55: 415, 57: 421, 420, 426, 427, 265, 265, 265, 265, 265, 265, 416, 265, 265, 429, 428, 424, 265, 265, 265, 265, 265, 265, 425, 265, 265, 265},
		{414, 419, 3: 266, 422, 423, 418, 417, 266, 266, 266, 12: 413, 41: 266, 266, 50: 266, 266, 53: 266, 55: 415, 57: 421, 420, 426, 427, 266, 266, 266, 266, 266, 266, 416, 266, 266, 429, 428, 424, 266, 266, 266, 266, 266, 266, 425, 266, 266, 266},
		{414, 419, 3: 267, 422, 423, 418, 417, 267, 267, 267, 12: 413, 41: 267, 267, 50: 267, 267, 53: 267, 55: 415, 57: 421, 420, 267, 267, 267, 267, 267, 267, 267, 267, 416, 267, 267, 267, 267, 424, 267, 267, 267, 267, 267, 267, 425, 267, 267, 267},
		// 160
		{414, 419, 3: 268, 422, 423, 418, 417, 268, 268, 268, 12: 413, 41: 268, 268, 50: 268, 268, 53: 268, 55: 415, 57: 421, 420, 268, 268, 268, 268, 268, 268, 268, 268, 416, 268, 268, 268, 268, 424, 268, 268, 268, 268, 268, 268, 425, 268, 268, 268},
		{414, 419, 3: 269, 422, 423, 418, 417, 269, 269, 269, 12: 413, 41: 269, 269, 50: 269, 269, 53: 269, 55: 415, 57: 421, 420, 269, 269, 269, 269, 269, 269, 269, 269, 416, 269, 269, 269, 269, 424, 269, 269, 269, 269, 269, 269, 425, 269, 269, 269},
		{414, 419, 3: 270, 422, 423, 418, 417, 270, 270, 270, 12: 413, 41: 270, 270, 50: 270, 270, 53: 270, 55: 415, 57: 421, 420, 270, 270, 270, 270, 270, 270, 270, 270, 416, 270, 270, 270, 270, 424, 270, 270, 270, 270, 270, 270, 425, 270, 270, 270},
		{414, 419, 3: 271, 422, 423, 418, 417, 271, 271, 271, 12: 413, 41: 271, 271, 50: 271, 271, 53: 271, 55: 415, 57: 421, 420, 271, 271, 271, 271, 271, 271, 271, 271, 416, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{414, 419, 3: 272, 422, 423, 418, 417, 272, 272, 272, 12: 413, 41: 272, 272, 50: 272, 272, 53: 272, 55: 415, 57: 421, 420, 272, 272, 272, 272, 272, 272, 272, 272, 416, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272},
		// 165
		{414, 419, 3: 273, 273, 273, 418, 417, 273, 273, 273, 12: 413, 41: 273, 273, 50: 273, 273, 53: 273, 55: 415, 57: 421, 420, 273, 273, 273, 273, 273, 273, 273, 273, 416, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273},
		{414, 419, 3: 274, 274, 274, 418, 417, 274, 274, 274, 12: 413, 41: 274, 274, 50: 274, 274, 53: 274, 55: 415, 57: 421, 420, 274, 274, 274, 274, 274, 274, 274, 274, 416, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274},
		{414, 275, 3: 275, 275, 275, 418, 417, 275, 275, 275, 12: 413, 41: 275, 275, 50: 275, 275, 53: 275, 55: 415, 57: 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 416, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275},
		{414, 276, 3: 276, 276, 276, 418, 417, 276, 276, 276, 12: 413, 41: 276, 276, 50: 276, 276, 53: 276, 55: 415, 57: 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 416, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276},
		{414, 277, 3: 277, 277, 277, 418, 417, 277, 277, 277, 12: 413, 41: 277, 277, 50: 277, 277, 53: 277, 55: 415, 57: 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 416, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277},
		// 170
		{292, 292, 3: 292, 292, 292, 292, 292, 292, 292, 292, 12: 292, 41: 292, 292, 50: 292, 292, 53: 292, 55: 292, 57: 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292},
		{293, 293, 3: 293, 293, 293, 293, 293, 293, 293, 293, 12: 293, 41: 293, 293, 50: 293, 293, 53: 293, 55: 293, 57: 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293},
		{414, 419, 3: 432, 422, 423, 418, 417, 307, 10: 307, 12: 413, 51: 438, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{8: 304, 10: 491},
		{8: 490},
		// 175
		{294, 294, 3: 294, 294, 294, 294, 294, 294, 294, 294, 12: 294, 41: 294, 294, 50: 294, 294, 53: 294, 55: 294, 57: 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 492},
		{414, 419, 3: 432, 422, 423, 418, 417, 306, 10: 306, 12: 413, 51: 438, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{10: 464, 53: 494},
		{295, 295, 3: 295, 295, 295, 295, 295, 295, 295, 295, 12: 295, 41: 295, 295, 50: 295, 295, 53: 295, 55: 295, 57: 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295},
		// 180
		{154, 154, 154, 154, 154, 154, 154, 154, 154, 10: 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 43: 154, 53: 154},
		{130, 8: 130, 10: 130, 12: 130},
		{53: 501},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 499},
		{414, 419, 3: 432, 422, 423, 418, 417, 12: 413, 51: 438, 53: 500, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		// 185
		{127, 8: 127, 10: 127, 12: 127},
		{129, 8: 129, 10: 129, 12: 129},
		{414, 285, 3: 285, 285, 285, 418, 417, 285, 285, 285, 12: 413, 41: 285, 285, 50: 285, 285, 53: 285, 55: 415, 57: 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 416, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285},
		{126, 8: 126, 10: 126, 12: 126},
		{8: 547},
		// 190
		{8: 149, 22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 112: 325, 116: 344, 347, 343, 324, 122: 509, 326, 126: 323, 158: 508, 171: 506, 507, 192: 510},
		{8: 151, 10: 544},
		{8: 148},
		{8: 147, 10: 147},
		{157, 395, 157, 8: 133, 10: 133, 12: 157, 135: 396, 512, 138: 513, 153: 399, 175: 514},
		// 195
		{8: 511},
		{124, 8: 124, 10: 124, 12: 124},
		{517, 2: 516, 12: 121, 179: 400, 402, 515},
		{8: 145, 10: 145},
		{8: 144, 10: 144},
		// 200
		{521, 8: 171, 171, 171, 12: 520, 22: 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 43: 171, 171, 171, 171, 171, 171, 171, 51: 171, 54: 171, 56: 171},
		{168, 8: 168, 168, 168, 12: 168, 22: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 43: 168, 168, 168, 168, 168, 168, 168, 51: 168, 54: 168, 56: 168},
		{157, 395, 157, 8: 125, 12: 157, 22: 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 43: 125, 125, 125, 125, 125, 125, 125, 135: 396, 512, 138: 518, 153: 504, 173: 505},
		{8: 519},
		{167, 8: 167, 167, 167, 12: 167, 22: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 43: 167, 167, 167, 167, 167, 167, 167, 51: 167, 54: 167, 56: 167},
		// 205
		{153, 153, 153, 153, 153, 153, 153, 153, 11: 153, 13: 153, 153, 153, 153, 153, 153, 153, 153, 153, 353, 351, 352, 43: 532, 53: 153, 112: 406, 151: 533, 160: 531},
		{2: 524, 8: 141, 22: 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 43: 162, 162, 162, 162, 162, 162, 162, 168: 525, 187: 523, 210: 522},
		{22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 112: 325, 116: 344, 347, 343, 324, 122: 509, 326, 126: 323, 158: 508, 171: 506, 529},
		{8: 528},
		{8: 143, 10: 143, 146: 143},
		// 210
		{8: 140, 10: 526},
		{2: 527},
		{8: 142, 10: 142, 146: 142},
		{160, 8: 160, 160, 160, 12: 160, 22: 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 43: 160, 160, 160, 160, 160, 160, 160, 51: 160, 54: 160, 56: 160},
		{8: 530},
		// 215
		{161, 8: 161, 161, 161, 12: 161, 22: 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 43: 161, 161, 161, 161, 161, 161, 161, 51: 161, 54: 161, 56: 161},
		{377, 540, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 404, 246, 167: 541},
		{153, 153, 153, 153, 153, 153, 153, 153, 11: 153, 13: 153, 153, 153, 153, 153, 153, 153, 153, 153, 353, 351, 352, 112: 406, 151: 410, 160: 537},
		{152, 152, 152, 152, 152, 152, 152, 152, 11: 152, 13: 152, 152, 152, 152, 152, 152, 152, 152, 152, 353, 351, 352, 43: 534, 53: 152, 112: 495},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 535},
		// 220
		{414, 419, 3: 432, 422, 423, 418, 417, 12: 413, 51: 438, 53: 536, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{164, 8: 164, 164, 164, 12: 164, 22: 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 43: 164, 164, 164, 164, 164, 164, 164, 51: 164, 54: 164, 56: 164},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 538},
		{414, 419, 3: 432, 422, 423, 418, 417, 12: 413, 51: 438, 53: 539, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{165, 8: 165, 165, 165, 12: 165, 22: 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 43: 165, 165, 165, 165, 165, 165, 165, 51: 165, 54: 165, 56: 165},
		// 225
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 502, 543},
		{53: 542},
		{166, 8: 166, 166, 166, 12: 166, 22: 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 43: 166, 166, 166, 166, 166, 166, 166, 51: 166, 54: 166, 56: 166},
		{163, 8: 163, 163, 163, 12: 163, 22: 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 43: 163, 163, 163, 163, 163, 163, 163, 51: 163, 54: 163, 56: 163},
		{22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 112: 325, 116: 344, 347, 343, 324, 122: 509, 326, 126: 323, 146: 545, 158: 546},
		// 230
		{8: 150},
		{8: 146, 10: 146},
		{131, 8: 131, 10: 131, 12: 131},
		{8: 123, 22: 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 43: 123, 123, 123, 123, 123, 123, 123, 199: 549},
		{8: 149, 22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 112: 325, 116: 344, 347, 343, 324, 122: 509, 326, 126: 323, 158: 508, 171: 506, 507, 192: 550},
		// 235
		{8: 551},
		{122, 8: 122, 10: 122, 12: 122},
		{159, 395, 159, 8: 159, 10: 159, 12: 159, 135: 553},
		{158, 2: 158, 8: 158, 10: 158, 12: 158},
		{192, 192, 192, 8: 192, 192, 12: 192, 41: 192},
		// 240
		{190, 190, 190, 8: 190, 190, 12: 190, 41: 190},
		{193, 193, 193, 8: 193, 193, 12: 193, 41: 193},
		{247, 247, 3: 247, 247, 247, 247, 247, 247, 247, 247, 12: 247, 41: 247, 247, 50: 247, 247, 53: 247, 55: 247, 57: 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 52: 461, 113: 560, 152: 391, 159: 561},
		{414, 280, 3: 280, 280, 280, 418, 417, 280, 280, 280, 12: 413, 41: 280, 280, 50: 280, 280, 53: 280, 55: 415, 57: 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 416, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280},
		// 245
		{8: 588, 10: 464},
		{8: 562},
		{377, 279, 370, 279, 279, 279, 380, 379, 279, 279, 279, 376, 279, 386, 385, 388, 371, 372, 373, 374, 375, 387, 41: 279, 279, 50: 279, 279, 563, 279, 564, 279, 57: 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279},
		{414, 278, 3: 278, 278, 278, 418, 417, 278, 278, 278, 12: 413, 41: 278, 278, 50: 278, 278, 53: 278, 55: 415, 57: 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 416, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278},
		{114, 114, 114, 114, 114, 114, 114, 114, 11: 114, 570, 114, 114, 114, 114, 114, 114, 114, 114, 114, 54: 114, 571, 157: 569, 164: 568, 566, 567, 191: 565},
		// 250
		{10: 581, 42: 185, 163: 586},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 577, 54: 578, 170: 579},
		{12: 570, 51: 575, 55: 571, 157: 576},
		{113, 113, 113, 113, 113, 113, 113, 113, 11: 113, 13: 113, 113, 113, 113, 113, 113, 113, 113, 113, 54: 113},
		{12: 112, 51: 112, 55: 112},
		// 255
		{240, 240, 240, 240, 240, 240, 240, 240, 11: 240, 13: 240, 240, 240, 240, 240, 240, 240, 240, 240, 144: 364, 573},
		{2: 572},
		{12: 109, 51: 109, 55: 109},
		{53: 574},
		{12: 110, 51: 110, 55: 110},
		// 260
		{115, 115, 115, 115, 115, 115, 115, 115, 11: 115, 13: 115, 115, 115, 115, 115, 115, 115, 115, 115, 54: 115},
		{12: 111, 51: 111, 55: 111},
		{414, 419, 3: 432, 422, 423, 418, 417, 9: 119, 119, 12: 413, 42: 119, 51: 438, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{114, 114, 114, 114, 114, 114, 114, 114, 11: 114, 570, 114, 114, 114, 114, 114, 114, 114, 114, 114, 54: 114, 571, 157: 569, 164: 568, 566, 567, 191: 580},
		{10: 117, 42: 117},
		// 265
		{10: 581, 42: 185, 163: 582},
		{114, 114, 114, 114, 114, 114, 114, 114, 11: 114, 570, 114, 114, 114, 114, 114, 114, 114, 114, 114, 42: 184, 54: 114, 571, 157: 569, 164: 568, 584, 567},
		{42: 583},
		{9: 118, 118, 42: 118},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 577, 54: 578, 170: 585},
		// 270
		{10: 116, 42: 116},
		{42: 587},
		{289, 289, 3: 289, 289, 289, 289, 289, 289, 289, 289, 12: 289, 41: 289, 289, 50: 289, 289, 53: 289, 55: 289, 57: 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289},
		{296, 296, 3: 296, 296, 296, 296, 296, 296, 296, 296, 12: 296, 41: 296, 296, 50: 296, 296, 53: 296, 55: 296, 57: 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296},
		{414, 281, 3: 281, 281, 281, 418, 417, 281, 281, 281, 12: 413, 41: 281, 281, 50: 281, 281, 53: 281, 55: 415, 57: 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 416, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281},
		// 275
		{414, 282, 3: 282, 282, 282, 418, 417, 282, 282, 282, 12: 413, 41: 282, 282, 50: 282, 282, 53: 282, 55: 415, 57: 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 416, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282},
		{414, 283, 3: 283, 283, 283, 418, 417, 283, 283, 283, 12: 413, 41: 283, 283, 50: 283, 283, 53: 283, 55: 415, 57: 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 416, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283},
		{414, 284, 3: 284, 284, 284, 418, 417, 284, 284, 284, 12: 413, 41: 284, 284, 50: 284, 284, 53: 284, 55: 415, 57: 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 416, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284},
		{414, 286, 3: 286, 286, 286, 418, 417, 286, 286, 286, 12: 413, 41: 286, 286, 50: 286, 286, 53: 286, 55: 415, 57: 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 416, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286},
		{414, 287, 3: 287, 287, 287, 418, 417, 287, 287, 287, 12: 413, 41: 287, 287, 50: 287, 287, 53: 287, 55: 415, 57: 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 416, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287},
		// 280
		{414, 288, 3: 288, 288, 288, 418, 417, 288, 288, 288, 12: 413, 41: 288, 288, 50: 288, 288, 53: 288, 55: 415, 57: 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 416, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288},
		{8: 597},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 563, 54: 564},
		{600},
		{57, 85: 57},
		// 285
		{11: 601, 161: 602},
		{8: 61, 11: 61, 41: 61},
		{8: 604, 11: 603},
		{8: 60, 11: 60, 41: 60},
		{59, 59, 59, 59, 59, 59, 59, 59, 9: 59, 11: 59, 13: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 42: 59, 59, 59, 59, 59, 59, 59, 59, 54: 59, 56: 59, 85: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 111: 59},
		// 290
		{22: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 43: 74, 74, 74, 74, 74, 74, 74, 74, 56: 74, 125: 74},
		{54: 183, 209: 608},
		{181, 181, 181, 8: 181, 181, 181, 12: 181, 22: 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 43: 181, 181, 181, 181, 181, 181, 181, 54: 138},
		{54: 609},
		{2: 610, 183: 613, 612, 223: 611},
		// 295
		{10: 308, 42: 308, 51: 308},
		{10: 616, 42: 185, 163: 617},
		{10: 180, 42: 180},
		{10: 178, 42: 178, 51: 614},
		{240, 240, 240, 240, 240, 240, 240, 240, 11: 240, 13: 240, 240, 240, 240, 240, 240, 240, 240, 240, 144: 364, 615},
		// 300
		{10: 177, 42: 177},
		{2: 610, 42: 184, 183: 613, 619},
		{42: 618},
		{182, 182, 182, 8: 182, 182, 182, 12: 182, 22: 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 43: 182, 182, 182, 182, 182, 182, 182},
		{10: 179, 42: 179},
		// 305
		{54: 203, 207: 622},
		{200, 200, 200, 8: 200, 200, 200, 12: 200, 22: 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 43: 200, 200, 200, 200, 200, 200, 200, 54: 138},
		{54: 623},
		{22: 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 208: 624},
		{22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 112: 393, 116: 344, 347, 343, 392, 150: 627, 195: 626, 229: 625},
		// 310
		{22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 42: 641, 112: 393, 116: 344, 347, 343, 392, 150: 627, 195: 642},
		{22: 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 42: 197},
		{157, 395, 157, 9: 629, 41: 170, 135: 634, 633, 138: 631, 178: 632, 196: 630, 230: 628},
		{9: 638, 639},
		{22: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 42: 194},
		// 315
		{9: 189, 189},
		{9: 187, 187, 41: 169},
		{41: 636},
		{635, 2: 516, 181: 515},
		{156, 2: 156},
		// 320
		{157, 395, 157, 135: 634, 633, 138: 518},
		{240, 240, 240, 240, 240, 240, 240, 240, 11: 240, 13: 240, 240, 240, 240, 240, 240, 240, 240, 240, 144: 364, 637},
		{9: 186, 186},
		{22: 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 42: 195},
		{157, 395, 157, 41: 170, 135: 634, 633, 138: 631, 178: 632, 196: 640},
		// 325
		{9: 188, 188},
		{201, 201, 201, 8: 201, 201, 201, 12: 201, 22: 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 43: 201, 201, 201, 201, 201, 201, 201},
		{22: 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 42: 196},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 52: 644, 152: 391, 159: 645},
		{414, 419, 3: 432, 422, 423, 418, 417, 647, 12: 413, 51: 438, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		// 330
		{8: 646},
		{204, 204, 204, 8: 204, 204, 204, 12: 204, 22: 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 43: 204, 204, 204, 204, 204, 204, 204},
		{205, 205, 205, 8: 205, 205, 205, 12: 205, 22: 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 43: 205, 205, 205, 205, 205, 205, 205},
		{234, 234, 234, 8: 234, 234, 234, 12: 234},
		{232, 232, 232, 8: 232, 232, 232, 12: 232},
		// 335
		{235, 235, 235, 8: 235, 235, 235, 12: 235},
		{236, 236, 236, 8: 236, 236, 236, 12: 236},
		{237, 237, 237, 8: 237, 237, 237, 12: 237},
		{9: 793},
		{9: 231, 231},
		// 340
		{9: 228, 791},
		{9: 227, 227, 22: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 43: 63, 63, 63, 63, 63, 63, 63, 51: 226, 54: 64, 56: 64, 174: 657, 204: 659, 217: 658},
		{51: 789},
		{54: 72, 56: 72, 201: 665},
		{22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 112: 325, 116: 344, 347, 343, 324, 122: 660, 326, 126: 323, 143: 661, 216: 662},
		// 345
		{157, 395, 157, 9: 229, 135: 634, 633, 138: 664, 169: 654, 189: 655, 653},
		{22: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 43: 66, 66, 66, 66, 66, 66, 66, 54: 66, 56: 66},
		{22: 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 54: 62, 56: 62, 112: 325, 116: 344, 347, 343, 324, 122: 660, 326, 126: 323, 143: 663},
		{22: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 43: 65, 65, 65, 65, 65, 65, 65, 54: 65, 56: 65},
		{9: 227, 227, 51: 226, 174: 657},
		// 350
		{54: 70, 56: 68, 202: 667, 668, 224: 666},
		{22: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 43: 71, 71, 71, 71, 71, 71, 71, 71, 56: 71, 125: 71},
		{54: 710, 128: 711},
		{56: 670, 121: 671, 127: 669},
		{9: 709},
		// 355
		{58, 22: 599, 85: 58, 197: 672},
		{49, 49, 49, 49, 49, 49, 49, 49, 9: 49, 11: 49, 13: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 42: 49, 49, 49, 49, 49, 49, 49, 49, 54: 49, 56: 49, 85: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 111: 49},
		{673, 85: 674},
		{11: 601, 161: 699},
		{675},
		// 360
		{11: 601, 161: 676},
		{11: 603, 41: 677},
		{41: 678},
		{11: 53, 682, 154: 680, 679, 162: 681},
		{11: 695},
		// 365
		{8: 55, 10: 55, 41: 55},
		{10: 685, 41: 686},
		{2: 683},
		{53: 684},
		{11: 52},
		// 370
		{11: 53, 682, 154: 694, 679},
		{11: 687, 177: 688},
		{8: 51, 10: 51, 41: 51},
		{10: 689, 41: 690},
		{11: 693},
		// 375
		{2: 524, 168: 691},
		{8: 692, 10: 526},
		{45, 45, 45, 45, 45, 45, 45, 45, 9: 45, 11: 45, 13: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 42: 45, 45, 45, 45, 45, 45, 45, 45, 54: 45, 56: 45, 85: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 111: 45},
		{8: 50, 10: 50, 41: 50},
		{8: 54, 10: 54, 41: 54},
		// 380
		{696},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 697},
		{414, 419, 3: 432, 422, 423, 418, 417, 698, 12: 413, 51: 438, 55: 415, 57: 421, 420, 426, 427, 437, 433, 434, 442, 435, 446, 416, 440, 430, 429, 428, 424, 444, 441, 439, 431, 448, 436, 425, 445, 443, 447},
		{8: 56, 10: 56, 41: 56},
		{8: 604, 11: 603, 41: 700},
		// 385
		{11: 53, 682, 154: 680, 679, 162: 701},
		{8: 702, 10: 685, 41: 703},
		{48, 48, 48, 48, 48, 48, 48, 48, 9: 48, 11: 48, 13: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 42: 48, 48, 48, 48, 48, 48, 48, 48, 54: 48, 56: 48, 85: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 111: 48},
		{11: 53, 682, 154: 680, 679, 162: 704},
		{8: 705, 10: 685, 41: 706},
		// 390
		{47, 47, 47, 47, 47, 47, 47, 47, 9: 47, 11: 47, 13: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 42: 47, 47, 47, 47, 47, 47, 47, 47, 54: 47, 56: 47, 85: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 111: 47},
		{11: 687, 177: 707},
		{8: 708, 10: 689},
		{46, 46, 46, 46, 46, 46, 46, 46, 9: 46, 11: 46, 13: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 42: 46, 46, 46, 46, 46, 46, 46, 46, 54: 46, 56: 46, 85: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 111: 46},
		{22: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 43: 67, 67, 67, 67, 67, 67, 67, 67, 56: 67, 125: 67},
		// 395
		{98, 98, 98, 98, 98, 98, 98, 98, 9: 98, 11: 98, 13: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 42: 98, 98, 98, 98, 98, 98, 98, 98, 54: 98, 56: 98, 85: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 200: 712},
		{22: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 43: 69, 69, 69, 69, 69, 69, 69, 69, 56: 69, 125: 69},
		{377, 382, 713, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 42: 94, 329, 330, 328, 354, 355, 331, 327, 52: 461, 54: 710, 56: 670, 85: 735, 732, 737, 722, 736, 723, 733, 734, 730, 738, 731, 112: 325, 714, 116: 344, 347, 343, 324, 729, 671, 660, 326, 126: 323, 721, 716, 717, 719, 720, 715, 718, 728, 143: 727, 176: 725, 213: 726, 724},
		{303, 303, 3: 303, 303, 303, 303, 303, 9: 303, 303, 12: 303, 41: 787, 51: 303, 55: 303, 57: 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303},
		{8: 241, 241, 464},
		// 400
		{108, 108, 108, 108, 108, 108, 108, 108, 9: 108, 11: 108, 13: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 42: 108, 108, 108, 108, 108, 108, 108, 108, 54: 108, 56: 108, 85: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 111: 108},
		{107, 107, 107, 107, 107, 107, 107, 107, 9: 107, 11: 107, 13: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 42: 107, 107, 107, 107, 107, 107, 107, 107, 54: 107, 56: 107, 85: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 111: 107},
		{106, 106, 106, 106, 106, 106, 106, 106, 9: 106, 11: 106, 13: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 42: 106, 106, 106, 106, 106, 106, 106, 106, 54: 106, 56: 106, 85: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 111: 106},
		{105, 105, 105, 105, 105, 105, 105, 105, 9: 105, 11: 105, 13: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 42: 105, 105, 105, 105, 105, 105, 105, 105, 54: 105, 56: 105, 85: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 111: 105},
		{104, 104, 104, 104, 104, 104, 104, 104, 9: 104, 11: 104, 13: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 42: 104, 104, 104, 104, 104, 104, 104, 104, 54: 104, 56: 104, 85: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 111: 104},
		// 405
		{103, 103, 103, 103, 103, 103, 103, 103, 9: 103, 11: 103, 13: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 42: 103, 103, 103, 103, 103, 103, 103, 103, 54: 103, 56: 103, 85: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 111: 103},
		{102, 102, 102, 102, 102, 102, 102, 102, 9: 102, 11: 102, 13: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 42: 102, 102, 102, 102, 102, 102, 102, 102, 54: 102, 56: 102, 85: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 111: 102},
		{240, 240, 240, 240, 240, 240, 240, 240, 11: 240, 13: 240, 240, 240, 240, 240, 240, 240, 240, 240, 144: 364, 784},
		{41: 782},
		{42: 781},
		// 410
		{96, 96, 96, 96, 96, 96, 96, 96, 9: 96, 11: 96, 13: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 42: 96, 96, 96, 96, 96, 96, 96, 96, 54: 96, 56: 96, 85: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96},
		{377, 382, 713, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 42: 93, 329, 330, 328, 354, 355, 331, 327, 52: 461, 54: 710, 56: 670, 85: 735, 732, 737, 722, 736, 723, 733, 734, 730, 738, 731, 112: 325, 714, 116: 344, 347, 343, 324, 729, 671, 660, 326, 126: 323, 721, 716, 717, 719, 720, 715, 718, 728, 143: 727, 176: 780},
		{92, 92, 92, 92, 92, 92, 92, 92, 9: 92, 11: 92, 13: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 42: 92, 92, 92, 92, 92, 92, 92, 92, 54: 92, 56: 92, 85: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92},
		{91, 91, 91, 91, 91, 91, 91, 91, 9: 91, 11: 91, 13: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 42: 91, 91, 91, 91, 91, 91, 91, 91, 54: 91, 56: 91, 85: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91},
		{9: 779},
		// 415
		{773},
		{769},
		{765},
		{377, 382, 713, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 54: 710, 56: 670, 85: 735, 732, 737, 722, 736, 723, 733, 734, 730, 738, 731, 113: 714, 120: 729, 671, 127: 721, 716, 717, 719, 720, 715, 718, 759},
		{745},
		// 420
		{2: 743},
		{9: 742},
		{9: 741},
		{377, 382, 370, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 113: 714, 120: 739},
		{9: 740},
		// 425
		{79, 79, 79, 79, 79, 79, 79, 79, 9: 79, 11: 79, 13: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 42: 79, 79, 79, 79, 79, 79, 79, 79, 54: 79, 56: 79, 85: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 111: 79},
		{80, 80, 80, 80, 80, 80, 80, 80, 9: 80, 11: 80, 13: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 42: 80, 80, 80, 80, 80, 80, 80, 80, 54: 80, 56: 80, 85: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 111: 80},
		{81, 81, 81, 81, 81, 81, 81, 81, 9: 81, 11: 81, 13: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 42: 81, 81, 81, 81, 81, 81, 81, 81, 54: 81, 56: 81, 85: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 111: 81},
		{9: 744},
		{82, 82, 82, 82, 82, 82, 82, 82, 9: 82, 11: 82, 13: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 42: 82, 82, 82, 82, 82, 82, 82, 82, 54: 82, 56: 82, 85: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 111: 82},
		// 430
		{377, 382, 370, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 353, 351, 352, 341, 333, 342, 338, 350, 337, 335, 336, 334, 339, 348, 345, 346, 349, 340, 332, 43: 329, 330, 328, 354, 355, 331, 327, 52: 461, 112: 325, 714, 116: 344, 347, 343, 324, 746, 122: 660, 326, 126: 323, 143: 747},
		{9: 753},
		{377, 382, 370, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 113: 714, 120: 748},
		{9: 749},
		{377, 382, 370, 381, 383, 384, 380, 379, 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 113: 714, 120: 750},
		// 435
		{8: 751},
		{377, 382, 713, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 54: 710, 56: 670, 85: 735, 732, 737, 722, 736, 723, 733, 734, 730, 738, 731, 113: 714, 120: 729, 671, 127: 721, 716, 717, 719, 720, 715, 718, 752},
		{83, 83, 83, 83, 83, 83, 83, 83, 9: 83, 11: 83, 13: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 42: 83, 83, 83, 83, 83, 83, 83, 83, 54: 83, 56: 83, 85: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 111: 83},
		{377, 382, 370, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 113: 714, 120: 754},
		{9: 755},
		// 440
		{377, 382, 370, 381, 383, 384, 380, 379, 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 113: 714, 120: 756},
		{8: 757},
		{377, 382, 713, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 54: 710, 56: 670, 85: 735, 732, 737, 722, 736, 723, 733, 734, 730, 738, 731, 113: 714, 120: 729, 671, 127: 721, 716, 717, 719, 720, 715, 718, 758},
		{84, 84, 84, 84, 84, 84, 84, 84, 9: 84, 11: 84, 13: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 42: 84, 84, 84, 84, 84, 84, 84, 84, 54: 84, 56: 84, 85: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 111: 84},
		{86: 760},
		// 445
		{761},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 113: 762},
		{8: 763, 10: 464},
		{9: 764},
		{85, 85, 85, 85, 85, 85, 85, 85, 9: 85, 11: 85, 13: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 42: 85, 85, 85, 85, 85, 85, 85, 85, 54: 85, 56: 85, 85: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 111: 85},
		// 450
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 113: 766},
		{8: 767, 10: 464},
		{377, 382, 713, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 54: 710, 56: 670, 85: 735, 732, 737, 722, 736, 723, 733, 734, 730, 738, 731, 113: 714, 120: 729, 671, 127: 721, 716, 717, 719, 720, 715, 718, 768},
		{86, 86, 86, 86, 86, 86, 86, 86, 9: 86, 11: 86, 13: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 42: 86, 86, 86, 86, 86, 86, 86, 86, 54: 86, 56: 86, 85: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 111: 86},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 113: 770},
		// 455
		{8: 771, 10: 464},
		{377, 382, 713, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 54: 710, 56: 670, 85: 735, 732, 737, 722, 736, 723, 733, 734, 730, 738, 731, 113: 714, 120: 729, 671, 127: 721, 716, 717, 719, 720, 715, 718, 772},
		{87, 87, 87, 87, 87, 87, 87, 87, 9: 87, 11: 87, 13: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 42: 87, 87, 87, 87, 87, 87, 87, 87, 54: 87, 56: 87, 85: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 111: 87},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 113: 774},
		{8: 775, 10: 464},
		// 460
		{377, 382, 713, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 54: 710, 56: 670, 85: 735, 732, 737, 722, 736, 723, 733, 734, 730, 738, 731, 113: 714, 120: 729, 671, 127: 721, 716, 717, 719, 720, 715, 718, 776},
		{89, 89, 89, 89, 89, 89, 89, 89, 9: 89, 11: 89, 13: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 42: 89, 89, 89, 89, 89, 89, 89, 89, 54: 89, 56: 89, 85: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 111: 777},
		{377, 382, 713, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 54: 710, 56: 670, 85: 735, 732, 737, 722, 736, 723, 733, 734, 730, 738, 731, 113: 714, 120: 729, 671, 127: 721, 716, 717, 719, 720, 715, 718, 778},
		{88, 88, 88, 88, 88, 88, 88, 88, 9: 88, 11: 88, 13: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 42: 88, 88, 88, 88, 88, 88, 88, 88, 54: 88, 56: 88, 85: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 111: 88},
		{90, 90, 90, 90, 90, 90, 90, 90, 9: 90, 11: 90, 13: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 42: 90, 90, 90, 90, 90, 90, 90, 90, 54: 90, 56: 90, 85: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 111: 90},
		// 465
		{95, 95, 95, 95, 95, 95, 95, 95, 9: 95, 11: 95, 13: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 42: 95, 95, 95, 95, 95, 95, 95, 95, 54: 95, 56: 95, 85: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95},
		{97, 97, 97, 97, 97, 97, 97, 97, 9: 97, 11: 97, 13: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 42: 97, 97, 97, 97, 97, 97, 97, 97, 97, 54: 97, 56: 97, 85: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 111: 97, 125: 97},
		{377, 382, 713, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 54: 710, 56: 670, 85: 735, 732, 737, 722, 736, 723, 733, 734, 730, 738, 731, 113: 714, 120: 729, 671, 127: 721, 716, 717, 719, 720, 715, 718, 783},
		{99, 99, 99, 99, 99, 99, 99, 99, 9: 99, 11: 99, 13: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 42: 99, 99, 99, 99, 99, 99, 99, 99, 54: 99, 56: 99, 85: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 111: 99},
		{41: 785},
		// 470
		{377, 382, 713, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 54: 710, 56: 670, 85: 735, 732, 737, 722, 736, 723, 733, 734, 730, 738, 731, 113: 714, 120: 729, 671, 127: 721, 716, 717, 719, 720, 715, 718, 786},
		{100, 100, 100, 100, 100, 100, 100, 100, 9: 100, 11: 100, 13: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 42: 100, 100, 100, 100, 100, 100, 100, 100, 54: 100, 56: 100, 85: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 111: 100},
		{377, 382, 713, 381, 383, 384, 380, 379, 9: 242, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 461, 54: 710, 56: 670, 85: 735, 732, 737, 722, 736, 723, 733, 734, 730, 738, 731, 113: 714, 120: 729, 671, 127: 721, 716, 717, 719, 720, 715, 718, 788},
		{101, 101, 101, 101, 101, 101, 101, 101, 9: 101, 11: 101, 13: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 42: 101, 101, 101, 101, 101, 101, 101, 101, 54: 101, 56: 101, 85: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 111: 101},
		{377, 382, 370, 381, 383, 384, 380, 379, 11: 376, 13: 386, 385, 388, 371, 372, 373, 374, 375, 387, 52: 577, 54: 578, 170: 790},
		// 475
		{9: 225, 225},
		{157, 395, 157, 135: 634, 633, 138: 664, 169: 792},
		{9: 230, 230},
		{238, 238, 238, 238, 238, 238, 238, 238, 9: 238, 11: 238, 13: 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 42: 238, 238, 238, 238, 238, 238, 238, 238, 238, 54: 238, 56: 238, 85: 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 125: 238},
		{22: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 43: 77, 77, 77, 77, 77, 77, 77, 77, 56: 77, 125: 77},
		// 480
		{240, 240, 240, 240, 240, 240, 240, 240, 11: 240, 13: 240, 240, 240, 240, 240, 240, 240, 240, 240, 144: 364, 796},
		{50: 311},
		{83: 819, 821, 99: 809, 810, 811, 806, 807, 808, 812, 816, 813, 803, 814, 815, 114: 820, 818, 124: 817, 137: 801, 139: 800, 805, 802, 804, 147: 799, 227: 798},
		{50: 313},
		{50: 43, 83: 819, 821, 99: 809, 810, 811, 806, 807, 808, 812, 816, 813, 803, 814, 815, 114: 820, 818, 124: 817, 137: 801, 139: 859, 805, 802, 804},
		// 485
		{50: 42, 83: 42, 42, 96: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{50: 38, 83: 38, 38, 96: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{50: 37, 83: 37, 37, 96: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{84: 821, 114: 881, 818},
		{50: 35, 83: 35, 35, 96: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		// 490
		{96: 28, 28, 869, 182: 867, 218: 868, 866},
		{84: 821, 114: 863, 818},
		{2: 860},
		{2: 855},
		{2: 836, 83: 838, 225: 837},
		// 495
		{83: 819, 821, 114: 820, 818, 124: 835},
		{50: 16, 83: 16, 16, 96: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{84: 821, 114: 833, 818},
		{84: 821, 114: 831, 818},
		{83: 819, 821, 114: 820, 818, 124: 830},
		// 500
		{2: 826},
		{84: 821, 114: 824, 818},
		{50: 7, 83: 7, 7, 96: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{83: 5, 823},
		{50: 4, 83: 4, 4, 96: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		// 505
		{83: 822},
		{83: 2, 2},
		{50: 3, 83: 3, 3, 96: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{83: 1, 1},
		{83: 825},
		// 510
		{50: 8, 83: 8, 8, 96: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{83: 827, 821, 114: 828, 818},
		{50: 12, 83: 12, 12, 96: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{83: 829},
		{50: 9, 83: 9, 9, 96: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		// 515
		{50: 13, 83: 13, 13, 96: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{83: 832},
		{50: 14, 83: 14, 14, 96: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{83: 834},
		{50: 15, 83: 15, 15, 96: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		// 520
		{50: 17, 83: 17, 17, 96: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{83: 819, 821, 114: 820, 818, 124: 844, 149: 854},
		{2: 524, 8: 141, 146: 840, 168: 839, 187: 841},
		{50: 10, 83: 10, 10, 96: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{8: 140, 10: 847, 146: 848},
		// 525
		{8: 845},
		{8: 842},
		{83: 819, 821, 114: 820, 818, 124: 844, 149: 843},
		{50: 18, 83: 18, 18, 96: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{50: 6, 83: 6, 6, 96: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		// 530
		{83: 819, 821, 114: 820, 818, 124: 844, 149: 846},
		{50: 20, 83: 20, 20, 96: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{2: 527, 146: 851},
		{8: 849},
		{83: 819, 821, 114: 820, 818, 124: 844, 149: 850},
		// 535
		{50: 11, 83: 11, 11, 96: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{8: 852},
		{83: 819, 821, 114: 820, 818, 124: 844, 149: 853},
		{50: 19, 83: 19, 19, 96: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{50: 21, 83: 21, 21, 96: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		// 540
		{83: 856},
		{83: 819, 821, 96: 40, 40, 40, 809, 810, 811, 806, 807, 808, 812, 816, 813, 803, 814, 815, 114: 820, 818, 124: 817, 137: 801, 139: 800, 805, 802, 804, 147: 857, 858},
		{83: 819, 821, 96: 39, 39, 39, 809, 810, 811, 806, 807, 808, 812, 816, 813, 803, 814, 815, 114: 820, 818, 124: 817, 137: 801, 139: 859, 805, 802, 804},
		{96: 31, 31, 31},
		{50: 41, 83: 41, 41, 96: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		// 545
		{83: 861},
		{83: 819, 821, 96: 40, 40, 40, 809, 810, 811, 806, 807, 808, 812, 816, 813, 803, 814, 815, 114: 820, 818, 124: 817, 137: 801, 139: 800, 805, 802, 804, 147: 857, 862},
		{96: 32, 32, 32},
		{83: 864},
		{83: 819, 821, 96: 40, 40, 40, 809, 810, 811, 806, 807, 808, 812, 816, 813, 803, 814, 815, 114: 820, 818, 124: 817, 137: 801, 139: 800, 805, 802, 804, 147: 857, 865},
		// 550
		{96: 33, 33, 33},
		{96: 24, 875, 220: 876, 874},
		{96: 30, 30, 30},
		{96: 27, 27, 869, 182: 873},
		{84: 821, 114: 870, 818},
		// 555
		{83: 871},
		{83: 819, 821, 96: 40, 40, 40, 809, 810, 811, 806, 807, 808, 812, 816, 813, 803, 814, 815, 114: 820, 818, 124: 817, 137: 801, 139: 800, 805, 802, 804, 147: 857, 872},
		{96: 26, 26, 26},
		{96: 29, 29, 29},
		{96: 880, 222: 879},
		// 560
		{83: 877},
		{96: 23},
		{83: 819, 821, 96: 40, 99: 809, 810, 811, 806, 807, 808, 812, 816, 813, 803, 814, 815, 114: 820, 818, 124: 817, 137: 801, 139: 800, 805, 802, 804, 147: 857, 878},
		{96: 25},
		{50: 34, 83: 34, 34, 96: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		// 565
		{50: 22, 83: 22, 22, 96: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{83: 882},
		{50: 36, 83: 36, 36, 96: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
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

	if c < 0x7f {
		return __yyfmt__.Sprintf("'%c'", c)
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
	const yyError = 235

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
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
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
			yyVAL.node = &Expression{
				Case:     56,
				Token:    yyS[yypt-3].Token,
				Token2:   yyS[yypt-2].Token,
				TypeName: yyS[yypt-1].node.(*TypeName),
				Token3:   yyS[yypt-0].Token,
			}
		}
	case 69:
		{
			yyVAL.node = (*ExpressionOpt)(nil)
		}
	case 70:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionOpt{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 71:
		{
			yyVAL.node = &ExpressionList{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 72:
		{
			yyVAL.node = &ExpressionList{
				Case:           1,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList),
				Token:          yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 73:
		{
			yyVAL.node = (*ExpressionListOpt)(nil)
		}
	case 74:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionListOpt{
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 75:
		{
			lx := yylex.(*lexer)
			lx.constExprToks = []xc.Token{lx.last}
		}
	case 76:
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
	case 77:
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
	case 78:
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
	case 79:
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
	case 80:
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
	case 81:
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
	case 82:
		{
			yyVAL.node = (*DeclarationSpecifiersOpt)(nil)
		}
	case 83:
		{
			lhs := &DeclarationSpecifiersOpt{
				DeclarationSpecifiers: yyS[yypt-0].node.(*DeclarationSpecifiers),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.DeclarationSpecifiers.attr
			lhs.typeSpecifier = lhs.DeclarationSpecifiers.typeSpecifier
		}
	case 84:
		{
			yyVAL.node = &InitDeclaratorList{
				InitDeclarator: yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 85:
		{
			yyVAL.node = &InitDeclaratorList{
				Case:               1,
				InitDeclaratorList: yyS[yypt-2].node.(*InitDeclaratorList),
				Token:              yyS[yypt-1].Token,
				InitDeclarator:     yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 86:
		{
			yyVAL.node = (*InitDeclaratorListOpt)(nil)
		}
	case 87:
		{
			yyVAL.node = &InitDeclaratorListOpt{
				InitDeclaratorList: yyS[yypt-0].node.(*InitDeclaratorList).reverse(),
			}
		}
	case 88:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
		}
	case 89:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-0].node.(*Declarator)
			d.setFull(lx)
		}
	case 90:
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
	case 91:
		{
			lhs := &StorageClassSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saTypedef
		}
	case 92:
		{
			lhs := &StorageClassSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saExtern
		}
	case 93:
		{
			lhs := &StorageClassSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saStatic
		}
	case 94:
		{
			lhs := &StorageClassSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saAuto
		}
	case 95:
		{
			lhs := &StorageClassSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRegister
		}
	case 96:
		{
			lhs := &TypeSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsVoid)
		}
	case 97:
		{
			lhs := &TypeSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsChar)
		}
	case 98:
		{
			lhs := &TypeSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsShort)
		}
	case 99:
		{
			lhs := &TypeSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsInt)
		}
	case 100:
		{
			lhs := &TypeSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsLong)
		}
	case 101:
		{
			lhs := &TypeSpecifier{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsFloat)
		}
	case 102:
		{
			lhs := &TypeSpecifier{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsDouble)
		}
	case 103:
		{
			lhs := &TypeSpecifier{
				Case:  7,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsSigned)
		}
	case 104:
		{
			lhs := &TypeSpecifier{
				Case:  8,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsUnsigned)
		}
	case 105:
		{
			lhs := &TypeSpecifier{
				Case:  9,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsBool)
		}
	case 106:
		{
			lhs := &TypeSpecifier{
				Case:  10,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsComplex)
		}
	case 107:
		{
			lhs := &TypeSpecifier{
				Case: 11,
				StructOrUnionSpecifier: yyS[yypt-0].node.(*StructOrUnionSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(lhs.StructOrUnionSpecifier.typeSpecifiers())
		}
	case 108:
		{
			lhs := &TypeSpecifier{
				Case:          12,
				EnumSpecifier: yyS[yypt-0].node.(*EnumSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsEnumSpecifier)
		}
	case 109:
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
	case 110:
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
	case 111:
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
	case 112:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareStructTag(o.Token, lx.report)
			}
		}
	case 113:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeMembers)
			lx.scope.isUnion = yyS[yypt-3].node.(*StructOrUnion).Case == 1 // "union"
			lx.scope.prevStructDeclarator = nil
		}
	case 114:
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
	case 115:
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
	case 116:
		{
			yyVAL.node = &StructOrUnion{
				Token: yyS[yypt-0].Token,
			}
		}
	case 117:
		{
			yyVAL.node = &StructOrUnion{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 118:
		{
			yyVAL.node = &StructDeclarationList{
				StructDeclaration: yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 119:
		{
			yyVAL.node = &StructDeclarationList{
				Case: 1,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList),
				StructDeclaration:     yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 120:
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
	case 121:
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
	case 122:
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
	case 123:
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
	case 124:
		{
			yyVAL.node = (*SpecifierQualifierListOpt)(nil)
		}
	case 125:
		{
			lhs := &SpecifierQualifierListOpt{
				SpecifierQualifierList: yyS[yypt-0].node.(*SpecifierQualifierList),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.SpecifierQualifierList.attr
			lhs.typeSpecifier = lhs.SpecifierQualifierList.typeSpecifier
		}
	case 126:
		{
			yyVAL.node = &StructDeclaratorList{
				StructDeclarator: yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 127:
		{
			yyVAL.node = &StructDeclaratorList{
				Case:                 1,
				StructDeclaratorList: yyS[yypt-2].node.(*StructDeclaratorList),
				Token:                yyS[yypt-1].Token,
				StructDeclarator:     yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 128:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.post(lx)
		}
	case 129:
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
			if !IsIntType(e.Type) {
				lx.report.Err(e.Pos(), "bit field width not an integer (have '%s')", e.Type)
				e.Value, e.Type = m.value2(1, m.IntType)
			}
			if o := lhs.DeclaratorOpt; o != nil {
				o.Declarator.setFull(lx)
			}
			lhs.post(lx)
		}
	case 130:
		{
			yyVAL.node = (*CommaOpt)(nil)
		}
	case 131:
		{
			yyVAL.node = &CommaOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 132:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareEnumTag(o.Token, lx.report)
			}
			lx.iota = 0
		}
	case 133:
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
	case 134:
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
	case 135:
		{
			yyVAL.node = &EnumeratorList{
				Enumerator: yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 136:
		{
			yyVAL.node = &EnumeratorList{
				Case:           1,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList),
				Token:          yyS[yypt-1].Token,
				Enumerator:     yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 137:
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
	case 138:
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
	case 139:
		{
			lhs := &TypeQualifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saConst
		}
	case 140:
		{
			lhs := &TypeQualifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRestrict
		}
	case 141:
		{
			lhs := &TypeQualifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saVolatile
		}
	case 142:
		{
			lhs := &FunctionSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saInline
		}
	case 143:
		{
			lhs := &FunctionSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saNoreturn
		}
	case 144:
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
	case 145:
		{
			yyVAL.node = (*DeclaratorOpt)(nil)
		}
	case 146:
		{
			yyVAL.node = &DeclaratorOpt{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
		}
	case 147:
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
	case 148:
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
	case 149:
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
				var err error
				if lhs.elements, err = elements(o.Expression.eval(lx)); err != nil {
					lx.report.Err(o.Expression.Pos(), "%s", err)
				}

			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 150:
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
			var err error
			if lhs.elements, err = elements(lhs.Expression.eval(lx)); err != nil {
				lx.report.Err(lhs.Expression.Pos(), "%s", err)
			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 151:
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
			var err error
			if lhs.elements, err = elements(lhs.Expression.eval(lx)); err != nil {
				lx.report.Err(lhs.Expression.Pos(), "%s", err)
			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 152:
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
	case 153:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 154:
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
	case 155:
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
	case 156:
		{
			yyVAL.node = &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].node.(*TypeQualifierListOpt),
			}
		}
	case 157:
		{
			yyVAL.node = &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].node.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].node.(*Pointer),
			}
		}
	case 158:
		{
			yyVAL.node = (*PointerOpt)(nil)
		}
	case 159:
		{
			yyVAL.node = &PointerOpt{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
		}
	case 160:
		{
			lhs := &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.TypeQualifier.attr
		}
	case 161:
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
	case 162:
		{
			yyVAL.node = (*TypeQualifierListOpt)(nil)
		}
	case 163:
		{
			yyVAL.node = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].node.(*TypeQualifierList).reverse(),
			}
		}
	case 164:
		{
			lhs := &ParameterTypeList{
				ParameterList: yyS[yypt-0].node.(*ParameterList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 165:
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
	case 166:
		{
			yyVAL.node = (*ParameterTypeListOpt)(nil)
		}
	case 167:
		{
			yyVAL.node = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].node.(*ParameterTypeList),
			}
		}
	case 168:
		{
			yyVAL.node = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 169:
		{
			yyVAL.node = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].node.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 170:
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
	case 171:
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
	case 172:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 173:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 174:
		{
			yyVAL.node = (*IdentifierListOpt)(nil)
		}
	case 175:
		{
			yyVAL.node = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 176:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 177:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 178:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeBlock)
		}
	case 179:
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
	case 180:
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
	case 181:
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
	case 182:
		{
			yyVAL.node = (*AbstractDeclaratorOpt)(nil)
		}
	case 183:
		{
			yyVAL.node = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].node.(*AbstractDeclarator),
			}
		}
	case 184:
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
	case 185:
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
	case 186:
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
	case 187:
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
	case 188:
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
	case 189:
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
	case 190:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 191:
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
	case 192:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 193:
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
	case 194:
		{
			yyVAL.node = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 195:
		{
			yyVAL.node = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
		}
	case 196:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 197:
		{
			yyVAL.node = &Initializer{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 198:
		{
			yyVAL.node = &InitializerList{
				DesignationOpt: yyS[yypt-1].node.(*DesignationOpt),
				Initializer:    yyS[yypt-0].node.(*Initializer),
			}
		}
	case 199:
		{
			yyVAL.node = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].node.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].node.(*DesignationOpt),
				Initializer:     yyS[yypt-0].node.(*Initializer),
			}
		}
	case 200:
		{
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 201:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 202:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 203:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 204:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 205:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 206:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 207:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 208:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 209:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 210:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 211:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 212:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 213:
		{
			yyVAL.node = &Statement{
				Case:               6,
				AssemblerStatement: yyS[yypt-0].node.(*AssemblerStatement),
			}
		}
	case 214:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 215:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 216:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 217:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 218:
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
	case 219:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 220:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 221:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 222:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 223:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 224:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 225:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 226:
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
	case 227:
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
	case 228:
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
	case 229:
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
	case 230:
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
	case 231:
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
	case 232:
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
	case 233:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 234:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 235:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 236:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 237:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 238:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 239:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 240:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 241:
		{
			yyVAL.node = &ExternalDeclaration{
				Case: 2,
				BasicAssemblerStatement: yyS[yypt-1].node.(*BasicAssemblerStatement),
				Token: yyS[yypt-0].Token,
			}
		}
	case 242:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:         3,
				StaticAssert: yyS[yypt-0].node.(*StaticAssert),
			}
		}
	case 243:
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
			lx.fnDeclarator = d
		}
	case 244:
		{
			yyVAL.node = &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:          yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 245:
		{
			lx := yylex.(*lexer)
			// Handle __func__, [0], 6.4.2.2.
			id, _ := lx.fnDeclarator.Identifier()
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
	case 246:
		{
			lhs := &FunctionBody{
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
			yyVAL.node = lhs
			lhs.scope = lhs.CompoundStatement.scope
		}
	case 247:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 248:
		{
			lx := yylex.(*lexer)
			lhs := &FunctionBody{
				Case:               1,
				AssemblerStatement: yyS[yypt-1].node.(*AssemblerStatement),
				Token:              yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.scope = lx.scope
			lx.popScope(lx.tokPrev)
		}
	case 249:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 250:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 251:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 252:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 253:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScopePos(lhs.Pos())
		}
	case 254:
		{
			yyVAL.node = &AssemblerInstructions{
				Token: yyS[yypt-0].Token,
			}
		}
	case 255:
		{
			yyVAL.node = &AssemblerInstructions{
				Case: 1,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions),
				Token: yyS[yypt-0].Token,
			}
		}
	case 256:
		{
			yyVAL.node = &BasicAssemblerStatement{
				Token:                 yyS[yypt-4].Token,
				VolatileOpt:           yyS[yypt-3].node.(*VolatileOpt),
				Token2:                yyS[yypt-2].Token,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-0].Token,
			}
		}
	case 257:
		{
			yyVAL.node = (*VolatileOpt)(nil)
		}
	case 258:
		{
			yyVAL.node = &VolatileOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 259:
		{
			yyVAL.node = &AssemblerOperand{
				AssemblerSymbolicNameOpt: yyS[yypt-4].node.(*AssemblerSymbolicNameOpt),
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
		}
	case 260:
		{
			yyVAL.node = &AssemblerOperands{
				AssemblerOperand: yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 261:
		{
			yyVAL.node = &AssemblerOperands{
				Case:              1,
				AssemblerOperands: yyS[yypt-2].node.(*AssemblerOperands),
				Token:             yyS[yypt-1].Token,
				AssemblerOperand:  yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 262:
		{
			yyVAL.node = (*AssemblerSymbolicNameOpt)(nil)
		}
	case 263:
		{
			yyVAL.node = &AssemblerSymbolicNameOpt{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 264:
		{
			yyVAL.node = &Clobbers{
				Token: yyS[yypt-0].Token,
			}
		}
	case 265:
		{
			yyVAL.node = &Clobbers{
				Case:     1,
				Clobbers: yyS[yypt-2].node.(*Clobbers),
				Token:    yyS[yypt-1].Token,
				Token2:   yyS[yypt-0].Token,
			}
		}
	case 266:
		{
			yyVAL.node = &AssemblerStatement{
				BasicAssemblerStatement: yyS[yypt-0].node.(*BasicAssemblerStatement),
			}
		}
	case 267:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  1,
				Token:                 yyS[yypt-6].Token,
				VolatileOpt:           yyS[yypt-5].node.(*VolatileOpt),
				Token2:                yyS[yypt-4].Token,
				AssemblerInstructions: yyS[yypt-3].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-2].Token,
				AssemblerOperands:     yyS[yypt-1].node.(*AssemblerOperands).reverse(),
				Token4:                yyS[yypt-0].Token,
			}
		}
	case 268:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  2,
				Token:                 yyS[yypt-8].Token,
				VolatileOpt:           yyS[yypt-7].node.(*VolatileOpt),
				Token2:                yyS[yypt-6].Token,
				AssemblerInstructions: yyS[yypt-5].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-4].Token,
				AssemblerOperands:     yyS[yypt-3].node.(*AssemblerOperands).reverse(),
				Token4:                yyS[yypt-2].Token,
				AssemblerOperands2:    yyS[yypt-1].node.(*AssemblerOperands).reverse(),
				Token5:                yyS[yypt-0].Token,
			}
		}
	case 269:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  3,
				Token:                 yyS[yypt-10].Token,
				VolatileOpt:           yyS[yypt-9].node.(*VolatileOpt),
				Token2:                yyS[yypt-8].Token,
				AssemblerInstructions: yyS[yypt-7].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-6].Token,
				AssemblerOperands:     yyS[yypt-5].node.(*AssemblerOperands).reverse(),
				Token4:                yyS[yypt-4].Token,
				AssemblerOperands2:    yyS[yypt-3].node.(*AssemblerOperands).reverse(),
				Token5:                yyS[yypt-2].Token,
				Clobbers:              yyS[yypt-1].node.(*Clobbers).reverse(),
				Token6:                yyS[yypt-0].Token,
			}
		}
	case 270:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  4,
				Token:                 yyS[yypt-12].Token,
				VolatileOpt:           yyS[yypt-11].node.(*VolatileOpt),
				Token2:                yyS[yypt-10].Token,
				Token3:                yyS[yypt-9].Token,
				AssemblerInstructions: yyS[yypt-8].node.(*AssemblerInstructions).reverse(),
				Token4:                yyS[yypt-7].Token,
				Token5:                yyS[yypt-6].Token,
				AssemblerOperands:     yyS[yypt-5].node.(*AssemblerOperands).reverse(),
				Token6:                yyS[yypt-4].Token,
				Clobbers:              yyS[yypt-3].node.(*Clobbers).reverse(),
				Token7:                yyS[yypt-2].Token,
				IdentifierList:        yyS[yypt-1].node.(*IdentifierList).reverse(),
				Token8:                yyS[yypt-0].Token,
			}
		}
	case 271:
		{
			lx := yylex.(*lexer)
			lhs := &StaticAssert{
				Token:              yyS[yypt-6].Token,
				Token2:             yyS[yypt-5].Token,
				ConstantExpression: yyS[yypt-4].node.(*ConstantExpression),
				Token3:             yyS[yypt-3].Token,
				Token4:             yyS[yypt-2].Token,
				Token5:             yyS[yypt-1].Token,
				Token6:             yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			ce := lhs.ConstantExpression
			if ce.Type == nil || ce.Type.Kind() == Undefined || ce.Value == nil || !IsIntType(ce.Type) {
				lx.report.Err(ce.Pos(), "invalid static assert expression (have '%v')", ce.Type)
				break
			}

			if !isNonZero(ce.Value) {
				lx.report.ErrTok(lhs.Token, "%s", lhs.Token4.S())
			}
		}
	case 272:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 273:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 274:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 275:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 276:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 277:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 278:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 279:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 280:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 281:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 282:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 283:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 284:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 285:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 286:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 287:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 288:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 289:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 290:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 291:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 292:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 293:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 294:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 295:
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
	case 296:
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
	case 297:
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
	case 298:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 299:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 300:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 301:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 302:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 303:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 304:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:            10,
				Token:           yyS[yypt-5].Token,
				Token2:          yyS[yypt-4].Token,
				IdentifierList:  yyS[yypt-3].node.(*IdentifierList).reverse(),
				Token3:          yyS[yypt-2].Token,
				Token4:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableDefineOmitCommaBeforeDDD {
				lx.report.ErrTok(lhs.Token4, "missing comma before \"...\"")
			}
		}
	case 305:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:   11,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableEmptyDefine {
				lx.report.ErrTok(lhs.Token2, "expected identifier")
			}
		}
	case 306:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:        12,
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
	case 307:
		{
			yyVAL.node = &ControlLine{
				Case:        13,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 310:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 311:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
