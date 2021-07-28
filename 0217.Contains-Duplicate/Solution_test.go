package leetcode

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{[]int{}},
			false,
		},
		{
			"test-2",
			args{[]int{1}},
			false,
		},
		{
			"test-3",
			args{[]int{2}},
			false,
		},
		{
			"test-4",
			args{[]int{1, 2}},
			false,
		},
		{
			"test-5",
			args{[]int{1, 2, 1}},
			true,
		},
		{
			"test-6",
			args{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
