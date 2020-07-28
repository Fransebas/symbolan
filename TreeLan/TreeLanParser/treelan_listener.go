// Code generated from TreeLan.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // TreeLan

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TreeLanListener is a complete listener for a parse tree produced by TreeLanParser.
type TreeLanListener interface {
	antlr.ParseTreeListener

	// EnterTreeRule is called when entering the treeRule production.
	EnterTreeRule(c *TreeRuleContext)

	// EnterTreeDescriptor is called when entering the treeDescriptor production.
	EnterTreeDescriptor(c *TreeDescriptorContext)

	// EnterRootNode is called when entering the rootNode production.
	EnterRootNode(c *RootNodeContext)

	// EnterLeafNode is called when entering the leafNode production.
	EnterLeafNode(c *LeafNodeContext)

	// EnterClassValue is called when entering the classValue production.
	EnterClassValue(c *ClassValueContext)

	// EnterOperationClass is called when entering the operationClass production.
	EnterOperationClass(c *OperationClassContext)

	// ExitTreeRule is called when exiting the treeRule production.
	ExitTreeRule(c *TreeRuleContext)

	// ExitTreeDescriptor is called when exiting the treeDescriptor production.
	ExitTreeDescriptor(c *TreeDescriptorContext)

	// ExitRootNode is called when exiting the rootNode production.
	ExitRootNode(c *RootNodeContext)

	// ExitLeafNode is called when exiting the leafNode production.
	ExitLeafNode(c *LeafNodeContext)

	// ExitClassValue is called when exiting the classValue production.
	ExitClassValue(c *ClassValueContext)

	// ExitOperationClass is called when exiting the operationClass production.
	ExitOperationClass(c *OperationClassContext)
}
