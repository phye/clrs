package bst

import (
	"fmt"
)

func (bst *BST) levelOrder(root *Node) [][]interface{} {
	ret := make([][]interface{}, 0)
	if root == nil {
		return ret
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	queue = append(queue, nil)
	for {
		if len(queue) == 0 || queue[0] == nil {
			break
		}
		k := 0 // Index of last nil
		for j := 0; j < len(queue); j++ {
			if queue[j] != nil {
				k++
				if queue[j].Left != nil {
					queue = append(queue, queue[j].Left)
				}
				if queue[j].Right != nil {
					queue = append(queue, queue[j].Right)
				}
			} else {
				queue = append(queue, nil)
				break
			}
		}
		ret = append(ret, make([]interface{}, k))
		for j := 0; j < k; j++ {
			ret[len(ret)-1][j] = queue[j].Value
		}
		queue = queue[k+1:]
		fmt.Printf("Level Order: %v\n", ret[len(ret)-1])
	}
	return ret
}
