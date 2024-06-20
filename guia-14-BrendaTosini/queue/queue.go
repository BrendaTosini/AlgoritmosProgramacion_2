package queue

import (
	"errors"
)

type Queue[T any] struct {
	datos []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{datos: make([]T, 0)}
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.datos) == 0
}

func (q *Queue[T]) Enqueue(dato T) {
	q.datos = append(q.datos, dato)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var dato T
	if q.IsEmpty() {
		return dato, errors.New("No hay más elementos")
	}
	dato = q.datos[0]
	q.datos = q.datos[1:]
	return dato, nil
}
