package pathing

import "fmt"

func printPath(w world, p []tile) {
	for _, r := range w.tiles {
		for _, c := range r {
			if c.row == p[0].row && c.col == p[0].col {
				fmt.Print("X")
				if len(p) > 1 {
					p = p[1:]
				}
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
