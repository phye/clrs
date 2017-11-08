package main

import (
	"fmt"
	"sort"
)

const (
	INVALID int = -1
)

func CoinChanges(N int, S []int) int {
	sort.Ints(S)
	counts := make(map[int]int)
	for j := 0; j < len(S); j++ {
		counts[S[j]] = 1
	}
	coinChanges(N, S, counts)
	return counts[N]
}

func coinChanges(N int, S []int, counts map[int]int) {
	if _, ok := counts[N]; ok {
		return
	}
	if N < S[0] {
		counts[N] = 0
		return
	}
	counts[N] = 0
	for j := 0; j < len(S); j++ {
		coinChanges(N-S[j], S, counts)
		if counts[N-S[j]] != 0 {
			counts[N] += counts[N-S[j]]
		}
	}
}

func main() {
	fmt.Printf("%v\n", CoinChanges(10, []int{2, 5, 3, 6}))
}
