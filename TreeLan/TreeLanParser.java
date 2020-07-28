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
		T__0=1, T__1=2, T__2=3, ADDITION=4, SUBSTRACTION=5, MULTIPLICATION=6, 
		DIVISION=7, EXPONENTIAL=8, LOGARITHM=9, SIN=10, COS=11, TAN=12, TRANSFORMATION=13, 
		EQUATION=14, LEAF=15, ANY=16, NUMERIC_CONSTANT=17, CONSTANT=18, VARIABLE=19, 
		CONSTANT_EXPRESSION=20, VARIABLE_EXPRESSION=21, NUMERIC_CONSTANT_EXPRESSION=22, 
		SYSTEM_FUNCTION=23, ID=24, ASSIGN=25, NUMBER=26, WS=27;
	public static final int
		RULE_treeRule = 0, RULE_treeDescriptor = 1, RULE_rootNode = 2, RULE_leafNode = 3, 
		RULE_classValue = 4, RULE_operationClass = 5;
	public static final String[] ruleNames = {
		"treeRule", "treeDescriptor", "rootNode", "leafNode", "classValue", "operationClass"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'('", "')'", "'_'", "'ADDITION'", "'SUBSTRACTION'", "'MULTIPLICATION'", 
		"'DIVISION'", "'EXPONENTIAL'", "'LOGARITHM'", "'SIN'", "'COS'", "'TAN'", 
		"'TRANSFORMATION'", "'EQUATION'", "'LEAF'", "'ANY'", "'NUMERIC_CONSTANT'", 
		"'CONSTANT'", "'VARIABLE'", "'CONSTANT_EXPRESSION'", "'VARIABLE_EXPRESSION'", 
		"'NUMERIC_CONSTANT_EXPRESSION'", "'SYSTEM_FUNCTION'", null, "'=>'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, "ADDITION", "SUBSTRACTION", "MULTIPLICATION", 
		"DIVISION", "EXPONENTIAL", "LOGARITHM", "SIN", "COS", "TAN", "TRANSFORMATION", 
		"EQUATION", "LEAF", "ANY", "NUMERIC_CONSTANT", "CONSTANT", "VARIABLE", 
		"CONSTANT_EXPRESSION", "VARIABLE_EXPRESSION", "NUMERIC_CONSTANT_EXPRESSION", 
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
			enterOuterAlt(_localctx, 1);
			{
			setState(12);
			treeDescriptor();
			setState(13);
			match(ASSIGN);
			setState(14);
			treeDescriptor();
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
		enterRule(_localctx, 2, RULE_treeDescriptor);
		try {
			setState(25);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(16);
				match(T__0);
				setState(17);
				treeDescriptor();
				setState(18);
				match(T__1);
				setState(19);
				rootNode();
				setState(20);
				match(T__0);
				setState(21);
				treeDescriptor();
				setState(22);
				match(T__1);
				}
				break;
			case ANY:
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
				setState(24);
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
		enterRule(_localctx, 4, RULE_rootNode);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(27);
			operationClass();
			setState(28);
			match(T__2);
			setState(29);
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
		enterRule(_localctx, 6, RULE_leafNode);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(31);
			classValue();
			setState(32);
			match(T__2);
			setState(33);
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
		enterRule(_localctx, 8, RULE_classValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(35);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ANY) | (1L << NUMERIC_CONSTANT) | (1L << CONSTANT) | (1L << VARIABLE) | (1L << CONSTANT_EXPRESSION) | (1L << VARIABLE_EXPRESSION) | (1L << NUMERIC_CONSTANT_EXPRESSION) | (1L << SYSTEM_FUNCTION) | (1L << NUMBER))) != 0)) ) {
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
		public TerminalNode SUBSTRACTION() { return getToken(TreeLanParser.SUBSTRACTION, 0); }
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
		enterRule(_localctx, 10, RULE_operationClass);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(37);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ADDITION) | (1L << SUBSTRACTION) | (1L << MULTIPLICATION) | (1L << DIVISION) | (1L << EXPONENTIAL) | (1L << LOGARITHM) | (1L << SIN) | (1L << COS) | (1L << TAN) | (1L << TRANSFORMATION) | (1L << EQUATION) | (1L << LEAF) | (1L << ANY))) != 0)) ) {
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\35*\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\5\3\34\n\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3"+
		"\7\3\7\3\7\2\2\b\2\4\6\b\n\f\2\4\4\2\22\31\34\34\3\2\6\22\2$\2\16\3\2"+
		"\2\2\4\33\3\2\2\2\6\35\3\2\2\2\b!\3\2\2\2\n%\3\2\2\2\f\'\3\2\2\2\16\17"+
		"\5\4\3\2\17\20\7\33\2\2\20\21\5\4\3\2\21\3\3\2\2\2\22\23\7\3\2\2\23\24"+
		"\5\4\3\2\24\25\7\4\2\2\25\26\5\6\4\2\26\27\7\3\2\2\27\30\5\4\3\2\30\31"+
		"\7\4\2\2\31\34\3\2\2\2\32\34\5\b\5\2\33\22\3\2\2\2\33\32\3\2\2\2\34\5"+
		"\3\2\2\2\35\36\5\f\7\2\36\37\7\5\2\2\37 \7\32\2\2 \7\3\2\2\2!\"\5\n\6"+
		"\2\"#\7\5\2\2#$\7\32\2\2$\t\3\2\2\2%&\t\2\2\2&\13\3\2\2\2\'(\t\3\2\2("+
		"\r\3\2\2\2\3\33";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}