package main

import (
	"reflect"
	"testing"
)

func Test_mostInRange(t *testing.T) {
	type args struct {
		nanos []nanobot
	}
	tests := []struct {
		name string
		args args
		want xyz
	}{
		{
			name: "",
			args: args{
				nanos: []nanobot{
					nanobot{xyz: xyz{10, 12, 12}, r: 2},
					nanobot{xyz: xyz{12, 14, 12}, r: 2},
					nanobot{xyz: xyz{16, 12, 12}, r: 4},
					nanobot{xyz: xyz{14, 14, 14}, r: 6},
					nanobot{xyz: xyz{50, 50, 50}, r: 200},
					nanobot{xyz: xyz{10, 10, 10}, r: 5},
				},
			},
			want: xyz{12, 12, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostInRange(tt.args.nanos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
