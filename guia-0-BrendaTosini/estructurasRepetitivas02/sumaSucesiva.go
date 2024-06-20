package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumaSucesiva(2, 1))

}

func sumaSucesiva(numero1, numero2 int) (producto int) {

	for i := 0; i < AbsInt(numero2); i++ {
		producto += numero1
	}
	if numero1 > 0 && numero2 < 0 || numero1 < 0 && numero2 < 0 {
		producto *= (-1)
	}

	return
}

func AbsInt(numero int) int {
	if numero < 0 {
		return numero * (-1)
	}
	return numero
}
