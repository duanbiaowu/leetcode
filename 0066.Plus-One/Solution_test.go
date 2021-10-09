package leetcode

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
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
			args{[]int{0}},
			[]int{1},
		},
		{
			"test-3",
			args{[]int{9, 9}},
			[]int{1, 0, 0},
		},
		{
			"test-4",
			args{[]int{1, 2, 3}},
			[]int{1, 2, 4},
		},
		{
			"test-5",
			args{[]int{4, 3, 2, 1}},
			[]int{4, 3, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
