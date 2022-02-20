package leetcode

import (
	"reflect"
	"testing"
)

func Test_findLongestSubarray(t *testing.T) {
	type args struct {
		array []string
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
			args{[]string{"A", "A"}},
			[]string{},
		},
		{
			"test-3",
			args{[]string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7", "H", "I", "J", "K", "L", "M"}},
			[]string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestSubarray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLongestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
