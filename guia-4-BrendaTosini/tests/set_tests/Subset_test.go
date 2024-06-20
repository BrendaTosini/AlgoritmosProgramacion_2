package tests

import (
	"guia4/set"
	"testing"
)

func TestSubset(t *testing.T) {
	s := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4)
	if !set.Subset(s, s2) {
		t.Errorf("Subset(%v, %v) = %t, debería ser %t", s, s2, set.Subset(s, s2), true)
	}

	if set.Subset(s2, s) {
		t.Errorf("Subset(%v, %v) = %t, debería ser %t", s2, s, set.Subset(s2, s), false)
	}
}
