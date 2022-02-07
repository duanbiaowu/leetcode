package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_removeDuplicateNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{structures.GenerateListNodesByArray([]int{1})},
			structures.GenerateListNodesByArray([]int{1}),
		},
		{
			"test-3",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3})},
			structures.GenerateListNodesByArray([]int{1, 2, 3}),
		},
		{
			"test-4",
			args{structures.GenerateListNodesByArray([]int{1, 1, 2, 3, 3})},
			structures.GenerateListNodesByArray([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicateNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
