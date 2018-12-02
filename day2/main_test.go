package main

import (
	"testing"
)

func Test_repetitions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 bool
	}{
		{
			"abcdef contains no letters that appear exactly two or three times.",
			args{
				"abcdef",
			},
			false, false,
		},
		{
			"bababc contains two a and three b, so it counts for both.",
			args{
				"bababc",
			},
			true, true,
		},
		{
			"abbcde contains two b, but no letter appears exactly three times.",
			args{
				"abbcde",
			},
			true, false,
		},
		{
			"abcccd contains three c, but no letter appears exactly two times.",
			args{
				"abcccd",
			},
			false, true,
		},
		{
			"aabcdd contains two a and two d, but it only counts once.",
			args{
				"aabcdd",
			},
			true, false,
		},
		{
			"abcdee contains two e.",
			args{
				"abcdee",
			},
			true, false,
		},
		{
			"ababab contains three a and three b, but it only counts once.",
			args{
				"ababab",
			},
			false, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := repetitions(tt.args.s)
			if got != tt.want {
				t.Errorf("twoLetter repetitions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("threeLetter repetitions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_levenshteinDistance(t *testing.T) {
	type args struct {
		source string
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"aixwcbzrmdvpsjfgllthdyeoqe and aixwcbzrmdvpsjfgllthdyioqe",
			args{
				"aixwcbzrmdvpsjfgllthdyeoqe",
				"aixwcbzrmdvpsjfgllthdyioqe",
			},
			1,
		},
		{
			"fghij and fguij",
			args{
				"fghij",
				"fguij",
			},
			1,
		},
		{
			"abcde and axcye",
			args{
				"abcde",
				"axcye",
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levenshteinDistance(tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("levenshteinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
