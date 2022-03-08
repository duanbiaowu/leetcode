package leetcode

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{2, [][]int{
				{1, 0},
				{0, 1},
			}},
			false,
		},
		{
			"test-2",
			args{2, [][]int{
				{1, 0},
			}},
			true,
		},
		{
			"test-2",
			args{6, [][]int{
				{3, 0},
				{3, 1},
				{4, 1},
				{4, 2},
				{5, 3},
				{5, 4},
			}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish2(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
