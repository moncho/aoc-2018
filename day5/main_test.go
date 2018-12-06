package main

import (
	"testing"
)

func Test_triggerReaction(t *testing.T) {
	reaction := reactionByPolarity()
	type args struct {
		polymer string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"aAsdaA",
			args{
				"aAsdaA",
			},
			"sd",
		},
		{
			"AasdaA",
			args{
				"AasdaA",
			},
			"sd",
		},
		{
			"AbBa",
			args{
				"AbBa",
			},
			"",
		},
		{
			"aAsdaADc",
			args{
				"aAsdaADc",
			},
			"sc",
		},
		{
			"dabAcCaCBAcCcaDA",
			args{
				"dabAcCaCBAcCcaDA",
			},
			"dabCBAcaDA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triggerReaction(tt.args.polymer, reaction); got != tt.want {
				t.Errorf("triggerReaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"-1",
			args{
				-1,
			},
			1,
		},
		{
			"aA",
			args{
				int("Aa"[0]) - int("Aa"[1]),
			},
			32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cleanPolymer(t *testing.T) {
	type args struct {
		polymer string
		unit    rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"dabAcCaCBAcCcaDA removing a/A",
			args{
				"dabAcCaCBAcCcaDA",
				'a',
			},
			"dbcCCBcCcD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanPolymer(tt.args.polymer, tt.args.unit); got != tt.want {
				t.Errorf("cleanPolymer() = %v, want %v", got, tt.want)
			}
		})
	}
}
