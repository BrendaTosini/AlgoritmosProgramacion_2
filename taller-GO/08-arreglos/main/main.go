package main

import (
	"arreglos/arreglos"
	"fmt"
)

func main() {
	a1 := []int{-1, 5, -6, 10, 7, 2, -10, 5, 3}
	a2 := a1
	a3 := []int{-1, 5, -6, 10, 7, 2, -10, 5, 3, 4}

	fmt.Println(a1)
	arreglos.Insertion_sort(a1)
	fmt.Println(a1)
	fmt.Println(a2)

	fmt.Println("la suma de los elementos del arreglo es: ", arreglos.Sumar(a2))

	fmt.Println("Arreglo: ", a3)
	suma_maxima, inicial, final := arreglos.SubsecuenciaSumaMaxima(a3)
	fmt.Printf("La subsecuencia de suma maxima empieza en %v, termina en %v "+
		"y su suma es %v\n", inicial, final, suma_maxima)

	// Tajadas o Slices
	tajada := a3[inicial:]
	fmt.Println(tajada)

	for i := 0; i < len(tajada); i++ {
		tajada[i] = i
	}
	fmt.Println(tajada)
	fmt.Println(a3)

	t1 := a3[5:]
	t2 := a3[:4]
	t3 := a3[:]

	fmt.Println("t1 = ", t1)
	fmt.Println("t2 = ", t2)
	fmt.Println("t3 = ", t3)

	t3[0] = 100

	fmt.Println("a3 = ", a3)
	//append toma una tajada y le agrega elementos, devolviendo un nuevo arreglo
	t1 = append(t1, 33, 44)
	fmt.Println("t1 = ", t1)

	//con la sintaxis a3... puedo pasarle todos los elementos de un arreglo
	//a una función variádica
	t1 = append(t1, a3...)
	fmt.Println("t1 = ", t1)
}
