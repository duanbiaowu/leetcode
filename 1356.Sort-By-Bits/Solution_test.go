package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortByBits(t *testing.T) {
	type args struct {
		arr []int
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
			args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
			[]int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			"test-3",
			args{[]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}},
			[]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByBitsOpts(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
