// Code generated from C:/Users/noni4/Desktop/miniGoCompiler/serverAndCompiler/MiniGoParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // MiniGoParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MiniGoParser.
type MiniGoParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MiniGoParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by MiniGoParser#topDeclarationList.
	VisitTopDeclarationList(ctx *TopDeclarationListContext) interface{}

	// Visit a parse tree produced by MiniGoParser#variableDeclAST.
	VisitVariableDeclAST(ctx *VariableDeclASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#variableDeclBlockAST.
	VisitVariableDeclBlockAST(ctx *VariableDeclBlockASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#variableDeclEmptyAST.
	VisitVariableDeclEmptyAST(ctx *VariableDeclEmptyASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#innerVarDecls.
	VisitInnerVarDecls(ctx *InnerVarDeclsContext) interface{}

	// Visit a parse tree produced by MiniGoParser#singleVarDeclAST.
	VisitSingleVarDeclAST(ctx *SingleVarDeclASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#singleVarDeclAssignAST.
	VisitSingleVarDeclAssignAST(ctx *SingleVarDeclAssignASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#singleVarDeclNoExpsAST.
	VisitSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#singleVarDeclNoExps.
	VisitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) interface{}

	// Visit a parse tree produced by MiniGoParser#typeDeclAST.
	VisitTypeDeclAST(ctx *TypeDeclASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#typeDeclBlockAST.
	VisitTypeDeclBlockAST(ctx *TypeDeclBlockASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#typeDeclEmptyAST.
	VisitTypeDeclEmptyAST(ctx *TypeDeclEmptyASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#innerTypeDecls.
	VisitInnerTypeDecls(ctx *InnerTypeDeclsContext) interface{}

	// Visit a parse tree produced by MiniGoParser#singleTypeDecl.
	VisitSingleTypeDecl(ctx *SingleTypeDeclContext) interface{}

	// Visit a parse tree produced by MiniGoParser#funcDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by MiniGoParser#funcFrontDecl.
	VisitFuncFrontDecl(ctx *FuncFrontDeclContext) interface{}

	// Visit a parse tree produced by MiniGoParser#multipleReturnTypes.
	VisitMultipleReturnTypes(ctx *MultipleReturnTypesContext) interface{}

	// Visit a parse tree produced by MiniGoParser#returnTypeList.
	VisitReturnTypeList(ctx *ReturnTypeListContext) interface{}

	// Visit a parse tree produced by MiniGoParser#singleReturnTypeAST.
	VisitSingleReturnTypeAST(ctx *SingleReturnTypeASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#singleReturnTypeEmptyAST.
	VisitSingleReturnTypeEmptyAST(ctx *SingleReturnTypeEmptyASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#funcArgDecls.
	VisitFuncArgDecls(ctx *FuncArgDeclsContext) interface{}

	// Visit a parse tree produced by MiniGoParser#declTypeParenAST.
	VisitDeclTypeParenAST(ctx *DeclTypeParenASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#declTypeIdentifierAST.
	VisitDeclTypeIdentifierAST(ctx *DeclTypeIdentifierASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#declTypeSliceAST.
	VisitDeclTypeSliceAST(ctx *DeclTypeSliceASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#declTypeArrayAST.
	VisitDeclTypeArrayAST(ctx *DeclTypeArrayASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#declTypeStructAST.
	VisitDeclTypeStructAST(ctx *DeclTypeStructASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#sliceDeclType.
	VisitSliceDeclType(ctx *SliceDeclTypeContext) interface{}

	// Visit a parse tree produced by MiniGoParser#arrayDeclType.
	VisitArrayDeclType(ctx *ArrayDeclTypeContext) interface{}

	// Visit a parse tree produced by MiniGoParser#structDeclType.
	VisitStructDeclType(ctx *StructDeclTypeContext) interface{}

	// Visit a parse tree produced by MiniGoParser#structMemDecls.
	VisitStructMemDecls(ctx *StructMemDeclsContext) interface{}

	// Visit a parse tree produced by MiniGoParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionNotUnaryAST.
	VisitExpressionNotUnaryAST(ctx *ExpressionNotUnaryASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionMultiplyAST.
	VisitExpressionMultiplyAST(ctx *ExpressionMultiplyASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionPlusAST.
	VisitExpressionPlusAST(ctx *ExpressionPlusASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionModuloAST.
	VisitExpressionModuloAST(ctx *ExpressionModuloASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionAndAST.
	VisitExpressionAndAST(ctx *ExpressionAndASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionBitwiseXorAST.
	VisitExpressionBitwiseXorAST(ctx *ExpressionBitwiseXorASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionMinusAST.
	VisitExpressionMinusAST(ctx *ExpressionMinusASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionBitwiseXorUnaryAST.
	VisitExpressionBitwiseXorUnaryAST(ctx *ExpressionBitwiseXorUnaryASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionEqualAST.
	VisitExpressionEqualAST(ctx *ExpressionEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionMinusUnaryAST.
	VisitExpressionMinusUnaryAST(ctx *ExpressionMinusUnaryASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionBitwiseClearAST.
	VisitExpressionBitwiseClearAST(ctx *ExpressionBitwiseClearASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionGreaterEqualAST.
	VisitExpressionGreaterEqualAST(ctx *ExpressionGreaterEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionLessEqualAST.
	VisitExpressionLessEqualAST(ctx *ExpressionLessEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionNotEqualAST.
	VisitExpressionNotEqualAST(ctx *ExpressionNotEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionPrimaryAST.
	VisitExpressionPrimaryAST(ctx *ExpressionPrimaryASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionBitwiseAndAST.
	VisitExpressionBitwiseAndAST(ctx *ExpressionBitwiseAndASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionGreaterAST.
	VisitExpressionGreaterAST(ctx *ExpressionGreaterASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionDivideAST.
	VisitExpressionDivideAST(ctx *ExpressionDivideASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionOrAST.
	VisitExpressionOrAST(ctx *ExpressionOrASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionPlusUnaryAST.
	VisitExpressionPlusUnaryAST(ctx *ExpressionPlusUnaryASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionBitwiseOrAST.
	VisitExpressionBitwiseOrAST(ctx *ExpressionBitwiseOrASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionLessAST.
	VisitExpressionLessAST(ctx *ExpressionLessASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionShiftRightAST.
	VisitExpressionShiftRightAST(ctx *ExpressionShiftRightASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionShiftLeftAST.
	VisitExpressionShiftLeftAST(ctx *ExpressionShiftLeftASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionLengthAST.
	VisitPrimaryExpressionLengthAST(ctx *PrimaryExpressionLengthASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionOperandAST.
	VisitPrimaryExpressionOperandAST(ctx *PrimaryExpressionOperandASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionIndexAST.
	VisitPrimaryExpressionIndexAST(ctx *PrimaryExpressionIndexASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionAppendAST.
	VisitPrimaryExpressionAppendAST(ctx *PrimaryExpressionAppendASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionArgumentsAST.
	VisitPrimaryExpressionArgumentsAST(ctx *PrimaryExpressionArgumentsASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionCapAST.
	VisitPrimaryExpressionCapAST(ctx *PrimaryExpressionCapASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#primaryExpressionSelectorAST.
	VisitPrimaryExpressionSelectorAST(ctx *PrimaryExpressionSelectorASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#operandLiteralAST.
	VisitOperandLiteralAST(ctx *OperandLiteralASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#operandIdentifierAST.
	VisitOperandIdentifierAST(ctx *OperandIdentifierASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#operandParenAST.
	VisitOperandParenAST(ctx *OperandParenASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#literalIntAST.
	VisitLiteralIntAST(ctx *LiteralIntASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#literalFloatAST.
	VisitLiteralFloatAST(ctx *LiteralFloatASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#literalRuneAST.
	VisitLiteralRuneAST(ctx *LiteralRuneASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#literalRawStringAST.
	VisitLiteralRawStringAST(ctx *LiteralRawStringASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#literalInterpretedStringAST.
	VisitLiteralInterpretedStringAST(ctx *LiteralInterpretedStringASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by MiniGoParser#argumentsAST.
	VisitArgumentsAST(ctx *ArgumentsASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#argumentsEmptyAST.
	VisitArgumentsEmptyAST(ctx *ArgumentsEmptyASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by MiniGoParser#appendExpression.
	VisitAppendExpression(ctx *AppendExpressionContext) interface{}

	// Visit a parse tree produced by MiniGoParser#lengthExpression.
	VisitLengthExpression(ctx *LengthExpressionContext) interface{}

	// Visit a parse tree produced by MiniGoParser#capExpression.
	VisitCapExpression(ctx *CapExpressionContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by MiniGoParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementPrintAST.
	VisitStatementPrintAST(ctx *StatementPrintASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementPrintlnAST.
	VisitStatementPrintlnAST(ctx *StatementPrintlnASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementReturnAST.
	VisitStatementReturnAST(ctx *StatementReturnASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementBreakAST.
	VisitStatementBreakAST(ctx *StatementBreakASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementContinueAST.
	VisitStatementContinueAST(ctx *StatementContinueASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementSimpleAST.
	VisitStatementSimpleAST(ctx *StatementSimpleASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementBlockAST.
	VisitStatementBlockAST(ctx *StatementBlockASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementSwitchAST.
	VisitStatementSwitchAST(ctx *StatementSwitchASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementIfAST.
	VisitStatementIfAST(ctx *StatementIfASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementLoopAST.
	VisitStatementLoopAST(ctx *StatementLoopASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementTypeDeclAST.
	VisitStatementTypeDeclAST(ctx *StatementTypeDeclASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#statementVariableDeclAST.
	VisitStatementVariableDeclAST(ctx *StatementVariableDeclASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#simpleStatementEmptyAST.
	VisitSimpleStatementEmptyAST(ctx *SimpleStatementEmptyASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#simpleStatementExpressionAST.
	VisitSimpleStatementExpressionAST(ctx *SimpleStatementExpressionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#simpleStatementAssignmentAST.
	VisitSimpleStatementAssignmentAST(ctx *SimpleStatementAssignmentASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#simpleStatementExpressionListAssignAST.
	VisitSimpleStatementExpressionListAssignAST(ctx *SimpleStatementExpressionListAssignASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementAST.
	VisitAssignmentStatementAST(ctx *AssignmentStatementASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementPlusEqualAST.
	VisitAssignmentStatementPlusEqualAST(ctx *AssignmentStatementPlusEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementMinusEqualAST.
	VisitAssignmentStatementMinusEqualAST(ctx *AssignmentStatementMinusEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementBitwiseAndEqualAST.
	VisitAssignmentStatementBitwiseAndEqualAST(ctx *AssignmentStatementBitwiseAndEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementBitwiseOrEqualAST.
	VisitAssignmentStatementBitwiseOrEqualAST(ctx *AssignmentStatementBitwiseOrEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementMultiplyEqualAST.
	VisitAssignmentStatementMultiplyEqualAST(ctx *AssignmentStatementMultiplyEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementBitwiseXorEqualAST.
	VisitAssignmentStatementBitwiseXorEqualAST(ctx *AssignmentStatementBitwiseXorEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementShiftLeftEqualAST.
	VisitAssignmentStatementShiftLeftEqualAST(ctx *AssignmentStatementShiftLeftEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementShiftRightEqualAST.
	VisitAssignmentStatementShiftRightEqualAST(ctx *AssignmentStatementShiftRightEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementBitwiseClearEqualAST.
	VisitAssignmentStatementBitwiseClearEqualAST(ctx *AssignmentStatementBitwiseClearEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementModuloEqualAST.
	VisitAssignmentStatementModuloEqualAST(ctx *AssignmentStatementModuloEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#assignmentStatementDivideEqualAST.
	VisitAssignmentStatementDivideEqualAST(ctx *AssignmentStatementDivideEqualASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifStatementAST.
	VisitIfStatementAST(ctx *IfStatementASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifElseIfStatementAST.
	VisitIfElseIfStatementAST(ctx *IfElseIfStatementASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifElseStatementAST.
	VisitIfElseStatementAST(ctx *IfElseStatementASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifSimpleStatementAST.
	VisitIfSimpleStatementAST(ctx *IfSimpleStatementASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifSimpleElseIfStatementAST.
	VisitIfSimpleElseIfStatementAST(ctx *IfSimpleElseIfStatementASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#ifSimpleElseStatementAST.
	VisitIfSimpleElseStatementAST(ctx *IfSimpleElseStatementASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#loopBlockAST.
	VisitLoopBlockAST(ctx *LoopBlockASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#loopExpressionBlockAST.
	VisitLoopExpressionBlockAST(ctx *LoopExpressionBlockASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#loopSimpleStatementExpressionSimpleStatementBlockAST.
	VisitLoopSimpleStatementExpressionSimpleStatementBlockAST(ctx *LoopSimpleStatementExpressionSimpleStatementBlockASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#loopSimpleStatementSimpleStatementBlockAST.
	VisitLoopSimpleStatementSimpleStatementBlockAST(ctx *LoopSimpleStatementSimpleStatementBlockASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#switchStmtSimpleStatementAST.
	VisitSwitchStmtSimpleStatementAST(ctx *SwitchStmtSimpleStatementASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#switchStmtExpressionAST.
	VisitSwitchStmtExpressionAST(ctx *SwitchStmtExpressionASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#switchStmtSimpleStatementBlockAST.
	VisitSwitchStmtSimpleStatementBlockAST(ctx *SwitchStmtSimpleStatementBlockASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#switchStmtBlockAST.
	VisitSwitchStmtBlockAST(ctx *SwitchStmtBlockASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionCaseClauseListEmptyAST.
	VisitExpressionCaseClauseListEmptyAST(ctx *ExpressionCaseClauseListEmptyASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionCaseClauseListAST.
	VisitExpressionCaseClauseListAST(ctx *ExpressionCaseClauseListASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionCaseClause.
	VisitExpressionCaseClause(ctx *ExpressionCaseClauseContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionSwitchCaseAST.
	VisitExpressionSwitchCaseAST(ctx *ExpressionSwitchCaseASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#expressionSwitchDefaultAST.
	VisitExpressionSwitchDefaultAST(ctx *ExpressionSwitchDefaultASTContext) interface{}

	// Visit a parse tree produced by MiniGoParser#epsilon.
	VisitEpsilon(ctx *EpsilonContext) interface{}
}
