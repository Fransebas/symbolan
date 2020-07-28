package Tree

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {
	input, e := ioutil.ReadFile("./TestParsing")

	if e != nil {
		panic(e)
	}

	config := TreeParsingConfig{
		Simplify: NUMERIC_FUNCTIONS_2,
	}

	data := string(input)

	lines := strings.Split(data, "\n")
	for _, value := range lines {
		vals := strings.Split(value, "#")
		in := vals[0]
		expectedOut := vals[1]

		root := ProcessString(in, &config)
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

	config := TreeParsingConfig{
		Simplify: NUMERIC_FUNCTIONS_2,
	}

	data := string(input)

	lines := strings.Split(data, "\n")
	for _, value := range lines {
		vals := strings.Split(value, "#")
		in := vals[0]
		expectedOut := vals[1]

		root := ProcessString(in, &config)
		out := root.String()

		if out != expectedOut {
			t.Fail()
		}
	}
}

func Test1Simplify(t *testing.T) {
	config := TreeParsingConfig{
		Simplify: NUMERIC_FUNCTIONS_2,
	}

	data := "3sin(5)"

	in := data

	root := ProcessString(in, &config)
	out := root.String()

	if out != "-2.876773" {
		t.Fail()
	}
}
