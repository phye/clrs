package bst

import (
	"fmt"
)

type (
	Node struct {
		// Value must be comparable!
		Value  interface{}
		Left   *Node
		Right  *Node
		Parent *Node
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
	var pp *Node
	pc := &bst.Root
	for {
		if *pc == nil {
			*pc = &Node{v, nil, nil, pp}
			return
		}
		if bst.less(v, (*pc).Value) {
			pp = *pc
			pc = &((*pc).Left)
		} else {
			pp = *pc
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

func (bst *BST) Minimum() *Node {
	c := bst.Root
	for {
		if c == nil || c.Left == nil {
			return c
		}
		c = c.Left
	}
}

func (bst *BST) Maximum() *Node {
	c := bst.Root
	for {
		if c == nil || c.Right == nil {
			return c
		}
		c = c.Right
	}
}

func (bst *BST) Predessor(v interface{}) *Node {
	n := bst.Search(v)
	if n == nil {
		return nil
	}
	if n.Left != nil {
		var c *Node
		for c = n.Left; c.Right != nil; c = c.Right {
		}
		return c
	} else {
		var c *Node
		for c = n.Parent; c != nil && c.Left == n; c, n = c.Parent, n.Parent {
		}
		return c
	}
}

func (bst *BST) Successor(v interface{}) *Node {
	n := bst.Search(v)
	if n == nil {
		return nil
	}
	if n.Right != nil {
		var c *Node
		for c = n.Right; c.Left != nil; c = c.Left {
		}
		return c
	} else {
		var c *Node
		for c = n.Parent; c != nil && c.Right == n; c, n = c.Parent, n.Parent {
		}
		return c
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
