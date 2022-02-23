package leetcode

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"test-1",
			args{0},
			nil,
		},
		{
			"test-2",
			args{4},
			[][]string{
				{
					".Q..", // 解法 1
					"...Q",
					"Q...",
					"..Q.",
				},
				{
					"..Q.", // 解法 2
					"Q...",
					"...Q",
					".Q..",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
