package graph

import (
	"github.com/phye/clrs/ds"
)

type (
	// Single-Source Shortest Paths
	// Note the similarity between BFS and SSS problem:
	// BFS solves the problem of minimum number of edges, while SSS solves the problem of
	// minimum weight of edges
	SSSAux struct {
		isolated bool
		NodeAux
	}

	SSSAuxMap map[interface{}]*SSSAux
)

func initSingleSourceAuxMap(g Graph, s interface{}) SSSAuxMap {
	sam := make(SSSAuxMap)
	for _, n := range g.Nodes() {
		sam[n] = &SSSAux{true, NodeAux{n, nil, INF}}
	}
	sam[s].d = 0
	return sam
}

func relax(g Graph, sam SSSAuxMap, sv, ev interface{}) {
	wt, _ := g.Weight(sv, ev)
	if sam[ev].d > sam[sv].d+wt {
		sam[ev].d = sam[sv].d + wt
		sam[ev].pi = sv
	}
}

// Return Single-Source Shortest Paths graph started from s via BellmanFord Algorithm O(VE)
// Advantage: Versatile
// Shortage: Slow
func BellmanFord(g Graph, s interface{}) Graph {
	sam := initSingleSourceAuxMap(g, s)
	for i := 1; i < len(g.Nodes())-1; i++ {
		for _, n1 := range g.Nodes() {
			for _, n2 := range g.Nodes() {
				if n1 != n2 {
					if _, err := g.Weight(n1, n2); err == nil {
						relax(g, sam, n1, n2)
					}
				}
			}
		}
	}

	for _, n1 := range g.Nodes() {
		for _, n2 := range g.Nodes() {
			if n1 != n2 {
				if wt, err := g.Weight(n1, n2); err == nil {
					if sam[n2].d > sam[n1].d+wt {
						return nil
					}
				}
			}
		}
	}

	ng := NewGraph(g.Directed(), g.GraphType())
	for n, _ := range sam {
		ng.AddNode(n)
	}
	for n, sa := range sam {
		if sa.pi != nil {
			pi := sa.pi
			wt, _ := g.Weight(pi, n)
			ng.AddEdge(pi, n, wt)
		}
	}
	return ng
}

// Return Single Source Shortest Paths started from s via DAG Toposort first O(V)
// Advantage: Fast
// Shortage: DAG only (Directed Acyclic Graph)
func DAGShortestPaths(g Graph, s interface{}) Graph {
	sg := TopoSort(g)
	sam := initSingleSourceAuxMap(sg, s)
	for _, u := range sg.Nodes() {
		adjs, _ := g.Adj(u)
		for _, v := range adjs {
			relax(g, sam, u, v)
		}
	}

	ng := NewGraph(g.Directed(), g.GraphType())
	for n, _ := range sam {
		ng.AddNode(n)
	}
	for n, sa := range sam {
		if sa.pi != nil {
			pi := sa.pi
			wt, _ := g.Weight(pi, n)
			ng.AddEdge(pi, n, wt)
		}
	}
	return ng
}

// Return Single Source Shortest Paths started from s via Dijkstra's Algorithm O((V+E)lgV)
// Advantage: Fast
// Shortage: No negative weight edge supported
func Dijkstra(g Graph, s interface{}) Graph {
	Q := ds.NewMinPriorityQueue()
	sam := initSingleSourceAuxMap(g, s)
	for _, sa := range sam {
		Q.Push(sa, sa.d)
	}

	ng := NewGraph(g.Directed(), g.GraphType())
	for {
		if Q.Empty() {
			break
		}
		uitem := Q.Pop().(*ds.Item)
		sa := uitem.Value.(*SSSAux)
		u := sa.c
		sa.isolated = false
		ng.AddNode(u)
		if sa.pi != nil {
			pi := sa.pi
			wt, _ := g.Weight(pi, u)
			ng.AddEdge(pi, u, wt)
		}
		adjs, _ := g.Adj(u)
		for _, v := range adjs {
			if sam[v].isolated {
				relax(g, sam, u, v)
				vitem := Q.Item(sam[v])
				Q.Update(vitem, sam[v].d)
			}
		}
	}
	return ng
}
