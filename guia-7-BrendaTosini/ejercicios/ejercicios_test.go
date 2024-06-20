package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEjercicio01(t *testing.T) {
	assert.Equal(t, 1, Suma(1))
	assert.Equal(t, 21, Suma(6))
	assert.Equal(t, 55, Suma(10))
}

func TestEjercicio02(t *testing.T) {
	assert.Equal(t, 2, Factorial(2))
	assert.Equal(t, 24, Factorial(4))
	assert.Equal(t, 120, Factorial(5))
}

func TestEjercicio03(t *testing.T) {
	assert.Equal(t, 0, CantidadDeUnos(0))
	assert.Equal(t, 1, CantidadDeUnos(4))
	assert.Equal(t, 2, CantidadDeUnos(5))
	assert.Equal(t, 3, CantidadDeUnos(42))
	assert.Equal(t, 2, CantidadDeUnos(6))

}

func TestEjercicio04(t *testing.T) {
	assert.Equal(t, "aloH", Invertir("Hola"))
	assert.Equal(t, "neuqueN", Invertir("Neuquen"))
	assert.Equal(t, "!odnuM ,aloH", Invertir("Hola, Mundo!"))
}

func TestEjercicio05(t *testing.T) {
	var pila Stack
	pila.Push(1)
	pila.Push(2)
	pila.Push(3)

	assert.Equal(t, 3, CantidadDeElementos(pila))
}

func TestEjercicio06(t *testing.T) {
	assert.Equal(t, 12, MCD(12, 60))
	assert.Equal(t, 12, MCD(60, 12))
	assert.Equal(t, 1, MCD(7, 13))
	assert.Equal(t, 1, MCD(13, 7))
	assert.Equal(t, 2, MCD(2, 4))
	assert.Equal(t, 2, MCD(4, 2))
	assert.Equal(t, 3, MCD(3, 9))
	assert.Equal(t, 3, MCD(9, 3))
}

func TestEjercicio07(t *testing.T) {
	assert.Equal(t, 16, Multiplicar(2, 8))
	assert.Equal(t, 16, Multiplicar(8, 2))
	assert.Equal(t, 0, Multiplicar(0, 8))
	assert.Equal(t, 0, Multiplicar(8, 0))
}

func TestEjercicio08(t *testing.T) {
	cociente, resto := DivisionEntera(8, 2)
	assert.Equal(t, 4, cociente)
	assert.Equal(t, 0, resto)

}

func TestEjercicio09(t *testing.T) {
	assert.Equal(t, 0, SumaArray([]int{}))
	assert.Equal(t, 1, SumaArray([]int{1}))
	assert.Equal(t, 3, SumaArray([]int{1, 2}))
	assert.Equal(t, 6, SumaArray([]int{1, 2, 3}))
}

func TestEjercicio10(t *testing.T) {
	assert.Equal(t, "101010", DecimalABinario(42))
	assert.Equal(t, "10000000000000000000000000000000", DecimalABinario(2147483648))
	assert.Equal(t, "0", DecimalABinario(0))
	assert.Equal(t, "1", DecimalABinario(1))
	assert.Equal(t, "1100", DecimalABinario(12))
}

func TestEjercicio11(t *testing.T) {
	assert.PanicsWithValue(t, "No es potencia de 0", func() { EsPotencia(2, 0) })
	assert.False(t, EsPotencia(10, 3))
	assert.False(t, EsPotencia(2, 9))
	assert.True(t, EsPotencia(8, 2))
	assert.True(t, EsPotencia(1, 2))
	assert.False(t, EsPotencia(7, 2))

}

func TestEjercicio12(t *testing.T) {
	assert.Equal(t, 5, Fibonacci(5))
	assert.Equal(t, 8, Fibonacci(6))
	assert.Equal(t, 13, Fibonacci(7))
	assert.Equal(t, 21, Fibonacci(8))
}

func TestEjercicio13(t *testing.T) {
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 1))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 2))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 3))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 4))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 5))
	assert.False(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 6))
}

func TestEjercicio14(t *testing.T) {
	assert.False(t, SumaElementos([]int{7, 4, 6, 8}, []int{3, 1, 6, 6}, 2))
	assert.True(t, SumaElementos([]int{7, 4, 6, 8}, []int{3, 1, 6, 6}, 7))

	assert.True(t, SumaElementos([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 4))
	assert.False(t, SumaElementos([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 11))
}

func TestEjercicio15(t *testing.T) {
	assert.Equal(t, 4, Pico([]int{1, 2, 3, 4, 5, 4, 3, 2, 1}))
	assert.Equal(t, 8, Pico([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	assert.Equal(t, 0, Pico([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}
