package leetcode

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{"aa", "a"},
			false,
		},
		{
			"test-2",
			args{"aa", "a*"},
			true,
		},
		{
			"test-3",
			args{"ab", ".*"},
			true,
		},
		{
			"test-4",
			args{"aab", "c*a*b*"},
			true,
		},
		{
			"test-5",
			args{"mississippi", "mis*is*p*."},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
