package main

import "fmt"

func main() {

	var arreglo1 = [3]int{1, 2, 3}
	var arreglo2 = [3]int{1, 2, 3}

	fmt.Println(sumaYproductoEscalar(arreglo1, arreglo2))

}

func sumaYproductoEscalar(arr1 [3]int, arr2 [3]int) (producto int, arregloSuma [3]int) {

	for i := 0; i < len(arr1); i++ {

		producto += arr1[i] * arr2[i]
		arregloSuma[i] = arr1[i] + arr2[i]

	}
	return
}
