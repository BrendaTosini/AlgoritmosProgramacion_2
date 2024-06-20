package ejercicios

import (
	"guia2/queue"
	"guia2/stack"
)

// O(2n) = O(n)
func InvertirCadena(cadena string) string {

	var pila stack.Stack
	var invertida string

	for _, v := range cadena {
		pila.Push(string(v))
	}

	for i := 0; i < len(cadena); i++ {
		v, _ := pila.Pop()
		invertida += v.(string)
	}

	return invertida
}

// O(2n)= O(n)
func Palindromo(cadena string) bool {

	inversa := InvertirCadena(cadena)

	return cadena == inversa
}

// O(n)
func Balanceada(cadena string) bool {

	var pila stack.Stack
	balanceada := true

	i := 0

	for i < len(cadena) && balanceada {

		simbolo := string(cadena[i])

		if simbolo == "{" || simbolo == "[" || simbolo == "(" {

			pila.Push(simbolo)
		} else {

			a, _ := pila.Pop()
			simbolo2 := a.(string)

			switch simbolo {

			case ")":
				if simbolo2 == "(" {
					balanceada = true
				} else {
					balanceada = false
				}
			case "]":
				if simbolo2 == "[" {
					balanceada = true
				} else {
					balanceada = false
				}
			case "}":
				if simbolo2 == "{" {
					balanceada = true
				} else {
					balanceada = false
				}
			}
		}
		i++
	}
	return balanceada

}

// O(n)
func UnirColas(q1, q2 queue.Queue) queue.Queue {

	for !q2.IsEmpty() {

		valor, _ := q2.Dequeue()

		q1.Enqueue(valor)

	}
	return q1

}
