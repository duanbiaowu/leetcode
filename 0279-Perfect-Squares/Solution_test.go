package leetcode

import "testing"

func Test_numSquares(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	"test-1",
		//	args{0},
		//	0,
		//},
		//{
		//	"test-2",
		//	args{1},
		//	1,
		//},
		//{
		//	"test-3",
		//	args{2},
		//	2,
		//},
		//{
		//	"test-4",
		//	args{3},
		//	3,
		//},
		//{
		//	"test-5",
		//	args{4},
		//	1,
		//},
		{
			"test-6",
			args{12},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.args.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
