package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"symbolan/main/Node"
)

func main() {

	input, e := ioutil.ReadFile(os.Args[1])

	if e != nil {
		panic(e)
	}

	config := Node.SymbolanProcessorConfig{
		Simplify: Node.NUMERIC_FUNCTIONS_2,
	}

	symbolanProcessor := Node.NewSymbolanProcessor(&config)

	node := symbolanProcessor.ProcessString(string(input))

	fmt.Println(node.String())
}
