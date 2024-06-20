package main

import "fmt"

func main() {

	i := 20
	j := 50
	swap(&i, &j)
	fmt.Println(i, j)

}
func swap(px, py *int) {
	aux := *px
	*px = *py
	*py = aux

}
