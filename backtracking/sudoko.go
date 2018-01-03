package main

import (
	"fmt"
)

var count int

func SudokuAux(m [][]int) bool {
	count++
	if IsSolved(m) {
		return true
	}
	var r, c int
loop:
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 0 {
				r = i
				c = j
				break loop
			}
		}
	}
	for k := 1; k <= 9; k++ {
		m[r][c] = k
		if IsSafe(m, r, c) {
			if SudokuAux(m) {
				return true
			}
		}
		m[r][c] = 0
	}
	return false
}

func IsSolved(m [][]int) bool {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func IsSafe(m [][]int, r, c int) bool {
	for i := 0; i < len(m); i++ {
		if i != r && m[i][c] == m[r][c] {
			return false
		}
		if i != c && m[r][i] == m[r][c] {
			return false
		}
	}
	R := r / 3
	C := c / 3
	for k := 0; k < 9; k++ {
		i := R*3 + k/3
		j := C*3 + k%3
		if i != r && j != c && m[i][j] == m[r][c] {
			return false
		}
	}
	return true
}

func main() {
	count = 0
	m := [][]int{
		[]int{3, 0, 6, 5, 0, 8, 4, 0, 0},
		[]int{5, 2, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 8, 7, 0, 0, 0, 0, 3, 1},
		[]int{0, 0, 3, 0, 1, 0, 0, 8, 0},
		[]int{9, 0, 0, 8, 6, 3, 0, 0, 5},
		[]int{0, 5, 0, 0, 9, 0, 6, 0, 0},
		[]int{1, 3, 0, 0, 0, 0, 2, 5, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 7, 4},
		[]int{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}
	res := SudokuAux(m)
	fmt.Printf("%v %v\n", count, res)
	for i := 0; i < len(m); i++ {
		fmt.Printf("%v \n", m[i])
	}
}
