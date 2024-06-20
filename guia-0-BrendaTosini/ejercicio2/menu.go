package main

import (
	"fmt"
)

func main() {
	Opciones(1)
}

func Opciones(numero int) {

	switch numero {

	case 1:
		fmt.Println("Opción 1")

	case 2:
		fmt.Println("Opción 2")

	case 3:
		fmt.Println("Opción 3")

	case 4:
		fmt.Println("Opción 4")

	default:
		fmt.Println("Opción incorrecta")
	}

}
