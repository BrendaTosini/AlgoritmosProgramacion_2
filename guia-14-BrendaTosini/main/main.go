package main

import (
	"fmt"
	"guia14/avltree"
)

func main() {

	fmt.Println("AVLTree")
	avl := avltree.NewAVLTree[int]()
	avl.Insert(7)
	avl.Insert(5)
	avl.Insert(9)
	avl.Insert(2)
	avl.Insert(6)
	avl.Insert(8)
	avl.Insert(12)
	avl.Insert(3)
	avl.Insert(10)
	avl.Insert(10)
	avl.Insert(10)
	avl.Insert(10)
	fmt.Println("Raiz: ", avl.GetRoot().GetData())
	fmt.Println("Altura: ", avl.GetHeight())
	fmt.Println("Factor de balanceo: ", avl.GetBalance())
	fmt.Println("Max", avl.FindMax())
	fmt.Println("Min", avl.FindMin())
	fmt.Println("Recorrido inorden: ")
	it := avltree.NewAVLInOrderIterator[int](avl.GetRoot())
	for it.HasNext() {
		nodo, _ := it.Next()
		fmt.Printf("%v ", nodo.GetData())
	}
	fmt.Println()
	fmt.Println("Recorrido por niveles: ")
	itl := avltree.NewAVLLevelIterator[int](avl.GetRoot())
	anterior := avl.GetRoot().GetData()
	for itl.HasNext() {
		nodo, _ := itl.Next()
		if anterior > nodo.GetData() {
			fmt.Println()
		}
		anterior = nodo.GetData()
		fmt.Printf("%v ", nodo.GetData())
	}
	fmt.Println()

	fmt.Println("Buscar 3: ", avl.Search(3))
	fmt.Println("Buscar 2: ", avl.Search(2))
	fmt.Println("Buscar 5: ", avl.Search(5))
	fmt.Println("Buscar 12: ", avl.Search(12))
	fmt.Println("Buscar 10: ", avl.Search(10))
	fmt.Println("Buscar 4: ", avl.Search(4))
	fmt.Println("Borrar 7 y 12")
	avl.Remove(7)
	avl.Remove(12)
	fmt.Println("Raiz: ", avl.GetRoot().GetData())
	fmt.Println("Altura: ", avl.GetHeight())
	fmt.Println("Factor de balanceo: ", avl.GetBalance())
	fmt.Println("Max", avl.FindMax())
	fmt.Println("Min", avl.FindMin())
	fmt.Println("Recorrido inorden: ")
	it = avltree.NewAVLInOrderIterator[int](avl.GetRoot())
	for it.HasNext() {
		nodo, _ := it.Next()
		fmt.Printf("%v ", nodo.GetData())
	}
	fmt.Println()

	fmt.Println("Recorrido por niveles: ")
	itl = avltree.NewAVLLevelIterator[int](avl.GetRoot())
	anterior = avl.GetRoot().GetData()
	for itl.HasNext() {
		nodo, _ := itl.Next()
		if anterior > nodo.GetData() {
			fmt.Println()
		}
		anterior = nodo.GetData()
		fmt.Printf("%v ", nodo.GetData())
	}
	fmt.Println()
}
