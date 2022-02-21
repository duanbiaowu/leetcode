package leetcode

import (
	"math"
	"testing"
)

func Test_smallestDifference(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{nil, nil},
			math.MaxInt32,
		},
		{
			"test-2",
			args{[]int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestDifference(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("smallestDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
