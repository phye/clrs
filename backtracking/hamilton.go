package main

import (
	"fmt"
)

func Hamilton(mg [][]int) ([]int, bool) {
	ret := make([]int, 0)
	if HamiltonAux(mg, &ret) {
		return ret, true
	} else {
		return ret, false
	}
}

func HamiltonAux(mg [][]int, ret *[]int) bool {
	if len(*ret) == len(mg)+1 {
		return true
	}
	for i := 0; i < len(mg); i++ {
		*ret = append(*ret, i)
		if IsSafe(mg, *ret) {
			if HamiltonAux(mg, ret) {
				return true
			}
		}
		*ret = (*ret)[:len(*ret)-1]
	}
	return false
}

func IsSafe(mg [][]int, path []int) bool {
	if len(path) <= 1 {
		return true
	}
	n := len(path)
	if n < len(mg)+1 {
		for i := 0; i < n-1; i++ {
			if path[i] == path[n-1] {
				return false
			}
		}
	} else {
		if path[len(path)-1] != path[0] {
			return false
		}
	}
	if mg[path[n-2]][path[n-1]] == 0 {
		return false
	}
	return true
}

func main() {
	mg := [][]int{
		[]int{0, 1, 0, 1, 0},
		[]int{1, 0, 1, 1, 1},
		[]int{0, 1, 0, 0, 1},
		[]int{1, 1, 0, 0, 1},
		[]int{0, 1, 1, 1, 0},
	}
	ret, b := Hamilton(mg)
	fmt.Printf("%v %v\n", ret, b)

	mg = [][]int{
		[]int{0, 1, 0, 1, 0},
		[]int{1, 0, 1, 1, 1},
		[]int{0, 1, 0, 0, 1},
		[]int{1, 1, 0, 0, 0},
		[]int{0, 1, 1, 0, 0},
	}
	ret, b = Hamilton(mg)
	fmt.Printf("%v, %v\n", ret, b)
}
