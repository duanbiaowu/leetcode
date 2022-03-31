package leetcode

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
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
			args{[]int{1}},
			[]int{1},
		},
		{
			"test-3",
			args{[]int{1, 2}},
			[]int{2, 1},
		},
		{
			"test-4",
			args{[]int{1, 2, 3, 4}},
			[]int{24, 12, 8, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
