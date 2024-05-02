// Code generated from C:/Users/noni4/Desktop/miniGoCompiler/serverAndCompiler/MiniGoParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // MiniGoParser
import "github.com/antlr4-go/antlr/v4"

// MiniGoParserListener is a complete listener for a parse tree produced by MiniGoParser.
type MiniGoParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterTopDeclarationList is called when entering the topDeclarationList production.
	EnterTopDeclarationList(c *TopDeclarationListContext)

	// EnterVariableDeclAST is called when entering the variableDeclAST production.
	EnterVariableDeclAST(c *VariableDeclASTContext)

	// EnterVariableDeclBlockAST is called when entering the variableDeclBlockAST production.
	EnterVariableDeclBlockAST(c *VariableDeclBlockASTContext)

	// EnterVariableDeclEmptyAST is called when entering the variableDeclEmptyAST production.
	EnterVariableDeclEmptyAST(c *VariableDeclEmptyASTContext)

	// EnterInnerVarDecls is called when entering the innerVarDecls production.
	EnterInnerVarDecls(c *InnerVarDeclsContext)

	// EnterSingleVarDeclAST is called when entering the singleVarDeclAST production.
	EnterSingleVarDeclAST(c *SingleVarDeclASTContext)

	// EnterSingleVarDeclAssignAST is called when entering the singleVarDeclAssignAST production.
	EnterSingleVarDeclAssignAST(c *SingleVarDeclAssignASTContext)

	// EnterSingleVarDeclNoExpsAST is called when entering the singleVarDeclNoExpsAST production.
	EnterSingleVarDeclNoExpsAST(c *SingleVarDeclNoExpsASTContext)

	// EnterSingleVarDeclNoExps is called when entering the singleVarDeclNoExps production.
	EnterSingleVarDeclNoExps(c *SingleVarDeclNoExpsContext)

	// EnterTypeDeclAST is called when entering the typeDeclAST production.
	EnterTypeDeclAST(c *TypeDeclASTContext)

	// EnterTypeDeclBlockAST is called when entering the typeDeclBlockAST production.
	EnterTypeDeclBlockAST(c *TypeDeclBlockASTContext)

	// EnterTypeDeclEmptyAST is called when entering the typeDeclEmptyAST production.
	EnterTypeDeclEmptyAST(c *TypeDeclEmptyASTContext)

	// EnterInnerTypeDecls is called when entering the innerTypeDecls production.
	EnterInnerTypeDecls(c *InnerTypeDeclsContext)

	// EnterSingleTypeDecl is called when entering the singleTypeDecl production.
	EnterSingleTypeDecl(c *SingleTypeDeclContext)

	// EnterFuncDecl is called when entering the funcDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterFuncFrontDecl is called when entering the funcFrontDecl production.
	EnterFuncFrontDecl(c *FuncFrontDeclContext)

	// EnterMultipleReturnTypes is called when entering the multipleReturnTypes production.
	EnterMultipleReturnTypes(c *MultipleReturnTypesContext)

	// EnterReturnTypeList is called when entering the returnTypeList production.
	EnterReturnTypeList(c *ReturnTypeListContext)

	// EnterSingleReturnTypeAST is called when entering the singleReturnTypeAST production.
	EnterSingleReturnTypeAST(c *SingleReturnTypeASTContext)

	// EnterSingleReturnTypeEmptyAST is called when entering the singleReturnTypeEmptyAST production.
	EnterSingleReturnTypeEmptyAST(c *SingleReturnTypeEmptyASTContext)

	// EnterFuncArgDecls is called when entering the funcArgDecls production.
	EnterFuncArgDecls(c *FuncArgDeclsContext)

	// EnterDeclTypeParenAST is called when entering the declTypeParenAST production.
	EnterDeclTypeParenAST(c *DeclTypeParenASTContext)

	// EnterDeclTypeIdentifierAST is called when entering the declTypeIdentifierAST production.
	EnterDeclTypeIdentifierAST(c *DeclTypeIdentifierASTContext)

	// EnterDeclTypeSliceAST is called when entering the declTypeSliceAST production.
	EnterDeclTypeSliceAST(c *DeclTypeSliceASTContext)

	// EnterDeclTypeArrayAST is called when entering the declTypeArrayAST production.
	EnterDeclTypeArrayAST(c *DeclTypeArrayASTContext)

	// EnterDeclTypeStructAST is called when entering the declTypeStructAST production.
	EnterDeclTypeStructAST(c *DeclTypeStructASTContext)

	// EnterSliceDeclType is called when entering the sliceDeclType production.
	EnterSliceDeclType(c *SliceDeclTypeContext)

	// EnterArrayDeclType is called when entering the arrayDeclType production.
	EnterArrayDeclType(c *ArrayDeclTypeContext)

	// EnterStructDeclType is called when entering the structDeclType production.
	EnterStructDeclType(c *StructDeclTypeContext)

	// EnterStructMemDecls is called when entering the structMemDecls production.
	EnterStructMemDecls(c *StructMemDeclsContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterExpressionNotUnaryAST is called when entering the expressionNotUnaryAST production.
	EnterExpressionNotUnaryAST(c *ExpressionNotUnaryASTContext)

	// EnterExpressionMultiplyAST is called when entering the expressionMultiplyAST production.
	EnterExpressionMultiplyAST(c *ExpressionMultiplyASTContext)

	// EnterExpressionPlusAST is called when entering the expressionPlusAST production.
	EnterExpressionPlusAST(c *ExpressionPlusASTContext)

	// EnterExpressionModuloAST is called when entering the expressionModuloAST production.
	EnterExpressionModuloAST(c *ExpressionModuloASTContext)

	// EnterExpressionAndAST is called when entering the expressionAndAST production.
	EnterExpressionAndAST(c *ExpressionAndASTContext)

	// EnterExpressionBitwiseXorAST is called when entering the expressionBitwiseXorAST production.
	EnterExpressionBitwiseXorAST(c *ExpressionBitwiseXorASTContext)

	// EnterExpressionMinusAST is called when entering the expressionMinusAST production.
	EnterExpressionMinusAST(c *ExpressionMinusASTContext)

	// EnterExpressionBitwiseXorUnaryAST is called when entering the expressionBitwiseXorUnaryAST production.
	EnterExpressionBitwiseXorUnaryAST(c *ExpressionBitwiseXorUnaryASTContext)

	// EnterExpressionEqualAST is called when entering the expressionEqualAST production.
	EnterExpressionEqualAST(c *ExpressionEqualASTContext)

	// EnterExpressionMinusUnaryAST is called when entering the expressionMinusUnaryAST production.
	EnterExpressionMinusUnaryAST(c *ExpressionMinusUnaryASTContext)

	// EnterExpressionBitwiseClearAST is called when entering the expressionBitwiseClearAST production.
	EnterExpressionBitwiseClearAST(c *ExpressionBitwiseClearASTContext)

	// EnterExpressionGreaterEqualAST is called when entering the expressionGreaterEqualAST production.
	EnterExpressionGreaterEqualAST(c *ExpressionGreaterEqualASTContext)

	// EnterExpressionLessEqualAST is called when entering the expressionLessEqualAST production.
	EnterExpressionLessEqualAST(c *ExpressionLessEqualASTContext)

	// EnterExpressionNotEqualAST is called when entering the expressionNotEqualAST production.
	EnterExpressionNotEqualAST(c *ExpressionNotEqualASTContext)

	// EnterExpressionPrimaryAST is called when entering the expressionPrimaryAST production.
	EnterExpressionPrimaryAST(c *ExpressionPrimaryASTContext)

	// EnterExpressionBitwiseAndAST is called when entering the expressionBitwiseAndAST production.
	EnterExpressionBitwiseAndAST(c *ExpressionBitwiseAndASTContext)

	// EnterExpressionGreaterAST is called when entering the expressionGreaterAST production.
	EnterExpressionGreaterAST(c *ExpressionGreaterASTContext)

	// EnterExpressionDivideAST is called when entering the expressionDivideAST production.
	EnterExpressionDivideAST(c *ExpressionDivideASTContext)

	// EnterExpressionOrAST is called when entering the expressionOrAST production.
	EnterExpressionOrAST(c *ExpressionOrASTContext)

	// EnterExpressionPlusUnaryAST is called when entering the expressionPlusUnaryAST production.
	EnterExpressionPlusUnaryAST(c *ExpressionPlusUnaryASTContext)

	// EnterExpressionBitwiseOrAST is called when entering the expressionBitwiseOrAST production.
	EnterExpressionBitwiseOrAST(c *ExpressionBitwiseOrASTContext)

	// EnterExpressionLessAST is called when entering the expressionLessAST production.
	EnterExpressionLessAST(c *ExpressionLessASTContext)

	// EnterExpressionShiftRightAST is called when entering the expressionShiftRightAST production.
	EnterExpressionShiftRightAST(c *ExpressionShiftRightASTContext)

	// EnterExpressionShiftLeftAST is called when entering the expressionShiftLeftAST production.
	EnterExpressionShiftLeftAST(c *ExpressionShiftLeftASTContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterPrimaryExpressionLengthAST is called when entering the primaryExpressionLengthAST production.
	EnterPrimaryExpressionLengthAST(c *PrimaryExpressionLengthASTContext)

	// EnterPrimaryExpressionOperandAST is called when entering the primaryExpressionOperandAST production.
	EnterPrimaryExpressionOperandAST(c *PrimaryExpressionOperandASTContext)

	// EnterPrimaryExpressionIndexAST is called when entering the primaryExpressionIndexAST production.
	EnterPrimaryExpressionIndexAST(c *PrimaryExpressionIndexASTContext)

	// EnterPrimaryExpressionAppendAST is called when entering the primaryExpressionAppendAST production.
	EnterPrimaryExpressionAppendAST(c *PrimaryExpressionAppendASTContext)

	// EnterPrimaryExpressionArgumentsAST is called when entering the primaryExpressionArgumentsAST production.
	EnterPrimaryExpressionArgumentsAST(c *PrimaryExpressionArgumentsASTContext)

	// EnterPrimaryExpressionCapAST is called when entering the primaryExpressionCapAST production.
	EnterPrimaryExpressionCapAST(c *PrimaryExpressionCapASTContext)

	// EnterPrimaryExpressionSelectorAST is called when entering the primaryExpressionSelectorAST production.
	EnterPrimaryExpressionSelectorAST(c *PrimaryExpressionSelectorASTContext)

	// EnterOperandLiteralAST is called when entering the operandLiteralAST production.
	EnterOperandLiteralAST(c *OperandLiteralASTContext)

	// EnterOperandIdentifierAST is called when entering the operandIdentifierAST production.
	EnterOperandIdentifierAST(c *OperandIdentifierASTContext)

	// EnterOperandParenAST is called when entering the operandParenAST production.
	EnterOperandParenAST(c *OperandParenASTContext)

	// EnterLiteralIntAST is called when entering the literalIntAST production.
	EnterLiteralIntAST(c *LiteralIntASTContext)

	// EnterLiteralFloatAST is called when entering the literalFloatAST production.
	EnterLiteralFloatAST(c *LiteralFloatASTContext)

	// EnterLiteralRuneAST is called when entering the literalRuneAST production.
	EnterLiteralRuneAST(c *LiteralRuneASTContext)

	// EnterLiteralRawStringAST is called when entering the literalRawStringAST production.
	EnterLiteralRawStringAST(c *LiteralRawStringASTContext)

	// EnterLiteralInterpretedStringAST is called when entering the literalInterpretedStringAST production.
	EnterLiteralInterpretedStringAST(c *LiteralInterpretedStringASTContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterArgumentsAST is called when entering the argumentsAST production.
	EnterArgumentsAST(c *ArgumentsASTContext)

	// EnterArgumentsEmptyAST is called when entering the argumentsEmptyAST production.
	EnterArgumentsEmptyAST(c *ArgumentsEmptyASTContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterAppendExpression is called when entering the appendExpression production.
	EnterAppendExpression(c *AppendExpressionContext)

	// EnterLengthExpression is called when entering the lengthExpression production.
	EnterLengthExpression(c *LengthExpressionContext)

	// EnterCapExpression is called when entering the capExpression production.
	EnterCapExpression(c *CapExpressionContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatementPrintAST is called when entering the statementPrintAST production.
	EnterStatementPrintAST(c *StatementPrintASTContext)

	// EnterStatementPrintlnAST is called when entering the statementPrintlnAST production.
	EnterStatementPrintlnAST(c *StatementPrintlnASTContext)

	// EnterStatementReturnAST is called when entering the statementReturnAST production.
	EnterStatementReturnAST(c *StatementReturnASTContext)

	// EnterStatementBreakAST is called when entering the statementBreakAST production.
	EnterStatementBreakAST(c *StatementBreakASTContext)

	// EnterStatementContinueAST is called when entering the statementContinueAST production.
	EnterStatementContinueAST(c *StatementContinueASTContext)

	// EnterStatementSimpleAST is called when entering the statementSimpleAST production.
	EnterStatementSimpleAST(c *StatementSimpleASTContext)

	// EnterStatementBlockAST is called when entering the statementBlockAST production.
	EnterStatementBlockAST(c *StatementBlockASTContext)

	// EnterStatementSwitchAST is called when entering the statementSwitchAST production.
	EnterStatementSwitchAST(c *StatementSwitchASTContext)

	// EnterStatementIfAST is called when entering the statementIfAST production.
	EnterStatementIfAST(c *StatementIfASTContext)

	// EnterStatementLoopAST is called when entering the statementLoopAST production.
	EnterStatementLoopAST(c *StatementLoopASTContext)

	// EnterStatementTypeDeclAST is called when entering the statementTypeDeclAST production.
	EnterStatementTypeDeclAST(c *StatementTypeDeclASTContext)

	// EnterStatementVariableDeclAST is called when entering the statementVariableDeclAST production.
	EnterStatementVariableDeclAST(c *StatementVariableDeclASTContext)

	// EnterSimpleStatementEmptyAST is called when entering the simpleStatementEmptyAST production.
	EnterSimpleStatementEmptyAST(c *SimpleStatementEmptyASTContext)

	// EnterSimpleStatementExpressionAST is called when entering the simpleStatementExpressionAST production.
	EnterSimpleStatementExpressionAST(c *SimpleStatementExpressionASTContext)

	// EnterSimpleStatementAssignmentAST is called when entering the simpleStatementAssignmentAST production.
	EnterSimpleStatementAssignmentAST(c *SimpleStatementAssignmentASTContext)

	// EnterSimpleStatementExpressionListAssignAST is called when entering the simpleStatementExpressionListAssignAST production.
	EnterSimpleStatementExpressionListAssignAST(c *SimpleStatementExpressionListAssignASTContext)

	// EnterAssignmentStatementAST is called when entering the assignmentStatementAST production.
	EnterAssignmentStatementAST(c *AssignmentStatementASTContext)

	// EnterAssignmentStatementPlusEqualAST is called when entering the assignmentStatementPlusEqualAST production.
	EnterAssignmentStatementPlusEqualAST(c *AssignmentStatementPlusEqualASTContext)

	// EnterAssignmentStatementMinusEqualAST is called when entering the assignmentStatementMinusEqualAST production.
	EnterAssignmentStatementMinusEqualAST(c *AssignmentStatementMinusEqualASTContext)

	// EnterAssignmentStatementBitwiseAndEqualAST is called when entering the assignmentStatementBitwiseAndEqualAST production.
	EnterAssignmentStatementBitwiseAndEqualAST(c *AssignmentStatementBitwiseAndEqualASTContext)

	// EnterAssignmentStatementBitwiseOrEqualAST is called when entering the assignmentStatementBitwiseOrEqualAST production.
	EnterAssignmentStatementBitwiseOrEqualAST(c *AssignmentStatementBitwiseOrEqualASTContext)

	// EnterAssignmentStatementMultiplyEqualAST is called when entering the assignmentStatementMultiplyEqualAST production.
	EnterAssignmentStatementMultiplyEqualAST(c *AssignmentStatementMultiplyEqualASTContext)

	// EnterAssignmentStatementBitwiseXorEqualAST is called when entering the assignmentStatementBitwiseXorEqualAST production.
	EnterAssignmentStatementBitwiseXorEqualAST(c *AssignmentStatementBitwiseXorEqualASTContext)

	// EnterAssignmentStatementShiftLeftEqualAST is called when entering the assignmentStatementShiftLeftEqualAST production.
	EnterAssignmentStatementShiftLeftEqualAST(c *AssignmentStatementShiftLeftEqualASTContext)

	// EnterAssignmentStatementShiftRightEqualAST is called when entering the assignmentStatementShiftRightEqualAST production.
	EnterAssignmentStatementShiftRightEqualAST(c *AssignmentStatementShiftRightEqualASTContext)

	// EnterAssignmentStatementBitwiseClearEqualAST is called when entering the assignmentStatementBitwiseClearEqualAST production.
	EnterAssignmentStatementBitwiseClearEqualAST(c *AssignmentStatementBitwiseClearEqualASTContext)

	// EnterAssignmentStatementModuloEqualAST is called when entering the assignmentStatementModuloEqualAST production.
	EnterAssignmentStatementModuloEqualAST(c *AssignmentStatementModuloEqualASTContext)

	// EnterAssignmentStatementDivideEqualAST is called when entering the assignmentStatementDivideEqualAST production.
	EnterAssignmentStatementDivideEqualAST(c *AssignmentStatementDivideEqualASTContext)

	// EnterIfStatementAST is called when entering the ifStatementAST production.
	EnterIfStatementAST(c *IfStatementASTContext)

	// EnterIfElseIfStatementAST is called when entering the ifElseIfStatementAST production.
	EnterIfElseIfStatementAST(c *IfElseIfStatementASTContext)

	// EnterIfElseStatementAST is called when entering the ifElseStatementAST production.
	EnterIfElseStatementAST(c *IfElseStatementASTContext)

	// EnterIfSimpleStatementAST is called when entering the ifSimpleStatementAST production.
	EnterIfSimpleStatementAST(c *IfSimpleStatementASTContext)

	// EnterIfSimpleElseIfStatementAST is called when entering the ifSimpleElseIfStatementAST production.
	EnterIfSimpleElseIfStatementAST(c *IfSimpleElseIfStatementASTContext)

	// EnterIfSimpleElseStatementAST is called when entering the ifSimpleElseStatementAST production.
	EnterIfSimpleElseStatementAST(c *IfSimpleElseStatementASTContext)

	// EnterLoopBlockAST is called when entering the loopBlockAST production.
	EnterLoopBlockAST(c *LoopBlockASTContext)

	// EnterLoopExpressionBlockAST is called when entering the loopExpressionBlockAST production.
	EnterLoopExpressionBlockAST(c *LoopExpressionBlockASTContext)

	// EnterLoopSimpleStatementExpressionSimpleStatementBlockAST is called when entering the loopSimpleStatementExpressionSimpleStatementBlockAST production.
	EnterLoopSimpleStatementExpressionSimpleStatementBlockAST(c *LoopSimpleStatementExpressionSimpleStatementBlockASTContext)

	// EnterLoopSimpleStatementSimpleStatementBlockAST is called when entering the loopSimpleStatementSimpleStatementBlockAST production.
	EnterLoopSimpleStatementSimpleStatementBlockAST(c *LoopSimpleStatementSimpleStatementBlockASTContext)

	// EnterSwitchStmtSimpleStatementAST is called when entering the switchStmtSimpleStatementAST production.
	EnterSwitchStmtSimpleStatementAST(c *SwitchStmtSimpleStatementASTContext)

	// EnterSwitchStmtExpressionAST is called when entering the switchStmtExpressionAST production.
	EnterSwitchStmtExpressionAST(c *SwitchStmtExpressionASTContext)

	// EnterSwitchStmtSimpleStatementBlockAST is called when entering the switchStmtSimpleStatementBlockAST production.
	EnterSwitchStmtSimpleStatementBlockAST(c *SwitchStmtSimpleStatementBlockASTContext)

	// EnterSwitchStmtBlockAST is called when entering the switchStmtBlockAST production.
	EnterSwitchStmtBlockAST(c *SwitchStmtBlockASTContext)

	// EnterExpressionCaseClauseListEmptyAST is called when entering the expressionCaseClauseListEmptyAST production.
	EnterExpressionCaseClauseListEmptyAST(c *ExpressionCaseClauseListEmptyASTContext)

	// EnterExpressionCaseClauseListAST is called when entering the expressionCaseClauseListAST production.
	EnterExpressionCaseClauseListAST(c *ExpressionCaseClauseListASTContext)

	// EnterExpressionCaseClause is called when entering the expressionCaseClause production.
	EnterExpressionCaseClause(c *ExpressionCaseClauseContext)

	// EnterExpressionSwitchCaseAST is called when entering the expressionSwitchCaseAST production.
	EnterExpressionSwitchCaseAST(c *ExpressionSwitchCaseASTContext)

	// EnterExpressionSwitchDefaultAST is called when entering the expressionSwitchDefaultAST production.
	EnterExpressionSwitchDefaultAST(c *ExpressionSwitchDefaultASTContext)

	// EnterEpsilon is called when entering the epsilon production.
	EnterEpsilon(c *EpsilonContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitTopDeclarationList is called when exiting the topDeclarationList production.
	ExitTopDeclarationList(c *TopDeclarationListContext)

	// ExitVariableDeclAST is called when exiting the variableDeclAST production.
	ExitVariableDeclAST(c *VariableDeclASTContext)

	// ExitVariableDeclBlockAST is called when exiting the variableDeclBlockAST production.
	ExitVariableDeclBlockAST(c *VariableDeclBlockASTContext)

	// ExitVariableDeclEmptyAST is called when exiting the variableDeclEmptyAST production.
	ExitVariableDeclEmptyAST(c *VariableDeclEmptyASTContext)

	// ExitInnerVarDecls is called when exiting the innerVarDecls production.
	ExitInnerVarDecls(c *InnerVarDeclsContext)

	// ExitSingleVarDeclAST is called when exiting the singleVarDeclAST production.
	ExitSingleVarDeclAST(c *SingleVarDeclASTContext)

	// ExitSingleVarDeclAssignAST is called when exiting the singleVarDeclAssignAST production.
	ExitSingleVarDeclAssignAST(c *SingleVarDeclAssignASTContext)

	// ExitSingleVarDeclNoExpsAST is called when exiting the singleVarDeclNoExpsAST production.
	ExitSingleVarDeclNoExpsAST(c *SingleVarDeclNoExpsASTContext)

	// ExitSingleVarDeclNoExps is called when exiting the singleVarDeclNoExps production.
	ExitSingleVarDeclNoExps(c *SingleVarDeclNoExpsContext)

	// ExitTypeDeclAST is called when exiting the typeDeclAST production.
	ExitTypeDeclAST(c *TypeDeclASTContext)

	// ExitTypeDeclBlockAST is called when exiting the typeDeclBlockAST production.
	ExitTypeDeclBlockAST(c *TypeDeclBlockASTContext)

	// ExitTypeDeclEmptyAST is called when exiting the typeDeclEmptyAST production.
	ExitTypeDeclEmptyAST(c *TypeDeclEmptyASTContext)

	// ExitInnerTypeDecls is called when exiting the innerTypeDecls production.
	ExitInnerTypeDecls(c *InnerTypeDeclsContext)

	// ExitSingleTypeDecl is called when exiting the singleTypeDecl production.
	ExitSingleTypeDecl(c *SingleTypeDeclContext)

	// ExitFuncDecl is called when exiting the funcDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitFuncFrontDecl is called when exiting the funcFrontDecl production.
	ExitFuncFrontDecl(c *FuncFrontDeclContext)

	// ExitMultipleReturnTypes is called when exiting the multipleReturnTypes production.
	ExitMultipleReturnTypes(c *MultipleReturnTypesContext)

	// ExitReturnTypeList is called when exiting the returnTypeList production.
	ExitReturnTypeList(c *ReturnTypeListContext)

	// ExitSingleReturnTypeAST is called when exiting the singleReturnTypeAST production.
	ExitSingleReturnTypeAST(c *SingleReturnTypeASTContext)

	// ExitSingleReturnTypeEmptyAST is called when exiting the singleReturnTypeEmptyAST production.
	ExitSingleReturnTypeEmptyAST(c *SingleReturnTypeEmptyASTContext)

	// ExitFuncArgDecls is called when exiting the funcArgDecls production.
	ExitFuncArgDecls(c *FuncArgDeclsContext)

	// ExitDeclTypeParenAST is called when exiting the declTypeParenAST production.
	ExitDeclTypeParenAST(c *DeclTypeParenASTContext)

	// ExitDeclTypeIdentifierAST is called when exiting the declTypeIdentifierAST production.
	ExitDeclTypeIdentifierAST(c *DeclTypeIdentifierASTContext)

	// ExitDeclTypeSliceAST is called when exiting the declTypeSliceAST production.
	ExitDeclTypeSliceAST(c *DeclTypeSliceASTContext)

	// ExitDeclTypeArrayAST is called when exiting the declTypeArrayAST production.
	ExitDeclTypeArrayAST(c *DeclTypeArrayASTContext)

	// ExitDeclTypeStructAST is called when exiting the declTypeStructAST production.
	ExitDeclTypeStructAST(c *DeclTypeStructASTContext)

	// ExitSliceDeclType is called when exiting the sliceDeclType production.
	ExitSliceDeclType(c *SliceDeclTypeContext)

	// ExitArrayDeclType is called when exiting the arrayDeclType production.
	ExitArrayDeclType(c *ArrayDeclTypeContext)

	// ExitStructDeclType is called when exiting the structDeclType production.
	ExitStructDeclType(c *StructDeclTypeContext)

	// ExitStructMemDecls is called when exiting the structMemDecls production.
	ExitStructMemDecls(c *StructMemDeclsContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitExpressionNotUnaryAST is called when exiting the expressionNotUnaryAST production.
	ExitExpressionNotUnaryAST(c *ExpressionNotUnaryASTContext)

	// ExitExpressionMultiplyAST is called when exiting the expressionMultiplyAST production.
	ExitExpressionMultiplyAST(c *ExpressionMultiplyASTContext)

	// ExitExpressionPlusAST is called when exiting the expressionPlusAST production.
	ExitExpressionPlusAST(c *ExpressionPlusASTContext)

	// ExitExpressionModuloAST is called when exiting the expressionModuloAST production.
	ExitExpressionModuloAST(c *ExpressionModuloASTContext)

	// ExitExpressionAndAST is called when exiting the expressionAndAST production.
	ExitExpressionAndAST(c *ExpressionAndASTContext)

	// ExitExpressionBitwiseXorAST is called when exiting the expressionBitwiseXorAST production.
	ExitExpressionBitwiseXorAST(c *ExpressionBitwiseXorASTContext)

	// ExitExpressionMinusAST is called when exiting the expressionMinusAST production.
	ExitExpressionMinusAST(c *ExpressionMinusASTContext)

	// ExitExpressionBitwiseXorUnaryAST is called when exiting the expressionBitwiseXorUnaryAST production.
	ExitExpressionBitwiseXorUnaryAST(c *ExpressionBitwiseXorUnaryASTContext)

	// ExitExpressionEqualAST is called when exiting the expressionEqualAST production.
	ExitExpressionEqualAST(c *ExpressionEqualASTContext)

	// ExitExpressionMinusUnaryAST is called when exiting the expressionMinusUnaryAST production.
	ExitExpressionMinusUnaryAST(c *ExpressionMinusUnaryASTContext)

	// ExitExpressionBitwiseClearAST is called when exiting the expressionBitwiseClearAST production.
	ExitExpressionBitwiseClearAST(c *ExpressionBitwiseClearASTContext)

	// ExitExpressionGreaterEqualAST is called when exiting the expressionGreaterEqualAST production.
	ExitExpressionGreaterEqualAST(c *ExpressionGreaterEqualASTContext)

	// ExitExpressionLessEqualAST is called when exiting the expressionLessEqualAST production.
	ExitExpressionLessEqualAST(c *ExpressionLessEqualASTContext)

	// ExitExpressionNotEqualAST is called when exiting the expressionNotEqualAST production.
	ExitExpressionNotEqualAST(c *ExpressionNotEqualASTContext)

	// ExitExpressionPrimaryAST is called when exiting the expressionPrimaryAST production.
	ExitExpressionPrimaryAST(c *ExpressionPrimaryASTContext)

	// ExitExpressionBitwiseAndAST is called when exiting the expressionBitwiseAndAST production.
	ExitExpressionBitwiseAndAST(c *ExpressionBitwiseAndASTContext)

	// ExitExpressionGreaterAST is called when exiting the expressionGreaterAST production.
	ExitExpressionGreaterAST(c *ExpressionGreaterASTContext)

	// ExitExpressionDivideAST is called when exiting the expressionDivideAST production.
	ExitExpressionDivideAST(c *ExpressionDivideASTContext)

	// ExitExpressionOrAST is called when exiting the expressionOrAST production.
	ExitExpressionOrAST(c *ExpressionOrASTContext)

	// ExitExpressionPlusUnaryAST is called when exiting the expressionPlusUnaryAST production.
	ExitExpressionPlusUnaryAST(c *ExpressionPlusUnaryASTContext)

	// ExitExpressionBitwiseOrAST is called when exiting the expressionBitwiseOrAST production.
	ExitExpressionBitwiseOrAST(c *ExpressionBitwiseOrASTContext)

	// ExitExpressionLessAST is called when exiting the expressionLessAST production.
	ExitExpressionLessAST(c *ExpressionLessASTContext)

	// ExitExpressionShiftRightAST is called when exiting the expressionShiftRightAST production.
	ExitExpressionShiftRightAST(c *ExpressionShiftRightASTContext)

	// ExitExpressionShiftLeftAST is called when exiting the expressionShiftLeftAST production.
	ExitExpressionShiftLeftAST(c *ExpressionShiftLeftASTContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitPrimaryExpressionLengthAST is called when exiting the primaryExpressionLengthAST production.
	ExitPrimaryExpressionLengthAST(c *PrimaryExpressionLengthASTContext)

	// ExitPrimaryExpressionOperandAST is called when exiting the primaryExpressionOperandAST production.
	ExitPrimaryExpressionOperandAST(c *PrimaryExpressionOperandASTContext)

	// ExitPrimaryExpressionIndexAST is called when exiting the primaryExpressionIndexAST production.
	ExitPrimaryExpressionIndexAST(c *PrimaryExpressionIndexASTContext)

	// ExitPrimaryExpressionAppendAST is called when exiting the primaryExpressionAppendAST production.
	ExitPrimaryExpressionAppendAST(c *PrimaryExpressionAppendASTContext)

	// ExitPrimaryExpressionArgumentsAST is called when exiting the primaryExpressionArgumentsAST production.
	ExitPrimaryExpressionArgumentsAST(c *PrimaryExpressionArgumentsASTContext)

	// ExitPrimaryExpressionCapAST is called when exiting the primaryExpressionCapAST production.
	ExitPrimaryExpressionCapAST(c *PrimaryExpressionCapASTContext)

	// ExitPrimaryExpressionSelectorAST is called when exiting the primaryExpressionSelectorAST production.
	ExitPrimaryExpressionSelectorAST(c *PrimaryExpressionSelectorASTContext)

	// ExitOperandLiteralAST is called when exiting the operandLiteralAST production.
	ExitOperandLiteralAST(c *OperandLiteralASTContext)

	// ExitOperandIdentifierAST is called when exiting the operandIdentifierAST production.
	ExitOperandIdentifierAST(c *OperandIdentifierASTContext)

	// ExitOperandParenAST is called when exiting the operandParenAST production.
	ExitOperandParenAST(c *OperandParenASTContext)

	// ExitLiteralIntAST is called when exiting the literalIntAST production.
	ExitLiteralIntAST(c *LiteralIntASTContext)

	// ExitLiteralFloatAST is called when exiting the literalFloatAST production.
	ExitLiteralFloatAST(c *LiteralFloatASTContext)

	// ExitLiteralRuneAST is called when exiting the literalRuneAST production.
	ExitLiteralRuneAST(c *LiteralRuneASTContext)

	// ExitLiteralRawStringAST is called when exiting the literalRawStringAST production.
	ExitLiteralRawStringAST(c *LiteralRawStringASTContext)

	// ExitLiteralInterpretedStringAST is called when exiting the literalInterpretedStringAST production.
	ExitLiteralInterpretedStringAST(c *LiteralInterpretedStringASTContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitArgumentsAST is called when exiting the argumentsAST production.
	ExitArgumentsAST(c *ArgumentsASTContext)

	// ExitArgumentsEmptyAST is called when exiting the argumentsEmptyAST production.
	ExitArgumentsEmptyAST(c *ArgumentsEmptyASTContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitAppendExpression is called when exiting the appendExpression production.
	ExitAppendExpression(c *AppendExpressionContext)

	// ExitLengthExpression is called when exiting the lengthExpression production.
	ExitLengthExpression(c *LengthExpressionContext)

	// ExitCapExpression is called when exiting the capExpression production.
	ExitCapExpression(c *CapExpressionContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatementPrintAST is called when exiting the statementPrintAST production.
	ExitStatementPrintAST(c *StatementPrintASTContext)

	// ExitStatementPrintlnAST is called when exiting the statementPrintlnAST production.
	ExitStatementPrintlnAST(c *StatementPrintlnASTContext)

	// ExitStatementReturnAST is called when exiting the statementReturnAST production.
	ExitStatementReturnAST(c *StatementReturnASTContext)

	// ExitStatementBreakAST is called when exiting the statementBreakAST production.
	ExitStatementBreakAST(c *StatementBreakASTContext)

	// ExitStatementContinueAST is called when exiting the statementContinueAST production.
	ExitStatementContinueAST(c *StatementContinueASTContext)

	// ExitStatementSimpleAST is called when exiting the statementSimpleAST production.
	ExitStatementSimpleAST(c *StatementSimpleASTContext)

	// ExitStatementBlockAST is called when exiting the statementBlockAST production.
	ExitStatementBlockAST(c *StatementBlockASTContext)

	// ExitStatementSwitchAST is called when exiting the statementSwitchAST production.
	ExitStatementSwitchAST(c *StatementSwitchASTContext)

	// ExitStatementIfAST is called when exiting the statementIfAST production.
	ExitStatementIfAST(c *StatementIfASTContext)

	// ExitStatementLoopAST is called when exiting the statementLoopAST production.
	ExitStatementLoopAST(c *StatementLoopASTContext)

	// ExitStatementTypeDeclAST is called when exiting the statementTypeDeclAST production.
	ExitStatementTypeDeclAST(c *StatementTypeDeclASTContext)

	// ExitStatementVariableDeclAST is called when exiting the statementVariableDeclAST production.
	ExitStatementVariableDeclAST(c *StatementVariableDeclASTContext)

	// ExitSimpleStatementEmptyAST is called when exiting the simpleStatementEmptyAST production.
	ExitSimpleStatementEmptyAST(c *SimpleStatementEmptyASTContext)

	// ExitSimpleStatementExpressionAST is called when exiting the simpleStatementExpressionAST production.
	ExitSimpleStatementExpressionAST(c *SimpleStatementExpressionASTContext)

	// ExitSimpleStatementAssignmentAST is called when exiting the simpleStatementAssignmentAST production.
	ExitSimpleStatementAssignmentAST(c *SimpleStatementAssignmentASTContext)

	// ExitSimpleStatementExpressionListAssignAST is called when exiting the simpleStatementExpressionListAssignAST production.
	ExitSimpleStatementExpressionListAssignAST(c *SimpleStatementExpressionListAssignASTContext)

	// ExitAssignmentStatementAST is called when exiting the assignmentStatementAST production.
	ExitAssignmentStatementAST(c *AssignmentStatementASTContext)

	// ExitAssignmentStatementPlusEqualAST is called when exiting the assignmentStatementPlusEqualAST production.
	ExitAssignmentStatementPlusEqualAST(c *AssignmentStatementPlusEqualASTContext)

	// ExitAssignmentStatementMinusEqualAST is called when exiting the assignmentStatementMinusEqualAST production.
	ExitAssignmentStatementMinusEqualAST(c *AssignmentStatementMinusEqualASTContext)

	// ExitAssignmentStatementBitwiseAndEqualAST is called when exiting the assignmentStatementBitwiseAndEqualAST production.
	ExitAssignmentStatementBitwiseAndEqualAST(c *AssignmentStatementBitwiseAndEqualASTContext)

	// ExitAssignmentStatementBitwiseOrEqualAST is called when exiting the assignmentStatementBitwiseOrEqualAST production.
	ExitAssignmentStatementBitwiseOrEqualAST(c *AssignmentStatementBitwiseOrEqualASTContext)

	// ExitAssignmentStatementMultiplyEqualAST is called when exiting the assignmentStatementMultiplyEqualAST production.
	ExitAssignmentStatementMultiplyEqualAST(c *AssignmentStatementMultiplyEqualASTContext)

	// ExitAssignmentStatementBitwiseXorEqualAST is called when exiting the assignmentStatementBitwiseXorEqualAST production.
	ExitAssignmentStatementBitwiseXorEqualAST(c *AssignmentStatementBitwiseXorEqualASTContext)

	// ExitAssignmentStatementShiftLeftEqualAST is called when exiting the assignmentStatementShiftLeftEqualAST production.
	ExitAssignmentStatementShiftLeftEqualAST(c *AssignmentStatementShiftLeftEqualASTContext)

	// ExitAssignmentStatementShiftRightEqualAST is called when exiting the assignmentStatementShiftRightEqualAST production.
	ExitAssignmentStatementShiftRightEqualAST(c *AssignmentStatementShiftRightEqualASTContext)

	// ExitAssignmentStatementBitwiseClearEqualAST is called when exiting the assignmentStatementBitwiseClearEqualAST production.
	ExitAssignmentStatementBitwiseClearEqualAST(c *AssignmentStatementBitwiseClearEqualASTContext)

	// ExitAssignmentStatementModuloEqualAST is called when exiting the assignmentStatementModuloEqualAST production.
	ExitAssignmentStatementModuloEqualAST(c *AssignmentStatementModuloEqualASTContext)

	// ExitAssignmentStatementDivideEqualAST is called when exiting the assignmentStatementDivideEqualAST production.
	ExitAssignmentStatementDivideEqualAST(c *AssignmentStatementDivideEqualASTContext)

	// ExitIfStatementAST is called when exiting the ifStatementAST production.
	ExitIfStatementAST(c *IfStatementASTContext)

	// ExitIfElseIfStatementAST is called when exiting the ifElseIfStatementAST production.
	ExitIfElseIfStatementAST(c *IfElseIfStatementASTContext)

	// ExitIfElseStatementAST is called when exiting the ifElseStatementAST production.
	ExitIfElseStatementAST(c *IfElseStatementASTContext)

	// ExitIfSimpleStatementAST is called when exiting the ifSimpleStatementAST production.
	ExitIfSimpleStatementAST(c *IfSimpleStatementASTContext)

	// ExitIfSimpleElseIfStatementAST is called when exiting the ifSimpleElseIfStatementAST production.
	ExitIfSimpleElseIfStatementAST(c *IfSimpleElseIfStatementASTContext)

	// ExitIfSimpleElseStatementAST is called when exiting the ifSimpleElseStatementAST production.
	ExitIfSimpleElseStatementAST(c *IfSimpleElseStatementASTContext)

	// ExitLoopBlockAST is called when exiting the loopBlockAST production.
	ExitLoopBlockAST(c *LoopBlockASTContext)

	// ExitLoopExpressionBlockAST is called when exiting the loopExpressionBlockAST production.
	ExitLoopExpressionBlockAST(c *LoopExpressionBlockASTContext)

	// ExitLoopSimpleStatementExpressionSimpleStatementBlockAST is called when exiting the loopSimpleStatementExpressionSimpleStatementBlockAST production.
	ExitLoopSimpleStatementExpressionSimpleStatementBlockAST(c *LoopSimpleStatementExpressionSimpleStatementBlockASTContext)

	// ExitLoopSimpleStatementSimpleStatementBlockAST is called when exiting the loopSimpleStatementSimpleStatementBlockAST production.
	ExitLoopSimpleStatementSimpleStatementBlockAST(c *LoopSimpleStatementSimpleStatementBlockASTContext)

	// ExitSwitchStmtSimpleStatementAST is called when exiting the switchStmtSimpleStatementAST production.
	ExitSwitchStmtSimpleStatementAST(c *SwitchStmtSimpleStatementASTContext)

	// ExitSwitchStmtExpressionAST is called when exiting the switchStmtExpressionAST production.
	ExitSwitchStmtExpressionAST(c *SwitchStmtExpressionASTContext)

	// ExitSwitchStmtSimpleStatementBlockAST is called when exiting the switchStmtSimpleStatementBlockAST production.
	ExitSwitchStmtSimpleStatementBlockAST(c *SwitchStmtSimpleStatementBlockASTContext)

	// ExitSwitchStmtBlockAST is called when exiting the switchStmtBlockAST production.
	ExitSwitchStmtBlockAST(c *SwitchStmtBlockASTContext)

	// ExitExpressionCaseClauseListEmptyAST is called when exiting the expressionCaseClauseListEmptyAST production.
	ExitExpressionCaseClauseListEmptyAST(c *ExpressionCaseClauseListEmptyASTContext)

	// ExitExpressionCaseClauseListAST is called when exiting the expressionCaseClauseListAST production.
	ExitExpressionCaseClauseListAST(c *ExpressionCaseClauseListASTContext)

	// ExitExpressionCaseClause is called when exiting the expressionCaseClause production.
	ExitExpressionCaseClause(c *ExpressionCaseClauseContext)

	// ExitExpressionSwitchCaseAST is called when exiting the expressionSwitchCaseAST production.
	ExitExpressionSwitchCaseAST(c *ExpressionSwitchCaseASTContext)

	// ExitExpressionSwitchDefaultAST is called when exiting the expressionSwitchDefaultAST production.
	ExitExpressionSwitchDefaultAST(c *ExpressionSwitchDefaultASTContext)

	// ExitEpsilon is called when exiting the epsilon production.
	ExitEpsilon(c *EpsilonContext)
}
