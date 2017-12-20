package bst

import (
	"fmt"
)

type (
	Node struct {
		// Value must be comparable!
		Value interface{}
		Left  *Node
		Right *Node
	}

	BST struct {
		Root *Node
		less Comparator
	}

	Comparator func(a, b interface{}) bool
)

func NewBST(less Comparator) *BST {
	return &BST{nil, less}
}

func (bst *BST) Insert(v interface{}) {
	// Note that we can not simply use c := bst.Root, as the assignment to c blow will not be
	// shown to bst.Root
	pc := &bst.Root
	for {
		if *pc == nil {
			*pc = &Node{v, nil, nil}
			return
		}
		if bst.less(v, (*pc).Value) {
			pc = &((*pc).Left)
		} else {
			pc = &((*pc).Right)
		}
	}
}

func (bst *BST) Search(v interface{}) *Node {
	c := bst.Root
	for {
		if c == nil || c.Value == v {
			return c
		}
		if bst.less(v, c.Value) {
			c = c.Left
		} else {
			c = c.Right
		}
	}
}

func (bst *BST) String() string {
	ret := "\n"
	bst.genString(bst.Root, &ret)
	ret += "\n"
	return ret
}

func (bst *BST) genString(c *Node, str *string) {
	if c == nil {
		return
	}
	*str += fmt.Sprintf("%d ", c.Value)
	bst.genString(c.Left, str)
	bst.genString(c.Right, str)
}

// Delete a node from BST should be the most difficult task
func (bst *BST) Delete(v interface{}) error {
	return nil
}
