package ej04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCasoBase(t *testing.T) {
	valores := []int{20, 30, 15, 25, 10}
	pesos := []int{6, 13, 5, 10, 3}
	k := 20

	assert.Equal(t, 55, Mochila(valores, pesos, k))
}
