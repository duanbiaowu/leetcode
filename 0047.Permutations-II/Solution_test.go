package leetcode

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{[]int{1, 1, 2}},
			[][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{
			"test-2",
			args{[]int{1, 2, 2}},
			[][]int{
				{1, 2, 2},
				{2, 1, 2},
				{2, 2, 1},
			},
		},
		{
			"test-3",
			args{[]int{2, 2, 2}},
			[][]int{
				{2, 2, 2},
			},
		},
		{
			"test-4",
			args{[]int{1, 2, 3}},
			[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
