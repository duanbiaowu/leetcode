package leetcode

import (
	"reflect"
	"testing"
)

func Test_pondSizes(t *testing.T) {
	type args struct {
		land [][]int
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
			args{[][]int{
				{0, 2, 1, 0},
				{0, 1, 0, 1},
				{1, 1, 0, 1},
				{0, 1, 0, 1},
			}},
			[]int{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pondSizes(tt.args.land); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pondSizes() = %v, want %v", got, tt.want)
			}
		})
	}
}
