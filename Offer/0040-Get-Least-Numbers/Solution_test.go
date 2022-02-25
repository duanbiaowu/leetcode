package leetcode

import (
	"reflect"
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
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
			args{nil, 0},
			nil,
		},
		{
			"test-2",
			args{[]int{3, 2, 1}, 2},
			[]int{1, 2},
		},
		{
			"test-3",
			args{[]int{3, 2, 1}, 3},
			[]int{3, 2, 1},
		},
		{
			"test-4",
			args{[]int{3, 2, 1, 5, 6, -1, 0}, 3},
			[]int{-1, 0, 1},
		},
		{
			"test-5",
			args{[]int{0, 1, 2, 1}, 1},
			[]int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLeastNumbers(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
