// Generated from TreeLan.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class TreeLanParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, ADDITION=5, SUBTRACTION=6, MULTIPLICATION=7, 
		DIVISION=8, EXPONENTIAL=9, LOGARITHM=10, SIN=11, COS=12, TAN=13, TRANSFORMATION=14, 
		EQUATION=15, LEAF=16, ANY=17, SPECIAL_NUMBER=18, NUMERIC_CONSTANT=19, 
		CONSTANT=20, VARIABLE=21, CONSTANT_EXPRESSION=22, VARIABLE_EXPRESSION=23, 
		NUMERIC_CONSTANT_EXPRESSION=24, SYSTEM_FUNCTION=25, ID=26, ASSIGN=27, 
		NUMBER=28, WS=29;
	public static final int
		RULE_treeRule = 0, RULE_formula = 1, RULE_treeDescriptor = 2, RULE_rootNode = 3, 
		RULE_leafNode = 4, RULE_classValue = 5, RULE_operationClass = 6;
	public static final String[] ruleNames = {
		"treeRule", "formula", "treeDescriptor", "rootNode", "leafNode", "classValue", 
		"operationClass"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'F_'", "'('", "')'", "'_'", "'ADDITION'", "'SUBTRACTION'", "'MULTIPLICATION'", 
		"'DIVISION'", "'EXPONENTIAL'", "'LOGARITHM'", "'SIN'", "'COS'", "'TAN'", 
		"'TRANSFORMATION'", "'EQUATION'", "'LEAF'", "'ANY'", null, "'NUMERIC_CONSTANT'", 
		"'CONSTANT'", "'VARIABLE'", "'CONSTANT_EXPRESSION'", "'VARIABLE_EXPRESSION'", 
		"'NUMERIC_CONSTANT_EXPRESSION'", "'SYSTEM_FUNCTION'", null, "'=>'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, "ADDITION", "SUBTRACTION", "MULTIPLICATION", 
		"DIVISION", "EXPONENTIAL", "LOGARITHM", "SIN", "COS", "TAN", "TRANSFORMATION", 
		"EQUATION", "LEAF", "ANY", "SPECIAL_NUMBER", "NUMERIC_CONSTANT", "CONSTANT", 
		"VARIABLE", "CONSTANT_EXPRESSION", "VARIABLE_EXPRESSION", "NUMERIC_CONSTANT_EXPRESSION", 
		"SYSTEM_FUNCTION", "ID", "ASSIGN", "NUMBER", "WS"
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
	public String getGrammarFileName() { return "TreeLan.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public TreeLanParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class TreeRuleContext extends ParserRuleContext {
		public List<TreeDescriptorContext> treeDescriptor() {
			return getRuleContexts(TreeDescriptorContext.class);
		}
		public TreeDescriptorContext treeDescriptor(int i) {
			return getRuleContext(TreeDescriptorContext.class,i);
		}
		public TerminalNode ASSIGN() { return getToken(TreeLanParser.ASSIGN, 0); }
		public FormulaContext formula() {
			return getRuleContext(FormulaContext.class,0);
		}
		public TreeRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_treeRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).enterTreeRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).exitTreeRule(this);
		}
	}

	public final TreeRuleContext treeRule() throws RecognitionException {
		TreeRuleContext _localctx = new TreeRuleContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_treeRule);
		try {
			setState(23);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(14);
				treeDescriptor();
				setState(15);
				match(ASSIGN);
				setState(16);
				treeDescriptor();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(18);
				treeDescriptor();
				setState(19);
				match(ASSIGN);
				setState(20);
				match(T__0);
				setState(21);
				formula();
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

	public static class FormulaContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(TreeLanParser.ID, 0); }
		public FormulaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_formula; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).enterFormula(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).exitFormula(this);
		}
	}

	public final FormulaContext formula() throws RecognitionException {
		FormulaContext _localctx = new FormulaContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_formula);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(25);
			match(ID);
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

	public static class TreeDescriptorContext extends ParserRuleContext {
		public List<TreeDescriptorContext> treeDescriptor() {
			return getRuleContexts(TreeDescriptorContext.class);
		}
		public TreeDescriptorContext treeDescriptor(int i) {
			return getRuleContext(TreeDescriptorContext.class,i);
		}
		public RootNodeContext rootNode() {
			return getRuleContext(RootNodeContext.class,0);
		}
		public LeafNodeContext leafNode() {
			return getRuleContext(LeafNodeContext.class,0);
		}
		public TreeDescriptorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_treeDescriptor; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).enterTreeDescriptor(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).exitTreeDescriptor(this);
		}
	}

	public final TreeDescriptorContext treeDescriptor() throws RecognitionException {
		TreeDescriptorContext _localctx = new TreeDescriptorContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_treeDescriptor);
		try {
			setState(36);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
				enterOuterAlt(_localctx, 1);
				{
				setState(27);
				match(T__1);
				setState(28);
				treeDescriptor();
				setState(29);
				match(T__2);
				setState(30);
				rootNode();
				setState(31);
				match(T__1);
				setState(32);
				treeDescriptor();
				setState(33);
				match(T__2);
				}
				break;
			case ANY:
			case SPECIAL_NUMBER:
			case NUMERIC_CONSTANT:
			case CONSTANT:
			case VARIABLE:
			case CONSTANT_EXPRESSION:
			case VARIABLE_EXPRESSION:
			case NUMERIC_CONSTANT_EXPRESSION:
			case SYSTEM_FUNCTION:
			case NUMBER:
				enterOuterAlt(_localctx, 2);
				{
				setState(35);
				leafNode();
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

	public static class RootNodeContext extends ParserRuleContext {
		public OperationClassContext operationClass() {
			return getRuleContext(OperationClassContext.class,0);
		}
		public TerminalNode ID() { return getToken(TreeLanParser.ID, 0); }
		public RootNodeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rootNode; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).enterRootNode(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).exitRootNode(this);
		}
	}

	public final RootNodeContext rootNode() throws RecognitionException {
		RootNodeContext _localctx = new RootNodeContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_rootNode);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(38);
			operationClass();
			setState(39);
			match(T__3);
			setState(40);
			match(ID);
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

	public static class LeafNodeContext extends ParserRuleContext {
		public ClassValueContext classValue() {
			return getRuleContext(ClassValueContext.class,0);
		}
		public TerminalNode ID() { return getToken(TreeLanParser.ID, 0); }
		public LeafNodeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_leafNode; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).enterLeafNode(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).exitLeafNode(this);
		}
	}

	public final LeafNodeContext leafNode() throws RecognitionException {
		LeafNodeContext _localctx = new LeafNodeContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_leafNode);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(42);
			classValue();
			setState(43);
			match(T__3);
			setState(44);
			match(ID);
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

	public static class ClassValueContext extends ParserRuleContext {
		public TerminalNode NUMERIC_CONSTANT() { return getToken(TreeLanParser.NUMERIC_CONSTANT, 0); }
		public TerminalNode CONSTANT() { return getToken(TreeLanParser.CONSTANT, 0); }
		public TerminalNode VARIABLE() { return getToken(TreeLanParser.VARIABLE, 0); }
		public TerminalNode CONSTANT_EXPRESSION() { return getToken(TreeLanParser.CONSTANT_EXPRESSION, 0); }
		public TerminalNode VARIABLE_EXPRESSION() { return getToken(TreeLanParser.VARIABLE_EXPRESSION, 0); }
		public TerminalNode NUMERIC_CONSTANT_EXPRESSION() { return getToken(TreeLanParser.NUMERIC_CONSTANT_EXPRESSION, 0); }
		public TerminalNode SYSTEM_FUNCTION() { return getToken(TreeLanParser.SYSTEM_FUNCTION, 0); }
		public TerminalNode ANY() { return getToken(TreeLanParser.ANY, 0); }
		public TerminalNode NUMBER() { return getToken(TreeLanParser.NUMBER, 0); }
		public TerminalNode SPECIAL_NUMBER() { return getToken(TreeLanParser.SPECIAL_NUMBER, 0); }
		public ClassValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classValue; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).enterClassValue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).exitClassValue(this);
		}
	}

	public final ClassValueContext classValue() throws RecognitionException {
		ClassValueContext _localctx = new ClassValueContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_classValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ANY) | (1L << SPECIAL_NUMBER) | (1L << NUMERIC_CONSTANT) | (1L << CONSTANT) | (1L << VARIABLE) | (1L << CONSTANT_EXPRESSION) | (1L << VARIABLE_EXPRESSION) | (1L << NUMERIC_CONSTANT_EXPRESSION) | (1L << SYSTEM_FUNCTION) | (1L << NUMBER))) != 0)) ) {
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

	public static class OperationClassContext extends ParserRuleContext {
		public TerminalNode ADDITION() { return getToken(TreeLanParser.ADDITION, 0); }
		public TerminalNode SUBTRACTION() { return getToken(TreeLanParser.SUBTRACTION, 0); }
		public TerminalNode MULTIPLICATION() { return getToken(TreeLanParser.MULTIPLICATION, 0); }
		public TerminalNode DIVISION() { return getToken(TreeLanParser.DIVISION, 0); }
		public TerminalNode EXPONENTIAL() { return getToken(TreeLanParser.EXPONENTIAL, 0); }
		public TerminalNode LOGARITHM() { return getToken(TreeLanParser.LOGARITHM, 0); }
		public TerminalNode SIN() { return getToken(TreeLanParser.SIN, 0); }
		public TerminalNode COS() { return getToken(TreeLanParser.COS, 0); }
		public TerminalNode TAN() { return getToken(TreeLanParser.TAN, 0); }
		public TerminalNode TRANSFORMATION() { return getToken(TreeLanParser.TRANSFORMATION, 0); }
		public TerminalNode EQUATION() { return getToken(TreeLanParser.EQUATION, 0); }
		public TerminalNode LEAF() { return getToken(TreeLanParser.LEAF, 0); }
		public TerminalNode ANY() { return getToken(TreeLanParser.ANY, 0); }
		public OperationClassContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operationClass; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).enterOperationClass(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TreeLanListener ) ((TreeLanListener)listener).exitOperationClass(this);
		}
	}

	public final OperationClassContext operationClass() throws RecognitionException {
		OperationClassContext _localctx = new OperationClassContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_operationClass);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(48);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ADDITION) | (1L << SUBTRACTION) | (1L << MULTIPLICATION) | (1L << DIVISION) | (1L << EXPONENTIAL) | (1L << LOGARITHM) | (1L << SIN) | (1L << COS) | (1L << TAN) | (1L << TRANSFORMATION) | (1L << EQUATION) | (1L << LEAF) | (1L << ANY))) != 0)) ) {
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

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\37\65\4\2\t\2\4\3"+
		"\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\3\2\3\2\3\2\3\2\3\2\3\2\3"+
		"\2\3\2\3\2\5\2\32\n\2\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4"+
		"\'\n\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\b\3\b\3\b\2\2\t\2\4\6"+
		"\b\n\f\16\2\4\4\2\23\33\36\36\3\2\7\23\2/\2\31\3\2\2\2\4\33\3\2\2\2\6"+
		"&\3\2\2\2\b(\3\2\2\2\n,\3\2\2\2\f\60\3\2\2\2\16\62\3\2\2\2\20\21\5\6\4"+
		"\2\21\22\7\35\2\2\22\23\5\6\4\2\23\32\3\2\2\2\24\25\5\6\4\2\25\26\7\35"+
		"\2\2\26\27\7\3\2\2\27\30\5\4\3\2\30\32\3\2\2\2\31\20\3\2\2\2\31\24\3\2"+
		"\2\2\32\3\3\2\2\2\33\34\7\34\2\2\34\5\3\2\2\2\35\36\7\4\2\2\36\37\5\6"+
		"\4\2\37 \7\5\2\2 !\5\b\5\2!\"\7\4\2\2\"#\5\6\4\2#$\7\5\2\2$\'\3\2\2\2"+
		"%\'\5\n\6\2&\35\3\2\2\2&%\3\2\2\2\'\7\3\2\2\2()\5\16\b\2)*\7\6\2\2*+\7"+
		"\34\2\2+\t\3\2\2\2,-\5\f\7\2-.\7\6\2\2./\7\34\2\2/\13\3\2\2\2\60\61\t"+
		"\2\2\2\61\r\3\2\2\2\62\63\t\3\2\2\63\17\3\2\2\2\4\31&";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}