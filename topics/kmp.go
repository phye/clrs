package main

import (
	"fmt"
)

func KMPMatcher(T []byte, P []byte) {
	n := len(T)
	m := len(P)
	pf := ComputePrefixFunction(P)
	q := 0 // Number of characters matched
	for i := 0; i < n; i++ {
		for ; q > 0 && P[q] != T[i]; q = pf[q] {
		}
		if P[q] == T[i] {
			q = q + 1
		}
		if q == m {
			fmt.Printf("Pattern occurs with shift %v\n", i-m+1)
			q = pf[m-1]
		}
	}
}

func ComputePrefixFunction(P []byte) []int {
	m := len(P)
	ret := []int{0}
	k := 0
	for q := 1; q < m; q++ {
		for ; k > 0 && P[k] != P[q]; k = ret[k] {
		}
		if P[k] == P[q] {
			k = k + 1
		}
		ret = append(ret, k)
	}
	fmt.Printf("%v\n", ret)
	return ret
}

func main() {
	T := []byte("bacbababaabcbababa")
	P := []byte("ababa")

	KMPMatcher(T, P)
}
