package graph

import (
	"fmt"
	"testing"
)

func TestPrimMST(test *testing.T) {
	gh := NewGraph(false)
	a := gh.AddNode("a")
	b := gh.AddNode("b")
	c := gh.AddNode("c")
	d := gh.AddNode("d")
	e := gh.AddNode("e")
	f := gh.AddNode("f")
	g := gh.AddNode("g")
	h := gh.AddNode("h")
	i := gh.AddNode("i")

	gh.AddEdge(a, b, 4)
	gh.AddEdge(a, h, 8)
	gh.AddEdge(b, c, 8)
	gh.AddEdge(b, h, 11)
	gh.AddEdge(c, i, 2)
	gh.AddEdge(c, d, 7)
	gh.AddEdge(c, f, 4)
	gh.AddEdge(d, e, 9)
	gh.AddEdge(d, f, 14)
	gh.AddEdge(e, f, 10)
	gh.AddEdge(f, g, 2)
	gh.AddEdge(g, h, 1)
	gh.AddEdge(g, i, 6)
	gh.AddEdge(h, i, 7)

	ngh := gh.PrimMST(a)
	fmt.Println("\n----------------PRIM MST ----------------")
	fmt.Printf("Original graph is %s", gh)
	fmt.Printf("Prim MST is %s", ngh)
	fmt.Println("----------------PRIM MST ----------------")
}
