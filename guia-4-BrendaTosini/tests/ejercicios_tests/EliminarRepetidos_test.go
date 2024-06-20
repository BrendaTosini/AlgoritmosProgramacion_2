package tests

import (
	"guia4/ejercicios"
	"testing"
)

func TestEliminarRepetidos(t *testing.T) {
	arreglo := []int{1, 2, 3, 4, 4, 3, 2, 1}
	arreglo = ejercicios.EliminarRepetidos(arreglo)
	if len(arreglo) != 4 {
		t.Errorf("Tamaño %d, debería ser: %d", len(arreglo), 4)
	}
}
