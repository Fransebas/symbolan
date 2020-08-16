package Node

import (
	"fmt"
	"symbolan/OperationClass"
)

func (this *Node) Latex() string {
	beginLatex := ``
	endLatex := ``

	return beginLatex + this.latex() + endLatex
}

func (this *Node) latex() string {
	if this.IsLeaf {
		return this.Value
	}
	if this.classByOperation == OperationClass.SYSTEM_FUNCTION {
		return this.Left.latex() + "(" + this.Right.latex() + ")"
	} else {
		if this.OperationClass() == OperationClass.ADDITION {
			return this.Left.latex() + this.Operation + this.Right.latex()
		} else if this.OperationClass() == OperationClass.DIVISION {
			return fmt.Sprintf(`\frac{%v}{%v}`, this.Left.latex(), this.Right.latex())
		} else if this.OperationClass() == OperationClass.EXPONENTIAL {
			if this.Left.treeSize > 1 {
				return fmt.Sprintf(`{(%v)}^{%v}`, this.Left.latex(), this.Right.latex())
			}
			return fmt.Sprintf(`{%v}^{%v}`, this.Left.latex(), this.Right.latex())
		} else if this.OperationClass() == OperationClass.MULTIPLICATION {
			operation := ""
			left := ""
			if this.Left.OperationClass() == OperationClass.MULTIPLICATION || this.Left.OperationClass() == OperationClass.SYSTEM_FUNCTION || this.Left.IsLeaf {
				left = fmt.Sprintf("%v", this.Left.latex())
			} else {
				left = fmt.Sprintf("(%v)", this.Left.latex())
				operation = "*"
			}

			right := ""
			if this.Right.OperationClass() == OperationClass.MULTIPLICATION || this.Right.OperationClass() == OperationClass.SYSTEM_FUNCTION || this.Right.IsLeaf {
				right = fmt.Sprintf("%v", this.Right.latex())
			} else {
				right = fmt.Sprintf("(%v)", this.Right.latex())
				operation = "*"
			}
			return left + operation + right

		} else if this.Right != nil {
			operation := this.Operation
			if this.OperationClass() == OperationClass.SIGN {
				operation = ""
			}
			left := ""
			if this.Left.treeSize > 1 {
				left = fmt.Sprintf("(%v)", this.Left.latex())
			} else {
				left = fmt.Sprintf("%v", this.Left.latex())
			}
			right := ""
			if this.Right.treeSize > 1 {
				right = fmt.Sprintf("(%v)", this.Right.latex())
			} else {
				right = fmt.Sprintf("%v", this.Right.latex())
			}
			return left + operation + right
		} else {
			return this.Left.latex() + this.Operation
		}
	}

}
