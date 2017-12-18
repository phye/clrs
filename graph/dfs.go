package graph

import (
	"container/list"
	"fmt"
)

type (
	DFSAux struct {
		color int
		NodeAux
		f int // Time at which the node is fully explored
	}
	DFSStatus map[interface{}]*DFSAux
)

// DFS will update nodes to be sorted in topological order
func DFS(g Graph) []interface{} {
	dfsStatus := make(DFSStatus, 0)
	time := 0
	l := list.New()
	for _, n := range g.Nodes() {
		dfsStatus[n] = &DFSAux{WHITE, NodeAux{nil, nil, INF}, INF}
	}
	for _, n := range g.Nodes() {
		if dfsStatus[n].color == WHITE {
			dfsVisit(g, n, dfsStatus, l, &time)
		}
	}
	fmt.Printf("%s", dfsStatus)

	nodes := restoreNodes(l)
	return nodes
}

func dfsVisit(g Graph, u interface{}, dfsStatus DFSStatus, l *list.List, pt *int) {
	*pt++
	dfsStatus[u].d = *pt
	dfsStatus[u].color = GRAY
	adjs, _ := g.Adj(u)
	for _, v := range adjs {
		if dfsStatus[v].color == WHITE {
			dfsStatus[v].pi = u
			dfsVisit(g, v, dfsStatus, l, pt)
		}
	}
	l.PushFront(u)
	dfsStatus[u].color = BLACK
	*pt++
	dfsStatus[u].f = *pt
}

func restoreNodes(l *list.List) []interface{} {
	nodes := make([]interface{}, l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		n := e.Value
		nodes[i] = n
		i++
	}
	return nodes
}

// Stringer
func (dfss DFSStatus) String() string {
	var ret string
	ret += fmt.Sprintf("--------------------DFS--------------------\n")
	ret += fmt.Sprintf("%8s %8s %8s %8s\n", "current", "parent", "start", "finish")
	for n, s := range dfss {
		if s.pi != nil {
			ret += fmt.Sprintf("%8v %8v %8v %8v\n", n, s.pi, s.d, s.f)
		} else {
			ret += fmt.Sprintf("%8v %8v %8v %8v\n", n, "nil", s.d, s.f)
		}
	}
	ret += fmt.Sprintf("-------------------------------------------\n")
	return ret
}

// In place topo sort the graph, after sort, all nodes will be stored in decreasing u.f order
func TopoSort(g Graph) []interface{} {
	nodes := DFS(g)
	return nodes
}

// Divide graph into strongly connected components(subgraphs)
func Decomposition(g Graph) []*Graph {
	ret := []*Graph{}
	/*
		DFS(g)
		gt := g.Transpose()

		dfsStatus := make(DFSStatus, 0)
		time = 0
		l := list.New()
		for _, n := range gt.Nodes() {
			dfsStatus[n] = &DFSAux{WHITE, NodeAux{nil, nil, INF}, INF}
		}
		for _, n := range gt.Nodes() {
			if dfsStatus[n].color == WHITE {
				ng := NewGraph(g.directed)
				olh := l.Front()
				gt.dfsVisit(n, dfsStatus, l)
				for e := l.Front(); e != olh; e = e.Next() {
					n := e.Value
					ng.nodes = append(ng.nodes, n)
				}
				ret = append(ret, ng)
			}
		}
	*/
	return ret
}
