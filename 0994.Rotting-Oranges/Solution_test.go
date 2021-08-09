package leetcode

import "testing"

func Test_orangesRotting(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{[][]int{}},
			-1,
		},
		{
			"test-2",
			args{[][]int{
				{1, 0},
			}},
			-1,
		},
		{
			"test-3",
			args{[][]int{
				{0, 2},
			}},
			0,
		},
		{
			"test-4",
			args{[][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			}},
			4,
		},
		{
			"test-5",
			args{[][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orangesRotting(tt.args.grid); got != tt.want {
				t.Errorf("orangesRotting() = %v, want %v", got, tt.want)
			}
		})
	}
}
