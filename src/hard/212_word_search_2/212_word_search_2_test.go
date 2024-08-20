package hard

import (
	"reflect"
	"slices"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example 1",
			args: args{
				board: [][]byte{
					{'o', 'a', 'a', 'n'},
					{'e', 't', 'a', 'e'},
					{'i', 'h', 'k', 'r'},
					{'i', 'f', 'l', 'v'},
				},
				words: []string{
					"oath", "pea", "eat", "rain",
				},
			},
			want: []string{"eat", "oath"},
		},
		{
			name: "example 2",
			args: args{
				board: [][]byte{
					{'a', 'b'},
					{'c', 'd'},
				},
				words: []string{"abcb"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findWords(tt.args.board, tt.args.words)
			// Sort the slices before comparing
			slices.Sort(got)
			slices.Sort(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
