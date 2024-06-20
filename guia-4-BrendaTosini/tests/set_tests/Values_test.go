package tests

import (
	"guia4/set"
	"testing"
)

func TestValues(t *testing.T) {
	s := set.NewSet(1, 2, 3, 4)
	v := s.Values()
	if len(v) != 4 {
		t.Errorf("Tamaño %d, debería ser: %d", len(v), 4)
	}
	for i := 1; i <= 4; i++ {
		if !s.Contains(i) {
			t.Errorf("Contains(%d) = %t, debería ser %t", i, s.Contains(i), true)
		}
	}
}
