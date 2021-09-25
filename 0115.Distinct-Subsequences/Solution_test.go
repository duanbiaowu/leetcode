package leetcode

import "testing"

func Test_numDistinct(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{"", ""},
			1,
		},
		{
			"test-2",
			args{"", "a"},
			0,
		},
		{
			"test-3",
			args{"rabbbit", "rabbit"},
			3,
		},
		{
			"test-4",
			args{"babgbag", "bag"},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
