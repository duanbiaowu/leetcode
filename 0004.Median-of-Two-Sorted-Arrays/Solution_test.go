package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"test two empty arrays",
			args{nums1: []int{}, nums2: []int{}},
			0,
		},
		{
			"test an empty array and a non-empty array",
			args{nums1: []int{}, nums2: []int{1}},
			1,
		},
		{
			"test an empty array and a non-empty array",
			args{nums1: []int{1}, nums2: []int{}},
			1,
		},
		{
			"test an empty containing more than one number",
			args{nums1: []int{1, 3}, nums2: []int{2}},
			2,
		},
		{
			"test an empty containing more than one number",
			args{nums1: []int{1, 2}, nums2: []int{3, 4}},
			2.5000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
