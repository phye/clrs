package graph

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func extendShortestPaths(L, W [][]int) [][]int {
	n := len(L)
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ret[i][j] = INF
			for k := 0; k < n; k++ {
				ret[i][j] = min(ret[i][j], L[i][k]+W[k][j])
			}
		}
	}
	return ret
}

// Return All Pairs Shortest Paths via Slow Matrix Mulplication similar algo in O(V^4)
func (mg *MGraph) SlowAllPairsShortestPaths() [][]int {
	n := len(mg.matrix)
	Ls := make([][][]int, n)
	Ls[1] = mg.matrix
	for m := 2; m < n; m++ {
		displayMatrix(Ls[m-1])
		Ls[m] = extendShortestPaths(Ls[m-1], mg.matrix)
	}
	return Ls[n-1]
}

func displayMatrix(m [][]int) {
	fmt.Printf("------------------------------\n")
	for i := 0; i < len(m); i++ {
		fmt.Printf("\t")
		for j := 0; j < len(m[i]); j++ {
			fmt.Printf("%8x ", m[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("------------------------------\n")
}

// Return All Pairs Shortest Paths via Faster Matrix Mulplication similar algo in O(V^3lgV)
func (mg *MGraph) FasterAllPairsShortestPaths() [][]int {
	n := len(mg.matrix)
	Ls := make([][][]int, n)
	Ls[1] = mg.matrix
	m := 1
	for {
		if m >= n-1 {
			break
		}
		L := extendShortestPaths(Ls[m], Ls[m])
		if 2*m > n {
			Ls = append(Ls, L)
		} else {
			Ls[2*m] = L
		}
		m = 2 * m
	}
	return Ls[len(Ls)-1]
}

// Return All Pair Shortest Paths via FloydWarshall algorithm which runs in O(V^3)
func (mg *MGraph) FloydWarshall() [][]int {
	n := len(mg.matrix)
	Ds := make([][][]int, n+1)
	Ds[0] = mg.matrix
	// Note that k represents # of intermediate nodes
	for k := 1; k <= n; k++ {
		Ds[k] = make([][]int, n)
		for i := 0; i < n; i++ {
			Ds[k][i] = make([]int, n)
			for j := 0; j < n; j++ {
				// Note the k-1 in Ds[k-1][i][k-1], the first k-1 means # of intermediate nodes is
				// k-1, the latter k-1 is the index of last node of the k nodes in (0,1,..k-1)
				Ds[k][i][j] = min(Ds[k-1][i][j], Ds[k-1][i][k-1]+Ds[k-1][k-1][j])
			}
		}
	}
	return Ds[n]
}
