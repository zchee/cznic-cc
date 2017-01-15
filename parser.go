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
	yyTabOfs   = -321
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (322x)
		42:      1,   // '*' (283x)
		57377:   2,   // IDENTIFIER (236x)
		38:      3,   // '&' (215x)
		43:      4,   // '+' (215x)
		45:      5,   // '-' (215x)
		57363:   6,   // DEC (215x)
		57381:   7,   // INC (215x)
		59:      8,   // ';' (212x)
		41:      9,   // ')' (197x)
		44:      10,  // ',' (183x)
		57427:   11,  // STRINGLITERAL (163x)
		91:      12,  // '[' (160x)
		33:      13,  // '!' (145x)
		126:     14,  // '~' (145x)
		57347:   15,  // ALIGNOF (145x)
		57358:   16,  // CHARCONST (145x)
		57373:   17,  // FLOATCONST (145x)
		57384:   18,  // INTCONST (145x)
		57387:   19,  // LONGCHARCONST (145x)
		57388:   20,  // LONGSTRINGLITERAL (145x)
		57424:   21,  // SIZEOF (145x)
		57438:   22,  // VOLATILE (139x)
		57360:   23,  // CONST (137x)
		57418:   24,  // RESTRICT (137x)
		57353:   25,  // BOOL (127x)
		57357:   26,  // CHAR (127x)
		57359:   27,  // COMPLEX (127x)
		57367:   28,  // DOUBLE (127x)
		57369:   29,  // ENUM (127x)
		57372:   30,  // FLOAT (127x)
		57383:   31,  // INT (127x)
		57386:   32,  // LONG (127x)
		57422:   33,  // SHORT (127x)
		57423:   34,  // SIGNED (127x)
		57428:   35,  // STRUCT (127x)
		57432:   36,  // TYPEDEFNAME (127x)
		57433:   37,  // TYPEOF (127x)
		57435:   38,  // UNION (127x)
		57436:   39,  // UNSIGNED (127x)
		57437:   40,  // VOID (127x)
		125:     41,  // '}' (123x)
		58:      42,  // ':' (120x)
		57425:   43,  // STATIC (118x)
		57352:   44,  // AUTO (112x)
		57371:   45,  // EXTERN (112x)
		57382:   46,  // INLINE (112x)
		57395:   47,  // NORETURN (112x)
		57417:   48,  // REGISTER (112x)
		57431:   49,  // TYPEDEF (112x)
		57344:   50,  // $end (103x)
		61:      51,  // '=' (90x)
		123:     52,  // '{' (88x)
		57502:   53,  // Expression (85x)
		93:      54,  // ']' (83x)
		57351:   55,  // ASM (82x)
		46:      56,  // '.' (79x)
		57426:   57,  // STATIC_ASSERT (77x)
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
		57376:   85,  // GOTO (53x)
		57413:   86,  // PPOTHER (52x)
		57439:   87,  // WHILE (51x)
		57354:   88,  // BREAK (50x)
		57355:   89,  // CASE (50x)
		57361:   90,  // CONTINUE (50x)
		57364:   91,  // DEFAULT (50x)
		57366:   92,  // DO (50x)
		57374:   93,  // FOR (50x)
		57380:   94,  // IF (50x)
		57419:   95,  // RETURN (50x)
		57430:   96,  // SWITCH (50x)
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
		57368:   112, // ELSE (30x)
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
		57512:   149, // GroupListOpt (5x)
		57537:   150, // ReplacementList (5x)
		57539:   151, // SpecifierQualifierList (5x)
		57555:   152, // TypeQualifierList (5x)
		57458:   153, // $@9 (4x)
		57460:   154, // AbstractDeclarator (4x)
		57465:   155, // AssemblerOperand (4x)
		57468:   156, // AssemblerSymbolicNameOpt (4x)
		57482:   157, // DeclarationSpecifiersOpt (4x)
		57487:   158, // Designator (4x)
		57530:   159, // ParameterDeclaration (4x)
		57553:   160, // TypeName (4x)
		57556:   161, // TypeQualifierListOpt (4x)
		57464:   162, // AssemblerInstructions (3x)
		57466:   163, // AssemblerOperands (3x)
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
		"GroupListOpt",
		"ReplacementList",
		"SpecifierQualifierList",
		"TypeQualifierList",
		"$@9",
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
		83:  {157, 0},
		84:  {157, 1},
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
		124: {151, 2},
		125: {151, 2},
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
		162: {152, 1},
		163: {152, 2},
		164: {161, 0},
		165: {161, 1},
		166: {173, 1},
		167: {173, 3},
		168: {199, 0},
		169: {199, 1},
		170: {172, 1},
		171: {172, 3},
		172: {159, 2},
		173: {159, 2},
		174: {169, 1},
		175: {169, 3},
		176: {194, 0},
		177: {194, 1},
		178: {195, 0},
		179: {195, 1},
		180: {153, 0},
		181: {160, 3},
		182: {154, 1},
		183: {154, 2},
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
		202: {165, 2},
		203: {166, 0},
		204: {166, 1},
		205: {167, 1},
		206: {167, 2},
		207: {158, 3},
		208: {158, 2},
		209: {135, 1},
		210: {135, 1},
		211: {135, 1},
		212: {135, 1},
		213: {135, 1},
		214: {135, 1},
		215: {135, 1},
		216: {132, 3},
		217: {132, 4},
		218: {132, 3},
		219: {206, 0},
		220: {128, 4},
		221: {216, 1},
		222: {216, 2},
		223: {217, 0},
		224: {217, 1},
		225: {181, 1},
		226: {181, 1},
		227: {129, 2},
		228: {134, 5},
		229: {134, 7},
		230: {134, 5},
		231: {130, 5},
		232: {130, 7},
		233: {130, 9},
		234: {130, 8},
		235: {131, 3},
		236: {131, 2},
		237: {131, 2},
		238: {131, 3},
		239: {233, 1},
		240: {233, 2},
		241: {191, 1},
		242: {191, 1},
		243: {191, 2},
		244: {191, 1},
		245: {207, 0},
		246: {193, 5},
		247: {175, 0},
		248: {208, 0},
		249: {193, 5},
		250: {176, 0},
		251: {192, 2},
		252: {177, 0},
		253: {192, 3},
		254: {219, 1},
		255: {219, 2},
		256: {183, 0},
		257: {178, 0},
		258: {183, 2},
		259: {162, 1},
		260: {162, 2},
		261: {122, 5},
		262: {203, 0},
		263: {203, 1},
		264: {155, 5},
		265: {163, 1},
		266: {163, 3},
		267: {156, 0},
		268: {156, 3},
		269: {182, 1},
		270: {182, 3},
		271: {127, 1},
		272: {127, 7},
		273: {127, 9},
		274: {127, 11},
		275: {127, 13},
		276: {127, 6},
		277: {138, 7},
		278: {228, 1},
		279: {148, 1},
		280: {148, 2},
		281: {149, 0},
		282: {149, 1},
		283: {140, 1},
		284: {140, 1},
		285: {140, 3},
		286: {140, 1},
		287: {142, 4},
		288: {141, 4},
		289: {141, 4},
		290: {141, 4},
		291: {220, 1},
		292: {220, 2},
		293: {221, 0},
		294: {221, 1},
		295: {188, 4},
		296: {222, 3},
		297: {223, 0},
		298: {223, 1},
		299: {224, 1},
		300: {139, 3},
		301: {139, 5},
		302: {139, 7},
		303: {139, 5},
		304: {139, 2},
		305: {139, 1},
		306: {139, 3},
		307: {139, 3},
		308: {139, 2},
		309: {139, 3},
		310: {139, 6},
		311: {139, 2},
		312: {139, 4},
		313: {139, 3},
		314: {143, 1},
		315: {150, 1},
		316: {115, 1},
		317: {125, 1},
		318: {125, 2},
		319: {116, 1},
		320: {116, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 50}:   "invalid empty input",
		yyXError{569, -1}: "expected #endif",
		yyXError{571, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{489, -1}: "expected $end",
		yyXError{491, -1}: "expected $end",
		yyXError{32, -1}:  "expected '('",
		yyXError{49, -1}:  "expected '('",
		yyXError{75, -1}:  "expected '('",
		yyXError{285, -1}: "expected '('",
		yyXError{320, -1}: "expected '('",
		yyXError{341, -1}: "expected '('",
		yyXError{377, -1}: "expected '('",
		yyXError{378, -1}: "expected '('",
		yyXError{379, -1}: "expected '('",
		yyXError{381, -1}: "expected '('",
		yyXError{407, -1}: "expected '('",
		yyXError{54, -1}:  "expected ')'",
		yyXError{77, -1}:  "expected ')'",
		yyXError{84, -1}:  "expected ')'",
		yyXError{176, -1}: "expected ')'",
		yyXError{191, -1}: "expected ')'",
		yyXError{194, -1}: "expected ')'",
		yyXError{197, -1}: "expected ')'",
		yyXError{205, -1}: "expected ')'",
		yyXError{210, -1}: "expected ')'",
		yyXError{216, -1}: "expected ')'",
		yyXError{232, -1}: "expected ')'",
		yyXError{237, -1}: "expected ')'",
		yyXError{248, -1}: "expected ')'",
		yyXError{283, -1}: "expected ')'",
		yyXError{397, -1}: "expected ')'",
		yyXError{403, -1}: "expected ')'",
		yyXError{475, -1}: "expected ')'",
		yyXError{533, -1}: "expected ')'",
		yyXError{534, -1}: "expected ')'",
		yyXError{541, -1}: "expected ')'",
		yyXError{544, -1}: "expected ')'",
		yyXError{52, -1}:  "expected ','",
		yyXError{323, -1}: "expected ':'",
		yyXError{370, -1}: "expected ':'",
		yyXError{431, -1}: "expected ':'",
		yyXError{465, -1}: "expected ':'",
		yyXError{45, -1}:  "expected ';'",
		yyXError{55, -1}:  "expected ';'",
		yyXError{301, -1}: "expected ';'",
		yyXError{315, -1}: "expected ';'",
		yyXError{376, -1}: "expected ';'",
		yyXError{383, -1}: "expected ';'",
		yyXError{384, -1}: "expected ';'",
		yyXError{386, -1}: "expected ';'",
		yyXError{390, -1}: "expected ';'",
		yyXError{393, -1}: "expected ';'",
		yyXError{395, -1}: "expected ';'",
		yyXError{401, -1}: "expected ';'",
		yyXError{410, -1}: "expected ';'",
		yyXError{305, -1}: "expected '='",
		yyXError{89, -1}:  "expected '['",
		yyXError{513, -1}: "expected '\\n'",
		yyXError{517, -1}: "expected '\\n'",
		yyXError{521, -1}: "expected '\\n'",
		yyXError{524, -1}: "expected '\\n'",
		yyXError{526, -1}: "expected '\\n'",
		yyXError{548, -1}: "expected '\\n'",
		yyXError{553, -1}: "expected '\\n'",
		yyXError{556, -1}: "expected '\\n'",
		yyXError{563, -1}: "expected '\\n'",
		yyXError{568, -1}: "expected '\\n'",
		yyXError{574, -1}: "expected '\\n'",
		yyXError{95, -1}:  "expected ']'",
		yyXError{184, -1}: "expected ']'",
		yyXError{228, -1}: "expected ']'",
		yyXError{260, -1}: "expected ']'",
		yyXError{329, -1}: "expected ']'",
		yyXError{438, -1}: "expected '{'",
		yyXError{440, -1}: "expected '{'",
		yyXError{452, -1}: "expected '{'",
		yyXError{269, -1}: "expected '}'",
		yyXError{273, -1}: "expected '}'",
		yyXError{371, -1}: "expected '}'",
		yyXError{449, -1}: "expected '}'",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{204, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{88, -1}:  "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{287, -1}: "expected assembler instructions or string literal",
		yyXError{319, -1}: "expected assembler instructions or string literal",
		yyXError{321, -1}: "expected assembler instructions or string literal",
		yyXError{331, -1}: "expected assembler operand or one of ['[', string literal]",
		yyXError{346, -1}: "expected assembler operands or one of [')', '[', string literal]",
		yyXError{324, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{350, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{314, -1}: "expected assembler statement or asm",
		yyXError{373, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{332, -1}: "expected clobbers or string literal",
		yyXError{353, -1}: "expected clobbers or string literal",
		yyXError{313, -1}: "expected compound statement or '{'",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{50, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{257, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{369, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{446, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{466, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{488, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{296, -1}: "expected declaration list or one of [_Bool, _Complex, _Noreturn, _Static_assert, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{299, -1}: "expected declaration or one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{392, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{47, -1}:  "expected declarator or one of ['(', '*', identifier]",
		yyXError{436, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{196, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{254, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{199, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{85, -1}:  "expected direct abstract declarator or one of ['(', '[']",
		yyXError{292, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{561, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{567, -1}: "expected endif line or #endif",
		yyXError{498, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{559, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{441, -1}: "expected enumerator list or identifier",
		yyXError{448, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{100, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{124, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{408, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{412, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{416, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{420, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{64, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{245, -1}: "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{249, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		yyXError{92, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{227, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{284, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{51, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{66, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{67, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{68, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{69, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{70, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{71, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{72, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{73, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{74, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{98, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{122, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{123, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{125, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{126, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{127, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{128, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{129, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{130, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{131, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{132, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{133, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{134, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{135, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{150, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{151, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{178, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{185, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{221, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{224, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{342, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{96, -1}:  "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{219, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{473, -1}: "expected expression or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{6, -1}:   "expected external declaration or one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{295, -1}: "expected function body or one of ['{', asm]",
		yyXError{311, -1}: "expected function body or one of ['{', asm]",
		yyXError{484, -1}: "expected function body or one of ['{', asm]",
		yyXError{485, -1}: "expected function body or one of ['{', asm]",
		yyXError{483, -1}: "expected function body or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{294, -1}: "expected function body or optional declaration list or one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{550, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{492, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{102, -1}: "expected identifier",
		yyXError{103, -1}: "expected identifier",
		yyXError{213, -1}: "expected identifier",
		yyXError{258, -1}: "expected identifier",
		yyXError{328, -1}: "expected identifier",
		yyXError{382, -1}: "expected identifier",
		yyXError{500, -1}: "expected identifier",
		yyXError{501, -1}: "expected identifier",
		yyXError{508, -1}: "expected identifier",
		yyXError{336, -1}: "expected identifier list or identifier",
		yyXError{530, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{308, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{251, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{265, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{253, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{271, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{306, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{268, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{57, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{58, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{59, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{60, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{61, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{62, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{63, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{104, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{105, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
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
		yyXError{172, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{173, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{177, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{181, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{189, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{244, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{246, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{250, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{274, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{275, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{276, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{277, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{278, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{279, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{280, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{281, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{282, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{65, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{148, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{152, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{174, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{179, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{343, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{474, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{360, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{264, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{91, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{99, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{186, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{222, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{225, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{493, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{494, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{495, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{497, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{504, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{510, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{512, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{515, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{518, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{520, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{522, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{523, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{525, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{527, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{528, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{531, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{536, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{537, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{539, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{543, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{546, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{547, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{552, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{572, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{573, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{575, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{551, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{555, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{558, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{560, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{565, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{566, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{428, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{8, -1}:   "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{56, -1}:  "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{310, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{42, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{43, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{44, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{46, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{312, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{356, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{358, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{437, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{486, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{487, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{38, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{39, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{93, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{182, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{291, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{317, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{338, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{348, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{349, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{352, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{355, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{362, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{363, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{364, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{365, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{366, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{367, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{368, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{387, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{388, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{389, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{391, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{399, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{405, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{411, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{415, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{419, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{423, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{425, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{426, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{430, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{433, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{435, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{372, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{374, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{375, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{427, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{255, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{262, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{439, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{453, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
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
		yyXError{31, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{450, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{456, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{471, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{476, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{477, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{17, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{40, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{41, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{478, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{479, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{480, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{481, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{482, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{241, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{242, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{243, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{202, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{203, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{206, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{215, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{217, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{223, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{226, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{229, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{230, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{83, -1}:  "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{240, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{87, -1}:  "expected one of ['(', ')', ',', '[']",
		yyXError{136, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{183, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{187, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{188, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{190, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{198, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{234, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{238, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{286, -1}: "expected one of ['(', goto]",
		yyXError{318, -1}: "expected one of ['(', goto]",
		yyXError{293, -1}: "expected one of ['(', identifier]",
		yyXError{326, -1}: "expected one of [')', ',', ':']",
		yyXError{333, -1}: "expected one of [')', ',', ':']",
		yyXError{339, -1}: "expected one of [')', ',', ':']",
		yyXError{340, -1}: "expected one of [')', ',', ':']",
		yyXError{344, -1}: "expected one of [')', ',', ':']",
		yyXError{347, -1}: "expected one of [')', ',', ':']",
		yyXError{351, -1}: "expected one of [')', ',', ':']",
		yyXError{361, -1}: "expected one of [')', ',', ';']",
		yyXError{211, -1}: "expected one of [')', ',', ...]",
		yyXError{214, -1}: "expected one of [')', ',', ...]",
		yyXError{532, -1}: "expected one of [')', ',', ...]",
		yyXError{86, -1}:  "expected one of [')', ',']",
		yyXError{175, -1}: "expected one of [')', ',']",
		yyXError{193, -1}: "expected one of [')', ',']",
		yyXError{195, -1}: "expected one of [')', ',']",
		yyXError{200, -1}: "expected one of [')', ',']",
		yyXError{201, -1}: "expected one of [')', ',']",
		yyXError{212, -1}: "expected one of [')', ',']",
		yyXError{233, -1}: "expected one of [')', ',']",
		yyXError{247, -1}: "expected one of [')', ',']",
		yyXError{337, -1}: "expected one of [')', ',']",
		yyXError{354, -1}: "expected one of [')', ',']",
		yyXError{409, -1}: "expected one of [')', ',']",
		yyXError{413, -1}: "expected one of [')', ',']",
		yyXError{417, -1}: "expected one of [')', ',']",
		yyXError{421, -1}: "expected one of [')', ',']",
		yyXError{288, -1}: "expected one of [')', ':', string literal]",
		yyXError{290, -1}: "expected one of [')', ':', string literal]",
		yyXError{345, -1}: "expected one of [')', ':', string literal]",
		yyXError{289, -1}: "expected one of [')', string literal]",
		yyXError{464, -1}: "expected one of [',', ':', ';']",
		yyXError{149, -1}: "expected one of [',', ':']",
		yyXError{327, -1}: "expected one of [',', ':']",
		yyXError{334, -1}: "expected one of [',', ':']",
		yyXError{304, -1}: "expected one of [',', ';', '=']",
		yyXError{270, -1}: "expected one of [',', ';', '}']",
		yyXError{302, -1}: "expected one of [',', ';']",
		yyXError{303, -1}: "expected one of [',', ';']",
		yyXError{307, -1}: "expected one of [',', ';']",
		yyXError{309, -1}: "expected one of [',', ';']",
		yyXError{461, -1}: "expected one of [',', ';']",
		yyXError{463, -1}: "expected one of [',', ';']",
		yyXError{467, -1}: "expected one of [',', ';']",
		yyXError{470, -1}: "expected one of [',', ';']",
		yyXError{442, -1}: "expected one of [',', '=', '}']",
		yyXError{445, -1}: "expected one of [',', '=', '}']",
		yyXError{180, -1}: "expected one of [',', ']']",
		yyXError{266, -1}: "expected one of [',', '}']",
		yyXError{272, -1}: "expected one of [',', '}']",
		yyXError{444, -1}: "expected one of [',', '}']",
		yyXError{447, -1}: "expected one of [',', '}']",
		yyXError{451, -1}: "expected one of [',', '}']",
		yyXError{256, -1}: "expected one of ['.', '=', '[']",
		yyXError{259, -1}: "expected one of ['.', '=', '[']",
		yyXError{261, -1}: "expected one of ['.', '=', '[']",
		yyXError{263, -1}: "expected one of ['.', '=', '[']",
		yyXError{322, -1}: "expected one of [':', string literal]",
		yyXError{502, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		yyXError{511, -1}: "expected one of ['\\n', ppother]",
		yyXError{514, -1}: "expected one of ['\\n', ppother]",
		yyXError{516, -1}: "expected one of ['\\n', ppother]",
		yyXError{298, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{300, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{34, -1}:  "expected one of ['{', identifier]",
		yyXError{35, -1}:  "expected one of ['{', identifier]",
		yyXError{458, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{460, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{462, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{468, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{472, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{540, -1}: "expected one of [..., identifier]",
		yyXError{81, -1}:  "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{101, -1}: "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{357, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{359, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{252, -1}: "expected optional comma or one of [',', '}']",
		yyXError{267, -1}: "expected optional comma or one of [',', '}']",
		yyXError{443, -1}: "expected optional comma or one of [',', '}']",
		yyXError{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{12, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{396, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{402, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{385, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{394, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{400, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{218, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{207, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{90, -1}:  "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{94, -1}:  "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{549, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{554, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{557, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{564, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{570, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{208, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{33, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{36, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{297, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{192, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{235, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{236, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{79, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{80, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{503, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{507, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{82, -1}:  "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{316, -1}: "expected optional volatile or one of ['(', goto, volatile]",
		yyXError{48, -1}:  "expected optional volatile or one of ['(', volatile]",
		yyXError{231, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{209, -1}: "expected parameter type list or one of [_Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{239, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{490, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{529, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{535, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{538, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{542, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{545, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{78, -1}:  "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{380, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{398, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{404, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{414, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{418, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{422, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{424, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{429, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{432, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{434, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{53, -1}:  "expected string literal",
		yyXError{325, -1}: "expected string literal",
		yyXError{330, -1}: "expected string literal",
		yyXError{335, -1}: "expected string literal",
		yyXError{454, -1}: "expected struct declaration list or one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{455, -1}: "expected struct declaration list or one of [_Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{457, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{459, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		yyXError{469, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{519, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{496, -1}: "expected token list or ppother",
		yyXError{499, -1}: "expected token list or ppother",
		yyXError{505, -1}: "expected token list or ppother",
		yyXError{506, -1}: "expected token list or ppother",
		yyXError{509, -1}: "expected token list or ppother",
		yyXError{562, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of ['(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{5, -1}:   "expected translation unit or one of ['(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{76, -1}:  "expected type name or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{97, -1}:  "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{220, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{406, -1}: "expected while",
		yyXError{3, 50}:   "unexpected EOF",
		yyXError{2, 50}:   "unexpected EOF",
		yyXError{4, 50}:   "unexpected EOF",
	}

	yyParseTab = [576][]uint16{
		// 0
		{218: 324, 227: 323, 229: 322, 232: 325},
		{50: 321},
		{84: 320, 86: 320, 100: 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 204: 811},
		{318, 318, 318, 318, 318, 318, 318, 318, 11: 318, 13: 318, 318, 318, 318, 318, 318, 318, 318, 318, 209: 809},
		{316, 316, 316, 8: 316, 22: 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 43: 316, 316, 316, 316, 316, 316, 316, 55: 316, 57: 316, 210: 326},
		// 5
		{74, 74, 74, 8: 367, 22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 55: 369, 57: 370, 113: 332, 117: 351, 354, 350, 331, 122: 366, 328, 333, 126: 330, 138: 329, 144: 365, 175: 368, 191: 363, 193: 364, 233: 327},
		{74, 74, 74, 8: 367, 22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 315, 55: 369, 57: 370, 113: 332, 117: 351, 354, 350, 331, 122: 366, 328, 333, 126: 330, 138: 329, 144: 365, 175: 368, 191: 808, 193: 364},
		{161, 403, 161, 8: 234, 133: 614, 136: 613, 804, 170: 623, 196: 624, 622},
		{243, 243, 243, 243, 243, 243, 243, 243, 243, 11: 243, 13: 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 43: 243, 243, 243, 243, 243, 243, 243, 243, 52: 243, 55: 243, 57: 243, 85: 243, 87: 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{238, 238, 238, 8: 238, 238, 238, 12: 238, 22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 113: 332, 117: 351, 354, 350, 331, 123: 800, 333, 126: 330, 157: 803},
		// 10
		{238, 238, 238, 8: 238, 238, 238, 12: 238, 22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 113: 332, 117: 351, 354, 350, 331, 123: 800, 333, 126: 330, 157: 802},
		{238, 238, 238, 8: 238, 238, 238, 12: 238, 22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 113: 332, 117: 351, 354, 350, 331, 123: 800, 333, 126: 330, 157: 801},
		{238, 238, 238, 8: 238, 238, 238, 12: 238, 22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 113: 332, 117: 351, 354, 350, 331, 123: 800, 333, 126: 330, 157: 799},
		{229, 229, 229, 8: 229, 229, 229, 12: 229, 22: 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 43: 229, 229, 229, 229, 229, 229, 229},
		{228, 228, 228, 8: 228, 228, 228, 12: 228, 22: 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 43: 228, 228, 228, 228, 228, 228, 228},
		// 15
		{227, 227, 227, 8: 227, 227, 227, 12: 227, 22: 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 43: 227, 227, 227, 227, 227, 227, 227},
		{226, 226, 226, 8: 226, 226, 226, 12: 226, 22: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 43: 226, 226, 226, 226, 226, 226, 226},
		{225, 225, 225, 8: 225, 225, 225, 12: 225, 22: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 43: 225, 225, 225, 225, 225, 225, 225},
		{224, 224, 224, 8: 224, 224, 224, 12: 224, 22: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 42: 224, 224, 224, 224, 224, 224, 224, 224},
		{223, 223, 223, 8: 223, 223, 223, 12: 223, 22: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 42: 223, 223, 223, 223, 223, 223, 223, 223},
		// 20
		{222, 222, 222, 8: 222, 222, 222, 12: 222, 22: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 42: 222, 222, 222, 222, 222, 222, 222, 222},
		{221, 221, 221, 8: 221, 221, 221, 12: 221, 22: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 42: 221, 221, 221, 221, 221, 221, 221, 221},
		{220, 220, 220, 8: 220, 220, 220, 12: 220, 22: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 42: 220, 220, 220, 220, 220, 220, 220, 220},
		{219, 219, 219, 8: 219, 219, 219, 12: 219, 22: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 42: 219, 219, 219, 219, 219, 219, 219, 219},
		{218, 218, 218, 8: 218, 218, 218, 12: 218, 22: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 42: 218, 218, 218, 218, 218, 218, 218, 218},
		// 25
		{217, 217, 217, 8: 217, 217, 217, 12: 217, 22: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 42: 217, 217, 217, 217, 217, 217, 217, 217},
		{216, 216, 216, 8: 216, 216, 216, 12: 216, 22: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 42: 216, 216, 216, 216, 216, 216, 216, 216},
		{215, 215, 215, 8: 215, 215, 215, 12: 215, 22: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 42: 215, 215, 215, 215, 215, 215, 215, 215},
		{214, 214, 214, 8: 214, 214, 214, 12: 214, 22: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 42: 214, 214, 214, 214, 214, 214, 214, 214},
		{213, 213, 213, 8: 213, 213, 213, 12: 213, 22: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 42: 213, 213, 213, 213, 213, 213, 213, 213},
		// 30
		{212, 212, 212, 8: 212, 212, 212, 12: 212, 22: 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 42: 212, 212, 212, 212, 212, 212, 212, 212},
		{211, 211, 211, 8: 211, 211, 211, 12: 211, 22: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 42: 211, 211, 211, 211, 211, 211, 211, 211},
		{794},
		{2: 774, 52: 143, 195: 773},
		{2: 204, 52: 204},
		// 35
		{2: 203, 52: 203},
		{2: 760, 52: 143, 195: 759},
		{180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 42: 180, 180, 180, 180, 180, 180, 180, 180, 54: 180},
		{179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 42: 179, 179, 179, 179, 179, 179, 179, 179, 54: 179},
		{178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 42: 178, 178, 178, 178, 178, 178, 178, 178, 54: 178},
		// 40
		{177, 177, 177, 8: 177, 177, 177, 12: 177, 22: 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 43: 177, 177, 177, 177, 177, 177, 177},
		{176, 176, 176, 8: 176, 176, 176, 12: 176, 22: 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 43: 176, 176, 176, 176, 176, 176, 176},
		{82, 82, 82, 8: 82, 22: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 43: 82, 82, 82, 82, 82, 82, 82, 82, 55: 82, 57: 82},
		{80, 80, 80, 8: 80, 22: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 43: 80, 80, 80, 80, 80, 80, 80, 80, 55: 80, 57: 80},
		{79, 79, 79, 8: 79, 22: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 43: 79, 79, 79, 79, 79, 79, 79, 79, 55: 79, 57: 79},
		// 45
		{8: 758},
		{77, 77, 77, 8: 77, 22: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 43: 77, 77, 77, 77, 77, 77, 77, 77, 55: 77, 57: 77},
		{161, 403, 161, 133: 614, 136: 613, 615},
		{59, 22: 607, 203: 606},
		{371},
		// 50
		{246, 246, 246, 246, 246, 246, 246, 246, 11: 246, 13: 246, 246, 246, 246, 246, 246, 246, 246, 246, 145: 372, 373},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 386},
		{10: 374},
		{11: 375},
		{9: 376},
		// 55
		{8: 377},
		{44, 44, 44, 44, 44, 44, 44, 44, 44, 11: 44, 13: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 43: 44, 44, 44, 44, 44, 44, 44, 44, 52: 44, 55: 44, 57: 44, 85: 44, 87: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44},
		{309, 309, 3: 309, 309, 309, 309, 309, 309, 309, 309, 12: 309, 41: 309, 309, 50: 309, 309, 54: 309, 56: 309, 58: 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309},
		{308, 308, 3: 308, 308, 308, 308, 308, 308, 308, 308, 12: 308, 41: 308, 308, 50: 308, 308, 54: 308, 56: 308, 58: 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308},
		{307, 307, 3: 307, 307, 307, 307, 307, 307, 307, 307, 12: 307, 41: 307, 307, 50: 307, 307, 54: 307, 56: 307, 58: 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307},
		// 60
		{306, 306, 3: 306, 306, 306, 306, 306, 306, 306, 306, 12: 306, 41: 306, 306, 50: 306, 306, 54: 306, 56: 306, 58: 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306},
		{305, 305, 3: 305, 305, 305, 305, 305, 305, 305, 305, 12: 305, 41: 305, 305, 50: 305, 305, 54: 305, 56: 305, 58: 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305},
		{304, 304, 3: 304, 304, 304, 304, 304, 304, 304, 304, 12: 304, 41: 304, 304, 50: 304, 304, 54: 304, 56: 304, 58: 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304},
		{303, 303, 3: 303, 303, 303, 303, 303, 303, 303, 303, 12: 303, 41: 303, 303, 50: 303, 303, 54: 303, 56: 303, 58: 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 53: 469, 114: 568, 153: 399, 160: 604},
		// 65
		{422, 427, 3: 440, 430, 431, 426, 425, 245, 10: 245, 12: 421, 41: 245, 245, 50: 245, 446, 54: 245, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 603},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 602},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 601},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 510},
		// 70
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 600},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 599},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 598},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 597},
		{566, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 567},
		// 75
		{397},
		{22: 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 153: 399, 160: 398},
		{9: 565},
		{22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 113: 401, 117: 351, 354, 350, 400, 151: 402},
		{195, 195, 195, 8: 195, 195, 12: 195, 22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 42: 195, 113: 401, 117: 351, 354, 350, 400, 151: 563, 200: 564},
		// 80
		{195, 195, 195, 8: 195, 195, 12: 195, 22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 42: 195, 113: 401, 117: 351, 354, 350, 400, 151: 563, 200: 562},
		{161, 403, 9: 137, 12: 161, 133: 404, 136: 406, 154: 407, 180: 405},
		{157, 157, 157, 9: 157, 157, 12: 157, 22: 360, 358, 359, 113: 414, 152: 418, 161: 560},
		{160, 2: 160, 9: 139, 139, 12: 160},
		{9: 140},
		// 85
		{409, 12: 125, 185: 408, 410},
		{9: 136, 136},
		{556, 9: 138, 138, 12: 124},
		{161, 403, 9: 129, 12: 161, 22: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 43: 129, 129, 129, 129, 129, 129, 129, 133: 404, 136: 406, 154: 512, 174: 513},
		{12: 411},
		// 90
		{385, 413, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 360, 358, 359, 43: 417, 53: 412, 252, 113: 414, 152: 415, 168: 416},
		{422, 427, 3: 440, 430, 431, 426, 425, 12: 421, 51: 446, 54: 251, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 510, 511},
		{159, 159, 159, 159, 159, 159, 159, 159, 9: 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 43: 159, 54: 159},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 360, 358, 359, 43: 506, 53: 412, 252, 113: 503, 168: 505},
		// 95
		{54: 504},
		{157, 157, 157, 157, 157, 157, 157, 157, 11: 157, 13: 157, 157, 157, 157, 157, 157, 157, 157, 157, 360, 358, 359, 113: 414, 152: 418, 161: 419},
		{156, 156, 156, 156, 156, 156, 156, 156, 9: 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 360, 358, 359, 113: 503},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 420},
		{422, 427, 3: 440, 430, 431, 426, 425, 12: 421, 51: 446, 54: 457, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		// 100
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 469, 114: 501},
		{385, 390, 378, 389, 391, 392, 388, 387, 9: 311, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 495, 214: 496, 497},
		{2: 494},
		{2: 493},
		{297, 297, 3: 297, 297, 297, 297, 297, 297, 297, 297, 12: 297, 41: 297, 297, 50: 297, 297, 54: 297, 56: 297, 58: 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297},
		// 105
		{296, 296, 3: 296, 296, 296, 296, 296, 296, 296, 296, 12: 296, 41: 296, 296, 50: 296, 296, 54: 296, 56: 296, 58: 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 492},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 491},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 490},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 489},
		// 110
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 488},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 487},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 486},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 485},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 484},
		// 115
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 483},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 482},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 481},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 480},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 479},
		// 120
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 478},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 477},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 476},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 475},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 469, 114: 470},
		// 125
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 468},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 467},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 466},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 465},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 464},
		// 130
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 463},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 462},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 461},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 460},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 459},
		// 135
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 458},
		{132, 9: 132, 132, 12: 132},
		{422, 427, 3: 440, 430, 431, 426, 425, 254, 254, 254, 12: 421, 41: 254, 254, 50: 254, 446, 54: 254, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{422, 427, 3: 440, 430, 431, 426, 425, 255, 255, 255, 12: 421, 41: 255, 255, 50: 255, 446, 54: 255, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{422, 427, 3: 440, 430, 431, 426, 425, 256, 256, 256, 12: 421, 41: 256, 256, 50: 256, 446, 54: 256, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		// 140
		{422, 427, 3: 440, 430, 431, 426, 425, 257, 257, 257, 12: 421, 41: 257, 257, 50: 257, 446, 54: 257, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{422, 427, 3: 440, 430, 431, 426, 425, 258, 258, 258, 12: 421, 41: 258, 258, 50: 258, 446, 54: 258, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{422, 427, 3: 440, 430, 431, 426, 425, 259, 259, 259, 12: 421, 41: 259, 259, 50: 259, 446, 54: 259, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{422, 427, 3: 440, 430, 431, 426, 425, 260, 260, 260, 12: 421, 41: 260, 260, 50: 260, 446, 54: 260, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{422, 427, 3: 440, 430, 431, 426, 425, 261, 261, 261, 12: 421, 41: 261, 261, 50: 261, 446, 54: 261, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		// 145
		{422, 427, 3: 440, 430, 431, 426, 425, 262, 262, 262, 12: 421, 41: 262, 262, 50: 262, 446, 54: 262, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{422, 427, 3: 440, 430, 431, 426, 425, 263, 263, 263, 12: 421, 41: 263, 263, 50: 263, 446, 54: 263, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{422, 427, 3: 440, 430, 431, 426, 425, 264, 264, 264, 12: 421, 41: 264, 264, 50: 264, 446, 54: 264, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{422, 427, 3: 440, 430, 431, 426, 425, 250, 250, 250, 12: 421, 42: 250, 51: 446, 54: 250, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{10: 472, 42: 471},
		// 150
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 474},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 473},
		{422, 427, 3: 440, 430, 431, 426, 425, 249, 249, 249, 12: 421, 42: 249, 51: 446, 54: 249, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{422, 427, 3: 440, 430, 431, 426, 425, 265, 265, 265, 12: 421, 41: 265, 265, 50: 265, 265, 54: 265, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 265, 443, 265, 424, 265, 438, 437, 436, 432, 265, 265, 265, 439, 265, 444, 433, 265, 265, 265},
		{422, 427, 3: 440, 430, 431, 426, 425, 266, 266, 266, 12: 421, 41: 266, 266, 50: 266, 266, 54: 266, 56: 423, 58: 429, 428, 434, 435, 266, 441, 442, 266, 443, 266, 424, 266, 438, 437, 436, 432, 266, 266, 266, 439, 266, 266, 433, 266, 266, 266},
		// 155
		{422, 427, 3: 440, 430, 431, 426, 425, 267, 267, 267, 12: 421, 41: 267, 267, 50: 267, 267, 54: 267, 56: 423, 58: 429, 428, 434, 435, 267, 441, 442, 267, 267, 267, 424, 267, 438, 437, 436, 432, 267, 267, 267, 439, 267, 267, 433, 267, 267, 267},
		{422, 427, 3: 440, 430, 431, 426, 425, 268, 268, 268, 12: 421, 41: 268, 268, 50: 268, 268, 54: 268, 56: 423, 58: 429, 428, 434, 435, 268, 441, 268, 268, 268, 268, 424, 268, 438, 437, 436, 432, 268, 268, 268, 439, 268, 268, 433, 268, 268, 268},
		{422, 427, 3: 440, 430, 431, 426, 425, 269, 269, 269, 12: 421, 41: 269, 269, 50: 269, 269, 54: 269, 56: 423, 58: 429, 428, 434, 435, 269, 269, 269, 269, 269, 269, 424, 269, 438, 437, 436, 432, 269, 269, 269, 439, 269, 269, 433, 269, 269, 269},
		{422, 427, 3: 270, 430, 431, 426, 425, 270, 270, 270, 12: 421, 41: 270, 270, 50: 270, 270, 54: 270, 56: 423, 58: 429, 428, 434, 435, 270, 270, 270, 270, 270, 270, 424, 270, 438, 437, 436, 432, 270, 270, 270, 439, 270, 270, 433, 270, 270, 270},
		{422, 427, 3: 271, 430, 431, 426, 425, 271, 271, 271, 12: 421, 41: 271, 271, 50: 271, 271, 54: 271, 56: 423, 58: 429, 428, 434, 435, 271, 271, 271, 271, 271, 271, 424, 271, 271, 437, 436, 432, 271, 271, 271, 271, 271, 271, 433, 271, 271, 271},
		// 160
		{422, 427, 3: 272, 430, 431, 426, 425, 272, 272, 272, 12: 421, 41: 272, 272, 50: 272, 272, 54: 272, 56: 423, 58: 429, 428, 434, 435, 272, 272, 272, 272, 272, 272, 424, 272, 272, 437, 436, 432, 272, 272, 272, 272, 272, 272, 433, 272, 272, 272},
		{422, 427, 3: 273, 430, 431, 426, 425, 273, 273, 273, 12: 421, 41: 273, 273, 50: 273, 273, 54: 273, 56: 423, 58: 429, 428, 273, 273, 273, 273, 273, 273, 273, 273, 424, 273, 273, 273, 273, 432, 273, 273, 273, 273, 273, 273, 433, 273, 273, 273},
		{422, 427, 3: 274, 430, 431, 426, 425, 274, 274, 274, 12: 421, 41: 274, 274, 50: 274, 274, 54: 274, 56: 423, 58: 429, 428, 274, 274, 274, 274, 274, 274, 274, 274, 424, 274, 274, 274, 274, 432, 274, 274, 274, 274, 274, 274, 433, 274, 274, 274},
		{422, 427, 3: 275, 430, 431, 426, 425, 275, 275, 275, 12: 421, 41: 275, 275, 50: 275, 275, 54: 275, 56: 423, 58: 429, 428, 275, 275, 275, 275, 275, 275, 275, 275, 424, 275, 275, 275, 275, 432, 275, 275, 275, 275, 275, 275, 433, 275, 275, 275},
		{422, 427, 3: 276, 430, 431, 426, 425, 276, 276, 276, 12: 421, 41: 276, 276, 50: 276, 276, 54: 276, 56: 423, 58: 429, 428, 276, 276, 276, 276, 276, 276, 276, 276, 424, 276, 276, 276, 276, 432, 276, 276, 276, 276, 276, 276, 433, 276, 276, 276},
		// 165
		{422, 427, 3: 277, 430, 431, 426, 425, 277, 277, 277, 12: 421, 41: 277, 277, 50: 277, 277, 54: 277, 56: 423, 58: 429, 428, 277, 277, 277, 277, 277, 277, 277, 277, 424, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277},
		{422, 427, 3: 278, 430, 431, 426, 425, 278, 278, 278, 12: 421, 41: 278, 278, 50: 278, 278, 54: 278, 56: 423, 58: 429, 428, 278, 278, 278, 278, 278, 278, 278, 278, 424, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278},
		{422, 427, 3: 279, 279, 279, 426, 425, 279, 279, 279, 12: 421, 41: 279, 279, 50: 279, 279, 54: 279, 56: 423, 58: 429, 428, 279, 279, 279, 279, 279, 279, 279, 279, 424, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279},
		{422, 427, 3: 280, 280, 280, 426, 425, 280, 280, 280, 12: 421, 41: 280, 280, 50: 280, 280, 54: 280, 56: 423, 58: 429, 428, 280, 280, 280, 280, 280, 280, 280, 280, 424, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280},
		{422, 281, 3: 281, 281, 281, 426, 425, 281, 281, 281, 12: 421, 41: 281, 281, 50: 281, 281, 54: 281, 56: 423, 58: 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 424, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281},
		// 170
		{422, 282, 3: 282, 282, 282, 426, 425, 282, 282, 282, 12: 421, 41: 282, 282, 50: 282, 282, 54: 282, 56: 423, 58: 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 424, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282},
		{422, 283, 3: 283, 283, 283, 426, 425, 283, 283, 283, 12: 421, 41: 283, 283, 50: 283, 283, 54: 283, 56: 423, 58: 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 424, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283},
		{298, 298, 3: 298, 298, 298, 298, 298, 298, 298, 298, 12: 298, 41: 298, 298, 50: 298, 298, 54: 298, 56: 298, 58: 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298},
		{299, 299, 3: 299, 299, 299, 299, 299, 299, 299, 299, 12: 299, 41: 299, 299, 50: 299, 299, 54: 299, 56: 299, 58: 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299},
		{422, 427, 3: 440, 430, 431, 426, 425, 9: 313, 313, 12: 421, 51: 446, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		// 175
		{9: 310, 499},
		{9: 498},
		{300, 300, 3: 300, 300, 300, 300, 300, 300, 300, 300, 12: 300, 41: 300, 300, 50: 300, 300, 54: 300, 56: 300, 58: 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 500},
		{422, 427, 3: 440, 430, 431, 426, 425, 9: 312, 312, 12: 421, 51: 446, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		// 180
		{10: 472, 54: 502},
		{301, 301, 3: 301, 301, 301, 301, 301, 301, 301, 301, 12: 301, 41: 301, 301, 50: 301, 301, 54: 301, 56: 301, 58: 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301},
		{158, 158, 158, 158, 158, 158, 158, 158, 9: 158, 158, 158, 158, 158, 158, 158, 158, 158, 158, 158, 158, 158, 158, 158, 158, 43: 158, 54: 158},
		{134, 9: 134, 134, 12: 134},
		{54: 509},
		// 185
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 507},
		{422, 427, 3: 440, 430, 431, 426, 425, 12: 421, 51: 446, 54: 508, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{131, 9: 131, 131, 12: 131},
		{133, 9: 133, 133, 12: 133},
		{422, 291, 3: 291, 291, 291, 426, 425, 291, 291, 291, 12: 421, 41: 291, 291, 50: 291, 291, 54: 291, 56: 423, 58: 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 424, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291},
		// 190
		{130, 9: 130, 130, 12: 130},
		{9: 555},
		{9: 153, 22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 113: 332, 117: 351, 354, 350, 331, 123: 517, 333, 126: 330, 159: 516, 172: 514, 515, 199: 518},
		{9: 155, 552},
		{9: 152},
		// 195
		{9: 151, 151},
		{161, 403, 161, 9: 137, 137, 12: 161, 133: 404, 136: 520, 521, 154: 407, 180: 522},
		{9: 519},
		{128, 9: 128, 128, 12: 128},
		{525, 2: 524, 12: 125, 185: 408, 410, 523},
		// 200
		{9: 149, 149},
		{9: 148, 148},
		{529, 8: 175, 175, 175, 12: 528, 22: 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 42: 175, 175, 175, 175, 175, 175, 175, 175, 51: 175, 175, 55: 175, 57: 175},
		{172, 8: 172, 172, 172, 12: 172, 22: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 42: 172, 172, 172, 172, 172, 172, 172, 172, 51: 172, 172, 55: 172, 57: 172},
		{161, 403, 161, 9: 129, 12: 161, 22: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 43: 129, 129, 129, 129, 129, 129, 129, 133: 404, 136: 520, 526, 154: 512, 174: 513},
		// 205
		{9: 527},
		{171, 8: 171, 171, 171, 12: 171, 22: 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 42: 171, 171, 171, 171, 171, 171, 171, 171, 51: 171, 171, 55: 171, 57: 171},
		{157, 157, 157, 157, 157, 157, 157, 157, 11: 157, 13: 157, 157, 157, 157, 157, 157, 157, 157, 157, 360, 358, 359, 43: 540, 54: 157, 113: 414, 152: 541, 161: 539},
		{2: 532, 9: 145, 22: 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 43: 166, 166, 166, 166, 166, 166, 166, 169: 533, 194: 531, 213: 530},
		{22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 113: 332, 117: 351, 354, 350, 331, 123: 517, 333, 126: 330, 159: 516, 172: 514, 537},
		// 210
		{9: 536},
		{9: 147, 147, 147: 147},
		{9: 144, 534},
		{2: 535},
		{9: 146, 146, 147: 146},
		// 215
		{164, 8: 164, 164, 164, 12: 164, 22: 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 42: 164, 164, 164, 164, 164, 164, 164, 164, 51: 164, 164, 55: 164, 57: 164},
		{9: 538},
		{165, 8: 165, 165, 165, 12: 165, 22: 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 42: 165, 165, 165, 165, 165, 165, 165, 165, 51: 165, 165, 55: 165, 57: 165},
		{385, 548, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 412, 252, 168: 549},
		{157, 157, 157, 157, 157, 157, 157, 157, 11: 157, 13: 157, 157, 157, 157, 157, 157, 157, 157, 157, 360, 358, 359, 113: 414, 152: 418, 161: 545},
		// 220
		{156, 156, 156, 156, 156, 156, 156, 156, 11: 156, 13: 156, 156, 156, 156, 156, 156, 156, 156, 156, 360, 358, 359, 43: 542, 54: 156, 113: 503},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 543},
		{422, 427, 3: 440, 430, 431, 426, 425, 12: 421, 51: 446, 54: 544, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{168, 8: 168, 168, 168, 12: 168, 22: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 42: 168, 168, 168, 168, 168, 168, 168, 168, 51: 168, 168, 55: 168, 57: 168},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 546},
		// 225
		{422, 427, 3: 440, 430, 431, 426, 425, 12: 421, 51: 446, 54: 547, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{169, 8: 169, 169, 169, 12: 169, 22: 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 42: 169, 169, 169, 169, 169, 169, 169, 169, 51: 169, 169, 55: 169, 57: 169},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 510, 551},
		{54: 550},
		{170, 8: 170, 170, 170, 12: 170, 22: 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 42: 170, 170, 170, 170, 170, 170, 170, 170, 51: 170, 170, 55: 170, 57: 170},
		// 230
		{167, 8: 167, 167, 167, 12: 167, 22: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 42: 167, 167, 167, 167, 167, 167, 167, 167, 51: 167, 167, 55: 167, 57: 167},
		{22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 113: 332, 117: 351, 354, 350, 331, 123: 517, 333, 126: 330, 147: 553, 159: 554},
		{9: 154},
		{9: 150, 150},
		{135, 9: 135, 135, 12: 135},
		// 235
		{9: 127, 22: 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 43: 127, 127, 127, 127, 127, 127, 127, 205: 557},
		{9: 153, 22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 113: 332, 117: 351, 354, 350, 331, 123: 517, 333, 126: 330, 159: 516, 172: 514, 515, 199: 558},
		{9: 559},
		{126, 9: 126, 126, 12: 126},
		{163, 403, 163, 9: 163, 163, 12: 163, 133: 561},
		// 240
		{162, 2: 162, 9: 162, 162, 12: 162},
		{196, 196, 196, 8: 196, 196, 12: 196, 42: 196},
		{194, 194, 194, 8: 194, 194, 12: 194, 42: 194},
		{197, 197, 197, 8: 197, 197, 12: 197, 42: 197},
		{253, 253, 3: 253, 253, 253, 253, 253, 253, 253, 253, 12: 253, 41: 253, 253, 50: 253, 253, 54: 253, 56: 253, 58: 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253},
		// 245
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 53: 469, 114: 568, 153: 399, 160: 569},
		{422, 286, 3: 286, 286, 286, 426, 425, 286, 286, 286, 12: 421, 41: 286, 286, 50: 286, 286, 54: 286, 56: 423, 58: 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 424, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286},
		{9: 596, 472},
		{9: 570},
		{385, 285, 378, 285, 285, 285, 388, 387, 285, 285, 285, 384, 285, 394, 393, 396, 379, 380, 381, 382, 383, 395, 41: 285, 285, 50: 285, 285, 572, 571, 285, 56: 285, 58: 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285},
		// 250
		{422, 284, 3: 284, 284, 284, 426, 425, 284, 284, 284, 12: 421, 41: 284, 284, 50: 284, 284, 54: 284, 56: 423, 58: 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 424, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284},
		{118, 118, 118, 118, 118, 118, 118, 118, 11: 118, 578, 118, 118, 118, 118, 118, 118, 118, 118, 118, 52: 118, 56: 579, 158: 577, 165: 576, 574, 575, 198: 573},
		{10: 589, 41: 189, 164: 594},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 586, 585, 171: 587},
		{12: 578, 51: 583, 56: 579, 158: 584},
		// 255
		{117, 117, 117, 117, 117, 117, 117, 117, 11: 117, 13: 117, 117, 117, 117, 117, 117, 117, 117, 117, 52: 117},
		{12: 116, 51: 116, 56: 116},
		{246, 246, 246, 246, 246, 246, 246, 246, 11: 246, 13: 246, 246, 246, 246, 246, 246, 246, 246, 246, 145: 372, 581},
		{2: 580},
		{12: 113, 51: 113, 56: 113},
		// 260
		{54: 582},
		{12: 114, 51: 114, 56: 114},
		{119, 119, 119, 119, 119, 119, 119, 119, 11: 119, 13: 119, 119, 119, 119, 119, 119, 119, 119, 119, 52: 119},
		{12: 115, 51: 115, 56: 115},
		{422, 427, 3: 440, 430, 431, 426, 425, 123, 10: 123, 12: 421, 41: 123, 51: 446, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		// 265
		{118, 118, 118, 118, 118, 118, 118, 118, 11: 118, 578, 118, 118, 118, 118, 118, 118, 118, 118, 118, 52: 118, 56: 579, 158: 577, 165: 576, 574, 575, 198: 588},
		{10: 121, 41: 121},
		{10: 589, 41: 189, 164: 590},
		{118, 118, 118, 118, 118, 118, 118, 118, 11: 118, 578, 118, 118, 118, 118, 118, 118, 118, 118, 118, 41: 188, 52: 118, 56: 579, 158: 577, 165: 576, 592, 575},
		{41: 591},
		// 270
		{8: 122, 10: 122, 41: 122},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 586, 585, 171: 593},
		{10: 120, 41: 120},
		{41: 595},
		{295, 295, 3: 295, 295, 295, 295, 295, 295, 295, 295, 12: 295, 41: 295, 295, 50: 295, 295, 54: 295, 56: 295, 58: 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295},
		// 275
		{302, 302, 3: 302, 302, 302, 302, 302, 302, 302, 302, 12: 302, 41: 302, 302, 50: 302, 302, 54: 302, 56: 302, 58: 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302},
		{422, 287, 3: 287, 287, 287, 426, 425, 287, 287, 287, 12: 421, 41: 287, 287, 50: 287, 287, 54: 287, 56: 423, 58: 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 424, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287},
		{422, 288, 3: 288, 288, 288, 426, 425, 288, 288, 288, 12: 421, 41: 288, 288, 50: 288, 288, 54: 288, 56: 423, 58: 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 424, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288},
		{422, 289, 3: 289, 289, 289, 426, 425, 289, 289, 289, 12: 421, 41: 289, 289, 50: 289, 289, 54: 289, 56: 423, 58: 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 424, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289},
		{422, 290, 3: 290, 290, 290, 426, 425, 290, 290, 290, 12: 421, 41: 290, 290, 50: 290, 290, 54: 290, 56: 423, 58: 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 424, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290},
		// 280
		{422, 292, 3: 292, 292, 292, 426, 425, 292, 292, 292, 12: 421, 41: 292, 292, 50: 292, 292, 54: 292, 56: 423, 58: 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 424, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292},
		{422, 293, 3: 293, 293, 293, 426, 425, 293, 293, 293, 12: 421, 41: 293, 293, 50: 293, 293, 54: 293, 56: 423, 58: 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 424, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293},
		{422, 294, 3: 294, 294, 294, 426, 425, 294, 294, 294, 12: 421, 41: 294, 294, 50: 294, 294, 54: 294, 56: 423, 58: 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 424, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294},
		{9: 605},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 572, 571},
		// 285
		{608},
		{58, 85: 58},
		{11: 609, 162: 610},
		{9: 62, 11: 62, 42: 62},
		{9: 612, 11: 611},
		// 290
		{9: 61, 11: 61, 42: 61},
		{60, 60, 60, 60, 60, 60, 60, 60, 60, 11: 60, 13: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 43: 60, 60, 60, 60, 60, 60, 60, 52: 60, 55: 60, 57: 60, 85: 60, 87: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 112: 60},
		{757, 2: 524, 187: 523},
		{160, 2: 160},
		{22: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 43: 64, 64, 64, 64, 64, 64, 64, 52: 65, 55: 65, 57: 64, 178: 617, 183: 616},
		// 295
		{52: 73, 55: 73, 208: 632},
		{22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 57: 370, 113: 332, 117: 351, 354, 350, 331, 123: 618, 333, 126: 330, 138: 329, 144: 619, 219: 620},
		{161, 403, 161, 8: 234, 133: 614, 136: 613, 625, 170: 623, 196: 624, 622},
		{22: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 43: 67, 67, 67, 67, 67, 67, 67, 52: 67, 55: 67, 57: 67},
		{22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 52: 63, 55: 63, 57: 370, 113: 332, 117: 351, 354, 350, 331, 123: 618, 333, 126: 330, 138: 329, 144: 621},
		// 300
		{22: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 43: 66, 66, 66, 66, 66, 66, 66, 52: 66, 55: 66, 57: 66},
		{8: 631},
		{8: 236, 10: 236},
		{8: 233, 10: 629},
		{8: 232, 10: 232, 51: 231, 179: 626},
		// 305
		{51: 627},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 586, 585, 171: 628},
		{8: 230, 10: 230},
		{161, 403, 161, 133: 614, 136: 613, 625, 170: 630},
		{8: 235, 10: 235},
		// 310
		{244, 244, 244, 244, 244, 244, 244, 244, 244, 11: 244, 13: 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 43: 244, 244, 244, 244, 244, 244, 244, 244, 52: 244, 55: 244, 57: 244, 85: 244, 87: 244, 244, 244, 244, 244, 244, 244, 244, 244, 244},
		{52: 71, 55: 69, 176: 634, 635, 192: 633},
		{72, 72, 72, 8: 72, 22: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 43: 72, 72, 72, 72, 72, 72, 72, 72, 55: 72, 57: 72},
		{52: 678, 128: 679},
		{55: 637, 122: 638, 127: 636},
		// 315
		{8: 677},
		{59, 22: 607, 85: 59, 203: 639},
		{50, 50, 50, 50, 50, 50, 50, 50, 50, 11: 50, 13: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 43: 50, 50, 50, 50, 50, 50, 50, 52: 50, 55: 50, 57: 50, 85: 50, 87: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 112: 50},
		{640, 85: 641},
		{11: 609, 162: 666},
		// 320
		{642},
		{11: 609, 162: 643},
		{11: 611, 42: 644},
		{42: 645},
		{11: 54, 649, 155: 647, 646, 163: 648},
		// 325
		{11: 662},
		{9: 56, 56, 42: 56},
		{10: 652, 42: 653},
		{2: 650},
		{54: 651},
		// 330
		{11: 53},
		{11: 54, 649, 155: 661, 646},
		{11: 654, 182: 655},
		{9: 52, 52, 42: 52},
		{10: 656, 42: 657},
		// 335
		{11: 660},
		{2: 532, 169: 658},
		{9: 659, 534},
		{46, 46, 46, 46, 46, 46, 46, 46, 46, 11: 46, 13: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 43: 46, 46, 46, 46, 46, 46, 46, 52: 46, 55: 46, 57: 46, 85: 46, 87: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 112: 46},
		{9: 51, 51, 42: 51},
		// 340
		{9: 55, 55, 42: 55},
		{663},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 664},
		{422, 427, 3: 440, 430, 431, 426, 425, 9: 665, 12: 421, 51: 446, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		{9: 57, 57, 42: 57},
		// 345
		{9: 612, 11: 611, 42: 667},
		{9: 669, 11: 54, 649, 155: 647, 646, 163: 668},
		{9: 670, 652, 42: 671},
		{45, 45, 45, 45, 45, 45, 45, 45, 45, 11: 45, 13: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 43: 45, 45, 45, 45, 45, 45, 45, 52: 45, 55: 45, 57: 45, 85: 45, 87: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 112: 45},
		{49, 49, 49, 49, 49, 49, 49, 49, 49, 11: 49, 13: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 43: 49, 49, 49, 49, 49, 49, 49, 52: 49, 55: 49, 57: 49, 85: 49, 87: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 112: 49},
		// 350
		{11: 54, 649, 155: 647, 646, 163: 672},
		{9: 673, 652, 42: 674},
		{48, 48, 48, 48, 48, 48, 48, 48, 48, 11: 48, 13: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 43: 48, 48, 48, 48, 48, 48, 48, 52: 48, 55: 48, 57: 48, 85: 48, 87: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 112: 48},
		{11: 654, 182: 675},
		{9: 676, 656},
		// 355
		{47, 47, 47, 47, 47, 47, 47, 47, 47, 11: 47, 13: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 43: 47, 47, 47, 47, 47, 47, 47, 52: 47, 55: 47, 57: 47, 85: 47, 87: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 112: 47},
		{68, 68, 68, 8: 68, 22: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 43: 68, 68, 68, 68, 68, 68, 68, 68, 55: 68, 57: 68},
		{102, 102, 102, 102, 102, 102, 102, 102, 102, 11: 102, 13: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 43: 102, 102, 102, 102, 102, 102, 102, 52: 102, 55: 102, 57: 102, 85: 102, 87: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 206: 680},
		{70, 70, 70, 8: 70, 22: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 43: 70, 70, 70, 70, 70, 70, 70, 70, 55: 70, 57: 70},
		{385, 390, 681, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 98, 43: 336, 337, 335, 361, 362, 338, 334, 52: 678, 469, 55: 637, 57: 370, 85: 703, 87: 700, 705, 690, 704, 691, 701, 702, 698, 706, 699, 113: 332, 682, 117: 351, 354, 350, 331, 697, 638, 618, 333, 126: 330, 689, 684, 685, 687, 688, 683, 134: 686, 696, 138: 329, 144: 695, 181: 693, 216: 694, 692},
		// 360
		{309, 309, 3: 309, 309, 309, 309, 309, 309, 10: 309, 12: 309, 42: 755, 51: 309, 56: 309, 58: 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309},
		{8: 247, 247, 472},
		{112, 112, 112, 112, 112, 112, 112, 112, 112, 11: 112, 13: 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 43: 112, 112, 112, 112, 112, 112, 112, 52: 112, 55: 112, 57: 112, 85: 112, 87: 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112: 112},
		{111, 111, 111, 111, 111, 111, 111, 111, 111, 11: 111, 13: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 43: 111, 111, 111, 111, 111, 111, 111, 52: 111, 55: 111, 57: 111, 85: 111, 87: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 112: 111},
		{110, 110, 110, 110, 110, 110, 110, 110, 110, 11: 110, 13: 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 43: 110, 110, 110, 110, 110, 110, 110, 52: 110, 55: 110, 57: 110, 85: 110, 87: 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 112: 110},
		// 365
		{109, 109, 109, 109, 109, 109, 109, 109, 109, 11: 109, 13: 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 43: 109, 109, 109, 109, 109, 109, 109, 52: 109, 55: 109, 57: 109, 85: 109, 87: 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 112: 109},
		{108, 108, 108, 108, 108, 108, 108, 108, 108, 11: 108, 13: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 43: 108, 108, 108, 108, 108, 108, 108, 52: 108, 55: 108, 57: 108, 85: 108, 87: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 112: 108},
		{107, 107, 107, 107, 107, 107, 107, 107, 107, 11: 107, 13: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 43: 107, 107, 107, 107, 107, 107, 107, 52: 107, 55: 107, 57: 107, 85: 107, 87: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 112: 107},
		{106, 106, 106, 106, 106, 106, 106, 106, 106, 11: 106, 13: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 43: 106, 106, 106, 106, 106, 106, 106, 52: 106, 55: 106, 57: 106, 85: 106, 87: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 112: 106},
		{246, 246, 246, 246, 246, 246, 246, 246, 11: 246, 13: 246, 246, 246, 246, 246, 246, 246, 246, 246, 145: 372, 752},
		// 370
		{42: 750},
		{41: 749},
		{100, 100, 100, 100, 100, 100, 100, 100, 100, 11: 100, 13: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 43: 100, 100, 100, 100, 100, 100, 100, 52: 100, 55: 100, 57: 100, 85: 100, 87: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
		{385, 390, 681, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 97, 43: 336, 337, 335, 361, 362, 338, 334, 52: 678, 469, 55: 637, 57: 370, 85: 703, 87: 700, 705, 690, 704, 691, 701, 702, 698, 706, 699, 113: 332, 682, 117: 351, 354, 350, 331, 697, 638, 618, 333, 126: 330, 689, 684, 685, 687, 688, 683, 134: 686, 696, 138: 329, 144: 695, 181: 748},
		{96, 96, 96, 96, 96, 96, 96, 96, 96, 11: 96, 13: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 43: 96, 96, 96, 96, 96, 96, 96, 52: 96, 55: 96, 57: 96, 85: 96, 87: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96},
		// 375
		{95, 95, 95, 95, 95, 95, 95, 95, 95, 11: 95, 13: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 43: 95, 95, 95, 95, 95, 95, 95, 52: 95, 55: 95, 57: 95, 85: 95, 87: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95},
		{8: 747},
		{741},
		{737},
		{733},
		// 380
		{385, 390, 681, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 678, 469, 55: 637, 85: 703, 87: 700, 705, 690, 704, 691, 701, 702, 698, 706, 699, 114: 682, 121: 697, 638, 127: 689, 684, 685, 687, 688, 683, 134: 686, 727},
		{713},
		{2: 711},
		{8: 710},
		{8: 709},
		// 385
		{385, 390, 378, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 469, 114: 682, 121: 707},
		{8: 708},
		{83, 83, 83, 83, 83, 83, 83, 83, 83, 11: 83, 13: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 43: 83, 83, 83, 83, 83, 83, 83, 52: 83, 55: 83, 57: 83, 85: 83, 87: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 112: 83},
		{84, 84, 84, 84, 84, 84, 84, 84, 84, 11: 84, 13: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 43: 84, 84, 84, 84, 84, 84, 84, 52: 84, 55: 84, 57: 84, 85: 84, 87: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 112: 84},
		{85, 85, 85, 85, 85, 85, 85, 85, 85, 11: 85, 13: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 43: 85, 85, 85, 85, 85, 85, 85, 52: 85, 55: 85, 57: 85, 85: 85, 87: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 112: 85},
		// 390
		{8: 712},
		{86, 86, 86, 86, 86, 86, 86, 86, 86, 11: 86, 13: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 43: 86, 86, 86, 86, 86, 86, 86, 52: 86, 55: 86, 57: 86, 85: 86, 87: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 112: 86},
		{385, 390, 378, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 43: 336, 337, 335, 361, 362, 338, 334, 53: 469, 57: 370, 113: 332, 682, 117: 351, 354, 350, 331, 714, 123: 618, 333, 126: 330, 138: 329, 144: 715},
		{8: 721},
		{385, 390, 378, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 469, 114: 682, 121: 716},
		// 395
		{8: 717},
		{385, 390, 378, 389, 391, 392, 388, 387, 9: 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 469, 114: 682, 121: 718},
		{9: 719},
		{385, 390, 681, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 678, 469, 55: 637, 85: 703, 87: 700, 705, 690, 704, 691, 701, 702, 698, 706, 699, 114: 682, 121: 697, 638, 127: 689, 684, 685, 687, 688, 683, 134: 686, 720},
		{87, 87, 87, 87, 87, 87, 87, 87, 87, 11: 87, 13: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 43: 87, 87, 87, 87, 87, 87, 87, 52: 87, 55: 87, 57: 87, 85: 87, 87: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 112: 87},
		// 400
		{385, 390, 378, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 469, 114: 682, 121: 722},
		{8: 723},
		{385, 390, 378, 389, 391, 392, 388, 387, 9: 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 469, 114: 682, 121: 724},
		{9: 725},
		{385, 390, 681, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 678, 469, 55: 637, 85: 703, 87: 700, 705, 690, 704, 691, 701, 702, 698, 706, 699, 114: 682, 121: 697, 638, 127: 689, 684, 685, 687, 688, 683, 134: 686, 726},
		// 405
		{88, 88, 88, 88, 88, 88, 88, 88, 88, 11: 88, 13: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 43: 88, 88, 88, 88, 88, 88, 88, 52: 88, 55: 88, 57: 88, 85: 88, 87: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 112: 88},
		{87: 728},
		{729},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 469, 114: 730},
		{9: 731, 472},
		// 410
		{8: 732},
		{89, 89, 89, 89, 89, 89, 89, 89, 89, 11: 89, 13: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 43: 89, 89, 89, 89, 89, 89, 89, 52: 89, 55: 89, 57: 89, 85: 89, 87: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 112: 89},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 469, 114: 734},
		{9: 735, 472},
		{385, 390, 681, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 678, 469, 55: 637, 85: 703, 87: 700, 705, 690, 704, 691, 701, 702, 698, 706, 699, 114: 682, 121: 697, 638, 127: 689, 684, 685, 687, 688, 683, 134: 686, 736},
		// 415
		{90, 90, 90, 90, 90, 90, 90, 90, 90, 11: 90, 13: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 43: 90, 90, 90, 90, 90, 90, 90, 52: 90, 55: 90, 57: 90, 85: 90, 87: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 112: 90},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 469, 114: 738},
		{9: 739, 472},
		{385, 390, 681, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 678, 469, 55: 637, 85: 703, 87: 700, 705, 690, 704, 691, 701, 702, 698, 706, 699, 114: 682, 121: 697, 638, 127: 689, 684, 685, 687, 688, 683, 134: 686, 740},
		{91, 91, 91, 91, 91, 91, 91, 91, 91, 11: 91, 13: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 43: 91, 91, 91, 91, 91, 91, 91, 52: 91, 55: 91, 57: 91, 85: 91, 87: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 112: 91},
		// 420
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 53: 469, 114: 742},
		{9: 743, 472},
		{385, 390, 681, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 678, 469, 55: 637, 85: 703, 87: 700, 705, 690, 704, 691, 701, 702, 698, 706, 699, 114: 682, 121: 697, 638, 127: 689, 684, 685, 687, 688, 683, 134: 686, 744},
		{93, 93, 93, 93, 93, 93, 93, 93, 93, 11: 93, 13: 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 43: 93, 93, 93, 93, 93, 93, 93, 52: 93, 55: 93, 57: 93, 85: 93, 87: 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 112: 745},
		{385, 390, 681, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 678, 469, 55: 637, 85: 703, 87: 700, 705, 690, 704, 691, 701, 702, 698, 706, 699, 114: 682, 121: 697, 638, 127: 689, 684, 685, 687, 688, 683, 134: 686, 746},
		// 425
		{92, 92, 92, 92, 92, 92, 92, 92, 92, 11: 92, 13: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 43: 92, 92, 92, 92, 92, 92, 92, 52: 92, 55: 92, 57: 92, 85: 92, 87: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 112: 92},
		{94, 94, 94, 94, 94, 94, 94, 94, 94, 11: 94, 13: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 43: 94, 94, 94, 94, 94, 94, 94, 52: 94, 55: 94, 57: 94, 85: 94, 87: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 112: 94},
		{99, 99, 99, 99, 99, 99, 99, 99, 99, 11: 99, 13: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 43: 99, 99, 99, 99, 99, 99, 99, 52: 99, 55: 99, 57: 99, 85: 99, 87: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99},
		{101, 101, 101, 101, 101, 101, 101, 101, 101, 11: 101, 13: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 43: 101, 101, 101, 101, 101, 101, 101, 101, 52: 101, 55: 101, 57: 101, 85: 101, 87: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 112: 101},
		{385, 390, 681, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 678, 469, 55: 637, 85: 703, 87: 700, 705, 690, 704, 691, 701, 702, 698, 706, 699, 114: 682, 121: 697, 638, 127: 689, 684, 685, 687, 688, 683, 134: 686, 751},
		// 430
		{103, 103, 103, 103, 103, 103, 103, 103, 103, 11: 103, 13: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 43: 103, 103, 103, 103, 103, 103, 103, 52: 103, 55: 103, 57: 103, 85: 103, 87: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 112: 103},
		{42: 753},
		{385, 390, 681, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 678, 469, 55: 637, 85: 703, 87: 700, 705, 690, 704, 691, 701, 702, 698, 706, 699, 114: 682, 121: 697, 638, 127: 689, 684, 685, 687, 688, 683, 134: 686, 754},
		{104, 104, 104, 104, 104, 104, 104, 104, 104, 11: 104, 13: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 43: 104, 104, 104, 104, 104, 104, 104, 52: 104, 55: 104, 57: 104, 85: 104, 87: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 112: 104},
		{385, 390, 681, 389, 391, 392, 388, 387, 248, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 52: 678, 469, 55: 637, 85: 703, 87: 700, 705, 690, 704, 691, 701, 702, 698, 706, 699, 114: 682, 121: 697, 638, 127: 689, 684, 685, 687, 688, 683, 134: 686, 756},
		// 435
		{105, 105, 105, 105, 105, 105, 105, 105, 105, 11: 105, 13: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 43: 105, 105, 105, 105, 105, 105, 105, 52: 105, 55: 105, 57: 105, 85: 105, 87: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 112: 105},
		{161, 403, 161, 133: 614, 136: 613, 526},
		{78, 78, 78, 8: 78, 22: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 43: 78, 78, 78, 78, 78, 78, 78, 78, 55: 78, 57: 78},
		{52: 187, 212: 761},
		{185, 185, 185, 8: 185, 185, 185, 12: 185, 22: 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 42: 185, 185, 185, 185, 185, 185, 185, 185, 52: 142},
		// 440
		{52: 762},
		{2: 763, 189: 766, 765, 225: 764},
		{10: 314, 41: 314, 51: 314},
		{10: 769, 41: 189, 164: 770},
		{10: 184, 41: 184},
		// 445
		{10: 182, 41: 182, 51: 767},
		{246, 246, 246, 246, 246, 246, 246, 246, 11: 246, 13: 246, 246, 246, 246, 246, 246, 246, 246, 246, 145: 372, 768},
		{10: 181, 41: 181},
		{2: 763, 41: 188, 189: 766, 772},
		{41: 771},
		// 450
		{186, 186, 186, 8: 186, 186, 186, 12: 186, 22: 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 42: 186, 186, 186, 186, 186, 186, 186, 186},
		{10: 183, 41: 183},
		{52: 775},
		{206, 206, 206, 8: 206, 206, 206, 12: 206, 22: 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 42: 206, 206, 206, 206, 206, 206, 206, 206, 52: 142},
		{22: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 777, 57: 208, 211: 776},
		// 455
		{22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 57: 370, 113: 401, 117: 351, 354, 350, 400, 138: 781, 151: 780, 201: 779, 230: 778},
		{205, 205, 205, 8: 205, 205, 205, 12: 205, 22: 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 42: 205, 205, 205, 205, 205, 205, 205, 205},
		{22: 360, 358, 359, 348, 340, 349, 345, 357, 344, 342, 343, 341, 346, 355, 352, 353, 356, 347, 339, 792, 57: 370, 113: 401, 117: 351, 354, 350, 400, 138: 781, 151: 780, 201: 793},
		{22: 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 57: 202},
		{161, 403, 161, 8: 783, 42: 174, 133: 614, 136: 613, 785, 184: 786, 202: 784, 231: 782},
		// 460
		{22: 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 57: 198},
		{8: 789, 10: 790},
		{22: 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 57: 199},
		{8: 193, 10: 193},
		{8: 191, 10: 191, 42: 173},
		// 465
		{42: 787},
		{246, 246, 246, 246, 246, 246, 246, 246, 11: 246, 13: 246, 246, 246, 246, 246, 246, 246, 246, 246, 145: 372, 788},
		{8: 190, 10: 190},
		{22: 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 57: 200},
		{161, 403, 161, 42: 174, 133: 614, 136: 613, 785, 184: 786, 202: 791},
		// 470
		{8: 192, 10: 192},
		{207, 207, 207, 8: 207, 207, 207, 12: 207, 22: 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 42: 207, 207, 207, 207, 207, 207, 207, 207},
		{22: 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 57: 201},
		{385, 390, 378, 389, 391, 392, 388, 387, 11: 384, 13: 394, 393, 396, 379, 380, 381, 382, 383, 395, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 53: 795, 153: 399, 160: 796},
		{422, 427, 3: 440, 430, 431, 426, 425, 9: 798, 12: 421, 51: 446, 56: 423, 58: 429, 428, 434, 435, 445, 441, 442, 450, 443, 454, 424, 448, 438, 437, 436, 432, 452, 449, 447, 439, 456, 444, 433, 453, 451, 455},
		// 475
		{9: 797},
		{209, 209, 209, 8: 209, 209, 209, 12: 209, 22: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 42: 209, 209, 209, 209, 209, 209, 209, 209},
		{210, 210, 210, 8: 210, 210, 210, 12: 210, 22: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 42: 210, 210, 210, 210, 210, 210, 210, 210},
		{239, 239, 239, 8: 239, 239, 239, 12: 239},
		{237, 237, 237, 8: 237, 237, 237, 12: 237},
		// 480
		{240, 240, 240, 8: 240, 240, 240, 12: 240},
		{241, 241, 241, 8: 241, 241, 241, 12: 241},
		{242, 242, 242, 8: 242, 242, 242, 12: 242},
		{8: 232, 10: 232, 22: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 43: 64, 64, 64, 64, 64, 64, 64, 51: 231, 65, 55: 65, 57: 64, 178: 617, 626, 183: 805},
		{52: 76, 55: 76, 207: 806},
		// 485
		{52: 71, 55: 69, 176: 634, 635, 192: 807},
		{75, 75, 75, 8: 75, 22: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 43: 75, 75, 75, 75, 75, 75, 75, 75, 55: 75, 57: 75},
		{81, 81, 81, 8: 81, 22: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 43: 81, 81, 81, 81, 81, 81, 81, 81, 55: 81, 57: 81},
		{246, 246, 246, 246, 246, 246, 246, 246, 11: 246, 13: 246, 246, 246, 246, 246, 246, 246, 246, 246, 145: 372, 810},
		{50: 317},
		// 490
		{84: 833, 86: 835, 100: 823, 824, 825, 820, 821, 822, 826, 830, 827, 817, 828, 829, 115: 834, 832, 125: 831, 139: 815, 814, 819, 816, 818, 148: 813, 228: 812},
		{50: 319},
		{50: 43, 84: 833, 86: 835, 100: 823, 824, 825, 820, 821, 822, 826, 830, 827, 817, 828, 829, 115: 834, 832, 125: 831, 139: 815, 873, 819, 816, 818},
		{50: 42, 84: 42, 86: 42, 97: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{50: 38, 84: 38, 86: 38, 97: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		// 495
		{50: 37, 84: 37, 86: 37, 97: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{86: 835, 115: 895, 832},
		{50: 35, 84: 35, 86: 35, 97: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{97: 28, 28, 883, 188: 881, 220: 882, 880},
		{86: 835, 115: 877, 832},
		// 500
		{2: 874},
		{2: 869},
		{2: 850, 84: 852, 226: 851},
		{84: 833, 86: 835, 115: 834, 832, 125: 849},
		{50: 16, 84: 16, 86: 16, 97: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		// 505
		{86: 835, 115: 847, 832},
		{86: 835, 115: 845, 832},
		{84: 833, 86: 835, 115: 834, 832, 125: 844},
		{2: 840},
		{86: 835, 115: 838, 832},
		// 510
		{50: 7, 84: 7, 86: 7, 97: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{84: 5, 86: 837},
		{50: 4, 84: 4, 86: 4, 97: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{84: 836},
		{84: 2, 86: 2},
		// 515
		{50: 3, 84: 3, 86: 3, 97: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{84: 1, 86: 1},
		{84: 839},
		{50: 8, 84: 8, 86: 8, 97: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{84: 841, 86: 835, 115: 842, 832},
		// 520
		{50: 12, 84: 12, 86: 12, 97: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{84: 843},
		{50: 9, 84: 9, 86: 9, 97: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{50: 13, 84: 13, 86: 13, 97: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{84: 846},
		// 525
		{50: 14, 84: 14, 86: 14, 97: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{84: 848},
		{50: 15, 84: 15, 86: 15, 97: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{50: 17, 84: 17, 86: 17, 97: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{84: 833, 86: 835, 115: 834, 832, 125: 858, 150: 868},
		// 530
		{2: 532, 9: 145, 147: 854, 169: 853, 194: 855},
		{50: 10, 84: 10, 86: 10, 97: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{9: 144, 861, 147: 862},
		{9: 859},
		{9: 856},
		// 535
		{84: 833, 86: 835, 115: 834, 832, 125: 858, 150: 857},
		{50: 18, 84: 18, 86: 18, 97: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{50: 6, 84: 6, 86: 6, 97: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{84: 833, 86: 835, 115: 834, 832, 125: 858, 150: 860},
		{50: 20, 84: 20, 86: 20, 97: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		// 540
		{2: 535, 147: 865},
		{9: 863},
		{84: 833, 86: 835, 115: 834, 832, 125: 858, 150: 864},
		{50: 11, 84: 11, 86: 11, 97: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{9: 866},
		// 545
		{84: 833, 86: 835, 115: 834, 832, 125: 858, 150: 867},
		{50: 19, 84: 19, 86: 19, 97: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{50: 21, 84: 21, 86: 21, 97: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{84: 870},
		{84: 833, 86: 835, 97: 40, 40, 40, 823, 824, 825, 820, 821, 822, 826, 830, 827, 817, 828, 829, 115: 834, 832, 125: 831, 139: 815, 814, 819, 816, 818, 148: 871, 872},
		// 550
		{84: 833, 86: 835, 97: 39, 39, 39, 823, 824, 825, 820, 821, 822, 826, 830, 827, 817, 828, 829, 115: 834, 832, 125: 831, 139: 815, 873, 819, 816, 818},
		{97: 31, 31, 31},
		{50: 41, 84: 41, 86: 41, 97: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		{84: 875},
		{84: 833, 86: 835, 97: 40, 40, 40, 823, 824, 825, 820, 821, 822, 826, 830, 827, 817, 828, 829, 115: 834, 832, 125: 831, 139: 815, 814, 819, 816, 818, 148: 871, 876},
		// 555
		{97: 32, 32, 32},
		{84: 878},
		{84: 833, 86: 835, 97: 40, 40, 40, 823, 824, 825, 820, 821, 822, 826, 830, 827, 817, 828, 829, 115: 834, 832, 125: 831, 139: 815, 814, 819, 816, 818, 148: 871, 879},
		{97: 33, 33, 33},
		{97: 24, 889, 222: 890, 888},
		// 560
		{97: 30, 30, 30},
		{97: 27, 27, 883, 188: 887},
		{86: 835, 115: 884, 832},
		{84: 885},
		{84: 833, 86: 835, 97: 40, 40, 40, 823, 824, 825, 820, 821, 822, 826, 830, 827, 817, 828, 829, 115: 834, 832, 125: 831, 139: 815, 814, 819, 816, 818, 148: 871, 886},
		// 565
		{97: 26, 26, 26},
		{97: 29, 29, 29},
		{97: 894, 224: 893},
		{84: 891},
		{97: 23},
		// 570
		{84: 833, 86: 835, 97: 40, 100: 823, 824, 825, 820, 821, 822, 826, 830, 827, 817, 828, 829, 115: 834, 832, 125: 831, 139: 815, 814, 819, 816, 818, 148: 871, 892},
		{97: 25},
		{50: 34, 84: 34, 86: 34, 97: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{50: 22, 84: 22, 86: 22, 97: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{84: 896},
		// 575
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
			//TODO- outer2:
			//TODO- 	switch i := lhs.Initializer; i.Case {
			//TODO- 	case 0: // Expression
			//TODO- 		e := i.Expression
			//TODO- 		et := e.Type
			//TODO- 		d := lhs.Declarator
			//TODO- 		dt := d.Type
			//TODO- 		switch x := e.Value.(type) {
			//TODO- 		case StringLitID:
			//TODO- 			if dt.Kind() != Ptr && dt.Kind() != Array {
			//TODO- 				break
			//TODO- 			}

			//TODO- 			switch dt.Element().Kind() {
			//TODO- 			case Char, SChar, UChar:
			//TODO- 				if dt.Kind() != Array {
			//TODO- 					d.Type = d.Type.(*ctype).setElements(len(xc.Dict.S(int(x)))+1)
			//TODO- 				}
			//TODO- 				break outer2
			//TODO- 			}
			//TODO- 		case LongStringLitID:
			//TODO- 			if dt.Kind() != Ptr && dt.Kind() != Array {
			//TODO- 				break
			//TODO- 			}

			//TODO- 			if IsIntType(dt.Element()) {
			//TODO- 				if dt.Kind() != Array {
			//TODO- 					d.Type = d.Type.(*ctype).setElements(len(xc.Dict.S(int(x)))+1)
			//TODO- 				}
			//TODO- 				break outer2
			//TODO- 			}
			//TODO- 		}

			//TODO- 		if !et.CanAssignTo(dt) {
			//TODO- 			lx.report.Err(i.Pos(), "incompatible types when initializing type '%s' using type ‘%s'", dt, et)
			//TODO- 		}
			//TODO- 	case 1: // '{' InitializerList CommaOpt '}'
			//TODO- 		limit := -1
			//TODO- 		var checkType Type
			//TODO- 		var mb []Member
			//TODO- 		var incomplete bool
			//TODO- 		d := lhs.Declarator
			//TODO- 		k := d.Type.Kind()
			//TODO- 		switch k {
			//TODO- 		case Array:
			//TODO- 			checkType = d.Type.Element()
			//TODO- 			limit = nElem(d.Type)
			//TODO- 		case Ptr:
			//TODO- 			checkType = d.Type.Element()
			//TODO- 			d.Type = d.Type.(*ctype).setElements(i.InitializerList.Len())
			//TODO- 		case Struct, Union:
			//TODO- 			mb, incomplete = d.Type.Members()
			//TODO- 			if mb == nil {
			//TODO- 				panic("internal error")
			//TODO- 			}

			//TODO- 			limit = len(mb)
			//TODO- 			if k == Union {
			//TODO- 				limit = 1
			//TODO- 			}
			//TODO- 		default:
			//TODO- 			limit = 1
			//TODO- 			checkType = d.Type
			//TODO- 			mb = []Member{
			//TODO- 				{Declarator: d, Type: d.Type},
			//TODO- 			}
			//TODO- 		}

			//TODO- 		values := 0
			//TODO- 		for l := i.InitializerList; l != nil; l = l.InitializerList {
			//TODO- 			values++
			//TODO- 			if incomplete {
			//TODO- 				lx.report.Err(i.Pos(), "variable/field has initializer but incomplete type")
			//TODO- 				break
			//TODO- 			}

			//TODO- 			if o := l.DesignationOpt; o != nil {
			//TODO- 				var a []*Designator
			//TODO- 				for l := o.Designation.DesignatorList; l != nil; l = l.DesignatorList {
			//TODO- 					a = append(a, l.Designator)
			//TODO- 				}
			//TODO- 				if len(a) != 1 {
			//TODO- 					panic("TODO")
			//TODO- 				}

			//TODO- 			outer:
			//TODO- 				switch des := a[0]; des.Case {
			//TODO- 				case 0: // '[' ConstantExpression ']'
			//TODO- 					e := des.ConstantExpression
			//TODO- 					if !IsIntType(e.Type) {
			//TODO- 						lx.report.Err(e.Pos(), "index expression is not an integer type (have '%s')", e.Type)
			//TODO- 						break outer
			//TODO- 					}

			//TODO- 					var ix uint64
			//TODO- 					valid := true
			//TODO- 					switch x := e.Value.(type) {
			//TODO- 					case int32:
			//TODO- 						valid = x >= 0
			//TODO- 						ix = uint64(x)
			//TODO- 					case uint32:
			//TODO- 						ix = uint64(x)
			//TODO- 					case int64:
			//TODO- 						valid = x >= 0
			//TODO- 						ix = uint64(x)
			//TODO- 					case uint64:
			//TODO- 						ix = x
			//TODO- 					default:
			//TODO- 						panic("TODO")
			//TODO- 					}
			//TODO- 					if !valid {
			//TODO- 						lx.report.Err(e.Pos(), "index must be non-negative (have '%v')", e.Value)
			//TODO- 					}

			//TODO- 					if limit >= 0 && ix >= uint64(limit) {
			//TODO- 						lx.report.Err(e.Pos(), "index value out of bounds (have '%v', limit '%v')", e.Value, limit-1)
			//TODO- 					}
			//TODO- 					l.Initializer.typeCheck(checkType, mb, int(ix), limit, lx)
			//TODO- 				case 1: // '.' IDENTIFIER              // Case 1
			//TODO- 					id := des.Token2.Val
			//TODO- 					for i, v := range mb {
			//TODO- 						if v.Name == id {
			//TODO- 							l.Initializer.typeCheck(checkType, mb, i, limit, lx)
			//TODO- 							break outer
			//TODO- 						}
			//TODO- 					}

			//TODO- 					lx.report.Err(des.Token2.Pos(), "type '%s' has no member '%s'", d.Type, dict.S(id))
			//TODO- 				default:
			//TODO- 					panic("internal error")
			//TODO- 				}
			//TODO- 				continue
			//TODO- 			}

			//TODO- 			l.Initializer.typeCheck(checkType, mb, values-1, limit, lx)
			//TODO- 		}
			//TODO- 	default:
			//TODO- 		panic(i.Case)
			//TODO- 	}
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
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 203:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 204:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 205:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 206:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 207:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 208:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 209:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 210:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 211:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 212:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 213:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 214:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 215:
		{
			yyVAL.node = &Statement{
				Case:               6,
				AssemblerStatement: yyS[yypt-0].node.(*AssemblerStatement),
			}
		}
	case 216:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 217:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 218:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 219:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 220:
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
	case 221:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 222:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 223:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 224:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 225:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 226:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 227:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 228:
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
	case 229:
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
	case 230:
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
	case 231:
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
	case 232:
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
	case 233:
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
	case 234:
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
	case 235:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 236:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 237:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 238:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 239:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 240:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 241:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 242:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 243:
		{
			yyVAL.node = &ExternalDeclaration{
				Case: 2,
				BasicAssemblerStatement: yyS[yypt-1].node.(*BasicAssemblerStatement),
				Token: yyS[yypt-0].Token,
			}
		}
	case 244:
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
	case 245:
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
	case 246:
		{
			yyVAL.node = &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:          yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 247:
		{
			lx := yylex.(*lexer)
			lx.scope.specifier = &DeclarationSpecifiers{typeSpecifier: tsEncode(tsInt)}
		}
	case 248:
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
	case 249:
		{
			yyVAL.node = &FunctionDefinition{
				Case:               1,
				Declarator:         yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt: yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:       yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 250:
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
	case 251:
		{
			lhs := &FunctionBody{
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
			yyVAL.node = lhs
			lhs.scope = lhs.CompoundStatement.scope
		}
	case 252:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 253:
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
	case 254:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 255:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 256:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 257:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 258:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScopePos(lhs.Pos())
		}
	case 259:
		{
			yyVAL.node = &AssemblerInstructions{
				Token: yyS[yypt-0].Token,
			}
		}
	case 260:
		{
			yyVAL.node = &AssemblerInstructions{
				Case: 1,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions),
				Token: yyS[yypt-0].Token,
			}
		}
	case 261:
		{
			yyVAL.node = &BasicAssemblerStatement{
				Token:                 yyS[yypt-4].Token,
				VolatileOpt:           yyS[yypt-3].node.(*VolatileOpt),
				Token2:                yyS[yypt-2].Token,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-0].Token,
			}
		}
	case 262:
		{
			yyVAL.node = (*VolatileOpt)(nil)
		}
	case 263:
		{
			yyVAL.node = &VolatileOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 264:
		{
			yyVAL.node = &AssemblerOperand{
				AssemblerSymbolicNameOpt: yyS[yypt-4].node.(*AssemblerSymbolicNameOpt),
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
		}
	case 265:
		{
			yyVAL.node = &AssemblerOperands{
				AssemblerOperand: yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 266:
		{
			yyVAL.node = &AssemblerOperands{
				Case:              1,
				AssemblerOperands: yyS[yypt-2].node.(*AssemblerOperands),
				Token:             yyS[yypt-1].Token,
				AssemblerOperand:  yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 267:
		{
			yyVAL.node = (*AssemblerSymbolicNameOpt)(nil)
		}
	case 268:
		{
			yyVAL.node = &AssemblerSymbolicNameOpt{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 269:
		{
			yyVAL.node = &Clobbers{
				Token: yyS[yypt-0].Token,
			}
		}
	case 270:
		{
			yyVAL.node = &Clobbers{
				Case:     1,
				Clobbers: yyS[yypt-2].node.(*Clobbers),
				Token:    yyS[yypt-1].Token,
				Token2:   yyS[yypt-0].Token,
			}
		}
	case 271:
		{
			yyVAL.node = &AssemblerStatement{
				BasicAssemblerStatement: yyS[yypt-0].node.(*BasicAssemblerStatement),
			}
		}
	case 272:
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
	case 273:
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
	case 274:
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
	case 275:
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
	case 276:
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
	case 277:
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
	case 278:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 279:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 280:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 281:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 282:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 283:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 284:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 285:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 286:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 287:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 288:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 289:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 290:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 291:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 292:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 293:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 294:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 295:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 296:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 297:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 298:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 299:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 300:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 301:
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
	case 302:
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
	case 303:
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
	case 304:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 305:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 306:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 307:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 308:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 309:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 310:
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
	case 311:
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
	case 312:
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
	case 313:
		{
			yyVAL.node = &ControlLine{
				Case:        13,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 316:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 317:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
