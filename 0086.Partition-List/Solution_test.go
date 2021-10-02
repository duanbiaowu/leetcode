package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test-1",
			args{nil, 1},
			nil,
		},
		{
			"test-2",
			args{structures.GenerateListNodesByArray([]int{1, 2}), 3},
			structures.GenerateListNodesByArray([]int{1, 2}),
		},
		{
			"test-3",
			args{structures.GenerateListNodesByArray([]int{3, 1}), 2},
			structures.GenerateListNodesByArray([]int{1, 3}),
		},
		{
			"test-4",
			args{structures.GenerateListNodesByArray([]int{1, 1, 1}), 1},
			structures.GenerateListNodesByArray([]int{1, 1, 1}),
		},
		{
			"test-5",
			args{structures.GenerateListNodesByArray([]int{1, 4, 3, 2, 5, 2}), 3},
			structures.GenerateListNodesByArray([]int{1, 2, 2, 4, 3, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
