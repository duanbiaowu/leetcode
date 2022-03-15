package leetcode

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{"", ""},
			nil,
		},
		{
			"test-2",
			args{"abc", "d"},
			nil,
		},
		{
			"test-3",
			args{"", "abc"},
			nil,
		},
		{
			"test-4",
			args{"cbaebabacd", "abc"},
			[]int{0, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
