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
	yyDefault           = 57460
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
	NOSEMI              = 57396
	ORASSIGN            = 57397
	OROR                = 57398
	PPDEFINE            = 57399
	PPELIF              = 57400
	PPELSE              = 57401
	PPENDIF             = 57402
	PPERROR             = 57403
	PPHASH_NL           = 57404
	PPHEADER_NAME       = 57405
	PPIF                = 57406
	PPIFDEF             = 57407
	PPIFNDEF            = 57408
	PPINCLUDE           = 57409
	PPINCLUDE_NEXT      = 57410
	PPLINE              = 57411
	PPNONDIRECTIVE      = 57412
	PPNUMBER            = 57413
	PPOTHER             = 57414
	PPPASTE             = 57415
	PPPRAGMA            = 57416
	PPUNDEF             = 57417
	PREPROCESSING_FILE  = 1048576
	REGISTER            = 57418
	RESTRICT            = 57419
	RETURN              = 57420
	RSH                 = 57421
	RSHASSIGN           = 57422
	SHORT               = 57423
	SIGNED              = 57424
	SIZEOF              = 57425
	STATIC              = 57426
	STATIC_ASSERT       = 57427
	STRINGLITERAL       = 57428
	STRUCT              = 57429
	SUBASSIGN           = 57430
	SWITCH              = 57431
	TRANSLATION_UNIT    = 1048578
	TYPEDEF             = 57432
	TYPEDEFNAME         = 57433
	TYPEOF              = 57434
	UNARY               = 57435
	UNION               = 57436
	UNSIGNED            = 57437
	VOID                = 57438
	VOLATILE            = 57439
	WHILE               = 57440
	XORASSIGN           = 57441
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -327
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (331x)
		42:      1,   // '*' (292x)
		57377:   2,   // IDENTIFIER (240x)
		38:      3,   // '&' (224x)
		43:      4,   // '+' (224x)
		45:      5,   // '-' (224x)
		57348:   6,   // ANDAND (224x)
		57363:   7,   // DEC (224x)
		57381:   8,   // INC (224x)
		59:      9,   // ';' (219x)
		41:      10,  // ')' (202x)
		44:      11,  // ',' (190x)
		57428:   12,  // STRINGLITERAL (168x)
		91:      13,  // '[' (166x)
		33:      14,  // '!' (149x)
		126:     15,  // '~' (149x)
		57347:   16,  // ALIGNOF (149x)
		57358:   17,  // CHARCONST (149x)
		57373:   18,  // FLOATCONST (149x)
		57384:   19,  // INTCONST (149x)
		57387:   20,  // LONGCHARCONST (149x)
		57388:   21,  // LONGSTRINGLITERAL (149x)
		57425:   22,  // SIZEOF (149x)
		57439:   23,  // VOLATILE (141x)
		57360:   24,  // CONST (139x)
		57419:   25,  // RESTRICT (139x)
		125:     26,  // '}' (131x)
		57353:   27,  // BOOL (129x)
		57357:   28,  // CHAR (129x)
		57359:   29,  // COMPLEX (129x)
		57367:   30,  // DOUBLE (129x)
		57369:   31,  // ENUM (129x)
		57372:   32,  // FLOAT (129x)
		57383:   33,  // INT (129x)
		57386:   34,  // LONG (129x)
		57423:   35,  // SHORT (129x)
		57424:   36,  // SIGNED (129x)
		57429:   37,  // STRUCT (129x)
		57433:   38,  // TYPEDEFNAME (129x)
		57434:   39,  // TYPEOF (129x)
		57436:   40,  // UNION (129x)
		57437:   41,  // UNSIGNED (129x)
		57438:   42,  // VOID (129x)
		58:      43,  // ':' (124x)
		57426:   44,  // STATIC (120x)
		57352:   45,  // AUTO (114x)
		57371:   46,  // EXTERN (114x)
		57382:   47,  // INLINE (114x)
		57395:   48,  // NORETURN (114x)
		57418:   49,  // REGISTER (114x)
		57432:   50,  // TYPEDEF (114x)
		57344:   51,  // $end (105x)
		61:      52,  // '=' (95x)
		123:     53,  // '{' (93x)
		57503:   54,  // Expression (87x)
		93:      55,  // ']' (85x)
		46:      56,  // '.' (84x)
		57351:   57,  // ASM (84x)
		57427:   58,  // STATIC_ASSERT (79x)
		37:      59,  // '%' (76x)
		47:      60,  // '/' (76x)
		60:      61,  // '<' (76x)
		62:      62,  // '>' (76x)
		63:      63,  // '?' (76x)
		94:      64,  // '^' (76x)
		124:     65,  // '|' (76x)
		57346:   66,  // ADDASSIGN (76x)
		57349:   67,  // ANDASSIGN (76x)
		57350:   68,  // ARROW (76x)
		57365:   69,  // DIVASSIGN (76x)
		57370:   70,  // EQ (76x)
		57375:   71,  // GEQ (76x)
		57385:   72,  // LEQ (76x)
		57389:   73,  // LSH (76x)
		57390:   74,  // LSHASSIGN (76x)
		57391:   75,  // MODASSIGN (76x)
		57392:   76,  // MULASSIGN (76x)
		57393:   77,  // NEQ (76x)
		57397:   78,  // ORASSIGN (76x)
		57398:   79,  // OROR (76x)
		57421:   80,  // RSH (76x)
		57422:   81,  // RSHASSIGN (76x)
		57430:   82,  // SUBASSIGN (76x)
		57441:   83,  // XORASSIGN (76x)
		10:      84,  // '\n' (58x)
		57376:   85,  // GOTO (55x)
		57440:   86,  // WHILE (53x)
		57354:   87,  // BREAK (52x)
		57355:   88,  // CASE (52x)
		57361:   89,  // CONTINUE (52x)
		57364:   90,  // DEFAULT (52x)
		57366:   91,  // DO (52x)
		57374:   92,  // FOR (52x)
		57380:   93,  // IF (52x)
		57414:   94,  // PPOTHER (52x)
		57420:   95,  // RETURN (52x)
		57431:   96,  // SWITCH (52x)
		57402:   97,  // PPENDIF (44x)
		57401:   98,  // PPELSE (40x)
		57400:   99,  // PPELIF (39x)
		57399:   100, // PPDEFINE (35x)
		57403:   101, // PPERROR (35x)
		57404:   102, // PPHASH_NL (35x)
		57406:   103, // PPIF (35x)
		57407:   104, // PPIFDEF (35x)
		57408:   105, // PPIFNDEF (35x)
		57409:   106, // PPINCLUDE (35x)
		57410:   107, // PPINCLUDE_NEXT (35x)
		57411:   108, // PPLINE (35x)
		57412:   109, // PPNONDIRECTIVE (35x)
		57416:   110, // PPPRAGMA (35x)
		57417:   111, // PPUNDEF (35x)
		57368:   112, // ELSE (32x)
		57555:   113, // TypeQualifier (28x)
		57504:   114, // ExpressionList (26x)
		57528:   115, // PPTokenList (22x)
		57530:   116, // PPTokens (22x)
		57499:   117, // EnumSpecifier (20x)
		57550:   118, // StructOrUnion (20x)
		57551:   119, // StructOrUnionSpecifier (20x)
		57558:   120, // TypeSpecifier (20x)
		57505:   121, // ExpressionListOpt (18x)
		57470:   122, // BasicAssemblerStatement (15x)
		57476:   123, // CompoundStatement (15x)
		57482:   124, // DeclarationSpecifiers (15x)
		57511:   125, // FunctionSpecifier (15x)
		57529:   126, // PPTokenListOpt (15x)
		57545:   127, // StorageClassSpecifier (15x)
		57468:   128, // AssemblerStatement (13x)
		57507:   129, // ExpressionStatement (12x)
		57525:   130, // IterationStatement (12x)
		57526:   131, // JumpStatement (12x)
		57527:   132, // LabeledStatement (12x)
		57535:   133, // Pointer (12x)
		57539:   134, // SelectionStatement (12x)
		57543:   135, // Statement (12x)
		57536:   136, // PointerOpt (11x)
		57484:   137, // Declarator (9x)
		57544:   138, // StaticAssertDeclaration (9x)
		57478:   139, // ControlLine (8x)
		57514:   140, // GroupPart (8x)
		57518:   141, // IfGroup (8x)
		57519:   142, // IfSection (8x)
		57552:   143, // TextLine (8x)
		57479:   144, // Declaration (7x)
		57454:   145, // $@4 (6x)
		57477:   146, // ConstantExpression (6x)
		57362:   147, // DDD (6x)
		57512:   148, // GroupList (6x)
		57466:   149, // AssemblerOperand (5x)
		57469:   150, // AssemblerSymbolicNameOpt (5x)
		57513:   151, // GroupListOpt (5x)
		57538:   152, // ReplacementList (5x)
		57540:   153, // SpecifierQualifierList (5x)
		57556:   154, // TypeQualifierList (5x)
		57459:   155, // $@9 (4x)
		57461:   156, // AbstractDeclarator (4x)
		57467:   157, // AssemblerOperands (4x)
		57483:   158, // DeclarationSpecifiersOpt (4x)
		57488:   159, // Designator (4x)
		57523:   160, // Initializer (4x)
		57531:   161, // ParameterDeclaration (4x)
		57554:   162, // TypeName (4x)
		57557:   163, // TypeQualifierListOpt (4x)
		57465:   164, // AssemblerInstructions (3x)
		57475:   165, // CommaOpt (3x)
		57486:   166, // Designation (3x)
		57487:   167, // DesignationOpt (3x)
		57489:   168, // DesignatorList (3x)
		57506:   169, // ExpressionOpt (3x)
		57515:   170, // IdentifierList (3x)
		57520:   171, // InitDeclarator (3x)
		57532:   172, // ParameterList (3x)
		57533:   173, // ParameterTypeList (3x)
		57443:   174, // $@10 (2x)
		57447:   175, // $@14 (2x)
		57449:   176, // $@16 (2x)
		57450:   177, // $@17 (2x)
		57451:   178, // $@18 (2x)
		57455:   179, // $@5 (2x)
		57462:   180, // AbstractDeclaratorOpt (2x)
		57471:   181, // BlockItem (2x)
		57474:   182, // Clobbers (2x)
		57481:   183, // DeclarationListOpt (2x)
		57485:   184, // DeclaratorOpt (2x)
		57490:   185, // DirectAbstractDeclarator (2x)
		57491:   186, // DirectAbstractDeclaratorOpt (2x)
		57492:   187, // DirectDeclarator (2x)
		57493:   188, // ElifGroup (2x)
		57500:   189, // EnumerationConstant (2x)
		57501:   190, // Enumerator (2x)
		57508:   191, // ExternalDeclaration (2x)
		57509:   192, // FunctionBody (2x)
		57510:   193, // FunctionDefinition (2x)
		57516:   194, // IdentifierListOpt (2x)
		57517:   195, // IdentifierOpt (2x)
		57521:   196, // InitDeclaratorList (2x)
		57522:   197, // InitDeclaratorListOpt (2x)
		57524:   198, // InitializerList (2x)
		57534:   199, // ParameterTypeListOpt (2x)
		57541:   200, // SpecifierQualifierListOpt (2x)
		57546:   201, // StructDeclaration (2x)
		57548:   202, // StructDeclarator (2x)
		57559:   203, // VolatileOpt (2x)
		57442:   204, // $@1 (1x)
		57444:   205, // $@11 (1x)
		57445:   206, // $@12 (1x)
		57446:   207, // $@13 (1x)
		57448:   208, // $@15 (1x)
		57452:   209, // $@2 (1x)
		57453:   210, // $@3 (1x)
		57456:   211, // $@6 (1x)
		57457:   212, // $@7 (1x)
		57458:   213, // $@8 (1x)
		57463:   214, // ArgumentExpressionList (1x)
		57464:   215, // ArgumentExpressionListOpt (1x)
		57472:   216, // BlockItemList (1x)
		57473:   217, // BlockItemListOpt (1x)
		1048577: 218, // CONSTANT_EXPRESSION (1x)
		57480:   219, // DeclarationList (1x)
		57494:   220, // ElifGroupList (1x)
		57495:   221, // ElifGroupListOpt (1x)
		57496:   222, // ElseGroup (1x)
		57497:   223, // ElseGroupOpt (1x)
		57498:   224, // EndifLine (1x)
		57502:   225, // EnumeratorList (1x)
		57378:   226, // IDENTIFIER_LPAREN (1x)
		1048576: 227, // PREPROCESSING_FILE (1x)
		57537:   228, // PreprocessingFile (1x)
		57542:   229, // Start (1x)
		57547:   230, // StructDeclarationList (1x)
		57549:   231, // StructDeclaratorList (1x)
		1048578: 232, // TRANSLATION_UNIT (1x)
		57553:   233, // TranslationUnit (1x)
		57460:   234, // $default (0x)
		57356:   235, // CAST (0x)
		57345:   236, // error (0x)
		57379:   237, // IDENTIFIER_NONREPL (0x)
		57394:   238, // NOELSE (0x)
		57396:   239, // NOSEMI (0x)
		57405:   240, // PPHEADER_NAME (0x)
		57413:   241, // PPNUMBER (0x)
		57415:   242, // PPPASTE (0x)
		57435:   243, // UNARY (0x)
	}

	yySymNames = []string{
		"'('",
		"'*'",
		"IDENTIFIER",
		"'&'",
		"'+'",
		"'-'",
		"ANDAND",
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
		"'.'",
		"ASM",
		"STATIC_ASSERT",
		"'%'",
		"'/'",
		"'<'",
		"'>'",
		"'?'",
		"'^'",
		"'|'",
		"ADDASSIGN",
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
		"WHILE",
		"BREAK",
		"CASE",
		"CONTINUE",
		"DEFAULT",
		"DO",
		"FOR",
		"IF",
		"PPOTHER",
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
		"NOSEMI",
		"PPHEADER_NAME",
		"PPNUMBER",
		"PPPASTE",
		"UNARY",
	}

	yyTokenLiteralStrings = map[int]string{
		57377:   "identifier",
		57348:   "&&",
		57363:   "--",
		57381:   "++",
		57428:   "string literal",
		57347:   "_Alignof",
		57358:   "character constant",
		57373:   "floating-point constant",
		57384:   "integer constant",
		57387:   "long character constant",
		57388:   "long string constant",
		57425:   "sizeof",
		57439:   "volatile",
		57360:   "const",
		57419:   "restrict",
		57353:   "_Bool",
		57357:   "char",
		57359:   "_Complex",
		57367:   "double",
		57369:   "enum",
		57372:   "float",
		57383:   "int",
		57386:   "long",
		57423:   "short",
		57424:   "signed",
		57429:   "struct",
		57433:   "typedefname",
		57434:   "typeof",
		57436:   "union",
		57437:   "unsigned",
		57438:   "void",
		57426:   "static",
		57352:   "auto",
		57371:   "extern",
		57382:   "inline",
		57395:   "_Noreturn",
		57418:   "register",
		57432:   "typedef",
		57351:   "asm",
		57427:   "_Static_assert",
		57346:   "+=",
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
		57397:   "|=",
		57398:   "||",
		57421:   ">>",
		57422:   ">>=",
		57430:   "-=",
		57441:   "^=",
		57376:   "goto",
		57440:   "while",
		57354:   "break",
		57355:   "case",
		57361:   "continue",
		57364:   "default",
		57366:   "do",
		57374:   "for",
		57380:   "if",
		57414:   "ppother",
		57420:   "return",
		57431:   "switch",
		57402:   "#endif",
		57401:   "#else",
		57400:   "#elif",
		57399:   "#define",
		57403:   "#error",
		57404:   "#",
		57406:   "#if",
		57407:   "#ifdef",
		57408:   "#ifndef",
		57409:   "#include",
		57410:   "#include_next",
		57411:   "#line",
		57412:   "#foo",
		57416:   "#pragma",
		57417:   "#undef",
		57368:   "else",
		57362:   "...",
		1048577: "constant expression prefix",
		57378:   "identifier immediatelly followed by '('",
		1048576: "preprocessing file prefix",
		1048578: "translation unit prefix",
		57379:   "non replaceable identifier",
		57405:   "header name",
		57413:   "preprocessing number",
		57415:   "##",
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
		12:  {54, 1},
		13:  {54, 1},
		14:  {54, 1},
		15:  {54, 1},
		16:  {54, 1},
		17:  {54, 1},
		18:  {54, 1},
		19:  {54, 3},
		20:  {54, 4},
		21:  {54, 4},
		22:  {54, 3},
		23:  {54, 3},
		24:  {54, 2},
		25:  {54, 2},
		26:  {54, 7},
		27:  {54, 2},
		28:  {54, 2},
		29:  {54, 2},
		30:  {54, 2},
		31:  {54, 2},
		32:  {54, 2},
		33:  {54, 2},
		34:  {54, 2},
		35:  {54, 2},
		36:  {54, 4},
		37:  {54, 4},
		38:  {54, 3},
		39:  {54, 3},
		40:  {54, 3},
		41:  {54, 3},
		42:  {54, 3},
		43:  {54, 3},
		44:  {54, 3},
		45:  {54, 3},
		46:  {54, 3},
		47:  {54, 3},
		48:  {54, 3},
		49:  {54, 3},
		50:  {54, 3},
		51:  {54, 3},
		52:  {54, 3},
		53:  {54, 3},
		54:  {54, 3},
		55:  {54, 3},
		56:  {54, 5},
		57:  {54, 3},
		58:  {54, 3},
		59:  {54, 3},
		60:  {54, 3},
		61:  {54, 3},
		62:  {54, 3},
		63:  {54, 3},
		64:  {54, 3},
		65:  {54, 3},
		66:  {54, 3},
		67:  {54, 3},
		68:  {54, 4},
		69:  {54, 3},
		70:  {54, 2},
		71:  {169, 0},
		72:  {169, 1},
		73:  {114, 1},
		74:  {114, 3},
		75:  {121, 0},
		76:  {121, 1},
		77:  {145, 0},
		78:  {146, 2},
		79:  {144, 3},
		80:  {144, 1},
		81:  {124, 2},
		82:  {124, 2},
		83:  {124, 2},
		84:  {124, 2},
		85:  {158, 0},
		86:  {158, 1},
		87:  {196, 1},
		88:  {196, 3},
		89:  {197, 0},
		90:  {197, 1},
		91:  {171, 1},
		92:  {179, 0},
		93:  {171, 4},
		94:  {127, 1},
		95:  {127, 1},
		96:  {127, 1},
		97:  {127, 1},
		98:  {127, 1},
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
		112: {120, 1},
		113: {120, 4},
		114: {120, 4},
		115: {211, 0},
		116: {119, 6},
		117: {119, 2},
		118: {119, 4},
		119: {118, 1},
		120: {118, 1},
		121: {230, 1},
		122: {230, 2},
		123: {201, 3},
		124: {201, 2},
		125: {201, 1},
		126: {153, 2},
		127: {153, 2},
		128: {200, 0},
		129: {200, 1},
		130: {231, 1},
		131: {231, 3},
		132: {202, 1},
		133: {202, 3},
		134: {165, 0},
		135: {165, 1},
		136: {212, 0},
		137: {117, 7},
		138: {117, 2},
		139: {225, 1},
		140: {225, 3},
		141: {190, 1},
		142: {190, 3},
		143: {113, 1},
		144: {113, 1},
		145: {113, 1},
		146: {125, 1},
		147: {125, 1},
		148: {137, 2},
		149: {184, 0},
		150: {184, 1},
		151: {187, 1},
		152: {187, 3},
		153: {187, 5},
		154: {187, 6},
		155: {187, 6},
		156: {187, 5},
		157: {213, 0},
		158: {187, 5},
		159: {187, 4},
		160: {133, 2},
		161: {133, 3},
		162: {136, 0},
		163: {136, 1},
		164: {154, 1},
		165: {154, 2},
		166: {163, 0},
		167: {163, 1},
		168: {173, 1},
		169: {173, 3},
		170: {199, 0},
		171: {199, 1},
		172: {172, 1},
		173: {172, 3},
		174: {161, 2},
		175: {161, 2},
		176: {170, 1},
		177: {170, 3},
		178: {194, 0},
		179: {194, 1},
		180: {195, 0},
		181: {195, 1},
		182: {155, 0},
		183: {162, 3},
		184: {156, 1},
		185: {156, 2},
		186: {180, 0},
		187: {180, 1},
		188: {185, 3},
		189: {185, 4},
		190: {185, 5},
		191: {185, 6},
		192: {185, 6},
		193: {185, 4},
		194: {174, 0},
		195: {185, 4},
		196: {205, 0},
		197: {185, 5},
		198: {186, 0},
		199: {186, 1},
		200: {160, 1},
		201: {160, 4},
		202: {160, 3},
		203: {198, 2},
		204: {198, 4},
		205: {198, 0},
		206: {166, 2},
		207: {167, 0},
		208: {167, 1},
		209: {168, 1},
		210: {168, 2},
		211: {159, 3},
		212: {159, 2},
		213: {135, 1},
		214: {135, 1},
		215: {135, 1},
		216: {135, 1},
		217: {135, 1},
		218: {135, 1},
		219: {135, 1},
		220: {132, 3},
		221: {132, 4},
		222: {132, 3},
		223: {206, 0},
		224: {123, 4},
		225: {216, 1},
		226: {216, 2},
		227: {217, 0},
		228: {217, 1},
		229: {181, 1},
		230: {181, 1},
		231: {129, 2},
		232: {134, 5},
		233: {134, 7},
		234: {134, 5},
		235: {130, 5},
		236: {130, 7},
		237: {130, 9},
		238: {130, 8},
		239: {131, 3},
		240: {131, 2},
		241: {131, 2},
		242: {131, 3},
		243: {131, 3},
		244: {233, 1},
		245: {233, 2},
		246: {191, 1},
		247: {191, 1},
		248: {191, 2},
		249: {191, 1},
		250: {207, 0},
		251: {193, 5},
		252: {175, 0},
		253: {208, 0},
		254: {193, 5},
		255: {176, 0},
		256: {192, 2},
		257: {177, 0},
		258: {192, 3},
		259: {219, 1},
		260: {219, 2},
		261: {183, 0},
		262: {178, 0},
		263: {183, 2},
		264: {164, 1},
		265: {164, 2},
		266: {122, 5},
		267: {203, 0},
		268: {203, 1},
		269: {149, 5},
		270: {157, 1},
		271: {157, 3},
		272: {150, 0},
		273: {150, 3},
		274: {182, 1},
		275: {182, 3},
		276: {128, 1},
		277: {128, 7},
		278: {128, 9},
		279: {128, 11},
		280: {128, 13},
		281: {128, 6},
		282: {128, 8},
		283: {138, 7},
		284: {228, 1},
		285: {148, 1},
		286: {148, 2},
		287: {151, 0},
		288: {151, 1},
		289: {140, 1},
		290: {140, 1},
		291: {140, 3},
		292: {140, 1},
		293: {142, 4},
		294: {141, 4},
		295: {141, 4},
		296: {141, 4},
		297: {220, 1},
		298: {220, 2},
		299: {221, 0},
		300: {221, 1},
		301: {188, 4},
		302: {222, 3},
		303: {223, 0},
		304: {223, 1},
		305: {224, 1},
		306: {139, 3},
		307: {139, 5},
		308: {139, 7},
		309: {139, 5},
		310: {139, 2},
		311: {139, 1},
		312: {139, 3},
		313: {139, 3},
		314: {139, 2},
		315: {139, 3},
		316: {139, 6},
		317: {139, 2},
		318: {139, 4},
		319: {139, 3},
		320: {143, 1},
		321: {152, 1},
		322: {115, 1},
		323: {126, 1},
		324: {126, 2},
		325: {116, 1},
		326: {116, 2},
	}

	yyXErrors = map[yyXError]string{
		{0, 51}:   "invalid empty input",
		{581, -1}: "expected #endif",
		{583, -1}: "expected #endif",
		{1, -1}:   "expected $end",
		{501, -1}: "expected $end",
		{503, -1}: "expected $end",
		{32, -1}:  "expected '('",
		{49, -1}:  "expected '('",
		{75, -1}:  "expected '('",
		{272, -1}: "expected '('",
		{273, -1}: "expected '('",
		{274, -1}: "expected '('",
		{276, -1}: "expected '('",
		{286, -1}: "expected '('",
		{309, -1}: "expected '('",
		{351, -1}: "expected '('",
		{433, -1}: "expected '('",
		{54, -1}:  "expected ')'",
		{79, -1}:  "expected ')'",
		{86, -1}:  "expected ')'",
		{178, -1}: "expected ')'",
		{193, -1}: "expected ')'",
		{196, -1}: "expected ')'",
		{199, -1}: "expected ')'",
		{207, -1}: "expected ')'",
		{212, -1}: "expected ')'",
		{218, -1}: "expected ')'",
		{234, -1}: "expected ')'",
		{239, -1}: "expected ')'",
		{250, -1}: "expected ')'",
		{251, -1}: "expected ')'",
		{341, -1}: "expected ')'",
		{347, -1}: "expected ')'",
		{431, -1}: "expected ')'",
		{487, -1}: "expected ')'",
		{545, -1}: "expected ')'",
		{546, -1}: "expected ')'",
		{553, -1}: "expected ')'",
		{556, -1}: "expected ')'",
		{52, -1}:  "expected ','",
		{265, -1}: "expected ':'",
		{291, -1}: "expected ':'",
		{375, -1}: "expected ':'",
		{477, -1}: "expected ':'",
		{45, -1}:  "expected ';'",
		{55, -1}:  "expected ';'",
		{271, -1}: "expected ';'",
		{278, -1}: "expected ';'",
		{279, -1}: "expected ';'",
		{328, -1}: "expected ';'",
		{337, -1}: "expected ';'",
		{339, -1}: "expected ';'",
		{345, -1}: "expected ';'",
		{354, -1}: "expected ';'",
		{378, -1}: "expected ';'",
		{446, -1}: "expected ';'",
		{385, -1}: "expected '='",
		{91, -1}:  "expected '['",
		{525, -1}: "expected '\\n'",
		{529, -1}: "expected '\\n'",
		{533, -1}: "expected '\\n'",
		{536, -1}: "expected '\\n'",
		{538, -1}: "expected '\\n'",
		{560, -1}: "expected '\\n'",
		{565, -1}: "expected '\\n'",
		{568, -1}: "expected '\\n'",
		{575, -1}: "expected '\\n'",
		{580, -1}: "expected '\\n'",
		{586, -1}: "expected '\\n'",
		{97, -1}:  "expected ']'",
		{186, -1}: "expected ']'",
		{230, -1}: "expected ']'",
		{297, -1}: "expected ']'",
		{399, -1}: "expected ']'",
		{450, -1}: "expected '{'",
		{452, -1}: "expected '{'",
		{464, -1}: "expected '{'",
		{266, -1}: "expected '}'",
		{405, -1}: "expected '}'",
		{421, -1}: "expected '}'",
		{461, -1}: "expected '}'",
		{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		{206, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{90, -1}:  "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{285, -1}: "expected assembler instructions or string literal",
		{287, -1}: "expected assembler instructions or string literal",
		{434, -1}: "expected assembler instructions or string literal",
		{299, -1}: "expected assembler operand or one of ['[', string literal]",
		{315, -1}: "expected assembler operands or one of [')', ':', '[', string literal]",
		{292, -1}: "expected assembler operands or one of ['[', string literal]",
		{318, -1}: "expected assembler operands or one of ['[', string literal]",
		{322, -1}: "expected assembler operands or one of ['[', string literal]",
		{445, -1}: "expected assembler statement or asm",
		{268, -1}: "expected block item or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{300, -1}: "expected clobbers or string literal",
		{325, -1}: "expected clobbers or string literal",
		{444, -1}: "expected compound statement or '{'",
		{64, -1}:  "expected compound statement or expression list or type name or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{247, -1}: "expected compound statement or expression list or type name or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{3, -1}:   "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{50, -1}:  "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{264, -1}: "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{396, -1}: "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{458, -1}: "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{478, -1}: "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{500, -1}: "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{438, -1}: "expected declaration list or one of [_Bool, _Complex, _Noreturn, _Static_assert, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{440, -1}: "expected declaration or one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{336, -1}: "expected declaration or optional expression list or one of [&&, '!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{47, -1}:  "expected declarator or one of ['(', '*', identifier]",
		{384, -1}: "expected declarator or one of ['(', '*', identifier]",
		{198, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		{393, -1}: "expected designator or one of ['.', '=', '[']",
		{201, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		{87, -1}:  "expected direct abstract declarator or one of ['(', '[']",
		{382, -1}: "expected direct declarator or one of ['(', identifier]",
		{573, -1}: "expected elif group or one of [#elif, #else, #endif]",
		{579, -1}: "expected endif line or #endif",
		{510, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		{571, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		{453, -1}: "expected enumerator list or identifier",
		{460, -1}: "expected enumerator or one of ['}', identifier]",
		{102, -1}: "expected expression list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{126, -1}: "expected expression list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{352, -1}: "expected expression list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{356, -1}: "expected expression list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{360, -1}: "expected expression list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{364, -1}: "expected expression list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{417, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		{94, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{229, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{432, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{51, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{66, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{67, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{68, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{69, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{70, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{71, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{72, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{73, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{74, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{100, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{108, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{109, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{110, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{111, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{112, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{113, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{114, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{115, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{116, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{117, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{118, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{119, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{120, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{121, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{122, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{123, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{124, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{125, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{127, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{128, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{129, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{130, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{131, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{132, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{133, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{134, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{135, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{136, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{137, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{152, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{153, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{180, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{187, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{223, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{226, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{277, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{310, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{98, -1}:  "expected expression or optional type qualifier list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		{221, -1}: "expected expression or optional type qualifier list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		{485, -1}: "expected expression or type name or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{6, -1}:   "expected external declaration or one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{437, -1}: "expected function body or one of ['{', asm]",
		{442, -1}: "expected function body or one of ['{', asm]",
		{496, -1}: "expected function body or one of ['{', asm]",
		{497, -1}: "expected function body or one of ['{', asm]",
		{495, -1}: "expected function body or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{436, -1}: "expected function body or optional declaration list or one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{562, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{504, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{76, -1}:  "expected identifier",
		{104, -1}: "expected identifier",
		{105, -1}: "expected identifier",
		{215, -1}: "expected identifier",
		{296, -1}: "expected identifier",
		{397, -1}: "expected identifier",
		{512, -1}: "expected identifier",
		{513, -1}: "expected identifier",
		{520, -1}: "expected identifier",
		{304, -1}: "expected identifier list or identifier",
		{542, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		{411, -1}: "expected init declarator or one of ['(', '*', identifier]",
		{390, -1}: "expected initializer list or optional comma or one of [&&, '!', '&', '(', '*', '+', ',', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{419, -1}: "expected initializer list or optional comma or one of [&&, '!', '&', '(', '*', '+', ',', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{386, -1}: "expected initializer or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{392, -1}: "expected initializer or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{407, -1}: "expected initializer or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{409, -1}: "expected initializer or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{404, -1}: "expected initializer or optional designation or one of [&&, '!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{57, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{58, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{59, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{60, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{61, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{62, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{63, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{77, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{106, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{107, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{139, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{140, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{141, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{142, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{143, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{144, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{145, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{146, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{147, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{148, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{149, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
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
		{174, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{175, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{179, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{183, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{191, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{246, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{248, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{416, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{418, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{422, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{423, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{424, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{425, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{426, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{427, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{428, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{429, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{430, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{65, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{150, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{154, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{176, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{181, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{311, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{486, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{387, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{254, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{388, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{332, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{333, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{93, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{101, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{188, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{224, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{227, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{505, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{506, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{507, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{509, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{516, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{522, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{524, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{527, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{530, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{532, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{534, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{535, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{537, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{539, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{540, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{543, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{548, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{549, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{551, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{555, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{558, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{559, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{564, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{584, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{585, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{587, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{563, -1}: "expected one of [#elif, #else, #endif]",
		{567, -1}: "expected one of [#elif, #else, #endif]",
		{570, -1}: "expected one of [#elif, #else, #endif]",
		{572, -1}: "expected one of [#elif, #else, #endif]",
		{577, -1}: "expected one of [#elif, #else, #endif]",
		{578, -1}: "expected one of [#elif, #else, #endif]",
		{372, -1}: "expected one of [$end, &&, '!', '&', '(', ')', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{8, -1}:   "expected one of [$end, &&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{56, -1}:  "expected one of [$end, &&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{413, -1}: "expected one of [$end, &&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{42, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{43, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{44, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{46, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{443, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{447, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{448, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{449, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{498, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{499, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{37, -1}:  "expected one of [&&, '!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{38, -1}:  "expected one of [&&, '!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{39, -1}:  "expected one of [&&, '!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{95, -1}:  "expected one of [&&, '!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{184, -1}: "expected one of [&&, '!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{257, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{258, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{259, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{260, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{261, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{262, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{263, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{282, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{306, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{314, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{317, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{320, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{321, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{324, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{327, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{329, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{330, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{331, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{334, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{335, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{343, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{349, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{355, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{359, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{363, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{367, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{369, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{370, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{374, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{377, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{415, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{267, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{269, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{270, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{371, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{394, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{401, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{451, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{465, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
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
		{462, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{468, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{483, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{488, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{489, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{17, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{40, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{41, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{490, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{491, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{492, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{493, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{494, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{243, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		{244, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		{245, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		{204, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{205, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{208, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{217, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{219, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{225, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{228, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{231, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{232, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{85, -1}:  "expected one of ['(', ')', ',', '[', identifier]",
		{242, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		{89, -1}:  "expected one of ['(', ')', ',', '[']",
		{138, -1}: "expected one of ['(', ')', ',', '[']",
		{185, -1}: "expected one of ['(', ')', ',', '[']",
		{189, -1}: "expected one of ['(', ')', ',', '[']",
		{190, -1}: "expected one of ['(', ')', ',', '[']",
		{192, -1}: "expected one of ['(', ')', ',', '[']",
		{200, -1}: "expected one of ['(', ')', ',', '[']",
		{236, -1}: "expected one of ['(', ')', ',', '[']",
		{240, -1}: "expected one of ['(', ')', ',', '[']",
		{283, -1}: "expected one of ['(', goto]",
		{284, -1}: "expected one of ['(', goto]",
		{383, -1}: "expected one of ['(', identifier]",
		{294, -1}: "expected one of [')', ',', ':']",
		{301, -1}: "expected one of [')', ',', ':']",
		{307, -1}: "expected one of [')', ',', ':']",
		{308, -1}: "expected one of [')', ',', ':']",
		{312, -1}: "expected one of [')', ',', ':']",
		{316, -1}: "expected one of [')', ',', ':']",
		{323, -1}: "expected one of [')', ',', ':']",
		{255, -1}: "expected one of [')', ',', ';']",
		{213, -1}: "expected one of [')', ',', ...]",
		{216, -1}: "expected one of [')', ',', ...]",
		{544, -1}: "expected one of [')', ',', ...]",
		{88, -1}:  "expected one of [')', ',']",
		{177, -1}: "expected one of [')', ',']",
		{195, -1}: "expected one of [')', ',']",
		{197, -1}: "expected one of [')', ',']",
		{202, -1}: "expected one of [')', ',']",
		{203, -1}: "expected one of [')', ',']",
		{214, -1}: "expected one of [')', ',']",
		{235, -1}: "expected one of [')', ',']",
		{249, -1}: "expected one of [')', ',']",
		{305, -1}: "expected one of [')', ',']",
		{319, -1}: "expected one of [')', ',']",
		{326, -1}: "expected one of [')', ',']",
		{353, -1}: "expected one of [')', ',']",
		{357, -1}: "expected one of [')', ',']",
		{361, -1}: "expected one of [')', ',']",
		{365, -1}: "expected one of [')', ',']",
		{288, -1}: "expected one of [')', ':', string literal]",
		{290, -1}: "expected one of [')', ':', string literal]",
		{313, -1}: "expected one of [')', ':', string literal]",
		{435, -1}: "expected one of [')', string literal]",
		{476, -1}: "expected one of [',', ':', ';']",
		{151, -1}: "expected one of [',', ':']",
		{295, -1}: "expected one of [',', ':']",
		{302, -1}: "expected one of [',', ':']",
		{381, -1}: "expected one of [',', ';', '=']",
		{406, -1}: "expected one of [',', ';', '}']",
		{410, -1}: "expected one of [',', ';', '}']",
		{379, -1}: "expected one of [',', ';']",
		{380, -1}: "expected one of [',', ';']",
		{389, -1}: "expected one of [',', ';']",
		{412, -1}: "expected one of [',', ';']",
		{473, -1}: "expected one of [',', ';']",
		{475, -1}: "expected one of [',', ';']",
		{479, -1}: "expected one of [',', ';']",
		{482, -1}: "expected one of [',', ';']",
		{454, -1}: "expected one of [',', '=', '}']",
		{457, -1}: "expected one of [',', '=', '}']",
		{182, -1}: "expected one of [',', ']']",
		{403, -1}: "expected one of [',', '}']",
		{408, -1}: "expected one of [',', '}']",
		{456, -1}: "expected one of [',', '}']",
		{459, -1}: "expected one of [',', '}']",
		{463, -1}: "expected one of [',', '}']",
		{395, -1}: "expected one of ['.', '=', '[']",
		{398, -1}: "expected one of ['.', '=', '[']",
		{400, -1}: "expected one of ['.', '=', '[']",
		{402, -1}: "expected one of ['.', '=', '[']",
		{289, -1}: "expected one of [':', string literal]",
		{514, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		{523, -1}: "expected one of ['\\n', ppother]",
		{526, -1}: "expected one of ['\\n', ppother]",
		{528, -1}: "expected one of ['\\n', ppother]",
		{439, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{441, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{34, -1}:  "expected one of ['{', identifier]",
		{35, -1}:  "expected one of ['{', identifier]",
		{470, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{472, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{474, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{480, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{484, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{552, -1}: "expected one of [..., identifier]",
		{83, -1}:  "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		{103, -1}: "expected optional argument expression list or one of [&&, '!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{252, -1}: "expected optional block item list or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{253, -1}: "expected optional block item list or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{391, -1}: "expected optional comma or one of [',', '}']",
		{420, -1}: "expected optional comma or one of [',', '}']",
		{455, -1}: "expected optional comma or one of [',', '}']",
		{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{12, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{340, -1}: "expected optional expression list or one of [&&, '!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{346, -1}: "expected optional expression list or one of [&&, '!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{280, -1}: "expected optional expression list or one of [&&, '!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{338, -1}: "expected optional expression list or one of [&&, '!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{344, -1}: "expected optional expression list or one of [&&, '!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{220, -1}: "expected optional expression or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{209, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{92, -1}:  "expected optional expression or type qualifier list or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{96, -1}:  "expected optional expression or type qualifier or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{561, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{566, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{569, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{576, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{582, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{210, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{33, -1}:  "expected optional identifier or one of ['{', identifier]",
		{36, -1}:  "expected optional identifier or one of ['{', identifier]",
		{256, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		{194, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{237, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{238, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{81, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{82, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{515, -1}: "expected optional token list or one of ['\\n', ppother]",
		{519, -1}: "expected optional token list or one of ['\\n', ppother]",
		{84, -1}:  "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		{281, -1}: "expected optional volatile or one of ['(', goto, volatile]",
		{48, -1}:  "expected optional volatile or one of ['(', volatile]",
		{233, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{211, -1}: "expected parameter type list or one of [_Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{241, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{502, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{541, -1}: "expected replacement list or one of ['\\n', ppother]",
		{547, -1}: "expected replacement list or one of ['\\n', ppother]",
		{550, -1}: "expected replacement list or one of ['\\n', ppother]",
		{554, -1}: "expected replacement list or one of ['\\n', ppother]",
		{557, -1}: "expected replacement list or one of ['\\n', ppother]",
		{80, -1}:  "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{275, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{342, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{348, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{358, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{362, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{366, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{368, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{373, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{376, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{414, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{53, -1}:  "expected string literal",
		{293, -1}: "expected string literal",
		{298, -1}: "expected string literal",
		{303, -1}: "expected string literal",
		{466, -1}: "expected struct declaration list or one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{467, -1}: "expected struct declaration list or one of [_Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{469, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{471, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		{481, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		{531, -1}: "expected token list or one of ['\\n', ppother]",
		{508, -1}: "expected token list or ppother",
		{511, -1}: "expected token list or ppother",
		{517, -1}: "expected token list or ppother",
		{518, -1}: "expected token list or ppother",
		{521, -1}: "expected token list or ppother",
		{574, -1}: "expected token list or ppother",
		{4, -1}:   "expected translation unit or one of ['(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{5, -1}:   "expected translation unit or one of ['(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{78, -1}:  "expected type name or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{99, -1}:  "expected type qualifier or one of [&&, '!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		{222, -1}: "expected type qualifier or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{350, -1}: "expected while",
		{3, 51}:   "unexpected EOF",
		{2, 51}:   "unexpected EOF",
		{4, 51}:   "unexpected EOF",
	}

	yyParseTab = [588][]uint16{
		// 0
		{218: 330, 227: 329, 229: 328, 232: 331},
		{51: 327},
		{84: 326, 94: 326, 100: 326, 326, 326, 326, 326, 326, 326, 326, 326, 326, 326, 326, 204: 829},
		{324, 324, 324, 324, 324, 324, 324, 324, 324, 12: 324, 14: 324, 324, 324, 324, 324, 324, 324, 324, 324, 209: 827},
		{322, 322, 322, 9: 322, 23: 322, 322, 322, 27: 322, 322, 322, 322, 322, 322, 322, 322, 322, 322, 322, 322, 322, 322, 322, 322, 44: 322, 322, 322, 322, 322, 322, 322, 57: 322, 322, 210: 332},
		// 5
		{75, 75, 75, 9: 373, 23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 57: 375, 376, 113: 338, 117: 357, 360, 356, 337, 122: 372, 124: 334, 339, 127: 336, 138: 335, 144: 371, 175: 374, 191: 369, 193: 370, 233: 333},
		{75, 75, 75, 9: 373, 23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 321, 57: 375, 376, 113: 338, 117: 357, 360, 356, 337, 122: 372, 124: 334, 339, 127: 336, 138: 335, 144: 371, 175: 374, 191: 826, 193: 370},
		{165, 411, 165, 9: 238, 133: 710, 136: 709, 822, 171: 706, 196: 707, 705},
		{247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 12: 247, 14: 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 44: 247, 247, 247, 247, 247, 247, 247, 247, 53: 247, 57: 247, 247, 85: 247, 247, 247, 247, 247, 247, 247, 247, 247, 95: 247, 247},
		{242, 242, 242, 9: 242, 242, 242, 13: 242, 23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 113: 338, 117: 357, 360, 356, 337, 124: 818, 339, 127: 336, 158: 821},
		// 10
		{242, 242, 242, 9: 242, 242, 242, 13: 242, 23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 113: 338, 117: 357, 360, 356, 337, 124: 818, 339, 127: 336, 158: 820},
		{242, 242, 242, 9: 242, 242, 242, 13: 242, 23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 113: 338, 117: 357, 360, 356, 337, 124: 818, 339, 127: 336, 158: 819},
		{242, 242, 242, 9: 242, 242, 242, 13: 242, 23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 113: 338, 117: 357, 360, 356, 337, 124: 818, 339, 127: 336, 158: 817},
		{233, 233, 233, 9: 233, 233, 233, 13: 233, 23: 233, 233, 233, 27: 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 44: 233, 233, 233, 233, 233, 233, 233},
		{232, 232, 232, 9: 232, 232, 232, 13: 232, 23: 232, 232, 232, 27: 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 44: 232, 232, 232, 232, 232, 232, 232},
		// 15
		{231, 231, 231, 9: 231, 231, 231, 13: 231, 23: 231, 231, 231, 27: 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 44: 231, 231, 231, 231, 231, 231, 231},
		{230, 230, 230, 9: 230, 230, 230, 13: 230, 23: 230, 230, 230, 27: 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 44: 230, 230, 230, 230, 230, 230, 230},
		{229, 229, 229, 9: 229, 229, 229, 13: 229, 23: 229, 229, 229, 27: 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 44: 229, 229, 229, 229, 229, 229, 229},
		{228, 228, 228, 9: 228, 228, 228, 13: 228, 23: 228, 228, 228, 27: 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228},
		{227, 227, 227, 9: 227, 227, 227, 13: 227, 23: 227, 227, 227, 27: 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227},
		// 20
		{226, 226, 226, 9: 226, 226, 226, 13: 226, 23: 226, 226, 226, 27: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226},
		{225, 225, 225, 9: 225, 225, 225, 13: 225, 23: 225, 225, 225, 27: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225},
		{224, 224, 224, 9: 224, 224, 224, 13: 224, 23: 224, 224, 224, 27: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224},
		{223, 223, 223, 9: 223, 223, 223, 13: 223, 23: 223, 223, 223, 27: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223},
		{222, 222, 222, 9: 222, 222, 222, 13: 222, 23: 222, 222, 222, 27: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222},
		// 25
		{221, 221, 221, 9: 221, 221, 221, 13: 221, 23: 221, 221, 221, 27: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221},
		{220, 220, 220, 9: 220, 220, 220, 13: 220, 23: 220, 220, 220, 27: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220},
		{219, 219, 219, 9: 219, 219, 219, 13: 219, 23: 219, 219, 219, 27: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219},
		{218, 218, 218, 9: 218, 218, 218, 13: 218, 23: 218, 218, 218, 27: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218},
		{217, 217, 217, 9: 217, 217, 217, 13: 217, 23: 217, 217, 217, 27: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217},
		// 30
		{216, 216, 216, 9: 216, 216, 216, 13: 216, 23: 216, 216, 216, 27: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216},
		{215, 215, 215, 9: 215, 215, 215, 13: 215, 23: 215, 215, 215, 27: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215},
		{812},
		{2: 792, 53: 147, 195: 791},
		{2: 208, 53: 208},
		// 35
		{2: 207, 53: 207},
		{2: 778, 53: 147, 195: 777},
		{184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 27: 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 55: 184},
		{183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 27: 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 55: 183},
		{182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 27: 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 55: 182},
		// 40
		{181, 181, 181, 9: 181, 181, 181, 13: 181, 23: 181, 181, 181, 27: 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 44: 181, 181, 181, 181, 181, 181, 181},
		{180, 180, 180, 9: 180, 180, 180, 13: 180, 23: 180, 180, 180, 27: 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 44: 180, 180, 180, 180, 180, 180, 180},
		{83, 83, 83, 9: 83, 23: 83, 83, 83, 27: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 44: 83, 83, 83, 83, 83, 83, 83, 83, 57: 83, 83},
		{81, 81, 81, 9: 81, 23: 81, 81, 81, 27: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 44: 81, 81, 81, 81, 81, 81, 81, 81, 57: 81, 81},
		{80, 80, 80, 9: 80, 23: 80, 80, 80, 27: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 44: 80, 80, 80, 80, 80, 80, 80, 80, 57: 80, 80},
		// 45
		{9: 776},
		{78, 78, 78, 9: 78, 23: 78, 78, 78, 27: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 44: 78, 78, 78, 78, 78, 78, 78, 78, 57: 78, 78},
		{165, 411, 165, 133: 710, 136: 709, 763},
		{60, 23: 611, 203: 760},
		{377},
		// 50
		{250, 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 14: 250, 250, 250, 250, 250, 250, 250, 250, 250, 145: 378, 379},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 392},
		{11: 380},
		{12: 381},
		{10: 382},
		// 55
		{9: 383},
		{44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 12: 44, 14: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44: 44, 44, 44, 44, 44, 44, 44, 44, 53: 44, 57: 44, 44, 85: 44, 44, 44, 44, 44, 44, 44, 44, 44, 95: 44, 44},
		{315, 315, 3: 315, 315, 315, 315, 315, 315, 315, 315, 315, 13: 315, 26: 315, 43: 315, 51: 315, 315, 55: 315, 315, 59: 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315},
		{314, 314, 3: 314, 314, 314, 314, 314, 314, 314, 314, 314, 13: 314, 26: 314, 43: 314, 51: 314, 314, 55: 314, 314, 59: 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314},
		{313, 313, 3: 313, 313, 313, 313, 313, 313, 313, 313, 313, 13: 313, 26: 313, 43: 313, 51: 313, 313, 55: 313, 313, 59: 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313},
		// 60
		{312, 312, 3: 312, 312, 312, 312, 312, 312, 312, 312, 312, 13: 312, 26: 312, 43: 312, 51: 312, 312, 55: 312, 312, 59: 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312},
		{311, 311, 3: 311, 311, 311, 311, 311, 311, 311, 311, 311, 13: 311, 26: 311, 43: 311, 51: 311, 311, 55: 311, 311, 59: 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311},
		{310, 310, 3: 310, 310, 310, 310, 310, 310, 310, 310, 310, 13: 310, 26: 310, 43: 310, 51: 310, 310, 55: 310, 310, 59: 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310},
		{309, 309, 3: 309, 309, 309, 309, 309, 309, 309, 309, 309, 13: 309, 26: 309, 43: 309, 51: 309, 309, 55: 309, 309, 59: 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 145, 145, 145, 27: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 53: 579, 477, 114: 576, 123: 578, 155: 407, 162: 758},
		// 65
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 249, 11: 249, 13: 429, 26: 249, 43: 249, 51: 249, 454, 55: 249, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 757},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 756},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 755},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 518},
		// 70
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 754},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 753},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 752},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 751},
		{574, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 575},
		// 75
		{405},
		{2: 404},
		{257, 257, 3: 257, 257, 257, 257, 257, 257, 257, 257, 257, 13: 257, 26: 257, 43: 257, 51: 257, 257, 55: 257, 257, 59: 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257},
		{23: 145, 145, 145, 27: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 155: 407, 162: 406},
		{10: 573},
		// 80
		{23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 113: 409, 117: 357, 360, 356, 408, 153: 410},
		{199, 199, 199, 9: 199, 199, 13: 199, 23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 199, 113: 409, 117: 357, 360, 356, 408, 153: 571, 200: 572},
		{199, 199, 199, 9: 199, 199, 13: 199, 23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 199, 113: 409, 117: 357, 360, 356, 408, 153: 571, 200: 570},
		{165, 411, 10: 141, 13: 165, 133: 412, 136: 414, 156: 415, 180: 413},
		{161, 161, 161, 10: 161, 161, 13: 161, 23: 366, 364, 365, 113: 422, 154: 426, 163: 568},
		// 85
		{164, 2: 164, 10: 143, 143, 13: 164},
		{10: 144},
		{417, 13: 129, 185: 416, 418},
		{10: 140, 140},
		{564, 10: 142, 142, 13: 128},
		// 90
		{165, 411, 10: 133, 13: 165, 23: 133, 133, 133, 27: 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 44: 133, 133, 133, 133, 133, 133, 133, 133: 412, 136: 414, 156: 520, 174: 521},
		{13: 419},
		{391, 421, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 366, 364, 365, 44: 425, 54: 420, 256, 113: 422, 154: 423, 169: 424},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 13: 429, 52: 454, 55: 255, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 518, 519},
		// 95
		{163, 163, 163, 163, 163, 163, 163, 163, 163, 10: 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 44: 163, 55: 163},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 366, 364, 365, 44: 514, 54: 420, 256, 113: 511, 169: 513},
		{55: 512},
		{161, 161, 161, 161, 161, 161, 161, 161, 161, 12: 161, 14: 161, 161, 161, 161, 161, 161, 161, 161, 161, 366, 364, 365, 113: 422, 154: 426, 163: 427},
		{160, 160, 160, 160, 160, 160, 160, 160, 160, 10: 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 366, 364, 365, 113: 511},
		// 100
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 428},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 13: 429, 52: 454, 55: 465, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 477, 114: 509},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 10: 317, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 503, 214: 504, 505},
		{2: 502},
		// 105
		{2: 501},
		{303, 303, 3: 303, 303, 303, 303, 303, 303, 303, 303, 303, 13: 303, 26: 303, 43: 303, 51: 303, 303, 55: 303, 303, 59: 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303},
		{302, 302, 3: 302, 302, 302, 302, 302, 302, 302, 302, 302, 13: 302, 26: 302, 43: 302, 51: 302, 302, 55: 302, 302, 59: 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 500},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 499},
		// 110
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 498},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 497},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 496},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 495},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 494},
		// 115
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 493},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 492},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 491},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 490},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 489},
		// 120
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 488},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 487},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 486},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 485},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 484},
		// 125
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 483},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 477, 114: 478},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 476},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 475},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 474},
		// 130
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 473},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 472},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 471},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 470},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 469},
		// 135
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 468},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 467},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 466},
		{136, 10: 136, 136, 13: 136},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 260, 260, 260, 13: 429, 26: 260, 43: 260, 51: 260, 454, 55: 260, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		// 140
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 261, 261, 261, 13: 429, 26: 261, 43: 261, 51: 261, 454, 55: 261, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 262, 262, 262, 13: 429, 26: 262, 43: 262, 51: 262, 454, 55: 262, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 263, 263, 263, 13: 429, 26: 263, 43: 263, 51: 263, 454, 55: 263, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 264, 264, 264, 13: 429, 26: 264, 43: 264, 51: 264, 454, 55: 264, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 265, 265, 265, 13: 429, 26: 265, 43: 265, 51: 265, 454, 55: 265, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		// 145
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 266, 266, 266, 13: 429, 26: 266, 43: 266, 51: 266, 454, 55: 266, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 267, 267, 267, 13: 429, 26: 267, 43: 267, 51: 267, 454, 55: 267, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 268, 268, 268, 13: 429, 26: 268, 43: 268, 51: 268, 454, 55: 268, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 269, 269, 269, 13: 429, 26: 269, 43: 269, 51: 269, 454, 55: 269, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 270, 270, 270, 13: 429, 26: 270, 43: 270, 51: 270, 454, 55: 270, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		// 150
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 254, 254, 254, 13: 429, 43: 254, 52: 454, 55: 254, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{11: 480, 43: 479},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 482},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 481},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 253, 253, 253, 13: 429, 43: 253, 52: 454, 55: 253, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		// 155
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 271, 271, 271, 13: 429, 26: 271, 43: 271, 51: 271, 271, 55: 271, 431, 59: 437, 436, 442, 443, 453, 449, 450, 271, 271, 432, 271, 446, 445, 444, 440, 271, 271, 271, 447, 271, 452, 441, 271, 271, 271},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 272, 272, 272, 13: 429, 26: 272, 43: 272, 51: 272, 272, 55: 272, 431, 59: 437, 436, 442, 443, 272, 449, 450, 272, 272, 432, 272, 446, 445, 444, 440, 272, 272, 272, 447, 272, 272, 441, 272, 272, 272},
		{430, 435, 3: 448, 438, 439, 273, 434, 433, 273, 273, 273, 13: 429, 26: 273, 43: 273, 51: 273, 273, 55: 273, 431, 59: 437, 436, 442, 443, 273, 449, 450, 273, 273, 432, 273, 446, 445, 444, 440, 273, 273, 273, 447, 273, 273, 441, 273, 273, 273},
		{430, 435, 3: 448, 438, 439, 274, 434, 433, 274, 274, 274, 13: 429, 26: 274, 43: 274, 51: 274, 274, 55: 274, 431, 59: 437, 436, 442, 443, 274, 449, 274, 274, 274, 432, 274, 446, 445, 444, 440, 274, 274, 274, 447, 274, 274, 441, 274, 274, 274},
		{430, 435, 3: 448, 438, 439, 275, 434, 433, 275, 275, 275, 13: 429, 26: 275, 43: 275, 51: 275, 275, 55: 275, 431, 59: 437, 436, 442, 443, 275, 275, 275, 275, 275, 432, 275, 446, 445, 444, 440, 275, 275, 275, 447, 275, 275, 441, 275, 275, 275},
		// 160
		{430, 435, 3: 276, 438, 439, 276, 434, 433, 276, 276, 276, 13: 429, 26: 276, 43: 276, 51: 276, 276, 55: 276, 431, 59: 437, 436, 442, 443, 276, 276, 276, 276, 276, 432, 276, 446, 445, 444, 440, 276, 276, 276, 447, 276, 276, 441, 276, 276, 276},
		{430, 435, 3: 277, 438, 439, 277, 434, 433, 277, 277, 277, 13: 429, 26: 277, 43: 277, 51: 277, 277, 55: 277, 431, 59: 437, 436, 442, 443, 277, 277, 277, 277, 277, 432, 277, 277, 445, 444, 440, 277, 277, 277, 277, 277, 277, 441, 277, 277, 277},
		{430, 435, 3: 278, 438, 439, 278, 434, 433, 278, 278, 278, 13: 429, 26: 278, 43: 278, 51: 278, 278, 55: 278, 431, 59: 437, 436, 442, 443, 278, 278, 278, 278, 278, 432, 278, 278, 445, 444, 440, 278, 278, 278, 278, 278, 278, 441, 278, 278, 278},
		{430, 435, 3: 279, 438, 439, 279, 434, 433, 279, 279, 279, 13: 429, 26: 279, 43: 279, 51: 279, 279, 55: 279, 431, 59: 437, 436, 279, 279, 279, 279, 279, 279, 279, 432, 279, 279, 279, 279, 440, 279, 279, 279, 279, 279, 279, 441, 279, 279, 279},
		{430, 435, 3: 280, 438, 439, 280, 434, 433, 280, 280, 280, 13: 429, 26: 280, 43: 280, 51: 280, 280, 55: 280, 431, 59: 437, 436, 280, 280, 280, 280, 280, 280, 280, 432, 280, 280, 280, 280, 440, 280, 280, 280, 280, 280, 280, 441, 280, 280, 280},
		// 165
		{430, 435, 3: 281, 438, 439, 281, 434, 433, 281, 281, 281, 13: 429, 26: 281, 43: 281, 51: 281, 281, 55: 281, 431, 59: 437, 436, 281, 281, 281, 281, 281, 281, 281, 432, 281, 281, 281, 281, 440, 281, 281, 281, 281, 281, 281, 441, 281, 281, 281},
		{430, 435, 3: 282, 438, 439, 282, 434, 433, 282, 282, 282, 13: 429, 26: 282, 43: 282, 51: 282, 282, 55: 282, 431, 59: 437, 436, 282, 282, 282, 282, 282, 282, 282, 432, 282, 282, 282, 282, 440, 282, 282, 282, 282, 282, 282, 441, 282, 282, 282},
		{430, 435, 3: 283, 438, 439, 283, 434, 433, 283, 283, 283, 13: 429, 26: 283, 43: 283, 51: 283, 283, 55: 283, 431, 59: 437, 436, 283, 283, 283, 283, 283, 283, 283, 432, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283},
		{430, 435, 3: 284, 438, 439, 284, 434, 433, 284, 284, 284, 13: 429, 26: 284, 43: 284, 51: 284, 284, 55: 284, 431, 59: 437, 436, 284, 284, 284, 284, 284, 284, 284, 432, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284},
		{430, 435, 3: 285, 285, 285, 285, 434, 433, 285, 285, 285, 13: 429, 26: 285, 43: 285, 51: 285, 285, 55: 285, 431, 59: 437, 436, 285, 285, 285, 285, 285, 285, 285, 432, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285},
		// 170
		{430, 435, 3: 286, 286, 286, 286, 434, 433, 286, 286, 286, 13: 429, 26: 286, 43: 286, 51: 286, 286, 55: 286, 431, 59: 437, 436, 286, 286, 286, 286, 286, 286, 286, 432, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286},
		{430, 287, 3: 287, 287, 287, 287, 434, 433, 287, 287, 287, 13: 429, 26: 287, 43: 287, 51: 287, 287, 55: 287, 431, 59: 287, 287, 287, 287, 287, 287, 287, 287, 287, 432, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287},
		{430, 288, 3: 288, 288, 288, 288, 434, 433, 288, 288, 288, 13: 429, 26: 288, 43: 288, 51: 288, 288, 55: 288, 431, 59: 288, 288, 288, 288, 288, 288, 288, 288, 288, 432, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288},
		{430, 289, 3: 289, 289, 289, 289, 434, 433, 289, 289, 289, 13: 429, 26: 289, 43: 289, 51: 289, 289, 55: 289, 431, 59: 289, 289, 289, 289, 289, 289, 289, 289, 289, 432, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289},
		{304, 304, 3: 304, 304, 304, 304, 304, 304, 304, 304, 304, 13: 304, 26: 304, 43: 304, 51: 304, 304, 55: 304, 304, 59: 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304},
		// 175
		{305, 305, 3: 305, 305, 305, 305, 305, 305, 305, 305, 305, 13: 305, 26: 305, 43: 305, 51: 305, 305, 55: 305, 305, 59: 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 10: 319, 319, 13: 429, 52: 454, 56: 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{10: 316, 507},
		{10: 506},
		{306, 306, 3: 306, 306, 306, 306, 306, 306, 306, 306, 306, 13: 306, 26: 306, 43: 306, 51: 306, 306, 55: 306, 306, 59: 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306},
		// 180
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 508},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 10: 318, 318, 13: 429, 52: 454, 56: 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{11: 480, 55: 510},
		{307, 307, 3: 307, 307, 307, 307, 307, 307, 307, 307, 307, 13: 307, 26: 307, 43: 307, 51: 307, 307, 55: 307, 307, 59: 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307},
		{162, 162, 162, 162, 162, 162, 162, 162, 162, 10: 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 44: 162, 55: 162},
		// 185
		{138, 10: 138, 138, 13: 138},
		{55: 517},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 515},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 13: 429, 52: 454, 55: 516, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{135, 10: 135, 135, 13: 135},
		// 190
		{137, 10: 137, 137, 13: 137},
		{430, 297, 3: 297, 297, 297, 297, 434, 433, 297, 297, 297, 13: 429, 26: 297, 43: 297, 51: 297, 297, 55: 297, 431, 59: 297, 297, 297, 297, 297, 297, 297, 297, 297, 432, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297},
		{134, 10: 134, 134, 13: 134},
		{10: 563},
		{10: 157, 23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 113: 338, 117: 357, 360, 356, 337, 124: 525, 339, 127: 336, 161: 524, 172: 522, 523, 199: 526},
		// 195
		{10: 159, 560},
		{10: 156},
		{10: 155, 155},
		{165, 411, 165, 10: 141, 141, 13: 165, 133: 412, 136: 528, 529, 156: 415, 180: 530},
		{10: 527},
		// 200
		{132, 10: 132, 132, 13: 132},
		{533, 2: 532, 13: 129, 185: 416, 418, 531},
		{10: 153, 153},
		{10: 152, 152},
		{537, 9: 179, 179, 179, 13: 536, 23: 179, 179, 179, 27: 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 52: 179, 179, 57: 179, 179},
		// 205
		{176, 9: 176, 176, 176, 13: 176, 23: 176, 176, 176, 27: 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 52: 176, 176, 57: 176, 176},
		{165, 411, 165, 10: 133, 13: 165, 23: 133, 133, 133, 27: 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 44: 133, 133, 133, 133, 133, 133, 133, 133: 412, 136: 528, 534, 156: 520, 174: 521},
		{10: 535},
		{175, 9: 175, 175, 175, 13: 175, 23: 175, 175, 175, 27: 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 52: 175, 175, 57: 175, 175},
		{161, 161, 161, 161, 161, 161, 161, 161, 161, 12: 161, 14: 161, 161, 161, 161, 161, 161, 161, 161, 161, 366, 364, 365, 44: 548, 55: 161, 113: 422, 154: 549, 163: 547},
		// 210
		{2: 540, 10: 149, 23: 170, 170, 170, 27: 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 44: 170, 170, 170, 170, 170, 170, 170, 170: 541, 194: 539, 213: 538},
		{23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 113: 338, 117: 357, 360, 356, 337, 124: 525, 339, 127: 336, 161: 524, 172: 522, 545},
		{10: 544},
		{10: 151, 151, 147: 151},
		{10: 148, 542},
		// 215
		{2: 543},
		{10: 150, 150, 147: 150},
		{168, 9: 168, 168, 168, 13: 168, 23: 168, 168, 168, 27: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 52: 168, 168, 57: 168, 168},
		{10: 546},
		{169, 9: 169, 169, 169, 13: 169, 23: 169, 169, 169, 27: 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 52: 169, 169, 57: 169, 169},
		// 220
		{391, 556, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 420, 256, 169: 557},
		{161, 161, 161, 161, 161, 161, 161, 161, 161, 12: 161, 14: 161, 161, 161, 161, 161, 161, 161, 161, 161, 366, 364, 365, 113: 422, 154: 426, 163: 553},
		{160, 160, 160, 160, 160, 160, 160, 160, 160, 12: 160, 14: 160, 160, 160, 160, 160, 160, 160, 160, 160, 366, 364, 365, 44: 550, 55: 160, 113: 511},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 551},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 13: 429, 52: 454, 55: 552, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		// 225
		{172, 9: 172, 172, 172, 13: 172, 23: 172, 172, 172, 27: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 52: 172, 172, 57: 172, 172},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 554},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 13: 429, 52: 454, 55: 555, 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{173, 9: 173, 173, 173, 13: 173, 23: 173, 173, 173, 27: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 52: 173, 173, 57: 173, 173},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 518, 559},
		// 230
		{55: 558},
		{174, 9: 174, 174, 174, 13: 174, 23: 174, 174, 174, 27: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 52: 174, 174, 57: 174, 174},
		{171, 9: 171, 171, 171, 13: 171, 23: 171, 171, 171, 27: 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 52: 171, 171, 57: 171, 171},
		{23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 113: 338, 117: 357, 360, 356, 337, 124: 525, 339, 127: 336, 147: 561, 161: 562},
		{10: 158},
		// 235
		{10: 154, 154},
		{139, 10: 139, 139, 13: 139},
		{10: 131, 23: 131, 131, 131, 27: 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 44: 131, 131, 131, 131, 131, 131, 131, 205: 565},
		{10: 157, 23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 113: 338, 117: 357, 360, 356, 337, 124: 525, 339, 127: 336, 161: 524, 172: 522, 523, 199: 566},
		{10: 567},
		// 240
		{130, 10: 130, 130, 13: 130},
		{167, 411, 167, 10: 167, 167, 13: 167, 133: 569},
		{166, 2: 166, 10: 166, 166, 13: 166},
		{200, 200, 200, 9: 200, 200, 13: 200, 43: 200},
		{198, 198, 198, 9: 198, 198, 13: 198, 43: 198},
		// 245
		{201, 201, 201, 9: 201, 201, 13: 201, 43: 201},
		{259, 259, 3: 259, 259, 259, 259, 259, 259, 259, 259, 259, 13: 259, 26: 259, 43: 259, 51: 259, 259, 55: 259, 259, 59: 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 145, 145, 145, 27: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 53: 579, 477, 114: 576, 123: 578, 155: 407, 162: 577},
		{430, 292, 3: 292, 292, 292, 292, 434, 433, 292, 292, 292, 13: 429, 26: 292, 43: 292, 51: 292, 292, 55: 292, 431, 59: 292, 292, 292, 292, 292, 292, 292, 292, 292, 432, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292},
		{10: 750, 480},
		// 250
		{10: 744},
		{10: 743},
		{104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 12: 104, 14: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 44: 104, 104, 104, 104, 104, 104, 104, 53: 104, 57: 104, 104, 85: 104, 104, 104, 104, 104, 104, 104, 104, 104, 95: 104, 104, 206: 580},
		{391, 396, 581, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 366, 364, 365, 100, 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 53: 579, 477, 57: 608, 376, 85: 604, 601, 606, 591, 605, 592, 602, 603, 599, 95: 607, 600, 113: 338, 582, 117: 357, 360, 356, 337, 598, 609, 585, 583, 339, 127: 336, 590, 586, 588, 589, 584, 134: 587, 597, 138: 335, 144: 596, 181: 594, 216: 595, 593},
		{315, 315, 3: 315, 315, 315, 315, 315, 315, 315, 11: 315, 13: 315, 43: 741, 52: 315, 56: 315, 59: 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315},
		// 255
		{9: 251, 251, 480},
		{165, 411, 165, 9: 238, 133: 710, 136: 709, 708, 171: 706, 196: 707, 705},
		{114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 12: 114, 14: 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 44: 114, 114, 114, 114, 114, 114, 114, 53: 114, 57: 114, 114, 85: 114, 114, 114, 114, 114, 114, 114, 114, 114, 95: 114, 114, 112: 114},
		{113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 12: 113, 14: 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 44: 113, 113, 113, 113, 113, 113, 113, 53: 113, 57: 113, 113, 85: 113, 113, 113, 113, 113, 113, 113, 113, 113, 95: 113, 113, 112: 113},
		{112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 12: 112, 14: 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 44: 112, 112, 112, 112, 112, 112, 112, 53: 112, 57: 112, 112, 85: 112, 112, 112, 112, 112, 112, 112, 112, 112, 95: 112, 112, 112: 112},
		// 260
		{111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 12: 111, 14: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 44: 111, 111, 111, 111, 111, 111, 111, 53: 111, 57: 111, 111, 85: 111, 111, 111, 111, 111, 111, 111, 111, 111, 95: 111, 111, 112: 111},
		{110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 12: 110, 14: 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 44: 110, 110, 110, 110, 110, 110, 110, 53: 110, 57: 110, 110, 85: 110, 110, 110, 110, 110, 110, 110, 110, 110, 95: 110, 110, 112: 110},
		{109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 12: 109, 14: 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 44: 109, 109, 109, 109, 109, 109, 109, 53: 109, 57: 109, 109, 85: 109, 109, 109, 109, 109, 109, 109, 109, 109, 95: 109, 109, 112: 109},
		{108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 12: 108, 14: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 44: 108, 108, 108, 108, 108, 108, 108, 53: 108, 57: 108, 108, 85: 108, 108, 108, 108, 108, 108, 108, 108, 108, 95: 108, 108, 112: 108},
		{250, 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 14: 250, 250, 250, 250, 250, 250, 250, 250, 250, 145: 378, 702},
		// 265
		{43: 700},
		{26: 699},
		{102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 12: 102, 14: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 44: 102, 102, 102, 102, 102, 102, 102, 53: 102, 57: 102, 102, 85: 102, 102, 102, 102, 102, 102, 102, 102, 102, 95: 102, 102},
		{391, 396, 581, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 366, 364, 365, 99, 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 53: 579, 477, 57: 608, 376, 85: 604, 601, 606, 591, 605, 592, 602, 603, 599, 95: 607, 600, 113: 338, 582, 117: 357, 360, 356, 337, 598, 609, 585, 583, 339, 127: 336, 590, 586, 588, 589, 584, 134: 587, 597, 138: 335, 144: 596, 181: 698},
		{98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 12: 98, 14: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 44: 98, 98, 98, 98, 98, 98, 98, 53: 98, 57: 98, 98, 85: 98, 98, 98, 98, 98, 98, 98, 98, 98, 95: 98, 98},
		// 270
		{97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 12: 97, 14: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 44: 97, 97, 97, 97, 97, 97, 97, 53: 97, 57: 97, 97, 85: 97, 97, 97, 97, 97, 97, 97, 97, 97, 95: 97, 97},
		{9: 697},
		{691},
		{687},
		{683},
		// 275
		{391, 396, 581, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 579, 477, 57: 608, 85: 604, 601, 606, 591, 605, 592, 602, 603, 599, 95: 607, 600, 114: 582, 121: 598, 609, 585, 128: 590, 586, 588, 589, 584, 134: 587, 677},
		{663},
		{391, 396, 659, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 660},
		{9: 658},
		{9: 657},
		// 280
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 477, 114: 582, 121: 655},
		{60, 23: 611, 85: 60, 203: 610},
		{51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 12: 51, 14: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 44: 51, 51, 51, 51, 51, 51, 51, 53: 51, 57: 51, 51, 85: 51, 51, 51, 51, 51, 51, 51, 51, 51, 95: 51, 51, 112: 51},
		{612, 85: 613},
		{59, 85: 59},
		// 285
		{12: 615, 164: 640},
		{614},
		{12: 615, 164: 616},
		{10: 63, 12: 63, 43: 63},
		{12: 617, 43: 618},
		// 290
		{10: 62, 12: 62, 43: 62},
		{43: 619},
		{12: 55, 623, 149: 621, 620, 157: 622},
		{12: 636},
		{10: 57, 57, 43: 57},
		// 295
		{11: 626, 43: 627},
		{2: 624},
		{55: 625},
		{12: 54},
		{12: 55, 623, 149: 635, 620},
		// 300
		{12: 628, 182: 629},
		{10: 53, 53, 43: 53},
		{11: 630, 43: 631},
		{12: 634},
		{2: 540, 170: 632},
		// 305
		{10: 633, 542},
		{47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 12: 47, 14: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 44: 47, 47, 47, 47, 47, 47, 47, 53: 47, 57: 47, 47, 85: 47, 47, 47, 47, 47, 47, 47, 47, 47, 95: 47, 47, 112: 47},
		{10: 52, 52, 43: 52},
		{10: 56, 56, 43: 56},
		{637},
		// 310
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 638},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 10: 639, 13: 429, 52: 454, 56: 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{10: 58, 58, 43: 58},
		{10: 641, 12: 617, 43: 642},
		{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 12: 61, 14: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 44: 61, 61, 61, 61, 61, 61, 61, 53: 61, 57: 61, 61, 85: 61, 61, 61, 61, 61, 61, 61, 61, 61, 95: 61, 61, 112: 61},
		// 315
		{10: 644, 12: 55, 623, 43: 645, 149: 621, 620, 157: 643},
		{10: 648, 626, 43: 649},
		{46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 12: 46, 14: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 44: 46, 46, 46, 46, 46, 46, 46, 53: 46, 57: 46, 46, 85: 46, 46, 46, 46, 46, 46, 46, 46, 46, 95: 46, 46, 112: 46},
		{12: 55, 623, 149: 621, 620, 157: 646},
		{10: 647, 626},
		// 320
		{45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 12: 45, 14: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 44: 45, 45, 45, 45, 45, 45, 45, 53: 45, 57: 45, 45, 85: 45, 45, 45, 45, 45, 45, 45, 45, 45, 95: 45, 45, 112: 45},
		{50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 12: 50, 14: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 44: 50, 50, 50, 50, 50, 50, 50, 53: 50, 57: 50, 50, 85: 50, 50, 50, 50, 50, 50, 50, 50, 50, 95: 50, 50, 112: 50},
		{12: 55, 623, 149: 621, 620, 157: 650},
		{10: 651, 626, 43: 652},
		{49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 12: 49, 14: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 44: 49, 49, 49, 49, 49, 49, 49, 53: 49, 57: 49, 49, 85: 49, 49, 49, 49, 49, 49, 49, 49, 49, 95: 49, 49, 112: 49},
		// 325
		{12: 628, 182: 653},
		{10: 654, 630},
		{48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 12: 48, 14: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 44: 48, 48, 48, 48, 48, 48, 48, 53: 48, 57: 48, 48, 85: 48, 48, 48, 48, 48, 48, 48, 48, 48, 95: 48, 48, 112: 48},
		{9: 656},
		{85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 12: 85, 14: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 44: 85, 85, 85, 85, 85, 85, 85, 53: 85, 57: 85, 85, 85: 85, 85, 85, 85, 85, 85, 85, 85, 85, 95: 85, 85, 112: 85},
		// 330
		{86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 12: 86, 14: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 44: 86, 86, 86, 86, 86, 86, 86, 53: 86, 57: 86, 86, 85: 86, 86, 86, 86, 86, 86, 86, 86, 86, 95: 86, 86, 112: 86},
		{87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 12: 87, 14: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 44: 87, 87, 87, 87, 87, 87, 87, 53: 87, 57: 87, 87, 85: 87, 87, 87, 87, 87, 87, 87, 87, 87, 95: 87, 87, 112: 87},
		{315, 315, 3: 315, 315, 315, 315, 315, 315, 662, 13: 315, 52: 315, 56: 315, 59: 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 661, 13: 429, 52: 454, 56: 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 12: 84, 14: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 44: 84, 84, 84, 84, 84, 84, 84, 53: 84, 57: 84, 84, 85: 84, 84, 84, 84, 84, 84, 84, 84, 84, 95: 84, 84, 112: 84},
		// 335
		{88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 12: 88, 14: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 44: 88, 88, 88, 88, 88, 88, 88, 53: 88, 57: 88, 88, 85: 88, 88, 88, 88, 88, 88, 88, 88, 88, 95: 88, 88, 112: 88},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 54: 477, 58: 376, 113: 338, 582, 117: 357, 360, 356, 337, 664, 124: 583, 339, 127: 336, 138: 335, 144: 665},
		{9: 671},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 477, 114: 582, 121: 666},
		{9: 667},
		// 340
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 10: 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 477, 114: 582, 121: 668},
		{10: 669},
		{391, 396, 581, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 579, 477, 57: 608, 85: 604, 601, 606, 591, 605, 592, 602, 603, 599, 95: 607, 600, 114: 582, 121: 598, 609, 585, 128: 590, 586, 588, 589, 584, 134: 587, 670},
		{89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 12: 89, 14: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 44: 89, 89, 89, 89, 89, 89, 89, 53: 89, 57: 89, 89, 85: 89, 89, 89, 89, 89, 89, 89, 89, 89, 95: 89, 89, 112: 89},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 477, 114: 582, 121: 672},
		// 345
		{9: 673},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 10: 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 477, 114: 582, 121: 674},
		{10: 675},
		{391, 396, 581, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 579, 477, 57: 608, 85: 604, 601, 606, 591, 605, 592, 602, 603, 599, 95: 607, 600, 114: 582, 121: 598, 609, 585, 128: 590, 586, 588, 589, 584, 134: 587, 676},
		{90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 12: 90, 14: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 44: 90, 90, 90, 90, 90, 90, 90, 53: 90, 57: 90, 90, 85: 90, 90, 90, 90, 90, 90, 90, 90, 90, 95: 90, 90, 112: 90},
		// 350
		{86: 678},
		{679},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 477, 114: 680},
		{10: 681, 480},
		{9: 682},
		// 355
		{91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 12: 91, 14: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 44: 91, 91, 91, 91, 91, 91, 91, 53: 91, 57: 91, 91, 85: 91, 91, 91, 91, 91, 91, 91, 91, 91, 95: 91, 91, 112: 91},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 477, 114: 684},
		{10: 685, 480},
		{391, 396, 581, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 579, 477, 57: 608, 85: 604, 601, 606, 591, 605, 592, 602, 603, 599, 95: 607, 600, 114: 582, 121: 598, 609, 585, 128: 590, 586, 588, 589, 584, 134: 587, 686},
		{92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 12: 92, 14: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 44: 92, 92, 92, 92, 92, 92, 92, 53: 92, 57: 92, 92, 85: 92, 92, 92, 92, 92, 92, 92, 92, 92, 95: 92, 92, 112: 92},
		// 360
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 477, 114: 688},
		{10: 689, 480},
		{391, 396, 581, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 579, 477, 57: 608, 85: 604, 601, 606, 591, 605, 592, 602, 603, 599, 95: 607, 600, 114: 582, 121: 598, 609, 585, 128: 590, 586, 588, 589, 584, 134: 587, 690},
		{93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 12: 93, 14: 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 44: 93, 93, 93, 93, 93, 93, 93, 53: 93, 57: 93, 93, 85: 93, 93, 93, 93, 93, 93, 93, 93, 93, 95: 93, 93, 112: 93},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 54: 477, 114: 692},
		// 365
		{10: 693, 480},
		{391, 396, 581, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 579, 477, 57: 608, 85: 604, 601, 606, 591, 605, 592, 602, 603, 599, 95: 607, 600, 114: 582, 121: 598, 609, 585, 128: 590, 586, 588, 589, 584, 134: 587, 694},
		{95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 12: 95, 14: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 44: 95, 95, 95, 95, 95, 95, 95, 53: 95, 57: 95, 95, 85: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95: 95, 95, 112: 695},
		{391, 396, 581, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 579, 477, 57: 608, 85: 604, 601, 606, 591, 605, 592, 602, 603, 599, 95: 607, 600, 114: 582, 121: 598, 609, 585, 128: 590, 586, 588, 589, 584, 134: 587, 696},
		{94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 12: 94, 14: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 44: 94, 94, 94, 94, 94, 94, 94, 53: 94, 57: 94, 94, 85: 94, 94, 94, 94, 94, 94, 94, 94, 94, 95: 94, 94, 112: 94},
		// 370
		{96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 12: 96, 14: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 44: 96, 96, 96, 96, 96, 96, 96, 53: 96, 57: 96, 96, 85: 96, 96, 96, 96, 96, 96, 96, 96, 96, 95: 96, 96, 112: 96},
		{101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 12: 101, 14: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 44: 101, 101, 101, 101, 101, 101, 101, 53: 101, 57: 101, 101, 85: 101, 101, 101, 101, 101, 101, 101, 101, 101, 95: 101, 101},
		{103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 12: 103, 14: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 44: 103, 103, 103, 103, 103, 103, 103, 103, 53: 103, 57: 103, 103, 85: 103, 103, 103, 103, 103, 103, 103, 103, 103, 95: 103, 103, 112: 103},
		{391, 396, 581, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 579, 477, 57: 608, 85: 604, 601, 606, 591, 605, 592, 602, 603, 599, 95: 607, 600, 114: 582, 121: 598, 609, 585, 128: 590, 586, 588, 589, 584, 134: 587, 701},
		{105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 12: 105, 14: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 44: 105, 105, 105, 105, 105, 105, 105, 53: 105, 57: 105, 105, 85: 105, 105, 105, 105, 105, 105, 105, 105, 105, 95: 105, 105, 112: 105},
		// 375
		{43: 703},
		{391, 396, 581, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 579, 477, 57: 608, 85: 604, 601, 606, 591, 605, 592, 602, 603, 599, 95: 607, 600, 114: 582, 121: 598, 609, 585, 128: 590, 586, 588, 589, 584, 134: 587, 704},
		{106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 12: 106, 14: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 44: 106, 106, 106, 106, 106, 106, 106, 53: 106, 57: 106, 106, 85: 106, 106, 106, 106, 106, 106, 106, 106, 106, 95: 106, 106, 112: 106},
		{9: 740},
		{9: 240, 11: 240},
		// 380
		{9: 237, 11: 738},
		{9: 236, 11: 236, 52: 235, 179: 712},
		{711, 2: 532, 187: 531},
		{164, 2: 164},
		{165, 411, 165, 133: 710, 136: 709, 534},
		// 385
		{52: 713},
		{391, 396, 714, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 717, 715, 160: 716},
		{315, 315, 3: 315, 315, 315, 315, 315, 315, 315, 11: 315, 13: 315, 26: 315, 43: 736, 52: 315, 56: 315, 59: 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 127, 11: 127, 13: 429, 26: 127, 52: 454, 56: 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{9: 234, 11: 234},
		// 390
		{120, 120, 120, 120, 120, 120, 120, 120, 120, 11: 122, 120, 723, 120, 120, 120, 120, 120, 120, 120, 120, 120, 26: 122, 53: 120, 56: 724, 159: 722, 166: 721, 719, 720, 198: 718},
		{11: 731, 26: 193, 165: 732},
		{391, 396, 714, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 717, 715, 160: 730},
		{13: 723, 52: 728, 56: 724, 159: 729},
		{119, 119, 119, 119, 119, 119, 119, 119, 119, 12: 119, 14: 119, 119, 119, 119, 119, 119, 119, 119, 119, 53: 119},
		// 395
		{13: 118, 52: 118, 56: 118},
		{250, 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 14: 250, 250, 250, 250, 250, 250, 250, 250, 250, 145: 378, 726},
		{2: 725},
		{13: 115, 52: 115, 56: 115},
		{55: 727},
		// 400
		{13: 116, 52: 116, 56: 116},
		{121, 121, 121, 121, 121, 121, 121, 121, 121, 12: 121, 14: 121, 121, 121, 121, 121, 121, 121, 121, 121, 53: 121},
		{13: 117, 52: 117, 56: 117},
		{11: 124, 26: 124},
		{120, 120, 120, 120, 120, 120, 120, 120, 120, 12: 120, 723, 120, 120, 120, 120, 120, 120, 120, 120, 120, 26: 192, 53: 120, 56: 724, 159: 722, 166: 721, 734, 720},
		// 405
		{26: 733},
		{9: 126, 11: 126, 26: 126},
		{391, 396, 714, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 717, 715, 160: 735},
		{11: 123, 26: 123},
		{391, 396, 714, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 717, 715, 160: 737},
		// 410
		{9: 125, 11: 125, 26: 125},
		{165, 411, 165, 133: 710, 136: 709, 708, 171: 739},
		{9: 239, 11: 239},
		{248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 12: 248, 14: 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 44: 248, 248, 248, 248, 248, 248, 248, 248, 53: 248, 57: 248, 248, 85: 248, 248, 248, 248, 248, 248, 248, 248, 248, 95: 248, 248},
		{391, 396, 581, 395, 397, 398, 403, 394, 393, 252, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 579, 477, 57: 608, 85: 604, 601, 606, 591, 605, 592, 602, 603, 599, 95: 607, 600, 114: 582, 121: 598, 609, 585, 128: 590, 586, 588, 589, 584, 134: 587, 742},
		// 415
		{107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 12: 107, 14: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 44: 107, 107, 107, 107, 107, 107, 107, 53: 107, 57: 107, 107, 85: 107, 107, 107, 107, 107, 107, 107, 107, 107, 95: 107, 107, 112: 107},
		{258, 258, 3: 258, 258, 258, 258, 258, 258, 258, 258, 258, 13: 258, 26: 258, 43: 258, 51: 258, 258, 55: 258, 258, 59: 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258},
		{391, 291, 384, 291, 291, 291, 291, 394, 393, 291, 291, 291, 390, 291, 400, 399, 402, 385, 386, 387, 388, 389, 401, 26: 291, 43: 291, 51: 291, 291, 746, 745, 291, 291, 59: 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291},
		{430, 290, 3: 290, 290, 290, 290, 434, 433, 290, 290, 290, 13: 429, 26: 290, 43: 290, 51: 290, 290, 55: 290, 431, 59: 290, 290, 290, 290, 290, 290, 290, 290, 290, 432, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290},
		{120, 120, 120, 120, 120, 120, 120, 120, 120, 11: 122, 120, 723, 120, 120, 120, 120, 120, 120, 120, 120, 120, 26: 122, 53: 120, 56: 724, 159: 722, 166: 721, 719, 720, 198: 747},
		// 420
		{11: 731, 26: 193, 165: 748},
		{26: 749},
		{301, 301, 3: 301, 301, 301, 301, 301, 301, 301, 301, 301, 13: 301, 26: 301, 43: 301, 51: 301, 301, 55: 301, 301, 59: 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301},
		{308, 308, 3: 308, 308, 308, 308, 308, 308, 308, 308, 308, 13: 308, 26: 308, 43: 308, 51: 308, 308, 55: 308, 308, 59: 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308},
		{430, 293, 3: 293, 293, 293, 293, 434, 433, 293, 293, 293, 13: 429, 26: 293, 43: 293, 51: 293, 293, 55: 293, 431, 59: 293, 293, 293, 293, 293, 293, 293, 293, 293, 432, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293},
		// 425
		{430, 294, 3: 294, 294, 294, 294, 434, 433, 294, 294, 294, 13: 429, 26: 294, 43: 294, 51: 294, 294, 55: 294, 431, 59: 294, 294, 294, 294, 294, 294, 294, 294, 294, 432, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294},
		{430, 295, 3: 295, 295, 295, 295, 434, 433, 295, 295, 295, 13: 429, 26: 295, 43: 295, 51: 295, 295, 55: 295, 431, 59: 295, 295, 295, 295, 295, 295, 295, 295, 295, 432, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295},
		{430, 296, 3: 296, 296, 296, 296, 434, 433, 296, 296, 296, 13: 429, 26: 296, 43: 296, 51: 296, 296, 55: 296, 431, 59: 296, 296, 296, 296, 296, 296, 296, 296, 296, 432, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296},
		{430, 298, 3: 298, 298, 298, 298, 434, 433, 298, 298, 298, 13: 429, 26: 298, 43: 298, 51: 298, 298, 55: 298, 431, 59: 298, 298, 298, 298, 298, 298, 298, 298, 298, 432, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298},
		{430, 299, 3: 299, 299, 299, 299, 434, 433, 299, 299, 299, 13: 429, 26: 299, 43: 299, 51: 299, 299, 55: 299, 431, 59: 299, 299, 299, 299, 299, 299, 299, 299, 299, 432, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299},
		// 430
		{430, 300, 3: 300, 300, 300, 300, 434, 433, 300, 300, 300, 13: 429, 26: 300, 43: 300, 51: 300, 300, 55: 300, 431, 59: 300, 300, 300, 300, 300, 300, 300, 300, 300, 432, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300},
		{10: 759},
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 53: 746, 745},
		{761},
		{12: 615, 164: 762},
		// 435
		{10: 641, 12: 617},
		{23: 65, 65, 65, 27: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 44: 65, 65, 65, 65, 65, 65, 65, 53: 66, 57: 66, 65, 178: 765, 183: 764},
		{53: 74, 57: 74, 208: 769},
		{23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 58: 376, 113: 338, 117: 357, 360, 356, 337, 124: 583, 339, 127: 336, 138: 335, 144: 766, 219: 767},
		{23: 68, 68, 68, 27: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 44: 68, 68, 68, 68, 68, 68, 68, 53: 68, 57: 68, 68},
		// 440
		{23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 44: 342, 343, 341, 367, 368, 344, 340, 53: 64, 57: 64, 376, 113: 338, 117: 357, 360, 356, 337, 124: 583, 339, 127: 336, 138: 335, 144: 768},
		{23: 67, 67, 67, 27: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 44: 67, 67, 67, 67, 67, 67, 67, 53: 67, 57: 67, 67},
		{53: 72, 57: 70, 176: 771, 772, 192: 770},
		{73, 73, 73, 9: 73, 23: 73, 73, 73, 27: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 44: 73, 73, 73, 73, 73, 73, 73, 73, 57: 73, 73},
		{53: 579, 123: 775},
		// 445
		{57: 608, 122: 609, 128: 773},
		{9: 774},
		{69, 69, 69, 9: 69, 23: 69, 69, 69, 27: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 44: 69, 69, 69, 69, 69, 69, 69, 69, 57: 69, 69},
		{71, 71, 71, 9: 71, 23: 71, 71, 71, 27: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 44: 71, 71, 71, 71, 71, 71, 71, 71, 57: 71, 71},
		{79, 79, 79, 9: 79, 23: 79, 79, 79, 27: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 44: 79, 79, 79, 79, 79, 79, 79, 79, 57: 79, 79},
		// 450
		{53: 191, 212: 779},
		{189, 189, 189, 9: 189, 189, 189, 13: 189, 23: 189, 189, 189, 27: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 53: 146},
		{53: 780},
		{2: 781, 189: 784, 783, 225: 782},
		{11: 320, 26: 320, 52: 320},
		// 455
		{11: 787, 26: 193, 165: 788},
		{11: 188, 26: 188},
		{11: 186, 26: 186, 52: 785},
		{250, 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 14: 250, 250, 250, 250, 250, 250, 250, 250, 250, 145: 378, 786},
		{11: 185, 26: 185},
		// 460
		{2: 781, 26: 192, 189: 784, 790},
		{26: 789},
		{190, 190, 190, 9: 190, 190, 190, 13: 190, 23: 190, 190, 190, 27: 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190},
		{11: 187, 26: 187},
		{53: 793},
		// 465
		{210, 210, 210, 9: 210, 210, 210, 13: 210, 23: 210, 210, 210, 27: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 53: 146},
		{23: 212, 212, 212, 795, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 58: 212, 211: 794},
		{23: 366, 364, 365, 27: 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 58: 376, 113: 409, 117: 357, 360, 356, 408, 138: 799, 153: 798, 201: 797, 230: 796},
		{209, 209, 209, 9: 209, 209, 209, 13: 209, 23: 209, 209, 209, 27: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209},
		{23: 366, 364, 365, 810, 354, 346, 355, 351, 363, 350, 348, 349, 347, 352, 361, 358, 359, 362, 353, 345, 58: 376, 113: 409, 117: 357, 360, 356, 408, 138: 799, 153: 798, 201: 811},
		// 470
		{23: 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 58: 206},
		{165, 411, 165, 9: 801, 43: 178, 133: 710, 136: 709, 803, 184: 804, 202: 802, 231: 800},
		{23: 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 58: 202},
		{9: 807, 11: 808},
		{23: 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 58: 203},
		// 475
		{9: 197, 11: 197},
		{9: 195, 11: 195, 43: 177},
		{43: 805},
		{250, 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 14: 250, 250, 250, 250, 250, 250, 250, 250, 250, 145: 378, 806},
		{9: 194, 11: 194},
		// 480
		{23: 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 58: 204},
		{165, 411, 165, 43: 178, 133: 710, 136: 709, 803, 184: 804, 202: 809},
		{9: 196, 11: 196},
		{211, 211, 211, 9: 211, 211, 211, 13: 211, 23: 211, 211, 211, 27: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211},
		{23: 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 58: 205},
		// 485
		{391, 396, 384, 395, 397, 398, 403, 394, 393, 12: 390, 14: 400, 399, 402, 385, 386, 387, 388, 389, 401, 145, 145, 145, 27: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 54: 813, 155: 407, 162: 814},
		{430, 435, 3: 448, 438, 439, 451, 434, 433, 10: 816, 13: 429, 52: 454, 56: 431, 59: 437, 436, 442, 443, 453, 449, 450, 458, 462, 432, 456, 446, 445, 444, 440, 460, 457, 455, 447, 464, 452, 441, 461, 459, 463},
		{10: 815},
		{213, 213, 213, 9: 213, 213, 213, 13: 213, 23: 213, 213, 213, 27: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213},
		{214, 214, 214, 9: 214, 214, 214, 13: 214, 23: 214, 214, 214, 27: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214},
		// 490
		{243, 243, 243, 9: 243, 243, 243, 13: 243},
		{241, 241, 241, 9: 241, 241, 241, 13: 241},
		{244, 244, 244, 9: 244, 244, 244, 13: 244},
		{245, 245, 245, 9: 245, 245, 245, 13: 245},
		{246, 246, 246, 9: 246, 246, 246, 13: 246},
		// 495
		{9: 236, 11: 236, 23: 65, 65, 65, 27: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 44: 65, 65, 65, 65, 65, 65, 65, 52: 235, 66, 57: 66, 65, 178: 765, 712, 183: 823},
		{53: 77, 57: 77, 207: 824},
		{53: 72, 57: 70, 176: 771, 772, 192: 825},
		{76, 76, 76, 9: 76, 23: 76, 76, 76, 27: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 44: 76, 76, 76, 76, 76, 76, 76, 76, 57: 76, 76},
		{82, 82, 82, 9: 82, 23: 82, 82, 82, 27: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 44: 82, 82, 82, 82, 82, 82, 82, 82, 57: 82, 82},
		// 500
		{250, 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 14: 250, 250, 250, 250, 250, 250, 250, 250, 250, 145: 378, 828},
		{51: 323},
		{84: 851, 94: 853, 100: 841, 842, 843, 838, 839, 840, 844, 848, 845, 835, 846, 847, 115: 852, 850, 126: 849, 139: 833, 832, 837, 834, 836, 148: 831, 228: 830},
		{51: 325},
		{51: 43, 84: 851, 94: 853, 100: 841, 842, 843, 838, 839, 840, 844, 848, 845, 835, 846, 847, 115: 852, 850, 126: 849, 139: 833, 891, 837, 834, 836},
		// 505
		{51: 42, 84: 42, 94: 42, 97: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{51: 38, 84: 38, 94: 38, 97: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{51: 37, 84: 37, 94: 37, 97: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{94: 853, 115: 913, 850},
		{51: 35, 84: 35, 94: 35, 97: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		// 510
		{97: 28, 28, 901, 188: 899, 220: 900, 898},
		{94: 853, 115: 895, 850},
		{2: 892},
		{2: 887},
		{2: 868, 84: 870, 226: 869},
		// 515
		{84: 851, 94: 853, 115: 852, 850, 126: 867},
		{51: 16, 84: 16, 94: 16, 97: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{94: 853, 115: 865, 850},
		{94: 853, 115: 863, 850},
		{84: 851, 94: 853, 115: 852, 850, 126: 862},
		// 520
		{2: 858},
		{94: 853, 115: 856, 850},
		{51: 7, 84: 7, 94: 7, 97: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{84: 5, 94: 855},
		{51: 4, 84: 4, 94: 4, 97: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		// 525
		{84: 854},
		{84: 2, 94: 2},
		{51: 3, 84: 3, 94: 3, 97: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{84: 1, 94: 1},
		{84: 857},
		// 530
		{51: 8, 84: 8, 94: 8, 97: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{84: 859, 94: 853, 115: 860, 850},
		{51: 12, 84: 12, 94: 12, 97: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{84: 861},
		{51: 9, 84: 9, 94: 9, 97: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		// 535
		{51: 13, 84: 13, 94: 13, 97: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{84: 864},
		{51: 14, 84: 14, 94: 14, 97: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{84: 866},
		{51: 15, 84: 15, 94: 15, 97: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		// 540
		{51: 17, 84: 17, 94: 17, 97: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{84: 851, 94: 853, 115: 852, 850, 126: 876, 152: 886},
		{2: 540, 10: 149, 147: 872, 170: 871, 194: 873},
		{51: 10, 84: 10, 94: 10, 97: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{10: 148, 879, 147: 880},
		// 545
		{10: 877},
		{10: 874},
		{84: 851, 94: 853, 115: 852, 850, 126: 876, 152: 875},
		{51: 18, 84: 18, 94: 18, 97: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{51: 6, 84: 6, 94: 6, 97: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		// 550
		{84: 851, 94: 853, 115: 852, 850, 126: 876, 152: 878},
		{51: 20, 84: 20, 94: 20, 97: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{2: 543, 147: 883},
		{10: 881},
		{84: 851, 94: 853, 115: 852, 850, 126: 876, 152: 882},
		// 555
		{51: 11, 84: 11, 94: 11, 97: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{10: 884},
		{84: 851, 94: 853, 115: 852, 850, 126: 876, 152: 885},
		{51: 19, 84: 19, 94: 19, 97: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{51: 21, 84: 21, 94: 21, 97: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		// 560
		{84: 888},
		{84: 851, 94: 853, 97: 40, 40, 40, 841, 842, 843, 838, 839, 840, 844, 848, 845, 835, 846, 847, 115: 852, 850, 126: 849, 139: 833, 832, 837, 834, 836, 148: 889, 151: 890},
		{84: 851, 94: 853, 97: 39, 39, 39, 841, 842, 843, 838, 839, 840, 844, 848, 845, 835, 846, 847, 115: 852, 850, 126: 849, 139: 833, 891, 837, 834, 836},
		{97: 31, 31, 31},
		{51: 41, 84: 41, 94: 41, 97: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		// 565
		{84: 893},
		{84: 851, 94: 853, 97: 40, 40, 40, 841, 842, 843, 838, 839, 840, 844, 848, 845, 835, 846, 847, 115: 852, 850, 126: 849, 139: 833, 832, 837, 834, 836, 148: 889, 151: 894},
		{97: 32, 32, 32},
		{84: 896},
		{84: 851, 94: 853, 97: 40, 40, 40, 841, 842, 843, 838, 839, 840, 844, 848, 845, 835, 846, 847, 115: 852, 850, 126: 849, 139: 833, 832, 837, 834, 836, 148: 889, 151: 897},
		// 570
		{97: 33, 33, 33},
		{97: 24, 907, 222: 908, 906},
		{97: 30, 30, 30},
		{97: 27, 27, 901, 188: 905},
		{94: 853, 115: 902, 850},
		// 575
		{84: 903},
		{84: 851, 94: 853, 97: 40, 40, 40, 841, 842, 843, 838, 839, 840, 844, 848, 845, 835, 846, 847, 115: 852, 850, 126: 849, 139: 833, 832, 837, 834, 836, 148: 889, 151: 904},
		{97: 26, 26, 26},
		{97: 29, 29, 29},
		{97: 912, 224: 911},
		// 580
		{84: 909},
		{97: 23},
		{84: 851, 94: 853, 97: 40, 100: 841, 842, 843, 838, 839, 840, 844, 848, 845, 835, 846, 847, 115: 852, 850, 126: 849, 139: 833, 832, 837, 834, 836, 148: 889, 151: 910},
		{97: 25},
		{51: 34, 84: 34, 94: 34, 97: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		// 585
		{51: 22, 84: 22, 94: 22, 97: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{84: 914},
		{51: 36, 84: 36, 94: 36, 97: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
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
			lx := yylex.(*lexer)
			lhs := &Expression{
				Case:   58,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableComputedGotos {
				lx.report.Err(lhs.Pos(), "computed gotos not enabled")
			}
		}
	case 71:
		{
			yyVAL.node = (*ExpressionOpt)(nil)
		}
	case 72:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionOpt{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 73:
		{
			yyVAL.node = &ExpressionList{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 74:
		{
			yyVAL.node = &ExpressionList{
				Case:           1,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList),
				Token:          yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 75:
		{
			yyVAL.node = (*ExpressionListOpt)(nil)
		}
	case 76:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionListOpt{
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 77:
		{
			lx := yylex.(*lexer)
			lx.constExprToks = []xc.Token{lx.last}
		}
	case 78:
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
	case 79:
		{
			lx := yylex.(*lexer)
			lhs := &Declaration{
				DeclarationSpecifiers: yyS[yypt-2].node.(*DeclarationSpecifiers),
				InitDeclaratorListOpt: yyS[yypt-1].node.(*InitDeclaratorListOpt),
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			ts0 := lhs.DeclarationSpecifiers.typeSpecifiers()
			if ts0 == 0 && lx.tweaks.enableImplicitIntType {
				lhs.DeclarationSpecifiers.typeSpecifier = tsEncode(tsInt)
			}
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
	case 80:
		{
			yyVAL.node = &Declaration{
				Case: 1,
				StaticAssertDeclaration: yyS[yypt-0].node.(*StaticAssertDeclaration),
			}
		}
	case 81:
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
	case 82:
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
	case 83:
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
	case 84:
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
	case 85:
		{
			yyVAL.node = (*DeclarationSpecifiersOpt)(nil)
		}
	case 86:
		{
			lhs := &DeclarationSpecifiersOpt{
				DeclarationSpecifiers: yyS[yypt-0].node.(*DeclarationSpecifiers),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.DeclarationSpecifiers.attr
			lhs.typeSpecifier = lhs.DeclarationSpecifiers.typeSpecifier
		}
	case 87:
		{
			yyVAL.node = &InitDeclaratorList{
				InitDeclarator: yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 88:
		{
			yyVAL.node = &InitDeclaratorList{
				Case:               1,
				InitDeclaratorList: yyS[yypt-2].node.(*InitDeclaratorList),
				Token:              yyS[yypt-1].Token,
				InitDeclarator:     yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 89:
		{
			yyVAL.node = (*InitDeclaratorListOpt)(nil)
		}
	case 90:
		{
			yyVAL.node = &InitDeclaratorListOpt{
				InitDeclaratorList: yyS[yypt-0].node.(*InitDeclaratorList).reverse(),
			}
		}
	case 91:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
		}
	case 92:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-0].node.(*Declarator)
			d.setFull(lx)
		}
	case 93:
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
	case 94:
		{
			lhs := &StorageClassSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saTypedef
		}
	case 95:
		{
			lhs := &StorageClassSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saExtern
		}
	case 96:
		{
			lhs := &StorageClassSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saStatic
		}
	case 97:
		{
			lhs := &StorageClassSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saAuto
		}
	case 98:
		{
			lhs := &StorageClassSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRegister
		}
	case 99:
		{
			lhs := &TypeSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsVoid)
		}
	case 100:
		{
			lhs := &TypeSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsChar)
		}
	case 101:
		{
			lhs := &TypeSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsShort)
		}
	case 102:
		{
			lhs := &TypeSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsInt)
		}
	case 103:
		{
			lhs := &TypeSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsLong)
		}
	case 104:
		{
			lhs := &TypeSpecifier{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsFloat)
		}
	case 105:
		{
			lhs := &TypeSpecifier{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsDouble)
		}
	case 106:
		{
			lhs := &TypeSpecifier{
				Case:  7,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsSigned)
		}
	case 107:
		{
			lhs := &TypeSpecifier{
				Case:  8,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsUnsigned)
		}
	case 108:
		{
			lhs := &TypeSpecifier{
				Case:  9,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsBool)
		}
	case 109:
		{
			lhs := &TypeSpecifier{
				Case:  10,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsComplex)
		}
	case 110:
		{
			lhs := &TypeSpecifier{
				Case: 11,
				StructOrUnionSpecifier: yyS[yypt-0].node.(*StructOrUnionSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(lhs.StructOrUnionSpecifier.typeSpecifiers())
		}
	case 111:
		{
			lhs := &TypeSpecifier{
				Case:          12,
				EnumSpecifier: yyS[yypt-0].node.(*EnumSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsEnumSpecifier)
		}
	case 112:
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
	case 113:
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
	case 114:
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
	case 115:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-1].node.(*IdentifierOpt); o != nil {
				lx.scope.declareStructTag(o.Token, lx.report)
			}
			lx.pushScope(ScopeMembers)
			lx.scope.isUnion = yyS[yypt-2].node.(*StructOrUnion).Case == 1 // "union"
			lx.scope.prevStructDeclarator = nil
		}
	case 116:
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
				finishBitField(lhs, lx)
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
	case 117:
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
	case 118:
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
	case 119:
		{
			yyVAL.node = &StructOrUnion{
				Token: yyS[yypt-0].Token,
			}
		}
	case 120:
		{
			yyVAL.node = &StructOrUnion{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 121:
		{
			yyVAL.node = &StructDeclarationList{
				StructDeclaration: yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 122:
		{
			yyVAL.node = &StructDeclarationList{
				Case: 1,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList),
				StructDeclaration:     yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 123:
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
	case 124:
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
	case 125:
		{
			yyVAL.node = &StructDeclaration{
				Case: 2,
				StaticAssertDeclaration: yyS[yypt-0].node.(*StaticAssertDeclaration),
			}
		}
	case 126:
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
	case 127:
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
	case 128:
		{
			yyVAL.node = (*SpecifierQualifierListOpt)(nil)
		}
	case 129:
		{
			lhs := &SpecifierQualifierListOpt{
				SpecifierQualifierList: yyS[yypt-0].node.(*SpecifierQualifierList),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.SpecifierQualifierList.attr
			lhs.typeSpecifier = lhs.SpecifierQualifierList.typeSpecifier
		}
	case 130:
		{
			yyVAL.node = &StructDeclaratorList{
				StructDeclarator: yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 131:
		{
			yyVAL.node = &StructDeclaratorList{
				Case:                 1,
				StructDeclaratorList: yyS[yypt-2].node.(*StructDeclaratorList),
				Token:                yyS[yypt-1].Token,
				StructDeclarator:     yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 132:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.post(lx)
		}
	case 133:
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
	case 134:
		{
			yyVAL.node = (*CommaOpt)(nil)
		}
	case 135:
		{
			yyVAL.node = &CommaOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 136:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareEnumTag(o.Token, lx.report)
			}
			lx.iota = 0
		}
	case 137:
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
			if !lx.tweaks.enableUnsignedEnums {
				break
			}

			lhs.unsigned = true
		loop:
			for l := lhs.EnumeratorList; l != nil; l = l.EnumeratorList {
				switch e := l.Enumerator; x := e.Value.(type) {
				case int32:
					if x < 0 {
						lhs.unsigned = false
						break loop
					}
				case int64:
					if x < 0 {
						lhs.unsigned = false
						break loop
					}
				default:
					panic(fmt.Errorf("%s: TODO Enumerator.Value type %T", position(e.Pos()), x))
				}
			}
		}
	case 138:
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
	case 139:
		{
			yyVAL.node = &EnumeratorList{
				Enumerator: yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 140:
		{
			yyVAL.node = &EnumeratorList{
				Case:           1,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList),
				Token:          yyS[yypt-1].Token,
				Enumerator:     yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 141:
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
	case 142:
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
				v = m.MustConvert(int32(0), m.IntType)
			default:
				var ok bool
				if v, ok = m.enumValueToInt(e.Value); !ok {
					lx.report.Err(e.Pos(), "overflow in enumeration value: %v", e.Value)
				}
			}

			lhs.Value = v
			lx.scope.defineEnumConst(lx, lhs.EnumerationConstant.Token, v)
		}
	case 143:
		{
			lhs := &TypeQualifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saConst
		}
	case 144:
		{
			lhs := &TypeQualifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRestrict
		}
	case 145:
		{
			lhs := &TypeQualifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saVolatile
		}
	case 146:
		{
			lhs := &FunctionSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saInline
		}
	case 147:
		{
			lhs := &FunctionSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saNoreturn
		}
	case 148:
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
	case 149:
		{
			yyVAL.node = (*DeclaratorOpt)(nil)
		}
	case 150:
		{
			yyVAL.node = &DeclaratorOpt{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
		}
	case 151:
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
	case 152:
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
	case 153:
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
	case 154:
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
	case 155:
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
	case 156:
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
	case 157:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 158:
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
	case 159:
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
	case 160:
		{
			yyVAL.node = &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].node.(*TypeQualifierListOpt),
			}
		}
	case 161:
		{
			yyVAL.node = &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].node.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].node.(*Pointer),
			}
		}
	case 162:
		{
			yyVAL.node = (*PointerOpt)(nil)
		}
	case 163:
		{
			yyVAL.node = &PointerOpt{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
		}
	case 164:
		{
			lhs := &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.TypeQualifier.attr
		}
	case 165:
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
	case 166:
		{
			yyVAL.node = (*TypeQualifierListOpt)(nil)
		}
	case 167:
		{
			yyVAL.node = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].node.(*TypeQualifierList).reverse(),
			}
		}
	case 168:
		{
			lhs := &ParameterTypeList{
				ParameterList: yyS[yypt-0].node.(*ParameterList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 169:
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
	case 170:
		{
			yyVAL.node = (*ParameterTypeListOpt)(nil)
		}
	case 171:
		{
			yyVAL.node = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].node.(*ParameterTypeList),
			}
		}
	case 172:
		{
			yyVAL.node = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 173:
		{
			yyVAL.node = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].node.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 174:
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
	case 175:
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
	case 176:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 177:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 178:
		{
			yyVAL.node = (*IdentifierListOpt)(nil)
		}
	case 179:
		{
			yyVAL.node = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 180:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 181:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 182:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeBlock)
		}
	case 183:
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
	case 184:
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
	case 185:
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
	case 186:
		{
			yyVAL.node = (*AbstractDeclaratorOpt)(nil)
		}
	case 187:
		{
			yyVAL.node = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].node.(*AbstractDeclarator),
			}
		}
	case 188:
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
	case 189:
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
	case 190:
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
	case 191:
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
	case 192:
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
	case 193:
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
	case 194:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 195:
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
	case 196:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 197:
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
	case 198:
		{
			yyVAL.node = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 199:
		{
			yyVAL.node = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
		}
	case 200:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 201:
		{
			yyVAL.node = &Initializer{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 202:
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
	case 203:
		{
			yyVAL.node = &InitializerList{
				DesignationOpt: yyS[yypt-1].node.(*DesignationOpt),
				Initializer:    yyS[yypt-0].node.(*Initializer),
			}
		}
	case 204:
		{
			yyVAL.node = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].node.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].node.(*DesignationOpt),
				Initializer:     yyS[yypt-0].node.(*Initializer),
			}
		}
	case 205:
		{
			yyVAL.node = (*InitializerList)(nil)
		}
	case 206:
		{
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 207:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 208:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 209:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 210:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 211:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 212:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 213:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 214:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 215:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 216:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 217:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 218:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 219:
		{
			yyVAL.node = &Statement{
				Case:               6,
				AssemblerStatement: yyS[yypt-0].node.(*AssemblerStatement),
			}
		}
	case 220:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 221:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 222:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 223:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 224:
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
	case 225:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 226:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 227:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 228:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 229:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 230:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 231:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 232:
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
	case 233:
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
	case 234:
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
	case 235:
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
	case 236:
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
	case 237:
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
	case 238:
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
	case 239:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 240:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 241:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 242:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 243:
		{
			lx := yylex.(*lexer)
			lhs := &JumpStatement{
				Case:       4,
				Token:      yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token2:     yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			_, t := lhs.Expression.eval(lx)
			if t == nil {
				break
			}

			for t != nil && t.Kind() == Ptr {
				t = t.Element()
			}

			if t == nil || t.Kind() != Void {
				lx.report.Err(lhs.Pos(), "invalid computed goto argument type, have '%s'", t)
			}

			if !lx.tweaks.enableComputedGotos {
				lx.report.Err(lhs.Pos(), "computed gotos not enabled")
			}
		}
	case 244:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 245:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 246:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 247:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 248:
		{
			yyVAL.node = &ExternalDeclaration{
				Case: 2,
				BasicAssemblerStatement: yyS[yypt-1].node.(*BasicAssemblerStatement),
				Token: yyS[yypt-0].Token,
			}
		}
	case 249:
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
	case 250:
		{
			lx := yylex.(*lexer)
			if ds := yyS[yypt-2].node.(*DeclarationSpecifiers); ds.typeSpecifier == 0 {
				ds.typeSpecifier = tsEncode(tsInt)
				yyS[yypt-1].node.(*Declarator).Type = lx.model.IntType
				if !lx.tweaks.enableOmitFuncRetType {
					lx.report.Err(yyS[yypt-1].node.Pos(), "missing function return type")
				}
			}
			var fd *FunctionDefinition
			fd.post(lx, yyS[yypt-1].node.(*Declarator), yyS[yypt-0].node.(*DeclarationListOpt))
		}
	case 251:
		{
			yyVAL.node = &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:          yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 252:
		{
			lx := yylex.(*lexer)
			lx.scope.specifier = &DeclarationSpecifiers{typeSpecifier: tsEncode(tsInt)}
		}
	case 253:
		{
			lx := yylex.(*lexer)
			if !lx.tweaks.enableOmitFuncRetType {
				lx.report.Err(yyS[yypt-1].node.Pos(), "missing function return type")
			}
			var fd *FunctionDefinition
			fd.post(lx, yyS[yypt-1].node.(*Declarator), yyS[yypt-0].node.(*DeclarationListOpt))
		}
	case 254:
		{
			yyVAL.node = &FunctionDefinition{
				Case:               1,
				Declarator:         yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt: yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:       yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 255:
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
	case 256:
		{
			lhs := &FunctionBody{
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
			yyVAL.node = lhs
			lhs.scope = lhs.CompoundStatement.scope
		}
	case 257:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 258:
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
	case 259:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 260:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 261:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 262:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 263:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScopePos(lhs.Pos())
		}
	case 264:
		{
			yyVAL.node = &AssemblerInstructions{
				Token: yyS[yypt-0].Token,
			}
		}
	case 265:
		{
			yyVAL.node = &AssemblerInstructions{
				Case: 1,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions),
				Token: yyS[yypt-0].Token,
			}
		}
	case 266:
		{
			yyVAL.node = &BasicAssemblerStatement{
				Token:                 yyS[yypt-4].Token,
				VolatileOpt:           yyS[yypt-3].node.(*VolatileOpt),
				Token2:                yyS[yypt-2].Token,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-0].Token,
			}
		}
	case 267:
		{
			yyVAL.node = (*VolatileOpt)(nil)
		}
	case 268:
		{
			yyVAL.node = &VolatileOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 269:
		{
			yyVAL.node = &AssemblerOperand{
				AssemblerSymbolicNameOpt: yyS[yypt-4].node.(*AssemblerSymbolicNameOpt),
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
		}
	case 270:
		{
			yyVAL.node = &AssemblerOperands{
				AssemblerOperand: yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 271:
		{
			yyVAL.node = &AssemblerOperands{
				Case:              1,
				AssemblerOperands: yyS[yypt-2].node.(*AssemblerOperands),
				Token:             yyS[yypt-1].Token,
				AssemblerOperand:  yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 272:
		{
			yyVAL.node = (*AssemblerSymbolicNameOpt)(nil)
		}
	case 273:
		{
			yyVAL.node = &AssemblerSymbolicNameOpt{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 274:
		{
			yyVAL.node = &Clobbers{
				Token: yyS[yypt-0].Token,
			}
		}
	case 275:
		{
			yyVAL.node = &Clobbers{
				Case:     1,
				Clobbers: yyS[yypt-2].node.(*Clobbers),
				Token:    yyS[yypt-1].Token,
				Token2:   yyS[yypt-0].Token,
			}
		}
	case 276:
		{
			yyVAL.node = &AssemblerStatement{
				BasicAssemblerStatement: yyS[yypt-0].node.(*BasicAssemblerStatement),
			}
		}
	case 277:
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
	case 278:
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
	case 279:
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
	case 280:
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
	case 281:
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
	case 282:
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
	case 283:
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
	case 284:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 285:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 286:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 287:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 288:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 289:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 290:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 291:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 292:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 293:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 294:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 295:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 296:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 297:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 298:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 299:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 300:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 301:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 302:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 303:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 304:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 305:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 306:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 307:
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
	case 308:
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
	case 309:
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
	case 310:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 311:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 312:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 313:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 314:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 315:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 316:
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
	case 317:
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
	case 318:
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
	case 319:
		{
			yyVAL.node = &ControlLine{
				Case:        13,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 322:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 323:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
