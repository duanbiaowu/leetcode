package leetcode

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test a blank area",
			args{[]int{}},
			0,
		},
		{
			"No1. test an area without intervals",
			args{[]int{1}},
			0,
		},
		{
			"No2. test an area without intervals",
			args{[]int{100}},
			0,
		},

		{
			"No1. test an area with multiple intervals",
			args{[]int{1, 1}},
			1,
		},
		{
			"No2. test an area with multiple intervals",
			args{[]int{4, 3, 2, 1, 4}},
			16,
		},
		{
			"No3. test an area with multiple intervals",
			args{[]int{1, 2, 1}},
			2,
		},
		{
			"No4. test an area with multiple intervals",
			args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
