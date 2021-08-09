package leetcode

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{0},
			1,
		},
		{
			"test-2",
			args{1},
			1,
		},
		{
			"test-3",
			args{2},
			2,
		},
		{
			"test-4",
			args{3},
			3,
		},
		{
			"test-5",
			args{4},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
