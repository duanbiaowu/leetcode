package leetcode

import "testing"

func Test_maxAliveYear(t *testing.T) {
	type args struct {
		birth []int
		death []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{nil, nil},
			0,
		},
		{
			"test-2",
			args{[]int{1900, 1901, 1950}, []int{1948, 1951, 2000}},
			1901,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAliveYear(tt.args.birth, tt.args.death); got != tt.want {
				t.Errorf("maxAliveYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
