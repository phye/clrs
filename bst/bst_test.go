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

func TestMinMax(t *testing.T) {
	bst := NewBST(numCompare)
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(8)

	fmt.Printf("%s", bst)
	min := bst.Minimum()
	max := bst.Maximum()
	if min.Value != 2 {
		t.Fatalf("Tree Minimum Error")
	}
	if max.Value != 8 {
		t.Fatalf("Tree Maximum Error")
	}
	fmt.Printf("Minimum in tree is %v\n", min)
	fmt.Printf("Maximum in tree is %v\n", max)

	nbst := NewBST(numCompare)
	nmin := nbst.Minimum()
	nmax := nbst.Maximum()
	if nmin != nil {
		t.Fatalf("Empty Tree Minimum Error")
	}
	if nmax != nil {
		t.Fatalf("Empty Tree Maximum Error")
	}
	fmt.Printf("Minimum in empty tree is %v\n", nmin)
	fmt.Printf("Maximum in empty tree is %v\n", nmax)
}

func TestSuccessor(t *testing.T) {
	bst := NewBST(numCompare)
	nums := []int{100, 80, 120, 50, 90, 110, 40, 70, 85, 93, 86, 92}
	for i := 0; i < len(nums); i++ {
		bst.Insert(nums[i])
	}

	s1 := bst.Successor(80)
	s2 := bst.Successor(85)
	s3 := bst.Successor(93)
	s4 := bst.Successor(120)
	if s1.Value != 85 || s2.Value != 86 || s3.Value != 100 || s4 != nil {
		t.Fatalf("Error in BST successor")
	}
}

func TestPredessor(t *testing.T) {
	bst := NewBST(numCompare)
	nums := []int{100, 80, 120, 50, 90, 110, 40, 70, 85, 93, 86, 92}
	for i := 0; i < len(nums); i++ {
		bst.Insert(nums[i])
	}

	p1 := bst.Predessor(90)
	p2 := bst.Predessor(85)
	p3 := bst.Predessor(86)
	p4 := bst.Predessor(40)
	if p1.Value != 86 || p2.Value != 80 || p3.Value != 85 || p4 != nil {
		t.Fatalf("Error in BST Predessor")
	}
}
