package main

import (
	"strings"
	"testing"

	"github.com/beefsack/go-astar"
)

const straightLineMap = `#######
#.E..G#
#######`

const blockedStraightLineMap = `#######
#.E##G#
#.....#
#######`

const blockedStraightLineMap2 = `#######
#.....#
#.E##G#
#.....#
#######`

const circularPathMap = `#######
#.E##G#
#.EEE.#
#.....#
#######`

func TestTile_PathFinder(t *testing.T) {
	type args struct {
		game *game
		from xy
		to   xy
	}
	tests := []struct {
		name     string
		args     args
		found    bool
		distance float64
	}{
		{
			"straightLineMap",
			args{
				game: newGame(strings.NewReader(straightLineMap)),
				from: xy{2, 1},
				to:   xy{5, 1},
			},
			true,
			3.0,
		},
		{
			"blockedStraightLineMap",
			args{
				game: newGame(strings.NewReader(blockedStraightLineMap)),
				from: xy{2, 1},
				to:   xy{5, 1},
			},
			true,
			5.0,
		},
		{
			"blockedStraightLineMap2",
			args{
				game: newGame(strings.NewReader(blockedStraightLineMap2)),
				from: xy{2, 2},
				to:   xy{5, 2},
			},
			true,
			5.0,
		},
		{
			"circularPathMap",
			args{
				game: newGame(strings.NewReader(circularPathMap)),
				from: xy{2, 1},
				to:   xy{5, 1},
			},
			true,
			9.0,
		},
		{
			"circularPathMap reverse",
			args{
				game: newGame(strings.NewReader(circularPathMap)),
				from: xy{5, 1},
				to:   xy{2, 1},
			},
			true,
			9.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := tt.args.game
			pf := newPathFinder(game.battleMap, tt.args.from, tt.args.to)
			path, distance, found := astar.Path(pf.From(), pf.To())
			_ = len(path)
			if found != tt.found {
				t.Error("No path  found")
			}
			if distance != tt.distance {
				t.Errorf("Unexpected distance, wanted %f, got %f", tt.distance, distance)
			}
		})
	}
}
