package bst_test

import (
	"log"
	"testing"

	"github.com/alwindoss/goal/bst"
)

func TestBinarySearchTreeWhenTheTreeIsNewlyCreated(t *testing.T) {
	bst := bst.NewBST()
	bst.Insert(10)
	if bst.Root == nil {
		log.Println("ROOOTTTTT")
	}
	if bst.Root.Left != nil {
		t.Errorf("expected the Root node's Left node to be nil")
	}
	if bst.Root.Right != nil {
		t.Errorf("expected the Root node's Right node to be nil")
	}
	if bst.Root.Value != 10 {
		t.Errorf("expected the Root node's value to be %d but was %d", 10, bst.Root.Value)
	}
}

func TestBinarySearchTreeWhenTheMultipleValuesAreInserted(t *testing.T) {
	bst := bst.NewBST()
	bst.Insert(10)
	bst.Insert(8)
	bst.Insert(11)
	if bst.Root.Value != 10 {
		t.Errorf("expected root value to be %d but found %d", 10, bst.Root.Value)
	}
	if bst.Root.Left.Value != 8 {
		t.Errorf("expected root value to be %d but found %d", 8, bst.Root.Left.Value)
	}
	if bst.Root.Right.Value != 11 {
		t.Errorf("expected root value to be %d but found %d", 11, bst.Root.Right.Value)
	}
}
