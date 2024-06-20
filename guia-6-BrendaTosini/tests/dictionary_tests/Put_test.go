package tests

import (
	"guia6/dictionary"
	"testing"
)

func TestPut(t *testing.T) {
	d := dictionary.NewDictionary[string, int]()
	d.Put("Leo", 55)
	d.Put("Lucas", 38)
	if d.Size() != 2 {
		t.Errorf("Tamaño %d, debería ser %d", d.Size(), 2)
	}
}

func TestPutEqualsElements(t *testing.T) {
	d := dictionary.NewDictionary[string, int]()
	d.Put("Leo", 55)
	d.Put("Leo", 38)
	if d.Size() != 1 {
		t.Errorf("Tamaño %d, debería ser %d", d.Size(), 1)
	}
	if d.Get("Leo") != 38 {
		t.Errorf("La edad de Leo es de %d, pero debería ser de %d", d.Get("Leo"), 38)
	}
}

