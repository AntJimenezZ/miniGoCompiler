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
var declaringFunc = false
var forbool = false
var funcArgs []funcArg

func (v *EncoderLLVM) getValue(val value.Value) value.Value {
	if val, ok := val.(*ir.InstAlloca); ok {
		blockActual, _ := generalStackBlocks.Peek()
		return blockActual.NewLoad(val.ElemType, val)
	}
	return val
}

func getPointer(value value.Value, block *ir.Block) value.Value {
	if p, ok := value.Type().(*types.PointerType); ok {
		val := block.NewLoad(p.ElemType, value)
		return val
	}
	return value
}

type funcArg struct {
	name string
	typ  types.Type
}

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
			loadedValue := blockactual.NewLoad(value.Type(), vActual)
			return loadedValue
		}
	}
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitSingleVarDeclNoExpsAST(ctx *parser.SingleVarDeclNoExpsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitSingleVarDeclNoExps(ctx *parser.SingleVarDeclNoExpsContext) interface{} {
	if generalStackBlocks.IsEmpty() || declaringFunc {
		ids := ctx.IdentifierList().(*parser.IdentifierListContext).AllIDENTIFIER()
		value := v.Visit(ctx.DeclType()).(types.Type)
		for i := 0; i < len(ids); i++ {
			if !declaringFunc {
				v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewZeroInitializer(value))
			} else {
				funcArgs = append(funcArgs, funcArg{name: ids[i].GetText(), typ: value})
			}
		}
		return funcArgs
	} else {
		blockactual, _ := generalStackBlocks.Peek()
		ids := ctx.IdentifierList().(*parser.IdentifierListContext).AllIDENTIFIER()
		value := v.Visit(ctx.DeclType()).(types.Type)
		for i := 0; i < len(ids); i++ {
			vActual := blockactual.NewAlloca(value)
			blockactual.NewStore(constant.NewZeroInitializer(value), vActual)
			localVariables.AddVariable(blockactual, ids[i].GetText(), vActual)
		}
	}
	return nil
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
	v.Visit(ctx.FuncFrontDecl())
	v.Visit(ctx.Block())
	block, _ := generalStackBlocks.Peek()
	if block.Term == nil {
		block.NewRet(nil)
	}
	return nil
}

func (v *EncoderLLVM) VisitFuncFrontDecl(ctx *parser.FuncFrontDeclContext) interface{} {
	funcName := ctx.IDENTIFIER().GetText()

	var returnType types.Type
	if ctx.SingleReturnType() != nil {
		returnType = v.Visit(ctx.SingleReturnType()).(types.Type)
		funcActual = v.mainModule.NewFunc(funcName, returnType)
	} else if ctx.MultipleReturnTypes() != nil {
		returnType = v.Visit(ctx.MultipleReturnTypes()).(types.Type)
		funcActual = v.mainModule.NewFunc(funcName, returnType)
	} else {
		funcActual = v.mainModule.NewFunc(funcName, types.Void)
	}

	declaringFunc = true
	block := funcActual.NewBlock("")
	generalStackBlocks.Push(block)
	if ctx.FuncArgDecls() != nil {
		args, ok := v.Visit(ctx.FuncArgDecls()).([]funcArg)
		if ok {
			for _, arg := range args {
				param := ir.NewParam("", arg.typ)
				funcActual.Params = append(funcActual.Params, param)
				alloca := block.NewAlloca(arg.typ)
				block.NewStore(param, alloca)
				localVariables.AddVariable(block, arg.name, alloca)
			}
			funcArgs = []funcArg{}
		}
	}
	declaringFunc = false
	return funcActual
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
	return v.Visit(ctx.DeclType())
}

func (v *EncoderLLVM) VisitSingleReturnTypeEmptyAST(ctx *parser.SingleReturnTypeEmptyASTContext) interface{} {
	return types.Void
}

func (v *EncoderLLVM) VisitFuncArgDecls(ctx *parser.FuncArgDeclsContext) interface{} {
	return v.VisitChildren(ctx)
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
	blockActual, _ := generalStackBlocks.Peek()
	valExpression1 := getPointer(v.Visit(ctx.Expression(0)).(value.Value), blockActual)
	valExpression2 := getPointer(v.Visit(ctx.Expression(1)).(value.Value), blockActual)
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
	blockActual, _ := generalStackBlocks.Peek()
	valExpression1 := getPointer(v.Visit(ctx.Expression(0)).(value.Value), blockActual)
	valExpression2 := getPointer(v.Visit(ctx.Expression(1)).(value.Value), blockActual)
	valReturn := blockActual.NewICmp(enum.IPredSGE, valExpression1, valExpression2)

	return valReturn
}

func (v *EncoderLLVM) VisitExpressionLessEqualAST(ctx *parser.ExpressionLessEqualASTContext) interface{} {
	blockActual, _ := generalStackBlocks.Peek()
	valExpression1 := getPointer(v.Visit(ctx.Expression(0)).(value.Value), blockActual)
	valExpression2 := getPointer(v.Visit(ctx.Expression(1)).(value.Value), blockActual)
	valReturn := blockActual.NewICmp(enum.IPredSLE, valExpression1, valExpression2)

	return valReturn
}

func (v *EncoderLLVM) VisitExpressionNotEqualAST(ctx *parser.ExpressionNotEqualASTContext) interface{} {
	blockActual, _ := generalStackBlocks.Peek()
	valExpression1 := getPointer(v.Visit(ctx.Expression(0)).(value.Value), blockActual)
	valExpression2 := getPointer(v.Visit(ctx.Expression(1)).(value.Value), blockActual)
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
	blockActual, _ := generalStackBlocks.Peek()
	valExpression1 := getPointer(v.Visit(ctx.Expression(0)).(value.Value), blockActual)
	valExpression2 := getPointer(v.Visit(ctx.Expression(1)).(value.Value), blockActual)
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
	blockActual, _ := generalStackBlocks.Peek()
	valExpression1 := getPointer(v.Visit(ctx.Expression(0)).(value.Value), blockActual)
	valExpression2 := getPointer(v.Visit(ctx.Expression(1)).(value.Value), blockActual)
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
	return v.Visit(ctx.LengthExpression()).(value.Value)
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
	val, _ := localVariables.GetVariable(generalStackBlocks, ctx.IDENTIFIER().GetText())
	return val
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
	exprValue := v.Visit(ctx.Expression()).(value.Value)

	arraySize := exprValue.Type().(*types.PointerType).ElemType.(*types.ArrayType).Len

	return constant.NewInt(types.I32, int64(arraySize))
}

func (v *EncoderLLVM) VisitCapExpression(ctx *parser.CapExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.VisitChildren(ctx)
	return nil
}

func (v *EncoderLLVM) VisitStatementPrintAST(ctx *parser.StatementPrintASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementPrintlnAST(ctx *parser.StatementPrintlnASTContext) interface{} {
	exprList := ctx.ExpressionList().(*parser.ExpressionListContext).AllExpression()
	for _, expr := range exprList {
		value := v.getValue(v.Visit(expr).(value.Value))
		blockActual, _ := generalStackBlocks.Peek()
		if value.Type() == types.I32 {
			blockActual.NewCall(printFunction, v.intFormat, value)
		} else if value.Type() == types.Float {
			blockActual.NewCall(printFunction, v.floatFormat, value)
		} else if value.Type().Equal(types.I8) {
			blockActual.NewCall(printFunction, v.stringFormat, value)
		} else if value.Type() == types.I8 {
			blockActual.NewCall(printFunction, v.runeFormat, value)
		}
	}
	return nil
}

func (v *EncoderLLVM) VisitStatementReturnAST(ctx *parser.StatementReturnASTContext) interface{} {
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
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
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
	blockActual, _ := generalStackBlocks.Peek()
	tempAlloca, _ := localVariables.GetVariable(generalStackBlocks, ctx.Expression().GetText())
	tempLoad := blockActual.NewLoad(tempAlloca.Type().(*types.PointerType).ElemType, tempAlloca)

	if ctx.DECREMENT() != nil {
		tempAllocaDec := blockActual.NewSub(tempLoad, constant.NewInt(types.I32, 1))
		blockActual.NewStore(tempAllocaDec, tempAlloca)
	} else if ctx.INCREMENT() != nil {
		tempAllocaInc := blockActual.NewAdd(tempLoad, constant.NewInt(types.I32, 1))
		blockActual.NewStore(tempAllocaInc, tempAlloca)
	}
	return nil
}

func (v *EncoderLLVM) VisitSimpleStatementAssignmentAST(ctx *parser.SimpleStatementAssignmentASTContext) interface{} {
	assignment := ctx.AssignmentStatement().(*parser.AssignmentStatementASTContext)
	return v.Visit(assignment)
}

func (v *EncoderLLVM) VisitSimpleStatementExpressionListAssignAST(ctx *parser.SimpleStatementExpressionListAssignASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitAssignmentStatementAST(ctx *parser.AssignmentStatementASTContext) interface{} {
	ids := ctx.ExpressionList(0).AllExpression()
	assings := ctx.ExpressionList(1).AllExpression()
	blockActual, _ := generalStackBlocks.Peek()

	for i, id := range ids {
		if forbool {

			alloca := blockActual.NewAlloca(types.I32)
			localVariables.AddVariable(blockActual, id.GetText(), alloca)

		}
		ident := v.Visit(id).(value.Value)
		assign := v.Visit(assings[i]).(value.Value)
		fmt.Println(ident, assign)
		blockActual.NewStore(assign, ident)
	}
	return nil
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
	trueBlock := funcActual.NewBlock("")
	generalStackBlocks.Push(trueBlock)
	v.Visit(ctx.Block())

	generalStackBlocks.Push(funcActual.NewBlock(""))
	blockEndIf, _ := generalStackBlocks.Peek()
	trueBlock.NewBr(blockEndIf)

	blockAnteriorIf.NewCondBr(valCond, trueBlock, blockEndIf)

	return nil
}

func (v *EncoderLLVM) VisitIfElseIfStatementAST(ctx *parser.IfElseIfStatementASTContext) interface{} {
	elseIfCtx := ctx.IfStatement().(*parser.IfStatementASTContext)
	blockAnteriorIf, _ := generalStackBlocks.Peek()
	valCondIf := v.Visit(ctx.Expression()).(value.Value)

	trueBlock := funcActual.NewBlock("")
	generalStackBlocks.Push(trueBlock)
	v.Visit(ctx.Block())

	falseBlock := funcActual.NewBlock("")
	generalStackBlocks.Push(falseBlock)
	v.Visit(elseIfCtx.Block())

	valCondElseIf := v.Visit(elseIfCtx.Expression()).(value.Value)

	generalStackBlocks.Push(funcActual.NewBlock(""))
	blockEndIfElseIf, _ := generalStackBlocks.Peek()

	trueBlock.NewBr(blockEndIfElseIf)

	blockAnteriorIf.NewCondBr(valCondIf, trueBlock, falseBlock)
	falseBlock.NewCondBr(valCondElseIf, falseBlock, blockEndIfElseIf)

	return nil
}

func (v *EncoderLLVM) VisitIfElseStatementAST(ctx *parser.IfElseStatementASTContext) interface{} {
	blockAnteriorIf, _ := generalStackBlocks.Peek()
	valCond := v.Visit(ctx.Expression()).(value.Value)

	trueBlock := funcActual.NewBlock("")
	generalStackBlocks.Push(trueBlock)
	v.Visit(ctx.Block(0))

	falseBlock := funcActual.NewBlock("")
	generalStackBlocks.Push(falseBlock)
	v.Visit(ctx.Block(1))

	generalStackBlocks.Push(funcActual.NewBlock(""))
	blockEndIfElse, _ := generalStackBlocks.Peek()
	trueBlock.NewBr(blockEndIfElse)
	falseBlock.NewBr(blockEndIfElse)

	blockAnteriorIf.NewCondBr(valCond, trueBlock, falseBlock)

	return nil
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
	blockbeforeLoop, _ := generalStackBlocks.Peek()

	infinityLoop := funcActual.NewBlock("")
	generalStackBlocks.Push(infinityLoop)
	v.Visit(ctx.Block())

	blockbeforeLoop.NewBr(infinityLoop)
	infinityLoop.NewBr(infinityLoop)

	generalStackBlocks.Push(funcActual.NewBlock(""))

	return nil
}

func (v *EncoderLLVM) VisitLoopExpressionBlockAST(ctx *parser.LoopExpressionBlockASTContext) interface{} {
	blockBeforeLoop, _ := generalStackBlocks.Peek()
	conditionalBlock := funcActual.NewBlock("")
	generalStackBlocks.Push(conditionalBlock)
	forCond := v.Visit(ctx.Expression()).(value.Value)

	loopBlock := funcActual.NewBlock("")
	generalStackBlocks.Push(loopBlock)
	v.Visit(ctx.Block())

	generalStackBlocks.Push(funcActual.NewBlock(""))
	endBlock, _ := generalStackBlocks.Peek()

	blockBeforeLoop.NewBr(conditionalBlock)
	conditionalBlock.NewBr(loopBlock)
	loopBlock.NewBr(conditionalBlock)

	conditionalBlock.NewCondBr(forCond, conditionalBlock, endBlock)

	return nil

}

func (v *EncoderLLVM) VisitLoopSimpleStatementExpressionSimpleStatementBlockAST(ctx *parser.LoopSimpleStatementExpressionSimpleStatementBlockASTContext) interface{} {
	blockBeforeLoop, _ := generalStackBlocks.Peek()
	forbool = true
	conditionalBlock := funcActual.NewBlock("")
	generalStackBlocks.Push(conditionalBlock)

	v.Visit(ctx.SimpleStatement(0))
	forCond := v.Visit(ctx.Expression()).(value.Value)

	loopBlock := funcActual.NewBlock("")
	generalStackBlocks.Push(loopBlock)

	v.Visit(ctx.SimpleStatement(1))
	v.Visit(ctx.Block())

	generalStackBlocks.Push(funcActual.NewBlock(""))
	endBlock, _ := generalStackBlocks.Peek()

	blockBeforeLoop.NewBr(conditionalBlock)
	loopBlock.NewBr(conditionalBlock)

	conditionalBlock.NewCondBr(forCond, loopBlock, endBlock)

	return nil
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
