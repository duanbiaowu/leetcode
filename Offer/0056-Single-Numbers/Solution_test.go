package leetcode

import (
	"reflect"
	"testing"
)

func Test_singleNumbers(t *testing.T) {
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
			nil,
		},
		{
			"test-2",
			args{[]int{1, 1, 2, 2}},
			nil,
		},
		{
			"test-3",
			args{[]int{4, 1}},
			[]int{4, 1},
		},
		{
			"test-4",
			args{[]int{4, 1, 4, 6}},
			[]int{6, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
