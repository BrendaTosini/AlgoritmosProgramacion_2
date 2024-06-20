package tests

import (
	"guia6/ejercicios"
	"testing"
)

func TestInterseccion(t *testing.T) {
	s1:= []string{"A", "B", "C"}
	s2:= []string{"A", "D", "E"}
	lr := ejercicios.Interseccion(s1, s2)
	if lr.Size() != 1 {
		t.Errorf("Tamaño %d, debería ser %d", lr.Size(), 1)
	}
	if !lr.Contains("A"){
		t.Errorf("Contains(A) = %t, debería ser %t", lr.Contains("A"), true)
	}
}
