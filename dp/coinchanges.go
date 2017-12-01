package main

import (
	"fmt"
)

const (
	INVALID int = -1
)

func CoinChanges(N int, S []int) int {
	m := len(S)
	counts := make(map[int]map[int]int, 0)
	counts[0] = make(map[int]int, 0)
	counts[0][0] = 1
	for i := 1; i <= N; i++ {
		counts[0][i] = 0
	}
	coinChanges(N, S, counts)
	return counts[m][N]
}

func coinChanges(N int, S []int, counts map[int]map[int]int) {
	m := len(S)
	if _, ok := counts[m]; !ok {
		counts[m] = make(map[int]int, 0)
	}
	if _, ok := counts[m][N]; ok {
		return
	}
	if N == 0 {
		counts[m][N] = 1
		return
	} else if N < 0 {
		return
	} else {
		counts[m][N] = 0
		coinChanges(N, S[:m-1], counts)
		coinChanges(N-S[m-1], S, counts)
		counts[m][N] += counts[m-1][N]
		counts[m][N] += counts[m][N-S[m-1]]
	}
}

func CoinChanges2(N int, S []int) int {
	table := make([]int, N+1)
	table[0] = 1
	for i := 0; i < len(S); i++ {
		for j := S[i]; j <= N; j++ {
			table[j] += table[j-S[i]]
		}
	}
	return table[N]
}

func main() {
	fmt.Printf("%v\n", CoinChanges(10, []int{2, 5, 3, 6}))
	fmt.Printf("%v\n", CoinChanges2(10, []int{2, 5, 3, 6}))
}
