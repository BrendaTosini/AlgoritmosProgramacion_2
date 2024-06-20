package ejercicios

// Escriba un método recursivo que calcule Fibonacci de n.
func Fibonacci(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	resultado := Fibonacci(n-1) + Fibonacci(n-2)

	return resultado
}
