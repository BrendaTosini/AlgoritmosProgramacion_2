package avltree

import (
	"errors"
	"guia14/queue"

	"golang.org/x/exp/constraints"
)

type AVLLevelIterator[T constraints.Ordered] struct {
	queue *queue.Queue[AVLNode[T]] //cola de nodos
}

func NewAVLLevelIterator[T constraints.Ordered](root *AVLNode[T]) *AVLLevelIterator[T] {
	iterator := &AVLLevelIterator[T]{
		queue: queue.NewQueue[AVLNode[T]](),
	}
	iterator.queue.Enqueue(*root)
	return iterator
}

func (it *AVLLevelIterator[T]) HasNext() bool {
	return !it.queue.IsEmpty()
}

func (it *AVLLevelIterator[T]) Next() (AVLNode[T], error) {
	var dato T
	nodo := newAVLNode[T](dato, nil, nil)
	if it.queue.IsEmpty() {
		return *nodo, errors.New("No hay más elementos")
	}
	next, _ := it.queue.Dequeue()
	if next.getLeft() != nil {
		it.queue.Enqueue(*next.getLeft())
	}
	if next.getRight() != nil {
		it.queue.Enqueue(*next.getRight())
	}
	return next, nil

}
