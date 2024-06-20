package ej03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCasoBase(t *testing.T) {
	n := 10
	formas := []int{2, 4, 5, 8}

	escalera := Escalera(n, formas)

	assert.Equal(t, 11, escalera)
}
