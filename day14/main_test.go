package main

import (
	"reflect"
	"testing"
)

func Test_scores(t *testing.T) {
	type args struct {
		score1   int
		score2   int
		attempts int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"start with scores 3,7 - 9 attempts",
			args{
				3,
				7,
				9,
			},
			[]int{5, 1, 5, 8, 9, 1, 6, 7, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scores(tt.args.score1, tt.args.score2, tt.args.attempts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scores() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_index(t *testing.T) {
	type args struct {
		scores   []int
		expected []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				[]int{6, 5, 1, 5, 8, 9},
				[]int{5, 1, 5, 8, 9},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := index(tt.args.scores, tt.args.expected); got != tt.want {
				t.Errorf("index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2Scores(t *testing.T) {
	type args struct {
		score1   int
		score2   int
		expected int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"start with scores 3,7 - sequence 51589",
			args{
				3,
				7,
				51589,
			},

			[]int{3, 7, 1, 0, 1, 0, 1, 2, 4},
		},
		{
			"start with scores 3,7 - sequence 92510",
			args{
				3,
				7,
				92510,
			},
			[]int{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1, 5, 8, 9, 1, 6, 7, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2Scores(tt.args.score1, tt.args.score2, tt.args.expected); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part2Scores() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2Scores_resultCount(t *testing.T) {
	type args struct {
		score1   int
		score2   int
		expected int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"start with scores 3,7 - sequence 909441",
			args{
				3,
				7,
				909441,
			},
			20403320,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part2Scores(tt.args.score1, tt.args.score2, tt.args.expected)

			if len(got) != tt.want {
				t.Errorf("part2Scores() count = %v, want %v", got, tt.want)
			}
		})
	}
}
