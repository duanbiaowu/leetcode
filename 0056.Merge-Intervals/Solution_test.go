package leetcode

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{[][]int{}},
			[][]int{},
		},
		{
			"test-3",
			args{[][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			}},
			[][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			"test-4",
			args{[][]int{
				{1, 4},
				{4, 5},
			}},
			[][]int{
				{1, 5},
			},
		},
		{
			"test-5",
			args{[][]int{
				{1, 4},
				{1, 4},
			}},
			[][]int{
				{1, 4},
			},
		},
		{
			"test-6",
			args{[][]int{
				{4, 5},
				{1, 4},
				{0, 1},
			}},
			[][]int{
				{0, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
