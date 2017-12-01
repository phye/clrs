package main

import (
	"fmt"
)

func LongestCommonSubstring(A string, B string) int {
	// Let dp denotes the length of longest common suffix of two strings
	// Thus dp[i][j] = dp[i-1][j-1] if A[i] == B[j], otherwise, 0
	m := len(A)
	n := len(B)
	dp := make([][]int, m+1)
	longest := 0
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}
			// Note that i/j are iterating over length of substrings
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > longest {
					longest = dp[i][j]
				}
			}
		}
		//fmt.Printf("%v\n", dp[i])
	}
	return longest
}

func LongestCommonSubstringMemoized(A string, B string) int {
	dp := make([][]int, len(A)+1)
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]int, len(B)+1)
	}
	longest := 0
	longestCommonSubstring(A, B, len(A), len(B), dp, &longest)
	for i := 0; i <= len(A); i++ {
		fmt.Printf("%v\n", dp[i])
	}
	return longest
}

func longestCommonSubstring(A string, B string, m int, n int, dp [][]int, longest *int) {
	if dp[m][n] > 0 {
		return
	}
	if m == 0 || n == 0 {
		return
	}
	if A[m-1] == B[n-1] {
		longestCommonSubstring(A, B, m-1, n-1, dp, longest)
		dp[m][n] = dp[m-1][n-1] + 1
		if dp[m][n] > *longest {
			*longest = dp[m][n]
		}
		return
	} else {
		longestCommonSubstring(A, B, m, n-1, dp, longest)
		longestCommonSubstring(A, B, m-1, n, dp, longest)
		return
	}
}

func main() {
	fmt.Printf("%v\n", LongestCommonSubstring("Hello World!", "Hi World, abcabc"))
	fmt.Printf("%v\n", LongestCommonSubstringMemoized("Hello World!", "Hi World, abcabc"))
}
