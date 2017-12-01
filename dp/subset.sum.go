package main

import (
	"fmt"
)

func SubsetSum(nums []int, target int) int {
	sums := make([][]int, target+1)
	for i := 0; i <= target; i++ {
		sums[i] = make([]int, len(nums)+1)
		sums[i][0] = 1
	}
	for j := 0; j <= len(nums); j++ {
		sums[0][j] = 2
	}
	subsetSum(nums, target, len(nums), sums)
	for i := 0; i <= target; i++ {
		fmt.Printf("%v\n", sums[i])
	}
	return sums[target][len(nums)]
}

// l: length of nums
// sums: store the result
func subsetSum(nums []int, t int, l int, sums [][]int) {
	if sums[t][l] > 0 {
		return
	}
	if l == 0 || t == 0 {
		// already initialized
		return
	}
	if nums[l-1] == t {
		sums[t][l] = 2
		return
	}
	// Not including the last element
	subsetSum(nums, t, l-1, sums)
	sums[t][l] = sums[t][l-1]
	if t > nums[l-1] {
		// Including the last element
		subsetSum(nums, t-nums[l-1], l-1, sums)
		if sums[t][l] < sums[t-nums[l-1]][l-1] {
			sums[t][l] = sums[t-nums[l-1]][l-1]
		}
	}
	return
}

func main() {
	nums := []int{1, 3, 7, 5, 8, 2}
	fmt.Printf("%v\n", SubsetSum(nums, 9))
}
