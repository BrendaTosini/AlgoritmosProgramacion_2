package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeWayMergeSort(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sorted := ThreeWayMergeSort(array)
	assert.Equal(t, expected, sorted)
}
