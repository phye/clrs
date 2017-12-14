package graph

import (
	"fmt"
	"testing"
)

func TestBellmanFord(test *testing.T) {
	g := NewGraph(true)
	s := g.AddNode("s")
	t := g.AddNode("t")
	x := g.AddNode("x")
	y := g.AddNode("y")
	z := g.AddNode("z")

	g.AddEdge(s, t, 6)
	g.AddEdge(s, y, 7)
	g.AddEdge(t, y, 8)
	g.AddEdge(t, x, 5)
	g.AddEdge(t, z, -4)
	g.AddEdge(y, z, 9)
	g.AddEdge(y, x, -3)
	g.AddEdge(x, t, -2)
	g.AddEdge(z, s, 2)

	ng := g.BellmanFord(s)
	fmt.Printf("BellmanFord is %s\n", ng)
}

func TestDAGShortestPaths(test *testing.T) {
	g := NewGraph(true)

	x := g.AddNode("x")
	t := g.AddNode("t")
	y := g.AddNode("y")
	s := g.AddNode("s")
	z := g.AddNode("z")
	r := g.AddNode("r")

	g.AddEdge(s, x, 6)
	g.AddEdge(s, t, 2)
	g.AddEdge(t, x, 7)
	g.AddEdge(t, y, 4)
	g.AddEdge(t, z, 2)
	g.AddEdge(r, s, 5)
	g.AddEdge(r, t, 3)
	g.AddEdge(x, y, -1)
	g.AddEdge(x, z, 1)
	g.AddEdge(y, z, -2)

	ng := g.DAGShortestPaths(s)
	fmt.Printf("DAG shortest path is %s\n", ng)
}
