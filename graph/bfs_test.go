package graph

import (
	"testing"
)

func TestBFS(t *testing.T) {
	lg := NewGraph(false, LISTGRAPH)
	lg.AddNode("r")
	lg.AddNode("s")
	lg.AddNode("t")
	lg.AddNode("u")
	lg.AddNode("v")
	lg.AddNode("w")
	lg.AddNode("x")
	lg.AddNode("y")

	lg.AddEdge("r", "v", 1)
	lg.AddEdge("r", "s", 1)
	lg.AddEdge("s", "w", 1)
	lg.AddEdge("w", "t", 1)
	lg.AddEdge("w", "x", 1)
	lg.AddEdge("t", "x", 1)
	lg.AddEdge("t", "u", 1)
	lg.AddEdge("x", "u", 1)
	lg.AddEdge("x", "y", 1)
	lg.AddEdge("u", "y", 1)
	lbfs := BFS(lg, "s")

	mg := NewGraph(false, MATRIXGRAPH)
	mg.AddNode("r")
	mg.AddNode("s")
	mg.AddNode("t")
	mg.AddNode("u")
	mg.AddNode("v")
	mg.AddNode("w")
	mg.AddNode("x")
	mg.AddNode("y")

	mg.AddEdge("r", "v", 1)
	mg.AddEdge("r", "s", 1)
	mg.AddEdge("s", "w", 1)
	mg.AddEdge("w", "t", 1)
	mg.AddEdge("w", "x", 1)
	mg.AddEdge("t", "x", 1)
	mg.AddEdge("t", "u", 1)
	mg.AddEdge("x", "u", 1)
	mg.AddEdge("x", "y", 1)
	mg.AddEdge("u", "y", 1)
	mbfs := BFS(mg, "s")

	for v, s := range lbfs {
		if s.d != mbfs[v].d || s.pi != mbfs[v].pi {
			t.Fatalf("Incompatible BFS result!")
		}
	}
}

func TestPath(test *testing.T) {
	g := NewGraph(false, LISTGRAPH)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	g.AddNode(6)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 5, 1)
	g.AddEdge(2, 5, 1)
	g.AddEdge(2, 4, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 5, 1)
	g.AddEdge(4, 6, 1)

	Path(g, 1, 6)
}
