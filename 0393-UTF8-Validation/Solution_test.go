package leetcode

import "testing"

func Test_validUtf8(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{[]int{}},
			false,
		},
		{
			"test-2",
			args{[]int{197, 130, 1}},
			true,
		},
		{
			"test-3",
			args{[]int{235, 140, 4}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validUtf8(tt.args.data); got != tt.want {
				t.Errorf("validUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}
