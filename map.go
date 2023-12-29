package pathing

type tile struct {
	row, col int
	value    int
}

type world struct {
	tiles [][]tile
}
