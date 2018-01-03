package main

import (
	"fmt"
	"sort"
)

var tnodes int

func printSubset(A []int) {
	for i := 0; i < len(A); i++ {
		fmt.Printf("%4d", A[i])
	}
	fmt.Printf("\n")
}

func main() {
	weights := []int{10, 7, 5, 18, 12, 20, 15}
	generateSubsets(weights, 35)
	fmt.Printf("Nodes generated %d\n", tnodes)
}

func subset_sum(s, t []int, sum, ite int, tsum int) {
	tnodes++
	if tsum == sum {
		printSubset(t)
		t = t[:len(t)-1]
		if ite+1 < len(s) && sum-s[ite]+s[ite+1] <= tsum {
			subset_sum(s, t, sum-s[ite], ite+1, tsum)
		}
	} else {
		if ite < len(s) && sum+s[ite] < tsum {
			for i := ite; i < len(s); i++ {
				t = append(t, s[i])
				if sum+s[i] <= tsum {
					subset_sum(s, t, sum+s[i], i+1, tsum)
				}
				t = t[:len(t)-1]
			}
		}
	}
}

func generateSubsets(s []int, tsum int) {
	t := make([]int, 0)
	sort.Ints(s)
	var total int
	for i := 0; i < len(s); i++ {
		total += s[i]
	}
	if s[0] <= tsum && total >= tsum {
		subset_sum(s, t, 0, 0, tsum)
	}
}
