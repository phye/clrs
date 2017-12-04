package graph

import (
	"fmt"
)

type (
	// The following three fields are necessar for BFS
	BFSAux struct {
		color int
		depth int
		pi    *Node
	}

	BFSStatus map[*Node]*BFSAux
)

const (
	WHITE int = iota
	GRAY
	BLACK

	INF int = 0x7fffffff
)

func (g *Graph) BFS(s *Node) BFSStatus {
	bfsStatus := make(BFSStatus, 0)
	for _, n := range g.nodes {
		bfsStatus[n] = &BFSAux{WHITE, INF, nil}
	}
	bfsStatus[s].color = GRAY
	bfsStatus[s].depth = 0
	bfsStatus[s].pi = nil
	queue := make([]*Node, 0)
	queue = append(queue, s)
	for {
		if len(queue) == 0 {
			break
		}
		u := queue[0]
		queue = queue[1:]
		for _, v := range g.Adj(u) {
			if bfsStatus[v].color == WHITE {
				bfsStatus[v].color = GRAY
				bfsStatus[v].depth = bfsStatus[u].depth + 1
				bfsStatus[v].pi = u
				queue = append(queue, v)
			}
		}
		bfsStatus[u].color = BLACK
	}
	fmt.Printf("%s", bfsStatus)
	return bfsStatus
}

func (g *Graph) Path(sn *Node, tn *Node) []*Node {
	bfsStatus := g.BFS(sn)
	ret := g.path(sn, tn, bfsStatus)
	for _, n := range ret {
		fmt.Printf("%v ", n.value)
	}
	fmt.Printf("\n")
	return ret
}

func (g *Graph) path(sn *Node, tn *Node, bfsStatus BFSStatus) []*Node {
	if sn == tn {
		ret := []*Node{sn}
		return ret
	} else if bfsStatus[tn].pi == nil {
		return []*Node{}
	} else {
		ret := g.path(sn, bfsStatus[tn].pi, bfsStatus)
		if len(ret) > 0 {
			ret = append(ret, tn)
		}
		return ret
	}
}

func (bfss BFSStatus) String() string {
	var ret string
	ret += fmt.Sprintf("------------BFS------------\n")
	ret += fmt.Sprintf("%8s %8s %8s\n", "current", "parent", "depth")
	for n, s := range bfss {
		if s.pi != nil {
			ret += fmt.Sprintf("%8v %8v %8v\n", n.value, s.pi.value, s.depth)
		} else {
			ret += fmt.Sprintf("%8v %8v %8v\n", n.value, "nil", s.depth)
		}
	}
	ret += fmt.Sprintf("---------------------------\n")
	return ret
}
