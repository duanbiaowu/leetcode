package leetcode

import (
	"reflect"
	"testing"
)

func Test_pairSums(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{nil, 0},
			nil,
		},
		{
			"test-2",
			args{[]int{1}, 0},
			nil,
		},
		{
			"test-3",
			args{[]int{5, 6, 5}, 11},
			[][]int{
				{5, 6},
			},
		},
		{
			"test-4",
			args{[]int{5, 6, 5, 6}, 11},
			[][]int{
				{5, 6},
				{5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairSums2(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pairSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
