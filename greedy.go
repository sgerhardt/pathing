package pathing

import (
	"container/heap"
	"slices"
)

func greedy(start tile, target tile, w world) []tile {
	frontier := make(PriorityQueue, 0)
	heap.Init(&frontier)

	itemMap := make(map[tile]*Item)
	startItem := &Item{
		value:    start,
		priority: 0,
	}
	heap.Push(&frontier, startItem)
	itemMap[start] = startItem

	cameFrom := make(map[tile]tile)
	cameFrom[start] = start

	for frontier.Len() > 0 {
		currentItem := heap.Pop(&frontier).(*Item)
		current := currentItem.value

		if current == target {
			break
		}

		for _, next := range neighbors(current, w) {
			if _, ok := cameFrom[next]; !ok {
				cameFrom[next] = current
				priority := heuristic(next, target)
				if nextItem, ok := itemMap[next]; ok {
					frontier.UpdateItem(nextItem, next, priority)
				} else {
					nextItem = &Item{value: next, priority: priority}
					heap.Push(&frontier, nextItem)
					itemMap[next] = nextItem
				}
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

func heuristic(a tile, b tile) int {
	return abs(a.row-b.row) + abs(a.col-b.col)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
