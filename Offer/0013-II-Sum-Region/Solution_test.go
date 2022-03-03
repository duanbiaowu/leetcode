package leetcode

import "testing"

func TestNumMatrix_SumRegion(t *testing.T) {
	type fields struct {
		sums [][]int
	}
	type args struct {
		row1 int
		col1 int
		row2 int
		col2 int
	}

	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	tests := []struct {
		name   string
		martix [][]int
		args   args
		want   int
	}{
		{
			"test-1",
			matrix,
			args{-1, -1, -1, -1},
			0,
		},
		{
			"test-2",
			matrix,
			args{2, 1, 4, 3},
			8,
		},
		{
			"test-3",
			matrix,
			args{1, 1, 2, 2},
			11,
		},
		{
			"test-4",
			matrix,
			args{1, 2, 2, 4},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.martix)
			if got := this.SumRegion(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2); got != tt.want {
				t.Errorf("SumRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}
