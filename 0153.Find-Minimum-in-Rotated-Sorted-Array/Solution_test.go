package leetcode

import "testing"

func Test_findMin(t *testing.T) {
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
			args{[]int{1}},
			1,
		},
		{
			"test-2",
			args{[]int{1, 2}},
			1,
		},
		{
			"test-3",
			args{[]int{2, 3, 1}},
			1,
		},
		{
			"test-4",
			args{[]int{1, 2, 3}},
			1,
		},
		{
			"test-5",
			args{[]int{3, 4, 5, 1, 2}},
			1,
		},
		{
			"test-6",
			args{[]int{6, 5, 4, 3, 2}},
			2,
		},
		{
			"test-7",
			args{[]int{4, 5, 6, 7, 0, 1, 2}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
