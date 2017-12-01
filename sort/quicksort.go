package main

import (
	"fmt"
)

func quickSort(A []int, p int, r int) {
	if p < r {
		q := partition(A, p, r)
		quickSort(A, p, q-1)
		quickSort(A, q+1, r)
	}
}

func iterativeQuickSort(A []int, p int, r int) {
	stack := []int{}
	stack = append(stack, p, r)
	for {
		if len(stack) == 0 {
			break
		}
		p = stack[len(stack)-2]
		r = stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		q := partition(A, p, r)
		// Push rhs part to stack first
		if q+1 < r {
			stack = append(stack, q+1, r)
		}
		if p < q-1 {
			stack = append(stack, p, q-1)
		}
		fmt.Printf("%v\n", stack)
	}
}

func partition(A []int, p int, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j] < x {
			i = i + 1
			tmp := A[j]
			A[j] = A[i]
			A[i] = tmp
		}
	}
	tmp := A[i+1]
	A[i+1] = A[r]
	A[r] = tmp
	return i + 1
}

func main() {
	A := []int{11, 3, 7, 5, 8, 9, 20}
	fmt.Printf("%v\n", A)
	iterativeQuickSort(A, 0, len(A)-1)
	fmt.Printf("%v\n", A)
}
