package main

import (
	"fmt"
)

func main() {

	fmt.Println(Factorial(5))
}

func Factorial(numero int) (factorial int) {

	esPositivo(numero)
	factorial = esCero(numero, factorial)

	if numero != 0 {

		factorial = numero

		for i := numero - 1; i > 0; i-- {

			factorial *= i

		}
	}
	return
}

func esPositivo(numero int) {

	if numero < 0 {
		err := fmt.Errorf("%v No se puede hacer el factorial de un numero negativo", numero)
		panic(err)
	}

}

func esCero(numero, factorial int) (factorialFinal int) {

	if numero == 0 {
		factorial = 1
		factorialFinal = factorial
	}
	return factorialFinal
}
