package leetcode

import "testing"

func Test_longestConsecutive(t *testing.T) {
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
			args{nil},
			0,
		},
		{
			"test-2",
			args{[]int{0}},
			1,
		},
		{
			"test-3",
			args{[]int{1}},
			1,
		},
		{
			"test-4",
			args{[]int{100, 4, 200, 1, 3, 2}},
			4,
		},
		{
			"test-5",
			args{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}},
			7,
		},
		{
			"test-6",
			args{[]int{2147483646, -2147483647, 0, 2, 2147483644, -2147483645, 2147483645}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
