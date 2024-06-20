package monticulo

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Monticulo es una estructura de datos que permite almacenar elementos
// de forma ordenada. En este caso, se implementa un montículo máximo.
type MonticuloMax[T constraints.Ordered] struct {
	// Elementos del montículo
	elementos []T
}

// NewMonticulo crea un nuevo montículo vacío.
func NewMonticuloMax[T constraints.Ordered]() *MonticuloMax[T] {
	return &MonticuloMax[T]{}
}

// Len devuelve el número de elementos del montículo.
func (m *MonticuloMax[T]) Len() int {
	return len(m.elementos)
}

// Insertar inserta un nuevo elemento en el montículo.
func (m *MonticuloMax[T]) Insertar(elemento T) {
	m.elementos = append(m.elementos, elemento)
	m.upHeap(len(m.elementos) - 1)
}

// upHeap reordena el montículo hacia arriba.
// Costo: O(log n)
func (m *MonticuloMax[T]) upHeap(i int) {
	for i > 0 {
		padre := (i - 1) / 2
		if m.elementos[i] <= m.elementos[padre] {
			break
		}
		m.elementos[i], m.elementos[padre] = m.elementos[padre], m.elementos[i]
		i = padre
	}
}

// BorrarMax extrae el elemento máximo del montículo.
func (m *MonticuloMax[T]) BorrarMax() (T, error) {
	var elemento T
	if m.Len() == 0 {
		return elemento, fmt.Errorf("Montículo vacío")
	}
	elemento = m.elementos[0]
	m.elementos[0] = m.elementos[m.Len()-1]
	m.elementos = m.elementos[:m.Len()-1]
	m.downHeap(0)
	return elemento, nil
}

// downHeap reordena el montículo hacia abajo.
// Costo: O(log n)
func (m *MonticuloMax[T]) downHeap(i int) {
	for {
		izq := 2*i + 1
		der := 2*i + 2
		if izq >= m.Len() {
			break
		}
		max := izq
		if der < m.Len() && m.elementos[der] > m.elementos[izq] {
			max = der
		}
		if m.elementos[i] >= m.elementos[max] {
			break
		}
		m.elementos[i], m.elementos[max] = m.elementos[max], m.elementos[i]
		i = max
	}
}
