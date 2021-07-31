package leetcode

import (
	"reflect"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	type args struct {
		mat [][]int
		r   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{
				[][]int{},
				-1,
				-1,
			},
			[][]int{},
		},
		{
			"test-2",
			args{
				[][]int{},
				0,
				0,
			},
			[][]int{},
		},
		{
			"test-3",
			args{
				[][]int{},
				1,
				1,
			},
			[][]int{},
		},
		{
			"test-4",
			args{
				[][]int{
					{1},
					{2},
				},
				2,
				2,
			},
			[][]int{
				{1},
				{2},
			},
		},
		{
			"test-5",
			args{
				[][]int{
					{1},
					{2},
					{3},
					{4},
					{5},
					{6},
				},
				2,
				2,
			},
			[][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
				{6},
			},
		},
		{
			"test-6",
			args{
				[][]int{
					{1},
					{2},
					{3},
					{4},
				},
				2,
				2,
			},
			[][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			"test-7",
			args{
				[][]int{
					{1},
					{2},
					{3},
					{4},
				},
				1,
				4,
			},
			[][]int{
				{1, 2, 3, 4},
			},
		},
		{
			"test-8",
			args{
				[][]int{
					{1},
					{2},
					{3},
					{4},
					{5},
					{6},
				},
				3,
				2,
			},
			[][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixReshape(tt.args.mat, tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixReshape() = %v, want %v", got, tt.want)
			}
		})
	}
}
