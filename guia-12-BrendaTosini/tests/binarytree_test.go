package tests

import (
	"guia12/binarytree"
	"testing"
)

func TestTamañoYAlturaDeSoloRaiz(t *testing.T) {
	btree := binarytree.NewBinaryTree("-")
	if btree.Size() != 1 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 1, btree.Size())
	}
	if btree.Height() != 0 {
		t.Errorf("Error la altura deberia dar %v, pero dio %v", 0, btree.Height())
	}
}

func TestTamañoYAlturaDeSoloRaizConHijoIzquierdo(t *testing.T) {
	btree1 := binarytree.NewBinaryTree("+")
	btree2 := binarytree.NewBinaryTree("-")
	btree2.InsertLeft(btree1)
	if btree2.Size() != 2 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 2, btree2.Size())
	}
	if btree2.Height() != 1 {
		t.Errorf("Error la altura deberia dar %v, pero dio %v", 1, btree2.Height())
	}
}

func TestTamañoYAlturaDeSoloRaizConHijoDerecho(t *testing.T) {
	btree1 := binarytree.NewBinaryTree("+")
	btree2 := binarytree.NewBinaryTree("-")
	btree2.InsertRight(btree1)
	if btree2.Size() != 2 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 2, btree2.Size())
	}
	if btree2.Height() != 1 {
		t.Errorf("Error la altura deberia dar %v, pero dio %v", 1, btree2.Height())
	}
}

func TestTamañoYAlturaDeRaizYAmbosHijos(t *testing.T) {
	btree1 := binarytree.NewBinaryTree("+")
	btree2 := binarytree.NewBinaryTree("*")
	btree3 := binarytree.NewBinaryTree("-")
	btree3.InsertLeft(btree1)
	btree3.InsertRight(btree2)
	if btree3.Size() != 3 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 3, btree3.Size())
	}
	if btree3.Height() != 1 {
		t.Errorf("Error la altura deberia dar %v, pero dio %v", 1, btree3.Height())
	}
}
