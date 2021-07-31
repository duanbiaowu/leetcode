package leetcode

import (
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{-1},
			[][]int{},
		},
		{
			"test-2",
			args{0},
			[][]int{},
		},
		{
			"test-3",
			args{1},
			[][]int{
				{1},
			},
		},
		{
			"test-4",
			args{2},
			[][]int{
				{1},
				{1, 1},
			},
		},
		{
			"test-5",
			args{3},
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
			},
		},
		{
			"test-6",
			args{4},
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
			},
		},
		{
			"test-5",
			args{5},
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
