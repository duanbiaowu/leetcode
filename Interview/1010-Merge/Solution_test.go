package leetcode

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		A []int
		m int
		B []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{[]int{}, 0, []int{}, 0},
			[]int{},
		},
		{
			"test-2",
			args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			"test-3",
			args{[]int{1}, 1, []int{}, 0},
			[]int{1},
		},
		{
			"test-4",
			args{[]int{0}, 0, []int{1}, 1},
			[]int{1},
		},
		{
			"test-4",
			args{[]int{3, 0, 0}, 1, []int{1, 2}, 2},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.A, tt.args.m, tt.args.B, tt.args.n)
			if !reflect.DeepEqual(tt.args.A, tt.want) {
				t.Errorf("merge() = %v, want %v", tt.args.A, tt.want)
			}
		})
	}
}
