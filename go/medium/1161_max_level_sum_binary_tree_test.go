package main

import "testing"

func Test_maxLevelSum(t *testing.T) {
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
				arr: []any{1, 7, 0, 7, -8, nil, nil},
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				arr: []any{989, nil, 10250, nil, nil, 98693, -89388, nil, nil, nil, nil, nil, nil, nil, -32127},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := arrayToTree(tt.args.arr)
			if got := maxLevelSum(root); got != tt.want {
				t.Errorf("maxLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
