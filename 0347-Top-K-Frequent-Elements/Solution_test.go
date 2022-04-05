package leetcode

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil, 0},
			[]int{},
		},
		{
			"test-2",
			args{[]int{}, 1},
			[]int{},
		},
		{
			"test-3",
			args{[]int{1}, 1},
			[]int{1},
		},
		{
			"test-4",
			args{[]int{1, 1, 1, 2, 2, 3}, 2},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
