package tests

import (
	"guia4/set"
	"testing"
)

func TestIntersection(t *testing.T) {
	s := set.NewSet(1, 3, 4, 2)
	s2 := set.NewSet(6, 5, 4, 3)
	s3 := set.Intersection(s, s2)
	if s3.Size() != 2 {
		t.Errorf("Tamaño %d, debería ser %d", s.Size(), 2)
	}
	for i := 3; i <= 4; i++ {
		if !s3.Contains(i) {
			t.Errorf("Contains(%d) = %t, debería ser %t", i, s3.Contains(i), true)
		}
	}
}
