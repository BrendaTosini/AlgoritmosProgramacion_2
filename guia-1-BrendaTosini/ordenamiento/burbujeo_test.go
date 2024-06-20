package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBurbujeo(t *testing.T) {

	arreglo := []int{5, 2, 4, 1, 3}

	resultado := Burbujeo(arreglo)

	esperado := []int{1, 2, 3, 4, 5}

	assert.Equal(t, resultado, esperado)

	if !assert.Equal(t, resultado, esperado) {
		t.Errorf("resultado = %v, esperado %v", resultado, esperado)
	}

}

func TestBurbujeo2(t *testing.T) {

	arreglo := []int{-5, 800, 2, -9, 35}

	resultado := Burbujeo(arreglo)

	esperado := []int{-9, -5, 2, 35, 800}

	assert.Equal(t, resultado, esperado)

	if !assert.Equal(t, resultado, esperado) {
		t.Errorf("resultado = %v, esperado %v", resultado, esperado)
	}

}
