// Generated from Symbolan.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SymbolanParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, L_PAREN=9, 
		R_PAREN=10, L_CURLY=11, R_CURLY=12, L_BRACKET=13, R_BRACKET=14, ASSIGN=15, 
		COMMA=16, FACTORIAL=17, SIN=18, COS=19, TAN=20, COT=21, CSC=22, SEC=23, 
		SINH=24, COSH=25, TANH=26, COTH=27, CSCH=28, SECH=29, MIN=30, MAX=31, 
		LOG=32, SQRT=33, DERIVATIVE_OPERATOR=34, PI=35, DERIVATIVE=36, VAR=37, 
		CONST=38, EULER=39, INFINITE=40, UNDEFINED=41, DERIVATE_RULE=42, NUMERIC_RULE=43, 
		CONST_RULE=44, VAR_RULE=45, EXPR_RULE=46, NUMERIC_TREE_RULE=47, CONST_TREE_RULE=48, 
		VAR_TREE_RULE=49, SYSTEM_TREE_RULE=50, IMAGINARY=51, DECIMAL_LIT=52, OCTAL_LIT=53, 
		HEX_LIT=54, FLOAT_LIT=55, IMAGINARY_LIT=56, IDENTIFIER=57, NEWLINE=58, 
		WS=59, LINE_COMMENT=60;
	public static final int
		RULE_ruleSet = 0, RULE_eq = 1, RULE_system_functions = 2, RULE_rule_function = 3, 
		RULE_expressionRule = 4, RULE_expr = 5, RULE_ruleAtom = 6, RULE_sign = 7, 
		RULE_left_par = 8, RULE_right_par = 9, RULE_derivative_rule = 10, RULE_numeric_rule = 11, 
		RULE_const_rule = 12, RULE_var_rule = 13, RULE_expr_rule = 14, RULE_numeric_tree_rule = 15, 
		RULE_const_tree_rule = 16, RULE_var_tree_rule = 17, RULE_system_tree_rule = 18, 
		RULE_atom = 19, RULE_math_constant = 20, RULE_infinite = 21, RULE_imaginary = 22, 
		RULE_number = 23, RULE_derivative = 24, RULE_variable = 25, RULE_constant = 26, 
		RULE_bi_operators = 27;
	public static final String[] ruleNames = {
		"ruleSet", "eq", "system_functions", "rule_function", "expressionRule", 
		"expr", "ruleAtom", "sign", "left_par", "right_par", "derivative_rule", 
		"numeric_rule", "const_rule", "var_rule", "expr_rule", "numeric_tree_rule", 
		"const_tree_rule", "var_tree_rule", "system_tree_rule", "atom", "math_constant", 
		"infinite", "imaginary", "number", "derivative", "variable", "constant", 
		"bi_operators"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'()'", "'=>'", "'^'", "'**'", "'*'", "'/'", "'+'", "'-'", "'('", 
		"')'", "'{'", "'}'", "'['", "']'", "'='", "','", "'!'", "'sin'", "'cos'", 
		"'tan'", "'cot'", "'csc'", "'sec'", "'sinh'", "'cosh'", "'tanh'", "'coth'", 
		"'csch'", "'sech'", "'min'", "'max'", "'log'", "'sqrt'", "'D'", "'pi'", 
		null, null, null, null, null, "'undef'", null, null, null, null, null, 
		null, null, null, null, "'i'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, null, null, "L_PAREN", "R_PAREN", 
		"L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "FACTORIAL", 
		"SIN", "COS", "TAN", "COT", "CSC", "SEC", "SINH", "COSH", "TANH", "COTH", 
		"CSCH", "SECH", "MIN", "MAX", "LOG", "SQRT", "DERIVATIVE_OPERATOR", "PI", 
		"DERIVATIVE", "VAR", "CONST", "EULER", "INFINITE", "UNDEFINED", "DERIVATE_RULE", 
		"NUMERIC_RULE", "CONST_RULE", "VAR_RULE", "EXPR_RULE", "NUMERIC_TREE_RULE", 
		"CONST_TREE_RULE", "VAR_TREE_RULE", "SYSTEM_TREE_RULE", "IMAGINARY", "DECIMAL_LIT", 
		"OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "IMAGINARY_LIT", "IDENTIFIER", "NEWLINE", 
		"WS", "LINE_COMMENT"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Symbolan.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SymbolanParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class RuleSetContext extends ParserRuleContext {
		public List<ExpressionRuleContext> expressionRule() {
			return getRuleContexts(ExpressionRuleContext.class);
		}
		public ExpressionRuleContext expressionRule(int i) {
			return getRuleContext(ExpressionRuleContext.class,i);
		}
		public List<TerminalNode> NEWLINE() { return getTokens(SymbolanParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(SymbolanParser.NEWLINE, i);
		}
		public RuleSetContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ruleSet; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterRuleSet(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitRuleSet(this);
		}
	}

	public final RuleSetContext ruleSet() throws RecognitionException {
		RuleSetContext _localctx = new RuleSetContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_ruleSet);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(65);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__6) | (1L << T__7) | (1L << L_PAREN) | (1L << SIN) | (1L << COS) | (1L << TAN) | (1L << SINH) | (1L << COSH) | (1L << TANH) | (1L << COTH) | (1L << CSCH) | (1L << SECH) | (1L << MIN) | (1L << MAX) | (1L << LOG) | (1L << DERIVATIVE_OPERATOR) | (1L << PI) | (1L << DERIVATIVE) | (1L << VAR) | (1L << CONST) | (1L << EULER) | (1L << INFINITE) | (1L << UNDEFINED) | (1L << DERIVATE_RULE) | (1L << NUMERIC_RULE) | (1L << CONST_RULE) | (1L << VAR_RULE) | (1L << EXPR_RULE) | (1L << NUMERIC_TREE_RULE) | (1L << CONST_TREE_RULE) | (1L << VAR_TREE_RULE) | (1L << SYSTEM_TREE_RULE) | (1L << IMAGINARY) | (1L << DECIMAL_LIT) | (1L << OCTAL_LIT) | (1L << HEX_LIT) | (1L << FLOAT_LIT) | (1L << IMAGINARY_LIT))) != 0)) {
				{
				{
				setState(56);
				expressionRule();
				setState(60);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==NEWLINE) {
					{
					{
					setState(57);
					match(NEWLINE);
					}
					}
					setState(62);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				}
				setState(67);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EqContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public EqContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_eq; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterEq(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitEq(this);
		}
	}

	public final EqContext eq() throws RecognitionException {
		EqContext _localctx = new EqContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_eq);
		try {
			setState(73);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(68);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(69);
				expr(0);
				setState(70);
				match(ASSIGN);
				setState(71);
				expr(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class System_functionsContext extends ParserRuleContext {
		public TerminalNode MIN() { return getToken(SymbolanParser.MIN, 0); }
		public TerminalNode MAX() { return getToken(SymbolanParser.MAX, 0); }
		public TerminalNode SIN() { return getToken(SymbolanParser.SIN, 0); }
		public TerminalNode COS() { return getToken(SymbolanParser.COS, 0); }
		public TerminalNode TAN() { return getToken(SymbolanParser.TAN, 0); }
		public TerminalNode LOG() { return getToken(SymbolanParser.LOG, 0); }
		public TerminalNode SINH() { return getToken(SymbolanParser.SINH, 0); }
		public TerminalNode COSH() { return getToken(SymbolanParser.COSH, 0); }
		public TerminalNode TANH() { return getToken(SymbolanParser.TANH, 0); }
		public TerminalNode COTH() { return getToken(SymbolanParser.COTH, 0); }
		public TerminalNode CSCH() { return getToken(SymbolanParser.CSCH, 0); }
		public TerminalNode SECH() { return getToken(SymbolanParser.SECH, 0); }
		public TerminalNode DERIVATIVE_OPERATOR() { return getToken(SymbolanParser.DERIVATIVE_OPERATOR, 0); }
		public System_functionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_system_functions; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterSystem_functions(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitSystem_functions(this);
		}
	}

	public final System_functionsContext system_functions() throws RecognitionException {
		System_functionsContext _localctx = new System_functionsContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_system_functions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(75);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << SIN) | (1L << COS) | (1L << TAN) | (1L << SINH) | (1L << COSH) | (1L << TANH) | (1L << COTH) | (1L << CSCH) | (1L << SECH) | (1L << MIN) | (1L << MAX) | (1L << LOG) | (1L << DERIVATIVE_OPERATOR))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Rule_functionContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(SymbolanParser.IDENTIFIER, 0); }
		public Rule_functionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rule_function; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterRule_function(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitRule_function(this);
		}
	}

	public final Rule_functionContext rule_function() throws RecognitionException {
		Rule_functionContext _localctx = new Rule_functionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_rule_function);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(77);
			match(IDENTIFIER);
			setState(78);
			match(T__0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionRuleContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public Rule_functionContext rule_function() {
			return getRuleContext(Rule_functionContext.class,0);
		}
		public ExpressionRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterExpressionRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitExpressionRule(this);
		}
	}

	public final ExpressionRuleContext expressionRule() throws RecognitionException {
		ExpressionRuleContext _localctx = new ExpressionRuleContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_expressionRule);
		try {
			setState(88);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(80);
				expr(0);
				setState(81);
				match(T__1);
				setState(82);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(84);
				expr(0);
				setState(85);
				match(T__1);
				setState(86);
				rule_function();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExprContext extends ParserRuleContext {
		public Left_parContext left_par() {
			return getRuleContext(Left_parContext.class,0);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public Right_parContext right_par() {
			return getRuleContext(Right_parContext.class,0);
		}
		public System_functionsContext system_functions() {
			return getRuleContext(System_functionsContext.class,0);
		}
		public SignContext sign() {
			return getRuleContext(SignContext.class,0);
		}
		public AtomContext atom() {
			return getRuleContext(AtomContext.class,0);
		}
		public RuleAtomContext ruleAtom() {
			return getRuleContext(RuleAtomContext.class,0);
		}
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitExpr(this);
		}
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case L_PAREN:
				{
				setState(91);
				left_par();
				setState(92);
				expr(0);
				setState(93);
				right_par();
				}
				break;
			case SIN:
			case COS:
			case TAN:
			case SINH:
			case COSH:
			case TANH:
			case COTH:
			case CSCH:
			case SECH:
			case MIN:
			case MAX:
			case LOG:
			case DERIVATIVE_OPERATOR:
				{
				setState(95);
				system_functions();
				setState(96);
				left_par();
				setState(97);
				expr(0);
				setState(98);
				right_par();
				}
				break;
			case T__6:
			case T__7:
				{
				setState(100);
				sign();
				setState(101);
				expr(3);
				}
				break;
			case PI:
			case DERIVATIVE:
			case VAR:
			case CONST:
			case EULER:
			case INFINITE:
			case UNDEFINED:
			case IMAGINARY:
			case DECIMAL_LIT:
			case OCTAL_LIT:
			case HEX_LIT:
			case FLOAT_LIT:
			case IMAGINARY_LIT:
				{
				setState(103);
				atom();
				}
				break;
			case DERIVATE_RULE:
			case NUMERIC_RULE:
			case CONST_RULE:
			case VAR_RULE:
			case EXPR_RULE:
			case NUMERIC_TREE_RULE:
			case CONST_TREE_RULE:
			case VAR_TREE_RULE:
			case SYSTEM_TREE_RULE:
				{
				setState(104);
				ruleAtom();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(120);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(118);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(107);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(108);
						_la = _input.LA(1);
						if ( !(_la==T__2 || _la==T__3) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(109);
						expr(9);
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(110);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(111);
						_la = _input.LA(1);
						if ( !(_la==T__4 || _la==T__5) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(112);
						expr(8);
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(113);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(114);
						_la = _input.LA(1);
						if ( !(_la==T__6 || _la==T__7) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(115);
						expr(7);
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(116);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(117);
						expr(6);
						}
						break;
					}
					} 
				}
				setState(122);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class RuleAtomContext extends ParserRuleContext {
		public Const_ruleContext const_rule() {
			return getRuleContext(Const_ruleContext.class,0);
		}
		public Var_ruleContext var_rule() {
			return getRuleContext(Var_ruleContext.class,0);
		}
		public Expr_ruleContext expr_rule() {
			return getRuleContext(Expr_ruleContext.class,0);
		}
		public Numeric_ruleContext numeric_rule() {
			return getRuleContext(Numeric_ruleContext.class,0);
		}
		public Derivative_ruleContext derivative_rule() {
			return getRuleContext(Derivative_ruleContext.class,0);
		}
		public Numeric_tree_ruleContext numeric_tree_rule() {
			return getRuleContext(Numeric_tree_ruleContext.class,0);
		}
		public Const_tree_ruleContext const_tree_rule() {
			return getRuleContext(Const_tree_ruleContext.class,0);
		}
		public Var_tree_ruleContext var_tree_rule() {
			return getRuleContext(Var_tree_ruleContext.class,0);
		}
		public System_tree_ruleContext system_tree_rule() {
			return getRuleContext(System_tree_ruleContext.class,0);
		}
		public RuleAtomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ruleAtom; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterRuleAtom(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitRuleAtom(this);
		}
	}

	public final RuleAtomContext ruleAtom() throws RecognitionException {
		RuleAtomContext _localctx = new RuleAtomContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_ruleAtom);
		try {
			setState(132);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CONST_RULE:
				enterOuterAlt(_localctx, 1);
				{
				setState(123);
				const_rule();
				}
				break;
			case VAR_RULE:
				enterOuterAlt(_localctx, 2);
				{
				setState(124);
				var_rule();
				}
				break;
			case EXPR_RULE:
				enterOuterAlt(_localctx, 3);
				{
				setState(125);
				expr_rule();
				}
				break;
			case NUMERIC_RULE:
				enterOuterAlt(_localctx, 4);
				{
				setState(126);
				numeric_rule();
				}
				break;
			case DERIVATE_RULE:
				enterOuterAlt(_localctx, 5);
				{
				setState(127);
				derivative_rule();
				}
				break;
			case NUMERIC_TREE_RULE:
				enterOuterAlt(_localctx, 6);
				{
				setState(128);
				numeric_tree_rule();
				}
				break;
			case CONST_TREE_RULE:
				enterOuterAlt(_localctx, 7);
				{
				setState(129);
				const_tree_rule();
				}
				break;
			case VAR_TREE_RULE:
				enterOuterAlt(_localctx, 8);
				{
				setState(130);
				var_tree_rule();
				}
				break;
			case SYSTEM_TREE_RULE:
				enterOuterAlt(_localctx, 9);
				{
				setState(131);
				system_tree_rule();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SignContext extends ParserRuleContext {
		public SignContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sign; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterSign(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitSign(this);
		}
	}

	public final SignContext sign() throws RecognitionException {
		SignContext _localctx = new SignContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_sign);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			_la = _input.LA(1);
			if ( !(_la==T__6 || _la==T__7) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Left_parContext extends ParserRuleContext {
		public Left_parContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_left_par; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterLeft_par(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitLeft_par(this);
		}
	}

	public final Left_parContext left_par() throws RecognitionException {
		Left_parContext _localctx = new Left_parContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_left_par);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
			match(L_PAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Right_parContext extends ParserRuleContext {
		public Right_parContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_right_par; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterRight_par(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitRight_par(this);
		}
	}

	public final Right_parContext right_par() throws RecognitionException {
		Right_parContext _localctx = new Right_parContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_right_par);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(138);
			match(R_PAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Derivative_ruleContext extends ParserRuleContext {
		public TerminalNode DERIVATE_RULE() { return getToken(SymbolanParser.DERIVATE_RULE, 0); }
		public Derivative_ruleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_derivative_rule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterDerivative_rule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitDerivative_rule(this);
		}
	}

	public final Derivative_ruleContext derivative_rule() throws RecognitionException {
		Derivative_ruleContext _localctx = new Derivative_ruleContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_derivative_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(140);
			match(DERIVATE_RULE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Numeric_ruleContext extends ParserRuleContext {
		public TerminalNode NUMERIC_RULE() { return getToken(SymbolanParser.NUMERIC_RULE, 0); }
		public Numeric_ruleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_numeric_rule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterNumeric_rule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitNumeric_rule(this);
		}
	}

	public final Numeric_ruleContext numeric_rule() throws RecognitionException {
		Numeric_ruleContext _localctx = new Numeric_ruleContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_numeric_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(142);
			match(NUMERIC_RULE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Const_ruleContext extends ParserRuleContext {
		public TerminalNode CONST_RULE() { return getToken(SymbolanParser.CONST_RULE, 0); }
		public Const_ruleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_const_rule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterConst_rule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitConst_rule(this);
		}
	}

	public final Const_ruleContext const_rule() throws RecognitionException {
		Const_ruleContext _localctx = new Const_ruleContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_const_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			match(CONST_RULE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Var_ruleContext extends ParserRuleContext {
		public TerminalNode VAR_RULE() { return getToken(SymbolanParser.VAR_RULE, 0); }
		public Var_ruleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_rule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterVar_rule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitVar_rule(this);
		}
	}

	public final Var_ruleContext var_rule() throws RecognitionException {
		Var_ruleContext _localctx = new Var_ruleContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_var_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(146);
			match(VAR_RULE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Expr_ruleContext extends ParserRuleContext {
		public TerminalNode EXPR_RULE() { return getToken(SymbolanParser.EXPR_RULE, 0); }
		public Expr_ruleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_rule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterExpr_rule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitExpr_rule(this);
		}
	}

	public final Expr_ruleContext expr_rule() throws RecognitionException {
		Expr_ruleContext _localctx = new Expr_ruleContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_expr_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			match(EXPR_RULE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Numeric_tree_ruleContext extends ParserRuleContext {
		public TerminalNode NUMERIC_TREE_RULE() { return getToken(SymbolanParser.NUMERIC_TREE_RULE, 0); }
		public Numeric_tree_ruleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_numeric_tree_rule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterNumeric_tree_rule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitNumeric_tree_rule(this);
		}
	}

	public final Numeric_tree_ruleContext numeric_tree_rule() throws RecognitionException {
		Numeric_tree_ruleContext _localctx = new Numeric_tree_ruleContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_numeric_tree_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(150);
			match(NUMERIC_TREE_RULE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Const_tree_ruleContext extends ParserRuleContext {
		public TerminalNode CONST_TREE_RULE() { return getToken(SymbolanParser.CONST_TREE_RULE, 0); }
		public Const_tree_ruleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_const_tree_rule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterConst_tree_rule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitConst_tree_rule(this);
		}
	}

	public final Const_tree_ruleContext const_tree_rule() throws RecognitionException {
		Const_tree_ruleContext _localctx = new Const_tree_ruleContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_const_tree_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			match(CONST_TREE_RULE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Var_tree_ruleContext extends ParserRuleContext {
		public TerminalNode VAR_TREE_RULE() { return getToken(SymbolanParser.VAR_TREE_RULE, 0); }
		public Var_tree_ruleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_tree_rule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterVar_tree_rule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitVar_tree_rule(this);
		}
	}

	public final Var_tree_ruleContext var_tree_rule() throws RecognitionException {
		Var_tree_ruleContext _localctx = new Var_tree_ruleContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_var_tree_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(154);
			match(VAR_TREE_RULE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class System_tree_ruleContext extends ParserRuleContext {
		public TerminalNode SYSTEM_TREE_RULE() { return getToken(SymbolanParser.SYSTEM_TREE_RULE, 0); }
		public System_tree_ruleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_system_tree_rule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterSystem_tree_rule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitSystem_tree_rule(this);
		}
	}

	public final System_tree_ruleContext system_tree_rule() throws RecognitionException {
		System_tree_ruleContext _localctx = new System_tree_ruleContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_system_tree_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(156);
			match(SYSTEM_TREE_RULE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AtomContext extends ParserRuleContext {
		public NumberContext number() {
			return getRuleContext(NumberContext.class,0);
		}
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public ConstantContext constant() {
			return getRuleContext(ConstantContext.class,0);
		}
		public Math_constantContext math_constant() {
			return getRuleContext(Math_constantContext.class,0);
		}
		public ImaginaryContext imaginary() {
			return getRuleContext(ImaginaryContext.class,0);
		}
		public DerivativeContext derivative() {
			return getRuleContext(DerivativeContext.class,0);
		}
		public AtomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atom; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterAtom(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitAtom(this);
		}
	}

	public final AtomContext atom() throws RecognitionException {
		AtomContext _localctx = new AtomContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_atom);
		try {
			setState(165);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(158);
				number();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(159);
				variable();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(160);
				constant();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(161);
				number();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(162);
				math_constant();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(163);
				imaginary();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(164);
				derivative();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Math_constantContext extends ParserRuleContext {
		public TerminalNode PI() { return getToken(SymbolanParser.PI, 0); }
		public TerminalNode EULER() { return getToken(SymbolanParser.EULER, 0); }
		public TerminalNode INFINITE() { return getToken(SymbolanParser.INFINITE, 0); }
		public TerminalNode UNDEFINED() { return getToken(SymbolanParser.UNDEFINED, 0); }
		public TerminalNode IMAGINARY() { return getToken(SymbolanParser.IMAGINARY, 0); }
		public Math_constantContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_math_constant; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterMath_constant(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitMath_constant(this);
		}
	}

	public final Math_constantContext math_constant() throws RecognitionException {
		Math_constantContext _localctx = new Math_constantContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_math_constant);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PI) | (1L << EULER) | (1L << INFINITE) | (1L << UNDEFINED) | (1L << IMAGINARY))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InfiniteContext extends ParserRuleContext {
		public TerminalNode INFINITE() { return getToken(SymbolanParser.INFINITE, 0); }
		public InfiniteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_infinite; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterInfinite(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitInfinite(this);
		}
	}

	public final InfiniteContext infinite() throws RecognitionException {
		InfiniteContext _localctx = new InfiniteContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_infinite);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			match(INFINITE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImaginaryContext extends ParserRuleContext {
		public TerminalNode IMAGINARY_LIT() { return getToken(SymbolanParser.IMAGINARY_LIT, 0); }
		public ImaginaryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_imaginary; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterImaginary(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitImaginary(this);
		}
	}

	public final ImaginaryContext imaginary() throws RecognitionException {
		ImaginaryContext _localctx = new ImaginaryContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_imaginary);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(171);
			match(IMAGINARY_LIT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NumberContext extends ParserRuleContext {
		public TerminalNode DECIMAL_LIT() { return getToken(SymbolanParser.DECIMAL_LIT, 0); }
		public TerminalNode FLOAT_LIT() { return getToken(SymbolanParser.FLOAT_LIT, 0); }
		public TerminalNode OCTAL_LIT() { return getToken(SymbolanParser.OCTAL_LIT, 0); }
		public TerminalNode HEX_LIT() { return getToken(SymbolanParser.HEX_LIT, 0); }
		public NumberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_number; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterNumber(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitNumber(this);
		}
	}

	public final NumberContext number() throws RecognitionException {
		NumberContext _localctx = new NumberContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_number);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(173);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << DECIMAL_LIT) | (1L << OCTAL_LIT) | (1L << HEX_LIT) | (1L << FLOAT_LIT))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DerivativeContext extends ParserRuleContext {
		public TerminalNode DERIVATIVE() { return getToken(SymbolanParser.DERIVATIVE, 0); }
		public DerivativeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_derivative; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterDerivative(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitDerivative(this);
		}
	}

	public final DerivativeContext derivative() throws RecognitionException {
		DerivativeContext _localctx = new DerivativeContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_derivative);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			match(DERIVATIVE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VariableContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(SymbolanParser.VAR, 0); }
		public VariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterVariable(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitVariable(this);
		}
	}

	public final VariableContext variable() throws RecognitionException {
		VariableContext _localctx = new VariableContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_variable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(177);
			match(VAR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConstantContext extends ParserRuleContext {
		public TerminalNode CONST() { return getToken(SymbolanParser.CONST, 0); }
		public ConstantContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constant; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterConstant(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitConstant(this);
		}
	}

	public final ConstantContext constant() throws RecognitionException {
		ConstantContext _localctx = new ConstantContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_constant);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(179);
			match(CONST);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Bi_operatorsContext extends ParserRuleContext {
		public Bi_operatorsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bi_operators; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).enterBi_operators(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SymbolanListener ) ((SymbolanListener)listener).exitBi_operators(this);
		}
	}

	public final Bi_operatorsContext bi_operators() throws RecognitionException {
		Bi_operatorsContext _localctx = new Bi_operatorsContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_bi_operators);
		int _la;
		try {
			setState(184);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
			case T__3:
				enterOuterAlt(_localctx, 1);
				{
				setState(181);
				_la = _input.LA(1);
				if ( !(_la==T__2 || _la==T__3) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case T__4:
			case T__5:
				enterOuterAlt(_localctx, 2);
				{
				setState(182);
				_la = _input.LA(1);
				if ( !(_la==T__4 || _la==T__5) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case T__6:
			case T__7:
				enterOuterAlt(_localctx, 3);
				{
				setState(183);
				_la = _input.LA(1);
				if ( !(_la==T__6 || _la==T__7) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 5:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 8);
		case 1:
			return precpred(_ctx, 7);
		case 2:
			return precpred(_ctx, 6);
		case 3:
			return precpred(_ctx, 5);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3>\u00bd\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\3\2\3\2\7\2=\n\2\f\2\16\2@\13"+
		"\2\7\2B\n\2\f\2\16\2E\13\2\3\3\3\3\3\3\3\3\3\3\5\3L\n\3\3\4\3\4\3\5\3"+
		"\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6[\n\6\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7l\n\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\7\7y\n\7\f\7\16\7|\13\7\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\5\b\u0087\n\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3"+
		"\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3"+
		"\25\3\25\3\25\3\25\3\25\3\25\3\25\5\25\u00a8\n\25\3\26\3\26\3\27\3\27"+
		"\3\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\35\5\35"+
		"\u00bb\n\35\3\35\2\3\f\36\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&"+
		"(*,.\60\62\64\668\2\b\5\2\24\26\32\"$$\3\2\5\6\3\2\7\b\3\2\t\n\5\2%%)"+
		"+\65\65\3\2\669\2\u00bc\2C\3\2\2\2\4K\3\2\2\2\6M\3\2\2\2\bO\3\2\2\2\n"+
		"Z\3\2\2\2\fk\3\2\2\2\16\u0086\3\2\2\2\20\u0088\3\2\2\2\22\u008a\3\2\2"+
		"\2\24\u008c\3\2\2\2\26\u008e\3\2\2\2\30\u0090\3\2\2\2\32\u0092\3\2\2\2"+
		"\34\u0094\3\2\2\2\36\u0096\3\2\2\2 \u0098\3\2\2\2\"\u009a\3\2\2\2$\u009c"+
		"\3\2\2\2&\u009e\3\2\2\2(\u00a7\3\2\2\2*\u00a9\3\2\2\2,\u00ab\3\2\2\2."+
		"\u00ad\3\2\2\2\60\u00af\3\2\2\2\62\u00b1\3\2\2\2\64\u00b3\3\2\2\2\66\u00b5"+
		"\3\2\2\28\u00ba\3\2\2\2:>\5\n\6\2;=\7<\2\2<;\3\2\2\2=@\3\2\2\2><\3\2\2"+
		"\2>?\3\2\2\2?B\3\2\2\2@>\3\2\2\2A:\3\2\2\2BE\3\2\2\2CA\3\2\2\2CD\3\2\2"+
		"\2D\3\3\2\2\2EC\3\2\2\2FL\5\f\7\2GH\5\f\7\2HI\7\21\2\2IJ\5\f\7\2JL\3\2"+
		"\2\2KF\3\2\2\2KG\3\2\2\2L\5\3\2\2\2MN\t\2\2\2N\7\3\2\2\2OP\7;\2\2PQ\7"+
		"\3\2\2Q\t\3\2\2\2RS\5\f\7\2ST\7\4\2\2TU\5\f\7\2U[\3\2\2\2VW\5\f\7\2WX"+
		"\7\4\2\2XY\5\b\5\2Y[\3\2\2\2ZR\3\2\2\2ZV\3\2\2\2[\13\3\2\2\2\\]\b\7\1"+
		"\2]^\5\22\n\2^_\5\f\7\2_`\5\24\13\2`l\3\2\2\2ab\5\6\4\2bc\5\22\n\2cd\5"+
		"\f\7\2de\5\24\13\2el\3\2\2\2fg\5\20\t\2gh\5\f\7\5hl\3\2\2\2il\5(\25\2"+
		"jl\5\16\b\2k\\\3\2\2\2ka\3\2\2\2kf\3\2\2\2ki\3\2\2\2kj\3\2\2\2lz\3\2\2"+
		"\2mn\f\n\2\2no\t\3\2\2oy\5\f\7\13pq\f\t\2\2qr\t\4\2\2ry\5\f\7\nst\f\b"+
		"\2\2tu\t\5\2\2uy\5\f\7\tvw\f\7\2\2wy\5\f\7\bxm\3\2\2\2xp\3\2\2\2xs\3\2"+
		"\2\2xv\3\2\2\2y|\3\2\2\2zx\3\2\2\2z{\3\2\2\2{\r\3\2\2\2|z\3\2\2\2}\u0087"+
		"\5\32\16\2~\u0087\5\34\17\2\177\u0087\5\36\20\2\u0080\u0087\5\30\r\2\u0081"+
		"\u0087\5\26\f\2\u0082\u0087\5 \21\2\u0083\u0087\5\"\22\2\u0084\u0087\5"+
		"$\23\2\u0085\u0087\5&\24\2\u0086}\3\2\2\2\u0086~\3\2\2\2\u0086\177\3\2"+
		"\2\2\u0086\u0080\3\2\2\2\u0086\u0081\3\2\2\2\u0086\u0082\3\2\2\2\u0086"+
		"\u0083\3\2\2\2\u0086\u0084\3\2\2\2\u0086\u0085\3\2\2\2\u0087\17\3\2\2"+
		"\2\u0088\u0089\t\5\2\2\u0089\21\3\2\2\2\u008a\u008b\7\13\2\2\u008b\23"+
		"\3\2\2\2\u008c\u008d\7\f\2\2\u008d\25\3\2\2\2\u008e\u008f\7,\2\2\u008f"+
		"\27\3\2\2\2\u0090\u0091\7-\2\2\u0091\31\3\2\2\2\u0092\u0093\7.\2\2\u0093"+
		"\33\3\2\2\2\u0094\u0095\7/\2\2\u0095\35\3\2\2\2\u0096\u0097\7\60\2\2\u0097"+
		"\37\3\2\2\2\u0098\u0099\7\61\2\2\u0099!\3\2\2\2\u009a\u009b\7\62\2\2\u009b"+
		"#\3\2\2\2\u009c\u009d\7\63\2\2\u009d%\3\2\2\2\u009e\u009f\7\64\2\2\u009f"+
		"\'\3\2\2\2\u00a0\u00a8\5\60\31\2\u00a1\u00a8\5\64\33\2\u00a2\u00a8\5\66"+
		"\34\2\u00a3\u00a8\5\60\31\2\u00a4\u00a8\5*\26\2\u00a5\u00a8\5.\30\2\u00a6"+
		"\u00a8\5\62\32\2\u00a7\u00a0\3\2\2\2\u00a7\u00a1\3\2\2\2\u00a7\u00a2\3"+
		"\2\2\2\u00a7\u00a3\3\2\2\2\u00a7\u00a4\3\2\2\2\u00a7\u00a5\3\2\2\2\u00a7"+
		"\u00a6\3\2\2\2\u00a8)\3\2\2\2\u00a9\u00aa\t\6\2\2\u00aa+\3\2\2\2\u00ab"+
		"\u00ac\7*\2\2\u00ac-\3\2\2\2\u00ad\u00ae\7:\2\2\u00ae/\3\2\2\2\u00af\u00b0"+
		"\t\7\2\2\u00b0\61\3\2\2\2\u00b1\u00b2\7&\2\2\u00b2\63\3\2\2\2\u00b3\u00b4"+
		"\7\'\2\2\u00b4\65\3\2\2\2\u00b5\u00b6\7(\2\2\u00b6\67\3\2\2\2\u00b7\u00bb"+
		"\t\3\2\2\u00b8\u00bb\t\4\2\2\u00b9\u00bb\t\5\2\2\u00ba\u00b7\3\2\2\2\u00ba"+
		"\u00b8\3\2\2\2\u00ba\u00b9\3\2\2\2\u00bb9\3\2\2\2\f>CKZkxz\u0086\u00a7"+
		"\u00ba";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}