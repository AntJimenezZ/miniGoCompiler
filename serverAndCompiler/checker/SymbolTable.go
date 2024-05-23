package checker

import (
	"fmt"
	"os"
)

type Ident struct {
	token      string
	Type       int
	level      int
	objectType string
}

type VarIdent struct {
	Ident
}

type MethodIdent struct {
	Ident
	params []int
}

type TypeIdent struct {
	Ident
	baseType int
}

type StructIdent struct {
	Ident
	fields []string
}

type ObjectIdent struct {
	Ident
	_type string
}

type SymbolTable struct {
	table       []Ident
	actualLevel int
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		table:       make([]Ident, 0),
		actualLevel: 0,
	}

}

var (
	Bool = Ident{
		token: "bool",
		Type:  1,
		level: 0,
	}
	Int = Ident{
		token: "int",
		Type:  2,
		level: 0,
	}
	Float = Ident{
		token: "float",
		Type:  3,
		level: 0,
	}
	String = Ident{
		token: "string",
		Type:  4,
		level: 0,
	}
	Rune = Ident{
		token: "rune",
		Type:  5,
		level: 0,
	}
	Func = Ident{
		token: "func",
		Type:  6,
		level: 0,
	}
	Struct = Ident{
		token: "struct",
		Type:  7,
		level: 0,
	}
	Object = Ident{
		token: "Object",
		Type:  8,
		level: 0,
	}
)

func (t *SymbolTable) WriteErrorToFile(errMsg string) {
	// Abre el archivo en modo append
	f, err := os.OpenFile("errors.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Escribe el error en el archivo
	if _, err := f.WriteString(errMsg + "\n"); err != nil {
		fmt.Println(err)
		return
	}
}

func (t *SymbolTable) InsertMethod(token string, typ int, params []int) {
	i := MethodIdent{Ident: Ident{token: token, Type: typ, level: t.actualLevel}, params: params}
	t.table = append(t.table, i.Ident)
}

func (t *SymbolTable) InsertVar(token string, typ int) {
	i := VarIdent{Ident: Ident{token: token, Type: typ, level: t.actualLevel}}
	t.Find(token)
	if t.Find(token) == true {
		t.table = append(t.table, i.Ident)
	}
}

func (t *SymbolTable) InsertType(tok string, typ int, baseType int) {
	i := TypeIdent{Ident: Ident{token: tok, Type: typ, level: t.actualLevel}, baseType: baseType}
	t.Find(tok)
	t.table = append(t.table, i.Ident)
}

func (t *SymbolTable) InsertStruct(tok string, typ int, fields []string) {
	i := StructIdent{Ident: Ident{token: tok, Type: typ, level: t.actualLevel}, fields: fields}
	t.Find(tok)
	t.table = append(t.table, i.Ident)
}
func (t *SymbolTable) InsertObject(token string, _type string) {
	objectExist := false
	for _, id := range t.table {
		fmt.Println("comparacion1:", id.token)
		fmt.Println("comparacion2", token)
		if id.Type == 7 && id.token == _type {
			fmt.Println("el test definito", id.token)
			if t.Find(token) == true {
				i := ObjectIdent{Ident: Ident{token: token, Type: 8, level: t.actualLevel, objectType: _type}, _type: _type}

				t.table = append(t.table, i.Ident)
				objectExist = true

			}
		}
	}
	if objectExist == false {
		fmt.Println("ERROR, STRUCT DONT DEFINED \nOBJECT: '", _type, "' its not declared ") //TODO PASAR AL SERVER

		errorMsg := fmt.Sprintf("ERROR, STRUCT DONT DEFINED \nOBJECT: '%s' its not declared ", _type)
		t.WriteErrorToFile(errorMsg)
	}
}

func (t *SymbolTable) Find(name string) bool {
	for _, id := range t.table {
		if id.token == name && id.level == t.actualLevel{
			fmt.Printf("ERROR, MULTIPLE VAR DECLARATION \nVariable: '%s' its declared multiple times\n", id.token)
			errorMsg := fmt.Sprintf("ERROR, MULTIPLE VAR DECLARATION \nVariable: '%s' its declared multiple times", id.token)
			t.WriteErrorToFile(errorMsg)
			return false
		}
	}
	return true
}

func (t *SymbolTable) FindFunc(name string) *Ident {
	for _, id := range t.table {
		if id.token == name {
			fmt.Println("ERROR, MULTIPLE VAR DECLARATION \nVariable: '", id.token, "' its declared multiple times") //TODO PASAR AL SERVER
			errorMsg := fmt.Sprintf("ERROR, MULTIPLE VAR DECLARATION \nVariable: '%s' its declared multiple times", id.token)
			t.WriteErrorToFile(errorMsg)
			return &id
		}
	}
	return nil
}

func (t *SymbolTable) FindActualLevel(name string) *Ident {
	var temp *Ident
	tempLevel := t.actualLevel
	for _, id := range t.table {
		if tempLevel == id.level {
			if id.token == name {
				temp = &id
			}
		} else {
			break
		}
	}
	return temp
}

func (t *SymbolTable) OpenScope() {
	t.actualLevel++
}

func (t *SymbolTable) CloseScope() {
	for i := len(t.table) - 1; i >= 0; i-- {
		if t.table[i].level == t.actualLevel {
			t.table = append(t.table[:i], t.table[i+1:]...)
		}
	}
	t.actualLevel--
}

func (t *SymbolTable) PrintTable() {
	_type := ""
	fmt.Println("----- SYMBOL TABLE ------")
	fmt.Println("|    Name    |  Level  |  Type  |")
	fmt.Println("|------------|---------|--------|")
	for _, s := range t.table {
		if s.Type == 1 {
			_type = "boolean"
		} else if s.Type == 2 {
			_type = "int"
		} else if s.Type == 3 {
			_type = "float"
		} else if s.Type == 4 {
			_type = "string"
		} else if s.Type == 5 {
			_type = "rune"
		} else if s.Type == 6 {
			_type = "func"
		} else if s.Type == 7 {
			_type = "struct"
		} else if s.Type == 8 {
			_type = s.objectType

		} else {
			_type = "unknown"
		}

		fmt.Printf("|  %-10s|    %-4d|   %-4s|\n", s.token, s.level, _type)
	}
	fmt.Println("----- END TABLE ------")
}
func (t *SymbolTable) ExportTable() {
    _type := ""

    // Abre el archivo en modo de escritura y trunca cualquier contenido existente
    f, err := os.OpenFile("symbolTableView.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    fmt.Fprintln(f, "----- SYMBOL TABLE ------")
    fmt.Fprintln(f, "|    Name    |  Level  |  Type  |")
    fmt.Fprintln(f, "|------------|---------|--------|")
    for _, s := range t.table {
        if s.Type == 1 {
            _type = "boolean"
        } else if s.Type == 2 {
            _type = "int"
        } else if s.Type == 3 {
            _type = "float"
        } else if s.Type == 4 {
            _type = "string"
        } else if s.Type == 5 {
            _type = "rune"
        } else if s.Type == 6 {
            _type = "func"
        } else if s.Type == 7 {
            _type = "struct"
        } else if s.Type == 8 {
            _type = s.objectType
        } else {
            _type = "unknown"
        }

        fmt.Fprintf(f, "|  %-10s|    %-4d|   %-4s|\n", s.token, s.level, _type)
    }
    fmt.Fprintln(f, "----- END TABLE ------")
}