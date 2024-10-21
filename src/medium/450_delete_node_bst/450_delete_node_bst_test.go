package main

import (
	"reflect"
	"testing"

	"github.com/ericywl/leetcode/src/common"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		arr []any
		key int
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "example 1",
			args: args{
				arr: []any{5, 3, 6, 2, 4, nil, 7},
				key: 3,
			},
			want: []any{5, 4, 6, 2, nil, nil, 7},
		},
		{
			name: "example 2",
			args: args{
				arr: []any{5, 3, 6, 2, 4, nil, 7},
				key: 0,
			},
			want: []any{5, 3, 6, 2, 4, nil, 7},
		},
		{
			name: "example 3",
			args: args{
				arr: nil,
				key: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := deleteNode(common.ArrayToTree(tt.args.arr), tt.args.key)
			got := common.TreeToArray(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
