package main

import (
	"errors"
	"fmt"
)

type (
	Node struct {
		value interface{}
		edges []*Edge
	}

	Edge struct {
		weight int
		start  *Node
		end    *Node
	}

	Graph struct {
		nodes    []*Node
		directed bool
	}
)

func NewGraph(directed bool) *Graph {
	g := &Graph{
		make([]*Node, 0),
		directed,
	}
	return g
}

func (g *Graph) AddNode(v interface{}) *Node {
	n := new(Node)
	n.edges = make([]*Edge, 0)
	n.value = v
	g.nodes = append(g.nodes, n)
	return n
}

func (g *Graph) AddEdge(sn *Node, en *Node, weight int) (*Edge, error) {
	if sn == en {
		return nil, errors.New("Same node")
	}
	// Validate start node and end node
	cnt := 0
	for _, n := range g.nodes {
		if n == sn || n == en {
			cnt += 1
		}
	}
	if cnt != 2 {
		msg := "Node not found in graph"
		fmt.Printf("%v\n", msg)
		return nil, errors.New(msg)
	}

	// Validate existing edges
	if g.directed {
		for _, e := range sn.edges {
			if e.end == en {
				msg := "Edge already exist"
				//fmt.Printf("%v\n", msg)
				return e, errors.New(msg)
			}
		}
	} else {
		for _, e := range sn.edges {
			if e.end == en || e.start == en {
				msg := "Edge already exist"
				//fmt.Printf("%v\n", msg)
				return e, errors.New(msg)
			}
		}
	}

	ne := &Edge{
		weight,
		sn,
		en,
	}
	sn.edges = append(sn.edges, ne)
	en.edges = append(en.edges, ne)
	return ne, nil
}

func (g *Graph) RemoveEdge(e *Edge) error {
	if e == nil {
		msg := "Nil edge"
		fmt.Printf(msg)
		return errors.New(msg)
	}

	fmt.Printf("Removing edge <%v, %v>\n", e.start.value, e.end.value)
	// Validate edge
	cnt := 0
	for _, n := range g.nodes {
		if n == e.start || n == e.end {
			cnt += 1
		}
	}
	if cnt != 2 {
		msg := "Edge not found in graph"
		fmt.Printf(msg)
		return errors.New(msg)
	}

	// Remove edge
	cnt = 0
	for i := 0; i < len(e.start.edges); i++ {
		if e.start.edges[i] == e {
			e.start.edges = append(e.start.edges[:i], e.start.edges[i+1:]...)
			fmt.Printf("Removing edge <%v, %v> at node %v\n", e.start.value, e.end.value, e.start.value)
			cnt += 1
			break
		}
	}
	for i := 0; i < len(e.end.edges); i++ {
		if e.end.edges[i] == e {
			e.end.edges = append(e.end.edges[:i], e.end.edges[i+1:]...)
			fmt.Printf("Removing edge <%v, %v> at node %v\n", e.start.value, e.end.value, e.end.value)
			cnt += 1
			break
		}
	}
	if cnt != 2 {
		msg := fmt.Sprintf("Invalid endpoints for edge <%v, %v>", e.start.value, e.end.value)
		fmt.Printf(msg)
		return errors.New(msg)
	}
	return nil
}

func (g *Graph) RemoveNode(t *Node) error {
	var n *Node
	var idx int
	for i, nd := range g.nodes {
		if nd == t {
			n = nd
			idx = i
			break
		}
	}
	if n == nil {
		return errors.New("Node not found in graph")
	}

	// NOTE: Do remember to avoid the remove while iterate range loop pitfall!
	// If you do a for _, e := range n.edges { ... }, e will be pointed to some edge that's
	// already deleted
	for _ = range n.edges {
		e := n.edges[0]
		g.RemoveEdge(e)
		e = nil
	}
	g.nodes = append(g.nodes[:idx], g.nodes[idx+1:]...)
	return nil
}

func (g *Graph) String() string {
	ret := "\n\n"
	for _, n := range g.nodes {
		ret += fmt.Sprintf("%v: %p \t", n.value, n)
		for _, e := range n.edges {
			if g.directed {
				ret += fmt.Sprintf("%v --> %v ", e.start.value, e.end.value)
			} else {
				ret += fmt.Sprintf("%v --- %v ", e.start.value, e.end.value)
			}
		}
		ret += "\n"
	}
	return ret
}

func main() {
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
	fmt.Printf("%s", g)
}
