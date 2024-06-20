package ej02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCasoBase(t *testing.T) {
	s1 := "abdacbab"
	s2 := "acebfca"

	lcs := LCS(s1, s2)

	assert.Equal(t, 4, lcs)
}
