package leetcode

import "testing"

func Test_subarraySum(t *testing.T) {
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
			args{nil, 0},
			0,
		},
		{
			"test-2",
			args{[]int{1}, 0},
			0,
		},
		{
			"test-3",
			args{[]int{1, 1, 1}, 2},
			2,
		},
		{
			"test-4",
			args{[]int{1, 2, 3}, 3},
			2,
		},
		{
			"test-5",
			args{[]int{2, 1, 3}, 3},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
