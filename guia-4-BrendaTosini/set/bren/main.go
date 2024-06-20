package main

import (
	"fmt"
	"guia4/set"
)

func main() {

	var conjunto set.Set[int]
	conjunto.Add(1)
	fmt.Println(&conjunto)

}
