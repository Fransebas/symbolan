/** Taken from "The Definitive ANTLR 4 Reference" by Terence Parr */

grammar TreeLan;


ruleSet 
	: (treeRule NEWLINE)+
	;

treeRule
	: treeDescriptor ASSIGN treeDescriptor
	| treeDescriptor ASSIGN 'F_' formula
	;

formula
	: ID
	;

treeDescriptor 
	: '(' treeDescriptor ')' rootNode '(' treeDescriptor ')'
	| leafNode
	;

rootNode 
	: operationClass '_' ID
	;

leafNode
	: classValue '_' ID
	;

classValue 
	:  NUMERIC_CONSTANT            
	| CONSTANT                    
	| VARIABLE                    
	| CONSTANT_EXPRESSION         
	| VARIABLE_EXPRESSION         
	| NUMERIC_CONSTANT_EXPRESSION 
	| SYSTEM_FUNCTION 
	| ANY 
	| NUMBER   
	| SPECIAL_NUMBER 
	;       

operationClass 
    : ADDITION
    | SUBTRACTION
    | MULTIPLICATION
    | DIVISION
    | EXPONENTIAL
    | LOGARITHM
    | SIN
    | COS
    | TAN
    | TRANSFORMATION
    | EQUATION
    | LEAF
    | ANY
    ;

ADDITION 		: 'ADDITION';
SUBTRACTION 	: 'SUBTRACTION';
MULTIPLICATION 	: 'MULTIPLICATION';
DIVISION 		: 'DIVISION';
EXPONENTIAL 	: 'EXPONENTIAL';
LOGARITHM 		: 'LOGARITHM';
SIN 			: 'SIN';
COS 			: 'COS';
TAN 			: 'TAN';
TRANSFORMATION 	: 'TRANSFORMATION';
EQUATION 		: 'EQUATION';
LEAF 			: 'LEAF';
ANY 			: 'ANY';

SPECIAL_NUMBER 				: SP_NUMBER;
NUMERIC_CONSTANT            : 'NUMERIC_CONSTANT';
CONSTANT                    : 'CONSTANT';
VARIABLE                    : 'VARIABLE';
CONSTANT_EXPRESSION         : 'CONSTANT_EXPRESSION';
VARIABLE_EXPRESSION         : 'VARIABLE_EXPRESSION';
NUMERIC_CONSTANT_EXPRESSION : 'NUMERIC_CONSTANT_EXPRESSION';
SYSTEM_FUNCTION             : 'SYSTEM_FUNCTION';

ID : [a-zA-Z]+;

ASSIGN : '=>';

NUMBER 
	: INT
	;

fragment SP_NUMBER
	: 'e' | 'pi'
	;


fragment INT
   : '0' | [1-9] [0-9]*
   ;

// no leading zeros

fragment EXP
   : [Ee] [+\-]? INT // \- since - means "range" inside [...]
   ; 

NEWLINE:'\r'? '\n' ;
WS : [ \t]+ -> skip ;