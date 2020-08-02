// Code generated from Symbolan.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Symbolan

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SymbolanListener is a complete listener for a parse tree produced by SymbolanParser.
type SymbolanListener interface {
	antlr.ParseTreeListener

	// EnterRuleSet is called when entering the ruleSet production.
	EnterRuleSet(c *RuleSetContext)

	// EnterEq is called when entering the eq production.
	EnterEq(c *EqContext)

	// EnterSystem_functions is called when entering the system_functions production.
	EnterSystem_functions(c *System_functionsContext)

	// EnterRule_function is called when entering the rule_function production.
	EnterRule_function(c *Rule_functionContext)

	// EnterExpressionRule is called when entering the expressionRule production.
	EnterExpressionRule(c *ExpressionRuleContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterRuleAtom is called when entering the ruleAtom production.
	EnterRuleAtom(c *RuleAtomContext)

	// EnterNumeric_rule is called when entering the numeric_rule production.
	EnterNumeric_rule(c *Numeric_ruleContext)

	// EnterConst_rule is called when entering the const_rule production.
	EnterConst_rule(c *Const_ruleContext)

	// EnterVar_rule is called when entering the var_rule production.
	EnterVar_rule(c *Var_ruleContext)

	// EnterExpr_rule is called when entering the expr_rule production.
	EnterExpr_rule(c *Expr_ruleContext)

	// EnterLeft_par is called when entering the left_par production.
	EnterLeft_par(c *Left_parContext)

	// EnterRight_par is called when entering the right_par production.
	EnterRight_par(c *Right_parContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterMath_constant is called when entering the math_constant production.
	EnterMath_constant(c *Math_constantContext)

	// EnterInfinite is called when entering the infinite production.
	EnterInfinite(c *InfiniteContext)

	// EnterImaginary is called when entering the imaginary production.
	EnterImaginary(c *ImaginaryContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterBi_operators is called when entering the bi_operators production.
	EnterBi_operators(c *Bi_operatorsContext)

	// ExitRuleSet is called when exiting the ruleSet production.
	ExitRuleSet(c *RuleSetContext)

	// ExitEq is called when exiting the eq production.
	ExitEq(c *EqContext)

	// ExitSystem_functions is called when exiting the system_functions production.
	ExitSystem_functions(c *System_functionsContext)

	// ExitRule_function is called when exiting the rule_function production.
	ExitRule_function(c *Rule_functionContext)

	// ExitExpressionRule is called when exiting the expressionRule production.
	ExitExpressionRule(c *ExpressionRuleContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitRuleAtom is called when exiting the ruleAtom production.
	ExitRuleAtom(c *RuleAtomContext)

	// ExitNumeric_rule is called when exiting the numeric_rule production.
	ExitNumeric_rule(c *Numeric_ruleContext)

	// ExitConst_rule is called when exiting the const_rule production.
	ExitConst_rule(c *Const_ruleContext)

	// ExitVar_rule is called when exiting the var_rule production.
	ExitVar_rule(c *Var_ruleContext)

	// ExitExpr_rule is called when exiting the expr_rule production.
	ExitExpr_rule(c *Expr_ruleContext)

	// ExitLeft_par is called when exiting the left_par production.
	ExitLeft_par(c *Left_parContext)

	// ExitRight_par is called when exiting the right_par production.
	ExitRight_par(c *Right_parContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitMath_constant is called when exiting the math_constant production.
	ExitMath_constant(c *Math_constantContext)

	// ExitInfinite is called when exiting the infinite production.
	ExitInfinite(c *InfiniteContext)

	// ExitImaginary is called when exiting the imaginary production.
	ExitImaginary(c *ImaginaryContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitBi_operators is called when exiting the bi_operators production.
	ExitBi_operators(c *Bi_operatorsContext)
}
