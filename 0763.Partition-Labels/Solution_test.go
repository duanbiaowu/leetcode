package leetcode

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []int
	}{
		{
			"test-1",
			"ababcbacadefegdehijhklij",
			[]int{9, 7, 8},
		},
		{
			"test-2",
			"eccbbbbdec",
			[]int{10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLabels(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
