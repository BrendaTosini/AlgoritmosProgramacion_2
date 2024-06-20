package ejercicios

import (
	"fmt"
	"testing"
)

func TestMochila3(t *testing.T) {
	objetos := []Objeto{
		{Peso: 1, Valor: 1},
		{Peso: 2, Valor: 6},
		{Peso: 3, Valor: 18},
		{Peso: 4, Valor: 22},
		{Peso: 5, Valor: 28},
	}
	capacidad := 11
	valor, solucion := Mochila2(objetos, capacidad)
	fmt.Println(solucion)
	if valor != 56 {
		t.Errorf("Mochila3(objetos, capacidad) = %d; se esperaba 56", valor)
	}
}

func TestMochila4(t *testing.T) {
	objetos := []Objeto{
		{Peso: 1, Valor: 10},
		{Peso: 2, Valor: 6},
		{Peso: 3, Valor: 18},
		{Peso: 4, Valor: 22},
		{Peso: 5, Valor: 28},
	}
	capacidad := 11
	valor, solucion := Mochila2(objetos, capacidad)
	fmt.Println(solucion)
	if valor != 62 {
		t.Errorf("Mochila2(objetos, capacidad) = %d; se esperaba 62", valor)
	}
}
