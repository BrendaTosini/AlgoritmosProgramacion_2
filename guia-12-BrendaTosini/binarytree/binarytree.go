package binarytree

type BinaryTree[T any] struct {
	root *BinaryNode[T]
}

func NewBinaryTree[T any](data T) *BinaryTree[T] {
	node := NewBinaryNode(data, nil, nil)
	return &BinaryTree[T]{root: node}

}

func (t *BinaryTree[T]) InsertLeft(bt *BinaryTree[T]) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.left = bt.root
	}
}
func (t *BinaryTree[T]) InsertRight(bt *BinaryTree[T]) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.right = bt.root
	}
}

func (t *BinaryTree[T]) PrintPreOrder() {
	if t.root != nil {
		t.root.PrintPreOrder()
	}
}

func (t *BinaryTree[T]) PrintInOrder() {
	if t.root != nil {
		t.root.PrintInOrder()
	}
}

func (t *BinaryTree[T]) PrintPostOrder() {
	if t.root != nil {
		t.root.PrintPostOrder()
	}
}

func (t *BinaryTree[T]) Empty() {
	t.root = nil
}

func (t *BinaryTree[T]) IsEmpty() bool {
	return t.root == nil
}

func (t *BinaryTree[T]) Size() int {
	if t.root != nil {
		return t.root.Size()
	} else {
		return -1
	}
}

func (t *BinaryTree[T]) Height() int {
	if t.root != nil {
		return t.root.Height()
	} else {
		return 0
	}
}
