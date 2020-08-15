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
		COMMA=16, FACTORIAL=17, SIN=18, COS=19, TAN=20, COT=21, CSC=22, SEC=23, 
		SINH=24, COSH=25, TANH=26, COTH=27, CSCH=28, SECH=29, MIN=30, MAX=31, 
		LOG=32, SQRT=33, DERIVATIVE_OPERATOR=34, PI=35, DERIVATIVE=36, VAR=37, 
		CONST=38, EULER=39, INFINITE=40, UNDEFINED=41, DERIVATE_RULE=42, NUMERIC_RULE=43, 
		CONST_RULE=44, VAR_RULE=45, EXPR_RULE=46, NUMERIC_TREE_RULE=47, CONST_TREE_RULE=48, 
		VAR_TREE_RULE=49, SYSTEM_TREE_RULE=50, IMAGINARY=51, DECIMAL_LIT=52, OCTAL_LIT=53, 
		HEX_LIT=54, FLOAT_LIT=55, IMAGINARY_LIT=56, IDENTIFIER=57, NEWLINE=58, 
		WS=59, LINE_COMMENT=60;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "L_PAREN", 
		"R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", 
		"FACTORIAL", "SIN", "COS", "TAN", "COT", "CSC", "SEC", "SINH", "COSH", 
		"TANH", "COTH", "CSCH", "SECH", "MIN", "MAX", "LOG", "SQRT", "DERIVATIVE_OPERATOR", 
		"PI", "DERIVATIVE", "VAR", "CONST", "EULER", "INFINITE", "UNDEFINED", 
		"DERIVATE_RULE", "NUMERIC_RULE", "CONST_RULE", "VAR_RULE", "EXPR_RULE", 
		"NUMERIC_TREE_RULE", "CONST_TREE_RULE", "VAR_TREE_RULE", "SYSTEM_TREE_RULE", 
		"IMAGINARY", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "IMAGINARY_LIT", 
		"IDENTIFIER", "DECIMALS", "OCTAL_DIGIT", "HEX_DIGIT", "EXPONENT", "NEWLINE", 
		"WS", "LINE_COMMENT"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'()'", "'=>'", "'^'", "'*'", "'/'", "'+'", "'-'", "'**'", "'('", 
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2>\u01de\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\3\2\3\2\3\2\3\3\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3"+
		"\6\3\7\3\7\3\b\3\b\3\t\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16"+
		"\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\23\3\23\3\24"+
		"\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\27\3\27\3\27"+
		"\3\27\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32"+
		"\3\32\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35"+
		"\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3"+
		"!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#\3$\3$\3$\3%\3%\3%\3%\6%\u00f9\n%"+
		"\r%\16%\u00fa\5%\u00fd\n%\3&\3&\3&\6&\u0102\n&\r&\16&\u0103\5&\u0106\n"+
		"&\3\'\3\'\3\'\6\'\u010b\n\'\r\'\16\'\u010c\5\'\u010f\n\'\3(\3(\3)\5)\u0114"+
		"\n)\3)\3)\3)\3)\3*\3*\3*\3*\3*\3*\3+\3+\3+\3+\6+\u0124\n+\r+\16+\u0125"+
		"\5+\u0128\n+\3,\3,\3,\6,\u012d\n,\r,\16,\u012e\5,\u0131\n,\3-\3-\3-\6"+
		"-\u0136\n-\r-\16-\u0137\5-\u013a\n-\3.\3.\3.\6.\u013f\n.\r.\16.\u0140"+
		"\5.\u0143\n.\3/\3/\3/\6/\u0148\n/\r/\16/\u0149\5/\u014c\n/\3\60\5\60\u014f"+
		"\n\60\3\60\3\60\3\60\3\60\6\60\u0155\n\60\r\60\16\60\u0156\5\60\u0159"+
		"\n\60\3\61\5\61\u015c\n\61\3\61\3\61\3\61\3\61\6\61\u0162\n\61\r\61\16"+
		"\61\u0163\5\61\u0166\n\61\3\62\5\62\u0169\n\62\3\62\3\62\3\62\3\62\6\62"+
		"\u016f\n\62\r\62\16\62\u0170\5\62\u0173\n\62\3\63\5\63\u0176\n\63\3\63"+
		"\3\63\3\63\3\63\3\63\6\63\u017d\n\63\r\63\16\63\u017e\5\63\u0181\n\63"+
		"\3\64\3\64\3\65\3\65\7\65\u0187\n\65\f\65\16\65\u018a\13\65\3\66\3\66"+
		"\7\66\u018e\n\66\f\66\16\66\u0191\13\66\3\67\3\67\3\67\6\67\u0196\n\67"+
		"\r\67\16\67\u0197\38\38\38\58\u019d\n8\38\58\u01a0\n8\38\58\u01a3\n8\3"+
		"8\38\38\58\u01a8\n8\58\u01aa\n8\39\39\59\u01ae\n9\39\39\3:\3:\7:\u01b4"+
		"\n:\f:\16:\u01b7\13:\3;\6;\u01ba\n;\r;\16;\u01bb\3<\3<\3=\3=\3>\3>\5>"+
		"\u01c4\n>\3>\3>\3?\5?\u01c9\n?\3?\3?\3@\6@\u01ce\n@\r@\16@\u01cf\3@\3"+
		"@\3A\3A\3A\3A\7A\u01d8\nA\fA\16A\u01db\13A\3A\3A\2\2B\3\3\5\4\7\5\t\6"+
		"\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24"+
		"\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K"+
		"\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u\2w\2y\2{\2"+
		"}<\177=\u0081>\3\2\24\3\2z|\3\2\62;\3\2ce\3\2gg\4\2--//\3\2XX\3\2PP\3"+
		"\2EE\3\2HH\3\2\63;\4\2ZZzz\3\2C\\\6\2\62;C\\aac|\3\2\629\5\2\62;CHch\4"+
		"\2GGgg\4\2\13\13\"\"\4\2\f\f\17\17\2\u0205\2\3\3\2\2\2\2\5\3\2\2\2\2\7"+
		"\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2"+
		"\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2"+
		"\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2"+
		"\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2"+
		"\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2"+
		"\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M"+
		"\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2"+
		"\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2"+
		"\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s"+
		"\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\3\u0083\3\2\2\2\5\u0086"+
		"\3\2\2\2\7\u0089\3\2\2\2\t\u008b\3\2\2\2\13\u008d\3\2\2\2\r\u008f\3\2"+
		"\2\2\17\u0091\3\2\2\2\21\u0093\3\2\2\2\23\u0096\3\2\2\2\25\u0098\3\2\2"+
		"\2\27\u009a\3\2\2\2\31\u009c\3\2\2\2\33\u009e\3\2\2\2\35\u00a0\3\2\2\2"+
		"\37\u00a2\3\2\2\2!\u00a4\3\2\2\2#\u00a6\3\2\2\2%\u00a8\3\2\2\2\'\u00ac"+
		"\3\2\2\2)\u00b0\3\2\2\2+\u00b4\3\2\2\2-\u00b8\3\2\2\2/\u00bc\3\2\2\2\61"+
		"\u00c0\3\2\2\2\63\u00c5\3\2\2\2\65\u00ca\3\2\2\2\67\u00cf\3\2\2\29\u00d4"+
		"\3\2\2\2;\u00d9\3\2\2\2=\u00de\3\2\2\2?\u00e2\3\2\2\2A\u00e6\3\2\2\2C"+
		"\u00ea\3\2\2\2E\u00ef\3\2\2\2G\u00f1\3\2\2\2I\u00f4\3\2\2\2K\u00fe\3\2"+
		"\2\2M\u0107\3\2\2\2O\u0110\3\2\2\2Q\u0113\3\2\2\2S\u0119\3\2\2\2U\u011f"+
		"\3\2\2\2W\u0129\3\2\2\2Y\u0132\3\2\2\2[\u013b\3\2\2\2]\u0144\3\2\2\2_"+
		"\u014e\3\2\2\2a\u015b\3\2\2\2c\u0168\3\2\2\2e\u0175\3\2\2\2g\u0182\3\2"+
		"\2\2i\u0184\3\2\2\2k\u018b\3\2\2\2m\u0192\3\2\2\2o\u01a9\3\2\2\2q\u01ad"+
		"\3\2\2\2s\u01b1\3\2\2\2u\u01b9\3\2\2\2w\u01bd\3\2\2\2y\u01bf\3\2\2\2{"+
		"\u01c1\3\2\2\2}\u01c8\3\2\2\2\177\u01cd\3\2\2\2\u0081\u01d3\3\2\2\2\u0083"+
		"\u0084\7*\2\2\u0084\u0085\7+\2\2\u0085\4\3\2\2\2\u0086\u0087\7?\2\2\u0087"+
		"\u0088\7@\2\2\u0088\6\3\2\2\2\u0089\u008a\7`\2\2\u008a\b\3\2\2\2\u008b"+
		"\u008c\7,\2\2\u008c\n\3\2\2\2\u008d\u008e\7\61\2\2\u008e\f\3\2\2\2\u008f"+
		"\u0090\7-\2\2\u0090\16\3\2\2\2\u0091\u0092\7/\2\2\u0092\20\3\2\2\2\u0093"+
		"\u0094\7,\2\2\u0094\u0095\7,\2\2\u0095\22\3\2\2\2\u0096\u0097\7*\2\2\u0097"+
		"\24\3\2\2\2\u0098\u0099\7+\2\2\u0099\26\3\2\2\2\u009a\u009b\7}\2\2\u009b"+
		"\30\3\2\2\2\u009c\u009d\7\177\2\2\u009d\32\3\2\2\2\u009e\u009f\7]\2\2"+
		"\u009f\34\3\2\2\2\u00a0\u00a1\7_\2\2\u00a1\36\3\2\2\2\u00a2\u00a3\7?\2"+
		"\2\u00a3 \3\2\2\2\u00a4\u00a5\7.\2\2\u00a5\"\3\2\2\2\u00a6\u00a7\7#\2"+
		"\2\u00a7$\3\2\2\2\u00a8\u00a9\7u\2\2\u00a9\u00aa\7k\2\2\u00aa\u00ab\7"+
		"p\2\2\u00ab&\3\2\2\2\u00ac\u00ad\7e\2\2\u00ad\u00ae\7q\2\2\u00ae\u00af"+
		"\7u\2\2\u00af(\3\2\2\2\u00b0\u00b1\7v\2\2\u00b1\u00b2\7c\2\2\u00b2\u00b3"+
		"\7p\2\2\u00b3*\3\2\2\2\u00b4\u00b5\7e\2\2\u00b5\u00b6\7q\2\2\u00b6\u00b7"+
		"\7v\2\2\u00b7,\3\2\2\2\u00b8\u00b9\7e\2\2\u00b9\u00ba\7u\2\2\u00ba\u00bb"+
		"\7e\2\2\u00bb.\3\2\2\2\u00bc\u00bd\7u\2\2\u00bd\u00be\7g\2\2\u00be\u00bf"+
		"\7e\2\2\u00bf\60\3\2\2\2\u00c0\u00c1\7u\2\2\u00c1\u00c2\7k\2\2\u00c2\u00c3"+
		"\7p\2\2\u00c3\u00c4\7j\2\2\u00c4\62\3\2\2\2\u00c5\u00c6\7e\2\2\u00c6\u00c7"+
		"\7q\2\2\u00c7\u00c8\7u\2\2\u00c8\u00c9\7j\2\2\u00c9\64\3\2\2\2\u00ca\u00cb"+
		"\7v\2\2\u00cb\u00cc\7c\2\2\u00cc\u00cd\7p\2\2\u00cd\u00ce\7j\2\2\u00ce"+
		"\66\3\2\2\2\u00cf\u00d0\7e\2\2\u00d0\u00d1\7q\2\2\u00d1\u00d2\7v\2\2\u00d2"+
		"\u00d3\7j\2\2\u00d38\3\2\2\2\u00d4\u00d5\7e\2\2\u00d5\u00d6\7u\2\2\u00d6"+
		"\u00d7\7e\2\2\u00d7\u00d8\7j\2\2\u00d8:\3\2\2\2\u00d9\u00da\7u\2\2\u00da"+
		"\u00db\7g\2\2\u00db\u00dc\7e\2\2\u00dc\u00dd\7j\2\2\u00dd<\3\2\2\2\u00de"+
		"\u00df\7o\2\2\u00df\u00e0\7k\2\2\u00e0\u00e1\7p\2\2\u00e1>\3\2\2\2\u00e2"+
		"\u00e3\7o\2\2\u00e3\u00e4\7c\2\2\u00e4\u00e5\7z\2\2\u00e5@\3\2\2\2\u00e6"+
		"\u00e7\7n\2\2\u00e7\u00e8\7q\2\2\u00e8\u00e9\7i\2\2\u00e9B\3\2\2\2\u00ea"+
		"\u00eb\7u\2\2\u00eb\u00ec\7s\2\2\u00ec\u00ed\7t\2\2\u00ed\u00ee\7v\2\2"+
		"\u00eeD\3\2\2\2\u00ef\u00f0\7F\2\2\u00f0F\3\2\2\2\u00f1\u00f2\7r\2\2\u00f2"+
		"\u00f3\7k\2\2\u00f3H\3\2\2\2\u00f4\u00f5\7f\2\2\u00f5\u00fc\t\2\2\2\u00f6"+
		"\u00f8\7a\2\2\u00f7\u00f9\t\3\2\2\u00f8\u00f7\3\2\2\2\u00f9\u00fa\3\2"+
		"\2\2\u00fa\u00f8\3\2\2\2\u00fa\u00fb\3\2\2\2\u00fb\u00fd\3\2\2\2\u00fc"+
		"\u00f6\3\2\2\2\u00fc\u00fd\3\2\2\2\u00fdJ\3\2\2\2\u00fe\u0105\t\2\2\2"+
		"\u00ff\u0101\7a\2\2\u0100\u0102\t\3\2\2\u0101\u0100\3\2\2\2\u0102\u0103"+
		"\3\2\2\2\u0103\u0101\3\2\2\2\u0103\u0104\3\2\2\2\u0104\u0106\3\2\2\2\u0105"+
		"\u00ff\3\2\2\2\u0105\u0106\3\2\2\2\u0106L\3\2\2\2\u0107\u010e\t\4\2\2"+
		"\u0108\u010a\7a\2\2\u0109\u010b\t\3\2\2\u010a\u0109\3\2\2\2\u010b\u010c"+
		"\3\2\2\2\u010c\u010a\3\2\2\2\u010c\u010d\3\2\2\2\u010d\u010f\3\2\2\2\u010e"+
		"\u0108\3\2\2\2\u010e\u010f\3\2\2\2\u010fN\3\2\2\2\u0110\u0111\t\5\2\2"+
		"\u0111P\3\2\2\2\u0112\u0114\t\6\2\2\u0113\u0112\3\2\2\2\u0113\u0114\3"+
		"\2\2\2\u0114\u0115\3\2\2\2\u0115\u0116\7k\2\2\u0116\u0117\7p\2\2\u0117"+
		"\u0118\7h\2\2\u0118R\3\2\2\2\u0119\u011a\7w\2\2\u011a\u011b\7p\2\2\u011b"+
		"\u011c\7f\2\2\u011c\u011d\7g\2\2\u011d\u011e\7h\2\2\u011eT\3\2\2\2\u011f"+
		"\u0120\7f\2\2\u0120\u0127\t\7\2\2\u0121\u0123\7a\2\2\u0122\u0124\t\3\2"+
		"\2\u0123\u0122\3\2\2\2\u0124\u0125\3\2\2\2\u0125\u0123\3\2\2\2\u0125\u0126"+
		"\3\2\2\2\u0126\u0128\3\2\2\2\u0127\u0121\3\2\2\2\u0127\u0128\3\2\2\2\u0128"+
		"V\3\2\2\2\u0129\u0130\t\b\2\2\u012a\u012c\7a\2\2\u012b\u012d\t\3\2\2\u012c"+
		"\u012b\3\2\2\2\u012d\u012e\3\2\2\2\u012e\u012c\3\2\2\2\u012e\u012f\3\2"+
		"\2\2\u012f\u0131\3\2\2\2\u0130\u012a\3\2\2\2\u0130\u0131\3\2\2\2\u0131"+
		"X\3\2\2\2\u0132\u0139\t\t\2\2\u0133\u0135\7a\2\2\u0134\u0136\t\3\2\2\u0135"+
		"\u0134\3\2\2\2\u0136\u0137\3\2\2\2\u0137\u0135\3\2\2\2\u0137\u0138\3\2"+
		"\2\2\u0138\u013a\3\2\2\2\u0139\u0133\3\2\2\2\u0139\u013a\3\2\2\2\u013a"+
		"Z\3\2\2\2\u013b\u0142\t\7\2\2\u013c\u013e\7a\2\2\u013d\u013f\t\3\2\2\u013e"+
		"\u013d\3\2\2\2\u013f\u0140\3\2\2\2\u0140\u013e\3\2\2\2\u0140\u0141\3\2"+
		"\2\2\u0141\u0143\3\2\2\2\u0142\u013c\3\2\2\2\u0142\u0143\3\2\2\2\u0143"+
		"\\\3\2\2\2\u0144\u014b\t\n\2\2\u0145\u0147\7a\2\2\u0146\u0148\t\3\2\2"+
		"\u0147\u0146\3\2\2\2\u0148\u0149\3\2\2\2\u0149\u0147\3\2\2\2\u0149\u014a"+
		"\3\2\2\2\u014a\u014c\3\2\2\2\u014b\u0145\3\2\2\2\u014b\u014c\3\2\2\2\u014c"+
		"^\3\2\2\2\u014d\u014f\7#\2\2\u014e\u014d\3\2\2\2\u014e\u014f\3\2\2\2\u014f"+
		"\u0150\3\2\2\2\u0150\u0151\7H\2\2\u0151\u0158\t\b\2\2\u0152\u0154\7a\2"+
		"\2\u0153\u0155\t\3\2\2\u0154\u0153\3\2\2\2\u0155\u0156\3\2\2\2\u0156\u0154"+
		"\3\2\2\2\u0156\u0157\3\2\2\2\u0157\u0159\3\2\2\2\u0158\u0152\3\2\2\2\u0158"+
		"\u0159\3\2\2\2\u0159`\3\2\2\2\u015a\u015c\7#\2\2\u015b\u015a\3\2\2\2\u015b"+
		"\u015c\3\2\2\2\u015c\u015d\3\2\2\2\u015d\u015e\7H\2\2\u015e\u0165\t\t"+
		"\2\2\u015f\u0161\7a\2\2\u0160\u0162\t\3\2\2\u0161\u0160\3\2\2\2\u0162"+
		"\u0163\3\2\2\2\u0163\u0161\3\2\2\2\u0163\u0164\3\2\2\2\u0164\u0166\3\2"+
		"\2\2\u0165\u015f\3\2\2\2\u0165\u0166\3\2\2\2\u0166b\3\2\2\2\u0167\u0169"+
		"\7#\2\2\u0168\u0167\3\2\2\2\u0168\u0169\3\2\2\2\u0169\u016a\3\2\2\2\u016a"+
		"\u016b\7H\2\2\u016b\u0172\t\7\2\2\u016c\u016e\7a\2\2\u016d\u016f\t\3\2"+
		"\2\u016e\u016d\3\2\2\2\u016f\u0170\3\2\2\2\u0170\u016e\3\2\2\2\u0170\u0171"+
		"\3\2\2\2\u0171\u0173\3\2\2\2\u0172\u016c\3\2\2\2\u0172\u0173\3\2\2\2\u0173"+
		"d\3\2\2\2\u0174\u0176\7#\2\2\u0175\u0174\3\2\2\2\u0175\u0176\3\2\2\2\u0176"+
		"\u0177\3\2\2\2\u0177\u0178\7U\2\2\u0178\u0179\7H\2\2\u0179\u0180\3\2\2"+
		"\2\u017a\u017c\7a\2\2\u017b\u017d\t\3\2\2\u017c\u017b\3\2\2\2\u017d\u017e"+
		"\3\2\2\2\u017e\u017c\3\2\2\2\u017e\u017f\3\2\2\2\u017f\u0181\3\2\2\2\u0180"+
		"\u017a\3\2\2\2\u0180\u0181\3\2\2\2\u0181f\3\2\2\2\u0182\u0183\7k\2\2\u0183"+
		"h\3\2\2\2\u0184\u0188\t\13\2\2\u0185\u0187\t\3\2\2\u0186\u0185\3\2\2\2"+
		"\u0187\u018a\3\2\2\2\u0188\u0186\3\2\2\2\u0188\u0189\3\2\2\2\u0189j\3"+
		"\2\2\2\u018a\u0188\3\2\2\2\u018b\u018f\7\62\2\2\u018c\u018e\5w<\2\u018d"+
		"\u018c\3\2\2\2\u018e\u0191\3\2\2\2\u018f\u018d\3\2\2\2\u018f\u0190\3\2"+
		"\2\2\u0190l\3\2\2\2\u0191\u018f\3\2\2\2\u0192\u0193\7\62\2\2\u0193\u0195"+
		"\t\f\2\2\u0194\u0196\5y=\2\u0195\u0194\3\2\2\2\u0196\u0197\3\2\2\2\u0197"+
		"\u0195\3\2\2\2\u0197\u0198\3\2\2\2\u0198n\3\2\2\2\u0199\u01a2\5u;\2\u019a"+
		"\u019c\7\60\2\2\u019b\u019d\5u;\2\u019c\u019b\3\2\2\2\u019c\u019d\3\2"+
		"\2\2\u019d\u019f\3\2\2\2\u019e\u01a0\5{>\2\u019f\u019e\3\2\2\2\u019f\u01a0"+
		"\3\2\2\2\u01a0\u01a3\3\2\2\2\u01a1\u01a3\5{>\2\u01a2\u019a\3\2\2\2\u01a2"+
		"\u01a1\3\2\2\2\u01a3\u01aa\3\2\2\2\u01a4\u01a5\7\60\2\2\u01a5\u01a7\5"+
		"u;\2\u01a6\u01a8\5{>\2\u01a7\u01a6\3\2\2\2\u01a7\u01a8\3\2\2\2\u01a8\u01aa"+
		"\3\2\2\2\u01a9\u0199\3\2\2\2\u01a9\u01a4\3\2\2\2\u01aap\3\2\2\2\u01ab"+
		"\u01ae\5u;\2\u01ac\u01ae\5o8\2\u01ad\u01ab\3\2\2\2\u01ad\u01ac\3\2\2\2"+
		"\u01ae\u01af\3\2\2\2\u01af\u01b0\7k\2\2\u01b0r\3\2\2\2\u01b1\u01b5\t\r"+
		"\2\2\u01b2\u01b4\t\16\2\2\u01b3\u01b2\3\2\2\2\u01b4\u01b7\3\2\2\2\u01b5"+
		"\u01b3\3\2\2\2\u01b5\u01b6\3\2\2\2\u01b6t\3\2\2\2\u01b7\u01b5\3\2\2\2"+
		"\u01b8\u01ba\t\3\2\2\u01b9\u01b8\3\2\2\2\u01ba\u01bb\3\2\2\2\u01bb\u01b9"+
		"\3\2\2\2\u01bb\u01bc\3\2\2\2\u01bcv\3\2\2\2\u01bd\u01be\t\17\2\2\u01be"+
		"x\3\2\2\2\u01bf\u01c0\t\20\2\2\u01c0z\3\2\2\2\u01c1\u01c3\t\21\2\2\u01c2"+
		"\u01c4\t\6\2\2\u01c3\u01c2\3\2\2\2\u01c3\u01c4\3\2\2\2\u01c4\u01c5\3\2"+
		"\2\2\u01c5\u01c6\5u;\2\u01c6|\3\2\2\2\u01c7\u01c9\7\17\2\2\u01c8\u01c7"+
		"\3\2\2\2\u01c8\u01c9\3\2\2\2\u01c9\u01ca\3\2\2\2\u01ca\u01cb\7\f\2\2\u01cb"+
		"~\3\2\2\2\u01cc\u01ce\t\22\2\2\u01cd\u01cc\3\2\2\2\u01ce\u01cf\3\2\2\2"+
		"\u01cf\u01cd\3\2\2\2\u01cf\u01d0\3\2\2\2\u01d0\u01d1\3\2\2\2\u01d1\u01d2"+
		"\b@\2\2\u01d2\u0080\3\2\2\2\u01d3\u01d4\7\61\2\2\u01d4\u01d5\7\61\2\2"+
		"\u01d5\u01d9\3\2\2\2\u01d6\u01d8\n\23\2\2\u01d7\u01d6\3\2\2\2\u01d8\u01db"+
		"\3\2\2\2\u01d9\u01d7\3\2\2\2\u01d9\u01da\3\2\2\2\u01da\u01dc\3\2\2\2\u01db"+
		"\u01d9\3\2\2\2\u01dc\u01dd\bA\3\2\u01dd\u0082\3\2\2\2/\2\u00fa\u00fc\u0103"+
		"\u0105\u010c\u010e\u0113\u0125\u0127\u012e\u0130\u0137\u0139\u0140\u0142"+
		"\u0149\u014b\u014e\u0156\u0158\u015b\u0163\u0165\u0168\u0170\u0172\u0175"+
		"\u017e\u0180\u0188\u018f\u0197\u019c\u019f\u01a2\u01a7\u01a9\u01ad\u01b5"+
		"\u01bb\u01c3\u01c8\u01cf\u01d9\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}