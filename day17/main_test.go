package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_scanToXy(t *testing.T) {
	type args struct {
		s scanLine
	}
	tests := []struct {
		name string
		args args
		want []xy
	}{
		{
			"x=495, y=2..7",
			args{
				s: scanLine{
					fromX: 495,
					toX:   495,
					fromY: 2,
					toY:   7,
				},
			},
			[]xy{
				xy{x: 495, y: 2},
				xy{x: 495, y: 3},
				xy{x: 495, y: 4},
				xy{x: 495, y: 5},
				xy{x: 495, y: 6},
				xy{x: 495, y: 7},
			},
		},
		{
			"y=7, x=495..501",
			args{
				s: scanLine{
					fromX: 495,
					toX:   501,
					fromY: 7,
					toY:   7,
				},
			},
			[]xy{
				xy{x: 495, y: 7},
				xy{x: 496, y: 7},
				xy{x: 497, y: 7},
				xy{x: 498, y: 7},
				xy{x: 499, y: 7},
				xy{x: 500, y: 7},
				xy{x: 501, y: 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scanToXy(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scanToXy() = %v, want %v", got, tt.want)
			}
		})
	}
}

const expected = `......+.......
......|.....#.
.#..#||||...#.
.#..#~~#|.....
.#..#~~#|.....
.#~~~~~#|.....
.#~~~~~#|.....
.#######|.....
........|.....
...|||||||||..
...|#~~~~~#|..
...|#~~~~~#|..
...|#~~~~~#|..
...|#######|..
`

func Test_flow(t *testing.T) {

	minX := 495
	maxX := 506
	minY := 1
	maxY := 13
	input := []scanLine{
		scanLine{fromX: 495, toX: 495, fromY: 2, toY: 7},
		scanLine{fromY: 7, toY: 7, fromX: 495, toX: 501},
		scanLine{fromX: 501, toX: 501, fromY: 3, toY: 7},
		scanLine{fromX: 498, toX: 498, fromY: 2, toY: 4},
		scanLine{fromX: 506, toX: 506, fromY: 1, toY: 2},
		scanLine{fromX: 498, toX: 498, fromY: 10, toY: 13},
		scanLine{fromX: 504, toX: 504, fromY: 10, toY: 13},
		scanLine{fromY: 13, toY: 13, fromX: 498, toX: 504},
	}

	var scan []xy
	for _, s := range input {
		scan = append(scan, scanToXy(s)...)
	}
	groundMap := newGroundMap(scan, minX, minY, maxX, maxY)
	x, y := groundMap.spring()
	simulateWaterFlow(x, y, groundMap)
	var b strings.Builder
	print(&b, groundMap)

	if b.String() != expected {
		t.Errorf("Unexpected state after water flow stabilizes, got:\n%s\n expected:\n%s", b.String(), expected)
	}

	wc, fc := groundMap.waterReach()
	if wc+fc != 57 {
		t.Errorf("Expeted water reach to be 57, got: %d, %d\n", wc, fc)
	}

}
