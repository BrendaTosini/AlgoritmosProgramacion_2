package ej01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCasoBase(t *testing.T) {
	matriz := [][]int{
		{3, 2, 12, 15, 10},
		{6, 19, 7, 11, 17},
		{8, 5, 12, 32, 21},
		{3, 20, 2, 9, 7},
	}

	assert.Equal(t, 52, CostoCaminoMinimo(matriz))
}
