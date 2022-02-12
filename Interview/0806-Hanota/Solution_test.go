package leetcode

import (
	"reflect"
	"testing"
)

func Test_hanota(t *testing.T) {
	type args struct {
		A []int
		B []int
		C []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil, nil, nil},
			nil,
		},
		{
			"test-2",
			args{[]int{2, 1, 0}, []int{}, []int{}},
			[]int{2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hanota(tt.args.A, tt.args.B, tt.args.C); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hanota() = %v, want %v", got, tt.want)
			}
		})
	}
}
