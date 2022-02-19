package leetcode

import (
	"reflect"
	"testing"
)

func Test_getMaxMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//{
		//	"test-1",
		//	args{nil},
		//	nil,
		//},
		//{
		//	"test-2",
		//	args{[][]int{
		//		{-1,0},
		//		{0,-1},
		//	}},
		//	[]int{0,1,0,1},
		//},
		{
			"test-3",
			args{[][]int{
				{1, 2, -3},
				{4, 5, -6},
				{7, 8, -9},
			}},
			[]int{0, 0, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMaxMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
