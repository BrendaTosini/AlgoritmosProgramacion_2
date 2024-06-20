package ej05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEjemplo01(t *testing.T) {
	arr := []int{1, 2, 3, 1}
	k := 4

	assert.Equal(t, 3, SumaSubconjunto(arr, k))
}

func TestEjemplo02(t *testing.T) {
	arr := []int{1, 2, 3, 1, 4}
	k := 6

	assert.Equal(t, 4, SumaSubconjunto(arr, k))
}

func TestEjemplo03(t *testing.T) {
	arr := []int{2, 4, 2, 6, 8}
	k := 7

	assert.Equal(t, 0, SumaSubconjunto(arr, k))
}
