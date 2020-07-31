// Generated from TreeLan.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link TreeLanParser}.
 */
public interface TreeLanListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link TreeLanParser#treeRule}.
	 * @param ctx the parse tree
	 */
	void enterTreeRule(TreeLanParser.TreeRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link TreeLanParser#treeRule}.
	 * @param ctx the parse tree
	 */
	void exitTreeRule(TreeLanParser.TreeRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link TreeLanParser#formula}.
	 * @param ctx the parse tree
	 */
	void enterFormula(TreeLanParser.FormulaContext ctx);
	/**
	 * Exit a parse tree produced by {@link TreeLanParser#formula}.
	 * @param ctx the parse tree
	 */
	void exitFormula(TreeLanParser.FormulaContext ctx);
	/**
	 * Enter a parse tree produced by {@link TreeLanParser#treeDescriptor}.
	 * @param ctx the parse tree
	 */
	void enterTreeDescriptor(TreeLanParser.TreeDescriptorContext ctx);
	/**
	 * Exit a parse tree produced by {@link TreeLanParser#treeDescriptor}.
	 * @param ctx the parse tree
	 */
	void exitTreeDescriptor(TreeLanParser.TreeDescriptorContext ctx);
	/**
	 * Enter a parse tree produced by {@link TreeLanParser#rootNode}.
	 * @param ctx the parse tree
	 */
	void enterRootNode(TreeLanParser.RootNodeContext ctx);
	/**
	 * Exit a parse tree produced by {@link TreeLanParser#rootNode}.
	 * @param ctx the parse tree
	 */
	void exitRootNode(TreeLanParser.RootNodeContext ctx);
	/**
	 * Enter a parse tree produced by {@link TreeLanParser#leafNode}.
	 * @param ctx the parse tree
	 */
	void enterLeafNode(TreeLanParser.LeafNodeContext ctx);
	/**
	 * Exit a parse tree produced by {@link TreeLanParser#leafNode}.
	 * @param ctx the parse tree
	 */
	void exitLeafNode(TreeLanParser.LeafNodeContext ctx);
	/**
	 * Enter a parse tree produced by {@link TreeLanParser#classValue}.
	 * @param ctx the parse tree
	 */
	void enterClassValue(TreeLanParser.ClassValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link TreeLanParser#classValue}.
	 * @param ctx the parse tree
	 */
	void exitClassValue(TreeLanParser.ClassValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link TreeLanParser#operationClass}.
	 * @param ctx the parse tree
	 */
	void enterOperationClass(TreeLanParser.OperationClassContext ctx);
	/**
	 * Exit a parse tree produced by {@link TreeLanParser#operationClass}.
	 * @param ctx the parse tree
	 */
	void exitOperationClass(TreeLanParser.OperationClassContext ctx);
}