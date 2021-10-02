package leetcode

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{""},
			0,
		},
		{
			"test-2",
			args{"0"},
			0,
		},
		{
			"test-3",
			args{"12"},
			2,
		},
		{
			"test-4",
			args{"226"},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
