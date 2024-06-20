package ejercicios

import "math"

// Modificar el código de RadixSort, para que en lugar de ordenar los números de
// menor a mayor los ordene de mayor a menor.
// Escribir tests unitarios para la función RadixSort
func getDigit(num, pos int) int {
	return (num / int(math.Pow10(pos))) % 10
}

func RadixSort(arr []int) {
	// Obtenemos el máximo valor del arreglo para determinar el número de dígitos
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	// Determinamos la cantidad de dígitos necesarios
	digits := int(math.Log10(float64(max))) + 1

	// Realizamos el ordenamiento por cada dígito
	for i := 0; i < digits; i++ {
		// Creamos las "bandejas" para cada dígito
		buckets := make([][]int, 10)
		for j := 0; j < 10; j++ {
			buckets[j] = make([]int, 0)
		}

		// Distribuimos los números en las urnas según el dígito actual
		for _, num := range arr {
			digit := getDigit(num, i)
			buckets[digit] = append(buckets[digit], num)
		}

		// Copiamos los números de las urnas al arreglo original
		idx := 0
		for j := 0; j < 10; j++ {
			for _, num := range buckets[j] {
				arr[idx] = num
				idx++
			}
			buckets[j] = nil
		}
	}
}
