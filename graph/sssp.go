package graph

type (
	// Single-Source Shortest Paths
	// Note the similarity between BFS and SSS problem:
	// BFS solves the problem of minimum number of edges, while SSS solves the problem of
	// minimum weight of edges
	SSSAux struct {
		d  int   // Shortest-path estimate
		pi *Node // Parent Node in SSS tree
		c  *Node // Reverse Pointer to current node
	}

	SSSAuxMap map[*Node]*SSSAux
)

func (g *Graph) initSingleSourceAuxMap(s *Node) SSSAuxMap {
	sam := make(SSSAuxMap)
	for _, n := range g.nodes {
		sam[n] = &SSSAux{
			INF,
			nil,
			n,
		}
	}
	sam[s].d = 0
	return sam
}

func (g *Graph) relax(sam SSSAuxMap, e *Edge) {
	if sam[e.end].d > sam[e.start].d+e.weight {
		sam[e.end].d = sam[e.start].d + e.weight
		sam[e.end].pi = e.start
	}
}

func (g *Graph) BellmanFord(s *Node) *Graph {
	sam := g.initSingleSourceAuxMap(s)
	for i := 1; i < len(g.nodes)-1; i++ {
		for _, n := range g.nodes {
			for _, e := range n.edges {
				// Only relax edge that's started from the node
				if e.start == n {
					g.relax(sam, e)
				}
			}
		}
	}

	for _, n := range g.nodes {
		for _, e := range n.edges {
			if e.start == n {
				if sam[e.end].d > sam[e.start].d+e.weight {
					return nil
				}
			}
		}
	}

	ng := NewGraph(g.directed)
	mm := make(map[*Node]*Node)
	for n, _ := range sam {
		mm[n] = ng.AddNode(n.value)
	}
	for n, sa := range sam {
		if sa.pi != nil {
			pi := sa.pi
			ng.AddEdge(mm[pi], mm[n], g.Edge(pi, n).weight)
		}
	}
	return ng
}

func (g *Graph) DAGShortestPaths(s *Node) *Graph {
	g.TopoSort()
	sam := g.initSingleSourceAuxMap(s)
	for _, u := range g.nodes {
		for _, v := range g.Adj(u) {
			g.relax(sam, g.Edge(u, v))
		}
	}

	ng := NewGraph(g.directed)
	mm := make(map[*Node]*Node)
	for n, _ := range sam {
		mm[n] = ng.AddNode(n.value)
	}
	for n, sa := range sam {
		if sa.pi != nil {
			pi := sa.pi
			ng.AddEdge(mm[pi], mm[n], g.Edge(pi, n).weight)
		}
	}
	return ng
}
