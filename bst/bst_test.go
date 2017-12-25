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
	min := bst.Minimum(bst.Root)
	max := bst.Maximum(bst.Root)
	if min.Value != 2 {
		t.Fatalf("Tree Minimum Error")
	}
	if max.Value != 8 {
		t.Fatalf("Tree Maximum Error")
	}
	fmt.Printf("Minimum in tree is %v\n", min)
	fmt.Printf("Maximum in tree is %v\n", max)

	nbst := NewBST(numCompare)
	nmin := nbst.Minimum(nbst.Root)
	nmax := nbst.Maximum(nbst.Root)
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
	p5 := bst.Predessor(100)
	if p1.Value != 86 || p2.Value != 80 || p3.Value != 85 || p4 != nil || p5.Value != 93 {
		t.Fatalf("Error in BST Predessor")
	}
}

func TestDelete(t *testing.T) {
	bst := NewBST(numCompare)
	nums := []int{100, 80, 120, 50, 90, 110, 40, 70, 85, 93, 86, 92}
	for i := 0; i < len(nums); i++ {
		bst.Insert(nums[i])
	}

	if _, err := bst.Delete(77); err == nil {
		t.Fatalf("Delete non existing node should result in failure")
	}

	n1, _ := bst.Delete(90)
	fmt.Printf("After delete %v, bst is %s", n1.Value, bst)

	n2, _ := bst.Delete(100)
	fmt.Printf("After delete %v, bst is %s", n2.Value, bst)
}

func TestCLRSDelete(t *testing.T) {
	bst := NewBST(numCompare)
	nums := []int{100, 80, 120, 50, 90, 110, 40, 70, 85, 93, 86, 92}
	for i := 0; i < len(nums); i++ {
		bst.Insert(nums[i])
	}

	if _, err := bst.CLRSDelete(77); err == nil {
		t.Fatalf("Delete non existing node should result in failure")
	}

	n1, _ := bst.CLRSDelete(90)
	fmt.Printf("After delete %v, bst is %s", n1.Value, bst)

	n2, _ := bst.CLRSDelete(100)
	fmt.Printf("After delete %v, bst is %s", n2.Value, bst)
}

func TestBFS(t *testing.T) {
	bst := NewBST(numCompare)
	nums := []int{100, 80, 120, 50, 90, 110, 40, 70, 85, 93, 86, 92}
	for i := 0; i < len(nums); i++ {
		bst.Insert(nums[i])
	}
	bfs := bst.BFS(bst.Root)
	fmt.Printf("BFS: %v\n", bfs)
	expected := []int{100, 80, 120, 50, 90, 110, 40, 70, 85, 93, 86, 92}
	for i := 0; i < len(bfs); i++ {
		if bfs[i] != expected[i] {
			t.Fatalf("Unexpected BFS result!")
		}
	}
}

func TestDFS(t *testing.T) {
	bst := NewBST(numCompare)
	nums := []int{100, 80, 120, 50, 90, 110, 40, 70, 85, 93, 86, 92}
	for i := 0; i < len(nums); i++ {
		bst.Insert(nums[i])
	}
	dfs := bst.DFS(bst.Root)
	fmt.Printf("DFS: %v\n", dfs)
	expected := []int{100, 80, 50, 40, 70, 90, 85, 86, 93, 92, 120, 110}
	for i := 0; i < len(dfs); i++ {
		if dfs[i] != expected[i] {
			t.Fatalf("Unexpected DFS result!")
		}
	}
}

func TestPreOrder(t *testing.T) {
	bst := NewBST(numCompare)
	nums := []int{100, 80, 120, 50, 90, 110, 40, 70, 85, 93, 86, 92}
	for i := 0; i < len(nums); i++ {
		bst.Insert(nums[i])
	}
	res := bst.IterativePreOrder(bst.Root)
	expected := []int{100, 80, 50, 40, 70, 90, 85, 86, 93, 92, 120, 110}
	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Fatalf("Unexpected PreOrder result!")
		}
	}
}

func TestInOrder(t *testing.T) {
	bst := NewBST(numCompare)
	nums := []int{100, 80, 120, 50, 90, 110, 40, 70, 85, 93, 86, 92}
	for i := 0; i < len(nums); i++ {
		bst.Insert(nums[i])
	}
	res := bst.InOrder(bst.Root)
	expected := []int{40, 50, 70, 80, 85, 86, 90, 92, 93, 100, 110, 120}
	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Fatalf("Unexpected InOrder result!")
		}
	}

	res = bst.IterativeInOrder(bst.Root)
	fmt.Printf("InOrder1: %v \n", res)
	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Fatalf("Unexpected Iterative InOrder result!")
		}
	}

	res = bst.IterativeInOrder2(bst.Root)
	fmt.Printf("Inorder2: %v \n", res)
	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Fatalf("Unexpected Iterative InOrder result!")
		}
	}
}

func TestPostOrder(t *testing.T) {
	bst := NewBST(numCompare)
	nums := []int{100, 80, 120, 50, 90, 110, 40, 70, 85, 93, 86, 92}
	for i := 0; i < len(nums); i++ {
		bst.Insert(nums[i])
	}
	res := bst.IterativePostOrder(bst.Root)
	expected := []int{40, 70, 50, 86, 85, 92, 93, 90, 80, 110, 120, 100}
	fmt.Printf("PostOrder: %v \n", res)
	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Fatalf("Unexpected PostOrder result!")
		}
	}
}
