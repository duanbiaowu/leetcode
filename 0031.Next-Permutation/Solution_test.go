package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
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
			args{[]int{1}},
			[]int{1},
		},
		{
			"test-3",
			args{[]int{1, 2, 3}},
			[]int{1, 3, 2},
		},
		{
			"test-4",
			args{[]int{3, 2, 1}},
			[]int{1, 2, 3},
		},
		{
			"test-5",
			args{[]int{1, 1, 5}},
			[]int{1, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("nextPermutation() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
