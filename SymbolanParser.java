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
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, L_PAREN=8, R_PAREN=9, 
		L_CURLY=10, R_CURLY=11, L_BRACKET=12, R_BRACKET=13, ASSIGN=14, COMMA=15, 
		FACTORIAL=16, MIN=17, MAX=18, SIN=19, COS=20, TAN=21, LOG=22, PI=23, VAR=24, 
		CONST=25, EULER=26, CONST_RULE=27, VAR_RULE=28, EXP_RULE=29, IMAGINARY=30, 
		DECIMAL_LIT=31, OCTAL_LIT=32, HEX_LIT=33, FLOAT_LIT=34, IMAGINARY_LIT=35, 
		WS=36, LINE_COMMENT=37;
	public static final int
		RULE_eq = 0, RULE_system_functions = 1, RULE_expressionRule = 2, RULE_expr = 3, 
		RULE_left_par = 4, RULE_right_par = 5, RULE_atom = 6, RULE_number = 7, 
		RULE_variable = 8, RULE_constant = 9, RULE_bi_operators = 10;
	public static final String[] ruleNames = {
		"eq", "system_functions", "expressionRule", "expr", "left_par", "right_par", 
		"atom", "number", "variable", "constant", "bi_operators"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'=>'", "'^'", "'**'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", 
		"'{'", "'}'", "'['", "']'", "'='", "','", "'!'", "'min'", "'max'", "'sin'", 
		"'cos'", "'tan'", "'log'", "'pi'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, null, "L_PAREN", "R_PAREN", 
		"L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "FACTORIAL", 
		"MIN", "MAX", "SIN", "COS", "TAN", "LOG", "PI", "VAR", "CONST", "EULER", 
		"CONST_RULE", "VAR_RULE", "EXP_RULE", "IMAGINARY", "DECIMAL_LIT", "OCTAL_LIT", 
		"HEX_LIT", "FLOAT_LIT", "IMAGINARY_LIT", "WS", "LINE_COMMENT"
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
		enterRule(_localctx, 0, RULE_eq);
		try {
			setState(27);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(22);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(23);
				expr(0);
				setState(24);
				match(ASSIGN);
				setState(25);
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
		enterRule(_localctx, 2, RULE_system_functions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(29);
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

	public static class ExpressionRuleContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
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
		enterRule(_localctx, 4, RULE_expressionRule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(31);
			expr(0);
			setState(32);
			match(T__0);
			setState(33);
			expr(0);
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
		public AtomContext atom() {
			return getRuleContext(AtomContext.class,0);
		}
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
		int _startState = 6;
		enterRecursionRule(_localctx, 6, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PI:
			case VAR:
			case CONST:
			case EULER:
			case DECIMAL_LIT:
			case OCTAL_LIT:
			case HEX_LIT:
			case FLOAT_LIT:
			case IMAGINARY_LIT:
				{
				setState(36);
				atom();
				}
				break;
			case L_PAREN:
				{
				setState(37);
				left_par();
				setState(38);
				expr(0);
				setState(39);
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
				setState(41);
				system_functions();
				setState(42);
				left_par();
				setState(43);
				expr(0);
				setState(44);
				right_par();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(61);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(59);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(48);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(49);
						expr(8);
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(50);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(51);
						_la = _input.LA(1);
						if ( !(_la==T__1 || _la==T__2) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(52);
						expr(5);
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(53);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(54);
						_la = _input.LA(1);
						if ( !(_la==T__3 || _la==T__4) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(55);
						expr(4);
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(56);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(57);
						_la = _input.LA(1);
						if ( !(_la==T__5 || _la==T__6) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(58);
						expr(3);
						}
						break;
					}
					} 
				}
				setState(63);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
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
		enterRule(_localctx, 8, RULE_left_par);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(64);
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
		enterRule(_localctx, 10, RULE_right_par);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
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
		enterRule(_localctx, 12, RULE_atom);
		try {
			setState(72);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(68);
				number();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(69);
				variable();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(70);
				constant();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(71);
				number();
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

	public static class NumberContext extends ParserRuleContext {
		public TerminalNode DECIMAL_LIT() { return getToken(SymbolanParser.DECIMAL_LIT, 0); }
		public TerminalNode FLOAT_LIT() { return getToken(SymbolanParser.FLOAT_LIT, 0); }
		public TerminalNode IMAGINARY_LIT() { return getToken(SymbolanParser.IMAGINARY_LIT, 0); }
		public TerminalNode OCTAL_LIT() { return getToken(SymbolanParser.OCTAL_LIT, 0); }
		public TerminalNode HEX_LIT() { return getToken(SymbolanParser.HEX_LIT, 0); }
		public TerminalNode PI() { return getToken(SymbolanParser.PI, 0); }
		public TerminalNode EULER() { return getToken(SymbolanParser.EULER, 0); }
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
		enterRule(_localctx, 14, RULE_number);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(74);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PI) | (1L << EULER) | (1L << DECIMAL_LIT) | (1L << OCTAL_LIT) | (1L << HEX_LIT) | (1L << FLOAT_LIT) | (1L << IMAGINARY_LIT))) != 0)) ) {
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
		enterRule(_localctx, 16, RULE_variable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(76);
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
		enterRule(_localctx, 18, RULE_constant);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
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
		enterRule(_localctx, 20, RULE_bi_operators);
		int _la;
		try {
			setState(83);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
			case T__2:
				enterOuterAlt(_localctx, 1);
				{
				setState(80);
				_la = _input.LA(1);
				if ( !(_la==T__1 || _la==T__2) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case T__3:
			case T__4:
				enterOuterAlt(_localctx, 2);
				{
				setState(81);
				_la = _input.LA(1);
				if ( !(_la==T__3 || _la==T__4) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case T__5:
			case T__6:
				enterOuterAlt(_localctx, 3);
				{
				setState(82);
				_la = _input.LA(1);
				if ( !(_la==T__5 || _la==T__6) ) {
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
		case 3:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 7);
		case 1:
			return precpred(_ctx, 4);
		case 2:
			return precpred(_ctx, 3);
		case 3:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\'X\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t\13\4"+
		"\f\t\f\3\2\3\2\3\2\3\2\3\2\5\2\36\n\2\3\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5\61\n\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\7\5>\n\5\f\5\16\5A\13\5\3\6\3\6\3\7\3\7\3\b\3\b"+
		"\3\b\3\b\5\bK\n\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\f\5\fV\n\f\3\f\2"+
		"\3\b\r\2\4\6\b\n\f\16\20\22\24\26\2\7\3\2\23\30\3\2\4\5\3\2\6\7\3\2\b"+
		"\t\5\2\31\31\34\34!%\2X\2\35\3\2\2\2\4\37\3\2\2\2\6!\3\2\2\2\b\60\3\2"+
		"\2\2\nB\3\2\2\2\fD\3\2\2\2\16J\3\2\2\2\20L\3\2\2\2\22N\3\2\2\2\24P\3\2"+
		"\2\2\26U\3\2\2\2\30\36\5\b\5\2\31\32\5\b\5\2\32\33\7\20\2\2\33\34\5\b"+
		"\5\2\34\36\3\2\2\2\35\30\3\2\2\2\35\31\3\2\2\2\36\3\3\2\2\2\37 \t\2\2"+
		"\2 \5\3\2\2\2!\"\5\b\5\2\"#\7\3\2\2#$\5\b\5\2$\7\3\2\2\2%&\b\5\1\2&\61"+
		"\5\16\b\2\'(\5\n\6\2()\5\b\5\2)*\5\f\7\2*\61\3\2\2\2+,\5\4\3\2,-\5\n\6"+
		"\2-.\5\b\5\2./\5\f\7\2/\61\3\2\2\2\60%\3\2\2\2\60\'\3\2\2\2\60+\3\2\2"+
		"\2\61?\3\2\2\2\62\63\f\t\2\2\63>\5\b\5\n\64\65\f\6\2\2\65\66\t\3\2\2\66"+
		">\5\b\5\7\678\f\5\2\289\t\4\2\29>\5\b\5\6:;\f\4\2\2;<\t\5\2\2<>\5\b\5"+
		"\5=\62\3\2\2\2=\64\3\2\2\2=\67\3\2\2\2=:\3\2\2\2>A\3\2\2\2?=\3\2\2\2?"+
		"@\3\2\2\2@\t\3\2\2\2A?\3\2\2\2BC\7\n\2\2C\13\3\2\2\2DE\7\13\2\2E\r\3\2"+
		"\2\2FK\5\20\t\2GK\5\22\n\2HK\5\24\13\2IK\5\20\t\2JF\3\2\2\2JG\3\2\2\2"+
		"JH\3\2\2\2JI\3\2\2\2K\17\3\2\2\2LM\t\6\2\2M\21\3\2\2\2NO\7\32\2\2O\23"+
		"\3\2\2\2PQ\7\33\2\2Q\25\3\2\2\2RV\t\3\2\2SV\t\4\2\2TV\t\5\2\2UR\3\2\2"+
		"\2US\3\2\2\2UT\3\2\2\2V\27\3\2\2\2\b\35\60=?JU";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}