package ejercicios

// Escriba un método recursivo que calcule el máximo
// común divisor entre dos números enteros.
// Nota: Se puede usar el algoritmo de Euclides para
// resolver este problema.
func MCD(a, b int) int {

	var resultado int

	if a > b {

		if a%b == 0 {
			return b

		} else {

			resto := a % b

			resultado = MCD(b, resto)
		}
	} else {
		resultado = MCD(b, a)
	}
	return resultado
}
