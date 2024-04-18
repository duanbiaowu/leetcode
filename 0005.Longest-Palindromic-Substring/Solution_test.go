package leetcode

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test an empty string",
			args{""},
			"",
		},
		{
			"test a single character string",
			args{"a"},
			"a",
		},
		{
			"test a string that contains only one duplicate character",
			args{"bbbbb"},
			"bbbbb",
		},
		{
			"No.1 test a string that contains multiple duplicate characters",
			args{"bcabad"},
			"aba",
		},
		{
			"No.2 test a string that contains multiple duplicate characters",
			args{"cbbd"},
			"bb",
		},
		{
			"No.3 test a string that contains multiple duplicate characters",
			args{"abccbefg"},
			"bccb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeMemo(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
