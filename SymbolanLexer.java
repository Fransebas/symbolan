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
		SQRT=24, PI=25, VAR=26, CONST=27, EULER=28, INFINITE=29, UNDEFINED=30, 
		NUMERIC_RULE=31, CONST_RULE=32, VAR_RULE=33, EXPR_RULE=34, IMAGINARY=35, 
		DECIMAL_LIT=36, OCTAL_LIT=37, HEX_LIT=38, FLOAT_LIT=39, IMAGINARY_LIT=40, 
		IDENTIFIER=41, NEWLINE=42, WS=43, LINE_COMMENT=44;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "L_PAREN", 
		"R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", 
		"FACTORIAL", "MIN", "MAX", "SIN", "COS", "TAN", "LOG", "SQRT", "PI", "VAR", 
		"CONST", "EULER", "INFINITE", "UNDEFINED", "NUMERIC_RULE", "CONST_RULE", 
		"VAR_RULE", "EXPR_RULE", "IMAGINARY", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT", 
		"FLOAT_LIT", "IMAGINARY_LIT", "IDENTIFIER", "DECIMALS", "OCTAL_DIGIT", 
		"HEX_DIGIT", "EXPONENT", "NEWLINE", "WS", "LINE_COMMENT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2.\u0162\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\3\2\3\2\3\2\3\3\3\3\3\3\3\4"+
		"\3\4\3\5\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f"+
		"\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3"+
		"\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3"+
		"\26\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3"+
		"\32\3\32\3\32\3\33\3\33\3\33\6\33\u00ac\n\33\r\33\16\33\u00ad\5\33\u00b0"+
		"\n\33\3\34\3\34\3\34\6\34\u00b5\n\34\r\34\16\34\u00b6\5\34\u00b9\n\34"+
		"\3\35\3\35\3\36\5\36\u00be\n\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3 \5 \u00cb\n \3 \3 \3 \6 \u00d0\n \r \16 \u00d1\5 \u00d4\n"+
		" \3!\5!\u00d7\n!\3!\3!\3!\6!\u00dc\n!\r!\16!\u00dd\5!\u00e0\n!\3\"\5\""+
		"\u00e3\n\"\3\"\3\"\3\"\6\"\u00e8\n\"\r\"\16\"\u00e9\5\"\u00ec\n\"\3#\5"+
		"#\u00ef\n#\3#\3#\3#\6#\u00f4\n#\r#\16#\u00f5\5#\u00f8\n#\3$\3$\5$\u00fc"+
		"\n$\3$\3$\3%\3%\7%\u0102\n%\f%\16%\u0105\13%\3&\3&\7&\u0109\n&\f&\16&"+
		"\u010c\13&\3\'\3\'\3\'\6\'\u0111\n\'\r\'\16\'\u0112\3(\3(\3(\5(\u0118"+
		"\n(\3(\5(\u011b\n(\3(\5(\u011e\n(\3(\3(\3(\5(\u0123\n(\5(\u0125\n(\3)"+
		"\3)\5)\u0129\n)\3)\3)\3*\3*\7*\u012f\n*\f*\16*\u0132\13*\3+\5+\u0135\n"+
		"+\3+\6+\u0138\n+\r+\16+\u0139\3,\5,\u013d\n,\3,\3,\3-\5-\u0142\n-\3-\3"+
		"-\3.\3.\5.\u0148\n.\3.\3.\3/\5/\u014d\n/\3/\3/\3\60\6\60\u0152\n\60\r"+
		"\60\16\60\u0153\3\60\3\60\3\61\3\61\3\61\3\61\7\61\u015c\n\61\f\61\16"+
		"\61\u015f\13\61\3\61\3\61\2\2\62\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23"+
		"\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31"+
		"\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U\2W\2Y\2[\2"+
		"],_-a.\3\2\24\3\2z|\3\2\62;\3\2ce\3\2gg\4\2--//\3\2PP\3\2EE\3\2XX\3\2"+
		"HH\3\2\63;\4\2ZZzz\4\2C\\c|\6\2\62;C\\aac|\3\2\629\5\2\62;CHch\4\2GGg"+
		"g\4\2\13\13\"\"\4\2\f\f\17\17\2\u0181\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2"+
		"\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2"+
		"\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3"+
		"\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3"+
		"\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65"+
		"\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3"+
		"\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2"+
		"\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\3"+
		"c\3\2\2\2\5f\3\2\2\2\7i\3\2\2\2\tk\3\2\2\2\13n\3\2\2\2\rp\3\2\2\2\17r"+
		"\3\2\2\2\21t\3\2\2\2\23v\3\2\2\2\25x\3\2\2\2\27z\3\2\2\2\31|\3\2\2\2\33"+
		"~\3\2\2\2\35\u0080\3\2\2\2\37\u0082\3\2\2\2!\u0084\3\2\2\2#\u0086\3\2"+
		"\2\2%\u0088\3\2\2\2\'\u008c\3\2\2\2)\u0090\3\2\2\2+\u0094\3\2\2\2-\u0098"+
		"\3\2\2\2/\u009c\3\2\2\2\61\u00a0\3\2\2\2\63\u00a5\3\2\2\2\65\u00a8\3\2"+
		"\2\2\67\u00b1\3\2\2\29\u00ba\3\2\2\2;\u00bd\3\2\2\2=\u00c3\3\2\2\2?\u00ca"+
		"\3\2\2\2A\u00d6\3\2\2\2C\u00e2\3\2\2\2E\u00ee\3\2\2\2G\u00fb\3\2\2\2I"+
		"\u00ff\3\2\2\2K\u0106\3\2\2\2M\u010d\3\2\2\2O\u0124\3\2\2\2Q\u0128\3\2"+
		"\2\2S\u012c\3\2\2\2U\u0134\3\2\2\2W\u013c\3\2\2\2Y\u0141\3\2\2\2[\u0145"+
		"\3\2\2\2]\u014c\3\2\2\2_\u0151\3\2\2\2a\u0157\3\2\2\2cd\7*\2\2de\7+\2"+
		"\2e\4\3\2\2\2fg\7?\2\2gh\7@\2\2h\6\3\2\2\2ij\7`\2\2j\b\3\2\2\2kl\7,\2"+
		"\2lm\7,\2\2m\n\3\2\2\2no\7,\2\2o\f\3\2\2\2pq\7\61\2\2q\16\3\2\2\2rs\7"+
		"-\2\2s\20\3\2\2\2tu\7/\2\2u\22\3\2\2\2vw\7*\2\2w\24\3\2\2\2xy\7+\2\2y"+
		"\26\3\2\2\2z{\7}\2\2{\30\3\2\2\2|}\7\177\2\2}\32\3\2\2\2~\177\7]\2\2\177"+
		"\34\3\2\2\2\u0080\u0081\7_\2\2\u0081\36\3\2\2\2\u0082\u0083\7?\2\2\u0083"+
		" \3\2\2\2\u0084\u0085\7.\2\2\u0085\"\3\2\2\2\u0086\u0087\7#\2\2\u0087"+
		"$\3\2\2\2\u0088\u0089\7o\2\2\u0089\u008a\7k\2\2\u008a\u008b\7p\2\2\u008b"+
		"&\3\2\2\2\u008c\u008d\7o\2\2\u008d\u008e\7c\2\2\u008e\u008f\7z\2\2\u008f"+
		"(\3\2\2\2\u0090\u0091\7u\2\2\u0091\u0092\7k\2\2\u0092\u0093\7p\2\2\u0093"+
		"*\3\2\2\2\u0094\u0095\7e\2\2\u0095\u0096\7q\2\2\u0096\u0097\7u\2\2\u0097"+
		",\3\2\2\2\u0098\u0099\7v\2\2\u0099\u009a\7c\2\2\u009a\u009b\7p\2\2\u009b"+
		".\3\2\2\2\u009c\u009d\7n\2\2\u009d\u009e\7q\2\2\u009e\u009f\7i\2\2\u009f"+
		"\60\3\2\2\2\u00a0\u00a1\7u\2\2\u00a1\u00a2\7s\2\2\u00a2\u00a3\7t\2\2\u00a3"+
		"\u00a4\7v\2\2\u00a4\62\3\2\2\2\u00a5\u00a6\7r\2\2\u00a6\u00a7\7k\2\2\u00a7"+
		"\64\3\2\2\2\u00a8\u00af\t\2\2\2\u00a9\u00ab\7a\2\2\u00aa\u00ac\t\3\2\2"+
		"\u00ab\u00aa\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\u00ab\3\2\2\2\u00ad\u00ae"+
		"\3\2\2\2\u00ae\u00b0\3\2\2\2\u00af\u00a9\3\2\2\2\u00af\u00b0\3\2\2\2\u00b0"+
		"\66\3\2\2\2\u00b1\u00b8\t\4\2\2\u00b2\u00b4\7a\2\2\u00b3\u00b5\t\3\2\2"+
		"\u00b4\u00b3\3\2\2\2\u00b5\u00b6\3\2\2\2\u00b6\u00b4\3\2\2\2\u00b6\u00b7"+
		"\3\2\2\2\u00b7\u00b9\3\2\2\2\u00b8\u00b2\3\2\2\2\u00b8\u00b9\3\2\2\2\u00b9"+
		"8\3\2\2\2\u00ba\u00bb\t\5\2\2\u00bb:\3\2\2\2\u00bc\u00be\t\6\2\2\u00bd"+
		"\u00bc\3\2\2\2\u00bd\u00be\3\2\2\2\u00be\u00bf\3\2\2\2\u00bf\u00c0\7k"+
		"\2\2\u00c0\u00c1\7p\2\2\u00c1\u00c2\7h\2\2\u00c2<\3\2\2\2\u00c3\u00c4"+
		"\7w\2\2\u00c4\u00c5\7p\2\2\u00c5\u00c6\7f\2\2\u00c6\u00c7\7g\2\2\u00c7"+
		"\u00c8\7h\2\2\u00c8>\3\2\2\2\u00c9\u00cb\t\6\2\2\u00ca\u00c9\3\2\2\2\u00ca"+
		"\u00cb\3\2\2\2\u00cb\u00cc\3\2\2\2\u00cc\u00d3\t\7\2\2\u00cd\u00cf\7a"+
		"\2\2\u00ce\u00d0\t\3\2\2\u00cf\u00ce\3\2\2\2\u00d0\u00d1\3\2\2\2\u00d1"+
		"\u00cf\3\2\2\2\u00d1\u00d2\3\2\2\2\u00d2\u00d4\3\2\2\2\u00d3\u00cd\3\2"+
		"\2\2\u00d3\u00d4\3\2\2\2\u00d4@\3\2\2\2\u00d5\u00d7\t\6\2\2\u00d6\u00d5"+
		"\3\2\2\2\u00d6\u00d7\3\2\2\2\u00d7\u00d8\3\2\2\2\u00d8\u00df\t\b\2\2\u00d9"+
		"\u00db\7a\2\2\u00da\u00dc\t\3\2\2\u00db\u00da\3\2\2\2\u00dc\u00dd\3\2"+
		"\2\2\u00dd\u00db\3\2\2\2\u00dd\u00de\3\2\2\2\u00de\u00e0\3\2\2\2\u00df"+
		"\u00d9\3\2\2\2\u00df\u00e0\3\2\2\2\u00e0B\3\2\2\2\u00e1\u00e3\t\6\2\2"+
		"\u00e2\u00e1\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3\u00e4\3\2\2\2\u00e4\u00eb"+
		"\t\t\2\2\u00e5\u00e7\7a\2\2\u00e6\u00e8\t\3\2\2\u00e7\u00e6\3\2\2\2\u00e8"+
		"\u00e9\3\2\2\2\u00e9\u00e7\3\2\2\2\u00e9\u00ea\3\2\2\2\u00ea\u00ec\3\2"+
		"\2\2\u00eb\u00e5\3\2\2\2\u00eb\u00ec\3\2\2\2\u00ecD\3\2\2\2\u00ed\u00ef"+
		"\t\6\2\2\u00ee\u00ed\3\2\2\2\u00ee\u00ef\3\2\2\2\u00ef\u00f0\3\2\2\2\u00f0"+
		"\u00f7\t\n\2\2\u00f1\u00f3\7a\2\2\u00f2\u00f4\t\3\2\2\u00f3\u00f2\3\2"+
		"\2\2\u00f4\u00f5\3\2\2\2\u00f5\u00f3\3\2\2\2\u00f5\u00f6\3\2\2\2\u00f6"+
		"\u00f8\3\2\2\2\u00f7\u00f1\3\2\2\2\u00f7\u00f8\3\2\2\2\u00f8F\3\2\2\2"+
		"\u00f9\u00fc\5U+\2\u00fa\u00fc\5O(\2\u00fb\u00f9\3\2\2\2\u00fb\u00fa\3"+
		"\2\2\2\u00fc\u00fd\3\2\2\2\u00fd\u00fe\7k\2\2\u00feH\3\2\2\2\u00ff\u0103"+
		"\t\13\2\2\u0100\u0102\t\3\2\2\u0101\u0100\3\2\2\2\u0102\u0105\3\2\2\2"+
		"\u0103\u0101\3\2\2\2\u0103\u0104\3\2\2\2\u0104J\3\2\2\2\u0105\u0103\3"+
		"\2\2\2\u0106\u010a\7\62\2\2\u0107\u0109\5W,\2\u0108\u0107\3\2\2\2\u0109"+
		"\u010c\3\2\2\2\u010a\u0108\3\2\2\2\u010a\u010b\3\2\2\2\u010bL\3\2\2\2"+
		"\u010c\u010a\3\2\2\2\u010d\u010e\7\62\2\2\u010e\u0110\t\f\2\2\u010f\u0111"+
		"\5Y-\2\u0110\u010f\3\2\2\2\u0111\u0112\3\2\2\2\u0112\u0110\3\2\2\2\u0112"+
		"\u0113\3\2\2\2\u0113N\3\2\2\2\u0114\u011d\5U+\2\u0115\u0117\7\60\2\2\u0116"+
		"\u0118\5U+\2\u0117\u0116\3\2\2\2\u0117\u0118\3\2\2\2\u0118\u011a\3\2\2"+
		"\2\u0119\u011b\5[.\2\u011a\u0119\3\2\2\2\u011a\u011b\3\2\2\2\u011b\u011e"+
		"\3\2\2\2\u011c\u011e\5[.\2\u011d\u0115\3\2\2\2\u011d\u011c\3\2\2\2\u011e"+
		"\u0125\3\2\2\2\u011f\u0120\7\60\2\2\u0120\u0122\5U+\2\u0121\u0123\5[."+
		"\2\u0122\u0121\3\2\2\2\u0122\u0123\3\2\2\2\u0123\u0125\3\2\2\2\u0124\u0114"+
		"\3\2\2\2\u0124\u011f\3\2\2\2\u0125P\3\2\2\2\u0126\u0129\5U+\2\u0127\u0129"+
		"\5O(\2\u0128\u0126\3\2\2\2\u0128\u0127\3\2\2\2\u0129\u012a\3\2\2\2\u012a"+
		"\u012b\7k\2\2\u012bR\3\2\2\2\u012c\u0130\t\r\2\2\u012d\u012f\t\16\2\2"+
		"\u012e\u012d\3\2\2\2\u012f\u0132\3\2\2\2\u0130\u012e\3\2\2\2\u0130\u0131"+
		"\3\2\2\2\u0131T\3\2\2\2\u0132\u0130\3\2\2\2\u0133\u0135\t\6\2\2\u0134"+
		"\u0133\3\2\2\2\u0134\u0135\3\2\2\2\u0135\u0137\3\2\2\2\u0136\u0138\t\3"+
		"\2\2\u0137\u0136\3\2\2\2\u0138\u0139\3\2\2\2\u0139\u0137\3\2\2\2\u0139"+
		"\u013a\3\2\2\2\u013aV\3\2\2\2\u013b\u013d\t\6\2\2\u013c\u013b\3\2\2\2"+
		"\u013c\u013d\3\2\2\2\u013d\u013e\3\2\2\2\u013e\u013f\t\17\2\2\u013fX\3"+
		"\2\2\2\u0140\u0142\t\6\2\2\u0141\u0140\3\2\2\2\u0141\u0142\3\2\2\2\u0142"+
		"\u0143\3\2\2\2\u0143\u0144\t\20\2\2\u0144Z\3\2\2\2\u0145\u0147\t\21\2"+
		"\2\u0146\u0148\t\6\2\2\u0147\u0146\3\2\2\2\u0147\u0148\3\2\2\2\u0148\u0149"+
		"\3\2\2\2\u0149\u014a\5U+\2\u014a\\\3\2\2\2\u014b\u014d\7\17\2\2\u014c"+
		"\u014b\3\2\2\2\u014c\u014d\3\2\2\2\u014d\u014e\3\2\2\2\u014e\u014f\7\f"+
		"\2\2\u014f^\3\2\2\2\u0150\u0152\t\22\2\2\u0151\u0150\3\2\2\2\u0152\u0153"+
		"\3\2\2\2\u0153\u0151\3\2\2\2\u0153\u0154\3\2\2\2\u0154\u0155\3\2\2\2\u0155"+
		"\u0156\b\60\2\2\u0156`\3\2\2\2\u0157\u0158\7\61\2\2\u0158\u0159\7\61\2"+
		"\2\u0159\u015d\3\2\2\2\u015a\u015c\n\23\2\2\u015b\u015a\3\2\2\2\u015c"+
		"\u015f\3\2\2\2\u015d\u015b\3\2\2\2\u015d\u015e\3\2\2\2\u015e\u0160\3\2"+
		"\2\2\u015f\u015d\3\2\2\2\u0160\u0161\b\61\3\2\u0161b\3\2\2\2\'\2\u00ad"+
		"\u00af\u00b6\u00b8\u00bd\u00ca\u00d1\u00d3\u00d6\u00dd\u00df\u00e2\u00e9"+
		"\u00eb\u00ee\u00f5\u00f7\u00fb\u0103\u010a\u0112\u0117\u011a\u011d\u0122"+
		"\u0124\u0128\u0130\u0134\u0139\u013c\u0141\u0147\u014c\u0153\u015d\4\b"+
		"\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}