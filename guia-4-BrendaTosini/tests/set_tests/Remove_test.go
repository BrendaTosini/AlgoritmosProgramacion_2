package tests

import (
	"guia4/set"
	"testing"
)

func TestRemove(t *testing.T) {
	s := set.NewSet(1, 3, 2)
	s.Remove(1)
	s.Remove(3)
	s.Remove(1)
	s.Remove(3)
	if s.Size() != 1 {
		t.Errorf("Tamaño %d, debería ser: %d", s.Size(), 1)
	}
}
