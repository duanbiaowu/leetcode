package leetcode

import (
	"reflect"
	"testing"
)

func Test_permutation(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test-1",
			args{"ab"},
			[]string{"ab", "ba"},
		},
		{
			"test-2",
			args{"qqe"},
			[]string{"eqq", "qeq", "qqe"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutation(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
