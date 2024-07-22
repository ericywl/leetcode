package main

import "testing"

func arrayToTree(arr []any) *TreeNode {
	return arrayNode(arr, 0)
}

func arrayNode(arr []any, idx int) *TreeNode {
	if idx >= len(arr) || arr[idx] == nil {
		return nil
	}

	intVal, ok := arr[idx].(int)
	if !ok {
		panic("invalid array element")
	}

	node := &TreeNode{Val: intVal}
	node.Left = arrayNode(arr, (idx+1)*2-1)
	node.Right = arrayNode(arr, (idx+1)*2)
	return node
}

func Test_goodNodes(t *testing.T) {
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
				arr: []any{3, 1, 4, 3, nil, 1, 5},
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				arr: []any{3, 3, nil, 4, 2},
			},
			want: 3,
		},
		{
			name: "example 3",
			args: args{
				arr: []any{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := arrayToTree(tt.args.arr)
			if got := goodNodes(root); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
