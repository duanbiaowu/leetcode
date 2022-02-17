package leetcode

import (
	"reflect"
	"testing"
)

func Test_wiggleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{[]int{1}},
			[]int{1},
		},
		{
			"test-3",
			args{[]int{5, 3, 1, 2, 3}},
			[]int{5, 1, 3, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("wiggleSort() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
