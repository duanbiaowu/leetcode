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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
