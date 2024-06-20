package ejercicios

import (
	"testing"
)

func TestMochila01(t *testing.T) {
	objetos := []Objeto{
		{Peso: 1, Valor: 1},
		{Peso: 2, Valor: 6},
		{Peso: 3, Valor: 18},
		{Peso: 4, Valor: 22},
		{Peso: 5, Valor: 28},
	}
	capacidad := 11
	valor := Mochila01(objetos, capacidad)
	if valor != 40 {
		t.Errorf("Mochila01(objetos, capacidad) = %d; want 40", valor)
	}
}
