package main

import (
	"Proyecto_Compiladores/checker"
	"Proyecto_Compiladores/encoder"
	"Proyecto_Compiladores/parser"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type MyErrorListener struct {
	antlr.DefaultErrorListener
}

func NewMyErrorListener() *MyErrorListener {
	return new(MyErrorListener)
}

var errors []string

func (d *MyErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	err := fmt.Sprintf("Error in line %d:%d %s\n", line, column, msg)
	errors = append(errors, err)
	// Abre el archivo en modo append
	f, errFile := os.OpenFile("txtFiles/errors.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errFile != nil {
		fmt.Println(errFile)
		return
	}
	defer f.Close()

	// Escribe el error en el archivo
	if _, errFile := f.WriteString(err); errFile != nil {
		fmt.Println(errFile)
		return
	}
}

type Message struct {
	Text string `json:"text"`
}

func handler(w http.ResponseWriter, r *http.Request) { // In this function are all the logic of the server
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		// Preflight request
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return
	}
	// Decode the JSON
	var clientMsg Message
	err := json.NewDecoder(r.Body).Decode(&clientMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Save the message on test.txt
	file, err := os.Create("txtFiles/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file2, err := os.OpenFile("txtFiles/errors.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file2.Close()

	defer file.Close()
	_, _ = file.WriteString(clientMsg.Text)
	//Call the function where is all the compiler logic
	compiler()
	formatError := ""

	if len(errors) == 0 {
		formatError = "No errors.txt found!!! :)"
	} else {
		for i, err := range errors {
			formatError += fmt.Sprintf("\n%d) %s", i+1, err)
		}
	}

	consoleTests(clientMsg) //Call the function where is all the console prints

	// Send the response to the client
	response := Message{Text: formatError}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
	formatError = ""
	errors = []string{}
}

func compiler() *ir.Module {

	//Send the file to parser
	input, _ := antlr.NewFileStream("txtFiles/test.txt")
	lexer := parser.NewMiniGoScanner(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewMiniGoParser(stream)
	p.RemoveErrorListeners()         // Removes default listener
	listener := NewMyErrorListener() // This creates your own listener
	lexer.RemoveErrorListeners()     // Removes default listener
	p.AddErrorListener(listener)     // This creates your own listener
	lexer.AddErrorListener(listener)

	tree := p.Root()

	encoder := encoder.NewEncoderLLVM()
	module := encoder.Visit(tree).(*ir.Module)
	fmt.Println(module)

	//fmt.Println("Compilation failed")

	globalSymbolTable := checker.NewSymbolTable()
	check := &checker.Checker{
		SymbolTable: globalSymbolTable,
	}

	check.Visit(tree)
	return module
}

func consoleTests(clientMsg Message) {

	//DEPRECATED
	fmt.Printf("Message from client:\n%s\n", clientMsg.Text)
	fmt.Println("----------------------")
	fmt.Printf("Errors list:\n %v\n", errors)
}

func runModule(module *ir.Module) {
	// Escribir el módulo LLVM en un archivo
	f, err := os.Create("modules/module.ll")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer f.Close()
	if _, err := module.WriteTo(f); err != nil {
		fmt.Println("Error al escribir el módulo:", err)
		return
	}

	// Compilar el módulo LLVM a código objeto
	cmd := exec.Command("clang", "", "modules/module.ll", "-o", "modules/module.exe")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error al compilar el módulo:", err)
		return
	}

	fmt.Println("El archivo ejecutable .exe ha sido generado correctamente.")

	if _, err := os.Stat("modules/module.exe"); os.IsNotExist(err) {
		fmt.Println("Error: modules/module.exe no existe")
		return
	}

	//ejecute el programa
	cmd = exec.Command("./modules/module.exe")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		fmt.Println("Error al ejecutar el comando:", err)
		return
	}

}

func main() {

	module := compiler()
	if module != nil {
		runModule(module)
	} else {
		fmt.Println("No module generated")
	}
	/*
		http.HandleFunc("/json", handler)

		fmt.Println("Server is running on localhost")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("Error starting server: %s\n", err)
		}
	*/
}
