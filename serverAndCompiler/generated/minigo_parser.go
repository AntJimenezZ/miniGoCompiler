// Code generated from C:/Users/noni4/Desktop/miniGoProyect/serverAndCompiler/miniGoParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package generated // miniGoParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type miniGoParser struct {
	*antlr.BaseParser
}

var MiniGoParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func minigoparserParserInit() {
	staticData := &MiniGoParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "'package'", "'var'", "'type'", "'func'", "'struct'",
		"'len'", "'cap'", "'print'", "'println'", "'return'", "'break'", "'continue'",
		"'append'", "'if'", "'else'", "'for'", "'switch'", "'case'", "'default'",
		"", "'+'", "'-'", "'*'", "'/'", "'%'", "'<'", "'<='", "'>'", "'>='",
		"'=='", "'!='", "'&&'", "'||'", "'!'", "'&'", "'|'", "'^'", "'&^'",
		"'<<'", "'>>'", "'++'", "'--'", "'='", "'+='", "'-='", "'*='", "'/='",
		"'%='", "'&='", "'|='", "'^='", "'<<='", "'>>='", "'&^='", "':'", "';'",
		"','", "'.'", "'('", "')'", "'['", "']'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "INTLITERAL", "FLOATLITERAL", "RUNELITERAL", "RAWSTRINGLITERAL",
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
		"RBRACE", "SPACES", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"root", "topDeclarationList", "variableDecl", "innerVarDecls", "singleVarDecl",
		"singleVarDeclNoExps", "typeDecl", "innerTypeDecls", "singleTypeDecl",
		"funcDecl", "funcFrontDecl", "funcArgDecls", "declType", "sliceDeclType",
		"arrayDeclType", "structDeclType", "structMemDecls", "identifierList",
		"expression", "expressionList", "primaryExpression", "operand", "literal",
		"index", "arguments", "selector", "appendExpression", "lengthExpression",
		"capExpression", "statementList", "block", "statement", "simpleStatement",
		"assignmentStatement", "ifStatement", "loop", "switchStmt", "expressionCaseClauseList",
		"expressionCaseClause", "expressionSwitchCase", "epsilon",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 72, 615, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 91, 8, 1, 10, 1, 12, 1, 94, 9, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 110, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 117, 8,
		3, 10, 3, 12, 3, 120, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 132, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 151,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 158, 8, 7, 10, 7, 12, 7, 161,
		9, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 3, 10, 175, 8, 10, 1, 10, 1, 10, 1, 10, 3, 10, 180, 8, 10, 1,
		11, 1, 11, 1, 11, 5, 11, 185, 8, 11, 10, 11, 12, 11, 188, 9, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 198, 8, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 3, 15, 213, 8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 5, 16, 222, 8, 16, 10, 16, 12, 16, 225, 9, 16, 1, 17, 1, 17,
		1, 17, 5, 17, 230, 8, 17, 10, 17, 12, 17, 233, 9, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 245, 8, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 304, 8, 18, 10, 18, 12, 18, 307,
		9, 18, 1, 19, 1, 19, 1, 19, 5, 19, 312, 8, 19, 10, 19, 12, 19, 315, 9,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 322, 8, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 330, 8, 20, 10, 20, 12, 20, 333, 9,
		20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 341, 8, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 348, 8, 22, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 359, 8, 24, 1, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 5, 29, 382,
		8, 29, 10, 29, 12, 29, 385, 9, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 31, 1, 31, 3, 31, 395, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 3, 31, 404, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 3, 31, 412, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 437, 8, 31, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 3, 32, 444, 8, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 3, 32, 451, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33,
		501, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 541,
		8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 3, 35, 564, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 592,
		8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 599, 8, 37, 10, 37, 12,
		37, 602, 9, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 3, 39,
		611, 8, 39, 1, 40, 1, 40, 1, 40, 0, 3, 36, 40, 74, 41, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 0,
		0, 675, 0, 82, 1, 0, 0, 0, 2, 92, 1, 0, 0, 0, 4, 109, 1, 0, 0, 0, 6, 111,
		1, 0, 0, 0, 8, 131, 1, 0, 0, 0, 10, 133, 1, 0, 0, 0, 12, 150, 1, 0, 0,
		0, 14, 152, 1, 0, 0, 0, 16, 162, 1, 0, 0, 0, 18, 165, 1, 0, 0, 0, 20, 169,
		1, 0, 0, 0, 22, 181, 1, 0, 0, 0, 24, 197, 1, 0, 0, 0, 26, 199, 1, 0, 0,
		0, 28, 203, 1, 0, 0, 0, 30, 208, 1, 0, 0, 0, 32, 216, 1, 0, 0, 0, 34, 226,
		1, 0, 0, 0, 36, 244, 1, 0, 0, 0, 38, 308, 1, 0, 0, 0, 40, 321, 1, 0, 0,
		0, 42, 340, 1, 0, 0, 0, 44, 347, 1, 0, 0, 0, 46, 349, 1, 0, 0, 0, 48, 358,
		1, 0, 0, 0, 50, 360, 1, 0, 0, 0, 52, 363, 1, 0, 0, 0, 54, 370, 1, 0, 0,
		0, 56, 375, 1, 0, 0, 0, 58, 383, 1, 0, 0, 0, 60, 386, 1, 0, 0, 0, 62, 436,
		1, 0, 0, 0, 64, 450, 1, 0, 0, 0, 66, 500, 1, 0, 0, 0, 68, 540, 1, 0, 0,
		0, 70, 563, 1, 0, 0, 0, 72, 591, 1, 0, 0, 0, 74, 593, 1, 0, 0, 0, 76, 603,
		1, 0, 0, 0, 78, 610, 1, 0, 0, 0, 80, 612, 1, 0, 0, 0, 82, 83, 5, 6, 0,
		0, 83, 84, 5, 25, 0, 0, 84, 85, 5, 61, 0, 0, 85, 86, 3, 2, 1, 0, 86, 1,
		1, 0, 0, 0, 87, 91, 3, 4, 2, 0, 88, 91, 3, 12, 6, 0, 89, 91, 3, 18, 9,
		0, 90, 87, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 94,
		1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 3, 1, 0, 0, 0,
		94, 92, 1, 0, 0, 0, 95, 96, 5, 7, 0, 0, 96, 97, 3, 8, 4, 0, 97, 98, 5,
		61, 0, 0, 98, 110, 1, 0, 0, 0, 99, 100, 5, 7, 0, 0, 100, 101, 5, 64, 0,
		0, 101, 102, 3, 6, 3, 0, 102, 103, 5, 65, 0, 0, 103, 104, 5, 61, 0, 0,
		104, 110, 1, 0, 0, 0, 105, 106, 5, 7, 0, 0, 106, 107, 5, 64, 0, 0, 107,
		108, 5, 65, 0, 0, 108, 110, 5, 61, 0, 0, 109, 95, 1, 0, 0, 0, 109, 99,
		1, 0, 0, 0, 109, 105, 1, 0, 0, 0, 110, 5, 1, 0, 0, 0, 111, 112, 3, 8, 4,
		0, 112, 118, 5, 61, 0, 0, 113, 114, 3, 8, 4, 0, 114, 115, 5, 61, 0, 0,
		115, 117, 1, 0, 0, 0, 116, 113, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118,
		116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 7, 1, 0, 0, 0, 120, 118, 1,
		0, 0, 0, 121, 122, 3, 34, 17, 0, 122, 123, 3, 24, 12, 0, 123, 124, 5, 48,
		0, 0, 124, 125, 3, 38, 19, 0, 125, 132, 1, 0, 0, 0, 126, 127, 3, 34, 17,
		0, 127, 128, 5, 48, 0, 0, 128, 129, 3, 38, 19, 0, 129, 132, 1, 0, 0, 0,
		130, 132, 3, 10, 5, 0, 131, 121, 1, 0, 0, 0, 131, 126, 1, 0, 0, 0, 131,
		130, 1, 0, 0, 0, 132, 9, 1, 0, 0, 0, 133, 134, 3, 34, 17, 0, 134, 135,
		3, 24, 12, 0, 135, 11, 1, 0, 0, 0, 136, 137, 5, 8, 0, 0, 137, 138, 3, 16,
		8, 0, 138, 139, 5, 61, 0, 0, 139, 151, 1, 0, 0, 0, 140, 141, 5, 8, 0, 0,
		141, 142, 5, 64, 0, 0, 142, 143, 3, 14, 7, 0, 143, 144, 5, 65, 0, 0, 144,
		145, 5, 61, 0, 0, 145, 151, 1, 0, 0, 0, 146, 147, 5, 8, 0, 0, 147, 148,
		5, 64, 0, 0, 148, 149, 5, 65, 0, 0, 149, 151, 5, 61, 0, 0, 150, 136, 1,
		0, 0, 0, 150, 140, 1, 0, 0, 0, 150, 146, 1, 0, 0, 0, 151, 13, 1, 0, 0,
		0, 152, 153, 3, 16, 8, 0, 153, 159, 5, 61, 0, 0, 154, 155, 3, 16, 8, 0,
		155, 156, 5, 61, 0, 0, 156, 158, 1, 0, 0, 0, 157, 154, 1, 0, 0, 0, 158,
		161, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 15, 1,
		0, 0, 0, 161, 159, 1, 0, 0, 0, 162, 163, 5, 25, 0, 0, 163, 164, 3, 24,
		12, 0, 164, 17, 1, 0, 0, 0, 165, 166, 3, 20, 10, 0, 166, 167, 3, 60, 30,
		0, 167, 168, 5, 61, 0, 0, 168, 19, 1, 0, 0, 0, 169, 170, 5, 9, 0, 0, 170,
		171, 5, 25, 0, 0, 171, 174, 5, 64, 0, 0, 172, 175, 3, 22, 11, 0, 173, 175,
		3, 80, 40, 0, 174, 172, 1, 0, 0, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1,
		0, 0, 0, 176, 179, 5, 65, 0, 0, 177, 180, 3, 24, 12, 0, 178, 180, 3, 80,
		40, 0, 179, 177, 1, 0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 21, 1, 0, 0, 0,
		181, 186, 3, 10, 5, 0, 182, 183, 5, 62, 0, 0, 183, 185, 3, 10, 5, 0, 184,
		182, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187,
		1, 0, 0, 0, 187, 23, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 190, 5, 64,
		0, 0, 190, 191, 3, 24, 12, 0, 191, 192, 5, 65, 0, 0, 192, 198, 1, 0, 0,
		0, 193, 198, 5, 25, 0, 0, 194, 198, 3, 26, 13, 0, 195, 198, 3, 28, 14,
		0, 196, 198, 3, 30, 15, 0, 197, 189, 1, 0, 0, 0, 197, 193, 1, 0, 0, 0,
		197, 194, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197, 196, 1, 0, 0, 0, 198,
		25, 1, 0, 0, 0, 199, 200, 5, 66, 0, 0, 200, 201, 5, 67, 0, 0, 201, 202,
		3, 24, 12, 0, 202, 27, 1, 0, 0, 0, 203, 204, 5, 66, 0, 0, 204, 205, 5,
		1, 0, 0, 205, 206, 5, 67, 0, 0, 206, 207, 3, 24, 12, 0, 207, 29, 1, 0,
		0, 0, 208, 209, 5, 10, 0, 0, 209, 212, 5, 68, 0, 0, 210, 213, 3, 32, 16,
		0, 211, 213, 3, 80, 40, 0, 212, 210, 1, 0, 0, 0, 212, 211, 1, 0, 0, 0,
		213, 214, 1, 0, 0, 0, 214, 215, 5, 69, 0, 0, 215, 31, 1, 0, 0, 0, 216,
		217, 3, 10, 5, 0, 217, 223, 5, 61, 0, 0, 218, 219, 3, 10, 5, 0, 219, 220,
		5, 61, 0, 0, 220, 222, 1, 0, 0, 0, 221, 218, 1, 0, 0, 0, 222, 225, 1, 0,
		0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 33, 1, 0, 0, 0,
		225, 223, 1, 0, 0, 0, 226, 231, 5, 25, 0, 0, 227, 228, 5, 62, 0, 0, 228,
		230, 5, 25, 0, 0, 229, 227, 1, 0, 0, 0, 230, 233, 1, 0, 0, 0, 231, 229,
		1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 35, 1, 0, 0, 0, 233, 231, 1, 0,
		0, 0, 234, 235, 6, 18, -1, 0, 235, 245, 3, 40, 20, 0, 236, 237, 5, 26,
		0, 0, 237, 245, 3, 36, 18, 4, 238, 239, 5, 27, 0, 0, 239, 245, 3, 36, 18,
		3, 240, 241, 5, 39, 0, 0, 241, 245, 3, 36, 18, 2, 242, 243, 5, 42, 0, 0,
		243, 245, 3, 36, 18, 1, 244, 234, 1, 0, 0, 0, 244, 236, 1, 0, 0, 0, 244,
		238, 1, 0, 0, 0, 244, 240, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 245, 305,
		1, 0, 0, 0, 246, 247, 10, 23, 0, 0, 247, 248, 5, 28, 0, 0, 248, 304, 3,
		36, 18, 24, 249, 250, 10, 22, 0, 0, 250, 251, 5, 29, 0, 0, 251, 304, 3,
		36, 18, 23, 252, 253, 10, 21, 0, 0, 253, 254, 5, 30, 0, 0, 254, 304, 3,
		36, 18, 22, 255, 256, 10, 20, 0, 0, 256, 257, 5, 44, 0, 0, 257, 304, 3,
		36, 18, 21, 258, 259, 10, 19, 0, 0, 259, 260, 5, 45, 0, 0, 260, 304, 3,
		36, 18, 20, 261, 262, 10, 18, 0, 0, 262, 263, 5, 40, 0, 0, 263, 304, 3,
		36, 18, 19, 264, 265, 10, 17, 0, 0, 265, 266, 5, 43, 0, 0, 266, 304, 3,
		36, 18, 18, 267, 268, 10, 16, 0, 0, 268, 269, 5, 26, 0, 0, 269, 304, 3,
		36, 18, 17, 270, 271, 10, 15, 0, 0, 271, 272, 5, 27, 0, 0, 272, 304, 3,
		36, 18, 16, 273, 274, 10, 14, 0, 0, 274, 275, 5, 41, 0, 0, 275, 304, 3,
		36, 18, 15, 276, 277, 10, 13, 0, 0, 277, 278, 5, 42, 0, 0, 278, 304, 3,
		36, 18, 14, 279, 280, 10, 12, 0, 0, 280, 281, 5, 35, 0, 0, 281, 304, 3,
		36, 18, 13, 282, 283, 10, 11, 0, 0, 283, 284, 5, 36, 0, 0, 284, 304, 3,
		36, 18, 12, 285, 286, 10, 10, 0, 0, 286, 287, 5, 31, 0, 0, 287, 304, 3,
		36, 18, 11, 288, 289, 10, 9, 0, 0, 289, 290, 5, 32, 0, 0, 290, 304, 3,
		36, 18, 10, 291, 292, 10, 8, 0, 0, 292, 293, 5, 33, 0, 0, 293, 304, 3,
		36, 18, 9, 294, 295, 10, 7, 0, 0, 295, 296, 5, 34, 0, 0, 296, 304, 3, 36,
		18, 8, 297, 298, 10, 6, 0, 0, 298, 299, 5, 37, 0, 0, 299, 304, 3, 36, 18,
		7, 300, 301, 10, 5, 0, 0, 301, 302, 5, 38, 0, 0, 302, 304, 3, 36, 18, 6,
		303, 246, 1, 0, 0, 0, 303, 249, 1, 0, 0, 0, 303, 252, 1, 0, 0, 0, 303,
		255, 1, 0, 0, 0, 303, 258, 1, 0, 0, 0, 303, 261, 1, 0, 0, 0, 303, 264,
		1, 0, 0, 0, 303, 267, 1, 0, 0, 0, 303, 270, 1, 0, 0, 0, 303, 273, 1, 0,
		0, 0, 303, 276, 1, 0, 0, 0, 303, 279, 1, 0, 0, 0, 303, 282, 1, 0, 0, 0,
		303, 285, 1, 0, 0, 0, 303, 288, 1, 0, 0, 0, 303, 291, 1, 0, 0, 0, 303,
		294, 1, 0, 0, 0, 303, 297, 1, 0, 0, 0, 303, 300, 1, 0, 0, 0, 304, 307,
		1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 37, 1, 0,
		0, 0, 307, 305, 1, 0, 0, 0, 308, 313, 3, 36, 18, 0, 309, 310, 5, 62, 0,
		0, 310, 312, 3, 36, 18, 0, 311, 309, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0,
		313, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 39, 1, 0, 0, 0, 315, 313,
		1, 0, 0, 0, 316, 317, 6, 20, -1, 0, 317, 322, 3, 42, 21, 0, 318, 322, 3,
		52, 26, 0, 319, 322, 3, 54, 27, 0, 320, 322, 3, 56, 28, 0, 321, 316, 1,
		0, 0, 0, 321, 318, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 321, 320, 1, 0, 0,
		0, 322, 331, 1, 0, 0, 0, 323, 324, 10, 6, 0, 0, 324, 330, 3, 50, 25, 0,
		325, 326, 10, 5, 0, 0, 326, 330, 3, 46, 23, 0, 327, 328, 10, 4, 0, 0, 328,
		330, 3, 48, 24, 0, 329, 323, 1, 0, 0, 0, 329, 325, 1, 0, 0, 0, 329, 327,
		1, 0, 0, 0, 330, 333, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 331, 332, 1, 0,
		0, 0, 332, 41, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 334, 341, 3, 44, 22, 0,
		335, 341, 5, 25, 0, 0, 336, 337, 5, 64, 0, 0, 337, 338, 3, 36, 18, 0, 338,
		339, 5, 65, 0, 0, 339, 341, 1, 0, 0, 0, 340, 334, 1, 0, 0, 0, 340, 335,
		1, 0, 0, 0, 340, 336, 1, 0, 0, 0, 341, 43, 1, 0, 0, 0, 342, 348, 5, 1,
		0, 0, 343, 348, 5, 2, 0, 0, 344, 348, 5, 3, 0, 0, 345, 348, 5, 4, 0, 0,
		346, 348, 5, 5, 0, 0, 347, 342, 1, 0, 0, 0, 347, 343, 1, 0, 0, 0, 347,
		344, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 347, 346, 1, 0, 0, 0, 348, 45, 1,
		0, 0, 0, 349, 350, 5, 66, 0, 0, 350, 351, 3, 36, 18, 0, 351, 352, 5, 67,
		0, 0, 352, 47, 1, 0, 0, 0, 353, 354, 5, 64, 0, 0, 354, 359, 3, 38, 19,
		0, 355, 356, 3, 80, 40, 0, 356, 357, 5, 65, 0, 0, 357, 359, 1, 0, 0, 0,
		358, 353, 1, 0, 0, 0, 358, 355, 1, 0, 0, 0, 359, 49, 1, 0, 0, 0, 360, 361,
		5, 63, 0, 0, 361, 362, 5, 25, 0, 0, 362, 51, 1, 0, 0, 0, 363, 364, 5, 18,
		0, 0, 364, 365, 5, 64, 0, 0, 365, 366, 3, 36, 18, 0, 366, 367, 5, 62, 0,
		0, 367, 368, 3, 36, 18, 0, 368, 369, 5, 65, 0, 0, 369, 53, 1, 0, 0, 0,
		370, 371, 5, 11, 0, 0, 371, 372, 5, 64, 0, 0, 372, 373, 3, 36, 18, 0, 373,
		374, 5, 65, 0, 0, 374, 55, 1, 0, 0, 0, 375, 376, 5, 12, 0, 0, 376, 377,
		5, 64, 0, 0, 377, 378, 3, 36, 18, 0, 378, 379, 5, 65, 0, 0, 379, 57, 1,
		0, 0, 0, 380, 382, 3, 62, 31, 0, 381, 380, 1, 0, 0, 0, 382, 385, 1, 0,
		0, 0, 383, 381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 59, 1, 0, 0, 0,
		385, 383, 1, 0, 0, 0, 386, 387, 5, 68, 0, 0, 387, 388, 3, 58, 29, 0, 388,
		389, 5, 69, 0, 0, 389, 61, 1, 0, 0, 0, 390, 391, 5, 13, 0, 0, 391, 394,
		5, 64, 0, 0, 392, 395, 3, 38, 19, 0, 393, 395, 3, 80, 40, 0, 394, 392,
		1, 0, 0, 0, 394, 393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 397, 5, 65,
		0, 0, 397, 398, 5, 61, 0, 0, 398, 437, 1, 0, 0, 0, 399, 400, 5, 14, 0,
		0, 400, 403, 5, 64, 0, 0, 401, 404, 3, 38, 19, 0, 402, 404, 3, 80, 40,
		0, 403, 401, 1, 0, 0, 0, 403, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405,
		406, 5, 65, 0, 0, 406, 407, 5, 61, 0, 0, 407, 437, 1, 0, 0, 0, 408, 411,
		5, 15, 0, 0, 409, 412, 3, 36, 18, 0, 410, 412, 3, 80, 40, 0, 411, 409,
		1, 0, 0, 0, 411, 410, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 414, 5, 61,
		0, 0, 414, 437, 1, 0, 0, 0, 415, 416, 5, 16, 0, 0, 416, 437, 5, 61, 0,
		0, 417, 418, 5, 17, 0, 0, 418, 437, 5, 61, 0, 0, 419, 420, 3, 64, 32, 0,
		420, 421, 5, 61, 0, 0, 421, 437, 1, 0, 0, 0, 422, 423, 3, 60, 30, 0, 423,
		424, 5, 61, 0, 0, 424, 437, 1, 0, 0, 0, 425, 426, 3, 72, 36, 0, 426, 427,
		5, 61, 0, 0, 427, 437, 1, 0, 0, 0, 428, 429, 3, 68, 34, 0, 429, 430, 5,
		61, 0, 0, 430, 437, 1, 0, 0, 0, 431, 432, 3, 70, 35, 0, 432, 433, 5, 61,
		0, 0, 433, 437, 1, 0, 0, 0, 434, 437, 3, 12, 6, 0, 435, 437, 3, 4, 2, 0,
		436, 390, 1, 0, 0, 0, 436, 399, 1, 0, 0, 0, 436, 408, 1, 0, 0, 0, 436,
		415, 1, 0, 0, 0, 436, 417, 1, 0, 0, 0, 436, 419, 1, 0, 0, 0, 436, 422,
		1, 0, 0, 0, 436, 425, 1, 0, 0, 0, 436, 428, 1, 0, 0, 0, 436, 431, 1, 0,
		0, 0, 436, 434, 1, 0, 0, 0, 436, 435, 1, 0, 0, 0, 437, 63, 1, 0, 0, 0,
		438, 451, 3, 80, 40, 0, 439, 443, 3, 36, 18, 0, 440, 444, 5, 46, 0, 0,
		441, 444, 5, 47, 0, 0, 442, 444, 3, 80, 40, 0, 443, 440, 1, 0, 0, 0, 443,
		441, 1, 0, 0, 0, 443, 442, 1, 0, 0, 0, 444, 451, 1, 0, 0, 0, 445, 451,
		3, 66, 33, 0, 446, 447, 3, 38, 19, 0, 447, 448, 5, 48, 0, 0, 448, 449,
		3, 38, 19, 0, 449, 451, 1, 0, 0, 0, 450, 438, 1, 0, 0, 0, 450, 439, 1,
		0, 0, 0, 450, 445, 1, 0, 0, 0, 450, 446, 1, 0, 0, 0, 451, 65, 1, 0, 0,
		0, 452, 453, 3, 38, 19, 0, 453, 454, 5, 48, 0, 0, 454, 455, 3, 38, 19,
		0, 455, 501, 1, 0, 0, 0, 456, 457, 3, 36, 18, 0, 457, 458, 5, 49, 0, 0,
		458, 459, 3, 36, 18, 0, 459, 501, 1, 0, 0, 0, 460, 461, 3, 36, 18, 0, 461,
		462, 5, 50, 0, 0, 462, 463, 3, 36, 18, 0, 463, 501, 1, 0, 0, 0, 464, 465,
		3, 36, 18, 0, 465, 466, 5, 54, 0, 0, 466, 467, 3, 36, 18, 0, 467, 501,
		1, 0, 0, 0, 468, 469, 3, 36, 18, 0, 469, 470, 5, 55, 0, 0, 470, 471, 3,
		36, 18, 0, 471, 501, 1, 0, 0, 0, 472, 473, 3, 36, 18, 0, 473, 474, 5, 51,
		0, 0, 474, 475, 3, 36, 18, 0, 475, 501, 1, 0, 0, 0, 476, 477, 3, 36, 18,
		0, 477, 478, 5, 56, 0, 0, 478, 479, 3, 36, 18, 0, 479, 501, 1, 0, 0, 0,
		480, 481, 3, 36, 18, 0, 481, 482, 5, 57, 0, 0, 482, 483, 3, 36, 18, 0,
		483, 501, 1, 0, 0, 0, 484, 485, 3, 36, 18, 0, 485, 486, 5, 58, 0, 0, 486,
		487, 3, 36, 18, 0, 487, 501, 1, 0, 0, 0, 488, 489, 3, 36, 18, 0, 489, 490,
		5, 59, 0, 0, 490, 491, 3, 36, 18, 0, 491, 501, 1, 0, 0, 0, 492, 493, 3,
		36, 18, 0, 493, 494, 5, 53, 0, 0, 494, 495, 3, 36, 18, 0, 495, 501, 1,
		0, 0, 0, 496, 497, 3, 36, 18, 0, 497, 498, 5, 52, 0, 0, 498, 499, 3, 36,
		18, 0, 499, 501, 1, 0, 0, 0, 500, 452, 1, 0, 0, 0, 500, 456, 1, 0, 0, 0,
		500, 460, 1, 0, 0, 0, 500, 464, 1, 0, 0, 0, 500, 468, 1, 0, 0, 0, 500,
		472, 1, 0, 0, 0, 500, 476, 1, 0, 0, 0, 500, 480, 1, 0, 0, 0, 500, 484,
		1, 0, 0, 0, 500, 488, 1, 0, 0, 0, 500, 492, 1, 0, 0, 0, 500, 496, 1, 0,
		0, 0, 501, 67, 1, 0, 0, 0, 502, 503, 5, 19, 0, 0, 503, 504, 3, 36, 18,
		0, 504, 505, 3, 60, 30, 0, 505, 541, 1, 0, 0, 0, 506, 507, 5, 19, 0, 0,
		507, 508, 3, 36, 18, 0, 508, 509, 3, 60, 30, 0, 509, 510, 5, 20, 0, 0,
		510, 511, 3, 68, 34, 0, 511, 541, 1, 0, 0, 0, 512, 513, 5, 19, 0, 0, 513,
		514, 3, 36, 18, 0, 514, 515, 3, 60, 30, 0, 515, 516, 5, 20, 0, 0, 516,
		517, 3, 60, 30, 0, 517, 541, 1, 0, 0, 0, 518, 519, 5, 19, 0, 0, 519, 520,
		3, 64, 32, 0, 520, 521, 5, 61, 0, 0, 521, 522, 3, 36, 18, 0, 522, 523,
		3, 60, 30, 0, 523, 541, 1, 0, 0, 0, 524, 525, 5, 19, 0, 0, 525, 526, 3,
		64, 32, 0, 526, 527, 5, 61, 0, 0, 527, 528, 3, 36, 18, 0, 528, 529, 3,
		60, 30, 0, 529, 530, 5, 20, 0, 0, 530, 531, 3, 68, 34, 0, 531, 541, 1,
		0, 0, 0, 532, 533, 5, 19, 0, 0, 533, 534, 3, 64, 32, 0, 534, 535, 5, 61,
		0, 0, 535, 536, 3, 36, 18, 0, 536, 537, 3, 60, 30, 0, 537, 538, 5, 20,
		0, 0, 538, 539, 3, 60, 30, 0, 539, 541, 1, 0, 0, 0, 540, 502, 1, 0, 0,
		0, 540, 506, 1, 0, 0, 0, 540, 512, 1, 0, 0, 0, 540, 518, 1, 0, 0, 0, 540,
		524, 1, 0, 0, 0, 540, 532, 1, 0, 0, 0, 541, 69, 1, 0, 0, 0, 542, 543, 5,
		21, 0, 0, 543, 564, 3, 60, 30, 0, 544, 545, 5, 21, 0, 0, 545, 546, 3, 36,
		18, 0, 546, 547, 3, 60, 30, 0, 547, 564, 1, 0, 0, 0, 548, 549, 5, 21, 0,
		0, 549, 550, 3, 64, 32, 0, 550, 551, 5, 61, 0, 0, 551, 552, 3, 36, 18,
		0, 552, 553, 5, 61, 0, 0, 553, 554, 3, 64, 32, 0, 554, 555, 3, 60, 30,
		0, 555, 564, 1, 0, 0, 0, 556, 557, 5, 21, 0, 0, 557, 558, 3, 64, 32, 0,
		558, 559, 5, 61, 0, 0, 559, 560, 5, 61, 0, 0, 560, 561, 3, 64, 32, 0, 561,
		562, 3, 60, 30, 0, 562, 564, 1, 0, 0, 0, 563, 542, 1, 0, 0, 0, 563, 544,
		1, 0, 0, 0, 563, 548, 1, 0, 0, 0, 563, 556, 1, 0, 0, 0, 564, 71, 1, 0,
		0, 0, 565, 566, 5, 22, 0, 0, 566, 567, 3, 64, 32, 0, 567, 568, 5, 61, 0,
		0, 568, 569, 3, 36, 18, 0, 569, 570, 5, 68, 0, 0, 570, 571, 3, 74, 37,
		0, 571, 572, 5, 69, 0, 0, 572, 592, 1, 0, 0, 0, 573, 574, 5, 22, 0, 0,
		574, 575, 3, 36, 18, 0, 575, 576, 5, 68, 0, 0, 576, 577, 3, 74, 37, 0,
		577, 578, 5, 69, 0, 0, 578, 592, 1, 0, 0, 0, 579, 580, 5, 22, 0, 0, 580,
		581, 3, 64, 32, 0, 581, 582, 5, 61, 0, 0, 582, 583, 5, 68, 0, 0, 583, 584,
		3, 74, 37, 0, 584, 585, 5, 69, 0, 0, 585, 592, 1, 0, 0, 0, 586, 587, 5,
		22, 0, 0, 587, 588, 5, 68, 0, 0, 588, 589, 3, 74, 37, 0, 589, 590, 5, 69,
		0, 0, 590, 592, 1, 0, 0, 0, 591, 565, 1, 0, 0, 0, 591, 573, 1, 0, 0, 0,
		591, 579, 1, 0, 0, 0, 591, 586, 1, 0, 0, 0, 592, 73, 1, 0, 0, 0, 593, 594,
		6, 37, -1, 0, 594, 595, 3, 80, 40, 0, 595, 600, 1, 0, 0, 0, 596, 597, 10,
		1, 0, 0, 597, 599, 3, 76, 38, 0, 598, 596, 1, 0, 0, 0, 599, 602, 1, 0,
		0, 0, 600, 598, 1, 0, 0, 0, 600, 601, 1, 0, 0, 0, 601, 75, 1, 0, 0, 0,
		602, 600, 1, 0, 0, 0, 603, 604, 3, 78, 39, 0, 604, 605, 5, 60, 0, 0, 605,
		606, 3, 58, 29, 0, 606, 77, 1, 0, 0, 0, 607, 608, 5, 23, 0, 0, 608, 611,
		3, 38, 19, 0, 609, 611, 5, 24, 0, 0, 610, 607, 1, 0, 0, 0, 610, 609, 1,
		0, 0, 0, 611, 79, 1, 0, 0, 0, 612, 613, 1, 0, 0, 0, 613, 81, 1, 0, 0, 0,
		37, 90, 92, 109, 118, 131, 150, 159, 174, 179, 186, 197, 212, 223, 231,
		244, 303, 305, 313, 321, 329, 331, 340, 347, 358, 383, 394, 403, 411, 436,
		443, 450, 500, 540, 563, 591, 600, 610,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// miniGoParserInit initializes any static state used to implement miniGoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewminiGoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MiniGoParserInit() {
	staticData := &MiniGoParserParserStaticData
	staticData.once.Do(minigoparserParserInit)
}

// NewminiGoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewminiGoParser(input antlr.TokenStream) *miniGoParser {
	MiniGoParserInit()
	this := new(miniGoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MiniGoParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "miniGoParser.g4"

	return this
}

// miniGoParser tokens.
const (
	miniGoParserEOF                      = antlr.TokenEOF
	miniGoParserINTLITERAL               = 1
	miniGoParserFLOATLITERAL             = 2
	miniGoParserRUNELITERAL              = 3
	miniGoParserRAWSTRINGLITERAL         = 4
	miniGoParserINTERPRETEDSTRINGLITERAL = 5
	miniGoParserPACKAGE                  = 6
	miniGoParserVAR                      = 7
	miniGoParserTYPE                     = 8
	miniGoParserFUNC                     = 9
	miniGoParserSTRUCT                   = 10
	miniGoParserLEN                      = 11
	miniGoParserCAP                      = 12
	miniGoParserPRINT                    = 13
	miniGoParserPRINTLN                  = 14
	miniGoParserRETURN                   = 15
	miniGoParserBREAK                    = 16
	miniGoParserCONTINUE                 = 17
	miniGoParserAPPEND                   = 18
	miniGoParserIF                       = 19
	miniGoParserELSE                     = 20
	miniGoParserFOR                      = 21
	miniGoParserSWITCH                   = 22
	miniGoParserCASE                     = 23
	miniGoParserDEFAULT                  = 24
	miniGoParserIDENTIFIER               = 25
	miniGoParserPLUS                     = 26
	miniGoParserMINUS                    = 27
	miniGoParserMULTIPLY                 = 28
	miniGoParserDIVIDE                   = 29
	miniGoParserMODULO                   = 30
	miniGoParserLESS                     = 31
	miniGoParserLESSEQUAL                = 32
	miniGoParserGREATER                  = 33
	miniGoParserGREATEREQUAL             = 34
	miniGoParserEQUAL                    = 35
	miniGoParserNOTEQUAL                 = 36
	miniGoParserAND                      = 37
	miniGoParserOR                       = 38
	miniGoParserNOT                      = 39
	miniGoParserBITWISEAND               = 40
	miniGoParserBITWISEOR                = 41
	miniGoParserBITWISEXOR               = 42
	miniGoParserBITWISECLEAR             = 43
	miniGoParserSHIFTLEFT                = 44
	miniGoParserSHIFTRIGHT               = 45
	miniGoParserINCREMENT                = 46
	miniGoParserDECREMENT                = 47
	miniGoParserASSIGN                   = 48
	miniGoParserPLUSEQUAL                = 49
	miniGoParserMINUSEQUAL               = 50
	miniGoParserMULTIPLYEQUAL            = 51
	miniGoParserDIVIDEEQUAL              = 52
	miniGoParserMODULOEQUAL              = 53
	miniGoParserBITWISEANDEQUAL          = 54
	miniGoParserBITWISEOREQUAL           = 55
	miniGoParserBITWISEXOREQUAL          = 56
	miniGoParserSHIFTLEFTEQUAL           = 57
	miniGoParserSHIFTRIGHTEQUAL          = 58
	miniGoParserBITWISECLEAREQUAL        = 59
	miniGoParserCOLON                    = 60
	miniGoParserSEMICOLON                = 61
	miniGoParserCOMMA                    = 62
	miniGoParserDOT                      = 63
	miniGoParserLPAREN                   = 64
	miniGoParserRPAREN                   = 65
	miniGoParserLBRACKET                 = 66
	miniGoParserRBRACKET                 = 67
	miniGoParserLBRACE                   = 68
	miniGoParserRBRACE                   = 69
	miniGoParserSPACES                   = 70
	miniGoParserLINE_COMMENT             = 71
	miniGoParserBLOCK_COMMENT            = 72
)

// miniGoParser rules.
const (
	miniGoParserRULE_root                     = 0
	miniGoParserRULE_topDeclarationList       = 1
	miniGoParserRULE_variableDecl             = 2
	miniGoParserRULE_innerVarDecls            = 3
	miniGoParserRULE_singleVarDecl            = 4
	miniGoParserRULE_singleVarDeclNoExps      = 5
	miniGoParserRULE_typeDecl                 = 6
	miniGoParserRULE_innerTypeDecls           = 7
	miniGoParserRULE_singleTypeDecl           = 8
	miniGoParserRULE_funcDecl                 = 9
	miniGoParserRULE_funcFrontDecl            = 10
	miniGoParserRULE_funcArgDecls             = 11
	miniGoParserRULE_declType                 = 12
	miniGoParserRULE_sliceDeclType            = 13
	miniGoParserRULE_arrayDeclType            = 14
	miniGoParserRULE_structDeclType           = 15
	miniGoParserRULE_structMemDecls           = 16
	miniGoParserRULE_identifierList           = 17
	miniGoParserRULE_expression               = 18
	miniGoParserRULE_expressionList           = 19
	miniGoParserRULE_primaryExpression        = 20
	miniGoParserRULE_operand                  = 21
	miniGoParserRULE_literal                  = 22
	miniGoParserRULE_index                    = 23
	miniGoParserRULE_arguments                = 24
	miniGoParserRULE_selector                 = 25
	miniGoParserRULE_appendExpression         = 26
	miniGoParserRULE_lengthExpression         = 27
	miniGoParserRULE_capExpression            = 28
	miniGoParserRULE_statementList            = 29
	miniGoParserRULE_block                    = 30
	miniGoParserRULE_statement                = 31
	miniGoParserRULE_simpleStatement          = 32
	miniGoParserRULE_assignmentStatement      = 33
	miniGoParserRULE_ifStatement              = 34
	miniGoParserRULE_loop                     = 35
	miniGoParserRULE_switchStmt               = 36
	miniGoParserRULE_expressionCaseClauseList = 37
	miniGoParserRULE_expressionCaseClause     = 38
	miniGoParserRULE_expressionSwitchCase     = 39
	miniGoParserRULE_epsilon                  = 40
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) CopyAll(ctx *RootContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RootASTContext struct {
	RootContext
}

func NewRootASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RootASTContext {
	var p = new(RootASTContext)

	InitEmptyRootContext(&p.RootContext)
	p.parser = parser
	p.CopyAll(ctx.(*RootContext))

	return p
}

func (s *RootASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootASTContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(miniGoParserPACKAGE, 0)
}

func (s *RootASTContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(miniGoParserIDENTIFIER, 0)
}

func (s *RootASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *RootASTContext) TopDeclarationList() ITopDeclarationListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopDeclarationListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopDeclarationListContext)
}

func (s *RootASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterRootAST(s)
	}
}

func (s *RootASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitRootAST(s)
	}
}

func (s *RootASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitRootAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, miniGoParserRULE_root)
	localctx = NewRootASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(miniGoParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(miniGoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(miniGoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.TopDeclarationList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITopDeclarationListContext is an interface to support dynamic dispatch.
type ITopDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTopDeclarationListContext differentiates from other interfaces.
	IsTopDeclarationListContext()
}

type TopDeclarationListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopDeclarationListContext() *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_topDeclarationList
	return p
}

func InitEmptyTopDeclarationListContext(p *TopDeclarationListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_topDeclarationList
}

func (*TopDeclarationListContext) IsTopDeclarationListContext() {}

func NewTopDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_topDeclarationList

	return p
}

func (s *TopDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *TopDeclarationListContext) CopyAll(ctx *TopDeclarationListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TopDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TopDeclarationListASTContext struct {
	TopDeclarationListContext
}

func NewTopDeclarationListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopDeclarationListASTContext {
	var p = new(TopDeclarationListASTContext)

	InitEmptyTopDeclarationListContext(&p.TopDeclarationListContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopDeclarationListContext))

	return p
}

func (s *TopDeclarationListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDeclarationListASTContext) AllVariableDecl() []IVariableDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclContext); ok {
			tst[i] = t.(IVariableDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListASTContext) VariableDecl(i int) IVariableDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *TopDeclarationListASTContext) AllTypeDecl() []ITypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclContext); ok {
			tst[i] = t.(ITypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListASTContext) TypeDecl(i int) ITypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *TopDeclarationListASTContext) AllFuncDecl() []IFuncDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDeclContext); ok {
			len++
		}
	}

	tst := make([]IFuncDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDeclContext); ok {
			tst[i] = t.(IFuncDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListASTContext) FuncDecl(i int) IFuncDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclContext)
}

func (s *TopDeclarationListASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterTopDeclarationListAST(s)
	}
}

func (s *TopDeclarationListASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitTopDeclarationListAST(s)
	}
}

func (s *TopDeclarationListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitTopDeclarationListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) TopDeclarationList() (localctx ITopDeclarationListContext) {
	localctx = NewTopDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, miniGoParserRULE_topDeclarationList)
	var _la int

	localctx = NewTopDeclarationListASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0 {
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case miniGoParserVAR:
			{
				p.SetState(87)
				p.VariableDecl()
			}

		case miniGoParserTYPE:
			{
				p.SetState(88)
				p.TypeDecl()
			}

		case miniGoParserFUNC:
			{
				p.SetState(89)
				p.FuncDecl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclContext is an interface to support dynamic dispatch.
type IVariableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVariableDeclContext differentiates from other interfaces.
	IsVariableDeclContext()
}

type VariableDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclContext() *VariableDeclContext {
	var p = new(VariableDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_variableDecl
	return p
}

func InitEmptyVariableDeclContext(p *VariableDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_variableDecl
}

func (*VariableDeclContext) IsVariableDeclContext() {}

func NewVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclContext {
	var p = new(VariableDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_variableDecl

	return p
}

func (s *VariableDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclContext) CopyAll(ctx *VariableDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VariableDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VariableDeclBlockASTContext struct {
	VariableDeclContext
}

func NewVariableDeclBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclBlockASTContext {
	var p = new(VariableDeclBlockASTContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *VariableDeclBlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclBlockASTContext) VAR() antlr.TerminalNode {
	return s.GetToken(miniGoParserVAR, 0)
}

func (s *VariableDeclBlockASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *VariableDeclBlockASTContext) InnerVarDecls() IInnerVarDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerVarDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerVarDeclsContext)
}

func (s *VariableDeclBlockASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *VariableDeclBlockASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *VariableDeclBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterVariableDeclBlockAST(s)
	}
}

func (s *VariableDeclBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitVariableDeclBlockAST(s)
	}
}

func (s *VariableDeclBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitVariableDeclBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableDeclEmptyASTContext struct {
	VariableDeclContext
}

func NewVariableDeclEmptyASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclEmptyASTContext {
	var p = new(VariableDeclEmptyASTContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *VariableDeclEmptyASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclEmptyASTContext) VAR() antlr.TerminalNode {
	return s.GetToken(miniGoParserVAR, 0)
}

func (s *VariableDeclEmptyASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *VariableDeclEmptyASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *VariableDeclEmptyASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *VariableDeclEmptyASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterVariableDeclEmptyAST(s)
	}
}

func (s *VariableDeclEmptyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitVariableDeclEmptyAST(s)
	}
}

func (s *VariableDeclEmptyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitVariableDeclEmptyAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableDeclASTContext struct {
	VariableDeclContext
}

func NewVariableDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclASTContext {
	var p = new(VariableDeclASTContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *VariableDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclASTContext) VAR() antlr.TerminalNode {
	return s.GetToken(miniGoParserVAR, 0)
}

func (s *VariableDeclASTContext) SingleVarDecl() ISingleVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclContext)
}

func (s *VariableDeclASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *VariableDeclASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterVariableDeclAST(s)
	}
}

func (s *VariableDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitVariableDeclAST(s)
	}
}

func (s *VariableDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitVariableDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) VariableDecl() (localctx IVariableDeclContext) {
	localctx = NewVariableDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, miniGoParserRULE_variableDecl)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVariableDeclASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Match(miniGoParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.SingleVarDecl()
		}
		{
			p.SetState(97)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewVariableDeclBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Match(miniGoParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(miniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.InnerVarDecls()
		}
		{
			p.SetState(102)
			p.Match(miniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVariableDeclEmptyASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.Match(miniGoParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(miniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(miniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInnerVarDeclsContext is an interface to support dynamic dispatch.
type IInnerVarDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInnerVarDeclsContext differentiates from other interfaces.
	IsInnerVarDeclsContext()
}

type InnerVarDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerVarDeclsContext() *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_innerVarDecls
	return p
}

func InitEmptyInnerVarDeclsContext(p *InnerVarDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_innerVarDecls
}

func (*InnerVarDeclsContext) IsInnerVarDeclsContext() {}

func NewInnerVarDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_innerVarDecls

	return p
}

func (s *InnerVarDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerVarDeclsContext) CopyAll(ctx *InnerVarDeclsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InnerVarDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerVarDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InnerVarDeclsASTContext struct {
	InnerVarDeclsContext
}

func NewInnerVarDeclsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InnerVarDeclsASTContext {
	var p = new(InnerVarDeclsASTContext)

	InitEmptyInnerVarDeclsContext(&p.InnerVarDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InnerVarDeclsContext))

	return p
}

func (s *InnerVarDeclsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerVarDeclsASTContext) AllSingleVarDecl() []ISingleVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclContext); ok {
			tst[i] = t.(ISingleVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *InnerVarDeclsASTContext) SingleVarDecl(i int) ISingleVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclContext)
}

func (s *InnerVarDeclsASTContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(miniGoParserSEMICOLON)
}

func (s *InnerVarDeclsASTContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, i)
}

func (s *InnerVarDeclsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterInnerVarDeclsAST(s)
	}
}

func (s *InnerVarDeclsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitInnerVarDeclsAST(s)
	}
}

func (s *InnerVarDeclsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitInnerVarDeclsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) InnerVarDecls() (localctx IInnerVarDeclsContext) {
	localctx = NewInnerVarDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, miniGoParserRULE_innerVarDecls)
	var _la int

	localctx = NewInnerVarDeclsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.SingleVarDecl()
	}
	{
		p.SetState(112)
		p.Match(miniGoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == miniGoParserIDENTIFIER {
		{
			p.SetState(113)
			p.SingleVarDecl()
		}
		{
			p.SetState(114)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleVarDeclContext is an interface to support dynamic dispatch.
type ISingleVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingleVarDeclContext differentiates from other interfaces.
	IsSingleVarDeclContext()
}

type SingleVarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleVarDeclContext() *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_singleVarDecl
	return p
}

func InitEmptySingleVarDeclContext(p *SingleVarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_singleVarDecl
}

func (*SingleVarDeclContext) IsSingleVarDeclContext() {}

func NewSingleVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_singleVarDecl

	return p
}

func (s *SingleVarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclContext) CopyAll(ctx *SingleVarDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleVarDeclAssignASTContext struct {
	SingleVarDeclContext
}

func NewSingleVarDeclAssignASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleVarDeclAssignASTContext {
	var p = new(SingleVarDeclAssignASTContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *SingleVarDeclAssignASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclAssignASTContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *SingleVarDeclAssignASTContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(miniGoParserASSIGN, 0)
}

func (s *SingleVarDeclAssignASTContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *SingleVarDeclAssignASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSingleVarDeclAssignAST(s)
	}
}

func (s *SingleVarDeclAssignASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSingleVarDeclAssignAST(s)
	}
}

func (s *SingleVarDeclAssignASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSingleVarDeclAssignAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleVarDeclASTContext struct {
	SingleVarDeclContext
}

func NewSingleVarDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleVarDeclASTContext {
	var p = new(SingleVarDeclASTContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *SingleVarDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclASTContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *SingleVarDeclASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleVarDeclASTContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(miniGoParserASSIGN, 0)
}

func (s *SingleVarDeclASTContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *SingleVarDeclASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSingleVarDeclAST(s)
	}
}

func (s *SingleVarDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSingleVarDeclAST(s)
	}
}

func (s *SingleVarDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSingleVarDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleVarDeclNoExpsASTContext struct {
	SingleVarDeclContext
}

func NewSingleVarDeclNoExpsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleVarDeclNoExpsASTContext {
	var p = new(SingleVarDeclNoExpsASTContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *SingleVarDeclNoExpsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclNoExpsASTContext) SingleVarDeclNoExps() ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *SingleVarDeclNoExpsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSingleVarDeclNoExpsAST(s)
	}
}

func (s *SingleVarDeclNoExpsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSingleVarDeclNoExpsAST(s)
	}
}

func (s *SingleVarDeclNoExpsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSingleVarDeclNoExpsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) SingleVarDecl() (localctx ISingleVarDeclContext) {
	localctx = NewSingleVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, miniGoParserRULE_singleVarDecl)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleVarDeclASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.IdentifierList()
		}
		{
			p.SetState(122)
			p.DeclType()
		}
		{
			p.SetState(123)
			p.Match(miniGoParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.ExpressionList()
		}

	case 2:
		localctx = NewSingleVarDeclAssignASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.IdentifierList()
		}
		{
			p.SetState(127)
			p.Match(miniGoParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.ExpressionList()
		}

	case 3:
		localctx = NewSingleVarDeclNoExpsASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.SingleVarDeclNoExps()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleVarDeclNoExpsContext is an interface to support dynamic dispatch.
type ISingleVarDeclNoExpsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingleVarDeclNoExpsContext differentiates from other interfaces.
	IsSingleVarDeclNoExpsContext()
}

type SingleVarDeclNoExpsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleVarDeclNoExpsContext() *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_singleVarDeclNoExps
	return p
}

func InitEmptySingleVarDeclNoExpsContext(p *SingleVarDeclNoExpsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_singleVarDeclNoExps
}

func (*SingleVarDeclNoExpsContext) IsSingleVarDeclNoExpsContext() {}

func NewSingleVarDeclNoExpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_singleVarDeclNoExps

	return p
}

func (s *SingleVarDeclNoExpsContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclNoExpsContext) CopyAll(ctx *SingleVarDeclNoExpsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleVarDeclNoExpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclNoExpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleVarDeclNoExpsDeclTypeASTContext struct {
	SingleVarDeclNoExpsContext
}

func NewSingleVarDeclNoExpsDeclTypeASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleVarDeclNoExpsDeclTypeASTContext {
	var p = new(SingleVarDeclNoExpsDeclTypeASTContext)

	InitEmptySingleVarDeclNoExpsContext(&p.SingleVarDeclNoExpsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclNoExpsContext))

	return p
}

func (s *SingleVarDeclNoExpsDeclTypeASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclNoExpsDeclTypeASTContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *SingleVarDeclNoExpsDeclTypeASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleVarDeclNoExpsDeclTypeASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSingleVarDeclNoExpsDeclTypeAST(s)
	}
}

func (s *SingleVarDeclNoExpsDeclTypeASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSingleVarDeclNoExpsDeclTypeAST(s)
	}
}

func (s *SingleVarDeclNoExpsDeclTypeASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSingleVarDeclNoExpsDeclTypeAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) SingleVarDeclNoExps() (localctx ISingleVarDeclNoExpsContext) {
	localctx = NewSingleVarDeclNoExpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, miniGoParserRULE_singleVarDeclNoExps)
	localctx = NewSingleVarDeclNoExpsDeclTypeASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.IdentifierList()
	}
	{
		p.SetState(134)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_typeDecl
	return p
}

func InitEmptyTypeDeclContext(p *TypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_typeDecl
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) CopyAll(ctx *TypeDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeDeclEmptyASTContext struct {
	TypeDeclContext
}

func NewTypeDeclEmptyASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclEmptyASTContext {
	var p = new(TypeDeclEmptyASTContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *TypeDeclEmptyASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclEmptyASTContext) TYPE() antlr.TerminalNode {
	return s.GetToken(miniGoParserTYPE, 0)
}

func (s *TypeDeclEmptyASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *TypeDeclEmptyASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *TypeDeclEmptyASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *TypeDeclEmptyASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterTypeDeclEmptyAST(s)
	}
}

func (s *TypeDeclEmptyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitTypeDeclEmptyAST(s)
	}
}

func (s *TypeDeclEmptyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitTypeDeclEmptyAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeDeclASTContext struct {
	TypeDeclContext
}

func NewTypeDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclASTContext {
	var p = new(TypeDeclASTContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *TypeDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclASTContext) TYPE() antlr.TerminalNode {
	return s.GetToken(miniGoParserTYPE, 0)
}

func (s *TypeDeclASTContext) SingleTypeDecl() ISingleTypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleTypeDeclContext)
}

func (s *TypeDeclASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *TypeDeclASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterTypeDeclAST(s)
	}
}

func (s *TypeDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitTypeDeclAST(s)
	}
}

func (s *TypeDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitTypeDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeDeclBlockASTContext struct {
	TypeDeclContext
}

func NewTypeDeclBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclBlockASTContext {
	var p = new(TypeDeclBlockASTContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *TypeDeclBlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclBlockASTContext) TYPE() antlr.TerminalNode {
	return s.GetToken(miniGoParserTYPE, 0)
}

func (s *TypeDeclBlockASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *TypeDeclBlockASTContext) InnerTypeDecls() IInnerTypeDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerTypeDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerTypeDeclsContext)
}

func (s *TypeDeclBlockASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *TypeDeclBlockASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *TypeDeclBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterTypeDeclBlockAST(s)
	}
}

func (s *TypeDeclBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitTypeDeclBlockAST(s)
	}
}

func (s *TypeDeclBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitTypeDeclBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, miniGoParserRULE_typeDecl)
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeDeclASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Match(miniGoParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.SingleTypeDecl()
		}
		{
			p.SetState(138)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewTypeDeclBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Match(miniGoParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(miniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.InnerTypeDecls()
		}
		{
			p.SetState(143)
			p.Match(miniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewTypeDeclEmptyASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.Match(miniGoParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(miniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Match(miniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInnerTypeDeclsContext is an interface to support dynamic dispatch.
type IInnerTypeDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInnerTypeDeclsContext differentiates from other interfaces.
	IsInnerTypeDeclsContext()
}

type InnerTypeDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerTypeDeclsContext() *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_innerTypeDecls
	return p
}

func InitEmptyInnerTypeDeclsContext(p *InnerTypeDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_innerTypeDecls
}

func (*InnerTypeDeclsContext) IsInnerTypeDeclsContext() {}

func NewInnerTypeDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_innerTypeDecls

	return p
}

func (s *InnerTypeDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerTypeDeclsContext) CopyAll(ctx *InnerTypeDeclsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InnerTypeDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InnerTypeDeclsASTContext struct {
	InnerTypeDeclsContext
}

func NewInnerTypeDeclsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InnerTypeDeclsASTContext {
	var p = new(InnerTypeDeclsASTContext)

	InitEmptyInnerTypeDeclsContext(&p.InnerTypeDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InnerTypeDeclsContext))

	return p
}

func (s *InnerTypeDeclsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeDeclsASTContext) AllSingleTypeDecl() []ISingleTypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ISingleTypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleTypeDeclContext); ok {
			tst[i] = t.(ISingleTypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *InnerTypeDeclsASTContext) SingleTypeDecl(i int) ISingleTypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleTypeDeclContext)
}

func (s *InnerTypeDeclsASTContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(miniGoParserSEMICOLON)
}

func (s *InnerTypeDeclsASTContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, i)
}

func (s *InnerTypeDeclsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterInnerTypeDeclsAST(s)
	}
}

func (s *InnerTypeDeclsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitInnerTypeDeclsAST(s)
	}
}

func (s *InnerTypeDeclsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitInnerTypeDeclsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) InnerTypeDecls() (localctx IInnerTypeDeclsContext) {
	localctx = NewInnerTypeDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, miniGoParserRULE_innerTypeDecls)
	var _la int

	localctx = NewInnerTypeDeclsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.SingleTypeDecl()
	}
	{
		p.SetState(153)
		p.Match(miniGoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == miniGoParserIDENTIFIER {
		{
			p.SetState(154)
			p.SingleTypeDecl()
		}
		{
			p.SetState(155)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleTypeDeclContext is an interface to support dynamic dispatch.
type ISingleTypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingleTypeDeclContext differentiates from other interfaces.
	IsSingleTypeDeclContext()
}

type SingleTypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleTypeDeclContext() *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_singleTypeDecl
	return p
}

func InitEmptySingleTypeDeclContext(p *SingleTypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_singleTypeDecl
}

func (*SingleTypeDeclContext) IsSingleTypeDeclContext() {}

func NewSingleTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_singleTypeDecl

	return p
}

func (s *SingleTypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleTypeDeclContext) CopyAll(ctx *SingleTypeDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleTypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleTypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleTypeDeclASTContext struct {
	SingleTypeDeclContext
}

func NewSingleTypeDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleTypeDeclASTContext {
	var p = new(SingleTypeDeclASTContext)

	InitEmptySingleTypeDeclContext(&p.SingleTypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleTypeDeclContext))

	return p
}

func (s *SingleTypeDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleTypeDeclASTContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(miniGoParserIDENTIFIER, 0)
}

func (s *SingleTypeDeclASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleTypeDeclASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSingleTypeDeclAST(s)
	}
}

func (s *SingleTypeDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSingleTypeDeclAST(s)
	}
}

func (s *SingleTypeDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSingleTypeDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) SingleTypeDecl() (localctx ISingleTypeDeclContext) {
	localctx = NewSingleTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, miniGoParserRULE_singleTypeDecl)
	localctx = NewSingleTypeDeclASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(miniGoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_funcDecl
	return p
}

func InitEmptyFuncDeclContext(p *FuncDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_funcDecl
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) CopyAll(ctx *FuncDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncDeclASTContext struct {
	FuncDeclContext
}

func NewFuncDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncDeclASTContext {
	var p = new(FuncDeclASTContext)

	InitEmptyFuncDeclContext(&p.FuncDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncDeclContext))

	return p
}

func (s *FuncDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclASTContext) FuncFrontDecl() IFuncFrontDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncFrontDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncFrontDeclContext)
}

func (s *FuncDeclASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *FuncDeclASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterFuncDeclAST(s)
	}
}

func (s *FuncDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitFuncDeclAST(s)
	}
}

func (s *FuncDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitFuncDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, miniGoParserRULE_funcDecl)
	localctx = NewFuncDeclASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.FuncFrontDecl()
	}
	{
		p.SetState(166)
		p.Block()
	}
	{
		p.SetState(167)
		p.Match(miniGoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncFrontDeclContext is an interface to support dynamic dispatch.
type IFuncFrontDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncFrontDeclContext differentiates from other interfaces.
	IsFuncFrontDeclContext()
}

type FuncFrontDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncFrontDeclContext() *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_funcFrontDecl
	return p
}

func InitEmptyFuncFrontDeclContext(p *FuncFrontDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_funcFrontDecl
}

func (*FuncFrontDeclContext) IsFuncFrontDeclContext() {}

func NewFuncFrontDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_funcFrontDecl

	return p
}

func (s *FuncFrontDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncFrontDeclContext) CopyAll(ctx *FuncFrontDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncFrontDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFrontDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncFrontDeclASTContext struct {
	FuncFrontDeclContext
}

func NewFuncFrontDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncFrontDeclASTContext {
	var p = new(FuncFrontDeclASTContext)

	InitEmptyFuncFrontDeclContext(&p.FuncFrontDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncFrontDeclContext))

	return p
}

func (s *FuncFrontDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFrontDeclASTContext) FUNC() antlr.TerminalNode {
	return s.GetToken(miniGoParserFUNC, 0)
}

func (s *FuncFrontDeclASTContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(miniGoParserIDENTIFIER, 0)
}

func (s *FuncFrontDeclASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *FuncFrontDeclASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *FuncFrontDeclASTContext) FuncArgDecls() IFuncArgDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncArgDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncArgDeclsContext)
}

func (s *FuncFrontDeclASTContext) AllEpsilon() []IEpsilonContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEpsilonContext); ok {
			len++
		}
	}

	tst := make([]IEpsilonContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEpsilonContext); ok {
			tst[i] = t.(IEpsilonContext)
			i++
		}
	}

	return tst
}

func (s *FuncFrontDeclASTContext) Epsilon(i int) IEpsilonContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *FuncFrontDeclASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *FuncFrontDeclASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterFuncFrontDeclAST(s)
	}
}

func (s *FuncFrontDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitFuncFrontDeclAST(s)
	}
}

func (s *FuncFrontDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitFuncFrontDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) FuncFrontDecl() (localctx IFuncFrontDeclContext) {
	localctx = NewFuncFrontDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, miniGoParserRULE_funcFrontDecl)
	localctx = NewFuncFrontDeclASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(miniGoParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(miniGoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(miniGoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniGoParserIDENTIFIER:
		{
			p.SetState(172)
			p.FuncArgDecls()
		}

	case miniGoParserRPAREN:
		{
			p.SetState(173)
			p.Epsilon()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(176)
		p.Match(miniGoParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniGoParserSTRUCT, miniGoParserIDENTIFIER, miniGoParserLPAREN, miniGoParserLBRACKET:
		{
			p.SetState(177)
			p.DeclType()
		}

	case miniGoParserLBRACE:
		{
			p.SetState(178)
			p.Epsilon()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncArgDeclsContext is an interface to support dynamic dispatch.
type IFuncArgDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncArgDeclsContext differentiates from other interfaces.
	IsFuncArgDeclsContext()
}

type FuncArgDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgDeclsContext() *FuncArgDeclsContext {
	var p = new(FuncArgDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_funcArgDecls
	return p
}

func InitEmptyFuncArgDeclsContext(p *FuncArgDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_funcArgDecls
}

func (*FuncArgDeclsContext) IsFuncArgDeclsContext() {}

func NewFuncArgDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgDeclsContext {
	var p = new(FuncArgDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_funcArgDecls

	return p
}

func (s *FuncArgDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgDeclsContext) CopyAll(ctx *FuncArgDeclsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncArgDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncArgDeclsASTContext struct {
	FuncArgDeclsContext
}

func NewFuncArgDeclsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncArgDeclsASTContext {
	var p = new(FuncArgDeclsASTContext)

	InitEmptyFuncArgDeclsContext(&p.FuncArgDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncArgDeclsContext))

	return p
}

func (s *FuncArgDeclsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgDeclsASTContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclNoExpsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			tst[i] = t.(ISingleVarDeclNoExpsContext)
			i++
		}
	}

	return tst
}

func (s *FuncArgDeclsASTContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *FuncArgDeclsASTContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(miniGoParserCOMMA)
}

func (s *FuncArgDeclsASTContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(miniGoParserCOMMA, i)
}

func (s *FuncArgDeclsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterFuncArgDeclsAST(s)
	}
}

func (s *FuncArgDeclsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitFuncArgDeclsAST(s)
	}
}

func (s *FuncArgDeclsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitFuncArgDeclsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) FuncArgDecls() (localctx IFuncArgDeclsContext) {
	localctx = NewFuncArgDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, miniGoParserRULE_funcArgDecls)
	var _la int

	localctx = NewFuncArgDeclsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.SingleVarDeclNoExps()
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == miniGoParserCOMMA {
		{
			p.SetState(182)
			p.Match(miniGoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.SingleVarDeclNoExps()
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclTypeContext is an interface to support dynamic dispatch.
type IDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclTypeContext differentiates from other interfaces.
	IsDeclTypeContext()
}

type DeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclTypeContext() *DeclTypeContext {
	var p = new(DeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_declType
	return p
}

func InitEmptyDeclTypeContext(p *DeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_declType
}

func (*DeclTypeContext) IsDeclTypeContext() {}

func NewDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclTypeContext {
	var p = new(DeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_declType

	return p
}

func (s *DeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclTypeContext) CopyAll(ctx *DeclTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclTypeStructASTContext struct {
	DeclTypeContext
}

func NewDeclTypeStructASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclTypeStructASTContext {
	var p = new(DeclTypeStructASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *DeclTypeStructASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeStructASTContext) StructDeclType() IStructDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclTypeContext)
}

func (s *DeclTypeStructASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterDeclTypeStructAST(s)
	}
}

func (s *DeclTypeStructASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitDeclTypeStructAST(s)
	}
}

func (s *DeclTypeStructASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitDeclTypeStructAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclTypeSliceASTContext struct {
	DeclTypeContext
}

func NewDeclTypeSliceASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclTypeSliceASTContext {
	var p = new(DeclTypeSliceASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *DeclTypeSliceASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeSliceASTContext) SliceDeclType() ISliceDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceDeclTypeContext)
}

func (s *DeclTypeSliceASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterDeclTypeSliceAST(s)
	}
}

func (s *DeclTypeSliceASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitDeclTypeSliceAST(s)
	}
}

func (s *DeclTypeSliceASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitDeclTypeSliceAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclTypeArrayASTContext struct {
	DeclTypeContext
}

func NewDeclTypeArrayASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclTypeArrayASTContext {
	var p = new(DeclTypeArrayASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *DeclTypeArrayASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeArrayASTContext) ArrayDeclType() IArrayDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayDeclTypeContext)
}

func (s *DeclTypeArrayASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterDeclTypeArrayAST(s)
	}
}

func (s *DeclTypeArrayASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitDeclTypeArrayAST(s)
	}
}

func (s *DeclTypeArrayASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitDeclTypeArrayAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclTypeParenASTContext struct {
	DeclTypeContext
}

func NewDeclTypeParenASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclTypeParenASTContext {
	var p = new(DeclTypeParenASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *DeclTypeParenASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeParenASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *DeclTypeParenASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *DeclTypeParenASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *DeclTypeParenASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterDeclTypeParenAST(s)
	}
}

func (s *DeclTypeParenASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitDeclTypeParenAST(s)
	}
}

func (s *DeclTypeParenASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitDeclTypeParenAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclTypeIdentifierASTContext struct {
	DeclTypeContext
}

func NewDeclTypeIdentifierASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclTypeIdentifierASTContext {
	var p = new(DeclTypeIdentifierASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *DeclTypeIdentifierASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeIdentifierASTContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(miniGoParserIDENTIFIER, 0)
}

func (s *DeclTypeIdentifierASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterDeclTypeIdentifierAST(s)
	}
}

func (s *DeclTypeIdentifierASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitDeclTypeIdentifierAST(s)
	}
}

func (s *DeclTypeIdentifierASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitDeclTypeIdentifierAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) DeclType() (localctx IDeclTypeContext) {
	localctx = NewDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, miniGoParserRULE_declType)
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclTypeParenASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Match(miniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.DeclType()
		}
		{
			p.SetState(191)
			p.Match(miniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDeclTypeIdentifierASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Match(miniGoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewDeclTypeSliceASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(194)
			p.SliceDeclType()
		}

	case 4:
		localctx = NewDeclTypeArrayASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(195)
			p.ArrayDeclType()
		}

	case 5:
		localctx = NewDeclTypeStructASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(196)
			p.StructDeclType()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceDeclTypeContext is an interface to support dynamic dispatch.
type ISliceDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSliceDeclTypeContext differentiates from other interfaces.
	IsSliceDeclTypeContext()
}

type SliceDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceDeclTypeContext() *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_sliceDeclType
	return p
}

func InitEmptySliceDeclTypeContext(p *SliceDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_sliceDeclType
}

func (*SliceDeclTypeContext) IsSliceDeclTypeContext() {}

func NewSliceDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_sliceDeclType

	return p
}

func (s *SliceDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceDeclTypeContext) CopyAll(ctx *SliceDeclTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SliceDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SliceDeclTypeASTContext struct {
	SliceDeclTypeContext
}

func NewSliceDeclTypeASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceDeclTypeASTContext {
	var p = new(SliceDeclTypeASTContext)

	InitEmptySliceDeclTypeContext(&p.SliceDeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*SliceDeclTypeContext))

	return p
}

func (s *SliceDeclTypeASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDeclTypeASTContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(miniGoParserLBRACKET, 0)
}

func (s *SliceDeclTypeASTContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(miniGoParserRBRACKET, 0)
}

func (s *SliceDeclTypeASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SliceDeclTypeASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSliceDeclTypeAST(s)
	}
}

func (s *SliceDeclTypeASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSliceDeclTypeAST(s)
	}
}

func (s *SliceDeclTypeASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSliceDeclTypeAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) SliceDeclType() (localctx ISliceDeclTypeContext) {
	localctx = NewSliceDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, miniGoParserRULE_sliceDeclType)
	localctx = NewSliceDeclTypeASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(miniGoParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.Match(miniGoParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayDeclTypeContext is an interface to support dynamic dispatch.
type IArrayDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArrayDeclTypeContext differentiates from other interfaces.
	IsArrayDeclTypeContext()
}

type ArrayDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDeclTypeContext() *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_arrayDeclType
	return p
}

func InitEmptyArrayDeclTypeContext(p *ArrayDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_arrayDeclType
}

func (*ArrayDeclTypeContext) IsArrayDeclTypeContext() {}

func NewArrayDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_arrayDeclType

	return p
}

func (s *ArrayDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDeclTypeContext) CopyAll(ctx *ArrayDeclTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArrayDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayDeclTypeASTContext struct {
	ArrayDeclTypeContext
}

func NewArrayDeclTypeASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayDeclTypeASTContext {
	var p = new(ArrayDeclTypeASTContext)

	InitEmptyArrayDeclTypeContext(&p.ArrayDeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayDeclTypeContext))

	return p
}

func (s *ArrayDeclTypeASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclTypeASTContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(miniGoParserLBRACKET, 0)
}

func (s *ArrayDeclTypeASTContext) INTLITERAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserINTLITERAL, 0)
}

func (s *ArrayDeclTypeASTContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(miniGoParserRBRACKET, 0)
}

func (s *ArrayDeclTypeASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *ArrayDeclTypeASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterArrayDeclTypeAST(s)
	}
}

func (s *ArrayDeclTypeASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitArrayDeclTypeAST(s)
	}
}

func (s *ArrayDeclTypeASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitArrayDeclTypeAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) ArrayDeclType() (localctx IArrayDeclTypeContext) {
	localctx = NewArrayDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, miniGoParserRULE_arrayDeclType)
	localctx = NewArrayDeclTypeASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(miniGoParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Match(miniGoParserINTLITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(miniGoParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclTypeContext is an interface to support dynamic dispatch.
type IStructDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStructDeclTypeContext differentiates from other interfaces.
	IsStructDeclTypeContext()
}

type StructDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclTypeContext() *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_structDeclType
	return p
}

func InitEmptyStructDeclTypeContext(p *StructDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_structDeclType
}

func (*StructDeclTypeContext) IsStructDeclTypeContext() {}

func NewStructDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_structDeclType

	return p
}

func (s *StructDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclTypeContext) CopyAll(ctx *StructDeclTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StructDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructDeclTypeASTContext struct {
	StructDeclTypeContext
}

func NewStructDeclTypeASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclTypeASTContext {
	var p = new(StructDeclTypeASTContext)

	InitEmptyStructDeclTypeContext(&p.StructDeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructDeclTypeContext))

	return p
}

func (s *StructDeclTypeASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclTypeASTContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(miniGoParserSTRUCT, 0)
}

func (s *StructDeclTypeASTContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(miniGoParserLBRACE, 0)
}

func (s *StructDeclTypeASTContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(miniGoParserRBRACE, 0)
}

func (s *StructDeclTypeASTContext) StructMemDecls() IStructMemDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructMemDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructMemDeclsContext)
}

func (s *StructDeclTypeASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *StructDeclTypeASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStructDeclTypeAST(s)
	}
}

func (s *StructDeclTypeASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStructDeclTypeAST(s)
	}
}

func (s *StructDeclTypeASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStructDeclTypeAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) StructDeclType() (localctx IStructDeclTypeContext) {
	localctx = NewStructDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, miniGoParserRULE_structDeclType)
	localctx = NewStructDeclTypeASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(miniGoParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(miniGoParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniGoParserIDENTIFIER:
		{
			p.SetState(210)
			p.StructMemDecls()
		}

	case miniGoParserRBRACE:
		{
			p.SetState(211)
			p.Epsilon()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(214)
		p.Match(miniGoParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructMemDeclsContext is an interface to support dynamic dispatch.
type IStructMemDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStructMemDeclsContext differentiates from other interfaces.
	IsStructMemDeclsContext()
}

type StructMemDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructMemDeclsContext() *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_structMemDecls
	return p
}

func InitEmptyStructMemDeclsContext(p *StructMemDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_structMemDecls
}

func (*StructMemDeclsContext) IsStructMemDeclsContext() {}

func NewStructMemDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_structMemDecls

	return p
}

func (s *StructMemDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *StructMemDeclsContext) CopyAll(ctx *StructMemDeclsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StructMemDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructMemDeclsASTContext struct {
	StructMemDeclsContext
}

func NewStructMemDeclsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructMemDeclsASTContext {
	var p = new(StructMemDeclsASTContext)

	InitEmptyStructMemDeclsContext(&p.StructMemDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructMemDeclsContext))

	return p
}

func (s *StructMemDeclsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemDeclsASTContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclNoExpsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			tst[i] = t.(ISingleVarDeclNoExpsContext)
			i++
		}
	}

	return tst
}

func (s *StructMemDeclsASTContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *StructMemDeclsASTContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(miniGoParserSEMICOLON)
}

func (s *StructMemDeclsASTContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, i)
}

func (s *StructMemDeclsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStructMemDeclsAST(s)
	}
}

func (s *StructMemDeclsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStructMemDeclsAST(s)
	}
}

func (s *StructMemDeclsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStructMemDeclsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) StructMemDecls() (localctx IStructMemDeclsContext) {
	localctx = NewStructMemDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, miniGoParserRULE_structMemDecls)
	var _la int

	localctx = NewStructMemDeclsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.SingleVarDeclNoExps()
	}
	{
		p.SetState(217)
		p.Match(miniGoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == miniGoParserIDENTIFIER {
		{
			p.SetState(218)
			p.SingleVarDeclNoExps()
		}
		{
			p.SetState(219)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_identifierList
	return p
}

func InitEmptyIdentifierListContext(p *IdentifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_identifierList
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) CopyAll(ctx *IdentifierListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifierListASTContext struct {
	IdentifierListContext
}

func NewIdentifierListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierListASTContext {
	var p = new(IdentifierListASTContext)

	InitEmptyIdentifierListContext(&p.IdentifierListContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentifierListContext))

	return p
}

func (s *IdentifierListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListASTContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(miniGoParserIDENTIFIER)
}

func (s *IdentifierListASTContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(miniGoParserIDENTIFIER, i)
}

func (s *IdentifierListASTContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(miniGoParserCOMMA)
}

func (s *IdentifierListASTContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(miniGoParserCOMMA, i)
}

func (s *IdentifierListASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterIdentifierListAST(s)
	}
}

func (s *IdentifierListASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitIdentifierListAST(s)
	}
}

func (s *IdentifierListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitIdentifierListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, miniGoParserRULE_identifierList)
	var _la int

	localctx = NewIdentifierListASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(miniGoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == miniGoParserCOMMA {
		{
			p.SetState(227)
			p.Match(miniGoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Match(miniGoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionNotUnaryASTContext struct {
	ExpressionContext
}

func NewExpressionNotUnaryASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionNotUnaryASTContext {
	var p = new(ExpressionNotUnaryASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionNotUnaryASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionNotUnaryASTContext) NOT() antlr.TerminalNode {
	return s.GetToken(miniGoParserNOT, 0)
}

func (s *ExpressionNotUnaryASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionNotUnaryASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionNotUnaryAST(s)
	}
}

func (s *ExpressionNotUnaryASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionNotUnaryAST(s)
	}
}

func (s *ExpressionNotUnaryASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionNotUnaryAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionMultiplyASTContext struct {
	ExpressionContext
}

func NewExpressionMultiplyASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionMultiplyASTContext {
	var p = new(ExpressionMultiplyASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionMultiplyASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionMultiplyASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionMultiplyASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionMultiplyASTContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(miniGoParserMULTIPLY, 0)
}

func (s *ExpressionMultiplyASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionMultiplyAST(s)
	}
}

func (s *ExpressionMultiplyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionMultiplyAST(s)
	}
}

func (s *ExpressionMultiplyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionMultiplyAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionPlusASTContext struct {
	ExpressionContext
}

func NewExpressionPlusASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionPlusASTContext {
	var p = new(ExpressionPlusASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionPlusASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionPlusASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionPlusASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionPlusASTContext) PLUS() antlr.TerminalNode {
	return s.GetToken(miniGoParserPLUS, 0)
}

func (s *ExpressionPlusASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionPlusAST(s)
	}
}

func (s *ExpressionPlusASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionPlusAST(s)
	}
}

func (s *ExpressionPlusASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionPlusAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionModuloASTContext struct {
	ExpressionContext
}

func NewExpressionModuloASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionModuloASTContext {
	var p = new(ExpressionModuloASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionModuloASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionModuloASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionModuloASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionModuloASTContext) MODULO() antlr.TerminalNode {
	return s.GetToken(miniGoParserMODULO, 0)
}

func (s *ExpressionModuloASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionModuloAST(s)
	}
}

func (s *ExpressionModuloASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionModuloAST(s)
	}
}

func (s *ExpressionModuloASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionModuloAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionAndASTContext struct {
	ExpressionContext
}

func NewExpressionAndASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionAndASTContext {
	var p = new(ExpressionAndASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionAndASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAndASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionAndASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionAndASTContext) AND() antlr.TerminalNode {
	return s.GetToken(miniGoParserAND, 0)
}

func (s *ExpressionAndASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionAndAST(s)
	}
}

func (s *ExpressionAndASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionAndAST(s)
	}
}

func (s *ExpressionAndASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionAndAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionBitwiseXorASTContext struct {
	ExpressionContext
}

func NewExpressionBitwiseXorASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionBitwiseXorASTContext {
	var p = new(ExpressionBitwiseXorASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionBitwiseXorASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionBitwiseXorASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionBitwiseXorASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionBitwiseXorASTContext) BITWISEXOR() antlr.TerminalNode {
	return s.GetToken(miniGoParserBITWISEXOR, 0)
}

func (s *ExpressionBitwiseXorASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionBitwiseXorAST(s)
	}
}

func (s *ExpressionBitwiseXorASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionBitwiseXorAST(s)
	}
}

func (s *ExpressionBitwiseXorASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionBitwiseXorAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionMinusASTContext struct {
	ExpressionContext
}

func NewExpressionMinusASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionMinusASTContext {
	var p = new(ExpressionMinusASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionMinusASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionMinusASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionMinusASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionMinusASTContext) MINUS() antlr.TerminalNode {
	return s.GetToken(miniGoParserMINUS, 0)
}

func (s *ExpressionMinusASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionMinusAST(s)
	}
}

func (s *ExpressionMinusASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionMinusAST(s)
	}
}

func (s *ExpressionMinusASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionMinusAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionBitwiseXorUnaryASTContext struct {
	ExpressionContext
}

func NewExpressionBitwiseXorUnaryASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionBitwiseXorUnaryASTContext {
	var p = new(ExpressionBitwiseXorUnaryASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionBitwiseXorUnaryASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionBitwiseXorUnaryASTContext) BITWISEXOR() antlr.TerminalNode {
	return s.GetToken(miniGoParserBITWISEXOR, 0)
}

func (s *ExpressionBitwiseXorUnaryASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionBitwiseXorUnaryASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionBitwiseXorUnaryAST(s)
	}
}

func (s *ExpressionBitwiseXorUnaryASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionBitwiseXorUnaryAST(s)
	}
}

func (s *ExpressionBitwiseXorUnaryASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionBitwiseXorUnaryAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionEqualASTContext struct {
	ExpressionContext
}

func NewExpressionEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionEqualASTContext {
	var p = new(ExpressionEqualASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionEqualASTContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserEQUAL, 0)
}

func (s *ExpressionEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionEqualAST(s)
	}
}

func (s *ExpressionEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionEqualAST(s)
	}
}

func (s *ExpressionEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionMinusUnaryASTContext struct {
	ExpressionContext
}

func NewExpressionMinusUnaryASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionMinusUnaryASTContext {
	var p = new(ExpressionMinusUnaryASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionMinusUnaryASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionMinusUnaryASTContext) MINUS() antlr.TerminalNode {
	return s.GetToken(miniGoParserMINUS, 0)
}

func (s *ExpressionMinusUnaryASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionMinusUnaryASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionMinusUnaryAST(s)
	}
}

func (s *ExpressionMinusUnaryASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionMinusUnaryAST(s)
	}
}

func (s *ExpressionMinusUnaryASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionMinusUnaryAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionBitwiseClearASTContext struct {
	ExpressionContext
}

func NewExpressionBitwiseClearASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionBitwiseClearASTContext {
	var p = new(ExpressionBitwiseClearASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionBitwiseClearASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionBitwiseClearASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionBitwiseClearASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionBitwiseClearASTContext) BITWISECLEAR() antlr.TerminalNode {
	return s.GetToken(miniGoParserBITWISECLEAR, 0)
}

func (s *ExpressionBitwiseClearASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionBitwiseClearAST(s)
	}
}

func (s *ExpressionBitwiseClearASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionBitwiseClearAST(s)
	}
}

func (s *ExpressionBitwiseClearASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionBitwiseClearAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionGreaterEqualASTContext struct {
	ExpressionContext
}

func NewExpressionGreaterEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionGreaterEqualASTContext {
	var p = new(ExpressionGreaterEqualASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionGreaterEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionGreaterEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionGreaterEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionGreaterEqualASTContext) GREATEREQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserGREATEREQUAL, 0)
}

func (s *ExpressionGreaterEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionGreaterEqualAST(s)
	}
}

func (s *ExpressionGreaterEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionGreaterEqualAST(s)
	}
}

func (s *ExpressionGreaterEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionGreaterEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionLessEqualASTContext struct {
	ExpressionContext
}

func NewExpressionLessEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionLessEqualASTContext {
	var p = new(ExpressionLessEqualASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionLessEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionLessEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionLessEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionLessEqualASTContext) LESSEQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserLESSEQUAL, 0)
}

func (s *ExpressionLessEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionLessEqualAST(s)
	}
}

func (s *ExpressionLessEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionLessEqualAST(s)
	}
}

func (s *ExpressionLessEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionLessEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionNotEqualASTContext struct {
	ExpressionContext
}

func NewExpressionNotEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionNotEqualASTContext {
	var p = new(ExpressionNotEqualASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionNotEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionNotEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionNotEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionNotEqualASTContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserNOTEQUAL, 0)
}

func (s *ExpressionNotEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionNotEqualAST(s)
	}
}

func (s *ExpressionNotEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionNotEqualAST(s)
	}
}

func (s *ExpressionNotEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionNotEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionPrimaryASTContext struct {
	ExpressionContext
}

func NewExpressionPrimaryASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionPrimaryASTContext {
	var p = new(ExpressionPrimaryASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionPrimaryASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionPrimaryASTContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *ExpressionPrimaryASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionPrimaryAST(s)
	}
}

func (s *ExpressionPrimaryASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionPrimaryAST(s)
	}
}

func (s *ExpressionPrimaryASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionPrimaryAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionBitwiseAndASTContext struct {
	ExpressionContext
}

func NewExpressionBitwiseAndASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionBitwiseAndASTContext {
	var p = new(ExpressionBitwiseAndASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionBitwiseAndASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionBitwiseAndASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionBitwiseAndASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionBitwiseAndASTContext) BITWISEAND() antlr.TerminalNode {
	return s.GetToken(miniGoParserBITWISEAND, 0)
}

func (s *ExpressionBitwiseAndASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionBitwiseAndAST(s)
	}
}

func (s *ExpressionBitwiseAndASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionBitwiseAndAST(s)
	}
}

func (s *ExpressionBitwiseAndASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionBitwiseAndAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionGreaterASTContext struct {
	ExpressionContext
}

func NewExpressionGreaterASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionGreaterASTContext {
	var p = new(ExpressionGreaterASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionGreaterASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionGreaterASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionGreaterASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionGreaterASTContext) GREATER() antlr.TerminalNode {
	return s.GetToken(miniGoParserGREATER, 0)
}

func (s *ExpressionGreaterASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionGreaterAST(s)
	}
}

func (s *ExpressionGreaterASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionGreaterAST(s)
	}
}

func (s *ExpressionGreaterASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionGreaterAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionDivideASTContext struct {
	ExpressionContext
}

func NewExpressionDivideASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionDivideASTContext {
	var p = new(ExpressionDivideASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionDivideASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionDivideASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionDivideASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionDivideASTContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(miniGoParserDIVIDE, 0)
}

func (s *ExpressionDivideASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionDivideAST(s)
	}
}

func (s *ExpressionDivideASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionDivideAST(s)
	}
}

func (s *ExpressionDivideASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionDivideAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionOrASTContext struct {
	ExpressionContext
}

func NewExpressionOrASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionOrASTContext {
	var p = new(ExpressionOrASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionOrASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionOrASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionOrASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionOrASTContext) OR() antlr.TerminalNode {
	return s.GetToken(miniGoParserOR, 0)
}

func (s *ExpressionOrASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionOrAST(s)
	}
}

func (s *ExpressionOrASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionOrAST(s)
	}
}

func (s *ExpressionOrASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionOrAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionPlusUnaryASTContext struct {
	ExpressionContext
}

func NewExpressionPlusUnaryASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionPlusUnaryASTContext {
	var p = new(ExpressionPlusUnaryASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionPlusUnaryASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionPlusUnaryASTContext) PLUS() antlr.TerminalNode {
	return s.GetToken(miniGoParserPLUS, 0)
}

func (s *ExpressionPlusUnaryASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionPlusUnaryASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionPlusUnaryAST(s)
	}
}

func (s *ExpressionPlusUnaryASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionPlusUnaryAST(s)
	}
}

func (s *ExpressionPlusUnaryASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionPlusUnaryAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionBitwiseOrASTContext struct {
	ExpressionContext
}

func NewExpressionBitwiseOrASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionBitwiseOrASTContext {
	var p = new(ExpressionBitwiseOrASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionBitwiseOrASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionBitwiseOrASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionBitwiseOrASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionBitwiseOrASTContext) BITWISEOR() antlr.TerminalNode {
	return s.GetToken(miniGoParserBITWISEOR, 0)
}

func (s *ExpressionBitwiseOrASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionBitwiseOrAST(s)
	}
}

func (s *ExpressionBitwiseOrASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionBitwiseOrAST(s)
	}
}

func (s *ExpressionBitwiseOrASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionBitwiseOrAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionLessASTContext struct {
	ExpressionContext
}

func NewExpressionLessASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionLessASTContext {
	var p = new(ExpressionLessASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionLessASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionLessASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionLessASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionLessASTContext) LESS() antlr.TerminalNode {
	return s.GetToken(miniGoParserLESS, 0)
}

func (s *ExpressionLessASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionLessAST(s)
	}
}

func (s *ExpressionLessASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionLessAST(s)
	}
}

func (s *ExpressionLessASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionLessAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionShiftRightASTContext struct {
	ExpressionContext
}

func NewExpressionShiftRightASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionShiftRightASTContext {
	var p = new(ExpressionShiftRightASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionShiftRightASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionShiftRightASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionShiftRightASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionShiftRightASTContext) SHIFTRIGHT() antlr.TerminalNode {
	return s.GetToken(miniGoParserSHIFTRIGHT, 0)
}

func (s *ExpressionShiftRightASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionShiftRightAST(s)
	}
}

func (s *ExpressionShiftRightASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionShiftRightAST(s)
	}
}

func (s *ExpressionShiftRightASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionShiftRightAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionShiftLeftASTContext struct {
	ExpressionContext
}

func NewExpressionShiftLeftASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionShiftLeftASTContext {
	var p = new(ExpressionShiftLeftASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionShiftLeftASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionShiftLeftASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionShiftLeftASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionShiftLeftASTContext) SHIFTLEFT() antlr.TerminalNode {
	return s.GetToken(miniGoParserSHIFTLEFT, 0)
}

func (s *ExpressionShiftLeftASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionShiftLeftAST(s)
	}
}

func (s *ExpressionShiftLeftASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionShiftLeftAST(s)
	}
}

func (s *ExpressionShiftLeftASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionShiftLeftAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *miniGoParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, miniGoParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniGoParserINTLITERAL, miniGoParserFLOATLITERAL, miniGoParserRUNELITERAL, miniGoParserRAWSTRINGLITERAL, miniGoParserINTERPRETEDSTRINGLITERAL, miniGoParserLEN, miniGoParserCAP, miniGoParserAPPEND, miniGoParserIDENTIFIER, miniGoParserLPAREN:
		localctx = NewExpressionPrimaryASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(235)
			p.primaryExpression(0)
		}

	case miniGoParserPLUS:
		localctx = NewExpressionPlusUnaryASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(236)
			p.Match(miniGoParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.expression(4)
		}

	case miniGoParserMINUS:
		localctx = NewExpressionMinusUnaryASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.Match(miniGoParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.expression(3)
		}

	case miniGoParserNOT:
		localctx = NewExpressionNotUnaryASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(240)
			p.Match(miniGoParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.expression(2)
		}

	case miniGoParserBITWISEXOR:
		localctx = NewExpressionBitwiseXorUnaryASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(242)
			p.Match(miniGoParserBITWISEXOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.expression(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(303)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionMultiplyASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(246)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(247)
					p.Match(miniGoParserMULTIPLY)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(248)
					p.expression(24)
				}

			case 2:
				localctx = NewExpressionDivideASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(249)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(250)
					p.Match(miniGoParserDIVIDE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(251)
					p.expression(23)
				}

			case 3:
				localctx = NewExpressionModuloASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(252)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(253)
					p.Match(miniGoParserMODULO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(254)
					p.expression(22)
				}

			case 4:
				localctx = NewExpressionShiftLeftASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(255)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(256)
					p.Match(miniGoParserSHIFTLEFT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(257)
					p.expression(21)
				}

			case 5:
				localctx = NewExpressionShiftRightASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(258)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(259)
					p.Match(miniGoParserSHIFTRIGHT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(260)
					p.expression(20)
				}

			case 6:
				localctx = NewExpressionBitwiseAndASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(261)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(262)
					p.Match(miniGoParserBITWISEAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(263)
					p.expression(19)
				}

			case 7:
				localctx = NewExpressionBitwiseClearASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(264)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(265)
					p.Match(miniGoParserBITWISECLEAR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(266)
					p.expression(18)
				}

			case 8:
				localctx = NewExpressionPlusASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(267)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(268)
					p.Match(miniGoParserPLUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(269)
					p.expression(17)
				}

			case 9:
				localctx = NewExpressionMinusASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(270)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(271)
					p.Match(miniGoParserMINUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(272)
					p.expression(16)
				}

			case 10:
				localctx = NewExpressionBitwiseOrASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(273)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(274)
					p.Match(miniGoParserBITWISEOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(275)
					p.expression(15)
				}

			case 11:
				localctx = NewExpressionBitwiseXorASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(276)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(277)
					p.Match(miniGoParserBITWISEXOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(278)
					p.expression(14)
				}

			case 12:
				localctx = NewExpressionEqualASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(279)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(280)
					p.Match(miniGoParserEQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(281)
					p.expression(13)
				}

			case 13:
				localctx = NewExpressionNotEqualASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(282)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(283)
					p.Match(miniGoParserNOTEQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(284)
					p.expression(12)
				}

			case 14:
				localctx = NewExpressionLessASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(285)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(286)
					p.Match(miniGoParserLESS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(287)
					p.expression(11)
				}

			case 15:
				localctx = NewExpressionLessEqualASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(288)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(289)
					p.Match(miniGoParserLESSEQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(290)
					p.expression(10)
				}

			case 16:
				localctx = NewExpressionGreaterASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(291)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(292)
					p.Match(miniGoParserGREATER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(293)
					p.expression(9)
				}

			case 17:
				localctx = NewExpressionGreaterEqualASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(294)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(295)
					p.Match(miniGoParserGREATEREQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(296)
					p.expression(8)
				}

			case 18:
				localctx = NewExpressionAndASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(297)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(298)
					p.Match(miniGoParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(299)
					p.expression(7)
				}

			case 19:
				localctx = NewExpressionOrASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expression)
				p.SetState(300)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(301)
					p.Match(miniGoParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(302)
					p.expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) CopyAll(ctx *ExpressionListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionListASTContext struct {
	ExpressionListContext
}

func NewExpressionListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionListASTContext {
	var p = new(ExpressionListASTContext)

	InitEmptyExpressionListContext(&p.ExpressionListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionListContext))

	return p
}

func (s *ExpressionListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListASTContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(miniGoParserCOMMA)
}

func (s *ExpressionListASTContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(miniGoParserCOMMA, i)
}

func (s *ExpressionListASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionListAST(s)
	}
}

func (s *ExpressionListASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionListAST(s)
	}
}

func (s *ExpressionListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, miniGoParserRULE_expressionList)
	var _alt int

	localctx = NewExpressionListASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.expression(0)
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(309)
				p.Match(miniGoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(310)
				p.expression(0)
			}

		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) CopyAll(ctx *PrimaryExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrimaryExpressionLengthASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionLengthASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionLengthASTContext {
	var p = new(PrimaryExpressionLengthASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionLengthASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionLengthASTContext) LengthExpression() ILengthExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthExpressionContext)
}

func (s *PrimaryExpressionLengthASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionLengthAST(s)
	}
}

func (s *PrimaryExpressionLengthASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionLengthAST(s)
	}
}

func (s *PrimaryExpressionLengthASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitPrimaryExpressionLengthAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionOperandASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionOperandASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionOperandASTContext {
	var p = new(PrimaryExpressionOperandASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionOperandASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionOperandASTContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExpressionOperandASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionOperandAST(s)
	}
}

func (s *PrimaryExpressionOperandASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionOperandAST(s)
	}
}

func (s *PrimaryExpressionOperandASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitPrimaryExpressionOperandAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionIndexASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionIndexASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionIndexASTContext {
	var p = new(PrimaryExpressionIndexASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionIndexASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionIndexASTContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpressionIndexASTContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *PrimaryExpressionIndexASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionIndexAST(s)
	}
}

func (s *PrimaryExpressionIndexASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionIndexAST(s)
	}
}

func (s *PrimaryExpressionIndexASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitPrimaryExpressionIndexAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionAppendASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionAppendASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionAppendASTContext {
	var p = new(PrimaryExpressionAppendASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionAppendASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionAppendASTContext) AppendExpression() IAppendExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppendExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppendExpressionContext)
}

func (s *PrimaryExpressionAppendASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionAppendAST(s)
	}
}

func (s *PrimaryExpressionAppendASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionAppendAST(s)
	}
}

func (s *PrimaryExpressionAppendASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitPrimaryExpressionAppendAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionArgumentsASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionArgumentsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionArgumentsASTContext {
	var p = new(PrimaryExpressionArgumentsASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionArgumentsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionArgumentsASTContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpressionArgumentsASTContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *PrimaryExpressionArgumentsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionArgumentsAST(s)
	}
}

func (s *PrimaryExpressionArgumentsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionArgumentsAST(s)
	}
}

func (s *PrimaryExpressionArgumentsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitPrimaryExpressionArgumentsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionCapASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionCapASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionCapASTContext {
	var p = new(PrimaryExpressionCapASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionCapASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionCapASTContext) CapExpression() ICapExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICapExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICapExpressionContext)
}

func (s *PrimaryExpressionCapASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionCapAST(s)
	}
}

func (s *PrimaryExpressionCapASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionCapAST(s)
	}
}

func (s *PrimaryExpressionCapASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitPrimaryExpressionCapAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionSelectorASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionSelectorASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionSelectorASTContext {
	var p = new(PrimaryExpressionSelectorASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionSelectorASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionSelectorASTContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpressionSelectorASTContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *PrimaryExpressionSelectorASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionSelectorAST(s)
	}
}

func (s *PrimaryExpressionSelectorASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionSelectorAST(s)
	}
}

func (s *PrimaryExpressionSelectorASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitPrimaryExpressionSelectorAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	return p.primaryExpression(0)
}

func (p *miniGoParser) primaryExpression(_p int) (localctx IPrimaryExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, miniGoParserRULE_primaryExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniGoParserINTLITERAL, miniGoParserFLOATLITERAL, miniGoParserRUNELITERAL, miniGoParserRAWSTRINGLITERAL, miniGoParserINTERPRETEDSTRINGLITERAL, miniGoParserIDENTIFIER, miniGoParserLPAREN:
		localctx = NewPrimaryExpressionOperandASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(317)
			p.Operand()
		}

	case miniGoParserAPPEND:
		localctx = NewPrimaryExpressionAppendASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(318)
			p.AppendExpression()
		}

	case miniGoParserLEN:
		localctx = NewPrimaryExpressionLengthASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(319)
			p.LengthExpression()
		}

	case miniGoParserCAP:
		localctx = NewPrimaryExpressionCapASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(320)
			p.CapExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(329)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimaryExpressionSelectorASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_primaryExpression)
				p.SetState(323)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(324)
					p.Selector()
				}

			case 2:
				localctx = NewPrimaryExpressionIndexASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_primaryExpression)
				p.SetState(325)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(326)
					p.Index()
				}

			case 3:
				localctx = NewPrimaryExpressionArgumentsASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_primaryExpression)
				p.SetState(327)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(328)
					p.Arguments()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) CopyAll(ctx *OperandContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OperandLiteralASTContext struct {
	OperandContext
}

func NewOperandLiteralASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperandLiteralASTContext {
	var p = new(OperandLiteralASTContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *OperandLiteralASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandLiteralASTContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OperandLiteralASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterOperandLiteralAST(s)
	}
}

func (s *OperandLiteralASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitOperandLiteralAST(s)
	}
}

func (s *OperandLiteralASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitOperandLiteralAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type OperandIdentifierASTContext struct {
	OperandContext
}

func NewOperandIdentifierASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperandIdentifierASTContext {
	var p = new(OperandIdentifierASTContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *OperandIdentifierASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandIdentifierASTContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(miniGoParserIDENTIFIER, 0)
}

func (s *OperandIdentifierASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterOperandIdentifierAST(s)
	}
}

func (s *OperandIdentifierASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitOperandIdentifierAST(s)
	}
}

func (s *OperandIdentifierASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitOperandIdentifierAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type OperandParenASTContext struct {
	OperandContext
}

func NewOperandParenASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperandParenASTContext {
	var p = new(OperandParenASTContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *OperandParenASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandParenASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *OperandParenASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandParenASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *OperandParenASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterOperandParenAST(s)
	}
}

func (s *OperandParenASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitOperandParenAST(s)
	}
}

func (s *OperandParenASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitOperandParenAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, miniGoParserRULE_operand)
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniGoParserINTLITERAL, miniGoParserFLOATLITERAL, miniGoParserRUNELITERAL, miniGoParserRAWSTRINGLITERAL, miniGoParserINTERPRETEDSTRINGLITERAL:
		localctx = NewOperandLiteralASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(334)
			p.Literal()
		}

	case miniGoParserIDENTIFIER:
		localctx = NewOperandIdentifierASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(335)
			p.Match(miniGoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserLPAREN:
		localctx = NewOperandParenASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(336)
			p.Match(miniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.expression(0)
		}
		{
			p.SetState(338)
			p.Match(miniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LiteralIntASTContext struct {
	LiteralContext
}

func NewLiteralIntASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralIntASTContext {
	var p = new(LiteralIntASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LiteralIntASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralIntASTContext) INTLITERAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserINTLITERAL, 0)
}

func (s *LiteralIntASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterLiteralIntAST(s)
	}
}

func (s *LiteralIntASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitLiteralIntAST(s)
	}
}

func (s *LiteralIntASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitLiteralIntAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralInterpretedStringASTContext struct {
	LiteralContext
}

func NewLiteralInterpretedStringASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralInterpretedStringASTContext {
	var p = new(LiteralInterpretedStringASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LiteralInterpretedStringASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralInterpretedStringASTContext) INTERPRETEDSTRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserINTERPRETEDSTRINGLITERAL, 0)
}

func (s *LiteralInterpretedStringASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterLiteralInterpretedStringAST(s)
	}
}

func (s *LiteralInterpretedStringASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitLiteralInterpretedStringAST(s)
	}
}

func (s *LiteralInterpretedStringASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitLiteralInterpretedStringAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralFloatASTContext struct {
	LiteralContext
}

func NewLiteralFloatASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralFloatASTContext {
	var p = new(LiteralFloatASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LiteralFloatASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralFloatASTContext) FLOATLITERAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserFLOATLITERAL, 0)
}

func (s *LiteralFloatASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterLiteralFloatAST(s)
	}
}

func (s *LiteralFloatASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitLiteralFloatAST(s)
	}
}

func (s *LiteralFloatASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitLiteralFloatAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralRuneASTContext struct {
	LiteralContext
}

func NewLiteralRuneASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralRuneASTContext {
	var p = new(LiteralRuneASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LiteralRuneASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralRuneASTContext) RUNELITERAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserRUNELITERAL, 0)
}

func (s *LiteralRuneASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterLiteralRuneAST(s)
	}
}

func (s *LiteralRuneASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitLiteralRuneAST(s)
	}
}

func (s *LiteralRuneASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitLiteralRuneAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralRawStringASTContext struct {
	LiteralContext
}

func NewLiteralRawStringASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralRawStringASTContext {
	var p = new(LiteralRawStringASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LiteralRawStringASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralRawStringASTContext) RAWSTRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserRAWSTRINGLITERAL, 0)
}

func (s *LiteralRawStringASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterLiteralRawStringAST(s)
	}
}

func (s *LiteralRawStringASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitLiteralRawStringAST(s)
	}
}

func (s *LiteralRawStringASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitLiteralRawStringAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, miniGoParserRULE_literal)
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniGoParserINTLITERAL:
		localctx = NewLiteralIntASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(342)
			p.Match(miniGoParserINTLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserFLOATLITERAL:
		localctx = NewLiteralFloatASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(343)
			p.Match(miniGoParserFLOATLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserRUNELITERAL:
		localctx = NewLiteralRuneASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(344)
			p.Match(miniGoParserRUNELITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserRAWSTRINGLITERAL:
		localctx = NewLiteralRawStringASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(345)
			p.Match(miniGoParserRAWSTRINGLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserINTERPRETEDSTRINGLITERAL:
		localctx = NewLiteralInterpretedStringASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(346)
			p.Match(miniGoParserINTERPRETEDSTRINGLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) CopyAll(ctx *IndexContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IndexASTContext struct {
	IndexContext
}

func NewIndexASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexASTContext {
	var p = new(IndexASTContext)

	InitEmptyIndexContext(&p.IndexContext)
	p.parser = parser
	p.CopyAll(ctx.(*IndexContext))

	return p
}

func (s *IndexASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexASTContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(miniGoParserLBRACKET, 0)
}

func (s *IndexASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexASTContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(miniGoParserRBRACKET, 0)
}

func (s *IndexASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterIndexAST(s)
	}
}

func (s *IndexASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitIndexAST(s)
	}
}

func (s *IndexASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitIndexAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, miniGoParserRULE_index)
	localctx = NewIndexASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(miniGoParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(350)
		p.expression(0)
	}
	{
		p.SetState(351)
		p.Match(miniGoParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) CopyAll(ctx *ArgumentsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArgumentsASTContext struct {
	ArgumentsContext
}

func NewArgumentsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgumentsASTContext {
	var p = new(ArgumentsASTContext)

	InitEmptyArgumentsContext(&p.ArgumentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArgumentsContext))

	return p
}

func (s *ArgumentsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *ArgumentsASTContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterArgumentsAST(s)
	}
}

func (s *ArgumentsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitArgumentsAST(s)
	}
}

func (s *ArgumentsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitArgumentsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArgumentsEmptyASTContext struct {
	ArgumentsContext
}

func NewArgumentsEmptyASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgumentsEmptyASTContext {
	var p = new(ArgumentsEmptyASTContext)

	InitEmptyArgumentsContext(&p.ArgumentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArgumentsContext))

	return p
}

func (s *ArgumentsEmptyASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsEmptyASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *ArgumentsEmptyASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *ArgumentsEmptyASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterArgumentsEmptyAST(s)
	}
}

func (s *ArgumentsEmptyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitArgumentsEmptyAST(s)
	}
}

func (s *ArgumentsEmptyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitArgumentsEmptyAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, miniGoParserRULE_arguments)
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniGoParserLPAREN:
		localctx = NewArgumentsASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)
			p.Match(miniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.ExpressionList()
		}

	case miniGoParserRPAREN:
		localctx = NewArgumentsEmptyASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(355)
			p.Epsilon()
		}
		{
			p.SetState(356)
			p.Match(miniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) CopyAll(ctx *SelectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SelectorASTContext struct {
	SelectorContext
}

func NewSelectorASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorASTContext {
	var p = new(SelectorASTContext)

	InitEmptySelectorContext(&p.SelectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*SelectorContext))

	return p
}

func (s *SelectorASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorASTContext) DOT() antlr.TerminalNode {
	return s.GetToken(miniGoParserDOT, 0)
}

func (s *SelectorASTContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(miniGoParserIDENTIFIER, 0)
}

func (s *SelectorASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSelectorAST(s)
	}
}

func (s *SelectorASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSelectorAST(s)
	}
}

func (s *SelectorASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSelectorAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, miniGoParserRULE_selector)
	localctx = NewSelectorASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		p.Match(miniGoParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.Match(miniGoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAppendExpressionContext is an interface to support dynamic dispatch.
type IAppendExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAppendExpressionContext differentiates from other interfaces.
	IsAppendExpressionContext()
}

type AppendExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAppendExpressionContext() *AppendExpressionContext {
	var p = new(AppendExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_appendExpression
	return p
}

func InitEmptyAppendExpressionContext(p *AppendExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_appendExpression
}

func (*AppendExpressionContext) IsAppendExpressionContext() {}

func NewAppendExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppendExpressionContext {
	var p = new(AppendExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_appendExpression

	return p
}

func (s *AppendExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AppendExpressionContext) CopyAll(ctx *AppendExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AppendExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AppendExpressionASTContext struct {
	AppendExpressionContext
}

func NewAppendExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendExpressionASTContext {
	var p = new(AppendExpressionASTContext)

	InitEmptyAppendExpressionContext(&p.AppendExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AppendExpressionContext))

	return p
}

func (s *AppendExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpressionASTContext) APPEND() antlr.TerminalNode {
	return s.GetToken(miniGoParserAPPEND, 0)
}

func (s *AppendExpressionASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *AppendExpressionASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AppendExpressionASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AppendExpressionASTContext) COMMA() antlr.TerminalNode {
	return s.GetToken(miniGoParserCOMMA, 0)
}

func (s *AppendExpressionASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *AppendExpressionASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAppendExpressionAST(s)
	}
}

func (s *AppendExpressionASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAppendExpressionAST(s)
	}
}

func (s *AppendExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAppendExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) AppendExpression() (localctx IAppendExpressionContext) {
	localctx = NewAppendExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, miniGoParserRULE_appendExpression)
	localctx = NewAppendExpressionASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(miniGoParserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.Match(miniGoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(365)
		p.expression(0)
	}
	{
		p.SetState(366)
		p.Match(miniGoParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.expression(0)
	}
	{
		p.SetState(368)
		p.Match(miniGoParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILengthExpressionContext is an interface to support dynamic dispatch.
type ILengthExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLengthExpressionContext differentiates from other interfaces.
	IsLengthExpressionContext()
}

type LengthExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthExpressionContext() *LengthExpressionContext {
	var p = new(LengthExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_lengthExpression
	return p
}

func InitEmptyLengthExpressionContext(p *LengthExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_lengthExpression
}

func (*LengthExpressionContext) IsLengthExpressionContext() {}

func NewLengthExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthExpressionContext {
	var p = new(LengthExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_lengthExpression

	return p
}

func (s *LengthExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthExpressionContext) CopyAll(ctx *LengthExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LengthExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LengthExpressionASTContext struct {
	LengthExpressionContext
}

func NewLengthExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LengthExpressionASTContext {
	var p = new(LengthExpressionASTContext)

	InitEmptyLengthExpressionContext(&p.LengthExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*LengthExpressionContext))

	return p
}

func (s *LengthExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthExpressionASTContext) LEN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLEN, 0)
}

func (s *LengthExpressionASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *LengthExpressionASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LengthExpressionASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *LengthExpressionASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterLengthExpressionAST(s)
	}
}

func (s *LengthExpressionASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitLengthExpressionAST(s)
	}
}

func (s *LengthExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitLengthExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) LengthExpression() (localctx ILengthExpressionContext) {
	localctx = NewLengthExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, miniGoParserRULE_lengthExpression)
	localctx = NewLengthExpressionASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(miniGoParserLEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Match(miniGoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(372)
		p.expression(0)
	}
	{
		p.SetState(373)
		p.Match(miniGoParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICapExpressionContext is an interface to support dynamic dispatch.
type ICapExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCapExpressionContext differentiates from other interfaces.
	IsCapExpressionContext()
}

type CapExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCapExpressionContext() *CapExpressionContext {
	var p = new(CapExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_capExpression
	return p
}

func InitEmptyCapExpressionContext(p *CapExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_capExpression
}

func (*CapExpressionContext) IsCapExpressionContext() {}

func NewCapExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CapExpressionContext {
	var p = new(CapExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_capExpression

	return p
}

func (s *CapExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CapExpressionContext) CopyAll(ctx *CapExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CapExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CapExpressionASTContext struct {
	CapExpressionContext
}

func NewCapExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CapExpressionASTContext {
	var p = new(CapExpressionASTContext)

	InitEmptyCapExpressionContext(&p.CapExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*CapExpressionContext))

	return p
}

func (s *CapExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpressionASTContext) CAP() antlr.TerminalNode {
	return s.GetToken(miniGoParserCAP, 0)
}

func (s *CapExpressionASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *CapExpressionASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CapExpressionASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *CapExpressionASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterCapExpressionAST(s)
	}
}

func (s *CapExpressionASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitCapExpressionAST(s)
	}
}

func (s *CapExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitCapExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) CapExpression() (localctx ICapExpressionContext) {
	localctx = NewCapExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, miniGoParserRULE_capExpression)
	localctx = NewCapExpressionASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		p.Match(miniGoParserCAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(376)
		p.Match(miniGoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.expression(0)
	}
	{
		p.SetState(378)
		p.Match(miniGoParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_statementList
	return p
}

func InitEmptyStatementListContext(p *StatementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_statementList
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) CopyAll(ctx *StatementListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatementListASTContext struct {
	StatementListContext
}

func NewStatementListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementListASTContext {
	var p = new(StatementListASTContext)

	InitEmptyStatementListContext(&p.StatementListContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementListContext))

	return p
}

func (s *StatementListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListASTContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementListASTContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementListAST(s)
	}
}

func (s *StatementListASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementListAST(s)
	}
}

func (s *StatementListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, miniGoParserRULE_statementList)
	var _alt int

	localctx = NewStatementListASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(380)
				p.Statement()
			}

		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) CopyAll(ctx *BlockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlockASTContext struct {
	BlockContext
}

func NewBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockASTContext {
	var p = new(BlockASTContext)

	InitEmptyBlockContext(&p.BlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*BlockContext))

	return p
}

func (s *BlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockASTContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(miniGoParserLBRACE, 0)
}

func (s *BlockASTContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockASTContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(miniGoParserRBRACE, 0)
}

func (s *BlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterBlockAST(s)
	}
}

func (s *BlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitBlockAST(s)
	}
}

func (s *BlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, miniGoParserRULE_block)
	localctx = NewBlockASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.Match(miniGoParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.StatementList()
	}
	{
		p.SetState(388)
		p.Match(miniGoParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatementSimpleASTContext struct {
	StatementContext
}

func NewStatementSimpleASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementSimpleASTContext {
	var p = new(StatementSimpleASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementSimpleASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementSimpleASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *StatementSimpleASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *StatementSimpleASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementSimpleAST(s)
	}
}

func (s *StatementSimpleASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementSimpleAST(s)
	}
}

func (s *StatementSimpleASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementSimpleAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementSwitchASTContext struct {
	StatementContext
}

func NewStatementSwitchASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementSwitchASTContext {
	var p = new(StatementSwitchASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementSwitchASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementSwitchASTContext) SwitchStmt() ISwitchStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStmtContext)
}

func (s *StatementSwitchASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *StatementSwitchASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementSwitchAST(s)
	}
}

func (s *StatementSwitchASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementSwitchAST(s)
	}
}

func (s *StatementSwitchASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementSwitchAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementLoopASTContext struct {
	StatementContext
}

func NewStatementLoopASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementLoopASTContext {
	var p = new(StatementLoopASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementLoopASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementLoopASTContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementLoopASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *StatementLoopASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementLoopAST(s)
	}
}

func (s *StatementLoopASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementLoopAST(s)
	}
}

func (s *StatementLoopASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementLoopAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementBreakASTContext struct {
	StatementContext
}

func NewStatementBreakASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementBreakASTContext {
	var p = new(StatementBreakASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementBreakASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementBreakASTContext) BREAK() antlr.TerminalNode {
	return s.GetToken(miniGoParserBREAK, 0)
}

func (s *StatementBreakASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *StatementBreakASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementBreakAST(s)
	}
}

func (s *StatementBreakASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementBreakAST(s)
	}
}

func (s *StatementBreakASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementBreakAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementBlockASTContext struct {
	StatementContext
}

func NewStatementBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementBlockASTContext {
	var p = new(StatementBlockASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementBlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementBlockASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementBlockASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *StatementBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementBlockAST(s)
	}
}

func (s *StatementBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementBlockAST(s)
	}
}

func (s *StatementBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementTypeDeclASTContext struct {
	StatementContext
}

func NewStatementTypeDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementTypeDeclASTContext {
	var p = new(StatementTypeDeclASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementTypeDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementTypeDeclASTContext) TypeDecl() ITypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *StatementTypeDeclASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementTypeDeclAST(s)
	}
}

func (s *StatementTypeDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementTypeDeclAST(s)
	}
}

func (s *StatementTypeDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementTypeDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementReturnASTContext struct {
	StatementContext
}

func NewStatementReturnASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementReturnASTContext {
	var p = new(StatementReturnASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementReturnASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementReturnASTContext) RETURN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRETURN, 0)
}

func (s *StatementReturnASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *StatementReturnASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementReturnASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *StatementReturnASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementReturnAST(s)
	}
}

func (s *StatementReturnASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementReturnAST(s)
	}
}

func (s *StatementReturnASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementReturnAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementPrintASTContext struct {
	StatementContext
}

func NewStatementPrintASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementPrintASTContext {
	var p = new(StatementPrintASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementPrintASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementPrintASTContext) PRINT() antlr.TerminalNode {
	return s.GetToken(miniGoParserPRINT, 0)
}

func (s *StatementPrintASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *StatementPrintASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *StatementPrintASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *StatementPrintASTContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *StatementPrintASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *StatementPrintASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementPrintAST(s)
	}
}

func (s *StatementPrintASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementPrintAST(s)
	}
}

func (s *StatementPrintASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementPrintAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementVariableDeclASTContext struct {
	StatementContext
}

func NewStatementVariableDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementVariableDeclASTContext {
	var p = new(StatementVariableDeclASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementVariableDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementVariableDeclASTContext) VariableDecl() IVariableDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *StatementVariableDeclASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementVariableDeclAST(s)
	}
}

func (s *StatementVariableDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementVariableDeclAST(s)
	}
}

func (s *StatementVariableDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementVariableDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementPrintlnASTContext struct {
	StatementContext
}

func NewStatementPrintlnASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementPrintlnASTContext {
	var p = new(StatementPrintlnASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementPrintlnASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementPrintlnASTContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(miniGoParserPRINTLN, 0)
}

func (s *StatementPrintlnASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserLPAREN, 0)
}

func (s *StatementPrintlnASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(miniGoParserRPAREN, 0)
}

func (s *StatementPrintlnASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *StatementPrintlnASTContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *StatementPrintlnASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *StatementPrintlnASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementPrintlnAST(s)
	}
}

func (s *StatementPrintlnASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementPrintlnAST(s)
	}
}

func (s *StatementPrintlnASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementPrintlnAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementContinueASTContext struct {
	StatementContext
}

func NewStatementContinueASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementContinueASTContext {
	var p = new(StatementContinueASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementContinueASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContinueASTContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(miniGoParserCONTINUE, 0)
}

func (s *StatementContinueASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *StatementContinueASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementContinueAST(s)
	}
}

func (s *StatementContinueASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementContinueAST(s)
	}
}

func (s *StatementContinueASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementContinueAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementIfASTContext struct {
	StatementContext
}

func NewStatementIfASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementIfASTContext {
	var p = new(StatementIfASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StatementIfASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementIfASTContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementIfASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *StatementIfASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterStatementIfAST(s)
	}
}

func (s *StatementIfASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitStatementIfAST(s)
	}
}

func (s *StatementIfASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitStatementIfAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, miniGoParserRULE_statement)
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniGoParserPRINT:
		localctx = NewStatementPrintASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(390)
			p.Match(miniGoParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.Match(miniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(394)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case miniGoParserINTLITERAL, miniGoParserFLOATLITERAL, miniGoParserRUNELITERAL, miniGoParserRAWSTRINGLITERAL, miniGoParserINTERPRETEDSTRINGLITERAL, miniGoParserLEN, miniGoParserCAP, miniGoParserAPPEND, miniGoParserIDENTIFIER, miniGoParserPLUS, miniGoParserMINUS, miniGoParserNOT, miniGoParserBITWISEXOR, miniGoParserLPAREN:
			{
				p.SetState(392)
				p.ExpressionList()
			}

		case miniGoParserRPAREN:
			{
				p.SetState(393)
				p.Epsilon()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(396)
			p.Match(miniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserPRINTLN:
		localctx = NewStatementPrintlnASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(399)
			p.Match(miniGoParserPRINTLN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Match(miniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(403)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case miniGoParserINTLITERAL, miniGoParserFLOATLITERAL, miniGoParserRUNELITERAL, miniGoParserRAWSTRINGLITERAL, miniGoParserINTERPRETEDSTRINGLITERAL, miniGoParserLEN, miniGoParserCAP, miniGoParserAPPEND, miniGoParserIDENTIFIER, miniGoParserPLUS, miniGoParserMINUS, miniGoParserNOT, miniGoParserBITWISEXOR, miniGoParserLPAREN:
			{
				p.SetState(401)
				p.ExpressionList()
			}

		case miniGoParserRPAREN:
			{
				p.SetState(402)
				p.Epsilon()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(405)
			p.Match(miniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserRETURN:
		localctx = NewStatementReturnASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(408)
			p.Match(miniGoParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(411)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case miniGoParserINTLITERAL, miniGoParserFLOATLITERAL, miniGoParserRUNELITERAL, miniGoParserRAWSTRINGLITERAL, miniGoParserINTERPRETEDSTRINGLITERAL, miniGoParserLEN, miniGoParserCAP, miniGoParserAPPEND, miniGoParserIDENTIFIER, miniGoParserPLUS, miniGoParserMINUS, miniGoParserNOT, miniGoParserBITWISEXOR, miniGoParserLPAREN:
			{
				p.SetState(409)
				p.expression(0)
			}

		case miniGoParserSEMICOLON:
			{
				p.SetState(410)
				p.Epsilon()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(413)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserBREAK:
		localctx = NewStatementBreakASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(415)
			p.Match(miniGoParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserCONTINUE:
		localctx = NewStatementContinueASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(417)
			p.Match(miniGoParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserINTLITERAL, miniGoParserFLOATLITERAL, miniGoParserRUNELITERAL, miniGoParserRAWSTRINGLITERAL, miniGoParserINTERPRETEDSTRINGLITERAL, miniGoParserLEN, miniGoParserCAP, miniGoParserAPPEND, miniGoParserIDENTIFIER, miniGoParserPLUS, miniGoParserMINUS, miniGoParserNOT, miniGoParserBITWISEXOR, miniGoParserSEMICOLON, miniGoParserLPAREN:
		localctx = NewStatementSimpleASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(419)
			p.SimpleStatement()
		}
		{
			p.SetState(420)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserLBRACE:
		localctx = NewStatementBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(422)
			p.Block()
		}
		{
			p.SetState(423)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserSWITCH:
		localctx = NewStatementSwitchASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(425)
			p.SwitchStmt()
		}
		{
			p.SetState(426)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserIF:
		localctx = NewStatementIfASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(428)
			p.IfStatement()
		}
		{
			p.SetState(429)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserFOR:
		localctx = NewStatementLoopASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(431)
			p.Loop()
		}
		{
			p.SetState(432)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniGoParserTYPE:
		localctx = NewStatementTypeDeclASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(434)
			p.TypeDecl()
		}

	case miniGoParserVAR:
		localctx = NewStatementVariableDeclASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(435)
			p.VariableDecl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleStatementContext is an interface to support dynamic dispatch.
type ISimpleStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSimpleStatementContext differentiates from other interfaces.
	IsSimpleStatementContext()
}

type SimpleStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStatementContext() *SimpleStatementContext {
	var p = new(SimpleStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_simpleStatement
	return p
}

func InitEmptySimpleStatementContext(p *SimpleStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_simpleStatement
}

func (*SimpleStatementContext) IsSimpleStatementContext() {}

func NewSimpleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStatementContext {
	var p = new(SimpleStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_simpleStatement

	return p
}

func (s *SimpleStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStatementContext) CopyAll(ctx *SimpleStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SimpleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SimpleStatementExpressionASTContext struct {
	SimpleStatementContext
}

func NewSimpleStatementExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleStatementExpressionASTContext {
	var p = new(SimpleStatementExpressionASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *SimpleStatementExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementExpressionASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SimpleStatementExpressionASTContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(miniGoParserINCREMENT, 0)
}

func (s *SimpleStatementExpressionASTContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(miniGoParserDECREMENT, 0)
}

func (s *SimpleStatementExpressionASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *SimpleStatementExpressionASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSimpleStatementExpressionAST(s)
	}
}

func (s *SimpleStatementExpressionASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSimpleStatementExpressionAST(s)
	}
}

func (s *SimpleStatementExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSimpleStatementExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleStatementAssignmentASTContext struct {
	SimpleStatementContext
}

func NewSimpleStatementAssignmentASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleStatementAssignmentASTContext {
	var p = new(SimpleStatementAssignmentASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *SimpleStatementAssignmentASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementAssignmentASTContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *SimpleStatementAssignmentASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSimpleStatementAssignmentAST(s)
	}
}

func (s *SimpleStatementAssignmentASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSimpleStatementAssignmentAST(s)
	}
}

func (s *SimpleStatementAssignmentASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSimpleStatementAssignmentAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleStatementExpressionListAssignASTContext struct {
	SimpleStatementContext
}

func NewSimpleStatementExpressionListAssignASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleStatementExpressionListAssignASTContext {
	var p = new(SimpleStatementExpressionListAssignASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *SimpleStatementExpressionListAssignASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementExpressionListAssignASTContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *SimpleStatementExpressionListAssignASTContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *SimpleStatementExpressionListAssignASTContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(miniGoParserASSIGN, 0)
}

func (s *SimpleStatementExpressionListAssignASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSimpleStatementExpressionListAssignAST(s)
	}
}

func (s *SimpleStatementExpressionListAssignASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSimpleStatementExpressionListAssignAST(s)
	}
}

func (s *SimpleStatementExpressionListAssignASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSimpleStatementExpressionListAssignAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleStatementEmptyASTContext struct {
	SimpleStatementContext
}

func NewSimpleStatementEmptyASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleStatementEmptyASTContext {
	var p = new(SimpleStatementEmptyASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *SimpleStatementEmptyASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementEmptyASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *SimpleStatementEmptyASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSimpleStatementEmptyAST(s)
	}
}

func (s *SimpleStatementEmptyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSimpleStatementEmptyAST(s)
	}
}

func (s *SimpleStatementEmptyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSimpleStatementEmptyAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) SimpleStatement() (localctx ISimpleStatementContext) {
	localctx = NewSimpleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, miniGoParserRULE_simpleStatement)
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpleStatementEmptyASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(438)
			p.Epsilon()
		}

	case 2:
		localctx = NewSimpleStatementExpressionASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(439)
			p.expression(0)
		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case miniGoParserINCREMENT:
			{
				p.SetState(440)
				p.Match(miniGoParserINCREMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case miniGoParserDECREMENT:
			{
				p.SetState(441)
				p.Match(miniGoParserDECREMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case miniGoParserSEMICOLON, miniGoParserLBRACE:
			{
				p.SetState(442)
				p.Epsilon()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	case 3:
		localctx = NewSimpleStatementAssignmentASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(445)
			p.AssignmentStatement()
		}

	case 4:
		localctx = NewSimpleStatementExpressionListAssignASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(446)
			p.ExpressionList()
		}
		{
			p.SetState(447)
			p.Match(miniGoParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)
			p.ExpressionList()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) CopyAll(ctx *AssignmentStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignmentStatementBitwiseAndEqualASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementBitwiseAndEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementBitwiseAndEqualASTContext {
	var p = new(AssignmentStatementBitwiseAndEqualASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementBitwiseAndEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementBitwiseAndEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementBitwiseAndEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementBitwiseAndEqualASTContext) BITWISEANDEQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserBITWISEANDEQUAL, 0)
}

func (s *AssignmentStatementBitwiseAndEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAssignmentStatementBitwiseAndEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseAndEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAssignmentStatementBitwiseAndEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseAndEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAssignmentStatementBitwiseAndEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementBitwiseXorEqualASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementBitwiseXorEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementBitwiseXorEqualASTContext {
	var p = new(AssignmentStatementBitwiseXorEqualASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementBitwiseXorEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementBitwiseXorEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementBitwiseXorEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementBitwiseXorEqualASTContext) BITWISEXOREQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserBITWISEXOREQUAL, 0)
}

func (s *AssignmentStatementBitwiseXorEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAssignmentStatementBitwiseXorEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseXorEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAssignmentStatementBitwiseXorEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseXorEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAssignmentStatementBitwiseXorEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementShiftLeftEqualASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementShiftLeftEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementShiftLeftEqualASTContext {
	var p = new(AssignmentStatementShiftLeftEqualASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementShiftLeftEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementShiftLeftEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementShiftLeftEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementShiftLeftEqualASTContext) SHIFTLEFTEQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserSHIFTLEFTEQUAL, 0)
}

func (s *AssignmentStatementShiftLeftEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAssignmentStatementShiftLeftEqualAST(s)
	}
}

func (s *AssignmentStatementShiftLeftEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAssignmentStatementShiftLeftEqualAST(s)
	}
}

func (s *AssignmentStatementShiftLeftEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAssignmentStatementShiftLeftEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementPlusEqualASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementPlusEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementPlusEqualASTContext {
	var p = new(AssignmentStatementPlusEqualASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementPlusEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementPlusEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementPlusEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementPlusEqualASTContext) PLUSEQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserPLUSEQUAL, 0)
}

func (s *AssignmentStatementPlusEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAssignmentStatementPlusEqualAST(s)
	}
}

func (s *AssignmentStatementPlusEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAssignmentStatementPlusEqualAST(s)
	}
}

func (s *AssignmentStatementPlusEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAssignmentStatementPlusEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementMinusEqualASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementMinusEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementMinusEqualASTContext {
	var p = new(AssignmentStatementMinusEqualASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementMinusEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementMinusEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementMinusEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementMinusEqualASTContext) MINUSEQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserMINUSEQUAL, 0)
}

func (s *AssignmentStatementMinusEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAssignmentStatementMinusEqualAST(s)
	}
}

func (s *AssignmentStatementMinusEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAssignmentStatementMinusEqualAST(s)
	}
}

func (s *AssignmentStatementMinusEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAssignmentStatementMinusEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementModuloEqualASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementModuloEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementModuloEqualASTContext {
	var p = new(AssignmentStatementModuloEqualASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementModuloEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementModuloEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementModuloEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementModuloEqualASTContext) MODULOEQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserMODULOEQUAL, 0)
}

func (s *AssignmentStatementModuloEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAssignmentStatementModuloEqualAST(s)
	}
}

func (s *AssignmentStatementModuloEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAssignmentStatementModuloEqualAST(s)
	}
}

func (s *AssignmentStatementModuloEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAssignmentStatementModuloEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementASTContext {
	var p = new(AssignmentStatementASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementASTContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementASTContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *AssignmentStatementASTContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(miniGoParserASSIGN, 0)
}

func (s *AssignmentStatementASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAssignmentStatementAST(s)
	}
}

func (s *AssignmentStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAssignmentStatementAST(s)
	}
}

func (s *AssignmentStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAssignmentStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementBitwiseClearEqualASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementBitwiseClearEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementBitwiseClearEqualASTContext {
	var p = new(AssignmentStatementBitwiseClearEqualASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementBitwiseClearEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementBitwiseClearEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementBitwiseClearEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementBitwiseClearEqualASTContext) BITWISECLEAREQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserBITWISECLEAREQUAL, 0)
}

func (s *AssignmentStatementBitwiseClearEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAssignmentStatementBitwiseClearEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseClearEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAssignmentStatementBitwiseClearEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseClearEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAssignmentStatementBitwiseClearEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementDivideEqualASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementDivideEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementDivideEqualASTContext {
	var p = new(AssignmentStatementDivideEqualASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementDivideEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementDivideEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementDivideEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementDivideEqualASTContext) DIVIDEEQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserDIVIDEEQUAL, 0)
}

func (s *AssignmentStatementDivideEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAssignmentStatementDivideEqualAST(s)
	}
}

func (s *AssignmentStatementDivideEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAssignmentStatementDivideEqualAST(s)
	}
}

func (s *AssignmentStatementDivideEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAssignmentStatementDivideEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementShiftRightEqualASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementShiftRightEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementShiftRightEqualASTContext {
	var p = new(AssignmentStatementShiftRightEqualASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementShiftRightEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementShiftRightEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementShiftRightEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementShiftRightEqualASTContext) SHIFTRIGHTEQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserSHIFTRIGHTEQUAL, 0)
}

func (s *AssignmentStatementShiftRightEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAssignmentStatementShiftRightEqualAST(s)
	}
}

func (s *AssignmentStatementShiftRightEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAssignmentStatementShiftRightEqualAST(s)
	}
}

func (s *AssignmentStatementShiftRightEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAssignmentStatementShiftRightEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementMultiplyEqualASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementMultiplyEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementMultiplyEqualASTContext {
	var p = new(AssignmentStatementMultiplyEqualASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementMultiplyEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementMultiplyEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementMultiplyEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementMultiplyEqualASTContext) MULTIPLYEQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserMULTIPLYEQUAL, 0)
}

func (s *AssignmentStatementMultiplyEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAssignmentStatementMultiplyEqualAST(s)
	}
}

func (s *AssignmentStatementMultiplyEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAssignmentStatementMultiplyEqualAST(s)
	}
}

func (s *AssignmentStatementMultiplyEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAssignmentStatementMultiplyEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementBitwiseOrEqualASTContext struct {
	AssignmentStatementContext
}

func NewAssignmentStatementBitwiseOrEqualASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementBitwiseOrEqualASTContext {
	var p = new(AssignmentStatementBitwiseOrEqualASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignmentStatementBitwiseOrEqualASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementBitwiseOrEqualASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementBitwiseOrEqualASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementBitwiseOrEqualASTContext) BITWISEOREQUAL() antlr.TerminalNode {
	return s.GetToken(miniGoParserBITWISEOREQUAL, 0)
}

func (s *AssignmentStatementBitwiseOrEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterAssignmentStatementBitwiseOrEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseOrEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitAssignmentStatementBitwiseOrEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseOrEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitAssignmentStatementBitwiseOrEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, miniGoParserRULE_assignmentStatement)
	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignmentStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			p.ExpressionList()
		}
		{
			p.SetState(453)
			p.Match(miniGoParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)
			p.ExpressionList()
		}

	case 2:
		localctx = NewAssignmentStatementPlusEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(456)
			p.expression(0)
		}
		{
			p.SetState(457)
			p.Match(miniGoParserPLUSEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)
			p.expression(0)
		}

	case 3:
		localctx = NewAssignmentStatementMinusEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(460)
			p.expression(0)
		}
		{
			p.SetState(461)
			p.Match(miniGoParserMINUSEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.expression(0)
		}

	case 4:
		localctx = NewAssignmentStatementBitwiseAndEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(464)
			p.expression(0)
		}
		{
			p.SetState(465)
			p.Match(miniGoParserBITWISEANDEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(466)
			p.expression(0)
		}

	case 5:
		localctx = NewAssignmentStatementBitwiseOrEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(468)
			p.expression(0)
		}
		{
			p.SetState(469)
			p.Match(miniGoParserBITWISEOREQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
			p.expression(0)
		}

	case 6:
		localctx = NewAssignmentStatementMultiplyEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(472)
			p.expression(0)
		}
		{
			p.SetState(473)
			p.Match(miniGoParserMULTIPLYEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(474)
			p.expression(0)
		}

	case 7:
		localctx = NewAssignmentStatementBitwiseXorEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(476)
			p.expression(0)
		}
		{
			p.SetState(477)
			p.Match(miniGoParserBITWISEXOREQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(478)
			p.expression(0)
		}

	case 8:
		localctx = NewAssignmentStatementShiftLeftEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(480)
			p.expression(0)
		}
		{
			p.SetState(481)
			p.Match(miniGoParserSHIFTLEFTEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(482)
			p.expression(0)
		}

	case 9:
		localctx = NewAssignmentStatementShiftRightEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(484)
			p.expression(0)
		}
		{
			p.SetState(485)
			p.Match(miniGoParserSHIFTRIGHTEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(486)
			p.expression(0)
		}

	case 10:
		localctx = NewAssignmentStatementBitwiseClearEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(488)
			p.expression(0)
		}
		{
			p.SetState(489)
			p.Match(miniGoParserBITWISECLEAREQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.expression(0)
		}

	case 11:
		localctx = NewAssignmentStatementModuloEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(492)
			p.expression(0)
		}
		{
			p.SetState(493)
			p.Match(miniGoParserMODULOEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)
			p.expression(0)
		}

	case 12:
		localctx = NewAssignmentStatementDivideEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(496)
			p.expression(0)
		}
		{
			p.SetState(497)
			p.Match(miniGoParserDIVIDEEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(498)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) CopyAll(ctx *IfStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfSimpleStatementASTContext struct {
	IfStatementContext
}

func NewIfSimpleStatementASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfSimpleStatementASTContext {
	var p = new(IfSimpleStatementASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfSimpleStatementASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfSimpleStatementASTContext) IF() antlr.TerminalNode {
	return s.GetToken(miniGoParserIF, 0)
}

func (s *IfSimpleStatementASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IfSimpleStatementASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *IfSimpleStatementASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfSimpleStatementASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfSimpleStatementASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterIfSimpleStatementAST(s)
	}
}

func (s *IfSimpleStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitIfSimpleStatementAST(s)
	}
}

func (s *IfSimpleStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitIfSimpleStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseIfStatementASTContext struct {
	IfStatementContext
}

func NewIfElseIfStatementASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseIfStatementASTContext {
	var p = new(IfElseIfStatementASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfElseIfStatementASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseIfStatementASTContext) IF() antlr.TerminalNode {
	return s.GetToken(miniGoParserIF, 0)
}

func (s *IfElseIfStatementASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfElseIfStatementASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfElseIfStatementASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(miniGoParserELSE, 0)
}

func (s *IfElseIfStatementASTContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfElseIfStatementASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterIfElseIfStatementAST(s)
	}
}

func (s *IfElseIfStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitIfElseIfStatementAST(s)
	}
}

func (s *IfElseIfStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitIfElseIfStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStatementASTContext struct {
	IfStatementContext
}

func NewIfStatementASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementASTContext {
	var p = new(IfStatementASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfStatementASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementASTContext) IF() antlr.TerminalNode {
	return s.GetToken(miniGoParserIF, 0)
}

func (s *IfStatementASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterIfStatementAST(s)
	}
}

func (s *IfStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitIfStatementAST(s)
	}
}

func (s *IfStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitIfStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseStatementASTContext struct {
	IfStatementContext
}

func NewIfElseStatementASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseStatementASTContext {
	var p = new(IfElseStatementASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfElseStatementASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseStatementASTContext) IF() antlr.TerminalNode {
	return s.GetToken(miniGoParserIF, 0)
}

func (s *IfElseStatementASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfElseStatementASTContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfElseStatementASTContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfElseStatementASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(miniGoParserELSE, 0)
}

func (s *IfElseStatementASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterIfElseStatementAST(s)
	}
}

func (s *IfElseStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitIfElseStatementAST(s)
	}
}

func (s *IfElseStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitIfElseStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfSimpleElseStatementASTContext struct {
	IfStatementContext
}

func NewIfSimpleElseStatementASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfSimpleElseStatementASTContext {
	var p = new(IfSimpleElseStatementASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfSimpleElseStatementASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfSimpleElseStatementASTContext) IF() antlr.TerminalNode {
	return s.GetToken(miniGoParserIF, 0)
}

func (s *IfSimpleElseStatementASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IfSimpleElseStatementASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *IfSimpleElseStatementASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfSimpleElseStatementASTContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfSimpleElseStatementASTContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfSimpleElseStatementASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(miniGoParserELSE, 0)
}

func (s *IfSimpleElseStatementASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterIfSimpleElseStatementAST(s)
	}
}

func (s *IfSimpleElseStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitIfSimpleElseStatementAST(s)
	}
}

func (s *IfSimpleElseStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitIfSimpleElseStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfSimpleElseIfStatementASTContext struct {
	IfStatementContext
}

func NewIfSimpleElseIfStatementASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfSimpleElseIfStatementASTContext {
	var p = new(IfSimpleElseIfStatementASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfSimpleElseIfStatementASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfSimpleElseIfStatementASTContext) IF() antlr.TerminalNode {
	return s.GetToken(miniGoParserIF, 0)
}

func (s *IfSimpleElseIfStatementASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IfSimpleElseIfStatementASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *IfSimpleElseIfStatementASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfSimpleElseIfStatementASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfSimpleElseIfStatementASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(miniGoParserELSE, 0)
}

func (s *IfSimpleElseIfStatementASTContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfSimpleElseIfStatementASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterIfSimpleElseIfStatementAST(s)
	}
}

func (s *IfSimpleElseIfStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitIfSimpleElseIfStatementAST(s)
	}
}

func (s *IfSimpleElseIfStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitIfSimpleElseIfStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, miniGoParserRULE_ifStatement)
	p.SetState(540)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(502)
			p.Match(miniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(503)
			p.expression(0)
		}
		{
			p.SetState(504)
			p.Block()
		}

	case 2:
		localctx = NewIfElseIfStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(506)
			p.Match(miniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(507)
			p.expression(0)
		}
		{
			p.SetState(508)
			p.Block()
		}
		{
			p.SetState(509)
			p.Match(miniGoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(510)
			p.IfStatement()
		}

	case 3:
		localctx = NewIfElseStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(512)
			p.Match(miniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(513)
			p.expression(0)
		}
		{
			p.SetState(514)
			p.Block()
		}
		{
			p.SetState(515)
			p.Match(miniGoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(516)
			p.Block()
		}

	case 4:
		localctx = NewIfSimpleStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(518)
			p.Match(miniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(519)
			p.SimpleStatement()
		}
		{
			p.SetState(520)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(521)
			p.expression(0)
		}
		{
			p.SetState(522)
			p.Block()
		}

	case 5:
		localctx = NewIfSimpleElseIfStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(524)
			p.Match(miniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(525)
			p.SimpleStatement()
		}
		{
			p.SetState(526)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(527)
			p.expression(0)
		}
		{
			p.SetState(528)
			p.Block()
		}
		{
			p.SetState(529)
			p.Match(miniGoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(530)
			p.IfStatement()
		}

	case 6:
		localctx = NewIfSimpleElseStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(532)
			p.Match(miniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(533)
			p.SimpleStatement()
		}
		{
			p.SetState(534)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(535)
			p.expression(0)
		}
		{
			p.SetState(536)
			p.Block()
		}
		{
			p.SetState(537)
			p.Match(miniGoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(538)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_loop
	return p
}

func InitEmptyLoopContext(p *LoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_loop
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) CopyAll(ctx *LoopContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LoopExpressionBlockASTContext struct {
	LoopContext
}

func NewLoopExpressionBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LoopExpressionBlockASTContext {
	var p = new(LoopExpressionBlockASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *LoopExpressionBlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopExpressionBlockASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(miniGoParserFOR, 0)
}

func (s *LoopExpressionBlockASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopExpressionBlockASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopExpressionBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterLoopExpressionBlockAST(s)
	}
}

func (s *LoopExpressionBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitLoopExpressionBlockAST(s)
	}
}

func (s *LoopExpressionBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitLoopExpressionBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LoopBlockASTContext struct {
	LoopContext
}

func NewLoopBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LoopBlockASTContext {
	var p = new(LoopBlockASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *LoopBlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopBlockASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(miniGoParserFOR, 0)
}

func (s *LoopBlockASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterLoopBlockAST(s)
	}
}

func (s *LoopBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitLoopBlockAST(s)
	}
}

func (s *LoopBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitLoopBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LoopSimpleStatementExpressionSimpleStatementBlockASTContext struct {
	LoopContext
}

func NewLoopSimpleStatementExpressionSimpleStatementBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LoopSimpleStatementExpressionSimpleStatementBlockASTContext {
	var p = new(LoopSimpleStatementExpressionSimpleStatementBlockASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(miniGoParserFOR, 0)
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) AllSimpleStatement() []ISimpleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStatementContext); ok {
			tst[i] = t.(ISimpleStatementContext)
			i++
		}
	}

	return tst
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) SimpleStatement(i int) ISimpleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(miniGoParserSEMICOLON)
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, i)
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterLoopSimpleStatementExpressionSimpleStatementBlockAST(s)
	}
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitLoopSimpleStatementExpressionSimpleStatementBlockAST(s)
	}
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitLoopSimpleStatementExpressionSimpleStatementBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LoopSimpleStatementSimpleStatementBlockASTContext struct {
	LoopContext
}

func NewLoopSimpleStatementSimpleStatementBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LoopSimpleStatementSimpleStatementBlockASTContext {
	var p = new(LoopSimpleStatementSimpleStatementBlockASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(miniGoParserFOR, 0)
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) AllSimpleStatement() []ISimpleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStatementContext); ok {
			tst[i] = t.(ISimpleStatementContext)
			i++
		}
	}

	return tst
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) SimpleStatement(i int) ISimpleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(miniGoParserSEMICOLON)
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, i)
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterLoopSimpleStatementSimpleStatementBlockAST(s)
	}
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitLoopSimpleStatementSimpleStatementBlockAST(s)
	}
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitLoopSimpleStatementSimpleStatementBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, miniGoParserRULE_loop)
	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLoopBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(542)
			p.Match(miniGoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(543)
			p.Block()
		}

	case 2:
		localctx = NewLoopExpressionBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(544)
			p.Match(miniGoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(545)
			p.expression(0)
		}
		{
			p.SetState(546)
			p.Block()
		}

	case 3:
		localctx = NewLoopSimpleStatementExpressionSimpleStatementBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(548)
			p.Match(miniGoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(549)
			p.SimpleStatement()
		}
		{
			p.SetState(550)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(551)
			p.expression(0)
		}
		{
			p.SetState(552)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(553)
			p.SimpleStatement()
		}
		{
			p.SetState(554)
			p.Block()
		}

	case 4:
		localctx = NewLoopSimpleStatementSimpleStatementBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(556)
			p.Match(miniGoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(557)
			p.SimpleStatement()
		}
		{
			p.SetState(558)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(559)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(560)
			p.SimpleStatement()
		}
		{
			p.SetState(561)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchStmtContext is an interface to support dynamic dispatch.
type ISwitchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitchStmtContext differentiates from other interfaces.
	IsSwitchStmtContext()
}

type SwitchStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStmtContext() *SwitchStmtContext {
	var p = new(SwitchStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_switchStmt
	return p
}

func InitEmptySwitchStmtContext(p *SwitchStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_switchStmt
}

func (*SwitchStmtContext) IsSwitchStmtContext() {}

func NewSwitchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_switchStmt

	return p
}

func (s *SwitchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStmtContext) CopyAll(ctx *SwitchStmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchStmtSimpleStatementASTContext struct {
	SwitchStmtContext
}

func NewSwitchStmtSimpleStatementASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtSimpleStatementASTContext {
	var p = new(SwitchStmtSimpleStatementASTContext)

	InitEmptySwitchStmtContext(&p.SwitchStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchStmtContext))

	return p
}

func (s *SwitchStmtSimpleStatementASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtSimpleStatementASTContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(miniGoParserSWITCH, 0)
}

func (s *SwitchStmtSimpleStatementASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SwitchStmtSimpleStatementASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *SwitchStmtSimpleStatementASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchStmtSimpleStatementASTContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(miniGoParserLBRACE, 0)
}

func (s *SwitchStmtSimpleStatementASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SwitchStmtSimpleStatementASTContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(miniGoParserRBRACE, 0)
}

func (s *SwitchStmtSimpleStatementASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSwitchStmtSimpleStatementAST(s)
	}
}

func (s *SwitchStmtSimpleStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSwitchStmtSimpleStatementAST(s)
	}
}

func (s *SwitchStmtSimpleStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSwitchStmtSimpleStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchStmtExpressionASTContext struct {
	SwitchStmtContext
}

func NewSwitchStmtExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtExpressionASTContext {
	var p = new(SwitchStmtExpressionASTContext)

	InitEmptySwitchStmtContext(&p.SwitchStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchStmtContext))

	return p
}

func (s *SwitchStmtExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtExpressionASTContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(miniGoParserSWITCH, 0)
}

func (s *SwitchStmtExpressionASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchStmtExpressionASTContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(miniGoParserLBRACE, 0)
}

func (s *SwitchStmtExpressionASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SwitchStmtExpressionASTContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(miniGoParserRBRACE, 0)
}

func (s *SwitchStmtExpressionASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSwitchStmtExpressionAST(s)
	}
}

func (s *SwitchStmtExpressionASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSwitchStmtExpressionAST(s)
	}
}

func (s *SwitchStmtExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSwitchStmtExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchStmtSimpleStatementBlockASTContext struct {
	SwitchStmtContext
}

func NewSwitchStmtSimpleStatementBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtSimpleStatementBlockASTContext {
	var p = new(SwitchStmtSimpleStatementBlockASTContext)

	InitEmptySwitchStmtContext(&p.SwitchStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchStmtContext))

	return p
}

func (s *SwitchStmtSimpleStatementBlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtSimpleStatementBlockASTContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(miniGoParserSWITCH, 0)
}

func (s *SwitchStmtSimpleStatementBlockASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SwitchStmtSimpleStatementBlockASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserSEMICOLON, 0)
}

func (s *SwitchStmtSimpleStatementBlockASTContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(miniGoParserLBRACE, 0)
}

func (s *SwitchStmtSimpleStatementBlockASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SwitchStmtSimpleStatementBlockASTContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(miniGoParserRBRACE, 0)
}

func (s *SwitchStmtSimpleStatementBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSwitchStmtSimpleStatementBlockAST(s)
	}
}

func (s *SwitchStmtSimpleStatementBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSwitchStmtSimpleStatementBlockAST(s)
	}
}

func (s *SwitchStmtSimpleStatementBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSwitchStmtSimpleStatementBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchStmtBlockASTContext struct {
	SwitchStmtContext
}

func NewSwitchStmtBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtBlockASTContext {
	var p = new(SwitchStmtBlockASTContext)

	InitEmptySwitchStmtContext(&p.SwitchStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchStmtContext))

	return p
}

func (s *SwitchStmtBlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtBlockASTContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(miniGoParserSWITCH, 0)
}

func (s *SwitchStmtBlockASTContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(miniGoParserLBRACE, 0)
}

func (s *SwitchStmtBlockASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SwitchStmtBlockASTContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(miniGoParserRBRACE, 0)
}

func (s *SwitchStmtBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterSwitchStmtBlockAST(s)
	}
}

func (s *SwitchStmtBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitSwitchStmtBlockAST(s)
	}
}

func (s *SwitchStmtBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitSwitchStmtBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) SwitchStmt() (localctx ISwitchStmtContext) {
	localctx = NewSwitchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, miniGoParserRULE_switchStmt)
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSwitchStmtSimpleStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(565)
			p.Match(miniGoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(566)
			p.SimpleStatement()
		}
		{
			p.SetState(567)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(568)
			p.expression(0)
		}
		{
			p.SetState(569)
			p.Match(miniGoParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(570)
			p.expressionCaseClauseList(0)
		}
		{
			p.SetState(571)
			p.Match(miniGoParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSwitchStmtExpressionASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(573)
			p.Match(miniGoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(574)
			p.expression(0)
		}
		{
			p.SetState(575)
			p.Match(miniGoParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(576)
			p.expressionCaseClauseList(0)
		}
		{
			p.SetState(577)
			p.Match(miniGoParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSwitchStmtSimpleStatementBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(579)
			p.Match(miniGoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(580)
			p.SimpleStatement()
		}
		{
			p.SetState(581)
			p.Match(miniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(582)
			p.Match(miniGoParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(583)
			p.expressionCaseClauseList(0)
		}
		{
			p.SetState(584)
			p.Match(miniGoParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewSwitchStmtBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(586)
			p.Match(miniGoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(587)
			p.Match(miniGoParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(588)
			p.expressionCaseClauseList(0)
		}
		{
			p.SetState(589)
			p.Match(miniGoParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionCaseClauseListContext is an interface to support dynamic dispatch.
type IExpressionCaseClauseListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionCaseClauseListContext differentiates from other interfaces.
	IsExpressionCaseClauseListContext()
}

type ExpressionCaseClauseListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionCaseClauseListContext() *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_expressionCaseClauseList
	return p
}

func InitEmptyExpressionCaseClauseListContext(p *ExpressionCaseClauseListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_expressionCaseClauseList
}

func (*ExpressionCaseClauseListContext) IsExpressionCaseClauseListContext() {}

func NewExpressionCaseClauseListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_expressionCaseClauseList

	return p
}

func (s *ExpressionCaseClauseListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseListContext) CopyAll(ctx *ExpressionCaseClauseListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionCaseClauseListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionCaseClauseListEmptyASTContext struct {
	ExpressionCaseClauseListContext
}

func NewExpressionCaseClauseListEmptyASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionCaseClauseListEmptyASTContext {
	var p = new(ExpressionCaseClauseListEmptyASTContext)

	InitEmptyExpressionCaseClauseListContext(&p.ExpressionCaseClauseListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionCaseClauseListContext))

	return p
}

func (s *ExpressionCaseClauseListEmptyASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseListEmptyASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *ExpressionCaseClauseListEmptyASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionCaseClauseListEmptyAST(s)
	}
}

func (s *ExpressionCaseClauseListEmptyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionCaseClauseListEmptyAST(s)
	}
}

func (s *ExpressionCaseClauseListEmptyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionCaseClauseListEmptyAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionCaseClauseListASTContext struct {
	ExpressionCaseClauseListContext
}

func NewExpressionCaseClauseListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionCaseClauseListASTContext {
	var p = new(ExpressionCaseClauseListASTContext)

	InitEmptyExpressionCaseClauseListContext(&p.ExpressionCaseClauseListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionCaseClauseListContext))

	return p
}

func (s *ExpressionCaseClauseListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseListASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *ExpressionCaseClauseListASTContext) ExpressionCaseClause() IExpressionCaseClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseContext)
}

func (s *ExpressionCaseClauseListASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionCaseClauseListAST(s)
	}
}

func (s *ExpressionCaseClauseListASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionCaseClauseListAST(s)
	}
}

func (s *ExpressionCaseClauseListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionCaseClauseListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) ExpressionCaseClauseList() (localctx IExpressionCaseClauseListContext) {
	return p.expressionCaseClauseList(0)
}

func (p *miniGoParser) expressionCaseClauseList(_p int) (localctx IExpressionCaseClauseListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionCaseClauseListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionCaseClauseListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 74
	p.EnterRecursionRule(localctx, 74, miniGoParserRULE_expressionCaseClauseList, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewExpressionCaseClauseListEmptyASTContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(594)
		p.Epsilon()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(600)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionCaseClauseListASTContext(p, NewExpressionCaseClauseListContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, miniGoParserRULE_expressionCaseClauseList)
			p.SetState(596)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(597)
				p.ExpressionCaseClause()
			}

		}
		p.SetState(602)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionCaseClauseContext is an interface to support dynamic dispatch.
type IExpressionCaseClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionCaseClauseContext differentiates from other interfaces.
	IsExpressionCaseClauseContext()
}

type ExpressionCaseClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionCaseClauseContext() *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_expressionCaseClause
	return p
}

func InitEmptyExpressionCaseClauseContext(p *ExpressionCaseClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_expressionCaseClause
}

func (*ExpressionCaseClauseContext) IsExpressionCaseClauseContext() {}

func NewExpressionCaseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_expressionCaseClause

	return p
}

func (s *ExpressionCaseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseContext) CopyAll(ctx *ExpressionCaseClauseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionCaseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionCaseClauseASTContext struct {
	ExpressionCaseClauseContext
}

func NewExpressionCaseClauseASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionCaseClauseASTContext {
	var p = new(ExpressionCaseClauseASTContext)

	InitEmptyExpressionCaseClauseContext(&p.ExpressionCaseClauseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionCaseClauseContext))

	return p
}

func (s *ExpressionCaseClauseASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseASTContext) ExpressionSwitchCase() IExpressionSwitchCaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSwitchCaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSwitchCaseContext)
}

func (s *ExpressionCaseClauseASTContext) COLON() antlr.TerminalNode {
	return s.GetToken(miniGoParserCOLON, 0)
}

func (s *ExpressionCaseClauseASTContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *ExpressionCaseClauseASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionCaseClauseAST(s)
	}
}

func (s *ExpressionCaseClauseASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionCaseClauseAST(s)
	}
}

func (s *ExpressionCaseClauseASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionCaseClauseAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) ExpressionCaseClause() (localctx IExpressionCaseClauseContext) {
	localctx = NewExpressionCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, miniGoParserRULE_expressionCaseClause)
	localctx = NewExpressionCaseClauseASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(603)
		p.ExpressionSwitchCase()
	}
	{
		p.SetState(604)
		p.Match(miniGoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(605)
		p.StatementList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionSwitchCaseContext is an interface to support dynamic dispatch.
type IExpressionSwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionSwitchCaseContext differentiates from other interfaces.
	IsExpressionSwitchCaseContext()
}

type ExpressionSwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSwitchCaseContext() *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_expressionSwitchCase
	return p
}

func InitEmptyExpressionSwitchCaseContext(p *ExpressionSwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_expressionSwitchCase
}

func (*ExpressionSwitchCaseContext) IsExpressionSwitchCaseContext() {}

func NewExpressionSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_expressionSwitchCase

	return p
}

func (s *ExpressionSwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSwitchCaseContext) CopyAll(ctx *ExpressionSwitchCaseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionSwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionSwitchCaseASTContext struct {
	ExpressionSwitchCaseContext
}

func NewExpressionSwitchCaseASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionSwitchCaseASTContext {
	var p = new(ExpressionSwitchCaseASTContext)

	InitEmptyExpressionSwitchCaseContext(&p.ExpressionSwitchCaseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSwitchCaseContext))

	return p
}

func (s *ExpressionSwitchCaseASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSwitchCaseASTContext) CASE() antlr.TerminalNode {
	return s.GetToken(miniGoParserCASE, 0)
}

func (s *ExpressionSwitchCaseASTContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ExpressionSwitchCaseASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionSwitchCaseAST(s)
	}
}

func (s *ExpressionSwitchCaseASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionSwitchCaseAST(s)
	}
}

func (s *ExpressionSwitchCaseASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionSwitchCaseAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionSwitchDefaultASTContext struct {
	ExpressionSwitchCaseContext
}

func NewExpressionSwitchDefaultASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionSwitchDefaultASTContext {
	var p = new(ExpressionSwitchDefaultASTContext)

	InitEmptyExpressionSwitchCaseContext(&p.ExpressionSwitchCaseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSwitchCaseContext))

	return p
}

func (s *ExpressionSwitchDefaultASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSwitchDefaultASTContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(miniGoParserDEFAULT, 0)
}

func (s *ExpressionSwitchDefaultASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterExpressionSwitchDefaultAST(s)
	}
}

func (s *ExpressionSwitchDefaultASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitExpressionSwitchDefaultAST(s)
	}
}

func (s *ExpressionSwitchDefaultASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitExpressionSwitchDefaultAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) ExpressionSwitchCase() (localctx IExpressionSwitchCaseContext) {
	localctx = NewExpressionSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, miniGoParserRULE_expressionSwitchCase)
	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniGoParserCASE:
		localctx = NewExpressionSwitchCaseASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(607)
			p.Match(miniGoParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(608)
			p.ExpressionList()
		}

	case miniGoParserDEFAULT:
		localctx = NewExpressionSwitchDefaultASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(609)
			p.Match(miniGoParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEpsilonContext is an interface to support dynamic dispatch.
type IEpsilonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEpsilonContext differentiates from other interfaces.
	IsEpsilonContext()
}

type EpsilonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEpsilonContext() *EpsilonContext {
	var p = new(EpsilonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_epsilon
	return p
}

func InitEmptyEpsilonContext(p *EpsilonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniGoParserRULE_epsilon
}

func (*EpsilonContext) IsEpsilonContext() {}

func NewEpsilonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EpsilonContext {
	var p = new(EpsilonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniGoParserRULE_epsilon

	return p
}

func (s *EpsilonContext) GetParser() antlr.Parser { return s.parser }

func (s *EpsilonContext) CopyAll(ctx *EpsilonContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EpsilonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpsilonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EpsilonASTContext struct {
	EpsilonContext
}

func NewEpsilonASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EpsilonASTContext {
	var p = new(EpsilonASTContext)

	InitEmptyEpsilonContext(&p.EpsilonContext)
	p.parser = parser
	p.CopyAll(ctx.(*EpsilonContext))

	return p
}

func (s *EpsilonASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpsilonASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.EnterEpsilonAST(s)
	}
}

func (s *EpsilonASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniGoParserListener); ok {
		listenerT.ExitEpsilonAST(s)
	}
}

func (s *EpsilonASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case miniGoParserVisitor:
		return t.VisitEpsilonAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *miniGoParser) Epsilon() (localctx IEpsilonContext) {
	localctx = NewEpsilonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, miniGoParserRULE_epsilon)
	localctx = NewEpsilonASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *miniGoParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 20:
		var t *PrimaryExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExpressionContext)
		}
		return p.PrimaryExpression_Sempred(t, predIndex)

	case 37:
		var t *ExpressionCaseClauseListContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionCaseClauseListContext)
		}
		return p.ExpressionCaseClauseList_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *miniGoParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *miniGoParser) PrimaryExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 19:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 20:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 21:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *miniGoParser) ExpressionCaseClauseList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 22:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
