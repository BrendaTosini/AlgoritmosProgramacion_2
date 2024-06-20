package tests

import (
	"guia6/dictionary"
	"testing"
	"reflect"
	"sort"
)

func TestGetKeys(t *testing.T) {
	d := dictionary.NewDictionary[string, int]()
	d.Put("Leo", 55)
	d.Put("Lucas", 38)
	d.Put("Fede", 24)
	esperado := []string{"Fede", "Leo", "Lucas"}
	obtenido := d.GetKeys()
	sort.Strings(obtenido)
	if !reflect.DeepEqual(esperado, obtenido) {
		t.Errorf("Las claves son %v, pero deberían ser %v", obtenido, esperado)
	}
}


