package tests

import (
	"guia6/dictionary"
	"testing"
)

func TestRemove(t *testing.T) {
	d := dictionary.NewDictionary[string, int]()
	d.Put("Leo", 55)
	d.Put("Lucas", 38)
	d.Remove("Leo")
	if d.Size() != 1 {
		t.Errorf("Tamaño %d, debería ser: %d", d.Size(), 1)
	}
}