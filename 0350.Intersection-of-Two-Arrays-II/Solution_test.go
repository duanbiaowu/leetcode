package leetcode

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
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
				[]int{},
			},
			nil,
		},
		{
			"test-2",
			args{
				[]int{1},
				[]int{},
			},
			nil,
		},
		{
			"test-3",
			args{
				[]int{},
				[]int{1},
			},
			nil,
		},
		{
			"test-4",
			args{
				[]int{1},
				[]int{2},
			},
			nil,
		},
		{
			"test-5",
			args{
				[]int{1},
				[]int{1},
			},
			[]int{1},
		},
		{
			"test-6",
			args{
				[]int{1, 2, 2, 1},
				[]int{2, 2},
			},
			[]int{2, 2},
		},
		{
			"test-7",
			args{
				[]int{4, 9, 5},
				[]int{9, 4, 9, 8, 4},
			},
			[]int{9, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
