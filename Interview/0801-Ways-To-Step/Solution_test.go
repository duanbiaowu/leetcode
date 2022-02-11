package leetcode

import "testing"

func Test_waysToStep(t *testing.T) {
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
			args{2},
			2,
		},
		{
			"test-4",
			args{3},
			4,
		},
		{
			"test-5",
			args{5},
			13,
		},
		{
			"test-6",
			args{61},
			752119970,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToStep(tt.args.n); got != tt.want {
				t.Errorf("waysToStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
