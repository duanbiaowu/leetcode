package leetcode

import "testing"

func Test_firstUniqChar(t *testing.T) {
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
			-1,
		},
		{
			"test-2",
			args{"aa"},
			-1,
		},
		{
			"test-3",
			args{"ab"},
			0,
		},
		{
			"test-4",
			args{"abb"},
			0,
		},
		{
			"test-5",
			args{"abac"},
			1,
		},
		{
			"test-6",
			args{"leetcode"},
			0,
		},
		{
			"test-7",
			args{"loveleetcode"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
