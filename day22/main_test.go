package main

import "testing"

func Test_cave_riskLevel(t *testing.T) {
	type fields struct {
		depth   int
		targetX int
		targetY int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"",
			fields{
				depth:   510,
				targetX: 10,
				targetY: 10,
			},
			114,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cave{
				depth:           tt.fields.depth,
				targetX:         tt.fields.targetX,
				targetY:         tt.fields.targetY,
				geologicIndices: make(map[xy]int),
			}
			if got := c.riskLevel(); got != tt.want {
				t.Errorf("cave.riskLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
