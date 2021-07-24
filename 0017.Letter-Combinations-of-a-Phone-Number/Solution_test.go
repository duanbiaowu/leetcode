package leetcode

import (
	"reflect"
	"testing"
)

type args struct {
	digits string
}

var tests = []struct {
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
	{
		"No3. test a string with multiple digits",
		args{"234"},
		[]string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
	},
}

func Test_backtrackLetterCombinations(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backtrackLetterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtrackLetterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfsLetterCombinations(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dfsLetterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsLetterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bfsLetterCombinations(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bfsLetterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bfsLetterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
