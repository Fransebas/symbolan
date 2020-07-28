// Code generated from TreeLan.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // TreeLan

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 42, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 28, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 4, 4, 2, 18,
	25, 28, 28, 3, 2, 6, 18, 2, 36, 2, 14, 3, 2, 2, 2, 4, 27, 3, 2, 2, 2, 6,
	29, 3, 2, 2, 2, 8, 33, 3, 2, 2, 2, 10, 37, 3, 2, 2, 2, 12, 39, 3, 2, 2,
	2, 14, 15, 5, 4, 3, 2, 15, 16, 7, 27, 2, 2, 16, 17, 5, 4, 3, 2, 17, 3,
	3, 2, 2, 2, 18, 19, 7, 3, 2, 2, 19, 20, 5, 4, 3, 2, 20, 21, 7, 4, 2, 2,
	21, 22, 5, 6, 4, 2, 22, 23, 7, 3, 2, 2, 23, 24, 5, 4, 3, 2, 24, 25, 7,
	4, 2, 2, 25, 28, 3, 2, 2, 2, 26, 28, 5, 8, 5, 2, 27, 18, 3, 2, 2, 2, 27,
	26, 3, 2, 2, 2, 28, 5, 3, 2, 2, 2, 29, 30, 5, 12, 7, 2, 30, 31, 7, 5, 2,
	2, 31, 32, 7, 26, 2, 2, 32, 7, 3, 2, 2, 2, 33, 34, 5, 10, 6, 2, 34, 35,
	7, 5, 2, 2, 35, 36, 7, 26, 2, 2, 36, 9, 3, 2, 2, 2, 37, 38, 9, 2, 2, 2,
	38, 11, 3, 2, 2, 2, 39, 40, 9, 3, 2, 2, 40, 13, 3, 2, 2, 2, 3, 27,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'_'", "'ADDITION'", "'SUBSTRACTION'", "'MULTIPLICATION'",
	"'DIVISION'", "'EXPONENTIAL'", "'LOGARITHM'", "'SIN'", "'COS'", "'TAN'",
	"'TRANSFORMATION'", "'EQUATION'", "'LEAF'", "'ANY'", "'NUMERIC_CONSTANT'",
	"'CONSTANT'", "'VARIABLE'", "'CONSTANT_EXPRESSION'", "'VARIABLE_EXPRESSION'",
	"'NUMERIC_CONSTANT_EXPRESSION'", "'SYSTEM_FUNCTION'", "", "'=>'",
}
var symbolicNames = []string{
	"", "", "", "", "ADDITION", "SUBSTRACTION", "MULTIPLICATION", "DIVISION",
	"EXPONENTIAL", "LOGARITHM", "SIN", "COS", "TAN", "TRANSFORMATION", "EQUATION",
	"LEAF", "ANY", "NUMERIC_CONSTANT", "CONSTANT", "VARIABLE", "CONSTANT_EXPRESSION",
	"VARIABLE_EXPRESSION", "NUMERIC_CONSTANT_EXPRESSION", "SYSTEM_FUNCTION",
	"ID", "ASSIGN", "NUMBER", "WS",
}

var ruleNames = []string{
	"treeRule", "treeDescriptor", "rootNode", "leafNode", "classValue", "operationClass",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TreeLanParser struct {
	*antlr.BaseParser
}

func NewTreeLanParser(input antlr.TokenStream) *TreeLanParser {
	this := new(TreeLanParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TreeLan.g4"

	return this
}

// TreeLanParser tokens.
const (
	TreeLanParserEOF                         = antlr.TokenEOF
	TreeLanParserT__0                        = 1
	TreeLanParserT__1                        = 2
	TreeLanParserT__2                        = 3
	TreeLanParserADDITION                    = 4
	TreeLanParserSUBSTRACTION                = 5
	TreeLanParserMULTIPLICATION              = 6
	TreeLanParserDIVISION                    = 7
	TreeLanParserEXPONENTIAL                 = 8
	TreeLanParserLOGARITHM                   = 9
	TreeLanParserSIN                         = 10
	TreeLanParserCOS                         = 11
	TreeLanParserTAN                         = 12
	TreeLanParserTRANSFORMATION              = 13
	TreeLanParserEQUATION                    = 14
	TreeLanParserLEAF                        = 15
	TreeLanParserANY                         = 16
	TreeLanParserNUMERIC_CONSTANT            = 17
	TreeLanParserCONSTANT                    = 18
	TreeLanParserVARIABLE                    = 19
	TreeLanParserCONSTANT_EXPRESSION         = 20
	TreeLanParserVARIABLE_EXPRESSION         = 21
	TreeLanParserNUMERIC_CONSTANT_EXPRESSION = 22
	TreeLanParserSYSTEM_FUNCTION             = 23
	TreeLanParserID                          = 24
	TreeLanParserASSIGN                      = 25
	TreeLanParserNUMBER                      = 26
	TreeLanParserWS                          = 27
)

// TreeLanParser rules.
const (
	TreeLanParserRULE_treeRule       = 0
	TreeLanParserRULE_treeDescriptor = 1
	TreeLanParserRULE_rootNode       = 2
	TreeLanParserRULE_leafNode       = 3
	TreeLanParserRULE_classValue     = 4
	TreeLanParserRULE_operationClass = 5
)

// ITreeRuleContext is an interface to support dynamic dispatch.
type ITreeRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTreeRuleContext differentiates from other interfaces.
	IsTreeRuleContext()
}

type TreeRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTreeRuleContext() *TreeRuleContext {
	var p = new(TreeRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TreeLanParserRULE_treeRule
	return p
}

func (*TreeRuleContext) IsTreeRuleContext() {}

func NewTreeRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TreeRuleContext {
	var p = new(TreeRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TreeLanParserRULE_treeRule

	return p
}

func (s *TreeRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TreeRuleContext) AllTreeDescriptor() []ITreeDescriptorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITreeDescriptorContext)(nil)).Elem())
	var tst = make([]ITreeDescriptorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITreeDescriptorContext)
		}
	}

	return tst
}

func (s *TreeRuleContext) TreeDescriptor(i int) ITreeDescriptorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITreeDescriptorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITreeDescriptorContext)
}

func (s *TreeRuleContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(TreeLanParserASSIGN, 0)
}

func (s *TreeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TreeRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TreeRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreeLanListener); ok {
		listenerT.EnterTreeRule(s)
	}
}

func (s *TreeRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreeLanListener); ok {
		listenerT.ExitTreeRule(s)
	}
}

func (p *TreeLanParser) TreeRule() (localctx ITreeRuleContext) {
	localctx = NewTreeRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TreeLanParserRULE_treeRule)

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
		p.SetState(12)
		p.TreeDescriptor()
	}
	{
		p.SetState(13)
		p.Match(TreeLanParserASSIGN)
	}
	{
		p.SetState(14)
		p.TreeDescriptor()
	}

	return localctx
}

// ITreeDescriptorContext is an interface to support dynamic dispatch.
type ITreeDescriptorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTreeDescriptorContext differentiates from other interfaces.
	IsTreeDescriptorContext()
}

type TreeDescriptorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTreeDescriptorContext() *TreeDescriptorContext {
	var p = new(TreeDescriptorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TreeLanParserRULE_treeDescriptor
	return p
}

func (*TreeDescriptorContext) IsTreeDescriptorContext() {}

func NewTreeDescriptorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TreeDescriptorContext {
	var p = new(TreeDescriptorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TreeLanParserRULE_treeDescriptor

	return p
}

func (s *TreeDescriptorContext) GetParser() antlr.Parser { return s.parser }

func (s *TreeDescriptorContext) AllTreeDescriptor() []ITreeDescriptorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITreeDescriptorContext)(nil)).Elem())
	var tst = make([]ITreeDescriptorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITreeDescriptorContext)
		}
	}

	return tst
}

func (s *TreeDescriptorContext) TreeDescriptor(i int) ITreeDescriptorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITreeDescriptorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITreeDescriptorContext)
}

func (s *TreeDescriptorContext) RootNode() IRootNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRootNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRootNodeContext)
}

func (s *TreeDescriptorContext) LeafNode() ILeafNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeafNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeafNodeContext)
}

func (s *TreeDescriptorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TreeDescriptorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TreeDescriptorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreeLanListener); ok {
		listenerT.EnterTreeDescriptor(s)
	}
}

func (s *TreeDescriptorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreeLanListener); ok {
		listenerT.ExitTreeDescriptor(s)
	}
}

func (p *TreeLanParser) TreeDescriptor() (localctx ITreeDescriptorContext) {
	localctx = NewTreeDescriptorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TreeLanParserRULE_treeDescriptor)

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

	p.SetState(25)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TreeLanParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(16)
			p.Match(TreeLanParserT__0)
		}
		{
			p.SetState(17)
			p.TreeDescriptor()
		}
		{
			p.SetState(18)
			p.Match(TreeLanParserT__1)
		}
		{
			p.SetState(19)
			p.RootNode()
		}
		{
			p.SetState(20)
			p.Match(TreeLanParserT__0)
		}
		{
			p.SetState(21)
			p.TreeDescriptor()
		}
		{
			p.SetState(22)
			p.Match(TreeLanParserT__1)
		}

	case TreeLanParserANY, TreeLanParserNUMERIC_CONSTANT, TreeLanParserCONSTANT, TreeLanParserVARIABLE, TreeLanParserCONSTANT_EXPRESSION, TreeLanParserVARIABLE_EXPRESSION, TreeLanParserNUMERIC_CONSTANT_EXPRESSION, TreeLanParserSYSTEM_FUNCTION, TreeLanParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.LeafNode()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRootNodeContext is an interface to support dynamic dispatch.
type IRootNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootNodeContext differentiates from other interfaces.
	IsRootNodeContext()
}

type RootNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootNodeContext() *RootNodeContext {
	var p = new(RootNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TreeLanParserRULE_rootNode
	return p
}

func (*RootNodeContext) IsRootNodeContext() {}

func NewRootNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootNodeContext {
	var p = new(RootNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TreeLanParserRULE_rootNode

	return p
}

func (s *RootNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *RootNodeContext) OperationClass() IOperationClassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationClassContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationClassContext)
}

func (s *RootNodeContext) ID() antlr.TerminalNode {
	return s.GetToken(TreeLanParserID, 0)
}

func (s *RootNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreeLanListener); ok {
		listenerT.EnterRootNode(s)
	}
}

func (s *RootNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreeLanListener); ok {
		listenerT.ExitRootNode(s)
	}
}

func (p *TreeLanParser) RootNode() (localctx IRootNodeContext) {
	localctx = NewRootNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TreeLanParserRULE_rootNode)

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
		p.SetState(27)
		p.OperationClass()
	}
	{
		p.SetState(28)
		p.Match(TreeLanParserT__2)
	}
	{
		p.SetState(29)
		p.Match(TreeLanParserID)
	}

	return localctx
}

// ILeafNodeContext is an interface to support dynamic dispatch.
type ILeafNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeafNodeContext differentiates from other interfaces.
	IsLeafNodeContext()
}

type LeafNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeafNodeContext() *LeafNodeContext {
	var p = new(LeafNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TreeLanParserRULE_leafNode
	return p
}

func (*LeafNodeContext) IsLeafNodeContext() {}

func NewLeafNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeafNodeContext {
	var p = new(LeafNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TreeLanParserRULE_leafNode

	return p
}

func (s *LeafNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *LeafNodeContext) ClassValue() IClassValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassValueContext)
}

func (s *LeafNodeContext) ID() antlr.TerminalNode {
	return s.GetToken(TreeLanParserID, 0)
}

func (s *LeafNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeafNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeafNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreeLanListener); ok {
		listenerT.EnterLeafNode(s)
	}
}

func (s *LeafNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreeLanListener); ok {
		listenerT.ExitLeafNode(s)
	}
}

func (p *TreeLanParser) LeafNode() (localctx ILeafNodeContext) {
	localctx = NewLeafNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TreeLanParserRULE_leafNode)

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
		p.SetState(31)
		p.ClassValue()
	}
	{
		p.SetState(32)
		p.Match(TreeLanParserT__2)
	}
	{
		p.SetState(33)
		p.Match(TreeLanParserID)
	}

	return localctx
}

// IClassValueContext is an interface to support dynamic dispatch.
type IClassValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassValueContext differentiates from other interfaces.
	IsClassValueContext()
}

type ClassValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassValueContext() *ClassValueContext {
	var p = new(ClassValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TreeLanParserRULE_classValue
	return p
}

func (*ClassValueContext) IsClassValueContext() {}

func NewClassValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassValueContext {
	var p = new(ClassValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TreeLanParserRULE_classValue

	return p
}

func (s *ClassValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassValueContext) NUMERIC_CONSTANT() antlr.TerminalNode {
	return s.GetToken(TreeLanParserNUMERIC_CONSTANT, 0)
}

func (s *ClassValueContext) CONSTANT() antlr.TerminalNode {
	return s.GetToken(TreeLanParserCONSTANT, 0)
}

func (s *ClassValueContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(TreeLanParserVARIABLE, 0)
}

func (s *ClassValueContext) CONSTANT_EXPRESSION() antlr.TerminalNode {
	return s.GetToken(TreeLanParserCONSTANT_EXPRESSION, 0)
}

func (s *ClassValueContext) VARIABLE_EXPRESSION() antlr.TerminalNode {
	return s.GetToken(TreeLanParserVARIABLE_EXPRESSION, 0)
}

func (s *ClassValueContext) NUMERIC_CONSTANT_EXPRESSION() antlr.TerminalNode {
	return s.GetToken(TreeLanParserNUMERIC_CONSTANT_EXPRESSION, 0)
}

func (s *ClassValueContext) SYSTEM_FUNCTION() antlr.TerminalNode {
	return s.GetToken(TreeLanParserSYSTEM_FUNCTION, 0)
}

func (s *ClassValueContext) ANY() antlr.TerminalNode {
	return s.GetToken(TreeLanParserANY, 0)
}

func (s *ClassValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TreeLanParserNUMBER, 0)
}

func (s *ClassValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreeLanListener); ok {
		listenerT.EnterClassValue(s)
	}
}

func (s *ClassValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreeLanListener); ok {
		listenerT.ExitClassValue(s)
	}
}

func (p *TreeLanParser) ClassValue() (localctx IClassValueContext) {
	localctx = NewClassValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TreeLanParserRULE_classValue)
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
		p.SetState(35)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TreeLanParserANY)|(1<<TreeLanParserNUMERIC_CONSTANT)|(1<<TreeLanParserCONSTANT)|(1<<TreeLanParserVARIABLE)|(1<<TreeLanParserCONSTANT_EXPRESSION)|(1<<TreeLanParserVARIABLE_EXPRESSION)|(1<<TreeLanParserNUMERIC_CONSTANT_EXPRESSION)|(1<<TreeLanParserSYSTEM_FUNCTION)|(1<<TreeLanParserNUMBER))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperationClassContext is an interface to support dynamic dispatch.
type IOperationClassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationClassContext differentiates from other interfaces.
	IsOperationClassContext()
}

type OperationClassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationClassContext() *OperationClassContext {
	var p = new(OperationClassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TreeLanParserRULE_operationClass
	return p
}

func (*OperationClassContext) IsOperationClassContext() {}

func NewOperationClassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationClassContext {
	var p = new(OperationClassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TreeLanParserRULE_operationClass

	return p
}

func (s *OperationClassContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationClassContext) ADDITION() antlr.TerminalNode {
	return s.GetToken(TreeLanParserADDITION, 0)
}

func (s *OperationClassContext) SUBSTRACTION() antlr.TerminalNode {
	return s.GetToken(TreeLanParserSUBSTRACTION, 0)
}

func (s *OperationClassContext) MULTIPLICATION() antlr.TerminalNode {
	return s.GetToken(TreeLanParserMULTIPLICATION, 0)
}

func (s *OperationClassContext) DIVISION() antlr.TerminalNode {
	return s.GetToken(TreeLanParserDIVISION, 0)
}

func (s *OperationClassContext) EXPONENTIAL() antlr.TerminalNode {
	return s.GetToken(TreeLanParserEXPONENTIAL, 0)
}

func (s *OperationClassContext) LOGARITHM() antlr.TerminalNode {
	return s.GetToken(TreeLanParserLOGARITHM, 0)
}

func (s *OperationClassContext) SIN() antlr.TerminalNode {
	return s.GetToken(TreeLanParserSIN, 0)
}

func (s *OperationClassContext) COS() antlr.TerminalNode {
	return s.GetToken(TreeLanParserCOS, 0)
}

func (s *OperationClassContext) TAN() antlr.TerminalNode {
	return s.GetToken(TreeLanParserTAN, 0)
}

func (s *OperationClassContext) TRANSFORMATION() antlr.TerminalNode {
	return s.GetToken(TreeLanParserTRANSFORMATION, 0)
}

func (s *OperationClassContext) EQUATION() antlr.TerminalNode {
	return s.GetToken(TreeLanParserEQUATION, 0)
}

func (s *OperationClassContext) LEAF() antlr.TerminalNode {
	return s.GetToken(TreeLanParserLEAF, 0)
}

func (s *OperationClassContext) ANY() antlr.TerminalNode {
	return s.GetToken(TreeLanParserANY, 0)
}

func (s *OperationClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationClassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreeLanListener); ok {
		listenerT.EnterOperationClass(s)
	}
}

func (s *OperationClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreeLanListener); ok {
		listenerT.ExitOperationClass(s)
	}
}

func (p *TreeLanParser) OperationClass() (localctx IOperationClassContext) {
	localctx = NewOperationClassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TreeLanParserRULE_operationClass)
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
		p.SetState(37)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TreeLanParserADDITION)|(1<<TreeLanParserSUBSTRACTION)|(1<<TreeLanParserMULTIPLICATION)|(1<<TreeLanParserDIVISION)|(1<<TreeLanParserEXPONENTIAL)|(1<<TreeLanParserLOGARITHM)|(1<<TreeLanParserSIN)|(1<<TreeLanParserCOS)|(1<<TreeLanParserTAN)|(1<<TreeLanParserTRANSFORMATION)|(1<<TreeLanParserEQUATION)|(1<<TreeLanParserLEAF)|(1<<TreeLanParserANY))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
