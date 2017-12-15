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

		fmt.Stringer
	}

	NodeAux struct {
		c  interface{} // Value of current node
		pi interface{} // Value of parent node, follow the π notation in CLRS
		d  int         // Depth in BFS, Start time in DFS, path weight estimation in SSSP(Both Bellman-Ford and Dijkstra)
	}
)
