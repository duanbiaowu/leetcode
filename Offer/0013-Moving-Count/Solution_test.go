package leetcode

import "testing"

func Test_movingCount(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{0, 0, 0},
			0,
		},
		{
			"test-2",
			args{2, 3, 1},
			3,
		},
		{
			"test-3",
			args{3, 1, 0},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
