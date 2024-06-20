package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
Tipos de datos:
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64
float32 float64
complex64 complex128

byte // Alias de uint8
rune // Alias de int32, representa una posición en código Unicode
uintptr // Tipo utilizado para guardar una dirección de puntero

Nota: Los tamaños de int, uint y uintptr varian segun el sistema (normalmente 32 bits en SO de 32 bits y 64 bits en caso de SO de 64 bits)
*/

func main() {
	var i int
	fmt.Printf("Tamaño en bytes del tipo int): %d\n", reflect.TypeOf(i).Size())
	fmt.Printf("Tamaño en bytes de la variable i: %d\n", unsafe.Sizeof(i))
}
