package ejercicios

import (
	"fmt"
	"testing"
)

func TestEjercicio2(t *testing.T) {
	tareas := []Tarea{
		{"t1", 34},
		{"t2", 13},
		{"t3", 16},
		{"t4", 7},
		{"t5", 9},
		{"t6", 19},
		{"t7", 10},
		{"t8", 21},
		{"t9", 42},
		{"t10", 14},
		{"t11", 25},
	}
	fmt.Println("Tareas: ", tareas)
	Ejercicio2(tareas)
	fmt.Println("Tareas ordenadas: ", tareas)
	for i := 0; i < len(tareas)-1; i++ {
		if tareas[i].Tiempo > tareas[i+1].Tiempo {
			t.Error("Las tareas no están ordenadas")
		}
	}
}
