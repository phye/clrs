package graph

import (
	"container/heap"
)

type (
	Item struct {
		value    interface{}
		priority int
		index    int
	}
	MinPriorityQueue struct {
		items []*Item
	}
)

func (pq *MinPriorityQueue) Len() int {
	return len(pq.items)
}

func (pq *MinPriorityQueue) Less(i, j int) bool {
	return pq.items[i].priority < pq.items[j].priority
}

func (pq *MinPriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *MinPriorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(*Item)
	item.index = n
	pq.items = append(pq.items, item)
}

func (pq *MinPriorityQueue) Pop() interface{} {
	item := pq.items[len(pq.items)-1]
	item.index = -1
	pq.items = pq.items[:len(pq.items)-1]
	return item
}

func (pq *MinPriorityQueue) Update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
