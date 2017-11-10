package main

// Maximum Sum Increasing Subsequence
import (
	"fmt"
)

func MaximumSum(A []int) int {
	sums := make([]int, len(A))
	maximumSum(A, len(A)-1, sums)
	ret := sums[0]
	for i := 1; i < len(A); i++ {
		if ret < sums[i] {
			ret = sums[i]
		}
	}
	fmt.Printf("%v\n", sums)
	return ret
}

func maximumSum(A []int, i int, sums []int) {
	if sums[i] > 0 {
		return
	}
	if i == 0 {
		sums[i] = A[i]
		return
	}
	ret := 0
	for j := 0; j < i; j++ {
		maximumSum(A, j, sums)
		if A[j] < A[i] && sums[j] > ret {
			ret = sums[j]
		}
	}
	sums[i] = ret + A[i]
}

func main() {
	fmt.Printf("%v\n", MaximumSum([]int{1, 101, 2, 3, 100, 4, 5}))
}
