package leetcode

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{
				[]int{},
				0,
				[]int{},
				0,
			},
			[]int{},
		},
		{
			"test-2",
			args{
				[]int{1},
				1,
				[]int{},
				0,
			},
			[]int{1},
		},
		{
			"test-3",
			args{
				[]int{1, 0},
				1,
				[]int{1},
				1,
			},
			[]int{1, 1},
		},
		{
			"test-4",
			args{
				[]int{1, 0},
				1,
				[]int{2},
				1,
			},
			[]int{1, 2},
		},
		{
			"test-5",
			args{
				[]int{2, 0},
				1,
				[]int{1},
				1,
			},
			[]int{1, 2},
		},
		{
			"test-6",
			args{
				[]int{1, 3, 0, 0},
				2,
				[]int{2, 4},
				2,
			},
			[]int{1, 2, 3, 4},
		},
		{
			"test-7",
			args{
				[]int{1, 2, 3, 0, 0, 0},
				3,
				[]int{2, 5, 6},
				3,
			},
			[]int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("merge() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}
