// Code generated from Symbolan.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Symbolan

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSymbolanListener is a complete listener for a parse tree produced by SymbolanParser.
type BaseSymbolanListener struct{}

var _ SymbolanListener = &BaseSymbolanListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSymbolanListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSymbolanListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSymbolanListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSymbolanListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRuleSet is called when production ruleSet is entered.
func (s *BaseSymbolanListener) EnterRuleSet(ctx *RuleSetContext) {}

// ExitRuleSet is called when production ruleSet is exited.
func (s *BaseSymbolanListener) ExitRuleSet(ctx *RuleSetContext) {}

// EnterEq is called when production eq is entered.
func (s *BaseSymbolanListener) EnterEq(ctx *EqContext) {}

// ExitEq is called when production eq is exited.
func (s *BaseSymbolanListener) ExitEq(ctx *EqContext) {}

// EnterSystem_functions is called when production system_functions is entered.
func (s *BaseSymbolanListener) EnterSystem_functions(ctx *System_functionsContext) {}

// ExitSystem_functions is called when production system_functions is exited.
func (s *BaseSymbolanListener) ExitSystem_functions(ctx *System_functionsContext) {}

// EnterRule_function is called when production rule_function is entered.
func (s *BaseSymbolanListener) EnterRule_function(ctx *Rule_functionContext) {}

// ExitRule_function is called when production rule_function is exited.
func (s *BaseSymbolanListener) ExitRule_function(ctx *Rule_functionContext) {}

// EnterExpressionRule is called when production expressionRule is entered.
func (s *BaseSymbolanListener) EnterExpressionRule(ctx *ExpressionRuleContext) {}

// ExitExpressionRule is called when production expressionRule is exited.
func (s *BaseSymbolanListener) ExitExpressionRule(ctx *ExpressionRuleContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSymbolanListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSymbolanListener) ExitExpr(ctx *ExprContext) {}

// EnterRuleAtom is called when production ruleAtom is entered.
func (s *BaseSymbolanListener) EnterRuleAtom(ctx *RuleAtomContext) {}

// ExitRuleAtom is called when production ruleAtom is exited.
func (s *BaseSymbolanListener) ExitRuleAtom(ctx *RuleAtomContext) {}

// EnterSign is called when production sign is entered.
func (s *BaseSymbolanListener) EnterSign(ctx *SignContext) {}

// ExitSign is called when production sign is exited.
func (s *BaseSymbolanListener) ExitSign(ctx *SignContext) {}

// EnterLeft_par is called when production left_par is entered.
func (s *BaseSymbolanListener) EnterLeft_par(ctx *Left_parContext) {}

// ExitLeft_par is called when production left_par is exited.
func (s *BaseSymbolanListener) ExitLeft_par(ctx *Left_parContext) {}

// EnterRight_par is called when production right_par is entered.
func (s *BaseSymbolanListener) EnterRight_par(ctx *Right_parContext) {}

// ExitRight_par is called when production right_par is exited.
func (s *BaseSymbolanListener) ExitRight_par(ctx *Right_parContext) {}

// EnterDerivative_rule is called when production derivative_rule is entered.
func (s *BaseSymbolanListener) EnterDerivative_rule(ctx *Derivative_ruleContext) {}

// ExitDerivative_rule is called when production derivative_rule is exited.
func (s *BaseSymbolanListener) ExitDerivative_rule(ctx *Derivative_ruleContext) {}

// EnterNumeric_rule is called when production numeric_rule is entered.
func (s *BaseSymbolanListener) EnterNumeric_rule(ctx *Numeric_ruleContext) {}

// ExitNumeric_rule is called when production numeric_rule is exited.
func (s *BaseSymbolanListener) ExitNumeric_rule(ctx *Numeric_ruleContext) {}

// EnterConst_rule is called when production const_rule is entered.
func (s *BaseSymbolanListener) EnterConst_rule(ctx *Const_ruleContext) {}

// ExitConst_rule is called when production const_rule is exited.
func (s *BaseSymbolanListener) ExitConst_rule(ctx *Const_ruleContext) {}

// EnterVar_rule is called when production var_rule is entered.
func (s *BaseSymbolanListener) EnterVar_rule(ctx *Var_ruleContext) {}

// ExitVar_rule is called when production var_rule is exited.
func (s *BaseSymbolanListener) ExitVar_rule(ctx *Var_ruleContext) {}

// EnterExpr_rule is called when production expr_rule is entered.
func (s *BaseSymbolanListener) EnterExpr_rule(ctx *Expr_ruleContext) {}

// ExitExpr_rule is called when production expr_rule is exited.
func (s *BaseSymbolanListener) ExitExpr_rule(ctx *Expr_ruleContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseSymbolanListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseSymbolanListener) ExitAtom(ctx *AtomContext) {}

// EnterMath_constant is called when production math_constant is entered.
func (s *BaseSymbolanListener) EnterMath_constant(ctx *Math_constantContext) {}

// ExitMath_constant is called when production math_constant is exited.
func (s *BaseSymbolanListener) ExitMath_constant(ctx *Math_constantContext) {}

// EnterInfinite is called when production infinite is entered.
func (s *BaseSymbolanListener) EnterInfinite(ctx *InfiniteContext) {}

// ExitInfinite is called when production infinite is exited.
func (s *BaseSymbolanListener) ExitInfinite(ctx *InfiniteContext) {}

// EnterImaginary is called when production imaginary is entered.
func (s *BaseSymbolanListener) EnterImaginary(ctx *ImaginaryContext) {}

// ExitImaginary is called when production imaginary is exited.
func (s *BaseSymbolanListener) ExitImaginary(ctx *ImaginaryContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseSymbolanListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseSymbolanListener) ExitNumber(ctx *NumberContext) {}

// EnterDerivative is called when production derivative is entered.
func (s *BaseSymbolanListener) EnterDerivative(ctx *DerivativeContext) {}

// ExitDerivative is called when production derivative is exited.
func (s *BaseSymbolanListener) ExitDerivative(ctx *DerivativeContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseSymbolanListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseSymbolanListener) ExitVariable(ctx *VariableContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseSymbolanListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseSymbolanListener) ExitConstant(ctx *ConstantContext) {}

// EnterBi_operators is called when production bi_operators is entered.
func (s *BaseSymbolanListener) EnterBi_operators(ctx *Bi_operatorsContext) {}

// ExitBi_operators is called when production bi_operators is exited.
func (s *BaseSymbolanListener) ExitBi_operators(ctx *Bi_operatorsContext) {}
