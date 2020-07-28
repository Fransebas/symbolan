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

	config := Tree.TreeParsingConfig{
		Simplify: Tree.NUMERIC_FUNCTIONS_2,
	}

	node := Tree.ProcessString(string(input), &config)

	substitute := map[string]string{
		"x": "5",
		"y": "9",
	}

	node = Tree.Substitute(node, substitute, &config)

	fmt.Println(node.String())
}
