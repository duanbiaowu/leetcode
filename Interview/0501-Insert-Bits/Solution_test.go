package leetcode

import "testing"

func Test_insertBits(t *testing.T) {
	type args struct {
		N int
		M int
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{1024, 19, 2, 6},
			1100,
		},
		{
			"test-2",
			args{0, 31, 0, 4},
			31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertBits(tt.args.N, tt.args.M, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("insertBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
