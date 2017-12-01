package graph

import (
	"fmt"
)

const (
	WHITE int = iota
	GRAY
	BLACK

	INF int = 0x7fffffff
)

func BFS(g *Graph, s *Node) {
	for _, n := range g.nodes {
		if n != s {
			n.color = WHITE
			n.depth = INF
			n.pi = nil
		}
	}
	s.color = GRAY
	s.depth = 0
	s.pi = nil
	queue := make([]*Node, 0)
	queue = append(queue, s)
	for {
		if len(queue) == 0 {
			break
		}
		u := queue[0]
		queue = queue[1:]
		for _, v := range g.Adj(u) {
			if v.color == WHITE {
				v.color = GRAY
				v.depth = u.depth + 1
				v.pi = u
				queue = append(queue, v)
			}
		}
		u.color = BLACK
		fmt.Printf("%v\n", queue)
		fmt.Printf("%v %v %v\n", u.value, u.color, u.depth)
	}
}
