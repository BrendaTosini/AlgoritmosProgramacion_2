package main

import (
	"figuras/figuras"
	"fmt"
)

func main() {
	var figura string
	var arreglo []figuras.Figura

	for i := 0; i < 5; i++ {
		fmt.Println("Elija una figura: Cuadrado,Rectangulo o Circulo")
		fmt.Scan(&figura)
		arreglo = menuFiguras(figura, arreglo)
	}

	for i := 0; i < len(arreglo); i++ {
		fmt.Println(arreglo[i].ToString())
	}
}

func menuFiguras(f string, arreglo []figuras.Figura) []figuras.Figura {

	var x int
	var y int
	var lado int
	var radio int

	switch f {

	case "Circulo":

		fmt.Println("Ingrese coordenada x del punto")
		fmt.Scan(&x)
		fmt.Println("Ingrese coordenada y del punto")
		fmt.Scan(&y)
		fmt.Println("Ingrese el radio")
		fmt.Scan(&radio)
		pto := figuras.Punto{X: x, Y: y}
		circulo := figuras.Circulo{Punto: pto, Radio: radio}
		arreglo = append(arreglo, circulo)

	case "Cuadrado":

		fmt.Println("Ingrese coordenada x del punto")
		fmt.Scan(&x)
		fmt.Println("Ingrese coordenada y del punto")
		fmt.Scan(&y)
		fmt.Println("Ingrese el tamaño del lado")
		fmt.Scan(&lado)
		pto := figuras.Punto{X: x, Y: y}
		cuadrado := figuras.Cuadrado{Pto: pto, Lado: lado}
		arreglo = append(arreglo, cuadrado)

	case "Rectangulo":
		fmt.Println("Ingrese la coordenada x del punto 1")
		fmt.Scan(&x)
		fmt.Println("Ingrese la coordenada y del punto 1")
		fmt.Scan(&y)
		pto1 := figuras.Punto{X: x, Y: y}
		fmt.Println("Ingrese la coordenada x del punto 2")
		fmt.Scan(&x)
		fmt.Println("Ingrese la coordenada y del punto 2")
		fmt.Scan(&y)
		pto2 := figuras.Punto{X: x, Y: y}
		rectangulo := figuras.Rectangulo{P1: pto1, P2: pto2}
		arreglo = append(arreglo, rectangulo)
	}
	return arreglo
}
