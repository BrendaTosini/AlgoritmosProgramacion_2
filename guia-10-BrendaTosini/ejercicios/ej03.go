package ejercicios

func Mochila2(Objetos []Objeto, capacity int) (int, []Objeto) {
	bestValor := 0                                // Mejor valor encontrado
	bestSubset := make([]Objeto, len(Objetos))    // Conjunto de objetos a cargar en la mochila
	currentSubset := make([]Objeto, len(Objetos)) //Solucion parcial para el backtracking

	backtrack(Objetos, 0, capacity, 0, &bestValor, currentSubset, bestSubset)

	return bestValor, bestSubset
}

func backtrack(Objetos []Objeto, index, capacity, currentValor int, bestValor *int, currentSubset, bestSubset []Objeto) {
	panic("Implementar")
}
