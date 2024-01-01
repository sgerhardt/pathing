package pathing

import "fmt"

type coordinates struct {
	row int
	col int
}

func printPath(w world, p []tile) {
	if len(p) == 0 || len(w.tiles) == 0 {
		return
	}

	var pathMap = make(map[coordinates]bool)
	for _, t := range p {
		pathMap[coordinates{t.col, t.row}] = true
	}

	for _, r := range w.tiles {
		for _, c := range r {
			if pathMap[coordinates{c.col, c.row}] {
				fmt.Print("X")
			} else {
				fmt.Print(c.value)
			}
			fmt.Print(" ")
			if c.col == len(r)-1 {
				fmt.Print("\n")
			}
			if len(p) == 0 {
				return
			}
		}
	}
}
