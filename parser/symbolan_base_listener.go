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

// EnterEq is called when production eq is entered.
func (s *BaseSymbolanListener) EnterEq(ctx *EqContext) {}

// ExitEq is called when production eq is exited.
func (s *BaseSymbolanListener) ExitEq(ctx *EqContext) {}

// EnterSystem_functions is called when production system_functions is entered.
func (s *BaseSymbolanListener) EnterSystem_functions(ctx *System_functionsContext) {}

// ExitSystem_functions is called when production system_functions is exited.
func (s *BaseSymbolanListener) ExitSystem_functions(ctx *System_functionsContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSymbolanListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSymbolanListener) ExitExpr(ctx *ExprContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseSymbolanListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseSymbolanListener) ExitFunction(ctx *FunctionContext) {}

// EnterLeft_par is called when production left_par is entered.
func (s *BaseSymbolanListener) EnterLeft_par(ctx *Left_parContext) {}

// ExitLeft_par is called when production left_par is exited.
func (s *BaseSymbolanListener) ExitLeft_par(ctx *Left_parContext) {}

// EnterRight_par is called when production right_par is entered.
func (s *BaseSymbolanListener) EnterRight_par(ctx *Right_parContext) {}

// ExitRight_par is called when production right_par is exited.
func (s *BaseSymbolanListener) ExitRight_par(ctx *Right_parContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseSymbolanListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseSymbolanListener) ExitAtom(ctx *AtomContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseSymbolanListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseSymbolanListener) ExitNumber(ctx *NumberContext) {}

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
