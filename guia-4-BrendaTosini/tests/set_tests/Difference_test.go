package tests

import (
	"guia4/set"
	"testing"
)

func TestDifference(t *testing.T) {
	s := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4, 5, 6)
	s3 := set.Difference(s, s2)
	if s3.Size() != 2 {
		t.Errorf("Tamaño %d, debería ser %d", s.Size(), 2)
	}
	for i := 1; i <= 2; i++ {
		if !s3.Contains(i) {
			t.Errorf("Contains(%d) = %t, debería ser %t", i, s3.Contains(i), true)
		}
	}
}
