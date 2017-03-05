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
	yyTabOfs   = -323
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (323x)
		42:      1,   // '*' (284x)
		57377:   2,   // IDENTIFIER (237x)
		38:      3,   // '&' (216x)
		43:      4,   // '+' (216x)
		45:      5,   // '-' (216x)
		57363:   6,   // DEC (216x)
		57381:   7,   // INC (216x)
		59:      8,   // ';' (213x)
		41:      9,   // ')' (198x)
		44:      10,  // ',' (186x)
		57427:   11,  // STRINGLITERAL (165x)
		91:      12,  // '[' (161x)
		33:      13,  // '!' (146x)
		126:     14,  // '~' (146x)
		57347:   15,  // ALIGNOF (146x)
		57358:   16,  // CHARCONST (146x)
		57373:   17,  // FLOATCONST (146x)
		57384:   18,  // INTCONST (146x)
		57387:   19,  // LONGCHARCONST (146x)
		57388:   20,  // LONGSTRINGLITERAL (146x)
		57424:   21,  // SIZEOF (146x)
		57438:   22,  // VOLATILE (140x)
		57360:   23,  // CONST (138x)
		57418:   24,  // RESTRICT (138x)
		57353:   25,  // BOOL (128x)
		57357:   26,  // CHAR (128x)
		57359:   27,  // COMPLEX (128x)
		57367:   28,  // DOUBLE (128x)
		57369:   29,  // ENUM (128x)
		57372:   30,  // FLOAT (128x)
		57383:   31,  // INT (128x)
		57386:   32,  // LONG (128x)
		57422:   33,  // SHORT (128x)
		57423:   34,  // SIGNED (128x)
		57428:   35,  // STRUCT (128x)
		57432:   36,  // TYPEDEFNAME (128x)
		57433:   37,  // TYPEOF (128x)
		57435:   38,  // UNION (128x)
		57436:   39,  // UNSIGNED (128x)
		57437:   40,  // VOID (128x)
		125:     41,  // '}' (126x)
		58:      42,  // ':' (121x)
		57425:   43,  // STATIC (119x)
		57352:   44,  // AUTO (113x)
		57371:   45,  // EXTERN (113x)
		57382:   46,  // INLINE (113x)
		57395:   47,  // NORETURN (113x)
		57417:   48,  // REGISTER (113x)
		57431:   49,  // TYPEDEF (113x)
		57344:   50,  // $end (103x)
		61:      51,  // '=' (90x)
		123:     52,  // '{' (89x)
		57502:   53,  // Expression (85x)
		93:      54,  // ']' (83x)
		57351:   55,  // ASM (83x)
		46:      56,  // '.' (79x)
		57426:   57,  // STATIC_ASSERT (78x)
		37:      58,  // '%' (71x)
		47:      59,  // '/' (71x)
		60:      60,  // '<' (71x)
		62:      61,  // '>' (71x)
		63:      62,  // '?' (71x)
		94:      63,  // '^' (71x)
		124:     64,  // '|' (71x)
		57346:   65,  // ADDASSIGN (71x)
		57348:   66,  // ANDAND (71x)
		57349:   67,  // ANDASSIGN (71x)
		57350:   68,  // ARROW (71x)
		57365:   69,  // DIVASSIGN (71x)
		57370:   70,  // EQ (71x)
		57375:   71,  // GEQ (71x)
		57385:   72,  // LEQ (71x)
		57389:   73,  // LSH (71x)
		57390:   74,  // LSHASSIGN (71x)
		57391:   75,  // MODASSIGN (71x)
		57392:   76,  // MULASSIGN (71x)
		57393:   77,  // NEQ (71x)
		57396:   78,  // ORASSIGN (71x)
		57397:   79,  // OROR (71x)
		57420:   80,  // RSH (71x)
		57421:   81,  // RSHASSIGN (71x)
		57429:   82,  // SUBASSIGN (71x)
		57440:   83,  // XORASSIGN (71x)
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
		57481:   123, // DeclarationSpecifiers (15x)
		57510:   124, // FunctionSpecifier (15x)
		57528:   125, // PPTokenListOpt (15x)
		57544:   126, // StorageClassSpecifier (15x)
		57467:   127, // AssemblerStatement (13x)
		57475:   128, // CompoundStatement (13x)
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
		57530:   160, // ParameterDeclaration (4x)
		57553:   161, // TypeName (4x)
		57556:   162, // TypeQualifierListOpt (4x)
		57464:   163, // AssemblerInstructions (3x)
		57474:   164, // CommaOpt (3x)
		57485:   165, // Designation (3x)
		57486:   166, // DesignationOpt (3x)
		57488:   167, // DesignatorList (3x)
		57505:   168, // ExpressionOpt (3x)
		57514:   169, // IdentifierList (3x)
		57519:   170, // InitDeclarator (3x)
		57522:   171, // Initializer (3x)
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
		"'}'",
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
		"DeclarationSpecifiers",
		"FunctionSpecifier",
		"PPTokenListOpt",
		"StorageClassSpecifier",
		"AssemblerStatement",
		"CompoundStatement",
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
		"Initializer",
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
		69:  {168, 0},
		70:  {168, 1},
		71:  {114, 1},
		72:  {114, 3},
		73:  {121, 0},
		74:  {121, 1},
		75:  {145, 0},
		76:  {146, 2},
		77:  {144, 3},
		78:  {144, 1},
		79:  {123, 2},
		80:  {123, 2},
		81:  {123, 2},
		82:  {123, 2},
		83:  {158, 0},
		84:  {158, 1},
		85:  {196, 1},
		86:  {196, 3},
		87:  {197, 0},
		88:  {197, 1},
		89:  {170, 1},
		90:  {179, 0},
		91:  {170, 4},
		92:  {126, 1},
		93:  {126, 1},
		94:  {126, 1},
		95:  {126, 1},
		96:  {126, 1},
		97:  {120, 1},
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
		111: {120, 4},
		112: {120, 4},
		113: {211, 0},
		114: {119, 6},
		115: {119, 2},
		116: {119, 4},
		117: {118, 1},
		118: {118, 1},
		119: {230, 1},
		120: {230, 2},
		121: {201, 3},
		122: {201, 2},
		123: {201, 1},
		124: {153, 2},
		125: {153, 2},
		126: {200, 0},
		127: {200, 1},
		128: {231, 1},
		129: {231, 3},
		130: {202, 1},
		131: {202, 3},
		132: {164, 0},
		133: {164, 1},
		134: {212, 0},
		135: {117, 7},
		136: {117, 2},
		137: {225, 1},
		138: {225, 3},
		139: {190, 1},
		140: {190, 3},
		141: {113, 1},
		142: {113, 1},
		143: {113, 1},
		144: {124, 1},
		145: {124, 1},
		146: {137, 2},
		147: {184, 0},
		148: {184, 1},
		149: {187, 1},
		150: {187, 3},
		151: {187, 5},
		152: {187, 6},
		153: {187, 6},
		154: {187, 5},
		155: {213, 0},
		156: {187, 5},
		157: {187, 4},
		158: {133, 2},
		159: {133, 3},
		160: {136, 0},
		161: {136, 1},
		162: {154, 1},
		163: {154, 2},
		164: {162, 0},
		165: {162, 1},
		166: {173, 1},
		167: {173, 3},
		168: {199, 0},
		169: {199, 1},
		170: {172, 1},
		171: {172, 3},
		172: {160, 2},
		173: {160, 2},
		174: {169, 1},
		175: {169, 3},
		176: {194, 0},
		177: {194, 1},
		178: {195, 0},
		179: {195, 1},
		180: {155, 0},
		181: {161, 3},
		182: {156, 1},
		183: {156, 2},
		184: {180, 0},
		185: {180, 1},
		186: {185, 3},
		187: {185, 4},
		188: {185, 5},
		189: {185, 6},
		190: {185, 6},
		191: {185, 4},
		192: {174, 0},
		193: {185, 4},
		194: {205, 0},
		195: {185, 5},
		196: {186, 0},
		197: {186, 1},
		198: {171, 1},
		199: {171, 4},
		200: {198, 2},
		201: {198, 4},
		202: {198, 0},
		203: {165, 2},
		204: {166, 0},
		205: {166, 1},
		206: {167, 1},
		207: {167, 2},
		208: {159, 3},
		209: {159, 2},
		210: {135, 1},
		211: {135, 1},
		212: {135, 1},
		213: {135, 1},
		214: {135, 1},
		215: {135, 1},
		216: {135, 1},
		217: {132, 3},
		218: {132, 4},
		219: {132, 3},
		220: {206, 0},
		221: {128, 4},
		222: {216, 1},
		223: {216, 2},
		224: {217, 0},
		225: {217, 1},
		226: {181, 1},
		227: {181, 1},
		228: {129, 2},
		229: {134, 5},
		230: {134, 7},
		231: {134, 5},
		232: {130, 5},
		233: {130, 7},
		234: {130, 9},
		235: {130, 8},
		236: {131, 3},
		237: {131, 2},
		238: {131, 2},
		239: {131, 3},
		240: {233, 1},
		241: {233, 2},
		242: {191, 1},
		243: {191, 1},
		244: {191, 2},
		245: {191, 1},
		246: {207, 0},
		247: {193, 5},
		248: {175, 0},
		249: {208, 0},
		250: {193, 5},
		251: {176, 0},
		252: {192, 2},
		253: {177, 0},
		254: {192, 3},
		255: {219, 1},
		256: {219, 2},
		257: {183, 0},
		258: {178, 0},
		259: {183, 2},
		260: {163, 1},
		261: {163, 2},
		262: {122, 5},
		263: {203, 0},
		264: {203, 1},
		265: {149, 5},
		266: {157, 1},
		267: {157, 3},
		268: {150, 0},
		269: {150, 3},
		270: {182, 1},
		271: {182, 3},
		272: {127, 1},
		273: {127, 7},
		274: {127, 9},
		275: {127, 11},
		276: {127, 13},
		277: {127, 6},
		278: {127, 8},
		279: {138, 7},
		280: {228, 1},
		281: {148, 1},
		282: {148, 2},
		283: {151, 0},
		284: {151, 1},
		285: {140, 1},
		286: {140, 1},
		287: {140, 3},
		288: {140, 1},
		289: {142, 4},
		290: {141, 4},
		291: {141, 4},
		292: {141, 4},
		293: {220, 1},
		294: {220, 2},
		295: {221, 0},
		296: {221, 1},
		297: {188, 4},
		298: {222, 3},
		299: {223, 0},
		300: {223, 1},
		301: {224, 1},
		302: {139, 3},
		303: {139, 5},
		304: {139, 7},
		305: {139, 5},
		306: {139, 2},
		307: {139, 1},
		308: {139, 3},
		309: {139, 3},
		310: {139, 2},
		311: {139, 3},
		312: {139, 6},
		313: {139, 2},
		314: {139, 4},
		315: {139, 3},
		316: {143, 1},
		317: {152, 1},
		318: {115, 1},
		319: {125, 1},
		320: {125, 2},
		321: {116, 1},
		322: {116, 2},
	}

	yyXErrors = map[yyXError]string{
		{0, 50}:   "invalid empty input",
		{572, -1}: "expected #endif",
		{574, -1}: "expected #endif",
		{1, -1}:   "expected $end",
		{492, -1}: "expected $end",
		{494, -1}: "expected $end",
		{32, -1}:  "expected '('",
		{49, -1}:  "expected '('",
		{75, -1}:  "expected '('",
		{285, -1}: "expected '('",
		{320, -1}: "expected '('",
		{341, -1}: "expected '('",
		{380, -1}: "expected '('",
		{381, -1}: "expected '('",
		{382, -1}: "expected '('",
		{384, -1}: "expected '('",
		{410, -1}: "expected '('",
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
		{283, -1}: "expected ')'",
		{400, -1}: "expected ')'",
		{406, -1}: "expected ')'",
		{478, -1}: "expected ')'",
		{536, -1}: "expected ')'",
		{537, -1}: "expected ')'",
		{544, -1}: "expected ')'",
		{547, -1}: "expected ')'",
		{52, -1}:  "expected ','",
		{323, -1}: "expected ':'",
		{373, -1}: "expected ':'",
		{434, -1}: "expected ':'",
		{468, -1}: "expected ':'",
		{45, -1}:  "expected ';'",
		{55, -1}:  "expected ';'",
		{301, -1}: "expected ';'",
		{315, -1}: "expected ';'",
		{379, -1}: "expected ';'",
		{386, -1}: "expected ';'",
		{387, -1}: "expected ';'",
		{389, -1}: "expected ';'",
		{393, -1}: "expected ';'",
		{396, -1}: "expected ';'",
		{398, -1}: "expected ';'",
		{404, -1}: "expected ';'",
		{413, -1}: "expected ';'",
		{305, -1}: "expected '='",
		{89, -1}:  "expected '['",
		{516, -1}: "expected '\\n'",
		{520, -1}: "expected '\\n'",
		{524, -1}: "expected '\\n'",
		{527, -1}: "expected '\\n'",
		{529, -1}: "expected '\\n'",
		{551, -1}: "expected '\\n'",
		{556, -1}: "expected '\\n'",
		{559, -1}: "expected '\\n'",
		{566, -1}: "expected '\\n'",
		{571, -1}: "expected '\\n'",
		{577, -1}: "expected '\\n'",
		{95, -1}:  "expected ']'",
		{184, -1}: "expected ']'",
		{228, -1}: "expected ']'",
		{260, -1}: "expected ']'",
		{329, -1}: "expected ']'",
		{441, -1}: "expected '{'",
		{443, -1}: "expected '{'",
		{455, -1}: "expected '{'",
		{269, -1}: "expected '}'",
		{273, -1}: "expected '}'",
		{374, -1}: "expected '}'",
		{452, -1}: "expected '}'",
		{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		{204, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{88, -1}:  "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{287, -1}: "expected assembler instructions or string literal",
		{319, -1}: "expected assembler instructions or string literal",
		{321, -1}: "expected assembler instructions or string literal",
		{331, -1}: "expected assembler operand or one of ['[', string literal]",
		{346, -1}: "expected assembler operands or one of [')', ':', '[', string literal]",
		{324, -1}: "expected assembler operands or one of ['[', string literal]",
		{349, -1}: "expected assembler operands or one of ['[', string literal]",
		{353, -1}: "expected assembler operands or one of ['[', string literal]",
		{314, -1}: "expected assembler statement or asm",
		{376, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{332, -1}: "expected clobbers or string literal",
		{356, -1}: "expected clobbers or string literal",
		{313, -1}: "expected compound statement or '{'",
		{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{50, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{257, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{372, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{449, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{469, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{491, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{296, -1}: "expected declaration list or one of [_Bool, _Complex, _Noreturn, _Static_assert, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{299, -1}: "expected declaration or one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{395, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{47, -1}:  "expected declarator or one of ['(', '*', identifier]",
		{439, -1}: "expected declarator or one of ['(', '*', identifier]",
		{196, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		{254, -1}: "expected designator or one of ['.', '=', '[']",
		{199, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		{85, -1}:  "expected direct abstract declarator or one of ['(', '[']",
		{292, -1}: "expected direct declarator or one of ['(', identifier]",
		{564, -1}: "expected elif group or one of [#elif, #else, #endif]",
		{570, -1}: "expected endif line or #endif",
		{501, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		{562, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		{444, -1}: "expected enumerator list or identifier",
		{451, -1}: "expected enumerator or one of ['}', identifier]",
		{100, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{124, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{411, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{415, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{419, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{423, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{64, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{245, -1}: "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{249, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		{92, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{227, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{284, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		{342, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{96, -1}:  "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		{219, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		{476, -1}: "expected expression or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{6, -1}:   "expected external declaration or one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{295, -1}: "expected function body or one of ['{', asm]",
		{311, -1}: "expected function body or one of ['{', asm]",
		{487, -1}: "expected function body or one of ['{', asm]",
		{488, -1}: "expected function body or one of ['{', asm]",
		{486, -1}: "expected function body or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{294, -1}: "expected function body or optional declaration list or one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{553, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{495, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{102, -1}: "expected identifier",
		{103, -1}: "expected identifier",
		{213, -1}: "expected identifier",
		{258, -1}: "expected identifier",
		{328, -1}: "expected identifier",
		{385, -1}: "expected identifier",
		{503, -1}: "expected identifier",
		{504, -1}: "expected identifier",
		{511, -1}: "expected identifier",
		{336, -1}: "expected identifier list or identifier",
		{533, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		{308, -1}: "expected init declarator or one of ['(', '*', identifier]",
		{251, -1}: "expected initializer list or optional comma or one of ['!', '&', '(', '*', '+', ',', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{265, -1}: "expected initializer list or optional comma or one of ['!', '&', '(', '*', '+', ',', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{253, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{271, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{306, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{268, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		{250, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{274, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{275, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{276, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{277, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{278, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{279, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{280, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{281, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{282, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{65, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{148, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{152, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{174, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{179, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{343, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{477, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{363, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{264, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{91, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{99, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{186, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{222, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{225, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{496, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{497, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{498, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{500, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{507, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{513, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{515, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{518, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{521, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{523, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{525, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{526, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{528, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{530, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{531, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{534, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{539, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{540, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{542, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{546, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{549, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{550, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{555, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{575, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{576, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{578, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{554, -1}: "expected one of [#elif, #else, #endif]",
		{558, -1}: "expected one of [#elif, #else, #endif]",
		{561, -1}: "expected one of [#elif, #else, #endif]",
		{563, -1}: "expected one of [#elif, #else, #endif]",
		{568, -1}: "expected one of [#elif, #else, #endif]",
		{569, -1}: "expected one of [#elif, #else, #endif]",
		{431, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{8, -1}:   "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{56, -1}:  "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{310, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{42, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{43, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{44, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{46, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{312, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{359, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{361, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{440, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{489, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{490, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{38, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{39, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{93, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{182, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{291, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{317, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{338, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{348, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{351, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{352, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{355, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{358, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{365, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{366, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{367, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{368, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{369, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{370, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{371, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{390, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{391, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{392, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{394, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{402, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{408, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{414, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{418, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{422, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{426, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{428, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{429, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{433, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{436, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{438, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{375, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{377, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{378, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{430, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{255, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{262, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{442, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{456, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
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
		{453, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{459, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{474, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{479, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{480, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{17, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{40, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{41, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{481, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{482, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{483, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{484, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{485, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
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
		{286, -1}: "expected one of ['(', goto]",
		{318, -1}: "expected one of ['(', goto]",
		{293, -1}: "expected one of ['(', identifier]",
		{326, -1}: "expected one of [')', ',', ':']",
		{333, -1}: "expected one of [')', ',', ':']",
		{339, -1}: "expected one of [')', ',', ':']",
		{340, -1}: "expected one of [')', ',', ':']",
		{344, -1}: "expected one of [')', ',', ':']",
		{347, -1}: "expected one of [')', ',', ':']",
		{354, -1}: "expected one of [')', ',', ':']",
		{364, -1}: "expected one of [')', ',', ';']",
		{211, -1}: "expected one of [')', ',', ...]",
		{214, -1}: "expected one of [')', ',', ...]",
		{535, -1}: "expected one of [')', ',', ...]",
		{86, -1}:  "expected one of [')', ',']",
		{175, -1}: "expected one of [')', ',']",
		{193, -1}: "expected one of [')', ',']",
		{195, -1}: "expected one of [')', ',']",
		{200, -1}: "expected one of [')', ',']",
		{201, -1}: "expected one of [')', ',']",
		{212, -1}: "expected one of [')', ',']",
		{233, -1}: "expected one of [')', ',']",
		{247, -1}: "expected one of [')', ',']",
		{337, -1}: "expected one of [')', ',']",
		{350, -1}: "expected one of [')', ',']",
		{357, -1}: "expected one of [')', ',']",
		{412, -1}: "expected one of [')', ',']",
		{416, -1}: "expected one of [')', ',']",
		{420, -1}: "expected one of [')', ',']",
		{424, -1}: "expected one of [')', ',']",
		{288, -1}: "expected one of [')', ':', string literal]",
		{290, -1}: "expected one of [')', ':', string literal]",
		{345, -1}: "expected one of [')', ':', string literal]",
		{289, -1}: "expected one of [')', string literal]",
		{467, -1}: "expected one of [',', ':', ';']",
		{149, -1}: "expected one of [',', ':']",
		{327, -1}: "expected one of [',', ':']",
		{334, -1}: "expected one of [',', ':']",
		{304, -1}: "expected one of [',', ';', '=']",
		{270, -1}: "expected one of [',', ';', '}']",
		{302, -1}: "expected one of [',', ';']",
		{303, -1}: "expected one of [',', ';']",
		{307, -1}: "expected one of [',', ';']",
		{309, -1}: "expected one of [',', ';']",
		{464, -1}: "expected one of [',', ';']",
		{466, -1}: "expected one of [',', ';']",
		{470, -1}: "expected one of [',', ';']",
		{473, -1}: "expected one of [',', ';']",
		{445, -1}: "expected one of [',', '=', '}']",
		{448, -1}: "expected one of [',', '=', '}']",
		{180, -1}: "expected one of [',', ']']",
		{266, -1}: "expected one of [',', '}']",
		{272, -1}: "expected one of [',', '}']",
		{447, -1}: "expected one of [',', '}']",
		{450, -1}: "expected one of [',', '}']",
		{454, -1}: "expected one of [',', '}']",
		{256, -1}: "expected one of ['.', '=', '[']",
		{259, -1}: "expected one of ['.', '=', '[']",
		{261, -1}: "expected one of ['.', '=', '[']",
		{263, -1}: "expected one of ['.', '=', '[']",
		{322, -1}: "expected one of [':', string literal]",
		{505, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		{514, -1}: "expected one of ['\\n', ppother]",
		{517, -1}: "expected one of ['\\n', ppother]",
		{519, -1}: "expected one of ['\\n', ppother]",
		{298, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{300, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{34, -1}:  "expected one of ['{', identifier]",
		{35, -1}:  "expected one of ['{', identifier]",
		{461, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{463, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{465, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{471, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{475, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{543, -1}: "expected one of [..., identifier]",
		{81, -1}:  "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		{101, -1}: "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{360, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{362, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{252, -1}: "expected optional comma or one of [',', '}']",
		{267, -1}: "expected optional comma or one of [',', '}']",
		{446, -1}: "expected optional comma or one of [',', '}']",
		{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{12, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{399, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{405, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{388, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{397, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{403, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{218, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{207, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{90, -1}:  "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{94, -1}:  "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{552, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{557, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{560, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{567, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{573, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{208, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{33, -1}:  "expected optional identifier or one of ['{', identifier]",
		{36, -1}:  "expected optional identifier or one of ['{', identifier]",
		{297, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		{192, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{235, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{236, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{79, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{80, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{506, -1}: "expected optional token list or one of ['\\n', ppother]",
		{510, -1}: "expected optional token list or one of ['\\n', ppother]",
		{82, -1}:  "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		{316, -1}: "expected optional volatile or one of ['(', goto, volatile]",
		{48, -1}:  "expected optional volatile or one of ['(', volatile]",
		{231, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{209, -1}: "expected parameter type list or one of [_Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{239, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{493, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{532, -1}: "expected replacement list or one of ['\\n', ppother]",
		{538, -1}: "expected replacement list or one of ['\\n', ppother]",
		{541, -1}: "expected replacement list or one of ['\\n', ppother]",
		{545, -1}: "expected replacement list or one of ['\\n', ppother]",
		{548, -1}: "expected replacement list or one of ['\\n', ppother]",
		{78, -1}:  "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{383, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{401, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{407, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{417, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{421, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{425, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{427, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{432, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{435, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{437, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{53, -1}:  "expected string literal",
		{325, -1}: "expected string literal",
		{330, -1}: "expected string literal",
		{335, -1}: "expected string literal",
		{457, -1}: "expected struct declaration list or one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{458, -1}: "expected struct declaration list or one of [_Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{460, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{462, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		{472, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		{522, -1}: "expected token list or one of ['\\n', ppother]",
		{499, -1}: "expected token list or ppother",
		{502, -1}: "expected token list or ppother",
		{508, -1}: "expected token list or ppother",
		{509, -1}: "expected token list or ppother",
		{512, -1}: "expected token list or ppother",
		{565, -1}: "expected token list or ppother",
		{4, -1}:   "expected translation unit or one of ['(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{5, -1}:   "expected translation unit or one of ['(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{76, -1}:  "expected type name or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{97, -1}:  "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		{220, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{409, -1}: "expected while",
		{3, 50}:   "unexpected EOF",
		{2, 50}:   "unexpected EOF",
		{4, 50}:   "unexpected EOF",
	}

	yyParseTab = [579][]uint16{
		// 0
		{218: 326, 227: 325, 229: 324, 232: 327},
		{50: 323},
		{84: 322, 86: 322, 100: 322, 322, 322, 322, 322, 322, 322, 322, 322, 322, 322, 322, 204: 816},
		{320, 320, 320, 320, 320, 320, 320, 320, 11: 320, 13: 320, 320, 320, 320, 320, 320, 320, 320, 320, 209: 814},
		{318, 318, 318, 8: 318, 22: 318, 318, 318, 318, 318, 318, 318, 318, 318, 318, 318, 318, 318, 318, 318, 318, 318, 318, 318, 43: 318, 318, 318, 318, 318, 318, 318, 55: 318, 57: 318, 210: 328},
		// 5
		{75, 75, 75, 8: 369, 22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 55: 371, 57: 372, 113: 334, 117: 353, 356, 352, 333, 122: 368, 330, 335, 126: 332, 138: 331, 144: 367, 175: 370, 191: 365, 193: 366, 233: 329},
		{75, 75, 75, 8: 369, 22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 317, 55: 371, 57: 372, 113: 334, 117: 353, 356, 352, 333, 122: 368, 330, 335, 126: 332, 138: 331, 144: 367, 175: 370, 191: 813, 193: 366},
		{163, 405, 163, 8: 236, 133: 616, 136: 615, 809, 170: 625, 196: 626, 624},
		{245, 245, 245, 245, 245, 245, 245, 245, 245, 11: 245, 13: 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 43: 245, 245, 245, 245, 245, 245, 245, 245, 52: 245, 55: 245, 57: 245, 85: 245, 87: 245, 245, 245, 245, 245, 245, 245, 245, 245, 245},
		{240, 240, 240, 8: 240, 240, 240, 12: 240, 22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 113: 334, 117: 353, 356, 352, 333, 123: 805, 335, 126: 332, 158: 808},
		// 10
		{240, 240, 240, 8: 240, 240, 240, 12: 240, 22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 113: 334, 117: 353, 356, 352, 333, 123: 805, 335, 126: 332, 158: 807},
		{240, 240, 240, 8: 240, 240, 240, 12: 240, 22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 113: 334, 117: 353, 356, 352, 333, 123: 805, 335, 126: 332, 158: 806},
		{240, 240, 240, 8: 240, 240, 240, 12: 240, 22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 113: 334, 117: 353, 356, 352, 333, 123: 805, 335, 126: 332, 158: 804},
		{231, 231, 231, 8: 231, 231, 231, 12: 231, 22: 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 43: 231, 231, 231, 231, 231, 231, 231},
		{230, 230, 230, 8: 230, 230, 230, 12: 230, 22: 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 43: 230, 230, 230, 230, 230, 230, 230},
		// 15
		{229, 229, 229, 8: 229, 229, 229, 12: 229, 22: 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 43: 229, 229, 229, 229, 229, 229, 229},
		{228, 228, 228, 8: 228, 228, 228, 12: 228, 22: 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 43: 228, 228, 228, 228, 228, 228, 228},
		{227, 227, 227, 8: 227, 227, 227, 12: 227, 22: 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 43: 227, 227, 227, 227, 227, 227, 227},
		{226, 226, 226, 8: 226, 226, 226, 12: 226, 22: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 42: 226, 226, 226, 226, 226, 226, 226, 226},
		{225, 225, 225, 8: 225, 225, 225, 12: 225, 22: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 42: 225, 225, 225, 225, 225, 225, 225, 225},
		// 20
		{224, 224, 224, 8: 224, 224, 224, 12: 224, 22: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 42: 224, 224, 224, 224, 224, 224, 224, 224},
		{223, 223, 223, 8: 223, 223, 223, 12: 223, 22: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 42: 223, 223, 223, 223, 223, 223, 223, 223},
		{222, 222, 222, 8: 222, 222, 222, 12: 222, 22: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 42: 222, 222, 222, 222, 222, 222, 222, 222},
		{221, 221, 221, 8: 221, 221, 221, 12: 221, 22: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 42: 221, 221, 221, 221, 221, 221, 221, 221},
		{220, 220, 220, 8: 220, 220, 220, 12: 220, 22: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 42: 220, 220, 220, 220, 220, 220, 220, 220},
		// 25
		{219, 219, 219, 8: 219, 219, 219, 12: 219, 22: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 42: 219, 219, 219, 219, 219, 219, 219, 219},
		{218, 218, 218, 8: 218, 218, 218, 12: 218, 22: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 42: 218, 218, 218, 218, 218, 218, 218, 218},
		{217, 217, 217, 8: 217, 217, 217, 12: 217, 22: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 42: 217, 217, 217, 217, 217, 217, 217, 217},
		{216, 216, 216, 8: 216, 216, 216, 12: 216, 22: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 42: 216, 216, 216, 216, 216, 216, 216, 216},
		{215, 215, 215, 8: 215, 215, 215, 12: 215, 22: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 42: 215, 215, 215, 215, 215, 215, 215, 215},
		// 30
		{214, 214, 214, 8: 214, 214, 214, 12: 214, 22: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 42: 214, 214, 214, 214, 214, 214, 214, 214},
		{213, 213, 213, 8: 213, 213, 213, 12: 213, 22: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 42: 213, 213, 213, 213, 213, 213, 213, 213},
		{799},
		{2: 779, 52: 145, 195: 778},
		{2: 206, 52: 206},
		// 35
		{2: 205, 52: 205},
		{2: 765, 52: 145, 195: 764},
		{182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 42: 182, 182, 182, 182, 182, 182, 182, 182, 54: 182},
		{181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 42: 181, 181, 181, 181, 181, 181, 181, 181, 54: 181},
		{180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 42: 180, 180, 180, 180, 180, 180, 180, 180, 54: 180},
		// 40
		{179, 179, 179, 8: 179, 179, 179, 12: 179, 22: 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 43: 179, 179, 179, 179, 179, 179, 179},
		{178, 178, 178, 8: 178, 178, 178, 12: 178, 22: 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 43: 178, 178, 178, 178, 178, 178, 178},
		{83, 83, 83, 8: 83, 22: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 43: 83, 83, 83, 83, 83, 83, 83, 83, 55: 83, 57: 83},
		{81, 81, 81, 8: 81, 22: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 43: 81, 81, 81, 81, 81, 81, 81, 81, 55: 81, 57: 81},
		{80, 80, 80, 8: 80, 22: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 43: 80, 80, 80, 80, 80, 80, 80, 80, 55: 80, 57: 80},
		// 45
		{8: 763},
		{78, 78, 78, 8: 78, 22: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 43: 78, 78, 78, 78, 78, 78, 78, 78, 55: 78, 57: 78},
		{163, 405, 163, 133: 616, 136: 615, 617},
		{60, 22: 609, 203: 608},
		{373},
		// 50
		{248, 248, 248, 248, 248, 248, 248, 248, 11: 248, 13: 248, 248, 248, 248, 248, 248, 248, 248, 248, 145: 374, 375},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 388},
		{10: 376},
		{11: 377},
		{9: 378},
		// 55
		{8: 379},
		{44, 44, 44, 44, 44, 44, 44, 44, 44, 11: 44, 13: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 43: 44, 44, 44, 44, 44, 44, 44, 44, 52: 44, 55: 44, 57: 44, 85: 44, 87: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44},
		{311, 311, 3: 311, 311, 311, 311, 311, 311, 311, 311, 12: 311, 41: 311, 311, 50: 311, 311, 54: 311, 56: 311, 58: 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311},
		{310, 310, 3: 310, 310, 310, 310, 310, 310, 310, 310, 12: 310, 41: 310, 310, 50: 310, 310, 54: 310, 56: 310, 58: 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310},
		{309, 309, 3: 309, 309, 309, 309, 309, 309, 309, 309, 12: 309, 41: 309, 309, 50: 309, 309, 54: 309, 56: 309, 58: 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309},
		// 60
		{308, 308, 3: 308, 308, 308, 308, 308, 308, 308, 308, 12: 308, 41: 308, 308, 50: 308, 308, 54: 308, 56: 308, 58: 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308},
		{307, 307, 3: 307, 307, 307, 307, 307, 307, 307, 307, 12: 307, 41: 307, 307, 50: 307, 307, 54: 307, 56: 307, 58: 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307},
		{306, 306, 3: 306, 306, 306, 306, 306, 306, 306, 306, 12: 306, 41: 306, 306, 50: 306, 306, 54: 306, 56: 306, 58: 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306},
		{305, 305, 3: 305, 305, 305, 305, 305, 305, 305, 305, 12: 305, 41: 305, 305, 50: 305, 305, 54: 305, 56: 305, 58: 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 53: 471, 114: 570, 155: 401, 161: 606},
		// 65
		{424, 429, 3: 442, 432, 433, 428, 427, 247, 10: 247, 12: 423, 41: 247, 247, 50: 247, 448, 54: 247, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 605},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 604},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 603},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 512},
		// 70
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 602},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 601},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 600},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 599},
		{568, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 569},
		// 75
		{399},
		{22: 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 155: 401, 161: 400},
		{9: 567},
		{22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 113: 403, 117: 353, 356, 352, 402, 153: 404},
		{197, 197, 197, 8: 197, 197, 12: 197, 22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 42: 197, 113: 403, 117: 353, 356, 352, 402, 153: 565, 200: 566},
		// 80
		{197, 197, 197, 8: 197, 197, 12: 197, 22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 42: 197, 113: 403, 117: 353, 356, 352, 402, 153: 565, 200: 564},
		{163, 405, 9: 139, 12: 163, 133: 406, 136: 408, 156: 409, 180: 407},
		{159, 159, 159, 9: 159, 159, 12: 159, 22: 362, 360, 361, 113: 416, 154: 420, 162: 562},
		{162, 2: 162, 9: 141, 141, 12: 162},
		{9: 142},
		// 85
		{411, 12: 127, 185: 410, 412},
		{9: 138, 138},
		{558, 9: 140, 140, 12: 126},
		{163, 405, 9: 131, 12: 163, 22: 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 43: 131, 131, 131, 131, 131, 131, 131, 133: 406, 136: 408, 156: 514, 174: 515},
		{12: 413},
		// 90
		{387, 415, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 362, 360, 361, 43: 419, 53: 414, 254, 113: 416, 154: 417, 168: 418},
		{424, 429, 3: 442, 432, 433, 428, 427, 12: 423, 51: 448, 54: 253, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 512, 513},
		{161, 161, 161, 161, 161, 161, 161, 161, 9: 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 43: 161, 54: 161},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 362, 360, 361, 43: 508, 53: 414, 254, 113: 505, 168: 507},
		// 95
		{54: 506},
		{159, 159, 159, 159, 159, 159, 159, 159, 11: 159, 13: 159, 159, 159, 159, 159, 159, 159, 159, 159, 362, 360, 361, 113: 416, 154: 420, 162: 421},
		{158, 158, 158, 158, 158, 158, 158, 158, 9: 158, 158, 158, 158, 158, 158, 158, 158, 158, 158, 158, 158, 158, 362, 360, 361, 113: 505},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 422},
		{424, 429, 3: 442, 432, 433, 428, 427, 12: 423, 51: 448, 54: 459, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		// 100
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 471, 114: 503},
		{387, 392, 380, 391, 393, 394, 390, 389, 9: 313, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 497, 214: 498, 499},
		{2: 496},
		{2: 495},
		{299, 299, 3: 299, 299, 299, 299, 299, 299, 299, 299, 12: 299, 41: 299, 299, 50: 299, 299, 54: 299, 56: 299, 58: 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299},
		// 105
		{298, 298, 3: 298, 298, 298, 298, 298, 298, 298, 298, 12: 298, 41: 298, 298, 50: 298, 298, 54: 298, 56: 298, 58: 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 494},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 493},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 492},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 491},
		// 110
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 490},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 489},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 488},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 487},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 486},
		// 115
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 485},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 484},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 483},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 482},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 481},
		// 120
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 480},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 479},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 478},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 477},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 471, 114: 472},
		// 125
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 470},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 469},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 468},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 467},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 466},
		// 130
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 465},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 464},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 463},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 462},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 461},
		// 135
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 460},
		{134, 9: 134, 134, 12: 134},
		{424, 429, 3: 442, 432, 433, 428, 427, 256, 256, 256, 12: 423, 41: 256, 256, 50: 256, 448, 54: 256, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{424, 429, 3: 442, 432, 433, 428, 427, 257, 257, 257, 12: 423, 41: 257, 257, 50: 257, 448, 54: 257, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{424, 429, 3: 442, 432, 433, 428, 427, 258, 258, 258, 12: 423, 41: 258, 258, 50: 258, 448, 54: 258, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		// 140
		{424, 429, 3: 442, 432, 433, 428, 427, 259, 259, 259, 12: 423, 41: 259, 259, 50: 259, 448, 54: 259, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{424, 429, 3: 442, 432, 433, 428, 427, 260, 260, 260, 12: 423, 41: 260, 260, 50: 260, 448, 54: 260, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{424, 429, 3: 442, 432, 433, 428, 427, 261, 261, 261, 12: 423, 41: 261, 261, 50: 261, 448, 54: 261, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{424, 429, 3: 442, 432, 433, 428, 427, 262, 262, 262, 12: 423, 41: 262, 262, 50: 262, 448, 54: 262, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{424, 429, 3: 442, 432, 433, 428, 427, 263, 263, 263, 12: 423, 41: 263, 263, 50: 263, 448, 54: 263, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		// 145
		{424, 429, 3: 442, 432, 433, 428, 427, 264, 264, 264, 12: 423, 41: 264, 264, 50: 264, 448, 54: 264, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{424, 429, 3: 442, 432, 433, 428, 427, 265, 265, 265, 12: 423, 41: 265, 265, 50: 265, 448, 54: 265, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{424, 429, 3: 442, 432, 433, 428, 427, 266, 266, 266, 12: 423, 41: 266, 266, 50: 266, 448, 54: 266, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{424, 429, 3: 442, 432, 433, 428, 427, 252, 252, 252, 12: 423, 42: 252, 51: 448, 54: 252, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{10: 474, 42: 473},
		// 150
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 476},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 475},
		{424, 429, 3: 442, 432, 433, 428, 427, 251, 251, 251, 12: 423, 42: 251, 51: 448, 54: 251, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{424, 429, 3: 442, 432, 433, 428, 427, 267, 267, 267, 12: 423, 41: 267, 267, 50: 267, 267, 54: 267, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 267, 445, 267, 426, 267, 440, 439, 438, 434, 267, 267, 267, 441, 267, 446, 435, 267, 267, 267},
		{424, 429, 3: 442, 432, 433, 428, 427, 268, 268, 268, 12: 423, 41: 268, 268, 50: 268, 268, 54: 268, 56: 425, 58: 431, 430, 436, 437, 268, 443, 444, 268, 445, 268, 426, 268, 440, 439, 438, 434, 268, 268, 268, 441, 268, 268, 435, 268, 268, 268},
		// 155
		{424, 429, 3: 442, 432, 433, 428, 427, 269, 269, 269, 12: 423, 41: 269, 269, 50: 269, 269, 54: 269, 56: 425, 58: 431, 430, 436, 437, 269, 443, 444, 269, 269, 269, 426, 269, 440, 439, 438, 434, 269, 269, 269, 441, 269, 269, 435, 269, 269, 269},
		{424, 429, 3: 442, 432, 433, 428, 427, 270, 270, 270, 12: 423, 41: 270, 270, 50: 270, 270, 54: 270, 56: 425, 58: 431, 430, 436, 437, 270, 443, 270, 270, 270, 270, 426, 270, 440, 439, 438, 434, 270, 270, 270, 441, 270, 270, 435, 270, 270, 270},
		{424, 429, 3: 442, 432, 433, 428, 427, 271, 271, 271, 12: 423, 41: 271, 271, 50: 271, 271, 54: 271, 56: 425, 58: 431, 430, 436, 437, 271, 271, 271, 271, 271, 271, 426, 271, 440, 439, 438, 434, 271, 271, 271, 441, 271, 271, 435, 271, 271, 271},
		{424, 429, 3: 272, 432, 433, 428, 427, 272, 272, 272, 12: 423, 41: 272, 272, 50: 272, 272, 54: 272, 56: 425, 58: 431, 430, 436, 437, 272, 272, 272, 272, 272, 272, 426, 272, 440, 439, 438, 434, 272, 272, 272, 441, 272, 272, 435, 272, 272, 272},
		{424, 429, 3: 273, 432, 433, 428, 427, 273, 273, 273, 12: 423, 41: 273, 273, 50: 273, 273, 54: 273, 56: 425, 58: 431, 430, 436, 437, 273, 273, 273, 273, 273, 273, 426, 273, 273, 439, 438, 434, 273, 273, 273, 273, 273, 273, 435, 273, 273, 273},
		// 160
		{424, 429, 3: 274, 432, 433, 428, 427, 274, 274, 274, 12: 423, 41: 274, 274, 50: 274, 274, 54: 274, 56: 425, 58: 431, 430, 436, 437, 274, 274, 274, 274, 274, 274, 426, 274, 274, 439, 438, 434, 274, 274, 274, 274, 274, 274, 435, 274, 274, 274},
		{424, 429, 3: 275, 432, 433, 428, 427, 275, 275, 275, 12: 423, 41: 275, 275, 50: 275, 275, 54: 275, 56: 425, 58: 431, 430, 275, 275, 275, 275, 275, 275, 275, 275, 426, 275, 275, 275, 275, 434, 275, 275, 275, 275, 275, 275, 435, 275, 275, 275},
		{424, 429, 3: 276, 432, 433, 428, 427, 276, 276, 276, 12: 423, 41: 276, 276, 50: 276, 276, 54: 276, 56: 425, 58: 431, 430, 276, 276, 276, 276, 276, 276, 276, 276, 426, 276, 276, 276, 276, 434, 276, 276, 276, 276, 276, 276, 435, 276, 276, 276},
		{424, 429, 3: 277, 432, 433, 428, 427, 277, 277, 277, 12: 423, 41: 277, 277, 50: 277, 277, 54: 277, 56: 425, 58: 431, 430, 277, 277, 277, 277, 277, 277, 277, 277, 426, 277, 277, 277, 277, 434, 277, 277, 277, 277, 277, 277, 435, 277, 277, 277},
		{424, 429, 3: 278, 432, 433, 428, 427, 278, 278, 278, 12: 423, 41: 278, 278, 50: 278, 278, 54: 278, 56: 425, 58: 431, 430, 278, 278, 278, 278, 278, 278, 278, 278, 426, 278, 278, 278, 278, 434, 278, 278, 278, 278, 278, 278, 435, 278, 278, 278},
		// 165
		{424, 429, 3: 279, 432, 433, 428, 427, 279, 279, 279, 12: 423, 41: 279, 279, 50: 279, 279, 54: 279, 56: 425, 58: 431, 430, 279, 279, 279, 279, 279, 279, 279, 279, 426, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279},
		{424, 429, 3: 280, 432, 433, 428, 427, 280, 280, 280, 12: 423, 41: 280, 280, 50: 280, 280, 54: 280, 56: 425, 58: 431, 430, 280, 280, 280, 280, 280, 280, 280, 280, 426, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280},
		{424, 429, 3: 281, 281, 281, 428, 427, 281, 281, 281, 12: 423, 41: 281, 281, 50: 281, 281, 54: 281, 56: 425, 58: 431, 430, 281, 281, 281, 281, 281, 281, 281, 281, 426, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281},
		{424, 429, 3: 282, 282, 282, 428, 427, 282, 282, 282, 12: 423, 41: 282, 282, 50: 282, 282, 54: 282, 56: 425, 58: 431, 430, 282, 282, 282, 282, 282, 282, 282, 282, 426, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282},
		{424, 283, 3: 283, 283, 283, 428, 427, 283, 283, 283, 12: 423, 41: 283, 283, 50: 283, 283, 54: 283, 56: 425, 58: 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 426, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283},
		// 170
		{424, 284, 3: 284, 284, 284, 428, 427, 284, 284, 284, 12: 423, 41: 284, 284, 50: 284, 284, 54: 284, 56: 425, 58: 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 426, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284},
		{424, 285, 3: 285, 285, 285, 428, 427, 285, 285, 285, 12: 423, 41: 285, 285, 50: 285, 285, 54: 285, 56: 425, 58: 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 426, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285},
		{300, 300, 3: 300, 300, 300, 300, 300, 300, 300, 300, 12: 300, 41: 300, 300, 50: 300, 300, 54: 300, 56: 300, 58: 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300},
		{301, 301, 3: 301, 301, 301, 301, 301, 301, 301, 301, 12: 301, 41: 301, 301, 50: 301, 301, 54: 301, 56: 301, 58: 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301},
		{424, 429, 3: 442, 432, 433, 428, 427, 9: 315, 315, 12: 423, 51: 448, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		// 175
		{9: 312, 501},
		{9: 500},
		{302, 302, 3: 302, 302, 302, 302, 302, 302, 302, 302, 12: 302, 41: 302, 302, 50: 302, 302, 54: 302, 56: 302, 58: 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 502},
		{424, 429, 3: 442, 432, 433, 428, 427, 9: 314, 314, 12: 423, 51: 448, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		// 180
		{10: 474, 54: 504},
		{303, 303, 3: 303, 303, 303, 303, 303, 303, 303, 303, 12: 303, 41: 303, 303, 50: 303, 303, 54: 303, 56: 303, 58: 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303},
		{160, 160, 160, 160, 160, 160, 160, 160, 9: 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 43: 160, 54: 160},
		{136, 9: 136, 136, 12: 136},
		{54: 511},
		// 185
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 509},
		{424, 429, 3: 442, 432, 433, 428, 427, 12: 423, 51: 448, 54: 510, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{133, 9: 133, 133, 12: 133},
		{135, 9: 135, 135, 12: 135},
		{424, 293, 3: 293, 293, 293, 428, 427, 293, 293, 293, 12: 423, 41: 293, 293, 50: 293, 293, 54: 293, 56: 425, 58: 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 426, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293},
		// 190
		{132, 9: 132, 132, 12: 132},
		{9: 557},
		{9: 155, 22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 113: 334, 117: 353, 356, 352, 333, 123: 519, 335, 126: 332, 160: 518, 172: 516, 517, 199: 520},
		{9: 157, 554},
		{9: 154},
		// 195
		{9: 153, 153},
		{163, 405, 163, 9: 139, 139, 12: 163, 133: 406, 136: 522, 523, 156: 409, 180: 524},
		{9: 521},
		{130, 9: 130, 130, 12: 130},
		{527, 2: 526, 12: 127, 185: 410, 412, 525},
		// 200
		{9: 151, 151},
		{9: 150, 150},
		{531, 8: 177, 177, 177, 12: 530, 22: 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 42: 177, 177, 177, 177, 177, 177, 177, 177, 51: 177, 177, 55: 177, 57: 177},
		{174, 8: 174, 174, 174, 12: 174, 22: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 42: 174, 174, 174, 174, 174, 174, 174, 174, 51: 174, 174, 55: 174, 57: 174},
		{163, 405, 163, 9: 131, 12: 163, 22: 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 43: 131, 131, 131, 131, 131, 131, 131, 133: 406, 136: 522, 528, 156: 514, 174: 515},
		// 205
		{9: 529},
		{173, 8: 173, 173, 173, 12: 173, 22: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 42: 173, 173, 173, 173, 173, 173, 173, 173, 51: 173, 173, 55: 173, 57: 173},
		{159, 159, 159, 159, 159, 159, 159, 159, 11: 159, 13: 159, 159, 159, 159, 159, 159, 159, 159, 159, 362, 360, 361, 43: 542, 54: 159, 113: 416, 154: 543, 162: 541},
		{2: 534, 9: 147, 22: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 43: 168, 168, 168, 168, 168, 168, 168, 169: 535, 194: 533, 213: 532},
		{22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 113: 334, 117: 353, 356, 352, 333, 123: 519, 335, 126: 332, 160: 518, 172: 516, 539},
		// 210
		{9: 538},
		{9: 149, 149, 147: 149},
		{9: 146, 536},
		{2: 537},
		{9: 148, 148, 147: 148},
		// 215
		{166, 8: 166, 166, 166, 12: 166, 22: 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 42: 166, 166, 166, 166, 166, 166, 166, 166, 51: 166, 166, 55: 166, 57: 166},
		{9: 540},
		{167, 8: 167, 167, 167, 12: 167, 22: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 42: 167, 167, 167, 167, 167, 167, 167, 167, 51: 167, 167, 55: 167, 57: 167},
		{387, 550, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 414, 254, 168: 551},
		{159, 159, 159, 159, 159, 159, 159, 159, 11: 159, 13: 159, 159, 159, 159, 159, 159, 159, 159, 159, 362, 360, 361, 113: 416, 154: 420, 162: 547},
		// 220
		{158, 158, 158, 158, 158, 158, 158, 158, 11: 158, 13: 158, 158, 158, 158, 158, 158, 158, 158, 158, 362, 360, 361, 43: 544, 54: 158, 113: 505},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 545},
		{424, 429, 3: 442, 432, 433, 428, 427, 12: 423, 51: 448, 54: 546, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{170, 8: 170, 170, 170, 12: 170, 22: 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 42: 170, 170, 170, 170, 170, 170, 170, 170, 51: 170, 170, 55: 170, 57: 170},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 548},
		// 225
		{424, 429, 3: 442, 432, 433, 428, 427, 12: 423, 51: 448, 54: 549, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{171, 8: 171, 171, 171, 12: 171, 22: 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 42: 171, 171, 171, 171, 171, 171, 171, 171, 51: 171, 171, 55: 171, 57: 171},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 512, 553},
		{54: 552},
		{172, 8: 172, 172, 172, 12: 172, 22: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 42: 172, 172, 172, 172, 172, 172, 172, 172, 51: 172, 172, 55: 172, 57: 172},
		// 230
		{169, 8: 169, 169, 169, 12: 169, 22: 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 42: 169, 169, 169, 169, 169, 169, 169, 169, 51: 169, 169, 55: 169, 57: 169},
		{22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 113: 334, 117: 353, 356, 352, 333, 123: 519, 335, 126: 332, 147: 555, 160: 556},
		{9: 156},
		{9: 152, 152},
		{137, 9: 137, 137, 12: 137},
		// 235
		{9: 129, 22: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 43: 129, 129, 129, 129, 129, 129, 129, 205: 559},
		{9: 155, 22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 113: 334, 117: 353, 356, 352, 333, 123: 519, 335, 126: 332, 160: 518, 172: 516, 517, 199: 560},
		{9: 561},
		{128, 9: 128, 128, 12: 128},
		{165, 405, 165, 9: 165, 165, 12: 165, 133: 563},
		// 240
		{164, 2: 164, 9: 164, 164, 12: 164},
		{198, 198, 198, 8: 198, 198, 12: 198, 42: 198},
		{196, 196, 196, 8: 196, 196, 12: 196, 42: 196},
		{199, 199, 199, 8: 199, 199, 12: 199, 42: 199},
		{255, 255, 3: 255, 255, 255, 255, 255, 255, 255, 255, 12: 255, 41: 255, 255, 50: 255, 255, 54: 255, 56: 255, 58: 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		// 245
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 53: 471, 114: 570, 155: 401, 161: 571},
		{424, 288, 3: 288, 288, 288, 428, 427, 288, 288, 288, 12: 423, 41: 288, 288, 50: 288, 288, 54: 288, 56: 425, 58: 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 426, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288},
		{9: 598, 474},
		{9: 572},
		{387, 287, 380, 287, 287, 287, 390, 389, 287, 287, 287, 386, 287, 396, 395, 398, 381, 382, 383, 384, 385, 397, 41: 287, 287, 50: 287, 287, 574, 573, 287, 56: 287, 58: 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287},
		// 250
		{424, 286, 3: 286, 286, 286, 428, 427, 286, 286, 286, 12: 423, 41: 286, 286, 50: 286, 286, 54: 286, 56: 425, 58: 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 426, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286},
		{119, 119, 119, 119, 119, 119, 119, 119, 10: 121, 119, 580, 119, 119, 119, 119, 119, 119, 119, 119, 119, 41: 121, 52: 119, 56: 581, 159: 579, 165: 578, 576, 577, 198: 575},
		{10: 591, 41: 191, 164: 596},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 588, 587, 171: 589},
		{12: 580, 51: 585, 56: 581, 159: 586},
		// 255
		{118, 118, 118, 118, 118, 118, 118, 118, 11: 118, 13: 118, 118, 118, 118, 118, 118, 118, 118, 118, 52: 118},
		{12: 117, 51: 117, 56: 117},
		{248, 248, 248, 248, 248, 248, 248, 248, 11: 248, 13: 248, 248, 248, 248, 248, 248, 248, 248, 248, 145: 374, 583},
		{2: 582},
		{12: 114, 51: 114, 56: 114},
		// 260
		{54: 584},
		{12: 115, 51: 115, 56: 115},
		{120, 120, 120, 120, 120, 120, 120, 120, 11: 120, 13: 120, 120, 120, 120, 120, 120, 120, 120, 120, 52: 120},
		{12: 116, 51: 116, 56: 116},
		{424, 429, 3: 442, 432, 433, 428, 427, 125, 10: 125, 12: 423, 41: 125, 51: 448, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		// 265
		{119, 119, 119, 119, 119, 119, 119, 119, 10: 121, 119, 580, 119, 119, 119, 119, 119, 119, 119, 119, 119, 41: 121, 52: 119, 56: 581, 159: 579, 165: 578, 576, 577, 198: 590},
		{10: 123, 41: 123},
		{10: 591, 41: 191, 164: 592},
		{119, 119, 119, 119, 119, 119, 119, 119, 11: 119, 580, 119, 119, 119, 119, 119, 119, 119, 119, 119, 41: 190, 52: 119, 56: 581, 159: 579, 165: 578, 594, 577},
		{41: 593},
		// 270
		{8: 124, 10: 124, 41: 124},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 588, 587, 171: 595},
		{10: 122, 41: 122},
		{41: 597},
		{297, 297, 3: 297, 297, 297, 297, 297, 297, 297, 297, 12: 297, 41: 297, 297, 50: 297, 297, 54: 297, 56: 297, 58: 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297},
		// 275
		{304, 304, 3: 304, 304, 304, 304, 304, 304, 304, 304, 12: 304, 41: 304, 304, 50: 304, 304, 54: 304, 56: 304, 58: 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304},
		{424, 289, 3: 289, 289, 289, 428, 427, 289, 289, 289, 12: 423, 41: 289, 289, 50: 289, 289, 54: 289, 56: 425, 58: 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 426, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289},
		{424, 290, 3: 290, 290, 290, 428, 427, 290, 290, 290, 12: 423, 41: 290, 290, 50: 290, 290, 54: 290, 56: 425, 58: 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 426, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290},
		{424, 291, 3: 291, 291, 291, 428, 427, 291, 291, 291, 12: 423, 41: 291, 291, 50: 291, 291, 54: 291, 56: 425, 58: 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 426, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291},
		{424, 292, 3: 292, 292, 292, 428, 427, 292, 292, 292, 12: 423, 41: 292, 292, 50: 292, 292, 54: 292, 56: 425, 58: 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 426, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292},
		// 280
		{424, 294, 3: 294, 294, 294, 428, 427, 294, 294, 294, 12: 423, 41: 294, 294, 50: 294, 294, 54: 294, 56: 425, 58: 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 426, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294},
		{424, 295, 3: 295, 295, 295, 428, 427, 295, 295, 295, 12: 423, 41: 295, 295, 50: 295, 295, 54: 295, 56: 425, 58: 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 426, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295},
		{424, 296, 3: 296, 296, 296, 428, 427, 296, 296, 296, 12: 423, 41: 296, 296, 50: 296, 296, 54: 296, 56: 425, 58: 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 426, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296},
		{9: 607},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 574, 573},
		// 285
		{610},
		{59, 85: 59},
		{11: 611, 163: 612},
		{9: 63, 11: 63, 42: 63},
		{9: 614, 11: 613},
		// 290
		{9: 62, 11: 62, 42: 62},
		{61, 61, 61, 61, 61, 61, 61, 61, 61, 11: 61, 13: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 43: 61, 61, 61, 61, 61, 61, 61, 52: 61, 55: 61, 57: 61, 85: 61, 87: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 112: 61},
		{762, 2: 526, 187: 525},
		{162, 2: 162},
		{22: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 43: 65, 65, 65, 65, 65, 65, 65, 52: 66, 55: 66, 57: 65, 178: 619, 183: 618},
		// 295
		{52: 74, 55: 74, 208: 634},
		{22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 57: 372, 113: 334, 117: 353, 356, 352, 333, 123: 620, 335, 126: 332, 138: 331, 144: 621, 219: 622},
		{163, 405, 163, 8: 236, 133: 616, 136: 615, 627, 170: 625, 196: 626, 624},
		{22: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 43: 68, 68, 68, 68, 68, 68, 68, 52: 68, 55: 68, 57: 68},
		{22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 52: 64, 55: 64, 57: 372, 113: 334, 117: 353, 356, 352, 333, 123: 620, 335, 126: 332, 138: 331, 144: 623},
		// 300
		{22: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 43: 67, 67, 67, 67, 67, 67, 67, 52: 67, 55: 67, 57: 67},
		{8: 633},
		{8: 238, 10: 238},
		{8: 235, 10: 631},
		{8: 234, 10: 234, 51: 233, 179: 628},
		// 305
		{51: 629},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 588, 587, 171: 630},
		{8: 232, 10: 232},
		{163, 405, 163, 133: 616, 136: 615, 627, 170: 632},
		{8: 237, 10: 237},
		// 310
		{246, 246, 246, 246, 246, 246, 246, 246, 246, 11: 246, 13: 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 43: 246, 246, 246, 246, 246, 246, 246, 246, 52: 246, 55: 246, 57: 246, 85: 246, 87: 246, 246, 246, 246, 246, 246, 246, 246, 246, 246},
		{52: 72, 55: 70, 176: 636, 637, 192: 635},
		{73, 73, 73, 8: 73, 22: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 43: 73, 73, 73, 73, 73, 73, 73, 73, 55: 73, 57: 73},
		{52: 683, 128: 684},
		{55: 639, 122: 640, 127: 638},
		// 315
		{8: 682},
		{60, 22: 609, 85: 60, 203: 641},
		{51, 51, 51, 51, 51, 51, 51, 51, 51, 11: 51, 13: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 43: 51, 51, 51, 51, 51, 51, 51, 52: 51, 55: 51, 57: 51, 85: 51, 87: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 112: 51},
		{642, 85: 643},
		{11: 611, 163: 668},
		// 320
		{644},
		{11: 611, 163: 645},
		{11: 613, 42: 646},
		{42: 647},
		{11: 55, 651, 149: 649, 648, 157: 650},
		// 325
		{11: 664},
		{9: 57, 57, 42: 57},
		{10: 654, 42: 655},
		{2: 652},
		{54: 653},
		// 330
		{11: 54},
		{11: 55, 651, 149: 663, 648},
		{11: 656, 182: 657},
		{9: 53, 53, 42: 53},
		{10: 658, 42: 659},
		// 335
		{11: 662},
		{2: 534, 169: 660},
		{9: 661, 536},
		{47, 47, 47, 47, 47, 47, 47, 47, 47, 11: 47, 13: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 43: 47, 47, 47, 47, 47, 47, 47, 52: 47, 55: 47, 57: 47, 85: 47, 87: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 112: 47},
		{9: 52, 52, 42: 52},
		// 340
		{9: 56, 56, 42: 56},
		{665},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 666},
		{424, 429, 3: 442, 432, 433, 428, 427, 9: 667, 12: 423, 51: 448, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{9: 58, 58, 42: 58},
		// 345
		{9: 614, 11: 613, 42: 669},
		{9: 671, 11: 55, 651, 42: 672, 149: 649, 648, 157: 670},
		{9: 675, 654, 42: 676},
		{46, 46, 46, 46, 46, 46, 46, 46, 46, 11: 46, 13: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 43: 46, 46, 46, 46, 46, 46, 46, 52: 46, 55: 46, 57: 46, 85: 46, 87: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 112: 46},
		{11: 55, 651, 149: 649, 648, 157: 673},
		// 350
		{9: 674, 654},
		{45, 45, 45, 45, 45, 45, 45, 45, 45, 11: 45, 13: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 43: 45, 45, 45, 45, 45, 45, 45, 52: 45, 55: 45, 57: 45, 85: 45, 87: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 112: 45},
		{50, 50, 50, 50, 50, 50, 50, 50, 50, 11: 50, 13: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 43: 50, 50, 50, 50, 50, 50, 50, 52: 50, 55: 50, 57: 50, 85: 50, 87: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 112: 50},
		{11: 55, 651, 149: 649, 648, 157: 677},
		{9: 678, 654, 42: 679},
		// 355
		{49, 49, 49, 49, 49, 49, 49, 49, 49, 11: 49, 13: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 43: 49, 49, 49, 49, 49, 49, 49, 52: 49, 55: 49, 57: 49, 85: 49, 87: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 112: 49},
		{11: 656, 182: 680},
		{9: 681, 658},
		{48, 48, 48, 48, 48, 48, 48, 48, 48, 11: 48, 13: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 43: 48, 48, 48, 48, 48, 48, 48, 52: 48, 55: 48, 57: 48, 85: 48, 87: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 112: 48},
		{69, 69, 69, 8: 69, 22: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 43: 69, 69, 69, 69, 69, 69, 69, 69, 55: 69, 57: 69},
		// 360
		{103, 103, 103, 103, 103, 103, 103, 103, 103, 11: 103, 13: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 43: 103, 103, 103, 103, 103, 103, 103, 52: 103, 55: 103, 57: 103, 85: 103, 87: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 206: 685},
		{71, 71, 71, 8: 71, 22: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 43: 71, 71, 71, 71, 71, 71, 71, 71, 55: 71, 57: 71},
		{387, 392, 686, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 99, 43: 338, 339, 337, 363, 364, 340, 336, 52: 683, 471, 55: 639, 57: 372, 85: 708, 87: 705, 710, 695, 709, 696, 706, 707, 703, 711, 704, 113: 334, 687, 117: 353, 356, 352, 333, 702, 640, 620, 335, 126: 332, 694, 689, 690, 692, 693, 688, 134: 691, 701, 138: 331, 144: 700, 181: 698, 216: 699, 697},
		{311, 311, 3: 311, 311, 311, 311, 311, 311, 10: 311, 12: 311, 42: 760, 51: 311, 56: 311, 58: 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311},
		{8: 249, 249, 474},
		// 365
		{113, 113, 113, 113, 113, 113, 113, 113, 113, 11: 113, 13: 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 43: 113, 113, 113, 113, 113, 113, 113, 52: 113, 55: 113, 57: 113, 85: 113, 87: 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 112: 113},
		{112, 112, 112, 112, 112, 112, 112, 112, 112, 11: 112, 13: 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 43: 112, 112, 112, 112, 112, 112, 112, 52: 112, 55: 112, 57: 112, 85: 112, 87: 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112: 112},
		{111, 111, 111, 111, 111, 111, 111, 111, 111, 11: 111, 13: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 43: 111, 111, 111, 111, 111, 111, 111, 52: 111, 55: 111, 57: 111, 85: 111, 87: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 112: 111},
		{110, 110, 110, 110, 110, 110, 110, 110, 110, 11: 110, 13: 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 43: 110, 110, 110, 110, 110, 110, 110, 52: 110, 55: 110, 57: 110, 85: 110, 87: 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 112: 110},
		{109, 109, 109, 109, 109, 109, 109, 109, 109, 11: 109, 13: 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 43: 109, 109, 109, 109, 109, 109, 109, 52: 109, 55: 109, 57: 109, 85: 109, 87: 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 112: 109},
		// 370
		{108, 108, 108, 108, 108, 108, 108, 108, 108, 11: 108, 13: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 43: 108, 108, 108, 108, 108, 108, 108, 52: 108, 55: 108, 57: 108, 85: 108, 87: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 112: 108},
		{107, 107, 107, 107, 107, 107, 107, 107, 107, 11: 107, 13: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 43: 107, 107, 107, 107, 107, 107, 107, 52: 107, 55: 107, 57: 107, 85: 107, 87: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 112: 107},
		{248, 248, 248, 248, 248, 248, 248, 248, 11: 248, 13: 248, 248, 248, 248, 248, 248, 248, 248, 248, 145: 374, 757},
		{42: 755},
		{41: 754},
		// 375
		{101, 101, 101, 101, 101, 101, 101, 101, 101, 11: 101, 13: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 43: 101, 101, 101, 101, 101, 101, 101, 52: 101, 55: 101, 57: 101, 85: 101, 87: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101},
		{387, 392, 686, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 98, 43: 338, 339, 337, 363, 364, 340, 336, 52: 683, 471, 55: 639, 57: 372, 85: 708, 87: 705, 710, 695, 709, 696, 706, 707, 703, 711, 704, 113: 334, 687, 117: 353, 356, 352, 333, 702, 640, 620, 335, 126: 332, 694, 689, 690, 692, 693, 688, 134: 691, 701, 138: 331, 144: 700, 181: 753},
		{97, 97, 97, 97, 97, 97, 97, 97, 97, 11: 97, 13: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 43: 97, 97, 97, 97, 97, 97, 97, 52: 97, 55: 97, 57: 97, 85: 97, 87: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97},
		{96, 96, 96, 96, 96, 96, 96, 96, 96, 11: 96, 13: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 43: 96, 96, 96, 96, 96, 96, 96, 52: 96, 55: 96, 57: 96, 85: 96, 87: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96},
		{8: 752},
		// 380
		{746},
		{742},
		{738},
		{387, 392, 686, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 683, 471, 55: 639, 85: 708, 87: 705, 710, 695, 709, 696, 706, 707, 703, 711, 704, 114: 687, 121: 702, 640, 127: 694, 689, 690, 692, 693, 688, 134: 691, 732},
		{718},
		// 385
		{2: 716},
		{8: 715},
		{8: 714},
		{387, 392, 380, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 471, 114: 687, 121: 712},
		{8: 713},
		// 390
		{84, 84, 84, 84, 84, 84, 84, 84, 84, 11: 84, 13: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 43: 84, 84, 84, 84, 84, 84, 84, 52: 84, 55: 84, 57: 84, 85: 84, 87: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 112: 84},
		{85, 85, 85, 85, 85, 85, 85, 85, 85, 11: 85, 13: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 43: 85, 85, 85, 85, 85, 85, 85, 52: 85, 55: 85, 57: 85, 85: 85, 87: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 112: 85},
		{86, 86, 86, 86, 86, 86, 86, 86, 86, 11: 86, 13: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 43: 86, 86, 86, 86, 86, 86, 86, 52: 86, 55: 86, 57: 86, 85: 86, 87: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 112: 86},
		{8: 717},
		{87, 87, 87, 87, 87, 87, 87, 87, 87, 11: 87, 13: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 43: 87, 87, 87, 87, 87, 87, 87, 52: 87, 55: 87, 57: 87, 85: 87, 87: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 112: 87},
		// 395
		{387, 392, 380, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 43: 338, 339, 337, 363, 364, 340, 336, 53: 471, 57: 372, 113: 334, 687, 117: 353, 356, 352, 333, 719, 123: 620, 335, 126: 332, 138: 331, 144: 720},
		{8: 726},
		{387, 392, 380, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 471, 114: 687, 121: 721},
		{8: 722},
		{387, 392, 380, 391, 393, 394, 390, 389, 9: 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 471, 114: 687, 121: 723},
		// 400
		{9: 724},
		{387, 392, 686, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 683, 471, 55: 639, 85: 708, 87: 705, 710, 695, 709, 696, 706, 707, 703, 711, 704, 114: 687, 121: 702, 640, 127: 694, 689, 690, 692, 693, 688, 134: 691, 725},
		{88, 88, 88, 88, 88, 88, 88, 88, 88, 11: 88, 13: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 43: 88, 88, 88, 88, 88, 88, 88, 52: 88, 55: 88, 57: 88, 85: 88, 87: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 112: 88},
		{387, 392, 380, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 471, 114: 687, 121: 727},
		{8: 728},
		// 405
		{387, 392, 380, 391, 393, 394, 390, 389, 9: 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 471, 114: 687, 121: 729},
		{9: 730},
		{387, 392, 686, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 683, 471, 55: 639, 85: 708, 87: 705, 710, 695, 709, 696, 706, 707, 703, 711, 704, 114: 687, 121: 702, 640, 127: 694, 689, 690, 692, 693, 688, 134: 691, 731},
		{89, 89, 89, 89, 89, 89, 89, 89, 89, 11: 89, 13: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 43: 89, 89, 89, 89, 89, 89, 89, 52: 89, 55: 89, 57: 89, 85: 89, 87: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 112: 89},
		{87: 733},
		// 410
		{734},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 471, 114: 735},
		{9: 736, 474},
		{8: 737},
		{90, 90, 90, 90, 90, 90, 90, 90, 90, 11: 90, 13: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 43: 90, 90, 90, 90, 90, 90, 90, 52: 90, 55: 90, 57: 90, 85: 90, 87: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 112: 90},
		// 415
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 471, 114: 739},
		{9: 740, 474},
		{387, 392, 686, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 683, 471, 55: 639, 85: 708, 87: 705, 710, 695, 709, 696, 706, 707, 703, 711, 704, 114: 687, 121: 702, 640, 127: 694, 689, 690, 692, 693, 688, 134: 691, 741},
		{91, 91, 91, 91, 91, 91, 91, 91, 91, 11: 91, 13: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 43: 91, 91, 91, 91, 91, 91, 91, 52: 91, 55: 91, 57: 91, 85: 91, 87: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 112: 91},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 471, 114: 743},
		// 420
		{9: 744, 474},
		{387, 392, 686, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 683, 471, 55: 639, 85: 708, 87: 705, 710, 695, 709, 696, 706, 707, 703, 711, 704, 114: 687, 121: 702, 640, 127: 694, 689, 690, 692, 693, 688, 134: 691, 745},
		{92, 92, 92, 92, 92, 92, 92, 92, 92, 11: 92, 13: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 43: 92, 92, 92, 92, 92, 92, 92, 52: 92, 55: 92, 57: 92, 85: 92, 87: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 112: 92},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 53: 471, 114: 747},
		{9: 748, 474},
		// 425
		{387, 392, 686, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 683, 471, 55: 639, 85: 708, 87: 705, 710, 695, 709, 696, 706, 707, 703, 711, 704, 114: 687, 121: 702, 640, 127: 694, 689, 690, 692, 693, 688, 134: 691, 749},
		{94, 94, 94, 94, 94, 94, 94, 94, 94, 11: 94, 13: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 43: 94, 94, 94, 94, 94, 94, 94, 52: 94, 55: 94, 57: 94, 85: 94, 87: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 112: 750},
		{387, 392, 686, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 683, 471, 55: 639, 85: 708, 87: 705, 710, 695, 709, 696, 706, 707, 703, 711, 704, 114: 687, 121: 702, 640, 127: 694, 689, 690, 692, 693, 688, 134: 691, 751},
		{93, 93, 93, 93, 93, 93, 93, 93, 93, 11: 93, 13: 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 43: 93, 93, 93, 93, 93, 93, 93, 52: 93, 55: 93, 57: 93, 85: 93, 87: 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 112: 93},
		{95, 95, 95, 95, 95, 95, 95, 95, 95, 11: 95, 13: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 43: 95, 95, 95, 95, 95, 95, 95, 52: 95, 55: 95, 57: 95, 85: 95, 87: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 112: 95},
		// 430
		{100, 100, 100, 100, 100, 100, 100, 100, 100, 11: 100, 13: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 43: 100, 100, 100, 100, 100, 100, 100, 52: 100, 55: 100, 57: 100, 85: 100, 87: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
		{102, 102, 102, 102, 102, 102, 102, 102, 102, 11: 102, 13: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 43: 102, 102, 102, 102, 102, 102, 102, 102, 52: 102, 55: 102, 57: 102, 85: 102, 87: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 112: 102},
		{387, 392, 686, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 683, 471, 55: 639, 85: 708, 87: 705, 710, 695, 709, 696, 706, 707, 703, 711, 704, 114: 687, 121: 702, 640, 127: 694, 689, 690, 692, 693, 688, 134: 691, 756},
		{104, 104, 104, 104, 104, 104, 104, 104, 104, 11: 104, 13: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 43: 104, 104, 104, 104, 104, 104, 104, 52: 104, 55: 104, 57: 104, 85: 104, 87: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 112: 104},
		{42: 758},
		// 435
		{387, 392, 686, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 683, 471, 55: 639, 85: 708, 87: 705, 710, 695, 709, 696, 706, 707, 703, 711, 704, 114: 687, 121: 702, 640, 127: 694, 689, 690, 692, 693, 688, 134: 691, 759},
		{105, 105, 105, 105, 105, 105, 105, 105, 105, 11: 105, 13: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 43: 105, 105, 105, 105, 105, 105, 105, 52: 105, 55: 105, 57: 105, 85: 105, 87: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 112: 105},
		{387, 392, 686, 391, 393, 394, 390, 389, 250, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 52: 683, 471, 55: 639, 85: 708, 87: 705, 710, 695, 709, 696, 706, 707, 703, 711, 704, 114: 687, 121: 702, 640, 127: 694, 689, 690, 692, 693, 688, 134: 691, 761},
		{106, 106, 106, 106, 106, 106, 106, 106, 106, 11: 106, 13: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 43: 106, 106, 106, 106, 106, 106, 106, 52: 106, 55: 106, 57: 106, 85: 106, 87: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 112: 106},
		{163, 405, 163, 133: 616, 136: 615, 528},
		// 440
		{79, 79, 79, 8: 79, 22: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 43: 79, 79, 79, 79, 79, 79, 79, 79, 55: 79, 57: 79},
		{52: 189, 212: 766},
		{187, 187, 187, 8: 187, 187, 187, 12: 187, 22: 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 42: 187, 187, 187, 187, 187, 187, 187, 187, 52: 144},
		{52: 767},
		{2: 768, 189: 771, 770, 225: 769},
		// 445
		{10: 316, 41: 316, 51: 316},
		{10: 774, 41: 191, 164: 775},
		{10: 186, 41: 186},
		{10: 184, 41: 184, 51: 772},
		{248, 248, 248, 248, 248, 248, 248, 248, 11: 248, 13: 248, 248, 248, 248, 248, 248, 248, 248, 248, 145: 374, 773},
		// 450
		{10: 183, 41: 183},
		{2: 768, 41: 190, 189: 771, 777},
		{41: 776},
		{188, 188, 188, 8: 188, 188, 188, 12: 188, 22: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 42: 188, 188, 188, 188, 188, 188, 188, 188},
		{10: 185, 41: 185},
		// 455
		{52: 780},
		{208, 208, 208, 8: 208, 208, 208, 12: 208, 22: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 42: 208, 208, 208, 208, 208, 208, 208, 208, 52: 144},
		{22: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 782, 57: 210, 211: 781},
		{22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 57: 372, 113: 403, 117: 353, 356, 352, 402, 138: 786, 153: 785, 201: 784, 230: 783},
		{207, 207, 207, 8: 207, 207, 207, 12: 207, 22: 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 42: 207, 207, 207, 207, 207, 207, 207, 207},
		// 460
		{22: 362, 360, 361, 350, 342, 351, 347, 359, 346, 344, 345, 343, 348, 357, 354, 355, 358, 349, 341, 797, 57: 372, 113: 403, 117: 353, 356, 352, 402, 138: 786, 153: 785, 201: 798},
		{22: 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 57: 204},
		{163, 405, 163, 8: 788, 42: 176, 133: 616, 136: 615, 790, 184: 791, 202: 789, 231: 787},
		{22: 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 57: 200},
		{8: 794, 10: 795},
		// 465
		{22: 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 57: 201},
		{8: 195, 10: 195},
		{8: 193, 10: 193, 42: 175},
		{42: 792},
		{248, 248, 248, 248, 248, 248, 248, 248, 11: 248, 13: 248, 248, 248, 248, 248, 248, 248, 248, 248, 145: 374, 793},
		// 470
		{8: 192, 10: 192},
		{22: 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 57: 202},
		{163, 405, 163, 42: 176, 133: 616, 136: 615, 790, 184: 791, 202: 796},
		{8: 194, 10: 194},
		{209, 209, 209, 8: 209, 209, 209, 12: 209, 22: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 42: 209, 209, 209, 209, 209, 209, 209, 209},
		// 475
		{22: 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 57: 203},
		{387, 392, 380, 391, 393, 394, 390, 389, 11: 386, 13: 396, 395, 398, 381, 382, 383, 384, 385, 397, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 53: 800, 155: 401, 161: 801},
		{424, 429, 3: 442, 432, 433, 428, 427, 9: 803, 12: 423, 51: 448, 56: 425, 58: 431, 430, 436, 437, 447, 443, 444, 452, 445, 456, 426, 450, 440, 439, 438, 434, 454, 451, 449, 441, 458, 446, 435, 455, 453, 457},
		{9: 802},
		{211, 211, 211, 8: 211, 211, 211, 12: 211, 22: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 42: 211, 211, 211, 211, 211, 211, 211, 211},
		// 480
		{212, 212, 212, 8: 212, 212, 212, 12: 212, 22: 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 42: 212, 212, 212, 212, 212, 212, 212, 212},
		{241, 241, 241, 8: 241, 241, 241, 12: 241},
		{239, 239, 239, 8: 239, 239, 239, 12: 239},
		{242, 242, 242, 8: 242, 242, 242, 12: 242},
		{243, 243, 243, 8: 243, 243, 243, 12: 243},
		// 485
		{244, 244, 244, 8: 244, 244, 244, 12: 244},
		{8: 234, 10: 234, 22: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 43: 65, 65, 65, 65, 65, 65, 65, 51: 233, 66, 55: 66, 57: 65, 178: 619, 628, 183: 810},
		{52: 77, 55: 77, 207: 811},
		{52: 72, 55: 70, 176: 636, 637, 192: 812},
		{76, 76, 76, 8: 76, 22: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 43: 76, 76, 76, 76, 76, 76, 76, 76, 55: 76, 57: 76},
		// 490
		{82, 82, 82, 8: 82, 22: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 43: 82, 82, 82, 82, 82, 82, 82, 82, 55: 82, 57: 82},
		{248, 248, 248, 248, 248, 248, 248, 248, 11: 248, 13: 248, 248, 248, 248, 248, 248, 248, 248, 248, 145: 374, 815},
		{50: 319},
		{84: 838, 86: 840, 100: 828, 829, 830, 825, 826, 827, 831, 835, 832, 822, 833, 834, 115: 839, 837, 125: 836, 139: 820, 819, 824, 821, 823, 148: 818, 228: 817},
		{50: 321},
		// 495
		{50: 43, 84: 838, 86: 840, 100: 828, 829, 830, 825, 826, 827, 831, 835, 832, 822, 833, 834, 115: 839, 837, 125: 836, 139: 820, 878, 824, 821, 823},
		{50: 42, 84: 42, 86: 42, 97: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{50: 38, 84: 38, 86: 38, 97: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{50: 37, 84: 37, 86: 37, 97: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{86: 840, 115: 900, 837},
		// 500
		{50: 35, 84: 35, 86: 35, 97: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{97: 28, 28, 888, 188: 886, 220: 887, 885},
		{86: 840, 115: 882, 837},
		{2: 879},
		{2: 874},
		// 505
		{2: 855, 84: 857, 226: 856},
		{84: 838, 86: 840, 115: 839, 837, 125: 854},
		{50: 16, 84: 16, 86: 16, 97: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{86: 840, 115: 852, 837},
		{86: 840, 115: 850, 837},
		// 510
		{84: 838, 86: 840, 115: 839, 837, 125: 849},
		{2: 845},
		{86: 840, 115: 843, 837},
		{50: 7, 84: 7, 86: 7, 97: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{84: 5, 86: 842},
		// 515
		{50: 4, 84: 4, 86: 4, 97: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{84: 841},
		{84: 2, 86: 2},
		{50: 3, 84: 3, 86: 3, 97: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{84: 1, 86: 1},
		// 520
		{84: 844},
		{50: 8, 84: 8, 86: 8, 97: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{84: 846, 86: 840, 115: 847, 837},
		{50: 12, 84: 12, 86: 12, 97: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{84: 848},
		// 525
		{50: 9, 84: 9, 86: 9, 97: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{50: 13, 84: 13, 86: 13, 97: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{84: 851},
		{50: 14, 84: 14, 86: 14, 97: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{84: 853},
		// 530
		{50: 15, 84: 15, 86: 15, 97: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{50: 17, 84: 17, 86: 17, 97: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{84: 838, 86: 840, 115: 839, 837, 125: 863, 152: 873},
		{2: 534, 9: 147, 147: 859, 169: 858, 194: 860},
		{50: 10, 84: 10, 86: 10, 97: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		// 535
		{9: 146, 866, 147: 867},
		{9: 864},
		{9: 861},
		{84: 838, 86: 840, 115: 839, 837, 125: 863, 152: 862},
		{50: 18, 84: 18, 86: 18, 97: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		// 540
		{50: 6, 84: 6, 86: 6, 97: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{84: 838, 86: 840, 115: 839, 837, 125: 863, 152: 865},
		{50: 20, 84: 20, 86: 20, 97: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{2: 537, 147: 870},
		{9: 868},
		// 545
		{84: 838, 86: 840, 115: 839, 837, 125: 863, 152: 869},
		{50: 11, 84: 11, 86: 11, 97: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{9: 871},
		{84: 838, 86: 840, 115: 839, 837, 125: 863, 152: 872},
		{50: 19, 84: 19, 86: 19, 97: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		// 550
		{50: 21, 84: 21, 86: 21, 97: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{84: 875},
		{84: 838, 86: 840, 97: 40, 40, 40, 828, 829, 830, 825, 826, 827, 831, 835, 832, 822, 833, 834, 115: 839, 837, 125: 836, 139: 820, 819, 824, 821, 823, 148: 876, 151: 877},
		{84: 838, 86: 840, 97: 39, 39, 39, 828, 829, 830, 825, 826, 827, 831, 835, 832, 822, 833, 834, 115: 839, 837, 125: 836, 139: 820, 878, 824, 821, 823},
		{97: 31, 31, 31},
		// 555
		{50: 41, 84: 41, 86: 41, 97: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		{84: 880},
		{84: 838, 86: 840, 97: 40, 40, 40, 828, 829, 830, 825, 826, 827, 831, 835, 832, 822, 833, 834, 115: 839, 837, 125: 836, 139: 820, 819, 824, 821, 823, 148: 876, 151: 881},
		{97: 32, 32, 32},
		{84: 883},
		// 560
		{84: 838, 86: 840, 97: 40, 40, 40, 828, 829, 830, 825, 826, 827, 831, 835, 832, 822, 833, 834, 115: 839, 837, 125: 836, 139: 820, 819, 824, 821, 823, 148: 876, 151: 884},
		{97: 33, 33, 33},
		{97: 24, 894, 222: 895, 893},
		{97: 30, 30, 30},
		{97: 27, 27, 888, 188: 892},
		// 565
		{86: 840, 115: 889, 837},
		{84: 890},
		{84: 838, 86: 840, 97: 40, 40, 40, 828, 829, 830, 825, 826, 827, 831, 835, 832, 822, 833, 834, 115: 839, 837, 125: 836, 139: 820, 819, 824, 821, 823, 148: 876, 151: 891},
		{97: 26, 26, 26},
		{97: 29, 29, 29},
		// 570
		{97: 899, 224: 898},
		{84: 896},
		{97: 23},
		{84: 838, 86: 840, 97: 40, 100: 828, 829, 830, 825, 826, 827, 831, 835, 832, 822, 833, 834, 115: 839, 837, 125: 836, 139: 820, 819, 824, 821, 823, 148: 876, 151: 897},
		{97: 25},
		// 575
		{50: 34, 84: 34, 86: 34, 97: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{50: 22, 84: 22, 86: 22, 97: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{84: 901},
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
			yyVAL.node = &Declaration{
				Case: 1,
				StaticAssertDeclaration: yyS[yypt-0].node.(*StaticAssertDeclaration),
			}
		}
	case 79:
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
	case 80:
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
	case 81:
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
	case 82:
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
	case 83:
		{
			yyVAL.node = (*DeclarationSpecifiersOpt)(nil)
		}
	case 84:
		{
			lhs := &DeclarationSpecifiersOpt{
				DeclarationSpecifiers: yyS[yypt-0].node.(*DeclarationSpecifiers),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.DeclarationSpecifiers.attr
			lhs.typeSpecifier = lhs.DeclarationSpecifiers.typeSpecifier
		}
	case 85:
		{
			yyVAL.node = &InitDeclaratorList{
				InitDeclarator: yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 86:
		{
			yyVAL.node = &InitDeclaratorList{
				Case:               1,
				InitDeclaratorList: yyS[yypt-2].node.(*InitDeclaratorList),
				Token:              yyS[yypt-1].Token,
				InitDeclarator:     yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 87:
		{
			yyVAL.node = (*InitDeclaratorListOpt)(nil)
		}
	case 88:
		{
			yyVAL.node = &InitDeclaratorListOpt{
				InitDeclaratorList: yyS[yypt-0].node.(*InitDeclaratorList).reverse(),
			}
		}
	case 89:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
		}
	case 90:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-0].node.(*Declarator)
			d.setFull(lx)
		}
	case 91:
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
	case 92:
		{
			lhs := &StorageClassSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saTypedef
		}
	case 93:
		{
			lhs := &StorageClassSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saExtern
		}
	case 94:
		{
			lhs := &StorageClassSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saStatic
		}
	case 95:
		{
			lhs := &StorageClassSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saAuto
		}
	case 96:
		{
			lhs := &StorageClassSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRegister
		}
	case 97:
		{
			lhs := &TypeSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsVoid)
		}
	case 98:
		{
			lhs := &TypeSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsChar)
		}
	case 99:
		{
			lhs := &TypeSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsShort)
		}
	case 100:
		{
			lhs := &TypeSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsInt)
		}
	case 101:
		{
			lhs := &TypeSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsLong)
		}
	case 102:
		{
			lhs := &TypeSpecifier{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsFloat)
		}
	case 103:
		{
			lhs := &TypeSpecifier{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsDouble)
		}
	case 104:
		{
			lhs := &TypeSpecifier{
				Case:  7,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsSigned)
		}
	case 105:
		{
			lhs := &TypeSpecifier{
				Case:  8,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsUnsigned)
		}
	case 106:
		{
			lhs := &TypeSpecifier{
				Case:  9,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsBool)
		}
	case 107:
		{
			lhs := &TypeSpecifier{
				Case:  10,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsComplex)
		}
	case 108:
		{
			lhs := &TypeSpecifier{
				Case: 11,
				StructOrUnionSpecifier: yyS[yypt-0].node.(*StructOrUnionSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(lhs.StructOrUnionSpecifier.typeSpecifiers())
		}
	case 109:
		{
			lhs := &TypeSpecifier{
				Case:          12,
				EnumSpecifier: yyS[yypt-0].node.(*EnumSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsEnumSpecifier)
		}
	case 110:
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
	case 111:
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
	case 112:
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
	case 113:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-1].node.(*IdentifierOpt); o != nil {
				lx.scope.declareStructTag(o.Token, lx.report)
			}
			lx.pushScope(ScopeMembers)
			lx.scope.isUnion = yyS[yypt-2].node.(*StructOrUnion).Case == 1 // "union"
			lx.scope.prevStructDeclarator = nil
		}
	case 114:
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
	case 117:
		{
			yyVAL.node = &StructOrUnion{
				Token: yyS[yypt-0].Token,
			}
		}
	case 118:
		{
			yyVAL.node = &StructOrUnion{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 119:
		{
			yyVAL.node = &StructDeclarationList{
				StructDeclaration: yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 120:
		{
			yyVAL.node = &StructDeclarationList{
				Case: 1,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList),
				StructDeclaration:     yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 121:
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
	case 122:
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
	case 123:
		{
			yyVAL.node = &StructDeclaration{
				Case: 2,
				StaticAssertDeclaration: yyS[yypt-0].node.(*StaticAssertDeclaration),
			}
		}
	case 124:
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
	case 125:
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
	case 126:
		{
			yyVAL.node = (*SpecifierQualifierListOpt)(nil)
		}
	case 127:
		{
			lhs := &SpecifierQualifierListOpt{
				SpecifierQualifierList: yyS[yypt-0].node.(*SpecifierQualifierList),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.SpecifierQualifierList.attr
			lhs.typeSpecifier = lhs.SpecifierQualifierList.typeSpecifier
		}
	case 128:
		{
			yyVAL.node = &StructDeclaratorList{
				StructDeclarator: yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 129:
		{
			yyVAL.node = &StructDeclaratorList{
				Case:                 1,
				StructDeclaratorList: yyS[yypt-2].node.(*StructDeclaratorList),
				Token:                yyS[yypt-1].Token,
				StructDeclarator:     yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 130:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.post(lx)
		}
	case 131:
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
	case 132:
		{
			yyVAL.node = (*CommaOpt)(nil)
		}
	case 133:
		{
			yyVAL.node = &CommaOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 134:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareEnumTag(o.Token, lx.report)
			}
			lx.iota = 0
		}
	case 135:
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
	case 136:
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
	case 137:
		{
			yyVAL.node = &EnumeratorList{
				Enumerator: yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 138:
		{
			yyVAL.node = &EnumeratorList{
				Case:           1,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList),
				Token:          yyS[yypt-1].Token,
				Enumerator:     yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 139:
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
	case 140:
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
	case 141:
		{
			lhs := &TypeQualifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saConst
		}
	case 142:
		{
			lhs := &TypeQualifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRestrict
		}
	case 143:
		{
			lhs := &TypeQualifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saVolatile
		}
	case 144:
		{
			lhs := &FunctionSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saInline
		}
	case 145:
		{
			lhs := &FunctionSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saNoreturn
		}
	case 146:
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
	case 147:
		{
			yyVAL.node = (*DeclaratorOpt)(nil)
		}
	case 148:
		{
			yyVAL.node = &DeclaratorOpt{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
		}
	case 149:
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
	case 150:
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
	case 151:
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
	case 152:
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
	case 153:
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
	case 154:
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
	case 155:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 156:
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
	case 157:
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
	case 158:
		{
			yyVAL.node = &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].node.(*TypeQualifierListOpt),
			}
		}
	case 159:
		{
			yyVAL.node = &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].node.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].node.(*Pointer),
			}
		}
	case 160:
		{
			yyVAL.node = (*PointerOpt)(nil)
		}
	case 161:
		{
			yyVAL.node = &PointerOpt{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
		}
	case 162:
		{
			lhs := &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.TypeQualifier.attr
		}
	case 163:
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
	case 164:
		{
			yyVAL.node = (*TypeQualifierListOpt)(nil)
		}
	case 165:
		{
			yyVAL.node = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].node.(*TypeQualifierList).reverse(),
			}
		}
	case 166:
		{
			lhs := &ParameterTypeList{
				ParameterList: yyS[yypt-0].node.(*ParameterList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 167:
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
	case 168:
		{
			yyVAL.node = (*ParameterTypeListOpt)(nil)
		}
	case 169:
		{
			yyVAL.node = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].node.(*ParameterTypeList),
			}
		}
	case 170:
		{
			yyVAL.node = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 171:
		{
			yyVAL.node = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].node.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 172:
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
	case 173:
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
	case 174:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 175:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 176:
		{
			yyVAL.node = (*IdentifierListOpt)(nil)
		}
	case 177:
		{
			yyVAL.node = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 178:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 179:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 180:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeBlock)
		}
	case 181:
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
	case 182:
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
	case 183:
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
	case 184:
		{
			yyVAL.node = (*AbstractDeclaratorOpt)(nil)
		}
	case 185:
		{
			yyVAL.node = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].node.(*AbstractDeclarator),
			}
		}
	case 186:
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
	case 187:
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
	case 188:
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
	case 189:
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
	case 190:
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
	case 191:
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
	case 192:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 193:
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
	case 194:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 195:
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
	case 196:
		{
			yyVAL.node = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 197:
		{
			yyVAL.node = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
		}
	case 198:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 199:
		{
			yyVAL.node = &Initializer{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 200:
		{
			yyVAL.node = &InitializerList{
				DesignationOpt: yyS[yypt-1].node.(*DesignationOpt),
				Initializer:    yyS[yypt-0].node.(*Initializer),
			}
		}
	case 201:
		{
			yyVAL.node = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].node.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].node.(*DesignationOpt),
				Initializer:     yyS[yypt-0].node.(*Initializer),
			}
		}
	case 202:
		{
			yyVAL.node = (*InitializerList)(nil)
		}
	case 203:
		{
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 204:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 205:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 206:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 207:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 208:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 209:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 210:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 211:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 212:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 213:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 214:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 215:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 216:
		{
			yyVAL.node = &Statement{
				Case:               6,
				AssemblerStatement: yyS[yypt-0].node.(*AssemblerStatement),
			}
		}
	case 217:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 218:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 219:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 220:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 221:
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
	case 222:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 223:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 224:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 225:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 226:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 227:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 228:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 229:
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
	case 230:
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
	case 231:
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
	case 232:
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
	case 233:
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
	case 234:
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
	case 235:
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
	case 236:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 237:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 238:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 239:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 240:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 241:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 242:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 243:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 244:
		{
			yyVAL.node = &ExternalDeclaration{
				Case: 2,
				BasicAssemblerStatement: yyS[yypt-1].node.(*BasicAssemblerStatement),
				Token: yyS[yypt-0].Token,
			}
		}
	case 245:
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
	case 246:
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
	case 247:
		{
			yyVAL.node = &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:          yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 248:
		{
			lx := yylex.(*lexer)
			lx.scope.specifier = &DeclarationSpecifiers{typeSpecifier: tsEncode(tsInt)}
		}
	case 249:
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
	case 250:
		{
			yyVAL.node = &FunctionDefinition{
				Case:               1,
				Declarator:         yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt: yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:       yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 251:
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
	case 252:
		{
			lhs := &FunctionBody{
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
			yyVAL.node = lhs
			lhs.scope = lhs.CompoundStatement.scope
		}
	case 253:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 254:
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
	case 255:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 256:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 257:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 258:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 259:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScopePos(lhs.Pos())
		}
	case 260:
		{
			yyVAL.node = &AssemblerInstructions{
				Token: yyS[yypt-0].Token,
			}
		}
	case 261:
		{
			yyVAL.node = &AssemblerInstructions{
				Case: 1,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions),
				Token: yyS[yypt-0].Token,
			}
		}
	case 262:
		{
			yyVAL.node = &BasicAssemblerStatement{
				Token:                 yyS[yypt-4].Token,
				VolatileOpt:           yyS[yypt-3].node.(*VolatileOpt),
				Token2:                yyS[yypt-2].Token,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-0].Token,
			}
		}
	case 263:
		{
			yyVAL.node = (*VolatileOpt)(nil)
		}
	case 264:
		{
			yyVAL.node = &VolatileOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 265:
		{
			yyVAL.node = &AssemblerOperand{
				AssemblerSymbolicNameOpt: yyS[yypt-4].node.(*AssemblerSymbolicNameOpt),
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
		}
	case 266:
		{
			yyVAL.node = &AssemblerOperands{
				AssemblerOperand: yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 267:
		{
			yyVAL.node = &AssemblerOperands{
				Case:              1,
				AssemblerOperands: yyS[yypt-2].node.(*AssemblerOperands),
				Token:             yyS[yypt-1].Token,
				AssemblerOperand:  yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 268:
		{
			yyVAL.node = (*AssemblerSymbolicNameOpt)(nil)
		}
	case 269:
		{
			yyVAL.node = &AssemblerSymbolicNameOpt{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 270:
		{
			yyVAL.node = &Clobbers{
				Token: yyS[yypt-0].Token,
			}
		}
	case 271:
		{
			yyVAL.node = &Clobbers{
				Case:     1,
				Clobbers: yyS[yypt-2].node.(*Clobbers),
				Token:    yyS[yypt-1].Token,
				Token2:   yyS[yypt-0].Token,
			}
		}
	case 272:
		{
			yyVAL.node = &AssemblerStatement{
				BasicAssemblerStatement: yyS[yypt-0].node.(*BasicAssemblerStatement),
			}
		}
	case 273:
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
	case 274:
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
	case 275:
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
	case 276:
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
	case 277:
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
	case 278:
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
	case 279:
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
	case 280:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 281:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 282:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 283:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 284:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 285:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 286:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 287:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 288:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 289:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 290:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 291:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 292:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 293:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 294:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 295:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 296:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 297:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 298:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 299:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 300:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 301:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 302:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 303:
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
	case 304:
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
	case 305:
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
	case 306:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 307:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 308:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 309:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 310:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 311:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 312:
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
	case 313:
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
	case 314:
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
	case 315:
		{
			yyVAL.node = &ControlLine{
				Case:        13,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 318:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 319:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
