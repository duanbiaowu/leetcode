package leetcode

import "testing"

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{[]int{2, 2, 3, 2}},
			3,
		},
		{
			"test-2",
			args{[]int{0, 1, 0, 1, 0, 1, 99}},
			99,
		},
		{
			"test-3",
			args{[]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}},
			-4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberOpt2(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
