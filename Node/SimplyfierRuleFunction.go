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
	return *NewNodeFromValue(node.Left.NumericValue() * node.Right.NumericValue())
}

func division(node *Node) Node {
	return *NewNodeFromValue(node.Left.NumericValue() / node.Right.NumericValue())
}

func exponential(node *Node) Node {
	return *NewNodeFromValue(math.Pow(node.Left.NumericValue(), node.Right.NumericValue()))
}

func GetNumberFromSigned(node *Node) float64 {
	if !node.IsLeaf {
		isNegative := node.Left.Value == "-"
		if isNegative {
			return -1 * node.Right.NumericValue()
		}
		return node.Right.NumericValue()
	}
	return node.NumericValue()
}

func signedExponential(node *Node) Node {
	base := GetNumberFromSigned(node.Left)
	exp := GetNumberFromSigned(node.Right)
	return *NewNodeFromValue(math.Pow(base, exp))
}

func VarDerivative(node *Node) Node {
	return *NewNodeDerivatedVar(node)
}
