package leetcode

import (
	"reflect"
	"testing"
)

func Test_findSquare(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{[][]int{nil}},
			nil,
		},
		{
			"test-3",
			args{[][]int{
				{1, 0, 1},
				{0, 0, 1},
				{0, 0, 1},
			}},
			[]int{1, 0, 2},
		},
		{
			"test-3",
			args{[][]int{
				{0, 1, 1},
				{1, 0, 1},
				{1, 1, 0},
			}},
			[]int{0, 0, 1},
		},
		{
			"test-4",
			args{[][]int{
				{1},
			}},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSquare(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
