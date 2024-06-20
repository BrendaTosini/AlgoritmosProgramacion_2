package ejercicios

// Escriba un método recursivo para determinar si
// un elemento está en un arreglo utilizando búsqueda binaria.
func BusquedaBinaria(arreglo []int, elemento int) bool {

	inicio := 0
	final := len(arreglo) - 1
	medio := (inicio + final) / 2
	var resultado bool

	if len(arreglo) == 1 && arreglo[medio] != elemento {
		return resultado
	}

	switch {
	case arreglo[medio] > elemento:
		arreglo = arreglo[:medio]
		resultado = BusquedaBinaria(arreglo, elemento)
	case arreglo[medio] < elemento:
		arreglo = arreglo[medio+1:]
		resultado = BusquedaBinaria(arreglo, elemento)
	default:
		resultado = true
	}

	/*if arreglo[medio] > elemento {
		arreglo = arreglo[:medio]
		resultado = BusquedaBinaria(arreglo, elemento)
	} else if arreglo[medio] < elemento {
		arreglo = arreglo[medio+1:]
		resultado = BusquedaBinaria(arreglo, elemento)
	} else {
		resultado = true
	}
	*/

	return resultado
}
