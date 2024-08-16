package main

import (
	"reflect"
	"testing"

	"github.com/ericywl/leetcode/src/common"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		arr []any
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				arr: []any{1, 2, 3, nil, 5, nil, 4},
			},
			want: []int{1, 3, 4},
		},
		{
			name: "example 2",
			args: args{
				arr: []any{1, nil, 3},
			},
			want: []int{1, 3},
		},
		{
			name: "example 3",
			args: args{
				arr: []any{},
			},
			want: []int{},
		},
		{
			name: "example 4",
			args: args{
				arr: []any{1, 2, 3, nil, 5, 6, nil, nil, nil, 4},
			},
			want: []int{1, 3, 6, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := common.ArrayToTree(tt.args.arr)
			if got := rightSideView(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
