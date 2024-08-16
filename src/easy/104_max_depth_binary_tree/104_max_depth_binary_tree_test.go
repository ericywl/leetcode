package main

import (
	"testing"

	"github.com/ericywl/leetcode/src/common"
)

func Test_maxDepth(t *testing.T) {
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
				arr: []any{3, 9, 20, nil, nil, 15, 7},
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				arr: []any{1, nil, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := common.ArrayToTree(tt.args.arr)
			if got := maxDepthDFS(root); got != tt.want {
				t.Errorf("maxDepthDFS() = %v, want %v", got, tt.want)
			}
			if got := maxDepthBFS(root); got != tt.want {
				t.Errorf("maxDepthBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
