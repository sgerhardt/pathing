package pathing

import "testing"

func Test_printPath(t *testing.T) {
	type args struct {
		w world
		p []tile
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "simple print",
			args: args{
				w: populateWorld(t, simple),
				p: []tile{{
					row:   0,
					col:   1,
					value: 2,
				}, {
					row:   1,
					col:   1,
					value: 4,
				}},
			},
		}, {
			name: "forest print",
			args: args{
				w: populateWorld(t, forest),
				p: []tile{{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printPath(tt.args.w, tt.args.p)
		})
	}
}
