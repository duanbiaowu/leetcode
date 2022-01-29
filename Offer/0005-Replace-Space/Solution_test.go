package leetcode

import "testing"

func Test_replaceSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{"We are happy."},
			"We%20are%20happy.",
		},
		{
			"test-2",
			args{""},
			"",
		},
		{
			"test-3",
			args{" "},
			"%20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace(tt.args.s); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
