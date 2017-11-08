package main

import (
	"fmt"
)

// Calculate the edit distance of two strings with CUD operations

func EditDistance(str1, str2 string) int {
	m := len(str1)
	n := len(str2)
	if m == 0 {
		return n
	} else if n == 0 {
		return m
	}
	if str1[m-1] == str2[n-1] {
		return EditDistance(str1[:m-1], str2[:n-1])
	} else {
		c := 1 + EditDistance(str1, str2[:n-1])
		d := 1 + EditDistance(str1[:m-1], str2)
		u := 1 + EditDistance(str1[:m-1], str2[:n-1])
		min := c
		if min < d {
			min = d
		} else if min < u {
			min = u
		}
		return min
	}
}

func main() {
	fmt.Printf("%v", EditDistance("sunday", "saturday"))
}
