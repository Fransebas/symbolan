// Generated from Symbolan.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SymbolanParser}.
 */
public interface SymbolanListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SymbolanParser#eq}.
	 * @param ctx the parse tree
	 */
	void enterEq(SymbolanParser.EqContext ctx);
	/**
	 * Exit a parse tree produced by {@link SymbolanParser#eq}.
	 * @param ctx the parse tree
	 */
	void exitEq(SymbolanParser.EqContext ctx);
	/**
	 * Enter a parse tree produced by {@link SymbolanParser#system_functions}.
	 * @param ctx the parse tree
	 */
	void enterSystem_functions(SymbolanParser.System_functionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SymbolanParser#system_functions}.
	 * @param ctx the parse tree
	 */
	void exitSystem_functions(SymbolanParser.System_functionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SymbolanParser#expressionRule}.
	 * @param ctx the parse tree
	 */
	void enterExpressionRule(SymbolanParser.ExpressionRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link SymbolanParser#expressionRule}.
	 * @param ctx the parse tree
	 */
	void exitExpressionRule(SymbolanParser.ExpressionRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link SymbolanParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(SymbolanParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SymbolanParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(SymbolanParser.ExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SymbolanParser#left_par}.
	 * @param ctx the parse tree
	 */
	void enterLeft_par(SymbolanParser.Left_parContext ctx);
	/**
	 * Exit a parse tree produced by {@link SymbolanParser#left_par}.
	 * @param ctx the parse tree
	 */
	void exitLeft_par(SymbolanParser.Left_parContext ctx);
	/**
	 * Enter a parse tree produced by {@link SymbolanParser#right_par}.
	 * @param ctx the parse tree
	 */
	void enterRight_par(SymbolanParser.Right_parContext ctx);
	/**
	 * Exit a parse tree produced by {@link SymbolanParser#right_par}.
	 * @param ctx the parse tree
	 */
	void exitRight_par(SymbolanParser.Right_parContext ctx);
	/**
	 * Enter a parse tree produced by {@link SymbolanParser#atom}.
	 * @param ctx the parse tree
	 */
	void enterAtom(SymbolanParser.AtomContext ctx);
	/**
	 * Exit a parse tree produced by {@link SymbolanParser#atom}.
	 * @param ctx the parse tree
	 */
	void exitAtom(SymbolanParser.AtomContext ctx);
	/**
	 * Enter a parse tree produced by {@link SymbolanParser#number}.
	 * @param ctx the parse tree
	 */
	void enterNumber(SymbolanParser.NumberContext ctx);
	/**
	 * Exit a parse tree produced by {@link SymbolanParser#number}.
	 * @param ctx the parse tree
	 */
	void exitNumber(SymbolanParser.NumberContext ctx);
	/**
	 * Enter a parse tree produced by {@link SymbolanParser#variable}.
	 * @param ctx the parse tree
	 */
	void enterVariable(SymbolanParser.VariableContext ctx);
	/**
	 * Exit a parse tree produced by {@link SymbolanParser#variable}.
	 * @param ctx the parse tree
	 */
	void exitVariable(SymbolanParser.VariableContext ctx);
	/**
	 * Enter a parse tree produced by {@link SymbolanParser#constant}.
	 * @param ctx the parse tree
	 */
	void enterConstant(SymbolanParser.ConstantContext ctx);
	/**
	 * Exit a parse tree produced by {@link SymbolanParser#constant}.
	 * @param ctx the parse tree
	 */
	void exitConstant(SymbolanParser.ConstantContext ctx);
	/**
	 * Enter a parse tree produced by {@link SymbolanParser#bi_operators}.
	 * @param ctx the parse tree
	 */
	void enterBi_operators(SymbolanParser.Bi_operatorsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SymbolanParser#bi_operators}.
	 * @param ctx the parse tree
	 */
	void exitBi_operators(SymbolanParser.Bi_operatorsContext ctx);
}