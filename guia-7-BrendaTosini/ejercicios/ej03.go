package ejercicios

// Escriba un método recursivo que devuelva la cantidad de unos
// en la representación binaria de un número n. Use el hecho de
// que es igual al número de unos en la representación binaria
// de n/2, mas 1 si es impar.
func CantidadDeUnos(n int) int {

	if n == 0 || n == 1 {
		return n
	}
	cantidad := 0

	if n%2 == 1 {
		cantidad = 1 + CantidadDeUnos(n/2)
	} else {
		cantidad = CantidadDeUnos(n / 2)
	}
	return cantidad
}
