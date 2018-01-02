package main

import "fmt"

func isSafe(x, y int, sol [][]int) bool {
	N := len(sol)
	return x >= 0 && x < N && y >= 0 && y < N && sol[x][y] == -1
}

func printSolution(sol [][]int) {
	for i := 0; i < len(sol); i++ {
		for j := 0; j < len(sol[i]); j++ {
			fmt.Printf(" %2d ", sol[i][j])
		}
		fmt.Printf("\n")
	}
}

func solveKT(N int) bool {
	sol := make([][]int, N)
	for i := 0; i < N; i++ {
		sol[i] = make([]int, N)
		for j := 0; j < N; j++ {
			sol[i][j] = -1
		}
	}
	xMove := []int{2, 1, -1, -2, -2, -1, 1, 2}
	yMove := []int{1, 2, 2, 1, -1, -2, -2, -1}
	sol[0][0] = 0
	if !solveKTAux(0, 0, 1, sol, xMove, yMove) {
		fmt.Printf("Solution does not exist")
		return false
	} else {
		printSolution(sol)
	}
	return true
}

func solveKTAux(x, y int, movei int, sol [][]int, xMove, yMove []int) bool {
	var nx, ny int
	N := len(sol)
	if movei == N*N {
		return true
	}
	for k := 0; k < 8; k++ {
		nx = x + xMove[k]
		ny = y + yMove[k]
		if isSafe(nx, ny, sol) {
			sol[nx][ny] = movei
			if solveKTAux(nx, ny, movei+1, sol, xMove, yMove) {
				return true
			} else {
				sol[nx][ny] = -1
			}
		}
	}
	return false
}

func main() {
	solveKT(9)
}
