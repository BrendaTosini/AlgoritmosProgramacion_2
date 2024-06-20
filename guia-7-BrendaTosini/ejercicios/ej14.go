package ejercicios

// Sean as y bs dos listas de enteros de tamaño n.
// Escribir una función que reciba como parámetros
// as, bs y un entero x y decida si x se puede
// escribir como suma de un elemento de as más un
// elemento de bs.
func SumaElementos(as, bs []int, x int) bool {

	if len(as) == 0 || len(bs) == 0 {
		return false
	}

	if as[0]+bs[0] == x {
		return true
	}

	resultado := SumaElementos(as[1:], bs, x) || SumaElementos(as, bs[1:], x)

	return resultado

}
