package main

import (
	"reflect"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "example 1",
			args: args{
				equations: [][]string{
					{"a", "b"},
					{"b", "c"},
				},
				values: []float64{
					2.0,
					3.0,
				},
				queries: [][]string{
					{"a", "c"},
					{"b", "a"},
					{"a", "e"},
					{"a", "a"},
					{"x", "x"},
				},
			},
			want: []float64{
				6.0,
				0.5,
				-1.0,
				1.0,
				-1.0,
			},
		},
		{
			name: "example 2",
			args: args{
				equations: [][]string{
					{"a", "b"},
					{"b", "c"},
					{"bc", "cd"},
				},
				values: []float64{
					1.5,
					2.5,
					5.0,
				},
				queries: [][]string{
					{"a", "c"},
					{"c", "b"},
					{"bc", "cd"},
					{"cd", "bc"},
				},
			},
			want: []float64{
				3.75,
				0.4,
				5.0,
				0.2,
			},
		},
		{
			name: "example 2",
			args: args{
				equations: [][]string{
					{"a", "b"},
				},
				values: []float64{
					0.5,
				},
				queries: [][]string{
					{"a", "b"},
					{"b", "a"},
					{"a", "c"},
					{"x", "y"},
				},
			},
			want: []float64{
				0.5,
				2.0,
				-1.0,
				-1.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
