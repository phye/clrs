package graph

import (
	"fmt"
	"github.com/phye/clrs/ds"
)

type (
	PRIMAux struct {
		isolated bool // Is the node still in Q?
		NodeAux
	}
	PRIMStatus map[interface{}]*PRIMAux
)

// Return an MST graph constructed via PRIM algorithm
func PrimMST(g Graph, r interface{}) Graph {
	ng := NewGraph(g.Directed(), g.GraphType())

	psm := make(PRIMStatus)
	Q := ds.NewMinPriorityQueue()
	for _, n := range g.Nodes() {
		psm[n] = &PRIMAux{true, NodeAux{n, nil, 0}}
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
		ng.AddNode(u)
		if psm[u].pi != nil {
			pi := psm[u].pi
			wt, _ := g.Weight(pi, u)
			ng.AddEdge(pi, u, wt)
		}
		adjs, _ := g.Adj(u)
		for _, v := range adjs {
			// if v is in Q
			if psm[v].isolated {
				weight, _ := g.Weight(u, v)
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
			ret += fmt.Sprintf("%8s %8s\n", n.c, n.pi)
		} else {
			ret += fmt.Sprintf("%8s %8s\n", n.c, "NIL")
		}
	}
	return ret
}
