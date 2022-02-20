package leetcode

import (
	"reflect"
	"testing"
)

func Test_getValidT9Words(t *testing.T) {
	type args struct {
		num   string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test-1",
			args{"", nil},
			nil,
		},
		{
			"test-2",
			args{"8733", []string{"tree", "used"}},
			[]string{"tree", "used"},
		},
		{
			"test-3",
			args{"2", []string{"a", "b", "c", "d"}},
			[]string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValidT9Words(tt.args.num, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getValidT9Words() = %v, want %v", got, tt.want)
			}
		})
	}
}
