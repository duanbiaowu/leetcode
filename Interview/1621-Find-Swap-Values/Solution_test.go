package leetcode

import (
	"reflect"
	"testing"
)

func Test_findSwapValues(t *testing.T) {
	type args struct {
		array1 []int
		array2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil, nil},
			nil,
		},
		{
			"test-2",
			args{[]int{1, 2, 3}, []int{4, 5, 6}},
			nil,
		},
		{
			"test-3",
			args{[]int{4, 1, 2, 1, 1, 2}, []int{3, 6, 3, 3}},
			[]int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSwapValues(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSwapValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
