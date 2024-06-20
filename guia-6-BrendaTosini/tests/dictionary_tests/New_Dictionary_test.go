package tests

import (
	"guia6/dictionary"
	"testing"
)

func TestNewDictionary(t *testing.T) {
	d := dictionary.NewDictionary[string, int]()
	if &d == nil {
		t.Errorf("NewDictionary() = nil, want *Dictionary")
	}
}