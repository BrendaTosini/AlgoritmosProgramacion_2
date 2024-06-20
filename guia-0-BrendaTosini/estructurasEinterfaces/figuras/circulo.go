package figuras

import (
	"fmt"
)

type Circulo struct {
	Punto Punto
	Radio int
}

func (c Circulo) Area() int {
	return 3 * c.Radio * c.Radio
}
func (c Circulo) Perimetro() int {
	return 2 * 3 * c.Radio
}
func (c Circulo) ToString() string {
	cadena := fmt.Sprint("Circulo: ", c)
	return cadena
}
