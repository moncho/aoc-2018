package main

import (
	"strings"
	"testing"
)

var test = `.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.`

func Test_gameOfLife(t *testing.T) {
	type args struct {
		duration  int
		landscape [][]rune
	}
	tests := []struct {
		name  string
		args  args
		value int
	}{
		{
			"test",
			args{
				duration:  10,
				landscape: newLandscape(10, strings.NewReader(test)),
			},
			1147,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := simulate(tt.args.duration, tt.args.landscape)
			value := resourcesValue(got)
			if value != tt.value {
				t.Errorf("gameOfLife() value = %d, want %d", value, tt.value)
			}
		})
	}
}
