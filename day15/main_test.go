package main

import (
	"strings"
	"testing"
)

const gameMap1 = `#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`

const gameMap2 = `#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`

const gameMap3 = `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`

const gameMap4 = `#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`

const gameMap5 = `#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`

func Test_play(t *testing.T) {
	type args struct {
		game *game
	}
	tests := []struct {
		name  string
		args  args
		turns int
		score int
	}{

		{
			"gameMap1",
			args{
				game: newGame(strings.NewReader(gameMap1)),
			},
			54,
			536,
		},
		{
			"gameMap2",
			args{
				game: newGame(strings.NewReader(gameMap2)),
			},
			46,
			859,
		},
		{
			"gameMap3",
			args{
				game: newGame(strings.NewReader(gameMap3)),
			},
			46,
			590,
		},
		{
			"gameMap4",
			args{
				game: newGame(strings.NewReader(gameMap4)),
			},
			35,
			793,
		},
		{
			"gameMap5",
			args{
				game: newGame(strings.NewReader(gameMap5)),
			},
			37,
			982,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			turns := play(tt.args.game)
			score := tt.args.game.score()

			if turns != tt.turns {
				t.Errorf("turns after game finished = %v, want %v", turns, tt.turns)
			}
			if score != tt.score {
				t.Errorf("score after game finished = %v, want %v", score, tt.score)
			}
		})
	}
}
