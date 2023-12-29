package pathing

func heuristic(a tile, b tile) int {
	return abs(a.row-b.row) + abs(a.col-b.col)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
