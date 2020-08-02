package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"symbolan/main/Tree"
)

func main() {

	input, e := ioutil.ReadFile(os.Args[1])

	if e != nil {
		panic(e)
	}

	config := Tree.SymbolanProcessorConfig{
		Simplify: Tree.NUMERIC_FUNCTIONS_2,
	}

	symbolanProcessor := Tree.NewSymbolanProcessor(&config)

	node := symbolanProcessor.ProcessRuleString(string(input))

}
