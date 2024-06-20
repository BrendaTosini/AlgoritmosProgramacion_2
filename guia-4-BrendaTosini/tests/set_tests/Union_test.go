package tests

import (
	"guia4/set"
	"testing"
)

func TestUnion(t *testing.T) {
	s := set.NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s2 := set.NewSet[int]()
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)
	s2.Add(6)
	s3 := set.Union(s, s2)
	if s3.Size() != 6 {
		t.Errorf("Tamaño %d, debería ser %d", s.Size(), 6)
	}
	for i := 1; i <= 6; i++ {
		if !s3.Contains(i) {
			t.Errorf("Contains(%d) = %t, debería ser %t", i, s3.Contains(i), true)
		}
	}
}
