package avltree

import (
	"golang.org/x/exp/constraints"
)

// TAD Árbol AVL
type AVLTree[T constraints.Ordered] struct {
	root *AVLNode[T]
}

// Constructor. Devuelve un árbol AVL con la raíz inicializada en nil
func NewAVLTree[T constraints.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{root: nil}
}

// *********************************Métodos del TAD AVLTree*********************************
// Devuelve un string que representa el árbol AVL
func (avl AVLTree[T]) String() string {
	return avl.root.string()
}

// Devuelve la raíz del árbol
func (avl AVLTree[T]) GetRoot() *AVLNode[T] {
	return avl.root
}

// Insertar en un árbol AVL de manera de siempre conservar el equilibrio
func (avl *AVLTree[T]) Insert(k T) {
	avl.root = avl.root.insert(k)
}

// Eliminar un elemento en un árbol AVL
func (avl *AVLTree[T]) Remove(k T) {
	avl.root = avl.root.remove(k)
}

// Buscar un elemento en un árbol AVL
func (avl AVLTree[T]) Search(k T) bool {
	return avl.root.search(k)
}

// Devuelve el menor elemento del árbol
func (avl AVLTree[T]) FindMin() T {
	min := avl.root.findMin()
	return min.GetData()
}

// Devuelve el mayor elemento del árbol
func (avl AVLTree[T]) FindMax() T {
	max := avl.root.findMax()
	return max.GetData()
}

// Devuelve la altura del árbol
func (avl AVLTree[T]) GetHeight() int {
	return avl.root.getHeight()
}

// Devuelve el balance del árbol
func (avl AVLTree[T]) GetBalance() int {
	return avl.root.getBalance()
}

// Devuelve true si el árbol está vacío
func (avl AVLTree[T]) isEmpty() bool {
	return avl.root == nil
}

// Elimina todos los elementos del árbol
func (avl *AVLTree[T]) Clear() {
	avl.root = nil
}

func (avl *AVLTree[T]) InOrder() string {
	return avl.root.inOrder()

}
