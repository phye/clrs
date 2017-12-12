package ds

import (
	"container/heap"
)

type (
	Item struct {
		value    interface{}
		priority int
		index    int
	}
	minpq struct {
		items []*Item
	}
	// Thread-Unsafe Minimum Priority Queue
	MinPriorityQueue struct {
		pq *minpq
	}
)

// Required by heap.(sort).Interface
func (pq *minpq) Len() int {
	return len(pq.items)
}

// Required by heap.(sort).Interface
func (pq *minpq) Less(i, j int) bool {
	return pq.items[i].priority < pq.items[j].priority
}

// Required by heap.(sort).Interface
func (pq *minpq) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// Required by heap.Interface
func (pq *minpq) Push(x interface{}) {
	n := len(pq.items)
	item := x.(*Item)
	item.index = n
	pq.items = append(pq.items, item)
}

// Required by heap.Interface
func (pq *minpq) Pop() interface{} {
	item := pq.items[len(pq.items)-1]
	item.index = -1
	pq.items = pq.items[:len(pq.items)-1]
	return item
}

func NewMinPriorityQueue() *MinPriorityQueue {
	pq := &minpq{
		make([]*Item, 0),
	}
	heap.Init(pq)

	mpq := &MinPriorityQueue{pq}
	return mpq
}

// Push new value /x/ with priority /priority/ to the Min Priority Queue
func (mpq *MinPriorityQueue) Push(x interface{}, priority int) {
	heap.Push(mpq.pq, &Item{
		x,
		priority,
		mpq.pq.Len(),
	})
}

// Pop the item with minimum priority
func (mpq *MinPriorityQueue) Pop() interface{} {
	ret := heap.Pop(mpq.pq)
	return ret
}

// Update the item with new priority
func (mpq *MinPriorityQueue) Update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(mpq.pq, item.index)
}
