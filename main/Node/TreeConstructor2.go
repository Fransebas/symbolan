package Node

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"symbolan/main/OperationClass"
	"symbolan/main/ValueClass"
	"symbolan/parser"
)

type SymbolanProcessor struct {
	*parser.BaseSymbolanListener
	Root   *Node
	config *SymbolanProcessorConfig
}

type SymbolanProcessorConfig struct {
	Simplify SimplifyEnum
}

var DefaultConfig = SymbolanProcessorConfig{
	Simplify: NUMERIC_BASIC_OPS_1,
}

func NewSymbolanProcessor(config *SymbolanProcessorConfig) *SymbolanProcessor {
	symbolanProcessor := new(SymbolanProcessor)

	if config == nil {
		config = &DefaultConfig
	}
	symbolanProcessor.config = config

	return symbolanProcessor
}

func getANTLRparser(data string) *parser.SymbolanParser {
	input := antlr.NewInputStream(data)
	lexer := parser.NewSymbolanLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSymbolanParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	return p

}

func (this *SymbolanProcessor) ProcessString(data string) *Node {

	antlrParser := getANTLRparser(data)

	antrlTree := antlrParser.Expr()
	this.Root = ExprNode(antrlTree)

	this.initializeRoot()

	return this.Root
}

func (this *SymbolanProcessor) ProcessRuleString(data string) *RuleSet {

	antlrParser := getANTLRparser(data)

	antrlTree := antlrParser.RuleSet()

	rules := NewRuleSet(antrlTree)

	for _, rule := range rules.Rules {
		rule.calculateTreeClassByValue()
	}

	return rules
}

func (this *SymbolanProcessor) initializeRoot() {
	this.Root.calculateTreeClassByValue()
	//this.Root.calculateClassByOperation()
	simplyfier := NewSimplyfier()
	node := simplyfier.Simplify(*this.Root)
	this.Root = &node
	this.Root.CalculateHeightAndSize()
}

func EqNode(ctx antlr.Tree) *Node {
	node := NewNode()

	childrenCount := ctx.GetChildCount()
	if childrenCount == 1 {
		return node
	} else if childrenCount == 3 {
		node.Left = ExprNode(ctx.GetChild(0))
		node.Value = ctx.GetChild(1).(antlr.TerminalNode).GetText()
		node.Right = ExprNode(ctx.GetChild(2))
		node.classByOperation = OperationClass.EQUATION
		return node
	} else {
		panic(fmt.Sprintf("Eq node with %v children should never happen", childrenCount))
	}
}

func RuleFunction(ctx antlr.Tree) *Node {
	node := NewNode()

	node.Value = ctx.GetChild(0).(antlr.TerminalNode).GetText()
	node.IsLeaf = true
	node.classByValues = ValueClass.RULE_FUNCTION

	return node
}

func ExprRule(ctx antlr.Tree) *Node {
	node := NewNode()
	childrenCount := ctx.GetChildCount()

	if childrenCount == 3 {
		node.Left = ExprNode(ctx.GetChild(0))
		node.Value = ctx.GetChild(1).(antlr.TerminalNode).GetText()

		if ctx.GetChild(2).(antlr.ParserRuleContext).GetRuleIndex() == parser.SymbolanParserRULE_rule_function {
			node.Right = RuleFunction(ctx.GetChild(2))
		} else {
			node.Right = ExprNode(ctx.GetChild(2))
		}
		node.classByOperation = OperationClass.RULE
		return node
	} else {
		panic(fmt.Sprintf("ExprRule node with %v children should never happen", childrenCount))
	}
}

func isParenthesis(ctx antlr.Tree) bool {
	return ctx.(antlr.ParserRuleContext).GetRuleIndex() == parser.SymbolanParserRULE_left_par
}

func ExprNode(ctx antlr.Tree) *Node {
	node := NewNode()

	// Function
	childrenCount := ctx.GetChildCount()
	if childrenCount == 4 {
		node.classByOperation = OperationClass.SYSTEM_FUNCTION
		node.Left = FunctionNode(ctx.GetChild(0))
		node.Operation = "f"
		node.Right = ExprNode(ctx.GetChild(2))
	} else if childrenCount == 3 {
		if isParenthesis(ctx.GetChild(0)) {
			return ExprNode(ctx.GetChild(1))
		} else {
			node.Left = ExprNode(ctx.GetChild(0))
			node.Operation = ctx.GetChild(1).(antlr.TerminalNode).GetText()
			node.Right = ExprNode(ctx.GetChild(2))

			node.classByOperation = getOperationClass(node.Operation)

			return node
		}
	} else if childrenCount == 2 {
		node.classByOperation = OperationClass.MULTIPLICATION
		node.Left = ExprNode(ctx.GetChild(0))
		node.Right = ExprNode(ctx.GetChild(1))
	} else if childrenCount == 1 {
		return ProcessLeaf(ctx.GetChild(0), ctx)
	} else {
		panic(fmt.Sprintf("Expr node with %v children should never happen", childrenCount))
	}

	return node
}

func FunctionNode(ctx antlr.Tree) *Node {
	subNodeCtx := ctx.GetChild(0)
	node := NewNode()
	node.Value = subNodeCtx.(antlr.TerminalNode).GetText()
	node.IsLeaf = true

	node.classByValues = ValueClass.SYSTEM_FUNCTION

	return node
}

func ProcessLeaf(ctx antlr.Tree, parentCtx antlr.Tree) *Node {
	childrenCount := ctx.GetChildCount()
	if childrenCount == 1 {
		return ProcessLeaf(ctx.GetChild(0), ctx)
	} else if childrenCount == 0 {
		leafNode := NewNode()
		leafNode.Value = ctx.(antlr.TerminalNode).GetText()
		leafNode.IsLeaf = true

		leafNode.classByValues = getLeafClass(parentCtx.(antlr.ParserRuleContext).GetRuleIndex())

		if leafNode.classByValues == ValueClass.NUMERIC_CONSTANT {
			leafNode.setNumericValueFromValue()
		}

		return leafNode
	} else {
		panic(fmt.Sprintf("Leaf node with %v children should never happen", childrenCount))
	}
}
