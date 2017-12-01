package main

import (
	"fmt"
)

func Knapsack(val []int, wt []int, W int) int {
	P := make(map[int]map[int]int, len(val)+1)
	for i := 0; i <= len(val); i++ {
		P[i] = make(map[int]int, W+1)
		P[i][0] = 0
	}
	knapsack(val, wt, W, len(val), P)
	return P[len(val)][W]
}

func knapsack(val []int, wt []int, W int, n int, P map[int]map[int]int) {
	if _, ok := P[n][W]; ok {
		return
	}
	if n == 0 {
		P[n][W] = 0
		return
	}
	fmt.Printf("%v %v\n", n, W)
	knapsack(val, wt, W, n-1, P)
	// Since val[n-1] is not chosen, P[n][W] is equal to P[n-1][W]
	P[n][W] = P[n-1][W]
	if W >= wt[n-1] {
		knapsack(val, wt, W-wt[n-1], n-1, P)
		if P[n][W] < P[n-1][W-wt[n-1]]+val[n-1] {
			P[n][W] = P[n-1][W-wt[n-1]] + val[n-1]
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// This is a good example that use the memoized version is more efficient since the matrix is more
// sparse
func Knapsack2(val []int, wt []int, W int) int {
	P := make([][]int, len(val)+1)
	for i := 0; i <= len(val); i++ {
		P[i] = make([]int, W+1)
		for j := 0; j <= W; j++ {
			fmt.Printf("%v %v\n", i, j)
			if i == 0 || j == 0 {
				P[i][j] = 0
			} else if wt[i-1] <= j {
				P[i][j] = max(val[i-1]+P[i-1][j-wt[i-1]], P[i-1][j])
			} else {
				P[i][j] = P[i-1][j]
			}
		}
	}
	return P[len(val)][W]
}

func main() {
	fmt.Printf("%v \n", Knapsack([]int{60, 100, 120}, []int{10, 20, 30}, 50))
	fmt.Printf("%v \n", Knapsack2([]int{60, 100, 120}, []int{10, 20, 30}, 50))
}
