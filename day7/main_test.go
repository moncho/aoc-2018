package main

import "testing"

func Test_sortSteps(t *testing.T) {
	type args struct {
		req requirements
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"CABDFE",
			args{
				req: newRequirements(
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
				req: newRequirements(
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
				req: newRequirements(
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
			if got := tt.args.req.completionOrder(); got != tt.want {
				t.Errorf("complete() = %v, want %v", got, tt.want)
			}
		})
	}
}
