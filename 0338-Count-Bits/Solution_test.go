package leetcode

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{0},
			[]int{0},
		},
		{
			"test-2",
			args{1},
			[]int{0, 1},
		},
		{
			"test-3",
			args{2},
			[]int{0, 1, 1},
		},
		{
			"test-4",
			args{5},
			[]int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
