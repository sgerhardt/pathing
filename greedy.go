package pathing

import (
	"container/heap"
	"slices"
)

func greedy(start tile, target tile, w world) []tile {
	frontier := make(PriorityQueue, 0)
	heap.Init(&frontier)

	startItem := &Entry{
		value:    start,
		priority: 0,
	}
	heap.Push(&frontier, startItem)

	cameFrom := make(map[tile]tile)
	cameFrom[start] = start

	for frontier.Len() > 0 {
		currentItem := heap.Pop(&frontier).(*Entry)
		current := currentItem.value

		if current == target {
			break
		}

		for _, next := range neighbors(current, w) {
			if _, ok := cameFrom[next]; !ok {
				priority := heuristic(next, target)
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
	return path
}
