package leetcode

import "testing"

func Test_longestOnes(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{[]int{}, 0},
			0,
		},
		{
			"test-2",
			args{[]int{0}, 0},
			0,
		},
		{
			"test-3",
			args{[]int{1}, 0},
			1,
		},
		{
			"test-4",
			args{[]int{1, 0}, 0},
			1,
		},
		{
			"test-5",
			args{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2},
			6,
		},
		{
			"test-6",
			args{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
