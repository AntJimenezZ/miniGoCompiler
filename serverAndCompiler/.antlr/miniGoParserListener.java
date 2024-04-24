// Generated from c:/Users/noni4/Desktop/miniGoCompiler/serverAndCompiler/miniGoParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link miniGoParser}.
 */
public interface miniGoParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by the {@code rootAST}
	 * labeled alternative in {@link miniGoParser#root}.
	 * @param ctx the parse tree
	 */
	void enterRootAST(miniGoParser.RootASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code rootAST}
	 * labeled alternative in {@link miniGoParser#root}.
	 * @param ctx the parse tree
	 */
	void exitRootAST(miniGoParser.RootASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code topDeclarationListAST}
	 * labeled alternative in {@link miniGoParser#topDeclarationList}.
	 * @param ctx the parse tree
	 */
	void enterTopDeclarationListAST(miniGoParser.TopDeclarationListASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code topDeclarationListAST}
	 * labeled alternative in {@link miniGoParser#topDeclarationList}.
	 * @param ctx the parse tree
	 */
	void exitTopDeclarationListAST(miniGoParser.TopDeclarationListASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code variableDeclAST}
	 * labeled alternative in {@link miniGoParser#variableDecl}.
	 * @param ctx the parse tree
	 */
	void enterVariableDeclAST(miniGoParser.VariableDeclASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code variableDeclAST}
	 * labeled alternative in {@link miniGoParser#variableDecl}.
	 * @param ctx the parse tree
	 */
	void exitVariableDeclAST(miniGoParser.VariableDeclASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code variableDeclBlockAST}
	 * labeled alternative in {@link miniGoParser#variableDecl}.
	 * @param ctx the parse tree
	 */
	void enterVariableDeclBlockAST(miniGoParser.VariableDeclBlockASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code variableDeclBlockAST}
	 * labeled alternative in {@link miniGoParser#variableDecl}.
	 * @param ctx the parse tree
	 */
	void exitVariableDeclBlockAST(miniGoParser.VariableDeclBlockASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code variableDeclEmptyAST}
	 * labeled alternative in {@link miniGoParser#variableDecl}.
	 * @param ctx the parse tree
	 */
	void enterVariableDeclEmptyAST(miniGoParser.VariableDeclEmptyASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code variableDeclEmptyAST}
	 * labeled alternative in {@link miniGoParser#variableDecl}.
	 * @param ctx the parse tree
	 */
	void exitVariableDeclEmptyAST(miniGoParser.VariableDeclEmptyASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code innerVarDeclsAST}
	 * labeled alternative in {@link miniGoParser#innerVarDecls}.
	 * @param ctx the parse tree
	 */
	void enterInnerVarDeclsAST(miniGoParser.InnerVarDeclsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code innerVarDeclsAST}
	 * labeled alternative in {@link miniGoParser#innerVarDecls}.
	 * @param ctx the parse tree
	 */
	void exitInnerVarDeclsAST(miniGoParser.InnerVarDeclsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code singleVarDeclAST}
	 * labeled alternative in {@link miniGoParser#singleVarDecl}.
	 * @param ctx the parse tree
	 */
	void enterSingleVarDeclAST(miniGoParser.SingleVarDeclASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code singleVarDeclAST}
	 * labeled alternative in {@link miniGoParser#singleVarDecl}.
	 * @param ctx the parse tree
	 */
	void exitSingleVarDeclAST(miniGoParser.SingleVarDeclASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code singleVarDeclAssignAST}
	 * labeled alternative in {@link miniGoParser#singleVarDecl}.
	 * @param ctx the parse tree
	 */
	void enterSingleVarDeclAssignAST(miniGoParser.SingleVarDeclAssignASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code singleVarDeclAssignAST}
	 * labeled alternative in {@link miniGoParser#singleVarDecl}.
	 * @param ctx the parse tree
	 */
	void exitSingleVarDeclAssignAST(miniGoParser.SingleVarDeclAssignASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code singleVarDeclNoExpsAST}
	 * labeled alternative in {@link miniGoParser#singleVarDecl}.
	 * @param ctx the parse tree
	 */
	void enterSingleVarDeclNoExpsAST(miniGoParser.SingleVarDeclNoExpsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code singleVarDeclNoExpsAST}
	 * labeled alternative in {@link miniGoParser#singleVarDecl}.
	 * @param ctx the parse tree
	 */
	void exitSingleVarDeclNoExpsAST(miniGoParser.SingleVarDeclNoExpsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code singleVarDeclNoExpsDeclTypeAST}
	 * labeled alternative in {@link miniGoParser#singleVarDeclNoExps}.
	 * @param ctx the parse tree
	 */
	void enterSingleVarDeclNoExpsDeclTypeAST(miniGoParser.SingleVarDeclNoExpsDeclTypeASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code singleVarDeclNoExpsDeclTypeAST}
	 * labeled alternative in {@link miniGoParser#singleVarDeclNoExps}.
	 * @param ctx the parse tree
	 */
	void exitSingleVarDeclNoExpsDeclTypeAST(miniGoParser.SingleVarDeclNoExpsDeclTypeASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code typeDeclAST}
	 * labeled alternative in {@link miniGoParser#typeDecl}.
	 * @param ctx the parse tree
	 */
	void enterTypeDeclAST(miniGoParser.TypeDeclASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code typeDeclAST}
	 * labeled alternative in {@link miniGoParser#typeDecl}.
	 * @param ctx the parse tree
	 */
	void exitTypeDeclAST(miniGoParser.TypeDeclASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code typeDeclBlockAST}
	 * labeled alternative in {@link miniGoParser#typeDecl}.
	 * @param ctx the parse tree
	 */
	void enterTypeDeclBlockAST(miniGoParser.TypeDeclBlockASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code typeDeclBlockAST}
	 * labeled alternative in {@link miniGoParser#typeDecl}.
	 * @param ctx the parse tree
	 */
	void exitTypeDeclBlockAST(miniGoParser.TypeDeclBlockASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code typeDeclEmptyAST}
	 * labeled alternative in {@link miniGoParser#typeDecl}.
	 * @param ctx the parse tree
	 */
	void enterTypeDeclEmptyAST(miniGoParser.TypeDeclEmptyASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code typeDeclEmptyAST}
	 * labeled alternative in {@link miniGoParser#typeDecl}.
	 * @param ctx the parse tree
	 */
	void exitTypeDeclEmptyAST(miniGoParser.TypeDeclEmptyASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code innerTypeDeclsAST}
	 * labeled alternative in {@link miniGoParser#innerTypeDecls}.
	 * @param ctx the parse tree
	 */
	void enterInnerTypeDeclsAST(miniGoParser.InnerTypeDeclsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code innerTypeDeclsAST}
	 * labeled alternative in {@link miniGoParser#innerTypeDecls}.
	 * @param ctx the parse tree
	 */
	void exitInnerTypeDeclsAST(miniGoParser.InnerTypeDeclsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code singleTypeDeclAST}
	 * labeled alternative in {@link miniGoParser#singleTypeDecl}.
	 * @param ctx the parse tree
	 */
	void enterSingleTypeDeclAST(miniGoParser.SingleTypeDeclASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code singleTypeDeclAST}
	 * labeled alternative in {@link miniGoParser#singleTypeDecl}.
	 * @param ctx the parse tree
	 */
	void exitSingleTypeDeclAST(miniGoParser.SingleTypeDeclASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code funcDeclAST}
	 * labeled alternative in {@link miniGoParser#funcDecl}.
	 * @param ctx the parse tree
	 */
	void enterFuncDeclAST(miniGoParser.FuncDeclASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code funcDeclAST}
	 * labeled alternative in {@link miniGoParser#funcDecl}.
	 * @param ctx the parse tree
	 */
	void exitFuncDeclAST(miniGoParser.FuncDeclASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code funcFrontDeclAST}
	 * labeled alternative in {@link miniGoParser#funcFrontDecl}.
	 * @param ctx the parse tree
	 */
	void enterFuncFrontDeclAST(miniGoParser.FuncFrontDeclASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code funcFrontDeclAST}
	 * labeled alternative in {@link miniGoParser#funcFrontDecl}.
	 * @param ctx the parse tree
	 */
	void exitFuncFrontDeclAST(miniGoParser.FuncFrontDeclASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code funcArgDeclsAST}
	 * labeled alternative in {@link miniGoParser#funcArgDecls}.
	 * @param ctx the parse tree
	 */
	void enterFuncArgDeclsAST(miniGoParser.FuncArgDeclsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code funcArgDeclsAST}
	 * labeled alternative in {@link miniGoParser#funcArgDecls}.
	 * @param ctx the parse tree
	 */
	void exitFuncArgDeclsAST(miniGoParser.FuncArgDeclsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code declTypeParenAST}
	 * labeled alternative in {@link miniGoParser#declType}.
	 * @param ctx the parse tree
	 */
	void enterDeclTypeParenAST(miniGoParser.DeclTypeParenASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code declTypeParenAST}
	 * labeled alternative in {@link miniGoParser#declType}.
	 * @param ctx the parse tree
	 */
	void exitDeclTypeParenAST(miniGoParser.DeclTypeParenASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code declTypeIdentifierAST}
	 * labeled alternative in {@link miniGoParser#declType}.
	 * @param ctx the parse tree
	 */
	void enterDeclTypeIdentifierAST(miniGoParser.DeclTypeIdentifierASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code declTypeIdentifierAST}
	 * labeled alternative in {@link miniGoParser#declType}.
	 * @param ctx the parse tree
	 */
	void exitDeclTypeIdentifierAST(miniGoParser.DeclTypeIdentifierASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code declTypeSliceAST}
	 * labeled alternative in {@link miniGoParser#declType}.
	 * @param ctx the parse tree
	 */
	void enterDeclTypeSliceAST(miniGoParser.DeclTypeSliceASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code declTypeSliceAST}
	 * labeled alternative in {@link miniGoParser#declType}.
	 * @param ctx the parse tree
	 */
	void exitDeclTypeSliceAST(miniGoParser.DeclTypeSliceASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code declTypeArrayAST}
	 * labeled alternative in {@link miniGoParser#declType}.
	 * @param ctx the parse tree
	 */
	void enterDeclTypeArrayAST(miniGoParser.DeclTypeArrayASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code declTypeArrayAST}
	 * labeled alternative in {@link miniGoParser#declType}.
	 * @param ctx the parse tree
	 */
	void exitDeclTypeArrayAST(miniGoParser.DeclTypeArrayASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code declTypeStructAST}
	 * labeled alternative in {@link miniGoParser#declType}.
	 * @param ctx the parse tree
	 */
	void enterDeclTypeStructAST(miniGoParser.DeclTypeStructASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code declTypeStructAST}
	 * labeled alternative in {@link miniGoParser#declType}.
	 * @param ctx the parse tree
	 */
	void exitDeclTypeStructAST(miniGoParser.DeclTypeStructASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code sliceDeclTypeAST}
	 * labeled alternative in {@link miniGoParser#sliceDeclType}.
	 * @param ctx the parse tree
	 */
	void enterSliceDeclTypeAST(miniGoParser.SliceDeclTypeASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code sliceDeclTypeAST}
	 * labeled alternative in {@link miniGoParser#sliceDeclType}.
	 * @param ctx the parse tree
	 */
	void exitSliceDeclTypeAST(miniGoParser.SliceDeclTypeASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code arrayDeclTypeAST}
	 * labeled alternative in {@link miniGoParser#arrayDeclType}.
	 * @param ctx the parse tree
	 */
	void enterArrayDeclTypeAST(miniGoParser.ArrayDeclTypeASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code arrayDeclTypeAST}
	 * labeled alternative in {@link miniGoParser#arrayDeclType}.
	 * @param ctx the parse tree
	 */
	void exitArrayDeclTypeAST(miniGoParser.ArrayDeclTypeASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code structDeclTypeAST}
	 * labeled alternative in {@link miniGoParser#structDeclType}.
	 * @param ctx the parse tree
	 */
	void enterStructDeclTypeAST(miniGoParser.StructDeclTypeASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code structDeclTypeAST}
	 * labeled alternative in {@link miniGoParser#structDeclType}.
	 * @param ctx the parse tree
	 */
	void exitStructDeclTypeAST(miniGoParser.StructDeclTypeASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code structMemDeclsAST}
	 * labeled alternative in {@link miniGoParser#structMemDecls}.
	 * @param ctx the parse tree
	 */
	void enterStructMemDeclsAST(miniGoParser.StructMemDeclsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code structMemDeclsAST}
	 * labeled alternative in {@link miniGoParser#structMemDecls}.
	 * @param ctx the parse tree
	 */
	void exitStructMemDeclsAST(miniGoParser.StructMemDeclsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code identifierListAST}
	 * labeled alternative in {@link miniGoParser#identifierList}.
	 * @param ctx the parse tree
	 */
	void enterIdentifierListAST(miniGoParser.IdentifierListASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code identifierListAST}
	 * labeled alternative in {@link miniGoParser#identifierList}.
	 * @param ctx the parse tree
	 */
	void exitIdentifierListAST(miniGoParser.IdentifierListASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionNotUnaryAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionNotUnaryAST(miniGoParser.ExpressionNotUnaryASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionNotUnaryAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionNotUnaryAST(miniGoParser.ExpressionNotUnaryASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionMultiplyAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionMultiplyAST(miniGoParser.ExpressionMultiplyASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionMultiplyAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionMultiplyAST(miniGoParser.ExpressionMultiplyASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionPlusAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionPlusAST(miniGoParser.ExpressionPlusASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionPlusAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionPlusAST(miniGoParser.ExpressionPlusASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionModuloAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionModuloAST(miniGoParser.ExpressionModuloASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionModuloAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionModuloAST(miniGoParser.ExpressionModuloASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionAndAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionAndAST(miniGoParser.ExpressionAndASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionAndAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionAndAST(miniGoParser.ExpressionAndASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionBitwiseXorAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionBitwiseXorAST(miniGoParser.ExpressionBitwiseXorASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionBitwiseXorAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionBitwiseXorAST(miniGoParser.ExpressionBitwiseXorASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionMinusAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionMinusAST(miniGoParser.ExpressionMinusASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionMinusAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionMinusAST(miniGoParser.ExpressionMinusASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionBitwiseXorUnaryAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionBitwiseXorUnaryAST(miniGoParser.ExpressionBitwiseXorUnaryASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionBitwiseXorUnaryAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionBitwiseXorUnaryAST(miniGoParser.ExpressionBitwiseXorUnaryASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionEqualAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionEqualAST(miniGoParser.ExpressionEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionEqualAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionEqualAST(miniGoParser.ExpressionEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionMinusUnaryAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionMinusUnaryAST(miniGoParser.ExpressionMinusUnaryASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionMinusUnaryAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionMinusUnaryAST(miniGoParser.ExpressionMinusUnaryASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionBitwiseClearAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionBitwiseClearAST(miniGoParser.ExpressionBitwiseClearASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionBitwiseClearAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionBitwiseClearAST(miniGoParser.ExpressionBitwiseClearASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionGreaterEqualAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionGreaterEqualAST(miniGoParser.ExpressionGreaterEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionGreaterEqualAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionGreaterEqualAST(miniGoParser.ExpressionGreaterEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionLessEqualAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionLessEqualAST(miniGoParser.ExpressionLessEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionLessEqualAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionLessEqualAST(miniGoParser.ExpressionLessEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionNotEqualAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionNotEqualAST(miniGoParser.ExpressionNotEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionNotEqualAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionNotEqualAST(miniGoParser.ExpressionNotEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionPrimaryAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionPrimaryAST(miniGoParser.ExpressionPrimaryASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionPrimaryAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionPrimaryAST(miniGoParser.ExpressionPrimaryASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionBitwiseAndAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionBitwiseAndAST(miniGoParser.ExpressionBitwiseAndASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionBitwiseAndAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionBitwiseAndAST(miniGoParser.ExpressionBitwiseAndASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionGreaterAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionGreaterAST(miniGoParser.ExpressionGreaterASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionGreaterAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionGreaterAST(miniGoParser.ExpressionGreaterASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionDivideAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionDivideAST(miniGoParser.ExpressionDivideASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionDivideAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionDivideAST(miniGoParser.ExpressionDivideASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionOrAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionOrAST(miniGoParser.ExpressionOrASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionOrAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionOrAST(miniGoParser.ExpressionOrASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionPlusUnaryAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionPlusUnaryAST(miniGoParser.ExpressionPlusUnaryASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionPlusUnaryAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionPlusUnaryAST(miniGoParser.ExpressionPlusUnaryASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionBitwiseOrAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionBitwiseOrAST(miniGoParser.ExpressionBitwiseOrASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionBitwiseOrAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionBitwiseOrAST(miniGoParser.ExpressionBitwiseOrASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionLessAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionLessAST(miniGoParser.ExpressionLessASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionLessAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionLessAST(miniGoParser.ExpressionLessASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionShiftRightAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionShiftRightAST(miniGoParser.ExpressionShiftRightASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionShiftRightAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionShiftRightAST(miniGoParser.ExpressionShiftRightASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionShiftLeftAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionShiftLeftAST(miniGoParser.ExpressionShiftLeftASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionShiftLeftAST}
	 * labeled alternative in {@link miniGoParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionShiftLeftAST(miniGoParser.ExpressionShiftLeftASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionListAST}
	 * labeled alternative in {@link miniGoParser#expressionList}.
	 * @param ctx the parse tree
	 */
	void enterExpressionListAST(miniGoParser.ExpressionListASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionListAST}
	 * labeled alternative in {@link miniGoParser#expressionList}.
	 * @param ctx the parse tree
	 */
	void exitExpressionListAST(miniGoParser.ExpressionListASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code primaryExpressionLengthAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpressionLengthAST(miniGoParser.PrimaryExpressionLengthASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code primaryExpressionLengthAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpressionLengthAST(miniGoParser.PrimaryExpressionLengthASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code primaryExpressionOperandAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpressionOperandAST(miniGoParser.PrimaryExpressionOperandASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code primaryExpressionOperandAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpressionOperandAST(miniGoParser.PrimaryExpressionOperandASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code primaryExpressionIndexAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpressionIndexAST(miniGoParser.PrimaryExpressionIndexASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code primaryExpressionIndexAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpressionIndexAST(miniGoParser.PrimaryExpressionIndexASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code primaryExpressionAppendAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpressionAppendAST(miniGoParser.PrimaryExpressionAppendASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code primaryExpressionAppendAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpressionAppendAST(miniGoParser.PrimaryExpressionAppendASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code primaryExpressionArgumentsAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpressionArgumentsAST(miniGoParser.PrimaryExpressionArgumentsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code primaryExpressionArgumentsAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpressionArgumentsAST(miniGoParser.PrimaryExpressionArgumentsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code primaryExpressionCapAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpressionCapAST(miniGoParser.PrimaryExpressionCapASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code primaryExpressionCapAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpressionCapAST(miniGoParser.PrimaryExpressionCapASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code primaryExpressionSelectorAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpressionSelectorAST(miniGoParser.PrimaryExpressionSelectorASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code primaryExpressionSelectorAST}
	 * labeled alternative in {@link miniGoParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpressionSelectorAST(miniGoParser.PrimaryExpressionSelectorASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code operandLiteralAST}
	 * labeled alternative in {@link miniGoParser#operand}.
	 * @param ctx the parse tree
	 */
	void enterOperandLiteralAST(miniGoParser.OperandLiteralASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code operandLiteralAST}
	 * labeled alternative in {@link miniGoParser#operand}.
	 * @param ctx the parse tree
	 */
	void exitOperandLiteralAST(miniGoParser.OperandLiteralASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code operandIdentifierAST}
	 * labeled alternative in {@link miniGoParser#operand}.
	 * @param ctx the parse tree
	 */
	void enterOperandIdentifierAST(miniGoParser.OperandIdentifierASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code operandIdentifierAST}
	 * labeled alternative in {@link miniGoParser#operand}.
	 * @param ctx the parse tree
	 */
	void exitOperandIdentifierAST(miniGoParser.OperandIdentifierASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code operandParenAST}
	 * labeled alternative in {@link miniGoParser#operand}.
	 * @param ctx the parse tree
	 */
	void enterOperandParenAST(miniGoParser.OperandParenASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code operandParenAST}
	 * labeled alternative in {@link miniGoParser#operand}.
	 * @param ctx the parse tree
	 */
	void exitOperandParenAST(miniGoParser.OperandParenASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code literalIntAST}
	 * labeled alternative in {@link miniGoParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterLiteralIntAST(miniGoParser.LiteralIntASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code literalIntAST}
	 * labeled alternative in {@link miniGoParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitLiteralIntAST(miniGoParser.LiteralIntASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code literalFloatAST}
	 * labeled alternative in {@link miniGoParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterLiteralFloatAST(miniGoParser.LiteralFloatASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code literalFloatAST}
	 * labeled alternative in {@link miniGoParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitLiteralFloatAST(miniGoParser.LiteralFloatASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code literalRuneAST}
	 * labeled alternative in {@link miniGoParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterLiteralRuneAST(miniGoParser.LiteralRuneASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code literalRuneAST}
	 * labeled alternative in {@link miniGoParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitLiteralRuneAST(miniGoParser.LiteralRuneASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code literalRawStringAST}
	 * labeled alternative in {@link miniGoParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterLiteralRawStringAST(miniGoParser.LiteralRawStringASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code literalRawStringAST}
	 * labeled alternative in {@link miniGoParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitLiteralRawStringAST(miniGoParser.LiteralRawStringASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code literalInterpretedStringAST}
	 * labeled alternative in {@link miniGoParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterLiteralInterpretedStringAST(miniGoParser.LiteralInterpretedStringASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code literalInterpretedStringAST}
	 * labeled alternative in {@link miniGoParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitLiteralInterpretedStringAST(miniGoParser.LiteralInterpretedStringASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code indexAST}
	 * labeled alternative in {@link miniGoParser#index}.
	 * @param ctx the parse tree
	 */
	void enterIndexAST(miniGoParser.IndexASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code indexAST}
	 * labeled alternative in {@link miniGoParser#index}.
	 * @param ctx the parse tree
	 */
	void exitIndexAST(miniGoParser.IndexASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code argumentsAST}
	 * labeled alternative in {@link miniGoParser#arguments}.
	 * @param ctx the parse tree
	 */
	void enterArgumentsAST(miniGoParser.ArgumentsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code argumentsAST}
	 * labeled alternative in {@link miniGoParser#arguments}.
	 * @param ctx the parse tree
	 */
	void exitArgumentsAST(miniGoParser.ArgumentsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code argumentsEmptyAST}
	 * labeled alternative in {@link miniGoParser#arguments}.
	 * @param ctx the parse tree
	 */
	void enterArgumentsEmptyAST(miniGoParser.ArgumentsEmptyASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code argumentsEmptyAST}
	 * labeled alternative in {@link miniGoParser#arguments}.
	 * @param ctx the parse tree
	 */
	void exitArgumentsEmptyAST(miniGoParser.ArgumentsEmptyASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code selectorAST}
	 * labeled alternative in {@link miniGoParser#selector}.
	 * @param ctx the parse tree
	 */
	void enterSelectorAST(miniGoParser.SelectorASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code selectorAST}
	 * labeled alternative in {@link miniGoParser#selector}.
	 * @param ctx the parse tree
	 */
	void exitSelectorAST(miniGoParser.SelectorASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code appendExpressionAST}
	 * labeled alternative in {@link miniGoParser#appendExpression}.
	 * @param ctx the parse tree
	 */
	void enterAppendExpressionAST(miniGoParser.AppendExpressionASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code appendExpressionAST}
	 * labeled alternative in {@link miniGoParser#appendExpression}.
	 * @param ctx the parse tree
	 */
	void exitAppendExpressionAST(miniGoParser.AppendExpressionASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code lengthExpressionAST}
	 * labeled alternative in {@link miniGoParser#lengthExpression}.
	 * @param ctx the parse tree
	 */
	void enterLengthExpressionAST(miniGoParser.LengthExpressionASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code lengthExpressionAST}
	 * labeled alternative in {@link miniGoParser#lengthExpression}.
	 * @param ctx the parse tree
	 */
	void exitLengthExpressionAST(miniGoParser.LengthExpressionASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code capExpressionAST}
	 * labeled alternative in {@link miniGoParser#capExpression}.
	 * @param ctx the parse tree
	 */
	void enterCapExpressionAST(miniGoParser.CapExpressionASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code capExpressionAST}
	 * labeled alternative in {@link miniGoParser#capExpression}.
	 * @param ctx the parse tree
	 */
	void exitCapExpressionAST(miniGoParser.CapExpressionASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementListAST}
	 * labeled alternative in {@link miniGoParser#statementList}.
	 * @param ctx the parse tree
	 */
	void enterStatementListAST(miniGoParser.StatementListASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementListAST}
	 * labeled alternative in {@link miniGoParser#statementList}.
	 * @param ctx the parse tree
	 */
	void exitStatementListAST(miniGoParser.StatementListASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code blockAST}
	 * labeled alternative in {@link miniGoParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlockAST(miniGoParser.BlockASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code blockAST}
	 * labeled alternative in {@link miniGoParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlockAST(miniGoParser.BlockASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementPrintAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatementPrintAST(miniGoParser.StatementPrintASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementPrintAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatementPrintAST(miniGoParser.StatementPrintASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementPrintlnAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatementPrintlnAST(miniGoParser.StatementPrintlnASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementPrintlnAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatementPrintlnAST(miniGoParser.StatementPrintlnASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementReturnAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatementReturnAST(miniGoParser.StatementReturnASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementReturnAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatementReturnAST(miniGoParser.StatementReturnASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementBreakAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatementBreakAST(miniGoParser.StatementBreakASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementBreakAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatementBreakAST(miniGoParser.StatementBreakASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementContinueAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatementContinueAST(miniGoParser.StatementContinueASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementContinueAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatementContinueAST(miniGoParser.StatementContinueASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementSimpleAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatementSimpleAST(miniGoParser.StatementSimpleASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementSimpleAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatementSimpleAST(miniGoParser.StatementSimpleASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementBlockAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatementBlockAST(miniGoParser.StatementBlockASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementBlockAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatementBlockAST(miniGoParser.StatementBlockASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementSwitchAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatementSwitchAST(miniGoParser.StatementSwitchASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementSwitchAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatementSwitchAST(miniGoParser.StatementSwitchASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementIfAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatementIfAST(miniGoParser.StatementIfASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementIfAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatementIfAST(miniGoParser.StatementIfASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementLoopAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatementLoopAST(miniGoParser.StatementLoopASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementLoopAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatementLoopAST(miniGoParser.StatementLoopASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementTypeDeclAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatementTypeDeclAST(miniGoParser.StatementTypeDeclASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementTypeDeclAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatementTypeDeclAST(miniGoParser.StatementTypeDeclASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementVariableDeclAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatementVariableDeclAST(miniGoParser.StatementVariableDeclASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementVariableDeclAST}
	 * labeled alternative in {@link miniGoParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatementVariableDeclAST(miniGoParser.StatementVariableDeclASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code simpleStatementEmptyAST}
	 * labeled alternative in {@link miniGoParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void enterSimpleStatementEmptyAST(miniGoParser.SimpleStatementEmptyASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code simpleStatementEmptyAST}
	 * labeled alternative in {@link miniGoParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void exitSimpleStatementEmptyAST(miniGoParser.SimpleStatementEmptyASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code simpleStatementExpressionAST}
	 * labeled alternative in {@link miniGoParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void enterSimpleStatementExpressionAST(miniGoParser.SimpleStatementExpressionASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code simpleStatementExpressionAST}
	 * labeled alternative in {@link miniGoParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void exitSimpleStatementExpressionAST(miniGoParser.SimpleStatementExpressionASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code simpleStatementAssignmentAST}
	 * labeled alternative in {@link miniGoParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void enterSimpleStatementAssignmentAST(miniGoParser.SimpleStatementAssignmentASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code simpleStatementAssignmentAST}
	 * labeled alternative in {@link miniGoParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void exitSimpleStatementAssignmentAST(miniGoParser.SimpleStatementAssignmentASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code simpleStatementExpressionListAssignAST}
	 * labeled alternative in {@link miniGoParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void enterSimpleStatementExpressionListAssignAST(miniGoParser.SimpleStatementExpressionListAssignASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code simpleStatementExpressionListAssignAST}
	 * labeled alternative in {@link miniGoParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void exitSimpleStatementExpressionListAssignAST(miniGoParser.SimpleStatementExpressionListAssignASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assignmentStatementAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatementAST(miniGoParser.AssignmentStatementASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assignmentStatementAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatementAST(miniGoParser.AssignmentStatementASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assignmentStatementPlusEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatementPlusEqualAST(miniGoParser.AssignmentStatementPlusEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assignmentStatementPlusEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatementPlusEqualAST(miniGoParser.AssignmentStatementPlusEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assignmentStatementMinusEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatementMinusEqualAST(miniGoParser.AssignmentStatementMinusEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assignmentStatementMinusEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatementMinusEqualAST(miniGoParser.AssignmentStatementMinusEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assignmentStatementBitwiseAndEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatementBitwiseAndEqualAST(miniGoParser.AssignmentStatementBitwiseAndEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assignmentStatementBitwiseAndEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatementBitwiseAndEqualAST(miniGoParser.AssignmentStatementBitwiseAndEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assignmentStatementBitwiseOrEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatementBitwiseOrEqualAST(miniGoParser.AssignmentStatementBitwiseOrEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assignmentStatementBitwiseOrEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatementBitwiseOrEqualAST(miniGoParser.AssignmentStatementBitwiseOrEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assignmentStatementMultiplyEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatementMultiplyEqualAST(miniGoParser.AssignmentStatementMultiplyEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assignmentStatementMultiplyEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatementMultiplyEqualAST(miniGoParser.AssignmentStatementMultiplyEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assignmentStatementBitwiseXorEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatementBitwiseXorEqualAST(miniGoParser.AssignmentStatementBitwiseXorEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assignmentStatementBitwiseXorEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatementBitwiseXorEqualAST(miniGoParser.AssignmentStatementBitwiseXorEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assignmentStatementShiftLeftEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatementShiftLeftEqualAST(miniGoParser.AssignmentStatementShiftLeftEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assignmentStatementShiftLeftEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatementShiftLeftEqualAST(miniGoParser.AssignmentStatementShiftLeftEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assignmentStatementShiftRightEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatementShiftRightEqualAST(miniGoParser.AssignmentStatementShiftRightEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assignmentStatementShiftRightEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatementShiftRightEqualAST(miniGoParser.AssignmentStatementShiftRightEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assignmentStatementBitwiseClearEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatementBitwiseClearEqualAST(miniGoParser.AssignmentStatementBitwiseClearEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assignmentStatementBitwiseClearEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatementBitwiseClearEqualAST(miniGoParser.AssignmentStatementBitwiseClearEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assignmentStatementModuloEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatementModuloEqualAST(miniGoParser.AssignmentStatementModuloEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assignmentStatementModuloEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatementModuloEqualAST(miniGoParser.AssignmentStatementModuloEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assignmentStatementDivideEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatementDivideEqualAST(miniGoParser.AssignmentStatementDivideEqualASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assignmentStatementDivideEqualAST}
	 * labeled alternative in {@link miniGoParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatementDivideEqualAST(miniGoParser.AssignmentStatementDivideEqualASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ifStatementAST}
	 * labeled alternative in {@link miniGoParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatementAST(miniGoParser.IfStatementASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ifStatementAST}
	 * labeled alternative in {@link miniGoParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatementAST(miniGoParser.IfStatementASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ifElseIfStatementAST}
	 * labeled alternative in {@link miniGoParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfElseIfStatementAST(miniGoParser.IfElseIfStatementASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ifElseIfStatementAST}
	 * labeled alternative in {@link miniGoParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfElseIfStatementAST(miniGoParser.IfElseIfStatementASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ifElseStatementAST}
	 * labeled alternative in {@link miniGoParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfElseStatementAST(miniGoParser.IfElseStatementASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ifElseStatementAST}
	 * labeled alternative in {@link miniGoParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfElseStatementAST(miniGoParser.IfElseStatementASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ifSimpleStatementAST}
	 * labeled alternative in {@link miniGoParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfSimpleStatementAST(miniGoParser.IfSimpleStatementASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ifSimpleStatementAST}
	 * labeled alternative in {@link miniGoParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfSimpleStatementAST(miniGoParser.IfSimpleStatementASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ifSimpleElseIfStatementAST}
	 * labeled alternative in {@link miniGoParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfSimpleElseIfStatementAST(miniGoParser.IfSimpleElseIfStatementASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ifSimpleElseIfStatementAST}
	 * labeled alternative in {@link miniGoParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfSimpleElseIfStatementAST(miniGoParser.IfSimpleElseIfStatementASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ifSimpleElseStatementAST}
	 * labeled alternative in {@link miniGoParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfSimpleElseStatementAST(miniGoParser.IfSimpleElseStatementASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ifSimpleElseStatementAST}
	 * labeled alternative in {@link miniGoParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfSimpleElseStatementAST(miniGoParser.IfSimpleElseStatementASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code loopBlockAST}
	 * labeled alternative in {@link miniGoParser#loop}.
	 * @param ctx the parse tree
	 */
	void enterLoopBlockAST(miniGoParser.LoopBlockASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code loopBlockAST}
	 * labeled alternative in {@link miniGoParser#loop}.
	 * @param ctx the parse tree
	 */
	void exitLoopBlockAST(miniGoParser.LoopBlockASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code loopExpressionBlockAST}
	 * labeled alternative in {@link miniGoParser#loop}.
	 * @param ctx the parse tree
	 */
	void enterLoopExpressionBlockAST(miniGoParser.LoopExpressionBlockASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code loopExpressionBlockAST}
	 * labeled alternative in {@link miniGoParser#loop}.
	 * @param ctx the parse tree
	 */
	void exitLoopExpressionBlockAST(miniGoParser.LoopExpressionBlockASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code loopSimpleStatementExpressionSimpleStatementBlockAST}
	 * labeled alternative in {@link miniGoParser#loop}.
	 * @param ctx the parse tree
	 */
	void enterLoopSimpleStatementExpressionSimpleStatementBlockAST(miniGoParser.LoopSimpleStatementExpressionSimpleStatementBlockASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code loopSimpleStatementExpressionSimpleStatementBlockAST}
	 * labeled alternative in {@link miniGoParser#loop}.
	 * @param ctx the parse tree
	 */
	void exitLoopSimpleStatementExpressionSimpleStatementBlockAST(miniGoParser.LoopSimpleStatementExpressionSimpleStatementBlockASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code loopSimpleStatementSimpleStatementBlockAST}
	 * labeled alternative in {@link miniGoParser#loop}.
	 * @param ctx the parse tree
	 */
	void enterLoopSimpleStatementSimpleStatementBlockAST(miniGoParser.LoopSimpleStatementSimpleStatementBlockASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code loopSimpleStatementSimpleStatementBlockAST}
	 * labeled alternative in {@link miniGoParser#loop}.
	 * @param ctx the parse tree
	 */
	void exitLoopSimpleStatementSimpleStatementBlockAST(miniGoParser.LoopSimpleStatementSimpleStatementBlockASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code switchStmtSimpleStatementAST}
	 * labeled alternative in {@link miniGoParser#switchStmt}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStmtSimpleStatementAST(miniGoParser.SwitchStmtSimpleStatementASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code switchStmtSimpleStatementAST}
	 * labeled alternative in {@link miniGoParser#switchStmt}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStmtSimpleStatementAST(miniGoParser.SwitchStmtSimpleStatementASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code switchStmtExpressionAST}
	 * labeled alternative in {@link miniGoParser#switchStmt}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStmtExpressionAST(miniGoParser.SwitchStmtExpressionASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code switchStmtExpressionAST}
	 * labeled alternative in {@link miniGoParser#switchStmt}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStmtExpressionAST(miniGoParser.SwitchStmtExpressionASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code switchStmtSimpleStatementBlockAST}
	 * labeled alternative in {@link miniGoParser#switchStmt}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStmtSimpleStatementBlockAST(miniGoParser.SwitchStmtSimpleStatementBlockASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code switchStmtSimpleStatementBlockAST}
	 * labeled alternative in {@link miniGoParser#switchStmt}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStmtSimpleStatementBlockAST(miniGoParser.SwitchStmtSimpleStatementBlockASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code switchStmtBlockAST}
	 * labeled alternative in {@link miniGoParser#switchStmt}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStmtBlockAST(miniGoParser.SwitchStmtBlockASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code switchStmtBlockAST}
	 * labeled alternative in {@link miniGoParser#switchStmt}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStmtBlockAST(miniGoParser.SwitchStmtBlockASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionCaseClauseListEmptyAST}
	 * labeled alternative in {@link miniGoParser#expressionCaseClauseList}.
	 * @param ctx the parse tree
	 */
	void enterExpressionCaseClauseListEmptyAST(miniGoParser.ExpressionCaseClauseListEmptyASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionCaseClauseListEmptyAST}
	 * labeled alternative in {@link miniGoParser#expressionCaseClauseList}.
	 * @param ctx the parse tree
	 */
	void exitExpressionCaseClauseListEmptyAST(miniGoParser.ExpressionCaseClauseListEmptyASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionCaseClauseListAST}
	 * labeled alternative in {@link miniGoParser#expressionCaseClauseList}.
	 * @param ctx the parse tree
	 */
	void enterExpressionCaseClauseListAST(miniGoParser.ExpressionCaseClauseListASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionCaseClauseListAST}
	 * labeled alternative in {@link miniGoParser#expressionCaseClauseList}.
	 * @param ctx the parse tree
	 */
	void exitExpressionCaseClauseListAST(miniGoParser.ExpressionCaseClauseListASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionCaseClauseAST}
	 * labeled alternative in {@link miniGoParser#expressionCaseClause}.
	 * @param ctx the parse tree
	 */
	void enterExpressionCaseClauseAST(miniGoParser.ExpressionCaseClauseASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionCaseClauseAST}
	 * labeled alternative in {@link miniGoParser#expressionCaseClause}.
	 * @param ctx the parse tree
	 */
	void exitExpressionCaseClauseAST(miniGoParser.ExpressionCaseClauseASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionSwitchCaseAST}
	 * labeled alternative in {@link miniGoParser#expressionSwitchCase}.
	 * @param ctx the parse tree
	 */
	void enterExpressionSwitchCaseAST(miniGoParser.ExpressionSwitchCaseASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionSwitchCaseAST}
	 * labeled alternative in {@link miniGoParser#expressionSwitchCase}.
	 * @param ctx the parse tree
	 */
	void exitExpressionSwitchCaseAST(miniGoParser.ExpressionSwitchCaseASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionSwitchDefaultAST}
	 * labeled alternative in {@link miniGoParser#expressionSwitchCase}.
	 * @param ctx the parse tree
	 */
	void enterExpressionSwitchDefaultAST(miniGoParser.ExpressionSwitchDefaultASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionSwitchDefaultAST}
	 * labeled alternative in {@link miniGoParser#expressionSwitchCase}.
	 * @param ctx the parse tree
	 */
	void exitExpressionSwitchDefaultAST(miniGoParser.ExpressionSwitchDefaultASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code epsilonAST}
	 * labeled alternative in {@link miniGoParser#epsilon}.
	 * @param ctx the parse tree
	 */
	void enterEpsilonAST(miniGoParser.EpsilonASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code epsilonAST}
	 * labeled alternative in {@link miniGoParser#epsilon}.
	 * @param ctx the parse tree
	 */
	void exitEpsilonAST(miniGoParser.EpsilonASTContext ctx);
}