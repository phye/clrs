package graph

import (
	"fmt"
	"github.com/phye/clrs/ds"
)

type (
	PRIMAux struct {
		c        *Node // Current node
		pi       *Node // Parent node
		isolated bool
	}
	PRIMStatus map[*Node]*PRIMAux
)

func (g *Graph) PrimMST(r *Node) PRIMStatus {
	ps := make(PRIMStatus, 0)
	Q := ds.NewMinPriorityQueue()
	for _, n := range g.nodes {
		ps[n] = &PRIMAux{n, nil, true}
		Q.Push(ps[n], INF)
	}
	Q.Update(Q.Item(ps[r]), 0)
	for {
		if Q.Empty() {
			break
		}
		//uii := Q.Pop() // u item interface, interface{} which is often *ds.Item
		ui := Q.Pop().(*ds.Item)
		pma := ui.Value.(*PRIMAux) // Value is interface {} which is *PRIMAux in this case
		u := pma.c
		pma.isolated = false
		fmt.Printf("Processing node %s\n", u.value)
		for _, v := range g.Adj(u) {
			if ps[v].isolated && g.Edge(u, v).weight < Q.Item(ps[v]).Priority {
				ps[v].pi = u
				fmt.Printf("\tProcessing node %s\n", v.value)
				Q.Update(Q.Item(ps[v]), g.Edge(u, v).weight)
			}
		}
	}
	return ps
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
