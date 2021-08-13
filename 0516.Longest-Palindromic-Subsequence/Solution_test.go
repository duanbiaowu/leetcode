package leetcode

import "testing"

func Test_longestPalindromeSubseq(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{""},
			0,
		},
		{
			"test-2",
			args{"a"},
			1,
		},
		{
			"test-3",
			args{"ab"},
			1,
		},
		{
			"test-4",
			args{"bbbab"},
			4,
		},
		{
			"test-5",
			args{"cbbd"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
