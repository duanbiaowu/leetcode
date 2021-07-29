package leetcode

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test-1",
			args{0},
			[]string{},
		},
		{
			"test-2",
			args{1},
			[]string{"()"},
		},
		{
			"test-3",
			args{2},
			[]string{"(())", "()()"},
		},
		{
			"test-4",
			args{3},
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
