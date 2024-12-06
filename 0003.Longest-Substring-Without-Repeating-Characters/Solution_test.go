package leetcode

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test an empty string",
			args{""},
			0,
		},
		{
			"test a single character string",
			args{"a"},
			1,
		},
		{
			"test a string that contains only one duplicate character",
			args{"bbbbb"},
			1,
		},
		{
			"test a string that contains multiple duplicate characters",
			args{"abba"},
			2,
		},
		{
			"test a string that contains multiple duplicate characters",
			args{"abcabcbb"},
			3,
		},
		{
			"test a string that does not contain duplicate characters",
			args{"let's go"},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
