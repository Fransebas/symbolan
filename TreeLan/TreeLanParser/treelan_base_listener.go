// Code generated from TreeLan.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // TreeLan

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTreeLanListener is a complete listener for a parse tree produced by TreeLanParser.
type BaseTreeLanListener struct{}

var _ TreeLanListener = &BaseTreeLanListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTreeLanListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTreeLanListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTreeLanListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTreeLanListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTreeRule is called when production treeRule is entered.
func (s *BaseTreeLanListener) EnterTreeRule(ctx *TreeRuleContext) {}

// ExitTreeRule is called when production treeRule is exited.
func (s *BaseTreeLanListener) ExitTreeRule(ctx *TreeRuleContext) {}

// EnterTreeDescriptor is called when production treeDescriptor is entered.
func (s *BaseTreeLanListener) EnterTreeDescriptor(ctx *TreeDescriptorContext) {}

// ExitTreeDescriptor is called when production treeDescriptor is exited.
func (s *BaseTreeLanListener) ExitTreeDescriptor(ctx *TreeDescriptorContext) {}

// EnterRootNode is called when production rootNode is entered.
func (s *BaseTreeLanListener) EnterRootNode(ctx *RootNodeContext) {}

// ExitRootNode is called when production rootNode is exited.
func (s *BaseTreeLanListener) ExitRootNode(ctx *RootNodeContext) {}

// EnterLeafNode is called when production leafNode is entered.
func (s *BaseTreeLanListener) EnterLeafNode(ctx *LeafNodeContext) {}

// ExitLeafNode is called when production leafNode is exited.
func (s *BaseTreeLanListener) ExitLeafNode(ctx *LeafNodeContext) {}

// EnterClassValue is called when production classValue is entered.
func (s *BaseTreeLanListener) EnterClassValue(ctx *ClassValueContext) {}

// ExitClassValue is called when production classValue is exited.
func (s *BaseTreeLanListener) ExitClassValue(ctx *ClassValueContext) {}

// EnterOperationClass is called when production operationClass is entered.
func (s *BaseTreeLanListener) EnterOperationClass(ctx *OperationClassContext) {}

// ExitOperationClass is called when production operationClass is exited.
func (s *BaseTreeLanListener) ExitOperationClass(ctx *OperationClassContext) {}
