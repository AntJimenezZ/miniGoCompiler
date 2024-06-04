// Code generated from C:/Users/noni4/Desktop/miniGoCompiler/serverAndCompiler/MiniGoScanner.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MiniGoScanner struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MiniGoScannerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func minigoscannerLexerInit() {
	staticData := &MiniGoScannerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"BLOCK_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 72, 451, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 1, 0, 4, 0, 147,
		8, 0, 11, 0, 12, 0, 148, 1, 1, 4, 1, 152, 8, 1, 11, 1, 12, 1, 153, 1, 1,
		1, 1, 4, 1, 158, 8, 1, 11, 1, 12, 1, 159, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 5, 3, 170, 8, 3, 10, 3, 12, 3, 173, 9, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 181, 8, 4, 10, 4, 12, 4, 184, 9, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 5, 24, 302, 8, 24, 10, 24,
		12, 24, 305, 9, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33,
		1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1,
		36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41,
		1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1,
		45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48,
		1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 52, 1,
		52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55,
		1, 56, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1,
		58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63,
		1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66, 1, 66, 1, 67, 1, 67, 1, 68, 1,
		68, 1, 69, 4, 69, 421, 8, 69, 11, 69, 12, 69, 422, 1, 69, 1, 69, 1, 70,
		1, 70, 1, 70, 1, 70, 5, 70, 431, 8, 70, 10, 70, 12, 70, 434, 9, 70, 1,
		70, 1, 70, 1, 71, 1, 71, 1, 71, 1, 71, 5, 71, 442, 8, 71, 10, 71, 12, 71,
		445, 9, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 443, 0, 72, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30,
		61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39,
		79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48,
		97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 56, 113,
		57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62, 125, 63, 127, 64, 129,
		65, 131, 66, 133, 67, 135, 68, 137, 69, 139, 70, 141, 71, 143, 72, 1, 0,
		7, 1, 0, 48, 57, 2, 0, 92, 92, 96, 96, 2, 0, 34, 34, 92, 92, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10,
		13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 461, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0,
		0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1,
		0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73,
		1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0,
		81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0,
		0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0,
		0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1,
		0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0,
		111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0,
		0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125,
		1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 0, 131, 1, 0, 0, 0,
		0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 0, 137, 1, 0, 0, 0, 0, 139, 1,
		0, 0, 0, 0, 141, 1, 0, 0, 0, 0, 143, 1, 0, 0, 0, 1, 146, 1, 0, 0, 0, 3,
		151, 1, 0, 0, 0, 5, 161, 1, 0, 0, 0, 7, 165, 1, 0, 0, 0, 9, 176, 1, 0,
		0, 0, 11, 187, 1, 0, 0, 0, 13, 195, 1, 0, 0, 0, 15, 199, 1, 0, 0, 0, 17,
		204, 1, 0, 0, 0, 19, 209, 1, 0, 0, 0, 21, 216, 1, 0, 0, 0, 23, 220, 1,
		0, 0, 0, 25, 224, 1, 0, 0, 0, 27, 230, 1, 0, 0, 0, 29, 238, 1, 0, 0, 0,
		31, 245, 1, 0, 0, 0, 33, 251, 1, 0, 0, 0, 35, 260, 1, 0, 0, 0, 37, 267,
		1, 0, 0, 0, 39, 270, 1, 0, 0, 0, 41, 275, 1, 0, 0, 0, 43, 279, 1, 0, 0,
		0, 45, 286, 1, 0, 0, 0, 47, 291, 1, 0, 0, 0, 49, 299, 1, 0, 0, 0, 51, 306,
		1, 0, 0, 0, 53, 308, 1, 0, 0, 0, 55, 310, 1, 0, 0, 0, 57, 312, 1, 0, 0,
		0, 59, 314, 1, 0, 0, 0, 61, 316, 1, 0, 0, 0, 63, 318, 1, 0, 0, 0, 65, 321,
		1, 0, 0, 0, 67, 323, 1, 0, 0, 0, 69, 326, 1, 0, 0, 0, 71, 329, 1, 0, 0,
		0, 73, 332, 1, 0, 0, 0, 75, 335, 1, 0, 0, 0, 77, 338, 1, 0, 0, 0, 79, 340,
		1, 0, 0, 0, 81, 342, 1, 0, 0, 0, 83, 344, 1, 0, 0, 0, 85, 346, 1, 0, 0,
		0, 87, 349, 1, 0, 0, 0, 89, 352, 1, 0, 0, 0, 91, 355, 1, 0, 0, 0, 93, 358,
		1, 0, 0, 0, 95, 361, 1, 0, 0, 0, 97, 363, 1, 0, 0, 0, 99, 366, 1, 0, 0,
		0, 101, 369, 1, 0, 0, 0, 103, 372, 1, 0, 0, 0, 105, 375, 1, 0, 0, 0, 107,
		378, 1, 0, 0, 0, 109, 381, 1, 0, 0, 0, 111, 384, 1, 0, 0, 0, 113, 387,
		1, 0, 0, 0, 115, 391, 1, 0, 0, 0, 117, 395, 1, 0, 0, 0, 119, 399, 1, 0,
		0, 0, 121, 401, 1, 0, 0, 0, 123, 403, 1, 0, 0, 0, 125, 405, 1, 0, 0, 0,
		127, 407, 1, 0, 0, 0, 129, 409, 1, 0, 0, 0, 131, 411, 1, 0, 0, 0, 133,
		413, 1, 0, 0, 0, 135, 415, 1, 0, 0, 0, 137, 417, 1, 0, 0, 0, 139, 420,
		1, 0, 0, 0, 141, 426, 1, 0, 0, 0, 143, 437, 1, 0, 0, 0, 145, 147, 7, 0,
		0, 0, 146, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0,
		148, 149, 1, 0, 0, 0, 149, 2, 1, 0, 0, 0, 150, 152, 7, 0, 0, 0, 151, 150,
		1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0,
		0, 0, 154, 155, 1, 0, 0, 0, 155, 157, 5, 46, 0, 0, 156, 158, 7, 0, 0, 0,
		157, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 4, 1, 0, 0, 0, 161, 162, 5, 39, 0, 0, 162, 163, 9,
		0, 0, 0, 163, 164, 5, 39, 0, 0, 164, 6, 1, 0, 0, 0, 165, 171, 5, 96, 0,
		0, 166, 167, 5, 92, 0, 0, 167, 170, 9, 0, 0, 0, 168, 170, 8, 1, 0, 0, 169,
		166, 1, 0, 0, 0, 169, 168, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169,
		1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 174, 1, 0, 0, 0, 173, 171, 1, 0,
		0, 0, 174, 175, 5, 96, 0, 0, 175, 8, 1, 0, 0, 0, 176, 182, 5, 34, 0, 0,
		177, 178, 5, 92, 0, 0, 178, 181, 9, 0, 0, 0, 179, 181, 8, 2, 0, 0, 180,
		177, 1, 0, 0, 0, 180, 179, 1, 0, 0, 0, 181, 184, 1, 0, 0, 0, 182, 180,
		1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 185, 1, 0, 0, 0, 184, 182, 1, 0,
		0, 0, 185, 186, 5, 34, 0, 0, 186, 10, 1, 0, 0, 0, 187, 188, 5, 112, 0,
		0, 188, 189, 5, 97, 0, 0, 189, 190, 5, 99, 0, 0, 190, 191, 5, 107, 0, 0,
		191, 192, 5, 97, 0, 0, 192, 193, 5, 103, 0, 0, 193, 194, 5, 101, 0, 0,
		194, 12, 1, 0, 0, 0, 195, 196, 5, 118, 0, 0, 196, 197, 5, 97, 0, 0, 197,
		198, 5, 114, 0, 0, 198, 14, 1, 0, 0, 0, 199, 200, 5, 116, 0, 0, 200, 201,
		5, 121, 0, 0, 201, 202, 5, 112, 0, 0, 202, 203, 5, 101, 0, 0, 203, 16,
		1, 0, 0, 0, 204, 205, 5, 102, 0, 0, 205, 206, 5, 117, 0, 0, 206, 207, 5,
		110, 0, 0, 207, 208, 5, 99, 0, 0, 208, 18, 1, 0, 0, 0, 209, 210, 5, 115,
		0, 0, 210, 211, 5, 116, 0, 0, 211, 212, 5, 114, 0, 0, 212, 213, 5, 117,
		0, 0, 213, 214, 5, 99, 0, 0, 214, 215, 5, 116, 0, 0, 215, 20, 1, 0, 0,
		0, 216, 217, 5, 108, 0, 0, 217, 218, 5, 101, 0, 0, 218, 219, 5, 110, 0,
		0, 219, 22, 1, 0, 0, 0, 220, 221, 5, 99, 0, 0, 221, 222, 5, 97, 0, 0, 222,
		223, 5, 112, 0, 0, 223, 24, 1, 0, 0, 0, 224, 225, 5, 112, 0, 0, 225, 226,
		5, 114, 0, 0, 226, 227, 5, 105, 0, 0, 227, 228, 5, 110, 0, 0, 228, 229,
		5, 116, 0, 0, 229, 26, 1, 0, 0, 0, 230, 231, 5, 112, 0, 0, 231, 232, 5,
		114, 0, 0, 232, 233, 5, 105, 0, 0, 233, 234, 5, 110, 0, 0, 234, 235, 5,
		116, 0, 0, 235, 236, 5, 108, 0, 0, 236, 237, 5, 110, 0, 0, 237, 28, 1,
		0, 0, 0, 238, 239, 5, 114, 0, 0, 239, 240, 5, 101, 0, 0, 240, 241, 5, 116,
		0, 0, 241, 242, 5, 117, 0, 0, 242, 243, 5, 114, 0, 0, 243, 244, 5, 110,
		0, 0, 244, 30, 1, 0, 0, 0, 245, 246, 5, 98, 0, 0, 246, 247, 5, 114, 0,
		0, 247, 248, 5, 101, 0, 0, 248, 249, 5, 97, 0, 0, 249, 250, 5, 107, 0,
		0, 250, 32, 1, 0, 0, 0, 251, 252, 5, 99, 0, 0, 252, 253, 5, 111, 0, 0,
		253, 254, 5, 110, 0, 0, 254, 255, 5, 116, 0, 0, 255, 256, 5, 105, 0, 0,
		256, 257, 5, 110, 0, 0, 257, 258, 5, 117, 0, 0, 258, 259, 5, 101, 0, 0,
		259, 34, 1, 0, 0, 0, 260, 261, 5, 97, 0, 0, 261, 262, 5, 112, 0, 0, 262,
		263, 5, 112, 0, 0, 263, 264, 5, 101, 0, 0, 264, 265, 5, 110, 0, 0, 265,
		266, 5, 100, 0, 0, 266, 36, 1, 0, 0, 0, 267, 268, 5, 105, 0, 0, 268, 269,
		5, 102, 0, 0, 269, 38, 1, 0, 0, 0, 270, 271, 5, 101, 0, 0, 271, 272, 5,
		108, 0, 0, 272, 273, 5, 115, 0, 0, 273, 274, 5, 101, 0, 0, 274, 40, 1,
		0, 0, 0, 275, 276, 5, 102, 0, 0, 276, 277, 5, 111, 0, 0, 277, 278, 5, 114,
		0, 0, 278, 42, 1, 0, 0, 0, 279, 280, 5, 115, 0, 0, 280, 281, 5, 119, 0,
		0, 281, 282, 5, 105, 0, 0, 282, 283, 5, 116, 0, 0, 283, 284, 5, 99, 0,
		0, 284, 285, 5, 104, 0, 0, 285, 44, 1, 0, 0, 0, 286, 287, 5, 99, 0, 0,
		287, 288, 5, 97, 0, 0, 288, 289, 5, 115, 0, 0, 289, 290, 5, 101, 0, 0,
		290, 46, 1, 0, 0, 0, 291, 292, 5, 100, 0, 0, 292, 293, 5, 101, 0, 0, 293,
		294, 5, 102, 0, 0, 294, 295, 5, 97, 0, 0, 295, 296, 5, 117, 0, 0, 296,
		297, 5, 108, 0, 0, 297, 298, 5, 116, 0, 0, 298, 48, 1, 0, 0, 0, 299, 303,
		7, 3, 0, 0, 300, 302, 7, 4, 0, 0, 301, 300, 1, 0, 0, 0, 302, 305, 1, 0,
		0, 0, 303, 301, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 50, 1, 0, 0, 0,
		305, 303, 1, 0, 0, 0, 306, 307, 5, 43, 0, 0, 307, 52, 1, 0, 0, 0, 308,
		309, 5, 45, 0, 0, 309, 54, 1, 0, 0, 0, 310, 311, 5, 42, 0, 0, 311, 56,
		1, 0, 0, 0, 312, 313, 5, 47, 0, 0, 313, 58, 1, 0, 0, 0, 314, 315, 5, 37,
		0, 0, 315, 60, 1, 0, 0, 0, 316, 317, 5, 60, 0, 0, 317, 62, 1, 0, 0, 0,
		318, 319, 5, 60, 0, 0, 319, 320, 5, 61, 0, 0, 320, 64, 1, 0, 0, 0, 321,
		322, 5, 62, 0, 0, 322, 66, 1, 0, 0, 0, 323, 324, 5, 62, 0, 0, 324, 325,
		5, 61, 0, 0, 325, 68, 1, 0, 0, 0, 326, 327, 5, 61, 0, 0, 327, 328, 5, 61,
		0, 0, 328, 70, 1, 0, 0, 0, 329, 330, 5, 33, 0, 0, 330, 331, 5, 61, 0, 0,
		331, 72, 1, 0, 0, 0, 332, 333, 5, 38, 0, 0, 333, 334, 5, 38, 0, 0, 334,
		74, 1, 0, 0, 0, 335, 336, 5, 124, 0, 0, 336, 337, 5, 124, 0, 0, 337, 76,
		1, 0, 0, 0, 338, 339, 5, 33, 0, 0, 339, 78, 1, 0, 0, 0, 340, 341, 5, 38,
		0, 0, 341, 80, 1, 0, 0, 0, 342, 343, 5, 124, 0, 0, 343, 82, 1, 0, 0, 0,
		344, 345, 5, 94, 0, 0, 345, 84, 1, 0, 0, 0, 346, 347, 5, 38, 0, 0, 347,
		348, 5, 94, 0, 0, 348, 86, 1, 0, 0, 0, 349, 350, 5, 60, 0, 0, 350, 351,
		5, 60, 0, 0, 351, 88, 1, 0, 0, 0, 352, 353, 5, 62, 0, 0, 353, 354, 5, 62,
		0, 0, 354, 90, 1, 0, 0, 0, 355, 356, 5, 43, 0, 0, 356, 357, 5, 43, 0, 0,
		357, 92, 1, 0, 0, 0, 358, 359, 5, 45, 0, 0, 359, 360, 5, 45, 0, 0, 360,
		94, 1, 0, 0, 0, 361, 362, 5, 61, 0, 0, 362, 96, 1, 0, 0, 0, 363, 364, 5,
		43, 0, 0, 364, 365, 5, 61, 0, 0, 365, 98, 1, 0, 0, 0, 366, 367, 5, 45,
		0, 0, 367, 368, 5, 61, 0, 0, 368, 100, 1, 0, 0, 0, 369, 370, 5, 42, 0,
		0, 370, 371, 5, 61, 0, 0, 371, 102, 1, 0, 0, 0, 372, 373, 5, 47, 0, 0,
		373, 374, 5, 61, 0, 0, 374, 104, 1, 0, 0, 0, 375, 376, 5, 37, 0, 0, 376,
		377, 5, 61, 0, 0, 377, 106, 1, 0, 0, 0, 378, 379, 5, 38, 0, 0, 379, 380,
		5, 61, 0, 0, 380, 108, 1, 0, 0, 0, 381, 382, 5, 124, 0, 0, 382, 383, 5,
		61, 0, 0, 383, 110, 1, 0, 0, 0, 384, 385, 5, 94, 0, 0, 385, 386, 5, 61,
		0, 0, 386, 112, 1, 0, 0, 0, 387, 388, 5, 60, 0, 0, 388, 389, 5, 60, 0,
		0, 389, 390, 5, 61, 0, 0, 390, 114, 1, 0, 0, 0, 391, 392, 5, 62, 0, 0,
		392, 393, 5, 62, 0, 0, 393, 394, 5, 61, 0, 0, 394, 116, 1, 0, 0, 0, 395,
		396, 5, 38, 0, 0, 396, 397, 5, 94, 0, 0, 397, 398, 5, 61, 0, 0, 398, 118,
		1, 0, 0, 0, 399, 400, 5, 58, 0, 0, 400, 120, 1, 0, 0, 0, 401, 402, 5, 59,
		0, 0, 402, 122, 1, 0, 0, 0, 403, 404, 5, 44, 0, 0, 404, 124, 1, 0, 0, 0,
		405, 406, 5, 46, 0, 0, 406, 126, 1, 0, 0, 0, 407, 408, 5, 40, 0, 0, 408,
		128, 1, 0, 0, 0, 409, 410, 5, 41, 0, 0, 410, 130, 1, 0, 0, 0, 411, 412,
		5, 91, 0, 0, 412, 132, 1, 0, 0, 0, 413, 414, 5, 93, 0, 0, 414, 134, 1,
		0, 0, 0, 415, 416, 5, 123, 0, 0, 416, 136, 1, 0, 0, 0, 417, 418, 5, 125,
		0, 0, 418, 138, 1, 0, 0, 0, 419, 421, 7, 5, 0, 0, 420, 419, 1, 0, 0, 0,
		421, 422, 1, 0, 0, 0, 422, 420, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423,
		424, 1, 0, 0, 0, 424, 425, 6, 69, 0, 0, 425, 140, 1, 0, 0, 0, 426, 427,
		5, 47, 0, 0, 427, 428, 5, 47, 0, 0, 428, 432, 1, 0, 0, 0, 429, 431, 8,
		6, 0, 0, 430, 429, 1, 0, 0, 0, 431, 434, 1, 0, 0, 0, 432, 430, 1, 0, 0,
		0, 432, 433, 1, 0, 0, 0, 433, 435, 1, 0, 0, 0, 434, 432, 1, 0, 0, 0, 435,
		436, 6, 70, 0, 0, 436, 142, 1, 0, 0, 0, 437, 438, 5, 47, 0, 0, 438, 439,
		5, 42, 0, 0, 439, 443, 1, 0, 0, 0, 440, 442, 9, 0, 0, 0, 441, 440, 1, 0,
		0, 0, 442, 445, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0,
		444, 446, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 446, 447, 5, 42, 0, 0, 447,
		448, 5, 47, 0, 0, 448, 449, 1, 0, 0, 0, 449, 450, 6, 71, 0, 0, 450, 144,
		1, 0, 0, 0, 12, 0, 148, 153, 159, 169, 171, 180, 182, 303, 422, 432, 443,
		1, 6, 0, 0,
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

// MiniGoScannerInit initializes any static state used to implement MiniGoScanner. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMiniGoScanner(). You can call this function if you wish to initialize the static state ahead
// of time.
func MiniGoScannerInit() {
	staticData := &MiniGoScannerLexerStaticData
	staticData.once.Do(minigoscannerLexerInit)
}

// NewMiniGoScanner produces a new lexer instance for the optional input antlr.CharStream.
func NewMiniGoScanner(input antlr.CharStream) *MiniGoScanner {
	MiniGoScannerInit()
	l := new(MiniGoScanner)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MiniGoScannerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "MiniGoScanner.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MiniGoScanner tokens.
const (
	MiniGoScannerINTLITERAL               = 1
	MiniGoScannerFLOATLITERAL             = 2
	MiniGoScannerRUNELITERAL              = 3
	MiniGoScannerRAWSTRINGLITERAL         = 4
	MiniGoScannerINTERPRETEDSTRINGLITERAL = 5
	MiniGoScannerPACKAGE                  = 6
	MiniGoScannerVAR                      = 7
	MiniGoScannerTYPE                     = 8
	MiniGoScannerFUNC                     = 9
	MiniGoScannerSTRUCT                   = 10
	MiniGoScannerLEN                      = 11
	MiniGoScannerCAP                      = 12
	MiniGoScannerPRINT                    = 13
	MiniGoScannerPRINTLN                  = 14
	MiniGoScannerRETURN                   = 15
	MiniGoScannerBREAK                    = 16
	MiniGoScannerCONTINUE                 = 17
	MiniGoScannerAPPEND                   = 18
	MiniGoScannerIF                       = 19
	MiniGoScannerELSE                     = 20
	MiniGoScannerFOR                      = 21
	MiniGoScannerSWITCH                   = 22
	MiniGoScannerCASE                     = 23
	MiniGoScannerDEFAULT                  = 24
	MiniGoScannerIDENTIFIER               = 25
	MiniGoScannerPLUS                     = 26
	MiniGoScannerMINUS                    = 27
	MiniGoScannerMULTIPLY                 = 28
	MiniGoScannerDIVIDE                   = 29
	MiniGoScannerMODULO                   = 30
	MiniGoScannerLESS                     = 31
	MiniGoScannerLESSEQUAL                = 32
	MiniGoScannerGREATER                  = 33
	MiniGoScannerGREATEREQUAL             = 34
	MiniGoScannerEQUAL                    = 35
	MiniGoScannerNOTEQUAL                 = 36
	MiniGoScannerAND                      = 37
	MiniGoScannerOR                       = 38
	MiniGoScannerNOT                      = 39
	MiniGoScannerBITWISEAND               = 40
	MiniGoScannerBITWISEOR                = 41
	MiniGoScannerBITWISEXOR               = 42
	MiniGoScannerBITWISECLEAR             = 43
	MiniGoScannerSHIFTLEFT                = 44
	MiniGoScannerSHIFTRIGHT               = 45
	MiniGoScannerINCREMENT                = 46
	MiniGoScannerDECREMENT                = 47
	MiniGoScannerASSIGN                   = 48
	MiniGoScannerPLUSEQUAL                = 49
	MiniGoScannerMINUSEQUAL               = 50
	MiniGoScannerMULTIPLYEQUAL            = 51
	MiniGoScannerDIVIDEEQUAL              = 52
	MiniGoScannerMODULOEQUAL              = 53
	MiniGoScannerBITWISEANDEQUAL          = 54
	MiniGoScannerBITWISEOREQUAL           = 55
	MiniGoScannerBITWISEXOREQUAL          = 56
	MiniGoScannerSHIFTLEFTEQUAL           = 57
	MiniGoScannerSHIFTRIGHTEQUAL          = 58
	MiniGoScannerBITWISECLEAREQUAL        = 59
	MiniGoScannerCOLON                    = 60
	MiniGoScannerSEMICOLON                = 61
	MiniGoScannerCOMMA                    = 62
	MiniGoScannerDOT                      = 63
	MiniGoScannerLPAREN                   = 64
	MiniGoScannerRPAREN                   = 65
	MiniGoScannerLBRACKET                 = 66
	MiniGoScannerRBRACKET                 = 67
	MiniGoScannerLBRACE                   = 68
	MiniGoScannerRBRACE                   = 69
	MiniGoScannerSPACES                   = 70
	MiniGoScannerLINE_COMMENT             = 71
	MiniGoScannerBLOCK_COMMENT            = 72
)
