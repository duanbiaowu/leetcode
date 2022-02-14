package leetcode

import (
	"reflect"
	"testing"
)

func Test_divingBoard(t *testing.T) {
	type args struct {
		shorter int
		longer  int
		k       int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{0, 0, 0},
			nil,
		},
		{
			"test-2",
			args{1, 1, 1},
			[]int{1},
		},
		{
			"test-3",
			args{1, 2, 3},
			[]int{3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divingBoard(tt.args.shorter, tt.args.longer, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divingBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
