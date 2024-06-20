package ejercicios

// Escriba un método recursivo que tome un entero n
// y devuelva la suma de los primeros n números naturales.
func Suma(n int) int {

	if n == 0 || n == 1 {
		return n
	}

	resultado := n + Suma(n-1)
	return resultado
}
