package main

import (
	"testing"

	"github.com/ericywl/leetcode/src/common"
)

func Test_goodNodes(t *testing.T) {
	type args struct {
		arr []any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				arr: []any{3, 1, 4, 3, nil, 1, 5},
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				arr: []any{3, 3, nil, 4, 2},
			},
			want: 3,
		},
		{
			name: "example 3",
			args: args{
				arr: []any{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodNodes(common.ArrayToTree(tt.args.arr)); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
			}
			if got := goodNodesIter(common.ArrayToTree(tt.args.arr)); got != tt.want {
				t.Errorf("goodNodesIter() = %v, want %v", got, tt.want)
			}
		})
	}
}
