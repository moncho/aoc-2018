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
