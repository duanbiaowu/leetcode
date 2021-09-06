package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_reorderList(t *testing.T) {
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
			args{&ListNode{Val: 1}},
			&ListNode{Val: 1},
		},
		{
			"test-3",
			args{structures.GenerateListNodesByArray([]int{1, 2})},
			structures.GenerateListNodesByArray([]int{1, 2}),
		},
		{
			"test-4",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4})},
			structures.GenerateListNodesByArray([]int{1, 4, 2, 3}),
		},
		{
			"test-5",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4, 5})},
			structures.GenerateListNodesByArray([]int{1, 5, 2, 4, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if reorderList(tt.args.head); !reflect.DeepEqual(tt.args.head, tt.want) {
				t.Errorf("reorderList() = %v, want %v", tt.args.head, tt.want)
			}
		})
	}
}
