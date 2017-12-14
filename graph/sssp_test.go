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
