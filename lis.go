package main

// Longest Increasing Subsequence

import (
	"fmt"
)

func LIS0(A []int) int {
	lis := make(map[int]int, 0)
	_LIS0(A, len(A)-1, lis)
	max := -1
	for i := 0; i < len(A); i++ {
		if lis[i] > max {
			max = lis[i]
		}
	}
	return max
}

func _LIS0(A []int, i int, lis map[int]int) {
	if _, ok := lis[i]; ok {
		return
	}
	max := -1
	for j := i - 1; j >= 0; j-- {
		_LIS0(A, j, lis)
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

func LIS(A []int) int {
	max := 1
	_LIS(A, len(A)-1, &max)
	return max
}

// Return: length of increasing subsequence at index n
// max: Length of LIS in A
func _LIS(A []int, i int, pm *int) int {
	if i == 0 {
		return 1
	}
	ret := 1
	for j := 0; j < i; j++ {
		oret := _LIS(A, j, pm)
		if A[j] < A[i] && oret+1 > ret {
			ret = oret + 1
		}
	}
	if ret > *pm {
		*pm = ret
	}
	return ret
}

func main() {
	A := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	fmt.Printf("%v\n", LIS0(A))
	fmt.Printf("%v\n", LIS(A))
}
