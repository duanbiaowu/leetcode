package leetcode

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			"test-1",
			args{"abaccdeff"},
			'b',
		},
		{
			"test-2",
			args{""},
			' ',
		},
		{
			"test-3",
			args{"leetcode"},
			'l',
		},
		{
			"test-4",
			args{"aadadaad"},
			' ',
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
