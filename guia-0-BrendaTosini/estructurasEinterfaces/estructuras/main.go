package main

import (
	"figuras/figuras"
	"fmt"
)

func main() {
	p := figuras.Punto{X: 0, Y: 0}
	fmt.Println(p.ToString())
	p.Mover(10, -10)
	fmt.Println(p.ToString())

	p1 := figuras.Punto{0, 0}
	p2 := figuras.Punto{5, 3}
	r := figuras.Rectangulo{P1: p1, P2: p2}

	fmt.Println("Area del Rectangulo: ", r.Area())
	r.Mover(12, -10)
	fmt.Println("Area del Rectangulo: ", r.Area())
	fmt.Println("Perimetro del Rectangulo: ", r.Perimetro())
	fmt.Println(r.ToString())

	c := figuras.Cuadrado{Pto: p1, Lado: 5}
	fmt.Println("Area del Cuadrado: ", c.Area())
	fmt.Println("Perimetro del Cuadrado: ", c.Perimetro())
	fmt.Println(c.ToString())
}
