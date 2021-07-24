package leetcode

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test a empty array",
			args{[]int{}},
			[][]int{},
		},
		{
			"test an array with only one element",
			args{[]int{1}},
			[][]int{},
		},
		{
			"test an array of two elements",
			args{[]int{1, 2}},
			[][]int{},
		},

		{
			"test an array of three elements, but there is no answer",
			args{[]int{1, 2, 3}},
			[][]int{},
		},

		{
			"test an array of three elements, but there is only one answer",
			args{[]int{0, 0, 0}},
			[][]int{
				{0, 0, 0},
			},
		},

		{
			"test an array of three elements, but there are two answers",
			args{[]int{-1, 0, 1, 2, -1, -4}},
			[][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},

		{
			"No1. test an array of three elements, but there are multiple answers",
			args{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}},
			[][]int{
				{-4, -2, 6},
				{-4, 0, 4},
				{-4, 1, 3},
				{-4, 2, 2},
				{-2, -2, 4},
				{-2, 0, 2},
			},
		},
		{
			"No2. test an array of three elements, but there are multiple answers",
			args{[]int{5, -7, 3, -3, 5, -10, 4, 8, -3, -8, -3, -3, -1, -8, 6, 4, -4, 7, 2, -5, -2, -7, -3, 7, 2, 4, -6, 5}},
			[][]int{
				{-10, 2, 8},
				{-10, 3, 7},
				{-10, 4, 6},
				{-10, 5, 5},
				{-8, 2, 6},
				{-8, 3, 5},
				{-8, 4, 4},
				{-7, -1, 8},
				{-7, 2, 5},
				{-7, 3, 4},
				{-6, -2, 8},
				{-6, -1, 7},
				{-6, 2, 4},
				{-5, -3, 8},
				{-5, -2, 7},
				{-5, -1, 6},
				{-5, 2, 3},
				{-4, -3, 7},
				{-4, -2, 6},
				{-4, -1, 5},
				{-4, 2, 2},
				{-3, -3, 6},
				{-3, -2, 5},
				{-3, -1, 4},
				{-2, -1, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
