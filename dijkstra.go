package pathing

import (
	"container/heap"
	"slices"
)

// dijkstra algorithm with priority queue
func dijkstra(start tile, target tile, w world) ([]tile, int) {
	frontier := make(PriorityQueue, 0)
	heap.Init(&frontier)

	startItem := &Entry{
		value:    start,
		priority: 0,
	}
	heap.Push(&frontier, startItem)

	cameFrom := make(map[tile]tile)
	cameFrom[start] = start

	costSoFar := make(map[tile]int)
	costSoFar[start] = 0

	for frontier.Len() > 0 {
		currentItem := heap.Pop(&frontier).(*Entry)
		current := currentItem.value

		if current == target {
			break
		}

		for _, next := range neighbors(current, w) {
			newCost := costSoFar[current] + next.value
			if _, ok := costSoFar[next]; !ok || newCost < costSoFar[next] {
				costSoFar[next] = newCost
				priority := newCost
				heap.Push(&frontier, &Entry{value: next, priority: priority})
				cameFrom[next] = current
			}
		}
	}
	current := target
	var path []tile

	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	slices.Reverse(path)
	return path, len(cameFrom)
}
