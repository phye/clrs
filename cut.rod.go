package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func CutRod(P []int, n int) int {
	S := make([]int, n+1)
	cutRod(P, n, S)
	fmt.Printf("%v\n", S)
	return S[n]
}

func cutRod(P []int, n int, S []int) {
	if S[n] > 0 || n < 0 {
		return
	}
	if n == 0 {
		S[n] = 0
		return
	}
	if n == 1 {
		S[n] = P[n-1]
		return
	}
	max := 0
	for i := 1; i <= min(len(P), n); i++ {
		cutRod(P, n-i, S)
		if P[i-1]+S[n-i] > max {
			max = P[i-1] + S[n-i]
		}
	}
	S[n] = max
	return
}

func CutRodTabular(P []int, n int) int {
	S := make([]int, n+1)
	for i := 1; i <= n; i++ {
		max := 0
		for j := 1; j <= min(i, len(P)); j++ {
			if max < P[j-1]+S[i-j] {
				max = P[j-1] + S[i-j]
			}
		}
		S[i] = max
	}
	return S[n]
}

func main() {
	fmt.Printf("%v\n", CutRod([]int{1, 5, 8, 9, 10, 17, 17, 20}, 8))
	fmt.Printf("%v\n", CutRod([]int{3, 5, 8, 9, 10, 17, 17, 20}, 8))
	fmt.Printf("%v\n", CutRodTabular([]int{3, 5, 8, 9, 10, 17, 17, 20}, 8))
}
