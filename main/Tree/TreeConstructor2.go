package Tree

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
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

func (this *SymbolanProcessor) ProcessString(data string) *Node {

	input := antlr.NewInputStream(data)
	lexer := parser.NewSymbolanLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSymbolanParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	antrlTree := p.Eq()

	this.Root = ExprNode(antrlTree)
	this.initializeRoot()

	return this.Root
}

func (this *SymbolanProcessor) initializeRoot() {
	this.Root.calculateClassByValue()
	this.Root.calculateClassByOperation()
	this.Root.Simplify(this.config.Simplify)
}

func EqNode(ctx antlr.Tree) *Node {
	node := NewNode(&ctx)

	node.left = ExprNode(ctx.GetChild(0))
	node.Value = ctx.GetChild(1).(antlr.TerminalNode).GetText()
	node.right = ExprNode(ctx.GetChild(2))

	return node
}

func isParenthesis(ctx antlr.Tree) bool {
	return ctx.(antlr.ParserRuleContext).GetRuleIndex() == parser.SymbolanParserRULE_left_par
}

func ExprNode(ctx antlr.Tree) *Node {
	node := NewNode(&ctx)

	// Function
	if ctx.GetChildCount() == 4 {
		node.left = FuntionNode(ctx.GetChild(0))
		node.right = ExprNode(ctx.GetChild(2))
	} else if ctx.GetChildCount() == 3 {
		if isParenthesis(ctx.GetChild(0)) {
			return ExprNode(ctx.GetChild(1))
		} else {
			node.left = FuntionNode(ctx.GetChild(0))
			node.Value = ctx.GetChild(1).(antlr.TerminalNode).GetText()
			node.right = ExprNode(ctx.GetChild(2))
			return node
		}
	} else if ctx.GetChildCount() == 2 {
		node.left = ExprNode(ctx.GetChild(0))
		node.right = ExprNode(ctx.GetChild(1))
	} else if ctx.GetChildCount() == 1 {
		node.left = ProcessLeaf(ctx.GetChild(0), ctx)
	} else {
		panic(fmt.Printf("Expr node with %v children should never happen", ctx.GetChildCount()))
	}

	return node
}

func FuntionNode(ctx antlr.Tree) *Node {
	subNodeCtx := ctx.GetChild(0).GetChild(0)
	node := NewNode(&subNodeCtx)
	node.Value = subNodeCtx.(antlr.TerminalNode).GetText()

	return node
}

func ProcessLeaf(ctx antlr.Tree, parentCtx antlr.Tree) *Node {
	if ctx.GetChildCount() == 1 {
		return ProcessLeaf(ctx.GetChild(0), ctx)
	} else if ctx.GetChildCount() == 0 {
		leafCtx := ctx.GetChild(0)
		leafNode := NewNode(&parentCtx)
		leafNode.Value = leafCtx.(antlr.TerminalNode).GetText()
		leafNode.IsLeaf = true
		return leafNode
	} else {
		panic(fmt.Printf("Leaf node with %v children should never happen", ctx.GetChildCount()))
	}
}
