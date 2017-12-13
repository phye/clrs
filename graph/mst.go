package graph

import (
	"fmt"
	"github.com/phye/clrs/ds"
)

type (
	PRIMAux struct {
		c        *Node // Current node, reverse pointer for quicker lookup
		pi       *Node // Parent node, filled during PRIM MST execution
		isolated bool  // Is the node still in Q?
	}
	PRIMStatus map[*Node]*PRIMAux
)

// Return an MST graph constructed via PRIM algorithm
func (g *Graph) PrimMST(r *Node) *Graph {
	ng := NewGraph(g.directed)
	// Temp map to store g.node <-> ng.node
	nm := make(map[*Node]*Node)

	psm := make(PRIMStatus)
	Q := ds.NewMinPriorityQueue()
	for _, n := range g.nodes {
		psm[n] = &PRIMAux{n, nil, true}
		// NOTE: ds.Item.Value is a pointer to PRIMAux
		Q.Push(psm[n], INF)
	}
	Q.Update(Q.Item(psm[r]), 0)
	for {
		if Q.Empty() {
			break
		}
		uitem := Q.Pop().(*ds.Item)
		pma := uitem.Value.(*PRIMAux) // Value is interface {} which is *PRIMAux in this case
		pma.isolated = false          // Once a node is POPed from Q, it is no longer isolated
		u := pma.c
		nm[u] = ng.AddNode(u.value)
		if psm[u].pi != nil {
			pi := psm[u].pi
			ng.AddEdge(nm[pi], nm[u], g.Edge(pi, u).weight)
		}
		for _, v := range g.Adj(u) {
			// if v is in Q
			if psm[v].isolated {
				weight := g.Edge(u, v).weight
				vitem := Q.Item(psm[v])
				if weight < vitem.Priority {
					psm[v].pi = u
					Q.Update(vitem, weight)
				}
			}
		}
	}
	return ng
}

func (ps PRIMStatus) String() string {
	ret := fmt.Sprintf("\n----------------MST-PRIM------------------\n")
	ret += fmt.Sprintf("%8s %8s\n", "current", "parent")
	for _, n := range ps {
		if n.pi != nil {
			ret += fmt.Sprintf("%8s %8s\n", n.c.value, n.pi.value)
		} else {
			ret += fmt.Sprintf("%8s %8s\n", n.c.value, "NIL")
		}
	}
	return ret
}
