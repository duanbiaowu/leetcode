package leetcode

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
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
			args{[]int{}},
			[][]int{
				{},
			},
		},
		{
			"test-2",
			args{[]int{1}},
			[][]int{
				{1},
			},
		},
		{
			"test-3",
			args{[]int{0, 1}},
			[][]int{
				{0, 1},
				{1, 0},
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
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
