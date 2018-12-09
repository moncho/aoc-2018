package main

import (
	"reflect"
	"testing"
)

func Test_convexHull(t *testing.T) {
	type args struct {
		coords []coord
	}
	coords := make([]coord, 100)
	for i := 0; i < 100; i++ {
		coords[i] =
			coord{
				x: i / 10,
				y: i % 10,
			}
	}
	tests := []struct {
		name string
		args args
		want []coord
	}{
		{
			"convex hull of a 10-by-10 grid.",
			args{
				coords,
			},
			[]coord{
				coord{x: 0, y: 0},
				coord{x: 9, y: 0},
				coord{x: 9, y: 9},
				coord{x: 0, y: 9},
			},
		},
		{
			"convex hull test",
			args{
				[]coord{
					coord{x: 0, y: 3},
					coord{x: 2, y: 2},
					coord{x: 1, y: 1},
					coord{x: 2, y: 1},
					coord{x: 3, y: 0},
					coord{x: 0, y: 0}},
			},
			[]coord{
				coord{x: 0, y: 0},
				coord{x: 3, y: 0},
				coord{x: 2, y: 2},
				coord{x: 0, y: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convexHull(tt.args.coords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convexHull() = %v, want %v", got, tt.want)
			}
		})
	}
}
