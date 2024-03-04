package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
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
			[]int{},
		},
		{
			"test-2",
			args{[]int{1, 2, 1}},
			[]int{2, -1, 2},
		},
		{
			"test-3",
			args{[]int{1, 2, 3, 4, 3}},
			[]int{2, 3, 4, -1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElementsOpt(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
