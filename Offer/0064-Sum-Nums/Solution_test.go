package leetcode

import "testing"

func Test_sumNums(t *testing.T) {
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
			1,
		},
		{
			"test-3",
			args{10},
			55,
		},
		{
			"test-4",
			args{100},
			5050,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNums(tt.args.n); got != tt.want {
				t.Errorf("sumNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
