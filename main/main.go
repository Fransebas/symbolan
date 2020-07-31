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

	node := symbolanProcessor.ProcessString(string(input))

	fmt.Println(node.String())

	substitute := map[string]string{
		"x": "5",
		"y": "9",
	}

	node = Tree.Substitute(node, substitute, &config)

	fmt.Println(node.String())
}
