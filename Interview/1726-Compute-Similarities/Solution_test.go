package leetcode

import (
	"reflect"
	"testing"
)

func Test_computeSimilarities(t *testing.T) {
	type args struct {
		docs [][]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{[][]int{
				{14, 15, 100, 9, 3},
				{32, 1, 9, 3, 5},
				{15, 29, 2, 6, 8, 7},
				{7, 10},
			}},
			[]string{
				"0,1: 0.2500",
				"0,2: 0.1000",
				"2,3: 0.1429",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeSimilarities(tt.args.docs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeSimilarities() = %v, want %v", got, tt.want)
			}
		})
	}
}
