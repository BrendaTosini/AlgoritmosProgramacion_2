package main

import "fmt"

func main() {

	arreglo := []int{45, 13, 6, 2, 0, 598}
	fmt.Println(MayorYmenor(arreglo))

}

func MayorYmenor(arr []int) (menor, mayor int) {

	menor = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < menor {
			menor = arr[i]
		} else {
			mayor = arr[i]
		}
	}
	return menor, mayor

}
