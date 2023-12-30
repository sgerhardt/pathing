package pathing

import "slices"

func breadth(start tile, target tile, w world) ([]tile, int) {
	frontier := []tile{start}

	cameFrom := make(map[tile]tile)
	cameFrom[start] = start

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]
		if current == target {
			break
		}
		for _, next := range neighbors(current, w) {
			if _, ok := cameFrom[next]; !ok {
				frontier = append(frontier, next)
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
