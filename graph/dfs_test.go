package graph

import (
	"fmt"
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
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)
	n7 := g.AddNode(7)
	n8 := g.AddNode(8)
	n1 := g.AddNode(1)

	g.AddEdge(n1, n2, 1)
	g.AddEdge(n1, n3, 1)
	g.AddEdge(n2, n3, 1)

	g.AddEdge(n7, n8, 1)

	g.DFS()
}

func TestTopoSort(test *testing.T) {
	g := NewGraph(true)
	shirt := g.AddNode("shirt")
	tie := g.AddNode("tie")
	jacket := g.AddNode("jacket")
	undershorts := g.AddNode("undershorts")
	pants := g.AddNode("pants")
	belt := g.AddNode("belt")
	socks := g.AddNode("socks")
	shoes := g.AddNode("shoes")
	g.AddNode("watch")

	g.AddEdge(undershorts, pants, 1)
	g.AddEdge(undershorts, shoes, 1)
	g.AddEdge(pants, belt, 1)
	g.AddEdge(pants, shoes, 1)
	g.AddEdge(belt, jacket, 1)
	g.AddEdge(shirt, belt, 1)
	g.AddEdge(shirt, tie, 1)
	g.AddEdge(tie, jacket, 1)
	g.AddEdge(socks, shoes, 1)

	g.TopoSort()
	fmt.Printf("After Toposort, the graph is %s", g)
}

func TestPartition(test *testing.T) {
	gh := NewGraph(true)
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

	pgs := gh.Partition()
	fmt.Printf("After partition, strongly connected components are ")
	for _, pg := range pgs {
		fmt.Printf("%s", pg)
	}
}
