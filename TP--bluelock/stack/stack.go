package stack

import (
	"bluelock/linkedlist"
	"errors"
)

type Stack[T comparable] struct {
	list linkedlist.LinkedList[T]
}

func (p *Stack[T]) Push(value T) {
	p.list.Append(value)
}

func (p *Stack[T]) Pop() (T, error) {
	if p.IsEmpty() {
		var t T
		return t, errors.New("la pila esta vacia")
	}
	top, _ := p.Top()
	p.list.Remove(top)
	return top, nil
}

func (p *Stack[T]) Top() (T, error) {
	if p.IsEmpty() {
		var t T
		return t, errors.New("la pila esta vacia")
	}
	top_position := p.list.Size() - 1
	top := p.list.Get(top_position)
	return top, nil
}

func (p *Stack[T]) IsEmpty() bool {
	return p.list.Size() == 0
}
