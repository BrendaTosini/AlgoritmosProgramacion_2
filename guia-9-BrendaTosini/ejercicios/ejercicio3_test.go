package ejercicios

import (
	"fmt"
	"testing"
)

func contains(seleccionados map[Objeto]float32, objeto string, valor float32) bool {
	for o, v := range seleccionados {
		if o.Nombre == objeto && v == valor {
			return true
		}
	}
	return false
}
func TestEjercicio3(t *testing.T) {
	objetos := []Objeto{
		{"o1", 6, 20},
		{"o2", 3, 18},
		{"o3", 2, 14},
		{"o4", 5, 25},
	}
	capacidadMochila := 7
	salida := Ejercicio3(objetos, capacidadMochila)
	fmt.Println("Objetos: ", objetos)
	fmt.Println("Objetos seleccionados: ", salida)
	if len(salida) != 3 {
		t.Error("Se esperaban 3 objetos seleccionados")
	}

	if !contains(salida, "o2", 1) {
		t.Error("Se esperaba el objeto o1 con valor 1")
	}

	if !contains(salida, "o3", 1) {
		t.Error("Se esperaba el objeto o2 con valor 1")
	}

	if !contains(salida, "o4", 0.4) {
		t.Error("Se esperaba el objeto o3 con valor 0.4")
	}
}
