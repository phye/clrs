package ds

import (
	"fmt"
	"testing"
)

func TestMinPriorityQueue(t *testing.T) {
	l := 10
	mpq := NewMinPriorityQueue()
	for i := 0; i < l; i++ {
		mpq.Push(10-i, 10-i)
	}
	fmt.Printf("Priority Queue is \n")
	for i := 0; i < len(mpq.pq.items); i++ {
		fmt.Printf("%v ", mpq.pq.items[i])
	}
	fmt.Println()

	min := mpq.Pop()
	fmt.Printf("Minimum is %v\n", min)

	fmt.Printf("After Priority Queue Pop, pq is \n")
	for i := 0; i < len(mpq.pq.items); i++ {
		fmt.Printf("%v ", mpq.pq.items[i])
	}
	fmt.Println()
}
