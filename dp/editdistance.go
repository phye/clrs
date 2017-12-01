package main

import (
	"fmt"
)

// Calculate the edit distance of two strings with CUD operations

func editDistance(str1, str2 string, c map[int]map[int]int) {
	m := len(str1)
	n := len(str2)
	if _, ok := c[m][n]; ok {
		return
	}

	if m == 0 {
		c[m][n] = n
		return
	} else if n == 0 {
		c[m][n] = m
		return
	}
	if str1[m-1] == str2[n-1] {
		editDistance(str1[:m-1], str2[:n-1], c)
		c[m][n] = c[m-1][n-1]
		return
	} else {
		editDistance(str1, str2[:n-1], c)
		editDistance(str1[:m-1], str2, c)
		editDistance(str1[:m-1], str2[:n-1], c)
		min := c[m][n-1]
		if min > c[m-1][n] {
			min = c[m-1][n]
		}
		if min > c[m-1][n-1] {
			min = c[m-1][n-1]
		}
		c[m][n] = 1 + min
		return
	}
}

func EditDistance(str1, str2 string) int {
	c := make(map[int]map[int]int, 0)
	for i := 0; i <= len(str1); i++ {
		c[i] = make(map[int]int, 0)
	}
	editDistance(str1, str2, c)
	return c[len(str1)][len(str2)]
}

func main() {
	fmt.Printf("%v\n", EditDistance("satur", "sun"))
}
