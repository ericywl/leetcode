package main

import (
	"reflect"
	"testing"

	"github.com/ericywl/leetcode/src/common"
)

func Test_deleteMiddle(t *testing.T) {
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
				arr: []int{1, 3, 4, 7, 1, 2, 6},
			},
			want: []int{1, 3, 4, 1, 2, 6},
		},
		{
			name: "example 2",
			args: args{
				arr: []int{1, 2, 3, 4},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "example 3",
			args: args{
				arr: []int{2, 1},
			},
			want: []int{2},
		},
		{
			name: "example 4",
			args: args{
				arr: []int{1},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := deleteMiddle(common.ArrayToListNodes(tt.args.arr))

			var gotArr []int
			for head != nil {
				gotArr = append(gotArr, head.Val)
				head = head.Next
			}

			if !reflect.DeepEqual(gotArr, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", gotArr, tt.want)
			}
		})
	}
}
