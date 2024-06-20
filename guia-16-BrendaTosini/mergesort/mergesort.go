package main

import "fmt"

// Implementa el algoritmo de ordenamiento MergeSort para ordenar un arreglo de enteros de manera ascendente.
func MergeSort(array []int) []int {
	fmt.Println(":: ordenando", array) // Imprime el arreglo actual que está siendo ordenado

	// Caso base: si el tamaño del arreglo es menor a 2, no hay nada que ordenar, por lo que se retorna el arreglo original
	if len(array) < 2 {
		return array
	}

	// Se divide el arreglo en dos mitades aproximadamente iguales
	middle := len(array) / 2

	// Se realiza la llamada recursiva para ordenar cada mitad del arreglo
	left := MergeSort(array[:middle])  // Ordena la mitad izquierda del arreglo
	right := MergeSort(array[middle:]) // Ordena la mitad derecha del arreglo

	// Se combina (merge) las dos mitades ordenadas
	return merge(left, right)
}

// combina y mezcla dos arreglos ordenados en uno solo ordenado.
func merge(left, right []int) []int {
	fmt.Println("> mezclando", left, right) // Imprime los arreglos que están siendo mezclados

	size, i, j := len(left)+len(right), 0, 0
	array := make([]int, size) // Crea un nuevo arreglo con el tamaño necesario para almacenar los elementos de ambos arreglos

	// Combinar los elementos de los arreglos izquierdo y derecho en orden ascendente
	for k := 0; k < size; k++ {
		// Si se han recorrido todos los elementos del arreglo izquierdo y quedan elementos en el arreglo derecho
		if i > len(left)-1 && j <= len(right)-1 {
			array[k] = right[j] // Se agrega el siguiente elemento del arreglo derecho
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			array[k] = left[i] // Se agrega el siguiente elemento del arreglo izquierdo
			i++
		} else if left[i] < right[j] {
			array[k] = left[i] // Se agrega el siguiente elemento del arreglo izquierdo
			i++
		} else {
			array[k] = right[j] // Se agrega el siguiente elemento del arreglo derecho
			j++
		}
	}

	fmt.Println("< mezclado", array) // Imprime el arreglo mezclado resultante
	return array
}

func main() {
	//array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	array := []int{5, 1, 3, 4, 2}
	fmt.Println(array)
	fmt.Println(MergeSort(array))
}
