package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
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
			args{nil, nil},
			[]int{},
		},
		{
			"test-2",
			args{[]int{4, 1, 2}, []int{1, 3, 4, 2}},
			[]int{-1, 3, -1},
		},
		{
			"test-3",
			args{[]int{2, 4}, []int{1, 2, 3, 4}},
			[]int{3, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElementOpt(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
