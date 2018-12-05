package main

import (
	"testing"
)

func Test_guard_totalSleeping(t *testing.T) {
	tests := []struct {
		name  string
		guard guard
		want  int
	}{
		{
			"Total sleeping time in minutes",
			guard{
				id: "1",
				sleepPattern: map[int]int{
					1: 1,
					2: 1,
					3: 3,
				},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.guard.totalSleeping(); got != tt.want {
				t.Errorf("guard.totalSleeping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_guard_sleepiestMinute(t *testing.T) {
	tests := []struct {
		name  string
		guard guard
		want  int
	}{
		{
			"Sleepiest minute",
			guard{
				id: "1",
				sleepPattern: map[int]int{
					1: 1,
					2: 1,
					3: 3,
					4: 3,
				},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.guard.sleepiestMinute(); got != tt.want {
				t.Errorf("guard.sleepiestMinute() = %v, want %v", got, tt.want)
			}
		})
	}
}
