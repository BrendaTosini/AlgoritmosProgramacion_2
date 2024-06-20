package slicelist

import (
	"errors"
	"fmt"
)

// Implementar la lista sobre slices

type SliceList[T comparable] struct {
	lista []T
}

// O(1)
func NewSliceList[T comparable]() *SliceList[T] {
	return &SliceList[T]{lista: nil}
}

// O(1)
func (l *SliceList[T]) Append(value T) {
	l.lista = append(l.lista, value)
}

// O(1)
func (l *SliceList[T]) Prepend(value T) {
	l.lista = append([]T{value}, l.lista...) //lista de un solo valor.
}

// O(1)
func (l *SliceList[T]) InsertAt(value T, position int) {

	if position < len(l.lista) && position >= 0 {

		lista2 := l.lista[:position]
		lista3 := l.lista[position+1:]
		l.lista = append(lista2, l.lista[3])
		l.lista = append(l.lista, lista3...)
	}

}

// O(n)
func (l *SliceList[T]) Remove(value T) {

	for i := 0; i < len(l.lista); i++ {
		if l.lista[i] == value {
			lista2 := l.lista[:i]
			lista3 := l.lista[i+1:]
			l.lista = append(lista2, lista3...)
			return
		}
	}
}

// O(1)
func (l *SliceList[T]) String() string {
	return fmt.Sprintf("%v", l.lista)

}

// O(n)
func (l *SliceList[T]) Search(value T) int {

	for i := 0; i < len(l.lista); i++ {
		if l.lista[i] == value {
			return i
		}
	}
	return -1
}

// O(1)
func (l *SliceList[T]) Get(position int) (T, error) {

	if len(l.lista) == 0 {
		var t T
		return t, errors.New("la lista esta vacia")
	}

	if position < 0 || position >= len(l.lista) {
		var t T
		return t, errors.New("posicion invalida")
	} else {
		valor := l.lista[position]
		return valor, nil
	}

}

// O(1)
func (l *SliceList[T]) Size() int {
	return len(l.lista)

}
