package graph

import (
	"fmt"
)

type (
	// The following three fields are necessar for BFS
	BFSAux struct {
		color int
		NodeAux
	}

	BFSStatus map[interface{}]*BFSAux
)

const (
	WHITE int = iota
	GRAY
	BLACK

	INF int = 0x7fffffff
)

// Perform Bread-First-Search on Graph g and return the result in BFSStatus
func BFS(g Graph, s interface{}) BFSStatus {
	bfsStatus := make(BFSStatus, 0)
	for _, n := range g.Nodes() {
		bfsStatus[n] = &BFSAux{WHITE, NodeAux{nil, nil, INF}}
	}
	bfsStatus[s].color = GRAY
	bfsStatus[s].d = 0
	bfsStatus[s].pi = nil
	queue := make([]interface{}, 0)
	queue = append(queue, s)
	for {
		if len(queue) == 0 {
			break
		}
		u := queue[0]
		queue = queue[1:]
		adjs, _ := g.Adj(u)
		for _, v := range adjs {
			if bfsStatus[v].color == WHITE {
				bfsStatus[v].color = GRAY
				bfsStatus[v].d = bfsStatus[u].d + 1
				bfsStatus[v].pi = u
				queue = append(queue, v)
			}
		}
		bfsStatus[u].color = BLACK
	}
	fmt.Printf("%s", bfsStatus)
	return bfsStatus
}

// Return the shortest path from sv to ev in a graph which has same weights of all edges.
// Note the difference between Path and SSSP (single-source shortest path)
func Path(g Graph, sv, ev interface{}) []interface{} {
	bfsStatus := BFS(g, sv)
	ret := path(g, sv, ev, bfsStatus)
	fmt.Printf("Path from %v to %v is: \n\t", sv, ev)
	for _, v := range ret {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
	return ret
}

func path(g Graph, sv, ev interface{}, bfsStatus BFSStatus) []interface{} {
	if sv == ev {
		ret := []interface{}{sv}
		return ret
	} else if bfsStatus[ev].pi == nil {
		return []interface{}{}
	} else {
		ret := path(g, sv, bfsStatus[ev].pi, bfsStatus)
		if len(ret) > 0 {
			ret = append(ret, ev)
		}
		return ret
	}
}

func (bfss BFSStatus) String() string {
	var ret string
	ret += fmt.Sprintf("------------BFS------------\n")
	ret += fmt.Sprintf("%8s %8s %8s\n", "current", "parent", "depth")
	for n, s := range bfss {
		if s.pi != nil {
			ret += fmt.Sprintf("%8v %8v %8v\n", n, s.pi, s.d)
		} else {
			ret += fmt.Sprintf("%8v %8v %8v\n", n, "nil", s.d)
		}
	}
	ret += fmt.Sprintf("---------------------------\n")
	return ret
}
