package leetcode

import "testing"

func Test_numberOf2sInRange(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{0},
			0,
		},
		{
			"test-2",
			args{1},
			0,
		},
		{
			"test-3",
			args{2},
			1,
		},
		{
			"test-4",
			args{3},
			1,
		},
		{
			"test-5",
			args{12},
			2,
		},
		{
			"test-6",
			args{25},
			9,
		},
		{
			"test-7",
			args{121},
			24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOf2sInRange(tt.args.n); got != tt.want {
				t.Errorf("numberOf2sInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
