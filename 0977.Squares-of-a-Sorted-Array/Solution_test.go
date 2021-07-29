package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
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
			args{[]int{0}},
			[]int{0},
		},
		{
			"test-3",
			args{[]int{1, 2}},
			[]int{1, 4},
		},
		{
			"test-4",
			args{[]int{-2, 1}},
			[]int{1, 4},
		},
		{
			"test-5",
			args{[]int{-2, 1, 2}},
			[]int{1, 4, 4},
		},
		{
			"test-6",
			args{[]int{-4, -1, 0, 3, 10}},
			[]int{0, 1, 9, 16, 100},
		},
		{
			"test-7",
			args{[]int{-7, -3, 2, 3, 11}},
			[]int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
