package tests

import (
	"fmt"
	"guia5/mapadebits"
	"testing"
)

func TestEncenderTodosDeAUno(t *testing.T) {
	m := mapadebits.NewMapaDeBits()

	for i := uint8(0); i < 32; i++ {
		m.Encender(i)
	}

	var max uint32 = 0
	max--

	fmt.Printf("Aca %032b\n", m.GetMapa())
	fmt.Printf("Aca %032b\n", max)

	if m.GetMapa() != max {
		t.Error("Error en Encender todos")
	}
}

func TestQueSigaEncendido(t *testing.T) {
	m := mapadebits.NewMapaDeBits()

	m.Encender(1)
	m.Encender(1)

	aux, _ := m.EstaEncendido(1)
	if !aux {
		t.Error("Error en Encender multiple")
	}
}

func TestPosicionExcedida(t *testing.T) {
	m := mapadebits.NewMapaDeBits()

	err := m.Encender(32)

	if err == nil {
		t.Error("Error en exceso de posiciones")
	}
}
