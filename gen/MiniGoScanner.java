// Generated from C:/Users/noni4/Desktop/miniGoCompiler/serverAndCompiler/MiniGoScanner.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class MiniGoScanner extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INTLITERAL=1, FLOATLITERAL=2, RUNELITERAL=3, RAWSTRINGLITERAL=4, INTERPRETEDSTRINGLITERAL=5, 
		PACKAGE=6, VAR=7, TYPE=8, FUNC=9, STRUCT=10, LEN=11, CAP=12, PRINT=13, 
		PRINTLN=14, RETURN=15, BREAK=16, CONTINUE=17, APPEND=18, IF=19, ELSE=20, 
		FOR=21, SWITCH=22, CASE=23, DEFAULT=24, IDENTIFIER=25, PLUS=26, MINUS=27, 
		MULTIPLY=28, DIVIDE=29, MODULO=30, LESS=31, LESSEQUAL=32, GREATER=33, 
		GREATEREQUAL=34, EQUAL=35, NOTEQUAL=36, AND=37, OR=38, NOT=39, BITWISEAND=40, 
		BITWISEOR=41, BITWISEXOR=42, BITWISECLEAR=43, SHIFTLEFT=44, SHIFTRIGHT=45, 
		INCREMENT=46, DECREMENT=47, ASSIGN=48, PLUSEQUAL=49, MINUSEQUAL=50, MULTIPLYEQUAL=51, 
		DIVIDEEQUAL=52, MODULOEQUAL=53, BITWISEANDEQUAL=54, BITWISEOREQUAL=55, 
		BITWISEXOREQUAL=56, SHIFTLEFTEQUAL=57, SHIFTRIGHTEQUAL=58, BITWISECLEAREQUAL=59, 
		COLON=60, SEMICOLON=61, COMMA=62, DOT=63, LPAREN=64, RPAREN=65, LBRACKET=66, 
		RBRACKET=67, LBRACE=68, RBRACE=69, SPACES=70, LINE_COMMENT=71, BLOCK_COMMENT=72;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INTLITERAL", "FLOATLITERAL", "RUNELITERAL", "RAWSTRINGLITERAL", "INTERPRETEDSTRINGLITERAL", 
			"PACKAGE", "VAR", "TYPE", "FUNC", "STRUCT", "LEN", "CAP", "PRINT", "PRINTLN", 
			"RETURN", "BREAK", "CONTINUE", "APPEND", "IF", "ELSE", "FOR", "SWITCH", 
			"CASE", "DEFAULT", "IDENTIFIER", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", 
			"MODULO", "LESS", "LESSEQUAL", "GREATER", "GREATEREQUAL", "EQUAL", "NOTEQUAL", 
			"AND", "OR", "NOT", "BITWISEAND", "BITWISEOR", "BITWISEXOR", "BITWISECLEAR", 
			"SHIFTLEFT", "SHIFTRIGHT", "INCREMENT", "DECREMENT", "ASSIGN", "PLUSEQUAL", 
			"MINUSEQUAL", "MULTIPLYEQUAL", "DIVIDEEQUAL", "MODULOEQUAL", "BITWISEANDEQUAL", 
			"BITWISEOREQUAL", "BITWISEXOREQUAL", "SHIFTLEFTEQUAL", "SHIFTRIGHTEQUAL", 
			"BITWISECLEAREQUAL", "COLON", "SEMICOLON", "COMMA", "DOT", "LPAREN", 
			"RPAREN", "LBRACKET", "RBRACKET", "LBRACE", "RBRACE", "SPACES", "LINE_COMMENT", 
			"BLOCK_COMMENT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, null, "'package'", "'var'", "'type'", "'func'", 
			"'struct'", "'len'", "'cap'", "'print'", "'println'", "'return'", "'break'", 
			"'continue'", "'append'", "'if'", "'else'", "'for'", "'switch'", "'case'", 
			"'default'", null, "'+'", "'-'", "'*'", "'/'", "'%'", "'<'", "'<='", 
			"'>'", "'>='", "'=='", "'!='", "'&&'", "'||'", "'!'", "'&'", "'|'", "'^'", 
			"'&^'", "'<<'", "'>>'", "'++'", "'--'", "'='", "'+='", "'-='", "'*='", 
			"'/='", "'%='", "'&='", "'|='", "'^='", "'<<='", "'>>='", "'&^='", "':'", 
			"';'", "','", "'.'", "'('", "')'", "'['", "']'", "'{'", "'}'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INTLITERAL", "FLOATLITERAL", "RUNELITERAL", "RAWSTRINGLITERAL", 
			"INTERPRETEDSTRINGLITERAL", "PACKAGE", "VAR", "TYPE", "FUNC", "STRUCT", 
			"LEN", "CAP", "PRINT", "PRINTLN", "RETURN", "BREAK", "CONTINUE", "APPEND", 
			"IF", "ELSE", "FOR", "SWITCH", "CASE", "DEFAULT", "IDENTIFIER", "PLUS", 
			"MINUS", "MULTIPLY", "DIVIDE", "MODULO", "LESS", "LESSEQUAL", "GREATER", 
			"GREATEREQUAL", "EQUAL", "NOTEQUAL", "AND", "OR", "NOT", "BITWISEAND", 
			"BITWISEOR", "BITWISEXOR", "BITWISECLEAR", "SHIFTLEFT", "SHIFTRIGHT", 
			"INCREMENT", "DECREMENT", "ASSIGN", "PLUSEQUAL", "MINUSEQUAL", "MULTIPLYEQUAL", 
			"DIVIDEEQUAL", "MODULOEQUAL", "BITWISEANDEQUAL", "BITWISEOREQUAL", "BITWISEXOREQUAL", 
			"SHIFTLEFTEQUAL", "SHIFTRIGHTEQUAL", "BITWISECLEAREQUAL", "COLON", "SEMICOLON", 
			"COMMA", "DOT", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LBRACE", 
			"RBRACE", "SPACES", "LINE_COMMENT", "BLOCK_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
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


	public MiniGoScanner(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "MiniGoScanner.g4"; }

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
		"\u0004\u0000H\u01c3\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002"+
		"\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002"+
		"\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002"+
		"\u0015\u0007\u0015\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002"+
		"\u0018\u0007\u0018\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002"+
		"\u001b\u0007\u001b\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002"+
		"\u001e\u0007\u001e\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007"+
		"!\u0002\"\u0007\"\u0002#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007"+
		"&\u0002\'\u0007\'\u0002(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007"+
		"+\u0002,\u0007,\u0002-\u0007-\u0002.\u0007.\u0002/\u0007/\u00020\u0007"+
		"0\u00021\u00071\u00022\u00072\u00023\u00073\u00024\u00074\u00025\u0007"+
		"5\u00026\u00076\u00027\u00077\u00028\u00078\u00029\u00079\u0002:\u0007"+
		":\u0002;\u0007;\u0002<\u0007<\u0002=\u0007=\u0002>\u0007>\u0002?\u0007"+
		"?\u0002@\u0007@\u0002A\u0007A\u0002B\u0007B\u0002C\u0007C\u0002D\u0007"+
		"D\u0002E\u0007E\u0002F\u0007F\u0002G\u0007G\u0001\u0000\u0004\u0000\u0093"+
		"\b\u0000\u000b\u0000\f\u0000\u0094\u0001\u0001\u0004\u0001\u0098\b\u0001"+
		"\u000b\u0001\f\u0001\u0099\u0001\u0001\u0001\u0001\u0004\u0001\u009e\b"+
		"\u0001\u000b\u0001\f\u0001\u009f\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0005\u0003\u00aa"+
		"\b\u0003\n\u0003\f\u0003\u00ad\t\u0003\u0001\u0003\u0001\u0003\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0005\u0004\u00b5\b\u0004\n\u0004"+
		"\f\u0004\u00b8\t\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0018\u0001\u0018"+
		"\u0005\u0018\u012e\b\u0018\n\u0018\f\u0018\u0131\t\u0018\u0001\u0019\u0001"+
		"\u0019\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0001\u001c\u0001"+
		"\u001c\u0001\u001d\u0001\u001d\u0001\u001e\u0001\u001e\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001 \u0001 \u0001!\u0001!\u0001!\u0001\"\u0001\""+
		"\u0001\"\u0001#\u0001#\u0001#\u0001$\u0001$\u0001$\u0001%\u0001%\u0001"+
		"%\u0001&\u0001&\u0001\'\u0001\'\u0001(\u0001(\u0001)\u0001)\u0001*\u0001"+
		"*\u0001*\u0001+\u0001+\u0001+\u0001,\u0001,\u0001,\u0001-\u0001-\u0001"+
		"-\u0001.\u0001.\u0001.\u0001/\u0001/\u00010\u00010\u00010\u00011\u0001"+
		"1\u00011\u00012\u00012\u00012\u00013\u00013\u00013\u00014\u00014\u0001"+
		"4\u00015\u00015\u00015\u00016\u00016\u00016\u00017\u00017\u00017\u0001"+
		"8\u00018\u00018\u00018\u00019\u00019\u00019\u00019\u0001:\u0001:\u0001"+
		":\u0001:\u0001;\u0001;\u0001<\u0001<\u0001=\u0001=\u0001>\u0001>\u0001"+
		"?\u0001?\u0001@\u0001@\u0001A\u0001A\u0001B\u0001B\u0001C\u0001C\u0001"+
		"D\u0001D\u0001E\u0004E\u01a5\bE\u000bE\fE\u01a6\u0001E\u0001E\u0001F\u0001"+
		"F\u0001F\u0001F\u0005F\u01af\bF\nF\fF\u01b2\tF\u0001F\u0001F\u0001G\u0001"+
		"G\u0001G\u0001G\u0005G\u01ba\bG\nG\fG\u01bd\tG\u0001G\u0001G\u0001G\u0001"+
		"G\u0001G\u0001\u01bb\u0000H\u0001\u0001\u0003\u0002\u0005\u0003\u0007"+
		"\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b"+
		"\u0017\f\u0019\r\u001b\u000e\u001d\u000f\u001f\u0010!\u0011#\u0012%\u0013"+
		"\'\u0014)\u0015+\u0016-\u0017/\u00181\u00193\u001a5\u001b7\u001c9\u001d"+
		";\u001e=\u001f? A!C\"E#G$I%K&M\'O(Q)S*U+W,Y-[.]/_0a1c2e3g4i5k6m7o8q9s"+
		":u;w<y={>}?\u007f@\u0081A\u0083B\u0085C\u0087D\u0089E\u008bF\u008dG\u008f"+
		"H\u0001\u0000\u0007\u0001\u000009\u0002\u0000\\\\``\u0002\u0000\"\"\\"+
		"\\\u0003\u0000AZ__az\u0004\u000009AZ__az\u0003\u0000\t\n\r\r  \u0002\u0000"+
		"\n\n\r\r\u01cd\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000"+
		"\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000"+
		"\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000"+
		"\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000"+
		"\u0000\u0011\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000"+
		"\u0000\u0015\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000\u0000\u0000"+
		"\u0000\u0019\u0001\u0000\u0000\u0000\u0000\u001b\u0001\u0000\u0000\u0000"+
		"\u0000\u001d\u0001\u0000\u0000\u0000\u0000\u001f\u0001\u0000\u0000\u0000"+
		"\u0000!\u0001\u0000\u0000\u0000\u0000#\u0001\u0000\u0000\u0000\u0000%"+
		"\u0001\u0000\u0000\u0000\u0000\'\u0001\u0000\u0000\u0000\u0000)\u0001"+
		"\u0000\u0000\u0000\u0000+\u0001\u0000\u0000\u0000\u0000-\u0001\u0000\u0000"+
		"\u0000\u0000/\u0001\u0000\u0000\u0000\u00001\u0001\u0000\u0000\u0000\u0000"+
		"3\u0001\u0000\u0000\u0000\u00005\u0001\u0000\u0000\u0000\u00007\u0001"+
		"\u0000\u0000\u0000\u00009\u0001\u0000\u0000\u0000\u0000;\u0001\u0000\u0000"+
		"\u0000\u0000=\u0001\u0000\u0000\u0000\u0000?\u0001\u0000\u0000\u0000\u0000"+
		"A\u0001\u0000\u0000\u0000\u0000C\u0001\u0000\u0000\u0000\u0000E\u0001"+
		"\u0000\u0000\u0000\u0000G\u0001\u0000\u0000\u0000\u0000I\u0001\u0000\u0000"+
		"\u0000\u0000K\u0001\u0000\u0000\u0000\u0000M\u0001\u0000\u0000\u0000\u0000"+
		"O\u0001\u0000\u0000\u0000\u0000Q\u0001\u0000\u0000\u0000\u0000S\u0001"+
		"\u0000\u0000\u0000\u0000U\u0001\u0000\u0000\u0000\u0000W\u0001\u0000\u0000"+
		"\u0000\u0000Y\u0001\u0000\u0000\u0000\u0000[\u0001\u0000\u0000\u0000\u0000"+
		"]\u0001\u0000\u0000\u0000\u0000_\u0001\u0000\u0000\u0000\u0000a\u0001"+
		"\u0000\u0000\u0000\u0000c\u0001\u0000\u0000\u0000\u0000e\u0001\u0000\u0000"+
		"\u0000\u0000g\u0001\u0000\u0000\u0000\u0000i\u0001\u0000\u0000\u0000\u0000"+
		"k\u0001\u0000\u0000\u0000\u0000m\u0001\u0000\u0000\u0000\u0000o\u0001"+
		"\u0000\u0000\u0000\u0000q\u0001\u0000\u0000\u0000\u0000s\u0001\u0000\u0000"+
		"\u0000\u0000u\u0001\u0000\u0000\u0000\u0000w\u0001\u0000\u0000\u0000\u0000"+
		"y\u0001\u0000\u0000\u0000\u0000{\u0001\u0000\u0000\u0000\u0000}\u0001"+
		"\u0000\u0000\u0000\u0000\u007f\u0001\u0000\u0000\u0000\u0000\u0081\u0001"+
		"\u0000\u0000\u0000\u0000\u0083\u0001\u0000\u0000\u0000\u0000\u0085\u0001"+
		"\u0000\u0000\u0000\u0000\u0087\u0001\u0000\u0000\u0000\u0000\u0089\u0001"+
		"\u0000\u0000\u0000\u0000\u008b\u0001\u0000\u0000\u0000\u0000\u008d\u0001"+
		"\u0000\u0000\u0000\u0000\u008f\u0001\u0000\u0000\u0000\u0001\u0092\u0001"+
		"\u0000\u0000\u0000\u0003\u0097\u0001\u0000\u0000\u0000\u0005\u00a1\u0001"+
		"\u0000\u0000\u0000\u0007\u00a5\u0001\u0000\u0000\u0000\t\u00b0\u0001\u0000"+
		"\u0000\u0000\u000b\u00bb\u0001\u0000\u0000\u0000\r\u00c3\u0001\u0000\u0000"+
		"\u0000\u000f\u00c7\u0001\u0000\u0000\u0000\u0011\u00cc\u0001\u0000\u0000"+
		"\u0000\u0013\u00d1\u0001\u0000\u0000\u0000\u0015\u00d8\u0001\u0000\u0000"+
		"\u0000\u0017\u00dc\u0001\u0000\u0000\u0000\u0019\u00e0\u0001\u0000\u0000"+
		"\u0000\u001b\u00e6\u0001\u0000\u0000\u0000\u001d\u00ee\u0001\u0000\u0000"+
		"\u0000\u001f\u00f5\u0001\u0000\u0000\u0000!\u00fb\u0001\u0000\u0000\u0000"+
		"#\u0104\u0001\u0000\u0000\u0000%\u010b\u0001\u0000\u0000\u0000\'\u010e"+
		"\u0001\u0000\u0000\u0000)\u0113\u0001\u0000\u0000\u0000+\u0117\u0001\u0000"+
		"\u0000\u0000-\u011e\u0001\u0000\u0000\u0000/\u0123\u0001\u0000\u0000\u0000"+
		"1\u012b\u0001\u0000\u0000\u00003\u0132\u0001\u0000\u0000\u00005\u0134"+
		"\u0001\u0000\u0000\u00007\u0136\u0001\u0000\u0000\u00009\u0138\u0001\u0000"+
		"\u0000\u0000;\u013a\u0001\u0000\u0000\u0000=\u013c\u0001\u0000\u0000\u0000"+
		"?\u013e\u0001\u0000\u0000\u0000A\u0141\u0001\u0000\u0000\u0000C\u0143"+
		"\u0001\u0000\u0000\u0000E\u0146\u0001\u0000\u0000\u0000G\u0149\u0001\u0000"+
		"\u0000\u0000I\u014c\u0001\u0000\u0000\u0000K\u014f\u0001\u0000\u0000\u0000"+
		"M\u0152\u0001\u0000\u0000\u0000O\u0154\u0001\u0000\u0000\u0000Q\u0156"+
		"\u0001\u0000\u0000\u0000S\u0158\u0001\u0000\u0000\u0000U\u015a\u0001\u0000"+
		"\u0000\u0000W\u015d\u0001\u0000\u0000\u0000Y\u0160\u0001\u0000\u0000\u0000"+
		"[\u0163\u0001\u0000\u0000\u0000]\u0166\u0001\u0000\u0000\u0000_\u0169"+
		"\u0001\u0000\u0000\u0000a\u016b\u0001\u0000\u0000\u0000c\u016e\u0001\u0000"+
		"\u0000\u0000e\u0171\u0001\u0000\u0000\u0000g\u0174\u0001\u0000\u0000\u0000"+
		"i\u0177\u0001\u0000\u0000\u0000k\u017a\u0001\u0000\u0000\u0000m\u017d"+
		"\u0001\u0000\u0000\u0000o\u0180\u0001\u0000\u0000\u0000q\u0183\u0001\u0000"+
		"\u0000\u0000s\u0187\u0001\u0000\u0000\u0000u\u018b\u0001\u0000\u0000\u0000"+
		"w\u018f\u0001\u0000\u0000\u0000y\u0191\u0001\u0000\u0000\u0000{\u0193"+
		"\u0001\u0000\u0000\u0000}\u0195\u0001\u0000\u0000\u0000\u007f\u0197\u0001"+
		"\u0000\u0000\u0000\u0081\u0199\u0001\u0000\u0000\u0000\u0083\u019b\u0001"+
		"\u0000\u0000\u0000\u0085\u019d\u0001\u0000\u0000\u0000\u0087\u019f\u0001"+
		"\u0000\u0000\u0000\u0089\u01a1\u0001\u0000\u0000\u0000\u008b\u01a4\u0001"+
		"\u0000\u0000\u0000\u008d\u01aa\u0001\u0000\u0000\u0000\u008f\u01b5\u0001"+
		"\u0000\u0000\u0000\u0091\u0093\u0007\u0000\u0000\u0000\u0092\u0091\u0001"+
		"\u0000\u0000\u0000\u0093\u0094\u0001\u0000\u0000\u0000\u0094\u0092\u0001"+
		"\u0000\u0000\u0000\u0094\u0095\u0001\u0000\u0000\u0000\u0095\u0002\u0001"+
		"\u0000\u0000\u0000\u0096\u0098\u0007\u0000\u0000\u0000\u0097\u0096\u0001"+
		"\u0000\u0000\u0000\u0098\u0099\u0001\u0000\u0000\u0000\u0099\u0097\u0001"+
		"\u0000\u0000\u0000\u0099\u009a\u0001\u0000\u0000\u0000\u009a\u009b\u0001"+
		"\u0000\u0000\u0000\u009b\u009d\u0005.\u0000\u0000\u009c\u009e\u0007\u0000"+
		"\u0000\u0000\u009d\u009c\u0001\u0000\u0000\u0000\u009e\u009f\u0001\u0000"+
		"\u0000\u0000\u009f\u009d\u0001\u0000\u0000\u0000\u009f\u00a0\u0001\u0000"+
		"\u0000\u0000\u00a0\u0004\u0001\u0000\u0000\u0000\u00a1\u00a2\u0005\'\u0000"+
		"\u0000\u00a2\u00a3\t\u0000\u0000\u0000\u00a3\u00a4\u0005\'\u0000\u0000"+
		"\u00a4\u0006\u0001\u0000\u0000\u0000\u00a5\u00ab\u0005`\u0000\u0000\u00a6"+
		"\u00a7\u0005\\\u0000\u0000\u00a7\u00aa\t\u0000\u0000\u0000\u00a8\u00aa"+
		"\b\u0001\u0000\u0000\u00a9\u00a6\u0001\u0000\u0000\u0000\u00a9\u00a8\u0001"+
		"\u0000\u0000\u0000\u00aa\u00ad\u0001\u0000\u0000\u0000\u00ab\u00a9\u0001"+
		"\u0000\u0000\u0000\u00ab\u00ac\u0001\u0000\u0000\u0000\u00ac\u00ae\u0001"+
		"\u0000\u0000\u0000\u00ad\u00ab\u0001\u0000\u0000\u0000\u00ae\u00af\u0005"+
		"`\u0000\u0000\u00af\b\u0001\u0000\u0000\u0000\u00b0\u00b6\u0005\"\u0000"+
		"\u0000\u00b1\u00b2\u0005\\\u0000\u0000\u00b2\u00b5\t\u0000\u0000\u0000"+
		"\u00b3\u00b5\b\u0002\u0000\u0000\u00b4\u00b1\u0001\u0000\u0000\u0000\u00b4"+
		"\u00b3\u0001\u0000\u0000\u0000\u00b5\u00b8\u0001\u0000\u0000\u0000\u00b6"+
		"\u00b4\u0001\u0000\u0000\u0000\u00b6\u00b7\u0001\u0000\u0000\u0000\u00b7"+
		"\u00b9\u0001\u0000\u0000\u0000\u00b8\u00b6\u0001\u0000\u0000\u0000\u00b9"+
		"\u00ba\u0005\"\u0000\u0000\u00ba\n\u0001\u0000\u0000\u0000\u00bb\u00bc"+
		"\u0005p\u0000\u0000\u00bc\u00bd\u0005a\u0000\u0000\u00bd\u00be\u0005c"+
		"\u0000\u0000\u00be\u00bf\u0005k\u0000\u0000\u00bf\u00c0\u0005a\u0000\u0000"+
		"\u00c0\u00c1\u0005g\u0000\u0000\u00c1\u00c2\u0005e\u0000\u0000\u00c2\f"+
		"\u0001\u0000\u0000\u0000\u00c3\u00c4\u0005v\u0000\u0000\u00c4\u00c5\u0005"+
		"a\u0000\u0000\u00c5\u00c6\u0005r\u0000\u0000\u00c6\u000e\u0001\u0000\u0000"+
		"\u0000\u00c7\u00c8\u0005t\u0000\u0000\u00c8\u00c9\u0005y\u0000\u0000\u00c9"+
		"\u00ca\u0005p\u0000\u0000\u00ca\u00cb\u0005e\u0000\u0000\u00cb\u0010\u0001"+
		"\u0000\u0000\u0000\u00cc\u00cd\u0005f\u0000\u0000\u00cd\u00ce\u0005u\u0000"+
		"\u0000\u00ce\u00cf\u0005n\u0000\u0000\u00cf\u00d0\u0005c\u0000\u0000\u00d0"+
		"\u0012\u0001\u0000\u0000\u0000\u00d1\u00d2\u0005s\u0000\u0000\u00d2\u00d3"+
		"\u0005t\u0000\u0000\u00d3\u00d4\u0005r\u0000\u0000\u00d4\u00d5\u0005u"+
		"\u0000\u0000\u00d5\u00d6\u0005c\u0000\u0000\u00d6\u00d7\u0005t\u0000\u0000"+
		"\u00d7\u0014\u0001\u0000\u0000\u0000\u00d8\u00d9\u0005l\u0000\u0000\u00d9"+
		"\u00da\u0005e\u0000\u0000\u00da\u00db\u0005n\u0000\u0000\u00db\u0016\u0001"+
		"\u0000\u0000\u0000\u00dc\u00dd\u0005c\u0000\u0000\u00dd\u00de\u0005a\u0000"+
		"\u0000\u00de\u00df\u0005p\u0000\u0000\u00df\u0018\u0001\u0000\u0000\u0000"+
		"\u00e0\u00e1\u0005p\u0000\u0000\u00e1\u00e2\u0005r\u0000\u0000\u00e2\u00e3"+
		"\u0005i\u0000\u0000\u00e3\u00e4\u0005n\u0000\u0000\u00e4\u00e5\u0005t"+
		"\u0000\u0000\u00e5\u001a\u0001\u0000\u0000\u0000\u00e6\u00e7\u0005p\u0000"+
		"\u0000\u00e7\u00e8\u0005r\u0000\u0000\u00e8\u00e9\u0005i\u0000\u0000\u00e9"+
		"\u00ea\u0005n\u0000\u0000\u00ea\u00eb\u0005t\u0000\u0000\u00eb\u00ec\u0005"+
		"l\u0000\u0000\u00ec\u00ed\u0005n\u0000\u0000\u00ed\u001c\u0001\u0000\u0000"+
		"\u0000\u00ee\u00ef\u0005r\u0000\u0000\u00ef\u00f0\u0005e\u0000\u0000\u00f0"+
		"\u00f1\u0005t\u0000\u0000\u00f1\u00f2\u0005u\u0000\u0000\u00f2\u00f3\u0005"+
		"r\u0000\u0000\u00f3\u00f4\u0005n\u0000\u0000\u00f4\u001e\u0001\u0000\u0000"+
		"\u0000\u00f5\u00f6\u0005b\u0000\u0000\u00f6\u00f7\u0005r\u0000\u0000\u00f7"+
		"\u00f8\u0005e\u0000\u0000\u00f8\u00f9\u0005a\u0000\u0000\u00f9\u00fa\u0005"+
		"k\u0000\u0000\u00fa \u0001\u0000\u0000\u0000\u00fb\u00fc\u0005c\u0000"+
		"\u0000\u00fc\u00fd\u0005o\u0000\u0000\u00fd\u00fe\u0005n\u0000\u0000\u00fe"+
		"\u00ff\u0005t\u0000\u0000\u00ff\u0100\u0005i\u0000\u0000\u0100\u0101\u0005"+
		"n\u0000\u0000\u0101\u0102\u0005u\u0000\u0000\u0102\u0103\u0005e\u0000"+
		"\u0000\u0103\"\u0001\u0000\u0000\u0000\u0104\u0105\u0005a\u0000\u0000"+
		"\u0105\u0106\u0005p\u0000\u0000\u0106\u0107\u0005p\u0000\u0000\u0107\u0108"+
		"\u0005e\u0000\u0000\u0108\u0109\u0005n\u0000\u0000\u0109\u010a\u0005d"+
		"\u0000\u0000\u010a$\u0001\u0000\u0000\u0000\u010b\u010c\u0005i\u0000\u0000"+
		"\u010c\u010d\u0005f\u0000\u0000\u010d&\u0001\u0000\u0000\u0000\u010e\u010f"+
		"\u0005e\u0000\u0000\u010f\u0110\u0005l\u0000\u0000\u0110\u0111\u0005s"+
		"\u0000\u0000\u0111\u0112\u0005e\u0000\u0000\u0112(\u0001\u0000\u0000\u0000"+
		"\u0113\u0114\u0005f\u0000\u0000\u0114\u0115\u0005o\u0000\u0000\u0115\u0116"+
		"\u0005r\u0000\u0000\u0116*\u0001\u0000\u0000\u0000\u0117\u0118\u0005s"+
		"\u0000\u0000\u0118\u0119\u0005w\u0000\u0000\u0119\u011a\u0005i\u0000\u0000"+
		"\u011a\u011b\u0005t\u0000\u0000\u011b\u011c\u0005c\u0000\u0000\u011c\u011d"+
		"\u0005h\u0000\u0000\u011d,\u0001\u0000\u0000\u0000\u011e\u011f\u0005c"+
		"\u0000\u0000\u011f\u0120\u0005a\u0000\u0000\u0120\u0121\u0005s\u0000\u0000"+
		"\u0121\u0122\u0005e\u0000\u0000\u0122.\u0001\u0000\u0000\u0000\u0123\u0124"+
		"\u0005d\u0000\u0000\u0124\u0125\u0005e\u0000\u0000\u0125\u0126\u0005f"+
		"\u0000\u0000\u0126\u0127\u0005a\u0000\u0000\u0127\u0128\u0005u\u0000\u0000"+
		"\u0128\u0129\u0005l\u0000\u0000\u0129\u012a\u0005t\u0000\u0000\u012a0"+
		"\u0001\u0000\u0000\u0000\u012b\u012f\u0007\u0003\u0000\u0000\u012c\u012e"+
		"\u0007\u0004\u0000\u0000\u012d\u012c\u0001\u0000\u0000\u0000\u012e\u0131"+
		"\u0001\u0000\u0000\u0000\u012f\u012d\u0001\u0000\u0000\u0000\u012f\u0130"+
		"\u0001\u0000\u0000\u0000\u01302\u0001\u0000\u0000\u0000\u0131\u012f\u0001"+
		"\u0000\u0000\u0000\u0132\u0133\u0005+\u0000\u0000\u01334\u0001\u0000\u0000"+
		"\u0000\u0134\u0135\u0005-\u0000\u0000\u01356\u0001\u0000\u0000\u0000\u0136"+
		"\u0137\u0005*\u0000\u0000\u01378\u0001\u0000\u0000\u0000\u0138\u0139\u0005"+
		"/\u0000\u0000\u0139:\u0001\u0000\u0000\u0000\u013a\u013b\u0005%\u0000"+
		"\u0000\u013b<\u0001\u0000\u0000\u0000\u013c\u013d\u0005<\u0000\u0000\u013d"+
		">\u0001\u0000\u0000\u0000\u013e\u013f\u0005<\u0000\u0000\u013f\u0140\u0005"+
		"=\u0000\u0000\u0140@\u0001\u0000\u0000\u0000\u0141\u0142\u0005>\u0000"+
		"\u0000\u0142B\u0001\u0000\u0000\u0000\u0143\u0144\u0005>\u0000\u0000\u0144"+
		"\u0145\u0005=\u0000\u0000\u0145D\u0001\u0000\u0000\u0000\u0146\u0147\u0005"+
		"=\u0000\u0000\u0147\u0148\u0005=\u0000\u0000\u0148F\u0001\u0000\u0000"+
		"\u0000\u0149\u014a\u0005!\u0000\u0000\u014a\u014b\u0005=\u0000\u0000\u014b"+
		"H\u0001\u0000\u0000\u0000\u014c\u014d\u0005&\u0000\u0000\u014d\u014e\u0005"+
		"&\u0000\u0000\u014eJ\u0001\u0000\u0000\u0000\u014f\u0150\u0005|\u0000"+
		"\u0000\u0150\u0151\u0005|\u0000\u0000\u0151L\u0001\u0000\u0000\u0000\u0152"+
		"\u0153\u0005!\u0000\u0000\u0153N\u0001\u0000\u0000\u0000\u0154\u0155\u0005"+
		"&\u0000\u0000\u0155P\u0001\u0000\u0000\u0000\u0156\u0157\u0005|\u0000"+
		"\u0000\u0157R\u0001\u0000\u0000\u0000\u0158\u0159\u0005^\u0000\u0000\u0159"+
		"T\u0001\u0000\u0000\u0000\u015a\u015b\u0005&\u0000\u0000\u015b\u015c\u0005"+
		"^\u0000\u0000\u015cV\u0001\u0000\u0000\u0000\u015d\u015e\u0005<\u0000"+
		"\u0000\u015e\u015f\u0005<\u0000\u0000\u015fX\u0001\u0000\u0000\u0000\u0160"+
		"\u0161\u0005>\u0000\u0000\u0161\u0162\u0005>\u0000\u0000\u0162Z\u0001"+
		"\u0000\u0000\u0000\u0163\u0164\u0005+\u0000\u0000\u0164\u0165\u0005+\u0000"+
		"\u0000\u0165\\\u0001\u0000\u0000\u0000\u0166\u0167\u0005-\u0000\u0000"+
		"\u0167\u0168\u0005-\u0000\u0000\u0168^\u0001\u0000\u0000\u0000\u0169\u016a"+
		"\u0005=\u0000\u0000\u016a`\u0001\u0000\u0000\u0000\u016b\u016c\u0005+"+
		"\u0000\u0000\u016c\u016d\u0005=\u0000\u0000\u016db\u0001\u0000\u0000\u0000"+
		"\u016e\u016f\u0005-\u0000\u0000\u016f\u0170\u0005=\u0000\u0000\u0170d"+
		"\u0001\u0000\u0000\u0000\u0171\u0172\u0005*\u0000\u0000\u0172\u0173\u0005"+
		"=\u0000\u0000\u0173f\u0001\u0000\u0000\u0000\u0174\u0175\u0005/\u0000"+
		"\u0000\u0175\u0176\u0005=\u0000\u0000\u0176h\u0001\u0000\u0000\u0000\u0177"+
		"\u0178\u0005%\u0000\u0000\u0178\u0179\u0005=\u0000\u0000\u0179j\u0001"+
		"\u0000\u0000\u0000\u017a\u017b\u0005&\u0000\u0000\u017b\u017c\u0005=\u0000"+
		"\u0000\u017cl\u0001\u0000\u0000\u0000\u017d\u017e\u0005|\u0000\u0000\u017e"+
		"\u017f\u0005=\u0000\u0000\u017fn\u0001\u0000\u0000\u0000\u0180\u0181\u0005"+
		"^\u0000\u0000\u0181\u0182\u0005=\u0000\u0000\u0182p\u0001\u0000\u0000"+
		"\u0000\u0183\u0184\u0005<\u0000\u0000\u0184\u0185\u0005<\u0000\u0000\u0185"+
		"\u0186\u0005=\u0000\u0000\u0186r\u0001\u0000\u0000\u0000\u0187\u0188\u0005"+
		">\u0000\u0000\u0188\u0189\u0005>\u0000\u0000\u0189\u018a\u0005=\u0000"+
		"\u0000\u018at\u0001\u0000\u0000\u0000\u018b\u018c\u0005&\u0000\u0000\u018c"+
		"\u018d\u0005^\u0000\u0000\u018d\u018e\u0005=\u0000\u0000\u018ev\u0001"+
		"\u0000\u0000\u0000\u018f\u0190\u0005:\u0000\u0000\u0190x\u0001\u0000\u0000"+
		"\u0000\u0191\u0192\u0005;\u0000\u0000\u0192z\u0001\u0000\u0000\u0000\u0193"+
		"\u0194\u0005,\u0000\u0000\u0194|\u0001\u0000\u0000\u0000\u0195\u0196\u0005"+
		".\u0000\u0000\u0196~\u0001\u0000\u0000\u0000\u0197\u0198\u0005(\u0000"+
		"\u0000\u0198\u0080\u0001\u0000\u0000\u0000\u0199\u019a\u0005)\u0000\u0000"+
		"\u019a\u0082\u0001\u0000\u0000\u0000\u019b\u019c\u0005[\u0000\u0000\u019c"+
		"\u0084\u0001\u0000\u0000\u0000\u019d\u019e\u0005]\u0000\u0000\u019e\u0086"+
		"\u0001\u0000\u0000\u0000\u019f\u01a0\u0005{\u0000\u0000\u01a0\u0088\u0001"+
		"\u0000\u0000\u0000\u01a1\u01a2\u0005}\u0000\u0000\u01a2\u008a\u0001\u0000"+
		"\u0000\u0000\u01a3\u01a5\u0007\u0005\u0000\u0000\u01a4\u01a3\u0001\u0000"+
		"\u0000\u0000\u01a5\u01a6\u0001\u0000\u0000\u0000\u01a6\u01a4\u0001\u0000"+
		"\u0000\u0000\u01a6\u01a7\u0001\u0000\u0000\u0000\u01a7\u01a8\u0001\u0000"+
		"\u0000\u0000\u01a8\u01a9\u0006E\u0000\u0000\u01a9\u008c\u0001\u0000\u0000"+
		"\u0000\u01aa\u01ab\u0005/\u0000\u0000\u01ab\u01ac\u0005/\u0000\u0000\u01ac"+
		"\u01b0\u0001\u0000\u0000\u0000\u01ad\u01af\b\u0006\u0000\u0000\u01ae\u01ad"+
		"\u0001\u0000\u0000\u0000\u01af\u01b2\u0001\u0000\u0000\u0000\u01b0\u01ae"+
		"\u0001\u0000\u0000\u0000\u01b0\u01b1\u0001\u0000\u0000\u0000\u01b1\u01b3"+
		"\u0001\u0000\u0000\u0000\u01b2\u01b0\u0001\u0000\u0000\u0000\u01b3\u01b4"+
		"\u0006F\u0000\u0000\u01b4\u008e\u0001\u0000\u0000\u0000\u01b5\u01b6\u0005"+
		"/\u0000\u0000\u01b6\u01b7\u0005*\u0000\u0000\u01b7\u01bb\u0001\u0000\u0000"+
		"\u0000\u01b8\u01ba\t\u0000\u0000\u0000\u01b9\u01b8\u0001\u0000\u0000\u0000"+
		"\u01ba\u01bd\u0001\u0000\u0000\u0000\u01bb\u01bc\u0001\u0000\u0000\u0000"+
		"\u01bb\u01b9\u0001\u0000\u0000\u0000\u01bc\u01be\u0001\u0000\u0000\u0000"+
		"\u01bd\u01bb\u0001\u0000\u0000\u0000\u01be\u01bf\u0005*\u0000\u0000\u01bf"+
		"\u01c0\u0005/\u0000\u0000\u01c0\u01c1\u0001\u0000\u0000\u0000\u01c1\u01c2"+
		"\u0006G\u0000\u0000\u01c2\u0090\u0001\u0000\u0000\u0000\f\u0000\u0094"+
		"\u0099\u009f\u00a9\u00ab\u00b4\u00b6\u012f\u01a6\u01b0\u01bb\u0001\u0006"+
		"\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}