package tests

import (
	"guia6/ejercicios"
	"testing"
)
func TestFrecuencia(t *testing.T) {
	d := ejercicios.Frecuencia("ahora")
	if d.Size() != 4 {
		t.Errorf("Tamaño %d, debería ser %d", d.Size(), 4)
	}
	if d.Get("a") != 2{
		t.Errorf("La cantidad que aparece la a deberia ser %d, pero es %d", 2, d.Get("a"))
	}
	if d.Get("r") != 1{
		t.Errorf("La cantidad que aparece la r deberia ser %d, pero es %d", 1, d.Get("r"))
	}
	if d.Get("o") != 1{
		t.Errorf("La cantidad que aparece la o deberia ser %d, pero es %d", 1, d.Get("o"))
	}
	if d.Get("h") != 1{
		t.Errorf("La cantidad que aparece la h deberia ser %d, pero es %d", 1, d.Get("h"))
	}
}
