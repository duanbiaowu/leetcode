package leetcode

import "testing"

func Test_fourSumCount(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{nil, nil, nil, nil},
			0,
		},
		{
			"test-2",
			args{[]int{1}, []int{2}, []int{3}, []int{4}},
			0,
		},
		{
			"test-3",
			args{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumCount(tt.args.nums1, tt.args.nums2, tt.args.nums3, tt.args.nums4); got != tt.want {
				t.Errorf("fourSumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
