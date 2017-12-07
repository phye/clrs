package graph

import (
	"fmt"
	"testing"
)

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

	l := g.TopoSort()
	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Node)
		fmt.Printf("%v ", v.value)
	}
	fmt.Printf("\n")
}
