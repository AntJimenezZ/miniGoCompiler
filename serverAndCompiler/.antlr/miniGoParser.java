// Generated from c:/Users/noni4/Desktop/miniGoCompiler/serverAndCompiler/miniGoParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class miniGoParser extends Parser {
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
	public static final int
		RULE_root = 0, RULE_topDeclarationList = 1, RULE_variableDecl = 2, RULE_innerVarDecls = 3, 
		RULE_singleVarDecl = 4, RULE_singleVarDeclNoExps = 5, RULE_typeDecl = 6, 
		RULE_innerTypeDecls = 7, RULE_singleTypeDecl = 8, RULE_funcDecl = 9, RULE_funcFrontDecl = 10, 
		RULE_funcArgDecls = 11, RULE_declType = 12, RULE_sliceDeclType = 13, RULE_arrayDeclType = 14, 
		RULE_structDeclType = 15, RULE_structMemDecls = 16, RULE_identifierList = 17, 
		RULE_expression = 18, RULE_expressionList = 19, RULE_primaryExpression = 20, 
		RULE_operand = 21, RULE_literal = 22, RULE_index = 23, RULE_arguments = 24, 
		RULE_selector = 25, RULE_appendExpression = 26, RULE_lengthExpression = 27, 
		RULE_capExpression = 28, RULE_statementList = 29, RULE_block = 30, RULE_statement = 31, 
		RULE_simpleStatement = 32, RULE_assignmentStatement = 33, RULE_ifStatement = 34, 
		RULE_loop = 35, RULE_switchStmt = 36, RULE_expressionCaseClauseList = 37, 
		RULE_expressionCaseClause = 38, RULE_expressionSwitchCase = 39, RULE_epsilon = 40;
	private static String[] makeRuleNames() {
		return new String[] {
			"root", "topDeclarationList", "variableDecl", "innerVarDecls", "singleVarDecl", 
			"singleVarDeclNoExps", "typeDecl", "innerTypeDecls", "singleTypeDecl", 
			"funcDecl", "funcFrontDecl", "funcArgDecls", "declType", "sliceDeclType", 
			"arrayDeclType", "structDeclType", "structMemDecls", "identifierList", 
			"expression", "expressionList", "primaryExpression", "operand", "literal", 
			"index", "arguments", "selector", "appendExpression", "lengthExpression", 
			"capExpression", "statementList", "block", "statement", "simpleStatement", 
			"assignmentStatement", "ifStatement", "loop", "switchStmt", "expressionCaseClauseList", 
			"expressionCaseClause", "expressionSwitchCase", "epsilon"
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

	@Override
	public String getGrammarFileName() { return "miniGoParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public miniGoParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RootContext extends ParserRuleContext {
		public RootContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_root; }
	 
		public RootContext() { }
		public void copyFrom(RootContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RootASTContext extends RootContext {
		public TerminalNode PACKAGE() { return getToken(miniGoParser.PACKAGE, 0); }
		public TerminalNode IDENTIFIER() { return getToken(miniGoParser.IDENTIFIER, 0); }
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public TopDeclarationListContext topDeclarationList() {
			return getRuleContext(TopDeclarationListContext.class,0);
		}
		public RootASTContext(RootContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterRootAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitRootAST(this);
		}
	}

	public final RootContext root() throws RecognitionException {
		RootContext _localctx = new RootContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_root);
		try {
			_localctx = new RootASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(82);
			match(PACKAGE);
			setState(83);
			match(IDENTIFIER);
			setState(84);
			match(SEMICOLON);
			setState(85);
			topDeclarationList();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TopDeclarationListContext extends ParserRuleContext {
		public TopDeclarationListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_topDeclarationList; }
	 
		public TopDeclarationListContext() { }
		public void copyFrom(TopDeclarationListContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TopDeclarationListASTContext extends TopDeclarationListContext {
		public List<VariableDeclContext> variableDecl() {
			return getRuleContexts(VariableDeclContext.class);
		}
		public VariableDeclContext variableDecl(int i) {
			return getRuleContext(VariableDeclContext.class,i);
		}
		public List<TypeDeclContext> typeDecl() {
			return getRuleContexts(TypeDeclContext.class);
		}
		public TypeDeclContext typeDecl(int i) {
			return getRuleContext(TypeDeclContext.class,i);
		}
		public List<FuncDeclContext> funcDecl() {
			return getRuleContexts(FuncDeclContext.class);
		}
		public FuncDeclContext funcDecl(int i) {
			return getRuleContext(FuncDeclContext.class,i);
		}
		public TopDeclarationListASTContext(TopDeclarationListContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterTopDeclarationListAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitTopDeclarationListAST(this);
		}
	}

	public final TopDeclarationListContext topDeclarationList() throws RecognitionException {
		TopDeclarationListContext _localctx = new TopDeclarationListContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_topDeclarationList);
		int _la;
		try {
			_localctx = new TopDeclarationListASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(92);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 896L) != 0)) {
				{
				setState(90);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case VAR:
					{
					setState(87);
					variableDecl();
					}
					break;
				case TYPE:
					{
					setState(88);
					typeDecl();
					}
					break;
				case FUNC:
					{
					setState(89);
					funcDecl();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(94);
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

	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclContext extends ParserRuleContext {
		public VariableDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableDecl; }
	 
		public VariableDeclContext() { }
		public void copyFrom(VariableDeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclBlockASTContext extends VariableDeclContext {
		public TerminalNode VAR() { return getToken(miniGoParser.VAR, 0); }
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public InnerVarDeclsContext innerVarDecls() {
			return getRuleContext(InnerVarDeclsContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public VariableDeclBlockASTContext(VariableDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterVariableDeclBlockAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitVariableDeclBlockAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclEmptyASTContext extends VariableDeclContext {
		public TerminalNode VAR() { return getToken(miniGoParser.VAR, 0); }
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public VariableDeclEmptyASTContext(VariableDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterVariableDeclEmptyAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitVariableDeclEmptyAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclASTContext extends VariableDeclContext {
		public TerminalNode VAR() { return getToken(miniGoParser.VAR, 0); }
		public SingleVarDeclContext singleVarDecl() {
			return getRuleContext(SingleVarDeclContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public VariableDeclASTContext(VariableDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterVariableDeclAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitVariableDeclAST(this);
		}
	}

	public final VariableDeclContext variableDecl() throws RecognitionException {
		VariableDeclContext _localctx = new VariableDeclContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_variableDecl);
		try {
			setState(109);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				_localctx = new VariableDeclASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(95);
				match(VAR);
				setState(96);
				singleVarDecl();
				setState(97);
				match(SEMICOLON);
				}
				break;
			case 2:
				_localctx = new VariableDeclBlockASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(99);
				match(VAR);
				setState(100);
				match(LPAREN);
				setState(101);
				innerVarDecls();
				setState(102);
				match(RPAREN);
				setState(103);
				match(SEMICOLON);
				}
				break;
			case 3:
				_localctx = new VariableDeclEmptyASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(105);
				match(VAR);
				setState(106);
				match(LPAREN);
				setState(107);
				match(RPAREN);
				setState(108);
				match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class InnerVarDeclsContext extends ParserRuleContext {
		public InnerVarDeclsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_innerVarDecls; }
	 
		public InnerVarDeclsContext() { }
		public void copyFrom(InnerVarDeclsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class InnerVarDeclsASTContext extends InnerVarDeclsContext {
		public List<SingleVarDeclContext> singleVarDecl() {
			return getRuleContexts(SingleVarDeclContext.class);
		}
		public SingleVarDeclContext singleVarDecl(int i) {
			return getRuleContext(SingleVarDeclContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(miniGoParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(miniGoParser.SEMICOLON, i);
		}
		public InnerVarDeclsASTContext(InnerVarDeclsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterInnerVarDeclsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitInnerVarDeclsAST(this);
		}
	}

	public final InnerVarDeclsContext innerVarDecls() throws RecognitionException {
		InnerVarDeclsContext _localctx = new InnerVarDeclsContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_innerVarDecls);
		int _la;
		try {
			_localctx = new InnerVarDeclsASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(111);
			singleVarDecl();
			setState(112);
			match(SEMICOLON);
			setState(118);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(113);
				singleVarDecl();
				setState(114);
				match(SEMICOLON);
				}
				}
				setState(120);
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

	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarDeclContext extends ParserRuleContext {
		public SingleVarDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleVarDecl; }
	 
		public SingleVarDeclContext() { }
		public void copyFrom(SingleVarDeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarDeclAssignASTContext extends SingleVarDeclContext {
		public IdentifierListContext identifierList() {
			return getRuleContext(IdentifierListContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(miniGoParser.ASSIGN, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public SingleVarDeclAssignASTContext(SingleVarDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSingleVarDeclAssignAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSingleVarDeclAssignAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarDeclASTContext extends SingleVarDeclContext {
		public IdentifierListContext identifierList() {
			return getRuleContext(IdentifierListContext.class,0);
		}
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(miniGoParser.ASSIGN, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public SingleVarDeclASTContext(SingleVarDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSingleVarDeclAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSingleVarDeclAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarDeclNoExpsASTContext extends SingleVarDeclContext {
		public SingleVarDeclNoExpsContext singleVarDeclNoExps() {
			return getRuleContext(SingleVarDeclNoExpsContext.class,0);
		}
		public SingleVarDeclNoExpsASTContext(SingleVarDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSingleVarDeclNoExpsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSingleVarDeclNoExpsAST(this);
		}
	}

	public final SingleVarDeclContext singleVarDecl() throws RecognitionException {
		SingleVarDeclContext _localctx = new SingleVarDeclContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_singleVarDecl);
		try {
			setState(131);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				_localctx = new SingleVarDeclASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(121);
				identifierList();
				setState(122);
				declType();
				setState(123);
				match(ASSIGN);
				setState(124);
				expressionList();
				}
				break;
			case 2:
				_localctx = new SingleVarDeclAssignASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(126);
				identifierList();
				setState(127);
				match(ASSIGN);
				setState(128);
				expressionList();
				}
				break;
			case 3:
				_localctx = new SingleVarDeclNoExpsASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(130);
				singleVarDeclNoExps();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarDeclNoExpsContext extends ParserRuleContext {
		public SingleVarDeclNoExpsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleVarDeclNoExps; }
	 
		public SingleVarDeclNoExpsContext() { }
		public void copyFrom(SingleVarDeclNoExpsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarDeclNoExpsDeclTypeASTContext extends SingleVarDeclNoExpsContext {
		public IdentifierListContext identifierList() {
			return getRuleContext(IdentifierListContext.class,0);
		}
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public SingleVarDeclNoExpsDeclTypeASTContext(SingleVarDeclNoExpsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSingleVarDeclNoExpsDeclTypeAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSingleVarDeclNoExpsDeclTypeAST(this);
		}
	}

	public final SingleVarDeclNoExpsContext singleVarDeclNoExps() throws RecognitionException {
		SingleVarDeclNoExpsContext _localctx = new SingleVarDeclNoExpsContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_singleVarDeclNoExps);
		try {
			_localctx = new SingleVarDeclNoExpsDeclTypeASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			identifierList();
			setState(134);
			declType();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeDeclContext extends ParserRuleContext {
		public TypeDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeDecl; }
	 
		public TypeDeclContext() { }
		public void copyFrom(TypeDeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypeDeclEmptyASTContext extends TypeDeclContext {
		public TerminalNode TYPE() { return getToken(miniGoParser.TYPE, 0); }
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public TypeDeclEmptyASTContext(TypeDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterTypeDeclEmptyAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitTypeDeclEmptyAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypeDeclASTContext extends TypeDeclContext {
		public TerminalNode TYPE() { return getToken(miniGoParser.TYPE, 0); }
		public SingleTypeDeclContext singleTypeDecl() {
			return getRuleContext(SingleTypeDeclContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public TypeDeclASTContext(TypeDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterTypeDeclAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitTypeDeclAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypeDeclBlockASTContext extends TypeDeclContext {
		public TerminalNode TYPE() { return getToken(miniGoParser.TYPE, 0); }
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public InnerTypeDeclsContext innerTypeDecls() {
			return getRuleContext(InnerTypeDeclsContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public TypeDeclBlockASTContext(TypeDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterTypeDeclBlockAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitTypeDeclBlockAST(this);
		}
	}

	public final TypeDeclContext typeDecl() throws RecognitionException {
		TypeDeclContext _localctx = new TypeDeclContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_typeDecl);
		try {
			setState(150);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				_localctx = new TypeDeclASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(136);
				match(TYPE);
				setState(137);
				singleTypeDecl();
				setState(138);
				match(SEMICOLON);
				}
				break;
			case 2:
				_localctx = new TypeDeclBlockASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(140);
				match(TYPE);
				setState(141);
				match(LPAREN);
				setState(142);
				innerTypeDecls();
				setState(143);
				match(RPAREN);
				setState(144);
				match(SEMICOLON);
				}
				break;
			case 3:
				_localctx = new TypeDeclEmptyASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(146);
				match(TYPE);
				setState(147);
				match(LPAREN);
				setState(148);
				match(RPAREN);
				setState(149);
				match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class InnerTypeDeclsContext extends ParserRuleContext {
		public InnerTypeDeclsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_innerTypeDecls; }
	 
		public InnerTypeDeclsContext() { }
		public void copyFrom(InnerTypeDeclsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class InnerTypeDeclsASTContext extends InnerTypeDeclsContext {
		public List<SingleTypeDeclContext> singleTypeDecl() {
			return getRuleContexts(SingleTypeDeclContext.class);
		}
		public SingleTypeDeclContext singleTypeDecl(int i) {
			return getRuleContext(SingleTypeDeclContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(miniGoParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(miniGoParser.SEMICOLON, i);
		}
		public InnerTypeDeclsASTContext(InnerTypeDeclsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterInnerTypeDeclsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitInnerTypeDeclsAST(this);
		}
	}

	public final InnerTypeDeclsContext innerTypeDecls() throws RecognitionException {
		InnerTypeDeclsContext _localctx = new InnerTypeDeclsContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_innerTypeDecls);
		int _la;
		try {
			_localctx = new InnerTypeDeclsASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			singleTypeDecl();
			setState(153);
			match(SEMICOLON);
			setState(159);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(154);
				singleTypeDecl();
				setState(155);
				match(SEMICOLON);
				}
				}
				setState(161);
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

	@SuppressWarnings("CheckReturnValue")
	public static class SingleTypeDeclContext extends ParserRuleContext {
		public SingleTypeDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleTypeDecl; }
	 
		public SingleTypeDeclContext() { }
		public void copyFrom(SingleTypeDeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SingleTypeDeclASTContext extends SingleTypeDeclContext {
		public TerminalNode IDENTIFIER() { return getToken(miniGoParser.IDENTIFIER, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public SingleTypeDeclASTContext(SingleTypeDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSingleTypeDeclAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSingleTypeDeclAST(this);
		}
	}

	public final SingleTypeDeclContext singleTypeDecl() throws RecognitionException {
		SingleTypeDeclContext _localctx = new SingleTypeDeclContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_singleTypeDecl);
		try {
			_localctx = new SingleTypeDeclASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			match(IDENTIFIER);
			setState(163);
			declType();
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

	@SuppressWarnings("CheckReturnValue")
	public static class FuncDeclContext extends ParserRuleContext {
		public FuncDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcDecl; }
	 
		public FuncDeclContext() { }
		public void copyFrom(FuncDeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncDeclASTContext extends FuncDeclContext {
		public FuncFrontDeclContext funcFrontDecl() {
			return getRuleContext(FuncFrontDeclContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public FuncDeclASTContext(FuncDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterFuncDeclAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitFuncDeclAST(this);
		}
	}

	public final FuncDeclContext funcDecl() throws RecognitionException {
		FuncDeclContext _localctx = new FuncDeclContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_funcDecl);
		try {
			_localctx = new FuncDeclASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			funcFrontDecl();
			setState(166);
			block();
			setState(167);
			match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FuncFrontDeclContext extends ParserRuleContext {
		public FuncFrontDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcFrontDecl; }
	 
		public FuncFrontDeclContext() { }
		public void copyFrom(FuncFrontDeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncFrontDeclASTContext extends FuncFrontDeclContext {
		public TerminalNode FUNC() { return getToken(miniGoParser.FUNC, 0); }
		public TerminalNode IDENTIFIER() { return getToken(miniGoParser.IDENTIFIER, 0); }
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public FuncArgDeclsContext funcArgDecls() {
			return getRuleContext(FuncArgDeclsContext.class,0);
		}
		public List<EpsilonContext> epsilon() {
			return getRuleContexts(EpsilonContext.class);
		}
		public EpsilonContext epsilon(int i) {
			return getRuleContext(EpsilonContext.class,i);
		}
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public FuncFrontDeclASTContext(FuncFrontDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterFuncFrontDeclAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitFuncFrontDeclAST(this);
		}
	}

	public final FuncFrontDeclContext funcFrontDecl() throws RecognitionException {
		FuncFrontDeclContext _localctx = new FuncFrontDeclContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_funcFrontDecl);
		try {
			_localctx = new FuncFrontDeclASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			match(FUNC);
			setState(170);
			match(IDENTIFIER);
			setState(171);
			match(LPAREN);
			setState(174);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				{
				setState(172);
				funcArgDecls();
				}
				break;
			case RPAREN:
				{
				setState(173);
				epsilon();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(176);
			match(RPAREN);
			setState(179);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRUCT:
			case IDENTIFIER:
			case LPAREN:
			case LBRACKET:
				{
				setState(177);
				declType();
				}
				break;
			case LBRACE:
				{
				setState(178);
				epsilon();
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FuncArgDeclsContext extends ParserRuleContext {
		public FuncArgDeclsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcArgDecls; }
	 
		public FuncArgDeclsContext() { }
		public void copyFrom(FuncArgDeclsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncArgDeclsASTContext extends FuncArgDeclsContext {
		public List<SingleVarDeclNoExpsContext> singleVarDeclNoExps() {
			return getRuleContexts(SingleVarDeclNoExpsContext.class);
		}
		public SingleVarDeclNoExpsContext singleVarDeclNoExps(int i) {
			return getRuleContext(SingleVarDeclNoExpsContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(miniGoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(miniGoParser.COMMA, i);
		}
		public FuncArgDeclsASTContext(FuncArgDeclsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterFuncArgDeclsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitFuncArgDeclsAST(this);
		}
	}

	public final FuncArgDeclsContext funcArgDecls() throws RecognitionException {
		FuncArgDeclsContext _localctx = new FuncArgDeclsContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_funcArgDecls);
		int _la;
		try {
			_localctx = new FuncArgDeclsASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(181);
			singleVarDeclNoExps();
			setState(186);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(182);
				match(COMMA);
				setState(183);
				singleVarDeclNoExps();
				}
				}
				setState(188);
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

	@SuppressWarnings("CheckReturnValue")
	public static class DeclTypeContext extends ParserRuleContext {
		public DeclTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declType; }
	 
		public DeclTypeContext() { }
		public void copyFrom(DeclTypeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclTypeStructASTContext extends DeclTypeContext {
		public StructDeclTypeContext structDeclType() {
			return getRuleContext(StructDeclTypeContext.class,0);
		}
		public DeclTypeStructASTContext(DeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterDeclTypeStructAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitDeclTypeStructAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclTypeSliceASTContext extends DeclTypeContext {
		public SliceDeclTypeContext sliceDeclType() {
			return getRuleContext(SliceDeclTypeContext.class,0);
		}
		public DeclTypeSliceASTContext(DeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterDeclTypeSliceAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitDeclTypeSliceAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclTypeArrayASTContext extends DeclTypeContext {
		public ArrayDeclTypeContext arrayDeclType() {
			return getRuleContext(ArrayDeclTypeContext.class,0);
		}
		public DeclTypeArrayASTContext(DeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterDeclTypeArrayAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitDeclTypeArrayAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclTypeParenASTContext extends DeclTypeContext {
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public DeclTypeParenASTContext(DeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterDeclTypeParenAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitDeclTypeParenAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclTypeIdentifierASTContext extends DeclTypeContext {
		public TerminalNode IDENTIFIER() { return getToken(miniGoParser.IDENTIFIER, 0); }
		public DeclTypeIdentifierASTContext(DeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterDeclTypeIdentifierAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitDeclTypeIdentifierAST(this);
		}
	}

	public final DeclTypeContext declType() throws RecognitionException {
		DeclTypeContext _localctx = new DeclTypeContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_declType);
		try {
			setState(197);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				_localctx = new DeclTypeParenASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(189);
				match(LPAREN);
				setState(190);
				declType();
				setState(191);
				match(RPAREN);
				}
				break;
			case 2:
				_localctx = new DeclTypeIdentifierASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(193);
				match(IDENTIFIER);
				}
				break;
			case 3:
				_localctx = new DeclTypeSliceASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(194);
				sliceDeclType();
				}
				break;
			case 4:
				_localctx = new DeclTypeArrayASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(195);
				arrayDeclType();
				}
				break;
			case 5:
				_localctx = new DeclTypeStructASTContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(196);
				structDeclType();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SliceDeclTypeContext extends ParserRuleContext {
		public SliceDeclTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceDeclType; }
	 
		public SliceDeclTypeContext() { }
		public void copyFrom(SliceDeclTypeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceDeclTypeASTContext extends SliceDeclTypeContext {
		public TerminalNode LBRACKET() { return getToken(miniGoParser.LBRACKET, 0); }
		public TerminalNode RBRACKET() { return getToken(miniGoParser.RBRACKET, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public SliceDeclTypeASTContext(SliceDeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSliceDeclTypeAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSliceDeclTypeAST(this);
		}
	}

	public final SliceDeclTypeContext sliceDeclType() throws RecognitionException {
		SliceDeclTypeContext _localctx = new SliceDeclTypeContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_sliceDeclType);
		try {
			_localctx = new SliceDeclTypeASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(199);
			match(LBRACKET);
			setState(200);
			match(RBRACKET);
			setState(201);
			declType();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayDeclTypeContext extends ParserRuleContext {
		public ArrayDeclTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayDeclType; }
	 
		public ArrayDeclTypeContext() { }
		public void copyFrom(ArrayDeclTypeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayDeclTypeASTContext extends ArrayDeclTypeContext {
		public TerminalNode LBRACKET() { return getToken(miniGoParser.LBRACKET, 0); }
		public TerminalNode INTLITERAL() { return getToken(miniGoParser.INTLITERAL, 0); }
		public TerminalNode RBRACKET() { return getToken(miniGoParser.RBRACKET, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public ArrayDeclTypeASTContext(ArrayDeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterArrayDeclTypeAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitArrayDeclTypeAST(this);
		}
	}

	public final ArrayDeclTypeContext arrayDeclType() throws RecognitionException {
		ArrayDeclTypeContext _localctx = new ArrayDeclTypeContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_arrayDeclType);
		try {
			_localctx = new ArrayDeclTypeASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(203);
			match(LBRACKET);
			setState(204);
			match(INTLITERAL);
			setState(205);
			match(RBRACKET);
			setState(206);
			declType();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructDeclTypeContext extends ParserRuleContext {
		public StructDeclTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structDeclType; }
	 
		public StructDeclTypeContext() { }
		public void copyFrom(StructDeclTypeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructDeclTypeASTContext extends StructDeclTypeContext {
		public TerminalNode STRUCT() { return getToken(miniGoParser.STRUCT, 0); }
		public TerminalNode LBRACE() { return getToken(miniGoParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(miniGoParser.RBRACE, 0); }
		public StructMemDeclsContext structMemDecls() {
			return getRuleContext(StructMemDeclsContext.class,0);
		}
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public StructDeclTypeASTContext(StructDeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStructDeclTypeAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStructDeclTypeAST(this);
		}
	}

	public final StructDeclTypeContext structDeclType() throws RecognitionException {
		StructDeclTypeContext _localctx = new StructDeclTypeContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_structDeclType);
		try {
			_localctx = new StructDeclTypeASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(208);
			match(STRUCT);
			setState(209);
			match(LBRACE);
			setState(212);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				{
				setState(210);
				structMemDecls();
				}
				break;
			case RBRACE:
				{
				setState(211);
				epsilon();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(214);
			match(RBRACE);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructMemDeclsContext extends ParserRuleContext {
		public StructMemDeclsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structMemDecls; }
	 
		public StructMemDeclsContext() { }
		public void copyFrom(StructMemDeclsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructMemDeclsASTContext extends StructMemDeclsContext {
		public List<SingleVarDeclNoExpsContext> singleVarDeclNoExps() {
			return getRuleContexts(SingleVarDeclNoExpsContext.class);
		}
		public SingleVarDeclNoExpsContext singleVarDeclNoExps(int i) {
			return getRuleContext(SingleVarDeclNoExpsContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(miniGoParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(miniGoParser.SEMICOLON, i);
		}
		public StructMemDeclsASTContext(StructMemDeclsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStructMemDeclsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStructMemDeclsAST(this);
		}
	}

	public final StructMemDeclsContext structMemDecls() throws RecognitionException {
		StructMemDeclsContext _localctx = new StructMemDeclsContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_structMemDecls);
		int _la;
		try {
			_localctx = new StructMemDeclsASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(216);
			singleVarDeclNoExps();
			setState(217);
			match(SEMICOLON);
			setState(223);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(218);
				singleVarDeclNoExps();
				setState(219);
				match(SEMICOLON);
				}
				}
				setState(225);
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

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierListContext extends ParserRuleContext {
		public IdentifierListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifierList; }
	 
		public IdentifierListContext() { }
		public void copyFrom(IdentifierListContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierListASTContext extends IdentifierListContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(miniGoParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(miniGoParser.IDENTIFIER, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(miniGoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(miniGoParser.COMMA, i);
		}
		public IdentifierListASTContext(IdentifierListContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterIdentifierListAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitIdentifierListAST(this);
		}
	}

	public final IdentifierListContext identifierList() throws RecognitionException {
		IdentifierListContext _localctx = new IdentifierListContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_identifierList);
		int _la;
		try {
			_localctx = new IdentifierListASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(226);
			match(IDENTIFIER);
			setState(231);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(227);
				match(COMMA);
				setState(228);
				match(IDENTIFIER);
				}
				}
				setState(233);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	 
		public ExpressionContext() { }
		public void copyFrom(ExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionNotUnaryASTContext extends ExpressionContext {
		public TerminalNode NOT() { return getToken(miniGoParser.NOT, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ExpressionNotUnaryASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionNotUnaryAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionNotUnaryAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionMultiplyASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode MULTIPLY() { return getToken(miniGoParser.MULTIPLY, 0); }
		public ExpressionMultiplyASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionMultiplyAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionMultiplyAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionPlusASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode PLUS() { return getToken(miniGoParser.PLUS, 0); }
		public ExpressionPlusASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionPlusAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionPlusAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionModuloASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode MODULO() { return getToken(miniGoParser.MODULO, 0); }
		public ExpressionModuloASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionModuloAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionModuloAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionAndASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode AND() { return getToken(miniGoParser.AND, 0); }
		public ExpressionAndASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionAndAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionAndAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionBitwiseXorASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode BITWISEXOR() { return getToken(miniGoParser.BITWISEXOR, 0); }
		public ExpressionBitwiseXorASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionBitwiseXorAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionBitwiseXorAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionMinusASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode MINUS() { return getToken(miniGoParser.MINUS, 0); }
		public ExpressionMinusASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionMinusAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionMinusAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionBitwiseXorUnaryASTContext extends ExpressionContext {
		public TerminalNode BITWISEXOR() { return getToken(miniGoParser.BITWISEXOR, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ExpressionBitwiseXorUnaryASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionBitwiseXorUnaryAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionBitwiseXorUnaryAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionEqualASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode EQUAL() { return getToken(miniGoParser.EQUAL, 0); }
		public ExpressionEqualASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionMinusUnaryASTContext extends ExpressionContext {
		public TerminalNode MINUS() { return getToken(miniGoParser.MINUS, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ExpressionMinusUnaryASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionMinusUnaryAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionMinusUnaryAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionBitwiseClearASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode BITWISECLEAR() { return getToken(miniGoParser.BITWISECLEAR, 0); }
		public ExpressionBitwiseClearASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionBitwiseClearAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionBitwiseClearAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionGreaterEqualASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode GREATEREQUAL() { return getToken(miniGoParser.GREATEREQUAL, 0); }
		public ExpressionGreaterEqualASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionGreaterEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionGreaterEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionLessEqualASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode LESSEQUAL() { return getToken(miniGoParser.LESSEQUAL, 0); }
		public ExpressionLessEqualASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionLessEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionLessEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionNotEqualASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode NOTEQUAL() { return getToken(miniGoParser.NOTEQUAL, 0); }
		public ExpressionNotEqualASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionNotEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionNotEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionPrimaryASTContext extends ExpressionContext {
		public PrimaryExpressionContext primaryExpression() {
			return getRuleContext(PrimaryExpressionContext.class,0);
		}
		public ExpressionPrimaryASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionPrimaryAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionPrimaryAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionBitwiseAndASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode BITWISEAND() { return getToken(miniGoParser.BITWISEAND, 0); }
		public ExpressionBitwiseAndASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionBitwiseAndAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionBitwiseAndAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionGreaterASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode GREATER() { return getToken(miniGoParser.GREATER, 0); }
		public ExpressionGreaterASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionGreaterAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionGreaterAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionDivideASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode DIVIDE() { return getToken(miniGoParser.DIVIDE, 0); }
		public ExpressionDivideASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionDivideAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionDivideAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionOrASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode OR() { return getToken(miniGoParser.OR, 0); }
		public ExpressionOrASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionOrAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionOrAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionPlusUnaryASTContext extends ExpressionContext {
		public TerminalNode PLUS() { return getToken(miniGoParser.PLUS, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ExpressionPlusUnaryASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionPlusUnaryAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionPlusUnaryAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionBitwiseOrASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode BITWISEOR() { return getToken(miniGoParser.BITWISEOR, 0); }
		public ExpressionBitwiseOrASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionBitwiseOrAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionBitwiseOrAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionLessASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode LESS() { return getToken(miniGoParser.LESS, 0); }
		public ExpressionLessASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionLessAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionLessAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionShiftRightASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode SHIFTRIGHT() { return getToken(miniGoParser.SHIFTRIGHT, 0); }
		public ExpressionShiftRightASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionShiftRightAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionShiftRightAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionShiftLeftASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode SHIFTLEFT() { return getToken(miniGoParser.SHIFTLEFT, 0); }
		public ExpressionShiftLeftASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionShiftLeftAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionShiftLeftAST(this);
		}
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(244);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTLITERAL:
			case FLOATLITERAL:
			case RUNELITERAL:
			case RAWSTRINGLITERAL:
			case INTERPRETEDSTRINGLITERAL:
			case LEN:
			case CAP:
			case APPEND:
			case IDENTIFIER:
			case LPAREN:
				{
				_localctx = new ExpressionPrimaryASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(235);
				primaryExpression(0);
				}
				break;
			case PLUS:
				{
				_localctx = new ExpressionPlusUnaryASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(236);
				match(PLUS);
				setState(237);
				expression(4);
				}
				break;
			case MINUS:
				{
				_localctx = new ExpressionMinusUnaryASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(238);
				match(MINUS);
				setState(239);
				expression(3);
				}
				break;
			case NOT:
				{
				_localctx = new ExpressionNotUnaryASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(240);
				match(NOT);
				setState(241);
				expression(2);
				}
				break;
			case BITWISEXOR:
				{
				_localctx = new ExpressionBitwiseXorUnaryASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(242);
				match(BITWISEXOR);
				setState(243);
				expression(1);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(305);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(303);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionMultiplyASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(246);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(247);
						match(MULTIPLY);
						setState(248);
						expression(24);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionDivideASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(249);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(250);
						match(DIVIDE);
						setState(251);
						expression(23);
						}
						break;
					case 3:
						{
						_localctx = new ExpressionModuloASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(252);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(253);
						match(MODULO);
						setState(254);
						expression(22);
						}
						break;
					case 4:
						{
						_localctx = new ExpressionShiftLeftASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(255);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(256);
						match(SHIFTLEFT);
						setState(257);
						expression(21);
						}
						break;
					case 5:
						{
						_localctx = new ExpressionShiftRightASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(258);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(259);
						match(SHIFTRIGHT);
						setState(260);
						expression(20);
						}
						break;
					case 6:
						{
						_localctx = new ExpressionBitwiseAndASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(261);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(262);
						match(BITWISEAND);
						setState(263);
						expression(19);
						}
						break;
					case 7:
						{
						_localctx = new ExpressionBitwiseClearASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(264);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(265);
						match(BITWISECLEAR);
						setState(266);
						expression(18);
						}
						break;
					case 8:
						{
						_localctx = new ExpressionPlusASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(267);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(268);
						match(PLUS);
						setState(269);
						expression(17);
						}
						break;
					case 9:
						{
						_localctx = new ExpressionMinusASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(270);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(271);
						match(MINUS);
						setState(272);
						expression(16);
						}
						break;
					case 10:
						{
						_localctx = new ExpressionBitwiseOrASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(273);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(274);
						match(BITWISEOR);
						setState(275);
						expression(15);
						}
						break;
					case 11:
						{
						_localctx = new ExpressionBitwiseXorASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(276);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(277);
						match(BITWISEXOR);
						setState(278);
						expression(14);
						}
						break;
					case 12:
						{
						_localctx = new ExpressionEqualASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(279);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(280);
						match(EQUAL);
						setState(281);
						expression(13);
						}
						break;
					case 13:
						{
						_localctx = new ExpressionNotEqualASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(282);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(283);
						match(NOTEQUAL);
						setState(284);
						expression(12);
						}
						break;
					case 14:
						{
						_localctx = new ExpressionLessASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(285);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(286);
						match(LESS);
						setState(287);
						expression(11);
						}
						break;
					case 15:
						{
						_localctx = new ExpressionLessEqualASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(288);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(289);
						match(LESSEQUAL);
						setState(290);
						expression(10);
						}
						break;
					case 16:
						{
						_localctx = new ExpressionGreaterASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(291);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(292);
						match(GREATER);
						setState(293);
						expression(9);
						}
						break;
					case 17:
						{
						_localctx = new ExpressionGreaterEqualASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(294);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(295);
						match(GREATEREQUAL);
						setState(296);
						expression(8);
						}
						break;
					case 18:
						{
						_localctx = new ExpressionAndASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(297);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(298);
						match(AND);
						setState(299);
						expression(7);
						}
						break;
					case 19:
						{
						_localctx = new ExpressionOrASTContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(300);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(301);
						match(OR);
						setState(302);
						expression(6);
						}
						break;
					}
					} 
				}
				setState(307);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionListContext extends ParserRuleContext {
		public ExpressionListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionList; }
	 
		public ExpressionListContext() { }
		public void copyFrom(ExpressionListContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionListASTContext extends ExpressionListContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(miniGoParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(miniGoParser.COMMA, i);
		}
		public ExpressionListASTContext(ExpressionListContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionListAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionListAST(this);
		}
	}

	public final ExpressionListContext expressionList() throws RecognitionException {
		ExpressionListContext _localctx = new ExpressionListContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_expressionList);
		try {
			int _alt;
			_localctx = new ExpressionListASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(308);
			expression(0);
			setState(313);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(309);
					match(COMMA);
					setState(310);
					expression(0);
					}
					} 
				}
				setState(315);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionContext extends ParserRuleContext {
		public PrimaryExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primaryExpression; }
	 
		public PrimaryExpressionContext() { }
		public void copyFrom(PrimaryExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionLengthASTContext extends PrimaryExpressionContext {
		public LengthExpressionContext lengthExpression() {
			return getRuleContext(LengthExpressionContext.class,0);
		}
		public PrimaryExpressionLengthASTContext(PrimaryExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterPrimaryExpressionLengthAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitPrimaryExpressionLengthAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionOperandASTContext extends PrimaryExpressionContext {
		public OperandContext operand() {
			return getRuleContext(OperandContext.class,0);
		}
		public PrimaryExpressionOperandASTContext(PrimaryExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterPrimaryExpressionOperandAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitPrimaryExpressionOperandAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionIndexASTContext extends PrimaryExpressionContext {
		public PrimaryExpressionContext primaryExpression() {
			return getRuleContext(PrimaryExpressionContext.class,0);
		}
		public IndexContext index() {
			return getRuleContext(IndexContext.class,0);
		}
		public PrimaryExpressionIndexASTContext(PrimaryExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterPrimaryExpressionIndexAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitPrimaryExpressionIndexAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionAppendASTContext extends PrimaryExpressionContext {
		public AppendExpressionContext appendExpression() {
			return getRuleContext(AppendExpressionContext.class,0);
		}
		public PrimaryExpressionAppendASTContext(PrimaryExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterPrimaryExpressionAppendAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitPrimaryExpressionAppendAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionArgumentsASTContext extends PrimaryExpressionContext {
		public PrimaryExpressionContext primaryExpression() {
			return getRuleContext(PrimaryExpressionContext.class,0);
		}
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public PrimaryExpressionArgumentsASTContext(PrimaryExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterPrimaryExpressionArgumentsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitPrimaryExpressionArgumentsAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionCapASTContext extends PrimaryExpressionContext {
		public CapExpressionContext capExpression() {
			return getRuleContext(CapExpressionContext.class,0);
		}
		public PrimaryExpressionCapASTContext(PrimaryExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterPrimaryExpressionCapAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitPrimaryExpressionCapAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionSelectorASTContext extends PrimaryExpressionContext {
		public PrimaryExpressionContext primaryExpression() {
			return getRuleContext(PrimaryExpressionContext.class,0);
		}
		public SelectorContext selector() {
			return getRuleContext(SelectorContext.class,0);
		}
		public PrimaryExpressionSelectorASTContext(PrimaryExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterPrimaryExpressionSelectorAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitPrimaryExpressionSelectorAST(this);
		}
	}

	public final PrimaryExpressionContext primaryExpression() throws RecognitionException {
		return primaryExpression(0);
	}

	private PrimaryExpressionContext primaryExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PrimaryExpressionContext _localctx = new PrimaryExpressionContext(_ctx, _parentState);
		PrimaryExpressionContext _prevctx = _localctx;
		int _startState = 40;
		enterRecursionRule(_localctx, 40, RULE_primaryExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(321);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTLITERAL:
			case FLOATLITERAL:
			case RUNELITERAL:
			case RAWSTRINGLITERAL:
			case INTERPRETEDSTRINGLITERAL:
			case IDENTIFIER:
			case LPAREN:
				{
				_localctx = new PrimaryExpressionOperandASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(317);
				operand();
				}
				break;
			case APPEND:
				{
				_localctx = new PrimaryExpressionAppendASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(318);
				appendExpression();
				}
				break;
			case LEN:
				{
				_localctx = new PrimaryExpressionLengthASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(319);
				lengthExpression();
				}
				break;
			case CAP:
				{
				_localctx = new PrimaryExpressionCapASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(320);
				capExpression();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(331);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(329);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
					case 1:
						{
						_localctx = new PrimaryExpressionSelectorASTContext(new PrimaryExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_primaryExpression);
						setState(323);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(324);
						selector();
						}
						break;
					case 2:
						{
						_localctx = new PrimaryExpressionIndexASTContext(new PrimaryExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_primaryExpression);
						setState(325);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(326);
						index();
						}
						break;
					case 3:
						{
						_localctx = new PrimaryExpressionArgumentsASTContext(new PrimaryExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_primaryExpression);
						setState(327);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(328);
						arguments();
						}
						break;
					}
					} 
				}
				setState(333);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class OperandContext extends ParserRuleContext {
		public OperandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operand; }
	 
		public OperandContext() { }
		public void copyFrom(OperandContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OperandLiteralASTContext extends OperandContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public OperandLiteralASTContext(OperandContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterOperandLiteralAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitOperandLiteralAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OperandIdentifierASTContext extends OperandContext {
		public TerminalNode IDENTIFIER() { return getToken(miniGoParser.IDENTIFIER, 0); }
		public OperandIdentifierASTContext(OperandContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterOperandIdentifierAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitOperandIdentifierAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OperandParenASTContext extends OperandContext {
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public OperandParenASTContext(OperandContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterOperandParenAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitOperandParenAST(this);
		}
	}

	public final OperandContext operand() throws RecognitionException {
		OperandContext _localctx = new OperandContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_operand);
		try {
			setState(340);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTLITERAL:
			case FLOATLITERAL:
			case RUNELITERAL:
			case RAWSTRINGLITERAL:
			case INTERPRETEDSTRINGLITERAL:
				_localctx = new OperandLiteralASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(334);
				literal();
				}
				break;
			case IDENTIFIER:
				_localctx = new OperandIdentifierASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(335);
				match(IDENTIFIER);
				}
				break;
			case LPAREN:
				_localctx = new OperandParenASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(336);
				match(LPAREN);
				setState(337);
				expression(0);
				setState(338);
				match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ParserRuleContext {
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	 
		public LiteralContext() { }
		public void copyFrom(LiteralContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LiteralIntASTContext extends LiteralContext {
		public TerminalNode INTLITERAL() { return getToken(miniGoParser.INTLITERAL, 0); }
		public LiteralIntASTContext(LiteralContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterLiteralIntAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitLiteralIntAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LiteralInterpretedStringASTContext extends LiteralContext {
		public TerminalNode INTERPRETEDSTRINGLITERAL() { return getToken(miniGoParser.INTERPRETEDSTRINGLITERAL, 0); }
		public LiteralInterpretedStringASTContext(LiteralContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterLiteralInterpretedStringAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitLiteralInterpretedStringAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LiteralFloatASTContext extends LiteralContext {
		public TerminalNode FLOATLITERAL() { return getToken(miniGoParser.FLOATLITERAL, 0); }
		public LiteralFloatASTContext(LiteralContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterLiteralFloatAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitLiteralFloatAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LiteralRuneASTContext extends LiteralContext {
		public TerminalNode RUNELITERAL() { return getToken(miniGoParser.RUNELITERAL, 0); }
		public LiteralRuneASTContext(LiteralContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterLiteralRuneAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitLiteralRuneAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LiteralRawStringASTContext extends LiteralContext {
		public TerminalNode RAWSTRINGLITERAL() { return getToken(miniGoParser.RAWSTRINGLITERAL, 0); }
		public LiteralRawStringASTContext(LiteralContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterLiteralRawStringAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitLiteralRawStringAST(this);
		}
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_literal);
		try {
			setState(347);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTLITERAL:
				_localctx = new LiteralIntASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(342);
				match(INTLITERAL);
				}
				break;
			case FLOATLITERAL:
				_localctx = new LiteralFloatASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(343);
				match(FLOATLITERAL);
				}
				break;
			case RUNELITERAL:
				_localctx = new LiteralRuneASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(344);
				match(RUNELITERAL);
				}
				break;
			case RAWSTRINGLITERAL:
				_localctx = new LiteralRawStringASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(345);
				match(RAWSTRINGLITERAL);
				}
				break;
			case INTERPRETEDSTRINGLITERAL:
				_localctx = new LiteralInterpretedStringASTContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(346);
				match(INTERPRETEDSTRINGLITERAL);
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

	@SuppressWarnings("CheckReturnValue")
	public static class IndexContext extends ParserRuleContext {
		public IndexContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_index; }
	 
		public IndexContext() { }
		public void copyFrom(IndexContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IndexASTContext extends IndexContext {
		public TerminalNode LBRACKET() { return getToken(miniGoParser.LBRACKET, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RBRACKET() { return getToken(miniGoParser.RBRACKET, 0); }
		public IndexASTContext(IndexContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterIndexAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitIndexAST(this);
		}
	}

	public final IndexContext index() throws RecognitionException {
		IndexContext _localctx = new IndexContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_index);
		try {
			_localctx = new IndexASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(349);
			match(LBRACKET);
			setState(350);
			expression(0);
			setState(351);
			match(RBRACKET);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArgumentsContext extends ParserRuleContext {
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	 
		public ArgumentsContext() { }
		public void copyFrom(ArgumentsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArgumentsASTContext extends ArgumentsContext {
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public ArgumentsASTContext(ArgumentsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterArgumentsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitArgumentsAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArgumentsEmptyASTContext extends ArgumentsContext {
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public ArgumentsEmptyASTContext(ArgumentsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterArgumentsEmptyAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitArgumentsEmptyAST(this);
		}
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_arguments);
		try {
			setState(358);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
				_localctx = new ArgumentsASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(353);
				match(LPAREN);
				setState(354);
				expressionList();
				}
				break;
			case RPAREN:
				_localctx = new ArgumentsEmptyASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(355);
				epsilon();
				setState(356);
				match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class SelectorContext extends ParserRuleContext {
		public SelectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selector; }
	 
		public SelectorContext() { }
		public void copyFrom(SelectorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SelectorASTContext extends SelectorContext {
		public TerminalNode DOT() { return getToken(miniGoParser.DOT, 0); }
		public TerminalNode IDENTIFIER() { return getToken(miniGoParser.IDENTIFIER, 0); }
		public SelectorASTContext(SelectorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSelectorAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSelectorAST(this);
		}
	}

	public final SelectorContext selector() throws RecognitionException {
		SelectorContext _localctx = new SelectorContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_selector);
		try {
			_localctx = new SelectorASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(360);
			match(DOT);
			setState(361);
			match(IDENTIFIER);
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

	@SuppressWarnings("CheckReturnValue")
	public static class AppendExpressionContext extends ParserRuleContext {
		public AppendExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_appendExpression; }
	 
		public AppendExpressionContext() { }
		public void copyFrom(AppendExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AppendExpressionASTContext extends AppendExpressionContext {
		public TerminalNode APPEND() { return getToken(miniGoParser.APPEND, 0); }
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode COMMA() { return getToken(miniGoParser.COMMA, 0); }
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public AppendExpressionASTContext(AppendExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAppendExpressionAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAppendExpressionAST(this);
		}
	}

	public final AppendExpressionContext appendExpression() throws RecognitionException {
		AppendExpressionContext _localctx = new AppendExpressionContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_appendExpression);
		try {
			_localctx = new AppendExpressionASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			match(APPEND);
			setState(364);
			match(LPAREN);
			setState(365);
			expression(0);
			setState(366);
			match(COMMA);
			setState(367);
			expression(0);
			setState(368);
			match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class LengthExpressionContext extends ParserRuleContext {
		public LengthExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lengthExpression; }
	 
		public LengthExpressionContext() { }
		public void copyFrom(LengthExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LengthExpressionASTContext extends LengthExpressionContext {
		public TerminalNode LEN() { return getToken(miniGoParser.LEN, 0); }
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public LengthExpressionASTContext(LengthExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterLengthExpressionAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitLengthExpressionAST(this);
		}
	}

	public final LengthExpressionContext lengthExpression() throws RecognitionException {
		LengthExpressionContext _localctx = new LengthExpressionContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_lengthExpression);
		try {
			_localctx = new LengthExpressionASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(370);
			match(LEN);
			setState(371);
			match(LPAREN);
			setState(372);
			expression(0);
			setState(373);
			match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class CapExpressionContext extends ParserRuleContext {
		public CapExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_capExpression; }
	 
		public CapExpressionContext() { }
		public void copyFrom(CapExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CapExpressionASTContext extends CapExpressionContext {
		public TerminalNode CAP() { return getToken(miniGoParser.CAP, 0); }
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public CapExpressionASTContext(CapExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterCapExpressionAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitCapExpressionAST(this);
		}
	}

	public final CapExpressionContext capExpression() throws RecognitionException {
		CapExpressionContext _localctx = new CapExpressionContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_capExpression);
		try {
			_localctx = new CapExpressionASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(375);
			match(CAP);
			setState(376);
			match(LPAREN);
			setState(377);
			expression(0);
			setState(378);
			match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StatementListContext extends ParserRuleContext {
		public StatementListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statementList; }
	 
		public StatementListContext() { }
		public void copyFrom(StatementListContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementListASTContext extends StatementListContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public StatementListASTContext(StatementListContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementListAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementListAST(this);
		}
	}

	public final StatementListContext statementList() throws RecognitionException {
		StatementListContext _localctx = new StatementListContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_statementList);
		try {
			int _alt;
			_localctx = new StatementListASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(383);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(380);
					statement();
					}
					} 
				}
				setState(385);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	 
		public BlockContext() { }
		public void copyFrom(BlockContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BlockASTContext extends BlockContext {
		public TerminalNode LBRACE() { return getToken(miniGoParser.LBRACE, 0); }
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(miniGoParser.RBRACE, 0); }
		public BlockASTContext(BlockContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterBlockAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitBlockAST(this);
		}
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_block);
		try {
			_localctx = new BlockASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(386);
			match(LBRACE);
			setState(387);
			statementList();
			setState(388);
			match(RBRACE);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	 
		public StatementContext() { }
		public void copyFrom(StatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementSimpleASTContext extends StatementContext {
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public StatementSimpleASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementSimpleAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementSimpleAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementSwitchASTContext extends StatementContext {
		public SwitchStmtContext switchStmt() {
			return getRuleContext(SwitchStmtContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public StatementSwitchASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementSwitchAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementSwitchAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementLoopASTContext extends StatementContext {
		public LoopContext loop() {
			return getRuleContext(LoopContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public StatementLoopASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementLoopAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementLoopAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementBreakASTContext extends StatementContext {
		public TerminalNode BREAK() { return getToken(miniGoParser.BREAK, 0); }
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public StatementBreakASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementBreakAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementBreakAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementBlockASTContext extends StatementContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public StatementBlockASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementBlockAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementBlockAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementTypeDeclASTContext extends StatementContext {
		public TypeDeclContext typeDecl() {
			return getRuleContext(TypeDeclContext.class,0);
		}
		public StatementTypeDeclASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementTypeDeclAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementTypeDeclAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementReturnASTContext extends StatementContext {
		public TerminalNode RETURN() { return getToken(miniGoParser.RETURN, 0); }
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public StatementReturnASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementReturnAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementReturnAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementPrintASTContext extends StatementContext {
		public TerminalNode PRINT() { return getToken(miniGoParser.PRINT, 0); }
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public StatementPrintASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementPrintAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementPrintAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementVariableDeclASTContext extends StatementContext {
		public VariableDeclContext variableDecl() {
			return getRuleContext(VariableDeclContext.class,0);
		}
		public StatementVariableDeclASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementVariableDeclAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementVariableDeclAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementPrintlnASTContext extends StatementContext {
		public TerminalNode PRINTLN() { return getToken(miniGoParser.PRINTLN, 0); }
		public TerminalNode LPAREN() { return getToken(miniGoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(miniGoParser.RPAREN, 0); }
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public StatementPrintlnASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementPrintlnAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementPrintlnAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementContinueASTContext extends StatementContext {
		public TerminalNode CONTINUE() { return getToken(miniGoParser.CONTINUE, 0); }
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public StatementContinueASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementContinueAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementContinueAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementIfASTContext extends StatementContext {
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public StatementIfASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterStatementIfAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitStatementIfAST(this);
		}
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_statement);
		try {
			setState(436);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				_localctx = new StatementPrintASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(390);
				match(PRINT);
				setState(391);
				match(LPAREN);
				setState(394);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case INTLITERAL:
				case FLOATLITERAL:
				case RUNELITERAL:
				case RAWSTRINGLITERAL:
				case INTERPRETEDSTRINGLITERAL:
				case LEN:
				case CAP:
				case APPEND:
				case IDENTIFIER:
				case PLUS:
				case MINUS:
				case NOT:
				case BITWISEXOR:
				case LPAREN:
					{
					setState(392);
					expressionList();
					}
					break;
				case RPAREN:
					{
					setState(393);
					epsilon();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(396);
				match(RPAREN);
				setState(397);
				match(SEMICOLON);
				}
				break;
			case PRINTLN:
				_localctx = new StatementPrintlnASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(399);
				match(PRINTLN);
				setState(400);
				match(LPAREN);
				setState(403);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case INTLITERAL:
				case FLOATLITERAL:
				case RUNELITERAL:
				case RAWSTRINGLITERAL:
				case INTERPRETEDSTRINGLITERAL:
				case LEN:
				case CAP:
				case APPEND:
				case IDENTIFIER:
				case PLUS:
				case MINUS:
				case NOT:
				case BITWISEXOR:
				case LPAREN:
					{
					setState(401);
					expressionList();
					}
					break;
				case RPAREN:
					{
					setState(402);
					epsilon();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(405);
				match(RPAREN);
				setState(406);
				match(SEMICOLON);
				}
				break;
			case RETURN:
				_localctx = new StatementReturnASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(408);
				match(RETURN);
				setState(411);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case INTLITERAL:
				case FLOATLITERAL:
				case RUNELITERAL:
				case RAWSTRINGLITERAL:
				case INTERPRETEDSTRINGLITERAL:
				case LEN:
				case CAP:
				case APPEND:
				case IDENTIFIER:
				case PLUS:
				case MINUS:
				case NOT:
				case BITWISEXOR:
				case LPAREN:
					{
					setState(409);
					expression(0);
					}
					break;
				case SEMICOLON:
					{
					setState(410);
					epsilon();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(413);
				match(SEMICOLON);
				}
				break;
			case BREAK:
				_localctx = new StatementBreakASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(415);
				match(BREAK);
				setState(416);
				match(SEMICOLON);
				}
				break;
			case CONTINUE:
				_localctx = new StatementContinueASTContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(417);
				match(CONTINUE);
				setState(418);
				match(SEMICOLON);
				}
				break;
			case INTLITERAL:
			case FLOATLITERAL:
			case RUNELITERAL:
			case RAWSTRINGLITERAL:
			case INTERPRETEDSTRINGLITERAL:
			case LEN:
			case CAP:
			case APPEND:
			case IDENTIFIER:
			case PLUS:
			case MINUS:
			case NOT:
			case BITWISEXOR:
			case SEMICOLON:
			case LPAREN:
				_localctx = new StatementSimpleASTContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(419);
				simpleStatement();
				setState(420);
				match(SEMICOLON);
				}
				break;
			case LBRACE:
				_localctx = new StatementBlockASTContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(422);
				block();
				setState(423);
				match(SEMICOLON);
				}
				break;
			case SWITCH:
				_localctx = new StatementSwitchASTContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(425);
				switchStmt();
				setState(426);
				match(SEMICOLON);
				}
				break;
			case IF:
				_localctx = new StatementIfASTContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(428);
				ifStatement();
				setState(429);
				match(SEMICOLON);
				}
				break;
			case FOR:
				_localctx = new StatementLoopASTContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(431);
				loop();
				setState(432);
				match(SEMICOLON);
				}
				break;
			case TYPE:
				_localctx = new StatementTypeDeclASTContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(434);
				typeDecl();
				}
				break;
			case VAR:
				_localctx = new StatementVariableDeclASTContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(435);
				variableDecl();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SimpleStatementContext extends ParserRuleContext {
		public SimpleStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simpleStatement; }
	 
		public SimpleStatementContext() { }
		public void copyFrom(SimpleStatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SimpleStatementExpressionASTContext extends SimpleStatementContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode INCREMENT() { return getToken(miniGoParser.INCREMENT, 0); }
		public TerminalNode DECREMENT() { return getToken(miniGoParser.DECREMENT, 0); }
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public SimpleStatementExpressionASTContext(SimpleStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSimpleStatementExpressionAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSimpleStatementExpressionAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SimpleStatementAssignmentASTContext extends SimpleStatementContext {
		public AssignmentStatementContext assignmentStatement() {
			return getRuleContext(AssignmentStatementContext.class,0);
		}
		public SimpleStatementAssignmentASTContext(SimpleStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSimpleStatementAssignmentAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSimpleStatementAssignmentAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SimpleStatementExpressionListAssignASTContext extends SimpleStatementContext {
		public List<ExpressionListContext> expressionList() {
			return getRuleContexts(ExpressionListContext.class);
		}
		public ExpressionListContext expressionList(int i) {
			return getRuleContext(ExpressionListContext.class,i);
		}
		public TerminalNode ASSIGN() { return getToken(miniGoParser.ASSIGN, 0); }
		public SimpleStatementExpressionListAssignASTContext(SimpleStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSimpleStatementExpressionListAssignAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSimpleStatementExpressionListAssignAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SimpleStatementEmptyASTContext extends SimpleStatementContext {
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public SimpleStatementEmptyASTContext(SimpleStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSimpleStatementEmptyAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSimpleStatementEmptyAST(this);
		}
	}

	public final SimpleStatementContext simpleStatement() throws RecognitionException {
		SimpleStatementContext _localctx = new SimpleStatementContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_simpleStatement);
		try {
			setState(450);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				_localctx = new SimpleStatementEmptyASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(438);
				epsilon();
				}
				break;
			case 2:
				_localctx = new SimpleStatementExpressionASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(439);
				expression(0);
				setState(443);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case INCREMENT:
					{
					setState(440);
					match(INCREMENT);
					}
					break;
				case DECREMENT:
					{
					setState(441);
					match(DECREMENT);
					}
					break;
				case SEMICOLON:
				case LBRACE:
					{
					setState(442);
					epsilon();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				break;
			case 3:
				_localctx = new SimpleStatementAssignmentASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(445);
				assignmentStatement();
				}
				break;
			case 4:
				_localctx = new SimpleStatementExpressionListAssignASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(446);
				expressionList();
				setState(447);
				match(ASSIGN);
				setState(448);
				expressionList();
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

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementContext extends ParserRuleContext {
		public AssignmentStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentStatement; }
	 
		public AssignmentStatementContext() { }
		public void copyFrom(AssignmentStatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementBitwiseAndEqualASTContext extends AssignmentStatementContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode BITWISEANDEQUAL() { return getToken(miniGoParser.BITWISEANDEQUAL, 0); }
		public AssignmentStatementBitwiseAndEqualASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAssignmentStatementBitwiseAndEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAssignmentStatementBitwiseAndEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementBitwiseXorEqualASTContext extends AssignmentStatementContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode BITWISEXOREQUAL() { return getToken(miniGoParser.BITWISEXOREQUAL, 0); }
		public AssignmentStatementBitwiseXorEqualASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAssignmentStatementBitwiseXorEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAssignmentStatementBitwiseXorEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementShiftLeftEqualASTContext extends AssignmentStatementContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode SHIFTLEFTEQUAL() { return getToken(miniGoParser.SHIFTLEFTEQUAL, 0); }
		public AssignmentStatementShiftLeftEqualASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAssignmentStatementShiftLeftEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAssignmentStatementShiftLeftEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementPlusEqualASTContext extends AssignmentStatementContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode PLUSEQUAL() { return getToken(miniGoParser.PLUSEQUAL, 0); }
		public AssignmentStatementPlusEqualASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAssignmentStatementPlusEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAssignmentStatementPlusEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementMinusEqualASTContext extends AssignmentStatementContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode MINUSEQUAL() { return getToken(miniGoParser.MINUSEQUAL, 0); }
		public AssignmentStatementMinusEqualASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAssignmentStatementMinusEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAssignmentStatementMinusEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementModuloEqualASTContext extends AssignmentStatementContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode MODULOEQUAL() { return getToken(miniGoParser.MODULOEQUAL, 0); }
		public AssignmentStatementModuloEqualASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAssignmentStatementModuloEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAssignmentStatementModuloEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementASTContext extends AssignmentStatementContext {
		public List<ExpressionListContext> expressionList() {
			return getRuleContexts(ExpressionListContext.class);
		}
		public ExpressionListContext expressionList(int i) {
			return getRuleContext(ExpressionListContext.class,i);
		}
		public TerminalNode ASSIGN() { return getToken(miniGoParser.ASSIGN, 0); }
		public AssignmentStatementASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAssignmentStatementAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAssignmentStatementAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementBitwiseClearEqualASTContext extends AssignmentStatementContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode BITWISECLEAREQUAL() { return getToken(miniGoParser.BITWISECLEAREQUAL, 0); }
		public AssignmentStatementBitwiseClearEqualASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAssignmentStatementBitwiseClearEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAssignmentStatementBitwiseClearEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementDivideEqualASTContext extends AssignmentStatementContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode DIVIDEEQUAL() { return getToken(miniGoParser.DIVIDEEQUAL, 0); }
		public AssignmentStatementDivideEqualASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAssignmentStatementDivideEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAssignmentStatementDivideEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementShiftRightEqualASTContext extends AssignmentStatementContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode SHIFTRIGHTEQUAL() { return getToken(miniGoParser.SHIFTRIGHTEQUAL, 0); }
		public AssignmentStatementShiftRightEqualASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAssignmentStatementShiftRightEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAssignmentStatementShiftRightEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementMultiplyEqualASTContext extends AssignmentStatementContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode MULTIPLYEQUAL() { return getToken(miniGoParser.MULTIPLYEQUAL, 0); }
		public AssignmentStatementMultiplyEqualASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAssignmentStatementMultiplyEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAssignmentStatementMultiplyEqualAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementBitwiseOrEqualASTContext extends AssignmentStatementContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode BITWISEOREQUAL() { return getToken(miniGoParser.BITWISEOREQUAL, 0); }
		public AssignmentStatementBitwiseOrEqualASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterAssignmentStatementBitwiseOrEqualAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitAssignmentStatementBitwiseOrEqualAST(this);
		}
	}

	public final AssignmentStatementContext assignmentStatement() throws RecognitionException {
		AssignmentStatementContext _localctx = new AssignmentStatementContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_assignmentStatement);
		try {
			setState(500);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				_localctx = new AssignmentStatementASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(452);
				expressionList();
				setState(453);
				match(ASSIGN);
				setState(454);
				expressionList();
				}
				break;
			case 2:
				_localctx = new AssignmentStatementPlusEqualASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(456);
				expression(0);
				setState(457);
				match(PLUSEQUAL);
				setState(458);
				expression(0);
				}
				break;
			case 3:
				_localctx = new AssignmentStatementMinusEqualASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(460);
				expression(0);
				setState(461);
				match(MINUSEQUAL);
				setState(462);
				expression(0);
				}
				break;
			case 4:
				_localctx = new AssignmentStatementBitwiseAndEqualASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(464);
				expression(0);
				setState(465);
				match(BITWISEANDEQUAL);
				setState(466);
				expression(0);
				}
				break;
			case 5:
				_localctx = new AssignmentStatementBitwiseOrEqualASTContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(468);
				expression(0);
				setState(469);
				match(BITWISEOREQUAL);
				setState(470);
				expression(0);
				}
				break;
			case 6:
				_localctx = new AssignmentStatementMultiplyEqualASTContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(472);
				expression(0);
				setState(473);
				match(MULTIPLYEQUAL);
				setState(474);
				expression(0);
				}
				break;
			case 7:
				_localctx = new AssignmentStatementBitwiseXorEqualASTContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(476);
				expression(0);
				setState(477);
				match(BITWISEXOREQUAL);
				setState(478);
				expression(0);
				}
				break;
			case 8:
				_localctx = new AssignmentStatementShiftLeftEqualASTContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(480);
				expression(0);
				setState(481);
				match(SHIFTLEFTEQUAL);
				setState(482);
				expression(0);
				}
				break;
			case 9:
				_localctx = new AssignmentStatementShiftRightEqualASTContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(484);
				expression(0);
				setState(485);
				match(SHIFTRIGHTEQUAL);
				setState(486);
				expression(0);
				}
				break;
			case 10:
				_localctx = new AssignmentStatementBitwiseClearEqualASTContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(488);
				expression(0);
				setState(489);
				match(BITWISECLEAREQUAL);
				setState(490);
				expression(0);
				}
				break;
			case 11:
				_localctx = new AssignmentStatementModuloEqualASTContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(492);
				expression(0);
				setState(493);
				match(MODULOEQUAL);
				setState(494);
				expression(0);
				}
				break;
			case 12:
				_localctx = new AssignmentStatementDivideEqualASTContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(496);
				expression(0);
				setState(497);
				match(DIVIDEEQUAL);
				setState(498);
				expression(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementContext extends ParserRuleContext {
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
	 
		public IfStatementContext() { }
		public void copyFrom(IfStatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfSimpleStatementASTContext extends IfStatementContext {
		public TerminalNode IF() { return getToken(miniGoParser.IF, 0); }
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public IfSimpleStatementASTContext(IfStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterIfSimpleStatementAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitIfSimpleStatementAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfElseIfStatementASTContext extends IfStatementContext {
		public TerminalNode IF() { return getToken(miniGoParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(miniGoParser.ELSE, 0); }
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public IfElseIfStatementASTContext(IfStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterIfElseIfStatementAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitIfElseIfStatementAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementASTContext extends IfStatementContext {
		public TerminalNode IF() { return getToken(miniGoParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public IfStatementASTContext(IfStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterIfStatementAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitIfStatementAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfElseStatementASTContext extends IfStatementContext {
		public TerminalNode IF() { return getToken(miniGoParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(miniGoParser.ELSE, 0); }
		public IfElseStatementASTContext(IfStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterIfElseStatementAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitIfElseStatementAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfSimpleElseStatementASTContext extends IfStatementContext {
		public TerminalNode IF() { return getToken(miniGoParser.IF, 0); }
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(miniGoParser.ELSE, 0); }
		public IfSimpleElseStatementASTContext(IfStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterIfSimpleElseStatementAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitIfSimpleElseStatementAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfSimpleElseIfStatementASTContext extends IfStatementContext {
		public TerminalNode IF() { return getToken(miniGoParser.IF, 0); }
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(miniGoParser.ELSE, 0); }
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public IfSimpleElseIfStatementASTContext(IfStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterIfSimpleElseIfStatementAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitIfSimpleElseIfStatementAST(this);
		}
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_ifStatement);
		try {
			setState(540);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				_localctx = new IfStatementASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(502);
				match(IF);
				setState(503);
				expression(0);
				setState(504);
				block();
				}
				break;
			case 2:
				_localctx = new IfElseIfStatementASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(506);
				match(IF);
				setState(507);
				expression(0);
				setState(508);
				block();
				setState(509);
				match(ELSE);
				setState(510);
				ifStatement();
				}
				break;
			case 3:
				_localctx = new IfElseStatementASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(512);
				match(IF);
				setState(513);
				expression(0);
				setState(514);
				block();
				setState(515);
				match(ELSE);
				setState(516);
				block();
				}
				break;
			case 4:
				_localctx = new IfSimpleStatementASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(518);
				match(IF);
				setState(519);
				simpleStatement();
				setState(520);
				match(SEMICOLON);
				setState(521);
				expression(0);
				setState(522);
				block();
				}
				break;
			case 5:
				_localctx = new IfSimpleElseIfStatementASTContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(524);
				match(IF);
				setState(525);
				simpleStatement();
				setState(526);
				match(SEMICOLON);
				setState(527);
				expression(0);
				setState(528);
				block();
				setState(529);
				match(ELSE);
				setState(530);
				ifStatement();
				}
				break;
			case 6:
				_localctx = new IfSimpleElseStatementASTContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(532);
				match(IF);
				setState(533);
				simpleStatement();
				setState(534);
				match(SEMICOLON);
				setState(535);
				expression(0);
				setState(536);
				block();
				setState(537);
				match(ELSE);
				setState(538);
				block();
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

	@SuppressWarnings("CheckReturnValue")
	public static class LoopContext extends ParserRuleContext {
		public LoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loop; }
	 
		public LoopContext() { }
		public void copyFrom(LoopContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LoopExpressionBlockASTContext extends LoopContext {
		public TerminalNode FOR() { return getToken(miniGoParser.FOR, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public LoopExpressionBlockASTContext(LoopContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterLoopExpressionBlockAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitLoopExpressionBlockAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LoopBlockASTContext extends LoopContext {
		public TerminalNode FOR() { return getToken(miniGoParser.FOR, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public LoopBlockASTContext(LoopContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterLoopBlockAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitLoopBlockAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LoopSimpleStatementExpressionSimpleStatementBlockASTContext extends LoopContext {
		public TerminalNode FOR() { return getToken(miniGoParser.FOR, 0); }
		public List<SimpleStatementContext> simpleStatement() {
			return getRuleContexts(SimpleStatementContext.class);
		}
		public SimpleStatementContext simpleStatement(int i) {
			return getRuleContext(SimpleStatementContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(miniGoParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(miniGoParser.SEMICOLON, i);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public LoopSimpleStatementExpressionSimpleStatementBlockASTContext(LoopContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterLoopSimpleStatementExpressionSimpleStatementBlockAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitLoopSimpleStatementExpressionSimpleStatementBlockAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LoopSimpleStatementSimpleStatementBlockASTContext extends LoopContext {
		public TerminalNode FOR() { return getToken(miniGoParser.FOR, 0); }
		public List<SimpleStatementContext> simpleStatement() {
			return getRuleContexts(SimpleStatementContext.class);
		}
		public SimpleStatementContext simpleStatement(int i) {
			return getRuleContext(SimpleStatementContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(miniGoParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(miniGoParser.SEMICOLON, i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public LoopSimpleStatementSimpleStatementBlockASTContext(LoopContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterLoopSimpleStatementSimpleStatementBlockAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitLoopSimpleStatementSimpleStatementBlockAST(this);
		}
	}

	public final LoopContext loop() throws RecognitionException {
		LoopContext _localctx = new LoopContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_loop);
		try {
			setState(563);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				_localctx = new LoopBlockASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(542);
				match(FOR);
				setState(543);
				block();
				}
				break;
			case 2:
				_localctx = new LoopExpressionBlockASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(544);
				match(FOR);
				setState(545);
				expression(0);
				setState(546);
				block();
				}
				break;
			case 3:
				_localctx = new LoopSimpleStatementExpressionSimpleStatementBlockASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(548);
				match(FOR);
				setState(549);
				simpleStatement();
				setState(550);
				match(SEMICOLON);
				setState(551);
				expression(0);
				setState(552);
				match(SEMICOLON);
				setState(553);
				simpleStatement();
				setState(554);
				block();
				}
				break;
			case 4:
				_localctx = new LoopSimpleStatementSimpleStatementBlockASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(556);
				match(FOR);
				setState(557);
				simpleStatement();
				setState(558);
				match(SEMICOLON);
				setState(559);
				match(SEMICOLON);
				setState(560);
				simpleStatement();
				setState(561);
				block();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStmtContext extends ParserRuleContext {
		public SwitchStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchStmt; }
	 
		public SwitchStmtContext() { }
		public void copyFrom(SwitchStmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStmtSimpleStatementASTContext extends SwitchStmtContext {
		public TerminalNode SWITCH() { return getToken(miniGoParser.SWITCH, 0); }
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(miniGoParser.LBRACE, 0); }
		public ExpressionCaseClauseListContext expressionCaseClauseList() {
			return getRuleContext(ExpressionCaseClauseListContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(miniGoParser.RBRACE, 0); }
		public SwitchStmtSimpleStatementASTContext(SwitchStmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSwitchStmtSimpleStatementAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSwitchStmtSimpleStatementAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStmtExpressionASTContext extends SwitchStmtContext {
		public TerminalNode SWITCH() { return getToken(miniGoParser.SWITCH, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(miniGoParser.LBRACE, 0); }
		public ExpressionCaseClauseListContext expressionCaseClauseList() {
			return getRuleContext(ExpressionCaseClauseListContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(miniGoParser.RBRACE, 0); }
		public SwitchStmtExpressionASTContext(SwitchStmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSwitchStmtExpressionAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSwitchStmtExpressionAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStmtSimpleStatementBlockASTContext extends SwitchStmtContext {
		public TerminalNode SWITCH() { return getToken(miniGoParser.SWITCH, 0); }
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(miniGoParser.SEMICOLON, 0); }
		public TerminalNode LBRACE() { return getToken(miniGoParser.LBRACE, 0); }
		public ExpressionCaseClauseListContext expressionCaseClauseList() {
			return getRuleContext(ExpressionCaseClauseListContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(miniGoParser.RBRACE, 0); }
		public SwitchStmtSimpleStatementBlockASTContext(SwitchStmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSwitchStmtSimpleStatementBlockAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSwitchStmtSimpleStatementBlockAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStmtBlockASTContext extends SwitchStmtContext {
		public TerminalNode SWITCH() { return getToken(miniGoParser.SWITCH, 0); }
		public TerminalNode LBRACE() { return getToken(miniGoParser.LBRACE, 0); }
		public ExpressionCaseClauseListContext expressionCaseClauseList() {
			return getRuleContext(ExpressionCaseClauseListContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(miniGoParser.RBRACE, 0); }
		public SwitchStmtBlockASTContext(SwitchStmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterSwitchStmtBlockAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitSwitchStmtBlockAST(this);
		}
	}

	public final SwitchStmtContext switchStmt() throws RecognitionException {
		SwitchStmtContext _localctx = new SwitchStmtContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_switchStmt);
		try {
			setState(591);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				_localctx = new SwitchStmtSimpleStatementASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(565);
				match(SWITCH);
				setState(566);
				simpleStatement();
				setState(567);
				match(SEMICOLON);
				setState(568);
				expression(0);
				setState(569);
				match(LBRACE);
				setState(570);
				expressionCaseClauseList(0);
				setState(571);
				match(RBRACE);
				}
				break;
			case 2:
				_localctx = new SwitchStmtExpressionASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(573);
				match(SWITCH);
				setState(574);
				expression(0);
				setState(575);
				match(LBRACE);
				setState(576);
				expressionCaseClauseList(0);
				setState(577);
				match(RBRACE);
				}
				break;
			case 3:
				_localctx = new SwitchStmtSimpleStatementBlockASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(579);
				match(SWITCH);
				setState(580);
				simpleStatement();
				setState(581);
				match(SEMICOLON);
				setState(582);
				match(LBRACE);
				setState(583);
				expressionCaseClauseList(0);
				setState(584);
				match(RBRACE);
				}
				break;
			case 4:
				_localctx = new SwitchStmtBlockASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(586);
				match(SWITCH);
				setState(587);
				match(LBRACE);
				setState(588);
				expressionCaseClauseList(0);
				setState(589);
				match(RBRACE);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionCaseClauseListContext extends ParserRuleContext {
		public ExpressionCaseClauseListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionCaseClauseList; }
	 
		public ExpressionCaseClauseListContext() { }
		public void copyFrom(ExpressionCaseClauseListContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionCaseClauseListEmptyASTContext extends ExpressionCaseClauseListContext {
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public ExpressionCaseClauseListEmptyASTContext(ExpressionCaseClauseListContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionCaseClauseListEmptyAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionCaseClauseListEmptyAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionCaseClauseListASTContext extends ExpressionCaseClauseListContext {
		public ExpressionCaseClauseListContext expressionCaseClauseList() {
			return getRuleContext(ExpressionCaseClauseListContext.class,0);
		}
		public ExpressionCaseClauseContext expressionCaseClause() {
			return getRuleContext(ExpressionCaseClauseContext.class,0);
		}
		public ExpressionCaseClauseListASTContext(ExpressionCaseClauseListContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionCaseClauseListAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionCaseClauseListAST(this);
		}
	}

	public final ExpressionCaseClauseListContext expressionCaseClauseList() throws RecognitionException {
		return expressionCaseClauseList(0);
	}

	private ExpressionCaseClauseListContext expressionCaseClauseList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionCaseClauseListContext _localctx = new ExpressionCaseClauseListContext(_ctx, _parentState);
		ExpressionCaseClauseListContext _prevctx = _localctx;
		int _startState = 74;
		enterRecursionRule(_localctx, 74, RULE_expressionCaseClauseList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new ExpressionCaseClauseListEmptyASTContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(594);
			epsilon();
			}
			_ctx.stop = _input.LT(-1);
			setState(600);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExpressionCaseClauseListASTContext(new ExpressionCaseClauseListContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_expressionCaseClauseList);
					setState(596);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(597);
					expressionCaseClause();
					}
					} 
				}
				setState(602);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionCaseClauseContext extends ParserRuleContext {
		public ExpressionCaseClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionCaseClause; }
	 
		public ExpressionCaseClauseContext() { }
		public void copyFrom(ExpressionCaseClauseContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionCaseClauseASTContext extends ExpressionCaseClauseContext {
		public ExpressionSwitchCaseContext expressionSwitchCase() {
			return getRuleContext(ExpressionSwitchCaseContext.class,0);
		}
		public TerminalNode COLON() { return getToken(miniGoParser.COLON, 0); }
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public ExpressionCaseClauseASTContext(ExpressionCaseClauseContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionCaseClauseAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionCaseClauseAST(this);
		}
	}

	public final ExpressionCaseClauseContext expressionCaseClause() throws RecognitionException {
		ExpressionCaseClauseContext _localctx = new ExpressionCaseClauseContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_expressionCaseClause);
		try {
			_localctx = new ExpressionCaseClauseASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(603);
			expressionSwitchCase();
			setState(604);
			match(COLON);
			setState(605);
			statementList();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionSwitchCaseContext extends ParserRuleContext {
		public ExpressionSwitchCaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionSwitchCase; }
	 
		public ExpressionSwitchCaseContext() { }
		public void copyFrom(ExpressionSwitchCaseContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionSwitchCaseASTContext extends ExpressionSwitchCaseContext {
		public TerminalNode CASE() { return getToken(miniGoParser.CASE, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public ExpressionSwitchCaseASTContext(ExpressionSwitchCaseContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionSwitchCaseAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionSwitchCaseAST(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionSwitchDefaultASTContext extends ExpressionSwitchCaseContext {
		public TerminalNode DEFAULT() { return getToken(miniGoParser.DEFAULT, 0); }
		public ExpressionSwitchDefaultASTContext(ExpressionSwitchCaseContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterExpressionSwitchDefaultAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitExpressionSwitchDefaultAST(this);
		}
	}

	public final ExpressionSwitchCaseContext expressionSwitchCase() throws RecognitionException {
		ExpressionSwitchCaseContext _localctx = new ExpressionSwitchCaseContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_expressionSwitchCase);
		try {
			setState(610);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CASE:
				_localctx = new ExpressionSwitchCaseASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(607);
				match(CASE);
				setState(608);
				expressionList();
				}
				break;
			case DEFAULT:
				_localctx = new ExpressionSwitchDefaultASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(609);
				match(DEFAULT);
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

	@SuppressWarnings("CheckReturnValue")
	public static class EpsilonContext extends ParserRuleContext {
		public EpsilonContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_epsilon; }
	 
		public EpsilonContext() { }
		public void copyFrom(EpsilonContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class EpsilonASTContext extends EpsilonContext {
		public EpsilonASTContext(EpsilonContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).enterEpsilonAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof miniGoParserListener ) ((miniGoParserListener)listener).exitEpsilonAST(this);
		}
	}

	public final EpsilonContext epsilon() throws RecognitionException {
		EpsilonContext _localctx = new EpsilonContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_epsilon);
		try {
			_localctx = new EpsilonASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
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
		case 18:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 20:
			return primaryExpression_sempred((PrimaryExpressionContext)_localctx, predIndex);
		case 37:
			return expressionCaseClauseList_sempred((ExpressionCaseClauseListContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 23);
		case 1:
			return precpred(_ctx, 22);
		case 2:
			return precpred(_ctx, 21);
		case 3:
			return precpred(_ctx, 20);
		case 4:
			return precpred(_ctx, 19);
		case 5:
			return precpred(_ctx, 18);
		case 6:
			return precpred(_ctx, 17);
		case 7:
			return precpred(_ctx, 16);
		case 8:
			return precpred(_ctx, 15);
		case 9:
			return precpred(_ctx, 14);
		case 10:
			return precpred(_ctx, 13);
		case 11:
			return precpred(_ctx, 12);
		case 12:
			return precpred(_ctx, 11);
		case 13:
			return precpred(_ctx, 10);
		case 14:
			return precpred(_ctx, 9);
		case 15:
			return precpred(_ctx, 8);
		case 16:
			return precpred(_ctx, 7);
		case 17:
			return precpred(_ctx, 6);
		case 18:
			return precpred(_ctx, 5);
		}
		return true;
	}
	private boolean primaryExpression_sempred(PrimaryExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 19:
			return precpred(_ctx, 6);
		case 20:
			return precpred(_ctx, 5);
		case 21:
			return precpred(_ctx, 4);
		}
		return true;
	}
	private boolean expressionCaseClauseList_sempred(ExpressionCaseClauseListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 22:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001H\u0267\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0005\u0001[\b\u0001\n\u0001\f\u0001^\t"+
		"\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0003\u0002n\b\u0002\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0005\u0003u\b\u0003\n\u0003"+
		"\f\u0003x\t\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003"+
		"\u0004\u0084\b\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0003\u0006\u0097\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0005\u0007\u009e\b\u0007\n\u0007\f\u0007\u00a1\t\u0007"+
		"\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0003\n\u00af\b\n\u0001\n\u0001\n\u0001\n\u0003"+
		"\n\u00b4\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u00b9\b\u000b"+
		"\n\u000b\f\u000b\u00bc\t\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f"+
		"\u0001\f\u0001\f\u0001\f\u0003\f\u00c6\b\f\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u00d5\b\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0005\u0010\u00de\b\u0010\n\u0010\f\u0010\u00e1\t\u0010\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0005\u0011\u00e6\b\u0011\n\u0011\f\u0011\u00e9\t\u0011"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u00f5\b\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0005\u0012\u0130\b\u0012\n\u0012"+
		"\f\u0012\u0133\t\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0005\u0013"+
		"\u0138\b\u0013\n\u0013\f\u0013\u013b\t\u0013\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0003\u0014\u0142\b\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u014a"+
		"\b\u0014\n\u0014\f\u0014\u014d\t\u0014\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u0155\b\u0015\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u015c\b\u0016"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0003\u0018\u0167\b\u0018\u0001\u0019"+
		"\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c"+
		"\u0001\u001c\u0001\u001d\u0005\u001d\u017e\b\u001d\n\u001d\f\u001d\u0181"+
		"\t\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u018b\b\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0003"+
		"\u001f\u0194\b\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0003\u001f\u019c\b\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u01b5\b\u001f\u0001 \u0001"+
		" \u0001 \u0001 \u0001 \u0003 \u01bc\b \u0001 \u0001 \u0001 \u0001 \u0001"+
		" \u0003 \u01c3\b \u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0003!\u01f5\b!\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0003\"\u021d\b\"\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0003#\u0234"+
		"\b#\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0003$\u0250\b$\u0001%\u0001"+
		"%\u0001%\u0001%\u0001%\u0005%\u0257\b%\n%\f%\u025a\t%\u0001&\u0001&\u0001"+
		"&\u0001&\u0001\'\u0001\'\u0001\'\u0003\'\u0263\b\'\u0001(\u0001(\u0001"+
		"(\u0000\u0003$(J)\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014"+
		"\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDFHJLNP\u0000\u0000\u02a3"+
		"\u0000R\u0001\u0000\u0000\u0000\u0002\\\u0001\u0000\u0000\u0000\u0004"+
		"m\u0001\u0000\u0000\u0000\u0006o\u0001\u0000\u0000\u0000\b\u0083\u0001"+
		"\u0000\u0000\u0000\n\u0085\u0001\u0000\u0000\u0000\f\u0096\u0001\u0000"+
		"\u0000\u0000\u000e\u0098\u0001\u0000\u0000\u0000\u0010\u00a2\u0001\u0000"+
		"\u0000\u0000\u0012\u00a5\u0001\u0000\u0000\u0000\u0014\u00a9\u0001\u0000"+
		"\u0000\u0000\u0016\u00b5\u0001\u0000\u0000\u0000\u0018\u00c5\u0001\u0000"+
		"\u0000\u0000\u001a\u00c7\u0001\u0000\u0000\u0000\u001c\u00cb\u0001\u0000"+
		"\u0000\u0000\u001e\u00d0\u0001\u0000\u0000\u0000 \u00d8\u0001\u0000\u0000"+
		"\u0000\"\u00e2\u0001\u0000\u0000\u0000$\u00f4\u0001\u0000\u0000\u0000"+
		"&\u0134\u0001\u0000\u0000\u0000(\u0141\u0001\u0000\u0000\u0000*\u0154"+
		"\u0001\u0000\u0000\u0000,\u015b\u0001\u0000\u0000\u0000.\u015d\u0001\u0000"+
		"\u0000\u00000\u0166\u0001\u0000\u0000\u00002\u0168\u0001\u0000\u0000\u0000"+
		"4\u016b\u0001\u0000\u0000\u00006\u0172\u0001\u0000\u0000\u00008\u0177"+
		"\u0001\u0000\u0000\u0000:\u017f\u0001\u0000\u0000\u0000<\u0182\u0001\u0000"+
		"\u0000\u0000>\u01b4\u0001\u0000\u0000\u0000@\u01c2\u0001\u0000\u0000\u0000"+
		"B\u01f4\u0001\u0000\u0000\u0000D\u021c\u0001\u0000\u0000\u0000F\u0233"+
		"\u0001\u0000\u0000\u0000H\u024f\u0001\u0000\u0000\u0000J\u0251\u0001\u0000"+
		"\u0000\u0000L\u025b\u0001\u0000\u0000\u0000N\u0262\u0001\u0000\u0000\u0000"+
		"P\u0264\u0001\u0000\u0000\u0000RS\u0005\u0006\u0000\u0000ST\u0005\u0019"+
		"\u0000\u0000TU\u0005=\u0000\u0000UV\u0003\u0002\u0001\u0000V\u0001\u0001"+
		"\u0000\u0000\u0000W[\u0003\u0004\u0002\u0000X[\u0003\f\u0006\u0000Y[\u0003"+
		"\u0012\t\u0000ZW\u0001\u0000\u0000\u0000ZX\u0001\u0000\u0000\u0000ZY\u0001"+
		"\u0000\u0000\u0000[^\u0001\u0000\u0000\u0000\\Z\u0001\u0000\u0000\u0000"+
		"\\]\u0001\u0000\u0000\u0000]\u0003\u0001\u0000\u0000\u0000^\\\u0001\u0000"+
		"\u0000\u0000_`\u0005\u0007\u0000\u0000`a\u0003\b\u0004\u0000ab\u0005="+
		"\u0000\u0000bn\u0001\u0000\u0000\u0000cd\u0005\u0007\u0000\u0000de\u0005"+
		"@\u0000\u0000ef\u0003\u0006\u0003\u0000fg\u0005A\u0000\u0000gh\u0005="+
		"\u0000\u0000hn\u0001\u0000\u0000\u0000ij\u0005\u0007\u0000\u0000jk\u0005"+
		"@\u0000\u0000kl\u0005A\u0000\u0000ln\u0005=\u0000\u0000m_\u0001\u0000"+
		"\u0000\u0000mc\u0001\u0000\u0000\u0000mi\u0001\u0000\u0000\u0000n\u0005"+
		"\u0001\u0000\u0000\u0000op\u0003\b\u0004\u0000pv\u0005=\u0000\u0000qr"+
		"\u0003\b\u0004\u0000rs\u0005=\u0000\u0000su\u0001\u0000\u0000\u0000tq"+
		"\u0001\u0000\u0000\u0000ux\u0001\u0000\u0000\u0000vt\u0001\u0000\u0000"+
		"\u0000vw\u0001\u0000\u0000\u0000w\u0007\u0001\u0000\u0000\u0000xv\u0001"+
		"\u0000\u0000\u0000yz\u0003\"\u0011\u0000z{\u0003\u0018\f\u0000{|\u0005"+
		"0\u0000\u0000|}\u0003&\u0013\u0000}\u0084\u0001\u0000\u0000\u0000~\u007f"+
		"\u0003\"\u0011\u0000\u007f\u0080\u00050\u0000\u0000\u0080\u0081\u0003"+
		"&\u0013\u0000\u0081\u0084\u0001\u0000\u0000\u0000\u0082\u0084\u0003\n"+
		"\u0005\u0000\u0083y\u0001\u0000\u0000\u0000\u0083~\u0001\u0000\u0000\u0000"+
		"\u0083\u0082\u0001\u0000\u0000\u0000\u0084\t\u0001\u0000\u0000\u0000\u0085"+
		"\u0086\u0003\"\u0011\u0000\u0086\u0087\u0003\u0018\f\u0000\u0087\u000b"+
		"\u0001\u0000\u0000\u0000\u0088\u0089\u0005\b\u0000\u0000\u0089\u008a\u0003"+
		"\u0010\b\u0000\u008a\u008b\u0005=\u0000\u0000\u008b\u0097\u0001\u0000"+
		"\u0000\u0000\u008c\u008d\u0005\b\u0000\u0000\u008d\u008e\u0005@\u0000"+
		"\u0000\u008e\u008f\u0003\u000e\u0007\u0000\u008f\u0090\u0005A\u0000\u0000"+
		"\u0090\u0091\u0005=\u0000\u0000\u0091\u0097\u0001\u0000\u0000\u0000\u0092"+
		"\u0093\u0005\b\u0000\u0000\u0093\u0094\u0005@\u0000\u0000\u0094\u0095"+
		"\u0005A\u0000\u0000\u0095\u0097\u0005=\u0000\u0000\u0096\u0088\u0001\u0000"+
		"\u0000\u0000\u0096\u008c\u0001\u0000\u0000\u0000\u0096\u0092\u0001\u0000"+
		"\u0000\u0000\u0097\r\u0001\u0000\u0000\u0000\u0098\u0099\u0003\u0010\b"+
		"\u0000\u0099\u009f\u0005=\u0000\u0000\u009a\u009b\u0003\u0010\b\u0000"+
		"\u009b\u009c\u0005=\u0000\u0000\u009c\u009e\u0001\u0000\u0000\u0000\u009d"+
		"\u009a\u0001\u0000\u0000\u0000\u009e\u00a1\u0001\u0000\u0000\u0000\u009f"+
		"\u009d\u0001\u0000\u0000\u0000\u009f\u00a0\u0001\u0000\u0000\u0000\u00a0"+
		"\u000f\u0001\u0000\u0000\u0000\u00a1\u009f\u0001\u0000\u0000\u0000\u00a2"+
		"\u00a3\u0005\u0019\u0000\u0000\u00a3\u00a4\u0003\u0018\f\u0000\u00a4\u0011"+
		"\u0001\u0000\u0000\u0000\u00a5\u00a6\u0003\u0014\n\u0000\u00a6\u00a7\u0003"+
		"<\u001e\u0000\u00a7\u00a8\u0005=\u0000\u0000\u00a8\u0013\u0001\u0000\u0000"+
		"\u0000\u00a9\u00aa\u0005\t\u0000\u0000\u00aa\u00ab\u0005\u0019\u0000\u0000"+
		"\u00ab\u00ae\u0005@\u0000\u0000\u00ac\u00af\u0003\u0016\u000b\u0000\u00ad"+
		"\u00af\u0003P(\u0000\u00ae\u00ac\u0001\u0000\u0000\u0000\u00ae\u00ad\u0001"+
		"\u0000\u0000\u0000\u00af\u00b0\u0001\u0000\u0000\u0000\u00b0\u00b3\u0005"+
		"A\u0000\u0000\u00b1\u00b4\u0003\u0018\f\u0000\u00b2\u00b4\u0003P(\u0000"+
		"\u00b3\u00b1\u0001\u0000\u0000\u0000\u00b3\u00b2\u0001\u0000\u0000\u0000"+
		"\u00b4\u0015\u0001\u0000\u0000\u0000\u00b5\u00ba\u0003\n\u0005\u0000\u00b6"+
		"\u00b7\u0005>\u0000\u0000\u00b7\u00b9\u0003\n\u0005\u0000\u00b8\u00b6"+
		"\u0001\u0000\u0000\u0000\u00b9\u00bc\u0001\u0000\u0000\u0000\u00ba\u00b8"+
		"\u0001\u0000\u0000\u0000\u00ba\u00bb\u0001\u0000\u0000\u0000\u00bb\u0017"+
		"\u0001\u0000\u0000\u0000\u00bc\u00ba\u0001\u0000\u0000\u0000\u00bd\u00be"+
		"\u0005@\u0000\u0000\u00be\u00bf\u0003\u0018\f\u0000\u00bf\u00c0\u0005"+
		"A\u0000\u0000\u00c0\u00c6\u0001\u0000\u0000\u0000\u00c1\u00c6\u0005\u0019"+
		"\u0000\u0000\u00c2\u00c6\u0003\u001a\r\u0000\u00c3\u00c6\u0003\u001c\u000e"+
		"\u0000\u00c4\u00c6\u0003\u001e\u000f\u0000\u00c5\u00bd\u0001\u0000\u0000"+
		"\u0000\u00c5\u00c1\u0001\u0000\u0000\u0000\u00c5\u00c2\u0001\u0000\u0000"+
		"\u0000\u00c5\u00c3\u0001\u0000\u0000\u0000\u00c5\u00c4\u0001\u0000\u0000"+
		"\u0000\u00c6\u0019\u0001\u0000\u0000\u0000\u00c7\u00c8\u0005B\u0000\u0000"+
		"\u00c8\u00c9\u0005C\u0000\u0000\u00c9\u00ca\u0003\u0018\f\u0000\u00ca"+
		"\u001b\u0001\u0000\u0000\u0000\u00cb\u00cc\u0005B\u0000\u0000\u00cc\u00cd"+
		"\u0005\u0001\u0000\u0000\u00cd\u00ce\u0005C\u0000\u0000\u00ce\u00cf\u0003"+
		"\u0018\f\u0000\u00cf\u001d\u0001\u0000\u0000\u0000\u00d0\u00d1\u0005\n"+
		"\u0000\u0000\u00d1\u00d4\u0005D\u0000\u0000\u00d2\u00d5\u0003 \u0010\u0000"+
		"\u00d3\u00d5\u0003P(\u0000\u00d4\u00d2\u0001\u0000\u0000\u0000\u00d4\u00d3"+
		"\u0001\u0000\u0000\u0000\u00d5\u00d6\u0001\u0000\u0000\u0000\u00d6\u00d7"+
		"\u0005E\u0000\u0000\u00d7\u001f\u0001\u0000\u0000\u0000\u00d8\u00d9\u0003"+
		"\n\u0005\u0000\u00d9\u00df\u0005=\u0000\u0000\u00da\u00db\u0003\n\u0005"+
		"\u0000\u00db\u00dc\u0005=\u0000\u0000\u00dc\u00de\u0001\u0000\u0000\u0000"+
		"\u00dd\u00da\u0001\u0000\u0000\u0000\u00de\u00e1\u0001\u0000\u0000\u0000"+
		"\u00df\u00dd\u0001\u0000\u0000\u0000\u00df\u00e0\u0001\u0000\u0000\u0000"+
		"\u00e0!\u0001\u0000\u0000\u0000\u00e1\u00df\u0001\u0000\u0000\u0000\u00e2"+
		"\u00e7\u0005\u0019\u0000\u0000\u00e3\u00e4\u0005>\u0000\u0000\u00e4\u00e6"+
		"\u0005\u0019\u0000\u0000\u00e5\u00e3\u0001\u0000\u0000\u0000\u00e6\u00e9"+
		"\u0001\u0000\u0000\u0000\u00e7\u00e5\u0001\u0000\u0000\u0000\u00e7\u00e8"+
		"\u0001\u0000\u0000\u0000\u00e8#\u0001\u0000\u0000\u0000\u00e9\u00e7\u0001"+
		"\u0000\u0000\u0000\u00ea\u00eb\u0006\u0012\uffff\uffff\u0000\u00eb\u00f5"+
		"\u0003(\u0014\u0000\u00ec\u00ed\u0005\u001a\u0000\u0000\u00ed\u00f5\u0003"+
		"$\u0012\u0004\u00ee\u00ef\u0005\u001b\u0000\u0000\u00ef\u00f5\u0003$\u0012"+
		"\u0003\u00f0\u00f1\u0005\'\u0000\u0000\u00f1\u00f5\u0003$\u0012\u0002"+
		"\u00f2\u00f3\u0005*\u0000\u0000\u00f3\u00f5\u0003$\u0012\u0001\u00f4\u00ea"+
		"\u0001\u0000\u0000\u0000\u00f4\u00ec\u0001\u0000\u0000\u0000\u00f4\u00ee"+
		"\u0001\u0000\u0000\u0000\u00f4\u00f0\u0001\u0000\u0000\u0000\u00f4\u00f2"+
		"\u0001\u0000\u0000\u0000\u00f5\u0131\u0001\u0000\u0000\u0000\u00f6\u00f7"+
		"\n\u0017\u0000\u0000\u00f7\u00f8\u0005\u001c\u0000\u0000\u00f8\u0130\u0003"+
		"$\u0012\u0018\u00f9\u00fa\n\u0016\u0000\u0000\u00fa\u00fb\u0005\u001d"+
		"\u0000\u0000\u00fb\u0130\u0003$\u0012\u0017\u00fc\u00fd\n\u0015\u0000"+
		"\u0000\u00fd\u00fe\u0005\u001e\u0000\u0000\u00fe\u0130\u0003$\u0012\u0016"+
		"\u00ff\u0100\n\u0014\u0000\u0000\u0100\u0101\u0005,\u0000\u0000\u0101"+
		"\u0130\u0003$\u0012\u0015\u0102\u0103\n\u0013\u0000\u0000\u0103\u0104"+
		"\u0005-\u0000\u0000\u0104\u0130\u0003$\u0012\u0014\u0105\u0106\n\u0012"+
		"\u0000\u0000\u0106\u0107\u0005(\u0000\u0000\u0107\u0130\u0003$\u0012\u0013"+
		"\u0108\u0109\n\u0011\u0000\u0000\u0109\u010a\u0005+\u0000\u0000\u010a"+
		"\u0130\u0003$\u0012\u0012\u010b\u010c\n\u0010\u0000\u0000\u010c\u010d"+
		"\u0005\u001a\u0000\u0000\u010d\u0130\u0003$\u0012\u0011\u010e\u010f\n"+
		"\u000f\u0000\u0000\u010f\u0110\u0005\u001b\u0000\u0000\u0110\u0130\u0003"+
		"$\u0012\u0010\u0111\u0112\n\u000e\u0000\u0000\u0112\u0113\u0005)\u0000"+
		"\u0000\u0113\u0130\u0003$\u0012\u000f\u0114\u0115\n\r\u0000\u0000\u0115"+
		"\u0116\u0005*\u0000\u0000\u0116\u0130\u0003$\u0012\u000e\u0117\u0118\n"+
		"\f\u0000\u0000\u0118\u0119\u0005#\u0000\u0000\u0119\u0130\u0003$\u0012"+
		"\r\u011a\u011b\n\u000b\u0000\u0000\u011b\u011c\u0005$\u0000\u0000\u011c"+
		"\u0130\u0003$\u0012\f\u011d\u011e\n\n\u0000\u0000\u011e\u011f\u0005\u001f"+
		"\u0000\u0000\u011f\u0130\u0003$\u0012\u000b\u0120\u0121\n\t\u0000\u0000"+
		"\u0121\u0122\u0005 \u0000\u0000\u0122\u0130\u0003$\u0012\n\u0123\u0124"+
		"\n\b\u0000\u0000\u0124\u0125\u0005!\u0000\u0000\u0125\u0130\u0003$\u0012"+
		"\t\u0126\u0127\n\u0007\u0000\u0000\u0127\u0128\u0005\"\u0000\u0000\u0128"+
		"\u0130\u0003$\u0012\b\u0129\u012a\n\u0006\u0000\u0000\u012a\u012b\u0005"+
		"%\u0000\u0000\u012b\u0130\u0003$\u0012\u0007\u012c\u012d\n\u0005\u0000"+
		"\u0000\u012d\u012e\u0005&\u0000\u0000\u012e\u0130\u0003$\u0012\u0006\u012f"+
		"\u00f6\u0001\u0000\u0000\u0000\u012f\u00f9\u0001\u0000\u0000\u0000\u012f"+
		"\u00fc\u0001\u0000\u0000\u0000\u012f\u00ff\u0001\u0000\u0000\u0000\u012f"+
		"\u0102\u0001\u0000\u0000\u0000\u012f\u0105\u0001\u0000\u0000\u0000\u012f"+
		"\u0108\u0001\u0000\u0000\u0000\u012f\u010b\u0001\u0000\u0000\u0000\u012f"+
		"\u010e\u0001\u0000\u0000\u0000\u012f\u0111\u0001\u0000\u0000\u0000\u012f"+
		"\u0114\u0001\u0000\u0000\u0000\u012f\u0117\u0001\u0000\u0000\u0000\u012f"+
		"\u011a\u0001\u0000\u0000\u0000\u012f\u011d\u0001\u0000\u0000\u0000\u012f"+
		"\u0120\u0001\u0000\u0000\u0000\u012f\u0123\u0001\u0000\u0000\u0000\u012f"+
		"\u0126\u0001\u0000\u0000\u0000\u012f\u0129\u0001\u0000\u0000\u0000\u012f"+
		"\u012c\u0001\u0000\u0000\u0000\u0130\u0133\u0001\u0000\u0000\u0000\u0131"+
		"\u012f\u0001\u0000\u0000\u0000\u0131\u0132\u0001\u0000\u0000\u0000\u0132"+
		"%\u0001\u0000\u0000\u0000\u0133\u0131\u0001\u0000\u0000\u0000\u0134\u0139"+
		"\u0003$\u0012\u0000\u0135\u0136\u0005>\u0000\u0000\u0136\u0138\u0003$"+
		"\u0012\u0000\u0137\u0135\u0001\u0000\u0000\u0000\u0138\u013b\u0001\u0000"+
		"\u0000\u0000\u0139\u0137\u0001\u0000\u0000\u0000\u0139\u013a\u0001\u0000"+
		"\u0000\u0000\u013a\'\u0001\u0000\u0000\u0000\u013b\u0139\u0001\u0000\u0000"+
		"\u0000\u013c\u013d\u0006\u0014\uffff\uffff\u0000\u013d\u0142\u0003*\u0015"+
		"\u0000\u013e\u0142\u00034\u001a\u0000\u013f\u0142\u00036\u001b\u0000\u0140"+
		"\u0142\u00038\u001c\u0000\u0141\u013c\u0001\u0000\u0000\u0000\u0141\u013e"+
		"\u0001\u0000\u0000\u0000\u0141\u013f\u0001\u0000\u0000\u0000\u0141\u0140"+
		"\u0001\u0000\u0000\u0000\u0142\u014b\u0001\u0000\u0000\u0000\u0143\u0144"+
		"\n\u0006\u0000\u0000\u0144\u014a\u00032\u0019\u0000\u0145\u0146\n\u0005"+
		"\u0000\u0000\u0146\u014a\u0003.\u0017\u0000\u0147\u0148\n\u0004\u0000"+
		"\u0000\u0148\u014a\u00030\u0018\u0000\u0149\u0143\u0001\u0000\u0000\u0000"+
		"\u0149\u0145\u0001\u0000\u0000\u0000\u0149\u0147\u0001\u0000\u0000\u0000"+
		"\u014a\u014d\u0001\u0000\u0000\u0000\u014b\u0149\u0001\u0000\u0000\u0000"+
		"\u014b\u014c\u0001\u0000\u0000\u0000\u014c)\u0001\u0000\u0000\u0000\u014d"+
		"\u014b\u0001\u0000\u0000\u0000\u014e\u0155\u0003,\u0016\u0000\u014f\u0155"+
		"\u0005\u0019\u0000\u0000\u0150\u0151\u0005@\u0000\u0000\u0151\u0152\u0003"+
		"$\u0012\u0000\u0152\u0153\u0005A\u0000\u0000\u0153\u0155\u0001\u0000\u0000"+
		"\u0000\u0154\u014e\u0001\u0000\u0000\u0000\u0154\u014f\u0001\u0000\u0000"+
		"\u0000\u0154\u0150\u0001\u0000\u0000\u0000\u0155+\u0001\u0000\u0000\u0000"+
		"\u0156\u015c\u0005\u0001\u0000\u0000\u0157\u015c\u0005\u0002\u0000\u0000"+
		"\u0158\u015c\u0005\u0003\u0000\u0000\u0159\u015c\u0005\u0004\u0000\u0000"+
		"\u015a\u015c\u0005\u0005\u0000\u0000\u015b\u0156\u0001\u0000\u0000\u0000"+
		"\u015b\u0157\u0001\u0000\u0000\u0000\u015b\u0158\u0001\u0000\u0000\u0000"+
		"\u015b\u0159\u0001\u0000\u0000\u0000\u015b\u015a\u0001\u0000\u0000\u0000"+
		"\u015c-\u0001\u0000\u0000\u0000\u015d\u015e\u0005B\u0000\u0000\u015e\u015f"+
		"\u0003$\u0012\u0000\u015f\u0160\u0005C\u0000\u0000\u0160/\u0001\u0000"+
		"\u0000\u0000\u0161\u0162\u0005@\u0000\u0000\u0162\u0167\u0003&\u0013\u0000"+
		"\u0163\u0164\u0003P(\u0000\u0164\u0165\u0005A\u0000\u0000\u0165\u0167"+
		"\u0001\u0000\u0000\u0000\u0166\u0161\u0001\u0000\u0000\u0000\u0166\u0163"+
		"\u0001\u0000\u0000\u0000\u01671\u0001\u0000\u0000\u0000\u0168\u0169\u0005"+
		"?\u0000\u0000\u0169\u016a\u0005\u0019\u0000\u0000\u016a3\u0001\u0000\u0000"+
		"\u0000\u016b\u016c\u0005\u0012\u0000\u0000\u016c\u016d\u0005@\u0000\u0000"+
		"\u016d\u016e\u0003$\u0012\u0000\u016e\u016f\u0005>\u0000\u0000\u016f\u0170"+
		"\u0003$\u0012\u0000\u0170\u0171\u0005A\u0000\u0000\u01715\u0001\u0000"+
		"\u0000\u0000\u0172\u0173\u0005\u000b\u0000\u0000\u0173\u0174\u0005@\u0000"+
		"\u0000\u0174\u0175\u0003$\u0012\u0000\u0175\u0176\u0005A\u0000\u0000\u0176"+
		"7\u0001\u0000\u0000\u0000\u0177\u0178\u0005\f\u0000\u0000\u0178\u0179"+
		"\u0005@\u0000\u0000\u0179\u017a\u0003$\u0012\u0000\u017a\u017b\u0005A"+
		"\u0000\u0000\u017b9\u0001\u0000\u0000\u0000\u017c\u017e\u0003>\u001f\u0000"+
		"\u017d\u017c\u0001\u0000\u0000\u0000\u017e\u0181\u0001\u0000\u0000\u0000"+
		"\u017f\u017d\u0001\u0000\u0000\u0000\u017f\u0180\u0001\u0000\u0000\u0000"+
		"\u0180;\u0001\u0000\u0000\u0000\u0181\u017f\u0001\u0000\u0000\u0000\u0182"+
		"\u0183\u0005D\u0000\u0000\u0183\u0184\u0003:\u001d\u0000\u0184\u0185\u0005"+
		"E\u0000\u0000\u0185=\u0001\u0000\u0000\u0000\u0186\u0187\u0005\r\u0000"+
		"\u0000\u0187\u018a\u0005@\u0000\u0000\u0188\u018b\u0003&\u0013\u0000\u0189"+
		"\u018b\u0003P(\u0000\u018a\u0188\u0001\u0000\u0000\u0000\u018a\u0189\u0001"+
		"\u0000\u0000\u0000\u018b\u018c\u0001\u0000\u0000\u0000\u018c\u018d\u0005"+
		"A\u0000\u0000\u018d\u018e\u0005=\u0000\u0000\u018e\u01b5\u0001\u0000\u0000"+
		"\u0000\u018f\u0190\u0005\u000e\u0000\u0000\u0190\u0193\u0005@\u0000\u0000"+
		"\u0191\u0194\u0003&\u0013\u0000\u0192\u0194\u0003P(\u0000\u0193\u0191"+
		"\u0001\u0000\u0000\u0000\u0193\u0192\u0001\u0000\u0000\u0000\u0194\u0195"+
		"\u0001\u0000\u0000\u0000\u0195\u0196\u0005A\u0000\u0000\u0196\u0197\u0005"+
		"=\u0000\u0000\u0197\u01b5\u0001\u0000\u0000\u0000\u0198\u019b\u0005\u000f"+
		"\u0000\u0000\u0199\u019c\u0003$\u0012\u0000\u019a\u019c\u0003P(\u0000"+
		"\u019b\u0199\u0001\u0000\u0000\u0000\u019b\u019a\u0001\u0000\u0000\u0000"+
		"\u019c\u019d\u0001\u0000\u0000\u0000\u019d\u019e\u0005=\u0000\u0000\u019e"+
		"\u01b5\u0001\u0000\u0000\u0000\u019f\u01a0\u0005\u0010\u0000\u0000\u01a0"+
		"\u01b5\u0005=\u0000\u0000\u01a1\u01a2\u0005\u0011\u0000\u0000\u01a2\u01b5"+
		"\u0005=\u0000\u0000\u01a3\u01a4\u0003@ \u0000\u01a4\u01a5\u0005=\u0000"+
		"\u0000\u01a5\u01b5\u0001\u0000\u0000\u0000\u01a6\u01a7\u0003<\u001e\u0000"+
		"\u01a7\u01a8\u0005=\u0000\u0000\u01a8\u01b5\u0001\u0000\u0000\u0000\u01a9"+
		"\u01aa\u0003H$\u0000\u01aa\u01ab\u0005=\u0000\u0000\u01ab\u01b5\u0001"+
		"\u0000\u0000\u0000\u01ac\u01ad\u0003D\"\u0000\u01ad\u01ae\u0005=\u0000"+
		"\u0000\u01ae\u01b5\u0001\u0000\u0000\u0000\u01af\u01b0\u0003F#\u0000\u01b0"+
		"\u01b1\u0005=\u0000\u0000\u01b1\u01b5\u0001\u0000\u0000\u0000\u01b2\u01b5"+
		"\u0003\f\u0006\u0000\u01b3\u01b5\u0003\u0004\u0002\u0000\u01b4\u0186\u0001"+
		"\u0000\u0000\u0000\u01b4\u018f\u0001\u0000\u0000\u0000\u01b4\u0198\u0001"+
		"\u0000\u0000\u0000\u01b4\u019f\u0001\u0000\u0000\u0000\u01b4\u01a1\u0001"+
		"\u0000\u0000\u0000\u01b4\u01a3\u0001\u0000\u0000\u0000\u01b4\u01a6\u0001"+
		"\u0000\u0000\u0000\u01b4\u01a9\u0001\u0000\u0000\u0000\u01b4\u01ac\u0001"+
		"\u0000\u0000\u0000\u01b4\u01af\u0001\u0000\u0000\u0000\u01b4\u01b2\u0001"+
		"\u0000\u0000\u0000\u01b4\u01b3\u0001\u0000\u0000\u0000\u01b5?\u0001\u0000"+
		"\u0000\u0000\u01b6\u01c3\u0003P(\u0000\u01b7\u01bb\u0003$\u0012\u0000"+
		"\u01b8\u01bc\u0005.\u0000\u0000\u01b9\u01bc\u0005/\u0000\u0000\u01ba\u01bc"+
		"\u0003P(\u0000\u01bb\u01b8\u0001\u0000\u0000\u0000\u01bb\u01b9\u0001\u0000"+
		"\u0000\u0000\u01bb\u01ba\u0001\u0000\u0000\u0000\u01bc\u01c3\u0001\u0000"+
		"\u0000\u0000\u01bd\u01c3\u0003B!\u0000\u01be\u01bf\u0003&\u0013\u0000"+
		"\u01bf\u01c0\u00050\u0000\u0000\u01c0\u01c1\u0003&\u0013\u0000\u01c1\u01c3"+
		"\u0001\u0000\u0000\u0000\u01c2\u01b6\u0001\u0000\u0000\u0000\u01c2\u01b7"+
		"\u0001\u0000\u0000\u0000\u01c2\u01bd\u0001\u0000\u0000\u0000\u01c2\u01be"+
		"\u0001\u0000\u0000\u0000\u01c3A\u0001\u0000\u0000\u0000\u01c4\u01c5\u0003"+
		"&\u0013\u0000\u01c5\u01c6\u00050\u0000\u0000\u01c6\u01c7\u0003&\u0013"+
		"\u0000\u01c7\u01f5\u0001\u0000\u0000\u0000\u01c8\u01c9\u0003$\u0012\u0000"+
		"\u01c9\u01ca\u00051\u0000\u0000\u01ca\u01cb\u0003$\u0012\u0000\u01cb\u01f5"+
		"\u0001\u0000\u0000\u0000\u01cc\u01cd\u0003$\u0012\u0000\u01cd\u01ce\u0005"+
		"2\u0000\u0000\u01ce\u01cf\u0003$\u0012\u0000\u01cf\u01f5\u0001\u0000\u0000"+
		"\u0000\u01d0\u01d1\u0003$\u0012\u0000\u01d1\u01d2\u00056\u0000\u0000\u01d2"+
		"\u01d3\u0003$\u0012\u0000\u01d3\u01f5\u0001\u0000\u0000\u0000\u01d4\u01d5"+
		"\u0003$\u0012\u0000\u01d5\u01d6\u00057\u0000\u0000\u01d6\u01d7\u0003$"+
		"\u0012\u0000\u01d7\u01f5\u0001\u0000\u0000\u0000\u01d8\u01d9\u0003$\u0012"+
		"\u0000\u01d9\u01da\u00053\u0000\u0000\u01da\u01db\u0003$\u0012\u0000\u01db"+
		"\u01f5\u0001\u0000\u0000\u0000\u01dc\u01dd\u0003$\u0012\u0000\u01dd\u01de"+
		"\u00058\u0000\u0000\u01de\u01df\u0003$\u0012\u0000\u01df\u01f5\u0001\u0000"+
		"\u0000\u0000\u01e0\u01e1\u0003$\u0012\u0000\u01e1\u01e2\u00059\u0000\u0000"+
		"\u01e2\u01e3\u0003$\u0012\u0000\u01e3\u01f5\u0001\u0000\u0000\u0000\u01e4"+
		"\u01e5\u0003$\u0012\u0000\u01e5\u01e6\u0005:\u0000\u0000\u01e6\u01e7\u0003"+
		"$\u0012\u0000\u01e7\u01f5\u0001\u0000\u0000\u0000\u01e8\u01e9\u0003$\u0012"+
		"\u0000\u01e9\u01ea\u0005;\u0000\u0000\u01ea\u01eb\u0003$\u0012\u0000\u01eb"+
		"\u01f5\u0001\u0000\u0000\u0000\u01ec\u01ed\u0003$\u0012\u0000\u01ed\u01ee"+
		"\u00055\u0000\u0000\u01ee\u01ef\u0003$\u0012\u0000\u01ef\u01f5\u0001\u0000"+
		"\u0000\u0000\u01f0\u01f1\u0003$\u0012\u0000\u01f1\u01f2\u00054\u0000\u0000"+
		"\u01f2\u01f3\u0003$\u0012\u0000\u01f3\u01f5\u0001\u0000\u0000\u0000\u01f4"+
		"\u01c4\u0001\u0000\u0000\u0000\u01f4\u01c8\u0001\u0000\u0000\u0000\u01f4"+
		"\u01cc\u0001\u0000\u0000\u0000\u01f4\u01d0\u0001\u0000\u0000\u0000\u01f4"+
		"\u01d4\u0001\u0000\u0000\u0000\u01f4\u01d8\u0001\u0000\u0000\u0000\u01f4"+
		"\u01dc\u0001\u0000\u0000\u0000\u01f4\u01e0\u0001\u0000\u0000\u0000\u01f4"+
		"\u01e4\u0001\u0000\u0000\u0000\u01f4\u01e8\u0001\u0000\u0000\u0000\u01f4"+
		"\u01ec\u0001\u0000\u0000\u0000\u01f4\u01f0\u0001\u0000\u0000\u0000\u01f5"+
		"C\u0001\u0000\u0000\u0000\u01f6\u01f7\u0005\u0013\u0000\u0000\u01f7\u01f8"+
		"\u0003$\u0012\u0000\u01f8\u01f9\u0003<\u001e\u0000\u01f9\u021d\u0001\u0000"+
		"\u0000\u0000\u01fa\u01fb\u0005\u0013\u0000\u0000\u01fb\u01fc\u0003$\u0012"+
		"\u0000\u01fc\u01fd\u0003<\u001e\u0000\u01fd\u01fe\u0005\u0014\u0000\u0000"+
		"\u01fe\u01ff\u0003D\"\u0000\u01ff\u021d\u0001\u0000\u0000\u0000\u0200"+
		"\u0201\u0005\u0013\u0000\u0000\u0201\u0202\u0003$\u0012\u0000\u0202\u0203"+
		"\u0003<\u001e\u0000\u0203\u0204\u0005\u0014\u0000\u0000\u0204\u0205\u0003"+
		"<\u001e\u0000\u0205\u021d\u0001\u0000\u0000\u0000\u0206\u0207\u0005\u0013"+
		"\u0000\u0000\u0207\u0208\u0003@ \u0000\u0208\u0209\u0005=\u0000\u0000"+
		"\u0209\u020a\u0003$\u0012\u0000\u020a\u020b\u0003<\u001e\u0000\u020b\u021d"+
		"\u0001\u0000\u0000\u0000\u020c\u020d\u0005\u0013\u0000\u0000\u020d\u020e"+
		"\u0003@ \u0000\u020e\u020f\u0005=\u0000\u0000\u020f\u0210\u0003$\u0012"+
		"\u0000\u0210\u0211\u0003<\u001e\u0000\u0211\u0212\u0005\u0014\u0000\u0000"+
		"\u0212\u0213\u0003D\"\u0000\u0213\u021d\u0001\u0000\u0000\u0000\u0214"+
		"\u0215\u0005\u0013\u0000\u0000\u0215\u0216\u0003@ \u0000\u0216\u0217\u0005"+
		"=\u0000\u0000\u0217\u0218\u0003$\u0012\u0000\u0218\u0219\u0003<\u001e"+
		"\u0000\u0219\u021a\u0005\u0014\u0000\u0000\u021a\u021b\u0003<\u001e\u0000"+
		"\u021b\u021d\u0001\u0000\u0000\u0000\u021c\u01f6\u0001\u0000\u0000\u0000"+
		"\u021c\u01fa\u0001\u0000\u0000\u0000\u021c\u0200\u0001\u0000\u0000\u0000"+
		"\u021c\u0206\u0001\u0000\u0000\u0000\u021c\u020c\u0001\u0000\u0000\u0000"+
		"\u021c\u0214\u0001\u0000\u0000\u0000\u021dE\u0001\u0000\u0000\u0000\u021e"+
		"\u021f\u0005\u0015\u0000\u0000\u021f\u0234\u0003<\u001e\u0000\u0220\u0221"+
		"\u0005\u0015\u0000\u0000\u0221\u0222\u0003$\u0012\u0000\u0222\u0223\u0003"+
		"<\u001e\u0000\u0223\u0234\u0001\u0000\u0000\u0000\u0224\u0225\u0005\u0015"+
		"\u0000\u0000\u0225\u0226\u0003@ \u0000\u0226\u0227\u0005=\u0000\u0000"+
		"\u0227\u0228\u0003$\u0012\u0000\u0228\u0229\u0005=\u0000\u0000\u0229\u022a"+
		"\u0003@ \u0000\u022a\u022b\u0003<\u001e\u0000\u022b\u0234\u0001\u0000"+
		"\u0000\u0000\u022c\u022d\u0005\u0015\u0000\u0000\u022d\u022e\u0003@ \u0000"+
		"\u022e\u022f\u0005=\u0000\u0000\u022f\u0230\u0005=\u0000\u0000\u0230\u0231"+
		"\u0003@ \u0000\u0231\u0232\u0003<\u001e\u0000\u0232\u0234\u0001\u0000"+
		"\u0000\u0000\u0233\u021e\u0001\u0000\u0000\u0000\u0233\u0220\u0001\u0000"+
		"\u0000\u0000\u0233\u0224\u0001\u0000\u0000\u0000\u0233\u022c\u0001\u0000"+
		"\u0000\u0000\u0234G\u0001\u0000\u0000\u0000\u0235\u0236\u0005\u0016\u0000"+
		"\u0000\u0236\u0237\u0003@ \u0000\u0237\u0238\u0005=\u0000\u0000\u0238"+
		"\u0239\u0003$\u0012\u0000\u0239\u023a\u0005D\u0000\u0000\u023a\u023b\u0003"+
		"J%\u0000\u023b\u023c\u0005E\u0000\u0000\u023c\u0250\u0001\u0000\u0000"+
		"\u0000\u023d\u023e\u0005\u0016\u0000\u0000\u023e\u023f\u0003$\u0012\u0000"+
		"\u023f\u0240\u0005D\u0000\u0000\u0240\u0241\u0003J%\u0000\u0241\u0242"+
		"\u0005E\u0000\u0000\u0242\u0250\u0001\u0000\u0000\u0000\u0243\u0244\u0005"+
		"\u0016\u0000\u0000\u0244\u0245\u0003@ \u0000\u0245\u0246\u0005=\u0000"+
		"\u0000\u0246\u0247\u0005D\u0000\u0000\u0247\u0248\u0003J%\u0000\u0248"+
		"\u0249\u0005E\u0000\u0000\u0249\u0250\u0001\u0000\u0000\u0000\u024a\u024b"+
		"\u0005\u0016\u0000\u0000\u024b\u024c\u0005D\u0000\u0000\u024c\u024d\u0003"+
		"J%\u0000\u024d\u024e\u0005E\u0000\u0000\u024e\u0250\u0001\u0000\u0000"+
		"\u0000\u024f\u0235\u0001\u0000\u0000\u0000\u024f\u023d\u0001\u0000\u0000"+
		"\u0000\u024f\u0243\u0001\u0000\u0000\u0000\u024f\u024a\u0001\u0000\u0000"+
		"\u0000\u0250I\u0001\u0000\u0000\u0000\u0251\u0252\u0006%\uffff\uffff\u0000"+
		"\u0252\u0253\u0003P(\u0000\u0253\u0258\u0001\u0000\u0000\u0000\u0254\u0255"+
		"\n\u0001\u0000\u0000\u0255\u0257\u0003L&\u0000\u0256\u0254\u0001\u0000"+
		"\u0000\u0000\u0257\u025a\u0001\u0000\u0000\u0000\u0258\u0256\u0001\u0000"+
		"\u0000\u0000\u0258\u0259\u0001\u0000\u0000\u0000\u0259K\u0001\u0000\u0000"+
		"\u0000\u025a\u0258\u0001\u0000\u0000\u0000\u025b\u025c\u0003N\'\u0000"+
		"\u025c\u025d\u0005<\u0000\u0000\u025d\u025e\u0003:\u001d\u0000\u025eM"+
		"\u0001\u0000\u0000\u0000\u025f\u0260\u0005\u0017\u0000\u0000\u0260\u0263"+
		"\u0003&\u0013\u0000\u0261\u0263\u0005\u0018\u0000\u0000\u0262\u025f\u0001"+
		"\u0000\u0000\u0000\u0262\u0261\u0001\u0000\u0000\u0000\u0263O\u0001\u0000"+
		"\u0000\u0000\u0264\u0265\u0001\u0000\u0000\u0000\u0265Q\u0001\u0000\u0000"+
		"\u0000%Z\\mv\u0083\u0096\u009f\u00ae\u00b3\u00ba\u00c5\u00d4\u00df\u00e7"+
		"\u00f4\u012f\u0131\u0139\u0141\u0149\u014b\u0154\u015b\u0166\u017f\u018a"+
		"\u0193\u019b\u01b4\u01bb\u01c2\u01f4\u021c\u0233\u024f\u0258\u0262";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}