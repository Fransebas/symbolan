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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 35, 85, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 30, 10, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 41, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 54, 10, 4, 12, 4, 14, 4, 57, 11,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 5, 8, 72, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 5, 12, 83, 10, 12, 3, 12, 2, 3, 6, 13, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 2, 7, 3, 2, 18, 23, 3, 2, 3, 4, 3, 2, 5, 6, 3, 2, 7,
	8, 5, 2, 24, 24, 27, 27, 29, 33, 2, 85, 2, 29, 3, 2, 2, 2, 4, 31, 3, 2,
	2, 2, 6, 40, 3, 2, 2, 2, 8, 58, 3, 2, 2, 2, 10, 63, 3, 2, 2, 2, 12, 65,
	3, 2, 2, 2, 14, 71, 3, 2, 2, 2, 16, 73, 3, 2, 2, 2, 18, 75, 3, 2, 2, 2,
	20, 77, 3, 2, 2, 2, 22, 82, 3, 2, 2, 2, 24, 30, 5, 6, 4, 2, 25, 26, 5,
	6, 4, 2, 26, 27, 7, 15, 2, 2, 27, 28, 5, 6, 4, 2, 28, 30, 3, 2, 2, 2, 29,
	24, 3, 2, 2, 2, 29, 25, 3, 2, 2, 2, 30, 3, 3, 2, 2, 2, 31, 32, 9, 2, 2,
	2, 32, 5, 3, 2, 2, 2, 33, 34, 8, 4, 1, 2, 34, 41, 5, 14, 8, 2, 35, 36,
	5, 10, 6, 2, 36, 37, 5, 6, 4, 2, 37, 38, 5, 12, 7, 2, 38, 41, 3, 2, 2,
	2, 39, 41, 5, 8, 5, 2, 40, 33, 3, 2, 2, 2, 40, 35, 3, 2, 2, 2, 40, 39,
	3, 2, 2, 2, 41, 55, 3, 2, 2, 2, 42, 43, 12, 9, 2, 2, 43, 54, 5, 6, 4, 10,
	44, 45, 12, 6, 2, 2, 45, 46, 9, 3, 2, 2, 46, 54, 5, 6, 4, 7, 47, 48, 12,
	5, 2, 2, 48, 49, 9, 4, 2, 2, 49, 54, 5, 6, 4, 6, 50, 51, 12, 4, 2, 2, 51,
	52, 9, 5, 2, 2, 52, 54, 5, 6, 4, 5, 53, 42, 3, 2, 2, 2, 53, 44, 3, 2, 2,
	2, 53, 47, 3, 2, 2, 2, 53, 50, 3, 2, 2, 2, 54, 57, 3, 2, 2, 2, 55, 53,
	3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 7, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2,
	58, 59, 5, 4, 3, 2, 59, 60, 5, 10, 6, 2, 60, 61, 5, 6, 4, 2, 61, 62, 5,
	12, 7, 2, 62, 9, 3, 2, 2, 2, 63, 64, 7, 9, 2, 2, 64, 11, 3, 2, 2, 2, 65,
	66, 7, 10, 2, 2, 66, 13, 3, 2, 2, 2, 67, 72, 5, 16, 9, 2, 68, 72, 5, 18,
	10, 2, 69, 72, 5, 20, 11, 2, 70, 72, 5, 16, 9, 2, 71, 67, 3, 2, 2, 2, 71,
	68, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 15, 3, 2, 2,
	2, 73, 74, 9, 6, 2, 2, 74, 17, 3, 2, 2, 2, 75, 76, 7, 25, 2, 2, 76, 19,
	3, 2, 2, 2, 77, 78, 7, 26, 2, 2, 78, 21, 3, 2, 2, 2, 79, 83, 9, 3, 2, 2,
	80, 83, 9, 4, 2, 2, 81, 83, 9, 5, 2, 2, 82, 79, 3, 2, 2, 2, 82, 80, 3,
	2, 2, 2, 82, 81, 3, 2, 2, 2, 83, 23, 3, 2, 2, 2, 8, 29, 40, 53, 55, 71,
	82,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'^'", "'**'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "'='", "','", "'!'", "'min'", "'max'", "'sin'", "'cos'",
	"'tan'", "'log'", "'pi'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY",
	"L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "FACTORIAL", "MIN", "MAX",
	"SIN", "COS", "TAN", "LOG", "PI", "VAR", "CONST", "EULER", "IMAGINARY",
	"DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "IMAGINARY_LIT", "WS",
	"LINE_COMMENT",
}

var ruleNames = []string{
	"eq", "system_functions", "expr", "function", "left_par", "right_par",
	"atom", "number", "variable", "constant", "bi_operators",
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
	SymbolanParserEOF           = antlr.TokenEOF
	SymbolanParserT__0          = 1
	SymbolanParserT__1          = 2
	SymbolanParserT__2          = 3
	SymbolanParserT__3          = 4
	SymbolanParserT__4          = 5
	SymbolanParserT__5          = 6
	SymbolanParserL_PAREN       = 7
	SymbolanParserR_PAREN       = 8
	SymbolanParserL_CURLY       = 9
	SymbolanParserR_CURLY       = 10
	SymbolanParserL_BRACKET     = 11
	SymbolanParserR_BRACKET     = 12
	SymbolanParserASSIGN        = 13
	SymbolanParserCOMMA         = 14
	SymbolanParserFACTORIAL     = 15
	SymbolanParserMIN           = 16
	SymbolanParserMAX           = 17
	SymbolanParserSIN           = 18
	SymbolanParserCOS           = 19
	SymbolanParserTAN           = 20
	SymbolanParserLOG           = 21
	SymbolanParserPI            = 22
	SymbolanParserVAR           = 23
	SymbolanParserCONST         = 24
	SymbolanParserEULER         = 25
	SymbolanParserIMAGINARY     = 26
	SymbolanParserDECIMAL_LIT   = 27
	SymbolanParserOCTAL_LIT     = 28
	SymbolanParserHEX_LIT       = 29
	SymbolanParserFLOAT_LIT     = 30
	SymbolanParserIMAGINARY_LIT = 31
	SymbolanParserWS            = 32
	SymbolanParserLINE_COMMENT  = 33
)

// SymbolanParser rules.
const (
	SymbolanParserRULE_eq               = 0
	SymbolanParserRULE_system_functions = 1
	SymbolanParserRULE_expr             = 2
	SymbolanParserRULE_function         = 3
	SymbolanParserRULE_left_par         = 4
	SymbolanParserRULE_right_par        = 5
	SymbolanParserRULE_atom             = 6
	SymbolanParserRULE_number           = 7
	SymbolanParserRULE_variable         = 8
	SymbolanParserRULE_constant         = 9
	SymbolanParserRULE_bi_operators     = 10
)

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
	p.EnterRule(localctx, 0, SymbolanParserRULE_eq)

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

	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.expr(0)
		}
		{
			p.SetState(24)
			p.Match(SymbolanParserASSIGN)
		}
		{
			p.SetState(25)
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
	p.EnterRule(localctx, 2, SymbolanParserRULE_system_functions)
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
		p.SetState(29)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SymbolanParserMIN)|(1<<SymbolanParserMAX)|(1<<SymbolanParserSIN)|(1<<SymbolanParserCOS)|(1<<SymbolanParserTAN)|(1<<SymbolanParserLOG))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

func (s *ExprContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

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

func (s *ExprContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
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
	_startState := 4
	p.EnterRecursionRule(localctx, 4, SymbolanParserRULE_expr, _p)
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
	p.SetState(38)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SymbolanParserPI, SymbolanParserVAR, SymbolanParserCONST, SymbolanParserEULER, SymbolanParserDECIMAL_LIT, SymbolanParserOCTAL_LIT, SymbolanParserHEX_LIT, SymbolanParserFLOAT_LIT, SymbolanParserIMAGINARY_LIT:
		{
			p.SetState(32)
			p.Atom()
		}

	case SymbolanParserL_PAREN:
		{
			p.SetState(33)
			p.Left_par()
		}
		{
			p.SetState(34)
			p.expr(0)
		}
		{
			p.SetState(35)
			p.Right_par()
		}

	case SymbolanParserMIN, SymbolanParserMAX, SymbolanParserSIN, SymbolanParserCOS, SymbolanParserTAN, SymbolanParserLOG:
		{
			p.SetState(37)
			p.Function()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SymbolanParserRULE_expr)
				p.SetState(40)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(41)
					p.expr(8)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SymbolanParserRULE_expr)
				p.SetState(42)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(43)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SymbolanParserT__0 || _la == SymbolanParserT__1) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(44)
					p.expr(5)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SymbolanParserRULE_expr)
				p.SetState(45)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(46)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SymbolanParserT__2 || _la == SymbolanParserT__3) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(47)
					p.expr(4)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SymbolanParserRULE_expr)
				p.SetState(48)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(49)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SymbolanParserT__4 || _la == SymbolanParserT__5) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(50)
					p.expr(3)
				}

			}

		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SymbolanParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SymbolanParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) System_functions() ISystem_functionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISystem_functionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISystem_functionsContext)
}

func (s *FunctionContext) Left_par() ILeft_parContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeft_parContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeft_parContext)
}

func (s *FunctionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionContext) Right_par() IRight_parContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRight_parContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRight_parContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SymbolanListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *SymbolanParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SymbolanParserRULE_function)

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
		p.SetState(56)
		p.System_functions()
	}
	{
		p.SetState(57)
		p.Left_par()
	}
	{
		p.SetState(58)
		p.expr(0)
	}
	{
		p.SetState(59)
		p.Right_par()
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
	p.EnterRule(localctx, 8, SymbolanParserRULE_left_par)

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
		p.SetState(61)
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
	p.EnterRule(localctx, 10, SymbolanParserRULE_right_par)

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
		p.SetState(63)
		p.Match(SymbolanParserR_PAREN)
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
	p.EnterRule(localctx, 12, SymbolanParserRULE_atom)

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

	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Variable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Constant()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.Number()
		}

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

func (s *NumberContext) IMAGINARY_LIT() antlr.TerminalNode {
	return s.GetToken(SymbolanParserIMAGINARY_LIT, 0)
}

func (s *NumberContext) OCTAL_LIT() antlr.TerminalNode {
	return s.GetToken(SymbolanParserOCTAL_LIT, 0)
}

func (s *NumberContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(SymbolanParserHEX_LIT, 0)
}

func (s *NumberContext) PI() antlr.TerminalNode {
	return s.GetToken(SymbolanParserPI, 0)
}

func (s *NumberContext) EULER() antlr.TerminalNode {
	return s.GetToken(SymbolanParserEULER, 0)
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
	p.EnterRule(localctx, 14, SymbolanParserRULE_number)
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
		p.SetState(71)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SymbolanParserPI)|(1<<SymbolanParserEULER)|(1<<SymbolanParserDECIMAL_LIT)|(1<<SymbolanParserOCTAL_LIT)|(1<<SymbolanParserHEX_LIT)|(1<<SymbolanParserFLOAT_LIT)|(1<<SymbolanParserIMAGINARY_LIT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.EnterRule(localctx, 16, SymbolanParserRULE_variable)

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
		p.SetState(73)
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
	p.EnterRule(localctx, 18, SymbolanParserRULE_constant)

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
		p.SetState(75)
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
	p.EnterRule(localctx, 20, SymbolanParserRULE_bi_operators)
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

	p.SetState(80)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SymbolanParserT__0, SymbolanParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SymbolanParserT__0 || _la == SymbolanParserT__1) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case SymbolanParserT__2, SymbolanParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SymbolanParserT__2 || _la == SymbolanParserT__3) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case SymbolanParserT__4, SymbolanParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SymbolanParserT__4 || _la == SymbolanParserT__5) {
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
	case 2:
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
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
