package binarytree

import (
	"fmt"
	"math"
)

type BinaryNode[T any] struct {
	left  *BinaryNode[T]
	right *BinaryNode[T]
	data  T
}

func NewBinaryNode[T any](data T, left *BinaryNode[T], right *BinaryNode[T]) *BinaryNode[T] {
	return &BinaryNode[T]{left: left, right: right, data: data}
}

func (n *BinaryNode[T]) PrintPreOrder() {
	fmt.Println(n.data)
	if n.left != nil {
		n.left.PrintPreOrder()
	}
	if n.right != nil {
		n.right.PrintPreOrder()
	}
}

func (n *BinaryNode[T]) PrintInOrder() {
	if n.left != nil {
		n.left.PrintInOrder()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.PrintInOrder()
	}
}

func (n *BinaryNode[T]) PrintPostOrder() {
	if n.left != nil {
		n.left.PrintPostOrder()
	}
	if n.right != nil {
		n.right.PrintPostOrder()
	}
	fmt.Println(n.data)
}

func (n *BinaryNode[T]) Size() int {
	size := 1
	if n.left != nil {
		size += n.left.Size()
	}
	if n.right != nil {
		size += n.right.Size()
	}
	return size
}

func (n *BinaryNode[T]) Height() int {
	leftHeight := -1
	rightHeight := -1
	if n.left != nil {
		leftHeight = n.left.Height()
	}
	if n.right != nil {
		rightHeight = n.right.Height()
	}
	return int(1 + math.Max(float64(leftHeight), float64(rightHeight)))
}
