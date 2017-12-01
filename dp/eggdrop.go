package main

import (
	"fmt"
)

const (
	INF int = 10000
)

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// Return the minimum number of trials in worst case
// k: number of floors, n: number of eggs
func EggDrop(k int, n int) int {
	if k == 1 || k == 0 {
		return k
	}
	if n == 1 {
		return k
	}
	min := INF
	for i := 1; i <= k; i++ {
		low := EggDrop(i-1, n-1)
		high := EggDrop(k-i, n)
		res := max(low, high)
		if res < min {
			min = res
		}
	}
	return min + 1
}

func EggDropMemoized(k int, n int) (int, [][]int) {
	mem := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		mem[i] = make([]int, n+1)
	}
	choices := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		choices[i] = make([]int, n+1)
	}
	eggDropMemoized(k, n, mem, choices)
	return mem[k][n], choices
}

func eggDropMemoized(k int, n int, mem [][]int, choices [][]int) {
	if mem[k][n] > 0 {
		return
	}
	if k == 0 || k == 1 {
		mem[k][n] = k
		return
	}
	if n == 1 {
		mem[k][n] = k
		return
	}
	min := INF
	for i := 1; i <= k; i++ {
		eggDropMemoized(i-1, n-1, mem, choices)
		eggDropMemoized(k-i, n, mem, choices)
		res := max(mem[i-1][n-1], mem[k-i][n])
		if res < min {
			min = res
			choices[k][n] = i
		}
	}
	mem[k][n] = min + 1
	return
}

func main() {
	k := 36
	n := 2
	num, choices := EggDropMemoized(k, n)
	for i := 0; i <= k; i++ {
		//fmt.Printf("%v\n", choices[i])
	}
	fmt.Printf("%v \n", num)
	i := 0
	for {
		if k-i <= 1 {
			break
		}
		i += choices[k-i][n]
		fmt.Printf("%v ", i)
	}
	fmt.Printf("%v \n", k)
}
