package leetcode

import "testing"

func Test_findWhetherExistsPath(t *testing.T) {
	type args struct {
		n      int
		graph  [][]int
		start  int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{
				5,
				[][]int{
					{0, 1},
					{0, 2},
					{1, 2},
					{1, 2},
				},
				0,
				2,
			},
			true,
		},
		{
			"test-2",
			args{
				5,
				[][]int{
					{0, 1},
					{0, 2},
					{0, 4},
					{0, 4},

					{0, 1},
					{1, 3},
					{1, 4},
					{1, 3},

					{2, 3},
					{3, 4},
				},
				0,
				4,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWhetherExistsPath(tt.args.n, tt.args.graph, tt.args.start, tt.args.target); got != tt.want {
				t.Errorf("findWhetherExistsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
