package ejercicios

// Escriba un método recursivo que tome un entero n
// devuelva su factorial
func Factorial(n int) int {

	if n == 0 || n == 1 {
		return 1
	}

	resultado := n * Factorial(n-1)
	return resultado
}
