package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{[]int{}, 0},
			[][]int{
				{},
			},
		},
		{
			"test-2",
			args{[]int{1}, 0},
			[][]int{
				{},
			},
		},
		{
			"test-3",
			args{[]int{1}, 1},
			[][]int{
				{1},
			},
		},
		{
			"test-4",
			args{[]int{2, 5, 2, 1, 2}, 5},
			[][]int{
				{1, 2, 2},
				{5},
			},
		},
		{
			"test-5",
			args{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
			[][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
