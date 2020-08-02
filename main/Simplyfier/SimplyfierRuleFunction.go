package Simplyfier

import (
	"math"
	"symbolan/main/Tree"
)

func addition(node *Tree.Node) Tree.Node {
	return *Tree.NewNodeFromValue(node.Left.NumericValue() + node.Right.NumericValue())
}

func subtraction(node *Tree.Node) Tree.Node {
	return *Tree.NewNodeFromValue(node.Left.NumericValue() - node.Right.NumericValue())
}

func multiplication(node *Tree.Node) Tree.Node {
	return *Tree.NewNodeFromValue(node.Left.NumericValue() + node.Right.NumericValue())
}

func division(node *Tree.Node) Tree.Node {
	return *Tree.NewNodeFromValue(node.Left.NumericValue() / node.Right.NumericValue())
}

func exponential(node *Tree.Node) Tree.Node {
	return *Tree.NewNodeFromValue(math.Pow(node.Left.NumericValue(), node.Right.NumericValue()))
}
