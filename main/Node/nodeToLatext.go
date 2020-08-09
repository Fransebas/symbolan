package Node

import (
	"fmt"
	"symbolan/main/OperationClass"
)

func (this *Node) Latex() string {
	beginLatex := `\documentclass[12pt]{article} \begin{document} \begin{equation}` + "\n"
	endLatex := "\n" + `\end{equation}  \end{document}`

	return beginLatex + this.latex() + endLatex
}

func (this *Node) latex() string {
	if this.IsLeaf {
		return this.Value
	}
	if this.classByOperation == OperationClass.SYSTEM_FUNCTION {
		return this.Left.String() + "(" + this.Right.String() + ")"
	} else {
		if this.OperationClass() == OperationClass.ADDITION {
			return this.Left.String() + this.Operation + this.Right.String()
		} else if this.OperationClass() == OperationClass.DIVISION {
			return fmt.Sprintf(`\frac{%v}{%v}`, this.Left.String(), this.Right.String())
		} else if this.Right != nil {
			operation := this.Operation
			if this.OperationClass() == OperationClass.SIGN {
				operation = ""
			}
			left := ""
			if this.Left.treeSize > 1 {
				left = fmt.Sprintf("(%v)", this.Left.String())
			} else {
				left = fmt.Sprintf("%v", this.Left.String())
			}
			right := ""
			if this.Right.treeSize > 1 {
				right = fmt.Sprintf("(%v)", this.Right.String())
			} else {
				right = fmt.Sprintf("%v", this.Right.String())
			}
			return left + operation + right
		} else {
			return this.Left.String() + this.Operation
		}
	}

}
