package graph

import (
	"testing"
)

func TestBFS(test *testing.T) {
	g := NewListGraph(false)
	g.AddNode("r")
	g.AddNode("s")
	g.AddNode("t")
	g.AddNode("u")
	g.AddNode("v")
	g.AddNode("w")
	g.AddNode("x")
	g.AddNode("y")

	g.AddEdge("r", "v", 1)
	g.AddEdge("r", "s", 1)
	g.AddEdge("s", "w", 1)
	g.AddEdge("w", "t", 1)
	g.AddEdge("w", "x", 1)
	g.AddEdge("t", "x", 1)
	g.AddEdge("t", "u", 1)
	g.AddEdge("x", "u", 1)
	g.AddEdge("x", "y", 1)
	g.AddEdge("u", "y", 1)

	BFS(g, "s")
}

func TestPath(test *testing.T) {
	g := NewListGraph(false)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	g.AddNode(6)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 5, 1)
	g.AddEdge(2, 5, 1)
	g.AddEdge(2, 4, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 5, 1)
	g.AddEdge(4, 6, 1)

	Path(g, 1, 6)
}
