package leetcode

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		columnTitle string
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
			args{"A"},
			1,
		},
		{
			"test-3",
			args{"AB"},
			28,
		},
		{
			"test-4",
			args{"ZY"},
			701,
		},
		{
			"test-5",
			args{"FXSHRXW"},
			2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.columnTitle); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
