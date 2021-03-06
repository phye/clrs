package graph

import (
	"fmt"
)

type (
	Graph interface {
		// Add node v into graph, interface{} MUST be comparable! (thus cannot be slice)
		AddNode(v interface{})
		// Add edge (sv, ev) into graph, if graph is not directed, two edges will be added
		AddEdge(sv, ev interface{}, weight int) error
		// Return a slice of all nodes' keys in graph
		Nodes() []interface{}
		// Return adjacent nodes' keys of node v
		Adj(v interface{}) ([]interface{}, error)
		// Return weight of edge (sv, ev)
		Weight(sv, ev interface{}) (int, error)
		// Return a new graph with all edges direction reverted
		Transpose() Graph
		// Return a deep copy of the original graph
		Clone() Graph
		// Return graph type. FIXME: May not be necessary with support of reflection
		GraphType() GType
		// Is the graph directed
		Directed() bool

		fmt.Stringer
	}

	NodeAux struct {
		c  interface{} // Value of current node
		pi interface{} // Value of parent node, follow the π notation in CLRS
		d  int         // Depth in BFS, Start time in DFS, path weight estimation in SSSP(Both Bellman-Ford and Dijkstra)
	}

	GType int // Graph type
)

const (
	LISTGRAPH GType = iota
	MATRIXGRAPH
)

func NewGraph(directed bool, gt GType) Graph {
	var g Graph
	switch gt {
	case LISTGRAPH:
		g = &LGraph{
			make([]*Node, 0),
			directed,
		}
	case MATRIXGRAPH:
		g = &MGraph{
			make([][]int, 0),
			directed,
			make([]interface{}, 0),
			make(map[interface{}]int),
		}
	}
	return g
}

// Given Graph g and nodes of subsets, return a new graph with the corresponding edges in the
// original graph
func SubGraph(g Graph, nodes []interface{}) Graph {
	ng := NewGraph(g.Directed(), g.GraphType())
	for _, n := range nodes {
		ng.AddNode(n)
	}
	for _, u := range nodes {
		adjs, _ := g.Adj(u)
		for _, v := range adjs {
			// FIXME: Need to check if v is in nodes?
			wt, _ := g.Weight(u, v)
			ng.AddEdge(u, v, wt)
		}
	}
	return ng
}
