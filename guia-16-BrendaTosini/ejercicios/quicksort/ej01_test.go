package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedio(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	expected := 6
	kesimo := EncontrarKesimo(array, 5)
	assert.Equal(t, expected, kesimo)
}

func TestPrimero(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	expected := 1
	kesimo := EncontrarKesimo(array, 0)
	assert.Equal(t, expected, kesimo)
}

func TestUltimo(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	expected := 10
	kesimo := EncontrarKesimo(array, 9)
	assert.Equal(t, expected, kesimo)
}
