package main

import (
	"fmt"
)

func LongestPalindromicSubsequence(A string) int {
	L := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		L[i] = make([]int, len(A))
	}
	longestPalindromicSubsequence(A, 0, len(A)-1, L)
	for i := 0; i < len(A); i++ {
		fmt.Printf("%v\n", L[i])
	}
	return L[0][len(A)-1]
}

func longestPalindromicSubsequence(A string, start int, end int, L [][]int) {
	if start > end || L[start][end] > 0 {
		return
	}
	if start == end {
		L[start][end] = 1
		return
	}
	if A[start] == A[end] {
		longestPalindromicSubsequence(A, start+1, end-1, L)
		L[start][end] = L[start+1][end-1] + 2
	} else {
		longestPalindromicSubsequence(A, start+1, end, L)
		longestPalindromicSubsequence(A, start, end-1, L)
		if L[start+1][end] > L[start][end-1] {
			L[start][end] = L[start+1][end]
		} else {
			L[start][end] = L[start][end-1]
		}
	}
	return
}

func main() {
	fmt.Printf("%v\n", LongestPalindromicSubsequence("B"))
	fmt.Printf("%v\n", LongestPalindromicSubsequence("BB"))
	fmt.Printf("%v\n", LongestPalindromicSubsequence("BBA"))
	fmt.Printf("%v\n", LongestPalindromicSubsequence("BAB"))
	fmt.Printf("%v\n", LongestPalindromicSubsequence("BBBABCBCABBABCBCAB"))
}
