package leetcode

import (
	"reflect"
	"testing"
)

func Test_multiSearch(t *testing.T) {
	type args struct {
		big    string
		smalls []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{"", nil},
			nil,
		},
		{
			"test-2",
			args{"", []string{}},
			nil,
		},
		{
			"test-3",
			args{"a", nil},
			nil,
		},
		{
			"test-4",
			args{"mississippi", []string{"is", "ppi", "hi", "sis", "i", "ssippi"}},
			[][]int{
				{1, 4},
				{8},
				{},
				{3},
				{1, 4, 7, 10},
				{5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiSearch(tt.args.big, tt.args.smalls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multiSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
