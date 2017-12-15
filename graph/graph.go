package graph

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

func (g *Graph) AddNode(v interface{}) {
	n := new(Node)
	n.edges = make([]*Edge, 0)
	n.value = v
	g.nodes = append(g.nodes, n)
}

func (g *Graph) checkExist(v interface{}) error {
	if g.Node(v) == nil {
		msg := fmt.Sprintf("ERROR: %v does not exist in graph", v)
		fmt.Println(msg)
		return errors.New(msg)
	}
	return nil
}

func (g *Graph) AddEdge(sv, ev interface{}, weight int) error {
	if err := g.checkExist(sv); err != nil {
		return err
	}
	if err := g.checkExist(ev); err != nil {
		return err
	}
	sn, en := g.Node(sv), g.Node(ev)

	// Validate existing edges
	if g.directed {
		for _, e := range sn.edges {
			if e.end == en {
				msg := fmt.Sprintf("ERROR: Edge (%v, %v) already exist", sv, ev)
				fmt.Println(msg)
				return errors.New(msg)
			}
		}
	} else {
		for _, e := range sn.edges {
			if e.end == en || e.start == en {
				msg := fmt.Sprintf("ERROR: Edge (%v, %v) already exist", sv, ev)
				fmt.Println(msg)
				return errors.New(msg)
			}
		}
	}

	ne := &Edge{
		weight,
		sn,
		en,
	}
	sn.edges = append(sn.edges, ne)
	if sn != en {
		en.edges = append(en.edges, ne)
	}
	return nil
}

func (g *Graph) RemoveEdge(sv, ev interface{}) error {
	if err := g.checkExist(sv); err != nil {
		return err
	}
	if err := g.checkExist(ev); err != nil {
		return err
	}
	var e *Edge
	if e = g.Edge(sv, ev); e == nil {
		msg := fmt.Sprintf("ERROR: Edge (%v, %v) does not exist", sv, ev)
		fmt.Println(msg)
		return errors.New(msg)
	}

	// Remove edge
	cnt := 0
	for i := 0; i < len(e.start.edges); i++ {
		if e.start.edges[i] == e {
			e.start.edges = append(e.start.edges[:i], e.start.edges[i+1:]...)
			cnt += 1
			break
		}
	}
	for i := 0; i < len(e.end.edges); i++ {
		if e.end.edges[i] == e {
			e.end.edges = append(e.end.edges[:i], e.end.edges[i+1:]...)
			cnt += 1
			break
		}
	}
	if (e.start != e.end && cnt != 2) || (e.start == e.end && cnt != 1) {
		msg := fmt.Sprintf("Invalid endpoints for edge <%v, %v>", e.start.value, e.end.value)
		fmt.Printf(msg)
		return errors.New(msg)
	}
	return nil
}

func (g *Graph) RemoveNode(tv interface{}) error {
	if err := g.checkExist(tv); err != nil {
		return err
	}
	t := g.Node(tv)
	var n *Node
	var idx int
	for i, nd := range g.nodes {
		if nd == t {
			n = nd
			idx = i
			break
		}
	}

	// NOTE: Do remember to avoid the remove while iterate range loop pitfall!
	// If you do a for _, e := range n.edges { ... }, e will be pointed to some edge that's
	// already deleted
	for _ = range n.edges {
		e := n.edges[0]
		g.RemoveEdge(e.start.value, e.end.value)
		e = nil
	}
	g.nodes = append(g.nodes[:idx], g.nodes[idx+1:]...)
	return nil
}

func (g *Graph) Node(v interface{}) *Node {
	for _, n := range g.nodes {
		if n.value == v {
			return n
		}
	}
	return nil
}

func (g *Graph) Edge(sv, ev interface{}) *Edge {
	var n1, n2 *Node
	switch sv.(type) {
	case *Node:
		n1 = sv.(*Node)
		n2 = ev.(*Node)
	default:
		n1 = g.Node(sv)
		n2 = g.Node(ev)
	}
	for _, n := range g.nodes {
		if n == n1 {
			for _, e := range n.edges {
				if g.directed {
					if e.end == n2 {
						return e
					}
				} else {
					if e.start == n2 || e.end == n2 {
						return e
					}
				}
			}
		}
	}
	return nil
}

func (g *Graph) Adj(tv interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0)
	if err := g.checkExist(tv); err != nil {
		return ret, err
	}
	n := g.Node(tv)
	for _, e := range n.edges {
		if e.start == n {
			ret = append(ret, e.end.value)
		}
		if !g.directed && e.start != n {
			ret = append(ret, e.start.value)
		}
	}
	return ret, nil
}

func (g *Graph) String() string {
	ret := "\n\n"
	for _, n := range g.nodes {
		ret += fmt.Sprintf("%-12v: \t", n.value)
		for _, e := range n.edges {
			if g.directed {
				if e.start == n {
					ret += fmt.Sprintf("%v --(%d)--> %v ", e.start.value, e.weight, e.end.value)
				}
			} else {
				ret += fmt.Sprintf("%v --(%d)-- %v ", e.start.value, e.weight, e.end.value)
			}
		}
		ret += "\n"
	}
	ret += "\n"
	return ret
}

// For Directed Graph, return a new graph with every edge direction reverted
func (g *Graph) Transpose() *Graph {
	if !g.directed {
		return g
	}
	ng := NewGraph(g.directed)
	for _, n := range g.nodes {
		ng.AddNode(n.value)
	}
	for _, n := range g.nodes {
		for _, e := range n.edges {
			if e.start == n {
				ng.AddEdge(e.end.value, e.start.value, e.weight)
			}
		}
	}
	return ng
}
