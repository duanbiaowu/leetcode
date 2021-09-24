package leetcode

import "testing"

func Test_minimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{nil},
			0,
		},
		{
			"test-2",
			args{[][]int{
				{-10},
			}},
			-10,
		},
		{
			"test-3",
			args{[][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			}},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
