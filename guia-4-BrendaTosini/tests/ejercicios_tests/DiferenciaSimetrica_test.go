package tests

import (
	"guia4/ejercicios"
	"guia4/set"
	"testing"
)

func TestDiferenciaSimetrica(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4, 5, 6)
	s3 := ejercicios.DiferenciaSimetrica(s1, s2)
	if s3.Size() != 4 {
		t.Errorf("Tamaño %d, debería ser: %d", s3.Size(), 4)
	}
}

func TestNilpotnecia(t *testing.T) {
	panic("Not implemented")
}

func TestAsociativa(t *testing.T) {
	panic("Not implemented")
}

func TestConmutativa(t *testing.T) {
	panic("Not implemented")
}

func TestNulo(t *testing.T) {
	panic("Not implemented")
}
