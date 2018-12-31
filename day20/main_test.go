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
