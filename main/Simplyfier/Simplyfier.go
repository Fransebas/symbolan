package Simplyfier

import (
	"symbolan/main/Tree"
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

func constRuleCase(node *Tree.Node) bool {
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

func (this *Simplyfier) Simplify(node Tree.Node) Tree.Node {
	if node.IsLeaf {
		return node
	}

	this.Simplify(*node.Left)
	this.Simplify(*node.Right)

	rules := *this.manager.GetRules(&node)

	for _, rule := range rules {
		matchRule := rule.node.Left
		resultRule := rule.node.Right
		if resultRule.ValueClass() == ValueClass.RULE_FUNCTION {
			return this.manager.functions[resultRule.Value](&node)
		}
		if substitutions, ok := Compare(&node, matchRule); ok {
			return resultRule.Eval(substitutions)
		}
	}

	return node
}

func isNodeKnown(name string, knownNodes *(map[string]*Tree.Node)) bool {
	_, ok := (*knownNodes)[name]
	return ok
}

func equalToExisting(name string, node *Tree.Node, knownNodes *(map[string]*Tree.Node)) bool {
	storedNode := (*knownNodes)[name]
	return storedNode.Compare(node)
}

func Compare(node *Tree.Node, rule *Tree.Node) (*map[string]*Tree.Node, bool) {
	knownNodes := make(map[string]*Tree.Node)
	if compare(node, rule, &knownNodes) {
		return &knownNodes, true
	}
	return nil, false
}

func compare(node *Tree.Node, rule *Tree.Node, knownNodes *map[string]*Tree.Node) bool {
	if rule.ValueClass() == ValueClass.EXPR_RULE {
		ruleName := rule.Value

		if !isNodeKnown(ruleName, knownNodes) {
			(*knownNodes)[ruleName] = node
			return true
		}
		return equalToExisting(ruleName, node, knownNodes)
	}

	if rule.ValueClass() == ValueClass.VARIABLE_EXPR_RULE {
		return compare(node.Left, rule.Left, knownNodes)
	}

	if !node.IsLeaf && rule.IsLeaf {
		return false
	}

	if node.IsLeaf && rule.IsLeaf {
		ruleName := rule.Value

		switch rule.ValueClass() {

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
			if node.ValueClass() == ValueClass.VARIABLE_EXPRESSION {
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

		case ValueClass.NUMERIC_RULE:
			return node.ValueClass() == ValueClass.NUMERIC_CONSTANT

		case ValueClass.IMAGINARY_CONSTANT:
			return true

		case ValueClass.SYSTEM_FUNCTION:
			return node.Compare(rule)

		default:
			return false
		}
	}

	if rule.ValueClass() == ValueClass.CONST_RULE {
		if ValueClass.IsNumeric(node.ValueClass()) || ValueClass.IsConstant(node.ValueClass()) {
			return compare(node.Left, rule.Left, knownNodes)
		}
	}

	if rule.ValueClass() == ValueClass.VAR_RULE {
		if ValueClass.IsVariable(node.ValueClass()) {
			return compare(node.Left, rule.Left, knownNodes)
		}
	}

	return false
}
