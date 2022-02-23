package leetcode

import (
	"reflect"
	"testing"
)

func Test_printNumbers(t *testing.T) {
	type args struct {
		n int
	}

	N2 := make([]int, 99)
	for i := range N2 {
		N2[i] = i + 1
	}

	N3 := make([]int, 999)
	for i := range N3 {
		N3[i] = i + 1
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{0},
			nil,
		},
		{
			"test-2",
			args{1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"test-3",
			args{2},
			N2,
		},
		{
			"test-4",
			args{3},
			N3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printNumbers(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
