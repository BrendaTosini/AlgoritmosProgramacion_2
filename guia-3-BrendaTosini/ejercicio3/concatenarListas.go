package main

import (
	"fmt"
	"guia3/linkedlist"
)

func main() {

	var l linkedlist.LinkedList[any]
	l.Append(1)
	l.Append(2)
	l.Append(3)

	var l2 linkedlist.LinkedList[any]
	l2.Append(4)
	l2.Append(5)

	fmt.Println(Concatenar(l, l2))

}

func Concatenar(lista, lista2 linkedlist.LinkedList[any]) *linkedlist.LinkedList[any] {

	for i := 0; i < lista2.Size(); i++ {
		v, _ := lista2.Get(i)
		lista.Append(v)
	}
	return &lista

}
