package subsumamax

func SubsecuenciaSumaMaxima2(arreglo []int) (int, int, int) {

	sumaMaxima, sumaLocal := 0, 0
	posInicial, posFinal := 0, 0
	for i := 0; i < len(arreglo); i++ {
		sumaLocal += arreglo[i]
		if sumaLocal > sumaMaxima {
			sumaMaxima = sumaLocal
			posFinal = i
		}
		if sumaLocal < 0 {
			sumaLocal = 0
			posInicial = i + 1
		}
	}
	return sumaMaxima, posInicial, posFinal
}
