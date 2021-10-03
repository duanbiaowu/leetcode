package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
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
			args{&ListNode{Val: 0}},
			&ListNode{Val: 0},
		},
		{
			"test-3",
			args{&ListNode{Val: 1}},
			&ListNode{Val: 1},
		},
		{
			"test-4",
			args{&ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			}},
			&ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			},
		},
		{
			"test-5",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3})},
			structures.GenerateListNodesByArray([]int{1, 2, 3}),
		},
		{
			"test-6",
			args{structures.GenerateListNodesByArray([]int{1, 1, 2})},
			structures.GenerateListNodesByArray([]int{2}),
		},
		{
			"test-7",
			args{structures.GenerateListNodesByArray([]int{1, 2, 2, 3, 3, 4, 5})},
			structures.GenerateListNodesByArray([]int{1, 4, 5}),
		},
		{
			"test-8",
			args{structures.GenerateListNodesByArray([]int{1, 1, 1, 1, 1})},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicatesIteratively(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
