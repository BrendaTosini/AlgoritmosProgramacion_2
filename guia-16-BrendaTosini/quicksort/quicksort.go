package main

import "fmt"

// implementa el algoritmo de ordenamiento QuickSort para ordenar un arreglo de enteros de manera ascendente.
func QuickSort(array []int) {
	if len(array) < 2 {
		return // Si el tamaño de l arreglo es menor a 2, no hay nada que ordenar, por lo que se retorna
	}
	pivot := array[0]   // Se selecciona el primer elemento como pivote
	i := 1              // Índice para recorrer el arreglo desde el segundo elemento
	j := len(array) - 1 // Índice para recorrer el arreglo desde el último elemento

	// Se realiza el particionamiento del arreglo
	for i < j {
		// Se busca un elemento mayor al pivote desde la izquierda
		for array[i] < pivot && i < len(array)-1 {
			i++
		}
		// Se busca un elemento menor al pivote desde la derecha
		for array[j] > pivot && j > 0 {
			j--
		}
		if i < j {
			// Se intercambian los elementos encontrados
			array[i], array[j] = array[j], array[i]
		}
	}

	// Se coloca el pivote en su posición final
	if array[j] < pivot {
		array[0], array[j] = array[j], array[0]
	}

	// Se realiza la ordenación recursiva de las sublistas
	QuickSort(array[:j])   // Ordenar la sublista izquierda del pivote
	QuickSort(array[j+1:]) // Ordenar la sublista derecha del pivote
}

func main() {
	//array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	array := []int{4, 1, 5, 3, 2}
	fmt.Println(array)
	QuickSort(array)
	fmt.Println(array)
}
