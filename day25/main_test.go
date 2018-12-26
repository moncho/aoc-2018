package main

import "testing"

func Test_constellations(t *testing.T) {
	type args struct {
		points []point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{
				[]point{
					point{0, 0, 0, 0},
					point{3, 0, 0, 0},
					point{0, 3, 0, 0},
					point{0, 0, 3, 0},
					point{0, 0, 0, 3},
					point{0, 0, 0, 6},
					point{9, 0, 0, 0},
					point{12, 0, 0, 0},
				},
			},
			2,
		},
		{
			"Example 2",
			args{
				[]point{
					point{-1, 2, 2, 0},
					point{0, 0, 2, -2},
					point{0, 0, 0, -2},
					point{-1, 2, 0, 0},
					point{-2, -2, -2, 2},
					point{3, 0, 2, -1},
					point{-1, 3, 2, 2},
					point{-1, 0, -1, 0},
					point{0, 2, 1, -2},
					point{3, 0, 0, 0},
				},
			},
			4,
		},
		{
			"Example 3",
			args{
				[]point{
					point{1, -1, 0, 1},
					point{2, 0, -1, 0},
					point{3, 2, -1, 0},
					point{0, 0, 3, 1},
					point{0, 0, -1, -1},
					point{2, 3, -2, 0},
					point{-2, 2, 0, 0},
					point{2, -2, 0, -1},
					point{1, -1, 0, -1},
					point{3, 2, 0, 2},
				},
			},
			3,
		},
		{
			"Example 4",
			args{
				[]point{
					point{1, -1, -1, -2},
					point{-2, -2, 0, 1},
					point{0, 2, 1, 3},
					point{-2, 3, -2, 1},
					point{0, 2, 3, -2},
					point{-1, -1, 1, -2},
					point{0, -2, -1, 0},
					point{-2, 2, 3, -1},
					point{1, 2, 2, 0},
					point{-1, -2, 0, -2},
				},
			},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constellations(tt.args.points); got != tt.want {
				t.Errorf("constellations() = %v, want %v", got, tt.want)
			}
		})
	}
}
