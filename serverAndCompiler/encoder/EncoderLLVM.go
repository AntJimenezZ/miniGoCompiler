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
)

var v parser.MiniGoParserVisitor = &EncoderLLVM{}

type EncoderLLVM struct {
	*parser.BaseMiniGoParserVisitor
	mainModule *ir.Module
	aquiStr    *ir.Global
}

func NewEncoderLLVM() *EncoderLLVM {
	return &EncoderLLVM{
		BaseMiniGoParserVisitor: &parser.BaseMiniGoParserVisitor{},
		mainModule:              ir.NewModule(),
		aquiStr:                 nil,
	}
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

var funcActual *ir.Func //se usa para la función que se esté visitando de turno... solo una a la vez

var blocksFunActual Stack               //se usa almacenar los bloques de la función que se esté visitando de turno
var variablesLocales BlockVariableTable //se usa para almacenar las variables locales del bloque de turno. Variables X Bloque
//esta última tabla de variables locales tiene que pensarse mejor porque funciona bien cuando hablamos de identificadores declarados en el mismo bloque donde se usan
//si estamos hablando de identificadores más globales al bloque donde se usan, entonces no creo que funcione tan bien o al menos debería de pensarse mejor y cambiarse

func (v *EncoderLLVM) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (v *EncoderLLVM) VisitErrorNode(node antlr.ErrorNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitRoot(ctx *parser.RootContext) interface{} {
	//se inicializa la estructura de variables locales
	variablesLocales = *NewBlockVariableTable()
	v.aquiStr = v.mainModule.NewGlobalDef("aquiStr", constant.NewCharArrayFromString("Aqui\n\x00"))
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
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitVariableDeclEmptyAST(ctx *parser.VariableDeclEmptyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitInnerVarDecls(ctx *parser.InnerVarDeclsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitSingleVarDeclAST(ctx *parser.SingleVarDeclASTContext) interface{} {
	if blocksFunActual.IsEmpty() {
		ids := ctx.IdentifierList().(*parser.IdentifierListContext).AllIDENTIFIER()
		for i := 0; i < len(ids); i++ {
			if v.Visit(ctx.DeclType()) == "int" {
				v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewInt(types.I32, 0))
			} else if v.Visit(ctx.DeclType()) == "float" {
				v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewFloat(types.Float, 0.0))
			} else if v.Visit(ctx.DeclType()) == "string" {
				v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewCharArrayFromString(""))
			} else if v.Visit(ctx.DeclType()) == "rune" {
				v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewInt(types.I8, 0))
			}
			//v.mainModule.NewGlobalDef(ctx.IdentifierList().GetText(), constant.NewInt(types.I32, 0))
		}
	} else {
		blockactual, _ := blocksFunActual.Peek()

		//se obtienen todos los identificadores de la declaración para declararlos todos
		ids := ctx.IdentifierList().(*parser.IdentifierListContext).AllIDENTIFIER()

		// se recorre la lista para generar el código para solicitar memoria para el identificador
		// se usa quemado el tipo entero... debe obenerse del contexto
		for i := 0; i < len(ids); i++ {
			vActual := blockactual.NewAlloca(types.I32) //es por la prueba... depende del tipo de declaración
			blockactual.NewStore(constant.NewInt(types.I32, 0), vActual)
			//se agrega a una tabla de variablesXbloque la nueva variable creada
			variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
		}
	}
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitSingleVarDeclAssignAST(ctx *parser.SingleVarDeclAssignASTContext) interface{} {
	//TODO implement me

	panic("implement me")
}

func (v *EncoderLLVM) VisitSingleVarDeclNoExpsAST(ctx *parser.SingleVarDeclNoExpsASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitSingleVarDeclNoExps(ctx *parser.SingleVarDeclNoExpsContext) interface{} {
	if blocksFunActual.IsEmpty() {
		ids := ctx.IdentifierList().(*parser.IdentifierListContext).AllIDENTIFIER()
		for i := 0; i < len(ids); i++ {
			if v.Visit(ctx.DeclType()) == "int" {
				v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewInt(types.I32, 0))
			} else if v.Visit(ctx.DeclType()) == "float" {
				v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewFloat(types.Float, 0.0))
			} else if v.Visit(ctx.DeclType()) == "string" {
				v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewCharArrayFromString(""))
			} else if v.Visit(ctx.DeclType()) == "rune" {
				v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewInt(types.I8, 0))
			}
			//v.mainModule.NewGlobalDef(ctx.IdentifierList().GetText(), constant.NewInt(types.I32, 0))
		}
	} else {
		blockactual, _ := blocksFunActual.Peek()

		//se obtienen todos los identificadores de la declaración para declararlos todos
		ids := ctx.IdentifierList().(*parser.IdentifierListContext).AllIDENTIFIER()

		// se recorre la lista para generar el código para solicitar memoria para el identificador
		// se usa quemado el tipo entero... debe obenerse del contexto
		for i := 0; i < len(ids); i++ {
			vActual := blockactual.NewAlloca(types.I32) //es por la prueba... depende del tipo de declaración
			blockactual.NewStore(constant.NewInt(types.I32, 0), vActual)
			//se agrega a una tabla de variablesXbloque la nueva variable creada
			variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
		}
	}
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitTypeDeclAST(ctx *parser.TypeDeclASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitTypeDeclBlockAST(ctx *parser.TypeDeclBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitTypeDeclEmptyAST(ctx *parser.TypeDeclEmptyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
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
	// Obtener el nombre de la función
	funcName := ctx.IDENTIFIER().GetText()

	// Obtener los argumentos (si existen) y procesarlos
	if ctx.FuncArgDecls() != nil {
		v.Visit(ctx.FuncArgDecls())
	}

	// Procesar el tipo de retorno
	var returnType interface{}
	if ctx.SingleReturnType() != nil {
		returnType = v.Visit(ctx.SingleReturnType()).(string)
	} else if ctx.MultipleReturnTypes() != nil {
		// Manejo de múltiples tipos de retorno si es necesario en el futuro
		returnType = v.Visit(ctx.MultipleReturnTypes())
	}

	// Crear la función con el tipo de retorno procesado
	if returnType == "int" {
		funcActual = v.mainModule.NewFunc(funcName, types.I32)
	} else if returnType == "float" {
		funcActual = v.mainModule.NewFunc(funcName, types.Float)
	} else if returnType == "string" {
		funcActual = v.mainModule.NewFunc(funcName, types.I8Ptr)
	}

	// Visitar los hijos para procesar el bloque de la función
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
	return ctx.IDENTIFIER().GetText()
}

func (v *EncoderLLVM) VisitDeclTypeSliceAST(ctx *parser.DeclTypeSliceASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitDeclTypeArrayAST(ctx *parser.DeclTypeArrayASTContext) interface{} {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionPlusAST(ctx *parser.ExpressionPlusASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionModuloAST(ctx *parser.ExpressionModuloASTContext) interface{} {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionBitwiseXorUnaryAST(ctx *parser.ExpressionBitwiseXorUnaryASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionEqualAST(ctx *parser.ExpressionEqualASTContext) interface{} {
	//una comparación tiene dos expresiones que se deben visitar para obtener los Values
	valExpression1 := v.Visit(ctx.Expression(0)).(value.Value)
	valExpression2 := v.Visit(ctx.Expression(1)).(value.Value)
	//se obtiene el bloque actual de la pila de bloques
	blockActual, _ := blocksFunActual.Peek()
	//se crea la instrucción de comparación en el bloque actual esperando que sea el correcto y se devuelve
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
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionLessEqualAST(ctx *parser.ExpressionLessEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionNotEqualAST(ctx *parser.ExpressionNotEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionPrimaryAST(ctx *parser.ExpressionPrimaryASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EncoderLLVM) VisitExpressionBitwiseAndAST(ctx *parser.ExpressionBitwiseAndASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionGreaterAST(ctx *parser.ExpressionGreaterASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitExpressionDivideAST(ctx *parser.ExpressionDivideASTContext) interface{} {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
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
	//se busca en el almacen de variables locales del bloque actual la variable del contexto para devolverla
	//TODO: se debe considerar qué pasa si la variable no fue declarado en el bloque del tope, por ejemplo variables globales al contexto o parámetros
	blockActual, _ := blocksFunActual.Peek()
	val, _ := variablesLocales.GetVariable(blockActual, ctx.IDENTIFIER().GetText())
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
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitLiteralInterpretedStringAST(ctx *parser.LiteralInterpretedStringASTContext) interface{} {
	//TODO implement me
	panic("implement me")
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
	//se agrega el bloque actual creado a la lista de bloques.. sin nombre
	blocksFunActual.Push(funcActual.NewBlock(""))

	//se visita el contenido del Block
	v.VisitChildren(ctx)

	//que se devuelva el bloque creado para conocerlo a la vuelta del visit

	returnBlock, _ := blocksFunActual.Peek()
	return returnBlock
}

func (v *EncoderLLVM) VisitStatementPrintAST(ctx *parser.StatementPrintASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementPrintlnAST(ctx *parser.StatementPrintlnASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *EncoderLLVM) VisitStatementReturnAST(ctx *parser.StatementReturnASTContext) interface{} {
	if blocksFunActual.Size() == 1 {
		blockCuerpo, _ := blocksFunActual.Peek()
		blocksFunActual.Push(funcActual.NewBlock(""))
		blockEnd, _ := blocksFunActual.Peek()
		blockCuerpo.NewBr(blockEnd)
	}
	//se obtiene el bloque actual de la pila de bloques
	blockActual, _ := blocksFunActual.Peek()
	//se escribe la instrucción en el bloque correspondiente...
	//el tipo de retorno es quemado a nil..
	//debe cambiarse para que retorne el valor que se obtenga del ctx

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
	//se respada el bloque que precede al IF
	blockAnteriorIf, _ := blocksFunActual.Peek()
	//se visita la expresión para generar el código de la comparación
	valCond := v.Visit(ctx.Expression()).(value.Value)

	//blockAnteriorIf.NewLoad(types.I32, vActual)
	//valExpression2 := constant.NewInt(types.I32, 10)

	//se crea el cloque para el cuerpo del if...
	//luego se debe volver a poner las instrucciones del comparación y saldo al bloque Padre
	v.Visit(ctx.Block())

	//se respalda el bloque del cuerpo del IF
	blockCuerpoIf, _ := blocksFunActual.Peek()

	//se crea el bloque para el cierre del if
	blocksFunActual.Push(funcActual.NewBlock(""))
	blockEndIf, _ := blocksFunActual.Peek()
	//una vez creado el bloque de cierre, se debe agregar la instrucción de salto al final del bloque del cuerpo del if para terminar dicho bloque
	blockCuerpoIf.NewBr(blockEndIf)

	//se debe agregar la instrucción de salto al final del bloque anterior al if para que se ejecute el bloque de cierre
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
