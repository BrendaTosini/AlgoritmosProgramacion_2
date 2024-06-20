package tests

import (
	"guia6/dictionary"
	"testing"
)

func TestNewContains(t *testing.T) {
	d := dictionary.NewDictionary[string, int]()
	d.Put("Leo", 55)
	d.Put("lucas", 38)
	if !d.Contains("Leo") {
		t.Errorf("Contains(Leo) = %t, debería ser %t", d.Contains("Leo"), true)
	}
	if !d.Contains("lucas") {
		t.Errorf("Contains(lucas) = %t, debería ser %t", d.Contains("lucas"), true)
	}
	if d.Contains("Fede") {
		t.Errorf("Contains(Fede) = %t, debería ser %t", d.Contains("Fede"), false)
	}
}