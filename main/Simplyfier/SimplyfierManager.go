package Simplyfier

import (
	"fmt"
	"io/ioutil"
	"symbolan/main/OperationClass"
	"symbolan/main/Tree"
	"symbolan/main/ValueClass"
)

const rulesFile = ""

type Rule struct {
	node *Tree.Node
}

func NewRule(node *Tree.Node) Rule {
	return Rule{
		node: node,
	}
}

type Manager struct {
	index     map[OperationClass.OperationClass]*[]Rule
	functions map[string]func(node *Tree.Node) Tree.Node
}

func treeAgmentedValuesToString(treeSize, treeHeight int64) string {
	return fmt.Sprintf("(%v,%v)", treeSize, treeHeight)
}

func (this *Manager) GetRules(node *Tree.Node) *[]Rule {
	return this.getRules(node.OperationClass(), node.ValueClass(), node.Height(), node.TreeSize())
}

func (this *Manager) containsKeys(opClass OperationClass.OperationClass, vClass ValueClass.ValueClass, treeSize, treeHeight int64) bool {
	firstKey := treeAgmentedValuesToString(treeSize, treeHeight)
	if subMap, ok := this.index[firstKey]; ok {
		if _, ok := subMap[opClass]; ok {
			return true
		}
	}
	return false
}

func (this *Manager) containsNode(node *Tree.Node) bool {
	return this.containsKeys(node.OperationClass(), node.ValueClass(), node.Height(), node.TreeSize())
}

func (this *Manager) getRules(opClass OperationClass.OperationClass, vClass ValueClass.ValueClass, treeSize, treeHeight int64) *[]Rule {
	firstKey := treeAgmentedValuesToString(treeSize, treeHeight)
	return this.index[firstKey][opClass]
}

func (this *Manager) addRule(ruleNode *Tree.Node) {
	opClass := ruleNode.OperationClass()
	firstKey := treeAgmentedValuesToString(ruleNode.TreeSize(), ruleNode.Height())

	if _, ok := this.index[firstKey]; !ok {
		this.index[firstKey] = make(map[OperationClass.OperationClass]*[]Rule)
	}

	subMap := this.index[firstKey]
	var rules []Rule
	if _, ok := subMap[opClass]; !ok {
		rules = make([]Rule, 0)
		subMap[opClass] = &rules
	}
	rules = *subMap[opClass]
	rules = append(rules, NewRule(ruleNode))
}

func NewSimplyfierManager() *Manager {
	manager := new(Manager)
	manager.index = make(map[string]map[OperationClass.OperationClass]*[]Rule, 0)

	manager.readRules()
	manager.functions = map[string]func(node *Tree.Node) Tree.Node{
		"addition":       addition,
		"subtraction":    subtraction,
		"multiplication": multiplication,
		"division":       division,
		"exponential":    exponential,
	}

	return manager
}

func processRules(data string) *[]Tree.Node {
	config := Tree.SymbolanProcessorConfig{
		Simplify: Tree.NONE,
	}
	symbolanProcessor := Tree.NewSymbolanProcessor(&config)

	ruleSet := symbolanProcessor.ProcessRuleString(data)

	return &ruleSet.Rules
}

func (this *Manager) readRules() {
	byteData, e := ioutil.ReadFile(rulesFile)
	if e != nil {
		panic(e)
	}
	data := string(byteData)
	ruleNodes := processRules(data)

	for _, ruleNode := range *ruleNodes {
		this.addRule(&ruleNode)
	}
}
