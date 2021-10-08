package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
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
			args{nil},
			nil,
		},
		{
			"test-2",
			args{[]int{}},
			[]int{},
		},
		{
			"test-3",
			args{[]int{1}},
			[]int{1},
		},
		{
			"test-4",
			args{[]int{2, 0, 2, 1, 1, 0}},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			"test-5",
			args{[]int{2, 0, 1, 1, 2, 0, 2, 1, 2, 0, 0, 0, 1, 2, 2, 2, 0, 1, 1}},
			[]int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("sortColors() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
