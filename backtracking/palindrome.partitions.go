package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	ret := make([][]string, 0)
	sol := make([]string, 0)
	partitionAux(s, sol, 0, &ret)
	return ret
}

func partitionAux(s string, sol []string, ite int, ret *[][]string) {
	if ite >= len(s) {
		dst := make([]string, len(sol))
		copy(dst, sol)
		*ret = append(*ret, dst)
		return
	}
	for i := ite; i < len(s); i++ {
		if isPalindrome(s[ite : i+1]) {
			sol = append(sol, s[ite:i+1])
			partitionAux(s, sol, i+1, ret)
			sol = sol[:len(sol)-1]
		}
	}
}

func main() {
	ret := partition("cbbbcc")
	for i := 0; i < len(ret); i++ {
		fmt.Printf("%v\n", ret[i])
	}
}
