package graph

import (
	"testing"
)

func TestBFS(test *testing.T) {
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

	g.BFS(s)
}

func TestPath(test *testing.T) {
	g := NewGraph(false)
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)
	n4 := g.AddNode(4)
	n5 := g.AddNode(5)
	n6 := g.AddNode(6)
	g.AddEdge(n1, n2, 1)
	g.AddEdge(n1, n5, 1)
	g.AddEdge(n2, n1, 1)
	g.AddEdge(n2, n5, 1)
	g.AddEdge(n2, n4, 1)
	g.AddEdge(n2, n3, 1)
	g.AddEdge(n3, n2, 1)
	g.AddEdge(n3, n4, 1)
	g.AddEdge(n4, n3, 1)
	g.AddEdge(n4, n2, 1)
	g.AddEdge(n4, n5, 1)
	g.AddEdge(n4, n6, 1)
	g.AddEdge(n5, n1, 1)
	g.AddEdge(n5, n2, 1)
	g.AddEdge(n5, n4, 1)

	g.Path(n1, n6)
}
