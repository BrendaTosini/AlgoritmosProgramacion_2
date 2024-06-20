package stack

import (
	"fmt"
)

// TAD Pila sobre slices
type Stack[T any] struct {
	datos []T
}

// Constructor. Devuelve una pila vacía
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{datos: []T{}}
}

// *********************************Métodos del TAD Stack*********************************
// Devuelve un string que representa la pila
func (s Stack[T]) String() string {
	return fmt.Sprintf("%v", s.datos)
}

// Agrega un elemento a la pila
func (s *Stack[T]) Push(dato T) {
	s.datos = append(s.datos, dato)
}

// Elimina el último elemento de la pila
func (s *Stack[T]) Pop() T {
	dato := s.datos[len(s.datos)-1]
	s.datos = s.datos[:len(s.datos)-1]
	return dato
}

// Devuelve el último elemento de la pila
func (s Stack[T]) Top() T {
	return s.datos[len(s.datos)-1]
}

// Devuelve true si la pila está vacía
func (s Stack[T]) IsEmpty() bool {
	return len(s.datos) == 0
}
