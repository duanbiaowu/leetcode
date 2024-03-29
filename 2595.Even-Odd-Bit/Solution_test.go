package leetcode

import (
	"reflect"
	"testing"
)

func Test_evenOddBit(t *testing.T) {
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
			[]int{0, 0},
		},
		{
			"test-2",
			args{17},
			[]int{2, 0},
		},
		{
			"test-3",
			args{2},
			[]int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evenOddBit(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evenOddBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
