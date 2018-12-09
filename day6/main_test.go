package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func Test_generateAreaMap(t *testing.T) {
	type args struct {
		coords []coord
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want [][]*distance
	}{
		{
			"2x2 grid, distances to 0,0",
			args{
				coords: []coord{
					coord{id: "A", x: 0, y: 0},
				},
				width:  2,
				height: 2,
			},
			[][]*distance{
				{&distance{
					"A", 0,
				}, &distance{
					"A", 1,
				}},
				{&distance{
					"A", 1,
				}, &distance{
					"A", 2,
				}},
			},
		},
		{
			"3x3 grid, distances to (0,0), (2,2)",
			args{
				coords: []coord{
					coord{id: "A", x: 0, y: 0},
					coord{id: "C", x: 2, y: 2},
				},
				width:  3,
				height: 3,
			},
			[][]*distance{
				{
					&distance{
						"A", 0,
					},
					&distance{
						"A", 1,
					},
					&distance{
						".", 2,
					},
				},
				{
					&distance{
						"A", 1,
					},
					&distance{
						".", 2,
					},
					&distance{
						"C", 1,
					},
				},
				{
					&distance{
						".", 2,
					},
					&distance{
						"C", 1,
					},
					&distance{
						"C", 0,
					},
				},
			},
		},
		{
			"4x2 grid, distances to (0,0), (3,1)",
			args{
				coords: []coord{
					coord{id: "A", x: 0, y: 0},
					coord{id: "C", x: 3, y: 1},
				},
				width:  4,
				height: 2,
			},
			[][]*distance{
				{
					&distance{
						"A", 0,
					},
					&distance{
						"A", 1,
					},
					&distance{
						".", 2,
					},
					&distance{
						"C", 1,
					},
				},
				{
					&distance{
						"A", 1,
					},
					&distance{
						".", 2,
					},
					&distance{
						"C", 1,
					},
					&distance{
						"C", 0,
					},
				},
			},
		},

		{
			"2x4 grid, distances to (0,0), (0,1), (1,3)",
			args{
				coords: []coord{
					coord{id: "A", x: 0, y: 0},
					coord{id: "B", x: 0, y: 1},
					coord{id: "C", x: 1, y: 3},
				},
				width:  2,
				height: 4,
			},
			[][]*distance{
				{
					&distance{
						"A", 0,
					},
					&distance{
						"A", 1,
					},
				},
				{
					&distance{
						"B", 0,
					},
					&distance{
						"B", 1,
					},
				},
				{
					&distance{
						"B", 1,
					},
					&distance{
						"C", 1,
					},
				},
				{
					&distance{
						"C", 1,
					},
					&distance{
						"C", 0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateAreaMap(tt.args.coords, tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateAreaMap() = \n%s\n, want \n%s\n",
					printDistanceGrid(got), printDistanceGrid(tt.want))
			}
		})
	}
}

func printDistanceGrid(distances [][]*distance) string {
	var result bytes.Buffer

	for _, row := range distances {
		fmt.Fprint(&result, "\t")
		for _, d := range row {
			fmt.Fprintf(&result, " %s-%d ", d.coordID, d.distance)
		}
		fmt.Fprint(&result, "\n")
	}
	return result.String()
}
