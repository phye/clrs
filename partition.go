package main

import (
	"fmt"
)

// Determine whether a given set can be partitioned into two subsets such that the sum of elements
// in both sets is same

func Partition(set []int) bool {
	sum := 0
	for i := 0; i < len(set); i++ {
		sum += set[i]
	}
	if sum%2 != 0 {
		return false
	} else {
		return partition(set, sum/2)
	}
}

func partition(set []int, target int) bool {
	fmt.Printf("%v %v\n", set, target)
	if len(set) == 0 {
		if target != 0 {
			return false
		} else {
			return true
		}
	} else if len(set) == 1 {
		if set[0] == target {
			return true
		} else {
			return false
		}
	} else {
		inc := partition(set[:len(set)-1], target-set[len(set)-1])
		if inc {
			return true
		}
		exc := partition(set[:len(set)-1], target)
		if exc {
			return true
		}
		return false
	}
}

func main() {
	fmt.Printf("%v\n", Partition([]int{11, 5, 1, 5}))
	//fmt.Printf("%v\n", Partition([]int{1, 5, 11, 6}))
	//fmt.Printf("%v\n", Partition([]int{1, 5, 3}))
}
