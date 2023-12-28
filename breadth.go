package pathing

import "slices"

type tile struct {
	row, col int
	value    int
}

type world struct {
	tiles [][]tile
}

func breadth(start tile, target tile, w world) []tile {
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
	return path
}

// neighbors returns the neighbors of the given tile and bounds checks
func neighbors(t tile, w world) []tile {
	var ns []tile
	// up
	if t.row > 0 {
		ns = append(ns, tile{
			row:   t.row - 1,
			col:   t.col,
			value: w.tiles[t.row-1][t.col].value,
		})
	}

	// right
	if t.col < len(w.tiles[0])-1 {
		ns = append(ns, tile{
			row:   t.row,
			col:   t.col + 1,
			value: w.tiles[t.row][t.col+1].value,
		})
	}

	// down
	if t.row < len(w.tiles)-1 {
		ns = append(ns, tile{
			row:   t.row + 1,
			col:   t.col,
			value: w.tiles[t.row+1][t.col].value,
		})
	}

	// left
	if t.col > 0 {
		ns = append(ns, tile{
			row:   t.row,
			col:   t.col - 1,
			value: w.tiles[t.row][t.col-1].value,
		})
	}

	return ns
}
