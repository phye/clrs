package graph

import (
	"fmt"
	"testing"
)

func TestBellmanFord(test *testing.T) {
	g := NewGraph(true, LISTGRAPH)
	g.AddNode("s")
	g.AddNode("t")
	g.AddNode("x")
	g.AddNode("y")
	g.AddNode("z")

	g.AddEdge("s", "t", 6)
	g.AddEdge("s", "y", 7)
	g.AddEdge("t", "y", 8)
	g.AddEdge("t", "x", 5)
	g.AddEdge("t", "z", -4)
	g.AddEdge("y", "z", 9)
	g.AddEdge("y", "x", -3)
	g.AddEdge("x", "t", -2)
	g.AddEdge("z", "s", 2)

	ng := BellmanFord(g, "s")
	fmt.Printf("BellmanFord result is %s\n", ng)
}

func TestDAGShortestPaths(test *testing.T) {
	g := NewGraph(true, LISTGRAPH)

	g.AddNode("x")
	g.AddNode("t")
	g.AddNode("y")
	g.AddNode("s")
	g.AddNode("z")
	g.AddNode("r")

	g.AddEdge("s", "x", 6)
	g.AddEdge("s", "t", 2)
	g.AddEdge("t", "x", 7)
	g.AddEdge("t", "y", 4)
	g.AddEdge("t", "z", 2)
	g.AddEdge("r", "s", 5)
	g.AddEdge("r", "t", 3)
	g.AddEdge("x", "y", -1)
	g.AddEdge("x", "z", 1)
	g.AddEdge("y", "z", -2)

	ng := DAGShortestPaths(g, "s")
	fmt.Printf("DAG shortest path is %s\n", ng)
}

func TestDijkstra(test *testing.T) {
	g := NewGraph(true, LISTGRAPH)

	g.AddNode("x")
	g.AddNode("t")
	g.AddNode("y")
	g.AddNode("s")
	g.AddNode("z")

	g.AddEdge("s", "t", 10)
	g.AddEdge("s", "y", 5)
	g.AddEdge("t", "y", 2)
	g.AddEdge("t", "x", 1)
	g.AddEdge("x", "z", 4)
	g.AddEdge("y", "t", 3)
	g.AddEdge("y", "x", 9)
	g.AddEdge("y", "z", 2)
	g.AddEdge("z", "s", 2)
	g.AddEdge("z", "x", 6)

	ng := Dijkstra(g, "s")
	fmt.Printf("SSSP via Dijkstra is %s\n", ng)

	mg := NewGraph(true, MATRIXGRAPH)

	mg.AddNode("x")
	mg.AddNode("t")
	mg.AddNode("y")
	mg.AddNode("s")
	mg.AddNode("z")

	mg.AddEdge("s", "t", 10)
	mg.AddEdge("s", "y", 5)
	mg.AddEdge("t", "y", 2)
	mg.AddEdge("t", "x", 1)
	mg.AddEdge("x", "z", 4)
	mg.AddEdge("y", "t", 3)
	mg.AddEdge("y", "x", 9)
	mg.AddEdge("y", "z", 2)
	mg.AddEdge("z", "s", 2)
	mg.AddEdge("z", "x", 6)

	mng := Dijkstra(mg, "s")
	fmt.Printf("SSSP via Dijkstra of matrix graph is %s\n", mng)

}
