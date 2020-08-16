package Node

import (
	"fmt"
	"io/ioutil"
	"symbolan/OperationClass"
	"symbolan/ValueClass"
)

const rulesFile = "./Node/rules.sp"

type Rule struct {
	node *Node
}

func NewRule(node *Node) Rule {
	return Rule{
		node: node,
	}
}

type Manager struct {
	index     map[OperationClass.OperationClass][]Rule
	functions map[string]func(node *Node) Node
}

func treeAgmentedValuesToString(treeSize, treeHeight int64) string {
	return fmt.Sprintf("(%v,%v)", treeSize, treeHeight)
}

func (this *Manager) GetRules(node *Node) []Rule {
	return this.getRules(node.OperationClass(), node.ValueClass(), node.Height(), node.TreeSize())
}

func (this *Manager) containsKeys(opClass OperationClass.OperationClass, vClass ValueClass.ValueClass, treeSize, treeHeight int64) bool {
	_, ok := this.index[opClass]
	return ok
}

func (this *Manager) containsNode(node *Node) bool {
	return this.containsKeys(node.OperationClass(), node.ValueClass(), node.Height(), node.TreeSize())
}

func (this *Manager) getRules(opClass OperationClass.OperationClass, vClass ValueClass.ValueClass, treeSize, treeHeight int64) []Rule {
	return this.index[opClass]
}

func (this *Manager) addRule(ruleNode Node) {
	opClass := ruleNode.Left.OperationClass()

	//if _, ok := this.index[firstKey]; !ok {
	//	this.index[firstKey] = make(map[OperationClass.OperationClass]*[]Rule)
	//}

	var rules []Rule
	rulesAddr, ok := this.index[opClass]
	if !ok {
		rules = make([]Rule, 0)
		this.index[opClass] = rules
	} else {
		rules = rulesAddr
	}
	rules = append(rules, NewRule(&ruleNode))
	this.index[opClass] = rules
}

func NewSimplyfierManager() *Manager {
	manager := new(Manager)
	manager.index = make(map[OperationClass.OperationClass][]Rule, 0)

	manager.readRules()
	manager.functions = map[string]func(node *Node) Node{
		"Addition":          addition,
		"Subtraction":       subtraction,
		"Multiplication":    multiplication,
		"Division":          division,
		"Exponential":       exponential,
		"VarDerivative":     VarDerivative,
		"SignedExponential": signedExponential,
	}

	return manager
}

func processRules(data string) *[]Node {
	config := SymbolanProcessorConfig{
		Simplify: NONE,
	}
	symbolanProcessor := NewSymbolanProcessor(&config)

	ruleSet := symbolanProcessor.ProcessRuleString(data)

	return &ruleSet.Rules
}

func (this *Manager) readRules() {
	//path, _ := os.Getwd()
	byteData, e := ioutil.ReadFile(rulesFile)
	if e != nil {
		panic(e)
	}
	data := string(byteData)
	ruleNodes := processRules(data)

	for _, ruleNode := range *ruleNodes {
		this.addRule(ruleNode)
	}
}
