package ejercicios

// Escribir un método que calcule recursivamente cuántos
// elementos hay en una pila y no altere el contenido de
// la misma. La pila sólo tiene los métodos Push, Pop y isEmpty.
func CantidadDeElementos(pila Stack) int {

	var elementos int

	if !pila.IsEmpty() {

		elem, _ := pila.Pop()

		elementos = 1 + CantidadDeElementos(pila)

		pila.Push(elem)

	}
	return elementos
}
