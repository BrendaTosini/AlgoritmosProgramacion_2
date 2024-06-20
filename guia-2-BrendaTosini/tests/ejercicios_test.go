package tests

import (
	"guia2/ejercicios"
	"guia2/queue"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestInvertirCadena(t *testing.T) {
	salida := ejercicios.InvertirCadena("abcd")
	if "dcba" != salida {
		t.Error("InvertirCadena falla")
	}
}

func TestPalindromo(t *testing.T) {
	if true != ejercicios.Palindromo("1456541") ||
		true != ejercicios.Palindromo("145541") {
		t.Error("Palindromo falla")
	}
}

func TestBalanceada(t *testing.T) {
	if true != ejercicios.Balanceada("[()]{}{[()()]()}") ||
		false != ejercicios.Balanceada("[(])") {
		t.Error("Balanceada falla")
	}
}

func TestUnirColas(t *testing.T) {

	var q1 queue.Queue

	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)

	var q2 queue.Queue

	q2.Enqueue(5)
	q2.Enqueue(7)

	var esperado queue.Queue

	esperado.Enqueue(1)
	esperado.Enqueue(2)
	esperado.Enqueue(3)
	esperado.Enqueue(5)
	esperado.Enqueue(7)

	dado := ejercicios.UnirColas(q1, q2)

	if !cmp.Equal(dado, esperado, cmpopts.IgnoreFields(queue.Queue{}, "cola")) {
		t.Error("UnirColas falla")
	}

}
