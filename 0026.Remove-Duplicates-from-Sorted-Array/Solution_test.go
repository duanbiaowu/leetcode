package leetcode

import "testing"

func Test_removeDuplicates(t *testing.T) {
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
			args{[]int{}},
			0,
		},
		{
			"test-2",
			args{[]int{1}},
			1,
		},
		{
			"test-3",
			args{[]int{1, 1}},
			1,
		},
		{
			"test-4",
			args{[]int{1, 2}},
			2,
		},
		{
			"test-5",
			args{[]int{1, 2, 3}},
			3,
		},
		{
			"test-6",
			args{[]int{1, 2, 3, 3, 3}},
			3,
		},
		{
			"test-7",
			args{[]int{1, 2, 3, 3, 4}},
			4,
		},
		{
			"test-8",
			args{[]int{1, 2, 3, 4, 4, 4}},
			4,
		},
		{
			"test-8",
			args{[]int{1, 2, 3, 4, 5, 6, 7}},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
