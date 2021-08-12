package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{""},
			nil,
		},
		{
			"test-2",
			args{"1+2"},
			[]int{3},
		},
		{
			"test-3",
			args{"1+2+3"},
			[]int{6, 6},
		},
		{
			"test-4",
			args{"2-1-1"},
			[]int{0, 2},
		},
		{
			"test-5",
			args{"2*3-4*5"},
			[]int{-34, -14, -10, -10, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffWaysToCompute(tt.args.expression)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
