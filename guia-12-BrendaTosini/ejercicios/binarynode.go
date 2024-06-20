package ejercicios

import (
	"fmt"
	"math"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

func NewBinaryNode(data int, left *BinaryNode, right *BinaryNode) *BinaryNode {
	return &BinaryNode{left: left, right: right, data: data}
}

func (n *BinaryNode) PrintPreOrder() {
	fmt.Println(n.data)
	if n.left != nil {
		n.left.PrintPreOrder()
	}
	if n.right != nil {
		n.right.PrintPreOrder()
	}
}

func (n *BinaryNode) PrintInOrder() {
	if n.left != nil {
		n.left.PrintInOrder()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.PrintInOrder()
	}
}

func (n *BinaryNode) PrintPostOrder() {
	if n.left != nil {
		n.left.PrintPostOrder()
	}
	if n.right != nil {
		n.right.PrintPostOrder()
	}
	fmt.Println(n.data)
}

func (n *BinaryNode) Size() int {
	size := 1
	if n.left != nil {
		size += n.left.Size()
	}
	if n.right != nil {
		size += n.right.Size()
	}
	return size
}

func (n *BinaryNode) Height() int {
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
