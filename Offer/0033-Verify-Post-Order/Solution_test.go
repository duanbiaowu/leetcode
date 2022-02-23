package leetcode

import "testing"

func Test_verifyPostorder(t *testing.T) {
	type args struct {
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{nil},
			true,
		},
		{
			"test-2",
			args{[]int{1}},
			true,
		},
		{
			"test-3",
			args{[]int{1, 6, 3, 2, 5}},
			false,
		},
		{
			"test-4",
			args{[]int{1, 3, 2, 6, 5}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPostorder(tt.args.postorder); got != tt.want {
				t.Errorf("verifyPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
