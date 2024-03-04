package leetcode

import (
	"reflect"
	"testing"
)

func Test_finalPrices(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil},
			[]int{},
		},
		{
			"test-2",
			args{[]int{8, 4, 6, 2, 3}},
			[]int{4, 2, 4, 2, 3},
		},
		{
			"test-3",
			args{[]int{1, 2, 3, 4, 5}},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalPrices(tt.args.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("finalPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
