// Parser Rules
parser grammar MiniGoParser;

options {
  tokenVocab=MiniGoScanner;
}

root            : PACKAGE IDENTIFIER SEMICOLON topDeclarationList
                            ;
topDeclarationList : ( variableDecl | typeDecl | funcDecl )*
                            ;
variableDecl    :   VAR singleVarDecl SEMICOLON                                                                         #variableDeclAST
                  | VAR LPAREN innerVarDecls RPAREN SEMICOLON                                                           #variableDeclBlockAST
                  | VAR LPAREN RPAREN SEMICOLON                                                                         #variableDeclEmptyAST
                            ;
innerVarDecls   : singleVarDecl SEMICOLON ( singleVarDecl SEMICOLON )*
                            ;
singleVarDecl   : identifierList declType ASSIGN expressionList                                                         #singleVarDeclAST
                  | identifierList ASSIGN expressionList                                                                #singleVarDeclAssignAST
                  | singleVarDeclNoExps                                                                                 #singleVarDeclNoExpsAST
                            ;
singleVarDeclNoExps : identifierList declType
                            ;
typeDecl        :   TYPE singleTypeDecl SEMICOLON                                                                       #typeDeclAST
                  | TYPE LPAREN innerTypeDecls RPAREN SEMICOLON                                                         #typeDeclBlockAST
                  | TYPE LPAREN RPAREN SEMICOLON                                                                        #typeDeclEmptyAST
                            ;
innerTypeDecls  : singleTypeDecl SEMICOLON ( singleTypeDecl SEMICOLON )*
                            ;
singleTypeDecl  : IDENTIFIER declType
                            ;
funcDecl        : funcFrontDecl block SEMICOLON
                            ;
funcFrontDecl   : FUNC IDENTIFIER LPAREN ( funcArgDecls | epsilon ) RPAREN ( multipleReturnTypes | singleReturnType )
                ;

multipleReturnTypes : LPAREN returnTypeList RPAREN
                     ;

returnTypeList  : declType ( COMMA declType )*
                ;

singleReturnType: declType                                                                                              #singleReturnTypeAST
                 | epsilon                                                                                              #singleReturnTypeEmptyAST
                 ;

funcArgDecls    : singleVarDeclNoExps ( COMMA singleVarDeclNoExps )*
                            ;
declType        : LPAREN declType RPAREN                                                                                #declTypeParenAST
                  | IDENTIFIER                                                                                          #declTypeIdentifierAST
                  | sliceDeclType                                                                                       #declTypeSliceAST
                  | arrayDeclType                                                                                       #declTypeArrayAST
                  | structDeclType                                                                                      #declTypeStructAST
                            ;
sliceDeclType   : LBRACKET RBRACKET declType
                            ;
arrayDeclType   : LBRACKET INTLITERAL RBRACKET declType
                            ;
structDeclType  : STRUCT LBRACE ( structMemDecls | epsilon ) RBRACE
                            ;
structMemDecls  : singleVarDeclNoExps SEMICOLON ( singleVarDeclNoExps SEMICOLON )*
                            ;
identifierList  : IDENTIFIER ( COMMA IDENTIFIER )*
                            ;
expression      : primaryExpression                                                                                     #expressionPrimaryAST
                  | expression MULTIPLY expression                                                                      #expressionMultiplyAST
                  | expression DIVIDE expression                                                                        #expressionDivideAST
                  | expression MODULO expression                                                                        #expressionModuloAST
                  | expression SHIFTLEFT expression                                                                     #expressionShiftLeftAST
                  | expression SHIFTRIGHT expression                                                                    #expressionShiftRightAST
                  | expression BITWISEAND expression                                                                    #expressionBitwiseAndAST
                  | expression BITWISECLEAR expression                                                                  #expressionBitwiseClearAST
                  | expression PLUS expression                                                                          #expressionPlusAST
                  | expression MINUS expression                                                                         #expressionMinusAST
                  | expression BITWISEOR expression                                                                     #expressionBitwiseOrAST
                  | expression BITWISEXOR expression                                                                    #expressionBitwiseXorAST
                  | expression EQUAL expression                                                                         #expressionEqualAST
                  | expression NOTEQUAL expression                                                                      #expressionNotEqualAST
                  | expression LESS expression                                                                          #expressionLessAST
                  | expression LESSEQUAL expression                                                                     #expressionLessEqualAST
                  | expression GREATER expression                                                                       #expressionGreaterAST
                  | expression GREATEREQUAL expression                                                                  #expressionGreaterEqualAST
                  | expression AND expression                                                                           #expressionAndAST
                  | expression OR expression                                                                            #expressionOrAST
                  | PLUS expression                                                                                     #expressionPlusUnaryAST
                  | MINUS expression                                                                                    #expressionMinusUnaryAST
                  | NOT expression                                                                                      #expressionNotUnaryAST
                  | BITWISEXOR expression                                                                               #expressionBitwiseXorUnaryAST
                           ;
expressionList  : expression ( COMMA expression )*
                            ;
primaryExpression : operand                                                                                             #primaryExpressionOperandAST
                  | primaryExpression selector                                                                          #primaryExpressionSelectorAST
                  | primaryExpression index                                                                             #primaryExpressionIndexAST
                  | primaryExpression arguments                                                                         #primaryExpressionArgumentsAST
                  | appendExpression                                                                                    #primaryExpressionAppendAST
                  | lengthExpression                                                                                    #primaryExpressionLengthAST
                  | capExpression                                                                                       #primaryExpressionCapAST
                          ;
operand         :   literal                                                                                             #operandLiteralAST
                  | IDENTIFIER                                                                                          #operandIdentifierAST
                  | LPAREN expression RPAREN                                                                            #operandParenAST
                         ;
literal         :   INTLITERAL                                                                                          #literalIntAST
                  | FLOATLITERAL                                                                                        #literalFloatAST
                  | RUNELITERAL                                                                                         #literalRuneAST
                  | RAWSTRINGLITERAL                                                                                    #literalRawStringAST
                  | INTERPRETEDSTRINGLITERAL                                                                            #literalInterpretedStringAST
                        ;
index           : LBRACKET expression RBRACKET
                        ;
arguments       : LPAREN expressionList                                                                                 #argumentsAST
                  | epsilon RPAREN                                                                                      #argumentsEmptyAST
                        ;
selector        : DOT IDENTIFIER
                        ;
appendExpression : APPEND LPAREN expression COMMA expression RPAREN
                        ;
lengthExpression: LEN LPAREN expression RPAREN
                        ;
capExpression   : CAP LPAREN expression RPAREN
                        ;
statementList   : statement*
                        ;
block           : LBRACE statementList RBRACE
                        ;
statement       : PRINT LPAREN ( expressionList | epsilon ) RPAREN SEMICOLON                                            #statementPrintAST
                  | PRINTLN LPAREN ( expressionList | epsilon ) RPAREN SEMICOLON                                        #statementPrintlnAST
                  | RETURN ( expression | epsilon ) SEMICOLON                                                           #statementReturnAST
                  | BREAK SEMICOLON                                                                                     #statementBreakAST
                  | CONTINUE SEMICOLON                                                                                  #statementContinueAST
                  | simpleStatement SEMICOLON                                                                           #statementSimpleAST
                  | block SEMICOLON                                                                                     #statementBlockAST
                  | switchStmt SEMICOLON                                                                                #statementSwitchAST
                  | ifStatement SEMICOLON                                                                               #statementIfAST
                  | loop SEMICOLON                                                                                      #statementLoopAST
                  | typeDecl                                                                                            #statementTypeDeclAST
                  | variableDecl                                                                                        #statementVariableDeclAST
                        ;
simpleStatement :   epsilon                                                                                             #simpleStatementEmptyAST
                  | expression ( INCREMENT | DECREMENT | epsilon )                                                      #simpleStatementExpressionAST
                  | assignmentStatement                                                                                 #simpleStatementAssignmentAST
                  | expressionList ASSIGN expressionList                                                                #simpleStatementExpressionListAssignAST
                        ;
assignmentStatement : expressionList ASSIGN expressionList                                                              #assignmentStatementAST
                  | expression PLUSEQUAL expression                                                                     #assignmentStatementPlusEqualAST
                  | expression MINUSEQUAL expression                                                                    #assignmentStatementMinusEqualAST
                  | expression BITWISEANDEQUAL expression                                                               #assignmentStatementBitwiseAndEqualAST
                  | expression BITWISEOREQUAL expression                                                                #assignmentStatementBitwiseOrEqualAST
                  | expression MULTIPLYEQUAL expression                                                                 #assignmentStatementMultiplyEqualAST
                  | expression BITWISEXOREQUAL expression                                                               #assignmentStatementBitwiseXorEqualAST
                  | expression SHIFTLEFTEQUAL expression                                                                #assignmentStatementShiftLeftEqualAST
                  | expression SHIFTRIGHTEQUAL expression                                                               #assignmentStatementShiftRightEqualAST
                  | expression BITWISECLEAREQUAL expression                                                             #assignmentStatementBitwiseClearEqualAST
                  | expression MODULOEQUAL expression                                                                   #assignmentStatementModuloEqualAST
                  | expression DIVIDEEQUAL expression                                                                   #assignmentStatementDivideEqualAST
                         ;
ifStatement     :   IF expression block                                                                                 #ifStatementAST
                  | IF expression block ELSE ifStatement                                                                #ifElseIfStatementAST
                  | IF expression block ELSE block                                                                      #ifElseStatementAST
                  | IF simpleStatement SEMICOLON expression block                                                       #ifSimpleStatementAST
                  | IF simpleStatement SEMICOLON expression block ELSE ifStatement                                      #ifSimpleElseIfStatementAST
                  | IF simpleStatement SEMICOLON expression block ELSE block                                            #ifSimpleElseStatementAST
                        ;
loop            :   FOR block                                                                                           #loopBlockAST
                  | FOR expression block                                                                                #loopExpressionBlockAST
                  | FOR simpleStatement SEMICOLON expression SEMICOLON simpleStatement block                            #loopSimpleStatementExpressionSimpleStatementBlockAST
                  | FOR simpleStatement SEMICOLON SEMICOLON simpleStatement block                                       #loopSimpleStatementSimpleStatementBlockAST
                        ;
switchStmt      :   SWITCH simpleStatement SEMICOLON expression LBRACE expressionCaseClauseList RBRACE                  #switchStmtSimpleStatementAST
                  | SWITCH expression LBRACE expressionCaseClauseList RBRACE                                            #switchStmtExpressionAST
                  | SWITCH simpleStatement SEMICOLON LBRACE expressionCaseClauseList RBRACE                             #switchStmtSimpleStatementBlockAST
                  | SWITCH LBRACE expressionCaseClauseList RBRACE                                                       #switchStmtBlockAST
                        ;
expressionCaseClauseList : epsilon                                                                                      #expressionCaseClauseListEmptyAST
                  | expressionCaseClauseList expressionCaseClause                                                       #expressionCaseClauseListAST
                        ;
expressionCaseClause : expressionSwitchCase COLON statementList
                        ;
expressionSwitchCase : CASE expressionList                                                                              #expressionSwitchCaseAST
                  | DEFAULT                                                                                             #expressionSwitchDefaultAST
                        ;
epsilon           :
                        ;