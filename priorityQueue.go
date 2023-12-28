package pathing

import "container/heap"

type PriorityQueue []*Item

type Item struct {
	value    tile // The value of the item; arbitrary.
	priority int  // The priority of the item in the queue.
	index    int  // The index of the item in the heap.
}

func (pq *PriorityQueue) Len() int { return len(*pq) }

// Less reports whether the item with index i should sort before the item with index j.
func (pq *PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest priority so we use less than here.
	return (*pq)[i].priority < (*pq)[j].priority
}

// Swap swaps the items with indexes i and j.
func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// UpdateItem updates the priority of an item in the queue.
func (pq *PriorityQueue) UpdateItem(item *Item, value tile, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
