// Code generated from Symbolan.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Symbolan

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSymbolanVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSymbolanVisitor) VisitEq(ctx *EqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSymbolanVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSymbolanVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSymbolanVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSymbolanVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSymbolanVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSymbolanVisitor) VisitBi_operators(ctx *Bi_operatorsContext) interface{} {
	return v.VisitChildren(ctx)
}
