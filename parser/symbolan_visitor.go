// Code generated from Symbolan.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Symbolan

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SymbolanParser.
type SymbolanVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SymbolanParser#eq.
	VisitEq(ctx *EqContext) interface{}

	// Visit a parse tree produced by SymbolanParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SymbolanParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by SymbolanParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by SymbolanParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by SymbolanParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by SymbolanParser#bi_operators.
	VisitBi_operators(ctx *Bi_operatorsContext) interface{}
}
