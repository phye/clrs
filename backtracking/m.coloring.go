package main

import (
	"fmt"
)

type (
	Color int
)

const (
	TRANS Color = 0
)

func Coloring(mg [][]int, m Color) ([]Color, bool) {
	ret := make([]Color, 0)
	if ColoringAux(mg, m, &ret) {
		return ret, true
	} else {
		return ret, false
	}
}

func ColoringAux(mg [][]int, m Color, ret *[]Color) bool {
	if len(*ret) == len(mg) {
		return true
	}
	var c Color
	for c = 1; c <= m; c++ {
		*ret = append(*ret, c)
		if IsSafe(mg, *ret) {
			if ColoringAux(mg, m, ret) {
				return true
			}
		}
		*ret = (*ret)[:len(*ret)-1]
	}
	return false
}

func IsSafe(mg [][]int, col []Color) bool {
	for i := 0; i < len(mg); i++ {
		for j := 0; j < len(mg[i]); j++ {
			if mg[i][j] > 0 {
				if i < len(col) && j < len(col) {
					if col[i] == col[j] {
						return false
					}
				}
			}
		}
	}
	return true
}

func main() {
	mg := [][]int{
		[]int{0, 1, 1, 1},
		[]int{1, 0, 1, 0},
		[]int{1, 1, 0, 1},
		[]int{1, 0, 1, 0},
	}
	m := Color(3)
	ret, _ := Coloring(mg, m)
	fmt.Printf("%v\n", ret)
}
