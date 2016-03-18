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
	yyDefault           = 57457
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
	ORASSIGN            = 57395
	OROR                = 57396
	PPDEFINE            = 57397
	PPELIF              = 57398
	PPELSE              = 57399
	PPENDIF             = 57400
	PPERROR             = 57401
	PPHASH_NL           = 57402
	PPHEADER_NAME       = 57403
	PPIF                = 57404
	PPIFDEF             = 57405
	PPIFNDEF            = 57406
	PPINCLUDE           = 57407
	PPINCLUDE_NEXT      = 57408
	PPLINE              = 57409
	PPNONDIRECTIVE      = 57410
	PPNUMBER            = 57411
	PPOTHER             = 57412
	PPPASTE             = 57413
	PPPRAGMA            = 57414
	PPUNDEF             = 57415
	PREPROCESSING_FILE  = 1048576
	REGISTER            = 57416
	RESTRICT            = 57417
	RETURN              = 57418
	RSH                 = 57419
	RSHASSIGN           = 57420
	SHORT               = 57421
	SIGNED              = 57422
	SIZEOF              = 57423
	STATIC              = 57424
	STATIC_ASSERT       = 57425
	STRINGLITERAL       = 57426
	STRUCT              = 57427
	SUBASSIGN           = 57428
	SWITCH              = 57429
	TRANSLATION_UNIT    = 1048578
	TYPEDEF             = 57430
	TYPEDEFNAME         = 57431
	TYPEOF              = 57432
	UNARY               = 57433
	UNION               = 57434
	UNSIGNED            = 57435
	VOID                = 57436
	VOLATILE            = 57437
	WHILE               = 57438
	XORASSIGN           = 57439
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -314
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (303x)
		42:      1,   // '*' (264x)
		57377:   2,   // IDENTIFIER (217x)
		38:      3,   // '&' (212x)
		43:      4,   // '+' (212x)
		45:      5,   // '-' (212x)
		57363:   6,   // DEC (212x)
		57381:   7,   // INC (212x)
		41:      8,   // ')' (194x)
		59:      9,   // ';' (194x)
		44:      10,  // ',' (181x)
		57426:   11,  // STRINGLITERAL (160x)
		91:      12,  // '[' (158x)
		33:      13,  // '!' (142x)
		126:     14,  // '~' (142x)
		57347:   15,  // ALIGNOF (142x)
		57358:   16,  // CHARCONST (142x)
		57373:   17,  // FLOATCONST (142x)
		57384:   18,  // INTCONST (142x)
		57387:   19,  // LONGCHARCONST (142x)
		57388:   20,  // LONGSTRINGLITERAL (142x)
		57423:   21,  // SIZEOF (142x)
		57437:   22,  // VOLATILE (132x)
		57360:   23,  // CONST (130x)
		57417:   24,  // RESTRICT (130x)
		57353:   25,  // BOOL (120x)
		57357:   26,  // CHAR (120x)
		57359:   27,  // COMPLEX (120x)
		57367:   28,  // DOUBLE (120x)
		57369:   29,  // ENUM (120x)
		57372:   30,  // FLOAT (120x)
		57383:   31,  // INT (120x)
		57386:   32,  // LONG (120x)
		57421:   33,  // SHORT (120x)
		57422:   34,  // SIGNED (120x)
		57427:   35,  // STRUCT (120x)
		57431:   36,  // TYPEDEFNAME (120x)
		57432:   37,  // TYPEOF (120x)
		57434:   38,  // UNION (120x)
		57435:   39,  // UNSIGNED (120x)
		57436:   40,  // VOID (120x)
		58:      41,  // ':' (119x)
		125:     42,  // '}' (118x)
		57424:   43,  // STATIC (112x)
		57352:   44,  // AUTO (106x)
		57371:   45,  // EXTERN (106x)
		57382:   46,  // INLINE (106x)
		57416:   47,  // REGISTER (106x)
		57430:   48,  // TYPEDEF (106x)
		57344:   49,  // $end (101x)
		61:      50,  // '=' (90x)
		57500:   51,  // Expression (85x)
		93:      52,  // ']' (83x)
		123:     53,  // '{' (83x)
		46:      54,  // '.' (79x)
		57351:   55,  // ASM (76x)
		37:      56,  // '%' (71x)
		47:      57,  // '/' (71x)
		60:      58,  // '<' (71x)
		62:      59,  // '>' (71x)
		63:      60,  // '?' (71x)
		94:      61,  // '^' (71x)
		124:     62,  // '|' (71x)
		57346:   63,  // ADDASSIGN (71x)
		57348:   64,  // ANDAND (71x)
		57349:   65,  // ANDASSIGN (71x)
		57350:   66,  // ARROW (71x)
		57365:   67,  // DIVASSIGN (71x)
		57370:   68,  // EQ (71x)
		57375:   69,  // GEQ (71x)
		57385:   70,  // LEQ (71x)
		57389:   71,  // LSH (71x)
		57390:   72,  // LSHASSIGN (71x)
		57391:   73,  // MODASSIGN (71x)
		57392:   74,  // MULASSIGN (71x)
		57393:   75,  // NEQ (71x)
		57395:   76,  // ORASSIGN (71x)
		57396:   77,  // OROR (71x)
		57419:   78,  // RSH (71x)
		57420:   79,  // RSHASSIGN (71x)
		57428:   80,  // SUBASSIGN (71x)
		57439:   81,  // XORASSIGN (71x)
		10:      82,  // '\n' (58x)
		57412:   83,  // PPOTHER (52x)
		57376:   84,  // GOTO (50x)
		57438:   85,  // WHILE (48x)
		57354:   86,  // BREAK (47x)
		57355:   87,  // CASE (47x)
		57361:   88,  // CONTINUE (47x)
		57364:   89,  // DEFAULT (47x)
		57366:   90,  // DO (47x)
		57374:   91,  // FOR (47x)
		57380:   92,  // IF (47x)
		57418:   93,  // RETURN (47x)
		57429:   94,  // SWITCH (47x)
		57400:   95,  // PPENDIF (44x)
		57399:   96,  // PPELSE (40x)
		57398:   97,  // PPELIF (39x)
		57397:   98,  // PPDEFINE (35x)
		57401:   99,  // PPERROR (35x)
		57402:   100, // PPHASH_NL (35x)
		57404:   101, // PPIF (35x)
		57405:   102, // PPIFDEF (35x)
		57406:   103, // PPIFNDEF (35x)
		57407:   104, // PPINCLUDE (35x)
		57408:   105, // PPINCLUDE_NEXT (35x)
		57409:   106, // PPLINE (35x)
		57410:   107, // PPNONDIRECTIVE (35x)
		57414:   108, // PPPRAGMA (35x)
		57415:   109, // PPUNDEF (35x)
		57368:   110, // ELSE (29x)
		57552:   111, // TypeQualifier (28x)
		57501:   112, // ExpressionList (26x)
		57525:   113, // PPTokenList (22x)
		57527:   114, // PPTokens (22x)
		57496:   115, // EnumSpecifier (20x)
		57547:   116, // StructOrUnion (20x)
		57548:   117, // StructOrUnionSpecifier (20x)
		57555:   118, // TypeSpecifier (20x)
		57502:   119, // ExpressionListOpt (18x)
		57467:   120, // BasicAssemblerStatement (15x)
		57479:   121, // DeclarationSpecifiers (15x)
		57508:   122, // FunctionSpecifier (15x)
		57526:   123, // PPTokenListOpt (15x)
		57425:   124, // STATIC_ASSERT (15x)
		57542:   125, // StorageClassSpecifier (15x)
		57465:   126, // AssemblerStatement (13x)
		57473:   127, // CompoundStatement (13x)
		57504:   128, // ExpressionStatement (12x)
		57522:   129, // IterationStatement (12x)
		57523:   130, // JumpStatement (12x)
		57524:   131, // LabeledStatement (12x)
		57536:   132, // SelectionStatement (12x)
		57540:   133, // Statement (12x)
		57532:   134, // Pointer (11x)
		57533:   135, // PointerOpt (10x)
		57475:   136, // ControlLine (8x)
		57481:   137, // Declarator (8x)
		57511:   138, // GroupPart (8x)
		57515:   139, // IfGroup (8x)
		57516:   140, // IfSection (8x)
		57549:   141, // TextLine (8x)
		57476:   142, // Declaration (7x)
		57451:   143, // $@4 (6x)
		57474:   144, // ConstantExpression (6x)
		57362:   145, // DDD (6x)
		57509:   146, // GroupList (6x)
		57510:   147, // GroupListOpt (5x)
		57535:   148, // ReplacementList (5x)
		57537:   149, // SpecifierQualifierList (5x)
		57553:   150, // TypeQualifierList (5x)
		57441:   151, // $@10 (4x)
		57458:   152, // AbstractDeclarator (4x)
		57463:   153, // AssemblerOperand (4x)
		57466:   154, // AssemblerSymbolicNameOpt (4x)
		57480:   155, // DeclarationSpecifiersOpt (4x)
		57485:   156, // Designator (4x)
		57528:   157, // ParameterDeclaration (4x)
		57551:   158, // TypeName (4x)
		57554:   159, // TypeQualifierListOpt (4x)
		57462:   160, // AssemblerInstructions (3x)
		57464:   161, // AssemblerOperands (3x)
		57472:   162, // CommaOpt (3x)
		57483:   163, // Designation (3x)
		57484:   164, // DesignationOpt (3x)
		57486:   165, // DesignatorList (3x)
		57503:   166, // ExpressionOpt (3x)
		57512:   167, // IdentifierList (3x)
		57517:   168, // InitDeclarator (3x)
		57520:   169, // Initializer (3x)
		57529:   170, // ParameterList (3x)
		57530:   171, // ParameterTypeList (3x)
		57442:   172, // $@11 (2x)
		57452:   173, // $@5 (2x)
		57459:   174, // AbstractDeclaratorOpt (2x)
		57468:   175, // BlockItem (2x)
		57471:   176, // Clobbers (2x)
		57482:   177, // DeclaratorOpt (2x)
		57487:   178, // DirectAbstractDeclarator (2x)
		57488:   179, // DirectAbstractDeclaratorOpt (2x)
		57489:   180, // DirectDeclarator (2x)
		57490:   181, // ElifGroup (2x)
		57497:   182, // EnumerationConstant (2x)
		57498:   183, // Enumerator (2x)
		57505:   184, // ExternalDeclaration (2x)
		57507:   185, // FunctionDefinition (2x)
		57513:   186, // IdentifierListOpt (2x)
		57514:   187, // IdentifierOpt (2x)
		57518:   188, // InitDeclaratorList (2x)
		57519:   189, // InitDeclaratorListOpt (2x)
		57521:   190, // InitializerList (2x)
		57531:   191, // ParameterTypeListOpt (2x)
		57538:   192, // SpecifierQualifierListOpt (2x)
		57541:   193, // StaticAssert (2x)
		57543:   194, // StructDeclaration (2x)
		57545:   195, // StructDeclarator (2x)
		57556:   196, // VolatileOpt (2x)
		57440:   197, // $@1 (1x)
		57443:   198, // $@12 (1x)
		57444:   199, // $@13 (1x)
		57445:   200, // $@14 (1x)
		57446:   201, // $@15 (1x)
		57447:   202, // $@16 (1x)
		57448:   203, // $@17 (1x)
		57449:   204, // $@2 (1x)
		57450:   205, // $@3 (1x)
		57453:   206, // $@6 (1x)
		57454:   207, // $@7 (1x)
		57455:   208, // $@8 (1x)
		57456:   209, // $@9 (1x)
		57460:   210, // ArgumentExpressionList (1x)
		57461:   211, // ArgumentExpressionListOpt (1x)
		57469:   212, // BlockItemList (1x)
		57470:   213, // BlockItemListOpt (1x)
		1048577: 214, // CONSTANT_EXPRESSION (1x)
		57477:   215, // DeclarationList (1x)
		57478:   216, // DeclarationListOpt (1x)
		57491:   217, // ElifGroupList (1x)
		57492:   218, // ElifGroupListOpt (1x)
		57493:   219, // ElseGroup (1x)
		57494:   220, // ElseGroupOpt (1x)
		57495:   221, // EndifLine (1x)
		57499:   222, // EnumeratorList (1x)
		57506:   223, // FunctionBody (1x)
		57378:   224, // IDENTIFIER_LPAREN (1x)
		1048576: 225, // PREPROCESSING_FILE (1x)
		57534:   226, // PreprocessingFile (1x)
		57539:   227, // Start (1x)
		57544:   228, // StructDeclarationList (1x)
		57546:   229, // StructDeclaratorList (1x)
		1048578: 230, // TRANSLATION_UNIT (1x)
		57550:   231, // TranslationUnit (1x)
		57457:   232, // $default (0x)
		57356:   233, // CAST (0x)
		57345:   234, // error (0x)
		57379:   235, // IDENTIFIER_NONREPL (0x)
		57394:   236, // NOELSE (0x)
		57403:   237, // PPHEADER_NAME (0x)
		57411:   238, // PPNUMBER (0x)
		57413:   239, // PPPASTE (0x)
		57433:   240, // UNARY (0x)
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
		57426:   "string literal",
		57347:   "_Alignof",
		57358:   "character constant",
		57373:   "floating-point constant",
		57384:   "integer constant",
		57387:   "long character constant",
		57388:   "long string constant",
		57423:   "sizeof",
		57437:   "volatile",
		57360:   "const",
		57417:   "restrict",
		57353:   "_Bool",
		57357:   "char",
		57359:   "_Complex",
		57367:   "double",
		57369:   "enum",
		57372:   "float",
		57383:   "int",
		57386:   "long",
		57421:   "short",
		57422:   "signed",
		57427:   "struct",
		57431:   "typedefname",
		57432:   "typeof",
		57434:   "union",
		57435:   "unsigned",
		57436:   "void",
		57424:   "static",
		57352:   "auto",
		57371:   "extern",
		57382:   "inline",
		57416:   "register",
		57430:   "typedef",
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
		57395:   "|=",
		57396:   "||",
		57419:   ">>",
		57420:   ">>=",
		57428:   "-=",
		57439:   "^=",
		57412:   "ppother",
		57376:   "goto",
		57438:   "while",
		57354:   "break",
		57355:   "case",
		57361:   "continue",
		57364:   "default",
		57366:   "do",
		57374:   "for",
		57380:   "if",
		57418:   "return",
		57429:   "switch",
		57400:   "#endif",
		57399:   "#else",
		57398:   "#elif",
		57397:   "#define",
		57401:   "#error",
		57402:   "#",
		57404:   "#if",
		57405:   "#ifdef",
		57406:   "#ifndef",
		57407:   "#include",
		57408:   "#include_next",
		57409:   "#line",
		57410:   "#foo",
		57414:   "#pragma",
		57415:   "#undef",
		57368:   "else",
		57425:   "_Static_assert",
		57362:   "...",
		1048577: "constant expression prefix",
		57378:   "identifier immediatelly followed by '('",
		1048576: "preprocessing file prefix",
		1048578: "translation unit prefix",
		57379:   "non replaceable identifier",
		57403:   "header name",
		57411:   "preprocessing number",
		57413:   "##",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {197, 0},
		2:   {227, 3},
		3:   {204, 0},
		4:   {227, 3},
		5:   {205, 0},
		6:   {227, 3},
		7:   {182, 1},
		8:   {210, 1},
		9:   {210, 3},
		10:  {211, 0},
		11:  {211, 1},
		12:  {51, 1},
		13:  {51, 1},
		14:  {51, 1},
		15:  {51, 1},
		16:  {51, 1},
		17:  {51, 1},
		18:  {51, 1},
		19:  {51, 3},
		20:  {51, 4},
		21:  {51, 4},
		22:  {51, 3},
		23:  {51, 3},
		24:  {51, 2},
		25:  {51, 2},
		26:  {51, 7},
		27:  {51, 2},
		28:  {51, 2},
		29:  {51, 2},
		30:  {51, 2},
		31:  {51, 2},
		32:  {51, 2},
		33:  {51, 2},
		34:  {51, 2},
		35:  {51, 2},
		36:  {51, 4},
		37:  {51, 4},
		38:  {51, 3},
		39:  {51, 3},
		40:  {51, 3},
		41:  {51, 3},
		42:  {51, 3},
		43:  {51, 3},
		44:  {51, 3},
		45:  {51, 3},
		46:  {51, 3},
		47:  {51, 3},
		48:  {51, 3},
		49:  {51, 3},
		50:  {51, 3},
		51:  {51, 3},
		52:  {51, 3},
		53:  {51, 3},
		54:  {51, 3},
		55:  {51, 3},
		56:  {51, 5},
		57:  {51, 3},
		58:  {51, 3},
		59:  {51, 3},
		60:  {51, 3},
		61:  {51, 3},
		62:  {51, 3},
		63:  {51, 3},
		64:  {51, 3},
		65:  {51, 3},
		66:  {51, 3},
		67:  {51, 3},
		68:  {51, 4},
		69:  {166, 0},
		70:  {166, 1},
		71:  {112, 1},
		72:  {112, 3},
		73:  {119, 0},
		74:  {119, 1},
		75:  {143, 0},
		76:  {144, 2},
		77:  {142, 3},
		78:  {121, 2},
		79:  {121, 2},
		80:  {121, 2},
		81:  {121, 2},
		82:  {155, 0},
		83:  {155, 1},
		84:  {188, 1},
		85:  {188, 3},
		86:  {189, 0},
		87:  {189, 1},
		88:  {168, 1},
		89:  {173, 0},
		90:  {168, 4},
		91:  {125, 1},
		92:  {125, 1},
		93:  {125, 1},
		94:  {125, 1},
		95:  {125, 1},
		96:  {118, 1},
		97:  {118, 1},
		98:  {118, 1},
		99:  {118, 1},
		100: {118, 1},
		101: {118, 1},
		102: {118, 1},
		103: {118, 1},
		104: {118, 1},
		105: {118, 1},
		106: {118, 1},
		107: {118, 1},
		108: {118, 1},
		109: {118, 1},
		110: {118, 4},
		111: {118, 4},
		112: {206, 0},
		113: {207, 0},
		114: {117, 7},
		115: {117, 2},
		116: {116, 1},
		117: {116, 1},
		118: {228, 1},
		119: {228, 2},
		120: {194, 3},
		121: {194, 2},
		122: {149, 2},
		123: {149, 2},
		124: {192, 0},
		125: {192, 1},
		126: {229, 1},
		127: {229, 3},
		128: {195, 1},
		129: {195, 3},
		130: {162, 0},
		131: {162, 1},
		132: {208, 0},
		133: {115, 7},
		134: {115, 2},
		135: {222, 1},
		136: {222, 3},
		137: {183, 1},
		138: {183, 3},
		139: {111, 1},
		140: {111, 1},
		141: {111, 1},
		142: {122, 1},
		143: {137, 2},
		144: {177, 0},
		145: {177, 1},
		146: {180, 1},
		147: {180, 3},
		148: {180, 5},
		149: {180, 6},
		150: {180, 6},
		151: {180, 5},
		152: {209, 0},
		153: {180, 5},
		154: {180, 4},
		155: {134, 2},
		156: {134, 3},
		157: {135, 0},
		158: {135, 1},
		159: {150, 1},
		160: {150, 2},
		161: {159, 0},
		162: {159, 1},
		163: {171, 1},
		164: {171, 3},
		165: {191, 0},
		166: {191, 1},
		167: {170, 1},
		168: {170, 3},
		169: {157, 2},
		170: {157, 2},
		171: {167, 1},
		172: {167, 3},
		173: {186, 0},
		174: {186, 1},
		175: {187, 0},
		176: {187, 1},
		177: {151, 0},
		178: {158, 3},
		179: {152, 1},
		180: {152, 2},
		181: {174, 0},
		182: {174, 1},
		183: {178, 3},
		184: {178, 4},
		185: {178, 5},
		186: {178, 6},
		187: {178, 6},
		188: {178, 4},
		189: {172, 0},
		190: {178, 4},
		191: {198, 0},
		192: {178, 5},
		193: {179, 0},
		194: {179, 1},
		195: {169, 1},
		196: {169, 4},
		197: {190, 2},
		198: {190, 4},
		199: {163, 2},
		200: {164, 0},
		201: {164, 1},
		202: {165, 1},
		203: {165, 2},
		204: {156, 3},
		205: {156, 2},
		206: {133, 1},
		207: {133, 1},
		208: {133, 1},
		209: {133, 1},
		210: {133, 1},
		211: {133, 1},
		212: {133, 1},
		213: {131, 3},
		214: {131, 4},
		215: {131, 3},
		216: {199, 0},
		217: {127, 4},
		218: {212, 1},
		219: {212, 2},
		220: {213, 0},
		221: {213, 1},
		222: {175, 1},
		223: {175, 1},
		224: {128, 2},
		225: {132, 5},
		226: {132, 7},
		227: {132, 5},
		228: {129, 5},
		229: {129, 7},
		230: {129, 9},
		231: {129, 8},
		232: {130, 3},
		233: {130, 2},
		234: {130, 2},
		235: {130, 3},
		236: {231, 1},
		237: {231, 2},
		238: {184, 1},
		239: {184, 1},
		240: {184, 2},
		241: {184, 1},
		242: {200, 0},
		243: {185, 5},
		244: {201, 0},
		245: {223, 2},
		246: {202, 0},
		247: {223, 3},
		248: {215, 1},
		249: {215, 2},
		250: {216, 0},
		251: {203, 0},
		252: {216, 2},
		253: {160, 1},
		254: {160, 2},
		255: {120, 5},
		256: {196, 0},
		257: {196, 1},
		258: {153, 5},
		259: {161, 1},
		260: {161, 3},
		261: {154, 0},
		262: {154, 3},
		263: {176, 1},
		264: {176, 3},
		265: {126, 1},
		266: {126, 7},
		267: {126, 9},
		268: {126, 11},
		269: {126, 13},
		270: {193, 7},
		271: {226, 1},
		272: {146, 1},
		273: {146, 2},
		274: {147, 0},
		275: {147, 1},
		276: {138, 1},
		277: {138, 1},
		278: {138, 3},
		279: {138, 1},
		280: {140, 4},
		281: {139, 4},
		282: {139, 4},
		283: {139, 4},
		284: {217, 1},
		285: {217, 2},
		286: {218, 0},
		287: {218, 1},
		288: {181, 4},
		289: {219, 3},
		290: {220, 0},
		291: {220, 1},
		292: {221, 1},
		293: {136, 3},
		294: {136, 5},
		295: {136, 7},
		296: {136, 5},
		297: {136, 2},
		298: {136, 1},
		299: {136, 3},
		300: {136, 3},
		301: {136, 2},
		302: {136, 3},
		303: {136, 6},
		304: {136, 2},
		305: {136, 4},
		306: {136, 3},
		307: {141, 1},
		308: {148, 1},
		309: {113, 1},
		310: {123, 1},
		311: {123, 2},
		312: {114, 1},
		313: {114, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 49}:   "invalid empty input",
		yyXError{560, -1}: "expected #endif",
		yyXError{562, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{480, -1}: "expected $end",
		yyXError{482, -1}: "expected $end",
		yyXError{31, -1}:  "expected '('",
		yyXError{46, -1}:  "expected '('",
		yyXError{72, -1}:  "expected '('",
		yyXError{282, -1}: "expected '('",
		yyXError{358, -1}: "expected '('",
		yyXError{379, -1}: "expected '('",
		yyXError{414, -1}: "expected '('",
		yyXError{415, -1}: "expected '('",
		yyXError{416, -1}: "expected '('",
		yyXError{418, -1}: "expected '('",
		yyXError{444, -1}: "expected '('",
		yyXError{51, -1}:  "expected ')'",
		yyXError{74, -1}:  "expected ')'",
		yyXError{81, -1}:  "expected ')'",
		yyXError{173, -1}: "expected ')'",
		yyXError{188, -1}: "expected ')'",
		yyXError{191, -1}: "expected ')'",
		yyXError{194, -1}: "expected ')'",
		yyXError{202, -1}: "expected ')'",
		yyXError{207, -1}: "expected ')'",
		yyXError{213, -1}: "expected ')'",
		yyXError{229, -1}: "expected ')'",
		yyXError{234, -1}: "expected ')'",
		yyXError{245, -1}: "expected ')'",
		yyXError{280, -1}: "expected ')'",
		yyXError{329, -1}: "expected ')'",
		yyXError{434, -1}: "expected ')'",
		yyXError{440, -1}: "expected ')'",
		yyXError{524, -1}: "expected ')'",
		yyXError{525, -1}: "expected ')'",
		yyXError{532, -1}: "expected ')'",
		yyXError{535, -1}: "expected ')'",
		yyXError{49, -1}:  "expected ','",
		yyXError{316, -1}: "expected ':'",
		yyXError{361, -1}: "expected ':'",
		yyXError{407, -1}: "expected ':'",
		yyXError{468, -1}: "expected ':'",
		yyXError{43, -1}:  "expected ';'",
		yyXError{52, -1}:  "expected ';'",
		yyXError{337, -1}: "expected ';'",
		yyXError{353, -1}: "expected ';'",
		yyXError{413, -1}: "expected ';'",
		yyXError{420, -1}: "expected ';'",
		yyXError{421, -1}: "expected ';'",
		yyXError{423, -1}: "expected ';'",
		yyXError{427, -1}: "expected ';'",
		yyXError{430, -1}: "expected ';'",
		yyXError{432, -1}: "expected ';'",
		yyXError{438, -1}: "expected ';'",
		yyXError{447, -1}: "expected ';'",
		yyXError{341, -1}: "expected '='",
		yyXError{86, -1}:  "expected '['",
		yyXError{504, -1}: "expected '\\n'",
		yyXError{508, -1}: "expected '\\n'",
		yyXError{512, -1}: "expected '\\n'",
		yyXError{515, -1}: "expected '\\n'",
		yyXError{517, -1}: "expected '\\n'",
		yyXError{539, -1}: "expected '\\n'",
		yyXError{544, -1}: "expected '\\n'",
		yyXError{547, -1}: "expected '\\n'",
		yyXError{554, -1}: "expected '\\n'",
		yyXError{559, -1}: "expected '\\n'",
		yyXError{565, -1}: "expected '\\n'",
		yyXError{92, -1}:  "expected ']'",
		yyXError{181, -1}: "expected ']'",
		yyXError{225, -1}: "expected ']'",
		yyXError{257, -1}: "expected ']'",
		yyXError{367, -1}: "expected ']'",
		yyXError{290, -1}: "expected '{'",
		yyXError{292, -1}: "expected '{'",
		yyXError{304, -1}: "expected '{'",
		yyXError{306, -1}: "expected '{'",
		yyXError{266, -1}: "expected '}'",
		yyXError{270, -1}: "expected '}'",
		yyXError{301, -1}: "expected '}'",
		yyXError{408, -1}: "expected '}'",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{201, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{85, -1}:  "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{284, -1}: "expected assembler instructions or string literal",
		yyXError{357, -1}: "expected assembler instructions or string literal",
		yyXError{359, -1}: "expected assembler instructions or string literal",
		yyXError{369, -1}: "expected assembler operand or one of ['[', string literal]",
		yyXError{362, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{384, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{387, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{352, -1}: "expected assembler statement or asm",
		yyXError{410, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{370, -1}: "expected clobbers or string literal",
		yyXError{390, -1}: "expected clobbers or string literal",
		yyXError{351, -1}: "expected compound statement or '{'",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{47, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{254, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{298, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{320, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{406, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{479, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{343, -1}: "expected declaration list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{346, -1}: "expected declaration or one of ['{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{429, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{319, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{193, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{251, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{196, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{82, -1}:  "expected direct abstract declarator or one of ['(', '[']",
		yyXError{317, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{552, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{558, -1}: "expected endif line or #endif",
		yyXError{489, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{550, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{293, -1}: "expected enumerator list or identifier",
		yyXError{300, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{97, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{121, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{445, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{449, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{453, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{457, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{61, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{242, -1}: "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{246, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		yyXError{89, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{224, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{281, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{48, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{63, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{64, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{65, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{66, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{67, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{68, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{69, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{70, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{71, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{95, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{103, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{122, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{147, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{148, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{175, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{182, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{218, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{221, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{380, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{93, -1}:  "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{216, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{327, -1}: "expected expression or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{6, -1}:   "expected external declaration or one of [$end, _Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{342, -1}: "expected function body or one of ['{', asm]",
		yyXError{349, -1}: "expected function body or one of ['{', asm]",
		yyXError{340, -1}: "expected function body or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{541, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{483, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{99, -1}:  "expected identifier",
		yyXError{100, -1}: "expected identifier",
		yyXError{210, -1}: "expected identifier",
		yyXError{255, -1}: "expected identifier",
		yyXError{366, -1}: "expected identifier",
		yyXError{419, -1}: "expected identifier",
		yyXError{491, -1}: "expected identifier",
		yyXError{492, -1}: "expected identifier",
		yyXError{499, -1}: "expected identifier",
		yyXError{374, -1}: "expected identifier list or identifier",
		yyXError{521, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{475, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{248, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{262, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{250, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{268, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{473, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{265, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{54, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{55, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{56, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{57, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{58, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{59, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{60, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{101, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{102, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
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
		yyXError{150, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
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
		yyXError{174, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{178, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{186, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{241, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{243, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{247, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{271, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{272, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{273, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{274, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{275, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{276, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{277, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{278, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{279, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{62, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{145, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{149, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{171, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{176, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{328, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{381, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{397, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{261, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{88, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{96, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{183, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{219, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{222, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{484, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{485, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{486, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{488, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{495, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{501, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{503, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{506, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{509, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{511, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{513, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{514, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{516, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{518, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{519, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{522, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{527, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{528, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{530, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{534, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{537, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{538, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{543, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{563, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{564, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{566, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{542, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{546, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{549, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{551, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{556, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{557, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{465, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{477, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{40, -1}:  "expected one of [$end, _Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{41, -1}:  "expected one of [$end, _Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{42, -1}:  "expected one of [$end, _Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{44, -1}:  "expected one of [$end, _Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{53, -1}:  "expected one of [$end, _Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{289, -1}: "expected one of [$end, _Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{350, -1}: "expected one of [$end, _Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{393, -1}: "expected one of [$end, _Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{395, -1}: "expected one of [$end, _Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{478, -1}: "expected one of [$end, _Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{36, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{38, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{90, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{179, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{288, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{355, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{376, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{386, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{389, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{392, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{399, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{400, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{401, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{402, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{403, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{404, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{405, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{424, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{425, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{426, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{428, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{436, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{442, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{448, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{452, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{456, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{460, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{462, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{463, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{467, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{470, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{472, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{409, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{411, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{412, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{464, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{252, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{259, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{291, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{305, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
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
		yyXError{302, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{325, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{330, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{331, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{12, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{39, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{332, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{333, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{334, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{335, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{336, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{238, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{239, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{240, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{199, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{200, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{203, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{212, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{214, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{220, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{223, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{226, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{227, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{80, -1}:  "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{237, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{84, -1}:  "expected one of ['(', ')', ',', '[']",
		yyXError{133, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{180, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{184, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{185, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{187, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{195, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{231, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{235, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{283, -1}: "expected one of ['(', goto]",
		yyXError{356, -1}: "expected one of ['(', goto]",
		yyXError{318, -1}: "expected one of ['(', identifier]",
		yyXError{364, -1}: "expected one of [')', ',', ':']",
		yyXError{371, -1}: "expected one of [')', ',', ':']",
		yyXError{377, -1}: "expected one of [')', ',', ':']",
		yyXError{378, -1}: "expected one of [')', ',', ':']",
		yyXError{382, -1}: "expected one of [')', ',', ':']",
		yyXError{385, -1}: "expected one of [')', ',', ':']",
		yyXError{388, -1}: "expected one of [')', ',', ':']",
		yyXError{398, -1}: "expected one of [')', ',', ';']",
		yyXError{208, -1}: "expected one of [')', ',', ...]",
		yyXError{211, -1}: "expected one of [')', ',', ...]",
		yyXError{523, -1}: "expected one of [')', ',', ...]",
		yyXError{83, -1}:  "expected one of [')', ',']",
		yyXError{172, -1}: "expected one of [')', ',']",
		yyXError{190, -1}: "expected one of [')', ',']",
		yyXError{192, -1}: "expected one of [')', ',']",
		yyXError{197, -1}: "expected one of [')', ',']",
		yyXError{198, -1}: "expected one of [')', ',']",
		yyXError{209, -1}: "expected one of [')', ',']",
		yyXError{230, -1}: "expected one of [')', ',']",
		yyXError{244, -1}: "expected one of [')', ',']",
		yyXError{375, -1}: "expected one of [')', ',']",
		yyXError{391, -1}: "expected one of [')', ',']",
		yyXError{446, -1}: "expected one of [')', ',']",
		yyXError{450, -1}: "expected one of [')', ',']",
		yyXError{454, -1}: "expected one of [')', ',']",
		yyXError{458, -1}: "expected one of [')', ',']",
		yyXError{285, -1}: "expected one of [')', ':', string literal]",
		yyXError{287, -1}: "expected one of [')', ':', string literal]",
		yyXError{383, -1}: "expected one of [')', ':', string literal]",
		yyXError{286, -1}: "expected one of [')', string literal]",
		yyXError{315, -1}: "expected one of [',', ':', ';']",
		yyXError{146, -1}: "expected one of [',', ':']",
		yyXError{365, -1}: "expected one of [',', ':']",
		yyXError{372, -1}: "expected one of [',', ':']",
		yyXError{348, -1}: "expected one of [',', ';', '=']",
		yyXError{267, -1}: "expected one of [',', ';', '}']",
		yyXError{312, -1}: "expected one of [',', ';']",
		yyXError{314, -1}: "expected one of [',', ';']",
		yyXError{321, -1}: "expected one of [',', ';']",
		yyXError{324, -1}: "expected one of [',', ';']",
		yyXError{338, -1}: "expected one of [',', ';']",
		yyXError{339, -1}: "expected one of [',', ';']",
		yyXError{474, -1}: "expected one of [',', ';']",
		yyXError{476, -1}: "expected one of [',', ';']",
		yyXError{294, -1}: "expected one of [',', '=', '}']",
		yyXError{297, -1}: "expected one of [',', '=', '}']",
		yyXError{177, -1}: "expected one of [',', ']']",
		yyXError{263, -1}: "expected one of [',', '}']",
		yyXError{269, -1}: "expected one of [',', '}']",
		yyXError{296, -1}: "expected one of [',', '}']",
		yyXError{299, -1}: "expected one of [',', '}']",
		yyXError{303, -1}: "expected one of [',', '}']",
		yyXError{253, -1}: "expected one of ['.', '=', '[']",
		yyXError{256, -1}: "expected one of ['.', '=', '[']",
		yyXError{258, -1}: "expected one of ['.', '=', '[']",
		yyXError{260, -1}: "expected one of ['.', '=', '[']",
		yyXError{360, -1}: "expected one of [':', string literal]",
		yyXError{493, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		yyXError{502, -1}: "expected one of ['\\n', ppother]",
		yyXError{505, -1}: "expected one of ['\\n', ppother]",
		yyXError{507, -1}: "expected one of ['\\n', ppother]",
		yyXError{345, -1}: "expected one of ['{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{347, -1}: "expected one of ['{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{33, -1}:  "expected one of ['{', identifier]",
		yyXError{34, -1}:  "expected one of ['{', identifier]",
		yyXError{310, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{313, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{322, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{326, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{531, -1}: "expected one of [..., identifier]",
		yyXError{78, -1}:  "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{98, -1}:  "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{394, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{396, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{249, -1}: "expected optional comma or one of [',', '}']",
		yyXError{264, -1}: "expected optional comma or one of [',', '}']",
		yyXError{295, -1}: "expected optional comma or one of [',', '}']",
		yyXError{8, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{433, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{439, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{422, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{431, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{437, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{215, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{204, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{87, -1}:  "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{91, -1}:  "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{540, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{545, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{548, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{555, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{561, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{205, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{32, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{35, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{344, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{189, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{232, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{233, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{76, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{77, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{494, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{498, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{79, -1}:  "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{354, -1}: "expected optional volatile or one of ['(', goto, volatile]",
		yyXError{45, -1}:  "expected optional volatile or one of ['(', volatile]",
		yyXError{228, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{206, -1}: "expected parameter type list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{236, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{481, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{520, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{526, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{529, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{533, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{536, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{75, -1}:  "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{417, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{435, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{441, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{451, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{455, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{459, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{461, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{466, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{469, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{471, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{50, -1}:  "expected string literal",
		yyXError{363, -1}: "expected string literal",
		yyXError{368, -1}: "expected string literal",
		yyXError{373, -1}: "expected string literal",
		yyXError{307, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{308, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{309, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{311, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		yyXError{323, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{510, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{487, -1}: "expected token list or ppother",
		yyXError{490, -1}: "expected token list or ppother",
		yyXError{496, -1}: "expected token list or ppother",
		yyXError{497, -1}: "expected token list or ppother",
		yyXError{500, -1}: "expected token list or ppother",
		yyXError{553, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [_Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{5, -1}:   "expected translation unit or one of [_Bool, _Complex, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{73, -1}:  "expected type name or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{94, -1}:  "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{217, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{443, -1}: "expected while",
		yyXError{3, 49}:   "unexpected EOF",
		yyXError{2, 49}:   "unexpected EOF",
		yyXError{4, 49}:   "unexpected EOF",
	}

	yyParseTab = [567][]uint16{
		// 0
		{214: 317, 225: 316, 227: 315, 230: 318},
		{49: 314},
		{82: 313, 313, 98: 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 197: 795},
		{311, 311, 311, 311, 311, 311, 311, 311, 11: 311, 13: 311, 311, 311, 311, 311, 311, 311, 311, 311, 204: 793},
		{22: 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 43: 309, 309, 309, 309, 309, 309, 55: 309, 124: 309, 205: 319},
		// 5
		{22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 55: 359, 111: 324, 115: 343, 346, 342, 323, 120: 357, 321, 325, 124: 360, 322, 142: 356, 184: 354, 355, 193: 358, 231: 320},
		{22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 308, 55: 359, 111: 324, 115: 343, 346, 342, 323, 120: 357, 321, 325, 124: 360, 322, 142: 356, 184: 792, 355, 193: 358},
		{157, 393, 157, 9: 228, 134: 632, 631, 137: 654, 168: 652, 188: 653, 651},
		{232, 232, 232, 8: 232, 232, 232, 12: 232, 22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 111: 324, 115: 343, 346, 342, 323, 121: 647, 325, 125: 322, 155: 650},
		{232, 232, 232, 8: 232, 232, 232, 12: 232, 22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 111: 324, 115: 343, 346, 342, 323, 121: 647, 325, 125: 322, 155: 649},
		// 10
		{232, 232, 232, 8: 232, 232, 232, 12: 232, 22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 111: 324, 115: 343, 346, 342, 323, 121: 647, 325, 125: 322, 155: 648},
		{232, 232, 232, 8: 232, 232, 232, 12: 232, 22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 111: 324, 115: 343, 346, 342, 323, 121: 647, 325, 125: 322, 155: 646},
		{223, 223, 223, 8: 223, 223, 223, 12: 223, 22: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 43: 223, 223, 223, 223, 223, 223},
		{222, 222, 222, 8: 222, 222, 222, 12: 222, 22: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 43: 222, 222, 222, 222, 222, 222},
		{221, 221, 221, 8: 221, 221, 221, 12: 221, 22: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 43: 221, 221, 221, 221, 221, 221},
		// 15
		{220, 220, 220, 8: 220, 220, 220, 12: 220, 22: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 43: 220, 220, 220, 220, 220, 220},
		{219, 219, 219, 8: 219, 219, 219, 12: 219, 22: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 43: 219, 219, 219, 219, 219, 219},
		{218, 218, 218, 8: 218, 218, 218, 12: 218, 22: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 43: 218, 218, 218, 218, 218, 218},
		{217, 217, 217, 8: 217, 217, 217, 12: 217, 22: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 43: 217, 217, 217, 217, 217, 217},
		{216, 216, 216, 8: 216, 216, 216, 12: 216, 22: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 43: 216, 216, 216, 216, 216, 216},
		// 20
		{215, 215, 215, 8: 215, 215, 215, 12: 215, 22: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 43: 215, 215, 215, 215, 215, 215},
		{214, 214, 214, 8: 214, 214, 214, 12: 214, 22: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 43: 214, 214, 214, 214, 214, 214},
		{213, 213, 213, 8: 213, 213, 213, 12: 213, 22: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 43: 213, 213, 213, 213, 213, 213},
		{212, 212, 212, 8: 212, 212, 212, 12: 212, 22: 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 43: 212, 212, 212, 212, 212, 212},
		{211, 211, 211, 8: 211, 211, 211, 12: 211, 22: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 43: 211, 211, 211, 211, 211, 211},
		// 25
		{210, 210, 210, 8: 210, 210, 210, 12: 210, 22: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 43: 210, 210, 210, 210, 210, 210},
		{209, 209, 209, 8: 209, 209, 209, 12: 209, 22: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 43: 209, 209, 209, 209, 209, 209},
		{208, 208, 208, 8: 208, 208, 208, 12: 208, 22: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 43: 208, 208, 208, 208, 208, 208},
		{207, 207, 207, 8: 207, 207, 207, 12: 207, 22: 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 43: 207, 207, 207, 207, 207, 207},
		{206, 206, 206, 8: 206, 206, 206, 12: 206, 22: 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 43: 206, 206, 206, 206, 206, 206},
		// 30
		{205, 205, 205, 8: 205, 205, 205, 12: 205, 22: 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 43: 205, 205, 205, 205, 205, 205},
		{641},
		{2: 619, 53: 139, 187: 618},
		{2: 198, 53: 198},
		{2: 197, 53: 197},
		// 35
		{2: 605, 53: 139, 187: 604},
		{175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 43: 175, 175, 175, 175, 175, 175, 52: 175},
		{174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 43: 174, 174, 174, 174, 174, 174, 52: 174},
		{173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 43: 173, 173, 173, 173, 173, 173, 52: 173},
		{172, 172, 172, 8: 172, 172, 172, 12: 172, 22: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 43: 172, 172, 172, 172, 172, 172},
		// 40
		{22: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 43: 78, 78, 78, 78, 78, 78, 78, 55: 78, 124: 78},
		{22: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 43: 76, 76, 76, 76, 76, 76, 76, 55: 76, 124: 76},
		{22: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 43: 75, 75, 75, 75, 75, 75, 75, 55: 75, 124: 75},
		{9: 603},
		{22: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 43: 73, 73, 73, 73, 73, 73, 73, 55: 73, 124: 73},
		// 45
		{58, 22: 597, 196: 596},
		{361},
		{239, 239, 239, 239, 239, 239, 239, 239, 11: 239, 13: 239, 239, 239, 239, 239, 239, 239, 239, 239, 143: 362, 363},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 376},
		{10: 364},
		// 50
		{11: 365},
		{8: 366},
		{9: 367},
		{22: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 43: 44, 44, 44, 44, 44, 44, 44, 55: 44, 124: 44},
		{302, 302, 3: 302, 302, 302, 302, 302, 302, 302, 302, 12: 302, 41: 302, 302, 49: 302, 302, 52: 302, 54: 302, 56: 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302},
		// 55
		{301, 301, 3: 301, 301, 301, 301, 301, 301, 301, 301, 12: 301, 41: 301, 301, 49: 301, 301, 52: 301, 54: 301, 56: 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301},
		{300, 300, 3: 300, 300, 300, 300, 300, 300, 300, 300, 12: 300, 41: 300, 300, 49: 300, 300, 52: 300, 54: 300, 56: 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300},
		{299, 299, 3: 299, 299, 299, 299, 299, 299, 299, 299, 12: 299, 41: 299, 299, 49: 299, 299, 52: 299, 54: 299, 56: 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299},
		{298, 298, 3: 298, 298, 298, 298, 298, 298, 298, 298, 12: 298, 41: 298, 298, 49: 298, 298, 52: 298, 54: 298, 56: 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298},
		{297, 297, 3: 297, 297, 297, 297, 297, 297, 297, 297, 12: 297, 41: 297, 297, 49: 297, 297, 52: 297, 54: 297, 56: 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297},
		// 60
		{296, 296, 3: 296, 296, 296, 296, 296, 296, 296, 296, 12: 296, 41: 296, 296, 49: 296, 296, 52: 296, 54: 296, 56: 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 51: 459, 112: 558, 151: 389, 158: 594},
		{412, 417, 3: 430, 420, 421, 416, 415, 9: 238, 238, 12: 411, 41: 238, 238, 49: 238, 436, 52: 238, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 593},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 592},
		// 65
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 591},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 500},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 590},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 589},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 588},
		// 70
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 587},
		{556, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 557},
		{387},
		{22: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 151: 389, 158: 388},
		{8: 555},
		// 75
		{22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 111: 391, 115: 343, 346, 342, 390, 149: 392},
		{190, 190, 190, 8: 190, 190, 12: 190, 22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 190, 111: 391, 115: 343, 346, 342, 390, 149: 553, 192: 554},
		{190, 190, 190, 8: 190, 190, 12: 190, 22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 190, 111: 391, 115: 343, 346, 342, 390, 149: 553, 192: 552},
		{157, 393, 8: 133, 12: 157, 134: 394, 396, 152: 397, 174: 395},
		{153, 153, 153, 8: 153, 10: 153, 12: 153, 22: 352, 350, 351, 111: 404, 150: 408, 159: 550},
		// 80
		{156, 2: 156, 8: 135, 10: 135, 12: 156},
		{8: 136},
		{399, 12: 121, 178: 398, 400},
		{8: 132, 10: 132},
		{546, 8: 134, 10: 134, 12: 120},
		// 85
		{157, 393, 8: 125, 12: 157, 22: 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 43: 125, 125, 125, 125, 125, 125, 134: 394, 396, 152: 502, 172: 503},
		{12: 401},
		{375, 403, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 352, 350, 351, 43: 407, 51: 402, 245, 111: 404, 150: 405, 166: 406},
		{412, 417, 3: 430, 420, 421, 416, 415, 12: 411, 50: 436, 52: 244, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 500, 501},
		// 90
		{155, 155, 155, 155, 155, 155, 155, 155, 155, 10: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 43: 155, 52: 155},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 352, 350, 351, 43: 496, 51: 402, 245, 111: 493, 166: 495},
		{52: 494},
		{153, 153, 153, 153, 153, 153, 153, 153, 11: 153, 13: 153, 153, 153, 153, 153, 153, 153, 153, 153, 352, 350, 351, 111: 404, 150: 408, 159: 409},
		{152, 152, 152, 152, 152, 152, 152, 152, 152, 10: 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 352, 350, 351, 111: 493},
		// 95
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 410},
		{412, 417, 3: 430, 420, 421, 416, 415, 12: 411, 50: 436, 52: 447, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 112: 491},
		{375, 380, 368, 379, 381, 382, 378, 377, 304, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 485, 210: 486, 487},
		{2: 484},
		// 100
		{2: 483},
		{290, 290, 3: 290, 290, 290, 290, 290, 290, 290, 290, 12: 290, 41: 290, 290, 49: 290, 290, 52: 290, 54: 290, 56: 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290},
		{289, 289, 3: 289, 289, 289, 289, 289, 289, 289, 289, 12: 289, 41: 289, 289, 49: 289, 289, 52: 289, 54: 289, 56: 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 482},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 481},
		// 105
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 480},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 479},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 478},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 477},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 476},
		// 110
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 475},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 474},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 473},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 472},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 471},
		// 115
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 470},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 469},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 468},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 467},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 466},
		// 120
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 465},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 112: 460},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 458},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 457},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 456},
		// 125
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 455},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 454},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 453},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 452},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 451},
		// 130
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 450},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 449},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 448},
		{128, 8: 128, 10: 128, 12: 128},
		{412, 417, 3: 430, 420, 421, 416, 415, 247, 247, 247, 12: 411, 41: 247, 247, 49: 247, 436, 52: 247, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		// 135
		{412, 417, 3: 430, 420, 421, 416, 415, 248, 248, 248, 12: 411, 41: 248, 248, 49: 248, 436, 52: 248, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{412, 417, 3: 430, 420, 421, 416, 415, 249, 249, 249, 12: 411, 41: 249, 249, 49: 249, 436, 52: 249, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{412, 417, 3: 430, 420, 421, 416, 415, 250, 250, 250, 12: 411, 41: 250, 250, 49: 250, 436, 52: 250, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{412, 417, 3: 430, 420, 421, 416, 415, 251, 251, 251, 12: 411, 41: 251, 251, 49: 251, 436, 52: 251, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{412, 417, 3: 430, 420, 421, 416, 415, 252, 252, 252, 12: 411, 41: 252, 252, 49: 252, 436, 52: 252, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		// 140
		{412, 417, 3: 430, 420, 421, 416, 415, 253, 253, 253, 12: 411, 41: 253, 253, 49: 253, 436, 52: 253, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{412, 417, 3: 430, 420, 421, 416, 415, 254, 254, 254, 12: 411, 41: 254, 254, 49: 254, 436, 52: 254, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{412, 417, 3: 430, 420, 421, 416, 415, 255, 255, 255, 12: 411, 41: 255, 255, 49: 255, 436, 52: 255, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{412, 417, 3: 430, 420, 421, 416, 415, 256, 256, 256, 12: 411, 41: 256, 256, 49: 256, 436, 52: 256, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{412, 417, 3: 430, 420, 421, 416, 415, 257, 257, 257, 12: 411, 41: 257, 257, 49: 257, 436, 52: 257, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		// 145
		{412, 417, 3: 430, 420, 421, 416, 415, 243, 243, 243, 12: 411, 41: 243, 50: 436, 52: 243, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{10: 462, 41: 461},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 464},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 463},
		{412, 417, 3: 430, 420, 421, 416, 415, 242, 242, 242, 12: 411, 41: 242, 50: 436, 52: 242, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		// 150
		{412, 417, 3: 430, 420, 421, 416, 415, 258, 258, 258, 12: 411, 41: 258, 258, 49: 258, 258, 52: 258, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 258, 433, 258, 414, 258, 428, 427, 426, 422, 258, 258, 258, 429, 258, 434, 423, 258, 258, 258},
		{412, 417, 3: 430, 420, 421, 416, 415, 259, 259, 259, 12: 411, 41: 259, 259, 49: 259, 259, 52: 259, 54: 413, 56: 419, 418, 424, 425, 259, 431, 432, 259, 433, 259, 414, 259, 428, 427, 426, 422, 259, 259, 259, 429, 259, 259, 423, 259, 259, 259},
		{412, 417, 3: 430, 420, 421, 416, 415, 260, 260, 260, 12: 411, 41: 260, 260, 49: 260, 260, 52: 260, 54: 413, 56: 419, 418, 424, 425, 260, 431, 432, 260, 260, 260, 414, 260, 428, 427, 426, 422, 260, 260, 260, 429, 260, 260, 423, 260, 260, 260},
		{412, 417, 3: 430, 420, 421, 416, 415, 261, 261, 261, 12: 411, 41: 261, 261, 49: 261, 261, 52: 261, 54: 413, 56: 419, 418, 424, 425, 261, 431, 261, 261, 261, 261, 414, 261, 428, 427, 426, 422, 261, 261, 261, 429, 261, 261, 423, 261, 261, 261},
		{412, 417, 3: 430, 420, 421, 416, 415, 262, 262, 262, 12: 411, 41: 262, 262, 49: 262, 262, 52: 262, 54: 413, 56: 419, 418, 424, 425, 262, 262, 262, 262, 262, 262, 414, 262, 428, 427, 426, 422, 262, 262, 262, 429, 262, 262, 423, 262, 262, 262},
		// 155
		{412, 417, 3: 263, 420, 421, 416, 415, 263, 263, 263, 12: 411, 41: 263, 263, 49: 263, 263, 52: 263, 54: 413, 56: 419, 418, 424, 425, 263, 263, 263, 263, 263, 263, 414, 263, 428, 427, 426, 422, 263, 263, 263, 429, 263, 263, 423, 263, 263, 263},
		{412, 417, 3: 264, 420, 421, 416, 415, 264, 264, 264, 12: 411, 41: 264, 264, 49: 264, 264, 52: 264, 54: 413, 56: 419, 418, 424, 425, 264, 264, 264, 264, 264, 264, 414, 264, 264, 427, 426, 422, 264, 264, 264, 264, 264, 264, 423, 264, 264, 264},
		{412, 417, 3: 265, 420, 421, 416, 415, 265, 265, 265, 12: 411, 41: 265, 265, 49: 265, 265, 52: 265, 54: 413, 56: 419, 418, 424, 425, 265, 265, 265, 265, 265, 265, 414, 265, 265, 427, 426, 422, 265, 265, 265, 265, 265, 265, 423, 265, 265, 265},
		{412, 417, 3: 266, 420, 421, 416, 415, 266, 266, 266, 12: 411, 41: 266, 266, 49: 266, 266, 52: 266, 54: 413, 56: 419, 418, 266, 266, 266, 266, 266, 266, 266, 266, 414, 266, 266, 266, 266, 422, 266, 266, 266, 266, 266, 266, 423, 266, 266, 266},
		{412, 417, 3: 267, 420, 421, 416, 415, 267, 267, 267, 12: 411, 41: 267, 267, 49: 267, 267, 52: 267, 54: 413, 56: 419, 418, 267, 267, 267, 267, 267, 267, 267, 267, 414, 267, 267, 267, 267, 422, 267, 267, 267, 267, 267, 267, 423, 267, 267, 267},
		// 160
		{412, 417, 3: 268, 420, 421, 416, 415, 268, 268, 268, 12: 411, 41: 268, 268, 49: 268, 268, 52: 268, 54: 413, 56: 419, 418, 268, 268, 268, 268, 268, 268, 268, 268, 414, 268, 268, 268, 268, 422, 268, 268, 268, 268, 268, 268, 423, 268, 268, 268},
		{412, 417, 3: 269, 420, 421, 416, 415, 269, 269, 269, 12: 411, 41: 269, 269, 49: 269, 269, 52: 269, 54: 413, 56: 419, 418, 269, 269, 269, 269, 269, 269, 269, 269, 414, 269, 269, 269, 269, 422, 269, 269, 269, 269, 269, 269, 423, 269, 269, 269},
		{412, 417, 3: 270, 420, 421, 416, 415, 270, 270, 270, 12: 411, 41: 270, 270, 49: 270, 270, 52: 270, 54: 413, 56: 419, 418, 270, 270, 270, 270, 270, 270, 270, 270, 414, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270},
		{412, 417, 3: 271, 420, 421, 416, 415, 271, 271, 271, 12: 411, 41: 271, 271, 49: 271, 271, 52: 271, 54: 413, 56: 419, 418, 271, 271, 271, 271, 271, 271, 271, 271, 414, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{412, 417, 3: 272, 272, 272, 416, 415, 272, 272, 272, 12: 411, 41: 272, 272, 49: 272, 272, 52: 272, 54: 413, 56: 419, 418, 272, 272, 272, 272, 272, 272, 272, 272, 414, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272},
		// 165
		{412, 417, 3: 273, 273, 273, 416, 415, 273, 273, 273, 12: 411, 41: 273, 273, 49: 273, 273, 52: 273, 54: 413, 56: 419, 418, 273, 273, 273, 273, 273, 273, 273, 273, 414, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273},
		{412, 274, 3: 274, 274, 274, 416, 415, 274, 274, 274, 12: 411, 41: 274, 274, 49: 274, 274, 52: 274, 54: 413, 56: 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 414, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274},
		{412, 275, 3: 275, 275, 275, 416, 415, 275, 275, 275, 12: 411, 41: 275, 275, 49: 275, 275, 52: 275, 54: 413, 56: 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 414, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275},
		{412, 276, 3: 276, 276, 276, 416, 415, 276, 276, 276, 12: 411, 41: 276, 276, 49: 276, 276, 52: 276, 54: 413, 56: 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 414, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276},
		{291, 291, 3: 291, 291, 291, 291, 291, 291, 291, 291, 12: 291, 41: 291, 291, 49: 291, 291, 52: 291, 54: 291, 56: 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291},
		// 170
		{292, 292, 3: 292, 292, 292, 292, 292, 292, 292, 292, 12: 292, 41: 292, 292, 49: 292, 292, 52: 292, 54: 292, 56: 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292},
		{412, 417, 3: 430, 420, 421, 416, 415, 306, 10: 306, 12: 411, 50: 436, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{8: 303, 10: 489},
		{8: 488},
		{293, 293, 3: 293, 293, 293, 293, 293, 293, 293, 293, 12: 293, 41: 293, 293, 49: 293, 293, 52: 293, 54: 293, 56: 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293},
		// 175
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 490},
		{412, 417, 3: 430, 420, 421, 416, 415, 305, 10: 305, 12: 411, 50: 436, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{10: 462, 52: 492},
		{294, 294, 3: 294, 294, 294, 294, 294, 294, 294, 294, 12: 294, 41: 294, 294, 49: 294, 294, 52: 294, 54: 294, 56: 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294},
		{154, 154, 154, 154, 154, 154, 154, 154, 154, 10: 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 43: 154, 52: 154},
		// 180
		{130, 8: 130, 10: 130, 12: 130},
		{52: 499},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 497},
		{412, 417, 3: 430, 420, 421, 416, 415, 12: 411, 50: 436, 52: 498, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{127, 8: 127, 10: 127, 12: 127},
		// 185
		{129, 8: 129, 10: 129, 12: 129},
		{412, 284, 3: 284, 284, 284, 416, 415, 284, 284, 284, 12: 411, 41: 284, 284, 49: 284, 284, 52: 284, 54: 413, 56: 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 414, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284},
		{126, 8: 126, 10: 126, 12: 126},
		{8: 545},
		{8: 149, 22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 111: 324, 115: 343, 346, 342, 323, 121: 507, 325, 125: 322, 157: 506, 170: 504, 505, 191: 508},
		// 190
		{8: 151, 10: 542},
		{8: 148},
		{8: 147, 10: 147},
		{157, 393, 157, 8: 133, 10: 133, 12: 157, 134: 394, 510, 137: 511, 152: 397, 174: 512},
		{8: 509},
		// 195
		{124, 8: 124, 10: 124, 12: 124},
		{515, 2: 514, 12: 121, 178: 398, 400, 513},
		{8: 145, 10: 145},
		{8: 144, 10: 144},
		{519, 8: 171, 171, 171, 12: 518, 22: 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 43: 171, 171, 171, 171, 171, 171, 50: 171, 53: 171, 55: 171},
		// 200
		{168, 8: 168, 168, 168, 12: 168, 22: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 43: 168, 168, 168, 168, 168, 168, 50: 168, 53: 168, 55: 168},
		{157, 393, 157, 8: 125, 12: 157, 22: 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 43: 125, 125, 125, 125, 125, 125, 134: 394, 510, 137: 516, 152: 502, 172: 503},
		{8: 517},
		{167, 8: 167, 167, 167, 12: 167, 22: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 43: 167, 167, 167, 167, 167, 167, 50: 167, 53: 167, 55: 167},
		{153, 153, 153, 153, 153, 153, 153, 153, 11: 153, 13: 153, 153, 153, 153, 153, 153, 153, 153, 153, 352, 350, 351, 43: 530, 52: 153, 111: 404, 150: 531, 159: 529},
		// 205
		{2: 522, 8: 141, 22: 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 43: 162, 162, 162, 162, 162, 162, 167: 523, 186: 521, 209: 520},
		{22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 111: 324, 115: 343, 346, 342, 323, 121: 507, 325, 125: 322, 157: 506, 170: 504, 527},
		{8: 526},
		{8: 143, 10: 143, 145: 143},
		{8: 140, 10: 524},
		// 210
		{2: 525},
		{8: 142, 10: 142, 145: 142},
		{160, 8: 160, 160, 160, 12: 160, 22: 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 43: 160, 160, 160, 160, 160, 160, 50: 160, 53: 160, 55: 160},
		{8: 528},
		{161, 8: 161, 161, 161, 12: 161, 22: 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 43: 161, 161, 161, 161, 161, 161, 50: 161, 53: 161, 55: 161},
		// 215
		{375, 538, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 402, 245, 166: 539},
		{153, 153, 153, 153, 153, 153, 153, 153, 11: 153, 13: 153, 153, 153, 153, 153, 153, 153, 153, 153, 352, 350, 351, 111: 404, 150: 408, 159: 535},
		{152, 152, 152, 152, 152, 152, 152, 152, 11: 152, 13: 152, 152, 152, 152, 152, 152, 152, 152, 152, 352, 350, 351, 43: 532, 52: 152, 111: 493},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 533},
		{412, 417, 3: 430, 420, 421, 416, 415, 12: 411, 50: 436, 52: 534, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		// 220
		{164, 8: 164, 164, 164, 12: 164, 22: 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 43: 164, 164, 164, 164, 164, 164, 50: 164, 53: 164, 55: 164},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 536},
		{412, 417, 3: 430, 420, 421, 416, 415, 12: 411, 50: 436, 52: 537, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{165, 8: 165, 165, 165, 12: 165, 22: 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 43: 165, 165, 165, 165, 165, 165, 50: 165, 53: 165, 55: 165},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 500, 541},
		// 225
		{52: 540},
		{166, 8: 166, 166, 166, 12: 166, 22: 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 43: 166, 166, 166, 166, 166, 166, 50: 166, 53: 166, 55: 166},
		{163, 8: 163, 163, 163, 12: 163, 22: 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 43: 163, 163, 163, 163, 163, 163, 50: 163, 53: 163, 55: 163},
		{22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 111: 324, 115: 343, 346, 342, 323, 121: 507, 325, 125: 322, 145: 543, 157: 544},
		{8: 150},
		// 230
		{8: 146, 10: 146},
		{131, 8: 131, 10: 131, 12: 131},
		{8: 123, 22: 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 43: 123, 123, 123, 123, 123, 123, 198: 547},
		{8: 149, 22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 111: 324, 115: 343, 346, 342, 323, 121: 507, 325, 125: 322, 157: 506, 170: 504, 505, 191: 548},
		{8: 549},
		// 235
		{122, 8: 122, 10: 122, 12: 122},
		{159, 393, 159, 8: 159, 10: 159, 12: 159, 134: 551},
		{158, 2: 158, 8: 158, 10: 158, 12: 158},
		{191, 191, 191, 8: 191, 191, 12: 191, 41: 191},
		{189, 189, 189, 8: 189, 189, 12: 189, 41: 189},
		// 240
		{192, 192, 192, 8: 192, 192, 12: 192, 41: 192},
		{246, 246, 3: 246, 246, 246, 246, 246, 246, 246, 246, 12: 246, 41: 246, 246, 49: 246, 246, 52: 246, 54: 246, 56: 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 51: 459, 112: 558, 151: 389, 158: 559},
		{412, 279, 3: 279, 279, 279, 416, 415, 279, 279, 279, 12: 411, 41: 279, 279, 49: 279, 279, 52: 279, 54: 413, 56: 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 414, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279},
		{8: 586, 10: 462},
		// 245
		{8: 560},
		{375, 278, 368, 278, 278, 278, 378, 377, 278, 278, 278, 374, 278, 384, 383, 386, 369, 370, 371, 372, 373, 385, 41: 278, 278, 49: 278, 278, 561, 278, 562, 278, 56: 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278},
		{412, 277, 3: 277, 277, 277, 416, 415, 277, 277, 277, 12: 411, 41: 277, 277, 49: 277, 277, 52: 277, 54: 413, 56: 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 414, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277},
		{114, 114, 114, 114, 114, 114, 114, 114, 11: 114, 568, 114, 114, 114, 114, 114, 114, 114, 114, 114, 53: 114, 569, 156: 567, 163: 566, 564, 565, 190: 563},
		{10: 579, 42: 184, 162: 584},
		// 250
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 575, 53: 576, 169: 577},
		{12: 568, 50: 573, 54: 569, 156: 574},
		{113, 113, 113, 113, 113, 113, 113, 113, 11: 113, 13: 113, 113, 113, 113, 113, 113, 113, 113, 113, 53: 113},
		{12: 112, 50: 112, 54: 112},
		{239, 239, 239, 239, 239, 239, 239, 239, 11: 239, 13: 239, 239, 239, 239, 239, 239, 239, 239, 239, 143: 362, 571},
		// 255
		{2: 570},
		{12: 109, 50: 109, 54: 109},
		{52: 572},
		{12: 110, 50: 110, 54: 110},
		{115, 115, 115, 115, 115, 115, 115, 115, 11: 115, 13: 115, 115, 115, 115, 115, 115, 115, 115, 115, 53: 115},
		// 260
		{12: 111, 50: 111, 54: 111},
		{412, 417, 3: 430, 420, 421, 416, 415, 9: 119, 119, 12: 411, 42: 119, 50: 436, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{114, 114, 114, 114, 114, 114, 114, 114, 11: 114, 568, 114, 114, 114, 114, 114, 114, 114, 114, 114, 53: 114, 569, 156: 567, 163: 566, 564, 565, 190: 578},
		{10: 117, 42: 117},
		{10: 579, 42: 184, 162: 580},
		// 265
		{114, 114, 114, 114, 114, 114, 114, 114, 11: 114, 568, 114, 114, 114, 114, 114, 114, 114, 114, 114, 42: 183, 53: 114, 569, 156: 567, 163: 566, 582, 565},
		{42: 581},
		{9: 118, 118, 42: 118},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 575, 53: 576, 169: 583},
		{10: 116, 42: 116},
		// 270
		{42: 585},
		{288, 288, 3: 288, 288, 288, 288, 288, 288, 288, 288, 12: 288, 41: 288, 288, 49: 288, 288, 52: 288, 54: 288, 56: 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288},
		{295, 295, 3: 295, 295, 295, 295, 295, 295, 295, 295, 12: 295, 41: 295, 295, 49: 295, 295, 52: 295, 54: 295, 56: 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295},
		{412, 280, 3: 280, 280, 280, 416, 415, 280, 280, 280, 12: 411, 41: 280, 280, 49: 280, 280, 52: 280, 54: 413, 56: 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 414, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280},
		{412, 281, 3: 281, 281, 281, 416, 415, 281, 281, 281, 12: 411, 41: 281, 281, 49: 281, 281, 52: 281, 54: 413, 56: 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 414, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281},
		// 275
		{412, 282, 3: 282, 282, 282, 416, 415, 282, 282, 282, 12: 411, 41: 282, 282, 49: 282, 282, 52: 282, 54: 413, 56: 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 414, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282},
		{412, 283, 3: 283, 283, 283, 416, 415, 283, 283, 283, 12: 411, 41: 283, 283, 49: 283, 283, 52: 283, 54: 413, 56: 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 414, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283},
		{412, 285, 3: 285, 285, 285, 416, 415, 285, 285, 285, 12: 411, 41: 285, 285, 49: 285, 285, 52: 285, 54: 413, 56: 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 414, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285},
		{412, 286, 3: 286, 286, 286, 416, 415, 286, 286, 286, 12: 411, 41: 286, 286, 49: 286, 286, 52: 286, 54: 413, 56: 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 414, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286},
		{412, 287, 3: 287, 287, 287, 416, 415, 287, 287, 287, 12: 411, 41: 287, 287, 49: 287, 287, 52: 287, 54: 413, 56: 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 414, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287},
		// 280
		{8: 595},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 561, 53: 562},
		{598},
		{57, 84: 57},
		{11: 599, 160: 600},
		// 285
		{8: 61, 11: 61, 41: 61},
		{8: 602, 11: 601},
		{8: 60, 11: 60, 41: 60},
		{59, 59, 59, 59, 59, 59, 59, 59, 9: 59, 11: 59, 13: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 42: 59, 59, 59, 59, 59, 59, 59, 53: 59, 55: 59, 84: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 110: 59},
		{22: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 43: 74, 74, 74, 74, 74, 74, 74, 55: 74, 124: 74},
		// 290
		{53: 182, 208: 606},
		{180, 180, 180, 8: 180, 180, 180, 12: 180, 22: 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 43: 180, 180, 180, 180, 180, 180, 53: 138},
		{53: 607},
		{2: 608, 182: 611, 610, 222: 609},
		{10: 307, 42: 307, 50: 307},
		// 295
		{10: 614, 42: 184, 162: 615},
		{10: 179, 42: 179},
		{10: 177, 42: 177, 50: 612},
		{239, 239, 239, 239, 239, 239, 239, 239, 11: 239, 13: 239, 239, 239, 239, 239, 239, 239, 239, 239, 143: 362, 613},
		{10: 176, 42: 176},
		// 300
		{2: 608, 42: 183, 182: 611, 617},
		{42: 616},
		{181, 181, 181, 8: 181, 181, 181, 12: 181, 22: 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 43: 181, 181, 181, 181, 181, 181},
		{10: 178, 42: 178},
		{53: 202, 206: 620},
		// 305
		{199, 199, 199, 8: 199, 199, 199, 12: 199, 22: 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 43: 199, 199, 199, 199, 199, 199, 53: 138},
		{53: 621},
		{22: 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 207: 622},
		{22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 111: 391, 115: 343, 346, 342, 390, 149: 625, 194: 624, 228: 623},
		{22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 42: 639, 111: 391, 115: 343, 346, 342, 390, 149: 625, 194: 640},
		// 310
		{22: 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 42: 196},
		{157, 393, 157, 9: 627, 41: 170, 134: 632, 631, 137: 629, 177: 630, 195: 628, 229: 626},
		{9: 636, 637},
		{22: 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 42: 193},
		{9: 188, 188},
		// 315
		{9: 186, 186, 41: 169},
		{41: 634},
		{633, 2: 514, 180: 513},
		{156, 2: 156},
		{157, 393, 157, 134: 632, 631, 137: 516},
		// 320
		{239, 239, 239, 239, 239, 239, 239, 239, 11: 239, 13: 239, 239, 239, 239, 239, 239, 239, 239, 239, 143: 362, 635},
		{9: 185, 185},
		{22: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 42: 194},
		{157, 393, 157, 41: 170, 134: 632, 631, 137: 629, 177: 630, 195: 638},
		{9: 187, 187},
		// 325
		{200, 200, 200, 8: 200, 200, 200, 12: 200, 22: 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 43: 200, 200, 200, 200, 200, 200},
		{22: 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 42: 195},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 51: 642, 151: 389, 158: 643},
		{412, 417, 3: 430, 420, 421, 416, 415, 645, 12: 411, 50: 436, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{8: 644},
		// 330
		{203, 203, 203, 8: 203, 203, 203, 12: 203, 22: 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 43: 203, 203, 203, 203, 203, 203},
		{204, 204, 204, 8: 204, 204, 204, 12: 204, 22: 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 43: 204, 204, 204, 204, 204, 204},
		{233, 233, 233, 8: 233, 233, 233, 12: 233},
		{231, 231, 231, 8: 231, 231, 231, 12: 231},
		{234, 234, 234, 8: 234, 234, 234, 12: 234},
		// 335
		{235, 235, 235, 8: 235, 235, 235, 12: 235},
		{236, 236, 236, 8: 236, 236, 236, 12: 236},
		{9: 791},
		{9: 230, 230},
		{9: 227, 789},
		// 340
		{9: 226, 226, 22: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 43: 63, 63, 63, 63, 63, 63, 50: 225, 53: 64, 55: 64, 173: 655, 203: 657, 216: 656},
		{50: 787},
		{53: 72, 55: 72, 200: 663},
		{22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 111: 324, 115: 343, 346, 342, 323, 121: 658, 325, 125: 322, 142: 659, 215: 660},
		{157, 393, 157, 9: 228, 134: 632, 631, 137: 662, 168: 652, 188: 653, 651},
		// 345
		{22: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 43: 66, 66, 66, 66, 66, 66, 53: 66, 55: 66},
		{22: 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 53: 62, 55: 62, 111: 324, 115: 343, 346, 342, 323, 121: 658, 325, 125: 322, 142: 661},
		{22: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 43: 65, 65, 65, 65, 65, 65, 53: 65, 55: 65},
		{9: 226, 226, 50: 225, 173: 655},
		{53: 70, 55: 68, 201: 665, 666, 223: 664},
		// 350
		{22: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 43: 71, 71, 71, 71, 71, 71, 71, 55: 71, 124: 71},
		{53: 708, 127: 709},
		{55: 668, 120: 669, 126: 667},
		{9: 707},
		{58, 22: 597, 84: 58, 196: 670},
		// 355
		{49, 49, 49, 49, 49, 49, 49, 49, 9: 49, 11: 49, 13: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 42: 49, 49, 49, 49, 49, 49, 49, 53: 49, 55: 49, 84: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 110: 49},
		{671, 84: 672},
		{11: 599, 160: 697},
		{673},
		{11: 599, 160: 674},
		// 360
		{11: 601, 41: 675},
		{41: 676},
		{11: 53, 680, 153: 678, 677, 161: 679},
		{11: 693},
		{8: 55, 10: 55, 41: 55},
		// 365
		{10: 683, 41: 684},
		{2: 681},
		{52: 682},
		{11: 52},
		{11: 53, 680, 153: 692, 677},
		// 370
		{11: 685, 176: 686},
		{8: 51, 10: 51, 41: 51},
		{10: 687, 41: 688},
		{11: 691},
		{2: 522, 167: 689},
		// 375
		{8: 690, 10: 524},
		{45, 45, 45, 45, 45, 45, 45, 45, 9: 45, 11: 45, 13: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 42: 45, 45, 45, 45, 45, 45, 45, 53: 45, 55: 45, 84: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 110: 45},
		{8: 50, 10: 50, 41: 50},
		{8: 54, 10: 54, 41: 54},
		{694},
		// 380
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 695},
		{412, 417, 3: 430, 420, 421, 416, 415, 696, 12: 411, 50: 436, 54: 413, 56: 419, 418, 424, 425, 435, 431, 432, 440, 433, 444, 414, 438, 428, 427, 426, 422, 442, 439, 437, 429, 446, 434, 423, 443, 441, 445},
		{8: 56, 10: 56, 41: 56},
		{8: 602, 11: 601, 41: 698},
		{11: 53, 680, 153: 678, 677, 161: 699},
		// 385
		{8: 700, 10: 683, 41: 701},
		{48, 48, 48, 48, 48, 48, 48, 48, 9: 48, 11: 48, 13: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 42: 48, 48, 48, 48, 48, 48, 48, 53: 48, 55: 48, 84: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 110: 48},
		{11: 53, 680, 153: 678, 677, 161: 702},
		{8: 703, 10: 683, 41: 704},
		{47, 47, 47, 47, 47, 47, 47, 47, 9: 47, 11: 47, 13: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 42: 47, 47, 47, 47, 47, 47, 47, 53: 47, 55: 47, 84: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 110: 47},
		// 390
		{11: 685, 176: 705},
		{8: 706, 10: 687},
		{46, 46, 46, 46, 46, 46, 46, 46, 9: 46, 11: 46, 13: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 42: 46, 46, 46, 46, 46, 46, 46, 53: 46, 55: 46, 84: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 110: 46},
		{22: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 43: 67, 67, 67, 67, 67, 67, 67, 55: 67, 124: 67},
		{98, 98, 98, 98, 98, 98, 98, 98, 9: 98, 11: 98, 13: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 42: 98, 98, 98, 98, 98, 98, 98, 53: 98, 55: 98, 84: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 199: 710},
		// 395
		{22: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 43: 69, 69, 69, 69, 69, 69, 69, 55: 69, 124: 69},
		{375, 380, 711, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 42: 94, 328, 329, 327, 353, 330, 326, 51: 459, 53: 708, 55: 668, 84: 733, 730, 735, 720, 734, 721, 731, 732, 728, 736, 729, 111: 324, 712, 115: 343, 346, 342, 323, 727, 669, 658, 325, 125: 322, 719, 714, 715, 717, 718, 713, 716, 726, 142: 725, 175: 723, 212: 724, 722},
		{302, 302, 3: 302, 302, 302, 302, 302, 9: 302, 302, 12: 302, 41: 785, 50: 302, 54: 302, 56: 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302},
		{8: 240, 240, 462},
		{108, 108, 108, 108, 108, 108, 108, 108, 9: 108, 11: 108, 13: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 42: 108, 108, 108, 108, 108, 108, 108, 53: 108, 55: 108, 84: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 110: 108},
		// 400
		{107, 107, 107, 107, 107, 107, 107, 107, 9: 107, 11: 107, 13: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 42: 107, 107, 107, 107, 107, 107, 107, 53: 107, 55: 107, 84: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 110: 107},
		{106, 106, 106, 106, 106, 106, 106, 106, 9: 106, 11: 106, 13: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 42: 106, 106, 106, 106, 106, 106, 106, 53: 106, 55: 106, 84: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 110: 106},
		{105, 105, 105, 105, 105, 105, 105, 105, 9: 105, 11: 105, 13: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 42: 105, 105, 105, 105, 105, 105, 105, 53: 105, 55: 105, 84: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 110: 105},
		{104, 104, 104, 104, 104, 104, 104, 104, 9: 104, 11: 104, 13: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 42: 104, 104, 104, 104, 104, 104, 104, 53: 104, 55: 104, 84: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 110: 104},
		{103, 103, 103, 103, 103, 103, 103, 103, 9: 103, 11: 103, 13: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 42: 103, 103, 103, 103, 103, 103, 103, 53: 103, 55: 103, 84: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 110: 103},
		// 405
		{102, 102, 102, 102, 102, 102, 102, 102, 9: 102, 11: 102, 13: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 42: 102, 102, 102, 102, 102, 102, 102, 53: 102, 55: 102, 84: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 110: 102},
		{239, 239, 239, 239, 239, 239, 239, 239, 11: 239, 13: 239, 239, 239, 239, 239, 239, 239, 239, 239, 143: 362, 782},
		{41: 780},
		{42: 779},
		{96, 96, 96, 96, 96, 96, 96, 96, 9: 96, 11: 96, 13: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 42: 96, 96, 96, 96, 96, 96, 96, 53: 96, 55: 96, 84: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96},
		// 410
		{375, 380, 711, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 42: 93, 328, 329, 327, 353, 330, 326, 51: 459, 53: 708, 55: 668, 84: 733, 730, 735, 720, 734, 721, 731, 732, 728, 736, 729, 111: 324, 712, 115: 343, 346, 342, 323, 727, 669, 658, 325, 125: 322, 719, 714, 715, 717, 718, 713, 716, 726, 142: 725, 175: 778},
		{92, 92, 92, 92, 92, 92, 92, 92, 9: 92, 11: 92, 13: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 42: 92, 92, 92, 92, 92, 92, 92, 53: 92, 55: 92, 84: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92},
		{91, 91, 91, 91, 91, 91, 91, 91, 9: 91, 11: 91, 13: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 42: 91, 91, 91, 91, 91, 91, 91, 53: 91, 55: 91, 84: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91},
		{9: 777},
		{771},
		// 415
		{767},
		{763},
		{375, 380, 711, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 53: 708, 55: 668, 84: 733, 730, 735, 720, 734, 721, 731, 732, 728, 736, 729, 112: 712, 119: 727, 669, 126: 719, 714, 715, 717, 718, 713, 716, 757},
		{743},
		{2: 741},
		// 420
		{9: 740},
		{9: 739},
		{375, 380, 368, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 112: 712, 119: 737},
		{9: 738},
		{79, 79, 79, 79, 79, 79, 79, 79, 9: 79, 11: 79, 13: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 42: 79, 79, 79, 79, 79, 79, 79, 53: 79, 55: 79, 84: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 110: 79},
		// 425
		{80, 80, 80, 80, 80, 80, 80, 80, 9: 80, 11: 80, 13: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 42: 80, 80, 80, 80, 80, 80, 80, 53: 80, 55: 80, 84: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 110: 80},
		{81, 81, 81, 81, 81, 81, 81, 81, 9: 81, 11: 81, 13: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 42: 81, 81, 81, 81, 81, 81, 81, 53: 81, 55: 81, 84: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 110: 81},
		{9: 742},
		{82, 82, 82, 82, 82, 82, 82, 82, 9: 82, 11: 82, 13: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 42: 82, 82, 82, 82, 82, 82, 82, 53: 82, 55: 82, 84: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 110: 82},
		{375, 380, 368, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 352, 350, 351, 340, 332, 341, 337, 349, 336, 334, 335, 333, 338, 347, 344, 345, 348, 339, 331, 43: 328, 329, 327, 353, 330, 326, 51: 459, 111: 324, 712, 115: 343, 346, 342, 323, 744, 121: 658, 325, 125: 322, 142: 745},
		// 430
		{9: 751},
		{375, 380, 368, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 112: 712, 119: 746},
		{9: 747},
		{375, 380, 368, 379, 381, 382, 378, 377, 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 112: 712, 119: 748},
		{8: 749},
		// 435
		{375, 380, 711, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 53: 708, 55: 668, 84: 733, 730, 735, 720, 734, 721, 731, 732, 728, 736, 729, 112: 712, 119: 727, 669, 126: 719, 714, 715, 717, 718, 713, 716, 750},
		{83, 83, 83, 83, 83, 83, 83, 83, 9: 83, 11: 83, 13: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 42: 83, 83, 83, 83, 83, 83, 83, 53: 83, 55: 83, 84: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 110: 83},
		{375, 380, 368, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 112: 712, 119: 752},
		{9: 753},
		{375, 380, 368, 379, 381, 382, 378, 377, 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 112: 712, 119: 754},
		// 440
		{8: 755},
		{375, 380, 711, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 53: 708, 55: 668, 84: 733, 730, 735, 720, 734, 721, 731, 732, 728, 736, 729, 112: 712, 119: 727, 669, 126: 719, 714, 715, 717, 718, 713, 716, 756},
		{84, 84, 84, 84, 84, 84, 84, 84, 9: 84, 11: 84, 13: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 42: 84, 84, 84, 84, 84, 84, 84, 53: 84, 55: 84, 84: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 110: 84},
		{85: 758},
		{759},
		// 445
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 112: 760},
		{8: 761, 10: 462},
		{9: 762},
		{85, 85, 85, 85, 85, 85, 85, 85, 9: 85, 11: 85, 13: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 42: 85, 85, 85, 85, 85, 85, 85, 53: 85, 55: 85, 84: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 110: 85},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 112: 764},
		// 450
		{8: 765, 10: 462},
		{375, 380, 711, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 53: 708, 55: 668, 84: 733, 730, 735, 720, 734, 721, 731, 732, 728, 736, 729, 112: 712, 119: 727, 669, 126: 719, 714, 715, 717, 718, 713, 716, 766},
		{86, 86, 86, 86, 86, 86, 86, 86, 9: 86, 11: 86, 13: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 42: 86, 86, 86, 86, 86, 86, 86, 53: 86, 55: 86, 84: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 110: 86},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 112: 768},
		{8: 769, 10: 462},
		// 455
		{375, 380, 711, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 53: 708, 55: 668, 84: 733, 730, 735, 720, 734, 721, 731, 732, 728, 736, 729, 112: 712, 119: 727, 669, 126: 719, 714, 715, 717, 718, 713, 716, 770},
		{87, 87, 87, 87, 87, 87, 87, 87, 9: 87, 11: 87, 13: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 42: 87, 87, 87, 87, 87, 87, 87, 53: 87, 55: 87, 84: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 110: 87},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 112: 772},
		{8: 773, 10: 462},
		{375, 380, 711, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 53: 708, 55: 668, 84: 733, 730, 735, 720, 734, 721, 731, 732, 728, 736, 729, 112: 712, 119: 727, 669, 126: 719, 714, 715, 717, 718, 713, 716, 774},
		// 460
		{89, 89, 89, 89, 89, 89, 89, 89, 9: 89, 11: 89, 13: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 42: 89, 89, 89, 89, 89, 89, 89, 53: 89, 55: 89, 84: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 110: 775},
		{375, 380, 711, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 53: 708, 55: 668, 84: 733, 730, 735, 720, 734, 721, 731, 732, 728, 736, 729, 112: 712, 119: 727, 669, 126: 719, 714, 715, 717, 718, 713, 716, 776},
		{88, 88, 88, 88, 88, 88, 88, 88, 9: 88, 11: 88, 13: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 42: 88, 88, 88, 88, 88, 88, 88, 53: 88, 55: 88, 84: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 110: 88},
		{90, 90, 90, 90, 90, 90, 90, 90, 9: 90, 11: 90, 13: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 42: 90, 90, 90, 90, 90, 90, 90, 53: 90, 55: 90, 84: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 110: 90},
		{95, 95, 95, 95, 95, 95, 95, 95, 9: 95, 11: 95, 13: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 42: 95, 95, 95, 95, 95, 95, 95, 53: 95, 55: 95, 84: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95},
		// 465
		{97, 97, 97, 97, 97, 97, 97, 97, 9: 97, 11: 97, 13: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 42: 97, 97, 97, 97, 97, 97, 97, 97, 53: 97, 55: 97, 84: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 110: 97, 124: 97},
		{375, 380, 711, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 53: 708, 55: 668, 84: 733, 730, 735, 720, 734, 721, 731, 732, 728, 736, 729, 112: 712, 119: 727, 669, 126: 719, 714, 715, 717, 718, 713, 716, 781},
		{99, 99, 99, 99, 99, 99, 99, 99, 9: 99, 11: 99, 13: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 42: 99, 99, 99, 99, 99, 99, 99, 53: 99, 55: 99, 84: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 110: 99},
		{41: 783},
		{375, 380, 711, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 53: 708, 55: 668, 84: 733, 730, 735, 720, 734, 721, 731, 732, 728, 736, 729, 112: 712, 119: 727, 669, 126: 719, 714, 715, 717, 718, 713, 716, 784},
		// 470
		{100, 100, 100, 100, 100, 100, 100, 100, 9: 100, 11: 100, 13: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 42: 100, 100, 100, 100, 100, 100, 100, 53: 100, 55: 100, 84: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 110: 100},
		{375, 380, 711, 379, 381, 382, 378, 377, 9: 241, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 459, 53: 708, 55: 668, 84: 733, 730, 735, 720, 734, 721, 731, 732, 728, 736, 729, 112: 712, 119: 727, 669, 126: 719, 714, 715, 717, 718, 713, 716, 786},
		{101, 101, 101, 101, 101, 101, 101, 101, 9: 101, 11: 101, 13: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 42: 101, 101, 101, 101, 101, 101, 101, 53: 101, 55: 101, 84: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 110: 101},
		{375, 380, 368, 379, 381, 382, 378, 377, 11: 374, 13: 384, 383, 386, 369, 370, 371, 372, 373, 385, 51: 575, 53: 576, 169: 788},
		{9: 224, 224},
		// 475
		{157, 393, 157, 134: 632, 631, 137: 662, 168: 790},
		{9: 229, 229},
		{237, 237, 237, 237, 237, 237, 237, 237, 9: 237, 11: 237, 13: 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 42: 237, 237, 237, 237, 237, 237, 237, 237, 53: 237, 55: 237, 84: 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 237, 124: 237},
		{22: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 43: 77, 77, 77, 77, 77, 77, 77, 55: 77, 124: 77},
		{239, 239, 239, 239, 239, 239, 239, 239, 11: 239, 13: 239, 239, 239, 239, 239, 239, 239, 239, 239, 143: 362, 794},
		// 480
		{49: 310},
		{82: 817, 819, 98: 807, 808, 809, 804, 805, 806, 810, 814, 811, 801, 812, 813, 113: 818, 816, 123: 815, 136: 799, 138: 798, 803, 800, 802, 146: 797, 226: 796},
		{49: 312},
		{49: 43, 82: 817, 819, 98: 807, 808, 809, 804, 805, 806, 810, 814, 811, 801, 812, 813, 113: 818, 816, 123: 815, 136: 799, 138: 857, 803, 800, 802},
		{49: 42, 82: 42, 42, 95: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		// 485
		{49: 38, 82: 38, 38, 95: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{49: 37, 82: 37, 37, 95: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{83: 819, 113: 879, 816},
		{49: 35, 82: 35, 35, 95: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{95: 28, 28, 867, 181: 865, 217: 866, 864},
		// 490
		{83: 819, 113: 861, 816},
		{2: 858},
		{2: 853},
		{2: 834, 82: 836, 224: 835},
		{82: 817, 819, 113: 818, 816, 123: 833},
		// 495
		{49: 16, 82: 16, 16, 95: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{83: 819, 113: 831, 816},
		{83: 819, 113: 829, 816},
		{82: 817, 819, 113: 818, 816, 123: 828},
		{2: 824},
		// 500
		{83: 819, 113: 822, 816},
		{49: 7, 82: 7, 7, 95: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{82: 5, 821},
		{49: 4, 82: 4, 4, 95: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{82: 820},
		// 505
		{82: 2, 2},
		{49: 3, 82: 3, 3, 95: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{82: 1, 1},
		{82: 823},
		{49: 8, 82: 8, 8, 95: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		// 510
		{82: 825, 819, 113: 826, 816},
		{49: 12, 82: 12, 12, 95: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{82: 827},
		{49: 9, 82: 9, 9, 95: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{49: 13, 82: 13, 13, 95: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		// 515
		{82: 830},
		{49: 14, 82: 14, 14, 95: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{82: 832},
		{49: 15, 82: 15, 15, 95: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{49: 17, 82: 17, 17, 95: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		// 520
		{82: 817, 819, 113: 818, 816, 123: 842, 148: 852},
		{2: 522, 8: 141, 145: 838, 167: 837, 186: 839},
		{49: 10, 82: 10, 10, 95: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{8: 140, 10: 845, 145: 846},
		{8: 843},
		// 525
		{8: 840},
		{82: 817, 819, 113: 818, 816, 123: 842, 148: 841},
		{49: 18, 82: 18, 18, 95: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{49: 6, 82: 6, 6, 95: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{82: 817, 819, 113: 818, 816, 123: 842, 148: 844},
		// 530
		{49: 20, 82: 20, 20, 95: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{2: 525, 145: 849},
		{8: 847},
		{82: 817, 819, 113: 818, 816, 123: 842, 148: 848},
		{49: 11, 82: 11, 11, 95: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		// 535
		{8: 850},
		{82: 817, 819, 113: 818, 816, 123: 842, 148: 851},
		{49: 19, 82: 19, 19, 95: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{49: 21, 82: 21, 21, 95: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{82: 854},
		// 540
		{82: 817, 819, 95: 40, 40, 40, 807, 808, 809, 804, 805, 806, 810, 814, 811, 801, 812, 813, 113: 818, 816, 123: 815, 136: 799, 138: 798, 803, 800, 802, 146: 855, 856},
		{82: 817, 819, 95: 39, 39, 39, 807, 808, 809, 804, 805, 806, 810, 814, 811, 801, 812, 813, 113: 818, 816, 123: 815, 136: 799, 138: 857, 803, 800, 802},
		{95: 31, 31, 31},
		{49: 41, 82: 41, 41, 95: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		{82: 859},
		// 545
		{82: 817, 819, 95: 40, 40, 40, 807, 808, 809, 804, 805, 806, 810, 814, 811, 801, 812, 813, 113: 818, 816, 123: 815, 136: 799, 138: 798, 803, 800, 802, 146: 855, 860},
		{95: 32, 32, 32},
		{82: 862},
		{82: 817, 819, 95: 40, 40, 40, 807, 808, 809, 804, 805, 806, 810, 814, 811, 801, 812, 813, 113: 818, 816, 123: 815, 136: 799, 138: 798, 803, 800, 802, 146: 855, 863},
		{95: 33, 33, 33},
		// 550
		{95: 24, 873, 219: 874, 872},
		{95: 30, 30, 30},
		{95: 27, 27, 867, 181: 871},
		{83: 819, 113: 868, 816},
		{82: 869},
		// 555
		{82: 817, 819, 95: 40, 40, 40, 807, 808, 809, 804, 805, 806, 810, 814, 811, 801, 812, 813, 113: 818, 816, 123: 815, 136: 799, 138: 798, 803, 800, 802, 146: 855, 870},
		{95: 26, 26, 26},
		{95: 29, 29, 29},
		{95: 878, 221: 877},
		{82: 875},
		// 560
		{95: 23},
		{82: 817, 819, 95: 40, 98: 807, 808, 809, 804, 805, 806, 810, 814, 811, 801, 812, 813, 113: 818, 816, 123: 815, 136: 799, 138: 798, 803, 800, 802, 146: 855, 876},
		{95: 25},
		{49: 34, 82: 34, 34, 95: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{49: 22, 82: 22, 22, 95: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		// 565
		{82: 880},
		{49: 36, 82: 36, 36, 95: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
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
	const yyError = 234

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
			lx := yylex.(*lexer)
			lhs := &Declarator{
				PointerOpt:       yyS[yypt-1].node.(*PointerOpt),
				DirectDeclarator: yyS[yypt-0].node.(*DirectDeclarator),
			}
			yyVAL.node = lhs
			lhs.specifier = lx.scope.specifier
			lhs.DirectDeclarator.declarator = lhs
		}
	case 144:
		{
			yyVAL.node = (*DeclaratorOpt)(nil)
		}
	case 145:
		{
			yyVAL.node = &DeclaratorOpt{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
		}
	case 146:
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
	case 147:
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
	case 148:
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
	case 149:
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
	case 150:
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
	case 151:
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
	case 152:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 153:
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
	case 154:
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
	case 155:
		{
			yyVAL.node = &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].node.(*TypeQualifierListOpt),
			}
		}
	case 156:
		{
			yyVAL.node = &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].node.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].node.(*Pointer),
			}
		}
	case 157:
		{
			yyVAL.node = (*PointerOpt)(nil)
		}
	case 158:
		{
			yyVAL.node = &PointerOpt{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
		}
	case 159:
		{
			lhs := &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.TypeQualifier.attr
		}
	case 160:
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
	case 161:
		{
			yyVAL.node = (*TypeQualifierListOpt)(nil)
		}
	case 162:
		{
			yyVAL.node = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].node.(*TypeQualifierList).reverse(),
			}
		}
	case 163:
		{
			lhs := &ParameterTypeList{
				ParameterList: yyS[yypt-0].node.(*ParameterList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 164:
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
	case 165:
		{
			yyVAL.node = (*ParameterTypeListOpt)(nil)
		}
	case 166:
		{
			yyVAL.node = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].node.(*ParameterTypeList),
			}
		}
	case 167:
		{
			yyVAL.node = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 168:
		{
			yyVAL.node = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].node.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 169:
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
	case 170:
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
	case 171:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 172:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 173:
		{
			yyVAL.node = (*IdentifierListOpt)(nil)
		}
	case 174:
		{
			yyVAL.node = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 175:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 176:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 177:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeBlock)
		}
	case 178:
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
	case 179:
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
	case 180:
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
	case 181:
		{
			yyVAL.node = (*AbstractDeclaratorOpt)(nil)
		}
	case 182:
		{
			yyVAL.node = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].node.(*AbstractDeclarator),
			}
		}
	case 183:
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
	case 184:
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
	case 185:
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
	case 186:
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
	case 187:
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
	case 188:
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
	case 189:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 190:
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
	case 191:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 192:
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
	case 193:
		{
			yyVAL.node = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 194:
		{
			yyVAL.node = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
		}
	case 195:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 196:
		{
			yyVAL.node = &Initializer{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 197:
		{
			yyVAL.node = &InitializerList{
				DesignationOpt: yyS[yypt-1].node.(*DesignationOpt),
				Initializer:    yyS[yypt-0].node.(*Initializer),
			}
		}
	case 198:
		{
			yyVAL.node = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].node.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].node.(*DesignationOpt),
				Initializer:     yyS[yypt-0].node.(*Initializer),
			}
		}
	case 199:
		{
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 200:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 201:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 202:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 203:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 204:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 205:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 206:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 207:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 208:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 209:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 210:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 211:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 212:
		{
			yyVAL.node = &Statement{
				Case:               6,
				AssemblerStatement: yyS[yypt-0].node.(*AssemblerStatement),
			}
		}
	case 213:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 214:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 215:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 216:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 217:
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
	case 218:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 219:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 220:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 221:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 222:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 223:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 224:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 225:
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
	case 226:
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
	case 227:
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
	case 228:
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
	case 229:
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
	case 230:
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
	case 231:
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
	case 232:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 233:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 234:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 235:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 236:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 237:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 238:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 239:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 240:
		{
			yyVAL.node = &ExternalDeclaration{
				Case: 2,
				BasicAssemblerStatement: yyS[yypt-1].node.(*BasicAssemblerStatement),
				Token: yyS[yypt-0].Token,
			}
		}
	case 241:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:         3,
				StaticAssert: yyS[yypt-0].node.(*StaticAssert),
			}
		}
	case 242:
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
	case 243:
		{
			yyVAL.node = &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:          yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 244:
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
	case 245:
		{
			lhs := &FunctionBody{
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
			yyVAL.node = lhs
			lhs.scope = lhs.CompoundStatement.scope
		}
	case 246:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 247:
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
	case 248:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 249:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 250:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 251:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 252:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScopePos(lhs.Pos())
		}
	case 253:
		{
			yyVAL.node = &AssemblerInstructions{
				Token: yyS[yypt-0].Token,
			}
		}
	case 254:
		{
			yyVAL.node = &AssemblerInstructions{
				Case: 1,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions),
				Token: yyS[yypt-0].Token,
			}
		}
	case 255:
		{
			yyVAL.node = &BasicAssemblerStatement{
				Token:                 yyS[yypt-4].Token,
				VolatileOpt:           yyS[yypt-3].node.(*VolatileOpt),
				Token2:                yyS[yypt-2].Token,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-0].Token,
			}
		}
	case 256:
		{
			yyVAL.node = (*VolatileOpt)(nil)
		}
	case 257:
		{
			yyVAL.node = &VolatileOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 258:
		{
			yyVAL.node = &AssemblerOperand{
				AssemblerSymbolicNameOpt: yyS[yypt-4].node.(*AssemblerSymbolicNameOpt),
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
		}
	case 259:
		{
			yyVAL.node = &AssemblerOperands{
				AssemblerOperand: yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 260:
		{
			yyVAL.node = &AssemblerOperands{
				Case:              1,
				AssemblerOperands: yyS[yypt-2].node.(*AssemblerOperands),
				Token:             yyS[yypt-1].Token,
				AssemblerOperand:  yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 261:
		{
			yyVAL.node = (*AssemblerSymbolicNameOpt)(nil)
		}
	case 262:
		{
			yyVAL.node = &AssemblerSymbolicNameOpt{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 263:
		{
			yyVAL.node = &Clobbers{
				Token: yyS[yypt-0].Token,
			}
		}
	case 264:
		{
			yyVAL.node = &Clobbers{
				Case:     1,
				Clobbers: yyS[yypt-2].node.(*Clobbers),
				Token:    yyS[yypt-1].Token,
				Token2:   yyS[yypt-0].Token,
			}
		}
	case 265:
		{
			yyVAL.node = &AssemblerStatement{
				BasicAssemblerStatement: yyS[yypt-0].node.(*BasicAssemblerStatement),
			}
		}
	case 266:
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
	case 267:
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
	case 268:
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
	case 269:
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
	case 270:
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
	case 271:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 272:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 273:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 274:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 275:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 276:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 277:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 278:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 279:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 280:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 281:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 282:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 283:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 284:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 285:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 286:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 287:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 288:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 289:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 290:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 291:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 292:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 293:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 294:
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
	case 295:
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
	case 296:
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
	case 297:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 298:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 299:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 300:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 301:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 302:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 303:
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
	case 304:
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
	case 305:
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
	case 306:
		{
			yyVAL.node = &ControlLine{
				Case:        13,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 309:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 310:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
