package ejercicios

import "fmt"

// Escribir un método recursivo que dado un número
// entero decimal devuelva su equivalente en
// binario en forma de string

func DecimalABinario(n int) string {

	if n == 0 {
		return fmt.Sprint(n)
	}

	if n == 1 {
		return fmt.Sprint(n)
	}

	binario := DecimalABinario(n/2) + fmt.Sprint(n%2)

	return binario
}
