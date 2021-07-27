package leetcode

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
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
			args{[]int{0}, 1},
			1,
		},
		{
			"test-4",
			args{[]int{1}, 0},
			1,
		},
		{
			"test-5",
			args{[]int{1, 1}, 1},
			0,
		},
		{
			"test-6",
			args{[]int{1, 1}, 0},
			2,
		},
		{
			"test-7",
			args{[]int{1, 2, 3, 2, 1}, 1},
			3,
		},
		{
			"test-8",
			args{[]int{3, 2, 2, 3}, 3},
			2,
		},
		{
			"test-9",
			args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
