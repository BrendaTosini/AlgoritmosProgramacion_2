package ejercicios

// Escribir un método recursivo que reciba 2 enteros
// n y b y devuelva true si n es potencia de b.
// Por ejemplo: esPotencia(8, 2) devuelve true.
func EsPotencia(n, b int) bool {

	if b == 0 {
		panic("No es potencia de 0")
	}

	if n == 1 {
		return true
	}

	if n%b != 0 {
		return false
	}

	resultado := EsPotencia(n/b, b)

	return resultado
}
