/** Taken from "The Definitive ANTLR 4 Reference" by Terence Parr */

// Derived from http://json.org
grammar Symbolan;

// Parser i.e. Rules

ruleSet 
	: (expressionRule NEWLINE*)*
	;

eq  
    : expr
    | expr '=' expr
    ;

system_functions
    : MIN
    | MAX
    | SIN
    | COS
    | TAN
    | LOG
    | SINH
	| COSH
	| TANH
	| COTH
	| CSCH
	| SECH
    | DERIVATIVE_OPERATOR
    ;

	

rule_function
	: IDENTIFIER '()'
	;

expressionRule 
	: expr '=>' expr
	| expr '=>' rule_function
	;

expr:<assoc=right>  expr '^'expr 
	| atom atom

    | expr ('*'|'/') expr

    | expr ('+'|'-') expr
    | left_par expr right_par
    | system_functions left_par expr right_par
    | atom
    | ruleAtom
    | sign expr 
    ;

//
//
//
//expr 
//	: left_par expr right_par
//    | expr ('^'|'**') expr 
//    | expr ('*'|'/') expr 
//    | expr ('+'|'-') expr
//    | expr expr
//    | system_functions left_par expr right_par
//    | sign atom
//    | sign ruleAtom
//    | sign left_par expr right_par
//    | atom
//    | ruleAtom
//    ;

ruleAtom 
	: const_rule
	| var_rule
	| expr_rule
	| numeric_rule
	| derivative_rule
	| numeric_tree_rule
	| const_tree_rule
	| var_tree_rule
	| system_tree_rule
	;

sign: '+' | '-';
left_par: '(';
right_par: ')';


derivative_rule: DERIVATE_RULE; 
numeric_rule: NUMERIC_RULE;
const_rule: CONST_RULE;
var_rule: VAR_RULE;
expr_rule: EXPR_RULE;

numeric_tree_rule: NUMERIC_TREE_RULE;
const_tree_rule: CONST_TREE_RULE;
var_tree_rule: VAR_TREE_RULE;
system_tree_rule: SYSTEM_TREE_RULE;


atom 
    : number
    | variable
    | constant
    | number
    | math_constant
    | imaginary
    | derivative
    ;

math_constant 
	: PI
    | EULER
    | INFINITE
    | UNDEFINED
    | IMAGINARY
    ;

infinite
	: INFINITE
	;

imaginary
	: IMAGINARY_LIT
	;

number 
    : DECIMAL_LIT 
    | FLOAT_LIT
    | OCTAL_LIT
    | HEX_LIT
    ;

derivative
	: DERIVATIVE
	;

variable 
    : VAR
    ;

constant 
    : CONST 
    ;

bi_operators
    : ('^'|'**') 
    | ('*'|'/') 
    | ('+'|'-')  
    ;

// Lexer i.e. Tokens


L_PAREN                : '(';
R_PAREN                : ')';
L_CURLY                : '{';
R_CURLY                : '}';
L_BRACKET              : '[';
R_BRACKET              : ']';
ASSIGN                 : '=';
COMMA                  : ',';

// operators

FACTORIAL              : '!';

// Keywords
SIN                  : 'sin';
COS                  : 'cos';
TAN                  : 'tan';
COT                  : 'cot'; 
CSC                  : 'csc';
SEC                  : 'sec';

SINH                  : 'sinh';
COSH                  : 'cosh';
TANH                  : 'tanh';
COTH                  : 'coth'; 
CSCH                  : 'csch';
SECH                  : 'sech'; 

MIN                  : 'min';
MAX                  : 'max';
LOG                  : 'log';
SQRT                 : 'sqrt';
DERIVATIVE_OPERATOR  : 'D';


PI                  : 'pi';

DERIVATIVE : 'd' [x-z] (('_')[0-9]+)? ;
VAR : [x-z] (('_')[0-9]+)? ;
CONST : [a-c] (('_')[0-9]+)? ;
EULER : [e-e] ;
INFINITE : [+-]?'inf';
UNDEFINED : 'undef';

DERIVATE_RULE	: 'd'   [V] (('_')[0-9]+)? ;
NUMERIC_RULE 	: [N] (('_')[0-9]+)? ;
CONST_RULE 		: [C] (('_')[0-9]+)? ;
VAR_RULE 		: [V] (('_')[0-9]+)? ;
EXPR_RULE 		: [F] (('_')[0-9]+)? ;

NUMERIC_TREE_RULE	 	: ('!')? 'F' [N] (('_')[0-9]+)? ;
CONST_TREE_RULE 		: ('!')? 'F' [C] (('_')[0-9]+)? ;
VAR_TREE_RULE 			: ('!')? 'F' [V] (('_')[0-9]+)? ;
SYSTEM_TREE_RULE 		: ('!')? 'SF'    (('_')[0-9]+)? ;

IMAGINARY  : 'i';

// Numbers

//DECIMAL_LIT            : [+-]? [1-9] [0-9]*;
DECIMAL_LIT            :  [1-9] [0-9]*;
OCTAL_LIT              : '0' OCTAL_DIGIT*;
HEX_LIT                : '0' [xX] HEX_DIGIT+;

FLOAT_LIT              
    : DECIMALS ('.' DECIMALS? EXPONENT? | EXPONENT) 
    | '.' DECIMALS EXPONENT?
    ;

IMAGINARY_LIT
    : (DECIMALS | FLOAT_LIT) 'i'
    ;

IDENTIFIER
   : [A-Z] [a-zA-Z0-9_]*
   ;

// Fragments

fragment DECIMALS
    : [0-9]+
    ;
fragment OCTAL_DIGIT
    : [0-7]
    ;
fragment HEX_DIGIT
    : [0-9a-fA-F]
    ;
fragment EXPONENT
    : [eE] [+-]? DECIMALS
    ;

// \- since - means "range" inside [...]

NEWLINE:'\r'? '\n' ;
WS : [ \t]+ -> skip ;
   
LINE_COMMENT           : '//' ~[\r\n]*      -> channel(HIDDEN);

