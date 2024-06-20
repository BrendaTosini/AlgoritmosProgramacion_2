package avltree

import (
	"errors"
	"guia14/stack"

	"golang.org/x/exp/constraints"
)

type AVLInOrderIterator[T constraints.Ordered] struct {
	stack *stack.Stack[AVLNode[T]] //pila de nodos
}

func NewAVLInOrderIterator[T constraints.Ordered](root *AVLNode[T]) *AVLInOrderIterator[T] {
	iterator := &AVLInOrderIterator[T]{
		stack: stack.NewStack[AVLNode[T]](),
	}
	iterator.apilarHijosIzquierdos(root)
	return iterator
}

func (it *AVLInOrderIterator[T]) apilarHijosIzquierdos(nodo *AVLNode[T]) {
	for nodo != nil {
		it.stack.Push(*nodo)
		nodo = nodo.getLeft()
	}
}

func (it *AVLInOrderIterator[T]) HasNext() bool {
	return !it.stack.IsEmpty()
}

func (it *AVLInOrderIterator[T]) Next() (AVLNode[T], error) {
	var dato T
	nodo := newAVLNode[T](dato, nil, nil)
	if it.stack.IsEmpty() {
		return *nodo, errors.New("No hay más elementos")
	}
	next := it.stack.Pop()
	if next.getRight() != nil {
		it.apilarHijosIzquierdos(next.getRight())
	}
	return next, nil

}
