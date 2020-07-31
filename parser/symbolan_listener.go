// Code generated from Symbolan.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Symbolan

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SymbolanListener is a complete listener for a parse tree produced by SymbolanParser.
type SymbolanListener interface {
	antlr.ParseTreeListener

	// EnterEq is called when entering the eq production.
	EnterEq(c *EqContext)

	// EnterSystem_functions is called when entering the system_functions production.
	EnterSystem_functions(c *System_functionsContext)

	// EnterExpressionRule is called when entering the expressionRule production.
	EnterExpressionRule(c *ExpressionRuleContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterLeft_par is called when entering the left_par production.
	EnterLeft_par(c *Left_parContext)

	// EnterRight_par is called when entering the right_par production.
	EnterRight_par(c *Right_parContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterBi_operators is called when entering the bi_operators production.
	EnterBi_operators(c *Bi_operatorsContext)

	// ExitEq is called when exiting the eq production.
	ExitEq(c *EqContext)

	// ExitSystem_functions is called when exiting the system_functions production.
	ExitSystem_functions(c *System_functionsContext)

	// ExitExpressionRule is called when exiting the expressionRule production.
	ExitExpressionRule(c *ExpressionRuleContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitLeft_par is called when exiting the left_par production.
	ExitLeft_par(c *Left_parContext)

	// ExitRight_par is called when exiting the right_par production.
	ExitRight_par(c *Right_parContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitBi_operators is called when exiting the bi_operators production.
	ExitBi_operators(c *Bi_operatorsContext)
}
