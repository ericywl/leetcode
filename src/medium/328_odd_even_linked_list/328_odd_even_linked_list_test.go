package main

import (
	"reflect"
	"testing"

	"github.com/ericywl/leetcode/src/common"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 3, 5, 2, 4},
		},
		{
			name: "example 2",
			args: args{
				arr: []int{2, 1, 3, 5, 6, 4, 7},
			},
			want: []int{2, 3, 6, 7, 1, 5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := oddEvenList(common.ArrayToListNodes(tt.args.arr))

			var gotArr []int
			for head != nil {
				gotArr = append(gotArr, head.Val)
				head = head.Next
			}

			if !reflect.DeepEqual(gotArr, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", gotArr, tt.want)
			}
		})
	}
}
