package main

import (
	"testing"

	"github.com/ericywl/leetcode/src/common"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		arr       []any
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				arr:       []any{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1},
				targetSum: 8,
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				arr:       []any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1},
				targetSum: 22,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := common.ArrayToTree(tt.args.arr)
			if got := pathSum(root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
