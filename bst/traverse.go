package bst

import ()

func (bst *BST) BFS(n *Node) []interface{} {
	m := make(map[*Node]bool)
	queue := make([]*Node, 0)
	queue = append(queue, n)
	m[n] = true
	ret := make([]interface{}, 0)
	for {
		if len(queue) == 0 {
			break
		}
		c := queue[0]
		queue = queue[1:]
		ret = append(ret, c.Value)
		if c.Left != nil {
			if _, ok := m[c.Left]; !ok {
				queue = append(queue, c.Left)
				m[c.Left] = true
			}
		}
		if c.Right != nil {
			if _, ok := m[c.Right]; !ok {
				queue = append(queue, c.Right)
				m[c.Right] = true
			}
		}
	}
	return ret
}

// PreOrder DFS
func (bst *BST) DFS(n *Node) []interface{} {
	ret := make([]interface{}, 0)
	bst.dfs(bst.Root, &ret)
	return ret
}

func (bst *BST) dfs(n *Node, ret *[]interface{}) {
	if n == nil {
		return
	}
	*ret = append(*ret, n.Value)
	bst.dfs(n.Left, ret)
	bst.dfs(n.Right, ret)
}

func (bst *BST) IterativePreOrder(n *Node) []interface{} {
	ret := make([]interface{}, 0)
	stack := make([]*Node, 0)
	stack = append(stack, n)
	for {
		if len(stack) == 0 {
			break
		}
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, c.Value)
		if c.Right != nil {
			stack = append(stack, c.Right)
		}
		if c.Left != nil {
			stack = append(stack, c.Left)
		}
	}
	return ret
}

func (bst *BST) InOrder(n *Node) []interface{} {
	ret := make([]interface{}, 0)
	bst.inorder(n, &ret)
	return ret
}

func (bst *BST) inorder(n *Node, ret *[]interface{}) {
	if n == nil {
		return
	}
	if n.Left != nil {
		bst.inorder(n.Left, ret)
	}
	*ret = append(*ret, n.Value)
	if n.Right != nil {
		bst.inorder(n.Right, ret)
	}
}

func (bst *BST) IterativeInOrder(n *Node) []interface{} {
	ret := make([]interface{}, 0)
	m := make(map[*Node]int) // 0: init 1: in stack, 2: visited
	stack := make([]*Node, 0)
	stack = append(stack, n)
	i := 0
	for {
		i++
		if len(stack) == 0 || i >= 100 {
			break
		}
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		m[c] = 0
		if c.Left == nil || m[c.Left] > 0 {
			ret = append(ret, c.Value)
			m[c] = 2
			if c.Right != nil && m[c.Right] == 0 {
				stack = append(stack, c.Right)
				m[c.Right] = 1
			}
		} else {
			if c.Right != nil && m[c.Right] == 0 {
				stack = append(stack, c.Right)
				m[c.Right] = 1
			}
			stack = append(stack, c)
			m[c] = 1
			if c.Left != nil && m[c.Left] == 0 {
				stack = append(stack, c.Left)
				m[c.Left] = 1
			}
		}
	}
	return ret
}

func (bst *BST) IterativeInOrder2(n *Node) []interface{} {
	ret := make([]interface{}, 0)
	stack := make([]*Node, 0)
	c := n
	for {
		if c != nil {
			stack = append(stack, c)
			c = c.Left
		} else if len(stack) > 0 {
			c = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, c.Value)
			c = c.Right
		} else {
			break
		}
	}
	return ret
}

func (bst *BST) IterativePostOrder(n *Node) []interface{} {
	ret := make([]interface{}, 0)
	stack := make([]*Node, 0)
	stack = append(stack, n)
	m := make(map[*Node]bool)
	for {
		if len(stack) == 0 {
			break
		}
		c := stack[len(stack)-1]
		if (c.Left == nil || m[c.Left] == true) && (c.Right == nil || m[c.Right] == true) {
			ret = append(ret, c.Value)
			stack = stack[:len(stack)-1]
			m[c] = true
		} else {
			if c.Right != nil && m[c.Right] == false {
				stack = append(stack, c.Right)
			}
			if c.Left != nil && m[c.Left] == false {
				stack = append(stack, c.Left)
			}
		}
	}
	return ret
}
