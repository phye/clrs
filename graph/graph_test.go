package graph

import (
	"fmt"
	"testing"
)

func TestAddNode(t *testing.T) {
	g := NewGraph(false)
	g.AddNode(1)
}

func TestRemoveNode(t *testing.T) {
	g := NewGraph(false)
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)
	n4 := g.AddNode(4)
	n5 := g.AddNode(5)
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
	g.AddEdge(n5, n1, 1)
	g.AddEdge(n5, n2, 1)
	g.AddEdge(n5, n4, 1)
	//fmt.Printf("%s", g)
	g.RemoveNode(n2)
}

func TestGetNode(t *testing.T) {
	g := NewGraph(false)
	n5 := g.AddNode(5)
	n := g.Node(5)
	//fmt.Printf("%p\n", n)
	if n != n5 {
		t.Fatal("g.Node incorrect ptr!")
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph(false)
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)
	n4 := g.AddNode(4)
	n5 := g.AddNode(5)
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
	g.AddEdge(n5, n1, 1)
	g.AddEdge(n5, n2, 1)
	g.AddEdge(n5, n4, 1)
	//fmt.Printf("%s", g)
}

func TestGetEdge(t *testing.T) {
	g := NewGraph(false)
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	e12, _ := g.AddEdge(n1, n2, 3)
	e := g.Edge(n1, n2)
	if e != e12 {
		t.Fatalf("Incorrect Edge method")
	}
}

func TestAdj(t *testing.T) {
	g := NewGraph(false)
	n1 := g.AddNode(1)
	n2 := g.AddNode(2)
	n3 := g.AddNode(3)
	n4 := g.AddNode(4)
	n5 := g.AddNode(5)
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
	g.AddEdge(n5, n1, 1)
	g.AddEdge(n5, n2, 1)
	g.AddEdge(n5, n4, 1)

	adj1 := g.Adj(n1)
	adj2 := g.Adj(n2)
	adj3 := g.Adj(n3)
	adj4 := g.Adj(n4)
	adj5 := g.Adj(n5)
	if len(adj1) != 2 || len(adj2) != 4 || len(adj3) != 2 || len(adj4) != 3 || len(adj5) != 3 {
		t.Fatalf("Incorrect adjacent nodes")
	}
}

func TestDirectedGraph(t *testing.T) {
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
	g.AddEdge(n3, n7, 1)
	g.AddEdge(n3, n8, 1)
	adj1 := g.Adj(n1)
	adj2 := g.Adj(n2)
	adj3 := g.Adj(n3)
	adj7 := g.Adj(n7)
	adj8 := g.Adj(n8)
	if len(adj1) != 2 || len(adj2) != 2 || len(adj3) != 2 || len(adj7) != 0 || len(adj8) != 0 {
		t.Fatalf("Incorrect adjacent ndoes for directed graph")
	}
	fmt.Printf("%s", g)
	g.RemoveNode(n3)
	fmt.Printf("node 3 removed %s", g)
	g.AddEdge(n8, n8, 1)
	fmt.Printf("%s", g)
	e27 := g.Edge(n2, n7)
	g.RemoveEdge(e27)
	fmt.Printf("Edge (2,7) removed %s", g)
	e21 := g.Edge(n2, n1)
	g.RemoveEdge(e21)
	fmt.Printf("After removing nil Edge (2,1) %s", g)
}

func TestTranspose(t *testing.T) {
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
	g.AddEdge(n3, n7, 1)
	g.AddEdge(n3, n8, 1)
	adj1 := g.Adj(n1)
	adj2 := g.Adj(n2)
	adj3 := g.Adj(n3)
	adj7 := g.Adj(n7)
	adj8 := g.Adj(n8)
	if len(adj1) != 2 || len(adj2) != 2 || len(adj3) != 2 || len(adj7) != 0 || len(adj8) != 0 {
		t.Fatalf("Incorrect adjacent ndoes for directed graph")
	}
	fmt.Printf("Original graph: %s", g)
	fmt.Printf("After transpose: %s", g.Transpose())
}
