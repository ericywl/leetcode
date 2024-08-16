package main

import (
	"testing"

	"github.com/ericywl/leetcode/src/common"
)

func Test_leafSimilar(t *testing.T) {
	type args struct {
		arr1 []any
		arr2 []any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				arr1: []any{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4},
				arr2: []any{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8},
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				arr1: []any{1, 2, 3},
				arr2: []any{1, 3, 2},
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				arr1: []any{1, 2},
				arr2: []any{2, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root1 := common.ArrayToTree(tt.args.arr1)
			root2 := common.ArrayToTree(tt.args.arr2)
			if got := leafSimilar(root1, root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
