package leetcode

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{[]int{0, 0, 0, 0, 0}, 1},
			0,
		},
		{
			"test-2",
			args{[]int{-1, 0, 1, 1, 55}, 3},
			2,
		},
		{
			"test-3",
			args{[]int{-1, 2, 1, -4}, 1},
			2,
		},
		{
			"test-4",
			args{[]int{-1, 2, 1, -4}, 2},
			2,
		},
		{
			"test-5",
			args{[]int{-1, 2, 1, 4, 10}, 17},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
