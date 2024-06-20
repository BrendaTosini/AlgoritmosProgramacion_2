package main

import "fmt"

func main() {

	var numero int

	fmt.Scanln(&numero)

	fmt.Println(esPrimo(numero))

}

func esPrimo(numero int) (esPrimo bool) {

	numerosDivisibles := 0

	for i := 1; i <= numero; i++ {

		if numero%i == 0 {

			numerosDivisibles++
		}

	}
	if numerosDivisibles == 2 {
		esPrimo = true
	} else {
		esPrimo = false
	}
	return

}
