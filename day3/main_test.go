package main

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want claim
	}{
		{
			"#123 @ 3,2: 5x4",
			args{
				"#123 @ 3,2: 5x4",
			},
			claim{
				id:           "123",
				leftDistance: 3,
				topDistance:  2,
				width:        5,
				height:       4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
