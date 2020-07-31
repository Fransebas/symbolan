package TreeLan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"symbolan/TreeLan/TreeLanParser"
)

type NodeType string

const (
	ROOT_NODE = "ROOT_NODE"
	LEAF_NODE = "LEAF_NODE"
	TREE_RULE = "TREE_RULE"
	FORMULA   = "FORMULA"
)

type TreeLanNode struct {
	ctx    *antlr.Tree
	IsLeaf bool
	Value  string

	ID string

	Type NodeType

	left  *TreeLanNode
	main  *TreeLanNode
	right *TreeLanNode
}

func NewTreeLanNode(ctx *antlr.Tree) *TreeLanNode {
	node := new(TreeLanNode)
	node.ctx = ctx
	node.IsLeaf = false
	return node
}

func ProcessString(data string) *TreeLanNode {

	input := antlr.NewInputStream(data)
	lexer := parser.NewTreeLanLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewTreeLanParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	antrlTree := p.TreeRule()

	ruleNode := getRuleNode(antrlTree)

	return ruleNode
}

func getRuleNode(ctx antlr.Tree) *TreeLanNode {
	node := NewTreeLanNode(&ctx)

	// Formula on the right
	if ctx.GetChildCount() == 4 {
		node.left = getTreeDescriptorNode(ctx.GetChild(0))
		node.right = getFormulaNode(ctx.GetChild(3))
	} else {
		node.left = getTreeDescriptorNode(ctx.GetChild(0))
		node.right = getTreeDescriptorNode(ctx.GetChild(2))

		node.Type = TREE_RULE
	}

	return node
}

func getFormulaNode(ctx antlr.Tree) *TreeLanNode {
	node := NewTreeLanNode(&ctx)

	node.main = getFormula(ctx.GetChild(0))

	node.Type = FORMULA

	return node
}

func getFormula(ctx antlr.Tree) *TreeLanNode {
	node := NewTreeLanNode(&ctx)

	node.IsLeaf = true
	node.ID = (ctx.(antlr.TerminalNode)).GetText()

	return node
}

func getTreeDescriptorNode(ctx antlr.Tree) *TreeLanNode {
	node := NewTreeLanNode(&ctx)

	if ctx.GetChildCount() == 1 {
		node.main = getLeafNode(ctx.GetChild(0))
		node.Type = LEAF_NODE
	} else {
		node.left = getTreeDescriptorNode(ctx.GetChild(1))
		node.main = getRootNode(ctx.GetChild(3))
		node.right = getTreeDescriptorNode(ctx.GetChild(5))
		node.Type = ROOT_NODE
	}

	return node
}

func getRootNode(ctx antlr.Tree) *TreeLanNode {
	node := NewTreeLanNode(&ctx)

	node.IsLeaf = true
	node.Value = getOperationClass(ctx.GetChild(0))
	node.ID = getID(ctx.GetChild(2))

	return node
}

func getLeafNode(ctx antlr.Tree) *TreeLanNode {
	node := NewTreeLanNode(&ctx)

	node.IsLeaf = true
	node.Value = getClassValue(ctx.GetChild(0))
	node.ID = getID(ctx.GetChild(2))

	return node
}

func getOperationClass(ctx antlr.Tree) string {
	return getID(ctx.GetChild(0))
}

func getClassValue(ctx antlr.Tree) string {
	return getID(ctx.GetChild(0))
}

func getID(ctx antlr.Tree) string {
	return (ctx.(antlr.TerminalNode)).GetText()
}
