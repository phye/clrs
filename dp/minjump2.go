package main

import (
	"fmt"
)

const (
	INF int = 10000
)

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func MinJump(A []int) int {
	m := len(A)
	if m == 0 || A[0] == 0 {
		return INF
	}

	jumps := []int{0}
	for i := 1; i < m; i++ {
		jumps = append(jumps, INF)
		for j := 0; j < i; j++ {
			if i <= j+A[j] && jumps[j] != INF {
				jumps[i] = min(jumps[i], jumps[j]+1)
				break
			}
		}
	}
	return jumps[m-1]
}

func main() {
	fmt.Printf("%v\n", MinJump([]int{1, 3, 6, 1, 0, 9}))
}
