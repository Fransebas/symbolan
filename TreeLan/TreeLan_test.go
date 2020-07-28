package TreeLan

import (
	"fmt"
	"testing"
)

func Test1Simplify(t *testing.T) {

	data := "(1_left)  MULTIPLICATION_root (ANY_right) => ANY_right"

	root := ProcessString(data)

	fmt.Println(root)
}
