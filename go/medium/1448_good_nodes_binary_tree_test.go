package main

import "testing"

func createNodes(arr []any) []*TreeNode {
	nodes := make([]*TreeNode, 0, len(arr))
	for _, ele := range arr {
		if ele == nil {
			nodes = append(nodes, nil)
			continue
		}

		intVal, ok := ele.(int)
		if !ok {
			panic("invalid array element")
		}

		nodes = append(nodes, &TreeNode{Val: intVal})
	}

	return nodes
}

func arrayToTree(arr []any) *TreeNode {
	nodes := createNodes(arr)
	return arrayNode(nodes, 0)
}

func arrayNode(nodes []*TreeNode, idx int) *TreeNode {
	if idx >= len(nodes) || nodes[idx] == nil {
		return nil
	}

	nodes[idx].Left = arrayNode(nodes, (idx+1)*2-1)
	nodes[idx].Right = arrayNode(nodes, (idx+1)*2)
	return nodes[idx]
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
