package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_twoOutOfThree(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil, nil, nil},
			nil,
		},
		{
			"test-2",
			args{[]int{1, 1, 3, 2}, []int{2, 3}, []int{3}},
			[]int{3, 2},
		},
		{
			"test-3",
			args{[]int{3, 1}, []int{2, 3}, []int{1, 2}},
			[]int{2, 3, 1},
		},
		{
			"test-4",
			args{[]int{1, 2, 2}, []int{4, 3, 3}, []int{5}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoOutOfThreeOpt(tt.args.nums1, tt.args.nums2, tt.args.nums3)

			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoOutOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
