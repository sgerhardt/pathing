package pathing

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func Test_printPath(t *testing.T) {
	type args struct {
		w world
		p []tile
	}
	tests := []struct {
		name           string
		args           args
		expectedOutput string
	}{
		{
			name:           "simple print",
			expectedOutput: "1 X \n3 X \n",
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
			name:           "forest print",
			expectedOutput: "1 X X X \n1 8 8 X \n1 8 8 X \n1 1 1 X \n",
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
			// Create a pipe to capture stdout.
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			originalStdout := os.Stdout
			defer func() {
				os.Stdout = originalStdout
				w.Close()
				r.Close()
			}()

			os.Stdout = w

			printPath(tt.args.w, tt.args.p)

			// Close the write end of the pipe to finish writing.
			w.Close()

			// Read the captured output from the read end of the pipe.
			var buf bytes.Buffer
			if _, err := io.Copy(&buf, r); err != nil {
				t.Fatal(err)
			}

			// Check if the output is what you expect.
			assert.Equal(t, tt.expectedOutput, buf.String())
		})
	}
}
