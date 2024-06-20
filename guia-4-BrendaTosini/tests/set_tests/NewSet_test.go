package tests

import (
	"guia4/set"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := set.NewSet[int]()
	if s == nil {
		t.Errorf("NewSet() = nil, want *Set")
	}

	s = set.NewSet(1, 2, 3, 4)
	if s.Size() != 4 {
		t.Errorf("Tamaño %d, debería ser %d", s.Size(), 4)
	}
	for i := 1; i <= 4; i++ {
		if !s.Contains(i) {
			t.Errorf("Contains(%d) = %t, debería ser %t", i, s.Contains(i), true)
		}
	}
}
