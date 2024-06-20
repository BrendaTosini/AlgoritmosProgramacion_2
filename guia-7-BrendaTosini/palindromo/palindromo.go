package palindromo

func EsPalindromo(texto string) bool {
	if len(texto) <= 1 {
		return true
	}

	primero := string(texto[0])
	ultimo := string(texto[len(texto)-1])
	interior := texto[1 : len(texto)-1]

	return primero == ultimo && EsPalindromo(interior)
}

func EsPalindromo2(texto string) bool {
	if len(texto) <= 1 {
		return true
	}

	primero := string(texto[0])
	ultimo := string(texto[len(texto)-1])
	interior := texto[1 : len(texto)-1]

	if primero == " " {
		primero = string(texto[1])
		interior = texto[2 : len(texto)-1]
	}
	if ultimo == " " {
		ultimo = string(texto[len(texto)-2])
		interior = texto[1 : len(texto)-2]
	}

	return primero == ultimo && EsPalindromo2(interior)
}
