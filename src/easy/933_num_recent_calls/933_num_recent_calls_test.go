package main

import (
	"reflect"
	"testing"
)

func TestRecentCounter(t *testing.T) {
	type args struct {
		ts []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				ts: []int{1, 100, 3001, 3002},
			},
			want: []int{1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Constructor()
			output := make([]int, 0, len(tt.args.ts))
			for _, t := range tt.args.ts {
				output = append(output, r.Ping(t))
			}

			if !reflect.DeepEqual(output, tt.want) {
				t.Errorf("RecentCounter.Ping() = %v, want %v", output, tt.want)
			}
		})
	}
}
