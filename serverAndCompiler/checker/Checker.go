package checker

import (
	parser "Proyecto_Compiladores/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"regexp"
	"strconv"
	"strings"
)

var _ parser.MiniGoParserVisitor = &Checker{}

// Checker realiza el análisis contextual utilizando una table de símbolos

type Checker struct {
	SymbolTable *SymbolTable // Tabla de símbolos utilizada para el análisis contextual
	errorList   []string     // Lista de errores encontrados durante el análisis
}

func (c *Checker) VisitChildren(node antlr.RuleNode) interface{} {

	var result interface{}
	children := node.GetChildren()

	for _, child := range children {
		childResult := child.(antlr.ParseTree).Accept(c)
		result = childResult
	}

	return result
}

func (c *Checker) VisitTerminal(node antlr.TerminalNode) interface{} {

	return nil
}

func (c *Checker) VisitErrorNode(node antlr.ErrorNode) interface{} {
	//TODO implement me
	return c.errorList
}

func (c *Checker) VisitRoot(ctx *parser.RootContext) interface{} {

	funcs := ctx.TopDeclarationList().AllFuncDecl()

	for _, fn := range funcs {
		fun := fn.FuncFrontDecl()
		fmt.Println("FUNC", fun.IDENTIFIER())

		fmt.Println("ARGS", fun.FuncArgDecls().GetText())

	}

	return c.VisitChildren(ctx)
}

func (c *Checker) VisitTopDeclarationList(ctx *parser.TopDeclarationListContext) interface{} {

	fmt.Println("primer lugar de debug")

	return c.VisitChildren(ctx)
}

func (c *Checker) VisitVariableDeclBlockAST(ctx *parser.VariableDeclBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitVariableDeclEmptyAST(ctx *parser.VariableDeclEmptyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitInnerVarDecls(ctx *parser.InnerVarDeclsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitSingleVarDeclAST(ctx *parser.SingleVarDeclASTContext) interface{} {
	idents := ctx.IdentifierList().AllIDENTIFIER()
	expressions := ctx.ExpressionList().AllExpression()
	for i, ident := range idents {
		expression := expressions[i]
		var result = c.Visit(expression).(int)
		//c.SymbolTable.OpenScope()
		c.SymbolTable.InsertVar(ident.GetText(), result)
		//c.SymbolTable.CloseScope()
	}

	return nil
}

func (c *Checker) VisitSingleVarDeclAssignAST(ctx *parser.SingleVarDeclAssignASTContext) interface{} {

	idents := ctx.IdentifierList().AllIDENTIFIER()
	expressions := ctx.ExpressionList().AllExpression()
	for i, ident := range idents {

		fmt.Println("jummm 3", ident)
		expression := expressions[i]

		var result = c.Visit(expression)

		castResult := fmt.Sprintf("%v", result)
		if result, err := strconv.Atoi(castResult); err == nil {
			// Si no hay error, se puede convertir a entero
			c.SymbolTable.InsertVar(ident.GetText(), result)

			// Hacer algo aquí
		} else {
			// Si hay error, no se puede convertir a entero
			fmt.Println("digaleeees: ", ident)
			c.SymbolTable.InsertObject(ident.GetText(), castResult)

			// Hacer otra cosa aquí
		}

	}
	c.SymbolTable.PrintTable()
	c.SymbolTable.ExportTable()

	return nil
}

func (c *Checker) VisitSingleVarDeclNoExpsAST(ctx *parser.SingleVarDeclNoExpsASTContext) interface{} {

	return ctx.GetChildren() //sospechoso
}

func (c *Checker) VisitSingleVarDeclNoExps(ctx *parser.SingleVarDeclNoExpsContext) interface{} {
	//TODO implement me

	fmt.Println("test")
	return c.VisitChildren(ctx)
}

func (c *Checker) VisitTypeDeclAST(ctx *parser.TypeDeclASTContext) interface{} {

	//SUS

	return c.VisitChildren(ctx)
}

func (c *Checker) VisitTypeDeclBlockAST(ctx *parser.TypeDeclBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitTypeDeclEmptyAST(ctx *parser.TypeDeclEmptyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitInnerTypeDecls(ctx *parser.InnerTypeDeclsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitSingleTypeDecl(ctx *parser.SingleTypeDeclContext) interface{} {

	structName := ctx.IDENTIFIER()
	fmt.Println("NOMBRE STRUCT:", structName)

	structValues := ctx.DeclType().GetText()
	structValues = strings.TrimPrefix(structValues, "struct{")
	structValues = strings.TrimSuffix(structValues, ";}")

	newStructValues := strings.Split(structValues, ";")

	var structValuesList []string

	// Agregar cada campo a la slice
	for _, field := range newStructValues {
		// Eliminar espacios en blanco y agregar a la slice
		structValuesList = append(structValuesList, strings.TrimSpace(field))
	}

	fmt.Println("I ESTO: ", structValues)
	fmt.Println("testing: ", structValuesList)

	//METER LOS STRUCTS A LA SYMBOL TABLE, Y DESPUÉS FIJARSE QUE NO SE PUEDA HACER p.edad si es un "string"
	//vamos noni que se puede :V

	structNameString := fmt.Sprintf("%v", structName)

	c.SymbolTable.InsertStruct(structNameString, 7, structValuesList)
	c.SymbolTable.PrintTable()
	c.SymbolTable.ExportTable()
	return c.VisitChildren(ctx)

	//panic("implement me")
}

func (c *Checker) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {

	fmt.Println("Soy peruano: " + ctx.GetText())

	return c.VisitChildren(ctx)
}

func (c *Checker) VisitFuncFrontDecl(ctx *parser.FuncFrontDeclContext) interface{} {

	// Get the function name
	funcName := ctx.IDENTIFIER().GetText()

	// Get the function arguments
	args := ctx.FuncArgDecls().GetText()

	fmt.Println("Func args: " + args)
	fmt.Println("func name: " + funcName)

	return c.VisitChildren(ctx)
}

func (c *Checker) VisitMultipleReturnTypes(ctx *parser.MultipleReturnTypesContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitReturnTypeList(ctx *parser.ReturnTypeListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitSingleReturnTypeAST(ctx *parser.SingleReturnTypeASTContext) interface{} {
	//TODO implement me

	fmt.Println("visit single return typeast: " + ctx.GetText())
	return nil
}

func (c *Checker) VisitSingleReturnTypeEmptyAST(ctx *parser.SingleReturnTypeEmptyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitFuncArgDecls(ctx *parser.FuncArgDeclsContext) interface{} {
	//TODO implement me

	fmt.Println("sexo tilin:" + ctx.GetText())

	return c.VisitChildren(ctx)
	//panic("implement me")
	//TODO WIP

}

func (c *Checker) VisitDeclTypeParenAST(ctx *parser.DeclTypeParenASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitDeclTypeIdentifierAST(ctx *parser.DeclTypeIdentifierASTContext) interface{} {
	//TODO implement me
	fmt.Println("esto que es 2:", ctx.IDENTIFIER())

	return ctx.IDENTIFIER()
}

func (c *Checker) VisitDeclTypeSliceAST(ctx *parser.DeclTypeSliceASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitDeclTypeArrayAST(ctx *parser.DeclTypeArrayASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitDeclTypeStructAST(ctx *parser.DeclTypeStructASTContext) interface{} {

	fmt.Println("struct con llaves1: ", ctx.GetText())

	return c.VisitChildren(ctx)
}

func (c *Checker) VisitSliceDeclType(ctx *parser.SliceDeclTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitArrayDeclType(ctx *parser.ArrayDeclTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitStructDeclType(ctx *parser.StructDeclTypeContext) interface{} {

	fmt.Println("struct con llaves2: ", ctx.GetText())

	return c.VisitChildren(ctx)
	//WIP
	//Aqui lo que ocupo hacer es guardarlo, si uno intenta hacer
	//Person.age = "caca" le tire error porque no es del tipo
}

func (c *Checker) VisitStructMemDecls(ctx *parser.StructMemDeclsContext) interface{} {
	//TODO implement me
	fmt.Println("datos del struct sin llaves:", ctx.GetText())

	//testing := ctx.GetText()

	return c.VisitChildren(ctx)
}

func (c *Checker) VisitExpressionNotUnaryAST(ctx *parser.ExpressionNotUnaryASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionMultiplyAST(ctx *parser.ExpressionMultiplyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionPlusAST(ctx *parser.ExpressionPlusASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionModuloAST(ctx *parser.ExpressionModuloASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionAndAST(ctx *parser.ExpressionAndASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionBitwiseXorAST(ctx *parser.ExpressionBitwiseXorASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionMinusAST(ctx *parser.ExpressionMinusASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionBitwiseXorUnaryAST(ctx *parser.ExpressionBitwiseXorUnaryASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionEqualAST(ctx *parser.ExpressionEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionMinusUnaryAST(ctx *parser.ExpressionMinusUnaryASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionBitwiseClearAST(ctx *parser.ExpressionBitwiseClearASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionGreaterEqualAST(ctx *parser.ExpressionGreaterEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionLessEqualAST(ctx *parser.ExpressionLessEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionNotEqualAST(ctx *parser.ExpressionNotEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionPrimaryAST(ctx *parser.ExpressionPrimaryASTContext) interface{} {

	return c.VisitChildren(ctx)
}

func (c *Checker) VisitExpressionBitwiseAndAST(ctx *parser.ExpressionBitwiseAndASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionGreaterAST(ctx *parser.ExpressionGreaterASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionDivideAST(ctx *parser.ExpressionDivideASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionOrAST(ctx *parser.ExpressionOrASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionPlusUnaryAST(ctx *parser.ExpressionPlusUnaryASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionBitwiseOrAST(ctx *parser.ExpressionBitwiseOrASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionLessAST(ctx *parser.ExpressionLessASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionShiftRightAST(ctx *parser.ExpressionShiftRightASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionShiftLeftAST(ctx *parser.ExpressionShiftLeftASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitPrimaryExpressionLengthAST(ctx *parser.PrimaryExpressionLengthASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitPrimaryExpressionOperandAST(ctx *parser.PrimaryExpressionOperandASTContext) interface{} {

	return c.VisitChildren(ctx)
}

func (c *Checker) VisitPrimaryExpressionIndexAST(ctx *parser.PrimaryExpressionIndexASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitPrimaryExpressionAppendAST(ctx *parser.PrimaryExpressionAppendASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitPrimaryExpressionArgumentsAST(ctx *parser.PrimaryExpressionArgumentsASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitPrimaryExpressionCapAST(ctx *parser.PrimaryExpressionCapASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitPrimaryExpressionSelectorAST(ctx *parser.PrimaryExpressionSelectorASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitOperandLiteralAST(ctx *parser.OperandLiteralASTContext) interface{} {

	return c.VisitChildren(ctx)
}

func (c *Checker) VisitOperandIdentifierAST(ctx *parser.OperandIdentifierASTContext) interface{} {
	//TODO implement me

	fmt.Println("uuuuuuuuuuuuuuuuuuuuu", ctx.GetText())

	fmt.Println("eeeeeeeeeeeee", ctx.IDENTIFIER())

	return ctx.GetText()
	//panic("implement me")
}

func (c *Checker) VisitOperandParenAST(ctx *parser.OperandParenASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitLiteralIntAST(ctx *parser.LiteralIntASTContext) interface{} {
	//TODO implement me
	return 2
}

func (c *Checker) VisitLiteralFloatAST(ctx *parser.LiteralFloatASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitLiteralRuneAST(ctx *parser.LiteralRuneASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitLiteralRawStringAST(ctx *parser.LiteralRawStringASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitLiteralInterpretedStringAST(ctx *parser.LiteralInterpretedStringASTContext) interface{} {

	return String.Type
}

func (c *Checker) VisitIndex(ctx *parser.IndexContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitArgumentsAST(ctx *parser.ArgumentsASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitArgumentsEmptyAST(ctx *parser.ArgumentsEmptyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitSelector(ctx *parser.SelectorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAppendExpression(ctx *parser.AppendExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitLengthExpression(ctx *parser.LengthExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitCapExpression(ctx *parser.CapExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	// Obtén los statements del bloque
	statements := ctx.GetText()

	// Divide los statements en base al punto y coma
	statementsList := strings.Split(statements, ";")

	// Define las palabras clave y símbolos para dividir los statements
	keywords := []string{"var", ":=", "\\+", "-", "\\*", "/", "%", "return", "int", "float", "string", "bool", "print", "println"}

	// Crea una expresión regular que coincida con cualquier palabra clave o símbolo
	re := regexp.MustCompile(`(\s*(` + strings.Join(keywords, "|") + `)\s*)`)

	// Itera sobre los statements divididos
	for _, stmt := range statementsList {
		// Encuentra todos los índices de inicio y fin de las coincidencias en el statement
		indexes := re.FindAllStringIndex(stmt, -1)

		// Divide el statement en base a los índices encontrados
		var splitElements []string
		lastIndex := 0
		for _, indexPair := range indexes {
			splitElements = append(splitElements, stmt[lastIndex:indexPair[0]])
			splitElements = append(splitElements, stmt[indexPair[0]:indexPair[1]])
			lastIndex = indexPair[1]
		}
		splitElements = append(splitElements, stmt[lastIndex:])

		// Elimina elementos vacíos y espacios adicionales
		var cleanedElements []string
		for _, elem := range splitElements {
			if elem != "" && elem != " " {
				cleanedElements = append(cleanedElements, strings.TrimSpace(elem))
			}
		}

		// Une los elementos divididos en un solo string, separado por espacios
		finalStatement := strings.Join(cleanedElements, " ")

		if strings.Contains(finalStatement, "var") {
			// Divide el statement en base a la palabra clave "var"
			varElements := strings.Split(finalStatement, "var ")
			// Itera sobre los elementos divididos
			for _, elem := range varElements {
				// Elimina los espacios adicionales
				elem = strings.TrimSpace(elem)
				// Si el elemento no es vacío
				if elem != "" {
					// Divide el elemento en base a los espacios
					elemSplit := strings.Split(elem, " ")
					// Si el elemento tiene más de 2 partes
					if len(elemSplit) >= 2 {
						// Si la segunda parte del elemento es ":="
						if elemSplit[1] == ":=" {
							// Divide la tercera parte del elemento en base a las comas
							identifier := elemSplit[:1]
							// Si la primera parte del elemento es "int"
							if elemSplit[1] == "bool" {
								for _, id := range identifier {
									c.SymbolTable.InsertVar(id, 1)
								}
							} else if elemSplit[1] == "int" {
								for _, id := range identifier {
									c.SymbolTable.InsertVar(id, 2)
								}
							} else if elemSplit[1] == "float" {
								for _, id := range identifier {
									c.SymbolTable.InsertVar(id, 3)
								}
							} else if elemSplit[1] == "string" {
								for _, id := range identifier {
									c.SymbolTable.InsertVar(id, 4)
								}
							}
						} else {
							// Divide la tercera parte del elemento en base a las comas
							identifier := elemSplit[:1]
							// Si la primera parte del elemento es "int"
							if elemSplit[1] == "bool" {
								for _, id := range identifier {
									c.SymbolTable.InsertVar(id, 1)
								}
							} else if elemSplit[1] == "int" {
								for _, id := range identifier {
									c.SymbolTable.InsertVar(id, 2)
								}
							} else if elemSplit[1] == "float" {
								for _, id := range identifier {
									c.SymbolTable.InsertVar(id, 3)
								}
							} else if elemSplit[1] == "string" {
								for _, id := range identifier {
									c.SymbolTable.InsertVar(id, 4)
								}
							}
						}
					}
				}
			}
		} else if strings.Contains(finalStatement, "return") {
			//comprobar tipo de retorno

			returnElements := strings.Split(finalStatement, "return ")

			returnVar := c.SymbolTable.FindFunc(returnElements[1])
			if returnVar == nil {
				return nil
			}
			if len(returnElements) > 1 {
				returnValues := strings.Split(returnElements[1], ",")
				for _, value := range returnValues {
					if returnVar.Type == 1 {
						fmt.Println("Tipo de retorno: " + value)
					} else if returnVar.Type == 2 {
						fmt.Println("Tipo de retorno: " + value)
					} else if returnVar.Type == 3 {
						fmt.Println("Tipo de retorno: " + value)
					} else if returnVar.Type == 4 {
						fmt.Println("Tipo de retorno: " + value)
					} else if returnVar.Type == 5 {
						fmt.Println("Tipo de retorno: " + value)
					} else {
						fmt.Println("Error: Tipo de retorno incorrecto")
					}
				}
			}
		} else if strings.Contains(finalStatement, "print") {
			// Divide el statement en base a "print"
			printElements := strings.Split(finalStatement, "print")
			// Si el statement tiene más de 1 parte
			if len(printElements) > 1 {
				// Divide la segunda parte del statement en base a las comas
				printValues := strings.Split(printElements[1], ",")
				// Itera sobre los valores divididos
				for _, value := range printValues {
					// Elimina los espacios adicionales
					value = strings.TrimSpace(value)
					value = strings.Trim(value, "()")
					// Si el valor no es vacío
					if value != "" {
						// Si el valor es un número
						if _, err := strconv.Atoi(value); err == nil {
							// Imprime el valor
							fmt.Println(value)
						} else {
							// Si el valor es una variable
							if c.SymbolTable.FindFunc(value) != nil {
								// Imprime el valor de la variable
								fmt.Println(c.SymbolTable.FindFunc(value))
							}
						}
					}
				}
			}

		} else if strings.Contains(finalStatement, "println") {
			// Divide el statement en base a "println"
			printElements := strings.Split(finalStatement, "println")
			// Si el statement tiene más de 1 parte
			if len(printElements) > 1 {
				// Divide la segunda parte del statement en base a las comas
				printValues := strings.Split(printElements[1], ",")
				// Itera sobre los valores divididos
				for _, value := range printValues {
					// Elimina los espacios adicionales
					value = strings.TrimSpace(value)
					value = strings.Trim(value, "()")
					// Si el valor no es vacío
					if value != "" {
						// Si el valor es un número
						if _, err := strconv.Atoi(value); err == nil {
							// Imprime el valor
							fmt.Println(value)
						} else {
							// Si el valor es una variable
							if c.SymbolTable.FindFunc(value) != nil {
								// Imprime el valor de la variable
								fmt.Println(c.SymbolTable.Find(value))
							}
						}
					}
				}
			}
		} else if strings.Contains(finalStatement, ":=") {
			// Divide el statement en base a ":="
			assignElements := strings.Split(finalStatement, " := ")
			// Si el statement tiene más de 1 parte
			if len(assignElements) > 1 {
				id := c.SymbolTable.FindFunc(assignElements[0])

				if id != nil {
					idType := id.Type
					assign := assignElements[1]
					re := regexp.MustCompile(`[\+\-\*/]`)
					assigns := re.Split(assign, -1)

					for _, value := range assigns {
						// Elimina los espacios adicionales
						value = strings.TrimSpace(value)
						// Si el valor no es vacío
						varAssign := c.SymbolTable.FindFunc(value)
						varType := varAssign.Type
						if varType != idType {
							fmt.Println("Error: The variable type does not match the assigned value type")
						}
					}
				}
			}
		}
	}

	c.SymbolTable.PrintTable()
	c.SymbolTable.ExportTable()

	return c.VisitChildren(ctx)

}

func (c *Checker) VisitBlock(ctx *parser.BlockContext) interface{} {

	fmt.Println("GG")
	fmt.Println(ctx.GetText())
	return c.VisitChildren(ctx)
}

func (c *Checker) VisitStatementPrintAST(ctx *parser.StatementPrintASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitStatementPrintlnAST(ctx *parser.StatementPrintlnASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitStatementReturnAST(ctx *parser.StatementReturnASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitStatementBreakAST(ctx *parser.StatementBreakASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitStatementContinueAST(ctx *parser.StatementContinueASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitStatementSimpleAST(ctx *parser.StatementSimpleASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitStatementBlockAST(ctx *parser.StatementBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitStatementSwitchAST(ctx *parser.StatementSwitchASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitStatementIfAST(ctx *parser.StatementIfASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitStatementLoopAST(ctx *parser.StatementLoopASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitStatementTypeDeclAST(ctx *parser.StatementTypeDeclASTContext) interface{} {

	panic("implement me")
}

func (c *Checker) VisitStatementVariableDeclAST(ctx *parser.StatementVariableDeclASTContext) interface{} {
	//TODO implement me
	//JUM
	panic("implement me")
}

func (c *Checker) VisitSimpleStatementEmptyAST(ctx *parser.SimpleStatementEmptyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitSimpleStatementExpressionAST(ctx *parser.SimpleStatementExpressionASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitSimpleStatementAssignmentAST(ctx *parser.SimpleStatementAssignmentASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitSimpleStatementExpressionListAssignAST(ctx *parser.SimpleStatementExpressionListAssignASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAssignmentStatementAST(ctx *parser.AssignmentStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAssignmentStatementPlusEqualAST(ctx *parser.AssignmentStatementPlusEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAssignmentStatementMinusEqualAST(ctx *parser.AssignmentStatementMinusEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAssignmentStatementBitwiseAndEqualAST(ctx *parser.AssignmentStatementBitwiseAndEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAssignmentStatementBitwiseOrEqualAST(ctx *parser.AssignmentStatementBitwiseOrEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAssignmentStatementMultiplyEqualAST(ctx *parser.AssignmentStatementMultiplyEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAssignmentStatementBitwiseXorEqualAST(ctx *parser.AssignmentStatementBitwiseXorEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAssignmentStatementShiftLeftEqualAST(ctx *parser.AssignmentStatementShiftLeftEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAssignmentStatementShiftRightEqualAST(ctx *parser.AssignmentStatementShiftRightEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAssignmentStatementBitwiseClearEqualAST(ctx *parser.AssignmentStatementBitwiseClearEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAssignmentStatementModuloEqualAST(ctx *parser.AssignmentStatementModuloEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitAssignmentStatementDivideEqualAST(ctx *parser.AssignmentStatementDivideEqualASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitIfStatementAST(ctx *parser.IfStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitIfElseIfStatementAST(ctx *parser.IfElseIfStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitIfElseStatementAST(ctx *parser.IfElseStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitIfSimpleStatementAST(ctx *parser.IfSimpleStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitIfSimpleElseIfStatementAST(ctx *parser.IfSimpleElseIfStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitIfSimpleElseStatementAST(ctx *parser.IfSimpleElseStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitLoopBlockAST(ctx *parser.LoopBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitLoopExpressionBlockAST(ctx *parser.LoopExpressionBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitLoopSimpleStatementExpressionSimpleStatementBlockAST(ctx *parser.LoopSimpleStatementExpressionSimpleStatementBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitLoopSimpleStatementSimpleStatementBlockAST(ctx *parser.LoopSimpleStatementSimpleStatementBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitSwitchStmtSimpleStatementAST(ctx *parser.SwitchStmtSimpleStatementASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitSwitchStmtExpressionAST(ctx *parser.SwitchStmtExpressionASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitSwitchStmtSimpleStatementBlockAST(ctx *parser.SwitchStmtSimpleStatementBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitSwitchStmtBlockAST(ctx *parser.SwitchStmtBlockASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionCaseClauseListEmptyAST(ctx *parser.ExpressionCaseClauseListEmptyASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionCaseClauseListAST(ctx *parser.ExpressionCaseClauseListASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionCaseClause(ctx *parser.ExpressionCaseClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionSwitchCaseAST(ctx *parser.ExpressionSwitchCaseASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitExpressionSwitchDefaultAST(ctx *parser.ExpressionSwitchDefaultASTContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitEpsilon(ctx *parser.EpsilonContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Checker) VisitIdentifierList(ctx *parser.IdentifierListContext) interface{} {

	fmt.Println("esto que es: ", ctx.IDENTIFIER(0))

	return ctx.IDENTIFIER(0)
}

func (c *Checker) Visit(tree antlr.ParseTree) interface{} {

	return tree.Accept(c)
}

func (c *Checker) VisitVariableDeclAST(ctx *parser.VariableDeclASTContext) interface{} {

	return c.Visit(ctx.SingleVarDecl())
}
