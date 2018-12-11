package main

import (
	"reflect"
	"testing"
)

func Test_nextPosition(t *testing.T) {
	type args struct {
		currentPos int
		count      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			" -> (0)",
			args{
				0,
				0,
			},
			0,
		},
		{
			"(0) -> 0 (1)",
			args{
				0,
				1,
			},
			1,
		},
		{
			"0 (1) -> 0 (2) 1",
			args{
				1,
				2,
			},
			1,
		},
		{
			"0 (2) 1 ->	0  2  1 (3)",
			args{
				1,
				3,
			},
			3,
		},
		{
			"0 (2) 1 ->	0  2  1 (3)",
			args{
				1,
				3,
			},
			3,
		},
		{
			"0  2  1 (3) -> 0 (4) 2  1  3",
			args{
				3,
				4,
			},
			1,
		},
		{
			"0 (4) 2  1  3  -> 0  4  2 (5) 1  3",
			args{
				1,
				5,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPosition(tt.args.currentPos, tt.args.count); got != tt.want {
				t.Errorf("nextPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_highestScore(t *testing.T) {
	type args struct {
		players int
		plays   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"9 players, 25 points",
			args{
				9,
				25,
			},
			32,
		},
		{
			"10 players, 1618 points",
			args{
				10,
				1618,
			},
			8317,
		},
		{
			"13 players; last marble is worth 7999 points: high score is 146373",
			args{
				13,
				7999,
			},
			146373,
		},
		{
			"17 players; last marble is worth 1104 points: high score is 2764",
			args{
				17,
				1104,
			},
			2764,
		},
		{
			"21 players; last marble is worth 6111 points: high score is 54718",
			args{
				21,
				6111,
			},
			54718,
		},
		{
			"30 players; last marble is worth 5807 points: high score is 37305",
			args{
				30,
				5807,
			},
			37305,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highestScore(scores(tt.args.players, tt.args.plays)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("highestScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
