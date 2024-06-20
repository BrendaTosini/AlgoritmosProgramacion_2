package main

import (
	"fmt"
	"guia3/linkedlist"
)

func main() {
	var lista linkedlist.LinkedList[int]
	lista.Append(1)
	lista.Append(3)
	var lista2 linkedlist.LinkedList[int]
	lista2.Append(2)
	lista2.Append(4)
	listafinal := IntercalarNumeros(lista, lista2)
	fmt.Println(listafinal.String())
}

func IntercalarNumeros(lista, lista2 linkedlist.LinkedList[int]) (listafinal linkedlist.LinkedList[int]) {

	if lista.Size() == 0 && lista2.Size() == 0 {
		return listafinal
	}
	var listaGrande linkedlist.LinkedList[int]
	var listaChica linkedlist.LinkedList[int]

	if lista.Size() > lista2.Size() {
		listaGrande = lista
		listaChica = lista2
	} else {
		listaGrande = lista2
		listaChica = lista
	}

	for i := 0; i < listaGrande.Size(); i++ {

		v, _ := listaGrande.Get(i)
		listafinal.Append(v)
		if listaChica.Size() > i {
			v, _ = listaChica.Get(i)
			listafinal.Append(v)
		}
	}
	return listafinal
}
