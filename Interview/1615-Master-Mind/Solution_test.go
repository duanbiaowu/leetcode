package leetcode

import (
	"reflect"
	"testing"
)

func Test_masterMind(t *testing.T) {
	type args struct {
		solution string
		guess    string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{"", ""},
			[]int{0, 0},
		},
		{
			"test-2",
			args{"RGBY", "GGRR"},
			[]int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := masterMind(tt.args.solution, tt.args.guess); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("masterMind() = %v, want %v", got, tt.want)
			}
		})
	}
}
