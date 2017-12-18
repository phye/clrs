package graph

import (
	"fmt"
	"testing"
)

func TestAddNode(t *testing.T) {
	g := NewGraph(false, LISTGRAPH)
	g.AddNode(1)
}

func TestRemoveNode(t *testing.T) {
	gi := NewGraph(false, LISTGRAPH)
	g := gi.(*LGraph)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 5, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 4, 1)
	g.AddEdge(2, 5, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 5, 1)
	g.RemoveNode(2)
	if len(g.nodes) != 4 {
		t.Fatalf("Node remove failed!")
	}
}

func TestGetNode(t *testing.T) {
	gi := NewGraph(false, LISTGRAPH)
	g := gi.(*LGraph)
	g.AddNode(5)
	n := g.Node(5)
	//fmt.Printf("%p\n", n)
	if n.value != 5 {
		t.Fatal("g.Node incorrect ptr!")
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph(false, LISTGRAPH)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	if err := g.AddEdge(1, 2, 1); err != nil {
		t.Fatalf(err.Error())
	}
	if err := g.AddEdge(1, 5, 1); err != nil {
		t.Fatalf(err.Error())
	}
	if err := g.AddEdge(2, 5, 1); err != nil {
		t.Fatalf(err.Error())
	}
	if err := g.AddEdge(2, 4, 1); err != nil {
		t.Fatalf(err.Error())
	}
	if err := g.AddEdge(2, 3, 1); err != nil {
		t.Fatalf(err.Error())
	}
	if err := g.AddEdge(3, 4, 1); err != nil {
		t.Fatalf(err.Error())
	}
	// 6 does not exist
	if err := g.AddEdge(4, 6, 1); err == nil {
		t.Fatalf(err.Error())
	}
}

func TestGetEdge(t *testing.T) {
	gi := NewGraph(false, LISTGRAPH)
	g := gi.(*LGraph)
	g.AddNode(1)
	g.AddNode(2)
	g.AddEdge(1, 2, 3)
	e := g.Edge(1, 2)
	if e.weight != 3 {
		t.Fatalf("Incorrect Edge method")
	}
}

func TestAdj(t *testing.T) {
	g := NewGraph(false, LISTGRAPH)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 5, 1)
	g.AddEdge(2, 5, 1)
	g.AddEdge(2, 4, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 5, 1)

	adj1, _ := g.Adj(1)
	adj2, _ := g.Adj(2)
	adj3, _ := g.Adj(3)
	adj4, _ := g.Adj(4)
	adj5, _ := g.Adj(5)
	if len(adj1) != 2 || len(adj2) != 4 || len(adj3) != 2 || len(adj4) != 3 || len(adj5) != 3 {
		t.Fatalf("Incorrect adjacent nodes")
	}
}

func TestDirectedLGraph(t *testing.T) {
	/*
			1  -->      2  -->    7

		      \         |    -/|
		      _\|      \|/  /

					    3   -->   8
	*/
	gi := NewGraph(true, LISTGRAPH)
	g := gi.(*LGraph)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(7)
	g.AddNode(8)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 7, 1)
	g.AddEdge(3, 7, 1)
	g.AddEdge(3, 8, 1)
	adj1, _ := g.Adj(1)
	adj2, _ := g.Adj(2)
	adj3, _ := g.Adj(3)
	adj7, _ := g.Adj(7)
	adj8, _ := g.Adj(8)
	if len(adj1) != 2 || len(adj2) != 2 || len(adj3) != 2 || len(adj7) != 0 || len(adj8) != 0 {
		t.Fatalf("Incorrect adjacent ndoes for directed graph")
	}
	fmt.Printf("%s", g)
	g.RemoveNode(3)
	fmt.Printf("node 3 removed %s", g)
	g.AddEdge(8, 8, 1)
	fmt.Printf("%s", g)
	g.RemoveEdge(2, 7)
	fmt.Printf("Edge (2,7) removed %s", g)
	if err := g.RemoveEdge(2, 1); err == nil {
		t.Fatalf("Removing nonexising edge should return error")
	}
	fmt.Printf("After removing nil Edge (2,1) %s", g)
}

func TestTranspose(t *testing.T) {
	g := NewGraph(true, LISTGRAPH)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(7)
	g.AddNode(8)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 7, 1)
	g.AddEdge(3, 7, 1)
	g.AddEdge(3, 8, 1)
	adj1, _ := g.Adj(1)
	adj2, _ := g.Adj(2)
	adj3, _ := g.Adj(3)
	adj7, _ := g.Adj(7)
	adj8, _ := g.Adj(8)
	if len(adj1) != 2 || len(adj2) != 2 || len(adj3) != 2 || len(adj7) != 0 || len(adj8) != 0 {
		t.Fatalf("Incorrect adjacent ndoes for directed graph")
	}
	fmt.Printf("Original graph: %s", g)
	fmt.Printf("After transpose: %s", g.Transpose())
}
