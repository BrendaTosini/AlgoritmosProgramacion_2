package main

import (
	"fmt"
)

// arrUnion := []int{} // para agregar a un slice vacio se usa append. El slice es una tajada de un array original.

func main() {

	arreglo1 := []int{1, 2, 3}
	arreglo2 := []int{1, 5, 6, 7}

	fmt.Println(UnionEinterseccion(arreglo1, arreglo2))
}

func UnionEinterseccion(arr1 []int, arr2 []int) (arrUnion, arrInterseccion []int) {

	arrUnion = agregarArray(arr1, arrUnion)

	//comparo arr2 con arr1. Si encuentro que un elemento es igual al del otro array lo agrego en arrInterseccion. Si hay un elemento que no es igual y no esta en el arrUnion lo agrego.

	for i := 0; i < len(arr2); i++ {

		for j := 0; j < len(arr1); j++ {

			if arr2[i] == arr1[j] && !estaEnArreglo(arrInterseccion, arr2[i]) {
				arrInterseccion = append(arrInterseccion, arr2[i])
				j = len(arr1)

			} else if !estaEnArreglo(arrUnion, arr2[i]) {
				arrUnion = append(arrUnion, arr2[i])
				j = len(arr1)

			}
		}
	}
	return
}

//agrego primer array a arrUnion sin numeros repetidos.

func agregarArray(arr, arrU []int) []int {

	arrU = append(arrU, arr[0])

	i := 0

	for i < len(arr) {

		for j := 0; j < len(arr); j++ {

			if arr[i] != arr[j] {
				arrU = append(arrU, arr[j])
			}
		}
		i = len(arr)
	}
	return arrU

}

//pregunto si un numero esta en un array.

func estaEnArreglo(arr []int, n int) bool {

	resultado := false

	for i := 0; i < len(arr); i++ {
		if n == arr[i] {
			resultado = true
		}
	}
	return resultado
}
