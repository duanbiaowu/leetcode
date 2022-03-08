package leetcode

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{2, [][]int{
				{1, 0},
				{0, 1},
			}},
			[]int{},
		},
		{
			"test-2",
			args{2, [][]int{
				{1, 0},
			}},
			[]int{0, 1},
		},
		{
			"test-3",
			args{6, [][]int{
				{3, 0},
				{3, 1},
				{4, 1},
				{4, 2},
				{5, 3},
				{5, 4},
			}},
			[]int{0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
