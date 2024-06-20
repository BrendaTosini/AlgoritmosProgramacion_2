package ejercicios

import (
	"testing"
)

func TestDados(t *testing.T) {
	n, k, x := 3, 6, 7
	formas := Dados(n, k, x)
	if formas != 15 {
		t.Errorf("Dados(n, k, x) = %d; se esperaba 15", formas)
	}
}
