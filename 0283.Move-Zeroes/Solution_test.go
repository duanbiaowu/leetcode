package leetcode

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{[]int{}},
			[]int{},
		},
		{
			"test-2",
			args{[]int{1}},
			[]int{1},
		},
		{
			"test-3",
			args{[]int{0, 1}},
			[]int{1, 0},
		},
		{
			"test-4",
			args{[]int{1, 0}},
			[]int{1, 0},
		},
		{
			"test-5",
			args{[]int{0, 1, 0, 3, 12}},
			[]int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
