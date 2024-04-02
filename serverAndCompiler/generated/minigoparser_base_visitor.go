// Code generated from C:/Users/noni4/Desktop/miniGoProyect/serverAndCompiler/miniGoParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package generated // miniGoParser
import "github.com/antlr4-go/antlr/v4"

type BaseminiGoParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseminiGoParserVisitor) VisitRootAST(ctx *RootASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitTopDeclarationListAST(ctx *TopDeclarationListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitVariableDeclAST(ctx *VariableDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitVariableDeclBlockAST(ctx *VariableDeclBlockASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitVariableDeclEmptyAST(ctx *VariableDeclEmptyASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitInnerVarDeclsAST(ctx *InnerVarDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSingleVarDeclAST(ctx *SingleVarDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSingleVarDeclAssignAST(ctx *SingleVarDeclAssignASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSingleVarDeclNoExpsDeclTypeAST(ctx *SingleVarDeclNoExpsDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitTypeDeclAST(ctx *TypeDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitTypeDeclBlockAST(ctx *TypeDeclBlockASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitTypeDeclEmptyAST(ctx *TypeDeclEmptyASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitInnerTypeDeclsAST(ctx *InnerTypeDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSingleTypeDeclAST(ctx *SingleTypeDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitFuncDeclAST(ctx *FuncDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitFuncFrontDeclAST(ctx *FuncFrontDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitFuncArgDeclsAST(ctx *FuncArgDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitDeclTypeParenAST(ctx *DeclTypeParenASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitDeclTypeIdentifierAST(ctx *DeclTypeIdentifierASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitDeclTypeSliceAST(ctx *DeclTypeSliceASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitDeclTypeArrayAST(ctx *DeclTypeArrayASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitDeclTypeStructAST(ctx *DeclTypeStructASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSliceDeclTypeAST(ctx *SliceDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitArrayDeclTypeAST(ctx *ArrayDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStructDeclTypeAST(ctx *StructDeclTypeASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStructMemDeclsAST(ctx *StructMemDeclsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitIdentifierListAST(ctx *IdentifierListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionNotUnaryAST(ctx *ExpressionNotUnaryASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionMultiplyAST(ctx *ExpressionMultiplyASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionPlusAST(ctx *ExpressionPlusASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionModuloAST(ctx *ExpressionModuloASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionAndAST(ctx *ExpressionAndASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionBitwiseXorAST(ctx *ExpressionBitwiseXorASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionMinusAST(ctx *ExpressionMinusASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionBitwiseXorUnaryAST(ctx *ExpressionBitwiseXorUnaryASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionEqualAST(ctx *ExpressionEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionMinusUnaryAST(ctx *ExpressionMinusUnaryASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionBitwiseClearAST(ctx *ExpressionBitwiseClearASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionGreaterEqualAST(ctx *ExpressionGreaterEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionLessEqualAST(ctx *ExpressionLessEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionNotEqualAST(ctx *ExpressionNotEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionPrimaryAST(ctx *ExpressionPrimaryASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionBitwiseAndAST(ctx *ExpressionBitwiseAndASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionGreaterAST(ctx *ExpressionGreaterASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionDivideAST(ctx *ExpressionDivideASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionOrAST(ctx *ExpressionOrASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionPlusUnaryAST(ctx *ExpressionPlusUnaryASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionBitwiseOrAST(ctx *ExpressionBitwiseOrASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionLessAST(ctx *ExpressionLessASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionShiftRightAST(ctx *ExpressionShiftRightASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionShiftLeftAST(ctx *ExpressionShiftLeftASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionListAST(ctx *ExpressionListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitPrimaryExpressionLengthAST(ctx *PrimaryExpressionLengthASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitPrimaryExpressionOperandAST(ctx *PrimaryExpressionOperandASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitPrimaryExpressionIndexAST(ctx *PrimaryExpressionIndexASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitPrimaryExpressionAppendAST(ctx *PrimaryExpressionAppendASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitPrimaryExpressionArgumentsAST(ctx *PrimaryExpressionArgumentsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitPrimaryExpressionCapAST(ctx *PrimaryExpressionCapASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitPrimaryExpressionSelectorAST(ctx *PrimaryExpressionSelectorASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitOperandLiteralAST(ctx *OperandLiteralASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitOperandIdentifierAST(ctx *OperandIdentifierASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitOperandParenAST(ctx *OperandParenASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitLiteralIntAST(ctx *LiteralIntASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitLiteralFloatAST(ctx *LiteralFloatASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitLiteralRuneAST(ctx *LiteralRuneASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitLiteralRawStringAST(ctx *LiteralRawStringASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitLiteralInterpretedStringAST(ctx *LiteralInterpretedStringASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitIndexAST(ctx *IndexASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitArgumentsAST(ctx *ArgumentsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitArgumentsEmptyAST(ctx *ArgumentsEmptyASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSelectorAST(ctx *SelectorASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAppendExpressionAST(ctx *AppendExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitLengthExpressionAST(ctx *LengthExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitCapExpressionAST(ctx *CapExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementListAST(ctx *StatementListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitBlockAST(ctx *BlockASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementPrintAST(ctx *StatementPrintASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementPrintlnAST(ctx *StatementPrintlnASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementReturnAST(ctx *StatementReturnASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementBreakAST(ctx *StatementBreakASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementContinueAST(ctx *StatementContinueASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementSimpleAST(ctx *StatementSimpleASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementBlockAST(ctx *StatementBlockASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementSwitchAST(ctx *StatementSwitchASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementIfAST(ctx *StatementIfASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementLoopAST(ctx *StatementLoopASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementTypeDeclAST(ctx *StatementTypeDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitStatementVariableDeclAST(ctx *StatementVariableDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSimpleStatementEmptyAST(ctx *SimpleStatementEmptyASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSimpleStatementExpressionAST(ctx *SimpleStatementExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSimpleStatementAssignmentAST(ctx *SimpleStatementAssignmentASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSimpleStatementExpressionListAssignAST(ctx *SimpleStatementExpressionListAssignASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAssignmentStatementAST(ctx *AssignmentStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAssignmentStatementPlusEqualAST(ctx *AssignmentStatementPlusEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAssignmentStatementMinusEqualAST(ctx *AssignmentStatementMinusEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAssignmentStatementBitwiseAndEqualAST(ctx *AssignmentStatementBitwiseAndEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAssignmentStatementBitwiseOrEqualAST(ctx *AssignmentStatementBitwiseOrEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAssignmentStatementMultiplyEqualAST(ctx *AssignmentStatementMultiplyEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAssignmentStatementBitwiseXorEqualAST(ctx *AssignmentStatementBitwiseXorEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAssignmentStatementShiftLeftEqualAST(ctx *AssignmentStatementShiftLeftEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAssignmentStatementShiftRightEqualAST(ctx *AssignmentStatementShiftRightEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAssignmentStatementBitwiseClearEqualAST(ctx *AssignmentStatementBitwiseClearEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAssignmentStatementModuloEqualAST(ctx *AssignmentStatementModuloEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitAssignmentStatementDivideEqualAST(ctx *AssignmentStatementDivideEqualASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitIfStatementAST(ctx *IfStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitIfElseIfStatementAST(ctx *IfElseIfStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitIfElseStatementAST(ctx *IfElseStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitIfSimpleStatementAST(ctx *IfSimpleStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitIfSimpleElseIfStatementAST(ctx *IfSimpleElseIfStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitIfSimpleElseStatementAST(ctx *IfSimpleElseStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitLoopBlockAST(ctx *LoopBlockASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitLoopExpressionBlockAST(ctx *LoopExpressionBlockASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitLoopSimpleStatementExpressionSimpleStatementBlockAST(ctx *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitLoopSimpleStatementSimpleStatementBlockAST(ctx *LoopSimpleStatementSimpleStatementBlockASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSwitchStmtSimpleStatementAST(ctx *SwitchStmtSimpleStatementASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSwitchStmtExpressionAST(ctx *SwitchStmtExpressionASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSwitchStmtSimpleStatementBlockAST(ctx *SwitchStmtSimpleStatementBlockASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitSwitchStmtBlockAST(ctx *SwitchStmtBlockASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionCaseClauseListEmptyAST(ctx *ExpressionCaseClauseListEmptyASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionCaseClauseListAST(ctx *ExpressionCaseClauseListASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionCaseClauseAST(ctx *ExpressionCaseClauseASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionSwitchCaseAST(ctx *ExpressionSwitchCaseASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitExpressionSwitchDefaultAST(ctx *ExpressionSwitchDefaultASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseminiGoParserVisitor) VisitEpsilonAST(ctx *EpsilonASTContext) interface{} {
	return v.VisitChildren(ctx)
}
