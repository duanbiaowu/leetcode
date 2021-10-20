package leetcode

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{nil, nil},
			nil,
		},
		{
			"test-2",
			args{[][]int{
				{1, 3},
				{6, 9},
			}, nil},
			[][]int{
				{1, 3},
				{6, 9},
			},
		},
		{
			"test-3",
			args{nil, []int{2, 5}},
			[][]int{
				{2, 5},
			},
		},
		{
			"test-4",
			args{[][]int{
				{1, 3},
				{6, 9},
			}, []int{2, 5}},
			[][]int{
				{1, 5},
				{6, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
