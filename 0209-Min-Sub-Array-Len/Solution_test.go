package leetcode

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{0, nil},
			0,
		},
		{
			"test-2",
			args{7, []int{2, 3, 1, 2, 4, 3}},
			2,
		},
		{
			"test-3",
			args{11, []int{1, 1, 1, 1, 1, 1, 1, 1}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
