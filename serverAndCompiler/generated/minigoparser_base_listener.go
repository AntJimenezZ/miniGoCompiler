// Code generated from C:/Users/noni4/Desktop/miniGoProyect/serverAndCompiler/miniGoParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package generated // miniGoParser
import "github.com/antlr4-go/antlr/v4"

// BaseminiGoParserListener is a complete listener for a parse tree produced by miniGoParser.
type BaseminiGoParserListener struct{}

var _ miniGoParserListener = &BaseminiGoParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseminiGoParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseminiGoParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseminiGoParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseminiGoParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRootAST is called when production rootAST is entered.
func (s *BaseminiGoParserListener) EnterRootAST(ctx *RootASTContext) {}

// ExitRootAST is called when production rootAST is exited.
func (s *BaseminiGoParserListener) ExitRootAST(ctx *RootASTContext) {}

// EnterTopDeclarationListAST is called when production topDeclarationListAST is entered.
func (s *BaseminiGoParserListener) EnterTopDeclarationListAST(ctx *TopDeclarationListASTContext) {}

// ExitTopDeclarationListAST is called when production topDeclarationListAST is exited.
func (s *BaseminiGoParserListener) ExitTopDeclarationListAST(ctx *TopDeclarationListASTContext) {}

// EnterVariableDeclAST is called when production variableDeclAST is entered.
func (s *BaseminiGoParserListener) EnterVariableDeclAST(ctx *VariableDeclASTContext) {}

// ExitVariableDeclAST is called when production variableDeclAST is exited.
func (s *BaseminiGoParserListener) ExitVariableDeclAST(ctx *VariableDeclASTContext) {}

// EnterVariableDeclBlockAST is called when production variableDeclBlockAST is entered.
func (s *BaseminiGoParserListener) EnterVariableDeclBlockAST(ctx *VariableDeclBlockASTContext) {}

// ExitVariableDeclBlockAST is called when production variableDeclBlockAST is exited.
func (s *BaseminiGoParserListener) ExitVariableDeclBlockAST(ctx *VariableDeclBlockASTContext) {}

// EnterVariableDeclEmptyAST is called when production variableDeclEmptyAST is entered.
func (s *BaseminiGoParserListener) EnterVariableDeclEmptyAST(ctx *VariableDeclEmptyASTContext) {}

// ExitVariableDeclEmptyAST is called when production variableDeclEmptyAST is exited.
func (s *BaseminiGoParserListener) ExitVariableDeclEmptyAST(ctx *VariableDeclEmptyASTContext) {}

// EnterInnerVarDeclsAST is called when production innerVarDeclsAST is entered.
func (s *BaseminiGoParserListener) EnterInnerVarDeclsAST(ctx *InnerVarDeclsASTContext) {}

// ExitInnerVarDeclsAST is called when production innerVarDeclsAST is exited.
func (s *BaseminiGoParserListener) ExitInnerVarDeclsAST(ctx *InnerVarDeclsASTContext) {}

// EnterSingleVarDeclAST is called when production singleVarDeclAST is entered.
func (s *BaseminiGoParserListener) EnterSingleVarDeclAST(ctx *SingleVarDeclASTContext) {}

// ExitSingleVarDeclAST is called when production singleVarDeclAST is exited.
func (s *BaseminiGoParserListener) ExitSingleVarDeclAST(ctx *SingleVarDeclASTContext) {}

// EnterSingleVarDeclAssignAST is called when production singleVarDeclAssignAST is entered.
func (s *BaseminiGoParserListener) EnterSingleVarDeclAssignAST(ctx *SingleVarDeclAssignASTContext) {}

// ExitSingleVarDeclAssignAST is called when production singleVarDeclAssignAST is exited.
func (s *BaseminiGoParserListener) ExitSingleVarDeclAssignAST(ctx *SingleVarDeclAssignASTContext) {}

// EnterSingleVarDeclNoExpsAST is called when production singleVarDeclNoExpsAST is entered.
func (s *BaseminiGoParserListener) EnterSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) {}

// ExitSingleVarDeclNoExpsAST is called when production singleVarDeclNoExpsAST is exited.
func (s *BaseminiGoParserListener) ExitSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) {}

// EnterSingleVarDeclNoExpsDeclTypeAST is called when production singleVarDeclNoExpsDeclTypeAST is entered.
func (s *BaseminiGoParserListener) EnterSingleVarDeclNoExpsDeclTypeAST(ctx *SingleVarDeclNoExpsDeclTypeASTContext) {
}

// ExitSingleVarDeclNoExpsDeclTypeAST is called when production singleVarDeclNoExpsDeclTypeAST is exited.
func (s *BaseminiGoParserListener) ExitSingleVarDeclNoExpsDeclTypeAST(ctx *SingleVarDeclNoExpsDeclTypeASTContext) {
}

// EnterTypeDeclAST is called when production typeDeclAST is entered.
func (s *BaseminiGoParserListener) EnterTypeDeclAST(ctx *TypeDeclASTContext) {}

// ExitTypeDeclAST is called when production typeDeclAST is exited.
func (s *BaseminiGoParserListener) ExitTypeDeclAST(ctx *TypeDeclASTContext) {}

// EnterTypeDeclBlockAST is called when production typeDeclBlockAST is entered.
func (s *BaseminiGoParserListener) EnterTypeDeclBlockAST(ctx *TypeDeclBlockASTContext) {}

// ExitTypeDeclBlockAST is called when production typeDeclBlockAST is exited.
func (s *BaseminiGoParserListener) ExitTypeDeclBlockAST(ctx *TypeDeclBlockASTContext) {}

// EnterTypeDeclEmptyAST is called when production typeDeclEmptyAST is entered.
func (s *BaseminiGoParserListener) EnterTypeDeclEmptyAST(ctx *TypeDeclEmptyASTContext) {}

// ExitTypeDeclEmptyAST is called when production typeDeclEmptyAST is exited.
func (s *BaseminiGoParserListener) ExitTypeDeclEmptyAST(ctx *TypeDeclEmptyASTContext) {}

// EnterInnerTypeDeclsAST is called when production innerTypeDeclsAST is entered.
func (s *BaseminiGoParserListener) EnterInnerTypeDeclsAST(ctx *InnerTypeDeclsASTContext) {}

// ExitInnerTypeDeclsAST is called when production innerTypeDeclsAST is exited.
func (s *BaseminiGoParserListener) ExitInnerTypeDeclsAST(ctx *InnerTypeDeclsASTContext) {}

// EnterSingleTypeDeclAST is called when production singleTypeDeclAST is entered.
func (s *BaseminiGoParserListener) EnterSingleTypeDeclAST(ctx *SingleTypeDeclASTContext) {}

// ExitSingleTypeDeclAST is called when production singleTypeDeclAST is exited.
func (s *BaseminiGoParserListener) ExitSingleTypeDeclAST(ctx *SingleTypeDeclASTContext) {}

// EnterFuncDeclAST is called when production funcDeclAST is entered.
func (s *BaseminiGoParserListener) EnterFuncDeclAST(ctx *FuncDeclASTContext) {}

// ExitFuncDeclAST is called when production funcDeclAST is exited.
func (s *BaseminiGoParserListener) ExitFuncDeclAST(ctx *FuncDeclASTContext) {}

// EnterFuncFrontDeclAST is called when production funcFrontDeclAST is entered.
func (s *BaseminiGoParserListener) EnterFuncFrontDeclAST(ctx *FuncFrontDeclASTContext) {}

// ExitFuncFrontDeclAST is called when production funcFrontDeclAST is exited.
func (s *BaseminiGoParserListener) ExitFuncFrontDeclAST(ctx *FuncFrontDeclASTContext) {}

// EnterFuncArgDeclsAST is called when production funcArgDeclsAST is entered.
func (s *BaseminiGoParserListener) EnterFuncArgDeclsAST(ctx *FuncArgDeclsASTContext) {}

// ExitFuncArgDeclsAST is called when production funcArgDeclsAST is exited.
func (s *BaseminiGoParserListener) ExitFuncArgDeclsAST(ctx *FuncArgDeclsASTContext) {}

// EnterDeclTypeParenAST is called when production declTypeParenAST is entered.
func (s *BaseminiGoParserListener) EnterDeclTypeParenAST(ctx *DeclTypeParenASTContext) {}

// ExitDeclTypeParenAST is called when production declTypeParenAST is exited.
func (s *BaseminiGoParserListener) ExitDeclTypeParenAST(ctx *DeclTypeParenASTContext) {}

// EnterDeclTypeIdentifierAST is called when production declTypeIdentifierAST is entered.
func (s *BaseminiGoParserListener) EnterDeclTypeIdentifierAST(ctx *DeclTypeIdentifierASTContext) {}

// ExitDeclTypeIdentifierAST is called when production declTypeIdentifierAST is exited.
func (s *BaseminiGoParserListener) ExitDeclTypeIdentifierAST(ctx *DeclTypeIdentifierASTContext) {}

// EnterDeclTypeSliceAST is called when production declTypeSliceAST is entered.
func (s *BaseminiGoParserListener) EnterDeclTypeSliceAST(ctx *DeclTypeSliceASTContext) {}

// ExitDeclTypeSliceAST is called when production declTypeSliceAST is exited.
func (s *BaseminiGoParserListener) ExitDeclTypeSliceAST(ctx *DeclTypeSliceASTContext) {}

// EnterDeclTypeArrayAST is called when production declTypeArrayAST is entered.
func (s *BaseminiGoParserListener) EnterDeclTypeArrayAST(ctx *DeclTypeArrayASTContext) {}

// ExitDeclTypeArrayAST is called when production declTypeArrayAST is exited.
func (s *BaseminiGoParserListener) ExitDeclTypeArrayAST(ctx *DeclTypeArrayASTContext) {}

// EnterDeclTypeStructAST is called when production declTypeStructAST is entered.
func (s *BaseminiGoParserListener) EnterDeclTypeStructAST(ctx *DeclTypeStructASTContext) {}

// ExitDeclTypeStructAST is called when production declTypeStructAST is exited.
func (s *BaseminiGoParserListener) ExitDeclTypeStructAST(ctx *DeclTypeStructASTContext) {}

// EnterSliceDeclTypeAST is called when production sliceDeclTypeAST is entered.
func (s *BaseminiGoParserListener) EnterSliceDeclTypeAST(ctx *SliceDeclTypeASTContext) {}

// ExitSliceDeclTypeAST is called when production sliceDeclTypeAST is exited.
func (s *BaseminiGoParserListener) ExitSliceDeclTypeAST(ctx *SliceDeclTypeASTContext) {}

// EnterArrayDeclTypeAST is called when production arrayDeclTypeAST is entered.
func (s *BaseminiGoParserListener) EnterArrayDeclTypeAST(ctx *ArrayDeclTypeASTContext) {}

// ExitArrayDeclTypeAST is called when production arrayDeclTypeAST is exited.
func (s *BaseminiGoParserListener) ExitArrayDeclTypeAST(ctx *ArrayDeclTypeASTContext) {}

// EnterStructDeclTypeAST is called when production structDeclTypeAST is entered.
func (s *BaseminiGoParserListener) EnterStructDeclTypeAST(ctx *StructDeclTypeASTContext) {}

// ExitStructDeclTypeAST is called when production structDeclTypeAST is exited.
func (s *BaseminiGoParserListener) ExitStructDeclTypeAST(ctx *StructDeclTypeASTContext) {}

// EnterStructMemDeclsAST is called when production structMemDeclsAST is entered.
func (s *BaseminiGoParserListener) EnterStructMemDeclsAST(ctx *StructMemDeclsASTContext) {}

// ExitStructMemDeclsAST is called when production structMemDeclsAST is exited.
func (s *BaseminiGoParserListener) ExitStructMemDeclsAST(ctx *StructMemDeclsASTContext) {}

// EnterIdentifierListAST is called when production identifierListAST is entered.
func (s *BaseminiGoParserListener) EnterIdentifierListAST(ctx *IdentifierListASTContext) {}

// ExitIdentifierListAST is called when production identifierListAST is exited.
func (s *BaseminiGoParserListener) ExitIdentifierListAST(ctx *IdentifierListASTContext) {}

// EnterExpressionNotUnaryAST is called when production expressionNotUnaryAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionNotUnaryAST(ctx *ExpressionNotUnaryASTContext) {}

// ExitExpressionNotUnaryAST is called when production expressionNotUnaryAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionNotUnaryAST(ctx *ExpressionNotUnaryASTContext) {}

// EnterExpressionMultiplyAST is called when production expressionMultiplyAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionMultiplyAST(ctx *ExpressionMultiplyASTContext) {}

// ExitExpressionMultiplyAST is called when production expressionMultiplyAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionMultiplyAST(ctx *ExpressionMultiplyASTContext) {}

// EnterExpressionPlusAST is called when production expressionPlusAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionPlusAST(ctx *ExpressionPlusASTContext) {}

// ExitExpressionPlusAST is called when production expressionPlusAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionPlusAST(ctx *ExpressionPlusASTContext) {}

// EnterExpressionModuloAST is called when production expressionModuloAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionModuloAST(ctx *ExpressionModuloASTContext) {}

// ExitExpressionModuloAST is called when production expressionModuloAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionModuloAST(ctx *ExpressionModuloASTContext) {}

// EnterExpressionAndAST is called when production expressionAndAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionAndAST(ctx *ExpressionAndASTContext) {}

// ExitExpressionAndAST is called when production expressionAndAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionAndAST(ctx *ExpressionAndASTContext) {}

// EnterExpressionBitwiseXorAST is called when production expressionBitwiseXorAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionBitwiseXorAST(ctx *ExpressionBitwiseXorASTContext) {
}

// ExitExpressionBitwiseXorAST is called when production expressionBitwiseXorAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionBitwiseXorAST(ctx *ExpressionBitwiseXorASTContext) {}

// EnterExpressionMinusAST is called when production expressionMinusAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionMinusAST(ctx *ExpressionMinusASTContext) {}

// ExitExpressionMinusAST is called when production expressionMinusAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionMinusAST(ctx *ExpressionMinusASTContext) {}

// EnterExpressionBitwiseXorUnaryAST is called when production expressionBitwiseXorUnaryAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionBitwiseXorUnaryAST(ctx *ExpressionBitwiseXorUnaryASTContext) {
}

// ExitExpressionBitwiseXorUnaryAST is called when production expressionBitwiseXorUnaryAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionBitwiseXorUnaryAST(ctx *ExpressionBitwiseXorUnaryASTContext) {
}

// EnterExpressionEqualAST is called when production expressionEqualAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionEqualAST(ctx *ExpressionEqualASTContext) {}

// ExitExpressionEqualAST is called when production expressionEqualAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionEqualAST(ctx *ExpressionEqualASTContext) {}

// EnterExpressionMinusUnaryAST is called when production expressionMinusUnaryAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionMinusUnaryAST(ctx *ExpressionMinusUnaryASTContext) {
}

// ExitExpressionMinusUnaryAST is called when production expressionMinusUnaryAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionMinusUnaryAST(ctx *ExpressionMinusUnaryASTContext) {}

// EnterExpressionBitwiseClearAST is called when production expressionBitwiseClearAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionBitwiseClearAST(ctx *ExpressionBitwiseClearASTContext) {
}

// ExitExpressionBitwiseClearAST is called when production expressionBitwiseClearAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionBitwiseClearAST(ctx *ExpressionBitwiseClearASTContext) {
}

// EnterExpressionGreaterEqualAST is called when production expressionGreaterEqualAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionGreaterEqualAST(ctx *ExpressionGreaterEqualASTContext) {
}

// ExitExpressionGreaterEqualAST is called when production expressionGreaterEqualAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionGreaterEqualAST(ctx *ExpressionGreaterEqualASTContext) {
}

// EnterExpressionLessEqualAST is called when production expressionLessEqualAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionLessEqualAST(ctx *ExpressionLessEqualASTContext) {}

// ExitExpressionLessEqualAST is called when production expressionLessEqualAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionLessEqualAST(ctx *ExpressionLessEqualASTContext) {}

// EnterExpressionNotEqualAST is called when production expressionNotEqualAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionNotEqualAST(ctx *ExpressionNotEqualASTContext) {}

// ExitExpressionNotEqualAST is called when production expressionNotEqualAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionNotEqualAST(ctx *ExpressionNotEqualASTContext) {}

// EnterExpressionPrimaryAST is called when production expressionPrimaryAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionPrimaryAST(ctx *ExpressionPrimaryASTContext) {}

// ExitExpressionPrimaryAST is called when production expressionPrimaryAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionPrimaryAST(ctx *ExpressionPrimaryASTContext) {}

// EnterExpressionBitwiseAndAST is called when production expressionBitwiseAndAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionBitwiseAndAST(ctx *ExpressionBitwiseAndASTContext) {
}

// ExitExpressionBitwiseAndAST is called when production expressionBitwiseAndAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionBitwiseAndAST(ctx *ExpressionBitwiseAndASTContext) {}

// EnterExpressionGreaterAST is called when production expressionGreaterAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionGreaterAST(ctx *ExpressionGreaterASTContext) {}

// ExitExpressionGreaterAST is called when production expressionGreaterAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionGreaterAST(ctx *ExpressionGreaterASTContext) {}

// EnterExpressionDivideAST is called when production expressionDivideAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionDivideAST(ctx *ExpressionDivideASTContext) {}

// ExitExpressionDivideAST is called when production expressionDivideAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionDivideAST(ctx *ExpressionDivideASTContext) {}

// EnterExpressionOrAST is called when production expressionOrAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionOrAST(ctx *ExpressionOrASTContext) {}

// ExitExpressionOrAST is called when production expressionOrAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionOrAST(ctx *ExpressionOrASTContext) {}

// EnterExpressionPlusUnaryAST is called when production expressionPlusUnaryAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionPlusUnaryAST(ctx *ExpressionPlusUnaryASTContext) {}

// ExitExpressionPlusUnaryAST is called when production expressionPlusUnaryAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionPlusUnaryAST(ctx *ExpressionPlusUnaryASTContext) {}

// EnterExpressionBitwiseOrAST is called when production expressionBitwiseOrAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionBitwiseOrAST(ctx *ExpressionBitwiseOrASTContext) {}

// ExitExpressionBitwiseOrAST is called when production expressionBitwiseOrAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionBitwiseOrAST(ctx *ExpressionBitwiseOrASTContext) {}

// EnterExpressionLessAST is called when production expressionLessAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionLessAST(ctx *ExpressionLessASTContext) {}

// ExitExpressionLessAST is called when production expressionLessAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionLessAST(ctx *ExpressionLessASTContext) {}

// EnterExpressionShiftRightAST is called when production expressionShiftRightAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionShiftRightAST(ctx *ExpressionShiftRightASTContext) {
}

// ExitExpressionShiftRightAST is called when production expressionShiftRightAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionShiftRightAST(ctx *ExpressionShiftRightASTContext) {}

// EnterExpressionShiftLeftAST is called when production expressionShiftLeftAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionShiftLeftAST(ctx *ExpressionShiftLeftASTContext) {}

// ExitExpressionShiftLeftAST is called when production expressionShiftLeftAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionShiftLeftAST(ctx *ExpressionShiftLeftASTContext) {}

// EnterExpressionListAST is called when production expressionListAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionListAST(ctx *ExpressionListASTContext) {}

// ExitExpressionListAST is called when production expressionListAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionListAST(ctx *ExpressionListASTContext) {}

// EnterPrimaryExpressionLengthAST is called when production primaryExpressionLengthAST is entered.
func (s *BaseminiGoParserListener) EnterPrimaryExpressionLengthAST(ctx *PrimaryExpressionLengthASTContext) {
}

// ExitPrimaryExpressionLengthAST is called when production primaryExpressionLengthAST is exited.
func (s *BaseminiGoParserListener) ExitPrimaryExpressionLengthAST(ctx *PrimaryExpressionLengthASTContext) {
}

// EnterPrimaryExpressionOperandAST is called when production primaryExpressionOperandAST is entered.
func (s *BaseminiGoParserListener) EnterPrimaryExpressionOperandAST(ctx *PrimaryExpressionOperandASTContext) {
}

// ExitPrimaryExpressionOperandAST is called when production primaryExpressionOperandAST is exited.
func (s *BaseminiGoParserListener) ExitPrimaryExpressionOperandAST(ctx *PrimaryExpressionOperandASTContext) {
}

// EnterPrimaryExpressionIndexAST is called when production primaryExpressionIndexAST is entered.
func (s *BaseminiGoParserListener) EnterPrimaryExpressionIndexAST(ctx *PrimaryExpressionIndexASTContext) {
}

// ExitPrimaryExpressionIndexAST is called when production primaryExpressionIndexAST is exited.
func (s *BaseminiGoParserListener) ExitPrimaryExpressionIndexAST(ctx *PrimaryExpressionIndexASTContext) {
}

// EnterPrimaryExpressionAppendAST is called when production primaryExpressionAppendAST is entered.
func (s *BaseminiGoParserListener) EnterPrimaryExpressionAppendAST(ctx *PrimaryExpressionAppendASTContext) {
}

// ExitPrimaryExpressionAppendAST is called when production primaryExpressionAppendAST is exited.
func (s *BaseminiGoParserListener) ExitPrimaryExpressionAppendAST(ctx *PrimaryExpressionAppendASTContext) {
}

// EnterPrimaryExpressionArgumentsAST is called when production primaryExpressionArgumentsAST is entered.
func (s *BaseminiGoParserListener) EnterPrimaryExpressionArgumentsAST(ctx *PrimaryExpressionArgumentsASTContext) {
}

// ExitPrimaryExpressionArgumentsAST is called when production primaryExpressionArgumentsAST is exited.
func (s *BaseminiGoParserListener) ExitPrimaryExpressionArgumentsAST(ctx *PrimaryExpressionArgumentsASTContext) {
}

// EnterPrimaryExpressionCapAST is called when production primaryExpressionCapAST is entered.
func (s *BaseminiGoParserListener) EnterPrimaryExpressionCapAST(ctx *PrimaryExpressionCapASTContext) {
}

// ExitPrimaryExpressionCapAST is called when production primaryExpressionCapAST is exited.
func (s *BaseminiGoParserListener) ExitPrimaryExpressionCapAST(ctx *PrimaryExpressionCapASTContext) {}

// EnterPrimaryExpressionSelectorAST is called when production primaryExpressionSelectorAST is entered.
func (s *BaseminiGoParserListener) EnterPrimaryExpressionSelectorAST(ctx *PrimaryExpressionSelectorASTContext) {
}

// ExitPrimaryExpressionSelectorAST is called when production primaryExpressionSelectorAST is exited.
func (s *BaseminiGoParserListener) ExitPrimaryExpressionSelectorAST(ctx *PrimaryExpressionSelectorASTContext) {
}

// EnterOperandLiteralAST is called when production operandLiteralAST is entered.
func (s *BaseminiGoParserListener) EnterOperandLiteralAST(ctx *OperandLiteralASTContext) {}

// ExitOperandLiteralAST is called when production operandLiteralAST is exited.
func (s *BaseminiGoParserListener) ExitOperandLiteralAST(ctx *OperandLiteralASTContext) {}

// EnterOperandIdentifierAST is called when production operandIdentifierAST is entered.
func (s *BaseminiGoParserListener) EnterOperandIdentifierAST(ctx *OperandIdentifierASTContext) {}

// ExitOperandIdentifierAST is called when production operandIdentifierAST is exited.
func (s *BaseminiGoParserListener) ExitOperandIdentifierAST(ctx *OperandIdentifierASTContext) {}

// EnterOperandParenAST is called when production operandParenAST is entered.
func (s *BaseminiGoParserListener) EnterOperandParenAST(ctx *OperandParenASTContext) {}

// ExitOperandParenAST is called when production operandParenAST is exited.
func (s *BaseminiGoParserListener) ExitOperandParenAST(ctx *OperandParenASTContext) {}

// EnterLiteralIntAST is called when production literalIntAST is entered.
func (s *BaseminiGoParserListener) EnterLiteralIntAST(ctx *LiteralIntASTContext) {}

// ExitLiteralIntAST is called when production literalIntAST is exited.
func (s *BaseminiGoParserListener) ExitLiteralIntAST(ctx *LiteralIntASTContext) {}

// EnterLiteralFloatAST is called when production literalFloatAST is entered.
func (s *BaseminiGoParserListener) EnterLiteralFloatAST(ctx *LiteralFloatASTContext) {}

// ExitLiteralFloatAST is called when production literalFloatAST is exited.
func (s *BaseminiGoParserListener) ExitLiteralFloatAST(ctx *LiteralFloatASTContext) {}

// EnterLiteralRuneAST is called when production literalRuneAST is entered.
func (s *BaseminiGoParserListener) EnterLiteralRuneAST(ctx *LiteralRuneASTContext) {}

// ExitLiteralRuneAST is called when production literalRuneAST is exited.
func (s *BaseminiGoParserListener) ExitLiteralRuneAST(ctx *LiteralRuneASTContext) {}

// EnterLiteralRawStringAST is called when production literalRawStringAST is entered.
func (s *BaseminiGoParserListener) EnterLiteralRawStringAST(ctx *LiteralRawStringASTContext) {}

// ExitLiteralRawStringAST is called when production literalRawStringAST is exited.
func (s *BaseminiGoParserListener) ExitLiteralRawStringAST(ctx *LiteralRawStringASTContext) {}

// EnterLiteralInterpretedStringAST is called when production literalInterpretedStringAST is entered.
func (s *BaseminiGoParserListener) EnterLiteralInterpretedStringAST(ctx *LiteralInterpretedStringASTContext) {
}

// ExitLiteralInterpretedStringAST is called when production literalInterpretedStringAST is exited.
func (s *BaseminiGoParserListener) ExitLiteralInterpretedStringAST(ctx *LiteralInterpretedStringASTContext) {
}

// EnterIndexAST is called when production indexAST is entered.
func (s *BaseminiGoParserListener) EnterIndexAST(ctx *IndexASTContext) {}

// ExitIndexAST is called when production indexAST is exited.
func (s *BaseminiGoParserListener) ExitIndexAST(ctx *IndexASTContext) {}

// EnterArgumentsAST is called when production argumentsAST is entered.
func (s *BaseminiGoParserListener) EnterArgumentsAST(ctx *ArgumentsASTContext) {}

// ExitArgumentsAST is called when production argumentsAST is exited.
func (s *BaseminiGoParserListener) ExitArgumentsAST(ctx *ArgumentsASTContext) {}

// EnterArgumentsEmptyAST is called when production argumentsEmptyAST is entered.
func (s *BaseminiGoParserListener) EnterArgumentsEmptyAST(ctx *ArgumentsEmptyASTContext) {}

// ExitArgumentsEmptyAST is called when production argumentsEmptyAST is exited.
func (s *BaseminiGoParserListener) ExitArgumentsEmptyAST(ctx *ArgumentsEmptyASTContext) {}

// EnterSelectorAST is called when production selectorAST is entered.
func (s *BaseminiGoParserListener) EnterSelectorAST(ctx *SelectorASTContext) {}

// ExitSelectorAST is called when production selectorAST is exited.
func (s *BaseminiGoParserListener) ExitSelectorAST(ctx *SelectorASTContext) {}

// EnterAppendExpressionAST is called when production appendExpressionAST is entered.
func (s *BaseminiGoParserListener) EnterAppendExpressionAST(ctx *AppendExpressionASTContext) {}

// ExitAppendExpressionAST is called when production appendExpressionAST is exited.
func (s *BaseminiGoParserListener) ExitAppendExpressionAST(ctx *AppendExpressionASTContext) {}

// EnterLengthExpressionAST is called when production lengthExpressionAST is entered.
func (s *BaseminiGoParserListener) EnterLengthExpressionAST(ctx *LengthExpressionASTContext) {}

// ExitLengthExpressionAST is called when production lengthExpressionAST is exited.
func (s *BaseminiGoParserListener) ExitLengthExpressionAST(ctx *LengthExpressionASTContext) {}

// EnterCapExpressionAST is called when production capExpressionAST is entered.
func (s *BaseminiGoParserListener) EnterCapExpressionAST(ctx *CapExpressionASTContext) {}

// ExitCapExpressionAST is called when production capExpressionAST is exited.
func (s *BaseminiGoParserListener) ExitCapExpressionAST(ctx *CapExpressionASTContext) {}

// EnterStatementListAST is called when production statementListAST is entered.
func (s *BaseminiGoParserListener) EnterStatementListAST(ctx *StatementListASTContext) {}

// ExitStatementListAST is called when production statementListAST is exited.
func (s *BaseminiGoParserListener) ExitStatementListAST(ctx *StatementListASTContext) {}

// EnterBlockAST is called when production blockAST is entered.
func (s *BaseminiGoParserListener) EnterBlockAST(ctx *BlockASTContext) {}

// ExitBlockAST is called when production blockAST is exited.
func (s *BaseminiGoParserListener) ExitBlockAST(ctx *BlockASTContext) {}

// EnterStatementPrintAST is called when production statementPrintAST is entered.
func (s *BaseminiGoParserListener) EnterStatementPrintAST(ctx *StatementPrintASTContext) {}

// ExitStatementPrintAST is called when production statementPrintAST is exited.
func (s *BaseminiGoParserListener) ExitStatementPrintAST(ctx *StatementPrintASTContext) {}

// EnterStatementPrintlnAST is called when production statementPrintlnAST is entered.
func (s *BaseminiGoParserListener) EnterStatementPrintlnAST(ctx *StatementPrintlnASTContext) {}

// ExitStatementPrintlnAST is called when production statementPrintlnAST is exited.
func (s *BaseminiGoParserListener) ExitStatementPrintlnAST(ctx *StatementPrintlnASTContext) {}

// EnterStatementReturnAST is called when production statementReturnAST is entered.
func (s *BaseminiGoParserListener) EnterStatementReturnAST(ctx *StatementReturnASTContext) {}

// ExitStatementReturnAST is called when production statementReturnAST is exited.
func (s *BaseminiGoParserListener) ExitStatementReturnAST(ctx *StatementReturnASTContext) {}

// EnterStatementBreakAST is called when production statementBreakAST is entered.
func (s *BaseminiGoParserListener) EnterStatementBreakAST(ctx *StatementBreakASTContext) {}

// ExitStatementBreakAST is called when production statementBreakAST is exited.
func (s *BaseminiGoParserListener) ExitStatementBreakAST(ctx *StatementBreakASTContext) {}

// EnterStatementContinueAST is called when production statementContinueAST is entered.
func (s *BaseminiGoParserListener) EnterStatementContinueAST(ctx *StatementContinueASTContext) {}

// ExitStatementContinueAST is called when production statementContinueAST is exited.
func (s *BaseminiGoParserListener) ExitStatementContinueAST(ctx *StatementContinueASTContext) {}

// EnterStatementSimpleAST is called when production statementSimpleAST is entered.
func (s *BaseminiGoParserListener) EnterStatementSimpleAST(ctx *StatementSimpleASTContext) {}

// ExitStatementSimpleAST is called when production statementSimpleAST is exited.
func (s *BaseminiGoParserListener) ExitStatementSimpleAST(ctx *StatementSimpleASTContext) {}

// EnterStatementBlockAST is called when production statementBlockAST is entered.
func (s *BaseminiGoParserListener) EnterStatementBlockAST(ctx *StatementBlockASTContext) {}

// ExitStatementBlockAST is called when production statementBlockAST is exited.
func (s *BaseminiGoParserListener) ExitStatementBlockAST(ctx *StatementBlockASTContext) {}

// EnterStatementSwitchAST is called when production statementSwitchAST is entered.
func (s *BaseminiGoParserListener) EnterStatementSwitchAST(ctx *StatementSwitchASTContext) {}

// ExitStatementSwitchAST is called when production statementSwitchAST is exited.
func (s *BaseminiGoParserListener) ExitStatementSwitchAST(ctx *StatementSwitchASTContext) {}

// EnterStatementIfAST is called when production statementIfAST is entered.
func (s *BaseminiGoParserListener) EnterStatementIfAST(ctx *StatementIfASTContext) {}

// ExitStatementIfAST is called when production statementIfAST is exited.
func (s *BaseminiGoParserListener) ExitStatementIfAST(ctx *StatementIfASTContext) {}

// EnterStatementLoopAST is called when production statementLoopAST is entered.
func (s *BaseminiGoParserListener) EnterStatementLoopAST(ctx *StatementLoopASTContext) {}

// ExitStatementLoopAST is called when production statementLoopAST is exited.
func (s *BaseminiGoParserListener) ExitStatementLoopAST(ctx *StatementLoopASTContext) {}

// EnterStatementTypeDeclAST is called when production statementTypeDeclAST is entered.
func (s *BaseminiGoParserListener) EnterStatementTypeDeclAST(ctx *StatementTypeDeclASTContext) {}

// ExitStatementTypeDeclAST is called when production statementTypeDeclAST is exited.
func (s *BaseminiGoParserListener) ExitStatementTypeDeclAST(ctx *StatementTypeDeclASTContext) {}

// EnterStatementVariableDeclAST is called when production statementVariableDeclAST is entered.
func (s *BaseminiGoParserListener) EnterStatementVariableDeclAST(ctx *StatementVariableDeclASTContext) {
}

// ExitStatementVariableDeclAST is called when production statementVariableDeclAST is exited.
func (s *BaseminiGoParserListener) ExitStatementVariableDeclAST(ctx *StatementVariableDeclASTContext) {
}

// EnterSimpleStatementEmptyAST is called when production simpleStatementEmptyAST is entered.
func (s *BaseminiGoParserListener) EnterSimpleStatementEmptyAST(ctx *SimpleStatementEmptyASTContext) {
}

// ExitSimpleStatementEmptyAST is called when production simpleStatementEmptyAST is exited.
func (s *BaseminiGoParserListener) ExitSimpleStatementEmptyAST(ctx *SimpleStatementEmptyASTContext) {}

// EnterSimpleStatementExpressionAST is called when production simpleStatementExpressionAST is entered.
func (s *BaseminiGoParserListener) EnterSimpleStatementExpressionAST(ctx *SimpleStatementExpressionASTContext) {
}

// ExitSimpleStatementExpressionAST is called when production simpleStatementExpressionAST is exited.
func (s *BaseminiGoParserListener) ExitSimpleStatementExpressionAST(ctx *SimpleStatementExpressionASTContext) {
}

// EnterSimpleStatementAssignmentAST is called when production simpleStatementAssignmentAST is entered.
func (s *BaseminiGoParserListener) EnterSimpleStatementAssignmentAST(ctx *SimpleStatementAssignmentASTContext) {
}

// ExitSimpleStatementAssignmentAST is called when production simpleStatementAssignmentAST is exited.
func (s *BaseminiGoParserListener) ExitSimpleStatementAssignmentAST(ctx *SimpleStatementAssignmentASTContext) {
}

// EnterSimpleStatementExpressionListAssignAST is called when production simpleStatementExpressionListAssignAST is entered.
func (s *BaseminiGoParserListener) EnterSimpleStatementExpressionListAssignAST(ctx *SimpleStatementExpressionListAssignASTContext) {
}

// ExitSimpleStatementExpressionListAssignAST is called when production simpleStatementExpressionListAssignAST is exited.
func (s *BaseminiGoParserListener) ExitSimpleStatementExpressionListAssignAST(ctx *SimpleStatementExpressionListAssignASTContext) {
}

// EnterAssignmentStatementAST is called when production assignmentStatementAST is entered.
func (s *BaseminiGoParserListener) EnterAssignmentStatementAST(ctx *AssignmentStatementASTContext) {}

// ExitAssignmentStatementAST is called when production assignmentStatementAST is exited.
func (s *BaseminiGoParserListener) ExitAssignmentStatementAST(ctx *AssignmentStatementASTContext) {}

// EnterAssignmentStatementPlusEqualAST is called when production assignmentStatementPlusEqualAST is entered.
func (s *BaseminiGoParserListener) EnterAssignmentStatementPlusEqualAST(ctx *AssignmentStatementPlusEqualASTContext) {
}

// ExitAssignmentStatementPlusEqualAST is called when production assignmentStatementPlusEqualAST is exited.
func (s *BaseminiGoParserListener) ExitAssignmentStatementPlusEqualAST(ctx *AssignmentStatementPlusEqualASTContext) {
}

// EnterAssignmentStatementMinusEqualAST is called when production assignmentStatementMinusEqualAST is entered.
func (s *BaseminiGoParserListener) EnterAssignmentStatementMinusEqualAST(ctx *AssignmentStatementMinusEqualASTContext) {
}

// ExitAssignmentStatementMinusEqualAST is called when production assignmentStatementMinusEqualAST is exited.
func (s *BaseminiGoParserListener) ExitAssignmentStatementMinusEqualAST(ctx *AssignmentStatementMinusEqualASTContext) {
}

// EnterAssignmentStatementBitwiseAndEqualAST is called when production assignmentStatementBitwiseAndEqualAST is entered.
func (s *BaseminiGoParserListener) EnterAssignmentStatementBitwiseAndEqualAST(ctx *AssignmentStatementBitwiseAndEqualASTContext) {
}

// ExitAssignmentStatementBitwiseAndEqualAST is called when production assignmentStatementBitwiseAndEqualAST is exited.
func (s *BaseminiGoParserListener) ExitAssignmentStatementBitwiseAndEqualAST(ctx *AssignmentStatementBitwiseAndEqualASTContext) {
}

// EnterAssignmentStatementBitwiseOrEqualAST is called when production assignmentStatementBitwiseOrEqualAST is entered.
func (s *BaseminiGoParserListener) EnterAssignmentStatementBitwiseOrEqualAST(ctx *AssignmentStatementBitwiseOrEqualASTContext) {
}

// ExitAssignmentStatementBitwiseOrEqualAST is called when production assignmentStatementBitwiseOrEqualAST is exited.
func (s *BaseminiGoParserListener) ExitAssignmentStatementBitwiseOrEqualAST(ctx *AssignmentStatementBitwiseOrEqualASTContext) {
}

// EnterAssignmentStatementMultiplyEqualAST is called when production assignmentStatementMultiplyEqualAST is entered.
func (s *BaseminiGoParserListener) EnterAssignmentStatementMultiplyEqualAST(ctx *AssignmentStatementMultiplyEqualASTContext) {
}

// ExitAssignmentStatementMultiplyEqualAST is called when production assignmentStatementMultiplyEqualAST is exited.
func (s *BaseminiGoParserListener) ExitAssignmentStatementMultiplyEqualAST(ctx *AssignmentStatementMultiplyEqualASTContext) {
}

// EnterAssignmentStatementBitwiseXorEqualAST is called when production assignmentStatementBitwiseXorEqualAST is entered.
func (s *BaseminiGoParserListener) EnterAssignmentStatementBitwiseXorEqualAST(ctx *AssignmentStatementBitwiseXorEqualASTContext) {
}

// ExitAssignmentStatementBitwiseXorEqualAST is called when production assignmentStatementBitwiseXorEqualAST is exited.
func (s *BaseminiGoParserListener) ExitAssignmentStatementBitwiseXorEqualAST(ctx *AssignmentStatementBitwiseXorEqualASTContext) {
}

// EnterAssignmentStatementShiftLeftEqualAST is called when production assignmentStatementShiftLeftEqualAST is entered.
func (s *BaseminiGoParserListener) EnterAssignmentStatementShiftLeftEqualAST(ctx *AssignmentStatementShiftLeftEqualASTContext) {
}

// ExitAssignmentStatementShiftLeftEqualAST is called when production assignmentStatementShiftLeftEqualAST is exited.
func (s *BaseminiGoParserListener) ExitAssignmentStatementShiftLeftEqualAST(ctx *AssignmentStatementShiftLeftEqualASTContext) {
}

// EnterAssignmentStatementShiftRightEqualAST is called when production assignmentStatementShiftRightEqualAST is entered.
func (s *BaseminiGoParserListener) EnterAssignmentStatementShiftRightEqualAST(ctx *AssignmentStatementShiftRightEqualASTContext) {
}

// ExitAssignmentStatementShiftRightEqualAST is called when production assignmentStatementShiftRightEqualAST is exited.
func (s *BaseminiGoParserListener) ExitAssignmentStatementShiftRightEqualAST(ctx *AssignmentStatementShiftRightEqualASTContext) {
}

// EnterAssignmentStatementBitwiseClearEqualAST is called when production assignmentStatementBitwiseClearEqualAST is entered.
func (s *BaseminiGoParserListener) EnterAssignmentStatementBitwiseClearEqualAST(ctx *AssignmentStatementBitwiseClearEqualASTContext) {
}

// ExitAssignmentStatementBitwiseClearEqualAST is called when production assignmentStatementBitwiseClearEqualAST is exited.
func (s *BaseminiGoParserListener) ExitAssignmentStatementBitwiseClearEqualAST(ctx *AssignmentStatementBitwiseClearEqualASTContext) {
}

// EnterAssignmentStatementModuloEqualAST is called when production assignmentStatementModuloEqualAST is entered.
func (s *BaseminiGoParserListener) EnterAssignmentStatementModuloEqualAST(ctx *AssignmentStatementModuloEqualASTContext) {
}

// ExitAssignmentStatementModuloEqualAST is called when production assignmentStatementModuloEqualAST is exited.
func (s *BaseminiGoParserListener) ExitAssignmentStatementModuloEqualAST(ctx *AssignmentStatementModuloEqualASTContext) {
}

// EnterAssignmentStatementDivideEqualAST is called when production assignmentStatementDivideEqualAST is entered.
func (s *BaseminiGoParserListener) EnterAssignmentStatementDivideEqualAST(ctx *AssignmentStatementDivideEqualASTContext) {
}

// ExitAssignmentStatementDivideEqualAST is called when production assignmentStatementDivideEqualAST is exited.
func (s *BaseminiGoParserListener) ExitAssignmentStatementDivideEqualAST(ctx *AssignmentStatementDivideEqualASTContext) {
}

// EnterIfStatementAST is called when production ifStatementAST is entered.
func (s *BaseminiGoParserListener) EnterIfStatementAST(ctx *IfStatementASTContext) {}

// ExitIfStatementAST is called when production ifStatementAST is exited.
func (s *BaseminiGoParserListener) ExitIfStatementAST(ctx *IfStatementASTContext) {}

// EnterIfElseIfStatementAST is called when production ifElseIfStatementAST is entered.
func (s *BaseminiGoParserListener) EnterIfElseIfStatementAST(ctx *IfElseIfStatementASTContext) {}

// ExitIfElseIfStatementAST is called when production ifElseIfStatementAST is exited.
func (s *BaseminiGoParserListener) ExitIfElseIfStatementAST(ctx *IfElseIfStatementASTContext) {}

// EnterIfElseStatementAST is called when production ifElseStatementAST is entered.
func (s *BaseminiGoParserListener) EnterIfElseStatementAST(ctx *IfElseStatementASTContext) {}

// ExitIfElseStatementAST is called when production ifElseStatementAST is exited.
func (s *BaseminiGoParserListener) ExitIfElseStatementAST(ctx *IfElseStatementASTContext) {}

// EnterIfSimpleStatementAST is called when production ifSimpleStatementAST is entered.
func (s *BaseminiGoParserListener) EnterIfSimpleStatementAST(ctx *IfSimpleStatementASTContext) {}

// ExitIfSimpleStatementAST is called when production ifSimpleStatementAST is exited.
func (s *BaseminiGoParserListener) ExitIfSimpleStatementAST(ctx *IfSimpleStatementASTContext) {}

// EnterIfSimpleElseIfStatementAST is called when production ifSimpleElseIfStatementAST is entered.
func (s *BaseminiGoParserListener) EnterIfSimpleElseIfStatementAST(ctx *IfSimpleElseIfStatementASTContext) {
}

// ExitIfSimpleElseIfStatementAST is called when production ifSimpleElseIfStatementAST is exited.
func (s *BaseminiGoParserListener) ExitIfSimpleElseIfStatementAST(ctx *IfSimpleElseIfStatementASTContext) {
}

// EnterIfSimpleElseStatementAST is called when production ifSimpleElseStatementAST is entered.
func (s *BaseminiGoParserListener) EnterIfSimpleElseStatementAST(ctx *IfSimpleElseStatementASTContext) {
}

// ExitIfSimpleElseStatementAST is called when production ifSimpleElseStatementAST is exited.
func (s *BaseminiGoParserListener) ExitIfSimpleElseStatementAST(ctx *IfSimpleElseStatementASTContext) {
}

// EnterLoopBlockAST is called when production loopBlockAST is entered.
func (s *BaseminiGoParserListener) EnterLoopBlockAST(ctx *LoopBlockASTContext) {}

// ExitLoopBlockAST is called when production loopBlockAST is exited.
func (s *BaseminiGoParserListener) ExitLoopBlockAST(ctx *LoopBlockASTContext) {}

// EnterLoopExpressionBlockAST is called when production loopExpressionBlockAST is entered.
func (s *BaseminiGoParserListener) EnterLoopExpressionBlockAST(ctx *LoopExpressionBlockASTContext) {}

// ExitLoopExpressionBlockAST is called when production loopExpressionBlockAST is exited.
func (s *BaseminiGoParserListener) ExitLoopExpressionBlockAST(ctx *LoopExpressionBlockASTContext) {}

// EnterLoopSimpleStatementExpressionSimpleStatementBlockAST is called when production loopSimpleStatementExpressionSimpleStatementBlockAST is entered.
func (s *BaseminiGoParserListener) EnterLoopSimpleStatementExpressionSimpleStatementBlockAST(ctx *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) {
}

// ExitLoopSimpleStatementExpressionSimpleStatementBlockAST is called when production loopSimpleStatementExpressionSimpleStatementBlockAST is exited.
func (s *BaseminiGoParserListener) ExitLoopSimpleStatementExpressionSimpleStatementBlockAST(ctx *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) {
}

// EnterLoopSimpleStatementSimpleStatementBlockAST is called when production loopSimpleStatementSimpleStatementBlockAST is entered.
func (s *BaseminiGoParserListener) EnterLoopSimpleStatementSimpleStatementBlockAST(ctx *LoopSimpleStatementSimpleStatementBlockASTContext) {
}

// ExitLoopSimpleStatementSimpleStatementBlockAST is called when production loopSimpleStatementSimpleStatementBlockAST is exited.
func (s *BaseminiGoParserListener) ExitLoopSimpleStatementSimpleStatementBlockAST(ctx *LoopSimpleStatementSimpleStatementBlockASTContext) {
}

// EnterSwitchStmtSimpleStatementAST is called when production switchStmtSimpleStatementAST is entered.
func (s *BaseminiGoParserListener) EnterSwitchStmtSimpleStatementAST(ctx *SwitchStmtSimpleStatementASTContext) {
}

// ExitSwitchStmtSimpleStatementAST is called when production switchStmtSimpleStatementAST is exited.
func (s *BaseminiGoParserListener) ExitSwitchStmtSimpleStatementAST(ctx *SwitchStmtSimpleStatementASTContext) {
}

// EnterSwitchStmtExpressionAST is called when production switchStmtExpressionAST is entered.
func (s *BaseminiGoParserListener) EnterSwitchStmtExpressionAST(ctx *SwitchStmtExpressionASTContext) {
}

// ExitSwitchStmtExpressionAST is called when production switchStmtExpressionAST is exited.
func (s *BaseminiGoParserListener) ExitSwitchStmtExpressionAST(ctx *SwitchStmtExpressionASTContext) {}

// EnterSwitchStmtSimpleStatementBlockAST is called when production switchStmtSimpleStatementBlockAST is entered.
func (s *BaseminiGoParserListener) EnterSwitchStmtSimpleStatementBlockAST(ctx *SwitchStmtSimpleStatementBlockASTContext) {
}

// ExitSwitchStmtSimpleStatementBlockAST is called when production switchStmtSimpleStatementBlockAST is exited.
func (s *BaseminiGoParserListener) ExitSwitchStmtSimpleStatementBlockAST(ctx *SwitchStmtSimpleStatementBlockASTContext) {
}

// EnterSwitchStmtBlockAST is called when production switchStmtBlockAST is entered.
func (s *BaseminiGoParserListener) EnterSwitchStmtBlockAST(ctx *SwitchStmtBlockASTContext) {}

// ExitSwitchStmtBlockAST is called when production switchStmtBlockAST is exited.
func (s *BaseminiGoParserListener) ExitSwitchStmtBlockAST(ctx *SwitchStmtBlockASTContext) {}

// EnterExpressionCaseClauseListEmptyAST is called when production expressionCaseClauseListEmptyAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionCaseClauseListEmptyAST(ctx *ExpressionCaseClauseListEmptyASTContext) {
}

// ExitExpressionCaseClauseListEmptyAST is called when production expressionCaseClauseListEmptyAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionCaseClauseListEmptyAST(ctx *ExpressionCaseClauseListEmptyASTContext) {
}

// EnterExpressionCaseClauseListAST is called when production expressionCaseClauseListAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionCaseClauseListAST(ctx *ExpressionCaseClauseListASTContext) {
}

// ExitExpressionCaseClauseListAST is called when production expressionCaseClauseListAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionCaseClauseListAST(ctx *ExpressionCaseClauseListASTContext) {
}

// EnterExpressionCaseClauseAST is called when production expressionCaseClauseAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionCaseClauseAST(ctx *ExpressionCaseClauseASTContext) {
}

// ExitExpressionCaseClauseAST is called when production expressionCaseClauseAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionCaseClauseAST(ctx *ExpressionCaseClauseASTContext) {}

// EnterExpressionSwitchCaseAST is called when production expressionSwitchCaseAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionSwitchCaseAST(ctx *ExpressionSwitchCaseASTContext) {
}

// ExitExpressionSwitchCaseAST is called when production expressionSwitchCaseAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionSwitchCaseAST(ctx *ExpressionSwitchCaseASTContext) {}

// EnterExpressionSwitchDefaultAST is called when production expressionSwitchDefaultAST is entered.
func (s *BaseminiGoParserListener) EnterExpressionSwitchDefaultAST(ctx *ExpressionSwitchDefaultASTContext) {
}

// ExitExpressionSwitchDefaultAST is called when production expressionSwitchDefaultAST is exited.
func (s *BaseminiGoParserListener) ExitExpressionSwitchDefaultAST(ctx *ExpressionSwitchDefaultASTContext) {
}

// EnterEpsilonAST is called when production epsilonAST is entered.
func (s *BaseminiGoParserListener) EnterEpsilonAST(ctx *EpsilonASTContext) {}

// ExitEpsilonAST is called when production epsilonAST is exited.
func (s *BaseminiGoParserListener) ExitEpsilonAST(ctx *EpsilonASTContext) {}
