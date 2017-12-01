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

	BFS(g, s)
}
