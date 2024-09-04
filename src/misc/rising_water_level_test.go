package main

import "testing"

func Test_yearsLeft(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				grid: [][]int{
					{6, 5, 4, 5, 5},
					{4, 2, 5, 1, 1},
					{5, 5, 2, 1, 5},
					{2, 3, 2, 4, 4},
					{5, 4, 5, 5, 6},
				},
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				grid: [][]int{
					{3, 2, 1},
					{1, 1, 3},
				},
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				grid: [][]int{
					{4, 0},
					{0, 4},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := yearsLeft(tt.args.grid); got != tt.want {
				t.Errorf("yearsLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
