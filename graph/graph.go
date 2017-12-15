package graph

import (
	"fmt"
)

type (
	Graph interface {
		AddNode(v interface{})
		AddEdge(sv, ev interface{}, weight int) error
		Nodes() []interface{}
		Adj(v interface{}) ([]interface{}, error)
		Weight(sv, ev interface{}) (int, error)

		fmt.Stringer
	}

	NodeAux struct {
		c  interface{} // Value of current node
		pi interface{} // Value of parent node, follow the π notation in CLRS
		d  int         // Depth in BFS, Start time in DFS, path weight estimation in SSSP(Both Bellman-Ford and Dijkstra)
	}
)
