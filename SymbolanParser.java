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
		COMMA=16, FACTORIAL=17, MIN=18, MAX=19, SIN=20, COS=21, TAN=22, LOG=23, 
		SQRT=24, PI=25, VAR=26, CONST=27, EULER=28, INFINITE=29, UNDEFINED=30, 
		NUMERIC_RULE=31, CONST_RULE=32, VAR_RULE=33, EXPR_RULE=34, IMAGINARY=35, 
		DECIMAL_LIT=36, OCTAL_LIT=37, HEX_LIT=38, FLOAT_LIT=39, IMAGINARY_LIT=40, 
		IDENTIFIER=41, NEWLINE=42, WS=43, LINE_COMMENT=44;
	public static final int
		RULE_ruleSet = 0, RULE_eq = 1, RULE_system_functions = 2, RULE_rule_function = 3, 
		RULE_expressionRule = 4, RULE_expr = 5, RULE_ruleAtom = 6, RULE_numeric_rule = 7, 
		RULE_const_rule = 8, RULE_var_rule = 9, RULE_expr_rule = 10, RULE_left_par = 11, 
		RULE_right_par = 12, RULE_atom = 13, RULE_math_constant = 14, RULE_infinite = 15, 
		RULE_imaginary = 16, RULE_number = 17, RULE_variable = 18, RULE_constant = 19, 
		RULE_bi_operators = 20;
	public static final String[] ruleNames = {
		"ruleSet", "eq", "system_functions", "rule_function", "expressionRule", 
		"expr", "ruleAtom", "numeric_rule", "const_rule", "var_rule", "expr_rule", 
		"left_par", "right_par", "atom", "math_constant", "infinite", "imaginary", 
		"number", "variable", "constant", "bi_operators"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'()'", "'=>'", "'^'", "'**'", "'*'", "'/'", "'+'", "'-'", "'('", 
		"')'", "'{'", "'}'", "'['", "']'", "'='", "','", "'!'", "'min'", "'max'", 
		"'sin'", "'cos'", "'tan'", "'log'", "'sqrt'", "'pi'", null, null, null, 
		null, "'undef'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, null, null, "L_PAREN", "R_PAREN", 
		"L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "FACTORIAL", 
		"MIN", "MAX", "SIN", "COS", "TAN", "LOG", "SQRT", "PI", "VAR", "CONST", 
		"EULER", "INFINITE", "UNDEFINED", "NUMERIC_RULE", "CONST_RULE", "VAR_RULE", 
		"EXPR_RULE", "IMAGINARY", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", 
		"IMAGINARY_LIT", "IDENTIFIER", "NEWLINE", "WS", "LINE_COMMENT"
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
			setState(51);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << L_PAREN) | (1L << MIN) | (1L << MAX) | (1L << SIN) | (1L << COS) | (1L << TAN) | (1L << LOG) | (1L << PI) | (1L << VAR) | (1L << CONST) | (1L << EULER) | (1L << INFINITE) | (1L << UNDEFINED) | (1L << NUMERIC_RULE) | (1L << CONST_RULE) | (1L << VAR_RULE) | (1L << EXPR_RULE) | (1L << DECIMAL_LIT) | (1L << OCTAL_LIT) | (1L << HEX_LIT) | (1L << FLOAT_LIT) | (1L << IMAGINARY_LIT))) != 0)) {
				{
				{
				setState(42);
				expressionRule();
				setState(46);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==NEWLINE) {
					{
					{
					setState(43);
					match(NEWLINE);
					}
					}
					setState(48);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				}
				setState(53);
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
			setState(59);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(54);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(55);
				expr(0);
				setState(56);
				match(ASSIGN);
				setState(57);
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
			setState(61);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MIN) | (1L << MAX) | (1L << SIN) | (1L << COS) | (1L << TAN) | (1L << LOG))) != 0)) ) {
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
			setState(63);
			match(IDENTIFIER);
			setState(64);
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
			setState(74);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(66);
				expr(0);
				setState(67);
				match(T__1);
				setState(68);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(70);
				expr(0);
				setState(71);
				match(T__1);
				setState(72);
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
			setState(88);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case L_PAREN:
				{
				setState(77);
				left_par();
				setState(78);
				expr(0);
				setState(79);
				right_par();
				}
				break;
			case MIN:
			case MAX:
			case SIN:
			case COS:
			case TAN:
			case LOG:
				{
				setState(81);
				system_functions();
				setState(82);
				left_par();
				setState(83);
				expr(0);
				setState(84);
				right_par();
				}
				break;
			case PI:
			case VAR:
			case CONST:
			case EULER:
			case INFINITE:
			case UNDEFINED:
			case DECIMAL_LIT:
			case OCTAL_LIT:
			case HEX_LIT:
			case FLOAT_LIT:
			case IMAGINARY_LIT:
				{
				setState(86);
				atom();
				}
				break;
			case NUMERIC_RULE:
			case CONST_RULE:
			case VAR_RULE:
			case EXPR_RULE:
				{
				setState(87);
				ruleAtom();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(103);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(101);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(90);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(91);
						expr(9);
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(92);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(93);
						_la = _input.LA(1);
						if ( !(_la==T__2 || _la==T__3) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(94);
						expr(7);
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(95);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(96);
						_la = _input.LA(1);
						if ( !(_la==T__4 || _la==T__5) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(97);
						expr(6);
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(98);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(99);
						_la = _input.LA(1);
						if ( !(_la==T__6 || _la==T__7) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(100);
						expr(5);
						}
						break;
					}
					} 
				}
				setState(105);
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
			setState(110);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CONST_RULE:
				enterOuterAlt(_localctx, 1);
				{
				setState(106);
				const_rule();
				}
				break;
			case VAR_RULE:
				enterOuterAlt(_localctx, 2);
				{
				setState(107);
				var_rule();
				}
				break;
			case EXPR_RULE:
				enterOuterAlt(_localctx, 3);
				{
				setState(108);
				expr_rule();
				}
				break;
			case NUMERIC_RULE:
				enterOuterAlt(_localctx, 4);
				{
				setState(109);
				numeric_rule();
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
		enterRule(_localctx, 14, RULE_numeric_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(112);
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
		enterRule(_localctx, 16, RULE_const_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
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
		enterRule(_localctx, 18, RULE_var_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
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
		enterRule(_localctx, 20, RULE_expr_rule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(118);
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
		enterRule(_localctx, 22, RULE_left_par);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(120);
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
		enterRule(_localctx, 24, RULE_right_par);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(122);
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
		enterRule(_localctx, 26, RULE_atom);
		try {
			setState(130);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(124);
				number();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(125);
				variable();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(126);
				constant();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(127);
				number();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(128);
				math_constant();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(129);
				imaginary();
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
		enterRule(_localctx, 28, RULE_math_constant);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PI) | (1L << EULER) | (1L << INFINITE) | (1L << UNDEFINED))) != 0)) ) {
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
		enterRule(_localctx, 30, RULE_infinite);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
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
		enterRule(_localctx, 32, RULE_imaginary);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
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
		enterRule(_localctx, 34, RULE_number);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(138);
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
		enterRule(_localctx, 36, RULE_variable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(140);
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
		enterRule(_localctx, 38, RULE_constant);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(142);
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
		enterRule(_localctx, 40, RULE_bi_operators);
		int _la;
		try {
			setState(147);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
			case T__3:
				enterOuterAlt(_localctx, 1);
				{
				setState(144);
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
				setState(145);
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
				setState(146);
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
			return precpred(_ctx, 6);
		case 2:
			return precpred(_ctx, 5);
		case 3:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3.\u0098\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\3\2\3\2\7\2/\n\2\f\2\16\2\62"+
		"\13\2\7\2\64\n\2\f\2\16\2\67\13\2\3\3\3\3\3\3\3\3\3\3\5\3>\n\3\3\4\3\4"+
		"\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6M\n\6\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7[\n\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\7\7h\n\7\f\7\16\7k\13\7\3\b\3\b\3\b\3\b\5\bq\n\b\3\t"+
		"\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\17\3\17\3"+
		"\17\3\17\5\17\u0085\n\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24"+
		"\3\24\3\25\3\25\3\26\3\26\3\26\5\26\u0096\n\26\3\26\2\3\f\27\2\4\6\b\n"+
		"\f\16\20\22\24\26\30\32\34\36 \"$&(*\2\b\3\2\24\31\3\2\5\6\3\2\7\b\3\2"+
		"\t\n\4\2\33\33\36 \3\2&)\2\u0097\2\65\3\2\2\2\4=\3\2\2\2\6?\3\2\2\2\b"+
		"A\3\2\2\2\nL\3\2\2\2\fZ\3\2\2\2\16p\3\2\2\2\20r\3\2\2\2\22t\3\2\2\2\24"+
		"v\3\2\2\2\26x\3\2\2\2\30z\3\2\2\2\32|\3\2\2\2\34\u0084\3\2\2\2\36\u0086"+
		"\3\2\2\2 \u0088\3\2\2\2\"\u008a\3\2\2\2$\u008c\3\2\2\2&\u008e\3\2\2\2"+
		"(\u0090\3\2\2\2*\u0095\3\2\2\2,\60\5\n\6\2-/\7,\2\2.-\3\2\2\2/\62\3\2"+
		"\2\2\60.\3\2\2\2\60\61\3\2\2\2\61\64\3\2\2\2\62\60\3\2\2\2\63,\3\2\2\2"+
		"\64\67\3\2\2\2\65\63\3\2\2\2\65\66\3\2\2\2\66\3\3\2\2\2\67\65\3\2\2\2"+
		"8>\5\f\7\29:\5\f\7\2:;\7\21\2\2;<\5\f\7\2<>\3\2\2\2=8\3\2\2\2=9\3\2\2"+
		"\2>\5\3\2\2\2?@\t\2\2\2@\7\3\2\2\2AB\7+\2\2BC\7\3\2\2C\t\3\2\2\2DE\5\f"+
		"\7\2EF\7\4\2\2FG\5\f\7\2GM\3\2\2\2HI\5\f\7\2IJ\7\4\2\2JK\5\b\5\2KM\3\2"+
		"\2\2LD\3\2\2\2LH\3\2\2\2M\13\3\2\2\2NO\b\7\1\2OP\5\30\r\2PQ\5\f\7\2QR"+
		"\5\32\16\2R[\3\2\2\2ST\5\6\4\2TU\5\30\r\2UV\5\f\7\2VW\5\32\16\2W[\3\2"+
		"\2\2X[\5\34\17\2Y[\5\16\b\2ZN\3\2\2\2ZS\3\2\2\2ZX\3\2\2\2ZY\3\2\2\2[i"+
		"\3\2\2\2\\]\f\n\2\2]h\5\f\7\13^_\f\b\2\2_`\t\3\2\2`h\5\f\7\tab\f\7\2\2"+
		"bc\t\4\2\2ch\5\f\7\bde\f\6\2\2ef\t\5\2\2fh\5\f\7\7g\\\3\2\2\2g^\3\2\2"+
		"\2ga\3\2\2\2gd\3\2\2\2hk\3\2\2\2ig\3\2\2\2ij\3\2\2\2j\r\3\2\2\2ki\3\2"+
		"\2\2lq\5\22\n\2mq\5\24\13\2nq\5\26\f\2oq\5\20\t\2pl\3\2\2\2pm\3\2\2\2"+
		"pn\3\2\2\2po\3\2\2\2q\17\3\2\2\2rs\7!\2\2s\21\3\2\2\2tu\7\"\2\2u\23\3"+
		"\2\2\2vw\7#\2\2w\25\3\2\2\2xy\7$\2\2y\27\3\2\2\2z{\7\13\2\2{\31\3\2\2"+
		"\2|}\7\f\2\2}\33\3\2\2\2~\u0085\5$\23\2\177\u0085\5&\24\2\u0080\u0085"+
		"\5(\25\2\u0081\u0085\5$\23\2\u0082\u0085\5\36\20\2\u0083\u0085\5\"\22"+
		"\2\u0084~\3\2\2\2\u0084\177\3\2\2\2\u0084\u0080\3\2\2\2\u0084\u0081\3"+
		"\2\2\2\u0084\u0082\3\2\2\2\u0084\u0083\3\2\2\2\u0085\35\3\2\2\2\u0086"+
		"\u0087\t\6\2\2\u0087\37\3\2\2\2\u0088\u0089\7\37\2\2\u0089!\3\2\2\2\u008a"+
		"\u008b\7*\2\2\u008b#\3\2\2\2\u008c\u008d\t\7\2\2\u008d%\3\2\2\2\u008e"+
		"\u008f\7\34\2\2\u008f\'\3\2\2\2\u0090\u0091\7\35\2\2\u0091)\3\2\2\2\u0092"+
		"\u0096\t\3\2\2\u0093\u0096\t\4\2\2\u0094\u0096\t\5\2\2\u0095\u0092\3\2"+
		"\2\2\u0095\u0093\3\2\2\2\u0095\u0094\3\2\2\2\u0096+\3\2\2\2\f\60\65=L"+
		"Zgip\u0084\u0095";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}