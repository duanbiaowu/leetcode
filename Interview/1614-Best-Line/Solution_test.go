package leetcode

import (
	"reflect"
	"testing"
)

func Test_bestLine(t *testing.T) {
	type args struct {
		points [][]int
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
				{0, 0},
				{1, 1},
				{1, 0},
				{2, 0},
			}},
			[]int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestLine(tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bestLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
