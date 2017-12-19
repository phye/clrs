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
