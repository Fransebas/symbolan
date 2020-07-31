// Generated from TreeLan.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class TreeLanLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "T__3", "ADDITION", "SUBTRACTION", "MULTIPLICATION", 
		"DIVISION", "EXPONENTIAL", "LOGARITHM", "SIN", "COS", "TAN", "TRANSFORMATION", 
		"EQUATION", "LEAF", "ANY", "SPECIAL_NUMBER", "NUMERIC_CONSTANT", "CONSTANT", 
		"VARIABLE", "CONSTANT_EXPRESSION", "VARIABLE_EXPRESSION", "NUMERIC_CONSTANT_EXPRESSION", 
		"SYSTEM_FUNCTION", "ID", "ASSIGN", "NUMBER", "SP_NUMBER", "INT", "EXP", 
		"WS"
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


	public TreeLanLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "TreeLan.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\37\u015b\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\3\2\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\r\3\r\3"+
		"\r\3\r\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\23\3\23\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\33\6\33\u0137\n\33\r\33\16\33\u0138\3\34\3"+
		"\34\3\34\3\35\3\35\3\36\3\36\3\36\5\36\u0143\n\36\3\37\3\37\3\37\7\37"+
		"\u0148\n\37\f\37\16\37\u014b\13\37\5\37\u014d\n\37\3 \3 \5 \u0151\n \3"+
		" \3 \3!\6!\u0156\n!\r!\16!\u0157\3!\3!\2\2\"\3\3\5\4\7\5\t\6\13\7\r\b"+
		"\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26"+
		"+\27-\30/\31\61\32\63\33\65\34\67\359\36;\2=\2?\2A\37\3\2\b\4\2C\\c|\3"+
		"\2\63;\3\2\62;\4\2GGgg\4\2--//\5\2\13\f\17\17\"\"\2\u015d\2\3\3\2\2\2"+
		"\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2"+
		"\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2"+
		"\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2"+
		"\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2"+
		"\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2A\3\2\2\2\3C\3\2"+
		"\2\2\5F\3\2\2\2\7H\3\2\2\2\tJ\3\2\2\2\13L\3\2\2\2\rU\3\2\2\2\17a\3\2\2"+
		"\2\21p\3\2\2\2\23y\3\2\2\2\25\u0085\3\2\2\2\27\u008f\3\2\2\2\31\u0093"+
		"\3\2\2\2\33\u0097\3\2\2\2\35\u009b\3\2\2\2\37\u00aa\3\2\2\2!\u00b3\3\2"+
		"\2\2#\u00b8\3\2\2\2%\u00bc\3\2\2\2\'\u00be\3\2\2\2)\u00cf\3\2\2\2+\u00d8"+
		"\3\2\2\2-\u00e1\3\2\2\2/\u00f5\3\2\2\2\61\u0109\3\2\2\2\63\u0125\3\2\2"+
		"\2\65\u0136\3\2\2\2\67\u013a\3\2\2\29\u013d\3\2\2\2;\u0142\3\2\2\2=\u014c"+
		"\3\2\2\2?\u014e\3\2\2\2A\u0155\3\2\2\2CD\7H\2\2DE\7a\2\2E\4\3\2\2\2FG"+
		"\7*\2\2G\6\3\2\2\2HI\7+\2\2I\b\3\2\2\2JK\7a\2\2K\n\3\2\2\2LM\7C\2\2MN"+
		"\7F\2\2NO\7F\2\2OP\7K\2\2PQ\7V\2\2QR\7K\2\2RS\7Q\2\2ST\7P\2\2T\f\3\2\2"+
		"\2UV\7U\2\2VW\7W\2\2WX\7D\2\2XY\7V\2\2YZ\7T\2\2Z[\7C\2\2[\\\7E\2\2\\]"+
		"\7V\2\2]^\7K\2\2^_\7Q\2\2_`\7P\2\2`\16\3\2\2\2ab\7O\2\2bc\7W\2\2cd\7N"+
		"\2\2de\7V\2\2ef\7K\2\2fg\7R\2\2gh\7N\2\2hi\7K\2\2ij\7E\2\2jk\7C\2\2kl"+
		"\7V\2\2lm\7K\2\2mn\7Q\2\2no\7P\2\2o\20\3\2\2\2pq\7F\2\2qr\7K\2\2rs\7X"+
		"\2\2st\7K\2\2tu\7U\2\2uv\7K\2\2vw\7Q\2\2wx\7P\2\2x\22\3\2\2\2yz\7G\2\2"+
		"z{\7Z\2\2{|\7R\2\2|}\7Q\2\2}~\7P\2\2~\177\7G\2\2\177\u0080\7P\2\2\u0080"+
		"\u0081\7V\2\2\u0081\u0082\7K\2\2\u0082\u0083\7C\2\2\u0083\u0084\7N\2\2"+
		"\u0084\24\3\2\2\2\u0085\u0086\7N\2\2\u0086\u0087\7Q\2\2\u0087\u0088\7"+
		"I\2\2\u0088\u0089\7C\2\2\u0089\u008a\7T\2\2\u008a\u008b\7K\2\2\u008b\u008c"+
		"\7V\2\2\u008c\u008d\7J\2\2\u008d\u008e\7O\2\2\u008e\26\3\2\2\2\u008f\u0090"+
		"\7U\2\2\u0090\u0091\7K\2\2\u0091\u0092\7P\2\2\u0092\30\3\2\2\2\u0093\u0094"+
		"\7E\2\2\u0094\u0095\7Q\2\2\u0095\u0096\7U\2\2\u0096\32\3\2\2\2\u0097\u0098"+
		"\7V\2\2\u0098\u0099\7C\2\2\u0099\u009a\7P\2\2\u009a\34\3\2\2\2\u009b\u009c"+
		"\7V\2\2\u009c\u009d\7T\2\2\u009d\u009e\7C\2\2\u009e\u009f\7P\2\2\u009f"+
		"\u00a0\7U\2\2\u00a0\u00a1\7H\2\2\u00a1\u00a2\7Q\2\2\u00a2\u00a3\7T\2\2"+
		"\u00a3\u00a4\7O\2\2\u00a4\u00a5\7C\2\2\u00a5\u00a6\7V\2\2\u00a6\u00a7"+
		"\7K\2\2\u00a7\u00a8\7Q\2\2\u00a8\u00a9\7P\2\2\u00a9\36\3\2\2\2\u00aa\u00ab"+
		"\7G\2\2\u00ab\u00ac\7S\2\2\u00ac\u00ad\7W\2\2\u00ad\u00ae\7C\2\2\u00ae"+
		"\u00af\7V\2\2\u00af\u00b0\7K\2\2\u00b0\u00b1\7Q\2\2\u00b1\u00b2\7P\2\2"+
		"\u00b2 \3\2\2\2\u00b3\u00b4\7N\2\2\u00b4\u00b5\7G\2\2\u00b5\u00b6\7C\2"+
		"\2\u00b6\u00b7\7H\2\2\u00b7\"\3\2\2\2\u00b8\u00b9\7C\2\2\u00b9\u00ba\7"+
		"P\2\2\u00ba\u00bb\7[\2\2\u00bb$\3\2\2\2\u00bc\u00bd\5;\36\2\u00bd&\3\2"+
		"\2\2\u00be\u00bf\7P\2\2\u00bf\u00c0\7W\2\2\u00c0\u00c1\7O\2\2\u00c1\u00c2"+
		"\7G\2\2\u00c2\u00c3\7T\2\2\u00c3\u00c4\7K\2\2\u00c4\u00c5\7E\2\2\u00c5"+
		"\u00c6\7a\2\2\u00c6\u00c7\7E\2\2\u00c7\u00c8\7Q\2\2\u00c8\u00c9\7P\2\2"+
		"\u00c9\u00ca\7U\2\2\u00ca\u00cb\7V\2\2\u00cb\u00cc\7C\2\2\u00cc\u00cd"+
		"\7P\2\2\u00cd\u00ce\7V\2\2\u00ce(\3\2\2\2\u00cf\u00d0\7E\2\2\u00d0\u00d1"+
		"\7Q\2\2\u00d1\u00d2\7P\2\2\u00d2\u00d3\7U\2\2\u00d3\u00d4\7V\2\2\u00d4"+
		"\u00d5\7C\2\2\u00d5\u00d6\7P\2\2\u00d6\u00d7\7V\2\2\u00d7*\3\2\2\2\u00d8"+
		"\u00d9\7X\2\2\u00d9\u00da\7C\2\2\u00da\u00db\7T\2\2\u00db\u00dc\7K\2\2"+
		"\u00dc\u00dd\7C\2\2\u00dd\u00de\7D\2\2\u00de\u00df\7N\2\2\u00df\u00e0"+
		"\7G\2\2\u00e0,\3\2\2\2\u00e1\u00e2\7E\2\2\u00e2\u00e3\7Q\2\2\u00e3\u00e4"+
		"\7P\2\2\u00e4\u00e5\7U\2\2\u00e5\u00e6\7V\2\2\u00e6\u00e7\7C\2\2\u00e7"+
		"\u00e8\7P\2\2\u00e8\u00e9\7V\2\2\u00e9\u00ea\7a\2\2\u00ea\u00eb\7G\2\2"+
		"\u00eb\u00ec\7Z\2\2\u00ec\u00ed\7R\2\2\u00ed\u00ee\7T\2\2\u00ee\u00ef"+
		"\7G\2\2\u00ef\u00f0\7U\2\2\u00f0\u00f1\7U\2\2\u00f1\u00f2\7K\2\2\u00f2"+
		"\u00f3\7Q\2\2\u00f3\u00f4\7P\2\2\u00f4.\3\2\2\2\u00f5\u00f6\7X\2\2\u00f6"+
		"\u00f7\7C\2\2\u00f7\u00f8\7T\2\2\u00f8\u00f9\7K\2\2\u00f9\u00fa\7C\2\2"+
		"\u00fa\u00fb\7D\2\2\u00fb\u00fc\7N\2\2\u00fc\u00fd\7G\2\2\u00fd\u00fe"+
		"\7a\2\2\u00fe\u00ff\7G\2\2\u00ff\u0100\7Z\2\2\u0100\u0101\7R\2\2\u0101"+
		"\u0102\7T\2\2\u0102\u0103\7G\2\2\u0103\u0104\7U\2\2\u0104\u0105\7U\2\2"+
		"\u0105\u0106\7K\2\2\u0106\u0107\7Q\2\2\u0107\u0108\7P\2\2\u0108\60\3\2"+
		"\2\2\u0109\u010a\7P\2\2\u010a\u010b\7W\2\2\u010b\u010c\7O\2\2\u010c\u010d"+
		"\7G\2\2\u010d\u010e\7T\2\2\u010e\u010f\7K\2\2\u010f\u0110\7E\2\2\u0110"+
		"\u0111\7a\2\2\u0111\u0112\7E\2\2\u0112\u0113\7Q\2\2\u0113\u0114\7P\2\2"+
		"\u0114\u0115\7U\2\2\u0115\u0116\7V\2\2\u0116\u0117\7C\2\2\u0117\u0118"+
		"\7P\2\2\u0118\u0119\7V\2\2\u0119\u011a\7a\2\2\u011a\u011b\7G\2\2\u011b"+
		"\u011c\7Z\2\2\u011c\u011d\7R\2\2\u011d\u011e\7T\2\2\u011e\u011f\7G\2\2"+
		"\u011f\u0120\7U\2\2\u0120\u0121\7U\2\2\u0121\u0122\7K\2\2\u0122\u0123"+
		"\7Q\2\2\u0123\u0124\7P\2\2\u0124\62\3\2\2\2\u0125\u0126\7U\2\2\u0126\u0127"+
		"\7[\2\2\u0127\u0128\7U\2\2\u0128\u0129\7V\2\2\u0129\u012a\7G\2\2\u012a"+
		"\u012b\7O\2\2\u012b\u012c\7a\2\2\u012c\u012d\7H\2\2\u012d\u012e\7W\2\2"+
		"\u012e\u012f\7P\2\2\u012f\u0130\7E\2\2\u0130\u0131\7V\2\2\u0131\u0132"+
		"\7K\2\2\u0132\u0133\7Q\2\2\u0133\u0134\7P\2\2\u0134\64\3\2\2\2\u0135\u0137"+
		"\t\2\2\2\u0136\u0135\3\2\2\2\u0137\u0138\3\2\2\2\u0138\u0136\3\2\2\2\u0138"+
		"\u0139\3\2\2\2\u0139\66\3\2\2\2\u013a\u013b\7?\2\2\u013b\u013c\7@\2\2"+
		"\u013c8\3\2\2\2\u013d\u013e\5=\37\2\u013e:\3\2\2\2\u013f\u0143\7g\2\2"+
		"\u0140\u0141\7r\2\2\u0141\u0143\7k\2\2\u0142\u013f\3\2\2\2\u0142\u0140"+
		"\3\2\2\2\u0143<\3\2\2\2\u0144\u014d\7\62\2\2\u0145\u0149\t\3\2\2\u0146"+
		"\u0148\t\4\2\2\u0147\u0146\3\2\2\2\u0148\u014b\3\2\2\2\u0149\u0147\3\2"+
		"\2\2\u0149\u014a\3\2\2\2\u014a\u014d\3\2\2\2\u014b\u0149\3\2\2\2\u014c"+
		"\u0144\3\2\2\2\u014c\u0145\3\2\2\2\u014d>\3\2\2\2\u014e\u0150\t\5\2\2"+
		"\u014f\u0151\t\6\2\2\u0150\u014f\3\2\2\2\u0150\u0151\3\2\2\2\u0151\u0152"+
		"\3\2\2\2\u0152\u0153\5=\37\2\u0153@\3\2\2\2\u0154\u0156\t\7\2\2\u0155"+
		"\u0154\3\2\2\2\u0156\u0157\3\2\2\2\u0157\u0155\3\2\2\2\u0157\u0158\3\2"+
		"\2\2\u0158\u0159\3\2\2\2\u0159\u015a\b!\2\2\u015aB\3\2\2\2\t\2\u0138\u0142"+
		"\u0149\u014c\u0150\u0157\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}