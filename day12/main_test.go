package main

import (
	"reflect"
	"testing"
)

func Test_pots_grow_simulation(t *testing.T) {
	type fields struct {
		pots    []bool
		zeroPot int
	}
	type args struct {
		generations int
		notes       []note
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []bool
		wantSum int
	}{
		{
			"",
			fields{
				parse("#..#.#..##......###...###"),
				0,
			},
			args{
				generations: 20,
				notes: []note{
					note{parse("...##"), true},
					note{parse("..#.."), true},
					note{parse(".#..."), true},
					note{parse(".#.#."), true},
					note{parse(".#.##"), true},
					note{parse(".##.."), true},
					note{parse(".####"), true},
					note{parse("#.#.#"), true},
					note{parse("#.###"), true},
					note{parse("##.#."), true},
					note{parse("##.##"), true},
					note{parse("###.."), true},
					note{parse("###.#"), true},
					note{parse("####."), true},
				},
			},
			parse("#....##....#####...#######....#.#..##"),
			325,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pots{
				pots:          tt.fields.pots,
				zeroPotOffset: tt.fields.zeroPot,
			}
			simulate(tt.args.generations, p, tt.args.notes)
			if !reflect.DeepEqual(tt.want, p.pots) {
				t.Errorf("Unexpected state after simulating %d generations, want: %v, got:%v", tt.args.generations, tt.want, p.pots)
			}
			if tt.wantSum != p.potNumSum() {
				t.Errorf("Unexpected value after simulating %d generations, want: %v, got:%v", tt.args.generations, tt.wantSum, p.potNumSum())
			}
		})
	}
}
