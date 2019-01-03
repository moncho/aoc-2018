package main

import (
	"io"
	"strings"
	"testing"
)

func Test_route_longestRoute(t *testing.T) {
	type fields struct {
		regex io.Reader
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"^WNE$",
			fields{
				regex: strings.NewReader("^WNE$"),
			},
			3,
		},
		{
			"^N(EEENWWW|N)$",
			fields{
				regex: strings.NewReader("^N(EEENWWW|N)$"),
			},
			8,
		},

		{
			"^ENWWW(NEEE|SSE(EE|N))$",
			fields{
				regex: strings.NewReader("^ENWWW(NEEE|SSE(EE|N))$"),
			},
			10,
		},
		{
			"^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$",
			fields{
				regex: strings.NewReader("^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$"),
			},
			18,
		},
		{
			"^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$",
			fields{
				regex: strings.NewReader("^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"),
			},
			23,
		},
		{
			"^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$",
			fields{
				regex: strings.NewReader("^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$"),
			},
			31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := buildRoute(tt.fields.regex)
			if got := r.longestRoute(); got != tt.want {
				t.Errorf("route.longestRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_route_allRoutes(t *testing.T) {
	type fields struct {
		regex io.Reader
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			"^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$",
			fields{
				regex: strings.NewReader("^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$"),
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 9, 10, 11, 12, 13, 14, 15, 14, 15, 16, 17, 16, 17, 18},
		},

		{
			"^ENNWSWW(NEWS|NES)SSSEEN(WNSE|)EE(SWEN|)NNN$",
			fields{
				regex: strings.NewReader("^ENNWSWW(NEWS|NES)SSSEEN(WNSE|)EE(SWEN|)NNN$"),
			},
			[]int{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 17, 18, 19, 20, 19, 20, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := buildRoute(tt.fields.regex)
			got := r.allDistances()
			if len(got) != len(tt.want) {
				t.Errorf("route.allRoutes() = %v, want %v", len(got), len(tt.want))
			}
		})
	}
}
