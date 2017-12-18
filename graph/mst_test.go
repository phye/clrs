package graph

import (
	"fmt"
	"testing"
)

func TestPrimMST(test *testing.T) {
	gh := NewGraph(false, LISTGRAPH)
	gh.AddNode("a")
	gh.AddNode("b")
	gh.AddNode("c")
	gh.AddNode("d")
	gh.AddNode("e")
	gh.AddNode("f")
	gh.AddNode("g")
	gh.AddNode("h")
	gh.AddNode("i")

	gh.AddEdge("a", "b", 4)
	gh.AddEdge("a", "h", 8)
	gh.AddEdge("b", "c", 8)
	gh.AddEdge("b", "h", 11)
	gh.AddEdge("c", "i", 2)
	gh.AddEdge("c", "d", 7)
	gh.AddEdge("c", "f", 4)
	gh.AddEdge("d", "e", 9)
	gh.AddEdge("d", "f", 14)
	gh.AddEdge("e", "f", 10)
	gh.AddEdge("f", "g", 2)
	gh.AddEdge("g", "h", 1)
	gh.AddEdge("g", "i", 6)
	gh.AddEdge("h", "i", 7)

	ngh := PrimMST(gh, "a")
	fmt.Println("\n----------------PRIM MST ----------------")
	fmt.Printf("Original graph is %s", gh)
	fmt.Printf("Prim MST is %s", ngh)
	fmt.Println("----------------PRIM MST ----------------")

	mgh := NewGraph(false, MATRIXGRAPH)
	mgh.AddNode("a")
	mgh.AddNode("b")
	mgh.AddNode("c")
	mgh.AddNode("d")
	mgh.AddNode("e")
	mgh.AddNode("f")
	mgh.AddNode("g")
	mgh.AddNode("h")
	mgh.AddNode("i")

	mgh.AddEdge("a", "b", 4)
	mgh.AddEdge("a", "h", 8)
	mgh.AddEdge("b", "c", 8)
	mgh.AddEdge("b", "h", 11)
	mgh.AddEdge("c", "i", 2)
	mgh.AddEdge("c", "d", 7)
	mgh.AddEdge("c", "f", 4)
	mgh.AddEdge("d", "e", 9)
	mgh.AddEdge("d", "f", 14)
	mgh.AddEdge("e", "f", 10)
	mgh.AddEdge("f", "g", 2)
	mgh.AddEdge("g", "h", 1)
	mgh.AddEdge("g", "i", 6)
	mgh.AddEdge("h", "i", 7)

	ngh = PrimMST(mgh, "a")
	fmt.Println("\n----------------PRIM MST ----------------")
	fmt.Printf("Original graph is %s", mgh)
	fmt.Printf("Prim MST is %s", ngh)
	fmt.Println("----------------PRIM MST ----------------")

}
