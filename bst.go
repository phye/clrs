package main

import (
	"fmt"
)

type (
	TreeNode struct {
		Left  *TreeNode
		Right *TreeNode
		Val   int
	}

	BinarySearchTree struct {
		root *TreeNode
	}
)

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{nil}
}

func (bst *BinarySearchTree) Insert(Val int) {
	if bst.root == nil {
		bst.root = &TreeNode{nil, nil, Val}
		return
	}
	p := bst.root
	cur := bst.root
	for {
		p = cur
		if Val <= cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
		if cur == nil {
			if Val <= p.Val {
				p.Left = &TreeNode{nil, nil, Val}
			} else {
				p.Right = &TreeNode{nil, nil, Val}
			}
			return
		}
	}
}

func (bst *BinarySearchTree) InorderTraverse(tn *TreeNode) {
	if tn != nil {
		bst.InorderTraverse(tn.Left)
		fmt.Printf("%v ", tn.Val)
		bst.InorderTraverse(tn.Right)
	}
}

func (bst *BinarySearchTree) IterativeInorderTraverse(tn *TreeNode) {
	stack := []*TreeNode{}
	cur := tn
	for {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else if len(stack) > 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Printf("%v ", cur.Val)
			cur = cur.Right
		} else {
			return
		}
	}
}

func (bst *BinarySearchTree) PreorderTraverse(tn *TreeNode) {
	if tn != nil {
		fmt.Printf("%v ", tn.Val)
		bst.PreorderTraverse(tn.Left)
		bst.PreorderTraverse(tn.Right)
	}
}

func (bst *BinarySearchTree) IterativePreorderTraverse(tn *TreeNode) {
	stack := []*TreeNode{}
	cur := tn
	for {
		fmt.Printf("%v ", cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if len(stack) > 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			return
		}
	}
}

func (bst *BinarySearchTree) PostorderTraverse(tn *TreeNode) {
	if tn != nil {
		bst.PostorderTraverse(tn.Left)
		bst.PostorderTraverse(tn.Right)
		fmt.Printf("%v ", tn.Val)
	}
}

func (bst *BinarySearchTree) IterativePostorderTraverse(tn *TreeNode) {
	stack := []*TreeNode{tn}
	pstack := []*TreeNode{}
	for {
		if len(stack) == 0 {
			break
		}
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pstack = append(pstack, cur)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}
	for i := len(pstack) - 1; i >= 0; i-- {
		fmt.Printf("%v ", pstack[i].Val)
	}
}

func (bst *BinarySearchTree) IterativePostorderTraverse2(tn *TreeNode) {
	stack := []*TreeNode{tn}
	processed := map[*TreeNode]bool{}
	for {
		if len(stack) == 0 {
			return
		}
		cur := stack[len(stack)-1]
		if processed[cur] != true {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			processed[cur] = true
		} else {
			fmt.Printf("%v ", cur.Val)
			stack = stack[:len(stack)-1]
		}
	}
}

func (bst *BinarySearchTree) BuildFromPreInTraverse(pre []int, in []int) {
}

func (bst *BinarySearchTree) BuildFromPrePostTraverse(pre []int, post []int) {
}

func (bst *BinarySearchTree) BuildFromInPostTraverse(in []int, post []int) {
}

func DumpStack(stack []*TreeNode) {
	for _, tn := range stack {
		fmt.Printf("%v ", tn.Val)
	}
	fmt.Printf("\n")
}

func main() {
	A := []int{3, 7, 2, 5, 11, 9, 6, 20, 1}
	bst := NewBinarySearchTree()
	for _, v := range A {
		bst.Insert(v)
	}
	bst.InorderTraverse(bst.root)
	fmt.Printf("\n")
	bst.IterativeInorderTraverse(bst.root)
	fmt.Printf("\n")
	bst.PreorderTraverse(bst.root)
	fmt.Printf("\n")
	bst.IterativePreorderTraverse(bst.root)
	fmt.Printf("\n")
	bst.PostorderTraverse(bst.root)
	fmt.Printf("\n")
	bst.IterativePostorderTraverse2(bst.root)
	fmt.Printf("\n")
}
