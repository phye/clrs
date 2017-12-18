package graph

import (
	"fmt"
	//"reflect"
)

type (
	DFSAux struct {
		color int
		f     int // Time at which the node is fully explored
		NodeAux
	}
	DFSStatus map[interface{}]*DFSAux
)

// Return all nodes in Depth-First Order
func DFS(g Graph) []interface{} {
	dfsStatus := make(DFSStatus, 0)
	time := 0
	ret := make([]interface{}, 0)
	for _, n := range g.Nodes() {
		dfsStatus[n] = &DFSAux{WHITE, INF, NodeAux{nil, nil, INF}}
	}
	for _, n := range g.Nodes() {
		if dfsStatus[n].color == WHITE {
			dfsVisit(g, n, dfsStatus, &ret, &time)
		}
	}
	return ret
}

func dfsVisit(g Graph, u interface{}, dfsStatus DFSStatus, res *[]interface{}, pt *int) {
	*pt++
	dfsStatus[u].d = *pt
	dfsStatus[u].color = GRAY
	adjs, _ := g.Adj(u)
	for _, v := range adjs {
		if dfsStatus[v].color == WHITE {
			dfsStatus[v].pi = u
			dfsVisit(g, v, dfsStatus, res, pt)
		}
	}
	*res = append([]interface{}{u}, *res...)
	dfsStatus[u].color = BLACK
	*pt++
	dfsStatus[u].f = *pt
}

// Stringer of DFSStatus
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
func TopoSort(g Graph) Graph {
	nodes := DFS(g)
	ng := SubGraph(g, nodes)
	return ng
}

// Divide graph into strongly connected components(subgraphs)
func Decomposition(g Graph) []Graph {
	ret := []Graph{}
	sg := TopoSort(g)
	gt := sg.Transpose()

	dfsStatus := make(DFSStatus, 0)
	time := 0
	for _, n := range gt.Nodes() {
		dfsStatus[n] = &DFSAux{WHITE, INF, NodeAux{nil, nil, INF}}
	}
	for _, n := range gt.Nodes() {
		if dfsStatus[n].color == WHITE {
			dret := make([]interface{}, 0)
			dfsVisit(gt, n, dfsStatus, &dret, &time)
			ng := SubGraph(gt, dret)
			ret = append(ret, ng)
		}
	}
	return ret
}
