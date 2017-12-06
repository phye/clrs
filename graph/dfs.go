package graph

import (
	"fmt"
)

type (
	DFSAux struct {
		color int
		pi    *Node
		d     int // Time at which the node is discovered
		f     int // Time at which the node is fully explored
	}
	DFSStatus map[*Node]*DFSAux
)

var (
	time int = 0
)

func (g *Graph) DFS() DFSStatus {
	dfsStatus := make(DFSStatus, 0)
	time = 0
	for _, n := range g.nodes {
		dfsStatus[n] = &DFSAux{WHITE, nil, INF, INF}
	}
	for _, n := range g.nodes {
		if dfsStatus[n].color == WHITE {
			g.dfsVisit(n, dfsStatus)
		}
	}
	fmt.Printf("%s", dfsStatus)
	return dfsStatus
}

func (g *Graph) dfsVisit(u *Node, dfsStatus DFSStatus) {
	time++
	dfsStatus[u].d = time
	dfsStatus[u].color = GRAY
	for _, v := range g.Adj(u) {
		if dfsStatus[v].color == WHITE {
			dfsStatus[v].pi = u
			g.dfsVisit(v, dfsStatus)
		}
	}
	dfsStatus[u].color = BLACK
	time++
	dfsStatus[u].f = time
}

func (dfss DFSStatus) String() string {
	var ret string
	ret += fmt.Sprintf("--------------------DFS--------------------\n")
	ret += fmt.Sprintf("%8s %8s %8s %8s\n", "current", "parent", "start", "finish")
	for n, s := range dfss {
		if s.pi != nil {
			ret += fmt.Sprintf("%8v %8v %8v %8v\n", n.value, s.pi.value, s.d, s.f)
		} else {
			ret += fmt.Sprintf("%8v %8v %8v %8v\n", n.value, "nil", s.d, s.f)
		}
	}
	ret += fmt.Sprintf("-------------------------------------------\n")
	return ret
}
