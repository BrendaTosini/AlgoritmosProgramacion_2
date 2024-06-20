package mapadebits

import (
	"errors"
)

type MapaDeBits struct {
	posiciones uint32
}

// Crea un mapa y lo inicializa con el entero 0
func NewMapaDeBits() *MapaDeBits {
	return &MapaDeBits{posiciones: 0}
}

// Enciende el bit de una posición indicada
func (m *MapaDeBits) Encender(pos uint8) error {
	if valida(pos) {
		m.posiciones |= 1 << pos
		return nil
	}

	return errors.New("posición inválida")
}

// Apaga el bit de una posición indicada
func (m *MapaDeBits) Apagar(pos uint8) error {
	if valida(pos) {
		m.posiciones &= ^(1 << pos)
		return nil
	}

	return errors.New("posición inválida")
}

// Chequea si el bit de una posición indicada está encendido
func (m *MapaDeBits) EstaEncendido(pos uint8) (bool, error) {
	if valida(pos) {
		return m.posiciones&(1<<pos) != 0, nil
	}

	return false, errors.New("posición inválida")
}

func valida(pos uint8) bool {
	return pos < 32
}

func (m *MapaDeBits) GetMapa() uint32 {
	return m.posiciones
}
