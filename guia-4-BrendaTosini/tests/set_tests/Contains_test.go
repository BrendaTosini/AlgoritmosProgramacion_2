package tests

import (
	"guia4/set"
	"testing"
)

func TestContains(t *testing.T) {
	s := set.NewSet(1, 3, 4, 2)
	if !s.Contains(4) {
		t.Errorf("Contains(4) = %t, debería ser %t", s.Contains(4), true)
	}
	if !s.Contains(1) {
		t.Errorf("Contains(1) = %t, debería ser %t", s.Contains(1), true)
	}
	if s.Contains(5) {
		t.Errorf("Contains(5) = %t, debería ser %t", s.Contains(5), false)
	}
}
