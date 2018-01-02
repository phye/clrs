package main

import "fmt"

func printSololution(sol [][]int) {
	for i := 0; i < len(sol); i++ {
		for j := 0; j < len(sol[i]); j++ {
			fmt.Printf(" %2d ", sol[i][j])
		}
		fmt.Printf("\n")
	}
}

func isSafe(i, j int, sol [][]int) bool {
	N := len(sol)
	if i >= 0 && i < N && j >= 0 && j < N {
		// Check the same coloumn
		for x := 0; x < i; x++ {
			if sol[x][j] > 0 {
				return false
			}
		}
		// Check Left Top
		for x, y := i, j; x >= 0 && y >= 0; x, y = x-1, y-1 {
			if sol[x][y] > 0 {
				return false
			}
		}
		// Check Right Top
		for x, y := i, j; x >= 0 && y < N; x, y = x-1, y+1 {
			if sol[x][y] > 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func NQueen(N int) {
	sol := make([][]int, N)
	for i := 0; i < N; i++ {
		sol[i] = make([]int, N)
	}
	if NQueenAux(0, sol) {
		printSololution(sol)
	} else {
		fmt.Printf("No solution exist")
	}
}

func NQueenAux(r int, sol [][]int) bool {
	N := len(sol)
	if r == N {
		return true
	}
	for c := 0; c < len(sol[r]); c++ {
		if isSafe(r, c, sol) {
			sol[r][c] = 1
			if NQueenAux(r+1, sol) {
				return true
			} else {
				sol[r][c] = 0
			}
		}
	}
	return false
}

func main() {
	NQueen(8)
}
