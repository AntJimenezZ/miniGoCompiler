// Code generated from C:/Users/noni4/Desktop/miniGoProyect/serverAndCompiler/miniGoParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package generated // miniGoParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by miniGoParser.
type miniGoParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by miniGoParser#rootAST.
	VisitRootAST(ctx *RootASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#topDeclarationListAST.
	VisitTopDeclarationListAST(ctx *TopDeclarationListASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#variableDeclAST.
	VisitVariableDeclAST(ctx *VariableDeclASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#variableDeclBlockAST.
	VisitVariableDeclBlockAST(ctx *VariableDeclBlockASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#variableDeclEmptyAST.
	VisitVariableDeclEmptyAST(ctx *VariableDeclEmptyASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#innerVarDeclsAST.
	VisitInnerVarDeclsAST(ctx *InnerVarDeclsASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#singleVarDeclAST.
	VisitSingleVarDeclAST(ctx *SingleVarDeclASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#singleVarDeclAssignAST.
	VisitSingleVarDeclAssignAST(ctx *SingleVarDeclAssignASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#singleVarDeclNoExpsAST.
	VisitSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#singleVarDeclNoExpsDeclTypeAST.
	VisitSingleVarDeclNoExpsDeclTypeAST(ctx *SingleVarDeclNoExpsDeclTypeASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#typeDeclAST.
	VisitTypeDeclAST(ctx *TypeDeclASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#typeDeclBlockAST.
	VisitTypeDeclBlockAST(ctx *TypeDeclBlockASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#typeDeclEmptyAST.
	VisitTypeDeclEmptyAST(ctx *TypeDeclEmptyASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#innerTypeDeclsAST.
	VisitInnerTypeDeclsAST(ctx *InnerTypeDeclsASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#singleTypeDeclAST.
	VisitSingleTypeDeclAST(ctx *SingleTypeDeclASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#funcDeclAST.
	VisitFuncDeclAST(ctx *FuncDeclASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#funcFrontDeclAST.
	VisitFuncFrontDeclAST(ctx *FuncFrontDeclASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#funcArgDeclsAST.
	VisitFuncArgDeclsAST(ctx *FuncArgDeclsASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#declTypeParenAST.
	VisitDeclTypeParenAST(ctx *DeclTypeParenASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#declTypeIdentifierAST.
	VisitDeclTypeIdentifierAST(ctx *DeclTypeIdentifierASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#declTypeSliceAST.
	VisitDeclTypeSliceAST(ctx *DeclTypeSliceASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#declTypeArrayAST.
	VisitDeclTypeArrayAST(ctx *DeclTypeArrayASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#declTypeStructAST.
	VisitDeclTypeStructAST(ctx *DeclTypeStructASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#sliceDeclTypeAST.
	VisitSliceDeclTypeAST(ctx *SliceDeclTypeASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#arrayDeclTypeAST.
	VisitArrayDeclTypeAST(ctx *ArrayDeclTypeASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#structDeclTypeAST.
	VisitStructDeclTypeAST(ctx *StructDeclTypeASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#structMemDeclsAST.
	VisitStructMemDeclsAST(ctx *StructMemDeclsASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#identifierListAST.
	VisitIdentifierListAST(ctx *IdentifierListASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionNotUnaryAST.
	VisitExpressionNotUnaryAST(ctx *ExpressionNotUnaryASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionMultiplyAST.
	VisitExpressionMultiplyAST(ctx *ExpressionMultiplyASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionPlusAST.
	VisitExpressionPlusAST(ctx *ExpressionPlusASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionModuloAST.
	VisitExpressionModuloAST(ctx *ExpressionModuloASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionAndAST.
	VisitExpressionAndAST(ctx *ExpressionAndASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionBitwiseXorAST.
	VisitExpressionBitwiseXorAST(ctx *ExpressionBitwiseXorASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionMinusAST.
	VisitExpressionMinusAST(ctx *ExpressionMinusASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionBitwiseXorUnaryAST.
	VisitExpressionBitwiseXorUnaryAST(ctx *ExpressionBitwiseXorUnaryASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionEqualAST.
	VisitExpressionEqualAST(ctx *ExpressionEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionMinusUnaryAST.
	VisitExpressionMinusUnaryAST(ctx *ExpressionMinusUnaryASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionBitwiseClearAST.
	VisitExpressionBitwiseClearAST(ctx *ExpressionBitwiseClearASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionGreaterEqualAST.
	VisitExpressionGreaterEqualAST(ctx *ExpressionGreaterEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionLessEqualAST.
	VisitExpressionLessEqualAST(ctx *ExpressionLessEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionNotEqualAST.
	VisitExpressionNotEqualAST(ctx *ExpressionNotEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionPrimaryAST.
	VisitExpressionPrimaryAST(ctx *ExpressionPrimaryASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionBitwiseAndAST.
	VisitExpressionBitwiseAndAST(ctx *ExpressionBitwiseAndASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionGreaterAST.
	VisitExpressionGreaterAST(ctx *ExpressionGreaterASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionDivideAST.
	VisitExpressionDivideAST(ctx *ExpressionDivideASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionOrAST.
	VisitExpressionOrAST(ctx *ExpressionOrASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionPlusUnaryAST.
	VisitExpressionPlusUnaryAST(ctx *ExpressionPlusUnaryASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionBitwiseOrAST.
	VisitExpressionBitwiseOrAST(ctx *ExpressionBitwiseOrASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionLessAST.
	VisitExpressionLessAST(ctx *ExpressionLessASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionShiftRightAST.
	VisitExpressionShiftRightAST(ctx *ExpressionShiftRightASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionShiftLeftAST.
	VisitExpressionShiftLeftAST(ctx *ExpressionShiftLeftASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionListAST.
	VisitExpressionListAST(ctx *ExpressionListASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#primaryExpressionLengthAST.
	VisitPrimaryExpressionLengthAST(ctx *PrimaryExpressionLengthASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#primaryExpressionOperandAST.
	VisitPrimaryExpressionOperandAST(ctx *PrimaryExpressionOperandASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#primaryExpressionIndexAST.
	VisitPrimaryExpressionIndexAST(ctx *PrimaryExpressionIndexASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#primaryExpressionAppendAST.
	VisitPrimaryExpressionAppendAST(ctx *PrimaryExpressionAppendASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#primaryExpressionArgumentsAST.
	VisitPrimaryExpressionArgumentsAST(ctx *PrimaryExpressionArgumentsASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#primaryExpressionCapAST.
	VisitPrimaryExpressionCapAST(ctx *PrimaryExpressionCapASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#primaryExpressionSelectorAST.
	VisitPrimaryExpressionSelectorAST(ctx *PrimaryExpressionSelectorASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#operandLiteralAST.
	VisitOperandLiteralAST(ctx *OperandLiteralASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#operandIdentifierAST.
	VisitOperandIdentifierAST(ctx *OperandIdentifierASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#operandParenAST.
	VisitOperandParenAST(ctx *OperandParenASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#literalIntAST.
	VisitLiteralIntAST(ctx *LiteralIntASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#literalFloatAST.
	VisitLiteralFloatAST(ctx *LiteralFloatASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#literalRuneAST.
	VisitLiteralRuneAST(ctx *LiteralRuneASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#literalRawStringAST.
	VisitLiteralRawStringAST(ctx *LiteralRawStringASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#literalInterpretedStringAST.
	VisitLiteralInterpretedStringAST(ctx *LiteralInterpretedStringASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#indexAST.
	VisitIndexAST(ctx *IndexASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#argumentsAST.
	VisitArgumentsAST(ctx *ArgumentsASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#argumentsEmptyAST.
	VisitArgumentsEmptyAST(ctx *ArgumentsEmptyASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#selectorAST.
	VisitSelectorAST(ctx *SelectorASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#appendExpressionAST.
	VisitAppendExpressionAST(ctx *AppendExpressionASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#lengthExpressionAST.
	VisitLengthExpressionAST(ctx *LengthExpressionASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#capExpressionAST.
	VisitCapExpressionAST(ctx *CapExpressionASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementListAST.
	VisitStatementListAST(ctx *StatementListASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#blockAST.
	VisitBlockAST(ctx *BlockASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementPrintAST.
	VisitStatementPrintAST(ctx *StatementPrintASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementPrintlnAST.
	VisitStatementPrintlnAST(ctx *StatementPrintlnASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementReturnAST.
	VisitStatementReturnAST(ctx *StatementReturnASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementBreakAST.
	VisitStatementBreakAST(ctx *StatementBreakASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementContinueAST.
	VisitStatementContinueAST(ctx *StatementContinueASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementSimpleAST.
	VisitStatementSimpleAST(ctx *StatementSimpleASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementBlockAST.
	VisitStatementBlockAST(ctx *StatementBlockASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementSwitchAST.
	VisitStatementSwitchAST(ctx *StatementSwitchASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementIfAST.
	VisitStatementIfAST(ctx *StatementIfASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementLoopAST.
	VisitStatementLoopAST(ctx *StatementLoopASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementTypeDeclAST.
	VisitStatementTypeDeclAST(ctx *StatementTypeDeclASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#statementVariableDeclAST.
	VisitStatementVariableDeclAST(ctx *StatementVariableDeclASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#simpleStatementEmptyAST.
	VisitSimpleStatementEmptyAST(ctx *SimpleStatementEmptyASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#simpleStatementExpressionAST.
	VisitSimpleStatementExpressionAST(ctx *SimpleStatementExpressionASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#simpleStatementAssignmentAST.
	VisitSimpleStatementAssignmentAST(ctx *SimpleStatementAssignmentASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#simpleStatementExpressionListAssignAST.
	VisitSimpleStatementExpressionListAssignAST(ctx *SimpleStatementExpressionListAssignASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#assignmentStatementAST.
	VisitAssignmentStatementAST(ctx *AssignmentStatementASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#assignmentStatementPlusEqualAST.
	VisitAssignmentStatementPlusEqualAST(ctx *AssignmentStatementPlusEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#assignmentStatementMinusEqualAST.
	VisitAssignmentStatementMinusEqualAST(ctx *AssignmentStatementMinusEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#assignmentStatementBitwiseAndEqualAST.
	VisitAssignmentStatementBitwiseAndEqualAST(ctx *AssignmentStatementBitwiseAndEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#assignmentStatementBitwiseOrEqualAST.
	VisitAssignmentStatementBitwiseOrEqualAST(ctx *AssignmentStatementBitwiseOrEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#assignmentStatementMultiplyEqualAST.
	VisitAssignmentStatementMultiplyEqualAST(ctx *AssignmentStatementMultiplyEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#assignmentStatementBitwiseXorEqualAST.
	VisitAssignmentStatementBitwiseXorEqualAST(ctx *AssignmentStatementBitwiseXorEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#assignmentStatementShiftLeftEqualAST.
	VisitAssignmentStatementShiftLeftEqualAST(ctx *AssignmentStatementShiftLeftEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#assignmentStatementShiftRightEqualAST.
	VisitAssignmentStatementShiftRightEqualAST(ctx *AssignmentStatementShiftRightEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#assignmentStatementBitwiseClearEqualAST.
	VisitAssignmentStatementBitwiseClearEqualAST(ctx *AssignmentStatementBitwiseClearEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#assignmentStatementModuloEqualAST.
	VisitAssignmentStatementModuloEqualAST(ctx *AssignmentStatementModuloEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#assignmentStatementDivideEqualAST.
	VisitAssignmentStatementDivideEqualAST(ctx *AssignmentStatementDivideEqualASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#ifStatementAST.
	VisitIfStatementAST(ctx *IfStatementASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#ifElseIfStatementAST.
	VisitIfElseIfStatementAST(ctx *IfElseIfStatementASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#ifElseStatementAST.
	VisitIfElseStatementAST(ctx *IfElseStatementASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#ifSimpleStatementAST.
	VisitIfSimpleStatementAST(ctx *IfSimpleStatementASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#ifSimpleElseIfStatementAST.
	VisitIfSimpleElseIfStatementAST(ctx *IfSimpleElseIfStatementASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#ifSimpleElseStatementAST.
	VisitIfSimpleElseStatementAST(ctx *IfSimpleElseStatementASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#loopBlockAST.
	VisitLoopBlockAST(ctx *LoopBlockASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#loopExpressionBlockAST.
	VisitLoopExpressionBlockAST(ctx *LoopExpressionBlockASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#loopSimpleStatementExpressionSimpleStatementBlockAST.
	VisitLoopSimpleStatementExpressionSimpleStatementBlockAST(ctx *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#loopSimpleStatementSimpleStatementBlockAST.
	VisitLoopSimpleStatementSimpleStatementBlockAST(ctx *LoopSimpleStatementSimpleStatementBlockASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#switchStmtSimpleStatementAST.
	VisitSwitchStmtSimpleStatementAST(ctx *SwitchStmtSimpleStatementASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#switchStmtExpressionAST.
	VisitSwitchStmtExpressionAST(ctx *SwitchStmtExpressionASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#switchStmtSimpleStatementBlockAST.
	VisitSwitchStmtSimpleStatementBlockAST(ctx *SwitchStmtSimpleStatementBlockASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#switchStmtBlockAST.
	VisitSwitchStmtBlockAST(ctx *SwitchStmtBlockASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionCaseClauseListEmptyAST.
	VisitExpressionCaseClauseListEmptyAST(ctx *ExpressionCaseClauseListEmptyASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionCaseClauseListAST.
	VisitExpressionCaseClauseListAST(ctx *ExpressionCaseClauseListASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionCaseClauseAST.
	VisitExpressionCaseClauseAST(ctx *ExpressionCaseClauseASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionSwitchCaseAST.
	VisitExpressionSwitchCaseAST(ctx *ExpressionSwitchCaseASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#expressionSwitchDefaultAST.
	VisitExpressionSwitchDefaultAST(ctx *ExpressionSwitchDefaultASTContext) interface{}

	// Visit a parse tree produced by miniGoParser#epsilonAST.
	VisitEpsilonAST(ctx *EpsilonASTContext) interface{}
}
