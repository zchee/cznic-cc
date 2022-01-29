%{
package cc
%}

%union {
	Token			Token
	node			Node
}

%token
	/*yy:token "%c"		*/	IDENTIFIER
	/*yy:token "%c_e"	*/	ENUMCONST
	/*yy:token "%c_t"	*/	TYPENAME
	/*yy:token "%d"		*/	INTCONST
	/*yy:token "'%c'"	*/	CHARCONST
	/*yy:token "1.%d"	*/	FLOATCONST
	/*yy:token "L'%c'"	*/	LONGCHARCONST
	/*yy:token "L\"%c\""	*/	LONGSTRINGLITERAL
	/*yy:token "\"%c\""	*/	STRINGLITERAL

	ACCUM			"_Accum"
	ADDASSIGN		"+="
	ALIGNAS			"_Alignas"
	ALIGNOF			"_Alignof"
	ANDAND			"&&"
	ANDASSIGN		"&="
	ARROW			"->"
	ASM			"asm"
	ATOMIC			"_Atomic"
	ATTRIBUTE		"__attribute__"
	AUTO			"auto"
	BOOL			"_Bool"
	BREAK			"break"
	BUILTINCHOOSEXPR	"__builtin_choose_expr"
	BUILTINTYPESCOMPATIBLE	"__builtin_types_compatible_p"
	CASE			"case"
	CHAR			"char"
	COMPLEX			"_Complex"
	CONST			"const"
	CONTINUE		"continue"
	DDD			"..."
	DEC			"--"
	DECIMAL128		"_Decimal128"
	DECIMAL32		"_Decimal32"
	DECIMAL64		"_Decimal64"
	DEFAULT			"default"
	DIVASSIGN		"/="
	DO			"do"
	DOUBLE			"double"
	ELSE			"else"
	ENUM			"enum"
	EQ			"=="
	EXTERN			"extern"
	FLOAT			"float"
	FLOAT128		"_Float128"
	FLOAT16			"__fp16"
	FLOAT32			"_Float32"
	FLOAT32X		"_Float32x"
	FLOAT64			"_Float64"
	FLOAT64x		"_Float64x"
	FLOAT80			"__float80"
	FOR			"for"
	FRACT			"_Fract"
	GEQ			">="
	GOTO			"goto"
	IF			"if"
	IMAG			"__imag__"
	INC			"++"
	INLINE			"inline"
	INT			"int"
	INT8			"__int8"
	INT16			"__int16"
	INT32			"__int32"
	INT64			"__int64"
	INT128			"__int128"
	LABEL			"__label__"
	LEQ			"<="
	LONG			"long"
	LSH			"<<"
	LSHASSIGN		"<<="
	MODASSIGN		"%="
	MULASSIGN		"*="
	NEQ			"!="
	NORETURN		"_Noreturn"
	ORASSIGN		"|="
	OROR			"||"
	PPNUMBER		"preprocessing number"
	PPPASTE			"##"
	PRAGMASTDC		"__pragma_stdc"
	REAL			"__real__"
	REGISTER		"register"
	RESTRICT		"restrict"
	RETURN			"return"
	RSH			">>"
	RSHASSIGN		">>="
	SAT			"_Sat"
	SHORT			"short"
	SIGNED			"signed"
	SIZEOF			"sizeof"
	STATIC			"static"
	STRUCT			"struct"
	SUBASSIGN		"-="
	SWITCH			"switch"
	THREADLOCAL		"_Thread_local"
	TYPEDEF			"typedef"
	TYPEOF			"typeof"
	UNION			"union"
	UNSIGNED		"unsigned"
	VOID			"void"
	VOLATILE		"volatile"
	WHILE			"while"
	XORASSIGN		"^="


%%

ParameterDeclaration:
	DeclarationSpecifiers
|	DeclarationSpecifiers AbstractDeclarator
|	DeclarationSpecifiers Declarator

Declarator:
	DirectDeclarator
|	Pointer DirectDeclarator

Pointer:
	'*'
|	'*' Pointer

DirectDeclarator:
	IDENTIFIER
|	'(' Declarator ')'
|	DirectDeclarator '(' ')'
|	DirectDeclarator '(' IdentifierList ')'
|	DirectDeclarator '(' ParameterTypeList ')'
|	DirectDeclarator '[' "static" ']'

ParameterTypeList:
	ParameterList
|	ParameterList ',' "..."

ParameterList:
	ParameterDeclaration
|	ParameterList ',' ParameterDeclaration

IdentifierList:
	IDENTIFIER
|	IdentifierList ',' IDENTIFIER

DeclarationSpecifiers:
	"int"

AbstractDeclarator:
	Pointer
|	Pointer DirectAbstractDeclarator
|	DirectAbstractDeclarator

DirectAbstractDeclarator:
	'(' AbstractDeclarator ')'
|	'(' ')'
|	'(' ParameterTypeList ')'
|	'[' "static" ']'
|	DirectAbstractDeclarator '(' ')'
|	DirectAbstractDeclarator '(' ParameterTypeList ')'
|	DirectAbstractDeclarator '[' "static" ']'
