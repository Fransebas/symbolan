package Tree

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"symbolan/parser"
)

type TreeShapeListener struct {
	*parser.BaseSymbolanListener
	stack  []*Node
	Root   *Node
	config *TreeParsingConfig
}

type SimplifyEnum string

// Each level covers the previous
const (
	NUMERIC_BASIC_OPS_1 SimplifyEnum = "NUMERIC_BASIC_OPS"
	NUMERIC_FUNCTIONS_2 SimplifyEnum = "NUMERIC_FUNCTIONS"
)

type TreeParsingConfig struct {
	Simplify SimplifyEnum
}

func (this *TreeShapeListener) pop() *Node {
	node := this.peek()
	this.stack = this.stack[:len(this.stack)-1]
	return node
}

func (this *TreeShapeListener) peek() *Node {
	return this.stack[len(this.stack)-1]
}

func (this *TreeShapeListener) isStackEmpty() bool {
	return len(this.stack) == 0
}

func NewTreeShapeListener(config *TreeParsingConfig) *TreeShapeListener {
	treeListener := new(TreeShapeListener)
	treeListener.stack = make([]*Node, 0)
	treeListener.config = config
	return treeListener
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	childrenCount := ctx.GetChildCount()
	if childrenCount > 4 {
		panic("Node shouldn't have move than 3 children")
	}
	this.stack = append(this.stack, NewNode(&ctx))
}

func isParethesisNode(node *Node) bool {
	return node.Value == "(" || node.Value == ")"
}

var operations = map[string]bool{
	"+":  true,
	"-":  true,
	"*":  true,
	"/":  true,
	"^":  true,
	"**": true,
}

func isOperation(operation string) bool {
	return operations[operation]
}

func (this *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	node := this.pop()
	if this.isStackEmpty() {
		this.Root = node
		if node.right == nil {
			this.Root = node.left
		}
		this.onRootCreated()
	} else {
		parentNode := this.peek()

		// Parethesis Nodes are useless
		if !isParethesisNode(node) {
			parentNode.AddChild(getNodeToKeep(node))
		}
	}
}

func (this *TreeShapeListener) onRootCreated() {
	this.Root.calculateClassByValue()
	this.Root.calculateClassByOperation()
	this.Root.Simplify(this.config.Simplify)
}

// Leaf nodes
func (this *TreeShapeListener) VisitTerminal(leaf antlr.TerminalNode) {
	node := this.peek()
	node.IsLeaf = true
	if isOperation(leaf.GetText()) {
		node.IsOperation = true
		node.IsLeaf = false
		node.Operation = leaf.GetText()
	} else {
		node.Value = leaf.GetText()
	}
}
