package Node

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {
	input, e := ioutil.ReadFile("./TestParsing")

	if e != nil {
		panic(e)
	}

	config := SymbolanProcessorConfig{
		Simplify: NUMERIC_FUNCTIONS_2,
	}

	symbolanProcessor := NewSymbolanProcessor(&config)
	data := string(input)

	lines := strings.Split(data, "\n")
	for _, value := range lines {
		vals := strings.Split(value, "#")
		in := vals[0]
		expectedOut := vals[1]

		root := symbolanProcessor.ProcessString(in)
		out := root.String()

		if out != expectedOut {
			t.Fail()
		}
	}
}

func TestSimplify(t *testing.T) {
	input, e := ioutil.ReadFile("./TestSimplify")

	if e != nil {
		panic(e)
	}

	config := SymbolanProcessorConfig{
		Simplify: NUMERIC_BASIC_OPS_1,
	}
	symbolanProcessor := NewSymbolanProcessor(&config)

	data := string(input)

	lines := strings.Split(data, "\n")
	for _, value := range lines {
		vals := strings.Split(value, "#")
		in := vals[0]
		expectedOut := vals[1]

		root := symbolanProcessor.ProcessString(in)
		out := root.String()

		fmt.Println(out)

		if out != expectedOut {
			t.Fail()
		}
	}
}

func Test1Simplify(t *testing.T) {
	config := SymbolanProcessorConfig{
		Simplify: NUMERIC_FUNCTIONS_2,
	}
	symbolanProcessor := NewSymbolanProcessor(&config)

	data := "3sin(5)"

	in := data

	root := symbolanProcessor.ProcessString(in)
	out := root.String()

	if out != "-2.876773" {
		t.Fail()
	}
}
