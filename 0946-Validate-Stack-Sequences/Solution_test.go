package leetcode

import "testing"

func Test_validateStackSequences(t *testing.T) {
	type args struct {
		pushed []int
		popped []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{nil, nil},
			false,
		},
		{
			"test-2",
			args{[]int{1}, nil},
			false,
		},
		{
			"test-3",
			args{nil, []int{1}},
			false,
		},
		{
			"test-4",
			args{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}},
			true,
		},
		{
			"test-5",
			args{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateStackSequences(tt.args.pushed, tt.args.popped); got != tt.want {
				t.Errorf("validateStackSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
