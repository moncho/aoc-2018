package main

import "testing"

func Test_complete(t *testing.T) {
	type args struct {
		s     string
		steps map[string]*step
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"CABDFE",
			args{
				s: "E",
				steps: buildGraph(
					[]instruction{
						instruction{step: "F", precond: "C"},
						instruction{step: "D", precond: "A"},
						instruction{step: "E", precond: "B"},
						instruction{step: "B", precond: "A"},
						instruction{step: "A", precond: "C"},
						instruction{step: "E", precond: "F"},
						instruction{step: "E", precond: "D"},
					},
				),
			},
			"CABDFE",
		},
		{
			"BAXC",
			args{
				s: "C",
				steps: buildGraph(
					[]instruction{
						instruction{step: "C", precond: "A"},
						instruction{step: "C", precond: "X"},
						instruction{step: "X", precond: "B"},
						instruction{step: "A", precond: "B"},
					},
				),
			},
			"BAXC",
		},
		{
			"DABHJPFZ",
			args{
				s: "Z",
				steps: buildGraph(
					[]instruction{
						instruction{step: "A", precond: "D"},
						instruction{step: "H", precond: "D"},
						instruction{step: "J", precond: "D"},
						instruction{step: "B", precond: "A"},
						instruction{step: "Z", precond: "B"},
						instruction{step: "Z", precond: "F"},
						instruction{step: "F", precond: "P"},
						instruction{step: "F", precond: "H"},
						instruction{step: "P", precond: "J"},
					},
				),
			},
			"DABHJPFZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complete(tt.args.s, tt.args.steps); got != tt.want {
				t.Errorf("complete() = %v, want %v", got, tt.want)
			}
		})
	}
}
