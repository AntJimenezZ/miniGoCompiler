package encoder

import (
	"errors"
	"github.com/llir/llvm/ir"
)

// Stack define una estructura de pila
type ifStack struct {
	items []*ir.Block
}

// Push agrega un elemento al tope de la pila
func (s *ifStack) Push(item *ir.Block) {
	s.items = append(s.items, item)
}

// Pop elimina y devuelve el elemento al tope de la pila
func (s *ifStack) Pop() (*ir.Block, error) {
	if len(s.items) == 0 {
		return nil, errors.New("la pila está vacía")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek devuelve el elemento al tope de la pila sin eliminarlo
func (s *ifStack) Peek() (*ir.Block, error) {
	if len(s.items) == 0 {
		return nil, errors.New("la pila está vacía")
	}
	return s.items[len(s.items)-1], nil
}

// GetFromTop devuelve el elemento en el índice especificado desde el tope de la pila
func (s *ifStack) GetFromTop(index int) (*ir.Block, error) {
	if index < 0 || index >= len(s.items) {
		return nil, errors.New("índice fuera de rango")
	}
	return s.items[len(s.items)-1-index], nil
}

// IsEmpty verifica si la pila está vacía
func (s *ifStack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size devuelve el número de elementos en la pila
func (s *ifStack) Size() int {
	return len(s.items)
}

// Empty elimina todos los elementos de la pila
func (s *ifStack) Empty() {
	s.items = nil
}