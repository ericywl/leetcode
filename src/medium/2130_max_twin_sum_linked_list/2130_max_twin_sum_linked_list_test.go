package main

import (
	"testing"

	"github.com/ericywl/leetcode/src/common"
)

func Test_pairSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				arr: []int{5, 4, 2, 1},
			},
			want: 6,
		},
		{
			name: "example 2",
			args: args{
				arr: []int{4, 2, 2, 3},
			},
			want: 7,
		},
		{
			name: "example 3",
			args: args{
				arr: []int{1, 100000},
			},
			want: 100001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pairSum(common.ArrayToListNodes(tt.args.arr))
			// if got != tt.want {
			// 	t.Errorf("pairSum() = %v, want %v", got, tt.want)
			// }
		})
	}
}
