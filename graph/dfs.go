package graph

import (
	"container/list"
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

// DFS will update nodes to be sorted in topological order
func (g *Graph) DFS() {
	dfsStatus := make(DFSStatus, 0)
	time = 0
	l := list.New()
	for _, n := range g.nodes {
		dfsStatus[n] = &DFSAux{WHITE, nil, INF, INF}
	}
	for _, n := range g.nodes {
		if dfsStatus[n].color == WHITE {
			g.dfsVisit(n, dfsStatus, l)
		}
	}
	fmt.Printf("%s", dfsStatus)

	g.nodes = g.restoreNodes(l)
	return
}

func (g *Graph) restoreNodes(l *list.List) []*Node {
	nodes := make([]*Node, l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		n := e.Value.(*Node)
		nodes[i] = n
		i++
	}
	return nodes
}

func (g *Graph) dfsVisit(u *Node, dfsStatus DFSStatus, l *list.List) {
	time++
	dfsStatus[u].d = time
	dfsStatus[u].color = GRAY
	for _, v := range g.Adj(u) {
		if dfsStatus[v].color == WHITE {
			dfsStatus[v].pi = u
			g.dfsVisit(v, dfsStatus, l)
		}
	}
	l.PushFront(u)
	dfsStatus[u].color = BLACK
	time++
	dfsStatus[u].f = time
}

// Stringer
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

// In place topo sort the graph, after sort, all nodes will be stored in decreasing u.f order
func (g *Graph) TopoSort() {
	g.DFS()
	return
}

// Divide graph into strongly connected components(subgraphs)
func (g *Graph) Decomposition() []*Graph {
	ret := []*Graph{}
	g.DFS()
	gt := g.Transpose()

	dfsStatus := make(DFSStatus, 0)
	time = 0
	l := list.New()
	for _, n := range gt.nodes {
		dfsStatus[n] = &DFSAux{WHITE, nil, INF, INF}
	}
	for _, n := range gt.nodes {
		if dfsStatus[n].color == WHITE {
			ng := NewGraph(g.directed)
			olh := l.Front()
			gt.dfsVisit(n, dfsStatus, l)
			for e := l.Front(); e != olh; e = e.Next() {
				n := e.Value.(*Node)
				ng.nodes = append(ng.nodes, n)
			}
			ret = append(ret, ng)
		}
	}
	return ret
}
