package leetcode

import (
	"reflect"
	"testing"
)

func Test_missingTwo(t *testing.T) {
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
			[]int{1, 2},
		},
		{
			"test-2",
			args{[]int{1}},
			[]int{2, 3},
		},
		{
			"test-3",
			args{[]int{2, 4}},
			[]int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingTwo(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("missingTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
