package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_smallestK(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil, 1},
			nil,
		},
		{
			"test-2",
			args{[]int{1, 2, 3}, 0},
			nil,
		},
		{
			"test-3",
			args{[]int{1, 3, 5, 7, 2, 4, 6, 8}, 4},
			[]int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestK(tt.args.arr, tt.args.k)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestK() = %v, want %v", got, tt.want)
			}
		})
	}
}
