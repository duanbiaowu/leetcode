package leetcode

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{0, 1},
			[][]int{},
		},
		{
			"test-2",
			args{1, 0},
			[][]int{},
		},
		{
			"test-3",
			args{1, 2},
			[][]int{},
		},
		{
			"test-4",
			args{4, 2},
			[][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
