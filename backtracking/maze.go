package main

import "fmt"

/*
A Maze is given as N*N binary matrix of blocks where source block is the upper left most block i.e.,
maze[0][0] and destination block is lower rightmost block i.e., maze[N-1][N-1]. A rat starts from
source and has to reach destination. The rat can move only in two directions: forward and down.

In the maze matrix, 0 means the block is dead end and 1 means the block can be used in the path
from source to destination.
*/

func printSololution(sol [][]int) {
	for i := 0; i < len(sol); i++ {
		for j := 0; j < len(sol[i]); j++ {
			fmt.Printf(" %2d ", sol[i][j])
		}
		fmt.Printf("\n")
	}
}

func isSafe(i, j int, maze, sol [][]int) bool {
	N := len(sol)
	return i >= 0 && i < N && j >= 0 && j < N && maze[i][j] > 0 && sol[i][j] == 0
}

func MazeTravel(maze [][]int) {
	N := len(maze)
	sol := make([][]int, N)
	for i := 0; i < N; i++ {
		sol[i] = make([]int, N)
	}
	if MazeTravelAux(0, 0, maze, sol) {
		printSololution(sol)
	} else {
		fmt.Printf("No solution exist")
	}
}

func MazeTravelAux(i, j int, maze, sol [][]int) bool {
	N := len(sol)
	if i == N-1 && j == N-1 {
		sol[i][j] = 1
		return true
	}
	if isSafe(i, j, maze, sol) {
		sol[i][j] = 1
		if MazeTravelAux(i+1, j, maze, sol) {
			return true
		}

		if MazeTravelAux(i, j+1, maze, sol) {
			return true
		}
		sol[i][j] = 0
	}
	return false
}

func main() {
	maze := [][]int{
		[]int{1, 0, 0, 0},
		[]int{1, 1, 0, 1},
		[]int{0, 1, 0, 0},
		[]int{1, 1, 1, 1},
	}
	MazeTravel(maze)
}
