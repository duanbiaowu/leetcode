package leetcode

import (
	"testing"
)

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{"a", "a"},
			"a",
		},
		{
			"test-2",
			args{"a", "aa"},
			"",
		},
		{
			"test-3",
			args{"ADOBECODEBANC", "ABC"},
			"BANC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
