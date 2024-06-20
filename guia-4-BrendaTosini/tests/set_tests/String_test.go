package tests

import (
	"guia4/set"
	"testing"
)

func TestString(t *testing.T) {
	// Creo un conjunto vacío
	conjunto := set.NewSet("Hola", "Mundo", "!")
	// Como siempre agrega elementos al principio, el orden en que los imprime
	// es al revés
	if conjunto.String() != "Conjunto: {! Mundo Hola}" {
		t.Error("El conjunto no tiene los elementos correctos")
	}
}
