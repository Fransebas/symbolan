package Node

import (
	"fmt"
	"math"
	"strconv"
	"symbolan/OperationClass"
	"symbolan/ValueClass"
	"symbolan/parser"
	"symbolan/utils"
)

const delta = 0.000001

type Node struct {
	IsLeaf           bool
	Value            string
	classByValues    ValueClass.ValueClass
	classByOperation OperationClass.OperationClass
	numericValue     float64

	Left  *Node
	Right *Node

	IsOperation bool
	Operation   string

	height   int64
	treeSize int64
}

func isRule(class ValueClass.ValueClass) bool {
	return class == ValueClass.DERIVATIVE_RULE ||
		class == ValueClass.NUMERIC_RULE ||
		class == ValueClass.CONST_RULE ||
		class == ValueClass.VAR_RULE ||
		class == ValueClass.GENERIC_RULE ||
		class == ValueClass.EXPR_GENERIC_RULE ||
		class == ValueClass.EXPR_CONSTANT_RULE ||
		class == ValueClass.EXPR_VARIABLE_RULE ||

		class == ValueClass.CUSTOM_FUNCTION_RULE ||

		class == ValueClass.NUMERIC_TREE_RULE ||
		class == ValueClass.CONST_TREE_RULE ||
		class == ValueClass.VAR_TREE_RULE ||
		class == ValueClass.FUNCTION_TREE_RULE
}

func isTreeRule(class ValueClass.ValueClass) bool {
	return class == ValueClass.NUMERIC_TREE_RULE ||
		class == ValueClass.CONST_TREE_RULE ||
		class == ValueClass.VAR_TREE_RULE ||
		class == ValueClass.FUNCTION_TREE_RULE
}

func (this *Node) IsTreeRule() bool {
	return isTreeRule(this.classByValues)
}

func (this *Node) IsRule() bool {
	return isRule(this.classByValues)
}

func (this *Node) SetNumericValue(value float64) {
	this.numericValue = value
	this.setValueFromNumeric()
}

func (this *Node) IsNegativeRule() bool {
	if this.IsRule() {
		return this.Value[:1] == "!"
	}
	return false
}

func (this *Node) Height() int64 {
	return this.height
}

func (this *Node) TreeSize() int64 {
	return this.treeSize
}

func (this *Node) OperationClass() OperationClass.OperationClass {
	return this.classByOperation
}

func (this *Node) ValueClass() ValueClass.ValueClass {
	return this.classByValues
}

func (this *Node) NumericValue() float64 {
	if this.classByValues != ValueClass.NUMERIC_CONSTANT {
		panic("Can't get numeric value from non numeric constant")
	}
	return this.numericValue
}

func (this *Node) PrintTree(tabs string) {
	value := this.Value
	if this.OperationClass() == OperationClass.MULTIPLICATION {
		value = "*"
	}
	fmt.Println(tabs + value)
	if this.IsLeaf {
		return
	}
	this.Left.PrintTree(tabs + "_")
	this.Right.PrintTree(tabs + "_")
}

func (this *Node) String() string {
	if this.IsLeaf {
		return this.Value
	}
	if this.classByOperation == OperationClass.SYSTEM_FUNCTION {
		return this.Left.String() + "(" + this.Right.String() + ")"
	} else {
		if this.OperationClass() == OperationClass.ADDITION {
			return this.Left.String() + this.Operation + this.Right.String()
		} else if this.Right != nil {
			operation := this.Operation
			if this.OperationClass() == OperationClass.SIGN {
				operation = ""
			}
			left := ""
			if this.Left.treeSize > 1 {
				left = fmt.Sprintf("(%v)", this.Left.String())
			} else {
				left = fmt.Sprintf("%v", this.Left.String())
			}
			right := ""
			if this.Right.treeSize > 1 {
				right = fmt.Sprintf("(%v)", this.Right.String())
			} else {
				right = fmt.Sprintf("%v", this.Right.String())
			}
			return left + operation + right
		} else {
			return this.Left.String() + this.Operation
		}
	}

}

//func Substitute(node *Node, valueSubstitution map[string]string, symbolanProcessor *SymbolanProcessor) *Node {
//	processedSubstitutions := make(map[string]*Node)
//	for key, value := range valueSubstitution {
//		processedSubstitutions[key] = symbolanProcessor.ProcessString(value)
//	}
//	newNode := node.Eval(&processedSubstitutions)
//	newNode.Simplify(symbolanProcessor.config.Simplify)
//	return &newNode
//}

func (this Node) Eval(valueSubstitution *map[string]*Node) Node {
	nNode := this.eval(valueSubstitution)
	nNode.calculateTreeClassByValue()
	return nNode
}

func (this Node) eval(valueSubstitution *map[string]*Node) Node {
	if this.IsLeaf {
		// Do not substitute numbers, even though that could be interesting
		if ValueClass.IsNumeric(this.classByValues) {
			return this
		}
		if val, ok := (*valueSubstitution)[this.Value]; ok {
			return *val
		}
	} else {
		left := this.Left.eval(valueSubstitution)
		this.Left = &left
		if this.Right != nil {
			right := this.Right.eval(valueSubstitution)
			this.Right = &right
		}
	}
	return this
}

func (this *Node) CalculateTreeHeightAndSize() {
	if this.IsLeaf {
		this.CalculateHeightAndSize()
	} else {
		this.Left.CalculateHeightAndSize()
		this.Right.CalculateHeightAndSize()

		this.CalculateHeightAndSize()
	}
}

func (this *Node) CalculateHeightAndSize() {
	if this.IsLeaf {
		this.height = 0
		this.treeSize = 1
	} else {
		this.height = utils.Max(this.Left.height, this.Right.height) + 1
		this.treeSize = this.Left.treeSize + this.Right.treeSize + 1
	}
}

func (this *Node) Substitute(substituteNode *Node) {
	this.Value = substituteNode.Value
	this.Left = substituteNode.Left
	this.Right = substituteNode.Right
	this.IsLeaf = substituteNode.IsLeaf
	this.classByOperation = substituteNode.classByOperation
	this.classByValues = substituteNode.classByValues
	this.numericValue = substituteNode.numericValue
}

// Here the node/tree would be ( D ) f ( x )
func NewNodeDerivatedVar(node *Node) *Node {
	newNode := NewNode()

	newNode.classByValues = ValueClass.DERIVATIVE
	newNode.IsLeaf = true

	newNode.treeSize = 1
	newNode.height = 0

	newNode.Value = "d" + node.Right.Value

	return newNode
}

func NewSignNode(sign string) *Node {
	node := NewNode()
	node.height = 0
	node.treeSize = 1
	node.classByValues = ValueClass.SIGN
	node.IsLeaf = true
	node.Value = sign

	return node
}

func NewNodeFromValue(value float64) *Node {

	node := NewNode()

	node.SetNumericValue(value)
	node.classByValues = ValueClass.NUMERIC_CONSTANT
	node.IsLeaf = true

	node.treeSize = 1
	node.height = 0

	if value < 0 {
		node.SetNumericValue(math.Abs(value))

		expNode := NewNode()
		expNode.Right = node
		expNode.Left = NewSignNode("-")
		expNode.classByValues = ValueClass.NUMERIC_CONSTANT
		expNode.treeSize = 3
		expNode.height = 1
		expNode.classByOperation = OperationClass.SIGN
		expNode.Operation = "sign"
		return expNode
	} else {
		return node
	}

}

func NewNode() *Node {
	node := new(Node)
	node.IsLeaf = false
	return node
}

func getLeafClass(rule int) ValueClass.ValueClass {
	switch rule {
	case parser.SymbolanParserRULE_number:
		return ValueClass.NUMERIC_CONSTANT
	case parser.SymbolanParserRULE_variable:
		return ValueClass.VARIABLE
	case parser.SymbolanParserRULE_constant:
		return ValueClass.CONSTANT
	case parser.SymbolanParserRULE_system_functions:
		return ValueClass.SYSTEM_FUNCTION
	case parser.SymbolanParserRULE_numeric_rule:
		return ValueClass.NUMERIC_RULE
	case parser.SymbolanParserRULE_imaginary:
		return ValueClass.IMAGINARY_CONSTANT
	case parser.SymbolanParserRULE_math_constant:
		return ValueClass.MATH_SYMBOL
	case parser.SymbolanParserRULE_const_rule:
		return ValueClass.CONST_RULE
	case parser.SymbolanParserRULE_var_rule:
		return ValueClass.VAR_RULE
	case parser.SymbolanParserRULE_expr_rule:
		return ValueClass.GENERIC_RULE
	case parser.SymbolanParserRULE_derivative:
		return ValueClass.DERIVATIVE
	case parser.SymbolanParserRULE_derivative_rule:
		return ValueClass.DERIVATIVE_RULE

	case parser.SymbolanParserRULE_numeric_tree_rule:
		return ValueClass.NUMERIC_TREE_RULE
	case parser.SymbolanParserRULE_const_tree_rule:
		return ValueClass.CONST_TREE_RULE
	case parser.SymbolanParserRULE_var_tree_rule:
		return ValueClass.VAR_TREE_RULE
	case parser.SymbolanParserRULE_system_tree_rule:
		return ValueClass.FUNCTION_TREE_RULE

	case parser.SymbolanParserRULE_sign:
		return ValueClass.SIGN
	default:
		panic(fmt.Sprintf("Rule %v should't exist", rule))
	}
}

func isConstantRule(node *Node) bool {
	if node == nil {
		return false
	}
	return ValueClass.IsConstantRule(node.classByValues)
}

func isTreeNumericRule(node *Node) bool {
	if node == nil {
		return false
	}
	return ValueClass.IsTreeNumeric(node.classByValues)
}

func isTreeConstatnRule(node *Node) bool {
	if node == nil {
		return false
	}
	return ValueClass.IsTreeConstant(node.classByValues)
}

func isTreeVariableRule(node *Node) bool {
	if node == nil {
		return false
	}
	return ValueClass.IsTreeVar(node.classByValues)
}

func isTreeFuntionRule(node *Node) bool {
	if node == nil {
		return false
	}
	return ValueClass.IsTreeFunction(node.classByValues)
}

func isGenericRule(node *Node) bool {
	if node == nil {
		return false
	}
	return ValueClass.IsGenericRule(node.classByValues)
}

func isVariableRule(node *Node) bool {
	if node == nil {
		return false
	}
	return ValueClass.IsVariableRule(node.classByValues)
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
	this.numericValue, err = strconv.ParseFloat(this.Value, 64)
	if err != nil {
		panic(err)
	}
}

func (this *Node) setValueFromNumeric() {
	this.Value = fmt.Sprintf("%v", this.numericValue)
}

// Calculate the class by value of the node or in other words by leaf,
// to do that, first the class of the subnodes should be already calculated
func (this *Node) calculateTreeClassByValue() {

	if this.IsLeaf {
		if this.classByValues == ValueClass.NUMERIC_CONSTANT {
			this.setNumericValueFromValue()
		}
	} else {
		this.Left.calculateTreeClassByValue()
		if this.Right != nil {
			this.Right.calculateTreeClassByValue()
		}

		this.CalculateClassByValue()
	}
}

func (this *Node) CalculateClassByValue() {
	if !this.IsLeaf {
		if isGenericRule(this.Left) || isGenericRule(this.Right) {
			this.classByValues = ValueClass.EXPR_GENERIC_RULE

		} else if isTreeNumericRule(this.Left) || isTreeNumericRule(this.Right) {
			this.classByValues = ValueClass.EXPR_NUMERIC_TREE_RULE
		} else if isTreeConstatnRule(this.Left) || isTreeConstatnRule(this.Right) {
			this.classByValues = ValueClass.EXPR_CONST_TREE_RULE
		} else if isTreeVariableRule(this.Left) || isTreeVariableRule(this.Right) {
			this.classByValues = ValueClass.EXPR_VAR_TREE_RULE
		} else if isTreeFuntionRule(this.Left) || isTreeFuntionRule(this.Right) {
			this.classByValues = ValueClass.EXPR_FUNCTION_TREE_RULE

		} else if isVariableRule(this.Left) || isVariableRule(this.Right) {
			this.classByValues = ValueClass.EXPR_VARIABLE_RULE
		} else if isConstantRule(this.Left) || isConstantRule(this.Right) {
			this.classByValues = ValueClass.EXPR_CONSTANT_RULE
		} else if isVariable(this.Left) || isVariable(this.Right) {
			this.classByValues = ValueClass.VARIABLE_EXPRESSION
		} else if isConstant(this.Left) || isConstant(this.Right) {
			this.classByValues = ValueClass.CONSTANT_EXPRESSION
		} else if isNumeric(this.Left) || isNumeric(this.Right) {
			this.classByValues = ValueClass.NUMERIC_CONSTANT_EXPRESSION
		} else {
			panic("This case should happen")
		}
	}
}

var systemFunction = map[string]bool{
	"min":  true,
	"max":  true,
	"sin":  true,
	"cos":  true,
	"tan":  true,
	"log":  true,
	"sqrt": true,
}

func isSystemFunction(op string) bool {
	return systemFunction[op]
}

var OperationToEnum = map[string]OperationClass.OperationClass{
	"+":  OperationClass.ADDITION,
	"-":  OperationClass.SUBSTRACTION,
	"*":  OperationClass.MULTIPLICATION,
	"":   OperationClass.MULTIPLICATION,
	"/":  OperationClass.DIVISION,
	"^":  OperationClass.EXPONENTIAL,
	"**": OperationClass.EXPONENTIAL,
	"=":  OperationClass.EQUATION,
}

var OperationSimplifiedOnLeftOne = map[string]OperationClass.OperationClass{
	"*":  OperationClass.MULTIPLICATION,
	"":   OperationClass.MULTIPLICATION,
	"^":  OperationClass.EXPONENTIAL,
	"**": OperationClass.EXPONENTIAL,
}

var OperationSimplifiedOnRightOne = map[string]OperationClass.OperationClass{
	"*":  OperationClass.MULTIPLICATION,
	"":   OperationClass.MULTIPLICATION,
	"^":  OperationClass.EXPONENTIAL,
	"**": OperationClass.EXPONENTIAL,
}

func getOperationClass(op string) OperationClass.OperationClass {
	return OperationToEnum[op]
}

func (this *Node) Compare(node *Node) bool {

	if this.IsRule() || node.IsRule() {
		panic("Comparing rules with this functions is not supported")
	}

	if this.height != node.height || this.treeSize != node.treeSize {
		return false
	}

	if this.OperationClass() != node.OperationClass() {
		return false
	}

	if this.ValueClass() != node.ValueClass() {
		return false
	}

	if this.IsLeaf && node.IsLeaf {
		sameClass := this.ValueClass() == node.ValueClass()
		sameNumericValue := utils.Compare(this.numericValue, node.numericValue)
		sameValue := this.Value == node.Value

		if !sameClass {
			return false
		}

		if this.ValueClass() == ValueClass.NUMERIC_CONSTANT {
			return sameNumericValue
		} else {
			return sameValue
		}
	} else {
		return this.Left.Compare(node.Left) && this.Right.Compare(node.Right)
	}
}
