package main

// Longest Increasing Subsequence

import (
	"fmt"
)

func LIS(A []int) int {
	lis := make(map[int]int, 0)
	_LIS(A, len(A)-1, lis)
	max := -1
	for i := 0; i < len(A); i++ {
		if lis[i] > max {
			max = lis[i]
		}
	}
	return max
}

func _LIS(A []int, i int, lis map[int]int) {
	if _, ok := lis[i]; ok {
		return
	}
	max := -1
	for j := i - 1; j >= 0; j-- {
		_LIS(A, j, lis)
		if A[j] < A[i] && lis[j] > max {
			max = lis[j]
		}
	}
	if max != -1 {
		lis[i] = 1 + max
	} else {
		lis[i] = 1
	}
	return
}

func main() {
	A := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	fmt.Printf("%v\n", LIS(A))
}
