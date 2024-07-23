package main

import "testing"

func Test_longestZigZag(t *testing.T) {
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
				arr: []any{1, nil, 1, nil, nil, 1, nil, nil, nil, nil, nil, nil, 1, nil, nil},
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				arr: []any{1, 1, 1, nil, nil, 1, nil},
			},
			want: 2,
		},
		{
			name: "example 3",
			args: args{
				arr: []any{1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := arrayToTree(tt.args.arr)
			if got := longestZigZagDFS(root); got != tt.want {
				t.Errorf("longestZigZagDFS() = %v, want %v", got, tt.want)
			}
			if got := longestZigZagBFS(root); got != tt.want {
				t.Errorf("longestZigZagBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
