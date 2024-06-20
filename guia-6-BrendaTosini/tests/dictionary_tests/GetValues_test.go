package tests

import (
	"guia6/dictionary"
	"testing"
	"sort"
	"reflect"
)

func TestGetValues(t *testing.T) {
	d := dictionary.NewDictionary[string, int]()
	d.Put("Leo", 55)
	d.Put("Lucas", 38)
	d.Put("Fede", 24)
	esperado := []int{24, 38, 55}
	obtenido := d.GetValues()
	sort.Ints(obtenido)
	if !reflect.DeepEqual(esperado, obtenido) {
		t.Errorf("Los valores son %v, pero deberían ser %v", obtenido, esperado)
	}
}