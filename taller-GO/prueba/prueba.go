// Write a function that takes a number and returns a list of its digits.
package prueba

func lista(numero int) []int {

	arr := make([]int, 3)
	numeroString := string(numero)

	for i, numero := range numeroString {

		arr[i] = int(numero)
	}

	return arr

}

func main() {

	lista(123)
}
