package main

import (
	"testing"

	"github.com/beefsack/go-astar"
)

func TestTile_PathFinder(t *testing.T) {
	type args struct {
		cave *cave
	}
	tests := []struct {
		name     string
		args     args
		found    bool
		distance float64
	}{

		{
			"Target at 10,10 depth 510",
			args{
				cave: &cave{
					depth:           510,
					targetX:         10,
					targetY:         10,
					geologicIndices: make(map[xy]int),
				},
			},
			true,
			45.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cave := tt.args.cave
			e := rescueMission(cave)
			p, distance, found := astar.Path(e, explorer{
				cave: cave,
				tool: torch,
				x:    cave.targetX,
				y:    cave.targetY,
			})
			_ = p
			if found != tt.found {
				t.Error("No path  found")
			}
			if distance != tt.distance {
				t.Errorf("Unexpected distance, wanted %f, got %f", tt.distance, distance)
			}
		})
	}
}
