package leetcode

import (
	"reflect"
	"testing"
)

func Test_replaceNonCoprimes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{[]int{4, 6, 3, 2, 7, 6, 2}},
			[]int{12, 7, 6},
		},
		{
			"test-3",
			args{[]int{2, 2, 1, 1, 3, 3, 3}},
			[]int{2, 1, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceNonCoprimes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceNonCoprimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
