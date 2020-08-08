// Generated from Symbolan.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SymbolanLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, L_PAREN=9, 
		R_PAREN=10, L_CURLY=11, R_CURLY=12, L_BRACKET=13, R_BRACKET=14, ASSIGN=15, 
		COMMA=16, FACTORIAL=17, MIN=18, MAX=19, SIN=20, COS=21, TAN=22, LOG=23, 
		SQRT=24, DERIVATIVE_OPERATOR=25, PI=26, DERIVATIVE=27, VAR=28, CONST=29, 
		EULER=30, INFINITE=31, UNDEFINED=32, DERIVATE_RULE=33, NUMERIC_RULE=34, 
		CONST_RULE=35, VAR_RULE=36, EXPR_RULE=37, IMAGINARY=38, DECIMAL_LIT=39, 
		OCTAL_LIT=40, HEX_LIT=41, FLOAT_LIT=42, IMAGINARY_LIT=43, IDENTIFIER=44, 
		NEWLINE=45, WS=46, LINE_COMMENT=47;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "L_PAREN", 
		"R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", 
		"FACTORIAL", "MIN", "MAX", "SIN", "COS", "TAN", "LOG", "SQRT", "DERIVATIVE_OPERATOR", 
		"PI", "DERIVATIVE", "VAR", "CONST", "EULER", "INFINITE", "UNDEFINED", 
		"DERIVATE_RULE", "NUMERIC_RULE", "CONST_RULE", "VAR_RULE", "EXPR_RULE", 
		"IMAGINARY", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "IMAGINARY_LIT", 
		"IDENTIFIER", "DECIMALS", "OCTAL_DIGIT", "HEX_DIGIT", "EXPONENT", "NEWLINE", 
		"WS", "LINE_COMMENT"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'()'", "'=>'", "'^'", "'**'", "'*'", "'/'", "'+'", "'-'", "'('", 
		"')'", "'{'", "'}'", "'['", "']'", "'='", "','", "'!'", "'min'", "'max'", 
		"'sin'", "'cos'", "'tan'", "'log'", "'sqrt'", "'D'", "'pi'", null, null, 
		null, null, null, "'undef'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, null, null, "L_PAREN", "R_PAREN", 
		"L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "FACTORIAL", 
		"MIN", "MAX", "SIN", "COS", "TAN", "LOG", "SQRT", "DERIVATIVE_OPERATOR", 
		"PI", "DERIVATIVE", "VAR", "CONST", "EULER", "INFINITE", "UNDEFINED", 
		"DERIVATE_RULE", "NUMERIC_RULE", "CONST_RULE", "VAR_RULE", "EXPR_RULE", 
		"IMAGINARY", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "IMAGINARY_LIT", 
		"IDENTIFIER", "NEWLINE", "WS", "LINE_COMMENT"
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


	public SymbolanLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Symbolan.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\61\u017e\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\3\2\3\2\3\2\3\3\3\3\3\3\3\4\3\4\3\5\3\5\3\5\3\6\3\6\3\7\3\7\3\b"+
		"\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20"+
		"\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25"+
		"\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\30\3\30\3\30"+
		"\3\30\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\33\3\33\3\33\3\34\3\34\3\34"+
		"\3\34\6\34\u00b5\n\34\r\34\16\34\u00b6\5\34\u00b9\n\34\3\35\3\35\3\35"+
		"\6\35\u00be\n\35\r\35\16\35\u00bf\5\35\u00c2\n\35\3\36\3\36\3\36\6\36"+
		"\u00c7\n\36\r\36\16\36\u00c8\5\36\u00cb\n\36\3\37\3\37\3 \5 \u00d0\n "+
		"\3 \3 \3 \3 \3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\6\"\u00e0\n\"\r\"\16\""+
		"\u00e1\5\"\u00e4\n\"\3#\5#\u00e7\n#\3#\3#\3#\6#\u00ec\n#\r#\16#\u00ed"+
		"\5#\u00f0\n#\3$\5$\u00f3\n$\3$\3$\3$\6$\u00f8\n$\r$\16$\u00f9\5$\u00fc"+
		"\n$\3%\5%\u00ff\n%\3%\3%\3%\6%\u0104\n%\r%\16%\u0105\5%\u0108\n%\3&\5"+
		"&\u010b\n&\3&\3&\3&\6&\u0110\n&\r&\16&\u0111\5&\u0114\n&\3\'\3\'\5\'\u0118"+
		"\n\'\3\'\3\'\3(\3(\7(\u011e\n(\f(\16(\u0121\13(\3)\3)\7)\u0125\n)\f)\16"+
		")\u0128\13)\3*\3*\3*\6*\u012d\n*\r*\16*\u012e\3+\3+\3+\5+\u0134\n+\3+"+
		"\5+\u0137\n+\3+\5+\u013a\n+\3+\3+\3+\5+\u013f\n+\5+\u0141\n+\3,\3,\5,"+
		"\u0145\n,\3,\3,\3-\3-\7-\u014b\n-\f-\16-\u014e\13-\3.\5.\u0151\n.\3.\6"+
		".\u0154\n.\r.\16.\u0155\3/\5/\u0159\n/\3/\3/\3\60\5\60\u015e\n\60\3\60"+
		"\3\60\3\61\3\61\5\61\u0164\n\61\3\61\3\61\3\62\5\62\u0169\n\62\3\62\3"+
		"\62\3\63\6\63\u016e\n\63\r\63\16\63\u016f\3\63\3\63\3\64\3\64\3\64\3\64"+
		"\7\64\u0178\n\64\f\64\16\64\u017b\13\64\3\64\3\64\2\2\65\3\3\5\4\7\5\t"+
		"\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23"+
		"%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G"+
		"%I&K\'M(O)Q*S+U,W-Y.[\2]\2_\2a\2c/e\60g\61\3\2\24\3\2z|\3\2\62;\3\2ce"+
		"\3\2gg\4\2--//\3\2XX\3\2PP\3\2EE\3\2HH\3\2\63;\4\2ZZzz\3\2C\\\6\2\62;"+
		"C\\aac|\3\2\629\5\2\62;CHch\4\2GGgg\4\2\13\13\"\"\4\2\f\f\17\17\2\u01a1"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2"+
		"\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3"+
		"\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2"+
		"\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2"+
		"U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\3i\3"+
		"\2\2\2\5l\3\2\2\2\7o\3\2\2\2\tq\3\2\2\2\13t\3\2\2\2\rv\3\2\2\2\17x\3\2"+
		"\2\2\21z\3\2\2\2\23|\3\2\2\2\25~\3\2\2\2\27\u0080\3\2\2\2\31\u0082\3\2"+
		"\2\2\33\u0084\3\2\2\2\35\u0086\3\2\2\2\37\u0088\3\2\2\2!\u008a\3\2\2\2"+
		"#\u008c\3\2\2\2%\u008e\3\2\2\2\'\u0092\3\2\2\2)\u0096\3\2\2\2+\u009a\3"+
		"\2\2\2-\u009e\3\2\2\2/\u00a2\3\2\2\2\61\u00a6\3\2\2\2\63\u00ab\3\2\2\2"+
		"\65\u00ad\3\2\2\2\67\u00b0\3\2\2\29\u00ba\3\2\2\2;\u00c3\3\2\2\2=\u00cc"+
		"\3\2\2\2?\u00cf\3\2\2\2A\u00d5\3\2\2\2C\u00db\3\2\2\2E\u00e6\3\2\2\2G"+
		"\u00f2\3\2\2\2I\u00fe\3\2\2\2K\u010a\3\2\2\2M\u0117\3\2\2\2O\u011b\3\2"+
		"\2\2Q\u0122\3\2\2\2S\u0129\3\2\2\2U\u0140\3\2\2\2W\u0144\3\2\2\2Y\u0148"+
		"\3\2\2\2[\u0150\3\2\2\2]\u0158\3\2\2\2_\u015d\3\2\2\2a\u0161\3\2\2\2c"+
		"\u0168\3\2\2\2e\u016d\3\2\2\2g\u0173\3\2\2\2ij\7*\2\2jk\7+\2\2k\4\3\2"+
		"\2\2lm\7?\2\2mn\7@\2\2n\6\3\2\2\2op\7`\2\2p\b\3\2\2\2qr\7,\2\2rs\7,\2"+
		"\2s\n\3\2\2\2tu\7,\2\2u\f\3\2\2\2vw\7\61\2\2w\16\3\2\2\2xy\7-\2\2y\20"+
		"\3\2\2\2z{\7/\2\2{\22\3\2\2\2|}\7*\2\2}\24\3\2\2\2~\177\7+\2\2\177\26"+
		"\3\2\2\2\u0080\u0081\7}\2\2\u0081\30\3\2\2\2\u0082\u0083\7\177\2\2\u0083"+
		"\32\3\2\2\2\u0084\u0085\7]\2\2\u0085\34\3\2\2\2\u0086\u0087\7_\2\2\u0087"+
		"\36\3\2\2\2\u0088\u0089\7?\2\2\u0089 \3\2\2\2\u008a\u008b\7.\2\2\u008b"+
		"\"\3\2\2\2\u008c\u008d\7#\2\2\u008d$\3\2\2\2\u008e\u008f\7o\2\2\u008f"+
		"\u0090\7k\2\2\u0090\u0091\7p\2\2\u0091&\3\2\2\2\u0092\u0093\7o\2\2\u0093"+
		"\u0094\7c\2\2\u0094\u0095\7z\2\2\u0095(\3\2\2\2\u0096\u0097\7u\2\2\u0097"+
		"\u0098\7k\2\2\u0098\u0099\7p\2\2\u0099*\3\2\2\2\u009a\u009b\7e\2\2\u009b"+
		"\u009c\7q\2\2\u009c\u009d\7u\2\2\u009d,\3\2\2\2\u009e\u009f\7v\2\2\u009f"+
		"\u00a0\7c\2\2\u00a0\u00a1\7p\2\2\u00a1.\3\2\2\2\u00a2\u00a3\7n\2\2\u00a3"+
		"\u00a4\7q\2\2\u00a4\u00a5\7i\2\2\u00a5\60\3\2\2\2\u00a6\u00a7\7u\2\2\u00a7"+
		"\u00a8\7s\2\2\u00a8\u00a9\7t\2\2\u00a9\u00aa\7v\2\2\u00aa\62\3\2\2\2\u00ab"+
		"\u00ac\7F\2\2\u00ac\64\3\2\2\2\u00ad\u00ae\7r\2\2\u00ae\u00af\7k\2\2\u00af"+
		"\66\3\2\2\2\u00b0\u00b1\7f\2\2\u00b1\u00b8\t\2\2\2\u00b2\u00b4\7a\2\2"+
		"\u00b3\u00b5\t\3\2\2\u00b4\u00b3\3\2\2\2\u00b5\u00b6\3\2\2\2\u00b6\u00b4"+
		"\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7\u00b9\3\2\2\2\u00b8\u00b2\3\2\2\2\u00b8"+
		"\u00b9\3\2\2\2\u00b98\3\2\2\2\u00ba\u00c1\t\2\2\2\u00bb\u00bd\7a\2\2\u00bc"+
		"\u00be\t\3\2\2\u00bd\u00bc\3\2\2\2\u00be\u00bf\3\2\2\2\u00bf\u00bd\3\2"+
		"\2\2\u00bf\u00c0\3\2\2\2\u00c0\u00c2\3\2\2\2\u00c1\u00bb\3\2\2\2\u00c1"+
		"\u00c2\3\2\2\2\u00c2:\3\2\2\2\u00c3\u00ca\t\4\2\2\u00c4\u00c6\7a\2\2\u00c5"+
		"\u00c7\t\3\2\2\u00c6\u00c5\3\2\2\2\u00c7\u00c8\3\2\2\2\u00c8\u00c6\3\2"+
		"\2\2\u00c8\u00c9\3\2\2\2\u00c9\u00cb\3\2\2\2\u00ca\u00c4\3\2\2\2\u00ca"+
		"\u00cb\3\2\2\2\u00cb<\3\2\2\2\u00cc\u00cd\t\5\2\2\u00cd>\3\2\2\2\u00ce"+
		"\u00d0\t\6\2\2\u00cf\u00ce\3\2\2\2\u00cf\u00d0\3\2\2\2\u00d0\u00d1\3\2"+
		"\2\2\u00d1\u00d2\7k\2\2\u00d2\u00d3\7p\2\2\u00d3\u00d4\7h\2\2\u00d4@\3"+
		"\2\2\2\u00d5\u00d6\7w\2\2\u00d6\u00d7\7p\2\2\u00d7\u00d8\7f\2\2\u00d8"+
		"\u00d9\7g\2\2\u00d9\u00da\7h\2\2\u00daB\3\2\2\2\u00db\u00dc\7f\2\2\u00dc"+
		"\u00e3\t\7\2\2\u00dd\u00df\7a\2\2\u00de\u00e0\t\3\2\2\u00df\u00de\3\2"+
		"\2\2\u00e0\u00e1\3\2\2\2\u00e1\u00df\3\2\2\2\u00e1\u00e2\3\2\2\2\u00e2"+
		"\u00e4\3\2\2\2\u00e3\u00dd\3\2\2\2\u00e3\u00e4\3\2\2\2\u00e4D\3\2\2\2"+
		"\u00e5\u00e7\t\6\2\2\u00e6\u00e5\3\2\2\2\u00e6\u00e7\3\2\2\2\u00e7\u00e8"+
		"\3\2\2\2\u00e8\u00ef\t\b\2\2\u00e9\u00eb\7a\2\2\u00ea\u00ec\t\3\2\2\u00eb"+
		"\u00ea\3\2\2\2\u00ec\u00ed\3\2\2\2\u00ed\u00eb\3\2\2\2\u00ed\u00ee\3\2"+
		"\2\2\u00ee\u00f0\3\2\2\2\u00ef\u00e9\3\2\2\2\u00ef\u00f0\3\2\2\2\u00f0"+
		"F\3\2\2\2\u00f1\u00f3\t\6\2\2\u00f2\u00f1\3\2\2\2\u00f2\u00f3\3\2\2\2"+
		"\u00f3\u00f4\3\2\2\2\u00f4\u00fb\t\t\2\2\u00f5\u00f7\7a\2\2\u00f6\u00f8"+
		"\t\3\2\2\u00f7\u00f6\3\2\2\2\u00f8\u00f9\3\2\2\2\u00f9\u00f7\3\2\2\2\u00f9"+
		"\u00fa\3\2\2\2\u00fa\u00fc\3\2\2\2\u00fb\u00f5\3\2\2\2\u00fb\u00fc\3\2"+
		"\2\2\u00fcH\3\2\2\2\u00fd\u00ff\t\6\2\2\u00fe\u00fd\3\2\2\2\u00fe\u00ff"+
		"\3\2\2\2\u00ff\u0100\3\2\2\2\u0100\u0107\t\7\2\2\u0101\u0103\7a\2\2\u0102"+
		"\u0104\t\3\2\2\u0103\u0102\3\2\2\2\u0104\u0105\3\2\2\2\u0105\u0103\3\2"+
		"\2\2\u0105\u0106\3\2\2\2\u0106\u0108\3\2\2\2\u0107\u0101\3\2\2\2\u0107"+
		"\u0108\3\2\2\2\u0108J\3\2\2\2\u0109\u010b\t\6\2\2\u010a\u0109\3\2\2\2"+
		"\u010a\u010b\3\2\2\2\u010b\u010c\3\2\2\2\u010c\u0113\t\n\2\2\u010d\u010f"+
		"\7a\2\2\u010e\u0110\t\3\2\2\u010f\u010e\3\2\2\2\u0110\u0111\3\2\2\2\u0111"+
		"\u010f\3\2\2\2\u0111\u0112\3\2\2\2\u0112\u0114\3\2\2\2\u0113\u010d\3\2"+
		"\2\2\u0113\u0114\3\2\2\2\u0114L\3\2\2\2\u0115\u0118\5[.\2\u0116\u0118"+
		"\5U+\2\u0117\u0115\3\2\2\2\u0117\u0116\3\2\2\2\u0118\u0119\3\2\2\2\u0119"+
		"\u011a\7k\2\2\u011aN\3\2\2\2\u011b\u011f\t\13\2\2\u011c\u011e\t\3\2\2"+
		"\u011d\u011c\3\2\2\2\u011e\u0121\3\2\2\2\u011f\u011d\3\2\2\2\u011f\u0120"+
		"\3\2\2\2\u0120P\3\2\2\2\u0121\u011f\3\2\2\2\u0122\u0126\7\62\2\2\u0123"+
		"\u0125\5]/\2\u0124\u0123\3\2\2\2\u0125\u0128\3\2\2\2\u0126\u0124\3\2\2"+
		"\2\u0126\u0127\3\2\2\2\u0127R\3\2\2\2\u0128\u0126\3\2\2\2\u0129\u012a"+
		"\7\62\2\2\u012a\u012c\t\f\2\2\u012b\u012d\5_\60\2\u012c\u012b\3\2\2\2"+
		"\u012d\u012e\3\2\2\2\u012e\u012c\3\2\2\2\u012e\u012f\3\2\2\2\u012fT\3"+
		"\2\2\2\u0130\u0139\5[.\2\u0131\u0133\7\60\2\2\u0132\u0134\5[.\2\u0133"+
		"\u0132\3\2\2\2\u0133\u0134\3\2\2\2\u0134\u0136\3\2\2\2\u0135\u0137\5a"+
		"\61\2\u0136\u0135\3\2\2\2\u0136\u0137\3\2\2\2\u0137\u013a\3\2\2\2\u0138"+
		"\u013a\5a\61\2\u0139\u0131\3\2\2\2\u0139\u0138\3\2\2\2\u013a\u0141\3\2"+
		"\2\2\u013b\u013c\7\60\2\2\u013c\u013e\5[.\2\u013d\u013f\5a\61\2\u013e"+
		"\u013d\3\2\2\2\u013e\u013f\3\2\2\2\u013f\u0141\3\2\2\2\u0140\u0130\3\2"+
		"\2\2\u0140\u013b\3\2\2\2\u0141V\3\2\2\2\u0142\u0145\5[.\2\u0143\u0145"+
		"\5U+\2\u0144\u0142\3\2\2\2\u0144\u0143\3\2\2\2\u0145\u0146\3\2\2\2\u0146"+
		"\u0147\7k\2\2\u0147X\3\2\2\2\u0148\u014c\t\r\2\2\u0149\u014b\t\16\2\2"+
		"\u014a\u0149\3\2\2\2\u014b\u014e\3\2\2\2\u014c\u014a\3\2\2\2\u014c\u014d"+
		"\3\2\2\2\u014dZ\3\2\2\2\u014e\u014c\3\2\2\2\u014f\u0151\t\6\2\2\u0150"+
		"\u014f\3\2\2\2\u0150\u0151\3\2\2\2\u0151\u0153\3\2\2\2\u0152\u0154\t\3"+
		"\2\2\u0153\u0152\3\2\2\2\u0154\u0155\3\2\2\2\u0155\u0153\3\2\2\2\u0155"+
		"\u0156\3\2\2\2\u0156\\\3\2\2\2\u0157\u0159\t\6\2\2\u0158\u0157\3\2\2\2"+
		"\u0158\u0159\3\2\2\2\u0159\u015a\3\2\2\2\u015a\u015b\t\17\2\2\u015b^\3"+
		"\2\2\2\u015c\u015e\t\6\2\2\u015d\u015c\3\2\2\2\u015d\u015e\3\2\2\2\u015e"+
		"\u015f\3\2\2\2\u015f\u0160\t\20\2\2\u0160`\3\2\2\2\u0161\u0163\t\21\2"+
		"\2\u0162\u0164\t\6\2\2\u0163\u0162\3\2\2\2\u0163\u0164\3\2\2\2\u0164\u0165"+
		"\3\2\2\2\u0165\u0166\5[.\2\u0166b\3\2\2\2\u0167\u0169\7\17\2\2\u0168\u0167"+
		"\3\2\2\2\u0168\u0169\3\2\2\2\u0169\u016a\3\2\2\2\u016a\u016b\7\f\2\2\u016b"+
		"d\3\2\2\2\u016c\u016e\t\22\2\2\u016d\u016c\3\2\2\2\u016e\u016f\3\2\2\2"+
		"\u016f\u016d\3\2\2\2\u016f\u0170\3\2\2\2\u0170\u0171\3\2\2\2\u0171\u0172"+
		"\b\63\2\2\u0172f\3\2\2\2\u0173\u0174\7\61\2\2\u0174\u0175\7\61\2\2\u0175"+
		"\u0179\3\2\2\2\u0176\u0178\n\23\2\2\u0177\u0176\3\2\2\2\u0178\u017b\3"+
		"\2\2\2\u0179\u0177\3\2\2\2\u0179\u017a\3\2\2\2\u017a\u017c\3\2\2\2\u017b"+
		"\u0179\3\2\2\2\u017c\u017d\b\64\3\2\u017dh\3\2\2\2+\2\u00b6\u00b8\u00bf"+
		"\u00c1\u00c8\u00ca\u00cf\u00e1\u00e3\u00e6\u00ed\u00ef\u00f2\u00f9\u00fb"+
		"\u00fe\u0105\u0107\u010a\u0111\u0113\u0117\u011f\u0126\u012e\u0133\u0136"+
		"\u0139\u013e\u0140\u0144\u014c\u0150\u0155\u0158\u015d\u0163\u0168\u016f"+
		"\u0179\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}