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
	yyDefault           = 57455
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
	yyTabOfs   = -312
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (299x)
		42:      1,   // '*' (262x)
		57376:   2,   // IDENTIFIER (216x)
		38:      3,   // '&' (210x)
		43:      4,   // '+' (210x)
		45:      5,   // '-' (210x)
		57362:   6,   // DEC (210x)
		57380:   7,   // INC (210x)
		41:      8,   // ')' (194x)
		59:      9,   // ';' (191x)
		44:      10,  // ',' (181x)
		57424:   11,  // STRINGLITERAL (158x)
		91:      12,  // '[' (157x)
		33:      13,  // '!' (141x)
		126:     14,  // '~' (141x)
		57357:   15,  // CHARCONST (141x)
		57372:   16,  // FLOATCONST (141x)
		57383:   17,  // INTCONST (141x)
		57386:   18,  // LONGCHARCONST (141x)
		57387:   19,  // LONGSTRINGLITERAL (141x)
		57422:   20,  // SIZEOF (141x)
		57435:   21,  // VOLATILE (129x)
		57359:   22,  // CONST (127x)
		57416:   23,  // RESTRICT (127x)
		58:      24,  // ':' (118x)
		125:     25,  // '}' (117x)
		57352:   26,  // BOOL (117x)
		57356:   27,  // CHAR (117x)
		57358:   28,  // COMPLEX (117x)
		57366:   29,  // DOUBLE (117x)
		57368:   30,  // ENUM (117x)
		57371:   31,  // FLOAT (117x)
		57382:   32,  // INT (117x)
		57385:   33,  // LONG (117x)
		57420:   34,  // SHORT (117x)
		57421:   35,  // SIGNED (117x)
		57425:   36,  // STRUCT (117x)
		57429:   37,  // TYPEDEFNAME (117x)
		57430:   38,  // TYPEOF (117x)
		57432:   39,  // UNION (117x)
		57433:   40,  // UNSIGNED (117x)
		57434:   41,  // VOID (117x)
		57423:   42,  // STATIC (110x)
		57351:   43,  // AUTO (104x)
		57370:   44,  // EXTERN (104x)
		57381:   45,  // INLINE (104x)
		57415:   46,  // REGISTER (104x)
		57428:   47,  // TYPEDEF (104x)
		57344:   48,  // $end (100x)
		61:      49,  // '=' (89x)
		57498:   50,  // Expression (85x)
		123:     51,  // '{' (83x)
		93:      52,  // ']' (82x)
		46:      53,  // '.' (78x)
		57350:   54,  // ASM (74x)
		37:      55,  // '%' (70x)
		47:      56,  // '/' (70x)
		60:      57,  // '<' (70x)
		62:      58,  // '>' (70x)
		63:      59,  // '?' (70x)
		94:      60,  // '^' (70x)
		124:     61,  // '|' (70x)
		57346:   62,  // ADDASSIGN (70x)
		57347:   63,  // ANDAND (70x)
		57348:   64,  // ANDASSIGN (70x)
		57349:   65,  // ARROW (70x)
		57364:   66,  // DIVASSIGN (70x)
		57369:   67,  // EQ (70x)
		57374:   68,  // GEQ (70x)
		57384:   69,  // LEQ (70x)
		57388:   70,  // LSH (70x)
		57389:   71,  // LSHASSIGN (70x)
		57390:   72,  // MODASSIGN (70x)
		57391:   73,  // MULASSIGN (70x)
		57392:   74,  // NEQ (70x)
		57394:   75,  // ORASSIGN (70x)
		57395:   76,  // OROR (70x)
		57418:   77,  // RSH (70x)
		57419:   78,  // RSHASSIGN (70x)
		57426:   79,  // SUBASSIGN (70x)
		57437:   80,  // XORASSIGN (70x)
		10:      81,  // '\n' (60x)
		57411:   82,  // PPOTHER (54x)
		57375:   83,  // GOTO (50x)
		57436:   84,  // WHILE (48x)
		57353:   85,  // BREAK (47x)
		57354:   86,  // CASE (47x)
		57360:   87,  // CONTINUE (47x)
		57363:   88,  // DEFAULT (47x)
		57365:   89,  // DO (47x)
		57373:   90,  // FOR (47x)
		57379:   91,  // IF (47x)
		57417:   92,  // RETURN (47x)
		57427:   93,  // SWITCH (47x)
		57399:   94,  // PPENDIF (45x)
		57398:   95,  // PPELSE (41x)
		57397:   96,  // PPELIF (40x)
		57396:   97,  // PPDEFINE (36x)
		57400:   98,  // PPERROR (36x)
		57401:   99,  // PPHASH_NL (36x)
		57403:   100, // PPIF (36x)
		57404:   101, // PPIFDEF (36x)
		57405:   102, // PPIFNDEF (36x)
		57406:   103, // PPINCLUDE (36x)
		57407:   104, // PPINCLUDE_NEXT (36x)
		57408:   105, // PPLINE (36x)
		57409:   106, // PPNONDIRECTIVE (36x)
		57413:   107, // PPPRAGMA (36x)
		57414:   108, // PPUNDEF (36x)
		57367:   109, // ELSE (29x)
		57549:   110, // TypeQualifier (28x)
		57499:   111, // ExpressionList (26x)
		57523:   112, // PPTokenList (23x)
		57525:   113, // PPTokens (23x)
		57494:   114, // EnumSpecifier (20x)
		57544:   115, // StructOrUnion (20x)
		57545:   116, // StructOrUnionSpecifier (20x)
		57552:   117, // TypeSpecifier (20x)
		57500:   118, // ExpressionListOpt (18x)
		57524:   119, // PPTokenListOpt (16x)
		57465:   120, // BasicAssemblerStatement (15x)
		57477:   121, // DeclarationSpecifiers (15x)
		57506:   122, // FunctionSpecifier (15x)
		57539:   123, // StorageClassSpecifier (15x)
		57463:   124, // AssemblerStatement (13x)
		57471:   125, // CompoundStatement (13x)
		57502:   126, // ExpressionStatement (12x)
		57520:   127, // IterationStatement (12x)
		57521:   128, // JumpStatement (12x)
		57522:   129, // LabeledStatement (12x)
		57534:   130, // SelectionStatement (12x)
		57538:   131, // Statement (12x)
		57530:   132, // Pointer (11x)
		57531:   133, // PointerOpt (10x)
		57473:   134, // ControlLine (8x)
		57479:   135, // Declarator (8x)
		57509:   136, // GroupPart (8x)
		57513:   137, // IfGroup (8x)
		57514:   138, // IfSection (8x)
		57546:   139, // TextLine (8x)
		57474:   140, // Declaration (7x)
		57507:   141, // GroupList (6x)
		57533:   142, // ReplacementList (6x)
		57449:   143, // $@4 (5x)
		57472:   144, // ConstantExpression (5x)
		57361:   145, // DDD (5x)
		57508:   146, // GroupListOpt (5x)
		57535:   147, // SpecifierQualifierList (5x)
		57550:   148, // TypeQualifierList (5x)
		57456:   149, // AbstractDeclarator (4x)
		57461:   150, // AssemblerOperand (4x)
		57464:   151, // AssemblerSymbolicNameOpt (4x)
		57478:   152, // DeclarationSpecifiersOpt (4x)
		57483:   153, // Designator (4x)
		57526:   154, // ParameterDeclaration (4x)
		57551:   155, // TypeQualifierListOpt (4x)
		57439:   156, // $@10 (3x)
		57460:   157, // AssemblerInstructions (3x)
		57462:   158, // AssemblerOperands (3x)
		57470:   159, // CommaOpt (3x)
		57481:   160, // Designation (3x)
		57482:   161, // DesignationOpt (3x)
		57484:   162, // DesignatorList (3x)
		57501:   163, // ExpressionOpt (3x)
		57510:   164, // IdentifierList (3x)
		57515:   165, // InitDeclarator (3x)
		57518:   166, // Initializer (3x)
		57527:   167, // ParameterList (3x)
		57528:   168, // ParameterTypeList (3x)
		57548:   169, // TypeName (3x)
		57440:   170, // $@11 (2x)
		57450:   171, // $@5 (2x)
		57457:   172, // AbstractDeclaratorOpt (2x)
		57466:   173, // BlockItem (2x)
		57469:   174, // Clobbers (2x)
		57480:   175, // DeclaratorOpt (2x)
		57485:   176, // DirectAbstractDeclarator (2x)
		57486:   177, // DirectAbstractDeclaratorOpt (2x)
		57487:   178, // DirectDeclarator (2x)
		57488:   179, // ElifGroup (2x)
		57495:   180, // EnumerationConstant (2x)
		57496:   181, // Enumerator (2x)
		57503:   182, // ExternalDeclaration (2x)
		57505:   183, // FunctionDefinition (2x)
		57511:   184, // IdentifierListOpt (2x)
		57512:   185, // IdentifierOpt (2x)
		57516:   186, // InitDeclaratorList (2x)
		57517:   187, // InitDeclaratorListOpt (2x)
		57519:   188, // InitializerList (2x)
		57529:   189, // ParameterTypeListOpt (2x)
		57536:   190, // SpecifierQualifierListOpt (2x)
		57540:   191, // StructDeclaration (2x)
		57542:   192, // StructDeclarator (2x)
		57553:   193, // VolatileOpt (2x)
		57438:   194, // $@1 (1x)
		57441:   195, // $@12 (1x)
		57442:   196, // $@13 (1x)
		57443:   197, // $@14 (1x)
		57444:   198, // $@15 (1x)
		57445:   199, // $@16 (1x)
		57446:   200, // $@17 (1x)
		57447:   201, // $@2 (1x)
		57448:   202, // $@3 (1x)
		57451:   203, // $@6 (1x)
		57452:   204, // $@7 (1x)
		57453:   205, // $@8 (1x)
		57454:   206, // $@9 (1x)
		57458:   207, // ArgumentExpressionList (1x)
		57459:   208, // ArgumentExpressionListOpt (1x)
		57467:   209, // BlockItemList (1x)
		57468:   210, // BlockItemListOpt (1x)
		1048577: 211, // CONSTANT_EXPRESSION (1x)
		57475:   212, // DeclarationList (1x)
		57476:   213, // DeclarationListOpt (1x)
		57489:   214, // ElifGroupList (1x)
		57490:   215, // ElifGroupListOpt (1x)
		57491:   216, // ElseGroup (1x)
		57492:   217, // ElseGroupOpt (1x)
		57493:   218, // EndifLine (1x)
		57497:   219, // EnumeratorList (1x)
		57504:   220, // FunctionBody (1x)
		57377:   221, // IDENTIFIER_LPAREN (1x)
		1048576: 222, // PREPROCESSING_FILE (1x)
		57532:   223, // PreprocessingFile (1x)
		57537:   224, // Start (1x)
		57541:   225, // StructDeclarationList (1x)
		57543:   226, // StructDeclaratorList (1x)
		1048578: 227, // TRANSLATION_UNIT (1x)
		57547:   228, // TranslationUnit (1x)
		57455:   229, // $default (0x)
		57355:   230, // CAST (0x)
		57345:   231, // error (0x)
		57378:   232, // IDENTIFIER_NONREPL (0x)
		57393:   233, // NOELSE (0x)
		57402:   234, // PPHEADER_NAME (0x)
		57410:   235, // PPNUMBER (0x)
		57412:   236, // PPPASTE (0x)
		57431:   237, // UNARY (0x)
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
		"CHARCONST",
		"FLOATCONST",
		"INTCONST",
		"LONGCHARCONST",
		"LONGSTRINGLITERAL",
		"SIZEOF",
		"VOLATILE",
		"CONST",
		"RESTRICT",
		"':'",
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
		"STATIC",
		"AUTO",
		"EXTERN",
		"INLINE",
		"REGISTER",
		"TYPEDEF",
		"$end",
		"'='",
		"Expression",
		"'{'",
		"']'",
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
		"PPTokenListOpt",
		"BasicAssemblerStatement",
		"DeclarationSpecifiers",
		"FunctionSpecifier",
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
		"GroupList",
		"ReplacementList",
		"$@4",
		"ConstantExpression",
		"DDD",
		"GroupListOpt",
		"SpecifierQualifierList",
		"TypeQualifierList",
		"AbstractDeclarator",
		"AssemblerOperand",
		"AssemblerSymbolicNameOpt",
		"DeclarationSpecifiersOpt",
		"Designator",
		"ParameterDeclaration",
		"TypeQualifierListOpt",
		"$@10",
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
		"TypeName",
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
		57376:   "identifier",
		57362:   "--",
		57380:   "++",
		57424:   "string literal",
		57357:   "character constant",
		57372:   "floating-point constant",
		57383:   "integer constant",
		57386:   "long character constant",
		57387:   "long string constant",
		57422:   "sizeof",
		57435:   "volatile",
		57359:   "const",
		57416:   "restrict",
		57352:   "_Bool",
		57356:   "char",
		57358:   "_Complex",
		57366:   "double",
		57368:   "enum",
		57371:   "float",
		57382:   "int",
		57385:   "long",
		57420:   "short",
		57421:   "signed",
		57425:   "struct",
		57429:   "typedefname",
		57430:   "typeof",
		57432:   "union",
		57433:   "unsigned",
		57434:   "void",
		57423:   "static",
		57351:   "auto",
		57370:   "extern",
		57381:   "inline",
		57415:   "register",
		57428:   "typedef",
		57350:   "asm",
		57346:   "+=",
		57347:   "&&",
		57348:   "&=",
		57349:   "->",
		57364:   "/=",
		57369:   "==",
		57374:   ">=",
		57384:   "<=",
		57388:   "<<",
		57389:   "<<=",
		57390:   "%=",
		57391:   "*=",
		57392:   "!=",
		57394:   "|=",
		57395:   "||",
		57418:   ">>",
		57419:   ">>=",
		57426:   "-=",
		57437:   "^=",
		57411:   "ppother",
		57375:   "goto",
		57436:   "while",
		57353:   "break",
		57354:   "case",
		57360:   "continue",
		57363:   "default",
		57365:   "do",
		57373:   "for",
		57379:   "if",
		57417:   "return",
		57427:   "switch",
		57399:   "#endif",
		57398:   "#else",
		57397:   "#elif",
		57396:   "#define",
		57400:   "#error",
		57401:   "#",
		57403:   "#if",
		57404:   "#ifdef",
		57405:   "#ifndef",
		57406:   "#include",
		57407:   "#include_next",
		57408:   "#line",
		57409:   "#foo",
		57413:   "#pragma",
		57414:   "#undef",
		57367:   "else",
		57361:   "...",
		1048577: "constant expression prefix",
		57377:   "identifier immediatelly followed by '('",
		1048576: "preprocessing file prefix",
		1048578: "translation unit prefix",
		57378:   "non replaceable identifier",
		57402:   "header name",
		57410:   "preprocessing number",
		57412:   "##",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {194, 0},
		2:   {224, 3},
		3:   {201, 0},
		4:   {224, 3},
		5:   {202, 0},
		6:   {224, 3},
		7:   {180, 1},
		8:   {207, 1},
		9:   {207, 3},
		10:  {208, 0},
		11:  {208, 1},
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
		68:  {163, 0},
		69:  {163, 1},
		70:  {111, 1},
		71:  {111, 3},
		72:  {118, 0},
		73:  {118, 1},
		74:  {143, 0},
		75:  {144, 2},
		76:  {140, 3},
		77:  {121, 2},
		78:  {121, 2},
		79:  {121, 2},
		80:  {121, 2},
		81:  {152, 0},
		82:  {152, 1},
		83:  {186, 1},
		84:  {186, 3},
		85:  {187, 0},
		86:  {187, 1},
		87:  {165, 1},
		88:  {171, 0},
		89:  {165, 4},
		90:  {123, 1},
		91:  {123, 1},
		92:  {123, 1},
		93:  {123, 1},
		94:  {123, 1},
		95:  {117, 1},
		96:  {117, 1},
		97:  {117, 1},
		98:  {117, 1},
		99:  {117, 1},
		100: {117, 1},
		101: {117, 1},
		102: {117, 1},
		103: {117, 1},
		104: {117, 1},
		105: {117, 1},
		106: {117, 1},
		107: {117, 1},
		108: {117, 1},
		109: {117, 4},
		110: {117, 4},
		111: {203, 0},
		112: {204, 0},
		113: {116, 7},
		114: {116, 2},
		115: {115, 1},
		116: {115, 1},
		117: {225, 1},
		118: {225, 2},
		119: {191, 3},
		120: {191, 2},
		121: {147, 2},
		122: {147, 2},
		123: {190, 0},
		124: {190, 1},
		125: {226, 1},
		126: {226, 3},
		127: {192, 1},
		128: {192, 3},
		129: {159, 0},
		130: {159, 1},
		131: {205, 0},
		132: {114, 7},
		133: {114, 2},
		134: {219, 1},
		135: {219, 3},
		136: {181, 1},
		137: {181, 3},
		138: {110, 1},
		139: {110, 1},
		140: {110, 1},
		141: {122, 1},
		142: {135, 2},
		143: {175, 0},
		144: {175, 1},
		145: {178, 1},
		146: {178, 3},
		147: {178, 5},
		148: {178, 6},
		149: {178, 6},
		150: {178, 5},
		151: {206, 0},
		152: {178, 5},
		153: {178, 4},
		154: {132, 2},
		155: {132, 3},
		156: {133, 0},
		157: {133, 1},
		158: {148, 1},
		159: {148, 2},
		160: {155, 0},
		161: {155, 1},
		162: {168, 1},
		163: {168, 3},
		164: {189, 0},
		165: {189, 1},
		166: {167, 1},
		167: {167, 3},
		168: {154, 2},
		169: {154, 2},
		170: {164, 1},
		171: {164, 3},
		172: {184, 0},
		173: {184, 1},
		174: {185, 0},
		175: {185, 1},
		176: {156, 0},
		177: {169, 3},
		178: {149, 1},
		179: {149, 2},
		180: {172, 0},
		181: {172, 1},
		182: {176, 3},
		183: {176, 4},
		184: {176, 5},
		185: {176, 6},
		186: {176, 6},
		187: {176, 4},
		188: {170, 0},
		189: {176, 4},
		190: {195, 0},
		191: {176, 5},
		192: {177, 0},
		193: {177, 1},
		194: {166, 1},
		195: {166, 4},
		196: {188, 2},
		197: {188, 4},
		198: {160, 2},
		199: {161, 0},
		200: {161, 1},
		201: {162, 1},
		202: {162, 2},
		203: {153, 3},
		204: {153, 2},
		205: {131, 1},
		206: {131, 1},
		207: {131, 1},
		208: {131, 1},
		209: {131, 1},
		210: {131, 1},
		211: {131, 1},
		212: {129, 3},
		213: {129, 4},
		214: {129, 3},
		215: {196, 0},
		216: {125, 4},
		217: {209, 1},
		218: {209, 2},
		219: {210, 0},
		220: {210, 1},
		221: {173, 1},
		222: {173, 1},
		223: {126, 2},
		224: {130, 5},
		225: {130, 7},
		226: {130, 5},
		227: {127, 5},
		228: {127, 7},
		229: {127, 9},
		230: {127, 8},
		231: {128, 3},
		232: {128, 2},
		233: {128, 2},
		234: {128, 3},
		235: {228, 1},
		236: {228, 2},
		237: {182, 1},
		238: {182, 1},
		239: {182, 1},
		240: {197, 0},
		241: {183, 5},
		242: {198, 0},
		243: {220, 2},
		244: {199, 0},
		245: {220, 3},
		246: {212, 1},
		247: {212, 2},
		248: {213, 0},
		249: {200, 0},
		250: {213, 2},
		251: {157, 1},
		252: {157, 2},
		253: {120, 5},
		254: {193, 0},
		255: {193, 1},
		256: {150, 5},
		257: {158, 1},
		258: {158, 3},
		259: {151, 0},
		260: {151, 3},
		261: {174, 1},
		262: {174, 3},
		263: {124, 1},
		264: {124, 7},
		265: {124, 9},
		266: {124, 11},
		267: {124, 13},
		268: {223, 1},
		269: {141, 1},
		270: {141, 2},
		271: {146, 0},
		272: {146, 1},
		273: {136, 1},
		274: {136, 1},
		275: {136, 3},
		276: {136, 1},
		277: {138, 4},
		278: {137, 4},
		279: {137, 4},
		280: {137, 4},
		281: {214, 1},
		282: {214, 2},
		283: {215, 0},
		284: {215, 1},
		285: {179, 4},
		286: {216, 3},
		287: {217, 0},
		288: {217, 1},
		289: {218, 1},
		290: {134, 3},
		291: {134, 5},
		292: {134, 7},
		293: {134, 5},
		294: {134, 2},
		295: {134, 1},
		296: {134, 3},
		297: {134, 3},
		298: {134, 2},
		299: {134, 3},
		300: {134, 6},
		301: {134, 8},
		302: {134, 2},
		303: {134, 4},
		304: {134, 3},
		305: {139, 1},
		306: {142, 1},
		307: {112, 1},
		308: {119, 1},
		309: {119, 2},
		310: {113, 1},
		311: {113, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 48}:   "invalid empty input",
		yyXError{552, -1}: "expected #endif",
		yyXError{554, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{467, -1}: "expected $end",
		yyXError{469, -1}: "expected $end",
		yyXError{31, -1}:  "expected '('",
		yyXError{45, -1}:  "expected '('",
		yyXError{345, -1}: "expected '('",
		yyXError{366, -1}: "expected '('",
		yyXError{401, -1}: "expected '('",
		yyXError{402, -1}: "expected '('",
		yyXError{403, -1}: "expected '('",
		yyXError{405, -1}: "expected '('",
		yyXError{431, -1}: "expected '('",
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
		yyXError{421, -1}: "expected ')'",
		yyXError{427, -1}: "expected ')'",
		yyXError{512, -1}: "expected ')'",
		yyXError{513, -1}: "expected ')'",
		yyXError{521, -1}: "expected ')'",
		yyXError{524, -1}: "expected ')'",
		yyXError{527, -1}: "expected ')'",
		yyXError{303, -1}: "expected ':'",
		yyXError{348, -1}: "expected ':'",
		yyXError{394, -1}: "expected ':'",
		yyXError{455, -1}: "expected ':'",
		yyXError{324, -1}: "expected ';'",
		yyXError{340, -1}: "expected ';'",
		yyXError{400, -1}: "expected ';'",
		yyXError{407, -1}: "expected ';'",
		yyXError{408, -1}: "expected ';'",
		yyXError{410, -1}: "expected ';'",
		yyXError{414, -1}: "expected ';'",
		yyXError{417, -1}: "expected ';'",
		yyXError{419, -1}: "expected ';'",
		yyXError{425, -1}: "expected ';'",
		yyXError{434, -1}: "expected ';'",
		yyXError{328, -1}: "expected '='",
		yyXError{177, -1}: "expected '['",
		yyXError{491, -1}: "expected '\\n'",
		yyXError{495, -1}: "expected '\\n'",
		yyXError{499, -1}: "expected '\\n'",
		yyXError{502, -1}: "expected '\\n'",
		yyXError{504, -1}: "expected '\\n'",
		yyXError{531, -1}: "expected '\\n'",
		yyXError{536, -1}: "expected '\\n'",
		yyXError{539, -1}: "expected '\\n'",
		yyXError{546, -1}: "expected '\\n'",
		yyXError{551, -1}: "expected '\\n'",
		yyXError{557, -1}: "expected '\\n'",
		yyXError{183, -1}: "expected ']'",
		yyXError{191, -1}: "expected ']'",
		yyXError{235, -1}: "expected ']'",
		yyXError{262, -1}: "expected ']'",
		yyXError{354, -1}: "expected ']'",
		yyXError{52, -1}:  "expected '{'",
		yyXError{54, -1}:  "expected '{'",
		yyXError{291, -1}: "expected '{'",
		yyXError{293, -1}: "expected '{'",
		yyXError{271, -1}: "expected '}'",
		yyXError{275, -1}: "expected '}'",
		yyXError{288, -1}: "expected '}'",
		yyXError{395, -1}: "expected '}'",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{211, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{176, -1}: "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{47, -1}:  "expected assembler instructions or string literal",
		yyXError{344, -1}: "expected assembler instructions or string literal",
		yyXError{346, -1}: "expected assembler instructions or string literal",
		yyXError{356, -1}: "expected assembler operand or one of ['[', string literal]",
		yyXError{349, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{371, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{374, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{339, -1}: "expected assembler statement or asm",
		yyXError{397, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{357, -1}: "expected clobbers or string literal",
		yyXError{377, -1}: "expected clobbers or string literal",
		yyXError{338, -1}: "expected compound statement or '{'",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{60, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{259, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{307, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{393, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{466, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{330, -1}: "expected declaration list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{333, -1}: "expected declaration or one of ['{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{416, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{306, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{203, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{256, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{206, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{173, -1}: "expected direct abstract declarator or one of ['(', '[']",
		yyXError{304, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{544, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{550, -1}: "expected endif line or #endif",
		yyXError{476, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{542, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{55, -1}:  "expected enumerator list or identifier",
		yyXError{287, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{83, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{107, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{432, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{436, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{440, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{444, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{367, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{184, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{226, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{314, -1}: "expected expression or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{6, -1}:   "expected external declaration or one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{329, -1}: "expected function body or one of ['{', asm]",
		yyXError{336, -1}: "expected function body or one of ['{', asm]",
		yyXError{327, -1}: "expected function body or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{533, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{470, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{85, -1}:  "expected identifier",
		yyXError{86, -1}:  "expected identifier",
		yyXError{220, -1}: "expected identifier",
		yyXError{260, -1}: "expected identifier",
		yyXError{353, -1}: "expected identifier",
		yyXError{406, -1}: "expected identifier",
		yyXError{478, -1}: "expected identifier",
		yyXError{479, -1}: "expected identifier",
		yyXError{486, -1}: "expected identifier",
		yyXError{361, -1}: "expected identifier list or identifier",
		yyXError{508, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{462, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{253, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{267, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{255, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{273, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{460, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{368, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{384, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{266, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{179, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{187, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{193, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{229, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{232, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{471, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{472, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{473, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{475, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{482, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{488, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{490, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{493, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{496, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{498, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{500, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{501, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{503, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{505, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{506, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{509, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{515, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{516, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{518, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{523, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{526, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{529, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{530, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{535, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{555, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{556, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{558, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{534, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{538, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{541, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{543, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{548, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{549, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{51, -1}:  "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{452, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{464, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{40, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{41, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{42, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{43, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{337, -1}: "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{380, -1}: "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{382, -1}: "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{465, -1}: "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{36, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{38, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{181, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{189, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{342, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{363, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{373, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{376, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{379, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{386, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{387, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{388, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{389, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{390, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{391, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{392, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{411, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{412, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{413, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{415, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{423, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{429, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{435, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{439, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{443, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{447, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{449, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{450, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{454, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{457, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{459, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{396, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{398, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{399, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{451, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
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
		yyXError{209, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{210, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{213, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{222, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{224, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{230, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{233, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{236, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{237, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
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
		yyXError{46, -1}:  "expected one of ['(', goto]",
		yyXError{343, -1}: "expected one of ['(', goto]",
		yyXError{305, -1}: "expected one of ['(', identifier]",
		yyXError{351, -1}: "expected one of [')', ',', ':']",
		yyXError{358, -1}: "expected one of [')', ',', ':']",
		yyXError{364, -1}: "expected one of [')', ',', ':']",
		yyXError{365, -1}: "expected one of [')', ',', ':']",
		yyXError{369, -1}: "expected one of [')', ',', ':']",
		yyXError{372, -1}: "expected one of [')', ',', ':']",
		yyXError{375, -1}: "expected one of [')', ',', ':']",
		yyXError{385, -1}: "expected one of [')', ',', ';']",
		yyXError{510, -1}: "expected one of [')', ',', ...]",
		yyXError{520, -1}: "expected one of [')', ',', ...]",
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
		yyXError{362, -1}: "expected one of [')', ',']",
		yyXError{378, -1}: "expected one of [')', ',']",
		yyXError{433, -1}: "expected one of [')', ',']",
		yyXError{437, -1}: "expected one of [')', ',']",
		yyXError{441, -1}: "expected one of [')', ',']",
		yyXError{445, -1}: "expected one of [')', ',']",
		yyXError{511, -1}: "expected one of [')', ',']",
		yyXError{48, -1}:  "expected one of [')', ':', string literal]",
		yyXError{50, -1}:  "expected one of [')', ':', string literal]",
		yyXError{370, -1}: "expected one of [')', ':', string literal]",
		yyXError{49, -1}:  "expected one of [')', string literal]",
		yyXError{302, -1}: "expected one of [',', ':', ';']",
		yyXError{131, -1}: "expected one of [',', ':']",
		yyXError{352, -1}: "expected one of [',', ':']",
		yyXError{359, -1}: "expected one of [',', ':']",
		yyXError{335, -1}: "expected one of [',', ';', '=']",
		yyXError{272, -1}: "expected one of [',', ';', '}']",
		yyXError{299, -1}: "expected one of [',', ';']",
		yyXError{301, -1}: "expected one of [',', ';']",
		yyXError{308, -1}: "expected one of [',', ';']",
		yyXError{311, -1}: "expected one of [',', ';']",
		yyXError{325, -1}: "expected one of [',', ';']",
		yyXError{326, -1}: "expected one of [',', ';']",
		yyXError{461, -1}: "expected one of [',', ';']",
		yyXError{463, -1}: "expected one of [',', ';']",
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
		yyXError{347, -1}: "expected one of [':', string literal]",
		yyXError{480, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		yyXError{489, -1}: "expected one of ['\\n', ppother]",
		yyXError{492, -1}: "expected one of ['\\n', ppother]",
		yyXError{494, -1}: "expected one of ['\\n', ppother]",
		yyXError{332, -1}: "expected one of ['{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{334, -1}: "expected one of ['{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{33, -1}:  "expected one of ['{', identifier]",
		yyXError{34, -1}:  "expected one of ['{', identifier]",
		yyXError{297, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{300, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{309, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{313, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{519, -1}: "expected one of [..., identifier]",
		yyXError{169, -1}: "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{84, -1}:  "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{381, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{383, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{57, -1}:  "expected optional comma or one of [',', '}']",
		yyXError{254, -1}: "expected optional comma or one of [',', '}']",
		yyXError{269, -1}: "expected optional comma or one of [',', '}']",
		yyXError{8, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{420, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{426, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{409, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{418, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{424, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{225, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{214, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{178, -1}: "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{182, -1}: "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{532, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{537, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{540, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{547, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{553, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{215, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{32, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{35, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{331, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{199, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{242, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{243, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{167, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{168, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{481, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{485, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{170, -1}: "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{341, -1}: "expected optional volatile or one of ['(', goto, volatile]",
		yyXError{44, -1}:  "expected optional volatile or one of ['(', volatile]",
		yyXError{238, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{216, -1}: "expected parameter type list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{246, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{468, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{507, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{514, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{517, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{522, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{525, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{528, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{166, -1}: "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{404, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{422, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{428, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{438, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{442, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{446, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{448, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{453, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{456, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{458, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{350, -1}: "expected string literal",
		yyXError{355, -1}: "expected string literal",
		yyXError{360, -1}: "expected string literal",
		yyXError{294, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{295, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{296, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{298, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		yyXError{310, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{497, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{474, -1}: "expected token list or ppother",
		yyXError{477, -1}: "expected token list or ppother",
		yyXError{483, -1}: "expected token list or ppother",
		yyXError{484, -1}: "expected token list or ppother",
		yyXError{487, -1}: "expected token list or ppother",
		yyXError{545, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [_Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{5, -1}:   "expected translation unit or one of [_Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{185, -1}: "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{227, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{430, -1}: "expected while",
		yyXError{3, 48}:   "unexpected EOF",
		yyXError{2, 48}:   "unexpected EOF",
		yyXError{4, 48}:   "unexpected EOF",
	}

	yyParseTab = [559][]uint16{
		// 0
		{211: 315, 222: 314, 224: 313, 227: 316},
		{48: 312},
		{81: 311, 311, 97: 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 194: 780},
		{309, 309, 309, 309, 309, 309, 309, 309, 11: 309, 13: 309, 309, 309, 309, 309, 309, 309, 309, 201: 778},
		{21: 307, 307, 307, 26: 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 54: 307, 202: 317},
		// 5
		{21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 54: 356, 110: 322, 114: 341, 344, 340, 321, 120: 355, 319, 323, 320, 140: 354, 182: 352, 353, 228: 318},
		{21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 306, 54: 356, 110: 322, 114: 341, 344, 340, 321, 120: 355, 319, 323, 320, 140: 354, 182: 777, 353},
		{156, 482, 156, 9: 227, 132: 617, 616, 135: 639, 165: 637, 186: 638, 636},
		{231, 231, 231, 8: 231, 231, 231, 12: 231, 21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 110: 322, 114: 341, 344, 340, 321, 121: 632, 323, 320, 152: 635},
		{231, 231, 231, 8: 231, 231, 231, 12: 231, 21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 110: 322, 114: 341, 344, 340, 321, 121: 632, 323, 320, 152: 634},
		// 10
		{231, 231, 231, 8: 231, 231, 231, 12: 231, 21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 110: 322, 114: 341, 344, 340, 321, 121: 632, 323, 320, 152: 633},
		{231, 231, 231, 8: 231, 231, 231, 12: 231, 21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 110: 322, 114: 341, 344, 340, 321, 121: 632, 323, 320, 152: 631},
		{222, 222, 222, 8: 222, 222, 222, 12: 222, 21: 222, 222, 222, 26: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222},
		{221, 221, 221, 8: 221, 221, 221, 12: 221, 21: 221, 221, 221, 26: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221},
		{220, 220, 220, 8: 220, 220, 220, 12: 220, 21: 220, 220, 220, 26: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220},
		// 15
		{219, 219, 219, 8: 219, 219, 219, 12: 219, 21: 219, 219, 219, 26: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219},
		{218, 218, 218, 8: 218, 218, 218, 12: 218, 21: 218, 218, 218, 26: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218},
		{217, 217, 217, 8: 217, 217, 217, 12: 217, 21: 217, 217, 217, 217, 26: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217},
		{216, 216, 216, 8: 216, 216, 216, 12: 216, 21: 216, 216, 216, 216, 26: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216},
		{215, 215, 215, 8: 215, 215, 215, 12: 215, 21: 215, 215, 215, 215, 26: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215},
		// 20
		{214, 214, 214, 8: 214, 214, 214, 12: 214, 21: 214, 214, 214, 214, 26: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214},
		{213, 213, 213, 8: 213, 213, 213, 12: 213, 21: 213, 213, 213, 213, 26: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213},
		{212, 212, 212, 8: 212, 212, 212, 12: 212, 21: 212, 212, 212, 212, 26: 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212},
		{211, 211, 211, 8: 211, 211, 211, 12: 211, 21: 211, 211, 211, 211, 26: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211},
		{210, 210, 210, 8: 210, 210, 210, 12: 210, 21: 210, 210, 210, 210, 26: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210},
		// 25
		{209, 209, 209, 8: 209, 209, 209, 12: 209, 21: 209, 209, 209, 209, 26: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209},
		{208, 208, 208, 8: 208, 208, 208, 12: 208, 21: 208, 208, 208, 208, 26: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208},
		{207, 207, 207, 8: 207, 207, 207, 12: 207, 21: 207, 207, 207, 207, 26: 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207},
		{206, 206, 206, 8: 206, 206, 206, 12: 206, 21: 206, 206, 206, 206, 26: 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206},
		{205, 205, 205, 8: 205, 205, 205, 12: 205, 21: 205, 205, 205, 205, 26: 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205},
		// 30
		{204, 204, 204, 8: 204, 204, 204, 12: 204, 21: 204, 204, 204, 204, 26: 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204},
		{626},
		{2: 604, 51: 138, 185: 603},
		{2: 197, 51: 197},
		{2: 196, 51: 196},
		// 35
		{2: 365, 51: 138, 185: 364},
		{174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 26: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 52: 174},
		{173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 26: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 52: 173},
		{172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 26: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 52: 172},
		{171, 171, 171, 8: 171, 171, 171, 12: 171, 21: 171, 171, 171, 26: 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171},
		// 40
		{21: 77, 77, 77, 26: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 54: 77},
		{21: 75, 75, 75, 26: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 54: 75},
		{21: 74, 74, 74, 26: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 54: 74},
		{21: 73, 73, 73, 26: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 54: 73},
		{58, 21: 358, 193: 357},
		// 45
		{359},
		{57, 83: 57},
		{11: 360, 157: 361},
		{8: 61, 11: 61, 24: 61},
		{8: 363, 11: 362},
		// 50
		{8: 60, 11: 60, 24: 60},
		{59, 59, 59, 59, 59, 59, 59, 59, 9: 59, 11: 59, 13: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 25: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 51: 59, 54: 59, 83: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 109: 59},
		{51: 181, 205: 366},
		{179, 179, 179, 8: 179, 179, 179, 12: 179, 21: 179, 179, 179, 179, 26: 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 51: 137},
		{51: 367},
		// 55
		{2: 368, 180: 371, 370, 219: 369},
		{10: 305, 25: 305, 49: 305},
		{10: 599, 25: 183, 159: 600},
		{10: 178, 25: 178},
		{10: 176, 25: 176, 49: 372},
		// 60
		{238, 238, 238, 238, 238, 238, 238, 238, 11: 238, 13: 238, 238, 238, 238, 238, 238, 238, 238, 143: 373, 374},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 383},
		{10: 175, 25: 175},
		{300, 300, 3: 300, 300, 300, 300, 300, 300, 300, 300, 12: 300, 24: 300, 300, 48: 300, 300, 52: 300, 300, 55: 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300},
		{299, 299, 3: 299, 299, 299, 299, 299, 299, 299, 299, 12: 299, 24: 299, 299, 48: 299, 299, 52: 299, 299, 55: 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299},
		// 65
		{298, 298, 3: 298, 298, 298, 298, 298, 298, 298, 298, 12: 298, 24: 298, 298, 48: 298, 298, 52: 298, 298, 55: 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298},
		{297, 297, 3: 297, 297, 297, 297, 297, 297, 297, 297, 12: 297, 24: 297, 297, 48: 297, 297, 52: 297, 297, 55: 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297},
		{296, 296, 3: 296, 296, 296, 296, 296, 296, 296, 296, 12: 296, 24: 296, 296, 48: 296, 296, 52: 296, 296, 55: 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296},
		{295, 295, 3: 295, 295, 295, 295, 295, 295, 295, 295, 12: 295, 24: 295, 295, 48: 295, 295, 52: 295, 295, 55: 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295},
		{294, 294, 3: 294, 294, 294, 294, 294, 294, 294, 294, 12: 294, 24: 294, 294, 48: 294, 294, 52: 294, 294, 55: 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294},
		// 70
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 136, 136, 136, 26: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 50: 442, 111: 476, 156: 478, 169: 597},
		{396, 401, 3: 414, 404, 405, 400, 399, 9: 237, 237, 12: 395, 24: 237, 237, 48: 237, 420, 52: 237, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 596},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 595},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 594},
		// 75
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 508},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 593},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 592},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 591},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 590},
		// 80
		{393, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 394},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 136, 136, 136, 26: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 50: 442, 111: 476, 156: 478, 169: 477},
		{396, 277, 3: 277, 277, 277, 400, 399, 277, 277, 277, 12: 395, 24: 277, 277, 48: 277, 277, 52: 277, 397, 55: 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 398, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 111: 474},
		{382, 387, 375, 386, 388, 389, 385, 384, 302, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 468, 207: 469, 470},
		// 85
		{2: 467},
		{2: 466},
		{288, 288, 3: 288, 288, 288, 288, 288, 288, 288, 288, 12: 288, 24: 288, 288, 48: 288, 288, 52: 288, 288, 55: 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288},
		{287, 287, 3: 287, 287, 287, 287, 287, 287, 287, 287, 12: 287, 24: 287, 287, 48: 287, 287, 52: 287, 287, 55: 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 465},
		// 90
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 464},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 463},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 462},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 461},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 460},
		// 95
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 459},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 458},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 457},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 456},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 455},
		// 100
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 454},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 453},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 452},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 451},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 450},
		// 105
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 449},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 448},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 111: 443},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 441},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 440},
		// 110
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 439},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 438},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 437},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 436},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 435},
		// 115
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 434},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 433},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 432},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 431},
		{396, 401, 3: 414, 404, 405, 400, 399, 245, 245, 245, 12: 395, 24: 245, 245, 48: 245, 420, 52: 245, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		// 120
		{396, 401, 3: 414, 404, 405, 400, 399, 246, 246, 246, 12: 395, 24: 246, 246, 48: 246, 420, 52: 246, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{396, 401, 3: 414, 404, 405, 400, 399, 247, 247, 247, 12: 395, 24: 247, 247, 48: 247, 420, 52: 247, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{396, 401, 3: 414, 404, 405, 400, 399, 248, 248, 248, 12: 395, 24: 248, 248, 48: 248, 420, 52: 248, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{396, 401, 3: 414, 404, 405, 400, 399, 249, 249, 249, 12: 395, 24: 249, 249, 48: 249, 420, 52: 249, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{396, 401, 3: 414, 404, 405, 400, 399, 250, 250, 250, 12: 395, 24: 250, 250, 48: 250, 420, 52: 250, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		// 125
		{396, 401, 3: 414, 404, 405, 400, 399, 251, 251, 251, 12: 395, 24: 251, 251, 48: 251, 420, 52: 251, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{396, 401, 3: 414, 404, 405, 400, 399, 252, 252, 252, 12: 395, 24: 252, 252, 48: 252, 420, 52: 252, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{396, 401, 3: 414, 404, 405, 400, 399, 253, 253, 253, 12: 395, 24: 253, 253, 48: 253, 420, 52: 253, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{396, 401, 3: 414, 404, 405, 400, 399, 254, 254, 254, 12: 395, 24: 254, 254, 48: 254, 420, 52: 254, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{396, 401, 3: 414, 404, 405, 400, 399, 255, 255, 255, 12: 395, 24: 255, 255, 48: 255, 420, 52: 255, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		// 130
		{396, 401, 3: 414, 404, 405, 400, 399, 242, 242, 242, 12: 395, 24: 242, 49: 420, 52: 242, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{10: 445, 24: 444},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 447},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 446},
		{396, 401, 3: 414, 404, 405, 400, 399, 241, 241, 241, 12: 395, 24: 241, 49: 420, 52: 241, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		// 135
		{396, 401, 3: 414, 404, 405, 400, 399, 256, 256, 256, 12: 395, 24: 256, 256, 48: 256, 256, 52: 256, 397, 55: 403, 402, 408, 409, 419, 415, 416, 256, 417, 256, 398, 256, 412, 411, 410, 406, 256, 256, 256, 413, 256, 418, 407, 256, 256, 256},
		{396, 401, 3: 414, 404, 405, 400, 399, 257, 257, 257, 12: 395, 24: 257, 257, 48: 257, 257, 52: 257, 397, 55: 403, 402, 408, 409, 257, 415, 416, 257, 417, 257, 398, 257, 412, 411, 410, 406, 257, 257, 257, 413, 257, 257, 407, 257, 257, 257},
		{396, 401, 3: 414, 404, 405, 400, 399, 258, 258, 258, 12: 395, 24: 258, 258, 48: 258, 258, 52: 258, 397, 55: 403, 402, 408, 409, 258, 415, 416, 258, 258, 258, 398, 258, 412, 411, 410, 406, 258, 258, 258, 413, 258, 258, 407, 258, 258, 258},
		{396, 401, 3: 414, 404, 405, 400, 399, 259, 259, 259, 12: 395, 24: 259, 259, 48: 259, 259, 52: 259, 397, 55: 403, 402, 408, 409, 259, 415, 259, 259, 259, 259, 398, 259, 412, 411, 410, 406, 259, 259, 259, 413, 259, 259, 407, 259, 259, 259},
		{396, 401, 3: 414, 404, 405, 400, 399, 260, 260, 260, 12: 395, 24: 260, 260, 48: 260, 260, 52: 260, 397, 55: 403, 402, 408, 409, 260, 260, 260, 260, 260, 260, 398, 260, 412, 411, 410, 406, 260, 260, 260, 413, 260, 260, 407, 260, 260, 260},
		// 140
		{396, 401, 3: 261, 404, 405, 400, 399, 261, 261, 261, 12: 395, 24: 261, 261, 48: 261, 261, 52: 261, 397, 55: 403, 402, 408, 409, 261, 261, 261, 261, 261, 261, 398, 261, 412, 411, 410, 406, 261, 261, 261, 413, 261, 261, 407, 261, 261, 261},
		{396, 401, 3: 262, 404, 405, 400, 399, 262, 262, 262, 12: 395, 24: 262, 262, 48: 262, 262, 52: 262, 397, 55: 403, 402, 408, 409, 262, 262, 262, 262, 262, 262, 398, 262, 262, 411, 410, 406, 262, 262, 262, 262, 262, 262, 407, 262, 262, 262},
		{396, 401, 3: 263, 404, 405, 400, 399, 263, 263, 263, 12: 395, 24: 263, 263, 48: 263, 263, 52: 263, 397, 55: 403, 402, 408, 409, 263, 263, 263, 263, 263, 263, 398, 263, 263, 411, 410, 406, 263, 263, 263, 263, 263, 263, 407, 263, 263, 263},
		{396, 401, 3: 264, 404, 405, 400, 399, 264, 264, 264, 12: 395, 24: 264, 264, 48: 264, 264, 52: 264, 397, 55: 403, 402, 264, 264, 264, 264, 264, 264, 264, 264, 398, 264, 264, 264, 264, 406, 264, 264, 264, 264, 264, 264, 407, 264, 264, 264},
		{396, 401, 3: 265, 404, 405, 400, 399, 265, 265, 265, 12: 395, 24: 265, 265, 48: 265, 265, 52: 265, 397, 55: 403, 402, 265, 265, 265, 265, 265, 265, 265, 265, 398, 265, 265, 265, 265, 406, 265, 265, 265, 265, 265, 265, 407, 265, 265, 265},
		// 145
		{396, 401, 3: 266, 404, 405, 400, 399, 266, 266, 266, 12: 395, 24: 266, 266, 48: 266, 266, 52: 266, 397, 55: 403, 402, 266, 266, 266, 266, 266, 266, 266, 266, 398, 266, 266, 266, 266, 406, 266, 266, 266, 266, 266, 266, 407, 266, 266, 266},
		{396, 401, 3: 267, 404, 405, 400, 399, 267, 267, 267, 12: 395, 24: 267, 267, 48: 267, 267, 52: 267, 397, 55: 403, 402, 267, 267, 267, 267, 267, 267, 267, 267, 398, 267, 267, 267, 267, 406, 267, 267, 267, 267, 267, 267, 407, 267, 267, 267},
		{396, 401, 3: 268, 404, 405, 400, 399, 268, 268, 268, 12: 395, 24: 268, 268, 48: 268, 268, 52: 268, 397, 55: 403, 402, 268, 268, 268, 268, 268, 268, 268, 268, 398, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268},
		{396, 401, 3: 269, 404, 405, 400, 399, 269, 269, 269, 12: 395, 24: 269, 269, 48: 269, 269, 52: 269, 397, 55: 403, 402, 269, 269, 269, 269, 269, 269, 269, 269, 398, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269},
		{396, 401, 3: 270, 270, 270, 400, 399, 270, 270, 270, 12: 395, 24: 270, 270, 48: 270, 270, 52: 270, 397, 55: 403, 402, 270, 270, 270, 270, 270, 270, 270, 270, 398, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270},
		// 150
		{396, 401, 3: 271, 271, 271, 400, 399, 271, 271, 271, 12: 395, 24: 271, 271, 48: 271, 271, 52: 271, 397, 55: 403, 402, 271, 271, 271, 271, 271, 271, 271, 271, 398, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{396, 272, 3: 272, 272, 272, 400, 399, 272, 272, 272, 12: 395, 24: 272, 272, 48: 272, 272, 52: 272, 397, 55: 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 398, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272},
		{396, 273, 3: 273, 273, 273, 400, 399, 273, 273, 273, 12: 395, 24: 273, 273, 48: 273, 273, 52: 273, 397, 55: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 398, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273},
		{396, 274, 3: 274, 274, 274, 400, 399, 274, 274, 274, 12: 395, 24: 274, 274, 48: 274, 274, 52: 274, 397, 55: 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 398, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274},
		{289, 289, 3: 289, 289, 289, 289, 289, 289, 289, 289, 12: 289, 24: 289, 289, 48: 289, 289, 52: 289, 289, 55: 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289},
		// 155
		{290, 290, 3: 290, 290, 290, 290, 290, 290, 290, 290, 12: 290, 24: 290, 290, 48: 290, 290, 52: 290, 290, 55: 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290},
		{396, 401, 3: 414, 404, 405, 400, 399, 304, 10: 304, 12: 395, 49: 420, 53: 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{8: 301, 10: 472},
		{8: 471},
		{291, 291, 3: 291, 291, 291, 291, 291, 291, 291, 291, 12: 291, 24: 291, 291, 48: 291, 291, 52: 291, 291, 55: 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291},
		// 160
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 473},
		{396, 401, 3: 414, 404, 405, 400, 399, 303, 10: 303, 12: 395, 49: 420, 53: 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{10: 445, 52: 475},
		{292, 292, 3: 292, 292, 292, 292, 292, 292, 292, 292, 12: 292, 24: 292, 292, 48: 292, 292, 52: 292, 292, 55: 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292},
		{8: 589, 10: 445},
		// 165
		{8: 563},
		{21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 110: 480, 114: 341, 344, 340, 479, 147: 481},
		{189, 189, 189, 8: 189, 189, 12: 189, 21: 350, 348, 349, 189, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 110: 480, 114: 341, 344, 340, 479, 147: 561, 190: 562},
		{189, 189, 189, 8: 189, 189, 12: 189, 21: 350, 348, 349, 189, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 110: 480, 114: 341, 344, 340, 479, 147: 561, 190: 560},
		{156, 482, 8: 132, 12: 156, 132: 483, 485, 149: 486, 172: 484},
		// 170
		{152, 152, 152, 8: 152, 10: 152, 12: 152, 21: 350, 348, 349, 110: 493, 148: 497, 155: 558},
		{155, 2: 155, 8: 134, 10: 134, 12: 155},
		{8: 135},
		{488, 12: 120, 176: 487, 489},
		{8: 131, 10: 131},
		// 175
		{554, 8: 133, 10: 133, 12: 119},
		{156, 482, 8: 124, 12: 156, 21: 124, 124, 124, 26: 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 132: 483, 485, 149: 510, 170: 511},
		{12: 490},
		{382, 492, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 350, 348, 349, 42: 496, 50: 491, 52: 244, 110: 493, 148: 494, 163: 495},
		{396, 401, 3: 414, 404, 405, 400, 399, 12: 395, 49: 420, 52: 243, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		// 180
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 508, 52: 509},
		{154, 154, 154, 154, 154, 154, 154, 154, 154, 10: 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 42: 154, 52: 154},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 350, 348, 349, 42: 504, 50: 491, 52: 244, 110: 501, 163: 503},
		{52: 502},
		{152, 152, 152, 152, 152, 152, 152, 152, 11: 152, 13: 152, 152, 152, 152, 152, 152, 152, 152, 350, 348, 349, 110: 493, 148: 497, 155: 498},
		// 185
		{151, 151, 151, 151, 151, 151, 151, 151, 151, 10: 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 350, 348, 349, 110: 501},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 499},
		{396, 401, 3: 414, 404, 405, 400, 399, 12: 395, 49: 420, 52: 500, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{127, 8: 127, 10: 127, 12: 127},
		{153, 153, 153, 153, 153, 153, 153, 153, 153, 10: 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 42: 153, 52: 153},
		// 190
		{129, 8: 129, 10: 129, 12: 129},
		{52: 507},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 505},
		{396, 401, 3: 414, 404, 405, 400, 399, 12: 395, 49: 420, 52: 506, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{126, 8: 126, 10: 126, 12: 126},
		// 195
		{128, 8: 128, 10: 128, 12: 128},
		{396, 282, 3: 282, 282, 282, 400, 399, 282, 282, 282, 12: 395, 24: 282, 282, 48: 282, 282, 52: 282, 397, 55: 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 398, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282},
		{125, 8: 125, 10: 125, 12: 125},
		{8: 553},
		{8: 148, 21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 110: 322, 114: 341, 344, 340, 321, 121: 515, 323, 320, 154: 514, 167: 512, 513, 189: 516},
		// 200
		{8: 150, 10: 550},
		{8: 147},
		{8: 146, 10: 146},
		{156, 482, 156, 8: 132, 10: 132, 12: 156, 132: 483, 518, 135: 519, 149: 486, 172: 520},
		{8: 517},
		// 205
		{123, 8: 123, 10: 123, 12: 123},
		{523, 2: 522, 12: 120, 176: 487, 489, 521},
		{8: 144, 10: 144},
		{8: 143, 10: 143},
		{527, 8: 170, 170, 170, 12: 526, 21: 170, 170, 170, 170, 26: 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 49: 170, 51: 170, 54: 170},
		// 210
		{167, 8: 167, 167, 167, 12: 167, 21: 167, 167, 167, 167, 26: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 49: 167, 51: 167, 54: 167},
		{156, 482, 156, 8: 124, 12: 156, 21: 124, 124, 124, 26: 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 132: 483, 518, 135: 524, 149: 510, 170: 511},
		{8: 525},
		{166, 8: 166, 166, 166, 12: 166, 21: 166, 166, 166, 166, 26: 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 49: 166, 51: 166, 54: 166},
		{152, 152, 152, 152, 152, 152, 152, 152, 11: 152, 13: 152, 152, 152, 152, 152, 152, 152, 152, 350, 348, 349, 42: 538, 52: 152, 110: 493, 148: 539, 155: 537},
		// 215
		{2: 530, 8: 140, 21: 161, 161, 161, 26: 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 164: 531, 184: 529, 206: 528},
		{21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 110: 322, 114: 341, 344, 340, 321, 121: 515, 323, 320, 154: 514, 167: 512, 535},
		{8: 534},
		{8: 142, 10: 142},
		{8: 139, 10: 532},
		// 220
		{2: 533},
		{8: 141, 10: 141},
		{159, 8: 159, 159, 159, 12: 159, 21: 159, 159, 159, 159, 26: 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 49: 159, 51: 159, 54: 159},
		{8: 536},
		{160, 8: 160, 160, 160, 12: 160, 21: 160, 160, 160, 160, 26: 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 49: 160, 51: 160, 54: 160},
		// 225
		{382, 546, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 491, 52: 244, 163: 547},
		{152, 152, 152, 152, 152, 152, 152, 152, 11: 152, 13: 152, 152, 152, 152, 152, 152, 152, 152, 350, 348, 349, 110: 493, 148: 497, 155: 543},
		{151, 151, 151, 151, 151, 151, 151, 151, 11: 151, 13: 151, 151, 151, 151, 151, 151, 151, 151, 350, 348, 349, 42: 540, 52: 151, 110: 501},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 541},
		{396, 401, 3: 414, 404, 405, 400, 399, 12: 395, 49: 420, 52: 542, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		// 230
		{163, 8: 163, 163, 163, 12: 163, 21: 163, 163, 163, 163, 26: 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 49: 163, 51: 163, 54: 163},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 544},
		{396, 401, 3: 414, 404, 405, 400, 399, 12: 395, 49: 420, 52: 545, 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{164, 8: 164, 164, 164, 12: 164, 21: 164, 164, 164, 164, 26: 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 49: 164, 51: 164, 54: 164},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 508, 52: 549},
		// 235
		{52: 548},
		{165, 8: 165, 165, 165, 12: 165, 21: 165, 165, 165, 165, 26: 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 49: 165, 51: 165, 54: 165},
		{162, 8: 162, 162, 162, 12: 162, 21: 162, 162, 162, 162, 26: 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 49: 162, 51: 162, 54: 162},
		{21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 110: 322, 114: 341, 344, 340, 321, 121: 515, 323, 320, 145: 551, 154: 552},
		{8: 149},
		// 240
		{8: 145, 10: 145},
		{130, 8: 130, 10: 130, 12: 130},
		{8: 122, 21: 122, 122, 122, 26: 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 195: 555},
		{8: 148, 21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 110: 322, 114: 341, 344, 340, 321, 121: 515, 323, 320, 154: 514, 167: 512, 513, 189: 556},
		{8: 557},
		// 245
		{121, 8: 121, 10: 121, 12: 121},
		{158, 482, 158, 8: 158, 10: 158, 12: 158, 132: 559},
		{157, 2: 157, 8: 157, 10: 157, 12: 157},
		{190, 190, 190, 8: 190, 190, 12: 190, 24: 190},
		{188, 188, 188, 8: 188, 188, 12: 188, 24: 188},
		// 250
		{191, 191, 191, 8: 191, 191, 12: 191, 24: 191},
		{382, 276, 375, 276, 276, 276, 385, 384, 276, 276, 276, 381, 276, 391, 390, 376, 377, 378, 379, 380, 392, 24: 276, 276, 48: 276, 276, 564, 565, 276, 276, 55: 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276},
		{396, 275, 3: 275, 275, 275, 400, 399, 275, 275, 275, 12: 395, 24: 275, 275, 48: 275, 275, 52: 275, 397, 55: 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 398, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275},
		{113, 113, 113, 113, 113, 113, 113, 113, 11: 113, 571, 113, 113, 113, 113, 113, 113, 113, 113, 51: 113, 53: 572, 153: 570, 160: 569, 567, 568, 188: 566},
		{10: 582, 25: 183, 159: 587},
		// 255
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 578, 579, 166: 580},
		{12: 571, 49: 576, 53: 572, 153: 577},
		{112, 112, 112, 112, 112, 112, 112, 112, 11: 112, 13: 112, 112, 112, 112, 112, 112, 112, 112, 51: 112},
		{12: 111, 49: 111, 53: 111},
		{238, 238, 238, 238, 238, 238, 238, 238, 11: 238, 13: 238, 238, 238, 238, 238, 238, 238, 238, 143: 373, 574},
		// 260
		{2: 573},
		{12: 108, 49: 108, 53: 108},
		{52: 575},
		{12: 109, 49: 109, 53: 109},
		{114, 114, 114, 114, 114, 114, 114, 114, 11: 114, 13: 114, 114, 114, 114, 114, 114, 114, 114, 51: 114},
		// 265
		{12: 110, 49: 110, 53: 110},
		{396, 401, 3: 414, 404, 405, 400, 399, 9: 118, 118, 12: 395, 25: 118, 49: 420, 53: 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{113, 113, 113, 113, 113, 113, 113, 113, 11: 113, 571, 113, 113, 113, 113, 113, 113, 113, 113, 51: 113, 53: 572, 153: 570, 160: 569, 567, 568, 188: 581},
		{10: 116, 25: 116},
		{10: 582, 25: 183, 159: 583},
		// 270
		{113, 113, 113, 113, 113, 113, 113, 113, 11: 113, 571, 113, 113, 113, 113, 113, 113, 113, 113, 25: 182, 51: 113, 53: 572, 153: 570, 160: 569, 585, 568},
		{25: 584},
		{9: 117, 117, 25: 117},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 578, 579, 166: 586},
		{10: 115, 25: 115},
		// 275
		{25: 588},
		{286, 286, 3: 286, 286, 286, 286, 286, 286, 286, 286, 12: 286, 24: 286, 286, 48: 286, 286, 52: 286, 286, 55: 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286},
		{293, 293, 3: 293, 293, 293, 293, 293, 293, 293, 293, 12: 293, 24: 293, 293, 48: 293, 293, 52: 293, 293, 55: 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293},
		{396, 278, 3: 278, 278, 278, 400, 399, 278, 278, 278, 12: 395, 24: 278, 278, 48: 278, 278, 52: 278, 397, 55: 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 398, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278},
		{396, 279, 3: 279, 279, 279, 400, 399, 279, 279, 279, 12: 395, 24: 279, 279, 48: 279, 279, 52: 279, 397, 55: 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 398, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279},
		// 280
		{396, 280, 3: 280, 280, 280, 400, 399, 280, 280, 280, 12: 395, 24: 280, 280, 48: 280, 280, 52: 280, 397, 55: 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 398, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280},
		{396, 281, 3: 281, 281, 281, 400, 399, 281, 281, 281, 12: 395, 24: 281, 281, 48: 281, 281, 52: 281, 397, 55: 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 398, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281},
		{396, 283, 3: 283, 283, 283, 400, 399, 283, 283, 283, 12: 395, 24: 283, 283, 48: 283, 283, 52: 283, 397, 55: 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 398, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283},
		{396, 284, 3: 284, 284, 284, 400, 399, 284, 284, 284, 12: 395, 24: 284, 284, 48: 284, 284, 52: 284, 397, 55: 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 398, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284},
		{396, 285, 3: 285, 285, 285, 400, 399, 285, 285, 285, 12: 395, 24: 285, 285, 48: 285, 285, 52: 285, 397, 55: 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 398, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285},
		// 285
		{8: 598},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 564, 565},
		{2: 368, 25: 182, 180: 371, 602},
		{25: 601},
		{180, 180, 180, 8: 180, 180, 180, 12: 180, 21: 180, 180, 180, 180, 26: 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180},
		// 290
		{10: 177, 25: 177},
		{51: 201, 203: 605},
		{198, 198, 198, 8: 198, 198, 198, 12: 198, 21: 198, 198, 198, 198, 26: 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 51: 137},
		{51: 606},
		{21: 200, 200, 200, 26: 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 204: 607},
		// 295
		{21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 110: 480, 114: 341, 344, 340, 479, 147: 610, 191: 609, 225: 608},
		{21: 350, 348, 349, 25: 624, 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 110: 480, 114: 341, 344, 340, 479, 147: 610, 191: 625},
		{21: 195, 195, 195, 25: 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195},
		{156, 482, 156, 9: 612, 24: 169, 132: 617, 616, 135: 614, 175: 615, 192: 613, 226: 611},
		{9: 621, 622},
		// 300
		{21: 192, 192, 192, 25: 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192},
		{9: 187, 187},
		{9: 185, 185, 24: 168},
		{24: 619},
		{618, 2: 522, 178: 521},
		// 305
		{155, 2: 155},
		{156, 482, 156, 132: 617, 616, 135: 524},
		{238, 238, 238, 238, 238, 238, 238, 238, 11: 238, 13: 238, 238, 238, 238, 238, 238, 238, 238, 143: 373, 620},
		{9: 184, 184},
		{21: 193, 193, 193, 25: 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193},
		// 310
		{156, 482, 156, 24: 169, 132: 617, 616, 135: 614, 175: 615, 192: 623},
		{9: 186, 186},
		{199, 199, 199, 8: 199, 199, 199, 12: 199, 21: 199, 199, 199, 199, 26: 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199},
		{21: 194, 194, 194, 25: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 136, 136, 136, 26: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 50: 627, 156: 478, 169: 628},
		// 315
		{396, 401, 3: 414, 404, 405, 400, 399, 630, 12: 395, 49: 420, 53: 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{8: 629},
		{202, 202, 202, 8: 202, 202, 202, 12: 202, 21: 202, 202, 202, 202, 26: 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202},
		{203, 203, 203, 8: 203, 203, 203, 12: 203, 21: 203, 203, 203, 203, 26: 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203},
		{232, 232, 232, 8: 232, 232, 232, 12: 232},
		// 320
		{230, 230, 230, 8: 230, 230, 230, 12: 230},
		{233, 233, 233, 8: 233, 233, 233, 12: 233},
		{234, 234, 234, 8: 234, 234, 234, 12: 234},
		{235, 235, 235, 8: 235, 235, 235, 12: 235},
		{9: 776},
		// 325
		{9: 229, 229},
		{9: 226, 774},
		{9: 225, 225, 21: 63, 63, 63, 26: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 49: 224, 51: 64, 54: 64, 171: 640, 200: 642, 213: 641},
		{49: 772},
		{51: 72, 54: 72, 197: 648},
		// 330
		{21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 110: 322, 114: 341, 344, 340, 321, 121: 643, 323, 320, 140: 644, 212: 645},
		{156, 482, 156, 9: 227, 132: 617, 616, 135: 647, 165: 637, 186: 638, 636},
		{21: 66, 66, 66, 26: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 51: 66, 54: 66},
		{21: 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 51: 62, 54: 62, 110: 322, 114: 341, 344, 340, 321, 121: 643, 323, 320, 140: 646},
		{21: 65, 65, 65, 26: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 51: 65, 54: 65},
		// 335
		{9: 225, 225, 49: 224, 171: 640},
		{51: 70, 54: 68, 198: 650, 651, 220: 649},
		{21: 71, 71, 71, 26: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 54: 71},
		{51: 693, 125: 694},
		{54: 653, 120: 654, 124: 652},
		// 340
		{9: 692},
		{58, 21: 358, 83: 58, 193: 655},
		{49, 49, 49, 49, 49, 49, 49, 49, 9: 49, 11: 49, 13: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 25: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 51: 49, 54: 49, 83: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 109: 49},
		{656, 83: 657},
		{11: 360, 157: 682},
		// 345
		{658},
		{11: 360, 157: 659},
		{11: 362, 24: 660},
		{24: 661},
		{11: 53, 665, 150: 663, 662, 158: 664},
		// 350
		{11: 678},
		{8: 55, 10: 55, 24: 55},
		{10: 668, 24: 669},
		{2: 666},
		{52: 667},
		// 355
		{11: 52},
		{11: 53, 665, 150: 677, 662},
		{11: 670, 174: 671},
		{8: 51, 10: 51, 24: 51},
		{10: 672, 24: 673},
		// 360
		{11: 676},
		{2: 530, 164: 674},
		{8: 675, 10: 532},
		{45, 45, 45, 45, 45, 45, 45, 45, 9: 45, 11: 45, 13: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 25: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 51: 45, 54: 45, 83: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 109: 45},
		{8: 50, 10: 50, 24: 50},
		// 365
		{8: 54, 10: 54, 24: 54},
		{679},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 680},
		{396, 401, 3: 414, 404, 405, 400, 399, 681, 12: 395, 49: 420, 53: 397, 55: 403, 402, 408, 409, 419, 415, 416, 424, 417, 428, 398, 422, 412, 411, 410, 406, 426, 423, 421, 413, 430, 418, 407, 427, 425, 429},
		{8: 56, 10: 56, 24: 56},
		// 370
		{8: 363, 11: 362, 24: 683},
		{11: 53, 665, 150: 663, 662, 158: 684},
		{8: 685, 10: 668, 24: 686},
		{48, 48, 48, 48, 48, 48, 48, 48, 9: 48, 11: 48, 13: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 25: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 51: 48, 54: 48, 83: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 109: 48},
		{11: 53, 665, 150: 663, 662, 158: 687},
		// 375
		{8: 688, 10: 668, 24: 689},
		{47, 47, 47, 47, 47, 47, 47, 47, 9: 47, 11: 47, 13: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 25: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 51: 47, 54: 47, 83: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 109: 47},
		{11: 670, 174: 690},
		{8: 691, 10: 672},
		{46, 46, 46, 46, 46, 46, 46, 46, 9: 46, 11: 46, 13: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 25: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 51: 46, 54: 46, 83: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 109: 46},
		// 380
		{21: 67, 67, 67, 26: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 54: 67},
		{97, 97, 97, 97, 97, 97, 97, 97, 9: 97, 11: 97, 13: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 25: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 51: 97, 54: 97, 83: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 196: 695},
		{21: 69, 69, 69, 26: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 54: 69},
		{382, 387, 696, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 350, 348, 349, 25: 93, 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 50: 442, 693, 54: 653, 83: 718, 715, 720, 705, 719, 706, 716, 717, 713, 721, 714, 110: 322, 697, 114: 341, 344, 340, 321, 712, 120: 654, 643, 323, 320, 704, 699, 700, 702, 703, 698, 701, 711, 140: 710, 173: 708, 209: 709, 707},
		{300, 300, 3: 300, 300, 300, 300, 300, 9: 300, 300, 12: 300, 24: 770, 49: 300, 53: 300, 55: 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300},
		// 385
		{8: 239, 239, 445},
		{107, 107, 107, 107, 107, 107, 107, 107, 9: 107, 11: 107, 13: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 25: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 51: 107, 54: 107, 83: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 109: 107},
		{106, 106, 106, 106, 106, 106, 106, 106, 9: 106, 11: 106, 13: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 25: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 51: 106, 54: 106, 83: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 109: 106},
		{105, 105, 105, 105, 105, 105, 105, 105, 9: 105, 11: 105, 13: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 25: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 51: 105, 54: 105, 83: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 109: 105},
		{104, 104, 104, 104, 104, 104, 104, 104, 9: 104, 11: 104, 13: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 25: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 51: 104, 54: 104, 83: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 109: 104},
		// 390
		{103, 103, 103, 103, 103, 103, 103, 103, 9: 103, 11: 103, 13: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 25: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 51: 103, 54: 103, 83: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 109: 103},
		{102, 102, 102, 102, 102, 102, 102, 102, 9: 102, 11: 102, 13: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 25: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 51: 102, 54: 102, 83: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 109: 102},
		{101, 101, 101, 101, 101, 101, 101, 101, 9: 101, 11: 101, 13: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 25: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 51: 101, 54: 101, 83: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 109: 101},
		{238, 238, 238, 238, 238, 238, 238, 238, 11: 238, 13: 238, 238, 238, 238, 238, 238, 238, 238, 143: 373, 767},
		{24: 765},
		// 395
		{25: 764},
		{95, 95, 95, 95, 95, 95, 95, 95, 9: 95, 11: 95, 13: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 25: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 51: 95, 54: 95, 83: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95},
		{382, 387, 696, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 350, 348, 349, 25: 92, 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 50: 442, 693, 54: 653, 83: 718, 715, 720, 705, 719, 706, 716, 717, 713, 721, 714, 110: 322, 697, 114: 341, 344, 340, 321, 712, 120: 654, 643, 323, 320, 704, 699, 700, 702, 703, 698, 701, 711, 140: 710, 173: 763},
		{91, 91, 91, 91, 91, 91, 91, 91, 9: 91, 11: 91, 13: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 25: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 51: 91, 54: 91, 83: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91},
		{90, 90, 90, 90, 90, 90, 90, 90, 9: 90, 11: 90, 13: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 25: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 51: 90, 54: 90, 83: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90},
		// 400
		{9: 762},
		{756},
		{752},
		{748},
		{382, 387, 696, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 693, 54: 653, 83: 718, 715, 720, 705, 719, 706, 716, 717, 713, 721, 714, 111: 697, 118: 712, 120: 654, 124: 704, 699, 700, 702, 703, 698, 701, 742},
		// 405
		{728},
		{2: 726},
		{9: 725},
		{9: 724},
		{382, 387, 375, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 111: 697, 118: 722},
		// 410
		{9: 723},
		{78, 78, 78, 78, 78, 78, 78, 78, 9: 78, 11: 78, 13: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 25: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 51: 78, 54: 78, 83: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 109: 78},
		{79, 79, 79, 79, 79, 79, 79, 79, 9: 79, 11: 79, 13: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 25: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 51: 79, 54: 79, 83: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 109: 79},
		{80, 80, 80, 80, 80, 80, 80, 80, 9: 80, 11: 80, 13: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 25: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 51: 80, 54: 80, 83: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 109: 80},
		{9: 727},
		// 415
		{81, 81, 81, 81, 81, 81, 81, 81, 9: 81, 11: 81, 13: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 25: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 51: 81, 54: 81, 83: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 109: 81},
		{382, 387, 375, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 350, 348, 349, 26: 338, 330, 339, 335, 347, 334, 332, 333, 331, 336, 345, 342, 343, 346, 337, 329, 326, 327, 325, 351, 328, 324, 50: 442, 110: 322, 697, 114: 341, 344, 340, 321, 729, 121: 643, 323, 320, 140: 730},
		{9: 736},
		{382, 387, 375, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 111: 697, 118: 731},
		{9: 732},
		// 420
		{382, 387, 375, 386, 388, 389, 385, 384, 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 111: 697, 118: 733},
		{8: 734},
		{382, 387, 696, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 693, 54: 653, 83: 718, 715, 720, 705, 719, 706, 716, 717, 713, 721, 714, 111: 697, 118: 712, 120: 654, 124: 704, 699, 700, 702, 703, 698, 701, 735},
		{82, 82, 82, 82, 82, 82, 82, 82, 9: 82, 11: 82, 13: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 25: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 51: 82, 54: 82, 83: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 109: 82},
		{382, 387, 375, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 111: 697, 118: 737},
		// 425
		{9: 738},
		{382, 387, 375, 386, 388, 389, 385, 384, 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 111: 697, 118: 739},
		{8: 740},
		{382, 387, 696, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 693, 54: 653, 83: 718, 715, 720, 705, 719, 706, 716, 717, 713, 721, 714, 111: 697, 118: 712, 120: 654, 124: 704, 699, 700, 702, 703, 698, 701, 741},
		{83, 83, 83, 83, 83, 83, 83, 83, 9: 83, 11: 83, 13: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 25: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 51: 83, 54: 83, 83: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 109: 83},
		// 430
		{84: 743},
		{744},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 111: 745},
		{8: 746, 10: 445},
		{9: 747},
		// 435
		{84, 84, 84, 84, 84, 84, 84, 84, 9: 84, 11: 84, 13: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 25: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 51: 84, 54: 84, 83: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 109: 84},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 111: 749},
		{8: 750, 10: 445},
		{382, 387, 696, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 693, 54: 653, 83: 718, 715, 720, 705, 719, 706, 716, 717, 713, 721, 714, 111: 697, 118: 712, 120: 654, 124: 704, 699, 700, 702, 703, 698, 701, 751},
		{85, 85, 85, 85, 85, 85, 85, 85, 9: 85, 11: 85, 13: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 25: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 51: 85, 54: 85, 83: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 109: 85},
		// 440
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 111: 753},
		{8: 754, 10: 445},
		{382, 387, 696, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 693, 54: 653, 83: 718, 715, 720, 705, 719, 706, 716, 717, 713, 721, 714, 111: 697, 118: 712, 120: 654, 124: 704, 699, 700, 702, 703, 698, 701, 755},
		{86, 86, 86, 86, 86, 86, 86, 86, 9: 86, 11: 86, 13: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 25: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 51: 86, 54: 86, 83: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 109: 86},
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 111: 757},
		// 445
		{8: 758, 10: 445},
		{382, 387, 696, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 693, 54: 653, 83: 718, 715, 720, 705, 719, 706, 716, 717, 713, 721, 714, 111: 697, 118: 712, 120: 654, 124: 704, 699, 700, 702, 703, 698, 701, 759},
		{88, 88, 88, 88, 88, 88, 88, 88, 9: 88, 11: 88, 13: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 25: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 51: 88, 54: 88, 83: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 109: 760},
		{382, 387, 696, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 693, 54: 653, 83: 718, 715, 720, 705, 719, 706, 716, 717, 713, 721, 714, 111: 697, 118: 712, 120: 654, 124: 704, 699, 700, 702, 703, 698, 701, 761},
		{87, 87, 87, 87, 87, 87, 87, 87, 9: 87, 11: 87, 13: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 25: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 51: 87, 54: 87, 83: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 109: 87},
		// 450
		{89, 89, 89, 89, 89, 89, 89, 89, 9: 89, 11: 89, 13: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 25: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 51: 89, 54: 89, 83: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 109: 89},
		{94, 94, 94, 94, 94, 94, 94, 94, 9: 94, 11: 94, 13: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 25: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 51: 94, 54: 94, 83: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94},
		{96, 96, 96, 96, 96, 96, 96, 96, 9: 96, 11: 96, 13: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 25: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 51: 96, 54: 96, 83: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 109: 96},
		{382, 387, 696, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 693, 54: 653, 83: 718, 715, 720, 705, 719, 706, 716, 717, 713, 721, 714, 111: 697, 118: 712, 120: 654, 124: 704, 699, 700, 702, 703, 698, 701, 766},
		{98, 98, 98, 98, 98, 98, 98, 98, 9: 98, 11: 98, 13: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 25: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 51: 98, 54: 98, 83: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 109: 98},
		// 455
		{24: 768},
		{382, 387, 696, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 693, 54: 653, 83: 718, 715, 720, 705, 719, 706, 716, 717, 713, 721, 714, 111: 697, 118: 712, 120: 654, 124: 704, 699, 700, 702, 703, 698, 701, 769},
		{99, 99, 99, 99, 99, 99, 99, 99, 9: 99, 11: 99, 13: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 25: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 51: 99, 54: 99, 83: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 109: 99},
		{382, 387, 696, 386, 388, 389, 385, 384, 9: 240, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 442, 693, 54: 653, 83: 718, 715, 720, 705, 719, 706, 716, 717, 713, 721, 714, 111: 697, 118: 712, 120: 654, 124: 704, 699, 700, 702, 703, 698, 701, 771},
		{100, 100, 100, 100, 100, 100, 100, 100, 9: 100, 11: 100, 13: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 25: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 51: 100, 54: 100, 83: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 109: 100},
		// 460
		{382, 387, 375, 386, 388, 389, 385, 384, 11: 381, 13: 391, 390, 376, 377, 378, 379, 380, 392, 50: 578, 579, 166: 773},
		{9: 223, 223},
		{156, 482, 156, 132: 617, 616, 135: 647, 165: 775},
		{9: 228, 228},
		{236, 236, 236, 236, 236, 236, 236, 236, 9: 236, 11: 236, 13: 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 25: 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 51: 236, 54: 236, 83: 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236},
		// 465
		{21: 76, 76, 76, 26: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 54: 76},
		{238, 238, 238, 238, 238, 238, 238, 238, 11: 238, 13: 238, 238, 238, 238, 238, 238, 238, 238, 143: 373, 779},
		{48: 308},
		{81: 802, 804, 97: 792, 793, 794, 789, 790, 791, 795, 799, 796, 786, 797, 798, 112: 803, 801, 119: 800, 134: 784, 136: 783, 788, 785, 787, 141: 782, 223: 781},
		{48: 310},
		// 470
		{48: 44, 81: 802, 804, 97: 792, 793, 794, 789, 790, 791, 795, 799, 796, 786, 797, 798, 112: 803, 801, 119: 800, 134: 784, 136: 847, 788, 785, 787},
		{48: 43, 81: 43, 43, 94: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43},
		{48: 39, 81: 39, 39, 94: 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39},
		{48: 38, 81: 38, 38, 94: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{82: 804, 112: 869, 801},
		// 475
		{48: 36, 81: 36, 36, 94: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
		{94: 29, 29, 857, 179: 855, 214: 856, 854},
		{82: 804, 112: 851, 801},
		{2: 848},
		{2: 843},
		// 480
		{2: 819, 81: 821, 221: 820},
		{81: 802, 804, 112: 803, 801, 119: 818},
		{48: 17, 81: 17, 17, 94: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{82: 804, 112: 816, 801},
		{82: 804, 112: 814, 801},
		// 485
		{81: 802, 804, 112: 803, 801, 119: 813},
		{2: 809},
		{82: 804, 112: 807, 801},
		{48: 7, 81: 7, 7, 94: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{81: 5, 806},
		// 490
		{48: 4, 81: 4, 4, 94: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{81: 805},
		{81: 2, 2},
		{48: 3, 81: 3, 3, 94: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{81: 1, 1},
		// 495
		{81: 808},
		{48: 8, 81: 8, 8, 94: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{81: 810, 804, 112: 811, 801},
		{48: 13, 81: 13, 13, 94: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{81: 812},
		// 500
		{48: 9, 81: 9, 9, 94: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{48: 14, 81: 14, 14, 94: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{81: 815},
		{48: 15, 81: 15, 15, 94: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{81: 817},
		// 505
		{48: 16, 81: 16, 16, 94: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{48: 18, 81: 18, 18, 94: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{81: 802, 804, 112: 803, 801, 119: 828, 142: 842},
		{2: 822, 8: 140, 145: 824, 164: 823, 184: 825},
		{48: 10, 81: 10, 10, 94: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		// 510
		{8: 142, 10: 142, 145: 839},
		{8: 139, 10: 831},
		{8: 829},
		{8: 826},
		{81: 802, 804, 112: 803, 801, 119: 828, 142: 827},
		// 515
		{48: 19, 81: 19, 19, 94: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{48: 6, 81: 6, 6, 94: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{81: 802, 804, 112: 803, 801, 119: 828, 142: 830},
		{48: 21, 81: 21, 21, 94: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{2: 832, 145: 833},
		// 520
		{8: 141, 10: 141, 145: 836},
		{8: 834},
		{81: 802, 804, 112: 803, 801, 119: 828, 142: 835},
		{48: 20, 81: 20, 20, 94: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{8: 837},
		// 525
		{81: 802, 804, 112: 803, 801, 119: 828, 142: 838},
		{48: 11, 81: 11, 11, 94: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{8: 840},
		{81: 802, 804, 112: 803, 801, 119: 828, 142: 841},
		{48: 12, 81: 12, 12, 94: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		// 530
		{48: 22, 81: 22, 22, 94: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{81: 844},
		{81: 802, 804, 94: 41, 41, 41, 792, 793, 794, 789, 790, 791, 795, 799, 796, 786, 797, 798, 112: 803, 801, 119: 800, 134: 784, 136: 783, 788, 785, 787, 141: 845, 146: 846},
		{81: 802, 804, 94: 40, 40, 40, 792, 793, 794, 789, 790, 791, 795, 799, 796, 786, 797, 798, 112: 803, 801, 119: 800, 134: 784, 136: 847, 788, 785, 787},
		{94: 32, 32, 32},
		// 535
		{48: 42, 81: 42, 42, 94: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{81: 849},
		{81: 802, 804, 94: 41, 41, 41, 792, 793, 794, 789, 790, 791, 795, 799, 796, 786, 797, 798, 112: 803, 801, 119: 800, 134: 784, 136: 783, 788, 785, 787, 141: 845, 146: 850},
		{94: 33, 33, 33},
		{81: 852},
		// 540
		{81: 802, 804, 94: 41, 41, 41, 792, 793, 794, 789, 790, 791, 795, 799, 796, 786, 797, 798, 112: 803, 801, 119: 800, 134: 784, 136: 783, 788, 785, 787, 141: 845, 146: 853},
		{94: 34, 34, 34},
		{94: 25, 863, 216: 864, 862},
		{94: 31, 31, 31},
		{94: 28, 28, 857, 179: 861},
		// 545
		{82: 804, 112: 858, 801},
		{81: 859},
		{81: 802, 804, 94: 41, 41, 41, 792, 793, 794, 789, 790, 791, 795, 799, 796, 786, 797, 798, 112: 803, 801, 119: 800, 134: 784, 136: 783, 788, 785, 787, 141: 845, 146: 860},
		{94: 27, 27, 27},
		{94: 30, 30, 30},
		// 550
		{94: 868, 218: 867},
		{81: 865},
		{94: 24},
		{81: 802, 804, 94: 41, 97: 792, 793, 794, 789, 790, 791, 795, 799, 796, 786, 797, 798, 112: 803, 801, 119: 800, 134: 784, 136: 783, 788, 785, 787, 141: 845, 146: 866},
		{94: 26},
		// 555
		{48: 35, 81: 35, 35, 94: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{48: 23, 81: 23, 23, 94: 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23},
		{81: 870},
		{48: 37, 81: 37, 37, 94: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
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
	const yyError = 231

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
			if !IsIntType(e.Type) {
				lx.report.Err(e.Pos(), "bit field width not an integer (have '%s')", e.Type)
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
				var err error
				if lhs.elements, err = elements(o.Expression.eval(lx)); err != nil {
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
			var err error
			if lhs.elements, err = elements(lhs.Expression.eval(lx)); err != nil {
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
			var err error
			if lhs.elements, err = elements(lhs.Expression.eval(lx)); err != nil {
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
			yyVAL.node = &Statement{
				Case:               6,
				AssemblerStatement: yyS[yypt-0].node.(*AssemblerStatement),
			}
		}
	case 212:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 213:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 214:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 215:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 216:
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
	case 217:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 218:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 219:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 220:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 221:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 222:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 223:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 224:
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
	case 225:
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
	case 226:
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
	case 227:
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
	case 228:
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
	case 229:
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
	case 230:
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
	case 231:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 232:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 233:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 234:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 235:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 236:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 237:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 238:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 239:
		{
			yyVAL.node = &ExternalDeclaration{
				Case: 2,
				BasicAssemblerStatement: yyS[yypt-0].node.(*BasicAssemblerStatement),
			}
		}
	case 240:
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
	case 241:
		{
			yyVAL.node = &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:          yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 242:
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
	case 243:
		{
			lhs := &FunctionBody{
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
			yyVAL.node = lhs
			lhs.scope = lhs.CompoundStatement.scope
		}
	case 244:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 245:
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
	case 246:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 247:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 248:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 249:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 250:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScopePos(lhs.Pos())
		}
	case 251:
		{
			yyVAL.node = &AssemblerInstructions{
				Token: yyS[yypt-0].Token,
			}
		}
	case 252:
		{
			yyVAL.node = &AssemblerInstructions{
				Case: 1,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions),
				Token: yyS[yypt-0].Token,
			}
		}
	case 253:
		{
			yyVAL.node = &BasicAssemblerStatement{
				Token:                 yyS[yypt-4].Token,
				VolatileOpt:           yyS[yypt-3].node.(*VolatileOpt),
				Token2:                yyS[yypt-2].Token,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-0].Token,
			}
		}
	case 254:
		{
			yyVAL.node = (*VolatileOpt)(nil)
		}
	case 255:
		{
			yyVAL.node = &VolatileOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 256:
		{
			yyVAL.node = &AssemblerOperand{
				AssemblerSymbolicNameOpt: yyS[yypt-4].node.(*AssemblerSymbolicNameOpt),
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
		}
	case 257:
		{
			yyVAL.node = &AssemblerOperands{
				AssemblerOperand: yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 258:
		{
			yyVAL.node = &AssemblerOperands{
				Case:              1,
				AssemblerOperands: yyS[yypt-2].node.(*AssemblerOperands),
				Token:             yyS[yypt-1].Token,
				AssemblerOperand:  yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 259:
		{
			yyVAL.node = (*AssemblerSymbolicNameOpt)(nil)
		}
	case 260:
		{
			yyVAL.node = &AssemblerSymbolicNameOpt{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 261:
		{
			yyVAL.node = &Clobbers{
				Token: yyS[yypt-0].Token,
			}
		}
	case 262:
		{
			yyVAL.node = &Clobbers{
				Case:     1,
				Clobbers: yyS[yypt-2].node.(*Clobbers),
				Token:    yyS[yypt-1].Token,
				Token2:   yyS[yypt-0].Token,
			}
		}
	case 263:
		{
			yyVAL.node = &AssemblerStatement{
				BasicAssemblerStatement: yyS[yypt-0].node.(*BasicAssemblerStatement),
			}
		}
	case 264:
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
	case 265:
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
	case 266:
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
	case 267:
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
	case 268:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 269:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 270:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 271:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 272:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 273:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 274:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 275:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 276:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 277:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 278:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 279:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 280:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 281:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 282:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 283:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 284:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 285:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 286:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 287:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 288:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 289:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 290:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 291:
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
	case 292:
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
	case 293:
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
	case 294:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 295:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 296:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 297:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 298:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 299:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 300:
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
	case 301:
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
	case 302:
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
	case 303:
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
	case 304:
		{
			yyVAL.node = &ControlLine{
				Case:        14,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 307:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 308:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
