package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{[]int{}, 0},
			[]int{},
		},
		{
			"test-2",
			args{[]int{}, 1},
			[]int{},
		},
		{
			"test-3",
			args{[]int{1}, 0},
			[]int{1},
		},
		{
			"test-4",
			args{[]int{1}, 1},
			[]int{1},
		},
		{
			"test-5",
			args{[]int{1, 2, 3, 4, 5}, 4},
			[]int{2, 3, 4, 5, 1},
		},
		{
			"test-6",
			args{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"test-7",
			args{[]int{-1, -100, 3, 99}, 2},
			[]int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
