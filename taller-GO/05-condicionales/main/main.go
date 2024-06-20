package main

import (
	"condicionales/condicionales"
	"fmt"
	"time"
)

func main() {
	var numero int

	fmt.Print("Ingrese un entero: ")
	fmt.Scanln(&numero)

	condicionales.Condicional(numero)

	condicionales.Swtich_basico()
	condicionales.Switch_sin_condicion(time.Now().Local().Hour())
	condicionales.Switch_multiple(' ')
	condicionales.Switch_fallthrough(2)
}
