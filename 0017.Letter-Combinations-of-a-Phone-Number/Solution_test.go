package leetcode

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test a empty string",
			args{""},
			[]string{},
		},

		{
			"No1. test a string with only one digit",
			args{"2"},
			[]string{"a", "b", "c"},
		},
		{
			"No2. test a string with only one digit",
			args{"5"},
			[]string{"j", "k", "l"},
		},

		{
			"No1. test a string with multiple digits",
			args{"23"},
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"No2. test a string with multiple digits",
			args{"56"},
			[]string{"jm", "jn", "jo", "km", "kn", "ko", "lm", "ln", "lo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
