package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
func Binomial(n, k int) int {
	if k > n {
		return 0
	}
	c := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		c[i] = make([]int, k+1)
		c[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, k); j++ {
			if j == i {
				c[i][j] = 1
			} else {
				c[i][j] = c[i-1][j-1] + c[i-1][j]
			}
		}
	}
	return c[n][k]
}

func main() {
	fmt.Printf("%v\n", Binomial(10, 2))
}
