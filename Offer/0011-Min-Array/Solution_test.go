package leetcode

import "testing"

func Test_minArray(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{[]int{3, 4, 5, 1, 2}},
			1,
		},
		{
			"test-2",
			args{[]int{2, 2, 2, 0, 1}},
			0,
		},
		{
			"test-3",
			args{[]int{}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray(tt.args.numbers); got != tt.want {
				t.Errorf("minArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
