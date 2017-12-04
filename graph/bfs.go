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

func (g *Graph) BFS(s *Node) {
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
		if u.pi != nil {
			//fmt.Printf("%v %v %v\n", u.value, u.pi.value, u.depth)
		}
	}
}

func (g *Graph) Path(sn *Node, tn *Node) []*Node {
	g.BFS(sn)
	ret := g.path(sn, tn)
	for _, n := range ret {
		fmt.Printf("%v ", n.value)
	}
	fmt.Printf("\n")
	return ret
}

func (g *Graph) path(sn *Node, tn *Node) []*Node {
	if sn == tn {
		ret := []*Node{sn}
		return ret
	} else if tn.pi == nil {
		return []*Node{}
	} else {
		ret := g.path(sn, tn.pi)
		if len(ret) > 0 {
			ret = append(ret, tn)
		}
		return ret
	}
}
