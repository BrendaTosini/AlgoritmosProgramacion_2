package tests

import (
	"guia6/dictionary"
	"testing"
)

func TestGet(t *testing.T) {
	d := dictionary.NewDictionary[string, int]()
	d.Put("Lucas", 38)
	if d.Get("Lucas") != 38 {
		t.Errorf("La edad de Lucas es de %d, pero debería ser de %d", d.Get("Leo"), 38)
	}
}
