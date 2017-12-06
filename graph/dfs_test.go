package graph

import (
	"testing"
)

func TestDFS(test *testing.T) {
	g := NewGraph(false)
	r := g.AddNode("r")
	s := g.AddNode("s")
	t := g.AddNode("t")
	u := g.AddNode("u")
	v := g.AddNode("v")
	w := g.AddNode("w")
	x := g.AddNode("x")
	y := g.AddNode("y")

	g.AddEdge(r, v, 1)
	g.AddEdge(r, s, 1)
	g.AddEdge(s, w, 1)
	g.AddEdge(w, t, 1)
	g.AddEdge(w, x, 1)
	g.AddEdge(t, x, 1)
	g.AddEdge(t, u, 1)
	g.AddEdge(x, u, 1)
	g.AddEdge(x, y, 1)
	g.AddEdge(u, y, 1)

	g.DFS()
}

func TestDFSDirected(test *testing.T) {
	g := NewGraph(true)
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)
	n7 := g.AddNode(7)
	n8 := g.AddNode(8)

	g.AddEdge(n1, n2, 1)
	g.AddEdge(n1, n3, 1)
	g.AddEdge(n2, n3, 1)
	g.AddEdge(n2, n7, 1)
	g.AddEdge(n3, n8, 1)
	g.AddEdge(n7, n3, 1)

	g.DFS()
}

func TestDFSDirectedNotConnected(test *testing.T) {
	g := NewGraph(true)
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)
	n7 := g.AddNode(7)
	n8 := g.AddNode(8)

	g.AddEdge(n1, n2, 1)
	g.AddEdge(n1, n3, 1)
	g.AddEdge(n2, n3, 1)

	g.AddEdge(n7, n8, 1)

	g.DFS()
}
