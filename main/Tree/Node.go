package Tree

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"math"
	"strconv"
	"symbolan/main/OperationClass"
	"symbolan/main/ValueClass"
	"symbolan/parser"
)

const delta = 0.000001

type Node struct {
	ctx              *antlr.ParserRuleContext
	IsLeaf           bool
	Value            string
	classByValues    string
	classByOperation string
	numericValue     float64

	left  *Node
	right *Node

	IsOperation bool
	Operation   string
}

func (this *Node) NumericValue() float64 {
	if this.classByValues != ValueClass.NUMERIC_CONSTANT {
		panic("Can't get numeric value from non numeric constant")
	}
	return this.numericValue
}

func (this *Node) String() string {
	if this.IsLeaf {
		return this.Value
	}

	if this.classByOperation == OperationClass.SYSTEM_FUNCTION {
		return this.left.String() + "(" + this.right.String() + ")"
	} else {
		if this.right != nil {
			return this.left.String() + this.Operation + this.right.String()
		} else {
			return this.left.String() + this.Operation
		}
	}

}

func Substitute(node *Node, valueSubstitution map[string]string, config *TreeParsingConfig) *Node {
	processedSubstitutions := make(map[string]*Node)
	for key, value := range valueSubstitution {
		processedSubstitutions[key] = ProcessString(value, config)
	}
	node.Eval(processedSubstitutions)
	node.Simplify(config.Simplify)
	return node
}

func (this *Node) Eval(valueSubstitution map[string]*Node) {
	if this.IsLeaf {
		// Do not subtitute number, even though that could be interesting
		if ValueClass.IsNumeric(this.classByValues) {
			return
		}
		if val, ok := valueSubstitution[this.Value]; ok {
			this.Substitute(val)
		}
	} else {
		this.left.Eval(valueSubstitution)
		if this.right != nil {
			this.right.Eval(valueSubstitution)
		}
	}
}

func performFunctionOperation(function string, rightValue float64) float64 {

	switch function {
	case "sin":
		return math.Sin(rightValue)
	case "cos":
		return math.Cos(rightValue)
	case "tan":
		return math.Tan(rightValue)
	case "log":
		return math.Log(rightValue)
	default:
		panic(fmt.Sprintf("%v is not a valid system function", function))
	}
}

func performOperation(op string, leftValue, rightValue float64) float64 {

	switch op {
	case "+":
		return leftValue + rightValue
	case "-":
		return leftValue - rightValue
	case "*":
		return leftValue * rightValue
	case "":
		return leftValue * rightValue
	case "/":
		return leftValue / rightValue
	case "^":
		return math.Pow(leftValue, rightValue)
	case "**":
		return math.Pow(leftValue, rightValue)
	default:
		panic(fmt.Sprintf("%v is not a valid numeric operation", op))
	}
}

func (this *Node) simplifyBasicOps() {
	if this.left.classByValues == ValueClass.NUMERIC_CONSTANT && this.right.classByValues == ValueClass.NUMERIC_CONSTANT {
		this.numericValue = performOperation(this.Operation, this.left.NumericValue(), this.right.NumericValue())
		this.Value = fmt.Sprintf("%f", this.numericValue)

		this.left = nil
		this.right = nil
		this.classByValues = ValueClass.NUMERIC_CONSTANT
		this.IsOperation = false
		this.IsLeaf = true
		this.Operation = ""
		this.ctx = nil
	}
}

func isEqual(aNumber, bNumber float64) bool {
	return aNumber > bNumber-delta || aNumber < bNumber+delta
}

//Remove multiplications by 0 and multiplication by 1 and sum by 0
func (this *Node) simplifyZeroes() {
	isLeftZero := this.left.classByValues == ValueClass.NUMERIC_CONSTANT && isEqual(this.left.numericValue, 0)
	isRightZero := this.right.classByValues == ValueClass.NUMERIC_CONSTANT && isEqual(this.right.numericValue, 0)
	if isLeftZero || isRightZero {
		this.numericValue = 0
		this.Value = "0"

		this.left = nil
		this.right = nil
		this.classByValues = ValueClass.NUMERIC_CONSTANT
		this.IsOperation = false
		this.IsLeaf = true
		this.Operation = ""
		this.ctx = nil
	}
}


//Remove multiplications by 0 and multiplication by 1 and sum by 0
func (this *Node) simplifyOnes() {
	isLeftOne := this.left.classByValues == ValueClass.NUMERIC_CONSTANT && isEqual(this.left.numericValue, 1)
	isRightOne := this.right.classByValues == ValueClass.NUMERIC_CONSTANT && isEqual(this.right.numericValue, 1)

	if isLeftOne || isRightOne {
		if
		this.numericValue = 0
		this.Value = "0"

		this.left = nil
		this.right = nil
		this.classByValues = ValueClass.NUMERIC_CONSTANT
		this.IsOperation = false
		this.IsLeaf = true
		this.Operation = ""
		this.ctx = nil
	}
}

func (this *Node) simplifyNumericFunctions() {

	if this.left == nil {
		if this.IsLeaf {
			fmt.Println("WTF")
		}
		fmt.Println(this.Value)
	}
	if this.right == nil {
		fmt.Println("right " + this.Value)
	}
	if this.left.classByValues == ValueClass.SYSTEM_FUNCTION && this.right.classByValues == ValueClass.NUMERIC_CONSTANT {
		this.numericValue = performFunctionOperation(this.left.Value, this.right.NumericValue())
		this.Value = fmt.Sprintf("%f", this.numericValue)

		this.left = nil
		this.right = nil
		this.classByValues = ValueClass.NUMERIC_CONSTANT
		this.IsOperation = false
		this.IsLeaf = true
		this.Operation = ""
		this.ctx = nil
	}
}

// Modifies the tree
func (this *Node) Simplify(simplifyLevel SimplifyEnum) {

	// there is nothing to simplify
	if this.IsLeaf {
		return
	}

	this.left.Simplify(simplifyLevel)
	this.right.Simplify(simplifyLevel)

	if simplifyLevel == NUMERIC_BASIC_OPS_1 {
		this.simplifyBasicOps()
	}

	if simplifyLevel == NUMERIC_FUNCTIONS_2 {
		this.simplifyBasicOps()
		if !this.IsLeaf {
			this.simplifyNumericFunctions()
		}
	}
}

func (this *Node) Substitute(substituteNode *Node) {
	this.ctx = substituteNode.ctx
	this.Value = substituteNode.Value
	this.left = substituteNode.left
	this.right = substituteNode.right
	this.IsLeaf = substituteNode.IsLeaf
	this.classByOperation = substituteNode.classByOperation
	this.classByValues = substituteNode.classByValues
}

func NewNode(ctx *antlr.ParserRuleContext) *Node {
	node := new(Node)
	node.ctx = ctx
	node.IsLeaf = false
	return node
}

func getLeafClass(rule int) string {
	switch rule {
	case parser.SymbolanParserRULE_number:
		return ValueClass.NUMERIC_CONSTANT
	case parser.SymbolanParserRULE_variable:
		return ValueClass.VARIABLE
	case parser.SymbolanParserRULE_constant:
		return ValueClass.CONSTANT
	case parser.SymbolanParserRULE_system_functions:
		return ValueClass.SYSTEM_FUNCTION
	default:
		panic(fmt.Sprintf("Rule %v should't exist", rule))
	}
}

func isNumeric(node *Node) bool {
	if node == nil {
		return false
	}
	return ValueClass.IsNumeric(node.classByValues)
}

func isVariable(node *Node) bool {
	if node == nil {
		return false
	}
	return ValueClass.IsVariable(node.classByValues)
}

func isConstant(node *Node) bool {
	if node == nil {
		return false
	}
	return ValueClass.IsConstant(node.classByValues)
}

func (this *Node) setNumericValueFromValue() {
	var err error
	this.numericValue, err = strconv.ParseFloat(this.Value, 64);
	if err != nil {
		panic(err)
	}
}

// Calculate the class by value of the node or in other words by leaf,
// to do that, first the class of the subnodes should be already calculated
func (this *Node) calculateClassByValue() {

	if this.IsLeaf {
		rule := (*this.ctx).GetRuleIndex()
		this.classByValues = getLeafClass(rule)

		if this.classByValues == ValueClass.NUMERIC_CONSTANT {
			this.setNumericValueFromValue()
		}
	} else {
		this.left.calculateClassByValue()
		if this.right != nil {
			this.right.calculateClassByValue()
		}

		if isVariable(this.left) || isVariable(this.right) {
			this.classByValues = ValueClass.VARIABLE_EXPRESSION
		} else if isConstant(this.left) || isConstant(this.right) {
			this.classByValues = ValueClass.CONSTANT_EXPRESSION
		} else if isNumeric(this.left) || isNumeric(this.right) {
			this.classByValues = ValueClass.NUMERIC_CONSTANT_EXPRESSION
		} else {
			panic("This case should happen");
		}
	}
}

var systemFunction = map[string]bool{
	"min": true,
	"max": true,
	"sin": true,
	"cos": true,
	"tan": true,
	"log": true,
}

func isSystemFunction(op string) bool {
	return systemFunction[op]
}

var OperationToEnum = map[string]string{
	"+" : OperationClass.ADDITION,
	"-":  OperationClass.SUBSTRACTION,
	"*":  OperationClass.MULTIPLICATION,
	"":  OperationClass.MULTIPLICATION,
	"/":  OperationClass.DIVISION,
	"^":  OperationClass.EXPONENTIAL,
	"**":  OperationClass.EXPONENTIAL,
	"=":  OperationClass.EQUATION,
}

var OperationSimplifiedOnLeftOne = map[string]string{
	"*":  OperationClass.MULTIPLICATION,
	"":  OperationClass.MULTIPLICATION,
	"^":  OperationClass.EXPONENTIAL,
	"**":  OperationClass.EXPONENTIAL,
}

var OperationSimplifiedOnRightOne = map[string]string{
	"*":  OperationClass.MULTIPLICATION,
	"":  OperationClass.MULTIPLICATION,
	"^":  OperationClass.EXPONENTIAL,
	"**":  OperationClass.EXPONENTIAL,
}

func getOperationClass(op string) string {
	return OperationToEnum[op]
}

// Calculate the class of the node by operation or in other words by root value,
// to do that, first the class of the subnodes should be already calculated
func (this *Node) calculateClassByOperation() {

	if this.IsLeaf {
		this.classByOperation = OperationClass.LEAF
	} else {
		this.left.calculateClassByOperation()
		if this.right != nil {
			this.right.calculateClassByOperation()
		}
		// Edge case, this is for the root
		if this.right == nil {
			this.classByOperation = this.left.classByOperation
		} else if (*this.left.ctx).GetRuleIndex() == parser.SymbolanParserRULE_system_functions {
			this.classByOperation = OperationClass.SYSTEM_FUNCTION
		} else {
			this.classByOperation = getOperationClass(this.Operation)
		}
	}
}

func (this *Node) AddChild(node *Node) {
	if this.left == nil {
		this.left = node
	} else if this.right == nil {
		this.right = node
	} else {
		panic("Tree should be a binary Tree")
	}
}

// if the node only has one children, then we dont need this node,
// we only need its child
func getNodeToKeep(node *Node) *Node {
	if node.IsLeaf {
		return node
	}

	// this means that it only has one child, left
	if node.right == nil {
		return node.left
	} else if node.left == nil {
		panic("This sceneario shouldn't happen")
	}
	return node
}
