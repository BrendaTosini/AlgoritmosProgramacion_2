package main

import "fmt"

func main() {

	polinomio := Polinomio(3.0, 2.0, 1.0)
	fmt.Println(polinomio)

}

func Polinomio(coeficientes ...float32) (total string) {

	for indices, coeficientes := range coeficientes {

		if indices > 0 {

			total += fmt.Sprintf(" + %.1fx^%v", coeficientes, indices)
		} else {

			total += fmt.Sprintf("%.1f", coeficientes)

		}

	}

	return total

}
