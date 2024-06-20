package ejercicios

// Escribir un método recursivo que devuelva el
// cociente y el resto de la división entera mediante
// restas sucesivas
func DivisionEntera(dividendo, divisor int) (cociente, resto int) {

	if dividendo >= divisor {
		cociente, _ = DivisionEntera(dividendo-divisor, divisor)
		cociente++
		resto = dividendo - (divisor * cociente)
	}
	return cociente, resto
}
