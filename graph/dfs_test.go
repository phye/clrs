package graph

import (
	"fmt"
	"testing"
)

func TestDFS(test *testing.T) {
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

	DFS(g)
}

func TestDFSDirected(test *testing.T) {
	g := NewListGraph(true)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(7)
	g.AddNode(8)

	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 7, 1)
	g.AddEdge(3, 8, 1)
	g.AddEdge(7, 3, 1)

	DFS(g)
}

func TestDFSDirectedNotConnected(test *testing.T) {
	g := NewListGraph(true)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(7)
	g.AddNode(8)
	g.AddNode(1)

	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 1)

	g.AddEdge(7, 8, 1)

	DFS(g)
}

func TestTopoSort(test *testing.T) {
	g := NewListGraph(true)
	g.AddNode("shirt")
	g.AddNode("tie")
	g.AddNode("jacket")
	g.AddNode("undershorts")
	g.AddNode("pants")
	g.AddNode("belt")
	g.AddNode("socks")
	g.AddNode("shoes")
	g.AddNode("watch")

	g.AddEdge("undershorts", "pants", 1)
	g.AddEdge("undershorts", "shoes", 1)
	g.AddEdge("pants", "belt", 1)
	g.AddEdge("pants", "shoes", 1)
	g.AddEdge("belt", "jacket", 1)
	g.AddEdge("shirt", "belt", 1)
	g.AddEdge("shirt", "tie", 1)
	g.AddEdge("tie", "jacket", 1)
	g.AddEdge("socks", "shoes", 1)

	ret := TopoSort(g)
	fmt.Printf("After Toposort, the graph is %s\n", ret)
}

/*
func TestDecomposition(test *testing.T) {
	gh := NewListGraph(true)
	a := gh.AddNode("a")
	b := gh.AddNode("b")
	c := gh.AddNode("c")
	d := gh.AddNode("d")
	e := gh.AddNode("e")
	f := gh.AddNode("f")
	g := gh.AddNode("g")
	h := gh.AddNode("h")

	gh.AddEdge(a, b, 1)
	gh.AddEdge(e, a, 1)
	gh.AddEdge(b, c, 1)
	gh.AddEdge(b, e, 1)
	gh.AddEdge(b, f, 1)
	gh.AddEdge(c, d, 1)
	gh.AddEdge(c, g, 1)
	gh.AddEdge(d, c, 1)
	gh.AddEdge(d, h, 1)
	gh.AddEdge(e, f, 1)
	gh.AddEdge(f, g, 1)
	gh.AddEdge(g, f, 1)
	gh.AddEdge(g, h, 1)
	gh.AddEdge(h, h, 1)

	pgs := gh.Decomposition()
	fmt.Printf("After decomposition, strongly connected components are ")
	for _, pg := range pgs {
		fmt.Printf("%s", pg)
	}
}
*/
