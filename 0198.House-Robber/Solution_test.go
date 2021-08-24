package leetcode

import "testing"

func Test_rob(t *testing.T) {
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
			args{[]int{0}},
			0,
		},
		{
			"test-2",
			args{[]int{1}},
			1,
		},
		{
			"test-3",
			args{[]int{1, 2}},
			2,
		},
		{
			"test-4",
			args{[]int{1, 2, 3}},
			4,
		},
		{
			"test-5",
			args{[]int{1, 2, 3, 4}},
			6,
		},
		{
			"test-6",
			args{[]int{2, 7, 9, 3, 1}},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
