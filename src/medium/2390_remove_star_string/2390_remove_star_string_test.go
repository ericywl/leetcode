package main

import "testing"

func Test_removeStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				s: "leet**cod*e",
			},
			want: "lecoe",
		},
		{
			name: "example 2",
			args: args{
				s: "erase*****",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeStars(tt.args.s); got != tt.want {
				t.Errorf("removeStars() = %v, want %v", got, tt.want)
			}
		})
	}
}
