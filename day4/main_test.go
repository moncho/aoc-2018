package main

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		logEntry string
	}
	tests := []struct {
		name string
		args args
		want recordEntry
	}{
		{
			"[1518-11-01 00:00] Guard #10 begins shift",
			args{
				"[1518-11-01 00:00] Guard #10 begins shift",
			},
			recordEntry{
				guardID:    "10",
				minute:     0,
				recordType: NewShift,
				timestamp:  parseTimeStamp("[1518-11-01 00:00]"),
			},
		},
		{
			"[1518-11-01 00:05] falls asleep",
			args{
				"[1518-11-01 00:05] falls asleep",
			},
			recordEntry{
				guardID:    "",
				minute:     5,
				recordType: Sleep,
				timestamp:  parseTimeStamp("[1518-11-01 00:05]"),
			},
		},
		{
			"[1518-11-01 00:25] wakes up",
			args{
				"[1518-11-01 00:25] wakes up",
			},
			recordEntry{
				guardID:    "",
				minute:     25,
				recordType: Awake,
				timestamp:  parseTimeStamp("[1518-11-01 00:25]"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.logEntry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
