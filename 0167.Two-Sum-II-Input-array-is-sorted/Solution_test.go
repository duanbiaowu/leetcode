package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{[]int{}, 0},
			[]int{},
		},
		{
			"test-2",
			args{[]int{1}, 0},
			[]int{},
		},
		{
			"test-3",
			args{[]int{-1, 0}, -1},
			[]int{1, 2},
		},
		{
			"test-4",
			args{[]int{2, 3, 4}, 6},
			[]int{1, 3},
		},
		{
			"test-5",
			args{[]int{2, 7, 11, 15}, 9},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
