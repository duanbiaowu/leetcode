package leetcode

import "testing"

func Test_waysToChange(t *testing.T) {
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
			args{-1},
			0,
		},
		{
			"test-2",
			args{0},
			0,
		},
		{
			"test-3",
			args{1},
			1,
		},
		{
			"test-4",
			args{5},
			2,
		},
		{
			"test-5",
			args{10},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToChange(tt.args.n); got != tt.want {
				t.Errorf("waysToChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
