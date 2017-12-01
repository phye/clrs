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
	fmt.Printf("%s", g)
	g.RemoveNode(n2)
}

func TestGetNode(t *testing.T) {
	g := NewGraph(false)
	n5 := g.AddNode(5)
	n := g.Node(5)
	fmt.Printf("%p\n", n)
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
	fmt.Printf("%s", g)
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
