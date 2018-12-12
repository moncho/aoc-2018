package main

import "testing"

func Test_cell_powerLevel(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		serNo int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			"3,5 grid 8",
			fields{
				x: 3,
				y: 5,
			},
			args{
				8,
			},
			4,
		},
		{
			"122,79 grid 57",
			fields{
				x: 122,
				y: 79,
			},
			args{
				57,
			},
			-5,
		},
		{
			"217,196 grid 39",
			fields{
				x: 217,
				y: 196,
			},
			args{
				39,
			},
			0,
		},
		{
			"101,153 grid 71",
			fields{
				x: 101,
				y: 153,
			},
			args{
				71,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cell{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := c.powerLevel(tt.args.serNo); got != tt.want {
				t.Errorf("cell.powerLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
