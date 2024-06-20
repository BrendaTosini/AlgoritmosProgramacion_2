package linkedlist

import (
	"fmt"
)

// node es el nodo de la lista enlazada
// contiene un valor y un puntero al siguiente nodo
// el valor es de tipo genérico, comparable
type node[T any] struct {
	value T
	next  *node[T]
}

// newNode crea un nuevo nodo, con el valor recibido
// y el puntero al siguiente nodo en nil
func newNode[T comparable](value T) *node[T] {
	return &node[T]{value: value, next: nil}
}

/************************************************************/

// LinkedList es la lista enlazada simple
// contiene punteros al primer nodo y al último
type LinkedList[T comparable] struct {
	head *node[T] // puntero al primer nodo
	tail *node[T] // puntero al último nodo
	size int
}

// NewLinkedList crea una nueva lista enlazada, vacía
// En nuestra implementación representamos la lista vacía
// con un puntero al primer nodo en nil
// y un puntero al último nodo en nil
// O(1)
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, tail: nil, size: 0}
}

// Append agrega un nuevo nodo, con el valor recibido, al final de la lista
// O(1)
func (l *LinkedList[T]) Append(value T) {
	l.size++
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.next = newNode
	l.tail = newNode
}

// Prepend agrega un nuevo nodo, con el valor recibido,
// al inicio de la lista
// O(1)
func (l *LinkedList[T]) Prepend(value T) {
	l.size++
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
}

// InsertAt agrega un nuevo nodo, con el valor recibido,
// en la posición recibida.
// Si la posición es inválida, no hace nada
// O(n)
func (l *LinkedList[T]) InsertAt(value T, position int) {
	l.size++
	if position < 0 {
		return
	}
	newNode := newNode(value)
	if position == 0 {
		l.Prepend(value)
		return
	}
	current := l.head
	for current != nil && position > 1 {
		current = current.next
		position--
	}
	if current == nil {
		return
	}
	newNode.next = current.next
	current.next = newNode
}

// Remove elimina el primer nodo que contenga el valor recibido
// O(n)
func (l *LinkedList[T]) Remove(value T) {
	if l.head == nil {
		return // no hay nada que eliminar
	}
	l.size--
	if l.head.value == value {
		l.head = l.head.next
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (l *LinkedList[T]) RemoveAt(position int) {
	if position < 0 || position >= l.size {
		return
	}
	// Eliminar el primer nodo
	// O(1)
	if position == 0 {
		l.head = l.head.next
		l.size--
		return
	}

	current := l.head
	for current != nil && position > 1 {
		current = current.next
		position--
	}

	current.next = current.next.next
	l.size--
}

// String devuelve una representación en cadena de la lista
// en el formato [1 2 3].
// Se puede usar para imprimir la lista con fmt.Println
// O(n)
func (l *LinkedList[T]) String() string {
	if l.head == nil {
		return "[]"
	}
	current := l.head
	result := "["
	for current != nil {
		result += fmt.Sprintf("%v", current.value)
		if current.next != nil {
			result += " "
		}
		current = current.next
	}
	result += "]"
	return result
}

// Contains verifica si el valor recibido está en la lista
// O(n)
func (l *LinkedList[T]) Contains(value T) bool {
	return l.Search(value) != -1
}

// Search busca el primer nodo que contenga el valor recibido
// y devuelve su posición en la lista o -1 si no lo encuentra
// O(n)
func (l *LinkedList[T]) Search(value T) int {
	if l.head == nil {
		return -1
	}
	current := l.head
	position := 0
	for current != nil {
		if current.value == value {
			return position
		}
		current = current.next
		position++
	}
	return -1
}

// Get devuelve el valor del nodo en la posición recibida
// o un valor nulo si la posición es inválida
// O(n)
func (l *LinkedList[T]) Get(position int) T {
	if position < 0 || position >= l.size {
		var t T // la variable t se inicializ con un valor nulo
		return t
	}

	current := l.head
	for current != nil && position > 0 {
		current = current.next
		position--
	}

	return current.value
}

// Get devuelve el nodo en la posición recibida
// o un valor nulo si la posición es inválida
// O(n)
func (l *LinkedList[T]) GetNode(position int) *node[T] {
	if position < 0 || position >= l.size {
		return nil
	}

	current := l.head
	for current != nil && position > 0 {
		current = current.next
		position--
	}
	return current
}

// UpdateNodeValue actualiza al value del nodo
// O(n)
func (l *LinkedList[T]) UpdateNodeValue(targetValue, newValue T) {
	current := l.head

	for current != nil {
		if current.value == targetValue {
			current.value = newValue
			return
		}
		current = current.next
	}
}

// Value devuelve al value del nodo
// O(1)
func (n *node[T]) Value() T {
	if n == nil {
		var t T // la variable t se inicializ con un valor nulo
		return t
	}
	return n.value
}

// Size devuelve la cantidad de nodos en la lista
// O(1)
func (l *LinkedList[T]) Size() int {
	if l.head == nil {
		return 0
	}
	return l.size
}

// Devuelve un Slice de tareas
func (l *LinkedList[T]) ToSlice() []T {
	var slice []T

	if l.head == nil {
		return slice
	}

	current := l.head

	for current != nil {
		//Append de cada tarea
		slice = append(slice, current.value)
		current = current.next
	}

	return slice
}
