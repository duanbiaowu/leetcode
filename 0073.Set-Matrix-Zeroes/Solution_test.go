package leetcode

import (
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{[][]int{}},
			[][]int{},
		},
		{
			"test-2",
			args{[][]int{
				{1},
				{2},
				{3},
				{4},
			}},
			[][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		{
			"test-3",
			args{[][]int{
				{1, 2},
				{3, 4},
			}},
			[][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			"test-4",
			args{[][]int{
				{1, 2},
				{3, 0},
			}},
			[][]int{
				{1, 0},
				{0, 0},
			},
		},
		{
			"test-5",
			args{[][]int{
				{1, 2, 3},
				{4, 0, 5},
			}},
			[][]int{
				{1, 0, 3},
				{0, 0, 0},
			},
		},
		{
			"test-6",
			args{[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			}},
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			"test-7",
			args{[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			}},
			[][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			if !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("setZeroes() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}
