package leetcode

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test a empty array",
			args{[]int{}, 1},
			[][]int{},
		},
		{
			"test an array with only one element",
			args{[]int{1}, 1},
			[][]int{},
		},
		{
			"test an array of two elements",
			args{[]int{1, 2}, 1},
			[][]int{},
		},
		{
			"test an array of three elements",
			args{[]int{1, 2, 3}, 1},
			[][]int{},
		},

		{
			"test an array of three elements, but there is no answer",
			args{[]int{1, 2, 3, 4}, 1},
			[][]int{},
		},

		{
			"No1. test an array of four elements, but there is only one answer",
			args{[]int{0, 0, 0, 0}, 0},
			[][]int{
				{0, 0, 0, 0},
			},
		},
		{
			"No2. test an array of four elements, but there is only one answer",
			args{[]int{1, 1, 1, 1}, 4},
			[][]int{
				{1, 1, 1, 1},
			},
		},

		{
			"test an array of multiple elements, but there are two answers",
			args{[]int{0, 1, 5, 0, 1, 5, 5, -4}, 11},
			[][]int{
				{-4, 5, 5, 5},
				{0, 1, 5, 5},
			},
		},

		{
			"No1. test an array of multiple elements, and there are multiple answers",
			args{[]int{1, 0, -1, 0, -2, 2}, 0},
			[][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			"No2. test an array of multiple elements, and there are multiple answers",
			args{[]int{1, 0, -1, 0, -2, 2, 0, 0, 0, 0}, 0},
			[][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
				{0, 0, 0, 0},
			},
		},
		{
			"No3. test an array of multiple elements, and there are multiple answers",
			args{[]int{1, 0, -1, 0, -2, 2, 0, 0, 0, 0}, 1},
			[][]int{
				{-2, 0, 1, 2},
				{-1, 0, 0, 2},
				{0, 0, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
