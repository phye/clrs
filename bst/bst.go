package bst

import (
	"errors"
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

func (bst *BST) Minimum(c *Node) *Node {
	for {
		if c == nil || c.Left == nil {
			return c
		}
		c = c.Left
	}
}

func (bst *BST) Maximum(c *Node) *Node {
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
		return bst.Maximum(n.Left)
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
		return bst.Minimum(n.Right)
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
func (bst *BST) Delete(v interface{}) (*Node, error) {
	n := bst.Search(v)
	if n == nil {
		return nil, errors.New(fmt.Sprintf("%v does not exist in BST", v))
	}
	p := n.Parent
	var pc **Node
	if n == bst.Root {
		pc = &bst.Root
	} else {
		if p.Left == n {
			pc = &p.Left
		} else {
			pc = &p.Right
		}
	}

	if n.Left == nil && n.Right == nil {
		*pc = nil
	} else if n.Left == nil {
		*pc = n.Right
		n.Right.Parent = p
	} else if n.Right == nil {
		*pc = n.Left
		n.Left.Parent = p
	} else {
		pre := bst.Predessor(v)
		// Tricky case when predessor is left child. In this case, simply replace the node to be
		// deleted with its left child (which has no right child)
		if pre.Parent == n {
			n.Parent.Left = pre
			pre.Parent = n.Parent
			pre.Right = n.Right
		} else {
			// Replace node to be deleted with its predessor node

			// Firstly, handle the links in the original place of predessor node
			pre.Parent.Right = pre.Left
			if pre.Left != nil {
				pre.Left.Parent = pre.Parent
			}

			// Then, handle the links in the target place for the predessor node
			pre.Left = n.Left
			pre.Right = n.Right
			n.Left.Parent = pre
			n.Right.Parent = pre
			pre.Parent = p
			*pc = pre
		}
	}
	return n, nil
}
