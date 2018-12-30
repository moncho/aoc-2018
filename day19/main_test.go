package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		p program
		r registers
	}
	tests := []struct {
		name string
		args args
		want registers
	}{
		{
			"test",
			args{
				program{
					0,
					[]operation{
						operation{"seti", 5, 0, 1},
						operation{"seti", 6, 0, 2},
						operation{"addi", 0, 1, 0},
						operation{"addr", 1, 2, 3},
						operation{"setr", 1, 0, 0},
						operation{"seti", 8, 0, 4},
						operation{"seti", 9, 0, 5},
					},
				},
				[]int{0, 0, 0, 0, 0, 0},
			},
			[]int{6, 5, 6, 0, 0, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.p, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
