package pathing

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_greedy(t *testing.T) {
	type args struct {
		start  tile
		target tile
		input  world
	}

	simpleWorld := populateWorld(t, simple)
	intermediateWorld := populateWorld(t, intermediate)
	forestWorld := populateWorld(t, forest)

	tests := []struct {
		name                 string
		args                 args
		expectedPath         []tile
		expectedNodesVisited int
	}{
		{
			name: "greedy search simple",
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
			expectedNodesVisited: 4,
		},
		{
			name: "greedy search simple - no path",
			args: args{
				start:  simpleWorld.tiles[0][0],
				target: simpleWorld.tiles[0][0],
				input:  simpleWorld,
			},
			expectedPath:         nil,
			expectedNodesVisited: 1,
		},
		{
			name: "greedy search intermediate",
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
			expectedNodesVisited: 7,
		},
		{
			name: "forest search intermediate - we go around the forest",
			args: args{
				start:  forestWorld.tiles[0][0],
				target: forestWorld.tiles[3][3],
				input:  forestWorld,
			},
			expectedPath: []tile{{
				row:   0,
				col:   1,
				value: 1,
			}, {
				row:   0,
				col:   2,
				value: 1,
			}, {
				row:   0,
				col:   3,
				value: 1,
			}, {
				row:   1,
				col:   3,
				value: 1,
			}, {
				row:   2,
				col:   3,
				value: 1,
			}, {
				row:   3,
				col:   3,
				value: 1,
			}},
			expectedNodesVisited: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, nodesVisited := greedy(tt.args.start, tt.args.target, tt.args.input)
			assert.Equal(t, tt.expectedPath, path)
			assert.Equal(t, tt.expectedNodesVisited, nodesVisited)
		})
	}
}
