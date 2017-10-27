package main

import (
	"fmt"
)

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func maxHeapify(A []int, i int, hpsize int) {
	l := left(i)
	r := right(i)
	var largest int
	if l < hpsize && A[l] > A[i] {
		largest = l
	} else {
		largest = i
	}
	if r < hpsize && A[r] > A[largest] {
		largest = r
	}
	if largest != i {
		tmp := A[i]
		A[i] = A[largest]
		A[largest] = tmp
		maxHeapify(A, largest, hpsize)
	}
}

// Note that building max heap is of time O(n)
func buildMaxHeap(A []int) {
	hpsize := len(A)
	for i := len(A)/2 - 1; i >= 0; i-- {
		maxHeapify(A, i, hpsize)
	}
}

func heapSort(A []int) {
	hpsize := len(A)
	buildMaxHeap(A)
	fmt.Printf("%2v\n", A)
	for i := len(A) - 1; i > 0; i-- {
		tmp := A[0]
		A[0] = A[i]
		A[i] = tmp
		hpsize = hpsize - 1
		maxHeapify(A, 0, hpsize)
		fmt.Printf("%2v\n", A)
	}
}

func main() {
	A := []int{5, 13, 2, 25, 7, 17, 20, 8, 4}
	fmt.Printf("%2v\n", A)
	heapSort(A)
	fmt.Printf("%2v\n", A)
}
