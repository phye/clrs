package bst

// Lowest Common Ancestor

func (bst *BST) LowestCommonAncestor(p, q *Node) *Node {
	m := make(map[*Node]bool, 0)
	return bst.lcaAux(p, q, m)
}

func (bst *BST) lcaAux(p, q *Node, m map[*Node]bool) *Node {
	if p == nil || q == nil {
		return nil
	}
	if _, ok := m[q]; ok {
		return q
	}
	m[p] = true
	if p.Parent == nil {
		return bst.lcaAux(bst.Root, q.Parent, m)
	} else if q.Parent == nil {
		return bst.lcaAux(p.Parent, bst.Root, m)
	} else {
		return bst.lcaAux(p.Parent, q.Parent, m)
	}
}
