// Code generated from Symbolan.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Symbolan

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 169,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 7, 2, 53, 10, 2, 12, 2, 14, 2,
	56, 11, 2, 7, 2, 58, 10, 2, 12, 2, 14, 2, 61, 11, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 3, 68, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 83, 10, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	5, 7, 100, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 7, 7, 113, 10, 7, 12, 7, 14, 7, 116, 11, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 5, 8, 123, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 148, 10, 17, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 5, 25, 167, 10, 25, 3, 25, 2, 3,
	12, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 2, 8, 4, 2, 20, 25, 27, 27, 3, 2, 5, 6, 3,
	2, 7, 8, 3, 2, 9, 10, 5, 2, 28, 28, 32, 34, 40, 40, 3, 2, 41, 44, 2, 168,
	2, 59, 3, 2, 2, 2, 4, 67, 3, 2, 2, 2, 6, 69, 3, 2, 2, 2, 8, 71, 3, 2, 2,
	2, 10, 82, 3, 2, 2, 2, 12, 99, 3, 2, 2, 2, 14, 122, 3, 2, 2, 2, 16, 124,
	3, 2, 2, 2, 18, 126, 3, 2, 2, 2, 20, 128, 3, 2, 2, 2, 22, 130, 3, 2, 2,
	2, 24, 132, 3, 2, 2, 2, 26, 134, 3, 2, 2, 2, 28, 136, 3, 2, 2, 2, 30, 138,
	3, 2, 2, 2, 32, 147, 3, 2, 2, 2, 34, 149, 3, 2, 2, 2, 36, 151, 3, 2, 2,
	2, 38, 153, 3, 2, 2, 2, 40, 155, 3, 2, 2, 2, 42, 157, 3, 2, 2, 2, 44, 159,
	3, 2, 2, 2, 46, 161, 3, 2, 2, 2, 48, 166, 3, 2, 2, 2, 50, 54, 5, 10, 6,
	2, 51, 53, 7, 47, 2, 2, 52, 51, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2, 54, 52,
	3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2,
	57, 50, 3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3,
	2, 2, 2, 60, 3, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 68, 5, 12, 7, 2, 63,
	64, 5, 12, 7, 2, 64, 65, 7, 17, 2, 2, 65, 66, 5, 12, 7, 2, 66, 68, 3, 2,
	2, 2, 67, 62, 3, 2, 2, 2, 67, 63, 3, 2, 2, 2, 68, 5, 3, 2, 2, 2, 69, 70,
	9, 2, 2, 2, 70, 7, 3, 2, 2, 2, 71, 72, 7, 46, 2, 2, 72, 73, 7, 3, 2, 2,
	73, 9, 3, 2, 2, 2, 74, 75, 5, 12, 7, 2, 75, 76, 7, 4, 2, 2, 76, 77, 5,
	12, 7, 2, 77, 83, 3, 2, 2, 2, 78, 79, 5, 12, 7, 2, 79, 80, 7, 4, 2, 2,
	80, 81, 5, 8, 5, 2, 81, 83, 3, 2, 2, 2, 82, 74, 3, 2, 2, 2, 82, 78, 3,
	2, 2, 2, 83, 11, 3, 2, 2, 2, 84, 85, 8, 7, 1, 2, 85, 86, 5, 18, 10, 2,
	86, 87, 5, 12, 7, 2, 87, 88, 5, 20, 11, 2, 88, 100, 3, 2, 2, 2, 89, 90,
	5, 6, 4, 2, 90, 91, 5, 18, 10, 2, 91, 92, 5, 12, 7, 2, 92, 93, 5, 20, 11,
	2, 93, 100, 3, 2, 2, 2, 94, 95, 5, 16, 9, 2, 95, 96, 5, 12, 7, 5, 96, 100,
	3, 2, 2, 2, 97, 100, 5, 32, 17, 2, 98, 100, 5, 14, 8, 2, 99, 84, 3, 2,
	2, 2, 99, 89, 3, 2, 2, 2, 99, 94, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 98,
	3, 2, 2, 2, 100, 114, 3, 2, 2, 2, 101, 102, 12, 10, 2, 2, 102, 103, 9,
	3, 2, 2, 103, 113, 5, 12, 7, 11, 104, 105, 12, 9, 2, 2, 105, 106, 9, 4,
	2, 2, 106, 113, 5, 12, 7, 10, 107, 108, 12, 8, 2, 2, 108, 109, 9, 5, 2,
	2, 109, 113, 5, 12, 7, 9, 110, 111, 12, 7, 2, 2, 111, 113, 5, 12, 7, 8,
	112, 101, 3, 2, 2, 2, 112, 104, 3, 2, 2, 2, 112, 107, 3, 2, 2, 2, 112,
	110, 3, 2, 2, 2, 113, 116, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114, 115,
	3, 2, 2, 2, 115, 13, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 117, 123, 5, 26,
	14, 2, 118, 123, 5, 28, 15, 2, 119, 123, 5, 30, 16, 2, 120, 123, 5, 24,
	13, 2, 121, 123, 5, 22, 12, 2, 122, 117, 3, 2, 2, 2, 122, 118, 3, 2, 2,
	2, 122, 119, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 121, 3, 2, 2, 2, 123,
	15, 3, 2, 2, 2, 124, 125, 9, 5, 2, 2, 125, 17, 3, 2, 2, 2, 126, 127, 7,
	11, 2, 2, 127, 19, 3, 2, 2, 2, 128, 129, 7, 12, 2, 2, 129, 21, 3, 2, 2,
	2, 130, 131, 7, 35, 2, 2, 131, 23, 3, 2, 2, 2, 132, 133, 7, 36, 2, 2, 133,
	25, 3, 2, 2, 2, 134, 135, 7, 37, 2, 2, 135, 27, 3, 2, 2, 2, 136, 137, 7,
	38, 2, 2, 137, 29, 3, 2, 2, 2, 138, 139, 7, 39, 2, 2, 139, 31, 3, 2, 2,
	2, 140, 148, 5, 40, 21, 2, 141, 148, 5, 44, 23, 2, 142, 148, 5, 46, 24,
	2, 143, 148, 5, 40, 21, 2, 144, 148, 5, 34, 18, 2, 145, 148, 5, 38, 20,
	2, 146, 148, 5, 42, 22, 2, 147, 140, 3, 2, 2, 2, 147, 141, 3, 2, 2, 2,
	147, 142, 3, 2, 2, 2, 147, 143, 3, 2, 2, 2, 147, 144, 3, 2, 2, 2, 147,
	145, 3, 2, 2, 2, 147, 146, 3, 2, 2, 2, 148, 33, 3, 2, 2, 2, 149, 150, 9,
	6, 2, 2, 150, 35, 3, 2, 2, 2, 151, 152, 7, 33, 2, 2, 152, 37, 3, 2, 2,
	2, 153, 154, 7, 45, 2, 2, 154, 39, 3, 2, 2, 2, 155, 156, 9, 7, 2, 2, 156,
	41, 3, 2, 2, 2, 157, 158, 7, 29, 2, 2, 158, 43, 3, 2, 2, 2, 159, 160, 7,
	30, 2, 2, 160, 45, 3, 2, 2, 2, 161, 162, 7, 31, 2, 2, 162, 47, 3, 2, 2,
	2, 163, 167, 9, 3, 2, 2, 164, 167, 9, 4, 2, 2, 165, 167, 9, 5, 2, 2, 166,
	163, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166, 165, 3, 2, 2, 2, 167, 49, 3,
	2, 2, 2, 12, 54, 59, 67, 82, 99, 112, 114, 122, 147, 166,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'()'", "'=>'", "'^'", "'**'", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
	"'{'", "'}'", "'['", "']'", "'='", "','", "'!'", "'min'", "'max'", "'sin'",
	"'cos'", "'tan'", "'log'", "'sqrt'", "'D'", "'pi'", "", "", "", "", "",
	"'undef'", "", "", "", "", "", "'i'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY",
	"L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "FACTORIAL", "MIN", "MAX",
	"SIN", "COS", "TAN", "LOG", "SQRT", "DERIVATIVE_OPERATOR", "PI", "DERIVATIVE",
	"VAR", "CONST", "EULER", "INFINITE", "UNDEFINED", "DERIVATE_RULE", "NUMERIC_RULE",
	"CONST_RULE", "VAR_RULE", "EXPR_RULE", "IMAGINARY", "DECIMAL_LIT", "OCTAL_LIT",
	"HEX_LIT", "FLOAT_LIT", "IMAGINARY_LIT", "IDENTIFIER", "NEWLINE", "WS",
	"LINE_COMMENT",
}

var ruleNames = []string{
	"ruleSet", "eq", "system_functions", "rule_function", "expressionRule",
	"expr", "ruleAtom", "sign", "left_par", "right_par", "derivative_rule",
	"numeric_rule", "const_rule", "var_rule", "expr_rule", "atom", "math_constant",
	"infinite", "imaginary", "number", "derivative", "variable", "constant",
	"bi_operators",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SymbolanParser struct {
	*antlr.BaseParser
}

func NewSymbolanParser(input antlr.TokenStream) *SymbolanParser {
	this := new(SymbolanParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Symbolan.g4"

	return this
}

// SymbolanParser tokens.
const (
	SymbolanParserEOF                 = antlr.TokenEOF
	SymbolanParserT__0                = 1
	SymbolanParserT__1                = 2
	SymbolanParserT__2                = 3
	SymbolanParserT__3                = 4
	SymbolanParserT__4                = 5
	SymbolanParserT__5                = 6
	SymbolanParserT__6                = 7
	SymbolanParserT__7                = 8
	SymbolanParserL_PAREN             = 9
	SymbolanParserR_PAREN             = 10
	SymbolanParserL_CURLY             = 11
	SymbolanParserR_CURLY             = 12
	SymbolanParserL_BRACKET           = 13
	SymbolanParserR_BRACKET           = 14
	SymbolanParserASSIGN              = 15
	SymbolanParserCOMMA               = 16
	SymbolanParserFACTORIAL           = 17
	SymbolanParserMIN                 = 18
	SymbolanParserMAX                 = 19
	SymbolanParserSIN                 = 20
	SymbolanParserCOS                 = 21
	SymbolanParserTAN                 = 22
	SymbolanParserLOG                 = 23
	SymbolanParserSQRT                = 24
	SymbolanParserDERIVATIVE_OPERATOR = 25
	SymbolanParserPI                  = 26
	SymbolanParserDERIVATIVE          = 27
	SymbolanParserVAR                 = 28
	SymbolanParserCONST               = 29
	SymbolanParserEULER               = 30
	SymbolanParserINFINITE            = 31
	SymbolanParserUNDEFINED           = 32
	SymbolanParserDERIVATE_RULE       = 33
	SymbolanParserNUMERIC_RULE        = 34
	SymbolanParserCONST_RULE          = 35
	SymbolanParserVAR_RULE            = 36
	SymbolanParserEXPR_RULE           = 37
	SymbolanParserIMAGINARY           = 38
	SymbolanParserDECIMAL_LIT         = 39
	SymbolanParserOCTAL_LIT           = 40
	SymbolanParserHEX_LIT             = 41
	SymbolanParserFLOAT_LIT           = 42
	SymbolanParserIMAGINARY_LIT       = 43
	SymbolanParserIDENTIFIER          = 44
	SymbolanParserNEWLINE             = 45
	SymbolanParserWS                  = 46
	SymbolanParserLINE_COMMENT        = 47
)

// SymbolanParser rules.
const (
	SymbolanParserRULE_ruleSet          = 0
	SymbolanParserRULE_eq               = 1
	SymbolanParserRULE_system_functions = 2
	SymbolanParserRULE_rule_function    = 3
	SymbolanParserRULE_expressionRule   = 4
	SymbolanParserRULE_expr             = 5
	SymbolanParserRULE_ruleAtom         = 6
	SymbolanParserRULE_sign             = 7
	SymbolanParserRULE_left_par         = 8
	SymbolanParserRULE_right_par        = 9
	SymbolanParserRULE_derivative_rule  = 10
	SymbolanParserRULE_numeric_rule     = 11
	SymbolanParserRULE_const_rule       = 12
	SymbolanParserRULE_var_rule         = 13
	SymbolanParserRULE_expr_rule        = 14
	SymbolanParserRULE_atom             = 15
	SymbolanParserRULE_math_constant    = 16
	SymbolanParserRULE_infinite         = 17
	SymbolanParserRULE_imaginary        = 18
	SymbolanParserRULE_number           = 19
	SymbolanParserRULE_derivative       = 20
	SymbolanParserRULE_variable         = 21
	SymbolanParserRULE_constant         = 22
	SymbolanParserRULE_bi_operators     = 23
)

// IRuleSetContext is an interface to support dynamic dispatch.
type IRuleSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleSetContext differentiates from other interfaces.
	IsRuleSetContext()
}

type RuleSetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleSetContext() *RuleSetContext {
	var p = new(RuleSetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_ruleSet
	return p
}

func (*RuleSetContext) IsRuleSetContext() {}

func NewRuleSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleSetContext {
	var p = new(RuleSetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_ruleSet

	return p
}

func (s *RuleSetContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleSetContext) AllExpressionRule() []IExpressionRuleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionRuleContext)(nil)).Elem())
	var tst = make([]IExpressionRuleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionRuleContext)
		}
	}

	return tst
}

func (s *RuleSetContext) ExpressionRule(i int) IExpressionRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionRuleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionRuleContext)
}

func (s *RuleSetContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(SymbolanParserNEWLINE)
}

func (s *RuleSetContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(SymbolanParserNEWLINE, i)
}

func (s *RuleSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterRuleSet(s)
	}
}

func (s *RuleSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitRuleSet(s)
	}
}

func (p *SymbolanParser) RuleSet() (localctx IRuleSetContext) {
	localctx = NewRuleSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SymbolanParserRULE_ruleSet)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SymbolanParserT__6)|(1<<SymbolanParserT__7)|(1<<SymbolanParserL_PAREN)|(1<<SymbolanParserMIN)|(1<<SymbolanParserMAX)|(1<<SymbolanParserSIN)|(1<<SymbolanParserCOS)|(1<<SymbolanParserTAN)|(1<<SymbolanParserLOG)|(1<<SymbolanParserDERIVATIVE_OPERATOR)|(1<<SymbolanParserPI)|(1<<SymbolanParserDERIVATIVE)|(1<<SymbolanParserVAR)|(1<<SymbolanParserCONST)|(1<<SymbolanParserEULER)|(1<<SymbolanParserINFINITE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(SymbolanParserUNDEFINED-32))|(1<<(SymbolanParserDERIVATE_RULE-32))|(1<<(SymbolanParserNUMERIC_RULE-32))|(1<<(SymbolanParserCONST_RULE-32))|(1<<(SymbolanParserVAR_RULE-32))|(1<<(SymbolanParserEXPR_RULE-32))|(1<<(SymbolanParserIMAGINARY-32))|(1<<(SymbolanParserDECIMAL_LIT-32))|(1<<(SymbolanParserOCTAL_LIT-32))|(1<<(SymbolanParserHEX_LIT-32))|(1<<(SymbolanParserFLOAT_LIT-32))|(1<<(SymbolanParserIMAGINARY_LIT-32)))) != 0) {
		{
			p.SetState(48)
			p.ExpressionRule()
		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SymbolanParserNEWLINE {
			{
				p.SetState(49)
				p.Match(SymbolanParserNEWLINE)
			}

			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEqContext is an interface to support dynamic dispatch.
type IEqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqContext differentiates from other interfaces.
	IsEqContext()
}

type EqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqContext() *EqContext {
	var p = new(EqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_eq
	return p
}

func (*EqContext) IsEqContext() {}

func NewEqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqContext {
	var p = new(EqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_eq

	return p
}

func (s *EqContext) GetParser() antlr.Parser { return s.parser }

func (s *EqContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterEq(s)
	}
}

func (s *EqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitEq(s)
	}
}

func (p *SymbolanParser) Eq() (localctx IEqContext) {
	localctx = NewEqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SymbolanParserRULE_eq)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.expr(0)
		}
		{
			p.SetState(62)
			p.Match(SymbolanParserASSIGN)
		}
		{
			p.SetState(63)
			p.expr(0)
		}

	}

	return localctx
}

// ISystem_functionsContext is an interface to support dynamic dispatch.
type ISystem_functionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSystem_functionsContext differentiates from other interfaces.
	IsSystem_functionsContext()
}

type System_functionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySystem_functionsContext() *System_functionsContext {
	var p = new(System_functionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_system_functions
	return p
}

func (*System_functionsContext) IsSystem_functionsContext() {}

func NewSystem_functionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *System_functionsContext {
	var p = new(System_functionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_system_functions

	return p
}

func (s *System_functionsContext) GetParser() antlr.Parser { return s.parser }

func (s *System_functionsContext) MIN() antlr.TerminalNode {
	return s.GetToken(SymbolanParserMIN, 0)
}

func (s *System_functionsContext) MAX() antlr.TerminalNode {
	return s.GetToken(SymbolanParserMAX, 0)
}

func (s *System_functionsContext) SIN() antlr.TerminalNode {
	return s.GetToken(SymbolanParserSIN, 0)
}

func (s *System_functionsContext) COS() antlr.TerminalNode {
	return s.GetToken(SymbolanParserCOS, 0)
}

func (s *System_functionsContext) TAN() antlr.TerminalNode {
	return s.GetToken(SymbolanParserTAN, 0)
}

func (s *System_functionsContext) LOG() antlr.TerminalNode {
	return s.GetToken(SymbolanParserLOG, 0)
}

func (s *System_functionsContext) DERIVATIVE_OPERATOR() antlr.TerminalNode {
	return s.GetToken(SymbolanParserDERIVATIVE_OPERATOR, 0)
}

func (s *System_functionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *System_functionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *System_functionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterSystem_functions(s)
	}
}

func (s *System_functionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitSystem_functions(s)
	}
}

func (p *SymbolanParser) System_functions() (localctx ISystem_functionsContext) {
	localctx = NewSystem_functionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SymbolanParserRULE_system_functions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SymbolanParserMIN)|(1<<SymbolanParserMAX)|(1<<SymbolanParserSIN)|(1<<SymbolanParserCOS)|(1<<SymbolanParserTAN)|(1<<SymbolanParserLOG)|(1<<SymbolanParserDERIVATIVE_OPERATOR))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRule_functionContext is an interface to support dynamic dispatch.
type IRule_functionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRule_functionContext differentiates from other interfaces.
	IsRule_functionContext()
}

type Rule_functionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_functionContext() *Rule_functionContext {
	var p = new(Rule_functionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_rule_function
	return p
}

func (*Rule_functionContext) IsRule_functionContext() {}

func NewRule_functionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_functionContext {
	var p = new(Rule_functionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_rule_function

	return p
}

func (s *Rule_functionContext) GetParser() antlr.Parser { return s.parser }

func (s *Rule_functionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SymbolanParserIDENTIFIER, 0)
}

func (s *Rule_functionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_functionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_functionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterRule_function(s)
	}
}

func (s *Rule_functionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitRule_function(s)
	}
}

func (p *SymbolanParser) Rule_function() (localctx IRule_functionContext) {
	localctx = NewRule_functionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SymbolanParserRULE_rule_function)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(SymbolanParserIDENTIFIER)
	}
	{
		p.SetState(70)
		p.Match(SymbolanParserT__0)
	}

	return localctx
}

// IExpressionRuleContext is an interface to support dynamic dispatch.
type IExpressionRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionRuleContext differentiates from other interfaces.
	IsExpressionRuleContext()
}

type ExpressionRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionRuleContext() *ExpressionRuleContext {
	var p = new(ExpressionRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_expressionRule
	return p
}

func (*ExpressionRuleContext) IsExpressionRuleContext() {}

func NewExpressionRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionRuleContext {
	var p = new(ExpressionRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_expressionRule

	return p
}

func (s *ExpressionRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionRuleContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExpressionRuleContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpressionRuleContext) Rule_function() IRule_functionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRule_functionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRule_functionContext)
}

func (s *ExpressionRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterExpressionRule(s)
	}
}

func (s *ExpressionRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitExpressionRule(s)
	}
}

func (p *SymbolanParser) ExpressionRule() (localctx IExpressionRuleContext) {
	localctx = NewExpressionRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SymbolanParserRULE_expressionRule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.expr(0)
		}
		{
			p.SetState(73)
			p.Match(SymbolanParserT__1)
		}
		{
			p.SetState(74)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.expr(0)
		}
		{
			p.SetState(77)
			p.Match(SymbolanParserT__1)
		}
		{
			p.SetState(78)
			p.Rule_function()
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Left_par() ILeft_parContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeft_parContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeft_parContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) Right_par() IRight_parContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRight_parContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRight_parContext)
}

func (s *ExprContext) System_functions() ISystem_functionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISystem_functionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISystem_functionsContext)
}

func (s *ExprContext) Sign() ISignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignContext)
}

func (s *ExprContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ExprContext) RuleAtom() IRuleAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleAtomContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *SymbolanParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SymbolanParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, SymbolanParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SymbolanParserL_PAREN:
		{
			p.SetState(83)
			p.Left_par()
		}
		{
			p.SetState(84)
			p.expr(0)
		}
		{
			p.SetState(85)
			p.Right_par()
		}

	case SymbolanParserMIN, SymbolanParserMAX, SymbolanParserSIN, SymbolanParserCOS, SymbolanParserTAN, SymbolanParserLOG, SymbolanParserDERIVATIVE_OPERATOR:
		{
			p.SetState(87)
			p.System_functions()
		}
		{
			p.SetState(88)
			p.Left_par()
		}
		{
			p.SetState(89)
			p.expr(0)
		}
		{
			p.SetState(90)
			p.Right_par()
		}

	case SymbolanParserT__6, SymbolanParserT__7:
		{
			p.SetState(92)
			p.Sign()
		}
		{
			p.SetState(93)
			p.expr(3)
		}

	case SymbolanParserPI, SymbolanParserDERIVATIVE, SymbolanParserVAR, SymbolanParserCONST, SymbolanParserEULER, SymbolanParserINFINITE, SymbolanParserUNDEFINED, SymbolanParserIMAGINARY, SymbolanParserDECIMAL_LIT, SymbolanParserOCTAL_LIT, SymbolanParserHEX_LIT, SymbolanParserFLOAT_LIT, SymbolanParserIMAGINARY_LIT:
		{
			p.SetState(95)
			p.Atom()
		}

	case SymbolanParserDERIVATE_RULE, SymbolanParserNUMERIC_RULE, SymbolanParserCONST_RULE, SymbolanParserVAR_RULE, SymbolanParserEXPR_RULE:
		{
			p.SetState(96)
			p.RuleAtom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SymbolanParserRULE_expr)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(100)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SymbolanParserT__2 || _la == SymbolanParserT__3) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(101)
					p.expr(9)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SymbolanParserRULE_expr)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(103)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SymbolanParserT__4 || _la == SymbolanParserT__5) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(104)
					p.expr(8)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SymbolanParserRULE_expr)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(106)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SymbolanParserT__6 || _la == SymbolanParserT__7) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(107)
					p.expr(7)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SymbolanParserRULE_expr)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(109)
					p.expr(6)
				}

			}

		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IRuleAtomContext is an interface to support dynamic dispatch.
type IRuleAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleAtomContext differentiates from other interfaces.
	IsRuleAtomContext()
}

type RuleAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleAtomContext() *RuleAtomContext {
	var p = new(RuleAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_ruleAtom
	return p
}

func (*RuleAtomContext) IsRuleAtomContext() {}

func NewRuleAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleAtomContext {
	var p = new(RuleAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_ruleAtom

	return p
}

func (s *RuleAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleAtomContext) Const_rule() IConst_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_ruleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConst_ruleContext)
}

func (s *RuleAtomContext) Var_rule() IVar_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_ruleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_ruleContext)
}

func (s *RuleAtomContext) Expr_rule() IExpr_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_ruleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_ruleContext)
}

func (s *RuleAtomContext) Numeric_rule() INumeric_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumeric_ruleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumeric_ruleContext)
}

func (s *RuleAtomContext) Derivative_rule() IDerivative_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDerivative_ruleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDerivative_ruleContext)
}

func (s *RuleAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterRuleAtom(s)
	}
}

func (s *RuleAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitRuleAtom(s)
	}
}

func (p *SymbolanParser) RuleAtom() (localctx IRuleAtomContext) {
	localctx = NewRuleAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SymbolanParserRULE_ruleAtom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SymbolanParserCONST_RULE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Const_rule()
		}

	case SymbolanParserVAR_RULE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Var_rule()
		}

	case SymbolanParserEXPR_RULE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(117)
			p.Expr_rule()
		}

	case SymbolanParserNUMERIC_RULE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(118)
			p.Numeric_rule()
		}

	case SymbolanParserDERIVATE_RULE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(119)
			p.Derivative_rule()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISignContext is an interface to support dynamic dispatch.
type ISignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSignContext differentiates from other interfaces.
	IsSignContext()
}

type SignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignContext() *SignContext {
	var p = new(SignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_sign
	return p
}

func (*SignContext) IsSignContext() {}

func NewSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignContext {
	var p = new(SignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_sign

	return p
}

func (s *SignContext) GetParser() antlr.Parser { return s.parser }
func (s *SignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterSign(s)
	}
}

func (s *SignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitSign(s)
	}
}

func (p *SymbolanParser) Sign() (localctx ISignContext) {
	localctx = NewSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SymbolanParserRULE_sign)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SymbolanParserT__6 || _la == SymbolanParserT__7) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILeft_parContext is an interface to support dynamic dispatch.
type ILeft_parContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeft_parContext differentiates from other interfaces.
	IsLeft_parContext()
}

type Left_parContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeft_parContext() *Left_parContext {
	var p = new(Left_parContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_left_par
	return p
}

func (*Left_parContext) IsLeft_parContext() {}

func NewLeft_parContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Left_parContext {
	var p = new(Left_parContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_left_par

	return p
}

func (s *Left_parContext) GetParser() antlr.Parser { return s.parser }
func (s *Left_parContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Left_parContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Left_parContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterLeft_par(s)
	}
}

func (s *Left_parContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitLeft_par(s)
	}
}

func (p *SymbolanParser) Left_par() (localctx ILeft_parContext) {
	localctx = NewLeft_parContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SymbolanParserRULE_left_par)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(SymbolanParserL_PAREN)
	}

	return localctx
}

// IRight_parContext is an interface to support dynamic dispatch.
type IRight_parContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRight_parContext differentiates from other interfaces.
	IsRight_parContext()
}

type Right_parContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRight_parContext() *Right_parContext {
	var p = new(Right_parContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_right_par
	return p
}

func (*Right_parContext) IsRight_parContext() {}

func NewRight_parContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Right_parContext {
	var p = new(Right_parContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_right_par

	return p
}

func (s *Right_parContext) GetParser() antlr.Parser { return s.parser }
func (s *Right_parContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Right_parContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Right_parContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterRight_par(s)
	}
}

func (s *Right_parContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitRight_par(s)
	}
}

func (p *SymbolanParser) Right_par() (localctx IRight_parContext) {
	localctx = NewRight_parContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SymbolanParserRULE_right_par)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(SymbolanParserR_PAREN)
	}

	return localctx
}

// IDerivative_ruleContext is an interface to support dynamic dispatch.
type IDerivative_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDerivative_ruleContext differentiates from other interfaces.
	IsDerivative_ruleContext()
}

type Derivative_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDerivative_ruleContext() *Derivative_ruleContext {
	var p = new(Derivative_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_derivative_rule
	return p
}

func (*Derivative_ruleContext) IsDerivative_ruleContext() {}

func NewDerivative_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Derivative_ruleContext {
	var p = new(Derivative_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_derivative_rule

	return p
}

func (s *Derivative_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Derivative_ruleContext) DERIVATE_RULE() antlr.TerminalNode {
	return s.GetToken(SymbolanParserDERIVATE_RULE, 0)
}

func (s *Derivative_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Derivative_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Derivative_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterDerivative_rule(s)
	}
}

func (s *Derivative_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitDerivative_rule(s)
	}
}

func (p *SymbolanParser) Derivative_rule() (localctx IDerivative_ruleContext) {
	localctx = NewDerivative_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SymbolanParserRULE_derivative_rule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(SymbolanParserDERIVATE_RULE)
	}

	return localctx
}

// INumeric_ruleContext is an interface to support dynamic dispatch.
type INumeric_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumeric_ruleContext differentiates from other interfaces.
	IsNumeric_ruleContext()
}

type Numeric_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumeric_ruleContext() *Numeric_ruleContext {
	var p = new(Numeric_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_numeric_rule
	return p
}

func (*Numeric_ruleContext) IsNumeric_ruleContext() {}

func NewNumeric_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numeric_ruleContext {
	var p = new(Numeric_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_numeric_rule

	return p
}

func (s *Numeric_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Numeric_ruleContext) NUMERIC_RULE() antlr.TerminalNode {
	return s.GetToken(SymbolanParserNUMERIC_RULE, 0)
}

func (s *Numeric_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numeric_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numeric_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterNumeric_rule(s)
	}
}

func (s *Numeric_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitNumeric_rule(s)
	}
}

func (p *SymbolanParser) Numeric_rule() (localctx INumeric_ruleContext) {
	localctx = NewNumeric_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SymbolanParserRULE_numeric_rule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(SymbolanParserNUMERIC_RULE)
	}

	return localctx
}

// IConst_ruleContext is an interface to support dynamic dispatch.
type IConst_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConst_ruleContext differentiates from other interfaces.
	IsConst_ruleContext()
}

type Const_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_ruleContext() *Const_ruleContext {
	var p = new(Const_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_const_rule
	return p
}

func (*Const_ruleContext) IsConst_ruleContext() {}

func NewConst_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_ruleContext {
	var p = new(Const_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_const_rule

	return p
}

func (s *Const_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_ruleContext) CONST_RULE() antlr.TerminalNode {
	return s.GetToken(SymbolanParserCONST_RULE, 0)
}

func (s *Const_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterConst_rule(s)
	}
}

func (s *Const_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitConst_rule(s)
	}
}

func (p *SymbolanParser) Const_rule() (localctx IConst_ruleContext) {
	localctx = NewConst_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SymbolanParserRULE_const_rule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(SymbolanParserCONST_RULE)
	}

	return localctx
}

// IVar_ruleContext is an interface to support dynamic dispatch.
type IVar_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_ruleContext differentiates from other interfaces.
	IsVar_ruleContext()
}

type Var_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_ruleContext() *Var_ruleContext {
	var p = new(Var_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_var_rule
	return p
}

func (*Var_ruleContext) IsVar_ruleContext() {}

func NewVar_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_ruleContext {
	var p = new(Var_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_var_rule

	return p
}

func (s *Var_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_ruleContext) VAR_RULE() antlr.TerminalNode {
	return s.GetToken(SymbolanParserVAR_RULE, 0)
}

func (s *Var_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterVar_rule(s)
	}
}

func (s *Var_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitVar_rule(s)
	}
}

func (p *SymbolanParser) Var_rule() (localctx IVar_ruleContext) {
	localctx = NewVar_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SymbolanParserRULE_var_rule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(SymbolanParserVAR_RULE)
	}

	return localctx
}

// IExpr_ruleContext is an interface to support dynamic dispatch.
type IExpr_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr_ruleContext differentiates from other interfaces.
	IsExpr_ruleContext()
}

type Expr_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_ruleContext() *Expr_ruleContext {
	var p = new(Expr_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_expr_rule
	return p
}

func (*Expr_ruleContext) IsExpr_ruleContext() {}

func NewExpr_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_ruleContext {
	var p = new(Expr_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_expr_rule

	return p
}

func (s *Expr_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_ruleContext) EXPR_RULE() antlr.TerminalNode {
	return s.GetToken(SymbolanParserEXPR_RULE, 0)
}

func (s *Expr_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterExpr_rule(s)
	}
}

func (s *Expr_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitExpr_rule(s)
	}
}

func (p *SymbolanParser) Expr_rule() (localctx IExpr_ruleContext) {
	localctx = NewExpr_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SymbolanParserRULE_expr_rule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(SymbolanParserEXPR_RULE)
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *AtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AtomContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *AtomContext) Math_constant() IMath_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMath_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMath_constantContext)
}

func (s *AtomContext) Imaginary() IImaginaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImaginaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImaginaryContext)
}

func (s *AtomContext) Derivative() IDerivativeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDerivativeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDerivativeContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *SymbolanParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SymbolanParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.Variable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.Constant()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(141)
			p.Number()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(142)
			p.Math_constant()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(143)
			p.Imaginary()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(144)
			p.Derivative()
		}

	}

	return localctx
}

// IMath_constantContext is an interface to support dynamic dispatch.
type IMath_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMath_constantContext differentiates from other interfaces.
	IsMath_constantContext()
}

type Math_constantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMath_constantContext() *Math_constantContext {
	var p = new(Math_constantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_math_constant
	return p
}

func (*Math_constantContext) IsMath_constantContext() {}

func NewMath_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Math_constantContext {
	var p = new(Math_constantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_math_constant

	return p
}

func (s *Math_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Math_constantContext) PI() antlr.TerminalNode {
	return s.GetToken(SymbolanParserPI, 0)
}

func (s *Math_constantContext) EULER() antlr.TerminalNode {
	return s.GetToken(SymbolanParserEULER, 0)
}

func (s *Math_constantContext) INFINITE() antlr.TerminalNode {
	return s.GetToken(SymbolanParserINFINITE, 0)
}

func (s *Math_constantContext) UNDEFINED() antlr.TerminalNode {
	return s.GetToken(SymbolanParserUNDEFINED, 0)
}

func (s *Math_constantContext) IMAGINARY() antlr.TerminalNode {
	return s.GetToken(SymbolanParserIMAGINARY, 0)
}

func (s *Math_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Math_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Math_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterMath_constant(s)
	}
}

func (s *Math_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitMath_constant(s)
	}
}

func (p *SymbolanParser) Math_constant() (localctx IMath_constantContext) {
	localctx = NewMath_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SymbolanParserRULE_math_constant)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(SymbolanParserPI-26))|(1<<(SymbolanParserEULER-26))|(1<<(SymbolanParserINFINITE-26))|(1<<(SymbolanParserUNDEFINED-26))|(1<<(SymbolanParserIMAGINARY-26)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IInfiniteContext is an interface to support dynamic dispatch.
type IInfiniteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInfiniteContext differentiates from other interfaces.
	IsInfiniteContext()
}

type InfiniteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInfiniteContext() *InfiniteContext {
	var p = new(InfiniteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_infinite
	return p
}

func (*InfiniteContext) IsInfiniteContext() {}

func NewInfiniteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfiniteContext {
	var p = new(InfiniteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_infinite

	return p
}

func (s *InfiniteContext) GetParser() antlr.Parser { return s.parser }

func (s *InfiniteContext) INFINITE() antlr.TerminalNode {
	return s.GetToken(SymbolanParserINFINITE, 0)
}

func (s *InfiniteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfiniteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InfiniteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterInfinite(s)
	}
}

func (s *InfiniteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitInfinite(s)
	}
}

func (p *SymbolanParser) Infinite() (localctx IInfiniteContext) {
	localctx = NewInfiniteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SymbolanParserRULE_infinite)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(SymbolanParserINFINITE)
	}

	return localctx
}

// IImaginaryContext is an interface to support dynamic dispatch.
type IImaginaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImaginaryContext differentiates from other interfaces.
	IsImaginaryContext()
}

type ImaginaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImaginaryContext() *ImaginaryContext {
	var p = new(ImaginaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_imaginary
	return p
}

func (*ImaginaryContext) IsImaginaryContext() {}

func NewImaginaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImaginaryContext {
	var p = new(ImaginaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_imaginary

	return p
}

func (s *ImaginaryContext) GetParser() antlr.Parser { return s.parser }

func (s *ImaginaryContext) IMAGINARY_LIT() antlr.TerminalNode {
	return s.GetToken(SymbolanParserIMAGINARY_LIT, 0)
}

func (s *ImaginaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImaginaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImaginaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterImaginary(s)
	}
}

func (s *ImaginaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitImaginary(s)
	}
}

func (p *SymbolanParser) Imaginary() (localctx IImaginaryContext) {
	localctx = NewImaginaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SymbolanParserRULE_imaginary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(SymbolanParserIMAGINARY_LIT)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(SymbolanParserDECIMAL_LIT, 0)
}

func (s *NumberContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(SymbolanParserFLOAT_LIT, 0)
}

func (s *NumberContext) OCTAL_LIT() antlr.TerminalNode {
	return s.GetToken(SymbolanParserOCTAL_LIT, 0)
}

func (s *NumberContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(SymbolanParserHEX_LIT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *SymbolanParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SymbolanParserRULE_number)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(SymbolanParserDECIMAL_LIT-39))|(1<<(SymbolanParserOCTAL_LIT-39))|(1<<(SymbolanParserHEX_LIT-39))|(1<<(SymbolanParserFLOAT_LIT-39)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDerivativeContext is an interface to support dynamic dispatch.
type IDerivativeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDerivativeContext differentiates from other interfaces.
	IsDerivativeContext()
}

type DerivativeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDerivativeContext() *DerivativeContext {
	var p = new(DerivativeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_derivative
	return p
}

func (*DerivativeContext) IsDerivativeContext() {}

func NewDerivativeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DerivativeContext {
	var p = new(DerivativeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_derivative

	return p
}

func (s *DerivativeContext) GetParser() antlr.Parser { return s.parser }

func (s *DerivativeContext) DERIVATIVE() antlr.TerminalNode {
	return s.GetToken(SymbolanParserDERIVATIVE, 0)
}

func (s *DerivativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DerivativeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DerivativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterDerivative(s)
	}
}

func (s *DerivativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitDerivative(s)
	}
}

func (p *SymbolanParser) Derivative() (localctx IDerivativeContext) {
	localctx = NewDerivativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SymbolanParserRULE_derivative)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(SymbolanParserDERIVATIVE)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VAR() antlr.TerminalNode {
	return s.GetToken(SymbolanParserVAR, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *SymbolanParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SymbolanParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(SymbolanParserVAR)
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) CONST() antlr.TerminalNode {
	return s.GetToken(SymbolanParserCONST, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *SymbolanParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SymbolanParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(SymbolanParserCONST)
	}

	return localctx
}

// IBi_operatorsContext is an interface to support dynamic dispatch.
type IBi_operatorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBi_operatorsContext differentiates from other interfaces.
	IsBi_operatorsContext()
}

type Bi_operatorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBi_operatorsContext() *Bi_operatorsContext {
	var p = new(Bi_operatorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_bi_operators
	return p
}

func (*Bi_operatorsContext) IsBi_operatorsContext() {}

func NewBi_operatorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bi_operatorsContext {
	var p = new(Bi_operatorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_bi_operators

	return p
}

func (s *Bi_operatorsContext) GetParser() antlr.Parser { return s.parser }
func (s *Bi_operatorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bi_operatorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bi_operatorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterBi_operators(s)
	}
}

func (s *Bi_operatorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitBi_operators(s)
	}
}

func (p *SymbolanParser) Bi_operators() (localctx IBi_operatorsContext) {
	localctx = NewBi_operatorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SymbolanParserRULE_bi_operators)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SymbolanParserT__2, SymbolanParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SymbolanParserT__2 || _la == SymbolanParserT__3) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case SymbolanParserT__4, SymbolanParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SymbolanParserT__4 || _la == SymbolanParserT__5) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case SymbolanParserT__6, SymbolanParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SymbolanParserT__6 || _la == SymbolanParserT__7) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *SymbolanParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SymbolanParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
