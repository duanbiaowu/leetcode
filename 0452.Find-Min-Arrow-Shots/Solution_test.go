package leetcode

import (
	"testing"
)

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{
				[][]int{
					{10, 16},
					{2, 8},
					{1, 6},
					{7, 12},
				},
			},
			2,
		},
		{
			"test-2",
			args{
				[][]int{
					{1, 2},
					{3, 4},
					{5, 6},
					{7, 7},
				},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
