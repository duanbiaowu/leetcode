package leetcode

import "testing"

func Test_longestValidParentheses(t *testing.T) {
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
			args{"(()"},
			2,
		},
		{
			"test-2",
			args{")()())"},
			4,
		},
		{
			"test-3",
			args{"()(())"},
			6,
		},
		{
			"test-4",
			args{"()(())))"},
			6,
		},
		{
			"test-5",
			args{"(()())"},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
