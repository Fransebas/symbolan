package Tree

import "github.com/antlr/antlr4/runtime/Go/antlr"

type RuleSet struct {
	Rules []Node
}

func NewRuleSet(ctx antlr.Tree) *RuleSet {

	ruleSet := new(RuleSet)
	ruleSet.Rules = make([]Node, 0)

	childrenCount := ctx.GetChildCount()
	for i := 0; i < childrenCount; i++ {
		childrenCntx := ctx.GetChild(i)
		if childrenCntx.GetChildCount() == 1 {
			continue
		}
		ruleSet.Rules = append(ruleSet.Rules, *ExprRule(childrenCntx))
	}

	return ruleSet
}
