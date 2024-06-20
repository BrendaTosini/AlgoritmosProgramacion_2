package tests

import (
	"guia4/set"
	"testing"
)

func TestAdd(t *testing.T) {
	s := set.NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(1)
	s.Add(2)
	s.Add(3)
	if s.Size() != 3 {
		t.Errorf("Tamaño %d, debería ser %d", s.Size(), 3)
	}
}
