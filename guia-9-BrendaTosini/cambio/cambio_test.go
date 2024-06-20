package cambio

import (
	"fmt"
	"testing"
)

func equal(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}

func TestCambio(t *testing.T) {
	cantidad := 1234
	cbio := map[int]int{
		1000: 1,
		200:  1,
		20:   1,
		10:   1,
		2:    2,
	}
	c := Cambiar(cantidad)
	if !equal(c, cbio) {
		t.Errorf("cambio(%d) = %v, want %v", cantidad, c, cbio)
	}
	fmt.Println(c)
}
