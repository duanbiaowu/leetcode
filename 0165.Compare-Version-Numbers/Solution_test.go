package leetcode

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{"", ""},
			0,
		},
		{
			"test-2",
			args{"1.01", "1.001"},
			0,
		},
		{
			"test-3",
			args{"0.1", "1.1"},
			-1,
		},
		{
			"test-4",
			args{"1.0.1", "1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
