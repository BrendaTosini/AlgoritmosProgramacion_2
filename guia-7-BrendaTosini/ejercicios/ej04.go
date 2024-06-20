package ejercicios

// Escriba un método recursivo que tome una cadena y
// devuelva como resultado la cadena invertida.
func Invertir(cadena string) string {

	if len(cadena) == 1 {
		return cadena
	}

	ultimaLetra := string(cadena[len(cadena)-1])

	cadena = cadena[0 : len(cadena)-1]

	invertida := ultimaLetra + Invertir(cadena)

	return invertida
}
