package stack

import "errors"

// Stack implementa una pila genérica sobre un arreglo dinámico.

type Stack struct {
	pila []any
}

// Push agrega un elemento a la pila. O(1)
func (s *Stack) Push(v any) {
	(*s).pila = append((*s).pila, v)
}

// Pop saca un elemento de la pila. O(1)
func (s *Stack) Pop() (any, error) {
	if (*s).IsEmpty() {
		return nil, errors.New("la pila esta vacia")
	}
	v := (*s).pila[len((*s).pila)-1] // s[1] se ponen los parentesis porque se tiene que desreferenciar a s y no a s.pila. Pila no es un puntero.
	(*s).pila = (*s).pila[:len((*s).pila)-1]
	return v, nil
}

// Top devuelve el elemento del tope de la pila. O(1)
func (s *Stack) Top() (any, error) {
	if (*s).IsEmpty() {
		return 0, errors.New("la pila esta vacia")
	}
	v := (*s).pila[len((*s).pila)-1]
	return v, nil
}

// IsEmpty verifica si la pila esta vacia. O(1)
func (s *Stack) IsEmpty() bool {
	return len((*s).pila) == 0
}
