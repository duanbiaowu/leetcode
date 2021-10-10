package leetcode

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{0, 0},
			0,
		},
		{
			"test-2",
			args{1, 2},
			1,
		},
		{
			"test-3",
			args{3, 2},
			3,
		},
		{
			"test-4",
			args{7, 3},
			28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
