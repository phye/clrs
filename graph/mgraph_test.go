package graph

import (
	"fmt"
	"testing"
)

func TestMGraphBasic(t *testing.T) {
	g := NewMatrixGraph(false)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 5, 1)
	g.AddEdge(2, 1, 1)
	g.AddEdge(2, 5, 1)
	g.AddEdge(2, 4, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 2, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 3, 1)
	g.AddEdge(4, 2, 1)
	g.AddEdge(4, 5, 1)
	g.AddEdge(5, 1, 1)
	g.AddEdge(5, 2, 1)
	g.AddEdge(5, 4, 1)
	fmt.Printf("%s\n", g)
	adjs, _ := g.Adj(5)
	fmt.Printf("%v\n", adjs)
	if len(adjs) != 3 {
		t.Fatalf("Incorrect Adj for graph")
	}
}

func TestDirectedMGraph(t *testing.T) {
	g := NewMatrixGraph(true)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	g.AddNode(6)

	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 4, 1)
	g.AddEdge(2, 5, 1)
	g.AddEdge(3, 5, 1)
	g.AddEdge(3, 6, 1)
	g.AddEdge(4, 2, 1)
	g.AddEdge(5, 4, 1)
	g.AddEdge(6, 6, 1)
	fmt.Printf("%s\n", g)
	adjs, _ := g.Adj(5)
	fmt.Printf("%v\n", adjs)
	if len(adjs) != 1 {
		t.Fatalf("Incorrect Adj for directed graph")
	}

	ng := g.Transpose()
	fmt.Printf("Transposed matrix graph is \n%s", ng)
}
