package pathing

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func Test_breadth(t *testing.T) {
	type args struct {
		start  tile
		target tile
		input  world
	}

	simpleWorld := populateWorld(t, simple)
	intermediateWorld := populateWorld(t, intermediate)

	tests := []struct {
		name                 string
		args                 args
		expectedPath         []tile
		expectedNodesVisited int
	}{
		{
			name: "breadth search simple",
			args: args{
				start:  simpleWorld.tiles[0][0],
				target: simpleWorld.tiles[1][1],
				input:  simpleWorld,
			},
			expectedPath: []tile{{
				row:   0,
				col:   1,
				value: 2,
			}, {
				row:   1,
				col:   1,
				value: 4,
			}},
			expectedNodesVisited: len(simpleWorld.tiles) * len(simpleWorld.tiles[0]),
		},
		{
			name: "breadth search simple - no path",
			args: args{
				start:  simpleWorld.tiles[0][0],
				target: simpleWorld.tiles[0][0],
				input:  simpleWorld,
			},
			expectedPath:         nil,
			expectedNodesVisited: 1,
		},
		{
			name: "breadth search intermediate",
			args: args{
				start:  intermediateWorld.tiles[0][0],
				target: intermediateWorld.tiles[2][1],
				input:  intermediateWorld,
			},
			expectedPath: []tile{{
				row:   0,
				col:   1,
				value: 2,
			}, {
				row:   1,
				col:   1,
				value: 5,
			}, {
				row:   2,
				col:   1,
				value: 8,
			}},
			expectedNodesVisited: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, nodesVisited := breadth(tt.args.start, tt.args.target, tt.args.input)
			assert.Equal(t, tt.expectedPath, path)
			assert.Equal(t, tt.expectedNodesVisited, nodesVisited)
		})
	}
}

func populateWorld(t *testing.T, input string) world {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	w := make([][]tile, len(rows))
	for i, row := range rows {
		cols := strings.Split(row, ",")
		w[i] = make([]tile, len(cols))
		for j, col := range cols {
			num, err := strconv.Atoi(strings.TrimSpace(col))
			if err != nil {
				t.Fatalf("Failed to parse integer: %v", err)
			}
			w[i][j] = tile{
				row:   i,
				col:   j,
				value: num,
			}
		}
	}

	return world{tiles: w}
}
