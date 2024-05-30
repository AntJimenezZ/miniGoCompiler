package encoder

import (
	"Proyecto_Compiladores/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"strconv"
)

var v parser.MiniGoParserVisitor = &EncoderLLVM{}

type EncoderLLVM struct {
	*parser.BaseMiniGoParserVisitor
	mainModule   *ir.Module
	aquiStr      *ir.Global
	intFormat    *ir.Global
	floatFormat  *ir.Global
	stringFormat *ir.Global
	runeFormat   *ir.Global
}

func NewEncoderLLVM() *EncoderLLVM {
	value := &EncoderLLVM{
		BaseMiniGoParserVisitor: &parser.BaseMiniGoParserVisitor{},
		mainModule:              ir.NewModule(),
		aquiStr:                 nil,
		intFormat:               nil,
		floatFormat:             nil,
		stringFormat:            nil,
		runeFormat:              nil,
	}
	value.intFormat = value.mainModule.NewGlobalDef("", constant.NewCharArrayFromString("%d\n\x00"))
	value.floatFormat = value.mainModule.NewGlobalDef("", constant.NewCharArrayFromString("%f\n\x00"))
	value.stringFormat = value.mainModule.NewGlobalDef("", constant.NewCharArrayFromString("%s\n\x00"))
	value.runeFormat = value.mainModule.NewGlobalDef("", constant.NewCharArrayFromString("%c\n\x00"))
	printFunction = value.mainModule.NewFunc("printf", types.I32, ir.NewParam("", types.I8Ptr))
	printFunction.Sig.Variadic = true
	return value
}

func (v *EncoderLLVM) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *EncoderLLVM) VisitChildren(tree antlr.RuleNode) any {
	var result any
	n := tree.GetChildCount()
	for i := 0; i < n; i++ {
		c := tree.GetChild(i)
		val := c.(antlr.ParseTree)
		result2 := result
		result = v.Visit(val)
		if result == nil {
			result = result2
		}
	}
	return result
}

var funcActual *ir.Func
var generalStackBlocks GeneralStack
var funcStackBlocks funcStack
var ifStackBlocks ifStack
var forStackBlocks forStack
var localVariables BlockVariableTable
var typeConversion map[string]types.Type = map[string]types.Type{
	"int":    types.I32,
	"float":  types.Float,
	"string": types.I8Ptr,
	"rune":   types.I8,
	"bool":   types.I1,
}
var printFunction *ir.Func

func (v *EncoderLLVM) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (v *EncoderLLVM) VisitErrorNode(node antlr.ErrorNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitRoot(ctx *parser.RootContext) interface{} {
	localVariables = *NewBlockVariableTable()
	v.VisitChildren(ctx)
	return v.mainModule
}

func (v *EncoderLLVM) VisitTopDeclarationList(ctx *parser.TopDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitVariableDeclAST(ctx *parser.VariableDeclASTContext) interface{} {
	//TODO implement me
	fmt.Println("VariableDeclAST", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitVariableDeclBlockAST(ctx *parser.VariableDeclBlockASTContext) interface{} {
	innerVarDecls := ctx.InnerVarDecls().(*parser.InnerVarDeclsContext)
	decls := innerVarDecls.AllSingleVarDecl()

	for _, decl := range decls {
		switch decl := decl.(type) {
		case *parser.SingleVarDeclASTContext:
			v.VisitSingleVarDeclAST(decl)
		case *parser.SingleVarDeclAssignASTContext:
			v.VisitSingleVarDeclAssignAST(decl)
		case *parser.SingleVarDeclNoExpsASTContext:
			v.VisitSingleVarDeclNoExpsAST(decl)
		}
	}
	return nil
}

func (v *EncoderLLVM) VisitVariableDeclEmptyAST(ctx *parser.VariableDeclEmptyASTContext) interface{} {
	return nil
}

func (v *EncoderLLVM) VisitInnerVarDecls(ctx *parser.InnerVarDeclsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSingleVarDeclAST(ctx *parser.SingleVarDeclASTContext) interface{} {
	if generalStackBlocks.IsEmpty() {
		ids := ctx.IdentifierList().(*parser.IdentifierListContext).AllIDENTIFIER()
		exprs := ctx.ExpressionList().(*parser.ExpressionListContext).AllExpression()
		for i := 0; i < len(ids); i++ {
			value := v.Visit(exprs[i]).(value.Value)
			v.mainModule.NewGlobalDef(ids[i].GetText(), value.(constant.Constant))
		}
	} else {
		blockactual, _ := generalStackBlocks.Peek()
		ids := ctx.IdentifierList().(*parser.IdentifierListContext).AllIDENTIFIER()
		exprs := ctx.ExpressionList().(*parser.ExpressionListContext).AllExpression()
		for i := 0; i < len(ids); i++ {
			value := v.Visit(exprs[i]).(value.Value)
			vActual := blockactual.NewAlloca(value.Type())
			blockactual.NewStore(value, vActual)
			localVariables.AddVariable(blockactual, ids[i].GetText(), vActual)
		}
	}
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitSingleVarDeclAssignAST(ctx *parser.SingleVarDeclAssignASTContext) interface{} {
	if generalStackBlocks.IsEmpty() {
		ids := ctx.IdentifierList().(*parser.IdentifierListContext).AllIDENTIFIER()
		exprs := ctx.ExpressionList().(*parser.ExpressionListContext).AllExpression()
		for i := 0; i < len(ids); i++ {
			value := v.Visit(exprs[i]).(value.Value)
			v.mainModule.NewGlobalDef(ids[i].GetText(), value.(constant.Constant))
		}
	} else {
		blockactual, _ := generalStackBlocks.Peek()
		ids := ctx.IdentifierList().(*parser.IdentifierListContext).AllIDENTIFIER()
		exprs := ctx.ExpressionList().(*parser.ExpressionListContext).AllExpression()
		for i := 0; i < len(ids); i++ {
			value := v.Visit(exprs[i]).(value.Value)
			vActual := blockactual.NewAlloca(value.Type())
			blockactual.NewStore(value, vActual)
			localVariables.AddVariable(blockactual, ids[i].GetText(), vActual)
		}
	}
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitSingleVarDeclNoExpsAST(ctx *parser.SingleVarDeclNoExpsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitSingleVarDeclNoExps(ctx *parser.SingleVarDeclNoExpsContext) interface{} {
	if generalStackBlocks.IsEmpty() {
		ids := ctx.IdentifierList().(*parser.IdentifierListContext).AllIDENTIFIER()
		for i := 0; i < len(ids); i++ {
			value := v.Visit(ctx.DeclType()).(types.Type)
			v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewZeroInitializer(value))
		}
	} else {
		blockactual, _ := generalStackBlocks.Peek()
		ids := ctx.IdentifierList().(*parser.IdentifierListContext).AllIDENTIFIER()
		for i := 0; i < len(ids); i++ {
			value := v.Visit(ctx.DeclType()).(types.Type)
			vActual := blockactual.NewAlloca(value)
			blockactual.NewStore(constant.NewZeroInitializer(value), vActual)
			localVariables.AddVariable(blockactual, ids[i].GetText(), vActual)
		}
	}
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitTypeDeclAST(ctx *parser.TypeDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitTypeDeclBlockAST(ctx *parser.TypeDeclBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitTypeDeclEmptyAST(ctx *parser.TypeDeclEmptyASTContext) interface{} {
	return nil
}

func (v *EncoderLLVM) VisitInnerTypeDecls(ctx *parser.InnerTypeDeclsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSingleTypeDecl(ctx *parser.SingleTypeDeclContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitFuncFrontDecl(ctx *parser.FuncFrontDeclContext) interface{} {
	funcName := ctx.IDENTIFIER().GetText()

	if ctx.FuncArgDecls() != nil {
		v.Visit(ctx.FuncArgDecls())
	}

	var returnType types.Type
	if ctx.SingleReturnType() != nil {
		returnType = v.Visit(ctx.SingleReturnType()).(types.Type)
	} else if ctx.MultipleReturnTypes() != nil {
		returnType = v.Visit(ctx.MultipleReturnTypes()).(types.Type)
	}

	funcActual = v.mainModule.NewFunc(funcName, returnType)

	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitMultipleReturnTypes(ctx *parser.MultipleReturnTypesContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitReturnTypeList(ctx *parser.ReturnTypeListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSingleReturnTypeAST(ctx *parser.SingleReturnTypeASTContext) interface{} {
	// Obtener el tipo de retorno del declType
	return v.Visit(ctx.DeclType())
}

func (v *EncoderLLVM) VisitSingleReturnTypeEmptyAST(ctx *parser.SingleReturnTypeEmptyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitFuncArgDecls(ctx *parser.FuncArgDeclsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitDeclTypeParenAST(ctx *parser.DeclTypeParenASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitDeclTypeIdentifierAST(ctx *parser.DeclTypeIdentifierASTContext) interface{} {
	return typeConversion[ctx.IDENTIFIER().GetText()]
}

func (v *EncoderLLVM) VisitDeclTypeSliceAST(ctx *parser.DeclTypeSliceASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitDeclTypeArrayAST(ctx *parser.DeclTypeArrayASTContext) interface{} {
	return v.Visit(ctx.ArrayDeclType())
}

func (v *EncoderLLVM) VisitDeclTypeStructAST(ctx *parser.DeclTypeStructASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSliceDeclType(ctx *parser.SliceDeclTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitArrayDeclType(ctx *parser.ArrayDeclTypeContext) interface{} {

	arrayType := v.Visit(ctx.DeclType()).(types.Type)
	size, _ := strconv.Atoi(ctx.INTLITERAL().GetText())

	return types.NewArray(uint64(size), arrayType)
}

func (v *EncoderLLVM) VisitStructDeclType(ctx *parser.StructDeclTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStructMemDecls(ctx *parser.StructMemDeclsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitIdentifierList(ctx *parser.IdentifierListContext) interface{} {
	return ctx.AllIDENTIFIER()
}

func (v *EncoderLLVM) VisitExpressionNotUnaryAST(ctx *parser.ExpressionNotUnaryASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionMultiplyAST(ctx *parser.ExpressionMultiplyASTContext) interface{} {

	exprs := ctx.AllExpression()
	valExpression1 := v.Visit(exprs[0]).(value.Value)
	valExpression2 := v.Visit(exprs[1]).(value.Value)

	blockActual, _ := generalStackBlocks.Peek()
	valReturn := blockActual.NewMul(valExpression1, valExpression2)
	return valReturn
}

func (v *EncoderLLVM) VisitExpressionPlusAST(ctx *parser.ExpressionPlusASTContext) interface{} {
	exprs := ctx.AllExpression()
	valExpression1 := v.Visit(exprs[0]).(value.Value)
	valExpression2 := v.Visit(exprs[1]).(value.Value)

	blockActual, _ := generalStackBlocks.Peek()
	valReturn := blockActual.NewAdd(valExpression1, valExpression2)
	return valReturn
}

func (v *EncoderLLVM) VisitExpressionModuloAST(ctx *parser.ExpressionModuloASTContext) interface{} {
	exprs := ctx.AllExpression()
	valExpression1 := v.Visit(exprs[0]).(value.Value)
	valExpression2 := v.Visit(exprs[1]).(value.Value)

	blockActual, _ := generalStackBlocks.Peek()
	valReturn := blockActual.NewSRem(valExpression1, valExpression2)
	return valReturn
}

func (v *EncoderLLVM) VisitExpressionAndAST(ctx *parser.ExpressionAndASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionBitwiseXorAST(ctx *parser.ExpressionBitwiseXorASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionMinusAST(ctx *parser.ExpressionMinusASTContext) interface{} {
	exprs := ctx.AllExpression()
	valExpression1 := v.Visit(exprs[0]).(value.Value)
	valExpression2 := v.Visit(exprs[1]).(value.Value)

	blockActual, _ := generalStackBlocks.Peek()
	valReturn := blockActual.NewSub(valExpression1, valExpression2)
	return valReturn
}

func (v *EncoderLLVM) VisitExpressionBitwiseXorUnaryAST(ctx *parser.ExpressionBitwiseXorUnaryASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionEqualAST(ctx *parser.ExpressionEqualASTContext) interface{} {
	valExpression1 := v.Visit(ctx.Expression(0)).(value.Value)
	valExpression2 := v.Visit(ctx.Expression(1)).(value.Value)
	blockActual, _ := generalStackBlocks.Peek()
	valReturn := blockActual.NewICmp(enum.IPredEQ, valExpression1, valExpression2)

	return valReturn
}

func (v *EncoderLLVM) VisitExpressionMinusUnaryAST(ctx *parser.ExpressionMinusUnaryASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionBitwiseClearAST(ctx *parser.ExpressionBitwiseClearASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionGreaterEqualAST(ctx *parser.ExpressionGreaterEqualASTContext) interface{} {
	valExpression1 := v.Visit(ctx.Expression(0)).(value.Value)
	valExpression2 := v.Visit(ctx.Expression(1)).(value.Value)
	blockActual, _ := generalStackBlocks.Peek()
	valReturn := blockActual.NewICmp(enum.IPredSGE, valExpression1, valExpression2)

	return valReturn
}

func (v *EncoderLLVM) VisitExpressionLessEqualAST(ctx *parser.ExpressionLessEqualASTContext) interface{} {
	valExpression1 := v.Visit(ctx.Expression(0)).(value.Value)
	valExpression2 := v.Visit(ctx.Expression(1)).(value.Value)
	blockActual, _ := generalStackBlocks.Peek()
	valReturn := blockActual.NewICmp(enum.IPredSLE, valExpression1, valExpression2)

	return valReturn
}

func (v *EncoderLLVM) VisitExpressionNotEqualAST(ctx *parser.ExpressionNotEqualASTContext) interface{} {
	valExpression1 := v.Visit(ctx.Expression(0)).(value.Value)
	valExpression2 := v.Visit(ctx.Expression(1)).(value.Value)
	blockActual, _ := generalStackBlocks.Peek()
	valReturn := blockActual.NewICmp(enum.IPredNE, valExpression1, valExpression2)

	return valReturn
}

func (v *EncoderLLVM) VisitExpressionPrimaryAST(ctx *parser.ExpressionPrimaryASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitExpressionBitwiseAndAST(ctx *parser.ExpressionBitwiseAndASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionGreaterAST(ctx *parser.ExpressionGreaterASTContext) interface{} {
	valExpression1 := v.Visit(ctx.Expression(0)).(value.Value)
	valExpression2 := v.Visit(ctx.Expression(1)).(value.Value)
	blockActual, _ := generalStackBlocks.Peek()
	valReturn := blockActual.NewICmp(enum.IPredSGT, valExpression1, valExpression2)

	return valReturn
}

func (v *EncoderLLVM) VisitExpressionDivideAST(ctx *parser.ExpressionDivideASTContext) interface{} {
	exprs := ctx.AllExpression()
	valExpression1 := v.Visit(exprs[0]).(value.Value)
	valExpression2 := v.Visit(exprs[1]).(value.Value)

	blockActual, _ := generalStackBlocks.Peek()
	valReturn := blockActual.NewSDiv(valExpression1, valExpression2)
	return valReturn
}

func (v *EncoderLLVM) VisitExpressionOrAST(ctx *parser.ExpressionOrASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionPlusUnaryAST(ctx *parser.ExpressionPlusUnaryASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionBitwiseOrAST(ctx *parser.ExpressionBitwiseOrASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionLessAST(ctx *parser.ExpressionLessASTContext) interface{} {
	valExpression1 := v.Visit(ctx.Expression(0)).(value.Value)
	valExpression2 := v.Visit(ctx.Expression(1)).(value.Value)
	blockActual, _ := generalStackBlocks.Peek()
	valReturn := blockActual.NewICmp(enum.IPredSLT, valExpression1, valExpression2)

	return valReturn
}

func (v *EncoderLLVM) VisitExpressionShiftRightAST(ctx *parser.ExpressionShiftRightASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionShiftLeftAST(ctx *parser.ExpressionShiftLeftASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitPrimaryExpressionLengthAST(ctx *parser.PrimaryExpressionLengthASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitPrimaryExpressionOperandAST(ctx *parser.PrimaryExpressionOperandASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitPrimaryExpressionIndexAST(ctx *parser.PrimaryExpressionIndexASTContext) interface{} {
	//primaryExp := v.Visit(ctx.PrimaryExpression()).(value.Value)
	//primaryExpName := ctx.PrimaryExpression().GetText()
	//index := v.Visit(ctx.Index()).(value.Value)
	//finalArray := constant.NewArray()
	panic("implement me")
}

func (v *EncoderLLVM) VisitPrimaryExpressionAppendAST(ctx *parser.PrimaryExpressionAppendASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitPrimaryExpressionArgumentsAST(ctx *parser.PrimaryExpressionArgumentsASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitPrimaryExpressionCapAST(ctx *parser.PrimaryExpressionCapASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitPrimaryExpressionSelectorAST(ctx *parser.PrimaryExpressionSelectorASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitOperandLiteralAST(ctx *parser.OperandLiteralASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitOperandIdentifierAST(ctx *parser.OperandIdentifierASTContext) interface{} {
	blockActual, _ := generalStackBlocks.Peek()
	val, _ := localVariables.GetVariable(blockActual, ctx.IDENTIFIER().GetText())
	valReturn := blockActual.NewLoad(types.I32, val)
	return valReturn
}

func (v *EncoderLLVM) VisitOperandParenAST(ctx *parser.OperandParenASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitLiteralIntAST(ctx *parser.LiteralIntASTContext) interface{} {
	valReturn, _ := constant.NewIntFromString(types.I32, ctx.INTLITERAL().GetText())
	return valReturn
}

func (v *EncoderLLVM) VisitLiteralFloatAST(ctx *parser.LiteralFloatASTContext) interface{} {
	valReturn, _ := constant.NewFloatFromString(types.Float, ctx.FLOATLITERAL().GetText())
	return valReturn
}

func (v *EncoderLLVM) VisitLiteralRuneAST(ctx *parser.LiteralRuneASTContext) interface{} {
	valReturn, _ := constant.NewIntFromString(types.I8, ctx.RUNELITERAL().GetText())
	return valReturn
}

func (v *EncoderLLVM) VisitLiteralRawStringAST(ctx *parser.LiteralRawStringASTContext) interface{} {
	valReturnRaw := constant.NewCharArrayFromString(ctx.RAWSTRINGLITERAL().GetText())
	return valReturnRaw
}

func (v *EncoderLLVM) VisitLiteralInterpretedStringAST(ctx *parser.LiteralInterpretedStringASTContext) interface{} {
	valReturn := constant.NewCharArrayFromString(ctx.INTERPRETEDSTRINGLITERAL().GetText())
	return valReturn
}

func (v *EncoderLLVM) VisitIndex(ctx *parser.IndexContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitArgumentsAST(ctx *parser.ArgumentsASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitArgumentsEmptyAST(ctx *parser.ArgumentsEmptyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSelector(ctx *parser.SelectorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAppendExpression(ctx *parser.AppendExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitLengthExpression(ctx *parser.LengthExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitCapExpression(ctx *parser.CapExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementList(ctx *parser.StatementListContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitBlock(ctx *parser.BlockContext) interface{} {
	generalStackBlocks.Push(funcActual.NewBlock(""))

	v.VisitChildren(ctx)

	returnBlock, _ := generalStackBlocks.Peek()
	return returnBlock
}

func (v *EncoderLLVM) VisitStatementPrintAST(ctx *parser.StatementPrintASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementPrintlnAST(ctx *parser.StatementPrintlnASTContext) interface{} {
	exprList := ctx.ExpressionList().(*parser.ExpressionListContext).AllExpression()

	for _, expr := range exprList {
		value := v.Visit(expr).(value.Value)
		blockActual, _ := generalStackBlocks.Peek()
		if value.Type() == types.I32 {
			blockActual.NewCall(printFunction, v.intFormat, value)
		} else if value.Type() == types.Float {
			blockActual.NewCall(printFunction, v.floatFormat, value)
		} else if value.Type() == types.I8Ptr {
			blockActual.NewCall(printFunction, v.stringFormat, value)
		} else if value.Type() == types.I8 {
			blockActual.NewCall(printFunction, v.runeFormat, value)
		}
	}
	return nil
}

func (v *EncoderLLVM) VisitStatementReturnAST(ctx *parser.StatementReturnASTContext) interface{} {
	if generalStackBlocks.Size() == 1 {
		blockCuerpo, _ := generalStackBlocks.Peek()
		generalStackBlocks.Push(funcActual.NewBlock(""))
		blockEnd, _ := generalStackBlocks.Peek()
		blockCuerpo.NewBr(blockEnd)
	}
	blockActual, _ := generalStackBlocks.Peek()

	retValue := v.Visit(ctx.Expression())
	return blockActual.NewRet(retValue.(value.Value))
}

func (v *EncoderLLVM) VisitStatementBreakAST(ctx *parser.StatementBreakASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementContinueAST(ctx *parser.StatementContinueASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementSimpleAST(ctx *parser.StatementSimpleASTContext) interface{} {
	/*statement := ctx.SimpleStatement().(*parser.SimpleStatementAssignmentASTContext)
	return v.Visit(statement)*/
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementBlockAST(ctx *parser.StatementBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementSwitchAST(ctx *parser.StatementSwitchASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementIfAST(ctx *parser.StatementIfASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitStatementLoopAST(ctx *parser.StatementLoopASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementTypeDeclAST(ctx *parser.StatementTypeDeclASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementVariableDeclAST(ctx *parser.StatementVariableDeclASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitSimpleStatementEmptyAST(ctx *parser.SimpleStatementEmptyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSimpleStatementExpressionAST(ctx *parser.SimpleStatementExpressionASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSimpleStatementAssignmentAST(ctx *parser.SimpleStatementAssignmentASTContext) interface{} {
	/*assignment := ctx.AssignmentStatement().(*parser.AssignmentStatementASTContext)
	return v.Visit(assignment)*/
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSimpleStatementExpressionListAssignAST(ctx *parser.SimpleStatementExpressionListAssignASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementAST(ctx *parser.AssignmentStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementPlusEqualAST(ctx *parser.AssignmentStatementPlusEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementMinusEqualAST(ctx *parser.AssignmentStatementMinusEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementBitwiseAndEqualAST(ctx *parser.AssignmentStatementBitwiseAndEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementBitwiseOrEqualAST(ctx *parser.AssignmentStatementBitwiseOrEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementMultiplyEqualAST(ctx *parser.AssignmentStatementMultiplyEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementBitwiseXorEqualAST(ctx *parser.AssignmentStatementBitwiseXorEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementShiftLeftEqualAST(ctx *parser.AssignmentStatementShiftLeftEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementShiftRightEqualAST(ctx *parser.AssignmentStatementShiftRightEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementBitwiseClearEqualAST(ctx *parser.AssignmentStatementBitwiseClearEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementModuloEqualAST(ctx *parser.AssignmentStatementModuloEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementDivideEqualAST(ctx *parser.AssignmentStatementDivideEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitIfStatementAST(ctx *parser.IfStatementASTContext) interface{} {
	blockAnteriorIf, _ := generalStackBlocks.Peek()
	valCond := v.Visit(ctx.Expression()).(value.Value)

	v.Visit(ctx.Block())

	blockCuerpoIf, _ := generalStackBlocks.Peek()

	generalStackBlocks.Push(funcActual.NewBlock(""))
	blockEndIf, _ := generalStackBlocks.Peek()
	blockCuerpoIf.NewBr(blockEndIf)

	blockAnteriorIf.NewCondBr(valCond, blockCuerpoIf, blockEndIf)

	return nil
}

func (v *EncoderLLVM) VisitIfElseIfStatementAST(ctx *parser.IfElseIfStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitIfElseStatementAST(ctx *parser.IfElseStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitIfSimpleStatementAST(ctx *parser.IfSimpleStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitIfSimpleElseIfStatementAST(ctx *parser.IfSimpleElseIfStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitIfSimpleElseStatementAST(ctx *parser.IfSimpleElseStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitLoopBlockAST(ctx *parser.LoopBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitLoopExpressionBlockAST(ctx *parser.LoopExpressionBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitLoopSimpleStatementExpressionSimpleStatementBlockAST(ctx *parser.LoopSimpleStatementExpressionSimpleStatementBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitLoopSimpleStatementSimpleStatementBlockAST(ctx *parser.LoopSimpleStatementSimpleStatementBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSwitchStmtSimpleStatementAST(ctx *parser.SwitchStmtSimpleStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSwitchStmtExpressionAST(ctx *parser.SwitchStmtExpressionASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSwitchStmtSimpleStatementBlockAST(ctx *parser.SwitchStmtSimpleStatementBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSwitchStmtBlockAST(ctx *parser.SwitchStmtBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionCaseClauseListEmptyAST(ctx *parser.ExpressionCaseClauseListEmptyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionCaseClauseListAST(ctx *parser.ExpressionCaseClauseListASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionCaseClause(ctx *parser.ExpressionCaseClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionSwitchCaseAST(ctx *parser.ExpressionSwitchCaseASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionSwitchDefaultAST(ctx *parser.ExpressionSwitchDefaultASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitEpsilon(ctx *parser.EpsilonContext) interface{} {
	return v.VisitChildren(ctx)
}
