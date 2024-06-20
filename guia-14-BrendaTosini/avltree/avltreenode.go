package avltree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Nodo del árbol AVL, además del dato y los hijos, registra la altura
type AVLNode[T constraints.Ordered] struct {
	data   T           // dato
	height int         // altura
	left   *AVLNode[T] // hijo izquierdo
	right  *AVLNode[T] // hijo derecho
}

// Constructor. Devuelve un nodo con la altura inicializada en 1
func newAVLNode[T constraints.Ordered](data T, left *AVLNode[T], right *AVLNode[T]) *AVLNode[T] {
	return &AVLNode[T]{left: left, right: right, data: data, height: 1}
}

// *********************************Métodos del TAD AVLNode*********************************
// Devuelve el dato del nodo
func (n *AVLNode[T]) GetData() T {
	return n.data
}

// Devuelve un string que representa el dato almacenado en el nodo
func (n *AVLNode[T]) string() string {
	return fmt.Sprintf("%v", n.data)
}

// Devuelve el hijo izquierdo del nodo
func (n *AVLNode[T]) getLeft() *AVLNode[T] {
	return n.left
}

// Devuelve el hijo derecho del nodo
func (n *AVLNode[T]) getRight() *AVLNode[T] {
	return n.right
}

// Devuelve la altura del nodo
func (n *AVLNode[T]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

// Devuelve el balance del nodo
// El balance es la diferencia entre la altura del hijo izquierdo y la del hijo derecho
// -1 <= balance <= 1
func (n *AVLNode[T]) getBalance() int {
	if n == nil {
		return 0
	}
	return n.left.getHeight() - n.right.getHeight()
}

// Actualiza la altura del nodo
func (n *AVLNode[T]) updateHeight() {
	if n == nil {
		return
	}
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
}

// Inserta un elemento en el árbol AVL
// Si el elemento ya existe, no hace nada
// Si el elemento no existe, lo inserta y reestructura el árbol si es necesario
func (n *AVLNode[T]) insert(value T) *AVLNode[T] {
	// Si el nodo es nil, lo crea
	if n == nil {
		return newAVLNode[T](value, nil, nil)
	}
	//Primero inserta el nodo como si fuera un BST común
	if value < n.data {
		n.left = n.left.insert(value)
	} else if value > n.data {
		n.right = n.right.insert(value)
	} else { // el elemento ya se encuentra en el árbol
		return n
	}
	// Actualiza la altura del nodo
	n.updateHeight()
	// Calcula el balance del nodo
	balance := n.getBalance() // Si el balance es mayor a 1, el árbol está desbalanceado

	// Desbalanceado a la izquierda izquierda  (rotación simple a derecha)
	if balance > 1 && value < n.left.GetData() {
		return n.rotateRight()
	}

	// Desbalanceado a la derecha derecha (rotación simple a izquierda)
	if balance < -1 && value > n.right.GetData() {
		return n.rotateLeft()
	}

	// Desbalanceado a la izquierda derecha (rotación doble a derecha)
	// La rotación doble se implementa primero rotando el hijo izquierdo a la izquierda
	// y luego rotando el nodo actual a la derecha
	if balance > 1 && value > n.left.GetData() {
		n.left = n.left.rotateLeft()
		return n.rotateRight()
	}

	// Desbalanceado a la derecha izquierda (rotación doble a izquierda)
	// La rotación doble se implementa primero rotando el hijo derecho a la derecha
	// y luego rotando el nodo actual a la izquierda
	if balance < -1 && value < n.right.GetData() {
		n.right = n.right.rotateRight()
		return n.rotateLeft()
	}

	return n
}

// Rotación simple a la derecha
func (n *AVLNode[T]) rotateRight() *AVLNode[T] {
	y := n.left   // y es el hijo izquierdo de n
	t2 := y.right // t2 es el hijo derecho de y

	// reasignamos los punteros
	y.right = n
	n.left = t2

	// Actualizamos las alturas
	n.updateHeight()
	y.updateHeight()

	return y
}

// Rotación simple a la izquierda
func (n *AVLNode[T]) rotateLeft() *AVLNode[T] {
	x := n.right // x es el hijo derecho de n
	t2 := x.left // t2 es el hijo izquierdo de x

	// reasignamos los punteros
	x.left = n
	n.right = t2

	// Actualizamos las alturas
	n.updateHeight()
	x.updateHeight()

	return x
}

// Elimina un elemento del árbol AVL
// Si el elemento no existe, no hace nada
// Si el elemento existe, lo elimina y reestructura el árbol si es necesario
func (n *AVLNode[T]) remove(value T) *AVLNode[T] {
	if n == nil {
		return n
	}

	// Primero elimina el nodo como si fuera un BST común
	if value < n.GetData() {
		n.left = n.left.remove(value)
	} else if value > n.data {
		n.right = n.right.remove(value)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		temp := n.right.findMin()
		n.data = temp.data
		n.right = n.right.remove(temp.data)
	}

	// Actualiza la altura del nodo
	n.updateHeight()
	// Calcula el balance del nodo
	balance := n.getBalance()

	// Si está desbalanceado, reestructura el árbol
	// Desblanceado izquierda izquierda (rotación simple a derecha)
	if balance > 1 && n.left.getBalance() >= 0 {
		return n.rotateRight()
	}

	// Desbalanceado izquierda derecha (rotación doble a derecha)
	if balance > 1 && n.left.getBalance() < 0 {
		n.left = n.left.rotateLeft()
		return n.rotateRight()
	}

	// Desbalanceado derecha derecha (rotación simple a izquierda)
	if balance < -1 && n.right.getBalance() <= 0 {
		return n.rotateLeft()
	}

	// Desbalanceado derecha izquierda (rotación doble a izquierda)
	if balance < -1 && n.right.getBalance() > 0 {
		n.right = n.right.rotateRight()
		return n.rotateLeft()
	}

	return n
}

// Encuentra el elemento mínimo del árbol
func (n *AVLNode[T]) findMin() *AVLNode[T] {
	if n.left == nil {
		return n
	}
	return n.left.findMin()
}

// Encuentra el elemento máximo del árbol
func (n *AVLNode[T]) findMax() *AVLNode[T] {
	if n.right == nil {
		return n
	}
	return n.right.findMax()
}

// Busca un elemento en el árbol
// Devuelve true si el elemento existe, false en caso contrario
func (n *AVLNode[T]) search(k T) bool {
	if n == nil {
		return false
	}

	if n.data == k {
		return true
	}

	if n.data < k {
		return n.right.search(k)
	}

	return n.left.search(k)
}

// Recorre el árbol en inorden
// Devuelve un string con los elementos del árbol
func (n *AVLNode[T]) inOrder() string {
	if n == nil {
		return ""
	}
	return n.left.inOrder() + " " + n.string() + " " + n.right.inOrder()
}

//*****************************************************************************************************

// función auxiliar para calcular el máximo entre dos enteros
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
