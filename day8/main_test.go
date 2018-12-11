package main

import "testing"

func Test_metadataSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2",
			args{
				[]int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2},
			},
			138,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, _ := treeNode(tt.args.numbers)
			if got := node.metadataSum(); got != tt.want {
				t.Errorf("metadataSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodeValue(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2",
			args{
				[]int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2},
			},
			66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, _ := treeNode(tt.args.numbers)
			if got := node.nodeValue(); got != tt.want {
				t.Errorf("nodeValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
