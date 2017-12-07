package graph

import (
	"container/list"
)

func (g *Graph) TopoSort() *list.List {
	l := list.New()
	dfsStatus := make(DFSStatus, 0)
	time = 0
	for _, n := range g.nodes {
		dfsStatus[n] = &DFSAux{WHITE, nil, INF, INF}
	}
	for _, n := range g.nodes {
		if dfsStatus[n].color == WHITE {
			g.topoSort(n, dfsStatus, l)
		}
	}
	return l
}

func (g *Graph) topoSort(u *Node, dfsStatus DFSStatus, l *list.List) {
	time++
	dfsStatus[u].d = time
	dfsStatus[u].color = GRAY
	for _, v := range g.Adj(u) {
		if dfsStatus[v].color == WHITE {
			dfsStatus[v].pi = u
			g.topoSort(v, dfsStatus, l)
		}
	}
	l.PushFront(u)
	dfsStatus[u].color = BLACK
	time++
	dfsStatus[u].f = time
}
