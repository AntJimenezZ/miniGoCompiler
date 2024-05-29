package encoder

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// BlockVariableTable define una estructura para almacenar nombres de variables y valores por bloque
type BlockVariableTable struct {
	table map[*ir.Block]map[string]value.Value
}

// NewBlockVariableTable crea una nueva instancia de BlockVariableTable
func NewBlockVariableTable() *BlockVariableTable {
	return &BlockVariableTable{
		table: make(map[*ir.Block]map[string]value.Value),
	}
}

// AddVariable agrega una variable a un bloque específico
func (bvt *BlockVariableTable) AddVariable(block *ir.Block, name string, val value.Value) {
	if _, exists := bvt.table[block]; !exists {
		bvt.table[block] = make(map[string]value.Value)
	}
	bvt.table[block][name] = val
}

// GetVariable obtiene el valor de una variable en un bloque específico
func (bvt *BlockVariableTable) GetVariable(block *ir.Block, name string) (value.Value, bool) {
	if vars, exists := bvt.table[block]; exists {
		val, found := vars[name]
		return val, found
	}
	return nil, false
}

// RemoveVariable elimina una variable de un bloque específico
func (bvt *BlockVariableTable) RemoveVariable(block *ir.Block, name string) {
	if vars, exists := bvt.table[block]; exists {
		delete(vars, name)
	}
}

// ListVariables lista todas las variables en un bloque específico
func (bvt *BlockVariableTable) ListVariables(block *ir.Block) map[string]value.Value {
	if vars, exists := bvt.table[block]; exists {
		return vars
	}
	return nil
}

// ClearBlock elimina todas las variables de un bloque específico
func (bvt *BlockVariableTable) ClearBlock(block *ir.Block) {
	delete(bvt.table, block)
}
