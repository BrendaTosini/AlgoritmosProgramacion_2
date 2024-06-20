package linkedlist_test

import (
	"guia4/linkedlist"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	l := linkedlist.NewLinkedList[int]()
	if l == nil {
		t.Errorf("NewLinkedList() = nil, want *LinkedList")
	}
}

func TestInsertAt(t *testing.T) {
	l := linkedlist.NewLinkedList[int]()
	l.InsertAt(1, l.Size())
	l.InsertAt(2, l.Size())
	l.InsertAt(3, l.Size())
	if l.Size() != 3 {
		t.Errorf("Error en tamaño de la lista: %d, debería ser %d", l.Size(), 2)
	}
	for i := 0; i < l.Size(); i++ {
		if l.Get(i) != i+1 {
			t.Errorf("Error en el elemento %d, debería ser %d", l.Get(i), i+1)
		}
	}
}

func TestRemoveAt(t *testing.T) {
	l := linkedlist.NewLinkedList[int]()
	l.InsertAt(1, 0)
	l.InsertAt(2, 0)
	l.InsertAt(3, 0)
	l.RemoveAt(0)
	if l.Size() != 2 {
		t.Errorf("Error en el tamaño de la lista: %d, debería ser %d", l.Size(), 2)
	}
	l.RemoveAt(l.Size() - 1)
	if l.Size() != 1 {
		t.Errorf("Error en el tamaño de la lista: %d, debería ser %d", l.Size(), 1)
	}
	l.RemoveAt(0)
	if !l.IsEmpty() {
		t.Errorf("La lista debería estar vacía")
	}
	if l.Size() != 0 {
		t.Errorf("Error en el tamaño de la lista: %d, debería ser %d", l.Size(), 0)
	}
}

func TestGet(t *testing.T) {
	l := linkedlist.NewLinkedList[int]()
	l.InsertAt(1, 0)
	l.InsertAt(2, 0)
	l.InsertAt(3, 0)
	l.InsertAt(4, 0)
	if l.Get(0) != 4 {
		t.Errorf("Get() = %d, want %d", l.Get(0), 4)
	}
	if l.Get(l.Size()-1) != 1 {
		t.Errorf("Get() = %d, want %d", l.Get(0), 1)
	}

}
