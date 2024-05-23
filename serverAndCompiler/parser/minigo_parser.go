// Code generated from C:/Users/noni/Desktop/miniGoProyect/serverAndCompiler/MiniGoParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // MiniGoParser
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

type MiniGoParser struct {
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
		"funcDecl", "funcFrontDecl", "multipleReturnTypes", "returnTypeList",
		"singleReturnType", "funcArgDecls", "declType", "sliceDeclType", "arrayDeclType",
		"structDeclType", "structMemDecls", "identifierList", "expression",
		"expressionList", "primaryExpression", "operand", "literal", "index",
		"arguments", "selector", "appendExpression", "lengthExpression", "capExpression",
		"statementList", "block", "statement", "simpleStatement", "assignmentStatement",
		"ifStatement", "loop", "switchStmt", "expressionCaseClauseList", "expressionCaseClause",
		"expressionSwitchCase", "epsilon",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 72, 637, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		5, 1, 97, 8, 1, 10, 1, 12, 1, 100, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 116, 8, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 123, 8, 3, 10, 3, 12, 3, 126, 9, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 138,
		8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 157, 8, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 5, 7, 164, 8, 7, 10, 7, 12, 7, 167, 9, 7, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 181,
		8, 10, 1, 10, 1, 10, 1, 10, 3, 10, 186, 8, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 5, 12, 195, 8, 12, 10, 12, 12, 12, 198, 9, 12,
		1, 13, 1, 13, 3, 13, 202, 8, 13, 1, 14, 1, 14, 1, 14, 5, 14, 207, 8, 14,
		10, 14, 12, 14, 210, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 3, 15, 220, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 235, 8, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 244, 8, 19, 10, 19,
		12, 19, 247, 9, 19, 1, 20, 1, 20, 1, 20, 5, 20, 252, 8, 20, 10, 20, 12,
		20, 255, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 3, 21, 267, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21,
		326, 8, 21, 10, 21, 12, 21, 329, 9, 21, 1, 22, 1, 22, 1, 22, 5, 22, 334,
		8, 22, 10, 22, 12, 22, 337, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3,
		23, 344, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 352, 8,
		23, 10, 23, 12, 23, 355, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		3, 24, 363, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 370, 8, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 381,
		8, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 32, 5, 32, 404, 8, 32, 10, 32, 12, 32, 407, 9, 32, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 417, 8, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 426, 8, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 3, 34, 434, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 459,
		8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 466, 8, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 3, 35, 473, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 3, 36, 523, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 3, 37, 563, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 586, 8, 38, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 3, 39, 614, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40,
		621, 8, 40, 10, 40, 12, 40, 624, 9, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		42, 1, 42, 1, 42, 3, 42, 633, 8, 42, 1, 43, 1, 43, 1, 43, 0, 3, 42, 46,
		80, 44, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 82, 84, 86, 0, 0, 696, 0, 88, 1, 0, 0, 0, 2, 98,
		1, 0, 0, 0, 4, 115, 1, 0, 0, 0, 6, 117, 1, 0, 0, 0, 8, 137, 1, 0, 0, 0,
		10, 139, 1, 0, 0, 0, 12, 156, 1, 0, 0, 0, 14, 158, 1, 0, 0, 0, 16, 168,
		1, 0, 0, 0, 18, 171, 1, 0, 0, 0, 20, 175, 1, 0, 0, 0, 22, 187, 1, 0, 0,
		0, 24, 191, 1, 0, 0, 0, 26, 201, 1, 0, 0, 0, 28, 203, 1, 0, 0, 0, 30, 219,
		1, 0, 0, 0, 32, 221, 1, 0, 0, 0, 34, 225, 1, 0, 0, 0, 36, 230, 1, 0, 0,
		0, 38, 238, 1, 0, 0, 0, 40, 248, 1, 0, 0, 0, 42, 266, 1, 0, 0, 0, 44, 330,
		1, 0, 0, 0, 46, 343, 1, 0, 0, 0, 48, 362, 1, 0, 0, 0, 50, 369, 1, 0, 0,
		0, 52, 371, 1, 0, 0, 0, 54, 380, 1, 0, 0, 0, 56, 382, 1, 0, 0, 0, 58, 385,
		1, 0, 0, 0, 60, 392, 1, 0, 0, 0, 62, 397, 1, 0, 0, 0, 64, 405, 1, 0, 0,
		0, 66, 408, 1, 0, 0, 0, 68, 458, 1, 0, 0, 0, 70, 472, 1, 0, 0, 0, 72, 522,
		1, 0, 0, 0, 74, 562, 1, 0, 0, 0, 76, 585, 1, 0, 0, 0, 78, 613, 1, 0, 0,
		0, 80, 615, 1, 0, 0, 0, 82, 625, 1, 0, 0, 0, 84, 632, 1, 0, 0, 0, 86, 634,
		1, 0, 0, 0, 88, 89, 5, 6, 0, 0, 89, 90, 5, 25, 0, 0, 90, 91, 5, 61, 0,
		0, 91, 92, 3, 2, 1, 0, 92, 1, 1, 0, 0, 0, 93, 97, 3, 4, 2, 0, 94, 97, 3,
		12, 6, 0, 95, 97, 3, 18, 9, 0, 96, 93, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0,
		96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1,
		0, 0, 0, 99, 3, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 102, 5, 7, 0, 0,
		102, 103, 3, 8, 4, 0, 103, 104, 5, 61, 0, 0, 104, 116, 1, 0, 0, 0, 105,
		106, 5, 7, 0, 0, 106, 107, 5, 64, 0, 0, 107, 108, 3, 6, 3, 0, 108, 109,
		5, 65, 0, 0, 109, 110, 5, 61, 0, 0, 110, 116, 1, 0, 0, 0, 111, 112, 5,
		7, 0, 0, 112, 113, 5, 64, 0, 0, 113, 114, 5, 65, 0, 0, 114, 116, 5, 61,
		0, 0, 115, 101, 1, 0, 0, 0, 115, 105, 1, 0, 0, 0, 115, 111, 1, 0, 0, 0,
		116, 5, 1, 0, 0, 0, 117, 118, 3, 8, 4, 0, 118, 124, 5, 61, 0, 0, 119, 120,
		3, 8, 4, 0, 120, 121, 5, 61, 0, 0, 121, 123, 1, 0, 0, 0, 122, 119, 1, 0,
		0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0,
		125, 7, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 128, 3, 40, 20, 0, 128,
		129, 3, 30, 15, 0, 129, 130, 5, 48, 0, 0, 130, 131, 3, 44, 22, 0, 131,
		138, 1, 0, 0, 0, 132, 133, 3, 40, 20, 0, 133, 134, 5, 48, 0, 0, 134, 135,
		3, 44, 22, 0, 135, 138, 1, 0, 0, 0, 136, 138, 3, 10, 5, 0, 137, 127, 1,
		0, 0, 0, 137, 132, 1, 0, 0, 0, 137, 136, 1, 0, 0, 0, 138, 9, 1, 0, 0, 0,
		139, 140, 3, 40, 20, 0, 140, 141, 3, 30, 15, 0, 141, 11, 1, 0, 0, 0, 142,
		143, 5, 8, 0, 0, 143, 144, 3, 16, 8, 0, 144, 145, 5, 61, 0, 0, 145, 157,
		1, 0, 0, 0, 146, 147, 5, 8, 0, 0, 147, 148, 5, 64, 0, 0, 148, 149, 3, 14,
		7, 0, 149, 150, 5, 65, 0, 0, 150, 151, 5, 61, 0, 0, 151, 157, 1, 0, 0,
		0, 152, 153, 5, 8, 0, 0, 153, 154, 5, 64, 0, 0, 154, 155, 5, 65, 0, 0,
		155, 157, 5, 61, 0, 0, 156, 142, 1, 0, 0, 0, 156, 146, 1, 0, 0, 0, 156,
		152, 1, 0, 0, 0, 157, 13, 1, 0, 0, 0, 158, 159, 3, 16, 8, 0, 159, 165,
		5, 61, 0, 0, 160, 161, 3, 16, 8, 0, 161, 162, 5, 61, 0, 0, 162, 164, 1,
		0, 0, 0, 163, 160, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165, 163, 1, 0, 0,
		0, 165, 166, 1, 0, 0, 0, 166, 15, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 168,
		169, 5, 25, 0, 0, 169, 170, 3, 30, 15, 0, 170, 17, 1, 0, 0, 0, 171, 172,
		3, 20, 10, 0, 172, 173, 3, 66, 33, 0, 173, 174, 5, 61, 0, 0, 174, 19, 1,
		0, 0, 0, 175, 176, 5, 9, 0, 0, 176, 177, 5, 25, 0, 0, 177, 180, 5, 64,
		0, 0, 178, 181, 3, 28, 14, 0, 179, 181, 3, 86, 43, 0, 180, 178, 1, 0, 0,
		0, 180, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 185, 5, 65, 0, 0, 183,
		186, 3, 22, 11, 0, 184, 186, 3, 26, 13, 0, 185, 183, 1, 0, 0, 0, 185, 184,
		1, 0, 0, 0, 186, 21, 1, 0, 0, 0, 187, 188, 5, 64, 0, 0, 188, 189, 3, 24,
		12, 0, 189, 190, 5, 65, 0, 0, 190, 23, 1, 0, 0, 0, 191, 196, 3, 30, 15,
		0, 192, 193, 5, 62, 0, 0, 193, 195, 3, 30, 15, 0, 194, 192, 1, 0, 0, 0,
		195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197,
		25, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 202, 3, 30, 15, 0, 200, 202,
		3, 86, 43, 0, 201, 199, 1, 0, 0, 0, 201, 200, 1, 0, 0, 0, 202, 27, 1, 0,
		0, 0, 203, 208, 3, 10, 5, 0, 204, 205, 5, 62, 0, 0, 205, 207, 3, 10, 5,
		0, 206, 204, 1, 0, 0, 0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208,
		209, 1, 0, 0, 0, 209, 29, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 211, 212, 5,
		64, 0, 0, 212, 213, 3, 30, 15, 0, 213, 214, 5, 65, 0, 0, 214, 220, 1, 0,
		0, 0, 215, 220, 5, 25, 0, 0, 216, 220, 3, 32, 16, 0, 217, 220, 3, 34, 17,
		0, 218, 220, 3, 36, 18, 0, 219, 211, 1, 0, 0, 0, 219, 215, 1, 0, 0, 0,
		219, 216, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 218, 1, 0, 0, 0, 220,
		31, 1, 0, 0, 0, 221, 222, 5, 66, 0, 0, 222, 223, 5, 67, 0, 0, 223, 224,
		3, 30, 15, 0, 224, 33, 1, 0, 0, 0, 225, 226, 5, 66, 0, 0, 226, 227, 5,
		1, 0, 0, 227, 228, 5, 67, 0, 0, 228, 229, 3, 30, 15, 0, 229, 35, 1, 0,
		0, 0, 230, 231, 5, 10, 0, 0, 231, 234, 5, 68, 0, 0, 232, 235, 3, 38, 19,
		0, 233, 235, 3, 86, 43, 0, 234, 232, 1, 0, 0, 0, 234, 233, 1, 0, 0, 0,
		235, 236, 1, 0, 0, 0, 236, 237, 5, 69, 0, 0, 237, 37, 1, 0, 0, 0, 238,
		239, 3, 10, 5, 0, 239, 245, 5, 61, 0, 0, 240, 241, 3, 10, 5, 0, 241, 242,
		5, 61, 0, 0, 242, 244, 1, 0, 0, 0, 243, 240, 1, 0, 0, 0, 244, 247, 1, 0,
		0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 39, 1, 0, 0, 0,
		247, 245, 1, 0, 0, 0, 248, 253, 5, 25, 0, 0, 249, 250, 5, 62, 0, 0, 250,
		252, 5, 25, 0, 0, 251, 249, 1, 0, 0, 0, 252, 255, 1, 0, 0, 0, 253, 251,
		1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 41, 1, 0, 0, 0, 255, 253, 1, 0,
		0, 0, 256, 257, 6, 21, -1, 0, 257, 267, 3, 46, 23, 0, 258, 259, 5, 26,
		0, 0, 259, 267, 3, 42, 21, 4, 260, 261, 5, 27, 0, 0, 261, 267, 3, 42, 21,
		3, 262, 263, 5, 39, 0, 0, 263, 267, 3, 42, 21, 2, 264, 265, 5, 42, 0, 0,
		265, 267, 3, 42, 21, 1, 266, 256, 1, 0, 0, 0, 266, 258, 1, 0, 0, 0, 266,
		260, 1, 0, 0, 0, 266, 262, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 267, 327,
		1, 0, 0, 0, 268, 269, 10, 23, 0, 0, 269, 270, 5, 28, 0, 0, 270, 326, 3,
		42, 21, 24, 271, 272, 10, 22, 0, 0, 272, 273, 5, 29, 0, 0, 273, 326, 3,
		42, 21, 23, 274, 275, 10, 21, 0, 0, 275, 276, 5, 30, 0, 0, 276, 326, 3,
		42, 21, 22, 277, 278, 10, 20, 0, 0, 278, 279, 5, 44, 0, 0, 279, 326, 3,
		42, 21, 21, 280, 281, 10, 19, 0, 0, 281, 282, 5, 45, 0, 0, 282, 326, 3,
		42, 21, 20, 283, 284, 10, 18, 0, 0, 284, 285, 5, 40, 0, 0, 285, 326, 3,
		42, 21, 19, 286, 287, 10, 17, 0, 0, 287, 288, 5, 43, 0, 0, 288, 326, 3,
		42, 21, 18, 289, 290, 10, 16, 0, 0, 290, 291, 5, 26, 0, 0, 291, 326, 3,
		42, 21, 17, 292, 293, 10, 15, 0, 0, 293, 294, 5, 27, 0, 0, 294, 326, 3,
		42, 21, 16, 295, 296, 10, 14, 0, 0, 296, 297, 5, 41, 0, 0, 297, 326, 3,
		42, 21, 15, 298, 299, 10, 13, 0, 0, 299, 300, 5, 42, 0, 0, 300, 326, 3,
		42, 21, 14, 301, 302, 10, 12, 0, 0, 302, 303, 5, 35, 0, 0, 303, 326, 3,
		42, 21, 13, 304, 305, 10, 11, 0, 0, 305, 306, 5, 36, 0, 0, 306, 326, 3,
		42, 21, 12, 307, 308, 10, 10, 0, 0, 308, 309, 5, 31, 0, 0, 309, 326, 3,
		42, 21, 11, 310, 311, 10, 9, 0, 0, 311, 312, 5, 32, 0, 0, 312, 326, 3,
		42, 21, 10, 313, 314, 10, 8, 0, 0, 314, 315, 5, 33, 0, 0, 315, 326, 3,
		42, 21, 9, 316, 317, 10, 7, 0, 0, 317, 318, 5, 34, 0, 0, 318, 326, 3, 42,
		21, 8, 319, 320, 10, 6, 0, 0, 320, 321, 5, 37, 0, 0, 321, 326, 3, 42, 21,
		7, 322, 323, 10, 5, 0, 0, 323, 324, 5, 38, 0, 0, 324, 326, 3, 42, 21, 6,
		325, 268, 1, 0, 0, 0, 325, 271, 1, 0, 0, 0, 325, 274, 1, 0, 0, 0, 325,
		277, 1, 0, 0, 0, 325, 280, 1, 0, 0, 0, 325, 283, 1, 0, 0, 0, 325, 286,
		1, 0, 0, 0, 325, 289, 1, 0, 0, 0, 325, 292, 1, 0, 0, 0, 325, 295, 1, 0,
		0, 0, 325, 298, 1, 0, 0, 0, 325, 301, 1, 0, 0, 0, 325, 304, 1, 0, 0, 0,
		325, 307, 1, 0, 0, 0, 325, 310, 1, 0, 0, 0, 325, 313, 1, 0, 0, 0, 325,
		316, 1, 0, 0, 0, 325, 319, 1, 0, 0, 0, 325, 322, 1, 0, 0, 0, 326, 329,
		1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 43, 1, 0,
		0, 0, 329, 327, 1, 0, 0, 0, 330, 335, 3, 42, 21, 0, 331, 332, 5, 62, 0,
		0, 332, 334, 3, 42, 21, 0, 333, 331, 1, 0, 0, 0, 334, 337, 1, 0, 0, 0,
		335, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 45, 1, 0, 0, 0, 337, 335,
		1, 0, 0, 0, 338, 339, 6, 23, -1, 0, 339, 344, 3, 48, 24, 0, 340, 344, 3,
		58, 29, 0, 341, 344, 3, 60, 30, 0, 342, 344, 3, 62, 31, 0, 343, 338, 1,
		0, 0, 0, 343, 340, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 343, 342, 1, 0, 0,
		0, 344, 353, 1, 0, 0, 0, 345, 346, 10, 6, 0, 0, 346, 352, 3, 56, 28, 0,
		347, 348, 10, 5, 0, 0, 348, 352, 3, 52, 26, 0, 349, 350, 10, 4, 0, 0, 350,
		352, 3, 54, 27, 0, 351, 345, 1, 0, 0, 0, 351, 347, 1, 0, 0, 0, 351, 349,
		1, 0, 0, 0, 352, 355, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 353, 354, 1, 0,
		0, 0, 354, 47, 1, 0, 0, 0, 355, 353, 1, 0, 0, 0, 356, 363, 3, 50, 25, 0,
		357, 363, 5, 25, 0, 0, 358, 359, 5, 64, 0, 0, 359, 360, 3, 42, 21, 0, 360,
		361, 5, 65, 0, 0, 361, 363, 1, 0, 0, 0, 362, 356, 1, 0, 0, 0, 362, 357,
		1, 0, 0, 0, 362, 358, 1, 0, 0, 0, 363, 49, 1, 0, 0, 0, 364, 370, 5, 1,
		0, 0, 365, 370, 5, 2, 0, 0, 366, 370, 5, 3, 0, 0, 367, 370, 5, 4, 0, 0,
		368, 370, 5, 5, 0, 0, 369, 364, 1, 0, 0, 0, 369, 365, 1, 0, 0, 0, 369,
		366, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0, 369, 368, 1, 0, 0, 0, 370, 51, 1,
		0, 0, 0, 371, 372, 5, 66, 0, 0, 372, 373, 3, 42, 21, 0, 373, 374, 5, 67,
		0, 0, 374, 53, 1, 0, 0, 0, 375, 376, 5, 64, 0, 0, 376, 381, 3, 44, 22,
		0, 377, 378, 3, 86, 43, 0, 378, 379, 5, 65, 0, 0, 379, 381, 1, 0, 0, 0,
		380, 375, 1, 0, 0, 0, 380, 377, 1, 0, 0, 0, 381, 55, 1, 0, 0, 0, 382, 383,
		5, 63, 0, 0, 383, 384, 5, 25, 0, 0, 384, 57, 1, 0, 0, 0, 385, 386, 5, 18,
		0, 0, 386, 387, 5, 64, 0, 0, 387, 388, 3, 42, 21, 0, 388, 389, 5, 62, 0,
		0, 389, 390, 3, 42, 21, 0, 390, 391, 5, 65, 0, 0, 391, 59, 1, 0, 0, 0,
		392, 393, 5, 11, 0, 0, 393, 394, 5, 64, 0, 0, 394, 395, 3, 42, 21, 0, 395,
		396, 5, 65, 0, 0, 396, 61, 1, 0, 0, 0, 397, 398, 5, 12, 0, 0, 398, 399,
		5, 64, 0, 0, 399, 400, 3, 42, 21, 0, 400, 401, 5, 65, 0, 0, 401, 63, 1,
		0, 0, 0, 402, 404, 3, 68, 34, 0, 403, 402, 1, 0, 0, 0, 404, 407, 1, 0,
		0, 0, 405, 403, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 65, 1, 0, 0, 0,
		407, 405, 1, 0, 0, 0, 408, 409, 5, 68, 0, 0, 409, 410, 3, 64, 32, 0, 410,
		411, 5, 69, 0, 0, 411, 67, 1, 0, 0, 0, 412, 413, 5, 13, 0, 0, 413, 416,
		5, 64, 0, 0, 414, 417, 3, 44, 22, 0, 415, 417, 3, 86, 43, 0, 416, 414,
		1, 0, 0, 0, 416, 415, 1, 0, 0, 0, 417, 418, 1, 0, 0, 0, 418, 419, 5, 65,
		0, 0, 419, 420, 5, 61, 0, 0, 420, 459, 1, 0, 0, 0, 421, 422, 5, 14, 0,
		0, 422, 425, 5, 64, 0, 0, 423, 426, 3, 44, 22, 0, 424, 426, 3, 86, 43,
		0, 425, 423, 1, 0, 0, 0, 425, 424, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427,
		428, 5, 65, 0, 0, 428, 429, 5, 61, 0, 0, 429, 459, 1, 0, 0, 0, 430, 433,
		5, 15, 0, 0, 431, 434, 3, 42, 21, 0, 432, 434, 3, 86, 43, 0, 433, 431,
		1, 0, 0, 0, 433, 432, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 436, 5, 61,
		0, 0, 436, 459, 1, 0, 0, 0, 437, 438, 5, 16, 0, 0, 438, 459, 5, 61, 0,
		0, 439, 440, 5, 17, 0, 0, 440, 459, 5, 61, 0, 0, 441, 442, 3, 70, 35, 0,
		442, 443, 5, 61, 0, 0, 443, 459, 1, 0, 0, 0, 444, 445, 3, 66, 33, 0, 445,
		446, 5, 61, 0, 0, 446, 459, 1, 0, 0, 0, 447, 448, 3, 78, 39, 0, 448, 449,
		5, 61, 0, 0, 449, 459, 1, 0, 0, 0, 450, 451, 3, 74, 37, 0, 451, 452, 5,
		61, 0, 0, 452, 459, 1, 0, 0, 0, 453, 454, 3, 76, 38, 0, 454, 455, 5, 61,
		0, 0, 455, 459, 1, 0, 0, 0, 456, 459, 3, 12, 6, 0, 457, 459, 3, 4, 2, 0,
		458, 412, 1, 0, 0, 0, 458, 421, 1, 0, 0, 0, 458, 430, 1, 0, 0, 0, 458,
		437, 1, 0, 0, 0, 458, 439, 1, 0, 0, 0, 458, 441, 1, 0, 0, 0, 458, 444,
		1, 0, 0, 0, 458, 447, 1, 0, 0, 0, 458, 450, 1, 0, 0, 0, 458, 453, 1, 0,
		0, 0, 458, 456, 1, 0, 0, 0, 458, 457, 1, 0, 0, 0, 459, 69, 1, 0, 0, 0,
		460, 473, 3, 86, 43, 0, 461, 465, 3, 42, 21, 0, 462, 466, 5, 46, 0, 0,
		463, 466, 5, 47, 0, 0, 464, 466, 3, 86, 43, 0, 465, 462, 1, 0, 0, 0, 465,
		463, 1, 0, 0, 0, 465, 464, 1, 0, 0, 0, 466, 473, 1, 0, 0, 0, 467, 473,
		3, 72, 36, 0, 468, 469, 3, 44, 22, 0, 469, 470, 5, 48, 0, 0, 470, 471,
		3, 44, 22, 0, 471, 473, 1, 0, 0, 0, 472, 460, 1, 0, 0, 0, 472, 461, 1,
		0, 0, 0, 472, 467, 1, 0, 0, 0, 472, 468, 1, 0, 0, 0, 473, 71, 1, 0, 0,
		0, 474, 475, 3, 44, 22, 0, 475, 476, 5, 48, 0, 0, 476, 477, 3, 44, 22,
		0, 477, 523, 1, 0, 0, 0, 478, 479, 3, 42, 21, 0, 479, 480, 5, 49, 0, 0,
		480, 481, 3, 42, 21, 0, 481, 523, 1, 0, 0, 0, 482, 483, 3, 42, 21, 0, 483,
		484, 5, 50, 0, 0, 484, 485, 3, 42, 21, 0, 485, 523, 1, 0, 0, 0, 486, 487,
		3, 42, 21, 0, 487, 488, 5, 54, 0, 0, 488, 489, 3, 42, 21, 0, 489, 523,
		1, 0, 0, 0, 490, 491, 3, 42, 21, 0, 491, 492, 5, 55, 0, 0, 492, 493, 3,
		42, 21, 0, 493, 523, 1, 0, 0, 0, 494, 495, 3, 42, 21, 0, 495, 496, 5, 51,
		0, 0, 496, 497, 3, 42, 21, 0, 497, 523, 1, 0, 0, 0, 498, 499, 3, 42, 21,
		0, 499, 500, 5, 56, 0, 0, 500, 501, 3, 42, 21, 0, 501, 523, 1, 0, 0, 0,
		502, 503, 3, 42, 21, 0, 503, 504, 5, 57, 0, 0, 504, 505, 3, 42, 21, 0,
		505, 523, 1, 0, 0, 0, 506, 507, 3, 42, 21, 0, 507, 508, 5, 58, 0, 0, 508,
		509, 3, 42, 21, 0, 509, 523, 1, 0, 0, 0, 510, 511, 3, 42, 21, 0, 511, 512,
		5, 59, 0, 0, 512, 513, 3, 42, 21, 0, 513, 523, 1, 0, 0, 0, 514, 515, 3,
		42, 21, 0, 515, 516, 5, 53, 0, 0, 516, 517, 3, 42, 21, 0, 517, 523, 1,
		0, 0, 0, 518, 519, 3, 42, 21, 0, 519, 520, 5, 52, 0, 0, 520, 521, 3, 42,
		21, 0, 521, 523, 1, 0, 0, 0, 522, 474, 1, 0, 0, 0, 522, 478, 1, 0, 0, 0,
		522, 482, 1, 0, 0, 0, 522, 486, 1, 0, 0, 0, 522, 490, 1, 0, 0, 0, 522,
		494, 1, 0, 0, 0, 522, 498, 1, 0, 0, 0, 522, 502, 1, 0, 0, 0, 522, 506,
		1, 0, 0, 0, 522, 510, 1, 0, 0, 0, 522, 514, 1, 0, 0, 0, 522, 518, 1, 0,
		0, 0, 523, 73, 1, 0, 0, 0, 524, 525, 5, 19, 0, 0, 525, 526, 3, 42, 21,
		0, 526, 527, 3, 66, 33, 0, 527, 563, 1, 0, 0, 0, 528, 529, 5, 19, 0, 0,
		529, 530, 3, 42, 21, 0, 530, 531, 3, 66, 33, 0, 531, 532, 5, 20, 0, 0,
		532, 533, 3, 74, 37, 0, 533, 563, 1, 0, 0, 0, 534, 535, 5, 19, 0, 0, 535,
		536, 3, 42, 21, 0, 536, 537, 3, 66, 33, 0, 537, 538, 5, 20, 0, 0, 538,
		539, 3, 66, 33, 0, 539, 563, 1, 0, 0, 0, 540, 541, 5, 19, 0, 0, 541, 542,
		3, 70, 35, 0, 542, 543, 5, 61, 0, 0, 543, 544, 3, 42, 21, 0, 544, 545,
		3, 66, 33, 0, 545, 563, 1, 0, 0, 0, 546, 547, 5, 19, 0, 0, 547, 548, 3,
		70, 35, 0, 548, 549, 5, 61, 0, 0, 549, 550, 3, 42, 21, 0, 550, 551, 3,
		66, 33, 0, 551, 552, 5, 20, 0, 0, 552, 553, 3, 74, 37, 0, 553, 563, 1,
		0, 0, 0, 554, 555, 5, 19, 0, 0, 555, 556, 3, 70, 35, 0, 556, 557, 5, 61,
		0, 0, 557, 558, 3, 42, 21, 0, 558, 559, 3, 66, 33, 0, 559, 560, 5, 20,
		0, 0, 560, 561, 3, 66, 33, 0, 561, 563, 1, 0, 0, 0, 562, 524, 1, 0, 0,
		0, 562, 528, 1, 0, 0, 0, 562, 534, 1, 0, 0, 0, 562, 540, 1, 0, 0, 0, 562,
		546, 1, 0, 0, 0, 562, 554, 1, 0, 0, 0, 563, 75, 1, 0, 0, 0, 564, 565, 5,
		21, 0, 0, 565, 586, 3, 66, 33, 0, 566, 567, 5, 21, 0, 0, 567, 568, 3, 42,
		21, 0, 568, 569, 3, 66, 33, 0, 569, 586, 1, 0, 0, 0, 570, 571, 5, 21, 0,
		0, 571, 572, 3, 70, 35, 0, 572, 573, 5, 61, 0, 0, 573, 574, 3, 42, 21,
		0, 574, 575, 5, 61, 0, 0, 575, 576, 3, 70, 35, 0, 576, 577, 3, 66, 33,
		0, 577, 586, 1, 0, 0, 0, 578, 579, 5, 21, 0, 0, 579, 580, 3, 70, 35, 0,
		580, 581, 5, 61, 0, 0, 581, 582, 5, 61, 0, 0, 582, 583, 3, 70, 35, 0, 583,
		584, 3, 66, 33, 0, 584, 586, 1, 0, 0, 0, 585, 564, 1, 0, 0, 0, 585, 566,
		1, 0, 0, 0, 585, 570, 1, 0, 0, 0, 585, 578, 1, 0, 0, 0, 586, 77, 1, 0,
		0, 0, 587, 588, 5, 22, 0, 0, 588, 589, 3, 70, 35, 0, 589, 590, 5, 61, 0,
		0, 590, 591, 3, 42, 21, 0, 591, 592, 5, 68, 0, 0, 592, 593, 3, 80, 40,
		0, 593, 594, 5, 69, 0, 0, 594, 614, 1, 0, 0, 0, 595, 596, 5, 22, 0, 0,
		596, 597, 3, 42, 21, 0, 597, 598, 5, 68, 0, 0, 598, 599, 3, 80, 40, 0,
		599, 600, 5, 69, 0, 0, 600, 614, 1, 0, 0, 0, 601, 602, 5, 22, 0, 0, 602,
		603, 3, 70, 35, 0, 603, 604, 5, 61, 0, 0, 604, 605, 5, 68, 0, 0, 605, 606,
		3, 80, 40, 0, 606, 607, 5, 69, 0, 0, 607, 614, 1, 0, 0, 0, 608, 609, 5,
		22, 0, 0, 609, 610, 5, 68, 0, 0, 610, 611, 3, 80, 40, 0, 611, 612, 5, 69,
		0, 0, 612, 614, 1, 0, 0, 0, 613, 587, 1, 0, 0, 0, 613, 595, 1, 0, 0, 0,
		613, 601, 1, 0, 0, 0, 613, 608, 1, 0, 0, 0, 614, 79, 1, 0, 0, 0, 615, 616,
		6, 40, -1, 0, 616, 617, 3, 86, 43, 0, 617, 622, 1, 0, 0, 0, 618, 619, 10,
		1, 0, 0, 619, 621, 3, 82, 41, 0, 620, 618, 1, 0, 0, 0, 621, 624, 1, 0,
		0, 0, 622, 620, 1, 0, 0, 0, 622, 623, 1, 0, 0, 0, 623, 81, 1, 0, 0, 0,
		624, 622, 1, 0, 0, 0, 625, 626, 3, 84, 42, 0, 626, 627, 5, 60, 0, 0, 627,
		628, 3, 64, 32, 0, 628, 83, 1, 0, 0, 0, 629, 630, 5, 23, 0, 0, 630, 633,
		3, 44, 22, 0, 631, 633, 5, 24, 0, 0, 632, 629, 1, 0, 0, 0, 632, 631, 1,
		0, 0, 0, 633, 85, 1, 0, 0, 0, 634, 635, 1, 0, 0, 0, 635, 87, 1, 0, 0, 0,
		39, 96, 98, 115, 124, 137, 156, 165, 180, 185, 196, 201, 208, 219, 234,
		245, 253, 266, 325, 327, 335, 343, 351, 353, 362, 369, 380, 405, 416, 425,
		433, 458, 465, 472, 522, 562, 585, 613, 622, 632,
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

// MiniGoParserInit initializes any static state used to implement MiniGoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMiniGoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MiniGoParserInit() {
	staticData := &MiniGoParserParserStaticData
	staticData.once.Do(minigoparserParserInit)
}

// NewMiniGoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMiniGoParser(input antlr.TokenStream) *MiniGoParser {
	MiniGoParserInit()
	this := new(MiniGoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MiniGoParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MiniGoParser.g4"

	return this
}

// MiniGoParser tokens.
const (
	MiniGoParserEOF                      = antlr.TokenEOF
	MiniGoParserINTLITERAL               = 1
	MiniGoParserFLOATLITERAL             = 2
	MiniGoParserRUNELITERAL              = 3
	MiniGoParserRAWSTRINGLITERAL         = 4
	MiniGoParserINTERPRETEDSTRINGLITERAL = 5
	MiniGoParserPACKAGE                  = 6
	MiniGoParserVAR                      = 7
	MiniGoParserTYPE                     = 8
	MiniGoParserFUNC                     = 9
	MiniGoParserSTRUCT                   = 10
	MiniGoParserLEN                      = 11
	MiniGoParserCAP                      = 12
	MiniGoParserPRINT                    = 13
	MiniGoParserPRINTLN                  = 14
	MiniGoParserRETURN                   = 15
	MiniGoParserBREAK                    = 16
	MiniGoParserCONTINUE                 = 17
	MiniGoParserAPPEND                   = 18
	MiniGoParserIF                       = 19
	MiniGoParserELSE                     = 20
	MiniGoParserFOR                      = 21
	MiniGoParserSWITCH                   = 22
	MiniGoParserCASE                     = 23
	MiniGoParserDEFAULT                  = 24
	MiniGoParserIDENTIFIER               = 25
	MiniGoParserPLUS                     = 26
	MiniGoParserMINUS                    = 27
	MiniGoParserMULTIPLY                 = 28
	MiniGoParserDIVIDE                   = 29
	MiniGoParserMODULO                   = 30
	MiniGoParserLESS                     = 31
	MiniGoParserLESSEQUAL                = 32
	MiniGoParserGREATER                  = 33
	MiniGoParserGREATEREQUAL             = 34
	MiniGoParserEQUAL                    = 35
	MiniGoParserNOTEQUAL                 = 36
	MiniGoParserAND                      = 37
	MiniGoParserOR                       = 38
	MiniGoParserNOT                      = 39
	MiniGoParserBITWISEAND               = 40
	MiniGoParserBITWISEOR                = 41
	MiniGoParserBITWISEXOR               = 42
	MiniGoParserBITWISECLEAR             = 43
	MiniGoParserSHIFTLEFT                = 44
	MiniGoParserSHIFTRIGHT               = 45
	MiniGoParserINCREMENT                = 46
	MiniGoParserDECREMENT                = 47
	MiniGoParserASSIGN                   = 48
	MiniGoParserPLUSEQUAL                = 49
	MiniGoParserMINUSEQUAL               = 50
	MiniGoParserMULTIPLYEQUAL            = 51
	MiniGoParserDIVIDEEQUAL              = 52
	MiniGoParserMODULOEQUAL              = 53
	MiniGoParserBITWISEANDEQUAL          = 54
	MiniGoParserBITWISEOREQUAL           = 55
	MiniGoParserBITWISEXOREQUAL          = 56
	MiniGoParserSHIFTLEFTEQUAL           = 57
	MiniGoParserSHIFTRIGHTEQUAL          = 58
	MiniGoParserBITWISECLEAREQUAL        = 59
	MiniGoParserCOLON                    = 60
	MiniGoParserSEMICOLON                = 61
	MiniGoParserCOMMA                    = 62
	MiniGoParserDOT                      = 63
	MiniGoParserLPAREN                   = 64
	MiniGoParserRPAREN                   = 65
	MiniGoParserLBRACKET                 = 66
	MiniGoParserRBRACKET                 = 67
	MiniGoParserLBRACE                   = 68
	MiniGoParserRBRACE                   = 69
	MiniGoParserSPACES                   = 70
	MiniGoParserLINE_COMMENT             = 71
	MiniGoParserBLOCK_COMMENT            = 72
)

// MiniGoParser rules.
const (
	MiniGoParserRULE_root                     = 0
	MiniGoParserRULE_topDeclarationList       = 1
	MiniGoParserRULE_variableDecl             = 2
	MiniGoParserRULE_innerVarDecls            = 3
	MiniGoParserRULE_singleVarDecl            = 4
	MiniGoParserRULE_singleVarDeclNoExps      = 5
	MiniGoParserRULE_typeDecl                 = 6
	MiniGoParserRULE_innerTypeDecls           = 7
	MiniGoParserRULE_singleTypeDecl           = 8
	MiniGoParserRULE_funcDecl                 = 9
	MiniGoParserRULE_funcFrontDecl            = 10
	MiniGoParserRULE_multipleReturnTypes      = 11
	MiniGoParserRULE_returnTypeList           = 12
	MiniGoParserRULE_singleReturnType         = 13
	MiniGoParserRULE_funcArgDecls             = 14
	MiniGoParserRULE_declType                 = 15
	MiniGoParserRULE_sliceDeclType            = 16
	MiniGoParserRULE_arrayDeclType            = 17
	MiniGoParserRULE_structDeclType           = 18
	MiniGoParserRULE_structMemDecls           = 19
	MiniGoParserRULE_identifierList           = 20
	MiniGoParserRULE_expression               = 21
	MiniGoParserRULE_expressionList           = 22
	MiniGoParserRULE_primaryExpression        = 23
	MiniGoParserRULE_operand                  = 24
	MiniGoParserRULE_literal                  = 25
	MiniGoParserRULE_index                    = 26
	MiniGoParserRULE_arguments                = 27
	MiniGoParserRULE_selector                 = 28
	MiniGoParserRULE_appendExpression         = 29
	MiniGoParserRULE_lengthExpression         = 30
	MiniGoParserRULE_capExpression            = 31
	MiniGoParserRULE_statementList            = 32
	MiniGoParserRULE_block                    = 33
	MiniGoParserRULE_statement                = 34
	MiniGoParserRULE_simpleStatement          = 35
	MiniGoParserRULE_assignmentStatement      = 36
	MiniGoParserRULE_ifStatement              = 37
	MiniGoParserRULE_loop                     = 38
	MiniGoParserRULE_switchStmt               = 39
	MiniGoParserRULE_expressionCaseClauseList = 40
	MiniGoParserRULE_expressionCaseClause     = 41
	MiniGoParserRULE_expressionSwitchCase     = 42
	MiniGoParserRULE_epsilon                  = 43
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	TopDeclarationList() ITopDeclarationListContext

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
	p.RuleIndex = MiniGoParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserPACKAGE, 0)
}

func (s *RootContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIDENTIFIER, 0)
}

func (s *RootContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *RootContext) TopDeclarationList() ITopDeclarationListContext {
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

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MiniGoParserRULE_root)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(MiniGoParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Match(MiniGoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Match(MiniGoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
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

	// Getter signatures
	AllVariableDecl() []IVariableDeclContext
	VariableDecl(i int) IVariableDeclContext
	AllTypeDecl() []ITypeDeclContext
	TypeDecl(i int) ITypeDeclContext
	AllFuncDecl() []IFuncDeclContext
	FuncDecl(i int) IFuncDeclContext

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
	p.RuleIndex = MiniGoParserRULE_topDeclarationList
	return p
}

func InitEmptyTopDeclarationListContext(p *TopDeclarationListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_topDeclarationList
}

func (*TopDeclarationListContext) IsTopDeclarationListContext() {}

func NewTopDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_topDeclarationList

	return p
}

func (s *TopDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *TopDeclarationListContext) AllVariableDecl() []IVariableDeclContext {
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

func (s *TopDeclarationListContext) VariableDecl(i int) IVariableDeclContext {
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

func (s *TopDeclarationListContext) AllTypeDecl() []ITypeDeclContext {
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

func (s *TopDeclarationListContext) TypeDecl(i int) ITypeDeclContext {
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

func (s *TopDeclarationListContext) AllFuncDecl() []IFuncDeclContext {
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

func (s *TopDeclarationListContext) FuncDecl(i int) IFuncDeclContext {
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

func (s *TopDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopDeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterTopDeclarationList(s)
	}
}

func (s *TopDeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitTopDeclarationList(s)
	}
}

func (s *TopDeclarationListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitTopDeclarationList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) TopDeclarationList() (localctx ITopDeclarationListContext) {
	localctx = NewTopDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MiniGoParserRULE_topDeclarationList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0 {
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MiniGoParserVAR:
			{
				p.SetState(93)
				p.VariableDecl()
			}

		case MiniGoParserTYPE:
			{
				p.SetState(94)
				p.TypeDecl()
			}

		case MiniGoParserFUNC:
			{
				p.SetState(95)
				p.FuncDecl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(100)
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
	p.RuleIndex = MiniGoParserRULE_variableDecl
	return p
}

func InitEmptyVariableDeclContext(p *VariableDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_variableDecl
}

func (*VariableDeclContext) IsVariableDeclContext() {}

func NewVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclContext {
	var p = new(VariableDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_variableDecl

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
	return s.GetToken(MiniGoParserVAR, 0)
}

func (s *VariableDeclBlockASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLPAREN, 0)
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
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *VariableDeclBlockASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *VariableDeclBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterVariableDeclBlockAST(s)
	}
}

func (s *VariableDeclBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitVariableDeclBlockAST(s)
	}
}

func (s *VariableDeclBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserVAR, 0)
}

func (s *VariableDeclEmptyASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLPAREN, 0)
}

func (s *VariableDeclEmptyASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *VariableDeclEmptyASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *VariableDeclEmptyASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterVariableDeclEmptyAST(s)
	}
}

func (s *VariableDeclEmptyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitVariableDeclEmptyAST(s)
	}
}

func (s *VariableDeclEmptyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserVAR, 0)
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
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *VariableDeclASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterVariableDeclAST(s)
	}
}

func (s *VariableDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitVariableDeclAST(s)
	}
}

func (s *VariableDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitVariableDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) VariableDecl() (localctx IVariableDeclContext) {
	localctx = NewVariableDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MiniGoParserRULE_variableDecl)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVariableDeclASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Match(MiniGoParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.SingleVarDecl()
		}
		{
			p.SetState(103)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewVariableDeclBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.Match(MiniGoParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(MiniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.InnerVarDecls()
		}
		{
			p.SetState(108)
			p.Match(MiniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVariableDeclEmptyASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.Match(MiniGoParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(MiniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(MiniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(MiniGoParserSEMICOLON)
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

	// Getter signatures
	AllSingleVarDecl() []ISingleVarDeclContext
	SingleVarDecl(i int) ISingleVarDeclContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_innerVarDecls
	return p
}

func InitEmptyInnerVarDeclsContext(p *InnerVarDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_innerVarDecls
}

func (*InnerVarDeclsContext) IsInnerVarDeclsContext() {}

func NewInnerVarDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_innerVarDecls

	return p
}

func (s *InnerVarDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerVarDeclsContext) AllSingleVarDecl() []ISingleVarDeclContext {
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

func (s *InnerVarDeclsContext) SingleVarDecl(i int) ISingleVarDeclContext {
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

func (s *InnerVarDeclsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserSEMICOLON)
}

func (s *InnerVarDeclsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, i)
}

func (s *InnerVarDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerVarDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerVarDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterInnerVarDecls(s)
	}
}

func (s *InnerVarDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitInnerVarDecls(s)
	}
}

func (s *InnerVarDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitInnerVarDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) InnerVarDecls() (localctx IInnerVarDeclsContext) {
	localctx = NewInnerVarDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MiniGoParserRULE_innerVarDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.SingleVarDecl()
	}
	{
		p.SetState(118)
		p.Match(MiniGoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniGoParserIDENTIFIER {
		{
			p.SetState(119)
			p.SingleVarDecl()
		}
		{
			p.SetState(120)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(126)
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
	p.RuleIndex = MiniGoParserRULE_singleVarDecl
	return p
}

func InitEmptySingleVarDeclContext(p *SingleVarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_singleVarDecl
}

func (*SingleVarDeclContext) IsSingleVarDeclContext() {}

func NewSingleVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_singleVarDecl

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
	return s.GetToken(MiniGoParserASSIGN, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSingleVarDeclAssignAST(s)
	}
}

func (s *SingleVarDeclAssignASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSingleVarDeclAssignAST(s)
	}
}

func (s *SingleVarDeclAssignASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserASSIGN, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSingleVarDeclAST(s)
	}
}

func (s *SingleVarDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSingleVarDeclAST(s)
	}
}

func (s *SingleVarDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSingleVarDeclNoExpsAST(s)
	}
}

func (s *SingleVarDeclNoExpsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSingleVarDeclNoExpsAST(s)
	}
}

func (s *SingleVarDeclNoExpsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSingleVarDeclNoExpsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) SingleVarDecl() (localctx ISingleVarDeclContext) {
	localctx = NewSingleVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MiniGoParserRULE_singleVarDecl)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleVarDeclASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.IdentifierList()
		}
		{
			p.SetState(128)
			p.DeclType()
		}
		{
			p.SetState(129)
			p.Match(MiniGoParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.ExpressionList()
		}

	case 2:
		localctx = NewSingleVarDeclAssignASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.IdentifierList()
		}
		{
			p.SetState(133)
			p.Match(MiniGoParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.ExpressionList()
		}

	case 3:
		localctx = NewSingleVarDeclNoExpsASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(136)
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

	// Getter signatures
	IdentifierList() IIdentifierListContext
	DeclType() IDeclTypeContext

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
	p.RuleIndex = MiniGoParserRULE_singleVarDeclNoExps
	return p
}

func InitEmptySingleVarDeclNoExpsContext(p *SingleVarDeclNoExpsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_singleVarDeclNoExps
}

func (*SingleVarDeclNoExpsContext) IsSingleVarDeclNoExpsContext() {}

func NewSingleVarDeclNoExpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_singleVarDeclNoExps

	return p
}

func (s *SingleVarDeclNoExpsContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclNoExpsContext) IdentifierList() IIdentifierListContext {
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

func (s *SingleVarDeclNoExpsContext) DeclType() IDeclTypeContext {
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

func (s *SingleVarDeclNoExpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclNoExpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleVarDeclNoExpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSingleVarDeclNoExps(s)
	}
}

func (s *SingleVarDeclNoExpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSingleVarDeclNoExps(s)
	}
}

func (s *SingleVarDeclNoExpsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSingleVarDeclNoExps(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) SingleVarDeclNoExps() (localctx ISingleVarDeclNoExpsContext) {
	localctx = NewSingleVarDeclNoExpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MiniGoParserRULE_singleVarDeclNoExps)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.IdentifierList()
	}
	{
		p.SetState(140)
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
	p.RuleIndex = MiniGoParserRULE_typeDecl
	return p
}

func InitEmptyTypeDeclContext(p *TypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_typeDecl
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_typeDecl

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
	return s.GetToken(MiniGoParserTYPE, 0)
}

func (s *TypeDeclEmptyASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLPAREN, 0)
}

func (s *TypeDeclEmptyASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *TypeDeclEmptyASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *TypeDeclEmptyASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterTypeDeclEmptyAST(s)
	}
}

func (s *TypeDeclEmptyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitTypeDeclEmptyAST(s)
	}
}

func (s *TypeDeclEmptyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserTYPE, 0)
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
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *TypeDeclASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterTypeDeclAST(s)
	}
}

func (s *TypeDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitTypeDeclAST(s)
	}
}

func (s *TypeDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserTYPE, 0)
}

func (s *TypeDeclBlockASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLPAREN, 0)
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
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *TypeDeclBlockASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *TypeDeclBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterTypeDeclBlockAST(s)
	}
}

func (s *TypeDeclBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitTypeDeclBlockAST(s)
	}
}

func (s *TypeDeclBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitTypeDeclBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MiniGoParserRULE_typeDecl)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeDeclASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Match(MiniGoParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.SingleTypeDecl()
		}
		{
			p.SetState(144)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewTypeDeclBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Match(MiniGoParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(MiniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.InnerTypeDecls()
		}
		{
			p.SetState(149)
			p.Match(MiniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewTypeDeclEmptyASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(152)
			p.Match(MiniGoParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(MiniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(MiniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(MiniGoParserSEMICOLON)
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

	// Getter signatures
	AllSingleTypeDecl() []ISingleTypeDeclContext
	SingleTypeDecl(i int) ISingleTypeDeclContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_innerTypeDecls
	return p
}

func InitEmptyInnerTypeDeclsContext(p *InnerTypeDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_innerTypeDecls
}

func (*InnerTypeDeclsContext) IsInnerTypeDeclsContext() {}

func NewInnerTypeDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_innerTypeDecls

	return p
}

func (s *InnerTypeDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerTypeDeclsContext) AllSingleTypeDecl() []ISingleTypeDeclContext {
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

func (s *InnerTypeDeclsContext) SingleTypeDecl(i int) ISingleTypeDeclContext {
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

func (s *InnerTypeDeclsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserSEMICOLON)
}

func (s *InnerTypeDeclsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, i)
}

func (s *InnerTypeDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerTypeDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterInnerTypeDecls(s)
	}
}

func (s *InnerTypeDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitInnerTypeDecls(s)
	}
}

func (s *InnerTypeDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitInnerTypeDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) InnerTypeDecls() (localctx IInnerTypeDeclsContext) {
	localctx = NewInnerTypeDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MiniGoParserRULE_innerTypeDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.SingleTypeDecl()
	}
	{
		p.SetState(159)
		p.Match(MiniGoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniGoParserIDENTIFIER {
		{
			p.SetState(160)
			p.SingleTypeDecl()
		}
		{
			p.SetState(161)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(167)
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

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	DeclType() IDeclTypeContext

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
	p.RuleIndex = MiniGoParserRULE_singleTypeDecl
	return p
}

func InitEmptySingleTypeDeclContext(p *SingleTypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_singleTypeDecl
}

func (*SingleTypeDeclContext) IsSingleTypeDeclContext() {}

func NewSingleTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_singleTypeDecl

	return p
}

func (s *SingleTypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleTypeDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIDENTIFIER, 0)
}

func (s *SingleTypeDeclContext) DeclType() IDeclTypeContext {
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

func (s *SingleTypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleTypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleTypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSingleTypeDecl(s)
	}
}

func (s *SingleTypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSingleTypeDecl(s)
	}
}

func (s *SingleTypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSingleTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) SingleTypeDecl() (localctx ISingleTypeDeclContext) {
	localctx = NewSingleTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MiniGoParserRULE_singleTypeDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(MiniGoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
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

	// Getter signatures
	FuncFrontDecl() IFuncFrontDeclContext
	Block() IBlockContext
	SEMICOLON() antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_funcDecl
	return p
}

func InitEmptyFuncDeclContext(p *FuncDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_funcDecl
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) FuncFrontDecl() IFuncFrontDeclContext {
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

func (s *FuncDeclContext) Block() IBlockContext {
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

func (s *FuncDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MiniGoParserRULE_funcDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.FuncFrontDecl()
	}
	{
		p.SetState(172)
		p.Block()
	}
	{
		p.SetState(173)
		p.Match(MiniGoParserSEMICOLON)
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

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	FuncArgDecls() IFuncArgDeclsContext
	Epsilon() IEpsilonContext
	MultipleReturnTypes() IMultipleReturnTypesContext
	SingleReturnType() ISingleReturnTypeContext

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
	p.RuleIndex = MiniGoParserRULE_funcFrontDecl
	return p
}

func InitEmptyFuncFrontDeclContext(p *FuncFrontDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_funcFrontDecl
}

func (*FuncFrontDeclContext) IsFuncFrontDeclContext() {}

func NewFuncFrontDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_funcFrontDecl

	return p
}

func (s *FuncFrontDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncFrontDeclContext) FUNC() antlr.TerminalNode {
	return s.GetToken(MiniGoParserFUNC, 0)
}

func (s *FuncFrontDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIDENTIFIER, 0)
}

func (s *FuncFrontDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLPAREN, 0)
}

func (s *FuncFrontDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *FuncFrontDeclContext) FuncArgDecls() IFuncArgDeclsContext {
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

func (s *FuncFrontDeclContext) Epsilon() IEpsilonContext {
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

func (s *FuncFrontDeclContext) MultipleReturnTypes() IMultipleReturnTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultipleReturnTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultipleReturnTypesContext)
}

func (s *FuncFrontDeclContext) SingleReturnType() ISingleReturnTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleReturnTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleReturnTypeContext)
}

func (s *FuncFrontDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFrontDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncFrontDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterFuncFrontDecl(s)
	}
}

func (s *FuncFrontDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitFuncFrontDecl(s)
	}
}

func (s *FuncFrontDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitFuncFrontDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) FuncFrontDecl() (localctx IFuncFrontDeclContext) {
	localctx = NewFuncFrontDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MiniGoParserRULE_funcFrontDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(MiniGoParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(MiniGoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(MiniGoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserIDENTIFIER:
		{
			p.SetState(178)
			p.FuncArgDecls()
		}

	case MiniGoParserRPAREN:
		{
			p.SetState(179)
			p.Epsilon()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(182)
		p.Match(MiniGoParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(183)
			p.MultipleReturnTypes()
		}

	case 2:
		{
			p.SetState(184)
			p.SingleReturnType()
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

// IMultipleReturnTypesContext is an interface to support dynamic dispatch.
type IMultipleReturnTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	ReturnTypeList() IReturnTypeListContext
	RPAREN() antlr.TerminalNode

	// IsMultipleReturnTypesContext differentiates from other interfaces.
	IsMultipleReturnTypesContext()
}

type MultipleReturnTypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultipleReturnTypesContext() *MultipleReturnTypesContext {
	var p = new(MultipleReturnTypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_multipleReturnTypes
	return p
}

func InitEmptyMultipleReturnTypesContext(p *MultipleReturnTypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_multipleReturnTypes
}

func (*MultipleReturnTypesContext) IsMultipleReturnTypesContext() {}

func NewMultipleReturnTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultipleReturnTypesContext {
	var p = new(MultipleReturnTypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_multipleReturnTypes

	return p
}

func (s *MultipleReturnTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *MultipleReturnTypesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLPAREN, 0)
}

func (s *MultipleReturnTypesContext) ReturnTypeList() IReturnTypeListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnTypeListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnTypeListContext)
}

func (s *MultipleReturnTypesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *MultipleReturnTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultipleReturnTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultipleReturnTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterMultipleReturnTypes(s)
	}
}

func (s *MultipleReturnTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitMultipleReturnTypes(s)
	}
}

func (s *MultipleReturnTypesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitMultipleReturnTypes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) MultipleReturnTypes() (localctx IMultipleReturnTypesContext) {
	localctx = NewMultipleReturnTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MiniGoParserRULE_multipleReturnTypes)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(MiniGoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.ReturnTypeList()
	}
	{
		p.SetState(189)
		p.Match(MiniGoParserRPAREN)
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

// IReturnTypeListContext is an interface to support dynamic dispatch.
type IReturnTypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclType() []IDeclTypeContext
	DeclType(i int) IDeclTypeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsReturnTypeListContext differentiates from other interfaces.
	IsReturnTypeListContext()
}

type ReturnTypeListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnTypeListContext() *ReturnTypeListContext {
	var p = new(ReturnTypeListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_returnTypeList
	return p
}

func InitEmptyReturnTypeListContext(p *ReturnTypeListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_returnTypeList
}

func (*ReturnTypeListContext) IsReturnTypeListContext() {}

func NewReturnTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeListContext {
	var p = new(ReturnTypeListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_returnTypeList

	return p
}

func (s *ReturnTypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnTypeListContext) AllDeclType() []IDeclTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclTypeContext); ok {
			len++
		}
	}

	tst := make([]IDeclTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclTypeContext); ok {
			tst[i] = t.(IDeclTypeContext)
			i++
		}
	}

	return tst
}

func (s *ReturnTypeListContext) DeclType(i int) IDeclTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
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

	return t.(IDeclTypeContext)
}

func (s *ReturnTypeListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserCOMMA)
}

func (s *ReturnTypeListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserCOMMA, i)
}

func (s *ReturnTypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnTypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterReturnTypeList(s)
	}
}

func (s *ReturnTypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitReturnTypeList(s)
	}
}

func (s *ReturnTypeListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitReturnTypeList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) ReturnTypeList() (localctx IReturnTypeListContext) {
	localctx = NewReturnTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MiniGoParserRULE_returnTypeList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.DeclType()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniGoParserCOMMA {
		{
			p.SetState(192)
			p.Match(MiniGoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.DeclType()
		}

		p.SetState(198)
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

// ISingleReturnTypeContext is an interface to support dynamic dispatch.
type ISingleReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingleReturnTypeContext differentiates from other interfaces.
	IsSingleReturnTypeContext()
}

type SingleReturnTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleReturnTypeContext() *SingleReturnTypeContext {
	var p = new(SingleReturnTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_singleReturnType
	return p
}

func InitEmptySingleReturnTypeContext(p *SingleReturnTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_singleReturnType
}

func (*SingleReturnTypeContext) IsSingleReturnTypeContext() {}

func NewSingleReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleReturnTypeContext {
	var p = new(SingleReturnTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_singleReturnType

	return p
}

func (s *SingleReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleReturnTypeContext) CopyAll(ctx *SingleReturnTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleReturnTypeEmptyASTContext struct {
	SingleReturnTypeContext
}

func NewSingleReturnTypeEmptyASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleReturnTypeEmptyASTContext {
	var p = new(SingleReturnTypeEmptyASTContext)

	InitEmptySingleReturnTypeContext(&p.SingleReturnTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleReturnTypeContext))

	return p
}

func (s *SingleReturnTypeEmptyASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleReturnTypeEmptyASTContext) Epsilon() IEpsilonContext {
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

func (s *SingleReturnTypeEmptyASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSingleReturnTypeEmptyAST(s)
	}
}

func (s *SingleReturnTypeEmptyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSingleReturnTypeEmptyAST(s)
	}
}

func (s *SingleReturnTypeEmptyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSingleReturnTypeEmptyAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleReturnTypeASTContext struct {
	SingleReturnTypeContext
}

func NewSingleReturnTypeASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleReturnTypeASTContext {
	var p = new(SingleReturnTypeASTContext)

	InitEmptySingleReturnTypeContext(&p.SingleReturnTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleReturnTypeContext))

	return p
}

func (s *SingleReturnTypeASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleReturnTypeASTContext) DeclType() IDeclTypeContext {
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

func (s *SingleReturnTypeASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSingleReturnTypeAST(s)
	}
}

func (s *SingleReturnTypeASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSingleReturnTypeAST(s)
	}
}

func (s *SingleReturnTypeASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSingleReturnTypeAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) SingleReturnType() (localctx ISingleReturnTypeContext) {
	localctx = NewSingleReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MiniGoParserRULE_singleReturnType)
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserSTRUCT, MiniGoParserIDENTIFIER, MiniGoParserLPAREN, MiniGoParserLBRACKET:
		localctx = NewSingleReturnTypeASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.DeclType()
		}

	case MiniGoParserLBRACE:
		localctx = NewSingleReturnTypeEmptyASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(200)
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

	// Getter signatures
	AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext
	SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_funcArgDecls
	return p
}

func InitEmptyFuncArgDeclsContext(p *FuncArgDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_funcArgDecls
}

func (*FuncArgDeclsContext) IsFuncArgDeclsContext() {}

func NewFuncArgDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgDeclsContext {
	var p = new(FuncArgDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_funcArgDecls

	return p
}

func (s *FuncArgDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgDeclsContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
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

func (s *FuncArgDeclsContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
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

func (s *FuncArgDeclsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserCOMMA)
}

func (s *FuncArgDeclsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserCOMMA, i)
}

func (s *FuncArgDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterFuncArgDecls(s)
	}
}

func (s *FuncArgDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitFuncArgDecls(s)
	}
}

func (s *FuncArgDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitFuncArgDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) FuncArgDecls() (localctx IFuncArgDeclsContext) {
	localctx = NewFuncArgDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MiniGoParserRULE_funcArgDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.SingleVarDeclNoExps()
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniGoParserCOMMA {
		{
			p.SetState(204)
			p.Match(MiniGoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.SingleVarDeclNoExps()
		}

		p.SetState(210)
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
	p.RuleIndex = MiniGoParserRULE_declType
	return p
}

func InitEmptyDeclTypeContext(p *DeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_declType
}

func (*DeclTypeContext) IsDeclTypeContext() {}

func NewDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclTypeContext {
	var p = new(DeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_declType

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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterDeclTypeStructAST(s)
	}
}

func (s *DeclTypeStructASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitDeclTypeStructAST(s)
	}
}

func (s *DeclTypeStructASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterDeclTypeSliceAST(s)
	}
}

func (s *DeclTypeSliceASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitDeclTypeSliceAST(s)
	}
}

func (s *DeclTypeSliceASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterDeclTypeArrayAST(s)
	}
}

func (s *DeclTypeArrayASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitDeclTypeArrayAST(s)
	}
}

func (s *DeclTypeArrayASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserLPAREN, 0)
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
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *DeclTypeParenASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterDeclTypeParenAST(s)
	}
}

func (s *DeclTypeParenASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitDeclTypeParenAST(s)
	}
}

func (s *DeclTypeParenASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserIDENTIFIER, 0)
}

func (s *DeclTypeIdentifierASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterDeclTypeIdentifierAST(s)
	}
}

func (s *DeclTypeIdentifierASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitDeclTypeIdentifierAST(s)
	}
}

func (s *DeclTypeIdentifierASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitDeclTypeIdentifierAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) DeclType() (localctx IDeclTypeContext) {
	localctx = NewDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MiniGoParserRULE_declType)
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclTypeParenASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Match(MiniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.DeclType()
		}
		{
			p.SetState(213)
			p.Match(MiniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDeclTypeIdentifierASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Match(MiniGoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewDeclTypeSliceASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.SliceDeclType()
		}

	case 4:
		localctx = NewDeclTypeArrayASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(217)
			p.ArrayDeclType()
		}

	case 5:
		localctx = NewDeclTypeStructASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(218)
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

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	DeclType() IDeclTypeContext

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
	p.RuleIndex = MiniGoParserRULE_sliceDeclType
	return p
}

func InitEmptySliceDeclTypeContext(p *SliceDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_sliceDeclType
}

func (*SliceDeclTypeContext) IsSliceDeclTypeContext() {}

func NewSliceDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_sliceDeclType

	return p
}

func (s *SliceDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceDeclTypeContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLBRACKET, 0)
}

func (s *SliceDeclTypeContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRBRACKET, 0)
}

func (s *SliceDeclTypeContext) DeclType() IDeclTypeContext {
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

func (s *SliceDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceDeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSliceDeclType(s)
	}
}

func (s *SliceDeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSliceDeclType(s)
	}
}

func (s *SliceDeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSliceDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) SliceDeclType() (localctx ISliceDeclTypeContext) {
	localctx = NewSliceDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MiniGoParserRULE_sliceDeclType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(MiniGoParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(MiniGoParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
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

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	INTLITERAL() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	DeclType() IDeclTypeContext

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
	p.RuleIndex = MiniGoParserRULE_arrayDeclType
	return p
}

func InitEmptyArrayDeclTypeContext(p *ArrayDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_arrayDeclType
}

func (*ArrayDeclTypeContext) IsArrayDeclTypeContext() {}

func NewArrayDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_arrayDeclType

	return p
}

func (s *ArrayDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDeclTypeContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLBRACKET, 0)
}

func (s *ArrayDeclTypeContext) INTLITERAL() antlr.TerminalNode {
	return s.GetToken(MiniGoParserINTLITERAL, 0)
}

func (s *ArrayDeclTypeContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRBRACKET, 0)
}

func (s *ArrayDeclTypeContext) DeclType() IDeclTypeContext {
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

func (s *ArrayDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterArrayDeclType(s)
	}
}

func (s *ArrayDeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitArrayDeclType(s)
	}
}

func (s *ArrayDeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitArrayDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) ArrayDeclType() (localctx IArrayDeclTypeContext) {
	localctx = NewArrayDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MiniGoParserRULE_arrayDeclType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(MiniGoParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.Match(MiniGoParserINTLITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
		p.Match(MiniGoParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
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

	// Getter signatures
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	StructMemDecls() IStructMemDeclsContext
	Epsilon() IEpsilonContext

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
	p.RuleIndex = MiniGoParserRULE_structDeclType
	return p
}

func InitEmptyStructDeclTypeContext(p *StructDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_structDeclType
}

func (*StructDeclTypeContext) IsStructDeclTypeContext() {}

func NewStructDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_structDeclType

	return p
}

func (s *StructDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclTypeContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSTRUCT, 0)
}

func (s *StructDeclTypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLBRACE, 0)
}

func (s *StructDeclTypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRBRACE, 0)
}

func (s *StructDeclTypeContext) StructMemDecls() IStructMemDeclsContext {
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

func (s *StructDeclTypeContext) Epsilon() IEpsilonContext {
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

func (s *StructDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStructDeclType(s)
	}
}

func (s *StructDeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStructDeclType(s)
	}
}

func (s *StructDeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitStructDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) StructDeclType() (localctx IStructDeclTypeContext) {
	localctx = NewStructDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MiniGoParserRULE_structDeclType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(MiniGoParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.Match(MiniGoParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserIDENTIFIER:
		{
			p.SetState(232)
			p.StructMemDecls()
		}

	case MiniGoParserRBRACE:
		{
			p.SetState(233)
			p.Epsilon()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(236)
		p.Match(MiniGoParserRBRACE)
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

	// Getter signatures
	AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext
	SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_structMemDecls
	return p
}

func InitEmptyStructMemDeclsContext(p *StructMemDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_structMemDecls
}

func (*StructMemDeclsContext) IsStructMemDeclsContext() {}

func NewStructMemDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_structMemDecls

	return p
}

func (s *StructMemDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *StructMemDeclsContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
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

func (s *StructMemDeclsContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
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

func (s *StructMemDeclsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserSEMICOLON)
}

func (s *StructMemDeclsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, i)
}

func (s *StructMemDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructMemDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStructMemDecls(s)
	}
}

func (s *StructMemDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStructMemDecls(s)
	}
}

func (s *StructMemDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitStructMemDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) StructMemDecls() (localctx IStructMemDeclsContext) {
	localctx = NewStructMemDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MiniGoParserRULE_structMemDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.SingleVarDeclNoExps()
	}
	{
		p.SetState(239)
		p.Match(MiniGoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniGoParserIDENTIFIER {
		{
			p.SetState(240)
			p.SingleVarDeclNoExps()
		}
		{
			p.SetState(241)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(247)
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

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_identifierList
	return p
}

func InitEmptyIdentifierListContext(p *IdentifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_identifierList
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserIDENTIFIER)
}

func (s *IdentifierListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserIDENTIFIER, i)
}

func (s *IdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserCOMMA)
}

func (s *IdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserCOMMA, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MiniGoParserRULE_identifierList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(MiniGoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniGoParserCOMMA {
		{
			p.SetState(249)
			p.Match(MiniGoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Match(MiniGoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(255)
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
	p.RuleIndex = MiniGoParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_expression

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
	return s.GetToken(MiniGoParserNOT, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionNotUnaryAST(s)
	}
}

func (s *ExpressionNotUnaryASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionNotUnaryAST(s)
	}
}

func (s *ExpressionNotUnaryASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserMULTIPLY, 0)
}

func (s *ExpressionMultiplyASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionMultiplyAST(s)
	}
}

func (s *ExpressionMultiplyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionMultiplyAST(s)
	}
}

func (s *ExpressionMultiplyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserPLUS, 0)
}

func (s *ExpressionPlusASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionPlusAST(s)
	}
}

func (s *ExpressionPlusASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionPlusAST(s)
	}
}

func (s *ExpressionPlusASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserMODULO, 0)
}

func (s *ExpressionModuloASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionModuloAST(s)
	}
}

func (s *ExpressionModuloASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionModuloAST(s)
	}
}

func (s *ExpressionModuloASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserAND, 0)
}

func (s *ExpressionAndASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionAndAST(s)
	}
}

func (s *ExpressionAndASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionAndAST(s)
	}
}

func (s *ExpressionAndASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserBITWISEXOR, 0)
}

func (s *ExpressionBitwiseXorASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionBitwiseXorAST(s)
	}
}

func (s *ExpressionBitwiseXorASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionBitwiseXorAST(s)
	}
}

func (s *ExpressionBitwiseXorASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserMINUS, 0)
}

func (s *ExpressionMinusASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionMinusAST(s)
	}
}

func (s *ExpressionMinusASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionMinusAST(s)
	}
}

func (s *ExpressionMinusASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserBITWISEXOR, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionBitwiseXorUnaryAST(s)
	}
}

func (s *ExpressionBitwiseXorUnaryASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionBitwiseXorUnaryAST(s)
	}
}

func (s *ExpressionBitwiseXorUnaryASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserEQUAL, 0)
}

func (s *ExpressionEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionEqualAST(s)
	}
}

func (s *ExpressionEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionEqualAST(s)
	}
}

func (s *ExpressionEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserMINUS, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionMinusUnaryAST(s)
	}
}

func (s *ExpressionMinusUnaryASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionMinusUnaryAST(s)
	}
}

func (s *ExpressionMinusUnaryASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserBITWISECLEAR, 0)
}

func (s *ExpressionBitwiseClearASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionBitwiseClearAST(s)
	}
}

func (s *ExpressionBitwiseClearASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionBitwiseClearAST(s)
	}
}

func (s *ExpressionBitwiseClearASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserGREATEREQUAL, 0)
}

func (s *ExpressionGreaterEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionGreaterEqualAST(s)
	}
}

func (s *ExpressionGreaterEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionGreaterEqualAST(s)
	}
}

func (s *ExpressionGreaterEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserLESSEQUAL, 0)
}

func (s *ExpressionLessEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionLessEqualAST(s)
	}
}

func (s *ExpressionLessEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionLessEqualAST(s)
	}
}

func (s *ExpressionLessEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserNOTEQUAL, 0)
}

func (s *ExpressionNotEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionNotEqualAST(s)
	}
}

func (s *ExpressionNotEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionNotEqualAST(s)
	}
}

func (s *ExpressionNotEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionPrimaryAST(s)
	}
}

func (s *ExpressionPrimaryASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionPrimaryAST(s)
	}
}

func (s *ExpressionPrimaryASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserBITWISEAND, 0)
}

func (s *ExpressionBitwiseAndASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionBitwiseAndAST(s)
	}
}

func (s *ExpressionBitwiseAndASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionBitwiseAndAST(s)
	}
}

func (s *ExpressionBitwiseAndASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserGREATER, 0)
}

func (s *ExpressionGreaterASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionGreaterAST(s)
	}
}

func (s *ExpressionGreaterASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionGreaterAST(s)
	}
}

func (s *ExpressionGreaterASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserDIVIDE, 0)
}

func (s *ExpressionDivideASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionDivideAST(s)
	}
}

func (s *ExpressionDivideASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionDivideAST(s)
	}
}

func (s *ExpressionDivideASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserOR, 0)
}

func (s *ExpressionOrASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionOrAST(s)
	}
}

func (s *ExpressionOrASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionOrAST(s)
	}
}

func (s *ExpressionOrASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserPLUS, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionPlusUnaryAST(s)
	}
}

func (s *ExpressionPlusUnaryASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionPlusUnaryAST(s)
	}
}

func (s *ExpressionPlusUnaryASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserBITWISEOR, 0)
}

func (s *ExpressionBitwiseOrASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionBitwiseOrAST(s)
	}
}

func (s *ExpressionBitwiseOrASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionBitwiseOrAST(s)
	}
}

func (s *ExpressionBitwiseOrASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserLESS, 0)
}

func (s *ExpressionLessASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionLessAST(s)
	}
}

func (s *ExpressionLessASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionLessAST(s)
	}
}

func (s *ExpressionLessASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserSHIFTRIGHT, 0)
}

func (s *ExpressionShiftRightASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionShiftRightAST(s)
	}
}

func (s *ExpressionShiftRightASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionShiftRightAST(s)
	}
}

func (s *ExpressionShiftRightASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserSHIFTLEFT, 0)
}

func (s *ExpressionShiftLeftASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionShiftLeftAST(s)
	}
}

func (s *ExpressionShiftLeftASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionShiftLeftAST(s)
	}
}

func (s *ExpressionShiftLeftASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExpressionShiftLeftAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *MiniGoParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, MiniGoParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserINTLITERAL, MiniGoParserFLOATLITERAL, MiniGoParserRUNELITERAL, MiniGoParserRAWSTRINGLITERAL, MiniGoParserINTERPRETEDSTRINGLITERAL, MiniGoParserLEN, MiniGoParserCAP, MiniGoParserAPPEND, MiniGoParserIDENTIFIER, MiniGoParserLPAREN:
		localctx = NewExpressionPrimaryASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(257)
			p.primaryExpression(0)
		}

	case MiniGoParserPLUS:
		localctx = NewExpressionPlusUnaryASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(258)
			p.Match(MiniGoParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.expression(4)
		}

	case MiniGoParserMINUS:
		localctx = NewExpressionMinusUnaryASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(260)
			p.Match(MiniGoParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.expression(3)
		}

	case MiniGoParserNOT:
		localctx = NewExpressionNotUnaryASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(262)
			p.Match(MiniGoParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
			p.expression(2)
		}

	case MiniGoParserBITWISEXOR:
		localctx = NewExpressionBitwiseXorUnaryASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(264)
			p.Match(MiniGoParserBITWISEXOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.expression(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(325)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionMultiplyASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(268)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(269)
					p.Match(MiniGoParserMULTIPLY)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(270)
					p.expression(24)
				}

			case 2:
				localctx = NewExpressionDivideASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(272)
					p.Match(MiniGoParserDIVIDE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(273)
					p.expression(23)
				}

			case 3:
				localctx = NewExpressionModuloASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(275)
					p.Match(MiniGoParserMODULO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(276)
					p.expression(22)
				}

			case 4:
				localctx = NewExpressionShiftLeftASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(278)
					p.Match(MiniGoParserSHIFTLEFT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(279)
					p.expression(21)
				}

			case 5:
				localctx = NewExpressionShiftRightASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(281)
					p.Match(MiniGoParserSHIFTRIGHT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(282)
					p.expression(20)
				}

			case 6:
				localctx = NewExpressionBitwiseAndASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(283)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(284)
					p.Match(MiniGoParserBITWISEAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(285)
					p.expression(19)
				}

			case 7:
				localctx = NewExpressionBitwiseClearASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(287)
					p.Match(MiniGoParserBITWISECLEAR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(288)
					p.expression(18)
				}

			case 8:
				localctx = NewExpressionPlusASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(290)
					p.Match(MiniGoParserPLUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(291)
					p.expression(17)
				}

			case 9:
				localctx = NewExpressionMinusASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(293)
					p.Match(MiniGoParserMINUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(294)
					p.expression(16)
				}

			case 10:
				localctx = NewExpressionBitwiseOrASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(296)
					p.Match(MiniGoParserBITWISEOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(297)
					p.expression(15)
				}

			case 11:
				localctx = NewExpressionBitwiseXorASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(298)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(299)
					p.Match(MiniGoParserBITWISEXOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(300)
					p.expression(14)
				}

			case 12:
				localctx = NewExpressionEqualASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(301)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(302)
					p.Match(MiniGoParserEQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(303)
					p.expression(13)
				}

			case 13:
				localctx = NewExpressionNotEqualASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(304)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(305)
					p.Match(MiniGoParserNOTEQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(306)
					p.expression(12)
				}

			case 14:
				localctx = NewExpressionLessASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(307)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(308)
					p.Match(MiniGoParserLESS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(309)
					p.expression(11)
				}

			case 15:
				localctx = NewExpressionLessEqualASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(311)
					p.Match(MiniGoParserLESSEQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(312)
					p.expression(10)
				}

			case 16:
				localctx = NewExpressionGreaterASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(313)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(314)
					p.Match(MiniGoParserGREATER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(315)
					p.expression(9)
				}

			case 17:
				localctx = NewExpressionGreaterEqualASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(316)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(317)
					p.Match(MiniGoParserGREATEREQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(318)
					p.expression(8)
				}

			case 18:
				localctx = NewExpressionAndASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(319)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(320)
					p.Match(MiniGoParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(321)
					p.expression(7)
				}

			case 19:
				localctx = NewExpressionOrASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expression)
				p.SetState(322)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(323)
					p.Match(MiniGoParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(324)
					p.expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
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

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MiniGoParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MiniGoParserRULE_expressionList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.expression(0)
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(331)
				p.Match(MiniGoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(332)
				p.expression(0)
			}

		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
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
	p.RuleIndex = MiniGoParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_primaryExpression

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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionLengthAST(s)
	}
}

func (s *PrimaryExpressionLengthASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionLengthAST(s)
	}
}

func (s *PrimaryExpressionLengthASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionOperandAST(s)
	}
}

func (s *PrimaryExpressionOperandASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionOperandAST(s)
	}
}

func (s *PrimaryExpressionOperandASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionIndexAST(s)
	}
}

func (s *PrimaryExpressionIndexASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionIndexAST(s)
	}
}

func (s *PrimaryExpressionIndexASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionAppendAST(s)
	}
}

func (s *PrimaryExpressionAppendASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionAppendAST(s)
	}
}

func (s *PrimaryExpressionAppendASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionArgumentsAST(s)
	}
}

func (s *PrimaryExpressionArgumentsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionArgumentsAST(s)
	}
}

func (s *PrimaryExpressionArgumentsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionCapAST(s)
	}
}

func (s *PrimaryExpressionCapASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionCapAST(s)
	}
}

func (s *PrimaryExpressionCapASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterPrimaryExpressionSelectorAST(s)
	}
}

func (s *PrimaryExpressionSelectorASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitPrimaryExpressionSelectorAST(s)
	}
}

func (s *PrimaryExpressionSelectorASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitPrimaryExpressionSelectorAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	return p.primaryExpression(0)
}

func (p *MiniGoParser) primaryExpression(_p int) (localctx IPrimaryExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, MiniGoParserRULE_primaryExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserINTLITERAL, MiniGoParserFLOATLITERAL, MiniGoParserRUNELITERAL, MiniGoParserRAWSTRINGLITERAL, MiniGoParserINTERPRETEDSTRINGLITERAL, MiniGoParserIDENTIFIER, MiniGoParserLPAREN:
		localctx = NewPrimaryExpressionOperandASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(339)
			p.Operand()
		}

	case MiniGoParserAPPEND:
		localctx = NewPrimaryExpressionAppendASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(340)
			p.AppendExpression()
		}

	case MiniGoParserLEN:
		localctx = NewPrimaryExpressionLengthASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(341)
			p.LengthExpression()
		}

	case MiniGoParserCAP:
		localctx = NewPrimaryExpressionCapASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(342)
			p.CapExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(351)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimaryExpressionSelectorASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_primaryExpression)
				p.SetState(345)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(346)
					p.Selector()
				}

			case 2:
				localctx = NewPrimaryExpressionIndexASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_primaryExpression)
				p.SetState(347)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(348)
					p.Index()
				}

			case 3:
				localctx = NewPrimaryExpressionArgumentsASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_primaryExpression)
				p.SetState(349)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(350)
					p.Arguments()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
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
	p.RuleIndex = MiniGoParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_operand

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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterOperandLiteralAST(s)
	}
}

func (s *OperandLiteralASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitOperandLiteralAST(s)
	}
}

func (s *OperandLiteralASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserIDENTIFIER, 0)
}

func (s *OperandIdentifierASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterOperandIdentifierAST(s)
	}
}

func (s *OperandIdentifierASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitOperandIdentifierAST(s)
	}
}

func (s *OperandIdentifierASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserLPAREN, 0)
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
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *OperandParenASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterOperandParenAST(s)
	}
}

func (s *OperandParenASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitOperandParenAST(s)
	}
}

func (s *OperandParenASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitOperandParenAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MiniGoParserRULE_operand)
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserINTLITERAL, MiniGoParserFLOATLITERAL, MiniGoParserRUNELITERAL, MiniGoParserRAWSTRINGLITERAL, MiniGoParserINTERPRETEDSTRINGLITERAL:
		localctx = NewOperandLiteralASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(356)
			p.Literal()
		}

	case MiniGoParserIDENTIFIER:
		localctx = NewOperandIdentifierASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.Match(MiniGoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserLPAREN:
		localctx = NewOperandParenASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(358)
			p.Match(MiniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)
			p.expression(0)
		}
		{
			p.SetState(360)
			p.Match(MiniGoParserRPAREN)
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
	p.RuleIndex = MiniGoParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_literal

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
	return s.GetToken(MiniGoParserINTLITERAL, 0)
}

func (s *LiteralIntASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterLiteralIntAST(s)
	}
}

func (s *LiteralIntASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitLiteralIntAST(s)
	}
}

func (s *LiteralIntASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserINTERPRETEDSTRINGLITERAL, 0)
}

func (s *LiteralInterpretedStringASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterLiteralInterpretedStringAST(s)
	}
}

func (s *LiteralInterpretedStringASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitLiteralInterpretedStringAST(s)
	}
}

func (s *LiteralInterpretedStringASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserFLOATLITERAL, 0)
}

func (s *LiteralFloatASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterLiteralFloatAST(s)
	}
}

func (s *LiteralFloatASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitLiteralFloatAST(s)
	}
}

func (s *LiteralFloatASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserRUNELITERAL, 0)
}

func (s *LiteralRuneASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterLiteralRuneAST(s)
	}
}

func (s *LiteralRuneASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitLiteralRuneAST(s)
	}
}

func (s *LiteralRuneASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserRAWSTRINGLITERAL, 0)
}

func (s *LiteralRawStringASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterLiteralRawStringAST(s)
	}
}

func (s *LiteralRawStringASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitLiteralRawStringAST(s)
	}
}

func (s *LiteralRawStringASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitLiteralRawStringAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MiniGoParserRULE_literal)
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserINTLITERAL:
		localctx = NewLiteralIntASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(364)
			p.Match(MiniGoParserINTLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserFLOATLITERAL:
		localctx = NewLiteralFloatASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(365)
			p.Match(MiniGoParserFLOATLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserRUNELITERAL:
		localctx = NewLiteralRuneASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(366)
			p.Match(MiniGoParserRUNELITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserRAWSTRINGLITERAL:
		localctx = NewLiteralRawStringASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(367)
			p.Match(MiniGoParserRAWSTRINGLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserINTERPRETEDSTRINGLITERAL:
		localctx = NewLiteralInterpretedStringASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(368)
			p.Match(MiniGoParserINTERPRETEDSTRINGLITERAL)
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

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	Expression() IExpressionContext
	RBRACKET() antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLBRACKET, 0)
}

func (s *IndexContext) Expression() IExpressionContext {
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

func (s *IndexContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRBRACKET, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MiniGoParserRULE_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Match(MiniGoParserLBRACKET)
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
		p.Match(MiniGoParserRBRACKET)
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
	p.RuleIndex = MiniGoParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_arguments

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
	return s.GetToken(MiniGoParserLPAREN, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterArgumentsAST(s)
	}
}

func (s *ArgumentsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitArgumentsAST(s)
	}
}

func (s *ArgumentsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *ArgumentsEmptyASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterArgumentsEmptyAST(s)
	}
}

func (s *ArgumentsEmptyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitArgumentsEmptyAST(s)
	}
}

func (s *ArgumentsEmptyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitArgumentsEmptyAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MiniGoParserRULE_arguments)
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserLPAREN:
		localctx = NewArgumentsASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(375)
			p.Match(MiniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.ExpressionList()
		}

	case MiniGoParserRPAREN:
		localctx = NewArgumentsEmptyASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(377)
			p.Epsilon()
		}
		{
			p.SetState(378)
			p.Match(MiniGoParserRPAREN)
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

	// Getter signatures
	DOT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) DOT() antlr.TerminalNode {
	return s.GetToken(MiniGoParserDOT, 0)
}

func (s *SelectorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MiniGoParserIDENTIFIER, 0)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (s *SelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, MiniGoParserRULE_selector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.Match(MiniGoParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
		p.Match(MiniGoParserIDENTIFIER)
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

	// Getter signatures
	APPEND() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COMMA() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_appendExpression
	return p
}

func InitEmptyAppendExpressionContext(p *AppendExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_appendExpression
}

func (*AppendExpressionContext) IsAppendExpressionContext() {}

func NewAppendExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppendExpressionContext {
	var p = new(AppendExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_appendExpression

	return p
}

func (s *AppendExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AppendExpressionContext) APPEND() antlr.TerminalNode {
	return s.GetToken(MiniGoParserAPPEND, 0)
}

func (s *AppendExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLPAREN, 0)
}

func (s *AppendExpressionContext) AllExpression() []IExpressionContext {
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

func (s *AppendExpressionContext) Expression(i int) IExpressionContext {
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

func (s *AppendExpressionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCOMMA, 0)
}

func (s *AppendExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *AppendExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppendExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAppendExpression(s)
	}
}

func (s *AppendExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAppendExpression(s)
	}
}

func (s *AppendExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAppendExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) AppendExpression() (localctx IAppendExpressionContext) {
	localctx = NewAppendExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MiniGoParserRULE_appendExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(MiniGoParserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Match(MiniGoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.expression(0)
	}
	{
		p.SetState(388)
		p.Match(MiniGoParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(389)
		p.expression(0)
	}
	{
		p.SetState(390)
		p.Match(MiniGoParserRPAREN)
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

	// Getter signatures
	LEN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_lengthExpression
	return p
}

func InitEmptyLengthExpressionContext(p *LengthExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_lengthExpression
}

func (*LengthExpressionContext) IsLengthExpressionContext() {}

func NewLengthExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthExpressionContext {
	var p = new(LengthExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_lengthExpression

	return p
}

func (s *LengthExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthExpressionContext) LEN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLEN, 0)
}

func (s *LengthExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLPAREN, 0)
}

func (s *LengthExpressionContext) Expression() IExpressionContext {
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

func (s *LengthExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *LengthExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterLengthExpression(s)
	}
}

func (s *LengthExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitLengthExpression(s)
	}
}

func (s *LengthExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitLengthExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) LengthExpression() (localctx ILengthExpressionContext) {
	localctx = NewLengthExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MiniGoParserRULE_lengthExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Match(MiniGoParserLEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(393)
		p.Match(MiniGoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(394)
		p.expression(0)
	}
	{
		p.SetState(395)
		p.Match(MiniGoParserRPAREN)
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

	// Getter signatures
	CAP() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_capExpression
	return p
}

func InitEmptyCapExpressionContext(p *CapExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_capExpression
}

func (*CapExpressionContext) IsCapExpressionContext() {}

func NewCapExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CapExpressionContext {
	var p = new(CapExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_capExpression

	return p
}

func (s *CapExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CapExpressionContext) CAP() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCAP, 0)
}

func (s *CapExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLPAREN, 0)
}

func (s *CapExpressionContext) Expression() IExpressionContext {
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

func (s *CapExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *CapExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CapExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterCapExpression(s)
	}
}

func (s *CapExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitCapExpression(s)
	}
}

func (s *CapExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitCapExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) CapExpression() (localctx ICapExpressionContext) {
	localctx = NewCapExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MiniGoParserRULE_capExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(MiniGoParserCAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Match(MiniGoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(399)
		p.expression(0)
	}
	{
		p.SetState(400)
		p.Match(MiniGoParserRPAREN)
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

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

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
	p.RuleIndex = MiniGoParserRULE_statementList
	return p
}

func InitEmptyStatementListContext(p *StatementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_statementList
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
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

func (s *StatementListContext) Statement(i int) IStatementContext {
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

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MiniGoParserRULE_statementList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(402)
				p.Statement()
			}

		}
		p.SetState(407)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
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

	// Getter signatures
	LBRACE() antlr.TerminalNode
	StatementList() IStatementListContext
	RBRACE() antlr.TerminalNode

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
	p.RuleIndex = MiniGoParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLBRACE, 0)
}

func (s *BlockContext) StatementList() IStatementListContext {
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

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRBRACE, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MiniGoParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(408)
		p.Match(MiniGoParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(409)
		p.StatementList()
	}
	{
		p.SetState(410)
		p.Match(MiniGoParserRBRACE)
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
	p.RuleIndex = MiniGoParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_statement

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
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *StatementSimpleASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementSimpleAST(s)
	}
}

func (s *StatementSimpleASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementSimpleAST(s)
	}
}

func (s *StatementSimpleASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *StatementSwitchASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementSwitchAST(s)
	}
}

func (s *StatementSwitchASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementSwitchAST(s)
	}
}

func (s *StatementSwitchASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *StatementLoopASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementLoopAST(s)
	}
}

func (s *StatementLoopASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementLoopAST(s)
	}
}

func (s *StatementLoopASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserBREAK, 0)
}

func (s *StatementBreakASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *StatementBreakASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementBreakAST(s)
	}
}

func (s *StatementBreakASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementBreakAST(s)
	}
}

func (s *StatementBreakASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *StatementBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementBlockAST(s)
	}
}

func (s *StatementBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementBlockAST(s)
	}
}

func (s *StatementBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementTypeDeclAST(s)
	}
}

func (s *StatementTypeDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementTypeDeclAST(s)
	}
}

func (s *StatementTypeDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserRETURN, 0)
}

func (s *StatementReturnASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementReturnAST(s)
	}
}

func (s *StatementReturnASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementReturnAST(s)
	}
}

func (s *StatementReturnASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserPRINT, 0)
}

func (s *StatementPrintASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLPAREN, 0)
}

func (s *StatementPrintASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *StatementPrintASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementPrintAST(s)
	}
}

func (s *StatementPrintASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementPrintAST(s)
	}
}

func (s *StatementPrintASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementVariableDeclAST(s)
	}
}

func (s *StatementVariableDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementVariableDeclAST(s)
	}
}

func (s *StatementVariableDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserPRINTLN, 0)
}

func (s *StatementPrintlnASTContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLPAREN, 0)
}

func (s *StatementPrintlnASTContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MiniGoParserRPAREN, 0)
}

func (s *StatementPrintlnASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementPrintlnAST(s)
	}
}

func (s *StatementPrintlnASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementPrintlnAST(s)
	}
}

func (s *StatementPrintlnASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserCONTINUE, 0)
}

func (s *StatementContinueASTContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *StatementContinueASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementContinueAST(s)
	}
}

func (s *StatementContinueASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementContinueAST(s)
	}
}

func (s *StatementContinueASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *StatementIfASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterStatementIfAST(s)
	}
}

func (s *StatementIfASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitStatementIfAST(s)
	}
}

func (s *StatementIfASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitStatementIfAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, MiniGoParserRULE_statement)
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserPRINT:
		localctx = NewStatementPrintASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(412)
			p.Match(MiniGoParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.Match(MiniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MiniGoParserINTLITERAL, MiniGoParserFLOATLITERAL, MiniGoParserRUNELITERAL, MiniGoParserRAWSTRINGLITERAL, MiniGoParserINTERPRETEDSTRINGLITERAL, MiniGoParserLEN, MiniGoParserCAP, MiniGoParserAPPEND, MiniGoParserIDENTIFIER, MiniGoParserPLUS, MiniGoParserMINUS, MiniGoParserNOT, MiniGoParserBITWISEXOR, MiniGoParserLPAREN:
			{
				p.SetState(414)
				p.ExpressionList()
			}

		case MiniGoParserRPAREN:
			{
				p.SetState(415)
				p.Epsilon()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(418)
			p.Match(MiniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserPRINTLN:
		localctx = NewStatementPrintlnASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(421)
			p.Match(MiniGoParserPRINTLN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.Match(MiniGoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(425)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MiniGoParserINTLITERAL, MiniGoParserFLOATLITERAL, MiniGoParserRUNELITERAL, MiniGoParserRAWSTRINGLITERAL, MiniGoParserINTERPRETEDSTRINGLITERAL, MiniGoParserLEN, MiniGoParserCAP, MiniGoParserAPPEND, MiniGoParserIDENTIFIER, MiniGoParserPLUS, MiniGoParserMINUS, MiniGoParserNOT, MiniGoParserBITWISEXOR, MiniGoParserLPAREN:
			{
				p.SetState(423)
				p.ExpressionList()
			}

		case MiniGoParserRPAREN:
			{
				p.SetState(424)
				p.Epsilon()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(427)
			p.Match(MiniGoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserRETURN:
		localctx = NewStatementReturnASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(430)
			p.Match(MiniGoParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(433)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MiniGoParserINTLITERAL, MiniGoParserFLOATLITERAL, MiniGoParserRUNELITERAL, MiniGoParserRAWSTRINGLITERAL, MiniGoParserINTERPRETEDSTRINGLITERAL, MiniGoParserLEN, MiniGoParserCAP, MiniGoParserAPPEND, MiniGoParserIDENTIFIER, MiniGoParserPLUS, MiniGoParserMINUS, MiniGoParserNOT, MiniGoParserBITWISEXOR, MiniGoParserLPAREN:
			{
				p.SetState(431)
				p.expression(0)
			}

		case MiniGoParserSEMICOLON:
			{
				p.SetState(432)
				p.Epsilon()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(435)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserBREAK:
		localctx = NewStatementBreakASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(437)
			p.Match(MiniGoParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserCONTINUE:
		localctx = NewStatementContinueASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(439)
			p.Match(MiniGoParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(440)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserINTLITERAL, MiniGoParserFLOATLITERAL, MiniGoParserRUNELITERAL, MiniGoParserRAWSTRINGLITERAL, MiniGoParserINTERPRETEDSTRINGLITERAL, MiniGoParserLEN, MiniGoParserCAP, MiniGoParserAPPEND, MiniGoParserIDENTIFIER, MiniGoParserPLUS, MiniGoParserMINUS, MiniGoParserNOT, MiniGoParserBITWISEXOR, MiniGoParserSEMICOLON, MiniGoParserLPAREN:
		localctx = NewStatementSimpleASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(441)
			p.SimpleStatement()
		}
		{
			p.SetState(442)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserLBRACE:
		localctx = NewStatementBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(444)
			p.Block()
		}
		{
			p.SetState(445)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserSWITCH:
		localctx = NewStatementSwitchASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(447)
			p.SwitchStmt()
		}
		{
			p.SetState(448)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserIF:
		localctx = NewStatementIfASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(450)
			p.IfStatement()
		}
		{
			p.SetState(451)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserFOR:
		localctx = NewStatementLoopASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(453)
			p.Loop()
		}
		{
			p.SetState(454)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniGoParserTYPE:
		localctx = NewStatementTypeDeclASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(456)
			p.TypeDecl()
		}

	case MiniGoParserVAR:
		localctx = NewStatementVariableDeclASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(457)
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
	p.RuleIndex = MiniGoParserRULE_simpleStatement
	return p
}

func InitEmptySimpleStatementContext(p *SimpleStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_simpleStatement
}

func (*SimpleStatementContext) IsSimpleStatementContext() {}

func NewSimpleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStatementContext {
	var p = new(SimpleStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_simpleStatement

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
	return s.GetToken(MiniGoParserINCREMENT, 0)
}

func (s *SimpleStatementExpressionASTContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(MiniGoParserDECREMENT, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSimpleStatementExpressionAST(s)
	}
}

func (s *SimpleStatementExpressionASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSimpleStatementExpressionAST(s)
	}
}

func (s *SimpleStatementExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSimpleStatementAssignmentAST(s)
	}
}

func (s *SimpleStatementAssignmentASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSimpleStatementAssignmentAST(s)
	}
}

func (s *SimpleStatementAssignmentASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserASSIGN, 0)
}

func (s *SimpleStatementExpressionListAssignASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSimpleStatementExpressionListAssignAST(s)
	}
}

func (s *SimpleStatementExpressionListAssignASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSimpleStatementExpressionListAssignAST(s)
	}
}

func (s *SimpleStatementExpressionListAssignASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSimpleStatementEmptyAST(s)
	}
}

func (s *SimpleStatementEmptyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSimpleStatementEmptyAST(s)
	}
}

func (s *SimpleStatementEmptyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSimpleStatementEmptyAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) SimpleStatement() (localctx ISimpleStatementContext) {
	localctx = NewSimpleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, MiniGoParserRULE_simpleStatement)
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpleStatementEmptyASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(460)
			p.Epsilon()
		}

	case 2:
		localctx = NewSimpleStatementExpressionASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(461)
			p.expression(0)
		}
		p.SetState(465)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MiniGoParserINCREMENT:
			{
				p.SetState(462)
				p.Match(MiniGoParserINCREMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case MiniGoParserDECREMENT:
			{
				p.SetState(463)
				p.Match(MiniGoParserDECREMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case MiniGoParserSEMICOLON, MiniGoParserLBRACE:
			{
				p.SetState(464)
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
			p.SetState(467)
			p.AssignmentStatement()
		}

	case 4:
		localctx = NewSimpleStatementExpressionListAssignASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(468)
			p.ExpressionList()
		}
		{
			p.SetState(469)
			p.Match(MiniGoParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
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
	p.RuleIndex = MiniGoParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_assignmentStatement

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
	return s.GetToken(MiniGoParserBITWISEANDEQUAL, 0)
}

func (s *AssignmentStatementBitwiseAndEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAssignmentStatementBitwiseAndEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseAndEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAssignmentStatementBitwiseAndEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseAndEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserBITWISEXOREQUAL, 0)
}

func (s *AssignmentStatementBitwiseXorEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAssignmentStatementBitwiseXorEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseXorEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAssignmentStatementBitwiseXorEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseXorEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserSHIFTLEFTEQUAL, 0)
}

func (s *AssignmentStatementShiftLeftEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAssignmentStatementShiftLeftEqualAST(s)
	}
}

func (s *AssignmentStatementShiftLeftEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAssignmentStatementShiftLeftEqualAST(s)
	}
}

func (s *AssignmentStatementShiftLeftEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserPLUSEQUAL, 0)
}

func (s *AssignmentStatementPlusEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAssignmentStatementPlusEqualAST(s)
	}
}

func (s *AssignmentStatementPlusEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAssignmentStatementPlusEqualAST(s)
	}
}

func (s *AssignmentStatementPlusEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserMINUSEQUAL, 0)
}

func (s *AssignmentStatementMinusEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAssignmentStatementMinusEqualAST(s)
	}
}

func (s *AssignmentStatementMinusEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAssignmentStatementMinusEqualAST(s)
	}
}

func (s *AssignmentStatementMinusEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserMODULOEQUAL, 0)
}

func (s *AssignmentStatementModuloEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAssignmentStatementModuloEqualAST(s)
	}
}

func (s *AssignmentStatementModuloEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAssignmentStatementModuloEqualAST(s)
	}
}

func (s *AssignmentStatementModuloEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserASSIGN, 0)
}

func (s *AssignmentStatementASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAssignmentStatementAST(s)
	}
}

func (s *AssignmentStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAssignmentStatementAST(s)
	}
}

func (s *AssignmentStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserBITWISECLEAREQUAL, 0)
}

func (s *AssignmentStatementBitwiseClearEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAssignmentStatementBitwiseClearEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseClearEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAssignmentStatementBitwiseClearEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseClearEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserDIVIDEEQUAL, 0)
}

func (s *AssignmentStatementDivideEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAssignmentStatementDivideEqualAST(s)
	}
}

func (s *AssignmentStatementDivideEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAssignmentStatementDivideEqualAST(s)
	}
}

func (s *AssignmentStatementDivideEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserSHIFTRIGHTEQUAL, 0)
}

func (s *AssignmentStatementShiftRightEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAssignmentStatementShiftRightEqualAST(s)
	}
}

func (s *AssignmentStatementShiftRightEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAssignmentStatementShiftRightEqualAST(s)
	}
}

func (s *AssignmentStatementShiftRightEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserMULTIPLYEQUAL, 0)
}

func (s *AssignmentStatementMultiplyEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAssignmentStatementMultiplyEqualAST(s)
	}
}

func (s *AssignmentStatementMultiplyEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAssignmentStatementMultiplyEqualAST(s)
	}
}

func (s *AssignmentStatementMultiplyEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserBITWISEOREQUAL, 0)
}

func (s *AssignmentStatementBitwiseOrEqualASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterAssignmentStatementBitwiseOrEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseOrEqualASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitAssignmentStatementBitwiseOrEqualAST(s)
	}
}

func (s *AssignmentStatementBitwiseOrEqualASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitAssignmentStatementBitwiseOrEqualAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, MiniGoParserRULE_assignmentStatement)
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignmentStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(474)
			p.ExpressionList()
		}
		{
			p.SetState(475)
			p.Match(MiniGoParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)
			p.ExpressionList()
		}

	case 2:
		localctx = NewAssignmentStatementPlusEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(478)
			p.expression(0)
		}
		{
			p.SetState(479)
			p.Match(MiniGoParserPLUSEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)
			p.expression(0)
		}

	case 3:
		localctx = NewAssignmentStatementMinusEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(482)
			p.expression(0)
		}
		{
			p.SetState(483)
			p.Match(MiniGoParserMINUSEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(484)
			p.expression(0)
		}

	case 4:
		localctx = NewAssignmentStatementBitwiseAndEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(486)
			p.expression(0)
		}
		{
			p.SetState(487)
			p.Match(MiniGoParserBITWISEANDEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(488)
			p.expression(0)
		}

	case 5:
		localctx = NewAssignmentStatementBitwiseOrEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(490)
			p.expression(0)
		}
		{
			p.SetState(491)
			p.Match(MiniGoParserBITWISEOREQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(492)
			p.expression(0)
		}

	case 6:
		localctx = NewAssignmentStatementMultiplyEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(494)
			p.expression(0)
		}
		{
			p.SetState(495)
			p.Match(MiniGoParserMULTIPLYEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(496)
			p.expression(0)
		}

	case 7:
		localctx = NewAssignmentStatementBitwiseXorEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(498)
			p.expression(0)
		}
		{
			p.SetState(499)
			p.Match(MiniGoParserBITWISEXOREQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(500)
			p.expression(0)
		}

	case 8:
		localctx = NewAssignmentStatementShiftLeftEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(502)
			p.expression(0)
		}
		{
			p.SetState(503)
			p.Match(MiniGoParserSHIFTLEFTEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(504)
			p.expression(0)
		}

	case 9:
		localctx = NewAssignmentStatementShiftRightEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(506)
			p.expression(0)
		}
		{
			p.SetState(507)
			p.Match(MiniGoParserSHIFTRIGHTEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(508)
			p.expression(0)
		}

	case 10:
		localctx = NewAssignmentStatementBitwiseClearEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(510)
			p.expression(0)
		}
		{
			p.SetState(511)
			p.Match(MiniGoParserBITWISECLEAREQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)
			p.expression(0)
		}

	case 11:
		localctx = NewAssignmentStatementModuloEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(514)
			p.expression(0)
		}
		{
			p.SetState(515)
			p.Match(MiniGoParserMODULOEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(516)
			p.expression(0)
		}

	case 12:
		localctx = NewAssignmentStatementDivideEqualASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(518)
			p.expression(0)
		}
		{
			p.SetState(519)
			p.Match(MiniGoParserDIVIDEEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(520)
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
	p.RuleIndex = MiniGoParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_ifStatement

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
	return s.GetToken(MiniGoParserIF, 0)
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
	return s.GetToken(MiniGoParserSEMICOLON, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterIfSimpleStatementAST(s)
	}
}

func (s *IfSimpleStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitIfSimpleStatementAST(s)
	}
}

func (s *IfSimpleStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserIF, 0)
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
	return s.GetToken(MiniGoParserELSE, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterIfElseIfStatementAST(s)
	}
}

func (s *IfElseIfStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitIfElseIfStatementAST(s)
	}
}

func (s *IfElseIfStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserIF, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterIfStatementAST(s)
	}
}

func (s *IfStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitIfStatementAST(s)
	}
}

func (s *IfStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserIF, 0)
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
	return s.GetToken(MiniGoParserELSE, 0)
}

func (s *IfElseStatementASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterIfElseStatementAST(s)
	}
}

func (s *IfElseStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitIfElseStatementAST(s)
	}
}

func (s *IfElseStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserIF, 0)
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
	return s.GetToken(MiniGoParserSEMICOLON, 0)
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
	return s.GetToken(MiniGoParserELSE, 0)
}

func (s *IfSimpleElseStatementASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterIfSimpleElseStatementAST(s)
	}
}

func (s *IfSimpleElseStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitIfSimpleElseStatementAST(s)
	}
}

func (s *IfSimpleElseStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserIF, 0)
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
	return s.GetToken(MiniGoParserSEMICOLON, 0)
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
	return s.GetToken(MiniGoParserELSE, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterIfSimpleElseIfStatementAST(s)
	}
}

func (s *IfSimpleElseIfStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitIfSimpleElseIfStatementAST(s)
	}
}

func (s *IfSimpleElseIfStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitIfSimpleElseIfStatementAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, MiniGoParserRULE_ifStatement)
	p.SetState(562)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(524)
			p.Match(MiniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(525)
			p.expression(0)
		}
		{
			p.SetState(526)
			p.Block()
		}

	case 2:
		localctx = NewIfElseIfStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(528)
			p.Match(MiniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(529)
			p.expression(0)
		}
		{
			p.SetState(530)
			p.Block()
		}
		{
			p.SetState(531)
			p.Match(MiniGoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(532)
			p.IfStatement()
		}

	case 3:
		localctx = NewIfElseStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(534)
			p.Match(MiniGoParserIF)
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
			p.Match(MiniGoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(538)
			p.Block()
		}

	case 4:
		localctx = NewIfSimpleStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(540)
			p.Match(MiniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(541)
			p.SimpleStatement()
		}
		{
			p.SetState(542)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(543)
			p.expression(0)
		}
		{
			p.SetState(544)
			p.Block()
		}

	case 5:
		localctx = NewIfSimpleElseIfStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(546)
			p.Match(MiniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(547)
			p.SimpleStatement()
		}
		{
			p.SetState(548)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(549)
			p.expression(0)
		}
		{
			p.SetState(550)
			p.Block()
		}
		{
			p.SetState(551)
			p.Match(MiniGoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(552)
			p.IfStatement()
		}

	case 6:
		localctx = NewIfSimpleElseStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(554)
			p.Match(MiniGoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(555)
			p.SimpleStatement()
		}
		{
			p.SetState(556)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(557)
			p.expression(0)
		}
		{
			p.SetState(558)
			p.Block()
		}
		{
			p.SetState(559)
			p.Match(MiniGoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(560)
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
	p.RuleIndex = MiniGoParserRULE_loop
	return p
}

func InitEmptyLoopContext(p *LoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_loop
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_loop

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
	return s.GetToken(MiniGoParserFOR, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterLoopExpressionBlockAST(s)
	}
}

func (s *LoopExpressionBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitLoopExpressionBlockAST(s)
	}
}

func (s *LoopExpressionBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserFOR, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterLoopBlockAST(s)
	}
}

func (s *LoopBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitLoopBlockAST(s)
	}
}

func (s *LoopBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserFOR, 0)
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
	return s.GetTokens(MiniGoParserSEMICOLON)
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, i)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterLoopSimpleStatementExpressionSimpleStatementBlockAST(s)
	}
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitLoopSimpleStatementExpressionSimpleStatementBlockAST(s)
	}
}

func (s *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserFOR, 0)
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
	return s.GetTokens(MiniGoParserSEMICOLON)
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(MiniGoParserSEMICOLON, i)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterLoopSimpleStatementSimpleStatementBlockAST(s)
	}
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitLoopSimpleStatementSimpleStatementBlockAST(s)
	}
}

func (s *LoopSimpleStatementSimpleStatementBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitLoopSimpleStatementSimpleStatementBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, MiniGoParserRULE_loop)
	p.SetState(585)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLoopBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(564)
			p.Match(MiniGoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(565)
			p.Block()
		}

	case 2:
		localctx = NewLoopExpressionBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(566)
			p.Match(MiniGoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(567)
			p.expression(0)
		}
		{
			p.SetState(568)
			p.Block()
		}

	case 3:
		localctx = NewLoopSimpleStatementExpressionSimpleStatementBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(570)
			p.Match(MiniGoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(571)
			p.SimpleStatement()
		}
		{
			p.SetState(572)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(573)
			p.expression(0)
		}
		{
			p.SetState(574)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(575)
			p.SimpleStatement()
		}
		{
			p.SetState(576)
			p.Block()
		}

	case 4:
		localctx = NewLoopSimpleStatementSimpleStatementBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(578)
			p.Match(MiniGoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(579)
			p.SimpleStatement()
		}
		{
			p.SetState(580)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(581)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(582)
			p.SimpleStatement()
		}
		{
			p.SetState(583)
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
	p.RuleIndex = MiniGoParserRULE_switchStmt
	return p
}

func InitEmptySwitchStmtContext(p *SwitchStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_switchStmt
}

func (*SwitchStmtContext) IsSwitchStmtContext() {}

func NewSwitchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_switchStmt

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
	return s.GetToken(MiniGoParserSWITCH, 0)
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
	return s.GetToken(MiniGoParserSEMICOLON, 0)
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
	return s.GetToken(MiniGoParserLBRACE, 0)
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
	return s.GetToken(MiniGoParserRBRACE, 0)
}

func (s *SwitchStmtSimpleStatementASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSwitchStmtSimpleStatementAST(s)
	}
}

func (s *SwitchStmtSimpleStatementASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSwitchStmtSimpleStatementAST(s)
	}
}

func (s *SwitchStmtSimpleStatementASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserSWITCH, 0)
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
	return s.GetToken(MiniGoParserLBRACE, 0)
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
	return s.GetToken(MiniGoParserRBRACE, 0)
}

func (s *SwitchStmtExpressionASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSwitchStmtExpressionAST(s)
	}
}

func (s *SwitchStmtExpressionASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSwitchStmtExpressionAST(s)
	}
}

func (s *SwitchStmtExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserSWITCH, 0)
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
	return s.GetToken(MiniGoParserSEMICOLON, 0)
}

func (s *SwitchStmtSimpleStatementBlockASTContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLBRACE, 0)
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
	return s.GetToken(MiniGoParserRBRACE, 0)
}

func (s *SwitchStmtSimpleStatementBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSwitchStmtSimpleStatementBlockAST(s)
	}
}

func (s *SwitchStmtSimpleStatementBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSwitchStmtSimpleStatementBlockAST(s)
	}
}

func (s *SwitchStmtSimpleStatementBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserSWITCH, 0)
}

func (s *SwitchStmtBlockASTContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MiniGoParserLBRACE, 0)
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
	return s.GetToken(MiniGoParserRBRACE, 0)
}

func (s *SwitchStmtBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterSwitchStmtBlockAST(s)
	}
}

func (s *SwitchStmtBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitSwitchStmtBlockAST(s)
	}
}

func (s *SwitchStmtBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitSwitchStmtBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) SwitchStmt() (localctx ISwitchStmtContext) {
	localctx = NewSwitchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, MiniGoParserRULE_switchStmt)
	p.SetState(613)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSwitchStmtSimpleStatementASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(587)
			p.Match(MiniGoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(588)
			p.SimpleStatement()
		}
		{
			p.SetState(589)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(590)
			p.expression(0)
		}
		{
			p.SetState(591)
			p.Match(MiniGoParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(592)
			p.expressionCaseClauseList(0)
		}
		{
			p.SetState(593)
			p.Match(MiniGoParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSwitchStmtExpressionASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(595)
			p.Match(MiniGoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(596)
			p.expression(0)
		}
		{
			p.SetState(597)
			p.Match(MiniGoParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(598)
			p.expressionCaseClauseList(0)
		}
		{
			p.SetState(599)
			p.Match(MiniGoParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSwitchStmtSimpleStatementBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(601)
			p.Match(MiniGoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(602)
			p.SimpleStatement()
		}
		{
			p.SetState(603)
			p.Match(MiniGoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(604)
			p.Match(MiniGoParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(605)
			p.expressionCaseClauseList(0)
		}
		{
			p.SetState(606)
			p.Match(MiniGoParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewSwitchStmtBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(608)
			p.Match(MiniGoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(609)
			p.Match(MiniGoParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(610)
			p.expressionCaseClauseList(0)
		}
		{
			p.SetState(611)
			p.Match(MiniGoParserRBRACE)
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
	p.RuleIndex = MiniGoParserRULE_expressionCaseClauseList
	return p
}

func InitEmptyExpressionCaseClauseListContext(p *ExpressionCaseClauseListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expressionCaseClauseList
}

func (*ExpressionCaseClauseListContext) IsExpressionCaseClauseListContext() {}

func NewExpressionCaseClauseListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_expressionCaseClauseList

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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionCaseClauseListEmptyAST(s)
	}
}

func (s *ExpressionCaseClauseListEmptyASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionCaseClauseListEmptyAST(s)
	}
}

func (s *ExpressionCaseClauseListEmptyASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionCaseClauseListAST(s)
	}
}

func (s *ExpressionCaseClauseListASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionCaseClauseListAST(s)
	}
}

func (s *ExpressionCaseClauseListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExpressionCaseClauseListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) ExpressionCaseClauseList() (localctx IExpressionCaseClauseListContext) {
	return p.expressionCaseClauseList(0)
}

func (p *MiniGoParser) expressionCaseClauseList(_p int) (localctx IExpressionCaseClauseListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionCaseClauseListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionCaseClauseListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 80
	p.EnterRecursionRule(localctx, 80, MiniGoParserRULE_expressionCaseClauseList, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewExpressionCaseClauseListEmptyASTContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(616)
		p.Epsilon()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(622)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
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
			p.PushNewRecursionContext(localctx, _startState, MiniGoParserRULE_expressionCaseClauseList)
			p.SetState(618)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(619)
				p.ExpressionCaseClause()
			}

		}
		p.SetState(624)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
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

	// Getter signatures
	ExpressionSwitchCase() IExpressionSwitchCaseContext
	COLON() antlr.TerminalNode
	StatementList() IStatementListContext

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
	p.RuleIndex = MiniGoParserRULE_expressionCaseClause
	return p
}

func InitEmptyExpressionCaseClauseContext(p *ExpressionCaseClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expressionCaseClause
}

func (*ExpressionCaseClauseContext) IsExpressionCaseClauseContext() {}

func NewExpressionCaseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_expressionCaseClause

	return p
}

func (s *ExpressionCaseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseContext) ExpressionSwitchCase() IExpressionSwitchCaseContext {
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

func (s *ExpressionCaseClauseContext) COLON() antlr.TerminalNode {
	return s.GetToken(MiniGoParserCOLON, 0)
}

func (s *ExpressionCaseClauseContext) StatementList() IStatementListContext {
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

func (s *ExpressionCaseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionCaseClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionCaseClause(s)
	}
}

func (s *ExpressionCaseClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionCaseClause(s)
	}
}

func (s *ExpressionCaseClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExpressionCaseClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) ExpressionCaseClause() (localctx IExpressionCaseClauseContext) {
	localctx = NewExpressionCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, MiniGoParserRULE_expressionCaseClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(625)
		p.ExpressionSwitchCase()
	}
	{
		p.SetState(626)
		p.Match(MiniGoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(627)
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
	p.RuleIndex = MiniGoParserRULE_expressionSwitchCase
	return p
}

func InitEmptyExpressionSwitchCaseContext(p *ExpressionSwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_expressionSwitchCase
}

func (*ExpressionSwitchCaseContext) IsExpressionSwitchCaseContext() {}

func NewExpressionSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_expressionSwitchCase

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
	return s.GetToken(MiniGoParserCASE, 0)
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
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionSwitchCaseAST(s)
	}
}

func (s *ExpressionSwitchCaseASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionSwitchCaseAST(s)
	}
}

func (s *ExpressionSwitchCaseASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
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
	return s.GetToken(MiniGoParserDEFAULT, 0)
}

func (s *ExpressionSwitchDefaultASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterExpressionSwitchDefaultAST(s)
	}
}

func (s *ExpressionSwitchDefaultASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitExpressionSwitchDefaultAST(s)
	}
}

func (s *ExpressionSwitchDefaultASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitExpressionSwitchDefaultAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) ExpressionSwitchCase() (localctx IExpressionSwitchCaseContext) {
	localctx = NewExpressionSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, MiniGoParserRULE_expressionSwitchCase)
	p.SetState(632)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniGoParserCASE:
		localctx = NewExpressionSwitchCaseASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(629)
			p.Match(MiniGoParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(630)
			p.ExpressionList()
		}

	case MiniGoParserDEFAULT:
		localctx = NewExpressionSwitchDefaultASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(631)
			p.Match(MiniGoParserDEFAULT)
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
	p.RuleIndex = MiniGoParserRULE_epsilon
	return p
}

func InitEmptyEpsilonContext(p *EpsilonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniGoParserRULE_epsilon
}

func (*EpsilonContext) IsEpsilonContext() {}

func NewEpsilonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EpsilonContext {
	var p = new(EpsilonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniGoParserRULE_epsilon

	return p
}

func (s *EpsilonContext) GetParser() antlr.Parser { return s.parser }
func (s *EpsilonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpsilonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EpsilonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.EnterEpsilon(s)
	}
}

func (s *EpsilonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniGoParserListener); ok {
		listenerT.ExitEpsilon(s)
	}
}

func (s *EpsilonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniGoParserVisitor:
		return t.VisitEpsilon(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MiniGoParser) Epsilon() (localctx IEpsilonContext) {
	localctx = NewEpsilonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, MiniGoParserRULE_epsilon)
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

func (p *MiniGoParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 21:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 23:
		var t *PrimaryExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExpressionContext)
		}
		return p.PrimaryExpression_Sempred(t, predIndex)

	case 40:
		var t *ExpressionCaseClauseListContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionCaseClauseListContext)
		}
		return p.ExpressionCaseClauseList_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MiniGoParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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

func (p *MiniGoParser) PrimaryExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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

func (p *MiniGoParser) ExpressionCaseClauseList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 22:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
