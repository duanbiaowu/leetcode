package leetcode

import "testing"

func Test_translateNum(t *testing.T) {
	type args struct {
		num int
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
			1,
		},
		{
			"test-3",
			args{9},
			1,
		},
		{
			"test-4",
			args{25},
			2,
		},
		{
			"test-5",
			args{56},
			1,
		},
		{
			"test-6",
			args{12258},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNum(tt.args.num); got != tt.want {
				t.Errorf("translateNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
