package Node

import (
	"symbolan/main/ValueClass"
)

type Simplyfier struct {
	manager *Manager
}

func NewSimplyfier() *Simplyfier {
	simplyfier := new(Simplyfier)
	simplyfier.manager = NewSimplyfierManager()
	return simplyfier
}

func constRuleCase(node *Node) bool {
	switch node.ValueClass() {
	case ValueClass.MATH_SYMBOL:
		return true
	case ValueClass.NUMERIC_CONSTANT:
		return true
	case ValueClass.IMAGINARY_CONSTANT:
		return true
	case ValueClass.CONSTANT:
		return true
	default:
		return false
	}
}

func (this *Simplyfier) Simplify(node Node) Node {
	if node.IsLeaf {
		return node
	}

	left := this.Simplify(*node.Left)
	node.Left = &left
	right := this.Simplify(*node.Right)
	node.Right = &right

	// This ensure the information in the current node stays accurate after
	// their children have been simplyfied
	node.CalculateClassByValue()
	node.CalculateHeightAndSize()

	rules := this.manager.GetRules(&node)

	for _, rule := range rules {
		matchRule := rule.node.Left
		resultRule := rule.node.Right
		//fmt.Println(i)
		substitutions, ok := Compare(&node, matchRule)
		if ok {
			if resultRule.ValueClass() == ValueClass.CUSTOM_FUNCTION_RULE {
				return this.manager.functions[resultRule.Value](&node)
			}
			nNode := resultRule.Eval(substitutions)
			return this.Simplify(nNode)
		}
	}

	return node
}

func isNodeKnown(name string, knownNodes *(map[string]*Node)) bool {
	_, ok := (*knownNodes)[name]
	return ok
}

func equalToExisting(name string, node *Node, knownNodes *(map[string]*Node)) bool {
	storedNode := (*knownNodes)[name]
	return storedNode.Compare(node)
}

func Compare(node *Node, rule *Node) (*map[string]*Node, bool) {
	knownNodes := make(map[string]*Node)
	if compare(node, rule, &knownNodes) {
		return &knownNodes, true
	}
	return nil, false
}

func compare(node *Node, rule *Node, knownNodes *map[string]*Node) (res bool) {
	//generic case
	if node.ValueClass() != ValueClass.SIGN && rule.ValueClass() == ValueClass.GENERIC_RULE {
		ruleName := rule.Value

		if !isNodeKnown(ruleName, knownNodes) {
			(*knownNodes)[ruleName] = node
			return true
		}
		return equalToExisting(ruleName, node, knownNodes)
	}

	if !rule.IsTreeRule() && node.OperationClass() != rule.OperationClass() {
		return false
	}

	//base case
	if !rule.IsTreeRule() && !node.IsLeaf && rule.IsLeaf {
		return false
	}

	if node.IsLeaf && rule.IsLeaf {
		res = compareLeafs(node, rule, knownNodes)
		if res {
			return res
		}
	}

	if rule.IsLeaf {
		if rule.IsNegativeRule() {
			return compareNegativeTreeRules(node, rule, knownNodes)
		}
		return compareTreeRules(node, rule, knownNodes)
	} else {
		return compareNonLeafs(node, rule, knownNodes)
	}

	return false
}

func compareLeafs(node *Node, rule *Node, knownNodes *map[string]*Node) bool {
	ruleName := rule.Value
	ruleValueClass := rule.ValueClass()
	switch ruleValueClass {

	case ValueClass.DERIVATIVE_RULE:
		if node.ValueClass() == ValueClass.DERIVATIVE {
			if !isNodeKnown(ruleName, knownNodes) {
				(*knownNodes)[ruleName] = node
				return true
			}
			return equalToExisting(ruleName, node, knownNodes)
		}

	case ValueClass.NUMERIC_RULE:
		if node.ValueClass() == ValueClass.NUMERIC_CONSTANT {
			if !isNodeKnown(ruleName, knownNodes) {
				(*knownNodes)[ruleName] = node
				return true
			}
			return equalToExisting(ruleName, node, knownNodes)
		}

	case ValueClass.CONST_RULE:
		if constRuleCase(node) {
			if !isNodeKnown(ruleName, knownNodes) {
				(*knownNodes)[ruleName] = node
				return true
			}
			return equalToExisting(ruleName, node, knownNodes)
		}
		return false

	case ValueClass.VAR_RULE:
		if node.ValueClass() == ValueClass.VARIABLE {
			if !isNodeKnown(ruleName, knownNodes) {
				(*knownNodes)[ruleName] = node
				return true
			}
			return equalToExisting(ruleName, node, knownNodes)
		}
		return false

	case ValueClass.MATH_SYMBOL:
		return node.Compare(rule)

	case ValueClass.NUMERIC_CONSTANT:
		return node.Compare(rule)

	case ValueClass.IMAGINARY_CONSTANT:
		return true

	case ValueClass.SYSTEM_FUNCTION:
		return node.Compare(rule)

	case ValueClass.SIGN:
		return node.Compare(rule)

	default:
		return false
	}
	return false
}

// In these case the comparison between knownNodes is removed because it can be really expensive
// but it could be added
// In other words rules with repeating tree rules wont work i.e. FC_
func compareTreeRules(node *Node, rule *Node, knownNodes *map[string]*Node) bool {
	ruleName := rule.Value
	ruleValueClass := rule.ValueClass()
	switch ruleValueClass {
	case ValueClass.NUMERIC_TREE_RULE:
		if isNumeric(node) {
			(*knownNodes)[ruleName] = node
			return true
		}

	case ValueClass.CONST_TREE_RULE:
		if isConstant(node) {
			(*knownNodes)[ruleName] = node
			return true
		}

	case ValueClass.VAR_TREE_RULE:
		if isVariable(node) {
			(*knownNodes)[ruleName] = node
			return true
		}

	default:
		return false
	}
	return false
}

func compareNegativeTreeRules(node *Node, rule *Node, knownNodes *map[string]*Node) bool {
	ruleName := rule.Value
	ruleValueClass := rule.ValueClass()
	switch ruleValueClass {
	case ValueClass.NUMERIC_TREE_RULE:
		if !isNumeric(node) {
			(*knownNodes)[ruleName] = node
			return true
		}

	case ValueClass.CONST_TREE_RULE:
		if !isConstant(node) {
			(*knownNodes)[ruleName] = node
			return true
		}

	case ValueClass.VAR_TREE_RULE:
		if !isVariable(node) {
			(*knownNodes)[ruleName] = node
			return true
		}

	default:
		return false
	}
	return false
}

func compareNonLeafs(node *Node, rule *Node, knownNodes *map[string]*Node) bool {

	ruleValueClass := rule.ValueClass()

	switch ruleValueClass {
	case ValueClass.EXPR_GENERIC_RULE:
		return compare(node.Left, rule.Left, knownNodes) && compare(node.Right, rule.Right, knownNodes)
	case ValueClass.EXPR_VARIABLE_RULE:
		if node.ValueClass() == ValueClass.VARIABLE_EXPRESSION {
			return compare(node.Left, rule.Left, knownNodes) && compare(node.Right, rule.Right, knownNodes)
		}
	case ValueClass.EXPR_CONSTANT_RULE:
		if node.ValueClass() == ValueClass.CONSTANT_EXPRESSION || node.ValueClass() == ValueClass.NUMERIC_CONSTANT_EXPRESSION {
			return compare(node.Left, rule.Left, knownNodes) && compare(node.Right, rule.Right, knownNodes)
		}
	case ValueClass.CONSTANT_EXPRESSION:
		if isConstant(node) {
			return compare(node.Left, rule.Left, knownNodes) && compare(node.Right, rule.Right, knownNodes)
		}

	case ValueClass.NUMERIC_CONSTANT_EXPRESSION:
		if node.ValueClass() == ValueClass.NUMERIC_CONSTANT_EXPRESSION {
			return compare(node.Left, rule.Left, knownNodes) && compare(node.Right, rule.Right, knownNodes)
		}

	case ValueClass.SYSTEM_FUNCTION:
		if node.ValueClass() == ValueClass.SYSTEM_FUNCTION {
			return compare(node.Left, rule.Left, knownNodes) && compare(node.Right, rule.Right, knownNodes)
		}

	case ValueClass.EXPR_NUMERIC_TREE_RULE:
		return compare(node.Left, rule.Left, knownNodes) && compare(node.Right, rule.Right, knownNodes)

	case ValueClass.EXPR_CONST_TREE_RULE:
		return compare(node.Left, rule.Left, knownNodes) && compare(node.Right, rule.Right, knownNodes)

	case ValueClass.EXPR_VAR_TREE_RULE:
		//if node.ValueClass() != ValueClass.EXPR_CONST_TREE_RULE {
		//	return compare(node.Left, rule.Left, knownNodes) && compare(node.Right, rule.Right, knownNodes)
		//}
		return compare(node.Left, rule.Left, knownNodes) && compare(node.Right, rule.Right, knownNodes)
	default:
		return false
	}
	return false
}
