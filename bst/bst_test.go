package bst

import (
	"fmt"
	"testing"
)

func numCompare(a, b interface{}) bool {
	if a.(int) < b.(int) {
		return true
	} else {
		return false
	}
}

func TestInsert(t *testing.T) {
	bst := NewBST(numCompare)
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(8)

	if bst.Root.Value != 5 && bst.Root.Left.Value != 3 && bst.Root.Left.Left.Value != 2 && bst.Root.Left.Right.Value != 4 {
		t.Fatalf("Incorrect Left hand tree")
	}
	if bst.Root.Right.Value != 8 {
		t.Fatalf("Incorrect Right hand tree")
	}

	fmt.Printf("%s", bst)
}

func TestSearch(t *testing.T) {
	bst := NewBST(numCompare)
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(8)

	fmt.Printf("%s", bst)
	n := bst.Search(8)
	fmt.Printf("%v\n", n)
}
