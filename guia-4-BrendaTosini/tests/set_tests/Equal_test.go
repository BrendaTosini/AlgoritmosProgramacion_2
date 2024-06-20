package tests

import (
	"guia4/set"
	"testing"
)

func TestEqual(t *testing.T) {
	s := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4, 5, 6)
	if set.Equal(s, s2) {
		t.Errorf("Equal(%v, %v) = %t, debería ser %t", s, s2, set.Equal(s, s2), false)
	}
	s.Remove(1)
	s.Remove(2)
	s.Add(5)
	s.Add(6)
	if !set.Equal(s, s2) {
		t.Errorf("Equal(%v, %v) = %t, debería ser %t", s, s2, set.Equal(s, s2), true)
	}

	if !set.Equal(s, s) {
		t.Errorf("Un conjunto debería ser igual a sí mismo")
	}
	s.Remove(3)
	s.Remove(4)
	s.Remove(5)
	s.Remove(6)
	s3 := set.NewSet[int]()
	if !set.Equal(s, s3) {
		t.Errorf("s debería ser igual al conjunto vacío")
	}

}
