package main

import (
	"fmt"

	"guia12/binarytree"
)

func main() {

	btree1 := binarytree.NewBinaryTree("+")
	btree1.InsertLeft(binarytree.NewBinaryTree("A"))
	btree1.InsertRight(binarytree.NewBinaryTree("B"))

	btree2 := binarytree.NewBinaryTree("*")
	btree2.InsertLeft(binarytree.NewBinaryTree("C"))
	btree2.InsertRight(binarytree.NewBinaryTree("D"))

	btree3 := binarytree.NewBinaryTree("-")
	btree3.InsertLeft(btree1)
	btree3.InsertRight(btree2)

	fmt.Println("-----------PrintInOrder-----------")
	btree3.PrintInOrder()
	fmt.Println("-----------PrintPreOrder----------")
	btree3.PrintPreOrder()
	fmt.Println("-----------PrintPostOrder----------")
	btree3.PrintPostOrder()
}
