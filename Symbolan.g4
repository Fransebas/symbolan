/** Taken from "The Definitive ANTLR 4 Reference" by Terence Parr */

// Derived from http://json.org
grammar Symbolan;

// Parser i.e. Rules

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
    ;

expr 
    : expr expr
    | atom
    | left_par expr right_par
    | expr ('^'|'**') expr 
    | expr ('*'|'/') expr 
    | expr ('+'|'-') expr
    | function
    ;

function 
    : system_functions left_par expr right_par
    ;

left_par
    : '('
    ;

right_par
    : ')'
    ;

atom 
    : number
    | variable
    | constant
    | number
    ;

number 
    : DECIMAL_LIT 
    | FLOAT_LIT
    | IMAGINARY_LIT
    | OCTAL_LIT
    | HEX_LIT
    | PI
    | EULER
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

MIN                  : 'min';
MAX                  : 'max';
SIN                  : 'sin';
COS                  : 'cos';
TAN                  : 'tan';
LOG                  : 'log';


PI                  : 'pi';

VAR : [x-z] (('_')[0-9]+)? ;

CONST : [a-c] (('_')[0-9]+)? ;

EULER : [e-e] ;

IMAGINARY  : (DECIMALS | FLOAT_LIT) 'i';

// Numbers

DECIMAL_LIT            : [1-9] [0-9]*;
OCTAL_LIT              : '0' OCTAL_DIGIT*;
HEX_LIT                : '0' [xX] HEX_DIGIT+;

FLOAT_LIT              
    : DECIMALS ('.' DECIMALS? EXPONENT? | EXPONENT) 
    | '.' DECIMALS EXPONENT?
    ;

IMAGINARY_LIT
    : (DECIMALS | FLOAT_LIT) 'i'
    ;

//IDENTIFIER             : [d-zA-Z] [a-zA-Z]+ ;

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

WS
   : [ \t\n\r] + -> skip
   ;
   
LINE_COMMENT           : '//' ~[\r\n]*      -> channel(HIDDEN);

