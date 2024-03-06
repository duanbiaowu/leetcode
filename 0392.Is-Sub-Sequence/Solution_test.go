package leetcode

import (
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{"", ""},
			true,
		},
		{
			"test-2",
			args{"abc", "ahbgdc"},
			true,
		},
		{
			"test-3",
			args{"axc", "ahbgdc"},
			false,
		},
		{
			"test-4",
			args{"hello", "hello world"},
			true,
		},
		{
			"test-5",
			args{"b", "c"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
