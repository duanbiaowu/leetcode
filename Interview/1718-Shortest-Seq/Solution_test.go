package leetcode

import (
	"reflect"
	"testing"
)

func Test_shortestSeq(t *testing.T) {
	type args struct {
		big   []int
		small []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil, nil},
			nil,
		},
		{
			"test-2",
			args{[]int{1, 2, 3}, []int{4}},
			nil,
		},
		{
			"test-3",
			args{nil, []int{1, 2, 3, 4, 5}},
			nil,
		},
		{
			"test-4",
			args{[]int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7}, []int{1, 5, 9}},
			[]int{7, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestSeq(tt.args.big, tt.args.small); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}
