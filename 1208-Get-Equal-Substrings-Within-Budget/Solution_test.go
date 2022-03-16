package leetcode

import "testing"

func Test_equalSubstring(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{"a", "b", 0},
			0,
		},
		{
			"test-2",
			args{"abcd", "bcdf", 3},
			3,
		},
		{
			"test-3",
			args{"abcd", "cdef", 3},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSubstring(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
