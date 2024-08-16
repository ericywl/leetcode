package main

import (
	"reflect"
	"testing"

	"github.com/ericywl/leetcode/src/common"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		arr  []any
		pIdx int
		qIdx int
	}
	tests := []struct {
		name    string
		args    args
		wantIdx int
	}{
		{
			name: "example 1",
			args: args{
				arr:  []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
				pIdx: 1,
				qIdx: 2,
			},
			wantIdx: 0,
		},
		{
			name: "example 2",
			args: args{
				arr:  []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
				pIdx: 1,
				qIdx: 10,
			},
			wantIdx: 1,
		},
		{
			name: "example 3",
			args: args{
				arr:  []any{1, 2},
				pIdx: 0,
				qIdx: 1,
			},
			wantIdx: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root, nodes := common.ArrayToTreeWithNodes(tt.args.arr)
			want := nodes[tt.wantIdx]
			if got := lowestCommonAncestor(root, nodes[tt.args.pIdx], nodes[tt.args.qIdx]); !reflect.DeepEqual(got, want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, want)
			}
		})
	}
}
