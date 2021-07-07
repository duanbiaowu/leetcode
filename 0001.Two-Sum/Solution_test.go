package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test empty array",
			args{[]int{}, 1},
			[]int{},
		},
		{
			"test an array with only one element",
			args{[]int{}, 1},
			[]int{},
		},
		{
			"test an array of two elements, but there is no answer",
			args{[]int{1, 2}, 5},
			[]int{},
		},
		{
			"test an array of two elements, and there is an answer",
			args{[]int{1, 2}, 3},
			[]int{0, 1},
		},
		{
			"test an array of multiple elements, but there is no answer",
			args{[]int{6, 7, 5, 4, 9, 0}, 1},
			[]int{},
		},
		{
			"test an array of multiple elements, and there is an answer",
			args{[]int{6, 7, 5, 4, 9, 0}, 10},
			[]int{0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
