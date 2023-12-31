package pathing

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
