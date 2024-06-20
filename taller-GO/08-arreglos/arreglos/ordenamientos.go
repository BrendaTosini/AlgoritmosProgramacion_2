package arreglos

// Método de ordenamiento por inserción, recibe como parámetro
// la referencia al arreglo a ordenar
func Insertion_sort(arreglo []int) {
	i := 1
	for i < len(arreglo) {
		j := i
		for j > 0 && arreglo[j-1] > arreglo[j] {
			swap(j-1, j, arreglo)
			j--
		}
		i++
	}
}

// Intercambia los elementos en las posiciones i y j de un arreglo
// recibe como parámetro un puntero al arreglo para que las modificaciones
// se realicen sobre el arreglo original y no sobre una copia
func swap(i int, j int, arreglo []int) {
	aux := arreglo[i]
	arreglo[i] = arreglo[j]
	arreglo[j] = aux
}
