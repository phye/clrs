package graph

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMinPriorityQueue(t *testing.T) {
	l := 10
	pq := &MinPriorityQueue{make([]*Item, 10)}
	for i := 0; i < l; i++ {
		pq.items[i] = &Item{10 - i, 10 - i, i}
	}
	fmt.Printf("Before Priority Queue init, pq is \n")
	for i := 0; i < len(pq.items); i++ {
		fmt.Printf("%v ", pq.items[i])
	}
	fmt.Println()

	heap.Init(pq)
	fmt.Printf("After Priority Queue init, pq is \n")
	for i := 0; i < len(pq.items); i++ {
		fmt.Printf("%v ", pq.items[i])
	}
	fmt.Println()

	min := heap.Pop(pq)
	fmt.Printf("Minimum is %v\n", min)

	fmt.Printf("After Priority Queue Pop, pq is \n")
	for i := 0; i < len(pq.items); i++ {
		fmt.Printf("%v ", pq.items[i])
	}
	fmt.Println()
}
