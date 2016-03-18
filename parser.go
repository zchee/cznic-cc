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
	yyDefault           = 57456
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
	STRINGLITERAL       = 57425
	STRUCT              = 57426
	SUBASSIGN           = 57427
	SWITCH              = 57428
	TRANSLATION_UNIT    = 1048578
	TYPEDEF             = 57429
	TYPEDEFNAME         = 57430
	TYPEOF              = 57431
	UNARY               = 57432
	UNION               = 57433
	UNSIGNED            = 57434
	VOID                = 57435
	VOLATILE            = 57436
	WHILE               = 57437
	XORASSIGN           = 57438
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -313
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (301x)
		42:      1,   // '*' (263x)
		57377:   2,   // IDENTIFIER (216x)
		38:      3,   // '&' (211x)
		43:      4,   // '+' (211x)
		45:      5,   // '-' (211x)
		57363:   6,   // DEC (211x)
		57381:   7,   // INC (211x)
		41:      8,   // ')' (196x)
		59:      9,   // ';' (192x)
		44:      10,  // ',' (182x)
		91:      11,  // '[' (158x)
		57425:   12,  // STRINGLITERAL (158x)
		33:      13,  // '!' (141x)
		126:     14,  // '~' (141x)
		57347:   15,  // ALIGNOF (141x)
		57358:   16,  // CHARCONST (141x)
		57373:   17,  // FLOATCONST (141x)
		57384:   18,  // INTCONST (141x)
		57387:   19,  // LONGCHARCONST (141x)
		57388:   20,  // LONGSTRINGLITERAL (141x)
		57423:   21,  // SIZEOF (141x)
		57436:   22,  // VOLATILE (130x)
		57360:   23,  // CONST (128x)
		57417:   24,  // RESTRICT (128x)
		58:      25,  // ':' (119x)
		125:     26,  // '}' (118x)
		57353:   27,  // BOOL (118x)
		57357:   28,  // CHAR (118x)
		57359:   29,  // COMPLEX (118x)
		57367:   30,  // DOUBLE (118x)
		57369:   31,  // ENUM (118x)
		57372:   32,  // FLOAT (118x)
		57383:   33,  // INT (118x)
		57386:   34,  // LONG (118x)
		57421:   35,  // SHORT (118x)
		57422:   36,  // SIGNED (118x)
		57426:   37,  // STRUCT (118x)
		57430:   38,  // TYPEDEFNAME (118x)
		57431:   39,  // TYPEOF (118x)
		57433:   40,  // UNION (118x)
		57434:   41,  // UNSIGNED (118x)
		57435:   42,  // VOID (118x)
		57424:   43,  // STATIC (110x)
		57352:   44,  // AUTO (104x)
		57371:   45,  // EXTERN (104x)
		57382:   46,  // INLINE (104x)
		57416:   47,  // REGISTER (104x)
		57429:   48,  // TYPEDEF (104x)
		57344:   49,  // $end (101x)
		61:      50,  // '=' (90x)
		57499:   51,  // Expression (85x)
		93:      52,  // ']' (83x)
		123:     53,  // '{' (83x)
		46:      54,  // '.' (79x)
		57351:   55,  // ASM (74x)
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
		57427:   80,  // SUBASSIGN (71x)
		57438:   81,  // XORASSIGN (71x)
		10:      82,  // '\n' (60x)
		57412:   83,  // PPOTHER (54x)
		57376:   84,  // GOTO (50x)
		57437:   85,  // WHILE (48x)
		57354:   86,  // BREAK (47x)
		57355:   87,  // CASE (47x)
		57361:   88,  // CONTINUE (47x)
		57364:   89,  // DEFAULT (47x)
		57366:   90,  // DO (47x)
		57374:   91,  // FOR (47x)
		57380:   92,  // IF (47x)
		57418:   93,  // RETURN (47x)
		57428:   94,  // SWITCH (47x)
		57400:   95,  // PPENDIF (45x)
		57399:   96,  // PPELSE (41x)
		57398:   97,  // PPELIF (40x)
		57397:   98,  // PPDEFINE (36x)
		57401:   99,  // PPERROR (36x)
		57402:   100, // PPHASH_NL (36x)
		57404:   101, // PPIF (36x)
		57405:   102, // PPIFDEF (36x)
		57406:   103, // PPIFNDEF (36x)
		57407:   104, // PPINCLUDE (36x)
		57408:   105, // PPINCLUDE_NEXT (36x)
		57409:   106, // PPLINE (36x)
		57410:   107, // PPNONDIRECTIVE (36x)
		57414:   108, // PPPRAGMA (36x)
		57415:   109, // PPUNDEF (36x)
		57368:   110, // ELSE (29x)
		57550:   111, // TypeQualifier (28x)
		57500:   112, // ExpressionList (26x)
		57524:   113, // PPTokenList (23x)
		57526:   114, // PPTokens (23x)
		57495:   115, // EnumSpecifier (20x)
		57545:   116, // StructOrUnion (20x)
		57546:   117, // StructOrUnionSpecifier (20x)
		57553:   118, // TypeSpecifier (20x)
		57501:   119, // ExpressionListOpt (18x)
		57525:   120, // PPTokenListOpt (16x)
		57466:   121, // BasicAssemblerStatement (15x)
		57478:   122, // DeclarationSpecifiers (15x)
		57507:   123, // FunctionSpecifier (15x)
		57540:   124, // StorageClassSpecifier (15x)
		57464:   125, // AssemblerStatement (13x)
		57472:   126, // CompoundStatement (13x)
		57503:   127, // ExpressionStatement (12x)
		57521:   128, // IterationStatement (12x)
		57522:   129, // JumpStatement (12x)
		57523:   130, // LabeledStatement (12x)
		57535:   131, // SelectionStatement (12x)
		57539:   132, // Statement (12x)
		57531:   133, // Pointer (11x)
		57532:   134, // PointerOpt (10x)
		57474:   135, // ControlLine (8x)
		57480:   136, // Declarator (8x)
		57510:   137, // GroupPart (8x)
		57514:   138, // IfGroup (8x)
		57515:   139, // IfSection (8x)
		57547:   140, // TextLine (8x)
		57475:   141, // Declaration (7x)
		57508:   142, // GroupList (6x)
		57534:   143, // ReplacementList (6x)
		57450:   144, // $@4 (5x)
		57473:   145, // ConstantExpression (5x)
		57362:   146, // DDD (5x)
		57509:   147, // GroupListOpt (5x)
		57536:   148, // SpecifierQualifierList (5x)
		57551:   149, // TypeQualifierList (5x)
		57440:   150, // $@10 (4x)
		57457:   151, // AbstractDeclarator (4x)
		57462:   152, // AssemblerOperand (4x)
		57465:   153, // AssemblerSymbolicNameOpt (4x)
		57479:   154, // DeclarationSpecifiersOpt (4x)
		57484:   155, // Designator (4x)
		57527:   156, // ParameterDeclaration (4x)
		57549:   157, // TypeName (4x)
		57552:   158, // TypeQualifierListOpt (4x)
		57461:   159, // AssemblerInstructions (3x)
		57463:   160, // AssemblerOperands (3x)
		57471:   161, // CommaOpt (3x)
		57482:   162, // Designation (3x)
		57483:   163, // DesignationOpt (3x)
		57485:   164, // DesignatorList (3x)
		57502:   165, // ExpressionOpt (3x)
		57511:   166, // IdentifierList (3x)
		57516:   167, // InitDeclarator (3x)
		57519:   168, // Initializer (3x)
		57528:   169, // ParameterList (3x)
		57529:   170, // ParameterTypeList (3x)
		57441:   171, // $@11 (2x)
		57451:   172, // $@5 (2x)
		57458:   173, // AbstractDeclaratorOpt (2x)
		57467:   174, // BlockItem (2x)
		57470:   175, // Clobbers (2x)
		57481:   176, // DeclaratorOpt (2x)
		57486:   177, // DirectAbstractDeclarator (2x)
		57487:   178, // DirectAbstractDeclaratorOpt (2x)
		57488:   179, // DirectDeclarator (2x)
		57489:   180, // ElifGroup (2x)
		57496:   181, // EnumerationConstant (2x)
		57497:   182, // Enumerator (2x)
		57504:   183, // ExternalDeclaration (2x)
		57506:   184, // FunctionDefinition (2x)
		57512:   185, // IdentifierListOpt (2x)
		57513:   186, // IdentifierOpt (2x)
		57517:   187, // InitDeclaratorList (2x)
		57518:   188, // InitDeclaratorListOpt (2x)
		57520:   189, // InitializerList (2x)
		57530:   190, // ParameterTypeListOpt (2x)
		57537:   191, // SpecifierQualifierListOpt (2x)
		57541:   192, // StructDeclaration (2x)
		57543:   193, // StructDeclarator (2x)
		57554:   194, // VolatileOpt (2x)
		57439:   195, // $@1 (1x)
		57442:   196, // $@12 (1x)
		57443:   197, // $@13 (1x)
		57444:   198, // $@14 (1x)
		57445:   199, // $@15 (1x)
		57446:   200, // $@16 (1x)
		57447:   201, // $@17 (1x)
		57448:   202, // $@2 (1x)
		57449:   203, // $@3 (1x)
		57452:   204, // $@6 (1x)
		57453:   205, // $@7 (1x)
		57454:   206, // $@8 (1x)
		57455:   207, // $@9 (1x)
		57459:   208, // ArgumentExpressionList (1x)
		57460:   209, // ArgumentExpressionListOpt (1x)
		57468:   210, // BlockItemList (1x)
		57469:   211, // BlockItemListOpt (1x)
		1048577: 212, // CONSTANT_EXPRESSION (1x)
		57476:   213, // DeclarationList (1x)
		57477:   214, // DeclarationListOpt (1x)
		57490:   215, // ElifGroupList (1x)
		57491:   216, // ElifGroupListOpt (1x)
		57492:   217, // ElseGroup (1x)
		57493:   218, // ElseGroupOpt (1x)
		57494:   219, // EndifLine (1x)
		57498:   220, // EnumeratorList (1x)
		57505:   221, // FunctionBody (1x)
		57378:   222, // IDENTIFIER_LPAREN (1x)
		1048576: 223, // PREPROCESSING_FILE (1x)
		57533:   224, // PreprocessingFile (1x)
		57538:   225, // Start (1x)
		57542:   226, // StructDeclarationList (1x)
		57544:   227, // StructDeclaratorList (1x)
		1048578: 228, // TRANSLATION_UNIT (1x)
		57548:   229, // TranslationUnit (1x)
		57456:   230, // $default (0x)
		57356:   231, // CAST (0x)
		57345:   232, // error (0x)
		57379:   233, // IDENTIFIER_NONREPL (0x)
		57394:   234, // NOELSE (0x)
		57403:   235, // PPHEADER_NAME (0x)
		57411:   236, // PPNUMBER (0x)
		57413:   237, // PPPASTE (0x)
		57432:   238, // UNARY (0x)
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
		57425:   "string literal",
		57347:   "_Alignof",
		57358:   "character constant",
		57373:   "floating-point constant",
		57384:   "integer constant",
		57387:   "long character constant",
		57388:   "long string constant",
		57423:   "sizeof",
		57436:   "volatile",
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
		57426:   "struct",
		57430:   "typedefname",
		57431:   "typeof",
		57433:   "union",
		57434:   "unsigned",
		57435:   "void",
		57424:   "static",
		57352:   "auto",
		57371:   "extern",
		57382:   "inline",
		57416:   "register",
		57429:   "typedef",
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
		57427:   "-=",
		57438:   "^=",
		57412:   "ppother",
		57376:   "goto",
		57437:   "while",
		57354:   "break",
		57355:   "case",
		57361:   "continue",
		57364:   "default",
		57366:   "do",
		57374:   "for",
		57380:   "if",
		57418:   "return",
		57428:   "switch",
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
		1:   {195, 0},
		2:   {225, 3},
		3:   {202, 0},
		4:   {225, 3},
		5:   {203, 0},
		6:   {225, 3},
		7:   {181, 1},
		8:   {208, 1},
		9:   {208, 3},
		10:  {209, 0},
		11:  {209, 1},
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
		69:  {165, 0},
		70:  {165, 1},
		71:  {112, 1},
		72:  {112, 3},
		73:  {119, 0},
		74:  {119, 1},
		75:  {144, 0},
		76:  {145, 2},
		77:  {141, 3},
		78:  {122, 2},
		79:  {122, 2},
		80:  {122, 2},
		81:  {122, 2},
		82:  {154, 0},
		83:  {154, 1},
		84:  {187, 1},
		85:  {187, 3},
		86:  {188, 0},
		87:  {188, 1},
		88:  {167, 1},
		89:  {172, 0},
		90:  {167, 4},
		91:  {124, 1},
		92:  {124, 1},
		93:  {124, 1},
		94:  {124, 1},
		95:  {124, 1},
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
		112: {204, 0},
		113: {205, 0},
		114: {117, 7},
		115: {117, 2},
		116: {116, 1},
		117: {116, 1},
		118: {226, 1},
		119: {226, 2},
		120: {192, 3},
		121: {192, 2},
		122: {148, 2},
		123: {148, 2},
		124: {191, 0},
		125: {191, 1},
		126: {227, 1},
		127: {227, 3},
		128: {193, 1},
		129: {193, 3},
		130: {161, 0},
		131: {161, 1},
		132: {206, 0},
		133: {115, 7},
		134: {115, 2},
		135: {220, 1},
		136: {220, 3},
		137: {182, 1},
		138: {182, 3},
		139: {111, 1},
		140: {111, 1},
		141: {111, 1},
		142: {123, 1},
		143: {136, 2},
		144: {176, 0},
		145: {176, 1},
		146: {179, 1},
		147: {179, 3},
		148: {179, 5},
		149: {179, 6},
		150: {179, 6},
		151: {179, 5},
		152: {207, 0},
		153: {179, 5},
		154: {179, 4},
		155: {133, 2},
		156: {133, 3},
		157: {134, 0},
		158: {134, 1},
		159: {149, 1},
		160: {149, 2},
		161: {158, 0},
		162: {158, 1},
		163: {170, 1},
		164: {170, 3},
		165: {190, 0},
		166: {190, 1},
		167: {169, 1},
		168: {169, 3},
		169: {156, 2},
		170: {156, 2},
		171: {166, 1},
		172: {166, 3},
		173: {185, 0},
		174: {185, 1},
		175: {186, 0},
		176: {186, 1},
		177: {150, 0},
		178: {157, 3},
		179: {151, 1},
		180: {151, 2},
		181: {173, 0},
		182: {173, 1},
		183: {177, 3},
		184: {177, 4},
		185: {177, 5},
		186: {177, 6},
		187: {177, 6},
		188: {177, 4},
		189: {171, 0},
		190: {177, 4},
		191: {196, 0},
		192: {177, 5},
		193: {178, 0},
		194: {178, 1},
		195: {168, 1},
		196: {168, 4},
		197: {189, 2},
		198: {189, 4},
		199: {162, 2},
		200: {163, 0},
		201: {163, 1},
		202: {164, 1},
		203: {164, 2},
		204: {155, 3},
		205: {155, 2},
		206: {132, 1},
		207: {132, 1},
		208: {132, 1},
		209: {132, 1},
		210: {132, 1},
		211: {132, 1},
		212: {132, 1},
		213: {130, 3},
		214: {130, 4},
		215: {130, 3},
		216: {197, 0},
		217: {126, 4},
		218: {210, 1},
		219: {210, 2},
		220: {211, 0},
		221: {211, 1},
		222: {174, 1},
		223: {174, 1},
		224: {127, 2},
		225: {131, 5},
		226: {131, 7},
		227: {131, 5},
		228: {128, 5},
		229: {128, 7},
		230: {128, 9},
		231: {128, 8},
		232: {129, 3},
		233: {129, 2},
		234: {129, 2},
		235: {129, 3},
		236: {229, 1},
		237: {229, 2},
		238: {183, 1},
		239: {183, 1},
		240: {183, 1},
		241: {198, 0},
		242: {184, 5},
		243: {199, 0},
		244: {221, 2},
		245: {200, 0},
		246: {221, 3},
		247: {213, 1},
		248: {213, 2},
		249: {214, 0},
		250: {201, 0},
		251: {214, 2},
		252: {159, 1},
		253: {159, 2},
		254: {121, 5},
		255: {194, 0},
		256: {194, 1},
		257: {152, 5},
		258: {160, 1},
		259: {160, 3},
		260: {153, 0},
		261: {153, 3},
		262: {175, 1},
		263: {175, 3},
		264: {125, 1},
		265: {125, 7},
		266: {125, 9},
		267: {125, 11},
		268: {125, 13},
		269: {224, 1},
		270: {142, 1},
		271: {142, 2},
		272: {147, 0},
		273: {147, 1},
		274: {137, 1},
		275: {137, 1},
		276: {137, 3},
		277: {137, 1},
		278: {139, 4},
		279: {138, 4},
		280: {138, 4},
		281: {138, 4},
		282: {215, 1},
		283: {215, 2},
		284: {216, 0},
		285: {216, 1},
		286: {180, 4},
		287: {217, 3},
		288: {218, 0},
		289: {218, 1},
		290: {219, 1},
		291: {135, 3},
		292: {135, 5},
		293: {135, 7},
		294: {135, 5},
		295: {135, 2},
		296: {135, 1},
		297: {135, 3},
		298: {135, 3},
		299: {135, 2},
		300: {135, 3},
		301: {135, 6},
		302: {135, 8},
		303: {135, 2},
		304: {135, 4},
		305: {135, 3},
		306: {140, 1},
		307: {143, 1},
		308: {113, 1},
		309: {120, 1},
		310: {120, 2},
		311: {114, 1},
		312: {114, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 49}:   "invalid empty input",
		yyXError{556, -1}: "expected #endif",
		yyXError{558, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{471, -1}: "expected $end",
		yyXError{473, -1}: "expected $end",
		yyXError{31, -1}:  "expected '('",
		yyXError{45, -1}:  "expected '('",
		yyXError{81, -1}:  "expected '('",
		yyXError{349, -1}: "expected '('",
		yyXError{370, -1}: "expected '('",
		yyXError{405, -1}: "expected '('",
		yyXError{406, -1}: "expected '('",
		yyXError{407, -1}: "expected '('",
		yyXError{409, -1}: "expected '('",
		yyXError{435, -1}: "expected '('",
		yyXError{83, -1}:  "expected ')'",
		yyXError{90, -1}:  "expected ')'",
		yyXError{182, -1}: "expected ')'",
		yyXError{197, -1}: "expected ')'",
		yyXError{200, -1}: "expected ')'",
		yyXError{203, -1}: "expected ')'",
		yyXError{211, -1}: "expected ')'",
		yyXError{216, -1}: "expected ')'",
		yyXError{222, -1}: "expected ')'",
		yyXError{238, -1}: "expected ')'",
		yyXError{243, -1}: "expected ')'",
		yyXError{254, -1}: "expected ')'",
		yyXError{289, -1}: "expected ')'",
		yyXError{320, -1}: "expected ')'",
		yyXError{425, -1}: "expected ')'",
		yyXError{431, -1}: "expected ')'",
		yyXError{516, -1}: "expected ')'",
		yyXError{517, -1}: "expected ')'",
		yyXError{525, -1}: "expected ')'",
		yyXError{528, -1}: "expected ')'",
		yyXError{531, -1}: "expected ')'",
		yyXError{307, -1}: "expected ':'",
		yyXError{352, -1}: "expected ':'",
		yyXError{398, -1}: "expected ':'",
		yyXError{459, -1}: "expected ':'",
		yyXError{328, -1}: "expected ';'",
		yyXError{344, -1}: "expected ';'",
		yyXError{404, -1}: "expected ';'",
		yyXError{411, -1}: "expected ';'",
		yyXError{412, -1}: "expected ';'",
		yyXError{414, -1}: "expected ';'",
		yyXError{418, -1}: "expected ';'",
		yyXError{421, -1}: "expected ';'",
		yyXError{423, -1}: "expected ';'",
		yyXError{429, -1}: "expected ';'",
		yyXError{438, -1}: "expected ';'",
		yyXError{332, -1}: "expected '='",
		yyXError{95, -1}:  "expected '['",
		yyXError{495, -1}: "expected '\\n'",
		yyXError{499, -1}: "expected '\\n'",
		yyXError{503, -1}: "expected '\\n'",
		yyXError{506, -1}: "expected '\\n'",
		yyXError{508, -1}: "expected '\\n'",
		yyXError{535, -1}: "expected '\\n'",
		yyXError{540, -1}: "expected '\\n'",
		yyXError{543, -1}: "expected '\\n'",
		yyXError{550, -1}: "expected '\\n'",
		yyXError{555, -1}: "expected '\\n'",
		yyXError{561, -1}: "expected '\\n'",
		yyXError{101, -1}: "expected ']'",
		yyXError{190, -1}: "expected ']'",
		yyXError{234, -1}: "expected ']'",
		yyXError{266, -1}: "expected ']'",
		yyXError{358, -1}: "expected ']'",
		yyXError{52, -1}:  "expected '{'",
		yyXError{54, -1}:  "expected '{'",
		yyXError{295, -1}: "expected '{'",
		yyXError{297, -1}: "expected '{'",
		yyXError{275, -1}: "expected '}'",
		yyXError{279, -1}: "expected '}'",
		yyXError{292, -1}: "expected '}'",
		yyXError{399, -1}: "expected '}'",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{210, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{94, -1}:  "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{47, -1}:  "expected assembler instructions or string literal",
		yyXError{348, -1}: "expected assembler instructions or string literal",
		yyXError{350, -1}: "expected assembler instructions or string literal",
		yyXError{360, -1}: "expected assembler operand or one of ['[', string literal]",
		yyXError{353, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{375, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{378, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{343, -1}: "expected assembler statement or asm",
		yyXError{401, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{361, -1}: "expected clobbers or string literal",
		yyXError{381, -1}: "expected clobbers or string literal",
		yyXError{342, -1}: "expected compound statement or '{'",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{60, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{263, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{311, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{397, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{470, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{334, -1}: "expected declaration list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{337, -1}: "expected declaration or one of ['{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{420, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{310, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{202, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{260, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{205, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{91, -1}:  "expected direct abstract declarator or one of ['(', '[']",
		yyXError{308, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{548, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{554, -1}: "expected endif line or #endif",
		yyXError{480, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{546, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{55, -1}:  "expected enumerator list or identifier",
		yyXError{291, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{106, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{130, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{436, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{440, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{444, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{448, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{70, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{251, -1}: "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{255, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		yyXError{98, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{233, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{290, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{61, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{72, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{73, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{74, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{75, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{76, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{77, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{78, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{79, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{80, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{104, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
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
		yyXError{124, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{125, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{126, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{127, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{128, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{129, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{131, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{132, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{133, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{134, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{135, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{136, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{137, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{138, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{139, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{140, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{141, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{156, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{157, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{184, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{191, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{227, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{230, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{371, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{102, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{225, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{318, -1}: "expected expression or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{6, -1}:   "expected external declaration or one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{333, -1}: "expected function body or one of ['{', asm]",
		yyXError{340, -1}: "expected function body or one of ['{', asm]",
		yyXError{331, -1}: "expected function body or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{537, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{474, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{108, -1}: "expected identifier",
		yyXError{109, -1}: "expected identifier",
		yyXError{219, -1}: "expected identifier",
		yyXError{264, -1}: "expected identifier",
		yyXError{357, -1}: "expected identifier",
		yyXError{410, -1}: "expected identifier",
		yyXError{482, -1}: "expected identifier",
		yyXError{483, -1}: "expected identifier",
		yyXError{490, -1}: "expected identifier",
		yyXError{365, -1}: "expected identifier list or identifier",
		yyXError{512, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{466, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{257, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{271, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{259, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{277, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{464, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{274, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{63, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{64, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{65, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{66, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{67, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{68, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{69, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{110, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{111, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
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
		yyXError{174, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{175, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{176, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{177, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{178, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{179, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{183, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{187, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{195, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{250, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{252, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{256, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{280, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{281, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{282, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{283, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{284, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{285, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{286, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{287, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{288, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{71, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{154, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{158, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{180, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{185, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{319, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{372, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{388, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{270, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{97, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{105, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{192, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{228, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{231, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{475, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{476, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{477, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{479, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{486, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{492, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{494, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{497, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{500, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{502, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{504, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{505, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{507, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{509, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{510, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{513, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{519, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{520, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{522, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{527, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{530, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{533, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{534, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{539, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{559, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{560, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{562, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{538, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{542, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{545, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{547, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{552, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{553, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{51, -1}:  "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{456, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{468, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{40, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{41, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{42, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{43, -1}:  "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{341, -1}: "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{384, -1}: "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{386, -1}: "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{469, -1}: "expected one of [$end, _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{36, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{38, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{99, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{188, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{346, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{367, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{377, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{380, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{383, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{390, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{391, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{392, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{393, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{394, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{395, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{396, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{415, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{416, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{417, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{419, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{427, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{433, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{439, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{443, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{447, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{451, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{453, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{454, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{458, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{461, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{463, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{400, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{402, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{403, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{455, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{261, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{268, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{53, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{296, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
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
		yyXError{293, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{316, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{321, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{322, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{12, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{39, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{323, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{324, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{325, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{326, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{327, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{247, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{248, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{249, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{208, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{209, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{212, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{221, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{223, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{229, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{232, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{235, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{236, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{89, -1}:  "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{246, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{93, -1}:  "expected one of ['(', ')', ',', '[']",
		yyXError{142, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{189, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{193, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{194, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{196, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{204, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{240, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{244, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{46, -1}:  "expected one of ['(', goto]",
		yyXError{347, -1}: "expected one of ['(', goto]",
		yyXError{309, -1}: "expected one of ['(', identifier]",
		yyXError{355, -1}: "expected one of [')', ',', ':']",
		yyXError{362, -1}: "expected one of [')', ',', ':']",
		yyXError{368, -1}: "expected one of [')', ',', ':']",
		yyXError{369, -1}: "expected one of [')', ',', ':']",
		yyXError{373, -1}: "expected one of [')', ',', ':']",
		yyXError{376, -1}: "expected one of [')', ',', ':']",
		yyXError{379, -1}: "expected one of [')', ',', ':']",
		yyXError{389, -1}: "expected one of [')', ',', ';']",
		yyXError{514, -1}: "expected one of [')', ',', ...]",
		yyXError{524, -1}: "expected one of [')', ',', ...]",
		yyXError{92, -1}:  "expected one of [')', ',']",
		yyXError{181, -1}: "expected one of [')', ',']",
		yyXError{199, -1}: "expected one of [')', ',']",
		yyXError{201, -1}: "expected one of [')', ',']",
		yyXError{206, -1}: "expected one of [')', ',']",
		yyXError{207, -1}: "expected one of [')', ',']",
		yyXError{217, -1}: "expected one of [')', ',']",
		yyXError{218, -1}: "expected one of [')', ',']",
		yyXError{220, -1}: "expected one of [')', ',']",
		yyXError{239, -1}: "expected one of [')', ',']",
		yyXError{253, -1}: "expected one of [')', ',']",
		yyXError{366, -1}: "expected one of [')', ',']",
		yyXError{382, -1}: "expected one of [')', ',']",
		yyXError{437, -1}: "expected one of [')', ',']",
		yyXError{441, -1}: "expected one of [')', ',']",
		yyXError{445, -1}: "expected one of [')', ',']",
		yyXError{449, -1}: "expected one of [')', ',']",
		yyXError{515, -1}: "expected one of [')', ',']",
		yyXError{48, -1}:  "expected one of [')', ':', string literal]",
		yyXError{50, -1}:  "expected one of [')', ':', string literal]",
		yyXError{374, -1}: "expected one of [')', ':', string literal]",
		yyXError{49, -1}:  "expected one of [')', string literal]",
		yyXError{306, -1}: "expected one of [',', ':', ';']",
		yyXError{155, -1}: "expected one of [',', ':']",
		yyXError{356, -1}: "expected one of [',', ':']",
		yyXError{363, -1}: "expected one of [',', ':']",
		yyXError{339, -1}: "expected one of [',', ';', '=']",
		yyXError{276, -1}: "expected one of [',', ';', '}']",
		yyXError{303, -1}: "expected one of [',', ';']",
		yyXError{305, -1}: "expected one of [',', ';']",
		yyXError{312, -1}: "expected one of [',', ';']",
		yyXError{315, -1}: "expected one of [',', ';']",
		yyXError{329, -1}: "expected one of [',', ';']",
		yyXError{330, -1}: "expected one of [',', ';']",
		yyXError{465, -1}: "expected one of [',', ';']",
		yyXError{467, -1}: "expected one of [',', ';']",
		yyXError{56, -1}:  "expected one of [',', '=', '}']",
		yyXError{59, -1}:  "expected one of [',', '=', '}']",
		yyXError{186, -1}: "expected one of [',', ']']",
		yyXError{58, -1}:  "expected one of [',', '}']",
		yyXError{62, -1}:  "expected one of [',', '}']",
		yyXError{272, -1}: "expected one of [',', '}']",
		yyXError{278, -1}: "expected one of [',', '}']",
		yyXError{294, -1}: "expected one of [',', '}']",
		yyXError{262, -1}: "expected one of ['.', '=', '[']",
		yyXError{265, -1}: "expected one of ['.', '=', '[']",
		yyXError{267, -1}: "expected one of ['.', '=', '[']",
		yyXError{269, -1}: "expected one of ['.', '=', '[']",
		yyXError{351, -1}: "expected one of [':', string literal]",
		yyXError{484, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		yyXError{493, -1}: "expected one of ['\\n', ppother]",
		yyXError{496, -1}: "expected one of ['\\n', ppother]",
		yyXError{498, -1}: "expected one of ['\\n', ppother]",
		yyXError{336, -1}: "expected one of ['{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{338, -1}: "expected one of ['{', _Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{33, -1}:  "expected one of ['{', identifier]",
		yyXError{34, -1}:  "expected one of ['{', identifier]",
		yyXError{301, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{304, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{313, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{317, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{523, -1}: "expected one of [..., identifier]",
		yyXError{87, -1}:  "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{107, -1}: "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{385, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{387, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{57, -1}:  "expected optional comma or one of [',', '}']",
		yyXError{258, -1}: "expected optional comma or one of [',', '}']",
		yyXError{273, -1}: "expected optional comma or one of [',', '}']",
		yyXError{8, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{424, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{430, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{413, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{422, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{428, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{224, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{213, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{96, -1}:  "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{100, -1}: "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{536, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{541, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{544, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{551, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{557, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{214, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{32, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{35, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{335, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{198, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{241, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{242, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{85, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{86, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{485, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{489, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{88, -1}:  "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{345, -1}: "expected optional volatile or one of ['(', goto, volatile]",
		yyXError{44, -1}:  "expected optional volatile or one of ['(', volatile]",
		yyXError{237, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{215, -1}: "expected parameter type list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{245, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{472, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{511, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{518, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{521, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{526, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{529, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{532, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{84, -1}:  "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{408, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{426, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{432, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{442, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{446, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{450, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{452, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{457, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{460, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{462, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{354, -1}: "expected string literal",
		yyXError{359, -1}: "expected string literal",
		yyXError{364, -1}: "expected string literal",
		yyXError{298, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{299, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{300, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{302, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		yyXError{314, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{501, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{478, -1}: "expected token list or ppother",
		yyXError{481, -1}: "expected token list or ppother",
		yyXError{487, -1}: "expected token list or ppother",
		yyXError{488, -1}: "expected token list or ppother",
		yyXError{491, -1}: "expected token list or ppother",
		yyXError{549, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [_Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{5, -1}:   "expected translation unit or one of [_Bool, _Complex, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{82, -1}:  "expected type name or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{103, -1}: "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{226, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{434, -1}: "expected while",
		yyXError{3, 49}:   "unexpected EOF",
		yyXError{2, 49}:   "unexpected EOF",
		yyXError{4, 49}:   "unexpected EOF",
	}

	yyParseTab = [563][]uint16{
		// 0
		{212: 316, 223: 315, 225: 314, 228: 317},
		{49: 313},
		{82: 312, 312, 98: 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 195: 785},
		{310, 310, 310, 310, 310, 310, 310, 310, 12: 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 202: 783},
		{22: 308, 308, 308, 27: 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 55: 308, 203: 318},
		// 5
		{22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 55: 357, 111: 323, 115: 342, 345, 341, 322, 121: 356, 320, 324, 321, 141: 355, 183: 353, 354, 229: 319},
		{22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 307, 55: 357, 111: 323, 115: 342, 345, 341, 322, 121: 356, 320, 324, 321, 141: 355, 183: 782, 354},
		{156, 401, 156, 9: 227, 133: 622, 621, 136: 644, 167: 642, 187: 643, 641},
		{231, 231, 231, 8: 231, 231, 231, 231, 22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 111: 323, 115: 342, 345, 341, 322, 122: 637, 324, 321, 154: 640},
		{231, 231, 231, 8: 231, 231, 231, 231, 22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 111: 323, 115: 342, 345, 341, 322, 122: 637, 324, 321, 154: 639},
		// 10
		{231, 231, 231, 8: 231, 231, 231, 231, 22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 111: 323, 115: 342, 345, 341, 322, 122: 637, 324, 321, 154: 638},
		{231, 231, 231, 8: 231, 231, 231, 231, 22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 111: 323, 115: 342, 345, 341, 322, 122: 637, 324, 321, 154: 636},
		{222, 222, 222, 8: 222, 222, 222, 222, 22: 222, 222, 222, 27: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222},
		{221, 221, 221, 8: 221, 221, 221, 221, 22: 221, 221, 221, 27: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221},
		{220, 220, 220, 8: 220, 220, 220, 220, 22: 220, 220, 220, 27: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220},
		// 15
		{219, 219, 219, 8: 219, 219, 219, 219, 22: 219, 219, 219, 27: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219},
		{218, 218, 218, 8: 218, 218, 218, 218, 22: 218, 218, 218, 27: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218},
		{217, 217, 217, 8: 217, 217, 217, 217, 22: 217, 217, 217, 217, 27: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217},
		{216, 216, 216, 8: 216, 216, 216, 216, 22: 216, 216, 216, 216, 27: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216},
		{215, 215, 215, 8: 215, 215, 215, 215, 22: 215, 215, 215, 215, 27: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215},
		// 20
		{214, 214, 214, 8: 214, 214, 214, 214, 22: 214, 214, 214, 214, 27: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214},
		{213, 213, 213, 8: 213, 213, 213, 213, 22: 213, 213, 213, 213, 27: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213},
		{212, 212, 212, 8: 212, 212, 212, 212, 22: 212, 212, 212, 212, 27: 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212},
		{211, 211, 211, 8: 211, 211, 211, 211, 22: 211, 211, 211, 211, 27: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211},
		{210, 210, 210, 8: 210, 210, 210, 210, 22: 210, 210, 210, 210, 27: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210},
		// 25
		{209, 209, 209, 8: 209, 209, 209, 209, 22: 209, 209, 209, 209, 27: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209},
		{208, 208, 208, 8: 208, 208, 208, 208, 22: 208, 208, 208, 208, 27: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208},
		{207, 207, 207, 8: 207, 207, 207, 207, 22: 207, 207, 207, 207, 27: 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207},
		{206, 206, 206, 8: 206, 206, 206, 206, 22: 206, 206, 206, 206, 27: 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206},
		{205, 205, 205, 8: 205, 205, 205, 205, 22: 205, 205, 205, 205, 27: 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205},
		// 30
		{204, 204, 204, 8: 204, 204, 204, 204, 22: 204, 204, 204, 204, 27: 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204},
		{631},
		{2: 609, 53: 138, 186: 608},
		{2: 197, 53: 197},
		{2: 196, 53: 196},
		// 35
		{2: 366, 53: 138, 186: 365},
		{174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 27: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 52: 174},
		{173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 27: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 52: 173},
		{172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 27: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 52: 172},
		{171, 171, 171, 8: 171, 171, 171, 171, 22: 171, 171, 171, 27: 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171},
		// 40
		{22: 77, 77, 77, 27: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 55: 77},
		{22: 75, 75, 75, 27: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 55: 75},
		{22: 74, 74, 74, 27: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 55: 74},
		{22: 73, 73, 73, 27: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 55: 73},
		{58, 22: 359, 194: 358},
		// 45
		{360},
		{57, 84: 57},
		{12: 361, 159: 362},
		{8: 61, 12: 61, 25: 61},
		{8: 364, 12: 363},
		// 50
		{8: 60, 12: 60, 25: 60},
		{59, 59, 59, 59, 59, 59, 59, 59, 9: 59, 12: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 26: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 53: 59, 55: 59, 84: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 110: 59},
		{53: 181, 206: 367},
		{179, 179, 179, 8: 179, 179, 179, 179, 22: 179, 179, 179, 179, 27: 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 53: 137},
		{53: 368},
		// 55
		{2: 369, 181: 372, 371, 220: 370},
		{10: 306, 26: 306, 50: 306},
		{10: 604, 26: 183, 161: 605},
		{10: 178, 26: 178},
		{10: 176, 26: 176, 50: 373},
		// 60
		{238, 238, 238, 238, 238, 238, 238, 238, 12: 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 144: 374, 375},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 384},
		{10: 175, 26: 175},
		{301, 301, 3: 301, 301, 301, 301, 301, 301, 301, 301, 301, 25: 301, 301, 49: 301, 301, 52: 301, 54: 301, 56: 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301},
		{300, 300, 3: 300, 300, 300, 300, 300, 300, 300, 300, 300, 25: 300, 300, 49: 300, 300, 52: 300, 54: 300, 56: 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300},
		// 65
		{299, 299, 3: 299, 299, 299, 299, 299, 299, 299, 299, 299, 25: 299, 299, 49: 299, 299, 52: 299, 54: 299, 56: 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299},
		{298, 298, 3: 298, 298, 298, 298, 298, 298, 298, 298, 298, 25: 298, 298, 49: 298, 298, 52: 298, 54: 298, 56: 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298},
		{297, 297, 3: 297, 297, 297, 297, 297, 297, 297, 297, 297, 25: 297, 297, 49: 297, 297, 52: 297, 54: 297, 56: 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297},
		{296, 296, 3: 296, 296, 296, 296, 296, 296, 296, 296, 296, 25: 296, 296, 49: 296, 296, 52: 296, 54: 296, 56: 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296},
		{295, 295, 3: 295, 295, 295, 295, 295, 295, 295, 295, 295, 25: 295, 295, 49: 295, 295, 52: 295, 54: 295, 56: 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295},
		// 70
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 136, 136, 136, 27: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 51: 467, 112: 566, 150: 397, 157: 602},
		{420, 425, 3: 438, 428, 429, 424, 423, 9: 237, 237, 419, 25: 237, 237, 49: 237, 444, 52: 237, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 601},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 600},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 599},
		// 75
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 508},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 598},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 597},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 596},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 595},
		// 80
		{564, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 565},
		{395},
		{22: 136, 136, 136, 27: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 150: 397, 157: 396},
		{8: 563},
		{22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 111: 399, 115: 342, 345, 341, 398, 148: 400},
		// 85
		{189, 189, 189, 8: 189, 189, 11: 189, 22: 351, 349, 350, 189, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 111: 399, 115: 342, 345, 341, 398, 148: 561, 191: 562},
		{189, 189, 189, 8: 189, 189, 11: 189, 22: 351, 349, 350, 189, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 111: 399, 115: 342, 345, 341, 398, 148: 561, 191: 560},
		{156, 401, 8: 132, 11: 156, 133: 402, 404, 151: 405, 173: 403},
		{152, 152, 152, 8: 152, 10: 152, 152, 22: 351, 349, 350, 111: 412, 149: 416, 158: 558},
		{155, 2: 155, 8: 134, 10: 134, 155},
		// 90
		{8: 135},
		{407, 11: 120, 177: 406, 408},
		{8: 131, 10: 131},
		{554, 8: 133, 10: 133, 119},
		{156, 401, 8: 124, 11: 156, 22: 124, 124, 124, 27: 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 133: 402, 404, 151: 510, 171: 511},
		// 95
		{11: 409},
		{383, 411, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 351, 349, 350, 43: 415, 51: 410, 244, 111: 412, 149: 413, 165: 414},
		{420, 425, 3: 438, 428, 429, 424, 423, 11: 419, 50: 444, 52: 243, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 508, 509},
		{154, 154, 154, 154, 154, 154, 154, 154, 154, 10: 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 43: 154, 52: 154},
		// 100
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 351, 349, 350, 43: 504, 51: 410, 244, 111: 501, 165: 503},
		{52: 502},
		{152, 152, 152, 152, 152, 152, 152, 152, 12: 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 351, 349, 350, 111: 412, 149: 416, 158: 417},
		{151, 151, 151, 151, 151, 151, 151, 151, 151, 10: 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 351, 349, 350, 111: 501},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 418},
		// 105
		{420, 425, 3: 438, 428, 429, 424, 423, 11: 419, 50: 444, 52: 455, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 112: 499},
		{383, 388, 376, 387, 389, 390, 386, 385, 303, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 493, 208: 494, 495},
		{2: 492},
		{2: 491},
		// 110
		{289, 289, 3: 289, 289, 289, 289, 289, 289, 289, 289, 289, 25: 289, 289, 49: 289, 289, 52: 289, 54: 289, 56: 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289},
		{288, 288, 3: 288, 288, 288, 288, 288, 288, 288, 288, 288, 25: 288, 288, 49: 288, 288, 52: 288, 54: 288, 56: 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 490},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 489},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 488},
		// 115
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 487},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 486},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 485},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 484},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 483},
		// 120
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 482},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 481},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 480},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 479},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 478},
		// 125
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 477},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 476},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 475},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 474},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 473},
		// 130
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 112: 468},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 466},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 465},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 464},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 463},
		// 135
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 462},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 461},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 460},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 459},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 458},
		// 140
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 457},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 456},
		{127, 8: 127, 10: 127, 127},
		{420, 425, 3: 438, 428, 429, 424, 423, 246, 246, 246, 419, 25: 246, 246, 49: 246, 444, 52: 246, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{420, 425, 3: 438, 428, 429, 424, 423, 247, 247, 247, 419, 25: 247, 247, 49: 247, 444, 52: 247, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		// 145
		{420, 425, 3: 438, 428, 429, 424, 423, 248, 248, 248, 419, 25: 248, 248, 49: 248, 444, 52: 248, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{420, 425, 3: 438, 428, 429, 424, 423, 249, 249, 249, 419, 25: 249, 249, 49: 249, 444, 52: 249, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{420, 425, 3: 438, 428, 429, 424, 423, 250, 250, 250, 419, 25: 250, 250, 49: 250, 444, 52: 250, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{420, 425, 3: 438, 428, 429, 424, 423, 251, 251, 251, 419, 25: 251, 251, 49: 251, 444, 52: 251, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{420, 425, 3: 438, 428, 429, 424, 423, 252, 252, 252, 419, 25: 252, 252, 49: 252, 444, 52: 252, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		// 150
		{420, 425, 3: 438, 428, 429, 424, 423, 253, 253, 253, 419, 25: 253, 253, 49: 253, 444, 52: 253, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{420, 425, 3: 438, 428, 429, 424, 423, 254, 254, 254, 419, 25: 254, 254, 49: 254, 444, 52: 254, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{420, 425, 3: 438, 428, 429, 424, 423, 255, 255, 255, 419, 25: 255, 255, 49: 255, 444, 52: 255, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{420, 425, 3: 438, 428, 429, 424, 423, 256, 256, 256, 419, 25: 256, 256, 49: 256, 444, 52: 256, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{420, 425, 3: 438, 428, 429, 424, 423, 242, 242, 242, 419, 25: 242, 50: 444, 52: 242, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		// 155
		{10: 470, 25: 469},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 472},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 471},
		{420, 425, 3: 438, 428, 429, 424, 423, 241, 241, 241, 419, 25: 241, 50: 444, 52: 241, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{420, 425, 3: 438, 428, 429, 424, 423, 257, 257, 257, 419, 25: 257, 257, 49: 257, 257, 52: 257, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 257, 441, 257, 422, 257, 436, 435, 434, 430, 257, 257, 257, 437, 257, 442, 431, 257, 257, 257},
		// 160
		{420, 425, 3: 438, 428, 429, 424, 423, 258, 258, 258, 419, 25: 258, 258, 49: 258, 258, 52: 258, 54: 421, 56: 427, 426, 432, 433, 258, 439, 440, 258, 441, 258, 422, 258, 436, 435, 434, 430, 258, 258, 258, 437, 258, 258, 431, 258, 258, 258},
		{420, 425, 3: 438, 428, 429, 424, 423, 259, 259, 259, 419, 25: 259, 259, 49: 259, 259, 52: 259, 54: 421, 56: 427, 426, 432, 433, 259, 439, 440, 259, 259, 259, 422, 259, 436, 435, 434, 430, 259, 259, 259, 437, 259, 259, 431, 259, 259, 259},
		{420, 425, 3: 438, 428, 429, 424, 423, 260, 260, 260, 419, 25: 260, 260, 49: 260, 260, 52: 260, 54: 421, 56: 427, 426, 432, 433, 260, 439, 260, 260, 260, 260, 422, 260, 436, 435, 434, 430, 260, 260, 260, 437, 260, 260, 431, 260, 260, 260},
		{420, 425, 3: 438, 428, 429, 424, 423, 261, 261, 261, 419, 25: 261, 261, 49: 261, 261, 52: 261, 54: 421, 56: 427, 426, 432, 433, 261, 261, 261, 261, 261, 261, 422, 261, 436, 435, 434, 430, 261, 261, 261, 437, 261, 261, 431, 261, 261, 261},
		{420, 425, 3: 262, 428, 429, 424, 423, 262, 262, 262, 419, 25: 262, 262, 49: 262, 262, 52: 262, 54: 421, 56: 427, 426, 432, 433, 262, 262, 262, 262, 262, 262, 422, 262, 436, 435, 434, 430, 262, 262, 262, 437, 262, 262, 431, 262, 262, 262},
		// 165
		{420, 425, 3: 263, 428, 429, 424, 423, 263, 263, 263, 419, 25: 263, 263, 49: 263, 263, 52: 263, 54: 421, 56: 427, 426, 432, 433, 263, 263, 263, 263, 263, 263, 422, 263, 263, 435, 434, 430, 263, 263, 263, 263, 263, 263, 431, 263, 263, 263},
		{420, 425, 3: 264, 428, 429, 424, 423, 264, 264, 264, 419, 25: 264, 264, 49: 264, 264, 52: 264, 54: 421, 56: 427, 426, 432, 433, 264, 264, 264, 264, 264, 264, 422, 264, 264, 435, 434, 430, 264, 264, 264, 264, 264, 264, 431, 264, 264, 264},
		{420, 425, 3: 265, 428, 429, 424, 423, 265, 265, 265, 419, 25: 265, 265, 49: 265, 265, 52: 265, 54: 421, 56: 427, 426, 265, 265, 265, 265, 265, 265, 265, 265, 422, 265, 265, 265, 265, 430, 265, 265, 265, 265, 265, 265, 431, 265, 265, 265},
		{420, 425, 3: 266, 428, 429, 424, 423, 266, 266, 266, 419, 25: 266, 266, 49: 266, 266, 52: 266, 54: 421, 56: 427, 426, 266, 266, 266, 266, 266, 266, 266, 266, 422, 266, 266, 266, 266, 430, 266, 266, 266, 266, 266, 266, 431, 266, 266, 266},
		{420, 425, 3: 267, 428, 429, 424, 423, 267, 267, 267, 419, 25: 267, 267, 49: 267, 267, 52: 267, 54: 421, 56: 427, 426, 267, 267, 267, 267, 267, 267, 267, 267, 422, 267, 267, 267, 267, 430, 267, 267, 267, 267, 267, 267, 431, 267, 267, 267},
		// 170
		{420, 425, 3: 268, 428, 429, 424, 423, 268, 268, 268, 419, 25: 268, 268, 49: 268, 268, 52: 268, 54: 421, 56: 427, 426, 268, 268, 268, 268, 268, 268, 268, 268, 422, 268, 268, 268, 268, 430, 268, 268, 268, 268, 268, 268, 431, 268, 268, 268},
		{420, 425, 3: 269, 428, 429, 424, 423, 269, 269, 269, 419, 25: 269, 269, 49: 269, 269, 52: 269, 54: 421, 56: 427, 426, 269, 269, 269, 269, 269, 269, 269, 269, 422, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269},
		{420, 425, 3: 270, 428, 429, 424, 423, 270, 270, 270, 419, 25: 270, 270, 49: 270, 270, 52: 270, 54: 421, 56: 427, 426, 270, 270, 270, 270, 270, 270, 270, 270, 422, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270},
		{420, 425, 3: 271, 271, 271, 424, 423, 271, 271, 271, 419, 25: 271, 271, 49: 271, 271, 52: 271, 54: 421, 56: 427, 426, 271, 271, 271, 271, 271, 271, 271, 271, 422, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{420, 425, 3: 272, 272, 272, 424, 423, 272, 272, 272, 419, 25: 272, 272, 49: 272, 272, 52: 272, 54: 421, 56: 427, 426, 272, 272, 272, 272, 272, 272, 272, 272, 422, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272},
		// 175
		{420, 273, 3: 273, 273, 273, 424, 423, 273, 273, 273, 419, 25: 273, 273, 49: 273, 273, 52: 273, 54: 421, 56: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 422, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273},
		{420, 274, 3: 274, 274, 274, 424, 423, 274, 274, 274, 419, 25: 274, 274, 49: 274, 274, 52: 274, 54: 421, 56: 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 422, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274},
		{420, 275, 3: 275, 275, 275, 424, 423, 275, 275, 275, 419, 25: 275, 275, 49: 275, 275, 52: 275, 54: 421, 56: 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 422, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275},
		{290, 290, 3: 290, 290, 290, 290, 290, 290, 290, 290, 290, 25: 290, 290, 49: 290, 290, 52: 290, 54: 290, 56: 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290},
		{291, 291, 3: 291, 291, 291, 291, 291, 291, 291, 291, 291, 25: 291, 291, 49: 291, 291, 52: 291, 54: 291, 56: 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291},
		// 180
		{420, 425, 3: 438, 428, 429, 424, 423, 305, 10: 305, 419, 50: 444, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{8: 302, 10: 497},
		{8: 496},
		{292, 292, 3: 292, 292, 292, 292, 292, 292, 292, 292, 292, 25: 292, 292, 49: 292, 292, 52: 292, 54: 292, 56: 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 498},
		// 185
		{420, 425, 3: 438, 428, 429, 424, 423, 304, 10: 304, 419, 50: 444, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{10: 470, 52: 500},
		{293, 293, 3: 293, 293, 293, 293, 293, 293, 293, 293, 293, 25: 293, 293, 49: 293, 293, 52: 293, 54: 293, 56: 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293},
		{153, 153, 153, 153, 153, 153, 153, 153, 153, 10: 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 43: 153, 52: 153},
		{129, 8: 129, 10: 129, 129},
		// 190
		{52: 507},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 505},
		{420, 425, 3: 438, 428, 429, 424, 423, 11: 419, 50: 444, 52: 506, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{126, 8: 126, 10: 126, 126},
		{128, 8: 128, 10: 128, 128},
		// 195
		{420, 283, 3: 283, 283, 283, 424, 423, 283, 283, 283, 419, 25: 283, 283, 49: 283, 283, 52: 283, 54: 421, 56: 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 422, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283},
		{125, 8: 125, 10: 125, 125},
		{8: 553},
		{8: 148, 22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 111: 323, 115: 342, 345, 341, 322, 122: 515, 324, 321, 156: 514, 169: 512, 513, 190: 516},
		{8: 150, 10: 550},
		// 200
		{8: 147},
		{8: 146, 10: 146},
		{156, 401, 156, 8: 132, 10: 132, 156, 133: 402, 518, 136: 519, 151: 405, 173: 520},
		{8: 517},
		{123, 8: 123, 10: 123, 123},
		// 205
		{523, 2: 522, 11: 120, 177: 406, 408, 521},
		{8: 144, 10: 144},
		{8: 143, 10: 143},
		{527, 8: 170, 170, 170, 526, 22: 170, 170, 170, 170, 27: 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 50: 170, 53: 170, 55: 170},
		{167, 8: 167, 167, 167, 167, 22: 167, 167, 167, 167, 27: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 50: 167, 53: 167, 55: 167},
		// 210
		{156, 401, 156, 8: 124, 11: 156, 22: 124, 124, 124, 27: 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 133: 402, 518, 136: 524, 151: 510, 171: 511},
		{8: 525},
		{166, 8: 166, 166, 166, 166, 22: 166, 166, 166, 166, 27: 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 50: 166, 53: 166, 55: 166},
		{152, 152, 152, 152, 152, 152, 152, 152, 12: 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 351, 349, 350, 43: 538, 52: 152, 111: 412, 149: 539, 158: 537},
		{2: 530, 8: 140, 22: 161, 161, 161, 27: 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 166: 531, 185: 529, 207: 528},
		// 215
		{22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 111: 323, 115: 342, 345, 341, 322, 122: 515, 324, 321, 156: 514, 169: 512, 535},
		{8: 534},
		{8: 142, 10: 142},
		{8: 139, 10: 532},
		{2: 533},
		// 220
		{8: 141, 10: 141},
		{159, 8: 159, 159, 159, 159, 22: 159, 159, 159, 159, 27: 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 159, 50: 159, 53: 159, 55: 159},
		{8: 536},
		{160, 8: 160, 160, 160, 160, 22: 160, 160, 160, 160, 27: 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 50: 160, 53: 160, 55: 160},
		{383, 546, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 410, 244, 165: 547},
		// 225
		{152, 152, 152, 152, 152, 152, 152, 152, 12: 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 351, 349, 350, 111: 412, 149: 416, 158: 543},
		{151, 151, 151, 151, 151, 151, 151, 151, 12: 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 351, 349, 350, 43: 540, 52: 151, 111: 501},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 541},
		{420, 425, 3: 438, 428, 429, 424, 423, 11: 419, 50: 444, 52: 542, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{163, 8: 163, 163, 163, 163, 22: 163, 163, 163, 163, 27: 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 50: 163, 53: 163, 55: 163},
		// 230
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 544},
		{420, 425, 3: 438, 428, 429, 424, 423, 11: 419, 50: 444, 52: 545, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{164, 8: 164, 164, 164, 164, 22: 164, 164, 164, 164, 27: 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 50: 164, 53: 164, 55: 164},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 508, 549},
		{52: 548},
		// 235
		{165, 8: 165, 165, 165, 165, 22: 165, 165, 165, 165, 27: 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 50: 165, 53: 165, 55: 165},
		{162, 8: 162, 162, 162, 162, 22: 162, 162, 162, 162, 27: 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 50: 162, 53: 162, 55: 162},
		{22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 111: 323, 115: 342, 345, 341, 322, 122: 515, 324, 321, 146: 551, 156: 552},
		{8: 149},
		{8: 145, 10: 145},
		// 240
		{130, 8: 130, 10: 130, 130},
		{8: 122, 22: 122, 122, 122, 27: 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 196: 555},
		{8: 148, 22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 111: 323, 115: 342, 345, 341, 322, 122: 515, 324, 321, 156: 514, 169: 512, 513, 190: 556},
		{8: 557},
		{121, 8: 121, 10: 121, 121},
		// 245
		{158, 401, 158, 8: 158, 10: 158, 158, 133: 559},
		{157, 2: 157, 8: 157, 10: 157, 157},
		{190, 190, 190, 8: 190, 190, 11: 190, 25: 190},
		{188, 188, 188, 8: 188, 188, 11: 188, 25: 188},
		{191, 191, 191, 8: 191, 191, 11: 191, 25: 191},
		// 250
		{245, 245, 3: 245, 245, 245, 245, 245, 245, 245, 245, 245, 25: 245, 245, 49: 245, 245, 52: 245, 54: 245, 56: 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 136, 136, 136, 27: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 51: 467, 112: 566, 150: 397, 157: 567},
		{420, 278, 3: 278, 278, 278, 424, 423, 278, 278, 278, 419, 25: 278, 278, 49: 278, 278, 52: 278, 54: 421, 56: 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 422, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278},
		{8: 594, 10: 470},
		{8: 568},
		// 255
		{383, 277, 376, 277, 277, 277, 386, 385, 277, 277, 277, 277, 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 25: 277, 277, 49: 277, 277, 569, 277, 570, 277, 56: 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277},
		{420, 276, 3: 276, 276, 276, 424, 423, 276, 276, 276, 419, 25: 276, 276, 49: 276, 276, 52: 276, 54: 421, 56: 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 422, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276},
		{113, 113, 113, 113, 113, 113, 113, 113, 11: 576, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 53: 113, 577, 155: 575, 162: 574, 572, 573, 189: 571},
		{10: 587, 26: 183, 161: 592},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 583, 53: 584, 168: 585},
		// 260
		{11: 576, 50: 581, 54: 577, 155: 582},
		{112, 112, 112, 112, 112, 112, 112, 112, 12: 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 53: 112},
		{11: 111, 50: 111, 54: 111},
		{238, 238, 238, 238, 238, 238, 238, 238, 12: 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 144: 374, 579},
		{2: 578},
		// 265
		{11: 108, 50: 108, 54: 108},
		{52: 580},
		{11: 109, 50: 109, 54: 109},
		{114, 114, 114, 114, 114, 114, 114, 114, 12: 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 53: 114},
		{11: 110, 50: 110, 54: 110},
		// 270
		{420, 425, 3: 438, 428, 429, 424, 423, 9: 118, 118, 419, 26: 118, 50: 444, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{113, 113, 113, 113, 113, 113, 113, 113, 11: 576, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 53: 113, 577, 155: 575, 162: 574, 572, 573, 189: 586},
		{10: 116, 26: 116},
		{10: 587, 26: 183, 161: 588},
		{113, 113, 113, 113, 113, 113, 113, 113, 11: 576, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 26: 182, 53: 113, 577, 155: 575, 162: 574, 590, 573},
		// 275
		{26: 589},
		{9: 117, 117, 26: 117},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 583, 53: 584, 168: 591},
		{10: 115, 26: 115},
		{26: 593},
		// 280
		{287, 287, 3: 287, 287, 287, 287, 287, 287, 287, 287, 287, 25: 287, 287, 49: 287, 287, 52: 287, 54: 287, 56: 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287},
		{294, 294, 3: 294, 294, 294, 294, 294, 294, 294, 294, 294, 25: 294, 294, 49: 294, 294, 52: 294, 54: 294, 56: 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294},
		{420, 279, 3: 279, 279, 279, 424, 423, 279, 279, 279, 419, 25: 279, 279, 49: 279, 279, 52: 279, 54: 421, 56: 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 422, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279},
		{420, 280, 3: 280, 280, 280, 424, 423, 280, 280, 280, 419, 25: 280, 280, 49: 280, 280, 52: 280, 54: 421, 56: 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 422, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280},
		{420, 281, 3: 281, 281, 281, 424, 423, 281, 281, 281, 419, 25: 281, 281, 49: 281, 281, 52: 281, 54: 421, 56: 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 422, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281},
		// 285
		{420, 282, 3: 282, 282, 282, 424, 423, 282, 282, 282, 419, 25: 282, 282, 49: 282, 282, 52: 282, 54: 421, 56: 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 422, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282},
		{420, 284, 3: 284, 284, 284, 424, 423, 284, 284, 284, 419, 25: 284, 284, 49: 284, 284, 52: 284, 54: 421, 56: 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 422, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284},
		{420, 285, 3: 285, 285, 285, 424, 423, 285, 285, 285, 419, 25: 285, 285, 49: 285, 285, 52: 285, 54: 421, 56: 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 422, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285},
		{420, 286, 3: 286, 286, 286, 424, 423, 286, 286, 286, 419, 25: 286, 286, 49: 286, 286, 52: 286, 54: 421, 56: 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 422, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286},
		{8: 603},
		// 290
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 569, 53: 570},
		{2: 369, 26: 182, 181: 372, 607},
		{26: 606},
		{180, 180, 180, 8: 180, 180, 180, 180, 22: 180, 180, 180, 180, 27: 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180},
		{10: 177, 26: 177},
		// 295
		{53: 201, 204: 610},
		{198, 198, 198, 8: 198, 198, 198, 198, 22: 198, 198, 198, 198, 27: 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 53: 137},
		{53: 611},
		{22: 200, 200, 200, 27: 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 205: 612},
		{22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 111: 399, 115: 342, 345, 341, 398, 148: 615, 192: 614, 226: 613},
		// 300
		{22: 351, 349, 350, 26: 629, 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 111: 399, 115: 342, 345, 341, 398, 148: 615, 192: 630},
		{22: 195, 195, 195, 26: 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195},
		{156, 401, 156, 9: 617, 25: 169, 133: 622, 621, 136: 619, 176: 620, 193: 618, 227: 616},
		{9: 626, 627},
		{22: 192, 192, 192, 26: 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192},
		// 305
		{9: 187, 187},
		{9: 185, 185, 25: 168},
		{25: 624},
		{623, 2: 522, 179: 521},
		{155, 2: 155},
		// 310
		{156, 401, 156, 133: 622, 621, 136: 524},
		{238, 238, 238, 238, 238, 238, 238, 238, 12: 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 144: 374, 625},
		{9: 184, 184},
		{22: 193, 193, 193, 26: 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193},
		{156, 401, 156, 25: 169, 133: 622, 621, 136: 619, 176: 620, 193: 628},
		// 315
		{9: 186, 186},
		{199, 199, 199, 8: 199, 199, 199, 199, 22: 199, 199, 199, 199, 27: 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199},
		{22: 194, 194, 194, 26: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 136, 136, 136, 27: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 51: 632, 150: 397, 157: 633},
		{420, 425, 3: 438, 428, 429, 424, 423, 635, 11: 419, 50: 444, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		// 320
		{8: 634},
		{202, 202, 202, 8: 202, 202, 202, 202, 22: 202, 202, 202, 202, 27: 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202},
		{203, 203, 203, 8: 203, 203, 203, 203, 22: 203, 203, 203, 203, 27: 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203},
		{232, 232, 232, 8: 232, 232, 232, 232},
		{230, 230, 230, 8: 230, 230, 230, 230},
		// 325
		{233, 233, 233, 8: 233, 233, 233, 233},
		{234, 234, 234, 8: 234, 234, 234, 234},
		{235, 235, 235, 8: 235, 235, 235, 235},
		{9: 781},
		{9: 229, 229},
		// 330
		{9: 226, 779},
		{9: 225, 225, 22: 63, 63, 63, 27: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 50: 224, 53: 64, 55: 64, 172: 645, 201: 647, 214: 646},
		{50: 777},
		{53: 72, 55: 72, 198: 653},
		{22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 111: 323, 115: 342, 345, 341, 322, 122: 648, 324, 321, 141: 649, 213: 650},
		// 335
		{156, 401, 156, 9: 227, 133: 622, 621, 136: 652, 167: 642, 187: 643, 641},
		{22: 66, 66, 66, 27: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 53: 66, 55: 66},
		{22: 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 53: 62, 55: 62, 111: 323, 115: 342, 345, 341, 322, 122: 648, 324, 321, 141: 651},
		{22: 65, 65, 65, 27: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 53: 65, 55: 65},
		{9: 225, 225, 50: 224, 172: 645},
		// 340
		{53: 70, 55: 68, 199: 655, 656, 221: 654},
		{22: 71, 71, 71, 27: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 55: 71},
		{53: 698, 126: 699},
		{55: 658, 121: 659, 125: 657},
		{9: 697},
		// 345
		{58, 22: 359, 84: 58, 194: 660},
		{49, 49, 49, 49, 49, 49, 49, 49, 9: 49, 12: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 26: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 53: 49, 55: 49, 84: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 110: 49},
		{661, 84: 662},
		{12: 361, 159: 687},
		{663},
		// 350
		{12: 361, 159: 664},
		{12: 363, 25: 665},
		{25: 666},
		{11: 670, 53, 152: 668, 667, 160: 669},
		{12: 683},
		// 355
		{8: 55, 10: 55, 25: 55},
		{10: 673, 25: 674},
		{2: 671},
		{52: 672},
		{12: 52},
		// 360
		{11: 670, 53, 152: 682, 667},
		{12: 675, 175: 676},
		{8: 51, 10: 51, 25: 51},
		{10: 677, 25: 678},
		{12: 681},
		// 365
		{2: 530, 166: 679},
		{8: 680, 10: 532},
		{45, 45, 45, 45, 45, 45, 45, 45, 9: 45, 12: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 26: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 53: 45, 55: 45, 84: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 110: 45},
		{8: 50, 10: 50, 25: 50},
		{8: 54, 10: 54, 25: 54},
		// 370
		{684},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 685},
		{420, 425, 3: 438, 428, 429, 424, 423, 686, 11: 419, 50: 444, 54: 421, 56: 427, 426, 432, 433, 443, 439, 440, 448, 441, 452, 422, 446, 436, 435, 434, 430, 450, 447, 445, 437, 454, 442, 431, 451, 449, 453},
		{8: 56, 10: 56, 25: 56},
		{8: 364, 12: 363, 25: 688},
		// 375
		{11: 670, 53, 152: 668, 667, 160: 689},
		{8: 690, 10: 673, 25: 691},
		{48, 48, 48, 48, 48, 48, 48, 48, 9: 48, 12: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 26: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 53: 48, 55: 48, 84: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 110: 48},
		{11: 670, 53, 152: 668, 667, 160: 692},
		{8: 693, 10: 673, 25: 694},
		// 380
		{47, 47, 47, 47, 47, 47, 47, 47, 9: 47, 12: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 26: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 53: 47, 55: 47, 84: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 110: 47},
		{12: 675, 175: 695},
		{8: 696, 10: 677},
		{46, 46, 46, 46, 46, 46, 46, 46, 9: 46, 12: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 26: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 53: 46, 55: 46, 84: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 110: 46},
		{22: 67, 67, 67, 27: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 55: 67},
		// 385
		{97, 97, 97, 97, 97, 97, 97, 97, 9: 97, 12: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 26: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 53: 97, 55: 97, 84: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 197: 700},
		{22: 69, 69, 69, 27: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 55: 69},
		{383, 388, 701, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 351, 349, 350, 26: 93, 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 51: 467, 53: 698, 55: 658, 84: 723, 720, 725, 710, 724, 711, 721, 722, 718, 726, 719, 111: 323, 702, 115: 342, 345, 341, 322, 717, 121: 659, 648, 324, 321, 709, 704, 705, 707, 708, 703, 706, 716, 141: 715, 174: 713, 210: 714, 712},
		{301, 301, 3: 301, 301, 301, 301, 301, 9: 301, 301, 301, 25: 775, 50: 301, 54: 301, 56: 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301},
		{8: 239, 239, 470},
		// 390
		{107, 107, 107, 107, 107, 107, 107, 107, 9: 107, 12: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 26: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 53: 107, 55: 107, 84: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 110: 107},
		{106, 106, 106, 106, 106, 106, 106, 106, 9: 106, 12: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 26: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 53: 106, 55: 106, 84: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 110: 106},
		{105, 105, 105, 105, 105, 105, 105, 105, 9: 105, 12: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 26: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 53: 105, 55: 105, 84: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 110: 105},
		{104, 104, 104, 104, 104, 104, 104, 104, 9: 104, 12: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 26: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 53: 104, 55: 104, 84: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 110: 104},
		{103, 103, 103, 103, 103, 103, 103, 103, 9: 103, 12: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 26: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 53: 103, 55: 103, 84: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 110: 103},
		// 395
		{102, 102, 102, 102, 102, 102, 102, 102, 9: 102, 12: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 26: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 53: 102, 55: 102, 84: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 110: 102},
		{101, 101, 101, 101, 101, 101, 101, 101, 9: 101, 12: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 26: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 53: 101, 55: 101, 84: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 110: 101},
		{238, 238, 238, 238, 238, 238, 238, 238, 12: 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 144: 374, 772},
		{25: 770},
		{26: 769},
		// 400
		{95, 95, 95, 95, 95, 95, 95, 95, 9: 95, 12: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 26: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 53: 95, 55: 95, 84: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95},
		{383, 388, 701, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 351, 349, 350, 26: 92, 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 51: 467, 53: 698, 55: 658, 84: 723, 720, 725, 710, 724, 711, 721, 722, 718, 726, 719, 111: 323, 702, 115: 342, 345, 341, 322, 717, 121: 659, 648, 324, 321, 709, 704, 705, 707, 708, 703, 706, 716, 141: 715, 174: 768},
		{91, 91, 91, 91, 91, 91, 91, 91, 9: 91, 12: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 26: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 53: 91, 55: 91, 84: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91},
		{90, 90, 90, 90, 90, 90, 90, 90, 9: 90, 12: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 26: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 53: 90, 55: 90, 84: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90},
		{9: 767},
		// 405
		{761},
		{757},
		{753},
		{383, 388, 701, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 53: 698, 55: 658, 84: 723, 720, 725, 710, 724, 711, 721, 722, 718, 726, 719, 112: 702, 119: 717, 121: 659, 125: 709, 704, 705, 707, 708, 703, 706, 747},
		{733},
		// 410
		{2: 731},
		{9: 730},
		{9: 729},
		{383, 388, 376, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 112: 702, 119: 727},
		{9: 728},
		// 415
		{78, 78, 78, 78, 78, 78, 78, 78, 9: 78, 12: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 26: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 53: 78, 55: 78, 84: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 110: 78},
		{79, 79, 79, 79, 79, 79, 79, 79, 9: 79, 12: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 26: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 53: 79, 55: 79, 84: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 110: 79},
		{80, 80, 80, 80, 80, 80, 80, 80, 9: 80, 12: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 26: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 53: 80, 55: 80, 84: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 110: 80},
		{9: 732},
		{81, 81, 81, 81, 81, 81, 81, 81, 9: 81, 12: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 26: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 53: 81, 55: 81, 84: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 110: 81},
		// 420
		{383, 388, 376, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 351, 349, 350, 27: 339, 331, 340, 336, 348, 335, 333, 334, 332, 337, 346, 343, 344, 347, 338, 330, 327, 328, 326, 352, 329, 325, 51: 467, 111: 323, 702, 115: 342, 345, 341, 322, 734, 122: 648, 324, 321, 141: 735},
		{9: 741},
		{383, 388, 376, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 112: 702, 119: 736},
		{9: 737},
		{383, 388, 376, 387, 389, 390, 386, 385, 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 112: 702, 119: 738},
		// 425
		{8: 739},
		{383, 388, 701, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 53: 698, 55: 658, 84: 723, 720, 725, 710, 724, 711, 721, 722, 718, 726, 719, 112: 702, 119: 717, 121: 659, 125: 709, 704, 705, 707, 708, 703, 706, 740},
		{82, 82, 82, 82, 82, 82, 82, 82, 9: 82, 12: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 26: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 53: 82, 55: 82, 84: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 110: 82},
		{383, 388, 376, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 112: 702, 119: 742},
		{9: 743},
		// 430
		{383, 388, 376, 387, 389, 390, 386, 385, 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 112: 702, 119: 744},
		{8: 745},
		{383, 388, 701, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 53: 698, 55: 658, 84: 723, 720, 725, 710, 724, 711, 721, 722, 718, 726, 719, 112: 702, 119: 717, 121: 659, 125: 709, 704, 705, 707, 708, 703, 706, 746},
		{83, 83, 83, 83, 83, 83, 83, 83, 9: 83, 12: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 26: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 53: 83, 55: 83, 84: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 110: 83},
		{85: 748},
		// 435
		{749},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 112: 750},
		{8: 751, 10: 470},
		{9: 752},
		{84, 84, 84, 84, 84, 84, 84, 84, 9: 84, 12: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 26: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 53: 84, 55: 84, 84: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 110: 84},
		// 440
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 112: 754},
		{8: 755, 10: 470},
		{383, 388, 701, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 53: 698, 55: 658, 84: 723, 720, 725, 710, 724, 711, 721, 722, 718, 726, 719, 112: 702, 119: 717, 121: 659, 125: 709, 704, 705, 707, 708, 703, 706, 756},
		{85, 85, 85, 85, 85, 85, 85, 85, 9: 85, 12: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 26: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 53: 85, 55: 85, 84: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 110: 85},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 112: 758},
		// 445
		{8: 759, 10: 470},
		{383, 388, 701, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 53: 698, 55: 658, 84: 723, 720, 725, 710, 724, 711, 721, 722, 718, 726, 719, 112: 702, 119: 717, 121: 659, 125: 709, 704, 705, 707, 708, 703, 706, 760},
		{86, 86, 86, 86, 86, 86, 86, 86, 9: 86, 12: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 26: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 53: 86, 55: 86, 84: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 110: 86},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 112: 762},
		{8: 763, 10: 470},
		// 450
		{383, 388, 701, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 53: 698, 55: 658, 84: 723, 720, 725, 710, 724, 711, 721, 722, 718, 726, 719, 112: 702, 119: 717, 121: 659, 125: 709, 704, 705, 707, 708, 703, 706, 764},
		{88, 88, 88, 88, 88, 88, 88, 88, 9: 88, 12: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 26: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 53: 88, 55: 88, 84: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 110: 765},
		{383, 388, 701, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 53: 698, 55: 658, 84: 723, 720, 725, 710, 724, 711, 721, 722, 718, 726, 719, 112: 702, 119: 717, 121: 659, 125: 709, 704, 705, 707, 708, 703, 706, 766},
		{87, 87, 87, 87, 87, 87, 87, 87, 9: 87, 12: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 26: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 53: 87, 55: 87, 84: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 110: 87},
		{89, 89, 89, 89, 89, 89, 89, 89, 9: 89, 12: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 26: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 53: 89, 55: 89, 84: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 110: 89},
		// 455
		{94, 94, 94, 94, 94, 94, 94, 94, 9: 94, 12: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 26: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 53: 94, 55: 94, 84: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94},
		{96, 96, 96, 96, 96, 96, 96, 96, 9: 96, 12: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 26: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 53: 96, 55: 96, 84: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 110: 96},
		{383, 388, 701, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 53: 698, 55: 658, 84: 723, 720, 725, 710, 724, 711, 721, 722, 718, 726, 719, 112: 702, 119: 717, 121: 659, 125: 709, 704, 705, 707, 708, 703, 706, 771},
		{98, 98, 98, 98, 98, 98, 98, 98, 9: 98, 12: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 26: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 53: 98, 55: 98, 84: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 110: 98},
		{25: 773},
		// 460
		{383, 388, 701, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 53: 698, 55: 658, 84: 723, 720, 725, 710, 724, 711, 721, 722, 718, 726, 719, 112: 702, 119: 717, 121: 659, 125: 709, 704, 705, 707, 708, 703, 706, 774},
		{99, 99, 99, 99, 99, 99, 99, 99, 9: 99, 12: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 26: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 53: 99, 55: 99, 84: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 110: 99},
		{383, 388, 701, 387, 389, 390, 386, 385, 9: 240, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 467, 53: 698, 55: 658, 84: 723, 720, 725, 710, 724, 711, 721, 722, 718, 726, 719, 112: 702, 119: 717, 121: 659, 125: 709, 704, 705, 707, 708, 703, 706, 776},
		{100, 100, 100, 100, 100, 100, 100, 100, 9: 100, 12: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 26: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 53: 100, 55: 100, 84: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 110: 100},
		{383, 388, 376, 387, 389, 390, 386, 385, 12: 382, 392, 391, 394, 377, 378, 379, 380, 381, 393, 51: 583, 53: 584, 168: 778},
		// 465
		{9: 223, 223},
		{156, 401, 156, 133: 622, 621, 136: 652, 167: 780},
		{9: 228, 228},
		{236, 236, 236, 236, 236, 236, 236, 236, 9: 236, 12: 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 26: 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 53: 236, 55: 236, 84: 236, 236, 236, 236, 236, 236, 236, 236, 236, 236, 236},
		{22: 76, 76, 76, 27: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 55: 76},
		// 470
		{238, 238, 238, 238, 238, 238, 238, 238, 12: 238, 238, 238, 238, 238, 238, 238, 238, 238, 238, 144: 374, 784},
		{49: 309},
		{82: 807, 809, 98: 797, 798, 799, 794, 795, 796, 800, 804, 801, 791, 802, 803, 113: 808, 806, 120: 805, 135: 789, 137: 788, 793, 790, 792, 142: 787, 224: 786},
		{49: 311},
		{49: 44, 82: 807, 809, 98: 797, 798, 799, 794, 795, 796, 800, 804, 801, 791, 802, 803, 113: 808, 806, 120: 805, 135: 789, 137: 852, 793, 790, 792},
		// 475
		{49: 43, 82: 43, 43, 95: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43},
		{49: 39, 82: 39, 39, 95: 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39},
		{49: 38, 82: 38, 38, 95: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{83: 809, 113: 874, 806},
		{49: 36, 82: 36, 36, 95: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
		// 480
		{95: 29, 29, 862, 180: 860, 215: 861, 859},
		{83: 809, 113: 856, 806},
		{2: 853},
		{2: 848},
		{2: 824, 82: 826, 222: 825},
		// 485
		{82: 807, 809, 113: 808, 806, 120: 823},
		{49: 17, 82: 17, 17, 95: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{83: 809, 113: 821, 806},
		{83: 809, 113: 819, 806},
		{82: 807, 809, 113: 808, 806, 120: 818},
		// 490
		{2: 814},
		{83: 809, 113: 812, 806},
		{49: 7, 82: 7, 7, 95: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{82: 5, 811},
		{49: 4, 82: 4, 4, 95: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		// 495
		{82: 810},
		{82: 2, 2},
		{49: 3, 82: 3, 3, 95: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{82: 1, 1},
		{82: 813},
		// 500
		{49: 8, 82: 8, 8, 95: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{82: 815, 809, 113: 816, 806},
		{49: 13, 82: 13, 13, 95: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{82: 817},
		{49: 9, 82: 9, 9, 95: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		// 505
		{49: 14, 82: 14, 14, 95: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{82: 820},
		{49: 15, 82: 15, 15, 95: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{82: 822},
		{49: 16, 82: 16, 16, 95: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		// 510
		{49: 18, 82: 18, 18, 95: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{82: 807, 809, 113: 808, 806, 120: 833, 143: 847},
		{2: 827, 8: 140, 146: 829, 166: 828, 185: 830},
		{49: 10, 82: 10, 10, 95: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{8: 142, 10: 142, 146: 844},
		// 515
		{8: 139, 10: 836},
		{8: 834},
		{8: 831},
		{82: 807, 809, 113: 808, 806, 120: 833, 143: 832},
		{49: 19, 82: 19, 19, 95: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		// 520
		{49: 6, 82: 6, 6, 95: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{82: 807, 809, 113: 808, 806, 120: 833, 143: 835},
		{49: 21, 82: 21, 21, 95: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{2: 837, 146: 838},
		{8: 141, 10: 141, 146: 841},
		// 525
		{8: 839},
		{82: 807, 809, 113: 808, 806, 120: 833, 143: 840},
		{49: 20, 82: 20, 20, 95: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{8: 842},
		{82: 807, 809, 113: 808, 806, 120: 833, 143: 843},
		// 530
		{49: 11, 82: 11, 11, 95: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{8: 845},
		{82: 807, 809, 113: 808, 806, 120: 833, 143: 846},
		{49: 12, 82: 12, 12, 95: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{49: 22, 82: 22, 22, 95: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		// 535
		{82: 849},
		{82: 807, 809, 95: 41, 41, 41, 797, 798, 799, 794, 795, 796, 800, 804, 801, 791, 802, 803, 113: 808, 806, 120: 805, 135: 789, 137: 788, 793, 790, 792, 142: 850, 147: 851},
		{82: 807, 809, 95: 40, 40, 40, 797, 798, 799, 794, 795, 796, 800, 804, 801, 791, 802, 803, 113: 808, 806, 120: 805, 135: 789, 137: 852, 793, 790, 792},
		{95: 32, 32, 32},
		{49: 42, 82: 42, 42, 95: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		// 540
		{82: 854},
		{82: 807, 809, 95: 41, 41, 41, 797, 798, 799, 794, 795, 796, 800, 804, 801, 791, 802, 803, 113: 808, 806, 120: 805, 135: 789, 137: 788, 793, 790, 792, 142: 850, 147: 855},
		{95: 33, 33, 33},
		{82: 857},
		{82: 807, 809, 95: 41, 41, 41, 797, 798, 799, 794, 795, 796, 800, 804, 801, 791, 802, 803, 113: 808, 806, 120: 805, 135: 789, 137: 788, 793, 790, 792, 142: 850, 147: 858},
		// 545
		{95: 34, 34, 34},
		{95: 25, 868, 217: 869, 867},
		{95: 31, 31, 31},
		{95: 28, 28, 862, 180: 866},
		{83: 809, 113: 863, 806},
		// 550
		{82: 864},
		{82: 807, 809, 95: 41, 41, 41, 797, 798, 799, 794, 795, 796, 800, 804, 801, 791, 802, 803, 113: 808, 806, 120: 805, 135: 789, 137: 788, 793, 790, 792, 142: 850, 147: 865},
		{95: 27, 27, 27},
		{95: 30, 30, 30},
		{95: 873, 219: 872},
		// 555
		{82: 870},
		{95: 24},
		{82: 807, 809, 95: 41, 98: 797, 798, 799, 794, 795, 796, 800, 804, 801, 791, 802, 803, 113: 808, 806, 120: 805, 135: 789, 137: 788, 793, 790, 792, 142: 850, 147: 871},
		{95: 26},
		{49: 35, 82: 35, 35, 95: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		// 560
		{49: 23, 82: 23, 23, 95: 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23},
		{82: 875},
		{49: 37, 82: 37, 37, 95: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
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
	const yyError = 232

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
				BasicAssemblerStatement: yyS[yypt-0].node.(*BasicAssemblerStatement),
			}
		}
	case 241:
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
	case 242:
		{
			yyVAL.node = &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:          yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 243:
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
	case 244:
		{
			lhs := &FunctionBody{
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
			yyVAL.node = lhs
			lhs.scope = lhs.CompoundStatement.scope
		}
	case 245:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 246:
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
	case 247:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 248:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 249:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 250:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 251:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScopePos(lhs.Pos())
		}
	case 252:
		{
			yyVAL.node = &AssemblerInstructions{
				Token: yyS[yypt-0].Token,
			}
		}
	case 253:
		{
			yyVAL.node = &AssemblerInstructions{
				Case: 1,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions),
				Token: yyS[yypt-0].Token,
			}
		}
	case 254:
		{
			yyVAL.node = &BasicAssemblerStatement{
				Token:                 yyS[yypt-4].Token,
				VolatileOpt:           yyS[yypt-3].node.(*VolatileOpt),
				Token2:                yyS[yypt-2].Token,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-0].Token,
			}
		}
	case 255:
		{
			yyVAL.node = (*VolatileOpt)(nil)
		}
	case 256:
		{
			yyVAL.node = &VolatileOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 257:
		{
			yyVAL.node = &AssemblerOperand{
				AssemblerSymbolicNameOpt: yyS[yypt-4].node.(*AssemblerSymbolicNameOpt),
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
		}
	case 258:
		{
			yyVAL.node = &AssemblerOperands{
				AssemblerOperand: yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 259:
		{
			yyVAL.node = &AssemblerOperands{
				Case:              1,
				AssemblerOperands: yyS[yypt-2].node.(*AssemblerOperands),
				Token:             yyS[yypt-1].Token,
				AssemblerOperand:  yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 260:
		{
			yyVAL.node = (*AssemblerSymbolicNameOpt)(nil)
		}
	case 261:
		{
			yyVAL.node = &AssemblerSymbolicNameOpt{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 262:
		{
			yyVAL.node = &Clobbers{
				Token: yyS[yypt-0].Token,
			}
		}
	case 263:
		{
			yyVAL.node = &Clobbers{
				Case:     1,
				Clobbers: yyS[yypt-2].node.(*Clobbers),
				Token:    yyS[yypt-1].Token,
				Token2:   yyS[yypt-0].Token,
			}
		}
	case 264:
		{
			yyVAL.node = &AssemblerStatement{
				BasicAssemblerStatement: yyS[yypt-0].node.(*BasicAssemblerStatement),
			}
		}
	case 265:
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
	case 266:
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
	case 267:
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
	case 268:
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
	case 269:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 270:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 271:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 272:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 273:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 274:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 275:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 276:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 277:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 278:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 279:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 280:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 281:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 282:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 283:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 284:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 285:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 286:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 287:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 288:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 289:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 290:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 291:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 292:
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
	case 293:
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
	case 294:
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
	case 295:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 296:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 297:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 298:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 299:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 300:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 301:
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
	case 302:
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
	case 303:
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
	case 304:
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
	case 305:
		{
			yyVAL.node = &ControlLine{
				Case:        14,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 308:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 309:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
