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
	yyDefault           = 57459
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
	yyTabOfs   = -325
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (326x)
		42:      1,   // '*' (287x)
		57377:   2,   // IDENTIFIER (238x)
		38:      3,   // '&' (219x)
		43:      4,   // '+' (219x)
		45:      5,   // '-' (219x)
		57363:   6,   // DEC (219x)
		57381:   7,   // INC (219x)
		59:      8,   // ';' (216x)
		41:      9,   // ')' (201x)
		44:      10,  // ',' (189x)
		57427:   11,  // STRINGLITERAL (166x)
		91:      12,  // '[' (163x)
		33:      13,  // '!' (147x)
		126:     14,  // '~' (147x)
		57347:   15,  // ALIGNOF (147x)
		57358:   16,  // CHARCONST (147x)
		57373:   17,  // FLOATCONST (147x)
		57384:   18,  // INTCONST (147x)
		57387:   19,  // LONGCHARCONST (147x)
		57388:   20,  // LONGSTRINGLITERAL (147x)
		57424:   21,  // SIZEOF (147x)
		57438:   22,  // VOLATILE (140x)
		57360:   23,  // CONST (138x)
		57418:   24,  // RESTRICT (138x)
		125:     25,  // '}' (129x)
		57353:   26,  // BOOL (128x)
		57357:   27,  // CHAR (128x)
		57359:   28,  // COMPLEX (128x)
		57367:   29,  // DOUBLE (128x)
		57369:   30,  // ENUM (128x)
		57372:   31,  // FLOAT (128x)
		57383:   32,  // INT (128x)
		57386:   33,  // LONG (128x)
		57422:   34,  // SHORT (128x)
		57423:   35,  // SIGNED (128x)
		57428:   36,  // STRUCT (128x)
		57432:   37,  // TYPEDEFNAME (128x)
		57433:   38,  // TYPEOF (128x)
		57435:   39,  // UNION (128x)
		57436:   40,  // UNSIGNED (128x)
		57437:   41,  // VOID (128x)
		58:      42,  // ':' (123x)
		57425:   43,  // STATIC (119x)
		57352:   44,  // AUTO (113x)
		57371:   45,  // EXTERN (113x)
		57382:   46,  // INLINE (113x)
		57395:   47,  // NORETURN (113x)
		57417:   48,  // REGISTER (113x)
		57431:   49,  // TYPEDEF (113x)
		57344:   50,  // $end (104x)
		61:      51,  // '=' (92x)
		123:     52,  // '{' (92x)
		57502:   53,  // Expression (86x)
		93:      54,  // ']' (84x)
		57351:   55,  // ASM (83x)
		46:      56,  // '.' (81x)
		57426:   57,  // STATIC_ASSERT (78x)
		37:      58,  // '%' (73x)
		47:      59,  // '/' (73x)
		60:      60,  // '<' (73x)
		62:      61,  // '>' (73x)
		63:      62,  // '?' (73x)
		94:      63,  // '^' (73x)
		124:     64,  // '|' (73x)
		57346:   65,  // ADDASSIGN (73x)
		57348:   66,  // ANDAND (73x)
		57349:   67,  // ANDASSIGN (73x)
		57350:   68,  // ARROW (73x)
		57365:   69,  // DIVASSIGN (73x)
		57370:   70,  // EQ (73x)
		57375:   71,  // GEQ (73x)
		57385:   72,  // LEQ (73x)
		57389:   73,  // LSH (73x)
		57390:   74,  // LSHASSIGN (73x)
		57391:   75,  // MODASSIGN (73x)
		57392:   76,  // MULASSIGN (73x)
		57393:   77,  // NEQ (73x)
		57396:   78,  // ORASSIGN (73x)
		57397:   79,  // OROR (73x)
		57420:   80,  // RSH (73x)
		57421:   81,  // RSHASSIGN (73x)
		57429:   82,  // SUBASSIGN (73x)
		57440:   83,  // XORASSIGN (73x)
		10:      84,  // '\n' (58x)
		57376:   85,  // GOTO (54x)
		57413:   86,  // PPOTHER (52x)
		57439:   87,  // WHILE (52x)
		57354:   88,  // BREAK (51x)
		57355:   89,  // CASE (51x)
		57361:   90,  // CONTINUE (51x)
		57364:   91,  // DEFAULT (51x)
		57366:   92,  // DO (51x)
		57374:   93,  // FOR (51x)
		57380:   94,  // IF (51x)
		57419:   95,  // RETURN (51x)
		57430:   96,  // SWITCH (51x)
		57401:   97,  // PPENDIF (44x)
		57400:   98,  // PPELSE (40x)
		57399:   99,  // PPELIF (39x)
		57398:   100, // PPDEFINE (35x)
		57402:   101, // PPERROR (35x)
		57403:   102, // PPHASH_NL (35x)
		57405:   103, // PPIF (35x)
		57406:   104, // PPIFDEF (35x)
		57407:   105, // PPIFNDEF (35x)
		57408:   106, // PPINCLUDE (35x)
		57409:   107, // PPINCLUDE_NEXT (35x)
		57410:   108, // PPLINE (35x)
		57411:   109, // PPNONDIRECTIVE (35x)
		57415:   110, // PPPRAGMA (35x)
		57416:   111, // PPUNDEF (35x)
		57368:   112, // ELSE (31x)
		57554:   113, // TypeQualifier (28x)
		57503:   114, // ExpressionList (26x)
		57527:   115, // PPTokenList (22x)
		57529:   116, // PPTokens (22x)
		57498:   117, // EnumSpecifier (20x)
		57549:   118, // StructOrUnion (20x)
		57550:   119, // StructOrUnionSpecifier (20x)
		57557:   120, // TypeSpecifier (20x)
		57504:   121, // ExpressionListOpt (18x)
		57469:   122, // BasicAssemblerStatement (15x)
		57475:   123, // CompoundStatement (15x)
		57481:   124, // DeclarationSpecifiers (15x)
		57510:   125, // FunctionSpecifier (15x)
		57528:   126, // PPTokenListOpt (15x)
		57544:   127, // StorageClassSpecifier (15x)
		57467:   128, // AssemblerStatement (13x)
		57506:   129, // ExpressionStatement (12x)
		57524:   130, // IterationStatement (12x)
		57525:   131, // JumpStatement (12x)
		57526:   132, // LabeledStatement (12x)
		57534:   133, // Pointer (12x)
		57538:   134, // SelectionStatement (12x)
		57542:   135, // Statement (12x)
		57535:   136, // PointerOpt (11x)
		57483:   137, // Declarator (9x)
		57543:   138, // StaticAssertDeclaration (9x)
		57477:   139, // ControlLine (8x)
		57513:   140, // GroupPart (8x)
		57517:   141, // IfGroup (8x)
		57518:   142, // IfSection (8x)
		57551:   143, // TextLine (8x)
		57478:   144, // Declaration (7x)
		57453:   145, // $@4 (6x)
		57476:   146, // ConstantExpression (6x)
		57362:   147, // DDD (6x)
		57511:   148, // GroupList (6x)
		57465:   149, // AssemblerOperand (5x)
		57468:   150, // AssemblerSymbolicNameOpt (5x)
		57512:   151, // GroupListOpt (5x)
		57537:   152, // ReplacementList (5x)
		57539:   153, // SpecifierQualifierList (5x)
		57555:   154, // TypeQualifierList (5x)
		57458:   155, // $@9 (4x)
		57460:   156, // AbstractDeclarator (4x)
		57466:   157, // AssemblerOperands (4x)
		57482:   158, // DeclarationSpecifiersOpt (4x)
		57487:   159, // Designator (4x)
		57522:   160, // Initializer (4x)
		57530:   161, // ParameterDeclaration (4x)
		57553:   162, // TypeName (4x)
		57556:   163, // TypeQualifierListOpt (4x)
		57464:   164, // AssemblerInstructions (3x)
		57474:   165, // CommaOpt (3x)
		57485:   166, // Designation (3x)
		57486:   167, // DesignationOpt (3x)
		57488:   168, // DesignatorList (3x)
		57505:   169, // ExpressionOpt (3x)
		57514:   170, // IdentifierList (3x)
		57519:   171, // InitDeclarator (3x)
		57531:   172, // ParameterList (3x)
		57532:   173, // ParameterTypeList (3x)
		57442:   174, // $@10 (2x)
		57446:   175, // $@14 (2x)
		57448:   176, // $@16 (2x)
		57449:   177, // $@17 (2x)
		57450:   178, // $@18 (2x)
		57454:   179, // $@5 (2x)
		57461:   180, // AbstractDeclaratorOpt (2x)
		57470:   181, // BlockItem (2x)
		57473:   182, // Clobbers (2x)
		57480:   183, // DeclarationListOpt (2x)
		57484:   184, // DeclaratorOpt (2x)
		57489:   185, // DirectAbstractDeclarator (2x)
		57490:   186, // DirectAbstractDeclaratorOpt (2x)
		57491:   187, // DirectDeclarator (2x)
		57492:   188, // ElifGroup (2x)
		57499:   189, // EnumerationConstant (2x)
		57500:   190, // Enumerator (2x)
		57507:   191, // ExternalDeclaration (2x)
		57508:   192, // FunctionBody (2x)
		57509:   193, // FunctionDefinition (2x)
		57515:   194, // IdentifierListOpt (2x)
		57516:   195, // IdentifierOpt (2x)
		57520:   196, // InitDeclaratorList (2x)
		57521:   197, // InitDeclaratorListOpt (2x)
		57523:   198, // InitializerList (2x)
		57533:   199, // ParameterTypeListOpt (2x)
		57540:   200, // SpecifierQualifierListOpt (2x)
		57545:   201, // StructDeclaration (2x)
		57547:   202, // StructDeclarator (2x)
		57558:   203, // VolatileOpt (2x)
		57441:   204, // $@1 (1x)
		57443:   205, // $@11 (1x)
		57444:   206, // $@12 (1x)
		57445:   207, // $@13 (1x)
		57447:   208, // $@15 (1x)
		57451:   209, // $@2 (1x)
		57452:   210, // $@3 (1x)
		57455:   211, // $@6 (1x)
		57456:   212, // $@7 (1x)
		57457:   213, // $@8 (1x)
		57462:   214, // ArgumentExpressionList (1x)
		57463:   215, // ArgumentExpressionListOpt (1x)
		57471:   216, // BlockItemList (1x)
		57472:   217, // BlockItemListOpt (1x)
		1048577: 218, // CONSTANT_EXPRESSION (1x)
		57479:   219, // DeclarationList (1x)
		57493:   220, // ElifGroupList (1x)
		57494:   221, // ElifGroupListOpt (1x)
		57495:   222, // ElseGroup (1x)
		57496:   223, // ElseGroupOpt (1x)
		57497:   224, // EndifLine (1x)
		57501:   225, // EnumeratorList (1x)
		57378:   226, // IDENTIFIER_LPAREN (1x)
		1048576: 227, // PREPROCESSING_FILE (1x)
		57536:   228, // PreprocessingFile (1x)
		57541:   229, // Start (1x)
		57546:   230, // StructDeclarationList (1x)
		57548:   231, // StructDeclaratorList (1x)
		1048578: 232, // TRANSLATION_UNIT (1x)
		57552:   233, // TranslationUnit (1x)
		57459:   234, // $default (0x)
		57356:   235, // CAST (0x)
		57345:   236, // error (0x)
		57379:   237, // IDENTIFIER_NONREPL (0x)
		57394:   238, // NOELSE (0x)
		57404:   239, // PPHEADER_NAME (0x)
		57412:   240, // PPNUMBER (0x)
		57414:   241, // PPPASTE (0x)
		57434:   242, // UNARY (0x)
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
		"AUTO",
		"EXTERN",
		"INLINE",
		"NORETURN",
		"REGISTER",
		"TYPEDEF",
		"$end",
		"'='",
		"'{'",
		"Expression",
		"']'",
		"ASM",
		"'.'",
		"STATIC_ASSERT",
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
		"GOTO",
		"PPOTHER",
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
		"CompoundStatement",
		"DeclarationSpecifiers",
		"FunctionSpecifier",
		"PPTokenListOpt",
		"StorageClassSpecifier",
		"AssemblerStatement",
		"ExpressionStatement",
		"IterationStatement",
		"JumpStatement",
		"LabeledStatement",
		"Pointer",
		"SelectionStatement",
		"Statement",
		"PointerOpt",
		"Declarator",
		"StaticAssertDeclaration",
		"ControlLine",
		"GroupPart",
		"IfGroup",
		"IfSection",
		"TextLine",
		"Declaration",
		"$@4",
		"ConstantExpression",
		"DDD",
		"GroupList",
		"AssemblerOperand",
		"AssemblerSymbolicNameOpt",
		"GroupListOpt",
		"ReplacementList",
		"SpecifierQualifierList",
		"TypeQualifierList",
		"$@9",
		"AbstractDeclarator",
		"AssemblerOperands",
		"DeclarationSpecifiersOpt",
		"Designator",
		"Initializer",
		"ParameterDeclaration",
		"TypeName",
		"TypeQualifierListOpt",
		"AssemblerInstructions",
		"CommaOpt",
		"Designation",
		"DesignationOpt",
		"DesignatorList",
		"ExpressionOpt",
		"IdentifierList",
		"InitDeclarator",
		"ParameterList",
		"ParameterTypeList",
		"$@10",
		"$@14",
		"$@16",
		"$@17",
		"$@18",
		"$@5",
		"AbstractDeclaratorOpt",
		"BlockItem",
		"Clobbers",
		"DeclarationListOpt",
		"DeclaratorOpt",
		"DirectAbstractDeclarator",
		"DirectAbstractDeclaratorOpt",
		"DirectDeclarator",
		"ElifGroup",
		"EnumerationConstant",
		"Enumerator",
		"ExternalDeclaration",
		"FunctionBody",
		"FunctionDefinition",
		"IdentifierListOpt",
		"IdentifierOpt",
		"InitDeclaratorList",
		"InitDeclaratorListOpt",
		"InitializerList",
		"ParameterTypeListOpt",
		"SpecifierQualifierListOpt",
		"StructDeclaration",
		"StructDeclarator",
		"VolatileOpt",
		"$@1",
		"$@11",
		"$@12",
		"$@13",
		"$@15",
		"$@2",
		"$@3",
		"$@6",
		"$@7",
		"$@8",
		"ArgumentExpressionList",
		"ArgumentExpressionListOpt",
		"BlockItemList",
		"BlockItemListOpt",
		"CONSTANT_EXPRESSION",
		"DeclarationList",
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
		57426:   "_Static_assert",
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
		57376:   "goto",
		57413:   "ppother",
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
		1:   {204, 0},
		2:   {229, 3},
		3:   {209, 0},
		4:   {229, 3},
		5:   {210, 0},
		6:   {229, 3},
		7:   {189, 1},
		8:   {214, 1},
		9:   {214, 3},
		10:  {215, 0},
		11:  {215, 1},
		12:  {53, 1},
		13:  {53, 1},
		14:  {53, 1},
		15:  {53, 1},
		16:  {53, 1},
		17:  {53, 1},
		18:  {53, 1},
		19:  {53, 3},
		20:  {53, 4},
		21:  {53, 4},
		22:  {53, 3},
		23:  {53, 3},
		24:  {53, 2},
		25:  {53, 2},
		26:  {53, 7},
		27:  {53, 2},
		28:  {53, 2},
		29:  {53, 2},
		30:  {53, 2},
		31:  {53, 2},
		32:  {53, 2},
		33:  {53, 2},
		34:  {53, 2},
		35:  {53, 2},
		36:  {53, 4},
		37:  {53, 4},
		38:  {53, 3},
		39:  {53, 3},
		40:  {53, 3},
		41:  {53, 3},
		42:  {53, 3},
		43:  {53, 3},
		44:  {53, 3},
		45:  {53, 3},
		46:  {53, 3},
		47:  {53, 3},
		48:  {53, 3},
		49:  {53, 3},
		50:  {53, 3},
		51:  {53, 3},
		52:  {53, 3},
		53:  {53, 3},
		54:  {53, 3},
		55:  {53, 3},
		56:  {53, 5},
		57:  {53, 3},
		58:  {53, 3},
		59:  {53, 3},
		60:  {53, 3},
		61:  {53, 3},
		62:  {53, 3},
		63:  {53, 3},
		64:  {53, 3},
		65:  {53, 3},
		66:  {53, 3},
		67:  {53, 3},
		68:  {53, 4},
		69:  {53, 3},
		70:  {169, 0},
		71:  {169, 1},
		72:  {114, 1},
		73:  {114, 3},
		74:  {121, 0},
		75:  {121, 1},
		76:  {145, 0},
		77:  {146, 2},
		78:  {144, 3},
		79:  {144, 1},
		80:  {124, 2},
		81:  {124, 2},
		82:  {124, 2},
		83:  {124, 2},
		84:  {158, 0},
		85:  {158, 1},
		86:  {196, 1},
		87:  {196, 3},
		88:  {197, 0},
		89:  {197, 1},
		90:  {171, 1},
		91:  {179, 0},
		92:  {171, 4},
		93:  {127, 1},
		94:  {127, 1},
		95:  {127, 1},
		96:  {127, 1},
		97:  {127, 1},
		98:  {120, 1},
		99:  {120, 1},
		100: {120, 1},
		101: {120, 1},
		102: {120, 1},
		103: {120, 1},
		104: {120, 1},
		105: {120, 1},
		106: {120, 1},
		107: {120, 1},
		108: {120, 1},
		109: {120, 1},
		110: {120, 1},
		111: {120, 1},
		112: {120, 4},
		113: {120, 4},
		114: {211, 0},
		115: {119, 6},
		116: {119, 2},
		117: {119, 4},
		118: {118, 1},
		119: {118, 1},
		120: {230, 1},
		121: {230, 2},
		122: {201, 3},
		123: {201, 2},
		124: {201, 1},
		125: {153, 2},
		126: {153, 2},
		127: {200, 0},
		128: {200, 1},
		129: {231, 1},
		130: {231, 3},
		131: {202, 1},
		132: {202, 3},
		133: {165, 0},
		134: {165, 1},
		135: {212, 0},
		136: {117, 7},
		137: {117, 2},
		138: {225, 1},
		139: {225, 3},
		140: {190, 1},
		141: {190, 3},
		142: {113, 1},
		143: {113, 1},
		144: {113, 1},
		145: {125, 1},
		146: {125, 1},
		147: {137, 2},
		148: {184, 0},
		149: {184, 1},
		150: {187, 1},
		151: {187, 3},
		152: {187, 5},
		153: {187, 6},
		154: {187, 6},
		155: {187, 5},
		156: {213, 0},
		157: {187, 5},
		158: {187, 4},
		159: {133, 2},
		160: {133, 3},
		161: {136, 0},
		162: {136, 1},
		163: {154, 1},
		164: {154, 2},
		165: {163, 0},
		166: {163, 1},
		167: {173, 1},
		168: {173, 3},
		169: {199, 0},
		170: {199, 1},
		171: {172, 1},
		172: {172, 3},
		173: {161, 2},
		174: {161, 2},
		175: {170, 1},
		176: {170, 3},
		177: {194, 0},
		178: {194, 1},
		179: {195, 0},
		180: {195, 1},
		181: {155, 0},
		182: {162, 3},
		183: {156, 1},
		184: {156, 2},
		185: {180, 0},
		186: {180, 1},
		187: {185, 3},
		188: {185, 4},
		189: {185, 5},
		190: {185, 6},
		191: {185, 6},
		192: {185, 4},
		193: {174, 0},
		194: {185, 4},
		195: {205, 0},
		196: {185, 5},
		197: {186, 0},
		198: {186, 1},
		199: {160, 1},
		200: {160, 4},
		201: {160, 3},
		202: {198, 2},
		203: {198, 4},
		204: {198, 0},
		205: {166, 2},
		206: {167, 0},
		207: {167, 1},
		208: {168, 1},
		209: {168, 2},
		210: {159, 3},
		211: {159, 2},
		212: {135, 1},
		213: {135, 1},
		214: {135, 1},
		215: {135, 1},
		216: {135, 1},
		217: {135, 1},
		218: {135, 1},
		219: {132, 3},
		220: {132, 4},
		221: {132, 3},
		222: {206, 0},
		223: {123, 4},
		224: {216, 1},
		225: {216, 2},
		226: {217, 0},
		227: {217, 1},
		228: {181, 1},
		229: {181, 1},
		230: {129, 2},
		231: {134, 5},
		232: {134, 7},
		233: {134, 5},
		234: {130, 5},
		235: {130, 7},
		236: {130, 9},
		237: {130, 8},
		238: {131, 3},
		239: {131, 2},
		240: {131, 2},
		241: {131, 3},
		242: {233, 1},
		243: {233, 2},
		244: {191, 1},
		245: {191, 1},
		246: {191, 2},
		247: {191, 1},
		248: {207, 0},
		249: {193, 5},
		250: {175, 0},
		251: {208, 0},
		252: {193, 5},
		253: {176, 0},
		254: {192, 2},
		255: {177, 0},
		256: {192, 3},
		257: {219, 1},
		258: {219, 2},
		259: {183, 0},
		260: {178, 0},
		261: {183, 2},
		262: {164, 1},
		263: {164, 2},
		264: {122, 5},
		265: {203, 0},
		266: {203, 1},
		267: {149, 5},
		268: {157, 1},
		269: {157, 3},
		270: {150, 0},
		271: {150, 3},
		272: {182, 1},
		273: {182, 3},
		274: {128, 1},
		275: {128, 7},
		276: {128, 9},
		277: {128, 11},
		278: {128, 13},
		279: {128, 6},
		280: {128, 8},
		281: {138, 7},
		282: {228, 1},
		283: {148, 1},
		284: {148, 2},
		285: {151, 0},
		286: {151, 1},
		287: {140, 1},
		288: {140, 1},
		289: {140, 3},
		290: {140, 1},
		291: {142, 4},
		292: {141, 4},
		293: {141, 4},
		294: {141, 4},
		295: {220, 1},
		296: {220, 2},
		297: {221, 0},
		298: {221, 1},
		299: {188, 4},
		300: {222, 3},
		301: {223, 0},
		302: {223, 1},
		303: {224, 1},
		304: {139, 3},
		305: {139, 5},
		306: {139, 7},
		307: {139, 5},
		308: {139, 2},
		309: {139, 1},
		310: {139, 3},
		311: {139, 3},
		312: {139, 2},
		313: {139, 3},
		314: {139, 6},
		315: {139, 2},
		316: {139, 4},
		317: {139, 3},
		318: {143, 1},
		319: {152, 1},
		320: {115, 1},
		321: {126, 1},
		322: {126, 2},
		323: {116, 1},
		324: {116, 2},
	}

	yyXErrors = map[yyXError]string{
		{0, 50}:   "invalid empty input",
		{577, -1}: "expected #endif",
		{579, -1}: "expected #endif",
		{1, -1}:   "expected $end",
		{497, -1}: "expected $end",
		{499, -1}: "expected $end",
		{32, -1}:  "expected '('",
		{49, -1}:  "expected '('",
		{75, -1}:  "expected '('",
		{270, -1}: "expected '('",
		{271, -1}: "expected '('",
		{272, -1}: "expected '('",
		{274, -1}: "expected '('",
		{284, -1}: "expected '('",
		{307, -1}: "expected '('",
		{347, -1}: "expected '('",
		{429, -1}: "expected '('",
		{54, -1}:  "expected ')'",
		{77, -1}:  "expected ')'",
		{84, -1}:  "expected ')'",
		{176, -1}: "expected ')'",
		{191, -1}: "expected ')'",
		{194, -1}: "expected ')'",
		{197, -1}: "expected ')'",
		{205, -1}: "expected ')'",
		{210, -1}: "expected ')'",
		{216, -1}: "expected ')'",
		{232, -1}: "expected ')'",
		{237, -1}: "expected ')'",
		{248, -1}: "expected ')'",
		{249, -1}: "expected ')'",
		{337, -1}: "expected ')'",
		{343, -1}: "expected ')'",
		{427, -1}: "expected ')'",
		{483, -1}: "expected ')'",
		{541, -1}: "expected ')'",
		{542, -1}: "expected ')'",
		{549, -1}: "expected ')'",
		{552, -1}: "expected ')'",
		{52, -1}:  "expected ','",
		{263, -1}: "expected ':'",
		{289, -1}: "expected ':'",
		{371, -1}: "expected ':'",
		{473, -1}: "expected ':'",
		{45, -1}:  "expected ';'",
		{55, -1}:  "expected ';'",
		{269, -1}: "expected ';'",
		{276, -1}: "expected ';'",
		{277, -1}: "expected ';'",
		{326, -1}: "expected ';'",
		{330, -1}: "expected ';'",
		{333, -1}: "expected ';'",
		{335, -1}: "expected ';'",
		{341, -1}: "expected ';'",
		{350, -1}: "expected ';'",
		{374, -1}: "expected ';'",
		{442, -1}: "expected ';'",
		{381, -1}: "expected '='",
		{89, -1}:  "expected '['",
		{521, -1}: "expected '\\n'",
		{525, -1}: "expected '\\n'",
		{529, -1}: "expected '\\n'",
		{532, -1}: "expected '\\n'",
		{534, -1}: "expected '\\n'",
		{556, -1}: "expected '\\n'",
		{561, -1}: "expected '\\n'",
		{564, -1}: "expected '\\n'",
		{571, -1}: "expected '\\n'",
		{576, -1}: "expected '\\n'",
		{582, -1}: "expected '\\n'",
		{95, -1}:  "expected ']'",
		{184, -1}: "expected ']'",
		{228, -1}: "expected ']'",
		{295, -1}: "expected ']'",
		{395, -1}: "expected ']'",
		{446, -1}: "expected '{'",
		{448, -1}: "expected '{'",
		{460, -1}: "expected '{'",
		{264, -1}: "expected '}'",
		{401, -1}: "expected '}'",
		{417, -1}: "expected '}'",
		{457, -1}: "expected '}'",
		{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		{204, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{88, -1}:  "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{283, -1}: "expected assembler instructions or string literal",
		{285, -1}: "expected assembler instructions or string literal",
		{430, -1}: "expected assembler instructions or string literal",
		{297, -1}: "expected assembler operand or one of ['[', string literal]",
		{313, -1}: "expected assembler operands or one of [')', ':', '[', string literal]",
		{290, -1}: "expected assembler operands or one of ['[', string literal]",
		{316, -1}: "expected assembler operands or one of ['[', string literal]",
		{320, -1}: "expected assembler operands or one of ['[', string literal]",
		{441, -1}: "expected assembler statement or asm",
		{266, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{298, -1}: "expected clobbers or string literal",
		{323, -1}: "expected clobbers or string literal",
		{440, -1}: "expected compound statement or '{'",
		{64, -1}:  "expected compound statement or expression list or type name or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{245, -1}: "expected compound statement or expression list or type name or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{50, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{262, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{392, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{454, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{474, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{496, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{434, -1}: "expected declaration list or one of [_Bool, _Complex, _Noreturn, _Static_assert, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{436, -1}: "expected declaration or one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{332, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{47, -1}:  "expected declarator or one of ['(', '*', identifier]",
		{380, -1}: "expected declarator or one of ['(', '*', identifier]",
		{196, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		{389, -1}: "expected designator or one of ['.', '=', '[']",
		{199, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		{85, -1}:  "expected direct abstract declarator or one of ['(', '[']",
		{378, -1}: "expected direct declarator or one of ['(', identifier]",
		{569, -1}: "expected elif group or one of [#elif, #else, #endif]",
		{575, -1}: "expected endif line or #endif",
		{506, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		{567, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		{449, -1}: "expected enumerator list or identifier",
		{456, -1}: "expected enumerator or one of ['}', identifier]",
		{100, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{124, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{348, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{352, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{356, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{360, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{413, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		{92, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{227, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{428, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{51, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{66, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{67, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{68, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{69, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{70, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{71, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{72, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{73, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{74, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{98, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{106, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{107, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{108, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{109, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{110, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{111, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{112, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{113, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{114, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{115, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{116, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{117, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{118, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{119, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{120, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{121, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{122, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{123, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{125, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{126, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{127, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{128, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{129, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{130, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{131, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{132, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{133, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{134, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{135, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{150, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{151, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{178, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{185, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{221, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{224, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{308, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{96, -1}:  "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		{219, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		{481, -1}: "expected expression or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{6, -1}:   "expected external declaration or one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{433, -1}: "expected function body or one of ['{', asm]",
		{438, -1}: "expected function body or one of ['{', asm]",
		{492, -1}: "expected function body or one of ['{', asm]",
		{493, -1}: "expected function body or one of ['{', asm]",
		{491, -1}: "expected function body or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{432, -1}: "expected function body or optional declaration list or one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{558, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{500, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{102, -1}: "expected identifier",
		{103, -1}: "expected identifier",
		{213, -1}: "expected identifier",
		{275, -1}: "expected identifier",
		{294, -1}: "expected identifier",
		{393, -1}: "expected identifier",
		{508, -1}: "expected identifier",
		{509, -1}: "expected identifier",
		{516, -1}: "expected identifier",
		{302, -1}: "expected identifier list or identifier",
		{538, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		{407, -1}: "expected init declarator or one of ['(', '*', identifier]",
		{386, -1}: "expected initializer list or optional comma or one of ['!', '&', '(', '*', '+', ',', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{415, -1}: "expected initializer list or optional comma or one of ['!', '&', '(', '*', '+', ',', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{382, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{388, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{403, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{405, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{400, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{57, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{58, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{59, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{60, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{61, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{62, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{63, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{104, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{105, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{137, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{138, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{139, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{140, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{141, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{142, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{143, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{144, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{145, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{146, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{147, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{153, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{154, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{155, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{156, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{157, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{158, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{159, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{160, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{161, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{162, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{163, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{164, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{165, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{166, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{167, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{168, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{169, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{170, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{171, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{172, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{173, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{177, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{181, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{189, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{244, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{246, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{412, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{414, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{418, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{419, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{420, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{421, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{422, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{423, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{424, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{425, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{426, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{65, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{148, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{152, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{174, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{179, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{309, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{482, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{383, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{252, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{384, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{91, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{99, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{186, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{222, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{225, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{501, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{502, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{503, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{505, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{512, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{518, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{520, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{523, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{526, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{528, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{530, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{531, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{533, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{535, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{536, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{539, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{544, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{545, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{547, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{551, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{554, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{555, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{560, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{580, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{581, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{583, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{559, -1}: "expected one of [#elif, #else, #endif]",
		{563, -1}: "expected one of [#elif, #else, #endif]",
		{566, -1}: "expected one of [#elif, #else, #endif]",
		{568, -1}: "expected one of [#elif, #else, #endif]",
		{573, -1}: "expected one of [#elif, #else, #endif]",
		{574, -1}: "expected one of [#elif, #else, #endif]",
		{368, -1}: "expected one of [$end, '!', '&', '(', ')', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{8, -1}:   "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{56, -1}:  "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{409, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{42, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{43, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{44, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{46, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{439, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{443, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{444, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{445, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{494, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{495, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{38, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{39, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{93, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{182, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{255, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{256, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{257, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{258, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{259, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{260, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{261, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{280, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{304, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{312, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{315, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{318, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{319, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{322, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{325, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{327, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{328, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{329, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{331, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{339, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{345, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{351, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{355, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{359, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{363, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{365, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{366, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{370, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{373, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{411, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{265, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{267, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{268, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{367, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{390, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{397, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{447, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{461, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{18, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{19, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{20, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{21, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{22, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{23, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{24, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{25, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{26, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{27, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{28, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{29, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{30, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{31, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{458, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{464, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{479, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{484, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{485, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{17, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{40, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{41, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{486, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{487, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{488, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{489, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{490, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{241, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		{242, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		{243, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		{202, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{203, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{206, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{215, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{217, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{223, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{226, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{229, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{230, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{83, -1}:  "expected one of ['(', ')', ',', '[', identifier]",
		{240, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		{87, -1}:  "expected one of ['(', ')', ',', '[']",
		{136, -1}: "expected one of ['(', ')', ',', '[']",
		{183, -1}: "expected one of ['(', ')', ',', '[']",
		{187, -1}: "expected one of ['(', ')', ',', '[']",
		{188, -1}: "expected one of ['(', ')', ',', '[']",
		{190, -1}: "expected one of ['(', ')', ',', '[']",
		{198, -1}: "expected one of ['(', ')', ',', '[']",
		{234, -1}: "expected one of ['(', ')', ',', '[']",
		{238, -1}: "expected one of ['(', ')', ',', '[']",
		{281, -1}: "expected one of ['(', goto]",
		{282, -1}: "expected one of ['(', goto]",
		{379, -1}: "expected one of ['(', identifier]",
		{292, -1}: "expected one of [')', ',', ':']",
		{299, -1}: "expected one of [')', ',', ':']",
		{305, -1}: "expected one of [')', ',', ':']",
		{306, -1}: "expected one of [')', ',', ':']",
		{310, -1}: "expected one of [')', ',', ':']",
		{314, -1}: "expected one of [')', ',', ':']",
		{321, -1}: "expected one of [')', ',', ':']",
		{253, -1}: "expected one of [')', ',', ';']",
		{211, -1}: "expected one of [')', ',', ...]",
		{214, -1}: "expected one of [')', ',', ...]",
		{540, -1}: "expected one of [')', ',', ...]",
		{86, -1}:  "expected one of [')', ',']",
		{175, -1}: "expected one of [')', ',']",
		{193, -1}: "expected one of [')', ',']",
		{195, -1}: "expected one of [')', ',']",
		{200, -1}: "expected one of [')', ',']",
		{201, -1}: "expected one of [')', ',']",
		{212, -1}: "expected one of [')', ',']",
		{233, -1}: "expected one of [')', ',']",
		{247, -1}: "expected one of [')', ',']",
		{303, -1}: "expected one of [')', ',']",
		{317, -1}: "expected one of [')', ',']",
		{324, -1}: "expected one of [')', ',']",
		{349, -1}: "expected one of [')', ',']",
		{353, -1}: "expected one of [')', ',']",
		{357, -1}: "expected one of [')', ',']",
		{361, -1}: "expected one of [')', ',']",
		{286, -1}: "expected one of [')', ':', string literal]",
		{288, -1}: "expected one of [')', ':', string literal]",
		{311, -1}: "expected one of [')', ':', string literal]",
		{431, -1}: "expected one of [')', string literal]",
		{472, -1}: "expected one of [',', ':', ';']",
		{149, -1}: "expected one of [',', ':']",
		{293, -1}: "expected one of [',', ':']",
		{300, -1}: "expected one of [',', ':']",
		{377, -1}: "expected one of [',', ';', '=']",
		{402, -1}: "expected one of [',', ';', '}']",
		{406, -1}: "expected one of [',', ';', '}']",
		{375, -1}: "expected one of [',', ';']",
		{376, -1}: "expected one of [',', ';']",
		{385, -1}: "expected one of [',', ';']",
		{408, -1}: "expected one of [',', ';']",
		{469, -1}: "expected one of [',', ';']",
		{471, -1}: "expected one of [',', ';']",
		{475, -1}: "expected one of [',', ';']",
		{478, -1}: "expected one of [',', ';']",
		{450, -1}: "expected one of [',', '=', '}']",
		{453, -1}: "expected one of [',', '=', '}']",
		{180, -1}: "expected one of [',', ']']",
		{399, -1}: "expected one of [',', '}']",
		{404, -1}: "expected one of [',', '}']",
		{452, -1}: "expected one of [',', '}']",
		{455, -1}: "expected one of [',', '}']",
		{459, -1}: "expected one of [',', '}']",
		{391, -1}: "expected one of ['.', '=', '[']",
		{394, -1}: "expected one of ['.', '=', '[']",
		{396, -1}: "expected one of ['.', '=', '[']",
		{398, -1}: "expected one of ['.', '=', '[']",
		{287, -1}: "expected one of [':', string literal]",
		{510, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		{519, -1}: "expected one of ['\\n', ppother]",
		{522, -1}: "expected one of ['\\n', ppother]",
		{524, -1}: "expected one of ['\\n', ppother]",
		{435, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{437, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{34, -1}:  "expected one of ['{', identifier]",
		{35, -1}:  "expected one of ['{', identifier]",
		{466, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{468, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{470, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{476, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{480, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{548, -1}: "expected one of [..., identifier]",
		{81, -1}:  "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		{101, -1}: "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{250, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{251, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{387, -1}: "expected optional comma or one of [',', '}']",
		{416, -1}: "expected optional comma or one of [',', '}']",
		{451, -1}: "expected optional comma or one of [',', '}']",
		{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{12, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{336, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{342, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{278, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{334, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{340, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{218, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{207, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{90, -1}:  "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{94, -1}:  "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{557, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{562, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{565, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{572, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{578, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{208, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{33, -1}:  "expected optional identifier or one of ['{', identifier]",
		{36, -1}:  "expected optional identifier or one of ['{', identifier]",
		{254, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		{192, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{235, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{236, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{79, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{80, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{511, -1}: "expected optional token list or one of ['\\n', ppother]",
		{515, -1}: "expected optional token list or one of ['\\n', ppother]",
		{82, -1}:  "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		{279, -1}: "expected optional volatile or one of ['(', goto, volatile]",
		{48, -1}:  "expected optional volatile or one of ['(', volatile]",
		{231, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{209, -1}: "expected parameter type list or one of [_Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{239, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{498, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{537, -1}: "expected replacement list or one of ['\\n', ppother]",
		{543, -1}: "expected replacement list or one of ['\\n', ppother]",
		{546, -1}: "expected replacement list or one of ['\\n', ppother]",
		{550, -1}: "expected replacement list or one of ['\\n', ppother]",
		{553, -1}: "expected replacement list or one of ['\\n', ppother]",
		{78, -1}:  "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{273, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{338, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{344, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{354, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{358, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{362, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{364, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{369, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{372, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{410, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{53, -1}:  "expected string literal",
		{291, -1}: "expected string literal",
		{296, -1}: "expected string literal",
		{301, -1}: "expected string literal",
		{462, -1}: "expected struct declaration list or one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{463, -1}: "expected struct declaration list or one of [_Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{465, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{467, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		{477, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		{527, -1}: "expected token list or one of ['\\n', ppother]",
		{504, -1}: "expected token list or ppother",
		{507, -1}: "expected token list or ppother",
		{513, -1}: "expected token list or ppother",
		{514, -1}: "expected token list or ppother",
		{517, -1}: "expected token list or ppother",
		{570, -1}: "expected token list or ppother",
		{4, -1}:   "expected translation unit or one of ['(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{5, -1}:   "expected translation unit or one of ['(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{76, -1}:  "expected type name or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{97, -1}:  "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		{220, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{346, -1}: "expected while",
		{3, 50}:   "unexpected EOF",
		{2, 50}:   "unexpected EOF",
		{4, 50}:   "unexpected EOF",
	}

	yyParseTab = [584][]uint16{
		// 0
		{218: 328, 227: 327, 229: 326, 232: 329},
		{50: 325},
		{84: 324, 86: 324, 100: 324, 324, 324, 324, 324, 324, 324, 324, 324, 324, 324, 324, 204: 823},
		{322, 322, 322, 322, 322, 322, 322, 322, 11: 322, 13: 322, 322, 322, 322, 322, 322, 322, 322, 322, 209: 821},
		{320, 320, 320, 8: 320, 22: 320, 320, 320, 26: 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 43: 320, 320, 320, 320, 320, 320, 320, 55: 320, 57: 320, 210: 330},
		// 5
		{75, 75, 75, 8: 371, 22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 55: 373, 57: 374, 113: 336, 117: 355, 358, 354, 335, 122: 370, 124: 332, 337, 127: 334, 138: 333, 144: 369, 175: 372, 191: 367, 193: 368, 233: 331},
		{75, 75, 75, 8: 371, 22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 319, 55: 373, 57: 374, 113: 336, 117: 355, 358, 354, 335, 122: 370, 124: 332, 337, 127: 334, 138: 333, 144: 369, 175: 372, 191: 820, 193: 368},
		{164, 407, 164, 8: 237, 133: 704, 136: 703, 816, 171: 700, 196: 701, 699},
		{246, 246, 246, 246, 246, 246, 246, 246, 246, 11: 246, 13: 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 43: 246, 246, 246, 246, 246, 246, 246, 246, 52: 246, 55: 246, 57: 246, 85: 246, 87: 246, 246, 246, 246, 246, 246, 246, 246, 246, 246},
		{241, 241, 241, 8: 241, 241, 241, 12: 241, 22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 113: 336, 117: 355, 358, 354, 335, 124: 812, 337, 127: 334, 158: 815},
		// 10
		{241, 241, 241, 8: 241, 241, 241, 12: 241, 22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 113: 336, 117: 355, 358, 354, 335, 124: 812, 337, 127: 334, 158: 814},
		{241, 241, 241, 8: 241, 241, 241, 12: 241, 22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 113: 336, 117: 355, 358, 354, 335, 124: 812, 337, 127: 334, 158: 813},
		{241, 241, 241, 8: 241, 241, 241, 12: 241, 22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 113: 336, 117: 355, 358, 354, 335, 124: 812, 337, 127: 334, 158: 811},
		{232, 232, 232, 8: 232, 232, 232, 12: 232, 22: 232, 232, 232, 26: 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 43: 232, 232, 232, 232, 232, 232, 232},
		{231, 231, 231, 8: 231, 231, 231, 12: 231, 22: 231, 231, 231, 26: 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 43: 231, 231, 231, 231, 231, 231, 231},
		// 15
		{230, 230, 230, 8: 230, 230, 230, 12: 230, 22: 230, 230, 230, 26: 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 43: 230, 230, 230, 230, 230, 230, 230},
		{229, 229, 229, 8: 229, 229, 229, 12: 229, 22: 229, 229, 229, 26: 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 43: 229, 229, 229, 229, 229, 229, 229},
		{228, 228, 228, 8: 228, 228, 228, 12: 228, 22: 228, 228, 228, 26: 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 43: 228, 228, 228, 228, 228, 228, 228},
		{227, 227, 227, 8: 227, 227, 227, 12: 227, 22: 227, 227, 227, 26: 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227},
		{226, 226, 226, 8: 226, 226, 226, 12: 226, 22: 226, 226, 226, 26: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226},
		// 20
		{225, 225, 225, 8: 225, 225, 225, 12: 225, 22: 225, 225, 225, 26: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225},
		{224, 224, 224, 8: 224, 224, 224, 12: 224, 22: 224, 224, 224, 26: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224},
		{223, 223, 223, 8: 223, 223, 223, 12: 223, 22: 223, 223, 223, 26: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223},
		{222, 222, 222, 8: 222, 222, 222, 12: 222, 22: 222, 222, 222, 26: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222},
		{221, 221, 221, 8: 221, 221, 221, 12: 221, 22: 221, 221, 221, 26: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221},
		// 25
		{220, 220, 220, 8: 220, 220, 220, 12: 220, 22: 220, 220, 220, 26: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220},
		{219, 219, 219, 8: 219, 219, 219, 12: 219, 22: 219, 219, 219, 26: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219},
		{218, 218, 218, 8: 218, 218, 218, 12: 218, 22: 218, 218, 218, 26: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218},
		{217, 217, 217, 8: 217, 217, 217, 12: 217, 22: 217, 217, 217, 26: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217},
		{216, 216, 216, 8: 216, 216, 216, 12: 216, 22: 216, 216, 216, 26: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216},
		// 30
		{215, 215, 215, 8: 215, 215, 215, 12: 215, 22: 215, 215, 215, 26: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215},
		{214, 214, 214, 8: 214, 214, 214, 12: 214, 22: 214, 214, 214, 26: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214},
		{806},
		{2: 786, 52: 146, 195: 785},
		{2: 207, 52: 207},
		// 35
		{2: 206, 52: 206},
		{2: 772, 52: 146, 195: 771},
		{183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 26: 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 54: 183},
		{182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 26: 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 54: 182},
		{181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 26: 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 54: 181},
		// 40
		{180, 180, 180, 8: 180, 180, 180, 12: 180, 22: 180, 180, 180, 26: 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 43: 180, 180, 180, 180, 180, 180, 180},
		{179, 179, 179, 8: 179, 179, 179, 12: 179, 22: 179, 179, 179, 26: 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 43: 179, 179, 179, 179, 179, 179, 179},
		{83, 83, 83, 8: 83, 22: 83, 83, 83, 26: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 43: 83, 83, 83, 83, 83, 83, 83, 83, 55: 83, 57: 83},
		{81, 81, 81, 8: 81, 22: 81, 81, 81, 26: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 43: 81, 81, 81, 81, 81, 81, 81, 81, 55: 81, 57: 81},
		{80, 80, 80, 8: 80, 22: 80, 80, 80, 26: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 43: 80, 80, 80, 80, 80, 80, 80, 80, 55: 80, 57: 80},
		// 45
		{8: 770},
		{78, 78, 78, 8: 78, 22: 78, 78, 78, 26: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 43: 78, 78, 78, 78, 78, 78, 78, 78, 55: 78, 57: 78},
		{164, 407, 164, 133: 704, 136: 703, 757},
		{60, 22: 607, 203: 754},
		{375},
		// 50
		{249, 249, 249, 249, 249, 249, 249, 249, 11: 249, 13: 249, 249, 249, 249, 249, 249, 249, 249, 249, 145: 376, 377},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 390},
		{10: 378},
		{11: 379},
		{9: 380},
		// 55
		{8: 381},
		{44, 44, 44, 44, 44, 44, 44, 44, 44, 11: 44, 13: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 43: 44, 44, 44, 44, 44, 44, 44, 44, 52: 44, 55: 44, 57: 44, 85: 44, 87: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44},
		{313, 313, 3: 313, 313, 313, 313, 313, 313, 313, 313, 12: 313, 25: 313, 42: 313, 50: 313, 313, 54: 313, 56: 313, 58: 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313},
		{312, 312, 3: 312, 312, 312, 312, 312, 312, 312, 312, 12: 312, 25: 312, 42: 312, 50: 312, 312, 54: 312, 56: 312, 58: 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312},
		{311, 311, 3: 311, 311, 311, 311, 311, 311, 311, 311, 12: 311, 25: 311, 42: 311, 50: 311, 311, 54: 311, 56: 311, 58: 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311},
		// 60
		{310, 310, 3: 310, 310, 310, 310, 310, 310, 310, 310, 12: 310, 25: 310, 42: 310, 50: 310, 310, 54: 310, 56: 310, 58: 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310},
		{309, 309, 3: 309, 309, 309, 309, 309, 309, 309, 309, 12: 309, 25: 309, 42: 309, 50: 309, 309, 54: 309, 56: 309, 58: 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309},
		{308, 308, 3: 308, 308, 308, 308, 308, 308, 308, 308, 12: 308, 25: 308, 42: 308, 50: 308, 308, 54: 308, 56: 308, 58: 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308},
		{307, 307, 3: 307, 307, 307, 307, 307, 307, 307, 307, 12: 307, 25: 307, 42: 307, 50: 307, 307, 54: 307, 56: 307, 58: 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 144, 144, 144, 26: 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 52: 575, 473, 114: 572, 123: 574, 155: 403, 162: 752},
		// 65
		{426, 431, 3: 444, 434, 435, 430, 429, 248, 10: 248, 12: 425, 25: 248, 42: 248, 50: 248, 450, 54: 248, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 751},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 750},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 749},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 514},
		// 70
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 748},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 747},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 746},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 745},
		{570, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 571},
		// 75
		{401},
		{22: 144, 144, 144, 26: 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 155: 403, 162: 402},
		{9: 569},
		{22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 113: 405, 117: 355, 358, 354, 404, 153: 406},
		{198, 198, 198, 8: 198, 198, 12: 198, 22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 198, 113: 405, 117: 355, 358, 354, 404, 153: 567, 200: 568},
		// 80
		{198, 198, 198, 8: 198, 198, 12: 198, 22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 198, 113: 405, 117: 355, 358, 354, 404, 153: 567, 200: 566},
		{164, 407, 9: 140, 12: 164, 133: 408, 136: 410, 156: 411, 180: 409},
		{160, 160, 160, 9: 160, 160, 12: 160, 22: 364, 362, 363, 113: 418, 154: 422, 163: 564},
		{163, 2: 163, 9: 142, 142, 12: 163},
		{9: 143},
		// 85
		{413, 12: 128, 185: 412, 414},
		{9: 139, 139},
		{560, 9: 141, 141, 12: 127},
		{164, 407, 9: 132, 12: 164, 22: 132, 132, 132, 26: 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 43: 132, 132, 132, 132, 132, 132, 132, 133: 408, 136: 410, 156: 516, 174: 517},
		{12: 415},
		// 90
		{389, 417, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 364, 362, 363, 43: 421, 53: 416, 255, 113: 418, 154: 419, 169: 420},
		{426, 431, 3: 444, 434, 435, 430, 429, 12: 425, 51: 450, 54: 254, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 514, 515},
		{162, 162, 162, 162, 162, 162, 162, 162, 9: 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 43: 162, 54: 162},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 364, 362, 363, 43: 510, 53: 416, 255, 113: 507, 169: 509},
		// 95
		{54: 508},
		{160, 160, 160, 160, 160, 160, 160, 160, 11: 160, 13: 160, 160, 160, 160, 160, 160, 160, 160, 160, 364, 362, 363, 113: 418, 154: 422, 163: 423},
		{159, 159, 159, 159, 159, 159, 159, 159, 9: 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 364, 362, 363, 113: 507},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 424},
		{426, 431, 3: 444, 434, 435, 430, 429, 12: 425, 51: 450, 54: 461, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		// 100
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 473, 114: 505},
		{389, 394, 382, 393, 395, 396, 392, 391, 9: 315, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 499, 214: 500, 501},
		{2: 498},
		{2: 497},
		{301, 301, 3: 301, 301, 301, 301, 301, 301, 301, 301, 12: 301, 25: 301, 42: 301, 50: 301, 301, 54: 301, 56: 301, 58: 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301},
		// 105
		{300, 300, 3: 300, 300, 300, 300, 300, 300, 300, 300, 12: 300, 25: 300, 42: 300, 50: 300, 300, 54: 300, 56: 300, 58: 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 496},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 495},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 494},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 493},
		// 110
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 492},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 491},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 490},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 489},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 488},
		// 115
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 487},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 486},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 485},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 484},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 483},
		// 120
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 482},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 481},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 480},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 479},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 473, 114: 474},
		// 125
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 472},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 471},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 470},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 469},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 468},
		// 130
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 467},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 466},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 465},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 464},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 463},
		// 135
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 462},
		{135, 9: 135, 135, 12: 135},
		{426, 431, 3: 444, 434, 435, 430, 429, 258, 258, 258, 12: 425, 25: 258, 42: 258, 50: 258, 450, 54: 258, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{426, 431, 3: 444, 434, 435, 430, 429, 259, 259, 259, 12: 425, 25: 259, 42: 259, 50: 259, 450, 54: 259, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{426, 431, 3: 444, 434, 435, 430, 429, 260, 260, 260, 12: 425, 25: 260, 42: 260, 50: 260, 450, 54: 260, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		// 140
		{426, 431, 3: 444, 434, 435, 430, 429, 261, 261, 261, 12: 425, 25: 261, 42: 261, 50: 261, 450, 54: 261, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{426, 431, 3: 444, 434, 435, 430, 429, 262, 262, 262, 12: 425, 25: 262, 42: 262, 50: 262, 450, 54: 262, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{426, 431, 3: 444, 434, 435, 430, 429, 263, 263, 263, 12: 425, 25: 263, 42: 263, 50: 263, 450, 54: 263, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{426, 431, 3: 444, 434, 435, 430, 429, 264, 264, 264, 12: 425, 25: 264, 42: 264, 50: 264, 450, 54: 264, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{426, 431, 3: 444, 434, 435, 430, 429, 265, 265, 265, 12: 425, 25: 265, 42: 265, 50: 265, 450, 54: 265, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		// 145
		{426, 431, 3: 444, 434, 435, 430, 429, 266, 266, 266, 12: 425, 25: 266, 42: 266, 50: 266, 450, 54: 266, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{426, 431, 3: 444, 434, 435, 430, 429, 267, 267, 267, 12: 425, 25: 267, 42: 267, 50: 267, 450, 54: 267, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{426, 431, 3: 444, 434, 435, 430, 429, 268, 268, 268, 12: 425, 25: 268, 42: 268, 50: 268, 450, 54: 268, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{426, 431, 3: 444, 434, 435, 430, 429, 253, 253, 253, 12: 425, 42: 253, 51: 450, 54: 253, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{10: 476, 42: 475},
		// 150
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 478},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 477},
		{426, 431, 3: 444, 434, 435, 430, 429, 252, 252, 252, 12: 425, 42: 252, 51: 450, 54: 252, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{426, 431, 3: 444, 434, 435, 430, 429, 269, 269, 269, 12: 425, 25: 269, 42: 269, 50: 269, 269, 54: 269, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 269, 447, 269, 428, 269, 442, 441, 440, 436, 269, 269, 269, 443, 269, 448, 437, 269, 269, 269},
		{426, 431, 3: 444, 434, 435, 430, 429, 270, 270, 270, 12: 425, 25: 270, 42: 270, 50: 270, 270, 54: 270, 56: 427, 58: 433, 432, 438, 439, 270, 445, 446, 270, 447, 270, 428, 270, 442, 441, 440, 436, 270, 270, 270, 443, 270, 270, 437, 270, 270, 270},
		// 155
		{426, 431, 3: 444, 434, 435, 430, 429, 271, 271, 271, 12: 425, 25: 271, 42: 271, 50: 271, 271, 54: 271, 56: 427, 58: 433, 432, 438, 439, 271, 445, 446, 271, 271, 271, 428, 271, 442, 441, 440, 436, 271, 271, 271, 443, 271, 271, 437, 271, 271, 271},
		{426, 431, 3: 444, 434, 435, 430, 429, 272, 272, 272, 12: 425, 25: 272, 42: 272, 50: 272, 272, 54: 272, 56: 427, 58: 433, 432, 438, 439, 272, 445, 272, 272, 272, 272, 428, 272, 442, 441, 440, 436, 272, 272, 272, 443, 272, 272, 437, 272, 272, 272},
		{426, 431, 3: 444, 434, 435, 430, 429, 273, 273, 273, 12: 425, 25: 273, 42: 273, 50: 273, 273, 54: 273, 56: 427, 58: 433, 432, 438, 439, 273, 273, 273, 273, 273, 273, 428, 273, 442, 441, 440, 436, 273, 273, 273, 443, 273, 273, 437, 273, 273, 273},
		{426, 431, 3: 274, 434, 435, 430, 429, 274, 274, 274, 12: 425, 25: 274, 42: 274, 50: 274, 274, 54: 274, 56: 427, 58: 433, 432, 438, 439, 274, 274, 274, 274, 274, 274, 428, 274, 442, 441, 440, 436, 274, 274, 274, 443, 274, 274, 437, 274, 274, 274},
		{426, 431, 3: 275, 434, 435, 430, 429, 275, 275, 275, 12: 425, 25: 275, 42: 275, 50: 275, 275, 54: 275, 56: 427, 58: 433, 432, 438, 439, 275, 275, 275, 275, 275, 275, 428, 275, 275, 441, 440, 436, 275, 275, 275, 275, 275, 275, 437, 275, 275, 275},
		// 160
		{426, 431, 3: 276, 434, 435, 430, 429, 276, 276, 276, 12: 425, 25: 276, 42: 276, 50: 276, 276, 54: 276, 56: 427, 58: 433, 432, 438, 439, 276, 276, 276, 276, 276, 276, 428, 276, 276, 441, 440, 436, 276, 276, 276, 276, 276, 276, 437, 276, 276, 276},
		{426, 431, 3: 277, 434, 435, 430, 429, 277, 277, 277, 12: 425, 25: 277, 42: 277, 50: 277, 277, 54: 277, 56: 427, 58: 433, 432, 277, 277, 277, 277, 277, 277, 277, 277, 428, 277, 277, 277, 277, 436, 277, 277, 277, 277, 277, 277, 437, 277, 277, 277},
		{426, 431, 3: 278, 434, 435, 430, 429, 278, 278, 278, 12: 425, 25: 278, 42: 278, 50: 278, 278, 54: 278, 56: 427, 58: 433, 432, 278, 278, 278, 278, 278, 278, 278, 278, 428, 278, 278, 278, 278, 436, 278, 278, 278, 278, 278, 278, 437, 278, 278, 278},
		{426, 431, 3: 279, 434, 435, 430, 429, 279, 279, 279, 12: 425, 25: 279, 42: 279, 50: 279, 279, 54: 279, 56: 427, 58: 433, 432, 279, 279, 279, 279, 279, 279, 279, 279, 428, 279, 279, 279, 279, 436, 279, 279, 279, 279, 279, 279, 437, 279, 279, 279},
		{426, 431, 3: 280, 434, 435, 430, 429, 280, 280, 280, 12: 425, 25: 280, 42: 280, 50: 280, 280, 54: 280, 56: 427, 58: 433, 432, 280, 280, 280, 280, 280, 280, 280, 280, 428, 280, 280, 280, 280, 436, 280, 280, 280, 280, 280, 280, 437, 280, 280, 280},
		// 165
		{426, 431, 3: 281, 434, 435, 430, 429, 281, 281, 281, 12: 425, 25: 281, 42: 281, 50: 281, 281, 54: 281, 56: 427, 58: 433, 432, 281, 281, 281, 281, 281, 281, 281, 281, 428, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281},
		{426, 431, 3: 282, 434, 435, 430, 429, 282, 282, 282, 12: 425, 25: 282, 42: 282, 50: 282, 282, 54: 282, 56: 427, 58: 433, 432, 282, 282, 282, 282, 282, 282, 282, 282, 428, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282},
		{426, 431, 3: 283, 283, 283, 430, 429, 283, 283, 283, 12: 425, 25: 283, 42: 283, 50: 283, 283, 54: 283, 56: 427, 58: 433, 432, 283, 283, 283, 283, 283, 283, 283, 283, 428, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283},
		{426, 431, 3: 284, 284, 284, 430, 429, 284, 284, 284, 12: 425, 25: 284, 42: 284, 50: 284, 284, 54: 284, 56: 427, 58: 433, 432, 284, 284, 284, 284, 284, 284, 284, 284, 428, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284},
		{426, 285, 3: 285, 285, 285, 430, 429, 285, 285, 285, 12: 425, 25: 285, 42: 285, 50: 285, 285, 54: 285, 56: 427, 58: 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 428, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285},
		// 170
		{426, 286, 3: 286, 286, 286, 430, 429, 286, 286, 286, 12: 425, 25: 286, 42: 286, 50: 286, 286, 54: 286, 56: 427, 58: 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 428, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286},
		{426, 287, 3: 287, 287, 287, 430, 429, 287, 287, 287, 12: 425, 25: 287, 42: 287, 50: 287, 287, 54: 287, 56: 427, 58: 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 428, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287},
		{302, 302, 3: 302, 302, 302, 302, 302, 302, 302, 302, 12: 302, 25: 302, 42: 302, 50: 302, 302, 54: 302, 56: 302, 58: 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302},
		{303, 303, 3: 303, 303, 303, 303, 303, 303, 303, 303, 12: 303, 25: 303, 42: 303, 50: 303, 303, 54: 303, 56: 303, 58: 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303},
		{426, 431, 3: 444, 434, 435, 430, 429, 9: 317, 317, 12: 425, 51: 450, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		// 175
		{9: 314, 503},
		{9: 502},
		{304, 304, 3: 304, 304, 304, 304, 304, 304, 304, 304, 12: 304, 25: 304, 42: 304, 50: 304, 304, 54: 304, 56: 304, 58: 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 504},
		{426, 431, 3: 444, 434, 435, 430, 429, 9: 316, 316, 12: 425, 51: 450, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		// 180
		{10: 476, 54: 506},
		{305, 305, 3: 305, 305, 305, 305, 305, 305, 305, 305, 12: 305, 25: 305, 42: 305, 50: 305, 305, 54: 305, 56: 305, 58: 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305},
		{161, 161, 161, 161, 161, 161, 161, 161, 9: 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 43: 161, 54: 161},
		{137, 9: 137, 137, 12: 137},
		{54: 513},
		// 185
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 511},
		{426, 431, 3: 444, 434, 435, 430, 429, 12: 425, 51: 450, 54: 512, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{134, 9: 134, 134, 12: 134},
		{136, 9: 136, 136, 12: 136},
		{426, 295, 3: 295, 295, 295, 430, 429, 295, 295, 295, 12: 425, 25: 295, 42: 295, 50: 295, 295, 54: 295, 56: 427, 58: 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 428, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295},
		// 190
		{133, 9: 133, 133, 12: 133},
		{9: 559},
		{9: 156, 22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 113: 336, 117: 355, 358, 354, 335, 124: 521, 337, 127: 334, 161: 520, 172: 518, 519, 199: 522},
		{9: 158, 556},
		{9: 155},
		// 195
		{9: 154, 154},
		{164, 407, 164, 9: 140, 140, 12: 164, 133: 408, 136: 524, 525, 156: 411, 180: 526},
		{9: 523},
		{131, 9: 131, 131, 12: 131},
		{529, 2: 528, 12: 128, 185: 412, 414, 527},
		// 200
		{9: 152, 152},
		{9: 151, 151},
		{533, 8: 178, 178, 178, 12: 532, 22: 178, 178, 178, 26: 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 51: 178, 178, 55: 178, 57: 178},
		{175, 8: 175, 175, 175, 12: 175, 22: 175, 175, 175, 26: 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 51: 175, 175, 55: 175, 57: 175},
		{164, 407, 164, 9: 132, 12: 164, 22: 132, 132, 132, 26: 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 43: 132, 132, 132, 132, 132, 132, 132, 133: 408, 136: 524, 530, 156: 516, 174: 517},
		// 205
		{9: 531},
		{174, 8: 174, 174, 174, 12: 174, 22: 174, 174, 174, 26: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 51: 174, 174, 55: 174, 57: 174},
		{160, 160, 160, 160, 160, 160, 160, 160, 11: 160, 13: 160, 160, 160, 160, 160, 160, 160, 160, 160, 364, 362, 363, 43: 544, 54: 160, 113: 418, 154: 545, 163: 543},
		{2: 536, 9: 148, 22: 169, 169, 169, 26: 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 43: 169, 169, 169, 169, 169, 169, 169, 170: 537, 194: 535, 213: 534},
		{22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 113: 336, 117: 355, 358, 354, 335, 124: 521, 337, 127: 334, 161: 520, 172: 518, 541},
		// 210
		{9: 540},
		{9: 150, 150, 147: 150},
		{9: 147, 538},
		{2: 539},
		{9: 149, 149, 147: 149},
		// 215
		{167, 8: 167, 167, 167, 12: 167, 22: 167, 167, 167, 26: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 51: 167, 167, 55: 167, 57: 167},
		{9: 542},
		{168, 8: 168, 168, 168, 12: 168, 22: 168, 168, 168, 26: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 51: 168, 168, 55: 168, 57: 168},
		{389, 552, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 416, 255, 169: 553},
		{160, 160, 160, 160, 160, 160, 160, 160, 11: 160, 13: 160, 160, 160, 160, 160, 160, 160, 160, 160, 364, 362, 363, 113: 418, 154: 422, 163: 549},
		// 220
		{159, 159, 159, 159, 159, 159, 159, 159, 11: 159, 13: 159, 159, 159, 159, 159, 159, 159, 159, 159, 364, 362, 363, 43: 546, 54: 159, 113: 507},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 547},
		{426, 431, 3: 444, 434, 435, 430, 429, 12: 425, 51: 450, 54: 548, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{171, 8: 171, 171, 171, 12: 171, 22: 171, 171, 171, 26: 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 51: 171, 171, 55: 171, 57: 171},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 550},
		// 225
		{426, 431, 3: 444, 434, 435, 430, 429, 12: 425, 51: 450, 54: 551, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{172, 8: 172, 172, 172, 12: 172, 22: 172, 172, 172, 26: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 51: 172, 172, 55: 172, 57: 172},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 514, 555},
		{54: 554},
		{173, 8: 173, 173, 173, 12: 173, 22: 173, 173, 173, 26: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 51: 173, 173, 55: 173, 57: 173},
		// 230
		{170, 8: 170, 170, 170, 12: 170, 22: 170, 170, 170, 26: 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 51: 170, 170, 55: 170, 57: 170},
		{22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 113: 336, 117: 355, 358, 354, 335, 124: 521, 337, 127: 334, 147: 557, 161: 558},
		{9: 157},
		{9: 153, 153},
		{138, 9: 138, 138, 12: 138},
		// 235
		{9: 130, 22: 130, 130, 130, 26: 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 43: 130, 130, 130, 130, 130, 130, 130, 205: 561},
		{9: 156, 22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 113: 336, 117: 355, 358, 354, 335, 124: 521, 337, 127: 334, 161: 520, 172: 518, 519, 199: 562},
		{9: 563},
		{129, 9: 129, 129, 12: 129},
		{166, 407, 166, 9: 166, 166, 12: 166, 133: 565},
		// 240
		{165, 2: 165, 9: 165, 165, 12: 165},
		{199, 199, 199, 8: 199, 199, 12: 199, 42: 199},
		{197, 197, 197, 8: 197, 197, 12: 197, 42: 197},
		{200, 200, 200, 8: 200, 200, 12: 200, 42: 200},
		{257, 257, 3: 257, 257, 257, 257, 257, 257, 257, 257, 12: 257, 25: 257, 42: 257, 50: 257, 257, 54: 257, 56: 257, 58: 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257},
		// 245
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 144, 144, 144, 26: 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 52: 575, 473, 114: 572, 123: 574, 155: 403, 162: 573},
		{426, 290, 3: 290, 290, 290, 430, 429, 290, 290, 290, 12: 425, 25: 290, 42: 290, 50: 290, 290, 54: 290, 56: 427, 58: 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 428, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290},
		{9: 744, 476},
		{9: 738},
		{9: 737},
		// 250
		{103, 103, 103, 103, 103, 103, 103, 103, 103, 11: 103, 13: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 43: 103, 103, 103, 103, 103, 103, 103, 52: 103, 55: 103, 57: 103, 85: 103, 87: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 206: 576},
		{389, 394, 577, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 364, 362, 363, 99, 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 52: 575, 473, 55: 604, 57: 374, 85: 600, 87: 597, 602, 587, 601, 588, 598, 599, 595, 603, 596, 113: 336, 578, 117: 355, 358, 354, 335, 594, 605, 581, 579, 337, 127: 334, 586, 582, 584, 585, 580, 134: 583, 593, 138: 333, 144: 592, 181: 590, 216: 591, 589},
		{313, 313, 3: 313, 313, 313, 313, 313, 313, 10: 313, 12: 313, 42: 735, 51: 313, 56: 313, 58: 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313},
		{8: 250, 250, 476},
		{164, 407, 164, 8: 237, 133: 704, 136: 703, 702, 171: 700, 196: 701, 699},
		// 255
		{113, 113, 113, 113, 113, 113, 113, 113, 113, 11: 113, 13: 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 43: 113, 113, 113, 113, 113, 113, 113, 52: 113, 55: 113, 57: 113, 85: 113, 87: 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 112: 113},
		{112, 112, 112, 112, 112, 112, 112, 112, 112, 11: 112, 13: 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 43: 112, 112, 112, 112, 112, 112, 112, 52: 112, 55: 112, 57: 112, 85: 112, 87: 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112: 112},
		{111, 111, 111, 111, 111, 111, 111, 111, 111, 11: 111, 13: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 43: 111, 111, 111, 111, 111, 111, 111, 52: 111, 55: 111, 57: 111, 85: 111, 87: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 112: 111},
		{110, 110, 110, 110, 110, 110, 110, 110, 110, 11: 110, 13: 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 43: 110, 110, 110, 110, 110, 110, 110, 52: 110, 55: 110, 57: 110, 85: 110, 87: 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 112: 110},
		{109, 109, 109, 109, 109, 109, 109, 109, 109, 11: 109, 13: 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 43: 109, 109, 109, 109, 109, 109, 109, 52: 109, 55: 109, 57: 109, 85: 109, 87: 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 112: 109},
		// 260
		{108, 108, 108, 108, 108, 108, 108, 108, 108, 11: 108, 13: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 43: 108, 108, 108, 108, 108, 108, 108, 52: 108, 55: 108, 57: 108, 85: 108, 87: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 112: 108},
		{107, 107, 107, 107, 107, 107, 107, 107, 107, 11: 107, 13: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 43: 107, 107, 107, 107, 107, 107, 107, 52: 107, 55: 107, 57: 107, 85: 107, 87: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 112: 107},
		{249, 249, 249, 249, 249, 249, 249, 249, 11: 249, 13: 249, 249, 249, 249, 249, 249, 249, 249, 249, 145: 376, 696},
		{42: 694},
		{25: 693},
		// 265
		{101, 101, 101, 101, 101, 101, 101, 101, 101, 11: 101, 13: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 43: 101, 101, 101, 101, 101, 101, 101, 52: 101, 55: 101, 57: 101, 85: 101, 87: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101},
		{389, 394, 577, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 364, 362, 363, 98, 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 52: 575, 473, 55: 604, 57: 374, 85: 600, 87: 597, 602, 587, 601, 588, 598, 599, 595, 603, 596, 113: 336, 578, 117: 355, 358, 354, 335, 594, 605, 581, 579, 337, 127: 334, 586, 582, 584, 585, 580, 134: 583, 593, 138: 333, 144: 592, 181: 692},
		{97, 97, 97, 97, 97, 97, 97, 97, 97, 11: 97, 13: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 43: 97, 97, 97, 97, 97, 97, 97, 52: 97, 55: 97, 57: 97, 85: 97, 87: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97},
		{96, 96, 96, 96, 96, 96, 96, 96, 96, 11: 96, 13: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 43: 96, 96, 96, 96, 96, 96, 96, 52: 96, 55: 96, 57: 96, 85: 96, 87: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96},
		{8: 691},
		// 270
		{685},
		{681},
		{677},
		{389, 394, 577, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 575, 473, 55: 604, 85: 600, 87: 597, 602, 587, 601, 588, 598, 599, 595, 603, 596, 114: 578, 121: 594, 605, 581, 128: 586, 582, 584, 585, 580, 134: 583, 671},
		{657},
		// 275
		{2: 655},
		{8: 654},
		{8: 653},
		{389, 394, 382, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 473, 114: 578, 121: 651},
		{60, 22: 607, 85: 60, 203: 606},
		// 280
		{51, 51, 51, 51, 51, 51, 51, 51, 51, 11: 51, 13: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 43: 51, 51, 51, 51, 51, 51, 51, 52: 51, 55: 51, 57: 51, 85: 51, 87: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 112: 51},
		{608, 85: 609},
		{59, 85: 59},
		{11: 611, 164: 636},
		{610},
		// 285
		{11: 611, 164: 612},
		{9: 63, 11: 63, 42: 63},
		{11: 613, 42: 614},
		{9: 62, 11: 62, 42: 62},
		{42: 615},
		// 290
		{11: 55, 619, 149: 617, 616, 157: 618},
		{11: 632},
		{9: 57, 57, 42: 57},
		{10: 622, 42: 623},
		{2: 620},
		// 295
		{54: 621},
		{11: 54},
		{11: 55, 619, 149: 631, 616},
		{11: 624, 182: 625},
		{9: 53, 53, 42: 53},
		// 300
		{10: 626, 42: 627},
		{11: 630},
		{2: 536, 170: 628},
		{9: 629, 538},
		{47, 47, 47, 47, 47, 47, 47, 47, 47, 11: 47, 13: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 43: 47, 47, 47, 47, 47, 47, 47, 52: 47, 55: 47, 57: 47, 85: 47, 87: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 112: 47},
		// 305
		{9: 52, 52, 42: 52},
		{9: 56, 56, 42: 56},
		{633},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 634},
		{426, 431, 3: 444, 434, 435, 430, 429, 9: 635, 12: 425, 51: 450, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		// 310
		{9: 58, 58, 42: 58},
		{9: 637, 11: 613, 42: 638},
		{61, 61, 61, 61, 61, 61, 61, 61, 61, 11: 61, 13: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 43: 61, 61, 61, 61, 61, 61, 61, 52: 61, 55: 61, 57: 61, 85: 61, 87: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 112: 61},
		{9: 640, 11: 55, 619, 42: 641, 149: 617, 616, 157: 639},
		{9: 644, 622, 42: 645},
		// 315
		{46, 46, 46, 46, 46, 46, 46, 46, 46, 11: 46, 13: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 43: 46, 46, 46, 46, 46, 46, 46, 52: 46, 55: 46, 57: 46, 85: 46, 87: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 112: 46},
		{11: 55, 619, 149: 617, 616, 157: 642},
		{9: 643, 622},
		{45, 45, 45, 45, 45, 45, 45, 45, 45, 11: 45, 13: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 43: 45, 45, 45, 45, 45, 45, 45, 52: 45, 55: 45, 57: 45, 85: 45, 87: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 112: 45},
		{50, 50, 50, 50, 50, 50, 50, 50, 50, 11: 50, 13: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 43: 50, 50, 50, 50, 50, 50, 50, 52: 50, 55: 50, 57: 50, 85: 50, 87: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 112: 50},
		// 320
		{11: 55, 619, 149: 617, 616, 157: 646},
		{9: 647, 622, 42: 648},
		{49, 49, 49, 49, 49, 49, 49, 49, 49, 11: 49, 13: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 43: 49, 49, 49, 49, 49, 49, 49, 52: 49, 55: 49, 57: 49, 85: 49, 87: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 112: 49},
		{11: 624, 182: 649},
		{9: 650, 626},
		// 325
		{48, 48, 48, 48, 48, 48, 48, 48, 48, 11: 48, 13: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 43: 48, 48, 48, 48, 48, 48, 48, 52: 48, 55: 48, 57: 48, 85: 48, 87: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 112: 48},
		{8: 652},
		{84, 84, 84, 84, 84, 84, 84, 84, 84, 11: 84, 13: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 43: 84, 84, 84, 84, 84, 84, 84, 52: 84, 55: 84, 57: 84, 85: 84, 87: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 112: 84},
		{85, 85, 85, 85, 85, 85, 85, 85, 85, 11: 85, 13: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 43: 85, 85, 85, 85, 85, 85, 85, 52: 85, 55: 85, 57: 85, 85: 85, 87: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 112: 85},
		{86, 86, 86, 86, 86, 86, 86, 86, 86, 11: 86, 13: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 43: 86, 86, 86, 86, 86, 86, 86, 52: 86, 55: 86, 57: 86, 85: 86, 87: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 112: 86},
		// 330
		{8: 656},
		{87, 87, 87, 87, 87, 87, 87, 87, 87, 11: 87, 13: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 43: 87, 87, 87, 87, 87, 87, 87, 52: 87, 55: 87, 57: 87, 85: 87, 87: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 112: 87},
		{389, 394, 382, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 53: 473, 57: 374, 113: 336, 578, 117: 355, 358, 354, 335, 658, 124: 579, 337, 127: 334, 138: 333, 144: 659},
		{8: 665},
		{389, 394, 382, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 473, 114: 578, 121: 660},
		// 335
		{8: 661},
		{389, 394, 382, 393, 395, 396, 392, 391, 9: 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 473, 114: 578, 121: 662},
		{9: 663},
		{389, 394, 577, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 575, 473, 55: 604, 85: 600, 87: 597, 602, 587, 601, 588, 598, 599, 595, 603, 596, 114: 578, 121: 594, 605, 581, 128: 586, 582, 584, 585, 580, 134: 583, 664},
		{88, 88, 88, 88, 88, 88, 88, 88, 88, 11: 88, 13: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 43: 88, 88, 88, 88, 88, 88, 88, 52: 88, 55: 88, 57: 88, 85: 88, 87: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 112: 88},
		// 340
		{389, 394, 382, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 473, 114: 578, 121: 666},
		{8: 667},
		{389, 394, 382, 393, 395, 396, 392, 391, 9: 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 473, 114: 578, 121: 668},
		{9: 669},
		{389, 394, 577, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 575, 473, 55: 604, 85: 600, 87: 597, 602, 587, 601, 588, 598, 599, 595, 603, 596, 114: 578, 121: 594, 605, 581, 128: 586, 582, 584, 585, 580, 134: 583, 670},
		// 345
		{89, 89, 89, 89, 89, 89, 89, 89, 89, 11: 89, 13: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 43: 89, 89, 89, 89, 89, 89, 89, 52: 89, 55: 89, 57: 89, 85: 89, 87: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 112: 89},
		{87: 672},
		{673},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 473, 114: 674},
		{9: 675, 476},
		// 350
		{8: 676},
		{90, 90, 90, 90, 90, 90, 90, 90, 90, 11: 90, 13: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 43: 90, 90, 90, 90, 90, 90, 90, 52: 90, 55: 90, 57: 90, 85: 90, 87: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 112: 90},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 473, 114: 678},
		{9: 679, 476},
		{389, 394, 577, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 575, 473, 55: 604, 85: 600, 87: 597, 602, 587, 601, 588, 598, 599, 595, 603, 596, 114: 578, 121: 594, 605, 581, 128: 586, 582, 584, 585, 580, 134: 583, 680},
		// 355
		{91, 91, 91, 91, 91, 91, 91, 91, 91, 11: 91, 13: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 43: 91, 91, 91, 91, 91, 91, 91, 52: 91, 55: 91, 57: 91, 85: 91, 87: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 112: 91},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 473, 114: 682},
		{9: 683, 476},
		{389, 394, 577, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 575, 473, 55: 604, 85: 600, 87: 597, 602, 587, 601, 588, 598, 599, 595, 603, 596, 114: 578, 121: 594, 605, 581, 128: 586, 582, 584, 585, 580, 134: 583, 684},
		{92, 92, 92, 92, 92, 92, 92, 92, 92, 11: 92, 13: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 43: 92, 92, 92, 92, 92, 92, 92, 52: 92, 55: 92, 57: 92, 85: 92, 87: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 112: 92},
		// 360
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 53: 473, 114: 686},
		{9: 687, 476},
		{389, 394, 577, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 575, 473, 55: 604, 85: 600, 87: 597, 602, 587, 601, 588, 598, 599, 595, 603, 596, 114: 578, 121: 594, 605, 581, 128: 586, 582, 584, 585, 580, 134: 583, 688},
		{94, 94, 94, 94, 94, 94, 94, 94, 94, 11: 94, 13: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 43: 94, 94, 94, 94, 94, 94, 94, 52: 94, 55: 94, 57: 94, 85: 94, 87: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 112: 689},
		{389, 394, 577, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 575, 473, 55: 604, 85: 600, 87: 597, 602, 587, 601, 588, 598, 599, 595, 603, 596, 114: 578, 121: 594, 605, 581, 128: 586, 582, 584, 585, 580, 134: 583, 690},
		// 365
		{93, 93, 93, 93, 93, 93, 93, 93, 93, 11: 93, 13: 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 43: 93, 93, 93, 93, 93, 93, 93, 52: 93, 55: 93, 57: 93, 85: 93, 87: 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 112: 93},
		{95, 95, 95, 95, 95, 95, 95, 95, 95, 11: 95, 13: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 43: 95, 95, 95, 95, 95, 95, 95, 52: 95, 55: 95, 57: 95, 85: 95, 87: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 112: 95},
		{100, 100, 100, 100, 100, 100, 100, 100, 100, 11: 100, 13: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 43: 100, 100, 100, 100, 100, 100, 100, 52: 100, 55: 100, 57: 100, 85: 100, 87: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
		{102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 11: 102, 13: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 43: 102, 102, 102, 102, 102, 102, 102, 102, 52: 102, 55: 102, 57: 102, 85: 102, 87: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 112: 102},
		{389, 394, 577, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 575, 473, 55: 604, 85: 600, 87: 597, 602, 587, 601, 588, 598, 599, 595, 603, 596, 114: 578, 121: 594, 605, 581, 128: 586, 582, 584, 585, 580, 134: 583, 695},
		// 370
		{104, 104, 104, 104, 104, 104, 104, 104, 104, 11: 104, 13: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 43: 104, 104, 104, 104, 104, 104, 104, 52: 104, 55: 104, 57: 104, 85: 104, 87: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 112: 104},
		{42: 697},
		{389, 394, 577, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 575, 473, 55: 604, 85: 600, 87: 597, 602, 587, 601, 588, 598, 599, 595, 603, 596, 114: 578, 121: 594, 605, 581, 128: 586, 582, 584, 585, 580, 134: 583, 698},
		{105, 105, 105, 105, 105, 105, 105, 105, 105, 11: 105, 13: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 43: 105, 105, 105, 105, 105, 105, 105, 52: 105, 55: 105, 57: 105, 85: 105, 87: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 112: 105},
		{8: 734},
		// 375
		{8: 239, 10: 239},
		{8: 236, 10: 732},
		{8: 235, 10: 235, 51: 234, 179: 706},
		{705, 2: 528, 187: 527},
		{163, 2: 163},
		// 380
		{164, 407, 164, 133: 704, 136: 703, 530},
		{51: 707},
		{389, 394, 708, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 711, 709, 160: 710},
		{313, 313, 3: 313, 313, 313, 313, 313, 313, 10: 313, 12: 313, 25: 313, 42: 730, 51: 313, 56: 313, 58: 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313},
		{426, 431, 3: 444, 434, 435, 430, 429, 126, 10: 126, 12: 425, 25: 126, 51: 450, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		// 385
		{8: 233, 10: 233},
		{119, 119, 119, 119, 119, 119, 119, 119, 10: 121, 119, 717, 119, 119, 119, 119, 119, 119, 119, 119, 119, 25: 121, 52: 119, 56: 718, 159: 716, 166: 715, 713, 714, 198: 712},
		{10: 725, 25: 192, 165: 726},
		{389, 394, 708, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 711, 709, 160: 724},
		{12: 717, 51: 722, 56: 718, 159: 723},
		// 390
		{118, 118, 118, 118, 118, 118, 118, 118, 11: 118, 13: 118, 118, 118, 118, 118, 118, 118, 118, 118, 52: 118},
		{12: 117, 51: 117, 56: 117},
		{249, 249, 249, 249, 249, 249, 249, 249, 11: 249, 13: 249, 249, 249, 249, 249, 249, 249, 249, 249, 145: 376, 720},
		{2: 719},
		{12: 114, 51: 114, 56: 114},
		// 395
		{54: 721},
		{12: 115, 51: 115, 56: 115},
		{120, 120, 120, 120, 120, 120, 120, 120, 11: 120, 13: 120, 120, 120, 120, 120, 120, 120, 120, 120, 52: 120},
		{12: 116, 51: 116, 56: 116},
		{10: 123, 25: 123},
		// 400
		{119, 119, 119, 119, 119, 119, 119, 119, 11: 119, 717, 119, 119, 119, 119, 119, 119, 119, 119, 119, 25: 191, 52: 119, 56: 718, 159: 716, 166: 715, 728, 714},
		{25: 727},
		{8: 125, 10: 125, 25: 125},
		{389, 394, 708, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 711, 709, 160: 729},
		{10: 122, 25: 122},
		// 405
		{389, 394, 708, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 711, 709, 160: 731},
		{8: 124, 10: 124, 25: 124},
		{164, 407, 164, 133: 704, 136: 703, 702, 171: 733},
		{8: 238, 10: 238},
		{247, 247, 247, 247, 247, 247, 247, 247, 247, 11: 247, 13: 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 43: 247, 247, 247, 247, 247, 247, 247, 247, 52: 247, 55: 247, 57: 247, 85: 247, 87: 247, 247, 247, 247, 247, 247, 247, 247, 247, 247},
		// 410
		{389, 394, 577, 393, 395, 396, 392, 391, 251, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 575, 473, 55: 604, 85: 600, 87: 597, 602, 587, 601, 588, 598, 599, 595, 603, 596, 114: 578, 121: 594, 605, 581, 128: 586, 582, 584, 585, 580, 134: 583, 736},
		{106, 106, 106, 106, 106, 106, 106, 106, 106, 11: 106, 13: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 43: 106, 106, 106, 106, 106, 106, 106, 52: 106, 55: 106, 57: 106, 85: 106, 87: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 112: 106},
		{256, 256, 3: 256, 256, 256, 256, 256, 256, 256, 256, 12: 256, 25: 256, 42: 256, 50: 256, 256, 54: 256, 56: 256, 58: 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256},
		{389, 289, 382, 289, 289, 289, 392, 391, 289, 289, 289, 388, 289, 398, 397, 400, 383, 384, 385, 386, 387, 399, 25: 289, 42: 289, 50: 289, 289, 740, 739, 289, 56: 289, 58: 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289},
		{426, 288, 3: 288, 288, 288, 430, 429, 288, 288, 288, 12: 425, 25: 288, 42: 288, 50: 288, 288, 54: 288, 56: 427, 58: 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 428, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288},
		// 415
		{119, 119, 119, 119, 119, 119, 119, 119, 10: 121, 119, 717, 119, 119, 119, 119, 119, 119, 119, 119, 119, 25: 121, 52: 119, 56: 718, 159: 716, 166: 715, 713, 714, 198: 741},
		{10: 725, 25: 192, 165: 742},
		{25: 743},
		{299, 299, 3: 299, 299, 299, 299, 299, 299, 299, 299, 12: 299, 25: 299, 42: 299, 50: 299, 299, 54: 299, 56: 299, 58: 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299},
		{306, 306, 3: 306, 306, 306, 306, 306, 306, 306, 306, 12: 306, 25: 306, 42: 306, 50: 306, 306, 54: 306, 56: 306, 58: 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306},
		// 420
		{426, 291, 3: 291, 291, 291, 430, 429, 291, 291, 291, 12: 425, 25: 291, 42: 291, 50: 291, 291, 54: 291, 56: 427, 58: 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 428, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291},
		{426, 292, 3: 292, 292, 292, 430, 429, 292, 292, 292, 12: 425, 25: 292, 42: 292, 50: 292, 292, 54: 292, 56: 427, 58: 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 428, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292},
		{426, 293, 3: 293, 293, 293, 430, 429, 293, 293, 293, 12: 425, 25: 293, 42: 293, 50: 293, 293, 54: 293, 56: 427, 58: 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 428, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293},
		{426, 294, 3: 294, 294, 294, 430, 429, 294, 294, 294, 12: 425, 25: 294, 42: 294, 50: 294, 294, 54: 294, 56: 427, 58: 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 428, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294},
		{426, 296, 3: 296, 296, 296, 430, 429, 296, 296, 296, 12: 425, 25: 296, 42: 296, 50: 296, 296, 54: 296, 56: 427, 58: 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 428, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296},
		// 425
		{426, 297, 3: 297, 297, 297, 430, 429, 297, 297, 297, 12: 425, 25: 297, 42: 297, 50: 297, 297, 54: 297, 56: 427, 58: 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 428, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297},
		{426, 298, 3: 298, 298, 298, 430, 429, 298, 298, 298, 12: 425, 25: 298, 42: 298, 50: 298, 298, 54: 298, 56: 427, 58: 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 428, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298},
		{9: 753},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 52: 740, 739},
		{755},
		// 430
		{11: 611, 164: 756},
		{9: 637, 11: 613},
		{22: 65, 65, 65, 26: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 43: 65, 65, 65, 65, 65, 65, 65, 52: 66, 55: 66, 57: 65, 178: 759, 183: 758},
		{52: 74, 55: 74, 208: 763},
		{22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 57: 374, 113: 336, 117: 355, 358, 354, 335, 124: 579, 337, 127: 334, 138: 333, 144: 760, 219: 761},
		// 435
		{22: 68, 68, 68, 26: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 43: 68, 68, 68, 68, 68, 68, 68, 52: 68, 55: 68, 57: 68},
		{22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 43: 340, 341, 339, 365, 366, 342, 338, 52: 64, 55: 64, 57: 374, 113: 336, 117: 355, 358, 354, 335, 124: 579, 337, 127: 334, 138: 333, 144: 762},
		{22: 67, 67, 67, 26: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 43: 67, 67, 67, 67, 67, 67, 67, 52: 67, 55: 67, 57: 67},
		{52: 72, 55: 70, 176: 765, 766, 192: 764},
		{73, 73, 73, 8: 73, 22: 73, 73, 73, 26: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 43: 73, 73, 73, 73, 73, 73, 73, 73, 55: 73, 57: 73},
		// 440
		{52: 575, 123: 769},
		{55: 604, 122: 605, 128: 767},
		{8: 768},
		{69, 69, 69, 8: 69, 22: 69, 69, 69, 26: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 43: 69, 69, 69, 69, 69, 69, 69, 69, 55: 69, 57: 69},
		{71, 71, 71, 8: 71, 22: 71, 71, 71, 26: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 43: 71, 71, 71, 71, 71, 71, 71, 71, 55: 71, 57: 71},
		// 445
		{79, 79, 79, 8: 79, 22: 79, 79, 79, 26: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 43: 79, 79, 79, 79, 79, 79, 79, 79, 55: 79, 57: 79},
		{52: 190, 212: 773},
		{188, 188, 188, 8: 188, 188, 188, 12: 188, 22: 188, 188, 188, 26: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 52: 145},
		{52: 774},
		{2: 775, 189: 778, 777, 225: 776},
		// 450
		{10: 318, 25: 318, 51: 318},
		{10: 781, 25: 192, 165: 782},
		{10: 187, 25: 187},
		{10: 185, 25: 185, 51: 779},
		{249, 249, 249, 249, 249, 249, 249, 249, 11: 249, 13: 249, 249, 249, 249, 249, 249, 249, 249, 249, 145: 376, 780},
		// 455
		{10: 184, 25: 184},
		{2: 775, 25: 191, 189: 778, 784},
		{25: 783},
		{189, 189, 189, 8: 189, 189, 189, 12: 189, 22: 189, 189, 189, 26: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189},
		{10: 186, 25: 186},
		// 460
		{52: 787},
		{209, 209, 209, 8: 209, 209, 209, 12: 209, 22: 209, 209, 209, 26: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 52: 145},
		{22: 211, 211, 211, 789, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 57: 211, 211: 788},
		{22: 364, 362, 363, 26: 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 57: 374, 113: 405, 117: 355, 358, 354, 404, 138: 793, 153: 792, 201: 791, 230: 790},
		{208, 208, 208, 8: 208, 208, 208, 12: 208, 22: 208, 208, 208, 26: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208},
		// 465
		{22: 364, 362, 363, 804, 352, 344, 353, 349, 361, 348, 346, 347, 345, 350, 359, 356, 357, 360, 351, 343, 57: 374, 113: 405, 117: 355, 358, 354, 404, 138: 793, 153: 792, 201: 805},
		{22: 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 57: 205},
		{164, 407, 164, 8: 795, 42: 177, 133: 704, 136: 703, 797, 184: 798, 202: 796, 231: 794},
		{22: 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 57: 201},
		{8: 801, 10: 802},
		// 470
		{22: 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 57: 202},
		{8: 196, 10: 196},
		{8: 194, 10: 194, 42: 176},
		{42: 799},
		{249, 249, 249, 249, 249, 249, 249, 249, 11: 249, 13: 249, 249, 249, 249, 249, 249, 249, 249, 249, 145: 376, 800},
		// 475
		{8: 193, 10: 193},
		{22: 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 57: 203},
		{164, 407, 164, 42: 177, 133: 704, 136: 703, 797, 184: 798, 202: 803},
		{8: 195, 10: 195},
		{210, 210, 210, 8: 210, 210, 210, 12: 210, 22: 210, 210, 210, 26: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210},
		// 480
		{22: 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 57: 204},
		{389, 394, 382, 393, 395, 396, 392, 391, 11: 388, 13: 398, 397, 400, 383, 384, 385, 386, 387, 399, 144, 144, 144, 26: 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 53: 807, 155: 403, 162: 808},
		{426, 431, 3: 444, 434, 435, 430, 429, 9: 810, 12: 425, 51: 450, 56: 427, 58: 433, 432, 438, 439, 449, 445, 446, 454, 447, 458, 428, 452, 442, 441, 440, 436, 456, 453, 451, 443, 460, 448, 437, 457, 455, 459},
		{9: 809},
		{212, 212, 212, 8: 212, 212, 212, 12: 212, 22: 212, 212, 212, 26: 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212},
		// 485
		{213, 213, 213, 8: 213, 213, 213, 12: 213, 22: 213, 213, 213, 26: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213},
		{242, 242, 242, 8: 242, 242, 242, 12: 242},
		{240, 240, 240, 8: 240, 240, 240, 12: 240},
		{243, 243, 243, 8: 243, 243, 243, 12: 243},
		{244, 244, 244, 8: 244, 244, 244, 12: 244},
		// 490
		{245, 245, 245, 8: 245, 245, 245, 12: 245},
		{8: 235, 10: 235, 22: 65, 65, 65, 26: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 43: 65, 65, 65, 65, 65, 65, 65, 51: 234, 66, 55: 66, 57: 65, 178: 759, 706, 183: 817},
		{52: 77, 55: 77, 207: 818},
		{52: 72, 55: 70, 176: 765, 766, 192: 819},
		{76, 76, 76, 8: 76, 22: 76, 76, 76, 26: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 43: 76, 76, 76, 76, 76, 76, 76, 76, 55: 76, 57: 76},
		// 495
		{82, 82, 82, 8: 82, 22: 82, 82, 82, 26: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 43: 82, 82, 82, 82, 82, 82, 82, 82, 55: 82, 57: 82},
		{249, 249, 249, 249, 249, 249, 249, 249, 11: 249, 13: 249, 249, 249, 249, 249, 249, 249, 249, 249, 145: 376, 822},
		{50: 321},
		{84: 845, 86: 847, 100: 835, 836, 837, 832, 833, 834, 838, 842, 839, 829, 840, 841, 115: 846, 844, 126: 843, 139: 827, 826, 831, 828, 830, 148: 825, 228: 824},
		{50: 323},
		// 500
		{50: 43, 84: 845, 86: 847, 100: 835, 836, 837, 832, 833, 834, 838, 842, 839, 829, 840, 841, 115: 846, 844, 126: 843, 139: 827, 885, 831, 828, 830},
		{50: 42, 84: 42, 86: 42, 97: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{50: 38, 84: 38, 86: 38, 97: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{50: 37, 84: 37, 86: 37, 97: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{86: 847, 115: 907, 844},
		// 505
		{50: 35, 84: 35, 86: 35, 97: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{97: 28, 28, 895, 188: 893, 220: 894, 892},
		{86: 847, 115: 889, 844},
		{2: 886},
		{2: 881},
		// 510
		{2: 862, 84: 864, 226: 863},
		{84: 845, 86: 847, 115: 846, 844, 126: 861},
		{50: 16, 84: 16, 86: 16, 97: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{86: 847, 115: 859, 844},
		{86: 847, 115: 857, 844},
		// 515
		{84: 845, 86: 847, 115: 846, 844, 126: 856},
		{2: 852},
		{86: 847, 115: 850, 844},
		{50: 7, 84: 7, 86: 7, 97: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{84: 5, 86: 849},
		// 520
		{50: 4, 84: 4, 86: 4, 97: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{84: 848},
		{84: 2, 86: 2},
		{50: 3, 84: 3, 86: 3, 97: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{84: 1, 86: 1},
		// 525
		{84: 851},
		{50: 8, 84: 8, 86: 8, 97: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{84: 853, 86: 847, 115: 854, 844},
		{50: 12, 84: 12, 86: 12, 97: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{84: 855},
		// 530
		{50: 9, 84: 9, 86: 9, 97: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{50: 13, 84: 13, 86: 13, 97: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{84: 858},
		{50: 14, 84: 14, 86: 14, 97: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{84: 860},
		// 535
		{50: 15, 84: 15, 86: 15, 97: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{50: 17, 84: 17, 86: 17, 97: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{84: 845, 86: 847, 115: 846, 844, 126: 870, 152: 880},
		{2: 536, 9: 148, 147: 866, 170: 865, 194: 867},
		{50: 10, 84: 10, 86: 10, 97: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		// 540
		{9: 147, 873, 147: 874},
		{9: 871},
		{9: 868},
		{84: 845, 86: 847, 115: 846, 844, 126: 870, 152: 869},
		{50: 18, 84: 18, 86: 18, 97: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		// 545
		{50: 6, 84: 6, 86: 6, 97: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{84: 845, 86: 847, 115: 846, 844, 126: 870, 152: 872},
		{50: 20, 84: 20, 86: 20, 97: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{2: 539, 147: 877},
		{9: 875},
		// 550
		{84: 845, 86: 847, 115: 846, 844, 126: 870, 152: 876},
		{50: 11, 84: 11, 86: 11, 97: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{9: 878},
		{84: 845, 86: 847, 115: 846, 844, 126: 870, 152: 879},
		{50: 19, 84: 19, 86: 19, 97: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		// 555
		{50: 21, 84: 21, 86: 21, 97: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{84: 882},
		{84: 845, 86: 847, 97: 40, 40, 40, 835, 836, 837, 832, 833, 834, 838, 842, 839, 829, 840, 841, 115: 846, 844, 126: 843, 139: 827, 826, 831, 828, 830, 148: 883, 151: 884},
		{84: 845, 86: 847, 97: 39, 39, 39, 835, 836, 837, 832, 833, 834, 838, 842, 839, 829, 840, 841, 115: 846, 844, 126: 843, 139: 827, 885, 831, 828, 830},
		{97: 31, 31, 31},
		// 560
		{50: 41, 84: 41, 86: 41, 97: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		{84: 887},
		{84: 845, 86: 847, 97: 40, 40, 40, 835, 836, 837, 832, 833, 834, 838, 842, 839, 829, 840, 841, 115: 846, 844, 126: 843, 139: 827, 826, 831, 828, 830, 148: 883, 151: 888},
		{97: 32, 32, 32},
		{84: 890},
		// 565
		{84: 845, 86: 847, 97: 40, 40, 40, 835, 836, 837, 832, 833, 834, 838, 842, 839, 829, 840, 841, 115: 846, 844, 126: 843, 139: 827, 826, 831, 828, 830, 148: 883, 151: 891},
		{97: 33, 33, 33},
		{97: 24, 901, 222: 902, 900},
		{97: 30, 30, 30},
		{97: 27, 27, 895, 188: 899},
		// 570
		{86: 847, 115: 896, 844},
		{84: 897},
		{84: 845, 86: 847, 97: 40, 40, 40, 835, 836, 837, 832, 833, 834, 838, 842, 839, 829, 840, 841, 115: 846, 844, 126: 843, 139: 827, 826, 831, 828, 830, 148: 883, 151: 898},
		{97: 26, 26, 26},
		{97: 29, 29, 29},
		// 575
		{97: 906, 224: 905},
		{84: 903},
		{97: 23},
		{84: 845, 86: 847, 97: 40, 100: 835, 836, 837, 832, 833, 834, 838, 842, 839, 829, 840, 841, 115: 846, 844, 126: 843, 139: 827, 826, 831, 828, 830, 148: 883, 151: 904},
		{97: 25},
		// 580
		{50: 34, 84: 34, 86: 34, 97: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{50: 22, 84: 22, 86: 22, 97: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{84: 908},
		{50: 36, 84: 36, 86: 36, 97: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
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
	const yyError = 236

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

			if lhs.Expression.Case == 0 { // IDENTIFIER
				if lx.tweaks.enableBuiltinConstantP && lhs.Expression.Token.Val == idBuiltinConstantP {
					break
				}

				b := lhs.Expression.scope.Lookup(NSIdentifiers, lhs.Expression.Token.Val)
				if b.Node == nil && lx.tweaks.enableImplicitFuncDef {
					for l := o.ArgumentExpressionList; l != nil; l = l.ArgumentExpressionList {
						l.Expression.eval(lx)
					}
					break
				}
			}

			lhs.Expression.eval(lx)
			for l := o.ArgumentExpressionList; l != nil; l = l.ArgumentExpressionList {
				l.Expression.eval(lx)
			}
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
			yyVAL.node = &Expression{
				Case:              57,
				Token:             yyS[yypt-2].Token,
				CompoundStatement: yyS[yypt-1].node.(*CompoundStatement),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 70:
		{
			yyVAL.node = (*ExpressionOpt)(nil)
		}
	case 71:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionOpt{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 72:
		{
			yyVAL.node = &ExpressionList{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 73:
		{
			yyVAL.node = &ExpressionList{
				Case:           1,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList),
				Token:          yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 74:
		{
			yyVAL.node = (*ExpressionListOpt)(nil)
		}
	case 75:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionListOpt{
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 76:
		{
			lx := yylex.(*lexer)
			lx.constExprToks = []xc.Token{lx.last}
		}
	case 77:
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
	case 78:
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
	case 79:
		{
			yyVAL.node = &Declaration{
				Case: 1,
				StaticAssertDeclaration: yyS[yypt-0].node.(*StaticAssertDeclaration),
			}
		}
	case 80:
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
	case 81:
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
			tsb := tsDecode(b.typeSpecifier)
			if len(tsb) == 1 && tsb[0] == tsTypedefName && lx.tweaks.allowCompatibleTypedefRedefinitions {
				tsb[0] = 0
			}
			ts := tsEncode(append(tsDecode(a.typeSpecifier), tsb...)...)
			if _, ok := tsValid[ts]; !ok {
				ts = tsEncode(tsInt)
				lx.report.Err(a.Pos(), "invalid type specifier")
			}
			lhs.typeSpecifier = ts
		}
	case 82:
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
	case 83:
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
	case 84:
		{
			yyVAL.node = (*DeclarationSpecifiersOpt)(nil)
		}
	case 85:
		{
			lhs := &DeclarationSpecifiersOpt{
				DeclarationSpecifiers: yyS[yypt-0].node.(*DeclarationSpecifiers),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.DeclarationSpecifiers.attr
			lhs.typeSpecifier = lhs.DeclarationSpecifiers.typeSpecifier
		}
	case 86:
		{
			yyVAL.node = &InitDeclaratorList{
				InitDeclarator: yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 87:
		{
			yyVAL.node = &InitDeclaratorList{
				Case:               1,
				InitDeclaratorList: yyS[yypt-2].node.(*InitDeclaratorList),
				Token:              yyS[yypt-1].Token,
				InitDeclarator:     yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 88:
		{
			yyVAL.node = (*InitDeclaratorListOpt)(nil)
		}
	case 89:
		{
			yyVAL.node = &InitDeclaratorListOpt{
				InitDeclaratorList: yyS[yypt-0].node.(*InitDeclaratorList).reverse(),
			}
		}
	case 90:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
		}
	case 91:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-0].node.(*Declarator)
			d.setFull(lx)
		}
	case 92:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Case:        1,
				Declarator:  yyS[yypt-3].node.(*Declarator),
				Token:       yyS[yypt-1].Token,
				Initializer: yyS[yypt-0].node.(*Initializer),
			}
			yyVAL.node = lhs
			d := lhs.Declarator
			lhs.Initializer.typeCheck(&d.Type, d.Type, lhs.Declarator.specifier.IsStatic(), lx)
			if d.Type.Specifier().IsExtern() {
				id, _ := d.Identifier()
				lx.report.Err(d.Pos(), "'%s' initialized and declared 'extern'", dict.S(id))
			}
		}
	case 93:
		{
			lhs := &StorageClassSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saTypedef
		}
	case 94:
		{
			lhs := &StorageClassSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saExtern
		}
	case 95:
		{
			lhs := &StorageClassSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saStatic
		}
	case 96:
		{
			lhs := &StorageClassSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saAuto
		}
	case 97:
		{
			lhs := &StorageClassSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRegister
		}
	case 98:
		{
			lhs := &TypeSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsVoid)
		}
	case 99:
		{
			lhs := &TypeSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsChar)
		}
	case 100:
		{
			lhs := &TypeSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsShort)
		}
	case 101:
		{
			lhs := &TypeSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsInt)
		}
	case 102:
		{
			lhs := &TypeSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsLong)
		}
	case 103:
		{
			lhs := &TypeSpecifier{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsFloat)
		}
	case 104:
		{
			lhs := &TypeSpecifier{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsDouble)
		}
	case 105:
		{
			lhs := &TypeSpecifier{
				Case:  7,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsSigned)
		}
	case 106:
		{
			lhs := &TypeSpecifier{
				Case:  8,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsUnsigned)
		}
	case 107:
		{
			lhs := &TypeSpecifier{
				Case:  9,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsBool)
		}
	case 108:
		{
			lhs := &TypeSpecifier{
				Case:  10,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsComplex)
		}
	case 109:
		{
			lhs := &TypeSpecifier{
				Case: 11,
				StructOrUnionSpecifier: yyS[yypt-0].node.(*StructOrUnionSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(lhs.StructOrUnionSpecifier.typeSpecifiers())
		}
	case 110:
		{
			lhs := &TypeSpecifier{
				Case:          12,
				EnumSpecifier: yyS[yypt-0].node.(*EnumSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsEnumSpecifier)
		}
	case 111:
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
	case 112:
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
	case 113:
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
	case 114:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-1].node.(*IdentifierOpt); o != nil {
				lx.scope.declareStructTag(o.Token, lx.report)
			}
			lx.pushScope(ScopeMembers)
			lx.scope.isUnion = yyS[yypt-2].node.(*StructOrUnion).Case == 1 // "union"
			lx.scope.prevStructDeclarator = nil
		}
	case 115:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				StructOrUnion: yyS[yypt-5].node.(*StructOrUnion),
				IdentifierOpt: yyS[yypt-4].node.(*IdentifierOpt),
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
	case 116:
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
	case 117:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				Case:          2,
				StructOrUnion: yyS[yypt-3].node.(*StructOrUnion),
				IdentifierOpt: yyS[yypt-2].node.(*IdentifierOpt),
				Token:         yyS[yypt-1].Token,
				Token2:        yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableEmptyStructs {
				lx.report.Err(lhs.Token.Pos(), "empty structs/unions not allowed")
			}
			if o := yyS[yypt-2].node.(*IdentifierOpt); o != nil {
				lx.scope.declareStructTag(o.Token, lx.report)
			}
			lx.scope.isUnion = yyS[yypt-3].node.(*StructOrUnion).Case == 1 // "union"
			lx.scope.prevStructDeclarator = nil
			lhs.alignOf = 1
			lhs.sizeOf = 0
			if o := lhs.IdentifierOpt; o != nil {
				lx.scope.defineStructTag(o.Token, lhs, lx.report)
			}
		}
	case 118:
		{
			yyVAL.node = &StructOrUnion{
				Token: yyS[yypt-0].Token,
			}
		}
	case 119:
		{
			yyVAL.node = &StructOrUnion{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 120:
		{
			yyVAL.node = &StructDeclarationList{
				StructDeclaration: yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 121:
		{
			yyVAL.node = &StructDeclarationList{
				Case: 1,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList),
				StructDeclaration:     yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 122:
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
	case 123:
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
	case 124:
		{
			yyVAL.node = &StructDeclaration{
				Case: 2,
				StaticAssertDeclaration: yyS[yypt-0].node.(*StaticAssertDeclaration),
			}
		}
	case 125:
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
	case 126:
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
	case 127:
		{
			yyVAL.node = (*SpecifierQualifierListOpt)(nil)
		}
	case 128:
		{
			lhs := &SpecifierQualifierListOpt{
				SpecifierQualifierList: yyS[yypt-0].node.(*SpecifierQualifierList),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.SpecifierQualifierList.attr
			lhs.typeSpecifier = lhs.SpecifierQualifierList.typeSpecifier
		}
	case 129:
		{
			yyVAL.node = &StructDeclaratorList{
				StructDeclarator: yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 130:
		{
			yyVAL.node = &StructDeclaratorList{
				Case:                 1,
				StructDeclaratorList: yyS[yypt-2].node.(*StructDeclaratorList),
				Token:                yyS[yypt-1].Token,
				StructDeclarator:     yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 131:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.post(lx)
		}
	case 132:
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
	case 133:
		{
			yyVAL.node = (*CommaOpt)(nil)
		}
	case 134:
		{
			yyVAL.node = &CommaOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 135:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareEnumTag(o.Token, lx.report)
			}
			lx.iota = 0
		}
	case 136:
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
	case 137:
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
	case 138:
		{
			yyVAL.node = &EnumeratorList{
				Enumerator: yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 139:
		{
			yyVAL.node = &EnumeratorList{
				Case:           1,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList),
				Token:          yyS[yypt-1].Token,
				Enumerator:     yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 140:
		{
			lx := yylex.(*lexer)
			lhs := &Enumerator{
				EnumerationConstant: yyS[yypt-0].node.(*EnumerationConstant),
			}
			yyVAL.node = lhs
			m := lx.model
			v := m.MustConvert(lx.iota, m.IntType)
			lhs.Value = v
			lx.scope.defineEnumConst(lx, lhs.EnumerationConstant.Token, v)
		}
	case 141:
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
				if v, ok = m.enumValueToInt(e.Value); !ok {
					lx.report.Err(e.Pos(), "overflow in enumeration value: %v", e.Value)
				}
			}

			lhs.Value = v
			lx.scope.defineEnumConst(lx, lhs.EnumerationConstant.Token, v)
		}
	case 142:
		{
			lhs := &TypeQualifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saConst
		}
	case 143:
		{
			lhs := &TypeQualifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRestrict
		}
	case 144:
		{
			lhs := &TypeQualifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saVolatile
		}
	case 145:
		{
			lhs := &FunctionSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saInline
		}
	case 146:
		{
			lhs := &FunctionSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saNoreturn
		}
	case 147:
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
	case 148:
		{
			yyVAL.node = (*DeclaratorOpt)(nil)
		}
	case 149:
		{
			yyVAL.node = &DeclaratorOpt{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
		}
	case 150:
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
	case 151:
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
	case 152:
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
	case 153:
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
	case 154:
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
	case 155:
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
	case 156:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 157:
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
	case 158:
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
	case 159:
		{
			yyVAL.node = &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].node.(*TypeQualifierListOpt),
			}
		}
	case 160:
		{
			yyVAL.node = &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].node.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].node.(*Pointer),
			}
		}
	case 161:
		{
			yyVAL.node = (*PointerOpt)(nil)
		}
	case 162:
		{
			yyVAL.node = &PointerOpt{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
		}
	case 163:
		{
			lhs := &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.TypeQualifier.attr
		}
	case 164:
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
	case 165:
		{
			yyVAL.node = (*TypeQualifierListOpt)(nil)
		}
	case 166:
		{
			yyVAL.node = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].node.(*TypeQualifierList).reverse(),
			}
		}
	case 167:
		{
			lhs := &ParameterTypeList{
				ParameterList: yyS[yypt-0].node.(*ParameterList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 168:
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
	case 169:
		{
			yyVAL.node = (*ParameterTypeListOpt)(nil)
		}
	case 170:
		{
			yyVAL.node = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].node.(*ParameterTypeList),
			}
		}
	case 171:
		{
			yyVAL.node = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 172:
		{
			yyVAL.node = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].node.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 173:
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
	case 174:
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
	case 175:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 176:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 177:
		{
			yyVAL.node = (*IdentifierListOpt)(nil)
		}
	case 178:
		{
			yyVAL.node = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 179:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 180:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 181:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeBlock)
		}
	case 182:
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
	case 183:
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
	case 184:
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
	case 185:
		{
			yyVAL.node = (*AbstractDeclaratorOpt)(nil)
		}
	case 186:
		{
			yyVAL.node = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].node.(*AbstractDeclarator),
			}
		}
	case 187:
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
	case 188:
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
			nElements := -1
			if o := lhs.ExpressionOpt; o != nil {
				var err error
				if nElements, err = elements(o.Expression.eval(lx)); err != nil {
					lx.report.Err(o.Expression.Pos(), "%s", err)
				}
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
				elements:         nElements,
			}
			dd.parent = lhs.directDeclarator
		}
	case 189:
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
	case 190:
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
	case 191:
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
	case 192:
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
	case 193:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 194:
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
	case 195:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 196:
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
	case 197:
		{
			yyVAL.node = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 198:
		{
			yyVAL.node = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
		}
	case 199:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 200:
		{
			yyVAL.node = &Initializer{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 201:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Case:        2,
				Token:       yyS[yypt-2].Token,
				Token2:      yyS[yypt-1].Token,
				Initializer: yyS[yypt-0].node.(*Initializer),
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableLegacyDesignators {
				lx.report.Err(lhs.Pos(), "legacy designators not enabled")
			}
		}
	case 202:
		{
			yyVAL.node = &InitializerList{
				DesignationOpt: yyS[yypt-1].node.(*DesignationOpt),
				Initializer:    yyS[yypt-0].node.(*Initializer),
			}
		}
	case 203:
		{
			yyVAL.node = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].node.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].node.(*DesignationOpt),
				Initializer:     yyS[yypt-0].node.(*Initializer),
			}
		}
	case 204:
		{
			yyVAL.node = (*InitializerList)(nil)
		}
	case 205:
		{
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 206:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 207:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 208:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 209:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 210:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 211:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 212:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 213:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 214:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 215:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 216:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 217:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 218:
		{
			yyVAL.node = &Statement{
				Case:               6,
				AssemblerStatement: yyS[yypt-0].node.(*AssemblerStatement),
			}
		}
	case 219:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 220:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 221:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 222:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 223:
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
	case 224:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 225:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 226:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 227:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 228:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 229:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 230:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 231:
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
	case 232:
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
	case 233:
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
	case 234:
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
	case 235:
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
	case 236:
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
	case 237:
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
	case 238:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 239:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 240:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 241:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 242:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 243:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 244:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 245:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 246:
		{
			yyVAL.node = &ExternalDeclaration{
				Case: 2,
				BasicAssemblerStatement: yyS[yypt-1].node.(*BasicAssemblerStatement),
				Token: yyS[yypt-0].Token,
			}
		}
	case 247:
		{
			lx := yylex.(*lexer)
			lhs := &ExternalDeclaration{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableEmptyDeclarations {
				lx.report.Err(lhs.Pos(), "C++11 empty declarations are illegal in C99.")
			}
		}
	case 248:
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
	case 249:
		{
			yyVAL.node = &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:          yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 250:
		{
			lx := yylex.(*lexer)
			lx.scope.specifier = &DeclarationSpecifiers{typeSpecifier: tsEncode(tsInt)}
		}
	case 251:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-1].node.(*Declarator)
			if !lx.tweaks.enableOmitFuncRetType {
				lx.report.Err(d.Pos(), "missing function return type")
			}
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
	case 252:
		{
			yyVAL.node = &FunctionDefinition{
				Case:               1,
				Declarator:         yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt: yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:       yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 253:
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
	case 254:
		{
			lhs := &FunctionBody{
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
			yyVAL.node = lhs
			lhs.scope = lhs.CompoundStatement.scope
		}
	case 255:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 256:
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
	case 257:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 258:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 259:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 260:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 261:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScopePos(lhs.Pos())
		}
	case 262:
		{
			yyVAL.node = &AssemblerInstructions{
				Token: yyS[yypt-0].Token,
			}
		}
	case 263:
		{
			yyVAL.node = &AssemblerInstructions{
				Case: 1,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions),
				Token: yyS[yypt-0].Token,
			}
		}
	case 264:
		{
			yyVAL.node = &BasicAssemblerStatement{
				Token:                 yyS[yypt-4].Token,
				VolatileOpt:           yyS[yypt-3].node.(*VolatileOpt),
				Token2:                yyS[yypt-2].Token,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-0].Token,
			}
		}
	case 265:
		{
			yyVAL.node = (*VolatileOpt)(nil)
		}
	case 266:
		{
			yyVAL.node = &VolatileOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 267:
		{
			yyVAL.node = &AssemblerOperand{
				AssemblerSymbolicNameOpt: yyS[yypt-4].node.(*AssemblerSymbolicNameOpt),
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
		}
	case 268:
		{
			yyVAL.node = &AssemblerOperands{
				AssemblerOperand: yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 269:
		{
			yyVAL.node = &AssemblerOperands{
				Case:              1,
				AssemblerOperands: yyS[yypt-2].node.(*AssemblerOperands),
				Token:             yyS[yypt-1].Token,
				AssemblerOperand:  yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 270:
		{
			yyVAL.node = (*AssemblerSymbolicNameOpt)(nil)
		}
	case 271:
		{
			yyVAL.node = &AssemblerSymbolicNameOpt{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 272:
		{
			yyVAL.node = &Clobbers{
				Token: yyS[yypt-0].Token,
			}
		}
	case 273:
		{
			yyVAL.node = &Clobbers{
				Case:     1,
				Clobbers: yyS[yypt-2].node.(*Clobbers),
				Token:    yyS[yypt-1].Token,
				Token2:   yyS[yypt-0].Token,
			}
		}
	case 274:
		{
			yyVAL.node = &AssemblerStatement{
				BasicAssemblerStatement: yyS[yypt-0].node.(*BasicAssemblerStatement),
			}
		}
	case 275:
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
	case 276:
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
	case 277:
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
	case 278:
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
	case 279:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  5,
				Token:                 yyS[yypt-5].Token,
				VolatileOpt:           yyS[yypt-4].node.(*VolatileOpt),
				Token2:                yyS[yypt-3].Token,
				AssemblerInstructions: yyS[yypt-2].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-1].Token,
				Token4:                yyS[yypt-0].Token,
			}
		}
	case 280:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  6,
				Token:                 yyS[yypt-7].Token,
				VolatileOpt:           yyS[yypt-6].node.(*VolatileOpt),
				Token2:                yyS[yypt-5].Token,
				AssemblerInstructions: yyS[yypt-4].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-3].Token,
				Token4:                yyS[yypt-2].Token,
				AssemblerOperands:     yyS[yypt-1].node.(*AssemblerOperands).reverse(),
				Token5:                yyS[yypt-0].Token,
			}
		}
	case 281:
		{
			lx := yylex.(*lexer)
			lhs := &StaticAssertDeclaration{
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
	case 282:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 283:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 284:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 285:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 286:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 287:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 288:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 289:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 290:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 291:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 292:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 293:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 294:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 295:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 296:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 297:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 298:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 299:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 300:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 301:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 302:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 303:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 304:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 305:
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
	case 306:
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
	case 307:
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
	case 308:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 309:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 310:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 311:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 312:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 313:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 314:
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
	case 315:
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
	case 316:
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
	case 317:
		{
			yyVAL.node = &ControlLine{
				Case:        13,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 320:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 321:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
