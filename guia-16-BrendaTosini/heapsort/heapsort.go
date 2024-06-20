package main

import "fmt"

// HeapSort ordena un arreglo utilizando el algoritmo HeapSort.
// Recibe el arreglo a ordenar como parámetro.
func HeapSort(array []int) {
	size := len(array)
	heapify(array) // Construye un heap máximo a partir del arreglo

	// En cada iteración, se extrae el elemento máximo del heap y se coloca al final del arreglo.
	// Luego, se ajusta el heap hacia abajo para mantener la propiedad del heap.
	end := size - 1
	for end > 0 {
		array[end], array[0] = array[0], array[end] // Intercambia el máximo actual con el último elemento del arreglo
		downHeap(array, 0, end-1)                   // Ajusta el heap hacia abajo (restaura la propiedad del heap)
		end--                                       // Reduce el tamaño efectivo del arreglo en 1 para excluir el elemento ya ordenado
	}
}

// heapify convierte el arreglo en un heap.
// Recibe el arreglo como parámetro y modifica el arreglo en su lugar.
func heapify(array []int) {
	size := len(array)
	// El primer nodo que tiene hijos se encuentra en la posición (size - 2) / 2.
	start := (size - 2) / 2

	// Comienza desde el último padre y ajusta cada subárbol hacia abajo para cumplir la propiedad del heap.
	for start >= 0 {
		downHeap(array, start, size-1)
		start--
	}
}

// downHeap realiza el ajuste descendente (downHeap) del nodo padre en un heap.
// Recibe un arreglo, un nodo padre y el último nodo del heap como parámetros.
func downHeap(array []int, start, end int) {
	father := start
	leftSon := father*2 + 1
	rightSon := leftSon + 1

	// Mientras el padre tenga al menos un hijo
	for leftSon <= end {
		// Si el padre tiene dos hijos, nos quedamos con el mayor
		if rightSon <= end && array[rightSon] > array[leftSon] {
			leftSon = rightSon
		}

		// Si el hijo es mayor que el padre, los intercambiamos
		if array[leftSon] > array[father] {
			array[leftSon], array[father] = array[father], array[leftSon]
			// El hijo se convierte en el padre
			father = leftSon
			leftSon = father*2 + 1
			rightSon = leftSon + 1
		} else {
			return
		}
	}
}

func main() {
	//array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	array := []int{2, -4, 7, 9, 0, 25, -5, 13, 17, 6}
	fmt.Println(array)
	HeapSort(array)
	fmt.Println(array)
}
