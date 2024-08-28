package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example 1",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4, 6},
			},
			want: [][]int{
				{1, 3},
				{4, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifference(tt.args.nums1, tt.args.nums2)
			sort.Ints(got[0])
			sort.Ints(got[1])
			sort.Ints(tt.want[0])
			sort.Ints(tt.want[1])
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
