package monticulo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrearMonticuloDeMaxVacio(t *testing.T) {
	m := NewMonticuloMax[int]()
	assert.Equal(t, 0, m.Len())
}

func TestCrearMonticuloInsertarYExtraer(t *testing.T) {

	/*
		 Gracias a visualgo.net/en/heap
		 por la ayuda para preparar este caso de prueba.

		 Insertando los siguientes elementos en orden:
		 44, 29, 58, 2, 98, 11, 65, 3, 68, 99

		 El arbol resultante debería ser:
		      		  [99]
			       [98]            [65]
		   [58]        [68]   [11]  [44]
		 [2]  [3]   [29]

		 Como arreglo:
		 [99, 98, 65, 58, 68, 11, 44, 2, 3, 29]
	*/
	secuenciaDeInsercion := []int{44, 29, 58, 2, 98, 11, 65, 3, 68, 99}

	ordenEsperadoDespuesDeInsertar := [][]int{
		[]int{44},
		[]int{44, 29},
		[]int{58, 29, 44},
		[]int{58, 29, 44, 2},
		[]int{98, 58, 44, 2, 29},
		[]int{98, 58, 44, 2, 29, 11},
		[]int{98, 58, 65, 2, 29, 11, 44},
		[]int{98, 58, 65, 3, 29, 11, 44, 2},
		[]int{98, 68, 65, 58, 29, 11, 44, 2, 3},
		[]int{99, 98, 65, 58, 68, 11, 44, 2, 3, 29},
	}

	// Verificaciones iniciales
	m := NewMonticuloMax[int]()
	assert.Equal(t, 0, m.Len())

	// Verificaciones a medida que vamos insertando
	for i := 0; i < len(secuenciaDeInsercion); i++ {
		m.Insertar(secuenciaDeInsercion[i])
		assert.Equal(t, ordenEsperadoDespuesDeInsertar[i], m.elementos)
	}

	ordenEsperadoDespuesDeEliminar := [][]int{
		[]int{98, 68, 65, 58, 29, 11, 44, 2, 3},
		[]int{68, 58, 65, 3, 29, 11, 44, 2},
		[]int{65, 58, 44, 3, 29, 11, 2},
		[]int{58, 29, 44, 3, 2, 11},
		[]int{44, 29, 11, 3, 2},
		[]int{29, 3, 11, 2},
		[]int{11, 3, 2},
		[]int{3, 2},
		[]int{2},
		[]int{},
	}

	for i := 0; i < m.Len(); i++ {
		m.BorrarMax()
		assert.Equal(t, ordenEsperadoDespuesDeEliminar[i], m.elementos)
	}
}
