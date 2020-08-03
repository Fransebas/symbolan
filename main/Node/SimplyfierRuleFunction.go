package Node

import (
	"math"
)

func addition(node *Node) Node {
	return *NewNodeFromValue(node.Left.NumericValue() + node.Right.NumericValue())
}

func subtraction(node *Node) Node {
	return *NewNodeFromValue(node.Left.NumericValue() - node.Right.NumericValue())
}

func multiplication(node *Node) Node {
	return *NewNodeFromValue(node.Left.NumericValue() + node.Right.NumericValue())
}

func division(node *Node) Node {
	return *NewNodeFromValue(node.Left.NumericValue() / node.Right.NumericValue())
}

func exponential(node *Node) Node {
	return *NewNodeFromValue(math.Pow(node.Left.NumericValue(), node.Right.NumericValue()))
}
