package main

import (
	"reflect"
	"testing"

	"github.com/ericywl/leetcode/src/common"
)

func Test_reverseList(t *testing.T) {
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
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "example 2",
			args: args{
				arr: []int{1, 2},
			},
			want: []int{2, 1},
		},
		{
			name: "example 3",
			args: args{
				arr: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := reverseList(common.ArrayToListNodes(tt.args.arr))
			headRec := reverseListRecursive(common.ArrayToListNodes(tt.args.arr))

			arr := common.ListNodesToArray(head)
			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("reverseList() = %v, want %v", arr, tt.want)
			}

			arrRec := common.ListNodesToArray(headRec)
			if !reflect.DeepEqual(arrRec, tt.want) {
				t.Errorf("reverseListRecursive() = %v, want %v", arrRec, tt.want)
			}
		})
	}
}
