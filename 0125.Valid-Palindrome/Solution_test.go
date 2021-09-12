package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{""},
			true,
		},
		{
			"test-2",
			args{"a"},
			true,
		},
		{
			"test-3",
			args{"aba"},
			true,
		},
		{
			"test-4",
			args{"A man, a plan, a canal: Panama"},
			true,
		},
		{
			"test-5",
			args{"race a car"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
