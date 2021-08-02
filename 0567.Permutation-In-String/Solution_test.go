package leetcode

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{"ab", "eidbaooo"},
			true,
		},
		{
			"test-2",
			args{"ab", "eidboaoo"},
			false,
		},
		{
			"test-3",
			args{"abcd", "abccdefg"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
