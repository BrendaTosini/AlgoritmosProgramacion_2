package main

import "fmt"

func main() {

	arreglo := []int{1, 2, 3}
	fmt.Println(suma(arreglo))

}

func suma(arreglo []int) (suma int) {

	for i := 0; i < len(arreglo); i++ {

		suma += arreglo[i]
	}
	return

}
