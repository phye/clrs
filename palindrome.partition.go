package main

// Minimum cuts to form palindrome partition

import (
	"fmt"
)

func PalindromePartition(S []byte) int {
	counts := make([][]int, len(S))
	for i := 0; i < len(S); i++ {
		counts[i] = make([]int, len(S))
	}
	palindromePartition(S, 0, len(S)-1, counts)
	return counts[0][len(S)-1]
}

func palindromePartition(S []byte, start int, end int, counts [][]int) {
	if counts[start][end] > 0 || start > end {
		return
	}
	if start == end {
		// No partition needed
		counts[start][end] = 0
		return
	}
}
