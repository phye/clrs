package graph

import (
	"fmt"
	"testing"
)

func TestSlowAllPairsShortestPaths(test *testing.T) {
	g := NewGraph(true, MATRIXGRAPH)
	mg := g.(*MGraph)
	mg.AddNode(1)
	mg.AddNode(2)
	mg.AddNode(3)
	mg.AddNode(4)
	mg.AddNode(5)

	mg.AddEdge(1, 2, 3)
	mg.AddEdge(1, 3, 8)
	mg.AddEdge(1, 5, -4)
	mg.AddEdge(2, 4, 1)
	mg.AddEdge(2, 5, 7)
	mg.AddEdge(3, 2, 4)
	mg.AddEdge(4, 1, 2)
	mg.AddEdge(4, 3, -5)
	mg.AddEdge(5, 4, 6)

	res := mg.SlowAllPairsShortestPaths()
	fmt.Printf("Slow All Pairs Shortest Paths matrix is :\n")
	displayMatrix(res)

	res2 := mg.FasterAllPairsShortestPaths()
	fmt.Printf("Faster All Pairs Shortest Paths matrix is :\n")
	displayMatrix(res2)
}
