package main

import (
	"fmt"
)

const (
	INF int = 10000
)

func MinJump(A []int) int {
	// c[i] denotes minimum jump from A[0] to A[i]
	c := make(map[int]int, 0)
	if len(A) == 0 || A[0] == 0 {
		return INF
	}
	c[0] = 0
	minJump(A, len(A)-1, c)
	return c[len(A)-1]
}

func minJump(A []int, i int, c map[int]int) {
	if _, ok := c[i]; ok {
		return
	}
	c[i] = INF
	for j := 0; j < i; j++ {
		minJump(A, j, c)
		if A[j]+j >= i {
			if c[j]+1 < c[i] {
				c[i] = c[j] + 1
				//break
			}
		}
	}
}

func main() {
	fmt.Printf("%v\n", MinJump([]int{1, 3, 6, 3, 2, 3, 6, 8, 9, 5}))
}
