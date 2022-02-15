package leetcode

import (
	"reflect"
	"testing"
)

func Test_subSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil},
			[]int{-1, -1},
		},
		{
			"test-2",
			args{[]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}},
			[]int{3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
